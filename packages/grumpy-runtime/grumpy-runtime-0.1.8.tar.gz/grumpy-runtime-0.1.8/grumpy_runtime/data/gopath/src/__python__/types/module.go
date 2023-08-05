package types
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/types.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßBooleanType := πg.InternStr("BooleanType")
		ßBuiltinFunctionType := πg.InternStr("BuiltinFunctionType")
		ßBuiltinMethodType := πg.InternStr("BuiltinMethodType")
		ßClassType := πg.InternStr("ClassType")
		ßComplexType := πg.InternStr("ComplexType")
		ßDictType := πg.InternStr("DictType")
		ßDictionaryType := πg.InternStr("DictionaryType")
		ßEllipsis := πg.InternStr("Ellipsis")
		ßEllipsisType := πg.InternStr("EllipsisType")
		ßFileType := πg.InternStr("FileType")
		ßFloatType := πg.InternStr("FloatType")
		ßFrameType := πg.InternStr("FrameType")
		ßFunctionType := πg.InternStr("FunctionType")
		ßGeneratorType := πg.InternStr("GeneratorType")
		ßIntType := πg.InternStr("IntType")
		ßListType := πg.InternStr("ListType")
		ßMethodType := πg.InternStr("MethodType")
		ßModuleType := πg.InternStr("ModuleType")
		ßNameError := πg.InternStr("NameError")
		ßNone := πg.InternStr("None")
		ßNoneType := πg.InternStr("NoneType")
		ßNotImplemented := πg.InternStr("NotImplemented")
		ßNotImplementedType := πg.InternStr("NotImplementedType")
		ßObjectType := πg.InternStr("ObjectType")
		ßSliceType := πg.InternStr("SliceType")
		ßStringType := πg.InternStr("StringType")
		ßStringTypes := πg.InternStr("StringTypes")
		ßTracebackType := πg.InternStr("TracebackType")
		ßTupleType := πg.InternStr("TupleType")
		ßTypeError := πg.InternStr("TypeError")
		ßTypeType := πg.InternStr("TypeType")
		ßUnboundMethodType := πg.InternStr("UnboundMethodType")
		ßUnicodeType := πg.InternStr("UnicodeType")
		ßXRangeType := πg.InternStr("XRangeType")
		ß_C := πg.InternStr("_C")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_f := πg.InternStr("_f")
		ß_g := πg.InternStr("_g")
		ß_m := πg.InternStr("_m")
		ß_x := πg.InternStr("_x")
		ßappend := πg.InternStr("append")
		ßbool := πg.InternStr("bool")
		ßcomplex := πg.InternStr("complex")
		ßdict := πg.InternStr("dict")
		ßexc_info := πg.InternStr("exc_info")
		ßfile := πg.InternStr("file")
		ßfloat := πg.InternStr("float")
		ßint := πg.InternStr("int")
		ßlen := πg.InternStr("len")
		ßlist := πg.InternStr("list")
		ßobject := πg.InternStr("object")
		ßslice := πg.InternStr("slice")
		ßstr := πg.InternStr("str")
		ßsys := πg.InternStr("sys")
		ßtb := πg.InternStr("tb")
		ßtb_frame := πg.InternStr("tb_frame")
		ßtuple := πg.InternStr("tuple")
		ßtype := πg.InternStr("type")
		ßunicode := πg.InternStr("unicode")
		ßxrange := πg.InternStr("xrange")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.BaseException
		_ = πTemp004
		var πTemp005 *πg.Traceback
		_ = πTemp005
		var πTemp006 bool
		_ = πTemp006
		var πTemp007 *πg.Object
		_ = πTemp007
		var πTemp008 []πg.Param
		_ = πTemp008
		var πTemp009 *πg.Object
		_ = πTemp009
		var πTemp010 *πg.Dict
		_ = πTemp010
		var πTemp011 *πg.Object
		_ = πTemp011
		var πTemp012 []*πg.Object
		_ = πTemp012
		var πTemp013 *πg.Object
		_ = πTemp013
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 8: goto Label8
			case 2: goto Label2
			case 5: goto Label5
			default: panic("unexpected function state")
			}
			// line 1: """Define names for all type symbols known in the standard interpreter.
			πF.SetLineno(1)
			// line 5: import sys
			πF.SetLineno(5)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 12: NoneType = type(None)
			πF.SetLineno(12)
			πTemp002 = πF.MakeArgs(1)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp002[0] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßNoneType.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 13: TypeType = type
			πF.SetLineno(13)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTypeType.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 14: ObjectType = object
			πF.SetLineno(14)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßObjectType.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 16: IntType = int
			πF.SetLineno(16)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßIntType.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: FloatType = float
			πF.SetLineno(18)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFloatType.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 19: BooleanType = bool
			πF.SetLineno(19)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßbool); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBooleanType.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 20: try:
			πF.SetLineno(20)
			πF.PushCheckpoint(2)
			// line 21: ComplexType = complex
			πF.SetLineno(21)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßcomplex); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßComplexType.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label1
		Label2:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp004, πTemp005 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNameError); πE != nil {
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
			// line 22: except NameError:
			πF.SetLineno(22)
		Label3:
			// line 23: pass
			πF.SetLineno(23)
			πF.RestoreExc(nil, nil)
			goto Label1
		Label1:
			// line 25: StringType = str
			πF.SetLineno(25)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStringType.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 30: try:
			πF.SetLineno(30)
			πF.PushCheckpoint(5)
			// line 31: UnicodeType = unicode
			πF.SetLineno(31)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßUnicodeType.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 32: StringTypes = (StringType, UnicodeType)
			πF.SetLineno(32)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßStringType); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßUnicodeType); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(πTemp003, πTemp007).ToObject()
			if πE = πF.Globals().SetItem(πF, ßStringTypes.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label4
		Label5:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp004, πTemp005 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNameError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label6
			}
			πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
			continue
			// line 33: except NameError:
			πF.SetLineno(33)
		Label6:
			// line 34: StringTypes = (StringType,)
			πF.SetLineno(34)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßStringType); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple1(πTemp003).ToObject()
			if πE = πF.Globals().SetItem(πF, ßStringTypes.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label4
		Label4:
			// line 38: TupleType = tuple
			πF.SetLineno(38)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTupleType.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 39: ListType = list
			πF.SetLineno(39)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßListType.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 40: DictType = DictionaryType = dict
			πF.SetLineno(40)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDictType.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDictionaryType.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 42: def _f(): pass
			πF.SetLineno(42)
			πTemp008 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("_f", "build/src/__python__/types.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 42: def _f(): pass
					πF.SetLineno(42)
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_f.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 43: FunctionType = type(_f)
			πF.SetLineno(43)
			πTemp002 = πF.MakeArgs(1)
			if πTemp003, πE = πg.ResolveGlobal(πF, ß_f); πE != nil {
				continue
			}
			πTemp002[0] = πTemp003
			if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßFunctionType.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 47: def _g():
			πF.SetLineno(47)
			πTemp008 = make([]πg.Param, 0)
			πTemp003 = πg.NewFunction(πg.NewCode("_g", "build/src/__python__/types.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 1: goto Label1
						default: panic("unexpected function state")
						}
						// line 48: yield 1
						πF.SetLineno(48)
						πF.PushCheckpoint(1)
						return πg.NewInt(1).ToObject(), nil
					Label1:
						πTemp001 = πSent
					}
					return nil, πE
				}).ToObject(), nil
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_g.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 49: GeneratorType = type(_g())
			πF.SetLineno(49)
			πTemp002 = πF.MakeArgs(1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ß_g); πE != nil {
				continue
			}
			if πTemp009, πE = πTemp007.Call(πF, nil, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp007, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßGeneratorType.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 51: class _C(object):
			πF.SetLineno(51)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp011, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp011
			πTemp010 = πg.NewDict()
			if πTemp007, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp010.SetItem(πF, ß__module__.ToObject(), πTemp007); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_C", "build/src/__python__/types.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp010
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
					// line 52: def _m(self): pass
					πF.SetLineno(52)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("_m", "build/src/__python__/types.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 52: def _m(self): pass
							πF.SetLineno(52)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_m.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp009, πE = πTemp010.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp009 == nil {
				πTemp009 = πg.TypeType.ToObject()
			}
			if πTemp011, πE = πTemp009.Call(πF, []*πg.Object{πg.NewStr("_C").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp010.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_C.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 53: ClassType = type(_C)
			πF.SetLineno(53)
			πTemp002 = πF.MakeArgs(1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ß_C); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßClassType.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 54: UnboundMethodType = type(_C._m)         # Same as MethodType
			πF.SetLineno(54)
			πTemp002 = πF.MakeArgs(1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ß_C); πE != nil {
				continue
			}
			if πTemp009, πE = πg.GetAttr(πF, πTemp007, ß_m, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp007, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßUnboundMethodType.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 55: _x = _C()
			πF.SetLineno(55)
			if πTemp007, πE = πg.ResolveGlobal(πF, ß_C); πE != nil {
				continue
			}
			if πTemp009, πE = πTemp007.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_x.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 57: MethodType = type(_x._m)
			πF.SetLineno(57)
			πTemp002 = πF.MakeArgs(1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ß_x); πE != nil {
				continue
			}
			if πTemp009, πE = πg.GetAttr(πF, πTemp007, ß_m, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp007, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßMethodType.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 59: BuiltinFunctionType = type(len)
			πF.SetLineno(59)
			πTemp002 = πF.MakeArgs(1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßBuiltinFunctionType.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 60: BuiltinMethodType = type([].append)     # Same as BuiltinFunctionType
			πF.SetLineno(60)
			πTemp002 = πF.MakeArgs(1)
			πTemp012 = make([]*πg.Object, 0)
			πTemp007 = πg.NewList(πTemp012...).ToObject()
			if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßappend, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp007, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßBuiltinMethodType.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 62: ModuleType = type(sys)
			πF.SetLineno(62)
			πTemp002 = πF.MakeArgs(1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßModuleType.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 63: FileType = file
			πF.SetLineno(63)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßfile); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFileType.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 64: XRangeType = xrange
			πF.SetLineno(64)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßXRangeType.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 66: try:
			πF.SetLineno(66)
			πF.PushCheckpoint(8)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
				continue
			}
			// line 67: raise TypeError
			πF.SetLineno(67)
			πE = πF.Raise(πTemp007, nil, nil)
			continue
			πF.PopCheckpoint()
			goto Label7
		Label8:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp004, πTemp005 = πF.ExcInfo()
			if πTemp007, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp007); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label9
			}
			πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
			continue
			// line 68: except TypeError:
			πF.SetLineno(68)
		Label9:
			// line 69: tb = sys.exc_info()[2]
			πF.SetLineno(69)
			πTemp007 = πg.NewInt(2).ToObject()
			if πTemp011, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp013, πE = πg.GetAttr(πF, πTemp011, ßexc_info, nil); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp013.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp009, πE = πg.GetItem(πF, πTemp011, πTemp007); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßtb.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 70: TracebackType = type(tb)
			πF.SetLineno(70)
			πTemp002 = πF.MakeArgs(1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßtb); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßTracebackType.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 71: FrameType = type(tb.tb_frame)
			πF.SetLineno(71)
			πTemp002 = πF.MakeArgs(1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßtb); πE != nil {
				continue
			}
			if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßtb_frame, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp007, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßFrameType.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 72: del tb
			πF.SetLineno(72)
			if πE = πg.DelVar(πF, πF.Globals(), ßtb); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label7
		Label7:
			// line 74: SliceType = slice
			πF.SetLineno(74)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSliceType.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 75: EllipsisType = type(Ellipsis)
			πF.SetLineno(75)
			πTemp002 = πF.MakeArgs(1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßEllipsis); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßEllipsisType.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 78: NotImplementedType = type(NotImplemented)
			πF.SetLineno(78)
			πTemp002 = πF.MakeArgs(1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßNotImplementedType.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 84: del sys, _C, _x                           # Not for export
			πF.SetLineno(84)
			if πE = πg.DelVar(πF, πF.Globals(), ßsys); πE != nil {
				continue
			}
			if πE = πg.DelVar(πF, πF.Globals(), ß_C); πE != nil {
				continue
			}
			if πE = πg.DelVar(πF, πF.Globals(), ß_x); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("types", Code)
}
