package random
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/random.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßBPF := πg.InternStr("BPF")
		ßGrumpyRandom := πg.InternStr("GrumpyRandom")
		ßNone := πg.InternStr("None")
		ßNotImplementedError := πg.InternStr("NotImplementedError")
		ßRECIP_BPF := πg.InternStr("RECIP_BPF")
		ßRandom := πg.InternStr("Random")
		ßSystemRandom := πg.InternStr("SystemRandom")
		ßVERSION := πg.InternStr("VERSION")
		ßValueError := πg.InternStr("ValueError")
		ßWichmannHill := πg.InternStr("WichmannHill")
		ß__all__ := πg.InternStr("__all__")
		ß__init__ := πg.InternStr("__init__")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_inst := πg.InternStr("_inst")
		ß_notimplemented := πg.InternStr("_notimplemented")
		ß_randbelow := πg.InternStr("_randbelow")
		ß_random := πg.InternStr("_random")
		ß_sqrt := πg.InternStr("_sqrt")
		ß_test := πg.InternStr("_test")
		ß_test_generator := πg.InternStr("_test_generator")
		ßbetavariate := πg.InternStr("betavariate")
		ßchoice := πg.InternStr("choice")
		ßchoices := πg.InternStr("choices")
		ßexpovariate := πg.InternStr("expovariate")
		ßgammavariate := πg.InternStr("gammavariate")
		ßgauss := πg.InternStr("gauss")
		ßgauss_next := πg.InternStr("gauss_next")
		ßgetrandbits := πg.InternStr("getrandbits")
		ßgetstate := πg.InternStr("getstate")
		ßint := πg.InternStr("int")
		ßjumpahead := πg.InternStr("jumpahead")
		ßlen := πg.InternStr("len")
		ßlognormvariate := πg.InternStr("lognormvariate")
		ßmax := πg.InternStr("max")
		ßmin := πg.InternStr("min")
		ßnormalvariate := πg.InternStr("normalvariate")
		ßparetovariate := πg.InternStr("paretovariate")
		ßrandint := πg.InternStr("randint")
		ßrandom := πg.InternStr("random")
		ßrandrange := πg.InternStr("randrange")
		ßrange := πg.InternStr("range")
		ßreversed := πg.InternStr("reversed")
		ßround := πg.InternStr("round")
		ßsample := πg.InternStr("sample")
		ßseed := πg.InternStr("seed")
		ßsetstate := πg.InternStr("setstate")
		ßshuffle := πg.InternStr("shuffle")
		ßsuper := πg.InternStr("super")
		ßtime := πg.InternStr("time")
		ßtimes := πg.InternStr("times")
		ßtriangular := πg.InternStr("triangular")
		ßuniform := πg.InternStr("uniform")
		ßvonmisesvariate := πg.InternStr("vonmisesvariate")
		ßweibullvariate := πg.InternStr("weibullvariate")
		ßxrange := πg.InternStr("xrange")
		var πTemp001 []*πg.Object
		_ = πTemp001
		var πTemp002 *πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.Dict
		_ = πTemp004
		var πTemp005 *πg.Object
		_ = πTemp005
		var πTemp006 *πg.Object
		_ = πTemp006
		var πTemp007 []πg.Param
		_ = πTemp007
		var πTemp008 *πg.Object
		_ = πTemp008
		var πTemp009 bool
		_ = πTemp009
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """Random variable generators.
			πF.SetLineno(1)
			// line 50: __all__ = ["Random","seed","random","uniform","randint","choice","sample",
			πF.SetLineno(50)
			πTemp001 = make([]*πg.Object, 25)
			πTemp001[0] = ßRandom.ToObject()
			πTemp001[1] = ßseed.ToObject()
			πTemp001[2] = ßrandom.ToObject()
			πTemp001[3] = ßuniform.ToObject()
			πTemp001[4] = ßrandint.ToObject()
			πTemp001[5] = ßchoice.ToObject()
			πTemp001[6] = ßsample.ToObject()
			πTemp001[7] = ßrandrange.ToObject()
			πTemp001[8] = ßshuffle.ToObject()
			πTemp001[9] = ßnormalvariate.ToObject()
			πTemp001[10] = ßlognormvariate.ToObject()
			πTemp001[11] = ßexpovariate.ToObject()
			πTemp001[12] = ßvonmisesvariate.ToObject()
			πTemp001[13] = ßgammavariate.ToObject()
			πTemp001[14] = ßtriangular.ToObject()
			πTemp001[15] = ßgauss.ToObject()
			πTemp001[16] = ßbetavariate.ToObject()
			πTemp001[17] = ßparetovariate.ToObject()
			πTemp001[18] = ßweibullvariate.ToObject()
			πTemp001[19] = ßgetstate.ToObject()
			πTemp001[20] = ßsetstate.ToObject()
			πTemp001[21] = ßjumpahead.ToObject()
			πTemp001[22] = ßWichmannHill.ToObject()
			πTemp001[23] = ßgetrandbits.ToObject()
			πTemp001[24] = ßSystemRandom.ToObject()
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 60: import _random
			πF.SetLineno(60)
			if πTemp001, πE = πg.ImportModule(πF, "_random"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ß_random.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 66: BPF = _random.BPF
			πF.SetLineno(66)
			if πTemp002, πE = πg.ResolveGlobal(πF, ß_random); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßBPF, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBPF.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 67: RECIP_BPF = _random.RECIP_BPF
			πF.SetLineno(67)
			if πTemp002, πE = πg.ResolveGlobal(πF, ß_random); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßRECIP_BPF, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßRECIP_BPF.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 70: class Random(_random.GrumpyRandom):
			πF.SetLineno(70)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_random); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßGrumpyRandom, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp006
			πTemp004 = πg.NewDict()
			if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Random", "build/src/__python__/random.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 71: """Random number generator base class used by bound module functions.
					πF.SetLineno(71)
					// line 87: VERSION = 3     # used by getstate/setstate
					πF.SetLineno(87)
					if πE = πClass.SetItem(πF, ßVERSION.ToObject(), πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					// line 89: def __init__(self, x=None):
					πF.SetLineno(89)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "x", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/random.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µx *πg.Object = πArgs[1]; _ = µx
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
							// line 90: """Initialize an instance.
							πF.SetLineno(90)
							// line 95: self.seed(x)
							πF.SetLineno(95)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßseed, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 96: self.gauss_next = None
							πF.SetLineno(96)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßgauss_next, πTemp003); πE != nil {
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
					// line 98: def seed(self, a=None):
					πF.SetLineno(98)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "a", Def: πTemp004}
					πTemp003 = πg.NewFunction(πg.NewCode("seed", "build/src/__python__/random.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πArgs[1]; _ = µa
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
							// line 99: """Initialize internal state of the random number generator.
							πF.SetLineno(99)
							// line 109: super(Random, self).seed(a)
							πF.SetLineno(109)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							πTemp002 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßRandom); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[1] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßseed, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 110: self.gauss_next = None
							πF.SetLineno(110)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßgauss_next, πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ßseed.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 117: def randrange(self, start, stop=None, step=1, _int=int, _maxwidth=1L<<BPF):
					πF.SetLineno(117)
					πTemp002 = make([]πg.Param, 6)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "start", Def: nil}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "stop", Def: πTemp005}
					πTemp002[3] = πg.Param{Name: "step", Def: πg.NewInt(1).ToObject()}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßint); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "_int", Def: πTemp005}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßBPF); πE != nil {
						continue
					}
					if πTemp005, πE = πg.LShift(πF, πg.NewLongFromBytes([]byte{0x1,}).ToObject(), πTemp006); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "_maxwidth", Def: πTemp005}
					πTemp004 = πg.NewFunction(πg.NewCode("randrange", "build/src/__python__/random.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstart *πg.Object = πArgs[1]; _ = µstart
						var µstop *πg.Object = πArgs[2]; _ = µstop
						var µstep *πg.Object = πArgs[3]; _ = µstep
						var µ_int *πg.Object = πArgs[4]; _ = µ_int
						var µ_maxwidth *πg.Object = πArgs[5]; _ = µ_maxwidth
						var µistart *πg.Object = πg.UnboundLocal; _ = µistart
						var µistop *πg.Object = πg.UnboundLocal; _ = µistop
						var µwidth *πg.Object = πg.UnboundLocal; _ = µwidth
						var µistep *πg.Object = πg.UnboundLocal; _ = µistep
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
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
							default: panic("unexpected function state")
							}
							// line 118: """Choose a random item from range(start, stop[, step]).
							πF.SetLineno(118)
							// line 127: istart = _int(start)
							πF.SetLineno(127)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							πTemp001[0] = µstart
							if πE = πg.CheckLocal(πF, µ_int, "_int"); πE != nil {
								continue
							}
							if πTemp002, πE = µ_int.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µistart = πTemp002
							if πE = πg.CheckLocal(πF, µistart, "istart"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.NE(πF, µistart, µstart); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 128: if istart != start:
							πF.SetLineno(128)
						Label1:
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							// line 129: raise ValueError, "non-integer arg 1 for randrange()"
							πF.SetLineno(129)
							πE = πF.Raise(πTemp002, πg.NewStr("non-integer arg 1 for randrange()").ToObject(), nil)
							continue
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µstop, "stop"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µstop == πTemp004).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 130: if stop is None:
							πF.SetLineno(130)
						Label3:
							if πE = πg.CheckLocal(πF, µistart, "istart"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GT(πF, µistart, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							goto Label6
							// line 131: if istart > 0:
							πF.SetLineno(131)
						Label5:
							if πE = πg.CheckLocal(πF, µistart, "istart"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µ_maxwidth, "_maxwidth"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GE(πF, µistart, µ_maxwidth); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label7
							}
							goto Label8
							// line 132: if istart >= _maxwidth:
							πF.SetLineno(132)
						Label7:
							// line 133: return self._randbelow(istart)
							πF.SetLineno(133)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µistart, "istart"); πE != nil {
								continue
							}
							πTemp001[0] = µistart
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_randbelow, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp004
							continue
							goto Label8
						Label8:
							// line 134: return _int(self.random() * istart)
							πF.SetLineno(134)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßrandom, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µistart, "istart"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mul(πF, πTemp005, µistart); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µ_int, "_int"); πE != nil {
								continue
							}
							if πTemp002, πE = µ_int.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp002
							continue
							goto Label6
						Label6:
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							// line 135: raise ValueError, "empty range for randrange()"
							πF.SetLineno(135)
							πE = πF.Raise(πTemp002, πg.NewStr("empty range for randrange()").ToObject(), nil)
							continue
							goto Label4
						Label4:
							// line 138: istop = _int(stop)
							πF.SetLineno(138)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstop, "stop"); πE != nil {
								continue
							}
							πTemp001[0] = µstop
							if πE = πg.CheckLocal(πF, µ_int, "_int"); πE != nil {
								continue
							}
							if πTemp002, πE = µ_int.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µistop = πTemp002
							if πE = πg.CheckLocal(πF, µistop, "istop"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstop, "stop"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.NE(πF, µistop, µstop); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label9
							}
							goto Label10
							// line 139: if istop != stop:
							πF.SetLineno(139)
						Label9:
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							// line 140: raise ValueError, "non-integer stop for randrange()"
							πF.SetLineno(140)
							πE = πF.Raise(πTemp002, πg.NewStr("non-integer stop for randrange()").ToObject(), nil)
							continue
							goto Label10
						Label10:
							// line 141: width = istop - istart
							πF.SetLineno(141)
							if πE = πg.CheckLocal(πF, µistop, "istop"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µistart, "istart"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, µistop, µistart); πE != nil {
								continue
							}
							µwidth = πTemp002
							if πE = πg.CheckLocal(πF, µstep, "step"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Eq(πF, µstep, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label11
							}
							if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GT(πF, µwidth, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004
						Label11:
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label12
							}
							goto Label13
							// line 142: if step == 1 and width > 0:
							πF.SetLineno(142)
						Label12:
							if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µ_maxwidth, "_maxwidth"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GE(πF, µwidth, µ_maxwidth); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label14
							}
							goto Label15
							// line 156: if width >= _maxwidth:
							πF.SetLineno(156)
						Label14:
							// line 157: return _int(istart + self._randbelow(width))
							πF.SetLineno(157)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µistart, "istart"); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
								continue
							}
							πTemp006[0] = µwidth
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_randbelow, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp002, πE = πg.Add(πF, µistart, πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µ_int, "_int"); πE != nil {
								continue
							}
							if πTemp002, πE = µ_int.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp002
							continue
							goto Label15
						Label15:
							// line 158: return _int(istart + _int(self.random()*width))
							πF.SetLineno(158)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µistart, "istart"); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßrandom, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mul(πF, πTemp007, µwidth); πE != nil {
								continue
							}
							πTemp006[0] = πTemp004
							if πE = πg.CheckLocal(πF, µ_int, "_int"); πE != nil {
								continue
							}
							if πTemp004, πE = µ_int.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp002, πE = πg.Add(πF, µistart, πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µ_int, "_int"); πE != nil {
								continue
							}
							if πTemp002, πE = µ_int.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp002
							continue
							goto Label13
						Label13:
							if πE = πg.CheckLocal(πF, µstep, "step"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µstep, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label16
							}
							goto Label17
							// line 159: if step == 1:
							πF.SetLineno(159)
						Label16:
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µistart, "istart"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µistop, "istop"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple3(µistart, µistop, µwidth).ToObject()
							if πTemp004, πE = πg.Mod(πF, πg.NewStr("empty range for randrange() (%d,%d, %d)").ToObject(), πTemp005); πE != nil {
								continue
							}
							// line 160: raise ValueError, "empty range for randrange() (%d,%d, %d)" % (istart, istop, width)
							πF.SetLineno(160)
							πE = πF.Raise(πTemp002, πTemp004, nil)
							continue
							goto Label17
						Label17:
							// line 163: istep = _int(step)
							πF.SetLineno(163)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstep, "step"); πE != nil {
								continue
							}
							πTemp001[0] = µstep
							if πE = πg.CheckLocal(πF, µ_int, "_int"); πE != nil {
								continue
							}
							if πTemp002, πE = µ_int.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µistep = πTemp002
							if πE = πg.CheckLocal(πF, µistep, "istep"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstep, "step"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.NE(πF, µistep, µstep); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label18
							}
							goto Label19
							// line 164: if istep != step:
							πF.SetLineno(164)
						Label18:
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							// line 165: raise ValueError, "non-integer step for randrange()"
							πF.SetLineno(165)
							πE = πF.Raise(πTemp002, πg.NewStr("non-integer step for randrange()").ToObject(), nil)
							continue
							goto Label19
						Label19:
							if πE = πg.CheckLocal(πF, µistep, "istep"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GT(πF, µistep, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label20
							}
							if πE = πg.CheckLocal(πF, µistep, "istep"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, µistep, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label21
							}
							goto Label22
							// line 166: if istep > 0:
							πF.SetLineno(166)
						Label20:
							// line 167: n = (width + istep - 1) // istep
							πF.SetLineno(167)
							if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µistep, "istep"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Add(πF, µwidth, µistep); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Sub(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µistep, "istep"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.FloorDiv(πF, πTemp004, µistep); πE != nil {
								continue
							}
							µn = πTemp002
							goto Label23
							// line 168: elif istep < 0:
							πF.SetLineno(168)
						Label21:
							// line 169: n = (width + istep + 1) // istep
							πF.SetLineno(169)
							if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µistep, "istep"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Add(πF, µwidth, µistep); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µistep, "istep"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.FloorDiv(πF, πTemp004, µistep); πE != nil {
								continue
							}
							µn = πTemp002
							goto Label23
						Label22:
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							// line 171: raise ValueError, "zero step for randrange()"
							πF.SetLineno(171)
							πE = πF.Raise(πTemp002, πg.NewStr("zero step for randrange()").ToObject(), nil)
							continue
							goto Label23
						Label23:
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LE(πF, µn, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label24
							}
							goto Label25
							// line 173: if n <= 0:
							πF.SetLineno(173)
						Label24:
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							// line 174: raise ValueError, "empty range for randrange()"
							πF.SetLineno(174)
							πE = πF.Raise(πTemp002, πg.NewStr("empty range for randrange()").ToObject(), nil)
							continue
							goto Label25
						Label25:
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µ_maxwidth, "_maxwidth"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GE(πF, µn, µ_maxwidth); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label26
							}
							goto Label27
							// line 176: if n >= _maxwidth:
							πF.SetLineno(176)
						Label26:
							// line 177: return istart + istep*self._randbelow(n)
							πF.SetLineno(177)
							if πE = πg.CheckLocal(πF, µistart, "istart"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µistep, "istep"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[0] = µn
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_randbelow, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.Mul(πF, µistep, πTemp007); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µistart, πTemp004); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label27
						Label27:
							// line 178: return istart + istep*_int(self.random() * n)
							πF.SetLineno(178)
							if πE = πg.CheckLocal(πF, µistart, "istart"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µistep, "istep"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßrandom, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Mul(πF, πTemp008, µn); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µ_int, "_int"); πE != nil {
								continue
							}
							if πTemp005, πE = µ_int.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.Mul(πF, µistep, πTemp005); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µistart, πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ßrandrange.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 180: def randint(self, a, b):
					πF.SetLineno(180)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "a", Def: nil}
					πTemp002[2] = πg.Param{Name: "b", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("randint", "build/src/__python__/random.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πArgs[1]; _ = µa
						var µb *πg.Object = πArgs[2]; _ = µb
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
							// line 181: """Return random integer in range [a, b], including both end points.
							πF.SetLineno(181)
							// line 184: return self.randrange(a, b+1)
							πF.SetLineno(184)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[0] = µa
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µb, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrandrange, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßrandint.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 188: def choice(self, seq):
					πF.SetLineno(188)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "seq", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("choice", "build/src/__python__/random.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µseq *πg.Object = πArgs[1]; _ = µseq
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
						var πTemp006 []*πg.Object
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
							// line 189: """Choose a random element from a non-empty sequence."""
							πF.SetLineno(189)
							// line 190: return seq[int(self.random() * len(seq))]  # raises IndexError if seq is empty
							πF.SetLineno(190)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßrandom, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µseq, "seq"); πE != nil {
								continue
							}
							πTemp006[0] = µseq
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp003, πE = πg.Mul(πF, πTemp005, πTemp007); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πTemp004
							if πE = πg.CheckLocal(πF, µseq, "seq"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µseq, πTemp001); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßchoice.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 192: def shuffle(self, x, random=None):
					πF.SetLineno(192)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "x", Def: nil}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "random", Def: πTemp008}
					πTemp007 = πg.NewFunction(πg.NewCode("shuffle", "build/src/__python__/random.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µx *πg.Object = πArgs[1]; _ = µx
						var µrandom *πg.Object = πArgs[2]; _ = µrandom
						var µ_int *πg.Object = πg.UnboundLocal; _ = µ_int
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µj *πg.Object = πg.UnboundLocal; _ = µj
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 3: goto Label3
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							// line 193: """x, random=random.random -> shuffle list x in place; return None.
							πF.SetLineno(193)
							if πE = πg.CheckLocal(πF, µrandom, "random"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µrandom == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 200: if random is None:
							πF.SetLineno(200)
						Label1:
							// line 201: random = self.random
							πF.SetLineno(201)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßrandom, nil); πE != nil {
								continue
							}
							µrandom = πTemp001
							goto Label2
						Label2:
							// line 202: _int = int
							πF.SetLineno(202)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							µ_int = πTemp001
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp006[0] = µx
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[1] = πTemp007
							if πTemp002, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[0] = πTemp007
							if πTemp002, πE = πg.ResolveGlobal(πF, ßreversed); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.Iter(πF, πTemp007); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp003 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
								πF.PopCheckpoint()
								goto Label5
							}
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µi = πTemp002
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 205: j = _int(random() * (i+1))
							πF.SetLineno(205)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrandom, "random"); πE != nil {
								continue
							}
							if πTemp007, πE = µrandom.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mul(πF, πTemp007, πTemp009); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µ_int, "_int"); πE != nil {
								continue
							}
							if πTemp002, πE = µ_int.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µj = πTemp002
							// line 206: x[i], x[j] = x[j], x[i]
							πF.SetLineno(206)
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							πTemp007 = µj
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µx, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp007 = µi
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µx, πTemp007); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp009, πTemp010).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp009}}}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp010 = µi
							if πE = πg.SetItem(πF, µx, πTemp010, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							πTemp007 = µj
							if πE = πg.SetItem(πF, µx, πTemp007, πTemp009); πE != nil {
								continue
							}
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßshuffle.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 273: def uniform(self, a, b):
					πF.SetLineno(273)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "a", Def: nil}
					πTemp002[2] = πg.Param{Name: "b", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("uniform", "build/src/__python__/random.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πArgs[1]; _ = µa
						var µb *πg.Object = πArgs[2]; _ = µb
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
							// line 274: "Get a random number in the range [a, b) or [a, b] depending on rounding."
							πF.SetLineno(274)
							// line 275: return a + (b-a) * self.random()
							πF.SetLineno(275)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, µb, µa); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßrandom, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mul(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µa, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßuniform.ToObject(), πTemp008); πE != nil {
						continue
					}
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Random").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßRandom.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 570: def _test_generator(n, func, args):
			πF.SetLineno(570)
			πTemp007 = make([]πg.Param, 3)
			πTemp007[0] = πg.Param{Name: "n", Def: nil}
			πTemp007[1] = πg.Param{Name: "func", Def: nil}
			πTemp007[2] = πg.Param{Name: "args", Def: nil}
			πTemp002 = πg.NewFunction(πg.NewCode("_test_generator", "build/src/__python__/random.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µn *πg.Object = πArgs[0]; _ = µn
				var µfunc *πg.Object = πArgs[1]; _ = µfunc
				var µargs *πg.Object = πArgs[2]; _ = µargs
				var µtime *πg.Object = πg.UnboundLocal; _ = µtime
				var µtotal *πg.Object = πg.UnboundLocal; _ = µtotal
				var µsqsum *πg.Object = πg.UnboundLocal; _ = µsqsum
				var µsmallest *πg.Object = πg.UnboundLocal; _ = µsmallest
				var µlargest *πg.Object = πg.UnboundLocal; _ = µlargest
				var µt0 *πg.Object = πg.UnboundLocal; _ = µt0
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µx *πg.Object = πg.UnboundLocal; _ = µx
				var µt1 *πg.Object = πg.UnboundLocal; _ = µt1
				var µavg *πg.Object = πg.UnboundLocal; _ = µavg
				var µstddev *πg.Object = πg.UnboundLocal; _ = µstddev
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
				var πTemp007 []*πg.Object
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
					// line 571: import time
					πF.SetLineno(571)
					if πTemp002, πE = πg.ImportModule(πF, "time"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					µtime = πTemp001
					// line 572: print n, 'times', func.__name__
					πF.SetLineno(572)
					πTemp002 = make([]*πg.Object, 3)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp002[0] = µn
					πTemp002[1] = ßtimes.ToObject()
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfunc, ß__name__, nil); πE != nil {
						continue
					}
					πTemp002[2] = πTemp001
					if πE = πg.Print(πF, πTemp002, true); πE != nil {
						continue
					}
					// line 573: total = 0.0
					πF.SetLineno(573)
					µtotal = πg.NewFloat(0.0).ToObject()
					// line 574: sqsum = 0.0
					πF.SetLineno(574)
					µsqsum = πg.NewFloat(0.0).ToObject()
					// line 575: smallest = 1e10
					πF.SetLineno(575)
					µsmallest = πg.NewFloat(10000000000.0).ToObject()
					// line 576: largest = -1e10
					πF.SetLineno(576)
					if πTemp001, πE = πg.Neg(πF, πg.NewFloat(10000000000.0).ToObject()); πE != nil {
						continue
					}
					µlargest = πTemp001
					// line 577: t0 = time.time()
					πF.SetLineno(577)
					if πE = πg.CheckLocal(πF, µtime, "time"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µtime, ßtime, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µt0 = πTemp003
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp002[0] = µn
					if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
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
						µi = πTemp003
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 579: x = func(*args)
					πF.SetLineno(579)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, µfunc, nil, µargs, nil, nil); πE != nil {
						continue
					}
					µx = πTemp003
					// line 580: total += x
					πF.SetLineno(580)
					if πE = πg.CheckLocal(πF, µtotal, "total"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IAdd(πF, µtotal, µx); πE != nil {
						continue
					}
					µtotal = πTemp003
					// line 581: sqsum = sqsum + x*x
					πF.SetLineno(581)
					if πE = πg.CheckLocal(πF, µsqsum, "sqsum"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mul(πF, µx, µx); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µsqsum, πTemp004); πE != nil {
						continue
					}
					µsqsum = πTemp003
					// line 582: smallest = min(x, smallest)
					πF.SetLineno(582)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πE = πg.CheckLocal(πF, µsmallest, "smallest"); πE != nil {
						continue
					}
					πTemp002[1] = µsmallest
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µsmallest = πTemp004
					// line 583: largest = max(x, largest)
					πF.SetLineno(583)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πE = πg.CheckLocal(πF, µlargest, "largest"); πE != nil {
						continue
					}
					πTemp002[1] = µlargest
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µlargest = πTemp004
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 584: t1 = time.time()
					πF.SetLineno(584)
					if πE = πg.CheckLocal(πF, µtime, "time"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µtime, ßtime, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µt1 = πTemp003
					// line 585: print round(t1-t0, 3), 'sec,',
					πF.SetLineno(585)
					πTemp002 = make([]*πg.Object, 2)
					πTemp007 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µt1, "t1"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µt0, "t0"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, µt1, µt0); πE != nil {
						continue
					}
					πTemp007[0] = πTemp001
					πTemp007[1] = πg.NewInt(3).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßround); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp002[0] = πTemp003
					πTemp002[1] = πg.NewStr("sec,").ToObject()
					if πE = πg.Print(πF, πTemp002, false); πE != nil {
						continue
					}
					// line 586: avg = total/n
					πF.SetLineno(586)
					if πE = πg.CheckLocal(πF, µtotal, "total"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Div(πF, µtotal, µn); πE != nil {
						continue
					}
					µavg = πTemp001
					// line 587: stddev = _sqrt(sqsum/n - avg*avg)
					πF.SetLineno(587)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsqsum, "sqsum"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Div(πF, µsqsum, µn); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µavg, "avg"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µavg, "avg"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mul(πF, µavg, µavg); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_sqrt); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µstddev = πTemp003
					// line 588: print 'avg %g, stddev %g, min %g, max %g' % \
					πF.SetLineno(588)
					πTemp002 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µavg, "avg"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstddev, "stddev"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsmallest, "smallest"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlargest, "largest"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple4(µavg, µstddev, µsmallest, µlargest).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("avg %g, stddev %g, min %g, max %g").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πE = πg.Print(πF, πTemp002, true); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_test_generator.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 592: def _test(N=2000):
			πF.SetLineno(592)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "N", Def: πg.NewInt(2000).ToObject()}
			πTemp003 = πg.NewFunction(πg.NewCode("_test", "build/src/__python__/random.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µN *πg.Object = πArgs[0]; _ = µN
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
					// line 593: _test_generator(N, random, ())
					πF.SetLineno(593)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
						continue
					}
					πTemp001[0] = µN
					if πTemp002, πE = πg.ResolveGlobal(πF, ßrandom); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					πTemp002 = πg.NewTuple0().ToObject()
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_test_generator); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 594: _test_generator(N, normalvariate, (0.0, 1.0))
					πF.SetLineno(594)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
						continue
					}
					πTemp001[0] = µN
					if πTemp002, πE = πg.ResolveGlobal(πF, ßnormalvariate); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					πTemp002 = πg.NewTuple2(πg.NewFloat(0.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_test_generator); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 595: _test_generator(N, lognormvariate, (0.0, 1.0))
					πF.SetLineno(595)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
						continue
					}
					πTemp001[0] = µN
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlognormvariate); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					πTemp002 = πg.NewTuple2(πg.NewFloat(0.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_test_generator); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 596: _test_generator(N, vonmisesvariate, (0.0, 1.0))
					πF.SetLineno(596)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
						continue
					}
					πTemp001[0] = µN
					if πTemp002, πE = πg.ResolveGlobal(πF, ßvonmisesvariate); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					πTemp002 = πg.NewTuple2(πg.NewFloat(0.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_test_generator); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 597: _test_generator(N, gammavariate, (0.01, 1.0))
					πF.SetLineno(597)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
						continue
					}
					πTemp001[0] = µN
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgammavariate); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					πTemp002 = πg.NewTuple2(πg.NewFloat(0.01).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_test_generator); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 598: _test_generator(N, gammavariate, (0.1, 1.0))
					πF.SetLineno(598)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
						continue
					}
					πTemp001[0] = µN
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgammavariate); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					πTemp002 = πg.NewTuple2(πg.NewFloat(0.1).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_test_generator); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 599: _test_generator(N, gammavariate, (0.1, 2.0))
					πF.SetLineno(599)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
						continue
					}
					πTemp001[0] = µN
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgammavariate); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					πTemp002 = πg.NewTuple2(πg.NewFloat(0.1).ToObject(), πg.NewFloat(2.0).ToObject()).ToObject()
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_test_generator); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 600: _test_generator(N, gammavariate, (0.5, 1.0))
					πF.SetLineno(600)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
						continue
					}
					πTemp001[0] = µN
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgammavariate); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					πTemp002 = πg.NewTuple2(πg.NewFloat(0.5).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_test_generator); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 601: _test_generator(N, gammavariate, (0.9, 1.0))
					πF.SetLineno(601)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
						continue
					}
					πTemp001[0] = µN
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgammavariate); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					πTemp002 = πg.NewTuple2(πg.NewFloat(0.9).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_test_generator); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 602: _test_generator(N, gammavariate, (1.0, 1.0))
					πF.SetLineno(602)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
						continue
					}
					πTemp001[0] = µN
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgammavariate); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					πTemp002 = πg.NewTuple2(πg.NewFloat(1.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_test_generator); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 603: _test_generator(N, gammavariate, (2.0, 1.0))
					πF.SetLineno(603)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
						continue
					}
					πTemp001[0] = µN
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgammavariate); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					πTemp002 = πg.NewTuple2(πg.NewFloat(2.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_test_generator); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 604: _test_generator(N, gammavariate, (20.0, 1.0))
					πF.SetLineno(604)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
						continue
					}
					πTemp001[0] = µN
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgammavariate); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					πTemp002 = πg.NewTuple2(πg.NewFloat(20.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_test_generator); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 605: _test_generator(N, gammavariate, (200.0, 1.0))
					πF.SetLineno(605)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
						continue
					}
					πTemp001[0] = µN
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgammavariate); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					πTemp002 = πg.NewTuple2(πg.NewFloat(200.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_test_generator); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 606: _test_generator(N, gauss, (0.0, 1.0))
					πF.SetLineno(606)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
						continue
					}
					πTemp001[0] = µN
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgauss); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					πTemp002 = πg.NewTuple2(πg.NewFloat(0.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_test_generator); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 607: _test_generator(N, betavariate, (3.0, 3.0))
					πF.SetLineno(607)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
						continue
					}
					πTemp001[0] = µN
					if πTemp002, πE = πg.ResolveGlobal(πF, ßbetavariate); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					πTemp002 = πg.NewTuple2(πg.NewFloat(3.0).ToObject(), πg.NewFloat(3.0).ToObject()).ToObject()
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_test_generator); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 608: _test_generator(N, triangular, (0.0, 1.0, 1.0/3.0))
					πF.SetLineno(608)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
						continue
					}
					πTemp001[0] = µN
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtriangular); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp003, πE = πg.Div(πF, πg.NewFloat(1.0).ToObject(), πg.NewFloat(3.0).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(πg.NewFloat(0.0).ToObject(), πg.NewFloat(1.0).ToObject(), πTemp003).ToObject()
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_test_generator); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_test.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 616: _inst = Random()
			πF.SetLineno(616)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßRandom); πE != nil {
				continue
			}
			if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_inst.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 617: seed = _inst.seed
			πF.SetLineno(617)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_inst); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßseed, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßseed.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 618: random = _inst.random
			πF.SetLineno(618)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_inst); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßrandom, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßrandom.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 619: randint = _inst.randint
			πF.SetLineno(619)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_inst); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßrandint, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßrandint.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 620: choice = _inst.choice
			πF.SetLineno(620)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_inst); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßchoice, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßchoice.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 621: randrange = _inst.randrange
			πF.SetLineno(621)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_inst); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßrandrange, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßrandrange.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 622: getrandbits = _inst.getrandbits
			πF.SetLineno(622)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_inst); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßgetrandbits, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßgetrandbits.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 623: getstate = _inst.getstate
			πF.SetLineno(623)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_inst); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßgetstate, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßgetstate.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 624: setstate = _inst.setstate
			πF.SetLineno(624)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_inst); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßsetstate, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßsetstate.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 625: uniform = _inst.uniform
			πF.SetLineno(625)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_inst); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßuniform, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßuniform.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 628: def _notimplemented(*args, **kwargs):
			πF.SetLineno(628)
			πTemp007 = make([]πg.Param, 0)
			πTemp005 = πg.NewFunction(πg.NewCode("_notimplemented", "build/src/__python__/random.py", πTemp007, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargs *πg.Object = πArgs[0]; _ = µargs
				var µkwargs *πg.Object = πArgs[1]; _ = µkwargs
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
						continue
					}
					// line 629: raise NotImplementedError
					πF.SetLineno(629)
					πE = πF.Raise(πTemp001, nil, nil)
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_notimplemented.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 632: shuffle = _notimplemented
			πF.SetLineno(632)
			if πTemp006, πE = πg.ResolveGlobal(πF, ß_notimplemented); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßshuffle.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 633: choices = _notimplemented
			πF.SetLineno(633)
			if πTemp006, πE = πg.ResolveGlobal(πF, ß_notimplemented); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßchoices.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 634: sample = _notimplemented
			πF.SetLineno(634)
			if πTemp006, πE = πg.ResolveGlobal(πF, ß_notimplemented); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßsample.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 635: triangular = _notimplemented
			πF.SetLineno(635)
			if πTemp006, πE = πg.ResolveGlobal(πF, ß_notimplemented); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßtriangular.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 636: normalvariate = _notimplemented
			πF.SetLineno(636)
			if πTemp006, πE = πg.ResolveGlobal(πF, ß_notimplemented); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßnormalvariate.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 637: lognormvariate = _notimplemented
			πF.SetLineno(637)
			if πTemp006, πE = πg.ResolveGlobal(πF, ß_notimplemented); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßlognormvariate.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 638: expovariate = _notimplemented
			πF.SetLineno(638)
			if πTemp006, πE = πg.ResolveGlobal(πF, ß_notimplemented); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßexpovariate.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 639: vonmisesvariate = _notimplemented
			πF.SetLineno(639)
			if πTemp006, πE = πg.ResolveGlobal(πF, ß_notimplemented); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßvonmisesvariate.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 640: gammavariate = _notimplemented
			πF.SetLineno(640)
			if πTemp006, πE = πg.ResolveGlobal(πF, ß_notimplemented); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßgammavariate.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 641: gauss = _notimplemented
			πF.SetLineno(641)
			if πTemp006, πE = πg.ResolveGlobal(πF, ß_notimplemented); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßgauss.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 642: betavariate = _notimplemented
			πF.SetLineno(642)
			if πTemp006, πE = πg.ResolveGlobal(πF, ß_notimplemented); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßbetavariate.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 643: paretovariate = _notimplemented
			πF.SetLineno(643)
			if πTemp006, πE = πg.ResolveGlobal(πF, ß_notimplemented); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßparetovariate.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 644: weibullvariate = _notimplemented
			πF.SetLineno(644)
			if πTemp006, πE = πg.ResolveGlobal(πF, ß_notimplemented); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßweibullvariate.ToObject(), πTemp006); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp006, πE = πg.Eq(πF, πTemp008, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp009, πE = πg.IsTrue(πF, πTemp006); πE != nil {
				continue
			}
			if πTemp009 {
				goto Label1
			}
			goto Label2
			// line 647: if __name__ == '__main__':
			πF.SetLineno(647)
		Label1:
			// line 648: pass
			πF.SetLineno(648)
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("random", Code)
}
