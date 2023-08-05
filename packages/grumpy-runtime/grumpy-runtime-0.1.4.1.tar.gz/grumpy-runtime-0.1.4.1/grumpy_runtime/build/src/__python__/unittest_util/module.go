package unittest_util
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/unittest_util.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßException := πg.InternStr("Exception")
		ßFalse := πg.InternStr("False")
		ßIndexError := πg.InternStr("IndexError")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßValueError := πg.InternStr("ValueError")
		ß_MAX_LENGTH := πg.InternStr("_MAX_LENGTH")
		ß_Mismatch := πg.InternStr("_Mismatch")
		ß__dict__ := πg.InternStr("__dict__")
		ß__getnewargs__ := πg.InternStr("__getnewargs__")
		ß__getstate__ := πg.InternStr("__getstate__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__new__ := πg.InternStr("__new__")
		ß__repr__ := πg.InternStr("__repr__")
		ß__slots__ := πg.InternStr("__slots__")
		ß__unittest := πg.InternStr("__unittest")
		ß_asdict := πg.InternStr("_asdict")
		ß_count_diff_all_purpose := πg.InternStr("_count_diff_all_purpose")
		ß_count_diff_hashable := πg.InternStr("_count_diff_hashable")
		ß_fields := πg.InternStr("_fields")
		ß_itemgetter := πg.InternStr("_itemgetter")
		ß_make := πg.InternStr("_make")
		ß_ordered_count := πg.InternStr("_ordered_count")
		ß_property := πg.InternStr("_property")
		ß_replace := πg.InternStr("_replace")
		ß_tuple := πg.InternStr("_tuple")
		ßactual := πg.InternStr("actual")
		ßappend := πg.InternStr("append")
		ßclassmethod := πg.InternStr("classmethod")
		ßdict := πg.InternStr("dict")
		ßenumerate := πg.InternStr("enumerate")
		ßexpected := πg.InternStr("expected")
		ßextend := πg.InternStr("extend")
		ßget := πg.InternStr("get")
		ßitemgetter := πg.InternStr("itemgetter")
		ßitems := πg.InternStr("items")
		ßkeys := πg.InternStr("keys")
		ßlen := πg.InternStr("len")
		ßlist := πg.InternStr("list")
		ßmap := πg.InternStr("map")
		ßobject := πg.InternStr("object")
		ßoperator := πg.InternStr("operator")
		ßpop := πg.InternStr("pop")
		ßproperty := πg.InternStr("property")
		ßrange := πg.InternStr("range")
		ßremove := πg.InternStr("remove")
		ßrepr := πg.InternStr("repr")
		ßsafe_repr := πg.InternStr("safe_repr")
		ßsorted_list_difference := πg.InternStr("sorted_list_difference")
		ßstrclass := πg.InternStr("strclass")
		ßtuple := πg.InternStr("tuple")
		ßunorderable_list_difference := πg.InternStr("unorderable_list_difference")
		ßvalue := πg.InternStr("value")
		ßzip := πg.InternStr("zip")
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
		var πTemp008 *πg.Dict
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
			// line 1: """Various utility functions."""
			πF.SetLineno(1)
			// line 4: import operator
			πF.SetLineno(4)
			if πTemp002, πE = πg.ImportModule(πF, "operator"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßoperator.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 5: _itemgetter = operator.itemgetter
			πF.SetLineno(5)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßitemgetter, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_itemgetter.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 6: _property = property
			πF.SetLineno(6)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßproperty); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_property.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 7: _tuple = tuple
			πF.SetLineno(7)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_tuple.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 9: __unittest = True
			πF.SetLineno(9)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__unittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 11: _MAX_LENGTH = 80
			πF.SetLineno(11)
			if πE = πF.Globals().SetItem(πF, ß_MAX_LENGTH.ToObject(), πg.NewInt(80).ToObject()); πE != nil {
				continue
			}
			// line 12: def safe_repr(obj, short=False):
			πF.SetLineno(12)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "obj", Def: nil}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[1] = πg.Param{Name: "short", Def: πTemp003}
			πTemp001 = πg.NewFunction(πg.NewCode("safe_repr", "build/src/__python__/unittest_util.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µshort *πg.Object = πArgs[1]; _ = µshort
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var πTemp001 []*πg.Object
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
				var πTemp007 bool
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
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 13: try:
					πF.SetLineno(13)
					πF.PushCheckpoint(2)
					// line 14: result = repr(obj)
					πF.SetLineno(14)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µresult = πTemp003
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 15: except Exception:
					πF.SetLineno(15)
				Label3:
					// line 16: result = object.__repr__(obj)
					πF.SetLineno(16)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					if πTemp002, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__repr__, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µresult = πTemp002
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					if πE = πg.CheckLocal(πF, µshort, "short"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, µshort); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp007).ToObject()
					πTemp002 = πTemp003
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label4
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp001[0] = µresult
					if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp008, πE = πg.ResolveGlobal(πF, ß_MAX_LENGTH); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LT(πF, πTemp009, πTemp008); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label4:
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label5
					}
					goto Label6
					// line 17: if not short or len(result) < _MAX_LENGTH:
					πF.SetLineno(17)
				Label5:
					// line 18: return result
					πF.SetLineno(18)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πR = µresult
					continue
					goto Label6
				Label6:
					// line 19: return result[:_MAX_LENGTH] + ' [truncated]...'
					πF.SetLineno(19)
					if πTemp008, πE = πg.ResolveGlobal(πF, ß_MAX_LENGTH); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp008, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µresult, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp008, πg.NewStr(" [truncated]...").ToObject()); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßsafe_repr.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 22: def strclass(cls):
			πF.SetLineno(22)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "cls", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("strclass", "build/src/__python__/unittest_util.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcls *πg.Object = πArgs[0]; _ = µcls
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
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
					// line 23: return "%s.%s" % (cls.__module__, cls.__name__)
					πF.SetLineno(23)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µcls, ß__module__, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µcls, ß__name__, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s.%s").ToObject(), πTemp002); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßstrclass.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 25: def sorted_list_difference(expected, actual):
			πF.SetLineno(25)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "expected", Def: nil}
			πTemp004[1] = πg.Param{Name: "actual", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("sorted_list_difference", "build/src/__python__/unittest_util.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µexpected *πg.Object = πArgs[0]; _ = µexpected
				var µactual *πg.Object = πArgs[1]; _ = µactual
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µj *πg.Object = πg.UnboundLocal; _ = µj
				var µmissing *πg.Object = πg.UnboundLocal; _ = µmissing
				var µunexpected *πg.Object = πg.UnboundLocal; _ = µunexpected
				var µe *πg.Object = πg.UnboundLocal; _ = µe
				var µa *πg.Object = πg.UnboundLocal; _ = µa
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
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.BaseException
				_ = πTemp008
				var πTemp009 *πg.Traceback
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					case 5: goto Label5
					case 10: goto Label10
					case 11: goto Label11
					case 13: goto Label13
					case 14: goto Label14
					case 16: goto Label16
					case 17: goto Label17
					case 18: goto Label18
					case 20: goto Label20
					case 21: goto Label21
					default: panic("unexpected function state")
					}
					// line 26: """Finds elements in only one or the other of two, sorted input lists.
					πF.SetLineno(26)
					// line 33: i = j = 0
					πF.SetLineno(33)
					µi = πg.NewInt(0).ToObject()
					µj = πg.NewInt(0).ToObject()
					// line 34: missing = []
					πF.SetLineno(34)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µmissing = πTemp002
					// line 35: unexpected = []
					πF.SetLineno(35)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µunexpected = πTemp002
					// line 36: while True:
					πF.SetLineno(36)
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
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 37: try:
					πF.SetLineno(37)
					πF.PushCheckpoint(5)
					// line 38: e = expected[i]
					πF.SetLineno(38)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp002 = µi
					if πE = πg.CheckLocal(πF, µexpected, "expected"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µexpected, πTemp002); πE != nil {
						continue
					}
					µe = πTemp005
					// line 39: a = actual[j]
					πF.SetLineno(39)
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					πTemp002 = µj
					if πE = πg.CheckLocal(πF, µactual, "actual"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µactual, πTemp002); πE != nil {
						continue
					}
					µa = πTemp005
					if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LT(πF, µe, µa); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label6
					}
					if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GT(πF, µe, µa); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label7
					}
					goto Label8
					// line 40: if e < a:
					πF.SetLineno(40)
				Label6:
					// line 41: missing.append(e)
					πF.SetLineno(41)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
						continue
					}
					πTemp001[0] = µe
					if πE = πg.CheckLocal(πF, µmissing, "missing"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µmissing, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 42: i += 1
					πF.SetLineno(42)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µi = πTemp002
					// line 43: while expected[i] == e:
					πF.SetLineno(43)
					πF.PushCheckpoint(11)
					πTemp004 = false
				Label10:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label12
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp005 = µi
					if πE = πg.CheckLocal(πF, µexpected, "expected"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µexpected, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp007, µe); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(10)            
					// line 44: i += 1
					πF.SetLineno(44)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µi = πTemp002
					continue
				Label11:
					if πE != nil || πR != nil {
						continue
					}
				Label12:
					goto Label9
					// line 45: elif e > a:
					πF.SetLineno(45)
				Label7:
					// line 46: unexpected.append(a)
					πF.SetLineno(46)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp001[0] = µa
					if πE = πg.CheckLocal(πF, µunexpected, "unexpected"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µunexpected, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 47: j += 1
					πF.SetLineno(47)
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µj, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µj = πTemp002
					// line 48: while actual[j] == a:
					πF.SetLineno(48)
					πF.PushCheckpoint(14)
					πTemp004 = false
				Label13:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label15
					}
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					πTemp005 = µj
					if πE = πg.CheckLocal(πF, µactual, "actual"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µactual, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp007, µa); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(13)            
					// line 49: j += 1
					πF.SetLineno(49)
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µj, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µj = πTemp002
					continue
				Label14:
					if πE != nil || πR != nil {
						continue
					}
				Label15:
					goto Label9
				Label8:
					// line 51: i += 1
					πF.SetLineno(51)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µi = πTemp002
					// line 52: try:
					πF.SetLineno(52)
					πF.PushCheckpoint(16)
					// line 53: while expected[i] == e:
					πF.SetLineno(53)
					πF.PushCheckpoint(18)
					πTemp004 = false
				Label17:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label19
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp005 = µi
					if πE = πg.CheckLocal(πF, µexpected, "expected"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µexpected, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp007, µe); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(17)            
					// line 54: i += 1
					πF.SetLineno(54)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µi = πTemp002
					continue
				Label18:
					if πE != nil || πR != nil {
						continue
					}
				Label19:
					πF.PopCheckpoint()
					goto Label16
				Label16:
					πTemp008, πTemp009 = πF.RestoreExc(nil, nil)
					// line 56: j += 1
					πF.SetLineno(56)
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µj, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µj = πTemp002
					// line 57: while actual[j] == a:
					πF.SetLineno(57)
					πF.PushCheckpoint(21)
					πTemp004 = false
				Label20:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label22
					}
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					πTemp005 = µj
					if πE = πg.CheckLocal(πF, µactual, "actual"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µactual, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp007, µa); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(20)            
					// line 58: j += 1
					πF.SetLineno(58)
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µj, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µj = πTemp002
					continue
				Label21:
					if πE != nil || πR != nil {
						continue
					}
				Label22:
					if πTemp008 != nil {
						πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
						continue
					}
					if πR != nil {
						continue
					}
					goto Label9
				Label9:
					πF.PopCheckpoint()
					goto Label4
				Label5:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label23
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 59: except IndexError:
					πF.SetLineno(59)
				Label23:
					// line 60: missing.extend(expected[i:])
					πF.SetLineno(60)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µi, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µexpected, "expected"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µexpected, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp005
					if πE = πg.CheckLocal(πF, µmissing, "missing"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µmissing, ßextend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 61: unexpected.extend(actual[j:])
					πF.SetLineno(61)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µj, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µactual, "actual"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µactual, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp005
					if πE = πg.CheckLocal(πF, µunexpected, "unexpected"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µunexpected, ßextend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 62: break
					πF.SetLineno(62)
					πTemp003 = true
					continue
					πF.RestoreExc(nil, nil)
					goto Label4
				Label4:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 63: return missing, unexpected
					πF.SetLineno(63)
					if πE = πg.CheckLocal(πF, µmissing, "missing"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µunexpected, "unexpected"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µmissing, µunexpected).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßsorted_list_difference.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 66: def unorderable_list_difference(expected, actual, ignore_duplicate=False):
			πF.SetLineno(66)
			πTemp004 = make([]πg.Param, 3)
			πTemp004[0] = πg.Param{Name: "expected", Def: nil}
			πTemp004[1] = πg.Param{Name: "actual", Def: nil}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp004[2] = πg.Param{Name: "ignore_duplicate", Def: πTemp007}
			πTemp006 = πg.NewFunction(πg.NewCode("unorderable_list_difference", "build/src/__python__/unittest_util.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µexpected *πg.Object = πArgs[0]; _ = µexpected
				var µactual *πg.Object = πArgs[1]; _ = µactual
				var µignore_duplicate *πg.Object = πArgs[2]; _ = µignore_duplicate
				var µmissing *πg.Object = πg.UnboundLocal; _ = µmissing
				var µunexpected *πg.Object = πg.UnboundLocal; _ = µunexpected
				var µitem *πg.Object = πg.UnboundLocal; _ = µitem
				var µlst *πg.Object = πg.UnboundLocal; _ = µlst
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
				var πTemp006 *πg.BaseException
				_ = πTemp006
				var πTemp007 *πg.Traceback
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 bool
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
					case 5: goto Label5
					case 9: goto Label9
					case 10: goto Label10
					case 13: goto Label13
					case 14: goto Label14
					case 15: goto Label15
					case 20: goto Label20
					case 21: goto Label21
					case 24: goto Label24
					case 25: goto Label25
					case 26: goto Label26
					default: panic("unexpected function state")
					}
					// line 67: """Same behavior as sorted_list_difference but
					πF.SetLineno(67)
					// line 73: missing = []
					πF.SetLineno(73)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µmissing = πTemp002
					// line 74: unexpected = []
					πF.SetLineno(74)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µunexpected = πTemp002
					// line 75: while expected:
					πF.SetLineno(75)
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
					if πE = πg.CheckLocal(πF, µexpected, "expected"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µexpected); πE != nil {
						continue
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 76: item = expected.pop()
					πF.SetLineno(76)
					if πE = πg.CheckLocal(πF, µexpected, "expected"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µexpected, ßpop, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µitem = πTemp005
					// line 77: try:
					πF.SetLineno(77)
					πF.PushCheckpoint(5)
					// line 78: actual.remove(item)
					πF.SetLineno(78)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp001[0] = µitem
					if πE = πg.CheckLocal(πF, µactual, "actual"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µactual, ßremove, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.PopCheckpoint()
					goto Label4
				Label5:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp006, πTemp007 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label6
					}
					πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
					continue
					// line 79: except ValueError:
					πF.SetLineno(79)
				Label6:
					// line 80: missing.append(item)
					πF.SetLineno(80)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp001[0] = µitem
					if πE = πg.CheckLocal(πF, µmissing, "missing"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µmissing, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.RestoreExc(nil, nil)
					goto Label4
				Label4:
					if πE = πg.CheckLocal(πF, µignore_duplicate, "ignore_duplicate"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µignore_duplicate); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label7
					}
					goto Label8
					// line 81: if ignore_duplicate:
					πF.SetLineno(81)
				Label7:
					if πE = πg.CheckLocal(πF, µexpected, "expected"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µactual, "actual"); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple2(µexpected, µactual).ToObject()
					if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
						continue
					}
					πF.PushCheckpoint(10)
					πTemp004 = false
				Label9:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label11
					}
					if πTemp005, πE = πg.Next(πF, πTemp002); πE != nil {
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
						µlst = πTemp005
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(9)            
					// line 83: try:
					πF.SetLineno(83)
					πF.PushCheckpoint(13)
					// line 84: while True:
					πF.SetLineno(84)
					πF.PushCheckpoint(15)
					πTemp008 = false
				Label14:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
						πF.PopCheckpoint()
						goto Label16
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(14)            
					// line 85: lst.remove(item)
					πF.SetLineno(85)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp001[0] = µitem
					if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µlst, ßremove, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label15:
					if πE != nil || πR != nil {
						continue
					}
				Label16:
					πF.PopCheckpoint()
					goto Label12
				Label13:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp006, πTemp007 = πF.ExcInfo()
					if πTemp005, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp005); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label17
					}
					πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
					continue
					// line 86: except ValueError:
					πF.SetLineno(86)
				Label17:
					// line 87: pass
					πF.SetLineno(87)
					πF.RestoreExc(nil, nil)
					goto Label12
				Label12:
					continue
				Label10:
					if πE != nil || πR != nil {
						continue
					}
				Label11:
					goto Label8
				Label8:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					if πE = πg.CheckLocal(πF, µignore_duplicate, "ignore_duplicate"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µignore_duplicate); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label18
					}
					goto Label19
					// line 88: if ignore_duplicate:
					πF.SetLineno(88)
				Label18:
					// line 89: while actual:
					πF.SetLineno(89)
					πF.PushCheckpoint(21)
					πTemp003 = false
				Label20:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp003 {
						πF.PopCheckpoint()
						goto Label22
					}
					if πE = πg.CheckLocal(πF, µactual, "actual"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µactual); πE != nil {
						continue
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(20)            
					// line 90: item = actual.pop()
					πF.SetLineno(90)
					if πE = πg.CheckLocal(πF, µactual, "actual"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µactual, ßpop, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µitem = πTemp005
					// line 91: unexpected.append(item)
					πF.SetLineno(91)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp001[0] = µitem
					if πE = πg.CheckLocal(πF, µunexpected, "unexpected"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µunexpected, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 92: try:
					πF.SetLineno(92)
					πF.PushCheckpoint(24)
					// line 93: while True:
					πF.SetLineno(93)
					πF.PushCheckpoint(26)
					πTemp004 = false
				Label25:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label27
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(25)            
					// line 94: actual.remove(item)
					πF.SetLineno(94)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp001[0] = µitem
					if πE = πg.CheckLocal(πF, µactual, "actual"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µactual, ßremove, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label26:
					if πE != nil || πR != nil {
						continue
					}
				Label27:
					πF.PopCheckpoint()
					goto Label23
				Label24:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp006, πTemp007 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label28
					}
					πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
					continue
					// line 95: except ValueError:
					πF.SetLineno(95)
				Label28:
					// line 96: pass
					πF.SetLineno(96)
					πF.RestoreExc(nil, nil)
					goto Label23
				Label23:
					continue
				Label21:
					if πE != nil || πR != nil {
						continue
					}
				Label22:
					// line 97: return missing, unexpected
					πF.SetLineno(97)
					if πE = πg.CheckLocal(πF, µmissing, "missing"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µunexpected, "unexpected"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µmissing, µunexpected).ToObject()
					πR = πTemp002
					continue
					goto Label19
				Label19:
					// line 100: return missing, actual
					πF.SetLineno(100)
					if πE = πg.CheckLocal(πF, µmissing, "missing"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µactual, "actual"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µmissing, µactual).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßunorderable_list_difference.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 103: class _Mismatch(tuple):
			πF.SetLineno(103)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp010, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
				continue
			}
			πTemp002[0] = πTemp010
			πTemp008 = πg.NewDict()
			if πTemp007, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, ß__module__.ToObject(), πTemp007); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_Mismatch", "build/src/__python__/unittest_util.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp008
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
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 []*πg.Object
				_ = πTemp011
				var πTemp012 *πg.Object
				_ = πTemp012
				var πTemp013 πg.KWArgs
				_ = πTemp013
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 104: 'Mismatch(actual, expected, value)'
					πF.SetLineno(104)
					// line 106: __slots__ = ()
					πF.SetLineno(106)
					πTemp001 = πg.NewTuple0().ToObject()
					if πE = πClass.SetItem(πF, ß__slots__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 108: _fields = ('actual', 'expected', 'value')
					πF.SetLineno(108)
					πTemp001 = πg.NewTuple3(ßactual.ToObject(), ßexpected.ToObject(), ßvalue.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ß_fields.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 110: def __new__(_cls, actual, expected, value):
					πF.SetLineno(110)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "_cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "actual", Def: nil}
					πTemp002[2] = πg.Param{Name: "expected", Def: nil}
					πTemp002[3] = πg.Param{Name: "value", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__new__", "build/src/__python__/unittest_util.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µ_cls *πg.Object = πArgs[0]; _ = µ_cls
						var µactual *πg.Object = πArgs[1]; _ = µactual
						var µexpected *πg.Object = πArgs[2]; _ = µexpected
						var µvalue *πg.Object = πArgs[3]; _ = µvalue
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
							// line 111: 'Create new instance of Mismatch(actual, expected, value)'
							πF.SetLineno(111)
							// line 112: return _tuple.__new__(_cls, (actual, expected, value))
							πF.SetLineno(112)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µ_cls, "_cls"); πE != nil {
								continue
							}
							πTemp001[0] = µ_cls
							if πE = πg.CheckLocal(πF, µactual, "actual"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µexpected, "expected"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(µactual, µexpected, µvalue).ToObject()
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_tuple); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__new__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__new__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 115: def _make(cls, iterable, new=tuple.__new__, len=len):
					πF.SetLineno(115)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "iterable", Def: nil}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßtuple); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß__new__, nil); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "new", Def: πTemp005}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßlen); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "len", Def: πTemp004}
					πTemp003 = πg.NewFunction(πg.NewCode("_make", "build/src/__python__/unittest_util.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µiterable *πg.Object = πArgs[1]; _ = µiterable
						var µnew *πg.Object = πArgs[2]; _ = µnew
						var µlen *πg.Object = πArgs[3]; _ = µlen
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
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
							default: panic("unexpected function state")
							}
							// line 116: 'Make a new Mismatch object from a sequence or iterable'
							πF.SetLineno(116)
							// line 117: result = new(cls, iterable)
							πF.SetLineno(117)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp001[0] = µcls
							if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
								continue
							}
							πTemp001[1] = µiterable
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							if πTemp002, πE = µnew.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µresult = πTemp002
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							if πE = πg.CheckLocal(πF, µlen, "len"); πE != nil {
								continue
							}
							if πTemp003, πE = µlen.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.NE(πF, πTemp003, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 118: if len(result) != 3:
							πF.SetLineno(118)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp005[0] = µresult
							if πE = πg.CheckLocal(πF, µlen, "len"); πE != nil {
								continue
							}
							if πTemp003, πE = µlen.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Expected 3 arguments, got %d").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 119: raise TypeError('Expected 3 arguments, got %d' % len(result))
							πF.SetLineno(119)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							// line 120: return result
							πF.SetLineno(120)
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
					if πE = πClass.SetItem(πF, ß_make.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 121: _make = classmethod(_make)
					πF.SetLineno(121)
					πTemp006 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß_make); πE != nil {
						continue
					}
					πTemp006[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ß_make.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 123: def __repr__(self):
					πF.SetLineno(123)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__repr__", "build/src/__python__/unittest_util.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 124: 'Return a nicely formatted representation string'
							πF.SetLineno(124)
							// line 125: return 'Mismatch(actual=%r, expected=%r, value=%r)' % self
							πF.SetLineno(125)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Mismatch(actual=%r, expected=%r, value=%r)").ToObject(), µself); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 127: def _asdict(self):
					πF.SetLineno(127)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("_asdict", "build/src/__python__/unittest_util.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 128: 'Return a new OrderedDict which maps field names to their values'
							πF.SetLineno(128)
							// line 130: return dict(zip(self._fields, self))
							πF.SetLineno(130)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_fields, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[1] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßzip); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_asdict.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 132: def _replace(_self, **kwds):
					πF.SetLineno(132)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "_self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("_replace", "build/src/__python__/unittest_util.py", πTemp002, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µ_self *πg.Object = πArgs[0]; _ = µ_self
						var µkwds *πg.Object = πArgs[1]; _ = µkwds
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
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
							// line 133: 'Return a new Mismatch object replacing specified fields with new values'
							πF.SetLineno(133)
							// line 134: result = _self._make(map(kwds.pop, ('actual', 'expected', 'value'), _self))
							πF.SetLineno(134)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µkwds, ßpop, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							πTemp003 = πg.NewTuple3(ßactual.ToObject(), ßexpected.ToObject(), ßvalue.ToObject()).ToObject()
							πTemp002[1] = πTemp003
							if πE = πg.CheckLocal(πF, µ_self, "_self"); πE != nil {
								continue
							}
							πTemp002[2] = µ_self
							if πTemp003, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µ_self, "_self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µ_self, ß_make, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µresult = πTemp004
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µkwds); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 135: if kwds:
							πF.SetLineno(135)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µkwds, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("Got unexpected field names: %r").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 136: raise ValueError('Got unexpected field names: %r' % kwds.keys())
							πF.SetLineno(136)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							// line 137: return result
							πF.SetLineno(137)
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
					if πE = πClass.SetItem(πF, ß_replace.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 139: def __getnewargs__(self):
					πF.SetLineno(139)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("__getnewargs__", "build/src/__python__/unittest_util.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 140: 'Return self as a plain tuple.  Used by copy and pickle.'
							πF.SetLineno(140)
							// line 141: return tuple(self)
							πF.SetLineno(141)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__getnewargs__.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 143: __dict__ = _property(_asdict)
					πF.SetLineno(143)
					πTemp006 = πF.MakeArgs(1)
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ß_asdict); πE != nil {
						continue
					}
					πTemp006[0] = πTemp009
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ß_property); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ß__dict__.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 145: def __getstate__(self):
					πF.SetLineno(145)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("__getstate__", "build/src/__python__/unittest_util.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 146: 'Exclude the OrderedDict from pickling'
							πF.SetLineno(146)
							// line 147: pass
							πF.SetLineno(147)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__getstate__.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 149: actual = _property(_itemgetter(0), doc='Alias for field number 0')
					πF.SetLineno(149)
					πTemp006 = πF.MakeArgs(1)
					πTemp011 = πF.MakeArgs(1)
					πTemp011[0] = πg.NewInt(0).ToObject()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_itemgetter); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp010.Call(πF, πTemp011, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp011)
					πTemp006[0] = πTemp012
					πTemp013 = πg.KWArgs{
						{"doc", πg.NewStr("Alias for field number 0").ToObject()},
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_property); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp010.Call(πF, πTemp006, πTemp013); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßactual.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 151: expected = _property(_itemgetter(1), doc='Alias for field number 1')
					πF.SetLineno(151)
					πTemp006 = πF.MakeArgs(1)
					πTemp011 = πF.MakeArgs(1)
					πTemp011[0] = πg.NewInt(1).ToObject()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_itemgetter); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp010.Call(πF, πTemp011, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp011)
					πTemp006[0] = πTemp012
					πTemp013 = πg.KWArgs{
						{"doc", πg.NewStr("Alias for field number 1").ToObject()},
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_property); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp010.Call(πF, πTemp006, πTemp013); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßexpected.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 153: value = _property(_itemgetter(2), doc='Alias for field number 2')
					πF.SetLineno(153)
					πTemp006 = πF.MakeArgs(1)
					πTemp011 = πF.MakeArgs(1)
					πTemp011[0] = πg.NewInt(2).ToObject()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_itemgetter); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp010.Call(πF, πTemp011, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp011)
					πTemp006[0] = πTemp012
					πTemp013 = πg.KWArgs{
						{"doc", πg.NewStr("Alias for field number 2").ToObject()},
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_property); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp010.Call(πF, πTemp006, πTemp013); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßvalue.ToObject(), πTemp012); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp009, πE = πTemp008.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp009 == nil {
				πTemp009 = πg.TypeType.ToObject()
			}
			if πTemp010, πE = πTemp009.Call(πF, []*πg.Object{πg.NewStr("_Mismatch").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp008.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Mismatch.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 155: def _count_diff_all_purpose(actual, expected):
			πF.SetLineno(155)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "actual", Def: nil}
			πTemp004[1] = πg.Param{Name: "expected", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("_count_diff_all_purpose", "build/src/__python__/unittest_util.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µactual *πg.Object = πArgs[0]; _ = µactual
				var µexpected *πg.Object = πArgs[1]; _ = µexpected
				var µs *πg.Object = πg.UnboundLocal; _ = µs
				var µt *πg.Object = πg.UnboundLocal; _ = µt
				var µm *πg.Object = πg.UnboundLocal; _ = µm
				var µn *πg.Object = πg.UnboundLocal; _ = µn
				var µNULL *πg.Object = πg.UnboundLocal; _ = µNULL
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µelem *πg.Object = πg.UnboundLocal; _ = µelem
				var µcnt_s *πg.Object = πg.UnboundLocal; _ = µcnt_s
				var µcnt_t *πg.Object = πg.UnboundLocal; _ = µcnt_t
				var µj *πg.Object = πg.UnboundLocal; _ = µj
				var µother_elem *πg.Object = πg.UnboundLocal; _ = µother_elem
				var µdiff *πg.Object = πg.UnboundLocal; _ = µdiff
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
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					case 6: goto Label6
					case 7: goto Label7
					case 11: goto Label11
					case 12: goto Label12
					case 18: goto Label18
					case 19: goto Label19
					case 23: goto Label23
					case 24: goto Label24
					default: panic("unexpected function state")
					}
					// line 156: 'Returns list of (cnt_act, cnt_exp, elem) triples where the counts differ'
					πF.SetLineno(156)
					// line 158: s, t = list(actual), list(expected)
					πF.SetLineno(158)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µactual, "actual"); πE != nil {
						continue
					}
					πTemp002[0] = µactual
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µexpected, "expected"); πE != nil {
						continue
					}
					πTemp002[0] = µexpected
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
						continue
					}
					µs = πTemp003
					µt = πTemp004
					// line 159: m, n = len(s), len(t)
					πF.SetLineno(159)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp002[0] = µs
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					πTemp002[0] = µt
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
						continue
					}
					µm = πTemp003
					µn = πTemp004
					// line 160: NULL = object()
					πF.SetLineno(160)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µNULL = πTemp003
					// line 161: result = []
					πF.SetLineno(161)
					πTemp002 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp002...).ToObject()
					µresult = πTemp001
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp002[0] = µs
					if πTemp003, πE = πg.ResolveGlobal(πF, ßenumerate); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
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
					if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
							continue
						}
						µi = πTemp004
						µelem = πTemp005
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(1)            
					if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µNULL, "NULL"); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µelem == µNULL).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label4
					}
					goto Label5
					// line 163: if elem is NULL:
					πF.SetLineno(163)
				Label4:
					// line 164: continue
					πF.SetLineno(164)
					continue
					goto Label5
				Label5:
					// line 165: cnt_s = cnt_t = 0
					πF.SetLineno(165)
					µcnt_s = πg.NewInt(0).ToObject()
					µcnt_t = πg.NewInt(0).ToObject()
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp002[0] = µi
					if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
						continue
					}
					πTemp002[1] = µm
					if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.Iter(πF, πTemp005); πE != nil {
						continue
					}
					πF.PushCheckpoint(7)
					πTemp007 = false
				Label6:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp007 {
						πF.PopCheckpoint()
						goto Label8
					}
					if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
						µj = πTemp004
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(6)            
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					πTemp005 = µj
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, µs, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, πTemp009, µelem); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label9
					}
					goto Label10
					// line 167: if s[j] == elem:
					πF.SetLineno(167)
				Label9:
					// line 168: cnt_s += 1
					πF.SetLineno(168)
					if πE = πg.CheckLocal(πF, µcnt_s, "cnt_s"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IAdd(πF, µcnt_s, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µcnt_s = πTemp004
					// line 169: s[j] = NULL
					πF.SetLineno(169)
					if πE = πg.CheckLocal(πF, µNULL, "NULL"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µNULL); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					πTemp005 = µj
					if πE = πg.SetItem(πF, µs, πTemp005, πTemp004); πE != nil {
						continue
					}
					goto Label10
				Label10:
					continue
				Label7:
					if πE != nil || πR != nil {
						continue
					}
				Label8:
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					πTemp002[0] = µt
					if πTemp004, πE = πg.ResolveGlobal(πF, ßenumerate); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.Iter(πF, πTemp005); πE != nil {
						continue
					}
					πF.PushCheckpoint(12)
					πTemp007 = false
				Label11:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp007 {
						πF.PopCheckpoint()
						goto Label13
					}
					if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp009}}}, πTemp004); πE != nil {
							continue
						}
						µj = πTemp005
						µother_elem = πTemp009
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(11)            
					if πE = πg.CheckLocal(πF, µother_elem, "other_elem"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, µother_elem, µelem); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label14
					}
					goto Label15
					// line 171: if other_elem == elem:
					πF.SetLineno(171)
				Label14:
					// line 172: cnt_t += 1
					πF.SetLineno(172)
					if πE = πg.CheckLocal(πF, µcnt_t, "cnt_t"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IAdd(πF, µcnt_t, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µcnt_t = πTemp004
					// line 173: t[j] = NULL
					πF.SetLineno(173)
					if πE = πg.CheckLocal(πF, µNULL, "NULL"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µNULL); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					πTemp005 = µj
					if πE = πg.SetItem(πF, µt, πTemp005, πTemp004); πE != nil {
						continue
					}
					goto Label15
				Label15:
					continue
				Label12:
					if πE != nil || πR != nil {
						continue
					}
				Label13:
					if πE = πg.CheckLocal(πF, µcnt_s, "cnt_s"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcnt_t, "cnt_t"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.NE(πF, µcnt_s, µcnt_t); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label16
					}
					goto Label17
					// line 174: if cnt_s != cnt_t:
					πF.SetLineno(174)
				Label16:
					// line 175: diff = _Mismatch(cnt_s, cnt_t, elem)
					πF.SetLineno(175)
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µcnt_s, "cnt_s"); πE != nil {
						continue
					}
					πTemp002[0] = µcnt_s
					if πE = πg.CheckLocal(πF, µcnt_t, "cnt_t"); πE != nil {
						continue
					}
					πTemp002[1] = µcnt_t
					if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
						continue
					}
					πTemp002[2] = µelem
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_Mismatch); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µdiff = πTemp004
					// line 176: result.append(diff)
					πF.SetLineno(176)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdiff, "diff"); πE != nil {
						continue
					}
					πTemp002[0] = µdiff
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
					goto Label17
				Label17:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					πTemp002[0] = µt
					if πTemp003, πE = πg.ResolveGlobal(πF, ßenumerate); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
						continue
					}
					πF.PushCheckpoint(19)
					πTemp006 = false
				Label18:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label20
					}
					if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
							continue
						}
						µi = πTemp004
						µelem = πTemp005
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(18)            
					if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µNULL, "NULL"); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µelem == µNULL).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label21
					}
					goto Label22
					// line 179: if elem is NULL:
					πF.SetLineno(179)
				Label21:
					// line 180: continue
					πF.SetLineno(180)
					continue
					goto Label22
				Label22:
					// line 181: cnt_t = 0
					πF.SetLineno(181)
					µcnt_t = πg.NewInt(0).ToObject()
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp002[0] = µi
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp002[1] = µn
					if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.Iter(πF, πTemp005); πE != nil {
						continue
					}
					πF.PushCheckpoint(24)
					πTemp007 = false
				Label23:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp007 {
						πF.PopCheckpoint()
						goto Label25
					}
					if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
						µj = πTemp004
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(23)            
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					πTemp005 = µj
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, µt, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, πTemp009, µelem); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label26
					}
					goto Label27
					// line 183: if t[j] == elem:
					πF.SetLineno(183)
				Label26:
					// line 184: cnt_t += 1
					πF.SetLineno(184)
					if πE = πg.CheckLocal(πF, µcnt_t, "cnt_t"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IAdd(πF, µcnt_t, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µcnt_t = πTemp004
					// line 185: t[j] = NULL
					πF.SetLineno(185)
					if πE = πg.CheckLocal(πF, µNULL, "NULL"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µNULL); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
						continue
					}
					πTemp005 = µj
					if πE = πg.SetItem(πF, µt, πTemp005, πTemp004); πE != nil {
						continue
					}
					goto Label27
				Label27:
					continue
				Label24:
					if πE != nil || πR != nil {
						continue
					}
				Label25:
					// line 186: diff = _Mismatch(0, cnt_t, elem)
					πF.SetLineno(186)
					πTemp002 = πF.MakeArgs(3)
					πTemp002[0] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µcnt_t, "cnt_t"); πE != nil {
						continue
					}
					πTemp002[1] = µcnt_t
					if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
						continue
					}
					πTemp002[2] = µelem
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_Mismatch); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µdiff = πTemp004
					// line 187: result.append(diff)
					πF.SetLineno(187)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdiff, "diff"); πE != nil {
						continue
					}
					πTemp002[0] = µdiff
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
					continue
				Label19:
					if πE != nil || πR != nil {
						continue
					}
				Label20:
					// line 188: return result
					πF.SetLineno(188)
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
			if πE = πF.Globals().SetItem(πF, ß_count_diff_all_purpose.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 190: def _ordered_count(iterable):
			πF.SetLineno(190)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "iterable", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("_ordered_count", "build/src/__python__/unittest_util.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µiterable *πg.Object = πArgs[0]; _ = µiterable
				var µc *πg.Object = πg.UnboundLocal; _ = µc
				var µelem *πg.Object = πg.UnboundLocal; _ = µelem
				var πTemp001 *πg.Dict
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
					// line 191: 'Return dict of element counts, in the order they were first seen'
					πF.SetLineno(191)
					// line 192: c = {} #OrderedDict()
					πF.SetLineno(192)
					πTemp001 = πg.NewDict()
					πTemp002 = πTemp001.ToObject()
					µc = πTemp002
					if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µiterable); πE != nil {
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
						µelem = πTemp005
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 194: c[elem] = c.get(elem, 0) + 1
					πF.SetLineno(194)
					πTemp006 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
						continue
					}
					πTemp006[0] = µelem
					πTemp006[1] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µc, ßget, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp005, πE = πg.Add(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
						continue
					}
					πTemp008 = µelem
					if πE = πg.SetItem(πF, µc, πTemp008, πTemp007); πE != nil {
						continue
					}
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 195: return c
					πF.SetLineno(195)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πR = µc
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_ordered_count.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 197: def _count_diff_hashable(actual, expected):
			πF.SetLineno(197)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "actual", Def: nil}
			πTemp004[1] = πg.Param{Name: "expected", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("_count_diff_hashable", "build/src/__python__/unittest_util.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µactual *πg.Object = πArgs[0]; _ = µactual
				var µexpected *πg.Object = πArgs[1]; _ = µexpected
				var µs *πg.Object = πg.UnboundLocal; _ = µs
				var µt *πg.Object = πg.UnboundLocal; _ = µt
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µelem *πg.Object = πg.UnboundLocal; _ = µelem
				var µcnt_s *πg.Object = πg.UnboundLocal; _ = µcnt_s
				var µcnt_t *πg.Object = πg.UnboundLocal; _ = µcnt_t
				var µdiff *πg.Object = πg.UnboundLocal; _ = µdiff
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
				var πTemp007 bool
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					case 6: goto Label6
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					// line 198: 'Returns list of (cnt_act, cnt_exp, elem) triples where the counts differ'
					πF.SetLineno(198)
					// line 200: s, t = _ordered_count(actual), _ordered_count(expected)
					πF.SetLineno(200)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µactual, "actual"); πE != nil {
						continue
					}
					πTemp002[0] = µactual
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_ordered_count); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µexpected, "expected"); πE != nil {
						continue
					}
					πTemp002[0] = µexpected
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_ordered_count); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
						continue
					}
					µs = πTemp003
					µt = πTemp004
					// line 201: result = []
					πF.SetLineno(201)
					πTemp002 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp002...).ToObject()
					µresult = πTemp001
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µs, ßitems, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
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
					if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
							continue
						}
						µelem = πTemp004
						µcnt_s = πTemp005
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 203: cnt_t = t.get(elem, 0)
					πF.SetLineno(203)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
						continue
					}
					πTemp002[0] = µelem
					πTemp002[1] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µt, ßget, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µcnt_t = πTemp004
					if πE = πg.CheckLocal(πF, µcnt_s, "cnt_s"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcnt_t, "cnt_t"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.NE(πF, µcnt_s, µcnt_t); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label4
					}
					goto Label5
					// line 204: if cnt_s != cnt_t:
					πF.SetLineno(204)
				Label4:
					// line 205: diff = _Mismatch(cnt_s, cnt_t, elem)
					πF.SetLineno(205)
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µcnt_s, "cnt_s"); πE != nil {
						continue
					}
					πTemp002[0] = µcnt_s
					if πE = πg.CheckLocal(πF, µcnt_t, "cnt_t"); πE != nil {
						continue
					}
					πTemp002[1] = µcnt_t
					if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
						continue
					}
					πTemp002[2] = µelem
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_Mismatch); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µdiff = πTemp004
					// line 206: result.append(diff)
					πF.SetLineno(206)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdiff, "diff"); πE != nil {
						continue
					}
					πTemp002[0] = µdiff
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
					goto Label5
				Label5:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µt, ßitems, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
						continue
					}
					πF.PushCheckpoint(7)
					πTemp006 = false
				Label6:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label8
					}
					if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
							continue
						}
						µelem = πTemp004
						µcnt_t = πTemp005
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(6)            
					if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Contains(πF, µs, µelem); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label9
					}
					goto Label10
					// line 208: if elem not in s:
					πF.SetLineno(208)
				Label9:
					// line 209: diff = _Mismatch(0, cnt_t, elem)
					πF.SetLineno(209)
					πTemp002 = πF.MakeArgs(3)
					πTemp002[0] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µcnt_t, "cnt_t"); πE != nil {
						continue
					}
					πTemp002[1] = µcnt_t
					if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
						continue
					}
					πTemp002[2] = µelem
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_Mismatch); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µdiff = πTemp004
					// line 210: result.append(diff)
					πF.SetLineno(210)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdiff, "diff"); πE != nil {
						continue
					}
					πTemp002[0] = µdiff
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
					goto Label10
				Label10:
					continue
				Label7:
					if πE != nil || πR != nil {
						continue
					}
				Label8:
					// line 211: return result
					πF.SetLineno(211)
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
			if πE = πF.Globals().SetItem(πF, ß_count_diff_hashable.ToObject(), πTemp010); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("unittest_util", Code)
}
