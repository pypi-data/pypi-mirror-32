package operator
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/operator.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßAbs := πg.InternStr("Abs")
		ßAttributeError := πg.InternStr("AttributeError")
		ßFalse := πg.InternStr("False")
		ßNotImplemented := πg.InternStr("NotImplemented")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßValueError := πg.InternStr("ValueError")
		ß__abs__ := πg.InternStr("__abs__")
		ß__add__ := πg.InternStr("__add__")
		ß__all__ := πg.InternStr("__all__")
		ß__and__ := πg.InternStr("__and__")
		ß__call__ := πg.InternStr("__call__")
		ß__concat__ := πg.InternStr("__concat__")
		ß__contains__ := πg.InternStr("__contains__")
		ß__delitem__ := πg.InternStr("__delitem__")
		ß__eq__ := πg.InternStr("__eq__")
		ß__floordiv__ := πg.InternStr("__floordiv__")
		ß__ge__ := πg.InternStr("__ge__")
		ß__getitem__ := πg.InternStr("__getitem__")
		ß__gt__ := πg.InternStr("__gt__")
		ß__iadd__ := πg.InternStr("__iadd__")
		ß__iand__ := πg.InternStr("__iand__")
		ß__iconcat__ := πg.InternStr("__iconcat__")
		ß__ifloordiv__ := πg.InternStr("__ifloordiv__")
		ß__ilshift__ := πg.InternStr("__ilshift__")
		ß__imod__ := πg.InternStr("__imod__")
		ß__imul__ := πg.InternStr("__imul__")
		ß__index__ := πg.InternStr("__index__")
		ß__init__ := πg.InternStr("__init__")
		ß__inv__ := πg.InternStr("__inv__")
		ß__invert__ := πg.InternStr("__invert__")
		ß__ior__ := πg.InternStr("__ior__")
		ß__ipow__ := πg.InternStr("__ipow__")
		ß__irshift__ := πg.InternStr("__irshift__")
		ß__isub__ := πg.InternStr("__isub__")
		ß__itruediv__ := πg.InternStr("__itruediv__")
		ß__ixor__ := πg.InternStr("__ixor__")
		ß__le__ := πg.InternStr("__le__")
		ß__length_hint__ := πg.InternStr("__length_hint__")
		ß__lshift__ := πg.InternStr("__lshift__")
		ß__lt__ := πg.InternStr("__lt__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__mod__ := πg.InternStr("__mod__")
		ß__module__ := πg.InternStr("__module__")
		ß__mul__ := πg.InternStr("__mul__")
		ß__name__ := πg.InternStr("__name__")
		ß__ne__ := πg.InternStr("__ne__")
		ß__neg__ := πg.InternStr("__neg__")
		ß__not__ := πg.InternStr("__not__")
		ß__or__ := πg.InternStr("__or__")
		ß__pos__ := πg.InternStr("__pos__")
		ß__pow__ := πg.InternStr("__pow__")
		ß__rshift__ := πg.InternStr("__rshift__")
		ß__setitem__ := πg.InternStr("__setitem__")
		ß__sub__ := πg.InternStr("__sub__")
		ß__truediv__ := πg.InternStr("__truediv__")
		ß__xor__ := πg.InternStr("__xor__")
		ß_abs := πg.InternStr("_abs")
		ß_args := πg.InternStr("_args")
		ß_call := πg.InternStr("_call")
		ß_kwargs := πg.InternStr("_kwargs")
		ß_name := πg.InternStr("_name")
		ßabs := πg.InternStr("abs")
		ßadd := πg.InternStr("add")
		ßand_ := πg.InternStr("and_")
		ßattrgetter := πg.InternStr("attrgetter")
		ßconcat := πg.InternStr("concat")
		ßcontains := πg.InternStr("contains")
		ßcountOf := πg.InternStr("countOf")
		ßdelitem := πg.InternStr("delitem")
		ßenumerate := πg.InternStr("enumerate")
		ßeq := πg.InternStr("eq")
		ßfloat := πg.InternStr("float")
		ßfloordiv := πg.InternStr("floordiv")
		ßge := πg.InternStr("ge")
		ßgetattr := πg.InternStr("getattr")
		ßgetitem := πg.InternStr("getitem")
		ßgt := πg.InternStr("gt")
		ßhasattr := πg.InternStr("hasattr")
		ßiadd := πg.InternStr("iadd")
		ßiand := πg.InternStr("iand")
		ßiconcat := πg.InternStr("iconcat")
		ßifloordiv := πg.InternStr("ifloordiv")
		ßilshift := πg.InternStr("ilshift")
		ßimod := πg.InternStr("imod")
		ßimul := πg.InternStr("imul")
		ßindex := πg.InternStr("index")
		ßindexOf := πg.InternStr("indexOf")
		ßint := πg.InternStr("int")
		ßinv := πg.InternStr("inv")
		ßinvert := πg.InternStr("invert")
		ßior := πg.InternStr("ior")
		ßipow := πg.InternStr("ipow")
		ßirshift := πg.InternStr("irshift")
		ßis_ := πg.InternStr("is_")
		ßis_not := πg.InternStr("is_not")
		ßisinstance := πg.InternStr("isinstance")
		ßisub := πg.InternStr("isub")
		ßitemgetter := πg.InternStr("itemgetter")
		ßitruediv := πg.InternStr("itruediv")
		ßixor := πg.InternStr("ixor")
		ßle := πg.InternStr("le")
		ßlen := πg.InternStr("len")
		ßlength_hint := πg.InternStr("length_hint")
		ßlong := πg.InternStr("long")
		ßlshift := πg.InternStr("lshift")
		ßlt := πg.InternStr("lt")
		ßmap := πg.InternStr("map")
		ßmethodcaller := πg.InternStr("methodcaller")
		ßmod := πg.InternStr("mod")
		ßmul := πg.InternStr("mul")
		ßne := πg.InternStr("ne")
		ßneg := πg.InternStr("neg")
		ßnot_ := πg.InternStr("not_")
		ßobject := πg.InternStr("object")
		ßor_ := πg.InternStr("or_")
		ßpos := πg.InternStr("pos")
		ßpow := πg.InternStr("pow")
		ßrshift := πg.InternStr("rshift")
		ßsetitem := πg.InternStr("setitem")
		ßsplit := πg.InternStr("split")
		ßstr := πg.InternStr("str")
		ßsub := πg.InternStr("sub")
		ßtruediv := πg.InternStr("truediv")
		ßtruth := πg.InternStr("truth")
		ßtuple := πg.InternStr("tuple")
		ßtype := πg.InternStr("type")
		ßxor := πg.InternStr("xor")
		var πTemp001 []*πg.Object
		_ = πTemp001
		var πTemp002 *πg.Object
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
		var πTemp038 *πg.Dict
		_ = πTemp038
		var πTemp039 *πg.Object
		_ = πTemp039
		var πTemp040 *πg.Object
		_ = πTemp040
		var πTemp041 *πg.Object
		_ = πTemp041
		var πTemp042 *πg.Object
		_ = πTemp042
		var πTemp043 *πg.Object
		_ = πTemp043
		var πTemp044 *πg.Object
		_ = πTemp044
		var πTemp045 *πg.Object
		_ = πTemp045
		var πTemp046 *πg.Object
		_ = πTemp046
		var πTemp047 *πg.Object
		_ = πTemp047
		var πTemp048 *πg.Object
		_ = πTemp048
		var πTemp049 *πg.Object
		_ = πTemp049
		var πTemp050 *πg.Object
		_ = πTemp050
		var πTemp051 *πg.Object
		_ = πTemp051
		var πTemp052 *πg.Object
		_ = πTemp052
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """
			πF.SetLineno(1)
			// line 11: __all__ = ['abs', 'add', 'and_', 'attrgetter', 'concat', 'contains', 'countOf',
			πF.SetLineno(11)
			πTemp001 = make([]*πg.Object, 52)
			πTemp001[0] = ßabs.ToObject()
			πTemp001[1] = ßadd.ToObject()
			πTemp001[2] = ßand_.ToObject()
			πTemp001[3] = ßattrgetter.ToObject()
			πTemp001[4] = ßconcat.ToObject()
			πTemp001[5] = ßcontains.ToObject()
			πTemp001[6] = ßcountOf.ToObject()
			πTemp001[7] = ßdelitem.ToObject()
			πTemp001[8] = ßeq.ToObject()
			πTemp001[9] = ßfloordiv.ToObject()
			πTemp001[10] = ßge.ToObject()
			πTemp001[11] = ßgetitem.ToObject()
			πTemp001[12] = ßgt.ToObject()
			πTemp001[13] = ßiadd.ToObject()
			πTemp001[14] = ßiand.ToObject()
			πTemp001[15] = ßiconcat.ToObject()
			πTemp001[16] = ßifloordiv.ToObject()
			πTemp001[17] = ßilshift.ToObject()
			πTemp001[18] = ßimod.ToObject()
			πTemp001[19] = ßimul.ToObject()
			πTemp001[20] = ßindex.ToObject()
			πTemp001[21] = ßindexOf.ToObject()
			πTemp001[22] = ßinv.ToObject()
			πTemp001[23] = ßinvert.ToObject()
			πTemp001[24] = ßior.ToObject()
			πTemp001[25] = ßipow.ToObject()
			πTemp001[26] = ßirshift.ToObject()
			πTemp001[27] = ßis_.ToObject()
			πTemp001[28] = ßis_not.ToObject()
			πTemp001[29] = ßisub.ToObject()
			πTemp001[30] = ßitemgetter.ToObject()
			πTemp001[31] = ßitruediv.ToObject()
			πTemp001[32] = ßixor.ToObject()
			πTemp001[33] = ßle.ToObject()
			πTemp001[34] = ßlength_hint.ToObject()
			πTemp001[35] = ßlshift.ToObject()
			πTemp001[36] = ßlt.ToObject()
			πTemp001[37] = ßmethodcaller.ToObject()
			πTemp001[38] = ßmod.ToObject()
			πTemp001[39] = ßmul.ToObject()
			πTemp001[40] = ßne.ToObject()
			πTemp001[41] = ßneg.ToObject()
			πTemp001[42] = ßnot_.ToObject()
			πTemp001[43] = ßor_.ToObject()
			πTemp001[44] = ßpos.ToObject()
			πTemp001[45] = ßpow.ToObject()
			πTemp001[46] = ßrshift.ToObject()
			πTemp001[47] = ßsetitem.ToObject()
			πTemp001[48] = ßsub.ToObject()
			πTemp001[49] = ßtruediv.ToObject()
			πTemp001[50] = ßtruth.ToObject()
			πTemp001[51] = ßxor.ToObject()
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 20: from '__go__/math' import Abs as _abs
			πF.SetLineno(20)
			if πTemp001, πE = πg.ImportModule(πF, "__go__/math"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßAbs, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_abs.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 25: def lt(a, b):
			πF.SetLineno(25)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp002 = πg.NewFunction(πg.NewCode("lt", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 26: "Same as a < b."
					πF.SetLineno(26)
					// line 27: return a < b
					πF.SetLineno(27)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µa, µb); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßlt.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 29: def le(a, b):
			πF.SetLineno(29)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("le", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 30: "Same as a <= b."
					πF.SetLineno(30)
					// line 31: return a <= b
					πF.SetLineno(31)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LE(πF, µa, µb); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßle.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 33: def eq(a, b):
			πF.SetLineno(33)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("eq", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 34: "Same as a == b."
					πF.SetLineno(34)
					// line 35: return a == b
					πF.SetLineno(35)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µa, µb); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßeq.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 37: def ne(a, b):
			πF.SetLineno(37)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("ne", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 38: "Same as a != b."
					πF.SetLineno(38)
					// line 39: return a != b
					πF.SetLineno(39)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, µa, µb); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßne.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 41: def ge(a, b):
			πF.SetLineno(41)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("ge", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 42: "Same as a >= b."
					πF.SetLineno(42)
					// line 43: return a >= b
					πF.SetLineno(43)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GE(πF, µa, µb); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßge.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 45: def gt(a, b):
			πF.SetLineno(45)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("gt", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 46: "Same as a > b."
					πF.SetLineno(46)
					// line 47: return a > b
					πF.SetLineno(47)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GT(πF, µa, µb); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßgt.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 51: def not_(a):
			πF.SetLineno(51)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("not_", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 52: "Same as not a."
					πF.SetLineno(52)
					// line 53: return not a
					πF.SetLineno(53)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µa); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßnot_.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 55: def truth(a):
			πF.SetLineno(55)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("truth", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
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
					// line 56: "Return True if a is true, False otherwise."
					πF.SetLineno(56)
					// line 57: return True if a else False
					πF.SetLineno(57)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µa); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					goto Label2
				Label1:
					if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp001 = πTemp003
				Label2:
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
			if πE = πF.Globals().SetItem(πF, ßtruth.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 59: def is_(a, b):
			πF.SetLineno(59)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp011 = πg.NewFunction(πg.NewCode("is_", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 60: "Same as a is b."
					πF.SetLineno(60)
					// line 61: return a is b
					πF.SetLineno(61)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µa == µb).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßis_.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 63: def is_not(a, b):
			πF.SetLineno(63)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("is_not", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 64: "Same as a is not b."
					πF.SetLineno(64)
					// line 65: return a is not b
					πF.SetLineno(65)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µa != µb).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßis_not.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 69: def abs(a):
			πF.SetLineno(69)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp013 = πg.NewFunction(πg.NewCode("abs", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var πTemp001 []*πg.Object
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
					// line 70: "Same as abs(a)."
					πF.SetLineno(70)
					// line 71: return _abs(a)
					πF.SetLineno(71)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp001[0] = µa
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_abs); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
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
			if πE = πF.Globals().SetItem(πF, ßabs.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 73: def add(a, b):
			πF.SetLineno(73)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp014 = πg.NewFunction(πg.NewCode("add", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 74: "Same as a + b."
					πF.SetLineno(74)
					// line 75: return a + b
					πF.SetLineno(75)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µa, µb); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßadd.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 77: def and_(a, b):
			πF.SetLineno(77)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp015 = πg.NewFunction(πg.NewCode("and_", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 78: "Same as a & b."
					πF.SetLineno(78)
					// line 79: return a & b
					πF.SetLineno(79)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.And(πF, µa, µb); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßand_.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 81: def floordiv(a, b):
			πF.SetLineno(81)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp016 = πg.NewFunction(πg.NewCode("floordiv", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 82: "Same as a // b."
					πF.SetLineno(82)
					// line 83: return a // b
					πF.SetLineno(83)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.FloorDiv(πF, µa, µb); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßfloordiv.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 85: def index(a):
			πF.SetLineno(85)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp017 = πg.NewFunction(πg.NewCode("index", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
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
					// line 86: "Same as a.__index__()."
					πF.SetLineno(86)
					// line 87: return a.__index__()
					πF.SetLineno(87)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µa, ß__index__, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßindex.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 89: def inv(a):
			πF.SetLineno(89)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp018 = πg.NewFunction(πg.NewCode("inv", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 90: "Same as ~a."
					πF.SetLineno(90)
					// line 91: return ~a
					πF.SetLineno(91)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Invert(πF, µa); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßinv.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 92: invert = inv
			πF.SetLineno(92)
			if πTemp019, πE = πg.ResolveGlobal(πF, ßinv); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßinvert.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 94: def lshift(a, b):
			πF.SetLineno(94)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp019 = πg.NewFunction(πg.NewCode("lshift", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 95: "Same as a << b."
					πF.SetLineno(95)
					// line 96: return a << b
					πF.SetLineno(96)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LShift(πF, µa, µb); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßlshift.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 98: def mod(a, b):
			πF.SetLineno(98)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp020 = πg.NewFunction(πg.NewCode("mod", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 99: "Same as a % b."
					πF.SetLineno(99)
					// line 100: return a % b
					πF.SetLineno(100)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, µa, µb); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßmod.ToObject(), πTemp020); πE != nil {
				continue
			}
			// line 102: def mul(a, b):
			πF.SetLineno(102)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp021 = πg.NewFunction(πg.NewCode("mul", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 103: "Same as a * b."
					πF.SetLineno(103)
					// line 104: return a * b
					πF.SetLineno(104)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mul(πF, µa, µb); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßmul.ToObject(), πTemp021); πE != nil {
				continue
			}
			// line 106: def neg(a):
			πF.SetLineno(106)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp022 = πg.NewFunction(πg.NewCode("neg", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 107: "Same as -a."
					πF.SetLineno(107)
					// line 108: return -a
					πF.SetLineno(108)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Neg(πF, µa); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßneg.ToObject(), πTemp022); πE != nil {
				continue
			}
			// line 110: def or_(a, b):
			πF.SetLineno(110)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp023 = πg.NewFunction(πg.NewCode("or_", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 111: "Same as a | b."
					πF.SetLineno(111)
					// line 112: return a | b
					πF.SetLineno(112)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Or(πF, µa, µb); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßor_.ToObject(), πTemp023); πE != nil {
				continue
			}
			// line 114: def pos(a):
			πF.SetLineno(114)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp024 = πg.NewFunction(πg.NewCode("pos", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 115: "Same as +a."
					πF.SetLineno(115)
					// line 116: return +a
					πF.SetLineno(116)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Pos(πF, µa); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßpos.ToObject(), πTemp024); πE != nil {
				continue
			}
			// line 118: def pow(a, b):
			πF.SetLineno(118)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp025 = πg.NewFunction(πg.NewCode("pow", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 119: "Same as a ** b."
					πF.SetLineno(119)
					// line 120: return a**b
					πF.SetLineno(120)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Pow(πF, µa, µb); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßpow.ToObject(), πTemp025); πE != nil {
				continue
			}
			// line 122: def rshift(a, b):
			πF.SetLineno(122)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp026 = πg.NewFunction(πg.NewCode("rshift", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 123: "Same as a >> b."
					πF.SetLineno(123)
					// line 124: return a >> b
					πF.SetLineno(124)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.RShift(πF, µa, µb); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßrshift.ToObject(), πTemp026); πE != nil {
				continue
			}
			// line 126: def sub(a, b):
			πF.SetLineno(126)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp027 = πg.NewFunction(πg.NewCode("sub", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 127: "Same as a - b."
					πF.SetLineno(127)
					// line 128: return a - b
					πF.SetLineno(128)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, µa, µb); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßsub.ToObject(), πTemp027); πE != nil {
				continue
			}
			// line 130: def truediv(a, b):
			πF.SetLineno(130)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp028 = πg.NewFunction(πg.NewCode("truediv", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
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
					// line 131: "Same as a / b."
					πF.SetLineno(131)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp004[0] = µa
					if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp006, πTemp005); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp004[0] = µa
					if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp006, πTemp005); πE != nil {
						continue
					}
					πTemp001 = πTemp003
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 132: if type(a) == int or type(a) == long:
					πF.SetLineno(132)
				Label2:
					// line 133: a = float(a)
					πF.SetLineno(133)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp004[0] = µa
					if πTemp001, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µa = πTemp003
					goto Label3
				Label3:
					// line 134: return a / b
					πF.SetLineno(134)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Div(πF, µa, µb); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßtruediv.ToObject(), πTemp028); πE != nil {
				continue
			}
			// line 136: def xor(a, b):
			πF.SetLineno(136)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp029 = πg.NewFunction(πg.NewCode("xor", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 137: "Same as a ^ b."
					πF.SetLineno(137)
					// line 138: return a ^ b
					πF.SetLineno(138)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Xor(πF, µa, µb); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßxor.ToObject(), πTemp029); πE != nil {
				continue
			}
			// line 142: def concat(a, b):
			πF.SetLineno(142)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp030 = πg.NewFunction(πg.NewCode("concat", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var µmsg *πg.Object = πg.UnboundLocal; _ = µmsg
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 143: "Same as a + b, for a and b sequences."
					πF.SetLineno(143)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp002[0] = µa
					πTemp002[1] = ß__getitem__.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 144: if not hasattr(a, '__getitem__'):
					πF.SetLineno(144)
				Label1:
					// line 145: msg = "'%s' object can't be concatenated" % type(a).__name__
					πF.SetLineno(145)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp002[0] = µa
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ß__name__, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("'%s' object can't be concatenated").ToObject(), πTemp003); πE != nil {
						continue
					}
					µmsg = πTemp001
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp002[0] = µmsg
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 146: raise TypeError(msg)
					πF.SetLineno(146)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label2
				Label2:
					// line 147: return a + b
					πF.SetLineno(147)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µa, µb); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßconcat.ToObject(), πTemp030); πE != nil {
				continue
			}
			// line 149: def contains(a, b):
			πF.SetLineno(149)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp031 = πg.NewFunction(πg.NewCode("contains", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 150: "Same as b in a (note reversed operands)."
					πF.SetLineno(150)
					// line 151: return b in a
					πF.SetLineno(151)
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, µa, µb); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßcontains.ToObject(), πTemp031); πE != nil {
				continue
			}
			// line 153: def countOf(a, b):
			πF.SetLineno(153)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp032 = πg.NewFunction(πg.NewCode("countOf", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var µcount *πg.Object = πg.UnboundLocal; _ = µcount
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 154: "Return the number of times b occurs in a."
					πF.SetLineno(154)
					// line 155: count = 0
					πF.SetLineno(155)
					µcount = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µa); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp002 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
						πF.PopCheckpoint()
						goto Label3
					}
					if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp003 = !isStop
					} else {
						πTemp003 = true
						µi = πTemp004
					}
					if πE != nil || !πTemp003 {
						continue
					}
					πF.PushCheckpoint(1)            
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, µi, µb); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label4
					}
					goto Label5
					// line 157: if i == b:
					πF.SetLineno(157)
				Label4:
					// line 158: count += 1
					πF.SetLineno(158)
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IAdd(πF, µcount, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µcount = πTemp004
					goto Label5
				Label5:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 159: return count
					πF.SetLineno(159)
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					πR = µcount
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßcountOf.ToObject(), πTemp032); πE != nil {
				continue
			}
			// line 161: def delitem(a, b):
			πF.SetLineno(161)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp033 = πg.NewFunction(πg.NewCode("delitem", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 162: "Same as del a[b]."
					πF.SetLineno(162)
					// line 163: del a[b]
					πF.SetLineno(163)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					πTemp001 = µb
					if πE = πg.DelItem(πF, µa, πTemp001); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßdelitem.ToObject(), πTemp033); πE != nil {
				continue
			}
			// line 165: def getitem(a, b):
			πF.SetLineno(165)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp034 = πg.NewFunction(πg.NewCode("getitem", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
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
					// line 166: "Same as a[b]."
					πF.SetLineno(166)
					// line 167: return a[b]
					πF.SetLineno(167)
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					πTemp001 = µb
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µa, πTemp001); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßgetitem.ToObject(), πTemp034); πE != nil {
				continue
			}
			// line 169: def indexOf(a, b):
			πF.SetLineno(169)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp035 = πg.NewFunction(πg.NewCode("indexOf", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µj *πg.Object = πg.UnboundLocal; _ = µj
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 bool
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
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 170: "Return the first index of b in a."
					πF.SetLineno(170)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp002[0] = µa
					if πTemp003, πE = πg.ResolveGlobal(πF, ßenumerate); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp005 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp005 {
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
						πTemp006 = !isStop
					} else {
						πTemp006 = true
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
							continue
						}
						µi = πTemp004
						µj = πTemp007
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(1)            
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µj, µb); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label4
					}
					goto Label5
					// line 172: if j == b:
					πF.SetLineno(172)
				Label4:
					// line 173: return i
					πF.SetLineno(173)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πR = µi
					continue
					goto Label5
				Label5:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("sequence.index(x): x not in sequence").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 175: raise ValueError('sequence.index(x): x not in sequence')
					πF.SetLineno(175)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
				Label3:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßindexOf.ToObject(), πTemp035); πE != nil {
				continue
			}
			// line 177: def setitem(a, b, c):
			πF.SetLineno(177)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp004[2] = πg.Param{Name: "c", Def: nil}
			πTemp036 = πg.NewFunction(πg.NewCode("setitem", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var µc *πg.Object = πArgs[2]; _ = µc
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
					// line 178: "Same as a[b] = c."
					πF.SetLineno(178)
					// line 179: a[b] = c
					πF.SetLineno(179)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µc); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					πTemp002 = µb
					if πE = πg.SetItem(πF, µa, πTemp002, πTemp001); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßsetitem.ToObject(), πTemp036); πE != nil {
				continue
			}
			// line 181: def length_hint(obj, default=0):
			πF.SetLineno(181)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			πTemp004[1] = πg.Param{Name: "default", Def: πg.NewInt(0).ToObject()}
			πTemp037 = πg.NewFunction(πg.NewCode("length_hint", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µdefault *πg.Object = πArgs[1]; _ = µdefault
				var µmsg *πg.Object = πg.UnboundLocal; _ = µmsg
				var µhint *πg.Object = πg.UnboundLocal; _ = µhint
				var µval *πg.Object = πg.UnboundLocal; _ = µval
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 *πg.BaseException
				_ = πTemp006
				var πTemp007 *πg.Traceback
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 10: goto Label10
					case 4: goto Label4
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					// line 182: """
					πF.SetLineno(182)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
						continue
					}
					πTemp002[0] = µdefault
					if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					πTemp002[1] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 189: if not isinstance(default, int):
					πF.SetLineno(189)
				Label1:
					// line 190: msg = ("'%s' object cannot be interpreted as an integer" %
					πF.SetLineno(190)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
						continue
					}
					πTemp002[0] = µdefault
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ß__name__, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("'%s' object cannot be interpreted as an integer").ToObject(), πTemp003); πE != nil {
						continue
					}
					µmsg = πTemp001
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp002[0] = µmsg
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 192: raise TypeError(msg)
					πF.SetLineno(192)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label2
				Label2:
					// line 194: try:
					πF.SetLineno(194)
					πF.PushCheckpoint(4)
					// line 195: return len(obj)
					πF.SetLineno(195)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πR = πTemp003
					continue
					πF.PopCheckpoint()
					goto Label3
				Label4:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp006, πTemp007 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label5
					}
					πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
					continue
					// line 196: except TypeError:
					πF.SetLineno(196)
				Label5:
					// line 197: pass
					πF.SetLineno(197)
					πF.RestoreExc(nil, nil)
					goto Label3
				Label3:
					// line 199: try:
					πF.SetLineno(199)
					πF.PushCheckpoint(7)
					// line 200: hint = type(obj).__length_hint__
					πF.SetLineno(200)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.GetAttr(πF, πTemp003, ß__length_hint__, nil); πE != nil {
						continue
					}
					µhint = πTemp001
					πF.PopCheckpoint()
					goto Label6
				Label7:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp006, πTemp007 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label8
					}
					πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
					continue
					// line 201: except AttributeError:
					πF.SetLineno(201)
				Label8:
					// line 202: return default
					πF.SetLineno(202)
					if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
						continue
					}
					πR = µdefault
					continue
					πF.RestoreExc(nil, nil)
					goto Label6
				Label6:
					// line 204: try:
					πF.SetLineno(204)
					πF.PushCheckpoint(10)
					// line 205: val = hint(obj)
					πF.SetLineno(205)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp002[0] = µobj
					if πE = πg.CheckLocal(πF, µhint, "hint"); πE != nil {
						continue
					}
					if πTemp001, πE = µhint.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µval = πTemp001
					πF.PopCheckpoint()
					goto Label9
				Label10:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp006, πTemp007 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label11
					}
					πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
					continue
					// line 206: except TypeError:
					πF.SetLineno(206)
				Label11:
					// line 207: return default
					πF.SetLineno(207)
					if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
						continue
					}
					πR = µdefault
					continue
					πF.RestoreExc(nil, nil)
					goto Label9
				Label9:
					if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µval == πTemp003).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label12
					}
					goto Label13
					// line 208: if val is NotImplemented:
					πF.SetLineno(208)
				Label12:
					// line 209: return default
					πF.SetLineno(209)
					if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
						continue
					}
					πR = µdefault
					continue
					goto Label13
				Label13:
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
						continue
					}
					πTemp002[0] = µval
					if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					πTemp002[1] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label14
					}
					goto Label15
					// line 210: if not isinstance(val, int):
					πF.SetLineno(210)
				Label14:
					// line 211: msg = ('__length_hint__ must be integer, not %s' %
					πF.SetLineno(211)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
						continue
					}
					πTemp002[0] = µval
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ß__name__, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("__length_hint__ must be integer, not %s").ToObject(), πTemp003); πE != nil {
						continue
					}
					µmsg = πTemp001
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp002[0] = µmsg
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 213: raise TypeError(msg)
					πF.SetLineno(213)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label15
				Label15:
					if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µval, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label16
					}
					goto Label17
					// line 214: if val < 0:
					πF.SetLineno(214)
				Label16:
					// line 215: msg = '__length_hint__() should return >= 0'
					πF.SetLineno(215)
					µmsg = πg.NewStr("__length_hint__() should return >= 0").ToObject()
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp002[0] = µmsg
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 216: raise ValueError(msg)
					πF.SetLineno(216)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label17
				Label17:
					// line 217: return val
					πF.SetLineno(217)
					if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
						continue
					}
					πR = µval
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßlength_hint.ToObject(), πTemp037); πE != nil {
				continue
			}
			// line 222: class attrgetter(object):
			πF.SetLineno(222)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp041, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp001[0] = πTemp041
			πTemp038 = πg.NewDict()
			if πTemp039, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp038.SetItem(πF, ß__module__.ToObject(), πTemp039); πE != nil {
				continue
			}
			_, πE = πg.NewCode("attrgetter", "build/src/__python__/operator.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp038
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 223: """
					πF.SetLineno(223)
					// line 230: def __init__(self, attr, *attrs):
					πF.SetLineno(230)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "attr", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/operator.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µattr *πg.Object = πArgs[1]; _ = µattr
						var µattrs *πg.Object = πArgs[2]; _ = µattrs
						var µnames *πg.Object = πg.UnboundLocal; _ = µnames
						var µfunc *πg.Object = πg.UnboundLocal; _ = µfunc
						var µgetters *πg.Object = πg.UnboundLocal; _ = µgetters
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
						var πTemp006 []πg.Param
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µattrs, "attrs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µattrs); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 231: if not attrs:
							πF.SetLineno(231)
						Label1:
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp003[0] = µattr
							if πTemp004, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							goto Label5
							// line 232: if not isinstance(attr, str):
							πF.SetLineno(232)
						Label4:
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("attribute name must be a string").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 233: raise TypeError('attribute name must be a string')
							πF.SetLineno(233)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label5
						Label5:
							// line 234: names = attr.split('.')
							πF.SetLineno(234)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr(".").ToObject()
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µattr, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µnames = πTemp004
							// line 235: def func(obj):
							πF.SetLineno(235)
							πTemp006 = make([]πg.Param, 1)
							πTemp006[0] = πg.Param{Name: "obj", Def: nil}
							πTemp001 = πg.NewFunction(πg.NewCode("func", "build/src/__python__/operator.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µobj *πg.Object = πArgs[0]; _ = µobj
								var µname *πg.Object = πg.UnboundLocal; _ = µname
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 bool
								_ = πTemp002
								var πTemp003 bool
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
									case 1: goto Label1
									case 2: goto Label2
									default: panic("unexpected function state")
									}
									if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.Iter(πF, µnames); πE != nil {
										continue
									}
									πF.PushCheckpoint(2)
									πTemp002 = false
								Label1:
									if πE != nil || πR != nil {
										continue
									}
									if πTemp002 {
										πF.PopCheckpoint()
										goto Label3
									}
									if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
										isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
										if exc != nil {
											πE = exc
										} else if isStop {
											πE = nil
											πF.RestoreExc(nil, nil)
										}
										πTemp003 = !isStop
									} else {
										πTemp003 = true
										µname = πTemp004
									}
									if πE != nil || !πTemp003 {
										continue
									}
									πF.PushCheckpoint(1)            
									// line 237: obj = getattr(obj, name)
									πF.SetLineno(237)
									πTemp005 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
										continue
									}
									πTemp005[0] = µobj
									if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
										continue
									}
									πTemp005[1] = µname
									if πTemp004, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
										continue
									}
									if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp005)
									µobj = πTemp006
									continue
								Label2:
									if πE != nil || πR != nil {
										continue
									}
								Label3:
									// line 238: return obj
									πF.SetLineno(238)
									if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
										continue
									}
									πR = µobj
									continue
								}
								if πE != nil {
									πR = nil
								} else if πR == nil {
									πR = πg.None
								}
								return πR, πE
							}), πF.Globals()).ToObject()
							µfunc = πTemp001
							// line 239: self._call = func
							πF.SetLineno(239)
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µfunc); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_call, πTemp004); πE != nil {
								continue
							}
							goto Label3
						Label2:
							// line 241: getters = tuple(map(attrgetter, (attr,) + attrs))
							πF.SetLineno(241)
							πTemp003 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(2)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßattrgetter); πE != nil {
								continue
							}
							πTemp007[0] = πTemp004
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple1(µattr).ToObject()
							if πE = πg.CheckLocal(πF, µattrs, "attrs"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp005, µattrs); πE != nil {
								continue
							}
							πTemp007[1] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp003[0] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µgetters = πTemp005
							// line 242: def func(obj):
							πF.SetLineno(242)
							πTemp006 = make([]πg.Param, 1)
							πTemp006[0] = πg.Param{Name: "obj", Def: nil}
							πTemp004 = πg.NewFunction(πg.NewCode("func", "build/src/__python__/operator.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µobj *πg.Object = πArgs[0]; _ = µobj
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
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 243: return tuple(getter(obj) for getter in getters)
									πF.SetLineno(243)
									πTemp001 = πF.MakeArgs(1)
									πTemp003 = make([]πg.Param, 0)
									πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/operator.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µgetter *πg.Object = πg.UnboundLocal; _ = µgetter
										var πTemp001 *πg.Object
										_ = πTemp001
										var πTemp002 bool
										_ = πTemp002
										var πTemp003 bool
										_ = πTemp003
										var πTemp004 *πg.Object
										_ = πTemp004
										var πTemp005 []*πg.Object
										_ = πTemp005
										var πTemp006 *πg.Object
										_ = πTemp006
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
											for ; πF.State() >= 0; πF.PopCheckpoint() {
												switch πF.State() {
												case 0:
												case 1: goto Label1
												case 2: goto Label2
												case 4: goto Label4
												default: panic("unexpected function state")
												}
												if πE = πg.CheckLocal(πF, µgetters, "getters"); πE != nil {
													continue
												}
												if πTemp001, πE = πg.Iter(πF, µgetters); πE != nil {
													continue
												}
												πF.PushCheckpoint(2)
												πTemp002 = false
											Label1:
												if πE != nil || πR != nil {
													continue
												}
												if πTemp002 {
													πF.PopCheckpoint()
													goto Label3
												}
												if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
													isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
													if exc != nil {
														πE = exc
													} else if isStop {
														πE = nil
														πF.RestoreExc(nil, nil)
													}
													πTemp003 = !isStop
												} else {
													πTemp003 = true
													µgetter = πTemp004
												}
												if πE != nil || !πTemp003 {
													continue
												}
												πF.PushCheckpoint(1)            
												// line 243: return tuple(getter(obj) for getter in getters)
												πF.SetLineno(243)
												πTemp005 = πF.MakeArgs(1)
												if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
													continue
												}
												πTemp005[0] = µobj
												if πE = πg.CheckLocal(πF, µgetter, "getter"); πE != nil {
													continue
												}
												if πTemp004, πE = µgetter.Call(πF, πTemp005, nil); πE != nil {
													continue
												}
												πF.FreeArgs(πTemp005)
												πF.PushCheckpoint(4)
												return πTemp004, nil
											Label4:
												πTemp006 = πSent
												continue
											Label2:
												if πE != nil || πR != nil {
													continue
												}
											Label3:
											}
											return nil, πE
										}).ToObject(), nil
									}), πF.Globals()).ToObject()
									if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
										continue
									}
									πTemp001[0] = πTemp004
									if πTemp004, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
										continue
									}
									if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									πR = πTemp005
									continue
								}
								if πE != nil {
									πR = nil
								} else if πR == nil {
									πR = πg.None
								}
								return πR, πE
							}), πF.Globals()).ToObject()
							µfunc = πTemp004
							// line 244: self._call = func
							πF.SetLineno(244)
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, µfunc); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_call, πTemp005); πE != nil {
								continue
							}
							goto Label3
						Label3:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 246: def __call__(self, obj):
					πF.SetLineno(246)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__call__", "build/src/__python__/operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µobj *πg.Object = πArgs[1]; _ = µobj
						var πTemp001 []*πg.Object
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
							// line 247: return self._call(obj)
							πF.SetLineno(247)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_call, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ß__call__.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp040, πE = πTemp038.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp040 == nil {
				πTemp040 = πg.TypeType.ToObject()
			}
			if πTemp041, πE = πTemp040.Call(πF, []*πg.Object{πg.NewStr("attrgetter").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp038.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßattrgetter.ToObject(), πTemp041); πE != nil {
				continue
			}
			// line 250: class itemgetter(object):
			πF.SetLineno(250)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp041, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp001[0] = πTemp041
			πTemp038 = πg.NewDict()
			if πTemp039, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp038.SetItem(πF, ß__module__.ToObject(), πTemp039); πE != nil {
				continue
			}
			_, πE = πg.NewCode("itemgetter", "build/src/__python__/operator.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp038
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 251: """
					πF.SetLineno(251)
					// line 256: def __init__(self, item, *items):
					πF.SetLineno(256)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "item", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/operator.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µitem *πg.Object = πArgs[1]; _ = µitem
						var µitems *πg.Object = πArgs[2]; _ = µitems
						var µfunc *πg.Object = πg.UnboundLocal; _ = µfunc
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 []πg.Param
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
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µitems); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 257: if not items:
							πF.SetLineno(257)
						Label1:
							// line 258: def func(obj):
							πF.SetLineno(258)
							πTemp003 = make([]πg.Param, 1)
							πTemp003[0] = πg.Param{Name: "obj", Def: nil}
							πTemp001 = πg.NewFunction(πg.NewCode("func", "build/src/__python__/operator.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µobj *πg.Object = πArgs[0]; _ = µobj
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
									// line 259: return obj[item]
									πF.SetLineno(259)
									if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
										continue
									}
									πTemp001 = µitem
									if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetItem(πF, µobj, πTemp001); πE != nil {
										continue
									}
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
							µfunc = πTemp001
							// line 260: self._call = func
							πF.SetLineno(260)
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µfunc); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_call, πTemp004); πE != nil {
								continue
							}
							goto Label3
						Label2:
							// line 262: items = (item,) + items
							πF.SetLineno(262)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple1(µitem).ToObject()
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp005, µitems); πE != nil {
								continue
							}
							µitems = πTemp004
							// line 263: def func(obj):
							πF.SetLineno(263)
							πTemp003 = make([]πg.Param, 1)
							πTemp003[0] = πg.Param{Name: "obj", Def: nil}
							πTemp004 = πg.NewFunction(πg.NewCode("func", "build/src/__python__/operator.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µobj *πg.Object = πArgs[0]; _ = µobj
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
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 264: return tuple(obj[i] for i in items)
									πF.SetLineno(264)
									πTemp001 = πF.MakeArgs(1)
									πTemp003 = make([]πg.Param, 0)
									πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/operator.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µi *πg.Object = πg.UnboundLocal; _ = µi
										var πTemp001 *πg.Object
										_ = πTemp001
										var πTemp002 bool
										_ = πTemp002
										var πTemp003 bool
										_ = πTemp003
										var πTemp004 *πg.Object
										_ = πTemp004
										var πTemp005 *πg.Object
										_ = πTemp005
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
											for ; πF.State() >= 0; πF.PopCheckpoint() {
												switch πF.State() {
												case 0:
												case 1: goto Label1
												case 2: goto Label2
												case 4: goto Label4
												default: panic("unexpected function state")
												}
												if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
													continue
												}
												if πTemp001, πE = πg.Iter(πF, µitems); πE != nil {
													continue
												}
												πF.PushCheckpoint(2)
												πTemp002 = false
											Label1:
												if πE != nil || πR != nil {
													continue
												}
												if πTemp002 {
													πF.PopCheckpoint()
													goto Label3
												}
												if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
													isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
													if exc != nil {
														πE = exc
													} else if isStop {
														πE = nil
														πF.RestoreExc(nil, nil)
													}
													πTemp003 = !isStop
												} else {
													πTemp003 = true
													µi = πTemp004
												}
												if πE != nil || !πTemp003 {
													continue
												}
												πF.PushCheckpoint(1)            
												// line 264: return tuple(obj[i] for i in items)
												πF.SetLineno(264)
												if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
													continue
												}
												πTemp004 = µi
												if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
													continue
												}
												if πTemp005, πE = πg.GetItem(πF, µobj, πTemp004); πE != nil {
													continue
												}
												πF.PushCheckpoint(4)
												return πTemp005, nil
											Label4:
												πTemp004 = πSent
												continue
											Label2:
												if πE != nil || πR != nil {
													continue
												}
											Label3:
											}
											return nil, πE
										}).ToObject(), nil
									}), πF.Globals()).ToObject()
									if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
										continue
									}
									πTemp001[0] = πTemp004
									if πTemp004, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
										continue
									}
									if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									πR = πTemp005
									continue
								}
								if πE != nil {
									πR = nil
								} else if πR == nil {
									πR = πg.None
								}
								return πR, πE
							}), πF.Globals()).ToObject()
							µfunc = πTemp004
							// line 265: self._call = func
							πF.SetLineno(265)
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, µfunc); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_call, πTemp005); πE != nil {
								continue
							}
							goto Label3
						Label3:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 267: def __call__(self, obj):
					πF.SetLineno(267)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__call__", "build/src/__python__/operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µobj *πg.Object = πArgs[1]; _ = µobj
						var πTemp001 []*πg.Object
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
							// line 268: return self._call(obj)
							πF.SetLineno(268)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_call, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ß__call__.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp040, πE = πTemp038.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp040 == nil {
				πTemp040 = πg.TypeType.ToObject()
			}
			if πTemp041, πE = πTemp040.Call(πF, []*πg.Object{πg.NewStr("itemgetter").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp038.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßitemgetter.ToObject(), πTemp041); πE != nil {
				continue
			}
			// line 271: class methodcaller(object):
			πF.SetLineno(271)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp041, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp001[0] = πTemp041
			πTemp038 = πg.NewDict()
			if πTemp039, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp038.SetItem(πF, ß__module__.ToObject(), πTemp039); πE != nil {
				continue
			}
			_, πE = πg.NewCode("methodcaller", "build/src/__python__/operator.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp038
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 272: """
					πF.SetLineno(272)
					// line 279: def __init__(*args, **kwargs):
					πF.SetLineno(279)
					πTemp002 = make([]πg.Param, 0)
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/operator.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µargs *πg.Object = πArgs[0]; _ = µargs
						var µkwargs *πg.Object = πArgs[1]; _ = µkwargs
						var µmsg *πg.Object = πg.UnboundLocal; _ = µmsg
						var µself *πg.Object = πg.UnboundLocal; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp002[0] = µargs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.LT(πF, πTemp004, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 280: if len(args) < 2:
							πF.SetLineno(280)
						Label1:
							// line 281: msg = "methodcaller needs at least one argument, the method name"
							πF.SetLineno(281)
							µmsg = πg.NewStr("methodcaller needs at least one argument, the method name").ToObject()
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp002[0] = µmsg
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 282: raise TypeError(msg)
							πF.SetLineno(282)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							// line 283: self = args[0]
							πF.SetLineno(283)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							µself = πTemp003
							// line 284: self._name = args[1]
							πF.SetLineno(284)
							πTemp001 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_name, πTemp001); πE != nil {
								continue
							}
							// line 285: self._args = args[2:]
							πF.SetLineno(285)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(2).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_args, πTemp001); πE != nil {
								continue
							}
							// line 286: self._kwargs = kwargs
							πF.SetLineno(286)
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µkwargs); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_kwargs, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 288: def __call__(self, obj):
					πF.SetLineno(288)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__call__", "build/src/__python__/operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µobj *πg.Object = πArgs[1]; _ = µobj
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
							// line 289: return getattr(obj, self._name)(*self._args, **self._kwargs)
							πF.SetLineno(289)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_args, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_kwargs, nil); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp003[0] = µobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_name, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.Invoke(πF, πTemp005, nil, πTemp001, nil, πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ß__call__.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp040, πE = πTemp038.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp040 == nil {
				πTemp040 = πg.TypeType.ToObject()
			}
			if πTemp041, πE = πTemp040.Call(πF, []*πg.Object{πg.NewStr("methodcaller").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp038.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmethodcaller.ToObject(), πTemp041); πE != nil {
				continue
			}
			// line 293: def iadd(a, b):
			πF.SetLineno(293)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp039 = πg.NewFunction(πg.NewCode("iadd", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 294: "Same as a += b."
					πF.SetLineno(294)
					// line 295: a += b
					πF.SetLineno(295)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µa, µb); πE != nil {
						continue
					}
					µa = πTemp001
					// line 296: return a
					πF.SetLineno(296)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πR = µa
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßiadd.ToObject(), πTemp039); πE != nil {
				continue
			}
			// line 298: def iand(a, b):
			πF.SetLineno(298)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp040 = πg.NewFunction(πg.NewCode("iand", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 299: "Same as a &= b."
					πF.SetLineno(299)
					// line 300: a &= b
					πF.SetLineno(300)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAnd(πF, µa, µb); πE != nil {
						continue
					}
					µa = πTemp001
					// line 301: return a
					πF.SetLineno(301)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πR = µa
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßiand.ToObject(), πTemp040); πE != nil {
				continue
			}
			// line 303: def iconcat(a, b):
			πF.SetLineno(303)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp041 = πg.NewFunction(πg.NewCode("iconcat", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var µmsg *πg.Object = πg.UnboundLocal; _ = µmsg
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 304: "Same as a += b, for a and b sequences."
					πF.SetLineno(304)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp002[0] = µa
					πTemp002[1] = ß__getitem__.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 305: if not hasattr(a, '__getitem__'):
					πF.SetLineno(305)
				Label1:
					// line 306: msg = "'%s' object can't be concatenated" % type(a).__name__
					πF.SetLineno(306)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp002[0] = µa
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ß__name__, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("'%s' object can't be concatenated").ToObject(), πTemp003); πE != nil {
						continue
					}
					µmsg = πTemp001
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp002[0] = µmsg
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 307: raise TypeError(msg)
					πF.SetLineno(307)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label2
				Label2:
					// line 308: a += b
					πF.SetLineno(308)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µa, µb); πE != nil {
						continue
					}
					µa = πTemp001
					// line 309: return a
					πF.SetLineno(309)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πR = µa
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßiconcat.ToObject(), πTemp041); πE != nil {
				continue
			}
			// line 311: def ifloordiv(a, b):
			πF.SetLineno(311)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp042 = πg.NewFunction(πg.NewCode("ifloordiv", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 312: "Same as a //= b."
					πF.SetLineno(312)
					// line 313: a //= b
					πF.SetLineno(313)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IFloorDiv(πF, µa, µb); πE != nil {
						continue
					}
					µa = πTemp001
					// line 314: return a
					πF.SetLineno(314)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πR = µa
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßifloordiv.ToObject(), πTemp042); πE != nil {
				continue
			}
			// line 316: def ilshift(a, b):
			πF.SetLineno(316)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp043 = πg.NewFunction(πg.NewCode("ilshift", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 317: "Same as a <<= b."
					πF.SetLineno(317)
					// line 318: a <<= b
					πF.SetLineno(318)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ILShift(πF, µa, µb); πE != nil {
						continue
					}
					µa = πTemp001
					// line 319: return a
					πF.SetLineno(319)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πR = µa
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßilshift.ToObject(), πTemp043); πE != nil {
				continue
			}
			// line 321: def imod(a, b):
			πF.SetLineno(321)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp044 = πg.NewFunction(πg.NewCode("imod", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 322: "Same as a %= b."
					πF.SetLineno(322)
					// line 323: a %= b
					πF.SetLineno(323)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IMod(πF, µa, µb); πE != nil {
						continue
					}
					µa = πTemp001
					// line 324: return a
					πF.SetLineno(324)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πR = µa
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßimod.ToObject(), πTemp044); πE != nil {
				continue
			}
			// line 326: def imul(a, b):
			πF.SetLineno(326)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp045 = πg.NewFunction(πg.NewCode("imul", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 327: "Same as a *= b."
					πF.SetLineno(327)
					// line 328: a *= b
					πF.SetLineno(328)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IMul(πF, µa, µb); πE != nil {
						continue
					}
					µa = πTemp001
					// line 329: return a
					πF.SetLineno(329)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πR = µa
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßimul.ToObject(), πTemp045); πE != nil {
				continue
			}
			// line 331: def ior(a, b):
			πF.SetLineno(331)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp046 = πg.NewFunction(πg.NewCode("ior", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 332: "Same as a |= b."
					πF.SetLineno(332)
					// line 333: a |= b
					πF.SetLineno(333)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IOr(πF, µa, µb); πE != nil {
						continue
					}
					µa = πTemp001
					// line 334: return a
					πF.SetLineno(334)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πR = µa
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßior.ToObject(), πTemp046); πE != nil {
				continue
			}
			// line 336: def ipow(a, b):
			πF.SetLineno(336)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp047 = πg.NewFunction(πg.NewCode("ipow", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 337: "Same as a **= b."
					πF.SetLineno(337)
					// line 338: a **= b
					πF.SetLineno(338)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IPow(πF, µa, µb); πE != nil {
						continue
					}
					µa = πTemp001
					// line 339: return a
					πF.SetLineno(339)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πR = µa
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßipow.ToObject(), πTemp047); πE != nil {
				continue
			}
			// line 341: def irshift(a, b):
			πF.SetLineno(341)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp048 = πg.NewFunction(πg.NewCode("irshift", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 342: "Same as a >>= b."
					πF.SetLineno(342)
					// line 343: a >>= b
					πF.SetLineno(343)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IRShift(πF, µa, µb); πE != nil {
						continue
					}
					µa = πTemp001
					// line 344: return a
					πF.SetLineno(344)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πR = µa
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßirshift.ToObject(), πTemp048); πE != nil {
				continue
			}
			// line 346: def isub(a, b):
			πF.SetLineno(346)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp049 = πg.NewFunction(πg.NewCode("isub", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 347: "Same as a -= b."
					πF.SetLineno(347)
					// line 348: a -= b
					πF.SetLineno(348)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ISub(πF, µa, µb); πE != nil {
						continue
					}
					µa = πTemp001
					// line 349: return a
					πF.SetLineno(349)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πR = µa
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßisub.ToObject(), πTemp049); πE != nil {
				continue
			}
			// line 351: def itruediv(a, b):
			πF.SetLineno(351)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp050 = πg.NewFunction(πg.NewCode("itruediv", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
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
					// line 352: "Same as a /= b."
					πF.SetLineno(352)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp004[0] = µa
					if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp006, πTemp005); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp004[0] = µa
					if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp006, πTemp005); πE != nil {
						continue
					}
					πTemp001 = πTemp003
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 353: if type(a) == int or type(a) == long:
					πF.SetLineno(353)
				Label2:
					// line 354: a = float(a)
					πF.SetLineno(354)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp004[0] = µa
					if πTemp001, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µa = πTemp003
					goto Label3
				Label3:
					// line 355: a /= b
					πF.SetLineno(355)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IDiv(πF, µa, µb); πE != nil {
						continue
					}
					µa = πTemp001
					// line 356: return a
					πF.SetLineno(356)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πR = µa
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßitruediv.ToObject(), πTemp050); πE != nil {
				continue
			}
			// line 358: def ixor(a, b):
			πF.SetLineno(358)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "a", Def: nil}
			πTemp004[1] = πg.Param{Name: "b", Def: nil}
			πTemp051 = πg.NewFunction(πg.NewCode("ixor", "build/src/__python__/operator.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µb *πg.Object = πArgs[1]; _ = µb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 359: "Same as a ^= b."
					πF.SetLineno(359)
					// line 360: a ^= b
					πF.SetLineno(360)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IXor(πF, µa, µb); πE != nil {
						continue
					}
					µa = πTemp001
					// line 361: return a
					πF.SetLineno(361)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πR = µa
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßixor.ToObject(), πTemp051); πE != nil {
				continue
			}
			// line 373: __lt__ = lt
			πF.SetLineno(373)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßlt); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__lt__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 374: __le__ = le
			πF.SetLineno(374)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßle); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__le__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 375: __eq__ = eq
			πF.SetLineno(375)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßeq); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__eq__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 376: __ne__ = ne
			πF.SetLineno(376)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßne); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__ne__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 377: __ge__ = ge
			πF.SetLineno(377)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßge); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__ge__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 378: __gt__ = gt
			πF.SetLineno(378)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßgt); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__gt__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 379: __not__ = not_
			πF.SetLineno(379)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßnot_); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__not__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 380: __abs__ = abs
			πF.SetLineno(380)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßabs); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__abs__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 381: __add__ = add
			πF.SetLineno(381)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßadd); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__add__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 382: __and__ = and_
			πF.SetLineno(382)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßand_); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__and__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 383: __floordiv__ = floordiv
			πF.SetLineno(383)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßfloordiv); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__floordiv__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 384: __index__ = index
			πF.SetLineno(384)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßindex); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__index__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 385: __inv__ = inv
			πF.SetLineno(385)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßinv); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__inv__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 386: __invert__ = invert
			πF.SetLineno(386)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßinvert); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__invert__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 387: __lshift__ = lshift
			πF.SetLineno(387)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßlshift); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__lshift__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 388: __mod__ = mod
			πF.SetLineno(388)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßmod); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__mod__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 389: __mul__ = mul
			πF.SetLineno(389)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßmul); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__mul__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 390: __neg__ = neg
			πF.SetLineno(390)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßneg); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__neg__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 391: __or__ = or_
			πF.SetLineno(391)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßor_); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__or__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 392: __pos__ = pos
			πF.SetLineno(392)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßpos); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__pos__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 393: __pow__ = pow
			πF.SetLineno(393)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßpow); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__pow__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 394: __rshift__ = rshift
			πF.SetLineno(394)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßrshift); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__rshift__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 395: __sub__ = sub
			πF.SetLineno(395)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßsub); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__sub__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 396: __truediv__ = truediv
			πF.SetLineno(396)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßtruediv); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__truediv__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 397: __xor__ = xor
			πF.SetLineno(397)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßxor); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__xor__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 398: __concat__ = concat
			πF.SetLineno(398)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßconcat); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__concat__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 399: __contains__ = contains
			πF.SetLineno(399)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßcontains); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__contains__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 400: __delitem__ = delitem
			πF.SetLineno(400)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßdelitem); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__delitem__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 401: __getitem__ = getitem
			πF.SetLineno(401)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßgetitem); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__getitem__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 402: __setitem__ = setitem
			πF.SetLineno(402)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßsetitem); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__setitem__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 403: __iadd__ = iadd
			πF.SetLineno(403)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßiadd); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__iadd__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 404: __iand__ = iand
			πF.SetLineno(404)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßiand); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__iand__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 405: __iconcat__ = iconcat
			πF.SetLineno(405)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßiconcat); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__iconcat__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 406: __ifloordiv__ = ifloordiv
			πF.SetLineno(406)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßifloordiv); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__ifloordiv__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 407: __ilshift__ = ilshift
			πF.SetLineno(407)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßilshift); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__ilshift__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 408: __imod__ = imod
			πF.SetLineno(408)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßimod); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__imod__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 409: __imul__ = imul
			πF.SetLineno(409)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßimul); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__imul__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 410: __ior__ = ior
			πF.SetLineno(410)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßior); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__ior__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 411: __ipow__ = ipow
			πF.SetLineno(411)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßipow); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__ipow__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 412: __irshift__ = irshift
			πF.SetLineno(412)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßirshift); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__irshift__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 413: __isub__ = isub
			πF.SetLineno(413)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßisub); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__isub__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 414: __itruediv__ = itruediv
			πF.SetLineno(414)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßitruediv); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__itruediv__.ToObject(), πTemp052); πE != nil {
				continue
			}
			// line 415: __ixor__ = ixor
			πF.SetLineno(415)
			if πTemp052, πE = πg.ResolveGlobal(πF, ßixor); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__ixor__.ToObject(), πTemp052); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("operator", Code)
}
