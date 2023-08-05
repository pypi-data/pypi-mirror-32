package itertools
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/itertools.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßException := πg.InternStr("Exception")
		ßFalse := πg.InternStr("False")
		ßNone := πg.InternStr("None")
		ßStopIteration := πg.InternStr("StopIteration")
		ßTrue := πg.InternStr("True")
		ßZipExhausted := πg.InternStr("ZipExhausted")
		ß__init__ := πg.InternStr("__init__")
		ß__iter__ := πg.InternStr("__iter__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_collections := πg.InternStr("_collections")
		ß_grouper := πg.InternStr("_grouper")
		ßappend := πg.InternStr("append")
		ßbool := πg.InternStr("bool")
		ßchain := πg.InternStr("chain")
		ßclassmethod := πg.InternStr("classmethod")
		ßcombinations := πg.InternStr("combinations")
		ßcombinations_with_replacement := πg.InternStr("combinations_with_replacement")
		ßcompress := πg.InternStr("compress")
		ßcount := πg.InternStr("count")
		ßcurriter := πg.InternStr("curriter")
		ßcurrkey := πg.InternStr("currkey")
		ßcurrvalue := πg.InternStr("currvalue")
		ßcycle := πg.InternStr("cycle")
		ßdeque := πg.InternStr("deque")
		ßdropwhile := πg.InternStr("dropwhile")
		ßenumerate := πg.InternStr("enumerate")
		ßfillvalue := πg.InternStr("fillvalue")
		ßfrom_iterable := πg.InternStr("from_iterable")
		ßget := πg.InternStr("get")
		ßgroupby := πg.InternStr("groupby")
		ßifilter := πg.InternStr("ifilter")
		ßifilterfalse := πg.InternStr("ifilterfalse")
		ßimap := πg.InternStr("imap")
		ßislice := πg.InternStr("islice")
		ßit := πg.InternStr("it")
		ßiter := πg.InternStr("iter")
		ßiterables := πg.InternStr("iterables")
		ßizip := πg.InternStr("izip")
		ßizip_longest := πg.InternStr("izip_longest")
		ßkeyfunc := πg.InternStr("keyfunc")
		ßlen := πg.InternStr("len")
		ßlist := πg.InternStr("list")
		ßmap := πg.InternStr("map")
		ßmaxint := πg.InternStr("maxint")
		ßnext := πg.InternStr("next")
		ßobject := πg.InternStr("object")
		ßpermutations := πg.InternStr("permutations")
		ßpopleft := πg.InternStr("popleft")
		ßproduct := πg.InternStr("product")
		ßrange := πg.InternStr("range")
		ßrepeat := πg.InternStr("repeat")
		ßset := πg.InternStr("set")
		ßslice := πg.InternStr("slice")
		ßsorted := πg.InternStr("sorted")
		ßstarmap := πg.InternStr("starmap")
		ßstart := πg.InternStr("start")
		ßstep := πg.InternStr("step")
		ßstop := πg.InternStr("stop")
		ßsys := πg.InternStr("sys")
		ßtakewhile := πg.InternStr("takewhile")
		ßtee := πg.InternStr("tee")
		ßtgtkey := πg.InternStr("tgtkey")
		ßtuple := πg.InternStr("tuple")
		ßxrange := πg.InternStr("xrange")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
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
		var πTemp017 *πg.Object
		_ = πTemp017
		var πTemp018 *πg.Object
		_ = πTemp018
		var πTemp019 *πg.Object
		_ = πTemp019
		var πTemp020 *πg.Object
		_ = πTemp020
		var πTemp021 *πg.Object
		_ = πTemp021
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 15: """Utilities for iterating over containers."""
			πF.SetLineno(15)
			// line 17: import _collections
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "_collections"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_collections.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: import sys
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 20: class chain(object):
			πF.SetLineno(20)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("chain", "build/src/__python__/itertools.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 []*πg.Object
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
					// line 22: def from_iterable(cls, iterables):
					πF.SetLineno(22)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "iterables", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("from_iterable", "build/src/__python__/itertools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µiterables *πg.Object = πArgs[1]; _ = µiterables
						var µit *πg.Object = πg.UnboundLocal; _ = µit
						var µelement *πg.Object = πg.UnboundLocal; _ = µelement
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
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
								case 4: goto Label4
								case 5: goto Label5
								case 7: goto Label7
								default: panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µiterables, "iterables"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µiterables); πE != nil {
									continue
								}
								πF.PushCheckpoint(2)
								πTemp002 = false
							Label1:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp002 {
									πF.PopCheckpoint()
									goto Label3
								}
								if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp003 = !isStop
								} else {
									πTemp003 = true
									µit = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.Iter(πF, µit); πE != nil {
									continue
								}
								πF.PushCheckpoint(5)
								πTemp003 = false
							Label4:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp003 {
									πF.PopCheckpoint()
									goto Label6
								}
								if πTemp006, πE = πg.Next(πF, πTemp004); πE != nil {
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
									µelement = πTemp006
								}
								if πE != nil || !πTemp005 {
									continue
								}
								πF.PushCheckpoint(4)            
								// line 25: yield element
								πF.SetLineno(25)
								if πE = πg.CheckLocal(πF, µelement, "element"); πE != nil {
									continue
								}
								πF.PushCheckpoint(7)
								return µelement, nil
							Label7:
								πTemp006 = πSent
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
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßfrom_iterable.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 27: from_iterable = classmethod(from_iterable)
					πF.SetLineno(27)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßfrom_iterable); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ßfrom_iterable.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 29: def __init__(self, *iterables):
					πF.SetLineno(29)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/itertools.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µiterables *πg.Object = πArgs[1]; _ = µiterables
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
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
							if πE = πg.CheckLocal(πF, µiterables, "iterables"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µiterables); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 30: if not iterables:
							πF.SetLineno(30)
						Label1:
							// line 31: self.iterables = iter([[]])
							πF.SetLineno(31)
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = make([]*πg.Object, 1)
							πTemp005 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp001
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßiterables, πTemp001); πE != nil {
								continue
							}
							goto Label3
						Label2:
							// line 33: self.iterables = iter(iterables)
							πF.SetLineno(33)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µiterables, "iterables"); πE != nil {
								continue
							}
							πTemp003[0] = µiterables
							if πTemp001, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßiterables, πTemp001); πE != nil {
								continue
							}
							goto Label3
						Label3:
							// line 34: self.curriter = iter(next(self.iterables))
							πF.SetLineno(34)
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßiterables, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnext); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp006
							if πTemp001, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcurriter, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 36: def __iter__(self):
					πF.SetLineno(36)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("__iter__", "build/src/__python__/itertools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 37: return self
							πF.SetLineno(37)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πR = µself
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 39: def next(self):
					πF.SetLineno(39)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("next", "build/src/__python__/itertools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µflag *πg.Object = πg.UnboundLocal; _ = µflag
						var µret *πg.Object = πg.UnboundLocal; _ = µret
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.BaseException
						_ = πTemp006
						var πTemp007 *πg.Traceback
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
							case 5: goto Label5
							default: panic("unexpected function state")
							}
							// line 40: flag = True
							πF.SetLineno(40)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µflag = πTemp001
							// line 41: while flag:
							πF.SetLineno(41)
							πF.PushCheckpoint(2)
							πTemp002 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µflag, "flag"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µflag); πE != nil {
								continue
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 42: try:
							πF.SetLineno(42)
							πF.PushCheckpoint(5)
							// line 43: ret = next(self.curriter)
							πF.SetLineno(43)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcurriter, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnext); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µret = πTemp005
							// line 44: flag = False
							πF.SetLineno(44)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							µflag = πTemp001
							πF.PopCheckpoint()
							goto Label4
						Label5:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
							continue
							// line 45: except StopIteration:
							πF.SetLineno(45)
						Label6:
							// line 46: self.curriter = iter(next(self.iterables))
							πF.SetLineno(46)
							πTemp004 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßiterables, nil); πE != nil {
								continue
							}
							πTemp008[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnext); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp004[0] = πTemp005
							if πTemp001, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcurriter, πTemp001); πE != nil {
								continue
							}
							πF.RestoreExc(nil, nil)
							goto Label4
						Label4:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 47: return ret
							πF.SetLineno(47)
							if πE = πg.CheckLocal(πF, µret, "ret"); πE != nil {
								continue
							}
							πR = µret
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßnext.ToObject(), πTemp006); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("chain").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßchain.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 50: def compress(data, selectors):
			πF.SetLineno(50)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "data", Def: nil}
			πTemp006[1] = πg.Param{Name: "selectors", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("compress", "build/src/__python__/itertools.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µdata *πg.Object = πArgs[0]; _ = µdata
				var µselectors *πg.Object = πArgs[1]; _ = µselectors
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
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
					// line 51: return (d for d,s in izip(data, selectors) if s)
					πF.SetLineno(51)
					πTemp002 = make([]πg.Param, 0)
					πTemp001 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/itertools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µs *πg.Object = πg.UnboundLocal; _ = µs
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
						var πTemp007 *πg.Object
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
								πTemp002 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
									continue
								}
								πTemp002[0] = µdata
								if πE = πg.CheckLocal(πF, µselectors, "selectors"); πE != nil {
									continue
								}
								πTemp002[1] = µselectors
								if πTemp003, πE = πg.ResolveGlobal(πF, ßizip); πE != nil {
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
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
										continue
									}
									µd = πTemp004
									µs = πTemp007
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)            
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.IsTrue(πF, µs); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label4
								}
								goto Label5
								// line 51: return (d for d,s in izip(data, selectors) if s)
								πF.SetLineno(51)
							Label4:
								// line 51: return (d for d,s in izip(data, selectors) if s)
								πF.SetLineno(51)
								if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
									continue
								}
								πF.PushCheckpoint(6)
								return µd, nil
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
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßcompress.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 54: def count(start=0, step=1):
			πF.SetLineno(54)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "start", Def: πg.NewInt(0).ToObject()}
			πTemp006[1] = πg.Param{Name: "step", Def: πg.NewInt(1).ToObject()}
			πTemp004 = πg.NewFunction(πg.NewCode("count", "build/src/__python__/itertools.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µstart *πg.Object = πArgs[0]; _ = µstart
				var µstep *πg.Object = πArgs[1]; _ = µstep
				var µn *πg.Object = πg.UnboundLocal; _ = µn
				var πTemp001 bool
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
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
						// line 55: n = start
						πF.SetLineno(55)
						if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
							continue
						}
						µn = µstart
						// line 56: while True:
						πF.SetLineno(56)
						πF.PushCheckpoint(2)
						πTemp001 = false
					Label1:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp001 {
							πF.PopCheckpoint()
							goto Label3
						}
						if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
							continue
						}
						if πE != nil || !πTemp002 {
							continue
						}
						πF.PushCheckpoint(1)            
						// line 57: yield n
						πF.SetLineno(57)
						if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
							continue
						}
						πF.PushCheckpoint(4)
						return µn, nil
					Label4:
						πTemp003 = πSent
						// line 58: n += step
						πF.SetLineno(58)
						if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µstep, "step"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.IAdd(πF, µn, µstep); πE != nil {
							continue
						}
						µn = πTemp003
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
			if πE = πF.Globals().SetItem(πF, ßcount.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 61: def cycle(iterable):
			πF.SetLineno(61)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "iterable", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("cycle", "build/src/__python__/itertools.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µiterable *πg.Object = πArgs[0]; _ = µiterable
				var µsaved *πg.Object = πg.UnboundLocal; _ = µsaved
				var µelement *πg.Object = πg.UnboundLocal; _ = µelement
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 1: goto Label1
						case 2: goto Label2
						case 4: goto Label4
						case 5: goto Label5
						case 6: goto Label6
						case 8: goto Label8
						case 9: goto Label9
						case 11: goto Label11
						default: panic("unexpected function state")
						}
						// line 62: saved = []
						πF.SetLineno(62)
						πTemp001 = make([]*πg.Object, 0)
						πTemp002 = πg.NewList(πTemp001...).ToObject()
						µsaved = πTemp002
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
							µelement = πTemp005
						}
						if πE != nil || !πTemp004 {
							continue
						}
						πF.PushCheckpoint(1)            
						// line 64: yield element
						πF.SetLineno(64)
						if πE = πg.CheckLocal(πF, µelement, "element"); πE != nil {
							continue
						}
						πF.PushCheckpoint(4)
						return µelement, nil
					Label4:
						πTemp005 = πSent
						// line 65: saved.append(element)
						πF.SetLineno(65)
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µelement, "element"); πE != nil {
							continue
						}
						πTemp001[0] = µelement
						if πE = πg.CheckLocal(πF, µsaved, "saved"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.GetAttr(πF, µsaved, ßappend, nil); πE != nil {
							continue
						}
						if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						continue
					Label2:
						if πE != nil || πR != nil {
							continue
						}
					Label3:
						// line 66: while saved:
						πF.SetLineno(66)
						πF.PushCheckpoint(6)
						πTemp003 = false
					Label5:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp003 {
							πF.PopCheckpoint()
							goto Label7
						}
						if πE = πg.CheckLocal(πF, µsaved, "saved"); πE != nil {
							continue
						}
						if πTemp004, πE = πg.IsTrue(πF, µsaved); πE != nil {
							continue
						}
						if πE != nil || !πTemp004 {
							continue
						}
						πF.PushCheckpoint(5)            
						if πE = πg.CheckLocal(πF, µsaved, "saved"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.Iter(πF, µsaved); πE != nil {
							continue
						}
						πF.PushCheckpoint(9)
						πTemp004 = false
					Label8:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp004 {
							πF.PopCheckpoint()
							goto Label10
						}
						if πTemp005, πE = πg.Next(πF, πTemp002); πE != nil {
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
							µelement = πTemp005
						}
						if πE != nil || !πTemp007 {
							continue
						}
						πF.PushCheckpoint(8)            
						// line 68: yield element
						πF.SetLineno(68)
						if πE = πg.CheckLocal(πF, µelement, "element"); πE != nil {
							continue
						}
						πF.PushCheckpoint(11)
						return µelement, nil
					Label11:
						πTemp005 = πSent
						continue
					Label9:
						if πE != nil || πR != nil {
							continue
						}
					Label10:
						continue
					Label6:
						if πE != nil || πR != nil {
							continue
						}
					Label7:
					}
					return nil, πE
				}).ToObject(), nil
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßcycle.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 71: def dropwhile(predicate, iterable):
			πF.SetLineno(71)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "predicate", Def: nil}
			πTemp006[1] = πg.Param{Name: "iterable", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("dropwhile", "build/src/__python__/itertools.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpredicate *πg.Object = πArgs[0]; _ = µpredicate
				var µiterable *πg.Object = πArgs[1]; _ = µiterable
				var µx *πg.Object = πg.UnboundLocal; _ = µx
				var πTemp001 []*πg.Object
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
						case 7: goto Label7
						case 8: goto Label8
						case 10: goto Label10
						default: panic("unexpected function state")
						}
						// line 72: iterable = iter(iterable)
						πF.SetLineno(72)
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
						µiterable = πTemp003
						if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.Iter(πF, µiterable); πE != nil {
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
						if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
							µx = πTemp003
						}
						if πE != nil || !πTemp005 {
							continue
						}
						πF.PushCheckpoint(1)            
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
							continue
						}
						πTemp001[0] = µx
						if πE = πg.CheckLocal(πF, µpredicate, "predicate"); πE != nil {
							continue
						}
						if πTemp006, πE = µpredicate.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
							continue
						}
						πTemp003 = πg.GetBool(!πTemp005).ToObject()
						if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
							continue
						}
						if πTemp005 {
							goto Label4
						}
						goto Label5
						// line 74: if not predicate(x):
						πF.SetLineno(74)
					Label4:
						// line 75: yield x
						πF.SetLineno(75)
						if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
							continue
						}
						πF.PushCheckpoint(6)
						return µx, nil
					Label6:
						πTemp003 = πSent
						// line 76: break
						πF.SetLineno(76)
						πTemp004 = true
						continue
						goto Label5
					Label5:
						continue
					Label2:
						if πE != nil || πR != nil {
							continue
						}
					Label3:
						if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.Iter(πF, µiterable); πE != nil {
							continue
						}
						πF.PushCheckpoint(8)
						πTemp004 = false
					Label7:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp004 {
							πF.PopCheckpoint()
							goto Label9
						}
						if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
							µx = πTemp003
						}
						if πE != nil || !πTemp005 {
							continue
						}
						πF.PushCheckpoint(7)            
						// line 78: yield x
						πF.SetLineno(78)
						if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
							continue
						}
						πF.PushCheckpoint(10)
						return µx, nil
					Label10:
						πTemp003 = πSent
						continue
					Label8:
						if πE != nil || πR != nil {
							continue
						}
					Label9:
					}
					return nil, πE
				}).ToObject(), nil
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßdropwhile.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 81: class groupby(object):
			πF.SetLineno(81)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp010, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp010
			πTemp003 = πg.NewDict()
			if πTemp008, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp008); πE != nil {
				continue
			}
			_, πE = πg.NewCode("groupby", "build/src/__python__/itertools.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 84: def __init__(self, iterable, key=None):
					πF.SetLineno(84)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "iterable", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "key", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/itertools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µiterable *πg.Object = πArgs[1]; _ = µiterable
						var µkey *πg.Object = πArgs[2]; _ = µkey
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 []πg.Param
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
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µkey == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 85: if key is None:
							πF.SetLineno(85)
						Label1:
							// line 86: key = lambda x: x
							πF.SetLineno(86)
							πTemp004 = make([]πg.Param, 1)
							πTemp004[0] = πg.Param{Name: "x", Def: nil}
							πTemp001 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/itertools.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µx *πg.Object = πArgs[0]; _ = µx
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 86: key = lambda x: x
									πF.SetLineno(86)
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
							µkey = πTemp001
							goto Label2
						Label2:
							// line 87: self.keyfunc = key
							πF.SetLineno(87)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µkey); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßkeyfunc, πTemp001); πE != nil {
								continue
							}
							// line 88: self.it = iter(iterable)
							πF.SetLineno(88)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
								continue
							}
							πTemp005[0] = µiterable
							if πTemp001, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßit, πTemp001); πE != nil {
								continue
							}
							// line 89: self.tgtkey = self.currkey = self.currvalue = object()
							πF.SetLineno(89)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßtgtkey, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcurrkey, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcurrvalue, πTemp001); πE != nil {
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
					// line 91: def __iter__(self):
					πF.SetLineno(91)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__iter__", "build/src/__python__/itertools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 92: return self
							πF.SetLineno(92)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πR = µself
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 94: def next(self):
					πF.SetLineno(94)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("next", "build/src/__python__/itertools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 bool
						_ = πTemp001
						var πTemp002 bool
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 95: while self.currkey == self.tgtkey:
							πF.SetLineno(95)
							πF.PushCheckpoint(2)
							πTemp001 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp001 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcurrkey, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßtgtkey, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πE != nil || !πTemp002 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 96: self.currvalue = next(self.it)    # Exit on StopIteration
							πF.SetLineno(96)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßit, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnext); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcurrvalue, πTemp003); πE != nil {
								continue
							}
							// line 97: self.currkey = self.keyfunc(self.currvalue)
							πF.SetLineno(97)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcurrvalue, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßkeyfunc, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcurrkey, πTemp003); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 98: self.tgtkey = self.currkey
							πF.SetLineno(98)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcurrkey, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtgtkey, πTemp004); πE != nil {
								continue
							}
							// line 99: return (self.currkey, self._grouper(self.tgtkey))
							πF.SetLineno(99)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcurrkey, nil); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßtgtkey, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_grouper, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003 = πg.NewTuple2(πTemp004, πTemp007).ToObject()
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
					if πE = πClass.SetItem(πF, ßnext.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 101: def _grouper(self, tgtkey):
					πF.SetLineno(101)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "tgtkey", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("_grouper", "build/src/__python__/itertools.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtgtkey *πg.Object = πArgs[1]; _ = µtgtkey
						var πTemp001 bool
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
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
								case 4: goto Label4
								default: panic("unexpected function state")
								}
								// line 102: while self.currkey == tgtkey:
								πF.SetLineno(102)
								πF.PushCheckpoint(2)
								πTemp001 = false
							Label1:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp001 {
									πF.PopCheckpoint()
									goto Label3
								}
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µself, ßcurrkey, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µtgtkey, "tgtkey"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Eq(πF, πTemp004, µtgtkey); πE != nil {
									continue
								}
								if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πE != nil || !πTemp002 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 103: yield self.currvalue
								πF.SetLineno(103)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µself, ßcurrvalue, nil); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return πTemp003, nil
							Label4:
								πTemp004 = πSent
								// line 104: self.currvalue = next(self.it)    # Exit on StopIteration
								πF.SetLineno(104)
								πTemp005 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µself, ßit, nil); πE != nil {
									continue
								}
								πTemp005[0] = πTemp004
								if πTemp004, πE = πg.ResolveGlobal(πF, ßnext); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp006); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µself, ßcurrvalue, πTemp004); πE != nil {
									continue
								}
								// line 105: self.currkey = self.keyfunc(self.currvalue)
								πF.SetLineno(105)
								πTemp005 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µself, ßcurrvalue, nil); πE != nil {
									continue
								}
								πTemp005[0] = πTemp004
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µself, ßkeyfunc, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp006); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µself, ßcurrkey, πTemp004); πE != nil {
									continue
								}
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
					if πE = πClass.SetItem(πF, ß_grouper.ToObject(), πTemp005); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp009, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp009 == nil {
				πTemp009 = πg.TypeType.ToObject()
			}
			if πTemp010, πE = πTemp009.Call(πF, []*πg.Object{πg.NewStr("groupby").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßgroupby.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 108: def ifilter(predicate, iterable):
			πF.SetLineno(108)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "predicate", Def: nil}
			πTemp006[1] = πg.Param{Name: "iterable", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("ifilter", "build/src/__python__/itertools.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpredicate *πg.Object = πArgs[0]; _ = µpredicate
				var µiterable *πg.Object = πArgs[1]; _ = µiterable
				var µx *πg.Object = πg.UnboundLocal; _ = µx
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
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 8: goto Label8
						case 3: goto Label3
						case 4: goto Label4
						default: panic("unexpected function state")
						}
						if πE = πg.CheckLocal(πF, µpredicate, "predicate"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
							continue
						}
						πTemp001 = πg.GetBool(µpredicate == πTemp002).ToObject()
						if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
							continue
						}
						if πTemp003 {
							goto Label1
						}
						goto Label2
						// line 109: if predicate is None:
						πF.SetLineno(109)
					Label1:
						// line 110: predicate = bool
						πF.SetLineno(110)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßbool); πE != nil {
							continue
						}
						µpredicate = πTemp001
						goto Label2
					Label2:
						if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.Iter(πF, µiterable); πE != nil {
							continue
						}
						πF.PushCheckpoint(4)
						πTemp003 = false
					Label3:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp003 {
							πF.PopCheckpoint()
							goto Label5
						}
						if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
							µx = πTemp002
						}
						if πE != nil || !πTemp004 {
							continue
						}
						πF.PushCheckpoint(3)            
						πTemp005 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
							continue
						}
						πTemp005[0] = µx
						if πE = πg.CheckLocal(πF, µpredicate, "predicate"); πE != nil {
							continue
						}
						if πTemp002, πE = µpredicate.Call(πF, πTemp005, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp005)
						if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
							continue
						}
						if πTemp004 {
							goto Label6
						}
						goto Label7
						// line 112: if predicate(x):
						πF.SetLineno(112)
					Label6:
						// line 113: yield x
						πF.SetLineno(113)
						if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
							continue
						}
						πF.PushCheckpoint(8)
						return µx, nil
					Label8:
						πTemp002 = πSent
						goto Label7
					Label7:
						continue
					Label4:
						if πE != nil || πR != nil {
							continue
						}
					Label5:
					}
					return nil, πE
				}).ToObject(), nil
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßifilter.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 116: def ifilterfalse(predicate, iterable):
			πF.SetLineno(116)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "predicate", Def: nil}
			πTemp006[1] = πg.Param{Name: "iterable", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("ifilterfalse", "build/src/__python__/itertools.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpredicate *πg.Object = πArgs[0]; _ = µpredicate
				var µiterable *πg.Object = πArgs[1]; _ = µiterable
				var µx *πg.Object = πg.UnboundLocal; _ = µx
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 8: goto Label8
						case 3: goto Label3
						case 4: goto Label4
						default: panic("unexpected function state")
						}
						if πE = πg.CheckLocal(πF, µpredicate, "predicate"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
							continue
						}
						πTemp001 = πg.GetBool(µpredicate == πTemp002).ToObject()
						if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
							continue
						}
						if πTemp003 {
							goto Label1
						}
						goto Label2
						// line 117: if predicate is None:
						πF.SetLineno(117)
					Label1:
						// line 118: predicate = bool
						πF.SetLineno(118)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßbool); πE != nil {
							continue
						}
						µpredicate = πTemp001
						goto Label2
					Label2:
						if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.Iter(πF, µiterable); πE != nil {
							continue
						}
						πF.PushCheckpoint(4)
						πTemp003 = false
					Label3:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp003 {
							πF.PopCheckpoint()
							goto Label5
						}
						if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
							µx = πTemp002
						}
						if πE != nil || !πTemp004 {
							continue
						}
						πF.PushCheckpoint(3)            
						πTemp005 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
							continue
						}
						πTemp005[0] = µx
						if πE = πg.CheckLocal(πF, µpredicate, "predicate"); πE != nil {
							continue
						}
						if πTemp006, πE = µpredicate.Call(πF, πTemp005, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp005)
						if πTemp004, πE = πg.IsTrue(πF, πTemp006); πE != nil {
							continue
						}
						πTemp002 = πg.GetBool(!πTemp004).ToObject()
						if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
							continue
						}
						if πTemp004 {
							goto Label6
						}
						goto Label7
						// line 120: if not predicate(x):
						πF.SetLineno(120)
					Label6:
						// line 121: yield x
						πF.SetLineno(121)
						if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
							continue
						}
						πF.PushCheckpoint(8)
						return µx, nil
					Label8:
						πTemp002 = πSent
						goto Label7
					Label7:
						continue
					Label4:
						if πE != nil || πR != nil {
							continue
						}
					Label5:
					}
					return nil, πE
				}).ToObject(), nil
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßifilterfalse.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 124: def imap(function, *iterables):
			πF.SetLineno(124)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "function", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("imap", "build/src/__python__/itertools.py", πTemp006, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfunction *πg.Object = πArgs[0]; _ = µfunction
				var µiterables *πg.Object = πArgs[1]; _ = µiterables
				var µargs *πg.Object = πg.UnboundLocal; _ = µargs
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 []πg.Param
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 8: goto Label8
						case 1: goto Label1
						case 2: goto Label2
						case 7: goto Label7
						default: panic("unexpected function state")
						}
						// line 125: iterables = map(iter, iterables)
						πF.SetLineno(125)
						πTemp001 = πF.MakeArgs(2)
						if πTemp002, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
							continue
						}
						πTemp001[0] = πTemp002
						if πE = πg.CheckLocal(πF, µiterables, "iterables"); πE != nil {
							continue
						}
						πTemp001[1] = µiterables
						if πTemp002, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						µiterables = πTemp003
						// line 126: while True:
						πF.SetLineno(126)
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
						if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
							continue
						}
						if πE != nil || !πTemp005 {
							continue
						}
						πF.PushCheckpoint(1)            
						// line 127: args = [next(it) for it in iterables]
						πF.SetLineno(127)
						πTemp006 = make([]πg.Param, 0)
						πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/itertools.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
							var µit *πg.Object = πg.UnboundLocal; _ = µit
							var πTemp001 *πg.Object
							_ = πTemp001
							var πTemp002 bool
							_ = πTemp002
							var πTemp003 bool
							_ = πTemp003
							var πTemp004 *πg.Object
							_ = πTemp004
							var πTemp005 []*πg.Object
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
									case 4: goto Label4
									default: panic("unexpected function state")
									}
									if πE = πg.CheckLocal(πF, µiterables, "iterables"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.Iter(πF, µiterables); πE != nil {
										continue
									}
									πF.PushCheckpoint(2)
									πTemp002 = false
								Label1:
									if πE != nil || πR != nil {
										continue
									}
									if πTemp002 {
										πF.PopCheckpoint()
										goto Label3
									}
									if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
										isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
										if exc != nil {
											πE = exc
										} else if isStop {
											πE = nil
											πF.RestoreExc(nil, nil)
										}
										πTemp003 = !isStop
									} else {
										πTemp003 = true
										µit = πTemp004
									}
									if πE != nil || !πTemp003 {
										continue
									}
									πF.PushCheckpoint(1)            
									// line 127: args = [next(it) for it in iterables]
									πF.SetLineno(127)
									πTemp005 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
										continue
									}
									πTemp005[0] = µit
									if πTemp004, πE = πg.ResolveGlobal(πF, ßnext); πE != nil {
										continue
									}
									if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp005)
									πF.PushCheckpoint(4)
									return πTemp006, nil
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
						if πTemp007, πE = πTemp003.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp007}, nil); πE != nil {
							continue
						}
						µargs = πTemp002
						if πE = πg.CheckLocal(πF, µfunction, "function"); πE != nil {
							continue
						}
						if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
							continue
						}
						πTemp002 = πg.GetBool(µfunction == πTemp007).ToObject()
						if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
							continue
						}
						if πTemp005 {
							goto Label4
						}
						goto Label5
						// line 128: if function is None:
						πF.SetLineno(128)
					Label4:
						// line 129: yield tuple(args)
						πF.SetLineno(129)
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
							continue
						}
						πTemp001[0] = µargs
						if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
							continue
						}
						if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						πF.PushCheckpoint(7)
						return πTemp007, nil
					Label7:
						πTemp002 = πSent
						goto Label6
					Label5:
						// line 131: yield function(*args)
						πF.SetLineno(131)
						if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µfunction, "function"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.Invoke(πF, µfunction, nil, µargs, nil, nil); πE != nil {
							continue
						}
						πF.PushCheckpoint(8)
						return πTemp002, nil
					Label8:
						πTemp008 = πSent
						goto Label6
					Label6:
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
			if πE = πF.Globals().SetItem(πF, ßimap.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 134: def islice(iterable, *args):
			πF.SetLineno(134)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "iterable", Def: nil}
			πTemp011 = πg.NewFunction(πg.NewCode("islice", "build/src/__python__/itertools.py", πTemp006, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µiterable *πg.Object = πArgs[0]; _ = µiterable
				var µargs *πg.Object = πArgs[1]; _ = µargs
				var µs *πg.Object = πg.UnboundLocal; _ = µs
				var µit *πg.Object = πg.UnboundLocal; _ = µit
				var µnexti *πg.Object = πg.UnboundLocal; _ = µnexti
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µelement *πg.Object = πg.UnboundLocal; _ = µelement
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 9: goto Label9
						case 4: goto Label4
						case 5: goto Label5
						default: panic("unexpected function state")
						}
						// line 135: s = slice(*args)
						πF.SetLineno(135)
						if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
							continue
						}
						if πTemp002, πE = πg.Invoke(πF, πTemp001, nil, µargs, nil, nil); πE != nil {
							continue
						}
						µs = πTemp002
						// line 136: it = iter(xrange(s.start or 0, s.stop or sys.maxint, s.step or 1))
						πF.SetLineno(136)
						πTemp003 = πF.MakeArgs(1)
						πTemp004 = πF.MakeArgs(3)
						if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µs, ßstart, nil); πE != nil {
							continue
						}
						πTemp001 = πTemp002
						if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
							continue
						}
						if πTemp005 {
							goto Label1
						}
						πTemp001 = πg.NewInt(0).ToObject()
					Label1:
						πTemp004[0] = πTemp001
						if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µs, ßstop, nil); πE != nil {
							continue
						}
						πTemp001 = πTemp002
						if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
							continue
						}
						if πTemp005 {
							goto Label2
						}
						if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
							continue
						}
						if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßmaxint, nil); πE != nil {
							continue
						}
						πTemp001 = πTemp006
					Label2:
						πTemp004[1] = πTemp001
						if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µs, ßstep, nil); πE != nil {
							continue
						}
						πTemp001 = πTemp002
						if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
							continue
						}
						if πTemp005 {
							goto Label3
						}
						πTemp001 = πg.NewInt(1).ToObject()
					Label3:
						πTemp004[2] = πTemp001
						if πTemp001, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp004)
						πTemp003[0] = πTemp002
						if πTemp001, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp003)
						µit = πTemp002
						// line 137: nexti = next(it)
						πF.SetLineno(137)
						πTemp003 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
							continue
						}
						πTemp003[0] = µit
						if πTemp001, πE = πg.ResolveGlobal(πF, ßnext); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp003)
						µnexti = πTemp002
						πTemp003 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
							continue
						}
						πTemp003[0] = µiterable
						if πTemp002, πE = πg.ResolveGlobal(πF, ßenumerate); πE != nil {
							continue
						}
						if πTemp006, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp003)
						if πTemp001, πE = πg.Iter(πF, πTemp006); πE != nil {
							continue
						}
						πF.PushCheckpoint(5)
						πTemp005 = false
					Label4:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp005 {
							πF.PopCheckpoint()
							goto Label6
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
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp008}}}, πTemp002); πE != nil {
								continue
							}
							µi = πTemp006
							µelement = πTemp008
						}
						if πE != nil || !πTemp007 {
							continue
						}
						πF.PushCheckpoint(4)            
						if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µnexti, "nexti"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.Eq(πF, µi, µnexti); πE != nil {
							continue
						}
						if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
							continue
						}
						if πTemp007 {
							goto Label7
						}
						goto Label8
						// line 139: if i == nexti:
						πF.SetLineno(139)
					Label7:
						// line 140: yield element
						πF.SetLineno(140)
						if πE = πg.CheckLocal(πF, µelement, "element"); πE != nil {
							continue
						}
						πF.PushCheckpoint(9)
						return µelement, nil
					Label9:
						πTemp002 = πSent
						// line 141: nexti = next(it)
						πF.SetLineno(141)
						πTemp003 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
							continue
						}
						πTemp003[0] = µit
						if πTemp002, πE = πg.ResolveGlobal(πF, ßnext); πE != nil {
							continue
						}
						if πTemp006, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp003)
						µnexti = πTemp006
						goto Label8
					Label8:
						continue
					Label5:
						if πE != nil || πR != nil {
							continue
						}
					Label6:
					}
					return nil, πE
				}).ToObject(), nil
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßislice.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 144: def izip(*iterables):
			πF.SetLineno(144)
			πTemp006 = make([]πg.Param, 0)
			πTemp012 = πg.NewFunction(πg.NewCode("izip", "build/src/__python__/itertools.py", πTemp006, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µiterables *πg.Object = πArgs[0]; _ = µiterables
				var µiterators *πg.Object = πg.UnboundLocal; _ = µiterators
				var πTemp001 []*πg.Object
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
						// line 145: iterators = map(iter, iterables)
						πF.SetLineno(145)
						πTemp001 = πF.MakeArgs(2)
						if πTemp002, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
							continue
						}
						πTemp001[0] = πTemp002
						if πE = πg.CheckLocal(πF, µiterables, "iterables"); πE != nil {
							continue
						}
						πTemp001[1] = µiterables
						if πTemp002, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						µiterators = πTemp003
						// line 146: while iterators:
						πF.SetLineno(146)
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
						if πE = πg.CheckLocal(πF, µiterators, "iterators"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.IsTrue(πF, µiterators); πE != nil {
							continue
						}
						if πE != nil || !πTemp005 {
							continue
						}
						πF.PushCheckpoint(1)            
						// line 147: yield tuple(map(next, iterators))
						πF.SetLineno(147)
						πTemp001 = πF.MakeArgs(1)
						πTemp006 = πF.MakeArgs(2)
						if πTemp002, πE = πg.ResolveGlobal(πF, ßnext); πE != nil {
							continue
						}
						πTemp006[0] = πTemp002
						if πE = πg.CheckLocal(πF, µiterators, "iterators"); πE != nil {
							continue
						}
						πTemp006[1] = µiterators
						if πTemp002, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp006)
						πTemp001[0] = πTemp003
						if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						πF.PushCheckpoint(4)
						return πTemp003, nil
					Label4:
						πTemp002 = πSent
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
			if πE = πF.Globals().SetItem(πF, ßizip.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 150: class ZipExhausted(Exception):
			πF.SetLineno(150)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp015, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
				continue
			}
			πTemp002[0] = πTemp015
			πTemp003 = πg.NewDict()
			if πTemp013, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp013); πE != nil {
				continue
			}
			_, πE = πg.NewCode("ZipExhausted", "build/src/__python__/itertools.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 151: pass
					πF.SetLineno(151)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp014, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp014 == nil {
				πTemp014 = πg.TypeType.ToObject()
			}
			if πTemp015, πE = πTemp014.Call(πF, []*πg.Object{πg.NewStr("ZipExhausted").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßZipExhausted.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 154: def izip_longest(*args, **kwds):
			πF.SetLineno(154)
			πTemp006 = make([]πg.Param, 0)
			πTemp013 = πg.NewFunction(πg.NewCode("izip_longest", "build/src/__python__/itertools.py", πTemp006, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargs *πg.Object = πArgs[0]; _ = µargs
				var µkwds *πg.Object = πArgs[1]; _ = µkwds
				var µfillvalue *πg.Object = πg.UnboundLocal; _ = µfillvalue
				var µcounter *πg.Object = πg.UnboundLocal; _ = µcounter
				var µsentinel *πg.Object = πg.UnboundLocal; _ = µsentinel
				var µfillers *πg.Object = πg.UnboundLocal; _ = µfillers
				var µiterators *πg.Object = πg.UnboundLocal; _ = µiterators
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 []πg.Param
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πTemp010 *πg.BaseException
				_ = πTemp010
				var πTemp011 *πg.Traceback
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 2: goto Label2
						case 3: goto Label3
						case 4: goto Label4
						case 6: goto Label6
						default: panic("unexpected function state")
						}
						// line 156: fillvalue = kwds.get('fillvalue')
						πF.SetLineno(156)
						πTemp001 = πF.MakeArgs(1)
						πTemp001[0] = ßfillvalue.ToObject()
						if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µkwds, ßget, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						µfillvalue = πTemp003
						// line 157: counter = [len(args) - 1]
						πF.SetLineno(157)
						πTemp001 = make([]*πg.Object, 1)
						πTemp004 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
							continue
						}
						πTemp004[0] = µargs
						if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp004)
						if πTemp002, πE = πg.Sub(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
							continue
						}
						πTemp001[0] = πTemp002
						πTemp002 = πg.NewList(πTemp001...).ToObject()
						µcounter = πTemp002
						// line 158: def sentinel():
						πF.SetLineno(158)
						πTemp006 = make([]πg.Param, 0)
						πTemp002 = πg.NewFunction(πg.NewCode("sentinel", "build/src/__python__/itertools.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
							var πTemp001 *πg.Object
							_ = πTemp001
							var πTemp002 *πg.Object
							_ = πTemp002
							var πTemp003 *πg.Object
							_ = πTemp003
							var πTemp004 bool
							_ = πTemp004
							var πR *πg.Object; _ = πR
							var πE *πg.BaseException; _ = πE
							return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									case 3: goto Label3
									default: panic("unexpected function state")
									}
									πTemp002 = πg.NewInt(0).ToObject()
									if πE = πg.CheckLocal(πF, µcounter, "counter"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetItem(πF, µcounter, πTemp002); πE != nil {
										continue
									}
									if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
										continue
									}
									πTemp001 = πg.GetBool(!πTemp004).ToObject()
									if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp004 {
										goto Label1
									}
									goto Label2
									// line 159: if not counter[0]:
									πF.SetLineno(159)
								Label1:
									if πTemp001, πE = πg.ResolveGlobal(πF, ßZipExhausted); πE != nil {
										continue
									}
									// line 160: raise ZipExhausted
									πF.SetLineno(160)
									πE = πF.Raise(πTemp001, nil, nil)
									continue
									goto Label2
								Label2:
									// line 161: counter[0] -= 1
									πF.SetLineno(161)
									πTemp001 = πg.NewInt(0).ToObject()
									if πE = πg.CheckLocal(πF, µcounter, "counter"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetItem(πF, µcounter, πTemp001); πE != nil {
										continue
									}
									if πTemp001, πE = πg.ISub(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µcounter, "counter"); πE != nil {
										continue
									}
									πTemp003 = πg.NewInt(0).ToObject()
									if πE = πg.SetItem(πF, µcounter, πTemp003, πTemp001); πE != nil {
										continue
									}
									// line 162: yield fillvalue
									πF.SetLineno(162)
									if πE = πg.CheckLocal(πF, µfillvalue, "fillvalue"); πE != nil {
										continue
									}
									πF.PushCheckpoint(3)
									return µfillvalue, nil
								Label3:
									πTemp001 = πSent
								}
								return nil, πE
							}).ToObject(), nil
						}), πF.Globals()).ToObject()
						µsentinel = πTemp002
						// line 163: fillers = repeat(fillvalue)
						πF.SetLineno(163)
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µfillvalue, "fillvalue"); πE != nil {
							continue
						}
						πTemp001[0] = µfillvalue
						if πTemp003, πE = πg.ResolveGlobal(πF, ßrepeat); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						µfillers = πTemp005
						// line 164: iterators = [chain(it, sentinel(), fillers) for it in args]
						πF.SetLineno(164)
						πTemp006 = make([]πg.Param, 0)
						πTemp005 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/itertools.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
							var µit *πg.Object = πg.UnboundLocal; _ = µit
							var πTemp001 *πg.Object
							_ = πTemp001
							var πTemp002 bool
							_ = πTemp002
							var πTemp003 bool
							_ = πTemp003
							var πTemp004 *πg.Object
							_ = πTemp004
							var πTemp005 []*πg.Object
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
									case 4: goto Label4
									default: panic("unexpected function state")
									}
									if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.Iter(πF, µargs); πE != nil {
										continue
									}
									πF.PushCheckpoint(2)
									πTemp002 = false
								Label1:
									if πE != nil || πR != nil {
										continue
									}
									if πTemp002 {
										πF.PopCheckpoint()
										goto Label3
									}
									if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
										isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
										if exc != nil {
											πE = exc
										} else if isStop {
											πE = nil
											πF.RestoreExc(nil, nil)
										}
										πTemp003 = !isStop
									} else {
										πTemp003 = true
										µit = πTemp004
									}
									if πE != nil || !πTemp003 {
										continue
									}
									πF.PushCheckpoint(1)            
									// line 164: iterators = [chain(it, sentinel(), fillers) for it in args]
									πF.SetLineno(164)
									πTemp005 = πF.MakeArgs(3)
									if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
										continue
									}
									πTemp005[0] = µit
									if πE = πg.CheckLocal(πF, µsentinel, "sentinel"); πE != nil {
										continue
									}
									if πTemp004, πE = µsentinel.Call(πF, nil, nil); πE != nil {
										continue
									}
									πTemp005[1] = πTemp004
									if πE = πg.CheckLocal(πF, µfillers, "fillers"); πE != nil {
										continue
									}
									πTemp005[2] = µfillers
									if πTemp004, πE = πg.ResolveGlobal(πF, ßchain); πE != nil {
										continue
									}
									if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp005)
									πF.PushCheckpoint(4)
									return πTemp006, nil
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
						if πTemp007, πE = πTemp005.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp007}, nil); πE != nil {
							continue
						}
						µiterators = πTemp003
						// line 165: try:
						πF.SetLineno(165)
						πF.PushCheckpoint(2)
						// line 166: while iterators:
						πF.SetLineno(166)
						πF.PushCheckpoint(4)
						πTemp008 = false
					Label3:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp008 {
							πF.PopCheckpoint()
							goto Label5
						}
						if πE = πg.CheckLocal(πF, µiterators, "iterators"); πE != nil {
							continue
						}
						if πTemp009, πE = πg.IsTrue(πF, µiterators); πE != nil {
							continue
						}
						if πE != nil || !πTemp009 {
							continue
						}
						πF.PushCheckpoint(3)            
						// line 167: yield tuple(map(next, iterators))
						πF.SetLineno(167)
						πTemp001 = πF.MakeArgs(1)
						πTemp004 = πF.MakeArgs(2)
						if πTemp003, πE = πg.ResolveGlobal(πF, ßnext); πE != nil {
							continue
						}
						πTemp004[0] = πTemp003
						if πE = πg.CheckLocal(πF, µiterators, "iterators"); πE != nil {
							continue
						}
						πTemp004[1] = µiterators
						if πTemp003, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
							continue
						}
						if πTemp007, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp004)
						πTemp001[0] = πTemp007
						if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
							continue
						}
						if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						πF.PushCheckpoint(6)
						return πTemp007, nil
					Label6:
						πTemp003 = πSent
						continue
					Label4:
						if πE != nil || πR != nil {
							continue
						}
					Label5:
						πF.PopCheckpoint()
						goto Label1
					Label2:
						if πE == nil {
						  continue
						}
						πE = nil
						πTemp010, πTemp011 = πF.ExcInfo()
						if πTemp003, πE = πg.ResolveGlobal(πF, ßZipExhausted); πE != nil {
							continue
						}
						if πTemp008, πE = πg.IsInstance(πF, πTemp010.ToObject(), πTemp003); πE != nil {
							continue
						}
						if πTemp008 {
							goto Label7
						}
						πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
						continue
						// line 168: except ZipExhausted:
						πF.SetLineno(168)
					Label7:
						// line 169: pass
						πF.SetLineno(169)
						πF.RestoreExc(nil, nil)
						goto Label1
					Label1:
					}
					return nil, πE
				}).ToObject(), nil
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßizip_longest.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 172: def product(*args, **kwds):
			πF.SetLineno(172)
			πTemp006 = make([]πg.Param, 0)
			πTemp014 = πg.NewFunction(πg.NewCode("product", "build/src/__python__/itertools.py", πTemp006, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargs *πg.Object = πArgs[0]; _ = µargs
				var µkwds *πg.Object = πArgs[1]; _ = µkwds
				var µpools *πg.Object = πg.UnboundLocal; _ = µpools
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µpool *πg.Object = πg.UnboundLocal; _ = µpool
				var µprod *πg.Object = πg.UnboundLocal; _ = µprod
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
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 []πg.Param
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 1: goto Label1
						case 2: goto Label2
						case 4: goto Label4
						case 5: goto Label5
						case 7: goto Label7
						default: panic("unexpected function state")
						}
						// line 175: pools = map(tuple, args) * kwds.get('repeat', 1)
						πF.SetLineno(175)
						πTemp002 = πF.MakeArgs(2)
						if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
							continue
						}
						πTemp002[0] = πTemp003
						if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
							continue
						}
						πTemp002[1] = µargs
						if πTemp003, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
							continue
						}
						if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp002)
						πTemp002 = πF.MakeArgs(2)
						πTemp002[0] = ßrepeat.ToObject()
						πTemp002[1] = πg.NewInt(1).ToObject()
						if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, µkwds, ßget, nil); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp002)
						if πTemp001, πE = πg.Mul(πF, πTemp004, πTemp005); πE != nil {
							continue
						}
						µpools = πTemp001
						// line 176: result = [[]]
						πF.SetLineno(176)
						πTemp002 = make([]*πg.Object, 1)
						πTemp006 = make([]*πg.Object, 0)
						πTemp001 = πg.NewList(πTemp006...).ToObject()
						πTemp002[0] = πTemp001
						πTemp001 = πg.NewList(πTemp002...).ToObject()
						µresult = πTemp001
						if πE = πg.CheckLocal(πF, µpools, "pools"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.Iter(πF, µpools); πE != nil {
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
						if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
							µpool = πTemp003
						}
						if πE != nil || !πTemp008 {
							continue
						}
						πF.PushCheckpoint(1)            
						// line 178: result = [x+[y] for x in result for y in pool]
						πF.SetLineno(178)
						πTemp009 = make([]πg.Param, 0)
						πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/itertools.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
							var µx *πg.Object = πg.UnboundLocal; _ = µx
							var µy *πg.Object = πg.UnboundLocal; _ = µy
							var πTemp001 *πg.Object
							_ = πTemp001
							var πTemp002 bool
							_ = πTemp002
							var πTemp003 bool
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
							var πR *πg.Object; _ = πR
							var πE *πg.BaseException; _ = πE
							return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									case 1: goto Label1
									case 2: goto Label2
									case 4: goto Label4
									case 5: goto Label5
									case 7: goto Label7
									default: panic("unexpected function state")
									}
									if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.Iter(πF, µresult); πE != nil {
										continue
									}
									πF.PushCheckpoint(2)
									πTemp002 = false
								Label1:
									if πE != nil || πR != nil {
										continue
									}
									if πTemp002 {
										πF.PopCheckpoint()
										goto Label3
									}
									if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
										isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
										if exc != nil {
											πE = exc
										} else if isStop {
											πE = nil
											πF.RestoreExc(nil, nil)
										}
										πTemp003 = !isStop
									} else {
										πTemp003 = true
										µx = πTemp004
									}
									if πE != nil || !πTemp003 {
										continue
									}
									πF.PushCheckpoint(1)            
									if πE = πg.CheckLocal(πF, µpool, "pool"); πE != nil {
										continue
									}
									if πTemp004, πE = πg.Iter(πF, µpool); πE != nil {
										continue
									}
									πF.PushCheckpoint(5)
									πTemp003 = false
								Label4:
									if πE != nil || πR != nil {
										continue
									}
									if πTemp003 {
										πF.PopCheckpoint()
										goto Label6
									}
									if πTemp006, πE = πg.Next(πF, πTemp004); πE != nil {
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
										µy = πTemp006
									}
									if πE != nil || !πTemp005 {
										continue
									}
									πF.PushCheckpoint(4)            
									// line 178: result = [x+[y] for x in result for y in pool]
									πF.SetLineno(178)
									if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
										continue
									}
									πTemp007 = make([]*πg.Object, 1)
									if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
										continue
									}
									πTemp007[0] = µy
									πTemp008 = πg.NewList(πTemp007...).ToObject()
									if πTemp006, πE = πg.Add(πF, µx, πTemp008); πE != nil {
										continue
									}
									πF.PushCheckpoint(7)
									return πTemp006, nil
								Label7:
									πTemp008 = πSent
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
								}
								return nil, πE
							}).ToObject(), nil
						}), πF.Globals()).ToObject()
						if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
							continue
						}
						µresult = πTemp003
						continue
					Label2:
						if πE != nil || πR != nil {
							continue
						}
					Label3:
						if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.Iter(πF, µresult); πE != nil {
							continue
						}
						πF.PushCheckpoint(5)
						πTemp007 = false
					Label4:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp007 {
							πF.PopCheckpoint()
							goto Label6
						}
						if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
							µprod = πTemp003
						}
						if πE != nil || !πTemp008 {
							continue
						}
						πF.PushCheckpoint(4)            
						// line 180: yield tuple(prod)
						πF.SetLineno(180)
						πTemp002 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µprod, "prod"); πE != nil {
							continue
						}
						πTemp002[0] = µprod
						if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp002)
						πF.PushCheckpoint(7)
						return πTemp005, nil
					Label7:
						πTemp003 = πSent
						continue
					Label5:
						if πE != nil || πR != nil {
							continue
						}
					Label6:
					}
					return nil, πE
				}).ToObject(), nil
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßproduct.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 183: def permutations(iterable, r=None):
			πF.SetLineno(183)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "iterable", Def: nil}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "r", Def: πTemp016}
			πTemp015 = πg.NewFunction(πg.NewCode("permutations", "build/src/__python__/itertools.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µiterable *πg.Object = πArgs[0]; _ = µiterable
				var µr *πg.Object = πArgs[1]; _ = µr
				var µpool *πg.Object = πg.UnboundLocal; _ = µpool
				var µn *πg.Object = πg.UnboundLocal; _ = µn
				var µindices *πg.Object = πg.UnboundLocal; _ = µindices
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
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 πg.KWArgs
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 []πg.Param
				_ = πTemp010
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 8: goto Label8
						case 3: goto Label3
						case 4: goto Label4
						default: panic("unexpected function state")
						}
						// line 184: pool = tuple(iterable)
						πF.SetLineno(184)
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
							continue
						}
						πTemp001[0] = µiterable
						if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						µpool = πTemp003
						// line 185: n = len(pool)
						πF.SetLineno(185)
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µpool, "pool"); πE != nil {
							continue
						}
						πTemp001[0] = µpool
						if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						µn = πTemp003
						// line 186: r = n if r is None else r
						πF.SetLineno(186)
						if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
							continue
						}
						if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
							continue
						}
						πTemp003 = πg.GetBool(µr == πTemp004).ToObject()
						if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
							continue
						}
						if !πTemp005 {
							goto Label1
						}
						if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
							continue
						}
						πTemp002 = µn
						goto Label2
					Label1:
						if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
							continue
						}
						πTemp002 = µr
					Label2:
						µr = πTemp002
						πTemp001 = πF.MakeArgs(1)
						πTemp006 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
							continue
						}
						πTemp006[0] = µn
						if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
							continue
						}
						if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp006)
						πTemp001[0] = πTemp004
						if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
							continue
						}
						πTemp007 = πg.KWArgs{
							{"repeat", µr},
						}
						if πTemp003, πE = πg.ResolveGlobal(πF, ßproduct); πE != nil {
							continue
						}
						if πTemp004, πE = πTemp003.Call(πF, πTemp001, πTemp007); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
							continue
						}
						πF.PushCheckpoint(4)
						πTemp005 = false
					Label3:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp005 {
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
							πTemp008 = !isStop
						} else {
							πTemp008 = true
							µindices = πTemp003
						}
						if πE != nil || !πTemp008 {
							continue
						}
						πF.PushCheckpoint(3)            
						πTemp001 = πF.MakeArgs(1)
						πTemp006 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µindices, "indices"); πE != nil {
							continue
						}
						πTemp006[0] = µindices
						if πTemp004, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
							continue
						}
						if πTemp009, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp006)
						πTemp001[0] = πTemp009
						if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
							continue
						}
						if πTemp009, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.Eq(πF, πTemp009, µr); πE != nil {
							continue
						}
						if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
							continue
						}
						if πTemp008 {
							goto Label6
						}
						goto Label7
						// line 188: if len(set(indices)) == r:
						πF.SetLineno(188)
					Label6:
						// line 189: yield tuple(pool[i] for i in indices)
						πF.SetLineno(189)
						πTemp001 = πF.MakeArgs(1)
						πTemp010 = make([]πg.Param, 0)
						πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/itertools.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
							var µi *πg.Object = πg.UnboundLocal; _ = µi
							var πTemp001 *πg.Object
							_ = πTemp001
							var πTemp002 bool
							_ = πTemp002
							var πTemp003 bool
							_ = πTemp003
							var πTemp004 *πg.Object
							_ = πTemp004
							var πTemp005 *πg.Object
							_ = πTemp005
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
									if πE = πg.CheckLocal(πF, µindices, "indices"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.Iter(πF, µindices); πE != nil {
										continue
									}
									πF.PushCheckpoint(2)
									πTemp002 = false
								Label1:
									if πE != nil || πR != nil {
										continue
									}
									if πTemp002 {
										πF.PopCheckpoint()
										goto Label3
									}
									if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
										isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
										if exc != nil {
											πE = exc
										} else if isStop {
											πE = nil
											πF.RestoreExc(nil, nil)
										}
										πTemp003 = !isStop
									} else {
										πTemp003 = true
										µi = πTemp004
									}
									if πE != nil || !πTemp003 {
										continue
									}
									πF.PushCheckpoint(1)            
									// line 189: yield tuple(pool[i] for i in indices)
									πF.SetLineno(189)
									if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
										continue
									}
									πTemp004 = µi
									if πE = πg.CheckLocal(πF, µpool, "pool"); πE != nil {
										continue
									}
									if πTemp005, πE = πg.GetItem(πF, µpool, πTemp004); πE != nil {
										continue
									}
									πF.PushCheckpoint(4)
									return πTemp005, nil
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
						if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
							continue
						}
						πTemp001[0] = πTemp004
						if πTemp004, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
							continue
						}
						if πTemp009, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						πF.PushCheckpoint(8)
						return πTemp009, nil
					Label8:
						πTemp004 = πSent
						goto Label7
					Label7:
						continue
					Label4:
						if πE != nil || πR != nil {
							continue
						}
					Label5:
					}
					return nil, πE
				}).ToObject(), nil
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßpermutations.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 192: def combinations(iterable, r):
			πF.SetLineno(192)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "iterable", Def: nil}
			πTemp006[1] = πg.Param{Name: "r", Def: nil}
			πTemp016 = πg.NewFunction(πg.NewCode("combinations", "build/src/__python__/itertools.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µiterable *πg.Object = πArgs[0]; _ = µiterable
				var µr *πg.Object = πArgs[1]; _ = µr
				var µpool *πg.Object = πg.UnboundLocal; _ = µpool
				var µn *πg.Object = πg.UnboundLocal; _ = µn
				var µindices *πg.Object = πg.UnboundLocal; _ = µindices
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 []πg.Param
				_ = πTemp010
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
						// line 193: pool = tuple(iterable)
						πF.SetLineno(193)
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
							continue
						}
						πTemp001[0] = µiterable
						if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						µpool = πTemp003
						// line 194: n = len(pool)
						πF.SetLineno(194)
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µpool, "pool"); πE != nil {
							continue
						}
						πTemp001[0] = µpool
						if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						µn = πTemp003
						πTemp001 = πF.MakeArgs(2)
						πTemp004 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
							continue
						}
						πTemp004[0] = µn
						if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp004)
						πTemp001[0] = πTemp005
						if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
							continue
						}
						πTemp001[1] = µr
						if πTemp003, πE = πg.ResolveGlobal(πF, ßpermutations); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
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
							µindices = πTemp003
						}
						if πE != nil || !πTemp007 {
							continue
						}
						πF.PushCheckpoint(1)            
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µindices, "indices"); πE != nil {
							continue
						}
						πTemp001[0] = µindices
						if πTemp005, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
							continue
						}
						if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µindices, "indices"); πE != nil {
							continue
						}
						πTemp001[0] = µindices
						if πTemp005, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
							continue
						}
						if πTemp009, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πTemp003, πE = πg.Eq(πF, πTemp008, πTemp009); πE != nil {
							continue
						}
						if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
							continue
						}
						if πTemp007 {
							goto Label4
						}
						goto Label5
						// line 196: if sorted(indices) == list(indices):
						πF.SetLineno(196)
					Label4:
						// line 197: yield tuple(pool[i] for i in indices)
						πF.SetLineno(197)
						πTemp001 = πF.MakeArgs(1)
						πTemp010 = make([]πg.Param, 0)
						πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/itertools.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
							var µi *πg.Object = πg.UnboundLocal; _ = µi
							var πTemp001 *πg.Object
							_ = πTemp001
							var πTemp002 bool
							_ = πTemp002
							var πTemp003 bool
							_ = πTemp003
							var πTemp004 *πg.Object
							_ = πTemp004
							var πTemp005 *πg.Object
							_ = πTemp005
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
									if πE = πg.CheckLocal(πF, µindices, "indices"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.Iter(πF, µindices); πE != nil {
										continue
									}
									πF.PushCheckpoint(2)
									πTemp002 = false
								Label1:
									if πE != nil || πR != nil {
										continue
									}
									if πTemp002 {
										πF.PopCheckpoint()
										goto Label3
									}
									if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
										isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
										if exc != nil {
											πE = exc
										} else if isStop {
											πE = nil
											πF.RestoreExc(nil, nil)
										}
										πTemp003 = !isStop
									} else {
										πTemp003 = true
										µi = πTemp004
									}
									if πE != nil || !πTemp003 {
										continue
									}
									πF.PushCheckpoint(1)            
									// line 197: yield tuple(pool[i] for i in indices)
									πF.SetLineno(197)
									if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
										continue
									}
									πTemp004 = µi
									if πE = πg.CheckLocal(πF, µpool, "pool"); πE != nil {
										continue
									}
									if πTemp005, πE = πg.GetItem(πF, µpool, πTemp004); πE != nil {
										continue
									}
									πF.PushCheckpoint(4)
									return πTemp005, nil
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
						if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
							continue
						}
						πTemp001[0] = πTemp005
						if πTemp005, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
							continue
						}
						if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						πF.PushCheckpoint(6)
						return πTemp008, nil
					Label6:
						πTemp005 = πSent
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
			if πE = πF.Globals().SetItem(πF, ßcombinations.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 200: def combinations_with_replacement(iterable, r):
			πF.SetLineno(200)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "iterable", Def: nil}
			πTemp006[1] = πg.Param{Name: "r", Def: nil}
			πTemp017 = πg.NewFunction(πg.NewCode("combinations_with_replacement", "build/src/__python__/itertools.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µiterable *πg.Object = πArgs[0]; _ = µiterable
				var µr *πg.Object = πArgs[1]; _ = µr
				var µpool *πg.Object = πg.UnboundLocal; _ = µpool
				var µn *πg.Object = πg.UnboundLocal; _ = µn
				var µindices *πg.Object = πg.UnboundLocal; _ = µindices
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 πg.KWArgs
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 []πg.Param
				_ = πTemp011
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
						// line 201: pool = tuple(iterable)
						πF.SetLineno(201)
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
							continue
						}
						πTemp001[0] = µiterable
						if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						µpool = πTemp003
						// line 202: n = len(pool)
						πF.SetLineno(202)
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µpool, "pool"); πE != nil {
							continue
						}
						πTemp001[0] = µpool
						if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						µn = πTemp003
						πTemp001 = πF.MakeArgs(1)
						πTemp004 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
							continue
						}
						πTemp004[0] = µn
						if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp004)
						πTemp001[0] = πTemp005
						if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
							continue
						}
						πTemp006 = πg.KWArgs{
							{"repeat", µr},
						}
						if πTemp003, πE = πg.ResolveGlobal(πF, ßproduct); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp003.Call(πF, πTemp001, πTemp006); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
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
							µindices = πTemp003
						}
						if πE != nil || !πTemp008 {
							continue
						}
						πF.PushCheckpoint(1)            
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µindices, "indices"); πE != nil {
							continue
						}
						πTemp001[0] = µindices
						if πTemp005, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
							continue
						}
						if πTemp009, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µindices, "indices"); πE != nil {
							continue
						}
						πTemp001[0] = µindices
						if πTemp005, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
							continue
						}
						if πTemp010, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πTemp003, πE = πg.Eq(πF, πTemp009, πTemp010); πE != nil {
							continue
						}
						if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
							continue
						}
						if πTemp008 {
							goto Label4
						}
						goto Label5
						// line 204: if sorted(indices) == list(indices):
						πF.SetLineno(204)
					Label4:
						// line 205: yield tuple(pool[i] for i in indices)
						πF.SetLineno(205)
						πTemp001 = πF.MakeArgs(1)
						πTemp011 = make([]πg.Param, 0)
						πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/itertools.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
							var µi *πg.Object = πg.UnboundLocal; _ = µi
							var πTemp001 *πg.Object
							_ = πTemp001
							var πTemp002 bool
							_ = πTemp002
							var πTemp003 bool
							_ = πTemp003
							var πTemp004 *πg.Object
							_ = πTemp004
							var πTemp005 *πg.Object
							_ = πTemp005
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
									if πE = πg.CheckLocal(πF, µindices, "indices"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.Iter(πF, µindices); πE != nil {
										continue
									}
									πF.PushCheckpoint(2)
									πTemp002 = false
								Label1:
									if πE != nil || πR != nil {
										continue
									}
									if πTemp002 {
										πF.PopCheckpoint()
										goto Label3
									}
									if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
										isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
										if exc != nil {
											πE = exc
										} else if isStop {
											πE = nil
											πF.RestoreExc(nil, nil)
										}
										πTemp003 = !isStop
									} else {
										πTemp003 = true
										µi = πTemp004
									}
									if πE != nil || !πTemp003 {
										continue
									}
									πF.PushCheckpoint(1)            
									// line 205: yield tuple(pool[i] for i in indices)
									πF.SetLineno(205)
									if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
										continue
									}
									πTemp004 = µi
									if πE = πg.CheckLocal(πF, µpool, "pool"); πE != nil {
										continue
									}
									if πTemp005, πE = πg.GetItem(πF, µpool, πTemp004); πE != nil {
										continue
									}
									πF.PushCheckpoint(4)
									return πTemp005, nil
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
						if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
							continue
						}
						πTemp001[0] = πTemp005
						if πTemp005, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
							continue
						}
						if πTemp009, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						πF.PushCheckpoint(6)
						return πTemp009, nil
					Label6:
						πTemp005 = πSent
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
			if πE = πF.Globals().SetItem(πF, ßcombinations_with_replacement.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 208: def repeat(object, times=None):
			πF.SetLineno(208)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "object", Def: nil}
			if πTemp019, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "times", Def: πTemp019}
			πTemp018 = πg.NewFunction(πg.NewCode("repeat", "build/src/__python__/itertools.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobject *πg.Object = πArgs[0]; _ = µobject
				var µtimes *πg.Object = πArgs[1]; _ = µtimes
				var µi *πg.Object = πg.UnboundLocal; _ = µi
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 4: goto Label4
						case 5: goto Label5
						case 7: goto Label7
						case 8: goto Label8
						case 9: goto Label9
						case 11: goto Label11
						default: panic("unexpected function state")
						}
						if πE = πg.CheckLocal(πF, µtimes, "times"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
							continue
						}
						πTemp001 = πg.GetBool(µtimes == πTemp002).ToObject()
						if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
							continue
						}
						if πTemp003 {
							goto Label1
						}
						goto Label2
						// line 209: if times is None:
						πF.SetLineno(209)
					Label1:
						// line 210: while True:
						πF.SetLineno(210)
						πF.PushCheckpoint(5)
						πTemp003 = false
					Label4:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp003 {
							πF.PopCheckpoint()
							goto Label6
						}
						if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
							continue
						}
						if πE != nil || !πTemp004 {
							continue
						}
						πF.PushCheckpoint(4)            
						// line 211: yield object
						πF.SetLineno(211)
						if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
							continue
						}
						πF.PushCheckpoint(7)
						return µobject, nil
					Label7:
						πTemp001 = πSent
						continue
					Label5:
						if πE != nil || πR != nil {
							continue
						}
					Label6:
						goto Label3
					Label2:
						πTemp005 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µtimes, "times"); πE != nil {
							continue
						}
						πTemp005[0] = µtimes
						if πTemp002, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
							continue
						}
						if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp005)
						if πTemp001, πE = πg.Iter(πF, πTemp006); πE != nil {
							continue
						}
						πF.PushCheckpoint(9)
						πTemp003 = false
					Label8:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp003 {
							πF.PopCheckpoint()
							goto Label10
						}
						if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
							µi = πTemp002
						}
						if πE != nil || !πTemp004 {
							continue
						}
						πF.PushCheckpoint(8)            
						// line 214: yield object
						πF.SetLineno(214)
						if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
							continue
						}
						πF.PushCheckpoint(11)
						return µobject, nil
					Label11:
						πTemp002 = πSent
						continue
					Label9:
						if πE != nil || πR != nil {
							continue
						}
					Label10:
						goto Label3
					Label3:
					}
					return nil, πE
				}).ToObject(), nil
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßrepeat.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 217: def starmap(function, iterable):
			πF.SetLineno(217)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "function", Def: nil}
			πTemp006[1] = πg.Param{Name: "iterable", Def: nil}
			πTemp019 = πg.NewFunction(πg.NewCode("starmap", "build/src/__python__/itertools.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfunction *πg.Object = πArgs[0]; _ = µfunction
				var µiterable *πg.Object = πArgs[1]; _ = µiterable
				var µargs *πg.Object = πg.UnboundLocal; _ = µargs
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
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
						if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.Iter(πF, µiterable); πE != nil {
							continue
						}
						πF.PushCheckpoint(2)
						πTemp002 = false
					Label1:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp002 {
							πF.PopCheckpoint()
							goto Label3
						}
						if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
							isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
							if exc != nil {
								πE = exc
							} else if isStop {
								πE = nil
								πF.RestoreExc(nil, nil)
							}
							πTemp003 = !isStop
						} else {
							πTemp003 = true
							µargs = πTemp004
						}
						if πE != nil || !πTemp003 {
							continue
						}
						πF.PushCheckpoint(1)            
						// line 219: yield function(*args)
						πF.SetLineno(219)
						if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µfunction, "function"); πE != nil {
							continue
						}
						if πTemp004, πE = πg.Invoke(πF, µfunction, nil, µargs, nil, nil); πE != nil {
							continue
						}
						πF.PushCheckpoint(4)
						return πTemp004, nil
					Label4:
						πTemp005 = πSent
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
			if πE = πF.Globals().SetItem(πF, ßstarmap.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 222: def takewhile(predicate, iterable):
			πF.SetLineno(222)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "predicate", Def: nil}
			πTemp006[1] = πg.Param{Name: "iterable", Def: nil}
			πTemp020 = πg.NewFunction(πg.NewCode("takewhile", "build/src/__python__/itertools.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpredicate *πg.Object = πArgs[0]; _ = µpredicate
				var µiterable *πg.Object = πArgs[1]; _ = µiterable
				var µx *πg.Object = πg.UnboundLocal; _ = µx
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 1: goto Label1
						case 2: goto Label2
						case 7: goto Label7
						default: panic("unexpected function state")
						}
						if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.Iter(πF, µiterable); πE != nil {
							continue
						}
						πF.PushCheckpoint(2)
						πTemp002 = false
					Label1:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp002 {
							πF.PopCheckpoint()
							goto Label3
						}
						if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
							isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
							if exc != nil {
								πE = exc
							} else if isStop {
								πE = nil
								πF.RestoreExc(nil, nil)
							}
							πTemp003 = !isStop
						} else {
							πTemp003 = true
							µx = πTemp004
						}
						if πE != nil || !πTemp003 {
							continue
						}
						πF.PushCheckpoint(1)            
						πTemp005 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
							continue
						}
						πTemp005[0] = µx
						if πE = πg.CheckLocal(πF, µpredicate, "predicate"); πE != nil {
							continue
						}
						if πTemp004, πE = µpredicate.Call(πF, πTemp005, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp005)
						if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
							continue
						}
						if πTemp003 {
							goto Label4
						}
						goto Label5
						// line 224: if predicate(x):
						πF.SetLineno(224)
					Label4:
						// line 225: yield x
						πF.SetLineno(225)
						if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
							continue
						}
						πF.PushCheckpoint(7)
						return µx, nil
					Label7:
						πTemp004 = πSent
						goto Label6
					Label5:
						// line 227: break
						πF.SetLineno(227)
						πTemp002 = true
						continue
						goto Label6
					Label6:
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
			if πE = πF.Globals().SetItem(πF, ßtakewhile.ToObject(), πTemp020); πE != nil {
				continue
			}
			// line 230: def tee(iterable, n=2):
			πF.SetLineno(230)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "iterable", Def: nil}
			πTemp006[1] = πg.Param{Name: "n", Def: πg.NewInt(2).ToObject()}
			πTemp021 = πg.NewFunction(πg.NewCode("tee", "build/src/__python__/itertools.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µiterable *πg.Object = πArgs[0]; _ = µiterable
				var µn *πg.Object = πArgs[1]; _ = µn
				var µit *πg.Object = πg.UnboundLocal; _ = µit
				var µdeques *πg.Object = πg.UnboundLocal; _ = µdeques
				var µgen *πg.Object = πg.UnboundLocal; _ = µgen
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 231: it = iter(iterable)
					πF.SetLineno(231)
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
					// line 232: deques = [_collections.deque() for i in range(n)]
					πF.SetLineno(232)
					πTemp004 = make([]πg.Param, 0)
					πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/itertools.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µi *πg.Object = πg.UnboundLocal; _ = µi
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
								πTemp002 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
									continue
								}
								πTemp002[0] = µn
								if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
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
									µi = πTemp003
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 232: deques = [_collections.deque() for i in range(n)]
								πF.SetLineno(232)
								if πTemp003, πE = πg.ResolveGlobal(πF, ß_collections); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßdeque, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
									continue
								}
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
					if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
						continue
					}
					µdeques = πTemp002
					// line 233: def gen(mydeque):
					πF.SetLineno(233)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "mydeque", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("gen", "build/src/__python__/itertools.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µmydeque *πg.Object = πArgs[0]; _ = µmydeque
						var µnewval *πg.Object = πg.UnboundLocal; _ = µnewval
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var πTemp001 bool
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 9: goto Label9
								case 6: goto Label6
								case 7: goto Label7
								default: panic("unexpected function state")
								}
								// line 234: while True:
								πF.SetLineno(234)
								πF.PushCheckpoint(2)
								πTemp001 = false
							Label1:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp001 {
									πF.PopCheckpoint()
									goto Label3
								}
								if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πE != nil || !πTemp002 {
									continue
								}
								πF.PushCheckpoint(1)            
								if πE = πg.CheckLocal(πF, µmydeque, "mydeque"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.IsTrue(πF, µmydeque); πE != nil {
									continue
								}
								πTemp003 = πg.GetBool(!πTemp002).ToObject()
								if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp002 {
									goto Label4
								}
								goto Label5
								// line 235: if not mydeque:
								πF.SetLineno(235)
							Label4:
								// line 236: newval = next(it)
								πF.SetLineno(236)
								πTemp004 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
									continue
								}
								πTemp004[0] = µit
								if πTemp003, πE = πg.ResolveGlobal(πF, ßnext); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								µnewval = πTemp005
								if πE = πg.CheckLocal(πF, µdeques, "deques"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Iter(πF, µdeques); πE != nil {
									continue
								}
								πF.PushCheckpoint(7)
								πTemp002 = false
							Label6:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp002 {
									πF.PopCheckpoint()
									goto Label8
								}
								if πTemp005, πE = πg.Next(πF, πTemp003); πE != nil {
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
									µd = πTemp005
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(6)            
								// line 238: d.append(newval)
								πF.SetLineno(238)
								πTemp004 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µnewval, "newval"); πE != nil {
									continue
								}
								πTemp004[0] = µnewval
								if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µd, ßappend, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								continue
							Label7:
								if πE != nil || πR != nil {
									continue
								}
							Label8:
								goto Label5
							Label5:
								// line 239: yield mydeque.popleft()
								πF.SetLineno(239)
								if πE = πg.CheckLocal(πF, µmydeque, "mydeque"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µmydeque, ßpopleft, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
									continue
								}
								πF.PushCheckpoint(9)
								return πTemp005, nil
							Label9:
								πTemp003 = πSent
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
					µgen = πTemp002
					// line 240: return tuple(gen(d) for d in deques)
					πF.SetLineno(240)
					πTemp001 = πF.MakeArgs(1)
					πTemp004 = make([]πg.Param, 0)
					πTemp005 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/itertools.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
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
								case 4: goto Label4
								default: panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µdeques, "deques"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µdeques); πE != nil {
									continue
								}
								πF.PushCheckpoint(2)
								πTemp002 = false
							Label1:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp002 {
									πF.PopCheckpoint()
									goto Label3
								}
								if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp003 = !isStop
								} else {
									πTemp003 = true
									µd = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 240: return tuple(gen(d) for d in deques)
								πF.SetLineno(240)
								πTemp005 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
									continue
								}
								πTemp005[0] = µd
								if πE = πg.CheckLocal(πF, µgen, "gen"); πE != nil {
									continue
								}
								if πTemp004, πE = µgen.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								πF.PushCheckpoint(4)
								return πTemp004, nil
							Label4:
								πTemp006 = πSent
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
					if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp006
					if πTemp006, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
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
			if πE = πF.Globals().SetItem(πF, ßtee.ToObject(), πTemp021); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("itertools", Code)
}
