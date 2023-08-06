# -*- coding: utf-8 -*-
#
# Copyright © 2017 Spyder Project Contributors
# Licensed under the terms of the MIT License
# (see LICENSE.txt for details)
"""Support for unittest framework."""

# Standard library imports
import re

# Local imports
from spyder_unittest.backend.runnerbase import Category, RunnerBase, TestResult


class UnittestRunner(RunnerBase):
    """Class for running tests with unittest module in standard library."""

    module = 'unittest'
    name = 'unittest'

    def create_argument_list(self):
        """Create argument list for testing process."""
        return ['-m', self.module, 'discover', '-v']

    def finished(self):
        """
        Called when the unit test process has finished.

        This function reads the results and emits `sig_finished`.
        """
        output = self.read_all_process_output()
        testresults = self.load_data(output)
        self.sig_finished.emit(testresults, output)

    def load_data(self, output):
        """
        Read and parse output from unittest module.

        Returns
        -------
        list of TestResult
            Unit test results.
        """
        res = []
        lines = output.splitlines()
        line_index = 0
        test_index = None

        while line_index < len(lines):
            data = self.try_parse_result(lines[line_index])
            if data:
                if data[2] == 'ok':
                    cat = Category.OK
                elif data[2] == 'FAIL' or data[2] == 'ERROR':
                    cat = Category.FAIL
                else:
                    cat = Category.SKIP
                name = '{}.{}'.format(data[1], data[0])
                tr = TestResult(category=cat, status=data[2], name=name,
                                message=data[3])
                res.append(tr)
                line_index += 1
                test_index = -1
                continue

            data = self.try_parse_exception_header(lines, line_index)
            if data:
                line_index = data[0]
                test_index = next(
                    i for i, tr in enumerate(res)
                    if tr.name == '{}.{}'.format(data[2], data[1]))

            data = self.try_parse_footer(lines, line_index)
            if data:
                line_index = data
                test_index = -1
                continue

            if test_index is not None:
                res[test_index].extra_text.append(lines[line_index] + '\n')
                line_index += 1

        return res

    def try_parse_result(self, line):
        """
        Try to parse a line of text as a test result.

        Returns
        -------
        tuple of str or None
            If line represents a test result, then return a tuple with four
            strings: the name of the test function, the name of the test class,
            the test result, and the reason (if no reason is given, the fourth
            string is empty). Otherwise, return None.
        """
        regexp = (r'([^\d\W]\w*) \(([^\d\W][\w.]*)\) \.\.\. '
                  '(ok|FAIL|ERROR|skipped|expected failure|unexpected success)'
                  "( '([^']*)')?\Z")
        match = re.match(regexp, line)
        if match:
            msg = match.groups()[4] or ''
            return match.groups()[:3] + (msg, )
        else:
            return None

    def try_parse_exception_header(self, lines, line_index):
        """
        Try to parse the header of an exception in unittest output.

        Returns
        -------
        (int, str, str) or None
            If an exception header is parsed successfully, then return a tuple
            with the new line index, the name of the test function, and the
            name of the test class. Otherwise, return None.
        """
        if lines[line_index] != '':
            return None
        if not all(char == '=' for char in lines[line_index + 1]):
            return None
        regexp = r'\w+: ([^\d\W]\w*) \(([^\d\W][\w.]*)\)\Z'
        match = re.match(regexp, lines[line_index + 2])
        if not match:
            return None
        if not all(char == '-' for char in lines[line_index + 3]):
            return None
        return (line_index + 4, ) + match.groups()

    def try_parse_footer(self, lines, line_index):
        """
        Try to parse footer of unittest output.

        Returns
        -------
        int or None
            New line index if footer is parsed successfully, None otherwise
        """
        if lines[line_index] != '':
            return None
        if not all(char == '-' for char in lines[line_index + 1]):
            return None
        if not re.match(r'^Ran [\d]+ tests? in', lines[line_index + 2]):
            return None
        if lines[line_index + 3] != '':
            return None
        return line_index + 5
