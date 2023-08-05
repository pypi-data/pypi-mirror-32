package sys
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/sys.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßArgs := πg.InternStr("Args")
		ßFalse := πg.InternStr("False")
		ßGOOS := πg.InternStr("GOOS")
		ßMaxInt := πg.InternStr("MaxInt")
		ßMaxRune := πg.InternStr("MaxRune")
		ßNone := πg.InternStr("None")
		ßStderr := πg.InternStr("Stderr")
		ßStdin := πg.InternStr("Stdin")
		ßStdout := πg.InternStr("Stdout")
		ßSysModules := πg.InternStr("SysModules")
		ßSystemExit := πg.InternStr("SystemExit")
		ßValueError := πg.InternStr("ValueError")
		ßVersion := πg.InternStr("Version")
		ß_Flags := πg.InternStr("_Flags")
		ß__exc_clear__ := πg.InternStr("__exc_clear__")
		ß__exc_info__ := πg.InternStr("__exc_info__")
		ß__frame__ := πg.InternStr("__frame__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_getframe := πg.InternStr("_getframe")
		ßappend := πg.InternStr("append")
		ßarg := πg.InternStr("arg")
		ßargv := πg.InternStr("argv")
		ßbyteorder := πg.InternStr("byteorder")
		ßbytes_warning := πg.InternStr("bytes_warning")
		ßdebug := πg.InternStr("debug")
		ßdivision_new := πg.InternStr("division_new")
		ßdivision_warning := πg.InternStr("division_warning")
		ßdont_write_bytecode := πg.InternStr("dont_write_bytecode")
		ßexc_clear := πg.InternStr("exc_clear")
		ßexc_info := πg.InternStr("exc_info")
		ßexit := πg.InternStr("exit")
		ßf_back := πg.InternStr("f_back")
		ßflags := πg.InternStr("flags")
		ßgoversion := πg.InternStr("goversion")
		ßhash_randomization := πg.InternStr("hash_randomization")
		ßignore_environment := πg.InternStr("ignore_environment")
		ßinspect := πg.InternStr("inspect")
		ßinteractive := πg.InternStr("interactive")
		ßlittle := πg.InternStr("little")
		ßmaxint := πg.InternStr("maxint")
		ßmaxsize := πg.InternStr("maxsize")
		ßmaxunicode := πg.InternStr("maxunicode")
		ßmodules := πg.InternStr("modules")
		ßno_site := πg.InternStr("no_site")
		ßno_user_site := πg.InternStr("no_user_site")
		ßobject := πg.InternStr("object")
		ßoptimize := πg.InternStr("optimize")
		ßplatform := πg.InternStr("platform")
		ßpy3k_warning := πg.InternStr("py3k_warning")
		ßpy3kwarning := πg.InternStr("py3kwarning")
		ßstderr := πg.InternStr("stderr")
		ßstdin := πg.InternStr("stdin")
		ßstdout := πg.InternStr("stdout")
		ßtabcheck := πg.InternStr("tabcheck")
		ßtype := πg.InternStr("type")
		ßunicode := πg.InternStr("unicode")
		ßverbose := πg.InternStr("verbose")
		ßversion := πg.InternStr("version")
		ßwarnoptions := πg.InternStr("warnoptions")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 bool
		_ = πTemp004
		var πTemp005 bool
		_ = πTemp005
		var πTemp006 *πg.Object
		_ = πTemp006
		var πTemp007 *πg.Dict
		_ = πTemp007
		var πTemp008 []πg.Param
		_ = πTemp008
		var πTemp009 *πg.Object
		_ = πTemp009
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 1: goto Label1
			case 2: goto Label2
			default: panic("unexpected function state")
			}
			// line 15: """System-specific parameters and functions."""
			πF.SetLineno(15)
			// line 17: from '__go__/os' import Args
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßArgs, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßArgs.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 18: from '__go__/grumpy' import SysModules, MaxInt, Stdin as stdin, Stdout as stdout, Stderr as stderr  # pylint: disable=g-multiple-import
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/grumpy"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßSysModules, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSysModules.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßMaxInt, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMaxInt.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßStdin, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßstdin.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßStdout, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßstdout.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßStderr, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßstderr.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 19: from '__go__/runtime' import (GOOS as platform, Version)
			πF.SetLineno(19)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/runtime"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßGOOS, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßplatform.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßVersion, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßVersion.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 20: from '__go__/unicode' import MaxRune
			πF.SetLineno(20)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/unicode"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßMaxRune, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMaxRune.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 22: argv = []
			πF.SetLineno(22)
			πTemp002 = make([]*πg.Object, 0)
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßargv.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßArgs); πE != nil {
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
			if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
				if πE = πF.Globals().SetItem(πF, ßarg.ToObject(), πTemp003); πE != nil {
					continue
				}
			}
			if πE != nil || !πTemp005 {
				continue
			}
			πF.PushCheckpoint(1)            
			// line 24: argv.append(arg)
			πF.SetLineno(24)
			πTemp002 = πF.MakeArgs(1)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßarg); πE != nil {
				continue
			}
			πTemp002[0] = πTemp003
			if πTemp003, πE = πg.ResolveGlobal(πF, ßargv); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			continue
		Label2:
			if πE != nil || πR != nil {
				continue
			}
		Label3:
			// line 26: goversion = Version()
			πF.SetLineno(26)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßVersion); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßgoversion.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 27: maxint = MaxInt
			πF.SetLineno(27)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßMaxInt); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmaxint.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 28: maxsize = maxint
			πF.SetLineno(28)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßmaxint); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmaxsize.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 29: maxunicode = MaxRune
			πF.SetLineno(29)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßMaxRune); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmaxunicode.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 30: modules = SysModules
			πF.SetLineno(30)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßSysModules); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmodules.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 31: py3kwarning = False
			πF.SetLineno(31)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßpy3kwarning.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 32: warnoptions = []
			πF.SetLineno(32)
			πTemp002 = make([]*πg.Object, 0)
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßwarnoptions.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 34: byteorder = 'little'
			πF.SetLineno(34)
			if πE = πF.Globals().SetItem(πF, ßbyteorder.ToObject(), ßlittle.ToObject()); πE != nil {
				continue
			}
			// line 35: version = '2.7.13'
			πF.SetLineno(35)
			if πE = πF.Globals().SetItem(πF, ßversion.ToObject(), πg.NewStr("2.7.13").ToObject()); πE != nil {
				continue
			}
			// line 37: class _Flags(object):
			πF.SetLineno(37)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp006, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp007 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_Flags", "build/src/__python__/sys.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp007
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 38: """Container class for sys.flags."""
					πF.SetLineno(38)
					// line 39: debug = 0
					πF.SetLineno(39)
					if πE = πClass.SetItem(πF, ßdebug.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 40: py3k_warning = 0
					πF.SetLineno(40)
					if πE = πClass.SetItem(πF, ßpy3k_warning.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 41: division_warning = 0
					πF.SetLineno(41)
					if πE = πClass.SetItem(πF, ßdivision_warning.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 42: division_new = 0
					πF.SetLineno(42)
					if πE = πClass.SetItem(πF, ßdivision_new.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 43: inspect = 0
					πF.SetLineno(43)
					if πE = πClass.SetItem(πF, ßinspect.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 44: interactive = 0
					πF.SetLineno(44)
					if πE = πClass.SetItem(πF, ßinteractive.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 45: optimize = 0
					πF.SetLineno(45)
					if πE = πClass.SetItem(πF, ßoptimize.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 46: dont_write_bytecode = 0
					πF.SetLineno(46)
					if πE = πClass.SetItem(πF, ßdont_write_bytecode.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 47: no_user_site = 0
					πF.SetLineno(47)
					if πE = πClass.SetItem(πF, ßno_user_site.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 48: no_site = 0
					πF.SetLineno(48)
					if πE = πClass.SetItem(πF, ßno_site.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 49: ignore_environment = 0
					πF.SetLineno(49)
					if πE = πClass.SetItem(πF, ßignore_environment.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 50: tabcheck = 0
					πF.SetLineno(50)
					if πE = πClass.SetItem(πF, ßtabcheck.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 51: verbose = 0
					πF.SetLineno(51)
					if πE = πClass.SetItem(πF, ßverbose.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 52: unicode = 0
					πF.SetLineno(52)
					if πE = πClass.SetItem(πF, ßunicode.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 53: bytes_warning = 0
					πF.SetLineno(53)
					if πE = πClass.SetItem(πF, ßbytes_warning.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 54: hash_randomization = 0
					πF.SetLineno(54)
					if πE = πClass.SetItem(πF, ßhash_randomization.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp006, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("_Flags").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Flags.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 57: flags = _Flags()
			πF.SetLineno(57)
			if πTemp001, πE = πg.ResolveGlobal(πF, ß_Flags); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßflags.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 60: def exc_clear():
			πF.SetLineno(60)
			πTemp008 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("exc_clear", "build/src/__python__/sys.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 61: __frame__().__exc_clear__()
					πF.SetLineno(61)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß__frame__); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp002, ß__exc_clear__, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßexc_clear.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 64: def exc_info():
			πF.SetLineno(64)
			πTemp008 = make([]πg.Param, 0)
			πTemp003 = πg.NewFunction(πg.NewCode("exc_info", "build/src/__python__/sys.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µe *πg.Object = πg.UnboundLocal; _ = µe
				var µtb *πg.Object = πg.UnboundLocal; _ = µtb
				var µt *πg.Object = πg.UnboundLocal; _ = µt
				var πTemp001 *πg.Object
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
					// line 65: e, tb = __frame__().__exc_info__()  # pylint: disable=undefined-variable
					πF.SetLineno(65)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß__frame__); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp002, ß__exc_info__, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
						continue
					}
					µe = πTemp001
					µtb = πTemp003
					// line 66: t = None
					πF.SetLineno(66)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µt = πTemp001
					if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µe); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 67: if e:
					πF.SetLineno(67)
				Label1:
					// line 68: t = type(e)
					πF.SetLineno(68)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
						continue
					}
					πTemp005[0] = µe
					if πTemp001, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µt = πTemp002
					goto Label2
				Label2:
					// line 69: return t, e, tb
					πF.SetLineno(69)
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple3(µt, µe, µtb).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßexc_info.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 72: def exit(code=None):  # pylint: disable=redefined-builtin
			πF.SetLineno(72)
			πTemp008 = make([]πg.Param, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp008[0] = πg.Param{Name: "code", Def: πTemp009}
			πTemp006 = πg.NewFunction(πg.NewCode("exit", "build/src/__python__/sys.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcode *πg.Object = πArgs[0]; _ = µcode
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
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp001[0] = µcode
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSystemExit); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 73: raise SystemExit(code)
					πF.SetLineno(73)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßexit.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 76: def _getframe(depth=0):
			πF.SetLineno(76)
			πTemp008 = make([]πg.Param, 1)
			πTemp008[0] = πg.Param{Name: "depth", Def: πg.NewInt(0).ToObject()}
			πTemp009 = πg.NewFunction(πg.NewCode("_getframe", "build/src/__python__/sys.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µdepth *πg.Object = πArgs[0]; _ = µdepth
				var µf *πg.Object = πg.UnboundLocal; _ = µf
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 77: f = __frame__()
					πF.SetLineno(77)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß__frame__); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µf = πTemp002
					// line 78: while depth > 0 and f is not None:
					πF.SetLineno(78)
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
					if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GT(πF, µdepth, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp002
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µf != πTemp006).ToObject()
					πTemp001 = πTemp002
				Label4:
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 79: f = f.f_back
					πF.SetLineno(79)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßf_back, nil); πE != nil {
						continue
					}
					µf = πTemp001
					// line 80: depth -= 1
					πF.SetLineno(80)
					if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ISub(πF, µdepth, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µdepth = πTemp001
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µf == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label5
					}
					goto Label6
					// line 81: if f is None:
					πF.SetLineno(81)
				Label5:
					πTemp007 = πF.MakeArgs(1)
					πTemp007[0] = πg.NewStr("call stack is not deep enough").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 82: raise ValueError('call stack is not deep enough')
					πF.SetLineno(82)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label6
				Label6:
					// line 83: return f
					πF.SetLineno(83)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					πR = µf
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_getframe.ToObject(), πTemp009); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("sys", Code)
}
