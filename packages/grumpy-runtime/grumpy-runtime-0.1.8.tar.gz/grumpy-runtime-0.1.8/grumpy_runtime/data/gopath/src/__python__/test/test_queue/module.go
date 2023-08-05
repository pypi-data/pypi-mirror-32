package test_queue
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/test/test_queue.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßBaseQueueTest := πg.InternStr("BaseQueueTest")
		ßBlockingTestMixin := πg.InternStr("BlockingTestMixin")
		ßEmpty := πg.InternStr("Empty")
		ßEvent := πg.InternStr("Event")
		ßException := πg.InternStr("Exception")
		ßFailingQueue := πg.InternStr("FailingQueue")
		ßFailingQueueException := πg.InternStr("FailingQueueException")
		ßFailingQueueTest := πg.InternStr("FailingQueueTest")
		ßFalse := πg.InternStr("False")
		ßFull := πg.InternStr("Full")
		ßLifoQueue := πg.InternStr("LifoQueue")
		ßLifoQueueTest := πg.InternStr("LifoQueueTest")
		ßLock := πg.InternStr("Lock")
		ßNone := πg.InternStr("None")
		ßPriorityQueue := πg.InternStr("PriorityQueue")
		ßPriorityQueueTest := πg.InternStr("PriorityQueueTest")
		ßQUEUE_SIZE := πg.InternStr("QUEUE_SIZE")
		ßQueue := πg.InternStr("Queue")
		ßQueueTest := πg.InternStr("QueueTest")
		ßRuntimeError := πg.InternStr("RuntimeError")
		ßTestCase := πg.InternStr("TestCase")
		ßThread := πg.InternStr("Thread")
		ßTrue := πg.InternStr("True")
		ßValueError := πg.InternStr("ValueError")
		ß_TriggerThread := πg.InternStr("_TriggerThread")
		ß__class__ := πg.InternStr("__class__")
		ß__enter__ := πg.InternStr("__enter__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__init__ := πg.InternStr("__init__")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_get := πg.InternStr("_get")
		ß_put := πg.InternStr("_put")
		ßargs := πg.InternStr("args")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertTrue := πg.InternStr("assertTrue")
		ßcum := πg.InternStr("cum")
		ßcumlock := πg.InternStr("cumlock")
		ßdict := πg.InternStr("dict")
		ßdo_blocking_test := πg.InternStr("do_blocking_test")
		ßdo_exceptional_blocking_test := πg.InternStr("do_exceptional_blocking_test")
		ßempty := πg.InternStr("empty")
		ßfail := πg.InternStr("fail")
		ßfail_next_get := πg.InternStr("fail_next_get")
		ßfail_next_put := πg.InternStr("fail_next_put")
		ßfailing_queue_test := πg.InternStr("failing_queue_test")
		ßfirst := πg.InternStr("first")
		ßfn := πg.InternStr("fn")
		ßfull := πg.InternStr("full")
		ßget := πg.InternStr("get")
		ßis_alive := πg.InternStr("is_alive")
		ßis_set := πg.InternStr("is_set")
		ßjoin := πg.InternStr("join")
		ßlast := πg.InternStr("last")
		ßobject := πg.InternStr("object")
		ßoops := πg.InternStr("oops")
		ßput := πg.InternStr("put")
		ßqueue_join_test := πg.InternStr("queue_join_test")
		ßrange := πg.InternStr("range")
		ßresult := πg.InternStr("result")
		ßrun := πg.InternStr("run")
		ßrun_unittest := πg.InternStr("run_unittest")
		ßset := πg.InternStr("set")
		ßsetUp := πg.InternStr("setUp")
		ßsimple_queue_test := πg.InternStr("simple_queue_test")
		ßsleep := πg.InternStr("sleep")
		ßstart := πg.InternStr("start")
		ßstartedEvent := πg.InternStr("startedEvent")
		ßsum := πg.InternStr("sum")
		ßt := πg.InternStr("t")
		ßtask_done := πg.InternStr("task_done")
		ßtearDown := πg.InternStr("tearDown")
		ßtest_failing_queue := πg.InternStr("test_failing_queue")
		ßtest_main := πg.InternStr("test_main")
		ßtest_queue_join := πg.InternStr("test_queue_join")
		ßtest_queue_task_done := πg.InternStr("test_queue_task_done")
		ßtest_simple_queue := πg.InternStr("test_simple_queue")
		ßtest_support := πg.InternStr("test_support")
		ßthreading := πg.InternStr("threading")
		ßtime := πg.InternStr("time")
		ßtype2test := πg.InternStr("type2test")
		ßunittest := πg.InternStr("unittest")
		ßworker := πg.InternStr("worker")
		ßxrange := πg.InternStr("xrange")
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
		var πTemp006 *πg.Object
		_ = πTemp006
		var πTemp007 []πg.Param
		_ = πTemp007
		var πTemp008 bool
		_ = πTemp008
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 3: import Queue
			πF.SetLineno(3)
			if πTemp002, πE = πg.ImportModule(πF, "Queue"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßQueue.ToObject(), πTemp001); πE != nil {
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
			// line 5: import unittest
			πF.SetLineno(5)
			if πTemp002, πE = πg.ImportModule(πF, "unittest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 6: from test import test_support
			πF.SetLineno(6)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßtest_support.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 8: import threading
			πF.SetLineno(8)
			if πTemp002, πE = πg.ImportModule(πF, "threading"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßthreading.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 10: QUEUE_SIZE = 5
			πF.SetLineno(10)
			if πE = πF.Globals().SetItem(πF, ßQUEUE_SIZE.ToObject(), πg.NewInt(5).ToObject()); πE != nil {
				continue
			}
			// line 13: class _TriggerThread(threading.Thread):
			πF.SetLineno(13)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßThread, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_TriggerThread", "build/src/__python__/test/test_queue.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 14: def __init__(self, fn, args):
					πF.SetLineno(14)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "fn", Def: nil}
					πTemp002[2] = πg.Param{Name: "args", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/test/test_queue.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfn *πg.Object = πArgs[1]; _ = µfn
						var µargs *πg.Object = πArgs[2]; _ = µargs
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
							// line 15: self.fn = fn
							πF.SetLineno(15)
							if πE = πg.CheckLocal(πF, µfn, "fn"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µfn); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfn, πTemp001); πE != nil {
								continue
							}
							// line 16: self.args = args
							πF.SetLineno(16)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µargs); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßargs, πTemp001); πE != nil {
								continue
							}
							// line 17: self.startedEvent = threading.Event()
							πF.SetLineno(17)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßEvent, nil); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßstartedEvent, πTemp002); πE != nil {
								continue
							}
							// line 18: threading.Thread.__init__(self)
							πF.SetLineno(18)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßThread, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 20: def run(self):
					πF.SetLineno(20)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("run", "build/src/__python__/test/test_queue.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 []*πg.Object
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
							// line 29: time.sleep(0.1)
							πF.SetLineno(29)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewFloat(0.1).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsleep, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 30: self.startedEvent.set()
							πF.SetLineno(30)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstartedEvent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 31: self.fn(*self.args)
							πF.SetLineno(31)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßargs, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßfn, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Invoke(πF, πTemp003, nil, πTemp002, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp003); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("_TriggerThread").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_TriggerThread.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 45: class BlockingTestMixin(object):
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
			_, πE = πg.NewCode("BlockingTestMixin", "build/src/__python__/test/test_queue.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 47: def tearDown(self):
					πF.SetLineno(47)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("tearDown", "build/src/__python__/test/test_queue.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 48: self.t = None
							πF.SetLineno(48)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßt, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtearDown.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 50: def do_blocking_test(self, block_func, block_args, trigger_func, trigger_args):
					πF.SetLineno(50)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "block_func", Def: nil}
					πTemp002[2] = πg.Param{Name: "block_args", Def: nil}
					πTemp002[3] = πg.Param{Name: "trigger_func", Def: nil}
					πTemp002[4] = πg.Param{Name: "trigger_args", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("do_blocking_test", "build/src/__python__/test/test_queue.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µblock_func *πg.Object = πArgs[1]; _ = µblock_func
						var µblock_args *πg.Object = πArgs[2]; _ = µblock_args
						var µtrigger_func *πg.Object = πArgs[3]; _ = µtrigger_func
						var µtrigger_args *πg.Object = πArgs[4]; _ = µtrigger_args
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 51: self.t = _TriggerThread(trigger_func, trigger_args)
							πF.SetLineno(51)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtrigger_func, "trigger_func"); πE != nil {
								continue
							}
							πTemp001[0] = µtrigger_func
							if πE = πg.CheckLocal(πF, µtrigger_args, "trigger_args"); πE != nil {
								continue
							}
							πTemp001[1] = µtrigger_args
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_TriggerThread); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßt, πTemp002); πE != nil {
								continue
							}
							// line 52: self.t.start()
							πF.SetLineno(52)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßt, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstart, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 53: self.result = block_func(*block_args)
							πF.SetLineno(53)
							if πE = πg.CheckLocal(πF, µblock_args, "block_args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µblock_func, "block_func"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, µblock_func, nil, µblock_args, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßresult, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßt, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßstartedEvent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßis_set, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 55: if not self.t.startedEvent.is_set():
							πF.SetLineno(55)
						Label1:
							// line 56: self.fail("blocking function '%r' appeared not to block" %
							πF.SetLineno(56)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µblock_func, "block_func"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("blocking function '%r' appeared not to block").ToObject(), µblock_func); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label2
						Label2:
							// line 58: self.t.join(10) # make sure the thread terminates
							πF.SetLineno(58)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(10).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßt, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßt, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßis_alive, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 59: if self.t.is_alive():
							πF.SetLineno(59)
						Label3:
							// line 60: self.fail("trigger function '%r' appeared to not return" %
							πF.SetLineno(60)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtrigger_func, "trigger_func"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("trigger function '%r' appeared to not return").ToObject(), µtrigger_func); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label4
						Label4:
							// line 62: return self.result
							πF.SetLineno(62)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßresult, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdo_blocking_test.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 65: def do_exceptional_blocking_test(self,block_func, block_args, trigger_func,
					πF.SetLineno(65)
					πTemp002 = make([]πg.Param, 6)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "block_func", Def: nil}
					πTemp002[2] = πg.Param{Name: "block_args", Def: nil}
					πTemp002[3] = πg.Param{Name: "trigger_func", Def: nil}
					πTemp002[4] = πg.Param{Name: "trigger_args", Def: nil}
					πTemp002[5] = πg.Param{Name: "expected_exception_class", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("do_exceptional_blocking_test", "build/src/__python__/test/test_queue.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µblock_func *πg.Object = πArgs[1]; _ = µblock_func
						var µblock_args *πg.Object = πArgs[2]; _ = µblock_args
						var µtrigger_func *πg.Object = πArgs[3]; _ = µtrigger_func
						var µtrigger_args *πg.Object = πArgs[4]; _ = µtrigger_args
						var µexpected_exception_class *πg.Object = πArgs[5]; _ = µexpected_exception_class
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
							case 3: goto Label3
							default: panic("unexpected function state")
							}
							// line 67: self.t = _TriggerThread(trigger_func, trigger_args)
							πF.SetLineno(67)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtrigger_func, "trigger_func"); πE != nil {
								continue
							}
							πTemp001[0] = µtrigger_func
							if πE = πg.CheckLocal(πF, µtrigger_args, "trigger_args"); πE != nil {
								continue
							}
							πTemp001[1] = µtrigger_args
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_TriggerThread); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßt, πTemp002); πE != nil {
								continue
							}
							// line 68: self.t.start()
							πF.SetLineno(68)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßt, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstart, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 69: try:
							πF.SetLineno(69)
							πF.PushCheckpoint(1)
							// line 70: try:
							πF.SetLineno(70)
							πF.PushCheckpoint(3)
							// line 71: block_func(*block_args)
							πF.SetLineno(71)
							if πE = πg.CheckLocal(πF, µblock_args, "block_args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µblock_func, "block_func"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, µblock_func, nil, µblock_args, nil, nil); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							// line 75: self.fail("expected exception of kind %r" %
							πF.SetLineno(75)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µexpected_exception_class, "expected_exception_class"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("expected exception of kind %r").ToObject(), µexpected_exception_class); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label2
						Label3:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πE = πg.CheckLocal(πF, µexpected_exception_class, "expected_exception_class"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), µexpected_exception_class); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 72: except expected_exception_class:
							πF.SetLineno(72)
						Label4:
							// line 73: raise
							πF.SetLineno(73)
							πE = πF.Raise(nil, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label2
						Label2:
							πF.PopCheckpoint()
							goto Label1
						Label1:
							πTemp004, πTemp005 = πF.RestoreExc(nil, nil)
							// line 78: self.t.join(10) # make sure the thread terminates
							πF.SetLineno(78)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(10).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßt, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßt, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßis_alive, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label5
							}
							goto Label6
							// line 79: if self.t.is_alive():
							πF.SetLineno(79)
						Label5:
							// line 80: self.fail("trigger function '%r' appeared to not return" %
							πF.SetLineno(80)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtrigger_func, "trigger_func"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("trigger function '%r' appeared to not return").ToObject(), µtrigger_func); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßt, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßstartedEvent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp007, ßis_set, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label7
							}
							goto Label8
							// line 82: if not self.t.startedEvent.is_set():
							πF.SetLineno(82)
						Label7:
							// line 83: self.fail("trigger thread ended but event never set")
							πF.SetLineno(83)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("trigger thread ended but event never set").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label8
						Label8:
							if πTemp004 != nil {
								πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
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
					if πE = πClass.SetItem(πF, ßdo_exceptional_blocking_test.ToObject(), πTemp004); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("BlockingTestMixin").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBlockingTestMixin.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 86: class BaseQueueTest(BlockingTestMixin):
			πF.SetLineno(86)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßBlockingTestMixin); πE != nil {
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
			_, πE = πg.NewCode("BaseQueueTest", "build/src/__python__/test/test_queue.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 87: def setUp(self):
					πF.SetLineno(87)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("setUp", "build/src/__python__/test/test_queue.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 88: self.cum = 0
							πF.SetLineno(88)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcum, πTemp001); πE != nil {
								continue
							}
							// line 89: self.cumlock = threading.Lock()
							πF.SetLineno(89)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßLock, nil); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßcumlock, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsetUp.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 91: def simple_queue_test(self, q):
					πF.SetLineno(91)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "q", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("simple_queue_test", "build/src/__python__/test/test_queue.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µq *πg.Object = πArgs[1]; _ = µq
						var µtarget_order *πg.Object = πg.UnboundLocal; _ = µtarget_order
						var µactual_order *πg.Object = πg.UnboundLocal; _ = µactual_order
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µlast *πg.Object = πg.UnboundLocal; _ = µlast
						var µfull *πg.Object = πg.UnboundLocal; _ = µfull
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 πg.KWArgs
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
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
							case 3: goto Label3
							case 4: goto Label4
							case 7: goto Label7
							case 10: goto Label10
							case 12: goto Label12
							case 13: goto Label13
							case 16: goto Label16
							case 19: goto Label19
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µq, ßempty, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 92: if not q.empty():
							πF.SetLineno(92)
						Label1:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							// line 93: raise RuntimeError, "Call this function with an empty queue"
							πF.SetLineno(93)
							πE = πF.Raise(πTemp001, πg.NewStr("Call this function with an empty queue").ToObject(), nil)
							continue
							goto Label2
						Label2:
							// line 95: q.put(111)
							πF.SetLineno(95)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewInt(111).ToObject()
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 96: q.put(333)
							πF.SetLineno(96)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewInt(333).ToObject()
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 97: q.put(222)
							πF.SetLineno(97)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewInt(222).ToObject()
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 98: target_order = dict(Queue = [111, 333, 222],
							πF.SetLineno(98)
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(111).ToObject()
							πTemp005[1] = πg.NewInt(333).ToObject()
							πTemp005[2] = πg.NewInt(222).ToObject()
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(222).ToObject()
							πTemp005[1] = πg.NewInt(333).ToObject()
							πTemp005[2] = πg.NewInt(111).ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(111).ToObject()
							πTemp005[1] = πg.NewInt(222).ToObject()
							πTemp005[2] = πg.NewInt(333).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp006 = πg.KWArgs{
								{"Queue", πTemp001},
								{"LifoQueue", πTemp002},
								{"PriorityQueue", πTemp003},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, πTemp006); πE != nil {
								continue
							}
							µtarget_order = πTemp002
							// line 101: actual_order = [q.get(), q.get(), q.get()]
							πF.SetLineno(101)
							πTemp005 = make([]*πg.Object, 3)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp002
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[2] = πTemp002
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							µactual_order = πTemp001
							// line 102: self.assertEqual(actual_order, target_order[q.__class__.__name__],
							πF.SetLineno(102)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µactual_order, "actual_order"); πE != nil {
								continue
							}
							πTemp005[0] = µactual_order
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µq, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__name__, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πE = πg.CheckLocal(πF, µtarget_order, "target_order"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µtarget_order, πTemp001); πE != nil {
								continue
							}
							πTemp005[1] = πTemp002
							πTemp005[2] = πg.NewStr("Didn't seem to queue the correct data!").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp005 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßQUEUE_SIZE); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
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
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µi = πTemp002
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 105: q.put(i)
							πF.SetLineno(105)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp005[0] = µi
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 106: self.assertTrue(not q.empty(), "Queue should not be empty")
							πF.SetLineno(106)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µq, ßempty, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp007).ToObject()
							πTemp005[0] = πTemp002
							πTemp005[1] = πg.NewStr("Queue should not be empty").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 107: self.assertTrue(not q.full(), "Queue should not be full")
							πF.SetLineno(107)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µq, ßfull, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							πTemp005[0] = πTemp001
							πTemp005[1] = πg.NewStr("Queue should not be full").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 108: last = 2 * QUEUE_SIZE
							πF.SetLineno(108)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßQUEUE_SIZE); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, πg.NewInt(2).ToObject(), πTemp002); πE != nil {
								continue
							}
							µlast = πTemp001
							// line 109: full = 3 * 2 * QUEUE_SIZE
							πF.SetLineno(109)
							if πTemp002, πE = πg.Mul(πF, πg.NewInt(3).ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßQUEUE_SIZE); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							µfull = πTemp001
							// line 110: q.put(last)
							πF.SetLineno(110)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlast, "last"); πE != nil {
								continue
							}
							πTemp005[0] = µlast
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 111: self.assertTrue(q.full(), "Queue should be full")
							πF.SetLineno(111)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßfull, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							πTemp005[1] = πg.NewStr("Queue should be full").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 112: try:
							πF.SetLineno(112)
							πF.PushCheckpoint(7)
							// line 113: q.put(full, block=0)
							πF.SetLineno(113)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfull, "full"); πE != nil {
								continue
							}
							πTemp005[0] = µfull
							πTemp006 = πg.KWArgs{
								{"block", πg.NewInt(0).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 114: self.fail("Didn't appear to block with a full queue")
							πF.SetLineno(114)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("Didn't appear to block with a full queue").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πF.PopCheckpoint()
							goto Label6
						Label7:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßQueue); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßFull, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 115: except Queue.Full:
							πF.SetLineno(115)
						Label8:
							// line 116: pass
							πF.SetLineno(116)
							πF.RestoreExc(nil, nil)
							goto Label6
						Label6:
							// line 117: try:
							πF.SetLineno(117)
							πF.PushCheckpoint(10)
							// line 118: q.put(full, timeout=0.01)
							πF.SetLineno(118)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfull, "full"); πE != nil {
								continue
							}
							πTemp005[0] = µfull
							πTemp006 = πg.KWArgs{
								{"timeout", πg.NewFloat(0.01).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 119: self.fail("Didn't appear to time-out with a full queue")
							πF.SetLineno(119)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("Didn't appear to time-out with a full queue").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πF.PopCheckpoint()
							goto Label9
						Label10:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßQueue); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßFull, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label11
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 120: except Queue.Full:
							πF.SetLineno(120)
						Label11:
							// line 121: pass
							πF.SetLineno(121)
							πF.RestoreExc(nil, nil)
							goto Label9
						Label9:
							// line 123: self.do_blocking_test(q.put, (full,), q.get, ())
							πF.SetLineno(123)
							πTemp005 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µfull, "full"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple1(µfull).ToObject()
							πTemp005[1] = πTemp001
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							πTemp005[2] = πTemp001
							πTemp001 = πg.NewTuple0().ToObject()
							πTemp005[3] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdo_blocking_test, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 124: self.do_blocking_test(q.put, (full, True, 10), q.get, ())
							πF.SetLineno(124)
							πTemp005 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µfull, "full"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(µfull, πTemp002, πg.NewInt(10).ToObject()).ToObject()
							πTemp005[1] = πTemp001
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							πTemp005[2] = πTemp001
							πTemp001 = πg.NewTuple0().ToObject()
							πTemp005[3] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdo_blocking_test, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp005 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßQUEUE_SIZE); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
								continue
							}
							πF.PushCheckpoint(13)
							πTemp004 = false
						Label12:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label14
							}
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µi = πTemp002
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(12)            
							// line 127: q.get()
							πF.SetLineno(127)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label13:
							if πE != nil || πR != nil {
								continue
							}
						Label14:
							// line 128: self.assertTrue(q.empty(), "Queue should be empty")
							πF.SetLineno(128)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßempty, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							πTemp005[1] = πg.NewStr("Queue should be empty").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 129: try:
							πF.SetLineno(129)
							πF.PushCheckpoint(16)
							// line 130: q.get(block=0)
							πF.SetLineno(130)
							πTemp006 = πg.KWArgs{
								{"block", πg.NewInt(0).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, πTemp006); πE != nil {
								continue
							}
							// line 131: self.fail("Didn't appear to block with an empty queue")
							πF.SetLineno(131)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("Didn't appear to block with an empty queue").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πF.PopCheckpoint()
							goto Label15
						Label16:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßQueue); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßEmpty, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label17
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 132: except Queue.Empty:
							πF.SetLineno(132)
						Label17:
							// line 133: pass
							πF.SetLineno(133)
							πF.RestoreExc(nil, nil)
							goto Label15
						Label15:
							// line 134: try:
							πF.SetLineno(134)
							πF.PushCheckpoint(19)
							// line 135: q.get(timeout=0.01)
							πF.SetLineno(135)
							πTemp006 = πg.KWArgs{
								{"timeout", πg.NewFloat(0.01).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, πTemp006); πE != nil {
								continue
							}
							// line 136: self.fail("Didn't appear to time-out with an empty queue")
							πF.SetLineno(136)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("Didn't appear to time-out with an empty queue").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πF.PopCheckpoint()
							goto Label18
						Label19:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßQueue); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßEmpty, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label20
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 137: except Queue.Empty:
							πF.SetLineno(137)
						Label20:
							// line 138: pass
							πF.SetLineno(138)
							πF.RestoreExc(nil, nil)
							goto Label18
						Label18:
							// line 140: self.do_blocking_test(q.get, (), q.put, ('empty',))
							πF.SetLineno(140)
							πTemp005 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							πTemp001 = πg.NewTuple0().ToObject()
							πTemp005[1] = πTemp001
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							πTemp005[2] = πTemp001
							πTemp001 = πg.NewTuple1(ßempty.ToObject()).ToObject()
							πTemp005[3] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdo_blocking_test, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 141: self.do_blocking_test(q.get, (True, 10), q.put, ('empty',))
							πF.SetLineno(141)
							πTemp005 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πg.NewInt(10).ToObject()).ToObject()
							πTemp005[1] = πTemp001
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							πTemp005[2] = πTemp001
							πTemp001 = πg.NewTuple1(ßempty.ToObject()).ToObject()
							πTemp005[3] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdo_blocking_test, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßsimple_queue_test.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 144: def worker(self, q):
					πF.SetLineno(144)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "q", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("worker", "build/src/__python__/test/test_queue.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µq *πg.Object = πArgs[1]; _ = µq
						var µx *πg.Object = πg.UnboundLocal; _ = µx
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
						var πTemp007 *πg.Object
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
							case 6: goto Label6
							default: panic("unexpected function state")
							}
							// line 145: while True:
							πF.SetLineno(145)
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
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πE != nil || !πTemp002 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 146: x = q.get()
							πF.SetLineno(146)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µx = πTemp004
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µx == πTemp004).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							goto Label5
							// line 147: if x is None:
							πF.SetLineno(147)
						Label4:
							// line 148: q.task_done()
							πF.SetLineno(148)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µq, ßtask_done, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 149: return
							πF.SetLineno(149)
							πR = πg.None
							continue
							goto Label5
						Label5:
							// line 150: with self.cumlock:
							πF.SetLineno(150)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcumlock, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp005.Call(πF, πg.Args{πTemp003}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(6)
							// line 151: self.cum += x
							πF.SetLineno(151)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßcum, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IAdd(πF, πTemp006, µx); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcum, πTemp007); πE != nil {
								continue
							}
							πF.PopCheckpoint()
						Label6:
							πTemp008, πTemp009 = nil, nil
							if πE != nil {
								πTemp008, πTemp009 = πF.ExcInfo()
							}
							if πTemp008 != nil {
								πTemp010 = πTemp008.Type()
								if πTemp006, πE = πTemp004.Call(πF, πg.Args{πTemp003, πTemp010.ToObject(), πTemp008.ToObject(), πTemp009.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp006, πE = πTemp004.Call(πF, πg.Args{πTemp003, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp008 != nil && πTemp002 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
								continue
							}
							// line 152: q.task_done()
							πF.SetLineno(152)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µq, ßtask_done, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßworker.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 154: def queue_join_test(self, q):
					πF.SetLineno(154)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "q", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("queue_join_test", "build/src/__python__/test/test_queue.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µq *πg.Object = πArgs[1]; _ = µq
						var µi *πg.Object = πg.UnboundLocal; _ = µi
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
						var πTemp006 πg.KWArgs
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							case 4: goto Label4
							case 5: goto Label5
							case 7: goto Label7
							case 8: goto Label8
							default: panic("unexpected function state")
							}
							// line 155: self.cum = 0
							πF.SetLineno(155)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcum, πTemp001); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject()).ToObject()
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
								continue
							}
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
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp004 = !isStop
							} else {
								πTemp004 = true
								µi = πTemp002
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 157: threading.Thread(target=self.worker, args=(q,)).start()
							πF.SetLineno(157)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßworker, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple1(µq).ToObject()
							πTemp006 = πg.KWArgs{
								{"target", πTemp002},
								{"args", πTemp005},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßThread, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, nil, πTemp006); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßstart, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = πg.NewInt(100).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp003 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
								πF.PopCheckpoint()
								goto Label6
							}
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp004 = !isStop
							} else {
								πTemp004 = true
								µi = πTemp002
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 159: q.put(i)
							πF.SetLineno(159)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp007[0] = µi
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 160: q.join()
							πF.SetLineno(160)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 161: self.assertEqual(self.cum, sum(range(100)),
							πF.SetLineno(161)
							πTemp007 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcum, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							πTemp008 = πF.MakeArgs(1)
							πTemp009 = πF.MakeArgs(1)
							πTemp009[0] = πg.NewInt(100).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp008[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsum); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp007[1] = πTemp002
							πTemp007[2] = πg.NewStr("q.join() did not block until all tasks were done").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp002 = πg.NewTuple2(πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject()).ToObject()
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
								continue
							}
							πF.PushCheckpoint(8)
							πTemp003 = false
						Label7:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
								πF.PopCheckpoint()
								goto Label9
							}
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp004 = !isStop
							} else {
								πTemp004 = true
								µi = πTemp002
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(7)            
							// line 164: q.put(None)         # instruct the threads to close
							πF.SetLineno(164)
							πTemp007 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							continue
						Label8:
							if πE != nil || πR != nil {
								continue
							}
						Label9:
							// line 165: q.join()                # verify that you can join twice
							πF.SetLineno(165)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßjoin, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßqueue_join_test.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 167: def test_queue_task_done(self):
					πF.SetLineno(167)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_queue_task_done", "build/src/__python__/test/test_queue.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µq *πg.Object = πg.UnboundLocal; _ = µq
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 169: q = self.type2test()
							πF.SetLineno(169)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µq = πTemp002
							// line 170: try:
							πF.SetLineno(170)
							πF.PushCheckpoint(2)
							// line 171: q.task_done()
							πF.SetLineno(171)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßtask_done, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							// line 175: self.fail("Did not detect task count going negative")
							πF.SetLineno(175)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("Did not detect task count going negative").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 172: except ValueError:
							πF.SetLineno(172)
						Label3:
							// line 173: pass
							πF.SetLineno(173)
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_queue_task_done.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 177: def test_queue_join(self):
					πF.SetLineno(177)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_queue_join", "build/src/__python__/test/test_queue.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µq *πg.Object = πg.UnboundLocal; _ = µq
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 180: q = self.type2test()
							πF.SetLineno(180)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µq = πTemp002
							// line 181: self.queue_join_test(q)
							πF.SetLineno(181)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							πTemp003[0] = µq
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßqueue_join_test, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 182: self.queue_join_test(q)
							πF.SetLineno(182)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							πTemp003[0] = µq
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßqueue_join_test, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 183: try:
							πF.SetLineno(183)
							πF.PushCheckpoint(2)
							// line 184: q.task_done()
							πF.SetLineno(184)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßtask_done, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							// line 188: self.fail("Did not detect task count going negative")
							πF.SetLineno(188)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("Did not detect task count going negative").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 185: except ValueError:
							πF.SetLineno(185)
						Label3:
							// line 186: pass
							πF.SetLineno(186)
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_queue_join.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 190: def test_simple_queue(self):
					πF.SetLineno(190)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("test_simple_queue", "build/src/__python__/test/test_queue.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µq *πg.Object = πg.UnboundLocal; _ = µq
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
							// line 193: q = self.type2test(QUEUE_SIZE)
							πF.SetLineno(193)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßQUEUE_SIZE); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µq = πTemp003
							// line 194: self.simple_queue_test(q)
							πF.SetLineno(194)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							πTemp001[0] = µq
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsimple_queue_test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 195: self.simple_queue_test(q)
							πF.SetLineno(195)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							πTemp001[0] = µq
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsimple_queue_test, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_simple_queue.ToObject(), πTemp008); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("BaseQueueTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBaseQueueTest.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 198: class QueueTest(BaseQueueTest, unittest.TestCase):
			πF.SetLineno(198)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßBaseQueueTest); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßTestCase, nil); πE != nil {
				continue
			}
			πTemp002[1] = πTemp006
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("QueueTest", "build/src/__python__/test/test_queue.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 199: type2test = Queue.Queue
					πF.SetLineno(199)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßQueue); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßQueue, nil); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßtype2test.ToObject(), πTemp002); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("QueueTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßQueueTest.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 201: class LifoQueueTest(BaseQueueTest, unittest.TestCase):
			πF.SetLineno(201)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßBaseQueueTest); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßTestCase, nil); πE != nil {
				continue
			}
			πTemp002[1] = πTemp006
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("LifoQueueTest", "build/src/__python__/test/test_queue.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 202: type2test = Queue.LifoQueue
					πF.SetLineno(202)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßQueue); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßLifoQueue, nil); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßtype2test.ToObject(), πTemp002); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("LifoQueueTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLifoQueueTest.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 204: class PriorityQueueTest(BaseQueueTest, unittest.TestCase):
			πF.SetLineno(204)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßBaseQueueTest); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßTestCase, nil); πE != nil {
				continue
			}
			πTemp002[1] = πTemp006
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("PriorityQueueTest", "build/src/__python__/test/test_queue.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 205: type2test = Queue.PriorityQueue
					πF.SetLineno(205)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßQueue); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßPriorityQueue, nil); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßtype2test.ToObject(), πTemp002); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("PriorityQueueTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPriorityQueueTest.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 210: class FailingQueueException(Exception):
			πF.SetLineno(210)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
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
			_, πE = πg.NewCode("FailingQueueException", "build/src/__python__/test/test_queue.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 211: pass
					πF.SetLineno(211)
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("FailingQueueException").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFailingQueueException.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 213: class FailingQueue(Queue.Queue):
			πF.SetLineno(213)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßQueue); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßQueue, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("FailingQueue", "build/src/__python__/test/test_queue.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 214: def __init__(self, *args):
					πF.SetLineno(214)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/test/test_queue.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µargs *πg.Object = πArgs[1]; _ = µargs
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
							// line 215: self.fail_next_put = False
							πF.SetLineno(215)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfail_next_put, πTemp002); πE != nil {
								continue
							}
							// line 216: self.fail_next_get = False
							πF.SetLineno(216)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfail_next_get, πTemp002); πE != nil {
								continue
							}
							// line 217: Queue.Queue.__init__(self, *args)
							πF.SetLineno(217)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßQueue); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßQueue, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, πTemp001, πTemp003, µargs, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 218: def _put(self, item):
					πF.SetLineno(218)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "item", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("_put", "build/src/__python__/test/test_queue.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µitem *πg.Object = πArgs[1]; _ = µitem
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
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail_next_put, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 219: if self.fail_next_put:
							πF.SetLineno(219)
						Label1:
							// line 220: self.fail_next_put = False
							πF.SetLineno(220)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfail_next_put, πTemp003); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFailingQueueException); πE != nil {
								continue
							}
							// line 221: raise FailingQueueException, "You Lose"
							πF.SetLineno(221)
							πE = πF.Raise(πTemp001, πg.NewStr("You Lose").ToObject(), nil)
							continue
							goto Label2
						Label2:
							// line 222: return Queue.Queue._put(self, item)
							πF.SetLineno(222)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp004[1] = µitem
							if πTemp001, πE = πg.ResolveGlobal(πF, ßQueue); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßQueue, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ß_put, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ß_put.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 223: def _get(self):
					πF.SetLineno(223)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("_get", "build/src/__python__/test/test_queue.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail_next_get, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 224: if self.fail_next_get:
							πF.SetLineno(224)
						Label1:
							// line 225: self.fail_next_get = False
							πF.SetLineno(225)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfail_next_get, πTemp003); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFailingQueueException); πE != nil {
								continue
							}
							// line 226: raise FailingQueueException, "You Lose"
							πF.SetLineno(226)
							πE = πF.Raise(πTemp001, πg.NewStr("You Lose").ToObject(), nil)
							continue
							goto Label2
						Label2:
							// line 227: return Queue.Queue._get(self)
							πF.SetLineno(227)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßQueue); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßQueue, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ß_get, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ß_get.ToObject(), πTemp004); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("FailingQueue").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFailingQueue.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 229: class FailingQueueTest(BlockingTestMixin, unittest.TestCase):
			πF.SetLineno(229)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßBlockingTestMixin); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßTestCase, nil); πE != nil {
				continue
			}
			πTemp002[1] = πTemp006
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("FailingQueueTest", "build/src/__python__/test/test_queue.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 231: def failing_queue_test(self, q):
					πF.SetLineno(231)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "q", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("failing_queue_test", "build/src/__python__/test/test_queue.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µq *πg.Object = πArgs[1]; _ = µq
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 πg.KWArgs
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
							case 3: goto Label3
							case 4: goto Label4
							case 7: goto Label7
							case 10: goto Label10
							case 13: goto Label13
							case 16: goto Label16
							case 18: goto Label18
							case 19: goto Label19
							case 22: goto Label22
							case 25: goto Label25
							case 28: goto Label28
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µq, ßempty, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 232: if not q.empty():
							πF.SetLineno(232)
						Label1:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							// line 233: raise RuntimeError, "Call this function with an empty queue"
							πF.SetLineno(233)
							πE = πF.Raise(πTemp001, πg.NewStr("Call this function with an empty queue").ToObject(), nil)
							continue
							goto Label2
						Label2:
							πTemp005 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßQUEUE_SIZE); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								µi = πTemp002
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 235: q.put(i)
							πF.SetLineno(235)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp005[0] = µi
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 237: q.fail_next_put = True
							πF.SetLineno(237)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µq, ßfail_next_put, πTemp002); πE != nil {
								continue
							}
							// line 238: try:
							πF.SetLineno(238)
							πF.PushCheckpoint(7)
							// line 239: q.put("oops", block=0)
							πF.SetLineno(239)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = ßoops.ToObject()
							πTemp007 = πg.KWArgs{
								{"block", πg.NewInt(0).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, πTemp007); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 240: self.fail("The queue didn't fail when it should have")
							πF.SetLineno(240)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("The queue didn't fail when it should have").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πF.PopCheckpoint()
							goto Label6
						Label7:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFailingQueueException); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 241: except FailingQueueException:
							πF.SetLineno(241)
						Label8:
							// line 242: pass
							πF.SetLineno(242)
							πF.RestoreExc(nil, nil)
							goto Label6
						Label6:
							// line 243: q.fail_next_put = True
							πF.SetLineno(243)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µq, ßfail_next_put, πTemp002); πE != nil {
								continue
							}
							// line 244: try:
							πF.SetLineno(244)
							πF.PushCheckpoint(10)
							// line 245: q.put("oops", timeout=0.1)
							πF.SetLineno(245)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = ßoops.ToObject()
							πTemp007 = πg.KWArgs{
								{"timeout", πg.NewFloat(0.1).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, πTemp007); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 246: self.fail("The queue didn't fail when it should have")
							πF.SetLineno(246)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("The queue didn't fail when it should have").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πF.PopCheckpoint()
							goto Label9
						Label10:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFailingQueueException); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label11
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 247: except FailingQueueException:
							πF.SetLineno(247)
						Label11:
							// line 248: pass
							πF.SetLineno(248)
							πF.RestoreExc(nil, nil)
							goto Label9
						Label9:
							// line 249: q.put("last")
							πF.SetLineno(249)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = ßlast.ToObject()
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 250: self.assertTrue(q.full(), "Queue should be full")
							πF.SetLineno(250)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßfull, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							πTemp005[1] = πg.NewStr("Queue should be full").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 252: q.fail_next_put = True
							πF.SetLineno(252)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µq, ßfail_next_put, πTemp002); πE != nil {
								continue
							}
							// line 253: try:
							πF.SetLineno(253)
							πF.PushCheckpoint(13)
							// line 254: self.do_blocking_test(q.put, ("full",), q.get, ())
							πF.SetLineno(254)
							πTemp005 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							πTemp001 = πg.NewTuple1(ßfull.ToObject()).ToObject()
							πTemp005[1] = πTemp001
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							πTemp005[2] = πTemp001
							πTemp001 = πg.NewTuple0().ToObject()
							πTemp005[3] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdo_blocking_test, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 255: self.fail("The queue didn't fail when it should have")
							πF.SetLineno(255)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("The queue didn't fail when it should have").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πF.PopCheckpoint()
							goto Label12
						Label13:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFailingQueueException); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label14
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 256: except FailingQueueException:
							πF.SetLineno(256)
						Label14:
							// line 257: pass
							πF.SetLineno(257)
							πF.RestoreExc(nil, nil)
							goto Label12
						Label12:
							// line 260: q.put("last")
							πF.SetLineno(260)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = ßlast.ToObject()
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 262: q.fail_next_put = True
							πF.SetLineno(262)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µq, ßfail_next_put, πTemp002); πE != nil {
								continue
							}
							// line 263: try:
							πF.SetLineno(263)
							πF.PushCheckpoint(16)
							// line 264: self.do_exceptional_blocking_test(q.put, ("full", True, 10), q.get, (),
							πF.SetLineno(264)
							πTemp005 = πF.MakeArgs(5)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(ßfull.ToObject(), πTemp002, πg.NewInt(10).ToObject()).ToObject()
							πTemp005[1] = πTemp001
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							πTemp005[2] = πTemp001
							πTemp001 = πg.NewTuple0().ToObject()
							πTemp005[3] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFailingQueueException); πE != nil {
								continue
							}
							πTemp005[4] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdo_exceptional_blocking_test, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 266: self.fail("The queue didn't fail when it should have")
							πF.SetLineno(266)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("The queue didn't fail when it should have").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πF.PopCheckpoint()
							goto Label15
						Label16:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFailingQueueException); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label17
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 267: except FailingQueueException:
							πF.SetLineno(267)
						Label17:
							// line 268: pass
							πF.SetLineno(268)
							πF.RestoreExc(nil, nil)
							goto Label15
						Label15:
							// line 271: q.put("last")
							πF.SetLineno(271)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = ßlast.ToObject()
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 272: self.assertTrue(q.full(), "Queue should be full")
							πF.SetLineno(272)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßfull, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							πTemp005[1] = πg.NewStr("Queue should be full").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 273: q.get()
							πF.SetLineno(273)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 274: self.assertTrue(not q.full(), "Queue should not be full")
							πF.SetLineno(274)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µq, ßfull, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							πTemp005[0] = πTemp001
							πTemp005[1] = πg.NewStr("Queue should not be full").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 275: q.put("last")
							πF.SetLineno(275)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = ßlast.ToObject()
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 276: self.assertTrue(q.full(), "Queue should be full")
							πF.SetLineno(276)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßfull, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							πTemp005[1] = πg.NewStr("Queue should be full").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 278: self.do_blocking_test(q.put, ("full",), q.get, ())
							πF.SetLineno(278)
							πTemp005 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							πTemp001 = πg.NewTuple1(ßfull.ToObject()).ToObject()
							πTemp005[1] = πTemp001
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							πTemp005[2] = πTemp001
							πTemp001 = πg.NewTuple0().ToObject()
							πTemp005[3] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdo_blocking_test, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp005 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßQUEUE_SIZE); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
								continue
							}
							πF.PushCheckpoint(19)
							πTemp004 = false
						Label18:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label20
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
								µi = πTemp002
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(18)            
							// line 281: q.get()
							πF.SetLineno(281)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label19:
							if πE != nil || πR != nil {
								continue
							}
						Label20:
							// line 282: self.assertTrue(q.empty(), "Queue should be empty")
							πF.SetLineno(282)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßempty, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							πTemp005[1] = πg.NewStr("Queue should be empty").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 283: q.put("first")
							πF.SetLineno(283)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = ßfirst.ToObject()
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 284: q.fail_next_get = True
							πF.SetLineno(284)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µq, ßfail_next_get, πTemp002); πE != nil {
								continue
							}
							// line 285: try:
							πF.SetLineno(285)
							πF.PushCheckpoint(22)
							// line 286: q.get()
							πF.SetLineno(286)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 287: self.fail("The queue didn't fail when it should have")
							πF.SetLineno(287)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("The queue didn't fail when it should have").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πF.PopCheckpoint()
							goto Label21
						Label22:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFailingQueueException); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label23
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 288: except FailingQueueException:
							πF.SetLineno(288)
						Label23:
							// line 289: pass
							πF.SetLineno(289)
							πF.RestoreExc(nil, nil)
							goto Label21
						Label21:
							// line 290: self.assertTrue(not q.empty(), "Queue should not be empty")
							πF.SetLineno(290)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µq, ßempty, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							πTemp005[0] = πTemp001
							πTemp005[1] = πg.NewStr("Queue should not be empty").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 291: q.fail_next_get = True
							πF.SetLineno(291)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µq, ßfail_next_get, πTemp002); πE != nil {
								continue
							}
							// line 292: try:
							πF.SetLineno(292)
							πF.PushCheckpoint(25)
							// line 293: q.get(timeout=0.1)
							πF.SetLineno(293)
							πTemp007 = πg.KWArgs{
								{"timeout", πg.NewFloat(0.1).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, πTemp007); πE != nil {
								continue
							}
							// line 294: self.fail("The queue didn't fail when it should have")
							πF.SetLineno(294)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("The queue didn't fail when it should have").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πF.PopCheckpoint()
							goto Label24
						Label25:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFailingQueueException); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label26
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 295: except FailingQueueException:
							πF.SetLineno(295)
						Label26:
							// line 296: pass
							πF.SetLineno(296)
							πF.RestoreExc(nil, nil)
							goto Label24
						Label24:
							// line 297: self.assertTrue(not q.empty(), "Queue should not be empty")
							πF.SetLineno(297)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µq, ßempty, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							πTemp005[0] = πTemp001
							πTemp005[1] = πg.NewStr("Queue should not be empty").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 298: q.get()
							πF.SetLineno(298)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 299: self.assertTrue(q.empty(), "Queue should be empty")
							πF.SetLineno(299)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßempty, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							πTemp005[1] = πg.NewStr("Queue should be empty").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 300: q.fail_next_get = True
							πF.SetLineno(300)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µq, ßfail_next_get, πTemp002); πE != nil {
								continue
							}
							// line 301: try:
							πF.SetLineno(301)
							πF.PushCheckpoint(28)
							// line 302: self.do_exceptional_blocking_test(q.get, (), q.put, ('empty',),
							πF.SetLineno(302)
							πTemp005 = πF.MakeArgs(5)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							πTemp001 = πg.NewTuple0().ToObject()
							πTemp005[1] = πTemp001
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßput, nil); πE != nil {
								continue
							}
							πTemp005[2] = πTemp001
							πTemp001 = πg.NewTuple1(ßempty.ToObject()).ToObject()
							πTemp005[3] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFailingQueueException); πE != nil {
								continue
							}
							πTemp005[4] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdo_exceptional_blocking_test, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 304: self.fail("The queue didn't fail when it should have")
							πF.SetLineno(304)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("The queue didn't fail when it should have").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πF.PopCheckpoint()
							goto Label27
						Label28:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFailingQueueException); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label29
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 305: except FailingQueueException:
							πF.SetLineno(305)
						Label29:
							// line 306: pass
							πF.SetLineno(306)
							πF.RestoreExc(nil, nil)
							goto Label27
						Label27:
							// line 308: self.assertTrue(not q.empty(), "Queue should not be empty")
							πF.SetLineno(308)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µq, ßempty, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							πTemp005[0] = πTemp001
							πTemp005[1] = πg.NewStr("Queue should not be empty").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 309: q.get()
							πF.SetLineno(309)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 310: self.assertTrue(q.empty(), "Queue should be empty")
							πF.SetLineno(310)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µq, ßempty, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							πTemp005[1] = πg.NewStr("Queue should be empty").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßfailing_queue_test.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 312: def test_failing_queue(self):
					πF.SetLineno(312)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_failing_queue", "build/src/__python__/test/test_queue.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µq *πg.Object = πg.UnboundLocal; _ = µq
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
							// line 315: q = FailingQueue(QUEUE_SIZE)
							πF.SetLineno(315)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßQUEUE_SIZE); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFailingQueue); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µq = πTemp003
							// line 316: self.failing_queue_test(q)
							πF.SetLineno(316)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							πTemp001[0] = µq
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfailing_queue_test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 317: self.failing_queue_test(q)
							πF.SetLineno(317)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							πTemp001[0] = µq
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfailing_queue_test, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_failing_queue.ToObject(), πTemp003); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("FailingQueueTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFailingQueueTest.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 320: def test_main():
			πF.SetLineno(320)
			πTemp007 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("test_main", "build/src/__python__/test/test_queue.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 321: test_support.run_unittest(QueueTest, LifoQueueTest, PriorityQueueTest,
					πF.SetLineno(321)
					πTemp001 = πF.MakeArgs(4)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßQueueTest); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßLifoQueueTest); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßPriorityQueueTest); πE != nil {
						continue
					}
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFailingQueueTest); πE != nil {
						continue
					}
					πTemp001[3] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrun_unittest, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßtest_main.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp004, πE = πg.Eq(πF, πTemp005, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
				continue
			}
			if πTemp008 {
				goto Label1
			}
			goto Label2
			// line 325: if __name__ == "__main__":
			πF.SetLineno(325)
		Label1:
			// line 326: test_main()
			πF.SetLineno(326)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßtest_main); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("test.test_queue", Code)
}
