"""
Copyright (c) 2017, Jairus Martin.

Distributed under the terms of the MIT License.

The full license is in the file LICENSE, distributed with this software.

Created on May 20, 2017

@author: jrm
"""
from atom.api import Typed, set_default

from enamlnative.widgets.edit_text import ProxyEditText

from .android_text_view import AndroidTextView, TextView
from .bridge import JavaMethod


class EditText(TextView):
    __nativeclass__ = set_default('android.widget.EditText')
    setSelection = JavaMethod('int', 'int')
    selectAll = JavaMethod()
    extendSelection = JavaMethod('int')
    setHint = JavaMethod('java.lang.CharSequence')


class AndroidEditText(AndroidTextView, ProxyEditText):
    """ An Android implementation of an Enaml ProxyEditText.

    """
    #: A reference to the widget created by the proxy.
    widget = Typed(EditText)

    # -------------------------------------------------------------------------
    # Initialization API
    # -------------------------------------------------------------------------
    def create_widget(self):
        """ Create the underlying widget.

        """
        d = self.declaration
        self.widget = EditText(self.get_context(), None,
                               d.style or "@attr/editTextStyle")

    # -------------------------------------------------------------------------
    # ProxyEditText API
    # -------------------------------------------------------------------------
    def set_selection(self, selection):
        self.widget.setSelection(*selection)

    def set_placeholder(self, placeholder):
        self.widget.setHint(placeholder)


