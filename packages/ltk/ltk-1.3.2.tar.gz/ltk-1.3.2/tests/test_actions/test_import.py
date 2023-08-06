from tests.test_actions import *
from ltk.actions.import_action import ImportAction
from ltk.actions.clean_action import CleanAction
from ltk.actions.rm_action import RmAction

import unittest

class TestImport(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        create_config()

    @classmethod
    def tearDownClass(cls):
        cleanup()

    def setUp(self):
        self.action = ImportAction(os.getcwd())
        self.clean_action = CleanAction(os.getcwd())
        self.clean_action.clean_action(False, False, None)
        self.files = ['sample.txt', 'sample1.txt', 'sample2.txt']
        for fn in self.files:
            create_txt_file(fn)
        self.doc_ids = []
        for fn in self.files:
            title = os.path.basename(os.path.normpath(fn))
            response = self.action.api.add_document(self.action.locale, fn, self.action.project_id, title)
            assert response.status_code == 202
            self.doc_ids.append(response.json()['properties']['id'])
        for doc_id in self.doc_ids:
            assert poll_doc(self.action, doc_id)
        for fn in self.files:
            delete_file(fn)
        self.imported = []

    def tearDown(self):
        self.rm_action = RmAction(os.getcwd())
        for doc_id in self.doc_ids:
            self.rm_action.rm_action(doc_id, id=True)
        for doc_id in self.imported:
            self.rm_action.rm_action(doc_id, id=True)
        self.clean_action.clean_action(True, False, None)
        self.rm_action.close()
        self.action.close()

    def test_import_all(self):
        files = os.listdir()
        self.action.import_action(True, True, None)
        for doc_id in self.doc_ids:
            doc = self.action.doc_manager.get_doc_by_prop('id', doc_id)
            assert doc
            self.imported.append(doc['id'])
        added_files = list(set(os.listdir()) - set(files))
        for file in added_files:
            os.remove(file)

    def test_import_locale(self):
        locale = "ja_JP"
        doc_id = self.doc_ids[0]
        response = self.action.api.document_add_target(doc_id, locale)
        assert response.status_code == 201
        self.action.import_action(False, True, None, doc_id)
        entry = self.action.doc_manager.get_doc_by_prop("id", doc_id)
        assert locale in entry["locales"]
        self.imported.append(entry['id'])

    def test_import_no_locale(self):
        self.action.import_action(False, True, None, self.doc_ids[0])
        entry = self.action.doc_manager.get_doc_by_prop("id", self.doc_ids[0])
        assert not entry.get("locales")
        self.imported.append(entry['id'])
