import os
import sys
import unittest
import uuid
from argparse import Namespace

import six
import tweak

from ... import CapturingIO, reset_tweak_changes

pkg_root = os.path.abspath(os.path.join(os.path.dirname(__file__), '..', '..'))  # noqa
sys.path.insert(0, pkg_root)  # noqa

import hca
from hca.upload.cli.select_command import SelectCommand


class TestUploadCliSelectCommand(unittest.TestCase):

    def setUp(self):
        self.area_uuid = str(uuid.uuid4())
        creds = "foo"
        self.urn = "dcp:upl:aws:dev:{}:{}".format(self.area_uuid, creds)

    @reset_tweak_changes
    def test_when_given_an_unrecognized_urn_it_stores_it_in_upload_area_list_and_sets_it_as_current_area(self):
        with CapturingIO('stdout') as stdout:
            args = Namespace(urn_or_alias=self.urn)
            SelectCommand(args)

        config = hca.get_config()
        self.assertIn(self.area_uuid, config.upload.areas)
        self.assertEqual(self.urn, config.upload.areas[self.area_uuid])
        self.assertEqual(self.area_uuid, config.upload.current_area)

    @reset_tweak_changes
    def test_when_given_a_urn_it_prints_an_alias(self):
        config = hca.get_config()
        config.upload = {
            'areas': {
                'deadbeef-dead-dead-dead-beeeeeeeeeef': 'dcp:upl:aws:dev:deadbeef-dead-dead-dead-beeeeeeeeeef:creds',
            }
        }
        config.save()
        area_uuid = 'deafbeef-deaf-deaf-deaf-beeeeeeeeeef'
        urn = 'dcp:upl:aws:dev:{}:creds'.format(area_uuid)

        with CapturingIO('stdout') as stdout:
            args = Namespace(urn_or_alias=urn)
            SelectCommand(args)

        six.assertRegex(self, stdout.captured(), "alias \"{}\"".format('deaf'))

    @reset_tweak_changes
    def test_when_given_an_alias_that_matches_no_areas_it_prints_a_warning(self):

        config = hca.get_config()
        config.upload = {'areas': {}}
        config.save()

        with CapturingIO('stdout') as stdout:
            args = Namespace(urn_or_alias='aaa')
            SelectCommand(args)

        six.assertRegex(self, stdout.captured(), "don't recognize area")

    @reset_tweak_changes
    def test_when_given_an_alias_that_matches_more_than_one_area_it_prints_a_warning(self):
        config = hca.get_config()
        config.upload = {
            'areas': {
                'deadbeef-dead-dead-dead-beeeeeeeeeef': 'dcp:upl:aws:dev:deadbeef-dead-dead-dead-beeeeeeeeeef:creds',
                'deafbeef-deaf-deaf-deaf-beeeeeeeeeef': 'dcp:upl:aws:dev:deafbeef-deaf-deaf-deaf-beeeeeeeeeef:creds',
            }
        }
        config.save()

        with CapturingIO('stdout') as stdout:
            args = Namespace(urn_or_alias='dea')
            SelectCommand(args)

        six.assertRegex(self, stdout.captured(), "matches more than one")

    @reset_tweak_changes
    def test_when_given_an_alias_that_matches_one_area_it_selects_it(self):
        a_uuid = 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa'
        b_uuid = 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb'
        config = hca.get_config()
        config.upload = {
            'areas': {
                a_uuid: "dcp:upl:aws:dev:%s" % (a_uuid,),
                b_uuid: "dcp:upl:aws:dev:%s" % (b_uuid,),
            }
        }
        config.save()

        with CapturingIO('stdout') as stdout:
            args = Namespace(urn_or_alias='bbb')
            SelectCommand(args)

        config = hca.get_config()
        self.assertEqual(b_uuid, config.upload.current_area)
