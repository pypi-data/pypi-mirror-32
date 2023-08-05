package test_mimetools
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/test/test_mimetools.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ß7bit := πg.InternStr("7bit")
		ß8bit := πg.InternStr("8bit")
		ßMessage := πg.InternStr("Message")
		ßMimeToolsTest := πg.InternStr("MimeToolsTest")
		ßNone := πg.InternStr("None")
		ßStringIO := πg.InternStr("StringIO")
		ßTestCase := πg.InternStr("TestCase")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßadd := πg.InternStr("add")
		ßascii_letters := πg.InternStr("ascii_letters")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertNotIn := πg.InternStr("assertNotIn")
		ßbase64 := πg.InternStr("base64")
		ßcharset := πg.InternStr("charset")
		ßchoose_boundary := πg.InternStr("choose_boundary")
		ßdecode := πg.InternStr("decode")
		ßdigits := πg.InternStr("digits")
		ßencode := πg.InternStr("encode")
		ßexpectedFailure := πg.InternStr("expectedFailure")
		ßflowed := πg.InternStr("flowed")
		ßformat := πg.InternStr("format")
		ßgetencoding := πg.InternStr("getencoding")
		ßgetmaintype := πg.InternStr("getmaintype")
		ßgetparam := πg.InternStr("getparam")
		ßgetparamnames := πg.InternStr("getparamnames")
		ßgetplist := πg.InternStr("getplist")
		ßgetsubtype := πg.InternStr("getsubtype")
		ßgettype := πg.InternStr("gettype")
		ßgetvalue := πg.InternStr("getvalue")
		ßmimetools := πg.InternStr("mimetools")
		ßmsgtext1 := πg.InternStr("msgtext1")
		ßplain := πg.InternStr("plain")
		ßrun_unittest := πg.InternStr("run_unittest")
		ßset := πg.InternStr("set")
		ßspam := πg.InternStr("spam")
		ßstring := πg.InternStr("string")
		ßtest_boundary := πg.InternStr("test_boundary")
		ßtest_decodeencode := πg.InternStr("test_decodeencode")
		ßtest_main := πg.InternStr("test_main")
		ßtest_message := πg.InternStr("test_message")
		ßtest_support := πg.InternStr("test_support")
		ßtext := πg.InternStr("text")
		ßunittest := πg.InternStr("unittest")
		ßuue := πg.InternStr("uue")
		ßuuencode := πg.InternStr("uuencode")
		ßxrange := πg.InternStr("xrange")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 []*πg.Object
		_ = πTemp003
		var πTemp004 *πg.Object
		_ = πTemp004
		var πTemp005 *πg.Dict
		_ = πTemp005
		var πTemp006 *πg.Object
		_ = πTemp006
		var πTemp007 *πg.Object
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
			// line 1: import unittest
			πF.SetLineno(1)
			if πTemp002, πE = πg.ImportModule(πF, "unittest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 2: from test import test_support
			πF.SetLineno(2)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßtest_support.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 4: import string
			πF.SetLineno(4)
			if πTemp002, πE = πg.ImportModule(πF, "string"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßstring.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 5: import StringIO
			πF.SetLineno(5)
			if πTemp002, πE = πg.ImportModule(πF, "StringIO"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßStringIO.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 8: import mimetools
			πF.SetLineno(8)
			if πTemp002, πE = πg.ImportModule(πF, "mimetools"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßmimetools.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 10: msgtext1 = mimetools.Message(StringIO.StringIO(
			πF.SetLineno(10)
			πTemp002 = πF.MakeArgs(1)
			πTemp003 = πF.MakeArgs(1)
			πTemp003[0] = πg.NewStr("Content-Type: text/plain; charset=iso-8859-1; format=flowed\nContent-Transfer-Encoding: 8bit\n\nFoo!\n").ToObject()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßStringIO, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp003)
			πTemp002[0] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßmimetools); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßMessage, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßmsgtext1.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 17: class MimeToolsTest(unittest.TestCase):
			πF.SetLineno(17)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp006, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßTestCase, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp005 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("MimeToolsTest", "build/src/__python__/test/test_mimetools.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
				_ = πClass
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 19: def test_decodeencode(self):
					πF.SetLineno(19)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_decodeencode", "build/src/__python__/test/test_mimetools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstart *πg.Object = πg.UnboundLocal; _ = µstart
						var µenc *πg.Object = πg.UnboundLocal; _ = µenc
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µo *πg.Object = πg.UnboundLocal; _ = µo
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 bool
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
							// line 20: start = string.ascii_letters + "=" + string.digits + "\n"
							πF.SetLineno(20)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßascii_letters, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp005, πg.NewStr("=").ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßdigits, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							µstart = πTemp001
							πTemp006 = make([]*πg.Object, 8)
							πTemp006[0] = ß7bit.ToObject()
							πTemp006[1] = ß8bit.ToObject()
							πTemp006[2] = ßbase64.ToObject()
							πTemp006[3] = πg.NewStr("quoted-printable").ToObject()
							πTemp006[4] = ßuuencode.ToObject()
							πTemp006[5] = πg.NewStr("x-uuencode").ToObject()
							πTemp006[6] = ßuue.ToObject()
							πTemp006[7] = πg.NewStr("x-uue").ToObject()
							πTemp002 = πg.NewList(πTemp006...).ToObject()
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µenc = πTemp002
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 23: i = StringIO.StringIO(start)
							πF.SetLineno(23)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							πTemp006[0] = µstart
							if πTemp002, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µi = πTemp002
							// line 24: o = StringIO.StringIO()
							πF.SetLineno(24)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µo = πTemp002
							// line 25: mimetools.encode(i, o, enc)
							πF.SetLineno(25)
							πTemp006 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp006[0] = µi
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp006[1] = µo
							if πE = πg.CheckLocal(πF, µenc, "enc"); πE != nil {
								continue
							}
							πTemp006[2] = µenc
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmimetools); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßencode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 26: i = StringIO.StringIO(o.getvalue())
							πF.SetLineno(26)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µo, ßgetvalue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µi = πTemp002
							// line 27: o = StringIO.StringIO()
							πF.SetLineno(27)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µo = πTemp002
							// line 28: mimetools.decode(i, o, enc)
							πF.SetLineno(28)
							πTemp006 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp006[0] = µi
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp006[1] = µo
							if πE = πg.CheckLocal(πF, µenc, "enc"); πE != nil {
								continue
							}
							πTemp006[2] = µenc
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmimetools); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdecode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 29: self.assertEqual(o.getvalue(), start)
							πF.SetLineno(29)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µo, ßgetvalue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp003
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							πTemp006[1] = µstart
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
					if πE = πClass.SetItem(πF, ßtest_decodeencode.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 32: def test_boundary(self):
					πF.SetLineno(32)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_boundary", "build/src/__python__/test/test_mimetools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µnb *πg.Object = πg.UnboundLocal; _ = µnb
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
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
							// line 33: s = set([""])
							πF.SetLineno(33)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 1)
							πTemp002[0] = ß.ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp004
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(100).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp006 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
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
								πTemp007 = !isStop
							} else {
								πTemp007 = true
								µi = πTemp004
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 35: nb = mimetools.choose_boundary()
							πF.SetLineno(35)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßmimetools); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßchoose_boundary, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							µnb = πTemp004
							// line 36: self.assertNotIn(nb, s)
							πF.SetLineno(36)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnb, "nb"); πE != nil {
								continue
							}
							πTemp001[0] = µnb
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertNotIn, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 37: s.add(nb)
							πF.SetLineno(37)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnb, "nb"); πE != nil {
								continue
							}
							πTemp001[0] = µnb
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µs, ßadd, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_boundary.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 31: @unittest.expectedFailure
					πF.SetLineno(31)
					πTemp004 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßtest_boundary); πE != nil {
						continue
					}
					πTemp004[0] = πTemp005
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_boundary.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 39: def test_message(self):
					πF.SetLineno(39)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_message", "build/src/__python__/test/test_mimetools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmsg *πg.Object = πg.UnboundLocal; _ = µmsg
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
							// line 40: msg = mimetools.Message(StringIO.StringIO(msgtext1))
							πF.SetLineno(40)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßmsgtext1); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßStringIO, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßmimetools); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßMessage, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmsg = πTemp003
							// line 41: self.assertEqual(msg.gettype(), "text/plain")
							πF.SetLineno(41)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmsg, ßgettype, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewStr("text/plain").ToObject()
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
							// line 42: self.assertEqual(msg.getmaintype(), "text")
							πF.SetLineno(42)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmsg, ßgetmaintype, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp001[1] = ßtext.ToObject()
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
							// line 43: self.assertEqual(msg.getsubtype(), "plain")
							πF.SetLineno(43)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmsg, ßgetsubtype, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp001[1] = ßplain.ToObject()
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
							// line 44: self.assertEqual(msg.getplist(), ["charset=iso-8859-1", "format=flowed"])
							πF.SetLineno(44)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmsg, ßgetplist, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewStr("charset=iso-8859-1").ToObject()
							πTemp002[1] = πg.NewStr("format=flowed").ToObject()
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
							// line 45: self.assertEqual(msg.getparamnames(), ["charset", "format"])
							πF.SetLineno(45)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmsg, ßgetparamnames, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = ßcharset.ToObject()
							πTemp002[1] = ßformat.ToObject()
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
							// line 46: self.assertEqual(msg.getparam("charset"), "iso-8859-1")
							πF.SetLineno(46)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßcharset.ToObject()
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmsg, ßgetparam, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewStr("iso-8859-1").ToObject()
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
							// line 47: self.assertEqual(msg.getparam("format"), "flowed")
							πF.SetLineno(47)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßformat.ToObject()
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmsg, ßgetparam, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = ßflowed.ToObject()
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
							// line 48: self.assertEqual(msg.getparam("spam"), None)
							πF.SetLineno(48)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßspam.ToObject()
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmsg, ßgetparam, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
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
							// line 49: self.assertEqual(msg.getencoding(), "8bit")
							πF.SetLineno(49)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmsg, ßgetencoding, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp001[1] = ß8bit.ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_message.ToObject(), πTemp005); πE != nil {
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
			if πTemp006, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("MimeToolsTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMimeToolsTest.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 51: def test_main():
			πF.SetLineno(51)
			πTemp008 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("test_main", "build/src/__python__/test/test_mimetools.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 52: test_support.run_unittest(MimeToolsTest)
					πF.SetLineno(52)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßMimeToolsTest); πE != nil {
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
			if πTemp006, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp004, πE = πg.Eq(πF, πTemp006, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp009, πE = πg.IsTrue(πF, πTemp004); πE != nil {
				continue
			}
			if πTemp009 {
				goto Label1
			}
			goto Label2
			// line 54: if __name__=="__main__":
			πF.SetLineno(54)
		Label1:
			// line 55: test_main()
			πF.SetLineno(55)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßtest_main); πE != nil {
				continue
			}
			if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("test.test_mimetools", Code)
}
