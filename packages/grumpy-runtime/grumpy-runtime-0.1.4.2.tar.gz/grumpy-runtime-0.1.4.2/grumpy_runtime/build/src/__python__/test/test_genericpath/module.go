package test_genericpath
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/test/test_genericpath.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßAttributeError := πg.InternStr("AttributeError")
		ßCommonTest := πg.InternStr("CommonTest")
		ßEnvironmentVarGuard := πg.InternStr("EnvironmentVarGuard")
		ßFS_NONASCII := πg.InternStr("FS_NONASCII")
		ßFalse := πg.InternStr("False")
		ßGenericTest := πg.InternStr("GenericTest")
		ßNone := πg.InternStr("None")
		ßOSError := πg.InternStr("OSError")
		ßTESTFN := πg.InternStr("TESTFN")
		ßTESTFN_ENCODING := πg.InternStr("TESTFN_ENCODING")
		ßTestCase := πg.InternStr("TestCase")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßUnicodeEncodeError := πg.InternStr("UnicodeEncodeError")
		ßXY := πg.InternStr("XY")
		ßXb := πg.InternStr("Xb")
		ßXbcd := πg.InternStr("Xbcd")
		ß__enter__ := πg.InternStr("__enter__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßaX := πg.InternStr("aX")
		ßaXc := πg.InternStr("aXc")
		ßab := πg.InternStr("ab")
		ßabc := πg.InternStr("abc")
		ßabcX := πg.InternStr("abcX")
		ßabcd := πg.InternStr("abcd")
		ßabd := πg.InternStr("abd")
		ßabspath := πg.InternStr("abspath")
		ßaltsep := πg.InternStr("altsep")
		ßascii := πg.InternStr("ascii")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertIn := πg.InternStr("assertIn")
		ßassertIs := πg.InternStr("assertIs")
		ßassertIsInstance := πg.InternStr("assertIsInstance")
		ßassertLessEqual := πg.InternStr("assertLessEqual")
		ßassertNotEqual := πg.InternStr("assertNotEqual")
		ßassertRaises := πg.InternStr("assertRaises")
		ßassertTrue := πg.InternStr("assertTrue")
		ßattributes := πg.InternStr("attributes")
		ßbar := πg.InternStr("bar")
		ßbarbar := πg.InternStr("barbar")
		ßbasename := πg.InternStr("basename")
		ßbaz1 := πg.InternStr("baz1")
		ßbaz2 := πg.InternStr("baz2")
		ßclear := πg.InternStr("clear")
		ßclose := πg.InternStr("close")
		ßclosed := πg.InternStr("closed")
		ßcommon_attributes := πg.InternStr("common_attributes")
		ßcommonprefix := πg.InternStr("commonprefix")
		ßcurdir := πg.InternStr("curdir")
		ßdarwin := πg.InternStr("darwin")
		ßdefpath := πg.InternStr("defpath")
		ßdevnull := πg.InternStr("devnull")
		ßdirname := πg.InternStr("dirname")
		ßencode := πg.InternStr("encode")
		ßexists := πg.InternStr("exists")
		ßexpanduser := πg.InternStr("expanduser")
		ßexpandvars := πg.InternStr("expandvars")
		ßextsep := πg.InternStr("extsep")
		ßfail := πg.InternStr("fail")
		ßfoo := πg.InternStr("foo")
		ßfoobar := πg.InternStr("foobar")
		ßformat := πg.InternStr("format")
		ßgenericpath := πg.InternStr("genericpath")
		ßgetatime := πg.InternStr("getatime")
		ßgetattr := πg.InternStr("getattr")
		ßgetctime := πg.InternStr("getctime")
		ßgetfilesystemencoding := πg.InternStr("getfilesystemencoding")
		ßgetmtime := πg.InternStr("getmtime")
		ßgetsize := πg.InternStr("getsize")
		ßgrumpy := πg.InternStr("grumpy")
		ßham := πg.InternStr("ham")
		ßisabs := πg.InternStr("isabs")
		ßisdir := πg.InternStr("isdir")
		ßisfile := πg.InternStr("isfile")
		ßislink := πg.InternStr("islink")
		ßismount := πg.InternStr("ismount")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßlexists := πg.InternStr("lexists")
		ßmacpath := πg.InternStr("macpath")
		ßmkdir := πg.InternStr("mkdir")
		ßnormcase := πg.InternStr("normcase")
		ßnormpath := πg.InternStr("normpath")
		ßopen := πg.InternStr("open")
		ßos := πg.InternStr("os")
		ßpardir := πg.InternStr("pardir")
		ßpathmodule := πg.InternStr("pathmodule")
		ßpathsep := πg.InternStr("pathsep")
		ßplatform := πg.InternStr("platform")
		ßrb := πg.InternStr("rb")
		ßread := πg.InternStr("read")
		ßrealpath := πg.InternStr("realpath")
		ßremove := πg.InternStr("remove")
		ßrequires_unicode := πg.InternStr("requires_unicode")
		ßrmdir := πg.InternStr("rmdir")
		ßrun_unittest := πg.InternStr("run_unittest")
		ßsafe_rmdir := πg.InternStr("safe_rmdir")
		ßsep := πg.InternStr("sep")
		ßskip := πg.InternStr("skip")
		ßskipIf := πg.InternStr("skipIf")
		ßskipTest := πg.InternStr("skipTest")
		ßskipUnless := πg.InternStr("skipUnless")
		ßspam := πg.InternStr("spam")
		ßsplit := πg.InternStr("split")
		ßsplitdrive := πg.InternStr("splitdrive")
		ßsplitext := πg.InternStr("splitext")
		ßstartswith := πg.InternStr("startswith")
		ßstr := πg.InternStr("str")
		ßsys := πg.InternStr("sys")
		ßtemp_cwd := πg.InternStr("temp_cwd")
		ßtest_abspath := πg.InternStr("test_abspath")
		ßtest_abspath_issue3426 := πg.InternStr("test_abspath_issue3426")
		ßtest_commonprefix := πg.InternStr("test_commonprefix")
		ßtest_exists := πg.InternStr("test_exists")
		ßtest_expandvars := πg.InternStr("test_expandvars")
		ßtest_expandvars_nonascii := πg.InternStr("test_expandvars_nonascii")
		ßtest_getsize := πg.InternStr("test_getsize")
		ßtest_isdir := πg.InternStr("test_isdir")
		ßtest_isfile := πg.InternStr("test_isfile")
		ßtest_main := πg.InternStr("test_main")
		ßtest_no_argument := πg.InternStr("test_no_argument")
		ßtest_nonascii_abspath := πg.InternStr("test_nonascii_abspath")
		ßtest_normcase := πg.InternStr("test_normcase")
		ßtest_normpath_issue5827 := πg.InternStr("test_normpath_issue5827")
		ßtest_realpath := πg.InternStr("test_realpath")
		ßtest_splitdrive := πg.InternStr("test_splitdrive")
		ßtest_support := πg.InternStr("test_support")
		ßtest_time := πg.InternStr("test_time")
		ßunicode := πg.InternStr("unicode")
		ßunittest := πg.InternStr("unittest")
		ßunlink := πg.InternStr("unlink")
		ßwb := πg.InternStr("wb")
		ßwrite := πg.InternStr("write")
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
			// line 1: """
			πF.SetLineno(1)
			// line 5: import unittest
			πF.SetLineno(5)
			if πTemp002, πE = πg.ImportModule(πF, "unittest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 6: from test import test_support
			πF.SetLineno(6)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßtest_support.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 7: import os
			πF.SetLineno(7)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 8: import genericpath
			πF.SetLineno(8)
			if πTemp002, πE = πg.ImportModule(πF, "genericpath"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßgenericpath.ToObject(), πTemp001); πE != nil {
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
			// line 12: def safe_rmdir(dirname):
			πF.SetLineno(12)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "dirname", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("safe_rmdir", "build/src/__python__/test/test_genericpath.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µdirname *πg.Object = πArgs[0]; _ = µdirname
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 13: try:
					πF.SetLineno(13)
					πF.PushCheckpoint(2)
					// line 14: os.rmdir(dirname)
					πF.SetLineno(14)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdirname, "dirname"); πE != nil {
						continue
					}
					πTemp001[0] = µdirname
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrmdir, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πTemp002, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 15: except OSError:
					πF.SetLineno(15)
				Label3:
					// line 16: pass
					πF.SetLineno(16)
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
			if πE = πF.Globals().SetItem(πF, ßsafe_rmdir.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 19: class GenericTest(unittest.TestCase):
			πF.SetLineno(19)
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
			_, πE = πg.NewCode("GenericTest", "build/src/__python__/test/test_genericpath.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 []πg.Param
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
					// line 21: pathmodule = genericpath
					πF.SetLineno(21)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßgenericpath); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßpathmodule.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 22: common_attributes = ['commonprefix', 'getsize', 'getatime', 'getctime',
					πF.SetLineno(22)
					πTemp002 = make([]*πg.Object, 8)
					πTemp002[0] = ßcommonprefix.ToObject()
					πTemp002[1] = ßgetsize.ToObject()
					πTemp002[2] = ßgetatime.ToObject()
					πTemp002[3] = ßgetctime.ToObject()
					πTemp002[4] = ßgetmtime.ToObject()
					πTemp002[5] = ßexists.ToObject()
					πTemp002[6] = ßisdir.ToObject()
					πTemp002[7] = ßisfile.ToObject()
					πTemp001 = πg.NewList(πTemp002...).ToObject()
					if πE = πClass.SetItem(πF, ßcommon_attributes.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 24: attributes = []
					πF.SetLineno(24)
					πTemp002 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp002...).ToObject()
					if πE = πClass.SetItem(πF, ßattributes.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 26: def test_no_argument(self):
					πF.SetLineno(26)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_no_argument", "build/src/__python__/test/test_genericpath.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µattr *πg.Object = πg.UnboundLocal; _ = µattr
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πTemp011 *πg.BaseException
						_ = πTemp011
						var πTemp012 *πg.Traceback
						_ = πTemp012
						var πTemp013 *πg.Type
						_ = πTemp013
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcommon_attributes, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßattributes, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µattr = πTemp002
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 28: with self.assertRaises(TypeError):
							πF.SetLineno(28)
							πTemp007 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp002, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp004.Call(πF, πg.Args{πTemp003}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							// line 29: getattr(self.pathmodule, attr)()
							πF.SetLineno(29)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp008
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp007[1] = µattr
							if πTemp008, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp008, πE = πTemp009.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp007 = πF.MakeArgs(1)
							πTemp010 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp008, ß__name__, nil); πE != nil {
								continue
							}
							πTemp010[0] = πTemp009
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp010[1] = µattr
							if πTemp008, πE = πg.GetAttr(πF, πg.NewStr("{}.{}() did not raise a TypeError").ToObject(), ßformat, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp007[0] = πTemp009
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 30: raise self.fail("{}.{}() did not raise a TypeError"
							πF.SetLineno(30)
							πE = πF.Raise(πTemp009, nil, nil)
							continue
							πF.PopCheckpoint()
						Label4:
							πTemp011, πTemp012 = nil, nil
							if πE != nil {
								πTemp011, πTemp012 = πF.ExcInfo()
							}
							if πTemp011 != nil {
								πTemp013 = πTemp011.Type()
								if πTemp008, πE = πTemp002.Call(πF, πg.Args{πTemp003, πTemp013.ToObject(), πTemp011.ToObject(), πTemp012.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp008, πE = πTemp002.Call(πF, πg.Args{πTemp003, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							if πTemp011 != nil && πTemp006 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßtest_no_argument.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 33: def test_commonprefix(self):
					πF.SetLineno(33)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_commonprefix", "build/src/__python__/test/test_genericpath.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µcommonprefix *πg.Object = πg.UnboundLocal; _ = µcommonprefix
						var µtestlist *πg.Object = πg.UnboundLocal; _ = µtestlist
						var µs1 *πg.Object = πg.UnboundLocal; _ = µs1
						var µs2 *πg.Object = πg.UnboundLocal; _ = µs2
						var µp *πg.Object = πg.UnboundLocal; _ = µp
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							case 4: goto Label4
							case 5: goto Label5
							default: panic("unexpected function state")
							}
							// line 34: commonprefix = self.pathmodule.commonprefix
							πF.SetLineno(34)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßcommonprefix, nil); πE != nil {
								continue
							}
							µcommonprefix = πTemp002
							// line 35: self.assertEqual(
							πF.SetLineno(35)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µcommonprefix, "commonprefix"); πE != nil {
								continue
							}
							if πTemp001, πE = µcommonprefix.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp001
							πTemp003[1] = ß.ToObject()
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
							// line 39: self.assertEqual(
							πF.SetLineno(39)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewStr("/home/swenson/spam").ToObject()
							πTemp005[1] = πg.NewStr("/home/swen/spam").ToObject()
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µcommonprefix, "commonprefix"); πE != nil {
								continue
							}
							if πTemp001, πE = µcommonprefix.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp001
							πTemp003[1] = πg.NewStr("/home/swen").ToObject()
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
							// line 43: self.assertEqual(
							πF.SetLineno(43)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewStr("/home/swen/spam").ToObject()
							πTemp005[1] = πg.NewStr("/home/swen/eggs").ToObject()
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µcommonprefix, "commonprefix"); πE != nil {
								continue
							}
							if πTemp001, πE = µcommonprefix.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp001
							πTemp003[1] = πg.NewStr("/home/swen/").ToObject()
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
							// line 47: self.assertEqual(
							πF.SetLineno(47)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewStr("/home/swen/spam").ToObject()
							πTemp005[1] = πg.NewStr("/home/swen/spam").ToObject()
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µcommonprefix, "commonprefix"); πE != nil {
								continue
							}
							if πTemp001, πE = µcommonprefix.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp001
							πTemp003[1] = πg.NewStr("/home/swen/spam").ToObject()
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
							// line 51: self.assertEqual(
							πF.SetLineno(51)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewStr("home:swenson:spam").ToObject()
							πTemp005[1] = πg.NewStr("home:swen:spam").ToObject()
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µcommonprefix, "commonprefix"); πE != nil {
								continue
							}
							if πTemp001, πE = µcommonprefix.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp001
							πTemp003[1] = πg.NewStr("home:swen").ToObject()
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
							// line 55: self.assertEqual(
							πF.SetLineno(55)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewStr(":home:swen:spam").ToObject()
							πTemp005[1] = πg.NewStr(":home:swen:eggs").ToObject()
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µcommonprefix, "commonprefix"); πE != nil {
								continue
							}
							if πTemp001, πE = µcommonprefix.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp001
							πTemp003[1] = πg.NewStr(":home:swen:").ToObject()
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
							// line 59: self.assertEqual(
							πF.SetLineno(59)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewStr(":home:swen:spam").ToObject()
							πTemp005[1] = πg.NewStr(":home:swen:spam").ToObject()
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µcommonprefix, "commonprefix"); πE != nil {
								continue
							}
							if πTemp001, πE = µcommonprefix.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp001
							πTemp003[1] = πg.NewStr(":home:swen:spam").ToObject()
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
							// line 64: testlist = ['', 'abc', 'Xbcd', 'Xb', 'XY', 'abcd',
							πF.SetLineno(64)
							πTemp003 = make([]*πg.Object, 11)
							πTemp003[0] = ß.ToObject()
							πTemp003[1] = ßabc.ToObject()
							πTemp003[2] = ßXbcd.ToObject()
							πTemp003[3] = ßXb.ToObject()
							πTemp003[4] = ßXY.ToObject()
							πTemp003[5] = ßabcd.ToObject()
							πTemp003[6] = ßaXc.ToObject()
							πTemp003[7] = ßabd.ToObject()
							πTemp003[8] = ßab.ToObject()
							πTemp003[9] = ßaX.ToObject()
							πTemp003[10] = ßabcX.ToObject()
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							µtestlist = πTemp001
							if πE = πg.CheckLocal(πF, µtestlist, "testlist"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µtestlist); πE != nil {
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
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µs1 = πTemp002
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µtestlist, "testlist"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µtestlist); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp007 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
								πF.PopCheckpoint()
								goto Label6
							}
							if πTemp009, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µs2 = πTemp009
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 68: p = commonprefix([s1, s2])
							πF.SetLineno(68)
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = make([]*πg.Object, 2)
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							πTemp004[0] = µs1
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							πTemp004[1] = µs2
							πTemp009 = πg.NewList(πTemp004...).ToObject()
							πTemp003[0] = πTemp009
							if πE = πg.CheckLocal(πF, µcommonprefix, "commonprefix"); πE != nil {
								continue
							}
							if πTemp009, πE = µcommonprefix.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µp = πTemp009
							// line 69: self.assertTrue(s1.startswith(p))
							πF.SetLineno(69)
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp004[0] = µp
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µs1, ßstartswith, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp010
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 70: self.assertTrue(s2.startswith(p))
							πF.SetLineno(70)
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp004[0] = µp
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µs2, ßstartswith, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp010
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.NE(πF, µs1, µs2); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp009); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label7
							}
							goto Label8
							// line 71: if s1 != s2:
							πF.SetLineno(71)
						Label7:
							// line 72: n = len(p)
							πF.SetLineno(72)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp003[0] = µp
							if πTemp009, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µn = πTemp010
							// line 73: self.assertNotEqual(s1[n:n+1], s2[n:n+1])
							πF.SetLineno(73)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Add(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp009, πE = πg.SliceType.Call(πF, πg.Args{µn, πTemp010, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µs1, πTemp009); πE != nil {
								continue
							}
							πTemp003[0] = πTemp010
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Add(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp009, πE = πg.SliceType.Call(πF, πg.Args{µn, πTemp010, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µs2, πTemp009); πE != nil {
								continue
							}
							πTemp003[1] = πTemp010
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßassertNotEqual, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label8
						Label8:
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
					if πE = πClass.SetItem(πF, ßtest_commonprefix.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 75: def test_getsize(self):
					πF.SetLineno(75)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_getsize", "build/src/__python__/test/test_genericpath.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 76: f = open(test_support.TESTFN, "wb")
							πF.SetLineno(76)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp001[1] = ßwb.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µf = πTemp003
							// line 77: try:
							πF.SetLineno(77)
							πF.PushCheckpoint(1)
							// line 78: f.write("foo")
							πF.SetLineno(78)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßfoo.ToObject()
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µf, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 79: f.close()
							πF.SetLineno(79)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 80: self.assertEqual(self.pathmodule.getsize(test_support.TESTFN), 3)
							πF.SetLineno(80)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßgetsize, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							πF.PopCheckpoint()
							goto Label1
						Label1:
							πTemp005, πTemp006 = πF.RestoreExc(nil, nil)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µf, ßclosed, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label2
							}
							goto Label3
							// line 82: if not f.closed:
							πF.SetLineno(82)
						Label2:
							// line 83: f.close()
							πF.SetLineno(83)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label3
						Label3:
							// line 84: test_support.unlink(test_support.TESTFN)
							πF.SetLineno(84)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunlink, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005 != nil {
								πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
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
					if πE = πClass.SetItem(πF, ßtest_getsize.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 87: def test_time(self):
					πF.SetLineno(87)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_time", "build/src/__python__/test/test_genericpath.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 88: f = open(test_support.TESTFN, "wb")
							πF.SetLineno(88)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp001[1] = ßwb.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µf = πTemp003
							// line 89: try:
							πF.SetLineno(89)
							πF.PushCheckpoint(1)
							// line 90: f.write("foo")
							πF.SetLineno(90)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßfoo.ToObject()
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µf, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 91: f.close()
							πF.SetLineno(91)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 92: f = open(test_support.TESTFN, "ab")
							πF.SetLineno(92)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp001[1] = ßab.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µf = πTemp003
							// line 93: f.write("bar")
							πF.SetLineno(93)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßbar.ToObject()
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µf, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 94: f.close()
							πF.SetLineno(94)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 95: f = open(test_support.TESTFN, "rb")
							πF.SetLineno(95)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp001[1] = ßrb.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µf = πTemp003
							// line 96: d = f.read()
							πF.SetLineno(96)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µf, ßread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp003
							// line 97: f.close()
							πF.SetLineno(97)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 98: self.assertEqual(d, "foobar")
							πF.SetLineno(98)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[0] = µd
							πTemp001[1] = ßfoobar.ToObject()
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
							// line 100: self.assertLessEqual(
							πF.SetLineno(100)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßgetctime, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							πTemp004 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßgetmtime, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertLessEqual, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πF.PopCheckpoint()
							goto Label1
						Label1:
							πTemp005, πTemp006 = πF.RestoreExc(nil, nil)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µf, ßclosed, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label2
							}
							goto Label3
							// line 105: if not f.closed:
							πF.SetLineno(105)
						Label2:
							// line 106: f.close()
							πF.SetLineno(106)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label3
						Label3:
							// line 107: test_support.unlink(test_support.TESTFN)
							πF.SetLineno(107)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunlink, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005 != nil {
								πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
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
					if πE = πClass.SetItem(πF, ßtest_time.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 86: @unittest.skip('grumpy')
					πF.SetLineno(86)
					πTemp002 = πF.MakeArgs(1)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßtest_time); πE != nil {
						continue
					}
					πTemp002[0] = πTemp007
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
					if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πClass.SetItem(πF, ßtest_time.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 109: def test_exists(self):
					πF.SetLineno(109)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_exists", "build/src/__python__/test/test_genericpath.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µf *πg.Object = πg.UnboundLocal; _ = µf
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
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.BaseException
						_ = πTemp008
						var πTemp009 *πg.Traceback
						_ = πTemp009
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 110: self.assertIs(self.pathmodule.exists(test_support.TESTFN), False)
							πF.SetLineno(110)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßexists, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertIs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 111: f = open(test_support.TESTFN, "wb")
							πF.SetLineno(111)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp001[1] = ßwb.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µf = πTemp004
							// line 112: try:
							πF.SetLineno(112)
							πF.PushCheckpoint(1)
							// line 113: f.write("foo")
							πF.SetLineno(113)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßfoo.ToObject()
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µf, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 114: f.close()
							πF.SetLineno(114)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 115: self.assertIs(self.pathmodule.exists(test_support.TESTFN), True)
							πF.SetLineno(115)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßexists, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertIs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßgenericpath); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Eq(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label2
							}
							goto Label3
							// line 116: if not self.pathmodule == genericpath:
							πF.SetLineno(116)
						Label2:
							// line 117: self.assertIs(self.pathmodule.lexists(test_support.TESTFN),
							πF.SetLineno(117)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßlexists, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertIs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label3
						Label3:
							πF.PopCheckpoint()
							goto Label1
						Label1:
							πTemp008, πTemp009 = πF.RestoreExc(nil, nil)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label4
							}
							goto Label5
							// line 120: if not f.close():
							πF.SetLineno(120)
						Label4:
							// line 121: f.close()
							πF.SetLineno(121)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label5
						Label5:
							// line 122: test_support.unlink(test_support.TESTFN)
							πF.SetLineno(122)
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßunlink, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp008 != nil {
								πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
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
					if πE = πClass.SetItem(πF, ßtest_exists.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 124: def test_isdir(self):
					πF.SetLineno(124)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("test_isdir", "build/src/__python__/test/test_genericpath.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µf *πg.Object = πg.UnboundLocal; _ = µf
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 125: self.assertIs(self.pathmodule.isdir(test_support.TESTFN), False)
							πF.SetLineno(125)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßisdir, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertIs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 126: f = open(test_support.TESTFN, "wb")
							πF.SetLineno(126)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp001[1] = ßwb.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µf = πTemp004
							// line 127: try:
							πF.SetLineno(127)
							πF.PushCheckpoint(1)
							// line 128: f.write("foo")
							πF.SetLineno(128)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßfoo.ToObject()
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µf, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 129: f.close()
							πF.SetLineno(129)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 130: self.assertIs(self.pathmodule.isdir(test_support.TESTFN), False)
							πF.SetLineno(130)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßisdir, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertIs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 131: os.remove(test_support.TESTFN)
							πF.SetLineno(131)
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßremove, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 132: os.mkdir(test_support.TESTFN)
							πF.SetLineno(132)
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßmkdir, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 133: self.assertIs(self.pathmodule.isdir(test_support.TESTFN), True)
							πF.SetLineno(133)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßisdir, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertIs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 134: os.rmdir(test_support.TESTFN)
							πF.SetLineno(134)
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßrmdir, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πF.PopCheckpoint()
							goto Label1
						Label1:
							πTemp005, πTemp006 = πF.RestoreExc(nil, nil)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp008).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label2
							}
							goto Label3
							// line 136: if not f.close():
							πF.SetLineno(136)
						Label2:
							// line 137: f.close()
							πF.SetLineno(137)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label3
						Label3:
							// line 138: test_support.unlink(test_support.TESTFN)
							πF.SetLineno(138)
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßunlink, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 139: safe_rmdir(test_support.TESTFN)
							πF.SetLineno(139)
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsafe_rmdir); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005 != nil {
								πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
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
					if πE = πClass.SetItem(πF, ßtest_isdir.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 142: def test_isfile(self):
					πF.SetLineno(142)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("test_isfile", "build/src/__python__/test/test_genericpath.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µf *πg.Object = πg.UnboundLocal; _ = µf
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 143: self.assertIs(self.pathmodule.isfile(test_support.TESTFN), False)
							πF.SetLineno(143)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßisfile, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertIs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 144: f = open(test_support.TESTFN, "wb")
							πF.SetLineno(144)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp001[1] = ßwb.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µf = πTemp004
							// line 145: try:
							πF.SetLineno(145)
							πF.PushCheckpoint(1)
							// line 146: f.write("foo")
							πF.SetLineno(146)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßfoo.ToObject()
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µf, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 147: f.close()
							πF.SetLineno(147)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 148: self.assertIs(self.pathmodule.isfile(test_support.TESTFN), True)
							πF.SetLineno(148)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßisfile, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertIs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 149: os.remove(test_support.TESTFN)
							πF.SetLineno(149)
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßremove, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 150: os.mkdir(test_support.TESTFN)
							πF.SetLineno(150)
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßmkdir, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 151: self.assertIs(self.pathmodule.isfile(test_support.TESTFN), False)
							πF.SetLineno(151)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßisfile, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertIs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 152: os.rmdir(test_support.TESTFN)
							πF.SetLineno(152)
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßrmdir, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πF.PopCheckpoint()
							goto Label1
						Label1:
							πTemp005, πTemp006 = πF.RestoreExc(nil, nil)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp008).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label2
							}
							goto Label3
							// line 154: if not f.close():
							πF.SetLineno(154)
						Label2:
							// line 155: f.close()
							πF.SetLineno(155)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label3
						Label3:
							// line 156: test_support.unlink(test_support.TESTFN)
							πF.SetLineno(156)
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßunlink, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 157: safe_rmdir(test_support.TESTFN)
							πF.SetLineno(157)
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTESTFN, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsafe_rmdir); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005 != nil {
								πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
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
					if πE = πClass.SetItem(πF, ßtest_isfile.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 141: @unittest.skip('grumpy')
					πF.SetLineno(141)
					πTemp002 = πF.MakeArgs(1)
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßtest_isfile); πE != nil {
						continue
					}
					πTemp002[0] = πTemp011
					πTemp008 = πF.MakeArgs(1)
					πTemp008[0] = ßgrumpy.ToObject()
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßskip, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp012.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πTemp012, πE = πTemp011.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πClass.SetItem(πF, ßtest_isfile.ToObject(), πTemp012); πE != nil {
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
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("GenericTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßGenericTest.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 163: class CommonTest(GenericTest):
			πF.SetLineno(163)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßGenericTest); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp004 = πg.NewDict()
			if πTemp005, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp005); πE != nil {
				continue
			}
			_, πE = πg.NewCode("CommonTest", "build/src/__python__/test/test_genericpath.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 []πg.Param
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 []*πg.Object
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 165: pathmodule = None
					πF.SetLineno(165)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßpathmodule.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 166: common_attributes = GenericTest.common_attributes + [
					πF.SetLineno(166)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßGenericTest); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcommon_attributes, nil); πE != nil {
						continue
					}
					πTemp004 = make([]*πg.Object, 25)
					πTemp004[0] = ßcurdir.ToObject()
					πTemp004[1] = ßpardir.ToObject()
					πTemp004[2] = ßextsep.ToObject()
					πTemp004[3] = ßsep.ToObject()
					πTemp004[4] = ßpathsep.ToObject()
					πTemp004[5] = ßdefpath.ToObject()
					πTemp004[6] = ßaltsep.ToObject()
					πTemp004[7] = ßdevnull.ToObject()
					πTemp004[8] = ßnormcase.ToObject()
					πTemp004[9] = ßsplitdrive.ToObject()
					πTemp004[10] = ßexpandvars.ToObject()
					πTemp004[11] = ßnormpath.ToObject()
					πTemp004[12] = ßabspath.ToObject()
					πTemp004[13] = ßjoin.ToObject()
					πTemp004[14] = ßsplit.ToObject()
					πTemp004[15] = ßsplitext.ToObject()
					πTemp004[16] = ßisabs.ToObject()
					πTemp004[17] = ßbasename.ToObject()
					πTemp004[18] = ßdirname.ToObject()
					πTemp004[19] = ßlexists.ToObject()
					πTemp004[20] = ßislink.ToObject()
					πTemp004[21] = ßismount.ToObject()
					πTemp004[22] = ßexpanduser.ToObject()
					πTemp004[23] = ßnormpath.ToObject()
					πTemp004[24] = ßrealpath.ToObject()
					πTemp002 = πg.NewList(πTemp004...).ToObject()
					if πTemp001, πE = πg.Add(πF, πTemp003, πTemp002); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßcommon_attributes.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 176: def test_normcase(self):
					πF.SetLineno(176)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_normcase", "build/src/__python__/test/test_genericpath.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µp *πg.Object = πg.UnboundLocal; _ = µp
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
							// line 178: p = "FoO/./BaR"
							πF.SetLineno(178)
							µp = πg.NewStr("FoO/./BaR").ToObject()
							// line 179: p = self.pathmodule.normcase(p)
							πF.SetLineno(179)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp001[0] = µp
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßnormcase, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µp = πTemp002
							// line 180: self.assertEqual(p, self.pathmodule.normcase(p))
							πF.SetLineno(180)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp001[0] = µp
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp004[0] = µp
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßnormcase, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßtest_normcase.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 182: def test_splitdrive(self):
					πF.SetLineno(182)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("test_splitdrive", "build/src/__python__/test/test_genericpath.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsplitdrive *πg.Object = πg.UnboundLocal; _ = µsplitdrive
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
							// line 184: splitdrive = self.pathmodule.splitdrive
							πF.SetLineno(184)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsplitdrive, nil); πE != nil {
								continue
							}
							µsplitdrive = πTemp002
							// line 185: self.assertEqual(splitdrive("/foo/bar"), ("", "/foo/bar"))
							πF.SetLineno(185)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("/foo/bar").ToObject()
							if πE = πg.CheckLocal(πF, µsplitdrive, "splitdrive"); πE != nil {
								continue
							}
							if πTemp001, πE = µsplitdrive.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp001
							πTemp001 = πg.NewTuple2(ß.ToObject(), πg.NewStr("/foo/bar").ToObject()).ToObject()
							πTemp003[1] = πTemp001
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
							// line 186: self.assertEqual(splitdrive("foo:bar"), ("", "foo:bar"))
							πF.SetLineno(186)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("foo:bar").ToObject()
							if πE = πg.CheckLocal(πF, µsplitdrive, "splitdrive"); πE != nil {
								continue
							}
							if πTemp001, πE = µsplitdrive.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp001
							πTemp001 = πg.NewTuple2(ß.ToObject(), πg.NewStr("foo:bar").ToObject()).ToObject()
							πTemp003[1] = πTemp001
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
							// line 187: self.assertEqual(splitdrive(":foo:bar"), ("", ":foo:bar"))
							πF.SetLineno(187)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr(":foo:bar").ToObject()
							if πE = πg.CheckLocal(πF, µsplitdrive, "splitdrive"); πE != nil {
								continue
							}
							if πTemp001, πE = µsplitdrive.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp001
							πTemp001 = πg.NewTuple2(ß.ToObject(), πg.NewStr(":foo:bar").ToObject()).ToObject()
							πTemp003[1] = πTemp001
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
					if πE = πClass.SetItem(πF, ßtest_splitdrive.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 189: def test_expandvars(self):
					πF.SetLineno(189)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_expandvars", "build/src/__python__/test/test_genericpath.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µexpandvars *πg.Object = πg.UnboundLocal; _ = µexpandvars
						var µenv *πg.Object = πg.UnboundLocal; _ = µenv
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 []*πg.Object
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
							case 3: goto Label3
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__name__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp003, ßmacpath.ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 190: if self.pathmodule.__name__ == 'macpath':
							πF.SetLineno(190)
						Label1:
							// line 191: self.skipTest('macpath.expandvars is a stub')
							πF.SetLineno(191)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("macpath.expandvars is a stub").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßskipTest, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label2
						Label2:
							// line 192: expandvars = self.pathmodule.expandvars
							πF.SetLineno(192)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßexpandvars, nil); πE != nil {
								continue
							}
							µexpandvars = πTemp002
							// line 193: with test_support.EnvironmentVarGuard() as env:
							πF.SetLineno(193)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßEnvironmentVarGuard, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp003.Call(πF, πg.Args{πTemp001}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(3)
							µenv = πTemp003
							// line 194: env.clear()
							πF.SetLineno(194)
							if πE = πg.CheckLocal(πF, µenv, "env"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µenv, ßclear, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 195: env["foo"] = "bar"
							πF.SetLineno(195)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, ßbar.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µenv, "env"); πE != nil {
								continue
							}
							πTemp007 = ßfoo.ToObject()
							if πE = πg.SetItem(πF, µenv, πTemp007, πTemp006); πE != nil {
								continue
							}
							// line 196: env["{foo"] = "baz1"
							πF.SetLineno(196)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, ßbaz1.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µenv, "env"); πE != nil {
								continue
							}
							πTemp007 = πg.NewStr("{foo").ToObject()
							if πE = πg.SetItem(πF, µenv, πTemp007, πTemp006); πE != nil {
								continue
							}
							// line 197: env["{foo}"] = "baz2"
							πF.SetLineno(197)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, ßbaz2.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µenv, "env"); πE != nil {
								continue
							}
							πTemp007 = πg.NewStr("{foo}").ToObject()
							if πE = πg.SetItem(πF, µenv, πTemp007, πTemp006); πE != nil {
								continue
							}
							// line 198: self.assertEqual(expandvars("foo"), "foo")
							πF.SetLineno(198)
							πTemp005 = πF.MakeArgs(2)
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = ßfoo.ToObject()
							if πE = πg.CheckLocal(πF, µexpandvars, "expandvars"); πE != nil {
								continue
							}
							if πTemp006, πE = µexpandvars.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp005[0] = πTemp006
							πTemp005[1] = ßfoo.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 199: self.assertEqual(expandvars("$foo bar"), "bar bar")
							πF.SetLineno(199)
							πTemp005 = πF.MakeArgs(2)
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = πg.NewStr("$foo bar").ToObject()
							if πE = πg.CheckLocal(πF, µexpandvars, "expandvars"); πE != nil {
								continue
							}
							if πTemp006, πE = µexpandvars.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp005[0] = πTemp006
							πTemp005[1] = πg.NewStr("bar bar").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 200: self.assertEqual(expandvars("${foo}bar"), "barbar")
							πF.SetLineno(200)
							πTemp005 = πF.MakeArgs(2)
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = πg.NewStr("${foo}bar").ToObject()
							if πE = πg.CheckLocal(πF, µexpandvars, "expandvars"); πE != nil {
								continue
							}
							if πTemp006, πE = µexpandvars.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp005[0] = πTemp006
							πTemp005[1] = ßbarbar.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 201: self.assertEqual(expandvars("$[foo]bar"), "$[foo]bar")
							πF.SetLineno(201)
							πTemp005 = πF.MakeArgs(2)
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = πg.NewStr("$[foo]bar").ToObject()
							if πE = πg.CheckLocal(πF, µexpandvars, "expandvars"); πE != nil {
								continue
							}
							if πTemp006, πE = µexpandvars.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp005[0] = πTemp006
							πTemp005[1] = πg.NewStr("$[foo]bar").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 202: self.assertEqual(expandvars("$bar bar"), "$bar bar")
							πF.SetLineno(202)
							πTemp005 = πF.MakeArgs(2)
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = πg.NewStr("$bar bar").ToObject()
							if πE = πg.CheckLocal(πF, µexpandvars, "expandvars"); πE != nil {
								continue
							}
							if πTemp006, πE = µexpandvars.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp005[0] = πTemp006
							πTemp005[1] = πg.NewStr("$bar bar").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 203: self.assertEqual(expandvars("$?bar"), "$?bar")
							πF.SetLineno(203)
							πTemp005 = πF.MakeArgs(2)
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = πg.NewStr("$?bar").ToObject()
							if πE = πg.CheckLocal(πF, µexpandvars, "expandvars"); πE != nil {
								continue
							}
							if πTemp006, πE = µexpandvars.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp005[0] = πTemp006
							πTemp005[1] = πg.NewStr("$?bar").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 204: self.assertEqual(expandvars("$foo}bar"), "bar}bar")
							πF.SetLineno(204)
							πTemp005 = πF.MakeArgs(2)
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = πg.NewStr("$foo}bar").ToObject()
							if πE = πg.CheckLocal(πF, µexpandvars, "expandvars"); πE != nil {
								continue
							}
							if πTemp006, πE = µexpandvars.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp005[0] = πTemp006
							πTemp005[1] = πg.NewStr("bar}bar").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 205: self.assertEqual(expandvars("${foo"), "${foo")
							πF.SetLineno(205)
							πTemp005 = πF.MakeArgs(2)
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = πg.NewStr("${foo").ToObject()
							if πE = πg.CheckLocal(πF, µexpandvars, "expandvars"); πE != nil {
								continue
							}
							if πTemp006, πE = µexpandvars.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp005[0] = πTemp006
							πTemp005[1] = πg.NewStr("${foo").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 206: self.assertEqual(expandvars("${{foo}}"), "baz1}")
							πF.SetLineno(206)
							πTemp005 = πF.MakeArgs(2)
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = πg.NewStr("${{foo}}").ToObject()
							if πE = πg.CheckLocal(πF, µexpandvars, "expandvars"); πE != nil {
								continue
							}
							if πTemp006, πE = µexpandvars.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp005[0] = πTemp006
							πTemp005[1] = πg.NewStr("baz1}").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 207: self.assertEqual(expandvars("$foo$foo"), "barbar")
							πF.SetLineno(207)
							πTemp005 = πF.MakeArgs(2)
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = πg.NewStr("$foo$foo").ToObject()
							if πE = πg.CheckLocal(πF, µexpandvars, "expandvars"); πE != nil {
								continue
							}
							if πTemp006, πE = µexpandvars.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp005[0] = πTemp006
							πTemp005[1] = ßbarbar.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 208: self.assertEqual(expandvars("$bar$bar"), "$bar$bar")
							πF.SetLineno(208)
							πTemp005 = πF.MakeArgs(2)
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = πg.NewStr("$bar$bar").ToObject()
							if πE = πg.CheckLocal(πF, µexpandvars, "expandvars"); πE != nil {
								continue
							}
							if πTemp006, πE = µexpandvars.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp005[0] = πTemp006
							πTemp005[1] = πg.NewStr("$bar$bar").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πF.PopCheckpoint()
						Label3:
							πTemp009, πTemp010 = nil, nil
							if πE != nil {
								πTemp009, πTemp010 = πF.ExcInfo()
							}
							if πTemp009 != nil {
								πTemp011 = πTemp009.Type()
								if πTemp006, πE = πTemp002.Call(πF, πg.Args{πTemp001, πTemp011.ToObject(), πTemp009.ToObject(), πTemp010.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp006, πE = πTemp002.Call(πF, πg.Args{πTemp001, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp009 != nil && πTemp004 != true {
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
					if πE = πClass.SetItem(πF, ßtest_expandvars.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 211: def test_expandvars_nonascii(self):
					πF.SetLineno(211)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_expandvars_nonascii", "build/src/__python__/test/test_genericpath.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µexpandvars *πg.Object = πg.UnboundLocal; _ = µexpandvars
						var µcheck *πg.Object = πg.UnboundLocal; _ = µcheck
						var µencoding *πg.Object = πg.UnboundLocal; _ = µencoding
						var µenv *πg.Object = πg.UnboundLocal; _ = µenv
						var µunonascii *πg.Object = πg.UnboundLocal; _ = µunonascii
						var µsnonascii *πg.Object = πg.UnboundLocal; _ = µsnonascii
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
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
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 *πg.BaseException
						_ = πTemp011
						var πTemp012 *πg.Traceback
						_ = πTemp012
						var πTemp013 *πg.Type
						_ = πTemp013
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 3: goto Label3
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__name__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp003, ßmacpath.ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 212: if self.pathmodule.__name__ == 'macpath':
							πF.SetLineno(212)
						Label1:
							// line 213: self.skipTest('macpath.expandvars is a stub')
							πF.SetLineno(213)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("macpath.expandvars is a stub").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßskipTest, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label2
						Label2:
							// line 214: expandvars = self.pathmodule.expandvars
							πF.SetLineno(214)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßexpandvars, nil); πE != nil {
								continue
							}
							µexpandvars = πTemp002
							// line 215: def check(value, expected):
							πF.SetLineno(215)
							πTemp006 = make([]πg.Param, 2)
							πTemp006[0] = πg.Param{Name: "value", Def: nil}
							πTemp006[1] = πg.Param{Name: "expected", Def: nil}
							πTemp001 = πg.NewFunction(πg.NewCode("check", "build/src/__python__/test/test_genericpath.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µvalue *πg.Object = πArgs[0]; _ = µvalue
								var µexpected *πg.Object = πArgs[1]; _ = µexpected
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
									// line 216: self.assertEqual(expandvars(value), expected)
									πF.SetLineno(216)
									πTemp001 = πF.MakeArgs(2)
									πTemp002 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
										continue
									}
									πTemp002[0] = µvalue
									if πE = πg.CheckLocal(πF, µexpandvars, "expandvars"); πE != nil {
										continue
									}
									if πTemp003, πE = µexpandvars.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									πTemp001[0] = πTemp003
									if πE = πg.CheckLocal(πF, µexpected, "expected"); πE != nil {
										continue
									}
									πTemp001[1] = µexpected
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
							µcheck = πTemp001
							// line 217: encoding = sys.getfilesystemencoding()
							πF.SetLineno(217)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßgetfilesystemencoding, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µencoding = πTemp002
							// line 218: with test_support.EnvironmentVarGuard() as env:
							πF.SetLineno(218)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßEnvironmentVarGuard, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp002.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp007.Call(πF, πg.Args{πTemp002}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(3)
							µenv = πTemp007
							// line 219: env.clear()
							πF.SetLineno(219)
							if πE = πg.CheckLocal(πF, µenv, "env"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µenv, ßclear, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 220: unonascii = test_support.FS_NONASCII
							πF.SetLineno(220)
							if πTemp008, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßFS_NONASCII, nil); πE != nil {
								continue
							}
							µunonascii = πTemp009
							// line 221: snonascii = unonascii.encode(encoding)
							πF.SetLineno(221)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
								continue
							}
							πTemp005[0] = µencoding
							if πE = πg.CheckLocal(πF, µunonascii, "unonascii"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µunonascii, ßencode, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µsnonascii = πTemp009
							// line 222: env['spam'] = snonascii
							πF.SetLineno(222)
							if πE = πg.CheckLocal(πF, µsnonascii, "snonascii"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, µsnonascii); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µenv, "env"); πE != nil {
								continue
							}
							πTemp009 = ßspam.ToObject()
							if πE = πg.SetItem(πF, µenv, πTemp009, πTemp008); πE != nil {
								continue
							}
							// line 223: env[snonascii] = 'ham' + snonascii
							πF.SetLineno(223)
							if πE = πg.CheckLocal(πF, µsnonascii, "snonascii"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Add(πF, ßham.ToObject(), µsnonascii); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πTemp008); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µenv, "env"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsnonascii, "snonascii"); πE != nil {
								continue
							}
							πTemp010 = µsnonascii
							if πE = πg.SetItem(πF, µenv, πTemp010, πTemp009); πE != nil {
								continue
							}
							// line 224: check(snonascii, snonascii)
							πF.SetLineno(224)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µsnonascii, "snonascii"); πE != nil {
								continue
							}
							πTemp005[0] = µsnonascii
							if πE = πg.CheckLocal(πF, µsnonascii, "snonascii"); πE != nil {
								continue
							}
							πTemp005[1] = µsnonascii
							if πE = πg.CheckLocal(πF, µcheck, "check"); πE != nil {
								continue
							}
							if πTemp008, πE = µcheck.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 225: check('$spam bar', '%s bar' % snonascii)
							πF.SetLineno(225)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = πg.NewStr("$spam bar").ToObject()
							if πE = πg.CheckLocal(πF, µsnonascii, "snonascii"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Mod(πF, πg.NewStr("%s bar").ToObject(), µsnonascii); πE != nil {
								continue
							}
							πTemp005[1] = πTemp008
							if πE = πg.CheckLocal(πF, µcheck, "check"); πE != nil {
								continue
							}
							if πTemp008, πE = µcheck.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 226: check('${spam}bar', '%sbar' % snonascii)
							πF.SetLineno(226)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = πg.NewStr("${spam}bar").ToObject()
							if πE = πg.CheckLocal(πF, µsnonascii, "snonascii"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Mod(πF, πg.NewStr("%sbar").ToObject(), µsnonascii); πE != nil {
								continue
							}
							πTemp005[1] = πTemp008
							if πE = πg.CheckLocal(πF, µcheck, "check"); πE != nil {
								continue
							}
							if πTemp008, πE = µcheck.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 227: check('${%s}bar' % snonascii, 'ham%sbar' % snonascii)
							πF.SetLineno(227)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µsnonascii, "snonascii"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Mod(πF, πg.NewStr("${%s}bar").ToObject(), µsnonascii); πE != nil {
								continue
							}
							πTemp005[0] = πTemp008
							if πE = πg.CheckLocal(πF, µsnonascii, "snonascii"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Mod(πF, πg.NewStr("ham%sbar").ToObject(), µsnonascii); πE != nil {
								continue
							}
							πTemp005[1] = πTemp008
							if πE = πg.CheckLocal(πF, µcheck, "check"); πE != nil {
								continue
							}
							if πTemp008, πE = µcheck.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 228: check('$bar%s bar' % snonascii, '$bar%s bar' % snonascii)
							πF.SetLineno(228)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µsnonascii, "snonascii"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Mod(πF, πg.NewStr("$bar%s bar").ToObject(), µsnonascii); πE != nil {
								continue
							}
							πTemp005[0] = πTemp008
							if πE = πg.CheckLocal(πF, µsnonascii, "snonascii"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Mod(πF, πg.NewStr("$bar%s bar").ToObject(), µsnonascii); πE != nil {
								continue
							}
							πTemp005[1] = πTemp008
							if πE = πg.CheckLocal(πF, µcheck, "check"); πE != nil {
								continue
							}
							if πTemp008, πE = µcheck.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 229: check('$spam}bar', '%s}bar' % snonascii)
							πF.SetLineno(229)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = πg.NewStr("$spam}bar").ToObject()
							if πE = πg.CheckLocal(πF, µsnonascii, "snonascii"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Mod(πF, πg.NewStr("%s}bar").ToObject(), µsnonascii); πE != nil {
								continue
							}
							πTemp005[1] = πTemp008
							if πE = πg.CheckLocal(πF, µcheck, "check"); πE != nil {
								continue
							}
							if πTemp008, πE = µcheck.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 231: check(unonascii, unonascii)
							πF.SetLineno(231)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µunonascii, "unonascii"); πE != nil {
								continue
							}
							πTemp005[0] = µunonascii
							if πE = πg.CheckLocal(πF, µunonascii, "unonascii"); πE != nil {
								continue
							}
							πTemp005[1] = µunonascii
							if πE = πg.CheckLocal(πF, µcheck, "check"); πE != nil {
								continue
							}
							if πTemp008, πE = µcheck.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 232: check(u'$spam bar', u'%s bar' % unonascii)
							πF.SetLineno(232)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = πg.NewUnicode("$spam bar").ToObject()
							if πE = πg.CheckLocal(πF, µunonascii, "unonascii"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Mod(πF, πg.NewUnicode("%s bar").ToObject(), µunonascii); πE != nil {
								continue
							}
							πTemp005[1] = πTemp008
							if πE = πg.CheckLocal(πF, µcheck, "check"); πE != nil {
								continue
							}
							if πTemp008, πE = µcheck.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 233: check(u'${spam}bar', u'%sbar' % unonascii)
							πF.SetLineno(233)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = πg.NewUnicode("${spam}bar").ToObject()
							if πE = πg.CheckLocal(πF, µunonascii, "unonascii"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Mod(πF, πg.NewUnicode("%sbar").ToObject(), µunonascii); πE != nil {
								continue
							}
							πTemp005[1] = πTemp008
							if πE = πg.CheckLocal(πF, µcheck, "check"); πE != nil {
								continue
							}
							if πTemp008, πE = µcheck.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 234: check(u'${%s}bar' % unonascii, u'ham%sbar' % unonascii)
							πF.SetLineno(234)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µunonascii, "unonascii"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Mod(πF, πg.NewUnicode("${%s}bar").ToObject(), µunonascii); πE != nil {
								continue
							}
							πTemp005[0] = πTemp008
							if πE = πg.CheckLocal(πF, µunonascii, "unonascii"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Mod(πF, πg.NewUnicode("ham%sbar").ToObject(), µunonascii); πE != nil {
								continue
							}
							πTemp005[1] = πTemp008
							if πE = πg.CheckLocal(πF, µcheck, "check"); πE != nil {
								continue
							}
							if πTemp008, πE = µcheck.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 235: check(u'$bar%s bar' % unonascii, u'$bar%s bar' % unonascii)
							πF.SetLineno(235)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µunonascii, "unonascii"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Mod(πF, πg.NewUnicode("$bar%s bar").ToObject(), µunonascii); πE != nil {
								continue
							}
							πTemp005[0] = πTemp008
							if πE = πg.CheckLocal(πF, µunonascii, "unonascii"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Mod(πF, πg.NewUnicode("$bar%s bar").ToObject(), µunonascii); πE != nil {
								continue
							}
							πTemp005[1] = πTemp008
							if πE = πg.CheckLocal(πF, µcheck, "check"); πE != nil {
								continue
							}
							if πTemp008, πE = µcheck.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 236: check(u'$spam}bar', u'%s}bar' % unonascii)
							πF.SetLineno(236)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = πg.NewUnicode("$spam}bar").ToObject()
							if πE = πg.CheckLocal(πF, µunonascii, "unonascii"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Mod(πF, πg.NewUnicode("%s}bar").ToObject(), µunonascii); πE != nil {
								continue
							}
							πTemp005[1] = πTemp008
							if πE = πg.CheckLocal(πF, µcheck, "check"); πE != nil {
								continue
							}
							if πTemp008, πE = µcheck.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πF.PopCheckpoint()
						Label3:
							πTemp011, πTemp012 = nil, nil
							if πE != nil {
								πTemp011, πTemp012 = πF.ExcInfo()
							}
							if πTemp011 != nil {
								πTemp013 = πTemp011.Type()
								if πTemp008, πE = πTemp003.Call(πF, πg.Args{πTemp002, πTemp013.ToObject(), πTemp011.ToObject(), πTemp012.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp008, πE = πTemp003.Call(πF, πg.Args{πTemp002, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							if πTemp011 != nil && πTemp004 != true {
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
					if πE = πClass.SetItem(πF, ßtest_expandvars_nonascii.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 210: @unittest.skipUnless(test_support.FS_NONASCII, 'need test_support.FS_NONASCII')
					πF.SetLineno(210)
					πTemp004 = πF.MakeArgs(1)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßtest_expandvars_nonascii); πE != nil {
						continue
					}
					πTemp004[0] = πTemp007
					πTemp008 = πF.MakeArgs(2)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßtest_support); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßFS_NONASCII, nil); πE != nil {
						continue
					}
					πTemp008[0] = πTemp009
					πTemp008[1] = πg.NewStr("need test_support.FS_NONASCII").ToObject()
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßskipUnless, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp009.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πTemp009, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_expandvars_nonascii.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 238: def test_abspath(self):
					πF.SetLineno(238)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_abspath", "build/src/__python__/test/test_genericpath.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µpath *πg.Object = πg.UnboundLocal; _ = µpath
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 239: self.assertIn("foo", self.pathmodule.abspath("foo"))
							πF.SetLineno(239)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßfoo.ToObject()
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßfoo.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßabspath, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertIn, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp004 = πg.NewTuple5(ß.ToObject(), ßfoo.ToObject(), πg.NewStr("f\xf2\xf2").ToObject(), πg.NewStr("/foo").ToObject(), πg.NewStr("C:\\").ToObject()).ToObject()
							if πTemp003, πE = πg.Iter(πF, πTemp004); πE != nil {
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
								µpath = πTemp004
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 243: self.assertIsInstance(self.pathmodule.abspath(path), str)
							πF.SetLineno(243)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp002[0] = µpath
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp004, ßabspath, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertIsInstance, nil); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_abspath.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 245: def test_realpath(self):
					πF.SetLineno(245)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("test_realpath", "build/src/__python__/test/test_genericpath.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 246: self.assertIn("foo", self.pathmodule.realpath("foo"))
							πF.SetLineno(246)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßfoo.ToObject()
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßfoo.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßrealpath, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertIn, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_realpath.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 249: def test_normpath_issue5827(self):
					πF.SetLineno(249)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("test_normpath_issue5827", "build/src/__python__/test/test_genericpath.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µpath *πg.Object = πg.UnboundLocal; _ = µpath
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
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
							πTemp002 = πg.NewTuple5(πg.NewUnicode("").ToObject(), πg.NewUnicode(".").ToObject(), πg.NewUnicode("/").ToObject(), πg.NewUnicode("\\").ToObject(), πg.NewUnicode("///foo/.//bar//").ToObject()).ToObject()
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
								µpath = πTemp002
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 252: self.assertIsInstance(self.pathmodule.normpath(path), unicode)
							πF.SetLineno(252)
							πTemp005 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp006[0] = µpath
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp002, ßnormpath, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							πTemp005[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIsInstance, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
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
					if πE = πClass.SetItem(πF, ßtest_normpath_issue5827.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 248: @test_support.requires_unicode
					πF.SetLineno(248)
					πTemp004 = πF.MakeArgs(1)
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßtest_normpath_issue5827); πE != nil {
						continue
					}
					πTemp004[0] = πTemp011
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßtest_support); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßrequires_unicode, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp012.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_normpath_issue5827.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 255: def test_abspath_issue3426(self):
					πF.SetLineno(255)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("test_abspath_issue3426", "build/src/__python__/test/test_genericpath.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µabspath *πg.Object = πg.UnboundLocal; _ = µabspath
						var µpath *πg.Object = πg.UnboundLocal; _ = µpath
						var µunicwd *πg.Object = πg.UnboundLocal; _ = µunicwd
						var µfsencoding *πg.Object = πg.UnboundLocal; _ = µfsencoding
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 []*πg.Object
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
						var πTemp011 *πg.BaseException
						_ = πTemp011
						var πTemp012 *πg.Traceback
						_ = πTemp012
						var πTemp013 *πg.Type
						_ = πTemp013
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							case 5: goto Label5
							case 7: goto Label7
							case 8: goto Label8
							case 9: goto Label9
							default: panic("unexpected function state")
							}
							// line 258: abspath = self.pathmodule.abspath
							πF.SetLineno(258)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßpathmodule, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßabspath, nil); πE != nil {
								continue
							}
							µabspath = πTemp002
							πTemp002 = πg.NewTuple5(πg.NewUnicode("").ToObject(), πg.NewUnicode("fuu").ToObject(), πg.NewUnicode("f\xc3\xb9\xc3\xb9").ToObject(), πg.NewUnicode("/fuu").ToObject(), πg.NewUnicode("U:\\").ToObject()).ToObject()
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
								µpath = πTemp002
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 260: self.assertIsInstance(abspath(path), unicode)
							πF.SetLineno(260)
							πTemp005 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp006[0] = µpath
							if πE = πg.CheckLocal(πF, µabspath, "abspath"); πE != nil {
								continue
							}
							if πTemp002, πE = µabspath.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							πTemp005[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIsInstance, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 262: unicwd = u'\xe7w\xf0'
							πF.SetLineno(262)
							µunicwd = πg.NewUnicode("\xc3\xa7w\xc3\xb0").ToObject()
							// line 263: try:
							πF.SetLineno(263)
							πF.PushCheckpoint(5)
							// line 264: fsencoding = test_support.TESTFN_ENCODING or "ascii"
							πF.SetLineno(264)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp002, ßTESTFN_ENCODING, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp007
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							πTemp001 = ßascii.ToObject()
						Label6:
							µfsencoding = πTemp001
							// line 265: unicwd.encode(fsencoding)
							πF.SetLineno(265)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfsencoding, "fsencoding"); πE != nil {
								continue
							}
							πTemp005[0] = µfsencoding
							if πE = πg.CheckLocal(πF, µunicwd, "unicwd"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µunicwd, ßencode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πF.PopCheckpoint()
							// line 270: with test_support.temp_cwd(unicwd):
							πF.SetLineno(270)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µunicwd, "unicwd"); πE != nil {
								continue
							}
							πTemp005[0] = µunicwd
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtemp_cwd, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πg.GetAttr(πF, πTemp001.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp001.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp007.Call(πF, πg.Args{πTemp001}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(7)
							πTemp009 = πg.NewTuple5(πg.NewUnicode("").ToObject(), πg.NewUnicode("fuu").ToObject(), πg.NewUnicode("f\xc3\xb9\xc3\xb9").ToObject(), πg.NewUnicode("/fuu").ToObject(), πg.NewUnicode("U:\\").ToObject()).ToObject()
							if πTemp008, πE = πg.Iter(πF, πTemp009); πE != nil {
								continue
							}
							πF.PushCheckpoint(9)
							πTemp003 = false
						Label8:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
								πF.PopCheckpoint()
								goto Label10
							}
							if πTemp009, πE = πg.Next(πF, πTemp008); πE != nil {
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
								µpath = πTemp009
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(8)            
							// line 272: self.assertIsInstance(abspath(path), unicode)
							πF.SetLineno(272)
							πTemp005 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp006[0] = µpath
							if πE = πg.CheckLocal(πF, µabspath, "abspath"); πE != nil {
								continue
							}
							if πTemp009, πE = µabspath.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp009
							if πTemp009, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							πTemp005[1] = πTemp009
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßassertIsInstance, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							continue
						Label9:
							if πE != nil || πR != nil {
								continue
							}
						Label10:
							πF.PopCheckpoint()
						Label7:
							πTemp011, πTemp012 = nil, nil
							if πE != nil {
								πTemp011, πTemp012 = πF.ExcInfo()
							}
							if πTemp011 != nil {
								πTemp013 = πTemp011.Type()
								if πTemp008, πE = πTemp002.Call(πF, πg.Args{πTemp001, πTemp013.ToObject(), πTemp011.ToObject(), πTemp012.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp008, πE = πTemp002.Call(πF, πg.Args{πTemp001, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							if πTemp011 != nil && πTemp003 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
								continue
							}
							goto Label4
						Label5:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp011, πTemp012 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßUnicodeEncodeError); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πTemp007).ToObject()
							if πTemp003, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label11
							}
							πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
							continue
							// line 266: except (AttributeError, UnicodeEncodeError):
							πF.SetLineno(266)
						Label11:
							// line 268: pass
							πF.SetLineno(268)
							πF.RestoreExc(nil, nil)
							goto Label4
						Label4:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_abspath_issue3426.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 254: @test_support.requires_unicode
					πF.SetLineno(254)
					πTemp004 = πF.MakeArgs(1)
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßtest_abspath_issue3426); πE != nil {
						continue
					}
					πTemp004[0] = πTemp012
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßtest_support); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßrequires_unicode, nil); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp013.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_abspath_issue3426.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 275: "Mac OS X denies the creation of a directory with an invalid utf8 name")
					πF.SetLineno(275)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("test_nonascii_abspath", "build/src/__python__/test/test_genericpath.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.BaseException
						_ = πTemp008
						var πTemp009 *πg.Traceback
						_ = πTemp009
						var πTemp010 *πg.Type
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 278: with test_support.temp_cwd('\xe7w\xf0'):
							πF.SetLineno(278)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("\xe7w\xf0").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtemp_cwd, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.GetAttr(πF, πTemp002.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp004.Call(πF, πg.Args{πTemp002}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(1)
							// line 279: self.test_abspath()
							πF.SetLineno(279)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßtest_abspath, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							πF.PopCheckpoint()
						Label1:
							πTemp008, πTemp009 = nil, nil
							if πE != nil {
								πTemp008, πTemp009 = πF.ExcInfo()
							}
							if πTemp008 != nil {
								πTemp010 = πTemp008.Type()
								if πTemp005, πE = πTemp003.Call(πF, πg.Args{πTemp002, πTemp010.ToObject(), πTemp008.ToObject(), πTemp009.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp005, πE = πTemp003.Call(πF, πg.Args{πTemp002, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp008 != nil && πTemp007 != true {
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
					if πE = πClass.SetItem(πF, ßtest_nonascii_abspath.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 274: @unittest.skipIf(sys.platform == 'darwin',
					πF.SetLineno(274)
					πTemp004 = πF.MakeArgs(1)
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßtest_nonascii_abspath); πE != nil {
						continue
					}
					πTemp004[0] = πTemp013
					πTemp008 = πF.MakeArgs(2)
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp015, πE = πg.GetAttr(πF, πTemp014, ßplatform, nil); πE != nil {
						continue
					}
					if πTemp013, πE = πg.Eq(πF, πTemp015, ßdarwin.ToObject()); πE != nil {
						continue
					}
					πTemp008[0] = πTemp013
					πTemp008[1] = πg.NewStr("Mac OS X denies the creation of a directory with an invalid utf8 name").ToObject()
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßskipIf, nil); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp014.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πTemp014, πE = πTemp013.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_nonascii_abspath.ToObject(), πTemp014); πE != nil {
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
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("CommonTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCommonTest.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 282: def test_main():
			πF.SetLineno(282)
			πTemp003 = make([]πg.Param, 0)
			πTemp005 = πg.NewFunction(πg.NewCode("test_main", "build/src/__python__/test/test_genericpath.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 283: test_support.run_unittest(GenericTest)
					πF.SetLineno(283)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßGenericTest); πE != nil {
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
			// line 286: if __name__=="__main__":
			πF.SetLineno(286)
		Label1:
			// line 287: test_main()
			πF.SetLineno(287)
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
	πg.RegisterModule("test.test_genericpath", Code)
}
