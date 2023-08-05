package warnings
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/warnings.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßAttributeError := πg.InternStr("AttributeError")
		ßBytesWarning := πg.InternStr("BytesWarning")
		ßDeprecationWarning := πg.InternStr("DeprecationWarning")
		ßException := πg.InternStr("Exception")
		ßFalse := πg.InternStr("False")
		ßI := πg.InternStr("I")
		ßIOError := πg.InternStr("IOError")
		ßImportError := πg.InternStr("ImportError")
		ßImportWarning := πg.InternStr("ImportWarning")
		ßNameError := πg.InternStr("NameError")
		ßNone := πg.InternStr("None")
		ßOverflowError := πg.InternStr("OverflowError")
		ßPendingDeprecationWarning := πg.InternStr("PendingDeprecationWarning")
		ßRuntimeError := πg.InternStr("RuntimeError")
		ßTrue := πg.InternStr("True")
		ßUnicodeDecodeError := πg.InternStr("UnicodeDecodeError")
		ßUnicodeEncodeError := πg.InternStr("UnicodeEncodeError")
		ßUnicodeError := πg.InternStr("UnicodeError")
		ßUserWarning := πg.InternStr("UserWarning")
		ßValueError := πg.InternStr("ValueError")
		ßWarning := πg.InternStr("Warning")
		ßWarningMessage := πg.InternStr("WarningMessage")
		ß_OptionError := πg.InternStr("_OptionError")
		ß_WARNING_DETAILS := πg.InternStr("_WARNING_DETAILS")
		ß__all__ := πg.InternStr("__all__")
		ß__class__ := πg.InternStr("__class__")
		ß__dict__ := πg.InternStr("__dict__")
		ß__enter__ := πg.InternStr("__enter__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__file__ := πg.InternStr("__file__")
		ß__import__ := πg.InternStr("__import__")
		ß__init__ := πg.InternStr("__init__")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__repr__ := πg.InternStr("__repr__")
		ß__str__ := πg.InternStr("__str__")
		ß__warningregistry__ := πg.InternStr("__warningregistry__")
		ß_category_name := πg.InternStr("_category_name")
		ß_entered := πg.InternStr("_entered")
		ß_filters := πg.InternStr("_filters")
		ß_getaction := πg.InternStr("_getaction")
		ß_getcategory := πg.InternStr("_getcategory")
		ß_getframe := πg.InternStr("_getframe")
		ß_module := πg.InternStr("_module")
		ß_processoptions := πg.InternStr("_processoptions")
		ß_record := πg.InternStr("_record")
		ß_setoption := πg.InternStr("_setoption")
		ß_show_warning := πg.InternStr("_show_warning")
		ß_showwarning := πg.InternStr("_showwarning")
		ß_warnings_defaults := πg.InternStr("_warnings_defaults")
		ßall := πg.InternStr("all")
		ßalways := πg.InternStr("always")
		ßappend := πg.InternStr("append")
		ßargv := πg.InternStr("argv")
		ßbasestring := πg.InternStr("basestring")
		ßbytes_action := πg.InternStr("bytes_action")
		ßbytes_warning := πg.InternStr("bytes_warning")
		ßcatch_warnings := πg.InternStr("catch_warnings")
		ßcategory := πg.InternStr("category")
		ßcls := πg.InternStr("cls")
		ßcompile := πg.InternStr("compile")
		ßdefault := πg.InternStr("default")
		ßdefaultaction := πg.InternStr("defaultaction")
		ßdivision_warning := πg.InternStr("division_warning")
		ßendswith := πg.InternStr("endswith")
		ßerror := πg.InternStr("error")
		ßescape := πg.InternStr("escape")
		ßeval := πg.InternStr("eval")
		ßf_globals := πg.InternStr("f_globals")
		ßf_lineno := πg.InternStr("f_lineno")
		ßfile := πg.InternStr("file")
		ßfilename := πg.InternStr("filename")
		ßfilters := πg.InternStr("filters")
		ßfilterwarnings := πg.InternStr("filterwarnings")
		ßflags := πg.InternStr("flags")
		ßformatwarning := πg.InternStr("formatwarning")
		ßget := πg.InternStr("get")
		ßgetattr := πg.InternStr("getattr")
		ßgetfilesystemencoding := πg.InternStr("getfilesystemencoding")
		ßgetline := πg.InternStr("getline")
		ßgetlines := πg.InternStr("getlines")
		ßignore := πg.InternStr("ignore")
		ßinsert := πg.InternStr("insert")
		ßint := πg.InternStr("int")
		ßisinstance := πg.InternStr("isinstance")
		ßissubclass := πg.InternStr("issubclass")
		ßjoin := πg.InternStr("join")
		ßlatin1 := πg.InternStr("latin1")
		ßlen := πg.InternStr("len")
		ßline := πg.InternStr("line")
		ßlinecache := πg.InternStr("linecache")
		ßlineno := πg.InternStr("lineno")
		ßlocals := πg.InternStr("locals")
		ßlower := πg.InternStr("lower")
		ßmatch := πg.InternStr("match")
		ßmessage := πg.InternStr("message")
		ßmodule := πg.InternStr("module")
		ßmodules := πg.InternStr("modules")
		ßobject := πg.InternStr("object")
		ßonce := πg.InternStr("once")
		ßonceregistry := πg.InternStr("onceregistry")
		ßpy3kwarning := πg.InternStr("py3kwarning")
		ßre := πg.InternStr("re")
		ßresetwarnings := πg.InternStr("resetwarnings")
		ßrfind := πg.InternStr("rfind")
		ßsetattr := πg.InternStr("setattr")
		ßsetdefault := πg.InternStr("setdefault")
		ßshowwarning := πg.InternStr("showwarning")
		ßsilence := πg.InternStr("silence")
		ßsimplefilter := πg.InternStr("simplefilter")
		ßsplit := πg.InternStr("split")
		ßstartswith := πg.InternStr("startswith")
		ßstderr := πg.InternStr("stderr")
		ßstr := πg.InternStr("str")
		ßstrip := πg.InternStr("strip")
		ßsys := πg.InternStr("sys")
		ßtype := πg.InternStr("type")
		ßtypes := πg.InternStr("types")
		ßunicode := πg.InternStr("unicode")
		ßwarn := πg.InternStr("warn")
		ßwarn_explicit := πg.InternStr("warn_explicit")
		ßwarnings := πg.InternStr("warnings")
		ßwarnoptions := πg.InternStr("warnoptions")
		ßwarnpy3k := πg.InternStr("warnpy3k")
		ßwrite := πg.InternStr("write")
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
		var πTemp019 bool
		_ = πTemp019
		var πTemp020 *πg.Object
		_ = πTemp020
		var πTemp021 bool
		_ = πTemp021
		var πTemp022 πg.KWArgs
		_ = πTemp022
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 6: goto Label6
			case 7: goto Label7
			default: panic("unexpected function state")
			}
			// line 1: """Python part of the warnings subsystem."""
			πF.SetLineno(1)
			// line 6: import linecache
			πF.SetLineno(6)
			if πTemp002, πE = πg.ImportModule(πF, "linecache"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßlinecache.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 7: import re
			πF.SetLineno(7)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 8: import sys
			πF.SetLineno(8)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 9: import types
			πF.SetLineno(9)
			if πTemp002, πE = πg.ImportModule(πF, "types"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtypes.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 11: __all__ = ["warn", "warn_explicit", "showwarning",
			πF.SetLineno(11)
			πTemp002 = make([]*πg.Object, 9)
			πTemp002[0] = ßwarn.ToObject()
			πTemp002[1] = ßwarn_explicit.ToObject()
			πTemp002[2] = ßshowwarning.ToObject()
			πTemp002[3] = ßformatwarning.ToObject()
			πTemp002[4] = ßfilterwarnings.ToObject()
			πTemp002[5] = ßsimplefilter.ToObject()
			πTemp002[6] = ßresetwarnings.ToObject()
			πTemp002[7] = ßcatch_warnings.ToObject()
			πTemp002[8] = ßwarnpy3k.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 16: def warnpy3k(message, category=None, stacklevel=1):
			πF.SetLineno(16)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "message", Def: nil}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[1] = πg.Param{Name: "category", Def: πTemp004}
			πTemp003[2] = πg.Param{Name: "stacklevel", Def: πg.NewInt(1).ToObject()}
			πTemp001 = πg.NewFunction(πg.NewCode("warnpy3k", "build/src/__python__/warnings.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmessage *πg.Object = πArgs[0]; _ = µmessage
				var µcategory *πg.Object = πArgs[1]; _ = µcategory
				var µstacklevel *πg.Object = πArgs[2]; _ = µstacklevel
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
					// line 17: """Issue a deprecation warning for Python 3.x related changes.
					πF.SetLineno(17)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpy3kwarning, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 21: if sys.py3kwarning:
					πF.SetLineno(21)
				Label1:
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µcategory == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label3
					}
					goto Label4
					// line 22: if category is None:
					πF.SetLineno(22)
				Label3:
					// line 23: category = DeprecationWarning
					πF.SetLineno(23)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßDeprecationWarning); πE != nil {
						continue
					}
					µcategory = πTemp001
					goto Label4
				Label4:
					// line 24: warn(message, category, stacklevel+1)
					πF.SetLineno(24)
					πTemp004 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
						continue
					}
					πTemp004[0] = µmessage
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					πTemp004[1] = µcategory
					if πE = πg.CheckLocal(πF, µstacklevel, "stacklevel"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µstacklevel, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp004[2] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßwarn); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					goto Label2
				Label2:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßwarnpy3k.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 26: def _show_warning(message, category, filename, lineno, file=None, line=None):
			πF.SetLineno(26)
			πTemp003 = make([]πg.Param, 6)
			πTemp003[0] = πg.Param{Name: "message", Def: nil}
			πTemp003[1] = πg.Param{Name: "category", Def: nil}
			πTemp003[2] = πg.Param{Name: "filename", Def: nil}
			πTemp003[3] = πg.Param{Name: "lineno", Def: nil}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[4] = πg.Param{Name: "file", Def: πTemp005}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[5] = πg.Param{Name: "line", Def: πTemp005}
			πTemp004 = πg.NewFunction(πg.NewCode("_show_warning", "build/src/__python__/warnings.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmessage *πg.Object = πArgs[0]; _ = µmessage
				var µcategory *πg.Object = πArgs[1]; _ = µcategory
				var µfilename *πg.Object = πArgs[2]; _ = µfilename
				var µlineno *πg.Object = πArgs[3]; _ = µlineno
				var µfile *πg.Object = πArgs[4]; _ = µfile
				var µline *πg.Object = πArgs[5]; _ = µline
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
				var πTemp008 *πg.Object
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 6: goto Label6
					default: panic("unexpected function state")
					}
					// line 27: """Hook to write a warning to a file; replace if you like."""
					πF.SetLineno(27)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µfile == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 28: if file is None:
					πF.SetLineno(28)
				Label1:
					// line 29: file = sys.stderr
					πF.SetLineno(29)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstderr, nil); πE != nil {
						continue
					}
					µfile = πTemp002
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µfile == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label3
					}
					goto Label4
					// line 30: if file is None:
					πF.SetLineno(30)
				Label3:
					// line 32: return
					πF.SetLineno(32)
					πR = πg.None
					continue
					goto Label4
				Label4:
					goto Label2
				Label2:
					// line 33: try:
					πF.SetLineno(33)
					πF.PushCheckpoint(6)
					// line 34: file.write(formatwarning(message, category, filename, lineno, line))
					πF.SetLineno(34)
					πTemp004 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(5)
					if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
						continue
					}
					πTemp005[0] = µmessage
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					πTemp005[1] = µcategory
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp005[2] = µfilename
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					πTemp005[3] = µlineno
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp005[4] = µline
					if πTemp001, πE = πg.ResolveGlobal(πF, ßformatwarning); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp004[0] = πTemp002
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfile, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πF.PopCheckpoint()
					goto Label5
				Label6:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp006, πTemp007 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveGlobal(πF, ßUnicodeError); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp002, πTemp008).ToObject()
					if πTemp003, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label7
					}
					πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
					continue
					// line 35: except (IOError, UnicodeError):
					πF.SetLineno(35)
				Label7:
					// line 36: pass # the file (probably stderr) is invalid - this warning gets lost.
					πF.SetLineno(36)
					πF.RestoreExc(nil, nil)
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
			if πE = πF.Globals().SetItem(πF, ß_show_warning.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 39: showwarning = _show_warning
			πF.SetLineno(39)
			if πTemp005, πE = πg.ResolveGlobal(πF, ß_show_warning); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßshowwarning.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 41: def formatwarning(message, category, filename, lineno, line=None):
			πF.SetLineno(41)
			πTemp003 = make([]πg.Param, 5)
			πTemp003[0] = πg.Param{Name: "message", Def: nil}
			πTemp003[1] = πg.Param{Name: "category", Def: nil}
			πTemp003[2] = πg.Param{Name: "filename", Def: nil}
			πTemp003[3] = πg.Param{Name: "lineno", Def: nil}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[4] = πg.Param{Name: "line", Def: πTemp006}
			πTemp005 = πg.NewFunction(πg.NewCode("formatwarning", "build/src/__python__/warnings.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmessage *πg.Object = πArgs[0]; _ = µmessage
				var µcategory *πg.Object = πArgs[1]; _ = µcategory
				var µfilename *πg.Object = πArgs[2]; _ = µfilename
				var µlineno *πg.Object = πArgs[3]; _ = µlineno
				var µline *πg.Object = πArgs[4]; _ = µline
				var µunicodetype *πg.Object = πg.UnboundLocal; _ = µunicodetype
				var µs *πg.Object = πg.UnboundLocal; _ = µs
				var µenc *πg.Object = πg.UnboundLocal; _ = µenc
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.BaseException
				_ = πTemp002
				var πTemp003 *πg.Traceback
				_ = πTemp003
				var πTemp004 bool
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
					case 2: goto Label2
					case 20: goto Label20
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 42: """Function to format a warning the standard way."""
					πF.SetLineno(42)
					// line 43: try:
					πF.SetLineno(43)
					πF.PushCheckpoint(2)
					// line 44: unicodetype = unicode
					πF.SetLineno(44)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
						continue
					}
					µunicodetype = πTemp001
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp002, πTemp003 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNameError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp002.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					πE = πF.Raise(πTemp002.ToObject(), nil, πTemp003.ToObject())
					continue
					// line 45: except NameError:
					πF.SetLineno(45)
				Label3:
					// line 46: unicodetype = ()
					πF.SetLineno(46)
					πTemp001 = πg.NewTuple0().ToObject()
					µunicodetype = πTemp001
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 47: try:
					πF.SetLineno(47)
					πF.PushCheckpoint(5)
					// line 48: message = str(message)
					πF.SetLineno(48)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
						continue
					}
					πTemp005[0] = µmessage
					if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µmessage = πTemp006
					πF.PopCheckpoint()
					goto Label4
				Label5:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp002, πTemp003 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßUnicodeEncodeError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp002.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label6
					}
					πE = πF.Raise(πTemp002.ToObject(), nil, πTemp003.ToObject())
					continue
					// line 49: except UnicodeEncodeError:
					πF.SetLineno(49)
				Label6:
					// line 50: pass
					πF.SetLineno(50)
					πF.RestoreExc(nil, nil)
					goto Label4
				Label4:
					// line 51: s =  "%s: %s: %s\n" % (lineno, category.__name__, message)
					πF.SetLineno(51)
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µcategory, ß__name__, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
						continue
					}
					πTemp006 = πg.NewTuple3(µlineno, πTemp007, µmessage).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s: %s: %s\n").ToObject(), πTemp006); πE != nil {
						continue
					}
					µs = πTemp001
					// line 52: line = linecache.getline(filename, lineno) if line is None else line
					πF.SetLineno(52)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp006 = πg.GetBool(µline == πTemp007).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label7
					}
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp005[0] = µfilename
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					πTemp005[1] = µlineno
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlinecache); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßgetline, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp007.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001 = πTemp006
					goto Label8
				Label7:
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp001 = µline
				Label8:
					µline = πTemp001
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µline); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label9
					}
					goto Label10
					// line 53: if line:
					πF.SetLineno(53)
				Label9:
					// line 54: line = line.strip()
					πF.SetLineno(54)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µline, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µline = πTemp006
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp005[0] = µs
					if πE = πg.CheckLocal(πF, µunicodetype, "unicodetype"); πE != nil {
						continue
					}
					πTemp005[1] = µunicodetype
					if πTemp006, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001 = πTemp007
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label11
					}
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp005[0] = µline
					if πTemp006, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					πTemp005[1] = πTemp006
					if πTemp006, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001 = πTemp007
				Label11:
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label12
					}
					goto Label13
					// line 55: if isinstance(s, unicodetype) and isinstance(line, str):
					πF.SetLineno(55)
				Label12:
					// line 56: line = unicode(line, 'latin1')
					πF.SetLineno(56)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp005[0] = µline
					πTemp005[1] = ßlatin1.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µline = πTemp006
					goto Label13
				Label13:
					// line 57: s += "  %s\n" % line
					πF.SetLineno(57)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("  %s\n").ToObject(), µline); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IAdd(πF, µs, πTemp001); πE != nil {
						continue
					}
					µs = πTemp006
					goto Label10
				Label10:
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp005[0] = µs
					if πE = πg.CheckLocal(πF, µunicodetype, "unicodetype"); πE != nil {
						continue
					}
					πTemp005[1] = µunicodetype
					if πTemp006, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001 = πTemp007
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label14
					}
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp005[0] = µfilename
					if πTemp006, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					πTemp005[1] = πTemp006
					if πTemp006, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001 = πTemp007
				Label14:
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label15
					}
					goto Label16
					// line 58: if isinstance(s, unicodetype) and isinstance(filename, str):
					πF.SetLineno(58)
				Label15:
					// line 59: enc = sys.getfilesystemencoding()
					πF.SetLineno(59)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp001, ßgetfilesystemencoding, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp006.Call(πF, nil, nil); πE != nil {
						continue
					}
					µenc = πTemp001
					if πE = πg.CheckLocal(πF, µenc, "enc"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µenc); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label17
					}
					goto Label18
					// line 60: if enc:
					πF.SetLineno(60)
				Label17:
					// line 61: try:
					πF.SetLineno(61)
					πF.PushCheckpoint(20)
					// line 62: filename = unicode(filename, enc)
					πF.SetLineno(62)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp005[0] = µfilename
					if πE = πg.CheckLocal(πF, µenc, "enc"); πE != nil {
						continue
					}
					πTemp005[1] = µenc
					if πTemp001, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µfilename = πTemp006
					πF.PopCheckpoint()
					goto Label19
				Label20:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp002, πTemp003 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßUnicodeDecodeError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp002.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label21
					}
					πE = πF.Raise(πTemp002.ToObject(), nil, πTemp003.ToObject())
					continue
					// line 63: except UnicodeDecodeError:
					πF.SetLineno(63)
				Label21:
					// line 64: pass
					πF.SetLineno(64)
					πF.RestoreExc(nil, nil)
					goto Label19
				Label19:
					goto Label18
				Label18:
					goto Label16
				Label16:
					// line 65: s = "%s:%s" % (filename, s)
					πF.SetLineno(65)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp006 = πg.NewTuple2(µfilename, µs).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s:%s").ToObject(), πTemp006); πE != nil {
						continue
					}
					µs = πTemp001
					// line 66: return s
					πF.SetLineno(66)
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
			if πE = πF.Globals().SetItem(πF, ßformatwarning.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 68: def filterwarnings(action, message="", category=Warning, module="", lineno=0,
			πF.SetLineno(68)
			πTemp003 = make([]πg.Param, 6)
			πTemp003[0] = πg.Param{Name: "action", Def: nil}
			πTemp003[1] = πg.Param{Name: "message", Def: ß.ToObject()}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßWarning); πE != nil {
				continue
			}
			πTemp003[2] = πg.Param{Name: "category", Def: πTemp007}
			πTemp003[3] = πg.Param{Name: "module", Def: ß.ToObject()}
			πTemp003[4] = πg.Param{Name: "lineno", Def: πg.NewInt(0).ToObject()}
			πTemp003[5] = πg.Param{Name: "append", Def: πg.NewInt(0).ToObject()}
			πTemp006 = πg.NewFunction(πg.NewCode("filterwarnings", "build/src/__python__/warnings.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µaction *πg.Object = πArgs[0]; _ = µaction
				var µmessage *πg.Object = πArgs[1]; _ = µmessage
				var µcategory *πg.Object = πArgs[2]; _ = µcategory
				var µmodule *πg.Object = πArgs[3]; _ = µmodule
				var µlineno *πg.Object = πArgs[4]; _ = µlineno
				var µappend *πg.Object = πArgs[5]; _ = µappend
				var µitem *πg.Object = πg.UnboundLocal; _ = µitem
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 70: """Insert an entry into the list of warnings filters (at the front).
					πF.SetLineno(70)
					// line 80: assert action in ("error", "ignore", "always", "default", "module",
					πF.SetLineno(80)
					if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple1(µaction).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("invalid action: %r").ToObject(), πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple6(ßerror.ToObject(), ßignore.ToObject(), ßalways.ToObject(), ßdefault.ToObject(), ßmodule.ToObject(), ßonce.ToObject()).ToObject()
					if πTemp004, πE = πg.Contains(πF, πTemp003, µaction); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp004).ToObject()
					if πE = πg.Assert(πF, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 82: assert isinstance(message, basestring), "message must be a string"
					πF.SetLineno(82)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
						continue
					}
					πTemp005[0] = µmessage
					if πTemp001, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
						continue
					}
					πTemp005[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πg.Assert(πF, πTemp002, πg.NewStr("message must be a string").ToObject()); πE != nil {
						continue
					}
					// line 83: assert isinstance(category, type), "category must be a class"
					πF.SetLineno(83)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					πTemp005[0] = µcategory
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					πTemp005[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πg.Assert(πF, πTemp002, πg.NewStr("category must be a class").ToObject()); πE != nil {
						continue
					}
					// line 84: assert issubclass(category, Warning), "category must be a Warning subclass"
					πF.SetLineno(84)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					πTemp005[0] = µcategory
					if πTemp001, πE = πg.ResolveGlobal(πF, ßWarning); πE != nil {
						continue
					}
					πTemp005[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πg.Assert(πF, πTemp002, πg.NewStr("category must be a Warning subclass").ToObject()); πE != nil {
						continue
					}
					// line 85: assert isinstance(module, basestring), "module must be a string"
					πF.SetLineno(85)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					πTemp005[0] = µmodule
					if πTemp001, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
						continue
					}
					πTemp005[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πg.Assert(πF, πTemp002, πg.NewStr("module must be a string").ToObject()); πE != nil {
						continue
					}
					// line 86: assert isinstance(lineno, int) and lineno >= 0, \
					πF.SetLineno(86)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					πTemp005[0] = µlineno
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					πTemp005[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001 = πTemp003
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GE(πF, µlineno, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp002
				Label1:
					if πE = πg.Assert(πF, πTemp001, πg.NewStr("lineno must be an int >= 0").ToObject()); πE != nil {
						continue
					}
					// line 88: item = (action, re.compile(message, re.I), category,
					πF.SetLineno(88)
					if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
						continue
					}
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
						continue
					}
					πTemp005[0] = µmessage
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßI, nil); πE != nil {
						continue
					}
					πTemp005[1] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					πTemp005[0] = µmodule
					if πTemp003, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple5(µaction, πTemp002, µcategory, πTemp003, µlineno).ToObject()
					µitem = πTemp001
					if πE = πg.CheckLocal(πF, µappend, "append"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µappend); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label2
					}
					goto Label3
					// line 90: if append:
					πF.SetLineno(90)
				Label2:
					// line 91: filters.append(item)
					πF.SetLineno(91)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp005[0] = µitem
					if πTemp001, πE = πg.ResolveGlobal(πF, ßfilters); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					goto Label4
				Label3:
					// line 93: filters.insert(0, item)
					πF.SetLineno(93)
					πTemp005 = πF.MakeArgs(2)
					πTemp005[0] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp005[1] = µitem
					if πTemp001, πE = πg.ResolveGlobal(πF, ßfilters); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinsert, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
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
			if πE = πF.Globals().SetItem(πF, ßfilterwarnings.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 95: def simplefilter(action, category=Warning, lineno=0, append=0):
			πF.SetLineno(95)
			πTemp003 = make([]πg.Param, 4)
			πTemp003[0] = πg.Param{Name: "action", Def: nil}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßWarning); πE != nil {
				continue
			}
			πTemp003[1] = πg.Param{Name: "category", Def: πTemp008}
			πTemp003[2] = πg.Param{Name: "lineno", Def: πg.NewInt(0).ToObject()}
			πTemp003[3] = πg.Param{Name: "append", Def: πg.NewInt(0).ToObject()}
			πTemp007 = πg.NewFunction(πg.NewCode("simplefilter", "build/src/__python__/warnings.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µaction *πg.Object = πArgs[0]; _ = µaction
				var µcategory *πg.Object = πArgs[1]; _ = µcategory
				var µlineno *πg.Object = πArgs[2]; _ = µlineno
				var µappend *πg.Object = πArgs[3]; _ = µappend
				var µitem *πg.Object = πg.UnboundLocal; _ = µitem
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 96: """Insert a simple entry into the list of warnings filters (at the front).
					πF.SetLineno(96)
					// line 105: assert action in ("error", "ignore", "always", "default", "module",
					πF.SetLineno(105)
					if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple1(µaction).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("invalid action: %r").ToObject(), πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple6(ßerror.ToObject(), ßignore.ToObject(), ßalways.ToObject(), ßdefault.ToObject(), ßmodule.ToObject(), ßonce.ToObject()).ToObject()
					if πTemp004, πE = πg.Contains(πF, πTemp003, µaction); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp004).ToObject()
					if πE = πg.Assert(πF, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 107: assert isinstance(lineno, int) and lineno >= 0, \
					πF.SetLineno(107)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					πTemp005[0] = µlineno
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					πTemp005[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001 = πTemp003
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GE(πF, µlineno, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp002
				Label1:
					if πE = πg.Assert(πF, πTemp001, πg.NewStr("lineno must be an int >= 0").ToObject()); πE != nil {
						continue
					}
					// line 109: item = (action, None, category, None, lineno)
					πF.SetLineno(109)
					if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple5(µaction, πTemp002, µcategory, πTemp003, µlineno).ToObject()
					µitem = πTemp001
					if πE = πg.CheckLocal(πF, µappend, "append"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µappend); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label2
					}
					goto Label3
					// line 110: if append:
					πF.SetLineno(110)
				Label2:
					// line 111: filters.append(item)
					πF.SetLineno(111)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp005[0] = µitem
					if πTemp001, πE = πg.ResolveGlobal(πF, ßfilters); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					goto Label4
				Label3:
					// line 113: filters.insert(0, item)
					πF.SetLineno(113)
					πTemp005 = πF.MakeArgs(2)
					πTemp005[0] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp005[1] = µitem
					if πTemp001, πE = πg.ResolveGlobal(πF, ßfilters); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinsert, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
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
			if πE = πF.Globals().SetItem(πF, ßsimplefilter.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 115: def resetwarnings():
			πF.SetLineno(115)
			πTemp003 = make([]πg.Param, 0)
			πTemp008 = πg.NewFunction(πg.NewCode("resetwarnings", "build/src/__python__/warnings.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 116: """Clear the list of warning filters, so that no filters are active."""
					πF.SetLineno(116)
					// line 117: filters[:] = []
					πF.SetLineno(117)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßfilters); πE != nil {
						continue
					}
					if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.SetItem(πF, πTemp004, πTemp005, πTemp003); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßresetwarnings.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 119: class _OptionError(Exception):
			πF.SetLineno(119)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp012, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
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
			_, πE = πg.NewCode("_OptionError", "build/src/__python__/warnings.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 120: """Exception used by option processing helpers."""
					πF.SetLineno(120)
					// line 121: pass
					πF.SetLineno(121)
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
			if πTemp012, πE = πTemp011.Call(πF, []*πg.Object{πg.NewStr("_OptionError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_OptionError.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 124: def _processoptions(args):
			πF.SetLineno(124)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "args", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("_processoptions", "build/src/__python__/warnings.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargs *πg.Object = πArgs[0]; _ = µargs
				var µarg *πg.Object = πg.UnboundLocal; _ = µarg
				var µmsg *πg.Object = πg.UnboundLocal; _ = µmsg
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
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µargs); πE != nil {
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
						µarg = πTemp004
					}
					if πE != nil || !πTemp003 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 126: try:
					πF.SetLineno(126)
					πF.PushCheckpoint(5)
					// line 127: _setoption(arg)
					πF.SetLineno(127)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
						continue
					}
					πTemp005[0] = µarg
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_setoption); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πF.PopCheckpoint()
					goto Label4
				Label5:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_OptionError); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp004); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label6
					}
					πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
					continue
					// line 128: except _OptionError, msg:
					πF.SetLineno(128)
				Label6:
					µmsg = πTemp007.ToObject()
					// line 129: print >>sys.stderr, "Invalid -W option ignored:", msg
					πF.SetLineno(129)
					πTemp005 = make([]*πg.Object, 2)
					πTemp005[0] = πg.NewStr("Invalid -W option ignored:").ToObject()
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp005[1] = µmsg
					if πE = πg.Print(πF, πTemp005, true); πE != nil {
						continue
					}
					πF.RestoreExc(nil, nil)
					goto Label4
				Label4:
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
			if πE = πF.Globals().SetItem(πF, ß_processoptions.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 132: def _setoption(arg):
			πF.SetLineno(132)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "arg", Def: nil}
			πTemp011 = πg.NewFunction(πg.NewCode("_setoption", "build/src/__python__/warnings.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µarg *πg.Object = πArgs[0]; _ = µarg
				var µparts *πg.Object = πg.UnboundLocal; _ = µparts
				var µaction *πg.Object = πg.UnboundLocal; _ = µaction
				var µmessage *πg.Object = πg.UnboundLocal; _ = µmessage
				var µcategory *πg.Object = πg.UnboundLocal; _ = µcategory
				var µmodule *πg.Object = πg.UnboundLocal; _ = µmodule
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
				var πTemp012 *πg.BaseException
				_ = πTemp012
				var πTemp013 *πg.Traceback
				_ = πTemp013
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 12: goto Label12
					case 3: goto Label3
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					// line 133: parts = arg.split(':')
					πF.SetLineno(133)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr(":").ToObject()
					if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µarg, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µparts = πTemp003
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
						continue
					}
					πTemp001[0] = µparts
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.GT(πF, πTemp004, πg.NewInt(5).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 134: if len(parts) > 5:
					πF.SetLineno(134)
				Label1:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple1(µarg).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("too many fields (max 5): %r").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_OptionError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 135: raise _OptionError("too many fields (max 5): %r" % (arg,))
					πF.SetLineno(135)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label2
				Label2:
					// line 136: while len(parts) < 5:
					πF.SetLineno(136)
					πF.PushCheckpoint(4)
					πTemp005 = false
				Label3:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp005 {
						πF.PopCheckpoint()
						goto Label5
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
						continue
					}
					πTemp001[0] = µparts
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.LT(πF, πTemp004, πg.NewInt(5).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(3)            
					// line 137: parts.append('')
					πF.SetLineno(137)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ß.ToObject()
					if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µparts, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label4:
					if πE != nil || πR != nil {
						continue
					}
				Label5:
					// line 138: action, message, category, module, lineno = [s.strip()
					πF.SetLineno(138)
					πTemp007 = make([]πg.Param, 0)
					πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/warnings.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
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
								if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µparts); πE != nil {
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
									µs = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 138: action, message, category, module, lineno = [s.strip()
								πF.SetLineno(138)
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µs, ßstrip, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return πTemp005, nil
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
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}, πg.TieTarget{Target: &πTemp010}, πg.TieTarget{Target: &πTemp011}}}, πTemp002); πE != nil {
						continue
					}
					µaction = πTemp004
					µmessage = πTemp008
					µcategory = πTemp009
					µmodule = πTemp010
					µlineno = πTemp011
					// line 140: action = _getaction(action)
					πF.SetLineno(140)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
						continue
					}
					πTemp001[0] = µaction
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_getaction); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µaction = πTemp004
					// line 141: message = re.escape(message)
					πF.SetLineno(141)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
						continue
					}
					πTemp001[0] = µmessage
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßescape, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µmessage = πTemp002
					// line 142: category = _getcategory(category)
					πF.SetLineno(142)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					πTemp001[0] = µcategory
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_getcategory); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcategory = πTemp004
					// line 143: module = re.escape(module)
					πF.SetLineno(143)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					πTemp001[0] = µmodule
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßescape, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µmodule = πTemp002
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µmodule); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label6
					}
					goto Label7
					// line 144: if module:
					πF.SetLineno(144)
				Label6:
					// line 145: module = module + '$'
					πF.SetLineno(145)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µmodule, πg.NewStr("$").ToObject()); πE != nil {
						continue
					}
					µmodule = πTemp002
					goto Label7
				Label7:
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µlineno); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label8
					}
					goto Label9
					// line 146: if lineno:
					πF.SetLineno(146)
				Label8:
					// line 147: try:
					πF.SetLineno(147)
					πF.PushCheckpoint(12)
					// line 148: lineno = int(lineno)
					πF.SetLineno(148)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					πTemp001[0] = µlineno
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µlineno = πTemp004
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LT(πF, µlineno, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label13
					}
					goto Label14
					// line 149: if lineno < 0:
					πF.SetLineno(149)
				Label13:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					// line 150: raise ValueError
					πF.SetLineno(150)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label14
				Label14:
					πF.PopCheckpoint()
					goto Label11
				Label12:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp012, πTemp013 = πF.ExcInfo()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveGlobal(πF, ßOverflowError); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp004, πTemp008).ToObject()
					if πTemp005, πE = πg.IsInstance(πF, πTemp012.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label15
					}
					πE = πF.Raise(πTemp012.ToObject(), nil, πTemp013.ToObject())
					continue
					// line 151: except (ValueError, OverflowError):
					πF.SetLineno(151)
				Label15:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple1(µlineno).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("invalid lineno %r").ToObject(), πTemp004); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_OptionError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 152: raise _OptionError("invalid lineno %r" % (lineno,))
					πF.SetLineno(152)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label11
				Label11:
					goto Label10
				Label9:
					// line 154: lineno = 0
					πF.SetLineno(154)
					µlineno = πg.NewInt(0).ToObject()
					goto Label10
				Label10:
					// line 155: filterwarnings(action, message, category, module, lineno)
					πF.SetLineno(155)
					πTemp001 = πF.MakeArgs(5)
					if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
						continue
					}
					πTemp001[0] = µaction
					if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
						continue
					}
					πTemp001[1] = µmessage
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					πTemp001[2] = µcategory
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					πTemp001[3] = µmodule
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					πTemp001[4] = µlineno
					if πTemp002, πE = πg.ResolveGlobal(πF, ßfilterwarnings); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_setoption.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 158: def _getaction(action):
			πF.SetLineno(158)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "action", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("_getaction", "build/src/__python__/warnings.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µaction *πg.Object = πArgs[0]; _ = µaction
				var µa *πg.Object = πg.UnboundLocal; _ = µa
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 5: goto Label5
					case 6: goto Label6
					default: panic("unexpected function state")
					}
					if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µaction); πE != nil {
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
					// line 159: if not action:
					πF.SetLineno(159)
				Label1:
					// line 160: return "default"
					πF.SetLineno(160)
					πR = ßdefault.ToObject()
					continue
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µaction, ßall.ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 161: if action == "all": return "always" # Alias
					πF.SetLineno(161)
				Label3:
					// line 161: if action == "all": return "always" # Alias
					πF.SetLineno(161)
					πR = ßalways.ToObject()
					continue
					goto Label4
				Label4:
					πTemp003 = πg.NewTuple6(ßdefault.ToObject(), ßalways.ToObject(), ßignore.ToObject(), ßmodule.ToObject(), ßonce.ToObject(), ßerror.ToObject()).ToObject()
					if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
						continue
					}
					πF.PushCheckpoint(6)
					πTemp002 = false
				Label5:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
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
						πTemp004 = !isStop
					} else {
						πTemp004 = true
						µa = πTemp003
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(5)            
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
						continue
					}
					πTemp005[0] = µaction
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µa, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp004, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label8
					}
					goto Label9
					// line 163: if a.startswith(action):
					πF.SetLineno(163)
				Label8:
					// line 164: return a
					πF.SetLineno(164)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πR = µa
					continue
					goto Label9
				Label9:
					continue
				Label6:
					if πE != nil || πR != nil {
						continue
					}
				Label7:
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple1(µaction).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("invalid action: %r").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp005[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_OptionError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 165: raise _OptionError("invalid action: %r" % (action,))
					πF.SetLineno(165)
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
			if πE = πF.Globals().SetItem(πF, ß_getaction.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 168: def _getcategory(category):
			πF.SetLineno(168)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "category", Def: nil}
			πTemp013 = πg.NewFunction(πg.NewCode("_getcategory", "build/src/__python__/warnings.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcategory *πg.Object = πArgs[0]; _ = µcategory
				var µcat *πg.Object = πg.UnboundLocal; _ = µcat
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µmodule *πg.Object = πg.UnboundLocal; _ = µmodule
				var µklass *πg.Object = πg.UnboundLocal; _ = µklass
				var µm *πg.Object = πg.UnboundLocal; _ = µm
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.BaseException
				_ = πTemp005
				var πTemp006 *πg.Traceback
				_ = πTemp006
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 10: goto Label10
					case 13: goto Label13
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µcategory); πE != nil {
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
					// line 169: if not category:
					πF.SetLineno(169)
				Label1:
					// line 170: return Warning
					πF.SetLineno(170)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßWarning); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label2
				Label2:
					πTemp003 = πF.MakeArgs(2)
					πTemp003[0] = πg.NewStr("^[a-zA-Z0-9_]+$").ToObject()
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					πTemp003[1] = µcategory
					if πTemp001, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßmatch, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 171: if re.match("^[a-zA-Z0-9_]+$", category):
					πF.SetLineno(171)
				Label3:
					// line 172: try:
					πF.SetLineno(172)
					πF.PushCheckpoint(7)
					// line 173: cat = eval(category)
					πF.SetLineno(173)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					πTemp003[0] = µcategory
					if πTemp001, πE = πg.ResolveGlobal(πF, ßeval); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µcat = πTemp004
					πF.PopCheckpoint()
					goto Label6
				Label7:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp005, πTemp006 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNameError); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label8
					}
					πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
					continue
					// line 174: except NameError:
					πF.SetLineno(174)
				Label8:
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple1(µcategory).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("unknown warning category: %r").ToObject(), πTemp004); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_OptionError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 175: raise _OptionError("unknown warning category: %r" % (category,))
					πF.SetLineno(175)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label6
				Label6:
					goto Label5
				Label4:
					// line 177: i = category.rfind(".")
					πF.SetLineno(177)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr(".").ToObject()
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µcategory, ßrfind, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µi = πTemp004
					// line 178: module = category[:i]
					πF.SetLineno(178)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µi, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µcategory, πTemp001); πE != nil {
						continue
					}
					µmodule = πTemp004
					// line 179: klass = category[i+1:]
					πF.SetLineno(179)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πTemp004, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µcategory, πTemp001); πE != nil {
						continue
					}
					µklass = πTemp004
					// line 180: try:
					πF.SetLineno(180)
					πF.PushCheckpoint(10)
					// line 181: m = __import__(module, None, None, [klass])
					πF.SetLineno(181)
					πTemp003 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					πTemp003[0] = µmodule
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003[2] = πTemp001
					πTemp007 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µklass, "klass"); πE != nil {
						continue
					}
					πTemp007[0] = µklass
					πTemp001 = πg.NewList(πTemp007...).ToObject()
					πTemp003[3] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ß__import__); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µm = πTemp004
					πF.PopCheckpoint()
					goto Label9
				Label10:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp005, πTemp006 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label11
					}
					πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
					continue
					// line 182: except ImportError:
					πF.SetLineno(182)
				Label11:
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple1(µmodule).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("invalid module name: %r").ToObject(), πTemp004); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_OptionError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 183: raise _OptionError("invalid module name: %r" % (module,))
					πF.SetLineno(183)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label9
				Label9:
					// line 184: try:
					πF.SetLineno(184)
					πF.PushCheckpoint(13)
					// line 185: cat = getattr(m, klass)
					πF.SetLineno(185)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
						continue
					}
					πTemp003[0] = µm
					if πE = πg.CheckLocal(πF, µklass, "klass"); πE != nil {
						continue
					}
					πTemp003[1] = µklass
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µcat = πTemp004
					πF.PopCheckpoint()
					goto Label12
				Label13:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp005, πTemp006 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label14
					}
					πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
					continue
					// line 186: except AttributeError:
					πF.SetLineno(186)
				Label14:
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple1(µcategory).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("unknown warning category: %r").ToObject(), πTemp004); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_OptionError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 187: raise _OptionError("unknown warning category: %r" % (category,))
					πF.SetLineno(187)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label12
				Label12:
					goto Label5
				Label5:
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µcat, "cat"); πE != nil {
						continue
					}
					πTemp003[0] = µcat
					if πTemp004, πE = πg.ResolveGlobal(πF, ßWarning); πE != nil {
						continue
					}
					πTemp003[1] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.IsTrue(πF, πTemp008); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label15
					}
					goto Label16
					// line 188: if not issubclass(cat, Warning):
					πF.SetLineno(188)
				Label15:
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple1(µcategory).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("invalid warning category: %r").ToObject(), πTemp004); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_OptionError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 189: raise _OptionError("invalid warning category: %r" % (category,))
					πF.SetLineno(189)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					goto Label16
				Label16:
					// line 190: return cat
					πF.SetLineno(190)
					if πE = πg.CheckLocal(πF, µcat, "cat"); πE != nil {
						continue
					}
					πR = µcat
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_getcategory.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 194: def warn(message, category=None, stacklevel=1):
			πF.SetLineno(194)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "message", Def: nil}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[1] = πg.Param{Name: "category", Def: πTemp015}
			πTemp003[2] = πg.Param{Name: "stacklevel", Def: πg.NewInt(1).ToObject()}
			πTemp014 = πg.NewFunction(πg.NewCode("warn", "build/src/__python__/warnings.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmessage *πg.Object = πArgs[0]; _ = µmessage
				var µcategory *πg.Object = πArgs[1]; _ = µcategory
				var µstacklevel *πg.Object = πArgs[2]; _ = µstacklevel
				var µcaller *πg.Object = πg.UnboundLocal; _ = µcaller
				var µglobals *πg.Object = πg.UnboundLocal; _ = µglobals
				var µlineno *πg.Object = πg.UnboundLocal; _ = µlineno
				var µmodule *πg.Object = πg.UnboundLocal; _ = µmodule
				var µfilename *πg.Object = πg.UnboundLocal; _ = µfilename
				var µfnl *πg.Object = πg.UnboundLocal; _ = µfnl
				var µregistry *πg.Object = πg.UnboundLocal; _ = µregistry
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 *πg.BaseException
				_ = πTemp005
				var πTemp006 *πg.Traceback
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Dict
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 19: goto Label19
					case 6: goto Label6
					default: panic("unexpected function state")
					}
					// line 195: """Issue a warning, or maybe ignore it or raise an exception."""
					πF.SetLineno(195)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
						continue
					}
					πTemp001[0] = µmessage
					if πTemp002, πE = πg.ResolveGlobal(πF, ßWarning); πE != nil {
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
					// line 197: if isinstance(message, Warning):
					πF.SetLineno(197)
				Label1:
					// line 198: category = message.__class__
					πF.SetLineno(198)
					if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µmessage, ß__class__, nil); πE != nil {
						continue
					}
					µcategory = πTemp002
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µcategory == πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 200: if category is None:
					πF.SetLineno(200)
				Label3:
					// line 201: category = UserWarning
					πF.SetLineno(201)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßUserWarning); πE != nil {
						continue
					}
					µcategory = πTemp002
					goto Label4
				Label4:
					// line 202: assert issubclass(category, Warning)
					πF.SetLineno(202)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					πTemp001[0] = µcategory
					if πTemp002, πE = πg.ResolveGlobal(πF, ßWarning); πE != nil {
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
					if πE = πg.Assert(πF, πTemp003, nil); πE != nil {
						continue
					}
					// line 204: try:
					πF.SetLineno(204)
					πF.PushCheckpoint(6)
					// line 205: caller = sys._getframe(stacklevel)
					πF.SetLineno(205)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstacklevel, "stacklevel"); πE != nil {
						continue
					}
					πTemp001[0] = µstacklevel
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß_getframe, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcaller = πTemp002
					πF.PopCheckpoint()
					// line 210: globals = caller.f_globals
					πF.SetLineno(210)
					if πE = πg.CheckLocal(πF, µcaller, "caller"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µcaller, ßf_globals, nil); πE != nil {
						continue
					}
					µglobals = πTemp002
					// line 211: lineno = caller.f_lineno
					πF.SetLineno(211)
					if πE = πg.CheckLocal(πF, µcaller, "caller"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µcaller, ßf_lineno, nil); πE != nil {
						continue
					}
					µlineno = πTemp002
					goto Label5
				Label6:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp005, πTemp006 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label7
					}
					πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
					continue
					// line 206: except ValueError:
					πF.SetLineno(206)
				Label7:
					// line 207: globals = sys.__dict__
					πF.SetLineno(207)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__dict__, nil); πE != nil {
						continue
					}
					µglobals = πTemp003
					// line 208: lineno = 1
					πF.SetLineno(208)
					µlineno = πg.NewInt(1).ToObject()
					πF.RestoreExc(nil, nil)
					goto Label5
				Label5:
					if πE = πg.CheckLocal(πF, µglobals, "globals"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, µglobals, ß__name__.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label8
					}
					goto Label9
					// line 212: if '__name__' in globals:
					πF.SetLineno(212)
				Label8:
					// line 213: module = globals['__name__']
					πF.SetLineno(213)
					πTemp002 = ß__name__.ToObject()
					if πE = πg.CheckLocal(πF, µglobals, "globals"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µglobals, πTemp002); πE != nil {
						continue
					}
					µmodule = πTemp003
					goto Label10
				Label9:
					// line 215: module = "<string>"
					πF.SetLineno(215)
					µmodule = πg.NewStr("<string>").ToObject()
					goto Label10
				Label10:
					// line 216: filename = globals.get('__file__')
					πF.SetLineno(216)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ß__file__.ToObject()
					if πE = πg.CheckLocal(πF, µglobals, "globals"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µglobals, ßget, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µfilename = πTemp003
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µfilename); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label11
					}
					goto Label12
					// line 217: if filename:
					πF.SetLineno(217)
				Label11:
					// line 218: fnl = filename.lower()
					πF.SetLineno(218)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µfilename, ßlower, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µfnl = πTemp003
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πg.NewTuple2(πg.NewStr(".pyc").ToObject(), πg.NewStr(".pyo").ToObject()).ToObject()
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µfnl, "fnl"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µfnl, ßendswith, nil); πE != nil {
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
						goto Label14
					}
					goto Label15
					// line 219: if fnl.endswith((".pyc", ".pyo")):
					πF.SetLineno(219)
				Label14:
					// line 220: filename = filename[:-1]
					πF.SetLineno(220)
					if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp003, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µfilename, πTemp002); πE != nil {
						continue
					}
					µfilename = πTemp003
					goto Label15
				Label15:
					goto Label13
				Label12:
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µmodule, ß__main__.ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label16
					}
					goto Label17
					// line 222: if module == "__main__":
					πF.SetLineno(222)
				Label16:
					// line 223: try:
					πF.SetLineno(223)
					πF.PushCheckpoint(19)
					// line 224: filename = sys.argv[0]
					πF.SetLineno(224)
					πTemp002 = πg.NewInt(0).ToObject()
					if πTemp007, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßargv, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp008, πTemp002); πE != nil {
						continue
					}
					µfilename = πTemp003
					πF.PopCheckpoint()
					goto Label18
				Label19:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp005, πTemp006 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label20
					}
					πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
					continue
					// line 225: except AttributeError:
					πF.SetLineno(225)
				Label20:
					// line 227: filename = '__main__'
					πF.SetLineno(227)
					µfilename = ß__main__.ToObject()
					πF.RestoreExc(nil, nil)
					goto Label18
				Label18:
					goto Label17
				Label17:
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µfilename); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label21
					}
					goto Label22
					// line 228: if not filename:
					πF.SetLineno(228)
				Label21:
					// line 229: filename = module
					πF.SetLineno(229)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					µfilename = µmodule
					goto Label22
				Label22:
					goto Label13
				Label13:
					// line 230: registry = globals.setdefault("__warningregistry__", {})
					πF.SetLineno(230)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = ß__warningregistry__.ToObject()
					πTemp009 = πg.NewDict()
					πTemp002 = πTemp009.ToObject()
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µglobals, "globals"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µglobals, ßsetdefault, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µregistry = πTemp003
					// line 231: warn_explicit(message, category, filename, lineno, module, registry,
					πF.SetLineno(231)
					πTemp001 = πF.MakeArgs(7)
					if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
						continue
					}
					πTemp001[0] = µmessage
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					πTemp001[1] = µcategory
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp001[2] = µfilename
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					πTemp001[3] = µlineno
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					πTemp001[4] = µmodule
					if πE = πg.CheckLocal(πF, µregistry, "registry"); πE != nil {
						continue
					}
					πTemp001[5] = µregistry
					if πE = πg.CheckLocal(πF, µglobals, "globals"); πE != nil {
						continue
					}
					πTemp001[6] = µglobals
					if πTemp002, πE = πg.ResolveGlobal(πF, ßwarn_explicit); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßwarn.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 234: def warn_explicit(message, category, filename, lineno,
			πF.SetLineno(234)
			πTemp003 = make([]πg.Param, 7)
			πTemp003[0] = πg.Param{Name: "message", Def: nil}
			πTemp003[1] = πg.Param{Name: "category", Def: nil}
			πTemp003[2] = πg.Param{Name: "filename", Def: nil}
			πTemp003[3] = πg.Param{Name: "lineno", Def: nil}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[4] = πg.Param{Name: "module", Def: πTemp016}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[5] = πg.Param{Name: "registry", Def: πTemp016}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[6] = πg.Param{Name: "module_globals", Def: πTemp016}
			πTemp015 = πg.NewFunction(πg.NewCode("warn_explicit", "build/src/__python__/warnings.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmessage *πg.Object = πArgs[0]; _ = µmessage
				var µcategory *πg.Object = πArgs[1]; _ = µcategory
				var µfilename *πg.Object = πArgs[2]; _ = µfilename
				var µlineno *πg.Object = πArgs[3]; _ = µlineno
				var µmodule *πg.Object = πArgs[4]; _ = µmodule
				var µregistry *πg.Object = πArgs[5]; _ = µregistry
				var µmodule_globals *πg.Object = πArgs[6]; _ = µmodule_globals
				var µtext *πg.Object = πg.UnboundLocal; _ = µtext
				var µkey *πg.Object = πg.UnboundLocal; _ = µkey
				var µitem *πg.Object = πg.UnboundLocal; _ = µitem
				var µaction *πg.Object = πg.UnboundLocal; _ = µaction
				var µmsg *πg.Object = πg.UnboundLocal; _ = µmsg
				var µcat *πg.Object = πg.UnboundLocal; _ = µcat
				var µmod *πg.Object = πg.UnboundLocal; _ = µmod
				var µln *πg.Object = πg.UnboundLocal; _ = µln
				var µoncekey *πg.Object = πg.UnboundLocal; _ = µoncekey
				var µaltkey *πg.Object = πg.UnboundLocal; _ = µaltkey
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
				var πTemp006 *πg.Dict
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 bool
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 13: goto Label13
					case 14: goto Label14
					default: panic("unexpected function state")
					}
					// line 236: lineno = int(lineno)
					πF.SetLineno(236)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					πTemp001[0] = µlineno
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µlineno = πTemp003
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µmodule == πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 237: if module is None:
					πF.SetLineno(237)
				Label1:
					// line 238: module = filename or "<unknown>"
					πF.SetLineno(238)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp002 = µfilename
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					πTemp002 = πg.NewStr("<unknown>").ToObject()
				Label3:
					µmodule = πTemp002
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πTemp005, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µmodule, πTemp003); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp005, ßlower, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp005, πg.NewStr(".py").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					goto Label5
					// line 239: if module[-3:].lower() == ".py":
					πF.SetLineno(239)
				Label4:
					// line 240: module = module[:-3] # XXX What about leading pathname?
					πF.SetLineno(240)
					if πTemp003, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp003, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µmodule, πTemp002); πE != nil {
						continue
					}
					µmodule = πTemp003
					goto Label5
				Label5:
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µregistry, "registry"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µregistry == πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label6
					}
					goto Label7
					// line 241: if registry is None:
					πF.SetLineno(241)
				Label6:
					// line 242: registry = {}
					πF.SetLineno(242)
					πTemp006 = πg.NewDict()
					πTemp002 = πTemp006.ToObject()
					µregistry = πTemp002
					goto Label7
				Label7:
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
						continue
					}
					πTemp001[0] = µmessage
					if πTemp002, πE = πg.ResolveGlobal(πF, ßWarning); πE != nil {
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
						goto Label8
					}
					goto Label9
					// line 243: if isinstance(message, Warning):
					πF.SetLineno(243)
				Label8:
					// line 244: text = str(message)
					πF.SetLineno(244)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
						continue
					}
					πTemp001[0] = µmessage
					if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp003
					// line 245: category = message.__class__
					πF.SetLineno(245)
					if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µmessage, ß__class__, nil); πE != nil {
						continue
					}
					µcategory = πTemp002
					goto Label10
				Label9:
					// line 247: text = message
					πF.SetLineno(247)
					if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
						continue
					}
					µtext = µmessage
					// line 248: message = category(message)
					πF.SetLineno(248)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
						continue
					}
					πTemp001[0] = µmessage
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					if πTemp002, πE = µcategory.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µmessage = πTemp002
					goto Label10
				Label10:
					// line 249: key = (text, category, lineno)
					πF.SetLineno(249)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(µtext, µcategory, µlineno).ToObject()
					µkey = πTemp002
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp001[0] = µkey
					if πE = πg.CheckLocal(πF, µregistry, "registry"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µregistry, ßget, nil); πE != nil {
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
						goto Label11
					}
					goto Label12
					// line 251: if registry.get(key):
					πF.SetLineno(251)
				Label11:
					// line 252: return
					πF.SetLineno(252)
					πR = πg.None
					continue
					goto Label12
				Label12:
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfilters); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
						continue
					}
					πF.PushCheckpoint(14)
					πTemp004 = false
				Label13:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label15
					}
					if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
						µitem = πTemp003
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(13)            
					// line 255: action, msg, cat, mod, ln = item
					πF.SetLineno(255)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}, πg.TieTarget{Target: &πTemp010}}}, µitem); πE != nil {
						continue
					}
					µaction = πTemp003
					µmsg = πTemp005
					µcat = πTemp008
					µmod = πTemp009
					µln = πTemp010
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp008 = πg.GetBool(µmsg == πTemp009).ToObject()
					πTemp005 = πTemp008
					if πTemp011, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp011 {
						goto Label17
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[0] = µtext
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, µmsg, ßmatch, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp005 = πTemp009
				Label17:
					πTemp003 = πTemp005
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label16
					}
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					πTemp001[0] = µcategory
					if πE = πg.CheckLocal(πF, µcat, "cat"); πE != nil {
						continue
					}
					πTemp001[1] = µcat
					if πTemp005, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp003 = πTemp008
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label16
					}
					if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp008 = πg.GetBool(µmod == πTemp009).ToObject()
					πTemp005 = πTemp008
					if πTemp011, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp011 {
						goto Label18
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					πTemp001[0] = µmodule
					if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, µmod, ßmatch, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp005 = πTemp009
				Label18:
					πTemp003 = πTemp005
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label16
					}
					if πE = πg.CheckLocal(πF, µln, "ln"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Eq(πF, µln, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp008
					if πTemp011, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp011 {
						goto Label19
					}
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µln, "ln"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Eq(πF, µlineno, µln); πE != nil {
						continue
					}
					πTemp005 = πTemp008
				Label19:
					πTemp003 = πTemp005
				Label16:
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label20
					}
					goto Label21
					// line 256: if ((msg is None or msg.match(text)) and
					πF.SetLineno(256)
				Label20:
					// line 260: break
					πF.SetLineno(260)
					πTemp004 = true
					continue
					goto Label21
				Label21:
					continue
				Label14:
					if πE != nil || πR != nil {
						continue
					}
					// line 262: action = defaultaction
					πF.SetLineno(262)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßdefaultaction); πE != nil {
						continue
					}
					µaction = πTemp003
				Label15:
					if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µaction, ßignore.ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label22
					}
					goto Label23
					// line 264: if action == "ignore":
					πF.SetLineno(264)
				Label22:
					// line 265: registry[key] = 1
					πF.SetLineno(265)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µregistry, "registry"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp003 = µkey
					if πE = πg.SetItem(πF, µregistry, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 266: return
					πF.SetLineno(266)
					πR = πg.None
					continue
					goto Label23
				Label23:
					// line 270: linecache.getlines(filename, module_globals)
					πF.SetLineno(270)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp001[0] = µfilename
					if πE = πg.CheckLocal(πF, µmodule_globals, "module_globals"); πE != nil {
						continue
					}
					πTemp001[1] = µmodule_globals
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlinecache); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßgetlines, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µaction, ßerror.ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label24
					}
					goto Label25
					// line 272: if action == "error":
					πF.SetLineno(272)
				Label24:
					if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
						continue
					}
					// line 273: raise message
					πF.SetLineno(273)
					πE = πF.Raise(µmessage, nil, nil)
					continue
					goto Label25
				Label25:
					if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µaction, ßonce.ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label26
					}
					if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µaction, ßalways.ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label27
					}
					if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µaction, ßmodule.ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label28
					}
					if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µaction, ßdefault.ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label29
					}
					goto Label30
					// line 275: if action == "once":
					πF.SetLineno(275)
				Label26:
					// line 276: registry[key] = 1
					πF.SetLineno(276)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µregistry, "registry"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp003 = µkey
					if πE = πg.SetItem(πF, µregistry, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 277: oncekey = (text, category)
					πF.SetLineno(277)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µtext, µcategory).ToObject()
					µoncekey = πTemp002
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µoncekey, "oncekey"); πE != nil {
						continue
					}
					πTemp001[0] = µoncekey
					if πTemp002, πE = πg.ResolveGlobal(πF, ßonceregistry); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label32
					}
					goto Label33
					// line 278: if onceregistry.get(oncekey):
					πF.SetLineno(278)
				Label32:
					// line 279: return
					πF.SetLineno(279)
					πR = πg.None
					continue
					goto Label33
				Label33:
					// line 280: onceregistry[oncekey] = 1
					πF.SetLineno(280)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßonceregistry); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µoncekey, "oncekey"); πE != nil {
						continue
					}
					πTemp005 = µoncekey
					if πE = πg.SetItem(πF, πTemp003, πTemp005, πTemp002); πE != nil {
						continue
					}
					goto Label31
					// line 281: elif action == "always":
					πF.SetLineno(281)
				Label27:
					// line 282: pass
					πF.SetLineno(282)
					goto Label31
					// line 283: elif action == "module":
					πF.SetLineno(283)
				Label28:
					// line 284: registry[key] = 1
					πF.SetLineno(284)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µregistry, "registry"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp003 = µkey
					if πE = πg.SetItem(πF, µregistry, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 285: altkey = (text, category, 0)
					πF.SetLineno(285)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(µtext, µcategory, πg.NewInt(0).ToObject()).ToObject()
					µaltkey = πTemp002
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µaltkey, "altkey"); πE != nil {
						continue
					}
					πTemp001[0] = µaltkey
					if πE = πg.CheckLocal(πF, µregistry, "registry"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µregistry, ßget, nil); πE != nil {
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
						goto Label34
					}
					goto Label35
					// line 286: if registry.get(altkey):
					πF.SetLineno(286)
				Label34:
					// line 287: return
					πF.SetLineno(287)
					πR = πg.None
					continue
					goto Label35
				Label35:
					// line 288: registry[altkey] = 1
					πF.SetLineno(288)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µregistry, "registry"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µaltkey, "altkey"); πE != nil {
						continue
					}
					πTemp003 = µaltkey
					if πE = πg.SetItem(πF, µregistry, πTemp003, πTemp002); πE != nil {
						continue
					}
					goto Label31
					// line 289: elif action == "default":
					πF.SetLineno(289)
				Label29:
					// line 290: registry[key] = 1
					πF.SetLineno(290)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µregistry, "registry"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp003 = µkey
					if πE = πg.SetItem(πF, µregistry, πTemp003, πTemp002); πE != nil {
						continue
					}
					goto Label31
				Label30:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µaction, µitem).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Unrecognized action (%r) in warnings.filters:\n %s").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 293: raise RuntimeError(
					πF.SetLineno(293)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label31
				Label31:
					// line 297: showwarning(message, category, filename, lineno)
					πF.SetLineno(297)
					πTemp001 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
						continue
					}
					πTemp001[0] = µmessage
					if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
						continue
					}
					πTemp001[1] = µcategory
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp001[2] = µfilename
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					πTemp001[3] = µlineno
					if πTemp002, πE = πg.ResolveGlobal(πF, ßshowwarning); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßwarn_explicit.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 300: class WarningMessage(object):
			πF.SetLineno(300)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp018, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp018
			πTemp009 = πg.NewDict()
			if πTemp016, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp016); πE != nil {
				continue
			}
			_, πE = πg.NewCode("WarningMessage", "build/src/__python__/warnings.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 302: """Holds the result of a single showwarning() call."""
					πF.SetLineno(302)
					// line 304: _WARNING_DETAILS = ("message", "category", "filename", "lineno", "file",
					πF.SetLineno(304)
					πTemp001 = πg.NewTuple6(ßmessage.ToObject(), ßcategory.ToObject(), ßfilename.ToObject(), ßlineno.ToObject(), ßfile.ToObject(), ßline.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ß_WARNING_DETAILS.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 307: def __init__(self, message, category, filename, lineno, file=None,
					πF.SetLineno(307)
					πTemp002 = make([]πg.Param, 7)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "message", Def: nil}
					πTemp002[2] = πg.Param{Name: "category", Def: nil}
					πTemp002[3] = πg.Param{Name: "filename", Def: nil}
					πTemp002[4] = πg.Param{Name: "lineno", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "file", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[6] = πg.Param{Name: "line", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/warnings.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmessage *πg.Object = πArgs[1]; _ = µmessage
						var µcategory *πg.Object = πArgs[2]; _ = µcategory
						var µfilename *πg.Object = πArgs[3]; _ = µfilename
						var µlineno *πg.Object = πArgs[4]; _ = µlineno
						var µfile *πg.Object = πArgs[5]; _ = µfile
						var µline *πg.Object = πArgs[6]; _ = µline
						var µlocal_values *πg.Object = πg.UnboundLocal; _ = µlocal_values
						var µattr *πg.Object = πg.UnboundLocal; _ = µattr
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
							// line 309: local_values = locals()
							πF.SetLineno(309)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlocals); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µlocal_values = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_WARNING_DETAILS, nil); πE != nil {
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
								µattr = πTemp002
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 311: setattr(self, attr, local_values[attr])
							πF.SetLineno(311)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[0] = µself
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp005[1] = µattr
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp002 = µattr
							if πE = πg.CheckLocal(πF, µlocal_values, "local_values"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µlocal_values, πTemp002); πE != nil {
								continue
							}
							πTemp005[2] = πTemp006
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 312: self._category_name = category.__name__ if category else None
							πF.SetLineno(312)
							if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µcategory); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µcategory, "category"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcategory, ß__name__, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							goto Label5
						Label4:
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label5:
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_category_name, πTemp002); πE != nil {
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
					// line 314: def __str__(self):
					πF.SetLineno(314)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__str__", "build/src/__python__/warnings.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 315: return ("{message : %r, category : %r, filename : %r, lineno : %s, "
							πF.SetLineno(315)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmessage, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_category_name, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßfilename, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßline, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple5(πTemp003, πTemp004, πTemp005, πTemp006, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("{message : %r, category : %r, filename : %r, lineno : %s, line : %r}").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__str__.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp017, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp017 == nil {
				πTemp017 = πg.TypeType.ToObject()
			}
			if πTemp018, πE = πTemp017.Call(πF, []*πg.Object{πg.NewStr("WarningMessage").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßWarningMessage.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 320: class catch_warnings(object):
			πF.SetLineno(320)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp018, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp018
			πTemp009 = πg.NewDict()
			if πTemp016, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp016); πE != nil {
				continue
			}
			_, πE = πg.NewCode("catch_warnings", "build/src/__python__/warnings.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 322: """A context manager that copies and restores the warnings filter upon
					πF.SetLineno(322)
					// line 337: def __init__(self, record=False, module=None):
					πF.SetLineno(337)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "record", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "module", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/warnings.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µrecord *πg.Object = πArgs[1]; _ = µrecord
						var µmodule *πg.Object = πArgs[2]; _ = µmodule
						var πTemp001 *πg.Object
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
							// line 338: """Specify whether to record warnings and if an alternative module
							πF.SetLineno(338)
							// line 345: self._record = record
							πF.SetLineno(345)
							if πE = πg.CheckLocal(πF, µrecord, "record"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µrecord); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_record, πTemp001); πE != nil {
								continue
							}
							// line 346: self._module = sys.modules['warnings'] if module is None else module
							πF.SetLineno(346)
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µmodule == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label1
							}
							πTemp002 = ßwarnings.ToObject()
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßmodules, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							goto Label2
						Label1:
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							πTemp001 = µmodule
						Label2:
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_module, πTemp002); πE != nil {
								continue
							}
							// line 347: self._entered = False
							πF.SetLineno(347)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_entered, πTemp002); πE != nil {
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
					// line 349: def __repr__(self):
					πF.SetLineno(349)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__repr__", "build/src/__python__/warnings.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µargs *πg.Object = πg.UnboundLocal; _ = µargs
						var µname *πg.Object = πg.UnboundLocal; _ = µname
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
						var πTemp008 *πg.Object
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 350: args = []
							πF.SetLineno(350)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µargs = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_record, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 351: if self._record:
							πF.SetLineno(351)
						Label1:
							// line 352: args.append("record=True")
							πF.SetLineno(352)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("record=True").ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µargs, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_module, nil); πE != nil {
								continue
							}
							πTemp005 = ßwarnings.ToObject()
							if πTemp007, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßmodules, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp008, πTemp005); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004 != πTemp006).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 353: if self._module is not sys.modules['warnings']:
							πF.SetLineno(353)
						Label3:
							// line 354: args.append("module=%r" % self._module)
							πF.SetLineno(354)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_module, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("module=%r").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µargs, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label4
						Label4:
							// line 355: name = type(self).__name__
							πF.SetLineno(355)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.GetAttr(πF, πTemp004, ß__name__, nil); πE != nil {
								continue
							}
							µname = πTemp002
							// line 356: return "%s(%s)" % (name, ", ".join(args))
							πF.SetLineno(356)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp001[0] = µargs
							if πTemp005, πE = πg.GetAttr(πF, πg.NewStr(", ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp004 = πg.NewTuple2(µname, πTemp006).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s(%s)").ToObject(), πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 358: def __enter__(self):
					πF.SetLineno(358)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__enter__", "build/src/__python__/warnings.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlog *πg.Object = πg.UnboundLocal; _ = µlog
						var µshowwarning *πg.Object = πg.UnboundLocal; _ = µshowwarning
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
						var πTemp006 []πg.Param
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_entered, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 359: if self._entered:
							πF.SetLineno(359)
						Label1:
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Cannot enter %r twice").ToObject(), µself); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 360: raise RuntimeError("Cannot enter %r twice" % self)
							πF.SetLineno(360)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							// line 361: self._entered = True
							πF.SetLineno(361)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_entered, πTemp004); πE != nil {
								continue
							}
							// line 362: self._filters = self._module.filters
							πF.SetLineno(362)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_module, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßfilters, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_filters, πTemp001); πE != nil {
								continue
							}
							// line 363: self._module.filters = self._filters[:]
							πF.SetLineno(363)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_filters, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_module, nil); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp005, ßfilters, πTemp001); πE != nil {
								continue
							}
							// line 364: self._showwarning = self._module.showwarning
							πF.SetLineno(364)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_module, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßshowwarning, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_showwarning, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_record, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 365: if self._record:
							πF.SetLineno(365)
						Label3:
							// line 366: log = []
							πF.SetLineno(366)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							µlog = πTemp001
							// line 367: def showwarning(*args, **kwargs):
							πF.SetLineno(367)
							πTemp006 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("showwarning", "build/src/__python__/warnings.py", πTemp006, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µargs *πg.Object = πArgs[0]; _ = µargs
								var µkwargs *πg.Object = πArgs[1]; _ = µkwargs
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
									// line 368: log.append(WarningMessage(*args, **kwargs))
									πF.SetLineno(368)
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.ResolveGlobal(πF, ßWarningMessage); πE != nil {
										continue
									}
									if πTemp003, πE = πg.Invoke(πF, πTemp002, nil, µargs, nil, µkwargs); πE != nil {
										continue
									}
									πTemp001[0] = πTemp003
									if πE = πg.CheckLocal(πF, µlog, "log"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µlog, ßappend, nil); πE != nil {
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
							µshowwarning = πTemp001
							// line 369: self._module.showwarning = showwarning
							πF.SetLineno(369)
							if πE = πg.CheckLocal(πF, µshowwarning, "showwarning"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µshowwarning); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_module, nil); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp005, ßshowwarning, πTemp004); πE != nil {
								continue
							}
							// line 370: return log
							πF.SetLineno(370)
							if πE = πg.CheckLocal(πF, µlog, "log"); πE != nil {
								continue
							}
							πR = µlog
							continue
							goto Label5
						Label4:
							// line 372: return None
							πF.SetLineno(372)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp004
							continue
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
					if πE = πClass.SetItem(πF, ß__enter__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 374: def __exit__(self, *exc_info):
					πF.SetLineno(374)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("__exit__", "build/src/__python__/warnings.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µexc_info *πg.Object = πArgs[1]; _ = µexc_info
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_entered, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 375: if not self._entered:
							πF.SetLineno(375)
						Label1:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Cannot exit %r without entering first").ToObject(), µself); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 376: raise RuntimeError("Cannot exit %r without entering first" % self)
							πF.SetLineno(376)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label2
						Label2:
							// line 377: self._module.filters = self._filters
							πF.SetLineno(377)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_filters, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_module, nil); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp005, ßfilters, πTemp002); πE != nil {
								continue
							}
							// line 378: self._module.showwarning = self._showwarning
							πF.SetLineno(378)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_showwarning, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_module, nil); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp005, ßshowwarning, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__exit__.ToObject(), πTemp005); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp017, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp017 == nil {
				πTemp017 = πg.TypeType.ToObject()
			}
			if πTemp018, πE = πTemp017.Call(πF, []*πg.Object{πg.NewStr("catch_warnings").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßcatch_warnings.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 389: _warnings_defaults = False
			πF.SetLineno(389)
			if πTemp016, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_warnings_defaults.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 397: filters = []
			πF.SetLineno(397)
			πTemp002 = make([]*πg.Object, 0)
			πTemp016 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßfilters.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 398: defaultaction = "default"
			πF.SetLineno(398)
			if πE = πF.Globals().SetItem(πF, ßdefaultaction.ToObject(), ßdefault.ToObject()); πE != nil {
				continue
			}
			// line 399: onceregistry = {}
			πF.SetLineno(399)
			πTemp009 = πg.NewDict()
			πTemp016 = πTemp009.ToObject()
			if πE = πF.Globals().SetItem(πF, ßonceregistry.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 403: _processoptions(sys.warnoptions)
			πF.SetLineno(403)
			πTemp002 = πF.MakeArgs(1)
			if πTemp016, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp017, πE = πg.GetAttr(πF, πTemp016, ßwarnoptions, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp017
			if πTemp016, πE = πg.ResolveGlobal(πF, ß_processoptions); πE != nil {
				continue
			}
			if πTemp017, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πTemp017, πE = πg.ResolveGlobal(πF, ß_warnings_defaults); πE != nil {
				continue
			}
			if πTemp019, πE = πg.IsTrue(πF, πTemp017); πE != nil {
				continue
			}
			πTemp016 = πg.GetBool(!πTemp019).ToObject()
			if πTemp019, πE = πg.IsTrue(πF, πTemp016); πE != nil {
				continue
			}
			if πTemp019 {
				goto Label1
			}
			goto Label2
			// line 404: if not _warnings_defaults:
			πF.SetLineno(404)
		Label1:
			// line 405: silence = [ImportWarning, PendingDeprecationWarning]
			πF.SetLineno(405)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp016, πE = πg.ResolveGlobal(πF, ßImportWarning); πE != nil {
				continue
			}
			πTemp002[0] = πTemp016
			if πTemp016, πE = πg.ResolveGlobal(πF, ßPendingDeprecationWarning); πE != nil {
				continue
			}
			πTemp002[1] = πTemp016
			πTemp016 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßsilence.ToObject(), πTemp016); πE != nil {
				continue
			}
			if πTemp018, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp020, πE = πg.GetAttr(πF, πTemp018, ßpy3kwarning, nil); πE != nil {
				continue
			}
			if πTemp021, πE = πg.IsTrue(πF, πTemp020); πE != nil {
				continue
			}
			πTemp017 = πg.GetBool(!πTemp021).ToObject()
			πTemp016 = πTemp017
			if πTemp019, πE = πg.IsTrue(πF, πTemp016); πE != nil {
				continue
			}
			if !πTemp019 {
				goto Label3
			}
			if πTemp018, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp020, πE = πg.GetAttr(πF, πTemp018, ßflags, nil); πE != nil {
				continue
			}
			if πTemp018, πE = πg.GetAttr(πF, πTemp020, ßdivision_warning, nil); πE != nil {
				continue
			}
			if πTemp021, πE = πg.IsTrue(πF, πTemp018); πE != nil {
				continue
			}
			πTemp017 = πg.GetBool(!πTemp021).ToObject()
			πTemp016 = πTemp017
		Label3:
			if πTemp019, πE = πg.IsTrue(πF, πTemp016); πE != nil {
				continue
			}
			if πTemp019 {
				goto Label4
			}
			goto Label5
			// line 407: if not sys.py3kwarning and not sys.flags.division_warning:
			πF.SetLineno(407)
		Label4:
			// line 408: silence.append(DeprecationWarning)
			πF.SetLineno(408)
			πTemp002 = πF.MakeArgs(1)
			if πTemp016, πE = πg.ResolveGlobal(πF, ßDeprecationWarning); πE != nil {
				continue
			}
			πTemp002[0] = πTemp016
			if πTemp016, πE = πg.ResolveGlobal(πF, ßsilence); πE != nil {
				continue
			}
			if πTemp017, πE = πg.GetAttr(πF, πTemp016, ßappend, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp017.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			goto Label5
		Label5:
			if πTemp017, πE = πg.ResolveGlobal(πF, ßsilence); πE != nil {
				continue
			}
			if πTemp016, πE = πg.Iter(πF, πTemp017); πE != nil {
				continue
			}
			πF.PushCheckpoint(7)
			πTemp019 = false
		Label6:
			if πE != nil || πR != nil {
				continue
			}
			if πTemp019 {
				πF.PopCheckpoint()
				goto Label8
			}
			if πTemp017, πE = πg.Next(πF, πTemp016); πE != nil {
				isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
				if exc != nil {
					πE = exc
				} else if isStop {
					πE = nil
					πF.RestoreExc(nil, nil)
				}
				πTemp021 = !isStop
			} else {
				πTemp021 = true
				if πE = πF.Globals().SetItem(πF, ßcls.ToObject(), πTemp017); πE != nil {
					continue
				}
			}
			if πE != nil || !πTemp021 {
				continue
			}
			πF.PushCheckpoint(6)            
			// line 410: simplefilter("ignore", category=cls)
			πF.SetLineno(410)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = ßignore.ToObject()
			if πTemp017, πE = πg.ResolveGlobal(πF, ßcls); πE != nil {
				continue
			}
			πTemp022 = πg.KWArgs{
				{"category", πTemp017},
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßsimplefilter); πE != nil {
				continue
			}
			if πTemp018, πE = πTemp017.Call(πF, πTemp002, πTemp022); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			continue
		Label7:
			if πE != nil || πR != nil {
				continue
			}
		Label8:
			// line 411: bytes_warning = sys.flags.bytes_warning
			πF.SetLineno(411)
			if πTemp016, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp017, πE = πg.GetAttr(πF, πTemp016, ßflags, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πg.GetAttr(πF, πTemp017, ßbytes_warning, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßbytes_warning.ToObject(), πTemp016); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßbytes_warning); πE != nil {
				continue
			}
			if πTemp016, πE = πg.GT(πF, πTemp017, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			if πTemp019, πE = πg.IsTrue(πF, πTemp016); πE != nil {
				continue
			}
			if πTemp019 {
				goto Label9
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßbytes_warning); πE != nil {
				continue
			}
			if πTemp019, πE = πg.IsTrue(πF, πTemp016); πE != nil {
				continue
			}
			if πTemp019 {
				goto Label10
			}
			goto Label11
			// line 412: if bytes_warning > 1:
			πF.SetLineno(412)
		Label9:
			// line 413: bytes_action = "error"
			πF.SetLineno(413)
			if πE = πF.Globals().SetItem(πF, ßbytes_action.ToObject(), ßerror.ToObject()); πE != nil {
				continue
			}
			goto Label12
			// line 414: elif bytes_warning:
			πF.SetLineno(414)
		Label10:
			// line 415: bytes_action = "default"
			πF.SetLineno(415)
			if πE = πF.Globals().SetItem(πF, ßbytes_action.ToObject(), ßdefault.ToObject()); πE != nil {
				continue
			}
			goto Label12
		Label11:
			// line 417: bytes_action = "ignore"
			πF.SetLineno(417)
			if πE = πF.Globals().SetItem(πF, ßbytes_action.ToObject(), ßignore.ToObject()); πE != nil {
				continue
			}
			goto Label12
		Label12:
			// line 418: simplefilter(bytes_action, category=BytesWarning, append=1)
			πF.SetLineno(418)
			πTemp002 = πF.MakeArgs(1)
			if πTemp016, πE = πg.ResolveGlobal(πF, ßbytes_action); πE != nil {
				continue
			}
			πTemp002[0] = πTemp016
			if πTemp016, πE = πg.ResolveGlobal(πF, ßBytesWarning); πE != nil {
				continue
			}
			πTemp022 = πg.KWArgs{
				{"category", πTemp016},
				{"append", πg.NewInt(1).ToObject()},
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßsimplefilter); πE != nil {
				continue
			}
			if πTemp017, πE = πTemp016.Call(πF, πTemp002, πTemp022); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			goto Label2
		Label2:
			// line 419: del _warnings_defaults
			πF.SetLineno(419)
			if πE = πg.DelVar(πF, πF.Globals(), ß_warnings_defaults); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("warnings", Code)
}
