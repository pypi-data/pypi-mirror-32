package random_test
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/random_test.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßAssertionError := πg.InternStr("AssertionError")
		ßGrumpyRandom := πg.InternStr("GrumpyRandom")
		ßIndexError := πg.InternStr("IndexError")
		ßRunTests := πg.InternStr("RunTests")
		ßTestGrumpyRandom := πg.InternStr("TestGrumpyRandom")
		ßTestRandom := πg.InternStr("TestRandom")
		ßTestRandomChoice := πg.InternStr("TestRandomChoice")
		ßTestRandomInt := πg.InternStr("TestRandomInt")
		ßTestRandomUniform := πg.InternStr("TestRandomUniform")
		ßTestSeed := πg.InternStr("TestSeed")
		ßTypeError := πg.InternStr("TypeError")
		ßValueError := πg.InternStr("ValueError")
		ß__main__ := πg.InternStr("__main__")
		ß__name__ := πg.InternStr("__name__")
		ß_gorandom := πg.InternStr("_gorandom")
		ß_int_bit_length := πg.InternStr("_int_bit_length")
		ß_int_from_bytes := πg.InternStr("_int_from_bytes")
		ß_randbelow := πg.InternStr("_randbelow")
		ß_random := πg.InternStr("_random")
		ßchoice := πg.InternStr("choice")
		ßfloat := πg.InternStr("float")
		ßgetrandbits := πg.InternStr("getrandbits")
		ßint := πg.InternStr("int")
		ßisinstance := πg.InternStr("isinstance")
		ßlen := πg.InternStr("len")
		ßrandint := πg.InternStr("randint")
		ßrandom := πg.InternStr("random")
		ßrange := πg.InternStr("range")
		ßseed := πg.InternStr("seed")
		ßuniform := πg.InternStr("uniform")
		ßweetest := πg.InternStr("weetest")
		ßwrongtype := πg.InternStr("wrongtype")
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
			// line 15: import _random
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "_random"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_random.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 16: import random
			πF.SetLineno(16)
			if πTemp002, πE = πg.ImportModule(πF, "random"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßrandom.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: import weetest
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "weetest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßweetest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 21: def TestGrumpyRandom():
			πF.SetLineno(21)
			πTemp003 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("TestGrumpyRandom", "build/src/__python__/random_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µr *πg.Object = πg.UnboundLocal; _ = µr
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 22: assert len(_random._gorandom(5)) == 5
					πF.SetLineno(22)
					πTemp002 = πF.MakeArgs(1)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewInt(5).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_random); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß_gorandom, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp002[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp005, πg.NewInt(5).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 24: assert _random._int_bit_length(0) == 0
					πF.SetLineno(24)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(0).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_random); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß_int_bit_length, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 25: assert _random._int_bit_length(1) == 1
					πF.SetLineno(25)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(1).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_random); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß_int_bit_length, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 26: assert _random._int_bit_length(8) == 4
					πF.SetLineno(26)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(8).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_random); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß_int_bit_length, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 27: assert _random._int_bit_length(256) == 9
					πF.SetLineno(27)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(256).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_random); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß_int_bit_length, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewInt(9).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 29: assert _random._int_from_bytes([1, 0, 0, 0]) == 1
					πF.SetLineno(29)
					πTemp002 = πF.MakeArgs(1)
					πTemp003 = make([]*πg.Object, 4)
					πTemp003[0] = πg.NewInt(1).ToObject()
					πTemp003[1] = πg.NewInt(0).ToObject()
					πTemp003[2] = πg.NewInt(0).ToObject()
					πTemp003[3] = πg.NewInt(0).ToObject()
					πTemp004 = πg.NewList(πTemp003...).ToObject()
					πTemp002[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_random); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß_int_from_bytes, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 30: assert _random._int_from_bytes([0, 0, 0, 0]) == 0
					πF.SetLineno(30)
					πTemp002 = πF.MakeArgs(1)
					πTemp003 = make([]*πg.Object, 4)
					πTemp003[0] = πg.NewInt(0).ToObject()
					πTemp003[1] = πg.NewInt(0).ToObject()
					πTemp003[2] = πg.NewInt(0).ToObject()
					πTemp003[3] = πg.NewInt(0).ToObject()
					πTemp004 = πg.NewList(πTemp003...).ToObject()
					πTemp002[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_random); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß_int_from_bytes, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 31: assert _random._int_from_bytes([255, 255, 0, 0]) == 65535
					πF.SetLineno(31)
					πTemp002 = πF.MakeArgs(1)
					πTemp003 = make([]*πg.Object, 4)
					πTemp003[0] = πg.NewInt(255).ToObject()
					πTemp003[1] = πg.NewInt(255).ToObject()
					πTemp003[2] = πg.NewInt(0).ToObject()
					πTemp003[3] = πg.NewInt(0).ToObject()
					πTemp004 = πg.NewList(πTemp003...).ToObject()
					πTemp002[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_random); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß_int_from_bytes, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewInt(65535).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 32: assert _random._int_from_bytes([0, 0, 0, 1]) == 16777216
					πF.SetLineno(32)
					πTemp002 = πF.MakeArgs(1)
					πTemp003 = make([]*πg.Object, 4)
					πTemp003[0] = πg.NewInt(0).ToObject()
					πTemp003[1] = πg.NewInt(0).ToObject()
					πTemp003[2] = πg.NewInt(0).ToObject()
					πTemp003[3] = πg.NewInt(1).ToObject()
					πTemp004 = πg.NewList(πTemp003...).ToObject()
					πTemp002[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_random); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß_int_from_bytes, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewInt(16777216).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 34: r = _random.GrumpyRandom()
					πF.SetLineno(34)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_random); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßGrumpyRandom, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					µr = πTemp001
					// line 35: assert 0.0 <= r.random() < 1.0
					πF.SetLineno(35)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µr, ßrandom, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LE(πF, πg.NewFloat(0.0).ToObject(), πTemp005); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label1
					}
					if πTemp001, πE = πg.LT(πF, πTemp005, πg.NewFloat(1.0).ToObject()); πE != nil {
						continue
					}
				Label1:
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 37: assert 0 <= r.getrandbits(1) <= 1
					πF.SetLineno(37)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µr, ßgetrandbits, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.LE(πF, πg.NewInt(0).ToObject(), πTemp005); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label2
					}
					if πTemp001, πE = πg.LE(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
				Label2:
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 38: assert 0 <= r.getrandbits(2) <= 3
					πF.SetLineno(38)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µr, ßgetrandbits, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.LE(πF, πg.NewInt(0).ToObject(), πTemp005); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label3
					}
					if πTemp001, πE = πg.LE(πF, πTemp005, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
				Label3:
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 39: assert 0 <= r.getrandbits(8) <= 255
					πF.SetLineno(39)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(8).ToObject()
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µr, ßgetrandbits, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.LE(πF, πg.NewInt(0).ToObject(), πTemp005); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label4
					}
					if πTemp001, πE = πg.LE(πF, πTemp005, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
				Label4:
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 41: assert 0 <= r._randbelow(1) < 1
					πF.SetLineno(41)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µr, ß_randbelow, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.LE(πF, πg.NewInt(0).ToObject(), πTemp005); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label5
					}
					if πTemp001, πE = πg.LT(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
				Label5:
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 42: assert 0 <= r._randbelow(3) < 3
					πF.SetLineno(42)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µr, ß_randbelow, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.LE(πF, πg.NewInt(0).ToObject(), πTemp005); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label6
					}
					if πTemp001, πE = πg.LT(πF, πTemp005, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
				Label6:
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 43: assert 0 <= r._randbelow(1000) < 1000
					πF.SetLineno(43)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(1000).ToObject()
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µr, ß_randbelow, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.LE(πF, πg.NewInt(0).ToObject(), πTemp005); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label7
					}
					if πTemp001, πE = πg.LT(πF, πTemp005, πg.NewInt(1000).ToObject()); πE != nil {
						continue
					}
				Label7:
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
			if πE = πF.Globals().SetItem(πF, ßTestGrumpyRandom.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 46: def TestSeed():
			πF.SetLineno(46)
			πTemp003 = make([]πg.Param, 0)
			πTemp004 = πg.NewFunction(πg.NewCode("TestSeed", "build/src/__python__/random_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
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
					default: panic("unexpected function state")
					}
					// line 47: random.seed()
					πF.SetLineno(47)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßrandom); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßseed, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 48: try:
					πF.SetLineno(48)
					πF.PushCheckpoint(2)
					// line 49: random.seed("wrongtype")
					πF.SetLineno(49)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = ßwrongtype.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßrandom); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßseed, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πF.PopCheckpoint()
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("TypeError not raised").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 53: raise AssertionError("TypeError not raised")
					πF.SetLineno(53)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 50: except TypeError:
					πF.SetLineno(50)
				Label3:
					// line 51: pass
					πF.SetLineno(51)
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
			if πE = πF.Globals().SetItem(πF, ßTestSeed.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 56: def TestRandom():
			πF.SetLineno(56)
			πTemp003 = make([]πg.Param, 0)
			πTemp005 = πg.NewFunction(πg.NewCode("TestRandom", "build/src/__python__/random_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πg.UnboundLocal; _ = µa
				var µb *πg.Object = πg.UnboundLocal; _ = µb
				var µc *πg.Object = πg.UnboundLocal; _ = µc
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
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
					// line 57: a = random.random()
					πF.SetLineno(57)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßrandom); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrandom, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µa = πTemp001
					// line 58: b = random.random()
					πF.SetLineno(58)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßrandom); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrandom, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µb = πTemp001
					// line 59: c = random.random()
					πF.SetLineno(59)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßrandom); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrandom, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µc = πTemp001
					// line 60: assert isinstance(a, float)
					πF.SetLineno(60)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp003[0] = µa
					if πTemp001, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					πTemp003[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
						continue
					}
					// line 61: assert 0.0 <= a < 1.0
					πF.SetLineno(61)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LE(πF, πg.NewFloat(0.0).ToObject(), µa); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label1
					}
					if πTemp001, πE = πg.LT(πF, µa, πg.NewFloat(1.0).ToObject()); πE != nil {
						continue
					}
				Label1:
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 62: assert not a == b == c
					πF.SetLineno(62)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µa, µb); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label2
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µb, µc); πE != nil {
						continue
					}
				Label2:
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp004).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßTestRandom.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 65: def TestRandomUniform():
			πF.SetLineno(65)
			πTemp003 = make([]πg.Param, 0)
			πTemp006 = πg.NewFunction(πg.NewCode("TestRandomUniform", "build/src/__python__/random_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
				var µa *πg.Object = πg.UnboundLocal; _ = µa
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
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(10).ToObject()
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
						µ_ = πTemp003
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 67: a = random.uniform(0, 1000)
					πF.SetLineno(67)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewInt(0).ToObject()
					πTemp002[1] = πg.NewInt(1000).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßrandom); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßuniform, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µa = πTemp003
					// line 68: assert isinstance(a, float)
					πF.SetLineno(68)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp002[0] = µa
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
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
					if πE = πg.Assert(πF, πTemp004, nil); πE != nil {
						continue
					}
					// line 69: assert 0 <= a <= 1000
					πF.SetLineno(69)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LE(πF, πg.NewInt(0).ToObject(), µa); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label4
					}
					if πTemp003, πE = πg.LE(πF, µa, πg.NewInt(1000).ToObject()); πE != nil {
						continue
					}
				Label4:
					if πE = πg.Assert(πF, πTemp003, nil); πE != nil {
						continue
					}
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßTestRandomUniform.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 72: def TestRandomInt():
			πF.SetLineno(72)
			πTemp003 = make([]πg.Param, 0)
			πTemp007 = πg.NewFunction(πg.NewCode("TestRandomInt", "build/src/__python__/random_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
				var µa *πg.Object = πg.UnboundLocal; _ = µa
				var µb *πg.Object = πg.UnboundLocal; _ = µb
				var µc *πg.Object = πg.UnboundLocal; _ = µc
				var µd *πg.Object = πg.UnboundLocal; _ = µd
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
				var πTemp007 *πg.BaseException
				_ = πTemp007
				var πTemp008 *πg.Traceback
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					case 6: goto Label6
					case 9: goto Label9
					default: panic("unexpected function state")
					}
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(10).ToObject()
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
						µ_ = πTemp003
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 74: a = random.randint(0, 1000000)
					πF.SetLineno(74)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewInt(0).ToObject()
					πTemp002[1] = πg.NewInt(1000000).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßrandom); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßrandint, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µa = πTemp003
					// line 75: assert isinstance(a, int)
					πF.SetLineno(75)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp002[0] = µa
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
					if πE = πg.Assert(πF, πTemp004, nil); πE != nil {
						continue
					}
					// line 76: assert 0 <= a <= 1000000
					πF.SetLineno(76)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LE(πF, πg.NewInt(0).ToObject(), µa); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label4
					}
					if πTemp003, πE = πg.LE(πF, µa, πg.NewInt(1000000).ToObject()); πE != nil {
						continue
					}
				Label4:
					if πE = πg.Assert(πF, πTemp003, nil); πE != nil {
						continue
					}
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 78: b = random.randint(1, 1)
					πF.SetLineno(78)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewInt(1).ToObject()
					πTemp002[1] = πg.NewInt(1).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßrandom); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßrandint, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µb = πTemp001
					// line 79: assert b == 1
					πF.SetLineno(79)
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µb, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 81: try:
					πF.SetLineno(81)
					πF.PushCheckpoint(6)
					// line 82: c = random.randint(0.1, 3)
					πF.SetLineno(82)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewFloat(0.1).ToObject()
					πTemp002[1] = πg.NewInt(3).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßrandom); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßrandint, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µc = πTemp001
					πF.PopCheckpoint()
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("ValueError not raised").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 86: raise AssertionError("ValueError not raised")
					πF.SetLineno(86)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label5
				Label6:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label7
					}
					πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
					continue
					// line 83: except ValueError:
					πF.SetLineno(83)
				Label7:
					// line 84: pass
					πF.SetLineno(84)
					πF.RestoreExc(nil, nil)
					goto Label5
				Label5:
					// line 88: try:
					πF.SetLineno(88)
					πF.PushCheckpoint(9)
					// line 89: d = random.randint(4, 3)
					πF.SetLineno(89)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewInt(4).ToObject()
					πTemp002[1] = πg.NewInt(3).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßrandom); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßrandint, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µd = πTemp001
					πF.PopCheckpoint()
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("ValueError not raised").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 93: raise AssertionError("ValueError not raised")
					πF.SetLineno(93)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label8
				Label9:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label10
					}
					πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
					continue
					// line 90: except ValueError:
					πF.SetLineno(90)
				Label10:
					// line 91: pass
					πF.SetLineno(91)
					πF.RestoreExc(nil, nil)
					goto Label8
				Label8:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßTestRandomInt.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 96: def TestRandomChoice():
			πF.SetLineno(96)
			πTemp003 = make([]πg.Param, 0)
			πTemp008 = πg.NewFunction(πg.NewCode("TestRandomChoice", "build/src/__python__/random_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µseq *πg.Object = πg.UnboundLocal; _ = µseq
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µitem *πg.Object = πg.UnboundLocal; _ = µitem
				var µitem_idx *πg.Object = πg.UnboundLocal; _ = µitem_idx
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []πg.Param
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 []*πg.Object
				_ = πTemp010
				var πTemp011 *πg.BaseException
				_ = πTemp011
				var πTemp012 *πg.Traceback
				_ = πTemp012
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 97: seq = [i*2 for i in range(5)]
					πF.SetLineno(97)
					πTemp003 = make([]πg.Param, 0)
					πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/random_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µi *πg.Object = πg.UnboundLocal; _ = µi
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
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 4: goto Label4
								default: panic("unexpected function state")
								}
								πTemp002 = πF.MakeArgs(1)
								πTemp002[0] = πg.NewInt(5).ToObject()
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
								// line 97: seq = [i*2 for i in range(5)]
								πF.SetLineno(97)
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Mul(πF, µi, πg.NewInt(2).ToObject()); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return πTemp003, nil
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
					if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
						continue
					}
					µseq = πTemp001
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = πg.NewInt(10).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp001, πE = πg.Iter(πF, πTemp006); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp007 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp007 {
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
						πTemp008 = !isStop
					} else {
						πTemp008 = true
						µi = πTemp004
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 99: item = random.choice(seq)
					πF.SetLineno(99)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µseq, "seq"); πE != nil {
						continue
					}
					πTemp005[0] = µseq
					if πTemp004, πE = πg.ResolveGlobal(πF, ßrandom); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßchoice, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µitem = πTemp004
					// line 100: item_idx = item/2
					πF.SetLineno(100)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Div(πF, µitem, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					µitem_idx = πTemp004
					// line 101: assert seq[item_idx] == item
					πF.SetLineno(101)
					if πE = πg.CheckLocal(πF, µitem_idx, "item_idx"); πE != nil {
						continue
					}
					πTemp006 = µitem_idx
					if πE = πg.CheckLocal(πF, µseq, "seq"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, µseq, πTemp006); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, πTemp009, µitem); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp004, nil); πE != nil {
						continue
					}
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 103: try:
					πF.SetLineno(103)
					πF.PushCheckpoint(5)
					// line 104: random.choice([])
					πF.SetLineno(104)
					πTemp005 = πF.MakeArgs(1)
					πTemp010 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp010...).ToObject()
					πTemp005[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßrandom); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßchoice, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πF.PopCheckpoint()
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = πg.NewStr("IndexError not raised").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 108: raise AssertionError("IndexError not raised")
					πF.SetLineno(108)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					goto Label4
				Label5:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp011, πTemp012 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label6
					}
					πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
					continue
					// line 105: except IndexError:
					πF.SetLineno(105)
				Label6:
					// line 106: pass
					πF.SetLineno(106)
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
			if πE = πF.Globals().SetItem(πF, ßTestRandomChoice.ToObject(), πTemp008); πE != nil {
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
			// line 111: if __name__ == '__main__':
			πF.SetLineno(111)
		Label1:
			// line 112: weetest.RunTests()
			πF.SetLineno(112)
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
	πg.RegisterModule("random_test", Code)
}
