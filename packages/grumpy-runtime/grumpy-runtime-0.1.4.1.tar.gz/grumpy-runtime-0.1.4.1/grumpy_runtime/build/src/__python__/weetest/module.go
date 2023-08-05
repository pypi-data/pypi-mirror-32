package weetest
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/weetest.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßAssertionError := πg.InternStr("AssertionError")
		ßBenchmark := πg.InternStr("Benchmark")
		ßERROR := πg.InternStr("ERROR")
		ßException := πg.InternStr("Exception")
		ßFAILED := πg.InternStr("FAILED")
		ßN := πg.InternStr("N")
		ßPASSED := πg.InternStr("PASSED")
		ßResetTimer := πg.InternStr("ResetTimer")
		ßRun := πg.InternStr("Run")
		ßRunBenchmarks := πg.InternStr("RunBenchmarks")
		ßRunTests := πg.InternStr("RunTests")
		ßTest := πg.InternStr("Test")
		ßWEETEST_TARGET := πg.InternStr("WEETEST_TARGET")
		ßXML_OUTPUT_FILE := πg.InternStr("XML_OUTPUT_FILE")
		ß_Benchmark := πg.InternStr("_Benchmark")
		ß_RunAll := πg.InternStr("_RunAll")
		ß_RunOnce := πg.InternStr("_RunOnce")
		ß_RunOneBenchmark := πg.InternStr("_RunOneBenchmark")
		ß_RunOneTest := πg.InternStr("_RunOneTest")
		ß_TestResult := πg.InternStr("_TestResult")
		ß_WriteXmlFile := πg.InternStr("_WriteXmlFile")
		ß__init__ := πg.InternStr("__init__")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßappend := πg.InternStr("append")
		ßargv := πg.InternStr("argv")
		ßbench_func := πg.InternStr("bench_func")
		ßclose := πg.InternStr("close")
		ßdir := πg.InternStr("dir")
		ßduration := πg.InternStr("duration")
		ßenviron := πg.InternStr("environ")
		ßerror := πg.InternStr("error")
		ßexit := πg.InternStr("exit")
		ßfailed := πg.InternStr("failed")
		ßfloat := πg.InternStr("float")
		ßget := πg.InternStr("get")
		ßgetattr := πg.InternStr("getattr")
		ßint := πg.InternStr("int")
		ßisinstance := πg.InternStr("isinstance")
		ßlen := πg.InternStr("len")
		ßmodules := πg.InternStr("modules")
		ßname := πg.InternStr("name")
		ßobject := πg.InternStr("object")
		ßopen := πg.InternStr("open")
		ßops_per_sec := πg.InternStr("ops_per_sec")
		ßos := πg.InternStr("os")
		ßpassed := πg.InternStr("passed")
		ßprint_exc := πg.InternStr("print_exc")
		ßproperties := πg.InternStr("properties")
		ßstart_time := πg.InternStr("start_time")
		ßstartswith := πg.InternStr("startswith")
		ßstatus := πg.InternStr("status")
		ßstr := πg.InternStr("str")
		ßsys := πg.InternStr("sys")
		ßtarget_duration := πg.InternStr("target_duration")
		ßtime := πg.InternStr("time")
		ßtraceback := πg.InternStr("traceback")
		ßw := πg.InternStr("w")
		ßwrite := πg.InternStr("write")
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 15: """Minimal framework for writing tests and benchmarks.
			πF.SetLineno(15)
			// line 38: import os
			πF.SetLineno(38)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 39: import sys
			πF.SetLineno(39)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 40: import time
			πF.SetLineno(40)
			if πTemp002, πE = πg.ImportModule(πF, "time"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtime.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 41: import traceback
			πF.SetLineno(41)
			if πTemp002, πE = πg.ImportModule(πF, "traceback"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtraceback.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 45: class _Benchmark(object):
			πF.SetLineno(45)
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
			_, πE = πg.NewCode("_Benchmark", "build/src/__python__/weetest.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp005 *πg.Object
				_ = πTemp005
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 46: """Wraps and runs a single user defined benchmark function."""
					πF.SetLineno(46)
					// line 48: def __init__(self, bench_func, target_duration):
					πF.SetLineno(48)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "bench_func", Def: nil}
					πTemp002[2] = πg.Param{Name: "target_duration", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/weetest.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µbench_func *πg.Object = πArgs[1]; _ = µbench_func
						var µtarget_duration *πg.Object = πArgs[2]; _ = µtarget_duration
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 49: """Set up this benchmark to run bench_func to be run for target_duration."""
							πF.SetLineno(49)
							// line 50: self.bench_func = bench_func
							πF.SetLineno(50)
							if πE = πg.CheckLocal(πF, µbench_func, "bench_func"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µbench_func); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbench_func, πTemp001); πE != nil {
								continue
							}
							// line 51: self.target_duration = target_duration
							πF.SetLineno(51)
							if πE = πg.CheckLocal(πF, µtarget_duration, "target_duration"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µtarget_duration); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtarget_duration, πTemp001); πE != nil {
								continue
							}
							// line 52: self.start_time = 0
							πF.SetLineno(52)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstart_time, πTemp001); πE != nil {
								continue
							}
							// line 53: self.N = 1
							πF.SetLineno(53)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßN, πTemp001); πE != nil {
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
					// line 55: def Run(self):
					πF.SetLineno(55)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("Run", "build/src/__python__/weetest.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsmall_duration *πg.Object = πg.UnboundLocal; _ = µsmall_duration
						var µN *πg.Object = πg.UnboundLocal; _ = µN
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 56: """Attempt to run this benchmark for the target duration."""
							πF.SetLineno(56)
							// line 57: small_duration = 0.05 * self.target_duration
							πF.SetLineno(57)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtarget_duration, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, πg.NewFloat(0.05).ToObject(), πTemp002); πE != nil {
								continue
							}
							µsmall_duration = πTemp001
							// line 58: self.N = 1
							πF.SetLineno(58)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßN, πTemp001); πE != nil {
								continue
							}
							// line 59: self._RunOnce()
							πF.SetLineno(59)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_RunOnce, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 60: while self.duration < self.target_duration:
							πF.SetLineno(60)
							πF.PushCheckpoint(2)
							πTemp003 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßduration, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßtarget_duration, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßduration, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsmall_duration, "small_duration"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, πTemp002, µsmall_duration); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 61: if self.duration < small_duration:
							πF.SetLineno(61)
						Label4:
							// line 63: N = self.N * 10
							πF.SetLineno(63)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßN, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, πTemp002, πg.NewInt(10).ToObject()); πE != nil {
								continue
							}
							µN = πTemp001
							goto Label6
						Label5:
							// line 68: N = int(self.N * (1.2 * self.target_duration / self.duration))
							πF.SetLineno(68)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßN, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßtarget_duration, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Mul(πF, πg.NewFloat(1.2).ToObject(), πTemp008); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßduration, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Div(πF, πTemp007, πTemp008); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µN = πTemp002
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßN, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, µN); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							goto Label8
							// line 69: if self.N == N:
							πF.SetLineno(69)
						Label7:
							// line 70: self.N += 1
							πF.SetLineno(70)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßN, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßN, πTemp002); πE != nil {
								continue
							}
							goto Label9
						Label8:
							// line 72: self.N = N
							πF.SetLineno(72)
							if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µN); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßN, πTemp001); πE != nil {
								continue
							}
							goto Label9
						Label9:
							// line 73: self._RunOnce()
							πF.SetLineno(73)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_RunOnce, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßRun.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 75: def _RunOnce(self):
					πF.SetLineno(75)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("_RunOnce", "build/src/__python__/weetest.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 76: self.start_time = time.time()
							πF.SetLineno(76)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtime, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstart_time, πTemp002); πE != nil {
								continue
							}
							// line 77: self.bench_func(self)
							πF.SetLineno(77)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbench_func, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 78: self.duration = time.time() - self.start_time
							πF.SetLineno(78)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßtime, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstart_time, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, πTemp002, πTemp004); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßduration, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_RunOnce.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 80: def ResetTimer(self):
					πF.SetLineno(80)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("ResetTimer", "build/src/__python__/weetest.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 81: """Clears the current elapsed time to discount expensive setup steps."""
							πF.SetLineno(81)
							// line 82: self.start_time = time.time()
							πF.SetLineno(82)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtime, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstart_time, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßResetTimer.ToObject(), πTemp005); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("_Benchmark").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Benchmark.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 85: class _TestResult(object):
			πF.SetLineno(85)
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
			_, πE = πg.NewCode("_TestResult", "build/src/__python__/weetest.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 86: """The outcome of running a particular benchmark function."""
					πF.SetLineno(86)
					// line 88: def __init__(self, name):
					πF.SetLineno(88)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/weetest.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Dict
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
							// line 89: self.name = name
							πF.SetLineno(89)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µname); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßname, πTemp001); πE != nil {
								continue
							}
							// line 90: self.status = 'not run'
							πF.SetLineno(90)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewStr("not run").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstatus, πTemp001); πE != nil {
								continue
							}
							// line 91: self.duration = 0
							πF.SetLineno(91)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßduration, πTemp001); πE != nil {
								continue
							}
							// line 92: self.properties = {}
							πF.SetLineno(92)
							πTemp002 = πg.NewDict()
							πTemp001 = πTemp002.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßproperties, πTemp003); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("_TestResult").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_TestResult.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 95: def _RunOneBenchmark(name, test_func):
			πF.SetLineno(95)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "name", Def: nil}
			πTemp006[1] = πg.Param{Name: "test_func", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("_RunOneBenchmark", "build/src/__python__/weetest.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]; _ = µname
				var µtest_func *πg.Object = πArgs[1]; _ = µtest_func
				var µb *πg.Object = πg.UnboundLocal; _ = µb
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µstart_time *πg.Object = πg.UnboundLocal; _ = µstart_time
				var µe *πg.Object = πg.UnboundLocal; _ = µe
				var µops_per_sec *πg.Object = πg.UnboundLocal; _ = µops_per_sec
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.BaseException
				_ = πTemp005
				var πTemp006 *πg.Traceback
				_ = πTemp006
				var πTemp007 bool
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
					// line 96: """Runs a single benchmark and returns a _TestResult."""
					πF.SetLineno(96)
					// line 97: b = _Benchmark(test_func, 1)
					πF.SetLineno(97)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µtest_func, "test_func"); πE != nil {
						continue
					}
					πTemp001[0] = µtest_func
					πTemp001[1] = πg.NewInt(1).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_Benchmark); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µb = πTemp003
					// line 98: result = _TestResult(name)
					πF.SetLineno(98)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001[0] = µname
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_TestResult); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µresult = πTemp003
					// line 99: print name,
					πF.SetLineno(99)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001[0] = µname
					if πE = πg.Print(πF, πTemp001, false); πE != nil {
						continue
					}
					// line 100: start_time = time.time()
					πF.SetLineno(100)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtime, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					µstart_time = πTemp002
					// line 101: try:
					πF.SetLineno(101)
					πF.PushCheckpoint(1)
					πF.PushCheckpoint(2)
					// line 102: b.Run()
					πF.SetLineno(102)
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µb, ßRun, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					// line 108: result.status = 'passed'
					πF.SetLineno(108)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ßpassed.ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µresult, ßstatus, πTemp002); πE != nil {
						continue
					}
					// line 109: ops_per_sec = b.N / b.duration
					πF.SetLineno(109)
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µb, ßN, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µb, ßduration, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Div(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					µops_per_sec = πTemp002
					// line 110: result.properties['ops_per_sec'] = ops_per_sec
					πF.SetLineno(110)
					if πE = πg.CheckLocal(πF, µops_per_sec, "ops_per_sec"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µops_per_sec); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µresult, ßproperties, nil); πE != nil {
						continue
					}
					πTemp004 = ßops_per_sec.ToObject()
					if πE = πg.SetItem(πF, πTemp003, πTemp004, πTemp002); πE != nil {
						continue
					}
					// line 111: print 'PASSED', ops_per_sec
					πF.SetLineno(111)
					πTemp001 = make([]*πg.Object, 2)
					πTemp001[0] = ßPASSED.ToObject()
					if πE = πg.CheckLocal(πF, µops_per_sec, "ops_per_sec"); πE != nil {
						continue
					}
					πTemp001[1] = µops_per_sec
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp005, πTemp006 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label3
					}
					πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
					continue
					// line 103: except Exception as e:  # pylint: disable=broad-except
					πF.SetLineno(103)
				Label3:
					µe = πTemp005.ToObject()
					// line 104: result.status = 'error'
					πF.SetLineno(104)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ßerror.ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µresult, ßstatus, πTemp002); πE != nil {
						continue
					}
					// line 105: print 'ERROR'
					πF.SetLineno(105)
					πTemp001 = make([]*πg.Object, 1)
					πTemp001[0] = ßERROR.ToObject()
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					// line 106: traceback.print_exc()
					πF.SetLineno(106)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtraceback); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßprint_exc, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πF.RestoreExc(nil, nil)
					πF.PopCheckpoint()
					goto Label1
				Label1:
					πTemp005, πTemp006 = πF.RestoreExc(nil, nil)
					// line 113: result.duration = time.time() - start_time
					πF.SetLineno(113)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßtime, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstart_time, "start_time"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, πTemp003, µstart_time); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µresult, ßduration, πTemp003); πE != nil {
						continue
					}
					if πTemp005 != nil {
						πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
						continue
					}
					if πR != nil {
						continue
					}
					// line 114: return result
					πF.SetLineno(114)
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
			if πE = πF.Globals().SetItem(πF, ß_RunOneBenchmark.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 117: def _RunOneTest(name, test_func):
			πF.SetLineno(117)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "name", Def: nil}
			πTemp006[1] = πg.Param{Name: "test_func", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("_RunOneTest", "build/src/__python__/weetest.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]; _ = µname
				var µtest_func *πg.Object = πArgs[1]; _ = µtest_func
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µstart_time *πg.Object = πg.UnboundLocal; _ = µstart_time
				var µe *πg.Object = πg.UnboundLocal; _ = µe
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
				var πTemp007 *πg.Object
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
					// line 118: """Runs a single test function and returns a _TestResult."""
					πF.SetLineno(118)
					// line 119: result = _TestResult(name)
					πF.SetLineno(119)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001[0] = µname
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_TestResult); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µresult = πTemp003
					// line 120: start_time = time.time()
					πF.SetLineno(120)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtime, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					µstart_time = πTemp002
					// line 121: try:
					πF.SetLineno(121)
					πF.PushCheckpoint(1)
					πF.PushCheckpoint(2)
					// line 122: test_func()
					πF.SetLineno(122)
					if πE = πg.CheckLocal(πF, µtest_func, "test_func"); πE != nil {
						continue
					}
					if πTemp002, πE = µtest_func.Call(πF, nil, nil); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					// line 132: result.status = 'passed'
					πF.SetLineno(132)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ßpassed.ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µresult, ßstatus, πTemp002); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label4
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 123: except AssertionError as e:
					πF.SetLineno(123)
				Label3:
					µe = πTemp004.ToObject()
					// line 124: result.status = 'failed'
					πF.SetLineno(124)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ßfailed.ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µresult, ßstatus, πTemp002); πE != nil {
						continue
					}
					// line 125: print name, 'FAILED'
					πF.SetLineno(125)
					πTemp001 = make([]*πg.Object, 2)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001[0] = µname
					πTemp001[1] = ßFAILED.ToObject()
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					// line 126: traceback.print_exc()
					πF.SetLineno(126)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtraceback); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßprint_exc, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πF.RestoreExc(nil, nil)
					πF.PopCheckpoint()
					goto Label1
					// line 127: except Exception as e:  # pylint: disable=broad-except
					πF.SetLineno(127)
				Label4:
					µe = πTemp004.ToObject()
					// line 128: result.status = 'error'
					πF.SetLineno(128)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ßerror.ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µresult, ßstatus, πTemp002); πE != nil {
						continue
					}
					// line 129: print name, 'ERROR'
					πF.SetLineno(129)
					πTemp001 = make([]*πg.Object, 2)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001[0] = µname
					πTemp001[1] = ßERROR.ToObject()
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					// line 130: traceback.print_exc()
					πF.SetLineno(130)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtraceback); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßprint_exc, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πF.RestoreExc(nil, nil)
					πF.PopCheckpoint()
					goto Label1
				Label1:
					πTemp004, πTemp005 = πF.RestoreExc(nil, nil)
					// line 134: result.duration = time.time() - start_time
					πF.SetLineno(134)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßtime, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp007.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstart_time, "start_time"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, πTemp003, µstart_time); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µresult, ßduration, πTemp003); πE != nil {
						continue
					}
					if πTemp004 != nil {
						πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
						continue
					}
					if πR != nil {
						continue
					}
					// line 135: return result
					πF.SetLineno(135)
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
			if πE = πF.Globals().SetItem(πF, ß_RunOneTest.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 138: def _WriteXmlFile(filename, suite_duration, results):
			πF.SetLineno(138)
			πTemp006 = make([]πg.Param, 3)
			πTemp006[0] = πg.Param{Name: "filename", Def: nil}
			πTemp006[1] = πg.Param{Name: "suite_duration", Def: nil}
			πTemp006[2] = πg.Param{Name: "results", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("_WriteXmlFile", "build/src/__python__/weetest.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfilename *πg.Object = πArgs[0]; _ = µfilename
				var µsuite_duration *πg.Object = πArgs[1]; _ = µsuite_duration
				var µresults *πg.Object = πArgs[2]; _ = µresults
				var µxml_file *πg.Object = πg.UnboundLocal; _ = µxml_file
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
				var µformatted *πg.Object = πg.UnboundLocal; _ = µformatted
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πTemp010 bool
				_ = πTemp010
				var πTemp011 bool
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					case 6: goto Label6
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					// line 139: """Given a list of _BenchmarkResults, writes XML test results to filename."""
					πF.SetLineno(139)
					// line 140: xml_file = open(filename, 'w')
					πF.SetLineno(140)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp001[0] = µfilename
					πTemp001[1] = ßw.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µxml_file = πTemp003
					// line 141: xml_file.write('<testsuite name="%s" tests="%s" '
					πF.SetLineno(141)
					πTemp001 = πF.MakeArgs(1)
					πTemp004 = πg.NewInt(0).ToObject()
					if πTemp006, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßargv, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp007, πTemp004); πE != nil {
						continue
					}
					πTemp008 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µresults, "results"); πE != nil {
						continue
					}
					πTemp008[0] = µresults
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πE = πg.CheckLocal(πF, µsuite_duration, "suite_duration"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple3(πTemp005, πTemp006, µsuite_duration).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("<testsuite name=\"%s\" tests=\"%s\" time=\"%f\" runner=\"weetest\">\n").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µxml_file, "xml_file"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µxml_file, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µresults, "results"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µresults); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp009 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp009 {
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
						πTemp010 = !isStop
					} else {
						πTemp010 = true
						µresult = πTemp003
					}
					if πE != nil || !πTemp010 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 145: xml_file.write('  <testcase name="%s" result="completed" '
					πF.SetLineno(145)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µresult, ßname, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µresult, ßduration, nil); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("  <testcase name=\"%s\" result=\"completed\" status=\"run\" time=\"%f\">\n").ToObject(), πTemp004); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µxml_file, "xml_file"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µxml_file, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µresult, ßproperties, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp010 {
						goto Label4
					}
					goto Label5
					// line 148: if result.properties:
					πF.SetLineno(148)
				Label4:
					// line 149: xml_file.write('    <properties>\n')
					πF.SetLineno(149)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("    <properties>\n").ToObject()
					if πE = πg.CheckLocal(πF, µxml_file, "xml_file"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µxml_file, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µresult, ßproperties, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Iter(πF, πTemp004); πE != nil {
						continue
					}
					πF.PushCheckpoint(7)
					πTemp010 = false
				Label6:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp010 {
						πF.PopCheckpoint()
						goto Label8
					}
					if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp011 = !isStop
					} else {
						πTemp011 = true
						µname = πTemp004
					}
					if πE != nil || !πTemp011 {
						continue
					}
					πF.PushCheckpoint(6)            
					// line 151: value = result.properties[name]
					πF.SetLineno(151)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp004 = µname
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µresult, ßproperties, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
						continue
					}
					µvalue = πTemp005
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[0] = µvalue
					if πTemp004, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					πTemp001[1] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp011, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp011 {
						goto Label9
					}
					goto Label10
					// line 152: if isinstance(value, float):
					πF.SetLineno(152)
				Label9:
					// line 153: formatted = '%f' % value
					πF.SetLineno(153)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mod(πF, πg.NewStr("%f").ToObject(), µvalue); πE != nil {
						continue
					}
					µformatted = πTemp004
					goto Label11
				Label10:
					// line 155: formatted = str(value)
					πF.SetLineno(155)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[0] = µvalue
					if πTemp004, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µformatted = πTemp005
					goto Label11
				Label11:
					// line 156: xml_file.write('      <property name="%s" value="%s"></property>\n' %
					πF.SetLineno(156)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µformatted, "formatted"); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple2(µname, µformatted).ToObject()
					if πTemp004, πE = πg.Mod(πF, πg.NewStr("      <property name=\"%s\" value=\"%s\"></property>\n").ToObject(), πTemp005); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					if πE = πg.CheckLocal(πF, µxml_file, "xml_file"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µxml_file, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label7:
					if πE != nil || πR != nil {
						continue
					}
				Label8:
					// line 158: xml_file.write('    </properties>\n')
					πF.SetLineno(158)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("    </properties>\n").ToObject()
					if πE = πg.CheckLocal(πF, µxml_file, "xml_file"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µxml_file, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label5
				Label5:
					// line 159: xml_file.write('  </testcase>\n')
					πF.SetLineno(159)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("  </testcase>\n").ToObject()
					if πE = πg.CheckLocal(πF, µxml_file, "xml_file"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µxml_file, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 160: xml_file.write('</testsuite>')
					πF.SetLineno(160)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("</testsuite>").ToObject()
					if πE = πg.CheckLocal(πF, µxml_file, "xml_file"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µxml_file, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 161: xml_file.close()
					πF.SetLineno(161)
					if πE = πg.CheckLocal(πF, µxml_file, "xml_file"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µxml_file, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_WriteXmlFile.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 164: def _RunAll(test_prefix, runner):
			πF.SetLineno(164)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "test_prefix", Def: nil}
			πTemp006[1] = πg.Param{Name: "runner", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("_RunAll", "build/src/__python__/weetest.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtest_prefix *πg.Object = πArgs[0]; _ = µtest_prefix
				var µrunner *πg.Object = πArgs[1]; _ = µrunner
				var µtarget *πg.Object = πg.UnboundLocal; _ = µtarget
				var µexit_status *πg.Object = πg.UnboundLocal; _ = µexit_status
				var µmod *πg.Object = πg.UnboundLocal; _ = µmod
				var µresults *πg.Object = πg.UnboundLocal; _ = µresults
				var µsuite_start_time *πg.Object = πg.UnboundLocal; _ = µsuite_start_time
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µsuite_duration *πg.Object = πg.UnboundLocal; _ = µsuite_duration
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
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πTemp010 []*πg.Object
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
					// line 165: """Runs all functions in __main__ matching test_prefix using runner."""
					πF.SetLineno(165)
					// line 166: target = os.environ.get('WEETEST_TARGET')
					πF.SetLineno(166)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ßWEETEST_TARGET.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßenviron, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßget, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtarget = πTemp003
					// line 167: exit_status = 0
					πF.SetLineno(167)
					µexit_status = πg.NewInt(0).ToObject()
					// line 168: mod = sys.modules['__main__']
					πF.SetLineno(168)
					πTemp002 = ß__main__.ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßmodules, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
						continue
					}
					µmod = πTemp003
					// line 169: results = []
					πF.SetLineno(169)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µresults = πTemp002
					// line 170: suite_start_time = time.time()
					πF.SetLineno(170)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtime, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					µsuite_start_time = πTemp002
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
						continue
					}
					πTemp001[0] = µmod
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
					πTemp006 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
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
						πTemp007 = !isStop
					} else {
						πTemp007 = true
						µname = πTemp003
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(1)            
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtest_prefix, "test_prefix"); πE != nil {
						continue
					}
					πTemp001[0] = µtest_prefix
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µname, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp003 = πTemp005
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, µtarget); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(!πTemp009).ToObject()
					πTemp004 = πTemp005
					if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, µname, µtarget); πE != nil {
						continue
					}
					πTemp004 = πTemp005
				Label5:
					πTemp003 = πTemp004
				Label4:
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label6
					}
					goto Label7
					// line 172: if name.startswith(test_prefix) and (not target or name == target):
					πF.SetLineno(172)
				Label6:
					// line 173: result = runner(name, getattr(mod, name))
					πF.SetLineno(173)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001[0] = µname
					πTemp010 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
						continue
					}
					πTemp010[0] = µmod
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp010[1] = µname
					if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					πTemp001[1] = πTemp004
					if πE = πg.CheckLocal(πF, µrunner, "runner"); πE != nil {
						continue
					}
					if πTemp003, πE = µrunner.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µresult = πTemp003
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µresult, ßstatus, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.NE(πF, πTemp004, ßpassed.ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label8
					}
					goto Label9
					// line 174: if result.status != 'passed':
					πF.SetLineno(174)
				Label8:
					// line 175: exit_status = 1
					πF.SetLineno(175)
					µexit_status = πg.NewInt(1).ToObject()
					goto Label9
				Label9:
					// line 176: results.append(result)
					πF.SetLineno(176)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp001[0] = µresult
					if πE = πg.CheckLocal(πF, µresults, "results"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µresults, ßappend, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label7
				Label7:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 177: suite_duration = time.time() - suite_start_time
					πF.SetLineno(177)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßtime, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsuite_start_time, "suite_start_time"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, πTemp003, µsuite_start_time); πE != nil {
						continue
					}
					µsuite_duration = πTemp002
					if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßenviron, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Contains(πF, πTemp004, ßXML_OUTPUT_FILE.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label10
					}
					goto Label11
					// line 178: if 'XML_OUTPUT_FILE' in os.environ:
					πF.SetLineno(178)
				Label10:
					// line 179: _WriteXmlFile(os.environ['XML_OUTPUT_FILE'], suite_duration, results)
					πF.SetLineno(179)
					πTemp001 = πF.MakeArgs(3)
					πTemp002 = ßXML_OUTPUT_FILE.ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßenviron, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µsuite_duration, "suite_duration"); πE != nil {
						continue
					}
					πTemp001[1] = µsuite_duration
					if πE = πg.CheckLocal(πF, µresults, "results"); πE != nil {
						continue
					}
					πTemp001[2] = µresults
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_WriteXmlFile); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label11
				Label11:
					// line 180: return exit_status
					πF.SetLineno(180)
					if πE = πg.CheckLocal(πF, µexit_status, "exit_status"); πE != nil {
						continue
					}
					πR = µexit_status
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_RunAll.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 183: def RunBenchmarks():
			πF.SetLineno(183)
			πTemp006 = make([]πg.Param, 0)
			πTemp008 = πg.NewFunction(πg.NewCode("RunBenchmarks", "build/src/__python__/weetest.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 184: """Benchmarks all functions in __main__ with names like BenchmarkXyz()."""
					πF.SetLineno(184)
					// line 185: sys.exit(_RunAll('Benchmark', _RunOneBenchmark))
					πF.SetLineno(185)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = ßBenchmark.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_RunOneBenchmark); πE != nil {
						continue
					}
					πTemp002[1] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_RunAll); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßexit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßRunBenchmarks.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 188: def RunTests():
			πF.SetLineno(188)
			πTemp006 = make([]πg.Param, 0)
			πTemp009 = πg.NewFunction(πg.NewCode("RunTests", "build/src/__python__/weetest.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 189: """Runs all functions in __main__ with names like TestXyz()."""
					πF.SetLineno(189)
					// line 190: sys.exit(_RunAll('Test', _RunOneTest))
					πF.SetLineno(190)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = ßTest.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_RunOneTest); πE != nil {
						continue
					}
					πTemp002[1] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_RunAll); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßexit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßRunTests.ToObject(), πTemp009); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("weetest", Code)
}
