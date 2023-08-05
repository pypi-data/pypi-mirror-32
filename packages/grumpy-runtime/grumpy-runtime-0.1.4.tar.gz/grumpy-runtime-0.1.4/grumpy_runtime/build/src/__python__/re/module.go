package re
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/re.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßBRANCH := πg.InternStr("BRANCH")
		ßDEBUG := πg.InternStr("DEBUG")
		ßDOTALL := πg.InternStr("DOTALL")
		ßI := πg.InternStr("I")
		ßIGNORECASE := πg.InternStr("IGNORECASE")
		ßL := πg.InternStr("L")
		ßLOCALE := πg.InternStr("LOCALE")
		ßM := πg.InternStr("M")
		ßMULTILINE := πg.InternStr("MULTILINE")
		ßNone := πg.InternStr("None")
		ßPattern := πg.InternStr("Pattern")
		ßS := πg.InternStr("S")
		ßSRE_FLAG_DEBUG := πg.InternStr("SRE_FLAG_DEBUG")
		ßSRE_FLAG_DOTALL := πg.InternStr("SRE_FLAG_DOTALL")
		ßSRE_FLAG_IGNORECASE := πg.InternStr("SRE_FLAG_IGNORECASE")
		ßSRE_FLAG_LOCALE := πg.InternStr("SRE_FLAG_LOCALE")
		ßSRE_FLAG_MULTILINE := πg.InternStr("SRE_FLAG_MULTILINE")
		ßSRE_FLAG_TEMPLATE := πg.InternStr("SRE_FLAG_TEMPLATE")
		ßSRE_FLAG_UNICODE := πg.InternStr("SRE_FLAG_UNICODE")
		ßSRE_FLAG_VERBOSE := πg.InternStr("SRE_FLAG_VERBOSE")
		ßSUBPATTERN := πg.InternStr("SUBPATTERN")
		ßScanner := πg.InternStr("Scanner")
		ßSubPattern := πg.InternStr("SubPattern")
		ßT := πg.InternStr("T")
		ßTEMPLATE := πg.InternStr("TEMPLATE")
		ßTypeError := πg.InternStr("TypeError")
		ßU := πg.InternStr("U")
		ßUNICODE := πg.InternStr("UNICODE")
		ßVERBOSE := πg.InternStr("VERBOSE")
		ßValueError := πg.InternStr("ValueError")
		ßX := πg.InternStr("X")
		ß_MAXCACHE := πg.InternStr("_MAXCACHE")
		ß__all__ := πg.InternStr("__all__")
		ß__call__ := πg.InternStr("__call__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__version__ := πg.InternStr("__version__")
		ß_alphanum := πg.InternStr("_alphanum")
		ß_cache := πg.InternStr("_cache")
		ß_cache_repl := πg.InternStr("_cache_repl")
		ß_compile := πg.InternStr("_compile")
		ß_compile_repl := πg.InternStr("_compile_repl")
		ß_expand := πg.InternStr("_expand")
		ß_locale := πg.InternStr("_locale")
		ß_pattern_type := πg.InternStr("_pattern_type")
		ß_pickle := πg.InternStr("_pickle")
		ß_subx := πg.InternStr("_subx")
		ßabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 := πg.InternStr("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
		ßappend := πg.InternStr("append")
		ßbranch := πg.InternStr("branch")
		ßclear := πg.InternStr("clear")
		ßcompile := πg.InternStr("compile")
		ßcopy_reg := πg.InternStr("copy_reg")
		ßend := πg.InternStr("end")
		ßenumerate := πg.InternStr("enumerate")
		ßerror := πg.InternStr("error")
		ßescape := πg.InternStr("escape")
		ßexpand_template := πg.InternStr("expand_template")
		ßfindall := πg.InternStr("findall")
		ßfinditer := πg.InternStr("finditer")
		ßflags := πg.InternStr("flags")
		ßfrozenset := πg.InternStr("frozenset")
		ßget := πg.InternStr("get")
		ßgroup := πg.InternStr("group")
		ßgroups := πg.InternStr("groups")
		ßhasattr := πg.InternStr("hasattr")
		ßisinstance := πg.InternStr("isinstance")
		ßisstring := πg.InternStr("isstring")
		ßjoin := πg.InternStr("join")
		ßlastindex := πg.InternStr("lastindex")
		ßlen := πg.InternStr("len")
		ßlexicon := πg.InternStr("lexicon")
		ßlist := πg.InternStr("list")
		ßmatch := πg.InternStr("match")
		ßobject := πg.InternStr("object")
		ßparse := πg.InternStr("parse")
		ßparse_template := πg.InternStr("parse_template")
		ßpattern := πg.InternStr("pattern")
		ßpickle := πg.InternStr("pickle")
		ßpurge := πg.InternStr("purge")
		ßscan := πg.InternStr("scan")
		ßscanner := πg.InternStr("scanner")
		ßsearch := πg.InternStr("search")
		ßsplit := πg.InternStr("split")
		ßsre_compile := πg.InternStr("sre_compile")
		ßsre_parse := πg.InternStr("sre_parse")
		ßsub := πg.InternStr("sub")
		ßsubn := πg.InternStr("subn")
		ßsubpattern := πg.InternStr("subpattern")
		ßtemplate := πg.InternStr("template")
		ßtype := πg.InternStr("type")
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
		var πTemp011 *πg.Object
		_ = πTemp011
		var πTemp012 *πg.Object
		_ = πTemp012
		var πTemp013 *πg.Object
		_ = πTemp013
		var πTemp014 *πg.Dict
		_ = πTemp014
		var πTemp015 []*πg.Object
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 17: r"""Support for regular expressions (RE).
			πF.SetLineno(17)
			// line 104: import sre_compile
			πF.SetLineno(104)
			if πTemp002, πE = πg.ImportModule(πF, "sre_compile"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsre_compile.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 105: import sre_parse
			πF.SetLineno(105)
			if πTemp002, πE = πg.ImportModule(πF, "sre_parse"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsre_parse.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 111: _locale = None
			πF.SetLineno(111)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_locale.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 112: BRANCH = "branch"
			πF.SetLineno(112)
			if πE = πF.Globals().SetItem(πF, ßBRANCH.ToObject(), ßbranch.ToObject()); πE != nil {
				continue
			}
			// line 113: SUBPATTERN = "subpattern"
			πF.SetLineno(113)
			if πE = πF.Globals().SetItem(πF, ßSUBPATTERN.ToObject(), ßsubpattern.ToObject()); πE != nil {
				continue
			}
			// line 116: __all__ = [ "match", "search", "sub", "subn", "split", "findall",
			πF.SetLineno(116)
			πTemp002 = make([]*πg.Object, 23)
			πTemp002[0] = ßmatch.ToObject()
			πTemp002[1] = ßsearch.ToObject()
			πTemp002[2] = ßsub.ToObject()
			πTemp002[3] = ßsubn.ToObject()
			πTemp002[4] = ßsplit.ToObject()
			πTemp002[5] = ßfindall.ToObject()
			πTemp002[6] = ßcompile.ToObject()
			πTemp002[7] = ßpurge.ToObject()
			πTemp002[8] = ßtemplate.ToObject()
			πTemp002[9] = ßescape.ToObject()
			πTemp002[10] = ßI.ToObject()
			πTemp002[11] = ßL.ToObject()
			πTemp002[12] = ßM.ToObject()
			πTemp002[13] = ßS.ToObject()
			πTemp002[14] = ßX.ToObject()
			πTemp002[15] = ßU.ToObject()
			πTemp002[16] = ßIGNORECASE.ToObject()
			πTemp002[17] = ßLOCALE.ToObject()
			πTemp002[18] = ßMULTILINE.ToObject()
			πTemp002[19] = ßDOTALL.ToObject()
			πTemp002[20] = ßVERBOSE.ToObject()
			πTemp002[21] = ßUNICODE.ToObject()
			πTemp002[22] = ßerror.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 121: __version__ = "2.2.1"
			πF.SetLineno(121)
			if πE = πF.Globals().SetItem(πF, ß__version__.ToObject(), πg.NewStr("2.2.1").ToObject()); πE != nil {
				continue
			}
			// line 124: I = IGNORECASE = sre_compile.SRE_FLAG_IGNORECASE # ignore case
			πF.SetLineno(124)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßsre_compile); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßSRE_FLAG_IGNORECASE, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßI.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßIGNORECASE.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 125: L = LOCALE = sre_compile.SRE_FLAG_LOCALE # assume current 8-bit locale
			πF.SetLineno(125)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßsre_compile); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßSRE_FLAG_LOCALE, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßL.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLOCALE.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 126: U = UNICODE = sre_compile.SRE_FLAG_UNICODE # assume unicode locale
			πF.SetLineno(126)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßsre_compile); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßSRE_FLAG_UNICODE, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßU.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßUNICODE.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 127: M = MULTILINE = sre_compile.SRE_FLAG_MULTILINE # make anchors look for newline
			πF.SetLineno(127)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßsre_compile); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßSRE_FLAG_MULTILINE, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßM.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMULTILINE.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 128: S = DOTALL = sre_compile.SRE_FLAG_DOTALL # make dot match newline
			πF.SetLineno(128)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßsre_compile); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßSRE_FLAG_DOTALL, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßS.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDOTALL.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 129: X = VERBOSE = sre_compile.SRE_FLAG_VERBOSE # ignore whitespace and comments
			πF.SetLineno(129)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßsre_compile); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßSRE_FLAG_VERBOSE, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßX.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßVERBOSE.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 132: T = TEMPLATE = sre_compile.SRE_FLAG_TEMPLATE # disable backtracking
			πF.SetLineno(132)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßsre_compile); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßSRE_FLAG_TEMPLATE, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßT.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTEMPLATE.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 133: DEBUG = sre_compile.SRE_FLAG_DEBUG # dump pattern after compilation
			πF.SetLineno(133)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßsre_compile); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßSRE_FLAG_DEBUG, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDEBUG.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 136: error = sre_compile.error
			πF.SetLineno(136)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßsre_compile); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßerror, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßerror.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 141: def match(pattern, string, flags=0):
			πF.SetLineno(141)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "pattern", Def: nil}
			πTemp004[1] = πg.Param{Name: "string", Def: nil}
			πTemp004[2] = πg.Param{Name: "flags", Def: πg.NewInt(0).ToObject()}
			πTemp001 = πg.NewFunction(πg.NewCode("match", "build/src/__python__/re.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpattern *πg.Object = πArgs[0]; _ = µpattern
				var µstring *πg.Object = πArgs[1]; _ = µstring
				var µflags *πg.Object = πArgs[2]; _ = µflags
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
					// line 142: """Try to apply the pattern at the start of the string, returning
					πF.SetLineno(142)
					// line 144: return _compile(pattern, flags).match(string)
					πF.SetLineno(144)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					πTemp001[0] = µstring
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					πTemp002[0] = µpattern
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp002[1] = µflags
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_compile); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßmatch, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßmatch.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 146: def search(pattern, string, flags=0):
			πF.SetLineno(146)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "pattern", Def: nil}
			πTemp004[1] = πg.Param{Name: "string", Def: nil}
			πTemp004[2] = πg.Param{Name: "flags", Def: πg.NewInt(0).ToObject()}
			πTemp003 = πg.NewFunction(πg.NewCode("search", "build/src/__python__/re.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpattern *πg.Object = πArgs[0]; _ = µpattern
				var µstring *πg.Object = πArgs[1]; _ = µstring
				var µflags *πg.Object = πArgs[2]; _ = µflags
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
					// line 147: """Scan through string looking for a match to the pattern, returning
					πF.SetLineno(147)
					// line 149: return _compile(pattern, flags).search(string)
					πF.SetLineno(149)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					πTemp001[0] = µstring
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					πTemp002[0] = µpattern
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp002[1] = µflags
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_compile); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßsearch, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßsearch.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 151: def sub(pattern, repl, string, count=0, flags=0):
			πF.SetLineno(151)
			πTemp004 = make([]πg.Param, 5)
			πTemp004[0] = πg.Param{Name: "pattern", Def: nil}
			πTemp004[1] = πg.Param{Name: "repl", Def: nil}
			πTemp004[2] = πg.Param{Name: "string", Def: nil}
			πTemp004[3] = πg.Param{Name: "count", Def: πg.NewInt(0).ToObject()}
			πTemp004[4] = πg.Param{Name: "flags", Def: πg.NewInt(0).ToObject()}
			πTemp005 = πg.NewFunction(πg.NewCode("sub", "build/src/__python__/re.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpattern *πg.Object = πArgs[0]; _ = µpattern
				var µrepl *πg.Object = πArgs[1]; _ = µrepl
				var µstring *πg.Object = πArgs[2]; _ = µstring
				var µcount *πg.Object = πArgs[3]; _ = µcount
				var µflags *πg.Object = πArgs[4]; _ = µflags
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
					// line 152: """Return the string obtained by replacing the leftmost
					πF.SetLineno(152)
					// line 158: return _compile(pattern, flags).sub(repl, string, count)
					πF.SetLineno(158)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µrepl, "repl"); πE != nil {
						continue
					}
					πTemp001[0] = µrepl
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					πTemp001[1] = µstring
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					πTemp001[2] = µcount
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					πTemp002[0] = µpattern
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp002[1] = µflags
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_compile); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßsub, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßsub.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 160: def subn(pattern, repl, string, count=0, flags=0):
			πF.SetLineno(160)
			πTemp004 = make([]πg.Param, 5)
			πTemp004[0] = πg.Param{Name: "pattern", Def: nil}
			πTemp004[1] = πg.Param{Name: "repl", Def: nil}
			πTemp004[2] = πg.Param{Name: "string", Def: nil}
			πTemp004[3] = πg.Param{Name: "count", Def: πg.NewInt(0).ToObject()}
			πTemp004[4] = πg.Param{Name: "flags", Def: πg.NewInt(0).ToObject()}
			πTemp006 = πg.NewFunction(πg.NewCode("subn", "build/src/__python__/re.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpattern *πg.Object = πArgs[0]; _ = µpattern
				var µrepl *πg.Object = πArgs[1]; _ = µrepl
				var µstring *πg.Object = πArgs[2]; _ = µstring
				var µcount *πg.Object = πArgs[3]; _ = µcount
				var µflags *πg.Object = πArgs[4]; _ = µflags
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
					// line 161: """Return a 2-tuple containing (new_string, number).
					πF.SetLineno(161)
					// line 169: return _compile(pattern, flags).subn(repl, string, count)
					πF.SetLineno(169)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µrepl, "repl"); πE != nil {
						continue
					}
					πTemp001[0] = µrepl
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					πTemp001[1] = µstring
					if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
						continue
					}
					πTemp001[2] = µcount
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					πTemp002[0] = µpattern
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp002[1] = µflags
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_compile); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßsubn, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßsubn.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 171: def split(pattern, string, maxsplit=0, flags=0):
			πF.SetLineno(171)
			πTemp004 = make([]πg.Param, 4)
			πTemp004[0] = πg.Param{Name: "pattern", Def: nil}
			πTemp004[1] = πg.Param{Name: "string", Def: nil}
			πTemp004[2] = πg.Param{Name: "maxsplit", Def: πg.NewInt(0).ToObject()}
			πTemp004[3] = πg.Param{Name: "flags", Def: πg.NewInt(0).ToObject()}
			πTemp007 = πg.NewFunction(πg.NewCode("split", "build/src/__python__/re.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpattern *πg.Object = πArgs[0]; _ = µpattern
				var µstring *πg.Object = πArgs[1]; _ = µstring
				var µmaxsplit *πg.Object = πArgs[2]; _ = µmaxsplit
				var µflags *πg.Object = πArgs[3]; _ = µflags
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
					// line 172: """Split the source string by the occurrences of the pattern,
					πF.SetLineno(172)
					// line 174: return _compile(pattern, flags).split(string, maxsplit)
					πF.SetLineno(174)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					πTemp001[0] = µstring
					if πE = πg.CheckLocal(πF, µmaxsplit, "maxsplit"); πE != nil {
						continue
					}
					πTemp001[1] = µmaxsplit
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					πTemp002[0] = µpattern
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp002[1] = µflags
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_compile); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßsplit, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßsplit.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 176: def findall(pattern, string, flags=0):
			πF.SetLineno(176)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "pattern", Def: nil}
			πTemp004[1] = πg.Param{Name: "string", Def: nil}
			πTemp004[2] = πg.Param{Name: "flags", Def: πg.NewInt(0).ToObject()}
			πTemp008 = πg.NewFunction(πg.NewCode("findall", "build/src/__python__/re.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpattern *πg.Object = πArgs[0]; _ = µpattern
				var µstring *πg.Object = πArgs[1]; _ = µstring
				var µflags *πg.Object = πArgs[2]; _ = µflags
				var µfinditer *πg.Object = πg.UnboundLocal; _ = µfinditer
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []πg.Param
				_ = πTemp005
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 177: """Return a list of all non-overlapping matches in the string.
					πF.SetLineno(177)
					// line 184: return _compile(pattern, flags).findall(string)
					πF.SetLineno(184)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					πTemp001[0] = µstring
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					πTemp002[0] = µpattern
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp002[1] = µflags
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_compile); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßfindall, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp004
					continue
					// line 188: def finditer(pattern, string, flags=0):
					πF.SetLineno(188)
					πTemp005 = make([]πg.Param, 3)
					πTemp005[0] = πg.Param{Name: "pattern", Def: nil}
					πTemp005[1] = πg.Param{Name: "string", Def: nil}
					πTemp005[2] = πg.Param{Name: "flags", Def: πg.NewInt(0).ToObject()}
					πTemp003 = πg.NewFunction(πg.NewCode("finditer", "build/src/__python__/re.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µpattern *πg.Object = πArgs[0]; _ = µpattern
						var µstring *πg.Object = πArgs[1]; _ = µstring
						var µflags *πg.Object = πArgs[2]; _ = µflags
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
							// line 189: """Return an iterator over all non-overlapping matches in the
							πF.SetLineno(189)
							// line 193: return _compile(pattern, flags).finditer(string)
							πF.SetLineno(193)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							πTemp001[0] = µstring
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
								continue
							}
							πTemp002[0] = µpattern
							if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
								continue
							}
							πTemp002[1] = µflags
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_compile); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßfinditer, nil); πE != nil {
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
					µfinditer = πTemp003
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßfindall.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 195: def compile(pattern, flags=0):
			πF.SetLineno(195)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pattern", Def: nil}
			πTemp004[1] = πg.Param{Name: "flags", Def: πg.NewInt(0).ToObject()}
			πTemp009 = πg.NewFunction(πg.NewCode("compile", "build/src/__python__/re.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpattern *πg.Object = πArgs[0]; _ = µpattern
				var µflags *πg.Object = πArgs[1]; _ = µflags
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
					// line 196: "Compile a regular expression pattern, returning a pattern object."
					πF.SetLineno(196)
					// line 197: return _compile(pattern, flags)
					πF.SetLineno(197)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					πTemp001[0] = µpattern
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp001[1] = µflags
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_compile); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßcompile.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 199: def purge():
			πF.SetLineno(199)
			πTemp004 = make([]πg.Param, 0)
			πTemp010 = πg.NewFunction(πg.NewCode("purge", "build/src/__python__/re.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 200: "Clear the regular expression cache"
					πF.SetLineno(200)
					// line 201: _cache.clear()
					πF.SetLineno(201)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_cache); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßclear, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 202: _cache_repl.clear()
					πF.SetLineno(202)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_cache_repl); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßclear, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßpurge.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 204: def template(pattern, flags=0):
			πF.SetLineno(204)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pattern", Def: nil}
			πTemp004[1] = πg.Param{Name: "flags", Def: πg.NewInt(0).ToObject()}
			πTemp011 = πg.NewFunction(πg.NewCode("template", "build/src/__python__/re.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpattern *πg.Object = πArgs[0]; _ = µpattern
				var µflags *πg.Object = πArgs[1]; _ = µflags
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
					// line 205: "Compile a template pattern, returning a pattern object"
					πF.SetLineno(205)
					// line 206: return _compile(pattern, flags|T)
					πF.SetLineno(206)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					πTemp001[0] = µpattern
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßT); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Or(πF, µflags, πTemp003); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_compile); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßtemplate.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 208: _alphanum = frozenset(
			πF.SetLineno(208)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = ßabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.ToObject()
			if πTemp012, πE = πg.ResolveGlobal(πF, ßfrozenset); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp012.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ß_alphanum.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 211: def escape(pattern):
			πF.SetLineno(211)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "pattern", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("escape", "build/src/__python__/re.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpattern *πg.Object = πArgs[0]; _ = µpattern
				var µs *πg.Object = πg.UnboundLocal; _ = µs
				var µalphanum *πg.Object = πg.UnboundLocal; _ = µalphanum
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µc *πg.Object = πg.UnboundLocal; _ = µc
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
					// line 212: "Escape all non-alphanumeric characters in pattern."
					πF.SetLineno(212)
					// line 213: s = list(pattern)
					πF.SetLineno(213)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					πTemp001[0] = µpattern
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µs = πTemp003
					// line 214: alphanum = _alphanum
					πF.SetLineno(214)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_alphanum); πE != nil {
						continue
					}
					µalphanum = πTemp002
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					πTemp001[0] = µpattern
					if πTemp003, πE = πg.ResolveGlobal(πF, ßenumerate); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
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
					if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
							continue
						}
						µi = πTemp004
						µc = πTemp007
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(1)            
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µalphanum, "alphanum"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Contains(πF, µalphanum, µc); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label4
					}
					goto Label5
					// line 216: if c not in alphanum:
					πF.SetLineno(216)
				Label4:
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µc, πg.NewStr("\x00").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label6
					}
					goto Label7
					// line 217: if c == "\000":
					πF.SetLineno(217)
				Label6:
					// line 218: s[i] = "\\000"
					πF.SetLineno(218)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\\000").ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp004 = µi
					if πE = πg.SetItem(πF, µs, πTemp004, πTemp003); πE != nil {
						continue
					}
					goto Label8
				Label7:
					// line 220: s[i] = "\\" + c
					πF.SetLineno(220)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πg.NewStr("\\").ToObject(), µc); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp007 = µi
					if πE = πg.SetItem(πF, µs, πTemp007, πTemp004); πE != nil {
						continue
					}
					goto Label8
				Label8:
					goto Label5
				Label5:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 221: return pattern[:0].join(s)
					πF.SetLineno(221)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(0).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µpattern, πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßjoin, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßescape.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 226: _cache = {}
			πF.SetLineno(226)
			πTemp014 = πg.NewDict()
			πTemp013 = πTemp014.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_cache.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 227: _cache_repl = {}
			πF.SetLineno(227)
			πTemp014 = πg.NewDict()
			πTemp013 = πTemp014.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_cache_repl.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 229: _pattern_type = type(sre_compile.compile("", 0))
			πF.SetLineno(229)
			πTemp002 = πF.MakeArgs(1)
			πTemp015 = πF.MakeArgs(2)
			πTemp015[0] = ß.ToObject()
			πTemp015[1] = πg.NewInt(0).ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßsre_compile); πE != nil {
				continue
			}
			if πTemp016, πE = πg.GetAttr(πF, πTemp013, ßcompile, nil); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp016.Call(πF, πTemp015, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp015)
			πTemp002[0] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp013.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ß_pattern_type.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 231: _MAXCACHE = 100
			πF.SetLineno(231)
			if πE = πF.Globals().SetItem(πF, ß_MAXCACHE.ToObject(), πg.NewInt(100).ToObject()); πE != nil {
				continue
			}
			// line 233: def _compile(*key):
			πF.SetLineno(233)
			πTemp004 = make([]πg.Param, 0)
			πTemp013 = πg.NewFunction(πg.NewCode("_compile", "build/src/__python__/re.py", πTemp004, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µkey *πg.Object = πArgs[0]; _ = µkey
				var µpattern *πg.Object = πg.UnboundLocal; _ = µpattern
				var µflags *πg.Object = πg.UnboundLocal; _ = µflags
				var µbypass_cache *πg.Object = πg.UnboundLocal; _ = µbypass_cache
				var µcachekey *πg.Object = πg.UnboundLocal; _ = µcachekey
				var µp *πg.Object = πg.UnboundLocal; _ = µp
				var µv *πg.Object = πg.UnboundLocal; _ = µv
				var µloc *πg.Object = πg.UnboundLocal; _ = µloc
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
					case 10: goto Label10
					default: panic("unexpected function state")
					}
					// line 235: pattern, flags = key
					πF.SetLineno(235)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp002}}}, µkey); πE != nil {
						continue
					}
					µpattern = πTemp001
					µflags = πTemp002
					// line 236: bypass_cache = flags & DEBUG
					πF.SetLineno(236)
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßDEBUG); πE != nil {
						continue
					}
					if πTemp001, πE = πg.And(πF, µflags, πTemp002); πE != nil {
						continue
					}
					µbypass_cache = πTemp001
					if πE = πg.CheckLocal(πF, µbypass_cache, "bypass_cache"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µbypass_cache); πE != nil {
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
					// line 237: if not bypass_cache:
					πF.SetLineno(237)
				Label1:
					// line 238: cachekey = (type(key[0]),) + key
					πF.SetLineno(238)
					πTemp004 = πF.MakeArgs(1)
					πTemp005 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µkey, πTemp005); πE != nil {
						continue
					}
					πTemp004[0] = πTemp006
					if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp002 = πg.NewTuple1(πTemp006).ToObject()
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp002, µkey); πE != nil {
						continue
					}
					µcachekey = πTemp001
					goto Label2
				Label2:
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					πTemp004[0] = µpattern
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_pattern_type); πE != nil {
						continue
					}
					πTemp004[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label3
					}
					goto Label4
					// line 245: if isinstance(pattern, _pattern_type):
					πF.SetLineno(245)
				Label3:
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µflags); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label5
					}
					goto Label6
					// line 246: if flags:
					πF.SetLineno(246)
				Label5:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("Cannot process flags argument with a compiled pattern").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 247: raise ValueError('Cannot process flags argument with a compiled pattern')
					πF.SetLineno(247)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label6
				Label6:
					// line 248: return pattern
					πF.SetLineno(248)
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					πR = µpattern
					continue
					goto Label4
				Label4:
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					πTemp004[0] = µpattern
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsre_compile); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßisstring, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp003).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label7
					}
					goto Label8
					// line 249: if not sre_compile.isstring(pattern):
					πF.SetLineno(249)
				Label7:
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					// line 250: raise TypeError, "first argument must be string or compiled pattern"
					πF.SetLineno(250)
					πE = πF.Raise(πTemp001, πg.NewStr("first argument must be string or compiled pattern").ToObject(), nil)
					continue
					goto Label8
				Label8:
					// line 251: try:
					πF.SetLineno(251)
					πF.PushCheckpoint(10)
					// line 252: p = sre_compile.compile(pattern, flags)
					πF.SetLineno(252)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					πTemp004[0] = µpattern
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp004[1] = µflags
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsre_compile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µp = πTemp001
					πF.PopCheckpoint()
					goto Label9
				Label10:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label11
					}
					πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
					continue
					// line 253: except error, v:
					πF.SetLineno(253)
				Label11:
					µv = πTemp007.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					// line 254: raise error, v # invalid expression
					πF.SetLineno(254)
					πE = πF.Raise(πTemp001, µv, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label9
				Label9:
					if πE = πg.CheckLocal(πF, µbypass_cache, "bypass_cache"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µbypass_cache); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp003).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label12
					}
					goto Label13
					// line 255: if not bypass_cache:
					πF.SetLineno(255)
				Label12:
					πTemp004 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_cache); πE != nil {
						continue
					}
					πTemp004[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_MAXCACHE); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GE(πF, πTemp005, πTemp002); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label14
					}
					goto Label15
					// line 256: if len(_cache) >= _MAXCACHE:
					πF.SetLineno(256)
				Label14:
					// line 257: _cache.clear()
					πF.SetLineno(257)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_cache); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßclear, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					goto Label15
				Label15:
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µp, ßflags, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßLOCALE); πE != nil {
						continue
					}
					if πTemp001, πE = πg.And(πF, πTemp002, πTemp005); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label16
					}
					goto Label17
					// line 258: if p.flags & LOCALE:
					πF.SetLineno(258)
				Label16:
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_locale); πE != nil {
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
						goto Label19
					}
					goto Label20
					// line 259: if not _locale:
					πF.SetLineno(259)
				Label19:
					// line 260: return p
					πF.SetLineno(260)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					πR = µp
					continue
					goto Label20
				Label20:
					goto Label18
				Label17:
					// line 263: loc = None
					πF.SetLineno(263)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µloc = πTemp001
					goto Label18
				Label18:
					// line 264: _cache[cachekey] = p, loc
					πF.SetLineno(264)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µloc, "loc"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µp, µloc).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_cache); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcachekey, "cachekey"); πE != nil {
						continue
					}
					πTemp006 = µcachekey
					if πE = πg.SetItem(πF, πTemp005, πTemp006, πTemp002); πE != nil {
						continue
					}
					goto Label13
				Label13:
					// line 265: return p
					πF.SetLineno(265)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					πR = µp
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_compile.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 267: def _compile_repl(*key):
			πF.SetLineno(267)
			πTemp004 = make([]πg.Param, 0)
			πTemp016 = πg.NewFunction(πg.NewCode("_compile_repl", "build/src/__python__/re.py", πTemp004, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µkey *πg.Object = πArgs[0]; _ = µkey
				var µp *πg.Object = πg.UnboundLocal; _ = µp
				var µrepl *πg.Object = πg.UnboundLocal; _ = µrepl
				var µpattern *πg.Object = πg.UnboundLocal; _ = µpattern
				var µv *πg.Object = πg.UnboundLocal; _ = µv
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					// line 269: p = _cache_repl.get(key)
					πF.SetLineno(269)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp001[0] = µkey
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_cache_repl); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µp = πTemp002
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µp != πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 270: if p is not None:
					πF.SetLineno(270)
				Label1:
					// line 271: return p
					πF.SetLineno(271)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					πR = µp
					continue
					goto Label2
				Label2:
					// line 272: repl, pattern = key
					πF.SetLineno(272)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}}}, µkey); πE != nil {
						continue
					}
					µrepl = πTemp002
					µpattern = πTemp003
					// line 273: try:
					πF.SetLineno(273)
					πF.PushCheckpoint(4)
					// line 274: p = sre_parse.parse_template(repl, pattern)
					πF.SetLineno(274)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µrepl, "repl"); πE != nil {
						continue
					}
					πTemp001[0] = µrepl
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					πTemp001[1] = µpattern
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsre_parse); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßparse_template, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µp = πTemp002
					πF.PopCheckpoint()
					goto Label3
				Label4:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp005, πTemp006 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label5
					}
					πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
					continue
					// line 275: except error, v:
					πF.SetLineno(275)
				Label5:
					µv = πTemp005.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					// line 276: raise error, v # invalid expression
					πF.SetLineno(276)
					πE = πF.Raise(πTemp002, µv, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label3
				Label3:
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_cache_repl); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_MAXCACHE); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GE(πF, πTemp007, πTemp003); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label6
					}
					goto Label7
					// line 277: if len(_cache_repl) >= _MAXCACHE:
					πF.SetLineno(277)
				Label6:
					// line 278: _cache_repl.clear()
					πF.SetLineno(278)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_cache_repl); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßclear, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					goto Label7
				Label7:
					// line 279: _cache_repl[key] = p
					πF.SetLineno(279)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µp); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_cache_repl); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp007 = µkey
					if πE = πg.SetItem(πF, πTemp003, πTemp007, πTemp002); πE != nil {
						continue
					}
					// line 280: return p
					πF.SetLineno(280)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					πR = µp
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_compile_repl.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 282: def _expand(pattern, match, template):
			πF.SetLineno(282)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "pattern", Def: nil}
			πTemp004[1] = πg.Param{Name: "match", Def: nil}
			πTemp004[2] = πg.Param{Name: "template", Def: nil}
			πTemp017 = πg.NewFunction(πg.NewCode("_expand", "build/src/__python__/re.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpattern *πg.Object = πArgs[0]; _ = µpattern
				var µmatch *πg.Object = πArgs[1]; _ = µmatch
				var µtemplate *πg.Object = πArgs[2]; _ = µtemplate
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
					// line 284: template = sre_parse.parse_template(template, pattern)
					πF.SetLineno(284)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µtemplate, "template"); πE != nil {
						continue
					}
					πTemp001[0] = µtemplate
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					πTemp001[1] = µpattern
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsre_parse); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßparse_template, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtemplate = πTemp002
					// line 285: return sre_parse.expand_template(template, match)
					πF.SetLineno(285)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µtemplate, "template"); πE != nil {
						continue
					}
					πTemp001[0] = µtemplate
					if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
						continue
					}
					πTemp001[1] = µmatch
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsre_parse); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßexpand_template, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_expand.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 287: def _subx(pattern, template):
			πF.SetLineno(287)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "pattern", Def: nil}
			πTemp004[1] = πg.Param{Name: "template", Def: nil}
			πTemp018 = πg.NewFunction(πg.NewCode("_subx", "build/src/__python__/re.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpattern *πg.Object = πArgs[0]; _ = µpattern
				var µtemplate *πg.Object = πArgs[1]; _ = µtemplate
				var µfilter *πg.Object = πg.UnboundLocal; _ = µfilter
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
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 []πg.Param
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 289: template = _compile_repl(template, pattern)
					πF.SetLineno(289)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µtemplate, "template"); πE != nil {
						continue
					}
					πTemp001[0] = µtemplate
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					πTemp001[1] = µpattern
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_compile_repl); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtemplate = πTemp003
					πTemp005 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µtemplate, "template"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µtemplate, πTemp005); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp007).ToObject()
					πTemp002 = πTemp003
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label1
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp005 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µtemplate, "template"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µtemplate, πTemp005); πE != nil {
						continue
					}
					πTemp001[0] = πTemp006
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.Eq(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label1:
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label2
					}
					goto Label3
					// line 290: if not template[0] and len(template[1]) == 1:
					πF.SetLineno(290)
				Label2:
					// line 292: return template[1][0]
					πF.SetLineno(292)
					πTemp002 = πg.NewInt(0).ToObject()
					πTemp005 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µtemplate, "template"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µtemplate, πTemp005); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
						continue
					}
					πR = πTemp003
					continue
					goto Label3
				Label3:
					// line 293: def filter(match, template=template):
					πF.SetLineno(293)
					πTemp008 = make([]πg.Param, 2)
					πTemp008[0] = πg.Param{Name: "match", Def: nil}
					if πE = πg.CheckLocal(πF, µtemplate, "template"); πE != nil {
						continue
					}
					πTemp008[1] = πg.Param{Name: "template", Def: µtemplate}
					πTemp002 = πg.NewFunction(πg.NewCode("filter", "build/src/__python__/re.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µmatch *πg.Object = πArgs[0]; _ = µmatch
						var µtemplate *πg.Object = πArgs[1]; _ = µtemplate
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
							// line 294: return sre_parse.expand_template(template, match)
							πF.SetLineno(294)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtemplate, "template"); πE != nil {
								continue
							}
							πTemp001[0] = µtemplate
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							πTemp001[1] = µmatch
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsre_parse); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßexpand_template, nil); πE != nil {
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
					µfilter = πTemp002
					// line 295: return filter
					πF.SetLineno(295)
					if πE = πg.CheckLocal(πF, µfilter, "filter"); πE != nil {
						continue
					}
					πR = µfilter
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_subx.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 299: import copy_reg
			πF.SetLineno(299)
			if πTemp002, πE = πg.ImportModule(πF, "copy_reg"); πE != nil {
				continue
			}
			πTemp019 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßcopy_reg.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 301: def _pickle(p):
			πF.SetLineno(301)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "p", Def: nil}
			πTemp019 = πg.NewFunction(πg.NewCode("_pickle", "build/src/__python__/re.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µp *πg.Object = πArgs[0]; _ = µp
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 302: return _compile, (p.pattern, p.flags)
					πF.SetLineno(302)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_compile); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µp, ßpattern, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µp, ßflags, nil); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
					πTemp001 = πg.NewTuple2(πTemp002, πTemp003).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ß_pickle.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 304: copy_reg.pickle(_pattern_type, _pickle, _compile)
			πF.SetLineno(304)
			πTemp002 = πF.MakeArgs(3)
			if πTemp020, πE = πg.ResolveGlobal(πF, ß_pattern_type); πE != nil {
				continue
			}
			πTemp002[0] = πTemp020
			if πTemp020, πE = πg.ResolveGlobal(πF, ß_pickle); πE != nil {
				continue
			}
			πTemp002[1] = πTemp020
			if πTemp020, πE = πg.ResolveGlobal(πF, ß_compile); πE != nil {
				continue
			}
			πTemp002[2] = πTemp020
			if πTemp020, πE = πg.ResolveGlobal(πF, ßcopy_reg); πE != nil {
				continue
			}
			if πTemp021, πE = πg.GetAttr(πF, πTemp020, ßpickle, nil); πE != nil {
				continue
			}
			if πTemp020, πE = πTemp021.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 309: class Scanner(object):
			πF.SetLineno(309)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp022, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp022
			πTemp014 = πg.NewDict()
			if πTemp020, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp014.SetItem(πF, ß__module__.ToObject(), πTemp020); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Scanner", "build/src/__python__/re.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp014
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
					// line 310: def __init__(self, lexicon, flags=0):
					πF.SetLineno(310)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "lexicon", Def: nil}
					πTemp002[2] = πg.Param{Name: "flags", Def: πg.NewInt(0).ToObject()}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/re.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlexicon *πg.Object = πArgs[1]; _ = µlexicon
						var µflags *πg.Object = πArgs[2]; _ = µflags
						var µp *πg.Object = πg.UnboundLocal; _ = µp
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µphrase *πg.Object = πg.UnboundLocal; _ = µphrase
						var µaction *πg.Object = πg.UnboundLocal; _ = µaction
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 []*πg.Object
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πTemp013 *πg.Object
						_ = πTemp013
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 311: self.lexicon = lexicon
							πF.SetLineno(311)
							if πE = πg.CheckLocal(πF, µlexicon, "lexicon"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µlexicon); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlexicon, πTemp001); πE != nil {
								continue
							}
							// line 313: p = []
							πF.SetLineno(313)
							πTemp002 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							µp = πTemp001
							// line 314: s = sre_parse.Pattern()
							πF.SetLineno(314)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsre_parse); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßPattern, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µs = πTemp001
							// line 315: s.flags = flags
							πF.SetLineno(315)
							if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µflags); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µs, ßflags, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlexicon, "lexicon"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µlexicon); πE != nil {
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
							if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
									continue
								}
								µphrase = πTemp006
								µaction = πTemp007
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 317: p.append(sre_parse.SubPattern(s, [
							πF.SetLineno(317)
							πTemp002 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp008[0] = µs
							πTemp009 = make([]*πg.Object, 1)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßSUBPATTERN); πE != nil {
								continue
							}
							πTemp011 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp011[0] = µp
							if πTemp012, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp012.Call(πF, πTemp011, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp011)
							if πTemp010, πE = πg.Add(πF, πTemp013, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp011 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µphrase, "phrase"); πE != nil {
								continue
							}
							πTemp011[0] = µphrase
							if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
								continue
							}
							πTemp011[1] = µflags
							if πTemp012, πE = πg.ResolveGlobal(πF, ßsre_parse); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßparse, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp013.Call(πF, πTemp011, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp011)
							πTemp007 = πg.NewTuple2(πTemp010, πTemp012).ToObject()
							πTemp003 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
							πTemp009[0] = πTemp003
							πTemp003 = πg.NewList(πTemp009...).ToObject()
							πTemp008[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsre_parse); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßSubPattern, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp006.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µp, ßappend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 320: s.groups = len(p)+1
							πF.SetLineno(320)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp002[0] = µp
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
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µs, ßgroups, πTemp003); πE != nil {
								continue
							}
							// line 321: p = sre_parse.SubPattern(s, [(BRANCH, (None, p))])
							πF.SetLineno(321)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp002[0] = µs
							πTemp008 = make([]*πg.Object, 1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßBRANCH); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp006 = πg.NewTuple2(πTemp007, µp).ToObject()
							πTemp001 = πg.NewTuple2(πTemp003, πTemp006).ToObject()
							πTemp008[0] = πTemp001
							πTemp001 = πg.NewList(πTemp008...).ToObject()
							πTemp002[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsre_parse); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßSubPattern, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µp = πTemp001
							// line 322: self.scanner = sre_compile.compile(p)
							πF.SetLineno(322)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp002[0] = µp
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsre_compile); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßcompile, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßscanner, πTemp003); πE != nil {
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
					// line 323: def scan(self, string):
					πF.SetLineno(323)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "string", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("scan", "build/src/__python__/re.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstring *πg.Object = πArgs[1]; _ = µstring
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
						var µappend *πg.Object = πg.UnboundLocal; _ = µappend
						var µmatch *πg.Object = πg.UnboundLocal; _ = µmatch
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µm *πg.Object = πg.UnboundLocal; _ = µm
						var µj *πg.Object = πg.UnboundLocal; _ = µj
						var µaction *πg.Object = πg.UnboundLocal; _ = µaction
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
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 324: result = []
							πF.SetLineno(324)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µresult = πTemp002
							// line 325: append = result.append
							πF.SetLineno(325)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
								continue
							}
							µappend = πTemp002
							// line 326: match = self.scanner.scanner(string).match
							πF.SetLineno(326)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							πTemp001[0] = µstring
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßscanner, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßscanner, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmatch, nil); πE != nil {
								continue
							}
							µmatch = πTemp003
							// line 327: i = 0
							πF.SetLineno(327)
							µi = πg.NewInt(0).ToObject()
							// line 328: while 1:
							πF.SetLineno(328)
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
							if πTemp005, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 329: m = match()
							πF.SetLineno(329)
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							if πTemp002, πE = µmatch.Call(πF, nil, nil); πE != nil {
								continue
							}
							µm = πTemp002
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µm); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 330: if not m:
							πF.SetLineno(330)
						Label4:
							// line 331: break
							πF.SetLineno(331)
							πTemp004 = true
							continue
							goto Label5
						Label5:
							// line 332: j = m.end()
							πF.SetLineno(332)
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µm, ßend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µj = πTemp003
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µi, µj); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label6
							}
							goto Label7
							// line 333: if i == j:
							πF.SetLineno(333)
						Label6:
							// line 334: break
							πF.SetLineno(334)
							πTemp004 = true
							continue
							goto Label7
						Label7:
							// line 335: action = self.lexicon[m.lastindex-1][1]
							πF.SetLineno(335)
							πTemp002 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µm, ßlastindex, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Sub(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp006 = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßlexicon, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp008, πTemp006); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
								continue
							}
							µaction = πTemp003
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
								continue
							}
							πTemp001[0] = µaction
							πTemp001[1] = ß__call__.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
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
								goto Label8
							}
							goto Label9
							// line 336: if hasattr(action, '__call__'):
							πF.SetLineno(336)
						Label8:
							// line 337: self.match = m
							πF.SetLineno(337)
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µm); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmatch, πTemp002); πE != nil {
								continue
							}
							// line 338: action = action(self, m.group())
							πF.SetLineno(338)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µm, ßgroup, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
								continue
							}
							if πTemp002, πE = µaction.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µaction = πTemp002
							goto Label9
						Label9:
							if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µaction != πTemp003).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label10
							}
							goto Label11
							// line 339: if action is not None:
							πF.SetLineno(339)
						Label10:
							// line 340: append(action)
							πF.SetLineno(340)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
								continue
							}
							πTemp001[0] = µaction
							if πE = πg.CheckLocal(πF, µappend, "append"); πE != nil {
								continue
							}
							if πTemp002, πE = µappend.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label11
						Label11:
							// line 341: i = j
							πF.SetLineno(341)
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							µi = µj
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 342: return result, string[i:]
							πF.SetLineno(342)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{µi, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µstring, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µresult, πTemp006).ToObject()
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
					if πE = πClass.SetItem(πF, ßscan.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp021, πE = πTemp014.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp021 == nil {
				πTemp021 = πg.TypeType.ToObject()
			}
			if πTemp022, πE = πTemp021.Call(πF, []*πg.Object{πg.NewStr("Scanner").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp014.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßScanner.ToObject(), πTemp022); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("re", Code)
}
