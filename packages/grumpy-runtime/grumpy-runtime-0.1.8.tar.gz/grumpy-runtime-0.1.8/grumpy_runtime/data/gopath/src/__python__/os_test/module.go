package os_test
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/os_test.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßAssertionError := πg.InternStr("AssertionError")
		ßHOME := πg.InternStr("HOME")
		ßOSError := πg.InternStr("OSError")
		ßRunTests := πg.InternStr("RunTests")
		ßS_IMODE := πg.InternStr("S_IMODE")
		ßS_ISDIR := πg.InternStr("S_ISDIR")
		ßTestChdirAndGetCwd := πg.InternStr("TestChdirAndGetCwd")
		ßTestChmod := πg.InternStr("TestChmod")
		ßTestChmodOSError := πg.InternStr("TestChmodOSError")
		ßTestClose := πg.InternStr("TestClose")
		ßTestCloseOSError := πg.InternStr("TestCloseOSError")
		ßTestEnviron := πg.InternStr("TestEnviron")
		ßTestFDOpen := πg.InternStr("TestFDOpen")
		ßTestFDOpenOSError := πg.InternStr("TestFDOpenOSError")
		ßTestMkdir := πg.InternStr("TestMkdir")
		ßTestPopenRead := πg.InternStr("TestPopenRead")
		ßTestPopenWrite := πg.InternStr("TestPopenWrite")
		ßTestRemove := πg.InternStr("TestRemove")
		ßTestRemoveDir := πg.InternStr("TestRemoveDir")
		ßTestRemoveNoExist := πg.InternStr("TestRemoveNoExist")
		ßTestRmDir := πg.InternStr("TestRmDir")
		ßTestRmDirFile := πg.InternStr("TestRmDirFile")
		ßTestRmDirNoExist := πg.InternStr("TestRmDirNoExist")
		ßTestStatDir := πg.InternStr("TestStatDir")
		ßTestStatFile := πg.InternStr("TestStatFile")
		ßTestStatNoExist := πg.InternStr("TestStatNoExist")
		ßTestWaitPid := πg.InternStr("TestWaitPid")
		ßWNOHANG := πg.InternStr("WNOHANG")
		ß__main__ := πg.InternStr("__main__")
		ß__name__ := πg.InternStr("__name__")
		ßcat := πg.InternStr("cat")
		ßchdir := πg.InternStr("chdir")
		ßchmod := πg.InternStr("chmod")
		ßclose := πg.InternStr("close")
		ßenviron := πg.InternStr("environ")
		ßfdopen := πg.InternStr("fdopen")
		ßfoobar := πg.InternStr("foobar")
		ßfoobarqux := πg.InternStr("foobarqux")
		ßgetcwd := πg.InternStr("getcwd")
		ßlower := πg.InternStr("lower")
		ßmkdir := πg.InternStr("mkdir")
		ßmkdtemp := πg.InternStr("mkdtemp")
		ßmkstemp := πg.InternStr("mkstemp")
		ßopen := πg.InternStr("open")
		ßos := πg.InternStr("os")
		ßpopen := πg.InternStr("popen")
		ßqux := πg.InternStr("qux")
		ßread := πg.InternStr("read")
		ßremove := πg.InternStr("remove")
		ßrmdir := πg.InternStr("rmdir")
		ßst_mode := πg.InternStr("st_mode")
		ßst_mtime := πg.InternStr("st_mtime")
		ßst_size := πg.InternStr("st_size")
		ßstat := πg.InternStr("stat")
		ßstr := πg.InternStr("str")
		ßtempfile := πg.InternStr("tempfile")
		ßtime := πg.InternStr("time")
		ßw := πg.InternStr("w")
		ßwaitpid := πg.InternStr("waitpid")
		ßweetest := πg.InternStr("weetest")
		ßwrite := πg.InternStr("write")
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
		var πTemp018 *πg.Object
		_ = πTemp018
		var πTemp019 *πg.Object
		_ = πTemp019
		var πTemp020 *πg.Object
		_ = πTemp020
		var πTemp021 *πg.Object
		_ = πTemp021
		var πTemp022 *πg.Object
		_ = πTemp022
		var πTemp023 *πg.Object
		_ = πTemp023
		var πTemp024 *πg.Object
		_ = πTemp024
		var πTemp025 *πg.Object
		_ = πTemp025
		var πTemp026 bool
		_ = πTemp026
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 15: import os
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 16: import stat
			πF.SetLineno(16)
			if πTemp002, πE = πg.ImportModule(πF, "stat"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßstat.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 17: import time
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "time"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtime.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: import tempfile
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "tempfile"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtempfile.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 20: import weetest
			πF.SetLineno(20)
			if πTemp002, πE = πg.ImportModule(πF, "weetest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßweetest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 23: def TestChdirAndGetCwd():
			πF.SetLineno(23)
			πTemp003 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("TestChdirAndGetCwd", "build/src/__python__/os_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var µtempdir *πg.Object = πg.UnboundLocal; _ = µtempdir
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
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
					// line 24: path = os.getcwd()
					πF.SetLineno(24)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßgetcwd, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µpath = πTemp001
					// line 25: os.chdir('.')
					πF.SetLineno(25)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr(".").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßchdir, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 26: assert os.getcwd() == path
					πF.SetLineno(26)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßgetcwd, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, πTemp002, µpath); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 27: tempdir = tempfile.mkdtemp()
					πF.SetLineno(27)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkdtemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µtempdir = πTemp001
					// line 28: try:
					πF.SetLineno(28)
					πF.PushCheckpoint(1)
					// line 29: os.chdir(tempdir)
					πF.SetLineno(29)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp003[0] = µtempdir
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßchdir, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 30: assert tempdir in os.getcwd()
					πF.SetLineno(30)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßgetcwd, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, πTemp002, µtempdir); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp005).ToObject()
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					goto Label1
				Label1:
					πTemp006, πTemp007 = πF.RestoreExc(nil, nil)
					// line 32: os.chdir(path)
					πF.SetLineno(32)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp003[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßchdir, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 33: os.rmdir(tempdir)
					πF.SetLineno(33)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp003[0] = µtempdir
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrmdir, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 34: assert os.getcwd() == path
					πF.SetLineno(34)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßgetcwd, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, πTemp002, µpath); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßTestChdirAndGetCwd.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 37: def TestChmod():
			πF.SetLineno(37)
			πTemp003 = make([]πg.Param, 0)
			πTemp004 = πg.NewFunction(πg.NewCode("TestChmod", "build/src/__python__/os_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfd *πg.Object = πg.UnboundLocal; _ = µfd
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var µmode *πg.Object = πg.UnboundLocal; _ = µmode
				var πTemp001 *πg.Object
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
					// line 38: fd, path = tempfile.mkstemp()
					πF.SetLineno(38)
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
					µfd = πTemp002
					µpath = πTemp003
					// line 39: os.close(fd)
					πF.SetLineno(39)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp004[0] = µfd
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßclose, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 40: os.chmod(path, 0o644)
					πF.SetLineno(40)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp004[0] = µpath
					πTemp004[1] = πg.NewInt(420).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßchmod, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 41: mode = os.stat(path).st_mode & 0o777
					πF.SetLineno(41)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp004[0] = µpath
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstat, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßst_mode, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.And(πF, πTemp003, πg.NewInt(511).ToObject()); πE != nil {
						continue
					}
					µmode = πTemp001
					// line 42: os.remove(path)
					πF.SetLineno(42)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp004[0] = µpath
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
					// line 43: assert mode == 0o644
					πF.SetLineno(43)
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µmode, πg.NewInt(420).ToObject()); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestChmod.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 46: def TestChmodOSError():
			πF.SetLineno(46)
			πTemp003 = make([]πg.Param, 0)
			πTemp005 = πg.NewFunction(πg.NewCode("TestChmodOSError", "build/src/__python__/os_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtempdir *πg.Object = πg.UnboundLocal; _ = µtempdir
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
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
					// line 47: tempdir = tempfile.mkdtemp()
					πF.SetLineno(47)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkdtemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µtempdir = πTemp001
					// line 48: try:
					πF.SetLineno(48)
					πF.PushCheckpoint(2)
					// line 49: os.chmod(tempdir + '/DoesNotExist', 0o644)
					πF.SetLineno(49)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µtempdir, πg.NewStr("/DoesNotExist").ToObject()); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					πTemp003[1] = πg.NewInt(420).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßchmod, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πF.PopCheckpoint()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					// line 53: raise AssertionError
					πF.SetLineno(53)
					πE = πF.Raise(πTemp001, nil, nil)
					continue
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 50: except OSError:
					πF.SetLineno(50)
				Label3:
					// line 51: pass
					πF.SetLineno(51)
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
			if πE = πF.Globals().SetItem(πF, ßTestChmodOSError.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 56: def TestClose():
			πF.SetLineno(56)
			πTemp003 = make([]πg.Param, 0)
			πTemp006 = πg.NewFunction(πg.NewCode("TestClose", "build/src/__python__/os_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfd *πg.Object = πg.UnboundLocal; _ = µfd
				var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
				var πTemp001 *πg.Object
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
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 57: fd, _ = tempfile.mkstemp()
					πF.SetLineno(57)
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
					µfd = πTemp002
					µ_ = πTemp003
					// line 58: os.close(fd)
					πF.SetLineno(58)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp004[0] = µfd
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßclose, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 59: try:
					πF.SetLineno(59)
					πF.PushCheckpoint(2)
					// line 60: os.fdopen(fd)
					πF.SetLineno(60)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp004[0] = µfd
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßfdopen, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πF.PopCheckpoint()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					// line 64: raise AssertionError
					πF.SetLineno(64)
					πE = πF.Raise(πTemp001, nil, nil)
					continue
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp005, πTemp006 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label3
					}
					πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
					continue
					// line 61: except OSError:
					πF.SetLineno(61)
				Label3:
					// line 62: pass
					πF.SetLineno(62)
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
			if πE = πF.Globals().SetItem(πF, ßTestClose.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 67: def TestCloseOSError():
			πF.SetLineno(67)
			πTemp003 = make([]πg.Param, 0)
			πTemp007 = πg.NewFunction(πg.NewCode("TestCloseOSError", "build/src/__python__/os_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfd *πg.Object = πg.UnboundLocal; _ = µfd
				var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
				var πTemp001 *πg.Object
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
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 68: fd, _ = tempfile.mkstemp()
					πF.SetLineno(68)
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
					µfd = πTemp002
					µ_ = πTemp003
					// line 69: os.close(fd)
					πF.SetLineno(69)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp004[0] = µfd
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßclose, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 70: try:
					πF.SetLineno(70)
					πF.PushCheckpoint(2)
					// line 71: os.close(fd)
					πF.SetLineno(71)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp004[0] = µfd
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßclose, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πF.PopCheckpoint()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					// line 75: raise AssertionError
					πF.SetLineno(75)
					πE = πF.Raise(πTemp001, nil, nil)
					continue
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp005, πTemp006 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label3
					}
					πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
					continue
					// line 72: except OSError:
					πF.SetLineno(72)
				Label3:
					// line 73: pass
					πF.SetLineno(73)
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
			if πE = πF.Globals().SetItem(πF, ßTestCloseOSError.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 78: def TestEnviron():
			πF.SetLineno(78)
			πTemp003 = make([]πg.Param, 0)
			πTemp008 = πg.NewFunction(πg.NewCode("TestEnviron", "build/src/__python__/os_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
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
					// line 79: assert 'HOME' in os.environ
					πF.SetLineno(79)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßenviron, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, πTemp003, ßHOME.ToObject()); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp004).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßTestEnviron.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 82: def TestFDOpen():
			πF.SetLineno(82)
			πTemp003 = make([]πg.Param, 0)
			πTemp009 = πg.NewFunction(πg.NewCode("TestFDOpen", "build/src/__python__/os_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfd *πg.Object = πg.UnboundLocal; _ = µfd
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var µf *πg.Object = πg.UnboundLocal; _ = µf
				var µcontents *πg.Object = πg.UnboundLocal; _ = µcontents
				var πTemp001 *πg.Object
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
					// line 83: fd, path = tempfile.mkstemp()
					πF.SetLineno(83)
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
					µfd = πTemp002
					µpath = πTemp003
					// line 84: f = os.fdopen(fd, 'w')
					πF.SetLineno(84)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp004[0] = µfd
					πTemp004[1] = ßw.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßfdopen, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µf = πTemp001
					// line 85: f.write('foobar')
					πF.SetLineno(85)
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = ßfoobar.ToObject()
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 86: f.close()
					πF.SetLineno(86)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 87: f = open(path)
					πF.SetLineno(87)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp004[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µf = πTemp002
					// line 88: contents = f.read()
					πF.SetLineno(88)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßread, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µcontents = πTemp002
					// line 89: f.close()
					πF.SetLineno(89)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 90: assert contents == 'foobar', contents
					πF.SetLineno(90)
					if πE = πg.CheckLocal(πF, µcontents, "contents"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcontents, "contents"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µcontents, ßfoobar.ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, µcontents); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestFDOpen.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 93: def TestFDOpenOSError():
			πF.SetLineno(93)
			πTemp003 = make([]πg.Param, 0)
			πTemp010 = πg.NewFunction(πg.NewCode("TestFDOpenOSError", "build/src/__python__/os_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfd *πg.Object = πg.UnboundLocal; _ = µfd
				var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
				var πTemp001 *πg.Object
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
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 94: fd, _ = tempfile.mkstemp()
					πF.SetLineno(94)
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
					µfd = πTemp002
					µ_ = πTemp003
					// line 95: os.close(fd)
					πF.SetLineno(95)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp004[0] = µfd
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßclose, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 96: try:
					πF.SetLineno(96)
					πF.PushCheckpoint(2)
					// line 97: os.fdopen(fd)
					πF.SetLineno(97)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp004[0] = µfd
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßfdopen, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πF.PopCheckpoint()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					// line 101: raise AssertionError
					πF.SetLineno(101)
					πE = πF.Raise(πTemp001, nil, nil)
					continue
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp005, πTemp006 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label3
					}
					πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
					continue
					// line 98: except OSError:
					πF.SetLineno(98)
				Label3:
					// line 99: pass
					πF.SetLineno(99)
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
			if πE = πF.Globals().SetItem(πF, ßTestFDOpenOSError.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 104: def TestMkdir():
			πF.SetLineno(104)
			πTemp003 = make([]πg.Param, 0)
			πTemp011 = πg.NewFunction(πg.NewCode("TestMkdir", "build/src/__python__/os_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
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
					case 4: goto Label4
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 105: path = 'foobarqux'
					πF.SetLineno(105)
					µpath = ßfoobarqux.ToObject()
					// line 106: try:
					πF.SetLineno(106)
					πF.PushCheckpoint(2)
					// line 107: os.stat(path)
					πF.SetLineno(107)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp001[0] = µpath
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstat, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.PopCheckpoint()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					// line 111: raise AssertionError
					πF.SetLineno(111)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
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
					// line 108: except OSError:
					πF.SetLineno(108)
				Label3:
					// line 109: pass
					πF.SetLineno(109)
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 112: try:
					πF.SetLineno(112)
					πF.PushCheckpoint(4)
					πF.PushCheckpoint(5)
					// line 113: os.mkdir(path)
					πF.SetLineno(113)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp001[0] = µpath
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmkdir, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 114: assert stat.S_ISDIR(os.stat(path).st_mode)
					πF.SetLineno(114)
					πTemp001 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp007[0] = µpath
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstat, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßst_mode, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßS_ISDIR, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					πF.PopCheckpoint()
					goto Label4
				Label5:
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
						goto Label6
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 115: except OSError:
					πF.SetLineno(115)
				Label6:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					// line 116: raise AssertionError
					πF.SetLineno(116)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					πF.PopCheckpoint()
					goto Label4
				Label4:
					πTemp004, πTemp005 = πF.RestoreExc(nil, nil)
					// line 118: os.rmdir(path)
					πF.SetLineno(118)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp001[0] = µpath
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
			if πE = πF.Globals().SetItem(πF, ßTestMkdir.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 121: def TestPopenRead():
			πF.SetLineno(121)
			πTemp003 = make([]πg.Param, 0)
			πTemp012 = πg.NewFunction(πg.NewCode("TestPopenRead", "build/src/__python__/os_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µf *πg.Object = πg.UnboundLocal; _ = µf
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.BaseException
				_ = πTemp005
				var πTemp006 *πg.Traceback
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					default: panic("unexpected function state")
					}
					// line 122: f = os.popen('qux')
					πF.SetLineno(122)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ßqux.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpopen, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µf = πTemp002
					// line 123: assert f.close() == 32512
					πF.SetLineno(123)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp004, πg.NewInt(32512).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
						continue
					}
					// line 124: f = os.popen('echo hello')
					πF.SetLineno(124)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("echo hello").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpopen, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µf = πTemp002
					// line 125: try:
					πF.SetLineno(125)
					πF.PushCheckpoint(1)
					// line 126: assert f.read() == 'hello\n'
					πF.SetLineno(126)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µf, ßread, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp004, πg.NewStr("hello\n").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					goto Label1
				Label1:
					πTemp005, πTemp006 = πF.RestoreExc(nil, nil)
					// line 128: assert f.close() == 0
					πF.SetLineno(128)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp004, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßTestPopenRead.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 131: def TestPopenWrite():
			πF.SetLineno(131)
			πTemp003 = make([]πg.Param, 0)
			πTemp013 = πg.NewFunction(πg.NewCode("TestPopenWrite", "build/src/__python__/os_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µf *πg.Object = πg.UnboundLocal; _ = µf
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
					// line 134: f = os.popen('cat', 'w')
					πF.SetLineno(134)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = ßcat.ToObject()
					πTemp001[1] = ßw.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpopen, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µf = πTemp002
					// line 135: f.write('popen write\n')
					πF.SetLineno(135)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("popen write\n").ToObject()
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
					// line 136: f.close()
					πF.SetLineno(136)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestPopenWrite.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 139: def TestRemove():
			πF.SetLineno(139)
			πTemp003 = make([]πg.Param, 0)
			πTemp014 = πg.NewFunction(πg.NewCode("TestRemove", "build/src/__python__/os_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfd *πg.Object = πg.UnboundLocal; _ = µfd
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var πTemp001 *πg.Object
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
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 140: fd, path = tempfile.mkstemp()
					πF.SetLineno(140)
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
					µfd = πTemp002
					µpath = πTemp003
					// line 141: os.close(fd)
					πF.SetLineno(141)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp004[0] = µfd
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßclose, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 142: os.stat(path)
					πF.SetLineno(142)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp004[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstat, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 143: os.remove(path)
					πF.SetLineno(143)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp004[0] = µpath
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
					// line 144: try:
					πF.SetLineno(144)
					πF.PushCheckpoint(2)
					// line 145: os.stat(path)
					πF.SetLineno(145)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp004[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstat, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πF.PopCheckpoint()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					// line 149: raise AssertionError
					πF.SetLineno(149)
					πE = πF.Raise(πTemp001, nil, nil)
					continue
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp005, πTemp006 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label3
					}
					πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
					continue
					// line 146: except OSError:
					πF.SetLineno(146)
				Label3:
					// line 147: pass
					πF.SetLineno(147)
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
			if πE = πF.Globals().SetItem(πF, ßTestRemove.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 152: def TestRemoveNoExist():
			πF.SetLineno(152)
			πTemp003 = make([]πg.Param, 0)
			πTemp015 = πg.NewFunction(πg.NewCode("TestRemoveNoExist", "build/src/__python__/os_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
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
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 153: path = tempfile.mkdtemp()
					πF.SetLineno(153)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkdtemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µpath = πTemp001
					// line 154: try:
					πF.SetLineno(154)
					πF.PushCheckpoint(1)
					πF.PushCheckpoint(2)
					// line 155: os.remove(path + '/nonexistent')
					πF.SetLineno(155)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µpath, πg.NewStr("/nonexistent").ToObject()); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßremove, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πF.PopCheckpoint()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					// line 159: raise AssertionError
					πF.SetLineno(159)
					πE = πF.Raise(πTemp001, nil, nil)
					continue
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 156: except OSError:
					πF.SetLineno(156)
				Label3:
					// line 157: pass
					πF.SetLineno(157)
					πF.RestoreExc(nil, nil)
					πF.PopCheckpoint()
					goto Label1
				Label1:
					πTemp004, πTemp005 = πF.RestoreExc(nil, nil)
					// line 161: os.rmdir(path)
					πF.SetLineno(161)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp003[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrmdir, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
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
			if πE = πF.Globals().SetItem(πF, ßTestRemoveNoExist.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 164: def TestRemoveDir():
			πF.SetLineno(164)
			πTemp003 = make([]πg.Param, 0)
			πTemp016 = πg.NewFunction(πg.NewCode("TestRemoveDir", "build/src/__python__/os_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
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
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 165: path = tempfile.mkdtemp()
					πF.SetLineno(165)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkdtemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µpath = πTemp001
					// line 166: try:
					πF.SetLineno(166)
					πF.PushCheckpoint(1)
					πF.PushCheckpoint(2)
					// line 167: os.remove(path)
					πF.SetLineno(167)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp003[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßremove, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πF.PopCheckpoint()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					// line 171: raise AssertionError
					πF.SetLineno(171)
					πE = πF.Raise(πTemp001, nil, nil)
					continue
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 168: except OSError:
					πF.SetLineno(168)
				Label3:
					// line 169: pass
					πF.SetLineno(169)
					πF.RestoreExc(nil, nil)
					πF.PopCheckpoint()
					goto Label1
				Label1:
					πTemp004, πTemp005 = πF.RestoreExc(nil, nil)
					// line 173: os.rmdir(path)
					πF.SetLineno(173)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp003[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrmdir, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
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
			if πE = πF.Globals().SetItem(πF, ßTestRemoveDir.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 176: def TestRmDir():
			πF.SetLineno(176)
			πTemp003 = make([]πg.Param, 0)
			πTemp017 = πg.NewFunction(πg.NewCode("TestRmDir", "build/src/__python__/os_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
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
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 177: path = tempfile.mkdtemp()
					πF.SetLineno(177)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkdtemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µpath = πTemp001
					// line 178: assert stat.S_ISDIR(os.stat(path).st_mode)
					πF.SetLineno(178)
					πTemp003 = πF.MakeArgs(1)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp004[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstat, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßst_mode, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßS_ISDIR, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 179: os.rmdir(path)
					πF.SetLineno(179)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp003[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrmdir, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 180: try:
					πF.SetLineno(180)
					πF.PushCheckpoint(2)
					// line 181: os.stat(path)
					πF.SetLineno(181)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp003[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstat, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πF.PopCheckpoint()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					// line 185: raise AssertionError
					πF.SetLineno(185)
					πE = πF.Raise(πTemp001, nil, nil)
					continue
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp005, πTemp006 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label3
					}
					πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
					continue
					// line 182: except OSError:
					πF.SetLineno(182)
				Label3:
					// line 183: pass
					πF.SetLineno(183)
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
			if πE = πF.Globals().SetItem(πF, ßTestRmDir.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 188: def TestRmDirNoExist():
			πF.SetLineno(188)
			πTemp003 = make([]πg.Param, 0)
			πTemp018 = πg.NewFunction(πg.NewCode("TestRmDirNoExist", "build/src/__python__/os_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
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
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 189: path = tempfile.mkdtemp()
					πF.SetLineno(189)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkdtemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µpath = πTemp001
					// line 190: try:
					πF.SetLineno(190)
					πF.PushCheckpoint(1)
					πF.PushCheckpoint(2)
					// line 191: os.rmdir(path + '/nonexistent')
					πF.SetLineno(191)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µpath, πg.NewStr("/nonexistent").ToObject()); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrmdir, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πF.PopCheckpoint()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					// line 195: raise AssertionError
					πF.SetLineno(195)
					πE = πF.Raise(πTemp001, nil, nil)
					continue
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 192: except OSError:
					πF.SetLineno(192)
				Label3:
					// line 193: pass
					πF.SetLineno(193)
					πF.RestoreExc(nil, nil)
					πF.PopCheckpoint()
					goto Label1
				Label1:
					πTemp004, πTemp005 = πF.RestoreExc(nil, nil)
					// line 197: os.rmdir(path)
					πF.SetLineno(197)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp003[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrmdir, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
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
			if πE = πF.Globals().SetItem(πF, ßTestRmDirNoExist.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 200: def TestRmDirFile():
			πF.SetLineno(200)
			πTemp003 = make([]πg.Param, 0)
			πTemp019 = πg.NewFunction(πg.NewCode("TestRmDirFile", "build/src/__python__/os_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfd *πg.Object = πg.UnboundLocal; _ = µfd
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var πTemp001 *πg.Object
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
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 201: fd, path = tempfile.mkstemp()
					πF.SetLineno(201)
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
					µfd = πTemp002
					µpath = πTemp003
					// line 202: os.close(fd)
					πF.SetLineno(202)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp004[0] = µfd
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßclose, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 203: try:
					πF.SetLineno(203)
					πF.PushCheckpoint(1)
					πF.PushCheckpoint(2)
					// line 204: os.rmdir(path)
					πF.SetLineno(204)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp004[0] = µpath
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
					πF.PopCheckpoint()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					// line 208: raise AssertionError
					πF.SetLineno(208)
					πE = πF.Raise(πTemp001, nil, nil)
					continue
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp005, πTemp006 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label3
					}
					πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
					continue
					// line 205: except OSError:
					πF.SetLineno(205)
				Label3:
					// line 206: pass
					πF.SetLineno(206)
					πF.RestoreExc(nil, nil)
					πF.PopCheckpoint()
					goto Label1
				Label1:
					πTemp005, πTemp006 = πF.RestoreExc(nil, nil)
					// line 210: os.remove(path)
					πF.SetLineno(210)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp004[0] = µpath
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
			if πE = πF.Globals().SetItem(πF, ßTestRmDirFile.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 213: def TestStatFile():
			πF.SetLineno(213)
			πTemp003 = make([]πg.Param, 0)
			πTemp020 = πg.NewFunction(πg.NewCode("TestStatFile", "build/src/__python__/os_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µt *πg.Object = πg.UnboundLocal; _ = µt
				var µfd *πg.Object = πg.UnboundLocal; _ = µfd
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var µst *πg.Object = πg.UnboundLocal; _ = µst
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 214: t = time.time()
					πF.SetLineno(214)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtime, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µt = πTemp001
					// line 215: fd, path = tempfile.mkstemp()
					πF.SetLineno(215)
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
					µfd = πTemp002
					µpath = πTemp003
					// line 216: os.close(fd)
					πF.SetLineno(216)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp004[0] = µfd
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßclose, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 217: st = os.stat(path)
					πF.SetLineno(217)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp004[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstat, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µst = πTemp001
					// line 218: os.remove(path)
					πF.SetLineno(218)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp004[0] = µpath
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
					// line 219: assert not stat.S_ISDIR(st.st_mode)
					πF.SetLineno(219)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µst, "st"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µst, ßst_mode, nil); πE != nil {
						continue
					}
					πTemp004[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßS_ISDIR, nil); πE != nil {
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
					// line 220: assert stat.S_IMODE(st.st_mode) == 0o600
					πF.SetLineno(220)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µst, "st"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µst, ßst_mode, nil); πE != nil {
						continue
					}
					πTemp004[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßS_IMODE, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp001, πE = πg.Eq(πF, πTemp002, πg.NewInt(384).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 222: assert st.st_mtime + 10 > t
					πF.SetLineno(222)
					if πE = πg.CheckLocal(πF, µst, "st"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µst, ßst_mtime, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp003, πg.NewInt(10).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GT(πF, πTemp002, µt); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 223: assert st.st_size == 0
					πF.SetLineno(223)
					if πE = πg.CheckLocal(πF, µst, "st"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µst, ßst_size, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, πTemp002, πg.NewInt(0).ToObject()); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestStatFile.ToObject(), πTemp020); πE != nil {
				continue
			}
			// line 226: def TestStatDir():
			πF.SetLineno(226)
			πTemp003 = make([]πg.Param, 0)
			πTemp021 = πg.NewFunction(πg.NewCode("TestStatDir", "build/src/__python__/os_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var µmode *πg.Object = πg.UnboundLocal; _ = µmode
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
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
					// line 227: path = tempfile.mkdtemp()
					πF.SetLineno(227)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkdtemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µpath = πTemp001
					// line 228: mode = os.stat(path).st_mode
					πF.SetLineno(228)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp003[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstat, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßst_mode, nil); πE != nil {
						continue
					}
					µmode = πTemp002
					// line 229: os.rmdir(path)
					πF.SetLineno(229)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp003[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrmdir, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 230: assert stat.S_ISDIR(mode)
					πF.SetLineno(230)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					πTemp003[0] = µmode
					if πTemp001, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßS_ISDIR, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 231: assert stat.S_IMODE(mode) == 0o700
					πF.SetLineno(231)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					πTemp003[0] = µmode
					if πTemp002, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßS_IMODE, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp001, πE = πg.Eq(πF, πTemp002, πg.NewInt(448).ToObject()); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestStatDir.ToObject(), πTemp021); πE != nil {
				continue
			}
			// line 234: def TestStatNoExist():
			πF.SetLineno(234)
			πTemp003 = make([]πg.Param, 0)
			πTemp022 = πg.NewFunction(πg.NewCode("TestStatNoExist", "build/src/__python__/os_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
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
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 235: path = tempfile.mkdtemp()
					πF.SetLineno(235)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkdtemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µpath = πTemp001
					// line 236: try:
					πF.SetLineno(236)
					πF.PushCheckpoint(1)
					πF.PushCheckpoint(2)
					// line 237: os.stat(path + '/nonexistent')
					πF.SetLineno(237)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µpath, πg.NewStr("/nonexistent").ToObject()); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstat, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πF.PopCheckpoint()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					// line 241: raise AssertionError
					πF.SetLineno(241)
					πE = πF.Raise(πTemp001, nil, nil)
					continue
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 238: except OSError:
					πF.SetLineno(238)
				Label3:
					// line 239: pass
					πF.SetLineno(239)
					πF.RestoreExc(nil, nil)
					πF.PopCheckpoint()
					goto Label1
				Label1:
					πTemp004, πTemp005 = πF.RestoreExc(nil, nil)
					// line 243: os.rmdir(path)
					πF.SetLineno(243)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp003[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrmdir, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
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
			if πE = πF.Globals().SetItem(πF, ßTestStatNoExist.ToObject(), πTemp022); πE != nil {
				continue
			}
			// line 246: def TestWaitPid():
			πF.SetLineno(246)
			πTemp003 = make([]πg.Param, 0)
			πTemp023 = πg.NewFunction(πg.NewCode("TestWaitPid", "build/src/__python__/os_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpid *πg.Object = πg.UnboundLocal; _ = µpid
				var µstatus *πg.Object = πg.UnboundLocal; _ = µstatus
				var µe *πg.Object = πg.UnboundLocal; _ = µe
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
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
					// line 247: try:
					πF.SetLineno(247)
					πF.PushCheckpoint(2)
					// line 248: pid, status = os.waitpid(-1, os.WNOHANG)
					πF.SetLineno(248)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßWNOHANG, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßwaitpid, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
						continue
					}
					µpid = πTemp003
					µstatus = πTemp004
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp005, πTemp006 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label3
					}
					πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
					continue
					// line 249: except OSError as e:
					πF.SetLineno(249)
				Label3:
					µe = πTemp005.ToObject()
					// line 250: assert 'no child processes' in str(e).lower()
					πF.SetLineno(250)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
						continue
					}
					πTemp001[0] = µe
					if πTemp003, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßlower, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Contains(πF, πTemp004, πg.NewStr("no child processes").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp007).ToObject()
					if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßTestWaitPid.ToObject(), πTemp023); πE != nil {
				continue
			}
			if πTemp025, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp024, πE = πg.Eq(πF, πTemp025, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp026, πE = πg.IsTrue(πF, πTemp024); πE != nil {
				continue
			}
			if πTemp026 {
				goto Label1
			}
			goto Label2
			// line 253: if __name__ == '__main__':
			πF.SetLineno(253)
		Label1:
			// line 254: weetest.RunTests()
			πF.SetLineno(254)
			if πTemp024, πE = πg.ResolveGlobal(πF, ßweetest); πE != nil {
				continue
			}
			if πTemp025, πE = πg.GetAttr(πF, πTemp024, ßRunTests, nil); πE != nil {
				continue
			}
			if πTemp024, πE = πTemp025.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("os_test", Code)
}
