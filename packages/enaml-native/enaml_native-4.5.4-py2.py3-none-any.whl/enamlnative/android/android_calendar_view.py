"""
Copyright (c) 2017, Jairus Martin.

Distributed under the terms of the MIT License.

The full license is in the file LICENSE, distributed with this software.

Created on May 20, 2017

@author: jrm
"""
from atom.api import Typed, set_default

from datetime import datetime
from enamlnative.widgets.calendar_view import ProxyCalendarView

from .android_frame_layout import AndroidFrameLayout, FrameLayout
from .bridge import JavaMethod, JavaCallback


class CalendarView(FrameLayout):
    __nativeclass__ = set_default('android.widget.CalendarView')
    setDate = JavaMethod('long')
    setMinDate = JavaMethod('long')
    setMaxDate = JavaMethod('long')
    setFirstDayOfWeek = JavaMethod('int')
    setOnDateChangeListener = JavaMethod(
        'android.widget.CalendarView$OnDateChangeListener')

    #: This is not actually a CalendarView method, but still works
    onSelectedDayChange = JavaCallback('android.widget.CalendarView', 'int',
                                       'int', 'int')

UTC_START = datetime(1970, 1, 1)


class AndroidCalendarView(AndroidFrameLayout, ProxyCalendarView):
    """ An Android implementation of an Enaml ProxyCalendarView.

    """
    #: A reference to the widget created by the proxy.
    widget = Typed(CalendarView)

    # -------------------------------------------------------------------------
    # Initialization API
    # -------------------------------------------------------------------------
    def create_widget(self):
        """ Create the underlying widget.

        """
        d = self.declaration
        self.widget = CalendarView(self.get_context(), None,
                                   d.style or "@attr/calendarViewStyle")

    def init_widget(self):
        """ Initialize the underlying widget.

        """
        super(AndroidCalendarView, self).init_widget()

        #: Setup listener
        w = self.widget
        w.setOnDateChangeListener(w.getId())
        w.onSelectedDayChange.connect(self.on_selected_day_change)

    # -------------------------------------------------------------------------
    # OnDateChangeListener API
    # -------------------------------------------------------------------------
    def on_selected_day_change(self, view, year, month, day):
        d = self.declaration
        with self.widget.setDate.suppressed():
            d.date = datetime(year, month+1, day)

    # -------------------------------------------------------------------------
    # ProxyCalendarView API
    # -------------------------------------------------------------------------
    def set_date(self, date):
        #: Convert date tuple to long
        s = long((date - UTC_START).total_seconds()*1000)
        self.widget.setDate(s)

    def set_min_date(self, date):
        #: Convert to long
        s = long((date - UTC_START).total_seconds()*1000)
        self.widget.setMinDate(s)

    def set_max_date(self, date):
        #: Convert to long
        s = long((date - UTC_START).total_seconds()*1000)
        self.widget.setMaxDate(s)

    def set_first_day_of_week(self, day):
        self.widget.setFirstDayOfWeek(day)
