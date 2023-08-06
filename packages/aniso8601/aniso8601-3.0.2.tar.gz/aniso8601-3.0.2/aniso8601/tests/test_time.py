# -*- coding: utf-8 -*-

# Copyright (c) 2018, Brandon Nielsen
# All rights reserved.
#
# This software may be modified and distributed under the terms
# of the BSD license.  See the LICENSE file for details.

import datetime
import unittest

from aniso8601.exceptions import HoursOutOfBoundsError, LeapSecondError, \
        MidnightBoundsError, MinutesOutOfBoundsError, SecondsOutOfBoundsError
from aniso8601.resolution import TimeResolution
from aniso8601.time import get_time_resolution, parse_datetime, parse_time, \
    _parse_time_naive, _parse_hour, _parse_minute_time, _parse_second_time, \
    _build_time, _split_tz

class TestTimeParserFunctions(unittest.TestCase):
    def test_get_time_resolution(self):
        self.assertEqual(get_time_resolution('01:23:45'), TimeResolution.Seconds)
        self.assertEqual(get_time_resolution('24:00:00'), TimeResolution.Seconds)
        self.assertEqual(get_time_resolution('23:21:28.512400'),
                         TimeResolution.Seconds)
        self.assertEqual(get_time_resolution('01:23'), TimeResolution.Minutes)
        self.assertEqual(get_time_resolution('24:00'), TimeResolution.Minutes)
        self.assertEqual(get_time_resolution('01:23.4567'), TimeResolution.Minutes)
        self.assertEqual(get_time_resolution('012345'), TimeResolution.Seconds)
        self.assertEqual(get_time_resolution('240000'), TimeResolution.Seconds)
        self.assertEqual(get_time_resolution('0123'), TimeResolution.Minutes)
        self.assertEqual(get_time_resolution('2400'), TimeResolution.Minutes)
        self.assertEqual(get_time_resolution('01'), TimeResolution.Hours)
        self.assertEqual(get_time_resolution('24'), TimeResolution.Hours)
        self.assertEqual(get_time_resolution('12.5'), TimeResolution.Hours)
        self.assertEqual(get_time_resolution('232128.512400+00:00'),
                         TimeResolution.Seconds)
        self.assertEqual(get_time_resolution('0123.4567+00:00'),
                         TimeResolution.Minutes)
        self.assertEqual(get_time_resolution('01.4567+00:00'), TimeResolution.Hours)
        self.assertEqual(get_time_resolution('01:23:45+00:00'),
                         TimeResolution.Seconds)
        self.assertEqual(get_time_resolution('24:00:00+00:00'),
                         TimeResolution.Seconds)
        self.assertEqual(get_time_resolution('23:21:28.512400+00:00'),
                         TimeResolution.Seconds)
        self.assertEqual(get_time_resolution('01:23+00:00'), TimeResolution.Minutes)
        self.assertEqual(get_time_resolution('24:00+00:00'), TimeResolution.Minutes)
        self.assertEqual(get_time_resolution('01:23.4567+00:00'),
                         TimeResolution.Minutes)
        self.assertEqual(get_time_resolution('23:21:28.512400+11:15'),
                         TimeResolution.Seconds)
        self.assertEqual(get_time_resolution('23:21:28.512400-12:34'),
                         TimeResolution.Seconds)
        self.assertEqual(get_time_resolution('23:21:28.512400Z'),
                         TimeResolution.Seconds)
        self.assertEqual(get_time_resolution('06:14:00.000123Z'),
                         TimeResolution.Seconds)

    def test_parse_time(self):
        time = parse_time('01:23:45')
        self.assertEqual(time, datetime.time(hour=1, minute=23, second=45))

        time = parse_time('24:00:00')
        self.assertEqual(time, datetime.time(hour=0))

        time = parse_time('23:21:28.512400')
        self.assertEqual(time, datetime.time(hour=23, minute=21, second=28, microsecond=512400))

        time = parse_time('14:43:59.9999997')
        self.assertEqual(time, datetime.time(hour=14, minute=43, second=59, microsecond=999999))

        time = parse_time('01:23')
        self.assertEqual(time, datetime.time(hour=1, minute=23))

        time = parse_time('24:00')
        self.assertEqual(time, datetime.time(hour=0))

        time = parse_time('01:23.4567')
        self.assertEqual(time, datetime.time(hour=1, minute=23, second=27, microsecond=402000))

        time = parse_time('012345')
        self.assertEqual(time, datetime.time(hour=1, minute=23, second=45))

        time = parse_time('240000')
        self.assertEqual(time, datetime.time(hour=0))

        time = parse_time('144359.9999997')
        self.assertEqual(time, datetime.time(hour=14, minute=43, second=59, microsecond=999999))

        time = parse_time('0123')
        self.assertEqual(time, datetime.time(hour=1, minute=23))

        time = parse_time('2400')
        self.assertEqual(time, datetime.time(hour=0))

        time = parse_time('01')
        self.assertEqual(time, datetime.time(hour=1))

        time = parse_time('24')
        self.assertEqual(time, datetime.time(hour=0))

        time = parse_time('12.5')
        self.assertEqual(time, datetime.time(hour=12, minute=30))

        time = parse_time('232128.512400+00:00')
        self.assertEqual(time.replace(tzinfo=None), datetime.time(hour=23, minute=21, second=28, microsecond=512400))

        tzinfoobject = time.tzinfo
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(hours=0))
        self.assertEqual(tzinfoobject.tzname(None), '+00:00')

        time = parse_time('0123.4567+00:00')
        self.assertEqual(time.replace(tzinfo=None), datetime.time(hour=1, minute=23, second=27, microsecond=402000))

        tzinfoobject = time.tzinfo
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(hours=0))
        self.assertEqual(tzinfoobject.tzname(None), '+00:00')

        time = parse_time('01.4567+00:00')
        self.assertEqual(time.replace(tzinfo=None), datetime.time(hour=1, minute=27, second=24, microsecond=120000))

        tzinfoobject = time.tzinfo
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(hours=0))
        self.assertEqual(tzinfoobject.tzname(None), '+00:00')

        time = parse_time('01:23:45+00:00')
        self.assertEqual(time.replace(tzinfo=None), datetime.time(hour=1, minute=23, second=45))

        tzinfoobject = time.tzinfo
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(hours=0))
        self.assertEqual(tzinfoobject.tzname(None), '+00:00')

        time = parse_time('24:00:00+00:00')
        self.assertEqual(time.replace(tzinfo=None), datetime.time(hour=0))

        tzinfoobject = time.tzinfo
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(hours=0))
        self.assertEqual(tzinfoobject.tzname(None), '+00:00')

        time = parse_time('23:21:28.512400+00:00')
        self.assertEqual(time.replace(tzinfo=None), datetime.time(hour=23, minute=21, second=28, microsecond=512400))

        tzinfoobject = time.tzinfo
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(hours=0))
        self.assertEqual(tzinfoobject.tzname(None), '+00:00')

        time = parse_time('01:23+00:00')
        self.assertEqual(time.replace(tzinfo=None), datetime.time(hour=1, minute=23))

        tzinfoobject = time.tzinfo
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(hours=0))
        self.assertEqual(tzinfoobject.tzname(None), '+00:00')

        time = parse_time('24:00+00:00')
        self.assertEqual(time.replace(tzinfo=None), datetime.time(hour=0))

        tzinfoobject = time.tzinfo
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(hours=0))
        self.assertEqual(tzinfoobject.tzname(None), '+00:00')

        time = parse_time('01:23.4567+00:00')
        self.assertEqual(time.replace(tzinfo=None), datetime.time(hour=1, minute=23, second=27, microsecond=402000))

        tzinfoobject = time.tzinfo
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(hours=0))
        self.assertEqual(tzinfoobject.tzname(None), '+00:00')

        time = parse_time('23:21:28.512400+11:15')
        self.assertEqual(time.replace(tzinfo=None), datetime.time(hour=23, minute=21, second=28, microsecond=512400))

        tzinfoobject = time.tzinfo
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(hours=11, minutes=15))
        self.assertEqual(tzinfoobject.tzname(None), '+11:15')

        time = parse_time('23:21:28.512400-12:34')
        self.assertEqual(time.replace(tzinfo=None), datetime.time(hour=23, minute=21, second=28, microsecond=512400))

        tzinfoobject = time.tzinfo
        self.assertEqual(tzinfoobject.utcoffset(None), -datetime.timedelta(hours=12, minutes=34))
        self.assertEqual(tzinfoobject.tzname(None), '-12:34')

        time = parse_time('23:21:28.512400Z')
        self.assertEqual(time.replace(tzinfo=None), datetime.time(hour=23, minute=21, second=28, microsecond=512400))

        tzinfoobject = time.tzinfo
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(hours=0))
        self.assertEqual(tzinfoobject.tzname(None), 'UTC')

        time = parse_time('06:14:00.000123Z')
        self.assertEqual(time.replace(tzinfo=None), datetime.time(hour=6, minute=14, microsecond=123))

        tzinfoobject = time.tzinfo
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(hours=0))
        self.assertEqual(tzinfoobject.tzname(None), 'UTC')

    def test_parse_time_bounds(self):
        #Leap seconds not supported
        #https://bitbucket.org/nielsenb/aniso8601/issues/10/sub-microsecond-precision-in-durations-is
        #https://bitbucket.org/nielsenb/aniso8601/issues/13/parsing-of-leap-second-gives-wildly
        with self.assertRaises(LeapSecondError):
            parse_time('23:59:60')

        with self.assertRaises(LeapSecondError):
            parse_time('23:59:60Z')

        with self.assertRaises(LeapSecondError):
            parse_time('23:59:60+00:00')

        #Leap seconds not supported, but exception type should be changed
        with self.assertRaises(SecondsOutOfBoundsError):
            parse_time('00:00:60')

        with self.assertRaises(SecondsOutOfBoundsError):
            parse_time('00:00:60Z')

        with self.assertRaises(SecondsOutOfBoundsError):
            parse_time('00:00:60+00:00')

        #Seconds must not be greater than or equal to 60
        with self.assertRaises(SecondsOutOfBoundsError):
            parse_time('00:00:61')

        with self.assertRaises(SecondsOutOfBoundsError):
            parse_time('00:00:61Z')

        with self.assertRaises(SecondsOutOfBoundsError):
            parse_time('00:00:61+00:00')

        #Minutes must not be greater than 60
        with self.assertRaises(MinutesOutOfBoundsError):
            parse_time('00:61')

        with self.assertRaises(MinutesOutOfBoundsError):
            parse_time('00:61Z')

        with self.assertRaises(MinutesOutOfBoundsError):
            parse_time('00:61+00:00')

    def test_parse_datetime(self):
        resultdatetime = parse_datetime('1981-04-05T23:21:28.512400Z')
        self.assertEqual(resultdatetime.replace(tzinfo=None), datetime.datetime(1981, 4, 5, hour=23, minute=21, second=28, microsecond=512400))

        tzinfoobject = resultdatetime.tzinfo
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(hours=0))
        self.assertEqual(tzinfoobject.tzname(None), 'UTC')

        resultdatetime = parse_datetime('1981095T23:21:28.512400-12:34')
        self.assertEqual(resultdatetime.replace(tzinfo=None), datetime.datetime(1981, 4, 5, hour=23, minute=21, second=28, microsecond=512400))

        tzinfoobject = resultdatetime.tzinfo
        self.assertEqual(tzinfoobject.utcoffset(None), -datetime.timedelta(hours=12, minutes=34))
        self.assertEqual(tzinfoobject.tzname(None), '-12:34')

        resultdatetime = parse_datetime('19810405T23:21:28+00')
        self.assertEqual(resultdatetime.replace(tzinfo=None), datetime.datetime(1981, 4, 5, hour=23, minute=21, second=28))

        tzinfoobject = resultdatetime.tzinfo
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(hours=0))
        self.assertEqual(tzinfoobject.tzname(None), '+00')

        resultdatetime = parse_datetime('19810405T23:21:28+00:00')
        self.assertEqual(resultdatetime.replace(tzinfo=None), datetime.datetime(1981, 4, 5, hour=23, minute=21, second=28))

        tzinfoobject = resultdatetime.tzinfo
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(hours=0))
        self.assertEqual(tzinfoobject.tzname(None), '+00:00')

    def test_parse_datetime_bounds(self):
        #Leap seconds not supported
        #https://bitbucket.org/nielsenb/aniso8601/issues/10/sub-microsecond-precision-in-durations-is
        #https://bitbucket.org/nielsenb/aniso8601/issues/13/parsing-of-leap-second-gives-wildly
        with self.assertRaises(LeapSecondError):
            parse_datetime('2016-12-31T23:59:60+00:00')

        with self.assertRaises(LeapSecondError):
            parse_datetime('2016-12-31T23:59:60')

        with self.assertRaises(LeapSecondError):
            parse_datetime('2016-12-31T23:59:60Z')

        #Seconds can't be greater than 60
        with self.assertRaises(SecondsOutOfBoundsError):
            parse_datetime('1981-04-05T00:00:60')

        with self.assertRaises(SecondsOutOfBoundsError):
            parse_datetime('1981-04-05T00:00:60Z')

        with self.assertRaises(SecondsOutOfBoundsError):
            parse_datetime('1981-04-05T00:00:60+00:00')

        with self.assertRaises(SecondsOutOfBoundsError):
            parse_datetime('2016-12-31T23:59:61+00:00')

        with self.assertRaises(SecondsOutOfBoundsError):
            parse_datetime('2016-12-31T23:59:61')

        with self.assertRaises(SecondsOutOfBoundsError):
            parse_datetime('1981-04-05T00:00:61')

        with self.assertRaises(SecondsOutOfBoundsError):
            parse_datetime('1981-04-05T00:00:61Z')

        with self.assertRaises(SecondsOutOfBoundsError):
            parse_datetime('1981-04-05T00:00:61+00:00')

        #Minutes can't be greater than 60
        with self.assertRaises(MinutesOutOfBoundsError):
            parse_datetime('1981-04-05T00:61')

        with self.assertRaises(MinutesOutOfBoundsError):
            parse_datetime('1981-04-05T00:61Z')

        with self.assertRaises(MinutesOutOfBoundsError):
            parse_datetime('1981-04-05T00:61+00:00')

    def test_parse_datetime_spaceseperated(self):
        resultdatetime = parse_datetime('2004-W53-6 23:21:28.512400-12:34', ' ')
        self.assertEqual(resultdatetime.replace(tzinfo=None), datetime.datetime(2005, 1, 1, hour=23, minute=21, second=28, microsecond=512400))

        tzinfoobject = resultdatetime.tzinfo
        self.assertEqual(tzinfoobject.utcoffset(None), -datetime.timedelta(hours=12, minutes=34))
        self.assertEqual(tzinfoobject.tzname(None), '-12:34')

    def test_parse_time_naive(self):
        time = _parse_time_naive('01:23:45')
        self.assertEqual(time, datetime.time(hour=1, minute=23, second=45))

        time = _parse_time_naive('24:00:00')
        self.assertEqual(time, datetime.time(hour=0))

        time = _parse_time_naive('23:21:28.512400')
        self.assertEqual(time, datetime.time(hour=23, minute=21, second=28, microsecond=512400))

        time = _parse_time_naive('01:23')
        self.assertEqual(time, datetime.time(hour=1, minute=23))

        time = _parse_time_naive('24:00')
        self.assertEqual(time, datetime.time(hour=0))

        time = _parse_time_naive('01:23.4567')
        self.assertEqual(time, datetime.time(hour=1, minute=23, second=27, microsecond=402000))

        time = _parse_time_naive('012345')
        self.assertEqual(time, datetime.time(hour=1, minute=23, second=45))

        time = _parse_time_naive('240000')
        self.assertEqual(time, datetime.time(hour=0))

        time = _parse_time_naive('0123')
        self.assertEqual(time, datetime.time(hour=1, minute=23))

        time = _parse_time_naive('2400')
        self.assertEqual(time, datetime.time(hour=0))

        time = _parse_time_naive('01')
        self.assertEqual(time, datetime.time(hour=1))

        time = _parse_time_naive('24')
        self.assertEqual(time, datetime.time(hour=0))

        time = _parse_time_naive('232128.512400')
        self.assertEqual(time, datetime.time(hour=23, minute=21, second=28, microsecond=512400))

        time = _parse_time_naive('0123.4567')
        self.assertEqual(time, datetime.time(hour=1, minute=23, second=27, microsecond=402000))

        time = _parse_time_naive('01.4567')
        self.assertEqual(time, datetime.time(hour=1, minute=27, second=24, microsecond=120000))

        with self.assertRaises(MinutesOutOfBoundsError):
            _parse_time_naive('00:61')

    def test_parse_time_naive_bounds(self):
        #Leap seconds not supporter
        with self.assertRaises(LeapSecondError):
            _parse_time_naive('23:59:60')

        #Seconds must not be greater than or equal to 60
        with self.assertRaises(SecondsOutOfBoundsError):
            _parse_time_naive('00:00:60')

        with self.assertRaises(SecondsOutOfBoundsError):
            _parse_time_naive('00:00:61')

        #Minutes must not be greater than 60
        with self.assertRaises(MinutesOutOfBoundsError):
            _parse_time_naive('00:61')

    def test_parse_time_naive_tz(self):
        #These fail because of the timezone being parsed by the naive parser
        with self.assertRaises(ValueError):
            _parse_time_naive('00:00:60+00:00')

        with self.assertRaises(ValueError):
            _parse_time_naive('00:61+00:00')

    def test_parse_hour(self):
        time = _parse_hour('01')
        self.assertEqual(time, datetime.time(hour=1))

        time = _parse_hour('24')
        self.assertEqual(time, datetime.time(hour=0))

        time = _parse_hour('01.4567')
        self.assertEqual(time, datetime.time(hour=1, minute=27, second=24, microsecond=120000))

        time = _parse_hour('12.5')
        self.assertEqual(time, datetime.time(hour=12, minute=30))

    def test_parse_hour_bounds(self):
        #Hour cannot be larger than 24
        with self.assertRaises(HoursOutOfBoundsError):
            _parse_hour('24.1')

    def test_parse_minute_time(self):
        time = _parse_minute_time('01:23')
        self.assertEqual(time, datetime.time(hour=1, minute=23))

        time = _parse_minute_time('24:00')
        self.assertEqual(time, datetime.time(hour=0))

        time = _parse_minute_time('01:23.4567')
        self.assertEqual(time, datetime.time(hour=1, minute=23, second=27, microsecond=402000))

        time = _parse_minute_time('0123')
        self.assertEqual(time, datetime.time(hour=1, minute=23))

        time = _parse_minute_time('2400')
        self.assertEqual(time, datetime.time(hour=0))

        time = _parse_minute_time('0123.4567')
        self.assertEqual(time, datetime.time(hour=1, minute=23, second=27, microsecond=402000))

    def test_parse_minute_time_bounds(self):
        #Minutes cannot be greater than 60
        with self.assertRaises(MinutesOutOfBoundsError):
            _parse_minute_time('0060')

        with self.assertRaises(MinutesOutOfBoundsError):
            _parse_minute_time('0060.1')

        with self.assertRaises(MinutesOutOfBoundsError):
            _parse_minute_time('00:60')

        with self.assertRaises(MinutesOutOfBoundsError):
            _parse_minute_time('00:60.1')

        #Hour 24 can only represent midnight
        with self.assertRaises(MidnightBoundsError):
            _parse_minute_time('2401')

        with self.assertRaises(MidnightBoundsError):
            _parse_minute_time('2400.1')

        with self.assertRaises(MidnightBoundsError):
            _parse_minute_time('24:01')

        with self.assertRaises(MidnightBoundsError):
            _parse_minute_time('24:00.1')

    def test_parse_second_time(self):
        time = _parse_second_time('01:23:45')
        self.assertEqual(time, datetime.time(hour=1, minute=23, second=45))

        time = _parse_second_time('24:00:00')
        self.assertEqual(time, datetime.time(hour=0))

        time = _parse_second_time('23:21:28.512400')
        self.assertEqual(time, datetime.time(hour=23, minute=21, second=28, microsecond=512400))

        time = _parse_second_time('14:43:59.9999997')
        self.assertEqual(time, datetime.time(hour=14, minute=43, second=59, microsecond=999999))

        time = _parse_second_time('012345')
        self.assertEqual(time, datetime.time(hour=1, minute=23, second=45))

        time = _parse_second_time('240000')
        self.assertEqual(time, datetime.time(hour=0))

        time = _parse_second_time('232128.512400')
        self.assertEqual(time, datetime.time(hour=23, minute=21, second=28, microsecond=512400))

        time = _parse_second_time('144359.9999997')
        self.assertEqual(time, datetime.time(hour=14, minute=43, second=59, microsecond=999999))

    def test_parse_second_time_bounds(self):
        #Leap seconds not supported
        with self.assertRaises(LeapSecondError):
            _parse_second_time('235960')

        with self.assertRaises(LeapSecondError):
            _parse_second_time('23:59:60')

        #Seconds must be less than 60
        #Leap seconds not supported
        with self.assertRaises(SecondsOutOfBoundsError):
            _parse_second_time('000060')

        with self.assertRaises(SecondsOutOfBoundsError):
            _parse_second_time('00:00:60')

        with self.assertRaises(SecondsOutOfBoundsError):
            _parse_second_time('000061')

        with self.assertRaises(SecondsOutOfBoundsError):
            _parse_second_time('00:00:61')

        #Minutes must be less than 60
        with self.assertRaises(MinutesOutOfBoundsError):
            _parse_second_time('006000')

        with self.assertRaises(MinutesOutOfBoundsError):
            _parse_second_time('00:60:00')

        #Hour 24 can only represent midnight
        with self.assertRaises(MidnightBoundsError):
            _parse_second_time('240001')

        with self.assertRaises(MidnightBoundsError):
            _parse_second_time('24:00:01')

        with self.assertRaises(MidnightBoundsError):
            _parse_second_time('240100')

        with self.assertRaises(MidnightBoundsError):
            _parse_second_time('24:01:00')

    def test_build_time(self):
        self.assertEqual(_build_time(datetime.time(hour=1), datetime.timedelta(hours=1.1, minutes=2.2, seconds=3.3)), datetime.time(hour=2, minute=8, second=15, microsecond=300000))

        #Make sure it overflows correctly
        self.assertEqual(_build_time(datetime.time.max, datetime.timedelta(microseconds=1)), datetime.time.min)

    def test_split_tz(self):
        self.assertEqual(_split_tz('01:23:45'), ('01:23:45', None))

        self.assertEqual(_split_tz('24:00:00'), ('24:00:00', None))

        self.assertEqual(_split_tz('23:21:28.512400'), ('23:21:28.512400', None))

        self.assertEqual(_split_tz('01:23'), ('01:23', None))

        self.assertEqual(_split_tz('24:00'), ('24:00', None))

        self.assertEqual(_split_tz('01:23.4567'), ('01:23.4567', None))

        self.assertEqual(_split_tz('012345'), ('012345', None))

        self.assertEqual(_split_tz('240000'), ('240000', None))

        self.assertEqual(_split_tz('0123'), ('0123', None))

        self.assertEqual(_split_tz('2400'), ('2400', None))

        self.assertEqual(_split_tz('01'), ('01', None))

        self.assertEqual(_split_tz('24'), ('24', None))

        self.assertEqual(_split_tz('12.5'), ('12.5', None))

        self.assertEqual(_split_tz('232128.512400+00:00'), ('232128.512400', '+00:00'))

        self.assertEqual(_split_tz('0123.4567+00:00'), ('0123.4567', '+00:00'))

        self.assertEqual(_split_tz('01.4567+00:00'), ('01.4567', '+00:00'))

        self.assertEqual(_split_tz('01:23:45+00:00'), ('01:23:45', '+00:00'))

        self.assertEqual(_split_tz('24:00:00+00:00'), ('24:00:00', '+00:00'))

        self.assertEqual(_split_tz('23:21:28.512400+00:00'), ('23:21:28.512400', '+00:00'))

        self.assertEqual(_split_tz('01:23+00:00'), ('01:23', '+00:00'))

        self.assertEqual(_split_tz('24:00+00:00'), ('24:00', '+00:00'))

        self.assertEqual(_split_tz('01:23.4567+00:00'), ('01:23.4567', '+00:00'))

        self.assertEqual(_split_tz('23:21:28.512400+11:15'), ('23:21:28.512400', '+11:15'))

        self.assertEqual(_split_tz('23:21:28.512400-12:34'), ('23:21:28.512400', '-12:34'))

        self.assertEqual(_split_tz('23:21:28.512400Z'), ('23:21:28.512400', 'Z'))

        self.assertEqual(_split_tz('06:14:00.000123Z'), ('06:14:00.000123', 'Z'))
