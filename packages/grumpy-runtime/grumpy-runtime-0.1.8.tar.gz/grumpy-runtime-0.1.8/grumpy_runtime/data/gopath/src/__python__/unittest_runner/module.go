package unittest_runner
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/unittest_runner.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßAttributeError := πg.InternStr("AttributeError")
		ßE := πg.InternStr("E")
		ßERROR := πg.InternStr("ERROR")
		ßF := πg.InternStr("F")
		ßFAIL := πg.InternStr("FAIL")
		ßFAILED := πg.InternStr("FAILED")
		ßFalse := πg.InternStr("False")
		ßNone := πg.InternStr("None")
		ßOK := πg.InternStr("OK")
		ßTestResult := πg.InternStr("TestResult")
		ßTextTestResult := πg.InternStr("TextTestResult")
		ßTextTestRunner := πg.InternStr("TextTestRunner")
		ßTrue := πg.InternStr("True")
		ß_WritelnDecorator := πg.InternStr("_WritelnDecorator")
		ß__getattr__ := πg.InternStr("__getattr__")
		ß__getstate__ := πg.InternStr("__getstate__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__unittest := πg.InternStr("__unittest")
		ß_makeResult := πg.InternStr("_makeResult")
		ßaddError := πg.InternStr("addError")
		ßaddExpectedFailure := πg.InternStr("addExpectedFailure")
		ßaddFailure := πg.InternStr("addFailure")
		ßaddSkip := πg.InternStr("addSkip")
		ßaddSuccess := πg.InternStr("addSuccess")
		ßaddUnexpectedSuccess := πg.InternStr("addUnexpectedSuccess")
		ßappend := πg.InternStr("append")
		ßbuffer := πg.InternStr("buffer")
		ßdescriptions := πg.InternStr("descriptions")
		ßdots := πg.InternStr("dots")
		ßerrors := πg.InternStr("errors")
		ßexpectedFailures := πg.InternStr("expectedFailures")
		ßfailfast := πg.InternStr("failfast")
		ßfailures := πg.InternStr("failures")
		ßflush := πg.InternStr("flush")
		ßgetDescription := πg.InternStr("getDescription")
		ßgetattr := πg.InternStr("getattr")
		ßhasattr := πg.InternStr("hasattr")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßmap := πg.InternStr("map")
		ßobject := πg.InternStr("object")
		ßok := πg.InternStr("ok")
		ßprintErrorList := πg.InternStr("printErrorList")
		ßprintErrors := πg.InternStr("printErrors")
		ßregisterResult := πg.InternStr("registerResult")
		ßresult := πg.InternStr("result")
		ßresultclass := πg.InternStr("resultclass")
		ßrun := πg.InternStr("run")
		ßs := πg.InternStr("s")
		ßseparator1 := πg.InternStr("separator1")
		ßseparator2 := πg.InternStr("separator2")
		ßshortDescription := πg.InternStr("shortDescription")
		ßshowAll := πg.InternStr("showAll")
		ßskipped := πg.InternStr("skipped")
		ßstartTest := πg.InternStr("startTest")
		ßstartTestRun := πg.InternStr("startTestRun")
		ßstderr := πg.InternStr("stderr")
		ßstopTestRun := πg.InternStr("stopTestRun")
		ßstr := πg.InternStr("str")
		ßstream := πg.InternStr("stream")
		ßsuper := πg.InternStr("super")
		ßsys := πg.InternStr("sys")
		ßtestsRun := πg.InternStr("testsRun")
		ßtime := πg.InternStr("time")
		ßu := πg.InternStr("u")
		ßunexpectedSuccesses := πg.InternStr("unexpectedSuccesses")
		ßunittest_signals := πg.InternStr("unittest_signals")
		ßverbosity := πg.InternStr("verbosity")
		ßwasSuccessful := πg.InternStr("wasSuccessful")
		ßwrite := πg.InternStr("write")
		ßwriteln := πg.InternStr("writeln")
		ßx := πg.InternStr("x")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.Dict
		_ = πTemp004
		var πTemp005 *πg.Object
		_ = πTemp005
		var πTemp006 *πg.Object
		_ = πTemp006
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """Running tests"""
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
			// line 4: import time
			πF.SetLineno(4)
			if πTemp002, πE = πg.ImportModule(πF, "time"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtime.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 8: import unittest_result as result
			πF.SetLineno(8)
			if πTemp002, πE = πg.ImportModule(πF, "unittest_result"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßresult.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 9: import unittest_signals
			πF.SetLineno(9)
			if πTemp002, πE = πg.ImportModule(πF, "unittest_signals"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunittest_signals.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 10: registerResult = unittest_signals.registerResult
			πF.SetLineno(10)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßunittest_signals); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßregisterResult, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßregisterResult.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 12: __unittest = True
			πF.SetLineno(12)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__unittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 15: class _WritelnDecorator(object):
			πF.SetLineno(15)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_WritelnDecorator", "build/src/__python__/unittest_runner.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 16: """Used to decorate file-like objects with a handy 'writeln' method"""
					πF.SetLineno(16)
					// line 17: def __init__(self, stream):
					πF.SetLineno(17)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "stream", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/unittest_runner.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstream *πg.Object = πArgs[1]; _ = µstream
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 18: self.stream = stream
							πF.SetLineno(18)
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µstream); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstream, πTemp001); πE != nil {
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
					// line 20: def __getattr__(self, attr):
					πF.SetLineno(20)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "attr", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__getattr__", "build/src/__python__/unittest_runner.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µattr *πg.Object = πArgs[1]; _ = µattr
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(ßstream.ToObject(), ß__getstate__.ToObject()).ToObject()
							if πTemp003, πE = πg.Contains(πF, πTemp002, µattr); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 21: if attr in ('stream', '__getstate__'):
							πF.SetLineno(21)
						Label1:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp004[0] = µattr
							if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 22: raise AttributeError(attr)
							πF.SetLineno(22)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label2
						Label2:
							// line 23: return getattr(self.stream,attr)
							πF.SetLineno(23)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp004[1] = µattr
							if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ß__getattr__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 25: def writeln(self, arg=None):
					πF.SetLineno(25)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "arg", Def: πTemp005}
					πTemp004 = πg.NewFunction(πg.NewCode("writeln", "build/src/__python__/unittest_runner.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µarg *πg.Object = πArgs[1]; _ = µarg
						var πTemp001 bool
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
							if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IsTrue(πF, µarg); πE != nil {
								continue
							}
							if πTemp001 {
								goto Label1
							}
							goto Label2
							// line 26: if arg:
							πF.SetLineno(26)
						Label1:
							// line 27: self.write(arg)
							πF.SetLineno(27)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
								continue
							}
							πTemp002[0] = µarg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label2
						Label2:
							// line 28: self.write('\n') # text-mode streams translate to \r\n if needed
							πF.SetLineno(28)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßwriteln.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 30: def write(self, arg):
					πF.SetLineno(30)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "arg", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("write", "build/src/__python__/unittest_runner.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µarg *πg.Object = πArgs[1]; _ = µarg
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
							// line 31: self.stream.write(arg)
							πF.SetLineno(31)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
								continue
							}
							πTemp001[0] = µarg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßwrite, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßwrite.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 33: def flush(self):
					πF.SetLineno(33)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("flush", "build/src/__python__/unittest_runner.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 34: pass
							πF.SetLineno(34)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßflush.ToObject(), πTemp006); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("_WritelnDecorator").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_WritelnDecorator.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 36: class TextTestResult(result.TestResult):
			πF.SetLineno(36)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßresult); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßTestResult, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TextTestResult", "build/src/__python__/unittest_runner.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 37: """A test result class that can print formatted text results to a stream.
					πF.SetLineno(37)
					// line 41: separator1 = '=' * 70
					πF.SetLineno(41)
					if πTemp001, πE = πg.Mul(πF, πg.NewStr("=").ToObject(), πg.NewInt(70).ToObject()); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßseparator1.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 42: separator2 = '-' * 70
					πF.SetLineno(42)
					if πTemp001, πE = πg.Mul(πF, πg.NewStr("-").ToObject(), πg.NewInt(70).ToObject()); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßseparator2.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 44: def __init__(self, stream, descriptions, verbosity):
					πF.SetLineno(44)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "stream", Def: nil}
					πTemp002[2] = πg.Param{Name: "descriptions", Def: nil}
					πTemp002[3] = πg.Param{Name: "verbosity", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/unittest_runner.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstream *πg.Object = πArgs[1]; _ = µstream
						var µdescriptions *πg.Object = πArgs[2]; _ = µdescriptions
						var µverbosity *πg.Object = πArgs[3]; _ = µverbosity
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
							// line 45: super(TextTestResult, self).__init__(stream, descriptions, verbosity)
							πF.SetLineno(45)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							πTemp001[0] = µstream
							if πE = πg.CheckLocal(πF, µdescriptions, "descriptions"); πE != nil {
								continue
							}
							πTemp001[1] = µdescriptions
							if πE = πg.CheckLocal(πF, µverbosity, "verbosity"); πE != nil {
								continue
							}
							πTemp001[2] = µverbosity
							πTemp002 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTextTestResult); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[1] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 46: self.stream = stream
							πF.SetLineno(46)
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µstream); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstream, πTemp003); πE != nil {
								continue
							}
							// line 47: self.showAll = verbosity > 1
							πF.SetLineno(47)
							if πE = πg.CheckLocal(πF, µverbosity, "verbosity"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, µverbosity, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßshowAll, πTemp004); πE != nil {
								continue
							}
							// line 48: self.dots = verbosity == 1
							πF.SetLineno(48)
							if πE = πg.CheckLocal(πF, µverbosity, "verbosity"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, µverbosity, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdots, πTemp004); πE != nil {
								continue
							}
							// line 49: self.descriptions = descriptions
							πF.SetLineno(49)
							if πE = πg.CheckLocal(πF, µdescriptions, "descriptions"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µdescriptions); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdescriptions, πTemp003); πE != nil {
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
					// line 51: def getDescription(self, test):
					πF.SetLineno(51)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "test", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("getDescription", "build/src/__python__/unittest_runner.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtest *πg.Object = πArgs[1]; _ = µtest
						var µdoc_first_line *πg.Object = πg.UnboundLocal; _ = µdoc_first_line
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
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
							// line 52: doc_first_line = test.shortDescription()
							πF.SetLineno(52)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtest, ßshortDescription, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µdoc_first_line = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdescriptions, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µdoc_first_line, "doc_first_line"); πE != nil {
								continue
							}
							πTemp001 = µdoc_first_line
						Label1:
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label2
							}
							goto Label3
							// line 53: if self.descriptions and doc_first_line:
							πF.SetLineno(53)
						Label2:
							// line 54: return '\n'.join((str(test), doc_first_line))
							πF.SetLineno(54)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp005[0] = µtest
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.CheckLocal(πF, µdoc_first_line, "doc_first_line"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp006, µdoc_first_line).ToObject()
							πTemp004[0] = πTemp001
							if πTemp001, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πR = πTemp002
							continue
							goto Label4
						Label3:
							// line 56: return str(test)
							πF.SetLineno(56)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp004[0] = µtest
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πR = πTemp002
							continue
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
					if πE = πClass.SetItem(πF, ßgetDescription.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 58: def startTest(self, test):
					πF.SetLineno(58)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "test", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("startTest", "build/src/__python__/unittest_runner.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtest *πg.Object = πArgs[1]; _ = µtest
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 59: super(TextTestResult, self).startTest(test)
							πF.SetLineno(59)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp001[0] = µtest
							πTemp002 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTextTestResult); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[1] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßstartTest, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßshowAll, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 60: if self.showAll:
							πF.SetLineno(60)
						Label1:
							// line 61: self.stream.write(self.getDescription(test))
							πF.SetLineno(61)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp002[0] = µtest
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßgetDescription, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 62: self.stream.write(" ... ")
							πF.SetLineno(62)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr(" ... ").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 63: self.stream.flush()
							πF.SetLineno(63)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßflush, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßstartTest.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 65: def addSuccess(self, test):
					πF.SetLineno(65)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "test", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("addSuccess", "build/src/__python__/unittest_runner.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtest *πg.Object = πArgs[1]; _ = µtest
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 66: super(TextTestResult, self).addSuccess(test)
							πF.SetLineno(66)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp001[0] = µtest
							πTemp002 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTextTestResult); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[1] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßaddSuccess, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßshowAll, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdots, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label2
							}
							goto Label3
							// line 67: if self.showAll:
							πF.SetLineno(67)
						Label1:
							// line 68: self.stream.writeln("ok")
							πF.SetLineno(68)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßok.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßwriteln, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label3
							// line 69: elif self.dots:
							πF.SetLineno(69)
						Label2:
							// line 70: self.stream.write('.')
							πF.SetLineno(70)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr(".").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 71: self.stream.flush()
							πF.SetLineno(71)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßflush, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßaddSuccess.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 73: def addError(self, test, err):
					πF.SetLineno(73)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "test", Def: nil}
					πTemp002[2] = πg.Param{Name: "err", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("addError", "build/src/__python__/unittest_runner.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtest *πg.Object = πArgs[1]; _ = µtest
						var µerr *πg.Object = πArgs[2]; _ = µerr
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 74: super(TextTestResult, self).addError(test, err)
							πF.SetLineno(74)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp001[0] = µtest
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp001[1] = µerr
							πTemp002 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTextTestResult); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[1] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßaddError, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßshowAll, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdots, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label2
							}
							goto Label3
							// line 75: if self.showAll:
							πF.SetLineno(75)
						Label1:
							// line 76: self.stream.writeln("ERROR")
							πF.SetLineno(76)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßERROR.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßwriteln, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label3
							// line 77: elif self.dots:
							πF.SetLineno(77)
						Label2:
							// line 78: self.stream.write('E')
							πF.SetLineno(78)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßE.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 79: self.stream.flush()
							πF.SetLineno(79)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßflush, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßaddError.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 81: def addFailure(self, test, err):
					πF.SetLineno(81)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "test", Def: nil}
					πTemp002[2] = πg.Param{Name: "err", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("addFailure", "build/src/__python__/unittest_runner.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtest *πg.Object = πArgs[1]; _ = µtest
						var µerr *πg.Object = πArgs[2]; _ = µerr
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 82: super(TextTestResult, self).addFailure(test, err)
							πF.SetLineno(82)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp001[0] = µtest
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp001[1] = µerr
							πTemp002 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTextTestResult); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[1] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßaddFailure, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßshowAll, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdots, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label2
							}
							goto Label3
							// line 83: if self.showAll:
							πF.SetLineno(83)
						Label1:
							// line 84: self.stream.writeln("FAIL")
							πF.SetLineno(84)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßFAIL.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßwriteln, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label3
							// line 85: elif self.dots:
							πF.SetLineno(85)
						Label2:
							// line 86: self.stream.write('F')
							πF.SetLineno(86)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßF.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 87: self.stream.flush()
							πF.SetLineno(87)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßflush, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßaddFailure.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 89: def addSkip(self, test, reason):
					πF.SetLineno(89)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "test", Def: nil}
					πTemp002[2] = πg.Param{Name: "reason", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("addSkip", "build/src/__python__/unittest_runner.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtest *πg.Object = πArgs[1]; _ = µtest
						var µreason *πg.Object = πArgs[2]; _ = µreason
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 90: super(TextTestResult, self).addSkip(test, reason)
							πF.SetLineno(90)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp001[0] = µtest
							if πE = πg.CheckLocal(πF, µreason, "reason"); πE != nil {
								continue
							}
							πTemp001[1] = µreason
							πTemp002 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTextTestResult); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[1] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßaddSkip, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßshowAll, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdots, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label2
							}
							goto Label3
							// line 91: if self.showAll:
							πF.SetLineno(91)
						Label1:
							// line 93: self.stream.writeln("skipped %r" % (reason))
							πF.SetLineno(93)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µreason, "reason"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("skipped %r").ToObject(), µreason); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßwriteln, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label3
							// line 94: elif self.dots:
							πF.SetLineno(94)
						Label2:
							// line 95: self.stream.write("s")
							πF.SetLineno(95)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßs.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 96: self.stream.flush()
							πF.SetLineno(96)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßflush, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßaddSkip.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 98: def addExpectedFailure(self, test, err):
					πF.SetLineno(98)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "test", Def: nil}
					πTemp002[2] = πg.Param{Name: "err", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("addExpectedFailure", "build/src/__python__/unittest_runner.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtest *πg.Object = πArgs[1]; _ = µtest
						var µerr *πg.Object = πArgs[2]; _ = µerr
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 99: super(TextTestResult, self).addExpectedFailure(test, err)
							πF.SetLineno(99)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp001[0] = µtest
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp001[1] = µerr
							πTemp002 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTextTestResult); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[1] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßaddExpectedFailure, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßshowAll, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdots, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label2
							}
							goto Label3
							// line 100: if self.showAll:
							πF.SetLineno(100)
						Label1:
							// line 101: self.stream.writeln("expected failure")
							πF.SetLineno(101)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("expected failure").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßwriteln, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label3
							// line 102: elif self.dots:
							πF.SetLineno(102)
						Label2:
							// line 103: self.stream.write("x")
							πF.SetLineno(103)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßx.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 104: self.stream.flush()
							πF.SetLineno(104)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßflush, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßaddExpectedFailure.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 106: def addUnexpectedSuccess(self, test):
					πF.SetLineno(106)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "test", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("addUnexpectedSuccess", "build/src/__python__/unittest_runner.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtest *πg.Object = πArgs[1]; _ = µtest
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 107: super(TextTestResult, self).addUnexpectedSuccess(test)
							πF.SetLineno(107)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp001[0] = µtest
							πTemp002 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTextTestResult); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[1] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßaddUnexpectedSuccess, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßshowAll, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdots, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label2
							}
							goto Label3
							// line 108: if self.showAll:
							πF.SetLineno(108)
						Label1:
							// line 109: self.stream.writeln("unexpected success")
							πF.SetLineno(109)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("unexpected success").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßwriteln, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label3
							// line 110: elif self.dots:
							πF.SetLineno(110)
						Label2:
							// line 111: self.stream.write("u")
							πF.SetLineno(111)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßu.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 112: self.stream.flush()
							πF.SetLineno(112)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßflush, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßaddUnexpectedSuccess.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 114: def printErrors(self):
					πF.SetLineno(114)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("printErrors", "build/src/__python__/unittest_runner.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdots, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßshowAll, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 115: if self.dots or self.showAll:
							πF.SetLineno(115)
						Label2:
							// line 116: self.stream.writeln()
							πF.SetLineno(116)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßwriteln, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label3
						Label3:
							// line 117: self.printErrorList('ERROR', self.errors)
							πF.SetLineno(117)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = ßERROR.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßerrors, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßprintErrorList, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 118: self.printErrorList('FAIL', self.failures)
							πF.SetLineno(118)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = ßFAIL.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfailures, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßprintErrorList, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßprintErrors.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 120: def printErrorList(self, flavour, errors):
					πF.SetLineno(120)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "flavour", Def: nil}
					πTemp002[2] = πg.Param{Name: "errors", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("printErrorList", "build/src/__python__/unittest_runner.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µflavour *πg.Object = πArgs[1]; _ = µflavour
						var µerrors *πg.Object = πArgs[2]; _ = µerrors
						var µtest *πg.Object = πg.UnboundLocal; _ = µtest
						var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
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
							if πE = πg.CheckLocal(πF, µerrors, "errors"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µerrors); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp004); πE != nil {
									continue
								}
								µtest = πTemp005
								µerr = πTemp006
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 122: self.stream.writeln(self.separator1)
							πF.SetLineno(122)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßseparator1, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßwriteln, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 123: self.stream.writeln("%s: %s" % (flavour,self.getDescription(test)))
							πF.SetLineno(123)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µflavour, "flavour"); πE != nil {
								continue
							}
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							πTemp008[0] = µtest
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßgetDescription, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp006.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp005 = πg.NewTuple2(µflavour, πTemp009).ToObject()
							if πTemp004, πE = πg.Mod(πF, πg.NewStr("%s: %s").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp007[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßwriteln, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 124: self.stream.writeln(self.separator2)
							πF.SetLineno(124)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßseparator2, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßwriteln, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 125: self.stream.writeln("%s" % err)
							πF.SetLineno(125)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mod(πF, πg.NewStr("%s").ToObject(), µerr); πE != nil {
								continue
							}
							πTemp007[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßwriteln, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
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
					if πE = πClass.SetItem(πF, ßprintErrorList.ToObject(), πTemp012); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("TextTestResult").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTextTestResult.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 128: class TextTestRunner(object):
			πF.SetLineno(128)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TextTestRunner", "build/src/__python__/unittest_runner.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 129: """A test runner class that displays results in textual form.
					πF.SetLineno(129)
					// line 134: resultclass = TextTestResult
					πF.SetLineno(134)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTextTestResult); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßresultclass.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 136: def __init__(self, stream=sys.stderr, descriptions=True, verbosity=1,
					πF.SetLineno(136)
					πTemp002 = make([]πg.Param, 7)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßstderr, nil); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "stream", Def: πTemp004}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "descriptions", Def: πTemp003}
					πTemp002[3] = πg.Param{Name: "verbosity", Def: πg.NewInt(1).ToObject()}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "failfast", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "buffer", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[6] = πg.Param{Name: "resultclass", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/unittest_runner.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstream *πg.Object = πArgs[1]; _ = µstream
						var µdescriptions *πg.Object = πArgs[2]; _ = µdescriptions
						var µverbosity *πg.Object = πArgs[3]; _ = µverbosity
						var µfailfast *πg.Object = πArgs[4]; _ = µfailfast
						var µbuffer *πg.Object = πArgs[5]; _ = µbuffer
						var µresultclass *πg.Object = πArgs[6]; _ = µresultclass
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
							// line 138: self.stream = _WritelnDecorator(stream)
							πF.SetLineno(138)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							πTemp001[0] = µstream
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_WritelnDecorator); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstream, πTemp002); πE != nil {
								continue
							}
							// line 139: self.descriptions = descriptions
							πF.SetLineno(139)
							if πE = πg.CheckLocal(πF, µdescriptions, "descriptions"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µdescriptions); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdescriptions, πTemp002); πE != nil {
								continue
							}
							// line 140: self.verbosity = verbosity
							πF.SetLineno(140)
							if πE = πg.CheckLocal(πF, µverbosity, "verbosity"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µverbosity); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßverbosity, πTemp002); πE != nil {
								continue
							}
							// line 141: self.failfast = failfast
							πF.SetLineno(141)
							if πE = πg.CheckLocal(πF, µfailfast, "failfast"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µfailfast); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfailfast, πTemp002); πE != nil {
								continue
							}
							// line 142: self.buffer = buffer
							πF.SetLineno(142)
							if πE = πg.CheckLocal(πF, µbuffer, "buffer"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µbuffer); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbuffer, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µresultclass, "resultclass"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µresultclass != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 143: if resultclass is not None:
							πF.SetLineno(143)
						Label1:
							// line 144: self.resultclass = resultclass
							πF.SetLineno(144)
							if πE = πg.CheckLocal(πF, µresultclass, "resultclass"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µresultclass); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßresultclass, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 146: def _makeResult(self):
					πF.SetLineno(146)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("_makeResult", "build/src/__python__/unittest_runner.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 147: return self.resultclass(self.stream, self.descriptions, self.verbosity)
							πF.SetLineno(147)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdescriptions, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßverbosity, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßresultclass, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_makeResult.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 149: def run(self, test):
					πF.SetLineno(149)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "test", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("run", "build/src/__python__/unittest_runner.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtest *πg.Object = πArgs[1]; _ = µtest
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
						var µstartTime *πg.Object = πg.UnboundLocal; _ = µstartTime
						var µstartTestRun *πg.Object = πg.UnboundLocal; _ = µstartTestRun
						var µstopTestRun *πg.Object = πg.UnboundLocal; _ = µstopTestRun
						var µstopTime *πg.Object = πg.UnboundLocal; _ = µstopTime
						var µtimeTaken *πg.Object = πg.UnboundLocal; _ = µtimeTaken
						var µrun *πg.Object = πg.UnboundLocal; _ = µrun
						var µexpectedFails *πg.Object = πg.UnboundLocal; _ = µexpectedFails
						var µunexpectedSuccesses *πg.Object = πg.UnboundLocal; _ = µunexpectedSuccesses
						var µskipped *πg.Object = πg.UnboundLocal; _ = µskipped
						var µresults *πg.Object = πg.UnboundLocal; _ = µresults
						var µinfos *πg.Object = πg.UnboundLocal; _ = µinfos
						var µfailed *πg.Object = πg.UnboundLocal; _ = µfailed
						var µerrored *πg.Object = πg.UnboundLocal; _ = µerrored
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.BaseException
						_ = πTemp005
						var πTemp006 *πg.Traceback
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 []*πg.Object
						_ = πTemp011
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 11: goto Label11
							case 3: goto Label3
							default: panic("unexpected function state")
							}
							// line 150: "Run the given test case or test suite."
							πF.SetLineno(150)
							// line 151: result = self._makeResult()
							πF.SetLineno(151)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_makeResult, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µresult = πTemp002
							// line 152: registerResult(result)
							πF.SetLineno(152)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp003[0] = µresult
							if πTemp001, πE = πg.ResolveGlobal(πF, ßregisterResult); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 153: result.failfast = self.failfast
							πF.SetLineno(153)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfailfast, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µresult, ßfailfast, πTemp002); πE != nil {
								continue
							}
							// line 154: result.buffer = self.buffer
							πF.SetLineno(154)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbuffer, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µresult, ßbuffer, πTemp002); πE != nil {
								continue
							}
							// line 155: startTime = time.time()
							πF.SetLineno(155)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtime, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µstartTime = πTemp001
							// line 156: startTestRun = getattr(result, 'startTestRun', None)
							πF.SetLineno(156)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp003[0] = µresult
							πTemp003[1] = ßstartTestRun.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µstartTestRun = πTemp002
							if πE = πg.CheckLocal(πF, µstartTestRun, "startTestRun"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µstartTestRun != πTemp002).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 157: if startTestRun is not None:
							πF.SetLineno(157)
						Label1:
							// line 158: startTestRun()
							πF.SetLineno(158)
							if πE = πg.CheckLocal(πF, µstartTestRun, "startTestRun"); πE != nil {
								continue
							}
							if πTemp001, πE = µstartTestRun.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 159: try:
							πF.SetLineno(159)
							πF.PushCheckpoint(3)
							// line 160: test(result)
							πF.SetLineno(160)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp003[0] = µresult
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							if πTemp001, πE = µtest.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πF.PopCheckpoint()
							goto Label3
						Label3:
							πTemp005, πTemp006 = πF.RestoreExc(nil, nil)
							// line 162: stopTestRun = getattr(result, 'stopTestRun', None)
							πF.SetLineno(162)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp003[0] = µresult
							πTemp003[1] = ßstopTestRun.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µstopTestRun = πTemp002
							if πE = πg.CheckLocal(πF, µstopTestRun, "stopTestRun"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µstopTestRun != πTemp002).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 163: if stopTestRun is not None:
							πF.SetLineno(163)
						Label4:
							// line 164: stopTestRun()
							πF.SetLineno(164)
							if πE = πg.CheckLocal(πF, µstopTestRun, "stopTestRun"); πE != nil {
								continue
							}
							if πTemp001, πE = µstopTestRun.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label5
						Label5:
							if πTemp005 != nil {
								πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
								continue
							}
							if πR != nil {
								continue
							}
							// line 165: stopTime = time.time()
							πF.SetLineno(165)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtime, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µstopTime = πTemp001
							// line 166: timeTaken = stopTime - startTime
							πF.SetLineno(166)
							if πE = πg.CheckLocal(πF, µstopTime, "stopTime"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstartTime, "startTime"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µstopTime, µstartTime); πE != nil {
								continue
							}
							µtimeTaken = πTemp001
							// line 167: result.printErrors()
							πF.SetLineno(167)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µresult, ßprintErrors, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp003[0] = µresult
							πTemp003[1] = ßseparator2.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							goto Label7
							// line 168: if hasattr(result, 'separator2'):
							πF.SetLineno(168)
						Label6:
							// line 169: self.stream.writeln(result.separator2)
							πF.SetLineno(169)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µresult, ßseparator2, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßwriteln, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label7
						Label7:
							// line 170: run = result.testsRun
							πF.SetLineno(170)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µresult, ßtestsRun, nil); πE != nil {
								continue
							}
							µrun = πTemp001
							// line 172: self.stream.writeln("Ran %d test%s in %fs" %
							πF.SetLineno(172)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrun, "run"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrun, "run"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.NE(πF, µrun, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp008 = πTemp010
							if πTemp009, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							if !πTemp009 {
								goto Label9
							}
							πTemp008 = ßs.ToObject()
						Label9:
							πTemp007 = πTemp008
							if πTemp004, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							πTemp007 = ß.ToObject()
						Label8:
							if πE = πg.CheckLocal(πF, µtimeTaken, "timeTaken"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(µrun, πTemp007, µtimeTaken).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Ran %d test%s in %fs").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßwriteln, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 174: self.stream.writeln()
							πF.SetLineno(174)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßwriteln, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 176: expectedFails = unexpectedSuccesses = skipped = 0
							πF.SetLineno(176)
							µexpectedFails = πg.NewInt(0).ToObject()
							µunexpectedSuccesses = πg.NewInt(0).ToObject()
							µskipped = πg.NewInt(0).ToObject()
							// line 177: try:
							πF.SetLineno(177)
							πF.PushCheckpoint(11)
							// line 178: results = map(len, (result.expectedFailures,
							πF.SetLineno(178)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µresult, ßexpectedFailures, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µresult, ßunexpectedSuccesses, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µresult, ßskipped, nil); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(πTemp002, πTemp007, πTemp008).ToObject()
							πTemp003[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µresults = πTemp002
							πF.PopCheckpoint()
							// line 184: expectedFails, unexpectedSuccesses, skipped = results
							πF.SetLineno(184)
							if πE = πg.CheckLocal(πF, µresults, "results"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp007}}}, µresults); πE != nil {
								continue
							}
							µexpectedFails = πTemp001
							µunexpectedSuccesses = πTemp002
							µskipped = πTemp007
							goto Label10
						Label11:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label12
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 181: except AttributeError:
							πF.SetLineno(181)
						Label12:
							// line 182: pass
							πF.SetLineno(182)
							πF.RestoreExc(nil, nil)
							goto Label10
						Label10:
							// line 186: infos = []
							πF.SetLineno(186)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							µinfos = πTemp001
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µresult, ßwasSuccessful, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label13
							}
							goto Label14
							// line 187: if not result.wasSuccessful():
							πF.SetLineno(187)
						Label13:
							// line 188: self.stream.write("FAILED")
							πF.SetLineno(188)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = ßFAILED.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 189: failed, errored = map(len, (result.failures, result.errors))
							πF.SetLineno(189)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µresult, ßfailures, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µresult, ßerrors, nil); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πTemp007).ToObject()
							πTemp003[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp007}}}, πTemp002); πE != nil {
								continue
							}
							µfailed = πTemp001
							µerrored = πTemp007
							if πE = πg.CheckLocal(πF, µfailed, "failed"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µfailed); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label16
							}
							goto Label17
							// line 190: if failed:
							πF.SetLineno(190)
						Label16:
							// line 191: infos.append("failures=%d" % failed)
							πF.SetLineno(191)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfailed, "failed"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("failures=%d").ToObject(), µfailed); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µinfos, "infos"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µinfos, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label17
						Label17:
							if πE = πg.CheckLocal(πF, µerrored, "errored"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µerrored); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label18
							}
							goto Label19
							// line 192: if errored:
							πF.SetLineno(192)
						Label18:
							// line 193: infos.append("errors=%d" % errored)
							πF.SetLineno(193)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerrored, "errored"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("errors=%d").ToObject(), µerrored); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µinfos, "infos"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µinfos, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label19
						Label19:
							goto Label15
						Label14:
							// line 195: self.stream.write("OK")
							πF.SetLineno(195)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = ßOK.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label15
						Label15:
							if πE = πg.CheckLocal(πF, µskipped, "skipped"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µskipped); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label20
							}
							goto Label21
							// line 196: if skipped:
							πF.SetLineno(196)
						Label20:
							// line 197: infos.append("skipped=%d" % skipped)
							πF.SetLineno(197)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µskipped, "skipped"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("skipped=%d").ToObject(), µskipped); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µinfos, "infos"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µinfos, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label21
						Label21:
							if πE = πg.CheckLocal(πF, µexpectedFails, "expectedFails"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µexpectedFails); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label22
							}
							goto Label23
							// line 198: if expectedFails:
							πF.SetLineno(198)
						Label22:
							// line 199: infos.append("expected failures=%d" % expectedFails)
							πF.SetLineno(199)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µexpectedFails, "expectedFails"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("expected failures=%d").ToObject(), µexpectedFails); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µinfos, "infos"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µinfos, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label23
						Label23:
							if πE = πg.CheckLocal(πF, µunexpectedSuccesses, "unexpectedSuccesses"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µunexpectedSuccesses); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label24
							}
							goto Label25
							// line 200: if unexpectedSuccesses:
							πF.SetLineno(200)
						Label24:
							// line 201: infos.append("unexpected successes=%d" % unexpectedSuccesses)
							πF.SetLineno(201)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µunexpectedSuccesses, "unexpectedSuccesses"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("unexpected successes=%d").ToObject(), µunexpectedSuccesses); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µinfos, "infos"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µinfos, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label25
						Label25:
							if πE = πg.CheckLocal(πF, µinfos, "infos"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µinfos); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label26
							}
							goto Label27
							// line 202: if infos:
							πF.SetLineno(202)
						Label26:
							// line 203: self.stream.writeln(" (%s)" % (", ".join(infos),))
							πF.SetLineno(203)
							πTemp003 = πF.MakeArgs(1)
							πTemp011 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µinfos, "infos"); πE != nil {
								continue
							}
							πTemp011[0] = µinfos
							if πTemp007, πE = πg.GetAttr(πF, πg.NewStr(", ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp011, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp011)
							πTemp002 = πg.NewTuple1(πTemp008).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr(" (%s)").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßwriteln, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label28
						Label27:
							// line 205: self.stream.write("\n")
							πF.SetLineno(205)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label28
						Label28:
							// line 206: return result
							πF.SetLineno(206)
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
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp004); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("TextTestRunner").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTextTestRunner.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("unittest_runner", Code)
}
