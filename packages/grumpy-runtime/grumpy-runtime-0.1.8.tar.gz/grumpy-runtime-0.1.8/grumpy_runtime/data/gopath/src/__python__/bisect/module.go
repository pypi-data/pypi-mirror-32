package bisect
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/bisect.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßNone := πg.InternStr("None")
		ßValueError := πg.InternStr("ValueError")
		ßbisect := πg.InternStr("bisect")
		ßbisect_left := πg.InternStr("bisect_left")
		ßbisect_right := πg.InternStr("bisect_right")
		ßinsert := πg.InternStr("insert")
		ßinsort := πg.InternStr("insort")
		ßinsort_left := πg.InternStr("insort_left")
		ßinsort_right := πg.InternStr("insort_right")
		ßlen := πg.InternStr("len")
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
			// line 1: """Bisection algorithms."""
			πF.SetLineno(1)
			// line 3: def insort_right(a, x, lo=0, hi=None):
			πF.SetLineno(3)
			πTemp002 = make([]πg.Param, 4)
			πTemp002[0] = πg.Param{Name: "a", Def: nil}
			πTemp002[1] = πg.Param{Name: "x", Def: nil}
			πTemp002[2] = πg.Param{Name: "lo", Def: πg.NewInt(0).ToObject()}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp002[3] = πg.Param{Name: "hi", Def: πTemp003}
			πTemp001 = πg.NewFunction(πg.NewCode("insort_right", "build/src/__python__/bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µx *πg.Object = πArgs[1]; _ = µx
				var µlo *πg.Object = πArgs[2]; _ = µlo
				var µhi *πg.Object = πArgs[3]; _ = µhi
				var µmid *πg.Object = πg.UnboundLocal; _ = µmid
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
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
					case 5: goto Label5
					case 6: goto Label6
					default: panic("unexpected function state")
					}
					// line 4: """Insert item x in list a, and keep it sorted assuming a is sorted.
					πF.SetLineno(4)
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µlo, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 12: if lo < 0:
					πF.SetLineno(12)
				Label1:
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("lo must be non-negative").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 13: raise ValueError('lo must be non-negative')
					πF.SetLineno(13)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µhi == πTemp004).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 14: if hi is None:
					πF.SetLineno(14)
				Label3:
					// line 15: hi = len(a)
					πF.SetLineno(15)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp003[0] = µa
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µhi = πTemp004
					goto Label4
				Label4:
					// line 16: while lo < hi:
					πF.SetLineno(16)
					πF.PushCheckpoint(6)
					πTemp002 = false
				Label5:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
						πF.PopCheckpoint()
						goto Label7
					}
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µlo, µhi); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(5)            
					// line 17: mid = (lo+hi)//2
					πF.SetLineno(17)
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, µlo, µhi); πE != nil {
						continue
					}
					if πTemp001, πE = πg.FloorDiv(πF, πTemp004, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					µmid = πTemp001
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmid, "mid"); πE != nil {
						continue
					}
					πTemp004 = µmid
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µa, πTemp004); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µx, πTemp006); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label8
					}
					goto Label9
					// line 18: if x < a[mid]: hi = mid
					πF.SetLineno(18)
				Label8:
					// line 18: if x < a[mid]: hi = mid
					πF.SetLineno(18)
					if πE = πg.CheckLocal(πF, µmid, "mid"); πE != nil {
						continue
					}
					µhi = µmid
					goto Label10
				Label9:
					// line 19: else: lo = mid+1
					πF.SetLineno(19)
					if πE = πg.CheckLocal(πF, µmid, "mid"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µmid, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µlo = πTemp001
					goto Label10
				Label10:
					continue
				Label6:
					if πE != nil || πR != nil {
						continue
					}
				Label7:
					// line 20: a.insert(lo, x)
					πF.SetLineno(20)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					πTemp003[0] = µlo
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp003[1] = µx
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µa, ßinsert, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßinsort_right.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 22: insort = insort_right   # backward compatibility
			πF.SetLineno(22)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßinsort_right); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßinsort.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 24: def bisect_right(a, x, lo=0, hi=None):
			πF.SetLineno(24)
			πTemp002 = make([]πg.Param, 4)
			πTemp002[0] = πg.Param{Name: "a", Def: nil}
			πTemp002[1] = πg.Param{Name: "x", Def: nil}
			πTemp002[2] = πg.Param{Name: "lo", Def: πg.NewInt(0).ToObject()}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp002[3] = πg.Param{Name: "hi", Def: πTemp004}
			πTemp003 = πg.NewFunction(πg.NewCode("bisect_right", "build/src/__python__/bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µx *πg.Object = πArgs[1]; _ = µx
				var µlo *πg.Object = πArgs[2]; _ = µlo
				var µhi *πg.Object = πArgs[3]; _ = µhi
				var µmid *πg.Object = πg.UnboundLocal; _ = µmid
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
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
					case 5: goto Label5
					case 6: goto Label6
					default: panic("unexpected function state")
					}
					// line 25: """Return the index where to insert item x in list a, assuming a is sorted.
					πF.SetLineno(25)
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µlo, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 35: if lo < 0:
					πF.SetLineno(35)
				Label1:
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("lo must be non-negative").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 36: raise ValueError('lo must be non-negative')
					πF.SetLineno(36)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µhi == πTemp004).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 37: if hi is None:
					πF.SetLineno(37)
				Label3:
					// line 38: hi = len(a)
					πF.SetLineno(38)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp003[0] = µa
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µhi = πTemp004
					goto Label4
				Label4:
					// line 39: while lo < hi:
					πF.SetLineno(39)
					πF.PushCheckpoint(6)
					πTemp002 = false
				Label5:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
						πF.PopCheckpoint()
						goto Label7
					}
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µlo, µhi); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(5)            
					// line 40: mid = (lo+hi)//2
					πF.SetLineno(40)
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, µlo, µhi); πE != nil {
						continue
					}
					if πTemp001, πE = πg.FloorDiv(πF, πTemp004, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					µmid = πTemp001
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmid, "mid"); πE != nil {
						continue
					}
					πTemp004 = µmid
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µa, πTemp004); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µx, πTemp006); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label8
					}
					goto Label9
					// line 41: if x < a[mid]: hi = mid
					πF.SetLineno(41)
				Label8:
					// line 41: if x < a[mid]: hi = mid
					πF.SetLineno(41)
					if πE = πg.CheckLocal(πF, µmid, "mid"); πE != nil {
						continue
					}
					µhi = µmid
					goto Label10
				Label9:
					// line 42: else: lo = mid+1
					πF.SetLineno(42)
					if πE = πg.CheckLocal(πF, µmid, "mid"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µmid, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µlo = πTemp001
					goto Label10
				Label10:
					continue
				Label6:
					if πE != nil || πR != nil {
						continue
					}
				Label7:
					// line 43: return lo
					πF.SetLineno(43)
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					πR = µlo
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßbisect_right.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 45: bisect = bisect_right   # backward compatibility
			πF.SetLineno(45)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßbisect_right); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßbisect.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 47: def insort_left(a, x, lo=0, hi=None):
			πF.SetLineno(47)
			πTemp002 = make([]πg.Param, 4)
			πTemp002[0] = πg.Param{Name: "a", Def: nil}
			πTemp002[1] = πg.Param{Name: "x", Def: nil}
			πTemp002[2] = πg.Param{Name: "lo", Def: πg.NewInt(0).ToObject()}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp002[3] = πg.Param{Name: "hi", Def: πTemp005}
			πTemp004 = πg.NewFunction(πg.NewCode("insort_left", "build/src/__python__/bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µx *πg.Object = πArgs[1]; _ = µx
				var µlo *πg.Object = πArgs[2]; _ = µlo
				var µhi *πg.Object = πArgs[3]; _ = µhi
				var µmid *πg.Object = πg.UnboundLocal; _ = µmid
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
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
					case 5: goto Label5
					case 6: goto Label6
					default: panic("unexpected function state")
					}
					// line 48: """Insert item x in list a, and keep it sorted assuming a is sorted.
					πF.SetLineno(48)
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µlo, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 56: if lo < 0:
					πF.SetLineno(56)
				Label1:
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("lo must be non-negative").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 57: raise ValueError('lo must be non-negative')
					πF.SetLineno(57)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µhi == πTemp004).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 58: if hi is None:
					πF.SetLineno(58)
				Label3:
					// line 59: hi = len(a)
					πF.SetLineno(59)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp003[0] = µa
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µhi = πTemp004
					goto Label4
				Label4:
					// line 60: while lo < hi:
					πF.SetLineno(60)
					πF.PushCheckpoint(6)
					πTemp002 = false
				Label5:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
						πF.PopCheckpoint()
						goto Label7
					}
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µlo, µhi); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(5)            
					// line 61: mid = (lo+hi)//2
					πF.SetLineno(61)
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, µlo, µhi); πE != nil {
						continue
					}
					if πTemp001, πE = πg.FloorDiv(πF, πTemp004, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					µmid = πTemp001
					if πE = πg.CheckLocal(πF, µmid, "mid"); πE != nil {
						continue
					}
					πTemp004 = µmid
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µa, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, πTemp006, µx); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label8
					}
					goto Label9
					// line 62: if a[mid] < x: lo = mid+1
					πF.SetLineno(62)
				Label8:
					// line 62: if a[mid] < x: lo = mid+1
					πF.SetLineno(62)
					if πE = πg.CheckLocal(πF, µmid, "mid"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µmid, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µlo = πTemp001
					goto Label10
				Label9:
					// line 63: else: hi = mid
					πF.SetLineno(63)
					if πE = πg.CheckLocal(πF, µmid, "mid"); πE != nil {
						continue
					}
					µhi = µmid
					goto Label10
				Label10:
					continue
				Label6:
					if πE != nil || πR != nil {
						continue
					}
				Label7:
					// line 64: a.insert(lo, x)
					πF.SetLineno(64)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					πTemp003[0] = µlo
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp003[1] = µx
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µa, ßinsert, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßinsort_left.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 67: def bisect_left(a, x, lo=0, hi=None):
			πF.SetLineno(67)
			πTemp002 = make([]πg.Param, 4)
			πTemp002[0] = πg.Param{Name: "a", Def: nil}
			πTemp002[1] = πg.Param{Name: "x", Def: nil}
			πTemp002[2] = πg.Param{Name: "lo", Def: πg.NewInt(0).ToObject()}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp002[3] = πg.Param{Name: "hi", Def: πTemp006}
			πTemp005 = πg.NewFunction(πg.NewCode("bisect_left", "build/src/__python__/bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πArgs[0]; _ = µa
				var µx *πg.Object = πArgs[1]; _ = µx
				var µlo *πg.Object = πArgs[2]; _ = µlo
				var µhi *πg.Object = πArgs[3]; _ = µhi
				var µmid *πg.Object = πg.UnboundLocal; _ = µmid
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
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
					case 5: goto Label5
					case 6: goto Label6
					default: panic("unexpected function state")
					}
					// line 68: """Return the index where to insert item x in list a, assuming a is sorted.
					πF.SetLineno(68)
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µlo, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 78: if lo < 0:
					πF.SetLineno(78)
				Label1:
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("lo must be non-negative").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 79: raise ValueError('lo must be non-negative')
					πF.SetLineno(79)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µhi == πTemp004).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 80: if hi is None:
					πF.SetLineno(80)
				Label3:
					// line 81: hi = len(a)
					πF.SetLineno(81)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp003[0] = µa
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µhi = πTemp004
					goto Label4
				Label4:
					// line 82: while lo < hi:
					πF.SetLineno(82)
					πF.PushCheckpoint(6)
					πTemp002 = false
				Label5:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
						πF.PopCheckpoint()
						goto Label7
					}
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µlo, µhi); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(5)            
					// line 83: mid = (lo+hi)//2
					πF.SetLineno(83)
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, µlo, µhi); πE != nil {
						continue
					}
					if πTemp001, πE = πg.FloorDiv(πF, πTemp004, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					µmid = πTemp001
					if πE = πg.CheckLocal(πF, µmid, "mid"); πE != nil {
						continue
					}
					πTemp004 = µmid
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µa, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, πTemp006, µx); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label8
					}
					goto Label9
					// line 84: if a[mid] < x: lo = mid+1
					πF.SetLineno(84)
				Label8:
					// line 84: if a[mid] < x: lo = mid+1
					πF.SetLineno(84)
					if πE = πg.CheckLocal(πF, µmid, "mid"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µmid, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µlo = πTemp001
					goto Label10
				Label9:
					// line 85: else: hi = mid
					πF.SetLineno(85)
					if πE = πg.CheckLocal(πF, µmid, "mid"); πE != nil {
						continue
					}
					µhi = µmid
					goto Label10
				Label10:
					continue
				Label6:
					if πE != nil || πR != nil {
						continue
					}
				Label7:
					// line 86: return lo
					πF.SetLineno(86)
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					πR = µlo
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßbisect_left.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("bisect", Code)
}
