package uu
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/uu.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßAttributeError := πg.InternStr("AttributeError")
		ßError := πg.InternStr("Error")
		ßException := πg.InternStr("Exception")
		ßFalse := πg.InternStr("False")
		ßNone := πg.InternStr("None")
		ßOptionParser := πg.InternStr("OptionParser")
		ßTrue := πg.InternStr("True")
		ßValueError := πg.InternStr("ValueError")
		ß__all__ := πg.InternStr("__all__")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßa2b_uu := πg.InternStr("a2b_uu")
		ßadd_option := πg.InternStr("add_option")
		ßappend := πg.InternStr("append")
		ßargv := πg.InternStr("argv")
		ßb2a_uu := πg.InternStr("b2a_uu")
		ßbasename := πg.InternStr("basename")
		ßbasestring := πg.InternStr("basestring")
		ßbegin := πg.InternStr("begin")
		ßbinascii := πg.InternStr("binascii")
		ßchmod := πg.InternStr("chmod")
		ßclose := πg.InternStr("close")
		ßdecode := πg.InternStr("decode")
		ßencode := πg.InternStr("encode")
		ßend := πg.InternStr("end")
		ßerror := πg.InternStr("error")
		ßexists := πg.InternStr("exists")
		ßexit := πg.InternStr("exit")
		ßint := πg.InternStr("int")
		ßisinstance := πg.InternStr("isinstance")
		ßlen := πg.InternStr("len")
		ßopen := πg.InternStr("open")
		ßord := πg.InternStr("ord")
		ßos := πg.InternStr("os")
		ßparse_args := πg.InternStr("parse_args")
		ßpath := πg.InternStr("path")
		ßr := πg.InternStr("r")
		ßrb := πg.InternStr("rb")
		ßread := πg.InternStr("read")
		ßreadline := πg.InternStr("readline")
		ßrstrip := πg.InternStr("rstrip")
		ßsplit := πg.InternStr("split")
		ßst_mode := πg.InternStr("st_mode")
		ßstartswith := πg.InternStr("startswith")
		ßstat := πg.InternStr("stat")
		ßstderr := πg.InternStr("stderr")
		ßstdin := πg.InternStr("stdin")
		ßstdout := πg.InternStr("stdout")
		ßstore_true := πg.InternStr("store_true")
		ßstrip := πg.InternStr("strip")
		ßsys := πg.InternStr("sys")
		ßtest := πg.InternStr("test")
		ßtext := πg.InternStr("text")
		ßw := πg.InternStr("w")
		ßwb := πg.InternStr("wb")
		ßwrite := πg.InternStr("write")
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
		var πTemp009 bool
		_ = πTemp009
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 27: """Implementation of the UUencode and UUdecode functions.
			πF.SetLineno(27)
			// line 33: import binascii
			πF.SetLineno(33)
			if πTemp002, πE = πg.ImportModule(πF, "binascii"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßbinascii.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 34: import os
			πF.SetLineno(34)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 35: import sys
			πF.SetLineno(35)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 37: __all__ = ["Error", "encode", "decode"]
			πF.SetLineno(37)
			πTemp002 = make([]*πg.Object, 3)
			πTemp002[0] = ßError.ToObject()
			πTemp002[1] = ßencode.ToObject()
			πTemp002[2] = ßdecode.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 39: class Error(Exception):
			πF.SetLineno(39)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
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
			_, πE = πg.NewCode("Error", "build/src/__python__/uu.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 40: pass
					πF.SetLineno(40)
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("Error").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßError.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 42: def encode(in_file, out_file, name=None, mode=None):
			πF.SetLineno(42)
			πTemp006 = make([]πg.Param, 4)
			πTemp006[0] = πg.Param{Name: "in_file", Def: nil}
			πTemp006[1] = πg.Param{Name: "out_file", Def: nil}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[2] = πg.Param{Name: "name", Def: πTemp004}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[3] = πg.Param{Name: "mode", Def: πTemp004}
			πTemp001 = πg.NewFunction(πg.NewCode("encode", "build/src/__python__/uu.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µin_file *πg.Object = πArgs[0]; _ = µin_file
				var µout_file *πg.Object = πArgs[1]; _ = µout_file
				var µname *πg.Object = πArgs[2]; _ = µname
				var µmode *πg.Object = πArgs[3]; _ = µmode
				var µopened_files *πg.Object = πg.UnboundLocal; _ = µopened_files
				var µdata *πg.Object = πg.UnboundLocal; _ = µdata
				var µf *πg.Object = πg.UnboundLocal; _ = µf
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.BaseException
				_ = πTemp005
				var πTemp006 *πg.Traceback
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 []*πg.Object
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 10: goto Label10
					case 19: goto Label19
					case 20: goto Label20
					case 22: goto Label22
					case 23: goto Label23
					default: panic("unexpected function state")
					}
					// line 43: """Uuencode file"""
					πF.SetLineno(43)
					// line 47: opened_files = []
					πF.SetLineno(47)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µopened_files = πTemp002
					// line 48: try:
					πF.SetLineno(48)
					πF.PushCheckpoint(1)
					if πE = πg.CheckLocal(πF, µin_file, "in_file"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µin_file, πg.NewStr("-").ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label2
					}
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µin_file, "in_file"); πE != nil {
						continue
					}
					πTemp001[0] = µin_file
					if πTemp002, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label3
					}
					goto Label4
					// line 49: if in_file == '-':
					πF.SetLineno(49)
				Label2:
					// line 50: in_file = sys.stdin
					πF.SetLineno(50)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßstdin, nil); πE != nil {
						continue
					}
					µin_file = πTemp004
					goto Label4
					// line 51: elif isinstance(in_file, basestring):
					πF.SetLineno(51)
				Label3:
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µname == πTemp004).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label5
					}
					goto Label6
					// line 52: if name is None:
					πF.SetLineno(52)
				Label5:
					// line 53: name = os.path.basename(in_file)
					πF.SetLineno(53)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µin_file, "in_file"); πE != nil {
						continue
					}
					πTemp001[0] = µin_file
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp004, ßbasename, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µname = πTemp004
					goto Label6
				Label6:
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µmode == πTemp004).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label7
					}
					goto Label8
					// line 54: if mode is None:
					πF.SetLineno(54)
				Label7:
					// line 55: try:
					πF.SetLineno(55)
					πF.PushCheckpoint(10)
					// line 56: mode = os.stat(in_file).st_mode
					πF.SetLineno(56)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µin_file, "in_file"); πE != nil {
						continue
					}
					πTemp001[0] = µin_file
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßstat, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßst_mode, nil); πE != nil {
						continue
					}
					µmode = πTemp004
					πF.PopCheckpoint()
					goto Label9
				Label10:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp005, πTemp006 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label11
					}
					πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
					continue
					// line 57: except AttributeError:
					πF.SetLineno(57)
				Label11:
					// line 58: pass
					πF.SetLineno(58)
					πF.RestoreExc(nil, nil)
					goto Label9
				Label9:
					goto Label8
				Label8:
					// line 59: in_file = open(in_file, 'rb')
					πF.SetLineno(59)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µin_file, "in_file"); πE != nil {
						continue
					}
					πTemp001[0] = µin_file
					πTemp001[1] = ßrb.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µin_file = πTemp004
					// line 60: opened_files.append(in_file)
					πF.SetLineno(60)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µin_file, "in_file"); πE != nil {
						continue
					}
					πTemp001[0] = µin_file
					if πE = πg.CheckLocal(πF, µopened_files, "opened_files"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µopened_files, ßappend, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label4
				Label4:
					if πE = πg.CheckLocal(πF, µout_file, "out_file"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µout_file, πg.NewStr("-").ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label12
					}
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µout_file, "out_file"); πE != nil {
						continue
					}
					πTemp001[0] = µout_file
					if πTemp002, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label13
					}
					goto Label14
					// line 64: if out_file == '-':
					πF.SetLineno(64)
				Label12:
					// line 65: out_file = sys.stdout
					πF.SetLineno(65)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßstdout, nil); πE != nil {
						continue
					}
					µout_file = πTemp004
					goto Label14
					// line 66: elif isinstance(out_file, basestring):
					πF.SetLineno(66)
				Label13:
					// line 67: out_file = open(out_file, 'wb')
					πF.SetLineno(67)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µout_file, "out_file"); πE != nil {
						continue
					}
					πTemp001[0] = µout_file
					πTemp001[1] = ßwb.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µout_file = πTemp004
					// line 68: opened_files.append(out_file)
					πF.SetLineno(68)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µout_file, "out_file"); πE != nil {
						continue
					}
					πTemp001[0] = µout_file
					if πE = πg.CheckLocal(πF, µopened_files, "opened_files"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µopened_files, ßappend, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label14
				Label14:
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µname == πTemp004).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label15
					}
					goto Label16
					// line 72: if name is None:
					πF.SetLineno(72)
				Label15:
					// line 73: name = '-'
					πF.SetLineno(73)
					µname = πg.NewStr("-").ToObject()
					goto Label16
				Label16:
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µmode == πTemp004).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label17
					}
					goto Label18
					// line 74: if mode is None:
					πF.SetLineno(74)
				Label17:
					// line 75: mode = 0666
					πF.SetLineno(75)
					µmode = πg.NewInt(438).ToObject()
					goto Label18
				Label18:
					// line 79: out_file.write('begin %o %s\n' % ((mode&0777),name))
					πF.SetLineno(79)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.And(πF, µmode, πg.NewInt(511).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple2(πTemp007, µname).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("begin %o %s\n").ToObject(), πTemp004); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µout_file, "out_file"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µout_file, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 80: data = in_file.read(45)
					πF.SetLineno(80)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewInt(45).ToObject()
					if πE = πg.CheckLocal(πF, µin_file, "in_file"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µin_file, ßread, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µdata = πTemp004
					// line 81: while len(data) > 0:
					πF.SetLineno(81)
					πF.PushCheckpoint(20)
					πTemp003 = false
				Label19:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp003 {
						πF.PopCheckpoint()
						goto Label21
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					πTemp001[0] = µdata
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.GT(πF, πTemp007, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(19)            
					// line 82: out_file.write(binascii.b2a_uu(data))
					πF.SetLineno(82)
					πTemp001 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					πTemp009[0] = µdata
					if πTemp002, πE = πg.ResolveGlobal(πF, ßbinascii); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßb2a_uu, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µout_file, "out_file"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µout_file, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 83: data = in_file.read(45)
					πF.SetLineno(83)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewInt(45).ToObject()
					if πE = πg.CheckLocal(πF, µin_file, "in_file"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µin_file, ßread, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µdata = πTemp004
					continue
				Label20:
					if πE != nil || πR != nil {
						continue
					}
				Label21:
					// line 84: out_file.write(' \nend\n')
					πF.SetLineno(84)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr(" \nend\n").ToObject()
					if πE = πg.CheckLocal(πF, µout_file, "out_file"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µout_file, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.PopCheckpoint()
					goto Label1
				Label1:
					πTemp005, πTemp006 = πF.RestoreExc(nil, nil)
					if πE = πg.CheckLocal(πF, µopened_files, "opened_files"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µopened_files); πE != nil {
						continue
					}
					πF.PushCheckpoint(23)
					πTemp003 = false
				Label22:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp003 {
						πF.PopCheckpoint()
						goto Label24
					}
					if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
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
						µf = πTemp004
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(22)            
					// line 87: f.close()
					πF.SetLineno(87)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					continue
				Label23:
					if πE != nil || πR != nil {
						continue
					}
				Label24:
					if πTemp005 != nil {
						πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
						continue
					}
					if πR != nil {
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
			if πE = πF.Globals().SetItem(πF, ßencode.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 90: def decode(in_file, out_file=None, mode=None, quiet=0):
			πF.SetLineno(90)
			πTemp006 = make([]πg.Param, 4)
			πTemp006[0] = πg.Param{Name: "in_file", Def: nil}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "out_file", Def: πTemp005}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[2] = πg.Param{Name: "mode", Def: πTemp005}
			πTemp006[3] = πg.Param{Name: "quiet", Def: πg.NewInt(0).ToObject()}
			πTemp004 = πg.NewFunction(πg.NewCode("decode", "build/src/__python__/uu.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µin_file *πg.Object = πArgs[0]; _ = µin_file
				var µout_file *πg.Object = πArgs[1]; _ = µout_file
				var µmode *πg.Object = πArgs[2]; _ = µmode
				var µquiet *πg.Object = πArgs[3]; _ = µquiet
				var µopened_files *πg.Object = πg.UnboundLocal; _ = µopened_files
				var µhdr *πg.Object = πg.UnboundLocal; _ = µhdr
				var µhdrfields *πg.Object = πg.UnboundLocal; _ = µhdrfields
				var µfp *πg.Object = πg.UnboundLocal; _ = µfp
				var µs *πg.Object = πg.UnboundLocal; _ = µs
				var µdata *πg.Object = πg.UnboundLocal; _ = µdata
				var µv *πg.Object = πg.UnboundLocal; _ = µv
				var µnbytes *πg.Object = πg.UnboundLocal; _ = µnbytes
				var µf *πg.Object = πg.UnboundLocal; _ = µf
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.BaseException
				_ = πTemp008
				var πTemp009 *πg.Traceback
				_ = πTemp009
				var πTemp010 bool
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 *πg.Object
				_ = πTemp012
				var πTemp013 *πg.Object
				_ = πTemp013
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 35: goto Label35
					case 4: goto Label4
					case 5: goto Label5
					case 6: goto Label6
					case 41: goto Label41
					case 42: goto Label42
					case 16: goto Label16
					case 28: goto Label28
					case 30: goto Label30
					case 31: goto Label31
					default: panic("unexpected function state")
					}
					// line 91: """Decode uuencoded file"""
					πF.SetLineno(91)
					// line 95: opened_files = []
					πF.SetLineno(95)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µopened_files = πTemp002
					if πE = πg.CheckLocal(πF, µin_file, "in_file"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µin_file, πg.NewStr("-").ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µin_file, "in_file"); πE != nil {
						continue
					}
					πTemp001[0] = µin_file
					if πTemp002, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label2
					}
					goto Label3
					// line 96: if in_file == '-':
					πF.SetLineno(96)
				Label1:
					// line 97: in_file = sys.stdin
					πF.SetLineno(97)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßstdin, nil); πE != nil {
						continue
					}
					µin_file = πTemp004
					goto Label3
					// line 98: elif isinstance(in_file, basestring):
					πF.SetLineno(98)
				Label2:
					// line 99: in_file = open(in_file)
					πF.SetLineno(99)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µin_file, "in_file"); πE != nil {
						continue
					}
					πTemp001[0] = µin_file
					if πTemp002, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µin_file = πTemp004
					// line 100: opened_files.append(in_file)
					πF.SetLineno(100)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µin_file, "in_file"); πE != nil {
						continue
					}
					πTemp001[0] = µin_file
					if πE = πg.CheckLocal(πF, µopened_files, "opened_files"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µopened_files, ßappend, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label3
				Label3:
					// line 101: try:
					πF.SetLineno(101)
					πF.PushCheckpoint(4)
					// line 105: while True:
					πF.SetLineno(105)
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
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(5)            
					// line 106: hdr = in_file.readline()
					πF.SetLineno(106)
					if πE = πg.CheckLocal(πF, µin_file, "in_file"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µin_file, ßreadline, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µhdr = πTemp004
					if πE = πg.CheckLocal(πF, µhdr, "hdr"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µhdr); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label8
					}
					goto Label9
					// line 107: if not hdr:
					πF.SetLineno(107)
				Label8:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("No valid begin line found in input file").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 108: raise Error('No valid begin line found in input file')
					πF.SetLineno(108)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					goto Label9
				Label9:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ßbegin.ToObject()
					if πE = πg.CheckLocal(πF, µhdr, "hdr"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µhdr, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label10
					}
					goto Label11
					// line 109: if not hdr.startswith('begin'):
					πF.SetLineno(109)
				Label10:
					// line 110: continue
					πF.SetLineno(110)
					continue
					goto Label11
				Label11:
					// line 111: hdrfields = hdr.split(' ', 2)
					πF.SetLineno(111)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewStr(" ").ToObject()
					πTemp001[1] = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µhdr, "hdr"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µhdr, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µhdrfields = πTemp004
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µhdrfields, "hdrfields"); πE != nil {
						continue
					}
					πTemp001[0] = µhdrfields
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp004, πE = πg.Eq(πF, πTemp007, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp004
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label12
					}
					πTemp006 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µhdrfields, "hdrfields"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µhdrfields, πTemp006); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, πTemp007, ßbegin.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp004
				Label12:
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label13
					}
					goto Label14
					// line 112: if len(hdrfields) == 3 and hdrfields[0] == 'begin':
					πF.SetLineno(112)
				Label13:
					// line 113: try:
					πF.SetLineno(113)
					πF.PushCheckpoint(16)
					// line 114: int(hdrfields[1], 8)
					πF.SetLineno(114)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µhdrfields, "hdrfields"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µhdrfields, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp001[1] = πg.NewInt(8).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 115: break
					πF.SetLineno(115)
					πTemp003 = true
					continue
					πF.PopCheckpoint()
					goto Label15
				Label16:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label17
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 116: except ValueError:
					πF.SetLineno(116)
				Label17:
					// line 117: pass
					πF.SetLineno(117)
					πF.RestoreExc(nil, nil)
					goto Label15
				Label15:
					goto Label14
				Label14:
					continue
				Label6:
					if πE != nil || πR != nil {
						continue
					}
				Label7:
					if πE = πg.CheckLocal(πF, µout_file, "out_file"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µout_file == πTemp004).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label18
					}
					goto Label19
					// line 118: if out_file is None:
					πF.SetLineno(118)
				Label18:
					// line 119: out_file = hdrfields[2].rstrip()
					πF.SetLineno(119)
					πTemp002 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µhdrfields, "hdrfields"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µhdrfields, πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp004, ßrstrip, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µout_file = πTemp004
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µout_file, "out_file"); πE != nil {
						continue
					}
					πTemp001[0] = µout_file
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp004, ßexists, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label20
					}
					goto Label21
					// line 120: if os.path.exists(out_file):
					πF.SetLineno(120)
				Label20:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µout_file, "out_file"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Cannot overwrite existing file: %s").ToObject(), µout_file); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 121: raise Error('Cannot overwrite existing file: %s' % out_file)
					πF.SetLineno(121)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					goto Label21
				Label21:
					goto Label19
				Label19:
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µmode == πTemp004).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label22
					}
					goto Label23
					// line 122: if mode is None:
					πF.SetLineno(122)
				Label22:
					// line 123: mode = int(hdrfields[1], 8)
					πF.SetLineno(123)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µhdrfields, "hdrfields"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µhdrfields, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					πTemp001[1] = πg.NewInt(8).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µmode = πTemp004
					goto Label23
				Label23:
					if πE = πg.CheckLocal(πF, µout_file, "out_file"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µout_file, πg.NewStr("-").ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label24
					}
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µout_file, "out_file"); πE != nil {
						continue
					}
					πTemp001[0] = µout_file
					if πTemp002, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label25
					}
					goto Label26
					// line 127: if out_file == '-':
					πF.SetLineno(127)
				Label24:
					// line 128: out_file = sys.stdout
					πF.SetLineno(128)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßstdout, nil); πE != nil {
						continue
					}
					µout_file = πTemp004
					goto Label26
					// line 129: elif isinstance(out_file, basestring):
					πF.SetLineno(129)
				Label25:
					// line 130: fp = open(out_file, 'wb')
					πF.SetLineno(130)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µout_file, "out_file"); πE != nil {
						continue
					}
					πTemp001[0] = µout_file
					πTemp001[1] = ßwb.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µfp = πTemp004
					// line 131: try:
					πF.SetLineno(131)
					πF.PushCheckpoint(28)
					// line 132: os.path.chmod(out_file, mode)
					πF.SetLineno(132)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µout_file, "out_file"); πE != nil {
						continue
					}
					πTemp001[0] = µout_file
					if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
						continue
					}
					πTemp001[1] = µmode
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp004, ßchmod, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.PopCheckpoint()
					goto Label27
				Label28:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label29
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 133: except AttributeError:
					πF.SetLineno(133)
				Label29:
					// line 134: pass
					πF.SetLineno(134)
					πF.RestoreExc(nil, nil)
					goto Label27
				Label27:
					// line 135: out_file = fp
					πF.SetLineno(135)
					if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
						continue
					}
					µout_file = µfp
					// line 136: opened_files.append(out_file)
					πF.SetLineno(136)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µout_file, "out_file"); πE != nil {
						continue
					}
					πTemp001[0] = µout_file
					if πE = πg.CheckLocal(πF, µopened_files, "opened_files"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µopened_files, ßappend, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label26
				Label26:
					// line 140: s = in_file.readline()
					πF.SetLineno(140)
					if πE = πg.CheckLocal(πF, µin_file, "in_file"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µin_file, ßreadline, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µs = πTemp004
					// line 141: while s and s.strip() != 'end':
					πF.SetLineno(141)
					πF.PushCheckpoint(31)
					πTemp003 = false
				Label30:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp003 {
						πF.PopCheckpoint()
						goto Label32
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp002 = µs
					if πTemp010, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp010 {
						goto Label33
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µs, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.NE(πF, πTemp007, ßend.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp004
				Label33:
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(30)            
					// line 142: try:
					πF.SetLineno(142)
					πF.PushCheckpoint(35)
					// line 143: data = binascii.a2b_uu(s)
					πF.SetLineno(143)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					if πTemp002, πE = πg.ResolveGlobal(πF, ßbinascii); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßa2b_uu, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µdata = πTemp002
					πF.PopCheckpoint()
					goto Label34
				Label35:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßbinascii); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßError, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp004); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label36
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 144: except binascii.Error, v:
					πF.SetLineno(144)
				Label36:
					µv = πTemp008.ToObject()
					// line 146: nbytes = (((ord(s[0])-32) & 63) * 4 + 5) // 3
					πF.SetLineno(146)
					πTemp001 = πF.MakeArgs(1)
					πTemp012 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetItem(πF, µs, πTemp012); πE != nil {
						continue
					}
					πTemp001[0] = πTemp013
					if πTemp012, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp012.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp011, πE = πg.Sub(πF, πTemp013, πg.NewInt(32).ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.And(πF, πTemp011, πg.NewInt(63).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Mul(πF, πTemp007, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, πTemp006, πg.NewInt(5).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.FloorDiv(πF, πTemp004, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					µnbytes = πTemp002
					// line 147: data = binascii.a2b_uu(s[:nbytes])
					πF.SetLineno(147)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µnbytes, "nbytes"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µnbytes, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µs, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					if πTemp002, πE = πg.ResolveGlobal(πF, ßbinascii); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßa2b_uu, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µdata = πTemp002
					if πE = πg.CheckLocal(πF, µquiet, "quiet"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µquiet); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label37
					}
					goto Label38
					// line 148: if not quiet:
					πF.SetLineno(148)
				Label37:
					// line 149: sys.stderr.write("Warning: %s\n" % v)
					πF.SetLineno(149)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Warning: %s\n").ToObject(), µv); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßstderr, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp004, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label38
				Label38:
					πF.RestoreExc(nil, nil)
					goto Label34
				Label34:
					// line 150: out_file.write(data)
					πF.SetLineno(150)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					πTemp001[0] = µdata
					if πE = πg.CheckLocal(πF, µout_file, "out_file"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µout_file, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 151: s = in_file.readline()
					πF.SetLineno(151)
					if πE = πg.CheckLocal(πF, µin_file, "in_file"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µin_file, ßreadline, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µs = πTemp004
					continue
				Label31:
					if πE != nil || πR != nil {
						continue
					}
				Label32:
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µs); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp003).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label39
					}
					goto Label40
					// line 152: if not s:
					πF.SetLineno(152)
				Label39:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("Truncated input file").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 153: raise Error('Truncated input file')
					πF.SetLineno(153)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					goto Label40
				Label40:
					πF.PopCheckpoint()
					goto Label4
				Label4:
					πTemp008, πTemp009 = πF.RestoreExc(nil, nil)
					if πE = πg.CheckLocal(πF, µopened_files, "opened_files"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µopened_files); πE != nil {
						continue
					}
					πF.PushCheckpoint(42)
					πTemp003 = false
				Label41:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp003 {
						πF.PopCheckpoint()
						goto Label43
					}
					if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
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
						µf = πTemp004
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(41)            
					// line 156: f.close()
					πF.SetLineno(156)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					continue
				Label42:
					if πE != nil || πR != nil {
						continue
					}
				Label43:
					if πTemp008 != nil {
						πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
						continue
					}
					if πR != nil {
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
			if πE = πF.Globals().SetItem(πF, ßdecode.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 158: def test():
			πF.SetLineno(158)
			πTemp006 = make([]πg.Param, 0)
			πTemp005 = πg.NewFunction(πg.NewCode("test", "build/src/__python__/uu.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µoptparse *πg.Object = πg.UnboundLocal; _ = µoptparse
				var µparser *πg.Object = πg.UnboundLocal; _ = µparser
				var µoptions *πg.Object = πg.UnboundLocal; _ = µoptions
				var µargs *πg.Object = πg.UnboundLocal; _ = µargs
				var µinput *πg.Object = πg.UnboundLocal; _ = µinput
				var µoutput *πg.Object = πg.UnboundLocal; _ = µoutput
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 πg.KWArgs
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
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
					default: panic("unexpected function state")
					}
					// line 159: """uuencode/uudecode main program"""
					πF.SetLineno(159)
					// line 161: import optparse
					πF.SetLineno(161)
					if πTemp002, πE = πg.ImportModule(πF, "optparse"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					µoptparse = πTemp001
					// line 162: parser = optparse.OptionParser(usage='usage: %prog [-d] [-t] [input [output]]')
					πF.SetLineno(162)
					πTemp003 = πg.KWArgs{
						{"usage", πg.NewStr("usage: %prog [-d] [-t] [input [output]]").ToObject()},
					}
					if πE = πg.CheckLocal(πF, µoptparse, "optparse"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µoptparse, ßOptionParser, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, nil, πTemp003); πE != nil {
						continue
					}
					µparser = πTemp004
					// line 163: parser.add_option('-d', '--decode', dest='decode', help='Decode (instead of encode)?', default=False, action='store_true')
					πF.SetLineno(163)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewStr("-d").ToObject()
					πTemp002[1] = πg.NewStr("--decode").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp003 = πg.KWArgs{
						{"dest", ßdecode.ToObject()},
						{"help", πg.NewStr("Decode (instead of encode)?").ToObject()},
						{"default", πTemp001},
						{"action", ßstore_true.ToObject()},
					}
					if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µparser, ßadd_option, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp002, πTemp003); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 164: parser.add_option('-t', '--text', dest='text', help='data is text, encoded format unix-compatible text?', default=False, action='store_true')
					πF.SetLineno(164)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewStr("-t").ToObject()
					πTemp002[1] = πg.NewStr("--text").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp003 = πg.KWArgs{
						{"dest", ßtext.ToObject()},
						{"help", πg.NewStr("data is text, encoded format unix-compatible text?").ToObject()},
						{"default", πTemp001},
						{"action", ßstore_true.ToObject()},
					}
					if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µparser, ßadd_option, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp002, πTemp003); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 166: (options, args) = parser.parse_args()
					πF.SetLineno(166)
					if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µparser, ßparse_args, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp005}}}, πTemp004); πE != nil {
						continue
					}
					µoptions = πTemp001
					µargs = πTemp005
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp002[0] = µargs
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.GT(πF, πTemp005, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label1
					}
					goto Label2
					// line 167: if len(args) > 2:
					πF.SetLineno(167)
				Label1:
					// line 168: parser.error('incorrect number of arguments')
					πF.SetLineno(168)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("incorrect number of arguments").ToObject()
					if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µparser, ßerror, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 169: sys.exit(1)
					πF.SetLineno(169)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(1).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßexit, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label2
				Label2:
					// line 171: input = sys.stdin
					πF.SetLineno(171)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßstdin, nil); πE != nil {
						continue
					}
					µinput = πTemp004
					// line 172: output = sys.stdout
					πF.SetLineno(172)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßstdout, nil); πE != nil {
						continue
					}
					µoutput = πTemp004
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp002[0] = µargs
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.GT(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					goto Label4
					// line 173: if len(args) > 0:
					πF.SetLineno(173)
				Label3:
					// line 174: input = args[0]
					πF.SetLineno(174)
					πTemp001 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
						continue
					}
					µinput = πTemp004
					goto Label4
				Label4:
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp002[0] = µargs
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.GT(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label5
					}
					goto Label6
					// line 175: if len(args) > 1:
					πF.SetLineno(175)
				Label5:
					// line 176: output = args[1]
					πF.SetLineno(176)
					πTemp001 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
						continue
					}
					µoutput = πTemp004
					goto Label6
				Label6:
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µoptions, ßdecode, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label7
					}
					goto Label8
					// line 178: if options.decode:
					πF.SetLineno(178)
				Label7:
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µoptions, ßtext, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label10
					}
					goto Label11
					// line 179: if options.text:
					πF.SetLineno(179)
				Label10:
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					πTemp002[0] = µoutput
					if πTemp001, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
						continue
					}
					πTemp002[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label12
					}
					goto Label13
					// line 180: if isinstance(output, basestring):
					πF.SetLineno(180)
				Label12:
					// line 181: output = open(output, 'w')
					πF.SetLineno(181)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					πTemp002[0] = µoutput
					πTemp002[1] = ßw.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µoutput = πTemp004
					goto Label14
				Label13:
					// line 183: print sys.argv[0], ': cannot do -t to stdout'
					πF.SetLineno(183)
					πTemp002 = make([]*πg.Object, 2)
					πTemp001 = πg.NewInt(0).ToObject()
					if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßargv, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp007, πTemp001); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					πTemp002[1] = πg.NewStr(": cannot do -t to stdout").ToObject()
					if πE = πg.Print(πF, πTemp002, true); πE != nil {
						continue
					}
					// line 184: sys.exit(1)
					πF.SetLineno(184)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(1).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßexit, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label14
				Label14:
					goto Label11
				Label11:
					// line 185: decode(input, output)
					πF.SetLineno(185)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µinput, "input"); πE != nil {
						continue
					}
					πTemp002[0] = µinput
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					πTemp002[1] = µoutput
					if πTemp001, πE = πg.ResolveGlobal(πF, ßdecode); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label9
				Label8:
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µoptions, ßtext, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label15
					}
					goto Label16
					// line 187: if options.text:
					πF.SetLineno(187)
				Label15:
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µinput, "input"); πE != nil {
						continue
					}
					πTemp002[0] = µinput
					if πTemp001, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
						continue
					}
					πTemp002[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label17
					}
					goto Label18
					// line 188: if isinstance(input, basestring):
					πF.SetLineno(188)
				Label17:
					// line 189: input = open(input, 'r')
					πF.SetLineno(189)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µinput, "input"); πE != nil {
						continue
					}
					πTemp002[0] = µinput
					πTemp002[1] = ßr.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µinput = πTemp004
					goto Label19
				Label18:
					// line 191: print sys.argv[0], ': cannot do -t from stdin'
					πF.SetLineno(191)
					πTemp002 = make([]*πg.Object, 2)
					πTemp001 = πg.NewInt(0).ToObject()
					if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßargv, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp007, πTemp001); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					πTemp002[1] = πg.NewStr(": cannot do -t from stdin").ToObject()
					if πE = πg.Print(πF, πTemp002, true); πE != nil {
						continue
					}
					// line 192: sys.exit(1)
					πF.SetLineno(192)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(1).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßexit, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label19
				Label19:
					goto Label16
				Label16:
					// line 193: encode(input, output)
					πF.SetLineno(193)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µinput, "input"); πE != nil {
						continue
					}
					πTemp002[0] = µinput
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					πTemp002[1] = µoutput
					if πTemp001, πE = πg.ResolveGlobal(πF, ßencode); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label9
				Label9:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßtest.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp007, πE = πg.Eq(πF, πTemp008, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp009, πE = πg.IsTrue(πF, πTemp007); πE != nil {
				continue
			}
			if πTemp009 {
				goto Label1
			}
			goto Label2
			// line 195: if __name__ == '__main__':
			πF.SetLineno(195)
		Label1:
			// line 196: test()
			πF.SetLineno(196)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßtest); πE != nil {
				continue
			}
			if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("uu", Code)
}
