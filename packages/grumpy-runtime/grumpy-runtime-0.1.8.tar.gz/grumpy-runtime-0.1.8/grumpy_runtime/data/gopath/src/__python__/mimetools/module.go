package mimetools
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/mimetools.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ß7bit := πg.InternStr("7bit")
		ß8bit := πg.InternStr("8bit")
		ßDeprecationWarning := πg.InternStr("DeprecationWarning")
		ßMessage := πg.InternStr("Message")
		ßNone := πg.InternStr("None")
		ßValueError := πg.InternStr("ValueError")
		ß__all__ := πg.InternStr("__all__")
		ß__enter__ := πg.InternStr("__enter__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_counter := πg.InternStr("_counter")
		ß_counter_lock := πg.InternStr("_counter_lock")
		ß_get_next_counter := πg.InternStr("_get_next_counter")
		ß_prefix := πg.InternStr("_prefix")
		ßacquire := πg.InternStr("acquire")
		ßallocate_lock := πg.InternStr("allocate_lock")
		ßappend := πg.InternStr("append")
		ßbase64 := πg.InternStr("base64")
		ßcatch_warnings := πg.InternStr("catch_warnings")
		ßchoose_boundary := πg.InternStr("choose_boundary")
		ßclose := πg.InternStr("close")
		ßcopybinary := πg.InternStr("copybinary")
		ßcopyliteral := πg.InternStr("copyliteral")
		ßdecode := πg.InternStr("decode")
		ßdecodetab := πg.InternStr("decodetab")
		ßencode := πg.InternStr("encode")
		ßencodetab := πg.InternStr("encodetab")
		ßencodingheader := πg.InternStr("encodingheader")
		ßfdopen := πg.InternStr("fdopen")
		ßfilterwarnings := πg.InternStr("filterwarnings")
		ßfind := πg.InternStr("find")
		ßgetencoding := πg.InternStr("getencoding")
		ßgetheader := πg.InternStr("getheader")
		ßgetmaintype := πg.InternStr("getmaintype")
		ßgetparam := πg.InternStr("getparam")
		ßgetparamnames := πg.InternStr("getparamnames")
		ßgetplist := πg.InternStr("getplist")
		ßgetsubtype := πg.InternStr("getsubtype")
		ßgettype := πg.InternStr("gettype")
		ßignore := πg.InternStr("ignore")
		ßindex := πg.InternStr("index")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßlower := πg.InternStr("lower")
		ßmaintype := πg.InternStr("maintype")
		ßmkstemp := πg.InternStr("mkstemp")
		ßos := πg.InternStr("os")
		ßparseplist := πg.InternStr("parseplist")
		ßparsetype := πg.InternStr("parsetype")
		ßpipethrough := πg.InternStr("pipethrough")
		ßpipeto := πg.InternStr("pipeto")
		ßplist := πg.InternStr("plist")
		ßplisttext := πg.InternStr("plisttext")
		ßpopen := πg.InternStr("popen")
		ßpy3kwarning := πg.InternStr("py3kwarning")
		ßr := πg.InternStr("r")
		ßrange := πg.InternStr("range")
		ßread := πg.InternStr("read")
		ßreadline := πg.InternStr("readline")
		ßrelease := πg.InternStr("release")
		ßrfc822 := πg.InternStr("rfc822")
		ßsplit := πg.InternStr("split")
		ßstrip := πg.InternStr("strip")
		ßsubtype := πg.InternStr("subtype")
		ßsys := πg.InternStr("sys")
		ßtempfile := πg.InternStr("tempfile")
		ßthread := πg.InternStr("thread")
		ßtype := πg.InternStr("type")
		ßtypeheader := πg.InternStr("typeheader")
		ßunlink := πg.InternStr("unlink")
		ßunquote := πg.InternStr("unquote")
		ßuudecode_pipe := πg.InternStr("uudecode_pipe")
		ßuue := πg.InternStr("uue")
		ßuuencode := πg.InternStr("uuencode")
		ßw := πg.InternStr("w")
		ßwarnpy3k := πg.InternStr("warnpy3k")
		ßwrite := πg.InternStr("write")
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
		var πTemp008 *πg.BaseException
		_ = πTemp008
		var πTemp009 *πg.Traceback
		_ = πTemp009
		var πTemp010 *πg.Type
		_ = πTemp010
		var πTemp011 πg.KWArgs
		_ = πTemp011
		var πTemp012 *πg.Dict
		_ = πTemp012
		var πTemp013 []πg.Param
		_ = πTemp013
		var πTemp014 *πg.Object
		_ = πTemp014
		var πTemp015 *πg.Object
		_ = πTemp015
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 1: goto Label1
			default: panic("unexpected function state")
			}
			// line 1: """Various tools used by MIME-reading or MIME-writing programs."""
			πF.SetLineno(1)
			// line 4: import os
			πF.SetLineno(4)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 5: import sys
			πF.SetLineno(5)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 6: import tempfile
			πF.SetLineno(6)
			if πTemp002, πE = πg.ImportModule(πF, "tempfile"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtempfile.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 7: from warnings import filterwarnings, catch_warnings
			πF.SetLineno(7)
			if πTemp002, πE = πg.ImportModule(πF, "warnings"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßfilterwarnings, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßfilterwarnings.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßcatch_warnings, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßcatch_warnings.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 8: with catch_warnings():
			πF.SetLineno(8)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßcatch_warnings); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__exit__, nil); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__enter__, nil); πE != nil {
				continue
			}
			if πTemp004, πE = πTemp004.Call(πF, πg.Args{πTemp003}, nil); πE != nil {
				continue
			}
			πF.PushCheckpoint(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßpy3kwarning, nil); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label2
			}
			goto Label3
			// line 9: if sys.py3kwarning:
			πF.SetLineno(9)
		Label2:
			// line 10: filterwarnings("ignore", ".*rfc822 has been removed", DeprecationWarning)
			πF.SetLineno(10)
			πTemp002 = πF.MakeArgs(3)
			πTemp002[0] = ßignore.ToObject()
			πTemp002[1] = πg.NewStr(".*rfc822 has been removed").ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßDeprecationWarning); πE != nil {
				continue
			}
			πTemp002[2] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßfilterwarnings); πE != nil {
				continue
			}
			if πTemp006, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			goto Label3
		Label3:
			// line 11: import rfc822
			πF.SetLineno(11)
			if πTemp002, πE = πg.ImportModule(πF, "rfc822"); πE != nil {
				continue
			}
			πTemp005 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßrfc822.ToObject(), πTemp005); πE != nil {
				continue
			}
			πF.PopCheckpoint()
		Label1:
			πTemp008, πTemp009 = nil, nil
			if πE != nil {
				πTemp008, πTemp009 = πF.ExcInfo()
			}
			if πTemp008 != nil {
				πTemp010 = πTemp008.Type()
				if πTemp005, πE = πTemp001.Call(πF, πg.Args{πTemp003, πTemp010.ToObject(), πTemp008.ToObject(), πTemp009.ToObject()}, nil); πE != nil {
					continue
				}
			} else {
				if πTemp005, πE = πTemp001.Call(πF, πg.Args{πTemp003, πg.None, πg.None, πg.None}, nil); πE != nil {
					continue
				}
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp008 != nil && πTemp007 != true {
				πE = πF.Raise(nil, nil, nil)
				continue
			}
			if πR != nil {
				continue
			}
			// line 13: from warnings import warnpy3k
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "warnings"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßwarnpy3k, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßwarnpy3k.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 14: warnpy3k("in 3.x, mimetools has been removed in favor of the email package",
			πF.SetLineno(14)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("in 3.x, mimetools has been removed in favor of the email package").ToObject()
			πTemp011 = πg.KWArgs{
				{"stacklevel", πg.NewInt(2).ToObject()},
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßwarnpy3k); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, πTemp011); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 17: __all__ = ["Message","choose_boundary","encode","decode","copyliteral",
			πF.SetLineno(17)
			πTemp002 = make([]*πg.Object, 6)
			πTemp002[0] = ßMessage.ToObject()
			πTemp002[1] = ßchoose_boundary.ToObject()
			πTemp002[2] = ßencode.ToObject()
			πTemp002[3] = ßdecode.ToObject()
			πTemp002[4] = ßcopyliteral.ToObject()
			πTemp002[5] = ßcopybinary.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 20: class Message(rfc822.Message):
			πF.SetLineno(20)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßrfc822); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßMessage, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp012 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp012.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Message", "build/src/__python__/mimetools.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp012
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
				var πTemp011 *πg.Object
				_ = πTemp011
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 21: """A derived class of rfc822.Message that knows about MIME headers and
					πF.SetLineno(21)
					// line 24: def __init__(self, fp, seekable = 1):
					πF.SetLineno(24)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "fp", Def: nil}
					πTemp002[2] = πg.Param{Name: "seekable", Def: πg.NewInt(1).ToObject()}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/mimetools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfp *πg.Object = πArgs[1]; _ = µfp
						var µseekable *πg.Object = πArgs[2]; _ = µseekable
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
							// line 25: rfc822.Message.__init__(self, fp, seekable)
							πF.SetLineno(25)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
								continue
							}
							πTemp001[1] = µfp
							if πE = πg.CheckLocal(πF, µseekable, "seekable"); πE != nil {
								continue
							}
							πTemp001[2] = µseekable
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrfc822); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßMessage, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 26: self.encodingheader = \
							πF.SetLineno(26)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("content-transfer-encoding").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgetheader, nil); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßencodingheader, πTemp002); πE != nil {
								continue
							}
							// line 28: self.typeheader = \
							πF.SetLineno(28)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("content-type").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgetheader, nil); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßtypeheader, πTemp002); πE != nil {
								continue
							}
							// line 30: self.parsetype()
							πF.SetLineno(30)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßparsetype, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 31: self.parseplist()
							πF.SetLineno(31)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßparseplist, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					// line 33: def parsetype(self):
					πF.SetLineno(33)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("parsetype", "build/src/__python__/mimetools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstr *πg.Object = πg.UnboundLocal; _ = µstr
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µfields *πg.Object = πg.UnboundLocal; _ = µfields
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
						var πTemp006 *πg.Object
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
							case 6: goto Label6
							case 7: goto Label7
							default: panic("unexpected function state")
							}
							// line 34: str = self.typeheader
							πF.SetLineno(34)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtypeheader, nil); πE != nil {
								continue
							}
							µstr = πTemp001
							if πE = πg.CheckLocal(πF, µstr, "str"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µstr == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 35: if str is None:
							πF.SetLineno(35)
						Label1:
							// line 36: str = 'text/plain'
							πF.SetLineno(36)
							µstr = πg.NewStr("text/plain").ToObject()
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µstr, "str"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, µstr, πg.NewStr(";").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 37: if ';' in str:
							πF.SetLineno(37)
						Label3:
							// line 38: i = str.index(';')
							πF.SetLineno(38)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr(";").ToObject()
							if πE = πg.CheckLocal(πF, µstr, "str"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstr, ßindex, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µi = πTemp002
							// line 39: self.plisttext = str[i:]
							πF.SetLineno(39)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µi, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstr, "str"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µstr, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßplisttext, πTemp001); πE != nil {
								continue
							}
							// line 40: str = str[:i]
							πF.SetLineno(40)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µi, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstr, "str"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µstr, πTemp001); πE != nil {
								continue
							}
							µstr = πTemp002
							goto Label5
						Label4:
							// line 42: self.plisttext = ''
							πF.SetLineno(42)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, ß.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßplisttext, πTemp001); πE != nil {
								continue
							}
							goto Label5
						Label5:
							// line 43: fields = str.split('/')
							πF.SetLineno(43)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("/").ToObject()
							if πE = πg.CheckLocal(πF, µstr, "str"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstr, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µfields = πTemp002
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfields, "fields"); πE != nil {
								continue
							}
							πTemp005[0] = µfields
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[0] = πTemp006
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.Iter(πF, πTemp006); πE != nil {
								continue
							}
							πF.PushCheckpoint(7)
							πTemp003 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
								πF.PopCheckpoint()
								goto Label8
							}
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µi = πTemp002
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(6)            
							// line 45: fields[i] = fields[i].strip().lower()
							πF.SetLineno(45)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp002 = µi
							if πE = πg.CheckLocal(πF, µfields, "fields"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µfields, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp006, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp006, ßlower, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfields, "fields"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp008 = µi
							if πE = πg.SetItem(πF, µfields, πTemp008, πTemp002); πE != nil {
								continue
							}
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							// line 46: self.type = '/'.join(fields)
							πF.SetLineno(46)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfields, "fields"); πE != nil {
								continue
							}
							πTemp004[0] = µfields
							if πTemp001, πE = πg.GetAttr(πF, πg.NewStr("/").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtype, πTemp001); πE != nil {
								continue
							}
							// line 47: self.maintype = fields[0]
							πF.SetLineno(47)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µfields, "fields"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µfields, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmaintype, πTemp001); πE != nil {
								continue
							}
							// line 48: self.subtype = '/'.join(fields[1:])
							πF.SetLineno(48)
							πTemp004 = πF.MakeArgs(1)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfields, "fields"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µfields, πTemp001); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πTemp001, πE = πg.GetAttr(πF, πg.NewStr("/").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsubtype, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßparsetype.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 50: def parseplist(self):
					πF.SetLineno(50)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("parseplist", "build/src/__python__/mimetools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstr *πg.Object = πg.UnboundLocal; _ = µstr
						var µend *πg.Object = πg.UnboundLocal; _ = µend
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var µi *πg.Object = πg.UnboundLocal; _ = µi
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 51: str = self.plisttext
							πF.SetLineno(51)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßplisttext, nil); πE != nil {
								continue
							}
							µstr = πTemp001
							// line 52: self.plist = []
							πF.SetLineno(52)
							πTemp002 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßplist, πTemp003); πE != nil {
								continue
							}
							// line 53: while str[:1] == ';':
							πF.SetLineno(53)
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
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(1).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstr, "str"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µstr, πTemp003); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp006, πg.NewStr(";").ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 54: str = str[1:]
							πF.SetLineno(54)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstr, "str"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µstr, πTemp001); πE != nil {
								continue
							}
							µstr = πTemp003
							if πE = πg.CheckLocal(πF, µstr, "str"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, µstr, πg.NewStr(";").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 55: if ';' in str:
							πF.SetLineno(55)
						Label4:
							// line 57: end = str.index(';')
							πF.SetLineno(57)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr(";").ToObject()
							if πE = πg.CheckLocal(πF, µstr, "str"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstr, ßindex, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µend = πTemp003
							goto Label6
						Label5:
							// line 59: end = len(str)
							πF.SetLineno(59)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstr, "str"); πE != nil {
								continue
							}
							πTemp002[0] = µstr
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µend = πTemp003
							goto Label6
						Label6:
							// line 60: f = str[:end]
							πF.SetLineno(60)
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µend, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstr, "str"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µstr, πTemp001); πE != nil {
								continue
							}
							µf = πTemp003
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, µf, πg.NewStr("=").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label7
							}
							goto Label8
							// line 61: if '=' in f:
							πF.SetLineno(61)
						Label7:
							// line 62: i = f.index('=')
							πF.SetLineno(62)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("=").ToObject()
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µf, ßindex, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µi = πTemp003
							// line 63: f = f[:i].strip().lower() + \
							πF.SetLineno(63)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µi, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µf, πTemp006); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp007, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp007, ßlower, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp007, πg.NewStr("=").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πTemp007, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µf, πTemp006); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp007, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, πTemp007); πE != nil {
								continue
							}
							µf = πTemp001
							goto Label8
						Label8:
							// line 65: self.plist.append(f.strip())
							πF.SetLineno(65)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µf, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßplist, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 66: str = str[end:]
							πF.SetLineno(66)
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µend, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstr, "str"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µstr, πTemp001); πE != nil {
								continue
							}
							µstr = πTemp003
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
					if πE = πClass.SetItem(πF, ßparseplist.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 68: def getplist(self):
					πF.SetLineno(68)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("getplist", "build/src/__python__/mimetools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 69: return self.plist
							πF.SetLineno(69)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßplist, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßgetplist.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 71: def getparam(self, name):
					πF.SetLineno(71)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("getparam", "build/src/__python__/mimetools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var µp *πg.Object = πg.UnboundLocal; _ = µp
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
							// line 72: name = name.lower() + '='
							πF.SetLineno(72)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µname, ßlower, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, πg.NewStr("=").ToObject()); πE != nil {
								continue
							}
							µname = πTemp001
							// line 73: n = len(name)
							πF.SetLineno(73)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp004[0] = µname
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µn = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßplist, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
								µp = πTemp002
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µn, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µp, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp007, µname); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 75: if p[:n] == name:
							πF.SetLineno(75)
						Label4:
							// line 76: return rfc822.unquote(p[n:])
							πF.SetLineno(76)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µn, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µp, πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrfc822); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunquote, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πR = πTemp002
							continue
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 77: return None
							πF.SetLineno(77)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
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
					if πE = πClass.SetItem(πF, ßgetparam.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 79: def getparamnames(self):
					πF.SetLineno(79)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("getparamnames", "build/src/__python__/mimetools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
						var µp *πg.Object = πg.UnboundLocal; _ = µp
						var µi *πg.Object = πg.UnboundLocal; _ = µi
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 80: result = []
							πF.SetLineno(80)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µresult = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßplist, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								µp = πTemp003
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 82: i = p.find('=')
							πF.SetLineno(82)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("=").ToObject()
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µp, ßfind, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µi = πTemp006
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GE(πF, µi, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 83: if i >= 0:
							πF.SetLineno(83)
						Label4:
							// line 84: result.append(p[:i].lower())
							πF.SetLineno(84)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µi, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µp, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp006, ßlower, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 85: return result
							πF.SetLineno(85)
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
					if πE = πClass.SetItem(πF, ßgetparamnames.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 87: def getencoding(self):
					πF.SetLineno(87)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("getencoding", "build/src/__python__/mimetools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßencodingheader, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 88: if self.encodingheader is None:
							πF.SetLineno(88)
						Label1:
							// line 89: return '7bit'
							πF.SetLineno(89)
							πR = ß7bit.ToObject()
							continue
							goto Label2
						Label2:
							// line 90: return self.encodingheader.lower()
							πF.SetLineno(90)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßencodingheader, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßlower, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßgetencoding.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 92: def gettype(self):
					πF.SetLineno(92)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("gettype", "build/src/__python__/mimetools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 93: return self.type
							πF.SetLineno(93)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtype, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßgettype.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 95: def getmaintype(self):
					πF.SetLineno(95)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("getmaintype", "build/src/__python__/mimetools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 96: return self.maintype
							πF.SetLineno(96)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmaintype, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßgetmaintype.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 98: def getsubtype(self):
					πF.SetLineno(98)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("getsubtype", "build/src/__python__/mimetools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 99: return self.subtype
							πF.SetLineno(99)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsubtype, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßgetsubtype.ToObject(), πTemp011); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp012.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Message").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp012.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMessage.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 108: import thread
			πF.SetLineno(108)
			if πTemp002, πE = πg.ImportModule(πF, "thread"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßthread.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 111: _counter_lock = thread.allocate_lock()
			πF.SetLineno(111)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßthread); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßallocate_lock, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_counter_lock.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 112: del thread
			πF.SetLineno(112)
			if πE = πg.DelVar(πF, πF.Globals(), ßthread); πE != nil {
				continue
			}
			// line 114: _counter = 0
			πF.SetLineno(114)
			if πE = πF.Globals().SetItem(πF, ß_counter.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
				continue
			}
			// line 115: def _get_next_counter():
			πF.SetLineno(115)
			πTemp013 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("_get_next_counter", "build/src/__python__/mimetools.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
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
					// line 116: global _counter
					πF.SetLineno(116)
					// line 117: _counter_lock.acquire()
					πF.SetLineno(117)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_counter_lock); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßacquire, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 118: _counter += 1
					πF.SetLineno(118)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_counter); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πF.Globals().SetItem(πF, ß_counter.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 119: result = _counter
					πF.SetLineno(119)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_counter); πE != nil {
						continue
					}
					µresult = πTemp001
					// line 120: _counter_lock.release()
					πF.SetLineno(120)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_counter_lock); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrelease, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 121: return result
					πF.SetLineno(121)
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
			if πE = πF.Globals().SetItem(πF, ß_get_next_counter.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 123: _prefix = None
			πF.SetLineno(123)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_prefix.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 157: def decode(input, output, encoding):
			πF.SetLineno(157)
			πTemp013 = make([]πg.Param, 3)
			πTemp013[0] = πg.Param{Name: "input", Def: nil}
			πTemp013[1] = πg.Param{Name: "output", Def: nil}
			πTemp013[2] = πg.Param{Name: "encoding", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("decode", "build/src/__python__/mimetools.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µinput *πg.Object = πArgs[0]; _ = µinput
				var µoutput *πg.Object = πArgs[1]; _ = µoutput
				var µencoding *πg.Object = πArgs[2]; _ = µencoding
				var µbase64 *πg.Object = πg.UnboundLocal; _ = µbase64
				var µquopri *πg.Object = πg.UnboundLocal; _ = µquopri
				var µuu *πg.Object = πg.UnboundLocal; _ = µuu
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
					// line 158: """Decode common content-transfer-encodings (base64, quopri, uuencode)."""
					πF.SetLineno(158)
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µencoding, ßbase64.ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 159: if encoding == 'base64':
					πF.SetLineno(159)
				Label1:
					// line 160: import base64
					πF.SetLineno(160)
					if πTemp003, πE = πg.ImportModule(πF, "base64"); πE != nil {
						continue
					}
					πTemp001 = πTemp003[0]
					µbase64 = πTemp001
					// line 161: return base64.decode(input, output)
					πF.SetLineno(161)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µinput, "input"); πE != nil {
						continue
					}
					πTemp003[0] = µinput
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					πTemp003[1] = µoutput
					if πE = πg.CheckLocal(πF, µbase64, "base64"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µbase64, ßdecode, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp004
					continue
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µencoding, πg.NewStr("quoted-printable").ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 162: if encoding == 'quoted-printable':
					πF.SetLineno(162)
				Label3:
					// line 163: import quopri
					πF.SetLineno(163)
					if πTemp003, πE = πg.ImportModule(πF, "quopri"); πE != nil {
						continue
					}
					πTemp001 = πTemp003[0]
					µquopri = πTemp001
					// line 164: return quopri.decode(input, output)
					πF.SetLineno(164)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µinput, "input"); πE != nil {
						continue
					}
					πTemp003[0] = µinput
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					πTemp003[1] = µoutput
					if πE = πg.CheckLocal(πF, µquopri, "quopri"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µquopri, ßdecode, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp004
					continue
					goto Label4
				Label4:
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple4(ßuuencode.ToObject(), πg.NewStr("x-uuencode").ToObject(), ßuue.ToObject(), πg.NewStr("x-uue").ToObject()).ToObject()
					if πTemp002, πE = πg.Contains(πF, πTemp004, µencoding); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label5
					}
					goto Label6
					// line 165: if encoding in ('uuencode', 'x-uuencode', 'uue', 'x-uue'):
					πF.SetLineno(165)
				Label5:
					// line 166: import uu
					πF.SetLineno(166)
					if πTemp003, πE = πg.ImportModule(πF, "uu"); πE != nil {
						continue
					}
					πTemp001 = πTemp003[0]
					µuu = πTemp001
					// line 167: return uu.decode(input, output)
					πF.SetLineno(167)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µinput, "input"); πE != nil {
						continue
					}
					πTemp003[0] = µinput
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					πTemp003[1] = µoutput
					if πE = πg.CheckLocal(πF, µuu, "uu"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µuu, ßdecode, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp004
					continue
					goto Label6
				Label6:
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple2(ß7bit.ToObject(), ß8bit.ToObject()).ToObject()
					if πTemp002, πE = πg.Contains(πF, πTemp004, µencoding); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label7
					}
					goto Label8
					// line 168: if encoding in ('7bit', '8bit'):
					πF.SetLineno(168)
				Label7:
					// line 169: return output.write(input.read())
					πF.SetLineno(169)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µinput, "input"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µinput, ßread, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µoutput, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp004
					continue
					goto Label8
				Label8:
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßdecodetab); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, πTemp004, µencoding); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label9
					}
					goto Label10
					// line 170: if encoding in decodetab:
					πF.SetLineno(170)
				Label9:
					// line 171: pipethrough(input, decodetab[encoding], output)
					πF.SetLineno(171)
					πTemp003 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µinput, "input"); πE != nil {
						continue
					}
					πTemp003[0] = µinput
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					πTemp001 = µencoding
					if πTemp005, πE = πg.ResolveGlobal(πF, ßdecodetab); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
						continue
					}
					πTemp003[1] = πTemp004
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					πTemp003[2] = µoutput
					if πTemp001, πE = πg.ResolveGlobal(πF, ßpipethrough); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label11
				Label10:
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mod(πF, πg.NewStr("unknown Content-Transfer-Encoding: %s").ToObject(), µencoding); πE != nil {
						continue
					}
					// line 173: raise ValueError, \
					πF.SetLineno(173)
					πE = πF.Raise(πTemp001, πTemp004, nil)
					continue
					goto Label11
				Label11:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßdecode.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 176: def encode(input, output, encoding):
			πF.SetLineno(176)
			πTemp013 = make([]πg.Param, 3)
			πTemp013[0] = πg.Param{Name: "input", Def: nil}
			πTemp013[1] = πg.Param{Name: "output", Def: nil}
			πTemp013[2] = πg.Param{Name: "encoding", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("encode", "build/src/__python__/mimetools.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µinput *πg.Object = πArgs[0]; _ = µinput
				var µoutput *πg.Object = πArgs[1]; _ = µoutput
				var µencoding *πg.Object = πArgs[2]; _ = µencoding
				var µbase64 *πg.Object = πg.UnboundLocal; _ = µbase64
				var µquopri *πg.Object = πg.UnboundLocal; _ = µquopri
				var µuu *πg.Object = πg.UnboundLocal; _ = µuu
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
					// line 177: """Encode common content-transfer-encodings (base64, quopri, uuencode)."""
					πF.SetLineno(177)
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µencoding, ßbase64.ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 178: if encoding == 'base64':
					πF.SetLineno(178)
				Label1:
					// line 179: import base64
					πF.SetLineno(179)
					if πTemp003, πE = πg.ImportModule(πF, "base64"); πE != nil {
						continue
					}
					πTemp001 = πTemp003[0]
					µbase64 = πTemp001
					// line 180: return base64.encode(input, output)
					πF.SetLineno(180)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µinput, "input"); πE != nil {
						continue
					}
					πTemp003[0] = µinput
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					πTemp003[1] = µoutput
					if πE = πg.CheckLocal(πF, µbase64, "base64"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µbase64, ßencode, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp004
					continue
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µencoding, πg.NewStr("quoted-printable").ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 181: if encoding == 'quoted-printable':
					πF.SetLineno(181)
				Label3:
					// line 182: import quopri
					πF.SetLineno(182)
					if πTemp003, πE = πg.ImportModule(πF, "quopri"); πE != nil {
						continue
					}
					πTemp001 = πTemp003[0]
					µquopri = πTemp001
					// line 183: return quopri.encode(input, output, 0)
					πF.SetLineno(183)
					πTemp003 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µinput, "input"); πE != nil {
						continue
					}
					πTemp003[0] = µinput
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					πTemp003[1] = µoutput
					πTemp003[2] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µquopri, "quopri"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µquopri, ßencode, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp004
					continue
					goto Label4
				Label4:
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple4(ßuuencode.ToObject(), πg.NewStr("x-uuencode").ToObject(), ßuue.ToObject(), πg.NewStr("x-uue").ToObject()).ToObject()
					if πTemp002, πE = πg.Contains(πF, πTemp004, µencoding); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label5
					}
					goto Label6
					// line 184: if encoding in ('uuencode', 'x-uuencode', 'uue', 'x-uue'):
					πF.SetLineno(184)
				Label5:
					// line 185: import uu
					πF.SetLineno(185)
					if πTemp003, πE = πg.ImportModule(πF, "uu"); πE != nil {
						continue
					}
					πTemp001 = πTemp003[0]
					µuu = πTemp001
					// line 186: return uu.encode(input, output)
					πF.SetLineno(186)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µinput, "input"); πE != nil {
						continue
					}
					πTemp003[0] = µinput
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					πTemp003[1] = µoutput
					if πE = πg.CheckLocal(πF, µuu, "uu"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µuu, ßencode, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp004
					continue
					goto Label6
				Label6:
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple2(ß7bit.ToObject(), ß8bit.ToObject()).ToObject()
					if πTemp002, πE = πg.Contains(πF, πTemp004, µencoding); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label7
					}
					goto Label8
					// line 187: if encoding in ('7bit', '8bit'):
					πF.SetLineno(187)
				Label7:
					// line 188: return output.write(input.read())
					πF.SetLineno(188)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µinput, "input"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µinput, ßread, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µoutput, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp004
					continue
					goto Label8
				Label8:
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßencodetab); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, πTemp004, µencoding); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label9
					}
					goto Label10
					// line 189: if encoding in encodetab:
					πF.SetLineno(189)
				Label9:
					// line 190: pipethrough(input, encodetab[encoding], output)
					πF.SetLineno(190)
					πTemp003 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µinput, "input"); πE != nil {
						continue
					}
					πTemp003[0] = µinput
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					πTemp001 = µencoding
					if πTemp005, πE = πg.ResolveGlobal(πF, ßencodetab); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
						continue
					}
					πTemp003[1] = πTemp004
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					πTemp003[2] = µoutput
					if πTemp001, πE = πg.ResolveGlobal(πF, ßpipethrough); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label11
				Label10:
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mod(πF, πg.NewStr("unknown Content-Transfer-Encoding: %s").ToObject(), µencoding); πE != nil {
						continue
					}
					// line 192: raise ValueError, \
					πF.SetLineno(192)
					πE = πF.Raise(πTemp001, πTemp004, nil)
					continue
					goto Label11
				Label11:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßencode.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 199: uudecode_pipe = '''(
			πF.SetLineno(199)
			if πE = πF.Globals().SetItem(πF, ßuudecode_pipe.ToObject(), πg.NewStr("(\nTEMP=/tmp/@uu.$$\nsed \"s%^begin [0-7][0-7]* .*%begin 600 $TEMP%\" | uudecode\ncat $TEMP\nrm $TEMP\n)").ToObject()); πE != nil {
				continue
			}
			// line 206: decodetab = {
			πF.SetLineno(206)
			πTemp012 = πg.NewDict()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßuudecode_pipe); πE != nil {
				continue
			}
			if πE = πTemp012.SetItem(πF, ßuuencode.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßuudecode_pipe); πE != nil {
				continue
			}
			if πE = πTemp012.SetItem(πF, πg.NewStr("x-uuencode").ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßuudecode_pipe); πE != nil {
				continue
			}
			if πE = πTemp012.SetItem(πF, ßuue.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßuudecode_pipe); πE != nil {
				continue
			}
			if πE = πTemp012.SetItem(πF, πg.NewStr("x-uue").ToObject(), πTemp005); πE != nil {
				continue
			}
			if πE = πTemp012.SetItem(πF, πg.NewStr("quoted-printable").ToObject(), πg.NewStr("mmencode -u -q").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp012.SetItem(πF, ßbase64.ToObject(), πg.NewStr("mmencode -u -b").ToObject()); πE != nil {
				continue
			}
			πTemp005 = πTemp012.ToObject()
			if πE = πF.Globals().SetItem(πF, ßdecodetab.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 215: encodetab = {
			πF.SetLineno(215)
			πTemp012 = πg.NewDict()
			if πE = πTemp012.SetItem(πF, πg.NewStr("x-uuencode").ToObject(), πg.NewStr("uuencode tempfile").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp012.SetItem(πF, ßuuencode.ToObject(), πg.NewStr("uuencode tempfile").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp012.SetItem(πF, πg.NewStr("x-uue").ToObject(), πg.NewStr("uuencode tempfile").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp012.SetItem(πF, ßuue.ToObject(), πg.NewStr("uuencode tempfile").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp012.SetItem(πF, πg.NewStr("quoted-printable").ToObject(), πg.NewStr("mmencode -q").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp012.SetItem(πF, ßbase64.ToObject(), πg.NewStr("mmencode -b").ToObject()); πE != nil {
				continue
			}
			πTemp005 = πTemp012.ToObject()
			if πE = πF.Globals().SetItem(πF, ßencodetab.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 224: def pipeto(input, command):
			πF.SetLineno(224)
			πTemp013 = make([]πg.Param, 2)
			πTemp013[0] = πg.Param{Name: "input", Def: nil}
			πTemp013[1] = πg.Param{Name: "command", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("pipeto", "build/src/__python__/mimetools.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µinput *πg.Object = πArgs[0]; _ = µinput
				var µcommand *πg.Object = πArgs[1]; _ = µcommand
				var µpipe *πg.Object = πg.UnboundLocal; _ = µpipe
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
					// line 225: pipe = os.popen(command, 'w')
					πF.SetLineno(225)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µcommand, "command"); πE != nil {
						continue
					}
					πTemp001[0] = µcommand
					πTemp001[1] = ßw.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpopen, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µpipe = πTemp002
					// line 226: copyliteral(input, pipe)
					πF.SetLineno(226)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µinput, "input"); πE != nil {
						continue
					}
					πTemp001[0] = µinput
					if πE = πg.CheckLocal(πF, µpipe, "pipe"); πE != nil {
						continue
					}
					πTemp001[1] = µpipe
					if πTemp002, πE = πg.ResolveGlobal(πF, ßcopyliteral); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 227: pipe.close()
					πF.SetLineno(227)
					if πE = πg.CheckLocal(πF, µpipe, "pipe"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpipe, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßpipeto.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 229: def pipethrough(input, command, output):
			πF.SetLineno(229)
			πTemp013 = make([]πg.Param, 3)
			πTemp013[0] = πg.Param{Name: "input", Def: nil}
			πTemp013[1] = πg.Param{Name: "command", Def: nil}
			πTemp013[2] = πg.Param{Name: "output", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("pipethrough", "build/src/__python__/mimetools.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µinput *πg.Object = πArgs[0]; _ = µinput
				var µcommand *πg.Object = πArgs[1]; _ = µcommand
				var µoutput *πg.Object = πArgs[2]; _ = µoutput
				var µfd *πg.Object = πg.UnboundLocal; _ = µfd
				var µtempname *πg.Object = πg.UnboundLocal; _ = µtempname
				var µtemp *πg.Object = πg.UnboundLocal; _ = µtemp
				var µpipe *πg.Object = πg.UnboundLocal; _ = µpipe
				var πTemp001 *πg.Object
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
					// line 230: (fd, tempname) = tempfile.mkstemp()
					πF.SetLineno(230)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtempfile); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmkstemp, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}}}, πTemp001); πE != nil {
						continue
					}
					µfd = πTemp002
					µtempname = πTemp003
					// line 231: temp = os.fdopen(fd, 'w')
					πF.SetLineno(231)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp004[0] = µfd
					πTemp004[1] = ßw.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßfdopen, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µtemp = πTemp001
					// line 232: copyliteral(input, temp)
					πF.SetLineno(232)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µinput, "input"); πE != nil {
						continue
					}
					πTemp004[0] = µinput
					if πE = πg.CheckLocal(πF, µtemp, "temp"); πE != nil {
						continue
					}
					πTemp004[1] = µtemp
					if πTemp001, πE = πg.ResolveGlobal(πF, ßcopyliteral); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 233: temp.close()
					πF.SetLineno(233)
					if πE = πg.CheckLocal(πF, µtemp, "temp"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µtemp, ßclose, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 234: pipe = os.popen(command + ' <' + tempname, 'r')
					πF.SetLineno(234)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µcommand, "command"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µcommand, πg.NewStr(" <").ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtempname, "tempname"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp002, µtempname); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					πTemp004[1] = ßr.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpopen, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µpipe = πTemp001
					// line 235: copybinary(pipe, output)
					πF.SetLineno(235)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpipe, "pipe"); πE != nil {
						continue
					}
					πTemp004[0] = µpipe
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					πTemp004[1] = µoutput
					if πTemp001, πE = πg.ResolveGlobal(πF, ßcopybinary); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 236: pipe.close()
					πF.SetLineno(236)
					if πE = πg.CheckLocal(πF, µpipe, "pipe"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpipe, ßclose, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 237: os.unlink(tempname)
					πF.SetLineno(237)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtempname, "tempname"); πE != nil {
						continue
					}
					πTemp004[0] = µtempname
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßunlink, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßpipethrough.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 239: def copyliteral(input, output):
			πF.SetLineno(239)
			πTemp013 = make([]πg.Param, 2)
			πTemp013[0] = πg.Param{Name: "input", Def: nil}
			πTemp013[1] = πg.Param{Name: "output", Def: nil}
			πTemp014 = πg.NewFunction(πg.NewCode("copyliteral", "build/src/__python__/mimetools.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µinput *πg.Object = πArgs[0]; _ = µinput
				var µoutput *πg.Object = πArgs[1]; _ = µoutput
				var µline *πg.Object = πg.UnboundLocal; _ = µline
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
					// line 240: while 1:
					πF.SetLineno(240)
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
					if πTemp002, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE != nil || !πTemp002 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 241: line = input.readline()
					πF.SetLineno(241)
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
					// line 242: if not line: break
					πF.SetLineno(242)
				Label4:
					// line 242: if not line: break
					πF.SetLineno(242)
					πTemp001 = true
					continue
					goto Label5
				Label5:
					// line 243: output.write(line)
					πF.SetLineno(243)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp005[0] = µline
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
			if πE = πF.Globals().SetItem(πF, ßcopyliteral.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 245: def copybinary(input, output):
			πF.SetLineno(245)
			πTemp013 = make([]πg.Param, 2)
			πTemp013[0] = πg.Param{Name: "input", Def: nil}
			πTemp013[1] = πg.Param{Name: "output", Def: nil}
			πTemp015 = πg.NewFunction(πg.NewCode("copybinary", "build/src/__python__/mimetools.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µinput *πg.Object = πArgs[0]; _ = µinput
				var µoutput *πg.Object = πArgs[1]; _ = µoutput
				var µBUFSIZE *πg.Object = πg.UnboundLocal; _ = µBUFSIZE
				var µline *πg.Object = πg.UnboundLocal; _ = µline
				var πTemp001 bool
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
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 246: BUFSIZE = 8192
					πF.SetLineno(246)
					µBUFSIZE = πg.NewInt(8192).ToObject()
					// line 247: while 1:
					πF.SetLineno(247)
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
					if πTemp002, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE != nil || !πTemp002 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 248: line = input.read(BUFSIZE)
					πF.SetLineno(248)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µBUFSIZE, "BUFSIZE"); πE != nil {
						continue
					}
					πTemp003[0] = µBUFSIZE
					if πE = πg.CheckLocal(πF, µinput, "input"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µinput, ßread, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µline = πTemp005
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µline); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					goto Label5
					// line 249: if not line: break
					πF.SetLineno(249)
				Label4:
					// line 249: if not line: break
					πF.SetLineno(249)
					πTemp001 = true
					continue
					goto Label5
				Label5:
					// line 250: output.write(line)
					πF.SetLineno(250)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp003[0] = µline
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µoutput, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
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
			if πE = πF.Globals().SetItem(πF, ßcopybinary.ToObject(), πTemp015); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("mimetools", Code)
}
