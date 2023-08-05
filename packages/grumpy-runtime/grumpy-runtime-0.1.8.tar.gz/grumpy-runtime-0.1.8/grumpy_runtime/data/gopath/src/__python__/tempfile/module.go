package tempfile
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/tempfile.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßClose := πg.InternStr("Close")
		ßDup := πg.InternStr("Dup")
		ßError := πg.InternStr("Error")
		ßFalse := πg.InternStr("False")
		ßFd := πg.InternStr("Fd")
		ßName := πg.InternStr("Name")
		ßNone := πg.InternStr("None")
		ßNotImplementedError := πg.InternStr("NotImplementedError")
		ßOSError := πg.InternStr("OSError")
		ßTempDir := πg.InternStr("TempDir")
		ßTempFile := πg.InternStr("TempFile")
		ßmkdtemp := πg.InternStr("mkdtemp")
		ßmkstemp := πg.InternStr("mkstemp")
		ßtmp := πg.InternStr("tmp")
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 15: """Generate temporary files and directories."""
			πF.SetLineno(15)
			// line 18: from '__go__/io/ioutil' import TempDir, TempFile
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/io/ioutil"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßTempDir, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTempDir.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßTempFile, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTempFile.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 19: from '__go__/syscall' import Dup
			πF.SetLineno(19)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/syscall"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßDup, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDup.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 23: def mkdtemp(suffix='', prefix='tmp', dir=None):
			πF.SetLineno(23)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "suffix", Def: ß.ToObject()}
			πTemp004[1] = πg.Param{Name: "prefix", Def: ßtmp.ToObject()}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "dir", Def: πTemp003}
			πTemp001 = πg.NewFunction(πg.NewCode("mkdtemp", "build/src/__python__/tempfile.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsuffix *πg.Object = πArgs[0]; _ = µsuffix
				var µprefix *πg.Object = πArgs[1]; _ = µprefix
				var µdir *πg.Object = πArgs[2]; _ = µdir
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var µerr *πg.Object = πg.UnboundLocal; _ = µerr
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 []*πg.Object
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
					if πE = πg.CheckLocal(πF, µdir, "dir"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µdir == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 24: if dir is None:
					πF.SetLineno(24)
				Label1:
					// line 25: dir = ''
					πF.SetLineno(25)
					µdir = ß.ToObject()
					goto Label2
				Label2:
					// line 27: path, err = TempDir(dir, prefix + '-' + suffix)
					πF.SetLineno(27)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µdir, "dir"); πE != nil {
						continue
					}
					πTemp004[0] = µdir
					if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µprefix, πg.NewStr("-").ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsuffix, "suffix"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp002, µsuffix); πE != nil {
						continue
					}
					πTemp004[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTempDir); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
						continue
					}
					µpath = πTemp001
					µerr = πTemp005
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µerr); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label3
					}
					goto Label4
					// line 28: if err:
					πF.SetLineno(28)
				Label3:
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp004[0] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 29: raise OSError(err.Error())
					πF.SetLineno(29)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label4
				Label4:
					// line 30: return path
					πF.SetLineno(30)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πR = µpath
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßmkdtemp.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 33: def mkstemp(suffix='', prefix='tmp', dir=None, text=False):
			πF.SetLineno(33)
			πTemp004 = make([]πg.Param, 4)
			πTemp004[0] = πg.Param{Name: "suffix", Def: ß.ToObject()}
			πTemp004[1] = πg.Param{Name: "prefix", Def: ßtmp.ToObject()}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "dir", Def: πTemp005}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[3] = πg.Param{Name: "text", Def: πTemp005}
			πTemp003 = πg.NewFunction(πg.NewCode("mkstemp", "build/src/__python__/tempfile.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsuffix *πg.Object = πArgs[0]; _ = µsuffix
				var µprefix *πg.Object = πArgs[1]; _ = µprefix
				var µdir *πg.Object = πArgs[2]; _ = µdir
				var µtext *πg.Object = πArgs[3]; _ = µtext
				var µf *πg.Object = πg.UnboundLocal; _ = µf
				var µerr *πg.Object = πg.UnboundLocal; _ = µerr
				var µfd *πg.Object = πg.UnboundLocal; _ = µfd
				var πTemp001 bool
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
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
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µtext); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label1
					}
					goto Label2
					// line 34: if text:
					πF.SetLineno(34)
				Label1:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
						continue
					}
					// line 35: raise NotImplementedError
					πF.SetLineno(35)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µdir, "dir"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µdir == πTemp003).ToObject()
					if πTemp001, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label3
					}
					goto Label4
					// line 36: if dir is None:
					πF.SetLineno(36)
				Label3:
					// line 37: dir = ''
					πF.SetLineno(37)
					µdir = ß.ToObject()
					goto Label4
				Label4:
					// line 39: f, err = TempFile(dir, prefix + '-' + suffix)
					πF.SetLineno(39)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µdir, "dir"); πE != nil {
						continue
					}
					πTemp004[0] = µdir
					if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µprefix, πg.NewStr("-").ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsuffix, "suffix"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp003, µsuffix); πE != nil {
						continue
					}
					πTemp004[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTempFile); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
						continue
					}
					µf = πTemp002
					µerr = πTemp005
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µerr); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label5
					}
					goto Label6
					// line 40: if err:
					πF.SetLineno(40)
				Label5:
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp004[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 41: raise OSError(err.Error())
					πF.SetLineno(41)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label6
				Label6:
					// line 42: try:
					πF.SetLineno(42)
					πF.PushCheckpoint(7)
					// line 43: fd, err = Dup(f.Fd())
					πF.SetLineno(43)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µf, ßFd, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp004[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßDup); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
						continue
					}
					µfd = πTemp002
					µerr = πTemp005
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µerr); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label8
					}
					goto Label9
					// line 44: if err:
					πF.SetLineno(44)
				Label8:
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp004[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 45: raise OSError(err.Error())
					πF.SetLineno(45)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label9
				Label9:
					// line 46: return fd, f.Name()
					πF.SetLineno(46)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µf, ßName, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µfd, πTemp005).ToObject()
					πR = πTemp002
					continue
					πF.PopCheckpoint()
					goto Label7
				Label7:
					πTemp006, πTemp007 = πF.RestoreExc(nil, nil)
					// line 48: f.Close()
					πF.SetLineno(48)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µf, ßClose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßmkstemp.ToObject(), πTemp003); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("tempfile", Code)
}
