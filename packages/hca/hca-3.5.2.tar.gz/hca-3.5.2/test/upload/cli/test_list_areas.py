import os
import sys
import unittest
import uuid
from argparse import Namespace

import tweak
import six

from ... import CapturingIO, reset_tweak_changes

pkg_root = os.path.abspath(os.path.join(os.path.dirname(__file__), '..', '..'))  # noqa
sys.path.insert(0, pkg_root)  # noqa

import hca
from hca.upload.cli.list_areas_command import ListAreasCommand


class TestUploadCliListAreasCommand(unittest.TestCase):

    def setUp(self):
        self.area_uuid = str(uuid.uuid4())
        creds = "foo"
        self.urn = "dcp:upl:aws:dev:{}:{}".format(self.area_uuid, creds)

    @reset_tweak_changes
    def test_it_lists_areas_when_there_are_some(self):
        a_uuid = 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa'
        b_uuid = 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb'
        config = hca.get_config()
        config.upload = {
            'areas': {
                a_uuid: "dcp:upl:aws:dev:%s" % (a_uuid,),
                b_uuid: "dcp:upl:aws:dev:%s" % (b_uuid,),
            },
            'current_area': a_uuid
        }
        config.save()

        with CapturingIO('stdout') as stdout:
            ListAreasCommand(Namespace())

        six.assertRegex(self, stdout.captured(), "%s <- selected" % a_uuid)
        six.assertRegex(self, stdout.captured(), b_uuid)
