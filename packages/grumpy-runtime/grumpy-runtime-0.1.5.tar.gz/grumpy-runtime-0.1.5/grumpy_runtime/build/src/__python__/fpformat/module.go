package fpformat
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/fpformat.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ß0 := πg.InternStr("0")
		ß5 := πg.InternStr("5")
		ß9 := πg.InternStr("9")
		ßEOFError := πg.InternStr("EOFError")
		ßKeyboardInterrupt := πg.InternStr("KeyboardInterrupt")
		ßNone := πg.InternStr("None")
		ßNotANumber := πg.InternStr("NotANumber")
		ßTypeError := πg.InternStr("TypeError")
		ßValueError := πg.InternStr("ValueError")
		ß__all__ := πg.InternStr("__all__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßabs := πg.InternStr("abs")
		ßchr := πg.InternStr("chr")
		ßcompile := πg.InternStr("compile")
		ßdecoder := πg.InternStr("decoder")
		ße := πg.InternStr("e")
		ßextract := πg.InternStr("extract")
		ßfix := πg.InternStr("fix")
		ßgroup := πg.InternStr("group")
		ßinput := πg.InternStr("input")
		ßint := πg.InternStr("int")
		ßlen := πg.InternStr("len")
		ßmatch := πg.InternStr("match")
		ßmax := πg.InternStr("max")
		ßord := πg.InternStr("ord")
		ßre := πg.InternStr("re")
		ßrepr := πg.InternStr("repr")
		ßroundfrac := πg.InternStr("roundfrac")
		ßsci := πg.InternStr("sci")
		ßtest := πg.InternStr("test")
		ßtype := πg.InternStr("type")
		ßunexpo := πg.InternStr("unexpo")
		ßwarnpy3k := πg.InternStr("warnpy3k")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 πg.KWArgs
		_ = πTemp004
		var πTemp005 *πg.Dict
		_ = πTemp005
		var πTemp006 *πg.Object
		_ = πTemp006
		var πTemp007 *πg.BaseException
		_ = πTemp007
		var πTemp008 *πg.Traceback
		_ = πTemp008
		var πTemp009 bool
		_ = πTemp009
		var πTemp010 []πg.Param
		_ = πTemp010
		var πTemp011 *πg.Object
		_ = πTemp011
		var πTemp012 *πg.Object
		_ = πTemp012
		var πTemp013 *πg.Object
		_ = πTemp013
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 2: goto Label2
			default: panic("unexpected function state")
			}
			// line 1: """General floating point formatting functions.
			πF.SetLineno(1)
			// line 13: from warnings import warnpy3k
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "warnings"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßwarnpy3k, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßwarnpy3k.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 14: warnpy3k("the fpformat module has been removed in Python 3.0", stacklevel=2)
			πF.SetLineno(14)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("the fpformat module has been removed in Python 3.0").ToObject()
			πTemp004 = πg.KWArgs{
				{"stacklevel", πg.NewInt(2).ToObject()},
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßwarnpy3k); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, πTemp004); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 15: del warnpy3k
			πF.SetLineno(15)
			if πE = πg.DelVar(πF, πF.Globals(), ßwarnpy3k); πE != nil {
				continue
			}
			// line 17: import re
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 19: __all__ = ["fix","sci","NotANumber"]
			πF.SetLineno(19)
			πTemp002 = make([]*πg.Object, 3)
			πTemp002[0] = ßfix.ToObject()
			πTemp002[1] = ßsci.ToObject()
			πTemp002[2] = ßNotANumber.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 22: decoder = re.compile(r'^([-+]?)0*(\d*)((?:\.\d*)?)(([eE][-+]?\d+)?)$')
			πF.SetLineno(22)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("^([-+]?)0*(\\d*)((?:\\.\\d*)?)(([eE][-+]?\\d+)?)$").ToObject()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßcompile, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßdecoder.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 29: try:
			πF.SetLineno(29)
			πF.PushCheckpoint(2)
			// line 30: class NotANumber(ValueError):
			πF.SetLineno(30)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp006, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp005 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("NotANumber", "build/src/__python__/fpformat.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 31: pass
					πF.SetLineno(31)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp006, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("NotANumber").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNotANumber.ToObject(), πTemp006); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label1
		Label2:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp007, πTemp008 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
				continue
			}
			if πTemp009, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp009 {
				goto Label3
			}
			πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
			continue
			// line 32: except TypeError:
			πF.SetLineno(32)
		Label3:
			// line 33: NotANumber = 'fpformat.NotANumber'
			πF.SetLineno(33)
			if πE = πF.Globals().SetItem(πF, ßNotANumber.ToObject(), πg.NewStr("fpformat.NotANumber").ToObject()); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label1
		Label1:
			// line 35: def extract(s):
			πF.SetLineno(35)
			πTemp010 = make([]πg.Param, 1)
			πTemp010[0] = πg.Param{Name: "s", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("extract", "build/src/__python__/fpformat.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µres *πg.Object = πg.UnboundLocal; _ = µres
				var µsign *πg.Object = πg.UnboundLocal; _ = µsign
				var µintpart *πg.Object = πg.UnboundLocal; _ = µintpart
				var µfraction *πg.Object = πg.UnboundLocal; _ = µfraction
				var µexppart *πg.Object = πg.UnboundLocal; _ = µexppart
				var µexpo *πg.Object = πg.UnboundLocal; _ = µexpo
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
				var πTemp007 *πg.Object
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 36: """Return (sign, intpart, fraction, expo) or raise an exception:
					πF.SetLineno(36)
					// line 41: res = decoder.match(s)
					πF.SetLineno(41)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdecoder); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmatch, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µres = πTemp002
					if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µres == πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 42: if res is None: raise NotANumber, s
					πF.SetLineno(42)
				Label1:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNotANumber); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					// line 42: if res is None: raise NotANumber, s
					πF.SetLineno(42)
					πE = πF.Raise(πTemp002, µs, nil)
					continue
					goto Label2
				Label2:
					// line 43: sign, intpart, fraction, exppart = res.group(1,2,3,4)
					πF.SetLineno(43)
					πTemp001 = πF.MakeArgs(4)
					πTemp001[0] = πg.NewInt(1).ToObject()
					πTemp001[1] = πg.NewInt(2).ToObject()
					πTemp001[2] = πg.NewInt(3).ToObject()
					πTemp001[3] = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µres, ßgroup, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					µsign = πTemp002
					µintpart = πTemp005
					µfraction = πTemp006
					µexppart = πTemp007
					if πE = πg.CheckLocal(πF, µsign, "sign"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µsign, πg.NewStr("+").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 44: if sign == '+': sign = ''
					πF.SetLineno(44)
				Label3:
					// line 44: if sign == '+': sign = ''
					πF.SetLineno(44)
					µsign = ß.ToObject()
					goto Label4
				Label4:
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µfraction); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label5
					}
					goto Label6
					// line 45: if fraction: fraction = fraction[1:]
					πF.SetLineno(45)
				Label5:
					// line 45: if fraction: fraction = fraction[1:]
					πF.SetLineno(45)
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µfraction, πTemp002); πE != nil {
						continue
					}
					µfraction = πTemp003
					goto Label6
				Label6:
					if πE = πg.CheckLocal(πF, µexppart, "exppart"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µexppart); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label7
					}
					goto Label8
					// line 46: if exppart: expo = int(exppart[1:])
					πF.SetLineno(46)
				Label7:
					// line 46: if exppart: expo = int(exppart[1:])
					πF.SetLineno(46)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µexppart, "exppart"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µexppart, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µexpo = πTemp003
					goto Label9
				Label8:
					// line 47: else: expo = 0
					πF.SetLineno(47)
					µexpo = πg.NewInt(0).ToObject()
					goto Label9
				Label9:
					// line 48: return sign, intpart, fraction, expo
					πF.SetLineno(48)
					if πE = πg.CheckLocal(πF, µsign, "sign"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µexpo, "expo"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple4(µsign, µintpart, µfraction, µexpo).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßextract.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 50: def unexpo(intpart, fraction, expo):
			πF.SetLineno(50)
			πTemp010 = make([]πg.Param, 3)
			πTemp010[0] = πg.Param{Name: "intpart", Def: nil}
			πTemp010[1] = πg.Param{Name: "fraction", Def: nil}
			πTemp010[2] = πg.Param{Name: "expo", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("unexpo", "build/src/__python__/fpformat.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µintpart *πg.Object = πArgs[0]; _ = µintpart
				var µfraction *πg.Object = πArgs[1]; _ = µfraction
				var µexpo *πg.Object = πArgs[2]; _ = µexpo
				var µf *πg.Object = πg.UnboundLocal; _ = µf
				var µi *πg.Object = πg.UnboundLocal; _ = µi
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
				var πTemp007 *πg.Object
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 51: """Remove the exponent by changing intpart and fraction."""
					πF.SetLineno(51)
					if πE = πg.CheckLocal(πF, µexpo, "expo"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GT(πF, µexpo, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µexpo, "expo"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µexpo, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 52: if expo > 0: # Move the point left
					πF.SetLineno(52)
				Label1:
					// line 53: f = len(fraction)
					πF.SetLineno(53)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					πTemp003[0] = µfraction
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µf = πTemp004
					// line 54: intpart, fraction = intpart + fraction[:expo], fraction[expo:]
					πF.SetLineno(54)
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µexpo, "expo"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µexpo, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µfraction, πTemp005); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, µintpart, πTemp006); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µexpo, "expo"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{µexpo, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µfraction, πTemp005); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp004, πTemp006).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp001); πE != nil {
						continue
					}
					µintpart = πTemp004
					µfraction = πTemp005
					if πE = πg.CheckLocal(πF, µexpo, "expo"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GT(πF, µexpo, µf); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					goto Label5
					// line 55: if expo > f:
					πF.SetLineno(55)
				Label4:
					// line 56: intpart = intpart + '0'*(expo-f)
					πF.SetLineno(56)
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µexpo, "expo"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Sub(πF, µexpo, µf); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mul(πF, ß0.ToObject(), πTemp005); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µintpart, πTemp004); πE != nil {
						continue
					}
					µintpart = πTemp001
					goto Label5
				Label5:
					goto Label3
					// line 57: elif expo < 0: # Move the point right
					πF.SetLineno(57)
				Label2:
					// line 58: i = len(intpart)
					πF.SetLineno(58)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					πTemp003[0] = µintpart
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µi = πTemp004
					// line 59: intpart, fraction = intpart[:expo], intpart[expo:] + fraction
					πF.SetLineno(59)
					if πE = πg.CheckLocal(πF, µexpo, "expo"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µexpo, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µintpart, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µexpo, "expo"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{µexpo, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µintpart, πTemp006); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, πTemp007, µfraction); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp005, πTemp004).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp001); πE != nil {
						continue
					}
					µintpart = πTemp004
					µfraction = πTemp005
					if πE = πg.CheckLocal(πF, µexpo, "expo"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Neg(πF, µi); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µexpo, πTemp004); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label6
					}
					goto Label7
					// line 60: if expo < -i:
					πF.SetLineno(60)
				Label6:
					// line 61: fraction = '0'*(-expo-i) + fraction
					πF.SetLineno(61)
					if πE = πg.CheckLocal(πF, µexpo, "expo"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Neg(πF, µexpo); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Sub(πF, πTemp006, µi); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mul(πF, ß0.ToObject(), πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp004, µfraction); πE != nil {
						continue
					}
					µfraction = πTemp001
					goto Label7
				Label7:
					goto Label3
				Label3:
					// line 62: return intpart, fraction
					πF.SetLineno(62)
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µintpart, µfraction).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßunexpo.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 64: def roundfrac(intpart, fraction, digs):
			πF.SetLineno(64)
			πTemp010 = make([]πg.Param, 3)
			πTemp010[0] = πg.Param{Name: "intpart", Def: nil}
			πTemp010[1] = πg.Param{Name: "fraction", Def: nil}
			πTemp010[2] = πg.Param{Name: "digs", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("roundfrac", "build/src/__python__/fpformat.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µintpart *πg.Object = πArgs[0]; _ = µintpart
				var µfraction *πg.Object = πArgs[1]; _ = µfraction
				var µdigs *πg.Object = πArgs[2]; _ = µdigs
				var µf *πg.Object = πg.UnboundLocal; _ = µf
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µtotal *πg.Object = πg.UnboundLocal; _ = µtotal
				var µnextdigit *πg.Object = πg.UnboundLocal; _ = µnextdigit
				var µn *πg.Object = πg.UnboundLocal; _ = µn
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
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 8: goto Label8
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					// line 65: """Round or extend the fraction to size digs."""
					πF.SetLineno(65)
					// line 66: f = len(fraction)
					πF.SetLineno(66)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					πTemp001[0] = µfraction
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µf = πTemp003
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdigs, "digs"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LE(πF, µf, µdigs); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 67: if f <= digs:
					πF.SetLineno(67)
				Label1:
					// line 68: return intpart, fraction + '0'*(digs-f)
					πF.SetLineno(68)
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdigs, "digs"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Sub(πF, µdigs, µf); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Mul(πF, ß0.ToObject(), πTemp006); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µfraction, πTemp005); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µintpart, πTemp003).ToObject()
					πR = πTemp002
					continue
					goto Label2
				Label2:
					// line 69: i = len(intpart)
					πF.SetLineno(69)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					πTemp001[0] = µintpart
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µi = πTemp003
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdigs, "digs"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µi, µdigs); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LT(πF, πTemp003, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 70: if i+digs < 0:
					πF.SetLineno(70)
				Label3:
					// line 71: return '0'*-digs, ''
					πF.SetLineno(71)
					if πE = πg.CheckLocal(πF, µdigs, "digs"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Neg(πF, µdigs); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, ß0.ToObject(), πTemp005); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, ß.ToObject()).ToObject()
					πR = πTemp002
					continue
					goto Label4
				Label4:
					// line 72: total = intpart + fraction
					πF.SetLineno(72)
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µintpart, µfraction); πE != nil {
						continue
					}
					µtotal = πTemp002
					// line 73: nextdigit = total[i+digs]
					πF.SetLineno(73)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdigs, "digs"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µi, µdigs); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πE = πg.CheckLocal(πF, µtotal, "total"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µtotal, πTemp002); πE != nil {
						continue
					}
					µnextdigit = πTemp003
					if πE = πg.CheckLocal(πF, µnextdigit, "nextdigit"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GE(πF, µnextdigit, ß5.ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label5
					}
					goto Label6
					// line 74: if nextdigit >= '5': # Hard case: increment last digit, may have carry!
					πF.SetLineno(74)
				Label5:
					// line 75: n = i + digs - 1
					πF.SetLineno(75)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdigs, "digs"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µi, µdigs); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µn = πTemp002
					// line 76: while n >= 0:
					πF.SetLineno(76)
					πF.PushCheckpoint(8)
					πTemp004 = false
				Label7:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label9
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GE(πF, µn, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(7)            
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp003 = µn
					if πE = πg.CheckLocal(πF, µtotal, "total"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µtotal, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.NE(πF, πTemp005, ß9.ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label10
					}
					goto Label11
					// line 77: if total[n] != '9': break
					πF.SetLineno(77)
				Label10:
					// line 77: if total[n] != '9': break
					πF.SetLineno(77)
					πTemp004 = true
					continue
					goto Label11
				Label11:
					// line 78: n = n-1
					πF.SetLineno(78)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µn = πTemp002
					continue
				Label8:
					if πE != nil || πR != nil {
						continue
					}
					// line 80: total = '0' + total
					πF.SetLineno(80)
					if πE = πg.CheckLocal(πF, µtotal, "total"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, ß0.ToObject(), µtotal); πE != nil {
						continue
					}
					µtotal = πTemp002
					// line 81: i = i+1
					πF.SetLineno(81)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µi = πTemp002
					// line 82: n = 0
					πF.SetLineno(82)
					µn = πg.NewInt(0).ToObject()
				Label9:
					// line 83: total = total[:n] + chr(ord(total[n]) + 1) + '0'*(len(total)-n-1)
					πF.SetLineno(83)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µn, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtotal, "total"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µtotal, πTemp005); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp008 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp009 = µn
					if πE = πg.CheckLocal(πF, µtotal, "total"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µtotal, πTemp009); πE != nil {
						continue
					}
					πTemp008[0] = πTemp010
					if πTemp009, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πTemp005, πE = πg.Add(πF, πTemp010, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001[0] = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.Add(πF, πTemp006, πTemp009); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtotal, "total"); πE != nil {
						continue
					}
					πTemp001[0] = µtotal
					if πTemp010, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Sub(πF, πTemp011, µn); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Sub(πF, πTemp009, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Mul(πF, ß0.ToObject(), πTemp006); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp003, πTemp005); πE != nil {
						continue
					}
					µtotal = πTemp002
					// line 84: intpart, fraction = total[:i], total[i:]
					πF.SetLineno(84)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µi, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtotal, "total"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µtotal, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{µi, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtotal, "total"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µtotal, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
						continue
					}
					µintpart = πTemp003
					µfraction = πTemp005
					goto Label6
				Label6:
					if πE = πg.CheckLocal(πF, µdigs, "digs"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GE(πF, µdigs, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label12
					}
					goto Label13
					// line 85: if digs >= 0:
					πF.SetLineno(85)
				Label12:
					// line 86: return intpart, fraction[:digs]
					πF.SetLineno(86)
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdigs, "digs"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µdigs, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µfraction, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µintpart, πTemp005).ToObject()
					πR = πTemp002
					continue
					goto Label14
				Label13:
					// line 88: return intpart[:digs] + '0'*-digs, ''
					πF.SetLineno(88)
					if πE = πg.CheckLocal(πF, µdigs, "digs"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µdigs, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µintpart, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdigs, "digs"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Neg(πF, µdigs); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Mul(πF, ß0.ToObject(), πTemp009); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πTemp006, πTemp005); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, ß.ToObject()).ToObject()
					πR = πTemp002
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
			if πE = πF.Globals().SetItem(πF, ßroundfrac.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 90: def fix(x, digs):
			πF.SetLineno(90)
			πTemp010 = make([]πg.Param, 2)
			πTemp010[0] = πg.Param{Name: "x", Def: nil}
			πTemp010[1] = πg.Param{Name: "digs", Def: nil}
			πTemp011 = πg.NewFunction(πg.NewCode("fix", "build/src/__python__/fpformat.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µdigs *πg.Object = πArgs[1]; _ = µdigs
				var µsign *πg.Object = πg.UnboundLocal; _ = µsign
				var µintpart *πg.Object = πg.UnboundLocal; _ = µintpart
				var µfraction *πg.Object = πg.UnboundLocal; _ = µfraction
				var µexpo *πg.Object = πg.UnboundLocal; _ = µexpo
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
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.BaseException
				_ = πTemp008
				var πTemp009 *πg.Traceback
				_ = πTemp009
				var πTemp010 bool
				_ = πTemp010
				var πTemp011 bool
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 4: goto Label4
					case 6: goto Label6
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					// line 91: """Format x as [-]ddd.ddd with 'digs' digits after the point
					πF.SetLineno(91)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = ß.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.NE(πF, πTemp004, πTemp005); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label1
					}
					goto Label2
					// line 94: if type(x) != type(''): x = repr(x)
					πF.SetLineno(94)
				Label1:
					// line 94: if type(x) != type(''): x = repr(x)
					πF.SetLineno(94)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp001, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µx = πTemp003
					goto Label2
				Label2:
					// line 95: try:
					πF.SetLineno(95)
					πF.PushCheckpoint(4)
					// line 96: sign, intpart, fraction, expo = extract(x)
					πF.SetLineno(96)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp001, πE = πg.ResolveGlobal(πF, ßextract); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					µsign = πTemp001
					µintpart = πTemp004
					µfraction = πTemp005
					µexpo = πTemp007
					πF.PopCheckpoint()
					goto Label3
				Label4:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNotANumber); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label5
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 97: except NotANumber:
					πF.SetLineno(97)
				Label5:
					// line 98: return x
					πF.SetLineno(98)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πR = µx
					continue
					πF.RestoreExc(nil, nil)
					goto Label3
				Label3:
					// line 99: intpart, fraction = unexpo(intpart, fraction, expo)
					πF.SetLineno(99)
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					πTemp002[0] = µintpart
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					πTemp002[1] = µfraction
					if πE = πg.CheckLocal(πF, µexpo, "expo"); πE != nil {
						continue
					}
					πTemp002[2] = µexpo
					if πTemp001, πE = πg.ResolveGlobal(πF, ßunexpo); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µintpart = πTemp001
					µfraction = πTemp004
					// line 100: intpart, fraction = roundfrac(intpart, fraction, digs)
					πF.SetLineno(100)
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					πTemp002[0] = µintpart
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					πTemp002[1] = µfraction
					if πE = πg.CheckLocal(πF, µdigs, "digs"); πE != nil {
						continue
					}
					πTemp002[2] = µdigs
					if πTemp001, πE = πg.ResolveGlobal(πF, ßroundfrac); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µintpart = πTemp001
					µfraction = πTemp004
					// line 101: while intpart and intpart[0] == '0': intpart = intpart[1:]
					πF.SetLineno(101)
					πF.PushCheckpoint(7)
					πTemp006 = false
				Label6:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label8
					}
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					πTemp001 = µintpart
					if πTemp011, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp011 {
						goto Label9
					}
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µintpart, πTemp004); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp005, ß0.ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
				Label9:
					if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp010 {
						continue
					}
					πF.PushCheckpoint(6)            
					// line 101: while intpart and intpart[0] == '0': intpart = intpart[1:]
					πF.SetLineno(101)
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µintpart, πTemp001); πE != nil {
						continue
					}
					µintpart = πTemp003
					continue
				Label7:
					if πE != nil || πR != nil {
						continue
					}
				Label8:
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µintpart, ß.ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label10
					}
					goto Label11
					// line 102: if intpart == '': intpart = '0'
					πF.SetLineno(102)
				Label10:
					// line 102: if intpart == '': intpart = '0'
					πF.SetLineno(102)
					µintpart = ß0.ToObject()
					goto Label11
				Label11:
					if πE = πg.CheckLocal(πF, µdigs, "digs"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GT(πF, µdigs, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label12
					}
					goto Label13
					// line 103: if digs > 0: return sign + intpart + '.' + fraction
					πF.SetLineno(103)
				Label12:
					// line 103: if digs > 0: return sign + intpart + '.' + fraction
					πF.SetLineno(103)
					if πE = πg.CheckLocal(πF, µsign, "sign"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, µsign, µintpart); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πTemp004, πg.NewStr(".").ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp003, µfraction); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label14
				Label13:
					// line 104: else: return sign + intpart
					πF.SetLineno(104)
					if πE = πg.CheckLocal(πF, µsign, "sign"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µsign, µintpart); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßfix.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 106: def sci(x, digs):
			πF.SetLineno(106)
			πTemp010 = make([]πg.Param, 2)
			πTemp010[0] = πg.Param{Name: "x", Def: nil}
			πTemp010[1] = πg.Param{Name: "digs", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("sci", "build/src/__python__/fpformat.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µdigs *πg.Object = πArgs[1]; _ = µdigs
				var µsign *πg.Object = πg.UnboundLocal; _ = µsign
				var µintpart *πg.Object = πg.UnboundLocal; _ = µintpart
				var µfraction *πg.Object = πg.UnboundLocal; _ = µfraction
				var µexpo *πg.Object = πg.UnboundLocal; _ = µexpo
				var µs *πg.Object = πg.UnboundLocal; _ = µs
				var µe *πg.Object = πg.UnboundLocal; _ = µe
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
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 []*πg.Object
				_ = πTemp012
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 6: goto Label6
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					// line 107: """Format x as [-]d.dddE[+-]ddd with 'digs' digits after the point
					πF.SetLineno(107)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = ß.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.NE(πF, πTemp004, πTemp005); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label1
					}
					goto Label2
					// line 110: if type(x) != type(''): x = repr(x)
					πF.SetLineno(110)
				Label1:
					// line 110: if type(x) != type(''): x = repr(x)
					πF.SetLineno(110)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp001, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µx = πTemp003
					goto Label2
				Label2:
					// line 111: sign, intpart, fraction, expo = extract(x)
					πF.SetLineno(111)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp001, πE = πg.ResolveGlobal(πF, ßextract); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					µsign = πTemp001
					µintpart = πTemp004
					µfraction = πTemp005
					µexpo = πTemp007
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µintpart); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					goto Label4
					// line 112: if not intpart:
					πF.SetLineno(112)
				Label3:
					// line 113: while fraction and fraction[0] == '0':
					πF.SetLineno(113)
					πF.PushCheckpoint(7)
					πTemp006 = false
				Label6:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label8
					}
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					πTemp001 = µfraction
					if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp009 {
						goto Label9
					}
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µfraction, πTemp004); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp005, ß0.ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
				Label9:
					if πTemp008, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(6)            
					// line 114: fraction = fraction[1:]
					πF.SetLineno(114)
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µfraction, πTemp001); πE != nil {
						continue
					}
					µfraction = πTemp003
					// line 115: expo = expo - 1
					πF.SetLineno(115)
					if πE = πg.CheckLocal(πF, µexpo, "expo"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, µexpo, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µexpo = πTemp001
					continue
				Label7:
					if πE != nil || πR != nil {
						continue
					}
				Label8:
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µfraction); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label10
					}
					goto Label11
					// line 116: if fraction:
					πF.SetLineno(116)
				Label10:
					// line 117: intpart, fraction = fraction[0], fraction[1:]
					πF.SetLineno(117)
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µfraction, πTemp003); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µfraction, πTemp003); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
						continue
					}
					µintpart = πTemp003
					µfraction = πTemp004
					// line 118: expo = expo - 1
					πF.SetLineno(118)
					if πE = πg.CheckLocal(πF, µexpo, "expo"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, µexpo, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µexpo = πTemp001
					goto Label12
				Label11:
					// line 120: intpart = '0'
					πF.SetLineno(120)
					µintpart = ß0.ToObject()
					goto Label12
				Label12:
					goto Label5
				Label4:
					// line 122: expo = expo + len(intpart) - 1
					πF.SetLineno(122)
					if πE = πg.CheckLocal(πF, µexpo, "expo"); πE != nil {
						continue
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					πTemp002[0] = µintpart
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.Add(πF, µexpo, πTemp005); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µexpo = πTemp001
					// line 123: intpart, fraction = intpart[0], intpart[1:] + fraction
					πF.SetLineno(123)
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µintpart, πTemp003); πE != nil {
						continue
					}
					if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µintpart, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πTemp007, µfraction); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp004, πTemp003).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
						continue
					}
					µintpart = πTemp003
					µfraction = πTemp004
					goto Label5
				Label5:
					// line 124: digs = max(0, digs)
					πF.SetLineno(124)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µdigs, "digs"); πE != nil {
						continue
					}
					πTemp002[1] = µdigs
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µdigs = πTemp003
					// line 125: intpart, fraction = roundfrac(intpart, fraction, digs)
					πF.SetLineno(125)
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					πTemp002[0] = µintpart
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					πTemp002[1] = µfraction
					if πE = πg.CheckLocal(πF, µdigs, "digs"); πE != nil {
						continue
					}
					πTemp002[2] = µdigs
					if πTemp001, πE = πg.ResolveGlobal(πF, ßroundfrac); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µintpart = πTemp001
					µfraction = πTemp004
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					πTemp002[0] = µintpart
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.GT(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label13
					}
					goto Label14
					// line 126: if len(intpart) > 1:
					πF.SetLineno(126)
				Label13:
					// line 127: intpart, fraction, expo = \
					πF.SetLineno(127)
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µintpart, πTemp003); πE != nil {
						continue
					}
					if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µintpart, πTemp005); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp010, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µfraction, πTemp005); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πTemp007, πTemp010); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µexpo, "expo"); πE != nil {
						continue
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					πTemp002[0] = µintpart
					if πTemp010, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp007, πE = πg.Add(πF, µexpo, πTemp011); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Sub(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple3(πTemp004, πTemp003, πTemp005).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp001); πE != nil {
						continue
					}
					µintpart = πTemp003
					µfraction = πTemp004
					µexpo = πTemp005
					goto Label14
				Label14:
					// line 130: s = sign + intpart
					πF.SetLineno(130)
					if πE = πg.CheckLocal(πF, µsign, "sign"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µintpart, "intpart"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µsign, µintpart); πE != nil {
						continue
					}
					µs = πTemp001
					if πE = πg.CheckLocal(πF, µdigs, "digs"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GT(πF, µdigs, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label15
					}
					goto Label16
					// line 131: if digs > 0: s = s + '.' + fraction
					πF.SetLineno(131)
				Label15:
					// line 131: if digs > 0: s = s + '.' + fraction
					πF.SetLineno(131)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µs, πg.NewStr(".").ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfraction, "fraction"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp003, µfraction); πE != nil {
						continue
					}
					µs = πTemp001
					goto Label16
				Label16:
					// line 132: e = repr(abs(expo))
					πF.SetLineno(132)
					πTemp002 = πF.MakeArgs(1)
					πTemp012 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µexpo, "expo"); πE != nil {
						continue
					}
					πTemp012[0] = µexpo
					if πTemp001, πE = πg.ResolveGlobal(πF, ßabs); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp002[0] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µe = πTemp003
					// line 133: e = '0'*(3-len(e)) + e
					πF.SetLineno(133)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
						continue
					}
					πTemp002[0] = µe
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp004, πE = πg.Sub(πF, πg.NewInt(3).ToObject(), πTemp007); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, ß0.ToObject(), πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp003, µe); πE != nil {
						continue
					}
					µe = πTemp001
					if πE = πg.CheckLocal(πF, µexpo, "expo"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µexpo, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label17
					}
					goto Label18
					// line 134: if expo < 0: e = '-' + e
					πF.SetLineno(134)
				Label17:
					// line 134: if expo < 0: e = '-' + e
					πF.SetLineno(134)
					if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πg.NewStr("-").ToObject(), µe); πE != nil {
						continue
					}
					µe = πTemp001
					goto Label19
				Label18:
					// line 135: else: e = '+' + e
					πF.SetLineno(135)
					if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πg.NewStr("+").ToObject(), µe); πE != nil {
						continue
					}
					µe = πTemp001
					goto Label19
				Label19:
					// line 136: return s + 'e' + e
					πF.SetLineno(136)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µs, ße.ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp003, µe); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßsci.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 138: def test():
			πF.SetLineno(138)
			πTemp010 = make([]πg.Param, 0)
			πTemp013 = πg.NewFunction(πg.NewCode("test", "build/src/__python__/fpformat.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πg.UnboundLocal; _ = µx
				var µdigs *πg.Object = πg.UnboundLocal; _ = µdigs
				var πTemp001 bool
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
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 *πg.BaseException
				_ = πTemp008
				var πTemp009 *πg.Traceback
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					case 3: goto Label3
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					// line 139: """Interactive test run."""
					πF.SetLineno(139)
					// line 140: try:
					πF.SetLineno(140)
					πF.PushCheckpoint(2)
					// line 141: while 1:
					πF.SetLineno(141)
					πF.PushCheckpoint(4)
					πTemp001 = false
				Label3:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp001 {
						πF.PopCheckpoint()
						goto Label5
					}
					if πTemp002, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE != nil || !πTemp002 {
						continue
					}
					πF.PushCheckpoint(3)            
					// line 142: x, digs = input('Enter (x, digs): ')
					πF.SetLineno(142)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("Enter (x, digs): ").ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßinput); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp006}}}, πTemp005); πE != nil {
						continue
					}
					µx = πTemp004
					µdigs = πTemp006
					// line 143: print x, fix(x, digs), sci(x, digs)
					πF.SetLineno(143)
					πTemp003 = make([]*πg.Object, 3)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp003[0] = µx
					πTemp007 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp007[0] = µx
					if πE = πg.CheckLocal(πF, µdigs, "digs"); πE != nil {
						continue
					}
					πTemp007[1] = µdigs
					if πTemp004, πE = πg.ResolveGlobal(πF, ßfix); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp003[1] = πTemp005
					πTemp007 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp007[0] = µx
					if πE = πg.CheckLocal(πF, µdigs, "digs"); πE != nil {
						continue
					}
					πTemp007[1] = µdigs
					if πTemp004, πE = πg.ResolveGlobal(πF, ßsci); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp003[2] = πTemp005
					if πE = πg.Print(πF, πTemp003, true); πE != nil {
						continue
					}
					continue
				Label4:
					if πE != nil || πR != nil {
						continue
					}
				Label5:
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp005, πE = πg.ResolveGlobal(πF, ßEOFError); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßKeyboardInterrupt); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
					if πTemp001, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp004); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label6
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 144: except (EOFError, KeyboardInterrupt):
					πF.SetLineno(144)
				Label6:
					// line 145: pass
					πF.SetLineno(145)
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
			if πE = πF.Globals().SetItem(πF, ßtest.ToObject(), πTemp013); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("fpformat", Code)
}
