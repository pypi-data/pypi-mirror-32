package _random
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/_random.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßBPF := πg.InternStr("BPF")
		ßGrumpyRandom := πg.InternStr("GrumpyRandom")
		ßNone := πg.InternStr("None")
		ßNotImplementedError := πg.InternStr("NotImplementedError")
		ßNow := πg.InternStr("Now")
		ßPow := πg.InternStr("Pow")
		ßRECIP_BPF := πg.InternStr("RECIP_BPF")
		ßSeed := πg.InternStr("Seed")
		ßTypeError := πg.InternStr("TypeError")
		ßUint32 := πg.InternStr("Uint32")
		ßUnixNano := πg.InternStr("UnixNano")
		ßValueError := πg.InternStr("ValueError")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_gorandom := πg.InternStr("_gorandom")
		ß_int_bit_length := πg.InternStr("_int_bit_length")
		ß_int_from_bytes := πg.InternStr("_int_from_bytes")
		ß_randbelow := πg.InternStr("_randbelow")
		ßappend := πg.InternStr("append")
		ßgetrandbits := πg.InternStr("getrandbits")
		ßgetstate := πg.InternStr("getstate")
		ßint := πg.InternStr("int")
		ßjumpahead := πg.InternStr("jumpahead")
		ßlen := πg.InternStr("len")
		ßobject := πg.InternStr("object")
		ßrandom := πg.InternStr("random")
		ßseed := πg.InternStr("seed")
		ßsetstate := πg.InternStr("setstate")
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
		var πTemp006 *πg.Dict
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
			// line 15: """Generate pseudo random numbers. Should not be used for security purposes."""
			πF.SetLineno(15)
			// line 17: from '__go__/math/rand' import Uint32, Seed
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/math/rand"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßUint32, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßUint32.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßSeed, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSeed.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 18: from '__go__/math' import Pow
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/math"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßPow, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPow.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 19: from '__go__/time' import Now
			πF.SetLineno(19)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/time"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßNow, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNow.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 22: BPF = 53  # Number of bits in a float
			πF.SetLineno(22)
			if πE = πF.Globals().SetItem(πF, ßBPF.ToObject(), πg.NewInt(53).ToObject()); πE != nil {
				continue
			}
			// line 23: RECIP_BPF = Pow(2, -BPF)
			πF.SetLineno(23)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewInt(2).ToObject()
			if πTemp003, πE = πg.ResolveGlobal(πF, ßBPF); πE != nil {
				continue
			}
			if πTemp001, πE = πg.Neg(πF, πTemp003); πE != nil {
				continue
			}
			πTemp002[1] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßPow); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßRECIP_BPF.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 31: def _gorandom(nbytes):
			πF.SetLineno(31)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "nbytes", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("_gorandom", "build/src/__python__/_random.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µnbytes *πg.Object = πArgs[0]; _ = µnbytes
				var µbyte_arr *πg.Object = πg.UnboundLocal; _ = µbyte_arr
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
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
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 32: byte_arr = []
					πF.SetLineno(32)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µbyte_arr = πTemp002
					// line 33: while len(byte_arr) < nbytes:
					πF.SetLineno(33)
					πF.PushCheckpoint(2)
					πTemp003 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp003 {
						πF.PopCheckpoint()
						goto Label3
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µbyte_arr, "byte_arr"); πE != nil {
						continue
					}
					πTemp001[0] = µbyte_arr
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µnbytes, "nbytes"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LT(πF, πTemp006, µnbytes); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 34: i = Uint32()
					πF.SetLineno(34)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßUint32); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µi = πTemp005
					// line 35: byte_arr.append(i & 0xff)
					πF.SetLineno(35)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, µi, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µbyte_arr, "byte_arr"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µbyte_arr, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 36: byte_arr.append(i >> 8 & 0xff)
					πF.SetLineno(36)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.RShift(πF, µi, πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, πTemp005, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µbyte_arr, "byte_arr"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µbyte_arr, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 37: byte_arr.append(i >> 16 & 0xff)
					πF.SetLineno(37)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.RShift(πF, µi, πg.NewInt(16).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, πTemp005, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µbyte_arr, "byte_arr"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µbyte_arr, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 38: byte_arr.append(i >> 24 & 0xff)
					πF.SetLineno(38)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.RShift(πF, µi, πg.NewInt(24).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, πTemp005, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µbyte_arr, "byte_arr"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µbyte_arr, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 39: byte_arr = byte_arr[0:nbytes]
					πF.SetLineno(39)
					if πE = πg.CheckLocal(πF, µnbytes, "nbytes"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(0).ToObject(), µnbytes, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µbyte_arr, "byte_arr"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µbyte_arr, πTemp002); πE != nil {
						continue
					}
					µbyte_arr = πTemp005
					// line 40: return byte_arr
					πF.SetLineno(40)
					if πE = πg.CheckLocal(πF, µbyte_arr, "byte_arr"); πE != nil {
						continue
					}
					πR = µbyte_arr
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_gorandom.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 45: def _int_bit_length(n):
			πF.SetLineno(45)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "n", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("_int_bit_length", "build/src/__python__/_random.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µn *πg.Object = πArgs[0]; _ = µn
				var µbits *πg.Object = πg.UnboundLocal; _ = µbits
				var πTemp001 bool
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
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 46: bits = 0
					πF.SetLineno(46)
					µbits = πg.NewInt(0).ToObject()
					// line 47: while n:  # 1 bit steps
					πF.SetLineno(47)
					πF.PushCheckpoint(2)
					πTemp001 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp001 {
						πF.PopCheckpoint()
						goto Label3
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µn); πE != nil {
						continue
					}
					if πE != nil || !πTemp002 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 48: n = n / 2
					πF.SetLineno(48)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Div(πF, µn, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					µn = πTemp003
					// line 49: bits += 1
					πF.SetLineno(49)
					if πE = πg.CheckLocal(πF, µbits, "bits"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IAdd(πF, µbits, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µbits = πTemp003
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 50: return bits
					πF.SetLineno(50)
					if πE = πg.CheckLocal(πF, µbits, "bits"); πE != nil {
						continue
					}
					πR = µbits
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_int_bit_length.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 54: def _int_from_bytes(bytes):
			πF.SetLineno(54)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "bytes", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("_int_from_bytes", "build/src/__python__/_random.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µbytes *πg.Object = πArgs[0]; _ = µbytes
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µn *πg.Object = πg.UnboundLocal; _ = µn
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 55: i = 0
					πF.SetLineno(55)
					µi = πg.NewInt(0).ToObject()
					// line 56: n = len(bytes) - 1
					πF.SetLineno(56)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µbytes, "bytes"); πE != nil {
						continue
					}
					πTemp002[0] = µbytes
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Sub(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µn = πTemp001
					// line 57: while n >= 0:
					πF.SetLineno(57)
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
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GE(πF, µn, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 58: i += bytes[n] << (8 * n)
					πF.SetLineno(58)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp003 = µn
					if πE = πg.CheckLocal(πF, µbytes, "bytes"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µbytes, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, πg.NewInt(8).ToObject(), µn); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LShift(πF, πTemp004, πTemp003); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IAdd(πF, µi, πTemp001); πE != nil {
						continue
					}
					µi = πTemp003
					// line 59: n -= 1
					πF.SetLineno(59)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ISub(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µn = πTemp001
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 60: return i
					πF.SetLineno(60)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πR = µi
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_int_from_bytes.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 63: class GrumpyRandom(object):
			πF.SetLineno(63)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp007, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp007); πE != nil {
				continue
			}
			_, πE = πg.NewCode("GrumpyRandom", "build/src/__python__/_random.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
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
					// line 64: """Random generator replacement for Grumpy.
					πF.SetLineno(64)
					// line 70: def random(self):
					πF.SetLineno(70)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("random", "build/src/__python__/_random.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 71: """Get the next random number in the range [0.0, 1.0)."""
							πF.SetLineno(71)
							// line 72: return (_int_from_bytes(_gorandom(7)) >> 3) * RECIP_BPF
							πF.SetLineno(72)
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewInt(7).ToObject()
							if πTemp005, πE = πg.ResolveGlobal(πF, ß_gorandom); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp006
							if πTemp005, πE = πg.ResolveGlobal(πF, ß_int_from_bytes); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.RShift(πF, πTemp006, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßRECIP_BPF); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, πTemp002, πTemp005); πE != nil {
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
					if πE = πClass.SetItem(πF, ßrandom.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 74: def getrandbits(self, k):
					πF.SetLineno(74)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "k", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("getrandbits", "build/src/__python__/_random.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µk *πg.Object = πArgs[1]; _ = µk
						var µnumbytes *πg.Object = πg.UnboundLocal; _ = µnumbytes
						var µx *πg.Object = πg.UnboundLocal; _ = µx
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 75: """getrandbits(k) -> x.  Generates an int with k random bits."""
							πF.SetLineno(75)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LE(πF, µk, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 76: if k <= 0:
							πF.SetLineno(76)
						Label1:
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("number of bits must be greater than zero").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 77: raise ValueError('number of bits must be greater than zero')
							πF.SetLineno(77)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp003[0] = µk
							if πTemp004, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.NE(πF, µk, πTemp005); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 78: if k != int(k):
							πF.SetLineno(78)
						Label3:
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("number of bits should be an integer").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 79: raise TypeError('number of bits should be an integer')
							πF.SetLineno(79)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label4
						Label4:
							// line 80: numbytes = (k + 7) // 8                       # bits / 8 and rounded up
							πF.SetLineno(80)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, µk, πg.NewInt(7).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.FloorDiv(πF, πTemp004, πg.NewInt(8).ToObject()); πE != nil {
								continue
							}
							µnumbytes = πTemp001
							// line 81: x = _int_from_bytes(_gorandom(numbytes))
							πF.SetLineno(81)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnumbytes, "numbytes"); πE != nil {
								continue
							}
							πTemp006[0] = µnumbytes
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_gorandom); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp004
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_int_from_bytes); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µx = πTemp004
							// line 82: return x >> (numbytes * 8 - k)                # trim excess bits
							πF.SetLineno(82)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnumbytes, "numbytes"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Mul(πF, µnumbytes, πg.NewInt(8).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Sub(πF, πTemp005, µk); πE != nil {
								continue
							}
							if πTemp001, πE = πg.RShift(πF, µx, πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ßgetrandbits.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 84: def seed(self, a=None):
					πF.SetLineno(84)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "a", Def: πTemp005}
					πTemp004 = πg.NewFunction(πg.NewCode("seed", "build/src/__python__/_random.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πArgs[1]; _ = µa
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 85: """Seed the golang.math.rand generator."""
							πF.SetLineno(85)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µa == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 86: if a is None:
							πF.SetLineno(86)
						Label1:
							// line 87: a = Now().UnixNano()
							πF.SetLineno(87)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNow); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßUnixNano, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µa = πTemp002
							goto Label2
						Label2:
							// line 88: Seed(a)
							πF.SetLineno(88)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp004[0] = µa
							if πTemp001, πE = πg.ResolveGlobal(πF, ßSeed); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßseed.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 90: def _randbelow(self, n):
					πF.SetLineno(90)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "n", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("_randbelow", "build/src/__python__/_random.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µn *πg.Object = πArgs[1]; _ = µn
						var µk *πg.Object = πg.UnboundLocal; _ = µk
						var µr *πg.Object = πg.UnboundLocal; _ = µr
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 91: """Return a random int in the range [0,n)."""
							πF.SetLineno(91)
							// line 95: k = _int_bit_length(n)
							πF.SetLineno(95)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[0] = µn
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_int_bit_length); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µk = πTemp003
							// line 96: r = self.getrandbits(k)
							πF.SetLineno(96)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp001[0] = µk
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgetrandbits, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µr = πTemp003
							// line 97: while r >= n:
							πF.SetLineno(97)
							πF.PushCheckpoint(2)
							πTemp004 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GE(πF, µr, µn); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 98: r = self.getrandbits(k)
							πF.SetLineno(98)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp001[0] = µk
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgetrandbits, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µr = πTemp003
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 99: return r
							πF.SetLineno(99)
							if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
								continue
							}
							πR = µr
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_randbelow.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 101: def getstate(self, *args, **kwargs):
					πF.SetLineno(101)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("getstate", "build/src/__python__/_random.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µargs *πg.Object = πArgs[1]; _ = µargs
						var µkwargs *πg.Object = πArgs[2]; _ = µkwargs
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
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("Entropy source does not have state.").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 102: raise NotImplementedError('Entropy source does not have state.')
							πF.SetLineno(102)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßgetstate.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 104: def setstate(self, *args, **kwargs):
					πF.SetLineno(104)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("setstate", "build/src/__python__/_random.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µargs *πg.Object = πArgs[1]; _ = µargs
						var µkwargs *πg.Object = πArgs[2]; _ = µkwargs
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
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("Entropy source does not have state.").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 105: raise NotImplementedError('Entropy source does not have state.')
							πF.SetLineno(105)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßsetstate.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 107: def jumpahead(self, *args, **kwargs):
					πF.SetLineno(107)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("jumpahead", "build/src/__python__/_random.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µargs *πg.Object = πArgs[1]; _ = µargs
						var µkwargs *πg.Object = πArgs[2]; _ = µkwargs
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
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("Entropy source does not have state.").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 108: raise NotImplementedError('Entropy source does not have state.')
							πF.SetLineno(108)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßjumpahead.ToObject(), πTemp008); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("GrumpyRandom").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßGrumpyRandom.ToObject(), πTemp009); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("_random", Code)
}
