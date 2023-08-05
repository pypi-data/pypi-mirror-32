package time
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/time.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ß01 := πg.InternStr("01")
		ß02 := πg.InternStr("02")
		ß03 := πg.InternStr("03")
		ß04 := πg.InternStr("04")
		ß05 := πg.InternStr("05")
		ß06 := πg.InternStr("06")
		ß15 := πg.InternStr("15")
		ß2006 := πg.InternStr("2006")
		ßA := πg.InternStr("A")
		ßB := πg.InternStr("B")
		ßDate := πg.InternStr("Date")
		ßDay := πg.InternStr("Day")
		ßFormat := πg.InternStr("Format")
		ßH := πg.InternStr("H")
		ßHour := πg.InternStr("Hour")
		ßI := πg.InternStr("I")
		ßJan := πg.InternStr("Jan")
		ßJanuary := πg.InternStr("January")
		ßL := πg.InternStr("L")
		ßLocal := πg.InternStr("Local")
		ßM := πg.InternStr("M")
		ßMST := πg.InternStr("MST")
		ßMinute := πg.InternStr("Minute")
		ßMon := πg.InternStr("Mon")
		ßMonday := πg.InternStr("Monday")
		ßMonth := πg.InternStr("Month")
		ßNone := πg.InternStr("None")
		ßNotImplemented := πg.InternStr("NotImplemented")
		ßNotImplementedError := πg.InternStr("NotImplementedError")
		ßNow := πg.InternStr("Now")
		ßPM := πg.InternStr("PM")
		ßS := πg.InternStr("S")
		ßSecond := πg.InternStr("Second")
		ßSleep := πg.InternStr("Sleep")
		ßU := πg.InternStr("U")
		ßUTC := πg.InternStr("UTC")
		ßUnix := πg.InternStr("Unix")
		ßUnixNano := πg.InternStr("UnixNano")
		ßW := πg.InternStr("W")
		ßWeekday := πg.InternStr("Weekday")
		ßX := πg.InternStr("X")
		ßY := πg.InternStr("Y")
		ßYear := πg.InternStr("Year")
		ßYearDay := πg.InternStr("YearDay")
		ßZ := πg.InternStr("Z")
		ßZone := πg.InternStr("Zone")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__repr__ := πg.InternStr("__repr__")
		ß__str__ := πg.InternStr("__str__")
		ß_strftime_directive_map := πg.InternStr("_strftime_directive_map")
		ßa := πg.InternStr("a")
		ßappend := πg.InternStr("append")
		ßb := πg.InternStr("b")
		ßc := πg.InternStr("c")
		ßd := πg.InternStr("d")
		ßdaylight := πg.InternStr("daylight")
		ßfind := πg.InternStr("find")
		ßfloat := πg.InternStr("float")
		ßget := πg.InternStr("get")
		ßgmtime := πg.InternStr("gmtime")
		ßint := πg.InternStr("int")
		ßj := πg.InternStr("j")
		ßjoin := πg.InternStr("join")
		ßlocaltime := πg.InternStr("localtime")
		ßm := πg.InternStr("m")
		ßmktime := πg.InternStr("mktime")
		ßp := πg.InternStr("p")
		ßrepr := πg.InternStr("repr")
		ßsleep := πg.InternStr("sleep")
		ßstrftime := πg.InternStr("strftime")
		ßstruct_time := πg.InternStr("struct_time")
		ßsuper := πg.InternStr("super")
		ßtime := πg.InternStr("time")
		ßtm_hour := πg.InternStr("tm_hour")
		ßtm_isdst := πg.InternStr("tm_isdst")
		ßtm_mday := πg.InternStr("tm_mday")
		ßtm_min := πg.InternStr("tm_min")
		ßtm_mon := πg.InternStr("tm_mon")
		ßtm_sec := πg.InternStr("tm_sec")
		ßtm_wday := πg.InternStr("tm_wday")
		ßtm_yday := πg.InternStr("tm_yday")
		ßtm_year := πg.InternStr("tm_year")
		ßtuple := πg.InternStr("tuple")
		ßtzname := πg.InternStr("tzname")
		ßw := πg.InternStr("w")
		ßx := πg.InternStr("x")
		ßy := πg.InternStr("y")
		ßz := πg.InternStr("z")
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 15: """Time access and conversions."""
			πF.SetLineno(15)
			// line 17: from '__go__/time' import Local, Now, Second, Sleep, Unix, Date, UTC # pylint: disable=g-multiple-import
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/time"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßLocal, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLocal.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßNow, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNow.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßSecond, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSecond.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßSleep, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSleep.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßUnix, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßUnix.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßDate, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDate.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßUTC, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßUTC.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 20: _strftime_directive_map = {
			πF.SetLineno(20)
			πTemp004 = πg.NewDict()
			if πE = πTemp004.SetItem(πF, πg.NewStr("%").ToObject(), πg.NewStr("%").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßa.ToObject(), ßMon.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßA.ToObject(), ßMonday.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßb.ToObject(), ßJan.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßB.ToObject(), ßJanuary.ToObject()); πE != nil {
				continue
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßc.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßd.ToObject(), ß02.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßH.ToObject(), ß15.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßI.ToObject(), ß03.ToObject()); πE != nil {
				continue
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßj.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßL.ToObject(), πg.NewStr(".000").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßm.ToObject(), ß01.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßM.ToObject(), ß04.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßp.ToObject(), ßPM.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßS.ToObject(), ß05.ToObject()); πE != nil {
				continue
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßU.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßW.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßw.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßX.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßx.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßy.ToObject(), ß06.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßY.ToObject(), ß2006.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßZ.ToObject(), ßMST.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßz.ToObject(), πg.NewStr("-0700").ToObject()); πE != nil {
				continue
			}
			πTemp001 = πTemp004.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_strftime_directive_map.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 48: class struct_time(tuple):  #pylint: disable=invalid-name,missing-docstring
			πF.SetLineno(48)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
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
			_, πE = πg.NewCode("struct_time", "build/src/__python__/time.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 50: def __init__(self, args):
					πF.SetLineno(50)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "args", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/time.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µargs *πg.Object = πArgs[1]; _ = µargs
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
							// line 51: super(struct_time, self).__init__(tuple, args)
							πF.SetLineno(51)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp001[1] = µargs
							πTemp003 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstruct_time); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[1] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.GetAttr(πF, πTemp004, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 52: self.tm_year = self[0]
							πF.SetLineno(52)
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µself, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtm_year, πTemp002); πE != nil {
								continue
							}
							// line 53: self.tm_mon = self[1]
							πF.SetLineno(53)
							πTemp002 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µself, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtm_mon, πTemp002); πE != nil {
								continue
							}
							// line 54: self.tm_mday = self[2]
							πF.SetLineno(54)
							πTemp002 = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µself, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtm_mday, πTemp002); πE != nil {
								continue
							}
							// line 55: self.tm_hour = self[3]
							πF.SetLineno(55)
							πTemp002 = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µself, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtm_hour, πTemp002); πE != nil {
								continue
							}
							// line 56: self.tm_min = self[4]
							πF.SetLineno(56)
							πTemp002 = πg.NewInt(4).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µself, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtm_min, πTemp002); πE != nil {
								continue
							}
							// line 57: self.tm_sec = self[5]
							πF.SetLineno(57)
							πTemp002 = πg.NewInt(5).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µself, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtm_sec, πTemp002); πE != nil {
								continue
							}
							// line 58: self.tm_wday = self[6]
							πF.SetLineno(58)
							πTemp002 = πg.NewInt(6).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µself, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtm_wday, πTemp002); πE != nil {
								continue
							}
							// line 59: self.tm_yday = self[7]
							πF.SetLineno(59)
							πTemp002 = πg.NewInt(7).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µself, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtm_yday, πTemp002); πE != nil {
								continue
							}
							// line 60: self.tm_isdst = self[8]
							πF.SetLineno(60)
							πTemp002 = πg.NewInt(8).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µself, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtm_isdst, πTemp002); πE != nil {
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
					// line 62: def __repr__(self):
					πF.SetLineno(62)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__repr__", "build/src/__python__/time.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 63: return ("time.struct_time(tm_year=%s, tm_mon=%s, tm_mday=%s, "
							πF.SetLineno(63)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("time.struct_time(tm_year=%s, tm_mon=%s, tm_mday=%s, tm_hour=%s, tm_min=%s, tm_sec=%s, tm_wday=%s, tm_yday=%s, tm_isdst=%s)").ToObject(), µself); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 67: def __str__(self):
					πF.SetLineno(67)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__str__", "build/src/__python__/time.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 68: return repr(self)
							πF.SetLineno(68)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__str__.ToObject(), πTemp004); πE != nil {
						continue
					}
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("struct_time").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßstruct_time.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 71: def gmtime(seconds=None):
			πF.SetLineno(71)
			πTemp006 = make([]πg.Param, 1)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[0] = πg.Param{Name: "seconds", Def: πTemp003}
			πTemp001 = πg.NewFunction(πg.NewCode("gmtime", "build/src/__python__/time.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µseconds *πg.Object = πArgs[0]; _ = µseconds
				var µt *πg.Object = πg.UnboundLocal; _ = µt
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
					// line 72: t = (Unix(seconds, 0) if seconds else Now()).UTC()
					πF.SetLineno(72)
					if πE = πg.CheckLocal(πF, µseconds, "seconds"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µseconds); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µseconds, "seconds"); πE != nil {
						continue
					}
					πTemp003[0] = µseconds
					πTemp003[1] = πg.NewInt(0).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßUnix); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001 = πTemp005
					goto Label2
				Label1:
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNow); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp005
				Label2:
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßUTC, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					µt = πTemp001
					// line 73: return struct_time((t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(),
					πF.SetLineno(73)
					πTemp003 = πF.MakeArgs(1)
					πTemp006 = make([]*πg.Object, 9)
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µt, ßYear, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp006[0] = πTemp005
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µt, ßMonth, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp006[1] = πTemp005
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µt, ßDay, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp006[2] = πTemp005
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µt, ßHour, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp006[3] = πTemp005
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µt, ßMinute, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp006[4] = πTemp005
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µt, ßSecond, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp006[5] = πTemp005
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µt, ßWeekday, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Add(πF, πTemp008, πg.NewInt(6).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mod(πF, πTemp005, πg.NewInt(7).ToObject()); πE != nil {
						continue
					}
					πTemp006[6] = πTemp004
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µt, ßYearDay, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp006[7] = πTemp005
					πTemp006[8] = πg.NewInt(0).ToObject()
					πTemp001 = πg.NewTuple(πTemp006...).ToObject()
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßstruct_time); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
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
			if πE = πF.Globals().SetItem(πF, ßgmtime.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 77: def localtime(seconds=None):
			πF.SetLineno(77)
			πTemp006 = make([]πg.Param, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[0] = πg.Param{Name: "seconds", Def: πTemp005}
			πTemp003 = πg.NewFunction(πg.NewCode("localtime", "build/src/__python__/time.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µseconds *πg.Object = πArgs[0]; _ = µseconds
				var µt *πg.Object = πg.UnboundLocal; _ = µt
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
					// line 78: t = (Unix(seconds, 0) if seconds else Now()).Local()
					πF.SetLineno(78)
					if πE = πg.CheckLocal(πF, µseconds, "seconds"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µseconds); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µseconds, "seconds"); πE != nil {
						continue
					}
					πTemp003[0] = µseconds
					πTemp003[1] = πg.NewInt(0).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßUnix); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001 = πTemp005
					goto Label2
				Label1:
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNow); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp005
				Label2:
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßLocal, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					µt = πTemp001
					// line 79: return struct_time((t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(),
					πF.SetLineno(79)
					πTemp003 = πF.MakeArgs(1)
					πTemp006 = make([]*πg.Object, 9)
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µt, ßYear, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp006[0] = πTemp005
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µt, ßMonth, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp006[1] = πTemp005
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µt, ßDay, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp006[2] = πTemp005
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µt, ßHour, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp006[3] = πTemp005
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µt, ßMinute, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp006[4] = πTemp005
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µt, ßSecond, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp006[5] = πTemp005
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µt, ßWeekday, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Add(πF, πTemp008, πg.NewInt(6).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mod(πF, πTemp005, πg.NewInt(7).ToObject()); πE != nil {
						continue
					}
					πTemp006[6] = πTemp004
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µt, ßYearDay, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp006[7] = πTemp005
					πTemp006[8] = πg.NewInt(0).ToObject()
					πTemp001 = πg.NewTuple(πTemp006...).ToObject()
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßstruct_time); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
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
			if πE = πF.Globals().SetItem(πF, ßlocaltime.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 83: def mktime(t):
			πF.SetLineno(83)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "t", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("mktime", "build/src/__python__/time.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µt *πg.Object = πArgs[0]; _ = µt
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
					// line 84: return float(Date(t[0], t[1], t[2], t[3], t[4], t[5], 0, Local).Unix())
					πF.SetLineno(84)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(8)
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µt, πTemp003); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µt, πTemp003); πE != nil {
						continue
					}
					πTemp002[1] = πTemp004
					πTemp003 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µt, πTemp003); πE != nil {
						continue
					}
					πTemp002[2] = πTemp004
					πTemp003 = πg.NewInt(3).ToObject()
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µt, πTemp003); πE != nil {
						continue
					}
					πTemp002[3] = πTemp004
					πTemp003 = πg.NewInt(4).ToObject()
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µt, πTemp003); πE != nil {
						continue
					}
					πTemp002[4] = πTemp004
					πTemp003 = πg.NewInt(5).ToObject()
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µt, πTemp003); πE != nil {
						continue
					}
					πTemp002[5] = πTemp004
					πTemp002[6] = πg.NewInt(0).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßLocal); πE != nil {
						continue
					}
					πTemp002[7] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßDate); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßUnix, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßmktime.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 87: def sleep(secs):
			πF.SetLineno(87)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "secs", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("sleep", "build/src/__python__/time.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsecs *πg.Object = πArgs[0]; _ = µsecs
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
					// line 88: Sleep(secs * Second)
					πF.SetLineno(88)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsecs, "secs"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßSecond); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mul(πF, µsecs, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSleep); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßsleep.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 91: def time():
			πF.SetLineno(91)
			πTemp006 = make([]πg.Param, 0)
			πTemp008 = πg.NewFunction(πg.NewCode("time", "build/src/__python__/time.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
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
					// line 92: return float(Now().UnixNano()) / Second
					πF.SetLineno(92)
					πTemp002 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNow); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßUnixNano, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßSecond); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Div(πF, πTemp004, πTemp003); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßtime.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 95: def strftime(format, tt=None):  # pylint: disable=missing-docstring,redefined-builtin
			πF.SetLineno(95)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "format", Def: nil}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "tt", Def: πTemp010}
			πTemp009 = πg.NewFunction(πg.NewCode("strftime", "build/src/__python__/time.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µformat *πg.Object = πArgs[0]; _ = µformat
				var µtt *πg.Object = πArgs[1]; _ = µtt
				var µt *πg.Object = πg.UnboundLocal; _ = µt
				var µret *πg.Object = πg.UnboundLocal; _ = µret
				var µprev *πg.Object = πg.UnboundLocal; _ = µprev
				var µn *πg.Object = πg.UnboundLocal; _ = µn
				var µnext_ch *πg.Object = πg.UnboundLocal; _ = µnext_ch
				var µc *πg.Object = πg.UnboundLocal; _ = µc
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
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 3: goto Label3
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					// line 96: t = Unix(int(mktime(tt)), 0) if tt else Now()
					πF.SetLineno(96)
					if πE = πg.CheckLocal(πF, µtt, "tt"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µtt); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					πTemp003 = πF.MakeArgs(2)
					πTemp004 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtt, "tt"); πE != nil {
						continue
					}
					πTemp005[0] = µtt
					if πTemp006, πE = πg.ResolveGlobal(πF, ßmktime); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp004[0] = πTemp007
					if πTemp006, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp003[0] = πTemp007
					πTemp003[1] = πg.NewInt(0).ToObject()
					if πTemp006, πE = πg.ResolveGlobal(πF, ßUnix); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001 = πTemp007
					goto Label2
				Label1:
					if πTemp006, πE = πg.ResolveGlobal(πF, ßNow); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp007
				Label2:
					µt = πTemp001
					// line 97: ret = []
					πF.SetLineno(97)
					πTemp003 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					µret = πTemp001
					// line 98: prev, n = 0, format.find('%', 0, -1)
					πF.SetLineno(98)
					πTemp003 = πF.MakeArgs(3)
					πTemp003[0] = πg.NewStr("%").ToObject()
					πTemp003[1] = πg.NewInt(0).ToObject()
					if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003[2] = πTemp006
					if πE = πg.CheckLocal(πF, µformat, "format"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µformat, ßfind, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001 = πg.NewTuple2(πg.NewInt(0).ToObject(), πTemp007).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp001); πE != nil {
						continue
					}
					µprev = πTemp006
					µn = πTemp007
					// line 99: while n != -1:
					πF.SetLineno(99)
					πF.PushCheckpoint(4)
					πTemp002 = false
				Label3:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
						πF.PopCheckpoint()
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, µn, πTemp006); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(3)            
					// line 100: ret.append(format[prev:n])
					πF.SetLineno(100)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µprev, "prev"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µprev, µn, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µformat, "format"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µformat, πTemp001); πE != nil {
						continue
					}
					πTemp003[0] = πTemp006
					if πE = πg.CheckLocal(πF, µret, "ret"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µret, ßappend, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 101: next_ch = format[n + 1]
					πF.SetLineno(101)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Add(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp006
					if πE = πg.CheckLocal(πF, µformat, "format"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µformat, πTemp001); πE != nil {
						continue
					}
					µnext_ch = πTemp006
					// line 102: c = _strftime_directive_map.get(next_ch)
					πF.SetLineno(102)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µnext_ch, "next_ch"); πE != nil {
						continue
					}
					πTemp003[0] = µnext_ch
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_strftime_directive_map); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µc = πTemp001
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µc == πTemp006).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label6
					}
					goto Label7
					// line 103: if c is NotImplemented:
					πF.SetLineno(103)
				Label6:
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µnext_ch, "next_ch"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Add(πF, πg.NewStr("Code: %").ToObject(), µnext_ch); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp006, πg.NewStr(" not yet supported").ToObject()); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 104: raise NotImplementedError('Code: %' + next_ch + ' not yet supported')
					πF.SetLineno(104)
					πE = πF.Raise(πTemp006, nil, nil)
					continue
					goto Label7
				Label7:
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µc); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label8
					}
					goto Label9
					// line 105: if c:
					πF.SetLineno(105)
				Label8:
					// line 106: ret.append(t.Format(c))
					πF.SetLineno(106)
					πTemp003 = πF.MakeArgs(1)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp004[0] = µc
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µt, ßFormat, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp003[0] = πTemp006
					if πE = πg.CheckLocal(πF, µret, "ret"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µret, ßappend, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label10
				Label9:
					// line 108: ret.append(format[n:n+2])
					πF.SetLineno(108)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Add(πF, µn, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µn, πTemp006, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µformat, "format"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µformat, πTemp001); πE != nil {
						continue
					}
					πTemp003[0] = πTemp006
					if πE = πg.CheckLocal(πF, µret, "ret"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µret, ßappend, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label10
				Label10:
					// line 109: n += 2
					πF.SetLineno(109)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µn, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					µn = πTemp001
					// line 110: prev, n = n, format.find('%', n, -1)
					πF.SetLineno(110)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp003 = πF.MakeArgs(3)
					πTemp003[0] = πg.NewStr("%").ToObject()
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp003[1] = µn
					if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003[2] = πTemp006
					if πE = πg.CheckLocal(πF, µformat, "format"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µformat, ßfind, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001 = πg.NewTuple2(µn, πTemp007).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp001); πE != nil {
						continue
					}
					µprev = πTemp006
					µn = πTemp007
					continue
				Label4:
					if πE != nil || πR != nil {
						continue
					}
				Label5:
					// line 111: ret.append(format[prev:])
					πF.SetLineno(111)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µprev, "prev"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µprev, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µformat, "format"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µformat, πTemp001); πE != nil {
						continue
					}
					πTemp003[0] = πTemp006
					if πE = πg.CheckLocal(πF, µret, "ret"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µret, ßappend, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 112: return ''.join(ret)
					πF.SetLineno(112)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µret, "ret"); πE != nil {
						continue
					}
					πTemp003[0] = µret
					if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp006
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßstrftime.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 116: daylight = 0
			πF.SetLineno(116)
			if πE = πF.Globals().SetItem(πF, ßdaylight.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
				continue
			}
			// line 119: tzname = (Now().Zone()[0], '')
			πF.SetLineno(119)
			πTemp011 = πg.NewInt(0).ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßNow); πE != nil {
				continue
			}
			if πTemp014, πE = πTemp013.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp013, πE = πg.GetAttr(πF, πTemp014, ßZone, nil); πE != nil {
				continue
			}
			if πTemp014, πE = πTemp013.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp012, πE = πg.GetItem(πF, πTemp014, πTemp011); πE != nil {
				continue
			}
			πTemp010 = πg.NewTuple2(πTemp012, ß.ToObject()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßtzname.ToObject(), πTemp010); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("time", Code)
}
