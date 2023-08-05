package sched
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/sched.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßEvent := πg.InternStr("Event")
		ß__all__ := πg.InternStr("__all__")
		ß__eq__ := πg.InternStr("__eq__")
		ß__ge__ := πg.InternStr("__ge__")
		ß__gt__ := πg.InternStr("__gt__")
		ß__init__ := πg.InternStr("__init__")
		ß__le__ := πg.InternStr("__le__")
		ß__lt__ := πg.InternStr("__lt__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__slots__ := πg.InternStr("__slots__")
		ß_queue := πg.InternStr("_queue")
		ßaction := πg.InternStr("action")
		ßargument := πg.InternStr("argument")
		ßcancel := πg.InternStr("cancel")
		ßdelayfunc := πg.InternStr("delayfunc")
		ßempty := πg.InternStr("empty")
		ßenter := πg.InternStr("enter")
		ßenterabs := πg.InternStr("enterabs")
		ßget_fields := πg.InternStr("get_fields")
		ßheapify := πg.InternStr("heapify")
		ßheappop := πg.InternStr("heappop")
		ßheappush := πg.InternStr("heappush")
		ßheapq := πg.InternStr("heapq")
		ßlen := πg.InternStr("len")
		ßmap := πg.InternStr("map")
		ßobject := πg.InternStr("object")
		ßpriority := πg.InternStr("priority")
		ßproperty := πg.InternStr("property")
		ßqueue := πg.InternStr("queue")
		ßremove := πg.InternStr("remove")
		ßrun := πg.InternStr("run")
		ßscheduler := πg.InternStr("scheduler")
		ßtime := πg.InternStr("time")
		ßtimefunc := πg.InternStr("timefunc")
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """A generally useful event scheduler class.
			πF.SetLineno(1)
			// line 28: import heapq
			πF.SetLineno(28)
			if πTemp002, πE = πg.ImportModule(πF, "heapq"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßheapq.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 32: __all__ = ["scheduler"]
			πF.SetLineno(32)
			πTemp002 = make([]*πg.Object, 1)
			πTemp002[0] = ßscheduler.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 37: class Event(object):
			πF.SetLineno(37)
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
			_, πE = πg.NewCode("Event", "build/src/__python__/sched.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 39: __slots__ = ['time', 'priority', 'action', 'argument']
					πF.SetLineno(39)
					πTemp001 = make([]*πg.Object, 4)
					πTemp001[0] = ßtime.ToObject()
					πTemp001[1] = ßpriority.ToObject()
					πTemp001[2] = ßaction.ToObject()
					πTemp001[3] = ßargument.ToObject()
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					if πE = πClass.SetItem(πF, ß__slots__.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 41: def __init__(self, time, priority, action, argument):
					πF.SetLineno(41)
					πTemp003 = make([]πg.Param, 5)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "time", Def: nil}
					πTemp003[2] = πg.Param{Name: "priority", Def: nil}
					πTemp003[3] = πg.Param{Name: "action", Def: nil}
					πTemp003[4] = πg.Param{Name: "argument", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/sched.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtime *πg.Object = πArgs[1]; _ = µtime
						var µpriority *πg.Object = πArgs[2]; _ = µpriority
						var µaction *πg.Object = πArgs[3]; _ = µaction
						var µargument *πg.Object = πArgs[4]; _ = µargument
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 42: self.time = time
							πF.SetLineno(42)
							if πE = πg.CheckLocal(πF, µtime, "time"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µtime); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtime, πTemp001); πE != nil {
								continue
							}
							// line 43: self.priority = priority
							πF.SetLineno(43)
							if πE = πg.CheckLocal(πF, µpriority, "priority"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µpriority); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpriority, πTemp001); πE != nil {
								continue
							}
							// line 44: self.action = action
							πF.SetLineno(44)
							if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µaction); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßaction, πTemp001); πE != nil {
								continue
							}
							// line 45: self.argument = argument
							πF.SetLineno(45)
							if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µargument); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßargument, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 47: def get_fields(self):
					πF.SetLineno(47)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("get_fields", "build/src/__python__/sched.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 48: return (self.time, self.priority, self.action, self.argument)
							πF.SetLineno(48)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtime, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpriority, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßaction, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßargument, nil); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple4(πTemp002, πTemp003, πTemp004, πTemp005).ToObject()
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
					if πE = πClass.SetItem(πF, ßget_fields.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 50: def __eq__(s, o): return (s.time, s.priority) == (o.time, o.priority)
					πF.SetLineno(50)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "s", Def: nil}
					πTemp003[1] = πg.Param{Name: "o", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("__eq__", "build/src/__python__/sched.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µs *πg.Object = πArgs[0]; _ = µs
						var µo *πg.Object = πArgs[1]; _ = µo
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
							// line 50: def __eq__(s, o): return (s.time, s.priority) == (o.time, o.priority)
							πF.SetLineno(50)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µs, ßtime, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µs, ßpriority, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µo, ßtime, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µo, ßpriority, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 51: def __lt__(s, o): return (s.time, s.priority) <  (o.time, o.priority)
					πF.SetLineno(51)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "s", Def: nil}
					πTemp003[1] = πg.Param{Name: "o", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("__lt__", "build/src/__python__/sched.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µs *πg.Object = πArgs[0]; _ = µs
						var µo *πg.Object = πArgs[1]; _ = µo
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
							// line 51: def __lt__(s, o): return (s.time, s.priority) <  (o.time, o.priority)
							πF.SetLineno(51)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µs, ßtime, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µs, ßpriority, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µo, ßtime, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µo, ßpriority, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
							if πTemp001, πE = πg.LT(πF, πTemp002, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__lt__.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 52: def __le__(s, o): return (s.time, s.priority) <= (o.time, o.priority)
					πF.SetLineno(52)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "s", Def: nil}
					πTemp003[1] = πg.Param{Name: "o", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("__le__", "build/src/__python__/sched.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µs *πg.Object = πArgs[0]; _ = µs
						var µo *πg.Object = πArgs[1]; _ = µo
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
							// line 52: def __le__(s, o): return (s.time, s.priority) <= (o.time, o.priority)
							πF.SetLineno(52)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µs, ßtime, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µs, ßpriority, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µo, ßtime, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µo, ßpriority, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
							if πTemp001, πE = πg.LE(πF, πTemp002, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__le__.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 53: def __gt__(s, o): return (s.time, s.priority) >  (o.time, o.priority)
					πF.SetLineno(53)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "s", Def: nil}
					πTemp003[1] = πg.Param{Name: "o", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("__gt__", "build/src/__python__/sched.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µs *πg.Object = πArgs[0]; _ = µs
						var µo *πg.Object = πArgs[1]; _ = µo
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
							// line 53: def __gt__(s, o): return (s.time, s.priority) >  (o.time, o.priority)
							πF.SetLineno(53)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µs, ßtime, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µs, ßpriority, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µo, ßtime, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µo, ßpriority, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
							if πTemp001, πE = πg.GT(πF, πTemp002, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__gt__.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 54: def __ge__(s, o): return (s.time, s.priority) >= (o.time, o.priority)
					πF.SetLineno(54)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "s", Def: nil}
					πTemp003[1] = πg.Param{Name: "o", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("__ge__", "build/src/__python__/sched.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µs *πg.Object = πArgs[0]; _ = µs
						var µo *πg.Object = πArgs[1]; _ = µo
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
							// line 54: def __ge__(s, o): return (s.time, s.priority) >= (o.time, o.priority)
							πF.SetLineno(54)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µs, ßtime, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µs, ßpriority, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µo, ßtime, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µo, ßpriority, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
							if πTemp001, πE = πg.GE(πF, πTemp002, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__ge__.ToObject(), πTemp009); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("Event").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßEvent.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 56: class scheduler(object):
			πF.SetLineno(56)
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
			_, πE = πg.NewCode("scheduler", "build/src/__python__/sched.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp009 []*πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 57: def __init__(self, timefunc, delayfunc):
					πF.SetLineno(57)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "timefunc", Def: nil}
					πTemp002[2] = πg.Param{Name: "delayfunc", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/sched.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtimefunc *πg.Object = πArgs[1]; _ = µtimefunc
						var µdelayfunc *πg.Object = πArgs[2]; _ = µdelayfunc
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
							// line 58: """Initialize a new instance, passing the time and delay
							πF.SetLineno(58)
							// line 60: self._queue = []
							πF.SetLineno(60)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_queue, πTemp003); πE != nil {
								continue
							}
							// line 61: self.timefunc = timefunc
							πF.SetLineno(61)
							if πE = πg.CheckLocal(πF, µtimefunc, "timefunc"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µtimefunc); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtimefunc, πTemp002); πE != nil {
								continue
							}
							// line 62: self.delayfunc = delayfunc
							πF.SetLineno(62)
							if πE = πg.CheckLocal(πF, µdelayfunc, "delayfunc"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µdelayfunc); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdelayfunc, πTemp002); πE != nil {
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
					// line 64: def enterabs(self, time, priority, action, argument):
					πF.SetLineno(64)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "time", Def: nil}
					πTemp002[2] = πg.Param{Name: "priority", Def: nil}
					πTemp002[3] = πg.Param{Name: "action", Def: nil}
					πTemp002[4] = πg.Param{Name: "argument", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("enterabs", "build/src/__python__/sched.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtime *πg.Object = πArgs[1]; _ = µtime
						var µpriority *πg.Object = πArgs[2]; _ = µpriority
						var µaction *πg.Object = πArgs[3]; _ = µaction
						var µargument *πg.Object = πArgs[4]; _ = µargument
						var µevent *πg.Object = πg.UnboundLocal; _ = µevent
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
							// line 65: """Enter a new event in the queue at an absolute time.
							πF.SetLineno(65)
							// line 69: event = Event(time, priority, action, argument)
							πF.SetLineno(69)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µtime, "time"); πE != nil {
								continue
							}
							πTemp001[0] = µtime
							if πE = πg.CheckLocal(πF, µpriority, "priority"); πE != nil {
								continue
							}
							πTemp001[1] = µpriority
							if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
								continue
							}
							πTemp001[2] = µaction
							if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
								continue
							}
							πTemp001[3] = µargument
							if πTemp002, πE = πg.ResolveGlobal(πF, ßEvent); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µevent = πTemp003
							// line 70: heapq.heappush(self._queue, event)
							πF.SetLineno(70)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_queue, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µevent, "event"); πE != nil {
								continue
							}
							πTemp001[1] = µevent
							if πTemp002, πE = πg.ResolveGlobal(πF, ßheapq); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßheappush, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 71: return event # The ID
							πF.SetLineno(71)
							if πE = πg.CheckLocal(πF, µevent, "event"); πE != nil {
								continue
							}
							πR = µevent
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßenterabs.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 73: def enter(self, delay, priority, action, argument):
					πF.SetLineno(73)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "delay", Def: nil}
					πTemp002[2] = πg.Param{Name: "priority", Def: nil}
					πTemp002[3] = πg.Param{Name: "action", Def: nil}
					πTemp002[4] = πg.Param{Name: "argument", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("enter", "build/src/__python__/sched.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdelay *πg.Object = πArgs[1]; _ = µdelay
						var µpriority *πg.Object = πArgs[2]; _ = µpriority
						var µaction *πg.Object = πArgs[3]; _ = µaction
						var µargument *πg.Object = πArgs[4]; _ = µargument
						var µtime *πg.Object = πg.UnboundLocal; _ = µtime
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 74: """A variant that specifies the time as a relative time.
							πF.SetLineno(74)
							// line 77: time = self.timefunc() + delay
							πF.SetLineno(77)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtimefunc, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdelay, "delay"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, µdelay); πE != nil {
								continue
							}
							µtime = πTemp001
							// line 78: return self.enterabs(time, priority, action, argument)
							πF.SetLineno(78)
							πTemp004 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µtime, "time"); πE != nil {
								continue
							}
							πTemp004[0] = µtime
							if πE = πg.CheckLocal(πF, µpriority, "priority"); πE != nil {
								continue
							}
							πTemp004[1] = µpriority
							if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
								continue
							}
							πTemp004[2] = µaction
							if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
								continue
							}
							πTemp004[3] = µargument
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßenterabs, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßenter.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 80: def cancel(self, event):
					πF.SetLineno(80)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "event", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("cancel", "build/src/__python__/sched.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µevent *πg.Object = πArgs[1]; _ = µevent
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
							// line 81: """Remove an event from the queue.
							πF.SetLineno(81)
							// line 85: self._queue.remove(event)
							πF.SetLineno(85)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µevent, "event"); πE != nil {
								continue
							}
							πTemp001[0] = µevent
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_queue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßremove, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 86: heapq.heapify(self._queue)
							πF.SetLineno(86)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_queue, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßheapq); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßheapify, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcancel.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 88: def empty(self):
					πF.SetLineno(88)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("empty", "build/src/__python__/sched.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 89: """Check whether the queue is empty."""
							πF.SetLineno(89)
							// line 90: return not self._queue
							πF.SetLineno(90)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_queue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
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
					if πE = πClass.SetItem(πF, ßempty.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 92: def run(self):
					πF.SetLineno(92)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("run", "build/src/__python__/sched.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µq *πg.Object = πg.UnboundLocal; _ = µq
						var µdelayfunc *πg.Object = πg.UnboundLocal; _ = µdelayfunc
						var µtimefunc *πg.Object = πg.UnboundLocal; _ = µtimefunc
						var µpop *πg.Object = πg.UnboundLocal; _ = µpop
						var µchecked_event *πg.Object = πg.UnboundLocal; _ = µchecked_event
						var µtime *πg.Object = πg.UnboundLocal; _ = µtime
						var µpriority *πg.Object = πg.UnboundLocal; _ = µpriority
						var µaction *πg.Object = πg.UnboundLocal; _ = µaction
						var µargument *πg.Object = πg.UnboundLocal; _ = µargument
						var µnow *πg.Object = πg.UnboundLocal; _ = µnow
						var µevent *πg.Object = πg.UnboundLocal; _ = µevent
						var πTemp001 *πg.Object
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
						var πTemp007 *πg.Object
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
							// line 93: """Execute events until the queue is empty.
							πF.SetLineno(93)
							// line 111: q = self._queue
							πF.SetLineno(111)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_queue, nil); πE != nil {
								continue
							}
							µq = πTemp001
							// line 112: delayfunc = self.delayfunc
							πF.SetLineno(112)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdelayfunc, nil); πE != nil {
								continue
							}
							µdelayfunc = πTemp001
							// line 113: timefunc = self.timefunc
							πF.SetLineno(113)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtimefunc, nil); πE != nil {
								continue
							}
							µtimefunc = πTemp001
							// line 114: pop = heapq.heappop
							πF.SetLineno(114)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßheapq); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßheappop, nil); πE != nil {
								continue
							}
							µpop = πTemp002
							// line 115: while q:
							πF.SetLineno(115)
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
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µq); πE != nil {
								continue
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 117: checked_event = q[0]
							πF.SetLineno(117)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µq, πTemp001); πE != nil {
								continue
							}
							µchecked_event = πTemp002
							// line 118: time, priority, action, argument = checked_event.get_fields()
							πF.SetLineno(118)
							if πE = πg.CheckLocal(πF, µchecked_event, "checked_event"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µchecked_event, ßget_fields, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp002); πE != nil {
								continue
							}
							µtime = πTemp001
							µpriority = πTemp005
							µaction = πTemp006
							µargument = πTemp007
							// line 119: now = timefunc()
							πF.SetLineno(119)
							if πE = πg.CheckLocal(πF, µtimefunc, "timefunc"); πE != nil {
								continue
							}
							if πTemp001, πE = µtimefunc.Call(πF, nil, nil); πE != nil {
								continue
							}
							µnow = πTemp001
							if πE = πg.CheckLocal(πF, µnow, "now"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtime, "time"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µnow, µtime); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 120: if now < time:
							πF.SetLineno(120)
						Label4:
							// line 121: delayfunc(time - now)
							πF.SetLineno(121)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtime, "time"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnow, "now"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µtime, µnow); πE != nil {
								continue
							}
							πTemp008[0] = πTemp001
							if πE = πg.CheckLocal(πF, µdelayfunc, "delayfunc"); πE != nil {
								continue
							}
							if πTemp001, πE = µdelayfunc.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							goto Label6
						Label5:
							// line 123: event = pop(q)
							πF.SetLineno(123)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							πTemp008[0] = µq
							if πE = πg.CheckLocal(πF, µpop, "pop"); πE != nil {
								continue
							}
							if πTemp001, πE = µpop.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							µevent = πTemp001
							if πE = πg.CheckLocal(πF, µevent, "event"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µchecked_event, "checked_event"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µevent == µchecked_event).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							goto Label8
							// line 126: if event is checked_event:
							πF.SetLineno(126)
						Label7:
							// line 127: action(*argument)
							πF.SetLineno(127)
							if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Invoke(πF, µaction, nil, µargument, nil, nil); πE != nil {
								continue
							}
							// line 128: delayfunc(0)   # Let other threads run
							πF.SetLineno(128)
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µdelayfunc, "delayfunc"); πE != nil {
								continue
							}
							if πTemp001, πE = µdelayfunc.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							goto Label9
						Label8:
							// line 130: heapq.heappush(q, event)
							πF.SetLineno(130)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							πTemp008[0] = µq
							if πE = πg.CheckLocal(πF, µevent, "event"); πE != nil {
								continue
							}
							πTemp008[1] = µevent
							if πTemp001, πE = πg.ResolveGlobal(πF, ßheapq); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßheappush, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							goto Label9
						Label9:
							goto Label6
						Label6:
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
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 133: def queue(self):
					πF.SetLineno(133)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("queue", "build/src/__python__/sched.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µevents *πg.Object = πg.UnboundLocal; _ = µevents
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
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
							// line 134: """An ordered list of upcoming events.
							πF.SetLineno(134)
							// line 141: events = self._queue[:]
							πF.SetLineno(141)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_queue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							µevents = πTemp002
							// line 142: return map(heapq.heappop, [events]*len(events))
							πF.SetLineno(142)
							πTemp004 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßheapq); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßheappop, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							πTemp005 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µevents, "events"); πE != nil {
								continue
							}
							πTemp005[0] = µevents
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µevents, "events"); πE != nil {
								continue
							}
							πTemp005[0] = µevents
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.Mul(πF, πTemp002, πTemp006); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßqueue.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 132: @property
					πF.SetLineno(132)
					πTemp009 = πF.MakeArgs(1)
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßqueue); πE != nil {
						continue
					}
					πTemp009[0] = πTemp010
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					if πE = πClass.SetItem(πF, ßqueue.ToObject(), πTemp011); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("scheduler").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßscheduler.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("sched", Code)
}
