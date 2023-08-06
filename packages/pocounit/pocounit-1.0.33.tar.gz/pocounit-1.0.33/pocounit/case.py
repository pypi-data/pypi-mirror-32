# coding=utf-8

import re
import six
import traceback
import unittest
import warnings

from unittest.util import safe_repr

from pocounit.fixture import FixtureUnit
from pocounit.result import PocoTestResult
from pocounit.result.metainfo import MetaInfo
from pocounit.result.collector import PocoResultCollector
from pocounit.result.trace import ScriptTracer
from pocounit.result.record import ScreenRecorder
from pocounit.result.action import ActionRecorder
from pocounit.result.app_runtime import AppRuntimeLog
from pocounit.result.runner_runtime import RunnerRuntimeLog
from pocounit.result.assertion import AssertionRecorder
from pocounit.result.site_snapshot import SiteSnapshot


SPECIAL_CHARS = re.compile(r'[\/\\\.:*?"<>|]')


class PocoTestCase(unittest.TestCase, FixtureUnit):
    """
    * longMessage: determines whether long messages (including repr of
        objects used in assert methods) will be printed on failure in *addition*
        to any explicit message passed.
    """
    longMessage = True

    _resule_collector = None
    _result_emitters = None     # {name -> PocoTestResultEmitter}
    _addons = None              # [addons]

    def __init__(self):
        unittest.TestCase.__init__(self)
        FixtureUnit.__init__(self)

        # check testcase name
        self.testcase_name = re.sub(SPECIAL_CHARS, lambda g: '-', self.name())

        collector = PocoResultCollector(self.project_root, [self.test_case_filename], self.testcase_name, self.test_case_dir)
        self.set_result_collector(collector)

        meta_info_emitter = MetaInfo(collector)
        runner_runtime_log = RunnerRuntimeLog(collector)
        tracer = ScriptTracer(collector)
        screen_recorder = ScreenRecorder(collector)
        action_recorder = ActionRecorder(collector)
        app_runtime_log = AppRuntimeLog(collector)
        assertion_recorder = AssertionRecorder(collector)
        site_snapshot = SiteSnapshot(collector)

        self.add_result_emitter('metaInfo', meta_info_emitter)
        self.add_result_emitter('runnerRuntimeLog', runner_runtime_log)
        self.add_result_emitter('tracer', tracer)
        self.add_result_emitter('screenRecorder', screen_recorder)
        self.add_result_emitter('actionRecorder', action_recorder)
        self.add_result_emitter('appRuntimeLog', app_runtime_log)
        self.add_result_emitter('assertionRecorder', assertion_recorder)
        self.add_result_emitter('siteSnapshot', site_snapshot)

        self.meta_info_emitter = meta_info_emitter

    @classmethod
    def name(cls):
        """
        改写此方法来自定义testcase的名字，允许中文，请使用utf-8编码的str或者unicode

        :return:
        """

        if type(cls) is type:
            return cls.__name__
        else:
            return cls.__class__.__name__

    @classmethod
    def getMetaInfo(cls):
        return {}

    @classmethod
    def setUpClass(cls):
        """
        用例预处理，加载自定义插件等

        :return: 
        """
        pass

    def runTest(self):
        """
        testcase的正文，把要执行的测试和包括断言语句都写到这里

        """

        raise NotImplementedError

    @classmethod
    def tearDownClass(cls):
        """
        用例后处理

        :return: 
        """
        pass

    def shortDescription(self):
        """
        描述这个testcase的细节，如果需要的话就override这个方法
        
        :return: <str>
        """

        return super(PocoTestCase, self).shortDescription()

    def defaultTestResult(self):
        return PocoTestResult()

    @classmethod
    def set_result_collector(cls, collector):
        if type(cls) is not type:
            cls = cls.__class__
        cls._resule_collector = collector

    @classmethod
    def get_result_collector(cls):
        return cls._resule_collector

    @classmethod
    def add_result_emitter(cls, name, emitter):
        if type(cls) is not type:
            cls = cls.__class__
        if not cls._result_emitters:
            cls._result_emitters = {}
        cls._result_emitters[name] = emitter

    @classmethod
    def get_result_emitter(cls, name):
        return cls._result_emitters.get(name)

    @classmethod
    def register_addon(cls, addon):
        if type(cls) is not type:
            cls = cls.__class__
        if not cls._addons:
            cls._addons = []
        cls._addons.append(addon)

    # deprecation warning
    @classmethod
    def register_addin(cls, v):
        warnings.warn('`register_addin` is deprecated. Please use `register_addon` instead.')
        return cls.register_addon(v)

    def run(self, result=None):
        result = result or self.defaultTestResult()
        if not isinstance(result, PocoTestResult):
            raise TypeError('Test result class should be subclass of PocoTestResult. '
                            'Current test result instance is "{}".'.format(result))

        # 自动把当前脚本add到collector的脚本watch列表里
        collector = self.get_result_collector()
        collector.add_testcase_file(self.test_case_filename)

        self.meta_info_emitter.set_testcase_metainfo(self.testcase_name, self.getMetaInfo())

        # register addon
        if not self.__class__._addons:
            self.__class__._addons = []
        for addon in self._addons:
            addon.initialize(self)

        self.meta_info_emitter.test_started(self.testcase_name)

        # start result emitter
        for name, emitter in self._result_emitters.items():
            try:
                emitter.start()
            except Exception as e:
                warnings.warn('Fail to start result emitter: "{}". You can report this error to the developers or just '
                              'ignore it. Error message: \n"{}"'
                              .format(emitter.__class__.__name__, traceback.format_exc()))

        # run test
        # super.run(...) will never raise
        ret = super(PocoTestCase, self).run(result)

        assertionRecorder = self.get_result_emitter('assertionRecorder')
        for _, exc_type, e, tb in result.detail_errors:
            assertionRecorder.traceback(exc_type, e, tb)

        # stop result emitter
        for name, emitter in self._result_emitters.items():
            try:
                emitter.stop()
            except Exception as e:
                warnings.warn('Fail to stop result emitter: "{}". You can report this error to the developers or just '
                              'ignore it. Error message: \n"{}"'
                              .format(emitter.__class__.__name__, traceback.format_exc()))

        self.meta_info_emitter.test_ended(self.testcase_name)

        # handle result
        if result.detail_errors or result.errors or result.failures:
            self.meta_info_emitter.test_fail(self.testcase_name)
        else:
            self.meta_info_emitter.test_succeed(self.testcase_name)
        return ret

    def fail(self, fail_msg=None):
        e = self.failureException(fail_msg)
        errmsg = six.text_type(e)
        err_typename = self.failureException.__name__
        assertionRecorder = self.get_result_emitter('assertionRecorder')
        assertionRecorder.fail(fail_msg, errmsg, 'Traceback: (failure cause)\n{}{}: {}'
                               .format(''.join(traceback.format_stack()), err_typename, fail_msg))
        raise super(PocoTestCase, self).fail(fail_msg)

    # assertions
    def assertTrue(self, expr, msg=None):
        expr = bool(expr)
        assertionRecorder = self.get_result_emitter('assertionRecorder')
        assertionRecorder.assert_('True', [expr], msg)

        if not expr:
            msg = self._formatMessage(msg, "%s is not true" % safe_repr(expr))
            self.fail(msg)

    def assertFalse(self, expr, msg=None):
        expr = bool(expr)
        assertionRecorder = self.get_result_emitter('assertionRecorder')
        assertionRecorder.assert_('False', [expr], msg)

        if expr:
            msg = self._formatMessage(msg, "%s is not false" % safe_repr(expr))
            self.fail(msg)

    def assertEqual(self, first, second, msg=None):
        assertionRecorder = self.get_result_emitter('assertionRecorder')
        assertionRecorder.assert_('Equal', [first, second], msg)
        return super(PocoTestCase, self).assertEqual(first, second, msg)

    def assertNotEqual(self, first, second, msg=None):
        assertionRecorder = self.get_result_emitter('assertionRecorder')
        assertionRecorder.assert_('NotEqual', [first, second], msg)
        return super(PocoTestCase, self).assertNotEqual(first, second, msg)

    def assertLess(self, a, b, msg=None):
        assertionRecorder = self.get_result_emitter('assertionRecorder')
        assertionRecorder.assert_('Less', [a, b], msg)
        return super(PocoTestCase, self).assertLess(a, b, msg)

    def assertLessEqual(self, a, b, msg=None):
        assertionRecorder = self.get_result_emitter('assertionRecorder')
        assertionRecorder.assert_('LessEqual', [a, b], msg)
        return super(PocoTestCase, self).assertLessEqual(a, b, msg)

    def assertGreater(self, a, b, msg=None):
        assertionRecorder = self.get_result_emitter('assertionRecorder')
        assertionRecorder.assert_('Greater', [a, b], msg)
        return super(PocoTestCase, self).assertGreater(a, b, msg)

    def assertGreaterEqual(self, a, b, msg=None):
        assertionRecorder = self.get_result_emitter('assertionRecorder')
        assertionRecorder.assert_('GreaterEqual', [a, b], msg)
        return super(PocoTestCase, self).assertGreaterEqual(a, b, msg)

    def assertIn(self, member, container, msg=None):
        assertionRecorder = self.get_result_emitter('assertionRecorder')
        assertionRecorder.assert_('In', [member, container], msg)
        return super(PocoTestCase, self).assertIn(member, container, msg)

    def assertNotIn(self, member, container, msg=None):
        assertionRecorder = self.get_result_emitter('assertionRecorder')
        assertionRecorder.assert_('NotIn', [member, container], msg)
        return super(PocoTestCase, self).assertNotIn(member, container, msg)
