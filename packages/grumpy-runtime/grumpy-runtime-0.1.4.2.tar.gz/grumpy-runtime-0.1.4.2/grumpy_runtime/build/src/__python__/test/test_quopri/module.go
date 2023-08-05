package test_quopri
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/test/test_quopri.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßDECSAMPLE := πg.InternStr("DECSAMPLE")
		ßENCSAMPLE := πg.InternStr("ENCSAMPLE")
		ßESTRINGS := πg.InternStr("ESTRINGS")
		ßFalse := πg.InternStr("False")
		ßHSTRINGS := πg.InternStr("HSTRINGS")
		ßNone := πg.InternStr("None")
		ßPIPE := πg.InternStr("PIPE")
		ßPopen := πg.InternStr("Popen")
		ßQuopriTestCase := πg.InternStr("QuopriTestCase")
		ßSTRINGS := πg.InternStr("STRINGS")
		ßStringIO := πg.InternStr("StringIO")
		ßTestCase := πg.InternStr("TestCase")
		ßTrue := πg.InternStr("True")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßa2b_qp := πg.InternStr("a2b_qp")
		ßaddCleanup := πg.InternStr("addCleanup")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertTrue := πg.InternStr("assertTrue")
		ßb2a_qp := πg.InternStr("b2a_qp")
		ßcStringIO := πg.InternStr("cStringIO")
		ßclose := πg.InternStr("close")
		ßcommunicate := πg.InternStr("communicate")
		ßdecode := πg.InternStr("decode")
		ßdecodestring := πg.InternStr("decodestring")
		ßencode := πg.InternStr("encode")
		ßencodestring := πg.InternStr("encodestring")
		ßexecutable := πg.InternStr("executable")
		ßexpectedFailure := πg.InternStr("expectedFailure")
		ßgetvalue := πg.InternStr("getvalue")
		ßhello := πg.InternStr("hello")
		ßhello_world := πg.InternStr("hello_world")
		ßquopri := πg.InternStr("quopri")
		ßrun_unittest := πg.InternStr("run_unittest")
		ßsplitlines := πg.InternStr("splitlines")
		ßstdout := πg.InternStr("stdout")
		ßsubprocess := πg.InternStr("subprocess")
		ßsys := πg.InternStr("sys")
		ßtest_decode := πg.InternStr("test_decode")
		ßtest_decode_header := πg.InternStr("test_decode_header")
		ßtest_decodestring := πg.InternStr("test_decodestring")
		ßtest_embedded_ws := πg.InternStr("test_embedded_ws")
		ßtest_encode := πg.InternStr("test_encode")
		ßtest_encode_header := πg.InternStr("test_encode_header")
		ßtest_encodestring := πg.InternStr("test_encodestring")
		ßtest_idempotent_string := πg.InternStr("test_idempotent_string")
		ßtest_main := πg.InternStr("test_main")
		ßtest_scriptdecode := πg.InternStr("test_scriptdecode")
		ßtest_scriptencode := πg.InternStr("test_scriptencode")
		ßtest_support := πg.InternStr("test_support")
		ßunittest := πg.InternStr("unittest")
		ßwithpythonimplementation := πg.InternStr("withpythonimplementation")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 []πg.Param
		_ = πTemp003
		var πTemp004 *πg.Dict
		_ = πTemp004
		var πTemp005 *πg.Object
		_ = πTemp005
		var πTemp006 *πg.Object
		_ = πTemp006
		var πTemp007 *πg.Object
		_ = πTemp007
		var πTemp008 *πg.Object
		_ = πTemp008
		var πTemp009 bool
		_ = πTemp009
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: from test import test_support
			πF.SetLineno(1)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßtest_support.ToObject(), πTemp001); πE != nil {
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
			// line 4: import sys, cStringIO #, subprocess
			πF.SetLineno(4)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "cStringIO"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßcStringIO.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 5: import quopri
			πF.SetLineno(5)
			if πTemp002, πE = πg.ImportModule(πF, "quopri"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßquopri.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 9: ENCSAMPLE = """\
			πF.SetLineno(9)
			if πE = πF.Globals().SetItem(πF, ßENCSAMPLE.ToObject(), πg.NewStr("Here's a bunch of special=20\n\n=A1=A2=A3=A4=A5=A6=A7=A8=A9\n=AA=AB=AC=AD=AE=AF=B0=B1=B2=B3\n=B4=B5=B6=B7=B8=B9=BA=BB=BC=BD=BE\n=BF=C0=C1=C2=C3=C4=C5=C6\n=C7=C8=C9=CA=CB=CC=CD=CE=CF\n=D0=D1=D2=D3=D4=D5=D6=D7\n=D8=D9=DA=DB=DC=DD=DE=DF\n=E0=E1=E2=E3=E4=E5=E6=E7\n=E8=E9=EA=EB=EC=ED=EE=EF\n=F0=F1=F2=F3=F4=F5=F6=F7\n=F8=F9=FA=FB=FC=FD=FE=FF\n\ncharacters... have fun!\n").ToObject()); πE != nil {
				continue
			}
			// line 28: DECSAMPLE = "Here's a bunch of special \n" + \
			πF.SetLineno(28)
			if πTemp001, πE = πg.Add(πF, πg.NewStr("Here's a bunch of special \n").ToObject(), πg.NewStr("\n\xa1\xa2\xa3\xa4\xa5\xa6\xa7\xa8\xa9\n\xaa\xab\xac\xad\xae\xaf\xb0\xb1\xb2\xb3\n\xb4\xb5\xb6\xb7\xb8\xb9\xba\xbb\xbc\xbd\xbe\n\xbf\xc0\xc1\xc2\xc3\xc4\xc5\xc6\n\xc7\xc8\xc9\xca\xcb\xcc\xcd\xce\xcf\n\xd0\xd1\xd2\xd3\xd4\xd5\xd6\xd7\n\xd8\xd9\xda\xdb\xdc\xdd\xde\xdf\n\xe0\xe1\xe2\xe3\xe4\xe5\xe6\xe7\n\xe8\xe9\xea\xeb\xec\xed\xee\xef\n\xf0\xf1\xf2\xf3\xf4\xf5\xf6\xf7\n\xf8\xf9\xfa\xfb\xfc\xfd\xfe\xff\n\ncharacters... have fun!\n").ToObject()); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDECSAMPLE.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 47: def withpythonimplementation(testfunc):
			πF.SetLineno(47)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "testfunc", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("withpythonimplementation", "build/src/__python__/test/test_quopri.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtestfunc *πg.Object = πArgs[0]; _ = µtestfunc
				var µnewtest *πg.Object = πg.UnboundLocal; _ = µnewtest
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 48: def newtest(self):
					πF.SetLineno(48)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("newtest", "build/src/__python__/test/test_quopri.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µoldencode *πg.Object = πg.UnboundLocal; _ = µoldencode
						var µolddecode *πg.Object = πg.UnboundLocal; _ = µolddecode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
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
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							// line 50: testfunc(self)
							πF.SetLineno(50)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µtestfunc, "testfunc"); πE != nil {
								continue
							}
							if πTemp002, πE = µtestfunc.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßquopri); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßb2a_qp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(πTemp006 != πTemp005).ToObject()
							πTemp002 = πTemp004
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßquopri); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßa2b_qp, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(πTemp006 != πTemp005).ToObject()
							πTemp002 = πTemp004
						Label1:
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label2
							}
							goto Label3
							// line 52: if quopri.b2a_qp is not None or quopri.a2b_qp is not None:
							πF.SetLineno(52)
						Label2:
							// line 53: oldencode = quopri.b2a_qp
							πF.SetLineno(53)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßquopri); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßb2a_qp, nil); πE != nil {
								continue
							}
							µoldencode = πTemp004
							// line 54: olddecode = quopri.a2b_qp
							πF.SetLineno(54)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßquopri); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßa2b_qp, nil); πE != nil {
								continue
							}
							µolddecode = πTemp004
							// line 55: try:
							πF.SetLineno(55)
							πF.PushCheckpoint(4)
							// line 56: quopri.b2a_qp = None
							πF.SetLineno(56)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßquopri); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp005, ßb2a_qp, πTemp004); πE != nil {
								continue
							}
							// line 57: quopri.a2b_qp = None
							πF.SetLineno(57)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßquopri); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp005, ßa2b_qp, πTemp004); πE != nil {
								continue
							}
							// line 58: testfunc(self)
							πF.SetLineno(58)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µtestfunc, "testfunc"); πE != nil {
								continue
							}
							if πTemp002, πE = µtestfunc.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πF.PopCheckpoint()
							goto Label4
						Label4:
							πTemp007, πTemp008 = πF.RestoreExc(nil, nil)
							// line 60: quopri.b2a_qp = oldencode
							πF.SetLineno(60)
							if πE = πg.CheckLocal(πF, µoldencode, "oldencode"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µoldencode); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßquopri); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp004, ßb2a_qp, πTemp002); πE != nil {
								continue
							}
							// line 61: quopri.a2b_qp = olddecode
							πF.SetLineno(61)
							if πE = πg.CheckLocal(πF, µolddecode, "olddecode"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µolddecode); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßquopri); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp004, ßa2b_qp, πTemp002); πE != nil {
								continue
							}
							if πTemp007 != nil {
								πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
								continue
							}
							if πR != nil {
								continue
							}
							goto Label3
						Label3:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					µnewtest = πTemp001
					// line 63: return newtest
					πF.SetLineno(63)
					if πE = πg.CheckLocal(πF, µnewtest, "newtest"); πE != nil {
						continue
					}
					πR = µnewtest
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßwithpythonimplementation.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 65: class QuopriTestCase(unittest.TestCase):
			πF.SetLineno(65)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßTestCase, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp004 = πg.NewDict()
			if πTemp005, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp005); πE != nil {
				continue
			}
			_, πE = πg.NewCode("QuopriTestCase", "build/src/__python__/test/test_quopri.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
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
				var πTemp006 []πg.Param
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 68: STRINGS = (
					πF.SetLineno(68)
					πTemp002 = make([]*πg.Object, 12)
					πTemp003 = πg.NewTuple2(ßhello.ToObject(), ßhello.ToObject()).ToObject()
					πTemp002[0] = πTemp003
					πTemp003 = πg.NewTuple2(πg.NewStr("hello\n        there\n        world").ToObject(), πg.NewStr("hello\n        there\n        world").ToObject()).ToObject()
					πTemp002[1] = πTemp003
					πTemp003 = πg.NewTuple2(πg.NewStr("hello\n        there\n        world\n").ToObject(), πg.NewStr("hello\n        there\n        world\n").ToObject()).ToObject()
					πTemp002[2] = πTemp003
					πTemp003 = πg.NewTuple2(πg.NewStr("\x81\x82\x83").ToObject(), πg.NewStr("=81=82=83").ToObject()).ToObject()
					πTemp002[3] = πTemp003
					πTemp003 = πg.NewTuple2(πg.NewStr("hello ").ToObject(), πg.NewStr("hello=20").ToObject()).ToObject()
					πTemp002[4] = πTemp003
					πTemp003 = πg.NewTuple2(πg.NewStr("hello\t").ToObject(), πg.NewStr("hello=09").ToObject()).ToObject()
					πTemp002[5] = πTemp003
					πTemp003 = πg.NewTuple2(πg.NewStr("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\xd8\xd9\xda\xdb\xdc\xdd\xde\xdfxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx").ToObject(), πg.NewStr("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx=D8=D9=DA=DB=DC=DD=DE=DFx=\nxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx").ToObject()).ToObject()
					πTemp002[6] = πTemp003
					πTemp003 = πg.NewTuple2(πg.NewStr("yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy").ToObject(), πg.NewStr("yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy").ToObject()).ToObject()
					πTemp002[7] = πTemp003
					πTemp003 = πg.NewTuple2(πg.NewStr("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz").ToObject(), πg.NewStr("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz=\nzz").ToObject()).ToObject()
					πTemp002[8] = πTemp003
					πTemp003 = πg.NewTuple2(πg.NewStr("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz").ToObject(), πg.NewStr("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz=\nzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz").ToObject()).ToObject()
					πTemp002[9] = πTemp003
					πTemp003 = πg.NewTuple2(πg.NewStr("yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy\nzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz").ToObject(), πg.NewStr("yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy=\nyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy\nzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz").ToObject()).ToObject()
					πTemp002[10] = πTemp003
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßDECSAMPLE); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßENCSAMPLE); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
					πTemp002[11] = πTemp003
					πTemp001 = πg.NewTuple(πTemp002...).ToObject()
					if πE = πClass.SetItem(πF, ßSTRINGS.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 119: ESTRINGS = (
					πF.SetLineno(119)
					πTemp003 = πg.NewTuple2(πg.NewStr("hello world").ToObject(), πg.NewStr("hello=20world").ToObject()).ToObject()
					πTemp004 = πg.NewTuple2(πg.NewStr("hello\tworld").ToObject(), πg.NewStr("hello=09world").ToObject()).ToObject()
					πTemp001 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					if πE = πClass.SetItem(πF, ßESTRINGS.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 125: HSTRINGS = (
					πF.SetLineno(125)
					πTemp003 = πg.NewTuple2(πg.NewStr("hello world").ToObject(), ßhello_world.ToObject()).ToObject()
					πTemp004 = πg.NewTuple2(ßhello_world.ToObject(), πg.NewStr("hello=5Fworld").ToObject()).ToObject()
					πTemp001 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					if πE = πClass.SetItem(πF, ßHSTRINGS.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 131: def test_encodestring(self):
					πF.SetLineno(131)
					πTemp006 = make([]πg.Param, 1)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_encodestring", "build/src/__python__/test/test_quopri.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µp *πg.Object = πg.UnboundLocal; _ = µp
						var µe *πg.Object = πg.UnboundLocal; _ = µe
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßSTRINGS, nil); πE != nil {
								continue
							}
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
									continue
								}
								µp = πTemp005
								µe = πTemp006
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 133: self.assertTrue(quopri.encodestring(p) == e)
							πF.SetLineno(133)
							πTemp007 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp008[0] = µp
							if πTemp005, πE = πg.ResolveGlobal(πF, ßquopri); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßencodestring, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp005, µe); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
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
					if πE = πClass.SetItem(πF, ßtest_encodestring.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 130: @withpythonimplementation
					πF.SetLineno(130)
					πTemp002 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßtest_encodestring); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßwithpythonimplementation); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πClass.SetItem(πF, ßtest_encodestring.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 136: def test_decodestring(self):
					πF.SetLineno(136)
					πTemp006 = make([]πg.Param, 1)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_decodestring", "build/src/__python__/test/test_quopri.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µp *πg.Object = πg.UnboundLocal; _ = µp
						var µe *πg.Object = πg.UnboundLocal; _ = µe
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßSTRINGS, nil); πE != nil {
								continue
							}
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
									continue
								}
								µp = πTemp005
								µe = πTemp006
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 138: self.assertTrue(quopri.decodestring(e) == p)
							πF.SetLineno(138)
							πTemp007 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							πTemp008[0] = µe
							if πTemp005, πE = πg.ResolveGlobal(πF, ßquopri); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßdecodestring, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp005, µp); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
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
					if πE = πClass.SetItem(πF, ßtest_decodestring.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 135: @withpythonimplementation
					πF.SetLineno(135)
					πTemp002 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßtest_decodestring); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßwithpythonimplementation); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πClass.SetItem(πF, ßtest_decodestring.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 141: def test_idempotent_string(self):
					πF.SetLineno(141)
					πTemp006 = make([]πg.Param, 1)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_idempotent_string", "build/src/__python__/test/test_quopri.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µp *πg.Object = πg.UnboundLocal; _ = µp
						var µe *πg.Object = πg.UnboundLocal; _ = µe
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
						var πTemp006 *πg.Object
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
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßSTRINGS, nil); πE != nil {
								continue
							}
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
									continue
								}
								µp = πTemp005
								µe = πTemp006
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 143: self.assertTrue(quopri.decodestring(quopri.encodestring(e)) == e)
							πF.SetLineno(143)
							πTemp007 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							πTemp009[0] = µe
							if πTemp005, πE = πg.ResolveGlobal(πF, ßquopri); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßencodestring, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp008[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßquopri); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßdecodestring, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp005, µe); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
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
					if πE = πClass.SetItem(πF, ßtest_idempotent_string.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 140: @withpythonimplementation
					πF.SetLineno(140)
					πTemp002 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßtest_idempotent_string); πE != nil {
						continue
					}
					πTemp002[0] = πTemp005
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßwithpythonimplementation); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πClass.SetItem(πF, ßtest_idempotent_string.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 146: def test_encode(self):
					πF.SetLineno(146)
					πTemp006 = make([]πg.Param, 1)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_encode", "build/src/__python__/test/test_quopri.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µp *πg.Object = πg.UnboundLocal; _ = µp
						var µe *πg.Object = πg.UnboundLocal; _ = µe
						var µinfp *πg.Object = πg.UnboundLocal; _ = µinfp
						var µoutfp *πg.Object = πg.UnboundLocal; _ = µoutfp
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 πg.KWArgs
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßSTRINGS, nil); πE != nil {
								continue
							}
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
									continue
								}
								µp = πTemp005
								µe = πTemp006
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 148: infp = cStringIO.StringIO(p)
							πF.SetLineno(148)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp007[0] = µp
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcStringIO); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µinfp = πTemp002
							// line 149: outfp = cStringIO.StringIO()
							πF.SetLineno(149)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcStringIO); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							µoutfp = πTemp002
							// line 150: quopri.encode(infp, outfp, quotetabs=False)
							πF.SetLineno(150)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µinfp, "infp"); πE != nil {
								continue
							}
							πTemp007[0] = µinfp
							if πE = πg.CheckLocal(πF, µoutfp, "outfp"); πE != nil {
								continue
							}
							πTemp007[1] = µoutfp
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
								{"quotetabs", πTemp002},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßquopri); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßencode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, πTemp007, πTemp008); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 151: self.assertTrue(outfp.getvalue() == e)
							πF.SetLineno(151)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µoutfp, "outfp"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µoutfp, ßgetvalue, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp006, µe); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
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
					if πE = πClass.SetItem(πF, ßtest_encode.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 145: @withpythonimplementation
					πF.SetLineno(145)
					πTemp002 = πF.MakeArgs(1)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßtest_encode); πE != nil {
						continue
					}
					πTemp002[0] = πTemp007
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßwithpythonimplementation); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πClass.SetItem(πF, ßtest_encode.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 154: def test_decode(self):
					πF.SetLineno(154)
					πTemp006 = make([]πg.Param, 1)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_decode", "build/src/__python__/test/test_quopri.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µp *πg.Object = πg.UnboundLocal; _ = µp
						var µe *πg.Object = πg.UnboundLocal; _ = µe
						var µinfp *πg.Object = πg.UnboundLocal; _ = µinfp
						var µoutfp *πg.Object = πg.UnboundLocal; _ = µoutfp
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßSTRINGS, nil); πE != nil {
								continue
							}
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
									continue
								}
								µp = πTemp005
								µe = πTemp006
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 156: infp = cStringIO.StringIO(e)
							πF.SetLineno(156)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							πTemp007[0] = µe
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcStringIO); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µinfp = πTemp002
							// line 157: outfp = cStringIO.StringIO()
							πF.SetLineno(157)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcStringIO); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							µoutfp = πTemp002
							// line 158: quopri.decode(infp, outfp)
							πF.SetLineno(158)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µinfp, "infp"); πE != nil {
								continue
							}
							πTemp007[0] = µinfp
							if πE = πg.CheckLocal(πF, µoutfp, "outfp"); πE != nil {
								continue
							}
							πTemp007[1] = µoutfp
							if πTemp002, πE = πg.ResolveGlobal(πF, ßquopri); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßdecode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 159: self.assertTrue(outfp.getvalue() == p)
							πF.SetLineno(159)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µoutfp, "outfp"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µoutfp, ßgetvalue, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp006, µp); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
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
					if πE = πClass.SetItem(πF, ßtest_decode.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 153: @withpythonimplementation
					πF.SetLineno(153)
					πTemp002 = πF.MakeArgs(1)
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßtest_decode); πE != nil {
						continue
					}
					πTemp002[0] = πTemp008
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßwithpythonimplementation); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp008.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πClass.SetItem(πF, ßtest_decode.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 162: def test_embedded_ws(self):
					πF.SetLineno(162)
					πTemp006 = make([]πg.Param, 1)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("test_embedded_ws", "build/src/__python__/test/test_quopri.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µp *πg.Object = πg.UnboundLocal; _ = µp
						var µe *πg.Object = πg.UnboundLocal; _ = µe
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 πg.KWArgs
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßESTRINGS, nil); πE != nil {
								continue
							}
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
									continue
								}
								µp = πTemp005
								µe = πTemp006
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 164: self.assertTrue(quopri.encodestring(p, quotetabs=True) == e)
							πF.SetLineno(164)
							πTemp007 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp008[0] = µp
							if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"quotetabs", πTemp005},
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßquopri); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßencodestring, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, πTemp008, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp005, µe); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 165: self.assertTrue(quopri.decodestring(e) == p)
							πF.SetLineno(165)
							πTemp007 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							πTemp008[0] = µe
							if πTemp005, πE = πg.ResolveGlobal(πF, ßquopri); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßdecodestring, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp005, µp); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
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
					if πE = πClass.SetItem(πF, ßtest_embedded_ws.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 161: @withpythonimplementation
					πF.SetLineno(161)
					πTemp002 = πF.MakeArgs(1)
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßtest_embedded_ws); πE != nil {
						continue
					}
					πTemp002[0] = πTemp009
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßwithpythonimplementation); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πClass.SetItem(πF, ßtest_embedded_ws.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 168: def test_encode_header(self):
					πF.SetLineno(168)
					πTemp006 = make([]πg.Param, 1)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("test_encode_header", "build/src/__python__/test/test_quopri.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µp *πg.Object = πg.UnboundLocal; _ = µp
						var µe *πg.Object = πg.UnboundLocal; _ = µe
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 πg.KWArgs
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßHSTRINGS, nil); πE != nil {
								continue
							}
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
									continue
								}
								µp = πTemp005
								µe = πTemp006
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 170: self.assertTrue(quopri.encodestring(p, header=True) == e)
							πF.SetLineno(170)
							πTemp007 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp008[0] = µp
							if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"header", πTemp005},
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßquopri); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßencodestring, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, πTemp008, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp005, µe); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
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
					if πE = πClass.SetItem(πF, ßtest_encode_header.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 167: @withpythonimplementation
					πF.SetLineno(167)
					πTemp002 = πF.MakeArgs(1)
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßtest_encode_header); πE != nil {
						continue
					}
					πTemp002[0] = πTemp010
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßwithpythonimplementation); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πClass.SetItem(πF, ßtest_encode_header.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 173: def test_decode_header(self):
					πF.SetLineno(173)
					πTemp006 = make([]πg.Param, 1)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("test_decode_header", "build/src/__python__/test/test_quopri.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µp *πg.Object = πg.UnboundLocal; _ = µp
						var µe *πg.Object = πg.UnboundLocal; _ = µe
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 πg.KWArgs
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßHSTRINGS, nil); πE != nil {
								continue
							}
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
									continue
								}
								µp = πTemp005
								µe = πTemp006
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 175: self.assertTrue(quopri.decodestring(e, header=True) == p)
							πF.SetLineno(175)
							πTemp007 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							πTemp008[0] = µe
							if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"header", πTemp005},
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßquopri); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßdecodestring, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, πTemp008, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp005, µp); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
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
					if πE = πClass.SetItem(πF, ßtest_decode_header.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 172: @withpythonimplementation
					πF.SetLineno(172)
					πTemp002 = πF.MakeArgs(1)
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßtest_decode_header); πE != nil {
						continue
					}
					πTemp002[0] = πTemp011
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßwithpythonimplementation); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp011.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πClass.SetItem(πF, ßtest_decode_header.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 178: def test_scriptencode(self):
					πF.SetLineno(178)
					πTemp006 = make([]πg.Param, 1)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("test_scriptencode", "build/src/__python__/test/test_quopri.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µp *πg.Object = πg.UnboundLocal; _ = µp
						var µe *πg.Object = πg.UnboundLocal; _ = µe
						var µprocess *πg.Object = πg.UnboundLocal; _ = µprocess
						var µcout *πg.Object = πg.UnboundLocal; _ = µcout
						var µcerr *πg.Object = πg.UnboundLocal; _ = µcerr
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 πg.KWArgs
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 179: (p, e) = self.STRINGS[-1]
							πF.SetLineno(179)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßSTRINGS, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
								continue
							}
							µp = πTemp001
							µe = πTemp003
							// line 180: process = subprocess.Popen([sys.executable, "-mquopri"],
							πF.SetLineno(180)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßexecutable, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							πTemp005[1] = πg.NewStr("-mquopri").ToObject()
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßPIPE, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßPIPE, nil); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"stdin", πTemp002},
								{"stdout", πTemp003},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßPopen, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µprocess = πTemp001
							// line 182: self.addCleanup(process.stdout.close)
							πF.SetLineno(182)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprocess, "process"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µprocess, ßstdout, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßclose, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßaddCleanup, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 183: cout, cerr = process.communicate(p)
							πF.SetLineno(183)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp004[0] = µp
							if πE = πg.CheckLocal(πF, µprocess, "process"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µprocess, ßcommunicate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
								continue
							}
							µcout = πTemp001
							µcerr = πTemp003
							// line 187: self.assertEqual(cout.splitlines(), e.splitlines())
							πF.SetLineno(187)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µcout, "cout"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcout, ßsplitlines, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µe, ßsplitlines, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_scriptencode.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 177: @unittest.expectedFailure
					πF.SetLineno(177)
					πTemp002 = πF.MakeArgs(1)
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßtest_scriptencode); πE != nil {
						continue
					}
					πTemp002[0] = πTemp012
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp013.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πClass.SetItem(πF, ßtest_scriptencode.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 190: def test_scriptdecode(self):
					πF.SetLineno(190)
					πTemp006 = make([]πg.Param, 1)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("test_scriptdecode", "build/src/__python__/test/test_quopri.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µp *πg.Object = πg.UnboundLocal; _ = µp
						var µe *πg.Object = πg.UnboundLocal; _ = µe
						var µprocess *πg.Object = πg.UnboundLocal; _ = µprocess
						var µcout *πg.Object = πg.UnboundLocal; _ = µcout
						var µcerr *πg.Object = πg.UnboundLocal; _ = µcerr
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 πg.KWArgs
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 191: (p, e) = self.STRINGS[-1]
							πF.SetLineno(191)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßSTRINGS, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
								continue
							}
							µp = πTemp001
							µe = πTemp003
							// line 192: process = subprocess.Popen([sys.executable, "-mquopri", "-d"],
							πF.SetLineno(192)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 3)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßexecutable, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							πTemp005[1] = πg.NewStr("-mquopri").ToObject()
							πTemp005[2] = πg.NewStr("-d").ToObject()
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßPIPE, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßPIPE, nil); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"stdin", πTemp002},
								{"stdout", πTemp003},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßPopen, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µprocess = πTemp001
							// line 194: self.addCleanup(process.stdout.close)
							πF.SetLineno(194)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprocess, "process"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µprocess, ßstdout, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßclose, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßaddCleanup, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 195: cout, cerr = process.communicate(e)
							πF.SetLineno(195)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							πTemp004[0] = µe
							if πE = πg.CheckLocal(πF, µprocess, "process"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µprocess, ßcommunicate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
								continue
							}
							µcout = πTemp001
							µcerr = πTemp003
							// line 196: self.assertEqual(cout.splitlines(), p.splitlines())
							πF.SetLineno(196)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µcout, "cout"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcout, ßsplitlines, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µp, ßsplitlines, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_scriptdecode.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 189: @unittest.expectedFailure
					πF.SetLineno(189)
					πTemp002 = πF.MakeArgs(1)
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßtest_scriptdecode); πE != nil {
						continue
					}
					πTemp002[0] = πTemp013
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp014.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πClass.SetItem(πF, ßtest_scriptdecode.ToObject(), πTemp013); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("QuopriTestCase").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßQuopriTestCase.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 198: def test_main():
			πF.SetLineno(198)
			πTemp003 = make([]πg.Param, 0)
			πTemp005 = πg.NewFunction(πg.NewCode("test_main", "build/src/__python__/test/test_quopri.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 199: test_support.run_unittest(QuopriTestCase)
					πF.SetLineno(199)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßQuopriTestCase); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßtest_main.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp006, πE = πg.Eq(πF, πTemp007, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp009, πE = πg.IsTrue(πF, πTemp006); πE != nil {
				continue
			}
			if πTemp009 {
				goto Label1
			}
			goto Label2
			// line 202: if __name__ == "__main__":
			πF.SetLineno(202)
		Label1:
			// line 203: test_main()
			πF.SetLineno(203)
			if πTemp006, πE = πg.ResolveGlobal(πF, ßtest_main); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("test.test_quopri", Code)
}
