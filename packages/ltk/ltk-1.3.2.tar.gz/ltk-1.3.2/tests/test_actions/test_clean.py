from tests.test_actions import *
from ltk.actions.clean_action import *
from ltk.actions.rm_action import RmAction
import subprocess

import unittest

class TestClean(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        create_config()

    @classmethod
    def tearDownClass(cls):
        cleanup()

    def setUp(self):
        self.action = CleanAction(os.getcwd())
        self.action.clean_action(False, False, None)
        self.files = ['sample.txt', 'sample1.txt', 'sample2.txt']
        self.forced = []
        for fn in self.files:
            create_txt_file(fn)
        subprocess.call(['ltk','add','-o','sample*.txt'])
        self.entries = self.action.doc_manager.get_all_entries()
        for entry in self.entries:
            assert poll_doc(self.action, entry['id'])

    def tearDown(self):
        self.rm_action = RmAction(os.getcwd())
        for curr_file in self.files:
            if curr_file in self.forced:
                continue
            self.rm_action.rm_action(curr_file, force=True)
        self.action.clean_action(True, False, None)
        self.action.close()

    def test_clean(self):
        delete_id = self.entries[0]['id']
        r = self.action.api.document_delete(delete_id)
        self.forced.append(self.entries[0]['file_name'])
        assert r.status_code == 204
        assert self.action.doc_manager.get_doc_by_prop('id', delete_id)
        self.action.clean_action(False, False, None)
        assert not self.action.doc_manager.get_doc_by_prop('id', delete_id)

    def test_clean_force(self):
        delete_id = self.entries[0]['id']
        doc_name = self.action.doc_manager.get_doc_by_prop('id', delete_id)['file_name']
        r = self.action.api.document_delete(delete_id)
        assert r.status_code == 204
        assert self.action.doc_manager.get_doc_by_prop('id', delete_id)
        self.action.clean_action(True, False, None)
        self.forced.append(self.entries[0]['file_name'])
        assert not self.action.doc_manager.get_doc_by_prop('id', delete_id)
        assert not os.path.isfile(os.path.join(self.action.path, doc_name))

    def test_disassociate(self):
        self.action.clean_action(False, True, None)
        for entry in self.entries:
            assert not self.action.doc_manager.get_doc_by_prop('id', entry['id'])
            delete_file(entry['file_name'])
            self.action.api.document_delete(entry['id'])
            self.forced.append(entry['file_name'])

    # Test that a specified file is disassociated
    def test_clean_single(self):
        delete_id = self.entries[0]['id']
        doc_name = self.entries[0]['file_name']
        self.action.clean_action(False, False, doc_name)
        delete_file(doc_name)
        self.action.api.document_delete(delete_id)
        self.forced.append(doc_name)
        assert not self.action.doc_manager.get_doc_by_prop('id', delete_id)
