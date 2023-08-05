package thread
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/thread.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßException := πg.InternStr("Exception")
		ßLock := πg.InternStr("Lock")
		ßLockType := πg.InternStr("LockType")
		ßNewTryableMutex := πg.InternStr("NewTryableMutex")
		ßNone := πg.InternStr("None")
		ßStartThread := πg.InternStr("StartThread")
		ßThreadCount := πg.InternStr("ThreadCount")
		ßTrue := πg.InternStr("True")
		ßTryLock := πg.InternStr("TryLock")
		ßUnlock := πg.InternStr("Unlock")
		ß__enter__ := πg.InternStr("__enter__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__frame__ := πg.InternStr("__frame__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_count := πg.InternStr("_count")
		ß_mutex := πg.InternStr("_mutex")
		ßacquire := πg.InternStr("acquire")
		ßallocate_lock := πg.InternStr("allocate_lock")
		ßappend := πg.InternStr("append")
		ßerror := πg.InternStr("error")
		ßf_back := πg.InternStr("f_back")
		ßget_ident := πg.InternStr("get_ident")
		ßid := πg.InternStr("id")
		ßobject := πg.InternStr("object")
		ßrelease := πg.InternStr("release")
		ßstack_size := πg.InternStr("stack_size")
		ßstart_new_thread := πg.InternStr("start_new_thread")
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
			// line 1: from '__go__/grumpy' import NewTryableMutex, StartThread, ThreadCount
			πF.SetLineno(1)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/grumpy"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßNewTryableMutex, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNewTryableMutex.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßStartThread, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStartThread.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßThreadCount, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßThreadCount.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 4: class error(Exception):
			πF.SetLineno(4)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
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
			_, πE = πg.NewCode("error", "build/src/__python__/thread.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 5: pass
					πF.SetLineno(5)
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("error").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßerror.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 8: def get_ident():
			πF.SetLineno(8)
			πTemp006 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("get_ident", "build/src/__python__/thread.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µf *πg.Object = πg.UnboundLocal; _ = µf
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 9: f = __frame__()
					πF.SetLineno(9)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß__frame__); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µf = πTemp002
					// line 10: while f.f_back:
					πF.SetLineno(10)
					πF.PushCheckpoint(2)
					πTemp003 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp003 {
						πF.PopCheckpoint()
						goto Label3
					}
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßf_back, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 11: f = f.f_back
					πF.SetLineno(11)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßf_back, nil); πE != nil {
						continue
					}
					µf = πTemp001
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 12: return id(f)
					πF.SetLineno(12)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					πTemp005[0] = µf
					if πTemp001, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
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
			if πE = πF.Globals().SetItem(πF, ßget_ident.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 15: class LockType(object):
			πF.SetLineno(15)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp004 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("LockType", "build/src/__python__/thread.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 16: def __init__(self):
					πF.SetLineno(16)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 17: self._mutex = NewTryableMutex()
							πF.SetLineno(17)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNewTryableMutex); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_mutex, πTemp001); πE != nil {
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
					// line 19: def acquire(self, waitflag=1):
					πF.SetLineno(19)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "waitflag", Def: πg.NewInt(1).ToObject()}
					πTemp003 = πg.NewFunction(πg.NewCode("acquire", "build/src/__python__/thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µwaitflag *πg.Object = πArgs[1]; _ = µwaitflag
						var πTemp001 bool
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
							if πE = πg.CheckLocal(πF, µwaitflag, "waitflag"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IsTrue(πF, µwaitflag); πE != nil {
								continue
							}
							if πTemp001 {
								goto Label1
							}
							goto Label2
							// line 20: if waitflag:
							πF.SetLineno(20)
						Label1:
							// line 21: self._mutex.Lock()
							πF.SetLineno(21)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_mutex, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßLock, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 22: return True
							πF.SetLineno(22)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label2
						Label2:
							// line 23: return self._mutex.TryLock()
							πF.SetLineno(23)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_mutex, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßTryLock, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßacquire.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 25: def release(self):
					πF.SetLineno(25)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("release", "build/src/__python__/thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 26: self._mutex.Unlock()
							πF.SetLineno(26)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_mutex, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßUnlock, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßrelease.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 28: def __enter__(self):
					πF.SetLineno(28)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("__enter__", "build/src/__python__/thread.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 29: self.acquire()
							πF.SetLineno(29)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßacquire, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__enter__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 31: def __exit__(self, *args):
					πF.SetLineno(31)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("__exit__", "build/src/__python__/thread.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µargs *πg.Object = πArgs[1]; _ = µargs
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
							// line 32: self.release()
							πF.SetLineno(32)
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
					if πE = πClass.SetItem(πF, ß__exit__.ToObject(), πTemp006); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp005, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp005 == nil {
				πTemp005 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp005.Call(πF, []*πg.Object{πg.NewStr("LockType").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLockType.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 35: def allocate_lock():
			πF.SetLineno(35)
			πTemp006 = make([]πg.Param, 0)
			πTemp003 = πg.NewFunction(πg.NewCode("allocate_lock", "build/src/__python__/thread.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 36: """Dummy implementation of thread.allocate_lock()."""
					πF.SetLineno(36)
					// line 37: return LockType()
					πF.SetLineno(37)
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
			if πE = πF.Globals().SetItem(πF, ßallocate_lock.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 40: def start_new_thread(func, args, kwargs=None):
			πF.SetLineno(40)
			πTemp006 = make([]πg.Param, 3)
			πTemp006[0] = πg.Param{Name: "func", Def: nil}
			πTemp006[1] = πg.Param{Name: "args", Def: nil}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[2] = πg.Param{Name: "kwargs", Def: πTemp007}
			πTemp005 = πg.NewFunction(πg.NewCode("start_new_thread", "build/src/__python__/thread.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfunc *πg.Object = πArgs[0]; _ = µfunc
				var µargs *πg.Object = πArgs[1]; _ = µargs
				var µkwargs *πg.Object = πArgs[2]; _ = µkwargs
				var µl *πg.Object = πg.UnboundLocal; _ = µl
				var µident *πg.Object = πg.UnboundLocal; _ = µident
				var µthread_func *πg.Object = πg.UnboundLocal; _ = µthread_func
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 *πg.Dict
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 []πg.Param
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
					if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µkwargs == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 41: if kwargs is None:
					πF.SetLineno(41)
				Label1:
					// line 42: kwargs = {}
					πF.SetLineno(42)
					πTemp004 = πg.NewDict()
					πTemp001 = πTemp004.ToObject()
					µkwargs = πTemp001
					goto Label2
				Label2:
					// line 43: l = allocate_lock()
					πF.SetLineno(43)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßallocate_lock); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µl = πTemp002
					// line 44: ident = []
					πF.SetLineno(44)
					πTemp005 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp005...).ToObject()
					µident = πTemp001
					// line 45: def thread_func():
					πF.SetLineno(45)
					πTemp006 = make([]πg.Param, 0)
					πTemp001 = πg.NewFunction(πg.NewCode("thread_func", "build/src/__python__/thread.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 46: ident.append(get_ident())
							πF.SetLineno(46)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßget_ident); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µident, "ident"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µident, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 47: l.release()
							πF.SetLineno(47)
							if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µl, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 48: func(*args, **kwargs)
							πF.SetLineno(48)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, µfunc, nil, µargs, nil, µkwargs); πE != nil {
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
					µthread_func = πTemp001
					// line 49: l.acquire()
					πF.SetLineno(49)
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µl, ßacquire, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 50: StartThread(thread_func)
					πF.SetLineno(50)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µthread_func, "thread_func"); πE != nil {
						continue
					}
					πTemp005[0] = µthread_func
					if πTemp002, πE = πg.ResolveGlobal(πF, ßStartThread); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 51: l.acquire()
					πF.SetLineno(51)
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µl, ßacquire, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 52: return ident[0]
					πF.SetLineno(52)
					πTemp002 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µident, "ident"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µident, πTemp002); πE != nil {
						continue
					}
					πR = πTemp007
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßstart_new_thread.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 55: def stack_size(n=0):
			πF.SetLineno(55)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "n", Def: πg.NewInt(0).ToObject()}
			πTemp007 = πg.NewFunction(πg.NewCode("stack_size", "build/src/__python__/thread.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µn *πg.Object = πArgs[0]; _ = µn
				var πTemp001 bool
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
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µn); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label1
					}
					goto Label2
					// line 56: if n:
					πF.SetLineno(56)
				Label1:
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("grumpy does not support setting stack size").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 57: raise error('grumpy does not support setting stack size')
					πF.SetLineno(57)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					goto Label2
				Label2:
					// line 58: return 0
					πF.SetLineno(58)
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
			if πE = πF.Globals().SetItem(πF, ßstack_size.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 61: def _count():
			πF.SetLineno(61)
			πTemp006 = make([]πg.Param, 0)
			πTemp008 = πg.NewFunction(πg.NewCode("_count", "build/src/__python__/thread.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 62: return ThreadCount
					πF.SetLineno(62)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßThreadCount); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_count.ToObject(), πTemp008); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("thread", Code)
}
