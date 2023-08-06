# -*- coding: utf-8 -*-

# Copyright (c) 2018, Brandon Nielsen
# All rights reserved.
#
# This software may be modified and distributed under the terms
# of the BSD license.  See the LICENSE file for details.

import datetime
import unittest

from aniso8601.exceptions import DayOutOfBoundsError, ISOFormatError, \
        WeekOutOfBoundsError, YearOutOfBoundsError
from aniso8601.date import parse_date, _parse_year, _parse_calendar_day, \
        _parse_calendar_month, _parse_week_day, _parse_week, \
        _parse_ordinal_date, get_date_resolution
from aniso8601.resolution import DateResolution

class TestDateResolutionFunctions(unittest.TestCase):
    def test_get_date_resolution_year(self):
        self.assertEqual(get_date_resolution('2013'), DateResolution.Year)
        self.assertEqual(get_date_resolution('0001'), DateResolution.Year)
        self.assertEqual(get_date_resolution('19'), DateResolution.Year)

    def test_get_date_resolution_month(self):
        self.assertEqual(get_date_resolution('1981-04'), DateResolution.Month)

    def test_get_date_resolution_week(self):
        self.assertEqual(get_date_resolution('2004-W53'), DateResolution.Week)
        self.assertEqual(get_date_resolution('2009-W01'), DateResolution.Week)
        self.assertEqual(get_date_resolution('2004W53'), DateResolution.Week)

    def test_get_date_resolution_year_weekday(self):
        self.assertEqual(get_date_resolution('2004-W53-6'), DateResolution.Weekday)
        self.assertEqual(get_date_resolution('2004W536'), DateResolution.Weekday)

    def test_get_date_resolution_year_ordinal(self):
        self.assertEqual(get_date_resolution('1981-095'), DateResolution.Ordinal)
        self.assertEqual(get_date_resolution('1981095'), DateResolution.Ordinal)

