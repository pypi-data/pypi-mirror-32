############################################################
# FlatCAM: 2D Post-processing for Manufacturing            #
# http://flatcam.org                                       #
# Author: Juan Pablo Caram (c)                             #
# Date: 2/5/2014                                           #
# MIT Licence                                              #
############################################################

import os
from pathlib import Path
from flatcam.FlatCAMObj import *
from flatcam import FlatCAMApp
from PyQt5 import QtWidgets, Qt, QtGui, QtCore


class KeySensitiveListView(QtWidgets.QListView):
    """
    QtGui.QListView extended to emit a signal on key press.
    """

    keyPressed = QtCore.pyqtSignal(int)

    def keyPressEvent(self, event):
        super(KeySensitiveListView, self).keyPressEvent(event)
        self.keyPressed.emit(event.key())


#class ObjectCollection(QtCore.QAbstractListModel):
class ObjectCollection():
    """
    Object storage and management.
    """

    classdict = {
        "gerber": FlatCAMGerber,
        "excellon": FlatCAMExcellon,
        "cncjob": FlatCAMCNCjob,
        "geometry": FlatCAMGeometry
    }

    icon_files = {
        "gerber": str( ICON_PATH /  'flatcam_icon16.png' ),
        "excellon": str( ICON_PATH /  'drill16.png' ),
        "cncjob": str( ICON_PATH /  'cnc16.png' ),
        "geometry": str( ICON_PATH /  'geometry16.png' )
    }

    def __init__(self, parent=None):
        ### Icons for the list view
        self.icons = {}
        for kind in ObjectCollection.icon_files:
            self.icons[kind] = QtGui.QPixmap(ObjectCollection.icon_files[kind])

        ### Data ###
        self.object_list = []
        self.checked_indexes = []

        # Names of objects that are expected to become available.
        # For example, when the creation of a new object will run
        # in the background and will complete some time in the
        # future. This is a way to reserve the name and to let other
        # tasks know that they have to wait until available.
        self.promises = set()

        ### View
        self.view = KeySensitiveListView()
        self.view.setSelectionMode(Qt.QAbstractItemView.ExtendedSelection)
        self.model = QtGui.QStandardItemModel(self.view)
        self.view.setModel(self.model)
        self.model.itemChanged.connect(self.on_item_changed)

        self.click_modifier = None

        ## GUI Events
        self.view.selectionModel().selectionChanged.connect(self.on_list_selection_change)
        self.view.activated.connect(self.on_item_activated)
        self.view.keyPressed.connect(self.on_key)
        self.view.clicked.connect(self.on_mouse_down)

    def promise(self, obj_name):
        FlatCAMApp.App.log.debug("Object %s has been promised." % obj_name)
        self.promises.add(obj_name)

    def has_promises(self):
        return len(self.promises) > 0

    def on_key(self, key):
        # Delete
        if key == QtCore.Qt.Key_Delete:
            # Delete via the application to
            # ensure cleanup of the GUI
            self.get_active().app.on_delete()
            return

        if key == QtCore.Qt.Key_Space:
            self.get_active().ui.plot_cb.toggle()
            return

    def print_list(self):
        for obj in self.object_list:
            print(obj)

    def on_mouse_down(self, event):
        FlatCAMApp.App.log.debug("Mouse button pressed on list")
        #self.print_list()

    def rowCount(self, parent=QtCore.QModelIndex(), *args, **kwargs):
        return len(self.object_list)

    def columnCount(self, *args, **kwargs):
        return 1

    def data(self, index, role=Qt.Qt.DisplayRole):
        if not index.isValid() or not 0 <= index.row() < self.rowCount():
            return QtCore.QVariant()
        row = index.row()
        if role == Qt.Qt.DisplayRole:
            return self.object_list[row].options["name"]
        if role == Qt.Qt.DecorationRole:
            return self.icons[self.object_list[row].kind]

    def append(self, obj, active=False):
        FlatCAMApp.App.log.debug(str(inspect.stack()[1][3]) + " --> OC.append()")

        name = obj.options["name"]

        # Check promises and clear if exists
        if name in self.promises:
            self.promises.remove(name)
            FlatCAMApp.App.log.debug("Promised object %s became available." % name)
            FlatCAMApp.App.log.debug("%d promised objects remaining." % len(self.promises))

        # Prevent same name
        while name in self.get_names():
            ## Create a new name
            # Ends with number?
            FlatCAMApp.App.log.debug("new_object(): Object name (%s) exists, changing." % name)
            match = re.search(r'(.*[^\d])?(\d+)$', name)
            if match:  # Yes: Increment the number!
                base = match.group(1) or ''
                num = int(match.group(2))
                name = base + str(num + 1)
            else:  # No: add a number!
                name += "_1"
        obj.options["name"] = name

        obj.set_ui(obj.ui_type())

        # Required before appending (Qt MVC)
        #self.beginInsertRows(QtCore.QModelIndex(), len(self.object_list), len(self.object_list))

        # Simply append to the python list
        self.object_list.append(obj)

        # Create the model item to insert into the QListView
        icon = QtGui.QIcon(self.icons[obj.kind])#self.icons["gerber"])
        item = QtGui.QStandardItem(icon, str(name))
        # Item is not editable, so that double click
        # does not allow cell value modification.
        item.setEditable(False)
        # The item is checkable, to add the checkbox.
        item.setCheckable(True)
        if obj.options["plot"] is True:
            item.setCheckState(2)   #Qt.Checked)
        else:
            item.setCheckState(0)   #Qt.Unchecked)

        self.model.appendRow(item)

        obj.option_changed.connect(self.on_object_option_changed)

        # Required after appending (Qt MVC)
        #self.endInsertRows()

    def on_object_option_changed(self, obj, key):
        if key == "plot":
            self.model.blockSignals(True)
            name = obj.options["name"]
            state = 0 #Qt.Unchecked
            for index in range(self.model.rowCount()):
                item = self.model.item(index)
                if self.object_list[item.row()].options["name"] == name:
                    if obj.options["plot"] == True:
                        state = 2 #Qt.Checked

                    item.setCheckState(state)
                    obj.ui.plot_cb.set_value(state)
                    break
            self.model.blockSignals(False)

    def get_names(self):
        """
        Gets a list of the names of all objects in the collection.

        :return: List of names.
        :rtype: list
        """

        FlatCAMApp.App.log.debug(str(inspect.stack()[1][3]) + " --> OC.get_names()")
        return [x.options['name'] for x in self.object_list]

    def get_bounds(self):
        """
        Finds coordinates bounding all objects in the collection.

        :return: [xmin, ymin, xmax, ymax]
        :rtype: list
        """
        FlatCAMApp.App.log.debug(str(inspect.stack()[1][3]) + "--> OC.get_bounds()")

        # TODO: Move the operation out of here.

        xmin = Inf
        ymin = Inf
        xmax = -Inf
        ymax = -Inf

        for obj in self.object_list:
            try:
                gxmin, gymin, gxmax, gymax = obj.bounds()
                xmin = min([xmin, gxmin])
                ymin = min([ymin, gymin])
                xmax = max([xmax, gxmax])
                ymax = max([ymax, gymax])
            except:
                FlatCAMApp.App.log.warning("DEV WARNING: Tried to get bounds of empty geometry.")

        return [xmin, ymin, xmax, ymax]

    def get_by_name(self, name):
        """
        Fetches the FlatCAMObj with the given `name`.

        :param name: The name of the object.
        :type name: str
        :return: The requested object or None if no such object.
        :rtype: FlatCAMObj or None
        """
        FlatCAMApp.App.log.debug(str(inspect.stack()[1][3]) + "--> OC.get_by_name()")

        for obj in self.object_list:
            if obj.options['name'] == name:
                return obj
        return None

    def delete_active(self):
        selections = self.view.selectedIndexes()
        if len(selections) == 0:
            return
        row = selections[0].row()

        #self.beginRemoveRows(QtCore.QModelIndex(), row, row)

        self.object_list.pop(row)
        self.model.removeRow(row)

        #self.endRemoveRows()

    def get_active(self):
        """
        Returns the active object or None

        :return: FlatCAMObj or None
        """
        selections = self.view.selectedIndexes()
        if len(selections) == 0:
            return None
        row = selections[0].row()
        return self.object_list[row]

    def get_selected(self):
        """
        Returns list of objects selected in the view.

        :return: List of objects
        """
        return [self.object_list[sel.row()] for sel in self.view.selectedIndexes()]

    def set_active(self, name):
        """
        Selects object by name from the project list. This triggers the
        list_selection_changed event and call on_list_selection_changed.

        :param name: Name of the FlatCAM Object
        :return: None
        """
        iobj = self.model.createIndex(self.get_names().index(name), 0)  # Column 0
        self.view.selectionModel().select(iobj, QtCore.QItemSelectionModel.Select)

    def set_inactive(self, name):
        """
        Unselect object by name from the project list. This triggers the
        list_selection_changed event and call on_list_selection_changed.

        :param name: Name of the FlatCAM Object
        :return: None
        """
        iobj = self.model.createIndex(self.get_names().index(name), 0)  # Column 0
        self.view.selectionModel().select(iobj, QtCore.QItemSelectionModel.Deselect)

    def set_all_inactive(self):
        """
        Unselect all objects from the project list. This triggers the
        list_selection_changed event and call on_list_selection_changed.

        :return: None
        """
        for name in self.get_names():
            self.set_inactive(name)

    def on_list_selection_change(self, current, previous):
        FlatCAMApp.App.log.debug("on_list_selection_change()")
        FlatCAMApp.App.log.debug("Current: %s, Previous %s" % (str(current), str(previous)))
        try:
            selection_index = current.indexes()[0].row()
        except IndexError:
            FlatCAMApp.App.log.debug("on_list_selection_change(): Index Error (Nothing selected?)")
            return

        self.object_list[selection_index].build_ui()

    def on_item_changed(self, item):
        FlatCAMApp.App.log.debug("on_item_changed(): " + str(item.row()) + " " + self.object_list[item.row()].options["name"])
        if item.checkState() == QtCore.Qt.Checked:
           self.object_list[item.row()].options["plot"] = True #(item.checkState() == QtCore.Qt.Checked)
        else:
           self.object_list[item.row()].options["plot"] = False #(item.checkState() == QtCore.Qt.Checked)

        self.object_list[item.row()].plot()
        return

    def on_item_activated(self, index):
        """
        Double-click or Enter on item.

        :param index: Index of the item in the list.
        :return: None
        """
        self.object_list[index.row()].build_ui()

    def delete_all(self):
        FlatCAMApp.App.log.debug(str(inspect.stack()[1][3]) + "--> OC.delete_all()")

#        self.beginResetModel()

        self.model.removeRows(0, self.model.rowCount())
        self.object_list = []
        self.checked_indexes = []

#        self.endResetModel()

    def get_list(self):
        return self.object_list

