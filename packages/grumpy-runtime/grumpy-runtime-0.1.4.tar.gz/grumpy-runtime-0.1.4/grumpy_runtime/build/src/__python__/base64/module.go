package base64
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/base64.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ß01 := πg.InternStr("01")
		ß2 := πg.InternStr("2")
		ß3 := πg.InternStr("3")
		ß4 := πg.InternStr("4")
		ß5 := πg.InternStr("5")
		ß6 := πg.InternStr("6")
		ß7 := πg.InternStr("7")
		ßA := πg.InternStr("A")
		ßB := πg.InternStr("B")
		ßC := πg.InternStr("C")
		ßD := πg.InternStr("D")
		ßE := πg.InternStr("E")
		ßEMPTYSTRING := πg.InternStr("EMPTYSTRING")
		ßError := πg.InternStr("Error")
		ßF := πg.InternStr("F")
		ßFalse := πg.InternStr("False")
		ßG := πg.InternStr("G")
		ßH := πg.InternStr("H")
		ßI := πg.InternStr("I")
		ßJ := πg.InternStr("J")
		ßK := πg.InternStr("K")
		ßL := πg.InternStr("L")
		ßM := πg.InternStr("M")
		ßMAXBINSIZE := πg.InternStr("MAXBINSIZE")
		ßMAXLINESIZE := πg.InternStr("MAXLINESIZE")
		ßN := πg.InternStr("N")
		ßNone := πg.InternStr("None")
		ßO := πg.InternStr("O")
		ßP := πg.InternStr("P")
		ßQ := πg.InternStr("Q")
		ßR := πg.InternStr("R")
		ßS := πg.InternStr("S")
		ßT := πg.InternStr("T")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßU := πg.InternStr("U")
		ßV := πg.InternStr("V")
		ßW := πg.InternStr("W")
		ßX := πg.InternStr("X")
		ßY := πg.InternStr("Y")
		ßZ := πg.InternStr("Z")
		ß__all__ := πg.InternStr("__all__")
		ß__enter__ := πg.InternStr("__enter__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__main__ := πg.InternStr("__main__")
		ß__name__ := πg.InternStr("__name__")
		ß_b32alphabet := πg.InternStr("_b32alphabet")
		ß_b32rev := πg.InternStr("_b32rev")
		ß_b32tab := πg.InternStr("_b32tab")
		ß_translate := πg.InternStr("_translate")
		ß_translation := πg.InternStr("_translation")
		ß_urlsafe_decode_translation := πg.InternStr("_urlsafe_decode_translation")
		ß_urlsafe_encode_translation := πg.InternStr("_urlsafe_encode_translation")
		ßa2b_base64 := πg.InternStr("a2b_base64")
		ßappend := πg.InternStr("append")
		ßargv := πg.InternStr("argv")
		ßb16decode := πg.InternStr("b16decode")
		ßb16encode := πg.InternStr("b16encode")
		ßb2a_base64 := πg.InternStr("b2a_base64")
		ßb32decode := πg.InternStr("b32decode")
		ßb32encode := πg.InternStr("b32encode")
		ßb64decode := πg.InternStr("b64decode")
		ßb64encode := πg.InternStr("b64encode")
		ßbinascii := πg.InternStr("binascii")
		ßchr := πg.InternStr("chr")
		ßdecode := πg.InternStr("decode")
		ßdecodestring := πg.InternStr("decodestring")
		ßdeut := πg.InternStr("deut")
		ßdict := πg.InternStr("dict")
		ßdivmod := πg.InternStr("divmod")
		ßencode := πg.InternStr("encode")
		ßencodestring := πg.InternStr("encodestring")
		ßerror := πg.InternStr("error")
		ßexit := πg.InternStr("exit")
		ßextend := πg.InternStr("extend")
		ßget := πg.InternStr("get")
		ßgetopt := πg.InternStr("getopt")
		ßgroup := πg.InternStr("group")
		ßhexlify := πg.InternStr("hexlify")
		ßitems := πg.InternStr("items")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßlong := πg.InternStr("long")
		ßmaketrans := πg.InternStr("maketrans")
		ßopen := πg.InternStr("open")
		ßord := πg.InternStr("ord")
		ßpad := πg.InternStr("pad")
		ßrange := πg.InternStr("range")
		ßrb := πg.InternStr("rb")
		ßre := πg.InternStr("re")
		ßread := πg.InternStr("read")
		ßreadline := πg.InternStr("readline")
		ßrepr := πg.InternStr("repr")
		ßsearch := πg.InternStr("search")
		ßsort := πg.InternStr("sort")
		ßstandard_b64decode := πg.InternStr("standard_b64decode")
		ßstandard_b64encode := πg.InternStr("standard_b64encode")
		ßstderr := πg.InternStr("stderr")
		ßstdin := πg.InternStr("stdin")
		ßstdout := πg.InternStr("stdout")
		ßstring := πg.InternStr("string")
		ßstruct := πg.InternStr("struct")
		ßtest := πg.InternStr("test")
		ßtest1 := πg.InternStr("test1")
		ßtranslate := πg.InternStr("translate")
		ßunhexlify := πg.InternStr("unhexlify")
		ßunpack := πg.InternStr("unpack")
		ßupper := πg.InternStr("upper")
		ßurlsafe_b64decode := πg.InternStr("urlsafe_b64decode")
		ßurlsafe_b64encode := πg.InternStr("urlsafe_b64encode")
		ßwrite := πg.InternStr("write")
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
		var πTemp011 *πg.Dict
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
		var πTemp026 bool
		_ = πTemp026
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 3: """RFC 3548: Base16, Base32, Base64 Data Encodings"""
			πF.SetLineno(3)
			// line 8: import re
			πF.SetLineno(8)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 10: import _struct as struct
			πF.SetLineno(10)
			if πTemp002, πE = πg.ImportModule(πF, "_struct"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßstruct.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 11: import string
			πF.SetLineno(11)
			if πTemp002, πE = πg.ImportModule(πF, "string"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßstring.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 12: import binascii
			πF.SetLineno(12)
			if πTemp002, πE = πg.ImportModule(πF, "binascii"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßbinascii.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 15: __all__ = [
			πF.SetLineno(15)
			πTemp002 = make([]*πg.Object, 14)
			πTemp002[0] = ßencode.ToObject()
			πTemp002[1] = ßdecode.ToObject()
			πTemp002[2] = ßencodestring.ToObject()
			πTemp002[3] = ßdecodestring.ToObject()
			πTemp002[4] = ßb64encode.ToObject()
			πTemp002[5] = ßb64decode.ToObject()
			πTemp002[6] = ßb32encode.ToObject()
			πTemp002[7] = ßb32decode.ToObject()
			πTemp002[8] = ßb16encode.ToObject()
			πTemp002[9] = ßb16decode.ToObject()
			πTemp002[10] = ßstandard_b64encode.ToObject()
			πTemp002[11] = ßstandard_b64decode.ToObject()
			πTemp002[12] = ßurlsafe_b64encode.ToObject()
			πTemp002[13] = ßurlsafe_b64decode.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 30: _translation = [chr(_x) for _x in range(256)]
			πF.SetLineno(30)
			πTemp004 = make([]πg.Param, 0)
			πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/base64.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µ_x *πg.Object = πg.UnboundLocal; _ = µ_x
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
						πTemp002[0] = πg.NewInt(256).ToObject()
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
							µ_x = πTemp003
						}
						if πE != nil || !πTemp006 {
							continue
						}
						πF.PushCheckpoint(1)            
						// line 30: _translation = [chr(_x) for _x in range(256)]
						πF.SetLineno(30)
						πTemp002 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µ_x, "_x"); πE != nil {
							continue
						}
						πTemp002[0] = µ_x
						if πTemp003, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
							continue
						}
						if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp002)
						πF.PushCheckpoint(4)
						return πTemp004, nil
					Label4:
						πTemp003 = πSent
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
			if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_translation.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 31: EMPTYSTRING = ''
			πF.SetLineno(31)
			if πE = πF.Globals().SetItem(πF, ßEMPTYSTRING.ToObject(), ß.ToObject()); πE != nil {
				continue
			}
			// line 34: def _translate(s, altchars):
			πF.SetLineno(34)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "s", Def: nil}
			πTemp004[1] = πg.Param{Name: "altchars", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("_translate", "build/src/__python__/base64.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µaltchars *πg.Object = πArgs[1]; _ = µaltchars
				var µtranslation *πg.Object = πg.UnboundLocal; _ = µtranslation
				var µk *πg.Object = πg.UnboundLocal; _ = µk
				var µv *πg.Object = πg.UnboundLocal; _ = µv
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
				var πTemp009 []*πg.Object
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 35: translation = _translation[:]
					πF.SetLineno(35)
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_translation); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
						continue
					}
					µtranslation = πTemp002
					if πE = πg.CheckLocal(πF, µaltchars, "altchars"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µaltchars, ßitems, nil); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
							continue
						}
						µk = πTemp003
						µv = πTemp006
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 37: translation[ord(k)] = v
					πF.SetLineno(37)
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µv); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtranslation, "translation"); πE != nil {
						continue
					}
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
						continue
					}
					πTemp007[0] = µk
					if πTemp006, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp006.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp003 = πTemp008
					if πE = πg.SetItem(πF, µtranslation, πTemp003, πTemp002); πE != nil {
						continue
					}
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 38: return s.translate(''.join(translation))
					πF.SetLineno(38)
					πTemp007 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtranslation, "translation"); πE != nil {
						continue
					}
					πTemp009[0] = µtranslation
					if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp007[0] = πTemp002
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µs, ßtranslate, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
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
			if πE = πF.Globals().SetItem(πF, ß_translate.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 44: def b64encode(s, altchars=None):
			πF.SetLineno(44)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "s", Def: nil}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "altchars", Def: πTemp006}
			πTemp005 = πg.NewFunction(πg.NewCode("b64encode", "build/src/__python__/base64.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µaltchars *πg.Object = πArgs[1]; _ = µaltchars
				var µencoded *πg.Object = πg.UnboundLocal; _ = µencoded
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
				var πTemp006 bool
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
					// line 45: """Encode a string using Base64.
					πF.SetLineno(45)
					// line 55: encoded = binascii.b2a_base64(s)[:-1]
					πF.SetLineno(55)
					if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp002, πg.None}, nil); πE != nil {
						continue
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp003[0] = µs
					if πTemp004, πE = πg.ResolveGlobal(πF, ßbinascii); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßb2a_base64, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
						continue
					}
					µencoded = πTemp002
					if πE = πg.CheckLocal(πF, µaltchars, "altchars"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µaltchars != πTemp002).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label1
					}
					goto Label2
					// line 56: if altchars is not None:
					πF.SetLineno(56)
				Label1:
					// line 57: return encoded.translate(string.maketrans(b'+/', altchars[:2]))
					πF.SetLineno(57)
					πTemp003 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(2)
					πTemp007[0] = πg.NewStr("+/").ToObject()
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(2).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µaltchars, "altchars"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µaltchars, πTemp001); πE != nil {
						continue
					}
					πTemp007[1] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmaketrans, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp003[0] = πTemp001
					if πE = πg.CheckLocal(πF, µencoded, "encoded"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µencoded, ßtranslate, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp002
					continue
					goto Label2
				Label2:
					// line 58: return encoded
					πF.SetLineno(58)
					if πE = πg.CheckLocal(πF, µencoded, "encoded"); πE != nil {
						continue
					}
					πR = µencoded
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßb64encode.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 61: def b64decode(s, altchars=None):
			πF.SetLineno(61)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "s", Def: nil}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "altchars", Def: πTemp007}
			πTemp006 = πg.NewFunction(πg.NewCode("b64decode", "build/src/__python__/base64.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µaltchars *πg.Object = πArgs[1]; _ = µaltchars
				var µmsg *πg.Object = πg.UnboundLocal; _ = µmsg
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
				var πTemp006 *πg.BaseException
				_ = πTemp006
				var πTemp007 *πg.Traceback
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					// line 62: """Decode a Base64 encoded string.
					πF.SetLineno(62)
					if πE = πg.CheckLocal(πF, µaltchars, "altchars"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µaltchars != πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 73: if altchars is not None:
					πF.SetLineno(73)
				Label1:
					// line 74: s = s.translate(string.maketrans(altchars[:2], '+/'))
					πF.SetLineno(74)
					πTemp004 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(2)
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(2).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µaltchars, "altchars"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µaltchars, πTemp001); πE != nil {
						continue
					}
					πTemp005[0] = πTemp002
					πTemp005[1] = πg.NewStr("+/").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmaketrans, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp004[0] = πTemp001
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µs, ßtranslate, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µs = πTemp002
					goto Label2
				Label2:
					// line 75: try:
					πF.SetLineno(75)
					πF.PushCheckpoint(4)
					// line 76: return binascii.a2b_base64(s)
					πF.SetLineno(76)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp004[0] = µs
					if πTemp001, πE = πg.ResolveGlobal(πF, ßbinascii); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßa2b_base64, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πR = πTemp001
					continue
					πF.PopCheckpoint()
					goto Label3
				Label4:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp006, πTemp007 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßbinascii); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßError, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label5
					}
					πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
					continue
					// line 77: except binascii.Error, msg:
					πF.SetLineno(77)
				Label5:
					µmsg = πTemp006.ToObject()
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp004[0] = µmsg
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 79: raise TypeError(msg)
					πF.SetLineno(79)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
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
			if πE = πF.Globals().SetItem(πF, ßb64decode.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 82: def standard_b64encode(s):
			πF.SetLineno(82)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "s", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("standard_b64encode", "build/src/__python__/base64.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
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
					// line 83: """Encode a string using the standard Base64 alphabet.
					πF.SetLineno(83)
					// line 87: return b64encode(s)
					πF.SetLineno(87)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					if πTemp002, πE = πg.ResolveGlobal(πF, ßb64encode); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßstandard_b64encode.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 89: def standard_b64decode(s):
			πF.SetLineno(89)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "s", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("standard_b64decode", "build/src/__python__/base64.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
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
					// line 90: """Decode a string encoded with the standard Base64 alphabet.
					πF.SetLineno(90)
					// line 97: return b64decode(s)
					πF.SetLineno(97)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					if πTemp002, πE = πg.ResolveGlobal(πF, ßb64decode); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßstandard_b64decode.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 99: _urlsafe_encode_translation = string.maketrans(b'+/', b'-_')
			πF.SetLineno(99)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("+/").ToObject()
			πTemp002[1] = πg.NewStr("-_").ToObject()
			if πTemp009, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
				continue
			}
			if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßmaketrans, nil); πE != nil {
				continue
			}
			if πTemp009, πE = πTemp010.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ß_urlsafe_encode_translation.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 100: _urlsafe_decode_translation = string.maketrans(b'-_', b'+/')
			πF.SetLineno(100)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("-_").ToObject()
			πTemp002[1] = πg.NewStr("+/").ToObject()
			if πTemp009, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
				continue
			}
			if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßmaketrans, nil); πE != nil {
				continue
			}
			if πTemp009, πE = πTemp010.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ß_urlsafe_decode_translation.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 102: def urlsafe_b64encode(s):
			πF.SetLineno(102)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "s", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("urlsafe_b64encode", "build/src/__python__/base64.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
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
					// line 103: """Encode a string using the URL- and filesystem-safe Base64 alphabet.
					πF.SetLineno(103)
					// line 108: return b64encode(s).translate(_urlsafe_encode_translation)
					πF.SetLineno(108)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_urlsafe_encode_translation); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp003[0] = µs
					if πTemp002, πE = πg.ResolveGlobal(πF, ßb64encode); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.GetAttr(πF, πTemp004, ßtranslate, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßurlsafe_b64encode.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 110: def urlsafe_b64decode(s):
			πF.SetLineno(110)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "s", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("urlsafe_b64decode", "build/src/__python__/base64.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
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
					// line 111: """Decode a string using the URL- and filesystem-safe Base64 alphabet.
					πF.SetLineno(111)
					// line 120: return b64decode(s.translate(_urlsafe_decode_translation))
					πF.SetLineno(120)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_urlsafe_decode_translation); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µs, ßtranslate, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßb64decode); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßurlsafe_b64decode.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 125: _b32alphabet = {
			πF.SetLineno(125)
			πTemp011 = πg.NewDict()
			if πE = πTemp011.SetItem(πF, πg.NewInt(0).ToObject(), ßA.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(9).ToObject(), ßJ.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(18).ToObject(), ßS.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(27).ToObject(), ß3.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(1).ToObject(), ßB.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(10).ToObject(), ßK.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(19).ToObject(), ßT.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(28).ToObject(), ß4.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(2).ToObject(), ßC.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(11).ToObject(), ßL.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(20).ToObject(), ßU.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(29).ToObject(), ß5.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(3).ToObject(), ßD.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(12).ToObject(), ßM.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(21).ToObject(), ßV.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(30).ToObject(), ß6.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(4).ToObject(), ßE.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(13).ToObject(), ßN.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(22).ToObject(), ßW.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(31).ToObject(), ß7.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(5).ToObject(), ßF.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(14).ToObject(), ßO.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(23).ToObject(), ßX.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(6).ToObject(), ßG.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(15).ToObject(), ßP.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(24).ToObject(), ßY.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(7).ToObject(), ßH.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(16).ToObject(), ßQ.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(25).ToObject(), ßZ.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(8).ToObject(), ßI.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(17).ToObject(), ßR.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, πg.NewInt(26).ToObject(), ß2.ToObject()); πE != nil {
				continue
			}
			πTemp012 = πTemp011.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_b32alphabet.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 137: _b32tab = _b32alphabet.items()
			πF.SetLineno(137)
			if πTemp012, πE = πg.ResolveGlobal(πF, ß_b32alphabet); πE != nil {
				continue
			}
			if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßitems, nil); πE != nil {
				continue
			}
			if πTemp012, πE = πTemp013.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_b32tab.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 138: _b32tab.sort()
			πF.SetLineno(138)
			if πTemp012, πE = πg.ResolveGlobal(πF, ß_b32tab); πE != nil {
				continue
			}
			if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßsort, nil); πE != nil {
				continue
			}
			if πTemp012, πE = πTemp013.Call(πF, nil, nil); πE != nil {
				continue
			}
			// line 139: _b32tab = [v for k, v in _b32tab]
			πF.SetLineno(139)
			πTemp004 = make([]πg.Param, 0)
			πTemp013 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/base64.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µk *πg.Object = πg.UnboundLocal; _ = µk
				var µv *πg.Object = πg.UnboundLocal; _ = µv
				var πTemp001 *πg.Object
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
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 1: goto Label1
						case 2: goto Label2
						case 4: goto Label4
						default: panic("unexpected function state")
						}
						if πTemp002, πE = πg.ResolveGlobal(πF, ß_b32tab); πE != nil {
							continue
						}
						if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
							continue
						}
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
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
								continue
							}
							µk = πTemp005
							µv = πTemp006
						}
						if πE != nil || !πTemp004 {
							continue
						}
						πF.PushCheckpoint(1)            
						// line 139: _b32tab = [v for k, v in _b32tab]
						πF.SetLineno(139)
						if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
							continue
						}
						πF.PushCheckpoint(4)
						return µv, nil
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
			if πTemp014, πE = πTemp013.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ListType.Call(πF, πg.Args{πTemp014}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_b32tab.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 140: _b32rev = dict([(v, long(k)) for k, v in _b32alphabet.items()])
			πF.SetLineno(140)
			πTemp002 = πF.MakeArgs(1)
			πTemp004 = make([]πg.Param, 0)
			πTemp014 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/base64.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µk *πg.Object = πg.UnboundLocal; _ = µk
				var µv *πg.Object = πg.UnboundLocal; _ = µv
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
						if πTemp002, πE = πg.ResolveGlobal(πF, ß_b32alphabet); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßitems, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
								continue
							}
							µk = πTemp003
							µv = πTemp006
						}
						if πE != nil || !πTemp005 {
							continue
						}
						πF.PushCheckpoint(1)            
						// line 140: _b32rev = dict([(v, long(k)) for k, v in _b32alphabet.items()])
						πF.SetLineno(140)
						if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
							continue
						}
						πTemp007 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
							continue
						}
						πTemp007[0] = µk
						if πTemp003, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
							continue
						}
						if πTemp006, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp007)
						πTemp002 = πg.NewTuple2(µv, πTemp006).ToObject()
						πF.PushCheckpoint(4)
						return πTemp002, nil
					Label4:
						πTemp003 = πSent
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
			if πTemp015, πE = πTemp014.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ListType.Call(πF, πg.Args{πTemp015}, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp012
			if πTemp012, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
				continue
			}
			if πTemp015, πE = πTemp012.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ß_b32rev.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 143: def b32encode(s):
			πF.SetLineno(143)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "s", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("b32encode", "build/src/__python__/base64.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µparts *πg.Object = πg.UnboundLocal; _ = µparts
				var µquanta *πg.Object = πg.UnboundLocal; _ = µquanta
				var µleftover *πg.Object = πg.UnboundLocal; _ = µleftover
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µc1 *πg.Object = πg.UnboundLocal; _ = µc1
				var µc2 *πg.Object = πg.UnboundLocal; _ = µc2
				var µc3 *πg.Object = πg.UnboundLocal; _ = µc3
				var µencoded *πg.Object = πg.UnboundLocal; _ = µencoded
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 bool
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
					case 3: goto Label3
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					// line 144: """Encode a string using Base32.
					πF.SetLineno(144)
					// line 148: parts = []
					πF.SetLineno(148)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µparts = πTemp002
					// line 149: quanta, leftover = divmod(len(s), 5)
					πF.SetLineno(149)
					πTemp001 = πF.MakeArgs(2)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp003[0] = µs
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001[0] = πTemp004
					πTemp001[1] = πg.NewInt(5).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp005}}}, πTemp004); πE != nil {
						continue
					}
					µquanta = πTemp002
					µleftover = πTemp005
					if πE = πg.CheckLocal(πF, µleftover, "leftover"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µleftover); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label1
					}
					goto Label2
					// line 151: if leftover:
					πF.SetLineno(151)
				Label1:
					// line 152: s += ('\0' * (5 - leftover))
					πF.SetLineno(152)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µleftover, "leftover"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Sub(πF, πg.NewInt(5).ToObject(), µleftover); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mul(πF, πg.NewStr("\x00").ToObject(), πTemp004); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IAdd(πF, µs, πTemp002); πE != nil {
						continue
					}
					µs = πTemp004
					// line 153: quanta += 1
					πF.SetLineno(153)
					if πE = πg.CheckLocal(πF, µquanta, "quanta"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µquanta, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µquanta = πTemp002
					goto Label2
				Label2:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µquanta, "quanta"); πE != nil {
						continue
					}
					πTemp001[0] = µquanta
					if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
						continue
					}
					πF.PushCheckpoint(4)
					πTemp006 = false
				Label3:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label5
					}
					if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp007 = !isStop
					} else {
						πTemp007 = true
						µi = πTemp004
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(3)            
					// line 160: c1, c2, c3 = struct.unpack('!HHB', s[i*5:(i+1)*5])
					πF.SetLineno(160)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewStr("!HHB").ToObject()
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Mul(πF, µi, πg.NewInt(5).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Mul(πF, πTemp009, πg.NewInt(5).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πTemp005, πTemp008, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µs, πTemp004); πE != nil {
						continue
					}
					πTemp001[1] = πTemp005
					if πTemp004, πE = πg.ResolveGlobal(πF, ßstruct); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßunpack, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}}}, πTemp004); πE != nil {
						continue
					}
					µc1 = πTemp005
					µc2 = πTemp008
					µc3 = πTemp009
					// line 161: c2 += (c1 & 1) << 16 # 17 bits wide
					πF.SetLineno(161)
					if πE = πg.CheckLocal(πF, µc2, "c2"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µc1, "c1"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.And(πF, µc1, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.LShift(πF, πTemp005, πg.NewInt(16).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IAdd(πF, µc2, πTemp004); πE != nil {
						continue
					}
					µc2 = πTemp005
					// line 162: c3 += (c2 & 3) << 8  # 10 bits wide
					πF.SetLineno(162)
					if πE = πg.CheckLocal(πF, µc3, "c3"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µc2, "c2"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.And(πF, µc2, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.LShift(πF, πTemp005, πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IAdd(πF, µc3, πTemp004); πE != nil {
						continue
					}
					µc3 = πTemp005
					// line 163: parts.extend([_b32tab[c1 >> 11],         # bits 1 - 5
					πF.SetLineno(163)
					πTemp001 = πF.MakeArgs(1)
					πTemp003 = make([]*πg.Object, 8)
					if πE = πg.CheckLocal(πF, µc1, "c1"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.RShift(πF, µc1, πg.NewInt(11).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp008, πE = πg.ResolveGlobal(πF, ß_b32tab); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp004); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					if πE = πg.CheckLocal(πF, µc1, "c1"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.RShift(πF, µc1, πg.NewInt(6).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.And(πF, πTemp008, πg.NewInt(31).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp008, πE = πg.ResolveGlobal(πF, ß_b32tab); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp004); πE != nil {
						continue
					}
					πTemp003[1] = πTemp005
					if πE = πg.CheckLocal(πF, µc1, "c1"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.RShift(πF, µc1, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.And(πF, πTemp008, πg.NewInt(31).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp008, πE = πg.ResolveGlobal(πF, ß_b32tab); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp004); πE != nil {
						continue
					}
					πTemp003[2] = πTemp005
					if πE = πg.CheckLocal(πF, µc2, "c2"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.RShift(πF, µc2, πg.NewInt(12).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp008, πE = πg.ResolveGlobal(πF, ß_b32tab); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp004); πE != nil {
						continue
					}
					πTemp003[3] = πTemp005
					if πE = πg.CheckLocal(πF, µc2, "c2"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.RShift(πF, µc2, πg.NewInt(7).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.And(πF, πTemp008, πg.NewInt(31).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp008, πE = πg.ResolveGlobal(πF, ß_b32tab); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp004); πE != nil {
						continue
					}
					πTemp003[4] = πTemp005
					if πE = πg.CheckLocal(πF, µc2, "c2"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.RShift(πF, µc2, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.And(πF, πTemp008, πg.NewInt(31).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp008, πE = πg.ResolveGlobal(πF, ß_b32tab); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp004); πE != nil {
						continue
					}
					πTemp003[5] = πTemp005
					if πE = πg.CheckLocal(πF, µc3, "c3"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.RShift(πF, µc3, πg.NewInt(5).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp008, πE = πg.ResolveGlobal(πF, ß_b32tab); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp004); πE != nil {
						continue
					}
					πTemp003[6] = πTemp005
					if πE = πg.CheckLocal(πF, µc3, "c3"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.And(πF, µc3, πg.NewInt(31).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp008, πE = πg.ResolveGlobal(πF, ß_b32tab); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp004); πE != nil {
						continue
					}
					πTemp003[7] = πTemp005
					πTemp004 = πg.NewList(πTemp003...).ToObject()
					πTemp001[0] = πTemp004
					if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µparts, ßextend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label4:
					if πE != nil || πR != nil {
						continue
					}
				Label5:
					// line 172: encoded = EMPTYSTRING.join(parts)
					πF.SetLineno(172)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
						continue
					}
					πTemp001[0] = µparts
					if πTemp002, πE = πg.ResolveGlobal(πF, ßEMPTYSTRING); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µencoded = πTemp002
					if πE = πg.CheckLocal(πF, µleftover, "leftover"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µleftover, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label6
					}
					if πE = πg.CheckLocal(πF, µleftover, "leftover"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µleftover, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label7
					}
					if πE = πg.CheckLocal(πF, µleftover, "leftover"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µleftover, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label8
					}
					if πE = πg.CheckLocal(πF, µleftover, "leftover"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µleftover, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label9
					}
					goto Label10
					// line 174: if leftover == 1:
					πF.SetLineno(174)
				Label6:
					// line 175: return encoded[:-6] + '======'
					πF.SetLineno(175)
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(6).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp005, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µencoded, "encoded"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µencoded, πTemp004); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp005, πg.NewStr("======").ToObject()); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					goto Label10
					// line 176: elif leftover == 2:
					πF.SetLineno(176)
				Label7:
					// line 177: return encoded[:-4] + '===='
					πF.SetLineno(177)
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp005, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µencoded, "encoded"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µencoded, πTemp004); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp005, πg.NewStr("====").ToObject()); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					goto Label10
					// line 178: elif leftover == 3:
					πF.SetLineno(178)
				Label8:
					// line 179: return encoded[:-3] + '==='
					πF.SetLineno(179)
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp005, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µencoded, "encoded"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µencoded, πTemp004); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp005, πg.NewStr("===").ToObject()); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					goto Label10
					// line 180: elif leftover == 4:
					πF.SetLineno(180)
				Label9:
					// line 181: return encoded[:-1] + '='
					πF.SetLineno(181)
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp005, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µencoded, "encoded"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µencoded, πTemp004); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp005, πg.NewStr("=").ToObject()); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					goto Label10
				Label10:
					// line 182: return encoded
					πF.SetLineno(182)
					if πE = πg.CheckLocal(πF, µencoded, "encoded"); πE != nil {
						continue
					}
					πR = µencoded
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßb32encode.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 185: def b32decode(s, casefold=False, map01=None):
			πF.SetLineno(185)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "s", Def: nil}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "casefold", Def: πTemp016}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "map01", Def: πTemp016}
			πTemp015 = πg.NewFunction(πg.NewCode("b32decode", "build/src/__python__/base64.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µcasefold *πg.Object = πArgs[1]; _ = µcasefold
				var µmap01 *πg.Object = πArgs[2]; _ = µmap01
				var µquanta *πg.Object = πg.UnboundLocal; _ = µquanta
				var µleftover *πg.Object = πg.UnboundLocal; _ = µleftover
				var µpadchars *πg.Object = πg.UnboundLocal; _ = µpadchars
				var µmo *πg.Object = πg.UnboundLocal; _ = µmo
				var µparts *πg.Object = πg.UnboundLocal; _ = µparts
				var µacc *πg.Object = πg.UnboundLocal; _ = µacc
				var µshift *πg.Object = πg.UnboundLocal; _ = µshift
				var µc *πg.Object = πg.UnboundLocal; _ = µc
				var µval *πg.Object = πg.UnboundLocal; _ = µval
				var µlast *πg.Object = πg.UnboundLocal; _ = µlast
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
				var πTemp006 bool
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
					case 11: goto Label11
					case 12: goto Label12
					default: panic("unexpected function state")
					}
					// line 186: """Decode a Base32 encoded string.
					πF.SetLineno(186)
					// line 204: quanta, leftover = divmod(len(s), 8)
					πF.SetLineno(204)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp002[0] = µs
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					πTemp001[1] = πg.NewInt(8).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}}}, πTemp004); πE != nil {
						continue
					}
					µquanta = πTemp003
					µleftover = πTemp005
					if πE = πg.CheckLocal(πF, µleftover, "leftover"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µleftover); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label1
					}
					goto Label2
					// line 205: if leftover:
					πF.SetLineno(205)
				Label1:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("Incorrect padding").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 206: raise TypeError('Incorrect padding')
					πF.SetLineno(206)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µmap01, "map01"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µmap01); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					goto Label4
					// line 210: if map01:
					πF.SetLineno(210)
				Label3:
					// line 211: s = s.translate(string.maketrans(b'01', b'O' + map01))
					πF.SetLineno(211)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = ß01.ToObject()
					if πE = πg.CheckLocal(πF, µmap01, "map01"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, ßO.ToObject(), µmap01); πE != nil {
						continue
					}
					πTemp002[1] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßmaketrans, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µs, ßtranslate, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µs = πTemp004
					goto Label4
				Label4:
					if πE = πg.CheckLocal(πF, µcasefold, "casefold"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µcasefold); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label5
					}
					goto Label6
					// line 212: if casefold:
					πF.SetLineno(212)
				Label5:
					// line 213: s = s.upper()
					πF.SetLineno(213)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µs, ßupper, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					µs = πTemp004
					goto Label6
				Label6:
					// line 217: padchars = 0
					πF.SetLineno(217)
					µpadchars = πg.NewInt(0).ToObject()
					// line 218: mo = re.search('(?P<pad>[=]*)$', s)
					πF.SetLineno(218)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewStr("(?P<pad>[=]*)$").ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[1] = µs
					if πTemp003, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsearch, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µmo = πTemp003
					if πE = πg.CheckLocal(πF, µmo, "mo"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µmo); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label7
					}
					goto Label8
					// line 219: if mo:
					πF.SetLineno(219)
				Label7:
					// line 220: padchars = len(mo.group('pad'))
					πF.SetLineno(220)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = ßpad.ToObject()
					if πE = πg.CheckLocal(πF, µmo, "mo"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µmo, ßgroup, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µpadchars = πTemp004
					if πE = πg.CheckLocal(πF, µpadchars, "padchars"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GT(πF, µpadchars, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label9
					}
					goto Label10
					// line 221: if padchars > 0:
					πF.SetLineno(221)
				Label9:
					// line 222: s = s[:-padchars]
					πF.SetLineno(222)
					if πE = πg.CheckLocal(πF, µpadchars, "padchars"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Neg(πF, µpadchars); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp004, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µs, πTemp003); πE != nil {
						continue
					}
					µs = πTemp004
					goto Label10
				Label10:
					goto Label8
				Label8:
					// line 224: parts = []
					πF.SetLineno(224)
					πTemp001 = make([]*πg.Object, 0)
					πTemp003 = πg.NewList(πTemp001...).ToObject()
					µparts = πTemp003
					// line 225: acc = 0
					πF.SetLineno(225)
					µacc = πg.NewInt(0).ToObject()
					// line 226: shift = 35
					πF.SetLineno(226)
					µshift = πg.NewInt(35).ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Iter(πF, µs); πE != nil {
						continue
					}
					πF.PushCheckpoint(12)
					πTemp006 = false
				Label11:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label13
					}
					if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp007 = !isStop
					} else {
						πTemp007 = true
						µc = πTemp004
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(11)            
					// line 228: val = _b32rev.get(c)
					πF.SetLineno(228)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp001[0] = µc
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_b32rev); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßget, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µval = πTemp004
					if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(µval == πTemp005).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label14
					}
					goto Label15
					// line 229: if val is None:
					πF.SetLineno(229)
				Label14:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("Non-base32 digit found").ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 230: raise TypeError('Non-base32 digit found')
					πF.SetLineno(230)
					πE = πF.Raise(πTemp005, nil, nil)
					continue
					goto Label15
				Label15:
					// line 231: acc += _b32rev[c] << shift
					πF.SetLineno(231)
					if πE = πg.CheckLocal(πF, µacc, "acc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp005 = µc
					if πTemp009, πE = πg.ResolveGlobal(πF, ß_b32rev); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µshift, "shift"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.LShift(πF, πTemp008, µshift); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IAdd(πF, µacc, πTemp004); πE != nil {
						continue
					}
					µacc = πTemp005
					// line 232: shift -= 5
					πF.SetLineno(232)
					if πE = πg.CheckLocal(πF, µshift, "shift"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ISub(πF, µshift, πg.NewInt(5).ToObject()); πE != nil {
						continue
					}
					µshift = πTemp004
					if πE = πg.CheckLocal(πF, µshift, "shift"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.LT(πF, µshift, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label16
					}
					goto Label17
					// line 233: if shift < 0:
					πF.SetLineno(233)
				Label16:
					// line 234: parts.append(binascii.unhexlify('%010x' % acc))
					πF.SetLineno(234)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µacc, "acc"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mod(πF, πg.NewStr("%010x").ToObject(), µacc); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßbinascii); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßunhexlify, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µparts, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 235: acc = 0
					πF.SetLineno(235)
					µacc = πg.NewInt(0).ToObject()
					// line 236: shift = 35
					πF.SetLineno(236)
					µshift = πg.NewInt(35).ToObject()
					goto Label17
				Label17:
					continue
				Label12:
					if πE != nil || πR != nil {
						continue
					}
				Label13:
					// line 238: last = binascii.unhexlify('%010x' % acc)
					πF.SetLineno(238)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µacc, "acc"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("%010x").ToObject(), µacc); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßbinascii); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßunhexlify, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µlast = πTemp003
					if πE = πg.CheckLocal(πF, µpadchars, "padchars"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µpadchars, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label18
					}
					if πE = πg.CheckLocal(πF, µpadchars, "padchars"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µpadchars, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label19
					}
					if πE = πg.CheckLocal(πF, µpadchars, "padchars"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µpadchars, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label20
					}
					if πE = πg.CheckLocal(πF, µpadchars, "padchars"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µpadchars, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label21
					}
					if πE = πg.CheckLocal(πF, µpadchars, "padchars"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µpadchars, πg.NewInt(6).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label22
					}
					goto Label23
					// line 239: if padchars == 0:
					πF.SetLineno(239)
				Label18:
					// line 240: last = ''                       # No characters
					πF.SetLineno(240)
					µlast = ß.ToObject()
					goto Label24
					// line 241: elif padchars == 1:
					πF.SetLineno(241)
				Label19:
					// line 242: last = last[:-1]
					πF.SetLineno(242)
					if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp004, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlast, "last"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µlast, πTemp003); πE != nil {
						continue
					}
					µlast = πTemp004
					goto Label24
					// line 243: elif padchars == 3:
					πF.SetLineno(243)
				Label20:
					// line 244: last = last[:-2]
					πF.SetLineno(244)
					if πTemp004, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp004, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlast, "last"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µlast, πTemp003); πE != nil {
						continue
					}
					µlast = πTemp004
					goto Label24
					// line 245: elif padchars == 4:
					πF.SetLineno(245)
				Label21:
					// line 246: last = last[:-3]
					πF.SetLineno(246)
					if πTemp004, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp004, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlast, "last"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µlast, πTemp003); πE != nil {
						continue
					}
					µlast = πTemp004
					goto Label24
					// line 247: elif padchars == 6:
					πF.SetLineno(247)
				Label22:
					// line 248: last = last[:-4]
					πF.SetLineno(248)
					if πTemp004, πE = πg.Neg(πF, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp004, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlast, "last"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µlast, πTemp003); πE != nil {
						continue
					}
					µlast = πTemp004
					goto Label24
				Label23:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("Incorrect padding").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 250: raise TypeError('Incorrect padding')
					πF.SetLineno(250)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					goto Label24
				Label24:
					// line 251: parts.append(last)
					πF.SetLineno(251)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlast, "last"); πE != nil {
						continue
					}
					πTemp001[0] = µlast
					if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µparts, ßappend, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 252: return EMPTYSTRING.join(parts)
					πF.SetLineno(252)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
						continue
					}
					πTemp001[0] = µparts
					if πTemp003, πE = πg.ResolveGlobal(πF, ßEMPTYSTRING); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßb32decode.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 259: def b16encode(s):
			πF.SetLineno(259)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "s", Def: nil}
			πTemp016 = πg.NewFunction(πg.NewCode("b16encode", "build/src/__python__/base64.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
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
					// line 260: """Encode a string using Base16.
					πF.SetLineno(260)
					// line 264: return binascii.hexlify(s).upper()
					πF.SetLineno(264)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					if πTemp002, πE = πg.ResolveGlobal(πF, ßbinascii); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßhexlify, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßupper, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßb16encode.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 267: def b16decode(s, casefold=False):
			πF.SetLineno(267)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "s", Def: nil}
			if πTemp018, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "casefold", Def: πTemp018}
			πTemp017 = πg.NewFunction(πg.NewCode("b16decode", "build/src/__python__/base64.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µcasefold *πg.Object = πArgs[1]; _ = µcasefold
				var πTemp001 bool
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
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
					// line 268: """Decode a Base16 encoded string.
					πF.SetLineno(268)
					if πE = πg.CheckLocal(πF, µcasefold, "casefold"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µcasefold); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label1
					}
					goto Label2
					// line 278: if casefold:
					πF.SetLineno(278)
				Label1:
					// line 279: s = s.upper()
					πF.SetLineno(279)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µs, ßupper, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µs = πTemp003
					goto Label2
				Label2:
					πTemp004 = πF.MakeArgs(2)
					πTemp004[0] = πg.NewStr("[^0-9A-F]").ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp004[1] = µs
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsearch, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp001, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label3
					}
					goto Label4
					// line 280: if re.search('[^0-9A-F]', s):
					πF.SetLineno(280)
				Label3:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("Non-base16 digit found").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 281: raise TypeError('Non-base16 digit found')
					πF.SetLineno(281)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label4
				Label4:
					// line 282: return binascii.unhexlify(s)
					πF.SetLineno(282)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp004[0] = µs
					if πTemp002, πE = πg.ResolveGlobal(πF, ßbinascii); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunhexlify, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
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
			if πE = πF.Globals().SetItem(πF, ßb16decode.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 290: MAXLINESIZE = 76 # Excluding the CRLF
			πF.SetLineno(290)
			if πE = πF.Globals().SetItem(πF, ßMAXLINESIZE.ToObject(), πg.NewInt(76).ToObject()); πE != nil {
				continue
			}
			// line 291: MAXBINSIZE = (MAXLINESIZE//4)*3
			πF.SetLineno(291)
			if πTemp020, πE = πg.ResolveGlobal(πF, ßMAXLINESIZE); πE != nil {
				continue
			}
			if πTemp019, πE = πg.FloorDiv(πF, πTemp020, πg.NewInt(4).ToObject()); πE != nil {
				continue
			}
			if πTemp018, πE = πg.Mul(πF, πTemp019, πg.NewInt(3).ToObject()); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMAXBINSIZE.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 293: def encode(input, output):
			πF.SetLineno(293)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "input", Def: nil}
			πTemp004[1] = πg.Param{Name: "output", Def: nil}
			πTemp018 = πg.NewFunction(πg.NewCode("encode", "build/src/__python__/base64.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µinput *πg.Object = πArgs[0]; _ = µinput
				var µoutput *πg.Object = πArgs[1]; _ = µoutput
				var µs *πg.Object = πg.UnboundLocal; _ = µs
				var µns *πg.Object = πg.UnboundLocal; _ = µns
				var µline *πg.Object = πg.UnboundLocal; _ = µline
				var πTemp001 bool
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					case 6: goto Label6
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					// line 294: """Encode a file."""
					πF.SetLineno(294)
					// line 295: while True:
					πF.SetLineno(295)
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
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πE != nil || !πTemp002 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 296: s = input.read(MAXBINSIZE)
					πF.SetLineno(296)
					πTemp004 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßMAXBINSIZE); πE != nil {
						continue
					}
					πTemp004[0] = πTemp003
					if πE = πg.CheckLocal(πF, µinput, "input"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µinput, ßread, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µs = πTemp005
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µs); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					goto Label5
					// line 297: if not s:
					πF.SetLineno(297)
				Label4:
					// line 298: break
					πF.SetLineno(298)
					πTemp001 = true
					continue
					goto Label5
				Label5:
					// line 299: while len(s) < MAXBINSIZE:
					πF.SetLineno(299)
					πF.PushCheckpoint(7)
					πTemp002 = false
				Label6:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
						πF.PopCheckpoint()
						goto Label8
					}
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp004[0] = µs
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßMAXBINSIZE); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LT(πF, πTemp007, πTemp005); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(6)            
					// line 300: ns = input.read(MAXBINSIZE-len(s))
					πF.SetLineno(300)
					πTemp004 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßMAXBINSIZE); πE != nil {
						continue
					}
					πTemp008 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp008[0] = µs
					if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp007.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πTemp003, πE = πg.Sub(πF, πTemp005, πTemp009); πE != nil {
						continue
					}
					πTemp004[0] = πTemp003
					if πE = πg.CheckLocal(πF, µinput, "input"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µinput, ßread, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µns = πTemp005
					if πE = πg.CheckLocal(πF, µns, "ns"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µns); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label9
					}
					goto Label10
					// line 301: if not ns:
					πF.SetLineno(301)
				Label9:
					// line 302: break
					πF.SetLineno(302)
					πTemp002 = true
					continue
					goto Label10
				Label10:
					// line 303: s += ns
					πF.SetLineno(303)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µns, "ns"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IAdd(πF, µs, µns); πE != nil {
						continue
					}
					µs = πTemp003
					continue
				Label7:
					if πE != nil || πR != nil {
						continue
					}
				Label8:
					// line 304: line = binascii.b2a_base64(s)
					πF.SetLineno(304)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp004[0] = µs
					if πTemp003, πE = πg.ResolveGlobal(πF, ßbinascii); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßb2a_base64, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µline = πTemp003
					// line 305: output.write(line)
					πF.SetLineno(305)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp004[0] = µline
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µoutput, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
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
			if πE = πF.Globals().SetItem(πF, ßencode.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 308: def decode(input, output):
			πF.SetLineno(308)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "input", Def: nil}
			πTemp004[1] = πg.Param{Name: "output", Def: nil}
			πTemp019 = πg.NewFunction(πg.NewCode("decode", "build/src/__python__/base64.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µinput *πg.Object = πArgs[0]; _ = µinput
				var µoutput *πg.Object = πArgs[1]; _ = µoutput
				var µline *πg.Object = πg.UnboundLocal; _ = µline
				var µs *πg.Object = πg.UnboundLocal; _ = µs
				var πTemp001 bool
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
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
					// line 309: """Decode a file."""
					πF.SetLineno(309)
					// line 310: while True:
					πF.SetLineno(310)
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
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πE != nil || !πTemp002 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 311: line = input.readline()
					πF.SetLineno(311)
					if πE = πg.CheckLocal(πF, µinput, "input"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µinput, ßreadline, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					µline = πTemp004
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µline); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					goto Label5
					// line 312: if not line:
					πF.SetLineno(312)
				Label4:
					// line 313: break
					πF.SetLineno(313)
					πTemp001 = true
					continue
					goto Label5
				Label5:
					// line 314: s = binascii.a2b_base64(line)
					πF.SetLineno(314)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp005[0] = µline
					if πTemp003, πE = πg.ResolveGlobal(πF, ßbinascii); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßa2b_base64, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µs = πTemp003
					// line 315: output.write(s)
					πF.SetLineno(315)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp005[0] = µs
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µoutput, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
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
			if πE = πF.Globals().SetItem(πF, ßdecode.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 318: def encodestring(s):
			πF.SetLineno(318)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "s", Def: nil}
			πTemp020 = πg.NewFunction(πg.NewCode("encodestring", "build/src/__python__/base64.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µpieces *πg.Object = πg.UnboundLocal; _ = µpieces
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µchunk *πg.Object = πg.UnboundLocal; _ = µchunk
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 319: """Encode a string into multiple lines of base-64 data."""
					πF.SetLineno(319)
					// line 320: pieces = []
					πF.SetLineno(320)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µpieces = πTemp002
					πTemp001 = πF.MakeArgs(3)
					πTemp001[0] = πg.NewInt(0).ToObject()
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp003[0] = µs
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001[1] = πTemp005
					if πTemp004, πE = πg.ResolveGlobal(πF, ßMAXBINSIZE); πE != nil {
						continue
					}
					πTemp001[2] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp006 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label3
					}
					if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp007 = !isStop
					} else {
						πTemp007 = true
						µi = πTemp004
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 322: chunk = s[i : i + MAXBINSIZE]
					πF.SetLineno(322)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveGlobal(πF, ßMAXBINSIZE); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Add(πF, µi, πTemp008); πE != nil {
						continue
					}
					if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{µi, πTemp005, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µs, πTemp004); πE != nil {
						continue
					}
					µchunk = πTemp005
					// line 323: pieces.append(binascii.b2a_base64(chunk))
					πF.SetLineno(323)
					πTemp001 = πF.MakeArgs(1)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchunk, "chunk"); πE != nil {
						continue
					}
					πTemp003[0] = µchunk
					if πTemp004, πE = πg.ResolveGlobal(πF, ßbinascii); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßb2a_base64, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001[0] = πTemp004
					if πE = πg.CheckLocal(πF, µpieces, "pieces"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µpieces, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 324: return "".join(pieces)
					πF.SetLineno(324)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpieces, "pieces"); πE != nil {
						continue
					}
					πTemp001[0] = µpieces
					if πTemp002, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßencodestring.ToObject(), πTemp020); πE != nil {
				continue
			}
			// line 327: def decodestring(s):
			πF.SetLineno(327)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "s", Def: nil}
			πTemp021 = πg.NewFunction(πg.NewCode("decodestring", "build/src/__python__/base64.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
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
					// line 328: """Decode a string."""
					πF.SetLineno(328)
					// line 329: return binascii.a2b_base64(s)
					πF.SetLineno(329)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					if πTemp002, πE = πg.ResolveGlobal(πF, ßbinascii); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßa2b_base64, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßdecodestring.ToObject(), πTemp021); πE != nil {
				continue
			}
			// line 334: def test():
			πF.SetLineno(334)
			πTemp004 = make([]πg.Param, 0)
			πTemp022 = πg.NewFunction(πg.NewCode("test", "build/src/__python__/base64.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsys *πg.Object = πg.UnboundLocal; _ = µsys
				var µgetopt *πg.Object = πg.UnboundLocal; _ = µgetopt
				var µopts *πg.Object = πg.UnboundLocal; _ = µopts
				var µargs *πg.Object = πg.UnboundLocal; _ = µargs
				var µmsg *πg.Object = πg.UnboundLocal; _ = µmsg
				var µfunc *πg.Object = πg.UnboundLocal; _ = µfunc
				var µo *πg.Object = πg.UnboundLocal; _ = µo
				var µa *πg.Object = πg.UnboundLocal; _ = µa
				var µf *πg.Object = πg.UnboundLocal; _ = µf
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.BaseException
				_ = πTemp005
				var πTemp006 *πg.Traceback
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πTemp010 *πg.Type
				_ = πTemp010
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					case 19: goto Label19
					case 4: goto Label4
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 335: """Small test program"""
					πF.SetLineno(335)
					// line 336: import sys, getopt
					πF.SetLineno(336)
					if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					µsys = πTemp001
					if πTemp002, πE = πg.ImportModule(πF, "getopt"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					µgetopt = πTemp001
					// line 337: try:
					πF.SetLineno(337)
					πF.PushCheckpoint(2)
					// line 338: opts, args = getopt.getopt(sys.argv[1:], 'deut')
					πF.SetLineno(338)
					πTemp002 = πF.MakeArgs(2)
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µsys, ßargv, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					πTemp002[1] = ßdeut.ToObject()
					if πE = πg.CheckLocal(πF, µgetopt, "getopt"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µgetopt, ßgetopt, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µopts = πTemp001
					µargs = πTemp004
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp005, πTemp006 = πF.ExcInfo()
					if πE = πg.CheckLocal(πF, µgetopt, "getopt"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µgetopt, ßerror, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label3
					}
					πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
					continue
					// line 339: except getopt.error, msg:
					πF.SetLineno(339)
				Label3:
					µmsg = πTemp005.ToObject()
					// line 340: sys.stdout = sys.stderr
					πF.SetLineno(340)
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µsys, ßstderr, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µsys, ßstdout, πTemp003); πE != nil {
						continue
					}
					// line 341: print msg
					πF.SetLineno(341)
					πTemp002 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp002[0] = µmsg
					if πE = πg.Print(πF, πTemp002, true); πE != nil {
						continue
					}
					// line 342: print """usage: %s [-d|-e|-u|-t] [file|-]
					πF.SetLineno(342)
					πTemp002 = make([]*πg.Object, 1)
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, µsys, ßargv, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp008, πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("usage: %s [-d|-e|-u|-t] [file|-]\n        -d, -u: decode\n        -e: encode (default)\n        -t: encode and decode string 'Aladdin:open sesame'").ToObject(), πTemp004); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πE = πg.Print(πF, πTemp002, true); πE != nil {
						continue
					}
					// line 346: sys.exit(2)
					πF.SetLineno(346)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µsys, ßexit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 347: func = encode
					πF.SetLineno(347)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßencode); πE != nil {
						continue
					}
					µfunc = πTemp001
					if πE = πg.CheckLocal(πF, µopts, "opts"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µopts); πE != nil {
						continue
					}
					πF.PushCheckpoint(5)
					πTemp007 = false
				Label4:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp007 {
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
						πTemp009 = !isStop
					} else {
						πTemp009 = true
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp008}}}, πTemp003); πE != nil {
							continue
						}
						µo = πTemp004
						µa = πTemp008
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(4)            
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µo, πg.NewStr("-e").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label7
					}
					goto Label8
					// line 349: if o == '-e': func = encode
					πF.SetLineno(349)
				Label7:
					// line 349: if o == '-e': func = encode
					πF.SetLineno(349)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßencode); πE != nil {
						continue
					}
					µfunc = πTemp003
					goto Label8
				Label8:
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µo, πg.NewStr("-d").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label9
					}
					goto Label10
					// line 350: if o == '-d': func = decode
					πF.SetLineno(350)
				Label9:
					// line 350: if o == '-d': func = decode
					πF.SetLineno(350)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßdecode); πE != nil {
						continue
					}
					µfunc = πTemp003
					goto Label10
				Label10:
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µo, πg.NewStr("-u").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label11
					}
					goto Label12
					// line 351: if o == '-u': func = decode
					πF.SetLineno(351)
				Label11:
					// line 351: if o == '-u': func = decode
					πF.SetLineno(351)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßdecode); πE != nil {
						continue
					}
					µfunc = πTemp003
					goto Label12
				Label12:
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µo, πg.NewStr("-t").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label13
					}
					goto Label14
					// line 352: if o == '-t': test1(); return
					πF.SetLineno(352)
				Label13:
					// line 352: if o == '-t': test1(); return
					πF.SetLineno(352)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtest1); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 352: if o == '-t': test1(); return
					πF.SetLineno(352)
					πR = πg.None
					continue
					goto Label14
				Label14:
					continue
				Label5:
					if πE != nil || πR != nil {
						continue
					}
				Label6:
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp001 = µargs
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label15
					}
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µargs, πTemp004); πE != nil {
						continue
					}
					if πTemp003, πE = πg.NE(πF, πTemp008, πg.NewStr("-").ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
				Label15:
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label16
					}
					goto Label17
					// line 353: if args and args[0] != '-':
					πF.SetLineno(353)
				Label16:
					// line 354: with open(args[0], 'rb') as f:
					πF.SetLineno(354)
					πTemp002 = πF.MakeArgs(2)
					πTemp001 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					πTemp002[1] = ßrb.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__exit__, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__enter__, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp004.Call(πF, πg.Args{πTemp003}, nil); πE != nil {
						continue
					}
					πF.PushCheckpoint(19)
					µf = πTemp004
					// line 355: func(f, sys.stdout)
					πF.SetLineno(355)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					πTemp002[0] = µf
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, µsys, ßstdout, nil); πE != nil {
						continue
					}
					πTemp002[1] = πTemp008
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp008, πE = µfunc.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πF.PopCheckpoint()
				Label19:
					πTemp005, πTemp006 = nil, nil
					if πE != nil {
						πTemp005, πTemp006 = πF.ExcInfo()
					}
					if πTemp005 != nil {
						πTemp010 = πTemp005.Type()
						if πTemp008, πE = πTemp001.Call(πF, πg.Args{πTemp003, πTemp010.ToObject(), πTemp005.ToObject(), πTemp006.ToObject()}, nil); πE != nil {
							continue
						}
					} else {
						if πTemp008, πE = πTemp001.Call(πF, πg.Args{πTemp003, πg.None, πg.None, πg.None}, nil); πE != nil {
							continue
						}
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp008); πE != nil {
						continue
					}
					if πTemp005 != nil && πTemp007 != true {
						πE = πF.Raise(nil, nil, nil)
						continue
					}
					if πR != nil {
						continue
					}
					goto Label18
				Label17:
					// line 357: func(sys.stdin, sys.stdout)
					πF.SetLineno(357)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µsys, ßstdin, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µsys, ßstdout, nil); πE != nil {
						continue
					}
					πTemp002[1] = πTemp001
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp001, πE = µfunc.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label18
				Label18:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßtest.ToObject(), πTemp022); πE != nil {
				continue
			}
			// line 360: def test1():
			πF.SetLineno(360)
			πTemp004 = make([]πg.Param, 0)
			πTemp023 = πg.NewFunction(πg.NewCode("test1", "build/src/__python__/base64.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs0 *πg.Object = πg.UnboundLocal; _ = µs0
				var µs1 *πg.Object = πg.UnboundLocal; _ = µs1
				var µs2 *πg.Object = πg.UnboundLocal; _ = µs2
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
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
					// line 361: s0 = "Aladdin:open sesame"
					πF.SetLineno(361)
					µs0 = πg.NewStr("Aladdin:open sesame").ToObject()
					// line 362: s1 = encodestring(s0)
					πF.SetLineno(362)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs0, "s0"); πE != nil {
						continue
					}
					πTemp001[0] = µs0
					if πTemp002, πE = πg.ResolveGlobal(πF, ßencodestring); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µs1 = πTemp003
					// line 363: s2 = decodestring(s1)
					πF.SetLineno(363)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
						continue
					}
					πTemp001[0] = µs1
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdecodestring); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µs2 = πTemp003
					// line 364: print s0, repr(s1), s2
					πF.SetLineno(364)
					πTemp001 = make([]*πg.Object, 3)
					if πE = πg.CheckLocal(πF, µs0, "s0"); πE != nil {
						continue
					}
					πTemp001[0] = µs0
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
						continue
					}
					πTemp004[0] = µs1
					if πTemp002, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp001[1] = πTemp003
					if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
						continue
					}
					πTemp001[2] = µs2
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßtest1.ToObject(), πTemp023); πE != nil {
				continue
			}
			if πTemp025, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp024, πE = πg.Eq(πF, πTemp025, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp026, πE = πg.IsTrue(πF, πTemp024); πE != nil {
				continue
			}
			if πTemp026 {
				goto Label1
			}
			goto Label2
			// line 367: if __name__ == '__main__':
			πF.SetLineno(367)
		Label1:
			// line 368: test()
			πF.SetLineno(368)
			if πTemp024, πE = πg.ResolveGlobal(πF, ßtest); πE != nil {
				continue
			}
			if πTemp025, πE = πTemp024.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("base64", Code)
}
