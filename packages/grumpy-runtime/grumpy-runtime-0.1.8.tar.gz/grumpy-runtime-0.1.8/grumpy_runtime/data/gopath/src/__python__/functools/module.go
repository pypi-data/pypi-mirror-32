package functools
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/functools.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßTypeError := πg.InternStr("TypeError")
		ßValueError := πg.InternStr("ValueError")
		ßWRAPPER_ASSIGNMENTS := πg.InternStr("WRAPPER_ASSIGNMENTS")
		ßWRAPPER_UPDATES := πg.InternStr("WRAPPER_UPDATES")
		ß__dict__ := πg.InternStr("__dict__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__eq__ := πg.InternStr("__eq__")
		ß__ge__ := πg.InternStr("__ge__")
		ß__gt__ := πg.InternStr("__gt__")
		ß__hash__ := πg.InternStr("__hash__")
		ß__init__ := πg.InternStr("__init__")
		ß__le__ := πg.InternStr("__le__")
		ß__lt__ := πg.InternStr("__lt__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__ne__ := πg.InternStr("__ne__")
		ß__slots__ := πg.InternStr("__slots__")
		ß_functools := πg.InternStr("_functools")
		ßcmp_to_key := πg.InternStr("cmp_to_key")
		ßdir := πg.InternStr("dir")
		ßgetattr := πg.InternStr("getattr")
		ßint := πg.InternStr("int")
		ßmax := πg.InternStr("max")
		ßobj := πg.InternStr("obj")
		ßobject := πg.InternStr("object")
		ßpartial := πg.InternStr("partial")
		ßreduce := πg.InternStr("reduce")
		ßset := πg.InternStr("set")
		ßsetattr := πg.InternStr("setattr")
		ßtotal_ordering := πg.InternStr("total_ordering")
		ßupdate := πg.InternStr("update")
		ßupdate_wrapper := πg.InternStr("update_wrapper")
		ßwraps := πg.InternStr("wraps")
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """functools.py - Tools for working with functions and callable objects
			πF.SetLineno(1)
			// line 11: import _functools
			πF.SetLineno(11)
			if πTemp002, πE = πg.ImportModule(πF, "_functools"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_functools.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 12: partial = _functools.partial
			πF.SetLineno(12)
			if πTemp001, πE = πg.ResolveGlobal(πF, ß_functools); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpartial, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßpartial.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 13: reduce = _functools.reduce
			πF.SetLineno(13)
			if πTemp001, πE = πg.ResolveGlobal(πF, ß_functools); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßreduce, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßreduce.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 15: def setattr(d, k, v):
			πF.SetLineno(15)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "d", Def: nil}
			πTemp004[1] = πg.Param{Name: "k", Def: nil}
			πTemp004[2] = πg.Param{Name: "v", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("setattr", "build/src/__python__/functools.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µd *πg.Object = πArgs[0]; _ = µd
				var µk *πg.Object = πArgs[1]; _ = µk
				var µv *πg.Object = πArgs[2]; _ = µv
				var πTemp001 *πg.Object
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
					// line 16: d.__dict__[k] = v
					πF.SetLineno(16)
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µv); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µd, ß__dict__, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
						continue
					}
					πTemp003 = µk
					if πE = πg.SetItem(πF, πTemp002, πTemp003, πTemp001); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßsetattr.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 21: WRAPPER_ASSIGNMENTS = ('__module__', '__name__') #, '__doc__'
			πF.SetLineno(21)
			πTemp003 = πg.NewTuple2(ß__module__.ToObject(), ß__name__.ToObject()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßWRAPPER_ASSIGNMENTS.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 22: WRAPPER_UPDATES = ('__dict__',)
			πF.SetLineno(22)
			πTemp003 = πg.NewTuple1(ß__dict__.ToObject()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßWRAPPER_UPDATES.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 23: def update_wrapper(wrapper,
			πF.SetLineno(23)
			πTemp004 = make([]πg.Param, 4)
			πTemp004[0] = πg.Param{Name: "wrapper", Def: nil}
			πTemp004[1] = πg.Param{Name: "wrapped", Def: nil}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßWRAPPER_ASSIGNMENTS); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "assigned", Def: πTemp005}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßWRAPPER_UPDATES); πE != nil {
				continue
			}
			πTemp004[3] = πg.Param{Name: "updated", Def: πTemp005}
			πTemp003 = πg.NewFunction(πg.NewCode("update_wrapper", "build/src/__python__/functools.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µwrapper *πg.Object = πArgs[0]; _ = µwrapper
				var µwrapped *πg.Object = πArgs[1]; _ = µwrapped
				var µassigned *πg.Object = πArgs[2]; _ = µassigned
				var µupdated *πg.Object = πArgs[3]; _ = µupdated
				var µattr *πg.Object = πg.UnboundLocal; _ = µattr
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
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Dict
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					case 4: goto Label4
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 27: """Update a wrapper function to look like the wrapped function
					πF.SetLineno(27)
					if πE = πg.CheckLocal(πF, µassigned, "assigned"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µassigned); πE != nil {
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
						µattr = πTemp004
					}
					if πE != nil || !πTemp003 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 39: setattr(wrapper, attr, getattr(wrapped, attr))
					πF.SetLineno(39)
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µwrapper, "wrapper"); πE != nil {
						continue
					}
					πTemp005[0] = µwrapper
					if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
						continue
					}
					πTemp005[1] = µattr
					πTemp006 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µwrapped, "wrapped"); πE != nil {
						continue
					}
					πTemp006[0] = µwrapped
					if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
						continue
					}
					πTemp006[1] = µattr
					if πTemp004, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp005[2] = πTemp007
					if πTemp004, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					if πE = πg.CheckLocal(πF, µupdated, "updated"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µupdated); πE != nil {
						continue
					}
					πF.PushCheckpoint(5)
					πTemp002 = false
				Label4:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
						πF.PopCheckpoint()
						goto Label6
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
						µattr = πTemp004
					}
					if πE != nil || !πTemp003 {
						continue
					}
					πF.PushCheckpoint(4)            
					// line 41: getattr(wrapper, attr).update(getattr(wrapped, attr, {}))
					πF.SetLineno(41)
					πTemp005 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µwrapped, "wrapped"); πE != nil {
						continue
					}
					πTemp006[0] = µwrapped
					if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
						continue
					}
					πTemp006[1] = µattr
					πTemp008 = πg.NewDict()
					πTemp004 = πTemp008.ToObject()
					πTemp006[2] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp005[0] = πTemp007
					πTemp006 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µwrapper, "wrapper"); πE != nil {
						continue
					}
					πTemp006[0] = µwrapper
					if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
						continue
					}
					πTemp006[1] = µattr
					if πTemp004, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp004, πE = πg.GetAttr(πF, πTemp007, ßupdate, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					continue
				Label5:
					if πE != nil || πR != nil {
						continue
					}
				Label6:
					// line 43: return wrapper
					πF.SetLineno(43)
					if πE = πg.CheckLocal(πF, µwrapper, "wrapper"); πE != nil {
						continue
					}
					πR = µwrapper
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßupdate_wrapper.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 45: def wraps(wrapped,
			πF.SetLineno(45)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "wrapped", Def: nil}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßWRAPPER_ASSIGNMENTS); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "assigned", Def: πTemp006}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßWRAPPER_UPDATES); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "updated", Def: πTemp006}
			πTemp005 = πg.NewFunction(πg.NewCode("wraps", "build/src/__python__/functools.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µwrapped *πg.Object = πArgs[0]; _ = µwrapped
				var µassigned *πg.Object = πArgs[1]; _ = µassigned
				var µupdated *πg.Object = πArgs[2]; _ = µupdated
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 πg.KWArgs
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
					// line 48: """Decorator factory to apply update_wrapper() to a wrapper function
					πF.SetLineno(48)
					// line 56: return partial(update_wrapper, wrapped=wrapped,
					πF.SetLineno(56)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßupdate_wrapper); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µwrapped, "wrapped"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µassigned, "assigned"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µupdated, "updated"); πE != nil {
						continue
					}
					πTemp003 = πg.KWArgs{
						{"wrapped", µwrapped},
						{"assigned", µassigned},
						{"updated", µupdated},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßpartial); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, πTemp003); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßwraps.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 59: def total_ordering(cls):
			πF.SetLineno(59)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "cls", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("total_ordering", "build/src/__python__/functools.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcls *πg.Object = πArgs[0]; _ = µcls
				var µconvert *πg.Object = πg.UnboundLocal; _ = µconvert
				var µroots *πg.Object = πg.UnboundLocal; _ = µroots
				var µroot *πg.Object = πg.UnboundLocal; _ = µroot
				var µopname *πg.Object = πg.UnboundLocal; _ = µopname
				var µopfunc *πg.Object = πg.UnboundLocal; _ = µopfunc
				var πTemp001 *πg.Dict
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []πg.Param
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πTemp010 bool
				_ = πTemp010
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 3: goto Label3
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					// line 60: """Class decorator that fills in missing ordering methods"""
					πF.SetLineno(60)
					// line 61: convert = {
					πF.SetLineno(61)
					πTemp001 = πg.NewDict()
					πTemp002 = make([]*πg.Object, 3)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "other", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/functools.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
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
							// line 62: '__lt__': [('__gt__', lambda self, other: not (self < other or self == other)),
							πF.SetLineno(62)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.LT(πF, µself, µother); πE != nil {
								continue
							}
							πTemp002 = πTemp004
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Eq(πF, µself, µother); πE != nil {
								continue
							}
							πTemp002 = πTemp004
						Label1:
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
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
					πTemp003 = πg.NewTuple2(ß__gt__.ToObject(), πTemp004).ToObject()
					πTemp002[0] = πTemp003
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "other", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/functools.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
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
							// line 63: ('__le__', lambda self, other: self < other or self == other),
							πF.SetLineno(63)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LT(πF, µself, µother); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, µself, µother); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label1:
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
					πTemp003 = πg.NewTuple2(ß__le__.ToObject(), πTemp004).ToObject()
					πTemp002[1] = πTemp003
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "other", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/functools.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 64: ('__ge__', lambda self, other: not self < other)],
							πF.SetLineno(64)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, µself, µother); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
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
					πTemp003 = πg.NewTuple2(ß__ge__.ToObject(), πTemp004).ToObject()
					πTemp002[2] = πTemp003
					πTemp003 = πg.NewList(πTemp002...).ToObject()
					if πE = πTemp001.SetItem(πF, ß__lt__.ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp002 = make([]*πg.Object, 3)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "other", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/functools.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 65: '__le__': [('__ge__', lambda self, other: not self <= other or self == other),
							πF.SetLineno(65)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.LE(πF, µself, µother); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp005).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, µself, µother); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label1:
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
					πTemp003 = πg.NewTuple2(ß__ge__.ToObject(), πTemp004).ToObject()
					πTemp002[0] = πTemp003
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "other", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/functools.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 66: ('__lt__', lambda self, other: self <= other and not self == other),
							πF.SetLineno(66)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LE(πF, µself, µother); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Eq(πF, µself, µother); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp005).ToObject()
							πTemp001 = πTemp003
						Label1:
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
					πTemp003 = πg.NewTuple2(ß__lt__.ToObject(), πTemp004).ToObject()
					πTemp002[1] = πTemp003
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "other", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/functools.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 67: ('__gt__', lambda self, other: not self <= other)],
							πF.SetLineno(67)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LE(πF, µself, µother); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
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
					πTemp003 = πg.NewTuple2(ß__gt__.ToObject(), πTemp004).ToObject()
					πTemp002[2] = πTemp003
					πTemp003 = πg.NewList(πTemp002...).ToObject()
					if πE = πTemp001.SetItem(πF, ß__le__.ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp002 = make([]*πg.Object, 3)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "other", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/functools.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
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
							// line 68: '__gt__': [('__lt__', lambda self, other: not (self > other or self == other)),
							πF.SetLineno(68)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GT(πF, µself, µother); πE != nil {
								continue
							}
							πTemp002 = πTemp004
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Eq(πF, µself, µother); πE != nil {
								continue
							}
							πTemp002 = πTemp004
						Label1:
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
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
					πTemp003 = πg.NewTuple2(ß__lt__.ToObject(), πTemp004).ToObject()
					πTemp002[0] = πTemp003
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "other", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/functools.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
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
							// line 69: ('__ge__', lambda self, other: self > other or self == other),
							πF.SetLineno(69)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, µself, µother); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, µself, µother); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label1:
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
					πTemp003 = πg.NewTuple2(ß__ge__.ToObject(), πTemp004).ToObject()
					πTemp002[1] = πTemp003
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "other", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/functools.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 70: ('__le__', lambda self, other: not self > other)],
							πF.SetLineno(70)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GT(πF, µself, µother); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
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
					πTemp003 = πg.NewTuple2(ß__le__.ToObject(), πTemp004).ToObject()
					πTemp002[2] = πTemp003
					πTemp003 = πg.NewList(πTemp002...).ToObject()
					if πE = πTemp001.SetItem(πF, ß__gt__.ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp002 = make([]*πg.Object, 3)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "other", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/functools.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 71: '__ge__': [('__le__', lambda self, other: (not self >= other) or self == other),
							πF.SetLineno(71)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GE(πF, µself, µother); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp005).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, µself, µother); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label1:
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
					πTemp003 = πg.NewTuple2(ß__le__.ToObject(), πTemp004).ToObject()
					πTemp002[0] = πTemp003
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "other", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/functools.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 72: ('__gt__', lambda self, other: self >= other and not self == other),
							πF.SetLineno(72)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GE(πF, µself, µother); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Eq(πF, µself, µother); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp005).ToObject()
							πTemp001 = πTemp003
						Label1:
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
					πTemp003 = πg.NewTuple2(ß__gt__.ToObject(), πTemp004).ToObject()
					πTemp002[1] = πTemp003
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "other", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/functools.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 73: ('__lt__', lambda self, other: not self >= other)]
							πF.SetLineno(73)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GE(πF, µself, µother); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
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
					πTemp003 = πg.NewTuple2(ß__lt__.ToObject(), πTemp004).ToObject()
					πTemp002[2] = πTemp003
					πTemp003 = πg.NewList(πTemp002...).ToObject()
					if πE = πTemp001.SetItem(πF, ß__ge__.ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp003 = πTemp001.ToObject()
					µconvert = πTemp003
					// line 75: roots = set(dir(cls)) & set(convert)
					πF.SetLineno(75)
					πTemp002 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					πTemp006[0] = µcls
					if πTemp004, πE = πg.ResolveGlobal(πF, ßdir); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp002[0] = πTemp007
					if πTemp004, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µconvert, "convert"); πE != nil {
						continue
					}
					πTemp002[0] = µconvert
					if πTemp004, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.And(πF, πTemp007, πTemp008); πE != nil {
						continue
					}
					µroots = πTemp003
					if πE = πg.CheckLocal(πF, µroots, "roots"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, µroots); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp009).ToObject()
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label1
					}
					goto Label2
					// line 76: if not roots:
					πF.SetLineno(76)
				Label1:
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("must define at least one ordering operation: < > <= >=").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 77: raise ValueError('must define at least one ordering operation: < > <= >=')
					πF.SetLineno(77)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					goto Label2
				Label2:
					// line 78: root = max(roots)       # prefer __lt__ to __le__ to __gt__ to __ge__
					πF.SetLineno(78)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µroots, "roots"); πE != nil {
						continue
					}
					πTemp002[0] = µroots
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µroot = πTemp004
					if πE = πg.CheckLocal(πF, µroot, "root"); πE != nil {
						continue
					}
					πTemp004 = µroot
					if πE = πg.CheckLocal(πF, µconvert, "convert"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µconvert, πTemp004); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Iter(πF, πTemp007); πE != nil {
						continue
					}
					πF.PushCheckpoint(4)
					πTemp009 = false
				Label3:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp009 {
						πF.PopCheckpoint()
						goto Label5
					}
					if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}}}, πTemp004); πE != nil {
							continue
						}
						µopname = πTemp007
						µopfunc = πTemp008
					}
					if πE != nil || !πTemp010 {
						continue
					}
					πF.PushCheckpoint(3)            
					if πE = πg.CheckLocal(πF, µopname, "opname"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µroots, "roots"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Contains(πF, µroots, µopname); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(!πTemp010).ToObject()
					if πTemp010, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp010 {
						goto Label6
					}
					goto Label7
					// line 80: if opname not in roots:
					πF.SetLineno(80)
				Label6:
					// line 81: opfunc.__name__ = opname
					πF.SetLineno(81)
					if πE = πg.CheckLocal(πF, µopname, "opname"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µopname); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µopfunc, "opfunc"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µopfunc, ß__name__, πTemp004); πE != nil {
						continue
					}
					// line 82: opfunc.__doc__ = getattr(int, opname).__doc__
					πF.SetLineno(82)
					πTemp002 = πF.MakeArgs(2)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πE = πg.CheckLocal(πF, µopname, "opname"); πE != nil {
						continue
					}
					πTemp002[1] = µopname
					if πTemp004, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp004, πE = πg.GetAttr(πF, πTemp007, ß__doc__, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µopfunc, "opfunc"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µopfunc, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 83: setattr(cls, opname, opfunc)
					πF.SetLineno(83)
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					πTemp002[0] = µcls
					if πE = πg.CheckLocal(πF, µopname, "opname"); πE != nil {
						continue
					}
					πTemp002[1] = µopname
					if πE = πg.CheckLocal(πF, µopfunc, "opfunc"); πE != nil {
						continue
					}
					πTemp002[2] = µopfunc
					if πTemp004, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label7
				Label7:
					continue
				Label4:
					if πE != nil || πR != nil {
						continue
					}
				Label5:
					// line 84: return cls
					πF.SetLineno(84)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					πR = µcls
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßtotal_ordering.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 86: def cmp_to_key(mycmp):
			πF.SetLineno(86)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "mycmp", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("cmp_to_key", "build/src/__python__/functools.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmycmp *πg.Object = πArgs[0]; _ = µmycmp
				var µK *πg.Object = πg.UnboundLocal; _ = µK
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 87: """Convert a cmp= function into a key= function"""
					πF.SetLineno(87)
					// line 88: class K(object):
					πF.SetLineno(88)
					πTemp003 = make([]*πg.Object, 1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					πTemp001 = πg.NewDict()
					if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
						continue
					}
					_, πE = πg.NewCode("K", "build/src/__python__/functools.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
						πClass := πTemp001
						_ = πClass
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 89: __slots__ = ['obj']
							πF.SetLineno(89)
							πTemp001 = make([]*πg.Object, 1)
							πTemp001[0] = ßobj.ToObject()
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πClass.SetItem(πF, ß__slots__.ToObject(), πTemp002); πE != nil {
								continue
							}
							// line 90: def __init__(self, obj, *args):
							πF.SetLineno(90)
							πTemp003 = make([]πg.Param, 2)
							πTemp003[0] = πg.Param{Name: "self", Def: nil}
							πTemp003[1] = πg.Param{Name: "obj", Def: nil}
							πTemp002 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/functools.py", πTemp003, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
								var µobj *πg.Object = πArgs[1]; _ = µobj
								var µargs *πg.Object = πArgs[2]; _ = µargs
								var πTemp001 *πg.Object
								_ = πTemp001
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 91: self.obj = obj
									πF.SetLineno(91)
									if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
										continue
									}
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µobj); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, µself, ßobj, πTemp001); πE != nil {
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
							// line 92: def __lt__(self, other):
							πF.SetLineno(92)
							πTemp003 = make([]πg.Param, 2)
							πTemp003[0] = πg.Param{Name: "self", Def: nil}
							πTemp003[1] = πg.Param{Name: "other", Def: nil}
							πTemp004 = πg.NewFunction(πg.NewCode("__lt__", "build/src/__python__/functools.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
								var µother *πg.Object = πArgs[1]; _ = µother
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
									// line 93: return mycmp(self.obj, other.obj) < 0
									πF.SetLineno(93)
									πTemp002 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µself, ßobj, nil); πE != nil {
										continue
									}
									πTemp002[0] = πTemp003
									if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µother, ßobj, nil); πE != nil {
										continue
									}
									πTemp002[1] = πTemp003
									if πE = πg.CheckLocal(πF, µmycmp, "mycmp"); πE != nil {
										continue
									}
									if πTemp003, πE = µmycmp.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									if πTemp001, πE = πg.LT(πF, πTemp003, πg.NewInt(0).ToObject()); πE != nil {
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
							if πE = πClass.SetItem(πF, ß__lt__.ToObject(), πTemp004); πE != nil {
								continue
							}
							// line 94: def __gt__(self, other):
							πF.SetLineno(94)
							πTemp003 = make([]πg.Param, 2)
							πTemp003[0] = πg.Param{Name: "self", Def: nil}
							πTemp003[1] = πg.Param{Name: "other", Def: nil}
							πTemp005 = πg.NewFunction(πg.NewCode("__gt__", "build/src/__python__/functools.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
								var µother *πg.Object = πArgs[1]; _ = µother
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
									// line 95: return mycmp(self.obj, other.obj) > 0
									πF.SetLineno(95)
									πTemp002 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µself, ßobj, nil); πE != nil {
										continue
									}
									πTemp002[0] = πTemp003
									if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µother, ßobj, nil); πE != nil {
										continue
									}
									πTemp002[1] = πTemp003
									if πE = πg.CheckLocal(πF, µmycmp, "mycmp"); πE != nil {
										continue
									}
									if πTemp003, πE = µmycmp.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									if πTemp001, πE = πg.GT(πF, πTemp003, πg.NewInt(0).ToObject()); πE != nil {
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
							if πE = πClass.SetItem(πF, ß__gt__.ToObject(), πTemp005); πE != nil {
								continue
							}
							// line 96: def __eq__(self, other):
							πF.SetLineno(96)
							πTemp003 = make([]πg.Param, 2)
							πTemp003[0] = πg.Param{Name: "self", Def: nil}
							πTemp003[1] = πg.Param{Name: "other", Def: nil}
							πTemp006 = πg.NewFunction(πg.NewCode("__eq__", "build/src/__python__/functools.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
								var µother *πg.Object = πArgs[1]; _ = µother
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
									// line 97: return mycmp(self.obj, other.obj) == 0
									πF.SetLineno(97)
									πTemp002 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µself, ßobj, nil); πE != nil {
										continue
									}
									πTemp002[0] = πTemp003
									if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µother, ßobj, nil); πE != nil {
										continue
									}
									πTemp002[1] = πTemp003
									if πE = πg.CheckLocal(πF, µmycmp, "mycmp"); πE != nil {
										continue
									}
									if πTemp003, πE = µmycmp.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewInt(0).ToObject()); πE != nil {
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
							if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp006); πE != nil {
								continue
							}
							// line 98: def __le__(self, other):
							πF.SetLineno(98)
							πTemp003 = make([]πg.Param, 2)
							πTemp003[0] = πg.Param{Name: "self", Def: nil}
							πTemp003[1] = πg.Param{Name: "other", Def: nil}
							πTemp007 = πg.NewFunction(πg.NewCode("__le__", "build/src/__python__/functools.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
								var µother *πg.Object = πArgs[1]; _ = µother
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
									// line 99: return mycmp(self.obj, other.obj) <= 0
									πF.SetLineno(99)
									πTemp002 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µself, ßobj, nil); πE != nil {
										continue
									}
									πTemp002[0] = πTemp003
									if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µother, ßobj, nil); πE != nil {
										continue
									}
									πTemp002[1] = πTemp003
									if πE = πg.CheckLocal(πF, µmycmp, "mycmp"); πE != nil {
										continue
									}
									if πTemp003, πE = µmycmp.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									if πTemp001, πE = πg.LE(πF, πTemp003, πg.NewInt(0).ToObject()); πE != nil {
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
							if πE = πClass.SetItem(πF, ß__le__.ToObject(), πTemp007); πE != nil {
								continue
							}
							// line 100: def __ge__(self, other):
							πF.SetLineno(100)
							πTemp003 = make([]πg.Param, 2)
							πTemp003[0] = πg.Param{Name: "self", Def: nil}
							πTemp003[1] = πg.Param{Name: "other", Def: nil}
							πTemp008 = πg.NewFunction(πg.NewCode("__ge__", "build/src/__python__/functools.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
								var µother *πg.Object = πArgs[1]; _ = µother
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
									// line 101: return mycmp(self.obj, other.obj) >= 0
									πF.SetLineno(101)
									πTemp002 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µself, ßobj, nil); πE != nil {
										continue
									}
									πTemp002[0] = πTemp003
									if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µother, ßobj, nil); πE != nil {
										continue
									}
									πTemp002[1] = πTemp003
									if πE = πg.CheckLocal(πF, µmycmp, "mycmp"); πE != nil {
										continue
									}
									if πTemp003, πE = µmycmp.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									if πTemp001, πE = πg.GE(πF, πTemp003, πg.NewInt(0).ToObject()); πE != nil {
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
							if πE = πClass.SetItem(πF, ß__ge__.ToObject(), πTemp008); πE != nil {
								continue
							}
							// line 102: def __ne__(self, other):
							πF.SetLineno(102)
							πTemp003 = make([]πg.Param, 2)
							πTemp003[0] = πg.Param{Name: "self", Def: nil}
							πTemp003[1] = πg.Param{Name: "other", Def: nil}
							πTemp009 = πg.NewFunction(πg.NewCode("__ne__", "build/src/__python__/functools.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
								var µother *πg.Object = πArgs[1]; _ = µother
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
									// line 103: return mycmp(self.obj, other.obj) != 0
									πF.SetLineno(103)
									πTemp002 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µself, ßobj, nil); πE != nil {
										continue
									}
									πTemp002[0] = πTemp003
									if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µother, ßobj, nil); πE != nil {
										continue
									}
									πTemp002[1] = πTemp003
									if πE = πg.CheckLocal(πF, µmycmp, "mycmp"); πE != nil {
										continue
									}
									if πTemp003, πE = µmycmp.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									if πTemp001, πE = πg.NE(πF, πTemp003, πg.NewInt(0).ToObject()); πE != nil {
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
							if πE = πClass.SetItem(πF, ß__ne__.ToObject(), πTemp009); πE != nil {
								continue
							}
							// line 104: def __hash__(self):
							πF.SetLineno(104)
							πTemp003 = make([]πg.Param, 1)
							πTemp003[0] = πg.Param{Name: "self", Def: nil}
							πTemp010 = πg.NewFunction(πg.NewCode("__hash__", "build/src/__python__/functools.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
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
									πTemp001 = πF.MakeArgs(1)
									πTemp001[0] = πg.NewStr("hash not implemented").ToObject()
									if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									// line 105: raise TypeError('hash not implemented')
									πF.SetLineno(105)
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
							if πE = πClass.SetItem(πF, ß__hash__.ToObject(), πTemp010); πE != nil {
								continue
							}
						}
						return nil, nil
					}).Eval(πF, πF.Globals(), nil, nil)
					if πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
						continue
					}
					if πTemp004 == nil {
						πTemp004 = πg.TypeType.ToObject()
					}
					if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("K").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
						continue
					}
					µK = πTemp005
					// line 106: return K
					πF.SetLineno(106)
					if πE = πg.CheckLocal(πF, µK, "K"); πE != nil {
						continue
					}
					πR = µK
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßcmp_to_key.ToObject(), πTemp007); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("functools", Code)
}
