package unittest_result
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/unittest_result.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßFalse := πg.InternStr("False")
		ßNone := πg.InternStr("None")
		ßSTDERR_LINE := πg.InternStr("STDERR_LINE")
		ßSTDOUT_LINE := πg.InternStr("STDOUT_LINE")
		ßStringIO := πg.InternStr("StringIO")
		ßTestResult := πg.InternStr("TestResult")
		ßTrue := πg.InternStr("True")
		ß_StringIO := πg.InternStr("_StringIO")
		ß__class__ := πg.InternStr("__class__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__repr__ := πg.InternStr("__repr__")
		ß__unittest := πg.InternStr("__unittest")
		ß_count_relevant_tb_levels := πg.InternStr("_count_relevant_tb_levels")
		ß_exc_info_to_string := πg.InternStr("_exc_info_to_string")
		ß_is_relevant_tb_level := πg.InternStr("_is_relevant_tb_level")
		ß_mirrorOutput := πg.InternStr("_mirrorOutput")
		ß_moduleSetUpFailed := πg.InternStr("_moduleSetUpFailed")
		ß_original_stderr := πg.InternStr("_original_stderr")
		ß_original_stdout := πg.InternStr("_original_stdout")
		ß_previousTestClass := πg.InternStr("_previousTestClass")
		ß_restoreStdout := πg.InternStr("_restoreStdout")
		ß_setupStdout := πg.InternStr("_setupStdout")
		ß_stderr_buffer := πg.InternStr("_stderr_buffer")
		ß_stdout_buffer := πg.InternStr("_stdout_buffer")
		ß_testRunEntered := πg.InternStr("_testRunEntered")
		ßaddError := πg.InternStr("addError")
		ßaddExpectedFailure := πg.InternStr("addExpectedFailure")
		ßaddFailure := πg.InternStr("addFailure")
		ßaddSkip := πg.InternStr("addSkip")
		ßaddSuccess := πg.InternStr("addSuccess")
		ßaddUnexpectedSuccess := πg.InternStr("addUnexpectedSuccess")
		ßappend := πg.InternStr("append")
		ßbuffer := πg.InternStr("buffer")
		ßendswith := πg.InternStr("endswith")
		ßerrors := πg.InternStr("errors")
		ßexpectedFailures := πg.InternStr("expectedFailures")
		ßf_globals := πg.InternStr("f_globals")
		ßfailfast := πg.InternStr("failfast")
		ßfailureException := πg.InternStr("failureException")
		ßfailures := πg.InternStr("failures")
		ßformat_exception := πg.InternStr("format_exception")
		ßfunctools := πg.InternStr("functools")
		ßgetattr := πg.InternStr("getattr")
		ßgetvalue := πg.InternStr("getvalue")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßobject := πg.InternStr("object")
		ßos := πg.InternStr("os")
		ßprintErrors := πg.InternStr("printErrors")
		ßseek := πg.InternStr("seek")
		ßshouldStop := πg.InternStr("shouldStop")
		ßskipped := πg.InternStr("skipped")
		ßstartTest := πg.InternStr("startTest")
		ßstartTestRun := πg.InternStr("startTestRun")
		ßstderr := πg.InternStr("stderr")
		ßstdout := πg.InternStr("stdout")
		ßstop := πg.InternStr("stop")
		ßstopTest := πg.InternStr("stopTest")
		ßstopTestRun := πg.InternStr("stopTestRun")
		ßstrclass := πg.InternStr("strclass")
		ßsys := πg.InternStr("sys")
		ßtb_frame := πg.InternStr("tb_frame")
		ßtb_next := πg.InternStr("tb_next")
		ßtestsRun := πg.InternStr("testsRun")
		ßtraceback := πg.InternStr("traceback")
		ßtruncate := πg.InternStr("truncate")
		ßunexpectedSuccesses := πg.InternStr("unexpectedSuccesses")
		ßutil := πg.InternStr("util")
		ßwasSuccessful := πg.InternStr("wasSuccessful")
		ßwraps := πg.InternStr("wraps")
		ßwrite := πg.InternStr("write")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 []πg.Param
		_ = πTemp004
		var πTemp005 *πg.Dict
		_ = πTemp005
		var πTemp006 *πg.Object
		_ = πTemp006
		var πTemp007 *πg.Object
		_ = πTemp007
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """Test result object"""
			πF.SetLineno(1)
			// line 3: import os
			πF.SetLineno(3)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 4: import sys
			πF.SetLineno(4)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 5: import traceback
			πF.SetLineno(5)
			if πTemp002, πE = πg.ImportModule(πF, "traceback"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtraceback.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 8: import StringIO as _StringIO
			πF.SetLineno(8)
			if πTemp002, πE = πg.ImportModule(πF, "StringIO"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_StringIO.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 9: StringIO = _StringIO.StringIO
			πF.SetLineno(9)
			if πTemp001, πE = πg.ResolveGlobal(πF, ß_StringIO); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßStringIO, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStringIO.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 13: import unittest_util as util
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "unittest_util"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßutil.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 14: import functools
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "functools"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßfunctools.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 15: wraps = functools.wraps
			πF.SetLineno(15)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßfunctools); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßwraps, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßwraps.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 17: __unittest = True
			πF.SetLineno(17)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__unittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 19: def failfast(method):
			πF.SetLineno(19)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "method", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("failfast", "build/src/__python__/unittest_result.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmethod *πg.Object = πArgs[0]; _ = µmethod
				var µinner *πg.Object = πg.UnboundLocal; _ = µinner
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 21: def inner(self, *args, **kw):
					πF.SetLineno(21)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("inner", "build/src/__python__/unittest_result.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µargs *πg.Object = πArgs[1]; _ = µargs
						var µkw *πg.Object = πArgs[2]; _ = µkw
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							πTemp001[1] = ßfailfast.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 22: if getattr(self, 'failfast', False):
							πF.SetLineno(22)
						Label1:
							// line 23: self.stop()
							πF.SetLineno(23)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstop, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 24: return method(self, *args, **kw)
							πF.SetLineno(24)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkw, "kw"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmethod, "method"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, µmethod, πTemp001, µargs, nil, µkw); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp002
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					µinner = πTemp001
					// line 25: inner = wraps(method)(inner)
					πF.SetLineno(25)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µinner, "inner"); πE != nil {
						continue
					}
					πTemp003[0] = µinner
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmethod, "method"); πE != nil {
						continue
					}
					πTemp004[0] = µmethod
					if πTemp005, πE = πg.ResolveGlobal(πF, ßwraps); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp005, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µinner = πTemp005
					// line 26: return inner
					πF.SetLineno(26)
					if πE = πg.CheckLocal(πF, µinner, "inner"); πE != nil {
						continue
					}
					πR = µinner
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßfailfast.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 28: STDOUT_LINE = '\nStdout:\n%s'
			πF.SetLineno(28)
			if πE = πF.Globals().SetItem(πF, ßSTDOUT_LINE.ToObject(), πg.NewStr("\nStdout:\n%s").ToObject()); πE != nil {
				continue
			}
			// line 29: STDERR_LINE = '\nStderr:\n%s'
			πF.SetLineno(29)
			if πE = πF.Globals().SetItem(πF, ßSTDERR_LINE.ToObject(), πg.NewStr("\nStderr:\n%s").ToObject()); πE != nil {
				continue
			}
			// line 32: class TestResult(object):
			πF.SetLineno(32)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp005 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TestResult", "build/src/__python__/unittest_result.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 []*πg.Object
				_ = πTemp011
				var πTemp012 *πg.Object
				_ = πTemp012
				var πTemp013 *πg.Object
				_ = πTemp013
				var πTemp014 *πg.Object
				_ = πTemp014
				var πTemp015 *πg.Object
				_ = πTemp015
				var πTemp016 *πg.Object
				_ = πTemp016
				var πTemp017 *πg.Object
				_ = πTemp017
				var πTemp018 *πg.Object
				_ = πTemp018
				var πTemp019 *πg.Object
				_ = πTemp019
				var πTemp020 *πg.Object
				_ = πTemp020
				var πTemp021 *πg.Object
				_ = πTemp021
				var πTemp022 *πg.Object
				_ = πTemp022
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 33: """Holder for test result information.
					πF.SetLineno(33)
					// line 43: _previousTestClass = None
					πF.SetLineno(43)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß_previousTestClass.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 44: _testRunEntered = False
					πF.SetLineno(44)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß_testRunEntered.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 45: _moduleSetUpFailed = False
					πF.SetLineno(45)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß_moduleSetUpFailed.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 46: def __init__(self, stream=None, descriptions=None, verbosity=None):
					πF.SetLineno(46)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "stream", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "descriptions", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "verbosity", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/unittest_result.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstream *πg.Object = πArgs[1]; _ = µstream
						var µdescriptions *πg.Object = πArgs[2]; _ = µdescriptions
						var µverbosity *πg.Object = πArgs[3]; _ = µverbosity
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 47: self.failfast = False
							πF.SetLineno(47)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfailfast, πTemp002); πE != nil {
								continue
							}
							// line 48: self.failures = []
							πF.SetLineno(48)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfailures, πTemp002); πE != nil {
								continue
							}
							// line 49: self.errors = []
							πF.SetLineno(49)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßerrors, πTemp002); πE != nil {
								continue
							}
							// line 50: self.testsRun = 0
							πF.SetLineno(50)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtestsRun, πTemp001); πE != nil {
								continue
							}
							// line 51: self.skipped = []
							πF.SetLineno(51)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßskipped, πTemp002); πE != nil {
								continue
							}
							// line 52: self.expectedFailures = []
							πF.SetLineno(52)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßexpectedFailures, πTemp002); πE != nil {
								continue
							}
							// line 53: self.unexpectedSuccesses = []
							πF.SetLineno(53)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßunexpectedSuccesses, πTemp002); πE != nil {
								continue
							}
							// line 54: self.shouldStop = False
							πF.SetLineno(54)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßshouldStop, πTemp002); πE != nil {
								continue
							}
							// line 55: self.buffer = False
							πF.SetLineno(55)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbuffer, πTemp002); πE != nil {
								continue
							}
							// line 56: self._stdout_buffer = None
							πF.SetLineno(56)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_stdout_buffer, πTemp002); πE != nil {
								continue
							}
							// line 57: self._stderr_buffer = None
							πF.SetLineno(57)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_stderr_buffer, πTemp002); πE != nil {
								continue
							}
							// line 58: self._original_stdout = sys.stdout
							πF.SetLineno(58)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstdout, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_original_stdout, πTemp001); πE != nil {
								continue
							}
							// line 59: self._original_stderr = sys.stderr
							πF.SetLineno(59)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstderr, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_original_stderr, πTemp001); πE != nil {
								continue
							}
							// line 60: self._mirrorOutput = False
							πF.SetLineno(60)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_mirrorOutput, πTemp002); πE != nil {
								continue
							}
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 62: def printErrors(self):
					πF.SetLineno(62)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("printErrors", "build/src/__python__/unittest_result.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 63: "Called by TestRunner after test run"
							πF.SetLineno(63)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßprintErrors.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 65: def startTest(self, test):
					πF.SetLineno(65)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "test", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("startTest", "build/src/__python__/unittest_result.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtest *πg.Object = πArgs[1]; _ = µtest
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 66: "Called when the given test is about to be run"
							πF.SetLineno(66)
							// line 67: self.testsRun += 1
							πF.SetLineno(67)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtestsRun, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtestsRun, πTemp002); πE != nil {
								continue
							}
							// line 68: self._mirrorOutput = False
							πF.SetLineno(68)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_mirrorOutput, πTemp002); πE != nil {
								continue
							}
							// line 69: self._setupStdout()
							πF.SetLineno(69)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_setupStdout, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßstartTest.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 71: def _setupStdout(self):
					πF.SetLineno(71)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("_setupStdout", "build/src/__python__/unittest_result.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbuffer, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 72: if self.buffer:
							πF.SetLineno(72)
						Label1:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_stderr_buffer, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003 == πTemp004).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 73: if self._stderr_buffer is None:
							πF.SetLineno(73)
						Label3:
							// line 74: self._stderr_buffer = StringIO()
							πF.SetLineno(74)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_stderr_buffer, πTemp001); πE != nil {
								continue
							}
							// line 75: self._stdout_buffer = StringIO()
							πF.SetLineno(75)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_stdout_buffer, πTemp001); πE != nil {
								continue
							}
							goto Label4
						Label4:
							// line 76: sys.stdout = self._stdout_buffer
							πF.SetLineno(76)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stdout_buffer, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp004, ßstdout, πTemp003); πE != nil {
								continue
							}
							// line 77: sys.stderr = self._stderr_buffer
							πF.SetLineno(77)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stderr_buffer, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp004, ßstderr, πTemp003); πE != nil {
								continue
							}
							goto Label2
						Label2:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_setupStdout.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 79: def startTestRun(self):
					πF.SetLineno(79)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("startTestRun", "build/src/__python__/unittest_result.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 80: """Called once before any tests are executed.
							πF.SetLineno(80)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßstartTestRun.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 85: def stopTest(self, test):
					πF.SetLineno(85)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "test", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("stopTest", "build/src/__python__/unittest_result.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtest *πg.Object = πArgs[1]; _ = µtest
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 86: """Called when the given test has been run"""
							πF.SetLineno(86)
							// line 87: self._restoreStdout()
							πF.SetLineno(87)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_restoreStdout, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 88: self._mirrorOutput = False
							πF.SetLineno(88)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_mirrorOutput, πTemp002); πE != nil {
								continue
							}
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßstopTest.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 90: def _restoreStdout(self):
					πF.SetLineno(90)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("_restoreStdout", "build/src/__python__/unittest_result.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µoutput *πg.Object = πg.UnboundLocal; _ = µoutput
						var µerror *πg.Object = πg.UnboundLocal; _ = µerror
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbuffer, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 91: if self.buffer:
							πF.SetLineno(91)
						Label1:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_mirrorOutput, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 92: if self._mirrorOutput:
							πF.SetLineno(92)
						Label3:
							// line 93: output = sys.stdout.getvalue()
							πF.SetLineno(93)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßstdout, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßgetvalue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µoutput = πTemp003
							// line 94: error = sys.stderr.getvalue()
							πF.SetLineno(94)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßstderr, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßgetvalue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µerror = πTemp003
							if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µoutput); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label5
							}
							goto Label6
							// line 95: if output:
							πF.SetLineno(95)
						Label5:
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("\n").ToObject()
							if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µoutput, ßendswith, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label7
							}
							goto Label8
							// line 96: if not output.endswith('\n'):
							πF.SetLineno(96)
						Label7:
							// line 97: output += '\n'
							πF.SetLineno(97)
							if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µoutput, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							µoutput = πTemp001
							goto Label8
						Label8:
							// line 98: self._original_stdout.write(STDOUT_LINE % output)
							πF.SetLineno(98)
							πTemp004 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSTDOUT_LINE); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πTemp003, µoutput); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_original_stdout, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µerror); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label9
							}
							goto Label10
							// line 99: if error:
							πF.SetLineno(99)
						Label9:
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("\n").ToObject()
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µerror, ßendswith, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label11
							}
							goto Label12
							// line 100: if not error.endswith('\n'):
							πF.SetLineno(100)
						Label11:
							// line 101: error += '\n'
							πF.SetLineno(101)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µerror, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							µerror = πTemp001
							goto Label12
						Label12:
							// line 102: self._original_stderr.write(STDERR_LINE % error)
							πF.SetLineno(102)
							πTemp004 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSTDERR_LINE); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πTemp003, µerror); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_original_stderr, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label10
						Label10:
							goto Label4
						Label4:
							// line 104: sys.stdout = self._original_stdout
							πF.SetLineno(104)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_original_stdout, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp005, ßstdout, πTemp003); πE != nil {
								continue
							}
							// line 105: sys.stderr = self._original_stderr
							πF.SetLineno(105)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_original_stderr, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp005, ßstderr, πTemp003); πE != nil {
								continue
							}
							// line 106: self._stdout_buffer.seek(0)
							πF.SetLineno(106)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stdout_buffer, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßseek, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 107: self._stdout_buffer.truncate()
							πF.SetLineno(107)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stdout_buffer, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßtruncate, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 108: self._stderr_buffer.seek(0)
							πF.SetLineno(108)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stderr_buffer, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßseek, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 109: self._stderr_buffer.truncate()
							πF.SetLineno(109)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stderr_buffer, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßtruncate, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label2
						Label2:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_restoreStdout.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 111: def stopTestRun(self):
					πF.SetLineno(111)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("stopTestRun", "build/src/__python__/unittest_result.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 112: """Called once after all tests are executed.
							πF.SetLineno(112)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßstopTestRun.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 118: def addError(self, test, err):
					πF.SetLineno(118)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "test", Def: nil}
					πTemp002[2] = πg.Param{Name: "err", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("addError", "build/src/__python__/unittest_result.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtest *πg.Object = πArgs[1]; _ = µtest
						var µerr *πg.Object = πArgs[2]; _ = µerr
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 119: """Called when an error has occurred. 'err' is a tuple of values as
							πF.SetLineno(119)
							// line 122: self.errors.append((test, self._exc_info_to_string(err, test)))
							πF.SetLineno(122)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp003[0] = µerr
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp003[1] = µtest
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_exc_info_to_string, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002 = πg.NewTuple2(µtest, πTemp005).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßerrors, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 123: self._mirrorOutput = True
							πF.SetLineno(123)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_mirrorOutput, πTemp004); πE != nil {
								continue
							}
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßaddError.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 124: addError = failfast(addError)
					πF.SetLineno(124)
					πTemp011 = πF.MakeArgs(1)
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßaddError); πE != nil {
						continue
					}
					πTemp011[0] = πTemp012
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßfailfast); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp012.Call(πF, πTemp011, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp011)
					if πE = πClass.SetItem(πF, ßaddError.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 127: def addFailure(self, test, err):
					πF.SetLineno(127)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "test", Def: nil}
					πTemp002[2] = πg.Param{Name: "err", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("addFailure", "build/src/__python__/unittest_result.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtest *πg.Object = πArgs[1]; _ = µtest
						var µerr *πg.Object = πArgs[2]; _ = µerr
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 128: """Called when an error has occurred. 'err' is a tuple of values as
							πF.SetLineno(128)
							// line 130: self.failures.append((test, self._exc_info_to_string(err, test)))
							πF.SetLineno(130)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp003[0] = µerr
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp003[1] = µtest
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_exc_info_to_string, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002 = πg.NewTuple2(µtest, πTemp005).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfailures, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 131: self._mirrorOutput = True
							πF.SetLineno(131)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_mirrorOutput, πTemp004); πE != nil {
								continue
							}
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßaddFailure.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 132: addFailure = failfast(addFailure)
					πF.SetLineno(132)
					πTemp011 = πF.MakeArgs(1)
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßaddFailure); πE != nil {
						continue
					}
					πTemp011[0] = πTemp013
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßfailfast); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp013.Call(πF, πTemp011, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp011)
					if πE = πClass.SetItem(πF, ßaddFailure.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 134: def addSuccess(self, test):
					πF.SetLineno(134)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "test", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("addSuccess", "build/src/__python__/unittest_result.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtest *πg.Object = πArgs[1]; _ = µtest
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 135: "Called when a test has completed successfully"
							πF.SetLineno(135)
							// line 136: pass
							πF.SetLineno(136)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßaddSuccess.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 138: def addSkip(self, test, reason):
					πF.SetLineno(138)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "test", Def: nil}
					πTemp002[2] = πg.Param{Name: "reason", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("addSkip", "build/src/__python__/unittest_result.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtest *πg.Object = πArgs[1]; _ = µtest
						var µreason *πg.Object = πArgs[2]; _ = µreason
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 139: """Called when a test is skipped."""
							πF.SetLineno(139)
							// line 140: self.skipped.append((test, reason))
							πF.SetLineno(140)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µreason, "reason"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µtest, µreason).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßskipped, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßaddSkip.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 142: def addExpectedFailure(self, test, err):
					πF.SetLineno(142)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "test", Def: nil}
					πTemp002[2] = πg.Param{Name: "err", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("addExpectedFailure", "build/src/__python__/unittest_result.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtest *πg.Object = πArgs[1]; _ = µtest
						var µerr *πg.Object = πArgs[2]; _ = µerr
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 143: """Called when an expected failure/error occurred."""
							πF.SetLineno(143)
							// line 144: self.expectedFailures.append(
							πF.SetLineno(144)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp003[0] = µerr
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp003[1] = µtest
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_exc_info_to_string, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002 = πg.NewTuple2(µtest, πTemp005).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßexpectedFailures, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßaddExpectedFailure.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 148: def addUnexpectedSuccess(self, test):
					πF.SetLineno(148)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "test", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("addUnexpectedSuccess", "build/src/__python__/unittest_result.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtest *πg.Object = πArgs[1]; _ = µtest
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 149: """Called when a test was expected to fail, but succeed."""
							πF.SetLineno(149)
							// line 150: self.unexpectedSuccesses.append(test)
							πF.SetLineno(150)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp001[0] = µtest
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßunexpectedSuccesses, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßaddUnexpectedSuccess.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 151: addUnexpectedSuccess = failfast(addUnexpectedSuccess)
					πF.SetLineno(151)
					πTemp011 = πF.MakeArgs(1)
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßaddUnexpectedSuccess); πE != nil {
						continue
					}
					πTemp011[0] = πTemp017
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßfailfast); πE != nil {
						continue
					}
					if πTemp018, πE = πTemp017.Call(πF, πTemp011, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp011)
					if πE = πClass.SetItem(πF, ßaddUnexpectedSuccess.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 153: def wasSuccessful(self):
					πF.SetLineno(153)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("wasSuccessful", "build/src/__python__/unittest_result.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 154: "Tells whether or not this result was a success"
							πF.SetLineno(154)
							// line 155: return len(self.failures) == len(self.errors) == 0
							πF.SetLineno(155)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßfailures, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßerrors, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Eq(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label1
							}
							if πTemp001, πE = πg.Eq(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
						Label1:
							πR = πTemp001
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßwasSuccessful.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 157: def stop(self):
					πF.SetLineno(157)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("stop", "build/src/__python__/unittest_result.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 158: "Indicates that the tests should be aborted"
							πF.SetLineno(158)
							// line 159: self.shouldStop = True
							πF.SetLineno(159)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßshouldStop, πTemp002); πE != nil {
								continue
							}
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßstop.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 161: def _exc_info_to_string(self, err, test):
					πF.SetLineno(161)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "err", Def: nil}
					πTemp002[2] = πg.Param{Name: "test", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("_exc_info_to_string", "build/src/__python__/unittest_result.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µerr *πg.Object = πArgs[1]; _ = µerr
						var µtest *πg.Object = πArgs[2]; _ = µtest
						var µexctype *πg.Object = πg.UnboundLocal; _ = µexctype
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
						var µtb *πg.Object = πg.UnboundLocal; _ = µtb
						var µlength *πg.Object = πg.UnboundLocal; _ = µlength
						var µmsgLines *πg.Object = πg.UnboundLocal; _ = µmsgLines
						var µoutput *πg.Object = πg.UnboundLocal; _ = µoutput
						var µerror *πg.Object = πg.UnboundLocal; _ = µerror
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 162: """Converts a sys.exc_info()-style tuple of values into a string."""
							πF.SetLineno(162)
							// line 163: exctype, value, tb = err
							πF.SetLineno(163)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}}}, µerr); πE != nil {
								continue
							}
							µexctype = πTemp001
							µvalue = πTemp002
							µtb = πTemp003
							// line 165: while tb and self._is_relevant_tb_level(tb):
							πF.SetLineno(165)
							πF.PushCheckpoint(2)
							πTemp004 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
								continue
							}
							πTemp001 = µtb
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label4
							}
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
								continue
							}
							πTemp007[0] = µtb
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_is_relevant_tb_level, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp001 = πTemp003
						Label4:
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 166: tb = tb.tb_next
							πF.SetLineno(166)
							if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtb, ßtb_next, nil); πE != nil {
								continue
							}
							µtb = πTemp001
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µexctype, "exctype"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtest, ßfailureException, nil); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µexctype == πTemp002).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							goto Label6
							// line 168: if exctype is test.failureException:
							πF.SetLineno(168)
						Label5:
							// line 170: length = self._count_relevant_tb_levels(tb)
							πF.SetLineno(170)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
								continue
							}
							πTemp007[0] = µtb
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_count_relevant_tb_levels, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µlength = πTemp002
							// line 171: msgLines = traceback.format_exception(exctype, value, tb, length)
							πF.SetLineno(171)
							πTemp007 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µexctype, "exctype"); πE != nil {
								continue
							}
							πTemp007[0] = µexctype
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp007[1] = µvalue
							if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
								continue
							}
							πTemp007[2] = µtb
							if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
								continue
							}
							πTemp007[3] = µlength
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtraceback); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßformat_exception, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µmsgLines = πTemp001
							goto Label7
						Label6:
							// line 173: msgLines = traceback.format_exception(exctype, value, tb)
							πF.SetLineno(173)
							πTemp007 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µexctype, "exctype"); πE != nil {
								continue
							}
							πTemp007[0] = µexctype
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp007[1] = µvalue
							if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
								continue
							}
							πTemp007[2] = µtb
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtraceback); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßformat_exception, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µmsgLines = πTemp001
							goto Label7
						Label7:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbuffer, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							goto Label9
							// line 175: if self.buffer:
							πF.SetLineno(175)
						Label8:
							// line 176: output = sys.stdout.getvalue()
							πF.SetLineno(176)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstdout, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßgetvalue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µoutput = πTemp002
							// line 177: error = sys.stderr.getvalue()
							πF.SetLineno(177)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstderr, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßgetvalue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µerror = πTemp002
							if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µoutput); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label10
							}
							goto Label11
							// line 178: if output:
							πF.SetLineno(178)
						Label10:
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = πg.NewStr("\n").ToObject()
							if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µoutput, ßendswith, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label12
							}
							goto Label13
							// line 179: if not output.endswith('\n'):
							πF.SetLineno(179)
						Label12:
							// line 180: output += '\n'
							πF.SetLineno(180)
							if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µoutput, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							µoutput = πTemp001
							goto Label13
						Label13:
							// line 181: msgLines.append(STDOUT_LINE % output)
							πF.SetLineno(181)
							πTemp007 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSTDOUT_LINE); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πTemp002, µoutput); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πE = πg.CheckLocal(πF, µmsgLines, "msgLines"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmsgLines, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							goto Label11
						Label11:
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µerror); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label14
							}
							goto Label15
							// line 182: if error:
							πF.SetLineno(182)
						Label14:
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = πg.NewStr("\n").ToObject()
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µerror, ßendswith, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label16
							}
							goto Label17
							// line 183: if not error.endswith('\n'):
							πF.SetLineno(183)
						Label16:
							// line 184: error += '\n'
							πF.SetLineno(184)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µerror, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							µerror = πTemp001
							goto Label17
						Label17:
							// line 185: msgLines.append(STDERR_LINE % error)
							πF.SetLineno(185)
							πTemp007 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSTDERR_LINE); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πTemp002, µerror); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πE = πg.CheckLocal(πF, µmsgLines, "msgLines"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmsgLines, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							goto Label15
						Label15:
							goto Label9
						Label9:
							// line 186: return ''.join(msgLines)
							πF.SetLineno(186)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsgLines, "msgLines"); πE != nil {
								continue
							}
							πTemp007[0] = µmsgLines
							if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πR = πTemp002
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_exc_info_to_string.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 189: def _is_relevant_tb_level(self, tb):
					πF.SetLineno(189)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "tb", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("_is_relevant_tb_level", "build/src/__python__/unittest_result.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtb *πg.Object = πArgs[1]; _ = µtb
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 190: return '__unittest' in tb.tb_frame.f_globals
							πF.SetLineno(190)
							if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtb, ßtb_frame, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßf_globals, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp003, ß__unittest.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004).ToObject()
							πR = πTemp001
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_is_relevant_tb_level.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 192: def _count_relevant_tb_levels(self, tb):
					πF.SetLineno(192)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "tb", Def: nil}
					πTemp021 = πg.NewFunction(πg.NewCode("_count_relevant_tb_levels", "build/src/__python__/unittest_result.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtb *πg.Object = πArgs[1]; _ = µtb
						var µlength *πg.Object = πg.UnboundLocal; _ = µlength
						var πTemp001 bool
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 193: length = 0
							πF.SetLineno(193)
							µlength = πg.NewInt(0).ToObject()
							// line 194: while tb and not self._is_relevant_tb_level(tb):
							πF.SetLineno(194)
							πF.PushCheckpoint(2)
							πTemp001 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp001 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
								continue
							}
							πTemp003 = µtb
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label4
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
								continue
							}
							πTemp006[0] = µtb
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ß_is_relevant_tb_level, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp009, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(!πTemp009).ToObject()
							πTemp003 = πTemp005
						Label4:
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πE != nil || !πTemp002 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 195: length += 1
							πF.SetLineno(195)
							if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, µlength, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µlength = πTemp003
							// line 196: tb = tb.tb_next
							πF.SetLineno(196)
							if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µtb, ßtb_next, nil); πE != nil {
								continue
							}
							µtb = πTemp003
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 197: return length
							πF.SetLineno(197)
							if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
								continue
							}
							πR = µlength
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_count_relevant_tb_levels.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 199: def __repr__(self):
					πF.SetLineno(199)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("__repr__", "build/src/__python__/unittest_result.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 200: return ("<%s run=%i errors=%i failures=%i>" %
							πF.SetLineno(200)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßutil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßstrclass, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßtestsRun, nil); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßerrors, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp006
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßfailures, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp006
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002 = πg.NewTuple4(πTemp004, πTemp005, πTemp007, πTemp008).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<%s run=%i errors=%i failures=%i>").ToObject(), πTemp002); πE != nil {
								continue
							}
							πR = πTemp001
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp022); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("TestResult").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTestResult.ToObject(), πTemp007); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("unittest_result", Code)
}
