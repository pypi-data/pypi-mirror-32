package glob
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/glob.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßNameError := πg.InternStr("NameError")
		ßNone := πg.InternStr("None")
		ß__all__ := πg.InternStr("__all__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_unicode := πg.InternStr("_unicode")
		ßcompile := πg.InternStr("compile")
		ßcurdir := πg.InternStr("curdir")
		ßerror := πg.InternStr("error")
		ßfilter := πg.InternStr("filter")
		ßfnmatch := πg.InternStr("fnmatch")
		ßgetdefaultencoding := πg.InternStr("getdefaultencoding")
		ßgetfilesystemencoding := πg.InternStr("getfilesystemencoding")
		ßglob := πg.InternStr("glob")
		ßglob0 := πg.InternStr("glob0")
		ßglob1 := πg.InternStr("glob1")
		ßhas_magic := πg.InternStr("has_magic")
		ßiglob := πg.InternStr("iglob")
		ßisdir := πg.InternStr("isdir")
		ßisinstance := πg.InternStr("isinstance")
		ßjoin := πg.InternStr("join")
		ßlexists := πg.InternStr("lexists")
		ßlist := πg.InternStr("list")
		ßlistdir := πg.InternStr("listdir")
		ßmagic_check := πg.InternStr("magic_check")
		ßobject := πg.InternStr("object")
		ßos := πg.InternStr("os")
		ßpath := πg.InternStr("path")
		ßre := πg.InternStr("re")
		ßsearch := πg.InternStr("search")
		ßsplit := πg.InternStr("split")
		ßsys := πg.InternStr("sys")
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 2: goto Label2
			default: panic("unexpected function state")
			}
			// line 1: """Filename globbing utility."""
			πF.SetLineno(1)
			// line 3: import sys
			πF.SetLineno(3)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 4: import os
			πF.SetLineno(4)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 5: import re
			πF.SetLineno(5)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 6: import fnmatch
			πF.SetLineno(6)
			if πTemp002, πE = πg.ImportModule(πF, "fnmatch"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßfnmatch.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 8: try:
			πF.SetLineno(8)
			πF.PushCheckpoint(2)
			// line 9: _unicode = unicode
			πF.SetLineno(9)
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
			// line 10: except NameError:
			πF.SetLineno(10)
		Label3:
			// line 13: class _unicode(object):
			πF.SetLineno(13)
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
			_, πE = πg.NewCode("_unicode", "build/src/__python__/glob.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 14: pass
					πF.SetLineno(14)
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
			// line 16: __all__ = ["glob", "iglob"]
			πF.SetLineno(16)
			πTemp002 = make([]*πg.Object, 2)
			πTemp002[0] = ßglob.ToObject()
			πTemp002[1] = ßiglob.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: def glob(pathname):
			πF.SetLineno(18)
			πTemp009 = make([]πg.Param, 1)
			πTemp009[0] = πg.Param{Name: "pathname", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("glob", "build/src/__python__/glob.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpathname *πg.Object = πArgs[0]; _ = µpathname
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
					// line 19: """Return a list of paths matching a pathname pattern.
					πF.SetLineno(19)
					// line 27: return list(iglob(pathname))
					πF.SetLineno(27)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpathname, "pathname"); πE != nil {
						continue
					}
					πTemp002[0] = µpathname
					if πTemp003, πE = πg.ResolveGlobal(πF, ßiglob); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßglob.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 29: def iglob(pathname):
			πF.SetLineno(29)
			πTemp009 = make([]πg.Param, 1)
			πTemp009[0] = πg.Param{Name: "pathname", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("iglob", "build/src/__python__/glob.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpathname *πg.Object = πArgs[0]; _ = µpathname
				var µdirname *πg.Object = πg.UnboundLocal; _ = µdirname
				var µbasename *πg.Object = πg.UnboundLocal; _ = µbasename
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µdirs *πg.Object = πg.UnboundLocal; _ = µdirs
				var µglob_in_dir *πg.Object = πg.UnboundLocal; _ = µglob_in_dir
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
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 8: goto Label8
						case 9: goto Label9
						case 12: goto Label12
						case 13: goto Label13
						case 15: goto Label15
						case 23: goto Label23
						case 24: goto Label24
						case 26: goto Label26
						case 27: goto Label27
						case 29: goto Label29
						default: panic("unexpected function state")
						}
						// line 30: """Return an iterator which yields the paths matching a pathname pattern.
						πF.SetLineno(30)
						// line 38: dirname, basename = os.path.split(pathname)
						πF.SetLineno(38)
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µpathname, "pathname"); πE != nil {
							continue
						}
						πTemp001[0] = µpathname
						if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßsplit, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
							continue
						}
						µdirname = πTemp002
						µbasename = πTemp004
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µpathname, "pathname"); πE != nil {
							continue
						}
						πTemp001[0] = µpathname
						if πTemp003, πE = πg.ResolveGlobal(πF, ßhas_magic); πE != nil {
							continue
						}
						if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
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
						// line 39: if not has_magic(pathname):
						πF.SetLineno(39)
					Label1:
						if πE = πg.CheckLocal(πF, µbasename, "basename"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.IsTrue(πF, µbasename); πE != nil {
							continue
						}
						if πTemp005 {
							goto Label3
						}
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µdirname, "dirname"); πE != nil {
							continue
						}
						πTemp001[0] = µdirname
						if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßisdir, nil); πE != nil {
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
							goto Label4
						}
						goto Label5
						// line 40: if basename:
						πF.SetLineno(40)
					Label3:
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µpathname, "pathname"); πE != nil {
							continue
						}
						πTemp001[0] = µpathname
						if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßlexists, nil); πE != nil {
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
							goto Label6
						}
						goto Label7
						// line 41: if os.path.lexists(pathname):
						πF.SetLineno(41)
					Label6:
						// line 42: yield pathname
						πF.SetLineno(42)
						if πE = πg.CheckLocal(πF, µpathname, "pathname"); πE != nil {
							continue
						}
						πF.PushCheckpoint(8)
						return µpathname, nil
					Label8:
						πTemp002 = πSent
						goto Label7
					Label7:
						goto Label5
						// line 45: if os.path.isdir(dirname):
						πF.SetLineno(45)
					Label4:
						// line 46: yield pathname
						πF.SetLineno(46)
						if πE = πg.CheckLocal(πF, µpathname, "pathname"); πE != nil {
							continue
						}
						πF.PushCheckpoint(9)
						return µpathname, nil
					Label9:
						πTemp002 = πSent
						goto Label5
					Label5:
						// line 47: return
						πF.SetLineno(47)
						πR = πg.None
						continue
						goto Label2
					Label2:
						if πE = πg.CheckLocal(πF, µdirname, "dirname"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.IsTrue(πF, µdirname); πE != nil {
							continue
						}
						πTemp002 = πg.GetBool(!πTemp005).ToObject()
						if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
							continue
						}
						if πTemp005 {
							goto Label10
						}
						goto Label11
						// line 48: if not dirname:
						πF.SetLineno(48)
					Label10:
						πTemp001 = πF.MakeArgs(2)
						if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
							continue
						}
						if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßcurdir, nil); πE != nil {
							continue
						}
						πTemp001[0] = πTemp004
						if πE = πg.CheckLocal(πF, µbasename, "basename"); πE != nil {
							continue
						}
						πTemp001[1] = µbasename
						if πTemp003, πE = πg.ResolveGlobal(πF, ßglob1); πE != nil {
							continue
						}
						if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
							continue
						}
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
							µname = πTemp003
						}
						if πE != nil || !πTemp006 {
							continue
						}
						πF.PushCheckpoint(12)            
						// line 50: yield name
						πF.SetLineno(50)
						if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
							continue
						}
						πF.PushCheckpoint(15)
						return µname, nil
					Label15:
						πTemp003 = πSent
						continue
					Label13:
						if πE != nil || πR != nil {
							continue
						}
					Label14:
						// line 51: return
						πF.SetLineno(51)
						πR = πg.None
						continue
						goto Label11
					Label11:
						if πE = πg.CheckLocal(πF, µdirname, "dirname"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µpathname, "pathname"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.NE(πF, µdirname, µpathname); πE != nil {
							continue
						}
						πTemp002 = πTemp003
						if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
							continue
						}
						if !πTemp005 {
							goto Label16
						}
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µdirname, "dirname"); πE != nil {
							continue
						}
						πTemp001[0] = µdirname
						if πTemp003, πE = πg.ResolveGlobal(πF, ßhas_magic); πE != nil {
							continue
						}
						if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						πTemp002 = πTemp004
					Label16:
						if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
							continue
						}
						if πTemp005 {
							goto Label17
						}
						goto Label18
						// line 55: if dirname != pathname and has_magic(dirname):
						πF.SetLineno(55)
					Label17:
						// line 56: dirs = iglob(dirname)
						πF.SetLineno(56)
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µdirname, "dirname"); πE != nil {
							continue
						}
						πTemp001[0] = µdirname
						if πTemp002, πE = πg.ResolveGlobal(πF, ßiglob); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						µdirs = πTemp003
						goto Label19
					Label18:
						// line 58: dirs = [dirname]
						πF.SetLineno(58)
						πTemp001 = make([]*πg.Object, 1)
						if πE = πg.CheckLocal(πF, µdirname, "dirname"); πE != nil {
							continue
						}
						πTemp001[0] = µdirname
						πTemp002 = πg.NewList(πTemp001...).ToObject()
						µdirs = πTemp002
						goto Label19
					Label19:
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µbasename, "basename"); πE != nil {
							continue
						}
						πTemp001[0] = µbasename
						if πTemp002, πE = πg.ResolveGlobal(πF, ßhas_magic); πE != nil {
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
							goto Label20
						}
						goto Label21
						// line 59: if has_magic(basename):
						πF.SetLineno(59)
					Label20:
						// line 60: glob_in_dir = glob1
						πF.SetLineno(60)
						if πTemp002, πE = πg.ResolveGlobal(πF, ßglob1); πE != nil {
							continue
						}
						µglob_in_dir = πTemp002
						goto Label22
					Label21:
						// line 62: glob_in_dir = glob0
						πF.SetLineno(62)
						if πTemp002, πE = πg.ResolveGlobal(πF, ßglob0); πE != nil {
							continue
						}
						µglob_in_dir = πTemp002
						goto Label22
					Label22:
						if πE = πg.CheckLocal(πF, µdirs, "dirs"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.Iter(πF, µdirs); πE != nil {
							continue
						}
						πF.PushCheckpoint(24)
						πTemp005 = false
					Label23:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp005 {
							πF.PopCheckpoint()
							goto Label25
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
							µdirname = πTemp003
						}
						if πE != nil || !πTemp006 {
							continue
						}
						πF.PushCheckpoint(23)            
						πTemp001 = πF.MakeArgs(2)
						if πE = πg.CheckLocal(πF, µdirname, "dirname"); πE != nil {
							continue
						}
						πTemp001[0] = µdirname
						if πE = πg.CheckLocal(πF, µbasename, "basename"); πE != nil {
							continue
						}
						πTemp001[1] = µbasename
						if πE = πg.CheckLocal(πF, µglob_in_dir, "glob_in_dir"); πE != nil {
							continue
						}
						if πTemp004, πE = µglob_in_dir.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πTemp003, πE = πg.Iter(πF, πTemp004); πE != nil {
							continue
						}
						πF.PushCheckpoint(27)
						πTemp006 = false
					Label26:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp006 {
							πF.PopCheckpoint()
							goto Label28
						}
						if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
							µname = πTemp004
						}
						if πE != nil || !πTemp007 {
							continue
						}
						πF.PushCheckpoint(26)            
						// line 65: yield os.path.join(dirname, name)
						πF.SetLineno(65)
						πTemp001 = πF.MakeArgs(2)
						if πE = πg.CheckLocal(πF, µdirname, "dirname"); πE != nil {
							continue
						}
						πTemp001[0] = µdirname
						if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
							continue
						}
						πTemp001[1] = µname
						if πTemp004, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
							continue
						}
						if πTemp008, πE = πg.GetAttr(πF, πTemp004, ßpath, nil); πE != nil {
							continue
						}
						if πTemp004, πE = πg.GetAttr(πF, πTemp008, ßjoin, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						πF.PushCheckpoint(29)
						return πTemp008, nil
					Label29:
						πTemp004 = πSent
						continue
					Label27:
						if πE != nil || πR != nil {
							continue
						}
					Label28:
						continue
					Label24:
						if πE != nil || πR != nil {
							continue
						}
					Label25:
					}
					return nil, πE
				}).ToObject(), nil
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßiglob.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 71: def glob1(dirname, pattern):
			πF.SetLineno(71)
			πTemp009 = make([]πg.Param, 2)
			πTemp009[0] = πg.Param{Name: "dirname", Def: nil}
			πTemp009[1] = πg.Param{Name: "pattern", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("glob1", "build/src/__python__/glob.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µdirname *πg.Object = πArgs[0]; _ = µdirname
				var µpattern *πg.Object = πArgs[1]; _ = µpattern
				var µnames *πg.Object = πg.UnboundLocal; _ = µnames
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 *πg.BaseException
				_ = πTemp008
				var πTemp009 *πg.Traceback
				_ = πTemp009
				var πTemp010 []πg.Param
				_ = πTemp010
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 8: goto Label8
					default: panic("unexpected function state")
					}
					if πE = πg.CheckLocal(πF, µdirname, "dirname"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µdirname); πE != nil {
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
					// line 72: if not dirname:
					πF.SetLineno(72)
				Label1:
					// line 73: dirname = os.curdir
					πF.SetLineno(73)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßcurdir, nil); πE != nil {
						continue
					}
					µdirname = πTemp003
					goto Label2
				Label2:
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					πTemp004[0] = µpattern
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_unicode); πE != nil {
						continue
					}
					πTemp004[1] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp001 = πTemp005
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label3
					}
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µdirname, "dirname"); πE != nil {
						continue
					}
					πTemp004[0] = µdirname
					if πTemp005, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
						continue
					}
					πTemp004[1] = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp007).ToObject()
					πTemp001 = πTemp003
				Label3:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					goto Label5
					// line 74: if isinstance(pattern, _unicode) and not isinstance(dirname, unicode):
					πF.SetLineno(74)
				Label4:
					// line 75: dirname = unicode(dirname, sys.getfilesystemencoding() or
					πF.SetLineno(75)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µdirname, "dirname"); πE != nil {
						continue
					}
					πTemp004[0] = µdirname
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßgetfilesystemencoding, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label6
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßgetdefaultencoding, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp003
				Label6:
					πTemp004[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µdirname = πTemp003
					goto Label5
				Label5:
					// line 77: try:
					πF.SetLineno(77)
					πF.PushCheckpoint(8)
					// line 78: names = os.listdir(dirname)
					πF.SetLineno(78)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdirname, "dirname"); πE != nil {
						continue
					}
					πTemp004[0] = µdirname
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßlistdir, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µnames = πTemp001
					πF.PopCheckpoint()
					goto Label7
				Label8:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßerror, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label9
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 79: except os.error:
					πF.SetLineno(79)
				Label9:
					// line 80: return []
					πF.SetLineno(80)
					πTemp004 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp004...).ToObject()
					πR = πTemp001
					continue
					πF.RestoreExc(nil, nil)
					goto Label7
				Label7:
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µpattern, πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, πTemp005, πg.NewStr(".").ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label10
					}
					goto Label11
					// line 81: if pattern[0] != '.':
					πF.SetLineno(81)
				Label10:
					// line 83: names = [x for x in names if x[0] != '.']
					πF.SetLineno(83)
					πTemp010 = make([]πg.Param, 0)
					πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/glob.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp006 *πg.Object
						_ = πTemp006
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
								if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µnames); πE != nil {
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
								πTemp005 = πg.NewInt(0).ToObject()
								if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetItem(πF, µx, πTemp005); πE != nil {
									continue
								}
								if πTemp004, πE = πg.NE(πF, πTemp006, πg.NewStr(".").ToObject()); πE != nil {
									continue
								}
								if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label4
								}
								goto Label5
								// line 83: names = [x for x in names if x[0] != '.']
								πF.SetLineno(83)
							Label4:
								// line 83: names = [x for x in names if x[0] != '.']
								πF.SetLineno(83)
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
					if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
						continue
					}
					µnames = πTemp001
					goto Label11
				Label11:
					// line 84: return fnmatch.filter(names, pattern)
					πF.SetLineno(84)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
						continue
					}
					πTemp004[0] = µnames
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					πTemp004[1] = µpattern
					if πTemp001, πE = πg.ResolveGlobal(πF, ßfnmatch); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßfilter, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
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
			if πE = πF.Globals().SetItem(πF, ßglob1.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 86: def glob0(dirname, basename):
			πF.SetLineno(86)
			πTemp009 = make([]πg.Param, 2)
			πTemp009[0] = πg.Param{Name: "dirname", Def: nil}
			πTemp009[1] = πg.Param{Name: "basename", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("glob0", "build/src/__python__/glob.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µdirname *πg.Object = πArgs[0]; _ = µdirname
				var µbasename *πg.Object = πArgs[1]; _ = µbasename
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					if πE = πg.CheckLocal(πF, µbasename, "basename"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µbasename, ß.ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					πTemp003 = πF.MakeArgs(1)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µdirname, "dirname"); πE != nil {
						continue
					}
					πTemp004[0] = µdirname
					if πE = πg.CheckLocal(πF, µbasename, "basename"); πE != nil {
						continue
					}
					πTemp004[1] = µbasename
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp003[0] = πTemp005
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßlexists, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 87: if basename == '':
					πF.SetLineno(87)
				Label1:
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdirname, "dirname"); πE != nil {
						continue
					}
					πTemp003[0] = µdirname
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßisdir, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					goto Label5
					// line 90: if os.path.isdir(dirname):
					πF.SetLineno(90)
				Label4:
					// line 91: return [basename]
					πF.SetLineno(91)
					πTemp003 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µbasename, "basename"); πE != nil {
						continue
					}
					πTemp003[0] = µbasename
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					πR = πTemp001
					continue
					goto Label5
				Label5:
					goto Label3
					// line 93: if os.path.lexists(os.path.join(dirname, basename)):
					πF.SetLineno(93)
				Label2:
					// line 94: return [basename]
					πF.SetLineno(94)
					πTemp003 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µbasename, "basename"); πE != nil {
						continue
					}
					πTemp003[0] = µbasename
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					πR = πTemp001
					continue
					goto Label3
				Label3:
					// line 95: return []
					πF.SetLineno(95)
					πTemp003 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp003...).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßglob0.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 98: magic_check = re.compile('[*?[]')
			πF.SetLineno(98)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("[*?[]").ToObject()
			if πTemp011, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßcompile, nil); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp012.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßmagic_check.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 100: def has_magic(s):
			πF.SetLineno(100)
			πTemp009 = make([]πg.Param, 1)
			πTemp009[0] = πg.Param{Name: "s", Def: nil}
			πTemp011 = πg.NewFunction(πg.NewCode("has_magic", "build/src/__python__/glob.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
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
					// line 101: return magic_check.search(s) is not None
					πF.SetLineno(101)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp002[0] = µs
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmagic_check); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsearch, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp003 != πTemp004).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßhas_magic.ToObject(), πTemp011); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("glob", Code)
}
