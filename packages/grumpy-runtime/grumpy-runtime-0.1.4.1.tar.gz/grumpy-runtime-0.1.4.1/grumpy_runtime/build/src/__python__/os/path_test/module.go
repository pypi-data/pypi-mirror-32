package path_test
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/os/path_test.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ß123 := πg.InternStr("123")
		ßRunTests := πg.InternStr("RunTests")
		ßTestAbspath := πg.InternStr("TestAbspath")
		ßTestBasename := πg.InternStr("TestBasename")
		ßTestDirname := πg.InternStr("TestDirname")
		ßTestExists := πg.InternStr("TestExists")
		ßTestIsAbs := πg.InternStr("TestIsAbs")
		ßTestIsDir := πg.InternStr("TestIsDir")
		ßTestIsFile := πg.InternStr("TestIsFile")
		ßTestJoin := πg.InternStr("TestJoin")
		ßTestNormPath := πg.InternStr("TestNormPath")
		ßTestSplit := πg.InternStr("TestSplit")
		ß_AssertEqual := πg.InternStr("_AssertEqual")
		ß__main__ := πg.InternStr("__main__")
		ß__name__ := πg.InternStr("__name__")
		ßa := πg.InternStr("a")
		ßabc := πg.InternStr("abc")
		ßabspath := πg.InternStr("abspath")
		ßb := πg.InternStr("b")
		ßbasename := πg.InternStr("basename")
		ßc := πg.InternStr("c")
		ßdirname := πg.InternStr("dirname")
		ßexists := πg.InternStr("exists")
		ßgetcwd := πg.InternStr("getcwd")
		ßisabs := πg.InternStr("isabs")
		ßisdir := πg.InternStr("isdir")
		ßisfile := πg.InternStr("isfile")
		ßjoin := πg.InternStr("join")
		ßmkdtemp := πg.InternStr("mkdtemp")
		ßmkstemp := πg.InternStr("mkstemp")
		ßnormpath := πg.InternStr("normpath")
		ßos := πg.InternStr("os")
		ßpath := πg.InternStr("path")
		ßremove := πg.InternStr("remove")
		ßrmdir := πg.InternStr("rmdir")
		ßsplit := πg.InternStr("split")
		ßtempfile := πg.InternStr("tempfile")
		ßtype := πg.InternStr("type")
		ßweetest := πg.InternStr("weetest")
		ßx := πg.InternStr("x")
		ßy := πg.InternStr("y")
		ßz := πg.InternStr("z")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
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
		var πTemp012 *πg.Object
		_ = πTemp012
		var πTemp013 *πg.Object
		_ = πTemp013
		var πTemp014 *πg.Object
		_ = πTemp014
		var πTemp015 *πg.Object
		_ = πTemp015
		var πTemp016 bool
		_ = πTemp016
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 17: import os
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: import os.path
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "os.path"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 19: path = os.path
			πF.SetLineno(19)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßpath.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 21: import weetest
			πF.SetLineno(21)
			if πTemp002, πE = πg.ImportModule(πF, "weetest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßweetest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 22: import tempfile
			πF.SetLineno(22)
			if πTemp002, πE = πg.ImportModule(πF, "tempfile"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtempfile.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 25: def _AssertEqual(a, b):
			πF.SetLineno(25)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("_AssertEqual", "build/src/__python__/os/path_test.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 26: assert a == b
					πF.SetLineno(26)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µa, µb); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 27: assert type(a) is type(b)
					πF.SetLineno(27)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp002[0] = µa
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					πTemp002[0] = µb
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001 = πg.GetBool(πTemp004 == πTemp005).ToObject()
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_AssertEqual.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 30: def TestAbspath():
			πF.SetLineno(30)
			πTemp004 = make([]πg.Param, 0)
			πTemp003 = πg.NewFunction(πg.NewCode("TestAbspath", "build/src/__python__/os/path_test.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 31: _AssertEqual(path.abspath('/a/b/c'), '/a/b/c')
					πF.SetLineno(31)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("/a/b/c").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßabspath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp003
					πTemp001[1] = πg.NewStr("/a/b/c").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_AssertEqual); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 32: _AssertEqual(path.abspath(u'/a/b/c'), u'/a/b/c')
					πF.SetLineno(32)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewUnicode("/a/b/c").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßabspath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp003
					πTemp001[1] = πg.NewUnicode("/a/b/c").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_AssertEqual); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 33: _AssertEqual(path.abspath('/a/b/c/'), '/a/b/c')
					πF.SetLineno(33)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("/a/b/c/").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßabspath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp003
					πTemp001[1] = πg.NewStr("/a/b/c").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_AssertEqual); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 34: _AssertEqual(path.abspath(u'/a/b/c/'), u'/a/b/c')
					πF.SetLineno(34)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewUnicode("/a/b/c/").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßabspath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp003
					πTemp001[1] = πg.NewUnicode("/a/b/c").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_AssertEqual); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 35: _AssertEqual(path.abspath('a/b/c'), path.normpath(os.getcwd() + '/a/b/c'))
					πF.SetLineno(35)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("a/b/c").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßabspath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp003
					πTemp002 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßgetcwd, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πTemp004, πg.NewStr("/a/b/c").ToObject()); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßnormpath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[1] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_AssertEqual); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestAbspath.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 38: def TestBasename():
			πF.SetLineno(38)
			πTemp004 = make([]πg.Param, 0)
			πTemp005 = πg.NewFunction(πg.NewCode("TestBasename", "build/src/__python__/os/path_test.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
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
					// line 39: assert path.basename('/a/b/c') == 'c'
					πF.SetLineno(39)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("/a/b/c").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbasename, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, ßc.ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 40: assert path.basename('/a/b/c/') == ''
					πF.SetLineno(40)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("/a/b/c/").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbasename, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, ß.ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestBasename.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 43: def TestDirname():
			πF.SetLineno(43)
			πTemp004 = make([]πg.Param, 0)
			πTemp006 = πg.NewFunction(πg.NewCode("TestDirname", "build/src/__python__/os/path_test.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
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
					// line 44: assert path.dirname('/a/b/c') == '/a/b'
					πF.SetLineno(44)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("/a/b/c").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßdirname, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewStr("/a/b").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 45: assert path.dirname('/a/b/c/') == '/a/b/c'
					πF.SetLineno(45)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("/a/b/c/").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßdirname, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewStr("/a/b/c").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestDirname.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 48: def TestExists():
			πF.SetLineno(48)
			πTemp004 = make([]πg.Param, 0)
			πTemp007 = πg.NewFunction(πg.NewCode("TestExists", "build/src/__python__/os/path_test.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
				var µfile_path *πg.Object = πg.UnboundLocal; _ = µfile_path
				var µdir_path *πg.Object = πg.UnboundLocal; _ = µdir_path
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 bool
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
					// line 49: _, file_path = tempfile.mkstemp()
					πF.SetLineno(49)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkstemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}}}, πTemp001); πE != nil {
						continue
					}
					µ_ = πTemp002
					µfile_path = πTemp003
					// line 50: dir_path = tempfile.mkdtemp()
					πF.SetLineno(50)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkdtemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µdir_path = πTemp001
					// line 51: try:
					πF.SetLineno(51)
					πF.PushCheckpoint(1)
					// line 52: assert path.exists(file_path)
					πF.SetLineno(52)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfile_path, "file_path"); πE != nil {
						continue
					}
					πTemp004[0] = µfile_path
					if πTemp001, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßexists, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 53: assert path.exists(dir_path)
					πF.SetLineno(53)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdir_path, "dir_path"); πE != nil {
						continue
					}
					πTemp004[0] = µdir_path
					if πTemp001, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßexists, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 54: assert not path.exists('path/does/not/exist')
					πF.SetLineno(54)
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("path/does/not/exist").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßexists, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp005).ToObject()
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					goto Label1
				Label1:
					πTemp006, πTemp007 = πF.RestoreExc(nil, nil)
					// line 56: os.remove(file_path)
					πF.SetLineno(56)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfile_path, "file_path"); πE != nil {
						continue
					}
					πTemp004[0] = µfile_path
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßremove, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 57: os.rmdir(dir_path)
					πF.SetLineno(57)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdir_path, "dir_path"); πE != nil {
						continue
					}
					πTemp004[0] = µdir_path
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrmdir, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
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
			if πE = πF.Globals().SetItem(πF, ßTestExists.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 60: def TestIsAbs():
			πF.SetLineno(60)
			πTemp004 = make([]πg.Param, 0)
			πTemp008 = πg.NewFunction(πg.NewCode("TestIsAbs", "build/src/__python__/os/path_test.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 61: assert path.isabs('/abc')
					πF.SetLineno(61)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("/abc").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßisabs, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
						continue
					}
					// line 62: assert not path.isabs('abc/123')
					πF.SetLineno(62)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("abc/123").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßisabs, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp005).ToObject()
					if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestIsAbs.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 65: def TestIsDir():
			πF.SetLineno(65)
			πTemp004 = make([]πg.Param, 0)
			πTemp009 = πg.NewFunction(πg.NewCode("TestIsDir", "build/src/__python__/os/path_test.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
				var µfile_path *πg.Object = πg.UnboundLocal; _ = µfile_path
				var µdir_path *πg.Object = πg.UnboundLocal; _ = µdir_path
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 bool
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
					// line 66: _, file_path = tempfile.mkstemp()
					πF.SetLineno(66)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkstemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}}}, πTemp001); πE != nil {
						continue
					}
					µ_ = πTemp002
					µfile_path = πTemp003
					// line 67: dir_path = tempfile.mkdtemp()
					πF.SetLineno(67)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkdtemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µdir_path = πTemp001
					// line 68: try:
					πF.SetLineno(68)
					πF.PushCheckpoint(1)
					// line 69: assert not path.isdir(file_path)
					πF.SetLineno(69)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfile_path, "file_path"); πE != nil {
						continue
					}
					πTemp004[0] = µfile_path
					if πTemp002, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßisdir, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp005).ToObject()
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 70: assert path.isdir(dir_path)
					πF.SetLineno(70)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdir_path, "dir_path"); πE != nil {
						continue
					}
					πTemp004[0] = µdir_path
					if πTemp001, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßisdir, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 71: assert not path.isdir('path/does/not/exist')
					πF.SetLineno(71)
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("path/does/not/exist").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßisdir, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp005).ToObject()
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					goto Label1
				Label1:
					πTemp006, πTemp007 = πF.RestoreExc(nil, nil)
					// line 73: os.remove(file_path)
					πF.SetLineno(73)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfile_path, "file_path"); πE != nil {
						continue
					}
					πTemp004[0] = µfile_path
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßremove, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 74: os.rmdir(dir_path)
					πF.SetLineno(74)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdir_path, "dir_path"); πE != nil {
						continue
					}
					πTemp004[0] = µdir_path
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrmdir, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
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
			if πE = πF.Globals().SetItem(πF, ßTestIsDir.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 77: def TestIsFile():
			πF.SetLineno(77)
			πTemp004 = make([]πg.Param, 0)
			πTemp010 = πg.NewFunction(πg.NewCode("TestIsFile", "build/src/__python__/os/path_test.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
				var µfile_path *πg.Object = πg.UnboundLocal; _ = µfile_path
				var µdir_path *πg.Object = πg.UnboundLocal; _ = µdir_path
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 bool
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
					// line 78: _, file_path = tempfile.mkstemp()
					πF.SetLineno(78)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkstemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}}}, πTemp001); πE != nil {
						continue
					}
					µ_ = πTemp002
					µfile_path = πTemp003
					// line 79: dir_path = tempfile.mkdtemp()
					πF.SetLineno(79)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkdtemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µdir_path = πTemp001
					// line 80: try:
					πF.SetLineno(80)
					πF.PushCheckpoint(1)
					// line 81: assert path.isfile(file_path)
					πF.SetLineno(81)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfile_path, "file_path"); πE != nil {
						continue
					}
					πTemp004[0] = µfile_path
					if πTemp001, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßisfile, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 82: assert not path.isfile(dir_path)
					πF.SetLineno(82)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdir_path, "dir_path"); πE != nil {
						continue
					}
					πTemp004[0] = µdir_path
					if πTemp002, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßisfile, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp005).ToObject()
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 83: assert not path.isfile('path/does/not/exist')
					πF.SetLineno(83)
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("path/does/not/exist").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßisfile, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp005).ToObject()
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					goto Label1
				Label1:
					πTemp006, πTemp007 = πF.RestoreExc(nil, nil)
					// line 85: os.remove(file_path)
					πF.SetLineno(85)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfile_path, "file_path"); πE != nil {
						continue
					}
					πTemp004[0] = µfile_path
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßremove, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 86: os.rmdir(dir_path)
					πF.SetLineno(86)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdir_path, "dir_path"); πE != nil {
						continue
					}
					πTemp004[0] = µdir_path
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrmdir, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
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
			if πE = πF.Globals().SetItem(πF, ßTestIsFile.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 89: def TestJoin():
			πF.SetLineno(89)
			πTemp004 = make([]πg.Param, 0)
			πTemp011 = πg.NewFunction(πg.NewCode("TestJoin", "build/src/__python__/os/path_test.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
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
					// line 90: assert path.join('') == ''
					πF.SetLineno(90)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = ß.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, ß.ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 91: assert path.join('', '') == ''
					πF.SetLineno(91)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = ß.ToObject()
					πTemp002[1] = ß.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, ß.ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 92: assert path.join('abc') == 'abc'
					πF.SetLineno(92)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = ßabc.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, ßabc.ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 93: assert path.join('abc', '') == 'abc/'
					πF.SetLineno(93)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = ßabc.ToObject()
					πTemp002[1] = ß.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewStr("abc/").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 94: assert path.join('abc', '', '') == 'abc/'
					πF.SetLineno(94)
					πTemp002 = πF.MakeArgs(3)
					πTemp002[0] = ßabc.ToObject()
					πTemp002[1] = ß.ToObject()
					πTemp002[2] = ß.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewStr("abc/").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 95: assert path.join('abc', '', '123') == 'abc/123'
					πF.SetLineno(95)
					πTemp002 = πF.MakeArgs(3)
					πTemp002[0] = ßabc.ToObject()
					πTemp002[1] = ß.ToObject()
					πTemp002[2] = ß123.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewStr("abc/123").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 96: assert path.normpath(path.join('abc', '.', '123')) == 'abc/123'
					πF.SetLineno(96)
					πTemp002 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(3)
					πTemp005[0] = ßabc.ToObject()
					πTemp005[1] = πg.NewStr(".").ToObject()
					πTemp005[2] = ß123.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp002[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßnormpath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewStr("abc/123").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 97: assert path.normpath(path.join('abc', '..', '123')) == '123'
					πF.SetLineno(97)
					πTemp002 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(3)
					πTemp005[0] = ßabc.ToObject()
					πTemp005[1] = πg.NewStr("..").ToObject()
					πTemp005[2] = ß123.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp002[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßnormpath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, ß123.ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 98: assert path.join('/abc', '123') == '/abc/123'
					πF.SetLineno(98)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewStr("/abc").ToObject()
					πTemp002[1] = ß123.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewStr("/abc/123").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 99: assert path.join('abc', '/123') == '/123'
					πF.SetLineno(99)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = ßabc.ToObject()
					πTemp002[1] = πg.NewStr("/123").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewStr("/123").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 100: assert path.join('abc/', '123') == 'abc/123'
					πF.SetLineno(100)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewStr("abc/").ToObject()
					πTemp002[1] = ß123.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewStr("abc/123").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 101: assert path.join('abc', 'x/y/z') == 'abc/x/y/z'
					πF.SetLineno(101)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = ßabc.ToObject()
					πTemp002[1] = πg.NewStr("x/y/z").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewStr("abc/x/y/z").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 102: assert path.join('abc', 'x', 'y', 'z') == 'abc/x/y/z'
					πF.SetLineno(102)
					πTemp002 = πF.MakeArgs(4)
					πTemp002[0] = ßabc.ToObject()
					πTemp002[1] = ßx.ToObject()
					πTemp002[2] = ßy.ToObject()
					πTemp002[3] = ßz.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewStr("abc/x/y/z").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestJoin.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 105: def TestNormPath():
			πF.SetLineno(105)
			πTemp004 = make([]πg.Param, 0)
			πTemp012 = πg.NewFunction(πg.NewCode("TestNormPath", "build/src/__python__/os/path_test.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 106: _AssertEqual(path.normpath('abc/'), 'abc')
					πF.SetLineno(106)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("abc/").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßnormpath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp003
					πTemp001[1] = ßabc.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_AssertEqual); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 107: _AssertEqual(path.normpath('/a//b'), '/a/b')
					πF.SetLineno(107)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("/a//b").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßnormpath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp003
					πTemp001[1] = πg.NewStr("/a/b").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_AssertEqual); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 108: _AssertEqual(path.normpath('abc/../123'), '123')
					πF.SetLineno(108)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("abc/../123").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßnormpath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp003
					πTemp001[1] = ß123.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_AssertEqual); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 109: _AssertEqual(path.normpath('../abc/123'), '../abc/123')
					πF.SetLineno(109)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("../abc/123").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßnormpath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp003
					πTemp001[1] = πg.NewStr("../abc/123").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_AssertEqual); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 110: _AssertEqual(path.normpath('x/y/./z'), 'x/y/z')
					πF.SetLineno(110)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("x/y/./z").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßnormpath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp003
					πTemp001[1] = πg.NewStr("x/y/z").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_AssertEqual); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 111: _AssertEqual(path.normpath(u'abc/'), u'abc')
					πF.SetLineno(111)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewUnicode("abc/").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßnormpath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp003
					πTemp001[1] = πg.NewUnicode("abc").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_AssertEqual); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 112: _AssertEqual(path.normpath(u'/a//b'), u'/a/b')
					πF.SetLineno(112)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewUnicode("/a//b").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßnormpath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp003
					πTemp001[1] = πg.NewUnicode("/a/b").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_AssertEqual); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 113: _AssertEqual(path.normpath(u'abc/../123'), u'123')
					πF.SetLineno(113)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewUnicode("abc/../123").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßnormpath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp003
					πTemp001[1] = πg.NewUnicode("123").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_AssertEqual); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 114: _AssertEqual(path.normpath(u'../abc/123'), u'../abc/123')
					πF.SetLineno(114)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewUnicode("../abc/123").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßnormpath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp003
					πTemp001[1] = πg.NewUnicode("../abc/123").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_AssertEqual); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 115: _AssertEqual(path.normpath(u'x/y/./z'), u'x/y/z')
					πF.SetLineno(115)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewUnicode("x/y/./z").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßnormpath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp003
					πTemp001[1] = πg.NewUnicode("x/y/z").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_AssertEqual); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestNormPath.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 118: def TestSplit():
			πF.SetLineno(118)
			πTemp004 = make([]πg.Param, 0)
			πTemp013 = πg.NewFunction(πg.NewCode("TestSplit", "build/src/__python__/os/path_test.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
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
					// line 119: assert path.split('a/b') == ('a', 'b')
					πF.SetLineno(119)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("a/b").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp004 = πg.NewTuple2(ßa.ToObject(), ßb.ToObject()).ToObject()
					if πTemp001, πE = πg.Eq(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 120: assert path.split('a/b/') == ('a/b', '')
					πF.SetLineno(120)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("a/b/").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp004 = πg.NewTuple2(πg.NewStr("a/b").ToObject(), ß.ToObject()).ToObject()
					if πTemp001, πE = πg.Eq(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 121: assert path.split('a/') == ('a', '')
					πF.SetLineno(121)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("a/").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp004 = πg.NewTuple2(ßa.ToObject(), ß.ToObject()).ToObject()
					if πTemp001, πE = πg.Eq(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 122: assert path.split('a') == ('', 'a')
					πF.SetLineno(122)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = ßa.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp004 = πg.NewTuple2(ß.ToObject(), ßa.ToObject()).ToObject()
					if πTemp001, πE = πg.Eq(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 123: assert path.split('/') == ('/', '')
					πF.SetLineno(123)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("/").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp004 = πg.NewTuple2(πg.NewStr("/").ToObject(), ß.ToObject()).ToObject()
					if πTemp001, πE = πg.Eq(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 124: assert path.split('/a/./b') == ('/a/.', 'b')
					πF.SetLineno(124)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("/a/./b").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp004 = πg.NewTuple2(πg.NewStr("/a/.").ToObject(), ßb.ToObject()).ToObject()
					if πTemp001, πE = πg.Eq(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestSplit.ToObject(), πTemp013); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp014, πE = πg.Eq(πF, πTemp015, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp016, πE = πg.IsTrue(πF, πTemp014); πE != nil {
				continue
			}
			if πTemp016 {
				goto Label1
			}
			goto Label2
			// line 127: if __name__ == '__main__':
			πF.SetLineno(127)
		Label1:
			// line 128: weetest.RunTests()
			πF.SetLineno(128)
			if πTemp014, πE = πg.ResolveGlobal(πF, ßweetest); πE != nil {
				continue
			}
			if πTemp015, πE = πg.GetAttr(πF, πTemp014, ßRunTests, nil); πE != nil {
				continue
			}
			if πTemp014, πE = πTemp015.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("os.path_test", Code)
}
