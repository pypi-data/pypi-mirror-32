package itertools_test
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/itertools_test.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßA := πg.InternStr("A")
		ßAB := πg.InternStr("AB")
		ßABC := πg.InternStr("ABC")
		ßB := πg.InternStr("B")
		ßC := πg.InternStr("C")
		ßFalse := πg.InternStr("False")
		ßNone := πg.InternStr("None")
		ßRunTests := πg.InternStr("RunTests")
		ßTestChain := πg.InternStr("TestChain")
		ßTestCombinations := πg.InternStr("TestCombinations")
		ßTestCombinationsWithReplacement := πg.InternStr("TestCombinationsWithReplacement")
		ßTestCycle := πg.InternStr("TestCycle")
		ßTestDropwhile := πg.InternStr("TestDropwhile")
		ßTestFromIterable := πg.InternStr("TestFromIterable")
		ßTestGroupBy := πg.InternStr("TestGroupBy")
		ßTestIFilter := πg.InternStr("TestIFilter")
		ßTestIFilterFalse := πg.InternStr("TestIFilterFalse")
		ßTestISlice := πg.InternStr("TestISlice")
		ßTestIZipLongest := πg.InternStr("TestIZipLongest")
		ßTestPermutations := πg.InternStr("TestPermutations")
		ßTestProduct := πg.InternStr("TestProduct")
		ßTestTakewhile := πg.InternStr("TestTakewhile")
		ßTrue := πg.InternStr("True")
		ß__main__ := πg.InternStr("__main__")
		ß__name__ := πg.InternStr("__name__")
		ßa := πg.InternStr("a")
		ßaa := πg.InternStr("aa")
		ßab := πg.InternStr("ab")
		ßabc := πg.InternStr("abc")
		ßabcde := πg.InternStr("abcde")
		ßappend := πg.InternStr("append")
		ßb := πg.InternStr("b")
		ßbcd := πg.InternStr("bcd")
		ßc := πg.InternStr("c")
		ßchain := πg.InternStr("chain")
		ßcombinations := πg.InternStr("combinations")
		ßcombinations_with_replacement := πg.InternStr("combinations_with_replacement")
		ßcycle := πg.InternStr("cycle")
		ßdropwhile := πg.InternStr("dropwhile")
		ßfrom_iterable := πg.InternStr("from_iterable")
		ßgroupby := πg.InternStr("groupby")
		ßifilter := πg.InternStr("ifilter")
		ßifilterfalse := πg.InternStr("ifilterfalse")
		ßislice := πg.InternStr("islice")
		ßitertools := πg.InternStr("itertools")
		ßizip_longest := πg.InternStr("izip_longest")
		ßlen := πg.InternStr("len")
		ßlist := πg.InternStr("list")
		ßpermutations := πg.InternStr("permutations")
		ßproduct := πg.InternStr("product")
		ßrange := πg.InternStr("range")
		ßtakewhile := πg.InternStr("takewhile")
		ßtuple := πg.InternStr("tuple")
		ßweetest := πg.InternStr("weetest")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 []πg.Param
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
		var πTemp016 *πg.Object
		_ = πTemp016
		var πTemp017 *πg.Object
		_ = πTemp017
		var πTemp018 *πg.Object
		_ = πTemp018
		var πTemp019 bool
		_ = πTemp019
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 15: import itertools
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "itertools"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßitertools.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 17: import weetest
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "weetest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßweetest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 19: def TestCycle():
			πF.SetLineno(19)
			πTemp003 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("TestCycle", "build/src/__python__/itertools_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µwant *πg.Object = πg.UnboundLocal; _ = µwant
				var µgot *πg.Object = πg.UnboundLocal; _ = µgot
				var µx *πg.Object = πg.UnboundLocal; _ = µx
				var µarg *πg.Object = πg.UnboundLocal; _ = µarg
				var µlimit *πg.Object = πg.UnboundLocal; _ = µlimit
				var µcounter *πg.Object = πg.UnboundLocal; _ = µcounter
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
				var πTemp006 bool
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
					case 4: goto Label4
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 20: want = []
					πF.SetLineno(20)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µwant = πTemp002
					// line 21: got = []
					πF.SetLineno(21)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µgot = πTemp002
					πTemp001 = πF.MakeArgs(1)
					πTemp003 = πg.NewTuple0().ToObject()
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßcycle, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
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
						µx = πTemp003
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 23: got.append(x)
					πF.SetLineno(23)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µgot, ßappend, nil); πE != nil {
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
					// line 24: assert got == want, 'empty cycle yields no elements'
					πF.SetLineno(24)
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µgot, µwant); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp002, πg.NewStr("empty cycle yields no elements").ToObject()); πE != nil {
						continue
					}
					// line 26: arg = (0, 1, 2)
					πF.SetLineno(26)
					πTemp002 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject()).ToObject()
					µarg = πTemp002
					// line 27: want = (0, 1, 2) * 10
					πF.SetLineno(27)
					πTemp003 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject()).ToObject()
					if πTemp002, πE = πg.Mul(πF, πTemp003, πg.NewInt(10).ToObject()); πE != nil {
						continue
					}
					µwant = πTemp002
					// line 28: got = []
					πF.SetLineno(28)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µgot = πTemp002
					// line 29: limit = 10 * len(arg)
					πF.SetLineno(29)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
						continue
					}
					πTemp001[0] = µarg
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Mul(πF, πg.NewInt(10).ToObject(), πTemp004); πE != nil {
						continue
					}
					µlimit = πTemp002
					// line 30: counter = 0
					πF.SetLineno(30)
					µcounter = πg.NewInt(0).ToObject()
					πTemp001 = πF.MakeArgs(1)
					πTemp003 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject()).ToObject()
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßcycle, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
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
						µx = πTemp003
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(4)            
					// line 32: got.append(x)
					πF.SetLineno(32)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µgot, ßappend, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 33: counter += 1
					πF.SetLineno(33)
					if πE = πg.CheckLocal(πF, µcounter, "counter"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IAdd(πF, µcounter, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µcounter = πTemp003
					if πE = πg.CheckLocal(πF, µcounter, "counter"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µcounter, µlimit); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label7
					}
					goto Label8
					// line 34: if counter == limit:
					πF.SetLineno(34)
				Label7:
					// line 35: break
					πF.SetLineno(35)
					πTemp005 = true
					continue
					goto Label8
				Label8:
					continue
				Label5:
					if πE != nil || πR != nil {
						continue
					}
				Label6:
					// line 36: assert tuple(got) == want, 'tuple(cycle%s) == %s, want %s' % (arg, tuple(got), want)
					πF.SetLineno(36)
					if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					πTemp001[0] = µgot
					if πTemp004, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple3(µarg, πTemp007, µwant).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("tuple(cycle%s) == %s, want %s").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					πTemp001[0] = µgot
					if πTemp004, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp007, µwant); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp003, πTemp002); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTestCycle.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 39: def TestDropwhile():
			πF.SetLineno(39)
			πTemp003 = make([]πg.Param, 0)
			πTemp004 = πg.NewFunction(πg.NewCode("TestDropwhile", "build/src/__python__/itertools_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µr *πg.Object = πg.UnboundLocal; _ = µr
				var µcases *πg.Object = πg.UnboundLocal; _ = µcases
				var µargs *πg.Object = πg.UnboundLocal; _ = µargs
				var µwant *πg.Object = πg.UnboundLocal; _ = µwant
				var µgot *πg.Object = πg.UnboundLocal; _ = µgot
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []πg.Param
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 40: r = range(10)
					πF.SetLineno(40)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewInt(10).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µr = πTemp003
					// line 41: cases = [
					πF.SetLineno(41)
					πTemp001 = make([]*πg.Object, 3)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "x", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/itertools_test.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µx *πg.Object = πArgs[0]; _ = µx
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 42: ((lambda x: x < 5, r), (5, 6, 7, 8, 9)),
							πF.SetLineno(42)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µx, πg.NewInt(5).ToObject()); πE != nil {
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
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp004, µr).ToObject()
					πTemp004 = πg.NewTuple5(πg.NewInt(5).ToObject(), πg.NewInt(6).ToObject(), πg.NewInt(7).ToObject(), πg.NewInt(8).ToObject(), πg.NewInt(9).ToObject()).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[0] = πTemp002
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "x", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/itertools_test.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µx *πg.Object = πArgs[0]; _ = µx
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 43: ((lambda x: True, r), ()),
							πF.SetLineno(43)
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
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp004, µr).ToObject()
					πTemp004 = πg.NewTuple0().ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[1] = πTemp002
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "x", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/itertools_test.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µx *πg.Object = πArgs[0]; _ = µx
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 44: ((lambda x: False, r), tuple(r)),
							πF.SetLineno(44)
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
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp004, µr).ToObject()
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp006[0] = µr
					if πTemp004, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp002 = πg.NewTuple2(πTemp003, πTemp007).ToObject()
					πTemp001[2] = πTemp002
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µcases = πTemp002
					if πE = πg.CheckLocal(πF, µcases, "cases"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µcases); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp008 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
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
						πTemp009 = !isStop
					} else {
						πTemp009 = true
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
							continue
						}
						µargs = πTemp004
						µwant = πTemp007
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 47: got = tuple(itertools.dropwhile(*args))
					πF.SetLineno(47)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßdropwhile, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp004, nil, µargs, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µgot = πTemp004
					// line 48: assert got == want, 'tuple(dropwhile%s) == %s, want %s' % (args, got, want)
					πF.SetLineno(48)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple3(µargs, µgot, µwant).ToObject()
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("tuple(dropwhile%s) == %s, want %s").ToObject(), πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, µgot, µwant); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp004, πTemp003); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßTestDropwhile.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 51: def TestChain():
			πF.SetLineno(51)
			πTemp003 = make([]πg.Param, 0)
			πTemp005 = πg.NewFunction(πg.NewCode("TestChain", "build/src/__python__/itertools_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µr *πg.Object = πg.UnboundLocal; _ = µr
				var µcases *πg.Object = πg.UnboundLocal; _ = µcases
				var µargs *πg.Object = πg.UnboundLocal; _ = µargs
				var µwant *πg.Object = πg.UnboundLocal; _ = µwant
				var µgot *πg.Object = πg.UnboundLocal; _ = µgot
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πTemp010 bool
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
					// line 52: r = range(10)
					πF.SetLineno(52)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewInt(10).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µr = πTemp003
					// line 53: cases = [
					πF.SetLineno(53)
					πTemp001 = make([]*πg.Object, 3)
					πTemp004 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp004[0] = µr
					πTemp003 = πg.NewList(πTemp004...).ToObject()
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp004[0] = µr
					if πTemp005, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp002 = πg.NewTuple2(πTemp003, πTemp006).ToObject()
					πTemp001[0] = πTemp002
					πTemp004 = make([]*πg.Object, 2)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp004[0] = µr
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp004[1] = µr
					πTemp003 = πg.NewList(πTemp004...).ToObject()
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp004[0] = µr
					if πTemp006, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp004[0] = µr
					if πTemp006, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp005, πE = πg.Add(πF, πTemp007, πTemp008); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
					πTemp001[1] = πTemp002
					πTemp004 = make([]*πg.Object, 0)
					πTemp003 = πg.NewList(πTemp004...).ToObject()
					πTemp005 = πg.NewTuple0().ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
					πTemp001[2] = πTemp002
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µcases = πTemp002
					if πE = πg.CheckLocal(πF, µcases, "cases"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µcases); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp009 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp009 {
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
						πTemp010 = !isStop
					} else {
						πTemp010 = true
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp003); πE != nil {
							continue
						}
						µargs = πTemp005
						µwant = πTemp006
					}
					if πE != nil || !πTemp010 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 59: got = tuple(itertools.chain(*args))
					πF.SetLineno(59)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßchain, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp005, nil, µargs, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µgot = πTemp005
					// line 60: assert got == want, 'tuple(chain%s) == %s, want %s' % (args, got, want)
					πF.SetLineno(60)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple3(µargs, µgot, µwant).ToObject()
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("tuple(chain%s) == %s, want %s").ToObject(), πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, µgot, µwant); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp005, πTemp003); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßTestChain.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 63: def TestFromIterable():
			πF.SetLineno(63)
			πTemp003 = make([]πg.Param, 0)
			πTemp006 = πg.NewFunction(πg.NewCode("TestFromIterable", "build/src/__python__/itertools_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µr *πg.Object = πg.UnboundLocal; _ = µr
				var µcases *πg.Object = πg.UnboundLocal; _ = µcases
				var µargs *πg.Object = πg.UnboundLocal; _ = µargs
				var µwant *πg.Object = πg.UnboundLocal; _ = µwant
				var µgot *πg.Object = πg.UnboundLocal; _ = µgot
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πTemp010 bool
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
					// line 64: r = range(10)
					πF.SetLineno(64)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewInt(10).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µr = πTemp003
					// line 65: cases = [
					πF.SetLineno(65)
					πTemp001 = make([]*πg.Object, 3)
					πTemp004 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp004[0] = µr
					πTemp003 = πg.NewList(πTemp004...).ToObject()
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp004[0] = µr
					if πTemp005, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp002 = πg.NewTuple2(πTemp003, πTemp006).ToObject()
					πTemp001[0] = πTemp002
					πTemp004 = make([]*πg.Object, 2)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp004[0] = µr
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp004[1] = µr
					πTemp003 = πg.NewList(πTemp004...).ToObject()
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp004[0] = µr
					if πTemp006, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp004[0] = µr
					if πTemp006, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp005, πE = πg.Add(πF, πTemp007, πTemp008); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
					πTemp001[1] = πTemp002
					πTemp004 = make([]*πg.Object, 0)
					πTemp003 = πg.NewList(πTemp004...).ToObject()
					πTemp005 = πg.NewTuple0().ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
					πTemp001[2] = πTemp002
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µcases = πTemp002
					if πE = πg.CheckLocal(πF, µcases, "cases"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µcases); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp009 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp009 {
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
						πTemp010 = !isStop
					} else {
						πTemp010 = true
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp003); πE != nil {
							continue
						}
						µargs = πTemp005
						µwant = πTemp006
					}
					if πE != nil || !πTemp010 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 71: got = tuple(itertools.chain.from_iterable(args))
					πF.SetLineno(71)
					πTemp001 = πF.MakeArgs(1)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp004[0] = µargs
					if πTemp003, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßchain, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp005, ßfrom_iterable, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp001[0] = πTemp005
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µgot = πTemp005
					// line 72: assert got == want, 'tuple(from_iterable%s) == %s, want %s' % (args, got, want)
					πF.SetLineno(72)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple3(µargs, µgot, µwant).ToObject()
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("tuple(from_iterable%s) == %s, want %s").ToObject(), πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, µgot, µwant); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp005, πTemp003); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßTestFromIterable.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 75: def TestIFilter():
			πF.SetLineno(75)
			πTemp003 = make([]πg.Param, 0)
			πTemp007 = πg.NewFunction(πg.NewCode("TestIFilter", "build/src/__python__/itertools_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µr *πg.Object = πg.UnboundLocal; _ = µr
				var µcases *πg.Object = πg.UnboundLocal; _ = µcases
				var µargs *πg.Object = πg.UnboundLocal; _ = µargs
				var µwant *πg.Object = πg.UnboundLocal; _ = µwant
				var µgot *πg.Object = πg.UnboundLocal; _ = µgot
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []πg.Param
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 76: r = range(10)
					πF.SetLineno(76)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewInt(10).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µr = πTemp003
					// line 77: cases = [
					πF.SetLineno(77)
					πTemp001 = make([]*πg.Object, 4)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "x", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/itertools_test.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µx *πg.Object = πArgs[0]; _ = µx
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 78: ((lambda x: x < 5, r), (0, 1, 2, 3, 4)),
							πF.SetLineno(78)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µx, πg.NewInt(5).ToObject()); πE != nil {
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
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp004, µr).ToObject()
					πTemp004 = πg.NewTuple5(πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject(), πg.NewInt(3).ToObject(), πg.NewInt(4).ToObject()).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[0] = πTemp002
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "x", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/itertools_test.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µx *πg.Object = πArgs[0]; _ = µx
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 79: ((lambda x: False, r), ()),
							πF.SetLineno(79)
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
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp004, µr).ToObject()
					πTemp004 = πg.NewTuple0().ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[1] = πTemp002
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "x", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/itertools_test.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µx *πg.Object = πArgs[0]; _ = µx
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 80: ((lambda x: True, r), tuple(r)),
							πF.SetLineno(80)
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
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp004, µr).ToObject()
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp006[0] = µr
					if πTemp004, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp002 = πg.NewTuple2(πTemp003, πTemp007).ToObject()
					πTemp001[2] = πTemp002
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp004, µr).ToObject()
					πTemp006 = make([]*πg.Object, 9)
					πTemp006[0] = πg.NewInt(1).ToObject()
					πTemp006[1] = πg.NewInt(2).ToObject()
					πTemp006[2] = πg.NewInt(3).ToObject()
					πTemp006[3] = πg.NewInt(4).ToObject()
					πTemp006[4] = πg.NewInt(5).ToObject()
					πTemp006[5] = πg.NewInt(6).ToObject()
					πTemp006[6] = πg.NewInt(7).ToObject()
					πTemp006[7] = πg.NewInt(8).ToObject()
					πTemp006[8] = πg.NewInt(9).ToObject()
					πTemp004 = πg.NewTuple(πTemp006...).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[3] = πTemp002
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µcases = πTemp002
					if πE = πg.CheckLocal(πF, µcases, "cases"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µcases); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp008 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
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
						πTemp009 = !isStop
					} else {
						πTemp009 = true
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
							continue
						}
						µargs = πTemp004
						µwant = πTemp007
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 84: got = tuple(itertools.ifilter(*args))
					πF.SetLineno(84)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßifilter, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp004, nil, µargs, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µgot = πTemp004
					// line 85: assert got == want, 'tuple(ifilter%s) == %s, want %s' % (args, got, want)
					πF.SetLineno(85)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple3(µargs, µgot, µwant).ToObject()
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("tuple(ifilter%s) == %s, want %s").ToObject(), πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, µgot, µwant); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp004, πTemp003); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßTestIFilter.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 88: def TestIFilterFalse():
			πF.SetLineno(88)
			πTemp003 = make([]πg.Param, 0)
			πTemp008 = πg.NewFunction(πg.NewCode("TestIFilterFalse", "build/src/__python__/itertools_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µr *πg.Object = πg.UnboundLocal; _ = µr
				var µcases *πg.Object = πg.UnboundLocal; _ = µcases
				var µargs *πg.Object = πg.UnboundLocal; _ = µargs
				var µwant *πg.Object = πg.UnboundLocal; _ = µwant
				var µgot *πg.Object = πg.UnboundLocal; _ = µgot
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []πg.Param
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 89: r = range(10)
					πF.SetLineno(89)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewInt(10).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µr = πTemp003
					// line 90: cases = [
					πF.SetLineno(90)
					πTemp001 = make([]*πg.Object, 4)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "x", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/itertools_test.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µx *πg.Object = πArgs[0]; _ = µx
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 91: ((lambda x: x < 5, r), (5, 6, 7, 8, 9)),
							πF.SetLineno(91)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µx, πg.NewInt(5).ToObject()); πE != nil {
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
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp004, µr).ToObject()
					πTemp004 = πg.NewTuple5(πg.NewInt(5).ToObject(), πg.NewInt(6).ToObject(), πg.NewInt(7).ToObject(), πg.NewInt(8).ToObject(), πg.NewInt(9).ToObject()).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[0] = πTemp002
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "x", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/itertools_test.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µx *πg.Object = πArgs[0]; _ = µx
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 92: ((lambda x: False, r), tuple(r)),
							πF.SetLineno(92)
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
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp004, µr).ToObject()
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp006[0] = µr
					if πTemp004, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp002 = πg.NewTuple2(πTemp003, πTemp007).ToObject()
					πTemp001[1] = πTemp002
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "x", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/itertools_test.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µx *πg.Object = πArgs[0]; _ = µx
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 93: ((lambda x: True, r), ()),
							πF.SetLineno(93)
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
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp004, µr).ToObject()
					πTemp004 = πg.NewTuple0().ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[2] = πTemp002
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp004, µr).ToObject()
					πTemp004 = πg.NewTuple1(πg.NewInt(0).ToObject()).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[3] = πTemp002
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µcases = πTemp002
					if πE = πg.CheckLocal(πF, µcases, "cases"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µcases); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp008 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
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
						πTemp009 = !isStop
					} else {
						πTemp009 = true
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
							continue
						}
						µargs = πTemp004
						µwant = πTemp007
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 97: got = tuple(itertools.ifilterfalse(*args))
					πF.SetLineno(97)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßifilterfalse, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp004, nil, µargs, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µgot = πTemp004
					// line 98: assert got == want, 'tuple(ifilterfalse%s) == %s, want %s' % (args, got, want)
					πF.SetLineno(98)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple3(µargs, µgot, µwant).ToObject()
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("tuple(ifilterfalse%s) == %s, want %s").ToObject(), πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, µgot, µwant); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp004, πTemp003); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßTestIFilterFalse.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 101: def TestISlice():
			πF.SetLineno(101)
			πTemp003 = make([]πg.Param, 0)
			πTemp009 = πg.NewFunction(πg.NewCode("TestISlice", "build/src/__python__/itertools_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µr *πg.Object = πg.UnboundLocal; _ = µr
				var µcases *πg.Object = πg.UnboundLocal; _ = µcases
				var µargs *πg.Object = πg.UnboundLocal; _ = µargs
				var µwant *πg.Object = πg.UnboundLocal; _ = µwant
				var µgot *πg.Object = πg.UnboundLocal; _ = µgot
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
				var πTemp006 bool
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
					// line 102: r = range(10)
					πF.SetLineno(102)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewInt(10).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µr = πTemp003
					// line 103: cases = [
					πF.SetLineno(103)
					πTemp001 = make([]*πg.Object, 3)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µr, πg.NewInt(5).ToObject()).ToObject()
					πTemp004 = πg.NewTuple5(πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject(), πg.NewInt(3).ToObject(), πg.NewInt(4).ToObject()).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple3(µr, πg.NewInt(25).ToObject(), πg.NewInt(30).ToObject()).ToObject()
					πTemp004 = πg.NewTuple0().ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple4(µr, πg.NewInt(1).ToObject(), πTemp004, πg.NewInt(3).ToObject()).ToObject()
					πTemp004 = πg.NewTuple3(πg.NewInt(1).ToObject(), πg.NewInt(4).ToObject(), πg.NewInt(7).ToObject()).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[2] = πTemp002
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µcases = πTemp002
					if πE = πg.CheckLocal(πF, µcases, "cases"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µcases); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
							continue
						}
						µargs = πTemp004
						µwant = πTemp007
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 109: got = tuple(itertools.islice(*args))
					πF.SetLineno(109)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßislice, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp004, nil, µargs, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µgot = πTemp004
					// line 110: assert got == want, 'tuple(islice%s) == %s, want %s' % (args, got, want)
					πF.SetLineno(110)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple3(µargs, µgot, µwant).ToObject()
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("tuple(islice%s) == %s, want %s").ToObject(), πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, µgot, µwant); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp004, πTemp003); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßTestISlice.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 113: def TestIZipLongest():
			πF.SetLineno(113)
			πTemp003 = make([]πg.Param, 0)
			πTemp010 = πg.NewFunction(πg.NewCode("TestIZipLongest", "build/src/__python__/itertools_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcases *πg.Object = πg.UnboundLocal; _ = µcases
				var µargs *πg.Object = πg.UnboundLocal; _ = µargs
				var µwant *πg.Object = πg.UnboundLocal; _ = µwant
				var µgot *πg.Object = πg.UnboundLocal; _ = µgot
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
				var πTemp013 bool
				_ = πTemp013
				var πTemp014 bool
				_ = πTemp014
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 114: cases = [
					πF.SetLineno(114)
					πTemp001 = make([]*πg.Object, 3)
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewInt(6).ToObject()
					if πTemp005, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp003 = πg.NewTuple2(ßabc.ToObject(), πTemp006).ToObject()
					πTemp006 = πg.NewTuple2(ßa.ToObject(), πg.NewInt(0).ToObject()).ToObject()
					πTemp007 = πg.NewTuple2(ßb.ToObject(), πg.NewInt(1).ToObject()).ToObject()
					πTemp008 = πg.NewTuple2(ßc.ToObject(), πg.NewInt(2).ToObject()).ToObject()
					if πTemp010, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp009 = πg.NewTuple2(πTemp010, πg.NewInt(3).ToObject()).ToObject()
					if πTemp011, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp010 = πg.NewTuple2(πTemp011, πg.NewInt(4).ToObject()).ToObject()
					if πTemp012, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp011 = πg.NewTuple2(πTemp012, πg.NewInt(5).ToObject()).ToObject()
					πTemp005 = πg.NewTuple6(πTemp006, πTemp007, πTemp008, πTemp009, πTemp010, πTemp011).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
					πTemp001[0] = πTemp002
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewInt(6).ToObject()
					if πTemp005, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp003 = πg.NewTuple2(πTemp006, ßabc.ToObject()).ToObject()
					πTemp006 = πg.NewTuple2(πg.NewInt(0).ToObject(), ßa.ToObject()).ToObject()
					πTemp007 = πg.NewTuple2(πg.NewInt(1).ToObject(), ßb.ToObject()).ToObject()
					πTemp008 = πg.NewTuple2(πg.NewInt(2).ToObject(), ßc.ToObject()).ToObject()
					if πTemp010, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp009 = πg.NewTuple2(πg.NewInt(3).ToObject(), πTemp010).ToObject()
					if πTemp011, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp010 = πg.NewTuple2(πg.NewInt(4).ToObject(), πTemp011).ToObject()
					if πTemp012, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp011 = πg.NewTuple2(πg.NewInt(5).ToObject(), πTemp012).ToObject()
					πTemp005 = πg.NewTuple6(πTemp006, πTemp007, πTemp008, πTemp009, πTemp010, πTemp011).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
					πTemp001[1] = πTemp002
					πTemp004 = make([]*πg.Object, 3)
					πTemp004[0] = πg.NewInt(1).ToObject()
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp004[1] = πTemp005
					πTemp004[2] = πg.NewInt(3).ToObject()
					πTemp005 = πg.NewList(πTemp004...).ToObject()
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewInt(1).ToObject()
					if πTemp006, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp003 = πg.NewTuple3(πTemp005, ßab.ToObject(), πTemp007).ToObject()
					πTemp006 = πg.NewTuple3(πg.NewInt(1).ToObject(), ßa.ToObject(), πg.NewInt(0).ToObject()).ToObject()
					if πTemp008, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp007 = πg.NewTuple3(πTemp008, ßb.ToObject(), πTemp009).ToObject()
					if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp008 = πg.NewTuple3(πg.NewInt(3).ToObject(), πTemp009, πTemp010).ToObject()
					πTemp005 = πg.NewTuple3(πTemp006, πTemp007, πTemp008).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
					πTemp001[2] = πTemp002
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µcases = πTemp002
					if πE = πg.CheckLocal(πF, µcases, "cases"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µcases); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp013 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp013 {
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
						πTemp014 = !isStop
					} else {
						πTemp014 = true
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp003); πE != nil {
							continue
						}
						µargs = πTemp005
						µwant = πTemp006
					}
					if πE != nil || !πTemp014 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 120: got = tuple(itertools.izip_longest(*args))
					πF.SetLineno(120)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßizip_longest, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp005, nil, µargs, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µgot = πTemp005
					// line 121: assert got == want, 'tuple(izip_longest%s) == %s, want %s' % (args, got, want)
					πF.SetLineno(121)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple3(µargs, µgot, µwant).ToObject()
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("tuple(izip_longest%s) == %s, want %s").ToObject(), πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, µgot, µwant); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp005, πTemp003); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßTestIZipLongest.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 124: def TestProduct():
			πF.SetLineno(124)
			πTemp003 = make([]πg.Param, 0)
			πTemp011 = πg.NewFunction(πg.NewCode("TestProduct", "build/src/__python__/itertools_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcases *πg.Object = πg.UnboundLocal; _ = µcases
				var µargs *πg.Object = πg.UnboundLocal; _ = µargs
				var µwant *πg.Object = πg.UnboundLocal; _ = µwant
				var µgot *πg.Object = πg.UnboundLocal; _ = µgot
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 bool
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
					default: panic("unexpected function state")
					}
					// line 125: cases = [
					πF.SetLineno(125)
					πTemp001 = make([]*πg.Object, 3)
					πTemp004 = make([]*πg.Object, 2)
					πTemp004[0] = πg.NewInt(1).ToObject()
					πTemp004[1] = πg.NewInt(2).ToObject()
					πTemp005 = πg.NewList(πTemp004...).ToObject()
					πTemp004 = make([]*πg.Object, 2)
					πTemp004[0] = ßa.ToObject()
					πTemp004[1] = ßb.ToObject()
					πTemp006 = πg.NewList(πTemp004...).ToObject()
					πTemp003 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
					πTemp006 = πg.NewTuple2(πg.NewInt(1).ToObject(), ßa.ToObject()).ToObject()
					πTemp007 = πg.NewTuple2(πg.NewInt(1).ToObject(), ßb.ToObject()).ToObject()
					πTemp008 = πg.NewTuple2(πg.NewInt(2).ToObject(), ßa.ToObject()).ToObject()
					πTemp009 = πg.NewTuple2(πg.NewInt(2).ToObject(), ßb.ToObject()).ToObject()
					πTemp005 = πg.NewTuple4(πTemp006, πTemp007, πTemp008, πTemp009).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
					πTemp001[0] = πTemp002
					πTemp004 = make([]*πg.Object, 1)
					πTemp004[0] = πg.NewInt(1).ToObject()
					πTemp005 = πg.NewList(πTemp004...).ToObject()
					πTemp004 = make([]*πg.Object, 2)
					πTemp004[0] = ßa.ToObject()
					πTemp004[1] = ßb.ToObject()
					πTemp006 = πg.NewList(πTemp004...).ToObject()
					πTemp003 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
					πTemp006 = πg.NewTuple2(πg.NewInt(1).ToObject(), ßa.ToObject()).ToObject()
					πTemp007 = πg.NewTuple2(πg.NewInt(1).ToObject(), ßb.ToObject()).ToObject()
					πTemp005 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
					πTemp001[1] = πTemp002
					πTemp004 = make([]*πg.Object, 0)
					πTemp005 = πg.NewList(πTemp004...).ToObject()
					πTemp003 = πg.NewTuple1(πTemp005).ToObject()
					πTemp005 = πg.NewTuple0().ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
					πTemp001[2] = πTemp002
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µcases = πTemp002
					if πE = πg.CheckLocal(πF, µcases, "cases"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µcases); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp010 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp010 {
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
						πTemp011 = !isStop
					} else {
						πTemp011 = true
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp003); πE != nil {
							continue
						}
						µargs = πTemp005
						µwant = πTemp006
					}
					if πE != nil || !πTemp011 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 131: got = tuple(itertools.product(*args))
					πF.SetLineno(131)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßproduct, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp005, nil, µargs, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µgot = πTemp005
					// line 132: assert got == want, 'tuple(product%s) == %s, want %s' % (args, got, want)
					πF.SetLineno(132)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple3(µargs, µgot, µwant).ToObject()
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("tuple(product%s) == %s, want %s").ToObject(), πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, µgot, µwant); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp005, πTemp003); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßTestProduct.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 135: def TestPermutations():
			πF.SetLineno(135)
			πTemp003 = make([]πg.Param, 0)
			πTemp012 = πg.NewFunction(πg.NewCode("TestPermutations", "build/src/__python__/itertools_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcases *πg.Object = πg.UnboundLocal; _ = µcases
				var µargs *πg.Object = πg.UnboundLocal; _ = µargs
				var µwant *πg.Object = πg.UnboundLocal; _ = µwant
				var µgot *πg.Object = πg.UnboundLocal; _ = µgot
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
				var πTemp011 []*πg.Object
				_ = πTemp011
				var πTemp012 bool
				_ = πTemp012
				var πTemp013 bool
				_ = πTemp013
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 136: cases = [
					πF.SetLineno(136)
					πTemp001 = make([]*πg.Object, 6)
					πTemp003 = πg.NewTuple1(ßAB.ToObject()).ToObject()
					πTemp005 = πg.NewTuple2(ßA.ToObject(), ßB.ToObject()).ToObject()
					πTemp006 = πg.NewTuple2(ßB.ToObject(), ßA.ToObject()).ToObject()
					πTemp004 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[0] = πTemp002
					πTemp003 = πg.NewTuple2(ßABC.ToObject(), πg.NewInt(2).ToObject()).ToObject()
					πTemp005 = πg.NewTuple2(ßA.ToObject(), ßB.ToObject()).ToObject()
					πTemp006 = πg.NewTuple2(ßA.ToObject(), ßC.ToObject()).ToObject()
					πTemp007 = πg.NewTuple2(ßB.ToObject(), ßA.ToObject()).ToObject()
					πTemp008 = πg.NewTuple2(ßB.ToObject(), ßC.ToObject()).ToObject()
					πTemp009 = πg.NewTuple2(ßC.ToObject(), ßA.ToObject()).ToObject()
					πTemp010 = πg.NewTuple2(ßC.ToObject(), ßB.ToObject()).ToObject()
					πTemp004 = πg.NewTuple6(πTemp005, πTemp006, πTemp007, πTemp008, πTemp009, πTemp010).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[1] = πTemp002
					πTemp011 = πF.MakeArgs(1)
					πTemp011[0] = πg.NewInt(3).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp011, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp011)
					πTemp003 = πg.NewTuple1(πTemp005).ToObject()
					πTemp005 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject()).ToObject()
					πTemp006 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewInt(2).ToObject(), πg.NewInt(1).ToObject()).ToObject()
					πTemp007 = πg.NewTuple3(πg.NewInt(1).ToObject(), πg.NewInt(0).ToObject(), πg.NewInt(2).ToObject()).ToObject()
					πTemp008 = πg.NewTuple3(πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject(), πg.NewInt(0).ToObject()).ToObject()
					πTemp009 = πg.NewTuple3(πg.NewInt(2).ToObject(), πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject()).ToObject()
					πTemp010 = πg.NewTuple3(πg.NewInt(2).ToObject(), πg.NewInt(1).ToObject(), πg.NewInt(0).ToObject()).ToObject()
					πTemp004 = πg.NewTuple6(πTemp005, πTemp006, πTemp007, πTemp008, πTemp009, πTemp010).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[2] = πTemp002
					πTemp011 = make([]*πg.Object, 0)
					πTemp004 = πg.NewList(πTemp011...).ToObject()
					πTemp003 = πg.NewTuple1(πTemp004).ToObject()
					πTemp005 = πg.NewTuple0().ToObject()
					πTemp004 = πg.NewTuple1(πTemp005).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[3] = πTemp002
					πTemp011 = make([]*πg.Object, 0)
					πTemp004 = πg.NewList(πTemp011...).ToObject()
					πTemp003 = πg.NewTuple2(πTemp004, πg.NewInt(0).ToObject()).ToObject()
					πTemp005 = πg.NewTuple0().ToObject()
					πTemp004 = πg.NewTuple1(πTemp005).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[4] = πTemp002
					πTemp011 = πF.MakeArgs(1)
					πTemp011[0] = πg.NewInt(3).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp011, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp011)
					πTemp003 = πg.NewTuple2(πTemp005, πg.NewInt(4).ToObject()).ToObject()
					πTemp004 = πg.NewTuple0().ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[5] = πTemp002
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µcases = πTemp002
					if πE = πg.CheckLocal(πF, µcases, "cases"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µcases); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp012 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp012 {
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
						πTemp013 = !isStop
					} else {
						πTemp013 = true
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
							continue
						}
						µargs = πTemp004
						µwant = πTemp005
					}
					if πE != nil || !πTemp013 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 145: got = tuple(itertools.permutations(*args))
					πF.SetLineno(145)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßpermutations, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp004, nil, µargs, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µgot = πTemp004
					// line 146: assert got == want, 'tuple(permutations%s) == %s, want %s' % (args, got, want)
					πF.SetLineno(146)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple3(µargs, µgot, µwant).ToObject()
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("tuple(permutations%s) == %s, want %s").ToObject(), πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, µgot, µwant); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp004, πTemp003); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßTestPermutations.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 149: def TestCombinations():
			πF.SetLineno(149)
			πTemp003 = make([]πg.Param, 0)
			πTemp013 = πg.NewFunction(πg.NewCode("TestCombinations", "build/src/__python__/itertools_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcases *πg.Object = πg.UnboundLocal; _ = µcases
				var µargs *πg.Object = πg.UnboundLocal; _ = µargs
				var µwant *πg.Object = πg.UnboundLocal; _ = µwant
				var µgot *πg.Object = πg.UnboundLocal; _ = µgot
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 bool
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
					default: panic("unexpected function state")
					}
					// line 150: cases = [
					πF.SetLineno(150)
					πTemp001 = make([]*πg.Object, 1)
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewInt(4).ToObject()
					if πTemp005, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp003 = πg.NewTuple2(πTemp006, πg.NewInt(3).ToObject()).ToObject()
					πTemp006 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject()).ToObject()
					πTemp007 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject(), πg.NewInt(3).ToObject()).ToObject()
					πTemp008 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewInt(2).ToObject(), πg.NewInt(3).ToObject()).ToObject()
					πTemp009 = πg.NewTuple3(πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject(), πg.NewInt(3).ToObject()).ToObject()
					πTemp005 = πg.NewTuple4(πTemp006, πTemp007, πTemp008, πTemp009).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
					πTemp001[0] = πTemp002
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µcases = πTemp002
					if πE = πg.CheckLocal(πF, µcases, "cases"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µcases); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp010 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp010 {
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
						πTemp011 = !isStop
					} else {
						πTemp011 = true
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp003); πE != nil {
							continue
						}
						µargs = πTemp005
						µwant = πTemp006
					}
					if πE != nil || !πTemp011 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 154: got = tuple(itertools.combinations(*args))
					πF.SetLineno(154)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßcombinations, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp005, nil, µargs, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µgot = πTemp005
					// line 155: assert got == want, 'tuple(combinations%s) == %s, want %s' % (args, got, want)
					πF.SetLineno(155)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple3(µargs, µgot, µwant).ToObject()
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("tuple(combinations%s) == %s, want %s").ToObject(), πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, µgot, µwant); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp005, πTemp003); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßTestCombinations.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 158: def TestCombinationsWithReplacement():
			πF.SetLineno(158)
			πTemp003 = make([]πg.Param, 0)
			πTemp014 = πg.NewFunction(πg.NewCode("TestCombinationsWithReplacement", "build/src/__python__/itertools_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcases *πg.Object = πg.UnboundLocal; _ = µcases
				var µargs *πg.Object = πg.UnboundLocal; _ = µargs
				var µwant *πg.Object = πg.UnboundLocal; _ = µwant
				var µgot *πg.Object = πg.UnboundLocal; _ = µgot
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 bool
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
					default: panic("unexpected function state")
					}
					// line 159: cases = [
					πF.SetLineno(159)
					πTemp001 = make([]*πg.Object, 4)
					πTemp004 = make([]*πg.Object, 1)
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(12).ToObject()); πE != nil {
						continue
					}
					πTemp004[0] = πTemp005
					πTemp005 = πg.NewList(πTemp004...).ToObject()
					πTemp003 = πg.NewTuple2(πTemp005, πg.NewInt(2).ToObject()).ToObject()
					if πTemp007, πE = πg.Neg(πF, πg.NewInt(12).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Neg(πF, πg.NewInt(12).ToObject()); πE != nil {
						continue
					}
					πTemp006 = πg.NewTuple2(πTemp007, πTemp008).ToObject()
					πTemp005 = πg.NewTuple1(πTemp006).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
					πTemp001[0] = πTemp002
					πTemp003 = πg.NewTuple2(ßAB.ToObject(), πg.NewInt(3).ToObject()).ToObject()
					πTemp006 = πg.NewTuple3(ßA.ToObject(), ßA.ToObject(), ßA.ToObject()).ToObject()
					πTemp007 = πg.NewTuple3(ßA.ToObject(), ßA.ToObject(), ßB.ToObject()).ToObject()
					πTemp008 = πg.NewTuple3(ßA.ToObject(), ßB.ToObject(), ßB.ToObject()).ToObject()
					πTemp009 = πg.NewTuple3(ßB.ToObject(), ßB.ToObject(), ßB.ToObject()).ToObject()
					πTemp005 = πg.NewTuple4(πTemp006, πTemp007, πTemp008, πTemp009).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
					πTemp001[1] = πTemp002
					πTemp004 = make([]*πg.Object, 0)
					πTemp005 = πg.NewList(πTemp004...).ToObject()
					πTemp003 = πg.NewTuple2(πTemp005, πg.NewInt(2).ToObject()).ToObject()
					πTemp005 = πg.NewTuple0().ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
					πTemp001[2] = πTemp002
					πTemp004 = make([]*πg.Object, 0)
					πTemp005 = πg.NewList(πTemp004...).ToObject()
					πTemp003 = πg.NewTuple2(πTemp005, πg.NewInt(0).ToObject()).ToObject()
					πTemp006 = πg.NewTuple0().ToObject()
					πTemp005 = πg.NewTuple1(πTemp006).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
					πTemp001[3] = πTemp002
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µcases = πTemp002
					if πE = πg.CheckLocal(πF, µcases, "cases"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µcases); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp010 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp010 {
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
						πTemp011 = !isStop
					} else {
						πTemp011 = true
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp003); πE != nil {
							continue
						}
						µargs = πTemp005
						µwant = πTemp006
					}
					if πE != nil || !πTemp011 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 166: got = tuple(itertools.combinations_with_replacement(*args))
					πF.SetLineno(166)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßcombinations_with_replacement, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp005, nil, µargs, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µgot = πTemp005
					// line 167: assert got == want, 'tuple(combinations_with_replacement%s) == %s, want %s' % (args, got, want)
					πF.SetLineno(167)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple3(µargs, µgot, µwant).ToObject()
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("tuple(combinations_with_replacement%s) == %s, want %s").ToObject(), πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, µgot, µwant); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp005, πTemp003); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßTestCombinationsWithReplacement.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 170: def TestGroupBy():
			πF.SetLineno(170)
			πTemp003 = make([]πg.Param, 0)
			πTemp015 = πg.NewFunction(πg.NewCode("TestGroupBy", "build/src/__python__/itertools_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcases *πg.Object = πg.UnboundLocal; _ = µcases
				var µargs *πg.Object = πg.UnboundLocal; _ = µargs
				var µwant *πg.Object = πg.UnboundLocal; _ = µwant
				var µgot *πg.Object = πg.UnboundLocal; _ = µgot
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
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πTemp010 []πg.Param
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
					// line 171: cases = [
					πF.SetLineno(171)
					πTemp001 = make([]*πg.Object, 2)
					πTemp004 = make([]*πg.Object, 10)
					πTemp004[0] = πg.NewInt(1).ToObject()
					πTemp004[1] = πg.NewInt(2).ToObject()
					πTemp004[2] = πg.NewInt(2).ToObject()
					πTemp004[3] = πg.NewInt(3).ToObject()
					πTemp004[4] = πg.NewInt(3).ToObject()
					πTemp004[5] = πg.NewInt(3).ToObject()
					πTemp004[6] = πg.NewInt(4).ToObject()
					πTemp004[7] = πg.NewInt(4).ToObject()
					πTemp004[8] = πg.NewInt(4).ToObject()
					πTemp004[9] = πg.NewInt(4).ToObject()
					πTemp005 = πg.NewList(πTemp004...).ToObject()
					πTemp003 = πg.NewTuple1(πTemp005).ToObject()
					πTemp004 = make([]*πg.Object, 4)
					πTemp006 = make([]*πg.Object, 1)
					πTemp006[0] = πg.NewInt(1).ToObject()
					πTemp007 = πg.NewList(πTemp006...).ToObject()
					πTemp005 = πg.NewTuple2(πg.NewInt(1).ToObject(), πTemp007).ToObject()
					πTemp004[0] = πTemp005
					πTemp006 = make([]*πg.Object, 2)
					πTemp006[0] = πg.NewInt(2).ToObject()
					πTemp006[1] = πg.NewInt(2).ToObject()
					πTemp007 = πg.NewList(πTemp006...).ToObject()
					πTemp005 = πg.NewTuple2(πg.NewInt(2).ToObject(), πTemp007).ToObject()
					πTemp004[1] = πTemp005
					πTemp006 = make([]*πg.Object, 3)
					πTemp006[0] = πg.NewInt(3).ToObject()
					πTemp006[1] = πg.NewInt(3).ToObject()
					πTemp006[2] = πg.NewInt(3).ToObject()
					πTemp007 = πg.NewList(πTemp006...).ToObject()
					πTemp005 = πg.NewTuple2(πg.NewInt(3).ToObject(), πTemp007).ToObject()
					πTemp004[2] = πTemp005
					πTemp006 = make([]*πg.Object, 4)
					πTemp006[0] = πg.NewInt(4).ToObject()
					πTemp006[1] = πg.NewInt(4).ToObject()
					πTemp006[2] = πg.NewInt(4).ToObject()
					πTemp006[3] = πg.NewInt(4).ToObject()
					πTemp007 = πg.NewList(πTemp006...).ToObject()
					πTemp005 = πg.NewTuple2(πg.NewInt(4).ToObject(), πTemp007).ToObject()
					πTemp004[3] = πTemp005
					πTemp005 = πg.NewList(πTemp004...).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
					πTemp001[0] = πTemp002
					πTemp004 = make([]*πg.Object, 5)
					πTemp004[0] = ßaa.ToObject()
					πTemp004[1] = ßab.ToObject()
					πTemp004[2] = ßabc.ToObject()
					πTemp004[3] = ßbcd.ToObject()
					πTemp004[4] = ßabcde.ToObject()
					πTemp005 = πg.NewList(πTemp004...).ToObject()
					if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp005, πTemp007).ToObject()
					πTemp004 = make([]*πg.Object, 3)
					πTemp006 = make([]*πg.Object, 2)
					πTemp006[0] = ßaa.ToObject()
					πTemp006[1] = ßab.ToObject()
					πTemp007 = πg.NewList(πTemp006...).ToObject()
					πTemp005 = πg.NewTuple2(πg.NewInt(2).ToObject(), πTemp007).ToObject()
					πTemp004[0] = πTemp005
					πTemp006 = make([]*πg.Object, 2)
					πTemp006[0] = ßabc.ToObject()
					πTemp006[1] = ßbcd.ToObject()
					πTemp007 = πg.NewList(πTemp006...).ToObject()
					πTemp005 = πg.NewTuple2(πg.NewInt(3).ToObject(), πTemp007).ToObject()
					πTemp004[1] = πTemp005
					πTemp006 = make([]*πg.Object, 1)
					πTemp006[0] = ßabcde.ToObject()
					πTemp007 = πg.NewList(πTemp006...).ToObject()
					πTemp005 = πg.NewTuple2(πg.NewInt(5).ToObject(), πTemp007).ToObject()
					πTemp004[2] = πTemp005
					πTemp005 = πg.NewList(πTemp004...).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
					πTemp001[1] = πTemp002
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µcases = πTemp002
					if πE = πg.CheckLocal(πF, µcases, "cases"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µcases); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp008 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
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
						πTemp009 = !isStop
					} else {
						πTemp009 = true
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
							continue
						}
						µargs = πTemp005
						µwant = πTemp007
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 176: got = [(k, list(v)) for k, v in itertools.groupby(*args)]
					πF.SetLineno(176)
					πTemp010 = make([]πg.Param, 0)
					πTemp005 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/itertools_test.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
								case 4: goto Label4
								default: panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßgroupby, nil); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Invoke(πF, πTemp003, nil, µargs, nil, nil); πE != nil {
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
								// line 176: got = [(k, list(v)) for k, v in itertools.groupby(*args)]
								πF.SetLineno(176)
								if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
									continue
								}
								πTemp007 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
									continue
								}
								πTemp007[0] = µv
								if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp007)
								πTemp002 = πg.NewTuple2(µk, πTemp006).ToObject()
								πF.PushCheckpoint(4)
								return πTemp002, nil
							Label4:
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
					if πTemp007, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp007}, nil); πE != nil {
						continue
					}
					µgot = πTemp003
					// line 177: assert got == want, 'groupby %s == %s, want %s' % (args, got, want)
					πF.SetLineno(177)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					πTemp007 = πg.NewTuple3(µargs, µgot, µwant).ToObject()
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("groupby %s == %s, want %s").ToObject(), πTemp007); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Eq(πF, µgot, µwant); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp007, πTemp003); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßTestGroupBy.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 180: def TestTakewhile():
			πF.SetLineno(180)
			πTemp003 = make([]πg.Param, 0)
			πTemp016 = πg.NewFunction(πg.NewCode("TestTakewhile", "build/src/__python__/itertools_test.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µr *πg.Object = πg.UnboundLocal; _ = µr
				var µcases *πg.Object = πg.UnboundLocal; _ = µcases
				var µargs *πg.Object = πg.UnboundLocal; _ = µargs
				var µwant *πg.Object = πg.UnboundLocal; _ = µwant
				var µgot *πg.Object = πg.UnboundLocal; _ = µgot
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []πg.Param
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 181: r = range(10)
					πF.SetLineno(181)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewInt(10).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µr = πTemp003
					// line 182: cases = [
					πF.SetLineno(182)
					πTemp001 = make([]*πg.Object, 3)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "x", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/itertools_test.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 183: ((lambda x: x % 2 == 0, r), (0,)),
							πF.SetLineno(183)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, µx, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πg.NewInt(0).ToObject()); πE != nil {
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
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp004, µr).ToObject()
					πTemp004 = πg.NewTuple1(πg.NewInt(0).ToObject()).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[0] = πTemp002
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "x", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/itertools_test.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µx *πg.Object = πArgs[0]; _ = µx
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 184: ((lambda x: True, r), tuple(r)),
							πF.SetLineno(184)
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
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp004, µr).ToObject()
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp006[0] = µr
					if πTemp004, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp002 = πg.NewTuple2(πTemp003, πTemp007).ToObject()
					πTemp001[1] = πTemp002
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "x", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/itertools_test.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µx *πg.Object = πArgs[0]; _ = µx
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 185: ((lambda x: False, r), ())
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
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp004, µr).ToObject()
					πTemp004 = πg.NewTuple0().ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					πTemp001[2] = πTemp002
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µcases = πTemp002
					if πE = πg.CheckLocal(πF, µcases, "cases"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µcases); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp008 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
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
						πTemp009 = !isStop
					} else {
						πTemp009 = true
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
							continue
						}
						µargs = πTemp004
						µwant = πTemp007
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 188: got = tuple(itertools.takewhile(*args))
					πF.SetLineno(188)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßtakewhile, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp004, nil, µargs, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µgot = πTemp004
					// line 189: assert got == want, 'tuple(takewhile%s) == %s, want %s' % (args, got, want)
					πF.SetLineno(189)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple3(µargs, µgot, µwant).ToObject()
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("tuple(takewhile%s) == %s, want %s").ToObject(), πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgot, "got"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwant, "want"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, µgot, µwant); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp004, πTemp003); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßTestTakewhile.ToObject(), πTemp016); πE != nil {
				continue
			}
			if πTemp018, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp017, πE = πg.Eq(πF, πTemp018, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp019, πE = πg.IsTrue(πF, πTemp017); πE != nil {
				continue
			}
			if πTemp019 {
				goto Label1
			}
			goto Label2
			// line 192: if __name__ == '__main__':
			πF.SetLineno(192)
		Label1:
			// line 193: weetest.RunTests()
			πF.SetLineno(193)
			if πTemp017, πE = πg.ResolveGlobal(πF, ßweetest); πE != nil {
				continue
			}
			if πTemp018, πE = πg.GetAttr(πF, πTemp017, ßRunTests, nil); πE != nil {
				continue
			}
			if πTemp017, πE = πTemp018.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("itertools_test", Code)
}
