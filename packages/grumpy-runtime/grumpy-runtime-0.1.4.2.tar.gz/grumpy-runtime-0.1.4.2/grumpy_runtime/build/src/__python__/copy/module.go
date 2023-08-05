package copy
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/copy.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßA := πg.InternStr("A")
		ßABC := πg.InternStr("ABC")
		ßAttributeError := πg.InternStr("AttributeError")
		ßB := πg.InternStr("B")
		ßBuiltinFunctionType := πg.InternStr("BuiltinFunctionType")
		ßCodeType := πg.InternStr("CodeType")
		ßComplexType := πg.InternStr("ComplexType")
		ßError := πg.InternStr("Error")
		ßException := πg.InternStr("Exception")
		ßFunctionType := πg.InternStr("FunctionType")
		ßKeyError := πg.InternStr("KeyError")
		ßMethodType := πg.InternStr("MethodType")
		ßNameError := πg.InternStr("NameError")
		ßNone := πg.InternStr("None")
		ßPyStringMap := πg.InternStr("PyStringMap")
		ßTypeError := πg.InternStr("TypeError")
		ßUnicodeType := πg.InternStr("UnicodeType")
		ßWeakRefType := πg.InternStr("WeakRefType")
		ß__all__ := πg.InternStr("__all__")
		ß__class__ := πg.InternStr("__class__")
		ß__copy__ := πg.InternStr("__copy__")
		ß__deepcopy__ := πg.InternStr("__deepcopy__")
		ß__dict__ := πg.InternStr("__dict__")
		ß__file__ := πg.InternStr("__file__")
		ß__getstate__ := πg.InternStr("__getstate__")
		ß__init__ := πg.InternStr("__init__")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__reduce__ := πg.InternStr("__reduce__")
		ß__reduce_ex__ := πg.InternStr("__reduce_ex__")
		ß__setitem__ := πg.InternStr("__setitem__")
		ß__setstate__ := πg.InternStr("__setstate__")
		ß_copy_dispatch := πg.InternStr("_copy_dispatch")
		ß_copy_immutable := πg.InternStr("_copy_immutable")
		ß_copy_with_constructor := πg.InternStr("_copy_with_constructor")
		ß_copy_with_copy_method := πg.InternStr("_copy_with_copy_method")
		ß_deepcopy_atomic := πg.InternStr("_deepcopy_atomic")
		ß_deepcopy_dict := πg.InternStr("_deepcopy_dict")
		ß_deepcopy_dispatch := πg.InternStr("_deepcopy_dispatch")
		ß_deepcopy_list := πg.InternStr("_deepcopy_list")
		ß_deepcopy_method := πg.InternStr("_deepcopy_method")
		ß_deepcopy_tuple := πg.InternStr("_deepcopy_tuple")
		ß_keep_alive := πg.InternStr("_keep_alive")
		ß_reconstruct := πg.InternStr("_reconstruct")
		ß_test := πg.InternStr("_test")
		ßa := πg.InternStr("a")
		ßabc := πg.InternStr("abc")
		ßappend := πg.InternStr("append")
		ßarg := πg.InternStr("arg")
		ßargv := πg.InternStr("argv")
		ßbool := πg.InternStr("bool")
		ßclose := πg.InternStr("close")
		ßcomplex := πg.InternStr("complex")
		ßcopy := πg.InternStr("copy")
		ßcopy_reg := πg.InternStr("copy_reg")
		ßd := πg.InternStr("d")
		ßdeepcopy := πg.InternStr("deepcopy")
		ßdict := πg.InternStr("dict")
		ßdispatch_table := πg.InternStr("dispatch_table")
		ßerror := πg.InternStr("error")
		ßfloat := πg.InternStr("float")
		ßfp := πg.InternStr("fp")
		ßfrozenset := πg.InternStr("frozenset")
		ßget := πg.InternStr("get")
		ßgetattr := πg.InternStr("getattr")
		ßhasattr := πg.InternStr("hasattr")
		ßid := πg.InternStr("id")
		ßim_class := πg.InternStr("im_class")
		ßim_func := πg.InternStr("im_func")
		ßim_self := πg.InternStr("im_self")
		ßint := πg.InternStr("int")
		ßisinstance := πg.InternStr("isinstance")
		ßissubclass := πg.InternStr("issubclass")
		ßiteritems := πg.InternStr("iteritems")
		ßlen := πg.InternStr("len")
		ßlist := πg.InternStr("list")
		ßlong := πg.InternStr("long")
		ßmap := πg.InternStr("map")
		ßname := πg.InternStr("name")
		ßopen := πg.InternStr("open")
		ßrange := πg.InternStr("range")
		ßrepr := πg.InternStr("repr")
		ßsetattr := πg.InternStr("setattr")
		ßstr := πg.InternStr("str")
		ßt := πg.InternStr("t")
		ßtuple := πg.InternStr("tuple")
		ßtype := πg.InternStr("type")
		ßtypes := πg.InternStr("types")
		ßunicode := πg.InternStr("unicode")
		ßupdate := πg.InternStr("update")
		ßxrange := πg.InternStr("xrange")
		ßxyz := πg.InternStr("xyz")
		ßxyzzy := πg.InternStr("xyzzy")
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
		var πTemp008 []*πg.Object
		_ = πTemp008
		var πTemp009 *πg.Object
		_ = πTemp009
		var πTemp010 *πg.Object
		_ = πTemp010
		var πTemp011 bool
		_ = πTemp011
		var πTemp012 bool
		_ = πTemp012
		var πTemp013 *πg.Object
		_ = πTemp013
		var πTemp014 *πg.Object
		_ = πTemp014
		var πTemp015 *πg.Object
		_ = πTemp015
		var πTemp016 *πg.Object
		_ = πTemp016
		var πTemp017 *πg.Object
		_ = πTemp017
		var πTemp018 *πg.Object
		_ = πTemp018
		var πTemp019 *πg.BaseException
		_ = πTemp019
		var πTemp020 *πg.Traceback
		_ = πTemp020
		var πTemp021 *πg.Object
		_ = πTemp021
		var πTemp022 *πg.Object
		_ = πTemp022
		var πTemp023 *πg.Object
		_ = πTemp023
		var πTemp024 *πg.Object
		_ = πTemp024
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 1: goto Label1
			case 2: goto Label2
			case 4: goto Label4
			case 5: goto Label5
			case 9: goto Label9
			case 10: goto Label10
			case 15: goto Label15
			case 18: goto Label18
			case 21: goto Label21
			default: panic("unexpected function state")
			}
			// line 1: """Generic (shallow and deep) copying operations.
			πF.SetLineno(1)
			// line 51: from '__go__/grumpy' import WeakRefType
			πF.SetLineno(51)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/grumpy"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßWeakRefType, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßWeakRefType.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 52: import types
			πF.SetLineno(52)
			if πTemp002, πE = πg.ImportModule(πF, "types"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtypes.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 54: import copy_reg
			πF.SetLineno(54)
			if πTemp002, πE = πg.ImportModule(πF, "copy_reg"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßcopy_reg.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 55: dispatch_table = copy_reg.dispatch_table
			πF.SetLineno(55)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßcopy_reg); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßdispatch_table, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdispatch_table.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 58: class Error(Exception):
			πF.SetLineno(58)
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
			_, πE = πg.NewCode("Error", "build/src/__python__/copy.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 59: pass
					πF.SetLineno(59)
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Error").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßError.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 60: error = Error   # backward compatibility
			πF.SetLineno(60)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßerror.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 65: PyStringMap = None
			πF.SetLineno(65)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPyStringMap.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 67: __all__ = ["Error", "copy", "deepcopy"]
			πF.SetLineno(67)
			πTemp002 = make([]*πg.Object, 3)
			πTemp002[0] = ßError.ToObject()
			πTemp002[1] = ßcopy.ToObject()
			πTemp002[2] = ßdeepcopy.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 69: def copy(x):
			πF.SetLineno(69)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "x", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("copy", "build/src/__python__/copy.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µcls *πg.Object = πg.UnboundLocal; _ = µcls
				var µcopier *πg.Object = πg.UnboundLocal; _ = µcopier
				var µreductor *πg.Object = πg.UnboundLocal; _ = µreductor
				var µrv *πg.Object = πg.UnboundLocal; _ = µrv
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 70: """Shallow copy operation on arbitrary Python objects.
					πF.SetLineno(70)
					// line 75: cls = type(x)
					πF.SetLineno(75)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcls = πTemp003
					// line 77: copier = _copy_dispatch.get(cls)
					πF.SetLineno(77)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					πTemp001[0] = µcls
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_copy_dispatch); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcopier = πTemp002
					if πE = πg.CheckLocal(πF, µcopier, "copier"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µcopier); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 78: if copier:
					πF.SetLineno(78)
				Label1:
					// line 79: return copier(x)
					πF.SetLineno(79)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					if πE = πg.CheckLocal(πF, µcopier, "copier"); πE != nil {
						continue
					}
					if πTemp002, πE = µcopier.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp002
					continue
					goto Label2
				Label2:
					// line 81: copier = getattr(cls, "__copy__", None)
					πF.SetLineno(81)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					πTemp001[0] = µcls
					πTemp001[1] = ß__copy__.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcopier = πTemp003
					if πE = πg.CheckLocal(πF, µcopier, "copier"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µcopier); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 82: if copier:
					πF.SetLineno(82)
				Label3:
					// line 83: return copier(x)
					πF.SetLineno(83)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					if πE = πg.CheckLocal(πF, µcopier, "copier"); πE != nil {
						continue
					}
					if πTemp002, πE = µcopier.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp002
					continue
					goto Label4
				Label4:
					// line 85: reductor = dispatch_table.get(cls)
					πF.SetLineno(85)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					πTemp001[0] = µcls
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdispatch_table); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µreductor = πTemp002
					if πE = πg.CheckLocal(πF, µreductor, "reductor"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µreductor); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label5
					}
					goto Label6
					// line 86: if reductor:
					πF.SetLineno(86)
				Label5:
					// line 87: rv = reductor(x)
					πF.SetLineno(87)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					if πE = πg.CheckLocal(πF, µreductor, "reductor"); πE != nil {
						continue
					}
					if πTemp002, πE = µreductor.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µrv = πTemp002
					goto Label7
				Label6:
					// line 89: reductor = getattr(x, "__reduce_ex__", None)
					πF.SetLineno(89)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					πTemp001[1] = ß__reduce_ex__.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µreductor = πTemp003
					if πE = πg.CheckLocal(πF, µreductor, "reductor"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µreductor); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label8
					}
					goto Label9
					// line 90: if reductor:
					πF.SetLineno(90)
				Label8:
					// line 91: rv = reductor(2)
					πF.SetLineno(91)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µreductor, "reductor"); πE != nil {
						continue
					}
					if πTemp002, πE = µreductor.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µrv = πTemp002
					goto Label10
				Label9:
					// line 93: reductor = getattr(x, "__reduce__", None)
					πF.SetLineno(93)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					πTemp001[1] = ß__reduce__.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µreductor = πTemp003
					if πE = πg.CheckLocal(πF, µreductor, "reductor"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µreductor); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label11
					}
					goto Label12
					// line 94: if reductor:
					πF.SetLineno(94)
				Label11:
					// line 95: rv = reductor()
					πF.SetLineno(95)
					if πE = πg.CheckLocal(πF, µreductor, "reductor"); πE != nil {
						continue
					}
					if πTemp002, πE = µreductor.Call(πF, nil, nil); πE != nil {
						continue
					}
					µrv = πTemp002
					goto Label13
				Label12:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("un(shallow)copyable object of type %s").ToObject(), µcls); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 97: raise Error("un(shallow)copyable object of type %s" % cls)
					πF.SetLineno(97)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label13
				Label13:
					goto Label10
				Label10:
					goto Label7
				Label7:
					// line 99: return _reconstruct(x, rv, 0)
					πF.SetLineno(99)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					if πE = πg.CheckLocal(πF, µrv, "rv"); πE != nil {
						continue
					}
					πTemp001[1] = µrv
					πTemp001[2] = πg.NewInt(0).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_reconstruct); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßcopy.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 102: _copy_dispatch = d = {}
			πF.SetLineno(102)
			πTemp004 = πg.NewDict()
			πTemp003 = πTemp004.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_copy_dispatch.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßd.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 104: def _copy_immutable(x):
			πF.SetLineno(104)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "x", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("_copy_immutable", "build/src/__python__/copy.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 105: return x
					πF.SetLineno(105)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πR = µx
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_copy_immutable.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp002 = make([]*πg.Object, 13)
			πTemp008 = πF.MakeArgs(1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp008[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp010, πE = πTemp009.Call(πF, πTemp008, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp008)
			πTemp002[0] = πTemp010
			if πTemp009, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
				continue
			}
			πTemp002[2] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßbool); πE != nil {
				continue
			}
			πTemp002[3] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
				continue
			}
			πTemp002[4] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
				continue
			}
			πTemp002[5] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
				continue
			}
			πTemp002[6] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			πTemp002[7] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
				continue
			}
			πTemp002[8] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßfrozenset); πE != nil {
				continue
			}
			πTemp002[9] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
				continue
			}
			if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßBuiltinFunctionType, nil); πE != nil {
				continue
			}
			πTemp002[10] = πTemp010
			if πTemp009, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
				continue
			}
			if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßFunctionType, nil); πE != nil {
				continue
			}
			πTemp002[11] = πTemp010
			if πTemp009, πE = πg.ResolveGlobal(πF, ßWeakRefType); πE != nil {
				continue
			}
			πTemp002[12] = πTemp009
			πTemp007 = πg.NewTuple(πTemp002...).ToObject()
			if πTemp005, πE = πg.Iter(πF, πTemp007); πE != nil {
				continue
			}
			πF.PushCheckpoint(2)
			πTemp011 = false
		Label1:
			if πE != nil || πR != nil {
				continue
			}
			if πTemp011 {
				πF.PopCheckpoint()
				goto Label3
			}
			if πTemp007, πE = πg.Next(πF, πTemp005); πE != nil {
				isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
				if exc != nil {
					πE = exc
				} else if isStop {
					πE = nil
					πF.RestoreExc(nil, nil)
				}
				πTemp012 = !isStop
			} else {
				πTemp012 = true
				if πE = πF.Globals().SetItem(πF, ßt.ToObject(), πTemp007); πE != nil {
					continue
				}
			}
			if πE != nil || !πTemp012 {
				continue
			}
			πF.PushCheckpoint(1)            
			// line 110: d[t] = _copy_immutable
			πF.SetLineno(110)
			if πTemp007, πE = πg.ResolveGlobal(πF, ß_copy_immutable); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πTemp007); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			if πTemp014, πE = πg.ResolveGlobal(πF, ßt); πE != nil {
				continue
			}
			πTemp013 = πTemp014
			if πE = πg.SetItem(πF, πTemp010, πTemp013, πTemp009); πE != nil {
				continue
			}
			continue
		Label2:
			if πE != nil || πR != nil {
				continue
			}
		Label3:
			πTemp007 = πg.NewTuple3(ßComplexType.ToObject(), ßUnicodeType.ToObject(), ßCodeType.ToObject()).ToObject()
			if πTemp005, πE = πg.Iter(πF, πTemp007); πE != nil {
				continue
			}
			πF.PushCheckpoint(5)
			πTemp011 = false
		Label4:
			if πE != nil || πR != nil {
				continue
			}
			if πTemp011 {
				πF.PopCheckpoint()
				goto Label6
			}
			if πTemp007, πE = πg.Next(πF, πTemp005); πE != nil {
				isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
				if exc != nil {
					πE = exc
				} else if isStop {
					πE = nil
					πF.RestoreExc(nil, nil)
				}
				πTemp012 = !isStop
			} else {
				πTemp012 = true
				if πE = πF.Globals().SetItem(πF, ßname.ToObject(), πTemp007); πE != nil {
					continue
				}
			}
			if πE != nil || !πTemp012 {
				continue
			}
			πF.PushCheckpoint(4)            
			// line 112: t = getattr(types, name, None)
			πF.SetLineno(112)
			πTemp002 = πF.MakeArgs(3)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßname); πE != nil {
				continue
			}
			πTemp002[1] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp002[2] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
				continue
			}
			if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßt.ToObject(), πTemp009); πE != nil {
				continue
			}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßt); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp007 = πg.GetBool(πTemp009 != πTemp010).ToObject()
			if πTemp012, πE = πg.IsTrue(πF, πTemp007); πE != nil {
				continue
			}
			if πTemp012 {
				goto Label7
			}
			goto Label8
			// line 113: if t is not None:
			πF.SetLineno(113)
		Label7:
			// line 114: d[t] = _copy_immutable
			πF.SetLineno(114)
			if πTemp007, πE = πg.ResolveGlobal(πF, ß_copy_immutable); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πTemp007); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			if πTemp014, πE = πg.ResolveGlobal(πF, ßt); πE != nil {
				continue
			}
			πTemp013 = πTemp014
			if πE = πg.SetItem(πF, πTemp010, πTemp013, πTemp009); πE != nil {
				continue
			}
			goto Label8
		Label8:
			continue
		Label5:
			if πE != nil || πR != nil {
				continue
			}
		Label6:
			// line 116: def _copy_with_constructor(x):
			πF.SetLineno(116)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "x", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("_copy_with_constructor", "build/src/__python__/copy.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var πTemp001 []*πg.Object
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
					// line 117: return type(x)(x)
					πF.SetLineno(117)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_copy_with_constructor.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
				continue
			}
			πTemp009 = πg.NewTuple2(πTemp010, πTemp013).ToObject()
			if πTemp007, πE = πg.Iter(πF, πTemp009); πE != nil {
				continue
			}
			πF.PushCheckpoint(10)
			πTemp011 = false
		Label9:
			if πE != nil || πR != nil {
				continue
			}
			if πTemp011 {
				πF.PopCheckpoint()
				goto Label11
			}
			if πTemp009, πE = πg.Next(πF, πTemp007); πE != nil {
				isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
				if exc != nil {
					πE = exc
				} else if isStop {
					πE = nil
					πF.RestoreExc(nil, nil)
				}
				πTemp012 = !isStop
			} else {
				πTemp012 = true
				if πE = πF.Globals().SetItem(πF, ßt.ToObject(), πTemp009); πE != nil {
					continue
				}
			}
			if πE != nil || !πTemp012 {
				continue
			}
			πF.PushCheckpoint(9)            
			// line 119: d[t] = _copy_with_constructor
			πF.SetLineno(119)
			if πTemp009, πE = πg.ResolveGlobal(πF, ß_copy_with_constructor); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πTemp009); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßt); πE != nil {
				continue
			}
			πTemp014 = πTemp015
			if πE = πg.SetItem(πF, πTemp013, πTemp014, πTemp010); πE != nil {
				continue
			}
			continue
		Label10:
			if πE != nil || πR != nil {
				continue
			}
		Label11:
			// line 121: def _copy_with_copy_method(x):
			πF.SetLineno(121)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "x", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("_copy_with_copy_method", "build/src/__python__/copy.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 122: return x.copy()
					πF.SetLineno(122)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µx, ßcopy, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_copy_with_copy_method.ToObject(), πTemp007); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßPyStringMap); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp009 = πg.GetBool(πTemp010 != πTemp013).ToObject()
			if πTemp011, πE = πg.IsTrue(πF, πTemp009); πE != nil {
				continue
			}
			if πTemp011 {
				goto Label12
			}
			goto Label13
			// line 123: if PyStringMap is not None:
			πF.SetLineno(123)
		Label12:
			// line 124: d[PyStringMap] = _copy_with_copy_method
			πF.SetLineno(124)
			if πTemp009, πE = πg.ResolveGlobal(πF, ß_copy_with_copy_method); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πTemp009); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßPyStringMap); πE != nil {
				continue
			}
			πTemp014 = πTemp015
			if πE = πg.SetItem(πF, πTemp013, πTemp014, πTemp010); πE != nil {
				continue
			}
			goto Label13
		Label13:
			// line 146: del d
			πF.SetLineno(146)
			if πE = πg.DelVar(πF, πF.Globals(), ßd); πE != nil {
				continue
			}
			// line 148: def deepcopy(x, memo=None, _nil=[]):
			πF.SetLineno(148)
			πTemp006 = make([]πg.Param, 3)
			πTemp006[0] = πg.Param{Name: "x", Def: nil}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "memo", Def: πTemp010}
			πTemp002 = make([]*πg.Object, 0)
			πTemp010 = πg.NewList(πTemp002...).ToObject()
			πTemp006[2] = πg.Param{Name: "_nil", Def: πTemp010}
			πTemp009 = πg.NewFunction(πg.NewCode("deepcopy", "build/src/__python__/copy.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µmemo *πg.Object = πArgs[1]; _ = µmemo
				var µ_nil *πg.Object = πArgs[2]; _ = µ_nil
				var µd *πg.Object = πg.UnboundLocal; _ = µd
				var µy *πg.Object = πg.UnboundLocal; _ = µy
				var µcls *πg.Object = πg.UnboundLocal; _ = µcls
				var µcopier *πg.Object = πg.UnboundLocal; _ = µcopier
				var µissc *πg.Object = πg.UnboundLocal; _ = µissc
				var µreductor *πg.Object = πg.UnboundLocal; _ = µreductor
				var µrv *πg.Object = πg.UnboundLocal; _ = µrv
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
				var πTemp006 *πg.BaseException
				_ = πTemp006
				var πTemp007 *πg.Traceback
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 9: goto Label9
					default: panic("unexpected function state")
					}
					// line 149: """Deep copy operation on arbitrary Python objects.
					πF.SetLineno(149)
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µmemo == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 154: if memo is None:
					πF.SetLineno(154)
				Label1:
					// line 155: memo = {}
					πF.SetLineno(155)
					πTemp004 = πg.NewDict()
					πTemp001 = πTemp004.ToObject()
					µmemo = πTemp001
					goto Label2
				Label2:
					// line 157: d = id(x)
					πF.SetLineno(157)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp005[0] = µx
					if πTemp001, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µd = πTemp002
					// line 158: y = memo.get(d, _nil)
					πF.SetLineno(158)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					πTemp005[0] = µd
					if πE = πg.CheckLocal(πF, µ_nil, "_nil"); πE != nil {
						continue
					}
					πTemp005[1] = µ_nil
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µmemo, ßget, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µy = πTemp002
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µ_nil, "_nil"); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µy != µ_nil).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label3
					}
					goto Label4
					// line 159: if y is not _nil:
					πF.SetLineno(159)
				Label3:
					// line 160: return y
					πF.SetLineno(160)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					πR = µy
					continue
					goto Label4
				Label4:
					// line 162: cls = type(x)
					πF.SetLineno(162)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp005[0] = µx
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µcls = πTemp002
					// line 164: copier = _deepcopy_dispatch.get(cls)
					πF.SetLineno(164)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					πTemp005[0] = µcls
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_deepcopy_dispatch); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µcopier = πTemp001
					if πE = πg.CheckLocal(πF, µcopier, "copier"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µcopier); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label5
					}
					goto Label6
					// line 165: if copier:
					πF.SetLineno(165)
				Label5:
					// line 166: y = copier(x, memo)
					πF.SetLineno(166)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp005[0] = µx
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					πTemp005[1] = µmemo
					if πE = πg.CheckLocal(πF, µcopier, "copier"); πE != nil {
						continue
					}
					if πTemp001, πE = µcopier.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µy = πTemp001
					goto Label7
				Label6:
					// line 168: try:
					πF.SetLineno(168)
					πF.PushCheckpoint(9)
					// line 169: issc = issubclass(cls, type)
					πF.SetLineno(169)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					πTemp005[0] = µcls
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					πTemp005[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µissc = πTemp002
					πF.PopCheckpoint()
					goto Label8
				Label9:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp006, πTemp007 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label10
					}
					πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
					continue
					// line 170: except TypeError: # cls is not a class (old Boost; see SF #502085)
					πF.SetLineno(170)
				Label10:
					// line 171: issc = 0
					πF.SetLineno(171)
					µissc = πg.NewInt(0).ToObject()
					πF.RestoreExc(nil, nil)
					goto Label8
				Label8:
					if πE = πg.CheckLocal(πF, µissc, "issc"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µissc); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label11
					}
					goto Label12
					// line 172: if issc:
					πF.SetLineno(172)
				Label11:
					// line 173: y = _deepcopy_atomic(x, memo)
					πF.SetLineno(173)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp005[0] = µx
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					πTemp005[1] = µmemo
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_deepcopy_atomic); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µy = πTemp002
					goto Label13
				Label12:
					// line 175: copier = getattr(x, "__deepcopy__", None)
					πF.SetLineno(175)
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp005[0] = µx
					πTemp005[1] = ß__deepcopy__.ToObject()
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
					µcopier = πTemp002
					if πE = πg.CheckLocal(πF, µcopier, "copier"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µcopier); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label14
					}
					goto Label15
					// line 176: if copier:
					πF.SetLineno(176)
				Label14:
					// line 177: y = copier(memo)
					πF.SetLineno(177)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					πTemp005[0] = µmemo
					if πE = πg.CheckLocal(πF, µcopier, "copier"); πE != nil {
						continue
					}
					if πTemp001, πE = µcopier.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µy = πTemp001
					goto Label16
				Label15:
					// line 179: reductor = dispatch_table.get(cls)
					πF.SetLineno(179)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					πTemp005[0] = µcls
					if πTemp001, πE = πg.ResolveGlobal(πF, ßdispatch_table); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µreductor = πTemp001
					if πE = πg.CheckLocal(πF, µreductor, "reductor"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µreductor); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label17
					}
					goto Label18
					// line 180: if reductor:
					πF.SetLineno(180)
				Label17:
					// line 181: rv = reductor(x)
					πF.SetLineno(181)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp005[0] = µx
					if πE = πg.CheckLocal(πF, µreductor, "reductor"); πE != nil {
						continue
					}
					if πTemp001, πE = µreductor.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µrv = πTemp001
					goto Label19
				Label18:
					// line 183: reductor = getattr(x, "__reduce_ex__", None)
					πF.SetLineno(183)
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp005[0] = µx
					πTemp005[1] = ß__reduce_ex__.ToObject()
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
					µreductor = πTemp002
					if πE = πg.CheckLocal(πF, µreductor, "reductor"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µreductor); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label20
					}
					goto Label21
					// line 184: if reductor:
					πF.SetLineno(184)
				Label20:
					// line 185: rv = reductor(2)
					πF.SetLineno(185)
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µreductor, "reductor"); πE != nil {
						continue
					}
					if πTemp001, πE = µreductor.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µrv = πTemp001
					goto Label22
				Label21:
					// line 187: reductor = getattr(x, "__reduce__", None)
					πF.SetLineno(187)
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp005[0] = µx
					πTemp005[1] = ß__reduce__.ToObject()
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
					µreductor = πTemp002
					if πE = πg.CheckLocal(πF, µreductor, "reductor"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µreductor); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label23
					}
					goto Label24
					// line 188: if reductor:
					πF.SetLineno(188)
				Label23:
					// line 189: rv = reductor()
					πF.SetLineno(189)
					if πE = πg.CheckLocal(πF, µreductor, "reductor"); πE != nil {
						continue
					}
					if πTemp001, πE = µreductor.Call(πF, nil, nil); πE != nil {
						continue
					}
					µrv = πTemp001
					goto Label25
				Label24:
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("un(deep)copyable object of type %s").ToObject(), µcls); πE != nil {
						continue
					}
					πTemp005[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 191: raise Error(
					πF.SetLineno(191)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label25
				Label25:
					goto Label22
				Label22:
					goto Label19
				Label19:
					// line 193: y = _reconstruct(x, rv, 1, memo)
					πF.SetLineno(193)
					πTemp005 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp005[0] = µx
					if πE = πg.CheckLocal(πF, µrv, "rv"); πE != nil {
						continue
					}
					πTemp005[1] = µrv
					πTemp005[2] = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					πTemp005[3] = µmemo
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_reconstruct); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µy = πTemp002
					goto Label16
				Label16:
					goto Label13
				Label13:
					goto Label7
				Label7:
					// line 195: memo[d] = y
					πF.SetLineno(195)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µy); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					πTemp002 = µd
					if πE = πg.SetItem(πF, µmemo, πTemp002, πTemp001); πE != nil {
						continue
					}
					// line 196: _keep_alive(x, memo) # Make sure x lives at least as long as d
					πF.SetLineno(196)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp005[0] = µx
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					πTemp005[1] = µmemo
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_keep_alive); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 197: return y
					πF.SetLineno(197)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					πR = µy
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßdeepcopy.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 199: _deepcopy_dispatch = d = {}
			πF.SetLineno(199)
			πTemp004 = πg.NewDict()
			πTemp010 = πTemp004.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_deepcopy_dispatch.ToObject(), πTemp010); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßd.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 201: def _deepcopy_atomic(x, memo):
			πF.SetLineno(201)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "x", Def: nil}
			πTemp006[1] = πg.Param{Name: "memo", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("_deepcopy_atomic", "build/src/__python__/copy.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µmemo *πg.Object = πArgs[1]; _ = µmemo
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 202: return x
					πF.SetLineno(202)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πR = µx
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_deepcopy_atomic.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 203: d[type(None)] = _deepcopy_atomic
			πF.SetLineno(203)
			if πTemp013, πE = πg.ResolveGlobal(πF, ß_deepcopy_atomic); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πTemp013); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			πTemp002 = πF.MakeArgs(1)
			if πTemp017, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp002[0] = πTemp017
			if πTemp017, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp018, πE = πTemp017.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			πTemp016 = πTemp018
			if πE = πg.SetItem(πF, πTemp015, πTemp016, πTemp014); πE != nil {
				continue
			}
			// line 205: d[int] = _deepcopy_atomic
			πF.SetLineno(205)
			if πTemp013, πE = πg.ResolveGlobal(πF, ß_deepcopy_atomic); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πTemp013); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
				continue
			}
			πTemp016 = πTemp017
			if πE = πg.SetItem(πF, πTemp015, πTemp016, πTemp014); πE != nil {
				continue
			}
			// line 206: d[long] = _deepcopy_atomic
			πF.SetLineno(206)
			if πTemp013, πE = πg.ResolveGlobal(πF, ß_deepcopy_atomic); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πTemp013); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
				continue
			}
			πTemp016 = πTemp017
			if πE = πg.SetItem(πF, πTemp015, πTemp016, πTemp014); πE != nil {
				continue
			}
			// line 207: d[float] = _deepcopy_atomic
			πF.SetLineno(207)
			if πTemp013, πE = πg.ResolveGlobal(πF, ß_deepcopy_atomic); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πTemp013); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
				continue
			}
			πTemp016 = πTemp017
			if πE = πg.SetItem(πF, πTemp015, πTemp016, πTemp014); πE != nil {
				continue
			}
			// line 208: d[bool] = _deepcopy_atomic
			πF.SetLineno(208)
			if πTemp013, πE = πg.ResolveGlobal(πF, ß_deepcopy_atomic); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πTemp013); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßbool); πE != nil {
				continue
			}
			πTemp016 = πTemp017
			if πE = πg.SetItem(πF, πTemp015, πTemp016, πTemp014); πE != nil {
				continue
			}
			// line 209: try:
			πF.SetLineno(209)
			πF.PushCheckpoint(15)
			// line 210: d[complex] = _deepcopy_atomic
			πF.SetLineno(210)
			if πTemp013, πE = πg.ResolveGlobal(πF, ß_deepcopy_atomic); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πTemp013); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßcomplex); πE != nil {
				continue
			}
			πTemp016 = πTemp017
			if πE = πg.SetItem(πF, πTemp015, πTemp016, πTemp014); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label14
		Label15:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp019, πTemp020 = πF.ExcInfo()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßNameError); πE != nil {
				continue
			}
			if πTemp011, πE = πg.IsInstance(πF, πTemp019.ToObject(), πTemp013); πE != nil {
				continue
			}
			if πTemp011 {
				goto Label16
			}
			πE = πF.Raise(πTemp019.ToObject(), nil, πTemp020.ToObject())
			continue
			// line 211: except NameError:
			πF.SetLineno(211)
		Label16:
			// line 212: pass
			πF.SetLineno(212)
			πF.RestoreExc(nil, nil)
			goto Label14
		Label14:
			// line 213: d[str] = _deepcopy_atomic
			πF.SetLineno(213)
			if πTemp013, πE = πg.ResolveGlobal(πF, ß_deepcopy_atomic); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πTemp013); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
				continue
			}
			πTemp016 = πTemp017
			if πE = πg.SetItem(πF, πTemp015, πTemp016, πTemp014); πE != nil {
				continue
			}
			// line 214: try:
			πF.SetLineno(214)
			πF.PushCheckpoint(18)
			// line 215: d[unicode] = _deepcopy_atomic
			πF.SetLineno(215)
			if πTemp013, πE = πg.ResolveGlobal(πF, ß_deepcopy_atomic); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πTemp013); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
				continue
			}
			πTemp016 = πTemp017
			if πE = πg.SetItem(πF, πTemp015, πTemp016, πTemp014); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label17
		Label18:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp019, πTemp020 = πF.ExcInfo()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßNameError); πE != nil {
				continue
			}
			if πTemp011, πE = πg.IsInstance(πF, πTemp019.ToObject(), πTemp013); πE != nil {
				continue
			}
			if πTemp011 {
				goto Label19
			}
			πE = πF.Raise(πTemp019.ToObject(), nil, πTemp020.ToObject())
			continue
			// line 216: except NameError:
			πF.SetLineno(216)
		Label19:
			// line 217: pass
			πF.SetLineno(217)
			πF.RestoreExc(nil, nil)
			goto Label17
		Label17:
			// line 218: try:
			πF.SetLineno(218)
			πF.PushCheckpoint(21)
			// line 219: d[types.CodeType] = _deepcopy_atomic
			πF.SetLineno(219)
			if πTemp013, πE = πg.ResolveGlobal(πF, ß_deepcopy_atomic); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πTemp013); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
				continue
			}
			if πTemp018, πE = πg.GetAttr(πF, πTemp017, ßCodeType, nil); πE != nil {
				continue
			}
			πTemp016 = πTemp018
			if πE = πg.SetItem(πF, πTemp015, πTemp016, πTemp014); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label20
		Label21:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp019, πTemp020 = πF.ExcInfo()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
				continue
			}
			if πTemp011, πE = πg.IsInstance(πF, πTemp019.ToObject(), πTemp013); πE != nil {
				continue
			}
			if πTemp011 {
				goto Label22
			}
			πE = πF.Raise(πTemp019.ToObject(), nil, πTemp020.ToObject())
			continue
			// line 220: except AttributeError:
			πF.SetLineno(220)
		Label22:
			// line 221: pass
			πF.SetLineno(221)
			πF.RestoreExc(nil, nil)
			goto Label20
		Label20:
			// line 222: d[type] = _deepcopy_atomic
			πF.SetLineno(222)
			if πTemp013, πE = πg.ResolveGlobal(πF, ß_deepcopy_atomic); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πTemp013); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			πTemp016 = πTemp017
			if πE = πg.SetItem(πF, πTemp015, πTemp016, πTemp014); πE != nil {
				continue
			}
			// line 223: d[xrange] = _deepcopy_atomic
			πF.SetLineno(223)
			if πTemp013, πE = πg.ResolveGlobal(πF, ß_deepcopy_atomic); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πTemp013); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
				continue
			}
			πTemp016 = πTemp017
			if πE = πg.SetItem(πF, πTemp015, πTemp016, πTemp014); πE != nil {
				continue
			}
			// line 225: d[types.BuiltinFunctionType] = _deepcopy_atomic
			πF.SetLineno(225)
			if πTemp013, πE = πg.ResolveGlobal(πF, ß_deepcopy_atomic); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πTemp013); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
				continue
			}
			if πTemp018, πE = πg.GetAttr(πF, πTemp017, ßBuiltinFunctionType, nil); πE != nil {
				continue
			}
			πTemp016 = πTemp018
			if πE = πg.SetItem(πF, πTemp015, πTemp016, πTemp014); πE != nil {
				continue
			}
			// line 226: d[types.FunctionType] = _deepcopy_atomic
			πF.SetLineno(226)
			if πTemp013, πE = πg.ResolveGlobal(πF, ß_deepcopy_atomic); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πTemp013); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
				continue
			}
			if πTemp018, πE = πg.GetAttr(πF, πTemp017, ßFunctionType, nil); πE != nil {
				continue
			}
			πTemp016 = πTemp018
			if πE = πg.SetItem(πF, πTemp015, πTemp016, πTemp014); πE != nil {
				continue
			}
			// line 227: d[WeakRefType] = _deepcopy_atomic
			πF.SetLineno(227)
			if πTemp013, πE = πg.ResolveGlobal(πF, ß_deepcopy_atomic); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πTemp013); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßWeakRefType); πE != nil {
				continue
			}
			πTemp016 = πTemp017
			if πE = πg.SetItem(πF, πTemp015, πTemp016, πTemp014); πE != nil {
				continue
			}
			// line 229: def _deepcopy_list(x, memo):
			πF.SetLineno(229)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "x", Def: nil}
			πTemp006[1] = πg.Param{Name: "memo", Def: nil}
			πTemp013 = πg.NewFunction(πg.NewCode("_deepcopy_list", "build/src/__python__/copy.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µmemo *πg.Object = πArgs[1]; _ = µmemo
				var µy *πg.Object = πg.UnboundLocal; _ = µy
				var µa *πg.Object = πg.UnboundLocal; _ = µa
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
				var πTemp008 []*πg.Object
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
					// line 230: y = []
					πF.SetLineno(230)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µy = πTemp002
					// line 231: memo[id(x)] = y
					πF.SetLineno(231)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µy); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					if πTemp004, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp003 = πTemp005
					if πE = πg.SetItem(πF, µmemo, πTemp003, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µx); πE != nil {
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
						µa = πTemp003
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 233: y.append(deepcopy(a, memo))
					πF.SetLineno(233)
					πTemp001 = πF.MakeArgs(1)
					πTemp008 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp008[0] = µa
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					πTemp008[1] = µmemo
					if πTemp003, πE = πg.ResolveGlobal(πF, ßdeepcopy); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					πTemp001[0] = πTemp004
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µy, ßappend, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 234: return y
					πF.SetLineno(234)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					πR = µy
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_deepcopy_list.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 235: d[list] = _deepcopy_list
			πF.SetLineno(235)
			if πTemp014, πE = πg.ResolveGlobal(πF, ß_deepcopy_list); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πTemp014); πE != nil {
				continue
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			if πTemp018, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
				continue
			}
			πTemp017 = πTemp018
			if πE = πg.SetItem(πF, πTemp016, πTemp017, πTemp015); πE != nil {
				continue
			}
			// line 237: def _deepcopy_tuple(x, memo):
			πF.SetLineno(237)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "x", Def: nil}
			πTemp006[1] = πg.Param{Name: "memo", Def: nil}
			πTemp014 = πg.NewFunction(πg.NewCode("_deepcopy_tuple", "build/src/__python__/copy.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µmemo *πg.Object = πArgs[1]; _ = µmemo
				var µy *πg.Object = πg.UnboundLocal; _ = µy
				var µa *πg.Object = πg.UnboundLocal; _ = µa
				var µd *πg.Object = πg.UnboundLocal; _ = µd
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.BaseException
				_ = πTemp008
				var πTemp009 *πg.Traceback
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
					case 8: goto Label8
					case 1: goto Label1
					case 2: goto Label2
					case 5: goto Label5
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					// line 238: y = []
					πF.SetLineno(238)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µy = πTemp002
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µx); πE != nil {
						continue
					}
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
					if πTemp005, πE = πg.Next(πF, πTemp002); πE != nil {
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
						µa = πTemp005
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 240: y.append(deepcopy(a, memo))
					πF.SetLineno(240)
					πTemp001 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp006[0] = µa
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					πTemp006[1] = µmemo
					if πTemp005, πE = πg.ResolveGlobal(πF, ßdeepcopy); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp001[0] = πTemp007
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µy, ßappend, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 241: d = id(x)
					πF.SetLineno(241)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					if πTemp002, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µd = πTemp005
					// line 242: try:
					πF.SetLineno(242)
					πF.PushCheckpoint(5)
					// line 243: return memo[d]
					πF.SetLineno(243)
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					πTemp002 = µd
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µmemo, πTemp002); πE != nil {
						continue
					}
					πR = πTemp005
					continue
					πF.PopCheckpoint()
					goto Label4
				Label5:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label6
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 244: except KeyError:
					πF.SetLineno(244)
				Label6:
					// line 245: pass
					πF.SetLineno(245)
					πF.RestoreExc(nil, nil)
					goto Label4
				Label4:
					πTemp001 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp006[0] = µx
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp001[0] = πTemp007
					if πTemp005, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Iter(πF, πTemp007); πE != nil {
						continue
					}
					πF.PushCheckpoint(8)
					πTemp003 = false
				Label7:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp003 {
						πF.PopCheckpoint()
						goto Label9
					}
					if πTemp005, πE = πg.Next(πF, πTemp002); πE != nil {
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
						µi = πTemp005
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(7)            
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp007 = µi
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µx, πTemp007); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp007 = µi
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetItem(πF, µy, πTemp007); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(πTemp010 != πTemp011).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label10
					}
					goto Label11
					// line 247: if x[i] is not y[i]:
					πF.SetLineno(247)
				Label10:
					// line 248: y = tuple(y)
					πF.SetLineno(248)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					πTemp001[0] = µy
					if πTemp005, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µy = πTemp007
					// line 249: break
					πF.SetLineno(249)
					πTemp003 = true
					continue
					goto Label11
				Label11:
					continue
				Label8:
					if πE != nil || πR != nil {
						continue
					}
					// line 251: y = x
					πF.SetLineno(251)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					µy = µx
				Label9:
					// line 252: memo[d] = y
					πF.SetLineno(252)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µy); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					πTemp005 = µd
					if πE = πg.SetItem(πF, µmemo, πTemp005, πTemp002); πE != nil {
						continue
					}
					// line 253: return y
					πF.SetLineno(253)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					πR = µy
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_deepcopy_tuple.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 254: d[tuple] = _deepcopy_tuple
			πF.SetLineno(254)
			if πTemp015, πE = πg.ResolveGlobal(πF, ß_deepcopy_tuple); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πTemp015); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			if πTemp021, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
				continue
			}
			πTemp018 = πTemp021
			if πE = πg.SetItem(πF, πTemp017, πTemp018, πTemp016); πE != nil {
				continue
			}
			// line 256: def _deepcopy_dict(x, memo):
			πF.SetLineno(256)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "x", Def: nil}
			πTemp006[1] = πg.Param{Name: "memo", Def: nil}
			πTemp015 = πg.NewFunction(πg.NewCode("_deepcopy_dict", "build/src/__python__/copy.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µmemo *πg.Object = πArgs[1]; _ = µmemo
				var µy *πg.Object = πg.UnboundLocal; _ = µy
				var µkey *πg.Object = πg.UnboundLocal; _ = µkey
				var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
				var πTemp001 *πg.Dict
				_ = πTemp001
				var πTemp002 *πg.Object
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
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 257: y = {}
					πF.SetLineno(257)
					πTemp001 = πg.NewDict()
					πTemp002 = πTemp001.ToObject()
					µy = πTemp002
					// line 258: memo[id(x)] = y
					πF.SetLineno(258)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µy); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp004[0] = µx
					if πTemp005, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp003 = πTemp006
					if πE = πg.SetItem(πF, µmemo, πTemp003, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µx, ßiteritems, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp007 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp007 {
						πF.PopCheckpoint()
						goto Label3
					}
					if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp003); πE != nil {
							continue
						}
						µkey = πTemp005
						µvalue = πTemp006
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 260: y[deepcopy(key, memo)] = deepcopy(value, memo)
					πF.SetLineno(260)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp004[0] = µvalue
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					πTemp004[1] = µmemo
					if πTemp003, πE = πg.ResolveGlobal(πF, ßdeepcopy); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp004[0] = µkey
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					πTemp004[1] = µmemo
					if πTemp009, πE = πg.ResolveGlobal(πF, ßdeepcopy); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp006 = πTemp010
					if πE = πg.SetItem(πF, µy, πTemp006, πTemp003); πE != nil {
						continue
					}
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 261: return y
					πF.SetLineno(261)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					πR = µy
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_deepcopy_dict.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 262: d[dict] = _deepcopy_dict
			πF.SetLineno(262)
			if πTemp016, πE = πg.ResolveGlobal(πF, ß_deepcopy_dict); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πTemp016); πE != nil {
				continue
			}
			if πTemp018, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			if πTemp022, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
				continue
			}
			πTemp021 = πTemp022
			if πE = πg.SetItem(πF, πTemp018, πTemp021, πTemp017); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßPyStringMap); πE != nil {
				continue
			}
			if πTemp018, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp016 = πg.GetBool(πTemp017 != πTemp018).ToObject()
			if πTemp011, πE = πg.IsTrue(πF, πTemp016); πE != nil {
				continue
			}
			if πTemp011 {
				goto Label23
			}
			goto Label24
			// line 263: if PyStringMap is not None:
			πF.SetLineno(263)
		Label23:
			// line 264: d[PyStringMap] = _deepcopy_dict
			πF.SetLineno(264)
			if πTemp016, πE = πg.ResolveGlobal(πF, ß_deepcopy_dict); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πTemp016); πE != nil {
				continue
			}
			if πTemp018, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			if πTemp022, πE = πg.ResolveGlobal(πF, ßPyStringMap); πE != nil {
				continue
			}
			πTemp021 = πTemp022
			if πE = πg.SetItem(πF, πTemp018, πTemp021, πTemp017); πE != nil {
				continue
			}
			goto Label24
		Label24:
			// line 266: def _deepcopy_method(x, memo): # Copy instance methods
			πF.SetLineno(266)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "x", Def: nil}
			πTemp006[1] = πg.Param{Name: "memo", Def: nil}
			πTemp016 = πg.NewFunction(πg.NewCode("_deepcopy_method", "build/src/__python__/copy.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µmemo *πg.Object = πArgs[1]; _ = µmemo
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
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
					// line 267: return type(x)(x.im_func, deepcopy(x.im_self, memo), x.im_class)
					πF.SetLineno(267)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µx, ßim_func, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µx, ßim_self, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp002
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					πTemp003[1] = µmemo
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdeepcopy); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001[1] = πTemp004
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µx, ßim_class, nil); πE != nil {
						continue
					}
					πTemp001[2] = πTemp002
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp003[0] = µx
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
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
			if πE = πF.Globals().SetItem(πF, ß_deepcopy_method.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 268: _deepcopy_dispatch[types.MethodType] = _deepcopy_method
			πF.SetLineno(268)
			if πTemp017, πE = πg.ResolveGlobal(πF, ß_deepcopy_method); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp018}, πTemp017); πE != nil {
				continue
			}
			if πTemp021, πE = πg.ResolveGlobal(πF, ß_deepcopy_dispatch); πE != nil {
				continue
			}
			if πTemp023, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
				continue
			}
			if πTemp024, πE = πg.GetAttr(πF, πTemp023, ßMethodType, nil); πE != nil {
				continue
			}
			πTemp022 = πTemp024
			if πE = πg.SetItem(πF, πTemp021, πTemp022, πTemp018); πE != nil {
				continue
			}
			// line 270: def _keep_alive(x, memo):
			πF.SetLineno(270)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "x", Def: nil}
			πTemp006[1] = πg.Param{Name: "memo", Def: nil}
			πTemp017 = πg.NewFunction(πg.NewCode("_keep_alive", "build/src/__python__/copy.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µmemo *πg.Object = πArgs[1]; _ = µmemo
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
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
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 271: """Keeps a reference to the object x in the memo.
					πF.SetLineno(271)
					// line 280: try:
					πF.SetLineno(280)
					πF.PushCheckpoint(2)
					// line 281: memo[id(memo)].append(x)
					πF.SetLineno(281)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					πTemp003[0] = µmemo
					if πTemp004, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp002 = πTemp005
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µmemo, πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp006, πTemp007 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
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
					// line 282: except KeyError:
					πF.SetLineno(282)
				Label3:
					// line 284: memo[id(memo)]=[x]
					πF.SetLineno(284)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					πTemp001[0] = µmemo
					if πTemp009, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp005 = πTemp010
					if πE = πg.SetItem(πF, µmemo, πTemp005, πTemp004); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ß_keep_alive.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 309: def _reconstruct(x, info, deep, memo=None):
			πF.SetLineno(309)
			πTemp006 = make([]πg.Param, 4)
			πTemp006[0] = πg.Param{Name: "x", Def: nil}
			πTemp006[1] = πg.Param{Name: "info", Def: nil}
			πTemp006[2] = πg.Param{Name: "deep", Def: nil}
			if πTemp021, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[3] = πg.Param{Name: "memo", Def: πTemp021}
			πTemp018 = πg.NewFunction(πg.NewCode("_reconstruct", "build/src/__python__/copy.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var µinfo *πg.Object = πArgs[1]; _ = µinfo
				var µdeep *πg.Object = πArgs[2]; _ = µdeep
				var µmemo *πg.Object = πArgs[3]; _ = µmemo
				var µn *πg.Object = πg.UnboundLocal; _ = µn
				var µcallable *πg.Object = πg.UnboundLocal; _ = µcallable
				var µargs *πg.Object = πg.UnboundLocal; _ = µargs
				var µstate *πg.Object = πg.UnboundLocal; _ = µstate
				var µlistiter *πg.Object = πg.UnboundLocal; _ = µlistiter
				var µdictiter *πg.Object = πg.UnboundLocal; _ = µdictiter
				var µy *πg.Object = πg.UnboundLocal; _ = µy
				var µslotstate *πg.Object = πg.UnboundLocal; _ = µslotstate
				var µkey *πg.Object = πg.UnboundLocal; _ = µkey
				var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
				var µitem *πg.Object = πg.UnboundLocal; _ = µitem
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 *πg.Dict
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 32: goto Label32
					case 36: goto Label36
					case 37: goto Label37
					case 43: goto Label43
					case 44: goto Label44
					case 31: goto Label31
					default: panic("unexpected function state")
					}
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µinfo, "info"); πE != nil {
						continue
					}
					πTemp001[0] = µinfo
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
					// line 310: if isinstance(info, str):
					πF.SetLineno(310)
				Label1:
					// line 311: return x
					πF.SetLineno(311)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πR = µx
					continue
					goto Label2
				Label2:
					// line 312: assert isinstance(info, tuple)
					πF.SetLineno(312)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µinfo, "info"); πE != nil {
						continue
					}
					πTemp001[0] = µinfo
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
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
					if πE = πg.Assert(πF, πTemp003, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µmemo == πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 313: if memo is None:
					πF.SetLineno(313)
				Label3:
					// line 314: memo = {}
					πF.SetLineno(314)
					πTemp005 = πg.NewDict()
					πTemp002 = πTemp005.ToObject()
					µmemo = πTemp002
					goto Label4
				Label4:
					// line 315: n = len(info)
					πF.SetLineno(315)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µinfo, "info"); πE != nil {
						continue
					}
					πTemp001[0] = µinfo
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µn = πTemp003
					// line 316: assert n in (2, 3, 4, 5)
					πF.SetLineno(316)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple4(πg.NewInt(2).ToObject(), πg.NewInt(3).ToObject(), πg.NewInt(4).ToObject(), πg.NewInt(5).ToObject()).ToObject()
					if πTemp004, πE = πg.Contains(πF, πTemp003, µn); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp004).ToObject()
					if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
						continue
					}
					// line 317: callable, args = info[:2]
					πF.SetLineno(317)
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(2).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µinfo, "info"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µinfo, πTemp002); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp006}}}, πTemp003); πE != nil {
						continue
					}
					µcallable = πTemp002
					µargs = πTemp006
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GT(πF, µn, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label5
					}
					goto Label6
					// line 318: if n > 2:
					πF.SetLineno(318)
				Label5:
					// line 319: state = info[2]
					πF.SetLineno(319)
					πTemp002 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µinfo, "info"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µinfo, πTemp002); πE != nil {
						continue
					}
					µstate = πTemp003
					goto Label7
				Label6:
					// line 321: state = None
					πF.SetLineno(321)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µstate = πTemp002
					goto Label7
				Label7:
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GT(πF, µn, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label8
					}
					goto Label9
					// line 322: if n > 3:
					πF.SetLineno(322)
				Label8:
					// line 323: listiter = info[3]
					πF.SetLineno(323)
					πTemp002 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µinfo, "info"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µinfo, πTemp002); πE != nil {
						continue
					}
					µlistiter = πTemp003
					goto Label10
				Label9:
					// line 325: listiter = None
					πF.SetLineno(325)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µlistiter = πTemp002
					goto Label10
				Label10:
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GT(πF, µn, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label11
					}
					goto Label12
					// line 326: if n > 4:
					πF.SetLineno(326)
				Label11:
					// line 327: dictiter = info[4]
					πF.SetLineno(327)
					πTemp002 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µinfo, "info"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µinfo, πTemp002); πE != nil {
						continue
					}
					µdictiter = πTemp003
					goto Label13
				Label12:
					// line 329: dictiter = None
					πF.SetLineno(329)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µdictiter = πTemp002
					goto Label13
				Label13:
					if πE = πg.CheckLocal(πF, µdeep, "deep"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µdeep); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label14
					}
					goto Label15
					// line 330: if deep:
					πF.SetLineno(330)
				Label14:
					// line 331: args = deepcopy(args, memo)
					πF.SetLineno(331)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp001[0] = µargs
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					πTemp001[1] = µmemo
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdeepcopy); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µargs = πTemp003
					goto Label15
				Label15:
					// line 332: y = callable(*args)
					πF.SetLineno(332)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcallable, "callable"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Invoke(πF, µcallable, nil, µargs, nil, nil); πE != nil {
						continue
					}
					µy = πTemp002
					// line 333: memo[id(x)] = y
					πF.SetLineno(333)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µy); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					if πTemp006, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp003 = πTemp007
					if πE = πg.SetItem(πF, µmemo, πTemp003, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µstate != πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label16
					}
					goto Label17
					// line 335: if state is not None:
					πF.SetLineno(335)
				Label16:
					if πE = πg.CheckLocal(πF, µdeep, "deep"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µdeep); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label18
					}
					goto Label19
					// line 336: if deep:
					πF.SetLineno(336)
				Label18:
					// line 337: state = deepcopy(state, memo)
					πF.SetLineno(337)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					πTemp001[0] = µstate
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					πTemp001[1] = µmemo
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdeepcopy); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µstate = πTemp003
					goto Label19
				Label19:
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					πTemp001[0] = µy
					πTemp001[1] = ß__setstate__.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
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
						goto Label20
					}
					goto Label21
					// line 338: if hasattr(y, '__setstate__'):
					πF.SetLineno(338)
				Label20:
					// line 339: y.__setstate__(state)
					πF.SetLineno(339)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					πTemp001[0] = µstate
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µy, ß__setstate__, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label22
				Label21:
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					πTemp001[0] = µstate
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp002 = πTemp006
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label23
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					πTemp001[0] = µstate
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.Eq(πF, πTemp007, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label23:
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label24
					}
					goto Label25
					// line 341: if isinstance(state, tuple) and len(state) == 2:
					πF.SetLineno(341)
				Label24:
					// line 342: state, slotstate = state
					πF.SetLineno(342)
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}}}, µstate); πE != nil {
						continue
					}
					µstate = πTemp002
					µslotstate = πTemp003
					goto Label26
				Label25:
					// line 344: slotstate = None
					πF.SetLineno(344)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µslotstate = πTemp002
					goto Label26
				Label26:
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µstate != πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label27
					}
					goto Label28
					// line 345: if state is not None:
					πF.SetLineno(345)
				Label27:
					// line 346: y.__dict__.update(state)
					πF.SetLineno(346)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					πTemp001[0] = µstate
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µy, ß__dict__, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßupdate, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label28
				Label28:
					if πE = πg.CheckLocal(πF, µslotstate, "slotstate"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µslotstate != πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label29
					}
					goto Label30
					// line 347: if slotstate is not None:
					πF.SetLineno(347)
				Label29:
					if πE = πg.CheckLocal(πF, µslotstate, "slotstate"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µslotstate, ßiteritems, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, πTemp006); πE != nil {
						continue
					}
					πF.PushCheckpoint(32)
					πTemp004 = false
				Label31:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label33
					}
					if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
							continue
						}
						µkey = πTemp006
						µvalue = πTemp007
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(31)            
					// line 349: setattr(y, key, value)
					πF.SetLineno(349)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					πTemp001[0] = µy
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp001[1] = µkey
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[2] = µvalue
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label32:
					if πE != nil || πR != nil {
						continue
					}
				Label33:
					goto Label30
				Label30:
					goto Label22
				Label22:
					goto Label17
				Label17:
					if πE = πg.CheckLocal(πF, µlistiter, "listiter"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µlistiter != πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label34
					}
					goto Label35
					// line 351: if listiter is not None:
					πF.SetLineno(351)
				Label34:
					if πE = πg.CheckLocal(πF, µlistiter, "listiter"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µlistiter); πE != nil {
						continue
					}
					πF.PushCheckpoint(37)
					πTemp004 = false
				Label36:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label38
					}
					if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
						µitem = πTemp003
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(36)            
					if πE = πg.CheckLocal(πF, µdeep, "deep"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µdeep); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label39
					}
					goto Label40
					// line 353: if deep:
					πF.SetLineno(353)
				Label39:
					// line 354: item = deepcopy(item, memo)
					πF.SetLineno(354)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp001[0] = µitem
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					πTemp001[1] = µmemo
					if πTemp003, πE = πg.ResolveGlobal(πF, ßdeepcopy); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µitem = πTemp006
					goto Label40
				Label40:
					// line 355: y.append(item)
					πF.SetLineno(355)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp001[0] = µitem
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µy, ßappend, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label37:
					if πE != nil || πR != nil {
						continue
					}
				Label38:
					goto Label35
				Label35:
					if πE = πg.CheckLocal(πF, µdictiter, "dictiter"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µdictiter != πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label41
					}
					goto Label42
					// line 356: if dictiter is not None:
					πF.SetLineno(356)
				Label41:
					if πE = πg.CheckLocal(πF, µdictiter, "dictiter"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µdictiter); πE != nil {
						continue
					}
					πF.PushCheckpoint(44)
					πTemp004 = false
				Label43:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label45
					}
					if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
							continue
						}
						µkey = πTemp006
						µvalue = πTemp007
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(43)            
					if πE = πg.CheckLocal(πF, µdeep, "deep"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µdeep); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label46
					}
					goto Label47
					// line 358: if deep:
					πF.SetLineno(358)
				Label46:
					// line 359: key = deepcopy(key, memo)
					πF.SetLineno(359)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp001[0] = µkey
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					πTemp001[1] = µmemo
					if πTemp003, πE = πg.ResolveGlobal(πF, ßdeepcopy); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µkey = πTemp006
					// line 360: value = deepcopy(value, memo)
					πF.SetLineno(360)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[0] = µvalue
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					πTemp001[1] = µmemo
					if πTemp003, πE = πg.ResolveGlobal(πF, ßdeepcopy); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µvalue = πTemp006
					goto Label47
				Label47:
					// line 361: y[key] = value
					πF.SetLineno(361)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µvalue); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp006 = µkey
					if πE = πg.SetItem(πF, µy, πTemp006, πTemp003); πE != nil {
						continue
					}
					continue
				Label44:
					if πE != nil || πR != nil {
						continue
					}
				Label45:
					goto Label42
				Label42:
					// line 362: return y
					πF.SetLineno(362)
					if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
						continue
					}
					πR = µy
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_reconstruct.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 364: del d
			πF.SetLineno(364)
			if πE = πg.DelVar(πF, πF.Globals(), ßd); πE != nil {
				continue
			}
			// line 366: del types
			πF.SetLineno(366)
			if πE = πg.DelVar(πF, πF.Globals(), ßtypes); πE != nil {
				continue
			}
			// line 372: def _test():
			πF.SetLineno(372)
			πTemp006 = make([]πg.Param, 0)
			πTemp021 = πg.NewFunction(πg.NewCode("_test", "build/src/__python__/copy.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µl *πg.Object = πg.UnboundLocal; _ = µl
				var µl1 *πg.Object = πg.UnboundLocal; _ = µl1
				var µC *πg.Object = πg.UnboundLocal; _ = µC
				var µc *πg.Object = πg.UnboundLocal; _ = µc
				var µl2 *πg.Object = πg.UnboundLocal; _ = µl2
				var µl3 *πg.Object = πg.UnboundLocal; _ = µl3
				var µrepr *πg.Object = πg.UnboundLocal; _ = µrepr
				var µodict *πg.Object = πg.UnboundLocal; _ = µodict
				var µo *πg.Object = πg.UnboundLocal; _ = µo
				var µx *πg.Object = πg.UnboundLocal; _ = µx
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Dict
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
					// line 374: l = [None, 1, 3.14, 'xyzzy', (1,), [3.14, 'abc'],
					πF.SetLineno(374)
					πTemp001 = make([]*πg.Object, 10)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					πTemp001[1] = πg.NewInt(1).ToObject()
					πTemp001[2] = πg.NewFloat(3.14).ToObject()
					πTemp001[3] = ßxyzzy.ToObject()
					πTemp002 = πg.NewTuple1(πg.NewInt(1).ToObject()).ToObject()
					πTemp001[4] = πTemp002
					πTemp003 = make([]*πg.Object, 2)
					πTemp003[0] = πg.NewFloat(3.14).ToObject()
					πTemp003[1] = ßabc.ToObject()
					πTemp002 = πg.NewList(πTemp003...).ToObject()
					πTemp001[5] = πTemp002
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßabc.ToObject(), ßABC.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp004.ToObject()
					πTemp001[6] = πTemp002
					πTemp002 = πg.NewTuple0().ToObject()
					πTemp001[7] = πTemp002
					πTemp003 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp003...).ToObject()
					πTemp001[8] = πTemp002
					πTemp004 = πg.NewDict()
					πTemp002 = πTemp004.ToObject()
					πTemp001[9] = πTemp002
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µl = πTemp002
					// line 376: l1 = copy(l)
					πF.SetLineno(376)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					πTemp001[0] = µl
					if πTemp002, πE = πg.ResolveGlobal(πF, ßcopy); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µl1 = πTemp005
					// line 377: print l1==l
					πF.SetLineno(377)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µl1, "l1"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µl1, µl); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					// line 378: l1 = map(copy, l)
					πF.SetLineno(378)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßcopy); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					πTemp001[1] = µl
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µl1 = πTemp005
					// line 379: print l1==l
					πF.SetLineno(379)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µl1, "l1"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µl1, µl); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					// line 380: l1 = deepcopy(l)
					πF.SetLineno(380)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					πTemp001[0] = µl
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdeepcopy); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µl1 = πTemp005
					// line 381: print l1==l
					πF.SetLineno(381)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µl1, "l1"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µl1, µl); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					// line 382: class C:
					πF.SetLineno(382)
					πTemp001 = make([]*πg.Object, 0)
					πTemp004 = πg.NewDict()
					if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
						continue
					}
					_, πE = πg.NewCode("C", "build/src/__python__/copy.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 383: def __init__(self, arg=None):
							πF.SetLineno(383)
							πTemp002 = make([]πg.Param, 2)
							πTemp002[0] = πg.Param{Name: "self", Def: nil}
							if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
								continue
							}
							πTemp002[1] = πg.Param{Name: "arg", Def: πTemp003}
							πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/copy.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
								var µarg *πg.Object = πArgs[1]; _ = µarg
								var µsys *πg.Object = πg.UnboundLocal; _ = µsys
								var µfile *πg.Object = πg.UnboundLocal; _ = µfile
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
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 384: self.a = 1
									πF.SetLineno(384)
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, µself, ßa, πTemp001); πE != nil {
										continue
									}
									// line 385: self.arg = arg
									πF.SetLineno(385)
									if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
										continue
									}
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µarg); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, µself, ßarg, πTemp001); πE != nil {
										continue
									}
									if πTemp002, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
										continue
									}
									if πTemp001, πE = πg.Eq(πF, πTemp002, ß__main__.ToObject()); πE != nil {
										continue
									}
									if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp003 {
										goto Label1
									}
									goto Label2
									// line 386: if __name__ == '__main__':
									πF.SetLineno(386)
								Label1:
									// line 387: import sys
									πF.SetLineno(387)
									if πTemp004, πE = πg.ImportModule(πF, "sys"); πE != nil {
										continue
									}
									πTemp001 = πTemp004[0]
									µsys = πTemp001
									// line 388: file = sys.argv[0]
									πF.SetLineno(388)
									πTemp001 = πg.NewInt(0).ToObject()
									if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
										continue
									}
									if πTemp005, πE = πg.GetAttr(πF, µsys, ßargv, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
										continue
									}
									µfile = πTemp002
									goto Label3
								Label2:
									// line 390: file = __file__
									πF.SetLineno(390)
									if πTemp001, πE = πg.ResolveGlobal(πF, ß__file__); πE != nil {
										continue
									}
									µfile = πTemp001
									goto Label3
								Label3:
									// line 391: self.fp = open(file)
									πF.SetLineno(391)
									πTemp004 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
										continue
									}
									πTemp004[0] = µfile
									if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp004)
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, µself, ßfp, πTemp001); πE != nil {
										continue
									}
									// line 392: self.fp.close()
									πF.SetLineno(392)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ßfp, nil); πE != nil {
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
							if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
								continue
							}
							// line 393: def __getstate__(self):
							πF.SetLineno(393)
							πTemp002 = make([]πg.Param, 1)
							πTemp002[0] = πg.Param{Name: "self", Def: nil}
							πTemp003 = πg.NewFunction(πg.NewCode("__getstate__", "build/src/__python__/copy.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
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
									// line 394: return {'a': self.a, 'arg': self.arg}
									πF.SetLineno(394)
									πTemp001 = πg.NewDict()
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µself, ßa, nil); πE != nil {
										continue
									}
									if πE = πTemp001.SetItem(πF, ßa.ToObject(), πTemp002); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µself, ßarg, nil); πE != nil {
										continue
									}
									if πE = πTemp001.SetItem(πF, ßarg.ToObject(), πTemp002); πE != nil {
										continue
									}
									πTemp002 = πTemp001.ToObject()
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
							if πE = πClass.SetItem(πF, ß__getstate__.ToObject(), πTemp003); πE != nil {
								continue
							}
							// line 395: def __setstate__(self, state):
							πF.SetLineno(395)
							πTemp002 = make([]πg.Param, 2)
							πTemp002[0] = πg.Param{Name: "self", Def: nil}
							πTemp002[1] = πg.Param{Name: "state", Def: nil}
							πTemp004 = πg.NewFunction(πg.NewCode("__setstate__", "build/src/__python__/copy.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
								var µstate *πg.Object = πArgs[1]; _ = µstate
								var µkey *πg.Object = πg.UnboundLocal; _ = µkey
								var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
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
									if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µstate, ßiteritems, nil); πE != nil {
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
										if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
											continue
										}
										µkey = πTemp003
										µvalue = πTemp006
									}
									if πE != nil || !πTemp005 {
										continue
									}
									πF.PushCheckpoint(1)            
									// line 397: setattr(self, key, value)
									πF.SetLineno(397)
									πTemp007 = πF.MakeArgs(3)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									πTemp007[0] = µself
									if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
										continue
									}
									πTemp007[1] = µkey
									if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
										continue
									}
									πTemp007[2] = µvalue
									if πTemp002, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp007)
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
							if πE = πClass.SetItem(πF, ß__setstate__.ToObject(), πTemp004); πE != nil {
								continue
							}
							// line 398: def __deepcopy__(self, memo=None):
							πF.SetLineno(398)
							πTemp002 = make([]πg.Param, 2)
							πTemp002[0] = πg.Param{Name: "self", Def: nil}
							if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
								continue
							}
							πTemp002[1] = πg.Param{Name: "memo", Def: πTemp006}
							πTemp005 = πg.NewFunction(πg.NewCode("__deepcopy__", "build/src/__python__/copy.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
								var µmemo *πg.Object = πArgs[1]; _ = µmemo
								var µnew *πg.Object = πg.UnboundLocal; _ = µnew
								var πTemp001 []*πg.Object
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
									// line 399: new = self.__class__(deepcopy(self.arg, memo))
									πF.SetLineno(399)
									πTemp001 = πF.MakeArgs(1)
									πTemp002 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µself, ßarg, nil); πE != nil {
										continue
									}
									πTemp002[0] = πTemp003
									if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
										continue
									}
									πTemp002[1] = µmemo
									if πTemp003, πE = πg.ResolveGlobal(πF, ßdeepcopy); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									πTemp001[0] = πTemp004
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									µnew = πTemp004
									// line 400: new.a = self.a
									πF.SetLineno(400)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µself, ßa, nil); πE != nil {
										continue
									}
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, µnew, ßa, πTemp004); πE != nil {
										continue
									}
									// line 401: return new
									πF.SetLineno(401)
									if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
										continue
									}
									πR = µnew
									continue
								}
								if πE != nil {
									πR = nil
								} else if πR == nil {
									πR = πg.None
								}
								return πR, πE
							}), πF.Globals()).ToObject()
							if πE = πClass.SetItem(πF, ß__deepcopy__.ToObject(), πTemp005); πE != nil {
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
					if πTemp006, πE = πTemp005.Call(πF, []*πg.Object{πg.NewStr("C").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
						continue
					}
					µC = πTemp006
					// line 402: c = C('argument sketch')
					πF.SetLineno(402)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("argument sketch").ToObject()
					if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
						continue
					}
					if πTemp002, πE = µC.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µc = πTemp002
					// line 403: l.append(c)
					πF.SetLineno(403)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp001[0] = µc
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µl, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 404: l2 = copy(l)
					πF.SetLineno(404)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					πTemp001[0] = µl
					if πTemp002, πE = πg.ResolveGlobal(πF, ßcopy); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µl2 = πTemp005
					// line 405: print l == l2
					πF.SetLineno(405)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µl2, "l2"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µl, µl2); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					// line 406: print l
					πF.SetLineno(406)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					πTemp001[0] = µl
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					// line 407: print l2
					πF.SetLineno(407)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µl2, "l2"); πE != nil {
						continue
					}
					πTemp001[0] = µl2
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					// line 408: l2 = deepcopy(l)
					πF.SetLineno(408)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					πTemp001[0] = µl
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdeepcopy); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µl2 = πTemp005
					// line 409: print l == l2
					πF.SetLineno(409)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µl2, "l2"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µl, µl2); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					// line 410: print l
					πF.SetLineno(410)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					πTemp001[0] = µl
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					// line 411: print l2
					πF.SetLineno(411)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µl2, "l2"); πE != nil {
						continue
					}
					πTemp001[0] = µl2
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					// line 412: l.append({l[1]: l, 'xyz': l[2]})
					πF.SetLineno(412)
					πTemp001 = πF.MakeArgs(1)
					πTemp004 = πg.NewDict()
					πTemp002 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µl, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, πTemp005, µl); πE != nil {
						continue
					}
					πTemp002 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µl, πTemp002); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßxyz.ToObject(), πTemp005); πE != nil {
						continue
					}
					πTemp002 = πTemp004.ToObject()
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µl, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 413: l3 = copy(l)
					πF.SetLineno(413)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					πTemp001[0] = µl
					if πTemp002, πE = πg.ResolveGlobal(πF, ßcopy); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µl3 = πTemp005
					// line 414: import repr
					πF.SetLineno(414)
					if πTemp001, πE = πg.ImportModule(πF, "repr"); πE != nil {
						continue
					}
					πTemp002 = πTemp001[0]
					µrepr = πTemp002
					// line 415: print map(repr.repr, l)
					πF.SetLineno(415)
					πTemp001 = make([]*πg.Object, 1)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µrepr, "repr"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µrepr, ßrepr, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp002
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					πTemp003[1] = µl
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001[0] = πTemp005
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					// line 416: print map(repr.repr, l1)
					πF.SetLineno(416)
					πTemp001 = make([]*πg.Object, 1)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µrepr, "repr"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µrepr, ßrepr, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp002
					if πE = πg.CheckLocal(πF, µl1, "l1"); πE != nil {
						continue
					}
					πTemp003[1] = µl1
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001[0] = πTemp005
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					// line 417: print map(repr.repr, l2)
					πF.SetLineno(417)
					πTemp001 = make([]*πg.Object, 1)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µrepr, "repr"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µrepr, ßrepr, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp002
					if πE = πg.CheckLocal(πF, µl2, "l2"); πE != nil {
						continue
					}
					πTemp003[1] = µl2
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001[0] = πTemp005
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					// line 418: print map(repr.repr, l3)
					πF.SetLineno(418)
					πTemp001 = make([]*πg.Object, 1)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µrepr, "repr"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µrepr, ßrepr, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp002
					if πE = πg.CheckLocal(πF, µl3, "l3"); πE != nil {
						continue
					}
					πTemp003[1] = µl3
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001[0] = πTemp005
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					// line 419: l3 = deepcopy(l)
					πF.SetLineno(419)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					πTemp001[0] = µl
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdeepcopy); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µl3 = πTemp005
					// line 420: import repr
					πF.SetLineno(420)
					if πTemp001, πE = πg.ImportModule(πF, "repr"); πE != nil {
						continue
					}
					πTemp002 = πTemp001[0]
					µrepr = πTemp002
					// line 421: print map(repr.repr, l)
					πF.SetLineno(421)
					πTemp001 = make([]*πg.Object, 1)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µrepr, "repr"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µrepr, ßrepr, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp002
					if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
						continue
					}
					πTemp003[1] = µl
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001[0] = πTemp005
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					// line 422: print map(repr.repr, l1)
					πF.SetLineno(422)
					πTemp001 = make([]*πg.Object, 1)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µrepr, "repr"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µrepr, ßrepr, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp002
					if πE = πg.CheckLocal(πF, µl1, "l1"); πE != nil {
						continue
					}
					πTemp003[1] = µl1
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001[0] = πTemp005
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					// line 423: print map(repr.repr, l2)
					πF.SetLineno(423)
					πTemp001 = make([]*πg.Object, 1)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µrepr, "repr"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µrepr, ßrepr, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp002
					if πE = πg.CheckLocal(πF, µl2, "l2"); πE != nil {
						continue
					}
					πTemp003[1] = µl2
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001[0] = πTemp005
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					// line 424: print map(repr.repr, l3)
					πF.SetLineno(424)
					πTemp001 = make([]*πg.Object, 1)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µrepr, "repr"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µrepr, ßrepr, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp002
					if πE = πg.CheckLocal(πF, µl3, "l3"); πE != nil {
						continue
					}
					πTemp003[1] = µl3
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001[0] = πTemp005
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					// line 425: class odict(dict):
					πF.SetLineno(425)
					πTemp001 = make([]*πg.Object, 1)
					if πTemp006, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
						continue
					}
					πTemp001[0] = πTemp006
					πTemp004 = πg.NewDict()
					if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
						continue
					}
					_, πE = πg.NewCode("odict", "build/src/__python__/copy.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
						πClass := πTemp004
						_ = πClass
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []πg.Param
						_ = πTemp002
						var πTemp003 *πg.Dict
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 426: def __init__(self, d = {}):
							πF.SetLineno(426)
							πTemp002 = make([]πg.Param, 2)
							πTemp002[0] = πg.Param{Name: "self", Def: nil}
							πTemp003 = πg.NewDict()
							πTemp004 = πTemp003.ToObject()
							πTemp002[1] = πg.Param{Name: "d", Def: πTemp004}
							πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/copy.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
								var µd *πg.Object = πArgs[1]; _ = µd
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
									// line 427: self.a = 99
									πF.SetLineno(427)
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(99).ToObject()); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, µself, ßa, πTemp001); πE != nil {
										continue
									}
									// line 428: dict.__init__(self, d)
									πF.SetLineno(428)
									πTemp002 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									πTemp002[0] = µself
									if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
										continue
									}
									πTemp002[1] = µd
									if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
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
							// line 429: def __setitem__(self, k, i):
							πF.SetLineno(429)
							πTemp002 = make([]πg.Param, 3)
							πTemp002[0] = πg.Param{Name: "self", Def: nil}
							πTemp002[1] = πg.Param{Name: "k", Def: nil}
							πTemp002[2] = πg.Param{Name: "i", Def: nil}
							πTemp004 = πg.NewFunction(πg.NewCode("__setitem__", "build/src/__python__/copy.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
								var µk *πg.Object = πArgs[1]; _ = µk
								var µi *πg.Object = πArgs[2]; _ = µi
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
									// line 430: dict.__setitem__(self, k, i)
									πF.SetLineno(430)
									πTemp001 = πF.MakeArgs(3)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									πTemp001[0] = µself
									if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
										continue
									}
									πTemp001[1] = µk
									if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
										continue
									}
									πTemp001[2] = µi
									if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__setitem__, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									// line 431: self.a
									πF.SetLineno(431)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µself, ßa, nil); πE != nil {
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
							if πE = πClass.SetItem(πF, ß__setitem__.ToObject(), πTemp004); πE != nil {
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
					if πTemp006, πE = πTemp005.Call(πF, []*πg.Object{πg.NewStr("odict").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
						continue
					}
					µodict = πTemp006
					// line 432: o = odict({"A" : "B"})
					πF.SetLineno(432)
					πTemp001 = πF.MakeArgs(1)
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßA.ToObject(), ßB.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp004.ToObject()
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µodict, "odict"); πE != nil {
						continue
					}
					if πTemp002, πE = µodict.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µo = πTemp002
					// line 433: x = deepcopy(o)
					πF.SetLineno(433)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					πTemp001[0] = µo
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdeepcopy); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µx = πTemp005
					// line 434: print(o, x)
					πF.SetLineno(434)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µo, µx).ToObject()
					πTemp001[0] = πTemp002
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_test.ToObject(), πTemp021); πE != nil {
				continue
			}
			if πTemp023, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp022, πE = πg.Eq(πF, πTemp023, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp011, πE = πg.IsTrue(πF, πTemp022); πE != nil {
				continue
			}
			if πTemp011 {
				goto Label25
			}
			goto Label26
			// line 436: if __name__ == '__main__':
			πF.SetLineno(436)
		Label25:
			// line 437: _test()
			πF.SetLineno(437)
			if πTemp022, πE = πg.ResolveGlobal(πF, ß_test); πE != nil {
				continue
			}
			if πTemp023, πE = πTemp022.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label26
		Label26:
		}
		return nil, πE
	})
	πg.RegisterModule("copy", Code)
}
