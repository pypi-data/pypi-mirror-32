package keyword
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/keyword.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßValueError := πg.InternStr("ValueError")
		ß__all__ := πg.InternStr("__all__")
		ß__contains__ := πg.InternStr("__contains__")
		ß__main__ := πg.InternStr("__main__")
		ß__name__ := πg.InternStr("__name__")
		ßand := πg.InternStr("and")
		ßappend := πg.InternStr("append")
		ßargv := πg.InternStr("argv")
		ßas := πg.InternStr("as")
		ßassert := πg.InternStr("assert")
		ßbreak := πg.InternStr("break")
		ßclass := πg.InternStr("class")
		ßclose := πg.InternStr("close")
		ßcompile := πg.InternStr("compile")
		ßcontinue := πg.InternStr("continue")
		ßdef := πg.InternStr("def")
		ßdel := πg.InternStr("del")
		ßelif := πg.InternStr("elif")
		ßelse := πg.InternStr("else")
		ßexcept := πg.InternStr("except")
		ßexec := πg.InternStr("exec")
		ßexit := πg.InternStr("exit")
		ßfinally := πg.InternStr("finally")
		ßfor := πg.InternStr("for")
		ßfrom := πg.InternStr("from")
		ßfrozenset := πg.InternStr("frozenset")
		ßglobal := πg.InternStr("global")
		ßgroup := πg.InternStr("group")
		ßif := πg.InternStr("if")
		ßimport := πg.InternStr("import")
		ßin := πg.InternStr("in")
		ßindex := πg.InternStr("index")
		ßis := πg.InternStr("is")
		ßiskeyword := πg.InternStr("iskeyword")
		ßjoin := πg.InternStr("join")
		ßkwlist := πg.InternStr("kwlist")
		ßlambda := πg.InternStr("lambda")
		ßlen := πg.InternStr("len")
		ßmain := πg.InternStr("main")
		ßnot := πg.InternStr("not")
		ßopen := πg.InternStr("open")
		ßor := πg.InternStr("or")
		ßpass := πg.InternStr("pass")
		ßprint := πg.InternStr("print")
		ßraise := πg.InternStr("raise")
		ßreadlines := πg.InternStr("readlines")
		ßreturn := πg.InternStr("return")
		ßsearch := πg.InternStr("search")
		ßsort := πg.InternStr("sort")
		ßstderr := πg.InternStr("stderr")
		ßtry := πg.InternStr("try")
		ßw := πg.InternStr("w")
		ßwhile := πg.InternStr("while")
		ßwith := πg.InternStr("with")
		ßwrite := πg.InternStr("write")
		ßyield := πg.InternStr("yield")
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
		var πTemp006 bool
		_ = πTemp006
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 3: """Keywords (from "graminit.c")
			πF.SetLineno(3)
			// line 13: __all__ = ["iskeyword", "kwlist"]
			πF.SetLineno(13)
			πTemp001 = make([]*πg.Object, 2)
			πTemp001[0] = ßiskeyword.ToObject()
			πTemp001[1] = ßkwlist.ToObject()
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 15: kwlist = [
			πF.SetLineno(15)
			πTemp001 = make([]*πg.Object, 31)
			πTemp001[0] = ßand.ToObject()
			πTemp001[1] = ßas.ToObject()
			πTemp001[2] = ßassert.ToObject()
			πTemp001[3] = ßbreak.ToObject()
			πTemp001[4] = ßclass.ToObject()
			πTemp001[5] = ßcontinue.ToObject()
			πTemp001[6] = ßdef.ToObject()
			πTemp001[7] = ßdel.ToObject()
			πTemp001[8] = ßelif.ToObject()
			πTemp001[9] = ßelse.ToObject()
			πTemp001[10] = ßexcept.ToObject()
			πTemp001[11] = ßexec.ToObject()
			πTemp001[12] = ßfinally.ToObject()
			πTemp001[13] = ßfor.ToObject()
			πTemp001[14] = ßfrom.ToObject()
			πTemp001[15] = ßglobal.ToObject()
			πTemp001[16] = ßif.ToObject()
			πTemp001[17] = ßimport.ToObject()
			πTemp001[18] = ßin.ToObject()
			πTemp001[19] = ßis.ToObject()
			πTemp001[20] = ßlambda.ToObject()
			πTemp001[21] = ßnot.ToObject()
			πTemp001[22] = ßor.ToObject()
			πTemp001[23] = ßpass.ToObject()
			πTemp001[24] = ßprint.ToObject()
			πTemp001[25] = ßraise.ToObject()
			πTemp001[26] = ßreturn.ToObject()
			πTemp001[27] = ßtry.ToObject()
			πTemp001[28] = ßwhile.ToObject()
			πTemp001[29] = ßwith.ToObject()
			πTemp001[30] = ßyield.ToObject()
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßkwlist.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 51: iskeyword = frozenset(kwlist).__contains__
			πF.SetLineno(51)
			πTemp001 = πF.MakeArgs(1)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßkwlist); πE != nil {
				continue
			}
			πTemp001[0] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßfrozenset); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πTemp002, πE = πg.GetAttr(πF, πTemp003, ß__contains__, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßiskeyword.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 53: def main():
			πF.SetLineno(53)
			πTemp004 = make([]πg.Param, 0)
			πTemp002 = πg.NewFunction(πg.NewCode("main", "build/src/__python__/keyword.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsys *πg.Object = πg.UnboundLocal; _ = µsys
				var µre *πg.Object = πg.UnboundLocal; _ = µre
				var µargs *πg.Object = πg.UnboundLocal; _ = µargs
				var µiptfile *πg.Object = πg.UnboundLocal; _ = µiptfile
				var µoptfile *πg.Object = πg.UnboundLocal; _ = µoptfile
				var µfp *πg.Object = πg.UnboundLocal; _ = µfp
				var µstrprog *πg.Object = πg.UnboundLocal; _ = µstrprog
				var µlines *πg.Object = πg.UnboundLocal; _ = µlines
				var µline *πg.Object = πg.UnboundLocal; _ = µline
				var µmatch *πg.Object = πg.UnboundLocal; _ = µmatch
				var µformat *πg.Object = πg.UnboundLocal; _ = µformat
				var µstart *πg.Object = πg.UnboundLocal; _ = µstart
				var µend *πg.Object = πg.UnboundLocal; _ = µend
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
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.BaseException
				_ = πTemp010
				var πTemp011 *πg.Traceback
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 14: goto Label14
					case 6: goto Label6
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					// line 54: import sys, re
					πF.SetLineno(54)
					if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					µsys = πTemp001
					if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					µre = πTemp001
					// line 56: args = sys.argv[1:]
					πF.SetLineno(56)
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µsys, ßargv, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
						continue
					}
					µargs = πTemp003
					// line 57: iptfile = args and args[0] or "Python/graminit.c"
					πF.SetLineno(57)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp003 = µargs
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label2
					}
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µargs, πTemp004); πE != nil {
						continue
					}
					πTemp003 = πTemp007
				Label2:
					πTemp001 = πTemp003
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					πTemp001 = πg.NewStr("Python/graminit.c").ToObject()
				Label1:
					µiptfile = πTemp001
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp002[0] = µargs
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.GT(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label3
					}
					goto Label4
					// line 58: if len(args) > 1: optfile = args[1]
					πF.SetLineno(58)
				Label3:
					// line 58: if len(args) > 1: optfile = args[1]
					πF.SetLineno(58)
					πTemp001 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
						continue
					}
					µoptfile = πTemp003
					goto Label5
				Label4:
					// line 59: else: optfile = "Lib/keyword.py"
					πF.SetLineno(59)
					µoptfile = πg.NewStr("Lib/keyword.py").ToObject()
					goto Label5
				Label5:
					// line 62: fp = open(iptfile)
					πF.SetLineno(62)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µiptfile, "iptfile"); πE != nil {
						continue
					}
					πTemp002[0] = µiptfile
					if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µfp = πTemp003
					// line 63: strprog = re.compile('"([^"]+)"')
					πF.SetLineno(63)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("\"([^\"]+)\"").ToObject()
					if πE = πg.CheckLocal(πF, µre, "re"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µre, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µstrprog = πTemp003
					// line 64: lines = []
					πF.SetLineno(64)
					πTemp002 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp002...).ToObject()
					µlines = πTemp001
					if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µfp); πE != nil {
						continue
					}
					πF.PushCheckpoint(7)
					πTemp005 = false
				Label6:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp005 {
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
						πTemp006 = !isStop
					} else {
						πTemp006 = true
						µline = πTemp003
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(6)            
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Contains(πF, µline, πg.NewStr("{1, \"").ToObject()); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label9
					}
					goto Label10
					// line 66: if '{1, "' in line:
					πF.SetLineno(66)
				Label9:
					// line 67: match = strprog.search(line)
					πF.SetLineno(67)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp002[0] = µline
					if πE = πg.CheckLocal(πF, µstrprog, "strprog"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µstrprog, ßsearch, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µmatch = πTemp004
					if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µmatch); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label11
					}
					goto Label12
					// line 68: if match:
					πF.SetLineno(68)
				Label11:
					// line 69: lines.append("        '" + match.group(1) + "',\n")
					πF.SetLineno(69)
					πTemp002 = πF.MakeArgs(1)
					πTemp008 = πF.MakeArgs(1)
					πTemp008[0] = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µmatch, ßgroup, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp007.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πTemp004, πE = πg.Add(πF, πg.NewStr("        '").ToObject(), πTemp009); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πTemp004, πg.NewStr("',\n").ToObject()); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µlines, ßappend, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label12
				Label12:
					goto Label10
				Label10:
					continue
				Label7:
					if πE != nil || πR != nil {
						continue
					}
				Label8:
					// line 70: fp.close()
					πF.SetLineno(70)
					if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfp, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 71: lines.sort()
					πF.SetLineno(71)
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µlines, ßsort, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 74: fp = open(optfile)
					πF.SetLineno(74)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µoptfile, "optfile"); πE != nil {
						continue
					}
					πTemp002[0] = µoptfile
					if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µfp = πTemp003
					// line 75: format = fp.readlines()
					πF.SetLineno(75)
					if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfp, ßreadlines, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µformat = πTemp003
					// line 76: fp.close()
					πF.SetLineno(76)
					if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfp, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 79: try:
					πF.SetLineno(79)
					πF.PushCheckpoint(14)
					// line 80: start = format.index("#--start keywords--\n") + 1
					πF.SetLineno(80)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("#--start keywords--\n").ToObject()
					if πE = πg.CheckLocal(πF, µformat, "format"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µformat, ßindex, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Add(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µstart = πTemp001
					// line 81: end = format.index("#--end keywords--\n")
					πF.SetLineno(81)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("#--end keywords--\n").ToObject()
					if πE = πg.CheckLocal(πF, µformat, "format"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µformat, ßindex, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µend = πTemp003
					// line 82: format[start:end] = lines
					πF.SetLineno(82)
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µlines); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µformat, "format"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{µstart, µend, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.SetItem(πF, µformat, πTemp003, πTemp001); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					goto Label13
				Label14:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp010, πTemp011 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsInstance(πF, πTemp010.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label15
					}
					πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
					continue
					// line 83: except ValueError:
					πF.SetLineno(83)
				Label15:
					// line 84: sys.stderr.write("target does not contain format markers\n")
					πF.SetLineno(84)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("target does not contain format markers\n").ToObject()
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µsys, ßstderr, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 85: sys.exit(1)
					πF.SetLineno(85)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µsys, ßexit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πF.RestoreExc(nil, nil)
					goto Label13
				Label13:
					// line 88: fp = open(optfile, 'w')
					πF.SetLineno(88)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µoptfile, "optfile"); πE != nil {
						continue
					}
					πTemp002[0] = µoptfile
					πTemp002[1] = ßw.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µfp = πTemp003
					// line 89: fp.write(''.join(format))
					πF.SetLineno(89)
					πTemp002 = πF.MakeArgs(1)
					πTemp008 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µformat, "format"); πE != nil {
						continue
					}
					πTemp008[0] = µformat
					if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					πTemp002[0] = πTemp003
					if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfp, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 90: fp.close()
					πF.SetLineno(90)
					if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfp, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßmain.ToObject(), πTemp002); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp003, πE = πg.Eq(πF, πTemp005, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label1
			}
			goto Label2
			// line 92: if __name__ == "__main__":
			πF.SetLineno(92)
		Label1:
			// line 93: main()
			πF.SetLineno(93)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßmain); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("keyword", Code)
}
