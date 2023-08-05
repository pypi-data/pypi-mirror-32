package getopt
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/getopt.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßException := πg.InternStr("Exception")
		ßFalse := πg.InternStr("False")
		ßGetoptError := πg.InternStr("GetoptError")
		ßNone := πg.InternStr("None")
		ßPOSIXLY_CORRECT := πg.InternStr("POSIXLY_CORRECT")
		ßTrue := πg.InternStr("True")
		ßValueError := πg.InternStr("ValueError")
		ß__all__ := πg.InternStr("__all__")
		ß__init__ := πg.InternStr("__init__")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__str__ := πg.InternStr("__str__")
		ßappend := πg.InternStr("append")
		ßargv := πg.InternStr("argv")
		ßbeta := πg.InternStr("beta")
		ßdo_longs := πg.InternStr("do_longs")
		ßdo_shorts := πg.InternStr("do_shorts")
		ßendswith := πg.InternStr("endswith")
		ßenviron := πg.InternStr("environ")
		ßerror := πg.InternStr("error")
		ßget := πg.InternStr("get")
		ßgetopt := πg.InternStr("getopt")
		ßgnu_getopt := πg.InternStr("gnu_getopt")
		ßindex := πg.InternStr("index")
		ßisinstance := πg.InternStr("isinstance")
		ßlen := πg.InternStr("len")
		ßlist := πg.InternStr("list")
		ßlong_has_args := πg.InternStr("long_has_args")
		ßmsg := πg.InternStr("msg")
		ßopt := πg.InternStr("opt")
		ßos := πg.InternStr("os")
		ßrange := πg.InternStr("range")
		ßshort_has_arg := πg.InternStr("short_has_arg")
		ßstartswith := πg.InternStr("startswith")
		ßstr := πg.InternStr("str")
		ßsys := πg.InternStr("sys")
		ßtype := πg.InternStr("type")
		var πTemp001 []*πg.Object
		_ = πTemp001
		var πTemp002 *πg.Object
		_ = πTemp002
		var πTemp003 *πg.Dict
		_ = πTemp003
		var πTemp004 *πg.Object
		_ = πTemp004
		var πTemp005 *πg.Object
		_ = πTemp005
		var πTemp006 []πg.Param
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
		var πTemp012 bool
		_ = πTemp012
		var πTemp013 []*πg.Object
		_ = πTemp013
		var πTemp014 *πg.Object
		_ = πTemp014
		var πTemp015 *πg.Object
		_ = πTemp015
		var πTemp016 []*πg.Object
		_ = πTemp016
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """Parser for command line options.
			πF.SetLineno(1)
			// line 34: __all__ = ["GetoptError","error","getopt","gnu_getopt"]
			πF.SetLineno(34)
			πTemp001 = make([]*πg.Object, 4)
			πTemp001[0] = ßGetoptError.ToObject()
			πTemp001[1] = ßerror.ToObject()
			πTemp001[2] = ßgetopt.ToObject()
			πTemp001[3] = ßgnu_getopt.ToObject()
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 36: import os
			πF.SetLineno(36)
			if πTemp001, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 38: class GetoptError(Exception):
			πF.SetLineno(38)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
				continue
			}
			πTemp001[0] = πTemp005
			πTemp003 = πg.NewDict()
			if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
				continue
			}
			_, πE = πg.NewCode("GetoptError", "build/src/__python__/getopt.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
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
					// line 39: opt = ''
					πF.SetLineno(39)
					if πE = πClass.SetItem(πF, ßopt.ToObject(), ß.ToObject()); πE != nil {
						continue
					}
					// line 40: msg = ''
					πF.SetLineno(40)
					if πE = πClass.SetItem(πF, ßmsg.ToObject(), ß.ToObject()); πE != nil {
						continue
					}
					// line 41: def __init__(self, msg, opt=''):
					πF.SetLineno(41)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "msg", Def: nil}
					πTemp002[2] = πg.Param{Name: "opt", Def: ß.ToObject()}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/getopt.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmsg *πg.Object = πArgs[1]; _ = µmsg
						var µopt *πg.Object = πArgs[2]; _ = µopt
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
							// line 42: self.msg = msg
							πF.SetLineno(42)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µmsg); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmsg, πTemp001); πE != nil {
								continue
							}
							// line 43: self.opt = opt
							πF.SetLineno(43)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µopt); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßopt, πTemp001); πE != nil {
								continue
							}
							// line 44: Exception.__init__(self, msg, opt)
							πF.SetLineno(44)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp002[1] = µmsg
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp002[2] = µopt
							if πTemp001, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 46: def __str__(self):
					πF.SetLineno(46)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__str__", "build/src/__python__/getopt.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 47: return self.msg
							πF.SetLineno(47)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmsg, nil); πE != nil {
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
			if πTemp004, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("GetoptError").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßGetoptError.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 49: error = GetoptError # backward compatibility
			πF.SetLineno(49)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßGetoptError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßerror.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 51: def getopt(args, shortopts, longopts = []):
			πF.SetLineno(51)
			πTemp006 = make([]πg.Param, 3)
			πTemp006[0] = πg.Param{Name: "args", Def: nil}
			πTemp006[1] = πg.Param{Name: "shortopts", Def: nil}
			πTemp001 = make([]*πg.Object, 0)
			πTemp004 = πg.NewList(πTemp001...).ToObject()
			πTemp006[2] = πg.Param{Name: "longopts", Def: πTemp004}
			πTemp002 = πg.NewFunction(πg.NewCode("getopt", "build/src/__python__/getopt.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargs *πg.Object = πArgs[0]; _ = µargs
				var µshortopts *πg.Object = πArgs[1]; _ = µshortopts
				var µlongopts *πg.Object = πArgs[2]; _ = µlongopts
				var µopts *πg.Object = πg.UnboundLocal; _ = µopts
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
				var πTemp006 bool
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
					case 4: goto Label4
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 52: """getopt(args, options[, long_options]) -> opts, args
					πF.SetLineno(52)
					// line 78: opts = []
					πF.SetLineno(78)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µopts = πTemp002
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlongopts, "longopts"); πE != nil {
						continue
					}
					πTemp001[0] = µlongopts
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ß.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Eq(πF, πTemp004, πTemp005); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label1
					}
					goto Label2
					// line 79: if type(longopts) == type(""):
					πF.SetLineno(79)
				Label1:
					// line 80: longopts = [longopts]
					πF.SetLineno(80)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µlongopts, "longopts"); πE != nil {
						continue
					}
					πTemp001[0] = µlongopts
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µlongopts = πTemp002
					goto Label3
				Label2:
					// line 82: longopts = list(longopts)
					πF.SetLineno(82)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlongopts, "longopts"); πE != nil {
						continue
					}
					πTemp001[0] = µlongopts
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µlongopts = πTemp003
					goto Label3
				Label3:
					// line 83: while args and args[0].startswith('-') and args[0] != '-':
					πF.SetLineno(83)
					πF.PushCheckpoint(5)
					πTemp006 = false
				Label4:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label6
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp002 = µargs
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label7
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("-").ToObject()
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µargs, πTemp003); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp002 = πTemp004
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label7
					}
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µargs, πTemp004); πE != nil {
						continue
					}
					if πTemp003, πE = πg.NE(πF, πTemp005, πg.NewStr("-").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label7:
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(4)            
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µargs, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp004, πg.NewStr("--").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label8
					}
					goto Label9
					// line 84: if args[0] == '--':
					πF.SetLineno(84)
				Label8:
					// line 85: args = args[1:]
					πF.SetLineno(85)
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp002); πE != nil {
						continue
					}
					µargs = πTemp003
					// line 86: break
					πF.SetLineno(86)
					πTemp006 = true
					continue
					goto Label9
				Label9:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("--").ToObject()
					πTemp002 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label10
					}
					goto Label11
					// line 87: if args[0].startswith('--'):
					πF.SetLineno(87)
				Label10:
					// line 88: opts, args = do_longs(opts, args[0][2:], longopts, args[1:])
					πF.SetLineno(88)
					πTemp001 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µopts, "opts"); πE != nil {
						continue
					}
					πTemp001[0] = µopts
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(2).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µargs, πTemp004); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πE = πg.CheckLocal(πF, µlongopts, "longopts"); πE != nil {
						continue
					}
					πTemp001[2] = µlongopts
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp002); πE != nil {
						continue
					}
					πTemp001[3] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdo_longs); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µopts = πTemp002
					µargs = πTemp004
					goto Label12
				Label11:
					// line 90: opts, args = do_shorts(opts, args[0][1:], shortopts, args[1:])
					πF.SetLineno(90)
					πTemp001 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µopts, "opts"); πE != nil {
						continue
					}
					πTemp001[0] = µopts
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µargs, πTemp004); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πE = πg.CheckLocal(πF, µshortopts, "shortopts"); πE != nil {
						continue
					}
					πTemp001[2] = µshortopts
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp002); πE != nil {
						continue
					}
					πTemp001[3] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdo_shorts); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µopts = πTemp002
					µargs = πTemp004
					goto Label12
				Label12:
					continue
				Label5:
					if πE != nil || πR != nil {
						continue
					}
				Label6:
					// line 92: return opts, args
					πF.SetLineno(92)
					if πE = πg.CheckLocal(πF, µopts, "opts"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µopts, µargs).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßgetopt.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 94: def gnu_getopt(args, shortopts, longopts = []):
			πF.SetLineno(94)
			πTemp006 = make([]πg.Param, 3)
			πTemp006[0] = πg.Param{Name: "args", Def: nil}
			πTemp006[1] = πg.Param{Name: "shortopts", Def: nil}
			πTemp001 = make([]*πg.Object, 0)
			πTemp005 = πg.NewList(πTemp001...).ToObject()
			πTemp006[2] = πg.Param{Name: "longopts", Def: πTemp005}
			πTemp004 = πg.NewFunction(πg.NewCode("gnu_getopt", "build/src/__python__/getopt.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargs *πg.Object = πArgs[0]; _ = µargs
				var µshortopts *πg.Object = πArgs[1]; _ = µshortopts
				var µlongopts *πg.Object = πArgs[2]; _ = µlongopts
				var µopts *πg.Object = πg.UnboundLocal; _ = µopts
				var µprog_args *πg.Object = πg.UnboundLocal; _ = µprog_args
				var µall_options_first *πg.Object = πg.UnboundLocal; _ = µall_options_first
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
				var πTemp009 *πg.Object
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 8: goto Label8
					case 9: goto Label9
					default: panic("unexpected function state")
					}
					// line 95: """getopt(args, options[, long_options]) -> opts, args
					πF.SetLineno(95)
					// line 109: opts = []
					πF.SetLineno(109)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µopts = πTemp002
					// line 110: prog_args = []
					πF.SetLineno(110)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µprog_args = πTemp002
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µlongopts, "longopts"); πE != nil {
						continue
					}
					πTemp001[0] = µlongopts
					if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
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
					// line 111: if isinstance(longopts, str):
					πF.SetLineno(111)
				Label1:
					// line 112: longopts = [longopts]
					πF.SetLineno(112)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µlongopts, "longopts"); πE != nil {
						continue
					}
					πTemp001[0] = µlongopts
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µlongopts = πTemp002
					goto Label3
				Label2:
					// line 114: longopts = list(longopts)
					πF.SetLineno(114)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlongopts, "longopts"); πE != nil {
						continue
					}
					πTemp001[0] = µlongopts
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µlongopts = πTemp003
					goto Label3
				Label3:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("+").ToObject()
					if πE = πg.CheckLocal(πF, µshortopts, "shortopts"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µshortopts, ßstartswith, nil); πE != nil {
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
						goto Label4
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ßPOSIXLY_CORRECT.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßenviron, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßget, nil); πE != nil {
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
						goto Label5
					}
					goto Label6
					// line 117: if shortopts.startswith('+'):
					πF.SetLineno(117)
				Label4:
					// line 118: shortopts = shortopts[1:]
					πF.SetLineno(118)
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µshortopts, "shortopts"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µshortopts, πTemp002); πE != nil {
						continue
					}
					µshortopts = πTemp003
					// line 119: all_options_first = True
					πF.SetLineno(119)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					µall_options_first = πTemp002
					goto Label7
					// line 120: elif os.environ.get("POSIXLY_CORRECT"):
					πF.SetLineno(120)
				Label5:
					// line 121: all_options_first = True
					πF.SetLineno(121)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					µall_options_first = πTemp002
					goto Label7
				Label6:
					// line 123: all_options_first = False
					πF.SetLineno(123)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					µall_options_first = πTemp002
					goto Label7
				Label7:
					// line 125: while args:
					πF.SetLineno(125)
					πF.PushCheckpoint(9)
					πTemp004 = false
				Label8:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label10
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µargs); πE != nil {
						continue
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(8)            
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µargs, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp006, πg.NewStr("--").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label11
					}
					goto Label12
					// line 126: if args[0] == '--':
					πF.SetLineno(126)
				Label11:
					// line 127: prog_args += args[1:]
					πF.SetLineno(127)
					if πE = πg.CheckLocal(πF, µprog_args, "prog_args"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µprog_args, πTemp003); πE != nil {
						continue
					}
					µprog_args = πTemp002
					// line 128: break
					πF.SetLineno(128)
					πTemp004 = true
					continue
					goto Label12
				Label12:
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(2).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					πTemp007 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µargs, πTemp007); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, πTemp008, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp006, πg.NewStr("--").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label13
					}
					if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(1).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					πTemp008 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, µargs, πTemp008); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, πTemp009, πTemp006); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp007, πg.NewStr("-").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label14
					}
					πTemp006 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µargs, πTemp006); πE != nil {
						continue
					}
					if πTemp003, πE = πg.NE(πF, πTemp007, πg.NewStr("-").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label14:
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label15
					}
					if πE = πg.CheckLocal(πF, µall_options_first, "all_options_first"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µall_options_first); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label16
					}
					goto Label17
					// line 130: if args[0][:2] == '--':
					πF.SetLineno(130)
				Label13:
					// line 131: opts, args = do_longs(opts, args[0][2:], longopts, args[1:])
					πF.SetLineno(131)
					πTemp001 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µopts, "opts"); πE != nil {
						continue
					}
					πTemp001[0] = µopts
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(2).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					πTemp006 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µargs, πTemp006); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πE = πg.CheckLocal(πF, µlongopts, "longopts"); πE != nil {
						continue
					}
					πTemp001[2] = µlongopts
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp002); πE != nil {
						continue
					}
					πTemp001[3] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdo_longs); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp006}}}, πTemp003); πE != nil {
						continue
					}
					µopts = πTemp002
					µargs = πTemp006
					goto Label18
					// line 132: elif args[0][:1] == '-' and args[0] != '-':
					πF.SetLineno(132)
				Label15:
					// line 133: opts, args = do_shorts(opts, args[0][1:], shortopts, args[1:])
					πF.SetLineno(133)
					πTemp001 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µopts, "opts"); πE != nil {
						continue
					}
					πTemp001[0] = µopts
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					πTemp006 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µargs, πTemp006); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πE = πg.CheckLocal(πF, µshortopts, "shortopts"); πE != nil {
						continue
					}
					πTemp001[2] = µshortopts
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp002); πE != nil {
						continue
					}
					πTemp001[3] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdo_shorts); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp006}}}, πTemp003); πE != nil {
						continue
					}
					µopts = πTemp002
					µargs = πTemp006
					goto Label18
					// line 135: if all_options_first:
					πF.SetLineno(135)
				Label16:
					// line 136: prog_args += args
					πF.SetLineno(136)
					if πE = πg.CheckLocal(πF, µprog_args, "prog_args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µprog_args, µargs); πE != nil {
						continue
					}
					µprog_args = πTemp002
					// line 137: break
					πF.SetLineno(137)
					πTemp004 = true
					continue
					goto Label18
				Label17:
					// line 139: prog_args.append(args[0])
					πF.SetLineno(139)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µprog_args, "prog_args"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µprog_args, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 140: args = args[1:]
					πF.SetLineno(140)
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp002); πE != nil {
						continue
					}
					µargs = πTemp003
					goto Label18
				Label18:
					continue
				Label9:
					if πE != nil || πR != nil {
						continue
					}
				Label10:
					// line 142: return opts, prog_args
					πF.SetLineno(142)
					if πE = πg.CheckLocal(πF, µopts, "opts"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µprog_args, "prog_args"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µopts, µprog_args).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßgnu_getopt.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 144: def do_longs(opts, opt, longopts, args):
			πF.SetLineno(144)
			πTemp006 = make([]πg.Param, 4)
			πTemp006[0] = πg.Param{Name: "opts", Def: nil}
			πTemp006[1] = πg.Param{Name: "opt", Def: nil}
			πTemp006[2] = πg.Param{Name: "longopts", Def: nil}
			πTemp006[3] = πg.Param{Name: "args", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("do_longs", "build/src/__python__/getopt.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µopts *πg.Object = πArgs[0]; _ = µopts
				var µopt *πg.Object = πArgs[1]; _ = µopt
				var µlongopts *πg.Object = πArgs[2]; _ = µlongopts
				var µargs *πg.Object = πArgs[3]; _ = µargs
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µoptarg *πg.Object = πg.UnboundLocal; _ = µoptarg
				var µhas_arg *πg.Object = πg.UnboundLocal; _ = µhas_arg
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
				var πTemp006 *πg.BaseException
				_ = πTemp006
				var πTemp007 *πg.Traceback
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 145: try:
					πF.SetLineno(145)
					πF.PushCheckpoint(2)
					// line 146: i = opt.index('=')
					πF.SetLineno(146)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("=").ToObject()
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µopt, ßindex, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µi = πTemp003
					πF.PopCheckpoint()
					// line 150: opt, optarg = opt[:i], opt[i+1:]
					πF.SetLineno(150)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µi, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µopt, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πTemp005, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µopt, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
						continue
					}
					µopt = πTemp003
					µoptarg = πTemp004
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp006, πTemp007 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label3
					}
					πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
					continue
					// line 147: except ValueError:
					πF.SetLineno(147)
				Label3:
					// line 148: optarg = None
					πF.SetLineno(148)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µoptarg = πTemp002
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 152: has_arg, opt = long_has_args(opt, longopts)
					πF.SetLineno(152)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					πTemp001[0] = µopt
					if πE = πg.CheckLocal(πF, µlongopts, "longopts"); πE != nil {
						continue
					}
					πTemp001[1] = µlongopts
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlong_has_args); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µhas_arg = πTemp002
					µopt = πTemp004
					if πE = πg.CheckLocal(πF, µhas_arg, "has_arg"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µhas_arg); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µoptarg, "optarg"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µoptarg != πTemp003).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label5
					}
					goto Label6
					// line 153: if has_arg:
					πF.SetLineno(153)
				Label4:
					if πE = πg.CheckLocal(πF, µoptarg, "optarg"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µoptarg == πTemp003).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label7
					}
					goto Label8
					// line 154: if optarg is None:
					πF.SetLineno(154)
				Label7:
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µargs); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp008).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label9
					}
					goto Label10
					// line 155: if not args:
					πF.SetLineno(155)
				Label9:
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("option --%s requires argument").ToObject(), µopt); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					πTemp001[1] = µopt
					if πTemp002, πE = πg.ResolveGlobal(πF, ßGetoptError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 156: raise GetoptError('option --%s requires argument' % opt, opt)
					πF.SetLineno(156)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label10
				Label10:
					// line 157: optarg, args = args[0], args[1:]
					πF.SetLineno(157)
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µargs, πTemp003); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µargs, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
						continue
					}
					µoptarg = πTemp003
					µargs = πTemp004
					goto Label8
				Label8:
					goto Label6
					// line 158: elif optarg is not None:
					πF.SetLineno(158)
				Label5:
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("option --%s must not have an argument").ToObject(), µopt); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					πTemp001[1] = µopt
					if πTemp002, πE = πg.ResolveGlobal(πF, ßGetoptError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 159: raise GetoptError('option --%s must not have an argument' % opt, opt)
					πF.SetLineno(159)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label6
				Label6:
					// line 160: opts.append(('--' + opt, optarg or ''))
					πF.SetLineno(160)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πg.NewStr("--").ToObject(), µopt); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µoptarg, "optarg"); πE != nil {
						continue
					}
					πTemp004 = µoptarg
					if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label11
					}
					πTemp004 = ß.ToObject()
				Label11:
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µopts, "opts"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µopts, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 161: return opts, args
					πF.SetLineno(161)
					if πE = πg.CheckLocal(πF, µopts, "opts"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µopts, µargs).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßdo_longs.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 166: def long_has_args(opt, longopts):
			πF.SetLineno(166)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "opt", Def: nil}
			πTemp006[1] = πg.Param{Name: "longopts", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("long_has_args", "build/src/__python__/getopt.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µopt *πg.Object = πArgs[0]; _ = µopt
				var µlongopts *πg.Object = πArgs[1]; _ = µlongopts
				var µpossibilities *πg.Object = πg.UnboundLocal; _ = µpossibilities
				var µunique_match *πg.Object = πg.UnboundLocal; _ = µunique_match
				var µhas_arg *πg.Object = πg.UnboundLocal; _ = µhas_arg
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []πg.Param
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 bool
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
					// line 167: possibilities = [o for o in longopts if o.startswith(opt)]
					πF.SetLineno(167)
					πTemp003 = make([]πg.Param, 0)
					πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/getopt.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µo *πg.Object = πg.UnboundLocal; _ = µo
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
								if πE = πg.CheckLocal(πF, µlongopts, "longopts"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µlongopts); πE != nil {
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
									µo = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								πTemp005 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
									continue
								}
								πTemp005[0] = µopt
								if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µo, ßstartswith, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								if πTemp003, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label4
								}
								goto Label5
								// line 167: possibilities = [o for o in longopts if o.startswith(opt)]
								πF.SetLineno(167)
							Label4:
								// line 167: possibilities = [o for o in longopts if o.startswith(opt)]
								πF.SetLineno(167)
								if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
									continue
								}
								πF.PushCheckpoint(6)
								return µo, nil
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
					if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
						continue
					}
					µpossibilities = πTemp001
					if πE = πg.CheckLocal(πF, µpossibilities, "possibilities"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µpossibilities); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 168: if not possibilities:
					πF.SetLineno(168)
				Label1:
					πTemp006 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("option --%s not recognized").ToObject(), µopt); πE != nil {
						continue
					}
					πTemp006[0] = πTemp001
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					πTemp006[1] = µopt
					if πTemp001, πE = πg.ResolveGlobal(πF, ßGetoptError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					// line 169: raise GetoptError('option --%s not recognized' % opt, opt)
					πF.SetLineno(169)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpossibilities, "possibilities"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, µpossibilities, µopt); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label3
					}
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, µopt, πg.NewStr("=").ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpossibilities, "possibilities"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, µpossibilities, πTemp004); πE != nil {
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
					// line 171: if opt in possibilities:
					πF.SetLineno(171)
				Label3:
					// line 172: return False, opt
					πF.SetLineno(172)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp004, µopt).ToObject()
					πR = πTemp001
					continue
					goto Label5
					// line 173: elif opt + '=' in possibilities:
					πF.SetLineno(173)
				Label4:
					// line 174: return True, opt
					πF.SetLineno(174)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp004, µopt).ToObject()
					πR = πTemp001
					continue
					goto Label5
				Label5:
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpossibilities, "possibilities"); πE != nil {
						continue
					}
					πTemp006[0] = µpossibilities
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp001, πE = πg.GT(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label6
					}
					goto Label7
					// line 176: if len(possibilities) > 1:
					πF.SetLineno(176)
				Label6:
					πTemp006 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("option --%s not a unique prefix").ToObject(), µopt); πE != nil {
						continue
					}
					πTemp006[0] = πTemp001
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					πTemp006[1] = µopt
					if πTemp001, πE = πg.ResolveGlobal(πF, ßGetoptError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					// line 179: raise GetoptError('option --%s not a unique prefix' % opt, opt)
					πF.SetLineno(179)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					goto Label7
				Label7:
					// line 180: assert len(possibilities) == 1
					πF.SetLineno(180)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpossibilities, "possibilities"); πE != nil {
						continue
					}
					πTemp006[0] = µpossibilities
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp001, πE = πg.Eq(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 181: unique_match = possibilities[0]
					πF.SetLineno(181)
					πTemp001 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µpossibilities, "possibilities"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µpossibilities, πTemp001); πE != nil {
						continue
					}
					µunique_match = πTemp004
					// line 182: has_arg = unique_match.endswith('=')
					πF.SetLineno(182)
					πTemp006 = πF.MakeArgs(1)
					πTemp006[0] = πg.NewStr("=").ToObject()
					if πE = πg.CheckLocal(πF, µunique_match, "unique_match"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µunique_match, ßendswith, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					µhas_arg = πTemp004
					if πE = πg.CheckLocal(πF, µhas_arg, "has_arg"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µhas_arg); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label8
					}
					goto Label9
					// line 183: if has_arg:
					πF.SetLineno(183)
				Label8:
					// line 184: unique_match = unique_match[:-1]
					πF.SetLineno(184)
					if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp004, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µunique_match, "unique_match"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µunique_match, πTemp001); πE != nil {
						continue
					}
					µunique_match = πTemp004
					goto Label9
				Label9:
					// line 185: return has_arg, unique_match
					πF.SetLineno(185)
					if πE = πg.CheckLocal(πF, µhas_arg, "has_arg"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µunique_match, "unique_match"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µhas_arg, µunique_match).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßlong_has_args.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 187: def do_shorts(opts, optstring, shortopts, args):
			πF.SetLineno(187)
			πTemp006 = make([]πg.Param, 4)
			πTemp006[0] = πg.Param{Name: "opts", Def: nil}
			πTemp006[1] = πg.Param{Name: "optstring", Def: nil}
			πTemp006[2] = πg.Param{Name: "shortopts", Def: nil}
			πTemp006[3] = πg.Param{Name: "args", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("do_shorts", "build/src/__python__/getopt.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µopts *πg.Object = πArgs[0]; _ = µopts
				var µoptstring *πg.Object = πArgs[1]; _ = µoptstring
				var µshortopts *πg.Object = πArgs[2]; _ = µshortopts
				var µargs *πg.Object = πArgs[3]; _ = µargs
				var µopt *πg.Object = πg.UnboundLocal; _ = µopt
				var µoptarg *πg.Object = πg.UnboundLocal; _ = µoptarg
				var πTemp001 bool
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 []*πg.Object
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
					// line 188: while optstring != '':
					πF.SetLineno(188)
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
					if πE = πg.CheckLocal(πF, µoptstring, "optstring"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.NE(πF, µoptstring, ß.ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πE != nil || !πTemp002 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 189: opt, optstring = optstring[0], optstring[1:]
					πF.SetLineno(189)
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µoptstring, "optstring"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µoptstring, πTemp004); πE != nil {
						continue
					}
					if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µoptstring, "optstring"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µoptstring, πTemp004); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
						continue
					}
					µopt = πTemp004
					µoptstring = πTemp005
					πTemp007 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					πTemp007[0] = µopt
					if πE = πg.CheckLocal(πF, µshortopts, "shortopts"); πE != nil {
						continue
					}
					πTemp007[1] = µshortopts
					if πTemp003, πE = πg.ResolveGlobal(πF, ßshort_has_arg); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					goto Label5
					// line 190: if short_has_arg(opt, shortopts):
					πF.SetLineno(190)
				Label4:
					if πE = πg.CheckLocal(πF, µoptstring, "optstring"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µoptstring, ß.ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label7
					}
					goto Label8
					// line 191: if optstring == '':
					πF.SetLineno(191)
				Label7:
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µargs); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label9
					}
					goto Label10
					// line 192: if not args:
					πF.SetLineno(192)
				Label9:
					πTemp007 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("option -%s requires argument").ToObject(), µopt); πE != nil {
						continue
					}
					πTemp007[0] = πTemp003
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					πTemp007[1] = µopt
					if πTemp003, πE = πg.ResolveGlobal(πF, ßGetoptError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 193: raise GetoptError('option -%s requires argument' % opt,
					πF.SetLineno(193)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					goto Label10
				Label10:
					// line 195: optstring, args = args[0], args[1:]
					πF.SetLineno(195)
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µargs, πTemp004); πE != nil {
						continue
					}
					if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µargs, πTemp004); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
						continue
					}
					µoptstring = πTemp004
					µargs = πTemp005
					goto Label8
				Label8:
					// line 196: optarg, optstring = optstring, ''
					πF.SetLineno(196)
					if πE = πg.CheckLocal(πF, µoptstring, "optstring"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µoptstring, ß.ToObject()).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
						continue
					}
					µoptarg = πTemp004
					µoptstring = πTemp005
					goto Label6
				Label5:
					// line 198: optarg = ''
					πF.SetLineno(198)
					µoptarg = ß.ToObject()
					goto Label6
				Label6:
					// line 199: opts.append(('-' + opt, optarg))
					πF.SetLineno(199)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, πg.NewStr("-").ToObject(), µopt); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µoptarg, "optarg"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp004, µoptarg).ToObject()
					πTemp007[0] = πTemp003
					if πE = πg.CheckLocal(πF, µopts, "opts"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µopts, ßappend, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 200: return opts, args
					πF.SetLineno(200)
					if πE = πg.CheckLocal(πF, µopts, "opts"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µopts, µargs).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßdo_shorts.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 202: def short_has_arg(opt, shortopts):
			πF.SetLineno(202)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "opt", Def: nil}
			πTemp006[1] = πg.Param{Name: "shortopts", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("short_has_arg", "build/src/__python__/getopt.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µopt *πg.Object = πArgs[0]; _ = µopt
				var µshortopts *πg.Object = πArgs[1]; _ = µshortopts
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
					πTemp002 = πF.MakeArgs(1)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µshortopts, "shortopts"); πE != nil {
						continue
					}
					πTemp003[0] = µshortopts
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
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp005 = µi
					if πE = πg.CheckLocal(πF, µshortopts, "shortopts"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µshortopts, πTemp005); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, µopt, πTemp008); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label4
					}
					if πTemp004, πE = πg.NE(πF, πTemp008, πg.NewStr(":").ToObject()); πE != nil {
						continue
					}
				Label4:
					if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label5
					}
					goto Label6
					// line 204: if opt == shortopts[i] != ':':
					πF.SetLineno(204)
				Label5:
					// line 205: return shortopts.startswith(':', i+1)
					πF.SetLineno(205)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewStr(":").ToObject()
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp002[1] = πTemp004
					if πE = πg.CheckLocal(πF, µshortopts, "shortopts"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µshortopts, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πR = πTemp005
					continue
					goto Label6
				Label6:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("option -%s not recognized").ToObject(), µopt); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					πTemp002[1] = µopt
					if πTemp001, πE = πg.ResolveGlobal(πF, ßGetoptError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 206: raise GetoptError('option -%s not recognized' % opt, opt)
					πF.SetLineno(206)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßshort_has_arg.ToObject(), πTemp009); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp010, πE = πg.Eq(πF, πTemp011, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp012, πE = πg.IsTrue(πF, πTemp010); πE != nil {
				continue
			}
			if πTemp012 {
				goto Label1
			}
			goto Label2
			// line 208: if __name__ == '__main__':
			πF.SetLineno(208)
		Label1:
			// line 209: import sys
			πF.SetLineno(209)
			if πTemp001, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp010 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 210: print getopt(sys.argv[1:], "a:b", ["alpha=", "beta"])
			πF.SetLineno(210)
			πTemp001 = make([]*πg.Object, 1)
			πTemp013 = πF.MakeArgs(3)
			if πTemp010, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
				continue
			}
			if πTemp014, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp015, πE = πg.GetAttr(πF, πTemp014, ßargv, nil); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetItem(πF, πTemp015, πTemp010); πE != nil {
				continue
			}
			πTemp013[0] = πTemp011
			πTemp013[1] = πg.NewStr("a:b").ToObject()
			πTemp016 = make([]*πg.Object, 2)
			πTemp016[0] = πg.NewStr("alpha=").ToObject()
			πTemp016[1] = ßbeta.ToObject()
			πTemp010 = πg.NewList(πTemp016...).ToObject()
			πTemp013[2] = πTemp010
			if πTemp010, πE = πg.ResolveGlobal(πF, ßgetopt); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp010.Call(πF, πTemp013, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp013)
			πTemp001[0] = πTemp011
			if πE = πg.Print(πF, πTemp001, true); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("getopt", Code)
}
