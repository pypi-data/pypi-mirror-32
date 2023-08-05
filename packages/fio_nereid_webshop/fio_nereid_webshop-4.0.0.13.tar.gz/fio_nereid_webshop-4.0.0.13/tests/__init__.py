# -*- coding: utf-8 -*-
import unittest

import trytond.tests.test_tryton

from .test_views_depends import TestViewsDepends
from .test_invoice import TestDownloadInvoice
from .test_css import CSSTest
from .test_templates import TestTemplates
from .test_gift_card import TestGiftCard
from .test_website import TestWebsite
from .test_tree import TestTree


def suite():
    """
    Define suite
    """
    test_suite = trytond.tests.test_tryton.suite()
    test_suite.addTests([
        unittest.TestLoader().loadTestsFromTestCase(CSSTest),
        unittest.TestLoader().loadTestsFromTestCase(TestViewsDepends),
        unittest.TestLoader().loadTestsFromTestCase(TestDownloadInvoice),
        unittest.TestLoader().loadTestsFromTestCase(TestTemplates),
        unittest.TestLoader().loadTestsFromTestCase(TestGiftCard),
        unittest.TestLoader().loadTestsFromTestCase(TestWebsite),
        unittest.TestLoader().loadTestsFromTestCase(TestTree),
    ])
    return test_suite


if __name__ == '__main__':
    unittest.TextTestRunner(verbosity=2).run(suite())
