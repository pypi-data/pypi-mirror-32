package fnmatch
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/fnmatch.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßKeyError := πg.InternStr("KeyError")
		ßNone := πg.InternStr("None")
		ß_MAXCACHE := πg.InternStr("_MAXCACHE")
		ß__all__ := πg.InternStr("__all__")
		ß_cache := πg.InternStr("_cache")
		ß_purge := πg.InternStr("_purge")
		ßappend := πg.InternStr("append")
		ßcompile := πg.InternStr("compile")
		ßescape := πg.InternStr("escape")
		ßfilter := πg.InternStr("filter")
		ßfnmatch := πg.InternStr("fnmatch")
		ßfnmatchcase := πg.InternStr("fnmatchcase")
		ßglobals := πg.InternStr("globals")
		ßlen := πg.InternStr("len")
		ßmatch := πg.InternStr("match")
		ßnormcase := πg.InternStr("normcase")
		ßpath := πg.InternStr("path")
		ßre := πg.InternStr("re")
		ßreplace := πg.InternStr("replace")
		ßtranslate := πg.InternStr("translate")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Dict
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """Filename matching with shell patterns.
			πF.SetLineno(1)
			// line 13: import re
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 15: __all__ = ["filter", "fnmatch", "fnmatchcase", "translate"]
			πF.SetLineno(15)
			πTemp002 = make([]*πg.Object, 4)
			πTemp002[0] = ßfilter.ToObject()
			πTemp002[1] = ßfnmatch.ToObject()
			πTemp002[2] = ßfnmatchcase.ToObject()
			πTemp002[3] = ßtranslate.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 17: _cache = {}
			πF.SetLineno(17)
			πTemp003 = πg.NewDict()
			πTemp001 = πTemp003.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_cache.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: _MAXCACHE = 100
			πF.SetLineno(18)
			if πE = πF.Globals().SetItem(πF, ß_MAXCACHE.ToObject(), πg.NewInt(100).ToObject()); πE != nil {
				continue
			}
			// line 20: def _purge():
			πF.SetLineno(20)
			πTemp004 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("_purge", "build/src/__python__/fnmatch.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Dict
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
					// line 21: """Clear the pattern cache"""
					πF.SetLineno(21)
					// line 23: globals()['_cache'] = {}
					πF.SetLineno(23)
					πTemp001 = πg.NewDict()
					πTemp002 = πTemp001.ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßglobals); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp004 = ß_cache.ToObject()
					if πE = πg.SetItem(πF, πTemp005, πTemp004, πTemp003); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_purge.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 25: def fnmatch(name, pat):
			πF.SetLineno(25)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "name", Def: nil}
			πTemp004[1] = πg.Param{Name: "pat", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("fnmatch", "build/src/__python__/fnmatch.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]; _ = µname
				var µpat *πg.Object = πArgs[1]; _ = µpat
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
					// line 26: """Test whether FILENAME matches PATTERN.
					πF.SetLineno(26)
					// line 44: return fnmatchcase(name, pat)
					πF.SetLineno(44)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001[0] = µname
					if πE = πg.CheckLocal(πF, µpat, "pat"); πE != nil {
						continue
					}
					πTemp001[1] = µpat
					if πTemp002, πE = πg.ResolveGlobal(πF, ßfnmatchcase); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßfnmatch.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 46: def filter(names, pat):
			πF.SetLineno(46)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "names", Def: nil}
			πTemp004[1] = πg.Param{Name: "pat", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("filter", "build/src/__python__/fnmatch.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µnames *πg.Object = πArgs[0]; _ = µnames
				var µpat *πg.Object = πArgs[1]; _ = µpat
				var µos *πg.Object = πg.UnboundLocal; _ = µos
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µre_pat *πg.Object = πg.UnboundLocal; _ = µre_pat
				var µres *πg.Object = πg.UnboundLocal; _ = µres
				var µmatch *πg.Object = πg.UnboundLocal; _ = µmatch
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.BaseException
				_ = πTemp005
				var πTemp006 *πg.Traceback
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 *πg.Dict
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 bool
				_ = πTemp010
				var πTemp011 []*πg.Object
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 9: goto Label9
					case 2: goto Label2
					case 10: goto Label10
					case 14: goto Label14
					case 15: goto Label15
					default: panic("unexpected function state")
					}
					// line 47: """Return the subset of the list NAMES that match PAT"""
					πF.SetLineno(47)
					// line 48: import os
					πF.SetLineno(48)
					if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					µos = πTemp001
					// line 50: result=[]
					πF.SetLineno(50)
					πTemp002 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp002...).ToObject()
					µresult = πTemp001
					// line 52: try:
					πF.SetLineno(52)
					πF.PushCheckpoint(2)
					// line 53: re_pat = _cache[pat]
					πF.SetLineno(53)
					if πE = πg.CheckLocal(πF, µpat, "pat"); πE != nil {
						continue
					}
					πTemp001 = µpat
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_cache); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
						continue
					}
					µre_pat = πTemp003
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp005, πTemp006 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label3
					}
					πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
					continue
					// line 54: except KeyError:
					πF.SetLineno(54)
				Label3:
					// line 55: res = translate(pat)
					πF.SetLineno(55)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpat, "pat"); πE != nil {
						continue
					}
					πTemp002[0] = µpat
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtranslate); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µres = πTemp003
					πTemp002 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_cache); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_MAXCACHE); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GE(πF, πTemp004, πTemp003); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label4
					}
					goto Label5
					// line 56: if len(_cache) >= _MAXCACHE:
					πF.SetLineno(56)
				Label4:
					// line 58: globals()['_cache'] = {}
					πF.SetLineno(58)
					πTemp008 = πg.NewDict()
					πTemp001 = πTemp008.ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßglobals); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp004 = ß_cache.ToObject()
					if πE = πg.SetItem(πF, πTemp009, πTemp004, πTemp003); πE != nil {
						continue
					}
					goto Label5
				Label5:
					// line 59: _cache[pat] = re_pat = re.compile(res)
					πF.SetLineno(59)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
						continue
					}
					πTemp002[0] = µres
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
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_cache); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpat, "pat"); πE != nil {
						continue
					}
					πTemp009 = µpat
					if πE = πg.SetItem(πF, πTemp004, πTemp009, πTemp003); πE != nil {
						continue
					}
					µre_pat = πTemp001
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 60: match = re_pat.match
					πF.SetLineno(60)
					if πE = πg.CheckLocal(πF, µre_pat, "re_pat"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µre_pat, ßmatch, nil); πE != nil {
						continue
					}
					µmatch = πTemp001
					if πTemp007, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label6
					}
					goto Label7
					// line 62: if 1:
					πF.SetLineno(62)
				Label6:
					if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µnames); πE != nil {
						continue
					}
					πF.PushCheckpoint(10)
					πTemp007 = false
				Label9:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp007 {
						πF.PopCheckpoint()
						goto Label11
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
						µname = πTemp003
					}
					if πE != nil || !πTemp010 {
						continue
					}
					πF.PushCheckpoint(9)            
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp002[0] = µname
					if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
						continue
					}
					if πTemp003, πE = µmatch.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp010, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp010 {
						goto Label12
					}
					goto Label13
					// line 65: if match(name):
					πF.SetLineno(65)
				Label12:
					// line 66: result.append(name)
					πF.SetLineno(66)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp002[0] = µname
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label13
				Label13:
					continue
				Label10:
					if πE != nil || πR != nil {
						continue
					}
				Label11:
					goto Label8
				Label7:
					if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µnames); πE != nil {
						continue
					}
					πF.PushCheckpoint(15)
					πTemp007 = false
				Label14:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp007 {
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
						µname = πTemp003
					}
					if πE != nil || !πTemp010 {
						continue
					}
					πF.PushCheckpoint(14)            
					πTemp002 = πF.MakeArgs(1)
					πTemp011 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp011[0] = µname
					if πE = πg.CheckLocal(πF, µos, "os"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µos, ßpath, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßnormcase, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp011, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp011)
					πTemp002[0] = πTemp003
					if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
						continue
					}
					if πTemp003, πE = µmatch.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp010, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp010 {
						goto Label17
					}
					goto Label18
					// line 69: if match(os.path.normcase(name)):
					πF.SetLineno(69)
				Label17:
					// line 70: result.append(name)
					πF.SetLineno(70)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp002[0] = µname
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label18
				Label18:
					continue
				Label15:
					if πE != nil || πR != nil {
						continue
					}
				Label16:
					goto Label8
				Label8:
					// line 71: return result
					πF.SetLineno(71)
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
			if πE = πF.Globals().SetItem(πF, ßfilter.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 73: def fnmatchcase(name, pat):
			πF.SetLineno(73)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "name", Def: nil}
			πTemp004[1] = πg.Param{Name: "pat", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("fnmatchcase", "build/src/__python__/fnmatch.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]; _ = µname
				var µpat *πg.Object = πArgs[1]; _ = µpat
				var µre_pat *πg.Object = πg.UnboundLocal; _ = µre_pat
				var µres *πg.Object = πg.UnboundLocal; _ = µres
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
				var πTemp008 *πg.Dict
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 74: """Test whether FILENAME matches PATTERN, including case.
					πF.SetLineno(74)
					// line 80: try:
					πF.SetLineno(80)
					πF.PushCheckpoint(2)
					// line 81: re_pat = _cache[pat]
					πF.SetLineno(81)
					if πE = πg.CheckLocal(πF, µpat, "pat"); πE != nil {
						continue
					}
					πTemp001 = µpat
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_cache); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
						continue
					}
					µre_pat = πTemp002
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
					// line 82: except KeyError:
					πF.SetLineno(82)
				Label3:
					// line 83: res = translate(pat)
					πF.SetLineno(83)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpat, "pat"); πE != nil {
						continue
					}
					πTemp007[0] = µpat
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtranslate); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					µres = πTemp002
					πTemp007 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_cache); πE != nil {
						continue
					}
					πTemp007[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_MAXCACHE); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GE(πF, πTemp003, πTemp002); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label4
					}
					goto Label5
					// line 84: if len(_cache) >= _MAXCACHE:
					πF.SetLineno(84)
				Label4:
					// line 86: globals()['_cache'] = {}
					πF.SetLineno(86)
					πTemp008 = πg.NewDict()
					πTemp001 = πTemp008.ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßglobals); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp003 = ß_cache.ToObject()
					if πE = πg.SetItem(πF, πTemp009, πTemp003, πTemp002); πE != nil {
						continue
					}
					goto Label5
				Label5:
					// line 87: _cache[pat] = re_pat = re.compile(res)
					πF.SetLineno(87)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
						continue
					}
					πTemp007[0] = µres
					if πTemp001, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_cache); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpat, "pat"); πE != nil {
						continue
					}
					πTemp009 = µpat
					if πE = πg.SetItem(πF, πTemp003, πTemp009, πTemp002); πE != nil {
						continue
					}
					µre_pat = πTemp001
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 88: return re_pat.match(name) is not None
					πF.SetLineno(88)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp007[0] = µname
					if πE = πg.CheckLocal(πF, µre_pat, "re_pat"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µre_pat, ßmatch, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp003 != πTemp002).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßfnmatchcase.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 90: def translate(pat):
			πF.SetLineno(90)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "pat", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("translate", "build/src/__python__/fnmatch.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpat *πg.Object = πArgs[0]; _ = µpat
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µn *πg.Object = πg.UnboundLocal; _ = µn
				var µres *πg.Object = πg.UnboundLocal; _ = µres
				var µc *πg.Object = πg.UnboundLocal; _ = µc
				var µj *πg.Object = πg.UnboundLocal; _ = µj
				var µstuff *πg.Object = πg.UnboundLocal; _ = µstuff
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
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
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 16: goto Label16
					case 1: goto Label1
					case 2: goto Label2
					case 15: goto Label15
					default: panic("unexpected function state")
					}
					// line 91: """Translate a shell PATTERN to a regular expression.
					πF.SetLineno(91)
					// line 96: i, n = 0, len(pat)
					πF.SetLineno(96)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpat, "pat"); πE != nil {
						continue
					}
					πTemp002[0] = µpat
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001 = πg.NewTuple2(πg.NewInt(0).ToObject(), πTemp004).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
						continue
					}
					µi = πTemp003
					µn = πTemp004
					// line 97: res = ''
					πF.SetLineno(97)
					µres = ß.ToObject()
					// line 98: while i < n:
					πF.SetLineno(98)
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
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µi, µn); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 99: c = pat[i]
					πF.SetLineno(99)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp001 = µi
					if πE = πg.CheckLocal(πF, µpat, "pat"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µpat, πTemp001); πE != nil {
						continue
					}
					µc = πTemp003
					// line 100: i = i+1
					πF.SetLineno(100)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µi = πTemp001
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µc, πg.NewStr("*").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µc, πg.NewStr("?").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µc, πg.NewStr("[").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label6
					}
					goto Label7
					// line 101: if c == '*':
					πF.SetLineno(101)
				Label4:
					// line 102: res = res + '.*'
					πF.SetLineno(102)
					if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µres, πg.NewStr(".*").ToObject()); πE != nil {
						continue
					}
					µres = πTemp001
					goto Label8
					// line 103: elif c == '?':
					πF.SetLineno(103)
				Label5:
					// line 104: res = res + '.'
					πF.SetLineno(104)
					if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µres, πg.NewStr(".").ToObject()); πE != nil {
						continue
					}
					µres = πTemp001
					goto Label8
					// line 105: elif c == '[':
					πF.SetLineno(105)
				Label6:
					// line 106: j = i
					πF.SetLineno(106)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					µj = µi
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LT(πF, µj, µn); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label9
					}
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					πTemp004 = µj
					if πE = πg.CheckLocal(πF, µpat, "pat"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µpat, πTemp004); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp007, πg.NewStr("!").ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
				Label9:
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label10
					}
					goto Label11
					// line 107: if j < n and pat[j] == '!':
					πF.SetLineno(107)
				Label10:
					// line 108: j = j+1
					πF.SetLineno(108)
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µj, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µj = πTemp001
					goto Label11
				Label11:
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LT(πF, µj, µn); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label12
					}
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					πTemp004 = µj
					if πE = πg.CheckLocal(πF, µpat, "pat"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µpat, πTemp004); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp007, πg.NewStr("]").ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
				Label12:
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label13
					}
					goto Label14
					// line 109: if j < n and pat[j] == ']':
					πF.SetLineno(109)
				Label13:
					// line 110: j = j+1
					πF.SetLineno(110)
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µj, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µj = πTemp001
					goto Label14
				Label14:
					// line 111: while j < n and pat[j] != ']':
					πF.SetLineno(111)
					πF.PushCheckpoint(16)
					πTemp006 = false
				Label15:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label17
					}
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LT(πF, µj, µn); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp009 {
						goto Label18
					}
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					πTemp004 = µj
					if πE = πg.CheckLocal(πF, µpat, "pat"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µpat, πTemp004); πE != nil {
						continue
					}
					if πTemp003, πE = πg.NE(πF, πTemp007, πg.NewStr("]").ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
				Label18:
					if πTemp008, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(15)            
					// line 112: j = j+1
					πF.SetLineno(112)
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µj, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µj = πTemp001
					continue
				Label16:
					if πE != nil || πR != nil {
						continue
					}
				Label17:
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GE(πF, µj, µn); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label19
					}
					goto Label20
					// line 113: if j >= n:
					πF.SetLineno(113)
				Label19:
					// line 114: res = res + '\\['
					πF.SetLineno(114)
					if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µres, πg.NewStr("\\[").ToObject()); πE != nil {
						continue
					}
					µres = πTemp001
					goto Label21
				Label20:
					// line 116: stuff = pat[i:j].replace('\\','\\\\')
					πF.SetLineno(116)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewStr("\\").ToObject()
					πTemp002[1] = πg.NewStr("\\\\").ToObject()
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µi, µj, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpat, "pat"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µpat, πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßreplace, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µstuff = πTemp003
					// line 117: i = j+1
					πF.SetLineno(117)
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µj, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µi = πTemp001
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µstuff, "stuff"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µstuff, πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewStr("!").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label22
					}
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µstuff, "stuff"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µstuff, πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewStr("^").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label23
					}
					goto Label24
					// line 118: if stuff[0] == '!':
					πF.SetLineno(118)
				Label22:
					// line 119: stuff = '^' + stuff[1:]
					πF.SetLineno(119)
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstuff, "stuff"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µstuff, πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πg.NewStr("^").ToObject(), πTemp004); πE != nil {
						continue
					}
					µstuff = πTemp001
					goto Label24
					// line 120: elif stuff[0] == '^':
					πF.SetLineno(120)
				Label23:
					// line 121: stuff = '\\' + stuff
					πF.SetLineno(121)
					if πE = πg.CheckLocal(πF, µstuff, "stuff"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πg.NewStr("\\").ToObject(), µstuff); πE != nil {
						continue
					}
					µstuff = πTemp001
					goto Label24
				Label24:
					// line 122: res = '%s[%s]' % (res, stuff)
					πF.SetLineno(122)
					if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstuff, "stuff"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µres, µstuff).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s[%s]").ToObject(), πTemp003); πE != nil {
						continue
					}
					µres = πTemp001
					goto Label21
				Label21:
					goto Label8
				Label7:
					// line 124: res = res + re.escape(c)
					πF.SetLineno(124)
					if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
						continue
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp002[0] = µc
					if πTemp003, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßescape, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Add(πF, µres, πTemp003); πE != nil {
						continue
					}
					µres = πTemp001
					goto Label8
				Label8:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 125: return res + '\Z(?ms)'
					πF.SetLineno(125)
					if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µres, πg.NewStr("\\Z(?ms)").ToObject()); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßtranslate.ToObject(), πTemp008); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("fnmatch", Code)
}
