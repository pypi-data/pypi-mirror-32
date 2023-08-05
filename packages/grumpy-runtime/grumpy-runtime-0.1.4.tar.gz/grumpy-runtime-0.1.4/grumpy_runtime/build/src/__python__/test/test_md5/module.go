package test_md5
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/test/test_md5.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ß0cc175b9c0f1b6a831c399e269772661 := πg.InternStr("0cc175b9c0f1b6a831c399e269772661")
		ß57edf4a22be3c955ac49da2e2107b67a := πg.InternStr("57edf4a22be3c955ac49da2e2107b67a")
		ß900150983cd24fb0d6963f7d28e17f72 := πg.InternStr("900150983cd24fb0d6963f7d28e17f72")
		ßABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789 := πg.InternStr("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
		ßDeprecationWarning := πg.InternStr("DeprecationWarning")
		ßMD5_Test := πg.InternStr("MD5_Test")
		ßTestCase := πg.InternStr("TestCase")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_md5 := πg.InternStr("_md5")
		ßa := πg.InternStr("a")
		ßabc := πg.InternStr("abc")
		ßabcdefghijklmnopqrstuvwxyz := πg.InternStr("abcdefghijklmnopqrstuvwxyz")
		ßassertEqual := πg.InternStr("assertEqual")
		ßb := πg.InternStr("b")
		ßc := πg.InternStr("c")
		ßc3fcd3d76192e4007dfb496cca67e13b := πg.InternStr("c3fcd3d76192e4007dfb496cca67e13b")
		ßd174ab98d277d9f5a5611c2c9f419d9f := πg.InternStr("d174ab98d277d9f5a5611c2c9f419d9f")
		ßd41d8cd98f00b204e9800998ecf8427e := πg.InternStr("d41d8cd98f00b204e9800998ecf8427e")
		ßdigest := πg.InternStr("digest")
		ßf96b697d7cb7938d525a2f31aaf161d0 := πg.InternStr("f96b697d7cb7938d525a2f31aaf161d0")
		ßfilterwarnings := πg.InternStr("filterwarnings")
		ßhexdigest := πg.InternStr("hexdigest")
		ßhexdigits := πg.InternStr("hexdigits")
		ßhexstr := πg.InternStr("hexstr")
		ßignore := πg.InternStr("ignore")
		ßmd5 := πg.InternStr("md5")
		ßmd5test := πg.InternStr("md5test")
		ßord := πg.InternStr("ord")
		ßrun_unittest := πg.InternStr("run_unittest")
		ßtest_basics := πg.InternStr("test_basics")
		ßtest_hexdigest := πg.InternStr("test_hexdigest")
		ßtest_large_update := πg.InternStr("test_large_update")
		ßtest_main := πg.InternStr("test_main")
		ßtest_support := πg.InternStr("test_support")
		ßunittest := πg.InternStr("unittest")
		ßupdate := πg.InternStr("update")
		ßwarnings := πg.InternStr("warnings")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 []πg.Param
		_ = πTemp004
		var πTemp005 *πg.Dict
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
			// line 2: import warnings
			πF.SetLineno(2)
			if πTemp002, πE = πg.ImportModule(πF, "warnings"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßwarnings.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 3: warnings.filterwarnings("ignore", "the md5 module is deprecated.*",
			πF.SetLineno(3)
			πTemp002 = πF.MakeArgs(3)
			πTemp002[0] = ßignore.ToObject()
			πTemp002[1] = πg.NewStr("the md5 module is deprecated.*").ToObject()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßDeprecationWarning); πE != nil {
				continue
			}
			πTemp002[2] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßwarnings); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßfilterwarnings, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 6: import unittest
			πF.SetLineno(6)
			if πTemp002, πE = πg.ImportModule(πF, "unittest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 8: import md5 as _md5
			πF.SetLineno(8)
			if πTemp002, πE = πg.ImportModule(πF, "md5"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_md5.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 9: md5 = _md5.md5
			πF.SetLineno(9)
			if πTemp001, πE = πg.ResolveGlobal(πF, ß_md5); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßmd5, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmd5.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 10: from test import test_support
			πF.SetLineno(10)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßtest_support.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 12: def hexstr(s):
			πF.SetLineno(12)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "s", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("hexstr", "build/src/__python__/test/test_md5.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µstring *πg.Object = πg.UnboundLocal; _ = µstring
				var µh *πg.Object = πg.UnboundLocal; _ = µh
				var µr *πg.Object = πg.UnboundLocal; _ = µr
				var µc *πg.Object = πg.UnboundLocal; _ = µc
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 bool
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 13: import string
					πF.SetLineno(13)
					if πTemp002, πE = πg.ImportModule(πF, "string"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					µstring = πTemp001
					// line 14: h = string.hexdigits
					πF.SetLineno(14)
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µstring, ßhexdigits, nil); πE != nil {
						continue
					}
					µh = πTemp001
					// line 15: r = ''
					πF.SetLineno(15)
					µr = ß.ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µs); πE != nil {
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
					if πTemp005, πE = πg.Next(πF, πTemp001); πE != nil {
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
						µc = πTemp005
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 17: i = ord(c)
					πF.SetLineno(17)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp002[0] = µc
					if πTemp005, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µi = πTemp006
					// line 18: r = r + h[(i >> 4) & 0xF] + h[i & 0xF]
					πF.SetLineno(18)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.RShift(πF, µi, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.And(πF, πTemp009, πg.NewInt(15).ToObject()); πE != nil {
						continue
					}
					πTemp007 = πTemp008
					if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µh, πTemp007); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Add(πF, µr, πTemp008); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.And(πF, µi, πg.NewInt(15).ToObject()); πE != nil {
						continue
					}
					πTemp007 = πTemp008
					if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µh, πTemp007); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Add(πF, πTemp006, πTemp008); πE != nil {
						continue
					}
					µr = πTemp005
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 19: return r
					πF.SetLineno(19)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πR = µr
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßhexstr.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 21: class MD5_Test(unittest.TestCase):
			πF.SetLineno(21)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßTestCase, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp005 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("MD5_Test", "build/src/__python__/test/test_md5.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
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
					// line 23: def md5test(self, s, expected):
					πF.SetLineno(23)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "s", Def: nil}
					πTemp002[2] = πg.Param{Name: "expected", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("md5test", "build/src/__python__/test/test_md5.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µs *πg.Object = πArgs[1]; _ = µs
						var µexpected *πg.Object = πArgs[2]; _ = µexpected
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
							// line 24: self.assertEqual(hexstr(md5(s).digest()), expected)
							πF.SetLineno(24)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp003[0] = µs
							if πTemp004, πE = πg.ResolveGlobal(πF, ßmd5); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßdigest, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßhexstr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µexpected, "expected"); πE != nil {
								continue
							}
							πTemp001[1] = µexpected
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
							// line 25: self.assertEqual(md5(s).hexdigest(), expected)
							πF.SetLineno(25)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp002[0] = µs
							if πTemp004, πE = πg.ResolveGlobal(πF, ßmd5); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßhexdigest, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µexpected, "expected"); πE != nil {
								continue
							}
							πTemp001[1] = µexpected
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
					if πE = πClass.SetItem(πF, ßmd5test.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 27: def test_basics(self):
					πF.SetLineno(27)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_basics", "build/src/__python__/test/test_md5.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µeq *πg.Object = πg.UnboundLocal; _ = µeq
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 28: eq = self.md5test
							πF.SetLineno(28)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmd5test, nil); πE != nil {
								continue
							}
							µeq = πTemp001
							// line 29: eq('', 'd41d8cd98f00b204e9800998ecf8427e')
							πF.SetLineno(29)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = ß.ToObject()
							πTemp002[1] = ßd41d8cd98f00b204e9800998ecf8427e.ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 30: eq('a', '0cc175b9c0f1b6a831c399e269772661')
							πF.SetLineno(30)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = ßa.ToObject()
							πTemp002[1] = ß0cc175b9c0f1b6a831c399e269772661.ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 31: eq('abc', '900150983cd24fb0d6963f7d28e17f72')
							πF.SetLineno(31)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = ßabc.ToObject()
							πTemp002[1] = ß900150983cd24fb0d6963f7d28e17f72.ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 32: eq('message digest', 'f96b697d7cb7938d525a2f31aaf161d0')
							πF.SetLineno(32)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewStr("message digest").ToObject()
							πTemp002[1] = ßf96b697d7cb7938d525a2f31aaf161d0.ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 33: eq('abcdefghijklmnopqrstuvwxyz', 'c3fcd3d76192e4007dfb496cca67e13b')
							πF.SetLineno(33)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = ßabcdefghijklmnopqrstuvwxyz.ToObject()
							πTemp002[1] = ßc3fcd3d76192e4007dfb496cca67e13b.ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 34: eq('ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789',
							πF.SetLineno(34)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = ßABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789.ToObject()
							πTemp002[1] = ßd174ab98d277d9f5a5611c2c9f419d9f.ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 36: eq('12345678901234567890123456789012345678901234567890123456789012345678901234567890',
							πF.SetLineno(36)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewStr("12345678901234567890123456789012345678901234567890123456789012345678901234567890").ToObject()
							πTemp002[1] = ß57edf4a22be3c955ac49da2e2107b67a.ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_basics.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 39: def test_hexdigest(self):
					πF.SetLineno(39)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_hexdigest", "build/src/__python__/test/test_md5.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µm *πg.Object = πg.UnboundLocal; _ = µm
						var µh *πg.Object = πg.UnboundLocal; _ = µh
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
							// line 41: m = md5('testing the hexdigest method')
							πF.SetLineno(41)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("testing the hexdigest method").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmd5); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µm = πTemp003
							// line 42: h = m.hexdigest()
							πF.SetLineno(42)
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µm, ßhexdigest, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µh = πTemp003
							// line 43: self.assertEqual(hexstr(m.digest()), h)
							πF.SetLineno(43)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µm, ßdigest, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhexstr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							πTemp001[1] = µh
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
					if πE = πClass.SetItem(πF, ßtest_hexdigest.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 45: def test_large_update(self):
					πF.SetLineno(45)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_large_update", "build/src/__python__/test/test_md5.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µaas *πg.Object = πg.UnboundLocal; _ = µaas
						var µbees *πg.Object = πg.UnboundLocal; _ = µbees
						var µcees *πg.Object = πg.UnboundLocal; _ = µcees
						var µm1 *πg.Object = πg.UnboundLocal; _ = µm1
						var µm2 *πg.Object = πg.UnboundLocal; _ = µm2
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
							// line 46: aas = 'a' * 64
							πF.SetLineno(46)
							if πTemp001, πE = πg.Mul(πF, ßa.ToObject(), πg.NewInt(64).ToObject()); πE != nil {
								continue
							}
							µaas = πTemp001
							// line 47: bees = 'b' * 64
							πF.SetLineno(47)
							if πTemp001, πE = πg.Mul(πF, ßb.ToObject(), πg.NewInt(64).ToObject()); πE != nil {
								continue
							}
							µbees = πTemp001
							// line 48: cees = 'c' * 64
							πF.SetLineno(48)
							if πTemp001, πE = πg.Mul(πF, ßc.ToObject(), πg.NewInt(64).ToObject()); πE != nil {
								continue
							}
							µcees = πTemp001
							// line 50: m1 = md5()
							πF.SetLineno(50)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmd5); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µm1 = πTemp002
							// line 51: m1.update(aas)
							πF.SetLineno(51)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µaas, "aas"); πE != nil {
								continue
							}
							πTemp003[0] = µaas
							if πE = πg.CheckLocal(πF, µm1, "m1"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µm1, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 52: m1.update(bees)
							πF.SetLineno(52)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbees, "bees"); πE != nil {
								continue
							}
							πTemp003[0] = µbees
							if πE = πg.CheckLocal(πF, µm1, "m1"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µm1, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 53: m1.update(cees)
							πF.SetLineno(53)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcees, "cees"); πE != nil {
								continue
							}
							πTemp003[0] = µcees
							if πE = πg.CheckLocal(πF, µm1, "m1"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µm1, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 55: m2 = md5()
							πF.SetLineno(55)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmd5); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µm2 = πTemp002
							// line 56: m2.update(aas + bees + cees)
							πF.SetLineno(56)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µaas, "aas"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbees, "bees"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µaas, µbees); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcees, "cees"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, µcees); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µm2, "m2"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µm2, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 57: self.assertEqual(m1.digest(), m2.digest())
							πF.SetLineno(57)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µm1, "m1"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µm1, ßdigest, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µm2, "m2"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µm2, ßdigest, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_large_update.ToObject(), πTemp005); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("MD5_Test").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMD5_Test.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 59: def test_main():
			πF.SetLineno(59)
			πTemp004 = make([]πg.Param, 0)
			πTemp003 = πg.NewFunction(πg.NewCode("test_main", "build/src/__python__/test/test_md5.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 60: test_support.run_unittest(MD5_Test)
					πF.SetLineno(60)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßMD5_Test); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßtest_main.ToObject(), πTemp003); πE != nil {
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
			// line 62: if __name__ == '__main__':
			πF.SetLineno(62)
		Label1:
			// line 63: test_main()
			πF.SetLineno(63)
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
	πg.RegisterModule("test.test_md5", Code)
}
