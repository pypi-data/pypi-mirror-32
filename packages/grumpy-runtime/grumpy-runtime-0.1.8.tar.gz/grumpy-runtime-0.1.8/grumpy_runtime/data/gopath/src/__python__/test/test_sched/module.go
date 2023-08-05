package test_sched
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/test/test_sched.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßCondition := πg.InternStr("Condition")
		ßImportError := πg.InternStr("ImportError")
		ßNone := πg.InternStr("None")
		ßQueue := πg.InternStr("Queue")
		ßTIMEOUT := πg.InternStr("TIMEOUT")
		ßTestCase := πg.InternStr("TestCase")
		ßThread := πg.InternStr("Thread")
		ßTimer := πg.InternStr("Timer")
		ß__enter__ := πg.InternStr("__enter__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__init__ := πg.InternStr("__init__")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_cond := πg.InternStr("_cond")
		ß_stop := πg.InternStr("_stop")
		ß_time := πg.InternStr("_time")
		ßadvance := πg.InternStr("advance")
		ßappend := πg.InternStr("append")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertFalse := πg.InternStr("assertFalse")
		ßassertTrue := πg.InternStr("assertTrue")
		ßcancel := πg.InternStr("cancel")
		ßempty := πg.InternStr("empty")
		ßenter := πg.InternStr("enter")
		ßenterabs := πg.InternStr("enterabs")
		ßget := πg.InternStr("get")
		ßgrumpy := πg.InternStr("grumpy")
		ßis_alive := πg.InternStr("is_alive")
		ßjoin := πg.InternStr("join")
		ßnotify_all := πg.InternStr("notify_all")
		ßobject := πg.InternStr("object")
		ßput := πg.InternStr("put")
		ßqueue := πg.InternStr("queue")
		ßrun := πg.InternStr("run")
		ßrun_unittest := πg.InternStr("run_unittest")
		ßsched := πg.InternStr("sched")
		ßscheduler := πg.InternStr("scheduler")
		ßskip := πg.InternStr("skip")
		ßsleep := πg.InternStr("sleep")
		ßstart := πg.InternStr("start")
		ßtest := πg.InternStr("test")
		ßtest_cancel := πg.InternStr("test_cancel")
		ßtest_cancel_concurrent := πg.InternStr("test_cancel_concurrent")
		ßtest_empty := πg.InternStr("test_empty")
		ßtest_enter := πg.InternStr("test_enter")
		ßtest_enter_concurrent := πg.InternStr("test_enter_concurrent")
		ßtest_enterabs := πg.InternStr("test_enterabs")
		ßtest_main := πg.InternStr("test_main")
		ßtest_priority := πg.InternStr("test_priority")
		ßtest_support := πg.InternStr("test_support")
		ßthreading := πg.InternStr("threading")
		ßtime := πg.InternStr("time")
		ßunittest := πg.InternStr("unittest")
		ßwait := πg.InternStr("wait")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.BaseException
		_ = πTemp003
		var πTemp004 *πg.Traceback
		_ = πTemp004
		var πTemp005 bool
		_ = πTemp005
		var πTemp006 *πg.Dict
		_ = πTemp006
		var πTemp007 *πg.Object
		_ = πTemp007
		var πTemp008 *πg.Object
		_ = πTemp008
		var πTemp009 *πg.Object
		_ = πTemp009
		var πTemp010 []πg.Param
		_ = πTemp010
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 2: goto Label2
			default: panic("unexpected function state")
			}
			// line 2: import sched
			πF.SetLineno(2)
			if πTemp002, πE = πg.ImportModule(πF, "sched"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsched.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 3: import time
			πF.SetLineno(3)
			if πTemp002, πE = πg.ImportModule(πF, "time"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtime.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 4: import unittest
			πF.SetLineno(4)
			if πTemp002, πE = πg.ImportModule(πF, "unittest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 5: import test.test_support
			πF.SetLineno(5)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 7: try:
			πF.SetLineno(7)
			πF.PushCheckpoint(2)
			// line 8: import threading
			πF.SetLineno(8)
			if πTemp002, πE = πg.ImportModule(πF, "threading"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßthreading.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label1
		Label2:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp003, πTemp004 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsInstance(πF, πTemp003.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label3
			}
			πE = πF.Raise(πTemp003.ToObject(), nil, πTemp004.ToObject())
			continue
			// line 9: except ImportError:
			πF.SetLineno(9)
		Label3:
			// line 10: threading = None
			πF.SetLineno(10)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßthreading.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label1
		Label1:
			// line 12: TIMEOUT = 10
			πF.SetLineno(12)
			if πE = πF.Globals().SetItem(πF, ßTIMEOUT.ToObject(), πg.NewInt(10).ToObject()); πE != nil {
				continue
			}
			// line 15: class Timer(object):
			πF.SetLineno(15)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Timer", "build/src/__python__/test/test_sched.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
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
					// line 16: def __init__(self):
					πF.SetLineno(16)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/test/test_sched.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 17: self._cond = threading.Condition()
							πF.SetLineno(17)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßCondition, nil); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ß_cond, πTemp002); πE != nil {
								continue
							}
							// line 18: self._time = 0
							πF.SetLineno(18)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_time, πTemp001); πE != nil {
								continue
							}
							// line 19: self._stop = 0
							πF.SetLineno(19)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_stop, πTemp001); πE != nil {
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
					// line 21: def time(self):
					πF.SetLineno(21)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("time", "build/src/__python__/test/test_sched.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 *πg.BaseException
						_ = πTemp006
						var πTemp007 *πg.Traceback
						_ = πTemp007
						var πTemp008 *πg.Type
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 22: with self._cond:
							πF.SetLineno(22)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_cond, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp003.Call(πF, πg.Args{πTemp001}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(1)
							// line 23: return self._time
							πF.SetLineno(23)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_time, nil); πE != nil {
								continue
							}
							πR = πTemp004
							continue
							πF.PopCheckpoint()
						Label1:
							πTemp006, πTemp007 = nil, nil
							if πE != nil {
								πTemp006, πTemp007 = πF.ExcInfo()
							}
							if πTemp006 != nil {
								πTemp008 = πTemp006.Type()
								if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πTemp008.ToObject(), πTemp006.ToObject(), πTemp007.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp006 != nil && πTemp005 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
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
					if πE = πClass.SetItem(πF, ßtime.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 26: def sleep(self, t):
					πF.SetLineno(26)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "t", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("sleep", "build/src/__python__/test/test_sched.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µt *πg.Object = πArgs[1]; _ = µt
						var πTemp001 *πg.Object
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
						var πTemp008 *πg.BaseException
						_ = πTemp008
						var πTemp009 *πg.Traceback
						_ = πTemp009
						var πTemp010 *πg.Type
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							case 3: goto Label3
							default: panic("unexpected function state")
							}
							// line 27: assert t >= 0
							πF.SetLineno(27)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GE(πF, µt, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
								continue
							}
							// line 28: with self._cond:
							πF.SetLineno(28)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_cond, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp003.Call(πF, πg.Args{πTemp001}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(1)
							// line 29: t += self._time
							πF.SetLineno(29)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_time, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IAdd(πF, µt, πTemp004); πE != nil {
								continue
							}
							µt = πTemp005
							// line 30: while self._stop < t:
							πF.SetLineno(30)
							πF.PushCheckpoint(3)
							πTemp006 = false
						Label2:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_stop, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.LT(πF, πTemp005, µt); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(2)            
							// line 31: self._time = self._stop
							πF.SetLineno(31)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_stop, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_time, πTemp005); πE != nil {
								continue
							}
							// line 32: self._cond.wait()
							πF.SetLineno(32)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_cond, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßwait, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label3:
							if πE != nil || πR != nil {
								continue
							}
						Label4:
							// line 33: self._time = t
							πF.SetLineno(33)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µt); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_time, πTemp004); πE != nil {
								continue
							}
							πF.PopCheckpoint()
						Label1:
							πTemp008, πTemp009 = nil, nil
							if πE != nil {
								πTemp008, πTemp009 = πF.ExcInfo()
							}
							if πTemp008 != nil {
								πTemp010 = πTemp008.Type()
								if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πTemp010.ToObject(), πTemp008.ToObject(), πTemp009.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 != nil && πTemp006 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
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
					if πE = πClass.SetItem(πF, ßsleep.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 36: def advance(self, t):
					πF.SetLineno(36)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "t", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("advance", "build/src/__python__/test/test_sched.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µt *πg.Object = πArgs[1]; _ = µt
						var πTemp001 *πg.Object
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
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
						_ = πTemp008
						var πTemp009 *πg.Type
						_ = πTemp009
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 37: assert t >= 0
							πF.SetLineno(37)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GE(πF, µt, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
								continue
							}
							// line 38: with self._cond:
							πF.SetLineno(38)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_cond, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp003.Call(πF, πg.Args{πTemp001}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(1)
							// line 39: self._stop += t
							πF.SetLineno(39)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_stop, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IAdd(πF, πTemp004, µt); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_stop, πTemp005); πE != nil {
								continue
							}
							// line 40: self._cond.notify_all()
							πF.SetLineno(40)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_cond, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßnotify_all, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							πF.PopCheckpoint()
						Label1:
							πTemp007, πTemp008 = nil, nil
							if πE != nil {
								πTemp007, πTemp008 = πF.ExcInfo()
							}
							if πTemp007 != nil {
								πTemp009 = πTemp007.Type()
								if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πTemp009.ToObject(), πTemp007.ToObject(), πTemp008.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 != nil && πTemp006 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
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
					if πE = πClass.SetItem(πF, ßadvance.ToObject(), πTemp005); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp007, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp007 == nil {
				πTemp007 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp007.Call(πF, []*πg.Object{πg.NewStr("Timer").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTimer.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 43: class TestCase(unittest.TestCase):
			πF.SetLineno(43)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßTestCase, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TestCase", "build/src/__python__/test/test_sched.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
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
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 45: def test_enter(self):
					πF.SetLineno(45)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_enter", "build/src/__python__/test/test_sched.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µl *πg.Object = πg.UnboundLocal; _ = µl
						var µfun *πg.Object = πg.UnboundLocal; _ = µfun
						var µscheduler *πg.Object = πg.UnboundLocal; _ = µscheduler
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var µz *πg.Object = πg.UnboundLocal; _ = µz
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 46: l = []
							πF.SetLineno(46)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µl = πTemp002
							// line 47: fun = lambda x: l.append(x)
							πF.SetLineno(47)
							πTemp003 = make([]πg.Param, 1)
							πTemp003[0] = πg.Param{Name: "x", Def: nil}
							πTemp002 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/test/test_sched.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µx *πg.Object = πArgs[0]; _ = µx
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
									// line 47: fun = lambda x: l.append(x)
									πF.SetLineno(47)
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
										continue
									}
									πTemp001[0] = µx
									if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µl, ßappend, nil); πE != nil {
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
							µfun = πTemp002
							// line 48: scheduler = sched.scheduler(time.time, time.sleep)
							πF.SetLineno(48)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßtime, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßsleep, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsched); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßscheduler, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µscheduler = πTemp002
							πTemp001 = make([]*πg.Object, 5)
							πTemp001[0] = πg.NewFloat(0.5).ToObject()
							πTemp001[1] = πg.NewFloat(0.4).ToObject()
							πTemp001[2] = πg.NewFloat(0.3).ToObject()
							πTemp001[3] = πg.NewFloat(0.2).ToObject()
							πTemp001[4] = πg.NewFloat(0.1).ToObject()
							πTemp004 = πg.NewList(πTemp001...).ToObject()
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
							if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µx = πTemp004
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 50: z = scheduler.enter(x, 1, fun, (x,))
							πF.SetLineno(50)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							πTemp001[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfun, "fun"); πE != nil {
								continue
							}
							πTemp001[2] = µfun
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple1(µx).ToObject()
							πTemp001[3] = πTemp004
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µscheduler, ßenter, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µz = πTemp007
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 51: scheduler.run()
							πF.SetLineno(51)
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µscheduler, ßrun, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 52: self.assertEqual(l, [0.1, 0.2, 0.3, 0.4, 0.5])
							πF.SetLineno(52)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
								continue
							}
							πTemp001[0] = µl
							πTemp008 = make([]*πg.Object, 5)
							πTemp008[0] = πg.NewFloat(0.1).ToObject()
							πTemp008[1] = πg.NewFloat(0.2).ToObject()
							πTemp008[2] = πg.NewFloat(0.3).ToObject()
							πTemp008[3] = πg.NewFloat(0.4).ToObject()
							πTemp008[4] = πg.NewFloat(0.5).ToObject()
							πTemp002 = πg.NewList(πTemp008...).ToObject()
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_enter.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 54: def test_enterabs(self):
					πF.SetLineno(54)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_enterabs", "build/src/__python__/test/test_sched.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µl *πg.Object = πg.UnboundLocal; _ = µl
						var µfun *πg.Object = πg.UnboundLocal; _ = µfun
						var µscheduler *πg.Object = πg.UnboundLocal; _ = µscheduler
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var µz *πg.Object = πg.UnboundLocal; _ = µz
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 55: l = []
							πF.SetLineno(55)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µl = πTemp002
							// line 56: fun = lambda x: l.append(x)
							πF.SetLineno(56)
							πTemp003 = make([]πg.Param, 1)
							πTemp003[0] = πg.Param{Name: "x", Def: nil}
							πTemp002 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/test/test_sched.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µx *πg.Object = πArgs[0]; _ = µx
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
									// line 56: fun = lambda x: l.append(x)
									πF.SetLineno(56)
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
										continue
									}
									πTemp001[0] = µx
									if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µl, ßappend, nil); πE != nil {
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
							µfun = πTemp002
							// line 57: scheduler = sched.scheduler(time.time, time.sleep)
							πF.SetLineno(57)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßtime, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßsleep, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsched); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßscheduler, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µscheduler = πTemp002
							πTemp001 = make([]*πg.Object, 5)
							πTemp001[0] = πg.NewFloat(0.05).ToObject()
							πTemp001[1] = πg.NewFloat(0.04).ToObject()
							πTemp001[2] = πg.NewFloat(0.03).ToObject()
							πTemp001[3] = πg.NewFloat(0.02).ToObject()
							πTemp001[4] = πg.NewFloat(0.01).ToObject()
							πTemp004 = πg.NewList(πTemp001...).ToObject()
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
							if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µx = πTemp004
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 59: z = scheduler.enterabs(x, 1, fun, (x,))
							πF.SetLineno(59)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							πTemp001[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfun, "fun"); πE != nil {
								continue
							}
							πTemp001[2] = µfun
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple1(µx).ToObject()
							πTemp001[3] = πTemp004
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µscheduler, ßenterabs, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µz = πTemp007
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 60: scheduler.run()
							πF.SetLineno(60)
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µscheduler, ßrun, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 61: self.assertEqual(l, [0.01, 0.02, 0.03, 0.04, 0.05])
							πF.SetLineno(61)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
								continue
							}
							πTemp001[0] = µl
							πTemp008 = make([]*πg.Object, 5)
							πTemp008[0] = πg.NewFloat(0.01).ToObject()
							πTemp008[1] = πg.NewFloat(0.02).ToObject()
							πTemp008[2] = πg.NewFloat(0.03).ToObject()
							πTemp008[3] = πg.NewFloat(0.04).ToObject()
							πTemp008[4] = πg.NewFloat(0.05).ToObject()
							πTemp002 = πg.NewList(πTemp008...).ToObject()
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_enterabs.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 65: def test_enter_concurrent(self):
					πF.SetLineno(65)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_enter_concurrent", "build/src/__python__/test/test_sched.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µq *πg.Object = πg.UnboundLocal; _ = µq
						var µfun *πg.Object = πg.UnboundLocal; _ = µfun
						var µtimer *πg.Object = πg.UnboundLocal; _ = µtimer
						var µscheduler *πg.Object = πg.UnboundLocal; _ = µscheduler
						var µt *πg.Object = πg.UnboundLocal; _ = µt
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var µz *πg.Object = πg.UnboundLocal; _ = µz
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 πg.KWArgs
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
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 66: q = queue.Queue()
							πF.SetLineno(66)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßqueue); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßQueue, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µq = πTemp001
							// line 67: fun = q.put
							πF.SetLineno(67)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							µfun = πTemp001
							// line 68: timer = Timer()
							πF.SetLineno(68)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTimer); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µtimer = πTemp002
							// line 69: scheduler = sched.scheduler(timer.time, timer.sleep)
							πF.SetLineno(69)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtimer, "timer"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtimer, ßtime, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µtimer, "timer"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtimer, ßsleep, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsched); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßscheduler, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µscheduler = πTemp001
							// line 70: scheduler.enter(1, 1, fun, (1,))
							πF.SetLineno(70)
							πTemp003 = πF.MakeArgs(4)
							πTemp003[0] = πg.NewInt(1).ToObject()
							πTemp003[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfun, "fun"); πE != nil {
								continue
							}
							πTemp003[2] = µfun
							πTemp001 = πg.NewTuple1(πg.NewInt(1).ToObject()).ToObject()
							πTemp003[3] = πTemp001
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µscheduler, ßenter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 71: scheduler.enter(3, 1, fun, (3,))
							πF.SetLineno(71)
							πTemp003 = πF.MakeArgs(4)
							πTemp003[0] = πg.NewInt(3).ToObject()
							πTemp003[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfun, "fun"); πE != nil {
								continue
							}
							πTemp003[2] = µfun
							πTemp001 = πg.NewTuple1(πg.NewInt(3).ToObject()).ToObject()
							πTemp003[3] = πTemp001
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µscheduler, ßenter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 72: t = threading.Thread(target=scheduler.run)
							πF.SetLineno(72)
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µscheduler, ßrun, nil); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"target", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßThread, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							µt = πTemp001
							// line 73: t.start()
							πF.SetLineno(73)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µt, ßstart, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 74: timer.advance(1)
							πF.SetLineno(74)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µtimer, "timer"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtimer, ßadvance, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 75: self.assertEqual(q.get(timeout=TIMEOUT), 1)
							πF.SetLineno(75)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTIMEOUT); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"timeout", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 76: self.assertTrue(q.empty())
							πF.SetLineno(76)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßempty, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp003 = make([]*πg.Object, 3)
							πTemp003[0] = πg.NewInt(4).ToObject()
							πTemp003[1] = πg.NewInt(5).ToObject()
							πTemp003[2] = πg.NewInt(2).ToObject()
							πTemp002 = πg.NewList(πTemp003...).ToObject()
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µx = πTemp002
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 78: z = scheduler.enter(x - 1, 1, fun, (x,))
							πF.SetLineno(78)
							πTemp003 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, µx, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfun, "fun"); πE != nil {
								continue
							}
							πTemp003[2] = µfun
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple1(µx).ToObject()
							πTemp003[3] = πTemp002
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µscheduler, ßenter, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µz = πTemp007
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 79: timer.advance(2)
							πF.SetLineno(79)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µtimer, "timer"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtimer, ßadvance, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 80: self.assertEqual(q.get(timeout=TIMEOUT), 2)
							πF.SetLineno(80)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTIMEOUT); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"timeout", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 81: self.assertEqual(q.get(timeout=TIMEOUT), 3)
							πF.SetLineno(81)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTIMEOUT); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"timeout", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 82: self.assertTrue(q.empty())
							πF.SetLineno(82)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßempty, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 83: timer.advance(1)
							πF.SetLineno(83)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µtimer, "timer"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtimer, ßadvance, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 84: self.assertEqual(q.get(timeout=TIMEOUT), 4)
							πF.SetLineno(84)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTIMEOUT); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"timeout", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewInt(4).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 85: self.assertTrue(q.empty())
							πF.SetLineno(85)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßempty, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 86: timer.advance(1)
							πF.SetLineno(86)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µtimer, "timer"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtimer, ßadvance, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 87: self.assertEqual(q.get(timeout=TIMEOUT), 5)
							πF.SetLineno(87)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTIMEOUT); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"timeout", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewInt(5).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 88: self.assertTrue(q.empty())
							πF.SetLineno(88)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßempty, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 89: timer.advance(1000)
							πF.SetLineno(89)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(1000).ToObject()
							if πE = πg.CheckLocal(πF, µtimer, "timer"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtimer, ßadvance, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 90: t.join(timeout=TIMEOUT)
							πF.SetLineno(90)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTIMEOUT); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"timeout", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µt, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							// line 91: self.assertFalse(t.is_alive())
							πF.SetLineno(91)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µt, ßis_alive, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 92: self.assertTrue(q.empty())
							πF.SetLineno(92)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßempty, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 93: self.assertEqual(timer.time(), 5)
							πF.SetLineno(93)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtimer, "timer"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtimer, ßtime, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewInt(5).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_enter_concurrent.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 64: @unittest.skip('grumpy')
					πF.SetLineno(64)
					πTemp005 = πF.MakeArgs(1)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßtest_enter_concurrent); πE != nil {
						continue
					}
					πTemp005[0] = πTemp006
					πTemp007 = πF.MakeArgs(1)
					πTemp007[0] = ßgrumpy.ToObject()
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp006, ßskip, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πTemp008, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πClass.SetItem(πF, ßtest_enter_concurrent.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 95: def test_priority(self):
					πF.SetLineno(95)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_priority", "build/src/__python__/test/test_sched.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µl *πg.Object = πg.UnboundLocal; _ = µl
						var µfun *πg.Object = πg.UnboundLocal; _ = µfun
						var µscheduler *πg.Object = πg.UnboundLocal; _ = µscheduler
						var µpriority *πg.Object = πg.UnboundLocal; _ = µpriority
						var µz *πg.Object = πg.UnboundLocal; _ = µz
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 96: l = []
							πF.SetLineno(96)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µl = πTemp002
							// line 97: fun = lambda x: l.append(x)
							πF.SetLineno(97)
							πTemp003 = make([]πg.Param, 1)
							πTemp003[0] = πg.Param{Name: "x", Def: nil}
							πTemp002 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/test/test_sched.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µx *πg.Object = πArgs[0]; _ = µx
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
									// line 97: fun = lambda x: l.append(x)
									πF.SetLineno(97)
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
										continue
									}
									πTemp001[0] = µx
									if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µl, ßappend, nil); πE != nil {
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
							µfun = πTemp002
							// line 98: scheduler = sched.scheduler(time.time, time.sleep)
							πF.SetLineno(98)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßtime, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßsleep, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsched); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßscheduler, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µscheduler = πTemp002
							πTemp001 = make([]*πg.Object, 5)
							πTemp001[0] = πg.NewInt(1).ToObject()
							πTemp001[1] = πg.NewInt(2).ToObject()
							πTemp001[2] = πg.NewInt(3).ToObject()
							πTemp001[3] = πg.NewInt(4).ToObject()
							πTemp001[4] = πg.NewInt(5).ToObject()
							πTemp004 = πg.NewList(πTemp001...).ToObject()
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
							if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µpriority = πTemp004
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 100: z = scheduler.enterabs(0.01, priority, fun, (priority,))
							πF.SetLineno(100)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewFloat(0.01).ToObject()
							if πE = πg.CheckLocal(πF, µpriority, "priority"); πE != nil {
								continue
							}
							πTemp001[1] = µpriority
							if πE = πg.CheckLocal(πF, µfun, "fun"); πE != nil {
								continue
							}
							πTemp001[2] = µfun
							if πE = πg.CheckLocal(πF, µpriority, "priority"); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple1(µpriority).ToObject()
							πTemp001[3] = πTemp004
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µscheduler, ßenterabs, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µz = πTemp007
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 101: scheduler.run()
							πF.SetLineno(101)
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µscheduler, ßrun, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 102: self.assertEqual(l, [1, 2, 3, 4, 5])
							πF.SetLineno(102)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
								continue
							}
							πTemp001[0] = µl
							πTemp008 = make([]*πg.Object, 5)
							πTemp008[0] = πg.NewInt(1).ToObject()
							πTemp008[1] = πg.NewInt(2).ToObject()
							πTemp008[2] = πg.NewInt(3).ToObject()
							πTemp008[3] = πg.NewInt(4).ToObject()
							πTemp008[4] = πg.NewInt(5).ToObject()
							πTemp002 = πg.NewList(πTemp008...).ToObject()
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_priority.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 105: def test_cancel(self):
					πF.SetLineno(105)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("test_cancel", "build/src/__python__/test/test_sched.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µl *πg.Object = πg.UnboundLocal; _ = µl
						var µfun *πg.Object = πg.UnboundLocal; _ = µfun
						var µscheduler *πg.Object = πg.UnboundLocal; _ = µscheduler
						var µnow *πg.Object = πg.UnboundLocal; _ = µnow
						var µevent1 *πg.Object = πg.UnboundLocal; _ = µevent1
						var µevent2 *πg.Object = πg.UnboundLocal; _ = µevent2
						var µevent3 *πg.Object = πg.UnboundLocal; _ = µevent3
						var µevent4 *πg.Object = πg.UnboundLocal; _ = µevent4
						var µevent5 *πg.Object = πg.UnboundLocal; _ = µevent5
						var πTemp001 []*πg.Object
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
							// line 106: l = []
							πF.SetLineno(106)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µl = πTemp002
							// line 107: fun = lambda x: l.append(x)
							πF.SetLineno(107)
							πTemp003 = make([]πg.Param, 1)
							πTemp003[0] = πg.Param{Name: "x", Def: nil}
							πTemp002 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/test/test_sched.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µx *πg.Object = πArgs[0]; _ = µx
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
									// line 107: fun = lambda x: l.append(x)
									πF.SetLineno(107)
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
										continue
									}
									πTemp001[0] = µx
									if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µl, ßappend, nil); πE != nil {
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
							µfun = πTemp002
							// line 108: scheduler = sched.scheduler(time.time, time.sleep)
							πF.SetLineno(108)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßtime, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßsleep, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsched); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßscheduler, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µscheduler = πTemp002
							// line 109: now = time.time()
							πF.SetLineno(109)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßtime, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							µnow = πTemp002
							// line 110: event1 = scheduler.enterabs(now + 0.01, 1, fun, (0.01,))
							πF.SetLineno(110)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µnow, "now"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µnow, πg.NewFloat(0.01).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfun, "fun"); πE != nil {
								continue
							}
							πTemp001[2] = µfun
							πTemp002 = πg.NewTuple1(πg.NewFloat(0.01).ToObject()).ToObject()
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µscheduler, ßenterabs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µevent1 = πTemp004
							// line 111: event2 = scheduler.enterabs(now + 0.02, 1, fun, (0.02,))
							πF.SetLineno(111)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µnow, "now"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µnow, πg.NewFloat(0.02).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfun, "fun"); πE != nil {
								continue
							}
							πTemp001[2] = µfun
							πTemp002 = πg.NewTuple1(πg.NewFloat(0.02).ToObject()).ToObject()
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µscheduler, ßenterabs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µevent2 = πTemp004
							// line 112: event3 = scheduler.enterabs(now + 0.03, 1, fun, (0.03,))
							πF.SetLineno(112)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µnow, "now"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µnow, πg.NewFloat(0.03).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfun, "fun"); πE != nil {
								continue
							}
							πTemp001[2] = µfun
							πTemp002 = πg.NewTuple1(πg.NewFloat(0.03).ToObject()).ToObject()
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µscheduler, ßenterabs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µevent3 = πTemp004
							// line 113: event4 = scheduler.enterabs(now + 0.04, 1, fun, (0.04,))
							πF.SetLineno(113)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µnow, "now"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µnow, πg.NewFloat(0.04).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfun, "fun"); πE != nil {
								continue
							}
							πTemp001[2] = µfun
							πTemp002 = πg.NewTuple1(πg.NewFloat(0.04).ToObject()).ToObject()
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µscheduler, ßenterabs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µevent4 = πTemp004
							// line 114: event5 = scheduler.enterabs(now + 0.05, 1, fun, (0.05,))
							πF.SetLineno(114)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µnow, "now"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µnow, πg.NewFloat(0.05).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfun, "fun"); πE != nil {
								continue
							}
							πTemp001[2] = µfun
							πTemp002 = πg.NewTuple1(πg.NewFloat(0.05).ToObject()).ToObject()
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µscheduler, ßenterabs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µevent5 = πTemp004
							// line 115: scheduler.cancel(event1)
							πF.SetLineno(115)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µevent1, "event1"); πE != nil {
								continue
							}
							πTemp001[0] = µevent1
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µscheduler, ßcancel, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 116: scheduler.cancel(event5)
							πF.SetLineno(116)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µevent5, "event5"); πE != nil {
								continue
							}
							πTemp001[0] = µevent5
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µscheduler, ßcancel, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 117: scheduler.run()
							πF.SetLineno(117)
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µscheduler, ßrun, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 118: self.assertEqual(l, [0.02, 0.03, 0.04])
							πF.SetLineno(118)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
								continue
							}
							πTemp001[0] = µl
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewFloat(0.02).ToObject()
							πTemp005[1] = πg.NewFloat(0.03).ToObject()
							πTemp005[2] = πg.NewFloat(0.04).ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_cancel.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 104: @unittest.skip('grumpy')
					πF.SetLineno(104)
					πTemp005 = πF.MakeArgs(1)
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßtest_cancel); πE != nil {
						continue
					}
					πTemp005[0] = πTemp009
					πTemp007 = πF.MakeArgs(1)
					πTemp007[0] = ßgrumpy.ToObject()
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßskip, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp010.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πTemp010, πE = πTemp009.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πClass.SetItem(πF, ßtest_cancel.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 122: def test_cancel_concurrent(self):
					πF.SetLineno(122)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("test_cancel_concurrent", "build/src/__python__/test/test_sched.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µq *πg.Object = πg.UnboundLocal; _ = µq
						var µfun *πg.Object = πg.UnboundLocal; _ = µfun
						var µtimer *πg.Object = πg.UnboundLocal; _ = µtimer
						var µscheduler *πg.Object = πg.UnboundLocal; _ = µscheduler
						var µnow *πg.Object = πg.UnboundLocal; _ = µnow
						var µevent1 *πg.Object = πg.UnboundLocal; _ = µevent1
						var µevent2 *πg.Object = πg.UnboundLocal; _ = µevent2
						var µevent4 *πg.Object = πg.UnboundLocal; _ = µevent4
						var µevent5 *πg.Object = πg.UnboundLocal; _ = µevent5
						var µevent3 *πg.Object = πg.UnboundLocal; _ = µevent3
						var µt *πg.Object = πg.UnboundLocal; _ = µt
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 πg.KWArgs
						_ = πTemp004
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 123: q = queue.Queue()
							πF.SetLineno(123)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßqueue); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßQueue, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µq = πTemp001
							// line 124: fun = q.put
							πF.SetLineno(124)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							µfun = πTemp001
							// line 125: timer = Timer()
							πF.SetLineno(125)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTimer); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µtimer = πTemp002
							// line 126: scheduler = sched.scheduler(timer.time, timer.sleep)
							πF.SetLineno(126)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtimer, "timer"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtimer, ßtime, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µtimer, "timer"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtimer, ßsleep, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsched); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßscheduler, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µscheduler = πTemp001
							// line 127: now = timer.time()
							πF.SetLineno(127)
							if πE = πg.CheckLocal(πF, µtimer, "timer"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtimer, ßtime, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µnow = πTemp002
							// line 128: event1 = scheduler.enterabs(now + 1, 1, fun, (1,))
							πF.SetLineno(128)
							πTemp003 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µnow, "now"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µnow, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							πTemp003[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfun, "fun"); πE != nil {
								continue
							}
							πTemp003[2] = µfun
							πTemp001 = πg.NewTuple1(πg.NewInt(1).ToObject()).ToObject()
							πTemp003[3] = πTemp001
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µscheduler, ßenterabs, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µevent1 = πTemp002
							// line 129: event2 = scheduler.enterabs(now + 2, 1, fun, (2,))
							πF.SetLineno(129)
							πTemp003 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µnow, "now"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µnow, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							πTemp003[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfun, "fun"); πE != nil {
								continue
							}
							πTemp003[2] = µfun
							πTemp001 = πg.NewTuple1(πg.NewInt(2).ToObject()).ToObject()
							πTemp003[3] = πTemp001
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µscheduler, ßenterabs, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µevent2 = πTemp002
							// line 130: event4 = scheduler.enterabs(now + 4, 1, fun, (4,))
							πF.SetLineno(130)
							πTemp003 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µnow, "now"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µnow, πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							πTemp003[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfun, "fun"); πE != nil {
								continue
							}
							πTemp003[2] = µfun
							πTemp001 = πg.NewTuple1(πg.NewInt(4).ToObject()).ToObject()
							πTemp003[3] = πTemp001
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µscheduler, ßenterabs, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µevent4 = πTemp002
							// line 131: event5 = scheduler.enterabs(now + 5, 1, fun, (5,))
							πF.SetLineno(131)
							πTemp003 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µnow, "now"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µnow, πg.NewInt(5).ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							πTemp003[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfun, "fun"); πE != nil {
								continue
							}
							πTemp003[2] = µfun
							πTemp001 = πg.NewTuple1(πg.NewInt(5).ToObject()).ToObject()
							πTemp003[3] = πTemp001
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µscheduler, ßenterabs, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µevent5 = πTemp002
							// line 132: event3 = scheduler.enterabs(now + 3, 1, fun, (3,))
							πF.SetLineno(132)
							πTemp003 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µnow, "now"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µnow, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							πTemp003[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfun, "fun"); πE != nil {
								continue
							}
							πTemp003[2] = µfun
							πTemp001 = πg.NewTuple1(πg.NewInt(3).ToObject()).ToObject()
							πTemp003[3] = πTemp001
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µscheduler, ßenterabs, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µevent3 = πTemp002
							// line 133: t = threading.Thread(target=scheduler.run)
							πF.SetLineno(133)
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µscheduler, ßrun, nil); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"target", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßThread, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							µt = πTemp001
							// line 134: t.start()
							πF.SetLineno(134)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µt, ßstart, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 135: timer.advance(1)
							πF.SetLineno(135)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µtimer, "timer"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtimer, ßadvance, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 136: self.assertEqual(q.get(timeout=TIMEOUT), 1)
							πF.SetLineno(136)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTIMEOUT); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"timeout", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 137: self.assertTrue(q.empty())
							πF.SetLineno(137)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßempty, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 138: scheduler.cancel(event2)
							πF.SetLineno(138)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µevent2, "event2"); πE != nil {
								continue
							}
							πTemp003[0] = µevent2
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µscheduler, ßcancel, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 139: scheduler.cancel(event5)
							πF.SetLineno(139)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µevent5, "event5"); πE != nil {
								continue
							}
							πTemp003[0] = µevent5
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µscheduler, ßcancel, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 140: timer.advance(1)
							πF.SetLineno(140)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µtimer, "timer"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtimer, ßadvance, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 141: self.assertTrue(q.empty())
							πF.SetLineno(141)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßempty, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 142: timer.advance(1)
							πF.SetLineno(142)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µtimer, "timer"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtimer, ßadvance, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 143: self.assertEqual(q.get(timeout=TIMEOUT), 3)
							πF.SetLineno(143)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTIMEOUT); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"timeout", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 144: self.assertTrue(q.empty())
							πF.SetLineno(144)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßempty, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 145: timer.advance(1)
							πF.SetLineno(145)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µtimer, "timer"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtimer, ßadvance, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 146: self.assertEqual(q.get(timeout=TIMEOUT), 4)
							πF.SetLineno(146)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTIMEOUT); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"timeout", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewInt(4).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 147: self.assertTrue(q.empty())
							πF.SetLineno(147)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßempty, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 148: timer.advance(1000)
							πF.SetLineno(148)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(1000).ToObject()
							if πE = πg.CheckLocal(πF, µtimer, "timer"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtimer, ßadvance, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 149: t.join(timeout=TIMEOUT)
							πF.SetLineno(149)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTIMEOUT); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"timeout", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µt, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							// line 150: self.assertFalse(t.is_alive())
							πF.SetLineno(150)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µt, ßis_alive, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 151: self.assertTrue(q.empty())
							πF.SetLineno(151)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßempty, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 152: self.assertEqual(timer.time(), 4)
							πF.SetLineno(152)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtimer, "timer"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtimer, ßtime, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewInt(4).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_cancel_concurrent.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 121: @unittest.skip('grumpy')
					πF.SetLineno(121)
					πTemp005 = πF.MakeArgs(1)
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßtest_cancel_concurrent); πE != nil {
						continue
					}
					πTemp005[0] = πTemp010
					πTemp007 = πF.MakeArgs(1)
					πTemp007[0] = ßgrumpy.ToObject()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßskip, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp011.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πTemp011, πE = πTemp010.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πClass.SetItem(πF, ßtest_cancel_concurrent.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 154: def test_empty(self):
					πF.SetLineno(154)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("test_empty", "build/src/__python__/test/test_sched.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µl *πg.Object = πg.UnboundLocal; _ = µl
						var µfun *πg.Object = πg.UnboundLocal; _ = µfun
						var µscheduler *πg.Object = πg.UnboundLocal; _ = µscheduler
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var µz *πg.Object = πg.UnboundLocal; _ = µz
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
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
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 155: l = []
							πF.SetLineno(155)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µl = πTemp002
							// line 156: fun = lambda x: l.append(x)
							πF.SetLineno(156)
							πTemp003 = make([]πg.Param, 1)
							πTemp003[0] = πg.Param{Name: "x", Def: nil}
							πTemp002 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/test/test_sched.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µx *πg.Object = πArgs[0]; _ = µx
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
									// line 156: fun = lambda x: l.append(x)
									πF.SetLineno(156)
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
										continue
									}
									πTemp001[0] = µx
									if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µl, ßappend, nil); πE != nil {
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
							µfun = πTemp002
							// line 157: scheduler = sched.scheduler(time.time, time.sleep)
							πF.SetLineno(157)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßtime, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßsleep, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsched); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßscheduler, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µscheduler = πTemp002
							// line 158: self.assertTrue(scheduler.empty())
							πF.SetLineno(158)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µscheduler, ßempty, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = make([]*πg.Object, 5)
							πTemp001[0] = πg.NewFloat(0.05).ToObject()
							πTemp001[1] = πg.NewFloat(0.04).ToObject()
							πTemp001[2] = πg.NewFloat(0.03).ToObject()
							πTemp001[3] = πg.NewFloat(0.02).ToObject()
							πTemp001[4] = πg.NewFloat(0.01).ToObject()
							πTemp004 = πg.NewList(πTemp001...).ToObject()
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
							if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µx = πTemp004
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 160: z = scheduler.enterabs(x, 1, fun, (x,))
							πF.SetLineno(160)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							πTemp001[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfun, "fun"); πE != nil {
								continue
							}
							πTemp001[2] = µfun
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple1(µx).ToObject()
							πTemp001[3] = πTemp004
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µscheduler, ßenterabs, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µz = πTemp007
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 161: self.assertFalse(scheduler.empty())
							πF.SetLineno(161)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µscheduler, ßempty, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 162: scheduler.run()
							πF.SetLineno(162)
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µscheduler, ßrun, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 163: self.assertTrue(scheduler.empty())
							πF.SetLineno(163)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µscheduler, "scheduler"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µscheduler, ßempty, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_empty.ToObject(), πTemp010); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp007, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp007 == nil {
				πTemp007 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp007.Call(πF, []*πg.Object{πg.NewStr("TestCase").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTestCase.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 165: def test_main():
			πF.SetLineno(165)
			πTemp010 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("test_main", "build/src/__python__/test/test_sched.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 166: test.test_support.run_unittest(TestCase)
					πF.SetLineno(166)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTestCase); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtest); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtest_support, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßrun_unittest, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßtest_main.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp007, πE = πg.Eq(πF, πTemp008, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsTrue(πF, πTemp007); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label4
			}
			goto Label5
			// line 168: if __name__ == "__main__":
			πF.SetLineno(168)
		Label4:
			// line 169: test_main()
			πF.SetLineno(169)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßtest_main); πE != nil {
				continue
			}
			if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label5
		Label5:
		}
		return nil, πE
	})
	πg.RegisterModule("test.test_sched", Code)
}
