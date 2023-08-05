package textwrap
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/textwrap.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßFalse := πg.InternStr("False")
		ßMULTILINE := πg.InternStr("MULTILINE")
		ßNameError := πg.InternStr("NameError")
		ßNone := πg.InternStr("None")
		ßTextWrapper := πg.InternStr("TextWrapper")
		ßTrue := πg.InternStr("True")
		ßU := πg.InternStr("U")
		ßValueError := πg.InternStr("ValueError")
		ß__all__ := πg.InternStr("__all__")
		ß__init__ := πg.InternStr("__init__")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__revision__ := πg.InternStr("__revision__")
		ß_fix_sentence_endings := πg.InternStr("_fix_sentence_endings")
		ß_handle_long_word := πg.InternStr("_handle_long_word")
		ß_leading_whitespace_re := πg.InternStr("_leading_whitespace_re")
		ß_munge_whitespace := πg.InternStr("_munge_whitespace")
		ß_split := πg.InternStr("_split")
		ß_unicode := πg.InternStr("_unicode")
		ß_whitespace := πg.InternStr("_whitespace")
		ß_whitespace_only_re := πg.InternStr("_whitespace_only_re")
		ß_wrap_chunks := πg.InternStr("_wrap_chunks")
		ßappend := πg.InternStr("append")
		ßbreak_long_words := πg.InternStr("break_long_words")
		ßbreak_on_hyphens := πg.InternStr("break_on_hyphens")
		ßcompile := πg.InternStr("compile")
		ßdedent := πg.InternStr("dedent")
		ßdrop_whitespace := πg.InternStr("drop_whitespace")
		ßenumerate := πg.InternStr("enumerate")
		ßexpand_tabs := πg.InternStr("expand_tabs")
		ßfill := πg.InternStr("fill")
		ßfindall := πg.InternStr("findall")
		ßfix_sentence_endings := πg.InternStr("fix_sentence_endings")
		ßinitial_indent := πg.InternStr("initial_indent")
		ßisinstance := πg.InternStr("isinstance")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßlowercase := πg.InternStr("lowercase")
		ßmap := πg.InternStr("map")
		ßobject := πg.InternStr("object")
		ßord := πg.InternStr("ord")
		ßpattern := πg.InternStr("pattern")
		ßpop := πg.InternStr("pop")
		ßre := πg.InternStr("re")
		ßreplace_whitespace := πg.InternStr("replace_whitespace")
		ßreverse := πg.InternStr("reverse")
		ßsearch := πg.InternStr("search")
		ßsentence_end_re := πg.InternStr("sentence_end_re")
		ßsplit := πg.InternStr("split")
		ßstartswith := πg.InternStr("startswith")
		ßstring := πg.InternStr("string")
		ßstrip := πg.InternStr("strip")
		ßsub := πg.InternStr("sub")
		ßsubsequent_indent := πg.InternStr("subsequent_indent")
		ßunicode := πg.InternStr("unicode")
		ßunicode_whitespace_trans := πg.InternStr("unicode_whitespace_trans")
		ßuspace := πg.InternStr("uspace")
		ßwhitespace_trans := πg.InternStr("whitespace_trans")
		ßwidth := πg.InternStr("width")
		ßwordsep_re := πg.InternStr("wordsep_re")
		ßwordsep_re_uni := πg.InternStr("wordsep_re_uni")
		ßwordsep_simple_re := πg.InternStr("wordsep_simple_re")
		ßwordsep_simple_re_uni := πg.InternStr("wordsep_simple_re_uni")
		ßwrap := πg.InternStr("wrap")
		ßx := πg.InternStr("x")
		ßzip := πg.InternStr("zip")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.BaseException
		_ = πTemp003
		var πTemp004 *πg.Traceback
		_ = πTemp004
		var πTemp005 bool
		_ = πTemp005
		var πTemp006 *πg.Dict
		_ = πTemp006
		var πTemp007 *πg.Object
		_ = πTemp007
		var πTemp008 *πg.Object
		_ = πTemp008
		var πTemp009 []πg.Param
		_ = πTemp009
		var πTemp010 *πg.Object
		_ = πTemp010
		var πTemp011 *πg.Object
		_ = πTemp011
		var πTemp012 []*πg.Object
		_ = πTemp012
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 2: goto Label2
			default: panic("unexpected function state")
			}
			// line 1: """Text wrapping and filling.
			πF.SetLineno(1)
			// line 8: __revision__ = "$Id$"
			πF.SetLineno(8)
			if πE = πF.Globals().SetItem(πF, ß__revision__.ToObject(), πg.NewStr("$Id$").ToObject()); πE != nil {
				continue
			}
			// line 10: import string, re
			πF.SetLineno(10)
			if πTemp002, πE = πg.ImportModule(πF, "string"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßstring.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 12: try:
			πF.SetLineno(12)
			πF.PushCheckpoint(2)
			// line 13: _unicode = unicode
			πF.SetLineno(13)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_unicode.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label1
		Label2:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp003, πTemp004 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNameError); πE != nil {
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
			// line 14: except NameError:
			πF.SetLineno(14)
		Label3:
			// line 17: class _unicode(object):
			πF.SetLineno(17)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_unicode", "build/src/__python__/textwrap.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 18: pass
					πF.SetLineno(18)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp007, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp007 == nil {
				πTemp007 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp007.Call(πF, []*πg.Object{πg.NewStr("_unicode").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_unicode.ToObject(), πTemp008); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label1
		Label1:
			// line 28: __all__ = ['TextWrapper', 'wrap', 'fill', 'dedent']
			πF.SetLineno(28)
			πTemp002 = make([]*πg.Object, 4)
			πTemp002[0] = ßTextWrapper.ToObject()
			πTemp002[1] = ßwrap.ToObject()
			πTemp002[2] = ßfill.ToObject()
			πTemp002[3] = ßdedent.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 38: _whitespace = '\t\n\x0b\x0c\r '
			πF.SetLineno(38)
			if πE = πF.Globals().SetItem(πF, ß_whitespace.ToObject(), πg.NewStr("\t\n\x0b\x0c\r ").ToObject()); πE != nil {
				continue
			}
			// line 40: class TextWrapper(object):
			πF.SetLineno(40)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TextWrapper", "build/src/__python__/textwrap.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Dict
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
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 []πg.Param
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
					// line 41: """
					πF.SetLineno(41)
					// line 83: whitespace_trans = '\x00\x01\x02\x03\x04\x05\x06\x07\x08     \x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f !"#$%&\'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~\x7f\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8a\x8b\x8c\x8d\x8e\x8f\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99\x9a\x9b\x9c\x9d\x9e\x9f\xa0\xa1\xa2\xa3\xa4\xa5\xa6\xa7\xa8\xa9\xaa\xab\xac\xad\xae\xaf\xb0\xb1\xb2\xb3\xb4\xb5\xb6\xb7\xb8\xb9\xba\xbb\xbc\xbd\xbe\xbf\xc0\xc1\xc2\xc3\xc4\xc5\xc6\xc7\xc8\xc9\xca\xcb\xcc\xcd\xce\xcf\xd0\xd1\xd2\xd3\xd4\xd5\xd6\xd7\xd8\xd9\xda\xdb\xdc\xdd\xde\xdf\xe0\xe1\xe2\xe3\xe4\xe5\xe6\xe7\xe8\xe9\xea\xeb\xec\xed\xee\xef\xf0\xf1\xf2\xf3\xf4\xf5\xf6\xf7\xf8\xf9\xfa\xfb\xfc\xfd\xfe\xff'
					πF.SetLineno(83)
					if πE = πClass.SetItem(πF, ßwhitespace_trans.ToObject(), πg.NewStr("\x00\x01\x02\x03\x04\x05\x06\x07\x08     \x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~\x7f\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8a\x8b\x8c\x8d\x8e\x8f\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99\x9a\x9b\x9c\x9d\x9e\x9f\xa0\xa1\xa2\xa3\xa4\xa5\xa6\xa7\xa8\xa9\xaa\xab\xac\xad\xae\xaf\xb0\xb1\xb2\xb3\xb4\xb5\xb6\xb7\xb8\xb9\xba\xbb\xbc\xbd\xbe\xbf\xc0\xc1\xc2\xc3\xc4\xc5\xc6\xc7\xc8\xc9\xca\xcb\xcc\xcd\xce\xcf\xd0\xd1\xd2\xd3\xd4\xd5\xd6\xd7\xd8\xd9\xda\xdb\xdc\xdd\xde\xdf\xe0\xe1\xe2\xe3\xe4\xe5\xe6\xe7\xe8\xe9\xea\xeb\xec\xed\xee\xef\xf0\xf1\xf2\xf3\xf4\xf5\xf6\xf7\xf8\xf9\xfa\xfb\xfc\xfd\xfe\xff").ToObject()); πE != nil {
						continue
					}
					// line 85: unicode_whitespace_trans = {}
					πF.SetLineno(85)
					πTemp001 = πg.NewDict()
					πTemp002 = πTemp001.ToObject()
					if πE = πClass.SetItem(πF, ßunicode_whitespace_trans.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 86: uspace = ord(u' ')
					πF.SetLineno(86)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewUnicode(" ").ToObject()
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßord); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ßuspace.ToObject(), πTemp004); πE != nil {
						continue
					}
					πTemp003 = πF.MakeArgs(2)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßord); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß_whitespace); πE != nil {
						continue
					}
					πTemp003[1] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßmap); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
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
						if πE = πClass.SetItem(πF, ßx.ToObject(), πTemp004); πE != nil {
							continue
						}
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 88: unicode_whitespace_trans[x] = uspace
					πF.SetLineno(88)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßuspace); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßunicode_whitespace_trans); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßx); πE != nil {
						continue
					}
					πTemp009 = πTemp010
					if πE = πg.SetItem(πF, πTemp008, πTemp009, πTemp005); πE != nil {
						continue
					}
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 96: wordsep_re = re.compile(
					πF.SetLineno(96)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("(\\s+|[^\\s\\w]*\\w+[^0-9\\W]-(?=\\w+[^0-9\\W])|(?<=[\\w\\!\\\"\\'\\&\\.\\,\\?])-{2,}(?=\\w))").ToObject()
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ßwordsep_re.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 105: wordsep_simple_re = re.compile(r'(\s+)')
					πF.SetLineno(105)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("(\\s+)").ToObject()
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ßwordsep_simple_re.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 109: sentence_end_re = re.compile(r'[%s]'              # lowercase letter
					πF.SetLineno(109)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßstring); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßlowercase, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("[%s][\\.\\!\\?][\\\"\\']?\\Z").ToObject(), πTemp005); πE != nil {
						continue
					}
					πTemp003[0] = πTemp002
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ßsentence_end_re.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 116: def __init__(self,
					πF.SetLineno(116)
					πTemp011 = make([]πg.Param, 10)
					πTemp011[0] = πg.Param{Name: "self", Def: nil}
					πTemp011[1] = πg.Param{Name: "width", Def: πg.NewInt(70).ToObject()}
					πTemp011[2] = πg.Param{Name: "initial_indent", Def: ß.ToObject()}
					πTemp011[3] = πg.Param{Name: "subsequent_indent", Def: ß.ToObject()}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp011[4] = πg.Param{Name: "expand_tabs", Def: πTemp004}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp011[5] = πg.Param{Name: "replace_whitespace", Def: πTemp004}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp011[6] = πg.Param{Name: "fix_sentence_endings", Def: πTemp004}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp011[7] = πg.Param{Name: "break_long_words", Def: πTemp004}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp011[8] = πg.Param{Name: "drop_whitespace", Def: πTemp004}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp011[9] = πg.Param{Name: "break_on_hyphens", Def: πTemp004}
					πTemp002 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/textwrap.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µwidth *πg.Object = πArgs[1]; _ = µwidth
						var µinitial_indent *πg.Object = πArgs[2]; _ = µinitial_indent
						var µsubsequent_indent *πg.Object = πArgs[3]; _ = µsubsequent_indent
						var µexpand_tabs *πg.Object = πArgs[4]; _ = µexpand_tabs
						var µreplace_whitespace *πg.Object = πArgs[5]; _ = µreplace_whitespace
						var µfix_sentence_endings *πg.Object = πArgs[6]; _ = µfix_sentence_endings
						var µbreak_long_words *πg.Object = πArgs[7]; _ = µbreak_long_words
						var µdrop_whitespace *πg.Object = πArgs[8]; _ = µdrop_whitespace
						var µbreak_on_hyphens *πg.Object = πArgs[9]; _ = µbreak_on_hyphens
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
							// line 126: self.width = width
							πF.SetLineno(126)
							if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µwidth); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßwidth, πTemp001); πE != nil {
								continue
							}
							// line 127: self.initial_indent = initial_indent
							πF.SetLineno(127)
							if πE = πg.CheckLocal(πF, µinitial_indent, "initial_indent"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µinitial_indent); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßinitial_indent, πTemp001); πE != nil {
								continue
							}
							// line 128: self.subsequent_indent = subsequent_indent
							πF.SetLineno(128)
							if πE = πg.CheckLocal(πF, µsubsequent_indent, "subsequent_indent"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µsubsequent_indent); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsubsequent_indent, πTemp001); πE != nil {
								continue
							}
							// line 129: self.expand_tabs = expand_tabs
							πF.SetLineno(129)
							if πE = πg.CheckLocal(πF, µexpand_tabs, "expand_tabs"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µexpand_tabs); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßexpand_tabs, πTemp001); πE != nil {
								continue
							}
							// line 130: self.replace_whitespace = replace_whitespace
							πF.SetLineno(130)
							if πE = πg.CheckLocal(πF, µreplace_whitespace, "replace_whitespace"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µreplace_whitespace); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßreplace_whitespace, πTemp001); πE != nil {
								continue
							}
							// line 131: self.fix_sentence_endings = fix_sentence_endings
							πF.SetLineno(131)
							if πE = πg.CheckLocal(πF, µfix_sentence_endings, "fix_sentence_endings"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µfix_sentence_endings); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfix_sentence_endings, πTemp001); πE != nil {
								continue
							}
							// line 132: self.break_long_words = break_long_words
							πF.SetLineno(132)
							if πE = πg.CheckLocal(πF, µbreak_long_words, "break_long_words"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µbreak_long_words); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbreak_long_words, πTemp001); πE != nil {
								continue
							}
							// line 133: self.drop_whitespace = drop_whitespace
							πF.SetLineno(133)
							if πE = πg.CheckLocal(πF, µdrop_whitespace, "drop_whitespace"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdrop_whitespace); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdrop_whitespace, πTemp001); πE != nil {
								continue
							}
							// line 134: self.break_on_hyphens = break_on_hyphens
							πF.SetLineno(134)
							if πE = πg.CheckLocal(πF, µbreak_on_hyphens, "break_on_hyphens"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µbreak_on_hyphens); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbreak_on_hyphens, πTemp001); πE != nil {
								continue
							}
							// line 139: self.wordsep_re_uni = re.compile(self.wordsep_re.pattern, re.U)
							πF.SetLineno(139)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwordsep_re, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpattern, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßU, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
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
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßwordsep_re_uni, πTemp003); πE != nil {
								continue
							}
							// line 140: self.wordsep_simple_re_uni = re.compile(
							πF.SetLineno(140)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwordsep_simple_re, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpattern, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßU, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
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
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßwordsep_simple_re_uni, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 147: def _munge_whitespace(self, text):
					πF.SetLineno(147)
					πTemp011 = make([]πg.Param, 2)
					πTemp011[0] = πg.Param{Name: "self", Def: nil}
					πTemp011[1] = πg.Param{Name: "text", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("_munge_whitespace", "build/src/__python__/textwrap.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtext *πg.Object = πArgs[1]; _ = µtext
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
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
							// line 148: """_munge_whitespace(text : string) -> string
							πF.SetLineno(148)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßexpand_tabs, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 154: if self.expand_tabs:
							πF.SetLineno(154)
						Label1:
							// line 156: text = ' '.join((' '.join(text.split('\n'))).split('\t'))
							πF.SetLineno(156)
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("\t").ToObject()
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr("\n").ToObject()
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtext, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp007
							if πTemp001, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.GetAttr(πF, πTemp007, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp007
							if πTemp001, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µtext = πTemp007
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreplace_whitespace, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 157: if self.replace_whitespace:
							πF.SetLineno(157)
						Label3:
							// line 162: text = ' '.join(' '.join(text.split('\n')).split('\t'))
							πF.SetLineno(162)
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("\t").ToObject()
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr("\n").ToObject()
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtext, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp007
							if πTemp001, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.GetAttr(πF, πTemp007, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp007
							if πTemp001, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µtext = πTemp007
							goto Label4
						Label4:
							// line 163: return text
							πF.SetLineno(163)
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
					if πE = πClass.SetItem(πF, ß_munge_whitespace.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 166: def _split(self, text):
					πF.SetLineno(166)
					πTemp011 = make([]πg.Param, 2)
					πTemp011[0] = πg.Param{Name: "self", Def: nil}
					πTemp011[1] = πg.Param{Name: "text", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("_split", "build/src/__python__/textwrap.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtext *πg.Object = πArgs[1]; _ = µtext
						var µpat *πg.Object = πg.UnboundLocal; _ = µpat
						var µchunks *πg.Object = πg.UnboundLocal; _ = µchunks
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 []πg.Param
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
							// line 167: """_split(text : string) -> [string]
							πF.SetLineno(167)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp001[0] = µtext
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_unicode); πE != nil {
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbreak_on_hyphens, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							goto Label3
							// line 181: if isinstance(text, _unicode):
							πF.SetLineno(181)
						Label1:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbreak_on_hyphens, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							goto Label6
							// line 182: if self.break_on_hyphens:
							πF.SetLineno(182)
						Label5:
							// line 183: pat = self.wordsep_re_uni
							πF.SetLineno(183)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßwordsep_re_uni, nil); πE != nil {
								continue
							}
							µpat = πTemp002
							goto Label7
						Label6:
							// line 185: pat = self.wordsep_simple_re_uni
							πF.SetLineno(185)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßwordsep_simple_re_uni, nil); πE != nil {
								continue
							}
							µpat = πTemp002
							goto Label7
						Label7:
							goto Label4
							// line 187: if self.break_on_hyphens:
							πF.SetLineno(187)
						Label2:
							// line 188: pat = self.wordsep_re
							πF.SetLineno(188)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßwordsep_re, nil); πE != nil {
								continue
							}
							µpat = πTemp002
							goto Label4
						Label3:
							// line 190: pat = self.wordsep_simple_re
							πF.SetLineno(190)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßwordsep_simple_re, nil); πE != nil {
								continue
							}
							µpat = πTemp002
							goto Label4
						Label4:
							// line 191: chunks = pat.split(text)
							πF.SetLineno(191)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp001[0] = µtext
							if πE = πg.CheckLocal(πF, µpat, "pat"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µpat, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µchunks = πTemp003
							// line 193: chunks = [x for x in chunks if x is not None]
							πF.SetLineno(193)
							πTemp005 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/textwrap.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µx *πg.Object = πg.UnboundLocal; _ = µx
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
										case 6: goto Label6
										default: panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µchunks); πE != nil {
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
											µx = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)            
										if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
											continue
										}
										if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
											continue
										}
										πTemp004 = πg.GetBool(µx != πTemp005).ToObject()
										if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
											continue
										}
										if πTemp003 {
											goto Label4
										}
										goto Label5
										// line 193: chunks = [x for x in chunks if x is not None]
										πF.SetLineno(193)
									Label4:
										// line 193: chunks = [x for x in chunks if x is not None]
										πF.SetLineno(193)
										if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
											continue
										}
										πF.PushCheckpoint(6)
										return µx, nil
									Label6:
										πTemp004 = πSent
										goto Label5
									Label5:
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
							if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp006}, nil); πE != nil {
								continue
							}
							µchunks = πTemp002
							// line 194: return chunks
							πF.SetLineno(194)
							if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
								continue
							}
							πR = µchunks
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_split.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 196: def _fix_sentence_endings(self, chunks):
					πF.SetLineno(196)
					πTemp011 = make([]πg.Param, 2)
					πTemp011[0] = πg.Param{Name: "self", Def: nil}
					πTemp011[1] = πg.Param{Name: "chunks", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("_fix_sentence_endings", "build/src/__python__/textwrap.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µchunks *πg.Object = πArgs[1]; _ = µchunks
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µpatsearch *πg.Object = πg.UnboundLocal; _ = µpatsearch
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
							// line 197: """_fix_sentence_endings(chunks : [string])
							πF.SetLineno(197)
							// line 205: i = 0
							πF.SetLineno(205)
							µi = πg.NewInt(0).ToObject()
							// line 206: patsearch = self.sentence_end_re.search
							πF.SetLineno(206)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsentence_end_re, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsearch, nil); πE != nil {
								continue
							}
							µpatsearch = πTemp002
							// line 207: while i < len(chunks)-1:
							πF.SetLineno(207)
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
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
								continue
							}
							πTemp005[0] = µchunks
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πg.Sub(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µi, πTemp002); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp006 = πTemp007
							if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µchunks, πTemp006); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp007, πg.NewStr(" ").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label4
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp002 = µi
							if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µchunks, πTemp002); πE != nil {
								continue
							}
							πTemp005[0] = πTemp006
							if πE = πg.CheckLocal(πF, µpatsearch, "patsearch"); πE != nil {
								continue
							}
							if πTemp002, πE = µpatsearch.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001 = πTemp002
						Label4:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							goto Label6
							// line 208: if chunks[i+1] == " " and patsearch(chunks[i]):
							πF.SetLineno(208)
						Label5:
							// line 209: chunks[i+1] = "  "
							πF.SetLineno(209)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewStr("  ").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp006
							if πE = πg.SetItem(πF, µchunks, πTemp002, πTemp001); πE != nil {
								continue
							}
							// line 210: i += 2
							πF.SetLineno(210)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µi, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							µi = πTemp001
							goto Label7
						Label6:
							// line 212: i += 1
							πF.SetLineno(212)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µi = πTemp001
							goto Label7
						Label7:
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
					if πE = πClass.SetItem(πF, ß_fix_sentence_endings.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 214: def _handle_long_word(self, reversed_chunks, cur_line, cur_len, width):
					πF.SetLineno(214)
					πTemp011 = make([]πg.Param, 5)
					πTemp011[0] = πg.Param{Name: "self", Def: nil}
					πTemp011[1] = πg.Param{Name: "reversed_chunks", Def: nil}
					πTemp011[2] = πg.Param{Name: "cur_line", Def: nil}
					πTemp011[3] = πg.Param{Name: "cur_len", Def: nil}
					πTemp011[4] = πg.Param{Name: "width", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("_handle_long_word", "build/src/__python__/textwrap.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µreversed_chunks *πg.Object = πArgs[1]; _ = µreversed_chunks
						var µcur_line *πg.Object = πArgs[2]; _ = µcur_line
						var µcur_len *πg.Object = πArgs[3]; _ = µcur_len
						var µwidth *πg.Object = πArgs[4]; _ = µwidth
						var µspace_left *πg.Object = πg.UnboundLocal; _ = µspace_left
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 215: """_handle_long_word(chunks : [string],
							πF.SetLineno(215)
							if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µwidth, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 224: if width < 1:
							πF.SetLineno(224)
						Label1:
							// line 225: space_left = 1
							πF.SetLineno(225)
							µspace_left = πg.NewInt(1).ToObject()
							goto Label3
						Label2:
							// line 227: space_left = width - cur_len
							πF.SetLineno(227)
							if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcur_len, "cur_len"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µwidth, µcur_len); πE != nil {
								continue
							}
							µspace_left = πTemp001
							goto Label3
						Label3:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbreak_long_words, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µcur_line, "cur_line"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µcur_line); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label5
							}
							goto Label6
							// line 231: if self.break_long_words:
							πF.SetLineno(231)
						Label4:
							// line 232: cur_line.append(reversed_chunks[-1][:space_left])
							πF.SetLineno(232)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µspace_left, "space_left"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µspace_left, πg.None}, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005 = πTemp006
							if πE = πg.CheckLocal(πF, µreversed_chunks, "reversed_chunks"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µreversed_chunks, πTemp005); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πE = πg.CheckLocal(πF, µcur_line, "cur_line"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcur_line, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 233: reversed_chunks[-1] = reversed_chunks[-1][space_left:]
							πF.SetLineno(233)
							if πE = πg.CheckLocal(πF, µspace_left, "space_left"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µspace_left, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005 = πTemp006
							if πE = πg.CheckLocal(πF, µreversed_chunks, "reversed_chunks"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µreversed_chunks, πTemp005); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µreversed_chunks, "reversed_chunks"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005 = πTemp006
							if πE = πg.SetItem(πF, µreversed_chunks, πTemp005, πTemp001); πE != nil {
								continue
							}
							goto Label6
							// line 238: elif not cur_line:
							πF.SetLineno(238)
						Label5:
							// line 239: cur_line.append(reversed_chunks.pop())
							πF.SetLineno(239)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µreversed_chunks, "reversed_chunks"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µreversed_chunks, ßpop, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πE = πg.CheckLocal(πF, µcur_line, "cur_line"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcur_line, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label6
						Label6:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_handle_long_word.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 247: def _wrap_chunks(self, chunks):
					πF.SetLineno(247)
					πTemp011 = make([]πg.Param, 2)
					πTemp011[0] = πg.Param{Name: "self", Def: nil}
					πTemp011[1] = πg.Param{Name: "chunks", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("_wrap_chunks", "build/src/__python__/textwrap.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µchunks *πg.Object = πArgs[1]; _ = µchunks
						var µlines *πg.Object = πg.UnboundLocal; _ = µlines
						var µcur_line *πg.Object = πg.UnboundLocal; _ = µcur_line
						var µcur_len *πg.Object = πg.UnboundLocal; _ = µcur_len
						var µindent *πg.Object = πg.UnboundLocal; _ = µindent
						var µwidth *πg.Object = πg.UnboundLocal; _ = µwidth
						var µl *πg.Object = πg.UnboundLocal; _ = µl
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
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 12: goto Label12
							case 3: goto Label3
							case 4: goto Label4
							case 13: goto Label13
							default: panic("unexpected function state")
							}
							// line 248: """_wrap_chunks(chunks : [string]) -> [string]
							πF.SetLineno(248)
							// line 260: lines = []
							πF.SetLineno(260)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µlines = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßwidth, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LE(πF, πTemp003, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 261: if self.width <= 0:
							πF.SetLineno(261)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßwidth, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("invalid width %r (must be > 0)").ToObject(), πTemp003); πE != nil {
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
							// line 262: raise ValueError("invalid width %r (must be > 0)" % self.width)
							πF.SetLineno(262)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							// line 266: chunks.reverse()
							πF.SetLineno(266)
							if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µchunks, ßreverse, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 268: while chunks:
							πF.SetLineno(268)
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
							if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µchunks); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 272: cur_line = []
							πF.SetLineno(272)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µcur_line = πTemp002
							// line 273: cur_len = 0
							πF.SetLineno(273)
							µcur_len = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µlines); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label6
							}
							goto Label7
							// line 276: if lines:
							πF.SetLineno(276)
						Label6:
							// line 277: indent = self.subsequent_indent
							πF.SetLineno(277)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsubsequent_indent, nil); πE != nil {
								continue
							}
							µindent = πTemp002
							goto Label8
						Label7:
							// line 279: indent = self.initial_indent
							πF.SetLineno(279)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßinitial_indent, nil); πE != nil {
								continue
							}
							µindent = πTemp002
							goto Label8
						Label8:
							// line 282: width = self.width - len(indent)
							πF.SetLineno(282)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßwidth, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							πTemp001[0] = µindent
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Sub(πF, πTemp003, πTemp007); πE != nil {
								continue
							}
							µwidth = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdrop_whitespace, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label9
							}
							if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp006 = πTemp007
							if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µchunks, πTemp006); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp007, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp007, ß.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label9
							}
							if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
								continue
							}
							πTemp002 = µlines
						Label9:
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label10
							}
							goto Label11
							// line 286: if self.drop_whitespace and chunks[-1].strip() == '' and lines:
							πF.SetLineno(286)
						Label10:
							// line 288: chunks.pop()
							πF.SetLineno(288)
							if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µchunks, ßpop, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label11
						Label11:
							// line 290: while chunks:
							πF.SetLineno(290)
							πF.PushCheckpoint(13)
							πTemp005 = false
						Label12:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label14
							}
							if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, µchunks); πE != nil {
								continue
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(12)            
							// line 291: l = len(chunks[-1])
							πF.SetLineno(291)
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µchunks, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µl = πTemp003
							if πE = πg.CheckLocal(πF, µcur_len, "cur_len"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µcur_len, µl); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LE(πF, πTemp003, µwidth); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label15
							}
							goto Label16
							// line 294: if cur_len + l <= width:
							πF.SetLineno(294)
						Label15:
							// line 295: cur_line.append(chunks.pop())
							πF.SetLineno(295)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µchunks, ßpop, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µcur_line, "cur_line"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcur_line, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 296: cur_len += l
							πF.SetLineno(296)
							if πE = πg.CheckLocal(πF, µcur_len, "cur_len"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µcur_len, µl); πE != nil {
								continue
							}
							µcur_len = πTemp002
							goto Label17
						Label16:
							// line 300: break
							πF.SetLineno(300)
							πTemp005 = true
							continue
							goto Label17
						Label17:
							continue
						Label13:
							if πE != nil || πR != nil {
								continue
							}
						Label14:
							if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
								continue
							}
							πTemp002 = µchunks
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label18
							}
							πTemp001 = πF.MakeArgs(1)
							if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp006 = πTemp007
							if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µchunks, πTemp006); πE != nil {
								continue
							}
							πTemp001[0] = πTemp007
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, πTemp007, µwidth); πE != nil {
								continue
							}
							πTemp002 = πTemp003
						Label18:
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label19
							}
							goto Label20
							// line 304: if chunks and len(chunks[-1]) > width:
							πF.SetLineno(304)
						Label19:
							// line 305: self._handle_long_word(chunks, cur_line, cur_len, width)
							πF.SetLineno(305)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
								continue
							}
							πTemp001[0] = µchunks
							if πE = πg.CheckLocal(πF, µcur_line, "cur_line"); πE != nil {
								continue
							}
							πTemp001[1] = µcur_line
							if πE = πg.CheckLocal(πF, µcur_len, "cur_len"); πE != nil {
								continue
							}
							πTemp001[2] = µcur_len
							if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
								continue
							}
							πTemp001[3] = µwidth
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_handle_long_word, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label20
						Label20:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdrop_whitespace, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label21
							}
							if πE = πg.CheckLocal(πF, µcur_line, "cur_line"); πE != nil {
								continue
							}
							πTemp002 = µcur_line
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label21
							}
							if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp006 = πTemp007
							if πE = πg.CheckLocal(πF, µcur_line, "cur_line"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µcur_line, πTemp006); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp007, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp007, ß.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
						Label21:
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label22
							}
							goto Label23
							// line 308: if self.drop_whitespace and cur_line and cur_line[-1].strip() == '':
							πF.SetLineno(308)
						Label22:
							// line 310: cur_line.pop()
							πF.SetLineno(310)
							if πE = πg.CheckLocal(πF, µcur_line, "cur_line"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcur_line, ßpop, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label23
						Label23:
							if πE = πg.CheckLocal(πF, µcur_line, "cur_line"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µcur_line); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label24
							}
							goto Label25
							// line 314: if cur_line:
							πF.SetLineno(314)
						Label24:
							// line 315: lines.append(indent + ''.join(cur_line))
							πF.SetLineno(315)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcur_line, "cur_line"); πE != nil {
								continue
							}
							πTemp009[0] = µcur_line
							if πTemp003, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp002, πE = πg.Add(πF, µindent, πTemp006); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlines, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label25
						Label25:
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 317: return lines
							πF.SetLineno(317)
							if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
								continue
							}
							πR = µlines
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_wrap_chunks.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 322: def wrap(self, text):
					πF.SetLineno(322)
					πTemp011 = make([]πg.Param, 2)
					πTemp011[0] = πg.Param{Name: "self", Def: nil}
					πTemp011[1] = πg.Param{Name: "text", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("wrap", "build/src/__python__/textwrap.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtext *πg.Object = πArgs[1]; _ = µtext
						var µchunks *πg.Object = πg.UnboundLocal; _ = µchunks
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
							// line 323: """wrap(text : string) -> [string]
							πF.SetLineno(323)
							// line 331: text = self._munge_whitespace(text)
							πF.SetLineno(331)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp001[0] = µtext
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_munge_whitespace, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtext = πTemp003
							// line 332: chunks = self._split(text)
							πF.SetLineno(332)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp001[0] = µtext
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_split, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µchunks = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfix_sentence_endings, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 333: if self.fix_sentence_endings:
							πF.SetLineno(333)
						Label1:
							// line 334: self._fix_sentence_endings(chunks)
							πF.SetLineno(334)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
								continue
							}
							πTemp001[0] = µchunks
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_fix_sentence_endings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label2
						Label2:
							// line 335: return self._wrap_chunks(chunks)
							πF.SetLineno(335)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
								continue
							}
							πTemp001[0] = µchunks
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_wrap_chunks, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßwrap.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 337: def fill(self, text):
					πF.SetLineno(337)
					πTemp011 = make([]πg.Param, 2)
					πTemp011[0] = πg.Param{Name: "self", Def: nil}
					πTemp011[1] = πg.Param{Name: "text", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("fill", "build/src/__python__/textwrap.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtext *πg.Object = πArgs[1]; _ = µtext
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
							// line 338: """fill(text : string) -> string
							πF.SetLineno(338)
							// line 344: return "\n".join(self.wrap(text))
							πF.SetLineno(344)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp002[0] = µtext
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßwrap, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßfill.ToObject(), πTemp013); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp007, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp007 == nil {
				πTemp007 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp007.Call(πF, []*πg.Object{πg.NewStr("TextWrapper").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTextWrapper.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 349: def wrap(text, width=70, **kwargs):
			πF.SetLineno(349)
			πTemp009 = make([]πg.Param, 2)
			πTemp009[0] = πg.Param{Name: "text", Def: nil}
			πTemp009[1] = πg.Param{Name: "width", Def: πg.NewInt(70).ToObject()}
			πTemp001 = πg.NewFunction(πg.NewCode("wrap", "build/src/__python__/textwrap.py", πTemp009, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtext *πg.Object = πArgs[0]; _ = µtext
				var µwidth *πg.Object = πArgs[1]; _ = µwidth
				var µkwargs *πg.Object = πArgs[2]; _ = µkwargs
				var µw *πg.Object = πg.UnboundLocal; _ = µw
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
					// line 350: """Wrap a single paragraph of text, returning a list of wrapped lines.
					πF.SetLineno(350)
					// line 359: w = TextWrapper(width=width, **kwargs)
					πF.SetLineno(359)
					if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
						continue
					}
					πTemp001 = πg.KWArgs{
						{"width", µwidth},
					}
					if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTextWrapper); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp002, nil, nil, πTemp001, µkwargs); πE != nil {
						continue
					}
					µw = πTemp003
					// line 360: return w.wrap(text)
					πF.SetLineno(360)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp004[0] = µtext
					if πE = πg.CheckLocal(πF, µw, "w"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µw, ßwrap, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
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
			if πE = πF.Globals().SetItem(πF, ßwrap.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 362: def fill(text, width=70, **kwargs):
			πF.SetLineno(362)
			πTemp009 = make([]πg.Param, 2)
			πTemp009[0] = πg.Param{Name: "text", Def: nil}
			πTemp009[1] = πg.Param{Name: "width", Def: πg.NewInt(70).ToObject()}
			πTemp007 = πg.NewFunction(πg.NewCode("fill", "build/src/__python__/textwrap.py", πTemp009, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtext *πg.Object = πArgs[0]; _ = µtext
				var µwidth *πg.Object = πArgs[1]; _ = µwidth
				var µkwargs *πg.Object = πArgs[2]; _ = µkwargs
				var µw *πg.Object = πg.UnboundLocal; _ = µw
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
					// line 363: """Fill a single paragraph of text, returning a new string.
					πF.SetLineno(363)
					// line 371: w = TextWrapper(width=width, **kwargs)
					πF.SetLineno(371)
					if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
						continue
					}
					πTemp001 = πg.KWArgs{
						{"width", µwidth},
					}
					if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTextWrapper); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp002, nil, nil, πTemp001, µkwargs); πE != nil {
						continue
					}
					µw = πTemp003
					// line 372: return w.fill(text)
					πF.SetLineno(372)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp004[0] = µtext
					if πE = πg.CheckLocal(πF, µw, "w"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µw, ßfill, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
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
			if πE = πF.Globals().SetItem(πF, ßfill.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 377: _whitespace_only_re = re.compile('^[ \t]+$', re.MULTILINE)
			πF.SetLineno(377)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("^[ \t]+$").ToObject()
			if πTemp008, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp010, πE = πg.GetAttr(πF, πTemp008, ßMULTILINE, nil); πE != nil {
				continue
			}
			πTemp002[1] = πTemp010
			if πTemp008, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp010, πE = πg.GetAttr(πF, πTemp008, ßcompile, nil); πE != nil {
				continue
			}
			if πTemp008, πE = πTemp010.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ß_whitespace_only_re.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 378: _leading_whitespace_re = re.compile('(^[ \t]*)(?:[^ \t\n])', re.MULTILINE)
			πF.SetLineno(378)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("(^[ \t]*)(?:[^ \t\n])").ToObject()
			if πTemp008, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp010, πE = πg.GetAttr(πF, πTemp008, ßMULTILINE, nil); πE != nil {
				continue
			}
			πTemp002[1] = πTemp010
			if πTemp008, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp010, πE = πg.GetAttr(πF, πTemp008, ßcompile, nil); πE != nil {
				continue
			}
			if πTemp008, πE = πTemp010.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ß_leading_whitespace_re.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 380: def dedent(text):
			πF.SetLineno(380)
			πTemp009 = make([]πg.Param, 1)
			πTemp009[0] = πg.Param{Name: "text", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("dedent", "build/src/__python__/textwrap.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtext *πg.Object = πArgs[0]; _ = µtext
				var µmargin *πg.Object = πg.UnboundLocal; _ = µmargin
				var µindents *πg.Object = πg.UnboundLocal; _ = µindents
				var µindent *πg.Object = πg.UnboundLocal; _ = µindent
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µx *πg.Object = πg.UnboundLocal; _ = µx
				var µy *πg.Object = πg.UnboundLocal; _ = µy
				var µline *πg.Object = πg.UnboundLocal; _ = µline
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
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					case 9: goto Label9
					case 10: goto Label10
					case 17: goto Label17
					case 18: goto Label18
					default: panic("unexpected function state")
					}
					// line 381: """Remove any common leading whitespace from every line in `text`.
					πF.SetLineno(381)
					// line 395: margin = None
					πF.SetLineno(395)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µmargin = πTemp001
					// line 396: text = _whitespace_only_re.sub('', text)
					πF.SetLineno(396)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = ß.ToObject()
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp002[1] = µtext
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_whitespace_only_re); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßsub, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µtext = πTemp001
					// line 397: indents = _leading_whitespace_re.findall(text)
					πF.SetLineno(397)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp002[0] = µtext
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_leading_whitespace_re); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßfindall, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µindents = πTemp001
					if πE = πg.CheckLocal(πF, µindents, "indents"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µindents); πE != nil {
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
						µindent = πTemp003
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(1)            
					if πE = πg.CheckLocal(πF, µmargin, "margin"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µmargin == πTemp006).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label4
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmargin, "margin"); πE != nil {
						continue
					}
					πTemp002[0] = µmargin
					if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µindent, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label5
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
						continue
					}
					πTemp002[0] = µindent
					if πE = πg.CheckLocal(πF, µmargin, "margin"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µmargin, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label6
					}
					goto Label7
					// line 399: if margin is None:
					πF.SetLineno(399)
				Label4:
					// line 400: margin = indent
					πF.SetLineno(400)
					if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
						continue
					}
					µmargin = µindent
					goto Label8
					// line 404: elif indent.startswith(margin):
					πF.SetLineno(404)
				Label5:
					// line 405: pass
					πF.SetLineno(405)
					goto Label8
					// line 409: elif margin.startswith(indent):
					πF.SetLineno(409)
				Label6:
					// line 410: margin = indent
					πF.SetLineno(410)
					if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
						continue
					}
					µmargin = µindent
					goto Label8
				Label7:
					πTemp002 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µmargin, "margin"); πE != nil {
						continue
					}
					πTemp007[0] = µmargin
					if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
						continue
					}
					πTemp007[1] = µindent
					if πTemp006, πE = πg.ResolveGlobal(πF, ßzip); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp006.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp002[0] = πTemp008
					if πTemp006, πE = πg.ResolveGlobal(πF, ßenumerate); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.Iter(πF, πTemp008); πE != nil {
						continue
					}
					πF.PushCheckpoint(10)
					πTemp005 = false
				Label9:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp005 {
						πF.PopCheckpoint()
						goto Label11
					}
					if πTemp006, πE = πg.Next(πF, πTemp003); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp010}, πg.TieTarget{Target: &πTemp011}}}}}, πTemp006); πE != nil {
							continue
						}
						µi = πTemp008
						µx = πTemp010
						µy = πTemp011
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(9)            
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.NE(πF, µx, µy); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label12
					}
					goto Label13
					// line 416: if x != y:
					πF.SetLineno(416)
				Label12:
					// line 417: margin = margin[:i]
					πF.SetLineno(417)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µi, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmargin, "margin"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µmargin, πTemp006); πE != nil {
						continue
					}
					µmargin = πTemp008
					// line 418: break
					πF.SetLineno(418)
					πTemp005 = true
					continue
					goto Label13
				Label13:
					continue
				Label10:
					if πE != nil || πR != nil {
						continue
					}
					// line 420: margin = margin[:len(indent)]
					πF.SetLineno(420)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
						continue
					}
					πTemp002[0] = µindent
					if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp008.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp010, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmargin, "margin"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µmargin, πTemp006); πE != nil {
						continue
					}
					µmargin = πTemp008
				Label11:
					goto Label8
				Label8:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					πTemp001 = πg.NewInt(0).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label14
					}
					if πE = πg.CheckLocal(πF, µmargin, "margin"); πE != nil {
						continue
					}
					πTemp001 = µmargin
				Label14:
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label15
					}
					goto Label16
					// line 423: if 0 and margin:
					πF.SetLineno(423)
				Label15:
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("\n").ToObject()
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µtext, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Iter(πF, πTemp006); πE != nil {
						continue
					}
					πF.PushCheckpoint(18)
					πTemp004 = false
				Label17:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label19
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
						µline = πTemp003
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(17)            
					// line 425: assert not line or line.startswith(margin), \
					πF.SetLineno(425)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmargin, "margin"); πE != nil {
						continue
					}
					πTemp006 = πg.NewTuple2(µline, µmargin).ToObject()
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("line = %r, margin = %r").ToObject(), πTemp006); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, µline); πE != nil {
						continue
					}
					πTemp008 = πg.GetBool(!πTemp009).ToObject()
					πTemp006 = πTemp008
					if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label20
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmargin, "margin"); πE != nil {
						continue
					}
					πTemp002[0] = µmargin
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, µline, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp008.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp006 = πTemp010
				Label20:
					if πE = πg.Assert(πF, πTemp006, πTemp003); πE != nil {
						continue
					}
					continue
				Label18:
					if πE != nil || πR != nil {
						continue
					}
				Label19:
					goto Label16
				Label16:
					if πE = πg.CheckLocal(πF, µmargin, "margin"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µmargin); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label21
					}
					goto Label22
					// line 428: if margin:
					πF.SetLineno(428)
				Label21:
					// line 429: text = re.sub(r'(?m)^' + margin, '', text)
					πF.SetLineno(429)
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µmargin, "margin"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πg.NewStr("(?m)^").ToObject(), µmargin); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					πTemp002[1] = ß.ToObject()
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp002[2] = µtext
					if πTemp001, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßsub, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µtext = πTemp001
					goto Label22
				Label22:
					// line 430: return text
					πF.SetLineno(430)
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
			if πE = πF.Globals().SetItem(πF, ßdedent.ToObject(), πTemp008); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp010, πE = πg.Eq(πF, πTemp011, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsTrue(πF, πTemp010); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label4
			}
			goto Label5
			// line 432: if __name__ == "__main__":
			πF.SetLineno(432)
		Label4:
			// line 435: print dedent("Hello there.\n  This is indented.")
			πF.SetLineno(435)
			πTemp002 = make([]*πg.Object, 1)
			πTemp012 = πF.MakeArgs(1)
			πTemp012[0] = πg.NewStr("Hello there.\n  This is indented.").ToObject()
			if πTemp010, πE = πg.ResolveGlobal(πF, ßdedent); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp010.Call(πF, πTemp012, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp012)
			πTemp002[0] = πTemp011
			if πE = πg.Print(πF, πTemp002, true); πE != nil {
				continue
			}
			goto Label5
		Label5:
		}
		return nil, πE
	})
	πg.RegisterModule("textwrap", Code)
}
