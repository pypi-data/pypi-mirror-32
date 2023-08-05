package exceptions
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/exceptions.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßExceptionTypes := πg.InternStr("ExceptionTypes")
		ß__name__ := πg.InternStr("__name__")
		ßg := πg.InternStr("g")
		ßglobals := πg.InternStr("globals")
		ßt := πg.InternStr("t")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 bool
		_ = πTemp004
		var πTemp005 bool
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
			case 1: goto Label1
			case 2: goto Label2
			default: panic("unexpected function state")
			}
			// line 15: """Built-in exception classes."""
			πF.SetLineno(15)
			// line 17: from '__go__/grumpy' import ExceptionTypes
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/grumpy"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßExceptionTypes, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßExceptionTypes.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 20: g = globals()
			πF.SetLineno(20)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßglobals); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßg.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßExceptionTypes); πE != nil {
				continue
			}
			if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
				continue
			}
			πF.PushCheckpoint(2)
			πTemp004 = false
		Label1:
			if πE != nil || πR != nil {
				continue
			}
			if πTemp004 {
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
				πTemp005 = !isStop
			} else {
				πTemp005 = true
				if πE = πF.Globals().SetItem(πF, ßt.ToObject(), πTemp003); πE != nil {
					continue
				}
			}
			if πE != nil || !πTemp005 {
				continue
			}
			πF.PushCheckpoint(1)            
			// line 22: g[t.__name__] = t
			πF.SetLineno(22)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßt); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp003); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßg); πE != nil {
				continue
			}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßt); πE != nil {
				continue
			}
			if πTemp010, πE = πg.GetAttr(πF, πTemp009, ß__name__, nil); πE != nil {
				continue
			}
			πTemp008 = πTemp010
			if πE = πg.SetItem(πF, πTemp007, πTemp008, πTemp006); πE != nil {
				continue
			}
			continue
		Label2:
			if πE != nil || πR != nil {
				continue
			}
		Label3:
		}
		return nil, πE
	})
	πg.RegisterModule("exceptions", Code)
}
