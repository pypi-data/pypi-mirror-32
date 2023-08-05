package test_uu
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/test/test_uu.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ß3 := πg.InternStr("3")
		ßError := πg.InternStr("Error")
		ßKeyboardInterrupt := πg.InternStr("KeyboardInterrupt")
		ßNone := πg.InternStr("None")
		ßStringIO := πg.InternStr("StringIO")
		ßSystemExit := πg.InternStr("SystemExit")
		ßTESTFN := πg.InternStr("TESTFN")
		ßTestCase := πg.InternStr("TestCase")
		ßTrue := πg.InternStr("True")
		ßUUFileTest := πg.InternStr("UUFileTest")
		ßUUStdIOTest := πg.InternStr("UUStdIOTest")
		ßUUTest := πg.InternStr("UUTest")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_kill := πg.InternStr("_kill")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertRaises := πg.InternStr("assertRaises")
		ßcStringIO := πg.InternStr("cStringIO")
		ßclose := πg.InternStr("close")
		ßdecode := πg.InternStr("decode")
		ßencode := πg.InternStr("encode")
		ßencodedtext := πg.InternStr("encodedtext")
		ßencodedtextwrapped := πg.InternStr("encodedtextwrapped")
		ßfail := πg.InternStr("fail")
		ßgetvalue := πg.InternStr("getvalue")
		ßi := πg.InternStr("i")
		ßname := πg.InternStr("name")
		ßo := πg.InternStr("o")
		ßopen := πg.InternStr("open")
		ßos := πg.InternStr("os")
		ßplaintext := πg.InternStr("plaintext")
		ßr := πg.InternStr("r")
		ßrb := πg.InternStr("rb")
		ßread := πg.InternStr("read")
		ßreplace := πg.InternStr("replace")
		ßrun_unittest := πg.InternStr("run_unittest")
		ßsetUp := πg.InternStr("setUp")
		ßstdin := πg.InternStr("stdin")
		ßstdout := πg.InternStr("stdout")
		ßstr := πg.InternStr("str")
		ßsys := πg.InternStr("sys")
		ßt1 := πg.InternStr("t1")
		ßtearDown := πg.InternStr("tearDown")
		ßtest_decode := πg.InternStr("test_decode")
		ßtest_decode_filename := πg.InternStr("test_decode_filename")
		ßtest_decodetwice := πg.InternStr("test_decodetwice")
		ßtest_encode := πg.InternStr("test_encode")
		ßtest_garbage_padding := πg.InternStr("test_garbage_padding")
		ßtest_main := πg.InternStr("test_main")
		ßtest_missingbegin := πg.InternStr("test_missingbegin")
		ßtest_support := πg.InternStr("test_support")
		ßtest_truncatedinput := πg.InternStr("test_truncatedinput")
		ßtmpin := πg.InternStr("tmpin")
		ßtmpout := πg.InternStr("tmpout")
		ßunittest := πg.InternStr("unittest")
		ßunlink := πg.InternStr("unlink")
		ßuu := πg.InternStr("uu")
		ßw := πg.InternStr("w")
		ßwb := πg.InternStr("wb")
		ßwrite := πg.InternStr("write")
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
		var πTemp006 *πg.Dict
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
			// line 1: """
			πF.SetLineno(1)
			// line 6: import unittest
			πF.SetLineno(6)
			if πTemp002, πE = πg.ImportModule(πF, "unittest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunittest.ToObject(), πTemp001); πE != nil {
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
			// line 9: import sys, os, uu, cStringIO
			πF.SetLineno(9)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "uu"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßuu.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "cStringIO"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßcStringIO.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 10: import uu
			πF.SetLineno(10)
			if πTemp002, πE = πg.ImportModule(πF, "uu"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßuu.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 12: plaintext = "The smooth-scaled python crept over the sleeping dog\n"
			πF.SetLineno(12)
			if πE = πF.Globals().SetItem(πF, ßplaintext.ToObject(), πg.NewStr("The smooth-scaled python crept over the sleeping dog\n").ToObject()); πE != nil {
				continue
			}
			// line 14: encodedtext = """\
			πF.SetLineno(14)
			if πE = πF.Globals().SetItem(πF, ßencodedtext.ToObject(), πg.NewStr("M5&AE('-M;V]T:\"US8V%L960@<'ET:&]N(&-R97!T(&]V97(@=&AE('-L965P\n(:6YG(&1O9PH ").ToObject()); πE != nil {
				continue
			}
			// line 18: encodedtextwrapped = "begin %03o %s\n" + encodedtext.replace("%", "%%") + "\n \nend\n"
			πF.SetLineno(18)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("%").ToObject()
			πTemp002[1] = πg.NewStr("%%").ToObject()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßencodedtext); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßreplace, nil); πE != nil {
				continue
			}
			if πTemp004, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πTemp003, πE = πg.Add(πF, πg.NewStr("begin %03o %s\n").ToObject(), πTemp004); πE != nil {
				continue
			}
			if πTemp001, πE = πg.Add(πF, πTemp003, πg.NewStr("\n \nend\n").ToObject()); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßencodedtextwrapped.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 20: class UUTest(unittest.TestCase):
			πF.SetLineno(20)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßTestCase, nil); πE != nil {
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
			_, πE = πg.NewCode("UUTest", "build/src/__python__/test/test_uu.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp006 *πg.Object
				_ = πTemp006
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 22: def test_encode(self):
					πF.SetLineno(22)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_encode", "build/src/__python__/test/test_uu.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µinp *πg.Object = πg.UnboundLocal; _ = µinp
						var µout *πg.Object = πg.UnboundLocal; _ = µout
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
							// line 23: inp = cStringIO.StringIO(plaintext)
							πF.SetLineno(23)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßplaintext); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µinp = πTemp002
							// line 24: out = cStringIO.StringIO()
							πF.SetLineno(24)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µout = πTemp002
							// line 25: uu.encode(inp, out, "t1")
							πF.SetLineno(25)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
								continue
							}
							πTemp001[0] = µinp
							if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
								continue
							}
							πTemp001[1] = µout
							πTemp001[2] = ßt1.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßuu); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßencode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 26: self.assertEqual(out.getvalue(), encodedtextwrapped % (0666, "t1"))
							πF.SetLineno(26)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µout, ßgetvalue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßencodedtextwrapped); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(πg.NewInt(438).ToObject(), ßt1.ToObject()).ToObject()
							if πTemp002, πE = πg.Mod(πF, πTemp003, πTemp004); πE != nil {
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
							// line 27: inp = cStringIO.StringIO(plaintext)
							πF.SetLineno(27)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßplaintext); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µinp = πTemp002
							// line 28: out = cStringIO.StringIO()
							πF.SetLineno(28)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µout = πTemp002
							// line 29: uu.encode(inp, out, "t1", 0644)
							πF.SetLineno(29)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
								continue
							}
							πTemp001[0] = µinp
							if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
								continue
							}
							πTemp001[1] = µout
							πTemp001[2] = ßt1.ToObject()
							πTemp001[3] = πg.NewInt(420).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßuu); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßencode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 30: self.assertEqual(out.getvalue(), encodedtextwrapped % (0644, "t1"))
							πF.SetLineno(30)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µout, ßgetvalue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßencodedtextwrapped); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(πg.NewInt(420).ToObject(), ßt1.ToObject()).ToObject()
							if πTemp002, πE = πg.Mod(πF, πTemp003, πTemp004); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_encode.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 32: def test_decode(self):
					πF.SetLineno(32)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_decode", "build/src/__python__/test/test_uu.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µinp *πg.Object = πg.UnboundLocal; _ = µinp
						var µout *πg.Object = πg.UnboundLocal; _ = µout
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 33: inp = cStringIO.StringIO(encodedtextwrapped % (0666, "t1"))
							πF.SetLineno(33)
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßencodedtextwrapped); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(πg.NewInt(438).ToObject(), ßt1.ToObject()).ToObject()
							if πTemp002, πE = πg.Mod(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µinp = πTemp002
							// line 34: out = cStringIO.StringIO()
							πF.SetLineno(34)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µout = πTemp002
							// line 35: uu.decode(inp, out)
							πF.SetLineno(35)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
								continue
							}
							πTemp001[0] = µinp
							if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
								continue
							}
							πTemp001[1] = µout
							if πTemp002, πE = πg.ResolveGlobal(πF, ßuu); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdecode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 36: self.assertEqual(out.getvalue(), plaintext)
							πF.SetLineno(36)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µout, ßgetvalue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßplaintext); πE != nil {
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
							// line 37: inp = cStringIO.StringIO(
							πF.SetLineno(37)
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.Add(πF, πg.NewStr("UUencoded files may contain many lines,\n").ToObject(), πg.NewStr("even some that have 'begin' in them.\n").ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßencodedtextwrapped); πE != nil {
								continue
							}
							πTemp006 = πg.NewTuple2(πg.NewInt(438).ToObject(), ßt1.ToObject()).ToObject()
							if πTemp004, πE = πg.Mod(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µinp = πTemp002
							// line 42: out = cStringIO.StringIO()
							πF.SetLineno(42)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µout = πTemp002
							// line 43: uu.decode(inp, out)
							πF.SetLineno(43)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
								continue
							}
							πTemp001[0] = µinp
							if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
								continue
							}
							πTemp001[1] = µout
							if πTemp002, πE = πg.ResolveGlobal(πF, ßuu); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdecode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 44: self.assertEqual(out.getvalue(), plaintext)
							πF.SetLineno(44)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µout, ßgetvalue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßplaintext); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_decode.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 46: def test_truncatedinput(self):
					πF.SetLineno(46)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_truncatedinput", "build/src/__python__/test/test_uu.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µinp *πg.Object = πg.UnboundLocal; _ = µinp
						var µout *πg.Object = πg.UnboundLocal; _ = µout
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 47: inp = cStringIO.StringIO("begin 644 t1\n" + encodedtext)
							πF.SetLineno(47)
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßencodedtext); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πg.NewStr("begin 644 t1\n").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µinp = πTemp002
							// line 48: out = cStringIO.StringIO()
							πF.SetLineno(48)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µout = πTemp002
							// line 49: try:
							πF.SetLineno(49)
							πF.PushCheckpoint(2)
							// line 50: uu.decode(inp, out)
							πF.SetLineno(50)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
								continue
							}
							πTemp001[0] = µinp
							if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
								continue
							}
							πTemp001[1] = µout
							if πTemp002, πE = πg.ResolveGlobal(πF, ßuu); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdecode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 51: self.fail("No exception raised")
							πF.SetLineno(51)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("No exception raised").ToObject()
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
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßuu); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßError, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 52: except uu.Error, e:
							πF.SetLineno(52)
						Label3:
							µe = πTemp004.ToObject()
							// line 53: self.assertEqual(str(e), "Truncated input file")
							πF.SetLineno(53)
							πTemp001 = πF.MakeArgs(2)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							πTemp007[0] = µe
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr("Truncated input file").ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_truncatedinput.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 55: def test_missingbegin(self):
					πF.SetLineno(55)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_missingbegin", "build/src/__python__/test/test_uu.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µinp *πg.Object = πg.UnboundLocal; _ = µinp
						var µout *πg.Object = πg.UnboundLocal; _ = µout
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 56: inp = cStringIO.StringIO("")
							πF.SetLineno(56)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ß.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µinp = πTemp002
							// line 57: out = cStringIO.StringIO()
							πF.SetLineno(57)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µout = πTemp002
							// line 58: try:
							πF.SetLineno(58)
							πF.PushCheckpoint(2)
							// line 59: uu.decode(inp, out)
							πF.SetLineno(59)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
								continue
							}
							πTemp001[0] = µinp
							if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
								continue
							}
							πTemp001[1] = µout
							if πTemp002, πE = πg.ResolveGlobal(πF, ßuu); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdecode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 60: self.fail("No exception raised")
							πF.SetLineno(60)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("No exception raised").ToObject()
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
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßuu); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßError, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 61: except uu.Error, e:
							πF.SetLineno(61)
						Label3:
							µe = πTemp004.ToObject()
							// line 62: self.assertEqual(str(e), "No valid begin line found in input file")
							πF.SetLineno(62)
							πTemp001 = πF.MakeArgs(2)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							πTemp007[0] = µe
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr("No valid begin line found in input file").ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_missingbegin.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 64: def test_garbage_padding(self):
					πF.SetLineno(64)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_garbage_padding", "build/src/__python__/test/test_uu.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µencodedtext *πg.Object = πg.UnboundLocal; _ = µencodedtext
						var µplaintext *πg.Object = πg.UnboundLocal; _ = µplaintext
						var µinp *πg.Object = πg.UnboundLocal; _ = µinp
						var µout *πg.Object = πg.UnboundLocal; _ = µout
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
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
							// line 66: encodedtext = (
							πF.SetLineno(66)
							µencodedtext = πg.NewStr("begin 644 file\n!,___\n \nend\n").ToObject()
							// line 73: plaintext = "\x33"  # 00110011
							πF.SetLineno(73)
							µplaintext = ß3.ToObject()
							// line 75: inp = cStringIO.StringIO(encodedtext)
							πF.SetLineno(75)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µencodedtext, "encodedtext"); πE != nil {
								continue
							}
							πTemp001[0] = µencodedtext
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µinp = πTemp002
							// line 76: out = cStringIO.StringIO()
							πF.SetLineno(76)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µout = πTemp002
							// line 77: uu.decode(inp, out, quiet=True)
							πF.SetLineno(77)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µinp, "inp"); πE != nil {
								continue
							}
							πTemp001[0] = µinp
							if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
								continue
							}
							πTemp001[1] = µout
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"quiet", πTemp002},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßuu); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdecode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, πTemp004); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 78: self.assertEqual(out.getvalue(), plaintext)
							πF.SetLineno(78)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µout, ßgetvalue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µplaintext, "plaintext"); πE != nil {
								continue
							}
							πTemp001[1] = µplaintext
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
					if πE = πClass.SetItem(πF, ßtest_garbage_padding.ToObject(), πTemp006); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("UUTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßUUTest.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 84: class UUStdIOTest(unittest.TestCase):
			πF.SetLineno(84)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßTestCase, nil); πE != nil {
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
			_, πE = πg.NewCode("UUStdIOTest", "build/src/__python__/test/test_uu.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 86: def setUp(self):
					πF.SetLineno(86)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("setUp", "build/src/__python__/test/test_uu.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 87: self.stdin = sys.stdin
							πF.SetLineno(87)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstdin, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstdin, πTemp001); πE != nil {
								continue
							}
							// line 88: self.stdout = sys.stdout
							πF.SetLineno(88)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstdout, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstdout, πTemp001); πE != nil {
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
					// line 90: def tearDown(self):
					πF.SetLineno(90)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("tearDown", "build/src/__python__/test/test_uu.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
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
							// line 91: sys.stdin = self.stdin
							πF.SetLineno(91)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstdin, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp003, ßstdin, πTemp002); πE != nil {
								continue
							}
							// line 92: sys.stdout = self.stdout
							πF.SetLineno(92)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstdout, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp003, ßstdout, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtearDown.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 94: def test_encode(self):
					πF.SetLineno(94)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_encode", "build/src/__python__/test/test_uu.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 95: sys.stdin = cStringIO.StringIO(plaintext)
							πF.SetLineno(95)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßplaintext); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp004, ßstdin, πTemp003); πE != nil {
								continue
							}
							// line 96: sys.stdout = cStringIO.StringIO()
							πF.SetLineno(96)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp004, ßstdout, πTemp003); πE != nil {
								continue
							}
							// line 97: uu.encode("-", "-", "t1", 0666)
							πF.SetLineno(97)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewStr("-").ToObject()
							πTemp001[1] = πg.NewStr("-").ToObject()
							πTemp001[2] = ßt1.ToObject()
							πTemp001[3] = πg.NewInt(438).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßuu); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßencode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 98: self.assertEqual(
							πF.SetLineno(98)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstdout, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßgetvalue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßencodedtextwrapped); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(πg.NewInt(438).ToObject(), ßt1.ToObject()).ToObject()
							if πTemp002, πE = πg.Mod(πF, πTemp003, πTemp004); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_encode.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 103: def test_decode(self):
					πF.SetLineno(103)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_decode", "build/src/__python__/test/test_uu.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 104: sys.stdin = cStringIO.StringIO(encodedtextwrapped % (0666, "t1"))
							πF.SetLineno(104)
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßencodedtextwrapped); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(πg.NewInt(438).ToObject(), ßt1.ToObject()).ToObject()
							if πTemp002, πE = πg.Mod(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp004, ßstdin, πTemp003); πE != nil {
								continue
							}
							// line 105: sys.stdout = cStringIO.StringIO()
							πF.SetLineno(105)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp004, ßstdout, πTemp003); πE != nil {
								continue
							}
							// line 106: uu.decode("-", "-")
							πF.SetLineno(106)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("-").ToObject()
							πTemp001[1] = πg.NewStr("-").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßuu); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdecode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 107: self.assertEqual(sys.stdout.getvalue(), plaintext)
							πF.SetLineno(107)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstdout, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßgetvalue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßplaintext); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_decode.ToObject(), πTemp005); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("UUStdIOTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßUUStdIOTest.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 109: class UUFileTest(unittest.TestCase):
			πF.SetLineno(109)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßTestCase, nil); πE != nil {
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
			_, πE = πg.NewCode("UUFileTest", "build/src/__python__/test/test_uu.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 111: def _kill(self, f):
					πF.SetLineno(111)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "f", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("_kill", "build/src/__python__/test/test_uu.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µf *πg.Object = πArgs[1]; _ = µf
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.BaseException
						_ = πTemp003
						var πTemp004 *πg.Traceback
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							case 6: goto Label6
							default: panic("unexpected function state")
							}
							// line 113: try:
							πF.SetLineno(113)
							πF.PushCheckpoint(2)
							// line 114: f.close()
							πF.SetLineno(114)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSystemExit); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßKeyboardInterrupt); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πTemp005).ToObject()
							if πTemp006, πE = πg.IsInstance(πF, πTemp003.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							goto Label4
							// line 115: except (SystemExit, KeyboardInterrupt):
							πF.SetLineno(115)
						Label3:
							// line 116: raise
							πF.SetLineno(116)
							πE = πF.Raise(nil, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
							// line 117: except:
							πF.SetLineno(117)
						Label4:
							// line 118: pass
							πF.SetLineno(118)
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 119: try:
							πF.SetLineno(119)
							πF.PushCheckpoint(6)
							// line 120: os.unlink(f.name)
							πF.SetLineno(120)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µf, ßname, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßunlink, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πF.PopCheckpoint()
							goto Label5
						Label6:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp003, πTemp004 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSystemExit); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßKeyboardInterrupt); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πTemp005).ToObject()
							if πTemp006, πE = πg.IsInstance(πF, πTemp003.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label7
							}
							πE = πF.Raise(πTemp003.ToObject(), nil, πTemp004.ToObject())
							continue
							// line 121: except (SystemExit, KeyboardInterrupt):
							πF.SetLineno(121)
						Label7:
							// line 122: raise
							πF.SetLineno(122)
							πE = πF.Raise(nil, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label5
						Label5:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_kill.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 126: def setUp(self):
					πF.SetLineno(126)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("setUp", "build/src/__python__/test/test_uu.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
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
							// line 127: self.tmpin  = test_support.TESTFN + "i"
							πF.SetLineno(127)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßTESTFN, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, ßi.ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtmpin, πTemp002); πE != nil {
								continue
							}
							// line 128: self.tmpout = test_support.TESTFN + "o"
							πF.SetLineno(128)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßTESTFN, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, ßo.ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtmpout, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsetUp.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 130: def tearDown(self):
					πF.SetLineno(130)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("tearDown", "build/src/__python__/test/test_uu.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 131: del self.tmpin
							πF.SetLineno(131)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.DelAttr(πF, µself, ßtmpin); πE != nil {
								continue
							}
							// line 132: del self.tmpout
							πF.SetLineno(132)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.DelAttr(πF, µself, ßtmpout); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtearDown.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 134: def test_encode(self):
					πF.SetLineno(134)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_encode", "build/src/__python__/test/test_uu.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfin *πg.Object = πg.UnboundLocal; _ = µfin
						var µfout *πg.Object = πg.UnboundLocal; _ = µfout
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 πg.KWArgs
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 135: fin = fout = None
							πF.SetLineno(135)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µfin = πTemp001
							µfout = πTemp001
							// line 136: try:
							πF.SetLineno(136)
							πF.PushCheckpoint(1)
							// line 137: test_support.unlink(self.tmpin)
							πF.SetLineno(137)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtmpin, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßunlink, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 138: fin = open(self.tmpin, 'wb')
							πF.SetLineno(138)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtmpin, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							πTemp002[1] = ßwb.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µfin = πTemp003
							// line 139: fin.write(plaintext)
							πF.SetLineno(139)
							πTemp002 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßplaintext); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µfin, "fin"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfin, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 140: fin.close()
							πF.SetLineno(140)
							if πE = πg.CheckLocal(πF, µfin, "fin"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfin, ßclose, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 142: fin = open(self.tmpin, 'rb')
							πF.SetLineno(142)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtmpin, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							πTemp002[1] = ßrb.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µfin = πTemp003
							// line 143: fout = open(self.tmpout, 'w')
							πF.SetLineno(143)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtmpout, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							πTemp002[1] = ßw.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µfout = πTemp003
							// line 144: uu.encode(fin, fout, self.tmpin, mode=0644)
							πF.SetLineno(144)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µfin, "fin"); πE != nil {
								continue
							}
							πTemp002[0] = µfin
							if πE = πg.CheckLocal(πF, µfout, "fout"); πE != nil {
								continue
							}
							πTemp002[1] = µfout
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtmpin, nil); πE != nil {
								continue
							}
							πTemp002[2] = πTemp001
							πTemp004 = πg.KWArgs{
								{"mode", πg.NewInt(420).ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßuu); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßencode, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, πTemp004); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 145: fin.close()
							πF.SetLineno(145)
							if πE = πg.CheckLocal(πF, µfin, "fin"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfin, ßclose, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 146: fout.close()
							πF.SetLineno(146)
							if πE = πg.CheckLocal(πF, µfout, "fout"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfout, ßclose, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 148: fout = open(self.tmpout, 'r')
							πF.SetLineno(148)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtmpout, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							πTemp002[1] = ßr.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µfout = πTemp003
							// line 149: s = fout.read()
							πF.SetLineno(149)
							if πE = πg.CheckLocal(πF, µfout, "fout"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfout, ßread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µs = πTemp003
							// line 150: fout.close()
							πF.SetLineno(150)
							if πE = πg.CheckLocal(πF, µfout, "fout"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfout, ßclose, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 151: self.assertEqual(s, encodedtextwrapped % (0644, self.tmpin))
							πF.SetLineno(151)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp002[0] = µs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßencodedtextwrapped); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßtmpin, nil); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(πg.NewInt(420).ToObject(), πTemp006).ToObject()
							if πTemp001, πE = πg.Mod(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 154: uu.encode(self.tmpin, self.tmpout, self.tmpin, mode=0644)
							πF.SetLineno(154)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtmpin, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtmpout, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtmpin, nil); πE != nil {
								continue
							}
							πTemp002[2] = πTemp001
							πTemp004 = πg.KWArgs{
								{"mode", πg.NewInt(420).ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßuu); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßencode, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, πTemp004); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 155: fout = open(self.tmpout, 'r')
							πF.SetLineno(155)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtmpout, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							πTemp002[1] = ßr.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µfout = πTemp003
							// line 156: s = fout.read()
							πF.SetLineno(156)
							if πE = πg.CheckLocal(πF, µfout, "fout"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfout, ßread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µs = πTemp003
							// line 157: fout.close()
							πF.SetLineno(157)
							if πE = πg.CheckLocal(πF, µfout, "fout"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfout, ßclose, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 158: self.assertEqual(s, encodedtextwrapped % (0644, self.tmpin))
							πF.SetLineno(158)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp002[0] = µs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßencodedtextwrapped); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßtmpin, nil); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(πg.NewInt(420).ToObject(), πTemp006).ToObject()
							if πTemp001, πE = πg.Mod(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πF.PopCheckpoint()
							goto Label1
						Label1:
							πTemp007, πTemp008 = πF.RestoreExc(nil, nil)
							// line 161: self._kill(fin)
							πF.SetLineno(161)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfin, "fin"); πE != nil {
								continue
							}
							πTemp002[0] = µfin
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_kill, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 162: self._kill(fout)
							πF.SetLineno(162)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfout, "fout"); πE != nil {
								continue
							}
							πTemp002[0] = µfout
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_kill, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp007 != nil {
								πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
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
					if πE = πClass.SetItem(πF, ßtest_encode.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 164: def test_decode(self):
					πF.SetLineno(164)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_decode", "build/src/__python__/test/test_uu.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var µs *πg.Object = πg.UnboundLocal; _ = µs
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
						var πTemp006 *πg.BaseException
						_ = πTemp006
						var πTemp007 *πg.Traceback
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 165: f = None
							πF.SetLineno(165)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µf = πTemp001
							// line 166: try:
							πF.SetLineno(166)
							πF.PushCheckpoint(1)
							// line 167: test_support.unlink(self.tmpin)
							πF.SetLineno(167)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtmpin, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßunlink, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 168: f = open(self.tmpin, 'w')
							πF.SetLineno(168)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtmpin, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							πTemp002[1] = ßw.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µf = πTemp003
							// line 169: f.write(encodedtextwrapped % (0644, self.tmpout))
							πF.SetLineno(169)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßencodedtextwrapped); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßtmpout, nil); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(πg.NewInt(420).ToObject(), πTemp005).ToObject()
							if πTemp001, πE = πg.Mod(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µf, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 170: f.close()
							πF.SetLineno(170)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 172: f = open(self.tmpin, 'r')
							πF.SetLineno(172)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtmpin, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							πTemp002[1] = ßr.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µf = πTemp003
							// line 173: uu.decode(f)
							πF.SetLineno(173)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp002[0] = µf
							if πTemp001, πE = πg.ResolveGlobal(πF, ßuu); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßdecode, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 174: f.close()
							πF.SetLineno(174)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 176: f = open(self.tmpout, 'r')
							πF.SetLineno(176)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtmpout, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							πTemp002[1] = ßr.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µf = πTemp003
							// line 177: s = f.read()
							πF.SetLineno(177)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µf, ßread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µs = πTemp003
							// line 178: f.close()
							πF.SetLineno(178)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 179: self.assertEqual(s, plaintext)
							πF.SetLineno(179)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp002[0] = µs
							if πTemp001, πE = πg.ResolveGlobal(πF, ßplaintext); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πF.PopCheckpoint()
							goto Label1
						Label1:
							πTemp006, πTemp007 = πF.RestoreExc(nil, nil)
							// line 182: self._kill(f)
							πF.SetLineno(182)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp002[0] = µf
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_kill, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp006 != nil {
								πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
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
					if πE = πClass.SetItem(πF, ßtest_decode.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 184: def test_decode_filename(self):
					πF.SetLineno(184)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_decode_filename", "build/src/__python__/test/test_uu.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var µs *πg.Object = πg.UnboundLocal; _ = µs
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
						var πTemp006 *πg.BaseException
						_ = πTemp006
						var πTemp007 *πg.Traceback
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 185: f = None
							πF.SetLineno(185)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µf = πTemp001
							// line 186: try:
							πF.SetLineno(186)
							πF.PushCheckpoint(1)
							// line 187: test_support.unlink(self.tmpin)
							πF.SetLineno(187)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtmpin, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßunlink, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 188: f = open(self.tmpin, 'w')
							πF.SetLineno(188)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtmpin, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							πTemp002[1] = ßw.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µf = πTemp003
							// line 189: f.write(encodedtextwrapped % (0644, self.tmpout))
							πF.SetLineno(189)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßencodedtextwrapped); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßtmpout, nil); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(πg.NewInt(420).ToObject(), πTemp005).ToObject()
							if πTemp001, πE = πg.Mod(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µf, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 190: f.close()
							πF.SetLineno(190)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 192: uu.decode(self.tmpin)
							πF.SetLineno(192)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtmpin, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßuu); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßdecode, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 194: f = open(self.tmpout, 'r')
							πF.SetLineno(194)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtmpout, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							πTemp002[1] = ßr.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µf = πTemp003
							// line 195: s = f.read()
							πF.SetLineno(195)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µf, ßread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µs = πTemp003
							// line 196: f.close()
							πF.SetLineno(196)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 197: self.assertEqual(s, plaintext)
							πF.SetLineno(197)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp002[0] = µs
							if πTemp001, πE = πg.ResolveGlobal(πF, ßplaintext); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πF.PopCheckpoint()
							goto Label1
						Label1:
							πTemp006, πTemp007 = πF.RestoreExc(nil, nil)
							// line 199: self._kill(f)
							πF.SetLineno(199)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp002[0] = µf
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_kill, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp006 != nil {
								πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
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
					if πE = πClass.SetItem(πF, ßtest_decode_filename.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 201: def test_decodetwice(self):
					πF.SetLineno(201)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("test_decodetwice", "build/src/__python__/test/test_uu.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µf *πg.Object = πg.UnboundLocal; _ = µf
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
						var πTemp006 *πg.BaseException
						_ = πTemp006
						var πTemp007 *πg.Traceback
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 203: f = None
							πF.SetLineno(203)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µf = πTemp001
							// line 204: try:
							πF.SetLineno(204)
							πF.PushCheckpoint(1)
							// line 205: f = cStringIO.StringIO(encodedtextwrapped % (0644, self.tmpout))
							πF.SetLineno(205)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßencodedtextwrapped); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßtmpout, nil); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(πg.NewInt(420).ToObject(), πTemp005).ToObject()
							if πTemp001, πE = πg.Mod(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßcStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µf = πTemp001
							// line 207: f = open(self.tmpin, 'r')
							πF.SetLineno(207)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtmpin, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							πTemp002[1] = ßr.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µf = πTemp003
							// line 208: uu.decode(f)
							πF.SetLineno(208)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp002[0] = µf
							if πTemp001, πE = πg.ResolveGlobal(πF, ßuu); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßdecode, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 209: f.close()
							πF.SetLineno(209)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 211: f = open(self.tmpin, 'r')
							πF.SetLineno(211)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtmpin, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							πTemp002[1] = ßr.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µf = πTemp003
							// line 212: self.assertRaises(uu.Error, uu.decode, f)
							πF.SetLineno(212)
							πTemp002 = πF.MakeArgs(3)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßuu); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßError, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ßuu); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßdecode, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp002[2] = µf
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 213: f.close()
							πF.SetLineno(213)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label1
						Label1:
							πTemp006, πTemp007 = πF.RestoreExc(nil, nil)
							// line 215: self._kill(f)
							πF.SetLineno(215)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp002[0] = µf
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_kill, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp006 != nil {
								πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
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
					if πE = πClass.SetItem(πF, ßtest_decodetwice.ToObject(), πTemp008); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("UUFileTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßUUFileTest.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 217: def test_main():
			πF.SetLineno(217)
			πTemp007 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("test_main", "build/src/__python__/test/test_uu.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 218: test_support.run_unittest(UUTest, UUStdIOTest, UUFileTest)
					πF.SetLineno(218)
					πTemp001 = πF.MakeArgs(3)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßUUTest); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßUUStdIOTest); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßUUFileTest); πE != nil {
						continue
					}
					πTemp001[2] = πTemp002
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
			if πTemp004, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp003, πE = πg.Eq(πF, πTemp004, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
				continue
			}
			if πTemp008 {
				goto Label1
			}
			goto Label2
			// line 220: if __name__=="__main__":
			πF.SetLineno(220)
		Label1:
			// line 221: test_main()
			πF.SetLineno(221)
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
	πg.RegisterModule("test.test_uu", Code)
}
