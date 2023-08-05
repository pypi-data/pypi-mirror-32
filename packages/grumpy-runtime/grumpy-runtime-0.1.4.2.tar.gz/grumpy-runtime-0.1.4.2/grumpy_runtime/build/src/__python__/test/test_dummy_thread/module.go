package test_dummy_thread
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/test/test_dummy_thread.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßDELAY := πg.InternStr("DELAY")
		ßFalse := πg.InternStr("False")
		ßKeyboardInterrupt := πg.InternStr("KeyboardInterrupt")
		ßLockTests := πg.InternStr("LockTests")
		ßLockType := πg.InternStr("LockType")
		ßMiscTests := πg.InternStr("MiscTests")
		ßNone := πg.InternStr("None")
		ßQueue := πg.InternStr("Queue")
		ßSystemExit := πg.InternStr("SystemExit")
		ßTestCase := πg.InternStr("TestCase")
		ßThreadTests := πg.InternStr("ThreadTests")
		ßTrue := πg.InternStr("True")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_thread := πg.InternStr("_thread")
		ßacquire := πg.InternStr("acquire")
		ßallocate_lock := πg.InternStr("allocate_lock")
		ßarg1 := πg.InternStr("arg1")
		ßarg2 := πg.InternStr("arg2")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertFalse := πg.InternStr("assertFalse")
		ßassertGreaterEqual := πg.InternStr("assertGreaterEqual")
		ßassertIs := πg.InternStr("assertIs")
		ßassertIsInstance := πg.InternStr("assertIsInstance")
		ßassertNotEqual := πg.InternStr("assertNotEqual")
		ßassertRaises := πg.InternStr("assertRaises")
		ßassertTrue := πg.InternStr("assertTrue")
		ßdone := πg.InternStr("done")
		ßerror := πg.InternStr("error")
		ßexit := πg.InternStr("exit")
		ßget := πg.InternStr("get")
		ßget_ident := πg.InternStr("get_ident")
		ßint := πg.InternStr("int")
		ßinterrupt_main := πg.InternStr("interrupt_main")
		ßlock := πg.InternStr("lock")
		ßlocked := πg.InternStr("locked")
		ßput := πg.InternStr("put")
		ßqsize := πg.InternStr("qsize")
		ßqueue := πg.InternStr("queue")
		ßrandom := πg.InternStr("random")
		ßrelease := πg.InternStr("release")
		ßround := πg.InternStr("round")
		ßrun_unittest := πg.InternStr("run_unittest")
		ßsetUp := πg.InternStr("setUp")
		ßsleep := πg.InternStr("sleep")
		ßstart_new_thread := πg.InternStr("start_new_thread")
		ßtest_LockType := πg.InternStr("test_LockType")
		ßtest_arg_passing := πg.InternStr("test_arg_passing")
		ßtest_cond_acquire_fail := πg.InternStr("test_cond_acquire_fail")
		ßtest_cond_acquire_success := πg.InternStr("test_cond_acquire_success")
		ßtest_exit := πg.InternStr("test_exit")
		ßtest_ident := πg.InternStr("test_ident")
		ßtest_improper_release := πg.InternStr("test_improper_release")
		ßtest_initlock := πg.InternStr("test_initlock")
		ßtest_interrupt_in_main := πg.InternStr("test_interrupt_in_main")
		ßtest_interrupt_main := πg.InternStr("test_interrupt_main")
		ßtest_main := πg.InternStr("test_main")
		ßtest_multi_creation := πg.InternStr("test_multi_creation")
		ßtest_release := πg.InternStr("test_release")
		ßtest_support := πg.InternStr("test_support")
		ßtest_uncond_acquire_blocking := πg.InternStr("test_uncond_acquire_blocking")
		ßtest_uncond_acquire_return_val := πg.InternStr("test_uncond_acquire_return_val")
		ßtest_uncond_acquire_success := πg.InternStr("test_uncond_acquire_success")
		ßtime := πg.InternStr("time")
		ßtuple := πg.InternStr("tuple")
		ßunittest := πg.InternStr("unittest")
		ßverbose := πg.InternStr("verbose")
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
			// line 1: """Generic thread tests.
			πF.SetLineno(1)
			// line 8: import dummy_thread as _thread
			πF.SetLineno(8)
			if πTemp002, πE = πg.ImportModule(πF, "dummy_thread"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_thread.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 9: import time
			πF.SetLineno(9)
			if πTemp002, πE = πg.ImportModule(πF, "time"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtime.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 10: import Queue
			πF.SetLineno(10)
			if πTemp002, πE = πg.ImportModule(πF, "Queue"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßQueue.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 11: import random
			πF.SetLineno(11)
			if πTemp002, πE = πg.ImportModule(πF, "random"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßrandom.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 12: import unittest
			πF.SetLineno(12)
			if πTemp002, πE = πg.ImportModule(πF, "unittest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 13: from test import test_support
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßtest_support.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 15: DELAY = 0 # Set > 0 when testing a module other than dummy_thread, such as
			πF.SetLineno(15)
			if πE = πF.Globals().SetItem(πF, ßDELAY.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
				continue
			}
			// line 18: class LockTests(unittest.TestCase):
			πF.SetLineno(18)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßTestCase, nil); πE != nil {
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
			_, πE = πg.NewCode("LockTests", "build/src/__python__/test/test_dummy_thread.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 19: """Test lock objects."""
					πF.SetLineno(19)
					// line 21: def setUp(self):
					πF.SetLineno(21)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("setUp", "build/src/__python__/test/test_dummy_thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 23: self.lock = _thread.allocate_lock()
							πF.SetLineno(23)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_thread); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßallocate_lock, nil); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßlock, πTemp002); πE != nil {
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
					// line 25: def test_initlock(self):
					πF.SetLineno(25)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_initlock", "build/src/__python__/test/test_dummy_thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 27: self.assertFalse(self.lock.locked(),
							πF.SetLineno(27)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlock, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlocked, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("Lock object is not initialized unlocked.").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_initlock.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 30: def test_release(self):
					πF.SetLineno(30)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_release", "build/src/__python__/test/test_dummy_thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 32: self.lock.acquire()
							πF.SetLineno(32)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlock, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 33: self.lock.release()
							πF.SetLineno(33)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlock, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 34: self.assertFalse(self.lock.locked(),
							πF.SetLineno(34)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlock, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßlocked, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							πTemp003[1] = πg.NewStr("Lock object did not release properly.").ToObject()
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_release.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 37: def test_improper_release(self):
					πF.SetLineno(37)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_improper_release", "build/src/__python__/test/test_dummy_thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 39: self.assertRaises(_thread.error, self.lock.release)
							πF.SetLineno(39)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_thread); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßerror, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlock, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrelease, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_improper_release.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 41: def test_cond_acquire_success(self):
					πF.SetLineno(41)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_cond_acquire_success", "build/src/__python__/test/test_dummy_thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 43: self.assertTrue(self.lock.acquire(0),
							πF.SetLineno(43)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßlock, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr("Conditional acquiring of the lock failed.").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_cond_acquire_success.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 46: def test_cond_acquire_fail(self):
					πF.SetLineno(46)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_cond_acquire_fail", "build/src/__python__/test/test_dummy_thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
							// line 48: self.lock.acquire(0)
							πF.SetLineno(48)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlock, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 49: self.assertFalse(self.lock.acquire(0),
							πF.SetLineno(49)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlock, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("Conditional acquiring of a locked lock incorrectly succeeded.").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_cond_acquire_fail.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 53: def test_uncond_acquire_success(self):
					πF.SetLineno(53)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("test_uncond_acquire_success", "build/src/__python__/test/test_dummy_thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 55: self.lock.acquire()
							πF.SetLineno(55)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlock, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 56: self.assertTrue(self.lock.locked(),
							πF.SetLineno(56)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlock, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßlocked, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							πTemp003[1] = πg.NewStr("Uncondional locking failed.").ToObject()
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_uncond_acquire_success.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 59: def test_uncond_acquire_return_val(self):
					πF.SetLineno(59)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("test_uncond_acquire_return_val", "build/src/__python__/test/test_dummy_thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 61: self.assertIs(self.lock.acquire(1), True,
							πF.SetLineno(61)
							πTemp001 = πF.MakeArgs(3)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßlock, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewStr("Unconditional locking did not return True.").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertIs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 63: self.assertIs(self.lock.acquire(), True)
							πF.SetLineno(63)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßlock, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertIs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_uncond_acquire_return_val.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 65: def test_uncond_acquire_blocking(self):
					πF.SetLineno(65)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("test_uncond_acquire_blocking", "build/src/__python__/test/test_dummy_thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdelay_unlock *πg.Object = πg.UnboundLocal; _ = µdelay_unlock
						var µstart_time *πg.Object = πg.UnboundLocal; _ = µstart_time
						var µend_time *πg.Object = πg.UnboundLocal; _ = µend_time
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
						var πTemp007 bool
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 67: def delay_unlock(to_unlock, delay):
							πF.SetLineno(67)
							πTemp002 = make([]πg.Param, 2)
							πTemp002[0] = πg.Param{Name: "to_unlock", Def: nil}
							πTemp002[1] = πg.Param{Name: "delay", Def: nil}
							πTemp001 = πg.NewFunction(πg.NewCode("delay_unlock", "build/src/__python__/test/test_dummy_thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µto_unlock *πg.Object = πArgs[0]; _ = µto_unlock
								var µdelay *πg.Object = πArgs[1]; _ = µdelay
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
									// line 68: """Hold on to lock for a set amount of time before unlocking."""
									πF.SetLineno(68)
									// line 69: time.sleep(delay)
									πF.SetLineno(69)
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µdelay, "delay"); πE != nil {
										continue
									}
									πTemp001[0] = µdelay
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
									// line 70: to_unlock.release()
									πF.SetLineno(70)
									if πE = πg.CheckLocal(πF, µto_unlock, "to_unlock"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µto_unlock, ßrelease, nil); πE != nil {
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
							µdelay_unlock = πTemp001
							// line 72: self.lock.acquire()
							πF.SetLineno(72)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßlock, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 73: start_time = int(time.time())
							πF.SetLineno(73)
							πTemp005 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßtime, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µstart_time = πTemp004
							// line 74: _thread.start_new_thread(delay_unlock,(self.lock, DELAY))
							πF.SetLineno(74)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdelay_unlock, "delay_unlock"); πE != nil {
								continue
							}
							πTemp005[0] = µdelay_unlock
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßlock, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßDELAY); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, πTemp006).ToObject()
							πTemp005[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_thread); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßstart_new_thread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßverbose, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label1
							}
							goto Label2
							// line 75: if test_support.verbose:
							πF.SetLineno(75)
						Label1:
							// line 76: print
							πF.SetLineno(76)
							πTemp005 = make([]*πg.Object, 0)
							if πE = πg.Print(πF, πTemp005, true); πE != nil {
								continue
							}
							// line 77: print "*** Waiting for thread to release the lock "\
							πF.SetLineno(77)
							πTemp005 = make([]*πg.Object, 1)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßDELAY); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("*** Waiting for thread to release the lock (approx. %s sec.) ***").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							if πE = πg.Print(πF, πTemp005, true); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 79: self.lock.acquire()
							πF.SetLineno(79)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßlock, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 80: end_time = int(time.time())
							πF.SetLineno(80)
							πTemp005 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßtime, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µend_time = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßverbose, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label3
							}
							goto Label4
							// line 81: if test_support.verbose:
							πF.SetLineno(81)
						Label3:
							// line 82: print "done"
							πF.SetLineno(82)
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = ßdone.ToObject()
							if πE = πg.Print(πF, πTemp005, true); πE != nil {
								continue
							}
							goto Label4
						Label4:
							// line 83: self.assertGreaterEqual(end_time - start_time, DELAY,
							πF.SetLineno(83)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µend_time, "end_time"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstart_time, "start_time"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, µend_time, µstart_time); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßDELAY); πE != nil {
								continue
							}
							πTemp005[1] = πTemp003
							πTemp005[2] = πg.NewStr("Blocking by unconditional acquiring failed.").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertGreaterEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_uncond_acquire_blocking.ToObject(), πTemp010); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("LockTests").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLockTests.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 86: class MiscTests(unittest.TestCase):
			πF.SetLineno(86)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßTestCase, nil); πE != nil {
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
			_, πE = πg.NewCode("MiscTests", "build/src/__python__/test/test_dummy_thread.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 87: """Miscellaneous tests."""
					πF.SetLineno(87)
					// line 89: def test_exit(self):
					πF.SetLineno(89)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_exit", "build/src/__python__/test/test_dummy_thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 91: self.assertRaises(SystemExit, _thread.exit)
							πF.SetLineno(91)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSystemExit); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_thread); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßexit, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_exit.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 93: def test_ident(self):
					πF.SetLineno(93)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_ident", "build/src/__python__/test/test_dummy_thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 95: self.assertIsInstance(_thread.get_ident(), int,
							πF.SetLineno(95)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_thread); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget_ident, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							πTemp001[2] = πg.NewStr("_thread.get_ident() returned a non-integer").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIsInstance, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 97: self.assertNotEqual(_thread.get_ident(), 0,
							πF.SetLineno(97)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_thread); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget_ident, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(0).ToObject()
							πTemp001[2] = πg.NewStr("_thread.get_ident() returned 0").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertNotEqual, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_ident.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 100: def test_LockType(self):
					πF.SetLineno(100)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_LockType", "build/src/__python__/test/test_dummy_thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 102: self.assertIsInstance(_thread.allocate_lock(), _thread.LockType,
							πF.SetLineno(102)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_thread); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßallocate_lock, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_thread); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßLockType, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewStr("_thread.LockType is not an instance of what is returned by _thread.allocate_lock()").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIsInstance, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_LockType.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 106: def test_interrupt_main(self):
					πF.SetLineno(106)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_interrupt_main", "build/src/__python__/test/test_dummy_thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µcall_interrupt *πg.Object = πg.UnboundLocal; _ = µcall_interrupt
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
							// line 109: def call_interrupt():
							πF.SetLineno(109)
							πTemp002 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("call_interrupt", "build/src/__python__/test/test_dummy_thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 110: _thread.interrupt_main()
									πF.SetLineno(110)
									if πTemp001, πE = πg.ResolveGlobal(πF, ß_thread); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinterrupt_main, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
							µcall_interrupt = πTemp001
							// line 111: self.assertRaises(KeyboardInterrupt, _thread.start_new_thread,
							πF.SetLineno(111)
							πTemp003 = πF.MakeArgs(4)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßKeyboardInterrupt); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_thread); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßstart_new_thread, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp005
							if πE = πg.CheckLocal(πF, µcall_interrupt, "call_interrupt"); πE != nil {
								continue
							}
							πTemp003[2] = µcall_interrupt
							if πTemp004, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[3] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_interrupt_main.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 114: def test_interrupt_in_main(self):
					πF.SetLineno(114)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_interrupt_in_main", "build/src/__python__/test/test_dummy_thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 117: self.assertRaises(KeyboardInterrupt, _thread.interrupt_main)
							πF.SetLineno(117)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyboardInterrupt); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_thread); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinterrupt_main, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_interrupt_in_main.ToObject(), πTemp006); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("MiscTests").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMiscTests.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 119: class ThreadTests(unittest.TestCase):
			πF.SetLineno(119)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßTestCase, nil); πE != nil {
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
			_, πE = πg.NewCode("ThreadTests", "build/src/__python__/test/test_dummy_thread.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 120: """Test thread creation."""
					πF.SetLineno(120)
					// line 122: def test_arg_passing(self):
					πF.SetLineno(122)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_arg_passing", "build/src/__python__/test/test_dummy_thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µarg_tester *πg.Object = πg.UnboundLocal; _ = µarg_tester
						var µtesting_queue *πg.Object = πg.UnboundLocal; _ = µtesting_queue
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []πg.Param
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Dict
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 124: def arg_tester(queue, arg1=False, arg2=False):
							πF.SetLineno(124)
							πTemp002 = make([]πg.Param, 3)
							πTemp002[0] = πg.Param{Name: "queue", Def: nil}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp002[1] = πg.Param{Name: "arg1", Def: πTemp003}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp002[2] = πg.Param{Name: "arg2", Def: πTemp003}
							πTemp001 = πg.NewFunction(πg.NewCode("arg_tester", "build/src/__python__/test/test_dummy_thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µqueue *πg.Object = πArgs[0]; _ = µqueue
								var µarg1 *πg.Object = πArgs[1]; _ = µarg1
								var µarg2 *πg.Object = πArgs[2]; _ = µarg2
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
									// line 125: """Use to test _thread.start_new_thread() passes args properly."""
									πF.SetLineno(125)
									// line 126: queue.put((arg1, arg2))
									πF.SetLineno(126)
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µarg1, "arg1"); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µarg2, "arg2"); πE != nil {
										continue
									}
									πTemp002 = πg.NewTuple2(µarg1, µarg2).ToObject()
									πTemp001[0] = πTemp002
									if πE = πg.CheckLocal(πF, µqueue, "queue"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µqueue, ßput, nil); πE != nil {
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
							µarg_tester = πTemp001
							// line 128: testing_queue = Queue.Queue(1)
							πF.SetLineno(128)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewInt(1).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßQueue); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßQueue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µtesting_queue = πTemp003
							// line 129: _thread.start_new_thread(arg_tester, (testing_queue, True, True))
							πF.SetLineno(129)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µarg_tester, "arg_tester"); πE != nil {
								continue
							}
							πTemp004[0] = µarg_tester
							if πE = πg.CheckLocal(πF, µtesting_queue, "testing_queue"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple3(µtesting_queue, πTemp005, πTemp006).ToObject()
							πTemp004[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_thread); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßstart_new_thread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 130: result = testing_queue.get()
							πF.SetLineno(130)
							if πE = πg.CheckLocal(πF, µtesting_queue, "testing_queue"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µtesting_queue, ßget, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µresult = πTemp005
							// line 131: self.assertTrue(result[0] and result[1],
							πF.SetLineno(131)
							πTemp004 = πF.MakeArgs(2)
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µresult, πTemp005); πE != nil {
								continue
							}
							πTemp003 = πTemp006
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp007 {
								goto Label1
							}
							πTemp005 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µresult, πTemp005); πE != nil {
								continue
							}
							πTemp003 = πTemp006
						Label1:
							πTemp004[0] = πTemp003
							πTemp004[1] = πg.NewStr("Argument passing for thread creation using tuple failed").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 133: _thread.start_new_thread(arg_tester, tuple(), {'queue':testing_queue,
							πF.SetLineno(133)
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µarg_tester, "arg_tester"); πE != nil {
								continue
							}
							πTemp004[0] = µarg_tester
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp005
							πTemp008 = πg.NewDict()
							if πE = πg.CheckLocal(πF, µtesting_queue, "testing_queue"); πE != nil {
								continue
							}
							if πE = πTemp008.SetItem(πF, ßqueue.ToObject(), µtesting_queue); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πTemp008.SetItem(πF, ßarg1.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πTemp008.SetItem(πF, ßarg2.ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp003 = πTemp008.ToObject()
							πTemp004[2] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_thread); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßstart_new_thread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 135: result = testing_queue.get()
							πF.SetLineno(135)
							if πE = πg.CheckLocal(πF, µtesting_queue, "testing_queue"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µtesting_queue, ßget, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µresult = πTemp005
							// line 136: self.assertTrue(result[0] and result[1],
							πF.SetLineno(136)
							πTemp004 = πF.MakeArgs(2)
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µresult, πTemp005); πE != nil {
								continue
							}
							πTemp003 = πTemp006
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp007 {
								goto Label2
							}
							πTemp005 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µresult, πTemp005); πE != nil {
								continue
							}
							πTemp003 = πTemp006
						Label2:
							πTemp004[0] = πTemp003
							πTemp004[1] = πg.NewStr("Argument passing for thread creation using kwargs failed").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 138: _thread.start_new_thread(arg_tester, (testing_queue, True), {'arg2':True})
							πF.SetLineno(138)
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µarg_tester, "arg_tester"); πE != nil {
								continue
							}
							πTemp004[0] = µarg_tester
							if πE = πg.CheckLocal(πF, µtesting_queue, "testing_queue"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(µtesting_queue, πTemp005).ToObject()
							πTemp004[1] = πTemp003
							πTemp008 = πg.NewDict()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πTemp008.SetItem(πF, ßarg2.ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp003 = πTemp008.ToObject()
							πTemp004[2] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_thread); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßstart_new_thread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 139: result = testing_queue.get()
							πF.SetLineno(139)
							if πE = πg.CheckLocal(πF, µtesting_queue, "testing_queue"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µtesting_queue, ßget, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µresult = πTemp005
							// line 140: self.assertTrue(result[0] and result[1],
							πF.SetLineno(140)
							πTemp004 = πF.MakeArgs(2)
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µresult, πTemp005); πE != nil {
								continue
							}
							πTemp003 = πTemp006
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp007 {
								goto Label3
							}
							πTemp005 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µresult, πTemp005); πE != nil {
								continue
							}
							πTemp003 = πTemp006
						Label3:
							πTemp004[0] = πTemp003
							πTemp004[1] = πg.NewStr("Argument passing for thread creation using both tuple and kwargs failed").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_arg_passing.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 144: def test_multi_creation(self):
					πF.SetLineno(144)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_multi_creation", "build/src/__python__/test/test_dummy_thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µqueue_mark *πg.Object = πg.UnboundLocal; _ = µqueue_mark
						var µthread_count *πg.Object = πg.UnboundLocal; _ = µthread_count
						var µtesting_queue *πg.Object = πg.UnboundLocal; _ = µtesting_queue
						var µcount *πg.Object = πg.UnboundLocal; _ = µcount
						var µlocal_delay *πg.Object = πg.UnboundLocal; _ = µlocal_delay
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 3: goto Label3
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							// line 146: def queue_mark(queue, delay):
							πF.SetLineno(146)
							πTemp002 = make([]πg.Param, 2)
							πTemp002[0] = πg.Param{Name: "queue", Def: nil}
							πTemp002[1] = πg.Param{Name: "delay", Def: nil}
							πTemp001 = πg.NewFunction(πg.NewCode("queue_mark", "build/src/__python__/test/test_dummy_thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µqueue *πg.Object = πArgs[0]; _ = µqueue
								var µdelay *πg.Object = πArgs[1]; _ = µdelay
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
									// line 147: """Wait for ``delay`` seconds and then put something into ``queue``"""
									πF.SetLineno(147)
									// line 148: time.sleep(delay)
									πF.SetLineno(148)
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µdelay, "delay"); πE != nil {
										continue
									}
									πTemp001[0] = µdelay
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
									// line 149: queue.put(_thread.get_ident())
									πF.SetLineno(149)
									πTemp001 = πF.MakeArgs(1)
									if πTemp002, πE = πg.ResolveGlobal(πF, ß_thread); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget_ident, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
										continue
									}
									πTemp001[0] = πTemp002
									if πE = πg.CheckLocal(πF, µqueue, "queue"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µqueue, ßput, nil); πE != nil {
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
							µqueue_mark = πTemp001
							// line 151: thread_count = 5
							πF.SetLineno(151)
							µthread_count = πg.NewInt(5).ToObject()
							// line 152: testing_queue = Queue.Queue(thread_count)
							πF.SetLineno(152)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µthread_count, "thread_count"); πE != nil {
								continue
							}
							πTemp003[0] = µthread_count
							if πTemp004, πE = πg.ResolveGlobal(πF, ßQueue); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßQueue, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µtesting_queue = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßverbose, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 153: if test_support.verbose:
							πF.SetLineno(153)
						Label1:
							// line 154: print
							πF.SetLineno(154)
							πTemp003 = make([]*πg.Object, 0)
							if πE = πg.Print(πF, πTemp003, true); πE != nil {
								continue
							}
							// line 155: print "*** Testing multiple thread creation "\
							πF.SetLineno(155)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp007, πE = πg.ResolveGlobal(πF, ßDELAY); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µthread_count, "thread_count"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(πTemp007, µthread_count).ToObject()
							if πTemp004, πE = πg.Mod(πF, πg.NewStr("*** Testing multiple thread creation (will take approx. %s to %s sec.) ***").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πE = πg.Print(πF, πTemp003, true); πE != nil {
								continue
							}
							goto Label2
						Label2:
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µthread_count, "thread_count"); πE != nil {
								continue
							}
							πTemp003[0] = µthread_count
							if πTemp005, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.Iter(πF, πTemp007); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp006 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label5
							}
							if πTemp005, πE = πg.Next(πF, πTemp004); πE != nil {
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
								µcount = πTemp005
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(3)            
							if πTemp005, πE = πg.ResolveGlobal(πF, ßDELAY); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label6
							}
							goto Label7
							// line 158: if DELAY:
							πF.SetLineno(158)
						Label6:
							// line 159: local_delay = round(random.random(), 1)
							πF.SetLineno(159)
							πTemp003 = πF.MakeArgs(2)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßrandom); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßrandom, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							πTemp003[1] = πg.NewInt(1).ToObject()
							if πTemp005, πE = πg.ResolveGlobal(πF, ßround); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µlocal_delay = πTemp007
							goto Label8
						Label7:
							// line 161: local_delay = 0
							πF.SetLineno(161)
							µlocal_delay = πg.NewInt(0).ToObject()
							goto Label8
						Label8:
							// line 162: _thread.start_new_thread(queue_mark,
							πF.SetLineno(162)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µqueue_mark, "queue_mark"); πE != nil {
								continue
							}
							πTemp003[0] = µqueue_mark
							if πE = πg.CheckLocal(πF, µtesting_queue, "testing_queue"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlocal_delay, "local_delay"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(µtesting_queue, µlocal_delay).ToObject()
							πTemp003[1] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ß_thread); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßstart_new_thread, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp007.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 164: time.sleep(DELAY)
							πF.SetLineno(164)
							πTemp003 = πF.MakeArgs(1)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßDELAY); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßsleep, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßverbose, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label9
							}
							goto Label10
							// line 165: if test_support.verbose:
							πF.SetLineno(165)
						Label9:
							// line 166: print 'done'
							πF.SetLineno(166)
							πTemp003 = make([]*πg.Object, 1)
							πTemp003[0] = ßdone.ToObject()
							if πE = πg.Print(πF, πTemp003, true); πE != nil {
								continue
							}
							goto Label10
						Label10:
							// line 167: self.assertEqual(testing_queue.qsize(), thread_count,
							πF.SetLineno(167)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µtesting_queue, "testing_queue"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µtesting_queue, ßqsize, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							if πE = πg.CheckLocal(πF, µthread_count, "thread_count"); πE != nil {
								continue
							}
							πTemp003[1] = µthread_count
							if πE = πg.CheckLocal(πF, µthread_count, "thread_count"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßDELAY); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(µthread_count, πTemp007).ToObject()
							if πTemp004, πE = πg.Mod(πF, πg.NewStr("Not all %s threads executed properly after %s sec.").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp003[2] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_multi_creation.ToObject(), πTemp003); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("ThreadTests").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßThreadTests.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 171: def test_main(imported_module=None):
			πF.SetLineno(171)
			πTemp007 = make([]πg.Param, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp007[0] = πg.Param{Name: "imported_module", Def: πTemp004}
			πTemp001 = πg.NewFunction(πg.NewCode("test_main", "build/src/__python__/test/test_dummy_thread.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µimported_module *πg.Object = πArgs[0]; _ = µimported_module
				var πTemp001 bool
				_ = πTemp001
				var πTemp002 *πg.Object
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
					// line 172: global _thread, DELAY
					πF.SetLineno(172)
					if πE = πg.CheckLocal(πF, µimported_module, "imported_module"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µimported_module); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label1
					}
					goto Label2
					// line 173: if imported_module:
					πF.SetLineno(173)
				Label1:
					// line 174: _thread = imported_module
					πF.SetLineno(174)
					if πE = πg.CheckLocal(πF, µimported_module, "imported_module"); πE != nil {
						continue
					}
					if πE = πF.Globals().SetItem(πF, ß_thread.ToObject(), µimported_module); πE != nil {
						continue
					}
					// line 175: DELAY = 2
					πF.SetLineno(175)
					if πE = πF.Globals().SetItem(πF, ßDELAY.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					goto Label2
				Label2:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßverbose, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label3
					}
					goto Label4
					// line 176: if test_support.verbose:
					πF.SetLineno(176)
				Label3:
					// line 177: print
					πF.SetLineno(177)
					πTemp004 = make([]*πg.Object, 0)
					if πE = πg.Print(πF, πTemp004, true); πE != nil {
						continue
					}
					// line 178: print "*** Using %s as _thread module ***" % _thread
					πF.SetLineno(178)
					πTemp004 = make([]*πg.Object, 1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_thread); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("*** Using %s as _thread module ***").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp004[0] = πTemp002
					if πE = πg.Print(πF, πTemp004, true); πE != nil {
						continue
					}
					goto Label4
				Label4:
					// line 179: test_support.run_unittest(LockTests, MiscTests, ThreadTests)
					πF.SetLineno(179)
					πTemp004 = πF.MakeArgs(3)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßLockTests); πE != nil {
						continue
					}
					πTemp004[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßMiscTests); πE != nil {
						continue
					}
					πTemp004[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßThreadTests); πE != nil {
						continue
					}
					πTemp004[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrun_unittest, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
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
			// line 181: if __name__ == '__main__':
			πF.SetLineno(181)
		Label1:
			// line 182: test_main()
			πF.SetLineno(182)
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
	πg.RegisterModule("test.test_dummy_thread", Code)
}
