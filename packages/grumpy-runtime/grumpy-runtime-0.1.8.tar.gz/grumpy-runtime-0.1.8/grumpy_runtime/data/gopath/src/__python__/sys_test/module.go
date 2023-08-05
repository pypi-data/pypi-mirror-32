package sys_test
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/sys_test.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßFalse := πg.InternStr("False")
		ßNone := πg.InternStr("None")
		ßRunTests := πg.InternStr("RunTests")
		ßRuntimeError := πg.InternStr("RuntimeError")
		ßSystemExit := πg.InternStr("SystemExit")
		ßTestArgv := πg.InternStr("TestArgv")
		ßTestExcClear := πg.InternStr("TestExcClear")
		ßTestExcInfoNoException := πg.InternStr("TestExcInfoNoException")
		ßTestExcInfoWithException := πg.InternStr("TestExcInfoWithException")
		ßTestExitCode := πg.InternStr("TestExitCode")
		ßTestExitEmpty := πg.InternStr("TestExitEmpty")
		ßTestExitInvalidArgs := πg.InternStr("TestExitInvalidArgs")
		ßTestGetFrame := πg.InternStr("TestGetFrame")
		ßTestMaxInt := πg.InternStr("TestMaxInt")
		ßTestSysModules := πg.InternStr("TestSysModules")
		ßTracebackType := πg.InternStr("TracebackType")
		ßTypeError := πg.InternStr("TypeError")
		ßValueError := πg.InternStr("ValueError")
		ß__main__ := πg.InternStr("__main__")
		ß__name__ := πg.InternStr("__name__")
		ß_getframe := πg.InternStr("_getframe")
		ßall := πg.InternStr("all")
		ßany := πg.InternStr("any")
		ßargv := πg.InternStr("argv")
		ßco_name := πg.InternStr("co_name")
		ßcode := πg.InternStr("code")
		ßexc_clear := πg.InternStr("exc_clear")
		ßexc_info := πg.InternStr("exc_info")
		ßexit := πg.InternStr("exit")
		ßf_code := πg.InternStr("f_code")
		ßisinstance := πg.InternStr("isinstance")
		ßmaxint := πg.InternStr("maxint")
		ßmodules := πg.InternStr("modules")
		ßstr := πg.InternStr("str")
		ßsys := πg.InternStr("sys")
		ßtypes := πg.InternStr("types")
		ßweetest := πg.InternStr("weetest")
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
		var πTemp015 bool
		_ = πTemp015
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 17: import sys
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: import types
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "types"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtypes.ToObject(), πTemp001); πE != nil {
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
			// line 23: def TestArgv():
			πF.SetLineno(23)
			πTemp003 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("TestArgv", "build/src/__python__/sys_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 24: assert sys.argv
					πF.SetLineno(24)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßargv, nil); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßTestArgv.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 27: def TestMaxInt():
			πF.SetLineno(27)
			πTemp003 = make([]πg.Param, 0)
			πTemp004 = πg.NewFunction(πg.NewCode("TestMaxInt", "build/src/__python__/sys_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 28: assert sys.maxint > 2000000000
					πF.SetLineno(28)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmaxint, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GT(πF, πTemp003, πg.NewInt(2000000000).ToObject()); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestMaxInt.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 31: def TestSysModules():
			πF.SetLineno(31)
			πTemp003 = make([]πg.Param, 0)
			πTemp005 = πg.NewFunction(πg.NewCode("TestSysModules", "build/src/__python__/sys_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
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
					// line 32: assert sys.modules['sys'] is not None
					πF.SetLineno(32)
					πTemp002 = ßsys.ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßmodules, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp003 != πTemp002).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßTestSysModules.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 35: def TestExcClear():
			πF.SetLineno(35)
			πTemp003 = make([]πg.Param, 0)
			πTemp006 = πg.NewFunction(πg.NewCode("TestExcClear", "build/src/__python__/sys_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.BaseException
				_ = πTemp002
				var πTemp003 *πg.Traceback
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
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
					// line 36: try:
					πF.SetLineno(36)
					πF.PushCheckpoint(2)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
						continue
					}
					// line 37: raise RuntimeError
					πF.SetLineno(37)
					πE = πF.Raise(πTemp001, nil, nil)
					continue
					πF.PopCheckpoint()
					// line 43: assert False
					πF.SetLineno(43)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp002, πTemp003 = πF.ExcInfo()
					goto Label3
					// line 38: except:
					πF.SetLineno(38)
				Label3:
					// line 39: assert all(sys.exc_info()), sys.exc_info()
					πF.SetLineno(39)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßexc_info, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp005 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßexc_info, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp006.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp005[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßall); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πg.Assert(πF, πTemp006, πTemp001); πE != nil {
						continue
					}
					// line 40: sys.exc_clear()
					πF.SetLineno(40)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßexc_clear, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 41: assert not any(sys.exc_info())
					πF.SetLineno(41)
					πTemp005 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßexc_info, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp006.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp005[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßany); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp007).ToObject()
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestExcClear.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 46: def TestExcInfoNoException():
			πF.SetLineno(46)
			πTemp003 = make([]πg.Param, 0)
			πTemp007 = πg.NewFunction(πg.NewCode("TestExcInfoNoException", "build/src/__python__/sys_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
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
					// line 47: assert sys.exc_info() == (None, None, None)
					πF.SetLineno(47)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßexc_info, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple3(πTemp004, πTemp005, πTemp006).ToObject()
					if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp003); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestExcInfoNoException.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 50: def TestExcInfoWithException():
			πF.SetLineno(50)
			πTemp003 = make([]πg.Param, 0)
			πTemp008 = πg.NewFunction(πg.NewCode("TestExcInfoWithException", "build/src/__python__/sys_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µt *πg.Object = πg.UnboundLocal; _ = µt
				var µe *πg.Object = πg.UnboundLocal; _ = µe
				var µtb *πg.Object = πg.UnboundLocal; _ = µtb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.BaseException
				_ = πTemp002
				var πTemp003 *πg.Traceback
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 51: try:
					πF.SetLineno(51)
					πF.PushCheckpoint(2)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
						continue
					}
					// line 52: raise RuntimeError
					πF.SetLineno(52)
					πE = πF.Raise(πTemp001, nil, nil)
					continue
					πF.PopCheckpoint()
					// line 56: assert False
					πF.SetLineno(56)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp002, πTemp003 = πF.ExcInfo()
					goto Label3
					// line 53: except:
					πF.SetLineno(53)
				Label3:
					// line 54: t, e, tb = sys.exc_info()
					πF.SetLineno(54)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßexc_info, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp001); πE != nil {
						continue
					}
					µt = πTemp004
					µe = πTemp005
					µtb = πTemp006
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 57: assert t is RuntimeError
					πF.SetLineno(57)
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µt == πTemp004).ToObject()
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 58: assert isinstance(e, t)
					πF.SetLineno(58)
					πTemp007 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
						continue
					}
					πTemp007[0] = µe
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					πTemp007[1] = µt
					if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πg.Assert(πF, πTemp004, nil); πE != nil {
						continue
					}
					// line 59: assert isinstance(tb, types.TracebackType)
					πF.SetLineno(59)
					πTemp007 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
						continue
					}
					πTemp007[0] = µtb
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßTracebackType, nil); πE != nil {
						continue
					}
					πTemp007[1] = πTemp004
					if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πg.Assert(πF, πTemp004, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestExcInfoWithException.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 62: def TestExitEmpty():
			πF.SetLineno(62)
			πTemp003 = make([]πg.Param, 0)
			πTemp009 = πg.NewFunction(πg.NewCode("TestExitEmpty", "build/src/__python__/sys_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µe *πg.Object = πg.UnboundLocal; _ = µe
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.BaseException
				_ = πTemp003
				var πTemp004 *πg.Traceback
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
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 63: try:
					πF.SetLineno(63)
					πF.PushCheckpoint(2)
					// line 64: sys.exit()
					πF.SetLineno(64)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßexit, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp003, πTemp004 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßSystemExit); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsInstance(πF, πTemp003.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label3
					}
					goto Label4
					// line 65: except SystemExit as e:
					πF.SetLineno(65)
				Label3:
					µe = πTemp003.ToObject()
					// line 66: assert e.code == None, e.code  # pylint: disable=g-equals-none
					πF.SetLineno(66)
					if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µe, ßcode, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µe, ßcode, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp006, πTemp007); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp002, πTemp001); πE != nil {
						continue
					}
					πF.RestoreExc(nil, nil)
					goto Label1
					// line 67: except:
					πF.SetLineno(67)
				Label4:
					// line 68: assert False
					πF.SetLineno(68)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestExitEmpty.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 71: def TestExitCode():
			πF.SetLineno(71)
			πTemp003 = make([]πg.Param, 0)
			πTemp010 = πg.NewFunction(πg.NewCode("TestExitCode", "build/src/__python__/sys_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µe *πg.Object = πg.UnboundLocal; _ = µe
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
				var πTemp007 *πg.Object
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 72: try:
					πF.SetLineno(72)
					πF.PushCheckpoint(2)
					// line 73: sys.exit(42)
					πF.SetLineno(73)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewInt(42).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßexit, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSystemExit); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					goto Label4
					// line 74: except SystemExit as e:
					πF.SetLineno(74)
				Label3:
					µe = πTemp004.ToObject()
					// line 75: assert e.code == 42, e.code
					πF.SetLineno(75)
					if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µe, ßcode, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µe, ßcode, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp007, πg.NewInt(42).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp003, πTemp002); πE != nil {
						continue
					}
					πF.RestoreExc(nil, nil)
					goto Label1
					// line 76: except:
					πF.SetLineno(76)
				Label4:
					// line 77: assert False
					πF.SetLineno(77)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßTestExitCode.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 80: def TestExitInvalidArgs():
			πF.SetLineno(80)
			πTemp003 = make([]πg.Param, 0)
			πTemp011 = πg.NewFunction(πg.NewCode("TestExitInvalidArgs", "build/src/__python__/sys_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µe *πg.Object = πg.UnboundLocal; _ = µe
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
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 81: try:
					πF.SetLineno(81)
					πF.PushCheckpoint(2)
					// line 82: sys.exit(1, 2, 3)
					πF.SetLineno(82)
					πTemp001 = πF.MakeArgs(3)
					πTemp001[0] = πg.NewInt(1).ToObject()
					πTemp001[1] = πg.NewInt(2).ToObject()
					πTemp001[2] = πg.NewInt(3).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßexit, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					goto Label4
					// line 83: except TypeError as e:
					πF.SetLineno(83)
				Label3:
					µe = πTemp004.ToObject()
					// line 84: assert str(e) == 'exit() takes 1 arguments (3 given)', str(e)
					πF.SetLineno(84)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
						continue
					}
					πTemp001[0] = µe
					if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
						continue
					}
					πTemp001[0] = µe
					if πTemp007, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Eq(πF, πTemp008, πg.NewStr("exit() takes 1 arguments (3 given)").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp002, πTemp003); πE != nil {
						continue
					}
					πF.RestoreExc(nil, nil)
					goto Label1
					// line 85: except:
					πF.SetLineno(85)
				Label4:
					// line 86: assert False
					πF.SetLineno(86)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßTestExitInvalidArgs.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 89: def TestGetFrame():
			πF.SetLineno(89)
			πTemp003 = make([]πg.Param, 0)
			πTemp012 = πg.NewFunction(πg.NewCode("TestGetFrame", "build/src/__python__/sys_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp007 *πg.Object
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 90: try:
					πF.SetLineno(90)
					πF.PushCheckpoint(2)
					// line 91: sys._getframe(42, 42)
					πF.SetLineno(91)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewInt(42).ToObject()
					πTemp001[1] = πg.NewInt(42).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß_getframe, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.PopCheckpoint()
					// line 95: assert False
					πF.SetLineno(95)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
						continue
					}
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
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
					// line 92: except TypeError:
					πF.SetLineno(92)
				Label3:
					// line 93: pass
					πF.SetLineno(93)
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 96: try:
					πF.SetLineno(96)
					πF.PushCheckpoint(5)
					// line 97: sys._getframe(2000000000)
					πF.SetLineno(97)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewInt(2000000000).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß_getframe, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.PopCheckpoint()
					// line 101: assert False
					πF.SetLineno(101)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
						continue
					}
					goto Label4
				Label5:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
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
					// line 98: except ValueError:
					πF.SetLineno(98)
				Label6:
					// line 99: pass
					πF.SetLineno(99)
					πF.RestoreExc(nil, nil)
					goto Label4
				Label4:
					// line 102: assert sys._getframe().f_code.co_name == '_getframe'
					πF.SetLineno(102)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp003, ß_getframe, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp007.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßf_code, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp007, ßco_name, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp003, ß_getframe.ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
						continue
					}
					// line 103: assert sys._getframe(1).f_code.co_name == 'TestGetFrame'
					πF.SetLineno(103)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewInt(1).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp003, ß_getframe, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßf_code, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp007, ßco_name, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp003, ßTestGetFrame.ToObject()); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßTestGetFrame.ToObject(), πTemp012); πE != nil {
				continue
			}
			if πTemp014, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp013, πE = πg.Eq(πF, πTemp014, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp015, πE = πg.IsTrue(πF, πTemp013); πE != nil {
				continue
			}
			if πTemp015 {
				goto Label1
			}
			goto Label2
			// line 106: if __name__ == '__main__':
			πF.SetLineno(106)
		Label1:
			// line 108: weetest.RunTests()
			πF.SetLineno(108)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßweetest); πE != nil {
				continue
			}
			if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßRunTests, nil); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp014.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("sys_test", Code)
}
