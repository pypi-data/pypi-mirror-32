"""
Copyright (c) 2017, Jairus Martin.

Distributed under the terms of the MIT License.

The full license is in the file LICENSE, distributed with this software.

Created on May 20, 2017

@author: jrm
"""
from atom.api import Typed, set_default

from enamlnative.widgets.radio_group import ProxyRadioGroup

from .android_linear_layout import AndroidLinearLayout, LinearLayout
from .bridge import JavaCallback, JavaMethod


class RadioGroup(LinearLayout):
    __nativeclass__ = set_default('android.widget.RadioGroup')
    __signature__ = set_default(('android.content.Context',))

    check = JavaMethod('int')
    clearCheck = JavaMethod()
    setOnCheckedChangeListener = JavaMethod(
        'android.widget.RadioGroup$OnCheckedChangeListener')
    onCheckedChanged = JavaCallback('android.widget.RadioGroup', 'int')


class AndroidRadioGroup(AndroidLinearLayout, ProxyRadioGroup):
    """ An Android implementation of an Enaml ProxyLinearLayout.

    """
    #: A reference to the widget created by the proxy.
    widget = Typed(RadioGroup)

    # -------------------------------------------------------------------------
    # Initialization API
    # -------------------------------------------------------------------------
    def create_widget(self):
        """ Create the underlying widget.

        """
        self.widget = RadioGroup(self.get_context())

    def init_layout(self):
        """ Set the checked state after all children have
        been populated.
        
        """
        super(AndroidRadioGroup, self).init_layout()
        d = self.declaration
        w = self.widget
        if d.checked:
            self.set_checked(d.checked)
        else:
            #: Check if any of the children have "checked = True"
            for c in d.children:
                if c.checked:
                    d.checked = c

        w.setOnCheckedChangeListener(w.getId())
        w.onCheckedChanged.connect(self.on_checked_changed)

    # -------------------------------------------------------------------------
    # OnCheckedChangeListener API
    # -------------------------------------------------------------------------
    def on_checked_changed(self, group, checked_id):
        """ Set the checked property based on the checked state
        of all the children
        
        """
        d = self.declaration
        if checked_id < 0:
            with self.widget.clearCheck.suppressed():
                d.checked = None
            return
        else:
            for c in self.children():
                if c.widget.getId() == checked_id:
                    with self.widget.check.suppressed():
                        d.checked = c.declaration
                    return

    # -------------------------------------------------------------------------
    # ProxyRadioGroup API
    # -------------------------------------------------------------------------
    def set_checked(self, checked):
        """ Properly check the correct radio button.

        """
        if not checked:
            self.widget.clearCheck()
        else:
            #: Checked is a reference to the radio declaration
            #: so we need to get the ID of it
            rb = checked.proxy.widget
            if not rb:
                return
            self.widget.check(rb.getId())
