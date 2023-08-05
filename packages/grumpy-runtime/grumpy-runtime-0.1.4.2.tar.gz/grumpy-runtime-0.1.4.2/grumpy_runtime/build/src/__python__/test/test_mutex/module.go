package test_mutex
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/test/test_mutex.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßFalse := πg.InternStr("False")
		ßMutexTest := πg.InternStr("MutexTest")
		ßTestCase := πg.InternStr("TestCase")
		ßTrue := πg.InternStr("True")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertFalse := πg.InternStr("assertFalse")
		ßassertTrue := πg.InternStr("assertTrue")
		ßeggs := πg.InternStr("eggs")
		ßlock := πg.InternStr("lock")
		ßmutex := πg.InternStr("mutex")
		ßrun_unittest := πg.InternStr("run_unittest")
		ßspam := πg.InternStr("spam")
		ßtest := πg.InternStr("test")
		ßtest_lock_and_unlock := πg.InternStr("test_lock_and_unlock")
		ßtest_main := πg.InternStr("test_main")
		ßtest_support := πg.InternStr("test_support")
		ßunittest := πg.InternStr("unittest")
		ßunlock := πg.InternStr("unlock")
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
			// line 1: import unittest
			πF.SetLineno(1)
			if πTemp002, πE = πg.ImportModule(πF, "unittest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 2: import test.test_support
			πF.SetLineno(2)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 5: import mutex
			πF.SetLineno(5)
			if πTemp002, πE = πg.ImportModule(πF, "mutex"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßmutex.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 7: class MutexTest(unittest.TestCase):
			πF.SetLineno(7)
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
			_, πE = πg.NewCode("MutexTest", "build/src/__python__/test/test_mutex.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 9: def test_lock_and_unlock(self):
					πF.SetLineno(9)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_lock_and_unlock", "build/src/__python__/test/test_mutex.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µcalled_by_mutex *πg.Object = πg.UnboundLocal; _ = µcalled_by_mutex
						var µcalled_by_mutex2 *πg.Object = πg.UnboundLocal; _ = µcalled_by_mutex2
						var µm *πg.Object = πg.UnboundLocal; _ = µm
						var µread_for_2 *πg.Object = πg.UnboundLocal; _ = µread_for_2
						var µready_for_2 *πg.Object = πg.UnboundLocal; _ = µready_for_2
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 11: def called_by_mutex(some_data):
							πF.SetLineno(11)
							πTemp002 = make([]πg.Param, 1)
							πTemp002[0] = πg.Param{Name: "some_data", Def: nil}
							πTemp001 = πg.NewFunction(πg.NewCode("called_by_mutex", "build/src/__python__/test/test_mutex.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µsome_data *πg.Object = πArgs[0]; _ = µsome_data
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
									// line 12: self.assertEqual(some_data, "spam")
									πF.SetLineno(12)
									πTemp001 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µsome_data, "some_data"); πE != nil {
										continue
									}
									πTemp001[0] = µsome_data
									πTemp001[1] = ßspam.ToObject()
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									// line 13: self.assertTrue(m.test(), "mutex not held")
									πF.SetLineno(13)
									πTemp001 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µm, ßtest, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
										continue
									}
									πTemp001[0] = πTemp003
									πTemp001[1] = πg.NewStr("mutex not held").ToObject()
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									// line 15: m.lock(called_by_mutex2, "eggs")
									πF.SetLineno(15)
									πTemp001 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µcalled_by_mutex2, "called_by_mutex2"); πE != nil {
										continue
									}
									πTemp001[0] = µcalled_by_mutex2
									πTemp001[1] = ßeggs.ToObject()
									if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µm, ßlock, nil); πE != nil {
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
							µcalled_by_mutex = πTemp001
							// line 17: def called_by_mutex2(some_data):
							πF.SetLineno(17)
							πTemp002 = make([]πg.Param, 1)
							πTemp002[0] = πg.Param{Name: "some_data", Def: nil}
							πTemp003 = πg.NewFunction(πg.NewCode("called_by_mutex2", "build/src/__python__/test/test_mutex.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µsome_data *πg.Object = πArgs[0]; _ = µsome_data
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
									// line 18: self.assertEqual(some_data, "eggs")
									πF.SetLineno(18)
									πTemp001 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µsome_data, "some_data"); πE != nil {
										continue
									}
									πTemp001[0] = µsome_data
									πTemp001[1] = ßeggs.ToObject()
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									// line 19: self.assertTrue(m.test(), "mutex not held")
									πF.SetLineno(19)
									πTemp001 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µm, ßtest, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
										continue
									}
									πTemp001[0] = πTemp003
									πTemp001[1] = πg.NewStr("mutex not held").ToObject()
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									// line 20: self.assertTrue(ready_for_2,
									πF.SetLineno(20)
									πTemp001 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µready_for_2, "ready_for_2"); πE != nil {
										continue
									}
									πTemp001[0] = µready_for_2
									πTemp001[1] = πg.NewStr("called_by_mutex2 called too soon").ToObject()
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
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
							µcalled_by_mutex2 = πTemp003
							// line 23: m = mutex.mutex()
							πF.SetLineno(23)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßmutex); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßmutex, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							µm = πTemp004
							// line 24: read_for_2 = False
							πF.SetLineno(24)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							µread_for_2 = πTemp004
							// line 25: m.lock(called_by_mutex, "spam")
							πF.SetLineno(25)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µcalled_by_mutex, "called_by_mutex"); πE != nil {
								continue
							}
							πTemp006[0] = µcalled_by_mutex
							πTemp006[1] = ßspam.ToObject()
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µm, ßlock, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 26: ready_for_2 = True
							πF.SetLineno(26)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µready_for_2 = πTemp004
							// line 28: m.unlock()
							πF.SetLineno(28)
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µm, ßunlock, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 29: m.unlock()
							πF.SetLineno(29)
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µm, ßunlock, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 30: self.assertFalse(m.test(), "mutex still held")
							πF.SetLineno(30)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µm, ßtest, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp005
							πTemp006[1] = πg.NewStr("mutex still held").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_lock_and_unlock.ToObject(), πTemp001); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("MutexTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMutexTest.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 32: def test_main():
			πF.SetLineno(32)
			πTemp007 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("test_main", "build/src/__python__/test/test_mutex.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 33: test.test_support.run_unittest(MutexTest)
					πF.SetLineno(33)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßMutexTest); πE != nil {
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
			// line 35: if __name__ == "__main__":
			πF.SetLineno(35)
		Label1:
			// line 36: test_main()
			πF.SetLineno(36)
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
	πg.RegisterModule("test.test_mutex", Code)
}
