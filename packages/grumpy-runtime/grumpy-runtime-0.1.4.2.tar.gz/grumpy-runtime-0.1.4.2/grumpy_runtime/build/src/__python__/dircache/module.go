package dircache
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/dircache.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßKeyError := πg.InternStr("KeyError")
		ß__all__ := πg.InternStr("__all__")
		ßannotate := πg.InternStr("annotate")
		ßcache := πg.InternStr("cache")
		ßisdir := πg.InternStr("isdir")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßlistdir := πg.InternStr("listdir")
		ßopendir := πg.InternStr("opendir")
		ßos := πg.InternStr("os")
		ßpath := πg.InternStr("path")
		ßrange := πg.InternStr("range")
		ßreset := πg.InternStr("reset")
		ßsort := πg.InternStr("sort")
		ßst_mtime := πg.InternStr("st_mtime")
		ßstat := πg.InternStr("stat")
		ßwarnpy3k := πg.InternStr("warnpy3k")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 πg.KWArgs
		_ = πTemp004
		var πTemp005 *πg.Dict
		_ = πTemp005
		var πTemp006 []πg.Param
		_ = πTemp006
		var πTemp007 *πg.Object
		_ = πTemp007
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """Read and cache directory listings.
			πF.SetLineno(1)
			// line 6: from warnings import warnpy3k
			πF.SetLineno(6)
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
			// line 7: warnpy3k("the dircache module has been removed in Python 3.0", stacklevel=2)
			πF.SetLineno(7)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("the dircache module has been removed in Python 3.0").ToObject()
			πTemp004 = πg.KWArgs{
				{"stacklevel", πg.NewInt(2).ToObject()},
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßwarnpy3k); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, πTemp004); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 8: del warnpy3k
			πF.SetLineno(8)
			if πE = πg.DelVar(πF, πF.Globals(), ßwarnpy3k); πE != nil {
				continue
			}
			// line 10: import os
			πF.SetLineno(10)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 12: __all__ = ["listdir", "opendir", "annotate", "reset"]
			πF.SetLineno(12)
			πTemp002 = make([]*πg.Object, 4)
			πTemp002[0] = ßlistdir.ToObject()
			πTemp002[1] = ßopendir.ToObject()
			πTemp002[2] = ßannotate.ToObject()
			πTemp002[3] = ßreset.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 14: cache = {}
			πF.SetLineno(14)
			πTemp005 = πg.NewDict()
			πTemp001 = πTemp005.ToObject()
			if πE = πF.Globals().SetItem(πF, ßcache.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 16: def reset():
			πF.SetLineno(16)
			πTemp006 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("reset", "build/src/__python__/dircache.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Dict
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
					// line 17: """Reset the cache completely."""
					πF.SetLineno(17)
					// line 18: global cache
					πF.SetLineno(18)
					// line 19: cache = {}
					πF.SetLineno(19)
					πTemp001 = πg.NewDict()
					πTemp002 = πTemp001.ToObject()
					if πE = πF.Globals().SetItem(πF, ßcache.ToObject(), πTemp002); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßreset.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 21: def listdir(path):
			πF.SetLineno(21)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "path", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("listdir", "build/src/__python__/dircache.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πArgs[0]; _ = µpath
				var µcached_mtime *πg.Object = πg.UnboundLocal; _ = µcached_mtime
				var µlist *πg.Object = πg.UnboundLocal; _ = µlist
				var µmtime *πg.Object = πg.UnboundLocal; _ = µmtime
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
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 22: """List directory contents, using cache."""
					πF.SetLineno(22)
					// line 23: try:
					πF.SetLineno(23)
					πF.PushCheckpoint(2)
					// line 24: cached_mtime, list = cache[path]
					πF.SetLineno(24)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp001 = µpath
					if πTemp003, πE = πg.ResolveGlobal(πF, ßcache); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
						continue
					}
					µcached_mtime = πTemp001
					µlist = πTemp003
					// line 25: del cache[path]
					πF.SetLineno(25)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßcache); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp002 = µpath
					if πE = πg.DelItem(πF, πTemp001, πTemp002); πE != nil {
						continue
					}
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
					// line 26: except KeyError:
					πF.SetLineno(26)
				Label3:
					// line 27: cached_mtime, list = -1, []
					πF.SetLineno(27)
					if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp007 = make([]*πg.Object, 0)
					πTemp003 = πg.NewList(πTemp007...).ToObject()
					πTemp001 = πg.NewTuple2(πTemp002, πTemp003).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}}}, πTemp001); πE != nil {
						continue
					}
					µcached_mtime = πTemp002
					µlist = πTemp003
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 28: mtime = os.stat(path).st_mtime
					πF.SetLineno(28)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp007[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstat, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßst_mtime, nil); πE != nil {
						continue
					}
					µmtime = πTemp002
					if πE = πg.CheckLocal(πF, µmtime, "mtime"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcached_mtime, "cached_mtime"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, µmtime, µcached_mtime); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label4
					}
					goto Label5
					// line 29: if mtime != cached_mtime:
					πF.SetLineno(29)
				Label4:
					// line 30: list = os.listdir(path)
					πF.SetLineno(30)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp007[0] = µpath
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßlistdir, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					µlist = πTemp001
					// line 31: list.sort()
					πF.SetLineno(31)
					if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µlist, ßsort, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					goto Label5
				Label5:
					// line 32: cache[path] = mtime, list
					πF.SetLineno(32)
					if πE = πg.CheckLocal(πF, µmtime, "mtime"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µmtime, µlist).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßcache); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp008 = µpath
					if πE = πg.SetItem(πF, πTemp003, πTemp008, πTemp002); πE != nil {
						continue
					}
					// line 33: return list
					πF.SetLineno(33)
					if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
						continue
					}
					πR = µlist
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßlistdir.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 35: opendir = listdir # XXX backward compatibility
			πF.SetLineno(35)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßlistdir); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßopendir.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 37: def annotate(head, list):
			πF.SetLineno(37)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "head", Def: nil}
			πTemp006[1] = πg.Param{Name: "list", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("annotate", "build/src/__python__/dircache.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µhead *πg.Object = πArgs[0]; _ = µhead
				var µlist *πg.Object = πArgs[1]; _ = µlist
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 38: """Add '/' suffixes to directories."""
					πF.SetLineno(38)
					πTemp002 = πF.MakeArgs(1)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
						continue
					}
					πTemp003[0] = µlist
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp002[0] = πTemp005
					if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
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
					if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
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
						µi = πTemp004
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(1)            
					πTemp002 = πF.MakeArgs(1)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µhead, "head"); πE != nil {
						continue
					}
					πTemp003[0] = µhead
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp004 = µi
					if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µlist, πTemp004); πE != nil {
						continue
					}
					πTemp003[1] = πTemp005
					if πTemp004, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßpath, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp002[0] = πTemp005
					if πTemp004, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßpath, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßisdir, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label4
					}
					goto Label5
					// line 40: if os.path.isdir(os.path.join(head, list[i])):
					πF.SetLineno(40)
				Label4:
					// line 41: list[i] = list[i] + '/'
					πF.SetLineno(41)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp005 = µi
					if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µlist, πTemp005); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, πTemp008, πg.NewStr("/").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp008 = µi
					if πE = πg.SetItem(πF, µlist, πTemp008, πTemp005); πE != nil {
						continue
					}
					goto Label5
				Label5:
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
			if πE = πF.Globals().SetItem(πF, ßannotate.ToObject(), πTemp007); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("dircache", Code)
}
