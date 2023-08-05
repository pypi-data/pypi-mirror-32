package weetest_test
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/weetest_test.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßAssertionError := πg.InternStr("AssertionError")
		ßBenchmark := πg.InternStr("Benchmark")
		ßBenchmarkBaz := πg.InternStr("BenchmarkBaz")
		ßBenchmarkFoo := πg.InternStr("BenchmarkFoo")
		ßBenchmarkQux := πg.InternStr("BenchmarkQux")
		ßN := πg.InternStr("N")
		ßNone := πg.InternStr("None")
		ßReset := πg.InternStr("Reset")
		ßResetTimer := πg.InternStr("ResetTimer")
		ßRun := πg.InternStr("Run")
		ßTest := πg.InternStr("Test")
		ßTestBar := πg.InternStr("TestBar")
		ßTestBenchmark := πg.InternStr("TestBenchmark")
		ßTestBenchmarkResetTimer := πg.InternStr("TestBenchmarkResetTimer")
		ßTestFoo := πg.InternStr("TestFoo")
		ßTestRunAll := πg.InternStr("TestRunAll")
		ßTestRunOneBenchmark := πg.InternStr("TestRunOneBenchmark")
		ßTestRunOneBenchmarkError := πg.InternStr("TestRunOneBenchmarkError")
		ßTestRunOneTest := πg.InternStr("TestRunOneTest")
		ßTestWriteXmlFile := πg.InternStr("TestWriteXmlFile")
		ßTrue := πg.InternStr("True")
		ßValueError := πg.InternStr("ValueError")
		ßXML_OUTPUT_FILE := πg.InternStr("XML_OUTPUT_FILE")
		ß_Benchmark := πg.InternStr("_Benchmark")
		ß_RunAll := πg.InternStr("_RunAll")
		ß_RunOneBenchmark := πg.InternStr("_RunOneBenchmark")
		ß_RunOneTest := πg.InternStr("_RunOneTest")
		ß_TestResult := πg.InternStr("_TestResult")
		ß_Timer := πg.InternStr("_Timer")
		ß_WriteXmlFile := πg.InternStr("_WriteXmlFile")
		ß__init__ := πg.InternStr("__init__")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_timer := πg.InternStr("_timer")
		ßbar := πg.InternStr("bar")
		ßbaz := πg.InternStr("baz")
		ßclose := πg.InternStr("close")
		ßduration := πg.InternStr("duration")
		ßenviron := πg.InternStr("environ")
		ßerror := πg.InternStr("error")
		ßfailed := πg.InternStr("failed")
		ßfoo := πg.InternStr("foo")
		ßget := πg.InternStr("get")
		ßglobals := πg.InternStr("globals")
		ßkeys := πg.InternStr("keys")
		ßmkstemp := πg.InternStr("mkstemp")
		ßmodules := πg.InternStr("modules")
		ßobject := πg.InternStr("object")
		ßopen := πg.InternStr("open")
		ßops_per_sec := πg.InternStr("ops_per_sec")
		ßos := πg.InternStr("os")
		ßpassed := πg.InternStr("passed")
		ßproperties := πg.InternStr("properties")
		ßread := πg.InternStr("read")
		ßremove := πg.InternStr("remove")
		ßrun := πg.InternStr("run")
		ßstartswith := πg.InternStr("startswith")
		ßstatus := πg.InternStr("status")
		ßsys := πg.InternStr("sys")
		ßt := πg.InternStr("t")
		ßtempfile := πg.InternStr("tempfile")
		ßtest_name := πg.InternStr("test_name")
		ßtime := πg.InternStr("time")
		ßweetest := πg.InternStr("weetest")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Dict
		_ = πTemp003
		var πTemp004 *πg.Object
		_ = πTemp004
		var πTemp005 *πg.Object
		_ = πTemp005
		var πTemp006 []πg.Param
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
		var πTemp013 bool
		_ = πTemp013
		var πTemp014 *πg.Object
		_ = πTemp014
		var πTemp015 bool
		_ = πTemp015
		var πTemp016 *πg.Object
		_ = πTemp016
		var πTemp017 *πg.Object
		_ = πTemp017
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 3: goto Label3
			case 4: goto Label4
			default: panic("unexpected function state")
			}
			// line 15: import os
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 16: import sys
			πF.SetLineno(16)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 17: import tempfile
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "tempfile"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtempfile.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: import time
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "time"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtime.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 20: import weetest
			πF.SetLineno(20)
			if πTemp002, πE = πg.ImportModule(πF, "weetest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßweetest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 23: class _Timer(object):
			πF.SetLineno(23)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_Timer", "build/src/__python__/weetest_test.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
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
					// line 25: def __init__(self):
					πF.SetLineno(25)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/weetest_test.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 26: self.t = 0
							πF.SetLineno(26)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßt, πTemp001); πE != nil {
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
					// line 28: def Reset(self):
					πF.SetLineno(28)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("Reset", "build/src/__python__/weetest_test.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 29: self.t = 0
							πF.SetLineno(29)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßt, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßReset.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 31: def time(self):  # pylint: disable=invalid-name
					πF.SetLineno(31)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("time", "build/src/__python__/weetest_test.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 32: return self.t
							πF.SetLineno(32)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßt, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtime.ToObject(), πTemp004); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp004, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("_Timer").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Timer.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 35: _timer = _Timer()
			πF.SetLineno(35)
			if πTemp001, πE = πg.ResolveGlobal(πF, ß_Timer); πE != nil {
				continue
			}
			if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_timer.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 37: time.time = _timer.time
			πF.SetLineno(37)
			if πTemp001, πE = πg.ResolveGlobal(πF, ß_timer); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßtime, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp004); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp005, ßtime, πTemp001); πE != nil {
				continue
			}
			// line 40: def TestBenchmark():
			πF.SetLineno(40)
			πTemp006 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("TestBenchmark", "build/src/__python__/weetest_test.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µBenchmarkFoo *πg.Object = πg.UnboundLocal; _ = µBenchmarkFoo
				var µb *πg.Object = πg.UnboundLocal; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
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
					// line 41: def BenchmarkFoo(b):
					πF.SetLineno(41)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "b", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("BenchmarkFoo", "build/src/__python__/weetest_test.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µb *πg.Object = πArgs[0]; _ = µb
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var πTemp001 bool
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
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
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 42: i = 0
							πF.SetLineno(42)
							µi = πg.NewInt(0).ToObject()
							// line 43: while i < b.N:
							πF.SetLineno(43)
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
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µb, ßN, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LT(πF, µi, πTemp004); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πE != nil || !πTemp002 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 44: i += 1.0
							πF.SetLineno(44)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, µi, πg.NewFloat(1.0).ToObject()); πE != nil {
								continue
							}
							µi = πTemp003
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 45: _timer.t += b.N
							πF.SetLineno(45)
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_timer); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßt, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µb, ßN, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IAdd(πF, πTemp004, πTemp003); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ß_timer); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp006, ßt, πTemp005); πE != nil {
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
					µBenchmarkFoo = πTemp001
					// line 47: b = weetest._Benchmark(BenchmarkFoo, 1000)
					πF.SetLineno(47)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µBenchmarkFoo, "BenchmarkFoo"); πE != nil {
						continue
					}
					πTemp003[0] = µBenchmarkFoo
					πTemp003[1] = πg.NewInt(1000).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßweetest); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß_Benchmark, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µb = πTemp004
					// line 48: _timer.Reset()
					πF.SetLineno(48)
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_timer); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßReset, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 49: b.Run()
					πF.SetLineno(49)
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µb, ßRun, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 50: assert b.duration >= 1000, b.duration
					πF.SetLineno(50)
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µb, ßduration, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µb, ßduration, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GE(πF, πTemp006, πg.NewInt(1000).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp005, πTemp004); πE != nil {
						continue
					}
					// line 51: assert b.N >= 1000
					πF.SetLineno(51)
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µb, ßN, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GE(πF, πTemp005, πg.NewInt(1000).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp004, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestBenchmark.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 54: def TestBenchmarkResetTimer():
			πF.SetLineno(54)
			πTemp006 = make([]πg.Param, 0)
			πTemp004 = πg.NewFunction(πg.NewCode("TestBenchmarkResetTimer", "build/src/__python__/weetest_test.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µBenchmarkFoo *πg.Object = πg.UnboundLocal; _ = µBenchmarkFoo
				var µb *πg.Object = πg.UnboundLocal; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
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
					// line 55: def BenchmarkFoo(b):
					πF.SetLineno(55)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "b", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("BenchmarkFoo", "build/src/__python__/weetest_test.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µb *πg.Object = πArgs[0]; _ = µb
						var µi *πg.Object = πg.UnboundLocal; _ = µi
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 56: _timer.t += 10000  # Do an "expensive" initialization.
							πF.SetLineno(56)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_timer); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßt, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, πTemp002, πg.NewInt(10000).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_timer); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp003, ßt, πTemp001); πE != nil {
								continue
							}
							// line 57: b.ResetTimer()
							πF.SetLineno(57)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µb, ßResetTimer, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 58: i = 0
							πF.SetLineno(58)
							µi = πg.NewInt(0).ToObject()
							// line 59: while i < b.N:
							πF.SetLineno(59)
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
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µb, ßN, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µi, πTemp002); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 60: i += 1.0
							πF.SetLineno(60)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µi, πg.NewFloat(1.0).ToObject()); πE != nil {
								continue
							}
							µi = πTemp001
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 61: _timer.t += b.N
							πF.SetLineno(61)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_timer); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßt, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µb, ßN, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, πTemp002, πTemp001); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ß_timer); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp006, ßt, πTemp003); πE != nil {
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
					µBenchmarkFoo = πTemp001
					// line 63: b = weetest._Benchmark(BenchmarkFoo, 1000)
					πF.SetLineno(63)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µBenchmarkFoo, "BenchmarkFoo"); πE != nil {
						continue
					}
					πTemp003[0] = µBenchmarkFoo
					πTemp003[1] = πg.NewInt(1000).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßweetest); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß_Benchmark, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µb = πTemp004
					// line 64: _timer.Reset()
					πF.SetLineno(64)
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_timer); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßReset, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 65: b.Run()
					πF.SetLineno(65)
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µb, ßRun, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 66: assert b.duration < 2000, b.duration
					πF.SetLineno(66)
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µb, ßduration, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µb, ßduration, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.LT(πF, πTemp006, πg.NewInt(2000).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp005, πTemp004); πE != nil {
						continue
					}
					// line 67: assert b.N < 2000, b.duration
					πF.SetLineno(67)
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µb, ßduration, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µb, ßN, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.LT(πF, πTemp006, πg.NewInt(2000).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp005, πTemp004); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestBenchmarkResetTimer.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 70: def TestRunOneBenchmark():
			πF.SetLineno(70)
			πTemp006 = make([]πg.Param, 0)
			πTemp005 = πg.NewFunction(πg.NewCode("TestRunOneBenchmark", "build/src/__python__/weetest_test.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µBenchmarkFoo *πg.Object = πg.UnboundLocal; _ = µBenchmarkFoo
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
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
					// line 71: def BenchmarkFoo(b):
					πF.SetLineno(71)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "b", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("BenchmarkFoo", "build/src/__python__/weetest_test.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µb *πg.Object = πArgs[0]; _ = µb
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var πTemp001 bool
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
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
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 72: i = 0
							πF.SetLineno(72)
							µi = πg.NewInt(0).ToObject()
							// line 73: while i < b.N:
							πF.SetLineno(73)
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
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µb, ßN, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LT(πF, µi, πTemp004); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πE != nil || !πTemp002 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 74: i += 1.0
							πF.SetLineno(74)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, µi, πg.NewFloat(1.0).ToObject()); πE != nil {
								continue
							}
							µi = πTemp003
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 75: _timer.t += b.N
							πF.SetLineno(75)
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_timer); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßt, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µb, ßN, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IAdd(πF, πTemp004, πTemp003); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ß_timer); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp006, ßt, πTemp005); πE != nil {
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
					µBenchmarkFoo = πTemp001
					// line 77: _timer.Reset()
					πF.SetLineno(77)
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_timer); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßReset, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 78: result = weetest._RunOneBenchmark('BenchmarkFoo', BenchmarkFoo)
					πF.SetLineno(78)
					πTemp005 = πF.MakeArgs(2)
					πTemp005[0] = ßBenchmarkFoo.ToObject()
					if πE = πg.CheckLocal(πF, µBenchmarkFoo, "BenchmarkFoo"); πE != nil {
						continue
					}
					πTemp005[1] = µBenchmarkFoo
					if πTemp003, πE = πg.ResolveGlobal(πF, ßweetest); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß_RunOneBenchmark, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µresult = πTemp003
					// line 80: assert result.status == 'passed'
					πF.SetLineno(80)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µresult, ßstatus, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp004, ßpassed.ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp003, nil); πE != nil {
						continue
					}
					// line 81: assert result.properties.get('ops_per_sec') == 1.0
					πF.SetLineno(81)
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = ßops_per_sec.ToObject()
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µresult, ßproperties, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßget, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp003, πE = πg.Eq(πF, πTemp004, πg.NewFloat(1.0).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp003, nil); πE != nil {
						continue
					}
					// line 82: assert result.duration == _timer.time()
					πF.SetLineno(82)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µresult, ßduration, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ß_timer); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßtime, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp007.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp004, πTemp006); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp003, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestRunOneBenchmark.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 85: def TestRunOneBenchmarkError():
			πF.SetLineno(85)
			πTemp006 = make([]πg.Param, 0)
			πTemp007 = πg.NewFunction(πg.NewCode("TestRunOneBenchmarkError", "build/src/__python__/weetest_test.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µBenchmarkFoo *πg.Object = πg.UnboundLocal; _ = µBenchmarkFoo
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
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
					// line 86: def BenchmarkFoo(unused_b):
					πF.SetLineno(86)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "unused_b", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("BenchmarkFoo", "build/src/__python__/weetest_test.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µunused_b *πg.Object = πArgs[0]; _ = µunused_b
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							// line 87: raise ValueError
							πF.SetLineno(87)
							πE = πF.Raise(πTemp001, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					µBenchmarkFoo = πTemp001
					// line 88: result = weetest._RunOneBenchmark('BenchmarkFoo', BenchmarkFoo)
					πF.SetLineno(88)
					πTemp003 = πF.MakeArgs(2)
					πTemp003[0] = ßBenchmarkFoo.ToObject()
					if πE = πg.CheckLocal(πF, µBenchmarkFoo, "BenchmarkFoo"); πE != nil {
						continue
					}
					πTemp003[1] = µBenchmarkFoo
					if πTemp004, πE = πg.ResolveGlobal(πF, ßweetest); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß_RunOneBenchmark, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µresult = πTemp004
					// line 89: assert result.status == 'error'
					πF.SetLineno(89)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µresult, ßstatus, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, πTemp005, ßerror.ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp004, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestRunOneBenchmarkError.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 92: def TestRunOneTest():
			πF.SetLineno(92)
			πTemp006 = make([]πg.Param, 0)
			πTemp008 = πg.NewFunction(πg.NewCode("TestRunOneTest", "build/src/__python__/weetest_test.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µTestFoo *πg.Object = πg.UnboundLocal; _ = µTestFoo
				var µcases *πg.Object = πg.UnboundLocal; _ = µcases
				var µcase *πg.Object = πg.UnboundLocal; _ = µcase
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 *πg.Object
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
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 93: def TestFoo():
					πF.SetLineno(93)
					πTemp002 = make([]πg.Param, 0)
					πTemp001 = πg.NewFunction(πg.NewCode("TestFoo", "build/src/__python__/weetest_test.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 95: _timer.t += 100
							πF.SetLineno(95)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_timer); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßt, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, πTemp002, πg.NewInt(100).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_timer); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp003, ßt, πTemp001); πE != nil {
								continue
							}
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µcase, "case"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µcase, πTemp001); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 96: if case[0]:
							πF.SetLineno(96)
						Label1:
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µcase, "case"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µcase, πTemp001); πE != nil {
								continue
							}
							// line 97: raise case[0]  # pylint: disable=raising-bad-type
							πF.SetLineno(97)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
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
					µTestFoo = πTemp001
					// line 99: cases = [(None, 'passed'),
					πF.SetLineno(99)
					πTemp003 = make([]*πg.Object, 3)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple2(πTemp005, ßpassed.ToObject()).ToObject()
					πTemp003[0] = πTemp004
					if πTemp005, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple2(πTemp005, ßfailed.ToObject()).ToObject()
					πTemp003[1] = πTemp004
					if πTemp005, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple2(πTemp005, ßerror.ToObject()).ToObject()
					πTemp003[2] = πTemp004
					πTemp004 = πg.NewList(πTemp003...).ToObject()
					µcases = πTemp004
					if πE = πg.CheckLocal(πF, µcases, "cases"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Iter(πF, µcases); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp006 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label3
					}
					if πTemp005, πE = πg.Next(πF, πTemp004); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp007 = !isStop
					} else {
						πTemp007 = true
						µcase = πTemp005
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 103: _timer.Reset()
					πF.SetLineno(103)
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_timer); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp005, ßReset, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp008.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 104: result = weetest._RunOneTest('TestFoo', TestFoo)
					πF.SetLineno(104)
					πTemp003 = πF.MakeArgs(2)
					πTemp003[0] = ßTestFoo.ToObject()
					if πE = πg.CheckLocal(πF, µTestFoo, "TestFoo"); πE != nil {
						continue
					}
					πTemp003[1] = µTestFoo
					if πTemp005, πE = πg.ResolveGlobal(πF, ßweetest); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp005, ß_RunOneTest, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp008.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µresult = πTemp005
					// line 105: assert result.status == case[1]
					πF.SetLineno(105)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, µresult, ßstatus, nil); πE != nil {
						continue
					}
					πTemp009 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µcase, "case"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µcase, πTemp009); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, πTemp008, πTemp010); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp005, nil); πE != nil {
						continue
					}
					// line 106: assert result.duration == 100
					πF.SetLineno(106)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, µresult, ßduration, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, πTemp008, πg.NewInt(100).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp005, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestRunOneTest.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 109: def TestWriteXmlFile():
			πF.SetLineno(109)
			πTemp006 = make([]πg.Param, 0)
			πTemp009 = πg.NewFunction(πg.NewCode("TestWriteXmlFile", "build/src/__python__/weetest_test.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µresult_with_properties *πg.Object = πg.UnboundLocal; _ = µresult_with_properties
				var µcases *πg.Object = πg.UnboundLocal; _ = µcases
				var µcase *πg.Object = πg.UnboundLocal; _ = µcase
				var µfd *πg.Object = πg.UnboundLocal; _ = µfd
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var µf *πg.Object = πg.UnboundLocal; _ = µf
				var µcontents *πg.Object = πg.UnboundLocal; _ = µcontents
				var µwant *πg.Object = πg.UnboundLocal; _ = µwant
				var πTemp001 []*πg.Object
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
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 bool
				_ = πTemp010
				var πTemp011 *πg.BaseException
				_ = πTemp011
				var πTemp012 *πg.Traceback
				_ = πTemp012
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					case 4: goto Label4
					case 5: goto Label5
					case 6: goto Label6
					default: panic("unexpected function state")
					}
					// line 110: result_with_properties = weetest._TestResult('foo')
					πF.SetLineno(110)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ßfoo.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßweetest); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß_TestResult, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µresult_with_properties = πTemp002
					// line 111: result_with_properties.properties['bar'] = 'baz'
					πF.SetLineno(111)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ßbaz.ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µresult_with_properties, "result_with_properties"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µresult_with_properties, ßproperties, nil); πE != nil {
						continue
					}
					πTemp004 = ßbar.ToObject()
					if πE = πg.SetItem(πF, πTemp003, πTemp004, πTemp002); πE != nil {
						continue
					}
					// line 112: cases = [([], ['<testsuite ', 'tests="0"', '</testsuite>']),
					πF.SetLineno(112)
					πTemp001 = make([]*πg.Object, 4)
					πTemp005 = make([]*πg.Object, 0)
					πTemp003 = πg.NewList(πTemp005...).ToObject()
					πTemp005 = make([]*πg.Object, 3)
					πTemp005[0] = πg.NewStr("<testsuite ").ToObject()
					πTemp005[1] = πg.NewStr("tests=\"0\"").ToObject()
					πTemp005[2] = πg.NewStr("</testsuite>").ToObject()
					πTemp004 = πg.NewList(πTemp005...).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[0] = πTemp002
					πTemp005 = make([]*πg.Object, 1)
					πTemp006 = πF.MakeArgs(1)
					πTemp006[0] = ßfoo.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßweetest); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß_TestResult, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp005[0] = πTemp003
					πTemp003 = πg.NewList(πTemp005...).ToObject()
					πTemp005 = make([]*πg.Object, 3)
					πTemp005[0] = πg.NewStr("<testsuite ").ToObject()
					πTemp005[1] = πg.NewStr("tests=\"1\"").ToObject()
					πTemp005[2] = πg.NewStr("<testcase name=\"foo\"").ToObject()
					πTemp004 = πg.NewList(πTemp005...).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[1] = πTemp002
					πTemp005 = make([]*πg.Object, 2)
					πTemp006 = πF.MakeArgs(1)
					πTemp006[0] = ßfoo.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßweetest); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß_TestResult, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp005[0] = πTemp003
					πTemp006 = πF.MakeArgs(1)
					πTemp006[0] = ßbar.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßweetest); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß_TestResult, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp005[1] = πTemp003
					πTemp003 = πg.NewList(πTemp005...).ToObject()
					πTemp005 = make([]*πg.Object, 3)
					πTemp005[0] = πg.NewStr("tests=\"2\"").ToObject()
					πTemp005[1] = πg.NewStr("<testcase name=\"foo\"").ToObject()
					πTemp005[2] = πg.NewStr("<testcase name=\"bar\"").ToObject()
					πTemp004 = πg.NewList(πTemp005...).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[2] = πTemp002
					πTemp005 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µresult_with_properties, "result_with_properties"); πE != nil {
						continue
					}
					πTemp005[0] = µresult_with_properties
					πTemp003 = πg.NewList(πTemp005...).ToObject()
					πTemp005 = make([]*πg.Object, 2)
					πTemp005[0] = πg.NewStr("<testcase name=\"foo\"").ToObject()
					πTemp005[1] = πg.NewStr("<property name=\"bar\" value=\"baz\"></property>").ToObject()
					πTemp004 = πg.NewList(πTemp005...).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[3] = πTemp002
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µcases = πTemp002
					if πE = πg.CheckLocal(πF, µcases, "cases"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µcases); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp007 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp007 {
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
						πTemp008 = !isStop
					} else {
						πTemp008 = true
						µcase = πTemp003
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 121: fd, path = tempfile.mkstemp()
					πF.SetLineno(121)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßmkstemp, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp009}}}, πTemp003); πE != nil {
						continue
					}
					µfd = πTemp004
					µpath = πTemp009
					// line 122: os.close(fd)
					πF.SetLineno(122)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp001[0] = µfd
					if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 123: try:
					πF.SetLineno(123)
					πF.PushCheckpoint(4)
					// line 124: weetest._WriteXmlFile(path, 100, case[0])
					πF.SetLineno(124)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp001[0] = µpath
					πTemp001[1] = πg.NewInt(100).ToObject()
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µcase, "case"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µcase, πTemp003); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßweetest); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß_WriteXmlFile, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 125: f = open(path)
					πF.SetLineno(125)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp001[0] = µpath
					if πTemp003, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µf = πTemp004
					// line 126: contents = f.read()
					πF.SetLineno(126)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µf, ßread, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					µcontents = πTemp004
					// line 127: f.close()
					πF.SetLineno(127)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µcase, "case"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, µcase, πTemp004); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Iter(πF, πTemp009); πE != nil {
						continue
					}
					πF.PushCheckpoint(6)
					πTemp008 = false
				Label5:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
						πF.PopCheckpoint()
						goto Label7
					}
					if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp010 = !isStop
					} else {
						πTemp010 = true
						µwant = πTemp004
					}
					if πE != nil || !πTemp010 {
						continue
					}
					πF.PushCheckpoint(5)            
					// line 129: assert want in contents, contents
					πF.SetLineno(129)
					if πE = πg.CheckLocal(πF, µcontents, "contents"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcontents, "contents"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Contains(πF, µcontents, µwant); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(πTemp010).ToObject()
					if πE = πg.Assert(πF, πTemp004, µcontents); πE != nil {
						continue
					}
					continue
				Label6:
					if πE != nil || πR != nil {
						continue
					}
				Label7:
					πF.PopCheckpoint()
					goto Label4
				Label4:
					πTemp011, πTemp012 = πF.RestoreExc(nil, nil)
					// line 131: os.remove(path)
					πF.SetLineno(131)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp001[0] = µpath
					if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßremove, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp011 != nil {
						πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
						continue
					}
					if πR != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestWriteXmlFile.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 134: def TestRunAll():
			πF.SetLineno(134)
			πTemp006 = make([]πg.Param, 0)
			πTemp010 = πg.NewFunction(πg.NewCode("TestRunAll", "build/src/__python__/weetest_test.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µMain *πg.Object = πg.UnboundLocal; _ = µMain
				var µfd *πg.Object = πg.UnboundLocal; _ = µfd
				var µtest_xml *πg.Object = πg.UnboundLocal; _ = µtest_xml
				var µbenchmark_xml *πg.Object = πg.UnboundLocal; _ = µbenchmark_xml
				var µcases *πg.Object = πg.UnboundLocal; _ = µcases
				var µprefix *πg.Object = πg.UnboundLocal; _ = µprefix
				var µrunner *πg.Object = πg.UnboundLocal; _ = µrunner
				var µxml_path *πg.Object = πg.UnboundLocal; _ = µxml_path
				var µwant_run *πg.Object = πg.UnboundLocal; _ = µwant_run
				var µold_main *πg.Object = πg.UnboundLocal; _ = µold_main
				var µmain *πg.Object = πg.UnboundLocal; _ = µmain
				var µf *πg.Object = πg.UnboundLocal; _ = µf
				var µxml *πg.Object = πg.UnboundLocal; _ = µxml
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var πTemp001 *πg.Dict
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
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.BaseException
				_ = πTemp011
				var πTemp012 *πg.Traceback
				_ = πTemp012
				var πTemp013 bool
				_ = πTemp013
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					case 11: goto Label11
					case 12: goto Label12
					case 6: goto Label6
					default: panic("unexpected function state")
					}
					// line 135: class Main(object):
					πF.SetLineno(135)
					πTemp003 = make([]*πg.Object, 1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					πTemp001 = πg.NewDict()
					if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
						continue
					}
					_, πE = πg.NewCode("Main", "build/src/__python__/weetest_test.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
						πClass := πTemp001
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
							// line 137: def __init__(self):
							πF.SetLineno(137)
							πTemp002 = make([]πg.Param, 1)
							πTemp002[0] = πg.Param{Name: "self", Def: nil}
							πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/weetest_test.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
								var πTemp001 *πg.Dict
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
									// line 138: self.run = {}
									πF.SetLineno(138)
									πTemp001 = πg.NewDict()
									πTemp002 = πTemp001.ToObject()
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, µself, ßrun, πTemp003); πE != nil {
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
							// line 140: def TestFoo(self):
							πF.SetLineno(140)
							πTemp002 = make([]πg.Param, 1)
							πTemp002[0] = πg.Param{Name: "self", Def: nil}
							πTemp003 = πg.NewFunction(πg.NewCode("TestFoo", "build/src/__python__/weetest_test.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
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
									// line 141: self.run['TestFoo'] = True
									πF.SetLineno(141)
									if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
										continue
									}
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µself, ßrun, nil); πE != nil {
										continue
									}
									πTemp004 = ßTestFoo.ToObject()
									if πE = πg.SetItem(πF, πTemp003, πTemp004, πTemp002); πE != nil {
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
							if πE = πClass.SetItem(πF, ßTestFoo.ToObject(), πTemp003); πE != nil {
								continue
							}
							// line 143: def TestBar(self):
							πF.SetLineno(143)
							πTemp002 = make([]πg.Param, 1)
							πTemp002[0] = πg.Param{Name: "self", Def: nil}
							πTemp004 = πg.NewFunction(πg.NewCode("TestBar", "build/src/__python__/weetest_test.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
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
									// line 144: self.run['TestBar'] = True
									πF.SetLineno(144)
									if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
										continue
									}
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µself, ßrun, nil); πE != nil {
										continue
									}
									πTemp004 = ßTestBar.ToObject()
									if πE = πg.SetItem(πF, πTemp003, πTemp004, πTemp002); πE != nil {
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
							if πE = πClass.SetItem(πF, ßTestBar.ToObject(), πTemp004); πE != nil {
								continue
							}
							// line 146: def BenchmarkBaz(self, b):
							πF.SetLineno(146)
							πTemp002 = make([]πg.Param, 2)
							πTemp002[0] = πg.Param{Name: "self", Def: nil}
							πTemp002[1] = πg.Param{Name: "b", Def: nil}
							πTemp005 = πg.NewFunction(πg.NewCode("BenchmarkBaz", "build/src/__python__/weetest_test.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
								var µb *πg.Object = πArgs[1]; _ = µb
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
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
									// line 147: self.run['BenchmarkBaz'] = True
									πF.SetLineno(147)
									if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
										continue
									}
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µself, ßrun, nil); πE != nil {
										continue
									}
									πTemp004 = ßBenchmarkBaz.ToObject()
									if πE = πg.SetItem(πF, πTemp003, πTemp004, πTemp002); πE != nil {
										continue
									}
									// line 148: _timer.t += b.N
									πF.SetLineno(148)
									if πTemp001, πE = πg.ResolveGlobal(πF, ß_timer); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßt, nil); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µb, ßN, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πg.IAdd(πF, πTemp002, πTemp001); πE != nil {
										continue
									}
									if πTemp004, πE = πg.ResolveGlobal(πF, ß_timer); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, πTemp004, ßt, πTemp003); πE != nil {
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
							if πE = πClass.SetItem(πF, ßBenchmarkBaz.ToObject(), πTemp005); πE != nil {
								continue
							}
							// line 150: def BenchmarkQux(self, b):
							πF.SetLineno(150)
							πTemp002 = make([]πg.Param, 2)
							πTemp002[0] = πg.Param{Name: "self", Def: nil}
							πTemp002[1] = πg.Param{Name: "b", Def: nil}
							πTemp006 = πg.NewFunction(πg.NewCode("BenchmarkQux", "build/src/__python__/weetest_test.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
								var µb *πg.Object = πArgs[1]; _ = µb
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
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
									// line 151: self.run['BenchmarkQux'] = True
									πF.SetLineno(151)
									if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
										continue
									}
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µself, ßrun, nil); πE != nil {
										continue
									}
									πTemp004 = ßBenchmarkQux.ToObject()
									if πE = πg.SetItem(πF, πTemp003, πTemp004, πTemp002); πE != nil {
										continue
									}
									// line 152: _timer.t += b.N
									πF.SetLineno(152)
									if πTemp001, πE = πg.ResolveGlobal(πF, ß_timer); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßt, nil); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µb, ßN, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πg.IAdd(πF, πTemp002, πTemp001); πE != nil {
										continue
									}
									if πTemp004, πE = πg.ResolveGlobal(πF, ß_timer); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, πTemp004, ßt, πTemp003); πE != nil {
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
							if πE = πClass.SetItem(πF, ßBenchmarkQux.ToObject(), πTemp006); πE != nil {
								continue
							}
						}
						return nil, nil
					}).Eval(πF, πF.Globals(), nil, nil)
					if πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
						continue
					}
					if πTemp004 == nil {
						πTemp004 = πg.TypeType.ToObject()
					}
					if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("Main").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
						continue
					}
					µMain = πTemp005
					// line 154: fd, test_xml = tempfile.mkstemp()
					πF.SetLineno(154)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßmkstemp, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
						continue
					}
					µfd = πTemp004
					µtest_xml = πTemp005
					// line 155: os.close(fd)
					πF.SetLineno(155)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp003[0] = µfd
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßclose, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 156: fd, benchmark_xml = tempfile.mkstemp()
					πF.SetLineno(156)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßmkstemp, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
						continue
					}
					µfd = πTemp004
					µbenchmark_xml = πTemp005
					// line 157: os.close(fd)
					πF.SetLineno(157)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp003[0] = µfd
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßclose, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 158: cases = [('Test', weetest._RunOneTest, None,
					πF.SetLineno(158)
					πTemp003 = make([]*πg.Object, 4)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßweetest); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß_RunOneTest, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.NewDict()
					if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßTestFoo.ToObject(), πTemp006); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßTestBar.ToObject(), πTemp006); πE != nil {
						continue
					}
					πTemp006 = πTemp001.ToObject()
					πTemp002 = πg.NewTuple4(ßTest.ToObject(), πTemp005, πTemp004, πTemp006).ToObject()
					πTemp003[0] = πTemp002
					if πTemp004, πE = πg.ResolveGlobal(πF, ßweetest); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß_RunOneTest, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtest_xml, "test_xml"); πE != nil {
						continue
					}
					πTemp001 = πg.NewDict()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßTestFoo.ToObject(), πTemp004); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßTestBar.ToObject(), πTemp004); πE != nil {
						continue
					}
					πTemp004 = πTemp001.ToObject()
					πTemp002 = πg.NewTuple4(ßTest.ToObject(), πTemp005, µtest_xml, πTemp004).ToObject()
					πTemp003[1] = πTemp002
					if πTemp004, πE = πg.ResolveGlobal(πF, ßweetest); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß_RunOneBenchmark, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.NewDict()
					if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßBenchmarkBaz.ToObject(), πTemp006); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßBenchmarkQux.ToObject(), πTemp006); πE != nil {
						continue
					}
					πTemp006 = πTemp001.ToObject()
					πTemp002 = πg.NewTuple4(ßBenchmark.ToObject(), πTemp005, πTemp004, πTemp006).ToObject()
					πTemp003[2] = πTemp002
					if πTemp004, πE = πg.ResolveGlobal(πF, ßweetest); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß_RunOneBenchmark, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µbenchmark_xml, "benchmark_xml"); πE != nil {
						continue
					}
					πTemp001 = πg.NewDict()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßBenchmarkBaz.ToObject(), πTemp004); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßBenchmarkQux.ToObject(), πTemp004); πE != nil {
						continue
					}
					πTemp004 = πTemp001.ToObject()
					πTemp002 = πg.NewTuple4(ßBenchmark.ToObject(), πTemp005, µbenchmark_xml, πTemp004).ToObject()
					πTemp003[3] = πTemp002
					πTemp002 = πg.NewList(πTemp003...).ToObject()
					µcases = πTemp002
					if πE = πg.CheckLocal(πF, µcases, "cases"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µcases); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp007 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp007 {
						πF.PopCheckpoint()
						goto Label3
					}
					if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp008 = !isStop
					} else {
						πTemp008 = true
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp009}, πg.TieTarget{Target: &πTemp010}}}, πTemp004); πE != nil {
							continue
						}
						µprefix = πTemp005
						µrunner = πTemp006
						µxml_path = πTemp009
						µwant_run = πTemp010
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 167: old_main = sys.modules['__main__']
					πF.SetLineno(167)
					πTemp004 = ß__main__.ToObject()
					if πTemp006, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp006, ßmodules, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					µold_main = πTemp005
					// line 168: sys.modules['__main__'] = main = Main()
					πF.SetLineno(168)
					if πE = πg.CheckLocal(πF, µMain, "Main"); πE != nil {
						continue
					}
					if πTemp004, πE = µMain.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp006, ßmodules, nil); πE != nil {
						continue
					}
					πTemp006 = ß__main__.ToObject()
					if πE = πg.SetItem(πF, πTemp009, πTemp006, πTemp005); πE != nil {
						continue
					}
					µmain = πTemp004
					if πE = πg.CheckLocal(πF, µxml_path, "xml_path"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µxml_path); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label4
					}
					goto Label5
					// line 169: if xml_path:
					πF.SetLineno(169)
				Label4:
					// line 170: os.environ['XML_OUTPUT_FILE'] = xml_path
					πF.SetLineno(170)
					if πE = πg.CheckLocal(πF, µxml_path, "xml_path"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µxml_path); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßenviron, nil); πE != nil {
						continue
					}
					πTemp005 = ßXML_OUTPUT_FILE.ToObject()
					if πE = πg.SetItem(πF, πTemp006, πTemp005, πTemp004); πE != nil {
						continue
					}
					goto Label5
				Label5:
					// line 171: try:
					πF.SetLineno(171)
					πF.PushCheckpoint(6)
					// line 172: weetest._RunAll(prefix, runner)
					πF.SetLineno(172)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
						continue
					}
					πTemp003[0] = µprefix
					if πE = πg.CheckLocal(πF, µrunner, "runner"); πE != nil {
						continue
					}
					πTemp003[1] = µrunner
					if πTemp004, πE = πg.ResolveGlobal(πF, ßweetest); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß_RunAll, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πF.PopCheckpoint()
					goto Label6
				Label6:
					πTemp011, πTemp012 = πF.RestoreExc(nil, nil)
					// line 174: sys.modules['__main__'] = old_main
					πF.SetLineno(174)
					if πE = πg.CheckLocal(πF, µold_main, "old_main"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µold_main); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßmodules, nil); πE != nil {
						continue
					}
					πTemp005 = ß__main__.ToObject()
					if πE = πg.SetItem(πF, πTemp006, πTemp005, πTemp004); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßenviron, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Contains(πF, πTemp006, ßXML_OUTPUT_FILE.ToObject()); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(πTemp008).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label7
					}
					goto Label8
					// line 175: if 'XML_OUTPUT_FILE' in os.environ:
					πF.SetLineno(175)
				Label7:
					// line 176: del os.environ['XML_OUTPUT_FILE']
					πF.SetLineno(176)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßenviron, nil); πE != nil {
						continue
					}
					πTemp004 = ßXML_OUTPUT_FILE.ToObject()
					if πE = πg.DelItem(πF, πTemp005, πTemp004); πE != nil {
						continue
					}
					goto Label8
				Label8:
					if πTemp011 != nil {
						πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
						continue
					}
					if πR != nil {
						continue
					}
					// line 177: assert main.run == want_run, main.run
					πF.SetLineno(177)
					if πE = πg.CheckLocal(πF, µmain, "main"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µmain, ßrun, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmain, "main"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µmain, ßrun, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant_run, "want_run"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, πTemp006, µwant_run); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp005, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µxml_path, "xml_path"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µxml_path); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label9
					}
					goto Label10
					// line 178: if xml_path:
					πF.SetLineno(178)
				Label9:
					// line 179: f = open(xml_path)
					πF.SetLineno(179)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µxml_path, "xml_path"); πE != nil {
						continue
					}
					πTemp003[0] = µxml_path
					if πTemp004, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µf = πTemp005
					// line 180: xml = f.read()
					πF.SetLineno(180)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µf, ßread, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					µxml = πTemp005
					// line 181: f.close()
					πF.SetLineno(181)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 182: os.remove(xml_path)
					πF.SetLineno(182)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µxml_path, "xml_path"); πE != nil {
						continue
					}
					πTemp003[0] = µxml_path
					if πTemp004, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßremove, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.CheckLocal(πF, µwant_run, "want_run"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Iter(πF, µwant_run); πE != nil {
						continue
					}
					πF.PushCheckpoint(12)
					πTemp008 = false
				Label11:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
						πF.PopCheckpoint()
						goto Label13
					}
					if πTemp005, πE = πg.Next(πF, πTemp004); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp013 = !isStop
					} else {
						πTemp013 = true
						µname = πTemp005
					}
					if πE != nil || !πTemp013 {
						continue
					}
					πF.PushCheckpoint(11)            
					// line 184: assert '<testcase name="%s"' % name in xml, xml
					πF.SetLineno(184)
					if πE = πg.CheckLocal(πF, µxml, "xml"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Mod(πF, πg.NewStr("<testcase name=\"%s\"").ToObject(), µname); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µxml, "xml"); πE != nil {
						continue
					}
					if πTemp013, πE = πg.Contains(πF, µxml, πTemp006); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(πTemp013).ToObject()
					if πE = πg.Assert(πF, πTemp005, µxml); πE != nil {
						continue
					}
					continue
				Label12:
					if πE != nil || πR != nil {
						continue
					}
				Label13:
					goto Label10
				Label10:
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
			if πE = πF.Globals().SetItem(πF, ßTestRunAll.ToObject(), πTemp010); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp011, πE = πg.Eq(πF, πTemp012, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp013, πE = πg.IsTrue(πF, πTemp011); πE != nil {
				continue
			}
			if πTemp013 {
				goto Label1
			}
			goto Label2
			// line 187: if __name__ == '__main__':
			πF.SetLineno(187)
		Label1:
			if πTemp012, πE = πg.ResolveGlobal(πF, ßglobals); πE != nil {
				continue
			}
			if πTemp014, πE = πTemp012.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp012, πE = πg.GetAttr(πF, πTemp014, ßkeys, nil); πE != nil {
				continue
			}
			if πTemp014, πE = πTemp012.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp011, πE = πg.Iter(πF, πTemp014); πE != nil {
				continue
			}
			πF.PushCheckpoint(4)
			πTemp013 = false
		Label3:
			if πE != nil || πR != nil {
				continue
			}
			if πTemp013 {
				πF.PopCheckpoint()
				goto Label5
			}
			if πTemp012, πE = πg.Next(πF, πTemp011); πE != nil {
				isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
				if exc != nil {
					πE = exc
				} else if isStop {
					πE = nil
					πF.RestoreExc(nil, nil)
				}
				πTemp015 = !isStop
			} else {
				πTemp015 = true
				if πE = πF.Globals().SetItem(πF, ßtest_name.ToObject(), πTemp012); πE != nil {
					continue
				}
			}
			if πE != nil || !πTemp015 {
				continue
			}
			πF.PushCheckpoint(3)            
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = ßTest.ToObject()
			if πTemp012, πE = πg.ResolveGlobal(πF, ßtest_name); πE != nil {
				continue
			}
			if πTemp014, πE = πg.GetAttr(πF, πTemp012, ßstartswith, nil); πE != nil {
				continue
			}
			if πTemp012, πE = πTemp014.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πTemp015, πE = πg.IsTrue(πF, πTemp012); πE != nil {
				continue
			}
			if πTemp015 {
				goto Label6
			}
			goto Label7
			// line 190: if test_name.startswith('Test'):
			πF.SetLineno(190)
		Label6:
			// line 191: globals()[test_name]()
			πF.SetLineno(191)
			if πTemp014, πE = πg.ResolveGlobal(πF, ßtest_name); πE != nil {
				continue
			}
			πTemp012 = πTemp014
			if πTemp016, πE = πg.ResolveGlobal(πF, ßglobals); πE != nil {
				continue
			}
			if πTemp017, πE = πTemp016.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp014, πE = πg.GetItem(πF, πTemp017, πTemp012); πE != nil {
				continue
			}
			if πTemp012, πE = πTemp014.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label7
		Label7:
			continue
		Label4:
			if πE != nil || πR != nil {
				continue
			}
		Label5:
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("weetest_test", Code)
}
