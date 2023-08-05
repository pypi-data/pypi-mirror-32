package list_tests
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/test/list_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßCommonTest := πg.InternStr("CommonTest")
		ßException := πg.InternStr("Exception")
		ßFalse := πg.InternStr("False")
		ßIndexError := πg.InternStr("IndexError")
		ßKeyboardInterrupt := πg.InternStr("KeyboardInterrupt")
		ßNone := πg.InternStr("None")
		ßTypeError := πg.InternStr("TypeError")
		ßValueError := πg.InternStr("ValueError")
		ß__delitem__ := πg.InternStr("__delitem__")
		ß__eq__ := πg.InternStr("__eq__")
		ß__getitem__ := πg.InternStr("__getitem__")
		ß__iadd__ := πg.InternStr("__iadd__")
		ß__iter__ := πg.InternStr("__iter__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__setitem__ := πg.InternStr("__setitem__")
		ß_test_sort := πg.InternStr("_test_sort")
		ßa := πg.InternStr("a")
		ßabcdefghcij := πg.InternStr("abcdefghcij")
		ßabdefghcij := πg.InternStr("abdefghcij")
		ßabdefghij := πg.InternStr("abdefghij")
		ßappend := πg.InternStr("append")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertIs := πg.InternStr("assertIs")
		ßassertRaises := πg.InternStr("assertRaises")
		ßb := πg.InternStr("b")
		ßc := πg.InternStr("c")
		ßcmp := πg.InternStr("cmp")
		ßcount := πg.InternStr("count")
		ßeggs := πg.InternStr("eggs")
		ßfoo := πg.InternStr("foo")
		ßh := πg.InternStr("h")
		ßham := πg.InternStr("ham")
		ßid := πg.InternStr("id")
		ßinsert := πg.InternStr("insert")
		ßiter := πg.InternStr("iter")
		ßleft := πg.InternStr("left")
		ßlist := πg.InternStr("list")
		ßnext := πg.InternStr("next")
		ßobject := πg.InternStr("object")
		ßos := πg.InternStr("os")
		ßpop := πg.InternStr("pop")
		ßrange := πg.InternStr("range")
		ßremove := πg.InternStr("remove")
		ßreverse := πg.InternStr("reverse")
		ßright := πg.InternStr("right")
		ßs := πg.InternStr("s")
		ßseq_tests := πg.InternStr("seq_tests")
		ßslice := πg.InternStr("slice")
		ßsort := πg.InternStr("sort")
		ßspam := πg.InternStr("spam")
		ßspameggs := πg.InternStr("spameggs")
		ßsuper := πg.InternStr("super")
		ßsys := πg.InternStr("sys")
		ßtest_append := πg.InternStr("test_append")
		ßtest_constructor_exception_handling := πg.InternStr("test_constructor_exception_handling")
		ßtest_count := πg.InternStr("test_count")
		ßtest_delitem := πg.InternStr("test_delitem")
		ßtest_delslice := πg.InternStr("test_delslice")
		ßtest_exhausted_iterator := πg.InternStr("test_exhausted_iterator")
		ßtest_iadd := πg.InternStr("test_iadd")
		ßtest_imul := πg.InternStr("test_imul")
		ßtest_insert := πg.InternStr("test_insert")
		ßtest_pop := πg.InternStr("test_pop")
		ßtest_remove := πg.InternStr("test_remove")
		ßtest_reverse := πg.InternStr("test_reverse")
		ßtest_set_subscript := πg.InternStr("test_set_subscript")
		ßtest_setitem := πg.InternStr("test_setitem")
		ßtest_slice := πg.InternStr("test_slice")
		ßtest_support := πg.InternStr("test_support")
		ßtype2test := πg.InternStr("type2test")
		ßx := πg.InternStr("x")
		ßzip := πg.InternStr("zip")
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """
			πF.SetLineno(1)
			// line 5: import sys
			πF.SetLineno(5)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 6: import os
			πF.SetLineno(6)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 8: from test import test_support, seq_tests
			πF.SetLineno(8)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßtest_support.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "test.seq_tests"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßseq_tests.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 10: class CommonTest(seq_tests.CommonTest):
			πF.SetLineno(10)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßseq_tests); πE != nil {
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
			_, πE = πg.NewCode("CommonTest", "build/src/__python__/test/list_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 *πg.Object
				_ = πTemp012
				var πTemp013 *πg.Object
				_ = πTemp013
				var πTemp014 *πg.Object
				_ = πTemp014
				var πTemp015 *πg.Object
				_ = πTemp015
				var πTemp016 *πg.Object
				_ = πTemp016
				var πTemp017 *πg.Object
				_ = πTemp017
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 67: def test_set_subscript(self):
					πF.SetLineno(67)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_set_subscript", "build/src/__python__/test/list_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πg.UnboundLocal; _ = µa
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 68: a = self.type2test(range(20))
							πF.SetLineno(68)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(20).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
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
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 69: self.assertRaises(ValueError, a.__setitem__, slice(0, 10, 0), [1,2,3])
							πF.SetLineno(69)
							πTemp001 = πF.MakeArgs(4)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__setitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp002 = πF.MakeArgs(3)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(10).ToObject()
							πTemp002[2] = πg.NewInt(0).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[2] = πTemp004
							πTemp002 = make([]*πg.Object, 3)
							πTemp002[0] = πg.NewInt(1).ToObject()
							πTemp002[1] = πg.NewInt(2).ToObject()
							πTemp002[2] = πg.NewInt(3).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[3] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 70: self.assertRaises(TypeError, a.__setitem__, slice(0, 10), 1)
							πF.SetLineno(70)
							πTemp001 = πF.MakeArgs(4)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__setitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(10).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[2] = πTemp004
							πTemp001[3] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 71: self.assertRaises(ValueError, a.__setitem__, slice(0, 10, 2), [1,2])
							πF.SetLineno(71)
							πTemp001 = πF.MakeArgs(4)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__setitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp002 = πF.MakeArgs(3)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(10).ToObject()
							πTemp002[2] = πg.NewInt(2).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[2] = πTemp004
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(1).ToObject()
							πTemp002[1] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[3] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 72: self.assertRaises(TypeError, a.__getitem__, 'x', 1)
							πF.SetLineno(72)
							πTemp001 = πF.MakeArgs(4)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__getitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = ßx.ToObject()
							πTemp001[3] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 73: a[slice(2,10,3)] = [1,2,3]
							πF.SetLineno(73)
							πTemp001 = make([]*πg.Object, 3)
							πTemp001[0] = πg.NewInt(1).ToObject()
							πTemp001[1] = πg.NewInt(2).ToObject()
							πTemp001[2] = πg.NewInt(3).ToObject()
							πTemp003 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = πg.NewInt(2).ToObject()
							πTemp001[1] = πg.NewInt(10).ToObject()
							πTemp001[2] = πg.NewInt(3).ToObject()
							if πTemp006, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp005 = πTemp007
							if πE = πg.SetItem(πF, µa, πTemp005, πTemp004); πE != nil {
								continue
							}
							// line 74: self.assertEqual(a, self.type2test([0, 1, 1, 3, 4, 2, 6, 7, 3,
							πF.SetLineno(74)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = πF.MakeArgs(1)
							πTemp008 = make([]*πg.Object, 20)
							πTemp008[0] = πg.NewInt(0).ToObject()
							πTemp008[1] = πg.NewInt(1).ToObject()
							πTemp008[2] = πg.NewInt(1).ToObject()
							πTemp008[3] = πg.NewInt(3).ToObject()
							πTemp008[4] = πg.NewInt(4).ToObject()
							πTemp008[5] = πg.NewInt(2).ToObject()
							πTemp008[6] = πg.NewInt(6).ToObject()
							πTemp008[7] = πg.NewInt(7).ToObject()
							πTemp008[8] = πg.NewInt(3).ToObject()
							πTemp008[9] = πg.NewInt(9).ToObject()
							πTemp008[10] = πg.NewInt(10).ToObject()
							πTemp008[11] = πg.NewInt(11).ToObject()
							πTemp008[12] = πg.NewInt(12).ToObject()
							πTemp008[13] = πg.NewInt(13).ToObject()
							πTemp008[14] = πg.NewInt(14).ToObject()
							πTemp008[15] = πg.NewInt(15).ToObject()
							πTemp008[16] = πg.NewInt(16).ToObject()
							πTemp008[17] = πg.NewInt(17).ToObject()
							πTemp008[18] = πg.NewInt(18).ToObject()
							πTemp008[19] = πg.NewInt(19).ToObject()
							πTemp003 = πg.NewList(πTemp008...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_set_subscript.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 88: def test_setitem(self):
					πF.SetLineno(88)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_setitem", "build/src/__python__/test/list_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πg.UnboundLocal; _ = µa
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
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
							// line 89: a = self.type2test([0, 1])
							πF.SetLineno(89)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 90: a[0] = 0
							πF.SetLineno(90)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.SetItem(πF, µa, πTemp004, πTemp003); πE != nil {
								continue
							}
							// line 91: a[1] = 100
							πF.SetLineno(91)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewInt(100).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp004 = πg.NewInt(1).ToObject()
							if πE = πg.SetItem(πF, µa, πTemp004, πTemp003); πE != nil {
								continue
							}
							// line 92: self.assertEqual(a, self.type2test([0, 100]))
							πF.SetLineno(92)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(0).ToObject()
							πTemp005[1] = πg.NewInt(100).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 93: a[-1] = 200
							πF.SetLineno(93)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewInt(200).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πTemp006
							if πE = πg.SetItem(πF, µa, πTemp004, πTemp003); πE != nil {
								continue
							}
							// line 94: self.assertEqual(a, self.type2test([0, 200]))
							πF.SetLineno(94)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(0).ToObject()
							πTemp005[1] = πg.NewInt(200).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 95: a[-2] = 100
							πF.SetLineno(95)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewInt(100).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πTemp006
							if πE = πg.SetItem(πF, µa, πTemp004, πTemp003); πE != nil {
								continue
							}
							// line 96: self.assertEqual(a, self.type2test([100, 200]))
							πF.SetLineno(96)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(100).ToObject()
							πTemp005[1] = πg.NewInt(200).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 97: self.assertRaises(IndexError, a.__setitem__, -3, 200)
							πF.SetLineno(97)
							πTemp001 = πF.MakeArgs(4)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__setitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp001[2] = πTemp003
							πTemp001[3] = πg.NewInt(200).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 98: self.assertRaises(IndexError, a.__setitem__, 2, 200)
							πF.SetLineno(98)
							πTemp001 = πF.MakeArgs(4)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__setitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewInt(2).ToObject()
							πTemp001[3] = πg.NewInt(200).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 100: a = self.type2test([])
							πF.SetLineno(100)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 101: self.assertRaises(IndexError, a.__setitem__, 0, 200)
							πF.SetLineno(101)
							πTemp001 = πF.MakeArgs(4)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__setitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewInt(0).ToObject()
							πTemp001[3] = πg.NewInt(200).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 102: self.assertRaises(IndexError, a.__setitem__, -1, 200)
							πF.SetLineno(102)
							πTemp001 = πF.MakeArgs(4)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__setitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[2] = πTemp003
							πTemp001[3] = πg.NewInt(200).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 103: self.assertRaises(TypeError, a.__setitem__)
							πF.SetLineno(103)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__setitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 105: a = self.type2test([0,1,2,3,4])
							πF.SetLineno(105)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 5)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp002[2] = πg.NewInt(2).ToObject()
							πTemp002[3] = πg.NewInt(3).ToObject()
							πTemp002[4] = πg.NewInt(4).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 106: a[0L] = 1
							πF.SetLineno(106)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.SetItem(πF, µa, πTemp004, πTemp003); πE != nil {
								continue
							}
							// line 107: a[1L] = 2
							πF.SetLineno(107)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp004 = πg.NewLongFromBytes([]byte{0x1,}).ToObject()
							if πE = πg.SetItem(πF, µa, πTemp004, πTemp003); πE != nil {
								continue
							}
							// line 108: a[2L] = 3
							πF.SetLineno(108)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp004 = πg.NewLongFromBytes([]byte{0x2,}).ToObject()
							if πE = πg.SetItem(πF, µa, πTemp004, πTemp003); πE != nil {
								continue
							}
							// line 109: self.assertEqual(a, self.type2test([1,2,3,3,4]))
							πF.SetLineno(109)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 5)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(3).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp005[4] = πg.NewInt(4).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 110: a[0] = 5
							πF.SetLineno(110)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewInt(5).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.SetItem(πF, µa, πTemp004, πTemp003); πE != nil {
								continue
							}
							// line 111: a[1] = 6
							πF.SetLineno(111)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewInt(6).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp004 = πg.NewInt(1).ToObject()
							if πE = πg.SetItem(πF, µa, πTemp004, πTemp003); πE != nil {
								continue
							}
							// line 112: a[2] = 7
							πF.SetLineno(112)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewInt(7).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp004 = πg.NewInt(2).ToObject()
							if πE = πg.SetItem(πF, µa, πTemp004, πTemp003); πE != nil {
								continue
							}
							// line 113: self.assertEqual(a, self.type2test([5,6,7,3,4]))
							πF.SetLineno(113)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 5)
							πTemp005[0] = πg.NewInt(5).ToObject()
							πTemp005[1] = πg.NewInt(6).ToObject()
							πTemp005[2] = πg.NewInt(7).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp005[4] = πg.NewInt(4).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 114: a[-2L] = 88
							πF.SetLineno(114)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewInt(88).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewLongFromBytes([]byte{0x2,}).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πTemp006
							if πE = πg.SetItem(πF, µa, πTemp004, πTemp003); πE != nil {
								continue
							}
							// line 115: a[-1L] = 99
							πF.SetLineno(115)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewInt(99).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewLongFromBytes([]byte{0x1,}).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πTemp006
							if πE = πg.SetItem(πF, µa, πTemp004, πTemp003); πE != nil {
								continue
							}
							// line 116: self.assertEqual(a, self.type2test([5,6,7,88,99]))
							πF.SetLineno(116)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 5)
							πTemp005[0] = πg.NewInt(5).ToObject()
							πTemp005[1] = πg.NewInt(6).ToObject()
							πTemp005[2] = πg.NewInt(7).ToObject()
							πTemp005[3] = πg.NewInt(88).ToObject()
							πTemp005[4] = πg.NewInt(99).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 117: a[-2] = 8
							πF.SetLineno(117)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewInt(8).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πTemp006
							if πE = πg.SetItem(πF, µa, πTemp004, πTemp003); πE != nil {
								continue
							}
							// line 118: a[-1] = 9
							πF.SetLineno(118)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewInt(9).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πTemp006
							if πE = πg.SetItem(πF, µa, πTemp004, πTemp003); πE != nil {
								continue
							}
							// line 119: self.assertEqual(a, self.type2test([5,6,7,8,9]))
							πF.SetLineno(119)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 5)
							πTemp005[0] = πg.NewInt(5).ToObject()
							πTemp005[1] = πg.NewInt(6).ToObject()
							πTemp005[2] = πg.NewInt(7).ToObject()
							πTemp005[3] = πg.NewInt(8).ToObject()
							πTemp005[4] = πg.NewInt(9).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_setitem.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 121: def test_delitem(self):
					πF.SetLineno(121)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_delitem", "build/src/__python__/test/list_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πg.UnboundLocal; _ = µa
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
							// line 122: a = self.type2test([0, 1])
							πF.SetLineno(122)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 123: del a[1]
							πF.SetLineno(123)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp003 = πg.NewInt(1).ToObject()
							if πE = πg.DelItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							// line 124: self.assertEqual(a, [0])
							πF.SetLineno(124)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = make([]*πg.Object, 1)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 125: del a[0]
							πF.SetLineno(125)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.DelItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							// line 126: self.assertEqual(a, [])
							πF.SetLineno(126)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 128: a = self.type2test([0, 1])
							πF.SetLineno(128)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 129: del a[-2]
							πF.SetLineno(129)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.DelItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							// line 130: self.assertEqual(a, [1])
							πF.SetLineno(130)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = make([]*πg.Object, 1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 131: del a[-1]
							πF.SetLineno(131)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.DelItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							// line 132: self.assertEqual(a, [])
							πF.SetLineno(132)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 134: a = self.type2test([0, 1])
							πF.SetLineno(134)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 135: self.assertRaises(IndexError, a.__delitem__, -3)
							πF.SetLineno(135)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__delitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp001[2] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 136: self.assertRaises(IndexError, a.__delitem__, 2)
							πF.SetLineno(136)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__delitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 138: a = self.type2test([])
							πF.SetLineno(138)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 139: self.assertRaises(IndexError, a.__delitem__, 0)
							πF.SetLineno(139)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__delitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 141: self.assertRaises(TypeError, a.__delitem__)
							πF.SetLineno(141)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__delitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_delitem.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 191: def test_delslice(self):
					πF.SetLineno(191)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_delslice", "build/src/__python__/test/list_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πg.UnboundLocal; _ = µa
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
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
							// line 192: a = self.type2test([0, 1])
							πF.SetLineno(192)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 193: del a[1:2]
							πF.SetLineno(193)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							// line 194: del a[0:1]
							πF.SetLineno(194)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							// line 195: self.assertEqual(a, self.type2test([]))
							πF.SetLineno(195)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 197: a = self.type2test([0, 1])
							πF.SetLineno(197)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 198: del a[1L:2L]
							πF.SetLineno(198)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewLongFromBytes([]byte{0x1,}).ToObject(), πg.NewLongFromBytes([]byte{0x2,}).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							// line 199: del a[0L:1L]
							πF.SetLineno(199)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(0).ToObject(), πg.NewLongFromBytes([]byte{0x1,}).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							// line 200: self.assertEqual(a, self.type2test([]))
							πF.SetLineno(200)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 202: a = self.type2test([0, 1])
							πF.SetLineno(202)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 203: del a[-2:-1]
							πF.SetLineno(203)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πTemp004, πTemp006, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							// line 204: self.assertEqual(a, self.type2test([1]))
							πF.SetLineno(204)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 206: a = self.type2test([0, 1])
							πF.SetLineno(206)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 207: del a[-2L:-1L]
							πF.SetLineno(207)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Neg(πF, πg.NewLongFromBytes([]byte{0x2,}).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewLongFromBytes([]byte{0x1,}).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πTemp004, πTemp006, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							// line 208: self.assertEqual(a, self.type2test([1]))
							πF.SetLineno(208)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 210: a = self.type2test([0, 1])
							πF.SetLineno(210)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 211: del a[1:]
							πF.SetLineno(211)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							// line 212: del a[:1]
							πF.SetLineno(212)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(1).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							// line 213: self.assertEqual(a, self.type2test([]))
							πF.SetLineno(213)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 215: a = self.type2test([0, 1])
							πF.SetLineno(215)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 216: del a[1L:]
							πF.SetLineno(216)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewLongFromBytes([]byte{0x1,}).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							// line 217: del a[:1L]
							πF.SetLineno(217)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewLongFromBytes([]byte{0x1,}).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							// line 218: self.assertEqual(a, self.type2test([]))
							πF.SetLineno(218)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 220: a = self.type2test([0, 1])
							πF.SetLineno(220)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 221: del a[-1:]
							πF.SetLineno(221)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πTemp004, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							// line 222: self.assertEqual(a, self.type2test([0]))
							πF.SetLineno(222)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewInt(0).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 224: a = self.type2test([0, 1])
							πF.SetLineno(224)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 225: del a[-1L:]
							πF.SetLineno(225)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Neg(πF, πg.NewLongFromBytes([]byte{0x1,}).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πTemp004, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							// line 226: self.assertEqual(a, self.type2test([0]))
							πF.SetLineno(226)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewInt(0).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 228: a = self.type2test([0, 1])
							πF.SetLineno(228)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 229: del a[:]
							πF.SetLineno(229)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							// line 230: self.assertEqual(a, self.type2test([]))
							πF.SetLineno(230)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_delslice.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 232: def test_append(self):
					πF.SetLineno(232)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_append", "build/src/__python__/test/list_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πg.UnboundLocal; _ = µa
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
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
							// line 233: a = self.type2test([])
							πF.SetLineno(233)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 234: a.append(0)
							πF.SetLineno(234)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 235: a.append(1)
							πF.SetLineno(235)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 236: a.append(2)
							πF.SetLineno(236)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 237: self.assertEqual(a, self.type2test([0, 1, 2]))
							πF.SetLineno(237)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(0).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 239: self.assertRaises(TypeError, a.append)
							πF.SetLineno(239)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßappend, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_append.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 262: def test_insert(self):
					πF.SetLineno(262)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_insert", "build/src/__python__/test/list_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πg.UnboundLocal; _ = µa
						var µb *πg.Object = πg.UnboundLocal; _ = µb
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
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
							// line 263: a = self.type2test([0, 1, 2])
							πF.SetLineno(263)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 3)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp002[2] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 264: a.insert(0, -2)
							πF.SetLineno(264)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 265: a.insert(1, -1)
							πF.SetLineno(265)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(1).ToObject()
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 266: a.insert(2, 0)
							πF.SetLineno(266)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(2).ToObject()
							πTemp001[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 267: self.assertEqual(a, [-2, -1, 0, 0, 1, 2])
							πF.SetLineno(267)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = make([]*πg.Object, 6)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							πTemp002[2] = πg.NewInt(0).ToObject()
							πTemp002[3] = πg.NewInt(0).ToObject()
							πTemp002[4] = πg.NewInt(1).ToObject()
							πTemp002[5] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 269: b = a[:]
							πF.SetLineno(269)
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							µb = πTemp004
							// line 270: b.insert(-2, "foo")
							πF.SetLineno(270)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp001[1] = ßfoo.ToObject()
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µb, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 271: b.insert(-200, "left")
							πF.SetLineno(271)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(200).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp001[1] = ßleft.ToObject()
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µb, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 272: b.insert(200, "right")
							πF.SetLineno(272)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(200).ToObject()
							πTemp001[1] = ßright.ToObject()
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µb, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 273: self.assertEqual(b, self.type2test(["left",-2,-1,0,0,"foo",1,2,"right"]))
							πF.SetLineno(273)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							πTemp001[0] = µb
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 9)
							πTemp005[0] = ßleft.ToObject()
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp005[1] = πTemp003
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005[2] = πTemp003
							πTemp005[3] = πg.NewInt(0).ToObject()
							πTemp005[4] = πg.NewInt(0).ToObject()
							πTemp005[5] = ßfoo.ToObject()
							πTemp005[6] = πg.NewInt(1).ToObject()
							πTemp005[7] = πg.NewInt(2).ToObject()
							πTemp005[8] = ßright.ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 275: self.assertRaises(TypeError, a.insert)
							πF.SetLineno(275)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßinsert, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_insert.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 277: def test_pop(self):
					πF.SetLineno(277)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("test_pop", "build/src/__python__/test/list_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πg.UnboundLocal; _ = µa
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
							// line 278: a = self.type2test([-1, 0, 1])
							πF.SetLineno(278)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 3)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewInt(0).ToObject()
							πTemp002[2] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 279: a.pop()
							πF.SetLineno(279)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßpop, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 280: self.assertEqual(a, [-1, 0])
							πF.SetLineno(280)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = make([]*πg.Object, 2)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewInt(0).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 281: a.pop(0)
							πF.SetLineno(281)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßpop, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 282: self.assertEqual(a, [0])
							πF.SetLineno(282)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = make([]*πg.Object, 1)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 283: self.assertRaises(IndexError, a.pop, 5)
							πF.SetLineno(283)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßpop, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewInt(5).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 284: a.pop(0)
							πF.SetLineno(284)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßpop, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 285: self.assertEqual(a, [])
							πF.SetLineno(285)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 286: self.assertRaises(IndexError, a.pop)
							πF.SetLineno(286)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßpop, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 287: self.assertRaises(TypeError, a.pop, 42, 42)
							πF.SetLineno(287)
							πTemp001 = πF.MakeArgs(4)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßpop, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewInt(42).ToObject()
							πTemp001[3] = πg.NewInt(42).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 288: a = self.type2test([0, 10, 20, 30, 40])
							πF.SetLineno(288)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 5)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(10).ToObject()
							πTemp002[2] = πg.NewInt(20).ToObject()
							πTemp002[3] = πg.NewInt(30).ToObject()
							πTemp002[4] = πg.NewInt(40).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_pop.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 290: def test_remove(self):
					πF.SetLineno(290)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("test_remove", "build/src/__python__/test/list_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πg.UnboundLocal; _ = µa
						var µBadExc *πg.Object = πg.UnboundLocal; _ = µBadExc
						var µBadCmp *πg.Object = πg.UnboundLocal; _ = µBadCmp
						var µBadCmp2 *πg.Object = πg.UnboundLocal; _ = µBadCmp2
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µe *πg.Object = πg.UnboundLocal; _ = µe
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var µy *πg.Object = πg.UnboundLocal; _ = µy
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Dict
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 bool
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
							// line 291: a = self.type2test([0, 0, 1])
							πF.SetLineno(291)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 3)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(0).ToObject()
							πTemp002[2] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 292: a.remove(1)
							πF.SetLineno(292)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßremove, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 293: self.assertEqual(a, [0, 0])
							πF.SetLineno(293)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(0).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 294: a.remove(0)
							πF.SetLineno(294)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßremove, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 295: self.assertEqual(a, [0])
							πF.SetLineno(295)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = make([]*πg.Object, 1)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 296: a.remove(0)
							πF.SetLineno(296)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßremove, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 297: self.assertEqual(a, [])
							πF.SetLineno(297)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 299: self.assertRaises(ValueError, a.remove, 0)
							πF.SetLineno(299)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßremove, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 301: self.assertRaises(TypeError, a.remove)
							πF.SetLineno(301)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßremove, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 303: class BadExc(Exception):
							πF.SetLineno(303)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							πTemp005 = πg.NewDict()
							if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
								continue
							}
							_, πE = πg.NewCode("BadExc", "build/src/__python__/test/list_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp005
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 304: pass
									πF.SetLineno(304)
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp004 == nil {
								πTemp004 = πg.TypeType.ToObject()
							}
							if πTemp006, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("BadExc").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µBadExc = πTemp006
							// line 306: class BadCmp(object):
							πF.SetLineno(306)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							πTemp005 = πg.NewDict()
							if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
								continue
							}
							_, πE = πg.NewCode("BadCmp", "build/src/__python__/test/list_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp005
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
									// line 307: def __eq__(self, other):
									πF.SetLineno(307)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__eq__", "build/src/__python__/test/list_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µother *πg.Object = πArgs[1]; _ = µother
										var πTemp001 *πg.Object
										_ = πTemp001
										var πTemp002 bool
										_ = πTemp002
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.Eq(πF, µother, πg.NewInt(2).ToObject()); πE != nil {
												continue
											}
											if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
												continue
											}
											if πTemp002 {
												goto Label1
											}
											goto Label2
											// line 308: if other == 2:
											πF.SetLineno(308)
										Label1:
											if πE = πg.CheckLocal(πF, µBadExc, "BadExc"); πE != nil {
												continue
											}
											if πTemp001, πE = µBadExc.Call(πF, nil, nil); πE != nil {
												continue
											}
											// line 309: raise BadExc()
											πF.SetLineno(309)
											πE = πF.Raise(πTemp001, nil, nil)
											continue
											goto Label2
										Label2:
											// line 310: return False
											πF.SetLineno(310)
											if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
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
									if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp001); πE != nil {
										continue
									}
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp004 == nil {
								πTemp004 = πg.TypeType.ToObject()
							}
							if πTemp006, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("BadCmp").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µBadCmp = πTemp006
							// line 312: a = self.type2test([0, 1, 2, 3])
							πF.SetLineno(312)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 4)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp002[2] = πg.NewInt(2).ToObject()
							πTemp002[3] = πg.NewInt(3).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 313: self.assertRaises(BadExc, a.remove, BadCmp())
							πF.SetLineno(313)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µBadExc, "BadExc"); πE != nil {
								continue
							}
							πTemp001[0] = µBadExc
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßremove, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µBadCmp, "BadCmp"); πE != nil {
								continue
							}
							if πTemp003, πE = µBadCmp.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 315: class BadCmp2(object):
							πF.SetLineno(315)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							πTemp005 = πg.NewDict()
							if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
								continue
							}
							_, πE = πg.NewCode("BadCmp2", "build/src/__python__/test/list_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp005
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
									// line 316: def __eq__(self, other):
									πF.SetLineno(316)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__eq__", "build/src/__python__/test/list_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µother *πg.Object = πArgs[1]; _ = µother
										var πTemp001 *πg.Object
										_ = πTemp001
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											if πE = πg.CheckLocal(πF, µBadExc, "BadExc"); πE != nil {
												continue
											}
											if πTemp001, πE = µBadExc.Call(πF, nil, nil); πE != nil {
												continue
											}
											// line 317: raise BadExc()
											πF.SetLineno(317)
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
									if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp001); πE != nil {
										continue
									}
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp004 == nil {
								πTemp004 = πg.TypeType.ToObject()
							}
							if πTemp006, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("BadCmp2").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µBadCmp2 = πTemp006
							// line 319: d = self.type2test('abcdefghcij')
							πF.SetLineno(319)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßabcdefghcij.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µd = πTemp004
							// line 320: d.remove('c')
							πF.SetLineno(320)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßc.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µd, ßremove, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 321: self.assertEqual(d, self.type2test('abdefghcij'))
							πF.SetLineno(321)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[0] = µd
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßabdefghcij.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 322: d.remove('c')
							πF.SetLineno(322)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßc.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µd, ßremove, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 323: self.assertEqual(d, self.type2test('abdefghij'))
							πF.SetLineno(323)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[0] = µd
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßabdefghij.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 324: self.assertRaises(ValueError, d.remove, 'c')
							πF.SetLineno(324)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µd, ßremove, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = ßc.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 325: self.assertEqual(d, self.type2test('abdefghij'))
							πF.SetLineno(325)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[0] = µd
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßabdefghij.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 328: d = self.type2test(['a', 'b', BadCmp2(), 'c'])
							πF.SetLineno(328)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 4)
							πTemp002[0] = ßa.ToObject()
							πTemp002[1] = ßb.ToObject()
							if πE = πg.CheckLocal(πF, µBadCmp2, "BadCmp2"); πE != nil {
								continue
							}
							if πTemp003, πE = µBadCmp2.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[2] = πTemp003
							πTemp002[3] = ßc.ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µd = πTemp004
							// line 329: e = self.type2test(d)
							πF.SetLineno(329)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[0] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µe = πTemp004
							// line 330: self.assertRaises(BadExc, d.remove, 'c')
							πF.SetLineno(330)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µBadExc, "BadExc"); πE != nil {
								continue
							}
							πTemp001[0] = µBadExc
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µd, ßremove, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = ßc.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[0] = µd
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							πTemp001[1] = µe
							if πTemp004, πE = πg.ResolveGlobal(πF, ßzip); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.Iter(πF, πTemp006); πE != nil {
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
							if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp009}}}, πTemp004); πE != nil {
									continue
								}
								µx = πTemp006
								µy = πTemp009
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 333: self.assertIs(x, y)
							πF.SetLineno(333)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							πTemp001[1] = µy
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertIs, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßtest_remove.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 335: def test_count(self):
					πF.SetLineno(335)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("test_count", "build/src/__python__/test/list_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πg.UnboundLocal; _ = µa
						var µBadExc *πg.Object = πg.UnboundLocal; _ = µBadExc
						var µBadCmp *πg.Object = πg.UnboundLocal; _ = µBadCmp
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Dict
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 336: a = self.type2test([0, 1, 2])*3
							πF.SetLineno(336)
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = make([]*πg.Object, 3)
							πTemp003[0] = πg.NewInt(0).ToObject()
							πTemp003[1] = πg.NewInt(1).ToObject()
							πTemp003[2] = πg.NewInt(2).ToObject()
							πTemp004 = πg.NewList(πTemp003...).ToObject()
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Mul(πF, πTemp005, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							µa = πTemp001
							// line 337: self.assertEqual(a.count(0), 3)
							πF.SetLineno(337)
							πTemp002 = πF.MakeArgs(2)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µa, ßcount, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002[0] = πTemp004
							πTemp002[1] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 338: self.assertEqual(a.count(1), 3)
							πF.SetLineno(338)
							πTemp002 = πF.MakeArgs(2)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µa, ßcount, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002[0] = πTemp004
							πTemp002[1] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 339: self.assertEqual(a.count(3), 0)
							πF.SetLineno(339)
							πTemp002 = πF.MakeArgs(2)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µa, ßcount, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002[0] = πTemp004
							πTemp002[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 341: self.assertRaises(TypeError, a.count)
							πF.SetLineno(341)
							πTemp002 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µa, ßcount, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 343: class BadExc(Exception):
							πF.SetLineno(343)
							πTemp002 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							πTemp002[0] = πTemp005
							πTemp006 = πg.NewDict()
							if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
								continue
							}
							_, πE = πg.NewCode("BadExc", "build/src/__python__/test/list_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp006
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 344: pass
									πF.SetLineno(344)
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp004, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp004 == nil {
								πTemp004 = πg.TypeType.ToObject()
							}
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("BadExc").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
								continue
							}
							µBadExc = πTemp005
							// line 346: class BadCmp(object):
							πF.SetLineno(346)
							πTemp002 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp002[0] = πTemp005
							πTemp006 = πg.NewDict()
							if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
								continue
							}
							_, πE = πg.NewCode("BadCmp", "build/src/__python__/test/list_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp006
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
									// line 347: def __eq__(self, other):
									πF.SetLineno(347)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__eq__", "build/src/__python__/test/list_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µother *πg.Object = πArgs[1]; _ = µother
										var πTemp001 *πg.Object
										_ = πTemp001
										var πTemp002 bool
										_ = πTemp002
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.Eq(πF, µother, πg.NewInt(2).ToObject()); πE != nil {
												continue
											}
											if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
												continue
											}
											if πTemp002 {
												goto Label1
											}
											goto Label2
											// line 348: if other == 2:
											πF.SetLineno(348)
										Label1:
											if πE = πg.CheckLocal(πF, µBadExc, "BadExc"); πE != nil {
												continue
											}
											if πTemp001, πE = µBadExc.Call(πF, nil, nil); πE != nil {
												continue
											}
											// line 349: raise BadExc()
											πF.SetLineno(349)
											πE = πF.Raise(πTemp001, nil, nil)
											continue
											goto Label2
										Label2:
											// line 350: return False
											πF.SetLineno(350)
											if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
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
									if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp001); πE != nil {
										continue
									}
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp004, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp004 == nil {
								πTemp004 = πg.TypeType.ToObject()
							}
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("BadCmp").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
								continue
							}
							µBadCmp = πTemp005
							// line 352: self.assertRaises(BadExc, a.count, BadCmp())
							πF.SetLineno(352)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µBadExc, "BadExc"); πE != nil {
								continue
							}
							πTemp002[0] = µBadExc
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µa, ßcount, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πE = πg.CheckLocal(πF, µBadCmp, "BadCmp"); πE != nil {
								continue
							}
							if πTemp001, πE = µBadCmp.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[2] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_count.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 411: def test_reverse(self):
					πF.SetLineno(411)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("test_reverse", "build/src/__python__/test/list_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µu *πg.Object = πg.UnboundLocal; _ = µu
						var µu2 *πg.Object = πg.UnboundLocal; _ = µu2
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
							// line 412: u = self.type2test([-2, -1, 0, 1, 2])
							πF.SetLineno(412)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 5)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							πTemp002[2] = πg.NewInt(0).ToObject()
							πTemp002[3] = πg.NewInt(1).ToObject()
							πTemp002[4] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µu = πTemp004
							// line 413: u2 = u[:]
							πF.SetLineno(413)
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µu, πTemp003); πE != nil {
								continue
							}
							µu2 = πTemp004
							// line 414: u.reverse()
							πF.SetLineno(414)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µu, ßreverse, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 415: self.assertEqual(u, [2, 1, 0, -1, -2])
							πF.SetLineno(415)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001[0] = µu
							πTemp002 = make([]*πg.Object, 5)
							πTemp002[0] = πg.NewInt(2).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp002[2] = πg.NewInt(0).ToObject()
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002[3] = πTemp003
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002[4] = πTemp003
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 416: u.reverse()
							πF.SetLineno(416)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µu, ßreverse, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 417: self.assertEqual(u, u2)
							πF.SetLineno(417)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001[0] = µu
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							πTemp001[1] = µu2
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 419: self.assertRaises(TypeError, u.reverse, 42)
							πF.SetLineno(419)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µu, ßreverse, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewInt(42).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_reverse.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 426: def _test_sort(self):
					πF.SetLineno(426)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("_test_sort", "build/src/__python__/test/list_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µu *πg.Object = πg.UnboundLocal; _ = µu
						var µrevcmp *πg.Object = πg.UnboundLocal; _ = µrevcmp
						var µmyComparison *πg.Object = πg.UnboundLocal; _ = µmyComparison
						var µz *πg.Object = πg.UnboundLocal; _ = µz
						var µselfmodifyingComparison *πg.Object = πg.UnboundLocal; _ = µselfmodifyingComparison
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 []πg.Param
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 427: u = self.type2test([1, 0])
							πF.SetLineno(427)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(1).ToObject()
							πTemp002[1] = πg.NewInt(0).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µu = πTemp004
							// line 428: u.sort()
							πF.SetLineno(428)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µu, ßsort, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 429: self.assertEqual(u, [0, 1])
							πF.SetLineno(429)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001[0] = µu
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 431: u = self.type2test([2,1,0,-1,-2])
							πF.SetLineno(431)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 5)
							πTemp002[0] = πg.NewInt(2).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp002[2] = πg.NewInt(0).ToObject()
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002[3] = πTemp003
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002[4] = πTemp003
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µu = πTemp004
							// line 432: u.sort()
							πF.SetLineno(432)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µu, ßsort, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 433: self.assertEqual(u, self.type2test([-2,-1,0,1,2]))
							πF.SetLineno(433)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001[0] = µu
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 5)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005[1] = πTemp003
							πTemp005[2] = πg.NewInt(0).ToObject()
							πTemp005[3] = πg.NewInt(1).ToObject()
							πTemp005[4] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 435: self.assertRaises(TypeError, u.sort, 42, 42)
							πF.SetLineno(435)
							πTemp001 = πF.MakeArgs(4)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µu, ßsort, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewInt(42).ToObject()
							πTemp001[3] = πg.NewInt(42).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 437: def revcmp(a, b):
							πF.SetLineno(437)
							πTemp006 = make([]πg.Param, 2)
							πTemp006[0] = πg.Param{Name: "a", Def: nil}
							πTemp006[1] = πg.Param{Name: "b", Def: nil}
							πTemp003 = πg.NewFunction(πg.NewCode("revcmp", "build/src/__python__/test/list_tests.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µa *πg.Object = πArgs[0]; _ = µa
								var µb *πg.Object = πArgs[1]; _ = µb
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
									// line 438: return cmp(b, a)
									πF.SetLineno(438)
									πTemp001 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
										continue
									}
									πTemp001[0] = µb
									if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
										continue
									}
									πTemp001[1] = µa
									if πTemp002, πE = πg.ResolveGlobal(πF, ßcmp); πE != nil {
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
							µrevcmp = πTemp003
							// line 439: u.sort(revcmp)
							πF.SetLineno(439)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrevcmp, "revcmp"); πE != nil {
								continue
							}
							πTemp001[0] = µrevcmp
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µu, ßsort, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 440: self.assertEqual(u, self.type2test([2,1,0,-1,-2]))
							πF.SetLineno(440)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001[0] = µu
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 5)
							πTemp005[0] = πg.NewInt(2).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(0).ToObject()
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005[3] = πTemp004
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp005[4] = πTemp004
							πTemp004 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 443: def myComparison(x,y):
							πF.SetLineno(443)
							πTemp006 = make([]πg.Param, 2)
							πTemp006[0] = πg.Param{Name: "x", Def: nil}
							πTemp006[1] = πg.Param{Name: "y", Def: nil}
							πTemp004 = πg.NewFunction(πg.NewCode("myComparison", "build/src/__python__/test/list_tests.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µx *πg.Object = πArgs[0]; _ = µx
								var µy *πg.Object = πArgs[1]; _ = µy
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
									// line 444: return cmp(x%3, y%7)
									πF.SetLineno(444)
									πTemp001 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.Mod(πF, µx, πg.NewInt(3).ToObject()); πE != nil {
										continue
									}
									πTemp001[0] = πTemp002
									if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.Mod(πF, µy, πg.NewInt(7).ToObject()); πE != nil {
										continue
									}
									πTemp001[1] = πTemp002
									if πTemp002, πE = πg.ResolveGlobal(πF, ßcmp); πE != nil {
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
							µmyComparison = πTemp004
							// line 445: z = self.type2test(range(12))
							πF.SetLineno(445)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(12).ToObject()
							if πTemp007, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp008
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µz = πTemp008
							// line 446: z.sort(myComparison)
							πF.SetLineno(446)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmyComparison, "myComparison"); πE != nil {
								continue
							}
							πTemp001[0] = µmyComparison
							if πE = πg.CheckLocal(πF, µz, "z"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µz, ßsort, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 448: self.assertRaises(TypeError, z.sort, 2)
							πF.SetLineno(448)
							πTemp001 = πF.MakeArgs(3)
							if πTemp007, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp007
							if πE = πg.CheckLocal(πF, µz, "z"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µz, ßsort, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp007
							πTemp001[2] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 450: def selfmodifyingComparison(x,y):
							πF.SetLineno(450)
							πTemp006 = make([]πg.Param, 2)
							πTemp006[0] = πg.Param{Name: "x", Def: nil}
							πTemp006[1] = πg.Param{Name: "y", Def: nil}
							πTemp007 = πg.NewFunction(πg.NewCode("selfmodifyingComparison", "build/src/__python__/test/list_tests.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µx *πg.Object = πArgs[0]; _ = µx
								var µy *πg.Object = πArgs[1]; _ = µy
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
									// line 451: z.append(1)
									πF.SetLineno(451)
									πTemp001 = πF.MakeArgs(1)
									πTemp001[0] = πg.NewInt(1).ToObject()
									if πE = πg.CheckLocal(πF, µz, "z"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µz, ßappend, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									// line 452: return cmp(x, y)
									πF.SetLineno(452)
									πTemp001 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
										continue
									}
									πTemp001[0] = µx
									if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
										continue
									}
									πTemp001[1] = µy
									if πTemp002, πE = πg.ResolveGlobal(πF, ßcmp); πE != nil {
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
							µselfmodifyingComparison = πTemp007
							// line 453: self.assertRaises(ValueError, z.sort, selfmodifyingComparison)
							πF.SetLineno(453)
							πTemp001 = πF.MakeArgs(3)
							if πTemp008, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp008
							if πE = πg.CheckLocal(πF, µz, "z"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µz, ßsort, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp008
							if πE = πg.CheckLocal(πF, µselfmodifyingComparison, "selfmodifyingComparison"); πE != nil {
								continue
							}
							πTemp001[2] = µselfmodifyingComparison
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 455: self.assertRaises(TypeError, z.sort, lambda x, y: 's')
							πF.SetLineno(455)
							πTemp001 = πF.MakeArgs(3)
							if πTemp008, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp008
							if πE = πg.CheckLocal(πF, µz, "z"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µz, ßsort, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp008
							πTemp006 = make([]πg.Param, 2)
							πTemp006[0] = πg.Param{Name: "x", Def: nil}
							πTemp006[1] = πg.Param{Name: "y", Def: nil}
							πTemp008 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/test/list_tests.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µx *πg.Object = πArgs[0]; _ = µx
								var µy *πg.Object = πArgs[1]; _ = µy
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 455: self.assertRaises(TypeError, z.sort, lambda x, y: 's')
									πF.SetLineno(455)
									πR = ßs.ToObject()
									continue
								}
								if πE != nil {
									πR = nil
								} else if πR == nil {
									πR = πg.None
								}
								return πR, πE
							}), πF.Globals()).ToObject()
							πTemp001[2] = πTemp008
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 457: self.assertRaises(TypeError, z.sort, 42, 42, 42, 42)
							πF.SetLineno(457)
							πTemp001 = πF.MakeArgs(6)
							if πTemp008, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp008
							if πE = πg.CheckLocal(πF, µz, "z"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µz, ßsort, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp008
							πTemp001[2] = πg.NewInt(42).ToObject()
							πTemp001[3] = πg.NewInt(42).ToObject()
							πTemp001[4] = πg.NewInt(42).ToObject()
							πTemp001[5] = πg.NewInt(42).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_test_sort.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 459: def test_slice(self):
					πF.SetLineno(459)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("test_slice", "build/src/__python__/test/list_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µu *πg.Object = πg.UnboundLocal; _ = µu
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
							// line 460: u = self.type2test("spam")
							πF.SetLineno(460)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßspam.ToObject()
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
							µu = πTemp003
							// line 461: u[:2] = "h"
							πF.SetLineno(461)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ßh.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(2).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, µu, πTemp003, πTemp002); πE != nil {
								continue
							}
							// line 462: self.assertEqual(u, list("ham"))
							πF.SetLineno(462)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001[0] = µu
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßham.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[1] = πTemp003
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
					if πE = πClass.SetItem(πF, ßtest_slice.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 464: def test_iadd(self):
					πF.SetLineno(464)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("test_iadd", "build/src/__python__/test/list_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µu *πg.Object = πg.UnboundLocal; _ = µu
						var µu2 *πg.Object = πg.UnboundLocal; _ = µu2
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
							// line 465: super(CommonTest, self).test_iadd()
							πF.SetLineno(465)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßCommonTest); πE != nil {
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
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßtest_iadd, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 466: u = self.type2test([0, 1])
							πF.SetLineno(466)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = make([]*πg.Object, 2)
							πTemp004[0] = πg.NewInt(0).ToObject()
							πTemp004[1] = πg.NewInt(1).ToObject()
							πTemp002 = πg.NewList(πTemp004...).ToObject()
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
							µu = πTemp003
							// line 467: u2 = u
							πF.SetLineno(467)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							µu2 = µu
							// line 468: u += [2, 3]
							πF.SetLineno(468)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001 = make([]*πg.Object, 2)
							πTemp001[0] = πg.NewInt(2).ToObject()
							πTemp001[1] = πg.NewInt(3).ToObject()
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πTemp003, πE = πg.IAdd(πF, µu, πTemp002); πE != nil {
								continue
							}
							µu = πTemp003
							// line 469: self.assertIs(u, u2)
							πF.SetLineno(469)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001[0] = µu
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							πTemp001[1] = µu2
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIs, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 471: u = self.type2test("spam")
							πF.SetLineno(471)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßspam.ToObject()
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
							µu = πTemp003
							// line 472: u += "eggs"
							πF.SetLineno(472)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µu, ßeggs.ToObject()); πE != nil {
								continue
							}
							µu = πTemp002
							// line 473: self.assertEqual(u, self.type2test("spameggs"))
							πF.SetLineno(473)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001[0] = µu
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßspameggs.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[1] = πTemp003
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
							// line 475: self.assertRaises(TypeError, u.__iadd__, None)
							πF.SetLineno(475)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µu, ß__iadd__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
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
					if πE = πClass.SetItem(πF, ßtest_iadd.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 477: def test_imul(self):
					πF.SetLineno(477)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("test_imul", "build/src/__python__/test/list_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µu *πg.Object = πg.UnboundLocal; _ = µu
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µoldid *πg.Object = πg.UnboundLocal; _ = µoldid
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
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
							// line 478: u = self.type2test([0, 1])
							πF.SetLineno(478)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µu = πTemp004
							// line 479: u *= 3
							πF.SetLineno(479)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IMul(πF, µu, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							µu = πTemp003
							// line 480: self.assertEqual(u, self.type2test([0, 1, 0, 1, 0, 1]))
							πF.SetLineno(480)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001[0] = µu
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 6)
							πTemp005[0] = πg.NewInt(0).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(0).ToObject()
							πTemp005[3] = πg.NewInt(1).ToObject()
							πTemp005[4] = πg.NewInt(0).ToObject()
							πTemp005[5] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 481: u *= 0
							πF.SetLineno(481)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IMul(πF, µu, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							µu = πTemp003
							// line 482: self.assertEqual(u, self.type2test([]))
							πF.SetLineno(482)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001[0] = µu
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 483: s = self.type2test([])
							πF.SetLineno(483)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp004
							// line 484: oldid = id(s)
							πF.SetLineno(484)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µoldid = πTemp004
							// line 485: s *= 10
							πF.SetLineno(485)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IMul(πF, µs, πg.NewInt(10).ToObject()); πE != nil {
								continue
							}
							µs = πTemp003
							// line 486: self.assertEqual(id(s), oldid)
							πF.SetLineno(486)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp002[0] = µs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µoldid, "oldid"); πE != nil {
								continue
							}
							πTemp001[1] = µoldid
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_imul.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 529: def test_constructor_exception_handling(self):
					πF.SetLineno(529)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("test_constructor_exception_handling", "build/src/__python__/test/list_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µF *πg.Object = πg.UnboundLocal; _ = µF
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
							// line 531: class F(object):
							πF.SetLineno(531)
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
							_, πE = πg.NewCode("F", "build/src/__python__/test/list_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp001
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
									// line 532: def __iter__(self):
									πF.SetLineno(532)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__iter__", "build/src/__python__/test/list_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyboardInterrupt); πE != nil {
												continue
											}
											// line 533: raise KeyboardInterrupt
											πF.SetLineno(533)
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
									if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp001); πE != nil {
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("F").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µF = πTemp005
							// line 534: self.assertRaises(KeyboardInterrupt, list, F())
							πF.SetLineno(534)
							πTemp003 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyboardInterrupt); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µF, "F"); πE != nil {
								continue
							}
							if πTemp002, πE = µF.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[2] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_constructor_exception_handling.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 536: def test_exhausted_iterator(self):
					πF.SetLineno(536)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("test_exhausted_iterator", "build/src/__python__/test/list_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πg.UnboundLocal; _ = µa
						var µexhit *πg.Object = πg.UnboundLocal; _ = µexhit
						var µempit *πg.Object = πg.UnboundLocal; _ = µempit
						var µx *πg.Object = πg.UnboundLocal; _ = µx
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
							// line 537: a = self.type2test([1, 2, 3])
							πF.SetLineno(537)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 3)
							πTemp002[0] = πg.NewInt(1).ToObject()
							πTemp002[1] = πg.NewInt(2).ToObject()
							πTemp002[2] = πg.NewInt(3).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 538: exhit = iter(a)
							πF.SetLineno(538)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							if πTemp003, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µexhit = πTemp004
							// line 539: empit = iter(a)
							πF.SetLineno(539)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							if πTemp003, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µempit = πTemp004
							if πE = πg.CheckLocal(πF, µexhit, "exhit"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Iter(πF, µexhit); πE != nil {
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
							if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
							// line 541: next(empit)  # not exhausted
							πF.SetLineno(541)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µempit, "empit"); πE != nil {
								continue
							}
							πTemp001[0] = µempit
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnext); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 542: a.append(9)
							πF.SetLineno(542)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(9).ToObject()
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 543: self.assertEqual(list(exhit), [])
							πF.SetLineno(543)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µexhit, "exhit"); πE != nil {
								continue
							}
							πTemp002[0] = µexhit
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 544: self.assertEqual(list(empit), [9])
							πF.SetLineno(544)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µempit, "empit"); πE != nil {
								continue
							}
							πTemp002[0] = µempit
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp002 = make([]*πg.Object, 1)
							πTemp002[0] = πg.NewInt(9).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 545: self.assertEqual(a, self.type2test([1, 2, 3, 9]))
							πF.SetLineno(545)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = πF.MakeArgs(1)
							πTemp008 = make([]*πg.Object, 4)
							πTemp008[0] = πg.NewInt(1).ToObject()
							πTemp008[1] = πg.NewInt(2).ToObject()
							πTemp008[2] = πg.NewInt(3).ToObject()
							πTemp008[3] = πg.NewInt(9).ToObject()
							πTemp003 = πg.NewList(πTemp008...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_exhausted_iterator.ToObject(), πTemp017); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("CommonTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCommonTest.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("test.list_tests", Code)
}
