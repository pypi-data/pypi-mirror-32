package tempfile_test
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/tempfile_test.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßAssertionError := πg.InternStr("AssertionError")
		ßOSError := πg.InternStr("OSError")
		ßRunTests := πg.InternStr("RunTests")
		ßS_IMODE := πg.InternStr("S_IMODE")
		ßS_ISDIR := πg.InternStr("S_ISDIR")
		ßTestMkdTemp := πg.InternStr("TestMkdTemp")
		ßTestMkdTempDir := πg.InternStr("TestMkdTempDir")
		ßTestMkdTempOSError := πg.InternStr("TestMkdTempOSError")
		ßTestMkdTempPrefixSuffix := πg.InternStr("TestMkdTempPrefixSuffix")
		ßTestMksTemp := πg.InternStr("TestMksTemp")
		ßTestMksTempDir := πg.InternStr("TestMksTempDir")
		ßTestMksTempOSError := πg.InternStr("TestMksTempOSError")
		ßTestMksTempPerms := πg.InternStr("TestMksTempPerms")
		ßTestMksTempPrefixSuffix := πg.InternStr("TestMksTempPrefixSuffix")
		ß__main__ := πg.InternStr("__main__")
		ß__name__ := πg.InternStr("__name__")
		ßbar := πg.InternStr("bar")
		ßchmod := πg.InternStr("chmod")
		ßclose := πg.InternStr("close")
		ßfdopen := πg.InternStr("fdopen")
		ßfoo := πg.InternStr("foo")
		ßfoobar := πg.InternStr("foobar")
		ßmkdtemp := πg.InternStr("mkdtemp")
		ßmkstemp := πg.InternStr("mkstemp")
		ßopen := πg.InternStr("open")
		ßos := πg.InternStr("os")
		ßread := πg.InternStr("read")
		ßremove := πg.InternStr("remove")
		ßrmdir := πg.InternStr("rmdir")
		ßst_mode := πg.InternStr("st_mode")
		ßstartswith := πg.InternStr("startswith")
		ßstat := πg.InternStr("stat")
		ßtempfile := πg.InternStr("tempfile")
		ßw := πg.InternStr("w")
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
		var πTemp014 bool
		_ = πTemp014
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
			// line 17: import tempfile
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "tempfile"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtempfile.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 19: import weetest
			πF.SetLineno(19)
			if πTemp002, πE = πg.ImportModule(πF, "weetest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßweetest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 22: def TestMkdTemp():
			πF.SetLineno(22)
			πTemp003 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("TestMkdTemp", "build/src/__python__/tempfile_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 23: path = tempfile.mkdtemp()
					πF.SetLineno(23)
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
					// line 24: mode = os.stat(path).st_mode
					πF.SetLineno(24)
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
					// line 25: os.rmdir(path)
					πF.SetLineno(25)
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
					// line 26: assert stat.S_ISDIR(mode), mode
					πF.SetLineno(26)
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
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
					if πE = πg.Assert(πF, πTemp001, µmode); πE != nil {
						continue
					}
					// line 27: assert stat.S_IMODE(mode) == 0o700, mode
					πF.SetLineno(27)
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
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
					if πE = πg.Assert(πF, πTemp001, µmode); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestMkdTemp.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 30: def TestMkdTempDir():
			πF.SetLineno(30)
			πTemp003 = make([]πg.Param, 0)
			πTemp004 = πg.NewFunction(πg.NewCode("TestMkdTempDir", "build/src/__python__/tempfile_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtempdir *πg.Object = πg.UnboundLocal; _ = µtempdir
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 πg.KWArgs
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
					// line 31: tempdir = tempfile.mkdtemp()
					πF.SetLineno(31)
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
					// line 32: path = tempfile.mkdtemp(dir=tempdir)
					πF.SetLineno(32)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp003 = πg.KWArgs{
						{"dir", µtempdir},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkdtemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, πTemp003); πE != nil {
						continue
					}
					µpath = πTemp001
					// line 33: os.rmdir(path)
					πF.SetLineno(33)
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
					// line 34: os.rmdir(tempdir)
					πF.SetLineno(34)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp004[0] = µtempdir
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
					// line 35: assert path.startswith(tempdir)
					πF.SetLineno(35)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp004[0] = µtempdir
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpath, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
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
			if πE = πF.Globals().SetItem(πF, ßTestMkdTempDir.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 38: def TestMkdTempOSError():
			πF.SetLineno(38)
			πTemp003 = make([]πg.Param, 0)
			πTemp005 = πg.NewFunction(πg.NewCode("TestMkdTempOSError", "build/src/__python__/tempfile_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtempdir *πg.Object = πg.UnboundLocal; _ = µtempdir
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 πg.KWArgs
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
					// line 39: tempdir = tempfile.mkdtemp()
					πF.SetLineno(39)
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
					// line 40: os.chmod(tempdir, 0o500)
					πF.SetLineno(40)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp003[0] = µtempdir
					πTemp003[1] = πg.NewInt(320).ToObject()
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
					// line 41: try:
					πF.SetLineno(41)
					πF.PushCheckpoint(2)
					// line 42: tempfile.mkdtemp(dir=tempdir)
					πF.SetLineno(42)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp004 = πg.KWArgs{
						{"dir", µtempdir},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkdtemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, πTemp004); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					// line 46: raise AssertionError
					πF.SetLineno(46)
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
					// line 43: except OSError:
					πF.SetLineno(43)
				Label3:
					// line 44: pass
					πF.SetLineno(44)
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 47: os.rmdir(tempdir)
					πF.SetLineno(47)
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
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßTestMkdTempOSError.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 50: def TestMkdTempPrefixSuffix():
			πF.SetLineno(50)
			πTemp003 = make([]πg.Param, 0)
			πTemp006 = πg.NewFunction(πg.NewCode("TestMkdTempPrefixSuffix", "build/src/__python__/tempfile_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var πTemp001 πg.KWArgs
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
					// line 51: path = tempfile.mkdtemp(prefix='foo', suffix='bar')
					πF.SetLineno(51)
					πTemp001 = πg.KWArgs{
						{"prefix", ßfoo.ToObject()},
						{"suffix", ßbar.ToObject()},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmkdtemp, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, nil, πTemp001); πE != nil {
						continue
					}
					µpath = πTemp002
					// line 52: os.rmdir(path)
					πF.SetLineno(52)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp004[0] = µpath
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrmdir, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 53: assert 'foo' in path
					πF.SetLineno(53)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, µpath, ßfoo.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp005).ToObject()
					if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
						continue
					}
					// line 54: assert 'bar' in path
					πF.SetLineno(54)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, µpath, ßbar.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp005).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßTestMkdTempPrefixSuffix.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 58: def TestMksTemp():
			πF.SetLineno(58)
			πTemp003 = make([]πg.Param, 0)
			πTemp007 = πg.NewFunction(πg.NewCode("TestMksTemp", "build/src/__python__/tempfile_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 59: fd, path = tempfile.mkstemp()
					πF.SetLineno(59)
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
					// line 60: f = os.fdopen(fd, 'w')
					πF.SetLineno(60)
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
					// line 61: f.write('foobar')
					πF.SetLineno(61)
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
					// line 62: f.close()
					πF.SetLineno(62)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 63: f = open(path)
					πF.SetLineno(63)
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
					// line 64: contents = f.read()
					πF.SetLineno(64)
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
					// line 65: f.close()
					πF.SetLineno(65)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 66: os.remove(path)
					πF.SetLineno(66)
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
					// line 67: assert contents == 'foobar', contents
					πF.SetLineno(67)
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
			if πE = πF.Globals().SetItem(πF, ßTestMksTemp.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 70: def TestMksTempDir():
			πF.SetLineno(70)
			πTemp003 = make([]πg.Param, 0)
			πTemp008 = πg.NewFunction(πg.NewCode("TestMksTempDir", "build/src/__python__/tempfile_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtempdir *πg.Object = πg.UnboundLocal; _ = µtempdir
				var µfd *πg.Object = πg.UnboundLocal; _ = µfd
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 πg.KWArgs
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
					// line 71: tempdir = tempfile.mkdtemp()
					πF.SetLineno(71)
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
					// line 72: fd, path = tempfile.mkstemp(dir=tempdir)
					πF.SetLineno(72)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp003 = πg.KWArgs{
						{"dir", µtempdir},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkstemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, πTemp003); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
						continue
					}
					µfd = πTemp002
					µpath = πTemp004
					// line 73: os.close(fd)
					πF.SetLineno(73)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp005[0] = µfd
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßclose, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 74: os.remove(path)
					πF.SetLineno(74)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp005[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßremove, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 75: os.rmdir(tempdir)
					πF.SetLineno(75)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp005[0] = µtempdir
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrmdir, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 76: assert path.startswith(tempdir)
					πF.SetLineno(76)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp005[0] = µtempdir
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpath, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
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
			if πE = πF.Globals().SetItem(πF, ßTestMksTempDir.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 79: def TestMksTempOSError():
			πF.SetLineno(79)
			πTemp003 = make([]πg.Param, 0)
			πTemp009 = πg.NewFunction(πg.NewCode("TestMksTempOSError", "build/src/__python__/tempfile_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtempdir *πg.Object = πg.UnboundLocal; _ = µtempdir
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 πg.KWArgs
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
					// line 80: tempdir = tempfile.mkdtemp()
					πF.SetLineno(80)
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
					// line 81: os.chmod(tempdir, 0o500)
					πF.SetLineno(81)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp003[0] = µtempdir
					πTemp003[1] = πg.NewInt(320).ToObject()
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
					// line 82: try:
					πF.SetLineno(82)
					πF.PushCheckpoint(2)
					// line 83: tempfile.mkstemp(dir=tempdir)
					πF.SetLineno(83)
					if πE = πg.CheckLocal(πF, µtempdir, "tempdir"); πE != nil {
						continue
					}
					πTemp004 = πg.KWArgs{
						{"dir", µtempdir},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkstemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, πTemp004); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					// line 87: raise AssertionError
					πF.SetLineno(87)
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
					// line 84: except OSError:
					πF.SetLineno(84)
				Label3:
					// line 85: pass
					πF.SetLineno(85)
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 88: os.rmdir(tempdir)
					πF.SetLineno(88)
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
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßTestMksTempOSError.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 91: def TestMksTempPerms():
			πF.SetLineno(91)
			πTemp003 = make([]πg.Param, 0)
			πTemp010 = πg.NewFunction(πg.NewCode("TestMksTempPerms", "build/src/__python__/tempfile_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 92: fd, path = tempfile.mkstemp()
					πF.SetLineno(92)
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
					// line 93: os.close(fd)
					πF.SetLineno(93)
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
					// line 94: mode = os.stat(path).st_mode
					πF.SetLineno(94)
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
					µmode = πTemp002
					// line 95: os.remove(path)
					πF.SetLineno(95)
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
					// line 96: assert stat.S_IMODE(mode) == 0o600, mode
					πF.SetLineno(96)
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					πTemp004[0] = µmode
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
					if πE = πg.Assert(πF, πTemp001, µmode); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestMksTempPerms.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 99: def TestMksTempPrefixSuffix():
			πF.SetLineno(99)
			πTemp003 = make([]πg.Param, 0)
			πTemp011 = πg.NewFunction(πg.NewCode("TestMksTempPrefixSuffix", "build/src/__python__/tempfile_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfd *πg.Object = πg.UnboundLocal; _ = µfd
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var πTemp001 πg.KWArgs
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 100: fd, path = tempfile.mkstemp(prefix='foo', suffix='bar')
					πF.SetLineno(100)
					πTemp001 = πg.KWArgs{
						{"prefix", ßfoo.ToObject()},
						{"suffix", ßbar.ToObject()},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmkstemp, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, nil, πTemp001); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
						continue
					}
					µfd = πTemp003
					µpath = πTemp004
					// line 101: os.close(fd)
					πF.SetLineno(101)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp005[0] = µfd
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßclose, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 102: os.remove(path)
					πF.SetLineno(102)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp005[0] = µpath
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßremove, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 103: assert 'foo' in path
					πF.SetLineno(103)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Contains(πF, µpath, ßfoo.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp006).ToObject()
					if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
						continue
					}
					// line 104: assert 'bar' in path
					πF.SetLineno(104)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Contains(πF, µpath, ßbar.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp006).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßTestMksTempPrefixSuffix.ToObject(), πTemp011); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp012, πE = πg.Eq(πF, πTemp013, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp014, πE = πg.IsTrue(πF, πTemp012); πE != nil {
				continue
			}
			if πTemp014 {
				goto Label1
			}
			goto Label2
			// line 108: if __name__ == '__main__':
			πF.SetLineno(108)
		Label1:
			// line 109: weetest.RunTests()
			πF.SetLineno(109)
			if πTemp012, πE = πg.ResolveGlobal(πF, ßweetest); πE != nil {
				continue
			}
			if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßRunTests, nil); πE != nil {
				continue
			}
			if πTemp012, πE = πTemp013.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("tempfile_test", Code)
}
