package genericpath
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/genericpath.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßFalse := πg.InternStr("False")
		ßNameError := πg.InternStr("NameError")
		ßS_ISDIR := πg.InternStr("S_ISDIR")
		ßS_ISREG := πg.InternStr("S_ISREG")
		ßTrue := πg.InternStr("True")
		ß__all__ := πg.InternStr("__all__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_splitext := πg.InternStr("_splitext")
		ß_unicode := πg.InternStr("_unicode")
		ßcommonprefix := πg.InternStr("commonprefix")
		ßenumerate := πg.InternStr("enumerate")
		ßerror := πg.InternStr("error")
		ßexists := πg.InternStr("exists")
		ßgetatime := πg.InternStr("getatime")
		ßgetctime := πg.InternStr("getctime")
		ßgetmtime := πg.InternStr("getmtime")
		ßgetsize := πg.InternStr("getsize")
		ßisdir := πg.InternStr("isdir")
		ßisfile := πg.InternStr("isfile")
		ßmax := πg.InternStr("max")
		ßmin := πg.InternStr("min")
		ßobject := πg.InternStr("object")
		ßos := πg.InternStr("os")
		ßrfind := πg.InternStr("rfind")
		ßst_atime := πg.InternStr("st_atime")
		ßst_ctime := πg.InternStr("st_ctime")
		ßst_mode := πg.InternStr("st_mode")
		ßst_mtime := πg.InternStr("st_mtime")
		ßst_size := πg.InternStr("st_size")
		ßstat := πg.InternStr("stat")
		ßunicode := πg.InternStr("unicode")
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
		var πTemp012 *πg.Object
		_ = πTemp012
		var πTemp013 *πg.Object
		_ = πTemp013
		var πTemp014 *πg.Object
		_ = πTemp014
		var πTemp015 *πg.Object
		_ = πTemp015
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 2: goto Label2
			default: panic("unexpected function state")
			}
			// line 1: """
			πF.SetLineno(1)
			// line 6: import os
			πF.SetLineno(6)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 7: import stat
			πF.SetLineno(7)
			if πTemp002, πE = πg.ImportModule(πF, "stat"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßstat.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 9: __all__ = ['commonprefix', 'exists', 'getatime', 'getctime', 'getmtime',
			πF.SetLineno(9)
			πTemp002 = make([]*πg.Object, 8)
			πTemp002[0] = ßcommonprefix.ToObject()
			πTemp002[1] = ßexists.ToObject()
			πTemp002[2] = ßgetatime.ToObject()
			πTemp002[3] = ßgetctime.ToObject()
			πTemp002[4] = ßgetmtime.ToObject()
			πTemp002[5] = ßgetsize.ToObject()
			πTemp002[6] = ßisdir.ToObject()
			πTemp002[7] = ßisfile.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 13: try:
			πF.SetLineno(13)
			πF.PushCheckpoint(2)
			// line 14: _unicode = unicode
			πF.SetLineno(14)
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
			// line 15: except NameError:
			πF.SetLineno(15)
		Label3:
			// line 18: class _unicode(object):
			πF.SetLineno(18)
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
			_, πE = πg.NewCode("_unicode", "build/src/__python__/genericpath.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 19: pass
					πF.SetLineno(19)
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
			// line 23: def exists(path):
			πF.SetLineno(23)
			πTemp009 = make([]πg.Param, 1)
			πTemp009[0] = πg.Param{Name: "path", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("exists", "build/src/__python__/genericpath.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πArgs[0]; _ = µpath
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 24: """Test whether a path exists.  Returns False for broken symbolic links"""
					πF.SetLineno(24)
					// line 25: try:
					πF.SetLineno(25)
					πF.PushCheckpoint(2)
					// line 26: os.stat(path)
					πF.SetLineno(26)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp001[0] = µpath
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstat, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßerror, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 27: except os.error:
					πF.SetLineno(27)
				Label3:
					// line 28: return False
					πF.SetLineno(28)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 29: return True
					πF.SetLineno(29)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßexists.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 34: def isfile(path):
			πF.SetLineno(34)
			πTemp009 = make([]πg.Param, 1)
			πTemp009[0] = πg.Param{Name: "path", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("isfile", "build/src/__python__/genericpath.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πArgs[0]; _ = µpath
				var µst *πg.Object = πg.UnboundLocal; _ = µst
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 35: """Test whether a path is a regular file"""
					πF.SetLineno(35)
					// line 36: try:
					πF.SetLineno(36)
					πF.PushCheckpoint(2)
					// line 37: st = os.stat(path)
					πF.SetLineno(37)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp001[0] = µpath
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstat, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µst = πTemp002
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßerror, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 38: except os.error:
					πF.SetLineno(38)
				Label3:
					// line 39: return False
					πF.SetLineno(39)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 40: return stat.S_ISREG(st.st_mode)
					πF.SetLineno(40)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µst, "st"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µst, ßst_mode, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßS_ISREG, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßisfile.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 46: def isdir(s):
			πF.SetLineno(46)
			πTemp009 = make([]πg.Param, 1)
			πTemp009[0] = πg.Param{Name: "s", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("isdir", "build/src/__python__/genericpath.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µst *πg.Object = πg.UnboundLocal; _ = µst
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 47: """Return true if the pathname refers to an existing directory."""
					πF.SetLineno(47)
					// line 48: try:
					πF.SetLineno(48)
					πF.PushCheckpoint(2)
					// line 49: st = os.stat(s)
					πF.SetLineno(49)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstat, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µst = πTemp002
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßerror, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 50: except os.error:
					πF.SetLineno(50)
				Label3:
					// line 51: return False
					πF.SetLineno(51)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 52: return stat.S_ISDIR(st.st_mode)
					πF.SetLineno(52)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µst, "st"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µst, ßst_mode, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßS_ISDIR, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßisdir.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 55: def getsize(filename):
			πF.SetLineno(55)
			πTemp009 = make([]πg.Param, 1)
			πTemp009[0] = πg.Param{Name: "filename", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("getsize", "build/src/__python__/genericpath.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfilename *πg.Object = πArgs[0]; _ = µfilename
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
					// line 56: """Return the size of a file, reported by os.stat()."""
					πF.SetLineno(56)
					// line 57: return os.stat(filename).st_size
					πF.SetLineno(57)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp001[0] = µfilename
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstat, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßst_size, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßgetsize.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 60: def getmtime(filename):
			πF.SetLineno(60)
			πTemp009 = make([]πg.Param, 1)
			πTemp009[0] = πg.Param{Name: "filename", Def: nil}
			πTemp011 = πg.NewFunction(πg.NewCode("getmtime", "build/src/__python__/genericpath.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfilename *πg.Object = πArgs[0]; _ = µfilename
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
					// line 61: """Return the last modification time of a file, reported by os.stat()."""
					πF.SetLineno(61)
					// line 62: return os.stat(filename).st_mtime
					πF.SetLineno(62)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp001[0] = µfilename
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstat, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßst_mtime, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßgetmtime.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 65: def getatime(filename):
			πF.SetLineno(65)
			πTemp009 = make([]πg.Param, 1)
			πTemp009[0] = πg.Param{Name: "filename", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("getatime", "build/src/__python__/genericpath.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfilename *πg.Object = πArgs[0]; _ = µfilename
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
					// line 66: """Return the last access time of a file, reported by os.stat()."""
					πF.SetLineno(66)
					// line 67: return os.stat(filename).st_atime
					πF.SetLineno(67)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp001[0] = µfilename
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstat, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßst_atime, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßgetatime.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 70: def getctime(filename):
			πF.SetLineno(70)
			πTemp009 = make([]πg.Param, 1)
			πTemp009[0] = πg.Param{Name: "filename", Def: nil}
			πTemp013 = πg.NewFunction(πg.NewCode("getctime", "build/src/__python__/genericpath.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfilename *πg.Object = πArgs[0]; _ = µfilename
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
					// line 71: """Return the metadata change time of a file, reported by os.stat()."""
					πF.SetLineno(71)
					// line 72: return os.stat(filename).st_ctime
					πF.SetLineno(72)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp001[0] = µfilename
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstat, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßst_ctime, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßgetctime.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 76: def commonprefix(m):
			πF.SetLineno(76)
			πTemp009 = make([]πg.Param, 1)
			πTemp009[0] = πg.Param{Name: "m", Def: nil}
			πTemp014 = πg.NewFunction(πg.NewCode("commonprefix", "build/src/__python__/genericpath.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µm *πg.Object = πArgs[0]; _ = µm
				var µs1 *πg.Object = πg.UnboundLocal; _ = µs1
				var µs2 *πg.Object = πg.UnboundLocal; _ = µs2
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µc *πg.Object = πg.UnboundLocal; _ = µc
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 3: goto Label3
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					// line 77: "Given a list of pathnames, returns the longest common leading component"
					πF.SetLineno(77)
					if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µm); πE != nil {
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
					// line 78: if not m: return ''
					πF.SetLineno(78)
				Label1:
					// line 78: if not m: return ''
					πF.SetLineno(78)
					πR = ß.ToObject()
					continue
					goto Label2
				Label2:
					// line 79: s1 = min(m)
					πF.SetLineno(79)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
						continue
					}
					πTemp003[0] = µm
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µs1 = πTemp004
					// line 80: s2 = max(m)
					πF.SetLineno(80)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
						continue
					}
					πTemp003[0] = µm
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µs2 = πTemp004
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
						continue
					}
					πTemp003[0] = µs1
					if πTemp004, πE = πg.ResolveGlobal(πF, ßenumerate); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp007}}}, πTemp004); πE != nil {
							continue
						}
						µi = πTemp005
						µc = πTemp007
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(3)            
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp005 = µi
					if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µs2, πTemp005); πE != nil {
						continue
					}
					if πTemp004, πE = πg.NE(πF, µc, πTemp007); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label6
					}
					goto Label7
					// line 82: if c != s2[i]:
					πF.SetLineno(82)
				Label6:
					// line 83: return s1[:i]
					πF.SetLineno(83)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µi, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µs1, πTemp004); πE != nil {
						continue
					}
					πR = πTemp005
					continue
					goto Label7
				Label7:
					continue
				Label4:
					if πE != nil || πR != nil {
						continue
					}
				Label5:
					// line 84: return s1
					πF.SetLineno(84)
					if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
						continue
					}
					πR = µs1
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßcommonprefix.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 93: def _splitext(p, sep, altsep, extsep):
			πF.SetLineno(93)
			πTemp009 = make([]πg.Param, 4)
			πTemp009[0] = πg.Param{Name: "p", Def: nil}
			πTemp009[1] = πg.Param{Name: "sep", Def: nil}
			πTemp009[2] = πg.Param{Name: "altsep", Def: nil}
			πTemp009[3] = πg.Param{Name: "extsep", Def: nil}
			πTemp015 = πg.NewFunction(πg.NewCode("_splitext", "build/src/__python__/genericpath.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µp *πg.Object = πArgs[0]; _ = µp
				var µsep *πg.Object = πArgs[1]; _ = µsep
				var µaltsep *πg.Object = πArgs[2]; _ = µaltsep
				var µextsep *πg.Object = πArgs[3]; _ = µextsep
				var µsepIndex *πg.Object = πg.UnboundLocal; _ = µsepIndex
				var µaltsepIndex *πg.Object = πg.UnboundLocal; _ = µaltsepIndex
				var µdotIndex *πg.Object = πg.UnboundLocal; _ = µdotIndex
				var µfilenameIndex *πg.Object = πg.UnboundLocal; _ = µfilenameIndex
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 5: goto Label5
					case 6: goto Label6
					default: panic("unexpected function state")
					}
					// line 94: """Split the extension from a pathname.
					πF.SetLineno(94)
					// line 99: sepIndex = p.rfind(sep)
					πF.SetLineno(99)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsep, "sep"); πE != nil {
						continue
					}
					πTemp001[0] = µsep
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µp, ßrfind, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µsepIndex = πTemp003
					if πE = πg.CheckLocal(πF, µaltsep, "altsep"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µaltsep); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 100: if altsep:
					πF.SetLineno(100)
				Label1:
					// line 101: altsepIndex = p.rfind(altsep)
					πF.SetLineno(101)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µaltsep, "altsep"); πE != nil {
						continue
					}
					πTemp001[0] = µaltsep
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µp, ßrfind, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µaltsepIndex = πTemp003
					// line 102: sepIndex = max(sepIndex, altsepIndex)
					πF.SetLineno(102)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µsepIndex, "sepIndex"); πE != nil {
						continue
					}
					πTemp001[0] = µsepIndex
					if πE = πg.CheckLocal(πF, µaltsepIndex, "altsepIndex"); πE != nil {
						continue
					}
					πTemp001[1] = µaltsepIndex
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µsepIndex = πTemp003
					goto Label2
				Label2:
					// line 104: dotIndex = p.rfind(extsep)
					πF.SetLineno(104)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µextsep, "extsep"); πE != nil {
						continue
					}
					πTemp001[0] = µextsep
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µp, ßrfind, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µdotIndex = πTemp003
					if πE = πg.CheckLocal(πF, µdotIndex, "dotIndex"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsepIndex, "sepIndex"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GT(πF, µdotIndex, µsepIndex); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 105: if dotIndex > sepIndex:
					πF.SetLineno(105)
				Label3:
					// line 107: filenameIndex = sepIndex + 1
					πF.SetLineno(107)
					if πE = πg.CheckLocal(πF, µsepIndex, "sepIndex"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µsepIndex, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µfilenameIndex = πTemp002
					// line 108: while filenameIndex < dotIndex:
					πF.SetLineno(108)
					πF.PushCheckpoint(6)
					πTemp004 = false
				Label5:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label7
					}
					if πE = πg.CheckLocal(πF, µfilenameIndex, "filenameIndex"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdotIndex, "dotIndex"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LT(πF, µfilenameIndex, µdotIndex); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(5)            
					if πE = πg.CheckLocal(πF, µfilenameIndex, "filenameIndex"); πE != nil {
						continue
					}
					πTemp003 = µfilenameIndex
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µp, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µextsep, "extsep"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.NE(πF, πTemp006, µextsep); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label8
					}
					goto Label9
					// line 109: if p[filenameIndex] != extsep:
					πF.SetLineno(109)
				Label8:
					// line 110: return p[:dotIndex], p[dotIndex:]
					πF.SetLineno(110)
					if πE = πg.CheckLocal(πF, µdotIndex, "dotIndex"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µdotIndex, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µp, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdotIndex, "dotIndex"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{µdotIndex, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µp, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
					πR = πTemp002
					continue
					goto Label9
				Label9:
					// line 111: filenameIndex += 1
					πF.SetLineno(111)
					if πE = πg.CheckLocal(πF, µfilenameIndex, "filenameIndex"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µfilenameIndex, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µfilenameIndex = πTemp002
					continue
				Label6:
					if πE != nil || πR != nil {
						continue
					}
				Label7:
					goto Label4
				Label4:
					// line 113: return p, ''
					πF.SetLineno(113)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µp, ß.ToObject()).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ß_splitext.ToObject(), πTemp015); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("genericpath", Code)
}
