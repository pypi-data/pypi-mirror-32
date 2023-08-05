package string
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/string.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ß01234567 := πg.InternStr("01234567")
		ß0123456789 := πg.InternStr("0123456789")
		ßABCDEF := πg.InternStr("ABCDEF")
		ßABCDEFGHIJKLMNOPQRSTUVWXYZ := πg.InternStr("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		ßFormatter := πg.InternStr("Formatter")
		ßIGNORECASE := πg.InternStr("IGNORECASE")
		ßIndexError := πg.InternStr("IndexError")
		ßKeyError := πg.InternStr("KeyError")
		ßNone := πg.InternStr("None")
		ßTemplate := πg.InternStr("Template")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßVERBOSE := πg.InternStr("VERBOSE")
		ßValueError := πg.InternStr("ValueError")
		ß_TemplateMetaclass := πg.InternStr("_TemplateMetaclass")
		ß__getitem__ := πg.InternStr("__getitem__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_float := πg.InternStr("_float")
		ß_formatter_field_name_split := πg.InternStr("_formatter_field_name_split")
		ß_formatter_parser := πg.InternStr("_formatter_parser")
		ß_idmap := πg.InternStr("_idmap")
		ß_idmapL := πg.InternStr("_idmapL")
		ß_int := πg.InternStr("_int")
		ß_invalid := πg.InternStr("_invalid")
		ß_long := πg.InternStr("_long")
		ß_multimap := πg.InternStr("_multimap")
		ß_primary := πg.InternStr("_primary")
		ß_re := πg.InternStr("_re")
		ß_secondary := πg.InternStr("_secondary")
		ß_vformat := πg.InternStr("_vformat")
		ßabcdef := πg.InternStr("abcdef")
		ßabcdefghijklmnopqrstuvwxyz := πg.InternStr("abcdefghijklmnopqrstuvwxyz")
		ßadd := πg.InternStr("add")
		ßappend := πg.InternStr("append")
		ßascii_letters := πg.InternStr("ascii_letters")
		ßascii_lowercase := πg.InternStr("ascii_lowercase")
		ßascii_uppercase := πg.InternStr("ascii_uppercase")
		ßatof := πg.InternStr("atof")
		ßatof_error := πg.InternStr("atof_error")
		ßatoi := πg.InternStr("atoi")
		ßatoi_error := πg.InternStr("atoi_error")
		ßatol := πg.InternStr("atol")
		ßatol_error := πg.InternStr("atol_error")
		ßbasestring := πg.InternStr("basestring")
		ßbraced := πg.InternStr("braced")
		ßcapitalize := πg.InternStr("capitalize")
		ßcapwords := πg.InternStr("capwords")
		ßcenter := πg.InternStr("center")
		ßcheck_unused_args := πg.InternStr("check_unused_args")
		ßchr := πg.InternStr("chr")
		ßcompile := πg.InternStr("compile")
		ßconvert_field := πg.InternStr("convert_field")
		ßcount := πg.InternStr("count")
		ßdelimiter := πg.InternStr("delimiter")
		ßdigits := πg.InternStr("digits")
		ßescape := πg.InternStr("escape")
		ßescaped := πg.InternStr("escaped")
		ßexpandtabs := πg.InternStr("expandtabs")
		ßfind := πg.InternStr("find")
		ßfloat := πg.InternStr("float")
		ßformat := πg.InternStr("format")
		ßformat_field := πg.InternStr("format_field")
		ßformat_string := πg.InternStr("format_string")
		ßget_field := πg.InternStr("get_field")
		ßget_value := πg.InternStr("get_value")
		ßgetattr := πg.InternStr("getattr")
		ßgroup := πg.InternStr("group")
		ßhexdigits := πg.InternStr("hexdigits")
		ßidpattern := πg.InternStr("idpattern")
		ßindex := πg.InternStr("index")
		ßindex_error := πg.InternStr("index_error")
		ßint := πg.InternStr("int")
		ßinvalid := πg.InternStr("invalid")
		ßisinstance := πg.InternStr("isinstance")
		ßjoin := πg.InternStr("join")
		ßjoinfields := πg.InternStr("joinfields")
		ßl := πg.InternStr("l")
		ßlen := πg.InternStr("len")
		ßletters := πg.InternStr("letters")
		ßlist := πg.InternStr("list")
		ßljust := πg.InternStr("ljust")
		ßlong := πg.InternStr("long")
		ßlower := πg.InternStr("lower")
		ßlowercase := πg.InternStr("lowercase")
		ßlstrip := πg.InternStr("lstrip")
		ßmaketrans := πg.InternStr("maketrans")
		ßmap := πg.InternStr("map")
		ßnamed := πg.InternStr("named")
		ßobject := πg.InternStr("object")
		ßoctdigits := πg.InternStr("octdigits")
		ßord := πg.InternStr("ord")
		ßparse := πg.InternStr("parse")
		ßpattern := πg.InternStr("pattern")
		ßpop := πg.InternStr("pop")
		ßprintable := πg.InternStr("printable")
		ßpunctuation := πg.InternStr("punctuation")
		ßr := πg.InternStr("r")
		ßrange := πg.InternStr("range")
		ßreplace := πg.InternStr("replace")
		ßrepr := πg.InternStr("repr")
		ßrfind := πg.InternStr("rfind")
		ßrindex := πg.InternStr("rindex")
		ßrjust := πg.InternStr("rjust")
		ßrsplit := πg.InternStr("rsplit")
		ßrstrip := πg.InternStr("rstrip")
		ßs := πg.InternStr("s")
		ßsafe_substitute := πg.InternStr("safe_substitute")
		ßset := πg.InternStr("set")
		ßsplit := πg.InternStr("split")
		ßsplitfields := πg.InternStr("splitfields")
		ßsplitlines := πg.InternStr("splitlines")
		ßstart := πg.InternStr("start")
		ßstr := πg.InternStr("str")
		ßstrip := πg.InternStr("strip")
		ßsub := πg.InternStr("sub")
		ßsubstitute := πg.InternStr("substitute")
		ßsuper := πg.InternStr("super")
		ßswapcase := πg.InternStr("swapcase")
		ßtemplate := πg.InternStr("template")
		ßtranslate := πg.InternStr("translate")
		ßtype := πg.InternStr("type")
		ßupper := πg.InternStr("upper")
		ßuppercase := πg.InternStr("uppercase")
		ßvformat := πg.InternStr("vformat")
		ßwhitespace := πg.InternStr("whitespace")
		ßxrange := πg.InternStr("xrange")
		ßzfill := πg.InternStr("zfill")
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
		var πTemp006 []πg.Param
		_ = πTemp006
		var πTemp007 []*πg.Object
		_ = πTemp007
		var πTemp008 []*πg.Object
		_ = πTemp008
		var πTemp009 *πg.Dict
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """A collection of string operations (most are no longer used).
			πF.SetLineno(1)
			// line 23: whitespace = ' \t\n\r\v\f'
			πF.SetLineno(23)
			if πE = πF.Globals().SetItem(πF, ßwhitespace.ToObject(), πg.NewStr(" \t\n\r\x0b\x0c").ToObject()); πE != nil {
				continue
			}
			// line 24: lowercase = 'abcdefghijklmnopqrstuvwxyz'
			πF.SetLineno(24)
			if πE = πF.Globals().SetItem(πF, ßlowercase.ToObject(), ßabcdefghijklmnopqrstuvwxyz.ToObject()); πE != nil {
				continue
			}
			// line 25: uppercase = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'
			πF.SetLineno(25)
			if πE = πF.Globals().SetItem(πF, ßuppercase.ToObject(), ßABCDEFGHIJKLMNOPQRSTUVWXYZ.ToObject()); πE != nil {
				continue
			}
			// line 26: letters = lowercase + uppercase
			πF.SetLineno(26)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßlowercase); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßuppercase); πE != nil {
				continue
			}
			if πTemp001, πE = πg.Add(πF, πTemp002, πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßletters.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 27: ascii_lowercase = lowercase
			πF.SetLineno(27)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßlowercase); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßascii_lowercase.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 28: ascii_uppercase = uppercase
			πF.SetLineno(28)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßuppercase); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßascii_uppercase.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 29: ascii_letters = ascii_lowercase + ascii_uppercase
			πF.SetLineno(29)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßascii_lowercase); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßascii_uppercase); πE != nil {
				continue
			}
			if πTemp001, πE = πg.Add(πF, πTemp002, πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßascii_letters.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 30: digits = '0123456789'
			πF.SetLineno(30)
			if πE = πF.Globals().SetItem(πF, ßdigits.ToObject(), ß0123456789.ToObject()); πE != nil {
				continue
			}
			// line 31: hexdigits = digits + 'abcdef' + 'ABCDEF'
			πF.SetLineno(31)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßdigits); πE != nil {
				continue
			}
			if πTemp002, πE = πg.Add(πF, πTemp003, ßabcdef.ToObject()); πE != nil {
				continue
			}
			if πTemp001, πE = πg.Add(πF, πTemp002, ßABCDEF.ToObject()); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßhexdigits.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 32: octdigits = '01234567'
			πF.SetLineno(32)
			if πE = πF.Globals().SetItem(πF, ßoctdigits.ToObject(), ß01234567.ToObject()); πE != nil {
				continue
			}
			// line 33: punctuation = """!"#$%&'()*+,-./:;<=>?@[\]^_`{|}~"""
			πF.SetLineno(33)
			if πE = πF.Globals().SetItem(πF, ßpunctuation.ToObject(), πg.NewStr("!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~").ToObject()); πE != nil {
				continue
			}
			// line 34: printable = digits + letters + punctuation + whitespace
			πF.SetLineno(34)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßdigits); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßletters); πE != nil {
				continue
			}
			if πTemp003, πE = πg.Add(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßpunctuation); πE != nil {
				continue
			}
			if πTemp002, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßwhitespace); πE != nil {
				continue
			}
			if πTemp001, πE = πg.Add(πF, πTemp002, πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßprintable.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 40: l = [chr(x) for x in xrange(256)]
			πF.SetLineno(40)
			πTemp006 = make([]πg.Param, 0)
			πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/string.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πg.UnboundLocal; _ = µx
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
						if πTemp003, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
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
							µx = πTemp003
						}
						if πE != nil || !πTemp006 {
							continue
						}
						πF.PushCheckpoint(1)            
						// line 40: l = [chr(x) for x in xrange(256)]
						πF.SetLineno(40)
						πTemp002 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
							continue
						}
						πTemp002[0] = µx
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
			if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp003}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßl.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 41: _idmap = str('').join(l)
			πF.SetLineno(41)
			πTemp007 = πF.MakeArgs(1)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßl); πE != nil {
				continue
			}
			πTemp007[0] = πTemp001
			πTemp008 = πF.MakeArgs(1)
			πTemp008[0] = ß.ToObject()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp008)
			if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßjoin, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp007)
			if πE = πF.Globals().SetItem(πF, ß_idmap.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 42: del l
			πF.SetLineno(42)
			if πE = πg.DelVar(πF, πF.Globals(), ßl); πE != nil {
				continue
			}
			// line 47: def capwords(s, sep=None):
			πF.SetLineno(47)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "sep", Def: πTemp003}
			πTemp001 = πg.NewFunction(πg.NewCode("capwords", "build/src/__python__/string.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µsep *πg.Object = πArgs[1]; _ = µsep
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []πg.Param
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 bool
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
					// line 48: """capwords(s [,sep]) -> string
					πF.SetLineno(48)
					// line 58: return (sep or ' ').join(x.capitalize() for x in s.split(sep))
					πF.SetLineno(58)
					πTemp001 = πF.MakeArgs(1)
					πTemp003 = make([]πg.Param, 0)
					πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/string.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µx *πg.Object = πg.UnboundLocal; _ = µx
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
								if πE = πg.CheckLocal(πF, µsep, "sep"); πE != nil {
									continue
								}
								πTemp002[0] = µsep
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µs, ßsplit, nil); πE != nil {
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
									µx = πTemp003
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 58: return (sep or ' ').join(x.capitalize() for x in s.split(sep))
								πF.SetLineno(58)
								if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µx, ßcapitalize, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
									continue
								}
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
					if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					if πE = πg.CheckLocal(πF, µsep, "sep"); πE != nil {
						continue
					}
					πTemp004 = µsep
					if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					πTemp004 = πg.NewStr(" ").ToObject()
				Label1:
					if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßcapwords.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 62: _idmapL = None
			πF.SetLineno(62)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_idmapL.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 63: def maketrans(fromstr, tostr):
			πF.SetLineno(63)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "fromstr", Def: nil}
			πTemp006[1] = πg.Param{Name: "tostr", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("maketrans", "build/src/__python__/string.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfromstr *πg.Object = πArgs[0]; _ = µfromstr
				var µtostr *πg.Object = πArgs[1]; _ = µtostr
				var µL *πg.Object = πg.UnboundLocal; _ = µL
				var µi *πg.Object = πg.UnboundLocal; _ = µi
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
				var πTemp007 []*πg.Object
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
					case 5: goto Label5
					case 6: goto Label6
					default: panic("unexpected function state")
					}
					// line 64: """maketrans(frm, to) -> string
					πF.SetLineno(64)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfromstr, "fromstr"); πE != nil {
						continue
					}
					πTemp002[0] = µfromstr
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtostr, "tostr"); πE != nil {
						continue
					}
					πTemp002[0] = µtostr
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
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
					// line 71: if len(fromstr) != len(tostr):
					πF.SetLineno(71)
				Label1:
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					// line 72: raise ValueError, "maketrans arguments must have same length"
					πF.SetLineno(72)
					πE = πF.Raise(πTemp001, πg.NewStr("maketrans arguments must have same length").ToObject(), nil)
					continue
					goto Label2
				Label2:
					// line 73: global _idmapL
					πF.SetLineno(73)
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_idmapL); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
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
					// line 74: if not _idmapL:
					πF.SetLineno(74)
				Label3:
					// line 75: _idmapL = list(_idmap)
					πF.SetLineno(75)
					πTemp002 = πF.MakeArgs(1)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_idmap); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πF.Globals().SetItem(πF, ß_idmapL.ToObject(), πTemp003); πE != nil {
						continue
					}
					goto Label4
				Label4:
					// line 76: L = _idmapL[:]
					πF.SetLineno(76)
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_idmapL); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
						continue
					}
					µL = πTemp003
					// line 77: fromstr = map(ord, fromstr)
					πF.SetLineno(77)
					πTemp002 = πF.MakeArgs(2)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πE = πg.CheckLocal(πF, µfromstr, "fromstr"); πE != nil {
						continue
					}
					πTemp002[1] = µfromstr
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µfromstr = πTemp003
					πTemp002 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfromstr, "fromstr"); πE != nil {
						continue
					}
					πTemp007[0] = µfromstr
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp002[0] = πTemp004
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
					πF.PushCheckpoint(6)
					πTemp006 = false
				Label5:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label7
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
						µi = πTemp003
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(5)            
					// line 79: L[fromstr[i]] = tostr[i]
					πF.SetLineno(79)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp003 = µi
					if πE = πg.CheckLocal(πF, µtostr, "tostr"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µtostr, πTemp003); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µL, "L"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp009 = µi
					if πE = πg.CheckLocal(πF, µfromstr, "fromstr"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µfromstr, πTemp009); πE != nil {
						continue
					}
					πTemp005 = πTemp010
					if πE = πg.SetItem(πF, µL, πTemp005, πTemp003); πE != nil {
						continue
					}
					continue
				Label6:
					if πE != nil || πR != nil {
						continue
					}
				Label7:
					// line 80: return ''.join(L)
					πF.SetLineno(80)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µL, "L"); πE != nil {
						continue
					}
					πTemp002[0] = µL
					if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
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
			if πE = πF.Globals().SetItem(πF, ßmaketrans.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 85: import re as _re
			πF.SetLineno(85)
			if πTemp007, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp004 = πTemp007[0]
			if πE = πF.Globals().SetItem(πF, ß_re.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 87: class _multimap(object):
			πF.SetLineno(87)
			πTemp007 = make([]*πg.Object, 1)
			if πTemp010, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp007[0] = πTemp010
			πTemp009 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_multimap", "build/src/__python__/string.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
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
					// line 88: """Helper class for combining multiple mappings.
					πF.SetLineno(88)
					// line 93: def __init__(self, primary, secondary):
					πF.SetLineno(93)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "primary", Def: nil}
					πTemp002[2] = πg.Param{Name: "secondary", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µprimary *πg.Object = πArgs[1]; _ = µprimary
						var µsecondary *πg.Object = πArgs[2]; _ = µsecondary
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 94: self._primary = primary
							πF.SetLineno(94)
							if πE = πg.CheckLocal(πF, µprimary, "primary"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µprimary); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_primary, πTemp001); πE != nil {
								continue
							}
							// line 95: self._secondary = secondary
							πF.SetLineno(95)
							if πE = πg.CheckLocal(πF, µsecondary, "secondary"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µsecondary); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_secondary, πTemp001); πE != nil {
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
					// line 97: def __getitem__(self, key):
					πF.SetLineno(97)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__getitem__", "build/src/__python__/string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var πTemp001 *πg.Object
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 98: try:
							πF.SetLineno(98)
							πF.PushCheckpoint(2)
							// line 99: return self._primary[key]
							πF.SetLineno(99)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001 = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_primary, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
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
							// line 100: except KeyError:
							πF.SetLineno(100)
						Label3:
							// line 101: return self._secondary[key]
							πF.SetLineno(101)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001 = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_secondary, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							πR = πTemp002
							continue
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
					if πE = πClass.SetItem(πF, ß__getitem__.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp005, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp005 == nil {
				πTemp005 = πg.TypeType.ToObject()
			}
			if πTemp010, πE = πTemp005.Call(πF, []*πg.Object{πg.NewStr("_multimap").ToObject(), πg.NewTuple(πTemp007...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_multimap.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 104: class _TemplateMetaclass(type):
			πF.SetLineno(104)
			πTemp007 = make([]*πg.Object, 1)
			if πTemp010, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			πTemp007[0] = πTemp010
			πTemp009 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_TemplateMetaclass", "build/src/__python__/string.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 113: pattern = r"""
					πF.SetLineno(113)
					if πE = πClass.SetItem(πF, ßpattern.ToObject(), πg.NewStr("\n    %s(?:\n      (?P<escaped>%s) |   # Escape sequence of two delimiters\n      (?P<named>%s)      |   # delimiter and a Python identifier\n      {(?P<braced>%s)}   |   # delimiter and a braced identifier\n      (?P<invalid>)              # Other ill-formed delimiter exprs\n    )\n    ").ToObject()); πE != nil {
						continue
					}
					// line 122: def __init__(cls, name, bases, dct):
					πF.SetLineno(122)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp002[2] = πg.Param{Name: "bases", Def: nil}
					πTemp002[3] = πg.Param{Name: "dct", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µname *πg.Object = πArgs[1]; _ = µname
						var µbases *πg.Object = πArgs[2]; _ = µbases
						var µdct *πg.Object = πArgs[3]; _ = µdct
						var µpattern *πg.Object = πg.UnboundLocal; _ = µpattern
						var µcls_delim *πg.Object = πg.UnboundLocal; _ = µcls_delim
						var µcls_id *πg.Object = πg.UnboundLocal; _ = µcls_id
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
							// line 124: super(_TemplateMetaclass, cls)
							πF.SetLineno(124)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_TemplateMetaclass); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp001[1] = µcls
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µdct, "dct"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, µdct, ßpattern.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 125: if 'pattern' in dct:
							πF.SetLineno(125)
						Label1:
							// line 126: pattern = cls.pattern
							πF.SetLineno(126)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcls, ßpattern, nil); πE != nil {
								continue
							}
							µpattern = πTemp002
							goto Label3
						Label2:
							// line 132: cls_delim, cls_id = _re.escape(cls.delimiter), cls.idpattern
							πF.SetLineno(132)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µcls, ßdelimiter, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_re); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßescape, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µcls, ßidpattern, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
								continue
							}
							µcls_delim = πTemp003
							µcls_id = πTemp005
							// line 133: pattern = _TemplateMetaclass.pattern % (cls_delim, cls_delim, cls_id, cls_id)
							πF.SetLineno(133)
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_TemplateMetaclass); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßpattern, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcls_delim, "cls_delim"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcls_delim, "cls_delim"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcls_id, "cls_id"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcls_id, "cls_id"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple4(µcls_delim, µcls_delim, µcls_id, µcls_id).ToObject()
							if πTemp002, πE = πg.Mod(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							µpattern = πTemp002
							goto Label3
						Label3:
							// line 134: cls.pattern = _re.compile(pattern, _re.IGNORECASE | _re.VERBOSE)
							πF.SetLineno(134)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
								continue
							}
							πTemp001[0] = µpattern
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_re); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßIGNORECASE, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_re); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßVERBOSE, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Or(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_re); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcompile, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µcls, ßpattern, πTemp003); πE != nil {
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
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp005, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp005 == nil {
				πTemp005 = πg.TypeType.ToObject()
			}
			if πTemp010, πE = πTemp005.Call(πF, []*πg.Object{πg.NewStr("_TemplateMetaclass").ToObject(), πg.NewTuple(πTemp007...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_TemplateMetaclass.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 137: class Template(object):
			πF.SetLineno(137)
			πTemp007 = make([]*πg.Object, 1)
			if πTemp010, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp007[0] = πTemp010
			πTemp009 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Template", "build/src/__python__/string.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 138: """A string class for supporting $-substitutions."""
					πF.SetLineno(138)
					// line 139: __metaclass__ = _TemplateMetaclass
					πF.SetLineno(139)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ß_TemplateMetaclass); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__metaclass__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 141: delimiter = '$'
					πF.SetLineno(141)
					if πE = πClass.SetItem(πF, ßdelimiter.ToObject(), πg.NewStr("$").ToObject()); πE != nil {
						continue
					}
					// line 142: idpattern = r'[_a-z][_a-z0-9]*'
					πF.SetLineno(142)
					if πE = πClass.SetItem(πF, ßidpattern.ToObject(), πg.NewStr("[_a-z][_a-z0-9]*").ToObject()); πE != nil {
						continue
					}
					// line 144: def __init__(self, template, *arg):
					πF.SetLineno(144)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "template", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/string.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtemplate *πg.Object = πArgs[1]; _ = µtemplate
						var µarg *πg.Object = πArgs[2]; _ = µarg
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 145: self.template = template
							πF.SetLineno(145)
							if πE = πg.CheckLocal(πF, µtemplate, "template"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µtemplate); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtemplate, πTemp001); πE != nil {
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
					// line 149: def _invalid(self, mo):
					πF.SetLineno(149)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "mo", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("_invalid", "build/src/__python__/string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmo *πg.Object = πArgs[1]; _ = µmo
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µlines *πg.Object = πg.UnboundLocal; _ = µlines
						var µcolno *πg.Object = πg.UnboundLocal; _ = µcolno
						var µlineno *πg.Object = πg.UnboundLocal; _ = µlineno
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
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
							// line 150: i = mo.start('invalid')
							πF.SetLineno(150)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßinvalid.ToObject()
							if πE = πg.CheckLocal(πF, µmo, "mo"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µmo, ßstart, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µi = πTemp003
							// line 151: lines = self.template[:i].splitlines(True)
							πF.SetLineno(151)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µi, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtemplate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßsplitlines, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µlines = πTemp003
							if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µlines); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 152: if not lines:
							πF.SetLineno(152)
						Label1:
							// line 153: colno = 1
							πF.SetLineno(153)
							µcolno = πg.NewInt(1).ToObject()
							// line 154: lineno = 1
							πF.SetLineno(154)
							µlineno = πg.NewInt(1).ToObject()
							goto Label3
						Label2:
							// line 156: colno = i - len(''.join(lines[:-1]))
							πF.SetLineno(156)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp004, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µlines, πTemp003); πE != nil {
								continue
							}
							πTemp006[0] = πTemp004
							if πTemp003, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Sub(πF, µi, πTemp004); πE != nil {
								continue
							}
							µcolno = πTemp002
							// line 157: lineno = len(lines)
							πF.SetLineno(157)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
								continue
							}
							πTemp001[0] = µlines
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µlineno = πTemp003
							goto Label3
						Label3:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcolno, "colno"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(µlineno, µcolno).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Invalid placeholder in string: line %d, col %d").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 158: raise ValueError('Invalid placeholder in string: line %d, col %d' %
							πF.SetLineno(158)
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
					if πE = πClass.SetItem(πF, ß_invalid.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 161: def substitute(*args, **kws):
					πF.SetLineno(161)
					πTemp002 = make([]πg.Param, 0)
					πTemp004 = πg.NewFunction(πg.NewCode("substitute", "build/src/__python__/string.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µargs *πg.Object = πArgs[0]; _ = µargs
						var µkws *πg.Object = πArgs[1]; _ = µkws
						var µself *πg.Object = πg.UnboundLocal; _ = µself
						var µmapping *πg.Object = πg.UnboundLocal; _ = µmapping
						var µconvert *πg.Object = πg.UnboundLocal; _ = µconvert
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
						var πTemp007 []πg.Param
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µargs); πE != nil {
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
							// line 162: if not args:
							πF.SetLineno(162)
						Label1:
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("descriptor 'substitute' of 'Template' object needs an argument").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 163: raise TypeError("descriptor 'substitute' of 'Template' object "
							πF.SetLineno(163)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							// line 165: self, args = args[0], args[1:]  # allow the "self" keyword be passed
							πF.SetLineno(165)
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µargs, πTemp004); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µargs, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp001); πE != nil {
								continue
							}
							µself = πTemp004
							µargs = πTemp005
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp003[0] = µargs
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.GT(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 166: if len(args) > 1:
							πF.SetLineno(166)
						Label3:
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("Too many positional arguments").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 167: raise TypeError('Too many positional arguments')
							πF.SetLineno(167)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µargs); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label5
							}
							if πE = πg.CheckLocal(πF, µkws, "kws"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µkws); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label6
							}
							goto Label7
							// line 168: if not args:
							πF.SetLineno(168)
						Label5:
							// line 169: mapping = kws
							πF.SetLineno(169)
							if πE = πg.CheckLocal(πF, µkws, "kws"); πE != nil {
								continue
							}
							µmapping = µkws
							goto Label8
							// line 170: elif kws:
							πF.SetLineno(170)
						Label6:
							// line 171: mapping = _multimap(kws, args[0])
							πF.SetLineno(171)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µkws, "kws"); πE != nil {
								continue
							}
							πTemp003[0] = µkws
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_multimap); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µmapping = πTemp004
							goto Label8
						Label7:
							// line 173: mapping = args[0]
							πF.SetLineno(173)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							µmapping = πTemp004
							goto Label8
						Label8:
							// line 175: def convert(mo):
							πF.SetLineno(175)
							πTemp007 = make([]πg.Param, 1)
							πTemp007[0] = πg.Param{Name: "mo", Def: nil}
							πTemp001 = πg.NewFunction(πg.NewCode("convert", "build/src/__python__/string.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µmo *πg.Object = πArgs[0]; _ = µmo
								var µnamed *πg.Object = πg.UnboundLocal; _ = µnamed
								var µval *πg.Object = πg.UnboundLocal; _ = µval
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
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 177: named = mo.group('named') or mo.group('braced')
									πF.SetLineno(177)
									πTemp003 = πF.MakeArgs(1)
									πTemp003[0] = ßnamed.ToObject()
									if πE = πg.CheckLocal(πF, µmo, "mo"); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetAttr(πF, µmo, ßgroup, nil); πE != nil {
										continue
									}
									if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									πTemp001 = πTemp005
									if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp002 {
										goto Label1
									}
									πTemp003 = πF.MakeArgs(1)
									πTemp003[0] = ßbraced.ToObject()
									if πE = πg.CheckLocal(πF, µmo, "mo"); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetAttr(πF, µmo, ßgroup, nil); πE != nil {
										continue
									}
									if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									πTemp001 = πTemp005
								Label1:
									µnamed = πTemp001
									if πE = πg.CheckLocal(πF, µnamed, "named"); πE != nil {
										continue
									}
									if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp001 = πg.GetBool(µnamed != πTemp004).ToObject()
									if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp002 {
										goto Label2
									}
									goto Label3
									// line 178: if named is not None:
									πF.SetLineno(178)
								Label2:
									// line 179: val = mapping[named]
									πF.SetLineno(179)
									if πE = πg.CheckLocal(πF, µnamed, "named"); πE != nil {
										continue
									}
									πTemp001 = µnamed
									if πE = πg.CheckLocal(πF, µmapping, "mapping"); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetItem(πF, µmapping, πTemp001); πE != nil {
										continue
									}
									µval = πTemp004
									// line 182: return '%s' % (val,)
									πF.SetLineno(182)
									if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
										continue
									}
									πTemp004 = πg.NewTuple1(µval).ToObject()
									if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s").ToObject(), πTemp004); πE != nil {
										continue
									}
									πR = πTemp001
									continue
									goto Label3
								Label3:
									πTemp003 = πF.MakeArgs(1)
									πTemp003[0] = ßescaped.ToObject()
									if πE = πg.CheckLocal(πF, µmo, "mo"); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetAttr(πF, µmo, ßgroup, nil); πE != nil {
										continue
									}
									if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp001 = πg.GetBool(πTemp005 != πTemp004).ToObject()
									if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp002 {
										goto Label4
									}
									goto Label5
									// line 183: if mo.group('escaped') is not None:
									πF.SetLineno(183)
								Label4:
									// line 184: return self.delimiter
									πF.SetLineno(184)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ßdelimiter, nil); πE != nil {
										continue
									}
									πR = πTemp001
									continue
									goto Label5
								Label5:
									πTemp003 = πF.MakeArgs(1)
									πTemp003[0] = ßinvalid.ToObject()
									if πE = πg.CheckLocal(πF, µmo, "mo"); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetAttr(πF, µmo, ßgroup, nil); πE != nil {
										continue
									}
									if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp001 = πg.GetBool(πTemp005 != πTemp004).ToObject()
									if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp002 {
										goto Label6
									}
									goto Label7
									// line 185: if mo.group('invalid') is not None:
									πF.SetLineno(185)
								Label6:
									// line 186: self._invalid(mo)
									πF.SetLineno(186)
									πTemp003 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µmo, "mo"); πE != nil {
										continue
									}
									πTemp003[0] = µmo
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ß_invalid, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									goto Label7
								Label7:
									πTemp003 = πF.MakeArgs(2)
									πTemp003[0] = πg.NewStr("Unrecognized named group in pattern").ToObject()
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ßpattern, nil); πE != nil {
										continue
									}
									πTemp003[1] = πTemp001
									if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									// line 187: raise ValueError('Unrecognized named group in pattern',
									πF.SetLineno(187)
									πE = πF.Raise(πTemp004, nil, nil)
									continue
								}
								if πE != nil {
									πR = nil
								} else if πR == nil {
									πR = πg.None
								}
								return πR, πE
							}), πF.Globals()).ToObject()
							µconvert = πTemp001
							// line 189: return self.pattern.sub(convert, self.template)
							πF.SetLineno(189)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µconvert, "convert"); πE != nil {
								continue
							}
							πTemp003[0] = µconvert
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtemplate, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßpattern, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßsub, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
					if πE = πClass.SetItem(πF, ßsubstitute.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 191: def safe_substitute(*args, **kws):
					πF.SetLineno(191)
					πTemp002 = make([]πg.Param, 0)
					πTemp005 = πg.NewFunction(πg.NewCode("safe_substitute", "build/src/__python__/string.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µargs *πg.Object = πArgs[0]; _ = µargs
						var µkws *πg.Object = πArgs[1]; _ = µkws
						var µself *πg.Object = πg.UnboundLocal; _ = µself
						var µmapping *πg.Object = πg.UnboundLocal; _ = µmapping
						var µconvert *πg.Object = πg.UnboundLocal; _ = µconvert
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
						var πTemp007 []πg.Param
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µargs); πE != nil {
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
							// line 192: if not args:
							πF.SetLineno(192)
						Label1:
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("descriptor 'safe_substitute' of 'Template' object needs an argument").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 193: raise TypeError("descriptor 'safe_substitute' of 'Template' object "
							πF.SetLineno(193)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							// line 195: self, args = args[0], args[1:]  # allow the "self" keyword be passed
							πF.SetLineno(195)
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µargs, πTemp004); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µargs, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp001); πE != nil {
								continue
							}
							µself = πTemp004
							µargs = πTemp005
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp003[0] = µargs
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.GT(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 196: if len(args) > 1:
							πF.SetLineno(196)
						Label3:
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("Too many positional arguments").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 197: raise TypeError('Too many positional arguments')
							πF.SetLineno(197)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µargs); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label5
							}
							if πE = πg.CheckLocal(πF, µkws, "kws"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µkws); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label6
							}
							goto Label7
							// line 198: if not args:
							πF.SetLineno(198)
						Label5:
							// line 199: mapping = kws
							πF.SetLineno(199)
							if πE = πg.CheckLocal(πF, µkws, "kws"); πE != nil {
								continue
							}
							µmapping = µkws
							goto Label8
							// line 200: elif kws:
							πF.SetLineno(200)
						Label6:
							// line 201: mapping = _multimap(kws, args[0])
							πF.SetLineno(201)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µkws, "kws"); πE != nil {
								continue
							}
							πTemp003[0] = µkws
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_multimap); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µmapping = πTemp004
							goto Label8
						Label7:
							// line 203: mapping = args[0]
							πF.SetLineno(203)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							µmapping = πTemp004
							goto Label8
						Label8:
							// line 205: def convert(mo):
							πF.SetLineno(205)
							πTemp007 = make([]πg.Param, 1)
							πTemp007[0] = πg.Param{Name: "mo", Def: nil}
							πTemp001 = πg.NewFunction(πg.NewCode("convert", "build/src/__python__/string.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µmo *πg.Object = πArgs[0]; _ = µmo
								var µnamed *πg.Object = πg.UnboundLocal; _ = µnamed
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
								var πTemp007 *πg.BaseException
								_ = πTemp007
								var πTemp008 *πg.Traceback
								_ = πTemp008
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									case 5: goto Label5
									default: panic("unexpected function state")
									}
									// line 206: named = mo.group('named') or mo.group('braced')
									πF.SetLineno(206)
									πTemp003 = πF.MakeArgs(1)
									πTemp003[0] = ßnamed.ToObject()
									if πE = πg.CheckLocal(πF, µmo, "mo"); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetAttr(πF, µmo, ßgroup, nil); πE != nil {
										continue
									}
									if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									πTemp001 = πTemp005
									if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp002 {
										goto Label1
									}
									πTemp003 = πF.MakeArgs(1)
									πTemp003[0] = ßbraced.ToObject()
									if πE = πg.CheckLocal(πF, µmo, "mo"); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetAttr(πF, µmo, ßgroup, nil); πE != nil {
										continue
									}
									if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									πTemp001 = πTemp005
								Label1:
									µnamed = πTemp001
									if πE = πg.CheckLocal(πF, µnamed, "named"); πE != nil {
										continue
									}
									if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp001 = πg.GetBool(µnamed != πTemp004).ToObject()
									if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp002 {
										goto Label2
									}
									goto Label3
									// line 207: if named is not None:
									πF.SetLineno(207)
								Label2:
									// line 208: try:
									πF.SetLineno(208)
									πF.PushCheckpoint(5)
									// line 211: return '%s' % (mapping[named],)
									πF.SetLineno(211)
									if πE = πg.CheckLocal(πF, µnamed, "named"); πE != nil {
										continue
									}
									πTemp005 = µnamed
									if πE = πg.CheckLocal(πF, µmapping, "mapping"); πE != nil {
										continue
									}
									if πTemp006, πE = πg.GetItem(πF, µmapping, πTemp005); πE != nil {
										continue
									}
									πTemp004 = πg.NewTuple1(πTemp006).ToObject()
									if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s").ToObject(), πTemp004); πE != nil {
										continue
									}
									πR = πTemp001
									continue
									πF.PopCheckpoint()
									goto Label4
								Label5:
									if πE == nil {
									  continue
									}
									πE = nil
									πTemp007, πTemp008 = πF.ExcInfo()
									if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
										continue
									}
									if πTemp002, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
										continue
									}
									if πTemp002 {
										goto Label6
									}
									πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
									continue
									// line 212: except KeyError:
									πF.SetLineno(212)
								Label6:
									// line 213: return mo.group()
									πF.SetLineno(213)
									if πE = πg.CheckLocal(πF, µmo, "mo"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µmo, ßgroup, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
										continue
									}
									πR = πTemp004
									continue
									πF.RestoreExc(nil, nil)
									goto Label4
								Label4:
									goto Label3
								Label3:
									πTemp003 = πF.MakeArgs(1)
									πTemp003[0] = ßescaped.ToObject()
									if πE = πg.CheckLocal(πF, µmo, "mo"); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetAttr(πF, µmo, ßgroup, nil); πE != nil {
										continue
									}
									if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp001 = πg.GetBool(πTemp005 != πTemp004).ToObject()
									if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp002 {
										goto Label7
									}
									goto Label8
									// line 214: if mo.group('escaped') is not None:
									πF.SetLineno(214)
								Label7:
									// line 215: return self.delimiter
									πF.SetLineno(215)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ßdelimiter, nil); πE != nil {
										continue
									}
									πR = πTemp001
									continue
									goto Label8
								Label8:
									πTemp003 = πF.MakeArgs(1)
									πTemp003[0] = ßinvalid.ToObject()
									if πE = πg.CheckLocal(πF, µmo, "mo"); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetAttr(πF, µmo, ßgroup, nil); πE != nil {
										continue
									}
									if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp001 = πg.GetBool(πTemp005 != πTemp004).ToObject()
									if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp002 {
										goto Label9
									}
									goto Label10
									// line 216: if mo.group('invalid') is not None:
									πF.SetLineno(216)
								Label9:
									// line 217: return mo.group()
									πF.SetLineno(217)
									if πE = πg.CheckLocal(πF, µmo, "mo"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µmo, ßgroup, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
										continue
									}
									πR = πTemp004
									continue
									goto Label10
								Label10:
									πTemp003 = πF.MakeArgs(2)
									πTemp003[0] = πg.NewStr("Unrecognized named group in pattern").ToObject()
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ßpattern, nil); πE != nil {
										continue
									}
									πTemp003[1] = πTemp001
									if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									// line 218: raise ValueError('Unrecognized named group in pattern',
									πF.SetLineno(218)
									πE = πF.Raise(πTemp004, nil, nil)
									continue
								}
								if πE != nil {
									πR = nil
								} else if πR == nil {
									πR = πg.None
								}
								return πR, πE
							}), πF.Globals()).ToObject()
							µconvert = πTemp001
							// line 220: return self.pattern.sub(convert, self.template)
							πF.SetLineno(220)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µconvert, "convert"); πE != nil {
								continue
							}
							πTemp003[0] = µconvert
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtemplate, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßpattern, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßsub, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
					if πE = πClass.SetItem(πF, ßsafe_substitute.ToObject(), πTemp005); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp005, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp005 == nil {
				πTemp005 = πg.TypeType.ToObject()
			}
			if πTemp010, πE = πTemp005.Call(πF, []*πg.Object{πg.NewStr("Template").ToObject(), πg.NewTuple(πTemp007...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTemplate.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 229: index_error = ValueError
			πF.SetLineno(229)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßindex_error.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 230: atoi_error = ValueError
			πF.SetLineno(230)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßatoi_error.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 231: atof_error = ValueError
			πF.SetLineno(231)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßatof_error.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 232: atol_error = ValueError
			πF.SetLineno(232)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßatol_error.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 235: def lower(s):
			πF.SetLineno(235)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("lower", "build/src/__python__/string.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
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
					// line 236: """lower(s) -> string
					πF.SetLineno(236)
					// line 241: return s.lower()
					πF.SetLineno(241)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µs, ßlower, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßlower.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 244: def upper(s):
			πF.SetLineno(244)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("upper", "build/src/__python__/string.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
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
					// line 245: """upper(s) -> string
					πF.SetLineno(245)
					// line 250: return s.upper()
					πF.SetLineno(250)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µs, ßupper, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßupper.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 253: def swapcase(s):
			πF.SetLineno(253)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("swapcase", "build/src/__python__/string.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
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
					// line 254: """swapcase(s) -> string
					πF.SetLineno(254)
					// line 260: return s.swapcase()
					πF.SetLineno(260)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µs, ßswapcase, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßswapcase.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 263: def strip(s, chars=None):
			πF.SetLineno(263)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "chars", Def: πTemp012}
			πTemp011 = πg.NewFunction(πg.NewCode("strip", "build/src/__python__/string.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µchars *πg.Object = πArgs[1]; _ = µchars
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
					// line 264: """strip(s [,chars]) -> string
					πF.SetLineno(264)
					// line 272: return s.strip(chars)
					πF.SetLineno(272)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchars, "chars"); πE != nil {
						continue
					}
					πTemp001[0] = µchars
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µs, ßstrip, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßstrip.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 275: def lstrip(s, chars=None):
			πF.SetLineno(275)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "chars", Def: πTemp013}
			πTemp012 = πg.NewFunction(πg.NewCode("lstrip", "build/src/__python__/string.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µchars *πg.Object = πArgs[1]; _ = µchars
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
					// line 276: """lstrip(s [,chars]) -> string
					πF.SetLineno(276)
					// line 282: return s.lstrip(chars)
					πF.SetLineno(282)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchars, "chars"); πE != nil {
						continue
					}
					πTemp001[0] = µchars
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µs, ßlstrip, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßlstrip.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 285: def rstrip(s, chars=None):
			πF.SetLineno(285)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			if πTemp014, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "chars", Def: πTemp014}
			πTemp013 = πg.NewFunction(πg.NewCode("rstrip", "build/src/__python__/string.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µchars *πg.Object = πArgs[1]; _ = µchars
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
					// line 286: """rstrip(s [,chars]) -> string
					πF.SetLineno(286)
					// line 292: return s.rstrip(chars)
					πF.SetLineno(292)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchars, "chars"); πE != nil {
						continue
					}
					πTemp001[0] = µchars
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µs, ßrstrip, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßrstrip.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 296: def split(s, sep=None, maxsplit=-1):
			πF.SetLineno(296)
			πTemp006 = make([]πg.Param, 3)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "sep", Def: πTemp015}
			if πTemp015, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp006[2] = πg.Param{Name: "maxsplit", Def: πTemp015}
			πTemp014 = πg.NewFunction(πg.NewCode("split", "build/src/__python__/string.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µsep *πg.Object = πArgs[1]; _ = µsep
				var µmaxsplit *πg.Object = πArgs[2]; _ = µmaxsplit
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
					// line 297: """split(s [,sep [,maxsplit]]) -> list of strings
					πF.SetLineno(297)
					// line 307: return s.split(sep, maxsplit)
					πF.SetLineno(307)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µsep, "sep"); πE != nil {
						continue
					}
					πTemp001[0] = µsep
					if πE = πg.CheckLocal(πF, µmaxsplit, "maxsplit"); πE != nil {
						continue
					}
					πTemp001[1] = µmaxsplit
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µs, ßsplit, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßsplit.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 308: splitfields = split
			πF.SetLineno(308)
			if πTemp015, πE = πg.ResolveGlobal(πF, ßsplit); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßsplitfields.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 311: def rsplit(s, sep=None, maxsplit=-1):
			πF.SetLineno(311)
			πTemp006 = make([]πg.Param, 3)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "sep", Def: πTemp016}
			if πTemp016, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp006[2] = πg.Param{Name: "maxsplit", Def: πTemp016}
			πTemp015 = πg.NewFunction(πg.NewCode("rsplit", "build/src/__python__/string.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µsep *πg.Object = πArgs[1]; _ = µsep
				var µmaxsplit *πg.Object = πArgs[2]; _ = µmaxsplit
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
					// line 312: """rsplit(s [,sep [,maxsplit]]) -> list of strings
					πF.SetLineno(312)
					// line 320: return s.rsplit(sep, maxsplit)
					πF.SetLineno(320)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µsep, "sep"); πE != nil {
						continue
					}
					πTemp001[0] = µsep
					if πE = πg.CheckLocal(πF, µmaxsplit, "maxsplit"); πE != nil {
						continue
					}
					πTemp001[1] = µmaxsplit
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µs, ßrsplit, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßrsplit.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 323: def join(words, sep = ' '):
			πF.SetLineno(323)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "words", Def: nil}
			πTemp006[1] = πg.Param{Name: "sep", Def: πg.NewStr(" ").ToObject()}
			πTemp016 = πg.NewFunction(πg.NewCode("join", "build/src/__python__/string.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µwords *πg.Object = πArgs[0]; _ = µwords
				var µsep *πg.Object = πArgs[1]; _ = µsep
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
					// line 324: """join(list [,sep]) -> string
					πF.SetLineno(324)
					// line 333: return sep.join(words)
					πF.SetLineno(333)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µwords, "words"); πE != nil {
						continue
					}
					πTemp001[0] = µwords
					if πE = πg.CheckLocal(πF, µsep, "sep"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsep, ßjoin, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßjoin.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 334: joinfields = join
			πF.SetLineno(334)
			if πTemp017, πE = πg.ResolveGlobal(πF, ßjoin); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßjoinfields.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 337: def index(s, *args):
			πF.SetLineno(337)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp017 = πg.NewFunction(πg.NewCode("index", "build/src/__python__/string.py", πTemp006, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µargs *πg.Object = πArgs[1]; _ = µargs
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
					// line 338: """index(s, sub [,start [,end]]) -> int
					πF.SetLineno(338)
					// line 343: return s.index(*args)
					πF.SetLineno(343)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µs, ßindex, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Invoke(πF, πTemp001, nil, µargs, nil, nil); πE != nil {
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
			// line 346: def rindex(s, *args):
			πF.SetLineno(346)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp018 = πg.NewFunction(πg.NewCode("rindex", "build/src/__python__/string.py", πTemp006, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µargs *πg.Object = πArgs[1]; _ = µargs
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
					// line 347: """rindex(s, sub [,start [,end]]) -> int
					πF.SetLineno(347)
					// line 352: return s.rindex(*args)
					πF.SetLineno(352)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µs, ßrindex, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Invoke(πF, πTemp001, nil, µargs, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßrindex.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 355: def count(s, *args):
			πF.SetLineno(355)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp019 = πg.NewFunction(πg.NewCode("count", "build/src/__python__/string.py", πTemp006, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µargs *πg.Object = πArgs[1]; _ = µargs
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
					// line 356: """count(s, sub[, start[,end]]) -> int
					πF.SetLineno(356)
					// line 363: return s.count(*args)
					πF.SetLineno(363)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µs, ßcount, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Invoke(πF, πTemp001, nil, µargs, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßcount.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 366: def find(s, *args):
			πF.SetLineno(366)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp020 = πg.NewFunction(πg.NewCode("find", "build/src/__python__/string.py", πTemp006, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µargs *πg.Object = πArgs[1]; _ = µargs
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
					// line 367: """find(s, sub [,start [,end]]) -> in
					πF.SetLineno(367)
					// line 376: return s.find(*args)
					πF.SetLineno(376)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µs, ßfind, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Invoke(πF, πTemp001, nil, µargs, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßfind.ToObject(), πTemp020); πE != nil {
				continue
			}
			// line 379: def rfind(s, *args):
			πF.SetLineno(379)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp021 = πg.NewFunction(πg.NewCode("rfind", "build/src/__python__/string.py", πTemp006, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µargs *πg.Object = πArgs[1]; _ = µargs
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
					// line 380: """rfind(s, sub [,start [,end]]) -> int
					πF.SetLineno(380)
					// line 389: return s.rfind(*args)
					πF.SetLineno(389)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µs, ßrfind, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Invoke(πF, πTemp001, nil, µargs, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßrfind.ToObject(), πTemp021); πE != nil {
				continue
			}
			// line 392: _float = float
			πF.SetLineno(392)
			if πTemp022, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_float.ToObject(), πTemp022); πE != nil {
				continue
			}
			// line 393: _int = int
			πF.SetLineno(393)
			if πTemp022, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_int.ToObject(), πTemp022); πE != nil {
				continue
			}
			// line 394: _long = long
			πF.SetLineno(394)
			if πTemp022, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_long.ToObject(), πTemp022); πE != nil {
				continue
			}
			// line 397: def atof(s):
			πF.SetLineno(397)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp022 = πg.NewFunction(πg.NewCode("atof", "build/src/__python__/string.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 398: """atof(s) -> float
					πF.SetLineno(398)
					// line 403: return _float(s)
					πF.SetLineno(403)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_float); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßatof.ToObject(), πTemp022); πE != nil {
				continue
			}
			// line 407: def atoi(s , base=10):
			πF.SetLineno(407)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp006[1] = πg.Param{Name: "base", Def: πg.NewInt(10).ToObject()}
			πTemp023 = πg.NewFunction(πg.NewCode("atoi", "build/src/__python__/string.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µbase *πg.Object = πArgs[1]; _ = µbase
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
					// line 408: """atoi(s [,base]) -> int
					πF.SetLineno(408)
					// line 418: return _int(s, base)
					πF.SetLineno(418)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
						continue
					}
					πTemp001[1] = µbase
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_int); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßatoi.ToObject(), πTemp023); πE != nil {
				continue
			}
			// line 422: def atol(s, base=10):
			πF.SetLineno(422)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp006[1] = πg.Param{Name: "base", Def: πg.NewInt(10).ToObject()}
			πTemp024 = πg.NewFunction(πg.NewCode("atol", "build/src/__python__/string.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µbase *πg.Object = πArgs[1]; _ = µbase
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
					// line 423: """atol(s [,base]) -> long
					πF.SetLineno(423)
					// line 434: return _long(s, base)
					πF.SetLineno(434)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
						continue
					}
					πTemp001[1] = µbase
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_long); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßatol.ToObject(), πTemp024); πE != nil {
				continue
			}
			// line 438: def ljust(s, width, *args):
			πF.SetLineno(438)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp006[1] = πg.Param{Name: "width", Def: nil}
			πTemp025 = πg.NewFunction(πg.NewCode("ljust", "build/src/__python__/string.py", πTemp006, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µwidth *πg.Object = πArgs[1]; _ = µwidth
				var µargs *πg.Object = πArgs[2]; _ = µargs
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
					// line 439: """ljust(s, width[, fillchar]) -> string
					πF.SetLineno(439)
					// line 446: return s.ljust(width, *args)
					πF.SetLineno(446)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
						continue
					}
					πTemp001[0] = µwidth
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µs, ßljust, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp001, µargs, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßljust.ToObject(), πTemp025); πE != nil {
				continue
			}
			// line 449: def rjust(s, width, *args):
			πF.SetLineno(449)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp006[1] = πg.Param{Name: "width", Def: nil}
			πTemp026 = πg.NewFunction(πg.NewCode("rjust", "build/src/__python__/string.py", πTemp006, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µwidth *πg.Object = πArgs[1]; _ = µwidth
				var µargs *πg.Object = πArgs[2]; _ = µargs
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
					// line 450: """rjust(s, width[, fillchar]) -> string
					πF.SetLineno(450)
					// line 457: return s.rjust(width, *args)
					πF.SetLineno(457)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
						continue
					}
					πTemp001[0] = µwidth
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µs, ßrjust, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp001, µargs, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßrjust.ToObject(), πTemp026); πE != nil {
				continue
			}
			// line 460: def center(s, width, *args):
			πF.SetLineno(460)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp006[1] = πg.Param{Name: "width", Def: nil}
			πTemp027 = πg.NewFunction(πg.NewCode("center", "build/src/__python__/string.py", πTemp006, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µwidth *πg.Object = πArgs[1]; _ = µwidth
				var µargs *πg.Object = πArgs[2]; _ = µargs
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
					// line 461: """center(s, width[, fillchar]) -> string
					πF.SetLineno(461)
					// line 468: return s.center(width, *args)
					πF.SetLineno(468)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
						continue
					}
					πTemp001[0] = µwidth
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µs, ßcenter, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp001, µargs, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßcenter.ToObject(), πTemp027); πE != nil {
				continue
			}
			// line 473: def zfill(x, width):
			πF.SetLineno(473)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "x", Def: nil}
			πTemp006[1] = πg.Param{Name: "width", Def: nil}
			πTemp028 = πg.NewFunction(πg.NewCode("zfill", "build/src/__python__/string.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µwidth *πg.Object = πArgs[1]; _ = µwidth
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
					// line 474: """zfill(x, width) -> string
					πF.SetLineno(474)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
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
					// line 480: if not isinstance(x, basestring):
					πF.SetLineno(480)
				Label1:
					// line 481: x = repr(x)
					πF.SetLineno(481)
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
					// line 482: return x.zfill(width)
					πF.SetLineno(482)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
						continue
					}
					πTemp002[0] = µwidth
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µx, ßzfill, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
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
			if πE = πF.Globals().SetItem(πF, ßzfill.ToObject(), πTemp028); πE != nil {
				continue
			}
			// line 486: def expandtabs(s, tabsize=8):
			πF.SetLineno(486)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp006[1] = πg.Param{Name: "tabsize", Def: πg.NewInt(8).ToObject()}
			πTemp029 = πg.NewFunction(πg.NewCode("expandtabs", "build/src/__python__/string.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µtabsize *πg.Object = πArgs[1]; _ = µtabsize
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
					// line 487: """expandtabs(s [,tabsize]) -> string
					πF.SetLineno(487)
					// line 494: return s.expandtabs(tabsize)
					πF.SetLineno(494)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtabsize, "tabsize"); πE != nil {
						continue
					}
					πTemp001[0] = µtabsize
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µs, ßexpandtabs, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßexpandtabs.ToObject(), πTemp029); πE != nil {
				continue
			}
			// line 497: def translate(s, table, deletions=""):
			πF.SetLineno(497)
			πTemp006 = make([]πg.Param, 3)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp006[1] = πg.Param{Name: "table", Def: nil}
			πTemp006[2] = πg.Param{Name: "deletions", Def: ß.ToObject()}
			πTemp030 = πg.NewFunction(πg.NewCode("translate", "build/src/__python__/string.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µtable *πg.Object = πArgs[1]; _ = µtable
				var µdeletions *πg.Object = πArgs[2]; _ = µdeletions
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 498: """translate(s,table [,deletions]) -> string
					πF.SetLineno(498)
					if πE = πg.CheckLocal(πF, µdeletions, "deletions"); πE != nil {
						continue
					}
					πTemp001 = µdeletions
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µtable, "table"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µtable == πTemp004).ToObject()
					πTemp001 = πTemp003
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 507: if deletions or table is None:
					πF.SetLineno(507)
				Label2:
					// line 508: return s.translate(table, deletions)
					πF.SetLineno(508)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µtable, "table"); πE != nil {
						continue
					}
					πTemp005[0] = µtable
					if πE = πg.CheckLocal(πF, µdeletions, "deletions"); πE != nil {
						continue
					}
					πTemp005[1] = µdeletions
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µs, ßtranslate, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πR = πTemp003
					continue
					goto Label4
				Label3:
					// line 513: return s.translate(table + s[:0])
					πF.SetLineno(513)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtable, "table"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(0).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µs, πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µtable, πTemp004); πE != nil {
						continue
					}
					πTemp005[0] = πTemp001
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µs, ßtranslate, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πR = πTemp003
					continue
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
			if πE = πF.Globals().SetItem(πF, ßtranslate.ToObject(), πTemp030); πE != nil {
				continue
			}
			// line 516: def capitalize(s):
			πF.SetLineno(516)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp031 = πg.NewFunction(πg.NewCode("capitalize", "build/src/__python__/string.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
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
					// line 517: """capitalize(s) -> string
					πF.SetLineno(517)
					// line 523: return s.capitalize()
					πF.SetLineno(523)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µs, ßcapitalize, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßcapitalize.ToObject(), πTemp031); πE != nil {
				continue
			}
			// line 526: def replace(s, old, new, maxreplace=-1):
			πF.SetLineno(526)
			πTemp006 = make([]πg.Param, 4)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp006[1] = πg.Param{Name: "old", Def: nil}
			πTemp006[2] = πg.Param{Name: "new", Def: nil}
			if πTemp033, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp006[3] = πg.Param{Name: "maxreplace", Def: πTemp033}
			πTemp032 = πg.NewFunction(πg.NewCode("replace", "build/src/__python__/string.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µold *πg.Object = πArgs[1]; _ = µold
				var µnew *πg.Object = πArgs[2]; _ = µnew
				var µmaxreplace *πg.Object = πArgs[3]; _ = µmaxreplace
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
					// line 527: """replace (str, old, new[, maxreplace]) -> string
					πF.SetLineno(527)
					// line 534: return s.replace(old, new, maxreplace)
					πF.SetLineno(534)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µold, "old"); πE != nil {
						continue
					}
					πTemp001[0] = µold
					if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
						continue
					}
					πTemp001[1] = µnew
					if πE = πg.CheckLocal(πF, µmaxreplace, "maxreplace"); πE != nil {
						continue
					}
					πTemp001[2] = µmaxreplace
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µs, ßreplace, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßreplace.ToObject(), πTemp032); πE != nil {
				continue
			}
			// line 558: class Formatter(object):
			πF.SetLineno(558)
			πTemp007 = make([]*πg.Object, 1)
			if πTemp035, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp007[0] = πTemp035
			πTemp009 = πg.NewDict()
			if πTemp033, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp033); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Formatter", "build/src/__python__/string.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
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
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 559: def format(*args, **kwargs):
					πF.SetLineno(559)
					πTemp002 = make([]πg.Param, 0)
					πTemp001 = πg.NewFunction(πg.NewCode("format", "build/src/__python__/string.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µargs *πg.Object = πArgs[0]; _ = µargs
						var µkwargs *πg.Object = πArgs[1]; _ = µkwargs
						var µself *πg.Object = πg.UnboundLocal; _ = µself
						var µformat_string *πg.Object = πg.UnboundLocal; _ = µformat_string
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
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µargs); πE != nil {
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
							// line 560: if not args:
							πF.SetLineno(560)
						Label1:
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("descriptor 'format' of 'Formatter' object needs an argument").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 561: raise TypeError("descriptor 'format' of 'Formatter' object "
							πF.SetLineno(561)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							// line 563: self, args = args[0], args[1:]  # allow the "self" keyword be passed
							πF.SetLineno(563)
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µargs, πTemp004); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µargs, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp001); πE != nil {
								continue
							}
							µself = πTemp004
							µargs = πTemp005
							// line 564: try:
							πF.SetLineno(564)
							πF.PushCheckpoint(4)
							// line 565: format_string, args = args[0], args[1:] # allow the "format_string" keyword be passed
							πF.SetLineno(565)
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µargs, πTemp004); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µargs, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp001); πE != nil {
								continue
							}
							µformat_string = πTemp004
							µargs = πTemp005
							πF.PopCheckpoint()
							goto Label3
						Label4:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label5
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 566: except IndexError:
							πF.SetLineno(566)
						Label5:
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, µkwargs, ßformat_string.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label6
							}
							goto Label7
							// line 567: if 'format_string' in kwargs:
							πF.SetLineno(567)
						Label6:
							// line 568: format_string = kwargs.pop('format_string')
							πF.SetLineno(568)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = ßformat_string.ToObject()
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µkwargs, ßpop, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µformat_string = πTemp004
							goto Label8
						Label7:
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("format() missing 1 required positional argument: 'format_string'").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 570: raise TypeError("format() missing 1 required positional "
							πF.SetLineno(570)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label8
						Label8:
							πF.RestoreExc(nil, nil)
							goto Label3
						Label3:
							// line 572: return self.vformat(format_string, args, kwargs)
							πF.SetLineno(572)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µformat_string, "format_string"); πE != nil {
								continue
							}
							πTemp003[0] = µformat_string
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp003[1] = µargs
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							πTemp003[2] = µkwargs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßvformat, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
					if πE = πClass.SetItem(πF, ßformat.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 574: def vformat(self, format_string, args, kwargs):
					πF.SetLineno(574)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "format_string", Def: nil}
					πTemp002[2] = πg.Param{Name: "args", Def: nil}
					πTemp002[3] = πg.Param{Name: "kwargs", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("vformat", "build/src/__python__/string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µformat_string *πg.Object = πArgs[1]; _ = µformat_string
						var µargs *πg.Object = πArgs[2]; _ = µargs
						var µkwargs *πg.Object = πArgs[3]; _ = µkwargs
						var µused_args *πg.Object = πg.UnboundLocal; _ = µused_args
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 575: used_args = set()
							πF.SetLineno(575)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µused_args = πTemp002
							// line 576: result = self._vformat(format_string, args, kwargs, used_args, 2)
							πF.SetLineno(576)
							πTemp003 = πF.MakeArgs(5)
							if πE = πg.CheckLocal(πF, µformat_string, "format_string"); πE != nil {
								continue
							}
							πTemp003[0] = µformat_string
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp003[1] = µargs
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							πTemp003[2] = µkwargs
							if πE = πg.CheckLocal(πF, µused_args, "used_args"); πE != nil {
								continue
							}
							πTemp003[3] = µused_args
							πTemp003[4] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_vformat, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µresult = πTemp002
							// line 577: self.check_unused_args(used_args, args, kwargs)
							πF.SetLineno(577)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µused_args, "used_args"); πE != nil {
								continue
							}
							πTemp003[0] = µused_args
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp003[1] = µargs
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							πTemp003[2] = µkwargs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcheck_unused_args, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 578: return result
							πF.SetLineno(578)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πR = µresult
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvformat.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 580: def _vformat(self, format_string, args, kwargs, used_args, recursion_depth):
					πF.SetLineno(580)
					πTemp002 = make([]πg.Param, 6)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "format_string", Def: nil}
					πTemp002[2] = πg.Param{Name: "args", Def: nil}
					πTemp002[3] = πg.Param{Name: "kwargs", Def: nil}
					πTemp002[4] = πg.Param{Name: "used_args", Def: nil}
					πTemp002[5] = πg.Param{Name: "recursion_depth", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("_vformat", "build/src/__python__/string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µformat_string *πg.Object = πArgs[1]; _ = µformat_string
						var µargs *πg.Object = πArgs[2]; _ = µargs
						var µkwargs *πg.Object = πArgs[3]; _ = µkwargs
						var µused_args *πg.Object = πArgs[4]; _ = µused_args
						var µrecursion_depth *πg.Object = πArgs[5]; _ = µrecursion_depth
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
						var µliteral_text *πg.Object = πg.UnboundLocal; _ = µliteral_text
						var µfield_name *πg.Object = πg.UnboundLocal; _ = µfield_name
						var µformat_spec *πg.Object = πg.UnboundLocal; _ = µformat_spec
						var µconversion *πg.Object = πg.UnboundLocal; _ = µconversion
						var µobj *πg.Object = πg.UnboundLocal; _ = µobj
						var µarg_used *πg.Object = πg.UnboundLocal; _ = µarg_used
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 []*πg.Object
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
							if πE = πg.CheckLocal(πF, µrecursion_depth, "recursion_depth"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µrecursion_depth, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 581: if recursion_depth < 0:
							πF.SetLineno(581)
						Label1:
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("Max string recursion exceeded").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 582: raise ValueError('Max string recursion exceeded')
							πF.SetLineno(582)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							// line 583: result = []
							πF.SetLineno(583)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							µresult = πTemp001
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µformat_string, "format_string"); πE != nil {
								continue
							}
							πTemp003[0] = µformat_string
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßparse, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp002 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002 {
								πF.PopCheckpoint()
								goto Label5
							}
							if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}}}, πTemp004); πE != nil {
									continue
								}
								µliteral_text = πTemp005
								µfield_name = πTemp007
								µformat_spec = πTemp008
								µconversion = πTemp009
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(3)            
							if πE = πg.CheckLocal(πF, µliteral_text, "literal_text"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µliteral_text); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label6
							}
							goto Label7
							// line 588: if literal_text:
							πF.SetLineno(588)
						Label6:
							// line 589: result.append(literal_text)
							πF.SetLineno(589)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µliteral_text, "literal_text"); πE != nil {
								continue
							}
							πTemp003[0] = µliteral_text
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label7
						Label7:
							if πE = πg.CheckLocal(πF, µfield_name, "field_name"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(µfield_name != πTemp005).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label8
							}
							goto Label9
							// line 592: if field_name is not None:
							πF.SetLineno(592)
						Label8:
							// line 598: obj, arg_used = self.get_field(field_name, args, kwargs)
							πF.SetLineno(598)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µfield_name, "field_name"); πE != nil {
								continue
							}
							πTemp003[0] = µfield_name
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp003[1] = µargs
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							πTemp003[2] = µkwargs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßget_field, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp005); πE != nil {
								continue
							}
							µobj = πTemp004
							µarg_used = πTemp007
							// line 599: used_args.add(arg_used)
							πF.SetLineno(599)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µarg_used, "arg_used"); πE != nil {
								continue
							}
							πTemp003[0] = µarg_used
							if πE = πg.CheckLocal(πF, µused_args, "used_args"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µused_args, ßadd, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 602: obj = self.convert_field(obj, conversion)
							πF.SetLineno(602)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp003[0] = µobj
							if πE = πg.CheckLocal(πF, µconversion, "conversion"); πE != nil {
								continue
							}
							πTemp003[1] = µconversion
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßconvert_field, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µobj = πTemp005
							// line 605: format_spec = self._vformat(format_spec, args, kwargs,
							πF.SetLineno(605)
							πTemp003 = πF.MakeArgs(5)
							if πE = πg.CheckLocal(πF, µformat_spec, "format_spec"); πE != nil {
								continue
							}
							πTemp003[0] = µformat_spec
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp003[1] = µargs
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							πTemp003[2] = µkwargs
							if πE = πg.CheckLocal(πF, µused_args, "used_args"); πE != nil {
								continue
							}
							πTemp003[3] = µused_args
							if πE = πg.CheckLocal(πF, µrecursion_depth, "recursion_depth"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Sub(πF, µrecursion_depth, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[4] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_vformat, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µformat_spec = πTemp005
							// line 609: result.append(self.format_field(obj, format_spec))
							πF.SetLineno(609)
							πTemp003 = πF.MakeArgs(1)
							πTemp010 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp010[0] = µobj
							if πE = πg.CheckLocal(πF, µformat_spec, "format_spec"); πE != nil {
								continue
							}
							πTemp010[1] = µformat_spec
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßformat_field, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp003[0] = πTemp005
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label9
						Label9:
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 611: return ''.join(result)
							πF.SetLineno(611)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp003[0] = µresult
							if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
					if πE = πClass.SetItem(πF, ß_vformat.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 614: def get_value(self, key, args, kwargs):
					πF.SetLineno(614)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					πTemp002[2] = πg.Param{Name: "args", Def: nil}
					πTemp002[3] = πg.Param{Name: "kwargs", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("get_value", "build/src/__python__/string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var µargs *πg.Object = πArgs[2]; _ = µargs
						var µkwargs *πg.Object = πArgs[3]; _ = µkwargs
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001[0] = µkey
							if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 615: if isinstance(key, (int, long)):
							πF.SetLineno(615)
						Label1:
							// line 616: return args[key]
							πF.SetLineno(616)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp002 = µkey
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µargs, πTemp002); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label3
						Label2:
							// line 618: return kwargs[key]
							πF.SetLineno(618)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp002 = µkey
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µkwargs, πTemp002); πE != nil {
								continue
							}
							πR = πTemp003
							continue
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
					if πE = πClass.SetItem(πF, ßget_value.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 621: def check_unused_args(self, used_args, args, kwargs):
					πF.SetLineno(621)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "used_args", Def: nil}
					πTemp002[2] = πg.Param{Name: "args", Def: nil}
					πTemp002[3] = πg.Param{Name: "kwargs", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("check_unused_args", "build/src/__python__/string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µused_args *πg.Object = πArgs[1]; _ = µused_args
						var µargs *πg.Object = πArgs[2]; _ = µargs
						var µkwargs *πg.Object = πArgs[3]; _ = µkwargs
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 622: pass
							πF.SetLineno(622)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcheck_unused_args.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 625: def format_field(self, value, format_spec):
					πF.SetLineno(625)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "value", Def: nil}
					πTemp002[2] = πg.Param{Name: "format_spec", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("format_field", "build/src/__python__/string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µvalue *πg.Object = πArgs[1]; _ = µvalue
						var µformat_spec *πg.Object = πArgs[2]; _ = µformat_spec
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
							// line 626: return format(value, format_spec)
							πF.SetLineno(626)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp001[0] = µvalue
							if πE = πg.CheckLocal(πF, µformat_spec, "format_spec"); πE != nil {
								continue
							}
							πTemp001[1] = µformat_spec
							if πTemp002, πE = πg.ResolveGlobal(πF, ßformat); πE != nil {
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
					if πE = πClass.SetItem(πF, ßformat_field.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 629: def convert_field(self, value, conversion):
					πF.SetLineno(629)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "value", Def: nil}
					πTemp002[2] = πg.Param{Name: "conversion", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("convert_field", "build/src/__python__/string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µvalue *πg.Object = πArgs[1]; _ = µvalue
						var µconversion *πg.Object = πArgs[2]; _ = µconversion
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
							if πE = πg.CheckLocal(πF, µconversion, "conversion"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µconversion == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µconversion, "conversion"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µconversion, ßs.ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label2
							}
							if πE = πg.CheckLocal(πF, µconversion, "conversion"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µconversion, ßr.ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 631: if conversion is None:
							πF.SetLineno(631)
						Label1:
							// line 632: return value
							πF.SetLineno(632)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πR = µvalue
							continue
							goto Label4
							// line 633: elif conversion == 's':
							πF.SetLineno(633)
						Label2:
							// line 634: return str(value)
							πF.SetLineno(634)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp004[0] = µvalue
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πR = πTemp002
							continue
							goto Label4
							// line 635: elif conversion == 'r':
							πF.SetLineno(635)
						Label3:
							// line 636: return repr(value)
							πF.SetLineno(636)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp004[0] = µvalue
							if πTemp001, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πR = πTemp002
							continue
							goto Label4
						Label4:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µconversion, "conversion"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Unknown conversion specifier %s").ToObject(), µconversion); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 637: raise ValueError("Unknown conversion specifier %s" % (conversion))
							πF.SetLineno(637)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßconvert_field.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 647: def parse(self, format_string):
					πF.SetLineno(647)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "format_string", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("parse", "build/src/__python__/string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µformat_string *πg.Object = πArgs[1]; _ = µformat_string
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
							// line 648: return format_string._formatter_parser()
							πF.SetLineno(648)
							if πE = πg.CheckLocal(πF, µformat_string, "format_string"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µformat_string, ß_formatter_parser, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßparse.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 656: def get_field(self, field_name, args, kwargs):
					πF.SetLineno(656)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "field_name", Def: nil}
					πTemp002[2] = πg.Param{Name: "args", Def: nil}
					πTemp002[3] = πg.Param{Name: "kwargs", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("get_field", "build/src/__python__/string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfield_name *πg.Object = πArgs[1]; _ = µfield_name
						var µargs *πg.Object = πArgs[2]; _ = µargs
						var µkwargs *πg.Object = πArgs[3]; _ = µkwargs
						var µfirst *πg.Object = πg.UnboundLocal; _ = µfirst
						var µrest *πg.Object = πg.UnboundLocal; _ = µrest
						var µobj *πg.Object = πg.UnboundLocal; _ = µobj
						var µis_attr *πg.Object = πg.UnboundLocal; _ = µis_attr
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							// line 657: first, rest = field_name._formatter_field_name_split()
							πF.SetLineno(657)
							if πE = πg.CheckLocal(πF, µfield_name, "field_name"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfield_name, ß_formatter_field_name_split, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
								continue
							}
							µfirst = πTemp001
							µrest = πTemp003
							// line 659: obj = self.get_value(first, args, kwargs)
							πF.SetLineno(659)
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							πTemp004[0] = µfirst
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp004[1] = µargs
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							πTemp004[2] = µkwargs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßget_value, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µobj = πTemp002
							if πE = πg.CheckLocal(πF, µrest, "rest"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µrest); πE != nil {
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
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp007}}}, πTemp002); πE != nil {
									continue
								}
								µis_attr = πTemp003
								µi = πTemp007
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µis_attr, "is_attr"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µis_attr); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 664: if is_attr:
							πF.SetLineno(664)
						Label4:
							// line 665: obj = getattr(obj, i)
							πF.SetLineno(665)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp004[0] = µobj
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004[1] = µi
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µobj = πTemp003
							goto Label6
						Label5:
							// line 667: obj = obj[i]
							πF.SetLineno(667)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp002 = µi
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µobj, πTemp002); πE != nil {
								continue
							}
							µobj = πTemp003
							goto Label6
						Label6:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 669: return obj, first
							πF.SetLineno(669)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µobj, µfirst).ToObject()
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
					if πE = πClass.SetItem(πF, ßget_field.ToObject(), πTemp010); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp034, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp034 == nil {
				πTemp034 = πg.TypeType.ToObject()
			}
			if πTemp035, πE = πTemp034.Call(πF, []*πg.Object{πg.NewStr("Formatter").ToObject(), πg.NewTuple(πTemp007...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFormatter.ToObject(), πTemp035); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("string", Code)
}
