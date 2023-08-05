package test_dircache
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/test/test_dircache.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßA := πg.InternStr("A")
		ßDircacheTests := πg.InternStr("DircacheTests")
		ßOSError := πg.InternStr("OSError")
		ßTestCase := πg.InternStr("TestCase")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_nonexistent := πg.InternStr("_nonexistent")
		ßannotate := πg.InternStr("annotate")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertRaises := πg.InternStr("assertRaises")
		ßassertTrue := πg.InternStr("assertTrue")
		ßclose := πg.InternStr("close")
		ßdelTemp := πg.InternStr("delTemp")
		ßdircache := πg.InternStr("dircache")
		ßisdir := πg.InternStr("isdir")
		ßjoin := πg.InternStr("join")
		ßlistdir := πg.InternStr("listdir")
		ßmkdir := πg.InternStr("mkdir")
		ßmkdirTemp := πg.InternStr("mkdirTemp")
		ßmkdtemp := πg.InternStr("mkdtemp")
		ßopen := πg.InternStr("open")
		ßos := πg.InternStr("os")
		ßos2 := πg.InternStr("os2")
		ßpath := πg.InternStr("path")
		ßplatform := πg.InternStr("platform")
		ßreset := πg.InternStr("reset")
		ßrmdir := πg.InternStr("rmdir")
		ßrun_unittest := πg.InternStr("run_unittest")
		ßsetUp := πg.InternStr("setUp")
		ßsleep := πg.InternStr("sleep")
		ßsys := πg.InternStr("sys")
		ßtearDown := πg.InternStr("tearDown")
		ßtempdir := πg.InternStr("tempdir")
		ßtempfile := πg.InternStr("tempfile")
		ßtest1 := πg.InternStr("test1")
		ßtest2 := πg.InternStr("test2")
		ßtest_annotate := πg.InternStr("test_annotate")
		ßtest_listdir := πg.InternStr("test_listdir")
		ßtest_main := πg.InternStr("test_main")
		ßtest_nonexistent := πg.InternStr("test_nonexistent")
		ßtime := πg.InternStr("time")
		ßunittest := πg.InternStr("unittest")
		ßunlink := πg.InternStr("unlink")
		ßw := πg.InternStr("w")
		ßwin := πg.InternStr("win")
		ßwriteTemp := πg.InternStr("writeTemp")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.Dict
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
			// line 7: from test.test_support import run_unittest # , import_module
			πF.SetLineno(7)
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
			// line 9: import dircache
			πF.SetLineno(9)
			if πTemp002, πE = πg.ImportModule(πF, "dircache"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßdircache.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 10: import os, time, sys, tempfile
			πF.SetLineno(10)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "time"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtime.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "tempfile"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtempfile.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 13: class DircacheTests(unittest.TestCase):
			πF.SetLineno(13)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßTestCase, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("DircacheTests", "build/src/__python__/test/test_dircache.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
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
					// line 14: def setUp(self):
					πF.SetLineno(14)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("setUp", "build/src/__python__/test/test_dircache.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 15: self.tempdir = tempfile.mkdtemp()
							πF.SetLineno(15)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkdtemp, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtempdir, πTemp002); πE != nil {
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
					// line 17: def tearDown(self):
					πF.SetLineno(17)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("tearDown", "build/src/__python__/test/test_dircache.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfname *πg.Object = πg.UnboundLocal; _ = µfname
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
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtempdir, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßlistdir, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								µfname = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 19: self.delTemp(fname)
							πF.SetLineno(19)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfname, "fname"); πE != nil {
								continue
							}
							πTemp002[0] = µfname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdelTemp, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 20: os.rmdir(self.tempdir)
							πF.SetLineno(20)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtempdir, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßrmdir, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtearDown.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 22: def writeTemp(self, fname):
					πF.SetLineno(22)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "fname", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("writeTemp", "build/src/__python__/test/test_dircache.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfname *πg.Object = πArgs[1]; _ = µfname
						var µf *πg.Object = πg.UnboundLocal; _ = µf
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
							// line 23: f = open(os.path.join(self.tempdir, fname), 'w')
							πF.SetLineno(23)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtempdir, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µfname, "fname"); πE != nil {
								continue
							}
							πTemp002[1] = µfname
							if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = ßw.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µf = πTemp004
							// line 24: f.close()
							πF.SetLineno(24)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßwriteTemp.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 26: def mkdirTemp(self, fname):
					πF.SetLineno(26)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "fname", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("mkdirTemp", "build/src/__python__/test/test_dircache.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfname *πg.Object = πArgs[1]; _ = µfname
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
							// line 27: os.mkdir(os.path.join(self.tempdir, fname))
							πF.SetLineno(27)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtempdir, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µfname, "fname"); πE != nil {
								continue
							}
							πTemp002[1] = µfname
							if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßmkdirTemp.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 29: def delTemp(self, fname):
					πF.SetLineno(29)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "fname", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("delTemp", "build/src/__python__/test/test_dircache.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfname *πg.Object = πArgs[1]; _ = µfname
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
							// line 30: fname = os.path.join(self.tempdir, fname)
							πF.SetLineno(30)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtempdir, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µfname, "fname"); πE != nil {
								continue
							}
							πTemp001[1] = µfname
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µfname = πTemp003
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfname, "fname"); πE != nil {
								continue
							}
							πTemp001[0] = µfname
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßisdir, nil); πE != nil {
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
							// line 31: if os.path.isdir(fname):
							πF.SetLineno(31)
						Label1:
							// line 32: os.rmdir(fname)
							πF.SetLineno(32)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfname, "fname"); πE != nil {
								continue
							}
							πTemp001[0] = µfname
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
							goto Label3
						Label2:
							// line 34: os.unlink(fname)
							πF.SetLineno(34)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfname, "fname"); πE != nil {
								continue
							}
							πTemp001[0] = µfname
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunlink, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßdelTemp.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 36: def test_listdir(self):
					πF.SetLineno(36)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_listdir", "build/src/__python__/test/test_dircache.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µentries *πg.Object = πg.UnboundLocal; _ = µentries
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
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
							default: panic("unexpected function state")
							}
							// line 38: entries = dircache.listdir(self.tempdir)
							πF.SetLineno(38)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtempdir, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdircache); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlistdir, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µentries = πTemp002
							// line 39: self.assertEqual(entries, [])
							πF.SetLineno(39)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µentries, "entries"); πE != nil {
								continue
							}
							πTemp001[0] = µentries
							πTemp004 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp004...).ToObject()
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
							// line 42: self.assertTrue(dircache.listdir(self.tempdir) is entries)
							πF.SetLineno(42)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtempdir, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßdircache); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßlistdir, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.CheckLocal(πF, µentries, "entries"); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp003 == µentries).ToObject()
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
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(3).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßplatform, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(ßwin.ToObject(), ßos2.ToObject()).ToObject()
							if πTemp008, πE = πg.Contains(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp008).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label1
							}
							goto Label2
							// line 49: if sys.platform[:3] not in ('win', 'os2'):
							πF.SetLineno(49)
						Label1:
							// line 53: time.sleep(1)
							πF.SetLineno(53)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(1).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsleep, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 54: self.writeTemp("test1")
							πF.SetLineno(54)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßtest1.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßwriteTemp, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 55: entries = dircache.listdir(self.tempdir)
							πF.SetLineno(55)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtempdir, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdircache); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlistdir, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µentries = πTemp002
							// line 56: self.assertEqual(entries, ['test1'])
							πF.SetLineno(56)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µentries, "entries"); πE != nil {
								continue
							}
							πTemp001[0] = µentries
							πTemp004 = make([]*πg.Object, 1)
							πTemp004[0] = ßtest1.ToObject()
							πTemp002 = πg.NewList(πTemp004...).ToObject()
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
							// line 57: self.assertTrue(dircache.listdir(self.tempdir) is entries)
							πF.SetLineno(57)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtempdir, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßdircache); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßlistdir, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.CheckLocal(πF, µentries, "entries"); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp003 == µentries).ToObject()
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
							goto Label2
						Label2:
							// line 60: self.assertRaises(OSError, dircache.listdir, self.tempdir+"_nonexistent")
							πF.SetLineno(60)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdircache); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlistdir, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtempdir, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, ß_nonexistent.ToObject()); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_listdir.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 62: def test_annotate(self):
					πF.SetLineno(62)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("test_annotate", "build/src/__python__/test/test_dircache.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlst *πg.Object = πg.UnboundLocal; _ = µlst
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
							// line 63: self.writeTemp("test2")
							πF.SetLineno(63)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßtest2.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßwriteTemp, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 64: self.mkdirTemp("A")
							πF.SetLineno(64)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßA.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmkdirTemp, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 65: lst = ['A', 'test2', 'test_nonexistent']
							πF.SetLineno(65)
							πTemp001 = make([]*πg.Object, 3)
							πTemp001[0] = ßA.ToObject()
							πTemp001[1] = ßtest2.ToObject()
							πTemp001[2] = ßtest_nonexistent.ToObject()
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µlst = πTemp002
							// line 66: dircache.annotate(self.tempdir, lst)
							πF.SetLineno(66)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtempdir, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
								continue
							}
							πTemp001[1] = µlst
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdircache); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßannotate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 67: self.assertEqual(lst, ['A/', 'test2', 'test_nonexistent'])
							πF.SetLineno(67)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
								continue
							}
							πTemp001[0] = µlst
							πTemp004 = make([]*πg.Object, 3)
							πTemp004[0] = πg.NewStr("A/").ToObject()
							πTemp004[1] = ßtest2.ToObject()
							πTemp004[2] = ßtest_nonexistent.ToObject()
							πTemp002 = πg.NewList(πTemp004...).ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_annotate.ToObject(), πTemp008); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("DircacheTests").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDircacheTests.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 70: def test_main():
			πF.SetLineno(70)
			πTemp007 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("test_main", "build/src/__python__/test/test_dircache.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					default: panic("unexpected function state")
					}
					// line 71: try:
					πF.SetLineno(71)
					πF.PushCheckpoint(1)
					// line 72: run_unittest(DircacheTests)
					πF.SetLineno(72)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßDircacheTests); πE != nil {
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
					πF.PopCheckpoint()
					goto Label1
				Label1:
					πTemp004, πTemp005 = πF.RestoreExc(nil, nil)
					// line 74: dircache.reset()
					πF.SetLineno(74)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdircache); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßreset, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp004 != nil {
						πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
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
			if πE = πF.Globals().SetItem(πF, ßtest_main.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp003, πE = πg.Eq(πF, πTemp005, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
				continue
			}
			if πTemp008 {
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
			if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("test.test_dircache", Code)
}
