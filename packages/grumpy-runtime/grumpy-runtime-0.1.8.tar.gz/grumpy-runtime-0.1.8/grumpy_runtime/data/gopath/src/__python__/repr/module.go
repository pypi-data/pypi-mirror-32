package repr
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/repr.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßException := πg.InternStr("Exception")
		ßRepr := πg.InternStr("Repr")
		ß_ := πg.InternStr("_")
		ß__all__ := πg.InternStr("__all__")
		ß__builtin__ := πg.InternStr("__builtin__")
		ß__class__ := πg.InternStr("__class__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_possibly_sorted := πg.InternStr("_possibly_sorted")
		ß_repr_iterable := πg.InternStr("_repr_iterable")
		ßaRepr := πg.InternStr("aRepr")
		ßappend := πg.InternStr("append")
		ßgetattr := πg.InternStr("getattr")
		ßhasattr := πg.InternStr("hasattr")
		ßid := πg.InternStr("id")
		ßislice := πg.InternStr("islice")
		ßitertools := πg.InternStr("itertools")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßlist := πg.InternStr("list")
		ßmax := πg.InternStr("max")
		ßmaxarray := πg.InternStr("maxarray")
		ßmaxdeque := πg.InternStr("maxdeque")
		ßmaxdict := πg.InternStr("maxdict")
		ßmaxfrozenset := πg.InternStr("maxfrozenset")
		ßmaxlevel := πg.InternStr("maxlevel")
		ßmaxlist := πg.InternStr("maxlist")
		ßmaxlong := πg.InternStr("maxlong")
		ßmaxother := πg.InternStr("maxother")
		ßmaxset := πg.InternStr("maxset")
		ßmaxstring := πg.InternStr("maxstring")
		ßmaxtuple := πg.InternStr("maxtuple")
		ßobject := πg.InternStr("object")
		ßrepr := πg.InternStr("repr")
		ßrepr1 := πg.InternStr("repr1")
		ßrepr_ := πg.InternStr("repr_")
		ßrepr_array := πg.InternStr("repr_array")
		ßrepr_deque := πg.InternStr("repr_deque")
		ßrepr_dict := πg.InternStr("repr_dict")
		ßrepr_frozenset := πg.InternStr("repr_frozenset")
		ßrepr_instance := πg.InternStr("repr_instance")
		ßrepr_list := πg.InternStr("repr_list")
		ßrepr_long := πg.InternStr("repr_long")
		ßrepr_set := πg.InternStr("repr_set")
		ßrepr_str := πg.InternStr("repr_str")
		ßrepr_tuple := πg.InternStr("repr_tuple")
		ßsorted := πg.InternStr("sorted")
		ßsplit := πg.InternStr("split")
		ßtype := πg.InternStr("type")
		ßtypecode := πg.InternStr("typecode")
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """Redo the builtin repr() (representation) but with limits on most sizes."""
			πF.SetLineno(1)
			// line 3: __all__ = ["Repr","repr"]
			πF.SetLineno(3)
			πTemp001 = make([]*πg.Object, 2)
			πTemp001[0] = ßRepr.ToObject()
			πTemp001[1] = ßrepr.ToObject()
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 5: import __builtin__
			πF.SetLineno(5)
			if πTemp001, πE = πg.ImportModule(πF, "__builtin__"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ß__builtin__.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 6: import itertools
			πF.SetLineno(6)
			if πTemp001, πE = πg.ImportModule(πF, "itertools"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßitertools.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 8: class Repr(object):
			πF.SetLineno(8)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
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
			_, πE = πg.NewCode("Repr", "build/src/__python__/repr.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp012 *πg.Object
				_ = πTemp012
				var πTemp013 *πg.Object
				_ = πTemp013
				var πTemp014 *πg.Object
				_ = πTemp014
				var πTemp015 *πg.Object
				_ = πTemp015
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 10: def __init__(self):
					πF.SetLineno(10)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/repr.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 11: self.maxlevel = 6
							πF.SetLineno(11)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(6).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmaxlevel, πTemp001); πE != nil {
								continue
							}
							// line 12: self.maxtuple = 6
							πF.SetLineno(12)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(6).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmaxtuple, πTemp001); πE != nil {
								continue
							}
							// line 13: self.maxlist = 6
							πF.SetLineno(13)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(6).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmaxlist, πTemp001); πE != nil {
								continue
							}
							// line 14: self.maxarray = 5
							πF.SetLineno(14)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(5).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmaxarray, πTemp001); πE != nil {
								continue
							}
							// line 15: self.maxdict = 4
							πF.SetLineno(15)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmaxdict, πTemp001); πE != nil {
								continue
							}
							// line 16: self.maxset = 6
							πF.SetLineno(16)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(6).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmaxset, πTemp001); πE != nil {
								continue
							}
							// line 17: self.maxfrozenset = 6
							πF.SetLineno(17)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(6).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmaxfrozenset, πTemp001); πE != nil {
								continue
							}
							// line 18: self.maxdeque = 6
							πF.SetLineno(18)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(6).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmaxdeque, πTemp001); πE != nil {
								continue
							}
							// line 19: self.maxstring = 30
							πF.SetLineno(19)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(30).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmaxstring, πTemp001); πE != nil {
								continue
							}
							// line 20: self.maxlong = 40
							πF.SetLineno(20)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(40).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmaxlong, πTemp001); πE != nil {
								continue
							}
							// line 21: self.maxother = 20
							πF.SetLineno(21)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(20).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmaxother, πTemp001); πE != nil {
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
					// line 23: def repr(self, x):
					πF.SetLineno(23)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "x", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("repr", "build/src/__python__/repr.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µx *πg.Object = πArgs[1]; _ = µx
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
							// line 24: return self.repr1(x, self.maxlevel)
							πF.SetLineno(24)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmaxlevel, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrepr1, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßrepr.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 26: def repr1(self, x, level):
					πF.SetLineno(26)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "x", Def: nil}
					πTemp002[2] = πg.Param{Name: "level", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("repr1", "build/src/__python__/repr.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µx *πg.Object = πArgs[1]; _ = µx
						var µlevel *πg.Object = πArgs[2]; _ = µlevel
						var µtypename *πg.Object = πg.UnboundLocal; _ = µtypename
						var µparts *πg.Object = πg.UnboundLocal; _ = µparts
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µj *πg.Object = πg.UnboundLocal; _ = µj
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
							default: panic("unexpected function state")
							}
							// line 27: typename = type(x).__name__
							πF.SetLineno(27)
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
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ß__name__, nil); πE != nil {
								continue
							}
							µtypename = πTemp002
							if πE = πg.CheckLocal(πF, µtypename, "typename"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, µtypename, πg.NewStr(" ").ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 28: if ' ' in typename:
							πF.SetLineno(28)
						Label1:
							// line 29: parts = typename.split()
							πF.SetLineno(29)
							if πE = πg.CheckLocal(πF, µtypename, "typename"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtypename, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µparts = πTemp003
							// line 30: typename = '_'.join(parts)
							πF.SetLineno(30)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
								continue
							}
							πTemp001[0] = µparts
							if πTemp002, πE = πg.GetAttr(πF, ß_.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtypename = πTemp003
							goto Label2
						Label2:
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µtypename, "typename"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, ßrepr_.ToObject(), µtypename); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
								goto Label3
							}
							goto Label4
							// line 31: if hasattr(self, 'repr_' + typename):
							πF.SetLineno(31)
						Label3:
							// line 32: return getattr(self, 'repr_' + typename)(x, level)
							πF.SetLineno(32)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							πTemp001[1] = µlevel
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[0] = µself
							if πE = πg.CheckLocal(πF, µtypename, "typename"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, ßrepr_.ToObject(), µtypename); πE != nil {
								continue
							}
							πTemp005[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp002
							continue
							goto Label5
						Label4:
							// line 34: s = __builtin__.repr(x)
							πF.SetLineno(34)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							if πTemp002, πE = πg.ResolveGlobal(πF, ß__builtin__); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrepr, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp002
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmaxother, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GT(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							goto Label7
							// line 35: if len(s) > self.maxother:
							πF.SetLineno(35)
						Label6:
							// line 36: i = max(0, (self.maxother-3)//2)
							πF.SetLineno(36)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßmaxother, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, πTemp006, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.FloorDiv(πF, πTemp003, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µi = πTemp003
							// line 37: j = max(0, self.maxother-3-i)
							πF.SetLineno(37)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßmaxother, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, πTemp006, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, πTemp003, µi); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µj = πTemp003
							// line 38: s = s[:i] + '...' + s[len(s)-j:]
							πF.SetLineno(38)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µi, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µs, πTemp006); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp007, πg.NewStr("...").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Sub(πF, πTemp009, µj); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πTemp007, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µs, πTemp006); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp007); πE != nil {
								continue
							}
							µs = πTemp002
							goto Label7
						Label7:
							// line 39: return s
							πF.SetLineno(39)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πR = µs
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
					if πE = πClass.SetItem(πF, ßrepr1.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 41: def _repr_iterable(self, x, level, left, right, maxiter, trail=''):
					πF.SetLineno(41)
					πTemp002 = make([]πg.Param, 7)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "x", Def: nil}
					πTemp002[2] = πg.Param{Name: "level", Def: nil}
					πTemp002[3] = πg.Param{Name: "left", Def: nil}
					πTemp002[4] = πg.Param{Name: "right", Def: nil}
					πTemp002[5] = πg.Param{Name: "maxiter", Def: nil}
					πTemp002[6] = πg.Param{Name: "trail", Def: ß.ToObject()}
					πTemp005 = πg.NewFunction(πg.NewCode("_repr_iterable", "build/src/__python__/repr.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µx *πg.Object = πArgs[1]; _ = µx
						var µlevel *πg.Object = πArgs[2]; _ = µlevel
						var µleft *πg.Object = πArgs[3]; _ = µleft
						var µright *πg.Object = πArgs[4]; _ = µright
						var µmaxiter *πg.Object = πArgs[5]; _ = µmaxiter
						var µtrail *πg.Object = πArgs[6]; _ = µtrail
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µnewlevel *πg.Object = πg.UnboundLocal; _ = µnewlevel
						var µrepr1 *πg.Object = πg.UnboundLocal; _ = µrepr1
						var µpieces *πg.Object = πg.UnboundLocal; _ = µpieces
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 []πg.Param
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
							// line 42: n = len(x)
							πF.SetLineno(42)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µn = πTemp003
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LE(πF, µlevel, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp002 = µn
						Label1:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							goto Label3
							// line 43: if level <= 0 and n:
							πF.SetLineno(43)
						Label2:
							// line 44: s = '...'
							πF.SetLineno(44)
							µs = πg.NewStr("...").ToObject()
							goto Label4
						Label3:
							// line 46: newlevel = level - 1
							πF.SetLineno(46)
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, µlevel, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µnewlevel = πTemp002
							// line 47: repr1 = self.repr1
							πF.SetLineno(47)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrepr1, nil); πE != nil {
								continue
							}
							µrepr1 = πTemp002
							// line 48: pieces = [repr1(elem, newlevel) for elem in itertools.islice(x, maxiter)]
							πF.SetLineno(48)
							πTemp005 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/repr.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µelem *πg.Object = πg.UnboundLocal; _ = µelem
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
								var πTemp006 bool
								_ = πTemp006
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
									for ; πF.State() >= 0; πF.PopCheckpoint() {
										switch πF.State() {
										case 0:
										case 1: goto Label1
										case 2: goto Label2
										case 4: goto Label4
										default: panic("unexpected function state")
										}
										πTemp002 = πF.MakeArgs(2)
										if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
											continue
										}
										πTemp002[0] = µx
										if πE = πg.CheckLocal(πF, µmaxiter, "maxiter"); πE != nil {
											continue
										}
										πTemp002[1] = µmaxiter
										if πTemp003, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
											continue
										}
										if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßislice, nil); πE != nil {
											continue
										}
										if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp002)
										if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
											continue
										}
										πF.PushCheckpoint(2)
										πTemp005 = false
									Label1:
										if πE != nil || πR != nil {
											continue
										}
										if πTemp005 {
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
											πTemp006 = !isStop
										} else {
											πTemp006 = true
											µelem = πTemp003
										}
										if πE != nil || !πTemp006 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 48: pieces = [repr1(elem, newlevel) for elem in itertools.islice(x, maxiter)]
										πF.SetLineno(48)
										πTemp002 = πF.MakeArgs(2)
										if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
											continue
										}
										πTemp002[0] = µelem
										if πE = πg.CheckLocal(πF, µnewlevel, "newlevel"); πE != nil {
											continue
										}
										πTemp002[1] = µnewlevel
										if πE = πg.CheckLocal(πF, µrepr1, "repr1"); πE != nil {
											continue
										}
										if πTemp003, πE = µrepr1.Call(πF, πTemp002, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp002)
										πF.PushCheckpoint(4)
										return πTemp003, nil
									Label4:
										πTemp004 = πSent
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
							if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp006}, nil); πE != nil {
								continue
							}
							µpieces = πTemp002
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmaxiter, "maxiter"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GT(πF, µn, µmaxiter); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							goto Label6
							// line 49: if n > maxiter:  pieces.append('...')
							πF.SetLineno(49)
						Label5:
							// line 49: if n > maxiter:  pieces.append('...')
							πF.SetLineno(49)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("...").ToObject()
							if πE = πg.CheckLocal(πF, µpieces, "pieces"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µpieces, ßappend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label6
						Label6:
							// line 50: s = ', '.join(pieces)
							πF.SetLineno(50)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpieces, "pieces"); πE != nil {
								continue
							}
							πTemp001[0] = µpieces
							if πTemp002, πE = πg.GetAttr(πF, πg.NewStr(", ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp006
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Eq(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp006
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µtrail, "trail"); πE != nil {
								continue
							}
							πTemp002 = µtrail
						Label7:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							goto Label9
							// line 51: if n == 1 and trail:  right = trail + right
							πF.SetLineno(51)
						Label8:
							// line 51: if n == 1 and trail:  right = trail + right
							πF.SetLineno(51)
							if πE = πg.CheckLocal(πF, µtrail, "trail"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µright, "right"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µtrail, µright); πE != nil {
								continue
							}
							µright = πTemp002
							goto Label9
						Label9:
							goto Label4
						Label4:
							// line 52: return '%s%s%s' % (left, s, right)
							πF.SetLineno(52)
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µright, "right"); πE != nil {
								continue
							}
							πTemp006 = πg.NewTuple3(µleft, µs, µright).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s%s%s").ToObject(), πTemp006); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_repr_iterable.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 54: def repr_tuple(self, x, level):
					πF.SetLineno(54)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "x", Def: nil}
					πTemp002[2] = πg.Param{Name: "level", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("repr_tuple", "build/src/__python__/repr.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µx *πg.Object = πArgs[1]; _ = µx
						var µlevel *πg.Object = πArgs[2]; _ = µlevel
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
							// line 55: return self._repr_iterable(x, level, '(', ')', self.maxtuple, ',')
							πF.SetLineno(55)
							πTemp001 = πF.MakeArgs(6)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							πTemp001[1] = µlevel
							πTemp001[2] = πg.NewStr("(").ToObject()
							πTemp001[3] = πg.NewStr(")").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmaxtuple, nil); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							πTemp001[5] = πg.NewStr(",").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_repr_iterable, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßrepr_tuple.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 57: def repr_list(self, x, level):
					πF.SetLineno(57)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "x", Def: nil}
					πTemp002[2] = πg.Param{Name: "level", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("repr_list", "build/src/__python__/repr.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µx *πg.Object = πArgs[1]; _ = µx
						var µlevel *πg.Object = πArgs[2]; _ = µlevel
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
							// line 58: return self._repr_iterable(x, level, '[', ']', self.maxlist)
							πF.SetLineno(58)
							πTemp001 = πF.MakeArgs(5)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							πTemp001[1] = µlevel
							πTemp001[2] = πg.NewStr("[").ToObject()
							πTemp001[3] = πg.NewStr("]").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmaxlist, nil); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_repr_iterable, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßrepr_list.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 60: def repr_array(self, x, level):
					πF.SetLineno(60)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "x", Def: nil}
					πTemp002[2] = πg.Param{Name: "level", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("repr_array", "build/src/__python__/repr.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µx *πg.Object = πArgs[1]; _ = µx
						var µlevel *πg.Object = πArgs[2]; _ = µlevel
						var µheader *πg.Object = πg.UnboundLocal; _ = µheader
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 61: header = "array('%s', [" % x.typecode
							πF.SetLineno(61)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µx, ßtypecode, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("array('%s', [").ToObject(), πTemp002); πE != nil {
								continue
							}
							µheader = πTemp001
							// line 62: return self._repr_iterable(x, level, header, '])', self.maxarray)
							πF.SetLineno(62)
							πTemp003 = πF.MakeArgs(5)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp003[0] = µx
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							πTemp003[1] = µlevel
							if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
								continue
							}
							πTemp003[2] = µheader
							πTemp003[3] = πg.NewStr("])").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmaxarray, nil); πE != nil {
								continue
							}
							πTemp003[4] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_repr_iterable, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
					if πE = πClass.SetItem(πF, ßrepr_array.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 64: def repr_set(self, x, level):
					πF.SetLineno(64)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "x", Def: nil}
					πTemp002[2] = πg.Param{Name: "level", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("repr_set", "build/src/__python__/repr.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µx *πg.Object = πArgs[1]; _ = µx
						var µlevel *πg.Object = πArgs[2]; _ = µlevel
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
							// line 65: x = _possibly_sorted(x)
							πF.SetLineno(65)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_possibly_sorted); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µx = πTemp003
							// line 66: return self._repr_iterable(x, level, 'set([', '])', self.maxset)
							πF.SetLineno(66)
							πTemp001 = πF.MakeArgs(5)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							πTemp001[1] = µlevel
							πTemp001[2] = πg.NewStr("set([").ToObject()
							πTemp001[3] = πg.NewStr("])").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmaxset, nil); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_repr_iterable, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßrepr_set.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 68: def repr_frozenset(self, x, level):
					πF.SetLineno(68)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "x", Def: nil}
					πTemp002[2] = πg.Param{Name: "level", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("repr_frozenset", "build/src/__python__/repr.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µx *πg.Object = πArgs[1]; _ = µx
						var µlevel *πg.Object = πArgs[2]; _ = µlevel
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
							// line 69: x = _possibly_sorted(x)
							πF.SetLineno(69)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_possibly_sorted); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µx = πTemp003
							// line 70: return self._repr_iterable(x, level, 'frozenset([', '])',
							πF.SetLineno(70)
							πTemp001 = πF.MakeArgs(5)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							πTemp001[1] = µlevel
							πTemp001[2] = πg.NewStr("frozenset([").ToObject()
							πTemp001[3] = πg.NewStr("])").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmaxfrozenset, nil); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_repr_iterable, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßrepr_frozenset.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 73: def repr_deque(self, x, level):
					πF.SetLineno(73)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "x", Def: nil}
					πTemp002[2] = πg.Param{Name: "level", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("repr_deque", "build/src/__python__/repr.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µx *πg.Object = πArgs[1]; _ = µx
						var µlevel *πg.Object = πArgs[2]; _ = µlevel
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
							// line 74: return self._repr_iterable(x, level, 'deque([', '])', self.maxdeque)
							πF.SetLineno(74)
							πTemp001 = πF.MakeArgs(5)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							πTemp001[1] = µlevel
							πTemp001[2] = πg.NewStr("deque([").ToObject()
							πTemp001[3] = πg.NewStr("])").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmaxdeque, nil); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_repr_iterable, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßrepr_deque.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 76: def repr_dict(self, x, level):
					πF.SetLineno(76)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "x", Def: nil}
					πTemp002[2] = πg.Param{Name: "level", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("repr_dict", "build/src/__python__/repr.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µx *πg.Object = πArgs[1]; _ = µx
						var µlevel *πg.Object = πArgs[2]; _ = µlevel
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var µnewlevel *πg.Object = πg.UnboundLocal; _ = µnewlevel
						var µrepr1 *πg.Object = πg.UnboundLocal; _ = µrepr1
						var µpieces *πg.Object = πg.UnboundLocal; _ = µpieces
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µkeyrepr *πg.Object = πg.UnboundLocal; _ = µkeyrepr
						var µvalrepr *πg.Object = πg.UnboundLocal; _ = µvalrepr
						var µs *πg.Object = πg.UnboundLocal; _ = µs
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 5: goto Label5
							case 6: goto Label6
							default: panic("unexpected function state")
							}
							// line 77: n = len(x)
							πF.SetLineno(77)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µn = πTemp003
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µn, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 78: if n == 0: return '{}'
							πF.SetLineno(78)
						Label1:
							// line 78: if n == 0: return '{}'
							πF.SetLineno(78)
							πR = πg.NewStr("{}").ToObject()
							continue
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LE(πF, µlevel, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 79: if level <= 0: return '{...}'
							πF.SetLineno(79)
						Label3:
							// line 79: if level <= 0: return '{...}'
							πF.SetLineno(79)
							πR = πg.NewStr("{...}").ToObject()
							continue
							goto Label4
						Label4:
							// line 80: newlevel = level - 1
							πF.SetLineno(80)
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, µlevel, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µnewlevel = πTemp002
							// line 81: repr1 = self.repr1
							πF.SetLineno(81)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrepr1, nil); πE != nil {
								continue
							}
							µrepr1 = πTemp002
							// line 82: pieces = []
							πF.SetLineno(82)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µpieces = πTemp002
							πTemp001 = πF.MakeArgs(2)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp005[0] = µx
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_possibly_sorted); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmaxdict, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßislice, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
								continue
							}
							πF.PushCheckpoint(6)
							πTemp004 = false
						Label5:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label7
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
								µkey = πTemp003
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(5)            
							// line 84: keyrepr = repr1(key, newlevel)
							πF.SetLineno(84)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001[0] = µkey
							if πE = πg.CheckLocal(πF, µnewlevel, "newlevel"); πE != nil {
								continue
							}
							πTemp001[1] = µnewlevel
							if πE = πg.CheckLocal(πF, µrepr1, "repr1"); πE != nil {
								continue
							}
							if πTemp003, πE = µrepr1.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µkeyrepr = πTemp003
							// line 85: valrepr = repr1(x[key], newlevel)
							πF.SetLineno(85)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp003 = µkey
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µx, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							if πE = πg.CheckLocal(πF, µnewlevel, "newlevel"); πE != nil {
								continue
							}
							πTemp001[1] = µnewlevel
							if πE = πg.CheckLocal(πF, µrepr1, "repr1"); πE != nil {
								continue
							}
							if πTemp003, πE = µrepr1.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µvalrepr = πTemp003
							// line 86: pieces.append('%s: %s' % (keyrepr, valrepr))
							πF.SetLineno(86)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkeyrepr, "keyrepr"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalrepr, "valrepr"); πE != nil {
								continue
							}
							πTemp006 = πg.NewTuple2(µkeyrepr, µvalrepr).ToObject()
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s: %s").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µpieces, "pieces"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µpieces, ßappend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label6:
							if πE != nil || πR != nil {
								continue
							}
						Label7:
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmaxdict, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GT(πF, µn, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							goto Label9
							// line 87: if n > self.maxdict: pieces.append('...')
							πF.SetLineno(87)
						Label8:
							// line 87: if n > self.maxdict: pieces.append('...')
							πF.SetLineno(87)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("...").ToObject()
							if πE = πg.CheckLocal(πF, µpieces, "pieces"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µpieces, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label9
						Label9:
							// line 88: s = ', '.join(pieces)
							πF.SetLineno(88)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpieces, "pieces"); πE != nil {
								continue
							}
							πTemp001[0] = µpieces
							if πTemp002, πE = πg.GetAttr(πF, πg.NewStr(", ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp003
							// line 89: return '{%s}' % (s,)
							πF.SetLineno(89)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple1(µs).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("{%s}").ToObject(), πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßrepr_dict.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 91: def repr_str(self, x, level):
					πF.SetLineno(91)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "x", Def: nil}
					πTemp002[2] = πg.Param{Name: "level", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("repr_str", "build/src/__python__/repr.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µx *πg.Object = πArgs[1]; _ = µx
						var µlevel *πg.Object = πArgs[2]; _ = µlevel
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µj *πg.Object = πg.UnboundLocal; _ = µj
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
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
							default: panic("unexpected function state")
							}
							// line 92: s = __builtin__.repr(x[:self.maxstring])
							πF.SetLineno(92)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmaxstring, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp003, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µx, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ß__builtin__); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrepr, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp002
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmaxstring, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GT(πF, πTemp004, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 93: if len(s) > self.maxstring:
							πF.SetLineno(93)
						Label1:
							// line 94: i = max(0, (self.maxstring-3)//2)
							πF.SetLineno(94)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßmaxstring, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, πTemp004, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.FloorDiv(πF, πTemp003, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µi = πTemp003
							// line 95: j = max(0, self.maxstring-3-i)
							πF.SetLineno(95)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßmaxstring, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, πTemp004, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, πTemp003, µi); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µj = πTemp003
							// line 96: s = __builtin__.repr(x[:i] + x[len(x)-j:])
							πF.SetLineno(96)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µi, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µx, πTemp003); πE != nil {
								continue
							}
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp007[0] = µx
							if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Sub(πF, πTemp009, µj); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πTemp006, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µx, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß__builtin__); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrepr, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp002
							// line 97: s = s[:i] + '...' + s[len(s)-j:]
							πF.SetLineno(97)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µi, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µs, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp006, πg.NewStr("...").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Sub(πF, πTemp009, µj); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πTemp006, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µs, πTemp004); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp006); πE != nil {
								continue
							}
							µs = πTemp002
							goto Label2
						Label2:
							// line 98: return s
							πF.SetLineno(98)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πR = µs
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßrepr_str.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 100: def repr_long(self, x, level):
					πF.SetLineno(100)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "x", Def: nil}
					πTemp002[2] = πg.Param{Name: "level", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("repr_long", "build/src/__python__/repr.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µx *πg.Object = πArgs[1]; _ = µx
						var µlevel *πg.Object = πArgs[2]; _ = µlevel
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µj *πg.Object = πg.UnboundLocal; _ = µj
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 *πg.Object
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
							// line 101: s = __builtin__.repr(x) # XXX Hope this isn't too slow...
							πF.SetLineno(101)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							if πTemp002, πE = πg.ResolveGlobal(πF, ß__builtin__); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrepr, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp002
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmaxlong, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GT(πF, πTemp004, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 102: if len(s) > self.maxlong:
							πF.SetLineno(102)
						Label1:
							// line 103: i = max(0, (self.maxlong-3)//2)
							πF.SetLineno(103)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßmaxlong, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, πTemp004, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.FloorDiv(πF, πTemp003, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µi = πTemp003
							// line 104: j = max(0, self.maxlong-3-i)
							πF.SetLineno(104)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßmaxlong, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, πTemp004, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, πTemp003, µi); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µj = πTemp003
							// line 105: s = s[:i] + '...' + s[len(s)-j:]
							πF.SetLineno(105)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µi, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µs, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp006, πg.NewStr("...").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Sub(πF, πTemp008, µj); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πTemp006, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µs, πTemp004); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp006); πE != nil {
								continue
							}
							µs = πTemp002
							goto Label2
						Label2:
							// line 106: return s
							πF.SetLineno(106)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πR = µs
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßrepr_long.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 108: def repr_instance(self, x, level):
					πF.SetLineno(108)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "x", Def: nil}
					πTemp002[2] = πg.Param{Name: "level", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("repr_instance", "build/src/__python__/repr.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µx *πg.Object = πArgs[1]; _ = µx
						var µlevel *πg.Object = πArgs[2]; _ = µlevel
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µj *πg.Object = πg.UnboundLocal; _ = µj
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
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
							// line 109: try:
							πF.SetLineno(109)
							πF.PushCheckpoint(2)
							// line 110: s = __builtin__.repr(x)
							πF.SetLineno(110)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							if πTemp002, πE = πg.ResolveGlobal(πF, ß__builtin__); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrepr, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp002
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
							// line 113: except Exception:
							πF.SetLineno(113)
						Label3:
							// line 114: return '<%s instance at %x>' % (x.__class__.__name__, id(x))
							πF.SetLineno(114)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µx, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ß__name__, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							if πTemp007, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp003 = πg.NewTuple2(πTemp008, πTemp009).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<%s instance at %x>").ToObject(), πTemp003); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmaxstring, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GT(πF, πTemp007, πTemp003); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 115: if len(s) > self.maxstring:
							πF.SetLineno(115)
						Label4:
							// line 116: i = max(0, (self.maxstring-3)//2)
							πF.SetLineno(116)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßmaxstring, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, πTemp007, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.FloorDiv(πF, πTemp003, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µi = πTemp003
							// line 117: j = max(0, self.maxstring-3-i)
							πF.SetLineno(117)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßmaxstring, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, πTemp007, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, πTemp003, µi); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µj = πTemp003
							// line 118: s = s[:i] + '...' + s[len(s)-j:]
							πF.SetLineno(118)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µi, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µs, πTemp007); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp008, πg.NewStr("...").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πTemp009, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Sub(πF, πTemp010, µj); πE != nil {
								continue
							}
							if πTemp007, πE = πg.SliceType.Call(πF, πg.Args{πTemp008, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µs, πTemp007); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp008); πE != nil {
								continue
							}
							µs = πTemp002
							goto Label5
						Label5:
							// line 119: return s
							πF.SetLineno(119)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πR = µs
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßrepr_instance.ToObject(), πTemp015); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("Repr").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßRepr.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 122: def _possibly_sorted(x):
			πF.SetLineno(122)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "x", Def: nil}
			πTemp002 = πg.NewFunction(πg.NewCode("_possibly_sorted", "build/src/__python__/repr.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 126: try:
					πF.SetLineno(126)
					πF.PushCheckpoint(2)
					// line 127: return sorted(x)
					πF.SetLineno(127)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
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
					// line 128: except Exception:
					πF.SetLineno(128)
				Label3:
					// line 129: return list(x)
					πF.SetLineno(129)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
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
			if πE = πF.Globals().SetItem(πF, ß_possibly_sorted.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 131: aRepr = Repr()
			πF.SetLineno(131)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßRepr); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßaRepr.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 132: repr = aRepr.repr
			πF.SetLineno(132)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßaRepr); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßrepr, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßrepr.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("repr", Code)
}
