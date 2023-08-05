package dummy_thread
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/dummy_thread.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßException := πg.InternStr("Exception")
		ßFalse := πg.InternStr("False")
		ßKeyboardInterrupt := πg.InternStr("KeyboardInterrupt")
		ßLockType := πg.InternStr("LockType")
		ßNone := πg.InternStr("None")
		ßSystemExit := πg.InternStr("SystemExit")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ß__all__ := πg.InternStr("__all__")
		ß__enter__ := πg.InternStr("__enter__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_interrupt := πg.InternStr("_interrupt")
		ß_main := πg.InternStr("_main")
		ß_traceback := πg.InternStr("_traceback")
		ßacquire := πg.InternStr("acquire")
		ßallocate_lock := πg.InternStr("allocate_lock")
		ßargs := πg.InternStr("args")
		ßdict := πg.InternStr("dict")
		ßerror := πg.InternStr("error")
		ßexit := πg.InternStr("exit")
		ßget_ident := πg.InternStr("get_ident")
		ßinterrupt_main := πg.InternStr("interrupt_main")
		ßlocked := πg.InternStr("locked")
		ßlocked_status := πg.InternStr("locked_status")
		ßobject := πg.InternStr("object")
		ßprint_exc := πg.InternStr("print_exc")
		ßrelease := πg.InternStr("release")
		ßstack_size := πg.InternStr("stack_size")
		ßstart_new_thread := πg.InternStr("start_new_thread")
		ßtuple := πg.InternStr("tuple")
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """Drop-in replacement for the thread module.
			πF.SetLineno(1)
			// line 16: __all__ = ['error', 'start_new_thread', 'exit', 'get_ident', 'allocate_lock',
			πF.SetLineno(16)
			πTemp001 = make([]*πg.Object, 7)
			πTemp001[0] = ßerror.ToObject()
			πTemp001[1] = ßstart_new_thread.ToObject()
			πTemp001[2] = ßexit.ToObject()
			πTemp001[3] = ßget_ident.ToObject()
			πTemp001[4] = ßallocate_lock.ToObject()
			πTemp001[5] = ßinterrupt_main.ToObject()
			πTemp001[6] = ßLockType.ToObject()
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 19: import traceback as _traceback
			πF.SetLineno(19)
			if πTemp001, πE = πg.ImportModule(πF, "traceback"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ß_traceback.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 21: class error(Exception):
			πF.SetLineno(21)
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
			_, πE = πg.NewCode("error", "build/src/__python__/dummy_thread.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 22: """Dummy implementation of thread.error."""
					πF.SetLineno(22)
					// line 24: def __init__(self, *args):
					πF.SetLineno(24)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/dummy_thread.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µargs *πg.Object = πArgs[1]; _ = µargs
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 25: self.args = args
							πF.SetLineno(25)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µargs); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßargs, πTemp001); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("error").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßerror.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 27: def start_new_thread(function, args, kwargs={}):
			πF.SetLineno(27)
			πTemp006 = make([]πg.Param, 3)
			πTemp006[0] = πg.Param{Name: "function", Def: nil}
			πTemp006[1] = πg.Param{Name: "args", Def: nil}
			πTemp003 = πg.NewDict()
			πTemp004 = πTemp003.ToObject()
			πTemp006[2] = πg.Param{Name: "kwargs", Def: πTemp004}
			πTemp002 = πg.NewFunction(πg.NewCode("start_new_thread", "build/src/__python__/dummy_thread.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfunction *πg.Object = πArgs[0]; _ = µfunction
				var µargs *πg.Object = πArgs[1]; _ = µargs
				var µkwargs *πg.Object = πArgs[2]; _ = µkwargs
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
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 *πg.BaseException
				_ = πTemp007
				var πTemp008 *πg.Traceback
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 6: goto Label6
					default: panic("unexpected function state")
					}
					// line 28: """Dummy implementation of thread.start_new_thread().
					πF.SetLineno(28)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp002[0] = µargs
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp002 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp005
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.NE(πF, πTemp004, πTemp005); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label1
					}
					goto Label2
					// line 40: if type(args) != type(tuple()):
					πF.SetLineno(40)
				Label1:
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("2nd arg must be a tuple").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 41: raise TypeError("2nd arg must be a tuple")
					πF.SetLineno(41)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label2
				Label2:
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
						continue
					}
					πTemp002[0] = µkwargs
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp002 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp005
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.NE(πF, πTemp004, πTemp005); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					goto Label4
					// line 42: if type(kwargs) != type(dict()):
					πF.SetLineno(42)
				Label3:
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("3rd arg must be a dict").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 43: raise TypeError("3rd arg must be a dict")
					πF.SetLineno(43)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label4
				Label4:
					// line 44: global _main
					πF.SetLineno(44)
					// line 45: _main = False
					πF.SetLineno(45)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					if πE = πF.Globals().SetItem(πF, ß_main.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 46: try:
					πF.SetLineno(46)
					πF.PushCheckpoint(6)
					// line 47: function(*args, **kwargs)
					πF.SetLineno(47)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfunction, "function"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Invoke(πF, µfunction, nil, µargs, nil, µkwargs); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					goto Label5
				Label6:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßSystemExit); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label7
					}
					goto Label8
					// line 48: except SystemExit:
					πF.SetLineno(48)
				Label7:
					// line 49: pass
					πF.SetLineno(49)
					πF.RestoreExc(nil, nil)
					goto Label5
					// line 50: except:
					πF.SetLineno(50)
				Label8:
					// line 51: _traceback.print_exc()
					πF.SetLineno(51)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_traceback); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßprint_exc, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πF.RestoreExc(nil, nil)
					goto Label5
				Label5:
					// line 52: _main = True
					πF.SetLineno(52)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πE = πF.Globals().SetItem(πF, ß_main.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 53: global _interrupt
					πF.SetLineno(53)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_interrupt); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label9
					}
					goto Label10
					// line 54: if _interrupt:
					πF.SetLineno(54)
				Label9:
					// line 55: _interrupt = False
					πF.SetLineno(55)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					if πE = πF.Globals().SetItem(πF, ß_interrupt.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyboardInterrupt); πE != nil {
						continue
					}
					// line 56: raise KeyboardInterrupt
					πF.SetLineno(56)
					πE = πF.Raise(πTemp001, nil, nil)
					continue
					goto Label10
				Label10:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßstart_new_thread.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 58: def exit():
			πF.SetLineno(58)
			πTemp006 = make([]πg.Param, 0)
			πTemp004 = πg.NewFunction(πg.NewCode("exit", "build/src/__python__/dummy_thread.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 59: """Dummy implementation of thread.exit()."""
					πF.SetLineno(59)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßSystemExit); πE != nil {
						continue
					}
					// line 60: raise SystemExit
					πF.SetLineno(60)
					πE = πF.Raise(πTemp001, nil, nil)
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßexit.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 62: def get_ident():
			πF.SetLineno(62)
			πTemp006 = make([]πg.Param, 0)
			πTemp005 = πg.NewFunction(πg.NewCode("get_ident", "build/src/__python__/dummy_thread.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 63: """Dummy implementation of thread.get_ident().
					πF.SetLineno(63)
					// line 69: return -1
					πF.SetLineno(69)
					if πTemp001, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßget_ident.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 71: def allocate_lock():
			πF.SetLineno(71)
			πTemp006 = make([]πg.Param, 0)
			πTemp007 = πg.NewFunction(πg.NewCode("allocate_lock", "build/src/__python__/dummy_thread.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 72: """Dummy implementation of thread.allocate_lock()."""
					πF.SetLineno(72)
					// line 73: return LockType()
					πF.SetLineno(73)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßLockType); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßallocate_lock.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 75: def stack_size(size=None):
			πF.SetLineno(75)
			πTemp006 = make([]πg.Param, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[0] = πg.Param{Name: "size", Def: πTemp009}
			πTemp008 = πg.NewFunction(πg.NewCode("stack_size", "build/src/__python__/dummy_thread.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsize *πg.Object = πArgs[0]; _ = µsize
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
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
					// line 76: """Dummy implementation of thread.stack_size()."""
					πF.SetLineno(76)
					if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µsize != πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 77: if size is not None:
					πF.SetLineno(77)
				Label1:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("setting thread stack size not supported").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 78: raise error("setting thread stack size not supported")
					πF.SetLineno(78)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label2
				Label2:
					// line 79: return 0
					πF.SetLineno(79)
					πR = πg.NewInt(0).ToObject()
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßstack_size.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 81: class LockType(object):
			πF.SetLineno(81)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp011, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp001[0] = πTemp011
			πTemp003 = πg.NewDict()
			if πTemp009, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp009); πE != nil {
				continue
			}
			_, πE = πg.NewCode("LockType", "build/src/__python__/dummy_thread.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 82: """Class implementing dummy implementation of thread.LockType.
					πF.SetLineno(82)
					// line 92: def __init__(self):
					πF.SetLineno(92)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/dummy_thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 93: self.locked_status = False
							πF.SetLineno(93)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlocked_status, πTemp002); πE != nil {
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
					// line 95: def acquire(self, waitflag=None):
					πF.SetLineno(95)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "waitflag", Def: πTemp004}
					πTemp003 = πg.NewFunction(πg.NewCode("acquire", "build/src/__python__/dummy_thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µwaitflag *πg.Object = πArgs[1]; _ = µwaitflag
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
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
							// line 96: """Dummy implementation of acquire().
							πF.SetLineno(96)
							if πE = πg.CheckLocal(πF, µwaitflag, "waitflag"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µwaitflag == πTemp004).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µwaitflag, "waitflag"); πE != nil {
								continue
							}
							πTemp001 = µwaitflag
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßlocked_status, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 106: if waitflag is None or waitflag:
							πF.SetLineno(106)
						Label2:
							// line 107: self.locked_status = True
							πF.SetLineno(107)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlocked_status, πTemp003); πE != nil {
								continue
							}
							// line 108: return True
							πF.SetLineno(108)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label5
							// line 110: if not self.locked_status:
							πF.SetLineno(110)
						Label3:
							// line 111: self.locked_status = True
							πF.SetLineno(111)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlocked_status, πTemp003); πE != nil {
								continue
							}
							// line 112: return True
							πF.SetLineno(112)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label5
						Label4:
							// line 114: return False
							πF.SetLineno(114)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label5
						Label5:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßacquire.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 116: __enter__ = acquire
					πF.SetLineno(116)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßacquire); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__enter__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 118: def __exit__(self, typ, val, tb):
					πF.SetLineno(118)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "typ", Def: nil}
					πTemp002[2] = πg.Param{Name: "val", Def: nil}
					πTemp002[3] = πg.Param{Name: "tb", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__exit__", "build/src/__python__/dummy_thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtyp *πg.Object = πArgs[1]; _ = µtyp
						var µval *πg.Object = πArgs[2]; _ = µval
						var µtb *πg.Object = πArgs[3]; _ = µtb
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
							// line 119: self.release()
							πF.SetLineno(119)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
					// line 121: def release(self):
					πF.SetLineno(121)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("release", "build/src/__python__/dummy_thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 122: """Release the dummy lock."""
							πF.SetLineno(122)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlocked_status, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 125: if not self.locked_status:
							πF.SetLineno(125)
						Label1:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
								continue
							}
							// line 126: raise error
							πF.SetLineno(126)
							πE = πF.Raise(πTemp001, nil, nil)
							continue
							goto Label2
						Label2:
							// line 127: self.locked_status = False
							πF.SetLineno(127)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlocked_status, πTemp002); πE != nil {
								continue
							}
							// line 128: return True
							πF.SetLineno(128)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
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
					if πE = πClass.SetItem(πF, ßrelease.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 130: def locked(self):
					πF.SetLineno(130)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("locked", "build/src/__python__/dummy_thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 131: return self.locked_status
							πF.SetLineno(131)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlocked_status, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßlocked.ToObject(), πTemp006); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp010, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp010 == nil {
				πTemp010 = πg.TypeType.ToObject()
			}
			if πTemp011, πE = πTemp010.Call(πF, []*πg.Object{πg.NewStr("LockType").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLockType.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 134: _interrupt = False
			πF.SetLineno(134)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_interrupt.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 136: _main = True
			πF.SetLineno(136)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_main.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 138: def interrupt_main():
			πF.SetLineno(138)
			πTemp006 = make([]πg.Param, 0)
			πTemp009 = πg.NewFunction(πg.NewCode("interrupt_main", "build/src/__python__/dummy_thread.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 139: """Set _interrupt flag to True to have start_new_thread raise
					πF.SetLineno(139)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_main); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 141: if _main:
					πF.SetLineno(141)
				Label1:
					if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyboardInterrupt); πE != nil {
						continue
					}
					// line 142: raise KeyboardInterrupt
					πF.SetLineno(142)
					πE = πF.Raise(πTemp001, nil, nil)
					continue
					goto Label3
				Label2:
					// line 144: global _interrupt
					πF.SetLineno(144)
					// line 145: _interrupt = True
					πF.SetLineno(145)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πE = πF.Globals().SetItem(πF, ß_interrupt.ToObject(), πTemp001); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßinterrupt_main.ToObject(), πTemp009); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("dummy_thread", Code)
}
