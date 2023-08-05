package contextlib
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/contextlib.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßDeprecationWarning := πg.InternStr("DeprecationWarning")
		ßGeneratorContextManager := πg.InternStr("GeneratorContextManager")
		ßNone := πg.InternStr("None")
		ßRuntimeError := πg.InternStr("RuntimeError")
		ßStopIteration := πg.InternStr("StopIteration")
		ß__all__ := πg.InternStr("__all__")
		ß__enter__ := πg.InternStr("__enter__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßappend := πg.InternStr("append")
		ßclose := πg.InternStr("close")
		ßclosing := πg.InternStr("closing")
		ßcontextmanager := πg.InternStr("contextmanager")
		ßexc_info := πg.InternStr("exc_info")
		ßfunctools := πg.InternStr("functools")
		ßgen := πg.InternStr("gen")
		ßnested := πg.InternStr("nested")
		ßnext := πg.InternStr("next")
		ßobject := πg.InternStr("object")
		ßpop := πg.InternStr("pop")
		ßsys := πg.InternStr("sys")
		ßthing := πg.InternStr("thing")
		ßwarn := πg.InternStr("warn")
		ßwarnings := πg.InternStr("warnings")
		ßwraps := πg.InternStr("wraps")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.Dict
		_ = πTemp004
		var πTemp005 *πg.Object
		_ = πTemp005
		var πTemp006 []πg.Param
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
			// line 1: """Utilities for with-statement contexts.  See PEP 343."""
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
			// line 5: import functools
			πF.SetLineno(5)
			if πTemp002, πE = πg.ImportModule(πF, "functools"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßfunctools.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 6: wraps = functools.wraps
			πF.SetLineno(6)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßfunctools); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßwraps, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßwraps.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 8: import warnings
			πF.SetLineno(8)
			if πTemp002, πE = πg.ImportModule(πF, "warnings"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßwarnings.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 9: warn = warnings.warn
			πF.SetLineno(9)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßwarnings); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßwarn, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßwarn.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 12: __all__ = ["contextmanager", "nested", "closing"]
			πF.SetLineno(12)
			πTemp002 = make([]*πg.Object, 3)
			πTemp002[0] = ßcontextmanager.ToObject()
			πTemp002[1] = ßnested.ToObject()
			πTemp002[2] = ßclosing.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 14: class GeneratorContextManager(object):
			πF.SetLineno(14)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("GeneratorContextManager", "build/src/__python__/contextlib.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 15: """Helper for @contextmanager decorator."""
					πF.SetLineno(15)
					// line 17: def __init__(self, gen):
					πF.SetLineno(17)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "gen", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/contextlib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µgen *πg.Object = πArgs[1]; _ = µgen
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 18: self.gen = gen
							πF.SetLineno(18)
							if πE = πg.CheckLocal(πF, µgen, "gen"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µgen); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßgen, πTemp001); πE != nil {
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
					// line 20: def __enter__(self):
					πF.SetLineno(20)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__enter__", "build/src/__python__/contextlib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.BaseException
						_ = πTemp003
						var πTemp004 *πg.Traceback
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 21: try:
							πF.SetLineno(21)
							πF.PushCheckpoint(2)
							// line 22: return self.gen.next()
							πF.SetLineno(22)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßgen, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnext, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp003, πTemp004 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
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
							// line 23: except StopIteration:
							πF.SetLineno(23)
						Label3:
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr("generator didn't yield").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 24: raise RuntimeError("generator didn't yield")
							πF.SetLineno(24)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__enter__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 26: def __exit__(self, t, value, tb):
					πF.SetLineno(26)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "t", Def: nil}
					πTemp002[2] = πg.Param{Name: "value", Def: nil}
					πTemp002[3] = πg.Param{Name: "tb", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__exit__", "build/src/__python__/contextlib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µt *πg.Object = πArgs[1]; _ = µt
						var µvalue *πg.Object = πArgs[2]; _ = µvalue
						var µtb *πg.Object = πArgs[3]; _ = µtb
						var µexc *πg.Object = πg.UnboundLocal; _ = µexc
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.BaseException
						_ = πTemp005
						var πTemp006 *πg.Traceback
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
							case 10: goto Label10
							case 5: goto Label5
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µt == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 27: if t is None:
							πF.SetLineno(27)
						Label1:
							// line 28: try:
							πF.SetLineno(28)
							πF.PushCheckpoint(5)
							// line 29: self.gen.next()
							πF.SetLineno(29)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßgen, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnext, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("generator didn't stop").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 33: raise RuntimeError("generator didn't stop")
							πF.SetLineno(33)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label4
						Label5:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 30: except StopIteration:
							πF.SetLineno(30)
						Label6:
							// line 31: return
							πF.SetLineno(31)
							πR = πg.None
							continue
							πF.RestoreExc(nil, nil)
							goto Label4
						Label4:
							goto Label3
						Label2:
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µvalue == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label7
							}
							goto Label8
							// line 35: if value is None:
							πF.SetLineno(35)
						Label7:
							// line 38: value = t()
							πF.SetLineno(38)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp001, πE = µt.Call(πF, nil, nil); πE != nil {
								continue
							}
							µvalue = πTemp001
							goto Label8
						Label8:
							// line 39: try:
							πF.SetLineno(39)
							πF.PushCheckpoint(10)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp004[0] = µvalue
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp001, πE = µt.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 42: raise t(value)
							πF.SetLineno(42)
							πE = πF.Raise(πTemp001, nil, nil)
							continue
							πF.PopCheckpoint()
							goto Label9
						Label10:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label11
							}
							goto Label12
							// line 43: except StopIteration, exc:
							πF.SetLineno(43)
						Label11:
							µexc = πTemp005.ToObject()
							// line 47: return exc is not value
							πF.SetLineno(47)
							if πE = πg.CheckLocal(πF, µexc, "exc"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µexc != µvalue).ToObject()
							πR = πTemp001
							continue
							πF.RestoreExc(nil, nil)
							goto Label9
							// line 48: except:
							πF.SetLineno(48)
						Label12:
							πTemp002 = πg.NewInt(1).ToObject()
							if πTemp008, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßexc_info, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp009.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp008, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp007 != µvalue).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label13
							}
							goto Label14
							// line 56: if sys.exc_info()[1] is not value:
							πF.SetLineno(56)
						Label13:
							// line 57: raise
							πF.SetLineno(57)
							πE = πF.Raise(nil, nil, nil)
							continue
							goto Label14
						Label14:
							πF.RestoreExc(nil, nil)
							goto Label9
						Label9:
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
					if πE = πClass.SetItem(πF, ß__exit__.ToObject(), πTemp004); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("GeneratorContextManager").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßGeneratorContextManager.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 60: def contextmanager(func):
			πF.SetLineno(60)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "func", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("contextmanager", "build/src/__python__/contextlib.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfunc *πg.Object = πArgs[0]; _ = µfunc
				var µhelper *πg.Object = πg.UnboundLocal; _ = µhelper
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
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
					// line 61: """@contextmanager decorator.
					πF.SetLineno(61)
					// line 89: def helper(*args, **kwds):
					πF.SetLineno(89)
					πTemp002 = make([]πg.Param, 0)
					πTemp001 = πg.NewFunction(πg.NewCode("helper", "build/src/__python__/contextlib.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µargs *πg.Object = πArgs[0]; _ = µargs
						var µkwds *πg.Object = πArgs[1]; _ = µkwds
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
							// line 90: return GeneratorContextManager(func(*args, **kwds))
							πF.SetLineno(90)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, µfunc, nil, µargs, nil, µkwds); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßGeneratorContextManager); πE != nil {
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
					µhelper = πTemp001
					// line 88: @wraps(func)
					πF.SetLineno(88)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µhelper, "helper"); πE != nil {
						continue
					}
					πTemp003[0] = µhelper
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp004[0] = µfunc
					if πTemp005, πE = πg.ResolveGlobal(πF, ßwraps); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp005, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µhelper = πTemp005
					// line 91: return helper
					πF.SetLineno(91)
					if πE = πg.CheckLocal(πF, µhelper, "helper"); πE != nil {
						continue
					}
					πR = µhelper
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßcontextmanager.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 95: def nested(*managers):
			πF.SetLineno(95)
			πTemp006 = make([]πg.Param, 0)
			πTemp003 = πg.NewFunction(πg.NewCode("nested", "build/src/__python__/contextlib.py", πTemp006, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmanagers *πg.Object = πArgs[0]; _ = µmanagers
				var µexits *πg.Object = πg.UnboundLocal; _ = µexits
				var µvars *πg.Object = πg.UnboundLocal; _ = µvars
				var µexc *πg.Object = πg.UnboundLocal; _ = µexc
				var µmgr *πg.Object = πg.UnboundLocal; _ = µmgr
				var µexit *πg.Object = πg.UnboundLocal; _ = µexit
				var µenter *πg.Object = πg.UnboundLocal; _ = µenter
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
				var πTemp008 *πg.BaseException
				_ = πTemp008
				var πTemp009 *πg.Traceback
				_ = πTemp009
				var πTemp010 *πg.BaseException
				_ = πTemp010
				var πTemp011 *πg.Traceback
				_ = πTemp011
				var πTemp012 *πg.Object
				_ = πTemp012
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 1: goto Label1
						case 2: goto Label2
						case 3: goto Label3
						case 4: goto Label4
						case 6: goto Label6
						case 8: goto Label8
						case 9: goto Label9
						case 12: goto Label12
						default: panic("unexpected function state")
						}
						// line 96: """Combine multiple context managers into a single nested context manager.
						πF.SetLineno(96)
						// line 109: warn("With-statements now directly support multiple context managers",
						πF.SetLineno(109)
						πTemp001 = πF.MakeArgs(3)
						πTemp001[0] = πg.NewStr("With-statements now directly support multiple context managers").ToObject()
						if πTemp002, πE = πg.ResolveGlobal(πF, ßDeprecationWarning); πE != nil {
							continue
						}
						πTemp001[1] = πTemp002
						πTemp001[2] = πg.NewInt(3).ToObject()
						if πTemp002, πE = πg.ResolveGlobal(πF, ßwarn); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						// line 111: exits = []
						πF.SetLineno(111)
						πTemp001 = make([]*πg.Object, 0)
						πTemp002 = πg.NewList(πTemp001...).ToObject()
						µexits = πTemp002
						// line 112: vars = []
						πF.SetLineno(112)
						πTemp001 = make([]*πg.Object, 0)
						πTemp002 = πg.NewList(πTemp001...).ToObject()
						µvars = πTemp002
						// line 113: exc = (None, None, None)
						πF.SetLineno(113)
						if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
							continue
						}
						if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
							continue
						}
						if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
							continue
						}
						πTemp002 = πg.NewTuple3(πTemp003, πTemp004, πTemp005).ToObject()
						µexc = πTemp002
						// line 114: try:
						πF.SetLineno(114)
						πF.PushCheckpoint(1)
						πF.PushCheckpoint(2)
						if πE = πg.CheckLocal(πF, µmanagers, "managers"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.Iter(πF, µmanagers); πE != nil {
							continue
						}
						πF.PushCheckpoint(4)
						πTemp006 = false
					Label3:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp006 {
							πF.PopCheckpoint()
							goto Label5
						}
						if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
							µmgr = πTemp003
						}
						if πE != nil || !πTemp007 {
							continue
						}
						πF.PushCheckpoint(3)            
						// line 116: exit = mgr.__exit__
						πF.SetLineno(116)
						if πE = πg.CheckLocal(πF, µmgr, "mgr"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, µmgr, ß__exit__, nil); πE != nil {
							continue
						}
						µexit = πTemp003
						// line 117: enter = mgr.__enter__
						πF.SetLineno(117)
						if πE = πg.CheckLocal(πF, µmgr, "mgr"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, µmgr, ß__enter__, nil); πE != nil {
							continue
						}
						µenter = πTemp003
						// line 118: vars.append(enter())
						πF.SetLineno(118)
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µenter, "enter"); πE != nil {
							continue
						}
						if πTemp003, πE = µenter.Call(πF, nil, nil); πE != nil {
							continue
						}
						πTemp001[0] = πTemp003
						if πE = πg.CheckLocal(πF, µvars, "vars"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, µvars, ßappend, nil); πE != nil {
							continue
						}
						if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						// line 119: exits.append(exit)
						πF.SetLineno(119)
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µexit, "exit"); πE != nil {
							continue
						}
						πTemp001[0] = µexit
						if πE = πg.CheckLocal(πF, µexits, "exits"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, µexits, ßappend, nil); πE != nil {
							continue
						}
						if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						continue
					Label4:
						if πE != nil || πR != nil {
							continue
						}
					Label5:
						// line 120: yield vars
						πF.SetLineno(120)
						if πE = πg.CheckLocal(πF, µvars, "vars"); πE != nil {
							continue
						}
						πF.PushCheckpoint(6)
						return µvars, nil
					Label6:
						πTemp002 = πSent
						πF.PopCheckpoint()
						πF.PopCheckpoint()
						goto Label1
					Label2:
						if πE == nil {
						  continue
						}
						πE = nil
						πTemp008, πTemp009 = πF.ExcInfo()
						goto Label7
						// line 121: except:
						πF.SetLineno(121)
					Label7:
						// line 122: exc = sys.exc_info()
						πF.SetLineno(122)
						if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßexc_info, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
							continue
						}
						µexc = πTemp002
						πF.RestoreExc(nil, nil)
						πF.PopCheckpoint()
						goto Label1
					Label1:
						πTemp008, πTemp009 = πF.RestoreExc(nil, nil)
						// line 124: while exits:
						πF.SetLineno(124)
						πF.PushCheckpoint(9)
						πTemp006 = false
					Label8:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp006 {
							πF.PopCheckpoint()
							goto Label10
						}
						if πE = πg.CheckLocal(πF, µexits, "exits"); πE != nil {
							continue
						}
						if πTemp007, πE = πg.IsTrue(πF, µexits); πE != nil {
							continue
						}
						if πE != nil || !πTemp007 {
							continue
						}
						πF.PushCheckpoint(8)            
						// line 125: exit = exits.pop()
						πF.SetLineno(125)
						if πE = πg.CheckLocal(πF, µexits, "exits"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µexits, ßpop, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
							continue
						}
						µexit = πTemp003
						// line 126: try:
						πF.SetLineno(126)
						πF.PushCheckpoint(12)
						if πE = πg.CheckLocal(πF, µexc, "exc"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µexit, "exit"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.Invoke(πF, µexit, nil, µexc, nil, nil); πE != nil {
							continue
						}
						if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
							continue
						}
						if πTemp007 {
							goto Label13
						}
						goto Label14
						// line 127: if exit(*exc):
						πF.SetLineno(127)
					Label13:
						// line 128: exc = (None, None, None)
						πF.SetLineno(128)
						if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
							continue
						}
						if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
							continue
						}
						if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
							continue
						}
						πTemp002 = πg.NewTuple3(πTemp003, πTemp004, πTemp005).ToObject()
						µexc = πTemp002
						goto Label14
					Label14:
						πF.PopCheckpoint()
						goto Label11
					Label12:
						if πE == nil {
						  continue
						}
						πE = nil
						πTemp010, πTemp011 = πF.ExcInfo()
						goto Label15
						// line 129: except:
						πF.SetLineno(129)
					Label15:
						// line 130: exc = sys.exc_info()
						πF.SetLineno(130)
						if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßexc_info, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
							continue
						}
						µexc = πTemp002
						πF.RestoreExc(nil, nil)
						goto Label11
					Label11:
						continue
					Label9:
						if πE != nil || πR != nil {
							continue
						}
					Label10:
						if πE = πg.CheckLocal(πF, µexc, "exc"); πE != nil {
							continue
						}
						if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
							continue
						}
						if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
							continue
						}
						if πTemp012, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
							continue
						}
						πTemp003 = πg.NewTuple3(πTemp004, πTemp005, πTemp012).ToObject()
						if πTemp002, πE = πg.NE(πF, µexc, πTemp003); πE != nil {
							continue
						}
						if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
							continue
						}
						if πTemp006 {
							goto Label16
						}
						goto Label17
						// line 131: if exc != (None, None, None):
						πF.SetLineno(131)
					Label16:
						πTemp002 = πg.NewInt(0).ToObject()
						if πE = πg.CheckLocal(πF, µexc, "exc"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetItem(πF, µexc, πTemp002); πE != nil {
							continue
						}
						πTemp002 = πg.NewInt(1).ToObject()
						if πE = πg.CheckLocal(πF, µexc, "exc"); πE != nil {
							continue
						}
						if πTemp004, πE = πg.GetItem(πF, µexc, πTemp002); πE != nil {
							continue
						}
						πTemp002 = πg.NewInt(2).ToObject()
						if πE = πg.CheckLocal(πF, µexc, "exc"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.GetItem(πF, µexc, πTemp002); πE != nil {
							continue
						}
						// line 135: raise exc[0], exc[1], exc[2]
						πF.SetLineno(135)
						πE = πF.Raise(πTemp003, πTemp004, πTemp005)
						continue
						goto Label17
					Label17:
						if πTemp008 != nil {
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
						}
						if πR != nil {
							continue
						}
					}
					return nil, πE
				}).ToObject(), nil
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßnested.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 94: @contextmanager
			πF.SetLineno(94)
			πTemp002 = πF.MakeArgs(1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßnested); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßcontextmanager); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßnested.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 138: class closing(object):
			πF.SetLineno(138)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp004 = πg.NewDict()
			if πTemp005, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp005); πE != nil {
				continue
			}
			_, πE = πg.NewCode("closing", "build/src/__python__/contextlib.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 139: """Context to automatically close something at the end of a block.
					πF.SetLineno(139)
					// line 155: def __init__(self, thing):
					πF.SetLineno(155)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "thing", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/contextlib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µthing *πg.Object = πArgs[1]; _ = µthing
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 156: self.thing = thing
							πF.SetLineno(156)
							if πE = πg.CheckLocal(πF, µthing, "thing"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µthing); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßthing, πTemp001); πE != nil {
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
					// line 157: def __enter__(self):
					πF.SetLineno(157)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__enter__", "build/src/__python__/contextlib.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 158: return self.thing
							πF.SetLineno(158)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßthing, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__enter__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 159: def __exit__(self, *exc_info):
					πF.SetLineno(159)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__exit__", "build/src/__python__/contextlib.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µexc_info *πg.Object = πArgs[1]; _ = µexc_info
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
							// line 160: self.thing.close()
							πF.SetLineno(160)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßthing, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßclose, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__exit__.ToObject(), πTemp004); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp007, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp007 == nil {
				πTemp007 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp007.Call(πF, []*πg.Object{πg.NewStr("closing").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßclosing.ToObject(), πTemp008); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("contextlib", Code)
}
