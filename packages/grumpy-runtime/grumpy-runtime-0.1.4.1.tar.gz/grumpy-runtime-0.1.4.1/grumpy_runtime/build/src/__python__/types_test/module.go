package types_test
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/types_test.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßFunctionType := πg.InternStr("FunctionType")
		ßMethodType := πg.InternStr("MethodType")
		ßModuleType := πg.InternStr("ModuleType")
		ßStrType := πg.InternStr("StrType")
		ßStringType := πg.InternStr("StringType")
		ßTracebackType := πg.InternStr("TracebackType")
		ßTypeType := πg.InternStr("TypeType")
		ßUnboundMethodType := πg.InternStr("UnboundMethodType")
		ßtypes := πg.InternStr("types")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.Object
		_ = πTemp004
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 15: import types
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "types"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtypes.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 17: from '__go__/grumpy' import (FunctionType, MethodType, ModuleType, StrType,  # pylint: disable=g-multiple-import
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/grumpy"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßFunctionType, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFunctionType.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßMethodType, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMethodType.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßModuleType, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßModuleType.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßStrType, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStrType.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßTracebackType, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTracebackType.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßTypeType, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTypeType.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 21: assert types.FunctionType is FunctionType
			πF.SetLineno(21)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßFunctionType, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFunctionType); πE != nil {
				continue
			}
			πTemp001 = πg.GetBool(πTemp004 == πTemp003).ToObject()
			if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
				continue
			}
			// line 22: assert types.MethodType is MethodType
			πF.SetLineno(22)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßMethodType, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßMethodType); πE != nil {
				continue
			}
			πTemp001 = πg.GetBool(πTemp004 == πTemp003).ToObject()
			if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
				continue
			}
			// line 23: assert types.UnboundMethodType is MethodType
			πF.SetLineno(23)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßUnboundMethodType, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßMethodType); πE != nil {
				continue
			}
			πTemp001 = πg.GetBool(πTemp004 == πTemp003).ToObject()
			if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
				continue
			}
			// line 24: assert types.ModuleType is ModuleType
			πF.SetLineno(24)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßModuleType, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßModuleType); πE != nil {
				continue
			}
			πTemp001 = πg.GetBool(πTemp004 == πTemp003).ToObject()
			if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
				continue
			}
			// line 25: assert types.StringType is StrType
			πF.SetLineno(25)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßStringType, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßStrType); πE != nil {
				continue
			}
			πTemp001 = πg.GetBool(πTemp004 == πTemp003).ToObject()
			if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
				continue
			}
			// line 26: assert types.TracebackType is TracebackType
			πF.SetLineno(26)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTracebackType, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßTracebackType); πE != nil {
				continue
			}
			πTemp001 = πg.GetBool(πTemp004 == πTemp003).ToObject()
			if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
				continue
			}
			// line 27: assert types.TypeType is TypeType
			πF.SetLineno(27)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTypeType, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeType); πE != nil {
				continue
			}
			πTemp001 = πg.GetBool(πTemp004 == πTemp003).ToObject()
			if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("types_test", Code)
}
