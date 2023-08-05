package quopri
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/quopri.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ß0 := πg.InternStr("0")
		ß0123456789ABCDEF := πg.InternStr("0123456789ABCDEF")
		ß9 := πg.InternStr("9")
		ßA := πg.InternStr("A")
		ßEMPTYSTRING := πg.InternStr("EMPTYSTRING")
		ßESCAPE := πg.InternStr("ESCAPE")
		ßF := πg.InternStr("F")
		ßHEX := πg.InternStr("HEX")
		ßIOError := πg.InternStr("IOError")
		ßImportError := πg.InternStr("ImportError")
		ßMAXLINESIZE := πg.InternStr("MAXLINESIZE")
		ßNone := πg.InternStr("None")
		ßStringIO := πg.InternStr("StringIO")
		ß_ := πg.InternStr("_")
		ß__all__ := πg.InternStr("__all__")
		ß__main__ := πg.InternStr("__main__")
		ß__name__ := πg.InternStr("__name__")
		ßa := πg.InternStr("a")
		ßa2b_qp := πg.InternStr("a2b_qp")
		ßappend := πg.InternStr("append")
		ßargv := πg.InternStr("argv")
		ßb2a_qp := πg.InternStr("b2a_qp")
		ßchr := πg.InternStr("chr")
		ßclose := πg.InternStr("close")
		ßdecode := πg.InternStr("decode")
		ßdecodestring := πg.InternStr("decodestring")
		ßencode := πg.InternStr("encode")
		ßencodestring := πg.InternStr("encodestring")
		ßerror := πg.InternStr("error")
		ßexit := πg.InternStr("exit")
		ßf := πg.InternStr("f")
		ßgetopt := πg.InternStr("getopt")
		ßgetvalue := πg.InternStr("getvalue")
		ßishex := πg.InternStr("ishex")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßmain := πg.InternStr("main")
		ßneedsquoting := πg.InternStr("needsquoting")
		ßopen := πg.InternStr("open")
		ßord := πg.InternStr("ord")
		ßquote := πg.InternStr("quote")
		ßread := πg.InternStr("read")
		ßreadline := πg.InternStr("readline")
		ßstderr := πg.InternStr("stderr")
		ßstdin := πg.InternStr("stdin")
		ßstdout := πg.InternStr("stdout")
		ßtd := πg.InternStr("td")
		ßunhex := πg.InternStr("unhex")
		ßwrite := πg.InternStr("write")
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
		var πTemp007 []πg.Param
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 2: goto Label2
			default: panic("unexpected function state")
			}
			// line 3: """Conversions to/from quoted-printable transport encoding as per RFC 1521."""
			πF.SetLineno(3)
			// line 7: __all__ = ["encode", "decode", "encodestring", "decodestring"]
			πF.SetLineno(7)
			πTemp001 = make([]*πg.Object, 4)
			πTemp001[0] = ßencode.ToObject()
			πTemp001[1] = ßdecode.ToObject()
			πTemp001[2] = ßencodestring.ToObject()
			πTemp001[3] = ßdecodestring.ToObject()
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 9: ESCAPE = '='
			πF.SetLineno(9)
			if πE = πF.Globals().SetItem(πF, ßESCAPE.ToObject(), πg.NewStr("=").ToObject()); πE != nil {
				continue
			}
			// line 10: MAXLINESIZE = 76
			πF.SetLineno(10)
			if πE = πF.Globals().SetItem(πF, ßMAXLINESIZE.ToObject(), πg.NewInt(76).ToObject()); πE != nil {
				continue
			}
			// line 11: HEX = '0123456789ABCDEF'
			πF.SetLineno(11)
			if πE = πF.Globals().SetItem(πF, ßHEX.ToObject(), ß0123456789ABCDEF.ToObject()); πE != nil {
				continue
			}
			// line 12: EMPTYSTRING = ''
			πF.SetLineno(12)
			if πE = πF.Globals().SetItem(πF, ßEMPTYSTRING.ToObject(), ß.ToObject()); πE != nil {
				continue
			}
			// line 14: try:
			πF.SetLineno(14)
			πF.PushCheckpoint(2)
			// line 15: from binascii import a2b_qp, b2a_qp
			πF.SetLineno(15)
			if πTemp001, πE = πg.ImportModule(πF, "binascii"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßa2b_qp, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßa2b_qp.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßb2a_qp, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßb2a_qp.ToObject(), πTemp003); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label1
		Label2:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp004, πTemp005 = πF.ExcInfo()
			if πTemp002, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
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
			// line 16: except ImportError:
			πF.SetLineno(16)
		Label3:
			// line 17: a2b_qp = None
			πF.SetLineno(17)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßa2b_qp.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 18: b2a_qp = None
			πF.SetLineno(18)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßb2a_qp.ToObject(), πTemp002); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label1
		Label1:
			// line 21: def needsquoting(c, quotetabs, header):
			πF.SetLineno(21)
			πTemp007 = make([]πg.Param, 3)
			πTemp007[0] = πg.Param{Name: "c", Def: nil}
			πTemp007[1] = πg.Param{Name: "quotetabs", Def: nil}
			πTemp007[2] = πg.Param{Name: "header", Def: nil}
			πTemp002 = πg.NewFunction(πg.NewCode("needsquoting", "build/src/__python__/quopri.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µc *πg.Object = πArgs[0]; _ = µc
				var µquotetabs *πg.Object = πArgs[1]; _ = µquotetabs
				var µheader *πg.Object = πArgs[2]; _ = µheader
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
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
					// line 22: """Decide whether a particular character needs to be quoted.
					πF.SetLineno(22)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, πg.NewStr(" \t").ToObject(), µc); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 28: if c in ' \t':
					πF.SetLineno(28)
				Label1:
					// line 29: return quotetabs
					πF.SetLineno(29)
					if πE = πg.CheckLocal(πF, µquotetabs, "quotetabs"); πE != nil {
						continue
					}
					πR = µquotetabs
					continue
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µc, ß_.ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 31: if c == '_':
					πF.SetLineno(31)
				Label3:
					// line 32: return header
					πF.SetLineno(32)
					if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
						continue
					}
					πR = µheader
					continue
					goto Label4
				Label4:
					// line 33: return c == ESCAPE or not (' ' <= c <= '~')
					πF.SetLineno(33)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßESCAPE); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µc, πTemp004); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.LE(πF, πg.NewStr(" ").ToObject(), µc); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label6
					}
					if πTemp004, πE = πg.LE(πF, µc, πg.NewStr("~").ToObject()); πE != nil {
						continue
					}
				Label6:
					if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp005).ToObject()
					πTemp001 = πTemp003
				Label5:
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
			if πE = πF.Globals().SetItem(πF, ßneedsquoting.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 35: def quote(c):
			πF.SetLineno(35)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "c", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("quote", "build/src/__python__/quopri.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µc *πg.Object = πArgs[0]; _ = µc
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
				var πTemp007 *πg.Object
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 36: """Quote a single character."""
					πF.SetLineno(36)
					// line 37: i = ord(c)
					πF.SetLineno(37)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp001[0] = µc
					if πTemp002, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µi = πTemp003
					// line 38: return ESCAPE + HEX[i//16] + HEX[i%16]
					πF.SetLineno(38)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßESCAPE); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.FloorDiv(πF, µi, πg.NewInt(16).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp006
					if πTemp007, πE = πg.ResolveGlobal(πF, ßHEX); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πTemp004, πTemp006); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Mod(πF, µi, πg.NewInt(16).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp006, πE = πg.ResolveGlobal(πF, ßHEX); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp003, πTemp005); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßquote.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 42: def encode(input, output, quotetabs, header = 0):
			πF.SetLineno(42)
			πTemp007 = make([]πg.Param, 4)
			πTemp007[0] = πg.Param{Name: "input", Def: nil}
			πTemp007[1] = πg.Param{Name: "output", Def: nil}
			πTemp007[2] = πg.Param{Name: "quotetabs", Def: nil}
			πTemp007[3] = πg.Param{Name: "header", Def: πg.NewInt(0).ToObject()}
			πTemp008 = πg.NewFunction(πg.NewCode("encode", "build/src/__python__/quopri.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µinput *πg.Object = πArgs[0]; _ = µinput
				var µoutput *πg.Object = πArgs[1]; _ = µoutput
				var µquotetabs *πg.Object = πArgs[2]; _ = µquotetabs
				var µheader *πg.Object = πArgs[3]; _ = µheader
				var µdata *πg.Object = πg.UnboundLocal; _ = µdata
				var µodata *πg.Object = πg.UnboundLocal; _ = µodata
				var µwrite *πg.Object = πg.UnboundLocal; _ = µwrite
				var µprevline *πg.Object = πg.UnboundLocal; _ = µprevline
				var µline *πg.Object = πg.UnboundLocal; _ = µline
				var µoutline *πg.Object = πg.UnboundLocal; _ = µoutline
				var µstripped *πg.Object = πg.UnboundLocal; _ = µstripped
				var µc *πg.Object = πg.UnboundLocal; _ = µc
				var µthisline *πg.Object = πg.UnboundLocal; _ = µthisline
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 πg.KWArgs
				_ = πTemp006
				var πTemp007 []πg.Param
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 bool
				_ = πTemp010
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 3: goto Label3
					case 4: goto Label4
					case 10: goto Label10
					case 11: goto Label11
					case 21: goto Label21
					case 22: goto Label22
					default: panic("unexpected function state")
					}
					// line 43: """Read 'input', apply quoted-printable encoding, and write to 'output'.
					πF.SetLineno(43)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßb2a_qp); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002 != πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 53: if b2a_qp is not None:
					πF.SetLineno(53)
				Label1:
					// line 54: data = input.read()
					πF.SetLineno(54)
					if πE = πg.CheckLocal(πF, µinput, "input"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µinput, ßread, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µdata = πTemp002
					// line 55: odata = b2a_qp(data, quotetabs = quotetabs, header = header)
					πF.SetLineno(55)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					πTemp005[0] = µdata
					if πE = πg.CheckLocal(πF, µquotetabs, "quotetabs"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"quotetabs", µquotetabs},
						{"header", µheader},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßb2a_qp); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µodata = πTemp002
					// line 56: output.write(odata)
					πF.SetLineno(56)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µodata, "odata"); πE != nil {
						continue
					}
					πTemp005[0] = µodata
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µoutput, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 57: return
					πF.SetLineno(57)
					πR = πg.None
					continue
					goto Label2
				Label2:
					// line 59: def write(s, output=output, lineEnd='\n'):
					πF.SetLineno(59)
					πTemp007 = make([]πg.Param, 3)
					πTemp007[0] = πg.Param{Name: "s", Def: nil}
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					πTemp007[1] = πg.Param{Name: "output", Def: µoutput}
					πTemp007[2] = πg.Param{Name: "lineEnd", Def: πg.NewStr("\n").ToObject()}
					πTemp001 = πg.NewFunction(πg.NewCode("write", "build/src/__python__/quopri.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µs *πg.Object = πArgs[0]; _ = µs
						var µoutput *πg.Object = πArgs[1]; _ = µoutput
						var µlineEnd *πg.Object = πArgs[2]; _ = µlineEnd
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 []*πg.Object
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
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001 = µs
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label1
							}
							if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πTemp005, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µs, πTemp004); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, πg.NewStr(" \t").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp006).ToObject()
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µs, πg.NewStr(".").ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 62: if s and s[-1:] in ' \t':
							πF.SetLineno(62)
						Label2:
							// line 63: output.write(s[:-1] + quote(s[-1]) + lineEnd)
							πF.SetLineno(63)
							πTemp007 = πF.MakeArgs(1)
							if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp005, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µs, πTemp004); πE != nil {
								continue
							}
							πTemp008 = πF.MakeArgs(1)
							if πTemp009, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πTemp009
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µs, πTemp004); πE != nil {
								continue
							}
							πTemp008[0] = πTemp009
							if πTemp004, πE = πg.ResolveGlobal(πF, ßquote); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πTemp003, πE = πg.Add(πF, πTemp005, πTemp009); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlineEnd, "lineEnd"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, µlineEnd); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µoutput, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							goto Label5
							// line 64: elif s == '.':
							πF.SetLineno(64)
						Label3:
							// line 65: output.write(quote(s) + lineEnd)
							πF.SetLineno(65)
							πTemp007 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp008[0] = µs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßquote); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πE = πg.CheckLocal(πF, µlineEnd, "lineEnd"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp004, µlineEnd); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µoutput, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							goto Label5
						Label4:
							// line 67: output.write(s + lineEnd)
							πF.SetLineno(67)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlineEnd, "lineEnd"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µs, µlineEnd); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µoutput, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
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
					µwrite = πTemp001
					// line 69: prevline = None
					πF.SetLineno(69)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µprevline = πTemp002
					// line 70: while 1:
					πF.SetLineno(70)
					πF.PushCheckpoint(4)
					πTemp004 = false
				Label3:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label5
					}
					if πTemp008, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(3)            
					// line 71: line = input.readline()
					πF.SetLineno(71)
					if πE = πg.CheckLocal(πF, µinput, "input"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µinput, ßreadline, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µline = πTemp003
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µline); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp008).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label6
					}
					goto Label7
					// line 72: if not line:
					πF.SetLineno(72)
				Label6:
					// line 73: break
					πF.SetLineno(73)
					πTemp004 = true
					continue
					goto Label7
				Label7:
					// line 74: outline = []
					πF.SetLineno(74)
					πTemp005 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp005...).ToObject()
					µoutline = πTemp002
					// line 76: stripped = ''
					πF.SetLineno(76)
					µstripped = ß.ToObject()
					if πTemp009, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πTemp009, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, µline, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp009, πg.NewStr("\n").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label8
					}
					goto Label9
					// line 77: if line[-1:] == '\n':
					πF.SetLineno(77)
				Label8:
					// line 78: line = line[:-1]
					πF.SetLineno(78)
					if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp003, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µline, πTemp002); πE != nil {
						continue
					}
					µline = πTemp003
					// line 79: stripped = '\n'
					πF.SetLineno(79)
					µstripped = πg.NewStr("\n").ToObject()
					goto Label9
				Label9:
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µline); πE != nil {
						continue
					}
					πF.PushCheckpoint(11)
					πTemp008 = false
				Label10:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
						πF.PopCheckpoint()
						goto Label12
					}
					if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp010 = !isStop
					} else {
						πTemp010 = true
						µc = πTemp003
					}
					if πE != nil || !πTemp010 {
						continue
					}
					πF.PushCheckpoint(10)            
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp005[0] = µc
					if πE = πg.CheckLocal(πF, µquotetabs, "quotetabs"); πE != nil {
						continue
					}
					πTemp005[1] = µquotetabs
					if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
						continue
					}
					πTemp005[2] = µheader
					if πTemp003, πE = πg.ResolveGlobal(πF, ßneedsquoting); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp010, πE = πg.IsTrue(πF, πTemp009); πE != nil {
						continue
					}
					if πTemp010 {
						goto Label13
					}
					goto Label14
					// line 82: if needsquoting(c, quotetabs, header):
					πF.SetLineno(82)
				Label13:
					// line 83: c = quote(c)
					πF.SetLineno(83)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp005[0] = µc
					if πTemp003, πE = πg.ResolveGlobal(πF, ßquote); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µc = πTemp009
					goto Label14
				Label14:
					if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
						continue
					}
					πTemp003 = µheader
					if πTemp010, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp010 {
						goto Label15
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Eq(πF, µc, πg.NewStr(" ").ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp009
				Label15:
					if πTemp010, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp010 {
						goto Label16
					}
					goto Label17
					// line 84: if header and c == ' ':
					πF.SetLineno(84)
				Label16:
					// line 85: outline.append('_')
					πF.SetLineno(85)
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = ß_.ToObject()
					if πE = πg.CheckLocal(πF, µoutline, "outline"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µoutline, ßappend, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					goto Label18
				Label17:
					// line 87: outline.append(c)
					πF.SetLineno(87)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp005[0] = µc
					if πE = πg.CheckLocal(πF, µoutline, "outline"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µoutline, ßappend, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					goto Label18
				Label18:
					continue
				Label11:
					if πE != nil || πR != nil {
						continue
					}
				Label12:
					if πE = πg.CheckLocal(πF, µprevline, "prevline"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µprevline != πTemp003).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label19
					}
					goto Label20
					// line 89: if prevline is not None:
					πF.SetLineno(89)
				Label19:
					// line 90: write(prevline)
					πF.SetLineno(90)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µprevline, "prevline"); πE != nil {
						continue
					}
					πTemp005[0] = µprevline
					if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
						continue
					}
					if πTemp002, πE = µwrite.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					goto Label20
				Label20:
					// line 93: thisline = EMPTYSTRING.join(outline)
					πF.SetLineno(93)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µoutline, "outline"); πE != nil {
						continue
					}
					πTemp005[0] = µoutline
					if πTemp002, πE = πg.ResolveGlobal(πF, ßEMPTYSTRING); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µthisline = πTemp002
					// line 94: while len(thisline) > MAXLINESIZE:
					πF.SetLineno(94)
					πF.PushCheckpoint(22)
					πTemp008 = false
				Label21:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
						πF.PopCheckpoint()
						goto Label23
					}
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µthisline, "thisline"); πE != nil {
						continue
					}
					πTemp005[0] = µthisline
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßMAXLINESIZE); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GT(πF, πTemp009, πTemp003); πE != nil {
						continue
					}
					if πTemp010, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp010 {
						continue
					}
					πF.PushCheckpoint(21)            
					// line 97: write(thisline[:MAXLINESIZE-1], lineEnd='=\n')
					πF.SetLineno(97)
					πTemp005 = πF.MakeArgs(1)
					if πTemp009, πE = πg.ResolveGlobal(πF, ßMAXLINESIZE); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Sub(πF, πTemp009, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp003, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µthisline, "thisline"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µthisline, πTemp002); πE != nil {
						continue
					}
					πTemp005[0] = πTemp003
					πTemp006 = πg.KWArgs{
						{"lineEnd", πg.NewStr("=\n").ToObject()},
					}
					if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
						continue
					}
					if πTemp002, πE = µwrite.Call(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 98: thisline = thisline[MAXLINESIZE-1:]
					πF.SetLineno(98)
					if πTemp009, πE = πg.ResolveGlobal(πF, ßMAXLINESIZE); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Sub(πF, πTemp009, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µthisline, "thisline"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µthisline, πTemp002); πE != nil {
						continue
					}
					µthisline = πTemp003
					continue
				Label22:
					if πE != nil || πR != nil {
						continue
					}
				Label23:
					// line 100: prevline = thisline
					πF.SetLineno(100)
					if πE = πg.CheckLocal(πF, µthisline, "thisline"); πE != nil {
						continue
					}
					µprevline = µthisline
					continue
				Label4:
					if πE != nil || πR != nil {
						continue
					}
				Label5:
					if πE = πg.CheckLocal(πF, µprevline, "prevline"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µprevline != πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label24
					}
					goto Label25
					// line 102: if prevline is not None:
					πF.SetLineno(102)
				Label24:
					// line 103: write(prevline, lineEnd=stripped)
					πF.SetLineno(103)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µprevline, "prevline"); πE != nil {
						continue
					}
					πTemp005[0] = µprevline
					if πE = πg.CheckLocal(πF, µstripped, "stripped"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"lineEnd", µstripped},
					}
					if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
						continue
					}
					if πTemp002, πE = µwrite.Call(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					goto Label25
				Label25:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßencode.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 105: def encodestring(s, quotetabs = 0, header = 0):
			πF.SetLineno(105)
			πTemp007 = make([]πg.Param, 3)
			πTemp007[0] = πg.Param{Name: "s", Def: nil}
			πTemp007[1] = πg.Param{Name: "quotetabs", Def: πg.NewInt(0).ToObject()}
			πTemp007[2] = πg.Param{Name: "header", Def: πg.NewInt(0).ToObject()}
			πTemp009 = πg.NewFunction(πg.NewCode("encodestring", "build/src/__python__/quopri.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µquotetabs *πg.Object = πArgs[1]; _ = µquotetabs
				var µheader *πg.Object = πArgs[2]; _ = µheader
				var µStringIO *πg.Object = πg.UnboundLocal; _ = µStringIO
				var µinfp *πg.Object = πg.UnboundLocal; _ = µinfp
				var µoutfp *πg.Object = πg.UnboundLocal; _ = µoutfp
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 πg.KWArgs
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßb2a_qp); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002 != πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 106: if b2a_qp is not None:
					πF.SetLineno(106)
				Label1:
					// line 107: return b2a_qp(s, quotetabs = quotetabs, header = header)
					πF.SetLineno(107)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp005[0] = µs
					if πE = πg.CheckLocal(πF, µquotetabs, "quotetabs"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"quotetabs", µquotetabs},
						{"header", µheader},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßb2a_qp); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πR = πTemp002
					continue
					goto Label2
				Label2:
					// line 108: from cStringIO import StringIO
					πF.SetLineno(108)
					if πTemp005, πE = πg.ImportModule(πF, "cStringIO"); πE != nil {
						continue
					}
					πTemp001 = πTemp005[0]
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßStringIO, nil); πE != nil {
						continue
					}
					µStringIO = πTemp002
					// line 109: infp = StringIO(s)
					πF.SetLineno(109)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp005[0] = µs
					if πE = πg.CheckLocal(πF, µStringIO, "StringIO"); πE != nil {
						continue
					}
					if πTemp001, πE = µStringIO.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µinfp = πTemp001
					// line 110: outfp = StringIO()
					πF.SetLineno(110)
					if πE = πg.CheckLocal(πF, µStringIO, "StringIO"); πE != nil {
						continue
					}
					if πTemp001, πE = µStringIO.Call(πF, nil, nil); πE != nil {
						continue
					}
					µoutfp = πTemp001
					// line 111: encode(infp, outfp, quotetabs, header)
					πF.SetLineno(111)
					πTemp005 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µinfp, "infp"); πE != nil {
						continue
					}
					πTemp005[0] = µinfp
					if πE = πg.CheckLocal(πF, µoutfp, "outfp"); πE != nil {
						continue
					}
					πTemp005[1] = µoutfp
					if πE = πg.CheckLocal(πF, µquotetabs, "quotetabs"); πE != nil {
						continue
					}
					πTemp005[2] = µquotetabs
					if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
						continue
					}
					πTemp005[3] = µheader
					if πTemp001, πE = πg.ResolveGlobal(πF, ßencode); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 112: return outfp.getvalue()
					πF.SetLineno(112)
					if πE = πg.CheckLocal(πF, µoutfp, "outfp"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µoutfp, ßgetvalue, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßencodestring.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 116: def decode(input, output, header = 0):
			πF.SetLineno(116)
			πTemp007 = make([]πg.Param, 3)
			πTemp007[0] = πg.Param{Name: "input", Def: nil}
			πTemp007[1] = πg.Param{Name: "output", Def: nil}
			πTemp007[2] = πg.Param{Name: "header", Def: πg.NewInt(0).ToObject()}
			πTemp010 = πg.NewFunction(πg.NewCode("decode", "build/src/__python__/quopri.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µinput *πg.Object = πArgs[0]; _ = µinput
				var µoutput *πg.Object = πArgs[1]; _ = µoutput
				var µheader *πg.Object = πArgs[2]; _ = µheader
				var µdata *πg.Object = πg.UnboundLocal; _ = µdata
				var µodata *πg.Object = πg.UnboundLocal; _ = µodata
				var µnew *πg.Object = πg.UnboundLocal; _ = µnew
				var µline *πg.Object = πg.UnboundLocal; _ = µline
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µn *πg.Object = πg.UnboundLocal; _ = µn
				var µpartial *πg.Object = πg.UnboundLocal; _ = µpartial
				var µc *πg.Object = πg.UnboundLocal; _ = µc
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 πg.KWArgs
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πTemp010 bool
				_ = πTemp010
				var πTemp011 bool
				_ = πTemp011
				var πTemp012 []*πg.Object
				_ = πTemp012
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 3: goto Label3
					case 4: goto Label4
					case 12: goto Label12
					case 13: goto Label13
					case 16: goto Label16
					case 17: goto Label17
					default: panic("unexpected function state")
					}
					// line 117: """Read 'input', apply quoted-printable decoding, and write to 'output'.
					πF.SetLineno(117)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßa2b_qp); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002 != πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 121: if a2b_qp is not None:
					πF.SetLineno(121)
				Label1:
					// line 122: data = input.read()
					πF.SetLineno(122)
					if πE = πg.CheckLocal(πF, µinput, "input"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µinput, ßread, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µdata = πTemp002
					// line 123: odata = a2b_qp(data, header = header)
					πF.SetLineno(123)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					πTemp005[0] = µdata
					if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"header", µheader},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßa2b_qp); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µodata = πTemp002
					// line 124: output.write(odata)
					πF.SetLineno(124)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µodata, "odata"); πE != nil {
						continue
					}
					πTemp005[0] = µodata
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µoutput, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 125: return
					πF.SetLineno(125)
					πR = πg.None
					continue
					goto Label2
				Label2:
					// line 127: new = ''
					πF.SetLineno(127)
					µnew = ß.ToObject()
					// line 128: while 1:
					πF.SetLineno(128)
					πF.PushCheckpoint(4)
					πTemp004 = false
				Label3:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label5
					}
					if πTemp007, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(3)            
					// line 129: line = input.readline()
					πF.SetLineno(129)
					if πE = πg.CheckLocal(πF, µinput, "input"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µinput, ßreadline, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µline = πTemp002
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, µline); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label6
					}
					goto Label7
					// line 130: if not line: break
					πF.SetLineno(130)
				Label6:
					// line 130: if not line: break
					πF.SetLineno(130)
					πTemp004 = true
					continue
					goto Label7
				Label7:
					// line 131: i, n = 0, len(line)
					πF.SetLineno(131)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp005[0] = µline
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001 = πg.NewTuple2(πg.NewInt(0).ToObject(), πTemp003).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}}}, πTemp001); πE != nil {
						continue
					}
					µi = πTemp002
					µn = πTemp003
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GT(πF, µn, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp002
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label8
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Sub(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp008
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µline, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp008, πg.NewStr("\n").ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp002
				Label8:
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label9
					}
					goto Label10
					// line 132: if n > 0 and line[n-1] == '\n':
					πF.SetLineno(132)
				Label9:
					// line 133: partial = 0; n = n-1
					πF.SetLineno(133)
					µpartial = πg.NewInt(0).ToObject()
					// line 133: partial = 0; n = n-1
					πF.SetLineno(133)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µn = πTemp001
					// line 135: while n > 0 and line[n-1] in " \t\r":
					πF.SetLineno(135)
					πF.PushCheckpoint(13)
					πTemp007 = false
				Label12:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp007 {
						πF.PopCheckpoint()
						goto Label14
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GT(πF, µn, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp002
					if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp010 {
						goto Label15
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Sub(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp008
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µline, πTemp003); πE != nil {
						continue
					}
					if πTemp011, πE = πg.Contains(πF, πg.NewStr(" \t\r").ToObject(), πTemp008); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp011).ToObject()
					πTemp001 = πTemp002
				Label15:
					if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(12)            
					// line 136: n = n-1
					πF.SetLineno(136)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µn = πTemp001
					continue
				Label13:
					if πE != nil || πR != nil {
						continue
					}
				Label14:
					goto Label11
				Label10:
					// line 138: partial = 1
					πF.SetLineno(138)
					µpartial = πg.NewInt(1).ToObject()
					goto Label11
				Label11:
					// line 139: while i < n:
					πF.SetLineno(139)
					πF.PushCheckpoint(17)
					πTemp007 = false
				Label16:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp007 {
						πF.PopCheckpoint()
						goto Label18
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µi, µn); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(16)            
					// line 140: c = line[i]
					πF.SetLineno(140)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp001 = µi
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µline, πTemp001); πE != nil {
						continue
					}
					µc = πTemp002
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µc, ß_.ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp002
					if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp009 {
						goto Label19
					}
					if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
						continue
					}
					πTemp001 = µheader
				Label19:
					if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label20
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßESCAPE); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, µc, πTemp002); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label21
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp003, µn); πE != nil {
						continue
					}
					πTemp001 = πTemp002
					if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp009 {
						goto Label22
					}
					if πE = πg.CheckLocal(πF, µpartial, "partial"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.IsTrue(πF, µpartial); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp010).ToObject()
					πTemp001 = πTemp002
				Label22:
					if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label23
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LT(πF, πTemp003, µn); πE != nil {
						continue
					}
					πTemp001 = πTemp002
					if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp009 {
						goto Label24
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp008
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µline, πTemp003); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßESCAPE); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp008, πTemp003); πE != nil {
						continue
					}
					πTemp001 = πTemp002
				Label24:
					if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label25
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µi, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LT(πF, πTemp003, µn); πE != nil {
						continue
					}
					πTemp001 = πTemp002
					if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp009 {
						goto Label26
					}
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µline, πTemp002); πE != nil {
						continue
					}
					πTemp005[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßishex); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001 = πTemp003
					if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp009 {
						goto Label26
					}
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µi, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µline, πTemp002); πE != nil {
						continue
					}
					πTemp005[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßishex); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001 = πTemp003
				Label26:
					if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label27
					}
					goto Label28
					// line 141: if c == '_' and header:
					πF.SetLineno(141)
				Label20:
					// line 142: new = new + ' '; i = i+1
					πF.SetLineno(142)
					if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µnew, πg.NewStr(" ").ToObject()); πE != nil {
						continue
					}
					µnew = πTemp001
					// line 142: new = new + ' '; i = i+1
					πF.SetLineno(142)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µi = πTemp001
					goto Label29
					// line 143: elif c != ESCAPE:
					πF.SetLineno(143)
				Label21:
					// line 144: new = new + c; i = i+1
					πF.SetLineno(144)
					if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µnew, µc); πE != nil {
						continue
					}
					µnew = πTemp001
					// line 144: new = new + c; i = i+1
					πF.SetLineno(144)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µi = πTemp001
					goto Label29
					// line 145: elif i+1 == n and not partial:
					πF.SetLineno(145)
				Label23:
					// line 146: partial = 1; break
					πF.SetLineno(146)
					µpartial = πg.NewInt(1).ToObject()
					// line 146: partial = 1; break
					πF.SetLineno(146)
					πTemp007 = true
					continue
					goto Label29
					// line 147: elif i+1 < n and line[i+1] == ESCAPE:
					πF.SetLineno(147)
				Label25:
					// line 148: new = new + ESCAPE; i = i+2
					πF.SetLineno(148)
					if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßESCAPE); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µnew, πTemp002); πE != nil {
						continue
					}
					µnew = πTemp001
					// line 148: new = new + ESCAPE; i = i+2
					πF.SetLineno(148)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µi, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					µi = πTemp001
					goto Label29
					// line 149: elif i+2 < n and ishex(line[i+1]) and ishex(line[i+2]):
					πF.SetLineno(149)
				Label27:
					// line 150: new = new + chr(unhex(line[i+1:i+3])); i = i+3
					πF.SetLineno(150)
					if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
						continue
					}
					πTemp005 = πF.MakeArgs(1)
					πTemp012 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Add(πF, µi, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πTemp008, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µline, πTemp002); πE != nil {
						continue
					}
					πTemp012[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßunhex); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp005[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp001, πE = πg.Add(πF, µnew, πTemp003); πE != nil {
						continue
					}
					µnew = πTemp001
					// line 150: new = new + chr(unhex(line[i+1:i+3])); i = i+3
					πF.SetLineno(150)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µi, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					µi = πTemp001
					goto Label29
				Label28:
					// line 152: new = new + c; i = i+1
					πF.SetLineno(152)
					if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µnew, µc); πE != nil {
						continue
					}
					µnew = πTemp001
					// line 152: new = new + c; i = i+1
					πF.SetLineno(152)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µi = πTemp001
					goto Label29
				Label29:
					continue
				Label17:
					if πE != nil || πR != nil {
						continue
					}
				Label18:
					if πE = πg.CheckLocal(πF, µpartial, "partial"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, µpartial); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label30
					}
					goto Label31
					// line 153: if not partial:
					πF.SetLineno(153)
				Label30:
					// line 154: output.write(new + '\n')
					πF.SetLineno(154)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µnew, πg.NewStr("\n").ToObject()); πE != nil {
						continue
					}
					πTemp005[0] = πTemp001
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µoutput, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 155: new = ''
					πF.SetLineno(155)
					µnew = ß.ToObject()
					goto Label31
				Label31:
					continue
				Label4:
					if πE != nil || πR != nil {
						continue
					}
				Label5:
					if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µnew); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label32
					}
					goto Label33
					// line 156: if new:
					πF.SetLineno(156)
				Label32:
					// line 157: output.write(new)
					πF.SetLineno(157)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
						continue
					}
					πTemp005[0] = µnew
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µoutput, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					goto Label33
				Label33:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßdecode.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 159: def decodestring(s, header = 0):
			πF.SetLineno(159)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "s", Def: nil}
			πTemp007[1] = πg.Param{Name: "header", Def: πg.NewInt(0).ToObject()}
			πTemp011 = πg.NewFunction(πg.NewCode("decodestring", "build/src/__python__/quopri.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µheader *πg.Object = πArgs[1]; _ = µheader
				var µStringIO *πg.Object = πg.UnboundLocal; _ = µStringIO
				var µinfp *πg.Object = πg.UnboundLocal; _ = µinfp
				var µoutfp *πg.Object = πg.UnboundLocal; _ = µoutfp
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 πg.KWArgs
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßa2b_qp); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002 != πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 160: if a2b_qp is not None:
					πF.SetLineno(160)
				Label1:
					// line 161: return a2b_qp(s, header = header)
					πF.SetLineno(161)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp005[0] = µs
					if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"header", µheader},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßa2b_qp); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πR = πTemp002
					continue
					goto Label2
				Label2:
					// line 162: from cStringIO import StringIO
					πF.SetLineno(162)
					if πTemp005, πE = πg.ImportModule(πF, "cStringIO"); πE != nil {
						continue
					}
					πTemp001 = πTemp005[0]
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßStringIO, nil); πE != nil {
						continue
					}
					µStringIO = πTemp002
					// line 163: infp = StringIO(s)
					πF.SetLineno(163)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp005[0] = µs
					if πE = πg.CheckLocal(πF, µStringIO, "StringIO"); πE != nil {
						continue
					}
					if πTemp001, πE = µStringIO.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µinfp = πTemp001
					// line 164: outfp = StringIO()
					πF.SetLineno(164)
					if πE = πg.CheckLocal(πF, µStringIO, "StringIO"); πE != nil {
						continue
					}
					if πTemp001, πE = µStringIO.Call(πF, nil, nil); πE != nil {
						continue
					}
					µoutfp = πTemp001
					// line 165: decode(infp, outfp, header = header)
					πF.SetLineno(165)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µinfp, "infp"); πE != nil {
						continue
					}
					πTemp005[0] = µinfp
					if πE = πg.CheckLocal(πF, µoutfp, "outfp"); πE != nil {
						continue
					}
					πTemp005[1] = µoutfp
					if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"header", µheader},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßdecode); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 166: return outfp.getvalue()
					πF.SetLineno(166)
					if πE = πg.CheckLocal(πF, µoutfp, "outfp"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µoutfp, ßgetvalue, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßdecodestring.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 171: def ishex(c):
			πF.SetLineno(171)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "c", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("ishex", "build/src/__python__/quopri.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µc *πg.Object = πArgs[0]; _ = µc
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
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
					// line 172: """Return true if the character 'c' is a hexadecimal digit."""
					πF.SetLineno(172)
					// line 173: return '0' <= c <= '9' or 'a' <= c <= 'f' or 'A' <= c <= 'F'
					πF.SetLineno(173)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LE(πF, ß0.ToObject(), µc); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label2
					}
					if πTemp003, πE = πg.LE(πF, µc, ß9.ToObject()); πE != nil {
						continue
					}
				Label2:
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LE(πF, ßa.ToObject(), µc); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label3
					}
					if πTemp003, πE = πg.LE(πF, µc, ßf.ToObject()); πE != nil {
						continue
					}
				Label3:
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LE(πF, ßA.ToObject(), µc); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label4
					}
					if πTemp003, πE = πg.LE(πF, µc, ßF.ToObject()); πE != nil {
						continue
					}
				Label4:
					πTemp001 = πTemp003
				Label1:
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
			if πE = πF.Globals().SetItem(πF, ßishex.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 175: def unhex(s):
			πF.SetLineno(175)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "s", Def: nil}
			πTemp013 = πg.NewFunction(πg.NewCode("unhex", "build/src/__python__/quopri.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µbits *πg.Object = πg.UnboundLocal; _ = µbits
				var µc *πg.Object = πg.UnboundLocal; _ = µc
				var µi *πg.Object = πg.UnboundLocal; _ = µi
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
				var πTemp007 *πg.Object
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
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 176: """Get the integer value of a hexadecimal number."""
					πF.SetLineno(176)
					// line 177: bits = 0
					πF.SetLineno(177)
					µbits = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µs); πE != nil {
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
						µc = πTemp004
					}
					if πE != nil || !πTemp003 {
						continue
					}
					πF.PushCheckpoint(1)            
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.LE(πF, ß0.ToObject(), µc); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label4
					}
					if πTemp004, πE = πg.LE(πF, µc, ß9.ToObject()); πE != nil {
						continue
					}
				Label4:
					if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.LE(πF, ßa.ToObject(), µc); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label6
					}
					if πTemp004, πE = πg.LE(πF, µc, ßf.ToObject()); πE != nil {
						continue
					}
				Label6:
					if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label7
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.LE(πF, ßA.ToObject(), µc); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label8
					}
					if πTemp004, πE = πg.LE(πF, µc, ßF.ToObject()); πE != nil {
						continue
					}
				Label8:
					if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label9
					}
					goto Label10
					// line 179: if '0' <= c <= '9':
					πF.SetLineno(179)
				Label5:
					// line 180: i = ord('0')
					πF.SetLineno(180)
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = ß0.ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µi = πTemp006
					goto Label11
					// line 181: elif 'a' <= c <= 'f':
					πF.SetLineno(181)
				Label7:
					// line 182: i = ord('a')-10
					πF.SetLineno(182)
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = ßa.ToObject()
					if πTemp006, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp004, πE = πg.Sub(πF, πTemp007, πg.NewInt(10).ToObject()); πE != nil {
						continue
					}
					µi = πTemp004
					goto Label11
					// line 183: elif 'A' <= c <= 'F':
					πF.SetLineno(183)
				Label9:
					// line 184: i = ord('A')-10
					πF.SetLineno(184)
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = ßA.ToObject()
					if πTemp006, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp004, πE = πg.Sub(πF, πTemp007, πg.NewInt(10).ToObject()); πE != nil {
						continue
					}
					µi = πTemp004
					goto Label11
				Label10:
					// line 186: break
					πF.SetLineno(186)
					πTemp002 = true
					continue
					goto Label11
				Label11:
					// line 187: bits = bits*16 + (ord(c) - i)
					πF.SetLineno(187)
					if πE = πg.CheckLocal(πF, µbits, "bits"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Mul(πF, µbits, πg.NewInt(16).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp005[0] = µc
					if πTemp008, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp008.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Sub(πF, πTemp009, µi); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, πTemp006, πTemp007); πE != nil {
						continue
					}
					µbits = πTemp004
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 188: return bits
					πF.SetLineno(188)
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
			if πE = πF.Globals().SetItem(πF, ßunhex.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 192: def main():
			πF.SetLineno(192)
			πTemp007 = make([]πg.Param, 0)
			πTemp014 = πg.NewFunction(πg.NewCode("main", "build/src/__python__/quopri.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsys *πg.Object = πg.UnboundLocal; _ = µsys
				var µgetopt *πg.Object = πg.UnboundLocal; _ = µgetopt
				var µopts *πg.Object = πg.UnboundLocal; _ = µopts
				var µargs *πg.Object = πg.UnboundLocal; _ = µargs
				var µmsg *πg.Object = πg.UnboundLocal; _ = µmsg
				var µdeco *πg.Object = πg.UnboundLocal; _ = µdeco
				var µtabs *πg.Object = πg.UnboundLocal; _ = µtabs
				var µo *πg.Object = πg.UnboundLocal; _ = µo
				var µa *πg.Object = πg.UnboundLocal; _ = µa
				var µsts *πg.Object = πg.UnboundLocal; _ = µsts
				var µfile *πg.Object = πg.UnboundLocal; _ = µfile
				var µfp *πg.Object = πg.UnboundLocal; _ = µfp
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
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					case 4: goto Label4
					case 5: goto Label5
					case 16: goto Label16
					case 17: goto Label17
					case 23: goto Label23
					default: panic("unexpected function state")
					}
					// line 193: import sys
					πF.SetLineno(193)
					if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					µsys = πTemp001
					// line 194: import getopt
					πF.SetLineno(194)
					if πTemp002, πE = πg.ImportModule(πF, "getopt"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					µgetopt = πTemp001
					// line 195: try:
					πF.SetLineno(195)
					πF.PushCheckpoint(2)
					// line 196: opts, args = getopt.getopt(sys.argv[1:], 'td')
					πF.SetLineno(196)
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
					πTemp002[1] = ßtd.ToObject()
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
					// line 197: except getopt.error, msg:
					πF.SetLineno(197)
				Label3:
					µmsg = πTemp005.ToObject()
					// line 198: sys.stdout = sys.stderr
					πF.SetLineno(198)
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
					// line 199: print msg
					πF.SetLineno(199)
					πTemp002 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp002[0] = µmsg
					if πE = πg.Print(πF, πTemp002, true); πE != nil {
						continue
					}
					// line 200: print "usage: quopri [-t | -d] [file] ..."
					πF.SetLineno(200)
					πTemp002 = make([]*πg.Object, 1)
					πTemp002[0] = πg.NewStr("usage: quopri [-t | -d] [file] ...").ToObject()
					if πE = πg.Print(πF, πTemp002, true); πE != nil {
						continue
					}
					// line 201: print "-t: quote tabs"
					πF.SetLineno(201)
					πTemp002 = make([]*πg.Object, 1)
					πTemp002[0] = πg.NewStr("-t: quote tabs").ToObject()
					if πE = πg.Print(πF, πTemp002, true); πE != nil {
						continue
					}
					// line 202: print "-d: decode; default encode"
					πF.SetLineno(202)
					πTemp002 = make([]*πg.Object, 1)
					πTemp002[0] = πg.NewStr("-d: decode; default encode").ToObject()
					if πE = πg.Print(πF, πTemp002, true); πE != nil {
						continue
					}
					// line 203: sys.exit(2)
					πF.SetLineno(203)
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
					// line 204: deco = 0
					πF.SetLineno(204)
					µdeco = πg.NewInt(0).ToObject()
					// line 205: tabs = 0
					πF.SetLineno(205)
					µtabs = πg.NewInt(0).ToObject()
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
						πTemp008 = !isStop
					} else {
						πTemp008 = true
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp009}}}, πTemp003); πE != nil {
							continue
						}
						µo = πTemp004
						µa = πTemp009
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(4)            
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µo, πg.NewStr("-t").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label7
					}
					goto Label8
					// line 207: if o == '-t': tabs = 1
					πF.SetLineno(207)
				Label7:
					// line 207: if o == '-t': tabs = 1
					πF.SetLineno(207)
					µtabs = πg.NewInt(1).ToObject()
					goto Label8
				Label8:
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µo, πg.NewStr("-d").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label9
					}
					goto Label10
					// line 208: if o == '-d': deco = 1
					πF.SetLineno(208)
				Label9:
					// line 208: if o == '-d': deco = 1
					πF.SetLineno(208)
					µdeco = πg.NewInt(1).ToObject()
					goto Label10
				Label10:
					continue
				Label5:
					if πE != nil || πR != nil {
						continue
					}
				Label6:
					if πE = πg.CheckLocal(πF, µtabs, "tabs"); πE != nil {
						continue
					}
					πTemp001 = µtabs
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label11
					}
					if πE = πg.CheckLocal(πF, µdeco, "deco"); πE != nil {
						continue
					}
					πTemp001 = µdeco
				Label11:
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label12
					}
					goto Label13
					// line 209: if tabs and deco:
					πF.SetLineno(209)
				Label12:
					// line 210: sys.stdout = sys.stderr
					πF.SetLineno(210)
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
					// line 211: print "-t and -d are mutually exclusive"
					πF.SetLineno(211)
					πTemp002 = make([]*πg.Object, 1)
					πTemp002[0] = πg.NewStr("-t and -d are mutually exclusive").ToObject()
					if πE = πg.Print(πF, πTemp002, true); πE != nil {
						continue
					}
					// line 212: sys.exit(2)
					πF.SetLineno(212)
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
					goto Label13
				Label13:
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, µargs); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label14
					}
					goto Label15
					// line 213: if not args: args = ['-']
					πF.SetLineno(213)
				Label14:
					// line 213: if not args: args = ['-']
					πF.SetLineno(213)
					πTemp002 = make([]*πg.Object, 1)
					πTemp002[0] = πg.NewStr("-").ToObject()
					πTemp001 = πg.NewList(πTemp002...).ToObject()
					µargs = πTemp001
					goto Label15
				Label15:
					// line 214: sts = 0
					πF.SetLineno(214)
					µsts = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µargs); πE != nil {
						continue
					}
					πF.PushCheckpoint(17)
					πTemp007 = false
				Label16:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp007 {
						πF.PopCheckpoint()
						goto Label18
					}
					if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
						µfile = πTemp003
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(16)            
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µfile, πg.NewStr("-").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label19
					}
					goto Label20
					// line 216: if file == '-':
					πF.SetLineno(216)
				Label19:
					// line 217: fp = sys.stdin
					πF.SetLineno(217)
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsys, ßstdin, nil); πE != nil {
						continue
					}
					µfp = πTemp003
					goto Label21
				Label20:
					// line 219: try:
					πF.SetLineno(219)
					πF.PushCheckpoint(23)
					// line 220: fp = open(file)
					πF.SetLineno(220)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp002[0] = µfile
					if πTemp003, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µfp = πTemp004
					πF.PopCheckpoint()
					goto Label22
				Label23:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp005, πTemp006 = πF.ExcInfo()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label24
					}
					πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
					continue
					// line 221: except IOError, msg:
					πF.SetLineno(221)
				Label24:
					µmsg = πTemp005.ToObject()
					// line 222: sys.stderr.write("%s: can't open (%s)\n" % (file, msg))
					πF.SetLineno(222)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple2(µfile, µmsg).ToObject()
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s: can't open (%s)\n").ToObject(), πTemp004); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsys, ßstderr, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 223: sts = 1
					πF.SetLineno(223)
					µsts = πg.NewInt(1).ToObject()
					// line 224: continue
					πF.SetLineno(224)
					continue
					πF.RestoreExc(nil, nil)
					goto Label22
				Label22:
					goto Label21
				Label21:
					if πE = πg.CheckLocal(πF, µdeco, "deco"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µdeco); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label25
					}
					goto Label26
					// line 225: if deco:
					πF.SetLineno(225)
				Label25:
					// line 226: decode(fp, sys.stdout)
					πF.SetLineno(226)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
						continue
					}
					πTemp002[0] = µfp
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsys, ßstdout, nil); πE != nil {
						continue
					}
					πTemp002[1] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßdecode); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label27
				Label26:
					// line 228: encode(fp, sys.stdout, tabs)
					πF.SetLineno(228)
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
						continue
					}
					πTemp002[0] = µfp
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsys, ßstdout, nil); πE != nil {
						continue
					}
					πTemp002[1] = πTemp003
					if πE = πg.CheckLocal(πF, µtabs, "tabs"); πE != nil {
						continue
					}
					πTemp002[2] = µtabs
					if πTemp003, πE = πg.ResolveGlobal(πF, ßencode); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label27
				Label27:
					if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µsys, ßstdin, nil); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µfp != πTemp004).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label28
					}
					goto Label29
					// line 229: if fp is not sys.stdin:
					πF.SetLineno(229)
				Label28:
					// line 230: fp.close()
					πF.SetLineno(230)
					if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µfp, ßclose, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					goto Label29
				Label29:
					continue
				Label17:
					if πE != nil || πR != nil {
						continue
					}
				Label18:
					if πE = πg.CheckLocal(πF, µsts, "sts"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, µsts); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label30
					}
					goto Label31
					// line 231: if sts:
					πF.SetLineno(231)
				Label30:
					// line 232: sys.exit(sts)
					πF.SetLineno(232)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsts, "sts"); πE != nil {
						continue
					}
					πTemp002[0] = µsts
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
					goto Label31
				Label31:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßmain.ToObject(), πTemp014); πE != nil {
				continue
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp015, πE = πg.Eq(πF, πTemp016, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp015); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label4
			}
			goto Label5
			// line 236: if __name__ == '__main__':
			πF.SetLineno(236)
		Label4:
			// line 237: main()
			πF.SetLineno(237)
			if πTemp015, πE = πg.ResolveGlobal(πF, ßmain); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp015.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label5
		Label5:
		}
		return nil, πE
	})
	πg.RegisterModule("quopri", Code)
}
