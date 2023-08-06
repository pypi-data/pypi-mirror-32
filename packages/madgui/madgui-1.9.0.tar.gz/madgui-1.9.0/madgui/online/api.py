"""
This module defines the API of any online control plugin. Note that the API
is subject to change (as is most parts of madgui…).

The interface contract is currently designed as follows:

    - A subclass of :class:`PluginLoader` is registered under the
      "madgui.online.PluginLoader" entry point.

    - It loads the DLL / connects the database when requested and returns a
      :class:`OnlinePlugin` instance.

    - An :class:`OnlinePlugin` instance is used to instanciate accessors for
      the actual elements.

    - There are two kinds of accessors (returned as tuple):

        - :class:`ElementBackend` performs the actual database I/O, i.e.
          reads/writes parameters from the database.

        - :class:`ElementBackendConverter` performs parameter conversions
          between internal and standard representation
"""

from abc import ABCMeta, abstractmethod

from madgui.core.unit import from_ui, to_ui

_Interface = ABCMeta('_Interface', (object,), {})


class PluginLoader(_Interface):

    """Loader interface for online control plugin."""

    @classmethod
    def check_avail(self):
        """Check if the plugin is available."""
        return True

    @classmethod
    def load(self, frame):
        """Get a :class:`OnlinePlugin` instance."""
        raise NotImplementedError


class OnlinePlugin(_Interface):

    """Interface for a connected online control plugin."""

    @abstractmethod
    def connect(self):
        """Connect the online plugin to the control system."""

    @abstractmethod
    def disconnect(self):
        """Unload the online plugin, free resources."""

    @abstractmethod
    def execute(self):
        """Commit transaction."""

    @abstractmethod
    def param_info(self, knob):
        """Get parameter info for backend key."""

    @abstractmethod
    def read_monitor(self, name):
        """
        Read out one monitor, return values as dict with keys:

            widthx:     Beam x width
            widthy:     Beam y width
            posx:       Beam x position
            posy:       Beam y position
        """

    @abstractmethod
    def get_knob(self, mad_knob):
        """Return a :class:`Knob` belonging to the given attribute."""

    @abstractmethod
    def read_param(self, param):
        """Read parameter. Return numeric value."""

    @abstractmethod
    def write_param(self, param, value):
        """Update parameter into control system."""

    @abstractmethod
    def get_beam(self):
        """
        Return a dict ``{name: value}`` for all beam properties, in MAD-X
        units. At least: particle, mass, charge, energy
        """


class Knob:

    """Base class for knobs."""

    def __init__(self, plug, elem, attr, param, unit):
        self.plug = plug
        self.elem = elem
        self.el_name = elem.node_name.lower()
        self.attr = attr.lower()
        self.param = param
        self.unit = unit

    def __str__(self):
        return self.param

    def __repr__(self):
        return "{}({})".format(self.__class__.__name__, self)

    # mixins:

    def read(self):
        """Read element attribute."""
        return from_ui(self.attr, self.unit, self.plug.read_param(self.param))

    def write(self, value):
        """Update element attribute into control system."""
        self.plug.write_param(self.param, to_ui(self.unit, self.attr, value))
