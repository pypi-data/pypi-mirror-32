package test_slice
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/test/test_slice.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßException := πg.InternStr("Exception")
		ßNone := πg.InternStr("None")
		ßOverflowError := πg.InternStr("OverflowError")
		ßSliceTest := πg.InternStr("SliceTest")
		ßTestCase := πg.InternStr("TestCase")
		ßTypeError := πg.InternStr("TypeError")
		ß__enter__ := πg.InternStr("__enter__")
		ß__eq__ := πg.InternStr("__eq__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__hash__ := πg.InternStr("__hash__")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__setslice__ := πg.InternStr("__setslice__")
		ßappend := πg.InternStr("append")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertIsNone := πg.InternStr("assertIsNone")
		ßassertNotEqual := πg.InternStr("assertNotEqual")
		ßassertRaises := πg.InternStr("assertRaises")
		ßassertTrue := πg.InternStr("assertTrue")
		ßcheck_py3k_warnings := πg.InternStr("check_py3k_warnings")
		ßcmp := πg.InternStr("cmp")
		ßexpectedFailure := πg.InternStr("expectedFailure")
		ßgc_collect := πg.InternStr("gc_collect")
		ßhash := πg.InternStr("hash")
		ßindices := πg.InternStr("indices")
		ßmaxint := πg.InternStr("maxint")
		ßobject := πg.InternStr("object")
		ßrange := πg.InternStr("range")
		ßref := πg.InternStr("ref")
		ßrepr := πg.InternStr("repr")
		ßrun_unittest := πg.InternStr("run_unittest")
		ßs := πg.InternStr("s")
		ßslice := πg.InternStr("slice")
		ßstart := πg.InternStr("start")
		ßstep := πg.InternStr("step")
		ßstop := πg.InternStr("stop")
		ßsys := πg.InternStr("sys")
		ßtest_cmp := πg.InternStr("test_cmp")
		ßtest_constructor := πg.InternStr("test_constructor")
		ßtest_cycle := πg.InternStr("test_cycle")
		ßtest_hash := πg.InternStr("test_hash")
		ßtest_indices := πg.InternStr("test_indices")
		ßtest_main := πg.InternStr("test_main")
		ßtest_members := πg.InternStr("test_members")
		ßtest_repr := πg.InternStr("test_repr")
		ßtest_setslice_without_getslice := πg.InternStr("test_setslice_without_getslice")
		ßtest_support := πg.InternStr("test_support")
		ßunittest := πg.InternStr("unittest")
		ßweakref := πg.InternStr("weakref")
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
			// line 3: import unittest
			πF.SetLineno(3)
			if πTemp002, πE = πg.ImportModule(πF, "unittest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 4: import weakref
			πF.SetLineno(4)
			if πTemp002, πE = πg.ImportModule(πF, "weakref"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßweakref.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 7: from test import test_support
			πF.SetLineno(7)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßtest_support.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 9: import sys
			πF.SetLineno(9)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 11: class SliceTest(unittest.TestCase):
			πF.SetLineno(11)
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
			_, πE = πg.NewCode("SliceTest", "build/src/__python__/test/test_slice.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp006 []*πg.Object
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
					// line 13: def test_constructor(self):
					πF.SetLineno(13)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_constructor", "build/src/__python__/test/test_slice.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 14: self.assertRaises(TypeError, slice)
							πF.SetLineno(14)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							// line 15: self.assertRaises(TypeError, slice, 1, 2, 3, 4)
							πF.SetLineno(15)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							πTemp001[2] = πg.NewInt(1).ToObject()
							πTemp001[3] = πg.NewInt(2).ToObject()
							πTemp001[4] = πg.NewInt(3).ToObject()
							πTemp001[5] = πg.NewInt(4).ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_constructor.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 17: def test_repr(self):
					πF.SetLineno(17)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_repr", "build/src/__python__/test/test_slice.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 18: self.assertEqual(repr(slice(1, 2, 3)), "slice(1, 2, 3)")
							πF.SetLineno(18)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(3)
							πTemp003[0] = πg.NewInt(1).ToObject()
							πTemp003[1] = πg.NewInt(2).ToObject()
							πTemp003[2] = πg.NewInt(3).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002[0] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							πTemp001[1] = πg.NewStr("slice(1, 2, 3)").ToObject()
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_repr.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 20: def test_hash(self):
					πF.SetLineno(20)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_hash", "build/src/__python__/test/test_slice.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 *πg.BaseException
						_ = πTemp009
						var πTemp010 *πg.Traceback
						_ = πTemp010
						var πTemp011 *πg.Type
						_ = πTemp011
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 22: self.assertRaises(TypeError, hash, slice(5))
							πF.SetLineno(22)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhash); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(5).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp001[2] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 23: with self.assertRaises(TypeError):
							πF.SetLineno(23)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.GetAttr(πF, πTemp004.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp005.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(1)
							// line 24: slice(5).__hash__()
							πF.SetLineno(24)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(5).ToObject()
							if πTemp006, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp006, πE = πg.GetAttr(πF, πTemp007, ß__hash__, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							πF.PopCheckpoint()
						Label1:
							πTemp009, πTemp010 = nil, nil
							if πE != nil {
								πTemp009, πTemp010 = πF.ExcInfo()
							}
							if πTemp009 != nil {
								πTemp011 = πTemp009.Type()
								if πTemp006, πE = πTemp002.Call(πF, πg.Args{πTemp004, πTemp011.ToObject(), πTemp009.ToObject(), πTemp010.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp006, πE = πTemp002.Call(πF, πg.Args{πTemp004, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp009 != nil && πTemp008 != true {
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
					if πE = πClass.SetItem(πF, ßtest_hash.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 27: def test_cmp(self):
					πF.SetLineno(27)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_cmp", "build/src/__python__/test/test_slice.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µs1 *πg.Object = πg.UnboundLocal; _ = µs1
						var µs2 *πg.Object = πg.UnboundLocal; _ = µs2
						var µs3 *πg.Object = πg.UnboundLocal; _ = µs3
						var µExc *πg.Object = πg.UnboundLocal; _ = µExc
						var µBadCmp *πg.Object = πg.UnboundLocal; _ = µBadCmp
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Dict
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
							// line 28: s1 = slice(1, 2, 3)
							πF.SetLineno(28)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = πg.NewInt(1).ToObject()
							πTemp001[1] = πg.NewInt(2).ToObject()
							πTemp001[2] = πg.NewInt(3).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs1 = πTemp003
							// line 29: s2 = slice(1, 2, 3)
							πF.SetLineno(29)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = πg.NewInt(1).ToObject()
							πTemp001[1] = πg.NewInt(2).ToObject()
							πTemp001[2] = πg.NewInt(3).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs2 = πTemp003
							// line 30: s3 = slice(1, 2, 4)
							πF.SetLineno(30)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = πg.NewInt(1).ToObject()
							πTemp001[1] = πg.NewInt(2).ToObject()
							πTemp001[2] = πg.NewInt(4).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs3 = πTemp003
							// line 31: self.assertEqual(s1, s2)
							πF.SetLineno(31)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							πTemp001[0] = µs1
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							πTemp001[1] = µs2
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
							// line 32: self.assertNotEqual(s1, s3)
							πF.SetLineno(32)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							πTemp001[0] = µs1
							if πE = πg.CheckLocal(πF, µs3, "s3"); πE != nil {
								continue
							}
							πTemp001[1] = µs3
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
							// line 34: class Exc(Exception):
							πF.SetLineno(34)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							πTemp004 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("Exc", "build/src/__python__/test/test_slice.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp004
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 35: pass
									πF.SetLineno(35)
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
							if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Exc").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
								continue
							}
							µExc = πTemp005
							// line 37: class BadCmp(object):
							πF.SetLineno(37)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							πTemp004 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("BadCmp", "build/src/__python__/test/test_slice.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp004
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
									// line 38: def __eq__(self, other):
									πF.SetLineno(38)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__eq__", "build/src/__python__/test/test_slice.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µother *πg.Object = πArgs[1]; _ = µother
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
												continue
											}
											// line 39: raise Exc
											πF.SetLineno(39)
											πE = πF.Raise(µExc, nil, nil)
											continue
										}
										if πE != nil {
											πR = nil
										} else if πR == nil {
											πR = πg.None
										}
										return πR, πE
									}), πF.Globals()).ToObject()
									if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp001); πE != nil {
										continue
									}
									// line 40: __hash__ = None # Silence Py3k warning
									πF.SetLineno(40)
									if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
										continue
									}
									if πE = πClass.SetItem(πF, ß__hash__.ToObject(), πTemp003); πE != nil {
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
							if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("BadCmp").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
								continue
							}
							µBadCmp = πTemp005
							// line 42: s1 = slice(BadCmp())
							πF.SetLineno(42)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µBadCmp, "BadCmp"); πE != nil {
								continue
							}
							if πTemp002, πE = µBadCmp.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs1 = πTemp003
							// line 43: s2 = slice(BadCmp())
							πF.SetLineno(43)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µBadCmp, "BadCmp"); πE != nil {
								continue
							}
							if πTemp002, πE = µBadCmp.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs2 = πTemp003
							// line 44: self.assertRaises(Exc, cmp, s1, s2)
							πF.SetLineno(44)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
								continue
							}
							πTemp001[0] = µExc
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcmp); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							πTemp001[2] = µs1
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							πTemp001[3] = µs2
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
							// line 45: self.assertEqual(s1, s1)
							πF.SetLineno(45)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							πTemp001[0] = µs1
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							πTemp001[1] = µs1
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
							// line 47: s1 = slice(1, BadCmp())
							πF.SetLineno(47)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µBadCmp, "BadCmp"); πE != nil {
								continue
							}
							if πTemp002, πE = µBadCmp.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs1 = πTemp003
							// line 48: s2 = slice(1, BadCmp())
							πF.SetLineno(48)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µBadCmp, "BadCmp"); πE != nil {
								continue
							}
							if πTemp002, πE = µBadCmp.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs2 = πTemp003
							// line 49: self.assertEqual(s1, s1)
							πF.SetLineno(49)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							πTemp001[0] = µs1
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							πTemp001[1] = µs1
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
							// line 50: self.assertRaises(Exc, cmp, s1, s2)
							πF.SetLineno(50)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
								continue
							}
							πTemp001[0] = µExc
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcmp); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							πTemp001[2] = µs1
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							πTemp001[3] = µs2
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
							// line 52: s1 = slice(1, 2, BadCmp())
							πF.SetLineno(52)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = πg.NewInt(1).ToObject()
							πTemp001[1] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µBadCmp, "BadCmp"); πE != nil {
								continue
							}
							if πTemp002, πE = µBadCmp.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs1 = πTemp003
							// line 53: s2 = slice(1, 2, BadCmp())
							πF.SetLineno(53)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = πg.NewInt(1).ToObject()
							πTemp001[1] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µBadCmp, "BadCmp"); πE != nil {
								continue
							}
							if πTemp002, πE = µBadCmp.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs2 = πTemp003
							// line 54: self.assertEqual(s1, s1)
							πF.SetLineno(54)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							πTemp001[0] = µs1
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							πTemp001[1] = µs1
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
							// line 55: self.assertRaises(Exc, cmp, s1, s2)
							πF.SetLineno(55)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
								continue
							}
							πTemp001[0] = µExc
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcmp); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							πTemp001[2] = µs1
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							πTemp001[3] = µs2
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
					if πE = πClass.SetItem(πF, ßtest_cmp.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 26: @unittest.expectedFailure
					πF.SetLineno(26)
					πTemp006 = πF.MakeArgs(1)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßtest_cmp); πE != nil {
						continue
					}
					πTemp006[0] = πTemp007
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp008.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßtest_cmp.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 57: def test_members(self):
					πF.SetLineno(57)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_members", "build/src/__python__/test/test_slice.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µAnyClass *πg.Object = πg.UnboundLocal; _ = µAnyClass
						var µobj *πg.Object = πg.UnboundLocal; _ = µobj
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Dict
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
							// line 58: s = slice(1)
							πF.SetLineno(58)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(1).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp003
							// line 59: self.assertEqual(s.start, None)
							πF.SetLineno(59)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßstart, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							// line 60: self.assertEqual(s.stop, 1)
							πF.SetLineno(60)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßstop, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
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
							// line 61: self.assertEqual(s.step, None)
							πF.SetLineno(61)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßstep, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							// line 63: s = slice(1, 2)
							πF.SetLineno(63)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(1).ToObject()
							πTemp001[1] = πg.NewInt(2).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp003
							// line 64: self.assertEqual(s.start, 1)
							πF.SetLineno(64)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßstart, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
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
							// line 65: self.assertEqual(s.stop, 2)
							πF.SetLineno(65)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßstop, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(2).ToObject()
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
							// line 66: self.assertEqual(s.step, None)
							πF.SetLineno(66)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßstep, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							// line 68: s = slice(1, 2, 3)
							πF.SetLineno(68)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = πg.NewInt(1).ToObject()
							πTemp001[1] = πg.NewInt(2).ToObject()
							πTemp001[2] = πg.NewInt(3).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp003
							// line 69: self.assertEqual(s.start, 1)
							πF.SetLineno(69)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßstart, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
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
							// line 70: self.assertEqual(s.stop, 2)
							πF.SetLineno(70)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßstop, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(2).ToObject()
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
							// line 71: self.assertEqual(s.step, 3)
							πF.SetLineno(71)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßstep, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
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
							// line 73: class AnyClass(object):
							πF.SetLineno(73)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							πTemp004 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("AnyClass", "build/src/__python__/test/test_slice.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp004
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 74: pass
									πF.SetLineno(74)
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
							if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("AnyClass").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
								continue
							}
							µAnyClass = πTemp005
							// line 76: obj = AnyClass()
							πF.SetLineno(76)
							if πE = πg.CheckLocal(πF, µAnyClass, "AnyClass"); πE != nil {
								continue
							}
							if πTemp002, πE = µAnyClass.Call(πF, nil, nil); πE != nil {
								continue
							}
							µobj = πTemp002
							// line 77: s = slice(obj)
							πF.SetLineno(77)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp003
							// line 78: self.assertTrue(s.stop is obj)
							πF.SetLineno(78)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µs, ßstop, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp003 == µobj).ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_members.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 81: def test_indices(self):
					πF.SetLineno(81)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("test_indices", "build/src/__python__/test/test_slice.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 82: self.assertEqual(slice(None           ).indices(10), (0, 10,  1))
							πF.SetLineno(82)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(10).ToObject()
							πTemp003 = πF.MakeArgs(1)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßindices, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							πTemp004 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewInt(10).ToObject(), πg.NewInt(1).ToObject()).ToObject()
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
							// line 83: self.assertEqual(slice(None,  None,  2).indices(10), (0, 10,  2))
							πF.SetLineno(83)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(10).ToObject()
							πTemp003 = πF.MakeArgs(3)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							πTemp003[2] = πg.NewInt(2).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßindices, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							πTemp004 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewInt(10).ToObject(), πg.NewInt(2).ToObject()).ToObject()
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
							// line 84: self.assertEqual(slice(1,     None,  2).indices(10), (1, 10,  2))
							πF.SetLineno(84)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(10).ToObject()
							πTemp003 = πF.MakeArgs(3)
							πTemp003[0] = πg.NewInt(1).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							πTemp003[2] = πg.NewInt(2).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßindices, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							πTemp004 = πg.NewTuple3(πg.NewInt(1).ToObject(), πg.NewInt(10).ToObject(), πg.NewInt(2).ToObject()).ToObject()
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
							// line 85: self.assertEqual(slice(None,  None, -1).indices(10), (9, -1, -1))
							πF.SetLineno(85)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(10).ToObject()
							πTemp003 = πF.MakeArgs(3)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[2] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßindices, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πg.NewInt(9).ToObject(), πTemp005, πTemp006).ToObject()
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
							// line 86: self.assertEqual(slice(None,  None, -2).indices(10), (9, -1, -2))
							πF.SetLineno(86)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(10).ToObject()
							πTemp003 = πF.MakeArgs(3)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp003[2] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßindices, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πg.NewInt(9).ToObject(), πTemp005, πTemp006).ToObject()
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
							// line 87: self.assertEqual(slice(3,     None, -2).indices(10), (3, -1, -2))
							πF.SetLineno(87)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(10).ToObject()
							πTemp003 = πF.MakeArgs(3)
							πTemp003[0] = πg.NewInt(3).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp003[2] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßindices, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πg.NewInt(3).ToObject(), πTemp005, πTemp006).ToObject()
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
							// line 89: self.assertEqual(slice(None, -9).indices(10), (0, 1, 1))
							πF.SetLineno(89)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(10).ToObject()
							πTemp003 = πF.MakeArgs(2)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(9).ToObject()); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßindices, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							πTemp004 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject(), πg.NewInt(1).ToObject()).ToObject()
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
							// line 90: self.assertEqual(slice(None, -10).indices(10), (0, 0, 1))
							πF.SetLineno(90)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(10).ToObject()
							πTemp003 = πF.MakeArgs(2)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(10).ToObject()); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßindices, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							πTemp004 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject()).ToObject()
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
							// line 91: self.assertEqual(slice(None, -11).indices(10), (0, 0, 1))
							πF.SetLineno(91)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(10).ToObject()
							πTemp003 = πF.MakeArgs(2)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(11).ToObject()); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßindices, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							πTemp004 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject()).ToObject()
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
							// line 92: self.assertEqual(slice(None, -10, -1).indices(10), (9, 0, -1))
							πF.SetLineno(92)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(10).ToObject()
							πTemp003 = πF.MakeArgs(3)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(10).ToObject()); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[2] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßindices, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πg.NewInt(9).ToObject(), πg.NewInt(0).ToObject(), πTemp005).ToObject()
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
							// line 93: self.assertEqual(slice(None, -11, -1).indices(10), (9, -1, -1))
							πF.SetLineno(93)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(10).ToObject()
							πTemp003 = πF.MakeArgs(3)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(11).ToObject()); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[2] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßindices, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πg.NewInt(9).ToObject(), πTemp005, πTemp006).ToObject()
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
							// line 94: self.assertEqual(slice(None, -12, -1).indices(10), (9, -1, -1))
							πF.SetLineno(94)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(10).ToObject()
							πTemp003 = πF.MakeArgs(3)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(12).ToObject()); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[2] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßindices, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πg.NewInt(9).ToObject(), πTemp005, πTemp006).ToObject()
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
							// line 95: self.assertEqual(slice(None, 9).indices(10), (0, 9, 1))
							πF.SetLineno(95)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(10).ToObject()
							πTemp003 = πF.MakeArgs(2)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							πTemp003[1] = πg.NewInt(9).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßindices, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							πTemp004 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewInt(9).ToObject(), πg.NewInt(1).ToObject()).ToObject()
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
							// line 96: self.assertEqual(slice(None, 10).indices(10), (0, 10, 1))
							πF.SetLineno(96)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(10).ToObject()
							πTemp003 = πF.MakeArgs(2)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							πTemp003[1] = πg.NewInt(10).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßindices, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							πTemp004 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewInt(10).ToObject(), πg.NewInt(1).ToObject()).ToObject()
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
							// line 97: self.assertEqual(slice(None, 11).indices(10), (0, 10, 1))
							πF.SetLineno(97)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(10).ToObject()
							πTemp003 = πF.MakeArgs(2)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							πTemp003[1] = πg.NewInt(11).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßindices, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							πTemp004 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewInt(10).ToObject(), πg.NewInt(1).ToObject()).ToObject()
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
							// line 98: self.assertEqual(slice(None, 8, -1).indices(10), (9, 8, -1))
							πF.SetLineno(98)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(10).ToObject()
							πTemp003 = πF.MakeArgs(3)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							πTemp003[1] = πg.NewInt(8).ToObject()
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[2] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßindices, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πg.NewInt(9).ToObject(), πg.NewInt(8).ToObject(), πTemp005).ToObject()
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
							// line 99: self.assertEqual(slice(None, 9, -1).indices(10), (9, 9, -1))
							πF.SetLineno(99)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(10).ToObject()
							πTemp003 = πF.MakeArgs(3)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							πTemp003[1] = πg.NewInt(9).ToObject()
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[2] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßindices, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πg.NewInt(9).ToObject(), πg.NewInt(9).ToObject(), πTemp005).ToObject()
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
							// line 100: self.assertEqual(slice(None, 10, -1).indices(10), (9, 9, -1))
							πF.SetLineno(100)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(10).ToObject()
							πTemp003 = πF.MakeArgs(3)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							πTemp003[1] = πg.NewInt(10).ToObject()
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[2] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßindices, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πg.NewInt(9).ToObject(), πg.NewInt(9).ToObject(), πTemp005).ToObject()
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
							// line 102: self.assertEqual(
							πF.SetLineno(102)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(10).ToObject()
							πTemp003 = πF.MakeArgs(2)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(100).ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							πTemp003[1] = πg.NewInt(100).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßindices, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(10).ToObject()
							πTemp003 = πF.MakeArgs(1)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßindices, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp005
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
							// line 106: self.assertEqual(
							πF.SetLineno(106)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(10).ToObject()
							πTemp003 = πF.MakeArgs(3)
							πTemp003[0] = πg.NewInt(100).ToObject()
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(100).ToObject()); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[2] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßindices, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(10).ToObject()
							πTemp003 = πF.MakeArgs(3)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[2] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßindices, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp005
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
							// line 110: self.assertEqual(slice(-100L, 100L, 2L).indices(10), (0, 10,  2))
							πF.SetLineno(110)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(10).ToObject()
							πTemp003 = πF.MakeArgs(3)
							if πTemp004, πE = πg.Neg(πF, πg.NewLongFromBytes([]byte{0x64,}).ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							πTemp003[1] = πg.NewLongFromBytes([]byte{0x64,}).ToObject()
							πTemp003[2] = πg.NewLongFromBytes([]byte{0x2,}).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßindices, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							πTemp004 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewInt(10).ToObject(), πg.NewInt(2).ToObject()).ToObject()
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
							// line 112: self.assertEqual(range(10)[::sys.maxint - 1], [0])
							πF.SetLineno(112)
							πTemp001 = πF.MakeArgs(2)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßmaxint, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Sub(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πTemp005}, nil); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(10).ToObject()
							if πTemp006, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.GetItem(πF, πTemp007, πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							πTemp002 = make([]*πg.Object, 1)
							πTemp002[0] = πg.NewInt(0).ToObject()
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
							// line 114: self.assertRaises(OverflowError, slice(None).indices, 1L<<100)
							πF.SetLineno(114)
							πTemp001 = πF.MakeArgs(3)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßOverflowError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp002 = πF.MakeArgs(1)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßindices, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							if πTemp004, πE = πg.LShift(πF, πg.NewLongFromBytes([]byte{0x1,}).ToObject(), πg.NewInt(100).ToObject()); πE != nil {
								continue
							}
							πTemp001[2] = πTemp004
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
					if πE = πClass.SetItem(πF, ßtest_indices.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 80: @unittest.expectedFailure
					πF.SetLineno(80)
					πTemp006 = πF.MakeArgs(1)
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßtest_indices); πE != nil {
						continue
					}
					πTemp006[0] = πTemp009
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp010.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßtest_indices.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 117: def test_setslice_without_getslice(self):
					πF.SetLineno(117)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("test_setslice_without_getslice", "build/src/__python__/test/test_slice.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtmp *πg.Object = πg.UnboundLocal; _ = µtmp
						var µX *πg.Object = πg.UnboundLocal; _ = µX
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Dict
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 *πg.BaseException
						_ = πTemp009
						var πTemp010 *πg.Traceback
						_ = πTemp010
						var πTemp011 *πg.Type
						_ = πTemp011
						var πTemp012 []*πg.Object
						_ = πTemp012
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 118: tmp = []
							πF.SetLineno(118)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µtmp = πTemp002
							// line 119: class X(object):
							πF.SetLineno(119)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							πTemp003 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("X", "build/src/__python__/test/test_slice.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 120: def __setslice__(self, i, j, k):
									πF.SetLineno(120)
									πTemp002 = make([]πg.Param, 4)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "i", Def: nil}
									πTemp002[2] = πg.Param{Name: "j", Def: nil}
									πTemp002[3] = πg.Param{Name: "k", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__setslice__", "build/src/__python__/test/test_slice.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µi *πg.Object = πArgs[1]; _ = µi
										var µj *πg.Object = πArgs[2]; _ = µj
										var µk *πg.Object = πArgs[3]; _ = µk
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
											// line 121: tmp.append((i, j, k))
											πF.SetLineno(121)
											πTemp001 = πF.MakeArgs(1)
											if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
												continue
											}
											πTemp002 = πg.NewTuple3(µi, µj, µk).ToObject()
											πTemp001[0] = πTemp002
											if πE = πg.CheckLocal(πF, µtmp, "tmp"); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, µtmp, ßappend, nil); πE != nil {
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
									if πE = πClass.SetItem(πF, ß__setslice__.ToObject(), πTemp001); πE != nil {
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("X").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
								continue
							}
							µX = πTemp005
							// line 123: x = X()
							πF.SetLineno(123)
							if πE = πg.CheckLocal(πF, µX, "X"); πE != nil {
								continue
							}
							if πTemp002, πE = µX.Call(πF, nil, nil); πE != nil {
								continue
							}
							µx = πTemp002
							// line 124: with test_support.check_py3k_warnings():
							πF.SetLineno(124)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßcheck_py3k_warnings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp005.Call(πF, πg.Args{πTemp002}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(1)
							// line 125: x[1:2] = 42
							πF.SetLineno(125)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewInt(42).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, µx, πTemp007, πTemp006); πE != nil {
								continue
							}
							πF.PopCheckpoint()
						Label1:
							πTemp009, πTemp010 = nil, nil
							if πE != nil {
								πTemp009, πTemp010 = πF.ExcInfo()
							}
							if πTemp009 != nil {
								πTemp011 = πTemp009.Type()
								if πTemp006, πE = πTemp004.Call(πF, πg.Args{πTemp002, πTemp011.ToObject(), πTemp009.ToObject(), πTemp010.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp006, πE = πTemp004.Call(πF, πg.Args{πTemp002, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp009 != nil && πTemp008 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
								continue
							}
							// line 126: self.assertEqual(tmp, [(1, 2, 42)])
							πF.SetLineno(126)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtmp, "tmp"); πE != nil {
								continue
							}
							πTemp001[0] = µtmp
							πTemp012 = make([]*πg.Object, 1)
							πTemp002 = πg.NewTuple3(πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject(), πg.NewInt(42).ToObject()).ToObject()
							πTemp012[0] = πTemp002
							πTemp002 = πg.NewList(πTemp012...).ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_setslice_without_getslice.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 116: @unittest.expectedFailure
					πF.SetLineno(116)
					πTemp006 = πF.MakeArgs(1)
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßtest_setslice_without_getslice); πE != nil {
						continue
					}
					πTemp006[0] = πTemp010
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp011.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßtest_setslice_without_getslice.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 137: def test_cycle(self):
					πF.SetLineno(137)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("test_cycle", "build/src/__python__/test/test_slice.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmyobj *πg.Object = πg.UnboundLocal; _ = µmyobj
						var µo *πg.Object = πg.UnboundLocal; _ = µo
						var µw *πg.Object = πg.UnboundLocal; _ = µw
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 138: class myobj(object): pass
							πF.SetLineno(138)
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
							_, πE = πg.NewCode("myobj", "build/src/__python__/test/test_slice.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp001
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 138: class myobj(object): pass
									πF.SetLineno(138)
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("myobj").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µmyobj = πTemp005
							// line 139: o = myobj()
							πF.SetLineno(139)
							if πE = πg.CheckLocal(πF, µmyobj, "myobj"); πE != nil {
								continue
							}
							if πTemp002, πE = µmyobj.Call(πF, nil, nil); πE != nil {
								continue
							}
							µo = πTemp002
							// line 140: o.s = slice(o)
							πF.SetLineno(140)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp003[0] = µo
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µo, ßs, πTemp002); πE != nil {
								continue
							}
							// line 141: w = weakref.ref(o)
							πF.SetLineno(141)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp003[0] = µo
							if πTemp002, πE = πg.ResolveGlobal(πF, ßweakref); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßref, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µw = πTemp002
							// line 142: o = None
							πF.SetLineno(142)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µo = πTemp002
							// line 143: test_support.gc_collect()
							πF.SetLineno(143)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßgc_collect, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 144: self.assertIsNone(w())
							πF.SetLineno(144)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µw, "w"); πE != nil {
								continue
							}
							if πTemp002, πE = µw.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIsNone, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_cycle.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 136: @unittest.expectedFailure
					πF.SetLineno(136)
					πTemp006 = πF.MakeArgs(1)
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßtest_cycle); πE != nil {
						continue
					}
					πTemp006[0] = πTemp011
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp012.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßtest_cycle.ToObject(), πTemp011); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("SliceTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSliceTest.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 146: def test_main():
			πF.SetLineno(146)
			πTemp007 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("test_main", "build/src/__python__/test/test_slice.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 147: test_support.run_unittest(SliceTest)
					πF.SetLineno(147)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSliceTest); πE != nil {
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
			// line 149: if __name__ == "__main__":
			πF.SetLineno(149)
		Label1:
			// line 150: test_main()
			πF.SetLineno(150)
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
	πg.RegisterModule("test.test_slice", Code)
}
