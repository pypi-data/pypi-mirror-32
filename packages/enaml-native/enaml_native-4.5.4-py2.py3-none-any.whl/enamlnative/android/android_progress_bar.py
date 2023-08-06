"""
Copyright (c) 2017, Jairus Martin.

Distributed under the terms of the MIT License.

The full license is in the file LICENSE, distributed with this software.

Created on May 26, 2017

@author: jrm
"""
from atom.api import Typed, Bool, set_default
from enamlnative.widgets.progress_bar import ProxyProgressBar

from .android_view import AndroidView, View
from .bridge import JavaMethod


class ProgressBar(View):
    __nativeclass__ = set_default('android.widget.ProgressBar')
    __signature__ = set_default(('android.content.Context',
                                 'android.util.AttributeSet',
                                 'android.R'))
    setIndeterminate = JavaMethod('boolean')
    setMax = JavaMethod('int')
    setMin = JavaMethod('int')
    setProgress = JavaMethod('int')#, 'boolean')
    setSecondaryProgress = JavaMethod('int')

    STYLE_HORIZONTAL = '@attr/progressBarStyleHorizontal'
    STYLE_INVERSE = '@attr/progressBarStyleInverse'
    STYLE_LARGE = '@attr/progressBarStyleLarge'
    STYLE_SMALL = '@attr/progressBarStyleSmall'
    STYLE_NORMAL = '@attr/progressBarStyle'

    STYLES = {
        'normal': STYLE_NORMAL,
        'small': STYLE_SMALL,
        'large': STYLE_LARGE,
    }


class AndroidProgressBar(AndroidView, ProxyProgressBar):
    """ An Android implementation of an Enaml ProxyProgressBar.

    For an indeterminate ProgressBar use the ActivityIndicator.

    """
    #: A reference to the widget created by the proxy.
    widget = Typed(ProgressBar)

    #: Set to True to make the progress bar an activity indicator
    indeterminate = Bool()

    # -------------------------------------------------------------------------
    # Initialization API
    # -------------------------------------------------------------------------
    def create_widget(self):
        """ Create the underlying widget.

        """
        d = self.declaration

        if self.indeterminate:
            #: Note: Style will only exist on activity indicators!
            style = ProgressBar.STYLES[d.size]
        else:
            style = ProgressBar.STYLE_HORIZONTAL
        self.widget = ProgressBar(self.get_context(), None, style)

    def init_widget(self):
        """ Initialize the underlying widget.

        """
        super(AndroidProgressBar, self).init_widget()
        d = self.declaration
        self.set_indeterminate(self.indeterminate)

        if not self.indeterminate:
            if d.max:
                self.set_max(d.max)
            if d.min:
                self.set_min(d.min)

            self.set_progress(d.progress)

            if d.secondary_progress:
                self.set_secondary_progress(d.secondary_progress)

    # -------------------------------------------------------------------------
    # ProxyProgressBar API
    # -------------------------------------------------------------------------
    def set_progress(self, progress):
        self.widget.setProgress(progress)

    def set_indeterminate(self, indeterminate):
        self.widget.setIndeterminate(indeterminate)

    def set_secondary_progress(self, progress):
        self.widget.setSecondaryProgress(progress)

    def set_max(self, value):
        self.widget.setMax(value)

    def set_min(self, value):
        self.widget.setMin(value)

    def set_size(self, size):
        """ Size cannot be changed dynamically. """
        pass
