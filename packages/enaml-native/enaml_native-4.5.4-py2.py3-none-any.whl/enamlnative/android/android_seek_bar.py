"""
Copyright (c) 2017, Jairus Martin.

Distributed under the terms of the MIT License.

The full license is in the file LICENSE, distributed with this software.

Created on June 7, 2017

@author: jrm
"""
from atom.api import Typed, set_default

from enamlnative.widgets.seek_bar import ProxySeekBar

from .android_progress_bar import AndroidProgressBar, ProgressBar
from .bridge import JavaMethod, JavaCallback


class SeekBar(ProgressBar):
    __nativeclass__ = set_default('android.widget.SeekBar')
    __signature__ = set_default(('android.content.Context',
                                 'android.util.AttributeSet', 'android.R'))
    setSplitTrack = JavaMethod('boolean')
    setOnSeekBarChangeListener = JavaMethod(
        'android.widget.SeekBar$OnSeekBarChangeListener')
    setKeyProgressIncrement = JavaMethod('int')

    onProgressChanged = JavaCallback('android.widget.SeekBar', 'int',
                                     'boolean')
    onStartTrackingTouch = JavaCallback('android.widget.SeekBar')
    onStopTrackingTouch = JavaCallback('android.widget.SeekBar')


class AndroidSeekBar(AndroidProgressBar, ProxySeekBar):
    """ An Android implementation of an Enaml ProxySeekBar.

    """
    #: A reference to the widget created by the proxy.
    widget = Typed(SeekBar)

    default_layout = set_default({
        'width': 'wrap_content',
        'height': 'wrap_content'
    })

    # -------------------------------------------------------------------------
    # Initialization API
    # -------------------------------------------------------------------------
    def create_widget(self):
        """ Create the underlying widget.

        """
        d = self.declaration
        self.widget = SeekBar(self.get_context(), None,
                              d.style or '@attr/seekBarStyle')

    def init_widget(self):
        """ Initialize the underlying widget.

        """
        super(AndroidSeekBar, self).init_widget()
        w = self.widget

        #: Setup listener
        w.setOnSeekBarChangeListener(w.getId())
        w.onProgressChanged.connect(self.on_progress_changed)

    # -------------------------------------------------------------------------
    # OnSeekBarChangeListener API
    # -------------------------------------------------------------------------
    def on_progress_changed(self, bar, progress, user):
        d = self.declaration
        with self.widget.setProgress.suppressed():
            d.progress = progress

    # -------------------------------------------------------------------------
    # ProxySeekBar API
    # -------------------------------------------------------------------------
    def set_key_progress_increment(self, value):
        self.widget.setKeyProgressIncrement(value)

    def set_split_track(self, split):
        self.widget.setSplitTrack(split)