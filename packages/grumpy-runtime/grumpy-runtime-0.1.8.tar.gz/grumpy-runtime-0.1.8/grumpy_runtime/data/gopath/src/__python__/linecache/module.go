package linecache
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/linecache.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßAttributeError := πg.InternStr("AttributeError")
		ßIOError := πg.InternStr("IOError")
		ßImportError := πg.InternStr("ImportError")
		ßMemoryError := πg.InternStr("MemoryError")
		ßNone := πg.InternStr("None")
		ßOSError := πg.InternStr("OSError")
		ßTypeError := πg.InternStr("TypeError")
		ß__all__ := πg.InternStr("__all__")
		ß__enter__ := πg.InternStr("__enter__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__loader__ := πg.InternStr("__loader__")
		ß__name__ := πg.InternStr("__name__")
		ßcache := πg.InternStr("cache")
		ßcheckcache := πg.InternStr("checkcache")
		ßclearcache := πg.InternStr("clearcache")
		ßendswith := πg.InternStr("endswith")
		ßerror := πg.InternStr("error")
		ßget := πg.InternStr("get")
		ßget_source := πg.InternStr("get_source")
		ßgetattr := πg.InternStr("getattr")
		ßgetline := πg.InternStr("getline")
		ßgetlines := πg.InternStr("getlines")
		ßisabs := πg.InternStr("isabs")
		ßjoin := πg.InternStr("join")
		ßkeys := πg.InternStr("keys")
		ßlen := πg.InternStr("len")
		ßopen := πg.InternStr("open")
		ßos := πg.InternStr("os")
		ßpath := πg.InternStr("path")
		ßrU := πg.InternStr("rU")
		ßreadlines := πg.InternStr("readlines")
		ßsplitlines := πg.InternStr("splitlines")
		ßst_mtime := πg.InternStr("st_mtime")
		ßst_size := πg.InternStr("st_size")
		ßstartswith := πg.InternStr("startswith")
		ßstat := πg.InternStr("stat")
		ßsys := πg.InternStr("sys")
		ßupdatecache := πg.InternStr("updatecache")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 []πg.Param
		_ = πTemp003
		var πTemp004 *πg.Object
		_ = πTemp004
		var πTemp005 *πg.Dict
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
			// line 1: """Cache lines from files.
			πF.SetLineno(1)
			// line 8: import sys
			πF.SetLineno(8)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 9: import os
			πF.SetLineno(9)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 11: __all__ = ["getline", "clearcache", "checkcache"]
			πF.SetLineno(11)
			πTemp002 = make([]*πg.Object, 3)
			πTemp002[0] = ßgetline.ToObject()
			πTemp002[1] = ßclearcache.ToObject()
			πTemp002[2] = ßcheckcache.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 13: def getline(filename, lineno, module_globals=None):
			πF.SetLineno(13)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "filename", Def: nil}
			πTemp003[1] = πg.Param{Name: "lineno", Def: nil}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[2] = πg.Param{Name: "module_globals", Def: πTemp004}
			πTemp001 = πg.NewFunction(πg.NewCode("getline", "build/src/__python__/linecache.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfilename *πg.Object = πArgs[0]; _ = µfilename
				var µlineno *πg.Object = πArgs[1]; _ = µlineno
				var µmodule_globals *πg.Object = πArgs[2]; _ = µmodule_globals
				var µlines *πg.Object = πg.UnboundLocal; _ = µlines
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
					// line 14: lines = getlines(filename, module_globals)
					πF.SetLineno(14)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp001[0] = µfilename
					if πE = πg.CheckLocal(πF, µmodule_globals, "module_globals"); πE != nil {
						continue
					}
					πTemp001[1] = µmodule_globals
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetlines); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µlines = πTemp003
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LE(πF, πg.NewInt(1).ToObject(), µlineno); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label1
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					πTemp001[0] = µlines
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.LE(πF, µlineno, πTemp005); πE != nil {
						continue
					}
				Label1:
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label2
					}
					goto Label3
					// line 15: if 1 <= lineno <= len(lines):
					πF.SetLineno(15)
				Label2:
					// line 16: return lines[lineno-1]
					πF.SetLineno(16)
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Sub(πF, µlineno, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µlines, πTemp002); πE != nil {
						continue
					}
					πR = πTemp003
					continue
					goto Label4
				Label3:
					// line 18: return ''
					πF.SetLineno(18)
					πR = ß.ToObject()
					continue
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
			if πE = πF.Globals().SetItem(πF, ßgetline.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 23: cache = {} # The cache
			πF.SetLineno(23)
			πTemp005 = πg.NewDict()
			πTemp004 = πTemp005.ToObject()
			if πE = πF.Globals().SetItem(πF, ßcache.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 26: def clearcache():
			πF.SetLineno(26)
			πTemp003 = make([]πg.Param, 0)
			πTemp004 = πg.NewFunction(πg.NewCode("clearcache", "build/src/__python__/linecache.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 27: """Clear the cache entirely."""
					πF.SetLineno(27)
					// line 29: global cache
					πF.SetLineno(29)
					// line 30: cache = {}
					πF.SetLineno(30)
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
			if πE = πF.Globals().SetItem(πF, ßclearcache.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 33: def getlines(filename, module_globals=None):
			πF.SetLineno(33)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "filename", Def: nil}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[1] = πg.Param{Name: "module_globals", Def: πTemp007}
			πTemp006 = πg.NewFunction(πg.NewCode("getlines", "build/src/__python__/linecache.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfilename *πg.Object = πArgs[0]; _ = µfilename
				var µmodule_globals *πg.Object = πArgs[1]; _ = µmodule_globals
				var πTemp001 *πg.Object
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
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 *πg.BaseException
				_ = πTemp008
				var πTemp009 *πg.Traceback
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					// line 34: """Get the lines for a file from the cache.
					πF.SetLineno(34)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßcache); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Contains(πF, πTemp002, µfilename); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp003).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 37: if filename in cache:
					πF.SetLineno(37)
				Label1:
					// line 38: return cache[filename][2]
					πF.SetLineno(38)
					πTemp001 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp004 = µfilename
					if πTemp006, πE = πg.ResolveGlobal(πF, ßcache); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					goto Label2
				Label2:
					// line 40: try:
					πF.SetLineno(40)
					πF.PushCheckpoint(4)
					// line 41: return updatecache(filename, module_globals)
					πF.SetLineno(41)
					πTemp007 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp007[0] = µfilename
					if πE = πg.CheckLocal(πF, µmodule_globals, "module_globals"); πE != nil {
						continue
					}
					πTemp007[1] = µmodule_globals
					if πTemp001, πE = πg.ResolveGlobal(πF, ßupdatecache); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πR = πTemp002
					continue
					πF.PopCheckpoint()
					goto Label3
				Label4:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßMemoryError); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label5
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 42: except MemoryError:
					πF.SetLineno(42)
				Label5:
					// line 43: clearcache()
					πF.SetLineno(43)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßclearcache); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 44: return []
					πF.SetLineno(44)
					πTemp007 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp007...).ToObject()
					πR = πTemp001
					continue
					πF.RestoreExc(nil, nil)
					goto Label3
				Label3:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßgetlines.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 47: def checkcache(filename=None):
			πF.SetLineno(47)
			πTemp003 = make([]πg.Param, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[0] = πg.Param{Name: "filename", Def: πTemp008}
			πTemp007 = πg.NewFunction(πg.NewCode("checkcache", "build/src/__python__/linecache.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfilename *πg.Object = πArgs[0]; _ = µfilename
				var µfilenames *πg.Object = πg.UnboundLocal; _ = µfilenames
				var µsize *πg.Object = πg.UnboundLocal; _ = µsize
				var µmtime *πg.Object = πg.UnboundLocal; _ = µmtime
				var µlines *πg.Object = πg.UnboundLocal; _ = µlines
				var µfullname *πg.Object = πg.UnboundLocal; _ = µfullname
				var µstat *πg.Object = πg.UnboundLocal; _ = µstat
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.BaseException
				_ = πTemp010
				var πTemp011 *πg.Traceback
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 11: goto Label11
					case 5: goto Label5
					case 6: goto Label6
					default: panic("unexpected function state")
					}
					// line 48: """Discard cache entries that are out of date.
					πF.SetLineno(48)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µfilename == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßcache); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Contains(πF, πTemp002, µfilename); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp003).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label2
					}
					goto Label3
					// line 51: if filename is None:
					πF.SetLineno(51)
				Label1:
					// line 52: filenames = cache.keys()
					πF.SetLineno(52)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßcache); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßkeys, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µfilenames = πTemp001
					goto Label4
					// line 54: if filename in cache:
					πF.SetLineno(54)
				Label2:
					// line 55: filenames = [filename]
					πF.SetLineno(55)
					πTemp004 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp004[0] = µfilename
					πTemp001 = πg.NewList(πTemp004...).ToObject()
					µfilenames = πTemp001
					goto Label4
				Label3:
					// line 57: return
					πF.SetLineno(57)
					πR = πg.None
					continue
					goto Label4
				Label4:
					if πE = πg.CheckLocal(πF, µfilenames, "filenames"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µfilenames); πE != nil {
						continue
					}
					πF.PushCheckpoint(6)
					πTemp003 = false
				Label5:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp003 {
						πF.PopCheckpoint()
						goto Label7
					}
					if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
						µfilename = πTemp002
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(5)            
					// line 60: size, mtime, lines, fullname = cache[filename]
					πF.SetLineno(60)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp002 = µfilename
					if πTemp007, πE = πg.ResolveGlobal(πF, ßcache); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}}}, πTemp006); πE != nil {
						continue
					}
					µsize = πTemp002
					µmtime = πTemp007
					µlines = πTemp008
					µfullname = πTemp009
					if πE = πg.CheckLocal(πF, µmtime, "mtime"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µmtime == πTemp006).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label8
					}
					goto Label9
					// line 61: if mtime is None:
					πF.SetLineno(61)
				Label8:
					// line 62: continue   # no-op for files loaded via a __loader__
					πF.SetLineno(62)
					continue
					goto Label9
				Label9:
					// line 63: try:
					πF.SetLineno(63)
					πF.PushCheckpoint(11)
					// line 64: stat = os.stat(fullname)
					πF.SetLineno(64)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfullname, "fullname"); πE != nil {
						continue
					}
					πTemp004[0] = µfullname
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßstat, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µstat = πTemp002
					πF.PopCheckpoint()
					goto Label10
				Label11:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp010, πTemp011 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßerror, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsInstance(πF, πTemp010.ToObject(), πTemp006); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label12
					}
					πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
					continue
					// line 65: except os.error:
					πF.SetLineno(65)
				Label12:
					// line 66: del cache[filename]
					πF.SetLineno(66)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßcache); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp006 = µfilename
					if πE = πg.DelItem(πF, πTemp002, πTemp006); πE != nil {
						continue
					}
					// line 67: continue
					πF.SetLineno(67)
					continue
					πF.RestoreExc(nil, nil)
					goto Label10
				Label10:
					if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstat, "stat"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µstat, ßst_size, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.NE(πF, µsize, πTemp007); πE != nil {
						continue
					}
					πTemp002 = πTemp006
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label13
					}
					if πE = πg.CheckLocal(πF, µmtime, "mtime"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstat, "stat"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µstat, ßst_mtime, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.NE(πF, µmtime, πTemp007); πE != nil {
						continue
					}
					πTemp002 = πTemp006
				Label13:
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label14
					}
					goto Label15
					// line 68: if size != stat.st_size or mtime != stat.st_mtime:
					πF.SetLineno(68)
				Label14:
					// line 69: del cache[filename]
					πF.SetLineno(69)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßcache); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp006 = µfilename
					if πE = πg.DelItem(πF, πTemp002, πTemp006); πE != nil {
						continue
					}
					goto Label15
				Label15:
					continue
				Label6:
					if πE != nil || πR != nil {
						continue
					}
				Label7:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßcheckcache.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 72: def updatecache(filename, module_globals=None):
			πF.SetLineno(72)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "filename", Def: nil}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[1] = πg.Param{Name: "module_globals", Def: πTemp009}
			πTemp008 = πg.NewFunction(πg.NewCode("updatecache", "build/src/__python__/linecache.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfilename *πg.Object = πArgs[0]; _ = µfilename
				var µmodule_globals *πg.Object = πArgs[1]; _ = µmodule_globals
				var µfullname *πg.Object = πg.UnboundLocal; _ = µfullname
				var µstat *πg.Object = πg.UnboundLocal; _ = µstat
				var µbasename *πg.Object = πg.UnboundLocal; _ = µbasename
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µloader *πg.Object = πg.UnboundLocal; _ = µloader
				var µget_source *πg.Object = πg.UnboundLocal; _ = µget_source
				var µdata *πg.Object = πg.UnboundLocal; _ = µdata
				var µdirname *πg.Object = πg.UnboundLocal; _ = µdirname
				var µfp *πg.Object = πg.UnboundLocal; _ = µfp
				var µlines *πg.Object = πg.UnboundLocal; _ = µlines
				var µsize *πg.Object = πg.UnboundLocal; _ = µsize
				var µmtime *πg.Object = πg.UnboundLocal; _ = µmtime
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
				var πTemp008 *πg.BaseException
				_ = πTemp008
				var πTemp009 *πg.Traceback
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 []πg.Param
				_ = πTemp011
				var πTemp012 *πg.Object
				_ = πTemp012
				var πTemp013 *πg.BaseException
				_ = πTemp013
				var πTemp014 *πg.Type
				_ = πTemp014
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 33: goto Label33
					case 34: goto Label34
					case 8: goto Label8
					case 17: goto Label17
					case 23: goto Label23
					case 24: goto Label24
					case 27: goto Label27
					case 30: goto Label30
					default: panic("unexpected function state")
					}
					// line 73: """Update a cache entry and return its list of lines.
					πF.SetLineno(73)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßcache); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Contains(πF, πTemp002, µfilename); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp003).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 77: if filename in cache:
					πF.SetLineno(77)
				Label1:
					// line 78: del cache[filename]
					πF.SetLineno(78)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßcache); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp002 = µfilename
					if πE = πg.DelItem(πF, πTemp001, πTemp002); πE != nil {
						continue
					}
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µfilename); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp004).ToObject()
					πTemp001 = πTemp002
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label3
					}
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = πg.NewStr("<").ToObject()
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µfilename, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp002 = πTemp007
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label4
					}
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = πg.NewStr(">").ToObject()
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µfilename, ßendswith, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp002 = πTemp007
				Label4:
					πTemp001 = πTemp002
				Label3:
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label5
					}
					goto Label6
					// line 79: if not filename or (filename.startswith('<') and filename.endswith('>')):
					πF.SetLineno(79)
				Label5:
					// line 80: return []
					πF.SetLineno(80)
					πTemp005 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp005...).ToObject()
					πR = πTemp001
					continue
					goto Label6
				Label6:
					// line 82: fullname = filename
					πF.SetLineno(82)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					µfullname = µfilename
					// line 83: try:
					πF.SetLineno(83)
					πF.PushCheckpoint(8)
					// line 84: stat = os.stat(fullname)
					πF.SetLineno(84)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfullname, "fullname"); πE != nil {
						continue
					}
					πTemp005[0] = µfullname
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstat, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µstat = πTemp001
					πF.PopCheckpoint()
					goto Label7
				Label8:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label9
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 85: except OSError:
					πF.SetLineno(85)
				Label9:
					// line 86: basename = filename
					πF.SetLineno(86)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					µbasename = µfilename
					if πE = πg.CheckLocal(πF, µmodule_globals, "module_globals"); πE != nil {
						continue
					}
					πTemp001 = µmodule_globals
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label10
					}
					if πE = πg.CheckLocal(πF, µmodule_globals, "module_globals"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, µmodule_globals, ß__loader__.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp004).ToObject()
					πTemp001 = πTemp002
				Label10:
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label11
					}
					goto Label12
					// line 89: if module_globals and '__loader__' in module_globals:
					πF.SetLineno(89)
				Label11:
					// line 90: name = module_globals.get('__name__')
					πF.SetLineno(90)
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = ß__name__.ToObject()
					if πE = πg.CheckLocal(πF, µmodule_globals, "module_globals"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µmodule_globals, ßget, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µname = πTemp002
					// line 91: loader = module_globals['__loader__']
					πF.SetLineno(91)
					πTemp001 = ß__loader__.ToObject()
					if πE = πg.CheckLocal(πF, µmodule_globals, "module_globals"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µmodule_globals, πTemp001); πE != nil {
						continue
					}
					µloader = πTemp002
					// line 92: get_source = getattr(loader, 'get_source', None)
					πF.SetLineno(92)
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µloader, "loader"); πE != nil {
						continue
					}
					πTemp005[0] = µloader
					πTemp005[1] = ßget_source.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp005[2] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µget_source = πTemp002
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001 = µname
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label13
					}
					if πE = πg.CheckLocal(πF, µget_source, "get_source"); πE != nil {
						continue
					}
					πTemp001 = µget_source
				Label13:
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label14
					}
					goto Label15
					// line 94: if name and get_source:
					πF.SetLineno(94)
				Label14:
					// line 95: try:
					πF.SetLineno(95)
					πF.PushCheckpoint(17)
					// line 96: data = get_source(name)
					πF.SetLineno(96)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp005[0] = µname
					if πE = πg.CheckLocal(πF, µget_source, "get_source"); πE != nil {
						continue
					}
					if πTemp001, πE = µget_source.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µdata = πTemp001
					πF.PopCheckpoint()
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µdata == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label18
					}
					goto Label19
					// line 100: if data is None:
					πF.SetLineno(100)
				Label18:
					// line 103: return []
					πF.SetLineno(103)
					πTemp005 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp005...).ToObject()
					πR = πTemp001
					continue
					goto Label19
				Label19:
					// line 104: cache[filename] = (
					πF.SetLineno(104)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					πTemp005[0] = µdata
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp011 = make([]πg.Param, 0)
					πTemp010 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/linecache.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µline *πg.Object = πg.UnboundLocal; _ = µline
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 bool
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
								if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µdata, ßsplitlines, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
									µline = πTemp002
								}
								if πE != nil || !πTemp005 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 106: [line+'\n' for line in data.splitlines()], fullname
								πF.SetLineno(106)
								if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Add(πF, µline, πg.NewStr("\n").ToObject()); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return πTemp002, nil
							Label4:
								πTemp003 = πSent
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
					if πTemp012, πE = πTemp010.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ListType.Call(πF, πg.Args{πTemp012}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfullname, "fullname"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple4(πTemp006, πTemp002, πTemp007, µfullname).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßcache); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp007 = µfilename
					if πE = πg.SetItem(πF, πTemp006, πTemp007, πTemp002); πE != nil {
						continue
					}
					// line 108: return cache[filename][2]
					πF.SetLineno(108)
					πTemp001 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp006 = µfilename
					if πTemp012, πE = πg.ResolveGlobal(πF, ßcache); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, πTemp012, πTemp006); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp007, πTemp001); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					goto Label16
				Label17:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp013, πTemp009 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp002, πTemp006).ToObject()
					if πTemp003, πE = πg.IsInstance(πF, πTemp013.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label20
					}
					πE = πF.Raise(πTemp013.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 97: except (ImportError, IOError):
					πF.SetLineno(97)
				Label20:
					// line 98: pass
					πF.SetLineno(98)
					πF.RestoreExc(nil, nil)
					goto Label16
				Label16:
					goto Label15
				Label15:
					goto Label12
				Label12:
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp005[0] = µfilename
					if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßisabs, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label21
					}
					goto Label22
					// line 112: if os.path.isabs(filename):
					πF.SetLineno(112)
				Label21:
					// line 113: return []
					πF.SetLineno(113)
					πTemp005 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp005...).ToObject()
					πR = πTemp001
					continue
					goto Label22
				Label22:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, πTemp006); πE != nil {
						continue
					}
					πF.PushCheckpoint(24)
					πTemp003 = false
				Label23:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp003 {
						πF.PopCheckpoint()
						goto Label25
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
						µdirname = πTemp002
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(23)            
					// line 118: try:
					πF.SetLineno(118)
					πF.PushCheckpoint(27)
					// line 119: fullname = os.path.join(dirname, basename)
					πF.SetLineno(119)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µdirname, "dirname"); πE != nil {
						continue
					}
					πTemp005[0] = µdirname
					if πE = πg.CheckLocal(πF, µbasename, "basename"); πE != nil {
						continue
					}
					πTemp005[1] = µbasename
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp006, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µfullname = πTemp006
					πF.PopCheckpoint()
					goto Label26
				Label27:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp013, πTemp009 = πF.ExcInfo()
					if πTemp006, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
					if πTemp004, πE = πg.IsInstance(πF, πTemp013.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label28
					}
					πE = πF.Raise(πTemp013.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 120: except (TypeError, AttributeError):
					πF.SetLineno(120)
				Label28:
					// line 122: continue
					πF.SetLineno(122)
					continue
					πF.RestoreExc(nil, nil)
					goto Label26
				Label26:
					// line 123: try:
					πF.SetLineno(123)
					πF.PushCheckpoint(30)
					// line 124: stat = os.stat(fullname)
					πF.SetLineno(124)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfullname, "fullname"); πE != nil {
						continue
					}
					πTemp005[0] = µfullname
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßstat, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µstat = πTemp002
					// line 125: break
					πF.SetLineno(125)
					πTemp003 = true
					continue
					πF.PopCheckpoint()
					goto Label29
				Label30:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp013, πTemp009 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßerror, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp013.ToObject(), πTemp006); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label31
					}
					πE = πF.Raise(πTemp013.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 126: except os.error:
					πF.SetLineno(126)
				Label31:
					// line 127: pass
					πF.SetLineno(127)
					πF.RestoreExc(nil, nil)
					goto Label29
				Label29:
					continue
				Label24:
					if πE != nil || πR != nil {
						continue
					}
					// line 129: return []
					πF.SetLineno(129)
					πTemp005 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp005...).ToObject()
					πR = πTemp002
					continue
				Label25:
					πF.RestoreExc(nil, nil)
					goto Label7
				Label7:
					// line 130: try:
					πF.SetLineno(130)
					πF.PushCheckpoint(33)
					// line 131: with open(fullname, 'rU') as fp:
					πF.SetLineno(131)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfullname, "fullname"); πE != nil {
						continue
					}
					πTemp005[0] = µfullname
					πTemp005[1] = ßrU.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp001, πE = πg.GetAttr(πF, πTemp002.Type().ToObject(), ß__exit__, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp002.Type().ToObject(), ß__enter__, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp006.Call(πF, πg.Args{πTemp002}, nil); πE != nil {
						continue
					}
					πF.PushCheckpoint(34)
					µfp = πTemp006
					// line 132: lines = fp.readlines()
					πF.SetLineno(132)
					if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µfp, ßreadlines, nil); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp007.Call(πF, nil, nil); πE != nil {
						continue
					}
					µlines = πTemp012
					πF.PopCheckpoint()
				Label34:
					πTemp008, πTemp009 = nil, nil
					if πE != nil {
						πTemp008, πTemp009 = πF.ExcInfo()
					}
					if πTemp008 != nil {
						πTemp014 = πTemp008.Type()
						if πTemp007, πE = πTemp001.Call(πF, πg.Args{πTemp002, πTemp014.ToObject(), πTemp008.ToObject(), πTemp009.ToObject()}, nil); πE != nil {
							continue
						}
					} else {
						if πTemp007, πE = πTemp001.Call(πF, πg.Args{πTemp002, πg.None, πg.None, πg.None}, nil); πE != nil {
							continue
						}
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					if πTemp008 != nil && πTemp003 != true {
						πE = πF.Raise(nil, nil, nil)
						continue
					}
					if πR != nil {
						continue
					}
					πF.PopCheckpoint()
					goto Label32
				Label33:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label35
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 133: except IOError:
					πF.SetLineno(133)
				Label35:
					// line 134: return []
					πF.SetLineno(134)
					πTemp005 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp005...).ToObject()
					πR = πTemp001
					continue
					πF.RestoreExc(nil, nil)
					goto Label32
				Label32:
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					πTemp001 = µlines
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label36
					}
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = πg.NewStr("\n").ToObject()
					if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp006 = πTemp007
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µlines, πTemp006); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp007, ßendswith, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp004, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp004).ToObject()
					πTemp001 = πTemp002
				Label36:
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label37
					}
					goto Label38
					// line 135: if lines and not lines[-1].endswith('\n'):
					πF.SetLineno(135)
				Label37:
					// line 136: lines[-1] += '\n'
					πF.SetLineno(136)
					if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp002
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µlines, πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, πTemp002, πg.NewStr("\n").ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp006 = πTemp007
					if πE = πg.SetItem(πF, µlines, πTemp006, πTemp001); πE != nil {
						continue
					}
					goto Label38
				Label38:
					// line 137: size, mtime = stat.st_size, stat.st_mtime
					πF.SetLineno(137)
					if πE = πg.CheckLocal(πF, µstat, "stat"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µstat, ßst_size, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstat, "stat"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µstat, ßst_mtime, nil); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp002, πTemp006).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp006}}}, πTemp001); πE != nil {
						continue
					}
					µsize = πTemp002
					µmtime = πTemp006
					// line 138: cache[filename] = size, mtime, lines, fullname
					πF.SetLineno(138)
					if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmtime, "mtime"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfullname, "fullname"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple4(µsize, µmtime, µlines, µfullname).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßcache); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp007 = µfilename
					if πE = πg.SetItem(πF, πTemp006, πTemp007, πTemp002); πE != nil {
						continue
					}
					// line 139: return lines
					πF.SetLineno(139)
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
			if πE = πF.Globals().SetItem(πF, ßupdatecache.ToObject(), πTemp008); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("linecache", Code)
}
