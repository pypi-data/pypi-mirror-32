package test_fpformat
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/test/test_fpformat.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ß0 := πg.InternStr("0")
		ß1 := πg.InternStr("1")
		ßFpformatTest := πg.InternStr("FpformatTest")
		ßNotANumber := πg.InternStr("NotANumber")
		ßStringType := πg.InternStr("StringType")
		ßTestCase := πg.InternStr("TestCase")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßassertEqual := πg.InternStr("assertEqual")
		ßcheckFix := πg.InternStr("checkFix")
		ßcheckSci := πg.InternStr("checkSci")
		ße := πg.InternStr("e")
		ßfail := πg.InternStr("fail")
		ßfix := πg.InternStr("fix")
		ßfloat := πg.InternStr("float")
		ßfpformat := πg.InternStr("fpformat")
		ßgrumpy := πg.InternStr("grumpy")
		ßisinstance := πg.InternStr("isinstance")
		ßlen := πg.InternStr("len")
		ßrange := πg.InternStr("range")
		ßrepr := πg.InternStr("repr")
		ßrun_unittest := πg.InternStr("run_unittest")
		ßsci := πg.InternStr("sci")
		ßskip := πg.InternStr("skip")
		ßsplit := πg.InternStr("split")
		ßtest_basic_cases := πg.InternStr("test_basic_cases")
		ßtest_failing_values := πg.InternStr("test_failing_values")
		ßtest_main := πg.InternStr("test_main")
		ßtest_reasonable_values := πg.InternStr("test_reasonable_values")
		ßtype := πg.InternStr("type")
		ßunittest := πg.InternStr("unittest")
		var πTemp001 *πg.Object
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
		var πTemp007 *πg.Dict
		_ = πTemp007
		var πTemp008 []πg.Param
		_ = πTemp008
		var πTemp009 bool
		_ = πTemp009
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: '''
			πF.SetLineno(1)
			// line 5: from test.test_support import run_unittest #, import_module
			πF.SetLineno(5)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßrun_unittest, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßrun_unittest.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 6: import unittest
			πF.SetLineno(6)
			if πTemp002, πE = πg.ImportModule(πF, "unittest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 8: import fpformat
			πF.SetLineno(8)
			if πTemp002, πE = πg.ImportModule(πF, "fpformat"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßfpformat.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 9: fix, sci, NotANumber = fpformat.fix, fpformat.sci, fpformat.NotANumber
			πF.SetLineno(9)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßfpformat); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßfix, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßfpformat); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßsci, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßfpformat); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßNotANumber, nil); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πTemp004, πTemp005, πTemp006).ToObject()
			if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp001); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßfix.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßsci.ToObject(), πTemp004); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNotANumber.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 11: StringType = type('')
			πF.SetLineno(11)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = ß.ToObject()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßStringType.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 22: class FpformatTest(unittest.TestCase):
			πF.SetLineno(22)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßTestCase, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp007 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("FpformatTest", "build/src/__python__/test/test_fpformat.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp007
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
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 24: def checkFix(self, n, digits):
					πF.SetLineno(24)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "n", Def: nil}
					πTemp002[2] = πg.Param{Name: "digits", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("checkFix", "build/src/__python__/test/test_fpformat.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µn *πg.Object = πArgs[1]; _ = µn
						var µdigits *πg.Object = πArgs[2]; _ = µdigits
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
						var µexpected *πg.Object = πg.UnboundLocal; _ = µexpected
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
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
							// line 25: result = fix(n, digits)
							πF.SetLineno(25)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[0] = µn
							if πE = πg.CheckLocal(πF, µdigits, "digits"); πE != nil {
								continue
							}
							πTemp001[1] = µdigits
							if πTemp002, πE = πg.ResolveGlobal(πF, ßfix); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µresult = πTemp003
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[0] = µn
							if πTemp002, πE = πg.ResolveGlobal(πF, ßStringType); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 26: if isinstance(n, StringType):
							πF.SetLineno(26)
						Label1:
							// line 27: n = repr(n)
							πF.SetLineno(27)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[0] = µn
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µn = πTemp003
							goto Label2
						Label2:
							// line 28: expected = "%.*f" % (digits, float(n))
							πF.SetLineno(28)
							if πE = πg.CheckLocal(πF, µdigits, "digits"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[0] = µn
							if πTemp005, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp003 = πg.NewTuple2(µdigits, πTemp006).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%.*f").ToObject(), πTemp003); πE != nil {
								continue
							}
							µexpected = πTemp002
							// line 30: self.assertEqual(result, expected)
							πF.SetLineno(30)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							if πE = πg.CheckLocal(πF, µexpected, "expected"); πE != nil {
								continue
							}
							πTemp001[1] = µexpected
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
					if πE = πClass.SetItem(πF, ßcheckFix.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 32: def checkSci(self, n, digits):
					πF.SetLineno(32)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "n", Def: nil}
					πTemp002[2] = πg.Param{Name: "digits", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("checkSci", "build/src/__python__/test/test_fpformat.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µn *πg.Object = πArgs[1]; _ = µn
						var µdigits *πg.Object = πArgs[2]; _ = µdigits
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
						var µexpected *πg.Object = πg.UnboundLocal; _ = µexpected
						var µnum *πg.Object = πg.UnboundLocal; _ = µnum
						var µexp *πg.Object = πg.UnboundLocal; _ = µexp
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
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
							// line 33: result = sci(n, digits)
							πF.SetLineno(33)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[0] = µn
							if πE = πg.CheckLocal(πF, µdigits, "digits"); πE != nil {
								continue
							}
							πTemp001[1] = µdigits
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsci); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µresult = πTemp003
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[0] = µn
							if πTemp002, πE = πg.ResolveGlobal(πF, ßStringType); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 34: if isinstance(n, StringType):
							πF.SetLineno(34)
						Label1:
							// line 35: n = repr(n)
							πF.SetLineno(35)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[0] = µn
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µn = πTemp003
							goto Label2
						Label2:
							// line 36: expected = "%.*e" % (digits, float(n))
							πF.SetLineno(36)
							if πE = πg.CheckLocal(πF, µdigits, "digits"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[0] = µn
							if πTemp005, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp003 = πg.NewTuple2(µdigits, πTemp006).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%.*e").ToObject(), πTemp003); πE != nil {
								continue
							}
							µexpected = πTemp002
							// line 38: num, exp = expected.split("e")
							πF.SetLineno(38)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ße.ToObject()
							if πE = πg.CheckLocal(πF, µexpected, "expected"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µexpected, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
								continue
							}
							µnum = πTemp002
							µexp = πTemp005
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µexp, "exp"); πE != nil {
								continue
							}
							πTemp001[0] = µexp
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.LT(πF, πTemp005, πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 39: if len(exp) < 4:
							πF.SetLineno(39)
						Label3:
							// line 40: exp = exp[0] + "0" + exp[1:]
							πF.SetLineno(40)
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µexp, "exp"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µexp, πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp006, ß0.ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µexp, "exp"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µexp, πTemp005); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp006); πE != nil {
								continue
							}
							µexp = πTemp002
							goto Label4
						Label4:
							// line 41: expected = "%se%s" % (num, exp)
							πF.SetLineno(41)
							if πE = πg.CheckLocal(πF, µnum, "num"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µexp, "exp"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(µnum, µexp).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%se%s").ToObject(), πTemp003); πE != nil {
								continue
							}
							µexpected = πTemp002
							// line 43: self.assertEqual(result, expected)
							πF.SetLineno(43)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							if πE = πg.CheckLocal(πF, µexpected, "expected"); πE != nil {
								continue
							}
							πTemp001[1] = µexpected
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
					if πE = πClass.SetItem(πF, ßcheckSci.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 45: def test_basic_cases(self):
					πF.SetLineno(45)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_basic_cases", "build/src/__python__/test/test_fpformat.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 46: self.assertEqual(fix(100.0/3, 3), '33.333')
							πF.SetLineno(46)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(2)
							if πTemp003, πE = πg.Div(πF, πg.NewFloat(100.0).ToObject(), πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewInt(3).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfix); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewStr("33.333").ToObject()
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
							// line 47: self.assertEqual(sci(100.0/3, 3), '3.333e+001')
							πF.SetLineno(47)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(2)
							if πTemp003, πE = πg.Div(πF, πg.NewFloat(100.0).ToObject(), πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewInt(3).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsci); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewStr("3.333e+001").ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_basic_cases.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 50: def test_reasonable_values(self):
					πF.SetLineno(50)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_reasonable_values", "build/src/__python__/test/test_fpformat.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µval *πg.Object = πg.UnboundLocal; _ = µval
						var µrealVal *πg.Object = πg.UnboundLocal; _ = µrealVal
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 *πg.Object
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
							case 4: goto Label4
							case 5: goto Label5
							case 7: goto Label7
							case 8: goto Label8
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(7).ToObject()
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
								µd = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πTemp007, πE = πg.Div(πF, πg.NewFloat(1000.0).ToObject(), πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Div(πF, πg.NewFloat(1.0).ToObject(), πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple6(πTemp007, πg.NewInt(1000).ToObject(), πg.NewFloat(1000.0).ToObject(), πg.NewFloat(0.002).ToObject(), πTemp008, πg.NewFloat(10000000000.0).ToObject()).ToObject()
							if πTemp003, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp006 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label6
							}
							if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp009 = !isStop
							} else {
								πTemp009 = true
								µval = πTemp004
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(4)            
							if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Div(πF, πg.NewFloat(1.0).ToObject(), µval); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Neg(πF, µval); πE != nil {
								continue
							}
							if πTemp012, πE = πg.Neg(πF, πg.NewFloat(1.0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Div(πF, πTemp012, µval); πE != nil {
								continue
							}
							πTemp007 = πg.NewTuple4(µval, πTemp008, πTemp010, πTemp011).ToObject()
							if πTemp004, πE = πg.Iter(πF, πTemp007); πE != nil {
								continue
							}
							πF.PushCheckpoint(8)
							πTemp009 = false
						Label7:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp009 {
								πF.PopCheckpoint()
								goto Label9
							}
							if πTemp007, πE = πg.Next(πF, πTemp004); πE != nil {
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
								µrealVal = πTemp007
							}
							if πE != nil || !πTemp013 {
								continue
							}
							πF.PushCheckpoint(7)            
							// line 54: self.checkFix(realVal, d)
							πF.SetLineno(54)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µrealVal, "realVal"); πE != nil {
								continue
							}
							πTemp002[0] = µrealVal
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp002[1] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßcheckFix, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 55: self.checkSci(realVal, d)
							πF.SetLineno(55)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µrealVal, "realVal"); πE != nil {
								continue
							}
							πTemp002[0] = µrealVal
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp002[1] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßcheckSci, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label8:
							if πE != nil || πR != nil {
								continue
							}
						Label9:
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
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_reasonable_values.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 49: @unittest.skip('grumpy')
					πF.SetLineno(49)
					πTemp006 = πF.MakeArgs(1)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßtest_reasonable_values); πE != nil {
						continue
					}
					πTemp006[0] = πTemp007
					πTemp008 = πF.MakeArgs(1)
					πTemp008[0] = ßgrumpy.ToObject()
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßskip, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp009.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πTemp009, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßtest_reasonable_values.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 57: def test_failing_values(self):
					πF.SetLineno(57)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_failing_values", "build/src/__python__/test/test_fpformat.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µyacht *πg.Object = πg.UnboundLocal; _ = µyacht
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 59: self.assertEqual(fix(1.0, 1000), '1.'+('0'*1000))
							πF.SetLineno(59)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewFloat(1.0).ToObject()
							πTemp002[1] = πg.NewInt(1000).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfix); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp004, πE = πg.Mul(πF, ß0.ToObject(), πg.NewInt(1000).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πg.NewStr("1.").ToObject(), πTemp004); πE != nil {
								continue
							}
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
							// line 60: self.assertEqual(sci("1"+('0'*1000), 0), '1e+1000')
							πF.SetLineno(60)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(2)
							if πTemp004, πE = πg.Mul(πF, ß0.ToObject(), πg.NewInt(1000).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, ß1.ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewInt(0).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsci); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewStr("1e+1000").ToObject()
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
							// line 63: yacht = "Throatwobbler Mangrove"
							πF.SetLineno(63)
							µyacht = πg.NewStr("Throatwobbler Mangrove").ToObject()
							// line 64: self.assertEqual(fix(yacht, 10), yacht)
							πF.SetLineno(64)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µyacht, "yacht"); πE != nil {
								continue
							}
							πTemp002[0] = µyacht
							πTemp002[1] = πg.NewInt(10).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfix); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µyacht, "yacht"); πE != nil {
								continue
							}
							πTemp001[1] = µyacht
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
							// line 65: try:
							πF.SetLineno(65)
							πF.PushCheckpoint(2)
							// line 66: sci(yacht, 10)
							πF.SetLineno(66)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µyacht, "yacht"); πE != nil {
								continue
							}
							πTemp001[0] = µyacht
							πTemp001[1] = πg.NewInt(10).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsci); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πF.PopCheckpoint()
							// line 70: self.fail("No exception on non-numeric sci")
							πF.SetLineno(70)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("No exception on non-numeric sci").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNotANumber); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label3
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 67: except NotANumber:
							πF.SetLineno(67)
						Label3:
							// line 68: pass
							πF.SetLineno(68)
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
					if πE = πClass.SetItem(πF, ßtest_failing_values.ToObject(), πTemp007); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("FpformatTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFpformatTest.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 73: def test_main():
			πF.SetLineno(73)
			πTemp008 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("test_main", "build/src/__python__/test/test_fpformat.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 74: run_unittest(FpformatTest)
					πF.SetLineno(74)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFpformatTest); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßrun_unittest); πE != nil {
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
			if πTemp004, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp003, πE = πg.Eq(πF, πTemp004, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
				continue
			}
			if πTemp009 {
				goto Label1
			}
			goto Label2
			// line 77: if __name__ == "__main__":
			πF.SetLineno(77)
		Label1:
			// line 78: test_main()
			πF.SetLineno(78)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_main); πE != nil {
				continue
			}
			if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("test.test_fpformat", Code)
}
