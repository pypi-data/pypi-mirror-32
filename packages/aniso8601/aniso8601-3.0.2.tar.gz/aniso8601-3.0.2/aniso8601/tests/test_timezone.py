# -*- coding: utf-8 -*-

# Copyright (c) 2018, Brandon Nielsen
# All rights reserved.
#
# This software may be modified and distributed under the terms
# of the BSD license.  See the LICENSE file for details.

import unittest
import datetime
import pickle

from aniso8601.exceptions import ISOFormatError
from aniso8601.timezone import parse_timezone, UTCOffset

class TestTimezoneParserFunctions(unittest.TestCase):
    def test_parse_timezone(self):
        tzinfoobject = parse_timezone('Z')
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(hours=0))
        self.assertEqual(tzinfoobject.tzname(None), 'UTC')

        tzinfoobject = parse_timezone('+00:00')
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(hours=0))
        self.assertEqual(tzinfoobject.tzname(None), '+00:00')

        tzinfoobject = parse_timezone('+01:00')
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(hours=1))
        self.assertEqual(tzinfoobject.tzname(None), '+01:00')

        tzinfoobject = parse_timezone('-01:00')
        self.assertEqual(tzinfoobject.utcoffset(None), -datetime.timedelta(hours=1))
        self.assertEqual(tzinfoobject.tzname(None), '-01:00')

        tzinfoobject = parse_timezone('+00:12')
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(minutes=12))
        self.assertEqual(tzinfoobject.tzname(None), '+00:12')

        tzinfoobject = parse_timezone('+01:23')
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(hours=1, minutes=23))
        self.assertEqual(tzinfoobject.tzname(None), '+01:23')

        tzinfoobject = parse_timezone('-01:23')
        self.assertEqual(tzinfoobject.utcoffset(None), -datetime.timedelta(hours=1, minutes=23))
        self.assertEqual(tzinfoobject.tzname(None), '-01:23')

        tzinfoobject = parse_timezone('+0000')
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(hours=0))
        self.assertEqual(tzinfoobject.tzname(None), '+0000')

        tzinfoobject = parse_timezone('+0100')
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(hours=1))
        self.assertEqual(tzinfoobject.tzname(None), '+0100')

        tzinfoobject = parse_timezone('-0100')
        self.assertEqual(tzinfoobject.utcoffset(None), -datetime.timedelta(hours=1))
        self.assertEqual(tzinfoobject.tzname(None), '-0100')

        tzinfoobject = parse_timezone('+0012')
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(minutes=12))
        self.assertEqual(tzinfoobject.tzname(None), '+0012')

        tzinfoobject = parse_timezone('+0123')
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(hours=1, minutes=23))
        self.assertEqual(tzinfoobject.tzname(None), '+0123')

        tzinfoobject = parse_timezone('-0123')
        self.assertEqual(tzinfoobject.utcoffset(None), -datetime.timedelta(hours=1, minutes=23))
        self.assertEqual(tzinfoobject.tzname(None), '-0123')

        tzinfoobject = parse_timezone('+00')
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(hours=0))
        self.assertEqual(tzinfoobject.tzname(None), '+00')

        tzinfoobject = parse_timezone('+01')
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(hours=1))
        self.assertEqual(tzinfoobject.tzname(None), '+01')

        tzinfoobject = parse_timezone('-01')
        self.assertEqual(tzinfoobject.utcoffset(None), -datetime.timedelta(hours=1))
        self.assertEqual(tzinfoobject.tzname(None), '-01')

        tzinfoobject = parse_timezone('+12')
        self.assertEqual(tzinfoobject.utcoffset(None), datetime.timedelta(hours=12))
        self.assertEqual(tzinfoobject.tzname(None), '+12')

        tzinfoobject = parse_timezone('-12')
        self.assertEqual(tzinfoobject.utcoffset(None), -datetime.timedelta(hours=12))
        self.assertEqual(tzinfoobject.tzname(None), '-12')

    def test_parse_timezone_tzstr(self):
        with self.assertRaises(ISOFormatError):
            parse_timezone('Y')

        with self.assertRaises(ISOFormatError):
            parse_timezone(' Z')

        with self.assertRaises(ISOFormatError):
            parse_timezone('Z ')

        with self.assertRaises(ISOFormatError):
            parse_timezone(' Z ')

    def test_parse_timezone_negativezero(self):
        #A 0 offset cannot be negative
        with self.assertRaises(ISOFormatError):
            parse_timezone('-00:00')

        with self.assertRaises(ISOFormatError):
            parse_timezone('-0000')

        with self.assertRaises(ISOFormatError):
            parse_timezone('-00')

    def test_pickle(self):
        #Make sure timezone objects are pickleable
        testutcoffset = UTCOffset(name='UTC', minutes=0)

        utcoffsetpickle = pickle.dumps(testutcoffset)

        resultutcoffset = pickle.loads(utcoffsetpickle)

        self.assertEqual(resultutcoffset._name, testutcoffset._name)
        self.assertEqual(resultutcoffset._utcdelta, testutcoffset._utcdelta)

    def test_string_representation(self):
        #Make sure UTC offsets can be printed out prettily
        tzinfoobject = parse_timezone('+00:00')
        self.assertEqual(str(tzinfoobject), '+0:00:00 UTC')

        tzinfoobject = parse_timezone('+01:00')
        self.assertEqual(str(tzinfoobject), '+1:00:00 UTC')

        tzinfoobject = parse_timezone('-01:00')
        self.assertEqual(str(tzinfoobject), '-1:00:00 UTC')

        tzinfoobject = parse_timezone('+00:12')
        self.assertEqual(str(tzinfoobject), '+0:12:00 UTC')

        tzinfoobject = parse_timezone('-00:12')
        self.assertEqual(str(tzinfoobject), '-0:12:00 UTC')

        tzinfoobject = parse_timezone('+01:23')
        self.assertEqual(str(tzinfoobject), '+1:23:00 UTC')

        tzinfoobject = parse_timezone('-01:23')
        self.assertEqual(str(tzinfoobject), '-1:23:00 UTC')

        tzinfoobject = parse_timezone('+24:00')
        self.assertEqual(str(tzinfoobject), '+1 day, 0:00:00 UTC')

        tzinfoobject = parse_timezone('-24:00')
        self.assertEqual(str(tzinfoobject), '-1 day, 0:00:00 UTC')

        tzinfoobject = parse_timezone('+49:27')
        self.assertEqual(str(tzinfoobject), '+2 days, 1:27:00 UTC')

        tzinfoobject = parse_timezone('-49:27')
        self.assertEqual(str(tzinfoobject), '-2 days, 1:27:00 UTC')

    def test_datetime_tzinfo_dst(self):
        tzinfoobject = parse_timezone('+04:00')
        #This would raise ISOFormatError or a TypeError if dst info is invalid
        result = datetime.datetime.now(tzinfoobject)
        #Hacky way to make sure the tzinfo is what we'd expect
        self.assertEqual(result.tzinfo.utcoffset(None), datetime.timedelta(hours=4))
