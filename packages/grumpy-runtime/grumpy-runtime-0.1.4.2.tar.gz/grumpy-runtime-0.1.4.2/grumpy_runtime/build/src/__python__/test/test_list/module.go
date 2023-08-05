package test_list
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/test/test_list.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßCommonTest := πg.InternStr("CommonTest")
		ßListTest := πg.InternStr("ListTest")
		ßMemoryError := πg.InternStr("MemoryError")
		ßNone := πg.InternStr("None")
		ßOverflowError := πg.InternStr("OverflowError")
		ßTrue := πg.InternStr("True")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßa := πg.InternStr("a")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertRaises := πg.InternStr("assertRaises")
		ßassertTrue := πg.InternStr("assertTrue")
		ßexpectedFailure := πg.InternStr("expectedFailure")
		ßextend := πg.InternStr("extend")
		ßint := πg.InternStr("int")
		ßlen := πg.InternStr("len")
		ßlist := πg.InternStr("list")
		ßlist_tests := πg.InternStr("list_tests")
		ßm := πg.InternStr("m")
		ßmaxint := πg.InternStr("maxint")
		ßmaxsize := πg.InternStr("maxsize")
		ßp := πg.InternStr("p")
		ßrun_unittest := πg.InternStr("run_unittest")
		ßs := πg.InternStr("s")
		ßspam := πg.InternStr("spam")
		ßsuper := πg.InternStr("super")
		ßsys := πg.InternStr("sys")
		ßtest_basic := πg.InternStr("test_basic")
		ßtest_identity := πg.InternStr("test_identity")
		ßtest_len := πg.InternStr("test_len")
		ßtest_main := πg.InternStr("test_main")
		ßtest_overflow := πg.InternStr("test_overflow")
		ßtest_support := πg.InternStr("test_support")
		ßtest_truth := πg.InternStr("test_truth")
		ßtype2test := πg.InternStr("type2test")
		ßunittest := πg.InternStr("unittest")
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
		var πTemp009 πg.KWArgs
		_ = πTemp009
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: import sys
			πF.SetLineno(1)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 2: import unittest
			πF.SetLineno(2)
			if πTemp002, πE = πg.ImportModule(πF, "unittest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 3: from test import test_support, list_tests
			πF.SetLineno(3)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßtest_support.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "test.list_tests"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßlist_tests.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 5: class ListTest(list_tests.CommonTest):
			πF.SetLineno(5)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlist_tests); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßCommonTest, nil); πE != nil {
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
			_, πE = πg.NewCode("ListTest", "build/src/__python__/test/test_list.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp007 []*πg.Object
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
					// line 6: type2test = list
					πF.SetLineno(6)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßlist); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßtype2test.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 8: def test_basic(self):
					πF.SetLineno(8)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_basic", "build/src/__python__/test/test_list.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µl0_3 *πg.Object = πg.UnboundLocal; _ = µl0_3
						var µl0_3_bis *πg.Object = πg.UnboundLocal; _ = µl0_3_bis
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
						var πTemp008 []πg.Param
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 9: self.assertEqual(list([]), [])
							πF.SetLineno(9)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp003...).ToObject()
							πTemp002[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							πTemp002 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 10: l0_3 = [0, 1, 2, 3]
							πF.SetLineno(10)
							πTemp001 = make([]*πg.Object, 4)
							πTemp001[0] = πg.NewInt(0).ToObject()
							πTemp001[1] = πg.NewInt(1).ToObject()
							πTemp001[2] = πg.NewInt(2).ToObject()
							πTemp001[3] = πg.NewInt(3).ToObject()
							πTemp004 = πg.NewList(πTemp001...).ToObject()
							µl0_3 = πTemp004
							// line 11: l0_3_bis = list(l0_3)
							πF.SetLineno(11)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µl0_3, "l0_3"); πE != nil {
								continue
							}
							πTemp001[0] = µl0_3
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µl0_3_bis = πTemp005
							// line 12: self.assertEqual(l0_3, l0_3_bis)
							πF.SetLineno(12)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µl0_3, "l0_3"); πE != nil {
								continue
							}
							πTemp001[0] = µl0_3
							if πE = πg.CheckLocal(πF, µl0_3_bis, "l0_3_bis"); πE != nil {
								continue
							}
							πTemp001[1] = µl0_3_bis
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 13: self.assertTrue(l0_3 is not l0_3_bis)
							πF.SetLineno(13)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µl0_3, "l0_3"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µl0_3_bis, "l0_3_bis"); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(µl0_3 != µl0_3_bis).ToObject()
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 14: self.assertEqual(list(()), [])
							πF.SetLineno(14)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp004 = πg.NewTuple0().ToObject()
							πTemp002[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							πTemp002 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 15: self.assertEqual(list((0, 1, 2, 3)), [0, 1, 2, 3])
							πF.SetLineno(15)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp004 = πg.NewTuple4(πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject(), πg.NewInt(3).ToObject()).ToObject()
							πTemp002[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							πTemp002 = make([]*πg.Object, 4)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp002[2] = πg.NewInt(2).ToObject()
							πTemp002[3] = πg.NewInt(3).ToObject()
							πTemp004 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 16: self.assertEqual(list(''), [])
							πF.SetLineno(16)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ß.ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							πTemp002 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 17: self.assertEqual(list('spam'), ['s', 'p', 'a', 'm'])
							πF.SetLineno(17)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßspam.ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							πTemp002 = make([]*πg.Object, 4)
							πTemp002[0] = ßs.ToObject()
							πTemp002[1] = ßp.ToObject()
							πTemp002[2] = ßa.ToObject()
							πTemp002[3] = ßm.ToObject()
							πTemp004 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßmaxsize, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Eq(πF, πTemp006, πg.NewInt(2147483647).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label1
							}
							goto Label2
							// line 19: if sys.maxsize == 0x7fffffff:
							πF.SetLineno(19)
						Label1:
							// line 34: self.assertRaises(MemoryError, list, xrange(sys.maxint // 2))
							πF.SetLineno(34)
							πTemp001 = πF.MakeArgs(3)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßMemoryError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							πTemp002 = πF.MakeArgs(1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßmaxint, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.FloorDiv(πF, πTemp006, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[2] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label2
						Label2:
							// line 37: x = []
							πF.SetLineno(37)
							πTemp001 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp001...).ToObject()
							µx = πTemp004
							// line 38: x.extend(-y for y in x)
							πF.SetLineno(38)
							πTemp001 = πF.MakeArgs(1)
							πTemp008 = make([]πg.Param, 0)
							πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/test/test_list.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µy *πg.Object = πg.UnboundLocal; _ = µy
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
								return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
									for ; πF.State() >= 0; πF.PopCheckpoint() {
										switch πF.State() {
										case 0:
										case 1: goto Label1
										case 2: goto Label2
										case 4: goto Label4
										default: panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µx); πE != nil {
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
											µy = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 38: x.extend(-y for y in x)
										πF.SetLineno(38)
										if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
											continue
										}
										if πTemp004, πE = πg.Neg(πF, µy); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return πTemp004, nil
									Label4:
										πTemp005 = πSent
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
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µx, ßextend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 39: self.assertEqual(x, [])
							πF.SetLineno(39)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							πTemp002 = make([]*πg.Object, 0)
							πTemp005 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_basic.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 41: def test_truth(self):
					πF.SetLineno(41)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_truth", "build/src/__python__/test/test_list.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							// line 42: super(ListTest, self).test_truth()
							πF.SetLineno(42)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßListTest); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[1] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßtest_truth, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 43: self.assertTrue(not [])
							πF.SetLineno(43)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp004...).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							πTemp001[0] = πTemp002
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
							// line 44: self.assertTrue([42])
							πF.SetLineno(44)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = make([]*πg.Object, 1)
							πTemp004[0] = πg.NewInt(42).ToObject()
							πTemp002 = πg.NewList(πTemp004...).ToObject()
							πTemp001[0] = πTemp002
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
					if πE = πClass.SetItem(πF, ßtest_truth.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 46: def test_identity(self):
					πF.SetLineno(46)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_identity", "build/src/__python__/test/test_list.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 47: self.assertTrue([] is not [])
							πF.SetLineno(47)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp003...).ToObject()
							πTemp003 = make([]*πg.Object, 0)
							πTemp005 = πg.NewList(πTemp003...).ToObject()
							πTemp002 = πg.GetBool(πTemp004 != πTemp005).ToObject()
							πTemp001[0] = πTemp002
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
					if πE = πClass.SetItem(πF, ßtest_identity.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 49: def test_len(self):
					πF.SetLineno(49)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_len", "build/src/__python__/test/test_list.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							// line 50: super(ListTest, self).test_len()
							πF.SetLineno(50)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßListTest); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[1] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßtest_len, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 51: self.assertEqual(len([]), 0)
							πF.SetLineno(51)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewInt(0).ToObject()
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
							// line 52: self.assertEqual(len([0]), 1)
							πF.SetLineno(52)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewInt(0).ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewInt(1).ToObject()
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
							// line 53: self.assertEqual(len([0, 1, 2]), 3)
							πF.SetLineno(53)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(0).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewInt(3).ToObject()
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_len.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 56: def test_overflow(self):
					πF.SetLineno(56)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_overflow", "build/src/__python__/test/test_list.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlst *πg.Object = πg.UnboundLocal; _ = µlst
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var µmul *πg.Object = πg.UnboundLocal; _ = µmul
						var µimul *πg.Object = πg.UnboundLocal; _ = µimul
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 []πg.Param
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 57: lst = [4, 5, 6, 7]
							πF.SetLineno(57)
							πTemp001 = make([]*πg.Object, 4)
							πTemp001[0] = πg.NewInt(4).ToObject()
							πTemp001[1] = πg.NewInt(5).ToObject()
							πTemp001[2] = πg.NewInt(6).ToObject()
							πTemp001[3] = πg.NewInt(7).ToObject()
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µlst = πTemp002
							// line 58: n = int((sys.maxsize*2+2) // len(lst))
							πF.SetLineno(58)
							πTemp001 = πF.MakeArgs(1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßmaxsize, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mul(πF, πTemp006, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
								continue
							}
							πTemp007[0] = µlst
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp002, πE = πg.FloorDiv(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µn = πTemp003
							// line 59: def mul(a, b): return a * b
							πF.SetLineno(59)
							πTemp008 = make([]πg.Param, 2)
							πTemp008[0] = πg.Param{Name: "a", Def: nil}
							πTemp008[1] = πg.Param{Name: "b", Def: nil}
							πTemp002 = πg.NewFunction(πg.NewCode("mul", "build/src/__python__/test/test_list.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µa *πg.Object = πArgs[0]; _ = µa
								var µb *πg.Object = πArgs[1]; _ = µb
								var πTemp001 *πg.Object
								_ = πTemp001
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 59: def mul(a, b): return a * b
									πF.SetLineno(59)
									if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.Mul(πF, µa, µb); πE != nil {
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
							µmul = πTemp002
							// line 60: def imul(a, b): a *= b
							πF.SetLineno(60)
							πTemp008 = make([]πg.Param, 2)
							πTemp008[0] = πg.Param{Name: "a", Def: nil}
							πTemp008[1] = πg.Param{Name: "b", Def: nil}
							πTemp003 = πg.NewFunction(πg.NewCode("imul", "build/src/__python__/test/test_list.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µa *πg.Object = πArgs[0]; _ = µa
								var µb *πg.Object = πArgs[1]; _ = µb
								var πTemp001 *πg.Object
								_ = πTemp001
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 60: def imul(a, b): a *= b
									πF.SetLineno(60)
									if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.IMul(πF, µa, µb); πE != nil {
										continue
									}
									µa = πTemp001
								}
								if πE != nil {
									πR = nil
								} else if πR == nil {
									πR = πg.None
								}
								return πR, πE
							}), πF.Globals()).ToObject()
							µimul = πTemp003
							// line 61: self.assertRaises((MemoryError, OverflowError), mul, lst, n)
							πF.SetLineno(61)
							πTemp001 = πF.MakeArgs(4)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßMemoryError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßOverflowError); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µmul, "mul"); πE != nil {
								continue
							}
							πTemp001[1] = µmul
							if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
								continue
							}
							πTemp001[2] = µlst
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[3] = µn
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 62: self.assertRaises((MemoryError, OverflowError), imul, lst, n)
							πF.SetLineno(62)
							πTemp001 = πF.MakeArgs(4)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßMemoryError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßOverflowError); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µimul, "imul"); πE != nil {
								continue
							}
							πTemp001[1] = µimul
							if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
								continue
							}
							πTemp001[2] = µlst
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[3] = µn
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_overflow.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 55: @unittest.expectedFailure
					πF.SetLineno(55)
					πTemp007 = πF.MakeArgs(1)
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßtest_overflow); πE != nil {
						continue
					}
					πTemp007[0] = πTemp008
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp009.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πClass.SetItem(πF, ßtest_overflow.ToObject(), πTemp008); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("ListTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßListTest.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 64: def test_main(verbose=None):
			πF.SetLineno(64)
			πTemp007 = make([]πg.Param, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp007[0] = πg.Param{Name: "verbose", Def: πTemp004}
			πTemp001 = πg.NewFunction(πg.NewCode("test_main", "build/src/__python__/test/test_list.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µverbose *πg.Object = πArgs[0]; _ = µverbose
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
					// line 65: test_support.run_unittest(ListTest)
					πF.SetLineno(65)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßListTest); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
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
			// line 79: if __name__ == "__main__":
			πF.SetLineno(79)
		Label1:
			// line 80: test_main(verbose=True)
			πF.SetLineno(80)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp009 = πg.KWArgs{
				{"verbose", πTemp004},
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßtest_main); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp004.Call(πF, nil, πTemp009); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("test.test_list", Code)
}
