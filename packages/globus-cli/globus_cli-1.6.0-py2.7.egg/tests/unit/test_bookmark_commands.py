import uuid
import json
import logging

from globus_sdk import GlobusAPIError, NetworkError

from tests.framework.constants import GO_EP1_ID
from tests.framework.cli_testcase import CliTestCase

log = logging.getLogger(__name__)


class BookmarkTests(CliTestCase):
    """
    Tests bookmark commands
    """
    def _clean(self):
        """
        delete all bookmarks on the class, catching 404s (and a host of other
        failures)
        """
        def safe_del(bmid):
            try:
                self.tc.delete_bookmark(bmid)
            except GlobusAPIError:
                log.exception('API error on bookmark tests cleanup')
            except NetworkError:
                log.exception('Network error on bookmark tests cleanup')

        # walk the bookmark list, get those that were defined here, and delete
        # them
        for bm in self.tc.bookmark_list():
            if bm['name'] in self.created_bookmark_names:
                safe_del(bm['id'])

    def gen_bookmark_name(self, name=""):
        if name:
            name = name + "-"
        bmname = "{}{}".format(name, str(uuid.uuid1()))
        self.created_bookmark_names.add(bmname)
        return bmname

    def setUp(self):
        super(BookmarkTests, self).setUp()

        self.created_bookmark_names = set()

        # use try-catch to ensure that cleanup this runs even if setUp crashes
        try:
            self.bm1name = self.gen_bookmark_name(name='bm1')
            res = self.tc.create_bookmark(
                {'endpoint_id': GO_EP1_ID,
                 'path': '/home/',
                 'name': self.bm1name})
            self.bm1id = res['id']
        except:
            self._clean()
            raise

    def tearDown(self):
        self._clean()
        super(BookmarkTests, self).tearDown()

    def test_bookmark_create(self):
        """
        Runs bookmark create, confirms simple things about text and json output
        """
        output = self.run_line(
            ("globus bookmark create "
             "{}:{} {}").format(GO_EP1_ID, '/share/',
                                self.gen_bookmark_name(name='sharebm')))
        self.assertIn('Bookmark ID: ', output)

        bm2name = self.gen_bookmark_name(name='share bookmark 2')
        json_output = json.loads(
            self.run_line(
                ('globus bookmark create -F json {}:{} "{}"'
                 ).format(GO_EP1_ID, '/share/dne/', bm2name)))
        self.assertEquals(json_output['name'], bm2name)
        self.assertEquals(json_output['path'], '/share/dne/')
        self.assertEquals(json_output['endpoint_id'], GO_EP1_ID)

    def test_bookmark_show(self):
        """
        Runs bookmark show on bm1's name and id.
        Confirms both inputs work, and verbose output is as expected.
        """
        # id
        output = self.run_line('globus bookmark show "{}"'.format(self.bm1id))
        self.assertEquals("{}:/home/\n".format(GO_EP1_ID), output)

        # name
        output = self.run_line(
            'globus bookmark show "{}"'.format(self.bm1name))
        self.assertEquals("{}:/home/\n".format(GO_EP1_ID), output)

        # verbose
        output = self.run_line("globus bookmark show -v {}".format(self.bm1id))
        self.assertIn("Endpoint ID: {}".format(GO_EP1_ID), output)

    def test_bookmark_rename_by_id(self):
        """
        Runs bookmark rename on bm1's id. Confirms can be shown by new name.
        """
        new_name = self.gen_bookmark_name(name="new_bm1")
        output = self.run_line(
            'globus bookmark rename "{}" "{}"'.format(self.bm1id, new_name))
        self.assertIn("Success", output)

        output = self.run_line('globus bookmark show -v "{}"'.format(new_name))
        self.assertIn("ID:          {}".format(self.bm1id), output)

    def test_bookmark_rename_by_name(self):
        """
        Runs bookmark rename on bm1's name. Confirms can be shown by new name.
        """
        new_name = self.gen_bookmark_name(name="new_bm1")
        output = self.run_line(
            'globus bookmark rename "{}" "{}"'.format(self.bm1name, new_name))
        self.assertIn("Success", output)

        output = self.run_line('globus bookmark show -v "{}"'.format(new_name))
        self.assertIn("ID:          {}".format(self.bm1id), output)

    def test_bookmark_delete_by_id(self):
        """
        Runs bookmark delete on bm1's id. Confirms success message.
        """
        output = self.run_line(
            'globus bookmark delete "{}"'.format(self.bm1id))
        self.assertIn("deleted successfully", output)

    def test_bookmark_delete_by_name(self):
        """
        Runs bookmark delete on bm1's name. Confirms success message.
        """
        output = self.run_line(
            'globus bookmark delete "{}"'.format(self.bm1name))
        self.assertIn("deleted successfully", output)
