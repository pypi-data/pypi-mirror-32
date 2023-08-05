package pprint
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/pprint.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßA := πg.InternStr("A")
		ßDeprecationWarning := πg.InternStr("DeprecationWarning")
		ßFalse := πg.InternStr("False")
		ßNone := πg.InternStr("None")
		ßPrettyPrinter := πg.InternStr("PrettyPrinter")
		ßStringIO := πg.InternStr("StringIO")
		ßTrue := πg.InternStr("True")
		ßZ := πg.InternStr("Z")
		ß_StringIO := πg.InternStr("_StringIO")
		ß__all__ := πg.InternStr("__all__")
		ß__enter__ := πg.InternStr("__enter__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__repr__ := πg.InternStr("__repr__")
		ß_commajoin := πg.InternStr("_commajoin")
		ß_depth := πg.InternStr("_depth")
		ß_format := πg.InternStr("_format")
		ß_id := πg.InternStr("_id")
		ß_indent_per_level := πg.InternStr("_indent_per_level")
		ß_len := πg.InternStr("_len")
		ß_readable := πg.InternStr("_readable")
		ß_recursion := πg.InternStr("_recursion")
		ß_recursive := πg.InternStr("_recursive")
		ß_repr := πg.InternStr("_repr")
		ß_safe_repr := πg.InternStr("_safe_repr")
		ß_sorted := πg.InternStr("_sorted")
		ß_stream := πg.InternStr("_stream")
		ß_sys := πg.InternStr("_sys")
		ß_type := πg.InternStr("_type")
		ß_width := πg.InternStr("_width")
		ßa := πg.InternStr("a")
		ßappend := πg.InternStr("append")
		ßcatch_warnings := πg.InternStr("catch_warnings")
		ßdict := πg.InternStr("dict")
		ßfilterwarnings := πg.InternStr("filterwarnings")
		ßformat := πg.InternStr("format")
		ßfrozenset := πg.InternStr("frozenset")
		ßget := πg.InternStr("get")
		ßgetattr := πg.InternStr("getattr")
		ßgetvalue := πg.InternStr("getvalue")
		ßid := πg.InternStr("id")
		ßignore := πg.InternStr("ignore")
		ßint := πg.InternStr("int")
		ßisreadable := πg.InternStr("isreadable")
		ßisrecursive := πg.InternStr("isrecursive")
		ßissubclass := πg.InternStr("issubclass")
		ßitems := πg.InternStr("items")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßlist := πg.InternStr("list")
		ßobject := πg.InternStr("object")
		ßpformat := πg.InternStr("pformat")
		ßpprint := πg.InternStr("pprint")
		ßpy3kwarning := πg.InternStr("py3kwarning")
		ßrepr := πg.InternStr("repr")
		ßsaferepr := πg.InternStr("saferepr")
		ßset := πg.InternStr("set")
		ßsorted := πg.InternStr("sorted")
		ßstartswith := πg.InternStr("startswith")
		ßstdout := πg.InternStr("stdout")
		ßstr := πg.InternStr("str")
		ßtuple := πg.InternStr("tuple")
		ßtype := πg.InternStr("type")
		ßwarnings := πg.InternStr("warnings")
		ßwrite := πg.InternStr("write")
		ßz := πg.InternStr("z")
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
		var πTemp009 *πg.Dict
		_ = πTemp009
		var πTemp010 *πg.Object
		_ = πTemp010
		var πTemp011 *πg.Object
		_ = πTemp011
		var πTemp012 *πg.Object
		_ = πTemp012
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 11: """Support to pretty-print lists, tuples, & dictionaries recursively.
			πF.SetLineno(11)
			// line 37: import sys as _sys
			πF.SetLineno(37)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_sys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 38: import warnings
			πF.SetLineno(38)
			if πTemp002, πE = πg.ImportModule(πF, "warnings"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßwarnings.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 39: import StringIO
			πF.SetLineno(39)
			if πTemp002, πE = πg.ImportModule(πF, "StringIO"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßStringIO.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 40: _StringIO = StringIO.StringIO
			πF.SetLineno(40)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßStringIO, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_StringIO.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 47: __all__ = ["pprint","pformat","isreadable","isrecursive","saferepr",
			πF.SetLineno(47)
			πTemp002 = make([]*πg.Object, 6)
			πTemp002[0] = ßpprint.ToObject()
			πTemp002[1] = ßpformat.ToObject()
			πTemp002[2] = ßisreadable.ToObject()
			πTemp002[3] = ßisrecursive.ToObject()
			πTemp002[4] = ßsaferepr.ToObject()
			πTemp002[5] = ßPrettyPrinter.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 51: _commajoin = ", ".join
			πF.SetLineno(51)
			if πTemp001, πE = πg.GetAttr(πF, πg.NewStr(", ").ToObject(), ßjoin, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_commajoin.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 52: _id = id
			πF.SetLineno(52)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_id.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 53: _len = len
			πF.SetLineno(53)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_len.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 54: _type = type
			πF.SetLineno(54)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_type.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 57: def pprint(o, stream=None, indent=1, width=80, depth=None):
			πF.SetLineno(57)
			πTemp004 = make([]πg.Param, 5)
			πTemp004[0] = πg.Param{Name: "o", Def: nil}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "stream", Def: πTemp003}
			πTemp004[2] = πg.Param{Name: "indent", Def: πg.NewInt(1).ToObject()}
			πTemp004[3] = πg.Param{Name: "width", Def: πg.NewInt(80).ToObject()}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[4] = πg.Param{Name: "depth", Def: πTemp003}
			πTemp001 = πg.NewFunction(πg.NewCode("pprint", "build/src/__python__/pprint.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µo *πg.Object = πArgs[0]; _ = µo
				var µstream *πg.Object = πArgs[1]; _ = µstream
				var µindent *πg.Object = πArgs[2]; _ = µindent
				var µwidth *πg.Object = πArgs[3]; _ = µwidth
				var µdepth *πg.Object = πArgs[4]; _ = µdepth
				var µprinter *πg.Object = πg.UnboundLocal; _ = µprinter
				var πTemp001 πg.KWArgs
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
					// line 58: """Pretty-print a Python o to a stream [default is sys.stdout]."""
					πF.SetLineno(58)
					// line 59: printer = PrettyPrinter(
					πF.SetLineno(59)
					if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
						continue
					}
					πTemp001 = πg.KWArgs{
						{"stream", µstream},
						{"indent", µindent},
						{"width", µwidth},
						{"depth", µdepth},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßPrettyPrinter); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, πTemp001); πE != nil {
						continue
					}
					µprinter = πTemp003
					// line 61: printer.pprint(o)
					πF.SetLineno(61)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					πTemp004[0] = µo
					if πE = πg.CheckLocal(πF, µprinter, "printer"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µprinter, ßpprint, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßpprint.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 63: def pformat(o, indent=1, width=80, depth=None):
			πF.SetLineno(63)
			πTemp004 = make([]πg.Param, 4)
			πTemp004[0] = πg.Param{Name: "o", Def: nil}
			πTemp004[1] = πg.Param{Name: "indent", Def: πg.NewInt(1).ToObject()}
			πTemp004[2] = πg.Param{Name: "width", Def: πg.NewInt(80).ToObject()}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[3] = πg.Param{Name: "depth", Def: πTemp005}
			πTemp003 = πg.NewFunction(πg.NewCode("pformat", "build/src/__python__/pprint.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µo *πg.Object = πArgs[0]; _ = µo
				var µindent *πg.Object = πArgs[1]; _ = µindent
				var µwidth *πg.Object = πArgs[2]; _ = µwidth
				var µdepth *πg.Object = πArgs[3]; _ = µdepth
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 πg.KWArgs
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
					// line 64: """Format a Python o into a pretty-printed representation."""
					πF.SetLineno(64)
					// line 65: return PrettyPrinter(indent=indent, width=width, depth=depth).pformat(o)
					πF.SetLineno(65)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					πTemp001[0] = µo
					if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
						continue
					}
					πTemp002 = πg.KWArgs{
						{"indent", µindent},
						{"width", µwidth},
						{"depth", µdepth},
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßPrettyPrinter); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, πTemp002); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßpformat, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßpformat.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 67: def saferepr(o):
			πF.SetLineno(67)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "o", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("saferepr", "build/src/__python__/pprint.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µo *πg.Object = πArgs[0]; _ = µo
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Dict
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
					// line 68: """Version of repr() which can handle recursive data structures."""
					πF.SetLineno(68)
					// line 69: return _safe_repr(o, {}, None, 0)[0]
					πF.SetLineno(69)
					πTemp001 = πg.NewInt(0).ToObject()
					πTemp003 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					πTemp003[0] = µo
					πTemp004 = πg.NewDict()
					πTemp005 = πTemp004.ToObject()
					πTemp003[1] = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003[2] = πTemp005
					πTemp003[3] = πg.NewInt(0).ToObject()
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_safe_repr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßsaferepr.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 71: def isreadable(o):
			πF.SetLineno(71)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "o", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("isreadable", "build/src/__python__/pprint.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µo *πg.Object = πArgs[0]; _ = µo
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Dict
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
					// line 72: """Determine if saferepr(o) is readable by eval()."""
					πF.SetLineno(72)
					// line 73: return _safe_repr(o, {}, None, 0)[1]
					πF.SetLineno(73)
					πTemp001 = πg.NewInt(1).ToObject()
					πTemp003 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					πTemp003[0] = µo
					πTemp004 = πg.NewDict()
					πTemp005 = πTemp004.ToObject()
					πTemp003[1] = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003[2] = πTemp005
					πTemp003[3] = πg.NewInt(0).ToObject()
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_safe_repr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßisreadable.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 75: def isrecursive(o):
			πF.SetLineno(75)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "o", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("isrecursive", "build/src/__python__/pprint.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µo *πg.Object = πArgs[0]; _ = µo
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Dict
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
					// line 76: """Determine if o requires a recursive representation."""
					πF.SetLineno(76)
					// line 77: return _safe_repr(o, {}, None, 0)[2]
					πF.SetLineno(77)
					πTemp001 = πg.NewInt(2).ToObject()
					πTemp003 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					πTemp003[0] = µo
					πTemp004 = πg.NewDict()
					πTemp005 = πTemp004.ToObject()
					πTemp003[1] = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003[2] = πTemp005
					πTemp003[3] = πg.NewInt(0).ToObject()
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_safe_repr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßisrecursive.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 79: def _sorted(iterable):
			πF.SetLineno(79)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "iterable", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("_sorted", "build/src/__python__/pprint.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µiterable *πg.Object = πArgs[0]; _ = µiterable
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
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 *πg.BaseException
				_ = πTemp008
				var πTemp009 *πg.Traceback
				_ = πTemp009
				var πTemp010 *πg.Type
				_ = πTemp010
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					default: panic("unexpected function state")
					}
					// line 80: with warnings.catch_warnings():
					πF.SetLineno(80)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßwarnings); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßcatch_warnings, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001.Type().ToObject(), ß__exit__, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001.Type().ToObject(), ß__enter__, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp003.Call(πF, πg.Args{πTemp001}, nil); πE != nil {
						continue
					}
					πF.PushCheckpoint(1)
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_sys); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßpy3kwarning, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label2
					}
					goto Label3
					// line 81: if _sys.py3kwarning:
					πF.SetLineno(81)
				Label2:
					// line 82: warnings.filterwarnings("ignore", "comparing unequal types "
					πF.SetLineno(82)
					πTemp007 = πF.MakeArgs(3)
					πTemp007[0] = ßignore.ToObject()
					πTemp007[1] = πg.NewStr("comparing unequal types not supported").ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßDeprecationWarning); πE != nil {
						continue
					}
					πTemp007[2] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßwarnings); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßfilterwarnings, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					goto Label3
				Label3:
					// line 84: return sorted(iterable)
					πF.SetLineno(84)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
						continue
					}
					πTemp007[0] = µiterable
					if πTemp004, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πR = πTemp005
					continue
					πF.PopCheckpoint()
				Label1:
					πTemp008, πTemp009 = nil, nil
					if πE != nil {
						πTemp008, πTemp009 = πF.ExcInfo()
					}
					if πTemp008 != nil {
						πTemp010 = πTemp008.Type()
						if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πTemp010.ToObject(), πTemp008.ToObject(), πTemp009.ToObject()}, nil); πE != nil {
							continue
						}
					} else {
						if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πg.None, πg.None, πg.None}, nil); πE != nil {
							continue
						}
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp008 != nil && πTemp006 != true {
						πE = πF.Raise(nil, nil, nil)
						continue
					}
					if πR != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_sorted.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 86: class PrettyPrinter(object):
			πF.SetLineno(86)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp012, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp012
			πTemp009 = πg.NewDict()
			if πTemp010, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp010); πE != nil {
				continue
			}
			_, πE = πg.NewCode("PrettyPrinter", "build/src/__python__/pprint.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 87: def __init__(self, indent=1, width=80, depth=None, stream=None):
					πF.SetLineno(87)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "indent", Def: πg.NewInt(1).ToObject()}
					πTemp002[2] = πg.Param{Name: "width", Def: πg.NewInt(80).ToObject()}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "depth", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "stream", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/pprint.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µindent *πg.Object = πArgs[1]; _ = µindent
						var µwidth *πg.Object = πArgs[2]; _ = µwidth
						var µdepth *πg.Object = πArgs[3]; _ = µdepth
						var µstream *πg.Object = πArgs[4]; _ = µstream
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 88: """Handle pretty printing operations onto a stream using a set of
							πF.SetLineno(88)
							// line 105: indent = int(indent)
							πF.SetLineno(105)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							πTemp001[0] = µindent
							if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µindent = πTemp003
							// line 106: width = int(width)
							πF.SetLineno(106)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
								continue
							}
							πTemp001[0] = µwidth
							if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µwidth = πTemp003
							// line 107: assert indent >= 0, "indent must be >= 0"
							πF.SetLineno(107)
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GE(πF, µindent, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp002, πg.NewStr("indent must be >= 0").ToObject()); πE != nil {
								continue
							}
							// line 108: assert depth is None or depth > 0, "depth must be > 0"
							πF.SetLineno(108)
							if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µdepth == πTemp005).ToObject()
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, µdepth, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
						Label1:
							if πE = πg.Assert(πF, πTemp002, πg.NewStr("depth must be > 0").ToObject()); πE != nil {
								continue
							}
							// line 109: assert width, "width must be != 0"
							πF.SetLineno(109)
							if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, µwidth, πg.NewStr("width must be != 0").ToObject()); πE != nil {
								continue
							}
							// line 110: self._depth = depth
							πF.SetLineno(110)
							if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µdepth); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_depth, πTemp002); πE != nil {
								continue
							}
							// line 111: self._indent_per_level = indent
							πF.SetLineno(111)
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µindent); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_indent_per_level, πTemp002); πE != nil {
								continue
							}
							// line 112: self._width = width
							πF.SetLineno(112)
							if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µwidth); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_width, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µstream != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							goto Label3
							// line 113: if stream is not None:
							πF.SetLineno(113)
						Label2:
							// line 114: self._stream = stream
							πF.SetLineno(114)
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µstream); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_stream, πTemp002); πE != nil {
								continue
							}
							goto Label4
						Label3:
							// line 116: self._stream = _sys.stdout
							πF.SetLineno(116)
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_sys); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstdout, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_stream, πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 118: def pprint(self, o):
					πF.SetLineno(118)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "o", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("pprint", "build/src/__python__/pprint.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µo *πg.Object = πArgs[1]; _ = µo
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Dict
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
							// line 119: self._format(o, self._stream, 0, 0, {}, 0)
							πF.SetLineno(119)
							πTemp001 = πF.MakeArgs(6)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp001[0] = µo
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_stream, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							πTemp001[2] = πg.NewInt(0).ToObject()
							πTemp001[3] = πg.NewInt(0).ToObject()
							πTemp003 = πg.NewDict()
							πTemp002 = πTemp003.ToObject()
							πTemp001[4] = πTemp002
							πTemp001[5] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_format, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 120: self._stream.write("\n")
							πF.SetLineno(120)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_stream, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßpprint.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 122: def pformat(self, o):
					πF.SetLineno(122)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "o", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("pformat", "build/src/__python__/pprint.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µo *πg.Object = πArgs[1]; _ = µo
						var µsio *πg.Object = πg.UnboundLocal; _ = µsio
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Dict
						_ = πTemp004
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 123: sio = _StringIO()
							πF.SetLineno(123)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_StringIO); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µsio = πTemp002
							// line 124: self._format(o, sio, 0, 0, {}, 0)
							πF.SetLineno(124)
							πTemp003 = πF.MakeArgs(6)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp003[0] = µo
							if πE = πg.CheckLocal(πF, µsio, "sio"); πE != nil {
								continue
							}
							πTemp003[1] = µsio
							πTemp003[2] = πg.NewInt(0).ToObject()
							πTemp003[3] = πg.NewInt(0).ToObject()
							πTemp004 = πg.NewDict()
							πTemp001 = πTemp004.ToObject()
							πTemp003[4] = πTemp001
							πTemp003[5] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_format, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 125: return sio.getvalue()
							πF.SetLineno(125)
							if πE = πg.CheckLocal(πF, µsio, "sio"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsio, ßgetvalue, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßpformat.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 127: def isrecursive(self, o):
					πF.SetLineno(127)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "o", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("isrecursive", "build/src/__python__/pprint.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µo *πg.Object = πArgs[1]; _ = µo
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Dict
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
							// line 128: return self.format(o, {}, 0, 0)[2]
							πF.SetLineno(128)
							πTemp001 = πg.NewInt(2).ToObject()
							πTemp003 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp003[0] = µo
							πTemp004 = πg.NewDict()
							πTemp005 = πTemp004.ToObject()
							πTemp003[1] = πTemp005
							πTemp003[2] = πg.NewInt(0).ToObject()
							πTemp003[3] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßformat, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßisrecursive.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 130: def isreadable(self, o):
					πF.SetLineno(130)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "o", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("isreadable", "build/src/__python__/pprint.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µo *πg.Object = πArgs[1]; _ = µo
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µreadable *πg.Object = πg.UnboundLocal; _ = µreadable
						var µrecursive *πg.Object = πg.UnboundLocal; _ = µrecursive
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Dict
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
						var πTemp008 bool
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 131: s, readable, recursive = self.format(o, {}, 0, 0)
							πF.SetLineno(131)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp001[0] = µo
							πTemp002 = πg.NewDict()
							πTemp003 = πTemp002.ToObject()
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewInt(0).ToObject()
							πTemp001[3] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßformat, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp004); πE != nil {
								continue
							}
							µs = πTemp003
							µreadable = πTemp005
							µrecursive = πTemp006
							// line 132: return readable and not recursive
							πF.SetLineno(132)
							if πE = πg.CheckLocal(πF, µreadable, "readable"); πE != nil {
								continue
							}
							πTemp003 = µreadable
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp007 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µrecursive, "recursive"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, µrecursive); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp008).ToObject()
							πTemp003 = πTemp004
						Label1:
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
					if πE = πClass.SetItem(πF, ßisreadable.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 134: def _format(self, o, stream, indent, allowance, context, level):
					πF.SetLineno(134)
					πTemp002 = make([]πg.Param, 7)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "o", Def: nil}
					πTemp002[2] = πg.Param{Name: "stream", Def: nil}
					πTemp002[3] = πg.Param{Name: "indent", Def: nil}
					πTemp002[4] = πg.Param{Name: "allowance", Def: nil}
					πTemp002[5] = πg.Param{Name: "context", Def: nil}
					πTemp002[6] = πg.Param{Name: "level", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("_format", "build/src/__python__/pprint.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µo *πg.Object = πArgs[1]; _ = µo
						var µstream *πg.Object = πArgs[2]; _ = µstream
						var µindent *πg.Object = πArgs[3]; _ = µindent
						var µallowance *πg.Object = πArgs[4]; _ = µallowance
						var µcontext *πg.Object = πArgs[5]; _ = µcontext
						var µlevel *πg.Object = πArgs[6]; _ = µlevel
						var µobjid *πg.Object = πg.UnboundLocal; _ = µobjid
						var µrep *πg.Object = πg.UnboundLocal; _ = µrep
						var µtyp *πg.Object = πg.UnboundLocal; _ = µtyp
						var µsepLines *πg.Object = πg.UnboundLocal; _ = µsepLines
						var µwrite *πg.Object = πg.UnboundLocal; _ = µwrite
						var µr *πg.Object = πg.UnboundLocal; _ = µr
						var µlength *πg.Object = πg.UnboundLocal; _ = µlength
						var µitems *πg.Object = πg.UnboundLocal; _ = µitems
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µent *πg.Object = πg.UnboundLocal; _ = µent
						var µendchar *πg.Object = πg.UnboundLocal; _ = µendchar
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
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
						var πTemp010 bool
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 36: goto Label36
							case 37: goto Label37
							case 14: goto Label14
							case 15: goto Label15
							default: panic("unexpected function state")
							}
							// line 135: level = level + 1
							πF.SetLineno(135)
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µlevel, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µlevel = πTemp001
							// line 136: objid = _id(o)
							πF.SetLineno(136)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp002[0] = µo
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_id); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µobjid = πTemp003
							if πE = πg.CheckLocal(πF, µobjid, "objid"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, µcontext, µobjid); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 137: if objid in context:
							πF.SetLineno(137)
						Label1:
							// line 138: stream.write(_recursion(o))
							πF.SetLineno(138)
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp005[0] = µo
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_recursion); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstream, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 139: self._recursive = True
							πF.SetLineno(139)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_recursive, πTemp003); πE != nil {
								continue
							}
							// line 140: self._readable = False
							πF.SetLineno(140)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_readable, πTemp003); πE != nil {
								continue
							}
							// line 141: return
							πF.SetLineno(141)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 142: rep = self._repr(o, context, level - 1)
							πF.SetLineno(142)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp002[0] = µo
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πTemp002[1] = µcontext
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µlevel, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002[2] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_repr, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µrep = πTemp003
							// line 143: typ = _type(o)
							πF.SetLineno(143)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp002[0] = µo
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_type); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µtyp = πTemp003
							// line 144: sepLines = _len(rep) > (self._width - 1 - indent - allowance)
							πF.SetLineno(144)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrep, "rep"); πE != nil {
								continue
							}
							πTemp002[0] = µrep
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_len); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ß_width, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Sub(πF, πTemp009, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Sub(πF, πTemp008, µindent); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µallowance, "allowance"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, πTemp007, µallowance); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GT(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							µsepLines = πTemp001
							// line 145: write = stream.write
							πF.SetLineno(145)
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstream, ßwrite, nil); πE != nil {
								continue
							}
							µwrite = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_depth, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß_depth, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, µlevel, πTemp006); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label3:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 147: if self._depth and level > self._depth:
							πF.SetLineno(147)
						Label4:
							// line 148: write(rep)
							πF.SetLineno(148)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrep, "rep"); πE != nil {
								continue
							}
							πTemp002[0] = µrep
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 149: return
							πF.SetLineno(149)
							πR = πg.None
							continue
							goto Label5
						Label5:
							// line 151: r = getattr(typ, "__repr__", None)
							πF.SetLineno(151)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µtyp, "typ"); πE != nil {
								continue
							}
							πTemp002[0] = µtyp
							πTemp002[1] = ß__repr__.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µr = πTemp003
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtyp, "typ"); πE != nil {
								continue
							}
							πTemp002[0] = µtyp
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							goto Label7
							// line 152: if issubclass(typ, dict):# and r == dict.__repr__:
							πF.SetLineno(152)
						Label6:
							// line 153: write('{')
							πF.SetLineno(153)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("{").ToObject()
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_indent_per_level, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GT(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							goto Label9
							// line 154: if self._indent_per_level > 1:
							πF.SetLineno(154)
						Label8:
							// line 155: write((self._indent_per_level - 1) * ' ')
							πF.SetLineno(155)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß_indent_per_level, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, πTemp003, πg.NewStr(" ").ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label9
						Label9:
							// line 156: length = _len(o)
							πF.SetLineno(156)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp002[0] = µo
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_len); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µlength = πTemp003
							if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µlength); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label10
							}
							goto Label11
							// line 157: if length:
							πF.SetLineno(157)
						Label10:
							// line 158: context[objid] = 1
							πF.SetLineno(158)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobjid, "objid"); πE != nil {
								continue
							}
							πTemp003 = µobjid
							if πE = πg.SetItem(πF, µcontext, πTemp003, πTemp001); πE != nil {
								continue
							}
							// line 159: indent = indent + self._indent_per_level
							πF.SetLineno(159)
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_indent_per_level, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µindent, πTemp003); πE != nil {
								continue
							}
							µindent = πTemp001
							// line 160: items = _sorted(o.items())
							πF.SetLineno(160)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µo, ßitems, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_sorted); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µitems = πTemp003
							// line 161: key, ent = items[0]
							πF.SetLineno(161)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µitems, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp006}}}, πTemp003); πE != nil {
								continue
							}
							µkey = πTemp001
							µent = πTemp006
							// line 162: rep = self._repr(key, context, level)
							πF.SetLineno(162)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp002[0] = µkey
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πTemp002[1] = µcontext
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							πTemp002[2] = µlevel
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_repr, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µrep = πTemp003
							// line 163: write(rep)
							πF.SetLineno(163)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrep, "rep"); πE != nil {
								continue
							}
							πTemp002[0] = µrep
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 164: write(': ')
							πF.SetLineno(164)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr(": ").ToObject()
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 165: self._format(ent, stream, indent + _len(rep) + 2,
							πF.SetLineno(165)
							πTemp002 = πF.MakeArgs(6)
							if πE = πg.CheckLocal(πF, µent, "ent"); πE != nil {
								continue
							}
							πTemp002[0] = µent
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							πTemp002[1] = µstream
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrep, "rep"); πE != nil {
								continue
							}
							πTemp005[0] = µrep
							if πTemp006, πE = πg.ResolveGlobal(πF, ß_len); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp003, πE = πg.Add(πF, µindent, πTemp007); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002[2] = πTemp001
							if πE = πg.CheckLocal(πF, µallowance, "allowance"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µallowance, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002[3] = πTemp001
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πTemp002[4] = µcontext
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							πTemp002[5] = µlevel
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_format, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GT(πF, µlength, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label12
							}
							goto Label13
							// line 167: if length > 1:
							πF.SetLineno(167)
						Label12:
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µitems, πTemp003); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp006); πE != nil {
								continue
							}
							πF.PushCheckpoint(15)
							πTemp004 = false
						Label14:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label16
							}
							if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
									continue
								}
								µkey = πTemp006
								µent = πTemp007
							}
							if πE != nil || !πTemp010 {
								continue
							}
							πF.PushCheckpoint(14)            
							// line 169: rep = self._repr(key, context, level)
							πF.SetLineno(169)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp002[0] = µkey
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πTemp002[1] = µcontext
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							πTemp002[2] = µlevel
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_repr, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µrep = πTemp006
							if πE = πg.CheckLocal(πF, µsepLines, "sepLines"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, µsepLines); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label17
							}
							goto Label18
							// line 170: if sepLines:
							πF.SetLineno(170)
						Label17:
							// line 171: write(',\n%s%s: ' % (' '*indent, rep))
							πF.SetLineno(171)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Mul(πF, πg.NewStr(" ").ToObject(), µindent); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrep, "rep"); πE != nil {
								continue
							}
							πTemp006 = πg.NewTuple2(πTemp007, µrep).ToObject()
							if πTemp003, πE = πg.Mod(πF, πg.NewStr(",\n%s%s: ").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp003, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label19
						Label18:
							// line 173: write(', %s: ' % rep)
							πF.SetLineno(173)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrep, "rep"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πg.NewStr(", %s: ").ToObject(), µrep); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp003, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label19
						Label19:
							// line 174: self._format(ent, stream, indent + _len(rep) + 2,
							πF.SetLineno(174)
							πTemp002 = πF.MakeArgs(6)
							if πE = πg.CheckLocal(πF, µent, "ent"); πE != nil {
								continue
							}
							πTemp002[0] = µent
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							πTemp002[1] = µstream
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrep, "rep"); πE != nil {
								continue
							}
							πTemp005[0] = µrep
							if πTemp007, πE = πg.ResolveGlobal(πF, ß_len); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp006, πE = πg.Add(πF, µindent, πTemp008); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp006, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002[2] = πTemp003
							if πE = πg.CheckLocal(πF, µallowance, "allowance"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µallowance, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002[3] = πTemp003
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πTemp002[4] = µcontext
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							πTemp002[5] = µlevel
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_format, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label15:
							if πE != nil || πR != nil {
								continue
							}
						Label16:
							goto Label13
						Label13:
							// line 176: indent = indent - self._indent_per_level
							πF.SetLineno(176)
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_indent_per_level, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µindent, πTemp003); πE != nil {
								continue
							}
							µindent = πTemp001
							// line 177: del context[objid]
							πF.SetLineno(177)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobjid, "objid"); πE != nil {
								continue
							}
							πTemp001 = µobjid
							if πE = πg.DelItem(πF, µcontext, πTemp001); πE != nil {
								continue
							}
							goto Label11
						Label11:
							// line 178: write('}')
							πF.SetLineno(178)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("}").ToObject()
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 179: return
							πF.SetLineno(179)
							πR = πg.None
							continue
							goto Label7
						Label7:
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtyp, "typ"); πE != nil {
								continue
							}
							πTemp002[0] = µtyp
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πTemp006
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label20
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtyp, "typ"); πE != nil {
								continue
							}
							πTemp002[0] = µtyp
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πTemp006
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label20
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtyp, "typ"); πE != nil {
								continue
							}
							πTemp002[0] = µtyp
							if πTemp003, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πTemp006
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label20
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtyp, "typ"); πE != nil {
								continue
							}
							πTemp002[0] = µtyp
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfrozenset); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πTemp006
						Label20:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label21
							}
							goto Label22
							// line 185: if ((issubclass(typ, list)) or
							πF.SetLineno(185)
						Label21:
							// line 190: length = _len(o)
							πF.SetLineno(190)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp002[0] = µo
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_len); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µlength = πTemp003
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtyp, "typ"); πE != nil {
								continue
							}
							πTemp002[0] = µtyp
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label23
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtyp, "typ"); πE != nil {
								continue
							}
							πTemp002[0] = µtyp
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label24
							}
							goto Label25
							// line 191: if issubclass(typ, list):
							πF.SetLineno(191)
						Label23:
							// line 192: write('[')
							πF.SetLineno(192)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("[").ToObject()
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 193: endchar = ']'
							πF.SetLineno(193)
							µendchar = πg.NewStr("]").ToObject()
							goto Label26
							// line 194: elif issubclass(typ, tuple):
							πF.SetLineno(194)
						Label24:
							// line 195: write('(')
							πF.SetLineno(195)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("(").ToObject()
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 196: endchar = ')'
							πF.SetLineno(196)
							µendchar = πg.NewStr(")").ToObject()
							goto Label26
						Label25:
							if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µlength); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label27
							}
							goto Label28
							// line 198: if not length:
							πF.SetLineno(198)
						Label27:
							// line 199: write(rep)
							πF.SetLineno(199)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrep, "rep"); πE != nil {
								continue
							}
							πTemp002[0] = µrep
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 200: return
							πF.SetLineno(200)
							πR = πg.None
							continue
							goto Label28
						Label28:
							// line 201: write(typ.__name__)
							πF.SetLineno(201)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtyp, "typ"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtyp, ß__name__, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 202: write('([')
							πF.SetLineno(202)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("([").ToObject()
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 203: endchar = '])'
							πF.SetLineno(203)
							µendchar = πg.NewStr("])").ToObject()
							// line 204: indent += len(typ.__name__) + 1
							πF.SetLineno(204)
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtyp, "typ"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µtyp, ß__name__, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Add(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, µindent, πTemp001); πE != nil {
								continue
							}
							µindent = πTemp003
							// line 205: o = _sorted(o)
							πF.SetLineno(205)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp002[0] = µo
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_sorted); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µo = πTemp003
							goto Label26
						Label26:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß_indent_per_level, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label29
							}
							if πE = πg.CheckLocal(πF, µsepLines, "sepLines"); πE != nil {
								continue
							}
							πTemp001 = µsepLines
						Label29:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label30
							}
							goto Label31
							// line 206: if self._indent_per_level > 1 and sepLines:
							πF.SetLineno(206)
						Label30:
							// line 207: write((self._indent_per_level - 1) * ' ')
							πF.SetLineno(207)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß_indent_per_level, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, πTemp003, πg.NewStr(" ").ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label31
						Label31:
							if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µlength); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label32
							}
							goto Label33
							// line 208: if length:
							πF.SetLineno(208)
						Label32:
							// line 209: context[objid] = 1
							πF.SetLineno(209)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobjid, "objid"); πE != nil {
								continue
							}
							πTemp003 = µobjid
							if πE = πg.SetItem(πF, µcontext, πTemp003, πTemp001); πE != nil {
								continue
							}
							// line 210: indent = indent + self._indent_per_level
							πF.SetLineno(210)
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_indent_per_level, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µindent, πTemp003); πE != nil {
								continue
							}
							µindent = πTemp001
							// line 211: self._format(o[0], stream, indent, allowance + 1,
							πF.SetLineno(211)
							πTemp002 = πF.MakeArgs(6)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µo, πTemp001); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							πTemp002[1] = µstream
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							πTemp002[2] = µindent
							if πE = πg.CheckLocal(πF, µallowance, "allowance"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µallowance, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002[3] = πTemp001
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πTemp002[4] = µcontext
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							πTemp002[5] = µlevel
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_format, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GT(πF, µlength, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label34
							}
							goto Label35
							// line 213: if length > 1:
							πF.SetLineno(213)
						Label34:
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µo, πTemp003); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp006); πE != nil {
								continue
							}
							πF.PushCheckpoint(37)
							πTemp004 = false
						Label36:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label38
							}
							if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µent = πTemp003
							}
							if πE != nil || !πTemp010 {
								continue
							}
							πF.PushCheckpoint(36)            
							if πE = πg.CheckLocal(πF, µsepLines, "sepLines"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, µsepLines); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label39
							}
							goto Label40
							// line 215: if sepLines:
							πF.SetLineno(215)
						Label39:
							// line 216: write(',\n' + ' '*indent)
							πF.SetLineno(216)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Mul(πF, πg.NewStr(" ").ToObject(), µindent); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πg.NewStr(",\n").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp003, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label41
						Label40:
							// line 218: write(', ')
							πF.SetLineno(218)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr(", ").ToObject()
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp003, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label41
						Label41:
							// line 219: self._format(ent, stream, indent,
							πF.SetLineno(219)
							πTemp002 = πF.MakeArgs(6)
							if πE = πg.CheckLocal(πF, µent, "ent"); πE != nil {
								continue
							}
							πTemp002[0] = µent
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							πTemp002[1] = µstream
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							πTemp002[2] = µindent
							if πE = πg.CheckLocal(πF, µallowance, "allowance"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µallowance, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002[3] = πTemp003
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πTemp002[4] = µcontext
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							πTemp002[5] = µlevel
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_format, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label37:
							if πE != nil || πR != nil {
								continue
							}
						Label38:
							goto Label35
						Label35:
							// line 221: indent = indent - self._indent_per_level
							πF.SetLineno(221)
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_indent_per_level, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µindent, πTemp003); πE != nil {
								continue
							}
							µindent = πTemp001
							// line 222: del context[objid]
							πF.SetLineno(222)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobjid, "objid"); πE != nil {
								continue
							}
							πTemp001 = µobjid
							if πE = πg.DelItem(πF, µcontext, πTemp001); πE != nil {
								continue
							}
							goto Label33
						Label33:
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtyp, "typ"); πE != nil {
								continue
							}
							πTemp002[0] = µtyp
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πTemp006
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label42
							}
							if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, µlength, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label42:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label43
							}
							goto Label44
							// line 223: if issubclass(typ, tuple) and length == 1:
							πF.SetLineno(223)
						Label43:
							// line 224: write(',')
							πF.SetLineno(224)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr(",").ToObject()
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label44
						Label44:
							// line 225: write(endchar)
							πF.SetLineno(225)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µendchar, "endchar"); πE != nil {
								continue
							}
							πTemp002[0] = µendchar
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 226: return
							πF.SetLineno(226)
							πR = πg.None
							continue
							goto Label22
						Label22:
							// line 228: write(rep)
							πF.SetLineno(228)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrep, "rep"); πE != nil {
								continue
							}
							πTemp002[0] = µrep
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_format.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 230: def _repr(self, o, context, level):
					πF.SetLineno(230)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "o", Def: nil}
					πTemp002[2] = πg.Param{Name: "context", Def: nil}
					πTemp002[3] = πg.Param{Name: "level", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("_repr", "build/src/__python__/pprint.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µo *πg.Object = πArgs[1]; _ = µo
						var µcontext *πg.Object = πArgs[2]; _ = µcontext
						var µlevel *πg.Object = πArgs[3]; _ = µlevel
						var µrepr1 *πg.Object = πg.UnboundLocal; _ = µrepr1
						var µreadable *πg.Object = πg.UnboundLocal; _ = µreadable
						var µrecursive *πg.Object = πg.UnboundLocal; _ = µrecursive
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 231: repr1, readable, recursive = self.format(o, dict(context),
							πF.SetLineno(231)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp001[0] = µo
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πTemp002[0] = µcontext
							if πTemp003, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_depth, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp003
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							πTemp001[3] = µlevel
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßformat, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp004); πE != nil {
								continue
							}
							µrepr1 = πTemp003
							µreadable = πTemp005
							µrecursive = πTemp006
							if πE = πg.CheckLocal(πF, µreadable, "readable"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µreadable); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label1
							}
							goto Label2
							// line 233: if not readable:
							πF.SetLineno(233)
						Label1:
							// line 234: self._readable = False
							πF.SetLineno(234)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_readable, πTemp004); πE != nil {
								continue
							}
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µrecursive, "recursive"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µrecursive); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label3
							}
							goto Label4
							// line 235: if recursive:
							πF.SetLineno(235)
						Label3:
							// line 236: self._recursive = True
							πF.SetLineno(236)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_recursive, πTemp004); πE != nil {
								continue
							}
							goto Label4
						Label4:
							// line 237: return repr1
							πF.SetLineno(237)
							if πE = πg.CheckLocal(πF, µrepr1, "repr1"); πE != nil {
								continue
							}
							πR = µrepr1
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_repr.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 239: def format(self, o, context, maxlevels, level):
					πF.SetLineno(239)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "o", Def: nil}
					πTemp002[2] = πg.Param{Name: "context", Def: nil}
					πTemp002[3] = πg.Param{Name: "maxlevels", Def: nil}
					πTemp002[4] = πg.Param{Name: "level", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("format", "build/src/__python__/pprint.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µo *πg.Object = πArgs[1]; _ = µo
						var µcontext *πg.Object = πArgs[2]; _ = µcontext
						var µmaxlevels *πg.Object = πArgs[3]; _ = µmaxlevels
						var µlevel *πg.Object = πArgs[4]; _ = µlevel
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
							// line 240: """Format o for a specific context, returning a string
							πF.SetLineno(240)
							// line 244: return _safe_repr(o, context, maxlevels, level)
							πF.SetLineno(244)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp001[0] = µo
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πTemp001[1] = µcontext
							if πE = πg.CheckLocal(πF, µmaxlevels, "maxlevels"); πE != nil {
								continue
							}
							πTemp001[2] = µmaxlevels
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							πTemp001[3] = µlevel
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_safe_repr); πE != nil {
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
					if πE = πClass.SetItem(πF, ßformat.ToObject(), πTemp009); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp011, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp011 == nil {
				πTemp011 = πg.TypeType.ToObject()
			}
			if πTemp012, πE = πTemp011.Call(πF, []*πg.Object{πg.NewStr("PrettyPrinter").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPrettyPrinter.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 249: def _safe_repr(o, context, maxlevels, level):
			πF.SetLineno(249)
			πTemp004 = make([]πg.Param, 4)
			πTemp004[0] = πg.Param{Name: "o", Def: nil}
			πTemp004[1] = πg.Param{Name: "context", Def: nil}
			πTemp004[2] = πg.Param{Name: "maxlevels", Def: nil}
			πTemp004[3] = πg.Param{Name: "level", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("_safe_repr", "build/src/__python__/pprint.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µo *πg.Object = πArgs[0]; _ = µo
				var µcontext *πg.Object = πArgs[1]; _ = µcontext
				var µmaxlevels *πg.Object = πArgs[2]; _ = µmaxlevels
				var µlevel *πg.Object = πArgs[3]; _ = µlevel
				var µtyp *πg.Object = πg.UnboundLocal; _ = µtyp
				var µclosure *πg.Object = πg.UnboundLocal; _ = µclosure
				var µquotes *πg.Object = πg.UnboundLocal; _ = µquotes
				var µqget *πg.Object = πg.UnboundLocal; _ = µqget
				var µsio *πg.Object = πg.UnboundLocal; _ = µsio
				var µwrite *πg.Object = πg.UnboundLocal; _ = µwrite
				var µchar *πg.Object = πg.UnboundLocal; _ = µchar
				var µr *πg.Object = πg.UnboundLocal; _ = µr
				var µobjid *πg.Object = πg.UnboundLocal; _ = µobjid
				var µreadable *πg.Object = πg.UnboundLocal; _ = µreadable
				var µrecursive *πg.Object = πg.UnboundLocal; _ = µrecursive
				var µcomponents *πg.Object = πg.UnboundLocal; _ = µcomponents
				var µappend *πg.Object = πg.UnboundLocal; _ = µappend
				var µsaferepr *πg.Object = πg.UnboundLocal; _ = µsaferepr
				var µk *πg.Object = πg.UnboundLocal; _ = µk
				var µv *πg.Object = πg.UnboundLocal; _ = µv
				var µkrepr *πg.Object = πg.UnboundLocal; _ = µkrepr
				var µkreadable *πg.Object = πg.UnboundLocal; _ = µkreadable
				var µkrecur *πg.Object = πg.UnboundLocal; _ = µkrecur
				var µvrepr *πg.Object = πg.UnboundLocal; _ = µvrepr
				var µvreadable *πg.Object = πg.UnboundLocal; _ = µvreadable
				var µvrecur *πg.Object = πg.UnboundLocal; _ = µvrecur
				var µformat *πg.Object = πg.UnboundLocal; _ = µformat
				var µx *πg.Object = πg.UnboundLocal; _ = µx
				var µorepr *πg.Object = πg.UnboundLocal; _ = µorepr
				var µoreadable *πg.Object = πg.UnboundLocal; _ = µoreadable
				var µorecur *πg.Object = πg.UnboundLocal; _ = µorecur
				var µrep *πg.Object = πg.UnboundLocal; _ = µrep
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
				var πTemp006 *πg.Dict
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 []*πg.Object
				_ = πTemp009
				var πTemp010 []*πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 *πg.Object
				_ = πTemp012
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 7: goto Label7
					case 8: goto Label8
					case 48: goto Label48
					case 49: goto Label49
					case 25: goto Label25
					case 26: goto Label26
					default: panic("unexpected function state")
					}
					// line 250: typ = _type(o)
					πF.SetLineno(250)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					πTemp001[0] = µo
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_type); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtyp = πTemp003
					if πE = πg.CheckLocal(πF, µtyp, "typ"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µtyp == πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 251: if typ is str:
					πF.SetLineno(251)
				Label1:
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, µo, πg.NewStr("'").ToObject()); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp005).ToObject()
					πTemp002 = πTemp003
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label3
					}
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, µo, πg.NewStr("\"").ToObject()); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp005).ToObject()
					πTemp002 = πTemp003
				Label3:
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					goto Label5
					// line 254: if "'" in o and '"' not in o:
					πF.SetLineno(254)
				Label4:
					// line 255: closure = '"'
					πF.SetLineno(255)
					µclosure = πg.NewStr("\"").ToObject()
					// line 256: quotes = {'"': '\\"'}
					πF.SetLineno(256)
					πTemp006 = πg.NewDict()
					if πE = πTemp006.SetItem(πF, πg.NewStr("\"").ToObject(), πg.NewStr("\\\"").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp006.ToObject()
					µquotes = πTemp002
					goto Label6
				Label5:
					// line 258: closure = "'"
					πF.SetLineno(258)
					µclosure = πg.NewStr("'").ToObject()
					// line 259: quotes = {"'": "\\'"}
					πF.SetLineno(259)
					πTemp006 = πg.NewDict()
					if πE = πTemp006.SetItem(πF, πg.NewStr("'").ToObject(), πg.NewStr("\\'").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp006.ToObject()
					µquotes = πTemp002
					goto Label6
				Label6:
					// line 260: qget = quotes.get
					πF.SetLineno(260)
					if πE = πg.CheckLocal(πF, µquotes, "quotes"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µquotes, ßget, nil); πE != nil {
						continue
					}
					µqget = πTemp002
					// line 261: sio = _StringIO()
					πF.SetLineno(261)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_StringIO); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µsio = πTemp003
					// line 262: write = sio.write
					πF.SetLineno(262)
					if πE = πg.CheckLocal(πF, µsio, "sio"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsio, ßwrite, nil); πE != nil {
						continue
					}
					µwrite = πTemp002
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µo); πE != nil {
						continue
					}
					πF.PushCheckpoint(8)
					πTemp004 = false
				Label7:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label9
					}
					if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
						µchar = πTemp003
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(7)            
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.LE(πF, ßa.ToObject(), µchar); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label11
					}
					if πTemp007, πE = πg.LE(πF, µchar, ßz.ToObject()); πE != nil {
						continue
					}
				Label11:
					πTemp003 = πTemp007
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label10
					}
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.LE(πF, ßA.ToObject(), µchar); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label12
					}
					if πTemp007, πE = πg.LE(πF, µchar, ßZ.ToObject()); πE != nil {
						continue
					}
				Label12:
					πTemp003 = πTemp007
				Label10:
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label13
					}
					goto Label14
					// line 265: if 'a' <= char <= 'z' or 'A' <= char <= 'Z':
					πF.SetLineno(265)
				Label13:
					// line 266: write(char)
					πF.SetLineno(266)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					πTemp001[0] = µchar
					if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
						continue
					}
					if πTemp003, πE = µwrite.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label15
				Label14:
					// line 268: write(qget(char, repr(char)[1:-1]))
					πF.SetLineno(268)
					πTemp001 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					πTemp009[0] = µchar
					if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πTemp007, πg.None}, nil); πE != nil {
						continue
					}
					πTemp010 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					πTemp010[0] = µchar
					if πTemp011, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp011.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp007, πE = πg.GetItem(πF, πTemp012, πTemp003); πE != nil {
						continue
					}
					πTemp009[1] = πTemp007
					if πE = πg.CheckLocal(πF, µqget, "qget"); πE != nil {
						continue
					}
					if πTemp003, πE = µqget.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
						continue
					}
					if πTemp003, πE = µwrite.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label15
				Label15:
					continue
				Label8:
					if πE != nil || πR != nil {
						continue
					}
				Label9:
					// line 269: return ("%s%s%s" % (closure, sio.getvalue(), closure)), True, False
					πF.SetLineno(269)
					if πE = πg.CheckLocal(πF, µclosure, "closure"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsio, "sio"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, µsio, ßgetvalue, nil); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp011.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µclosure, "closure"); πE != nil {
						continue
					}
					πTemp007 = πg.NewTuple3(µclosure, πTemp012, µclosure).ToObject()
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s%s%s").ToObject(), πTemp007); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(πTemp003, πTemp007, πTemp011).ToObject()
					πR = πTemp002
					continue
					goto Label2
				Label2:
					// line 271: r = getattr(typ, "__repr__", None)
					πF.SetLineno(271)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µtyp, "typ"); πE != nil {
						continue
					}
					πTemp001[0] = µtyp
					πTemp001[1] = ß__repr__.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µr = πTemp003
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µtyp, "typ"); πE != nil {
						continue
					}
					πTemp001[0] = µtyp
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
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
						goto Label16
					}
					goto Label17
					// line 272: if issubclass(typ, dict):# and r == dict.__repr__:
					πF.SetLineno(272)
				Label16:
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µo); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label18
					}
					goto Label19
					// line 273: if not o:
					πF.SetLineno(273)
				Label18:
					// line 274: return "{}", True, False
					πF.SetLineno(274)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(πg.NewStr("{}").ToObject(), πTemp003, πTemp007).ToObject()
					πR = πTemp002
					continue
					goto Label19
				Label19:
					// line 275: objid = _id(o)
					πF.SetLineno(275)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					πTemp001[0] = µo
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_id); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µobjid = πTemp003
					if πE = πg.CheckLocal(πF, µmaxlevels, "maxlevels"); πE != nil {
						continue
					}
					πTemp002 = µmaxlevels
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label20
					}
					if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmaxlevels, "maxlevels"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GE(πF, µlevel, µmaxlevels); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label20:
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label21
					}
					goto Label22
					// line 276: if maxlevels and level >= maxlevels:
					πF.SetLineno(276)
				Label21:
					// line 277: return "{...}", False, objid in context
					πF.SetLineno(277)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobjid, "objid"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, µcontext, µobjid); πE != nil {
						continue
					}
					πTemp007 = πg.GetBool(πTemp004).ToObject()
					πTemp002 = πg.NewTuple3(πg.NewStr("{...}").ToObject(), πTemp003, πTemp007).ToObject()
					πR = πTemp002
					continue
					goto Label22
				Label22:
					if πE = πg.CheckLocal(πF, µobjid, "objid"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, µcontext, µobjid); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label23
					}
					goto Label24
					// line 278: if objid in context:
					πF.SetLineno(278)
				Label23:
					// line 279: return _recursion(o), False, True
					πF.SetLineno(279)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					πTemp001[0] = µo
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_recursion); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(πTemp007, πTemp003, πTemp011).ToObject()
					πR = πTemp002
					continue
					goto Label24
				Label24:
					// line 280: context[objid] = 1
					πF.SetLineno(280)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobjid, "objid"); πE != nil {
						continue
					}
					πTemp003 = µobjid
					if πE = πg.SetItem(πF, µcontext, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 281: readable = True
					πF.SetLineno(281)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					µreadable = πTemp002
					// line 282: recursive = False
					πF.SetLineno(282)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					µrecursive = πTemp002
					// line 283: components = []
					πF.SetLineno(283)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µcomponents = πTemp002
					// line 284: append = components.append
					πF.SetLineno(284)
					if πE = πg.CheckLocal(πF, µcomponents, "components"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µcomponents, ßappend, nil); πE != nil {
						continue
					}
					µappend = πTemp002
					// line 285: level += 1
					πF.SetLineno(285)
					if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µlevel, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µlevel = πTemp002
					// line 286: saferepr = _safe_repr
					πF.SetLineno(286)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_safe_repr); πE != nil {
						continue
					}
					µsaferepr = πTemp002
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µo, ßitems, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp007
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_sorted); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Iter(πF, πTemp007); πE != nil {
						continue
					}
					πF.PushCheckpoint(26)
					πTemp004 = false
				Label25:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label27
					}
					if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp011}}}, πTemp003); πE != nil {
							continue
						}
						µk = πTemp007
						µv = πTemp011
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(25)            
					// line 288: krepr, kreadable, krecur = saferepr(k, context, maxlevels, level)
					πF.SetLineno(288)
					πTemp001 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
						continue
					}
					πTemp001[0] = µk
					if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
						continue
					}
					πTemp001[1] = µcontext
					if πE = πg.CheckLocal(πF, µmaxlevels, "maxlevels"); πE != nil {
						continue
					}
					πTemp001[2] = µmaxlevels
					if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
						continue
					}
					πTemp001[3] = µlevel
					if πE = πg.CheckLocal(πF, µsaferepr, "saferepr"); πE != nil {
						continue
					}
					if πTemp003, πE = µsaferepr.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp011}, πg.TieTarget{Target: &πTemp012}}}, πTemp003); πE != nil {
						continue
					}
					µkrepr = πTemp007
					µkreadable = πTemp011
					µkrecur = πTemp012
					// line 289: vrepr, vreadable, vrecur = saferepr(v, context, maxlevels, level)
					πF.SetLineno(289)
					πTemp001 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					πTemp001[0] = µv
					if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
						continue
					}
					πTemp001[1] = µcontext
					if πE = πg.CheckLocal(πF, µmaxlevels, "maxlevels"); πE != nil {
						continue
					}
					πTemp001[2] = µmaxlevels
					if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
						continue
					}
					πTemp001[3] = µlevel
					if πE = πg.CheckLocal(πF, µsaferepr, "saferepr"); πE != nil {
						continue
					}
					if πTemp003, πE = µsaferepr.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp011}, πg.TieTarget{Target: &πTemp012}}}, πTemp003); πE != nil {
						continue
					}
					µvrepr = πTemp007
					µvreadable = πTemp011
					µvrecur = πTemp012
					// line 290: append("%s: %s" % (krepr, vrepr))
					πF.SetLineno(290)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µkrepr, "krepr"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µvrepr, "vrepr"); πE != nil {
						continue
					}
					πTemp007 = πg.NewTuple2(µkrepr, µvrepr).ToObject()
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s: %s").ToObject(), πTemp007); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µappend, "append"); πE != nil {
						continue
					}
					if πTemp003, πE = µappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 291: readable = readable and kreadable and vreadable
					πF.SetLineno(291)
					if πE = πg.CheckLocal(πF, µreadable, "readable"); πE != nil {
						continue
					}
					πTemp003 = µreadable
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label28
					}
					if πE = πg.CheckLocal(πF, µkreadable, "kreadable"); πE != nil {
						continue
					}
					πTemp003 = µkreadable
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label28
					}
					if πE = πg.CheckLocal(πF, µvreadable, "vreadable"); πE != nil {
						continue
					}
					πTemp003 = µvreadable
				Label28:
					µreadable = πTemp003
					if πE = πg.CheckLocal(πF, µkrecur, "krecur"); πE != nil {
						continue
					}
					πTemp003 = µkrecur
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label29
					}
					if πE = πg.CheckLocal(πF, µvrecur, "vrecur"); πE != nil {
						continue
					}
					πTemp003 = µvrecur
				Label29:
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label30
					}
					goto Label31
					// line 292: if krecur or vrecur:
					πF.SetLineno(292)
				Label30:
					// line 293: recursive = True
					πF.SetLineno(293)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					µrecursive = πTemp003
					goto Label31
				Label31:
					continue
				Label26:
					if πE != nil || πR != nil {
						continue
					}
				Label27:
					// line 294: del context[objid]
					πF.SetLineno(294)
					if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobjid, "objid"); πE != nil {
						continue
					}
					πTemp002 = µobjid
					if πE = πg.DelItem(πF, µcontext, πTemp002); πE != nil {
						continue
					}
					// line 295: return "{%s}" % _commajoin(components), readable, recursive
					πF.SetLineno(295)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcomponents, "components"); πE != nil {
						continue
					}
					πTemp001[0] = µcomponents
					if πTemp007, πE = πg.ResolveGlobal(πF, ß_commajoin); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("{%s}").ToObject(), πTemp011); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µreadable, "readable"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µrecursive, "recursive"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(πTemp003, µreadable, µrecursive).ToObject()
					πR = πTemp002
					continue
					goto Label17
				Label17:
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µtyp, "typ"); πE != nil {
						continue
					}
					πTemp001[0] = µtyp
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp002 = πTemp007
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label32
					}
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µtyp, "typ"); πE != nil {
						continue
					}
					πTemp001[0] = µtyp
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp002 = πTemp007
				Label32:
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label33
					}
					goto Label34
					// line 299: if (issubclass(typ, list)) or \
					πF.SetLineno(299)
				Label33:
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µtyp, "typ"); πE != nil {
						continue
					}
					πTemp001[0] = µtyp
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
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
						goto Label35
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					πTemp001[0] = µo
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_len); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Eq(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label36
					}
					goto Label37
					// line 301: if issubclass(typ, list):
					πF.SetLineno(301)
				Label35:
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µo); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label39
					}
					goto Label40
					// line 302: if not o:
					πF.SetLineno(302)
				Label39:
					// line 303: return "[]", True, False
					πF.SetLineno(303)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(πg.NewStr("[]").ToObject(), πTemp003, πTemp007).ToObject()
					πR = πTemp002
					continue
					goto Label40
				Label40:
					// line 304: format = "[%s]"
					πF.SetLineno(304)
					µformat = πg.NewStr("[%s]").ToObject()
					goto Label38
					// line 305: elif _len(o) == 1:
					πF.SetLineno(305)
				Label36:
					// line 306: format = "(%s,)"
					πF.SetLineno(306)
					µformat = πg.NewStr("(%s,)").ToObject()
					goto Label38
				Label37:
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µo); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label41
					}
					goto Label42
					// line 308: if not o:
					πF.SetLineno(308)
				Label41:
					// line 309: return "()", True, False
					πF.SetLineno(309)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(πg.NewStr("()").ToObject(), πTemp003, πTemp007).ToObject()
					πR = πTemp002
					continue
					goto Label42
				Label42:
					// line 310: format = "(%s)"
					πF.SetLineno(310)
					µformat = πg.NewStr("(%s)").ToObject()
					goto Label38
				Label38:
					// line 311: objid = _id(o)
					πF.SetLineno(311)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					πTemp001[0] = µo
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_id); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µobjid = πTemp003
					if πE = πg.CheckLocal(πF, µmaxlevels, "maxlevels"); πE != nil {
						continue
					}
					πTemp002 = µmaxlevels
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label43
					}
					if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmaxlevels, "maxlevels"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GE(πF, µlevel, µmaxlevels); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label43:
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label44
					}
					goto Label45
					// line 312: if maxlevels and level >= maxlevels:
					πF.SetLineno(312)
				Label44:
					// line 313: return format % "...", False, objid in context
					πF.SetLineno(313)
					if πE = πg.CheckLocal(πF, µformat, "format"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mod(πF, µformat, πg.NewStr("...").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobjid, "objid"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, µcontext, µobjid); πE != nil {
						continue
					}
					πTemp011 = πg.GetBool(πTemp004).ToObject()
					πTemp002 = πg.NewTuple3(πTemp003, πTemp007, πTemp011).ToObject()
					πR = πTemp002
					continue
					goto Label45
				Label45:
					if πE = πg.CheckLocal(πF, µobjid, "objid"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, µcontext, µobjid); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label46
					}
					goto Label47
					// line 314: if objid in context:
					πF.SetLineno(314)
				Label46:
					// line 315: return _recursion(o), False, True
					πF.SetLineno(315)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					πTemp001[0] = µo
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_recursion); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(πTemp007, πTemp003, πTemp011).ToObject()
					πR = πTemp002
					continue
					goto Label47
				Label47:
					// line 316: context[objid] = 1
					πF.SetLineno(316)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobjid, "objid"); πE != nil {
						continue
					}
					πTemp003 = µobjid
					if πE = πg.SetItem(πF, µcontext, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 317: readable = True
					πF.SetLineno(317)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					µreadable = πTemp002
					// line 318: recursive = False
					πF.SetLineno(318)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					µrecursive = πTemp002
					// line 319: components = []
					πF.SetLineno(319)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µcomponents = πTemp002
					// line 320: append = components.append
					πF.SetLineno(320)
					if πE = πg.CheckLocal(πF, µcomponents, "components"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µcomponents, ßappend, nil); πE != nil {
						continue
					}
					µappend = πTemp002
					// line 321: level += 1
					πF.SetLineno(321)
					if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µlevel, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µlevel = πTemp002
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µo); πE != nil {
						continue
					}
					πF.PushCheckpoint(49)
					πTemp004 = false
				Label48:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label50
					}
					if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
						µx = πTemp003
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(48)            
					// line 323: orepr, oreadable, orecur = _safe_repr(x, context, maxlevels, level)
					πF.SetLineno(323)
					πTemp001 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
						continue
					}
					πTemp001[1] = µcontext
					if πE = πg.CheckLocal(πF, µmaxlevels, "maxlevels"); πE != nil {
						continue
					}
					πTemp001[2] = µmaxlevels
					if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
						continue
					}
					πTemp001[3] = µlevel
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_safe_repr); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp011}, πg.TieTarget{Target: &πTemp012}}}, πTemp007); πE != nil {
						continue
					}
					µorepr = πTemp003
					µoreadable = πTemp011
					µorecur = πTemp012
					// line 324: append(orepr)
					πF.SetLineno(324)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µorepr, "orepr"); πE != nil {
						continue
					}
					πTemp001[0] = µorepr
					if πE = πg.CheckLocal(πF, µappend, "append"); πE != nil {
						continue
					}
					if πTemp003, πE = µappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µoreadable, "oreadable"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µoreadable); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label51
					}
					goto Label52
					// line 325: if not oreadable:
					πF.SetLineno(325)
				Label51:
					// line 326: readable = False
					πF.SetLineno(326)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					µreadable = πTemp003
					goto Label52
				Label52:
					if πE = πg.CheckLocal(πF, µorecur, "orecur"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µorecur); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label53
					}
					goto Label54
					// line 327: if orecur:
					πF.SetLineno(327)
				Label53:
					// line 328: recursive = True
					πF.SetLineno(328)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					µrecursive = πTemp003
					goto Label54
				Label54:
					continue
				Label49:
					if πE != nil || πR != nil {
						continue
					}
				Label50:
					// line 329: del context[objid]
					πF.SetLineno(329)
					if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobjid, "objid"); πE != nil {
						continue
					}
					πTemp002 = µobjid
					if πE = πg.DelItem(πF, µcontext, πTemp002); πE != nil {
						continue
					}
					// line 330: return format % _commajoin(components), readable, recursive
					πF.SetLineno(330)
					if πE = πg.CheckLocal(πF, µformat, "format"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcomponents, "components"); πE != nil {
						continue
					}
					πTemp001[0] = µcomponents
					if πTemp007, πE = πg.ResolveGlobal(πF, ß_commajoin); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.Mod(πF, µformat, πTemp011); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µreadable, "readable"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µrecursive, "recursive"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(πTemp003, µreadable, µrecursive).ToObject()
					πR = πTemp002
					continue
					goto Label34
				Label34:
					// line 332: rep = repr(o)
					πF.SetLineno(332)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					πTemp001[0] = µo
					if πTemp002, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µrep = πTemp003
					// line 333: return rep, (rep and not rep.startswith('<')), False
					πF.SetLineno(333)
					if πE = πg.CheckLocal(πF, µrep, "rep"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µrep, "rep"); πE != nil {
						continue
					}
					πTemp003 = µrep
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label55
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("<").ToObject()
					if πE = πg.CheckLocal(πF, µrep, "rep"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, µrep, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp012); πE != nil {
						continue
					}
					πTemp007 = πg.GetBool(!πTemp005).ToObject()
					πTemp003 = πTemp007
				Label55:
					if πTemp007, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(µrep, πTemp003, πTemp007).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ß_safe_repr.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 336: def _recursion(o):
			πF.SetLineno(336)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "o", Def: nil}
			πTemp011 = πg.NewFunction(πg.NewCode("_recursion", "build/src/__python__/pprint.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µo *πg.Object = πArgs[0]; _ = µo
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 337: return ("<Recursion on %s with id=%s>"
					πF.SetLineno(337)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					πTemp003[0] = µo
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_type); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp004, πE = πg.GetAttr(πF, πTemp005, ß__name__, nil); πE != nil {
						continue
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					πTemp003[0] = µo
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_id); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp002 = πg.NewTuple2(πTemp004, πTemp006).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("<Recursion on %s with id=%s>").ToObject(), πTemp002); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_recursion.ToObject(), πTemp011); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("pprint", Code)
}
