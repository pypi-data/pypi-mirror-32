package select_
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/select_.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßBits := πg.InternStr("Bits")
		ßException := πg.InternStr("Exception")
		ßFD_SETSIZE := πg.InternStr("FD_SETSIZE")
		ßFdSet := πg.InternStr("FdSet")
		ßNone := πg.InternStr("None")
		ßSec := πg.InternStr("Sec")
		ßSelect := πg.InternStr("Select")
		ßTimeval := πg.InternStr("Timeval")
		ßUsec := πg.InternStr("Usec")
		ß_FD_SETSIZE := πg.InternStr("_FD_SETSIZE")
		ß_FdSet := πg.InternStr("_FdSet")
		ß_Select := πg.InternStr("_Select")
		ß_Timeval := πg.InternStr("_Timeval")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_fdset_isset := πg.InternStr("_fdset_isset")
		ß_fdset_set := πg.InternStr("_fdset_set")
		ß_make_fdset := πg.InternStr("_make_fdset")
		ß_normalize_fd_list := πg.InternStr("_normalize_fd_list")
		ß_syscall := πg.InternStr("_syscall")
		ßappend := πg.InternStr("append")
		ßbool := πg.InternStr("bool")
		ßenumerate := πg.InternStr("enumerate")
		ßerror := πg.InternStr("error")
		ßfileno := πg.InternStr("fileno")
		ßhasattr := πg.InternStr("hasattr")
		ßint := πg.InternStr("int")
		ßinvoke := πg.InternStr("invoke")
		ßlen := πg.InternStr("len")
		ßmath := πg.InternStr("math")
		ßmax := πg.InternStr("max")
		ßmodf := πg.InternStr("modf")
		ßnew := πg.InternStr("new")
		ßselect := πg.InternStr("select")
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 15: from '__go__/syscall' import (
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/syscall"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßFD_SETSIZE, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_FD_SETSIZE.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßSelect, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Select.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßFdSet, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_FdSet.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßTimeval, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Timeval.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 21: import _syscall
			πF.SetLineno(21)
			if πTemp002, πE = πg.ImportModule(πF, "_syscall"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_syscall.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 22: import math
			πF.SetLineno(22)
			if πTemp002, πE = πg.ImportModule(πF, "math"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßmath.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 25: class error(Exception):
			πF.SetLineno(25)
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
			_, πE = πg.NewCode("error", "build/src/__python__/select_.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 26: pass
					πF.SetLineno(26)
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("error").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßerror.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 29: def select(rlist, wlist, xlist, timeout=None):
			πF.SetLineno(29)
			πTemp006 = make([]πg.Param, 4)
			πTemp006[0] = πg.Param{Name: "rlist", Def: nil}
			πTemp006[1] = πg.Param{Name: "wlist", Def: nil}
			πTemp006[2] = πg.Param{Name: "xlist", Def: nil}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[3] = πg.Param{Name: "timeout", Def: πTemp003}
			πTemp001 = πg.NewFunction(πg.NewCode("select", "build/src/__python__/select_.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µrlist *πg.Object = πArgs[0]; _ = µrlist
				var µwlist *πg.Object = πArgs[1]; _ = µwlist
				var µxlist *πg.Object = πArgs[2]; _ = µxlist
				var µtimeout *πg.Object = πArgs[3]; _ = µtimeout
				var µrlist_norm *πg.Object = πg.UnboundLocal; _ = µrlist_norm
				var µwlist_norm *πg.Object = πg.UnboundLocal; _ = µwlist_norm
				var µxlist_norm *πg.Object = πg.UnboundLocal; _ = µxlist_norm
				var µall_fds *πg.Object = πg.UnboundLocal; _ = µall_fds
				var µnfd *πg.Object = πg.UnboundLocal; _ = µnfd
				var µrfds *πg.Object = πg.UnboundLocal; _ = µrfds
				var µwfds *πg.Object = πg.UnboundLocal; _ = µwfds
				var µxfds *πg.Object = πg.UnboundLocal; _ = µxfds
				var µtimeval *πg.Object = πg.UnboundLocal; _ = µtimeval
				var µfrac *πg.Object = πg.UnboundLocal; _ = µfrac
				var µinteger *πg.Object = πg.UnboundLocal; _ = µinteger
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 30: rlist_norm = _normalize_fd_list(rlist)
					πF.SetLineno(30)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µrlist, "rlist"); πE != nil {
						continue
					}
					πTemp001[0] = µrlist
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_normalize_fd_list); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µrlist_norm = πTemp003
					// line 31: wlist_norm = _normalize_fd_list(wlist)
					πF.SetLineno(31)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µwlist, "wlist"); πE != nil {
						continue
					}
					πTemp001[0] = µwlist
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_normalize_fd_list); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µwlist_norm = πTemp003
					// line 32: xlist_norm = _normalize_fd_list(xlist)
					πF.SetLineno(32)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µxlist, "xlist"); πE != nil {
						continue
					}
					πTemp001[0] = µxlist
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_normalize_fd_list); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µxlist_norm = πTemp003
					// line 33: all_fds = rlist_norm + wlist_norm + xlist_norm
					πF.SetLineno(33)
					if πE = πg.CheckLocal(πF, µrlist_norm, "rlist_norm"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwlist_norm, "wlist_norm"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µrlist_norm, µwlist_norm); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µxlist_norm, "xlist_norm"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp003, µxlist_norm); πE != nil {
						continue
					}
					µall_fds = πTemp002
					if πE = πg.CheckLocal(πF, µall_fds, "all_fds"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µall_fds); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 34: if not all_fds:
					πF.SetLineno(34)
				Label1:
					// line 35: nfd = 0
					πF.SetLineno(35)
					µnfd = πg.NewInt(0).ToObject()
					goto Label3
				Label2:
					// line 37: nfd = max(all_fds) + 1
					πF.SetLineno(37)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µall_fds, "all_fds"); πE != nil {
						continue
					}
					πTemp001[0] = µall_fds
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Add(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µnfd = πTemp002
					goto Label3
				Label3:
					// line 39: rfds = _make_fdset(rlist_norm)
					πF.SetLineno(39)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µrlist_norm, "rlist_norm"); πE != nil {
						continue
					}
					πTemp001[0] = µrlist_norm
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_make_fdset); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µrfds = πTemp003
					// line 40: wfds = _make_fdset(wlist_norm)
					πF.SetLineno(40)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µwlist_norm, "wlist_norm"); πE != nil {
						continue
					}
					πTemp001[0] = µwlist_norm
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_make_fdset); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µwfds = πTemp003
					// line 41: xfds = _make_fdset(xlist_norm)
					πF.SetLineno(41)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µxlist_norm, "xlist_norm"); πE != nil {
						continue
					}
					πTemp001[0] = µxlist_norm
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_make_fdset); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µxfds = πTemp003
					if πE = πg.CheckLocal(πF, µtimeout, "timeout"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µtimeout == πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					goto Label5
					// line 43: if timeout is None:
					πF.SetLineno(43)
				Label4:
					// line 44: timeval = None
					πF.SetLineno(44)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µtimeval = πTemp002
					goto Label6
				Label5:
					// line 46: timeval = _Timeval.new()
					πF.SetLineno(46)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_Timeval); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßnew, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					µtimeval = πTemp002
					// line 47: frac, integer = math.modf(timeout)
					πF.SetLineno(47)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtimeout, "timeout"); πE != nil {
						continue
					}
					πTemp001[0] = µtimeout
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmodf, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
						continue
					}
					µfrac = πTemp003
					µinteger = πTemp005
					// line 48: timeval.Sec = int(integer)
					πF.SetLineno(48)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µinteger, "integer"); πE != nil {
						continue
					}
					πTemp001[0] = µinteger
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtimeval, "timeval"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µtimeval, ßSec, πTemp002); πE != nil {
						continue
					}
					// line 49: timeval.Usec = int(frac * 1000000.0)
					πF.SetLineno(49)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfrac, "frac"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mul(πF, µfrac, πg.NewFloat(1000000.0).ToObject()); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtimeval, "timeval"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µtimeval, ßUsec, πTemp002); πE != nil {
						continue
					}
					goto Label6
				Label6:
					// line 50: _syscall.invoke(_Select, nfd, rfds, wfds, xfds, timeval)
					πF.SetLineno(50)
					πTemp001 = πF.MakeArgs(6)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_Select); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µnfd, "nfd"); πE != nil {
						continue
					}
					πTemp001[1] = µnfd
					if πE = πg.CheckLocal(πF, µrfds, "rfds"); πE != nil {
						continue
					}
					πTemp001[2] = µrfds
					if πE = πg.CheckLocal(πF, µwfds, "wfds"); πE != nil {
						continue
					}
					πTemp001[3] = µwfds
					if πE = πg.CheckLocal(πF, µxfds, "xfds"); πE != nil {
						continue
					}
					πTemp001[4] = µxfds
					if πE = πg.CheckLocal(πF, µtimeval, "timeval"); πE != nil {
						continue
					}
					πTemp001[5] = µtimeval
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_syscall); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinvoke, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 51: return ([rlist[i] for i, fd in enumerate(rlist_norm) if _fdset_isset(fd, rfds)],
					πF.SetLineno(51)
					πTemp006 = make([]πg.Param, 0)
					πTemp005 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/select_.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µfd *πg.Object = πg.UnboundLocal; _ = µfd
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
								πTemp002 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µrlist_norm, "rlist_norm"); πE != nil {
									continue
								}
								πTemp002[0] = µrlist_norm
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
									µi = πTemp004
									µfd = πTemp007
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)            
								πTemp002 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
									continue
								}
								πTemp002[0] = µfd
								if πE = πg.CheckLocal(πF, µrfds, "rfds"); πE != nil {
									continue
								}
								πTemp002[1] = µrfds
								if πTemp003, πE = πg.ResolveGlobal(πF, ß_fdset_isset); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label4
								}
								goto Label5
								// line 51: return ([rlist[i] for i, fd in enumerate(rlist_norm) if _fdset_isset(fd, rfds)],
								πF.SetLineno(51)
							Label4:
								// line 51: return ([rlist[i] for i, fd in enumerate(rlist_norm) if _fdset_isset(fd, rfds)],
								πF.SetLineno(51)
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								πTemp003 = µi
								if πE = πg.CheckLocal(πF, µrlist, "rlist"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetItem(πF, µrlist, πTemp003); πE != nil {
									continue
								}
								πF.PushCheckpoint(6)
								return πTemp004, nil
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
					if πTemp007, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp007}, nil); πE != nil {
						continue
					}
					πTemp006 = make([]πg.Param, 0)
					πTemp008 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/select_.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µfd *πg.Object = πg.UnboundLocal; _ = µfd
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
								πTemp002 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µwlist_norm, "wlist_norm"); πE != nil {
									continue
								}
								πTemp002[0] = µwlist_norm
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
									µi = πTemp004
									µfd = πTemp007
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)            
								πTemp002 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
									continue
								}
								πTemp002[0] = µfd
								if πE = πg.CheckLocal(πF, µwfds, "wfds"); πE != nil {
									continue
								}
								πTemp002[1] = µwfds
								if πTemp003, πE = πg.ResolveGlobal(πF, ß_fdset_isset); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label4
								}
								goto Label5
								// line 52: [wlist[i] for i, fd in enumerate(wlist_norm) if _fdset_isset(fd, wfds)],
								πF.SetLineno(52)
							Label4:
								// line 52: [wlist[i] for i, fd in enumerate(wlist_norm) if _fdset_isset(fd, wfds)],
								πF.SetLineno(52)
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								πTemp003 = µi
								if πE = πg.CheckLocal(πF, µwlist, "wlist"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetItem(πF, µwlist, πTemp003); πE != nil {
									continue
								}
								πF.PushCheckpoint(6)
								return πTemp004, nil
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
					if πTemp009, πE = πTemp008.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ListType.Call(πF, πg.Args{πTemp009}, nil); πE != nil {
						continue
					}
					πTemp006 = make([]πg.Param, 0)
					πTemp010 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/select_.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µfd *πg.Object = πg.UnboundLocal; _ = µfd
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
								πTemp002 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µxlist_norm, "xlist_norm"); πE != nil {
									continue
								}
								πTemp002[0] = µxlist_norm
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
									µi = πTemp004
									µfd = πTemp007
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)            
								πTemp002 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
									continue
								}
								πTemp002[0] = µfd
								if πE = πg.CheckLocal(πF, µxfds, "xfds"); πE != nil {
									continue
								}
								πTemp002[1] = µxfds
								if πTemp003, πE = πg.ResolveGlobal(πF, ß_fdset_isset); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label4
								}
								goto Label5
								// line 53: [xlist[i] for i, fd in enumerate(xlist_norm) if _fdset_isset(fd, xfds)])
								πF.SetLineno(53)
							Label4:
								// line 53: [xlist[i] for i, fd in enumerate(xlist_norm) if _fdset_isset(fd, xfds)])
								πF.SetLineno(53)
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								πTemp003 = µi
								if πE = πg.CheckLocal(πF, µxlist, "xlist"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetItem(πF, µxlist, πTemp003); πE != nil {
									continue
								}
								πF.PushCheckpoint(6)
								return πTemp004, nil
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
					if πTemp011, πE = πTemp010.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ListType.Call(πF, πg.Args{πTemp011}, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(πTemp003, πTemp007, πTemp009).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßselect.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 56: def _fdset_set(fd, fds):
			πF.SetLineno(56)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "fd", Def: nil}
			πTemp006[1] = πg.Param{Name: "fds", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("_fdset_set", "build/src/__python__/select_.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfd *πg.Object = πArgs[0]; _ = µfd
				var µfds *πg.Object = πArgs[1]; _ = µfds
				var µidx *πg.Object = πg.UnboundLocal; _ = µidx
				var µpos *πg.Object = πg.UnboundLocal; _ = µpos
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
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
					// line 57: idx = fd / (_FD_SETSIZE / len(fds.Bits)) % len(fds.Bits)
					πF.SetLineno(57)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_FD_SETSIZE); πE != nil {
						continue
					}
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfds, "fds"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µfds, ßBits, nil); πE != nil {
						continue
					}
					πTemp005[0] = πTemp006
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp003, πE = πg.Div(πF, πTemp004, πTemp007); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Div(πF, µfd, πTemp003); πE != nil {
						continue
					}
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfds, "fds"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µfds, ßBits, nil); πE != nil {
						continue
					}
					πTemp005[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp001, πE = πg.Mod(πF, πTemp002, πTemp004); πE != nil {
						continue
					}
					µidx = πTemp001
					// line 58: pos = fd % (_FD_SETSIZE / len(fds.Bits))
					πF.SetLineno(58)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_FD_SETSIZE); πE != nil {
						continue
					}
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfds, "fds"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µfds, ßBits, nil); πE != nil {
						continue
					}
					πTemp005[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp002, πE = πg.Div(πF, πTemp003, πTemp006); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, µfd, πTemp002); πE != nil {
						continue
					}
					µpos = πTemp001
					// line 59: fds.Bits[idx] |= 1 << pos
					πF.SetLineno(59)
					if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
						continue
					}
					πTemp001 = µidx
					if πE = πg.CheckLocal(πF, µfds, "fds"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µfds, ßBits, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LShift(πF, πg.NewInt(1).ToObject(), µpos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IOr(πF, πTemp002, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfds, "fds"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µfds, ßBits, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
						continue
					}
					πTemp006 = µidx
					if πE = πg.SetItem(πF, πTemp004, πTemp006, πTemp003); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_fdset_set.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 62: def _fdset_isset(fd, fds):
			πF.SetLineno(62)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "fd", Def: nil}
			πTemp006[1] = πg.Param{Name: "fds", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("_fdset_isset", "build/src/__python__/select_.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfd *πg.Object = πArgs[0]; _ = µfd
				var µfds *πg.Object = πArgs[1]; _ = µfds
				var µidx *πg.Object = πg.UnboundLocal; _ = µidx
				var µpos *πg.Object = πg.UnboundLocal; _ = µpos
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
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
					// line 63: idx = fd / (_FD_SETSIZE / len(fds.Bits)) % len(fds.Bits)
					πF.SetLineno(63)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_FD_SETSIZE); πE != nil {
						continue
					}
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfds, "fds"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µfds, ßBits, nil); πE != nil {
						continue
					}
					πTemp005[0] = πTemp006
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp003, πE = πg.Div(πF, πTemp004, πTemp007); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Div(πF, µfd, πTemp003); πE != nil {
						continue
					}
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfds, "fds"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µfds, ßBits, nil); πE != nil {
						continue
					}
					πTemp005[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp001, πE = πg.Mod(πF, πTemp002, πTemp004); πE != nil {
						continue
					}
					µidx = πTemp001
					// line 64: pos = fd % (_FD_SETSIZE / len(fds.Bits))
					πF.SetLineno(64)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_FD_SETSIZE); πE != nil {
						continue
					}
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfds, "fds"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µfds, ßBits, nil); πE != nil {
						continue
					}
					πTemp005[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp002, πE = πg.Div(πF, πTemp003, πTemp006); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, µfd, πTemp002); πE != nil {
						continue
					}
					µpos = πTemp001
					// line 65: return bool(fds.Bits[idx] & (1 << pos))
					πF.SetLineno(65)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
						continue
					}
					πTemp002 = µidx
					if πE = πg.CheckLocal(πF, µfds, "fds"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µfds, ßBits, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LShift(πF, πg.NewInt(1).ToObject(), µpos); πE != nil {
						continue
					}
					if πTemp001, πE = πg.And(πF, πTemp003, πTemp002); πE != nil {
						continue
					}
					πTemp005[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßbool); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
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
			if πE = πF.Globals().SetItem(πF, ß_fdset_isset.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 68: def _make_fdset(fd_list):
			πF.SetLineno(68)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "fd_list", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("_make_fdset", "build/src/__python__/select_.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfd_list *πg.Object = πArgs[0]; _ = µfd_list
				var µfds *πg.Object = πg.UnboundLocal; _ = µfds
				var µfd *πg.Object = πg.UnboundLocal; _ = µfd
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 69: fds = _FdSet.new()
					πF.SetLineno(69)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_FdSet); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnew, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µfds = πTemp001
					if πE = πg.CheckLocal(πF, µfd_list, "fd_list"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µfd_list); πE != nil {
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
						µfd = πTemp002
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 71: _fdset_set(fd, fds)
					πF.SetLineno(71)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp005[0] = µfd
					if πE = πg.CheckLocal(πF, µfds, "fds"); πE != nil {
						continue
					}
					πTemp005[1] = µfds
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_fdset_set); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 72: return fds
					πF.SetLineno(72)
					if πE = πg.CheckLocal(πF, µfds, "fds"); πE != nil {
						continue
					}
					πR = µfds
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_make_fdset.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 75: def _normalize_fd_list(fds):
			πF.SetLineno(75)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "fds", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("_normalize_fd_list", "build/src/__python__/select_.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfds *πg.Object = πArgs[0]; _ = µfds
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µfd *πg.Object = πg.UnboundLocal; _ = µfd
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 76: result = []
					πF.SetLineno(76)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µresult = πTemp002
					// line 80: i = 0
					πF.SetLineno(80)
					µi = πg.NewInt(0).ToObject()
					// line 81: while i < len(fds):
					πF.SetLineno(81)
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
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfds, "fds"); πE != nil {
						continue
					}
					πTemp001[0] = µfds
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.LT(πF, µi, πTemp006); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 82: fd = fds[i]
					πF.SetLineno(82)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp002 = µi
					if πE = πg.CheckLocal(πF, µfds, "fds"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µfds, πTemp002); πE != nil {
						continue
					}
					µfd = πTemp005
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp001[0] = µfd
					πTemp001[1] = ßfileno.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					goto Label5
					// line 83: if hasattr(fd, 'fileno'):
					πF.SetLineno(83)
				Label4:
					// line 84: fd = fd.fileno()
					πF.SetLineno(84)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µfd, ßfileno, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µfd = πTemp005
					goto Label5
				Label5:
					// line 85: result.append(fd)
					πF.SetLineno(85)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfd, "fd"); πE != nil {
						continue
					}
					πTemp001[0] = µfd
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 86: i += 1
					πF.SetLineno(86)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µi = πTemp002
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 87: return result
					πF.SetLineno(87)
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
			if πE = πF.Globals().SetItem(πF, ß_normalize_fd_list.ToObject(), πTemp008); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("select_", Code)
}
