package encoder
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/json/encoder.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßESCAPE := πg.InternStr("ESCAPE")
		ßESCAPE_ASCII := πg.InternStr("ESCAPE_ASCII")
		ßESCAPE_DCT := πg.InternStr("ESCAPE_DCT")
		ßFLOAT_REPR := πg.InternStr("FLOAT_REPR")
		ßFalse := πg.InternStr("False")
		ßHAS_UTF8 := πg.InternStr("HAS_UTF8")
		ßINFINITY := πg.InternStr("INFINITY")
		ßInfinity := πg.InternStr("Infinity")
		ßJSONEncoder := πg.InternStr("JSONEncoder")
		ßKeyError := πg.InternStr("KeyError")
		ßNaN := πg.InternStr("NaN")
		ßNone := πg.InternStr("None")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßValueError := πg.InternStr("ValueError")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_make_iterencode := πg.InternStr("_make_iterencode")
		ßallow_nan := πg.InternStr("allow_nan")
		ßbasestring := πg.InternStr("basestring")
		ßc_encode_basestring_ascii := πg.InternStr("c_encode_basestring_ascii")
		ßc_make_encoder := πg.InternStr("c_make_encoder")
		ßcheck_circular := πg.InternStr("check_circular")
		ßchr := πg.InternStr("chr")
		ßcompile := πg.InternStr("compile")
		ßdecode := πg.InternStr("decode")
		ßdefault := πg.InternStr("default")
		ßdict := πg.InternStr("dict")
		ßencode := πg.InternStr("encode")
		ßencode_basestring := πg.InternStr("encode_basestring")
		ßencode_basestring_ascii := πg.InternStr("encode_basestring_ascii")
		ßencoding := πg.InternStr("encoding")
		ßensure_ascii := πg.InternStr("ensure_ascii")
		ßfalse := πg.InternStr("false")
		ßfloat := πg.InternStr("float")
		ßgroup := πg.InternStr("group")
		ßi := πg.InternStr("i")
		ßid := πg.InternStr("id")
		ßindent := πg.InternStr("indent")
		ßinf := πg.InternStr("inf")
		ßint := πg.InternStr("int")
		ßisinstance := πg.InternStr("isinstance")
		ßitem_separator := πg.InternStr("item_separator")
		ßitems := πg.InternStr("items")
		ßiterencode := πg.InternStr("iterencode")
		ßiteritems := πg.InternStr("iteritems")
		ßjoin := πg.InternStr("join")
		ßkey_separator := πg.InternStr("key_separator")
		ßlist := πg.InternStr("list")
		ßlong := πg.InternStr("long")
		ßnull := πg.InternStr("null")
		ßobject := πg.InternStr("object")
		ßord := πg.InternStr("ord")
		ßpy_encode_basestring_ascii := πg.InternStr("py_encode_basestring_ascii")
		ßrange := πg.InternStr("range")
		ßre := πg.InternStr("re")
		ßrepr := πg.InternStr("repr")
		ßsearch := πg.InternStr("search")
		ßskipkeys := πg.InternStr("skipkeys")
		ßsort_keys := πg.InternStr("sort_keys")
		ßsorted := πg.InternStr("sorted")
		ßstr := πg.InternStr("str")
		ßsub := πg.InternStr("sub")
		ßtrue := πg.InternStr("true")
		ßtuple := πg.InternStr("tuple")
		ßx4 := πg.InternStr("x4")
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
		var πTemp006 *πg.Dict
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
		var πTemp012 *πg.Object
		_ = πTemp012
		var πTemp013 *πg.Object
		_ = πTemp013
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 1: goto Label1
			case 2: goto Label2
			default: panic("unexpected function state")
			}
			// line 1: """Implementation of JSONEncoder
			πF.SetLineno(1)
			// line 3: import re
			πF.SetLineno(3)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 9: c_encode_basestring_ascii = None
			πF.SetLineno(9)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßc_encode_basestring_ascii.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 15: c_make_encoder = None
			πF.SetLineno(15)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßc_make_encoder.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 17: def x4(i):
			πF.SetLineno(17)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "i", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("x4", "build/src/__python__/json/encoder.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µi *πg.Object = πArgs[0]; _ = µi
				var πTemp001 *πg.Object
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
					// line 18: return ("000%x" % i)[-4:]
					πF.SetLineno(18)
					if πTemp002, πE = πg.Neg(πF, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πTemp002, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("000%x").ToObject(), µi); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßx4.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 20: ESCAPE = re.compile(r'[\x00-\x1f\\"\b\f\n\r\t]')
			πF.SetLineno(20)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("[\\x00-\\x1f\\\\\"\\b\\f\\n\\r\\t]").ToObject()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßcompile, nil); πE != nil {
				continue
			}
			if πTemp004, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßESCAPE.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 21: ESCAPE_ASCII = re.compile(r'([\\"]|[^\ -~])')
			πF.SetLineno(21)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("([\\\\\"]|[^\\ -~])").ToObject()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßcompile, nil); πE != nil {
				continue
			}
			if πTemp004, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßESCAPE_ASCII.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 22: HAS_UTF8 = re.compile(r'[\x80-\xff]')
			πF.SetLineno(22)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("[\\x80-\\xff]").ToObject()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßcompile, nil); πE != nil {
				continue
			}
			if πTemp004, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßHAS_UTF8.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 23: ESCAPE_DCT = {
			πF.SetLineno(23)
			πTemp006 = πg.NewDict()
			if πE = πTemp006.SetItem(πF, πg.NewStr("\\").ToObject(), πg.NewStr("\\\\").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewStr("\"").ToObject(), πg.NewStr("\\\"").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewStr("\x08").ToObject(), πg.NewStr("\\b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewStr("\x0c").ToObject(), πg.NewStr("\\f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewStr("\n").ToObject(), πg.NewStr("\\n").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewStr("\r").ToObject(), πg.NewStr("\\r").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewStr("\t").ToObject(), πg.NewStr("\\t").ToObject()); πE != nil {
				continue
			}
			πTemp004 = πTemp006.ToObject()
			if πE = πF.Globals().SetItem(πF, ßESCAPE_DCT.ToObject(), πTemp004); πE != nil {
				continue
			}
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewInt(32).ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πTemp004, πE = πg.Iter(πF, πTemp007); πE != nil {
				continue
			}
			πF.PushCheckpoint(2)
			πTemp008 = false
		Label1:
			if πE != nil || πR != nil {
				continue
			}
			if πTemp008 {
				πF.PopCheckpoint()
				goto Label3
			}
			if πTemp005, πE = πg.Next(πF, πTemp004); πE != nil {
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
				if πE = πF.Globals().SetItem(πF, ßi.ToObject(), πTemp005); πE != nil {
					continue
				}
			}
			if πE != nil || !πTemp009 {
				continue
			}
			πF.PushCheckpoint(1)            
			// line 36: ESCAPE_DCT[chr(i)] = '\\u' + x4(i)
			πF.SetLineno(36)
			πTemp002 = πF.MakeArgs(1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßi); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßx4); πE != nil {
				continue
			}
			if πTemp010, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πTemp005, πE = πg.Add(πF, πg.NewStr("\\u").ToObject(), πTemp010); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πTemp005); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßESCAPE_DCT); πE != nil {
				continue
			}
			πTemp002 = πF.MakeArgs(1)
			if πTemp012, πE = πg.ResolveGlobal(πF, ßi); πE != nil {
				continue
			}
			πTemp002[0] = πTemp012
			if πTemp012, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp012.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			πTemp011 = πTemp013
			if πE = πg.SetItem(πF, πTemp010, πTemp011, πTemp007); πE != nil {
				continue
			}
			continue
		Label2:
			if πE != nil || πR != nil {
				continue
			}
		Label3:
			// line 38: INFINITY = float('inf')
			πF.SetLineno(38)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = ßinf.ToObject()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßINFINITY.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 39: FLOAT_REPR = repr
			πF.SetLineno(39)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFLOAT_REPR.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 41: def encode_basestring(s):
			πF.SetLineno(41)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "s", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("encode_basestring", "build/src/__python__/json/encoder.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µreplace *πg.Object = πg.UnboundLocal; _ = µreplace
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
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
					// line 42: """Return a JSON representation of a Python string
					πF.SetLineno(42)
					// line 45: def replace(match):
					πF.SetLineno(45)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "match", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("replace", "build/src/__python__/json/encoder.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µmatch *πg.Object = πArgs[0]; _ = µmatch
						var πTemp001 *πg.Object
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
							// line 46: return ESCAPE_DCT[match.group(0)]
							πF.SetLineno(46)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmatch, ßgroup, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßESCAPE_DCT); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
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
					µreplace = πTemp001
					// line 47: return '"' + ESCAPE.sub(replace, s) + '"'
					πF.SetLineno(47)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µreplace, "replace"); πE != nil {
						continue
					}
					πTemp005[0] = µreplace
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp005[1] = µs
					if πTemp006, πE = πg.ResolveGlobal(πF, ßESCAPE); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßsub, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp007.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp004, πE = πg.Add(πF, πg.NewStr("\"").ToObject(), πTemp006); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πTemp004, πg.NewStr("\"").ToObject()); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßencode_basestring.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 50: def py_encode_basestring_ascii(s):
			πF.SetLineno(50)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "s", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("py_encode_basestring_ascii", "build/src/__python__/json/encoder.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µreplace *πg.Object = πg.UnboundLocal; _ = µreplace
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
					// line 51: """Return an ASCII-only JSON representation of a Python string
					πF.SetLineno(51)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp003[0] = µs
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
					πTemp001 = πTemp005
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp003[0] = µs
					if πTemp005, πE = πg.ResolveGlobal(πF, ßHAS_UTF8); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßsearch, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(πTemp005 != πTemp006).ToObject()
					πTemp001 = πTemp004
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 54: if isinstance(s, str) and HAS_UTF8.search(s) is not None:
					πF.SetLineno(54)
				Label2:
					// line 55: s = s.decode('utf-8')
					πF.SetLineno(55)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("utf-8").ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µs, ßdecode, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µs = πTemp004
					goto Label3
				Label3:
					// line 56: def replace(match):
					πF.SetLineno(56)
					πTemp007 = make([]πg.Param, 1)
					πTemp007[0] = πg.Param{Name: "match", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("replace", "build/src/__python__/json/encoder.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µmatch *πg.Object = πArgs[0]; _ = µmatch
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var µs1 *πg.Object = πg.UnboundLocal; _ = µs1
						var µs2 *πg.Object = πg.UnboundLocal; _ = µs2
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
						var πTemp009 *πg.Object
						_ = πTemp009
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 57: s = match.group(0)
							πF.SetLineno(57)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µmatch, ßgroup, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp003
							// line 58: try:
							πF.SetLineno(58)
							πF.PushCheckpoint(2)
							// line 59: return ESCAPE_DCT[s]
							πF.SetLineno(59)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp002 = µs
							if πTemp004, πE = πg.ResolveGlobal(πF, ßESCAPE_DCT); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label3
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 60: except KeyError:
							πF.SetLineno(60)
						Label3:
							// line 61: n = ord(s)
							πF.SetLineno(61)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πTemp002, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µn = πTemp003
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, µn, πg.NewInt(65536).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label4
							}
							goto Label5
							// line 62: if n < 0x10000:
							πF.SetLineno(62)
						Label4:
							// line 65: return '\\u' + x4(n)
							πF.SetLineno(65)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[0] = µn
							if πTemp003, πE = πg.ResolveGlobal(πF, ßx4); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Add(πF, πg.NewStr("\\u").ToObject(), πTemp004); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label6
						Label5:
							// line 68: n -= 0x10000
							πF.SetLineno(68)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ISub(πF, µn, πg.NewInt(65536).ToObject()); πE != nil {
								continue
							}
							µn = πTemp002
							// line 69: s1 = 0xd800 | ((n >> 10) & 0x3ff)
							πF.SetLineno(69)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.RShift(πF, µn, πg.NewInt(10).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.And(πF, πTemp004, πg.NewInt(1023).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Or(πF, πg.NewInt(55296).ToObject(), πTemp003); πE != nil {
								continue
							}
							µs1 = πTemp002
							// line 70: s2 = 0xdc00 | (n & 0x3ff)
							πF.SetLineno(70)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.And(πF, µn, πg.NewInt(1023).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Or(πF, πg.NewInt(56320).ToObject(), πTemp003); πE != nil {
								continue
							}
							µs2 = πTemp002
							// line 73: return '\\u' + x4(s1) + '\\u' + x4(s2)
							πF.SetLineno(73)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							πTemp001[0] = µs1
							if πTemp008, πE = πg.ResolveGlobal(πF, ßx4); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.Add(πF, πg.NewStr("\\u").ToObject(), πTemp009); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, πg.NewStr("\\u").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							πTemp001[0] = µs2
							if πTemp004, πE = πg.ResolveGlobal(πF, ßx4); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp008); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label6
						Label6:
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
					µreplace = πTemp001
					// line 74: return '"' + str(ESCAPE_ASCII.sub(replace, s)) + '"'
					πF.SetLineno(74)
					πTemp003 = πF.MakeArgs(1)
					πTemp008 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µreplace, "replace"); πE != nil {
						continue
					}
					πTemp008[0] = µreplace
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp008[1] = µs
					if πTemp006, πE = πg.ResolveGlobal(πF, ßESCAPE_ASCII); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp006, ßsub, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp009.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					πTemp003[0] = πTemp006
					if πTemp006, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp005, πE = πg.Add(πF, πg.NewStr("\"").ToObject(), πTemp009); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, πTemp005, πg.NewStr("\"").ToObject()); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßpy_encode_basestring_ascii.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 77: encode_basestring_ascii = (
			πF.SetLineno(77)
			if πTemp010, πE = πg.ResolveGlobal(πF, ßc_encode_basestring_ascii); πE != nil {
				continue
			}
			πTemp007 = πTemp010
			if πTemp008, πE = πg.IsTrue(πF, πTemp007); πE != nil {
				continue
			}
			if πTemp008 {
				goto Label4
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßpy_encode_basestring_ascii); πE != nil {
				continue
			}
			πTemp007 = πTemp010
		Label4:
			if πE = πF.Globals().SetItem(πF, ßencode_basestring_ascii.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 80: class JSONEncoder(object):
			πF.SetLineno(80)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp011, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp011
			πTemp006 = πg.NewDict()
			if πTemp007, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp007); πE != nil {
				continue
			}
			_, πE = πg.NewCode("JSONEncoder", "build/src/__python__/json/encoder.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 81: """Extensible JSON <http://json.org> encoder for Python data structures.
					πF.SetLineno(81)
					// line 109: item_separator = ', '
					πF.SetLineno(109)
					if πE = πClass.SetItem(πF, ßitem_separator.ToObject(), πg.NewStr(", ").ToObject()); πE != nil {
						continue
					}
					// line 110: key_separator = ': '
					πF.SetLineno(110)
					if πE = πClass.SetItem(πF, ßkey_separator.ToObject(), πg.NewStr(": ").ToObject()); πE != nil {
						continue
					}
					// line 111: def __init__(self, skipkeys=False, ensure_ascii=True,
					πF.SetLineno(111)
					πTemp002 = make([]πg.Param, 10)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "skipkeys", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "ensure_ascii", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "check_circular", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "allow_nan", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "sort_keys", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[6] = πg.Param{Name: "indent", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[7] = πg.Param{Name: "separators", Def: πTemp003}
					πTemp002[8] = πg.Param{Name: "encoding", Def: πg.NewStr("utf-8").ToObject()}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[9] = πg.Param{Name: "default", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/json/encoder.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µskipkeys *πg.Object = πArgs[1]; _ = µskipkeys
						var µensure_ascii *πg.Object = πArgs[2]; _ = µensure_ascii
						var µcheck_circular *πg.Object = πArgs[3]; _ = µcheck_circular
						var µallow_nan *πg.Object = πArgs[4]; _ = µallow_nan
						var µsort_keys *πg.Object = πArgs[5]; _ = µsort_keys
						var µindent *πg.Object = πArgs[6]; _ = µindent
						var µseparators *πg.Object = πArgs[7]; _ = µseparators
						var µencoding *πg.Object = πArgs[8]; _ = µencoding
						var µdefault *πg.Object = πArgs[9]; _ = µdefault
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 114: """Constructor for JSONEncoder, with sensible defaults.
							πF.SetLineno(114)
							// line 163: self.skipkeys = skipkeys
							πF.SetLineno(163)
							if πE = πg.CheckLocal(πF, µskipkeys, "skipkeys"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µskipkeys); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßskipkeys, πTemp001); πE != nil {
								continue
							}
							// line 164: self.ensure_ascii = ensure_ascii
							πF.SetLineno(164)
							if πE = πg.CheckLocal(πF, µensure_ascii, "ensure_ascii"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µensure_ascii); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßensure_ascii, πTemp001); πE != nil {
								continue
							}
							// line 165: self.check_circular = check_circular
							πF.SetLineno(165)
							if πE = πg.CheckLocal(πF, µcheck_circular, "check_circular"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µcheck_circular); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcheck_circular, πTemp001); πE != nil {
								continue
							}
							// line 166: self.allow_nan = allow_nan
							πF.SetLineno(166)
							if πE = πg.CheckLocal(πF, µallow_nan, "allow_nan"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µallow_nan); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßallow_nan, πTemp001); πE != nil {
								continue
							}
							// line 167: self.sort_keys = sort_keys
							πF.SetLineno(167)
							if πE = πg.CheckLocal(πF, µsort_keys, "sort_keys"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µsort_keys); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsort_keys, πTemp001); πE != nil {
								continue
							}
							// line 168: self.indent = indent
							πF.SetLineno(168)
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µindent); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßindent, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µseparators, "separators"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µseparators != πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 169: if separators is not None:
							πF.SetLineno(169)
						Label1:
							// line 170: self.item_separator, self.key_separator = separators
							πF.SetLineno(170)
							if πE = πg.CheckLocal(πF, µseparators, "separators"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp002}}}, µseparators); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßitem_separator, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßkey_separator, πTemp002); πE != nil {
								continue
							}
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µdefault != πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 171: if default is not None:
							πF.SetLineno(171)
						Label3:
							// line 172: self.default = default
							πF.SetLineno(172)
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdefault); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdefault, πTemp001); πE != nil {
								continue
							}
							goto Label4
						Label4:
							// line 173: self.encoding = encoding
							πF.SetLineno(173)
							if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µencoding); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßencoding, πTemp001); πE != nil {
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
					// line 175: def default(self, o):
					πF.SetLineno(175)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "o", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("default", "build/src/__python__/json/encoder.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µo *πg.Object = πArgs[1]; _ = µo
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 176: """Implement this method in a subclass such that it returns
							πF.SetLineno(176)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp003[0] = µo
							if πTemp004, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.Add(πF, πTemp005, πg.NewStr(" is not JSON serializable").ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 194: raise TypeError(repr(o) + " is not JSON serializable")
							πF.SetLineno(194)
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
					if πE = πClass.SetItem(πF, ßdefault.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 196: def encode(self, o):
					πF.SetLineno(196)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "o", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("encode", "build/src/__python__/json/encoder.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µo *πg.Object = πArgs[1]; _ = µo
						var µ_encoding *πg.Object = πg.UnboundLocal; _ = µ_encoding
						var µchunks *πg.Object = πg.UnboundLocal; _ = µchunks
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 πg.KWArgs
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
							// line 197: """Return a JSON string representation of a Python data structure.
							πF.SetLineno(197)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp001[0] = µo
							if πTemp002, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 204: if isinstance(o, basestring):
							πF.SetLineno(204)
						Label1:
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp001[0] = µo
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 205: if isinstance(o, str):
							πF.SetLineno(205)
						Label3:
							// line 206: _encoding = self.encoding
							πF.SetLineno(206)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							µ_encoding = πTemp002
							if πE = πg.CheckLocal(πF, µ_encoding, "_encoding"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µ_encoding != πTemp005).ToObject()
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label5
							}
							if πE = πg.CheckLocal(πF, µ_encoding, "_encoding"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Eq(πF, µ_encoding, πg.NewStr("utf-8").ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp006).ToObject()
							πTemp002 = πTemp003
						Label5:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							goto Label7
							// line 207: if (_encoding is not None
							πF.SetLineno(207)
						Label6:
							// line 209: o = o.decode(_encoding)
							πF.SetLineno(209)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µ_encoding, "_encoding"); πE != nil {
								continue
							}
							πTemp001[0] = µ_encoding
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µo, ßdecode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µo = πTemp003
							goto Label7
						Label7:
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßensure_ascii, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							goto Label9
							// line 210: if self.ensure_ascii:
							πF.SetLineno(210)
						Label8:
							// line 211: return encode_basestring_ascii(o)
							πF.SetLineno(211)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp001[0] = µo
							if πTemp002, πE = πg.ResolveGlobal(πF, ßencode_basestring_ascii); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp003
							continue
							goto Label10
						Label9:
							// line 213: return encode_basestring(o)
							πF.SetLineno(213)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp001[0] = µo
							if πTemp002, πE = πg.ResolveGlobal(πF, ßencode_basestring); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp003
							continue
							goto Label10
						Label10:
							goto Label2
						Label2:
							// line 217: chunks = self.iterencode(o, _one_shot=True)
							πF.SetLineno(217)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp001[0] = µo
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp007 = πg.KWArgs{
								{"_one_shot", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßiterencode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp007); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µchunks = πTemp003
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
								continue
							}
							πTemp001[0] = µchunks
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp008, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp005, πTemp008).ToObject()
							πTemp001[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label11
							}
							goto Label12
							// line 218: if not isinstance(chunks, (list, tuple)):
							πF.SetLineno(218)
						Label11:
							// line 219: chunks = list(chunks)
							πF.SetLineno(219)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
								continue
							}
							πTemp001[0] = µchunks
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µchunks = πTemp003
							goto Label12
						Label12:
							// line 220: return ''.join(chunks)
							πF.SetLineno(220)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
								continue
							}
							πTemp001[0] = µchunks
							if πTemp002, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßencode.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 222: def iterencode(self, o, _one_shot=False):
					πF.SetLineno(222)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "o", Def: nil}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "_one_shot", Def: πTemp006}
					πTemp005 = πg.NewFunction(πg.NewCode("iterencode", "build/src/__python__/json/encoder.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µo *πg.Object = πArgs[1]; _ = µo
						var µ_one_shot *πg.Object = πArgs[2]; _ = µ_one_shot
						var µmarkers *πg.Object = πg.UnboundLocal; _ = µmarkers
						var µ_encoder *πg.Object = πg.UnboundLocal; _ = µ_encoder
						var µfloatstr *πg.Object = πg.UnboundLocal; _ = µfloatstr
						var µ_iterencode *πg.Object = πg.UnboundLocal; _ = µ_iterencode
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Dict
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []πg.Param
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 bool
						_ = πTemp010
						var πTemp011 []*πg.Object
						_ = πTemp011
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 223: """Encode the given object and yield each string
							πF.SetLineno(223)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcheck_circular, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 232: if self.check_circular:
							πF.SetLineno(232)
						Label1:
							// line 233: markers = {}
							πF.SetLineno(233)
							πTemp003 = πg.NewDict()
							πTemp001 = πTemp003.ToObject()
							µmarkers = πTemp001
							goto Label3
						Label2:
							// line 235: markers = None
							πF.SetLineno(235)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µmarkers = πTemp001
							goto Label3
						Label3:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßensure_ascii, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							goto Label5
							// line 236: if self.ensure_ascii:
							πF.SetLineno(236)
						Label4:
							// line 237: _encoder = encode_basestring_ascii
							πF.SetLineno(237)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßencode_basestring_ascii); πE != nil {
								continue
							}
							µ_encoder = πTemp001
							goto Label6
						Label5:
							// line 239: _encoder = encode_basestring
							πF.SetLineno(239)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßencode_basestring); πE != nil {
								continue
							}
							µ_encoder = πTemp001
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.NE(πF, πTemp004, πg.NewStr("utf-8").ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label7
							}
							goto Label8
							// line 240: if self.encoding != 'utf-8':
							πF.SetLineno(240)
						Label7:
							// line 241: def _encoder(o, _orig_encoder=_encoder, _encoding=self.encoding):
							πF.SetLineno(241)
							πTemp005 = make([]πg.Param, 3)
							πTemp005[0] = πg.Param{Name: "o", Def: nil}
							if πE = πg.CheckLocal(πF, µ_encoder, "_encoder"); πE != nil {
								continue
							}
							πTemp005[1] = πg.Param{Name: "_orig_encoder", Def: µ_encoder}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							πTemp005[2] = πg.Param{Name: "_encoding", Def: πTemp004}
							πTemp001 = πg.NewFunction(πg.NewCode("_encoder", "build/src/__python__/json/encoder.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µo *πg.Object = πArgs[0]; _ = µo
								var µ_orig_encoder *πg.Object = πArgs[1]; _ = µ_orig_encoder
								var µ_encoding *πg.Object = πArgs[2]; _ = µ_encoding
								var πTemp001 []*πg.Object
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
									πTemp001 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
										continue
									}
									πTemp001[0] = µo
									if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
										continue
									}
									πTemp001[1] = πTemp002
									if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
										continue
									}
									if πTemp004 {
										goto Label1
									}
									goto Label2
									// line 242: if isinstance(o, str):
									πF.SetLineno(242)
								Label1:
									// line 243: o = o.decode(_encoding)
									πF.SetLineno(243)
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µ_encoding, "_encoding"); πE != nil {
										continue
									}
									πTemp001[0] = µ_encoding
									if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µo, ßdecode, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									µo = πTemp003
									goto Label2
								Label2:
									// line 244: return _orig_encoder(o)
									πF.SetLineno(244)
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
										continue
									}
									πTemp001[0] = µo
									if πE = πg.CheckLocal(πF, µ_orig_encoder, "_orig_encoder"); πE != nil {
										continue
									}
									if πTemp002, πE = µ_orig_encoder.Call(πF, πTemp001, nil); πE != nil {
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
							µ_encoder = πTemp001
							goto Label8
						Label8:
							// line 246: def floatstr(o, allow_nan=self.allow_nan,
							πF.SetLineno(246)
							πTemp005 = make([]πg.Param, 5)
							πTemp005[0] = πg.Param{Name: "o", Def: nil}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßallow_nan, nil); πE != nil {
								continue
							}
							πTemp005[1] = πg.Param{Name: "allow_nan", Def: πTemp006}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßFLOAT_REPR); πE != nil {
								continue
							}
							πTemp005[2] = πg.Param{Name: "_repr", Def: πTemp006}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßINFINITY); πE != nil {
								continue
							}
							πTemp005[3] = πg.Param{Name: "_inf", Def: πTemp006}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßINFINITY); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πTemp007); πE != nil {
								continue
							}
							πTemp005[4] = πg.Param{Name: "_neginf", Def: πTemp006}
							πTemp004 = πg.NewFunction(πg.NewCode("floatstr", "build/src/__python__/json/encoder.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µo *πg.Object = πArgs[0]; _ = µo
								var µallow_nan *πg.Object = πArgs[1]; _ = µallow_nan
								var µ_repr *πg.Object = πArgs[2]; _ = µ_repr
								var µ_inf *πg.Object = πArgs[3]; _ = µ_inf
								var µ_neginf *πg.Object = πArgs[4]; _ = µ_neginf
								var µtext *πg.Object = πg.UnboundLocal; _ = µtext
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 bool
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
									if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.NE(πF, µo, µo); πE != nil {
										continue
									}
									if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp002 {
										goto Label1
									}
									if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µ_inf, "_inf"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.Eq(πF, µo, µ_inf); πE != nil {
										continue
									}
									if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp002 {
										goto Label2
									}
									if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µ_neginf, "_neginf"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.Eq(πF, µo, µ_neginf); πE != nil {
										continue
									}
									if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp002 {
										goto Label3
									}
									goto Label4
									// line 252: if o != o:
									πF.SetLineno(252)
								Label1:
									// line 253: text = 'NaN'
									πF.SetLineno(253)
									µtext = ßNaN.ToObject()
									goto Label5
									// line 254: elif o == _inf:
									πF.SetLineno(254)
								Label2:
									// line 255: text = 'Infinity'
									πF.SetLineno(255)
									µtext = ßInfinity.ToObject()
									goto Label5
									// line 256: elif o == _neginf:
									πF.SetLineno(256)
								Label3:
									// line 257: text = '-Infinity'
									πF.SetLineno(257)
									µtext = πg.NewStr("-Infinity").ToObject()
									goto Label5
								Label4:
									// line 259: return _repr(o)
									πF.SetLineno(259)
									πTemp003 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
										continue
									}
									πTemp003[0] = µo
									if πE = πg.CheckLocal(πF, µ_repr, "_repr"); πE != nil {
										continue
									}
									if πTemp001, πE = µ_repr.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									πR = πTemp001
									continue
									goto Label5
								Label5:
									if πE = πg.CheckLocal(πF, µallow_nan, "allow_nan"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.IsTrue(πF, µallow_nan); πE != nil {
										continue
									}
									πTemp001 = πg.GetBool(!πTemp002).ToObject()
									if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp002 {
										goto Label6
									}
									goto Label7
									// line 261: if not allow_nan:
									πF.SetLineno(261)
								Label6:
									πTemp003 = πF.MakeArgs(1)
									πTemp004 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
										continue
									}
									πTemp004[0] = µo
									if πTemp005, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
										continue
									}
									if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp004)
									if πTemp001, πE = πg.Add(πF, πg.NewStr("Out of range float values are not JSON compliant: ").ToObject(), πTemp006); πE != nil {
										continue
									}
									πTemp003[0] = πTemp001
									if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
										continue
									}
									if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									// line 262: raise ValueError(
									πF.SetLineno(262)
									πE = πF.Raise(πTemp005, nil, nil)
									continue
									goto Label7
								Label7:
									// line 266: return text
									πF.SetLineno(266)
									if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
										continue
									}
									πR = µtext
									continue
								}
								if πE != nil {
									πR = nil
								} else if πR == nil {
									πR = πg.None
								}
								return πR, πE
							}), πF.Globals()).ToObject()
							µfloatstr = πTemp004
							if πE = πg.CheckLocal(πF, µ_one_shot, "_one_shot"); πE != nil {
								continue
							}
							πTemp006 = µ_one_shot
							if πTemp002, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label9
							}
							if πTemp008, πE = πg.ResolveGlobal(πF, ßc_make_encoder); πE != nil {
								continue
							}
							if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp007 = πg.GetBool(πTemp008 != πTemp009).ToObject()
							πTemp006 = πTemp007
							if πTemp002, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label9
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßindent, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp007 = πg.GetBool(πTemp008 == πTemp009).ToObject()
							πTemp006 = πTemp007
							if πTemp002, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label9
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßsort_keys, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							πTemp007 = πg.GetBool(!πTemp010).ToObject()
							πTemp006 = πTemp007
						Label9:
							if πTemp002, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label10
							}
							goto Label11
							// line 269: if (_one_shot and c_make_encoder is not None
							πF.SetLineno(269)
						Label10:
							// line 271: _iterencode = c_make_encoder(
							πF.SetLineno(271)
							πTemp011 = πF.MakeArgs(9)
							if πE = πg.CheckLocal(πF, µmarkers, "markers"); πE != nil {
								continue
							}
							πTemp011[0] = µmarkers
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefault, nil); πE != nil {
								continue
							}
							πTemp011[1] = πTemp006
							if πE = πg.CheckLocal(πF, µ_encoder, "_encoder"); πE != nil {
								continue
							}
							πTemp011[2] = µ_encoder
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßindent, nil); πE != nil {
								continue
							}
							πTemp011[3] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßkey_separator, nil); πE != nil {
								continue
							}
							πTemp011[4] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßitem_separator, nil); πE != nil {
								continue
							}
							πTemp011[5] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßsort_keys, nil); πE != nil {
								continue
							}
							πTemp011[6] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßskipkeys, nil); πE != nil {
								continue
							}
							πTemp011[7] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßallow_nan, nil); πE != nil {
								continue
							}
							πTemp011[8] = πTemp006
							if πTemp006, πE = πg.ResolveGlobal(πF, ßc_make_encoder); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp011, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp011)
							µ_iterencode = πTemp007
							goto Label12
						Label11:
							// line 276: _iterencode = _make_iterencode(
							πF.SetLineno(276)
							πTemp011 = πF.MakeArgs(10)
							if πE = πg.CheckLocal(πF, µmarkers, "markers"); πE != nil {
								continue
							}
							πTemp011[0] = µmarkers
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefault, nil); πE != nil {
								continue
							}
							πTemp011[1] = πTemp006
							if πE = πg.CheckLocal(πF, µ_encoder, "_encoder"); πE != nil {
								continue
							}
							πTemp011[2] = µ_encoder
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßindent, nil); πE != nil {
								continue
							}
							πTemp011[3] = πTemp006
							if πE = πg.CheckLocal(πF, µfloatstr, "floatstr"); πE != nil {
								continue
							}
							πTemp011[4] = µfloatstr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßkey_separator, nil); πE != nil {
								continue
							}
							πTemp011[5] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßitem_separator, nil); πE != nil {
								continue
							}
							πTemp011[6] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßsort_keys, nil); πE != nil {
								continue
							}
							πTemp011[7] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßskipkeys, nil); πE != nil {
								continue
							}
							πTemp011[8] = πTemp006
							if πE = πg.CheckLocal(πF, µ_one_shot, "_one_shot"); πE != nil {
								continue
							}
							πTemp011[9] = µ_one_shot
							if πTemp006, πE = πg.ResolveGlobal(πF, ß_make_iterencode); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp011, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp011)
							µ_iterencode = πTemp007
							goto Label12
						Label12:
							// line 280: return _iterencode(o, 0)
							πF.SetLineno(280)
							πTemp011 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp011[0] = µo
							πTemp011[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µ_iterencode, "_iterencode"); πE != nil {
								continue
							}
							if πTemp006, πE = µ_iterencode.Call(πF, πTemp011, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp011)
							πR = πTemp006
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßiterencode.ToObject(), πTemp005); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp010, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp010 == nil {
				πTemp010 = πg.TypeType.ToObject()
			}
			if πTemp011, πE = πTemp010.Call(πF, []*πg.Object{πg.NewStr("JSONEncoder").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßJSONEncoder.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 282: def _make_iterencode(markers, _default, _encoder, _indent, _floatstr,
			πF.SetLineno(282)
			πTemp003 = make([]πg.Param, 21)
			πTemp003[0] = πg.Param{Name: "markers", Def: nil}
			πTemp003[1] = πg.Param{Name: "_default", Def: nil}
			πTemp003[2] = πg.Param{Name: "_encoder", Def: nil}
			πTemp003[3] = πg.Param{Name: "_indent", Def: nil}
			πTemp003[4] = πg.Param{Name: "_floatstr", Def: nil}
			πTemp003[5] = πg.Param{Name: "_key_separator", Def: nil}
			πTemp003[6] = πg.Param{Name: "_item_separator", Def: nil}
			πTemp003[7] = πg.Param{Name: "_sort_keys", Def: nil}
			πTemp003[8] = πg.Param{Name: "_skipkeys", Def: nil}
			πTemp003[9] = πg.Param{Name: "_one_shot", Def: nil}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
				continue
			}
			πTemp003[10] = πg.Param{Name: "ValueError", Def: πTemp010}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
				continue
			}
			πTemp003[11] = πg.Param{Name: "basestring", Def: πTemp010}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
				continue
			}
			πTemp003[12] = πg.Param{Name: "dict", Def: πTemp010}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
				continue
			}
			πTemp003[13] = πg.Param{Name: "float", Def: πTemp010}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
				continue
			}
			πTemp003[14] = πg.Param{Name: "id", Def: πTemp010}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
				continue
			}
			πTemp003[15] = πg.Param{Name: "int", Def: πTemp010}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
				continue
			}
			πTemp003[16] = πg.Param{Name: "isinstance", Def: πTemp010}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
				continue
			}
			πTemp003[17] = πg.Param{Name: "list", Def: πTemp010}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
				continue
			}
			πTemp003[18] = πg.Param{Name: "long", Def: πTemp010}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
				continue
			}
			πTemp003[19] = πg.Param{Name: "str", Def: πTemp010}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
				continue
			}
			πTemp003[20] = πg.Param{Name: "tuple", Def: πTemp010}
			πTemp007 = πg.NewFunction(πg.NewCode("_make_iterencode", "build/src/__python__/json/encoder.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmarkers *πg.Object = πArgs[0]; _ = µmarkers
				var µ_default *πg.Object = πArgs[1]; _ = µ_default
				var µ_encoder *πg.Object = πArgs[2]; _ = µ_encoder
				var µ_indent *πg.Object = πArgs[3]; _ = µ_indent
				var µ_floatstr *πg.Object = πArgs[4]; _ = µ_floatstr
				var µ_key_separator *πg.Object = πArgs[5]; _ = µ_key_separator
				var µ_item_separator *πg.Object = πArgs[6]; _ = µ_item_separator
				var µ_sort_keys *πg.Object = πArgs[7]; _ = µ_sort_keys
				var µ_skipkeys *πg.Object = πArgs[8]; _ = µ_skipkeys
				var µ_one_shot *πg.Object = πArgs[9]; _ = µ_one_shot
				var µValueError *πg.Object = πArgs[10]; _ = µValueError
				var µbasestring *πg.Object = πArgs[11]; _ = µbasestring
				var µdict *πg.Object = πArgs[12]; _ = µdict
				var µfloat *πg.Object = πArgs[13]; _ = µfloat
				var µid *πg.Object = πArgs[14]; _ = µid
				var µint *πg.Object = πArgs[15]; _ = µint
				var µisinstance *πg.Object = πArgs[16]; _ = µisinstance
				var µlist *πg.Object = πArgs[17]; _ = µlist
				var µlong *πg.Object = πArgs[18]; _ = µlong
				var µstr *πg.Object = πArgs[19]; _ = µstr
				var µtuple *πg.Object = πArgs[20]; _ = µtuple
				var µ_iterencode_list *πg.Object = πg.UnboundLocal; _ = µ_iterencode_list
				var µ_iterencode_dict *πg.Object = πg.UnboundLocal; _ = µ_iterencode_dict
				var µ_iterencode *πg.Object = πg.UnboundLocal; _ = µ_iterencode
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
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
					// line 298: def _iterencode_list(lst, _current_indent_level):
					πF.SetLineno(298)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "lst", Def: nil}
					πTemp002[1] = πg.Param{Name: "_current_indent_level", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("_iterencode_list", "build/src/__python__/json/encoder.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µlst *πg.Object = πArgs[0]; _ = µlst
						var µ_current_indent_level *πg.Object = πArgs[1]; _ = µ_current_indent_level
						var µmarkerid *πg.Object = πg.UnboundLocal; _ = µmarkerid
						var µbuf *πg.Object = πg.UnboundLocal; _ = µbuf
						var µnewline_indent *πg.Object = πg.UnboundLocal; _ = µnewline_indent
						var µseparator *πg.Object = πg.UnboundLocal; _ = µseparator
						var µfirst *πg.Object = πg.UnboundLocal; _ = µfirst
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
						var µchunks *πg.Object = πg.UnboundLocal; _ = µchunks
						var µchunk *πg.Object = πg.UnboundLocal; _ = µchunk
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
						var πTemp006 bool
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
						var πTemp012 bool
						_ = πTemp012
						var πTemp013 *πg.Object
						_ = πTemp013
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 3: goto Label3
								case 36: goto Label36
								case 37: goto Label37
								case 39: goto Label39
								case 42: goto Label42
								case 11: goto Label11
								case 12: goto Label12
								case 43: goto Label43
								case 25: goto Label25
								case 26: goto Label26
								case 27: goto Label27
								case 28: goto Label28
								case 29: goto Label29
								case 30: goto Label30
								case 31: goto Label31
								default: panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.IsTrue(πF, µlst); πE != nil {
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
								// line 299: if not lst:
								πF.SetLineno(299)
							Label1:
								// line 300: yield '[]'
								πF.SetLineno(300)
								πF.PushCheckpoint(3)
								return πg.NewStr("[]").ToObject(), nil
							Label3:
								πTemp001 = πSent
								// line 301: return
								πF.SetLineno(301)
								πR = πg.None
								continue
								goto Label2
							Label2:
								if πE = πg.CheckLocal(πF, µmarkers, "markers"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp001 = πg.GetBool(µmarkers != πTemp003).ToObject()
								if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp002 {
									goto Label4
								}
								goto Label5
								// line 302: if markers is not None:
								πF.SetLineno(302)
							Label4:
								// line 303: markerid = id(lst)
								πF.SetLineno(303)
								πTemp004 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
									continue
								}
								πTemp004[0] = µlst
								if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
									continue
								}
								if πTemp001, πE = µid.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								µmarkerid = πTemp001
								if πE = πg.CheckLocal(πF, µmarkerid, "markerid"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmarkers, "markers"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Contains(πF, µmarkers, µmarkerid); πE != nil {
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
								// line 304: if markerid in markers:
								πF.SetLineno(304)
							Label6:
								πTemp004 = πF.MakeArgs(1)
								πTemp004[0] = πg.NewStr("Circular reference detected").ToObject()
								if πE = πg.CheckLocal(πF, µValueError, "ValueError"); πE != nil {
									continue
								}
								if πTemp001, πE = µValueError.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								// line 305: raise ValueError("Circular reference detected")
								πF.SetLineno(305)
								πE = πF.Raise(πTemp001, nil, nil)
								continue
								goto Label7
							Label7:
								// line 306: markers[markerid] = lst
								πF.SetLineno(306)
								if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µlst); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmarkers, "markers"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmarkerid, "markerid"); πE != nil {
									continue
								}
								πTemp003 = µmarkerid
								if πE = πg.SetItem(πF, µmarkers, πTemp003, πTemp001); πE != nil {
									continue
								}
								goto Label5
							Label5:
								// line 307: buf = '['
								πF.SetLineno(307)
								µbuf = πg.NewStr("[").ToObject()
								if πE = πg.CheckLocal(πF, µ_indent, "_indent"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp001 = πg.GetBool(µ_indent != πTemp003).ToObject()
								if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp002 {
									goto Label8
								}
								goto Label9
								// line 308: if _indent is not None:
								πF.SetLineno(308)
							Label8:
								// line 309: _current_indent_level += 1
								πF.SetLineno(309)
								if πE = πg.CheckLocal(πF, µ_current_indent_level, "_current_indent_level"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.IAdd(πF, µ_current_indent_level, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								µ_current_indent_level = πTemp001
								// line 310: newline_indent = '\n' + (' ' * (_indent * _current_indent_level))
								πF.SetLineno(310)
								if πE = πg.CheckLocal(πF, µ_indent, "_indent"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µ_current_indent_level, "_current_indent_level"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.Mul(πF, µ_indent, µ_current_indent_level); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Mul(πF, πg.NewStr(" ").ToObject(), πTemp005); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Add(πF, πg.NewStr("\n").ToObject(), πTemp003); πE != nil {
									continue
								}
								µnewline_indent = πTemp001
								// line 311: separator = _item_separator + newline_indent
								πF.SetLineno(311)
								if πE = πg.CheckLocal(πF, µ_item_separator, "_item_separator"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µnewline_indent, "newline_indent"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Add(πF, µ_item_separator, µnewline_indent); πE != nil {
									continue
								}
								µseparator = πTemp001
								// line 312: buf += newline_indent
								πF.SetLineno(312)
								if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µnewline_indent, "newline_indent"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.IAdd(πF, µbuf, µnewline_indent); πE != nil {
									continue
								}
								µbuf = πTemp001
								goto Label10
							Label9:
								// line 314: newline_indent = None
								πF.SetLineno(314)
								if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								µnewline_indent = πTemp001
								// line 315: separator = _item_separator
								πF.SetLineno(315)
								if πE = πg.CheckLocal(πF, µ_item_separator, "_item_separator"); πE != nil {
									continue
								}
								µseparator = µ_item_separator
								goto Label10
							Label10:
								// line 316: first = True
								πF.SetLineno(316)
								if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								µfirst = πTemp001
								if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µlst); πE != nil {
									continue
								}
								πF.PushCheckpoint(12)
								πTemp002 = false
							Label11:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp002 {
									πF.PopCheckpoint()
									goto Label13
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
									µvalue = πTemp003
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(11)            
								if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.IsTrue(πF, µfirst); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label14
								}
								goto Label15
								// line 318: if first:
								πF.SetLineno(318)
							Label14:
								// line 319: first = False
								πF.SetLineno(319)
								if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								µfirst = πTemp003
								goto Label16
							Label15:
								// line 321: buf = separator
								πF.SetLineno(321)
								if πE = πg.CheckLocal(πF, µseparator, "separator"); πE != nil {
									continue
								}
								µbuf = µseparator
								goto Label16
							Label16:
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp004[0] = µvalue
								if πE = πg.CheckLocal(πF, µbasestring, "basestring"); πE != nil {
									continue
								}
								πTemp004[1] = µbasestring
								if πE = πg.CheckLocal(πF, µisinstance, "isinstance"); πE != nil {
									continue
								}
								if πTemp003, πE = µisinstance.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label17
								}
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp003 = πg.GetBool(µvalue == πTemp005).ToObject()
								if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label18
								}
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πTemp003 = πg.GetBool(µvalue == πTemp005).ToObject()
								if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label19
								}
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								πTemp003 = πg.GetBool(µvalue == πTemp005).ToObject()
								if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label20
								}
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp004[0] = µvalue
								if πE = πg.CheckLocal(πF, µint, "int"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µlong, "long"); πE != nil {
									continue
								}
								πTemp003 = πg.NewTuple2(µint, µlong).ToObject()
								πTemp004[1] = πTemp003
								if πE = πg.CheckLocal(πF, µisinstance, "isinstance"); πE != nil {
									continue
								}
								if πTemp003, πE = µisinstance.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label21
								}
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp004[0] = µvalue
								if πE = πg.CheckLocal(πF, µfloat, "float"); πE != nil {
									continue
								}
								πTemp004[1] = µfloat
								if πE = πg.CheckLocal(πF, µisinstance, "isinstance"); πE != nil {
									continue
								}
								if πTemp003, πE = µisinstance.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label22
								}
								goto Label23
								// line 322: if isinstance(value, basestring):
								πF.SetLineno(322)
							Label17:
								// line 323: yield buf + _encoder(value)
								πF.SetLineno(323)
								if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
									continue
								}
								πTemp004 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp004[0] = µvalue
								if πE = πg.CheckLocal(πF, µ_encoder, "_encoder"); πE != nil {
									continue
								}
								if πTemp005, πE = µ_encoder.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πTemp003, πE = πg.Add(πF, µbuf, πTemp005); πE != nil {
									continue
								}
								πF.PushCheckpoint(25)
								return πTemp003, nil
							Label25:
								πTemp005 = πSent
								goto Label24
								// line 324: elif value is None:
								πF.SetLineno(324)
							Label18:
								// line 325: yield buf + 'null'
								πF.SetLineno(325)
								if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.Add(πF, µbuf, ßnull.ToObject()); πE != nil {
									continue
								}
								πF.PushCheckpoint(26)
								return πTemp005, nil
							Label26:
								πTemp007 = πSent
								goto Label24
								// line 326: elif value is True:
								πF.SetLineno(326)
							Label19:
								// line 327: yield buf + 'true'
								πF.SetLineno(327)
								if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.Add(πF, µbuf, ßtrue.ToObject()); πE != nil {
									continue
								}
								πF.PushCheckpoint(27)
								return πTemp007, nil
							Label27:
								πTemp008 = πSent
								goto Label24
								// line 328: elif value is False:
								πF.SetLineno(328)
							Label20:
								// line 329: yield buf + 'false'
								πF.SetLineno(329)
								if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.Add(πF, µbuf, ßfalse.ToObject()); πE != nil {
									continue
								}
								πF.PushCheckpoint(28)
								return πTemp008, nil
							Label28:
								πTemp009 = πSent
								goto Label24
								// line 330: elif isinstance(value, (int, long)):
								πF.SetLineno(330)
							Label21:
								// line 331: yield buf + str(value)
								πF.SetLineno(331)
								if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
									continue
								}
								πTemp004 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp004[0] = µvalue
								if πE = πg.CheckLocal(πF, µstr, "str"); πE != nil {
									continue
								}
								if πTemp010, πE = µstr.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πTemp009, πE = πg.Add(πF, µbuf, πTemp010); πE != nil {
									continue
								}
								πF.PushCheckpoint(29)
								return πTemp009, nil
							Label29:
								πTemp010 = πSent
								goto Label24
								// line 332: elif isinstance(value, float):
								πF.SetLineno(332)
							Label22:
								// line 333: yield buf + _floatstr(value)
								πF.SetLineno(333)
								if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
									continue
								}
								πTemp004 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp004[0] = µvalue
								if πE = πg.CheckLocal(πF, µ_floatstr, "_floatstr"); πE != nil {
									continue
								}
								if πTemp011, πE = µ_floatstr.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πTemp010, πE = πg.Add(πF, µbuf, πTemp011); πE != nil {
									continue
								}
								πF.PushCheckpoint(30)
								return πTemp010, nil
							Label30:
								πTemp011 = πSent
								goto Label24
							Label23:
								// line 335: yield buf
								πF.SetLineno(335)
								if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
									continue
								}
								πF.PushCheckpoint(31)
								return µbuf, nil
							Label31:
								πTemp011 = πSent
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp004[0] = µvalue
								if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µtuple, "tuple"); πE != nil {
									continue
								}
								πTemp011 = πg.NewTuple2(µlist, µtuple).ToObject()
								πTemp004[1] = πTemp011
								if πE = πg.CheckLocal(πF, µisinstance, "isinstance"); πE != nil {
									continue
								}
								if πTemp011, πE = µisinstance.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πTemp006, πE = πg.IsTrue(πF, πTemp011); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label32
								}
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp004[0] = µvalue
								if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
									continue
								}
								πTemp004[1] = µdict
								if πE = πg.CheckLocal(πF, µisinstance, "isinstance"); πE != nil {
									continue
								}
								if πTemp011, πE = µisinstance.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πTemp006, πE = πg.IsTrue(πF, πTemp011); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label33
								}
								goto Label34
								// line 336: if isinstance(value, (list, tuple)):
								πF.SetLineno(336)
							Label32:
								// line 337: chunks = _iterencode_list(value, _current_indent_level)
								πF.SetLineno(337)
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp004[0] = µvalue
								if πE = πg.CheckLocal(πF, µ_current_indent_level, "_current_indent_level"); πE != nil {
									continue
								}
								πTemp004[1] = µ_current_indent_level
								if πE = πg.CheckLocal(πF, µ_iterencode_list, "_iterencode_list"); πE != nil {
									continue
								}
								if πTemp011, πE = µ_iterencode_list.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								µchunks = πTemp011
								goto Label35
								// line 338: elif isinstance(value, dict):
								πF.SetLineno(338)
							Label33:
								// line 339: chunks = _iterencode_dict(value, _current_indent_level)
								πF.SetLineno(339)
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp004[0] = µvalue
								if πE = πg.CheckLocal(πF, µ_current_indent_level, "_current_indent_level"); πE != nil {
									continue
								}
								πTemp004[1] = µ_current_indent_level
								if πE = πg.CheckLocal(πF, µ_iterencode_dict, "_iterencode_dict"); πE != nil {
									continue
								}
								if πTemp011, πE = µ_iterencode_dict.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								µchunks = πTemp011
								goto Label35
							Label34:
								// line 341: chunks = _iterencode(value, _current_indent_level)
								πF.SetLineno(341)
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp004[0] = µvalue
								if πE = πg.CheckLocal(πF, µ_current_indent_level, "_current_indent_level"); πE != nil {
									continue
								}
								πTemp004[1] = µ_current_indent_level
								if πE = πg.CheckLocal(πF, µ_iterencode, "_iterencode"); πE != nil {
									continue
								}
								if πTemp011, πE = µ_iterencode.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								µchunks = πTemp011
								goto Label35
							Label35:
								if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
									continue
								}
								if πTemp011, πE = πg.Iter(πF, µchunks); πE != nil {
									continue
								}
								πF.PushCheckpoint(37)
								πTemp006 = false
							Label36:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp006 {
									πF.PopCheckpoint()
									goto Label38
								}
								if πTemp013, πE = πg.Next(πF, πTemp011); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp012 = !isStop
								} else {
									πTemp012 = true
									µchunk = πTemp013
								}
								if πE != nil || !πTemp012 {
									continue
								}
								πF.PushCheckpoint(36)            
								// line 343: yield chunk
								πF.SetLineno(343)
								if πE = πg.CheckLocal(πF, µchunk, "chunk"); πE != nil {
									continue
								}
								πF.PushCheckpoint(39)
								return µchunk, nil
							Label39:
								πTemp013 = πSent
								continue
							Label37:
								if πE != nil || πR != nil {
									continue
								}
							Label38:
								goto Label24
							Label24:
								continue
							Label12:
								if πE != nil || πR != nil {
									continue
								}
							Label13:
								if πE = πg.CheckLocal(πF, µnewline_indent, "newline_indent"); πE != nil {
									continue
								}
								if πTemp011, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp001 = πg.GetBool(µnewline_indent != πTemp011).ToObject()
								if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp002 {
									goto Label40
								}
								goto Label41
								// line 344: if newline_indent is not None:
								πF.SetLineno(344)
							Label40:
								// line 345: _current_indent_level -= 1
								πF.SetLineno(345)
								if πE = πg.CheckLocal(πF, µ_current_indent_level, "_current_indent_level"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.ISub(πF, µ_current_indent_level, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								µ_current_indent_level = πTemp001
								// line 346: yield '\n' + (' ' * (_indent * _current_indent_level))
								πF.SetLineno(346)
								if πE = πg.CheckLocal(πF, µ_indent, "_indent"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µ_current_indent_level, "_current_indent_level"); πE != nil {
									continue
								}
								if πTemp013, πE = πg.Mul(πF, µ_indent, µ_current_indent_level); πE != nil {
									continue
								}
								if πTemp011, πE = πg.Mul(πF, πg.NewStr(" ").ToObject(), πTemp013); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Add(πF, πg.NewStr("\n").ToObject(), πTemp011); πE != nil {
									continue
								}
								πF.PushCheckpoint(42)
								return πTemp001, nil
							Label42:
								πTemp011 = πSent
								goto Label41
							Label41:
								// line 347: yield ']'
								πF.SetLineno(347)
								πF.PushCheckpoint(43)
								return πg.NewStr("]").ToObject(), nil
							Label43:
								πTemp011 = πSent
								if πE = πg.CheckLocal(πF, µmarkers, "markers"); πE != nil {
									continue
								}
								if πTemp013, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp011 = πg.GetBool(µmarkers != πTemp013).ToObject()
								if πTemp002, πE = πg.IsTrue(πF, πTemp011); πE != nil {
									continue
								}
								if πTemp002 {
									goto Label44
								}
								goto Label45
								// line 348: if markers is not None:
								πF.SetLineno(348)
							Label44:
								// line 349: del markers[markerid]
								πF.SetLineno(349)
								if πE = πg.CheckLocal(πF, µmarkers, "markers"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmarkerid, "markerid"); πE != nil {
									continue
								}
								πTemp011 = µmarkerid
								if πE = πg.DelItem(πF, µmarkers, πTemp011); πE != nil {
									continue
								}
								goto Label45
							Label45:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					µ_iterencode_list = πTemp001
					// line 351: def _iterencode_dict(dct, _current_indent_level):
					πF.SetLineno(351)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "dct", Def: nil}
					πTemp002[1] = πg.Param{Name: "_current_indent_level", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("_iterencode_dict", "build/src/__python__/json/encoder.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µdct *πg.Object = πArgs[0]; _ = µdct
						var µ_current_indent_level *πg.Object = πArgs[1]; _ = µ_current_indent_level
						var µmarkerid *πg.Object = πg.UnboundLocal; _ = µmarkerid
						var µnewline_indent *πg.Object = πg.UnboundLocal; _ = µnewline_indent
						var µitem_separator *πg.Object = πg.UnboundLocal; _ = µitem_separator
						var µfirst *πg.Object = πg.UnboundLocal; _ = µfirst
						var µitems *πg.Object = πg.UnboundLocal; _ = µitems
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
						var µchunks *πg.Object = πg.UnboundLocal; _ = µchunks
						var µchunk *πg.Object = πg.UnboundLocal; _ = µchunk
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
						var πTemp006 []πg.Param
						_ = πTemp006
						var πTemp007 πg.KWArgs
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πTemp013 bool
						_ = πTemp013
						var πTemp014 *πg.Object
						_ = πTemp014
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 32: goto Label32
								case 33: goto Label33
								case 3: goto Label3
								case 8: goto Label8
								case 44: goto Label44
								case 42: goto Label42
								case 43: goto Label43
								case 12: goto Label12
								case 45: goto Label45
								case 46: goto Label46
								case 47: goto Label47
								case 16: goto Label16
								case 17: goto Label17
								case 52: goto Label52
								case 53: goto Label53
								case 55: goto Label55
								case 58: goto Label58
								case 59: goto Label59
								case 31: goto Label31
								default: panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µdct, "dct"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.IsTrue(πF, µdct); πE != nil {
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
								// line 352: if not dct:
								πF.SetLineno(352)
							Label1:
								// line 353: yield '{}'
								πF.SetLineno(353)
								πF.PushCheckpoint(3)
								return πg.NewStr("{}").ToObject(), nil
							Label3:
								πTemp001 = πSent
								// line 354: return
								πF.SetLineno(354)
								πR = πg.None
								continue
								goto Label2
							Label2:
								if πE = πg.CheckLocal(πF, µmarkers, "markers"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp001 = πg.GetBool(µmarkers != πTemp003).ToObject()
								if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp002 {
									goto Label4
								}
								goto Label5
								// line 355: if markers is not None:
								πF.SetLineno(355)
							Label4:
								// line 356: markerid = id(dct)
								πF.SetLineno(356)
								πTemp004 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µdct, "dct"); πE != nil {
									continue
								}
								πTemp004[0] = µdct
								if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
									continue
								}
								if πTemp001, πE = µid.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								µmarkerid = πTemp001
								if πE = πg.CheckLocal(πF, µmarkerid, "markerid"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmarkers, "markers"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Contains(πF, µmarkers, µmarkerid); πE != nil {
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
								// line 357: if markerid in markers:
								πF.SetLineno(357)
							Label6:
								πTemp004 = πF.MakeArgs(1)
								πTemp004[0] = πg.NewStr("Circular reference detected").ToObject()
								if πE = πg.CheckLocal(πF, µValueError, "ValueError"); πE != nil {
									continue
								}
								if πTemp001, πE = µValueError.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								// line 358: raise ValueError("Circular reference detected")
								πF.SetLineno(358)
								πE = πF.Raise(πTemp001, nil, nil)
								continue
								goto Label7
							Label7:
								// line 359: markers[markerid] = dct
								πF.SetLineno(359)
								if πE = πg.CheckLocal(πF, µdct, "dct"); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdct); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmarkers, "markers"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmarkerid, "markerid"); πE != nil {
									continue
								}
								πTemp003 = µmarkerid
								if πE = πg.SetItem(πF, µmarkers, πTemp003, πTemp001); πE != nil {
									continue
								}
								goto Label5
							Label5:
								// line 360: yield '{'
								πF.SetLineno(360)
								πF.PushCheckpoint(8)
								return πg.NewStr("{").ToObject(), nil
							Label8:
								πTemp001 = πSent
								if πE = πg.CheckLocal(πF, µ_indent, "_indent"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp001 = πg.GetBool(µ_indent != πTemp003).ToObject()
								if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp002 {
									goto Label9
								}
								goto Label10
								// line 361: if _indent is not None:
								πF.SetLineno(361)
							Label9:
								// line 362: _current_indent_level += 1
								πF.SetLineno(362)
								if πE = πg.CheckLocal(πF, µ_current_indent_level, "_current_indent_level"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.IAdd(πF, µ_current_indent_level, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								µ_current_indent_level = πTemp001
								// line 363: newline_indent = '\n' + (' ' * (_indent * _current_indent_level))
								πF.SetLineno(363)
								if πE = πg.CheckLocal(πF, µ_indent, "_indent"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µ_current_indent_level, "_current_indent_level"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.Mul(πF, µ_indent, µ_current_indent_level); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Mul(πF, πg.NewStr(" ").ToObject(), πTemp005); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Add(πF, πg.NewStr("\n").ToObject(), πTemp003); πE != nil {
									continue
								}
								µnewline_indent = πTemp001
								// line 364: item_separator = _item_separator + newline_indent
								πF.SetLineno(364)
								if πE = πg.CheckLocal(πF, µ_item_separator, "_item_separator"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µnewline_indent, "newline_indent"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Add(πF, µ_item_separator, µnewline_indent); πE != nil {
									continue
								}
								µitem_separator = πTemp001
								// line 365: yield newline_indent
								πF.SetLineno(365)
								if πE = πg.CheckLocal(πF, µnewline_indent, "newline_indent"); πE != nil {
									continue
								}
								πF.PushCheckpoint(12)
								return µnewline_indent, nil
							Label12:
								πTemp001 = πSent
								goto Label11
							Label10:
								// line 367: newline_indent = None
								πF.SetLineno(367)
								if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								µnewline_indent = πTemp001
								// line 368: item_separator = _item_separator
								πF.SetLineno(368)
								if πE = πg.CheckLocal(πF, µ_item_separator, "_item_separator"); πE != nil {
									continue
								}
								µitem_separator = µ_item_separator
								goto Label11
							Label11:
								// line 369: first = True
								πF.SetLineno(369)
								if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								µfirst = πTemp001
								if πE = πg.CheckLocal(πF, µ_sort_keys, "_sort_keys"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.IsTrue(πF, µ_sort_keys); πE != nil {
									continue
								}
								if πTemp002 {
									goto Label13
								}
								goto Label14
								// line 370: if _sort_keys:
								πF.SetLineno(370)
							Label13:
								// line 371: items = sorted(dct.items(), key=lambda kv: kv[0])
								πF.SetLineno(371)
								πTemp004 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µdct, "dct"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µdct, ßitems, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
									continue
								}
								πTemp004[0] = πTemp003
								πTemp006 = make([]πg.Param, 1)
								πTemp006[0] = πg.Param{Name: "kv", Def: nil}
								πTemp001 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/json/encoder.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
									var µkv *πg.Object = πArgs[0]; _ = µkv
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
										// line 371: items = sorted(dct.items(), key=lambda kv: kv[0])
										πF.SetLineno(371)
										πTemp001 = πg.NewInt(0).ToObject()
										if πE = πg.CheckLocal(πF, µkv, "kv"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetItem(πF, µkv, πTemp001); πE != nil {
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
								πTemp007 = πg.KWArgs{
									{"key", πTemp001},
								}
								if πTemp001, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp001.Call(πF, πTemp004, πTemp007); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								µitems = πTemp003
								goto Label15
							Label14:
								// line 373: items = dct.iteritems()
								πF.SetLineno(373)
								if πE = πg.CheckLocal(πF, µdct, "dct"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µdct, ßiteritems, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
									continue
								}
								µitems = πTemp003
								goto Label15
							Label15:
								if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µitems); πE != nil {
									continue
								}
								πF.PushCheckpoint(17)
								πTemp002 = false
							Label16:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp002 {
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
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp009}}}, πTemp003); πE != nil {
										continue
									}
									µkey = πTemp005
									µvalue = πTemp009
								}
								if πE != nil || !πTemp008 {
									continue
								}
								πF.PushCheckpoint(16)            
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
									continue
								}
								πTemp004[0] = µkey
								if πE = πg.CheckLocal(πF, µbasestring, "basestring"); πE != nil {
									continue
								}
								πTemp004[1] = µbasestring
								if πE = πg.CheckLocal(πF, µisinstance, "isinstance"); πE != nil {
									continue
								}
								if πTemp003, πE = µisinstance.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp008 {
									goto Label19
								}
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
									continue
								}
								πTemp004[0] = µkey
								if πE = πg.CheckLocal(πF, µfloat, "float"); πE != nil {
									continue
								}
								πTemp004[1] = µfloat
								if πE = πg.CheckLocal(πF, µisinstance, "isinstance"); πE != nil {
									continue
								}
								if πTemp003, πE = µisinstance.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp008 {
									goto Label20
								}
								if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πTemp003 = πg.GetBool(µkey == πTemp005).ToObject()
								if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp008 {
									goto Label21
								}
								if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								πTemp003 = πg.GetBool(µkey == πTemp005).ToObject()
								if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp008 {
									goto Label22
								}
								if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp003 = πg.GetBool(µkey == πTemp005).ToObject()
								if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp008 {
									goto Label23
								}
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
									continue
								}
								πTemp004[0] = µkey
								if πE = πg.CheckLocal(πF, µint, "int"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µlong, "long"); πE != nil {
									continue
								}
								πTemp003 = πg.NewTuple2(µint, µlong).ToObject()
								πTemp004[1] = πTemp003
								if πE = πg.CheckLocal(πF, µisinstance, "isinstance"); πE != nil {
									continue
								}
								if πTemp003, πE = µisinstance.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp008 {
									goto Label24
								}
								if πE = πg.CheckLocal(πF, µ_skipkeys, "_skipkeys"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.IsTrue(πF, µ_skipkeys); πE != nil {
									continue
								}
								if πTemp008 {
									goto Label25
								}
								goto Label26
								// line 375: if isinstance(key, basestring):
								πF.SetLineno(375)
							Label19:
								// line 376: pass
								πF.SetLineno(376)
								goto Label27
								// line 379: elif isinstance(key, float):
								πF.SetLineno(379)
							Label20:
								// line 380: key = _floatstr(key)
								πF.SetLineno(380)
								πTemp004 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
									continue
								}
								πTemp004[0] = µkey
								if πE = πg.CheckLocal(πF, µ_floatstr, "_floatstr"); πE != nil {
									continue
								}
								if πTemp003, πE = µ_floatstr.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								µkey = πTemp003
								goto Label27
								// line 381: elif key is True:
								πF.SetLineno(381)
							Label21:
								// line 382: key = 'true'
								πF.SetLineno(382)
								µkey = ßtrue.ToObject()
								goto Label27
								// line 383: elif key is False:
								πF.SetLineno(383)
							Label22:
								// line 384: key = 'false'
								πF.SetLineno(384)
								µkey = ßfalse.ToObject()
								goto Label27
								// line 385: elif key is None:
								πF.SetLineno(385)
							Label23:
								// line 386: key = 'null'
								πF.SetLineno(386)
								µkey = ßnull.ToObject()
								goto Label27
								// line 387: elif isinstance(key, (int, long)):
								πF.SetLineno(387)
							Label24:
								// line 388: key = str(key)
								πF.SetLineno(388)
								πTemp004 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
									continue
								}
								πTemp004[0] = µkey
								if πE = πg.CheckLocal(πF, µstr, "str"); πE != nil {
									continue
								}
								if πTemp003, πE = µstr.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								µkey = πTemp003
								goto Label27
								// line 389: elif _skipkeys:
								πF.SetLineno(389)
							Label25:
								// line 390: continue
								πF.SetLineno(390)
								continue
								goto Label27
							Label26:
								πTemp004 = πF.MakeArgs(1)
								πTemp010 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
									continue
								}
								πTemp010[0] = µkey
								if πTemp009, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
									continue
								}
								if πTemp011, πE = πTemp009.Call(πF, πTemp010, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp010)
								if πTemp005, πE = πg.Add(πF, πg.NewStr("key ").ToObject(), πTemp011); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Add(πF, πTemp005, πg.NewStr(" is not a string").ToObject()); πE != nil {
									continue
								}
								πTemp004[0] = πTemp003
								if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								// line 392: raise TypeError("key " + repr(key) + " is not a string")
								πF.SetLineno(392)
								πE = πF.Raise(πTemp005, nil, nil)
								continue
								goto Label27
							Label27:
								if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.IsTrue(πF, µfirst); πE != nil {
									continue
								}
								if πTemp008 {
									goto Label28
								}
								goto Label29
								// line 393: if first:
								πF.SetLineno(393)
							Label28:
								// line 394: first = False
								πF.SetLineno(394)
								if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								µfirst = πTemp003
								goto Label30
							Label29:
								// line 396: yield item_separator
								πF.SetLineno(396)
								if πE = πg.CheckLocal(πF, µitem_separator, "item_separator"); πE != nil {
									continue
								}
								πF.PushCheckpoint(31)
								return µitem_separator, nil
							Label31:
								πTemp003 = πSent
								goto Label30
							Label30:
								// line 397: yield _encoder(key)
								πF.SetLineno(397)
								πTemp004 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
									continue
								}
								πTemp004[0] = µkey
								if πE = πg.CheckLocal(πF, µ_encoder, "_encoder"); πE != nil {
									continue
								}
								if πTemp003, πE = µ_encoder.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πF.PushCheckpoint(32)
								return πTemp003, nil
							Label32:
								πTemp005 = πSent
								// line 398: yield _key_separator
								πF.SetLineno(398)
								if πE = πg.CheckLocal(πF, µ_key_separator, "_key_separator"); πE != nil {
									continue
								}
								πF.PushCheckpoint(33)
								return µ_key_separator, nil
							Label33:
								πTemp005 = πSent
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp004[0] = µvalue
								if πE = πg.CheckLocal(πF, µbasestring, "basestring"); πE != nil {
									continue
								}
								πTemp004[1] = µbasestring
								if πE = πg.CheckLocal(πF, µisinstance, "isinstance"); πE != nil {
									continue
								}
								if πTemp005, πE = µisinstance.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								if πTemp008 {
									goto Label34
								}
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp005 = πg.GetBool(µvalue == πTemp009).ToObject()
								if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								if πTemp008 {
									goto Label35
								}
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πTemp005 = πg.GetBool(µvalue == πTemp009).ToObject()
								if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								if πTemp008 {
									goto Label36
								}
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								πTemp005 = πg.GetBool(µvalue == πTemp009).ToObject()
								if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								if πTemp008 {
									goto Label37
								}
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp004[0] = µvalue
								if πE = πg.CheckLocal(πF, µint, "int"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µlong, "long"); πE != nil {
									continue
								}
								πTemp005 = πg.NewTuple2(µint, µlong).ToObject()
								πTemp004[1] = πTemp005
								if πE = πg.CheckLocal(πF, µisinstance, "isinstance"); πE != nil {
									continue
								}
								if πTemp005, πE = µisinstance.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								if πTemp008 {
									goto Label38
								}
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp004[0] = µvalue
								if πE = πg.CheckLocal(πF, µfloat, "float"); πE != nil {
									continue
								}
								πTemp004[1] = µfloat
								if πE = πg.CheckLocal(πF, µisinstance, "isinstance"); πE != nil {
									continue
								}
								if πTemp005, πE = µisinstance.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								if πTemp008 {
									goto Label39
								}
								goto Label40
								// line 399: if isinstance(value, basestring):
								πF.SetLineno(399)
							Label34:
								// line 400: yield _encoder(value)
								πF.SetLineno(400)
								πTemp004 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp004[0] = µvalue
								if πE = πg.CheckLocal(πF, µ_encoder, "_encoder"); πE != nil {
									continue
								}
								if πTemp005, πE = µ_encoder.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πF.PushCheckpoint(42)
								return πTemp005, nil
							Label42:
								πTemp009 = πSent
								goto Label41
								// line 401: elif value is None:
								πF.SetLineno(401)
							Label35:
								// line 402: yield 'null'
								πF.SetLineno(402)
								πF.PushCheckpoint(43)
								return ßnull.ToObject(), nil
							Label43:
								πTemp009 = πSent
								goto Label41
								// line 403: elif value is True:
								πF.SetLineno(403)
							Label36:
								// line 404: yield 'true'
								πF.SetLineno(404)
								πF.PushCheckpoint(44)
								return ßtrue.ToObject(), nil
							Label44:
								πTemp009 = πSent
								goto Label41
								// line 405: elif value is False:
								πF.SetLineno(405)
							Label37:
								// line 406: yield 'false'
								πF.SetLineno(406)
								πF.PushCheckpoint(45)
								return ßfalse.ToObject(), nil
							Label45:
								πTemp009 = πSent
								goto Label41
								// line 407: elif isinstance(value, (int, long)):
								πF.SetLineno(407)
							Label38:
								// line 408: yield str(value)
								πF.SetLineno(408)
								πTemp004 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp004[0] = µvalue
								if πE = πg.CheckLocal(πF, µstr, "str"); πE != nil {
									continue
								}
								if πTemp009, πE = µstr.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πF.PushCheckpoint(46)
								return πTemp009, nil
							Label46:
								πTemp011 = πSent
								goto Label41
								// line 409: elif isinstance(value, float):
								πF.SetLineno(409)
							Label39:
								// line 410: yield _floatstr(value)
								πF.SetLineno(410)
								πTemp004 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp004[0] = µvalue
								if πE = πg.CheckLocal(πF, µ_floatstr, "_floatstr"); πE != nil {
									continue
								}
								if πTemp011, πE = µ_floatstr.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πF.PushCheckpoint(47)
								return πTemp011, nil
							Label47:
								πTemp012 = πSent
								goto Label41
							Label40:
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp004[0] = µvalue
								if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µtuple, "tuple"); πE != nil {
									continue
								}
								πTemp012 = πg.NewTuple2(µlist, µtuple).ToObject()
								πTemp004[1] = πTemp012
								if πE = πg.CheckLocal(πF, µisinstance, "isinstance"); πE != nil {
									continue
								}
								if πTemp012, πE = µisinstance.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πTemp008, πE = πg.IsTrue(πF, πTemp012); πE != nil {
									continue
								}
								if πTemp008 {
									goto Label48
								}
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp004[0] = µvalue
								if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
									continue
								}
								πTemp004[1] = µdict
								if πE = πg.CheckLocal(πF, µisinstance, "isinstance"); πE != nil {
									continue
								}
								if πTemp012, πE = µisinstance.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								if πTemp008, πE = πg.IsTrue(πF, πTemp012); πE != nil {
									continue
								}
								if πTemp008 {
									goto Label49
								}
								goto Label50
								// line 412: if isinstance(value, (list, tuple)):
								πF.SetLineno(412)
							Label48:
								// line 413: chunks = _iterencode_list(value, _current_indent_level)
								πF.SetLineno(413)
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp004[0] = µvalue
								if πE = πg.CheckLocal(πF, µ_current_indent_level, "_current_indent_level"); πE != nil {
									continue
								}
								πTemp004[1] = µ_current_indent_level
								if πE = πg.CheckLocal(πF, µ_iterencode_list, "_iterencode_list"); πE != nil {
									continue
								}
								if πTemp012, πE = µ_iterencode_list.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								µchunks = πTemp012
								goto Label51
								// line 414: elif isinstance(value, dict):
								πF.SetLineno(414)
							Label49:
								// line 415: chunks = _iterencode_dict(value, _current_indent_level)
								πF.SetLineno(415)
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp004[0] = µvalue
								if πE = πg.CheckLocal(πF, µ_current_indent_level, "_current_indent_level"); πE != nil {
									continue
								}
								πTemp004[1] = µ_current_indent_level
								if πE = πg.CheckLocal(πF, µ_iterencode_dict, "_iterencode_dict"); πE != nil {
									continue
								}
								if πTemp012, πE = µ_iterencode_dict.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								µchunks = πTemp012
								goto Label51
							Label50:
								// line 417: chunks = _iterencode(value, _current_indent_level)
								πF.SetLineno(417)
								πTemp004 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp004[0] = µvalue
								if πE = πg.CheckLocal(πF, µ_current_indent_level, "_current_indent_level"); πE != nil {
									continue
								}
								πTemp004[1] = µ_current_indent_level
								if πE = πg.CheckLocal(πF, µ_iterencode, "_iterencode"); πE != nil {
									continue
								}
								if πTemp012, πE = µ_iterencode.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								µchunks = πTemp012
								goto Label51
							Label51:
								if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
									continue
								}
								if πTemp012, πE = πg.Iter(πF, µchunks); πE != nil {
									continue
								}
								πF.PushCheckpoint(53)
								πTemp008 = false
							Label52:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp008 {
									πF.PopCheckpoint()
									goto Label54
								}
								if πTemp014, πE = πg.Next(πF, πTemp012); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp013 = !isStop
								} else {
									πTemp013 = true
									µchunk = πTemp014
								}
								if πE != nil || !πTemp013 {
									continue
								}
								πF.PushCheckpoint(52)            
								// line 419: yield chunk
								πF.SetLineno(419)
								if πE = πg.CheckLocal(πF, µchunk, "chunk"); πE != nil {
									continue
								}
								πF.PushCheckpoint(55)
								return µchunk, nil
							Label55:
								πTemp014 = πSent
								continue
							Label53:
								if πE != nil || πR != nil {
									continue
								}
							Label54:
								goto Label41
							Label41:
								continue
							Label17:
								if πE != nil || πR != nil {
									continue
								}
							Label18:
								if πE = πg.CheckLocal(πF, µnewline_indent, "newline_indent"); πE != nil {
									continue
								}
								if πTemp012, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp001 = πg.GetBool(µnewline_indent != πTemp012).ToObject()
								if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp002 {
									goto Label56
								}
								goto Label57
								// line 420: if newline_indent is not None:
								πF.SetLineno(420)
							Label56:
								// line 421: _current_indent_level -= 1
								πF.SetLineno(421)
								if πE = πg.CheckLocal(πF, µ_current_indent_level, "_current_indent_level"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.ISub(πF, µ_current_indent_level, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								µ_current_indent_level = πTemp001
								// line 422: yield '\n' + (' ' * (_indent * _current_indent_level))
								πF.SetLineno(422)
								if πE = πg.CheckLocal(πF, µ_indent, "_indent"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µ_current_indent_level, "_current_indent_level"); πE != nil {
									continue
								}
								if πTemp014, πE = πg.Mul(πF, µ_indent, µ_current_indent_level); πE != nil {
									continue
								}
								if πTemp012, πE = πg.Mul(πF, πg.NewStr(" ").ToObject(), πTemp014); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Add(πF, πg.NewStr("\n").ToObject(), πTemp012); πE != nil {
									continue
								}
								πF.PushCheckpoint(58)
								return πTemp001, nil
							Label58:
								πTemp012 = πSent
								goto Label57
							Label57:
								// line 423: yield '}'
								πF.SetLineno(423)
								πF.PushCheckpoint(59)
								return πg.NewStr("}").ToObject(), nil
							Label59:
								πTemp012 = πSent
								if πE = πg.CheckLocal(πF, µmarkers, "markers"); πE != nil {
									continue
								}
								if πTemp014, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp012 = πg.GetBool(µmarkers != πTemp014).ToObject()
								if πTemp002, πE = πg.IsTrue(πF, πTemp012); πE != nil {
									continue
								}
								if πTemp002 {
									goto Label60
								}
								goto Label61
								// line 424: if markers is not None:
								πF.SetLineno(424)
							Label60:
								// line 425: del markers[markerid]
								πF.SetLineno(425)
								if πE = πg.CheckLocal(πF, µmarkers, "markers"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmarkerid, "markerid"); πE != nil {
									continue
								}
								πTemp012 = µmarkerid
								if πE = πg.DelItem(πF, µmarkers, πTemp012); πE != nil {
									continue
								}
								goto Label61
							Label61:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					µ_iterencode_dict = πTemp003
					// line 427: def _iterencode(o, _current_indent_level):
					πF.SetLineno(427)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "o", Def: nil}
					πTemp002[1] = πg.Param{Name: "_current_indent_level", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("_iterencode", "build/src/__python__/json/encoder.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µo *πg.Object = πArgs[0]; _ = µo
						var µ_current_indent_level *πg.Object = πArgs[1]; _ = µ_current_indent_level
						var µchunk *πg.Object = πg.UnboundLocal; _ = µchunk
						var µmarkerid *πg.Object = πg.UnboundLocal; _ = µmarkerid
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 32: goto Label32
								case 11: goto Label11
								case 12: goto Label12
								case 13: goto Label13
								case 14: goto Label14
								case 15: goto Label15
								case 16: goto Label16
								case 17: goto Label17
								case 18: goto Label18
								case 20: goto Label20
								case 21: goto Label21
								case 22: goto Label22
								case 24: goto Label24
								case 29: goto Label29
								case 30: goto Label30
								default: panic("unexpected function state")
								}
								πTemp001 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
									continue
								}
								πTemp001[0] = µo
								if πE = πg.CheckLocal(πF, µbasestring, "basestring"); πE != nil {
									continue
								}
								πTemp001[1] = µbasestring
								if πE = πg.CheckLocal(πF, µisinstance, "isinstance"); πE != nil {
									continue
								}
								if πTemp002, πE = µisinstance.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label1
								}
								if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp002 = πg.GetBool(µo == πTemp004).ToObject()
								if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label2
								}
								if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πTemp002 = πg.GetBool(µo == πTemp004).ToObject()
								if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label3
								}
								if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								πTemp002 = πg.GetBool(µo == πTemp004).ToObject()
								if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label4
								}
								πTemp001 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
									continue
								}
								πTemp001[0] = µo
								if πE = πg.CheckLocal(πF, µint, "int"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µlong, "long"); πE != nil {
									continue
								}
								πTemp002 = πg.NewTuple2(µint, µlong).ToObject()
								πTemp001[1] = πTemp002
								if πE = πg.CheckLocal(πF, µisinstance, "isinstance"); πE != nil {
									continue
								}
								if πTemp002, πE = µisinstance.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label5
								}
								πTemp001 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
									continue
								}
								πTemp001[0] = µo
								if πE = πg.CheckLocal(πF, µfloat, "float"); πE != nil {
									continue
								}
								πTemp001[1] = µfloat
								if πE = πg.CheckLocal(πF, µisinstance, "isinstance"); πE != nil {
									continue
								}
								if πTemp002, πE = µisinstance.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label6
								}
								πTemp001 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
									continue
								}
								πTemp001[0] = µo
								if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µtuple, "tuple"); πE != nil {
									continue
								}
								πTemp002 = πg.NewTuple2(µlist, µtuple).ToObject()
								πTemp001[1] = πTemp002
								if πE = πg.CheckLocal(πF, µisinstance, "isinstance"); πE != nil {
									continue
								}
								if πTemp002, πE = µisinstance.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label7
								}
								πTemp001 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
									continue
								}
								πTemp001[0] = µo
								if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
									continue
								}
								πTemp001[1] = µdict
								if πE = πg.CheckLocal(πF, µisinstance, "isinstance"); πE != nil {
									continue
								}
								if πTemp002, πE = µisinstance.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label8
								}
								goto Label9
								// line 428: if isinstance(o, basestring):
								πF.SetLineno(428)
							Label1:
								// line 429: yield _encoder(o)
								πF.SetLineno(429)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
									continue
								}
								πTemp001[0] = µo
								if πE = πg.CheckLocal(πF, µ_encoder, "_encoder"); πE != nil {
									continue
								}
								if πTemp002, πE = µ_encoder.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								πF.PushCheckpoint(11)
								return πTemp002, nil
							Label11:
								πTemp004 = πSent
								goto Label10
								// line 430: elif o is None:
								πF.SetLineno(430)
							Label2:
								// line 431: yield 'null'
								πF.SetLineno(431)
								πF.PushCheckpoint(12)
								return ßnull.ToObject(), nil
							Label12:
								πTemp004 = πSent
								goto Label10
								// line 432: elif o is True:
								πF.SetLineno(432)
							Label3:
								// line 433: yield 'true'
								πF.SetLineno(433)
								πF.PushCheckpoint(13)
								return ßtrue.ToObject(), nil
							Label13:
								πTemp004 = πSent
								goto Label10
								// line 434: elif o is False:
								πF.SetLineno(434)
							Label4:
								// line 435: yield 'false'
								πF.SetLineno(435)
								πF.PushCheckpoint(14)
								return ßfalse.ToObject(), nil
							Label14:
								πTemp004 = πSent
								goto Label10
								// line 436: elif isinstance(o, (int, long)):
								πF.SetLineno(436)
							Label5:
								// line 437: yield str(o)
								πF.SetLineno(437)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
									continue
								}
								πTemp001[0] = µo
								if πE = πg.CheckLocal(πF, µstr, "str"); πE != nil {
									continue
								}
								if πTemp004, πE = µstr.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								πF.PushCheckpoint(15)
								return πTemp004, nil
							Label15:
								πTemp005 = πSent
								goto Label10
								// line 438: elif isinstance(o, float):
								πF.SetLineno(438)
							Label6:
								// line 439: yield _floatstr(o)
								πF.SetLineno(439)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
									continue
								}
								πTemp001[0] = µo
								if πE = πg.CheckLocal(πF, µ_floatstr, "_floatstr"); πE != nil {
									continue
								}
								if πTemp005, πE = µ_floatstr.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								πF.PushCheckpoint(16)
								return πTemp005, nil
							Label16:
								πTemp006 = πSent
								goto Label10
								// line 440: elif isinstance(o, (list, tuple)):
								πF.SetLineno(440)
							Label7:
								πTemp001 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
									continue
								}
								πTemp001[0] = µo
								if πE = πg.CheckLocal(πF, µ_current_indent_level, "_current_indent_level"); πE != nil {
									continue
								}
								πTemp001[1] = µ_current_indent_level
								if πE = πg.CheckLocal(πF, µ_iterencode_list, "_iterencode_list"); πE != nil {
									continue
								}
								if πTemp007, πE = µ_iterencode_list.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp006, πE = πg.Iter(πF, πTemp007); πE != nil {
									continue
								}
								πF.PushCheckpoint(18)
								πTemp003 = false
							Label17:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp003 {
									πF.PopCheckpoint()
									goto Label19
								}
								if πTemp007, πE = πg.Next(πF, πTemp006); πE != nil {
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
									µchunk = πTemp007
								}
								if πE != nil || !πTemp008 {
									continue
								}
								πF.PushCheckpoint(17)            
								// line 442: yield chunk
								πF.SetLineno(442)
								if πE = πg.CheckLocal(πF, µchunk, "chunk"); πE != nil {
									continue
								}
								πF.PushCheckpoint(20)
								return µchunk, nil
							Label20:
								πTemp007 = πSent
								continue
							Label18:
								if πE != nil || πR != nil {
									continue
								}
							Label19:
								goto Label10
								// line 443: elif isinstance(o, dict):
								πF.SetLineno(443)
							Label8:
								πTemp001 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
									continue
								}
								πTemp001[0] = µo
								if πE = πg.CheckLocal(πF, µ_current_indent_level, "_current_indent_level"); πE != nil {
									continue
								}
								πTemp001[1] = µ_current_indent_level
								if πE = πg.CheckLocal(πF, µ_iterencode_dict, "_iterencode_dict"); πE != nil {
									continue
								}
								if πTemp007, πE = µ_iterencode_dict.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp006, πE = πg.Iter(πF, πTemp007); πE != nil {
									continue
								}
								πF.PushCheckpoint(22)
								πTemp003 = false
							Label21:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp003 {
									πF.PopCheckpoint()
									goto Label23
								}
								if πTemp007, πE = πg.Next(πF, πTemp006); πE != nil {
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
									µchunk = πTemp007
								}
								if πE != nil || !πTemp008 {
									continue
								}
								πF.PushCheckpoint(21)            
								// line 445: yield chunk
								πF.SetLineno(445)
								if πE = πg.CheckLocal(πF, µchunk, "chunk"); πE != nil {
									continue
								}
								πF.PushCheckpoint(24)
								return µchunk, nil
							Label24:
								πTemp007 = πSent
								continue
							Label22:
								if πE != nil || πR != nil {
									continue
								}
							Label23:
								goto Label10
							Label9:
								if πE = πg.CheckLocal(πF, µmarkers, "markers"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp006 = πg.GetBool(µmarkers != πTemp007).ToObject()
								if πTemp003, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label25
								}
								goto Label26
								// line 447: if markers is not None:
								πF.SetLineno(447)
							Label25:
								// line 448: markerid = id(o)
								πF.SetLineno(448)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
									continue
								}
								πTemp001[0] = µo
								if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
									continue
								}
								if πTemp006, πE = µid.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µmarkerid = πTemp006
								if πE = πg.CheckLocal(πF, µmarkerid, "markerid"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmarkers, "markers"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Contains(πF, µmarkers, µmarkerid); πE != nil {
									continue
								}
								πTemp006 = πg.GetBool(πTemp003).ToObject()
								if πTemp003, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label27
								}
								goto Label28
								// line 449: if markerid in markers:
								πF.SetLineno(449)
							Label27:
								πTemp001 = πF.MakeArgs(1)
								πTemp001[0] = πg.NewStr("Circular reference detected").ToObject()
								if πE = πg.CheckLocal(πF, µValueError, "ValueError"); πE != nil {
									continue
								}
								if πTemp006, πE = µValueError.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								// line 450: raise ValueError("Circular reference detected")
								πF.SetLineno(450)
								πE = πF.Raise(πTemp006, nil, nil)
								continue
								goto Label28
							Label28:
								// line 451: markers[markerid] = o
								πF.SetLineno(451)
								if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, µo); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmarkers, "markers"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmarkerid, "markerid"); πE != nil {
									continue
								}
								πTemp007 = µmarkerid
								if πE = πg.SetItem(πF, µmarkers, πTemp007, πTemp006); πE != nil {
									continue
								}
								goto Label26
							Label26:
								// line 452: o = _default(o)
								πF.SetLineno(452)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
									continue
								}
								πTemp001[0] = µo
								if πE = πg.CheckLocal(πF, µ_default, "_default"); πE != nil {
									continue
								}
								if πTemp006, πE = µ_default.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µo = πTemp006
								πTemp001 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
									continue
								}
								πTemp001[0] = µo
								if πE = πg.CheckLocal(πF, µ_current_indent_level, "_current_indent_level"); πE != nil {
									continue
								}
								πTemp001[1] = µ_current_indent_level
								if πE = πg.CheckLocal(πF, µ_iterencode, "_iterencode"); πE != nil {
									continue
								}
								if πTemp007, πE = µ_iterencode.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp006, πE = πg.Iter(πF, πTemp007); πE != nil {
									continue
								}
								πF.PushCheckpoint(30)
								πTemp003 = false
							Label29:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp003 {
									πF.PopCheckpoint()
									goto Label31
								}
								if πTemp007, πE = πg.Next(πF, πTemp006); πE != nil {
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
									µchunk = πTemp007
								}
								if πE != nil || !πTemp008 {
									continue
								}
								πF.PushCheckpoint(29)            
								// line 454: yield chunk
								πF.SetLineno(454)
								if πE = πg.CheckLocal(πF, µchunk, "chunk"); πE != nil {
									continue
								}
								πF.PushCheckpoint(32)
								return µchunk, nil
							Label32:
								πTemp007 = πSent
								continue
							Label30:
								if πE != nil || πR != nil {
									continue
								}
							Label31:
								if πE = πg.CheckLocal(πF, µmarkers, "markers"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp006 = πg.GetBool(µmarkers != πTemp007).ToObject()
								if πTemp003, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label33
								}
								goto Label34
								// line 455: if markers is not None:
								πF.SetLineno(455)
							Label33:
								// line 456: del markers[markerid]
								πF.SetLineno(456)
								if πE = πg.CheckLocal(πF, µmarkers, "markers"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmarkerid, "markerid"); πE != nil {
									continue
								}
								πTemp006 = µmarkerid
								if πE = πg.DelItem(πF, µmarkers, πTemp006); πE != nil {
									continue
								}
								goto Label34
							Label34:
								goto Label10
							Label10:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					µ_iterencode = πTemp004
					// line 458: return _iterencode
					πF.SetLineno(458)
					if πE = πg.CheckLocal(πF, µ_iterencode, "_iterencode"); πE != nil {
						continue
					}
					πR = µ_iterencode
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_make_iterencode.ToObject(), πTemp007); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("json.encoder", Code)
}
