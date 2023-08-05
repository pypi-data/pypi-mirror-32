package json_scanner
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/json_scanner.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßDOTALL := πg.InternStr("DOTALL")
		ßFalse := πg.InternStr("False")
		ßI := πg.InternStr("I")
		ßIndexError := πg.InternStr("IndexError")
		ßInfinity := πg.InternStr("Infinity")
		ßMULTILINE := πg.InternStr("MULTILINE")
		ßN := πg.InternStr("N")
		ßNUMBER_RE := πg.InternStr("NUMBER_RE")
		ßNaN := πg.InternStr("NaN")
		ßNone := πg.InternStr("None")
		ßStopIteration := πg.InternStr("StopIteration")
		ßTrue := πg.InternStr("True")
		ßVERBOSE := πg.InternStr("VERBOSE")
		ß__all__ := πg.InternStr("__all__")
		ßc_make_scanner := πg.InternStr("c_make_scanner")
		ßcompile := πg.InternStr("compile")
		ßencoding := πg.InternStr("encoding")
		ßend := πg.InternStr("end")
		ßf := πg.InternStr("f")
		ßfalse := πg.InternStr("false")
		ßgroups := πg.InternStr("groups")
		ßmake_scanner := πg.InternStr("make_scanner")
		ßmatch := πg.InternStr("match")
		ßn := πg.InternStr("n")
		ßnull := πg.InternStr("null")
		ßobject_hook := πg.InternStr("object_hook")
		ßobject_pairs_hook := πg.InternStr("object_pairs_hook")
		ßparse_array := πg.InternStr("parse_array")
		ßparse_constant := πg.InternStr("parse_constant")
		ßparse_float := πg.InternStr("parse_float")
		ßparse_int := πg.InternStr("parse_int")
		ßparse_object := πg.InternStr("parse_object")
		ßparse_string := πg.InternStr("parse_string")
		ßpy_make_scanner := πg.InternStr("py_make_scanner")
		ßre := πg.InternStr("re")
		ßstrict := πg.InternStr("strict")
		ßt := πg.InternStr("t")
		ßtrue := πg.InternStr("true")
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
		var πTemp007 []πg.Param
		_ = πTemp007
		var πTemp008 bool
		_ = πTemp008
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """JSON token scanner
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
			// line 8: c_make_scanner = None
			πF.SetLineno(8)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßc_make_scanner.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 10: __all__ = ['make_scanner']
			πF.SetLineno(10)
			πTemp002 = make([]*πg.Object, 1)
			πTemp002[0] = ßmake_scanner.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 12: NUMBER_RE = re.compile(
			πF.SetLineno(12)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("(-?(?:0|[1-9]\\d*))(\\.\\d+)?([eE][-+]?\\d+)?").ToObject()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßVERBOSE, nil); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßMULTILINE, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.Or(πF, πTemp005, πTemp006); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßDOTALL, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πg.Or(πF, πTemp003, πTemp005); πE != nil {
				continue
			}
			πTemp002[1] = πTemp001
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
			if πE = πF.Globals().SetItem(πF, ßNUMBER_RE.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 16: def py_make_scanner(context):
			πF.SetLineno(16)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "context", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("py_make_scanner", "build/src/__python__/json_scanner.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcontext *πg.Object = πArgs[0]; _ = µcontext
				var µparse_object *πg.Object = πg.UnboundLocal; _ = µparse_object
				var µparse_array *πg.Object = πg.UnboundLocal; _ = µparse_array
				var µparse_string *πg.Object = πg.UnboundLocal; _ = µparse_string
				var µmatch_number *πg.Object = πg.UnboundLocal; _ = µmatch_number
				var µencoding *πg.Object = πg.UnboundLocal; _ = µencoding
				var µstrict *πg.Object = πg.UnboundLocal; _ = µstrict
				var µparse_float *πg.Object = πg.UnboundLocal; _ = µparse_float
				var µparse_int *πg.Object = πg.UnboundLocal; _ = µparse_int
				var µparse_constant *πg.Object = πg.UnboundLocal; _ = µparse_constant
				var µobject_hook *πg.Object = πg.UnboundLocal; _ = µobject_hook
				var µobject_pairs_hook *πg.Object = πg.UnboundLocal; _ = µobject_pairs_hook
				var µ_scan_once *πg.Object = πg.UnboundLocal; _ = µ_scan_once
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []πg.Param
				_ = πTemp003
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 17: parse_object = context.parse_object
					πF.SetLineno(17)
					if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µcontext, ßparse_object, nil); πE != nil {
						continue
					}
					µparse_object = πTemp001
					// line 18: parse_array = context.parse_array
					πF.SetLineno(18)
					if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µcontext, ßparse_array, nil); πE != nil {
						continue
					}
					µparse_array = πTemp001
					// line 19: parse_string = context.parse_string
					πF.SetLineno(19)
					if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µcontext, ßparse_string, nil); πE != nil {
						continue
					}
					µparse_string = πTemp001
					// line 20: match_number = NUMBER_RE.match
					πF.SetLineno(20)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNUMBER_RE); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmatch, nil); πE != nil {
						continue
					}
					µmatch_number = πTemp002
					// line 21: encoding = context.encoding
					πF.SetLineno(21)
					if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µcontext, ßencoding, nil); πE != nil {
						continue
					}
					µencoding = πTemp001
					// line 22: strict = context.strict
					πF.SetLineno(22)
					if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µcontext, ßstrict, nil); πE != nil {
						continue
					}
					µstrict = πTemp001
					// line 23: parse_float = context.parse_float
					πF.SetLineno(23)
					if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µcontext, ßparse_float, nil); πE != nil {
						continue
					}
					µparse_float = πTemp001
					// line 24: parse_int = context.parse_int
					πF.SetLineno(24)
					if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µcontext, ßparse_int, nil); πE != nil {
						continue
					}
					µparse_int = πTemp001
					// line 25: parse_constant = context.parse_constant
					πF.SetLineno(25)
					if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µcontext, ßparse_constant, nil); πE != nil {
						continue
					}
					µparse_constant = πTemp001
					// line 26: object_hook = context.object_hook
					πF.SetLineno(26)
					if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µcontext, ßobject_hook, nil); πE != nil {
						continue
					}
					µobject_hook = πTemp001
					// line 27: object_pairs_hook = context.object_pairs_hook
					πF.SetLineno(27)
					if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µcontext, ßobject_pairs_hook, nil); πE != nil {
						continue
					}
					µobject_pairs_hook = πTemp001
					// line 29: def _scan_once(string, idx):
					πF.SetLineno(29)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "string", Def: nil}
					πTemp003[1] = πg.Param{Name: "idx", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("_scan_once", "build/src/__python__/json_scanner.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µstring *πg.Object = πArgs[0]; _ = µstring
						var µidx *πg.Object = πArgs[1]; _ = µidx
						var µnextchar *πg.Object = πg.UnboundLocal; _ = µnextchar
						var µm *πg.Object = πg.UnboundLocal; _ = µm
						var µinteger *πg.Object = πg.UnboundLocal; _ = µinteger
						var µfrac *πg.Object = πg.UnboundLocal; _ = µfrac
						var µexp *πg.Object = πg.UnboundLocal; _ = µexp
						var µres *πg.Object = πg.UnboundLocal; _ = µres
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.BaseException
						_ = πTemp003
						var πTemp004 *πg.Traceback
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 30: try:
							πF.SetLineno(30)
							πF.PushCheckpoint(2)
							// line 31: nextchar = string[idx]
							πF.SetLineno(31)
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							πTemp001 = µidx
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µstring, πTemp001); πE != nil {
								continue
							}
							µnextchar = πTemp002
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp003, πTemp004 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsInstance(πF, πTemp003.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							πE = πF.Raise(πTemp003.ToObject(), nil, πTemp004.ToObject())
							continue
							// line 32: except IndexError:
							πF.SetLineno(32)
						Label3:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
								continue
							}
							// line 33: raise StopIteration
							πF.SetLineno(33)
							πE = πF.Raise(πTemp001, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							if πE = πg.CheckLocal(πF, µnextchar, "nextchar"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µnextchar, πg.NewStr("\"").ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µnextchar, "nextchar"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µnextchar, πg.NewStr("{").ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							if πE = πg.CheckLocal(πF, µnextchar, "nextchar"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µnextchar, πg.NewStr("[").ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µnextchar, "nextchar"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µnextchar, ßn.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Add(πF, µidx, πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{µidx, πTemp007, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µstring, πTemp006); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp007, ßnull.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label7:
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label8
							}
							if πE = πg.CheckLocal(πF, µnextchar, "nextchar"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µnextchar, ßt.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label9
							}
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Add(πF, µidx, πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{µidx, πTemp007, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µstring, πTemp006); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp007, ßtrue.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label9:
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label10
							}
							if πE = πg.CheckLocal(πF, µnextchar, "nextchar"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µnextchar, ßf.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label11
							}
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Add(πF, µidx, πg.NewInt(5).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{µidx, πTemp007, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µstring, πTemp006); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp007, ßfalse.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label11:
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label12
							}
							goto Label13
							// line 35: if nextchar == '"':
							πF.SetLineno(35)
						Label4:
							// line 36: return parse_string(string, idx + 1, encoding, strict)
							πF.SetLineno(36)
							πTemp008 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							πTemp008[0] = µstring
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µidx, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp008[1] = πTemp001
							if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
								continue
							}
							πTemp008[2] = µencoding
							if πE = πg.CheckLocal(πF, µstrict, "strict"); πE != nil {
								continue
							}
							πTemp008[3] = µstrict
							if πE = πg.CheckLocal(πF, µparse_string, "parse_string"); πE != nil {
								continue
							}
							if πTemp001, πE = µparse_string.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πR = πTemp001
							continue
							goto Label13
							// line 37: elif nextchar == '{':
							πF.SetLineno(37)
						Label5:
							// line 38: return parse_object((string, idx + 1), encoding, strict,
							πF.SetLineno(38)
							πTemp008 = πF.MakeArgs(6)
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µidx, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µstring, πTemp002).ToObject()
							πTemp008[0] = πTemp001
							if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
								continue
							}
							πTemp008[1] = µencoding
							if πE = πg.CheckLocal(πF, µstrict, "strict"); πE != nil {
								continue
							}
							πTemp008[2] = µstrict
							if πE = πg.CheckLocal(πF, µ_scan_once, "_scan_once"); πE != nil {
								continue
							}
							πTemp008[3] = µ_scan_once
							if πE = πg.CheckLocal(πF, µobject_hook, "object_hook"); πE != nil {
								continue
							}
							πTemp008[4] = µobject_hook
							if πE = πg.CheckLocal(πF, µobject_pairs_hook, "object_pairs_hook"); πE != nil {
								continue
							}
							πTemp008[5] = µobject_pairs_hook
							if πE = πg.CheckLocal(πF, µparse_object, "parse_object"); πE != nil {
								continue
							}
							if πTemp001, πE = µparse_object.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πR = πTemp001
							continue
							goto Label13
							// line 40: elif nextchar == '[':
							πF.SetLineno(40)
						Label6:
							// line 41: return parse_array((string, idx + 1), _scan_once)
							πF.SetLineno(41)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µidx, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µstring, πTemp002).ToObject()
							πTemp008[0] = πTemp001
							if πE = πg.CheckLocal(πF, µ_scan_once, "_scan_once"); πE != nil {
								continue
							}
							πTemp008[1] = µ_scan_once
							if πE = πg.CheckLocal(πF, µparse_array, "parse_array"); πE != nil {
								continue
							}
							if πTemp001, πE = µparse_array.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πR = πTemp001
							continue
							goto Label13
							// line 42: elif nextchar == 'n' and string[idx:idx + 4] == 'null':
							πF.SetLineno(42)
						Label8:
							// line 43: return None, idx + 4
							πF.SetLineno(43)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Add(πF, µidx, πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πTemp006).ToObject()
							πR = πTemp001
							continue
							goto Label13
							// line 44: elif nextchar == 't' and string[idx:idx + 4] == 'true':
							πF.SetLineno(44)
						Label10:
							// line 45: return True, idx + 4
							πF.SetLineno(45)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Add(πF, µidx, πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πTemp006).ToObject()
							πR = πTemp001
							continue
							goto Label13
							// line 46: elif nextchar == 'f' and string[idx:idx + 5] == 'false':
							πF.SetLineno(46)
						Label12:
							// line 47: return False, idx + 5
							πF.SetLineno(47)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Add(πF, µidx, πg.NewInt(5).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πTemp006).ToObject()
							πR = πTemp001
							continue
							goto Label13
						Label13:
							// line 49: m = match_number(string, idx)
							πF.SetLineno(49)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							πTemp008[0] = µstring
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							πTemp008[1] = µidx
							if πE = πg.CheckLocal(πF, µmatch_number, "match_number"); πE != nil {
								continue
							}
							if πTemp001, πE = µmatch_number.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							µm = πTemp001
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µm != πTemp002).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label14
							}
							if πE = πg.CheckLocal(πF, µnextchar, "nextchar"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µnextchar, ßN.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label15
							}
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Add(πF, µidx, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{µidx, πTemp007, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µstring, πTemp006); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp007, ßNaN.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label15:
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label16
							}
							if πE = πg.CheckLocal(πF, µnextchar, "nextchar"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µnextchar, ßI.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label17
							}
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Add(πF, µidx, πg.NewInt(8).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{µidx, πTemp007, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µstring, πTemp006); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp007, ßInfinity.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label17:
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label18
							}
							if πE = πg.CheckLocal(πF, µnextchar, "nextchar"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µnextchar, πg.NewStr("-").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label19
							}
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Add(πF, µidx, πg.NewInt(9).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{µidx, πTemp007, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µstring, πTemp006); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp007, πg.NewStr("-Infinity").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label19:
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label20
							}
							goto Label21
							// line 50: if m is not None:
							πF.SetLineno(50)
						Label14:
							// line 51: integer, frac, exp = m.groups()
							πF.SetLineno(51)
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µm, ßgroups, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp002); πE != nil {
								continue
							}
							µinteger = πTemp001
							µfrac = πTemp006
							µexp = πTemp007
							if πE = πg.CheckLocal(πF, µfrac, "frac"); πE != nil {
								continue
							}
							πTemp001 = µfrac
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label23
							}
							if πE = πg.CheckLocal(πF, µexp, "exp"); πE != nil {
								continue
							}
							πTemp001 = µexp
						Label23:
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label24
							}
							goto Label25
							// line 52: if frac or exp:
							πF.SetLineno(52)
						Label24:
							// line 53: res = parse_float(integer + (frac or '') + (exp or ''))
							πF.SetLineno(53)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µinteger, "integer"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfrac, "frac"); πE != nil {
								continue
							}
							πTemp006 = µfrac
							if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label27
							}
							πTemp006 = ß.ToObject()
						Label27:
							if πTemp002, πE = πg.Add(πF, µinteger, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µexp, "exp"); πE != nil {
								continue
							}
							πTemp006 = µexp
							if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label28
							}
							πTemp006 = ß.ToObject()
						Label28:
							if πTemp001, πE = πg.Add(πF, πTemp002, πTemp006); πE != nil {
								continue
							}
							πTemp008[0] = πTemp001
							if πE = πg.CheckLocal(πF, µparse_float, "parse_float"); πE != nil {
								continue
							}
							if πTemp001, πE = µparse_float.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							µres = πTemp001
							goto Label26
						Label25:
							// line 55: res = parse_int(integer)
							πF.SetLineno(55)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µinteger, "integer"); πE != nil {
								continue
							}
							πTemp008[0] = µinteger
							if πE = πg.CheckLocal(πF, µparse_int, "parse_int"); πE != nil {
								continue
							}
							if πTemp001, πE = µparse_int.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							µres = πTemp001
							goto Label26
						Label26:
							// line 56: return res, m.end()
							πF.SetLineno(56)
							if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µm, ßend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µres, πTemp006).ToObject()
							πR = πTemp001
							continue
							goto Label22
							// line 57: elif nextchar == 'N' and string[idx:idx + 3] == 'NaN':
							πF.SetLineno(57)
						Label16:
							// line 58: return parse_constant('NaN'), idx + 3
							πF.SetLineno(58)
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = ßNaN.ToObject()
							if πE = πg.CheckLocal(πF, µparse_constant, "parse_constant"); πE != nil {
								continue
							}
							if πTemp002, πE = µparse_constant.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Add(πF, µidx, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πTemp006).ToObject()
							πR = πTemp001
							continue
							goto Label22
							// line 59: elif nextchar == 'I' and string[idx:idx + 8] == 'Infinity':
							πF.SetLineno(59)
						Label18:
							// line 60: return parse_constant('Infinity'), idx + 8
							πF.SetLineno(60)
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = ßInfinity.ToObject()
							if πE = πg.CheckLocal(πF, µparse_constant, "parse_constant"); πE != nil {
								continue
							}
							if πTemp002, πE = µparse_constant.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Add(πF, µidx, πg.NewInt(8).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πTemp006).ToObject()
							πR = πTemp001
							continue
							goto Label22
							// line 61: elif nextchar == '-' and string[idx:idx + 9] == '-Infinity':
							πF.SetLineno(61)
						Label20:
							// line 62: return parse_constant('-Infinity'), idx + 9
							πF.SetLineno(62)
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = πg.NewStr("-Infinity").ToObject()
							if πE = πg.CheckLocal(πF, µparse_constant, "parse_constant"); πE != nil {
								continue
							}
							if πTemp002, πE = µparse_constant.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Add(πF, µidx, πg.NewInt(9).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πTemp006).ToObject()
							πR = πTemp001
							continue
							goto Label22
						Label21:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
								continue
							}
							// line 64: raise StopIteration
							πF.SetLineno(64)
							πE = πF.Raise(πTemp001, nil, nil)
							continue
							goto Label22
						Label22:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					µ_scan_once = πTemp001
					// line 66: return _scan_once
					πF.SetLineno(66)
					if πE = πg.CheckLocal(πF, µ_scan_once, "_scan_once"); πE != nil {
						continue
					}
					πR = µ_scan_once
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßpy_make_scanner.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 68: make_scanner = c_make_scanner or py_make_scanner
			πF.SetLineno(68)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßc_make_scanner); πE != nil {
				continue
			}
			πTemp003 = πTemp004
			if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
				continue
			}
			if πTemp008 {
				goto Label1
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßpy_make_scanner); πE != nil {
				continue
			}
			πTemp003 = πTemp004
		Label1:
			if πE = πF.Globals().SetItem(πF, ßmake_scanner.ToObject(), πTemp003); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("json_scanner", Code)
}
