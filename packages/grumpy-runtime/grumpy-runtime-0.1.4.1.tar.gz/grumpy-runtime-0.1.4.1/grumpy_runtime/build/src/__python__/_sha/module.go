package _sha
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/_sha.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßH0 := πg.InternStr("H0")
		ßH1 := πg.InternStr("H1")
		ßH2 := πg.InternStr("H2")
		ßH3 := πg.InternStr("H3")
		ßH4 := πg.InternStr("H4")
		ßK := πg.InternStr("K")
		ßNone := πg.InternStr("None")
		ß__date__ := πg.InternStr("__date__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__version__ := πg.InternStr("__version__")
		ß_bytelist2longBigEndian := πg.InternStr("_bytelist2longBigEndian")
		ß_long2bytesBigEndian := πg.InternStr("_long2bytesBigEndian")
		ß_rotateLeft := πg.InternStr("_rotateLeft")
		ß_transform := πg.InternStr("_transform")
		ßappend := πg.InternStr("append")
		ßblock_size := πg.InternStr("block_size")
		ßblocksize := πg.InternStr("blocksize")
		ßcopy := πg.InternStr("copy")
		ßcount := πg.InternStr("count")
		ßdeepcopy := πg.InternStr("deepcopy")
		ßdigest := πg.InternStr("digest")
		ßdigest_size := πg.InternStr("digest_size")
		ßdigestsize := πg.InternStr("digestsize")
		ßf := πg.InternStr("f")
		ßf0_19 := πg.InternStr("f0_19")
		ßf20_39 := πg.InternStr("f20_39")
		ßf40_59 := πg.InternStr("f40_59")
		ßf60_79 := πg.InternStr("f60_79")
		ßhexdigest := πg.InternStr("hexdigest")
		ßinit := πg.InternStr("init")
		ßinput := πg.InternStr("input")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßlength := πg.InternStr("length")
		ßlist := πg.InternStr("list")
		ßnew := πg.InternStr("new")
		ßobject := πg.InternStr("object")
		ßord := πg.InternStr("ord")
		ßpack := πg.InternStr("pack")
		ßrange := πg.InternStr("range")
		ßsha := πg.InternStr("sha")
		ßstruct := πg.InternStr("struct")
		ßupdate := πg.InternStr("update")
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
		var πTemp011 *πg.Dict
		_ = πTemp011
		var πTemp012 *πg.Object
		_ = πTemp012
		var πTemp013 *πg.Object
		_ = πTemp013
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 7: """A sample implementation of SHA-1 in pure Python.
			πF.SetLineno(7)
			// line 15: __date__    = '2004-11-17'
			πF.SetLineno(15)
			if πE = πF.Globals().SetItem(πF, ß__date__.ToObject(), πg.NewStr("2004-11-17").ToObject()); πE != nil {
				continue
			}
			// line 16: __version__ = 0.91 # Modernised by J. Hallï¿½n and L. Creighton for Pypy
			πF.SetLineno(16)
			if πE = πF.Globals().SetItem(πF, ß__version__.ToObject(), πg.NewFloat(0.91).ToObject()); πE != nil {
				continue
			}
			// line 19: import _struct as struct
			πF.SetLineno(19)
			if πTemp002, πE = πg.ImportModule(πF, "_struct"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßstruct.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 20: import copy
			πF.SetLineno(20)
			if πTemp002, πE = πg.ImportModule(πF, "copy"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßcopy.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 30: def _long2bytesBigEndian(n, blocksize=0):
			πF.SetLineno(30)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "n", Def: nil}
			πTemp003[1] = πg.Param{Name: "blocksize", Def: πg.NewInt(0).ToObject()}
			πTemp001 = πg.NewFunction(πg.NewCode("_long2bytesBigEndian", "build/src/__python__/_sha.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µn *πg.Object = πArgs[0]; _ = µn
				var µblocksize *πg.Object = πArgs[1]; _ = µblocksize
				var µs *πg.Object = πg.UnboundLocal; _ = µs
				var µpack *πg.Object = πg.UnboundLocal; _ = µpack
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
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
					case 1: goto Label1
					case 2: goto Label2
					case 4: goto Label4
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 31: """Convert a long integer to a byte string.
					πF.SetLineno(31)
					// line 39: s = b''
					πF.SetLineno(39)
					µs = ß.ToObject()
					// line 40: pack = struct.pack
					πF.SetLineno(40)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßstruct); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpack, nil); πE != nil {
						continue
					}
					µpack = πTemp002
					// line 41: while n > 0:
					πF.SetLineno(41)
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
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GT(πF, µn, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 42: s = pack('>I', n & 0xffffffff) + s
					πF.SetLineno(42)
					πTemp005 = πF.MakeArgs(2)
					πTemp005[0] = πg.NewStr(">I").ToObject()
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, µn, πg.NewInt(4294967295).ToObject()); πE != nil {
						continue
					}
					πTemp005[1] = πTemp002
					if πE = πg.CheckLocal(πF, µpack, "pack"); πE != nil {
						continue
					}
					if πTemp002, πE = µpack.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp002, µs); πE != nil {
						continue
					}
					µs = πTemp001
					// line 43: n = n >> 32
					πF.SetLineno(43)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.RShift(πF, µn, πg.NewInt(32).ToObject()); πE != nil {
						continue
					}
					µn = πTemp001
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					πTemp005 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp006[0] = µs
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp005[0] = πTemp007
					if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp001, πE = πg.Iter(πF, πTemp007); πE != nil {
						continue
					}
					πF.PushCheckpoint(5)
					πTemp003 = false
				Label4:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp003 {
						πF.PopCheckpoint()
						goto Label6
					}
					if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp004 = !isStop
					} else {
						πTemp004 = true
						µi = πTemp002
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(4)            
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp007 = µi
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µs, πTemp007); πE != nil {
						continue
					}
					if πTemp002, πE = πg.NE(πF, πTemp008, πg.NewStr("\x00").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label7
					}
					goto Label8
					// line 47: if s[i] != '\000':
					πF.SetLineno(47)
				Label7:
					// line 48: break
					πF.SetLineno(48)
					πTemp003 = true
					continue
					goto Label8
				Label8:
					continue
				Label5:
					if πE != nil || πR != nil {
						continue
					}
					// line 51: s = '\000'
					πF.SetLineno(51)
					µs = πg.NewStr("\x00").ToObject()
					// line 52: i = 0
					πF.SetLineno(52)
					µi = πg.NewInt(0).ToObject()
				Label6:
					// line 54: s = s[i:]
					πF.SetLineno(54)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µi, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µs, πTemp001); πE != nil {
						continue
					}
					µs = πTemp002
					if πE = πg.CheckLocal(πF, µblocksize, "blocksize"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GT(πF, µblocksize, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp002
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label9
					}
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp005[0] = µs
					if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πg.CheckLocal(πF, µblocksize, "blocksize"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πTemp008, µblocksize); πE != nil {
						continue
					}
					πTemp001 = πTemp002
				Label9:
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label10
					}
					goto Label11
					// line 58: if blocksize > 0 and len(s) % blocksize:
					πF.SetLineno(58)
				Label10:
					// line 59: s = (blocksize - len(s) % blocksize) * '\000' + s
					πF.SetLineno(59)
					if πE = πg.CheckLocal(πF, µblocksize, "blocksize"); πE != nil {
						continue
					}
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp005[0] = µs
					if πTemp009, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πg.CheckLocal(πF, µblocksize, "blocksize"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Mod(πF, πTemp010, µblocksize); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Sub(πF, µblocksize, πTemp008); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mul(πF, πTemp007, πg.NewStr("\x00").ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp002, µs); πE != nil {
						continue
					}
					µs = πTemp001
					goto Label11
				Label11:
					// line 61: return s
					πF.SetLineno(61)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πR = µs
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_long2bytesBigEndian.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 64: def _bytelist2longBigEndian(list):
			πF.SetLineno(64)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "list", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("_bytelist2longBigEndian", "build/src/__python__/_sha.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µlist *πg.Object = πArgs[0]; _ = µlist
				var µimax *πg.Object = πg.UnboundLocal; _ = µimax
				var µhl *πg.Object = πg.UnboundLocal; _ = µhl
				var µj *πg.Object = πg.UnboundLocal; _ = µj
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µb0 *πg.Object = πg.UnboundLocal; _ = µb0
				var µb1 *πg.Object = πg.UnboundLocal; _ = µb1
				var µb2 *πg.Object = πg.UnboundLocal; _ = µb2
				var µb3 *πg.Object = πg.UnboundLocal; _ = µb3
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
					// line 65: "Transform a list of characters into a list of longs."
					πF.SetLineno(65)
					// line 67: imax = len(list) // 4
					πF.SetLineno(67)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
						continue
					}
					πTemp002[0] = µlist
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.FloorDiv(πF, πTemp004, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					µimax = πTemp001
					// line 68: hl = [0] * imax
					πF.SetLineno(68)
					πTemp002 = make([]*πg.Object, 1)
					πTemp002[0] = πg.NewInt(0).ToObject()
					πTemp003 = πg.NewList(πTemp002...).ToObject()
					if πE = πg.CheckLocal(πF, µimax, "imax"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mul(πF, πTemp003, µimax); πE != nil {
						continue
					}
					µhl = πTemp001
					// line 70: j = 0
					πF.SetLineno(70)
					µj = πg.NewInt(0).ToObject()
					// line 71: i = 0
					πF.SetLineno(71)
					µi = πg.NewInt(0).ToObject()
					// line 72: while i < imax:
					πF.SetLineno(72)
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
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µimax, "imax"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µi, µimax); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 73: b0 = ord(list[j]) << 24
					πF.SetLineno(73)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					πTemp003 = µj
					if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µlist, πTemp003); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.LShift(πF, πTemp004, πg.NewInt(24).ToObject()); πE != nil {
						continue
					}
					µb0 = πTemp001
					// line 74: b1 = ord(list[j+1]) << 16
					πF.SetLineno(74)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, µj, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp004
					if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µlist, πTemp003); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.LShift(πF, πTemp004, πg.NewInt(16).ToObject()); πE != nil {
						continue
					}
					µb1 = πTemp001
					// line 75: b2 = ord(list[j+2]) << 8
					πF.SetLineno(75)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, µj, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp004
					if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µlist, πTemp003); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.LShift(πF, πTemp004, πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					µb2 = πTemp001
					// line 76: b3 = ord(list[j+3])
					πF.SetLineno(76)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µj, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µlist, πTemp001); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µb3 = πTemp003
					// line 77: hl[i] = b0 | b1 | b2 | b3
					πF.SetLineno(77)
					if πE = πg.CheckLocal(πF, µb0, "b0"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb1, "b1"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Or(πF, µb0, µb1); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb2, "b2"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Or(πF, πTemp004, µb2); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb3, "b3"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Or(πF, πTemp003, µb3); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µhl, "hl"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp004 = µi
					if πE = πg.SetItem(πF, µhl, πTemp004, πTemp003); πE != nil {
						continue
					}
					// line 78: i = i+1
					πF.SetLineno(78)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µi = πTemp001
					// line 79: j = j+4
					πF.SetLineno(79)
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µj, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					µj = πTemp001
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 81: return hl
					πF.SetLineno(81)
					if πE = πg.CheckLocal(πF, µhl, "hl"); πE != nil {
						continue
					}
					πR = µhl
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_bytelist2longBigEndian.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 84: def _rotateLeft(x, n):
			πF.SetLineno(84)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "x", Def: nil}
			πTemp003[1] = πg.Param{Name: "n", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("_rotateLeft", "build/src/__python__/_sha.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µn *πg.Object = πArgs[1]; _ = µn
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
					// line 85: "Rotate x (32 bit) left n bits circularly."
					πF.SetLineno(85)
					// line 87: return (x << n) | (x >> (32-n))
					πF.SetLineno(87)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LShift(πF, µx, µn); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Sub(πF, πg.NewInt(32).ToObject(), µn); πE != nil {
						continue
					}
					if πTemp003, πE = πg.RShift(πF, µx, πTemp004); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Or(πF, πTemp002, πTemp003); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_rotateLeft.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 95: def f0_19(B, C, D):
			πF.SetLineno(95)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "B", Def: nil}
			πTemp003[1] = πg.Param{Name: "C", Def: nil}
			πTemp003[2] = πg.Param{Name: "D", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("f0_19", "build/src/__python__/_sha.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µB *πg.Object = πArgs[0]; _ = µB
				var µC *πg.Object = πArgs[1]; _ = µC
				var µD *πg.Object = πArgs[2]; _ = µD
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
					// line 96: return (B & C) | ((~ B) & D)
					πF.SetLineno(96)
					if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, µB, µC); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Invert(πF, µB); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µD, "D"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.And(πF, πTemp004, µD); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Or(πF, πTemp002, πTemp003); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßf0_19.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 98: def f20_39(B, C, D):
			πF.SetLineno(98)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "B", Def: nil}
			πTemp003[1] = πg.Param{Name: "C", Def: nil}
			πTemp003[2] = πg.Param{Name: "D", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("f20_39", "build/src/__python__/_sha.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µB *πg.Object = πArgs[0]; _ = µB
				var µC *πg.Object = πArgs[1]; _ = µC
				var µD *πg.Object = πArgs[2]; _ = µD
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
					// line 99: return B ^ C ^ D
					πF.SetLineno(99)
					if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Xor(πF, µB, µC); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µD, "D"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Xor(πF, πTemp002, µD); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßf20_39.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 101: def f40_59(B, C, D):
			πF.SetLineno(101)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "B", Def: nil}
			πTemp003[1] = πg.Param{Name: "C", Def: nil}
			πTemp003[2] = πg.Param{Name: "D", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("f40_59", "build/src/__python__/_sha.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µB *πg.Object = πArgs[0]; _ = µB
				var µC *πg.Object = πArgs[1]; _ = µC
				var µD *πg.Object = πArgs[2]; _ = µD
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
					// line 102: return (B & C) | (B & D) | (C & D)
					πF.SetLineno(102)
					if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.And(πF, µB, µC); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µD, "D"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.And(πF, µB, µD); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Or(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µD, "D"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.And(πF, µC, µD); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Or(πF, πTemp002, πTemp003); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßf40_59.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 104: def f60_79(B, C, D):
			πF.SetLineno(104)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "B", Def: nil}
			πTemp003[1] = πg.Param{Name: "C", Def: nil}
			πTemp003[2] = πg.Param{Name: "D", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("f60_79", "build/src/__python__/_sha.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µB *πg.Object = πArgs[0]; _ = µB
				var µC *πg.Object = πArgs[1]; _ = µC
				var µD *πg.Object = πArgs[2]; _ = µD
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
					// line 105: return B ^ C ^ D
					πF.SetLineno(105)
					if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Xor(πF, µB, µC); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µD, "D"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Xor(πF, πTemp002, µD); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßf60_79.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 108: f = [f0_19, f20_39, f40_59, f60_79]
			πF.SetLineno(108)
			πTemp002 = make([]*πg.Object, 4)
			if πTemp010, πE = πg.ResolveGlobal(πF, ßf0_19); πE != nil {
				continue
			}
			πTemp002[0] = πTemp010
			if πTemp010, πE = πg.ResolveGlobal(πF, ßf20_39); πE != nil {
				continue
			}
			πTemp002[1] = πTemp010
			if πTemp010, πE = πg.ResolveGlobal(πF, ßf40_59); πE != nil {
				continue
			}
			πTemp002[2] = πTemp010
			if πTemp010, πE = πg.ResolveGlobal(πF, ßf60_79); πE != nil {
				continue
			}
			πTemp002[3] = πTemp010
			πTemp010 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßf.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 111: K = [
			πF.SetLineno(111)
			πTemp002 = make([]*πg.Object, 4)
			πTemp002[0] = πg.NewInt(1518500249).ToObject()
			πTemp002[1] = πg.NewInt(1859775393).ToObject()
			πTemp002[2] = πg.NewInt(2400959708).ToObject()
			πTemp002[3] = πg.NewInt(3395469782).ToObject()
			πTemp010 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßK.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 118: class sha(object):
			πF.SetLineno(118)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp013
			πTemp011 = πg.NewDict()
			if πTemp010, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, ß__module__.ToObject(), πTemp010); πE != nil {
				continue
			}
			_, πE = πg.NewCode("sha", "build/src/__python__/_sha.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp011
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
					// line 119: "An implementation of the SHA hash function in pure Python."
					πF.SetLineno(119)
					// line 121: digest_size = digestsize = 20
					πF.SetLineno(121)
					if πE = πClass.SetItem(πF, ßdigest_size.ToObject(), πg.NewInt(20).ToObject()); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdigestsize.ToObject(), πg.NewInt(20).ToObject()); πE != nil {
						continue
					}
					// line 122: block_size = 512 // 8
					πF.SetLineno(122)
					if πTemp001, πE = πg.FloorDiv(πF, πg.NewInt(512).ToObject(), πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßblock_size.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 124: def __init__(self):
					πF.SetLineno(124)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/_sha.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
							// line 125: "Initialisation."
							πF.SetLineno(125)
							// line 128: self.length = 0
							πF.SetLineno(128)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlength, πTemp001); πE != nil {
								continue
							}
							// line 129: self.count = [0, 0]
							πF.SetLineno(129)
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(0).ToObject()
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcount, πTemp003); πE != nil {
								continue
							}
							// line 132: self.input = []
							πF.SetLineno(132)
							πTemp002 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßinput, πTemp003); πE != nil {
								continue
							}
							// line 136: self.init()
							πF.SetLineno(136)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßinit, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
					// line 139: def init(self):
					πF.SetLineno(139)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("init", "build/src/__python__/_sha.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
							// line 140: "Initialize the message-digest and set all fields to zero."
							πF.SetLineno(140)
							// line 142: self.length = 0
							πF.SetLineno(142)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlength, πTemp001); πE != nil {
								continue
							}
							// line 143: self.input = []
							πF.SetLineno(143)
							πTemp002 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßinput, πTemp003); πE != nil {
								continue
							}
							// line 146: self.H0 = 0x67452301
							πF.SetLineno(146)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1732584193).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßH0, πTemp001); πE != nil {
								continue
							}
							// line 147: self.H1 = 0xEFCDAB89
							πF.SetLineno(147)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(4023233417).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßH1, πTemp001); πE != nil {
								continue
							}
							// line 148: self.H2 = 0x98BADCFE
							πF.SetLineno(148)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(2562383102).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßH2, πTemp001); πE != nil {
								continue
							}
							// line 149: self.H3 = 0x10325476
							πF.SetLineno(149)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(271733878).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßH3, πTemp001); πE != nil {
								continue
							}
							// line 150: self.H4 = 0xC3D2E1F0
							πF.SetLineno(150)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(3285377520).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßH4, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßinit.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 152: def _transform(self, W):
					πF.SetLineno(152)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "W", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("_transform", "build/src/__python__/_sha.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µW *πg.Object = πArgs[1]; _ = µW
						var µt *πg.Object = πg.UnboundLocal; _ = µt
						var µA *πg.Object = πg.UnboundLocal; _ = µA
						var µB *πg.Object = πg.UnboundLocal; _ = µB
						var µC *πg.Object = πg.UnboundLocal; _ = µC
						var µD *πg.Object = πg.UnboundLocal; _ = µD
						var µE *πg.Object = πg.UnboundLocal; _ = µE
						var µTEMP *πg.Object = πg.UnboundLocal; _ = µTEMP
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							case 4: goto Label4
							case 5: goto Label5
							case 7: goto Label7
							case 8: goto Label8
							case 10: goto Label10
							case 11: goto Label11
							case 13: goto Label13
							case 14: goto Label14
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewInt(16).ToObject()
							πTemp002[1] = πg.NewInt(80).ToObject()
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
								µt = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 155: W.append(_rotateLeft(
							πF.SetLineno(155)
							πTemp002 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Sub(πF, µt, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp010 = πTemp011
							if πE = πg.CheckLocal(πF, µW, "W"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetItem(πF, µW, πTemp010); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.Sub(πF, µt, πg.NewInt(8).ToObject()); πE != nil {
								continue
							}
							πTemp010 = πTemp012
							if πE = πg.CheckLocal(πF, µW, "W"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, µW, πTemp010); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Xor(πF, πTemp011, πTemp012); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Sub(πF, µt, πg.NewInt(14).ToObject()); πE != nil {
								continue
							}
							πTemp010 = πTemp011
							if πE = πg.CheckLocal(πF, µW, "W"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetItem(πF, µW, πTemp010); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Xor(πF, πTemp009, πTemp011); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Sub(πF, µt, πg.NewInt(16).ToObject()); πE != nil {
								continue
							}
							πTemp009 = πTemp010
							if πE = πg.CheckLocal(πF, µW, "W"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µW, πTemp009); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Xor(πF, πTemp008, πTemp010); πE != nil {
								continue
							}
							πTemp007[0] = πTemp004
							πTemp007[1] = πg.NewInt(1).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_rotateLeft); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp003, πE = πg.And(πF, πTemp008, πg.NewInt(4294967295).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µW, "W"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µW, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 158: A = self.H0
							πF.SetLineno(158)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßH0, nil); πE != nil {
								continue
							}
							µA = πTemp001
							// line 159: B = self.H1
							πF.SetLineno(159)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßH1, nil); πE != nil {
								continue
							}
							µB = πTemp001
							// line 160: C = self.H2
							πF.SetLineno(160)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßH2, nil); πE != nil {
								continue
							}
							µC = πTemp001
							// line 161: D = self.H3
							πF.SetLineno(161)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßH3, nil); πE != nil {
								continue
							}
							µD = πTemp001
							// line 162: E = self.H4
							πF.SetLineno(162)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßH4, nil); πE != nil {
								continue
							}
							µE = πTemp001
							// line 164: """
							πF.SetLineno(164)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(20).ToObject()
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
							πF.PushCheckpoint(5)
							πTemp005 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label6
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
								µt = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 176: TEMP = _rotateLeft(A, 5) + ((B & C) | ((~ B) & D)) + E + W[t] + K[0]
							πF.SetLineno(176)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µA, "A"); πE != nil {
								continue
							}
							πTemp002[0] = µA
							πTemp002[1] = πg.NewInt(5).ToObject()
							if πTemp010, πE = πg.ResolveGlobal(πF, ß_rotateLeft); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp010.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.And(πF, µB, µC); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
								continue
							}
							if πTemp014, πE = πg.Invert(πF, µB); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µD, "D"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.And(πF, πTemp014, µD); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Or(πF, πTemp012, πTemp013); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Add(πF, πTemp011, πTemp010); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µE, "E"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Add(πF, πTemp009, µE); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							πTemp009 = µt
							if πE = πg.CheckLocal(πF, µW, "W"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µW, πTemp009); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp008, πTemp010); πE != nil {
								continue
							}
							πTemp008 = πg.NewInt(0).ToObject()
							if πTemp010, πE = πg.ResolveGlobal(πF, ßK); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, πTemp010, πTemp008); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, πTemp009); πE != nil {
								continue
							}
							µTEMP = πTemp003
							// line 177: E = D
							πF.SetLineno(177)
							if πE = πg.CheckLocal(πF, µD, "D"); πE != nil {
								continue
							}
							µE = µD
							// line 178: D = C
							πF.SetLineno(178)
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							µD = µC
							// line 179: C = _rotateLeft(B, 30) & 0xffffffff
							πF.SetLineno(179)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
								continue
							}
							πTemp002[0] = µB
							πTemp002[1] = πg.NewInt(30).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_rotateLeft); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.And(πF, πTemp008, πg.NewInt(4294967295).ToObject()); πE != nil {
								continue
							}
							µC = πTemp003
							// line 180: B = A
							πF.SetLineno(180)
							if πE = πg.CheckLocal(πF, µA, "A"); πE != nil {
								continue
							}
							µB = µA
							// line 181: A = TEMP & 0xffffffff
							πF.SetLineno(181)
							if πE = πg.CheckLocal(πF, µTEMP, "TEMP"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.And(πF, µTEMP, πg.NewInt(4294967295).ToObject()); πE != nil {
								continue
							}
							µA = πTemp003
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewInt(20).ToObject()
							πTemp002[1] = πg.NewInt(40).ToObject()
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
							πF.PushCheckpoint(8)
							πTemp005 = false
						Label7:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label9
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
								µt = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(7)            
							// line 184: TEMP = _rotateLeft(A, 5) + (B ^ C ^ D) + E + W[t] + K[1]
							πF.SetLineno(184)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µA, "A"); πE != nil {
								continue
							}
							πTemp002[0] = µA
							πTemp002[1] = πg.NewInt(5).ToObject()
							if πTemp010, πE = πg.ResolveGlobal(πF, ß_rotateLeft); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp010.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.Xor(πF, µB, µC); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µD, "D"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Xor(πF, πTemp012, µD); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Add(πF, πTemp011, πTemp010); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µE, "E"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Add(πF, πTemp009, µE); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							πTemp009 = µt
							if πE = πg.CheckLocal(πF, µW, "W"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µW, πTemp009); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp008, πTemp010); πE != nil {
								continue
							}
							πTemp008 = πg.NewInt(1).ToObject()
							if πTemp010, πE = πg.ResolveGlobal(πF, ßK); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, πTemp010, πTemp008); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, πTemp009); πE != nil {
								continue
							}
							µTEMP = πTemp003
							// line 185: E = D
							πF.SetLineno(185)
							if πE = πg.CheckLocal(πF, µD, "D"); πE != nil {
								continue
							}
							µE = µD
							// line 186: D = C
							πF.SetLineno(186)
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							µD = µC
							// line 187: C = _rotateLeft(B, 30) & 0xffffffff
							πF.SetLineno(187)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
								continue
							}
							πTemp002[0] = µB
							πTemp002[1] = πg.NewInt(30).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_rotateLeft); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.And(πF, πTemp008, πg.NewInt(4294967295).ToObject()); πE != nil {
								continue
							}
							µC = πTemp003
							// line 188: B = A
							πF.SetLineno(188)
							if πE = πg.CheckLocal(πF, µA, "A"); πE != nil {
								continue
							}
							µB = µA
							// line 189: A = TEMP & 0xffffffff
							πF.SetLineno(189)
							if πE = πg.CheckLocal(πF, µTEMP, "TEMP"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.And(πF, µTEMP, πg.NewInt(4294967295).ToObject()); πE != nil {
								continue
							}
							µA = πTemp003
							continue
						Label8:
							if πE != nil || πR != nil {
								continue
							}
						Label9:
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewInt(40).ToObject()
							πTemp002[1] = πg.NewInt(60).ToObject()
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
							πF.PushCheckpoint(11)
							πTemp005 = false
						Label10:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label12
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
								µt = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(10)            
							// line 192: TEMP = _rotateLeft(A, 5) + ((B & C) | (B & D) | (C & D)) + E + W[t] + K[2]
							πF.SetLineno(192)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µA, "A"); πE != nil {
								continue
							}
							πTemp002[0] = µA
							πTemp002[1] = πg.NewInt(5).ToObject()
							if πTemp010, πE = πg.ResolveGlobal(πF, ß_rotateLeft); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp010.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.And(πF, µB, µC); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µD, "D"); πE != nil {
								continue
							}
							if πTemp014, πE = πg.And(πF, µB, µD); πE != nil {
								continue
							}
							if πTemp012, πE = πg.Or(πF, πTemp013, πTemp014); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µD, "D"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.And(πF, µC, µD); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Or(πF, πTemp012, πTemp013); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Add(πF, πTemp011, πTemp010); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µE, "E"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Add(πF, πTemp009, µE); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							πTemp009 = µt
							if πE = πg.CheckLocal(πF, µW, "W"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µW, πTemp009); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp008, πTemp010); πE != nil {
								continue
							}
							πTemp008 = πg.NewInt(2).ToObject()
							if πTemp010, πE = πg.ResolveGlobal(πF, ßK); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, πTemp010, πTemp008); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, πTemp009); πE != nil {
								continue
							}
							µTEMP = πTemp003
							// line 193: E = D
							πF.SetLineno(193)
							if πE = πg.CheckLocal(πF, µD, "D"); πE != nil {
								continue
							}
							µE = µD
							// line 194: D = C
							πF.SetLineno(194)
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							µD = µC
							// line 195: C = _rotateLeft(B, 30) & 0xffffffff
							πF.SetLineno(195)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
								continue
							}
							πTemp002[0] = µB
							πTemp002[1] = πg.NewInt(30).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_rotateLeft); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.And(πF, πTemp008, πg.NewInt(4294967295).ToObject()); πE != nil {
								continue
							}
							µC = πTemp003
							// line 196: B = A
							πF.SetLineno(196)
							if πE = πg.CheckLocal(πF, µA, "A"); πE != nil {
								continue
							}
							µB = µA
							// line 197: A = TEMP & 0xffffffff
							πF.SetLineno(197)
							if πE = πg.CheckLocal(πF, µTEMP, "TEMP"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.And(πF, µTEMP, πg.NewInt(4294967295).ToObject()); πE != nil {
								continue
							}
							µA = πTemp003
							continue
						Label11:
							if πE != nil || πR != nil {
								continue
							}
						Label12:
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewInt(60).ToObject()
							πTemp002[1] = πg.NewInt(80).ToObject()
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
							πF.PushCheckpoint(14)
							πTemp005 = false
						Label13:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label15
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
								µt = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(13)            
							// line 200: TEMP = _rotateLeft(A, 5) + (B ^ C ^ D)  + E + W[t] + K[3]
							πF.SetLineno(200)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µA, "A"); πE != nil {
								continue
							}
							πTemp002[0] = µA
							πTemp002[1] = πg.NewInt(5).ToObject()
							if πTemp010, πE = πg.ResolveGlobal(πF, ß_rotateLeft); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp010.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.Xor(πF, µB, µC); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µD, "D"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Xor(πF, πTemp012, µD); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Add(πF, πTemp011, πTemp010); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µE, "E"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Add(πF, πTemp009, µE); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							πTemp009 = µt
							if πE = πg.CheckLocal(πF, µW, "W"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µW, πTemp009); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp008, πTemp010); πE != nil {
								continue
							}
							πTemp008 = πg.NewInt(3).ToObject()
							if πTemp010, πE = πg.ResolveGlobal(πF, ßK); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, πTemp010, πTemp008); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, πTemp009); πE != nil {
								continue
							}
							µTEMP = πTemp003
							// line 201: E = D
							πF.SetLineno(201)
							if πE = πg.CheckLocal(πF, µD, "D"); πE != nil {
								continue
							}
							µE = µD
							// line 202: D = C
							πF.SetLineno(202)
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							µD = µC
							// line 203: C = _rotateLeft(B, 30) & 0xffffffff
							πF.SetLineno(203)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
								continue
							}
							πTemp002[0] = µB
							πTemp002[1] = πg.NewInt(30).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_rotateLeft); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.And(πF, πTemp008, πg.NewInt(4294967295).ToObject()); πE != nil {
								continue
							}
							µC = πTemp003
							// line 204: B = A
							πF.SetLineno(204)
							if πE = πg.CheckLocal(πF, µA, "A"); πE != nil {
								continue
							}
							µB = µA
							// line 205: A = TEMP & 0xffffffff
							πF.SetLineno(205)
							if πE = πg.CheckLocal(πF, µTEMP, "TEMP"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.And(πF, µTEMP, πg.NewInt(4294967295).ToObject()); πE != nil {
								continue
							}
							µA = πTemp003
							continue
						Label14:
							if πE != nil || πR != nil {
								continue
							}
						Label15:
							// line 208: self.H0 = (self.H0 + A) & 0xffffffff
							πF.SetLineno(208)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßH0, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µA, "A"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, µA); πE != nil {
								continue
							}
							if πTemp001, πE = πg.And(πF, πTemp003, πg.NewInt(4294967295).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßH0, πTemp003); πE != nil {
								continue
							}
							// line 209: self.H1 = (self.H1 + B) & 0xffffffff
							πF.SetLineno(209)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßH1, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, µB); πE != nil {
								continue
							}
							if πTemp001, πE = πg.And(πF, πTemp003, πg.NewInt(4294967295).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßH1, πTemp003); πE != nil {
								continue
							}
							// line 210: self.H2 = (self.H2 + C) & 0xffffffff
							πF.SetLineno(210)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßH2, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, µC); πE != nil {
								continue
							}
							if πTemp001, πE = πg.And(πF, πTemp003, πg.NewInt(4294967295).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßH2, πTemp003); πE != nil {
								continue
							}
							// line 211: self.H3 = (self.H3 + D) & 0xffffffff
							πF.SetLineno(211)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßH3, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µD, "D"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, µD); πE != nil {
								continue
							}
							if πTemp001, πE = πg.And(πF, πTemp003, πg.NewInt(4294967295).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßH3, πTemp003); πE != nil {
								continue
							}
							// line 212: self.H4 = (self.H4 + E) & 0xffffffff
							πF.SetLineno(212)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßH4, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µE, "E"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, µE); πE != nil {
								continue
							}
							if πTemp001, πE = πg.And(πF, πTemp003, πg.NewInt(4294967295).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßH4, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_transform.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 218: def update(self, inBuf):
					πF.SetLineno(218)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "inBuf", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("update", "build/src/__python__/_sha.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µinBuf *πg.Object = πArgs[1]; _ = µinBuf
						var µleninBuf *πg.Object = πg.UnboundLocal; _ = µleninBuf
						var µindex *πg.Object = πg.UnboundLocal; _ = µindex
						var µpartLen *πg.Object = πg.UnboundLocal; _ = µpartLen
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var πTemp001 []*πg.Object
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
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 6: goto Label6
							case 7: goto Label7
							default: panic("unexpected function state")
							}
							// line 219: """Add to the current message.
							πF.SetLineno(219)
							// line 234: leninBuf = len(inBuf)
							πF.SetLineno(234)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µinBuf, "inBuf"); πE != nil {
								continue
							}
							πTemp001[0] = µinBuf
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µleninBuf = πTemp003
							// line 237: index = (self.count[1] >> 3) & 0x3F
							πF.SetLineno(237)
							πTemp004 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßcount, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.RShift(πF, πTemp005, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.And(πF, πTemp003, πg.NewInt(63).ToObject()); πE != nil {
								continue
							}
							µindex = πTemp002
							// line 240: self.count[1] = self.count[1] + (leninBuf << 3)
							πF.SetLineno(240)
							πTemp003 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßcount, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µleninBuf, "leninBuf"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LShift(πF, µleninBuf, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp004, πTemp003); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcount, nil); πE != nil {
								continue
							}
							πTemp005 = πg.NewInt(1).ToObject()
							if πE = πg.SetItem(πF, πTemp004, πTemp005, πTemp003); πE != nil {
								continue
							}
							πTemp003 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßcount, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µleninBuf, "leninBuf"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LShift(πF, µleninBuf, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, πTemp004, πTemp003); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label1
							}
							goto Label2
							// line 241: if self.count[1] < (leninBuf << 3):
							πF.SetLineno(241)
						Label1:
							// line 242: self.count[0] = self.count[0] + 1
							πF.SetLineno(242)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßcount, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcount, nil); πE != nil {
								continue
							}
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.SetItem(πF, πTemp004, πTemp005, πTemp003); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 243: self.count[0] = self.count[0] + (leninBuf >> 29)
							πF.SetLineno(243)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßcount, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µleninBuf, "leninBuf"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.RShift(πF, µleninBuf, πg.NewInt(29).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp004, πTemp003); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcount, nil); πE != nil {
								continue
							}
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.SetItem(πF, πTemp004, πTemp005, πTemp003); πE != nil {
								continue
							}
							// line 245: partLen = 64 - index
							πF.SetLineno(245)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, πg.NewInt(64).ToObject(), µindex); πE != nil {
								continue
							}
							µpartLen = πTemp002
							if πE = πg.CheckLocal(πF, µleninBuf, "leninBuf"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpartLen, "partLen"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GE(πF, µleninBuf, µpartLen); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label3
							}
							goto Label4
							// line 247: if leninBuf >= partLen:
							πF.SetLineno(247)
						Label3:
							// line 248: self.input[index:] = list(inBuf[:partLen])
							πF.SetLineno(248)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpartLen, "partLen"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µpartLen, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µinBuf, "inBuf"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µinBuf, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßinput, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{µindex, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, πTemp004, πTemp005, πTemp002); πE != nil {
								continue
							}
							// line 249: self._transform(_bytelist2longBigEndian(self.input))
							πF.SetLineno(249)
							πTemp001 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßinput, nil); πE != nil {
								continue
							}
							πTemp008[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_bytelist2longBigEndian); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_transform, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 250: i = partLen
							πF.SetLineno(250)
							if πE = πg.CheckLocal(πF, µpartLen, "partLen"); πE != nil {
								continue
							}
							µi = µpartLen
							// line 251: while i + 63 < leninBuf:
							πF.SetLineno(251)
							πF.PushCheckpoint(7)
							πTemp007 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
								πF.PopCheckpoint()
								goto Label8
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µi, πg.NewInt(63).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µleninBuf, "leninBuf"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, πTemp003, µleninBuf); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(6)            
							// line 252: self._transform(_bytelist2longBigEndian(list(inBuf[i:i+64])))
							πF.SetLineno(252)
							πTemp001 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µi, πg.NewInt(64).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µi, πTemp003, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µinBuf, "inBuf"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µinBuf, πTemp002); πE != nil {
								continue
							}
							πTemp010[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp008[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_bytelist2longBigEndian); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_transform, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 253: i = i + 64
							πF.SetLineno(253)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µi, πg.NewInt(64).ToObject()); πE != nil {
								continue
							}
							µi = πTemp002
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
							// line 255: self.input = list(inBuf[i:leninBuf])
							πF.SetLineno(255)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µleninBuf, "leninBuf"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µi, µleninBuf, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µinBuf, "inBuf"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µinBuf, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßinput, πTemp002); πE != nil {
								continue
							}
						Label8:
							goto Label5
						Label4:
							// line 257: i = 0
							πF.SetLineno(257)
							µi = πg.NewInt(0).ToObject()
							// line 258: self.input = self.input + list(inBuf)
							πF.SetLineno(258)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßinput, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µinBuf, "inBuf"); πE != nil {
								continue
							}
							πTemp001[0] = µinBuf
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßinput, πTemp003); πE != nil {
								continue
							}
							goto Label5
						Label5:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßupdate.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 261: def digest(self):
					πF.SetLineno(261)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("digest", "build/src/__python__/_sha.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µH0 *πg.Object = πg.UnboundLocal; _ = µH0
						var µH1 *πg.Object = πg.UnboundLocal; _ = µH1
						var µH2 *πg.Object = πg.UnboundLocal; _ = µH2
						var µH3 *πg.Object = πg.UnboundLocal; _ = µH3
						var µH4 *πg.Object = πg.UnboundLocal; _ = µH4
						var µinput *πg.Object = πg.UnboundLocal; _ = µinput
						var µcount *πg.Object = πg.UnboundLocal; _ = µcount
						var µindex *πg.Object = πg.UnboundLocal; _ = µindex
						var µpadLen *πg.Object = πg.UnboundLocal; _ = µpadLen
						var µpadding *πg.Object = πg.UnboundLocal; _ = µpadding
						var µbits *πg.Object = πg.UnboundLocal; _ = µbits
						var µdigest *πg.Object = πg.UnboundLocal; _ = µdigest
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
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 262: """Terminate the message-digest computation and return digest.
							πF.SetLineno(262)
							// line 269: H0 = self.H0
							πF.SetLineno(269)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßH0, nil); πE != nil {
								continue
							}
							µH0 = πTemp001
							// line 270: H1 = self.H1
							πF.SetLineno(270)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßH1, nil); πE != nil {
								continue
							}
							µH1 = πTemp001
							// line 271: H2 = self.H2
							πF.SetLineno(271)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßH2, nil); πE != nil {
								continue
							}
							µH2 = πTemp001
							// line 272: H3 = self.H3
							πF.SetLineno(272)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßH3, nil); πE != nil {
								continue
							}
							µH3 = πTemp001
							// line 273: H4 = self.H4
							πF.SetLineno(273)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßH4, nil); πE != nil {
								continue
							}
							µH4 = πTemp001
							// line 274: input = [] + self.input
							πF.SetLineno(274)
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßinput, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							µinput = πTemp001
							// line 275: count = [] + self.count
							πF.SetLineno(275)
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcount, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							µcount = πTemp001
							// line 277: index = (self.count[1] >> 3) & 0x3f
							πF.SetLineno(277)
							πTemp004 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßcount, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.RShift(πF, πTemp005, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.And(πF, πTemp003, πg.NewInt(63).ToObject()); πE != nil {
								continue
							}
							µindex = πTemp001
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µindex, πg.NewInt(56).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label1
							}
							goto Label2
							// line 279: if index < 56:
							πF.SetLineno(279)
						Label1:
							// line 280: padLen = 56 - index
							πF.SetLineno(280)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, πg.NewInt(56).ToObject(), µindex); πE != nil {
								continue
							}
							µpadLen = πTemp001
							goto Label3
						Label2:
							// line 282: padLen = 120 - index
							πF.SetLineno(282)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, πg.NewInt(120).ToObject(), µindex); πE != nil {
								continue
							}
							µpadLen = πTemp001
							goto Label3
						Label3:
							// line 284: padding = ['\200'] + ['\000'] * 63
							πF.SetLineno(284)
							πTemp002 = make([]*πg.Object, 1)
							πTemp002[0] = πg.NewStr("\x80").ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp002 = make([]*πg.Object, 1)
							πTemp002[0] = πg.NewStr("\x00").ToObject()
							πTemp005 = πg.NewList(πTemp002...).ToObject()
							if πTemp004, πE = πg.Mul(πF, πTemp005, πg.NewInt(63).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							µpadding = πTemp001
							// line 285: self.update(padding[:padLen])
							πF.SetLineno(285)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpadLen, "padLen"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µpadLen, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpadding, "padding"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µpadding, πTemp001); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 288: bits = _bytelist2longBigEndian(self.input[:56]) + count
							πF.SetLineno(288)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(56).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßinput, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_bytelist2longBigEndian); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp004, µcount); πE != nil {
								continue
							}
							µbits = πTemp001
							// line 290: self._transform(bits)
							πF.SetLineno(290)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbits, "bits"); πE != nil {
								continue
							}
							πTemp002[0] = µbits
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_transform, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 293: digest = _long2bytesBigEndian(self.H0, 4) + \
							πF.SetLineno(293)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßH0, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp006
							πTemp002[1] = πg.NewInt(4).ToObject()
							if πTemp006, πE = πg.ResolveGlobal(πF, ß_long2bytesBigEndian); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßH1, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp006
							πTemp002[1] = πg.NewInt(4).ToObject()
							if πTemp006, πE = πg.ResolveGlobal(πF, ß_long2bytesBigEndian); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.Add(πF, πTemp008, πTemp009); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßH2, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp006
							πTemp002[1] = πg.NewInt(4).ToObject()
							if πTemp006, πE = πg.ResolveGlobal(πF, ß_long2bytesBigEndian); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp004, πE = πg.Add(πF, πTemp005, πTemp008); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßH3, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp005
							πTemp002[1] = πg.NewInt(4).ToObject()
							if πTemp005, πE = πg.ResolveGlobal(πF, ß_long2bytesBigEndian); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.Add(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßH4, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							πTemp002[1] = πg.NewInt(4).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_long2bytesBigEndian); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Add(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							µdigest = πTemp001
							// line 299: self.H0 = H0
							πF.SetLineno(299)
							if πE = πg.CheckLocal(πF, µH0, "H0"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µH0); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßH0, πTemp001); πE != nil {
								continue
							}
							// line 300: self.H1 = H1
							πF.SetLineno(300)
							if πE = πg.CheckLocal(πF, µH1, "H1"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µH1); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßH1, πTemp001); πE != nil {
								continue
							}
							// line 301: self.H2 = H2
							πF.SetLineno(301)
							if πE = πg.CheckLocal(πF, µH2, "H2"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µH2); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßH2, πTemp001); πE != nil {
								continue
							}
							// line 302: self.H3 = H3
							πF.SetLineno(302)
							if πE = πg.CheckLocal(πF, µH3, "H3"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µH3); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßH3, πTemp001); πE != nil {
								continue
							}
							// line 303: self.H4 = H4
							πF.SetLineno(303)
							if πE = πg.CheckLocal(πF, µH4, "H4"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µH4); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßH4, πTemp001); πE != nil {
								continue
							}
							// line 304: self.input = input
							πF.SetLineno(304)
							if πE = πg.CheckLocal(πF, µinput, "input"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µinput); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßinput, πTemp001); πE != nil {
								continue
							}
							// line 305: self.count = count
							πF.SetLineno(305)
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µcount); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcount, πTemp001); πE != nil {
								continue
							}
							// line 307: return digest
							πF.SetLineno(307)
							if πE = πg.CheckLocal(πF, µdigest, "digest"); πE != nil {
								continue
							}
							πR = µdigest
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdigest.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 310: def hexdigest(self):
					πF.SetLineno(310)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("hexdigest", "build/src/__python__/_sha.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 311: """Terminate and return digest in HEX form.
							πF.SetLineno(311)
							// line 318: return ''.join([('0%x' % ord(c))[-2:] for c in self.digest()])
							πF.SetLineno(318)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/_sha.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µc *πg.Object = πg.UnboundLocal; _ = µc
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
								_ = πTemp002
								var πTemp003 *πg.Object
								_ = πTemp003
								var πTemp004 bool
								_ = πTemp004
								var πTemp005 bool
								_ = πTemp005
								var πTemp006 *πg.Object
								_ = πTemp006
								var πTemp007 []*πg.Object
								_ = πTemp007
								var πTemp008 *πg.Object
								_ = πTemp008
								var πTemp009 *πg.Object
								_ = πTemp009
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
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µself, ßdigest, nil); πE != nil {
											continue
										}
										if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
											continue
										}
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
										if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
											isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
											if exc != nil {
												πE = exc
											} else if isStop {
												πE = nil
												πF.RestoreExc(nil, nil)
											}
											πTemp005 = !isStop
										} else {
											πTemp005 = true
											µc = πTemp002
										}
										if πE != nil || !πTemp005 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 318: return ''.join([('0%x' % ord(c))[-2:] for c in self.digest()])
										πF.SetLineno(318)
										if πTemp003, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
											continue
										}
										if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πg.None, πg.None}, nil); πE != nil {
											continue
										}
										πTemp007 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
											continue
										}
										πTemp007[0] = µc
										if πTemp008, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
											continue
										}
										if πTemp009, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp007)
										if πTemp006, πE = πg.Mod(πF, πg.NewStr("0%x").ToObject(), πTemp009); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return πTemp003, nil
									Label4:
										πTemp002 = πSent
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
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßhexdigest.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 320: def copy(self):
					πF.SetLineno(320)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("copy", "build/src/__python__/_sha.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 321: """Return a clone object.
							πF.SetLineno(321)
							// line 328: return copy.deepcopy(self)
							πF.SetLineno(328)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcopy); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdeepcopy, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßcopy.ToObject(), πTemp008); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp012, πE = πTemp011.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp012 == nil {
				πTemp012 = πg.TypeType.ToObject()
			}
			if πTemp013, πE = πTemp012.Call(πF, []*πg.Object{πg.NewStr("sha").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp011.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßsha.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 339: digest_size = 20
			πF.SetLineno(339)
			if πE = πF.Globals().SetItem(πF, ßdigest_size.ToObject(), πg.NewInt(20).ToObject()); πE != nil {
				continue
			}
			// line 340: digestsize = 20
			πF.SetLineno(340)
			if πE = πF.Globals().SetItem(πF, ßdigestsize.ToObject(), πg.NewInt(20).ToObject()); πE != nil {
				continue
			}
			// line 341: blocksize = 1
			πF.SetLineno(341)
			if πE = πF.Globals().SetItem(πF, ßblocksize.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			// line 343: def new(arg=None):
			πF.SetLineno(343)
			πTemp003 = make([]πg.Param, 1)
			if πTemp012, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[0] = πg.Param{Name: "arg", Def: πTemp012}
			πTemp010 = πg.NewFunction(πg.NewCode("new", "build/src/__python__/_sha.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µarg *πg.Object = πArgs[0]; _ = µarg
				var µcrypto *πg.Object = πg.UnboundLocal; _ = µcrypto
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
					// line 344: """Return a new sha crypto object.
					πF.SetLineno(344)
					// line 349: crypto = sha()
					πF.SetLineno(349)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsha); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µcrypto = πTemp002
					if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µarg); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 350: if arg:
					πF.SetLineno(350)
				Label1:
					// line 351: crypto.update(arg)
					πF.SetLineno(351)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
						continue
					}
					πTemp004[0] = µarg
					if πE = πg.CheckLocal(πF, µcrypto, "crypto"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µcrypto, ßupdate, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					goto Label2
				Label2:
					// line 353: return crypto
					πF.SetLineno(353)
					if πE = πg.CheckLocal(πF, µcrypto, "crypto"); πE != nil {
						continue
					}
					πR = µcrypto
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßnew.ToObject(), πTemp010); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("_sha", Code)
}
