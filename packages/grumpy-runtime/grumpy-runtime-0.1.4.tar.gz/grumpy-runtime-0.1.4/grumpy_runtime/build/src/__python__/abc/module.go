package abc
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/abc.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßABCMeta := πg.InternStr("ABCMeta")
		ßFalse := πg.InternStr("False")
		ßNone := πg.InternStr("None")
		ßNotImplemented := πg.InternStr("NotImplemented")
		ßRuntimeError := πg.InternStr("RuntimeError")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßWeakSet := πg.InternStr("WeakSet")
		ß__abstractmethods__ := πg.InternStr("__abstractmethods__")
		ß__class__ := πg.InternStr("__class__")
		ß__dict__ := πg.InternStr("__dict__")
		ß__instancecheck__ := πg.InternStr("__instancecheck__")
		ß__isabstractmethod__ := πg.InternStr("__isabstractmethod__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__mro__ := πg.InternStr("__mro__")
		ß__name__ := πg.InternStr("__name__")
		ß__new__ := πg.InternStr("__new__")
		ß__subclasscheck__ := πg.InternStr("__subclasscheck__")
		ß__subclasses__ := πg.InternStr("__subclasses__")
		ß__subclasshook__ := πg.InternStr("__subclasshook__")
		ß_abc_ := πg.InternStr("_abc_")
		ß_abc_cache := πg.InternStr("_abc_cache")
		ß_abc_invalidation_counter := πg.InternStr("_abc_invalidation_counter")
		ß_abc_negative_cache := πg.InternStr("_abc_negative_cache")
		ß_abc_negative_cache_version := πg.InternStr("_abc_negative_cache_version")
		ß_abc_registry := πg.InternStr("_abc_registry")
		ß_dump_registry := πg.InternStr("_dump_registry")
		ß_weakrefset := πg.InternStr("_weakrefset")
		ßabstractmethod := πg.InternStr("abstractmethod")
		ßabstractproperty := πg.InternStr("abstractproperty")
		ßadd := πg.InternStr("add")
		ßbool := πg.InternStr("bool")
		ßfrozenset := πg.InternStr("frozenset")
		ßgetattr := πg.InternStr("getattr")
		ßisinstance := πg.InternStr("isinstance")
		ßissubclass := πg.InternStr("issubclass")
		ßitems := πg.InternStr("items")
		ßkeys := πg.InternStr("keys")
		ßproperty := πg.InternStr("property")
		ßregister := πg.InternStr("register")
		ßset := πg.InternStr("set")
		ßsorted := πg.InternStr("sorted")
		ßstartswith := πg.InternStr("startswith")
		ßsuper := πg.InternStr("super")
		ßtype := πg.InternStr("type")
		ßtypes := πg.InternStr("types")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 []πg.Param
		_ = πTemp003
		var πTemp004 *πg.Dict
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
			// line 4: """Abstract Base Classes (ABCs) according to PEP 3119."""
			πF.SetLineno(4)
			// line 6: import _weakrefset
			πF.SetLineno(6)
			if πTemp002, πE = πg.ImportModule(πF, "_weakrefset"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_weakrefset.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 7: import types
			πF.SetLineno(7)
			if πTemp002, πE = πg.ImportModule(πF, "types"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtypes.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 15: def abstractmethod(funcobj):
			πF.SetLineno(15)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "funcobj", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("abstractmethod", "build/src/__python__/abc.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfuncobj *πg.Object = πArgs[0]; _ = µfuncobj
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
					// line 16: """A decorator indicating abstract methods.
					πF.SetLineno(16)
					// line 32: funcobj.__isabstractmethod__ = True
					πF.SetLineno(32)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfuncobj, "funcobj"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µfuncobj, ß__isabstractmethod__, πTemp002); πE != nil {
						continue
					}
					// line 33: return funcobj
					πF.SetLineno(33)
					if πE = πg.CheckLocal(πF, µfuncobj, "funcobj"); πE != nil {
						continue
					}
					πR = µfuncobj
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßabstractmethod.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 36: class abstractproperty(property):
			πF.SetLineno(36)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßproperty); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp004 = πg.NewDict()
			if πTemp005, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp005); πE != nil {
				continue
			}
			_, πE = πg.NewCode("abstractproperty", "build/src/__python__/abc.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 37: """A decorator indicating abstract properties.
					πF.SetLineno(37)
					// line 62: __isabstractmethod__ = True
					πF.SetLineno(62)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__isabstractmethod__.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("abstractproperty").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßabstractproperty.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 65: class ABCMeta(type):
			πF.SetLineno(65)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp004 = πg.NewDict()
			if πTemp005, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp005); πE != nil {
				continue
			}
			_, πE = πg.NewCode("ABCMeta", "build/src/__python__/abc.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 67: """Metaclass for defining Abstract Base Classes (ABCs).
					πF.SetLineno(67)
					// line 84: _abc_invalidation_counter = 0
					πF.SetLineno(84)
					if πE = πClass.SetItem(πF, ß_abc_invalidation_counter.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 86: def __new__(mcls, name, bases, namespace):
					πF.SetLineno(86)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "mcls", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp002[2] = πg.Param{Name: "bases", Def: nil}
					πTemp002[3] = πg.Param{Name: "namespace", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__new__", "build/src/__python__/abc.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µmcls *πg.Object = πArgs[0]; _ = µmcls
						var µname *πg.Object = πArgs[1]; _ = µname
						var µbases *πg.Object = πArgs[2]; _ = µbases
						var µnamespace *πg.Object = πArgs[3]; _ = µnamespace
						var µcls *πg.Object = πg.UnboundLocal; _ = µcls
						var µabstracts *πg.Object = πg.UnboundLocal; _ = µabstracts
						var µbase *πg.Object = πg.UnboundLocal; _ = µbase
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []πg.Param
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
						var πTemp011 bool
						_ = πTemp011
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
							// line 87: cls = super(ABCMeta, mcls).__new__(mcls, name, bases, namespace)
							πF.SetLineno(87)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µmcls, "mcls"); πE != nil {
								continue
							}
							πTemp001[0] = µmcls
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001[1] = µname
							if πE = πg.CheckLocal(πF, µbases, "bases"); πE != nil {
								continue
							}
							πTemp001[2] = µbases
							if πE = πg.CheckLocal(πF, µnamespace, "namespace"); πE != nil {
								continue
							}
							πTemp001[3] = µnamespace
							πTemp002 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßABCMeta); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µmcls, "mcls"); πE != nil {
								continue
							}
							πTemp002[1] = µmcls
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ß__new__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µcls = πTemp004
							// line 89: abstracts = set(name
							πF.SetLineno(89)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/abc.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µname *πg.Object = πg.UnboundLocal; _ = µname
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
								return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
									for ; πF.State() >= 0; πF.PopCheckpoint() {
										switch πF.State() {
										case 0:
										case 1: goto Label1
										case 2: goto Label2
										case 6: goto Label6
										default: panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µnamespace, "namespace"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µnamespace, ßitems, nil); πE != nil {
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
											µname = πTemp003
											µvalue = πTemp006
										}
										if πE != nil || !πTemp005 {
											continue
										}
										πF.PushCheckpoint(1)            
										πTemp007 = πF.MakeArgs(3)
										if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
											continue
										}
										πTemp007[0] = µvalue
										πTemp007[1] = ß__isabstractmethod__.ToObject()
										if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
											continue
										}
										πTemp007[2] = πTemp002
										if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
											continue
										}
										if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp007)
										if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
											continue
										}
										if πTemp005 {
											goto Label4
										}
										goto Label5
										// line 89: abstracts = set(name
										πF.SetLineno(89)
									Label4:
										// line 89: abstracts = set(name
										πF.SetLineno(89)
										if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
											continue
										}
										πF.PushCheckpoint(6)
										return µname, nil
									Label6:
										πTemp002 = πSent
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
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µabstracts = πTemp006
							if πE = πg.CheckLocal(πF, µbases, "bases"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Iter(πF, µbases); πE != nil {
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
							if πTemp006, πE = πg.Next(πF, πTemp004); πE != nil {
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
								µbase = πTemp006
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(1)            
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
								continue
							}
							πTemp001[0] = µbase
							πTemp001[1] = ß__abstractmethods__.ToObject()
							if πTemp009, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp010
							if πTemp009, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp006, πE = πg.Iter(πF, πTemp010); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp008 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label6
							}
							if πTemp009, πE = πg.Next(πF, πTemp006); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp011 = !isStop
							} else {
								πTemp011 = true
								µname = πTemp009
							}
							if πE != nil || !πTemp011 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 94: value = getattr(cls, name, None)
							πF.SetLineno(94)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp001[0] = µcls
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001[1] = µname
							if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp009
							if πTemp009, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µvalue = πTemp010
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp001[0] = µvalue
							πTemp001[1] = ß__isabstractmethod__.ToObject()
							if πTemp009, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[2] = πTemp009
							if πTemp009, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp011, πE = πg.IsTrue(πF, πTemp010); πE != nil {
								continue
							}
							if πTemp011 {
								goto Label7
							}
							goto Label8
							// line 95: if getattr(value, "__isabstractmethod__", False):
							πF.SetLineno(95)
						Label7:
							// line 96: abstracts.add(name)
							πF.SetLineno(96)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001[0] = µname
							if πE = πg.CheckLocal(πF, µabstracts, "abstracts"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µabstracts, ßadd, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label8
						Label8:
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 97: cls.__abstractmethods__ = frozenset(abstracts)
							πF.SetLineno(97)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µabstracts, "abstracts"); πE != nil {
								continue
							}
							πTemp001[0] = µabstracts
							if πTemp004, πE = πg.ResolveGlobal(πF, ßfrozenset); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µcls, ß__abstractmethods__, πTemp004); πE != nil {
								continue
							}
							// line 99: cls._abc_registry = _weakrefset.WeakSet()
							πF.SetLineno(99)
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_weakrefset); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßWeakSet, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µcls, ß_abc_registry, πTemp006); πE != nil {
								continue
							}
							// line 100: cls._abc_cache = _weakrefset.WeakSet()
							πF.SetLineno(100)
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_weakrefset); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßWeakSet, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µcls, ß_abc_cache, πTemp006); πE != nil {
								continue
							}
							// line 101: cls._abc_negative_cache = _weakrefset.WeakSet()
							πF.SetLineno(101)
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_weakrefset); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßWeakSet, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µcls, ß_abc_negative_cache, πTemp006); πE != nil {
								continue
							}
							// line 102: cls._abc_negative_cache_version = ABCMeta._abc_invalidation_counter
							πF.SetLineno(102)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßABCMeta); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ß_abc_invalidation_counter, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µcls, ß_abc_negative_cache_version, πTemp004); πE != nil {
								continue
							}
							// line 103: return cls
							πF.SetLineno(103)
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
					if πE = πClass.SetItem(πF, ß__new__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 105: def register(cls, subclass):
					πF.SetLineno(105)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "subclass", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("register", "build/src/__python__/abc.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µsubclass *πg.Object = πArgs[1]; _ = µsubclass
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
							// line 106: """Register a virtual subclass of an ABC."""
							πF.SetLineno(106)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µsubclass, "subclass"); πE != nil {
								continue
							}
							πTemp002[0] = µsubclass
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
							// line 107: if not isinstance(subclass, type):
							πF.SetLineno(107)
						Label1:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("Can only register classes").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 108: raise TypeError("Can only register classes")
							πF.SetLineno(108)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µsubclass, "subclass"); πE != nil {
								continue
							}
							πTemp002[0] = µsubclass
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp002[1] = µcls
							if πTemp001, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 109: if issubclass(subclass, cls):
							πF.SetLineno(109)
						Label3:
							// line 110: return  # Already a subclass
							πF.SetLineno(110)
							πR = πg.None
							continue
							goto Label4
						Label4:
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp002[0] = µcls
							if πE = πg.CheckLocal(πF, µsubclass, "subclass"); πE != nil {
								continue
							}
							πTemp002[1] = µsubclass
							if πTemp001, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							goto Label6
							// line 113: if issubclass(cls, subclass):
							πF.SetLineno(113)
						Label5:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("Refusing to create an inheritance cycle").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 115: raise RuntimeError("Refusing to create an inheritance cycle")
							πF.SetLineno(115)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label6
						Label6:
							// line 116: cls._abc_registry.add(subclass)
							πF.SetLineno(116)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsubclass, "subclass"); πE != nil {
								continue
							}
							πTemp002[0] = µsubclass
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcls, ß_abc_registry, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßadd, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 117: ABCMeta._abc_invalidation_counter += 1  # Invalidate negative cache
							πF.SetLineno(117)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßABCMeta); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ß_abc_invalidation_counter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßABCMeta); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp004, ß_abc_invalidation_counter, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßregister.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 119: def _dump_registry(cls, file=None):
					πF.SetLineno(119)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "file", Def: πTemp005}
					πTemp004 = πg.NewFunction(πg.NewCode("_dump_registry", "build/src/__python__/abc.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µfile *πg.Object = πArgs[1]; _ = µfile
						var µname *πg.Object = πg.UnboundLocal; _ = µname
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 120: """Debug helper to print the ABC registry."""
							πF.SetLineno(120)
							// line 121: print >> file, "Class: %s.%s" % (cls.__module__, cls.__name__)
							πF.SetLineno(121)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µcls, ß__module__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µcls, ß__name__, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Class: %s.%s").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.Print(πF, πTemp001, true); πE != nil {
								continue
							}
							// line 122: print >> file, "Inv.counter: %s" % ABCMeta._abc_invalidation_counter
							πF.SetLineno(122)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßABCMeta); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß_abc_invalidation_counter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Inv.counter: %s").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.Print(πF, πTemp001, true); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µcls, ß__dict__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
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
								µname = πTemp003
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)            
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ß_abc_.ToObject()
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µname, ßstartswith, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label4
							}
							goto Label5
							// line 124: if name.startswith("_abc_"):
							πF.SetLineno(124)
						Label4:
							// line 125: value = getattr(cls, name)
							πF.SetLineno(125)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp001[0] = µcls
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001[1] = µname
							if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µvalue = πTemp004
							// line 126: print >> file, "%s: %r" % (name, value)
							πF.SetLineno(126)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(µname, µvalue).ToObject()
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s: %r").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.Print(πF, πTemp001, true); πE != nil {
								continue
							}
							goto Label5
						Label5:
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
					if πE = πClass.SetItem(πF, ß_dump_registry.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 128: def __instancecheck__(cls, instance):
					πF.SetLineno(128)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "instance", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("__instancecheck__", "build/src/__python__/abc.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µinstance *πg.Object = πArgs[1]; _ = µinstance
						var µsubclass *πg.Object = πg.UnboundLocal; _ = µsubclass
						var µsubtype *πg.Object = πg.UnboundLocal; _ = µsubtype
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
						var πTemp008 *πg.Object
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 129: """Override for isinstance(instance, cls)."""
							πF.SetLineno(129)
							// line 131: subclass = getattr(instance, '__class__', None)
							πF.SetLineno(131)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µinstance, "instance"); πE != nil {
								continue
							}
							πTemp001[0] = µinstance
							πTemp001[1] = ß__class__.ToObject()
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
							µsubclass = πTemp003
							if πE = πg.CheckLocal(πF, µsubclass, "subclass"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µsubclass != πTemp005).ToObject()
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µsubclass, "subclass"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µcls, ß_abc_cache, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, πTemp005, µsubclass); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp006).ToObject()
							πTemp002 = πTemp003
						Label1:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							goto Label3
							// line 132: if subclass is not None and subclass in cls._abc_cache:
							πF.SetLineno(132)
						Label2:
							// line 133: return True
							πF.SetLineno(133)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label3:
							// line 134: subtype = type(instance)
							πF.SetLineno(134)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µinstance, "instance"); πE != nil {
								continue
							}
							πTemp001[0] = µinstance
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µsubtype = πTemp003
							if πE = πg.CheckLocal(πF, µsubtype, "subtype"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsubclass, "subclass"); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µsubtype == µsubclass).ToObject()
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µsubclass, "subclass"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µsubclass == πTemp005).ToObject()
							πTemp002 = πTemp003
						Label4:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							goto Label6
							// line 138: if subtype is subclass or subclass is None:
							πF.SetLineno(138)
						Label5:
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µcls, ß_abc_negative_cache_version, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßABCMeta); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ß_abc_invalidation_counter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp005, πTemp008); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µsubtype, "subtype"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µcls, ß_abc_negative_cache, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, πTemp005, µsubtype); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp006).ToObject()
							πTemp002 = πTemp003
						Label7:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							goto Label9
							// line 139: if (cls._abc_negative_cache_version ==
							πF.SetLineno(139)
						Label8:
							// line 142: return False
							πF.SetLineno(142)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label9
						Label9:
							// line 144: return cls.__subclasscheck__(subtype)
							πF.SetLineno(144)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsubtype, "subtype"); πE != nil {
								continue
							}
							πTemp001[0] = µsubtype
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcls, ß__subclasscheck__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp003
							continue
							goto Label6
						Label6:
							// line 145: return (cls.__subclasscheck__(subclass) or
							πF.SetLineno(145)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsubclass, "subclass"); πE != nil {
								continue
							}
							πTemp001[0] = µsubclass
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µcls, ß__subclasscheck__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πTemp005
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label10
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsubtype, "subtype"); πE != nil {
								continue
							}
							πTemp001[0] = µsubtype
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µcls, ß__subclasscheck__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πTemp005
						Label10:
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
					if πE = πClass.SetItem(πF, ß__instancecheck__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 148: def __subclasscheck__(cls, subclass):
					πF.SetLineno(148)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "subclass", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("__subclasscheck__", "build/src/__python__/abc.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µsubclass *πg.Object = πArgs[1]; _ = µsubclass
						var µok *πg.Object = πg.UnboundLocal; _ = µok
						var µrcls *πg.Object = πg.UnboundLocal; _ = µrcls
						var µscls *πg.Object = πg.UnboundLocal; _ = µscls
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 18: goto Label18
							case 19: goto Label19
							case 13: goto Label13
							case 14: goto Label14
							default: panic("unexpected function state")
							}
							// line 149: """Override for issubclass(subclass, cls)."""
							πF.SetLineno(149)
							if πE = πg.CheckLocal(πF, µsubclass, "subclass"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcls, ß_abc_cache, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, µsubclass); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 151: if subclass in cls._abc_cache:
							πF.SetLineno(151)
						Label1:
							// line 152: return True
							πF.SetLineno(152)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcls, ß_abc_negative_cache_version, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßABCMeta); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß_abc_invalidation_counter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µsubclass, "subclass"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcls, ß_abc_negative_cache, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, µsubclass); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							goto Label5
							// line 154: if cls._abc_negative_cache_version < ABCMeta._abc_invalidation_counter:
							πF.SetLineno(154)
						Label3:
							// line 156: cls._abc_negative_cache = _weakrefset.WeakSet()
							πF.SetLineno(156)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_weakrefset); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßWeakSet, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µcls, ß_abc_negative_cache, πTemp002); πE != nil {
								continue
							}
							// line 157: cls._abc_negative_cache_version = ABCMeta._abc_invalidation_counter
							πF.SetLineno(157)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßABCMeta); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß_abc_invalidation_counter, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µcls, ß_abc_negative_cache_version, πTemp001); πE != nil {
								continue
							}
							goto Label5
							// line 158: elif subclass in cls._abc_negative_cache:
							πF.SetLineno(158)
						Label4:
							// line 159: return False
							πF.SetLineno(159)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label5
						Label5:
							// line 161: ok = cls.__subclasshook__(subclass)
							πF.SetLineno(161)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsubclass, "subclass"); πE != nil {
								continue
							}
							πTemp006[0] = µsubclass
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcls, ß__subclasshook__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µok = πTemp002
							if πE = πg.CheckLocal(πF, µok, "ok"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µok != πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							goto Label7
							// line 162: if ok is not NotImplemented:
							πF.SetLineno(162)
						Label6:
							// line 163: assert isinstance(ok, bool)
							πF.SetLineno(163)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µok, "ok"); πE != nil {
								continue
							}
							πTemp006[0] = µok
							if πTemp001, πE = πg.ResolveGlobal(πF, ßbool); πE != nil {
								continue
							}
							πTemp006[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µok, "ok"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µok); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label8
							}
							goto Label9
							// line 164: if ok:
							πF.SetLineno(164)
						Label8:
							// line 165: cls._abc_cache.add(subclass)
							πF.SetLineno(165)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsubclass, "subclass"); πE != nil {
								continue
							}
							πTemp006[0] = µsubclass
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcls, ß_abc_cache, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßadd, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label10
						Label9:
							// line 167: cls._abc_negative_cache.add(subclass)
							πF.SetLineno(167)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsubclass, "subclass"); πE != nil {
								continue
							}
							πTemp006[0] = µsubclass
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcls, ß_abc_negative_cache, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßadd, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label10
						Label10:
							// line 168: return ok
							πF.SetLineno(168)
							if πE = πg.CheckLocal(πF, µok, "ok"); πE != nil {
								continue
							}
							πR = µok
							continue
							goto Label7
						Label7:
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µsubclass, "subclass"); πE != nil {
								continue
							}
							πTemp006[0] = µsubclass
							πTemp006[1] = ß__mro__.ToObject()
							πTemp002 = πg.NewTuple0().ToObject()
							πTemp006[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp003, πE = πg.Contains(πF, πTemp004, µcls); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label11
							}
							goto Label12
							// line 170: if cls in getattr(subclass, '__mro__', ()):
							πF.SetLineno(170)
						Label11:
							// line 171: cls._abc_cache.add(subclass)
							πF.SetLineno(171)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsubclass, "subclass"); πE != nil {
								continue
							}
							πTemp006[0] = µsubclass
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcls, ß_abc_cache, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßadd, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 172: return True
							πF.SetLineno(172)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label12
						Label12:
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcls, ß_abc_registry, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
								continue
							}
							πF.PushCheckpoint(14)
							πTemp003 = false
						Label13:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
								πF.PopCheckpoint()
								goto Label15
							}
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µrcls = πTemp002
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(13)            
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µsubclass, "subclass"); πE != nil {
								continue
							}
							πTemp006[0] = µsubclass
							if πE = πg.CheckLocal(πF, µrcls, "rcls"); πE != nil {
								continue
							}
							πTemp006[1] = µrcls
							if πTemp002, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label16
							}
							goto Label17
							// line 175: if issubclass(subclass, rcls):
							πF.SetLineno(175)
						Label16:
							// line 176: cls._abc_cache.add(subclass)
							πF.SetLineno(176)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsubclass, "subclass"); πE != nil {
								continue
							}
							πTemp006[0] = µsubclass
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcls, ß_abc_cache, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßadd, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 177: return True
							πF.SetLineno(177)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label17
						Label17:
							continue
						Label14:
							if πE != nil || πR != nil {
								continue
							}
						Label15:
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcls, ß__subclasses__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(19)
							πTemp003 = false
						Label18:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
								πF.PopCheckpoint()
								goto Label20
							}
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µscls = πTemp002
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(18)            
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µsubclass, "subclass"); πE != nil {
								continue
							}
							πTemp006[0] = µsubclass
							if πE = πg.CheckLocal(πF, µscls, "scls"); πE != nil {
								continue
							}
							πTemp006[1] = µscls
							if πTemp002, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label21
							}
							goto Label22
							// line 180: if issubclass(subclass, scls):
							πF.SetLineno(180)
						Label21:
							// line 181: cls._abc_cache.add(subclass)
							πF.SetLineno(181)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsubclass, "subclass"); πE != nil {
								continue
							}
							πTemp006[0] = µsubclass
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcls, ß_abc_cache, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßadd, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 182: return True
							πF.SetLineno(182)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label22
						Label22:
							continue
						Label19:
							if πE != nil || πR != nil {
								continue
							}
						Label20:
							// line 184: cls._abc_negative_cache.add(subclass)
							πF.SetLineno(184)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsubclass, "subclass"); πE != nil {
								continue
							}
							πTemp006[0] = µsubclass
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcls, ß_abc_negative_cache, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßadd, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 185: return False
							πF.SetLineno(185)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__subclasscheck__.ToObject(), πTemp006); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("ABCMeta").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßABCMeta.ToObject(), πTemp007); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("abc", Code)
}
