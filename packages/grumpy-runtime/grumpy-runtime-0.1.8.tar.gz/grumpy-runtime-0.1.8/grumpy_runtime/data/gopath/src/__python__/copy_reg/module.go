package copy_reg
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/copy_reg.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßAttributeError := πg.InternStr("AttributeError")
		ßNameError := πg.InternStr("NameError")
		ßNone := πg.InternStr("None")
		ßTypeError := πg.InternStr("TypeError")
		ßValueError := πg.InternStr("ValueError")
		ß_HEAPTYPE := πg.InternStr("_HEAPTYPE")
		ß__ := πg.InternStr("__")
		ß__all__ := πg.InternStr("__all__")
		ß__call__ := πg.InternStr("__call__")
		ß__class__ := πg.InternStr("__class__")
		ß__dict__ := πg.InternStr("__dict__")
		ß__flags__ := πg.InternStr("__flags__")
		ß__getstate__ := πg.InternStr("__getstate__")
		ß__init__ := πg.InternStr("__init__")
		ß__mro__ := πg.InternStr("__mro__")
		ß__name__ := πg.InternStr("__name__")
		ß__new__ := πg.InternStr("__new__")
		ß__newobj__ := πg.InternStr("__newobj__")
		ß__slotnames__ := πg.InternStr("__slotnames__")
		ß__slots__ := πg.InternStr("__slots__")
		ß__weakref__ := πg.InternStr("__weakref__")
		ß_extension_cache := πg.InternStr("_extension_cache")
		ß_extension_registry := πg.InternStr("_extension_registry")
		ß_inverted_registry := πg.InternStr("_inverted_registry")
		ß_reconstructor := πg.InternStr("_reconstructor")
		ß_reduce_ex := πg.InternStr("_reduce_ex")
		ß_slotnames := πg.InternStr("_slotnames")
		ßadd_extension := πg.InternStr("add_extension")
		ßappend := πg.InternStr("append")
		ßbasestring := πg.InternStr("basestring")
		ßclear := πg.InternStr("clear")
		ßclear_extension_cache := πg.InternStr("clear_extension_cache")
		ßcomplex := πg.InternStr("complex")
		ßconstructor := πg.InternStr("constructor")
		ßdispatch_table := πg.InternStr("dispatch_table")
		ßendswith := πg.InternStr("endswith")
		ßget := πg.InternStr("get")
		ßgetattr := πg.InternStr("getattr")
		ßhasattr := πg.InternStr("hasattr")
		ßimag := πg.InternStr("imag")
		ßint := πg.InternStr("int")
		ßisinstance := πg.InternStr("isinstance")
		ßobject := πg.InternStr("object")
		ßpickle := πg.InternStr("pickle")
		ßpickle_complex := πg.InternStr("pickle_complex")
		ßreal := πg.InternStr("real")
		ßremove_extension := πg.InternStr("remove_extension")
		ßstartswith := πg.InternStr("startswith")
		var πTemp001 []*πg.Object
		_ = πTemp001
		var πTemp002 *πg.Object
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
		var πTemp009 *πg.BaseException
		_ = πTemp009
		var πTemp010 *πg.Traceback
		_ = πTemp010
		var πTemp011 bool
		_ = πTemp011
		var πTemp012 *πg.Object
		_ = πTemp012
		var πTemp013 *πg.Object
		_ = πTemp013
		var πTemp014 *πg.Object
		_ = πTemp014
		var πTemp015 *πg.Object
		_ = πTemp015
		var πTemp016 *πg.Object
		_ = πTemp016
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 2: goto Label2
			default: panic("unexpected function state")
			}
			// line 1: """Helper to provide extensibility for pickle/cPickle.
			πF.SetLineno(1)
			// line 9: __all__ = ["pickle", "constructor",
			πF.SetLineno(9)
			πTemp001 = make([]*πg.Object, 5)
			πTemp001[0] = ßpickle.ToObject()
			πTemp001[1] = ßconstructor.ToObject()
			πTemp001[2] = ßadd_extension.ToObject()
			πTemp001[3] = ßremove_extension.ToObject()
			πTemp001[4] = ßclear_extension_cache.ToObject()
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 12: dispatch_table = {}
			πF.SetLineno(12)
			πTemp003 = πg.NewDict()
			πTemp002 = πTemp003.ToObject()
			if πE = πF.Globals().SetItem(πF, ßdispatch_table.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 14: def pickle(ob_type, pickle_function, constructor_ob=None):
			πF.SetLineno(14)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "ob_type", Def: nil}
			πTemp004[1] = πg.Param{Name: "pickle_function", Def: nil}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "constructor_ob", Def: πTemp005}
			πTemp002 = πg.NewFunction(πg.NewCode("pickle", "build/src/__python__/copy_reg.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µob_type *πg.Object = πArgs[0]; _ = µob_type
				var µpickle_function *πg.Object = πArgs[1]; _ = µpickle_function
				var µconstructor_ob *πg.Object = πArgs[2]; _ = µconstructor_ob
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpickle_function, "pickle_function"); πE != nil {
						continue
					}
					πTemp002[0] = µpickle_function
					πTemp002[1] = ß__call__.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
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
						goto Label1
					}
					goto Label2
					// line 18: if not hasattr(pickle_function, '__call__'):
					πF.SetLineno(18)
				Label1:
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("reduction functions must be callable").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 19: raise TypeError("reduction functions must be callable")
					πF.SetLineno(19)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label2
				Label2:
					// line 20: dispatch_table[ob_type] = pickle_function
					πF.SetLineno(20)
					if πE = πg.CheckLocal(πF, µpickle_function, "pickle_function"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µpickle_function); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßdispatch_table); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µob_type, "ob_type"); πE != nil {
						continue
					}
					πTemp004 = µob_type
					if πE = πg.SetItem(πF, πTemp003, πTemp004, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µconstructor_ob, "constructor_ob"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µconstructor_ob != πTemp003).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label3
					}
					goto Label4
					// line 24: if constructor_ob is not None:
					πF.SetLineno(24)
				Label3:
					// line 25: constructor(constructor_ob)
					πF.SetLineno(25)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µconstructor_ob, "constructor_ob"); πE != nil {
						continue
					}
					πTemp002[0] = µconstructor_ob
					if πTemp001, πE = πg.ResolveGlobal(πF, ßconstructor); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
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
			if πE = πF.Globals().SetItem(πF, ßpickle.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 27: def constructor(object):
			πF.SetLineno(27)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "object", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("constructor", "build/src/__python__/copy_reg.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobject *πg.Object = πArgs[0]; _ = µobject
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
						continue
					}
					πTemp002[0] = µobject
					πTemp002[1] = ß__call__.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
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
						goto Label1
					}
					goto Label2
					// line 28: if not hasattr(object, '__call__'):
					πF.SetLineno(28)
				Label1:
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("constructors must be callable").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 29: raise TypeError("constructors must be callable")
					πF.SetLineno(29)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label2
				Label2:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßconstructor.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 33: try:
			πF.SetLineno(33)
			πF.PushCheckpoint(2)
			// line 34: complex
			πF.SetLineno(34)
			if πTemp006, πE = πg.ResolveGlobal(πF, ßcomplex); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			// line 39: def pickle_complex(c):
			πF.SetLineno(39)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "c", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("pickle_complex", "build/src/__python__/copy_reg.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µc *πg.Object = πArgs[0]; _ = µc
				var πTemp001 *πg.Object
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
					// line 40: return complex, (c.real, c.imag)
					πF.SetLineno(40)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßcomplex); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µc, ßreal, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µc, ßimag, nil); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
					πTemp001 = πg.NewTuple2(πTemp002, πTemp003).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßpickle_complex.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 42: pickle(complex, pickle_complex, complex)
			πF.SetLineno(42)
			πTemp001 = πF.MakeArgs(3)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßcomplex); πE != nil {
				continue
			}
			πTemp001[0] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßpickle_complex); πE != nil {
				continue
			}
			πTemp001[1] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßcomplex); πE != nil {
				continue
			}
			πTemp001[2] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßpickle); πE != nil {
				continue
			}
			if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			goto Label1
		Label2:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp009, πTemp010 = πF.ExcInfo()
			if πTemp007, πE = πg.ResolveGlobal(πF, ßNameError); πE != nil {
				continue
			}
			if πTemp011, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp007); πE != nil {
				continue
			}
			if πTemp011 {
				goto Label3
			}
			πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
			continue
			// line 35: except NameError:
			πF.SetLineno(35)
		Label3:
			// line 36: pass
			πF.SetLineno(36)
			πF.RestoreExc(nil, nil)
			goto Label1
		Label1:
			// line 46: def _reconstructor(cls, base, state):
			πF.SetLineno(46)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "cls", Def: nil}
			πTemp004[1] = πg.Param{Name: "base", Def: nil}
			πTemp004[2] = πg.Param{Name: "state", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("_reconstructor", "build/src/__python__/copy_reg.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcls *πg.Object = πArgs[0]; _ = µcls
				var µbase *πg.Object = πArgs[1]; _ = µbase
				var µstate *πg.Object = πArgs[2]; _ = µstate
				var µobj *πg.Object = πg.UnboundLocal; _ = µobj
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
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
					if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µbase == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 47: if base is object:
					πF.SetLineno(47)
				Label1:
					// line 48: obj = object.__new__(cls)
					πF.SetLineno(48)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					πTemp004[0] = µcls
					if πTemp001, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß__new__, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µobj = πTemp001
					goto Label3
				Label2:
					// line 50: obj = base.__new__(cls, state)
					πF.SetLineno(50)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					πTemp004[0] = µcls
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					πTemp004[1] = µstate
					if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µbase, ß__new__, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µobj = πTemp002
					if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µbase, ß__init__, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ß__init__, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, πTemp002, πTemp006); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label4
					}
					goto Label5
					// line 51: if base.__init__ != object.__init__:
					πF.SetLineno(51)
				Label4:
					// line 52: base.__init__(obj, state)
					πF.SetLineno(52)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp004[0] = µobj
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					πTemp004[1] = µstate
					if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µbase, ß__init__, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					goto Label5
				Label5:
					goto Label3
				Label3:
					// line 53: return obj
					πF.SetLineno(53)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πR = µobj
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_reconstructor.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 55: _HEAPTYPE = 1<<9
			πF.SetLineno(55)
			if πTemp008, πE = πg.LShift(πF, πg.NewInt(1).ToObject(), πg.NewInt(9).ToObject()); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_HEAPTYPE.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 59: def _reduce_ex(self, proto):
			πF.SetLineno(59)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "self", Def: nil}
			πTemp004[1] = πg.Param{Name: "proto", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("_reduce_ex", "build/src/__python__/copy_reg.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µself *πg.Object = πArgs[0]; _ = µself
				var µproto *πg.Object = πArgs[1]; _ = µproto
				var µbase *πg.Object = πg.UnboundLocal; _ = µbase
				var µstate *πg.Object = πg.UnboundLocal; _ = µstate
				var µargs *πg.Object = πg.UnboundLocal; _ = µargs
				var µgetstate *πg.Object = πg.UnboundLocal; _ = µgetstate
				var µdict *πg.Object = πg.UnboundLocal; _ = µdict
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
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 bool
				_ = πTemp010
				var πTemp011 *πg.BaseException
				_ = πTemp011
				var πTemp012 *πg.Traceback
				_ = πTemp012
				var πTemp013 *πg.BaseException
				_ = πTemp013
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					case 18: goto Label18
					case 13: goto Label13
					default: panic("unexpected function state")
					}
					// line 60: assert proto < 2
					πF.SetLineno(60)
					if πE = πg.CheckLocal(πF, µproto, "proto"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µproto, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__mro__, nil); πE != nil {
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
						µbase = πTemp002
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(1)            
					πTemp006 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
						continue
					}
					πTemp006[0] = µbase
					πTemp006[1] = ß__flags__.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp002 = πTemp007
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, µbase, ß__flags__, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveGlobal(πF, ß_HEAPTYPE); πE != nil {
						continue
					}
					if πTemp007, πE = πg.And(πF, πTemp008, πTemp009); πE != nil {
						continue
					}
					if πTemp010, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp010).ToObject()
					πTemp002 = πTemp003
				Label4:
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label5
					}
					goto Label6
					// line 62: if hasattr(base, '__flags__') and not base.__flags__ & _HEAPTYPE:
					πF.SetLineno(62)
				Label5:
					// line 63: break
					πF.SetLineno(63)
					πTemp004 = true
					continue
					goto Label6
				Label6:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
					// line 65: base = object # not really reachable
					πF.SetLineno(65)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
						continue
					}
					µbase = πTemp002
				Label3:
					if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µbase == πTemp002).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label7
					}
					goto Label8
					// line 66: if base is object:
					πF.SetLineno(66)
				Label7:
					// line 67: state = None
					πF.SetLineno(67)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µstate = πTemp001
					goto Label9
				Label8:
					if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µbase == πTemp002).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label10
					}
					goto Label11
					// line 69: if base is self.__class__:
					πF.SetLineno(69)
				Label10:
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µbase, ß__name__, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("can't pickle %s objects").ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 70: raise TypeError, "can't pickle %s objects" % base.__name__
					πF.SetLineno(70)
					πE = πF.Raise(πTemp001, πTemp002, nil)
					continue
					goto Label11
				Label11:
					// line 71: state = base(self)
					πF.SetLineno(71)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					πTemp006[0] = µself
					if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
						continue
					}
					if πTemp001, πE = µbase.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					µstate = πTemp001
					goto Label9
				Label9:
					// line 72: args = (self.__class__, base, state)
					πF.SetLineno(72)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple3(πTemp002, µbase, µstate).ToObject()
					µargs = πTemp001
					// line 73: try:
					πF.SetLineno(73)
					πF.PushCheckpoint(13)
					// line 74: getstate = self.__getstate__
					πF.SetLineno(74)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µself, ß__getstate__, nil); πE != nil {
						continue
					}
					µgetstate = πTemp001
					πF.PopCheckpoint()
					// line 84: dict = getstate()
					πF.SetLineno(84)
					if πE = πg.CheckLocal(πF, µgetstate, "getstate"); πE != nil {
						continue
					}
					if πTemp001, πE = µgetstate.Call(πF, nil, nil); πE != nil {
						continue
					}
					µdict = πTemp001
					goto Label12
				Label13:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp011, πTemp012 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label14
					}
					πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
					continue
					// line 75: except AttributeError:
					πF.SetLineno(75)
				Label14:
					πTemp006 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					πTemp006[0] = µself
					πTemp006[1] = ß__slots__.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp006[2] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label15
					}
					goto Label16
					// line 76: if getattr(self, "__slots__", None):
					πF.SetLineno(76)
				Label15:
					πTemp006 = πF.MakeArgs(1)
					πTemp006[0] = πg.NewStr("a class that defines __slots__ without defining __getstate__ cannot be pickled").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					// line 77: raise TypeError("a class that defines __slots__ without "
					πF.SetLineno(77)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label16
				Label16:
					// line 79: try:
					πF.SetLineno(79)
					πF.PushCheckpoint(18)
					// line 80: dict = self.__dict__
					πF.SetLineno(80)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µself, ß__dict__, nil); πE != nil {
						continue
					}
					µdict = πTemp001
					πF.PopCheckpoint()
					goto Label17
				Label18:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp013, πTemp012 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp013.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label19
					}
					πE = πF.Raise(πTemp013.ToObject(), nil, πTemp012.ToObject())
					continue
					// line 81: except AttributeError:
					πF.SetLineno(81)
				Label19:
					// line 82: dict = None
					πF.SetLineno(82)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µdict = πTemp001
					πF.RestoreExc(nil, nil)
					goto Label17
				Label17:
					πF.RestoreExc(nil, nil)
					goto Label12
				Label12:
					if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µdict); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label20
					}
					goto Label21
					// line 85: if dict:
					πF.SetLineno(85)
				Label20:
					// line 86: return _reconstructor, args, dict
					πF.SetLineno(86)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_reconstructor); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple3(πTemp002, µargs, µdict).ToObject()
					πR = πTemp001
					continue
					goto Label22
				Label21:
					// line 88: return _reconstructor, args
					πF.SetLineno(88)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_reconstructor); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp002, µargs).ToObject()
					πR = πTemp001
					continue
					goto Label22
				Label22:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_reduce_ex.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 92: def __newobj__(cls, *args):
			πF.SetLineno(92)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "cls", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("__newobj__", "build/src/__python__/copy_reg.py", πTemp004, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcls *πg.Object = πArgs[0]; _ = µcls
				var µargs *πg.Object = πArgs[1]; _ = µargs
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
					// line 93: return cls.__new__(cls, *args)
					πF.SetLineno(93)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					πTemp001[0] = µcls
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µcls, ß__new__, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp001, µargs, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß__newobj__.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 95: def _slotnames(cls):
			πF.SetLineno(95)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "cls", Def: nil}
			πTemp013 = πg.NewFunction(πg.NewCode("_slotnames", "build/src/__python__/copy_reg.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcls *πg.Object = πArgs[0]; _ = µcls
				var µnames *πg.Object = πg.UnboundLocal; _ = µnames
				var µc *πg.Object = πg.UnboundLocal; _ = µc
				var µslots *πg.Object = πg.UnboundLocal; _ = µslots
				var µname *πg.Object = πg.UnboundLocal; _ = µname
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
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 bool
				_ = πTemp011
				var πTemp012 *πg.BaseException
				_ = πTemp012
				var πTemp013 *πg.Traceback
				_ = πTemp013
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 14: goto Label14
					case 22: goto Label22
					case 13: goto Label13
					case 6: goto Label6
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					// line 96: """Return a list of slot names for a given class.
					πF.SetLineno(96)
					// line 107: names = cls.__dict__.get("__slotnames__")
					πF.SetLineno(107)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ß__slotnames__.ToObject()
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µcls, ß__dict__, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µnames = πTemp002
					if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µnames != πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 108: if names is not None:
					πF.SetLineno(108)
				Label1:
					// line 109: return names
					πF.SetLineno(109)
					if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
						continue
					}
					πR = µnames
					continue
					goto Label2
				Label2:
					// line 112: names = []
					πF.SetLineno(112)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µnames = πTemp002
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					πTemp001[0] = µcls
					πTemp001[1] = ß__slots__.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 113: if not hasattr(cls, "__slots__"):
					πF.SetLineno(113)
				Label3:
					// line 115: pass
					πF.SetLineno(115)
					goto Label5
				Label4:
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µcls, ß__mro__, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
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
						πTemp006 = !isStop
					} else {
						πTemp006 = true
						µc = πTemp003
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(6)            
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µc, ß__dict__, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Contains(πF, πTemp005, ß__slots__.ToObject()); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label9
					}
					goto Label10
					// line 119: if "__slots__" in c.__dict__:
					πF.SetLineno(119)
				Label9:
					// line 120: slots = c.__dict__['__slots__']
					πF.SetLineno(120)
					πTemp003 = ß__slots__.ToObject()
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µc, ß__dict__, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
						continue
					}
					µslots = πTemp005
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µslots, "slots"); πE != nil {
						continue
					}
					πTemp001[0] = µslots
					if πTemp003, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label11
					}
					goto Label12
					// line 122: if isinstance(slots, basestring):
					πF.SetLineno(122)
				Label11:
					// line 123: slots = (slots,)
					πF.SetLineno(123)
					if πE = πg.CheckLocal(πF, µslots, "slots"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple1(µslots).ToObject()
					µslots = πTemp003
					goto Label12
				Label12:
					if πE = πg.CheckLocal(πF, µslots, "slots"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Iter(πF, µslots); πE != nil {
						continue
					}
					πF.PushCheckpoint(14)
					πTemp006 = false
				Label13:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label15
					}
					if πTemp005, πE = πg.Next(πF, πTemp003); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp008 = !isStop
					} else {
						πTemp008 = true
						µname = πTemp005
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(13)            
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp007 = πg.NewTuple2(ß__dict__.ToObject(), ß__weakref__.ToObject()).ToObject()
					if πTemp008, πE = πg.Contains(πF, πTemp007, µname); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(πTemp008).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label16
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ß__.ToObject()
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µname, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp005 = πTemp009
					if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label17
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ß__.ToObject()
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, µname, ßendswith, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp011, πE = πg.IsTrue(πF, πTemp010); πE != nil {
						continue
					}
					πTemp007 = πg.GetBool(!πTemp011).ToObject()
					πTemp005 = πTemp007
				Label17:
					if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label18
					}
					goto Label19
					// line 126: if name in ("__dict__", "__weakref__"):
					πF.SetLineno(126)
				Label16:
					// line 127: continue
					πF.SetLineno(127)
					continue
					goto Label20
					// line 129: elif name.startswith('__') and not name.endswith('__'):
					πF.SetLineno(129)
				Label18:
					// line 130: names.append('_%s%s' % (c.__name__, name))
					πF.SetLineno(130)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, µc, ß__name__, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp007 = πg.NewTuple2(πTemp009, µname).ToObject()
					if πTemp005, πE = πg.Mod(πF, πg.NewStr("_%s%s").ToObject(), πTemp007); πE != nil {
						continue
					}
					πTemp001[0] = πTemp005
					if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µnames, ßappend, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label20
				Label19:
					// line 132: names.append(name)
					πF.SetLineno(132)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001[0] = µname
					if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µnames, ßappend, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label20
				Label20:
					continue
				Label14:
					if πE != nil || πR != nil {
						continue
					}
				Label15:
					goto Label10
				Label10:
					continue
				Label7:
					if πE != nil || πR != nil {
						continue
					}
				Label8:
					goto Label5
				Label5:
					// line 135: try:
					πF.SetLineno(135)
					πF.PushCheckpoint(22)
					// line 136: cls.__slotnames__ = names
					πF.SetLineno(136)
					if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µnames); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µcls, ß__slotnames__, πTemp002); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					goto Label21
				Label22:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp012, πTemp013 = πF.ExcInfo()
					goto Label23
					// line 137: except:
					πF.SetLineno(137)
				Label23:
					// line 138: pass # But don't die if we can't
					πF.SetLineno(138)
					πF.RestoreExc(nil, nil)
					goto Label21
				Label21:
					// line 140: return names
					πF.SetLineno(140)
					if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
						continue
					}
					πR = µnames
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_slotnames.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 151: _extension_registry = {}                # key -> code
			πF.SetLineno(151)
			πTemp003 = πg.NewDict()
			πTemp014 = πTemp003.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_extension_registry.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 152: _inverted_registry = {}                 # code -> key
			πF.SetLineno(152)
			πTemp003 = πg.NewDict()
			πTemp014 = πTemp003.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_inverted_registry.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 153: _extension_cache = {}                   # code -> object
			πF.SetLineno(153)
			πTemp003 = πg.NewDict()
			πTemp014 = πTemp003.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_extension_cache.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 157: def add_extension(module, name, code):
			πF.SetLineno(157)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "module", Def: nil}
			πTemp004[1] = πg.Param{Name: "name", Def: nil}
			πTemp004[2] = πg.Param{Name: "code", Def: nil}
			πTemp014 = πg.NewFunction(πg.NewCode("add_extension", "build/src/__python__/copy_reg.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmodule *πg.Object = πArgs[0]; _ = µmodule
				var µname *πg.Object = πArgs[1]; _ = µname
				var µcode *πg.Object = πArgs[2]; _ = µcode
				var µkey *πg.Object = πg.UnboundLocal; _ = µkey
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
				var πTemp006 *πg.Object
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
					// line 158: """Register an extension code."""
					πF.SetLineno(158)
					// line 159: code = int(code)
					πF.SetLineno(159)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp001[0] = µcode
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcode = πTemp003
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LE(πF, πg.NewInt(1).ToObject(), µcode); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label1
					}
					if πTemp003, πE = πg.LE(πF, µcode, πg.NewInt(2147483647).ToObject()); πE != nil {
						continue
					}
				Label1:
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label2
					}
					goto Label3
					// line 160: if not 1 <= code <= 0x7fffffff:
					πF.SetLineno(160)
				Label2:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					// line 161: raise ValueError, "code out of range"
					πF.SetLineno(161)
					πE = πF.Raise(πTemp002, πg.NewStr("code out of range").ToObject(), nil)
					continue
					goto Label3
				Label3:
					// line 162: key = (module, name)
					πF.SetLineno(162)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µmodule, µname).ToObject()
					µkey = πTemp002
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp001[0] = µkey
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_extension_registry); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßget, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp005, µcode); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label4
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp001[0] = µcode
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_inverted_registry); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßget, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp005, µkey); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label4:
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label5
					}
					goto Label6
					// line 163: if (_extension_registry.get(key) == code and
					πF.SetLineno(163)
				Label5:
					// line 165: return # Redundant registrations are benign
					πF.SetLineno(165)
					πR = πg.None
					continue
					goto Label6
				Label6:
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_extension_registry); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, πTemp003, µkey); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label7
					}
					goto Label8
					// line 166: if key in _extension_registry:
					πF.SetLineno(166)
				Label7:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp005 = µkey
					if πTemp007, πE = πg.ResolveGlobal(πF, ß_extension_registry); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µkey, πTemp006).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("key %s is already registered with code %s").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 167: raise ValueError("key %s is already registered with code %s" %
					πF.SetLineno(167)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label8
				Label8:
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_inverted_registry); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, πTemp003, µcode); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label9
					}
					goto Label10
					// line 169: if code in _inverted_registry:
					πF.SetLineno(169)
				Label9:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp005 = µcode
					if πTemp007, πE = πg.ResolveGlobal(πF, ß_inverted_registry); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µcode, πTemp006).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("code %s is already in use for key %s").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 170: raise ValueError("code %s is already in use for key %s" %
					πF.SetLineno(170)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label10
				Label10:
					// line 172: _extension_registry[key] = code
					πF.SetLineno(172)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µcode); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_extension_registry); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp005 = µkey
					if πE = πg.SetItem(πF, πTemp003, πTemp005, πTemp002); πE != nil {
						continue
					}
					// line 173: _inverted_registry[code] = key
					πF.SetLineno(173)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µkey); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_inverted_registry); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp005 = µcode
					if πE = πg.SetItem(πF, πTemp003, πTemp005, πTemp002); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßadd_extension.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 175: def remove_extension(module, name, code):
			πF.SetLineno(175)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "module", Def: nil}
			πTemp004[1] = πg.Param{Name: "name", Def: nil}
			πTemp004[2] = πg.Param{Name: "code", Def: nil}
			πTemp015 = πg.NewFunction(πg.NewCode("remove_extension", "build/src/__python__/copy_reg.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmodule *πg.Object = πArgs[0]; _ = µmodule
				var µname *πg.Object = πArgs[1]; _ = µname
				var µcode *πg.Object = πArgs[2]; _ = µcode
				var µkey *πg.Object = πg.UnboundLocal; _ = µkey
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 176: """Unregister an extension code.  For testing only."""
					πF.SetLineno(176)
					// line 177: key = (module, name)
					πF.SetLineno(177)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µmodule, µname).ToObject()
					µkey = πTemp001
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp004[0] = µkey
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_extension_registry); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßget, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.NE(πF, πTemp005, µcode); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp004[0] = µcode
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_inverted_registry); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßget, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.NE(πF, πTemp005, µkey); πE != nil {
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
					// line 178: if (_extension_registry.get(key) != code or
					πF.SetLineno(178)
				Label2:
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µkey, µcode).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("key %s is not registered with code %s").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 180: raise ValueError("key %s is not registered with code %s" %
					πF.SetLineno(180)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label3
				Label3:
					// line 182: del _extension_registry[key]
					πF.SetLineno(182)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_extension_registry); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp003 = µkey
					if πE = πg.DelItem(πF, πTemp001, πTemp003); πE != nil {
						continue
					}
					// line 183: del _inverted_registry[code]
					πF.SetLineno(183)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_inverted_registry); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp003 = µcode
					if πE = πg.DelItem(πF, πTemp001, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_extension_cache); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, πTemp003, µcode); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					goto Label5
					// line 184: if code in _extension_cache:
					πF.SetLineno(184)
				Label4:
					// line 185: del _extension_cache[code]
					πF.SetLineno(185)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_extension_cache); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp003 = µcode
					if πE = πg.DelItem(πF, πTemp001, πTemp003); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßremove_extension.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 187: def clear_extension_cache():
			πF.SetLineno(187)
			πTemp004 = make([]πg.Param, 0)
			πTemp016 = πg.NewFunction(πg.NewCode("clear_extension_cache", "build/src/__python__/copy_reg.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 188: _extension_cache.clear()
					πF.SetLineno(188)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_extension_cache); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßclear, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßclear_extension_cache.ToObject(), πTemp016); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("copy_reg", Code)
}
