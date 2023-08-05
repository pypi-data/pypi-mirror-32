package math_test
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/math_test.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßAssertionError := πg.InternStr("AssertionError")
		ßRunTests := πg.InternStr("RunTests")
		ßTestDegrees := πg.InternStr("TestDegrees")
		ßTestFactorial := πg.InternStr("TestFactorial")
		ßTestFactorialError := πg.InternStr("TestFactorialError")
		ßTestLdexp := πg.InternStr("TestLdexp")
		ßTestLog := πg.InternStr("TestLog")
		ßTestRadians := πg.InternStr("TestRadians")
		ßValueError := πg.InternStr("ValueError")
		ß__main__ := πg.InternStr("__main__")
		ß__name__ := πg.InternStr("__name__")
		ßdegrees := πg.InternStr("degrees")
		ße := πg.InternStr("e")
		ßfactorial := πg.InternStr("factorial")
		ßldexp := πg.InternStr("ldexp")
		ßlog := πg.InternStr("log")
		ßmath := πg.InternStr("math")
		ßpi := πg.InternStr("pi")
		ßradians := πg.InternStr("radians")
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
		var πTemp011 bool
		_ = πTemp011
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 15: import math
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "math"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßmath.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 17: import weetest
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "weetest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßweetest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 23: def TestFactorial():
			πF.SetLineno(23)
			πTemp003 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("TestFactorial", "build/src/__python__/math_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
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
					// line 24: assert math.factorial(0) == 1
					πF.SetLineno(24)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(0).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßfactorial, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 25: assert math.factorial(1) == 1
					πF.SetLineno(25)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(1).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßfactorial, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 26: assert math.factorial(2) == 2
					πF.SetLineno(26)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(2).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßfactorial, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 27: assert math.factorial(3) == 6
					πF.SetLineno(27)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(3).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßfactorial, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewInt(6).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 28: assert math.factorial(4) == 24
					πF.SetLineno(28)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(4).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßfactorial, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewInt(24).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 29: assert math.factorial(5) == 120
					πF.SetLineno(29)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(5).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßfactorial, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewInt(120).ToObject()); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestFactorial.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 32: def TestFactorialError():
			πF.SetLineno(32)
			πTemp003 = make([]πg.Param, 0)
			πTemp004 = πg.NewFunction(πg.NewCode("TestFactorialError", "build/src/__python__/math_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 33: try:
					πF.SetLineno(33)
					πF.PushCheckpoint(2)
					// line 34: math.factorial(-1)
					πF.SetLineno(34)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßfactorial, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.PopCheckpoint()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					// line 38: raise AssertionError
					πF.SetLineno(38)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label1
				Label2:
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
						goto Label3
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 35: except ValueError:
					πF.SetLineno(35)
				Label3:
					// line 36: pass
					πF.SetLineno(36)
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 40: try:
					πF.SetLineno(40)
					πF.PushCheckpoint(5)
					// line 41: math.factorial(0.5)
					πF.SetLineno(41)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewFloat(0.5).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßfactorial, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.PopCheckpoint()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					// line 45: raise AssertionError
					πF.SetLineno(45)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
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
					// line 42: except ValueError:
					πF.SetLineno(42)
				Label6:
					// line 43: pass
					πF.SetLineno(43)
					πF.RestoreExc(nil, nil)
					goto Label4
				Label4:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßTestFactorialError.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 48: def TestLdexp():
			πF.SetLineno(48)
			πTemp003 = make([]πg.Param, 0)
			πTemp005 = πg.NewFunction(πg.NewCode("TestLdexp", "build/src/__python__/math_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
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
					// line 49: assert math.ldexp(1,1) == 2
					πF.SetLineno(49)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewInt(1).ToObject()
					πTemp002[1] = πg.NewInt(1).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßldexp, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 50: assert math.ldexp(1,2) == 4
					πF.SetLineno(50)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewInt(1).ToObject()
					πTemp002[1] = πg.NewInt(2).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßldexp, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 51: assert math.ldexp(1.5,1) == 3
					πF.SetLineno(51)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewFloat(1.5).ToObject()
					πTemp002[1] = πg.NewInt(1).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßldexp, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 52: assert math.ldexp(1.5,2) == 6
					πF.SetLineno(52)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewFloat(1.5).ToObject()
					πTemp002[1] = πg.NewInt(2).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßldexp, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewInt(6).ToObject()); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestLdexp.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 55: def TestLog():
			πF.SetLineno(55)
			πTemp003 = make([]πg.Param, 0)
			πTemp006 = πg.NewFunction(πg.NewCode("TestLog", "build/src/__python__/math_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
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
					// line 56: assert math.log(math.e) == 1
					πF.SetLineno(56)
					πTemp002 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ße, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßlog, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 57: assert math.log(2,2) == 1
					πF.SetLineno(57)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewInt(2).ToObject()
					πTemp002[1] = πg.NewInt(2).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßlog, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 58: assert math.log(10,10) == 1
					πF.SetLineno(58)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewInt(10).ToObject()
					πTemp002[1] = πg.NewInt(10).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßlog, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 59: assert math.log(100,10) == 2
					πF.SetLineno(59)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewInt(100).ToObject()
					πTemp002[1] = πg.NewInt(10).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßlog, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewInt(2).ToObject()); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestLog.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 62: def TestRadians():
			πF.SetLineno(62)
			πTemp003 = make([]πg.Param, 0)
			πTemp007 = πg.NewFunction(πg.NewCode("TestRadians", "build/src/__python__/math_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
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
					// line 63: assert math.radians(180) == math.pi
					πF.SetLineno(63)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(180).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßradians, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßpi, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, πTemp003, πTemp005); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 64: assert math.radians(360) == 2 * math.pi
					πF.SetLineno(64)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(360).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßradians, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßpi, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mul(πF, πg.NewInt(2).ToObject(), πTemp006); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, πTemp003, πTemp004); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestRadians.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 67: def TestDegrees():
			πF.SetLineno(67)
			πTemp003 = make([]πg.Param, 0)
			πTemp008 = πg.NewFunction(πg.NewCode("TestDegrees", "build/src/__python__/math_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
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
					// line 68: assert math.degrees(math.pi) == 180
					πF.SetLineno(68)
					πTemp002 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßpi, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßdegrees, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewInt(180).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 69: assert math.degrees(2 * math.pi) == 360
					πF.SetLineno(69)
					πTemp002 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßpi, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, πg.NewInt(2).ToObject(), πTemp005); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßdegrees, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewInt(360).ToObject()); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestDegrees.ToObject(), πTemp008); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp009, πE = πg.Eq(πF, πTemp010, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp011, πE = πg.IsTrue(πF, πTemp009); πE != nil {
				continue
			}
			if πTemp011 {
				goto Label1
			}
			goto Label2
			// line 72: if __name__ == '__main__':
			πF.SetLineno(72)
		Label1:
			// line 73: weetest.RunTests()
			πF.SetLineno(73)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßweetest); πE != nil {
				continue
			}
			if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßRunTests, nil); πE != nil {
				continue
			}
			if πTemp009, πE = πTemp010.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("math_test", Code)
}
