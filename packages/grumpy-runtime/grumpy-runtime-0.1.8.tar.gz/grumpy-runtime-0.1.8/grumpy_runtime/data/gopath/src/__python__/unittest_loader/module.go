package unittest_loader
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/unittest_loader.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßException := πg.InternStr("Exception")
		ßFalse := πg.InternStr("False")
		ßIGNORECASE := πg.InternStr("IGNORECASE")
		ßImportError := πg.InternStr("ImportError")
		ßLoadTestsFailure := πg.InternStr("LoadTestsFailure")
		ßModuleImportFailure := πg.InternStr("ModuleImportFailure")
		ßModuleType := πg.InternStr("ModuleType")
		ßNone := πg.InternStr("None")
		ßTestCase := πg.InternStr("TestCase")
		ßTestLoader := πg.InternStr("TestLoader")
		ßTestSuite := πg.InternStr("TestSuite")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßUnboundMethodType := πg.InternStr("UnboundMethodType")
		ßVALID_MODULE_NAME := πg.InternStr("VALID_MODULE_NAME")
		ß_CmpToKey := πg.InternStr("_CmpToKey")
		ß__call__ := πg.InternStr("__call__")
		ß__file__ := πg.InternStr("__file__")
		ß__import__ := πg.InternStr("__import__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__unittest := πg.InternStr("__unittest")
		ß_find_tests := πg.InternStr("_find_tests")
		ß_fnmatch := πg.InternStr("_fnmatch")
		ß_get_directory_containing_module := πg.InternStr("_get_directory_containing_module")
		ß_get_module_from_name := πg.InternStr("_get_module_from_name")
		ß_get_name_from_path := πg.InternStr("_get_name_from_path")
		ß_makeLoader := πg.InternStr("_makeLoader")
		ß_make_failed_import_test := πg.InternStr("_make_failed_import_test")
		ß_make_failed_load_tests := πg.InternStr("_make_failed_load_tests")
		ß_make_failed_test := πg.InternStr("_make_failed_test")
		ß_match_path := πg.InternStr("_match_path")
		ß_top_level_dir := πg.InternStr("_top_level_dir")
		ßabspath := πg.InternStr("abspath")
		ßappend := πg.InternStr("append")
		ßbasename := πg.InternStr("basename")
		ßcase := πg.InternStr("case")
		ßcmp := πg.InternStr("cmp")
		ßcmp_to_key := πg.InternStr("cmp_to_key")
		ßcompile := πg.InternStr("compile")
		ßdefaultTestLoader := πg.InternStr("defaultTestLoader")
		ßdir := πg.InternStr("dir")
		ßdirname := πg.InternStr("dirname")
		ßdiscover := πg.InternStr("discover")
		ßfindTestCases := πg.InternStr("findTestCases")
		ßfnmatch := πg.InternStr("fnmatch")
		ßformat_exc := πg.InternStr("format_exc")
		ßfunctools := πg.InternStr("functools")
		ßgetTestCaseNames := πg.InternStr("getTestCaseNames")
		ßgetattr := πg.InternStr("getattr")
		ßhasattr := πg.InternStr("hasattr")
		ßinsert := πg.InternStr("insert")
		ßisabs := πg.InternStr("isabs")
		ßisdir := πg.InternStr("isdir")
		ßisfile := πg.InternStr("isfile")
		ßisinstance := πg.InternStr("isinstance")
		ßissubclass := πg.InternStr("issubclass")
		ßjoin := πg.InternStr("join")
		ßlist := πg.InternStr("list")
		ßlistdir := πg.InternStr("listdir")
		ßloadTestsFromModule := πg.InternStr("loadTestsFromModule")
		ßloadTestsFromName := πg.InternStr("loadTestsFromName")
		ßloadTestsFromNames := πg.InternStr("loadTestsFromNames")
		ßloadTestsFromTestCase := πg.InternStr("loadTestsFromTestCase")
		ßload_tests := πg.InternStr("load_tests")
		ßlower := πg.InternStr("lower")
		ßmakeSuite := πg.InternStr("makeSuite")
		ßmap := πg.InternStr("map")
		ßmatch := πg.InternStr("match")
		ßmodules := πg.InternStr("modules")
		ßnormpath := πg.InternStr("normpath")
		ßobject := πg.InternStr("object")
		ßos := πg.InternStr("os")
		ßpath := πg.InternStr("path")
		ßre := πg.InternStr("re")
		ßrealpath := πg.InternStr("realpath")
		ßrelpath := πg.InternStr("relpath")
		ßremove := πg.InternStr("remove")
		ßreplace := πg.InternStr("replace")
		ßrunTest := πg.InternStr("runTest")
		ßsep := πg.InternStr("sep")
		ßsort := πg.InternStr("sort")
		ßsortTestMethodsUsing := πg.InternStr("sortTestMethodsUsing")
		ßsplit := πg.InternStr("split")
		ßsplitext := πg.InternStr("splitext")
		ßstartswith := πg.InternStr("startswith")
		ßsuite := πg.InternStr("suite")
		ßsuiteClass := πg.InternStr("suiteClass")
		ßsys := πg.InternStr("sys")
		ßtest := πg.InternStr("test")
		ßtestMethodPrefix := πg.InternStr("testMethodPrefix")
		ßtraceback := πg.InternStr("traceback")
		ßtype := πg.InternStr("type")
		ßtypes := πg.InternStr("types")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 []πg.Param
		_ = πTemp004
		var πTemp005 *πg.Object
		_ = πTemp005
		var πTemp006 *πg.Dict
		_ = πTemp006
		var πTemp007 *πg.Object
		_ = πTemp007
		var πTemp008 *πg.Object
		_ = πTemp008
		var πTemp009 *πg.Object
		_ = πTemp009
		var πTemp010 *πg.Object
		_ = πTemp010
		var πTemp011 *πg.Object
		_ = πTemp011
		var πTemp012 *πg.Object
		_ = πTemp012
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """Loading unittests."""
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
			// line 4: import re
			πF.SetLineno(4)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 5: import sys
			πF.SetLineno(5)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 6: import traceback
			πF.SetLineno(6)
			if πTemp002, πE = πg.ImportModule(πF, "traceback"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtraceback.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 7: import types
			πF.SetLineno(7)
			if πTemp002, πE = πg.ImportModule(πF, "types"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtypes.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 11: import functools
			πF.SetLineno(11)
			if πTemp002, πE = πg.ImportModule(πF, "functools"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßfunctools.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 12: import fnmatch as _fnmatch
			πF.SetLineno(12)
			if πTemp002, πE = πg.ImportModule(πF, "fnmatch"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_fnmatch.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 13: _CmpToKey = functools.cmp_to_key
			πF.SetLineno(13)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßfunctools); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßcmp_to_key, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_CmpToKey.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 14: fnmatch = _fnmatch.fnmatch
			πF.SetLineno(14)
			if πTemp001, πE = πg.ResolveGlobal(πF, ß_fnmatch); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßfnmatch, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßfnmatch.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 17: import unittest_case as case
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "unittest_case"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßcase.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: import unittest_suite as suite
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "unittest_suite"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsuite.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 20: __unittest = True
			πF.SetLineno(20)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__unittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 25: VALID_MODULE_NAME = re.compile(r'[_a-z]\w*\.py$', re.IGNORECASE)
			πF.SetLineno(25)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("[_a-z]\\w*\\.py$").ToObject()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßIGNORECASE, nil); πE != nil {
				continue
			}
			πTemp002[1] = πTemp003
			if πTemp001, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßcompile, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßVALID_MODULE_NAME.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 28: def _make_failed_import_test(name, suiteClass):
			πF.SetLineno(28)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "name", Def: nil}
			πTemp004[1] = πg.Param{Name: "suiteClass", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("_make_failed_import_test", "build/src/__python__/unittest_loader.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]; _ = µname
				var µsuiteClass *πg.Object = πArgs[1]; _ = µsuiteClass
				var µmessage *πg.Object = πg.UnboundLocal; _ = µmessage
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 29: message = 'Failed to import test module: %s\n%s' % (name, traceback.format_exc())
					πF.SetLineno(29)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtraceback); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßformat_exc, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µname, πTemp003).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("Failed to import test module: %s\n%s").ToObject(), πTemp002); πE != nil {
						continue
					}
					µmessage = πTemp001
					// line 30: return _make_failed_test('ModuleImportFailure', name, ImportError(message),
					πF.SetLineno(30)
					πTemp005 = πF.MakeArgs(4)
					πTemp005[0] = ßModuleImportFailure.ToObject()
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp005[1] = µname
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
						continue
					}
					πTemp006[0] = µmessage
					if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp005[2] = πTemp002
					if πE = πg.CheckLocal(πF, µsuiteClass, "suiteClass"); πE != nil {
						continue
					}
					πTemp005[3] = µsuiteClass
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_make_failed_test); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
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
			if πE = πF.Globals().SetItem(πF, ß_make_failed_import_test.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 33: def _make_failed_load_tests(name, exception, suiteClass):
			πF.SetLineno(33)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "name", Def: nil}
			πTemp004[1] = πg.Param{Name: "exception", Def: nil}
			πTemp004[2] = πg.Param{Name: "suiteClass", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("_make_failed_load_tests", "build/src/__python__/unittest_loader.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]; _ = µname
				var µexception *πg.Object = πArgs[1]; _ = µexception
				var µsuiteClass *πg.Object = πArgs[2]; _ = µsuiteClass
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
					// line 34: return _make_failed_test('LoadTestsFailure', name, exception, suiteClass)
					πF.SetLineno(34)
					πTemp001 = πF.MakeArgs(4)
					πTemp001[0] = ßLoadTestsFailure.ToObject()
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001[1] = µname
					if πE = πg.CheckLocal(πF, µexception, "exception"); πE != nil {
						continue
					}
					πTemp001[2] = µexception
					if πE = πg.CheckLocal(πF, µsuiteClass, "suiteClass"); πE != nil {
						continue
					}
					πTemp001[3] = µsuiteClass
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_make_failed_test); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_make_failed_load_tests.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 36: def _make_failed_test(classname, methodname, exception, suiteClass):
			πF.SetLineno(36)
			πTemp004 = make([]πg.Param, 4)
			πTemp004[0] = πg.Param{Name: "classname", Def: nil}
			πTemp004[1] = πg.Param{Name: "methodname", Def: nil}
			πTemp004[2] = πg.Param{Name: "exception", Def: nil}
			πTemp004[3] = πg.Param{Name: "suiteClass", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("_make_failed_test", "build/src/__python__/unittest_loader.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µclassname *πg.Object = πArgs[0]; _ = µclassname
				var µmethodname *πg.Object = πArgs[1]; _ = µmethodname
				var µexception *πg.Object = πArgs[2]; _ = µexception
				var µsuiteClass *πg.Object = πArgs[3]; _ = µsuiteClass
				var µtestFailure *πg.Object = πg.UnboundLocal; _ = µtestFailure
				var µattrs *πg.Object = πg.UnboundLocal; _ = µattrs
				var µTestClass *πg.Object = πg.UnboundLocal; _ = µTestClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Dict
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 37: def testFailure(self):
					πF.SetLineno(37)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("testFailure", "build/src/__python__/unittest_loader.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µexception, "exception"); πE != nil {
								continue
							}
							// line 38: raise exception
							πF.SetLineno(38)
							πE = πF.Raise(µexception, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					µtestFailure = πTemp001
					// line 39: attrs = {methodname: testFailure}
					πF.SetLineno(39)
					πTemp003 = πg.NewDict()
					if πE = πg.CheckLocal(πF, µmethodname, "methodname"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtestFailure, "testFailure"); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, µmethodname, µtestFailure); πE != nil {
						continue
					}
					πTemp004 = πTemp003.ToObject()
					µattrs = πTemp004
					// line 40: TestClass = type(classname, (case.TestCase,), attrs)
					πF.SetLineno(40)
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µclassname, "classname"); πE != nil {
						continue
					}
					πTemp005[0] = µclassname
					if πTemp006, πE = πg.ResolveGlobal(πF, ßcase); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßTestCase, nil); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple1(πTemp007).ToObject()
					πTemp005[1] = πTemp004
					if πE = πg.CheckLocal(πF, µattrs, "attrs"); πE != nil {
						continue
					}
					πTemp005[2] = µattrs
					if πTemp004, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µTestClass = πTemp006
					// line 41: return suiteClass((TestClass(methodname),))
					πF.SetLineno(41)
					πTemp005 = πF.MakeArgs(1)
					πTemp008 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmethodname, "methodname"); πE != nil {
						continue
					}
					πTemp008[0] = µmethodname
					if πE = πg.CheckLocal(πF, µTestClass, "TestClass"); πE != nil {
						continue
					}
					if πTemp006, πE = µTestClass.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					πTemp004 = πg.NewTuple1(πTemp006).ToObject()
					πTemp005[0] = πTemp004
					if πE = πg.CheckLocal(πF, µsuiteClass, "suiteClass"); πE != nil {
						continue
					}
					if πTemp004, πE = µsuiteClass.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πR = πTemp004
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_make_failed_test.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 44: class TestLoader(object):
			πF.SetLineno(44)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp007, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp007); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TestLoader", "build/src/__python__/unittest_loader.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []πg.Param
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
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 *πg.Object
				_ = πTemp012
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 45: """
					πF.SetLineno(45)
					// line 49: testMethodPrefix = 'test'
					πF.SetLineno(49)
					if πE = πClass.SetItem(πF, ßtestMethodPrefix.ToObject(), ßtest.ToObject()); πE != nil {
						continue
					}
					// line 50: sortTestMethodsUsing = cmp
					πF.SetLineno(50)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßcmp); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßsortTestMethodsUsing.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 51: suiteClass = suite.TestSuite
					πF.SetLineno(51)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßsuite); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßTestSuite, nil); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßsuiteClass.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 52: _top_level_dir = None
					πF.SetLineno(52)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß_top_level_dir.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 54: def loadTestsFromTestCase(self, testCaseClass):
					πF.SetLineno(54)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "testCaseClass", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("loadTestsFromTestCase", "build/src/__python__/unittest_loader.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtestCaseClass *πg.Object = πArgs[1]; _ = µtestCaseClass
						var µtestCaseNames *πg.Object = πg.UnboundLocal; _ = µtestCaseNames
						var µloaded_suite *πg.Object = πg.UnboundLocal; _ = µloaded_suite
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 55: """Return a suite of all tests cases contained in testCaseClass"""
							πF.SetLineno(55)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtestCaseClass, "testCaseClass"); πE != nil {
								continue
							}
							πTemp001[0] = µtestCaseClass
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsuite); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßTestSuite, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
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
							// line 56: if issubclass(testCaseClass, suite.TestSuite):
							πF.SetLineno(56)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("Test cases should not be derived from TestSuite. Maybe you meant to derive from TestCase?").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 57: raise TypeError("Test cases should not be derived from TestSuite." \
							πF.SetLineno(57)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							// line 59: testCaseNames = self.getTestCaseNames(testCaseClass)
							πF.SetLineno(59)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtestCaseClass, "testCaseClass"); πE != nil {
								continue
							}
							πTemp001[0] = µtestCaseClass
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgetTestCaseNames, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtestCaseNames = πTemp003
							if πE = πg.CheckLocal(πF, µtestCaseNames, "testCaseNames"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µtestCaseNames); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp005).ToObject()
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label3
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtestCaseClass, "testCaseClass"); πE != nil {
								continue
							}
							πTemp001[0] = µtestCaseClass
							πTemp001[1] = ßrunTest.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πTemp006
						Label3:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 60: if not testCaseNames and hasattr(testCaseClass, 'runTest'):
							πF.SetLineno(60)
						Label4:
							// line 61: testCaseNames = ['runTest']
							πF.SetLineno(61)
							πTemp001 = make([]*πg.Object, 1)
							πTemp001[0] = ßrunTest.ToObject()
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µtestCaseNames = πTemp002
							goto Label5
						Label5:
							// line 62: loaded_suite = self.suiteClass(map(testCaseClass, testCaseNames))
							πF.SetLineno(62)
							πTemp001 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtestCaseClass, "testCaseClass"); πE != nil {
								continue
							}
							πTemp007[0] = µtestCaseClass
							if πE = πg.CheckLocal(πF, µtestCaseNames, "testCaseNames"); πE != nil {
								continue
							}
							πTemp007[1] = µtestCaseNames
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsuiteClass, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µloaded_suite = πTemp003
							// line 63: return loaded_suite
							πF.SetLineno(63)
							if πE = πg.CheckLocal(πF, µloaded_suite, "loaded_suite"); πE != nil {
								continue
							}
							πR = µloaded_suite
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßloadTestsFromTestCase.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 65: def loadTestsFromModule(self, module, use_load_tests=True):
					πF.SetLineno(65)
					πTemp003 = make([]πg.Param, 3)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "module", Def: nil}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp003[2] = πg.Param{Name: "use_load_tests", Def: πTemp004}
					πTemp002 = πg.NewFunction(πg.NewCode("loadTestsFromModule", "build/src/__python__/unittest_loader.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmodule *πg.Object = πArgs[1]; _ = µmodule
						var µuse_load_tests *πg.Object = πArgs[2]; _ = µuse_load_tests
						var µtests *πg.Object = πg.UnboundLocal; _ = µtests
						var µname *πg.Object = πg.UnboundLocal; _ = µname
						var µobj *πg.Object = πg.UnboundLocal; _ = µobj
						var µload_tests *πg.Object = πg.UnboundLocal; _ = µload_tests
						var µe *πg.Object = πg.UnboundLocal; _ = µe
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 *πg.BaseException
						_ = πTemp009
						var πTemp010 *πg.Traceback
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							case 11: goto Label11
							default: panic("unexpected function state")
							}
							// line 66: """Return a suite of all tests cases contained in the given module"""
							πF.SetLineno(66)
							// line 67: tests = []
							πF.SetLineno(67)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µtests = πTemp002
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							πTemp001[0] = µmodule
							if πTemp003, πE = πg.ResolveGlobal(πF, ßdir); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp005 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp006 = !isStop
							} else {
								πTemp006 = true
								µname = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 69: obj = getattr(module, name)
							πF.SetLineno(69)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							πTemp001[0] = µmodule
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001[1] = µname
							if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µobj = πTemp004
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πTemp004, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp003 = πTemp007
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label4
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πTemp004, πE = πg.ResolveGlobal(πF, ßcase); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp004, ßTestCase, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp007
							if πTemp004, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp003 = πTemp007
						Label4:
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label5
							}
							goto Label6
							// line 70: if isinstance(obj, type) and issubclass(obj, case.TestCase):
							πF.SetLineno(70)
						Label5:
							// line 71: tests.append(self.loadTestsFromTestCase(obj))
							πF.SetLineno(71)
							πTemp001 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp008[0] = µobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßloadTestsFromTestCase, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µtests, "tests"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µtests, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label6
						Label6:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 73: load_tests = getattr(module, 'load_tests', None)
							πF.SetLineno(73)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							πTemp001[0] = µmodule
							πTemp001[1] = ßload_tests.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
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
							µload_tests = πTemp003
							// line 74: tests = self.suiteClass(tests)
							πF.SetLineno(74)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtests, "tests"); πE != nil {
								continue
							}
							πTemp001[0] = µtests
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsuiteClass, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtests = πTemp003
							if πE = πg.CheckLocal(πF, µuse_load_tests, "use_load_tests"); πE != nil {
								continue
							}
							πTemp002 = µuse_load_tests
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µload_tests, "load_tests"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µload_tests != πTemp004).ToObject()
							πTemp002 = πTemp003
						Label7:
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label8
							}
							goto Label9
							// line 75: if use_load_tests and load_tests is not None:
							πF.SetLineno(75)
						Label8:
							// line 76: try:
							πF.SetLineno(76)
							πF.PushCheckpoint(11)
							// line 77: return load_tests(self, tests, None)
							πF.SetLineno(77)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µtests, "tests"); πE != nil {
								continue
							}
							πTemp001[1] = µtests
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πE = πg.CheckLocal(πF, µload_tests, "load_tests"); πE != nil {
								continue
							}
							if πTemp002, πE = µload_tests.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp002
							continue
							πF.PopCheckpoint()
							goto Label10
						Label11:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label12
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 78: except Exception, e:
							πF.SetLineno(78)
						Label12:
							µe = πTemp009.ToObject()
							// line 79: return _make_failed_load_tests(module.__name__, e,
							πF.SetLineno(79)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µmodule, ß__name__, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							πTemp001[1] = µe
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsuiteClass, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_make_failed_load_tests); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp003
							continue
							πF.RestoreExc(nil, nil)
							goto Label10
						Label10:
							goto Label9
						Label9:
							// line 81: return tests
							πF.SetLineno(81)
							if πE = πg.CheckLocal(πF, µtests, "tests"); πE != nil {
								continue
							}
							πR = µtests
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßloadTestsFromModule.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 83: def loadTestsFromName(self, name, module=None):
					πF.SetLineno(83)
					πTemp003 = make([]πg.Param, 3)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "name", Def: nil}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp003[2] = πg.Param{Name: "module", Def: πTemp005}
					πTemp004 = πg.NewFunction(πg.NewCode("loadTestsFromName", "build/src/__python__/unittest_loader.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
						var µmodule *πg.Object = πArgs[2]; _ = µmodule
						var µparts *πg.Object = πg.UnboundLocal; _ = µparts
						var µparts_copy *πg.Object = πg.UnboundLocal; _ = µparts_copy
						var µobj *πg.Object = πg.UnboundLocal; _ = µobj
						var µpart *πg.Object = πg.UnboundLocal; _ = µpart
						var µparent *πg.Object = πg.UnboundLocal; _ = µparent
						var µinst *πg.Object = πg.UnboundLocal; _ = µinst
						var µtest *πg.Object = πg.UnboundLocal; _ = µtest
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 11: goto Label11
							case 12: goto Label12
							case 3: goto Label3
							case 4: goto Label4
							case 7: goto Label7
							default: panic("unexpected function state")
							}
							// line 84: """Return a suite of all tests cases given a string specifier.
							πF.SetLineno(84)
							// line 92: parts = name.split('.')
							πF.SetLineno(92)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr(".").ToObject()
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µname, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µparts = πTemp003
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µmodule == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 93: if module is None:
							πF.SetLineno(93)
						Label1:
							// line 94: parts_copy = parts[:]
							πF.SetLineno(94)
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µparts, πTemp002); πE != nil {
								continue
							}
							µparts_copy = πTemp003
							// line 95: while parts_copy:
							πF.SetLineno(95)
							πF.PushCheckpoint(4)
							πTemp004 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label5
							}
							if πE = πg.CheckLocal(πF, µparts_copy, "parts_copy"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µparts_copy); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 96: try:
							πF.SetLineno(96)
							πF.PushCheckpoint(7)
							// line 97: module = __import__('.'.join(parts_copy))
							πF.SetLineno(97)
							πTemp001 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µparts_copy, "parts_copy"); πE != nil {
								continue
							}
							πTemp006[0] = µparts_copy
							if πTemp002, πE = πg.GetAttr(πF, πg.NewStr(".").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ß__import__); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmodule = πTemp003
							// line 98: break
							πF.SetLineno(98)
							πTemp004 = true
							continue
							πF.PopCheckpoint()
							goto Label6
						Label7:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label8
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 99: except ImportError:
							πF.SetLineno(99)
						Label8:
							// line 100: del parts_copy[-1]
							πF.SetLineno(100)
							if πE = πg.CheckLocal(πF, µparts_copy, "parts_copy"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.DelItem(πF, µparts_copy, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µparts_copy, "parts_copy"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µparts_copy); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label9
							}
							goto Label10
							// line 101: if not parts_copy:
							πF.SetLineno(101)
						Label9:
							// line 102: raise
							πF.SetLineno(102)
							πE = πF.Raise(nil, nil, nil)
							continue
							goto Label10
						Label10:
							πF.RestoreExc(nil, nil)
							goto Label6
						Label6:
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 103: parts = parts[1:]
							πF.SetLineno(103)
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µparts, πTemp002); πE != nil {
								continue
							}
							µparts = πTemp003
							goto Label2
						Label2:
							// line 104: obj = module
							πF.SetLineno(104)
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							µobj = µmodule
							if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µparts); πE != nil {
								continue
							}
							πF.PushCheckpoint(12)
							πTemp004 = false
						Label11:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label13
							}
							if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp005 = !isStop
							} else {
								πTemp005 = true
								µpart = πTemp003
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(11)            
							// line 106: parent, obj = obj, getattr(obj, part)
							πF.SetLineno(106)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πE = πg.CheckLocal(πF, µpart, "part"); πE != nil {
								continue
							}
							πTemp001[1] = µpart
							if πTemp009, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp003 = πg.NewTuple2(µobj, πTemp010).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp009}, πg.TieTarget{Target: &πTemp010}}}, πTemp003); πE != nil {
								continue
							}
							µparent = πTemp009
							µobj = πTemp010
							continue
						Label12:
							if πE != nil || πR != nil {
								continue
							}
						Label13:
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßModuleType, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
								goto Label14
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πTemp009
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label15
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πTemp003, πE = πg.ResolveGlobal(πF, ßcase); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp003, ßTestCase, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp009
							if πTemp003, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πTemp009
						Label15:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label16
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp003, ßUnboundMethodType, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp009
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πTemp009
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label17
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							πTemp001[0] = µparent
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πTemp009
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label17
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							πTemp001[0] = µparent
							if πTemp003, πE = πg.ResolveGlobal(πF, ßcase); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp003, ßTestCase, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp009
							if πTemp003, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πTemp009
						Label17:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label18
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsuite); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßTestSuite, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
								goto Label19
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							πTemp001[1] = ß__call__.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
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
								goto Label20
							}
							goto Label21
							// line 108: if isinstance(obj, types.ModuleType):
							πF.SetLineno(108)
						Label14:
							// line 109: return self.loadTestsFromModule(obj)
							πF.SetLineno(109)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßloadTestsFromModule, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp003
							continue
							goto Label22
							// line 110: elif isinstance(obj, type) and issubclass(obj, case.TestCase):
							πF.SetLineno(110)
						Label16:
							// line 111: return self.loadTestsFromTestCase(obj)
							πF.SetLineno(111)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßloadTestsFromTestCase, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp003
							continue
							goto Label22
							// line 112: elif (isinstance(obj, types.UnboundMethodType) and
							πF.SetLineno(112)
						Label18:
							// line 115: name = parts[-1]
							πF.SetLineno(115)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µparts, πTemp002); πE != nil {
								continue
							}
							µname = πTemp003
							// line 116: inst = parent(name)
							πF.SetLineno(116)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001[0] = µname
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							if πTemp002, πE = µparent.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µinst = πTemp002
							// line 117: return self.suiteClass([inst])
							πF.SetLineno(117)
							πTemp001 = πF.MakeArgs(1)
							πTemp006 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µinst, "inst"); πE != nil {
								continue
							}
							πTemp006[0] = µinst
							πTemp002 = πg.NewList(πTemp006...).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsuiteClass, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp003
							continue
							goto Label22
							// line 118: elif isinstance(obj, suite.TestSuite):
							πF.SetLineno(118)
						Label19:
							// line 119: return obj
							πF.SetLineno(119)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πR = µobj
							continue
							goto Label22
							// line 120: elif hasattr(obj, '__call__'):
							πF.SetLineno(120)
						Label20:
							// line 121: test = obj()
							πF.SetLineno(121)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp002, πE = µobj.Call(πF, nil, nil); πE != nil {
								continue
							}
							µtest = πTemp002
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp001[0] = µtest
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsuite); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßTestSuite, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
								goto Label23
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp001[0] = µtest
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcase); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßTestCase, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
								goto Label24
							}
							goto Label25
							// line 122: if isinstance(test, suite.TestSuite):
							πF.SetLineno(122)
						Label23:
							// line 123: return test
							πF.SetLineno(123)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πR = µtest
							continue
							goto Label26
							// line 124: elif isinstance(test, case.TestCase):
							πF.SetLineno(124)
						Label24:
							// line 125: return self.suiteClass([test])
							πF.SetLineno(125)
							πTemp001 = πF.MakeArgs(1)
							πTemp006 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp006[0] = µtest
							πTemp002 = πg.NewList(πTemp006...).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsuiteClass, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp003
							continue
							goto Label26
						Label25:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(µobj, µtest).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("calling %s returned %s, not a test").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 127: raise TypeError("calling %s returned %s, not a test" %
							πF.SetLineno(127)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label26
						Label26:
							goto Label22
						Label21:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("don't know how to make test from: %s").ToObject(), µobj); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 130: raise TypeError("don't know how to make test from: %s" % obj)
							πF.SetLineno(130)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label22
						Label22:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßloadTestsFromName.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 132: def loadTestsFromNames(self, names, module=None):
					πF.SetLineno(132)
					πTemp003 = make([]πg.Param, 3)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "names", Def: nil}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp003[2] = πg.Param{Name: "module", Def: πTemp006}
					πTemp005 = πg.NewFunction(πg.NewCode("loadTestsFromNames", "build/src/__python__/unittest_loader.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µnames *πg.Object = πArgs[1]; _ = µnames
						var µmodule *πg.Object = πArgs[2]; _ = µmodule
						var µsuites *πg.Object = πg.UnboundLocal; _ = µsuites
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 133: """Return a suite of all tests cases found using the given sequence
							πF.SetLineno(133)
							// line 136: suites = [self.loadTestsFromName(name, module) for name in names]
							πF.SetLineno(136)
							πTemp003 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/unittest_loader.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µname *πg.Object = πg.UnboundLocal; _ = µname
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 bool
								_ = πTemp002
								var πTemp003 bool
								_ = πTemp003
								var πTemp004 *πg.Object
								_ = πTemp004
								var πTemp005 []*πg.Object
								_ = πTemp005
								var πTemp006 *πg.Object
								_ = πTemp006
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
									for ; πF.State() >= 0; πF.PopCheckpoint() {
										switch πF.State() {
										case 0:
										case 1: goto Label1
										case 2: goto Label2
										case 4: goto Label4
										default: panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µnames); πE != nil {
											continue
										}
										πF.PushCheckpoint(2)
										πTemp002 = false
									Label1:
										if πE != nil || πR != nil {
											continue
										}
										if πTemp002 {
											πF.PopCheckpoint()
											goto Label3
										}
										if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
											isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
											if exc != nil {
												πE = exc
											} else if isStop {
												πE = nil
												πF.RestoreExc(nil, nil)
											}
											πTemp003 = !isStop
										} else {
											πTemp003 = true
											µname = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 136: suites = [self.loadTestsFromName(name, module) for name in names]
										πF.SetLineno(136)
										πTemp005 = πF.MakeArgs(2)
										if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
											continue
										}
										πTemp005[0] = µname
										if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
											continue
										}
										πTemp005[1] = µmodule
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp004, πE = πg.GetAttr(πF, µself, ßloadTestsFromName, nil); πE != nil {
											continue
										}
										if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp005)
										πF.PushCheckpoint(4)
										return πTemp006, nil
									Label4:
										πTemp004 = πSent
										continue
									Label2:
										if πE != nil || πR != nil {
											continue
										}
									Label3:
									}
									return nil, πE
								}).ToObject(), nil
							}), πF.Globals()).ToObject()
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
								continue
							}
							µsuites = πTemp001
							// line 137: return self.suiteClass(suites)
							πF.SetLineno(137)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsuites, "suites"); πE != nil {
								continue
							}
							πTemp005[0] = µsuites
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsuiteClass, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πR = πTemp004
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßloadTestsFromNames.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 139: def getTestCaseNames(self, testCaseClass):
					πF.SetLineno(139)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "testCaseClass", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("getTestCaseNames", "build/src/__python__/unittest_loader.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtestCaseClass *πg.Object = πArgs[1]; _ = µtestCaseClass
						var µisTestMethod *πg.Object = πg.UnboundLocal; _ = µisTestMethod
						var µtestFnNames *πg.Object = πg.UnboundLocal; _ = µtestFnNames
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 πg.KWArgs
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 140: """Return a sorted sequence of method names found within testCaseClass
							πF.SetLineno(140)
							// line 142: def isTestMethod(attrname, testCaseClass=testCaseClass,
							πF.SetLineno(142)
							πTemp002 = make([]πg.Param, 3)
							πTemp002[0] = πg.Param{Name: "attrname", Def: nil}
							if πE = πg.CheckLocal(πF, µtestCaseClass, "testCaseClass"); πE != nil {
								continue
							}
							πTemp002[1] = πg.Param{Name: "testCaseClass", Def: µtestCaseClass}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtestMethodPrefix, nil); πE != nil {
								continue
							}
							πTemp002[2] = πg.Param{Name: "prefix", Def: πTemp003}
							πTemp001 = πg.NewFunction(πg.NewCode("isTestMethod", "build/src/__python__/unittest_loader.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µattrname *πg.Object = πArgs[0]; _ = µattrname
								var µtestCaseClass *πg.Object = πArgs[1]; _ = µtestCaseClass
								var µprefix *πg.Object = πArgs[2]; _ = µprefix
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 bool
								_ = πTemp002
								var πTemp003 []*πg.Object
								_ = πTemp003
								var πTemp004 *πg.Object
								_ = πTemp004
								var πTemp005 *πg.Object
								_ = πTemp005
								var πTemp006 []*πg.Object
								_ = πTemp006
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 144: return attrname.startswith(prefix) and \
									πF.SetLineno(144)
									πTemp003 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
										continue
									}
									πTemp003[0] = µprefix
									if πE = πg.CheckLocal(πF, µattrname, "attrname"); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetAttr(πF, µattrname, ßstartswith, nil); πE != nil {
										continue
									}
									if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									πTemp001 = πTemp005
									if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if !πTemp002 {
										goto Label1
									}
									πTemp003 = πF.MakeArgs(2)
									πTemp006 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µtestCaseClass, "testCaseClass"); πE != nil {
										continue
									}
									πTemp006[0] = µtestCaseClass
									if πE = πg.CheckLocal(πF, µattrname, "attrname"); πE != nil {
										continue
									}
									πTemp006[1] = µattrname
									if πTemp004, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
										continue
									}
									if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp006)
									πTemp003[0] = πTemp005
									πTemp003[1] = ß__call__.ToObject()
									if πTemp004, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
										continue
									}
									if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									πTemp001 = πTemp005
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
							µisTestMethod = πTemp001
							// line 147: testFnNames = [x for x in dir(testCaseClass) if isTestMethod(x)]
							πF.SetLineno(147)
							πTemp002 = make([]πg.Param, 0)
							πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/unittest_loader.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µx *πg.Object = πg.UnboundLocal; _ = µx
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 []*πg.Object
								_ = πTemp002
								var πTemp003 *πg.Object
								_ = πTemp003
								var πTemp004 *πg.Object
								_ = πTemp004
								var πTemp005 bool
								_ = πTemp005
								var πTemp006 bool
								_ = πTemp006
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
									for ; πF.State() >= 0; πF.PopCheckpoint() {
										switch πF.State() {
										case 0:
										case 1: goto Label1
										case 2: goto Label2
										case 6: goto Label6
										default: panic("unexpected function state")
										}
										πTemp002 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µtestCaseClass, "testCaseClass"); πE != nil {
											continue
										}
										πTemp002[0] = µtestCaseClass
										if πTemp003, πE = πg.ResolveGlobal(πF, ßdir); πE != nil {
											continue
										}
										if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp002)
										if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
											continue
										}
										πF.PushCheckpoint(2)
										πTemp005 = false
									Label1:
										if πE != nil || πR != nil {
											continue
										}
										if πTemp005 {
											πF.PopCheckpoint()
											goto Label3
										}
										if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
											isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
											if exc != nil {
												πE = exc
											} else if isStop {
												πE = nil
												πF.RestoreExc(nil, nil)
											}
											πTemp006 = !isStop
										} else {
											πTemp006 = true
											µx = πTemp003
										}
										if πE != nil || !πTemp006 {
											continue
										}
										πF.PushCheckpoint(1)            
										πTemp002 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
											continue
										}
										πTemp002[0] = µx
										if πE = πg.CheckLocal(πF, µisTestMethod, "isTestMethod"); πE != nil {
											continue
										}
										if πTemp003, πE = µisTestMethod.Call(πF, πTemp002, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp002)
										if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
											continue
										}
										if πTemp006 {
											goto Label4
										}
										goto Label5
										// line 147: testFnNames = [x for x in dir(testCaseClass) if isTestMethod(x)]
										πF.SetLineno(147)
									Label4:
										// line 147: testFnNames = [x for x in dir(testCaseClass) if isTestMethod(x)]
										πF.SetLineno(147)
										if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
											continue
										}
										πF.PushCheckpoint(6)
										return µx, nil
									Label6:
										πTemp003 = πSent
										goto Label5
									Label5:
										continue
									Label2:
										if πE != nil || πR != nil {
											continue
										}
									Label3:
									}
									return nil, πE
								}).ToObject(), nil
							}), πF.Globals()).ToObject()
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
								continue
							}
							µtestFnNames = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßsortTestMethodsUsing, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 148: if self.sortTestMethodsUsing:
							πF.SetLineno(148)
						Label1:
							// line 149: testFnNames.sort(key=_CmpToKey(self.sortTestMethodsUsing))
							πF.SetLineno(149)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßsortTestMethodsUsing, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_CmpToKey); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp008 = πg.KWArgs{
								{"key", πTemp005},
							}
							if πE = πg.CheckLocal(πF, µtestFnNames, "testFnNames"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µtestFnNames, ßsort, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, πTemp008); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 150: return testFnNames
							πF.SetLineno(150)
							if πE = πg.CheckLocal(πF, µtestFnNames, "testFnNames"); πE != nil {
								continue
							}
							πR = µtestFnNames
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßgetTestCaseNames.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 152: def discover(self, start_dir, pattern='test*.py', top_level_dir=None):
					πF.SetLineno(152)
					πTemp003 = make([]πg.Param, 4)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "start_dir", Def: nil}
					πTemp003[2] = πg.Param{Name: "pattern", Def: πg.NewStr("test*.py").ToObject()}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp003[3] = πg.Param{Name: "top_level_dir", Def: πTemp008}
					πTemp007 = πg.NewFunction(πg.NewCode("discover", "build/src/__python__/unittest_loader.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstart_dir *πg.Object = πArgs[1]; _ = µstart_dir
						var µpattern *πg.Object = πArgs[2]; _ = µpattern
						var µtop_level_dir *πg.Object = πArgs[3]; _ = µtop_level_dir
						var µset_implicit_top *πg.Object = πg.UnboundLocal; _ = µset_implicit_top
						var µis_not_importable *πg.Object = πg.UnboundLocal; _ = µis_not_importable
						var µthe_module *πg.Object = πg.UnboundLocal; _ = µthe_module
						var µtop_part *πg.Object = πg.UnboundLocal; _ = µtop_part
						var µtests *πg.Object = πg.UnboundLocal; _ = µtests
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 *πg.BaseException
						_ = πTemp008
						var πTemp009 *πg.Traceback
						_ = πTemp009
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 13: goto Label13
							default: panic("unexpected function state")
							}
							// line 153: """Find and return all test modules from the specified start
							πF.SetLineno(153)
							// line 173: set_implicit_top = False
							πF.SetLineno(173)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							µset_implicit_top = πTemp001
							if πE = πg.CheckLocal(πF, µtop_level_dir, "top_level_dir"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µtop_level_dir == πTemp004).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_top_level_dir, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp004 != πTemp005).ToObject()
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							if πE = πg.CheckLocal(πF, µtop_level_dir, "top_level_dir"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µtop_level_dir == πTemp003).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 174: if top_level_dir is None and self._top_level_dir is not None:
							πF.SetLineno(174)
						Label2:
							// line 176: top_level_dir = self._top_level_dir
							πF.SetLineno(176)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_top_level_dir, nil); πE != nil {
								continue
							}
							µtop_level_dir = πTemp001
							goto Label4
							// line 177: elif top_level_dir is None:
							πF.SetLineno(177)
						Label3:
							// line 178: set_implicit_top = True
							πF.SetLineno(178)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µset_implicit_top = πTemp001
							// line 179: top_level_dir = start_dir
							πF.SetLineno(179)
							if πE = πg.CheckLocal(πF, µstart_dir, "start_dir"); πE != nil {
								continue
							}
							µtop_level_dir = µstart_dir
							goto Label4
						Label4:
							// line 181: top_level_dir = os.path.abspath(top_level_dir)
							πF.SetLineno(181)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtop_level_dir, "top_level_dir"); πE != nil {
								continue
							}
							πTemp006[0] = µtop_level_dir
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßabspath, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µtop_level_dir = πTemp003
							if πE = πg.CheckLocal(πF, µtop_level_dir, "top_level_dir"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßpath, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, πTemp005, µtop_level_dir); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label5
							}
							goto Label6
							// line 183: if not top_level_dir in sys.path:
							πF.SetLineno(183)
						Label5:
							// line 188: sys.path.insert(0, top_level_dir)
							πF.SetLineno(188)
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µtop_level_dir, "top_level_dir"); πE != nil {
								continue
							}
							πTemp006[1] = µtop_level_dir
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label6
						Label6:
							// line 189: self._top_level_dir = top_level_dir
							πF.SetLineno(189)
							if πE = πg.CheckLocal(πF, µtop_level_dir, "top_level_dir"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µtop_level_dir); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_top_level_dir, πTemp001); πE != nil {
								continue
							}
							// line 191: is_not_importable = False
							πF.SetLineno(191)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							µis_not_importable = πTemp001
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstart_dir, "start_dir"); πE != nil {
								continue
							}
							πTemp007[0] = µstart_dir
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßabspath, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp006[0] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßisdir, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label7
							}
							goto Label8
							// line 192: if os.path.isdir(os.path.abspath(start_dir)):
							πF.SetLineno(192)
						Label7:
							// line 193: start_dir = os.path.abspath(start_dir)
							πF.SetLineno(193)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstart_dir, "start_dir"); πE != nil {
								continue
							}
							πTemp006[0] = µstart_dir
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßabspath, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µstart_dir = πTemp003
							if πE = πg.CheckLocal(πF, µstart_dir, "start_dir"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtop_level_dir, "top_level_dir"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.NE(πF, µstart_dir, µtop_level_dir); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label10
							}
							goto Label11
							// line 194: if start_dir != top_level_dir:
							πF.SetLineno(194)
						Label10:
							// line 195: is_not_importable = not os.path.isfile(os.path.join(start_dir, '__init__.py'))
							πF.SetLineno(195)
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µstart_dir, "start_dir"); πE != nil {
								continue
							}
							πTemp007[0] = µstart_dir
							πTemp007[1] = πg.NewStr("__init__.py").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp006[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßisfile, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp002).ToObject()
							µis_not_importable = πTemp001
							goto Label11
						Label11:
							goto Label9
						Label8:
							// line 198: try:
							πF.SetLineno(198)
							πF.PushCheckpoint(13)
							// line 199: __import__(start_dir)
							πF.SetLineno(199)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstart_dir, "start_dir"); πE != nil {
								continue
							}
							πTemp006[0] = µstart_dir
							if πTemp001, πE = πg.ResolveGlobal(πF, ß__import__); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πF.PopCheckpoint()
							// line 203: the_module = sys.modules[start_dir]
							πF.SetLineno(203)
							if πE = πg.CheckLocal(πF, µstart_dir, "start_dir"); πE != nil {
								continue
							}
							πTemp001 = µstart_dir
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßmodules, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							µthe_module = πTemp003
							// line 204: top_part = start_dir.split('.')[0]
							πF.SetLineno(204)
							πTemp001 = πg.NewInt(0).ToObject()
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr(".").ToObject()
							if πE = πg.CheckLocal(πF, µstart_dir, "start_dir"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µstart_dir, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							µtop_part = πTemp003
							// line 205: start_dir = os.path.abspath(os.path.dirname((the_module.__file__)))
							πF.SetLineno(205)
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µthe_module, "the_module"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µthe_module, ß__file__, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßdirname, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp006[0] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßabspath, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µstart_dir = πTemp003
							if πE = πg.CheckLocal(πF, µset_implicit_top, "set_implicit_top"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µset_implicit_top); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label14
							}
							goto Label15
							// line 206: if set_implicit_top:
							πF.SetLineno(206)
						Label14:
							// line 207: self._top_level_dir = self._get_directory_containing_module(top_part)
							πF.SetLineno(207)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtop_part, "top_part"); πE != nil {
								continue
							}
							πTemp006[0] = µtop_part
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_get_directory_containing_module, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_top_level_dir, πTemp001); πE != nil {
								continue
							}
							// line 208: sys.path.remove(top_level_dir)
							πF.SetLineno(208)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtop_level_dir, "top_level_dir"); πE != nil {
								continue
							}
							πTemp006[0] = µtop_level_dir
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßremove, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label15
						Label15:
							goto Label12
						Label13:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label16
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 200: except ImportError:
							πF.SetLineno(200)
						Label16:
							// line 201: is_not_importable = True
							πF.SetLineno(201)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µis_not_importable = πTemp001
							πF.RestoreExc(nil, nil)
							goto Label12
						Label12:
							goto Label9
						Label9:
							if πE = πg.CheckLocal(πF, µis_not_importable, "is_not_importable"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µis_not_importable); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label17
							}
							goto Label18
							// line 210: if is_not_importable:
							πF.SetLineno(210)
						Label17:
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstart_dir, "start_dir"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Start directory is not importable: %r").ToObject(), µstart_dir); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 211: raise ImportError('Start directory is not importable: %r' % start_dir)
							πF.SetLineno(211)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label18
						Label18:
							// line 213: tests = list(self._find_tests(start_dir, pattern))
							πF.SetLineno(213)
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µstart_dir, "start_dir"); πE != nil {
								continue
							}
							πTemp007[0] = µstart_dir
							if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
								continue
							}
							πTemp007[1] = µpattern
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_find_tests, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp006[0] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µtests = πTemp003
							// line 214: return self.suiteClass(tests)
							πF.SetLineno(214)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtests, "tests"); πE != nil {
								continue
							}
							πTemp006[0] = µtests
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsuiteClass, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πR = πTemp003
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdiscover.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 216: def _get_directory_containing_module(self, module_name):
					πF.SetLineno(216)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "module_name", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("_get_directory_containing_module", "build/src/__python__/unittest_loader.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmodule_name *πg.Object = πArgs[1]; _ = µmodule_name
						var µmodule *πg.Object = πg.UnboundLocal; _ = µmodule
						var µfull_path *πg.Object = πg.UnboundLocal; _ = µfull_path
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 217: module = sys.modules[module_name]
							πF.SetLineno(217)
							if πE = πg.CheckLocal(πF, µmodule_name, "module_name"); πE != nil {
								continue
							}
							πTemp001 = µmodule_name
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßmodules, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µmodule = πTemp002
							// line 218: full_path = os.path.abspath(module.__file__)
							πF.SetLineno(218)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmodule, ß__file__, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßabspath, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µfull_path = πTemp002
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("__init__.py").ToObject()
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfull_path, "full_path"); πE != nil {
								continue
							}
							πTemp006[0] = µfull_path
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßbasename, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßlower, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßstartswith, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label1
							}
							goto Label2
							// line 220: if os.path.basename(full_path).lower().startswith('__init__.py'):
							πF.SetLineno(220)
						Label1:
							// line 221: return os.path.dirname(os.path.dirname(full_path))
							πF.SetLineno(221)
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfull_path, "full_path"); πE != nil {
								continue
							}
							πTemp006[0] = µfull_path
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßdirname, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßdirname, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 226: return os.path.dirname(full_path)
							πF.SetLineno(226)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfull_path, "full_path"); πE != nil {
								continue
							}
							πTemp005[0] = µfull_path
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßdirname, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πR = πTemp002
							continue
							goto Label3
						Label3:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_get_directory_containing_module.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 228: def _get_name_from_path(self, path):
					πF.SetLineno(228)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "path", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("_get_name_from_path", "build/src/__python__/unittest_loader.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µpath *πg.Object = πArgs[1]; _ = µpath
						var µ_relpath *πg.Object = πg.UnboundLocal; _ = µ_relpath
						var µname *πg.Object = πg.UnboundLocal; _ = µname
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 229: path = os.path.splitext(os.path.normpath(path))[0]
							πF.SetLineno(229)
							πTemp001 = πg.NewInt(0).ToObject()
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp004[0] = µpath
							if πTemp005, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßpath, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßnormpath, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp006
							if πTemp005, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßpath, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßsplitext, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
								continue
							}
							µpath = πTemp002
							// line 231: _relpath = os.path.relpath(path, self._top_level_dir)
							πF.SetLineno(231)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp003[0] = µpath
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_top_level_dir, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßrelpath, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µ_relpath = πTemp002
							// line 232: assert not os.path.isabs(_relpath), "Path must be within the project"
							πF.SetLineno(232)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µ_relpath, "_relpath"); πE != nil {
								continue
							}
							πTemp003[0] = µ_relpath
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp005, ßisabs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp007).ToObject()
							if πE = πg.Assert(πF, πTemp001, πg.NewStr("Path must be within the project").ToObject()); πE != nil {
								continue
							}
							// line 233: assert not _relpath.startswith('..'), "Path must be within the project"
							πF.SetLineno(233)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("..").ToObject()
							if πE = πg.CheckLocal(πF, µ_relpath, "_relpath"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µ_relpath, ßstartswith, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp007).ToObject()
							if πE = πg.Assert(πF, πTemp001, πg.NewStr("Path must be within the project").ToObject()); πE != nil {
								continue
							}
							// line 235: name = _relpath.replace(os.path.sep, '.')
							πF.SetLineno(235)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßsep, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							πTemp003[1] = πg.NewStr(".").ToObject()
							if πE = πg.CheckLocal(πF, µ_relpath, "_relpath"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µ_relpath, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µname = πTemp002
							// line 236: return name
							πF.SetLineno(236)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πR = µname
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_get_name_from_path.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 238: def _get_module_from_name(self, name):
					πF.SetLineno(238)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "name", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("_get_module_from_name", "build/src/__python__/unittest_loader.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
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
							// line 239: __import__(name)
							πF.SetLineno(239)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001[0] = µname
							if πTemp002, πE = πg.ResolveGlobal(πF, ß__import__); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 240: return sys.modules[name]
							πF.SetLineno(240)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp002 = µname
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßmodules, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							πR = πTemp003
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_get_module_from_name.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 242: def _match_path(self, path, full_path, pattern):
					πF.SetLineno(242)
					πTemp003 = make([]πg.Param, 4)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "path", Def: nil}
					πTemp003[2] = πg.Param{Name: "full_path", Def: nil}
					πTemp003[3] = πg.Param{Name: "pattern", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("_match_path", "build/src/__python__/unittest_loader.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µpath *πg.Object = πArgs[1]; _ = µpath
						var µfull_path *πg.Object = πArgs[2]; _ = µfull_path
						var µpattern *πg.Object = πArgs[3]; _ = µpattern
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
							// line 244: return fnmatch(path, pattern)
							πF.SetLineno(244)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp001[0] = µpath
							if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
								continue
							}
							πTemp001[1] = µpattern
							if πTemp002, πE = πg.ResolveGlobal(πF, ßfnmatch); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp003
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_match_path.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 246: def _find_tests(self, start_dir, pattern):
					πF.SetLineno(246)
					πTemp003 = make([]πg.Param, 3)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "start_dir", Def: nil}
					πTemp003[2] = πg.Param{Name: "pattern", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("_find_tests", "build/src/__python__/unittest_loader.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstart_dir *πg.Object = πArgs[1]; _ = µstart_dir
						var µpattern *πg.Object = πArgs[2]; _ = µpattern
						var µpaths *πg.Object = πg.UnboundLocal; _ = µpaths
						var µpath *πg.Object = πg.UnboundLocal; _ = µpath
						var µfull_path *πg.Object = πg.UnboundLocal; _ = µfull_path
						var µname *πg.Object = πg.UnboundLocal; _ = µname
						var µmodule *πg.Object = πg.UnboundLocal; _ = µmodule
						var µmod_file *πg.Object = πg.UnboundLocal; _ = µmod_file
						var µrealpath *πg.Object = πg.UnboundLocal; _ = µrealpath
						var µfullpath_noext *πg.Object = πg.UnboundLocal; _ = µfullpath_noext
						var µmodule_dir *πg.Object = πg.UnboundLocal; _ = µmodule_dir
						var µmod_name *πg.Object = πg.UnboundLocal; _ = µmod_name
						var µexpected_dir *πg.Object = πg.UnboundLocal; _ = µexpected_dir
						var µmsg *πg.Object = πg.UnboundLocal; _ = µmsg
						var µload_tests *πg.Object = πg.UnboundLocal; _ = µload_tests
						var µtests *πg.Object = πg.UnboundLocal; _ = µtests
						var µpackage *πg.Object = πg.UnboundLocal; _ = µpackage
						var µtest *πg.Object = πg.UnboundLocal; _ = µtest
						var µe *πg.Object = πg.UnboundLocal; _ = µe
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.BaseException
						_ = πTemp010
						var πTemp011 *πg.Traceback
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πTemp013 πg.KWArgs
						_ = πTemp013
						var πTemp014 bool
						_ = πTemp014
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 36: goto Label36
								case 33: goto Label33
								case 12: goto Label12
								case 34: goto Label34
								case 15: goto Label15
								case 17: goto Label17
								case 27: goto Label27
								case 28: goto Label28
								case 29: goto Label29
								case 31: goto Label31
								default: panic("unexpected function state")
								}
								// line 247: """Used by discovery. Yields test suites it loads."""
								πF.SetLineno(247)
								// line 248: paths = os.listdir(start_dir)
								πF.SetLineno(248)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µstart_dir, "start_dir"); πE != nil {
									continue
								}
								πTemp001[0] = µstart_dir
								if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlistdir, nil); πE != nil {
									continue
								}
								if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µpaths = πTemp002
								if πE = πg.CheckLocal(πF, µpaths, "paths"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Iter(πF, µpaths); πE != nil {
									continue
								}
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
								if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp005 = !isStop
								} else {
									πTemp005 = true
									µpath = πTemp003
								}
								if πE != nil || !πTemp005 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 251: full_path = os.path.join(start_dir, path)
								πF.SetLineno(251)
								πTemp001 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µstart_dir, "start_dir"); πE != nil {
									continue
								}
								πTemp001[0] = µstart_dir
								if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
									continue
								}
								πTemp001[1] = µpath
								if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, πTemp006, ßjoin, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µfull_path = πTemp006
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µfull_path, "full_path"); πE != nil {
									continue
								}
								πTemp001[0] = µfull_path
								if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, πTemp006, ßisfile, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label4
								}
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µfull_path, "full_path"); πE != nil {
									continue
								}
								πTemp001[0] = µfull_path
								if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, πTemp006, ßisdir, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label5
								}
								goto Label6
								// line 252: if os.path.isfile(full_path):
								πF.SetLineno(252)
							Label4:
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
									continue
								}
								πTemp001[0] = µpath
								if πTemp006, πE = πg.ResolveGlobal(πF, ßVALID_MODULE_NAME); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßmatch, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								πTemp003 = πg.GetBool(!πTemp005).ToObject()
								if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label7
								}
								goto Label8
								// line 253: if not VALID_MODULE_NAME.match(path):
								πF.SetLineno(253)
							Label7:
								// line 255: continue
								πF.SetLineno(255)
								continue
								goto Label8
							Label8:
								πTemp001 = πF.MakeArgs(3)
								if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
									continue
								}
								πTemp001[0] = µpath
								if πE = πg.CheckLocal(πF, µfull_path, "full_path"); πE != nil {
									continue
								}
								πTemp001[1] = µfull_path
								if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
									continue
								}
								πTemp001[2] = µpattern
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µself, ß_match_path, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp005, πE = πg.IsTrue(πF, πTemp007); πE != nil {
									continue
								}
								πTemp003 = πg.GetBool(!πTemp005).ToObject()
								if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label9
								}
								goto Label10
								// line 256: if not self._match_path(path, full_path, pattern):
								πF.SetLineno(256)
							Label9:
								// line 257: continue
								πF.SetLineno(257)
								continue
								goto Label10
							Label10:
								// line 259: name = self._get_name_from_path(full_path)
								πF.SetLineno(259)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µfull_path, "full_path"); πE != nil {
									continue
								}
								πTemp001[0] = µfull_path
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µself, ß_get_name_from_path, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µname = πTemp006
								// line 260: try:
								πF.SetLineno(260)
								πF.PushCheckpoint(12)
								// line 261: module = self._get_module_from_name(name)
								πF.SetLineno(261)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
									continue
								}
								πTemp001[0] = µname
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µself, ß_get_module_from_name, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µmodule = πTemp006
								πF.PopCheckpoint()
								// line 265: mod_file = os.path.abspath(getattr(module, '__file__', full_path))
								πF.SetLineno(265)
								πTemp001 = πF.MakeArgs(1)
								πTemp008 = πF.MakeArgs(3)
								if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
									continue
								}
								πTemp008[0] = µmodule
								πTemp008[1] = ß__file__.ToObject()
								if πE = πg.CheckLocal(πF, µfull_path, "full_path"); πE != nil {
									continue
								}
								πTemp008[2] = µfull_path
								if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp008)
								πTemp001[0] = πTemp006
								if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, πTemp006, ßabspath, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µmod_file = πTemp006
								// line 266: realpath = os.path.splitext(os.path.realpath(mod_file))[0]
								πF.SetLineno(266)
								πTemp003 = πg.NewInt(0).ToObject()
								πTemp001 = πF.MakeArgs(1)
								πTemp008 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µmod_file, "mod_file"); πE != nil {
									continue
								}
								πTemp008[0] = µmod_file
								if πTemp007, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßpath, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, πTemp009, ßrealpath, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp007.Call(πF, πTemp008, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp008)
								πTemp001[0] = πTemp009
								if πTemp007, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßpath, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, πTemp009, ßsplitext, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp006, πE = πg.GetItem(πF, πTemp009, πTemp003); πE != nil {
									continue
								}
								µrealpath = πTemp006
								// line 267: fullpath_noext = os.path.splitext(os.path.realpath(full_path))[0]
								πF.SetLineno(267)
								πTemp003 = πg.NewInt(0).ToObject()
								πTemp001 = πF.MakeArgs(1)
								πTemp008 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µfull_path, "full_path"); πE != nil {
									continue
								}
								πTemp008[0] = µfull_path
								if πTemp007, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßpath, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, πTemp009, ßrealpath, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp007.Call(πF, πTemp008, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp008)
								πTemp001[0] = πTemp009
								if πTemp007, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßpath, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, πTemp009, ßsplitext, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp006, πE = πg.GetItem(πF, πTemp009, πTemp003); πE != nil {
									continue
								}
								µfullpath_noext = πTemp006
								if πE = πg.CheckLocal(πF, µrealpath, "realpath"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µrealpath, ßlower, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µfullpath_noext, "fullpath_noext"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µfullpath_noext, ßlower, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp006.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πg.NE(πF, πTemp007, πTemp009); πE != nil {
									continue
								}
								if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label13
								}
								goto Label14
								// line 268: if realpath.lower() != fullpath_noext.lower():
								πF.SetLineno(268)
							Label13:
								// line 269: module_dir = os.path.dirname(realpath)
								πF.SetLineno(269)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µrealpath, "realpath"); πE != nil {
									continue
								}
								πTemp001[0] = µrealpath
								if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, πTemp006, ßdirname, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µmodule_dir = πTemp006
								// line 270: mod_name = os.path.splitext(os.path.basename(full_path))[0]
								πF.SetLineno(270)
								πTemp003 = πg.NewInt(0).ToObject()
								πTemp001 = πF.MakeArgs(1)
								πTemp008 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µfull_path, "full_path"); πE != nil {
									continue
								}
								πTemp008[0] = µfull_path
								if πTemp007, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßpath, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, πTemp009, ßbasename, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp007.Call(πF, πTemp008, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp008)
								πTemp001[0] = πTemp009
								if πTemp007, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßpath, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, πTemp009, ßsplitext, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp006, πE = πg.GetItem(πF, πTemp009, πTemp003); πE != nil {
									continue
								}
								µmod_name = πTemp006
								// line 271: expected_dir = os.path.dirname(full_path)
								πF.SetLineno(271)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µfull_path, "full_path"); πE != nil {
									continue
								}
								πTemp001[0] = µfull_path
								if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, πTemp006, ßdirname, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µexpected_dir = πTemp006
								// line 272: msg = ("%r module incorrectly imported from %r. Expected %r. "
								πF.SetLineno(272)
								µmsg = πg.NewStr("%r module incorrectly imported from %r. Expected %r. Is this module globally installed?").ToObject()
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmod_name, "mod_name"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmodule_dir, "module_dir"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µexpected_dir, "expected_dir"); πE != nil {
									continue
								}
								πTemp006 = πg.NewTuple3(µmod_name, µmodule_dir, µexpected_dir).ToObject()
								if πTemp003, πE = πg.Mod(πF, µmsg, πTemp006); πE != nil {
									continue
								}
								πTemp001[0] = πTemp003
								if πTemp003, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								// line 274: raise ImportError(msg % (mod_name, module_dir, expected_dir))
								πF.SetLineno(274)
								πE = πF.Raise(πTemp006, nil, nil)
								continue
								goto Label14
							Label14:
								// line 275: yield self.loadTestsFromModule(module)
								πF.SetLineno(275)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
									continue
								}
								πTemp001[0] = µmodule
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µself, ßloadTestsFromModule, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								πF.PushCheckpoint(15)
								return πTemp006, nil
							Label15:
								πTemp003 = πSent
								goto Label11
							Label12:
								if πE == nil {
								  continue
								}
								πE = nil
								πTemp010, πTemp011 = πF.ExcInfo()
								goto Label16
								// line 262: except:
								πF.SetLineno(262)
							Label16:
								// line 263: yield _make_failed_import_test(name, self.suiteClass)
								πF.SetLineno(263)
								πTemp001 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
									continue
								}
								πTemp001[0] = µname
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µself, ßsuiteClass, nil); πE != nil {
									continue
								}
								πTemp001[1] = πTemp003
								if πTemp003, πE = πg.ResolveGlobal(πF, ß_make_failed_import_test); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								πF.PushCheckpoint(17)
								return πTemp007, nil
							Label17:
								πTemp003 = πSent
								πF.RestoreExc(nil, nil)
								goto Label11
							Label11:
								goto Label6
								// line 276: elif os.path.isdir(full_path):
								πF.SetLineno(276)
							Label5:
								πTemp001 = πF.MakeArgs(1)
								πTemp008 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µfull_path, "full_path"); πE != nil {
									continue
								}
								πTemp008[0] = µfull_path
								πTemp008[1] = πg.NewStr("__init__.py").ToObject()
								if πTemp009, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
									continue
								}
								if πTemp012, πE = πg.GetAttr(πF, πTemp009, ßpath, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, πTemp012, ßjoin, nil); πE != nil {
									continue
								}
								if πTemp012, πE = πTemp009.Call(πF, πTemp008, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp008)
								πTemp001[0] = πTemp012
								if πTemp009, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
									continue
								}
								if πTemp012, πE = πg.GetAttr(πF, πTemp009, ßpath, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, πTemp012, ßisfile, nil); πE != nil {
									continue
								}
								if πTemp012, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp005, πE = πg.IsTrue(πF, πTemp012); πE != nil {
									continue
								}
								πTemp003 = πg.GetBool(!πTemp005).ToObject()
								if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label18
								}
								goto Label19
								// line 277: if not os.path.isfile(os.path.join(full_path, '__init__.py')):
								πF.SetLineno(277)
							Label18:
								// line 278: continue
								πF.SetLineno(278)
								continue
								goto Label19
							Label19:
								// line 280: load_tests = None
								πF.SetLineno(280)
								if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								µload_tests = πTemp003
								// line 281: tests = None
								πF.SetLineno(281)
								if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								µtests = πTemp003
								πTemp001 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
									continue
								}
								πTemp001[0] = µpath
								if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
									continue
								}
								πTemp001[1] = µpattern
								if πTemp003, πE = πg.ResolveGlobal(πF, ßfnmatch); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp005, πE = πg.IsTrue(πF, πTemp009); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label20
								}
								goto Label21
								// line 282: if fnmatch(path, pattern):
								πF.SetLineno(282)
							Label20:
								// line 284: name = self._get_name_from_path(full_path)
								πF.SetLineno(284)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µfull_path, "full_path"); πE != nil {
									continue
								}
								πTemp001[0] = µfull_path
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µself, ß_get_name_from_path, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µname = πTemp009
								// line 285: package = self._get_module_from_name(name)
								πF.SetLineno(285)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
									continue
								}
								πTemp001[0] = µname
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µself, ß_get_module_from_name, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µpackage = πTemp009
								// line 286: load_tests = getattr(package, 'load_tests', None)
								πF.SetLineno(286)
								πTemp001 = πF.MakeArgs(3)
								if πE = πg.CheckLocal(πF, µpackage, "package"); πE != nil {
									continue
								}
								πTemp001[0] = µpackage
								πTemp001[1] = ßload_tests.ToObject()
								if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp001[2] = πTemp003
								if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µload_tests = πTemp009
								// line 287: tests = self.loadTestsFromModule(package, use_load_tests=False)
								πF.SetLineno(287)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µpackage, "package"); πE != nil {
									continue
								}
								πTemp001[0] = µpackage
								if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								πTemp013 = πg.KWArgs{
									{"use_load_tests", πTemp003},
								}
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µself, ßloadTestsFromModule, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp003.Call(πF, πTemp001, πTemp013); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µtests = πTemp009
								goto Label21
							Label21:
								if πE = πg.CheckLocal(πF, µload_tests, "load_tests"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp003 = πg.GetBool(µload_tests == πTemp009).ToObject()
								if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label22
								}
								goto Label23
								// line 289: if load_tests is None:
								πF.SetLineno(289)
							Label22:
								if πE = πg.CheckLocal(πF, µtests, "tests"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp003 = πg.GetBool(µtests != πTemp009).ToObject()
								if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label25
								}
								goto Label26
								// line 290: if tests is not None:
								πF.SetLineno(290)
							Label25:
								// line 292: yield tests
								πF.SetLineno(292)
								if πE = πg.CheckLocal(πF, µtests, "tests"); πE != nil {
									continue
								}
								πF.PushCheckpoint(27)
								return µtests, nil
							Label27:
								πTemp003 = πSent
								goto Label26
							Label26:
								πTemp001 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µfull_path, "full_path"); πE != nil {
									continue
								}
								πTemp001[0] = µfull_path
								if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
									continue
								}
								πTemp001[1] = µpattern
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, µself, ß_find_tests, nil); πE != nil {
									continue
								}
								if πTemp012, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp003, πE = πg.Iter(πF, πTemp012); πE != nil {
									continue
								}
								πF.PushCheckpoint(29)
								πTemp005 = false
							Label28:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp005 {
									πF.PopCheckpoint()
									goto Label30
								}
								if πTemp009, πE = πg.Next(πF, πTemp003); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp014 = !isStop
								} else {
									πTemp014 = true
									µtest = πTemp009
								}
								if πE != nil || !πTemp014 {
									continue
								}
								πF.PushCheckpoint(28)            
								// line 295: yield test
								πF.SetLineno(295)
								if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
									continue
								}
								πF.PushCheckpoint(31)
								return µtest, nil
							Label31:
								πTemp009 = πSent
								continue
							Label29:
								if πE != nil || πR != nil {
									continue
								}
							Label30:
								goto Label24
							Label23:
								// line 297: try:
								πF.SetLineno(297)
								πF.PushCheckpoint(33)
								// line 298: yield load_tests(self, tests, pattern)
								πF.SetLineno(298)
								πTemp001 = πF.MakeArgs(3)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								πTemp001[0] = µself
								if πE = πg.CheckLocal(πF, µtests, "tests"); πE != nil {
									continue
								}
								πTemp001[1] = µtests
								if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
									continue
								}
								πTemp001[2] = µpattern
								if πE = πg.CheckLocal(πF, µload_tests, "load_tests"); πE != nil {
									continue
								}
								if πTemp003, πE = µload_tests.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								πF.PushCheckpoint(34)
								return πTemp003, nil
							Label34:
								πTemp009 = πSent
								πF.PopCheckpoint()
								goto Label32
							Label33:
								if πE == nil {
								  continue
								}
								πE = nil
								πTemp010, πTemp011 = πF.ExcInfo()
								if πTemp009, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
									continue
								}
								if πTemp005, πE = πg.IsInstance(πF, πTemp010.ToObject(), πTemp009); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label35
								}
								πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
								continue
								// line 299: except Exception, e:
								πF.SetLineno(299)
							Label35:
								µe = πTemp010.ToObject()
								// line 300: yield _make_failed_load_tests(package.__name__, e,
								πF.SetLineno(300)
								πTemp001 = πF.MakeArgs(3)
								if πE = πg.CheckLocal(πF, µpackage, "package"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, µpackage, ß__name__, nil); πE != nil {
									continue
								}
								πTemp001[0] = πTemp009
								if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
									continue
								}
								πTemp001[1] = µe
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, µself, ßsuiteClass, nil); πE != nil {
									continue
								}
								πTemp001[2] = πTemp009
								if πTemp009, πE = πg.ResolveGlobal(πF, ß_make_failed_load_tests); πE != nil {
									continue
								}
								if πTemp012, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								πF.PushCheckpoint(36)
								return πTemp012, nil
							Label36:
								πTemp009 = πSent
								πF.RestoreExc(nil, nil)
								goto Label32
							Label32:
								goto Label24
							Label24:
								goto Label6
							Label6:
								continue
							Label2:
								if πE != nil || πR != nil {
									continue
								}
							Label3:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_find_tests.ToObject(), πTemp012); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("TestLoader").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTestLoader.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 303: defaultTestLoader = TestLoader()
			πF.SetLineno(303)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßTestLoader); πE != nil {
				continue
			}
			if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdefaultTestLoader.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 306: def _makeLoader(prefix, sortUsing, suiteClass=None):
			πF.SetLineno(306)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "prefix", Def: nil}
			πTemp004[1] = πg.Param{Name: "sortUsing", Def: nil}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "suiteClass", Def: πTemp008}
			πTemp007 = πg.NewFunction(πg.NewCode("_makeLoader", "build/src/__python__/unittest_loader.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µprefix *πg.Object = πArgs[0]; _ = µprefix
				var µsortUsing *πg.Object = πArgs[1]; _ = µsortUsing
				var µsuiteClass *πg.Object = πArgs[2]; _ = µsuiteClass
				var µloader *πg.Object = πg.UnboundLocal; _ = µloader
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 307: loader = TestLoader()
					πF.SetLineno(307)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTestLoader); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µloader = πTemp002
					// line 308: loader.sortTestMethodsUsing = sortUsing
					πF.SetLineno(308)
					if πE = πg.CheckLocal(πF, µsortUsing, "sortUsing"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µsortUsing); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µloader, "loader"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µloader, ßsortTestMethodsUsing, πTemp001); πE != nil {
						continue
					}
					// line 309: loader.testMethodPrefix = prefix
					πF.SetLineno(309)
					if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µprefix); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µloader, "loader"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µloader, ßtestMethodPrefix, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsuiteClass, "suiteClass"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µsuiteClass); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 310: if suiteClass:
					πF.SetLineno(310)
				Label1:
					// line 311: loader.suiteClass = suiteClass
					πF.SetLineno(311)
					if πE = πg.CheckLocal(πF, µsuiteClass, "suiteClass"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µsuiteClass); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µloader, "loader"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µloader, ßsuiteClass, πTemp001); πE != nil {
						continue
					}
					goto Label2
				Label2:
					// line 312: return loader
					πF.SetLineno(312)
					if πE = πg.CheckLocal(πF, µloader, "loader"); πE != nil {
						continue
					}
					πR = µloader
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_makeLoader.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 314: def getTestCaseNames(testCaseClass, prefix, sortUsing=cmp):
			πF.SetLineno(314)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "testCaseClass", Def: nil}
			πTemp004[1] = πg.Param{Name: "prefix", Def: nil}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßcmp); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "sortUsing", Def: πTemp009}
			πTemp008 = πg.NewFunction(πg.NewCode("getTestCaseNames", "build/src/__python__/unittest_loader.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtestCaseClass *πg.Object = πArgs[0]; _ = µtestCaseClass
				var µprefix *πg.Object = πArgs[1]; _ = µprefix
				var µsortUsing *πg.Object = πArgs[2]; _ = µsortUsing
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
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
					// line 315: return _makeLoader(prefix, sortUsing).getTestCaseNames(testCaseClass)
					πF.SetLineno(315)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtestCaseClass, "testCaseClass"); πE != nil {
						continue
					}
					πTemp001[0] = µtestCaseClass
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
						continue
					}
					πTemp002[0] = µprefix
					if πE = πg.CheckLocal(πF, µsortUsing, "sortUsing"); πE != nil {
						continue
					}
					πTemp002[1] = µsortUsing
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_makeLoader); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßgetTestCaseNames, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp004
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßgetTestCaseNames.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 317: def makeSuite(testCaseClass, prefix='test', sortUsing=cmp,
			πF.SetLineno(317)
			πTemp004 = make([]πg.Param, 4)
			πTemp004[0] = πg.Param{Name: "testCaseClass", Def: nil}
			πTemp004[1] = πg.Param{Name: "prefix", Def: ßtest.ToObject()}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßcmp); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "sortUsing", Def: πTemp010}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßsuite); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßTestSuite, nil); πE != nil {
				continue
			}
			πTemp004[3] = πg.Param{Name: "suiteClass", Def: πTemp011}
			πTemp009 = πg.NewFunction(πg.NewCode("makeSuite", "build/src/__python__/unittest_loader.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtestCaseClass *πg.Object = πArgs[0]; _ = µtestCaseClass
				var µprefix *πg.Object = πArgs[1]; _ = µprefix
				var µsortUsing *πg.Object = πArgs[2]; _ = µsortUsing
				var µsuiteClass *πg.Object = πArgs[3]; _ = µsuiteClass
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
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
					// line 319: return _makeLoader(prefix, sortUsing, suiteClass).loadTestsFromTestCase(testCaseClass)
					πF.SetLineno(319)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtestCaseClass, "testCaseClass"); πE != nil {
						continue
					}
					πTemp001[0] = µtestCaseClass
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
						continue
					}
					πTemp002[0] = µprefix
					if πE = πg.CheckLocal(πF, µsortUsing, "sortUsing"); πE != nil {
						continue
					}
					πTemp002[1] = µsortUsing
					if πE = πg.CheckLocal(πF, µsuiteClass, "suiteClass"); πE != nil {
						continue
					}
					πTemp002[2] = µsuiteClass
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_makeLoader); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßloadTestsFromTestCase, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp004
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßmakeSuite.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 321: def findTestCases(module, prefix='test', sortUsing=cmp,
			πF.SetLineno(321)
			πTemp004 = make([]πg.Param, 4)
			πTemp004[0] = πg.Param{Name: "module", Def: nil}
			πTemp004[1] = πg.Param{Name: "prefix", Def: ßtest.ToObject()}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßcmp); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "sortUsing", Def: πTemp011}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßsuite); πE != nil {
				continue
			}
			if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßTestSuite, nil); πE != nil {
				continue
			}
			πTemp004[3] = πg.Param{Name: "suiteClass", Def: πTemp012}
			πTemp010 = πg.NewFunction(πg.NewCode("findTestCases", "build/src/__python__/unittest_loader.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmodule *πg.Object = πArgs[0]; _ = µmodule
				var µprefix *πg.Object = πArgs[1]; _ = µprefix
				var µsortUsing *πg.Object = πArgs[2]; _ = µsortUsing
				var µsuiteClass *πg.Object = πArgs[3]; _ = µsuiteClass
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
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
					// line 323: return _makeLoader(prefix, sortUsing, suiteClass).loadTestsFromModule(module)
					πF.SetLineno(323)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					πTemp001[0] = µmodule
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
						continue
					}
					πTemp002[0] = µprefix
					if πE = πg.CheckLocal(πF, µsortUsing, "sortUsing"); πE != nil {
						continue
					}
					πTemp002[1] = µsortUsing
					if πE = πg.CheckLocal(πF, µsuiteClass, "suiteClass"); πE != nil {
						continue
					}
					πTemp002[2] = µsuiteClass
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_makeLoader); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßloadTestsFromModule, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp004
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßfindTestCases.ToObject(), πTemp010); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("unittest_loader", Code)
}
