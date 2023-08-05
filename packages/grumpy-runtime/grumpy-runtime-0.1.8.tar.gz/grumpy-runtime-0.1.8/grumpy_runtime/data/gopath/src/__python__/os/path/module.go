package path
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/os/path.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßAbs := πg.InternStr("Abs")
		ßBase := πg.InternStr("Base")
		ßClean := πg.InternStr("Clean")
		ßDir := πg.InternStr("Dir")
		ßError := πg.InternStr("Error")
		ßFalse := πg.InternStr("False")
		ßIsAbs := πg.InternStr("IsAbs")
		ßIsDir := πg.InternStr("IsDir")
		ßIsRegular := πg.InternStr("IsRegular")
		ßJoin := πg.InternStr("Join")
		ßMode := πg.InternStr("Mode")
		ßNone := πg.InternStr("None")
		ßOSError := πg.InternStr("OSError")
		ßSplit := πg.InternStr("Split")
		ßStat := πg.InternStr("Stat")
		ßTypeError := πg.InternStr("TypeError")
		ßabspath := πg.InternStr("abspath")
		ßappend := πg.InternStr("append")
		ßbasename := πg.InternStr("basename")
		ßdirname := πg.InternStr("dirname")
		ßendswith := πg.InternStr("endswith")
		ßexists := πg.InternStr("exists")
		ßisabs := πg.InternStr("isabs")
		ßisdir := πg.InternStr("isdir")
		ßisfile := πg.InternStr("isfile")
		ßisinstance := πg.InternStr("isinstance")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßnormpath := πg.InternStr("normpath")
		ßsplit := πg.InternStr("split")
		ßunicode := πg.InternStr("unicode")
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 15: """"Utilities for manipulating and inspecting OS paths."""
			πF.SetLineno(15)
			// line 17: from '__go__/os' import Stat
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßStat, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStat.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 18: from '__go__/path/filepath' import Abs, Base, Clean, Dir as dirname, IsAbs as isabs, Join, Split  # pylint: disable=g-multiple-import,unused-import
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/path/filepath"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßAbs, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßAbs.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßBase, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBase.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßClean, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßClean.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßDir, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdirname.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßIsAbs, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßisabs.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßJoin, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßJoin.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßSplit, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSplit.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 21: def abspath(path):
			πF.SetLineno(21)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "path", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("abspath", "build/src/__python__/os/path.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πArgs[0]; _ = µpath
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
					// line 22: result, err = Abs(path)
					πF.SetLineno(22)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp001[0] = µpath
					if πTemp002, πE = πg.ResolveGlobal(πF, ßAbs); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µresult = πTemp002
					µerr = πTemp004
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µerr); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 23: if err:
					πF.SetLineno(23)
				Label1:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 24: raise OSError(err.Error())
					πF.SetLineno(24)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label2
				Label2:
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp001[0] = µpath
					if πTemp002, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label3
					}
					goto Label4
					// line 25: if isinstance(path, unicode):
					πF.SetLineno(25)
				Label3:
					// line 28: return unicode(result, 'utf-8')
					πF.SetLineno(28)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp001[0] = µresult
					πTemp001[1] = πg.NewStr("utf-8").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
					goto Label4
				Label4:
					// line 29: return result
					πF.SetLineno(29)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πR = µresult
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßabspath.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 32: def basename(path):
			πF.SetLineno(32)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "path", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("basename", "build/src/__python__/os/path.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πArgs[0]; _ = µpath
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 33: return '' if path.endswith('/') else Base(path)
					πF.SetLineno(33)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("/").ToObject()
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µpath, ßendswith, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label1
					}
					πTemp001 = ß.ToObject()
					goto Label2
				Label1:
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp002[0] = µpath
					if πTemp003, πE = πg.ResolveGlobal(πF, ßBase); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001 = πTemp004
				Label2:
					πR = πTemp001
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßbasename.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 36: def exists(path):
			πF.SetLineno(36)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "path", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("exists", "build/src/__python__/os/path.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πArgs[0]; _ = µpath
				var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
				var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
					// line 37: _, err = Stat(path)
					πF.SetLineno(37)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp001[0] = µpath
					if πTemp002, πE = πg.ResolveGlobal(πF, ßStat); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µ_ = πTemp002
					µerr = πTemp004
					// line 38: return err is None
					πF.SetLineno(38)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µerr == πTemp003).ToObject()
					πR = πTemp002
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßexists.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 41: def isdir(path):
			πF.SetLineno(41)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "path", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("isdir", "build/src/__python__/os/path.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πArgs[0]; _ = µpath
				var µinfo *πg.Object = πg.UnboundLocal; _ = µinfo
				var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
					// line 42: info, err = Stat(path)
					πF.SetLineno(42)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp001[0] = µpath
					if πTemp002, πE = πg.ResolveGlobal(πF, ßStat); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µinfo = πTemp002
					µerr = πTemp004
					if πE = πg.CheckLocal(πF, µinfo, "info"); πE != nil {
						continue
					}
					πTemp002 = µinfo
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µerr == πTemp004).ToObject()
					πTemp002 = πTemp003
				Label1:
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label2
					}
					goto Label3
					// line 43: if info and err is None:
					πF.SetLineno(43)
				Label2:
					// line 44: return info.Mode().IsDir()
					πF.SetLineno(44)
					if πE = πg.CheckLocal(πF, µinfo, "info"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µinfo, ßMode, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßIsDir, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πR = πTemp003
					continue
					goto Label3
				Label3:
					// line 45: return False
					πF.SetLineno(45)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp002
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßisdir.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 48: def isfile(path):
			πF.SetLineno(48)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "path", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("isfile", "build/src/__python__/os/path.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πArgs[0]; _ = µpath
				var µinfo *πg.Object = πg.UnboundLocal; _ = µinfo
				var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
					// line 49: info, err = Stat(path)
					πF.SetLineno(49)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp001[0] = µpath
					if πTemp002, πE = πg.ResolveGlobal(πF, ßStat); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µinfo = πTemp002
					µerr = πTemp004
					if πE = πg.CheckLocal(πF, µinfo, "info"); πE != nil {
						continue
					}
					πTemp002 = µinfo
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µerr == πTemp004).ToObject()
					πTemp002 = πTemp003
				Label1:
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label2
					}
					goto Label3
					// line 50: if info and err is None:
					πF.SetLineno(50)
				Label2:
					// line 51: return info.Mode().IsRegular()
					πF.SetLineno(51)
					if πE = πg.CheckLocal(πF, µinfo, "info"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µinfo, ßMode, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßIsRegular, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πR = πTemp003
					continue
					goto Label3
				Label3:
					// line 52: return False
					πF.SetLineno(52)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp002
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßisfile.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 60: def join(*paths):
			πF.SetLineno(60)
			πTemp004 = make([]πg.Param, 0)
			πTemp008 = πg.NewFunction(πg.NewCode("join", "build/src/__python__/os/path.py", πTemp004, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpaths *πg.Object = πArgs[0]; _ = µpaths
				var µparts *πg.Object = πg.UnboundLocal; _ = µparts
				var µp *πg.Object = πg.UnboundLocal; _ = µp
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 3: goto Label3
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					if πE = πg.CheckLocal(πF, µpaths, "paths"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µpaths); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 61: if not paths:
					πF.SetLineno(61)
				Label1:
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("join() takes at least 1 argument (0 given)").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 62: raise TypeError('join() takes at least 1 argument (0 given)')
					πF.SetLineno(62)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					goto Label2
				Label2:
					// line 63: parts = []
					πF.SetLineno(63)
					πTemp003 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					µparts = πTemp001
					if πE = πg.CheckLocal(πF, µpaths, "paths"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µpaths); πE != nil {
						continue
					}
					πF.PushCheckpoint(4)
					πTemp002 = false
				Label3:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
						πF.PopCheckpoint()
						goto Label5
					}
					if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp005 = !isStop
					} else {
						πTemp005 = true
						µp = πTemp004
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(3)            
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					πTemp003[0] = µp
					if πTemp004, πE = πg.ResolveGlobal(πF, ßisabs); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label6
					}
					goto Label7
					// line 65: if isabs(p):
					πF.SetLineno(65)
				Label6:
					// line 66: parts = [p]
					πF.SetLineno(66)
					πTemp003 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					πTemp003[0] = µp
					πTemp004 = πg.NewList(πTemp003...).ToObject()
					µparts = πTemp004
					goto Label8
				Label7:
					// line 68: parts.append(p)
					πF.SetLineno(68)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					πTemp003[0] = µp
					if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µparts, ßappend, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label8
				Label8:
					continue
				Label4:
					if πE != nil || πR != nil {
						continue
					}
				Label5:
					// line 69: result = Join(*parts)
					πF.SetLineno(69)
					if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßJoin); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Invoke(πF, πTemp001, nil, µparts, nil, nil); πE != nil {
						continue
					}
					µresult = πTemp004
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp001 = µresult
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label9
					}
					if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp006 = πTemp007
					if πE = πg.CheckLocal(πF, µpaths, "paths"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µpaths, πTemp006); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(!πTemp005).ToObject()
					πTemp001 = πTemp004
				Label9:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label10
					}
					goto Label11
					// line 70: if result and not paths[-1]:
					πF.SetLineno(70)
				Label10:
					// line 71: result += '/'
					πF.SetLineno(71)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µresult, πg.NewStr("/").ToObject()); πE != nil {
						continue
					}
					µresult = πTemp001
					goto Label11
				Label11:
					// line 72: return result
					πF.SetLineno(72)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πR = µresult
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßjoin.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 75: def normpath(path):
			πF.SetLineno(75)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "path", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("normpath", "build/src/__python__/os/path.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πArgs[0]; _ = µpath
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
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
					// line 76: result = Clean(path)
					πF.SetLineno(76)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp001[0] = µpath
					if πTemp002, πE = πg.ResolveGlobal(πF, ßClean); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µresult = πTemp003
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp001[0] = µpath
					if πTemp002, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
					// line 77: if isinstance(path, unicode):
					πF.SetLineno(77)
				Label1:
					// line 78: return unicode(result, 'utf-8')
					πF.SetLineno(78)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp001[0] = µresult
					πTemp001[1] = πg.NewStr("utf-8").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
					goto Label2
				Label2:
					// line 79: return result
					πF.SetLineno(79)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πR = µresult
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßnormpath.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 82: def split(path):
			πF.SetLineno(82)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "path", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("split", "build/src/__python__/os/path.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πArgs[0]; _ = µpath
				var µhead *πg.Object = πg.UnboundLocal; _ = µhead
				var µtail *πg.Object = πg.UnboundLocal; _ = µtail
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 83: head, tail = Split(path)
					πF.SetLineno(83)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp001[0] = µpath
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSplit); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µhead = πTemp002
					µtail = πTemp004
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µhead, "head"); πE != nil {
						continue
					}
					πTemp001[0] = µhead
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.GT(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label1
					}
					if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp006
					if πE = πg.CheckLocal(πF, µhead, "head"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µhead, πTemp004); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp006, πg.NewStr("/").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label1:
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label2
					}
					goto Label3
					// line 84: if len(head) > 1 and head[-1] == '/':
					πF.SetLineno(84)
				Label2:
					// line 85: head = head[:-1]
					πF.SetLineno(85)
					if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp003, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µhead, "head"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µhead, πTemp002); πE != nil {
						continue
					}
					µhead = πTemp003
					goto Label3
				Label3:
					// line 86: return (head, tail)
					πF.SetLineno(86)
					if πE = πg.CheckLocal(πF, µhead, "head"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µhead, µtail).ToObject()
					πR = πTemp002
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßsplit.ToObject(), πTemp010); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("os.path", Code)
}
