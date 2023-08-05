package unittest_signals
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/unittest_signals.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßNone := πg.InternStr("None")
		ßTrue := πg.InternStr("True")
		ß_InterruptHandler := πg.InternStr("_InterruptHandler")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__unittest := πg.InternStr("__unittest")
		ß_interrupt_handler := πg.InternStr("_interrupt_handler")
		ß_results := πg.InternStr("_results")
		ßbool := πg.InternStr("bool")
		ßfunctools := πg.InternStr("functools")
		ßinstallHandler := πg.InternStr("installHandler")
		ßobject := πg.InternStr("object")
		ßpop := πg.InternStr("pop")
		ßregisterResult := πg.InternStr("registerResult")
		ßremoveHandler := πg.InternStr("removeHandler")
		ßremoveResult := πg.InternStr("removeResult")
		ßwraps := πg.InternStr("wraps")
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
		var πTemp006 []πg.Param
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
			// line 6: import functools
			πF.SetLineno(6)
			if πTemp002, πE = πg.ImportModule(πF, "functools"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßfunctools.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 7: wraps = functools.wraps
			πF.SetLineno(7)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßfunctools); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßwraps, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßwraps.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 9: __unittest = True
			πF.SetLineno(9)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__unittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 12: class _InterruptHandler(object):
			πF.SetLineno(12)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_InterruptHandler", "build/src/__python__/unittest_signals.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 13: pass
					πF.SetLineno(13)
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("_InterruptHandler").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_InterruptHandler.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 47: _results = {}
			πF.SetLineno(47)
			πTemp004 = πg.NewDict()
			πTemp001 = πTemp004.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_results.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 48: def registerResult(result):
			πF.SetLineno(48)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "result", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("registerResult", "build/src/__python__/unittest_signals.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µresult *πg.Object = πArgs[0]; _ = µresult
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
					// line 49: _results[result] = 1
					πF.SetLineno(49)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_results); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp003 = µresult
					if πE = πg.SetItem(πF, πTemp002, πTemp003, πTemp001); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßregisterResult.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 51: def removeResult(result):
			πF.SetLineno(51)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "result", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("removeResult", "build/src/__python__/unittest_signals.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µresult *πg.Object = πArgs[0]; _ = µresult
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
					// line 52: return bool(_results.pop(result, None))
					πF.SetLineno(52)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp002[0] = µresult
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_results); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßpop, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßbool); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp004
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßremoveResult.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 54: _interrupt_handler = None
			πF.SetLineno(54)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_interrupt_handler.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 55: def installHandler():
			πF.SetLineno(55)
			πTemp006 = make([]πg.Param, 0)
			πTemp005 = πg.NewFunction(πg.NewCode("installHandler", "build/src/__python__/unittest_signals.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 56: global _interrupt_handler
					πF.SetLineno(56)
					// line 57: pass
					πF.SetLineno(57)
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßinstallHandler.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 64: def removeHandler(method=None):
			πF.SetLineno(64)
			πTemp006 = make([]πg.Param, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[0] = πg.Param{Name: "method", Def: πTemp008}
			πTemp007 = πg.NewFunction(πg.NewCode("removeHandler", "build/src/__python__/unittest_signals.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmethod *πg.Object = πArgs[0]; _ = µmethod
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 65: pass
					πF.SetLineno(65)
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßremoveHandler.ToObject(), πTemp007); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("unittest_signals", Code)
}
