package test_tuple
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/test/test_tuple.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßCommonTest := πg.InternStr("CommonTest")
		ßTupleTest := πg.InternStr("TupleTest")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßa := πg.InternStr("a")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertTrue := πg.InternStr("assertTrue")
		ßhash := πg.InternStr("hash")
		ßlen := πg.InternStr("len")
		ßlist := πg.InternStr("list")
		ßm := πg.InternStr("m")
		ßmap := πg.InternStr("map")
		ßp := πg.InternStr("p")
		ßrange := πg.InternStr("range")
		ßrepr := πg.InternStr("repr")
		ßrun_unittest := πg.InternStr("run_unittest")
		ßs := πg.InternStr("s")
		ßseq_tests := πg.InternStr("seq_tests")
		ßset := πg.InternStr("set")
		ßspam := πg.InternStr("spam")
		ßstr := πg.InternStr("str")
		ßsuper := πg.InternStr("super")
		ßtest_constructors := πg.InternStr("test_constructors")
		ßtest_hash := πg.InternStr("test_hash")
		ßtest_iadd := πg.InternStr("test_iadd")
		ßtest_imul := πg.InternStr("test_imul")
		ßtest_len := πg.InternStr("test_len")
		ßtest_main := πg.InternStr("test_main")
		ßtest_repr := πg.InternStr("test_repr")
		ßtest_support := πg.InternStr("test_support")
		ßtest_truth := πg.InternStr("test_truth")
		ßtest_tupleresizebug := πg.InternStr("test_tupleresizebug")
		ßtuple := πg.InternStr("tuple")
		ßtype2test := πg.InternStr("type2test")
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
		var πTemp007 []πg.Param
		_ = πTemp007
		var πTemp008 bool
		_ = πTemp008
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: from test import test_support, seq_tests
			πF.SetLineno(1)
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
			// line 5: class TupleTest(seq_tests.CommonTest):
			πF.SetLineno(5)
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
			_, πE = πg.NewCode("TupleTest", "build/src/__python__/test/test_tuple.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 6: type2test = tuple
					πF.SetLineno(6)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßtuple); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßtype2test.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 8: def test_constructors(self):
					πF.SetLineno(8)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_constructors", "build/src/__python__/test/test_tuple.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µt0_3 *πg.Object = πg.UnboundLocal; _ = µt0_3
						var µt0_3_bis *πg.Object = πg.UnboundLocal; _ = µt0_3_bis
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
							// line 9: super(TupleTest, self).test_constructors()
							πF.SetLineno(9)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTupleTest); πE != nil {
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
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßtest_constructors, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 11: self.assertEqual(tuple(), ())
							πF.SetLineno(11)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp002 = πg.NewTuple0().ToObject()
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
							// line 12: t0_3 = (0, 1, 2, 3)
							πF.SetLineno(12)
							πTemp002 = πg.NewTuple4(πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject(), πg.NewInt(3).ToObject()).ToObject()
							µt0_3 = πTemp002
							// line 13: t0_3_bis = tuple(t0_3)
							πF.SetLineno(13)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µt0_3, "t0_3"); πE != nil {
								continue
							}
							πTemp001[0] = µt0_3
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µt0_3_bis = πTemp003
							// line 14: self.assertTrue(t0_3 is t0_3_bis)
							πF.SetLineno(14)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µt0_3, "t0_3"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µt0_3_bis, "t0_3_bis"); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µt0_3 == µt0_3_bis).ToObject()
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
							// line 15: self.assertEqual(tuple([]), ())
							πF.SetLineno(15)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							πTemp002 = πg.NewTuple0().ToObject()
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
							// line 16: self.assertEqual(tuple([0, 1, 2, 3]), (0, 1, 2, 3))
							πF.SetLineno(16)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 4)
							πTemp005[0] = πg.NewInt(0).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							πTemp002 = πg.NewTuple4(πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject(), πg.NewInt(3).ToObject()).ToObject()
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
							// line 17: self.assertEqual(tuple(''), ())
							πF.SetLineno(17)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ß.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							πTemp002 = πg.NewTuple0().ToObject()
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
							// line 18: self.assertEqual(tuple('spam'), ('s', 'p', 'a', 'm'))
							πF.SetLineno(18)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßspam.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							πTemp002 = πg.NewTuple4(ßs.ToObject(), ßp.ToObject(), ßa.ToObject(), ßm.ToObject()).ToObject()
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_constructors.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 20: def test_truth(self):
					πF.SetLineno(20)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_truth", "build/src/__python__/test/test_tuple.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 21: super(TupleTest, self).test_truth()
							πF.SetLineno(21)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTupleTest); πE != nil {
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
							// line 22: self.assertTrue(not ())
							πF.SetLineno(22)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πg.NewTuple0().ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp004).ToObject()
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
							// line 23: self.assertTrue((42, ))
							πF.SetLineno(23)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewTuple1(πg.NewInt(42).ToObject()).ToObject()
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
					// line 25: def test_len(self):
					πF.SetLineno(25)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_len", "build/src/__python__/test/test_tuple.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 26: super(TupleTest, self).test_len()
							πF.SetLineno(26)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTupleTest); πE != nil {
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
							// line 27: self.assertEqual(len(()), 0)
							πF.SetLineno(27)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp002 = πg.NewTuple0().ToObject()
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
							// line 28: self.assertEqual(len((0,)), 1)
							πF.SetLineno(28)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp002 = πg.NewTuple1(πg.NewInt(0).ToObject()).ToObject()
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
							// line 29: self.assertEqual(len((0, 1, 2)), 3)
							πF.SetLineno(29)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp002 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject()).ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_len.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 31: def test_iadd(self):
					πF.SetLineno(31)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_iadd", "build/src/__python__/test/test_tuple.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µu *πg.Object = πg.UnboundLocal; _ = µu
						var µu2 *πg.Object = πg.UnboundLocal; _ = µu2
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
							// line 32: super(TupleTest, self).test_iadd()
							πF.SetLineno(32)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTupleTest); πE != nil {
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
							// line 33: u = (0, 1)
							πF.SetLineno(33)
							πTemp002 = πg.NewTuple2(πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject()).ToObject()
							µu = πTemp002
							// line 34: u2 = u
							πF.SetLineno(34)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							µu2 = µu
							// line 35: u += (2, 3)
							πF.SetLineno(35)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πg.NewInt(2).ToObject(), πg.NewInt(3).ToObject()).ToObject()
							if πTemp003, πE = πg.IAdd(πF, µu, πTemp002); πE != nil {
								continue
							}
							µu = πTemp003
							// line 36: self.assertTrue(u is not u2)
							πF.SetLineno(36)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µu != µu2).ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_iadd.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 38: def test_imul(self):
					πF.SetLineno(38)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_imul", "build/src/__python__/test/test_tuple.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µu *πg.Object = πg.UnboundLocal; _ = µu
						var µu2 *πg.Object = πg.UnboundLocal; _ = µu2
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
							// line 39: super(TupleTest, self).test_imul()
							πF.SetLineno(39)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTupleTest); πE != nil {
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
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßtest_imul, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 40: u = (0, 1)
							πF.SetLineno(40)
							πTemp002 = πg.NewTuple2(πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject()).ToObject()
							µu = πTemp002
							// line 41: u2 = u
							πF.SetLineno(41)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							µu2 = µu
							// line 42: u *= 3
							πF.SetLineno(42)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IMul(πF, µu, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							µu = πTemp002
							// line 43: self.assertTrue(u is not u2)
							πF.SetLineno(43)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µu != µu2).ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_imul.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 45: def test_tupleresizebug(self):
					πF.SetLineno(45)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_tupleresizebug", "build/src/__python__/test/test_tuple.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []πg.Param
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							// line 47: def f():
							πF.SetLineno(47)
							πTemp002 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("f", "build/src/__python__/test/test_tuple.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µi *πg.Object = πg.UnboundLocal; _ = µi
								var πTemp001 *πg.Object
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
										πTemp002 = πF.MakeArgs(1)
										πTemp002[0] = πg.NewInt(1000).ToObject()
										if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
											continue
										}
										if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp002)
										if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
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
										if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
											µi = πTemp003
										}
										if πE != nil || !πTemp006 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 49: yield i
										πF.SetLineno(49)
										if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return µi, nil
									Label4:
										πTemp003 = πSent
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
							µf = πTemp001
							// line 50: self.assertEqual(list(tuple(f())), range(1000))
							πF.SetLineno(50)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp006, πE = µf.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp006
							if πTemp006, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[0] = πTemp007
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp007
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewInt(1000).ToObject()
							if πTemp006, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[1] = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_tupleresizebug.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 52: def test_hash(self):
					πF.SetLineno(52)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("test_hash", "build/src/__python__/test/test_tuple.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µN *πg.Object = πg.UnboundLocal; _ = µN
						var µbase *πg.Object = πg.UnboundLocal; _ = µbase
						var µxp *πg.Object = πg.UnboundLocal; _ = µxp
						var µinps *πg.Object = πg.UnboundLocal; _ = µinps
						var µcollisions *πg.Object = πg.UnboundLocal; _ = µcollisions
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []πg.Param
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
						var πTemp012 []*πg.Object
						_ = πTemp012
						var πTemp013 []*πg.Object
						_ = πTemp013
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 68: N=50
							πF.SetLineno(68)
							µN = πg.NewInt(50).ToObject()
							// line 69: base = range(N)
							πF.SetLineno(69)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
								continue
							}
							πTemp001[0] = µN
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µbase = πTemp003
							// line 70: xp = [(i, j) for i in base for j in base]
							πF.SetLineno(70)
							πTemp004 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/test/test_tuple.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µi *πg.Object = πg.UnboundLocal; _ = µi
								var µj *πg.Object = πg.UnboundLocal; _ = µj
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 bool
								_ = πTemp002
								var πTemp003 bool
								_ = πTemp003
								var πTemp004 *πg.Object
								_ = πTemp004
								var πTemp005 bool
								_ = πTemp005
								var πTemp006 *πg.Object
								_ = πTemp006
								var πTemp007 *πg.Object
								_ = πTemp007
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
									for ; πF.State() >= 0; πF.PopCheckpoint() {
										switch πF.State() {
										case 0:
										case 1: goto Label1
										case 2: goto Label2
										case 4: goto Label4
										case 5: goto Label5
										case 7: goto Label7
										default: panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µbase); πE != nil {
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
											µi = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)            
										if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
											continue
										}
										if πTemp004, πE = πg.Iter(πF, µbase); πE != nil {
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
										if πTemp006, πE = πg.Next(πF, πTemp004); πE != nil {
											isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
											if exc != nil {
												πE = exc
											} else if isStop {
												πE = nil
												πF.RestoreExc(nil, nil)
											}
											πTemp005 = !isStop
										} else {
											πTemp005 = true
											µj = πTemp006
										}
										if πE != nil || !πTemp005 {
											continue
										}
										πF.PushCheckpoint(4)            
										// line 70: xp = [(i, j) for i in base for j in base]
										πF.SetLineno(70)
										if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
											continue
										}
										πTemp006 = πg.NewTuple2(µi, µj).ToObject()
										πF.PushCheckpoint(7)
										return πTemp006, nil
									Label7:
										πTemp007 = πSent
										continue
									Label5:
										if πE != nil || πR != nil {
											continue
										}
									Label6:
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
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
								continue
							}
							µxp = πTemp002
							// line 71: inps = base + [(i, j) for i in base for j in xp] + \
							πF.SetLineno(71)
							if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
								continue
							}
							πTemp004 = make([]πg.Param, 0)
							πTemp009 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/test/test_tuple.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µi *πg.Object = πg.UnboundLocal; _ = µi
								var µj *πg.Object = πg.UnboundLocal; _ = µj
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 bool
								_ = πTemp002
								var πTemp003 bool
								_ = πTemp003
								var πTemp004 *πg.Object
								_ = πTemp004
								var πTemp005 bool
								_ = πTemp005
								var πTemp006 *πg.Object
								_ = πTemp006
								var πTemp007 *πg.Object
								_ = πTemp007
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
									for ; πF.State() >= 0; πF.PopCheckpoint() {
										switch πF.State() {
										case 0:
										case 1: goto Label1
										case 2: goto Label2
										case 4: goto Label4
										case 5: goto Label5
										case 7: goto Label7
										default: panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µbase); πE != nil {
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
											µi = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)            
										if πE = πg.CheckLocal(πF, µxp, "xp"); πE != nil {
											continue
										}
										if πTemp004, πE = πg.Iter(πF, µxp); πE != nil {
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
										if πTemp006, πE = πg.Next(πF, πTemp004); πE != nil {
											isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
											if exc != nil {
												πE = exc
											} else if isStop {
												πE = nil
												πF.RestoreExc(nil, nil)
											}
											πTemp005 = !isStop
										} else {
											πTemp005 = true
											µj = πTemp006
										}
										if πE != nil || !πTemp005 {
											continue
										}
										πF.PushCheckpoint(4)            
										// line 71: inps = base + [(i, j) for i in base for j in xp] + \
										πF.SetLineno(71)
										if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
											continue
										}
										πTemp006 = πg.NewTuple2(µi, µj).ToObject()
										πF.PushCheckpoint(7)
										return πTemp006, nil
									Label7:
										πTemp007 = πSent
										continue
									Label5:
										if πE != nil || πR != nil {
											continue
										}
									Label6:
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
							if πTemp010, πE = πTemp009.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.ListType.Call(πF, πg.Args{πTemp010}, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Add(πF, µbase, πTemp008); πE != nil {
								continue
							}
							πTemp004 = make([]πg.Param, 0)
							πTemp010 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/test/test_tuple.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µi *πg.Object = πg.UnboundLocal; _ = µi
								var µj *πg.Object = πg.UnboundLocal; _ = µj
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 bool
								_ = πTemp002
								var πTemp003 bool
								_ = πTemp003
								var πTemp004 *πg.Object
								_ = πTemp004
								var πTemp005 bool
								_ = πTemp005
								var πTemp006 *πg.Object
								_ = πTemp006
								var πTemp007 *πg.Object
								_ = πTemp007
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
									for ; πF.State() >= 0; πF.PopCheckpoint() {
										switch πF.State() {
										case 0:
										case 1: goto Label1
										case 2: goto Label2
										case 4: goto Label4
										case 5: goto Label5
										case 7: goto Label7
										default: panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µxp, "xp"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µxp); πE != nil {
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
											µi = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)            
										if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
											continue
										}
										if πTemp004, πE = πg.Iter(πF, µbase); πE != nil {
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
										if πTemp006, πE = πg.Next(πF, πTemp004); πE != nil {
											isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
											if exc != nil {
												πE = exc
											} else if isStop {
												πE = nil
												πF.RestoreExc(nil, nil)
											}
											πTemp005 = !isStop
										} else {
											πTemp005 = true
											µj = πTemp006
										}
										if πE != nil || !πTemp005 {
											continue
										}
										πF.PushCheckpoint(4)            
										// line 72: [(i, j) for i in xp for j in base] + xp + zip(base)
										πF.SetLineno(72)
										if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
											continue
										}
										πTemp006 = πg.NewTuple2(µi, µj).ToObject()
										πF.PushCheckpoint(7)
										return πTemp006, nil
									Label7:
										πTemp007 = πSent
										continue
									Label5:
										if πE != nil || πR != nil {
											continue
										}
									Label6:
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
							if πTemp011, πE = πTemp010.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.ListType.Call(πF, πg.Args{πTemp011}, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Add(πF, πTemp007, πTemp008); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µxp, "xp"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Add(πF, πTemp006, µxp); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
								continue
							}
							πTemp001[0] = µbase
							if πTemp006, πE = πg.ResolveGlobal(πF, ßzip); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Add(πF, πTemp005, πTemp007); πE != nil {
								continue
							}
							µinps = πTemp002
							// line 73: collisions = len(inps) - len(set(map(hash, inps)))
							πF.SetLineno(73)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µinps, "inps"); πE != nil {
								continue
							}
							πTemp001[0] = µinps
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							πTemp012 = πF.MakeArgs(1)
							πTemp013 = πF.MakeArgs(2)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßhash); πE != nil {
								continue
							}
							πTemp013[0] = πTemp005
							if πE = πg.CheckLocal(πF, µinps, "inps"); πE != nil {
								continue
							}
							πTemp013[1] = µinps
							if πTemp005, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp013, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp013)
							πTemp012[0] = πTemp007
							if πTemp005, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp012, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp012)
							πTemp001[0] = πTemp007
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Sub(πF, πTemp006, πTemp007); πE != nil {
								continue
							}
							µcollisions = πTemp002
							// line 74: self.assertTrue(collisions <= 15)
							πF.SetLineno(74)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcollisions, "collisions"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LE(πF, µcollisions, πg.NewInt(15).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_hash.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 76: def test_repr(self):
					πF.SetLineno(76)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("test_repr", "build/src/__python__/test/test_tuple.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µl0 *πg.Object = πg.UnboundLocal; _ = µl0
						var µl2 *πg.Object = πg.UnboundLocal; _ = µl2
						var µa0 *πg.Object = πg.UnboundLocal; _ = µa0
						var µa2 *πg.Object = πg.UnboundLocal; _ = µa2
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 77: l0 = tuple()
							πF.SetLineno(77)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µl0 = πTemp002
							// line 78: l2 = (0, 1, 2)
							πF.SetLineno(78)
							πTemp001 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject()).ToObject()
							µl2 = πTemp001
							// line 79: a0 = self.type2test(l0)
							πF.SetLineno(79)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µl0, "l0"); πE != nil {
								continue
							}
							πTemp003[0] = µl0
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µa0 = πTemp002
							// line 80: a2 = self.type2test(l2)
							πF.SetLineno(80)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µl2, "l2"); πE != nil {
								continue
							}
							πTemp003[0] = µl2
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µa2 = πTemp002
							// line 82: self.assertEqual(str(a0), repr(l0))
							πF.SetLineno(82)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa0, "a0"); πE != nil {
								continue
							}
							πTemp004[0] = µa0
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µl0, "l0"); πE != nil {
								continue
							}
							πTemp004[0] = µl0
							if πTemp001, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[1] = πTemp002
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
							// line 83: self.assertEqual(str(a2), repr(l2))
							πF.SetLineno(83)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa2, "a2"); πE != nil {
								continue
							}
							πTemp004[0] = µa2
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µl2, "l2"); πE != nil {
								continue
							}
							πTemp004[0] = µl2
							if πTemp001, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[1] = πTemp002
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
							// line 84: self.assertEqual(repr(a0), "()")
							πF.SetLineno(84)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa0, "a0"); πE != nil {
								continue
							}
							πTemp004[0] = µa0
							if πTemp001, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewStr("()").ToObject()
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
							// line 85: self.assertEqual(repr(a2), "(0, 1, 2)")
							πF.SetLineno(85)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa2, "a2"); πE != nil {
								continue
							}
							πTemp004[0] = µa2
							if πTemp001, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewStr("(0, 1, 2)").ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_repr.ToObject(), πTemp009); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("TupleTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTupleTest.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 157: def test_main():
			πF.SetLineno(157)
			πTemp007 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("test_main", "build/src/__python__/test/test_tuple.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 158: test_support.run_unittest(TupleTest)
					πF.SetLineno(158)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTupleTest); πE != nil {
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
			// line 160: if __name__=="__main__":
			πF.SetLineno(160)
		Label1:
			// line 161: test_main()
			πF.SetLineno(161)
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
	πg.RegisterModule("test.test_tuple", Code)
}
