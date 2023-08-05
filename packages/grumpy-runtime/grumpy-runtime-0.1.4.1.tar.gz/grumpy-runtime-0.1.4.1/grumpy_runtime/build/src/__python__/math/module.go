package math
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/math.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßAbs := πg.InternStr("Abs")
		ßAcos := πg.InternStr("Acos")
		ßAcosh := πg.InternStr("Acosh")
		ßAsin := πg.InternStr("Asin")
		ßAsinh := πg.InternStr("Asinh")
		ßAtan := πg.InternStr("Atan")
		ßAtan2 := πg.InternStr("Atan2")
		ßAtanh := πg.InternStr("Atanh")
		ßCeil := πg.InternStr("Ceil")
		ßCopysign := πg.InternStr("Copysign")
		ßCos := πg.InternStr("Cos")
		ßCosh := πg.InternStr("Cosh")
		ßE := πg.InternStr("E")
		ßErf := πg.InternStr("Erf")
		ßErfc := πg.InternStr("Erfc")
		ßExp := πg.InternStr("Exp")
		ßExp2 := πg.InternStr("Exp2")
		ßExpm1 := πg.InternStr("Expm1")
		ßFloor := πg.InternStr("Floor")
		ßFrexp := πg.InternStr("Frexp")
		ßGamma := πg.InternStr("Gamma")
		ßHypot := πg.InternStr("Hypot")
		ßIsInf := πg.InternStr("IsInf")
		ßIsNaN := πg.InternStr("IsNaN")
		ßLgamma := πg.InternStr("Lgamma")
		ßLog := πg.InternStr("Log")
		ßLog10 := πg.InternStr("Log10")
		ßLog1p := πg.InternStr("Log1p")
		ßMod := πg.InternStr("Mod")
		ßModf := πg.InternStr("Modf")
		ßNone := πg.InternStr("None")
		ßPi := πg.InternStr("Pi")
		ßPow := πg.InternStr("Pow")
		ßSin := πg.InternStr("Sin")
		ßSinh := πg.InternStr("Sinh")
		ßSqrt := πg.InternStr("Sqrt")
		ßTan := πg.InternStr("Tan")
		ßTanh := πg.InternStr("Tanh")
		ßTrunc := πg.InternStr("Trunc")
		ßTypeError := πg.InternStr("TypeError")
		ßValueError := πg.InternStr("ValueError")
		ßacos := πg.InternStr("acos")
		ßacosh := πg.InternStr("acosh")
		ßasin := πg.InternStr("asin")
		ßasinh := πg.InternStr("asinh")
		ßatan := πg.InternStr("atan")
		ßatan2 := πg.InternStr("atan2")
		ßatanh := πg.InternStr("atanh")
		ßceil := πg.InternStr("ceil")
		ßcopysign := πg.InternStr("copysign")
		ßcos := πg.InternStr("cos")
		ßcosh := πg.InternStr("cosh")
		ßdegrees := πg.InternStr("degrees")
		ße := πg.InternStr("e")
		ßerf := πg.InternStr("erf")
		ßerfc := πg.InternStr("erfc")
		ßexp := πg.InternStr("exp")
		ßexpm1 := πg.InternStr("expm1")
		ßfabs := πg.InternStr("fabs")
		ßfactorial := πg.InternStr("factorial")
		ßfloat := πg.InternStr("float")
		ßfloor := πg.InternStr("floor")
		ßfmod := πg.InternStr("fmod")
		ßfrexp := πg.InternStr("frexp")
		ßgamma := πg.InternStr("gamma")
		ßhypot := πg.InternStr("hypot")
		ßint := πg.InternStr("int")
		ßisinf := πg.InternStr("isinf")
		ßisnan := πg.InternStr("isnan")
		ßldexp := πg.InternStr("ldexp")
		ßlgamma := πg.InternStr("lgamma")
		ßlog := πg.InternStr("log")
		ßlog10 := πg.InternStr("log10")
		ßlog1p := πg.InternStr("log1p")
		ßmodf := πg.InternStr("modf")
		ßpi := πg.InternStr("pi")
		ßpow := πg.InternStr("pow")
		ßradians := πg.InternStr("radians")
		ßrange := πg.InternStr("range")
		ßsin := πg.InternStr("sin")
		ßsinh := πg.InternStr("sinh")
		ßsqrt := πg.InternStr("sqrt")
		ßtan := πg.InternStr("tan")
		ßtanh := πg.InternStr("tanh")
		ßtrunc := πg.InternStr("trunc")
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
		var πTemp011 *πg.Object
		_ = πTemp011
		var πTemp012 *πg.Object
		_ = πTemp012
		var πTemp013 *πg.Object
		_ = πTemp013
		var πTemp014 *πg.Object
		_ = πTemp014
		var πTemp015 *πg.Object
		_ = πTemp015
		var πTemp016 *πg.Object
		_ = πTemp016
		var πTemp017 *πg.Object
		_ = πTemp017
		var πTemp018 *πg.Object
		_ = πTemp018
		var πTemp019 *πg.Object
		_ = πTemp019
		var πTemp020 *πg.Object
		_ = πTemp020
		var πTemp021 *πg.Object
		_ = πTemp021
		var πTemp022 *πg.Object
		_ = πTemp022
		var πTemp023 *πg.Object
		_ = πTemp023
		var πTemp024 *πg.Object
		_ = πTemp024
		var πTemp025 *πg.Object
		_ = πTemp025
		var πTemp026 *πg.Object
		_ = πTemp026
		var πTemp027 *πg.Object
		_ = πTemp027
		var πTemp028 *πg.Object
		_ = πTemp028
		var πTemp029 *πg.Object
		_ = πTemp029
		var πTemp030 *πg.Object
		_ = πTemp030
		var πTemp031 *πg.Object
		_ = πTemp031
		var πTemp032 *πg.Object
		_ = πTemp032
		var πTemp033 *πg.Object
		_ = πTemp033
		var πTemp034 *πg.Object
		_ = πTemp034
		var πTemp035 *πg.Object
		_ = πTemp035
		var πTemp036 *πg.Object
		_ = πTemp036
		var πTemp037 *πg.Object
		_ = πTemp037
		var πTemp038 *πg.Object
		_ = πTemp038
		var πTemp039 *πg.Object
		_ = πTemp039
		var πTemp040 *πg.Object
		_ = πTemp040
		var πTemp041 *πg.Object
		_ = πTemp041
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 15: from '__go__/math' import (Pi, E, Ceil, Copysign, Abs, Floor, Mod, Frexp, IsInf,
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/math"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßPi, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPi.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßE, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßE.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßCeil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCeil.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßCopysign, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCopysign.ToObject(), πTemp003); πE != nil {
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
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßFloor, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFloor.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßMod, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMod.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßFrexp, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFrexp.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßIsInf, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßIsInf.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßIsNaN, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßIsNaN.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßExp2, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßExp2.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßModf, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßModf.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßTrunc, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTrunc.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßExp, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßExp.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßExpm1, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßExpm1.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßLog, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLog.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßLog1p, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLog1p.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßLog10, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLog10.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßPow, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPow.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßSqrt, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSqrt.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßAcos, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßAcos.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßAsin, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßAsin.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßAtan, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßAtan.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßAtan2, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßAtan2.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßHypot, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßHypot.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßSin, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSin.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßCos, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCos.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßTan, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTan.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßAcosh, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßAcosh.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßAsinh, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßAsinh.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßAtanh, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßAtanh.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßSinh, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSinh.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßCosh, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCosh.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßTanh, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTanh.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßErf, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßErf.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßErfc, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßErfc.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßGamma, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßGamma.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßLgamma, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLgamma.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 22: pi = Pi
			πF.SetLineno(22)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßPi); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßpi.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 25: e = E
			πF.SetLineno(25)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßE); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ße.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 30: def ceil(x):
			πF.SetLineno(30)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("ceil", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 31: return Ceil(float(x))
					πF.SetLineno(31)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßCeil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßceil.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 34: def copysign(x, y):
			πF.SetLineno(34)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp004[1] = πg.Param{Name: "y", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("copysign", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µy *πg.Object = πArgs[1]; _ = µy
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
					// line 35: return Copysign(float(x), float(y))
					πF.SetLineno(35)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					πTemp002[0] = µy
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[1] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßCopysign); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßcopysign.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 38: def fabs(x):
			πF.SetLineno(38)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("fabs", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 39: return Abs(float(x))
					πF.SetLineno(39)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßAbs); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßfabs.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 42: def factorial(x):
			πF.SetLineno(42)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("factorial", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µxi *πg.Object = πg.UnboundLocal; _ = µxi
				var µxf *πg.Object = πg.UnboundLocal; _ = µxf
				var µacc *πg.Object = πg.UnboundLocal; _ = µacc
				var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
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
				var πTemp008 bool
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 17: goto Label17
					case 2: goto Label2
					case 18: goto Label18
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 43: try:
					πF.SetLineno(43)
					πF.PushCheckpoint(2)
					// line 44: xi = int(x)
					πF.SetLineno(44)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µxi = πTemp003
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
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 45: except TypeError:
					πF.SetLineno(45)
				Label3:
					// line 46: xi = None
					πF.SetLineno(46)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µxi = πTemp002
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 48: try:
					πF.SetLineno(48)
					πF.PushCheckpoint(5)
					// line 49: xf = float(x)
					πF.SetLineno(49)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					if πTemp002, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µxf = πTemp003
					πF.PopCheckpoint()
					goto Label4
				Label5:
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
						goto Label6
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 50: except TypeError:
					πF.SetLineno(50)
				Label6:
					// line 51: xf = None
					πF.SetLineno(51)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µxf = πTemp002
					πF.RestoreExc(nil, nil)
					goto Label4
				Label4:
					if πE = πg.CheckLocal(πF, µxi, "xi"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µxi == πTemp003).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label7
					}
					if πE = πg.CheckLocal(πF, µxf, "xf"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µxf == πTemp007).ToObject()
					πTemp002 = πTemp003
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label8
					}
					if πE = πg.CheckLocal(πF, µxi, "xi"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µxi == πTemp007).ToObject()
					πTemp002 = πTemp003
				Label8:
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label9
					}
					if πE = πg.CheckLocal(πF, µxf, "xf"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µxf == πTemp003).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label10
					}
					if πE = πg.CheckLocal(πF, µxf, "xf"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µxi, "xi"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.NE(πF, µxf, µxi); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label11
					}
					goto Label12
					// line 53: if xi is None:
					πF.SetLineno(53)
				Label7:
					// line 54: xi = int(xf)
					πF.SetLineno(54)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µxf, "xf"); πE != nil {
						continue
					}
					πTemp001[0] = µxf
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µxi = πTemp003
					if πE = πg.CheckLocal(πF, µxi, "xi"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µxf, "xf"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.NE(πF, µxi, µxf); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label13
					}
					goto Label14
					// line 55: if xi != xf:
					πF.SetLineno(55)
				Label13:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("factorial() only accepts integral values").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 56: raise ValueError("factorial() only accepts integral values")
					πF.SetLineno(56)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label14
				Label14:
					goto Label12
					// line 57: elif xf is None and xi is None:
					πF.SetLineno(57)
				Label9:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("an integer is required").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 58: raise TypeError("an integer is required")
					πF.SetLineno(58)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label12
					// line 59: elif xf is None:
					πF.SetLineno(59)
				Label10:
					// line 60: pass
					πF.SetLineno(60)
					goto Label12
					// line 61: elif xf != xi:
					πF.SetLineno(61)
				Label11:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("factorial() only accepts integral values").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 62: raise ValueError("factorial() only accepts integral values")
					πF.SetLineno(62)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label12
				Label12:
					// line 64: x = xi
					πF.SetLineno(64)
					if πE = πg.CheckLocal(πF, µxi, "xi"); πE != nil {
						continue
					}
					µx = µxi
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LT(πF, µx, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label15
					}
					goto Label16
					// line 66: if x < 0:
					πF.SetLineno(66)
				Label15:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("factorial() not defined for negative values").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 67: raise ValueError("factorial() not defined for negative values")
					πF.SetLineno(67)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label16
				Label16:
					// line 69: acc = 1
					πF.SetLineno(69)
					µacc = πg.NewInt(1).ToObject()
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µx, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Iter(πF, πTemp007); πE != nil {
						continue
					}
					πF.PushCheckpoint(18)
					πTemp006 = false
				Label17:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label19
					}
					if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp008 = !isStop
					} else {
						πTemp008 = true
						µvalue = πTemp003
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(17)            
					// line 72: acc *= value
					πF.SetLineno(72)
					if πE = πg.CheckLocal(πF, µacc, "acc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IMul(πF, µacc, µvalue); πE != nil {
						continue
					}
					µacc = πTemp003
					continue
				Label18:
					if πE != nil || πR != nil {
						continue
					}
				Label19:
					// line 74: return acc
					πF.SetLineno(74)
					if πE = πg.CheckLocal(πF, µacc, "acc"); πE != nil {
						continue
					}
					πR = µacc
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßfactorial.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 77: def floor(x):
			πF.SetLineno(77)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("floor", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 78: return Floor(float(x))
					πF.SetLineno(78)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßFloor); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßfloor.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 81: def fmod(x):
			πF.SetLineno(81)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("fmod", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 82: return Mod(float(x))
					πF.SetLineno(82)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßMod); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßfmod.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 85: def frexp(x):
			πF.SetLineno(85)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("frexp", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 86: return Frexp(float(x))
					πF.SetLineno(86)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßFrexp); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßfrexp.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 94: def isinf(x):
			πF.SetLineno(94)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("isinf", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 95: return IsInf(float(x), 0)
					πF.SetLineno(95)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					πTemp001[1] = πg.NewInt(0).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßIsInf); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßisinf.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 98: def isnan(x):
			πF.SetLineno(98)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp011 = πg.NewFunction(πg.NewCode("isnan", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 99: return IsNaN(float(x))
					πF.SetLineno(99)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßIsNaN); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßisnan.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 102: def ldexp(x, i):
			πF.SetLineno(102)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp004[1] = πg.Param{Name: "i", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("ldexp", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µi *πg.Object = πArgs[1]; _ = µi
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 103: return float(x) * Exp2(float(i))
					πF.SetLineno(103)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp002 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp005[0] = µi
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp002[0] = πTemp006
					if πTemp003, πE = πg.ResolveGlobal(πF, ßExp2); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Mul(πF, πTemp004, πTemp006); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßldexp.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 106: def modf(x):
			πF.SetLineno(106)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp013 = πg.NewFunction(πg.NewCode("modf", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µa *πg.Object = πg.UnboundLocal; _ = µa
				var µb *πg.Object = πg.UnboundLocal; _ = µb
				var πTemp001 []*πg.Object
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
					// line 108: a, b = Modf(float(x))
					πF.SetLineno(108)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßModf); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}}}, πTemp004); πE != nil {
						continue
					}
					µa = πTemp003
					µb = πTemp005
					// line 109: return b, a
					πF.SetLineno(109)
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µb, µa).ToObject()
					πR = πTemp003
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßmodf.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 112: def trunc(x):
			πF.SetLineno(112)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp014 = πg.NewFunction(πg.NewCode("trunc", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 113: return Trunc(float(x))
					πF.SetLineno(113)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrunc); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßtrunc.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 118: def exp(x):
			πF.SetLineno(118)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp015 = πg.NewFunction(πg.NewCode("exp", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 119: return Exp(float(x))
					πF.SetLineno(119)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßExp); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßexp.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 122: def expm1(x):
			πF.SetLineno(122)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp016 = πg.NewFunction(πg.NewCode("expm1", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 123: return Expm1(float(x))
					πF.SetLineno(123)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßExpm1); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßexpm1.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 126: def log(x, b=None):
			πF.SetLineno(126)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			if πTemp018, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "b", Def: πTemp018}
			πTemp017 = πg.NewFunction(πg.NewCode("log", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
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
					default: panic("unexpected function state")
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µb == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 127: if b is None:
					πF.SetLineno(127)
				Label1:
					// line 128: return Log(float(x))
					πF.SetLineno(128)
					πTemp004 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp005[0] = µx
					if πTemp001, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp004[0] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßLog); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πR = πTemp002
					continue
					goto Label2
				Label2:
					// line 132: return Log(float(x)) / Log(float(b))
					πF.SetLineno(132)
					πTemp004 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp005[0] = µx
					if πTemp002, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp004[0] = πTemp006
					if πTemp002, πE = πg.ResolveGlobal(πF, ßLog); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp004 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					πTemp005[0] = µb
					if πTemp002, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp004[0] = πTemp007
					if πTemp002, πE = πg.ResolveGlobal(πF, ßLog); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp001, πE = πg.Div(πF, πTemp006, πTemp007); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßlog.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 135: def log1p(x):
			πF.SetLineno(135)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp018 = πg.NewFunction(πg.NewCode("log1p", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 136: return Log1p(float(x))
					πF.SetLineno(136)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßLog1p); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßlog1p.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 139: def log10(x):
			πF.SetLineno(139)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp019 = πg.NewFunction(πg.NewCode("log10", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 140: return Log10(float(x))
					πF.SetLineno(140)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßLog10); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßlog10.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 143: def pow(x, y):
			πF.SetLineno(143)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp004[1] = πg.Param{Name: "y", Def: nil}
			πTemp020 = πg.NewFunction(πg.NewCode("pow", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µy *πg.Object = πArgs[1]; _ = µy
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
					// line 144: return Pow(float(x), float(y))
					πF.SetLineno(144)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					πTemp002[0] = µy
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[1] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßPow); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßpow.ToObject(), πTemp020); πE != nil {
				continue
			}
			// line 147: def sqrt(x):
			πF.SetLineno(147)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp021 = πg.NewFunction(πg.NewCode("sqrt", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 148: return Sqrt(float(x))
					πF.SetLineno(148)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßSqrt); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßsqrt.ToObject(), πTemp021); πE != nil {
				continue
			}
			// line 153: def acos(x):
			πF.SetLineno(153)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp022 = πg.NewFunction(πg.NewCode("acos", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 154: return Acos(float(x))
					πF.SetLineno(154)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßAcos); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßacos.ToObject(), πTemp022); πE != nil {
				continue
			}
			// line 157: def asin(x):
			πF.SetLineno(157)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp023 = πg.NewFunction(πg.NewCode("asin", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 158: return Asin(float(x))
					πF.SetLineno(158)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßAsin); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßasin.ToObject(), πTemp023); πE != nil {
				continue
			}
			// line 161: def atan(x):
			πF.SetLineno(161)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp024 = πg.NewFunction(πg.NewCode("atan", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 162: return Atan(float(x))
					πF.SetLineno(162)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßAtan); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßatan.ToObject(), πTemp024); πE != nil {
				continue
			}
			// line 165: def atan2(y, x):
			πF.SetLineno(165)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "y", Def: nil}
			πTemp004[1] = πg.Param{Name: "x", Def: nil}
			πTemp025 = πg.NewFunction(πg.NewCode("atan2", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µy *πg.Object = πArgs[0]; _ = µy
				var µx *πg.Object = πArgs[1]; _ = µx
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
					// line 166: return Atan2(float(y), float(x))
					πF.SetLineno(166)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					πTemp002[0] = µy
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[1] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßAtan2); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßatan2.ToObject(), πTemp025); πE != nil {
				continue
			}
			// line 169: def cos(x):
			πF.SetLineno(169)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp026 = πg.NewFunction(πg.NewCode("cos", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 170: return Cos(float(x))
					πF.SetLineno(170)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßCos); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßcos.ToObject(), πTemp026); πE != nil {
				continue
			}
			// line 173: def hypot(x, y):
			πF.SetLineno(173)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp004[1] = πg.Param{Name: "y", Def: nil}
			πTemp027 = πg.NewFunction(πg.NewCode("hypot", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µy *πg.Object = πArgs[1]; _ = µy
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
					// line 174: return Hypot(float(x), float(y))
					πF.SetLineno(174)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					πTemp002[0] = µy
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[1] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßHypot); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßhypot.ToObject(), πTemp027); πE != nil {
				continue
			}
			// line 177: def sin(x):
			πF.SetLineno(177)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp028 = πg.NewFunction(πg.NewCode("sin", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 178: return Sin(float(x))
					πF.SetLineno(178)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßSin); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßsin.ToObject(), πTemp028); πE != nil {
				continue
			}
			// line 181: def tan(x):
			πF.SetLineno(181)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp029 = πg.NewFunction(πg.NewCode("tan", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 182: return Tan(float(x))
					πF.SetLineno(182)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTan); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßtan.ToObject(), πTemp029); πE != nil {
				continue
			}
			// line 187: def degrees(x):
			πF.SetLineno(187)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp030 = πg.NewFunction(πg.NewCode("degrees", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
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
					// line 188: return (float(x) * 180) / pi
					πF.SetLineno(188)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp003[0] = µx
					if πTemp004, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.Mul(πF, πTemp005, πg.NewInt(180).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßpi); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Div(πF, πTemp002, πTemp004); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßdegrees.ToObject(), πTemp030); πE != nil {
				continue
			}
			// line 191: def radians(x):
			πF.SetLineno(191)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp031 = πg.NewFunction(πg.NewCode("radians", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
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
					// line 192: return (float(x) * pi) / 180
					πF.SetLineno(192)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp003[0] = µx
					if πTemp004, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßpi); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mul(πF, πTemp005, πTemp004); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Div(πF, πTemp002, πg.NewInt(180).ToObject()); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßradians.ToObject(), πTemp031); πE != nil {
				continue
			}
			// line 197: def acosh(x):
			πF.SetLineno(197)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp032 = πg.NewFunction(πg.NewCode("acosh", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 198: return Acosh(float(x))
					πF.SetLineno(198)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßAcosh); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßacosh.ToObject(), πTemp032); πE != nil {
				continue
			}
			// line 201: def asinh(x):
			πF.SetLineno(201)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp033 = πg.NewFunction(πg.NewCode("asinh", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 202: return Asinh(float(x))
					πF.SetLineno(202)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßAsinh); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßasinh.ToObject(), πTemp033); πE != nil {
				continue
			}
			// line 205: def atanh(x):
			πF.SetLineno(205)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp034 = πg.NewFunction(πg.NewCode("atanh", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 206: return Atanh(float(x))
					πF.SetLineno(206)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßAtanh); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßatanh.ToObject(), πTemp034); πE != nil {
				continue
			}
			// line 209: def cosh(x):
			πF.SetLineno(209)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp035 = πg.NewFunction(πg.NewCode("cosh", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 210: return Cosh(float(x))
					πF.SetLineno(210)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßCosh); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßcosh.ToObject(), πTemp035); πE != nil {
				continue
			}
			// line 213: def sinh(x):
			πF.SetLineno(213)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp036 = πg.NewFunction(πg.NewCode("sinh", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 214: return Sinh(float(x))
					πF.SetLineno(214)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßSinh); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßsinh.ToObject(), πTemp036); πE != nil {
				continue
			}
			// line 217: def tanh(x):
			πF.SetLineno(217)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp037 = πg.NewFunction(πg.NewCode("tanh", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 218: return Tanh(float(x))
					πF.SetLineno(218)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTanh); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßtanh.ToObject(), πTemp037); πE != nil {
				continue
			}
			// line 223: def erf(x):
			πF.SetLineno(223)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp038 = πg.NewFunction(πg.NewCode("erf", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 224: return Erf(float(x))
					πF.SetLineno(224)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßErf); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßerf.ToObject(), πTemp038); πE != nil {
				continue
			}
			// line 227: def erfc(x):
			πF.SetLineno(227)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp039 = πg.NewFunction(πg.NewCode("erfc", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 228: return Erfc(float(x))
					πF.SetLineno(228)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßErfc); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßerfc.ToObject(), πTemp039); πE != nil {
				continue
			}
			// line 231: def gamma(x):
			πF.SetLineno(231)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp040 = πg.NewFunction(πg.NewCode("gamma", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 232: return Gamma(float(x))
					πF.SetLineno(232)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßGamma); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßgamma.ToObject(), πTemp040); πE != nil {
				continue
			}
			// line 235: def lgamma(x):
			πF.SetLineno(235)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "x", Def: nil}
			πTemp041 = πg.NewFunction(πg.NewCode("lgamma", "build/src/__python__/math.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 236: return Lgamma(float(x))
					πF.SetLineno(236)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßLgamma); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßlgamma.ToObject(), πTemp041); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("math", Code)
}
