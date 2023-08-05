package time_test
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/time_test.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßgot := πg.InternStr("got")
		ßlocaltime := πg.InternStr("localtime")
		ßmktime := πg.InternStr("mktime")
		ßtime := πg.InternStr("time")
		ßtime_struct := πg.InternStr("time_struct")
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
		var πTemp006 *πg.Object
		_ = πTemp006
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 15: import time
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "time"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtime.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 17: assert time.time() > 1000000000
			πF.SetLineno(17)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßtime, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πg.GT(πF, πTemp003, πg.NewInt(1000000000).ToObject()); πE != nil {
				continue
			}
			if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
				continue
			}
			// line 18: assert time.time() < 3000000000
			πF.SetLineno(18)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßtime, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πg.LT(πF, πTemp003, πg.NewInt(3000000000).ToObject()); πE != nil {
				continue
			}
			if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
				continue
			}
			// line 20: time_struct = (1999, 9, 19, 0, 0, 0, 6, 262, 0)
			πF.SetLineno(20)
			πTemp002 = make([]*πg.Object, 9)
			πTemp002[0] = πg.NewInt(1999).ToObject()
			πTemp002[1] = πg.NewInt(9).ToObject()
			πTemp002[2] = πg.NewInt(19).ToObject()
			πTemp002[3] = πg.NewInt(0).ToObject()
			πTemp002[4] = πg.NewInt(0).ToObject()
			πTemp002[5] = πg.NewInt(0).ToObject()
			πTemp002[6] = πg.NewInt(6).ToObject()
			πTemp002[7] = πg.NewInt(262).ToObject()
			πTemp002[8] = πg.NewInt(0).ToObject()
			πTemp001 = πg.NewTuple(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßtime_struct.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 21: got = time.localtime(time.mktime(time_struct))
			πF.SetLineno(21)
			πTemp002 = πF.MakeArgs(1)
			πTemp005 = πF.MakeArgs(1)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßtime_struct); πE != nil {
				continue
			}
			πTemp005[0] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßmktime, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp005)
			πTemp002[0] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßlocaltime, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßgot.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 22: assert got == time_struct, got
			πF.SetLineno(22)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßgot); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßgot); πE != nil {
				continue
			}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßtime_struct); πE != nil {
				continue
			}
			if πTemp003, πE = πg.Eq(πF, πTemp004, πTemp006); πE != nil {
				continue
			}
			if πE = πg.Assert(πF, πTemp003, πTemp001); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("time_test", Code)
}
