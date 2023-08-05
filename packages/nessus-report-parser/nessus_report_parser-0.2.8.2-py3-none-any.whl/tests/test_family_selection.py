import unittest

from lxml import etree

from nessus_report_parser import FamilySelection, FamilyItem


class TestFamilySelection(unittest.TestCase):
    def test_well_formed_node(self):
        node = "<FamilySelection><FamilyItem>" \
               "<FamilyName>MacOS X Local Security Checks</FamilyName>" \
               "<Status>enabled</Status></FamilyItem>" \
               "<FamilyItem><FamilyName>Incident Response</FamilyName>" \
               "<Status>enabled</Status></FamilyItem>" \
               "<FamilyItem>" \
               "<FamilyName>F5 Networks Local Security Checks</FamilyName>" \
               "<Status>enabled</Status></FamilyItem>" \
               "<FamilyItem><FamilyName>DNS</FamilyName>" \
               "<Status>enabled</Status></FamilyItem></FamilySelection>"

        expected_seq = [
            {
                'family_name': 'MacOS X Local Security Checks',
                'status': 'enabled'
            },
            {
                'family_name': 'Incident Response',
                'status': 'enabled'
            },
            {
                'family_name': 'F5 Networks Local Security Checks',
                'status': 'enabled'
            },
            {
                'family_name': 'DNS',
                'status': 'enabled'
            },
        ]

        self.assertEqual([FamilyItem(doc) for doc in expected_seq],
                         FamilySelection.from_etree(etree.XML(node))
                         )

    def test_family_selection(self):
        f = FamilySelection([FamilyItem({'family_name': 'f1', 'status': 's1'}),
                             FamilyItem({'family_name': 'f2', 'status': 's2'})])
        expected = [{'family_name': 'f1', 'status': 's1'},
                    {'family_name': 'f2', 'status': 's2'}]

        self.assertEqual(expected, f.data)
