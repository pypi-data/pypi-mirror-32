package colorsys
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/colorsys.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßONE_SIXTH := πg.InternStr("ONE_SIXTH")
		ßONE_THIRD := πg.InternStr("ONE_THIRD")
		ßTWO_THIRD := πg.InternStr("TWO_THIRD")
		ß__all__ := πg.InternStr("__all__")
		ß_v := πg.InternStr("_v")
		ßhls_to_rgb := πg.InternStr("hls_to_rgb")
		ßhsv_to_rgb := πg.InternStr("hsv_to_rgb")
		ßint := πg.InternStr("int")
		ßmax := πg.InternStr("max")
		ßmin := πg.InternStr("min")
		ßrgb_to_hls := πg.InternStr("rgb_to_hls")
		ßrgb_to_hsv := πg.InternStr("rgb_to_hsv")
		ßrgb_to_yiq := πg.InternStr("rgb_to_yiq")
		ßyiq_to_rgb := πg.InternStr("yiq_to_rgb")
		var πTemp001 []*πg.Object
		_ = πTemp001
		var πTemp002 *πg.Object
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """Conversion functions between RGB and other color systems.
			πF.SetLineno(1)
			// line 20: __all__ = ["rgb_to_yiq","yiq_to_rgb","rgb_to_hls","hls_to_rgb",
			πF.SetLineno(20)
			πTemp001 = make([]*πg.Object, 6)
			πTemp001[0] = ßrgb_to_yiq.ToObject()
			πTemp001[1] = ßyiq_to_rgb.ToObject()
			πTemp001[2] = ßrgb_to_hls.ToObject()
			πTemp001[3] = ßhls_to_rgb.ToObject()
			πTemp001[4] = ßrgb_to_hsv.ToObject()
			πTemp001[5] = ßhsv_to_rgb.ToObject()
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 25: ONE_THIRD = 1.0/3.0
			πF.SetLineno(25)
			if πTemp002, πE = πg.Div(πF, πg.NewFloat(1.0).ToObject(), πg.NewFloat(3.0).ToObject()); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßONE_THIRD.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 26: ONE_SIXTH = 1.0/6.0
			πF.SetLineno(26)
			if πTemp002, πE = πg.Div(πF, πg.NewFloat(1.0).ToObject(), πg.NewFloat(6.0).ToObject()); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßONE_SIXTH.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 27: TWO_THIRD = 2.0/3.0
			πF.SetLineno(27)
			if πTemp002, πE = πg.Div(πF, πg.NewFloat(2.0).ToObject(), πg.NewFloat(3.0).ToObject()); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTWO_THIRD.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 36: def rgb_to_yiq(r, g, b):
			πF.SetLineno(36)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "r", Def: nil}
			πTemp003[1] = πg.Param{Name: "g", Def: nil}
			πTemp003[2] = πg.Param{Name: "b", Def: nil}
			πTemp002 = πg.NewFunction(πg.NewCode("rgb_to_yiq", "build/src/__python__/colorsys.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µr *πg.Object = πArgs[0]; _ = µr
				var µg *πg.Object = πArgs[1]; _ = µg
				var µb *πg.Object = πArgs[2]; _ = µb
				var µy *πg.Object = πg.UnboundLocal; _ = µy
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µq *πg.Object = πg.UnboundLocal; _ = µq
				var πTemp001 *πg.Object
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
					// line 37: y = 0.30*r + 0.59*g + 0.11*b
					πF.SetLineno(37)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, πg.NewFloat(0.3).ToObject(), µr); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mul(πF, πg.NewFloat(0.59).ToObject(), µg); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, πg.NewFloat(0.11).ToObject(), µb); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp002, πTemp003); πE != nil {
						continue
					}
					µy = πTemp001
					// line 38: i = 0.74*(r-y) - 0.27*(b-y)
					πF.SetLineno(38)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Sub(πF, µr, µy); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mul(πF, πg.NewFloat(0.74).ToObject(), πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Sub(πF, µb, µy); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, πg.NewFloat(0.27).ToObject(), πTemp004); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, πTemp002, πTemp003); πE != nil {
						continue
					}
					µi = πTemp001
					// line 39: q = 0.48*(r-y) + 0.41*(b-y)
					πF.SetLineno(39)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Sub(πF, µr, µy); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mul(πF, πg.NewFloat(0.48).ToObject(), πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Sub(πF, µb, µy); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, πg.NewFloat(0.41).ToObject(), πTemp004); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp002, πTemp003); πE != nil {
						continue
					}
					µq = πTemp001
					// line 40: return (y, i, q)
					πF.SetLineno(40)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple3(µy, µi, µq).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßrgb_to_yiq.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 42: def yiq_to_rgb(y, i, q):
			πF.SetLineno(42)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "y", Def: nil}
			πTemp003[1] = πg.Param{Name: "i", Def: nil}
			πTemp003[2] = πg.Param{Name: "q", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("yiq_to_rgb", "build/src/__python__/colorsys.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µy *πg.Object = πArgs[0]; _ = µy
				var µi *πg.Object = πArgs[1]; _ = µi
				var µq *πg.Object = πArgs[2]; _ = µq
				var µr *πg.Object = πg.UnboundLocal; _ = µr
				var µg *πg.Object = πg.UnboundLocal; _ = µg
				var µb *πg.Object = πg.UnboundLocal; _ = µb
				var πTemp001 *πg.Object
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
					// line 47: r = y + 0.9468822170900693*i + 0.6235565819861433*q
					πF.SetLineno(47)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, πg.NewFloat(0.94688221709).ToObject(), µi); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µy, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, πg.NewFloat(0.623556581986).ToObject(), µq); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp002, πTemp003); πE != nil {
						continue
					}
					µr = πTemp001
					// line 48: g = y - 0.27478764629897834*i - 0.6356910791873801*q
					πF.SetLineno(48)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, πg.NewFloat(0.274787646299).ToObject(), µi); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, µy, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, πg.NewFloat(0.635691079187).ToObject(), µq); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, πTemp002, πTemp003); πE != nil {
						continue
					}
					µg = πTemp001
					// line 49: b = y - 1.1085450346420322*i + 1.7090069284064666*q
					πF.SetLineno(49)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, πg.NewFloat(1.10854503464).ToObject(), µi); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, µy, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, πg.NewFloat(1.70900692841).ToObject(), µq); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp002, πTemp003); πE != nil {
						continue
					}
					µb = πTemp001
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µr, πg.NewFloat(0.0).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 51: if r < 0.0:
					πF.SetLineno(51)
				Label1:
					// line 52: r = 0.0
					πF.SetLineno(52)
					µr = πg.NewFloat(0.0).ToObject()
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µg, πg.NewFloat(0.0).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 53: if g < 0.0:
					πF.SetLineno(53)
				Label3:
					// line 54: g = 0.0
					πF.SetLineno(54)
					µg = πg.NewFloat(0.0).ToObject()
					goto Label4
				Label4:
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µb, πg.NewFloat(0.0).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label5
					}
					goto Label6
					// line 55: if b < 0.0:
					πF.SetLineno(55)
				Label5:
					// line 56: b = 0.0
					πF.SetLineno(56)
					µb = πg.NewFloat(0.0).ToObject()
					goto Label6
				Label6:
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GT(πF, µr, πg.NewFloat(1.0).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label7
					}
					goto Label8
					// line 57: if r > 1.0:
					πF.SetLineno(57)
				Label7:
					// line 58: r = 1.0
					πF.SetLineno(58)
					µr = πg.NewFloat(1.0).ToObject()
					goto Label8
				Label8:
					if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GT(πF, µg, πg.NewFloat(1.0).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label9
					}
					goto Label10
					// line 59: if g > 1.0:
					πF.SetLineno(59)
				Label9:
					// line 60: g = 1.0
					πF.SetLineno(60)
					µg = πg.NewFloat(1.0).ToObject()
					goto Label10
				Label10:
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GT(πF, µb, πg.NewFloat(1.0).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label11
					}
					goto Label12
					// line 61: if b > 1.0:
					πF.SetLineno(61)
				Label11:
					// line 62: b = 1.0
					πF.SetLineno(62)
					µb = πg.NewFloat(1.0).ToObject()
					goto Label12
				Label12:
					// line 63: return (r, g, b)
					πF.SetLineno(63)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple3(µr, µg, µb).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßyiq_to_rgb.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 71: def rgb_to_hls(r, g, b):
			πF.SetLineno(71)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "r", Def: nil}
			πTemp003[1] = πg.Param{Name: "g", Def: nil}
			πTemp003[2] = πg.Param{Name: "b", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("rgb_to_hls", "build/src/__python__/colorsys.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µr *πg.Object = πArgs[0]; _ = µr
				var µg *πg.Object = πArgs[1]; _ = µg
				var µb *πg.Object = πArgs[2]; _ = µb
				var µmaxc *πg.Object = πg.UnboundLocal; _ = µmaxc
				var µminc *πg.Object = πg.UnboundLocal; _ = µminc
				var µl *πg.Object = πg.UnboundLocal; _ = µl
				var µs *πg.Object = πg.UnboundLocal; _ = µs
				var µrc *πg.Object = πg.UnboundLocal; _ = µrc
				var µgc *πg.Object = πg.UnboundLocal; _ = µgc
				var µbc *πg.Object = πg.UnboundLocal; _ = µbc
				var µh *πg.Object = πg.UnboundLocal; _ = µh
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
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
					// line 72: maxc = max(r, g, b)
					πF.SetLineno(72)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp001[0] = µr
					if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
						continue
					}
					πTemp001[1] = µg
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					πTemp001[2] = µb
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µmaxc = πTemp003
					// line 73: minc = min(r, g, b)
					πF.SetLineno(73)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp001[0] = µr
					if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
						continue
					}
					πTemp001[1] = µg
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					πTemp001[2] = µb
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µminc = πTemp003
					// line 75: l = (minc+maxc)/2.0
					πF.SetLineno(75)
					if πE = πg.CheckLocal(πF, µminc, "minc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µminc, µmaxc); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Div(πF, πTemp003, πg.NewFloat(2.0).ToObject()); πE != nil {
						continue
					}
					µl = πTemp002
					if πE = πg.CheckLocal(πF, µminc, "minc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µminc, µmaxc); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 76: if minc == maxc:
					πF.SetLineno(76)
				Label1:
					// line 77: return 0.0, l, 0.0
					πF.SetLineno(77)
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(πg.NewFloat(0.0).ToObject(), µl, πg.NewFloat(0.0).ToObject()).ToObject()
					πR = πTemp002
					continue
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LE(πF, µl, πg.NewFloat(0.5).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 78: if l <= 0.5:
					πF.SetLineno(78)
				Label3:
					// line 79: s = (maxc-minc) / (maxc+minc)
					πF.SetLineno(79)
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µminc, "minc"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Sub(πF, µmaxc, µminc); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µminc, "minc"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Add(πF, µmaxc, µminc); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Div(πF, πTemp003, πTemp005); πE != nil {
						continue
					}
					µs = πTemp002
					goto Label5
				Label4:
					// line 81: s = (maxc-minc) / (2.0-maxc-minc)
					πF.SetLineno(81)
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µminc, "minc"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Sub(πF, µmaxc, µminc); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Sub(πF, πg.NewFloat(2.0).ToObject(), µmaxc); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µminc, "minc"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Sub(πF, πTemp006, µminc); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Div(πF, πTemp003, πTemp005); πE != nil {
						continue
					}
					µs = πTemp002
					goto Label5
				Label5:
					// line 82: rc = (maxc-r) / (maxc-minc)
					πF.SetLineno(82)
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Sub(πF, µmaxc, µr); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µminc, "minc"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Sub(πF, µmaxc, µminc); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Div(πF, πTemp003, πTemp005); πE != nil {
						continue
					}
					µrc = πTemp002
					// line 83: gc = (maxc-g) / (maxc-minc)
					πF.SetLineno(83)
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Sub(πF, µmaxc, µg); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µminc, "minc"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Sub(πF, µmaxc, µminc); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Div(πF, πTemp003, πTemp005); πE != nil {
						continue
					}
					µgc = πTemp002
					// line 84: bc = (maxc-b) / (maxc-minc)
					πF.SetLineno(84)
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Sub(πF, µmaxc, µb); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µminc, "minc"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Sub(πF, µmaxc, µminc); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Div(πF, πTemp003, πTemp005); πE != nil {
						continue
					}
					µbc = πTemp002
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µr, µmaxc); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label6
					}
					if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µg, µmaxc); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label7
					}
					goto Label8
					// line 85: if r == maxc:
					πF.SetLineno(85)
				Label6:
					// line 86: h = bc-gc
					πF.SetLineno(86)
					if πE = πg.CheckLocal(πF, µbc, "bc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgc, "gc"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, µbc, µgc); πE != nil {
						continue
					}
					µh = πTemp002
					goto Label9
					// line 87: elif g == maxc:
					πF.SetLineno(87)
				Label7:
					// line 88: h = 2.0+rc-bc
					πF.SetLineno(88)
					if πE = πg.CheckLocal(πF, µrc, "rc"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πg.NewFloat(2.0).ToObject(), µrc); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µbc, "bc"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, πTemp003, µbc); πE != nil {
						continue
					}
					µh = πTemp002
					goto Label9
				Label8:
					// line 90: h = 4.0+gc-rc
					πF.SetLineno(90)
					if πE = πg.CheckLocal(πF, µgc, "gc"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πg.NewFloat(4.0).ToObject(), µgc); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µrc, "rc"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, πTemp003, µrc); πE != nil {
						continue
					}
					µh = πTemp002
					goto Label9
				Label9:
					// line 91: h = (h/6.0) % 1.0
					πF.SetLineno(91)
					if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Div(πF, µh, πg.NewFloat(6.0).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πTemp003, πg.NewFloat(1.0).ToObject()); πE != nil {
						continue
					}
					µh = πTemp002
					// line 92: return h, l, s
					πF.SetLineno(92)
					if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(µh, µl, µs).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßrgb_to_hls.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 94: def hls_to_rgb(h, l, s):
			πF.SetLineno(94)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "h", Def: nil}
			πTemp003[1] = πg.Param{Name: "l", Def: nil}
			πTemp003[2] = πg.Param{Name: "s", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("hls_to_rgb", "build/src/__python__/colorsys.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µh *πg.Object = πArgs[0]; _ = µh
				var µl *πg.Object = πArgs[1]; _ = µl
				var µs *πg.Object = πArgs[2]; _ = µs
				var µm2 *πg.Object = πg.UnboundLocal; _ = µm2
				var µm1 *πg.Object = πg.UnboundLocal; _ = µm1
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
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
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µs, πg.NewFloat(0.0).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 95: if s == 0.0:
					πF.SetLineno(95)
				Label1:
					// line 96: return l, l, l
					πF.SetLineno(96)
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple3(µl, µl, µl).ToObject()
					πR = πTemp001
					continue
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LE(πF, µl, πg.NewFloat(0.5).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 97: if l <= 0.5:
					πF.SetLineno(97)
				Label3:
					// line 98: m2 = l * (1.0+s)
					πF.SetLineno(98)
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πg.NewFloat(1.0).ToObject(), µs); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mul(πF, µl, πTemp003); πE != nil {
						continue
					}
					µm2 = πTemp001
					goto Label5
				Label4:
					// line 100: m2 = l+s-(l*s)
					πF.SetLineno(100)
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µl, µs); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mul(πF, µl, µs); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					µm2 = πTemp001
					goto Label5
				Label5:
					// line 101: m1 = 2.0*l - m2
					πF.SetLineno(101)
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, πg.NewFloat(2.0).ToObject(), µl); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µm2, "m2"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, πTemp003, µm2); πE != nil {
						continue
					}
					µm1 = πTemp001
					// line 102: return (_v(m1, m2, h+ONE_THIRD), _v(m1, m2, h), _v(m1, m2, h-ONE_THIRD))
					πF.SetLineno(102)
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µm1, "m1"); πE != nil {
						continue
					}
					πTemp005[0] = µm1
					if πE = πg.CheckLocal(πF, µm2, "m2"); πE != nil {
						continue
					}
					πTemp005[1] = µm2
					if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßONE_THIRD); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µh, πTemp004); πE != nil {
						continue
					}
					πTemp005[2] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_v); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µm1, "m1"); πE != nil {
						continue
					}
					πTemp005[0] = µm1
					if πE = πg.CheckLocal(πF, µm2, "m2"); πE != nil {
						continue
					}
					πTemp005[1] = µm2
					if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
						continue
					}
					πTemp005[2] = µh
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_v); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µm1, "m1"); πE != nil {
						continue
					}
					πTemp005[0] = µm1
					if πE = πg.CheckLocal(πF, µm2, "m2"); πE != nil {
						continue
					}
					πTemp005[1] = µm2
					if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßONE_THIRD); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Sub(πF, µh, πTemp007); πE != nil {
						continue
					}
					πTemp005[2] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_v); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001 = πg.NewTuple3(πTemp004, πTemp006, πTemp007).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßhls_to_rgb.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 104: def _v(m1, m2, hue):
			πF.SetLineno(104)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "m1", Def: nil}
			πTemp003[1] = πg.Param{Name: "m2", Def: nil}
			πTemp003[2] = πg.Param{Name: "hue", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("_v", "build/src/__python__/colorsys.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µm1 *πg.Object = πArgs[0]; _ = µm1
				var µm2 *πg.Object = πArgs[1]; _ = µm2
				var µhue *πg.Object = πArgs[2]; _ = µhue
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
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
					// line 105: hue = hue % 1.0
					πF.SetLineno(105)
					if πE = πg.CheckLocal(πF, µhue, "hue"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, µhue, πg.NewFloat(1.0).ToObject()); πE != nil {
						continue
					}
					µhue = πTemp001
					if πE = πg.CheckLocal(πF, µhue, "hue"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßONE_SIXTH); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µhue, πTemp002); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 106: if hue < ONE_SIXTH:
					πF.SetLineno(106)
				Label1:
					// line 107: return m1 + (m2-m1)*hue*6.0
					πF.SetLineno(107)
					if πE = πg.CheckLocal(πF, µm1, "m1"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µm2, "m2"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µm1, "m1"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Sub(πF, µm2, µm1); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µhue, "hue"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mul(πF, πTemp005, µhue); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mul(πF, πTemp004, πg.NewFloat(6.0).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µm1, πTemp002); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µhue, "hue"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µhue, πg.NewFloat(0.5).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label3
					}
					goto Label4
					// line 108: if hue < 0.5:
					πF.SetLineno(108)
				Label3:
					// line 109: return m2
					πF.SetLineno(109)
					if πE = πg.CheckLocal(πF, µm2, "m2"); πE != nil {
						continue
					}
					πR = µm2
					continue
					goto Label4
				Label4:
					if πE = πg.CheckLocal(πF, µhue, "hue"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTWO_THIRD); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µhue, πTemp002); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label5
					}
					goto Label6
					// line 110: if hue < TWO_THIRD:
					πF.SetLineno(110)
				Label5:
					// line 111: return m1 + (m2-m1)*(TWO_THIRD-hue)*6.0
					πF.SetLineno(111)
					if πE = πg.CheckLocal(πF, µm1, "m1"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µm2, "m2"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µm1, "m1"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Sub(πF, µm2, µm1); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßTWO_THIRD); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µhue, "hue"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Sub(πF, πTemp007, µhue); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mul(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mul(πF, πTemp004, πg.NewFloat(6.0).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µm1, πTemp002); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label6
				Label6:
					// line 112: return m1
					πF.SetLineno(112)
					if πE = πg.CheckLocal(πF, µm1, "m1"); πE != nil {
						continue
					}
					πR = µm1
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_v.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 120: def rgb_to_hsv(r, g, b):
			πF.SetLineno(120)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "r", Def: nil}
			πTemp003[1] = πg.Param{Name: "g", Def: nil}
			πTemp003[2] = πg.Param{Name: "b", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("rgb_to_hsv", "build/src/__python__/colorsys.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µr *πg.Object = πArgs[0]; _ = µr
				var µg *πg.Object = πArgs[1]; _ = µg
				var µb *πg.Object = πArgs[2]; _ = µb
				var µmaxc *πg.Object = πg.UnboundLocal; _ = µmaxc
				var µminc *πg.Object = πg.UnboundLocal; _ = µminc
				var µv *πg.Object = πg.UnboundLocal; _ = µv
				var µs *πg.Object = πg.UnboundLocal; _ = µs
				var µrc *πg.Object = πg.UnboundLocal; _ = µrc
				var µgc *πg.Object = πg.UnboundLocal; _ = µgc
				var µbc *πg.Object = πg.UnboundLocal; _ = µbc
				var µh *πg.Object = πg.UnboundLocal; _ = µh
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
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
					// line 121: maxc = max(r, g, b)
					πF.SetLineno(121)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp001[0] = µr
					if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
						continue
					}
					πTemp001[1] = µg
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					πTemp001[2] = µb
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µmaxc = πTemp003
					// line 122: minc = min(r, g, b)
					πF.SetLineno(122)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp001[0] = µr
					if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
						continue
					}
					πTemp001[1] = µg
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					πTemp001[2] = µb
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µminc = πTemp003
					// line 123: v = maxc
					πF.SetLineno(123)
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					µv = µmaxc
					if πE = πg.CheckLocal(πF, µminc, "minc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µminc, µmaxc); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 124: if minc == maxc:
					πF.SetLineno(124)
				Label1:
					// line 125: return 0.0, 0.0, v
					πF.SetLineno(125)
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(πg.NewFloat(0.0).ToObject(), πg.NewFloat(0.0).ToObject(), µv).ToObject()
					πR = πTemp002
					continue
					goto Label2
				Label2:
					// line 126: s = (maxc-minc) / maxc
					πF.SetLineno(126)
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µminc, "minc"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Sub(πF, µmaxc, µminc); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Div(πF, πTemp003, µmaxc); πE != nil {
						continue
					}
					µs = πTemp002
					// line 127: rc = (maxc-r) / (maxc-minc)
					πF.SetLineno(127)
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Sub(πF, µmaxc, µr); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µminc, "minc"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Sub(πF, µmaxc, µminc); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Div(πF, πTemp003, πTemp005); πE != nil {
						continue
					}
					µrc = πTemp002
					// line 128: gc = (maxc-g) / (maxc-minc)
					πF.SetLineno(128)
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Sub(πF, µmaxc, µg); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µminc, "minc"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Sub(πF, µmaxc, µminc); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Div(πF, πTemp003, πTemp005); πE != nil {
						continue
					}
					µgc = πTemp002
					// line 129: bc = (maxc-b) / (maxc-minc)
					πF.SetLineno(129)
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Sub(πF, µmaxc, µb); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µminc, "minc"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Sub(πF, µmaxc, µminc); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Div(πF, πTemp003, πTemp005); πE != nil {
						continue
					}
					µbc = πTemp002
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µr, µmaxc); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmaxc, "maxc"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µg, µmaxc); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					goto Label5
					// line 130: if r == maxc:
					πF.SetLineno(130)
				Label3:
					// line 131: h = bc-gc
					πF.SetLineno(131)
					if πE = πg.CheckLocal(πF, µbc, "bc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgc, "gc"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, µbc, µgc); πE != nil {
						continue
					}
					µh = πTemp002
					goto Label6
					// line 132: elif g == maxc:
					πF.SetLineno(132)
				Label4:
					// line 133: h = 2.0+rc-bc
					πF.SetLineno(133)
					if πE = πg.CheckLocal(πF, µrc, "rc"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πg.NewFloat(2.0).ToObject(), µrc); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µbc, "bc"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, πTemp003, µbc); πE != nil {
						continue
					}
					µh = πTemp002
					goto Label6
				Label5:
					// line 135: h = 4.0+gc-rc
					πF.SetLineno(135)
					if πE = πg.CheckLocal(πF, µgc, "gc"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πg.NewFloat(4.0).ToObject(), µgc); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µrc, "rc"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, πTemp003, µrc); πE != nil {
						continue
					}
					µh = πTemp002
					goto Label6
				Label6:
					// line 136: h = (h/6.0) % 1.0
					πF.SetLineno(136)
					if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Div(πF, µh, πg.NewFloat(6.0).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πTemp003, πg.NewFloat(1.0).ToObject()); πE != nil {
						continue
					}
					µh = πTemp002
					// line 137: return h, s, v
					πF.SetLineno(137)
					if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(µh, µs, µv).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßrgb_to_hsv.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 139: def hsv_to_rgb(h, s, v):
			πF.SetLineno(139)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "h", Def: nil}
			πTemp003[1] = πg.Param{Name: "s", Def: nil}
			πTemp003[2] = πg.Param{Name: "v", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("hsv_to_rgb", "build/src/__python__/colorsys.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µh *πg.Object = πArgs[0]; _ = µh
				var µs *πg.Object = πArgs[1]; _ = µs
				var µv *πg.Object = πArgs[2]; _ = µv
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µf *πg.Object = πg.UnboundLocal; _ = µf
				var µp *πg.Object = πg.UnboundLocal; _ = µp
				var µq *πg.Object = πg.UnboundLocal; _ = µq
				var µt *πg.Object = πg.UnboundLocal; _ = µt
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
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
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µs, πg.NewFloat(0.0).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 140: if s == 0.0:
					πF.SetLineno(140)
				Label1:
					// line 141: return v, v, v
					πF.SetLineno(141)
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple3(µv, µv, µv).ToObject()
					πR = πTemp001
					continue
					goto Label2
				Label2:
					// line 142: i = int(h*6.0) # XXX assume int() truncates!
					πF.SetLineno(142)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mul(πF, µh, πg.NewFloat(6.0).ToObject()); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µi = πTemp004
					// line 143: f = (h*6.0) - i
					πF.SetLineno(143)
					if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mul(πF, µh, πg.NewFloat(6.0).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, πTemp004, µi); πE != nil {
						continue
					}
					µf = πTemp001
					// line 144: p = v*(1.0 - s)
					πF.SetLineno(144)
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Sub(πF, πg.NewFloat(1.0).ToObject(), µs); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mul(πF, µv, πTemp004); πE != nil {
						continue
					}
					µp = πTemp001
					// line 145: q = v*(1.0 - s*f)
					πF.SetLineno(145)
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Mul(πF, µs, µf); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Sub(πF, πg.NewFloat(1.0).ToObject(), πTemp005); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mul(πF, µv, πTemp004); πE != nil {
						continue
					}
					µq = πTemp001
					// line 146: t = v*(1.0 - s*(1.0-f))
					πF.SetLineno(146)
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Sub(πF, πg.NewFloat(1.0).ToObject(), µf); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Mul(πF, µs, πTemp006); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Sub(πF, πg.NewFloat(1.0).ToObject(), πTemp005); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mul(πF, µv, πTemp004); πE != nil {
						continue
					}
					µt = πTemp001
					// line 147: i = i%6
					πF.SetLineno(147)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, µi, πg.NewInt(6).ToObject()); πE != nil {
						continue
					}
					µi = πTemp001
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µi, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 148: if i == 0:
					πF.SetLineno(148)
				Label3:
					// line 149: return v, t, p
					πF.SetLineno(149)
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple3(µv, µt, µp).ToObject()
					πR = πTemp001
					continue
					goto Label4
				Label4:
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label5
					}
					goto Label6
					// line 150: if i == 1:
					πF.SetLineno(150)
				Label5:
					// line 151: return q, v, p
					πF.SetLineno(151)
					if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple3(µq, µv, µp).ToObject()
					πR = πTemp001
					continue
					goto Label6
				Label6:
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µi, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label7
					}
					goto Label8
					// line 152: if i == 2:
					πF.SetLineno(152)
				Label7:
					// line 153: return p, v, t
					πF.SetLineno(153)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple3(µp, µv, µt).ToObject()
					πR = πTemp001
					continue
					goto Label8
				Label8:
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µi, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label9
					}
					goto Label10
					// line 154: if i == 3:
					πF.SetLineno(154)
				Label9:
					// line 155: return p, q, v
					πF.SetLineno(155)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple3(µp, µq, µv).ToObject()
					πR = πTemp001
					continue
					goto Label10
				Label10:
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µi, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label11
					}
					goto Label12
					// line 156: if i == 4:
					πF.SetLineno(156)
				Label11:
					// line 157: return t, p, v
					πF.SetLineno(157)
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple3(µt, µp, µv).ToObject()
					πR = πTemp001
					continue
					goto Label12
				Label12:
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µi, πg.NewInt(5).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label13
					}
					goto Label14
					// line 158: if i == 5:
					πF.SetLineno(158)
				Label13:
					// line 159: return v, p, q
					πF.SetLineno(159)
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple3(µv, µp, µq).ToObject()
					πR = πTemp001
					continue
					goto Label14
				Label14:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßhsv_to_rgb.ToObject(), πTemp009); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("colorsys", Code)
}
