package _functools
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/_functools.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßNone := πg.InternStr("None")
		ßStopIteration := πg.InternStr("StopIteration")
		ßTypeError := πg.InternStr("TypeError")
		ß__all__ := πg.InternStr("__all__")
		ß__call__ := πg.InternStr("__call__")
		ß__delattr__ := πg.InternStr("__delattr__")
		ß__dict__ := πg.InternStr("__dict__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__reduce__ := πg.InternStr("__reduce__")
		ß__setstate__ := πg.InternStr("__setstate__")
		ß__slots__ := πg.InternStr("__slots__")
		ß_args := πg.InternStr("_args")
		ß_func := πg.InternStr("_func")
		ß_keywords := πg.InternStr("_keywords")
		ßargs := πg.InternStr("args")
		ßcallable := πg.InternStr("callable")
		ßclear := πg.InternStr("clear")
		ßdict := πg.InternStr("dict")
		ßfunc := πg.InternStr("func")
		ßisinstance := πg.InternStr("isinstance")
		ßiter := πg.InternStr("iter")
		ßiteritems := πg.InternStr("iteritems")
		ßkeywords := πg.InternStr("keywords")
		ßlen := πg.InternStr("len")
		ßnext := πg.InternStr("next")
		ßobject := πg.InternStr("object")
		ßpartial := πg.InternStr("partial")
		ßproperty := πg.InternStr("property")
		ßreduce := πg.InternStr("reduce")
		ßtuple := πg.InternStr("tuple")
		ßtype := πg.InternStr("type")
		ßupdate := πg.InternStr("update")
		var πTemp001 []*πg.Object
		_ = πTemp001
		var πTemp002 *πg.Object
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """ Supplies the internal functions for functools.py in the standard library """
			πF.SetLineno(1)
			// line 3: __all__ = ['reduce', 'partial']
			πF.SetLineno(3)
			πTemp001 = make([]*πg.Object, 2)
			πTemp001[0] = ßreduce.ToObject()
			πTemp001[1] = ßpartial.ToObject()
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 6: def reduce(function, iterable, initializer=None):
			πF.SetLineno(6)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "function", Def: nil}
			πTemp003[1] = πg.Param{Name: "iterable", Def: nil}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[2] = πg.Param{Name: "initializer", Def: πTemp004}
			πTemp002 = πg.NewFunction(πg.NewCode("reduce", "build/src/__python__/_functools.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfunction *πg.Object = πArgs[0]; _ = µfunction
				var µiterable *πg.Object = πArgs[1]; _ = µiterable
				var µinitializer *πg.Object = πArgs[2]; _ = µinitializer
				var µit *πg.Object = πg.UnboundLocal; _ = µit
				var µaccum_value *πg.Object = πg.UnboundLocal; _ = µaccum_value
				var µx *πg.Object = πg.UnboundLocal; _ = µx
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 *πg.BaseException
				_ = πTemp005
				var πTemp006 *πg.Traceback
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 4: goto Label4
					case 6: goto Label6
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					// line 7: it = iter(iterable)
					πF.SetLineno(7)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
						continue
					}
					πTemp001[0] = µiterable
					if πTemp002, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µit = πTemp003
					if πE = πg.CheckLocal(πF, µinitializer, "initializer"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µinitializer == πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 8: if initializer is None:
					πF.SetLineno(8)
				Label1:
					// line 9: try:
					πF.SetLineno(9)
					πF.PushCheckpoint(4)
					// line 10: initializer = next(it)
					πF.SetLineno(10)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
						continue
					}
					πTemp001[0] = µit
					if πTemp002, πE = πg.ResolveGlobal(πF, ßnext); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µinitializer = πTemp003
					πF.PopCheckpoint()
					goto Label3
				Label4:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp005, πTemp006 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label5
					}
					πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
					continue
					// line 11: except StopIteration:
					πF.SetLineno(11)
				Label5:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("reduce() of empty sequence with no initial value").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 12: raise TypeError('reduce() of empty sequence with no initial value')
					πF.SetLineno(12)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label3
				Label3:
					goto Label2
				Label2:
					// line 13: accum_value = initializer
					πF.SetLineno(13)
					if πE = πg.CheckLocal(πF, µinitializer, "initializer"); πE != nil {
						continue
					}
					µaccum_value = µinitializer
					if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µit); πE != nil {
						continue
					}
					πF.PushCheckpoint(7)
					πTemp004 = false
				Label6:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label8
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
						µx = πTemp003
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(6)            
					// line 15: accum_value = function(accum_value, x)
					πF.SetLineno(15)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µaccum_value, "accum_value"); πE != nil {
						continue
					}
					πTemp001[0] = µaccum_value
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[1] = µx
					if πE = πg.CheckLocal(πF, µfunction, "function"); πE != nil {
						continue
					}
					if πTemp003, πE = µfunction.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µaccum_value = πTemp003
					continue
				Label7:
					if πE != nil || πR != nil {
						continue
					}
				Label8:
					// line 16: return accum_value
					πF.SetLineno(16)
					if πE = πg.CheckLocal(πF, µaccum_value, "accum_value"); πE != nil {
						continue
					}
					πR = µaccum_value
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßreduce.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 18: class partial(object):
			πF.SetLineno(18)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp001[0] = πTemp007
			πTemp005 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("partial", "build/src/__python__/_functools.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
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
					// line 19: """
					πF.SetLineno(19)
					// line 24: __slots__ = ('_func', '_args', '_keywords', '__dict__')
					πF.SetLineno(24)
					πTemp001 = πg.NewTuple4(ß_func.ToObject(), ß_args.ToObject(), ß_keywords.ToObject(), ß__dict__.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ß__slots__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 26: def __init__(*args, **keywords):
					πF.SetLineno(26)
					πTemp002 = make([]πg.Param, 0)
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/_functools.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µargs *πg.Object = πArgs[0]; _ = µargs
						var µkeywords *πg.Object = πArgs[1]; _ = µkeywords
						var µself *πg.Object = πg.UnboundLocal; _ = µself
						var µfunc *πg.Object = πg.UnboundLocal; _ = µfunc
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp002[0] = µargs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.LT(πF, πTemp004, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 27: if len(args) < 2:
							πF.SetLineno(27)
						Label1:
							πTemp002 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp006[0] = µargs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("__init__() takes at least 2 arguments (%d given)").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 28: raise TypeError('__init__() takes at least 2 arguments (%d given)'
							πF.SetLineno(28)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							// line 30: self, func, args = args[0], args[1], args[2:]
							πF.SetLineno(30)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µargs, πTemp003); πE != nil {
								continue
							}
							πTemp003 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µargs, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(2).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µargs, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(πTemp004, πTemp007, πTemp008).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp001); πE != nil {
								continue
							}
							µself = πTemp003
							µfunc = πTemp004
							µargs = πTemp007
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							πTemp002[0] = µfunc
							if πTemp003, πE = πg.ResolveGlobal(πF, ßcallable); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 31: if not callable(func):
							πF.SetLineno(31)
						Label3:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("the first argument must be callable").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 32: raise TypeError("the first argument must be callable")
							πF.SetLineno(32)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label4
						Label4:
							// line 33: self._func = func
							πF.SetLineno(33)
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µfunc); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_func, πTemp001); πE != nil {
								continue
							}
							// line 34: self._args = args
							πF.SetLineno(34)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µargs); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_args, πTemp001); πE != nil {
								continue
							}
							// line 35: self._keywords = keywords
							πF.SetLineno(35)
							if πE = πg.CheckLocal(πF, µkeywords, "keywords"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µkeywords); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_keywords, πTemp001); πE != nil {
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
					// line 37: def __delattr__(self, key):
					πF.SetLineno(37)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__delattr__", "build/src/__python__/_functools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µkey, ß__dict__.ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 38: if key == '__dict__':
							πF.SetLineno(38)
						Label1:
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("a partial object's dictionary may not be deleted").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 39: raise TypeError("a partial object's dictionary may not be deleted")
							πF.SetLineno(39)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							// line 40: object.__delattr__(self, key)
							πF.SetLineno(40)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp003[1] = µkey
							if πTemp001, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ß__delattr__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__delattr__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 43: def func(self):
					πF.SetLineno(43)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("func", "build/src/__python__/_functools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 44: return self._func
							πF.SetLineno(44)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_func, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßfunc.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 46: func = property(func)
					πF.SetLineno(46)
					πTemp005 = πF.MakeArgs(1)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßfunc); πE != nil {
						continue
					}
					πTemp005[0] = πTemp006
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πClass.SetItem(πF, ßfunc.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 49: def args(self):
					πF.SetLineno(49)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("args", "build/src/__python__/_functools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 50: return self._args
							πF.SetLineno(50)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_args, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßargs.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 52: args = property(args)
					πF.SetLineno(52)
					πTemp005 = πF.MakeArgs(1)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßargs); πE != nil {
						continue
					}
					πTemp005[0] = πTemp007
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πClass.SetItem(πF, ßargs.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 55: def keywords(self):
					πF.SetLineno(55)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("keywords", "build/src/__python__/_functools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 56: return self._keywords
							πF.SetLineno(56)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_keywords, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßkeywords.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 58: keywords = property(keywords)
					πF.SetLineno(58)
					πTemp005 = πF.MakeArgs(1)
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßkeywords); πE != nil {
						continue
					}
					πTemp005[0] = πTemp008
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp008.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πClass.SetItem(πF, ßkeywords.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 60: def __call__(self, *fargs, **fkeywords):
					πF.SetLineno(60)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("__call__", "build/src/__python__/_functools.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfargs *πg.Object = πArgs[1]; _ = µfargs
						var µfkeywords *πg.Object = πArgs[2]; _ = µfkeywords
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_keywords, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 61: if self._keywords:
							πF.SetLineno(61)
						Label1:
							// line 62: fkeywords = dict(self._keywords, **fkeywords)
							πF.SetLineno(62)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_keywords, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µfkeywords, "fkeywords"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Invoke(πF, πTemp001, πTemp003, nil, nil, µfkeywords); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µfkeywords = πTemp004
							goto Label2
						Label2:
							// line 63: return self._func(*(self._args + fargs), **fkeywords)
							πF.SetLineno(63)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_args, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfargs, "fargs"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp004, µfargs); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfkeywords, "fkeywords"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_func, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Invoke(πF, πTemp004, nil, πTemp001, nil, µfkeywords); πE != nil {
								continue
							}
							πR = πTemp005
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__call__.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 65: def __reduce__(self):
					πF.SetLineno(65)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("__reduce__", "build/src/__python__/_functools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
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
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 66: d = dict((k, v) for k, v in self.__dict__.iteritems() if k not in
							πF.SetLineno(66)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/_functools.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µk *πg.Object = πg.UnboundLocal; _ = µk
								var µv *πg.Object = πg.UnboundLocal; _ = µv
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
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µself, ß__dict__, nil); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßiteritems, nil); πE != nil {
											continue
										}
										if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
											if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
												continue
											}
											µk = πTemp003
											µv = πTemp006
										}
										if πE != nil || !πTemp005 {
											continue
										}
										πF.PushCheckpoint(1)            
										if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
											continue
										}
										πTemp003 = πg.NewTuple3(ß_func.ToObject(), ß_args.ToObject(), ß_keywords.ToObject()).ToObject()
										if πTemp005, πE = πg.Contains(πF, πTemp003, µk); πE != nil {
											continue
										}
										πTemp002 = πg.GetBool(!πTemp005).ToObject()
										if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
											continue
										}
										if πTemp005 {
											goto Label4
										}
										goto Label5
										// line 66: d = dict((k, v) for k, v in self.__dict__.iteritems() if k not in
										πF.SetLineno(66)
									Label4:
										// line 66: d = dict((k, v) for k, v in self.__dict__.iteritems() if k not in
										πF.SetLineno(66)
										if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
											continue
										}
										πTemp002 = πg.NewTuple2(µk, µv).ToObject()
										πF.PushCheckpoint(6)
										return πTemp002, nil
									Label6:
										πTemp003 = πSent
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
							πTemp001[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µd = πTemp005
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[0] = µd
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.Eq(πF, πTemp006, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label1
							}
							goto Label2
							// line 68: if len(d) == 0:
							πF.SetLineno(68)
						Label1:
							// line 69: d = None
							πF.SetLineno(69)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µd = πTemp004
							goto Label2
						Label2:
							// line 70: return (type(self), (self._func,),
							πF.SetLineno(70)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ß_func, nil); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple1(πTemp008).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ß_func, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µself, ß_args, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ß_keywords, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp008 = πg.NewTuple4(πTemp009, πTemp010, πTemp011, µd).ToObject()
							πTemp004 = πg.NewTuple3(πTemp006, πTemp005, πTemp008).ToObject()
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
					if πE = πClass.SetItem(πF, ß__reduce__.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 73: def __setstate__(self, state):
					πF.SetLineno(73)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "state", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("__setstate__", "build/src/__python__/_functools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstate *πg.Object = πArgs[1]; _ = µstate
						var µfunc *πg.Object = πg.UnboundLocal; _ = µfunc
						var µargs *πg.Object = πg.UnboundLocal; _ = µargs
						var µkeywords *πg.Object = πg.UnboundLocal; _ = µkeywords
						var µd *πg.Object = πg.UnboundLocal; _ = µd
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
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 *πg.Dict
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							πTemp004[0] = µstate
							if πTemp005, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
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
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							πTemp004[0] = µstate
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.NE(πF, πTemp006, πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 74: if not isinstance(state, tuple) or len(state) != 4:
							πF.SetLineno(74)
						Label2:
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("invalid partial state").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 75: raise TypeError("invalid partial state")
							πF.SetLineno(75)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label3
						Label3:
							// line 77: func, args, keywords, d = state
							πF.SetLineno(77)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, µstate); πE != nil {
								continue
							}
							µfunc = πTemp001
							µargs = πTemp003
							µkeywords = πTemp005
							µd = πTemp006
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							πTemp004[0] = µfunc
							if πTemp005, πE = πg.ResolveGlobal(πF, ßcallable); πE != nil {
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
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp004[0] = µargs
							if πTemp005, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
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
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µkeywords, "keywords"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(µkeywords != πTemp006).ToObject()
							πTemp003 = πTemp005
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp007 {
								goto Label5
							}
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µkeywords, "keywords"); πE != nil {
								continue
							}
							πTemp004[0] = µkeywords
							if πTemp006, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							πTemp004[1] = πTemp006
							if πTemp006, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp009, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(!πTemp009).ToObject()
							πTemp003 = πTemp005
						Label5:
							πTemp001 = πTemp003
						Label4:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label6
							}
							goto Label7
							// line 79: if (not callable(func) or not isinstance(args, tuple) or
							πF.SetLineno(79)
						Label6:
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("invalid partial state").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 81: raise TypeError("invalid partial state")
							πF.SetLineno(81)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label7
						Label7:
							// line 83: self._func = func
							πF.SetLineno(83)
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µfunc); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_func, πTemp001); πE != nil {
								continue
							}
							// line 84: self._args = tuple(args)
							πF.SetLineno(84)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp004[0] = µargs
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_args, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkeywords, "keywords"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µkeywords == πTemp003).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label8
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkeywords, "keywords"); πE != nil {
								continue
							}
							πTemp004[0] = µkeywords
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp005 != πTemp003).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label9
							}
							goto Label10
							// line 86: if keywords is None:
							πF.SetLineno(86)
						Label8:
							// line 87: keywords = {}
							πF.SetLineno(87)
							πTemp010 = πg.NewDict()
							πTemp001 = πTemp010.ToObject()
							µkeywords = πTemp001
							goto Label10
							// line 88: elif type(keywords) is not dict:
							πF.SetLineno(88)
						Label9:
							// line 89: keywords = dict(keywords)
							πF.SetLineno(89)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkeywords, "keywords"); πE != nil {
								continue
							}
							πTemp004[0] = µkeywords
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µkeywords = πTemp003
							goto Label10
						Label10:
							// line 90: self._keywords = keywords
							πF.SetLineno(90)
							if πE = πg.CheckLocal(πF, µkeywords, "keywords"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µkeywords); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_keywords, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µd == πTemp003).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label11
							}
							goto Label12
							// line 92: if d is None:
							πF.SetLineno(92)
						Label11:
							// line 93: self.__dict__.clear()
							πF.SetLineno(93)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__dict__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßclear, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label13
						Label12:
							// line 95: self.__dict__.update(d)
							πF.SetLineno(95)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp004[0] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__dict__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label13
						Label13:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__setstate__.ToObject(), πTemp010); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("partial").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßpartial.ToObject(), πTemp007); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("_functools", Code)
}
