package stat
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/stat.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßModeDir := πg.InternStr("ModeDir")
		ßModePerm := πg.InternStr("ModePerm")
		ßS_IMODE := πg.InternStr("S_IMODE")
		ßS_ISDIR := πg.InternStr("S_ISDIR")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 []πg.Param
		_ = πTemp004
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 15: """Interpreting stat() results."""
			πF.SetLineno(15)
			// line 18: from '__go__/os' import ModeDir, ModePerm
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßModeDir, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßModeDir.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßModePerm, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßModePerm.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 21: def S_ISDIR(mode):  # pylint: disable=invalid-name
			πF.SetLineno(21)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "mode", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("S_ISDIR", "build/src/__python__/stat.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmode *πg.Object = πArgs[0]; _ = µmode
				var πTemp001 *πg.Object
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
					// line 22: return mode & ModeDir != 0
					πF.SetLineno(22)
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßModeDir); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, µmode, πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, πTemp002, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßS_ISDIR.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 25: def S_IMODE(mode):  # pylint: disable=invalid-name
			πF.SetLineno(25)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "mode", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("S_IMODE", "build/src/__python__/stat.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmode *πg.Object = πArgs[0]; _ = µmode
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
					// line 26: return mode & ModePerm
					πF.SetLineno(26)
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßModePerm); πE != nil {
						continue
					}
					if πTemp001, πE = πg.And(πF, µmode, πTemp002); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßS_IMODE.ToObject(), πTemp003); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("stat", Code)
}
