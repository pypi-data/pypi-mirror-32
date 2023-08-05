package unittest_suite
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/unittest_suite.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßBaseTestSuite := πg.InternStr("BaseTestSuite")
		ßException := πg.InternStr("Exception")
		ßFalse := πg.InternStr("False")
		ßKeyError := πg.InternStr("KeyError")
		ßNone := πg.InternStr("None")
		ßNotImplemented := πg.InternStr("NotImplemented")
		ßSkipTest := πg.InternStr("SkipTest")
		ßTestCase := πg.InternStr("TestCase")
		ßTestSuite := πg.InternStr("TestSuite")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ß_DebugResult := πg.InternStr("_DebugResult")
		ß_ErrorHolder := πg.InternStr("_ErrorHolder")
		ß__call__ := πg.InternStr("__call__")
		ß__class__ := πg.InternStr("__class__")
		ß__eq__ := πg.InternStr("__eq__")
		ß__hash__ := πg.InternStr("__hash__")
		ß__init__ := πg.InternStr("__init__")
		ß__iter__ := πg.InternStr("__iter__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__ne__ := πg.InternStr("__ne__")
		ß__repr__ := πg.InternStr("__repr__")
		ß__str__ := πg.InternStr("__str__")
		ß__unittest := πg.InternStr("__unittest")
		ß__unittest_skip__ := πg.InternStr("__unittest_skip__")
		ß_addClassOrModuleLevelException := πg.InternStr("_addClassOrModuleLevelException")
		ß_call_if_exists := πg.InternStr("_call_if_exists")
		ß_classSetupFailed := πg.InternStr("_classSetupFailed")
		ß_get_previous_module := πg.InternStr("_get_previous_module")
		ß_handleClassSetUp := πg.InternStr("_handleClassSetUp")
		ß_handleModuleFixture := πg.InternStr("_handleModuleFixture")
		ß_handleModuleTearDown := πg.InternStr("_handleModuleTearDown")
		ß_isnotsuite := πg.InternStr("_isnotsuite")
		ß_moduleSetUpFailed := πg.InternStr("_moduleSetUpFailed")
		ß_previousTestClass := πg.InternStr("_previousTestClass")
		ß_restoreStdout := πg.InternStr("_restoreStdout")
		ß_setupStdout := πg.InternStr("_setupStdout")
		ß_tearDownPreviousClass := πg.InternStr("_tearDownPreviousClass")
		ß_testRunEntered := πg.InternStr("_testRunEntered")
		ß_tests := πg.InternStr("_tests")
		ßaddError := πg.InternStr("addError")
		ßaddSkip := πg.InternStr("addSkip")
		ßaddTest := πg.InternStr("addTest")
		ßaddTests := πg.InternStr("addTests")
		ßappend := πg.InternStr("append")
		ßbasestring := πg.InternStr("basestring")
		ßcase := πg.InternStr("case")
		ßcountTestCases := πg.InternStr("countTestCases")
		ßdebug := πg.InternStr("debug")
		ßdescription := πg.InternStr("description")
		ßexc_info := πg.InternStr("exc_info")
		ßfailureException := πg.InternStr("failureException")
		ßgetattr := πg.InternStr("getattr")
		ßhasattr := πg.InternStr("hasattr")
		ßid := πg.InternStr("id")
		ßisinstance := πg.InternStr("isinstance")
		ßissubclass := πg.InternStr("issubclass")
		ßiter := πg.InternStr("iter")
		ßlist := πg.InternStr("list")
		ßmodules := πg.InternStr("modules")
		ßobject := πg.InternStr("object")
		ßrepr := πg.InternStr("repr")
		ßrun := πg.InternStr("run")
		ßsetUpClass := πg.InternStr("setUpClass")
		ßsetUpModule := πg.InternStr("setUpModule")
		ßshortDescription := πg.InternStr("shortDescription")
		ßshouldStop := πg.InternStr("shouldStop")
		ßstr := πg.InternStr("str")
		ßstrclass := πg.InternStr("strclass")
		ßsys := πg.InternStr("sys")
		ßtearDownClass := πg.InternStr("tearDownClass")
		ßtearDownModule := πg.InternStr("tearDownModule")
		ßtype := πg.InternStr("type")
		ßutil := πg.InternStr("util")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 []πg.Param
		_ = πTemp003
		var πTemp004 *πg.Dict
		_ = πTemp004
		var πTemp005 *πg.Object
		_ = πTemp005
		var πTemp006 *πg.Object
		_ = πTemp006
		var πTemp007 *πg.Object
		_ = πTemp007
		var πTemp008 *πg.Object
		_ = πTemp008
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """TestSuite"""
			πF.SetLineno(1)
			// line 3: import sys
			πF.SetLineno(3)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 7: import unittest_case as case
			πF.SetLineno(7)
			if πTemp002, πE = πg.ImportModule(πF, "unittest_case"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßcase.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 8: import unittest_util as util
			πF.SetLineno(8)
			if πTemp002, πE = πg.ImportModule(πF, "unittest_util"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßutil.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 10: __unittest = True
			πF.SetLineno(10)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__unittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 13: def _call_if_exists(parent, attr):
			πF.SetLineno(13)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "parent", Def: nil}
			πTemp003[1] = πg.Param{Name: "attr", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("_call_if_exists", "build/src/__python__/unittest_suite.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µparent *πg.Object = πArgs[0]; _ = µparent
				var µattr *πg.Object = πArgs[1]; _ = µattr
				var µfunc *πg.Object = πg.UnboundLocal; _ = µfunc
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []πg.Param
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
					// line 14: func = getattr(parent, attr, lambda: None)
					πF.SetLineno(14)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
						continue
					}
					πTemp001[0] = µparent
					if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
						continue
					}
					πTemp001[1] = µattr
					πTemp003 = make([]πg.Param, 0)
					πTemp002 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/unittest_suite.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 14: func = getattr(parent, attr, lambda: None)
							πF.SetLineno(14)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
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
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µfunc = πTemp004
					// line 15: func()
					πF.SetLineno(15)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp002, πE = µfunc.Call(πF, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_call_if_exists.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: class BaseTestSuite(object):
			πF.SetLineno(18)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp004 = πg.NewDict()
			if πTemp005, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp005); πE != nil {
				continue
			}
			_, πE = πg.NewCode("BaseTestSuite", "build/src/__python__/unittest_suite.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
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
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 *πg.Object
				_ = πTemp012
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 19: """A simple test suite that doesn't provide class or module shared fixtures.
					πF.SetLineno(19)
					// line 21: def __init__(self, tests=()):
					πF.SetLineno(21)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewTuple0().ToObject()
					πTemp002[1] = πg.Param{Name: "tests", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtests *πg.Object = πArgs[1]; _ = µtests
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
							// line 22: self._tests = []
							πF.SetLineno(22)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_tests, πTemp003); πE != nil {
								continue
							}
							// line 23: self.addTests(tests)
							πF.SetLineno(23)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtests, "tests"); πE != nil {
								continue
							}
							πTemp001[0] = µtests
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßaddTests, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 25: def __repr__(self):
					πF.SetLineno(25)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__repr__", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 26: return "<%s tests=%s>" % (util.strclass(self.__class__), list(self))
							πF.SetLineno(26)
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
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002 = πg.NewTuple2(πTemp004, πTemp006).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<%s tests=%s>").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 28: def __eq__(self, other):
					πF.SetLineno(28)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__eq__", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 29: if not isinstance(other, self.__class__):
							πF.SetLineno(29)
						Label1:
							// line 30: return NotImplemented
							πF.SetLineno(30)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 31: return list(self) == list(other)
							πF.SetLineno(31)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Eq(πF, πTemp004, πTemp006); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 33: def __ne__(self, other):
					πF.SetLineno(33)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("__ne__", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							// line 34: return not self == other
							πF.SetLineno(34)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µself, µother); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
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
					if πE = πClass.SetItem(πF, ß__ne__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 37: __hash__ = None
					πF.SetLineno(37)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__hash__.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 39: def __iter__(self):
					πF.SetLineno(39)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("__iter__", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 40: return iter(self._tests)
							πF.SetLineno(40)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tests, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 42: def countTestCases(self):
					πF.SetLineno(42)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("countTestCases", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µcases *πg.Object = πg.UnboundLocal; _ = µcases
						var µtest *πg.Object = πg.UnboundLocal; _ = µtest
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 bool
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
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 43: cases = 0
							πF.SetLineno(43)
							µcases = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µself); πE != nil {
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
								µtest = πTemp004
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 45: cases += test.countTestCases()
							πF.SetLineno(45)
							if πE = πg.CheckLocal(πF, µcases, "cases"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µtest, ßcountTestCases, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IAdd(πF, µcases, πTemp005); πE != nil {
								continue
							}
							µcases = πTemp004
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 46: return cases
							πF.SetLineno(46)
							if πE = πg.CheckLocal(πF, µcases, "cases"); πE != nil {
								continue
							}
							πR = µcases
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcountTestCases.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 48: def addTest(self, test):
					πF.SetLineno(48)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "test", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("addTest", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtest *πg.Object = πArgs[1]; _ = µtest
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp002[0] = µtest
							πTemp002[1] = ß__call__.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 50: if not hasattr(test, '__call__'):
							πF.SetLineno(50)
						Label1:
							πTemp002 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp006[0] = µtest
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s is not callable").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 52: raise TypeError("%s is not callable" % (repr(test)))
							πF.SetLineno(52)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp002[0] = µtest
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πTemp004
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label3
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp002[0] = µtest
							if πTemp004, πE = πg.ResolveGlobal(πF, ßcase); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp004, ßTestCase, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßTestSuite); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp007, πTemp004).ToObject()
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πTemp004
						Label3:
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 53: if isinstance(test, type) and issubclass(test,
							πF.SetLineno(53)
						Label4:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("TestCases and TestSuites must be instantiated before passing them to addTest()").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 55: raise TypeError("TestCases and TestSuites must be instantiated "
							πF.SetLineno(55)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label5
						Label5:
							// line 57: self._tests.append(test)
							πF.SetLineno(57)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp002[0] = µtest
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_tests, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßaddTest.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 59: def addTests(self, tests):
					πF.SetLineno(59)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "tests", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("addTests", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtests *πg.Object = πArgs[1]; _ = µtests
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 3: goto Label3
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtests, "tests"); πE != nil {
								continue
							}
							πTemp001[0] = µtests
							if πTemp002, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
								goto Label1
							}
							goto Label2
							// line 60: if isinstance(tests, basestring):
							πF.SetLineno(60)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("tests must be an iterable of tests, not a string").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 61: raise TypeError("tests must be an iterable of tests, not a string")
							πF.SetLineno(61)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µtests, "tests"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µtests); πE != nil {
								continue
							}
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
								µtest = πTemp003
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 63: self.addTest(test)
							πF.SetLineno(63)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp001[0] = µtest
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßaddTest, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßaddTests.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 65: def run(self, result):
					πF.SetLineno(65)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "result", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("run", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µresult *πg.Object = πArgs[1]; _ = µresult
						var µtest *πg.Object = πg.UnboundLocal; _ = µtest
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µself); πE != nil {
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
								µtest = πTemp004
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µresult, ßshouldStop, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							goto Label5
							// line 67: if result.shouldStop:
							πF.SetLineno(67)
						Label4:
							// line 68: break
							πF.SetLineno(68)
							πTemp002 = true
							continue
							goto Label5
						Label5:
							// line 69: test(result)
							πF.SetLineno(69)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp005[0] = µresult
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							if πTemp004, πE = µtest.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 70: return result
							πF.SetLineno(70)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πR = µresult
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 72: def __call__(self, *args, **kwds):
					πF.SetLineno(72)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("__call__", "build/src/__python__/unittest_suite.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µargs *πg.Object = πArgs[1]; _ = µargs
						var µkwds *πg.Object = πArgs[2]; _ = µkwds
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
							// line 73: return self.run(*args, **kwds)
							πF.SetLineno(73)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßrun, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, πTemp001, nil, µargs, nil, µkwds); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ß__call__.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 75: def debug(self):
					πF.SetLineno(75)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("debug", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtest *πg.Object = πg.UnboundLocal; _ = µtest
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 bool
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
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 76: """Run the tests without collecting errors in a TestResult"""
							πF.SetLineno(76)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µself); πE != nil {
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
								µtest = πTemp004
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 78: test.debug()
							πF.SetLineno(78)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µtest, ßdebug, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdebug.ToObject(), πTemp012); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("BaseTestSuite").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBaseTestSuite.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 81: class TestSuite(BaseTestSuite):
			πF.SetLineno(81)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßBaseTestSuite); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp004 = πg.NewDict()
			if πTemp005, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp005); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TestSuite", "build/src/__python__/unittest_suite.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 82: """A test suite is a composite test consisting of a number of TestCases.
					πF.SetLineno(82)
					// line 91: def run(self, result, debug=False):
					πF.SetLineno(91)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "result", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "debug", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("run", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µresult *πg.Object = πArgs[1]; _ = µresult
						var µdebug *πg.Object = πArgs[2]; _ = µdebug
						var µtopLevel *πg.Object = πg.UnboundLocal; _ = µtopLevel
						var µtest *πg.Object = πg.UnboundLocal; _ = µtest
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 3: goto Label3
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							// line 92: topLevel = False
							πF.SetLineno(92)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							µtopLevel = πTemp001
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp002[0] = µresult
							πTemp002[1] = ß_testRunEntered.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp002[2] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004 == πTemp003).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 93: if getattr(result, '_testRunEntered', False) is False:
							πF.SetLineno(93)
						Label1:
							// line 94: result._testRunEntered = topLevel = True
							πF.SetLineno(94)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µresult, ß_testRunEntered, πTemp003); πE != nil {
								continue
							}
							µtopLevel = πTemp001
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µself); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp005 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label5
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
								µtest = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(3)            
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µresult, ßshouldStop, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label6
							}
							goto Label7
							// line 97: if result.shouldStop:
							πF.SetLineno(97)
						Label6:
							// line 98: break
							πF.SetLineno(98)
							πTemp005 = true
							continue
							goto Label7
						Label7:
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp002[0] = µtest
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_isnotsuite); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label8
							}
							goto Label9
							// line 100: if _isnotsuite(test):
							πF.SetLineno(100)
						Label8:
							// line 101: self._tearDownPreviousClass(test, result)
							πF.SetLineno(101)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp002[0] = µtest
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp002[1] = µresult
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_tearDownPreviousClass, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 102: self._handleModuleFixture(test, result)
							πF.SetLineno(102)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp002[0] = µtest
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp002[1] = µresult
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_handleModuleFixture, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 103: self._handleClassSetUp(test, result)
							πF.SetLineno(103)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp002[0] = µtest
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp002[1] = µresult
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_handleClassSetUp, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 104: result._previousTestClass = test.__class__
							πF.SetLineno(104)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µtest, ß__class__, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µresult, ß_previousTestClass, πTemp004); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µtest, ß__class__, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							πTemp002[1] = ß_classSetupFailed.ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp002[2] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp003 = πTemp007
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label10
							}
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp002[0] = µresult
							πTemp002[1] = ß_moduleSetUpFailed.ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp002[2] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp003 = πTemp007
						Label10:
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label11
							}
							goto Label12
							// line 106: if (getattr(test.__class__, '_classSetupFailed', False) or
							πF.SetLineno(106)
						Label11:
							// line 108: continue
							πF.SetLineno(108)
							continue
							goto Label12
						Label12:
							goto Label9
						Label9:
							if πE = πg.CheckLocal(πF, µdebug, "debug"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µdebug); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label13
							}
							goto Label14
							// line 110: if not debug:
							πF.SetLineno(110)
						Label13:
							// line 111: test(result)
							πF.SetLineno(111)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp002[0] = µresult
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							if πTemp003, πE = µtest.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label15
						Label14:
							// line 113: test.debug()
							πF.SetLineno(113)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µtest, ßdebug, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label15
						Label15:
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							if πE = πg.CheckLocal(πF, µtopLevel, "topLevel"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µtopLevel); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label16
							}
							goto Label17
							// line 115: if topLevel:
							πF.SetLineno(115)
						Label16:
							// line 116: self._tearDownPreviousClass(None, result)
							πF.SetLineno(116)
							πTemp002 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp002[1] = µresult
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_tearDownPreviousClass, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 117: self._handleModuleTearDown(result)
							πF.SetLineno(117)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp002[0] = µresult
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_handleModuleTearDown, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 118: result._testRunEntered = False
							πF.SetLineno(118)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µresult, ß_testRunEntered, πTemp003); πE != nil {
								continue
							}
							goto Label17
						Label17:
							// line 119: return result
							πF.SetLineno(119)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πR = µresult
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 121: def debug(self):
					πF.SetLineno(121)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("debug", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdebug *πg.Object = πg.UnboundLocal; _ = µdebug
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
							// line 122: """Run the tests without collecting errors in a TestResult"""
							πF.SetLineno(122)
							// line 123: debug = _DebugResult()
							πF.SetLineno(123)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_DebugResult); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µdebug = πTemp002
							// line 124: self.run(debug, True)
							πF.SetLineno(124)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdebug, "debug"); πE != nil {
								continue
							}
							πTemp003[0] = µdebug
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßrun, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdebug.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 128: def _handleClassSetUp(self, test, result):
					πF.SetLineno(128)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "test", Def: nil}
					πTemp002[2] = πg.Param{Name: "result", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("_handleClassSetUp", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtest *πg.Object = πArgs[1]; _ = µtest
						var µresult *πg.Object = πArgs[2]; _ = µresult
						var µpreviousClass *πg.Object = πg.UnboundLocal; _ = µpreviousClass
						var µcurrentClass *πg.Object = πg.UnboundLocal; _ = µcurrentClass
						var µsetUpClass *πg.Object = πg.UnboundLocal; _ = µsetUpClass
						var µe *πg.Object = πg.UnboundLocal; _ = µe
						var µclassName *πg.Object = πg.UnboundLocal; _ = µclassName
						var µerrorName *πg.Object = πg.UnboundLocal; _ = µerrorName
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.BaseException
						_ = πTemp005
						var πTemp006 *πg.Traceback
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 8: goto Label8
							case 12: goto Label12
							case 13: goto Label13
							default: panic("unexpected function state")
							}
							// line 129: previousClass = getattr(result, '_previousTestClass', None)
							πF.SetLineno(129)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							πTemp001[1] = ß_previousTestClass.ToObject()
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
							µpreviousClass = πTemp003
							// line 130: currentClass = test.__class__
							πF.SetLineno(130)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtest, ß__class__, nil); πE != nil {
								continue
							}
							µcurrentClass = πTemp002
							if πE = πg.CheckLocal(πF, µcurrentClass, "currentClass"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpreviousClass, "previousClass"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µcurrentClass, µpreviousClass); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 131: if currentClass == previousClass:
							πF.SetLineno(131)
						Label1:
							// line 132: return
							πF.SetLineno(132)
							πR = πg.None
							continue
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µresult, ß_moduleSetUpFailed, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 133: if result._moduleSetUpFailed:
							πF.SetLineno(133)
						Label3:
							// line 134: return
							πF.SetLineno(134)
							πR = πg.None
							continue
							goto Label4
						Label4:
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µcurrentClass, "currentClass"); πE != nil {
								continue
							}
							πTemp001[0] = µcurrentClass
							πTemp001[1] = ß__unittest_skip__.ToObject()
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
								goto Label5
							}
							goto Label6
							// line 135: if getattr(currentClass, "__unittest_skip__", False):
							πF.SetLineno(135)
						Label5:
							// line 136: return
							πF.SetLineno(136)
							πR = πg.None
							continue
							goto Label6
						Label6:
							// line 138: try:
							πF.SetLineno(138)
							πF.PushCheckpoint(8)
							// line 139: currentClass._classSetupFailed = False
							πF.SetLineno(139)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcurrentClass, "currentClass"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µcurrentClass, ß_classSetupFailed, πTemp003); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label7
						Label8:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label9
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 140: except TypeError:
							πF.SetLineno(140)
						Label9:
							// line 143: pass
							πF.SetLineno(143)
							πF.RestoreExc(nil, nil)
							goto Label7
						Label7:
							// line 145: setUpClass = getattr(currentClass, 'setUpClass', None)
							πF.SetLineno(145)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µcurrentClass, "currentClass"); πE != nil {
								continue
							}
							πTemp001[0] = µcurrentClass
							πTemp001[1] = ßsetUpClass.ToObject()
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
							µsetUpClass = πTemp003
							if πE = πg.CheckLocal(πF, µsetUpClass, "setUpClass"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µsetUpClass != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label10
							}
							goto Label11
							// line 146: if setUpClass is not None:
							πF.SetLineno(146)
						Label10:
							// line 147: _call_if_exists(result, '_setupStdout')
							πF.SetLineno(147)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							πTemp001[1] = ß_setupStdout.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_call_if_exists); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 148: try:
							πF.SetLineno(148)
							πF.PushCheckpoint(12)
							πF.PushCheckpoint(13)
							// line 149: setUpClass()
							πF.SetLineno(149)
							if πE = πg.CheckLocal(πF, µsetUpClass, "setUpClass"); πE != nil {
								continue
							}
							if πTemp002, πE = µsetUpClass.Call(πF, nil, nil); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							πF.PopCheckpoint()
							goto Label12
						Label13:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label14
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 150: except Exception as e:
							πF.SetLineno(150)
						Label14:
							µe = πTemp005.ToObject()
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_DebugResult); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
								goto Label15
							}
							goto Label16
							// line 151: if isinstance(result, _DebugResult):
							πF.SetLineno(151)
						Label15:
							// line 152: raise
							πF.SetLineno(152)
							πE = πF.Raise(nil, nil, nil)
							continue
							goto Label16
						Label16:
							// line 153: currentClass._classSetupFailed = True
							πF.SetLineno(153)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcurrentClass, "currentClass"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µcurrentClass, ß_classSetupFailed, πTemp003); πE != nil {
								continue
							}
							// line 154: className = util.strclass(currentClass)
							πF.SetLineno(154)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcurrentClass, "currentClass"); πE != nil {
								continue
							}
							πTemp001[0] = µcurrentClass
							if πTemp002, πE = πg.ResolveGlobal(πF, ßutil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstrclass, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µclassName = πTemp002
							// line 155: errorName = 'setUpClass (%s)' % className
							πF.SetLineno(155)
							if πE = πg.CheckLocal(πF, µclassName, "className"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("setUpClass (%s)").ToObject(), µclassName); πE != nil {
								continue
							}
							µerrorName = πTemp002
							// line 156: self._addClassOrModuleLevelException(result, e, errorName)
							πF.SetLineno(156)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							πTemp001[1] = µe
							if πE = πg.CheckLocal(πF, µerrorName, "errorName"); πE != nil {
								continue
							}
							πTemp001[2] = µerrorName
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_addClassOrModuleLevelException, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πF.RestoreExc(nil, nil)
							πF.PopCheckpoint()
							goto Label12
						Label12:
							πTemp005, πTemp006 = πF.RestoreExc(nil, nil)
							// line 158: _call_if_exists(result, '_restoreStdout')
							πF.SetLineno(158)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							πTemp001[1] = ß_restoreStdout.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_call_if_exists); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005 != nil {
								πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
								continue
							}
							if πR != nil {
								continue
							}
							goto Label11
						Label11:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_handleClassSetUp.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 160: def _get_previous_module(self, result):
					πF.SetLineno(160)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "result", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("_get_previous_module", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µresult *πg.Object = πArgs[1]; _ = µresult
						var µpreviousModule *πg.Object = πg.UnboundLocal; _ = µpreviousModule
						var µpreviousClass *πg.Object = πg.UnboundLocal; _ = µpreviousClass
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
							// line 161: previousModule = None
							πF.SetLineno(161)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µpreviousModule = πTemp001
							// line 162: previousClass = getattr(result, '_previousTestClass', None)
							πF.SetLineno(162)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp002[0] = µresult
							πTemp002[1] = ß_previousTestClass.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µpreviousClass = πTemp003
							if πE = πg.CheckLocal(πF, µpreviousClass, "previousClass"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µpreviousClass != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 163: if previousClass is not None:
							πF.SetLineno(163)
						Label1:
							// line 164: previousModule = previousClass.__module__
							πF.SetLineno(164)
							if πE = πg.CheckLocal(πF, µpreviousClass, "previousClass"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µpreviousClass, ß__module__, nil); πE != nil {
								continue
							}
							µpreviousModule = πTemp001
							goto Label2
						Label2:
							// line 165: return previousModule
							πF.SetLineno(165)
							if πE = πg.CheckLocal(πF, µpreviousModule, "previousModule"); πE != nil {
								continue
							}
							πR = µpreviousModule
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_get_previous_module.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 168: def _handleModuleFixture(self, test, result):
					πF.SetLineno(168)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "test", Def: nil}
					πTemp002[2] = πg.Param{Name: "result", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("_handleModuleFixture", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtest *πg.Object = πArgs[1]; _ = µtest
						var µresult *πg.Object = πArgs[2]; _ = µresult
						var µpreviousModule *πg.Object = πg.UnboundLocal; _ = µpreviousModule
						var µcurrentModule *πg.Object = πg.UnboundLocal; _ = µcurrentModule
						var µmodule *πg.Object = πg.UnboundLocal; _ = µmodule
						var µsetUpModule *πg.Object = πg.UnboundLocal; _ = µsetUpModule
						var µe *πg.Object = πg.UnboundLocal; _ = µe
						var µerrorName *πg.Object = πg.UnboundLocal; _ = µerrorName
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 8: goto Label8
							case 9: goto Label9
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							// line 169: previousModule = self._get_previous_module(result)
							πF.SetLineno(169)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_get_previous_module, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µpreviousModule = πTemp003
							// line 170: currentModule = test.__class__.__module__
							πF.SetLineno(170)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtest, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__module__, nil); πE != nil {
								continue
							}
							µcurrentModule = πTemp003
							if πE = πg.CheckLocal(πF, µcurrentModule, "currentModule"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpreviousModule, "previousModule"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µcurrentModule, µpreviousModule); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 171: if currentModule == previousModule:
							πF.SetLineno(171)
						Label1:
							// line 172: return
							πF.SetLineno(172)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 174: self._handleModuleTearDown(result)
							πF.SetLineno(174)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_handleModuleTearDown, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 176: result._moduleSetUpFailed = False
							πF.SetLineno(176)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µresult, ß_moduleSetUpFailed, πTemp003); πE != nil {
								continue
							}
							// line 177: try:
							πF.SetLineno(177)
							πF.PushCheckpoint(4)
							// line 178: module = sys.modules[currentModule]
							πF.SetLineno(178)
							if πE = πg.CheckLocal(πF, µcurrentModule, "currentModule"); πE != nil {
								continue
							}
							πTemp002 = µcurrentModule
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßmodules, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							µmodule = πTemp003
							πF.PopCheckpoint()
							goto Label3
						Label4:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 179: except KeyError:
							πF.SetLineno(179)
						Label5:
							// line 180: return
							πF.SetLineno(180)
							πR = πg.None
							continue
							πF.RestoreExc(nil, nil)
							goto Label3
						Label3:
							// line 181: setUpModule = getattr(module, 'setUpModule', None)
							πF.SetLineno(181)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							πTemp001[0] = µmodule
							πTemp001[1] = ßsetUpModule.ToObject()
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
							µsetUpModule = πTemp003
							if πE = πg.CheckLocal(πF, µsetUpModule, "setUpModule"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µsetUpModule != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							goto Label7
							// line 182: if setUpModule is not None:
							πF.SetLineno(182)
						Label6:
							// line 183: _call_if_exists(result, '_setupStdout')
							πF.SetLineno(183)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							πTemp001[1] = ß_setupStdout.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_call_if_exists); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 184: try:
							πF.SetLineno(184)
							πF.PushCheckpoint(8)
							πF.PushCheckpoint(9)
							// line 185: setUpModule()
							πF.SetLineno(185)
							if πE = πg.CheckLocal(πF, µsetUpModule, "setUpModule"); πE != nil {
								continue
							}
							if πTemp002, πE = µsetUpModule.Call(πF, nil, nil); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							πF.PopCheckpoint()
							goto Label8
						Label9:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label10
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 186: except Exception, e:
							πF.SetLineno(186)
						Label10:
							µe = πTemp007.ToObject()
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_DebugResult); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
								goto Label11
							}
							goto Label12
							// line 187: if isinstance(result, _DebugResult):
							πF.SetLineno(187)
						Label11:
							// line 188: raise
							πF.SetLineno(188)
							πE = πF.Raise(nil, nil, nil)
							continue
							goto Label12
						Label12:
							// line 189: result._moduleSetUpFailed = True
							πF.SetLineno(189)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µresult, ß_moduleSetUpFailed, πTemp003); πE != nil {
								continue
							}
							// line 190: errorName = 'setUpModule (%s)' % currentModule
							πF.SetLineno(190)
							if πE = πg.CheckLocal(πF, µcurrentModule, "currentModule"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("setUpModule (%s)").ToObject(), µcurrentModule); πE != nil {
								continue
							}
							µerrorName = πTemp002
							// line 191: self._addClassOrModuleLevelException(result, e, errorName)
							πF.SetLineno(191)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							πTemp001[1] = µe
							if πE = πg.CheckLocal(πF, µerrorName, "errorName"); πE != nil {
								continue
							}
							πTemp001[2] = µerrorName
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_addClassOrModuleLevelException, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πF.RestoreExc(nil, nil)
							πF.PopCheckpoint()
							goto Label8
						Label8:
							πTemp007, πTemp008 = πF.RestoreExc(nil, nil)
							// line 193: _call_if_exists(result, '_restoreStdout')
							πF.SetLineno(193)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							πTemp001[1] = ß_restoreStdout.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_call_if_exists); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp007 != nil {
								πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
								continue
							}
							if πR != nil {
								continue
							}
							goto Label7
						Label7:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_handleModuleFixture.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 195: def _addClassOrModuleLevelException(self, result, exception, errorName):
					πF.SetLineno(195)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "result", Def: nil}
					πTemp002[2] = πg.Param{Name: "exception", Def: nil}
					πTemp002[3] = πg.Param{Name: "errorName", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("_addClassOrModuleLevelException", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µresult *πg.Object = πArgs[1]; _ = µresult
						var µexception *πg.Object = πArgs[2]; _ = µexception
						var µerrorName *πg.Object = πArgs[3]; _ = µerrorName
						var µerror *πg.Object = πg.UnboundLocal; _ = µerror
						var µaddSkip *πg.Object = πg.UnboundLocal; _ = µaddSkip
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
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
							// line 196: error = _ErrorHolder(errorName)
							πF.SetLineno(196)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerrorName, "errorName"); πE != nil {
								continue
							}
							πTemp001[0] = µerrorName
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_ErrorHolder); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µerror = πTemp003
							// line 197: addSkip = getattr(result, 'addSkip', None)
							πF.SetLineno(197)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							πTemp001[1] = ßaddSkip.ToObject()
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
							µaddSkip = πTemp003
							if πE = πg.CheckLocal(πF, µaddSkip, "addSkip"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µaddSkip != πTemp005).ToObject()
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label1
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µexception, "exception"); πE != nil {
								continue
							}
							πTemp001[0] = µexception
							if πTemp003, πE = πg.ResolveGlobal(πF, ßcase); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßSkipTest, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πTemp005
						Label1:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							goto Label3
							// line 198: if addSkip is not None and isinstance(exception, case.SkipTest):
							πF.SetLineno(198)
						Label2:
							// line 199: addSkip(error, str(exception))
							πF.SetLineno(199)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp001[0] = µerror
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µexception, "exception"); πE != nil {
								continue
							}
							πTemp006[0] = µexception
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µaddSkip, "addSkip"); πE != nil {
								continue
							}
							if πTemp002, πE = µaddSkip.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label4
						Label3:
							// line 201: result.addError(error, sys.exc_info())
							πF.SetLineno(201)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp001[0] = µerror
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßexc_info, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µresult, ßaddError, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label4
						Label4:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_addClassOrModuleLevelException.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 203: def _handleModuleTearDown(self, result):
					πF.SetLineno(203)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "result", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("_handleModuleTearDown", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µresult *πg.Object = πArgs[1]; _ = µresult
						var µpreviousModule *πg.Object = πg.UnboundLocal; _ = µpreviousModule
						var µmodule *πg.Object = πg.UnboundLocal; _ = µmodule
						var µtearDownModule *πg.Object = πg.UnboundLocal; _ = µtearDownModule
						var µe *πg.Object = πg.UnboundLocal; _ = µe
						var µerrorName *πg.Object = πg.UnboundLocal; _ = µerrorName
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 10: goto Label10
							case 11: goto Label11
							case 6: goto Label6
							default: panic("unexpected function state")
							}
							// line 204: previousModule = self._get_previous_module(result)
							πF.SetLineno(204)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_get_previous_module, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µpreviousModule = πTemp003
							if πE = πg.CheckLocal(πF, µpreviousModule, "previousModule"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µpreviousModule == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 205: if previousModule is None:
							πF.SetLineno(205)
						Label1:
							// line 206: return
							πF.SetLineno(206)
							πR = πg.None
							continue
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µresult, ß_moduleSetUpFailed, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 207: if result._moduleSetUpFailed:
							πF.SetLineno(207)
						Label3:
							// line 208: return
							πF.SetLineno(208)
							πR = πg.None
							continue
							goto Label4
						Label4:
							// line 210: try:
							πF.SetLineno(210)
							πF.PushCheckpoint(6)
							// line 211: module = sys.modules[previousModule]
							πF.SetLineno(211)
							if πE = πg.CheckLocal(πF, µpreviousModule, "previousModule"); πE != nil {
								continue
							}
							πTemp002 = µpreviousModule
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßmodules, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							µmodule = πTemp003
							πF.PopCheckpoint()
							goto Label5
						Label6:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 212: except KeyError:
							πF.SetLineno(212)
						Label7:
							// line 213: return
							πF.SetLineno(213)
							πR = πg.None
							continue
							πF.RestoreExc(nil, nil)
							goto Label5
						Label5:
							// line 215: tearDownModule = getattr(module, 'tearDownModule', None)
							πF.SetLineno(215)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							πTemp001[0] = µmodule
							πTemp001[1] = ßtearDownModule.ToObject()
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
							µtearDownModule = πTemp003
							if πE = πg.CheckLocal(πF, µtearDownModule, "tearDownModule"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µtearDownModule != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							goto Label9
							// line 216: if tearDownModule is not None:
							πF.SetLineno(216)
						Label8:
							// line 217: _call_if_exists(result, '_setupStdout')
							πF.SetLineno(217)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							πTemp001[1] = ß_setupStdout.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_call_if_exists); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 218: try:
							πF.SetLineno(218)
							πF.PushCheckpoint(10)
							πF.PushCheckpoint(11)
							// line 219: tearDownModule()
							πF.SetLineno(219)
							if πE = πg.CheckLocal(πF, µtearDownModule, "tearDownModule"); πE != nil {
								continue
							}
							if πTemp002, πE = µtearDownModule.Call(πF, nil, nil); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							πF.PopCheckpoint()
							goto Label10
						Label11:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label12
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 220: except Exception as e:
							πF.SetLineno(220)
						Label12:
							µe = πTemp007.ToObject()
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_DebugResult); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
								goto Label13
							}
							goto Label14
							// line 221: if isinstance(result, _DebugResult):
							πF.SetLineno(221)
						Label13:
							// line 222: raise
							πF.SetLineno(222)
							πE = πF.Raise(nil, nil, nil)
							continue
							goto Label14
						Label14:
							// line 223: errorName = 'tearDownModule (%s)' % previousModule
							πF.SetLineno(223)
							if πE = πg.CheckLocal(πF, µpreviousModule, "previousModule"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("tearDownModule (%s)").ToObject(), µpreviousModule); πE != nil {
								continue
							}
							µerrorName = πTemp002
							// line 224: self._addClassOrModuleLevelException(result, e, errorName)
							πF.SetLineno(224)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							πTemp001[1] = µe
							if πE = πg.CheckLocal(πF, µerrorName, "errorName"); πE != nil {
								continue
							}
							πTemp001[2] = µerrorName
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_addClassOrModuleLevelException, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πF.RestoreExc(nil, nil)
							πF.PopCheckpoint()
							goto Label10
						Label10:
							πTemp007, πTemp008 = πF.RestoreExc(nil, nil)
							// line 226: _call_if_exists(result, '_restoreStdout')
							πF.SetLineno(226)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							πTemp001[1] = ß_restoreStdout.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_call_if_exists); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp007 != nil {
								πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
								continue
							}
							if πR != nil {
								continue
							}
							goto Label9
						Label9:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_handleModuleTearDown.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 228: def _tearDownPreviousClass(self, test, result):
					πF.SetLineno(228)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "test", Def: nil}
					πTemp002[2] = πg.Param{Name: "result", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("_tearDownPreviousClass", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtest *πg.Object = πArgs[1]; _ = µtest
						var µresult *πg.Object = πArgs[2]; _ = µresult
						var µpreviousClass *πg.Object = πg.UnboundLocal; _ = µpreviousClass
						var µcurrentClass *πg.Object = πg.UnboundLocal; _ = µcurrentClass
						var µtearDownClass *πg.Object = πg.UnboundLocal; _ = µtearDownClass
						var µe *πg.Object = πg.UnboundLocal; _ = µe
						var µclassName *πg.Object = πg.UnboundLocal; _ = µclassName
						var µerrorName *πg.Object = πg.UnboundLocal; _ = µerrorName
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.BaseException
						_ = πTemp005
						var πTemp006 *πg.Traceback
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 11: goto Label11
							case 12: goto Label12
							default: panic("unexpected function state")
							}
							// line 229: previousClass = getattr(result, '_previousTestClass', None)
							πF.SetLineno(229)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							πTemp001[1] = ß_previousTestClass.ToObject()
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
							µpreviousClass = πTemp003
							// line 230: currentClass = test.__class__
							πF.SetLineno(230)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtest, ß__class__, nil); πE != nil {
								continue
							}
							µcurrentClass = πTemp002
							if πE = πg.CheckLocal(πF, µcurrentClass, "currentClass"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpreviousClass, "previousClass"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µcurrentClass, µpreviousClass); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 231: if currentClass == previousClass:
							πF.SetLineno(231)
						Label1:
							// line 232: return
							πF.SetLineno(232)
							πR = πg.None
							continue
							goto Label2
						Label2:
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µpreviousClass, "previousClass"); πE != nil {
								continue
							}
							πTemp001[0] = µpreviousClass
							πTemp001[1] = ß_classSetupFailed.ToObject()
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
								goto Label3
							}
							goto Label4
							// line 233: if getattr(previousClass, '_classSetupFailed', False):
							πF.SetLineno(233)
						Label3:
							// line 234: return
							πF.SetLineno(234)
							πR = πg.None
							continue
							goto Label4
						Label4:
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							πTemp001[1] = ß_moduleSetUpFailed.ToObject()
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
								goto Label5
							}
							goto Label6
							// line 235: if getattr(result, '_moduleSetUpFailed', False):
							πF.SetLineno(235)
						Label5:
							// line 236: return
							πF.SetLineno(236)
							πR = πg.None
							continue
							goto Label6
						Label6:
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µpreviousClass, "previousClass"); πE != nil {
								continue
							}
							πTemp001[0] = µpreviousClass
							πTemp001[1] = ß__unittest_skip__.ToObject()
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
								goto Label7
							}
							goto Label8
							// line 237: if getattr(previousClass, "__unittest_skip__", False):
							πF.SetLineno(237)
						Label7:
							// line 238: return
							πF.SetLineno(238)
							πR = πg.None
							continue
							goto Label8
						Label8:
							// line 240: tearDownClass = getattr(previousClass, 'tearDownClass', None)
							πF.SetLineno(240)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µpreviousClass, "previousClass"); πE != nil {
								continue
							}
							πTemp001[0] = µpreviousClass
							πTemp001[1] = ßtearDownClass.ToObject()
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
							µtearDownClass = πTemp003
							if πE = πg.CheckLocal(πF, µtearDownClass, "tearDownClass"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µtearDownClass != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label9
							}
							goto Label10
							// line 241: if tearDownClass is not None:
							πF.SetLineno(241)
						Label9:
							// line 242: _call_if_exists(result, '_setupStdout')
							πF.SetLineno(242)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							πTemp001[1] = ß_setupStdout.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_call_if_exists); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 243: try:
							πF.SetLineno(243)
							πF.PushCheckpoint(11)
							πF.PushCheckpoint(12)
							// line 244: tearDownClass()
							πF.SetLineno(244)
							if πE = πg.CheckLocal(πF, µtearDownClass, "tearDownClass"); πE != nil {
								continue
							}
							if πTemp002, πE = µtearDownClass.Call(πF, nil, nil); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							πF.PopCheckpoint()
							goto Label11
						Label12:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label13
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 245: except Exception, e:
							πF.SetLineno(245)
						Label13:
							µe = πTemp005.ToObject()
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_DebugResult); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							goto Label15
							// line 246: if isinstance(result, _DebugResult):
							πF.SetLineno(246)
						Label14:
							// line 247: raise
							πF.SetLineno(247)
							πE = πF.Raise(nil, nil, nil)
							continue
							goto Label15
						Label15:
							// line 248: className = util.strclass(previousClass)
							πF.SetLineno(248)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpreviousClass, "previousClass"); πE != nil {
								continue
							}
							πTemp001[0] = µpreviousClass
							if πTemp002, πE = πg.ResolveGlobal(πF, ßutil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstrclass, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µclassName = πTemp002
							// line 249: errorName = 'tearDownClass (%s)' % className
							πF.SetLineno(249)
							if πE = πg.CheckLocal(πF, µclassName, "className"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("tearDownClass (%s)").ToObject(), µclassName); πE != nil {
								continue
							}
							µerrorName = πTemp002
							// line 250: self._addClassOrModuleLevelException(result, e, errorName)
							πF.SetLineno(250)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							πTemp001[1] = µe
							if πE = πg.CheckLocal(πF, µerrorName, "errorName"); πE != nil {
								continue
							}
							πTemp001[2] = µerrorName
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_addClassOrModuleLevelException, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πF.RestoreExc(nil, nil)
							πF.PopCheckpoint()
							goto Label11
						Label11:
							πTemp005, πTemp006 = πF.RestoreExc(nil, nil)
							// line 252: _call_if_exists(result, '_restoreStdout')
							πF.SetLineno(252)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							πTemp001[1] = ß_restoreStdout.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_call_if_exists); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005 != nil {
								πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
								continue
							}
							if πR != nil {
								continue
							}
							goto Label10
						Label10:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_tearDownPreviousClass.ToObject(), πTemp009); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("TestSuite").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTestSuite.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 255: class _ErrorHolder(object):
			πF.SetLineno(255)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp004 = πg.NewDict()
			if πTemp005, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp005); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_ErrorHolder", "build/src/__python__/unittest_suite.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 256: """
					πF.SetLineno(256)
					// line 265: failureException = None
					πF.SetLineno(265)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßfailureException.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 267: def __init__(self, description):
					πF.SetLineno(267)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "description", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdescription *πg.Object = πArgs[1]; _ = µdescription
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 268: self.description = description
							πF.SetLineno(268)
							if πE = πg.CheckLocal(πF, µdescription, "description"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdescription); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdescription, πTemp001); πE != nil {
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
					// line 270: def id(self):
					πF.SetLineno(270)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("id", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 271: return self.description
							πF.SetLineno(271)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdescription, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßid.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 273: def shortDescription(self):
					πF.SetLineno(273)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("shortDescription", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 274: return None
							πF.SetLineno(274)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
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
					if πE = πClass.SetItem(πF, ßshortDescription.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 276: def __repr__(self):
					πF.SetLineno(276)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("__repr__", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
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
							// line 277: return "<ErrorHolder description=%r>" % (self.description,)
							πF.SetLineno(277)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdescription, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple1(πTemp003).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<ErrorHolder description=%r>").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 279: def __str__(self):
					πF.SetLineno(279)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("__str__", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 280: return self.id()
							πF.SetLineno(280)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßid, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ß__str__.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 282: def run(self, result):
					πF.SetLineno(282)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "result", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("run", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µresult *πg.Object = πArgs[1]; _ = µresult
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 285: pass
							πF.SetLineno(285)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 287: def __call__(self, result):
					πF.SetLineno(287)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "result", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("__call__", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µresult *πg.Object = πArgs[1]; _ = µresult
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
							// line 288: return self.run(result)
							πF.SetLineno(288)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrun, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__call__.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 290: def countTestCases(self):
					πF.SetLineno(290)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("countTestCases", "build/src/__python__/unittest_suite.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 291: return 0
							πF.SetLineno(291)
							πR = πg.NewInt(0).ToObject()
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcountTestCases.ToObject(), πTemp009); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("_ErrorHolder").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_ErrorHolder.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 293: def _isnotsuite(test):
			πF.SetLineno(293)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "test", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("_isnotsuite", "build/src/__python__/unittest_suite.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtest *πg.Object = πArgs[0]; _ = µtest
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.BaseException
				_ = πTemp004
				var πTemp005 *πg.Traceback
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 294: "A crude way to tell apart testcases and suites with duck-typing"
					πF.SetLineno(294)
					// line 295: try:
					πF.SetLineno(295)
					πF.PushCheckpoint(2)
					// line 296: iter(test)
					πF.SetLineno(296)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
						continue
					}
					πTemp001[0] = µtest
					if πTemp002, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 297: except TypeError:
					πF.SetLineno(297)
				Label3:
					// line 298: return True
					πF.SetLineno(298)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 299: return False
					πF.SetLineno(299)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ß_isnotsuite.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 302: class _DebugResult(object):
			πF.SetLineno(302)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp004 = πg.NewDict()
			if πTemp006, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp006); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_DebugResult", "build/src/__python__/unittest_suite.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 303: "Used by the TestSuite to hold previous class when running in debug."
					πF.SetLineno(303)
					// line 304: _previousTestClass = None
					πF.SetLineno(304)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß_previousTestClass.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 305: _moduleSetUpFailed = False
					πF.SetLineno(305)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß_moduleSetUpFailed.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 306: shouldStop = False
					πF.SetLineno(306)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßshouldStop.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp007, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp007 == nil {
				πTemp007 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp007.Call(πF, []*πg.Object{πg.NewStr("_DebugResult").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_DebugResult.ToObject(), πTemp008); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("unittest_suite", Code)
}