class TestDateParserFunctions(unittest.TestCase):
    def test_parse_date(self):
        date = parse_date('2013')
        self.assertEqual(date, datetime.date(2013, 1, 1))

        date = parse_date('0001')
        self.assertEqual(date, datetime.date(1, 1, 1))

        date = parse_date('19')
        self.assertEqual(date, datetime.date(1900, 1, 1))

        date = parse_date('1981-04-05')
        self.assertEqual(date, datetime.date(1981, 4, 5))

        date = parse_date('19810405')
        self.assertEqual(date, datetime.date(1981, 4, 5))

        date = parse_date('1981-04')
        self.assertEqual(date, datetime.date(1981, 4, 1))

        date = parse_date('2004-W53')
        self.assertEqual(date, datetime.date(2004, 12, 27))
        self.assertEqual(date.weekday(), 0)

        date = parse_date('2009-W01')
        self.assertEqual(date, datetime.date(2008, 12, 29))
        self.assertEqual(date.weekday(), 0)

        date = parse_date('2004-W53-6')
        self.assertEqual(date, datetime.date(2005, 1, 1))
        self.assertEqual(date.weekday(), 5)

        date = parse_date('2004W53')
        self.assertEqual(date, datetime.date(2004, 12, 27))
        self.assertEqual(date.weekday(), 0)

        date = parse_date('2004W536')
        self.assertEqual(date, datetime.date(2005, 1, 1))
        self.assertEqual(date.weekday(), 5)

        date = parse_date('1981-095')
        self.assertEqual(date, datetime.date(1981, 4, 5))

        date = parse_date('1981095')
        self.assertEqual(date, datetime.date(1981, 4, 5))

    def test_parse_date_bounds(self):
        #0 isn't a valid week number
        with self.assertRaises(WeekOutOfBoundsError):
            parse_date('2003-W00')

        with self.assertRaises(WeekOutOfBoundsError):
            parse_date('2003W00')

        #Week must not be larger than 53
        with self.assertRaises(WeekOutOfBoundsError):
            parse_date('2004-W54')

        with self.assertRaises(WeekOutOfBoundsError):
            parse_date('2004W54')

        #0 isn't a valid day number
        with self.assertRaises(DayOutOfBoundsError):
            parse_date('2001-W02-0')

        with self.assertRaises(DayOutOfBoundsError):
            parse_date('2001W020')

        #Day must not be larger than 7
        with self.assertRaises(DayOutOfBoundsError):
            parse_date('2001-W02-8')

        with self.assertRaises(DayOutOfBoundsError):
            parse_date('2001W028')

    def test_parse_year(self):
        date = _parse_year('2013')
        self.assertEqual(date, datetime.date(2013, 1, 1))

        date = _parse_year('0001')
        self.assertEqual(date, datetime.date(1, 1, 1))

        date = _parse_year('19')
        self.assertEqual(date, datetime.date(1900, 1, 1))

    def test_parse_year_zero(self):
        #0 isn't a valid year
        with self.assertRaises(YearOutOfBoundsError):
            _parse_year('0')

    def test_parse_calendar_day(self):
        date = _parse_calendar_day('1981-04-05')
        self.assertEqual(date, datetime.date(1981, 4, 5))

        date = _parse_calendar_day('19810405')
        self.assertEqual(date, datetime.date(1981, 4, 5))

    def test_parse_calendar_month(self):
        date = _parse_calendar_month('1981-04')
        self.assertEqual(date, datetime.date(1981, 4, 1))

    def test_parse_calendar_month_nohyphen(self):
        #Hyphen is required
        with self.assertRaises(ISOFormatError):
            _parse_calendar_month('198104')

    def test_parse_week_day(self):
        date = _parse_week_day('2004-W53-6')
        self.assertEqual(date, datetime.date(2005, 1, 1))
        self.assertEqual(date.weekday(), 5)

        date = _parse_week_day('2009-W01-1')
        self.assertEqual(date, datetime.date(2008, 12, 29))
        self.assertEqual(date.weekday(), 0)

        date = _parse_week_day('2009-W53-7')
        self.assertEqual(date, datetime.date(2010, 1, 3))
        self.assertEqual(date.weekday(), 6)

        date = _parse_week_day('2010-W01-1')
        self.assertEqual(date, datetime.date(2010, 1, 4))
        self.assertEqual(date.weekday(), 0)

        date = _parse_week_day('2004W536')
        self.assertEqual(date, datetime.date(2005, 1, 1))
        self.assertEqual(date.weekday(), 5)

        date = _parse_week_day('2009W011')
        self.assertEqual(date, datetime.date(2008, 12, 29))
        self.assertEqual(date.weekday(), 0)

        date = _parse_week_day('2009W537')
        self.assertEqual(date, datetime.date(2010, 1, 3))
        self.assertEqual(date.weekday(), 6)

        date = _parse_week_day('2010W011')
        self.assertEqual(date, datetime.date(2010, 1, 4))
        self.assertEqual(date.weekday(), 0)

    def test_parse_week(self):
        date = _parse_week('2004-W53')
        self.assertEqual(date, datetime.date(2004, 12, 27))
        self.assertEqual(date.weekday(), 0)

        date = _parse_week('2009-W01')
        self.assertEqual(date, datetime.date(2008, 12, 29))
        self.assertEqual(date.weekday(), 0)

        date = _parse_week('2009-W53')
        self.assertEqual(date, datetime.date(2009, 12, 28))
        self.assertEqual(date.weekday(), 0)

        date = _parse_week('2010-W01')
        self.assertEqual(date, datetime.date(2010, 1, 4))
        self.assertEqual(date.weekday(), 0)

        date = _parse_week('2004W53')
        self.assertEqual(date, datetime.date(2004, 12, 27))
        self.assertEqual(date.weekday(), 0)

        date = _parse_week('2009W01')
        self.assertEqual(date, datetime.date(2008, 12, 29))
        self.assertEqual(date.weekday(), 0)

        date = _parse_week('2009W53')
        self.assertEqual(date, datetime.date(2009, 12, 28))
        self.assertEqual(date.weekday(), 0)

        date = _parse_week('2010W01')
        self.assertEqual(date, datetime.date(2010, 1, 4))
        self.assertEqual(date.weekday(), 0)

    def test_parse_ordinal_date(self):
        date = _parse_ordinal_date('1981-095')
        self.assertEqual(date, datetime.date(1981, 4, 5))

        date = _parse_ordinal_date('1981095')
        self.assertEqual(date, datetime.date(1981, 4, 5))

        date = _parse_ordinal_date('1981365')
        self.assertEqual(date, datetime.date(1981, 12, 31))

        date = _parse_ordinal_date('1980366')
        self.assertEqual(date, datetime.date(1980, 12, 31))

    def test_parse_ordinal_date_zero(self):
        #0 isn't a valid day of year
        with self.assertRaises(DayOutOfBoundsError):
            _parse_ordinal_date('1981000')

    def test_parse_ordinal_date_nonleap(self):
        #Day 366 is only valid on a leap year
        with self.assertRaises(DayOutOfBoundsError):
            _parse_ordinal_date('1981366')

    def test_parse_ordinal_date_overflow(self):
        #Day must me 365, or 366, not larger
        with self.assertRaises(DayOutOfBoundsError):
            _parse_ordinal_date('1981367')
