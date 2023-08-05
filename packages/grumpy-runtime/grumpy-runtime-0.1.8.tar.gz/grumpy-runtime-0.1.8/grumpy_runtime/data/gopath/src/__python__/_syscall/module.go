package _syscall
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/_syscall.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßEINTR := πg.InternStr("EINTR")
		ßError := πg.InternStr("Error")
		ßOSError := πg.InternStr("OSError")
		ßTrue := πg.InternStr("True")
		ßinvoke := πg.InternStr("invoke")
		ßisinstance := πg.InternStr("isinstance")
		ßtuple := πg.InternStr("tuple")
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
			// line 15: from '__go__/syscall' import EINTR
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/syscall"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßEINTR, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßEINTR.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 18: def invoke(func, *args):
			πF.SetLineno(18)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "func", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("invoke", "build/src/__python__/_syscall.py", πTemp004, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfunc *πg.Object = πArgs[0]; _ = µfunc
				var µargs *πg.Object = πArgs[1]; _ = µargs
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µerr *πg.Object = πg.UnboundLocal; _ = µerr
				var πTemp001 bool
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 *πg.Object
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
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 19: while True:
					πF.SetLineno(19)
					πF.PushCheckpoint(2)
					πTemp001 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp001 {
						πF.PopCheckpoint()
						goto Label3
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πE != nil || !πTemp002 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 20: result = func(*args)
					πF.SetLineno(20)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, µfunc, nil, µargs, nil, nil); πE != nil {
						continue
					}
					µresult = πTemp003
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp004[0] = µresult
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					πTemp004[1] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp002, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					goto Label5
					// line 21: if isinstance(result, tuple):
					πF.SetLineno(21)
				Label4:
					// line 22: err = result[-1]
					πF.SetLineno(22)
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp005
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µresult, πTemp003); πE != nil {
						continue
					}
					µerr = πTemp005
					// line 23: result = result[:-1]
					πF.SetLineno(23)
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp005, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µresult, πTemp003); πE != nil {
						continue
					}
					µresult = πTemp005
					goto Label6
				Label5:
					// line 25: err = result
					πF.SetLineno(25)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					µerr = µresult
					// line 26: result = ()
					πF.SetLineno(26)
					πTemp003 = πg.NewTuple0().ToObject()
					µresult = πTemp003
					goto Label6
				Label6:
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µerr); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label7
					}
					goto Label8
					// line 27: if err:
					πF.SetLineno(27)
				Label7:
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßEINTR); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µerr, πTemp005); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label9
					}
					goto Label10
					// line 28: if err == EINTR:
					πF.SetLineno(28)
				Label9:
					// line 29: continue
					πF.SetLineno(29)
					continue
					goto Label10
				Label10:
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µerr, ßError, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp004[0] = πTemp005
					if πTemp003, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 30: raise OSError(err.Error())
					πF.SetLineno(30)
					πE = πF.Raise(πTemp005, nil, nil)
					continue
					goto Label8
				Label8:
					// line 31: return result
					πF.SetLineno(31)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πR = µresult
					continue
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
			if πE = πF.Globals().SetItem(πF, ßinvoke.ToObject(), πTemp001); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("_syscall", Code)
}
