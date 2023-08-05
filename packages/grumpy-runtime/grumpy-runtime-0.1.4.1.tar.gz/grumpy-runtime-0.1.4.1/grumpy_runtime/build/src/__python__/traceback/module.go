package traceback
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/traceback.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßBaseException := πg.InternStr("BaseException")
		ßException := πg.InternStr("Exception")
		ßNone := πg.InternStr("None")
		ßSyntaxError := πg.InternStr("SyntaxError")
		ßValueError := πg.InternStr("ValueError")
		ßZeroDivisionError := πg.InternStr("ZeroDivisionError")
		ß__all__ := πg.InternStr("__all__")
		ß__name__ := πg.InternStr("__name__")
		ß_format_final_exc_line := πg.InternStr("_format_final_exc_line")
		ß_print := πg.InternStr("_print")
		ß_some_str := πg.InternStr("_some_str")
		ßappend := πg.InternStr("append")
		ßargs := πg.InternStr("args")
		ßascii := πg.InternStr("ascii")
		ßbackslashreplace := πg.InternStr("backslashreplace")
		ßcheckcache := πg.InternStr("checkcache")
		ßco_filename := πg.InternStr("co_filename")
		ßco_name := πg.InternStr("co_name")
		ßencode := πg.InternStr("encode")
		ßexc_info := πg.InternStr("exc_info")
		ßextract_stack := πg.InternStr("extract_stack")
		ßextract_tb := πg.InternStr("extract_tb")
		ßf_back := πg.InternStr("f_back")
		ßf_code := πg.InternStr("f_code")
		ßf_globals := πg.InternStr("f_globals")
		ßf_lineno := πg.InternStr("f_lineno")
		ßformat_exc := πg.InternStr("format_exc")
		ßformat_exception := πg.InternStr("format_exception")
		ßformat_exception_only := πg.InternStr("format_exception_only")
		ßformat_list := πg.InternStr("format_list")
		ßformat_stack := πg.InternStr("format_stack")
		ßformat_tb := πg.InternStr("format_tb")
		ßgetline := πg.InternStr("getline")
		ßhasattr := πg.InternStr("hasattr")
		ßisinstance := πg.InternStr("isinstance")
		ßisspace := πg.InternStr("isspace")
		ßissubclass := πg.InternStr("issubclass")
		ßjoin := πg.InternStr("join")
		ßlast_traceback := πg.InternStr("last_traceback")
		ßlast_type := πg.InternStr("last_type")
		ßlast_value := πg.InternStr("last_value")
		ßlen := πg.InternStr("len")
		ßlinecache := πg.InternStr("linecache")
		ßlstrip := πg.InternStr("lstrip")
		ßmin := πg.InternStr("min")
		ßopen := πg.InternStr("open")
		ßprint_exc := πg.InternStr("print_exc")
		ßprint_exception := πg.InternStr("print_exception")
		ßprint_last := πg.InternStr("print_last")
		ßprint_list := πg.InternStr("print_list")
		ßprint_stack := πg.InternStr("print_stack")
		ßprint_tb := πg.InternStr("print_tb")
		ßreverse := πg.InternStr("reverse")
		ßrstrip := πg.InternStr("rstrip")
		ßstderr := πg.InternStr("stderr")
		ßstr := πg.InternStr("str")
		ßstrip := πg.InternStr("strip")
		ßsys := πg.InternStr("sys")
		ßtb_frame := πg.InternStr("tb_frame")
		ßtb_lineno := πg.InternStr("tb_lineno")
		ßtb_next := πg.InternStr("tb_next")
		ßtracebacklimit := πg.InternStr("tracebacklimit")
		ßtype := πg.InternStr("type")
		ßtypes := πg.InternStr("types")
		ßunicode := πg.InternStr("unicode")
		ßw := πg.InternStr("w")
		ßwrite := πg.InternStr("write")
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
		var πTemp019 *πg.Object
		_ = πTemp019
		var πTemp020 *πg.Object
		_ = πTemp020
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """Extract, format and print information about Python stack traces."""
			πF.SetLineno(1)
			// line 3: import linecache
			πF.SetLineno(3)
			if πTemp002, πE = πg.ImportModule(πF, "linecache"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßlinecache.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 4: import sys
			πF.SetLineno(4)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 5: import types
			πF.SetLineno(5)
			if πTemp002, πE = πg.ImportModule(πF, "types"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtypes.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 7: __all__ = ['extract_stack', 'extract_tb', 'format_exception',
			πF.SetLineno(7)
			πTemp002 = make([]*πg.Object, 14)
			πTemp002[0] = ßextract_stack.ToObject()
			πTemp002[1] = ßextract_tb.ToObject()
			πTemp002[2] = ßformat_exception.ToObject()
			πTemp002[3] = ßformat_exception_only.ToObject()
			πTemp002[4] = ßformat_list.ToObject()
			πTemp002[5] = ßformat_stack.ToObject()
			πTemp002[6] = ßformat_tb.ToObject()
			πTemp002[7] = ßprint_exc.ToObject()
			πTemp002[8] = ßformat_exc.ToObject()
			πTemp002[9] = ßprint_exception.ToObject()
			πTemp002[10] = ßprint_last.ToObject()
			πTemp002[11] = ßprint_stack.ToObject()
			πTemp002[12] = ßprint_tb.ToObject()
			πTemp002[13] = ßtb_lineno.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 12: def _print(file, str='', terminator='\n'):
			πF.SetLineno(12)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "file", Def: nil}
			πTemp003[1] = πg.Param{Name: "str", Def: ß.ToObject()}
			πTemp003[2] = πg.Param{Name: "terminator", Def: πg.NewStr("\n").ToObject()}
			πTemp001 = πg.NewFunction(πg.NewCode("_print", "build/src/__python__/traceback.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfile *πg.Object = πArgs[0]; _ = µfile
				var µstr *πg.Object = πArgs[1]; _ = µstr
				var µterminator *πg.Object = πArgs[2]; _ = µterminator
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
					// line 13: file.write(str+terminator)
					πF.SetLineno(13)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstr, "str"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µterminator, "terminator"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µstr, µterminator); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µfile, ßwrite, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_print.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 16: def print_list(extracted_list, file=None):
			πF.SetLineno(16)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "extracted_list", Def: nil}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[1] = πg.Param{Name: "file", Def: πTemp005}
			πTemp004 = πg.NewFunction(πg.NewCode("print_list", "build/src/__python__/traceback.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µextracted_list *πg.Object = πArgs[0]; _ = µextracted_list
				var µfile *πg.Object = πArgs[1]; _ = µfile
				var µfilename *πg.Object = πg.UnboundLocal; _ = µfilename
				var µlineno *πg.Object = πg.UnboundLocal; _ = µlineno
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µline *πg.Object = πg.UnboundLocal; _ = µline
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
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 []*πg.Object
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 3: goto Label3
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					// line 17: """Print the list of tuples as returned by extract_tb() or
					πF.SetLineno(17)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µfile == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 19: if file is None:
					πF.SetLineno(19)
				Label1:
					// line 20: file = sys.stderr
					πF.SetLineno(20)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstderr, nil); πE != nil {
						continue
					}
					µfile = πTemp002
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µextracted_list, "extracted_list"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µextracted_list); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}}}, πTemp002); πE != nil {
							continue
						}
						µfilename = πTemp005
						µlineno = πTemp006
						µname = πTemp007
						µline = πTemp008
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(3)            
					// line 22: _print(file,
					πF.SetLineno(22)
					πTemp009 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp009[0] = µfile
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple3(µfilename, µlineno, µname).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("  File \"%s\", line %d, in %s").ToObject(), πTemp005); πE != nil {
						continue
					}
					πTemp009[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_print); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µline); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label6
					}
					goto Label7
					// line 24: if line:
					πF.SetLineno(24)
				Label6:
					// line 25: _print(file, '    %s' % line.strip())
					πF.SetLineno(25)
					πTemp009 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp009[0] = µfile
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µline, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("    %s").ToObject(), πTemp006); πE != nil {
						continue
					}
					πTemp009[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_print); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					goto Label7
				Label7:
					continue
				Label4:
					if πE != nil || πR != nil {
						continue
					}
				Label5:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßprint_list.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 27: def format_list(extracted_list):
			πF.SetLineno(27)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "extracted_list", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("format_list", "build/src/__python__/traceback.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µextracted_list *πg.Object = πArgs[0]; _ = µextracted_list
				var µlist *πg.Object = πg.UnboundLocal; _ = µlist
				var µfilename *πg.Object = πg.UnboundLocal; _ = µfilename
				var µlineno *πg.Object = πg.UnboundLocal; _ = µlineno
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µline *πg.Object = πg.UnboundLocal; _ = µline
				var µitem *πg.Object = πg.UnboundLocal; _ = µitem
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
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
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
					// line 28: """Format a list of traceback entry tuples for printing.
					πF.SetLineno(28)
					// line 37: list = []
					πF.SetLineno(37)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µlist = πTemp002
					if πE = πg.CheckLocal(πF, µextracted_list, "extracted_list"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µextracted_list); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}}}, πTemp005); πE != nil {
							continue
						}
						µfilename = πTemp006
						µlineno = πTemp007
						µname = πTemp008
						µline = πTemp009
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 39: item = '  File "%s", line %d, in %s\n' % (filename,lineno,name)
					πF.SetLineno(39)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp006 = πg.NewTuple3(µfilename, µlineno, µname).ToObject()
					if πTemp005, πE = πg.Mod(πF, πg.NewStr("  File \"%s\", line %d, in %s\n").ToObject(), πTemp006); πE != nil {
						continue
					}
					µitem = πTemp005
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µline); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					goto Label5
					// line 40: if line:
					πF.SetLineno(40)
				Label4:
					// line 41: item = item + '    %s\n' % line.strip()
					πF.SetLineno(41)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µline, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Mod(πF, πg.NewStr("    %s\n").ToObject(), πTemp008); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Add(πF, µitem, πTemp006); πE != nil {
						continue
					}
					µitem = πTemp005
					goto Label5
				Label5:
					// line 42: list.append(item)
					πF.SetLineno(42)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp001[0] = µitem
					if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µlist, ßappend, nil); πE != nil {
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
					// line 43: return list
					πF.SetLineno(43)
					if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
						continue
					}
					πR = µlist
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßformat_list.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 46: def print_tb(tb, limit=None, file=None):
			πF.SetLineno(46)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "tb", Def: nil}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[1] = πg.Param{Name: "limit", Def: πTemp007}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[2] = πg.Param{Name: "file", Def: πTemp007}
			πTemp006 = πg.NewFunction(πg.NewCode("print_tb", "build/src/__python__/traceback.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtb *πg.Object = πArgs[0]; _ = µtb
				var µlimit *πg.Object = πArgs[1]; _ = µlimit
				var µfile *πg.Object = πArgs[2]; _ = µfile
				var µn *πg.Object = πg.UnboundLocal; _ = µn
				var µf *πg.Object = πg.UnboundLocal; _ = µf
				var µlineno *πg.Object = πg.UnboundLocal; _ = µlineno
				var µco *πg.Object = πg.UnboundLocal; _ = µco
				var µfilename *πg.Object = πg.UnboundLocal; _ = µfilename
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µline *πg.Object = πg.UnboundLocal; _ = µline
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 8: goto Label8
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					// line 47: """Print up to 'limit' stack trace entries from the traceback 'tb'.
					πF.SetLineno(47)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µfile == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 54: if file is None:
					πF.SetLineno(54)
				Label1:
					// line 55: file = sys.stderr
					πF.SetLineno(55)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstderr, nil); πE != nil {
						continue
					}
					µfile = πTemp002
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µlimit == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label3
					}
					goto Label4
					// line 56: if limit is None:
					πF.SetLineno(56)
				Label3:
					πTemp004 = πF.MakeArgs(2)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					πTemp004[1] = ßtracebacklimit.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label5
					}
					goto Label6
					// line 57: if hasattr(sys, 'tracebacklimit'):
					πF.SetLineno(57)
				Label5:
					// line 58: limit = sys.tracebacklimit
					πF.SetLineno(58)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtracebacklimit, nil); πE != nil {
						continue
					}
					µlimit = πTemp002
					goto Label6
				Label6:
					goto Label4
				Label4:
					// line 59: n = 0
					πF.SetLineno(59)
					µn = πg.NewInt(0).ToObject()
					// line 60: while tb is not None and (limit is None or n < limit):
					πF.SetLineno(60)
					πF.PushCheckpoint(8)
					πTemp003 = false
				Label7:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp003 {
						πF.PopCheckpoint()
						goto Label9
					}
					if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µtb != πTemp007).ToObject()
					πTemp001 = πTemp002
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label10
					}
					if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp007 = πg.GetBool(µlimit == πTemp009).ToObject()
					πTemp002 = πTemp007
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label11
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.LT(πF, µn, µlimit); πE != nil {
						continue
					}
					πTemp002 = πTemp007
				Label11:
					πTemp001 = πTemp002
				Label10:
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(7)            
					// line 61: f = tb.tb_frame
					πF.SetLineno(61)
					if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µtb, ßtb_frame, nil); πE != nil {
						continue
					}
					µf = πTemp001
					// line 62: lineno = tb.tb_lineno
					πF.SetLineno(62)
					if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µtb, ßtb_lineno, nil); πE != nil {
						continue
					}
					µlineno = πTemp001
					// line 63: co = f.f_code
					πF.SetLineno(63)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßf_code, nil); πE != nil {
						continue
					}
					µco = πTemp001
					// line 64: filename = co.co_filename
					πF.SetLineno(64)
					if πE = πg.CheckLocal(πF, µco, "co"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µco, ßco_filename, nil); πE != nil {
						continue
					}
					µfilename = πTemp001
					// line 65: name = co.co_name
					πF.SetLineno(65)
					if πE = πg.CheckLocal(πF, µco, "co"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µco, ßco_name, nil); πE != nil {
						continue
					}
					µname = πTemp001
					// line 66: _print(file,
					πF.SetLineno(66)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp004[0] = µfile
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(µfilename, µlineno, µname).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("  File \"%s\", line %d, in %s").ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp004[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_print); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 68: linecache.checkcache(filename)
					πF.SetLineno(68)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp004[0] = µfilename
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlinecache); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßcheckcache, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 69: line = linecache.getline(filename, lineno, f.f_globals)
					πF.SetLineno(69)
					πTemp004 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp004[0] = µfilename
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					πTemp004[1] = µlineno
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßf_globals, nil); πE != nil {
						continue
					}
					πTemp004[2] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlinecache); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßgetline, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µline = πTemp001
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µline); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label12
					}
					goto Label13
					// line 70: if line: _print(file, '    ' + line.strip())
					πF.SetLineno(70)
				Label12:
					// line 70: if line: _print(file, '    ' + line.strip())
					πF.SetLineno(70)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp004[0] = µfile
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µline, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πg.NewStr("    ").ToObject(), πTemp007); πE != nil {
						continue
					}
					πTemp004[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_print); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					goto Label13
				Label13:
					// line 71: tb = tb.tb_next
					πF.SetLineno(71)
					if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µtb, ßtb_next, nil); πE != nil {
						continue
					}
					µtb = πTemp001
					// line 72: n = n+1
					πF.SetLineno(72)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µn = πTemp001
					continue
				Label8:
					if πE != nil || πR != nil {
						continue
					}
				Label9:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßprint_tb.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 74: def format_tb(tb, limit = None):
			πF.SetLineno(74)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "tb", Def: nil}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[1] = πg.Param{Name: "limit", Def: πTemp008}
			πTemp007 = πg.NewFunction(πg.NewCode("format_tb", "build/src/__python__/traceback.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtb *πg.Object = πArgs[0]; _ = µtb
				var µlimit *πg.Object = πArgs[1]; _ = µlimit
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
					// line 75: """A shorthand for 'format_list(extract_tb(tb, limit))'."""
					πF.SetLineno(75)
					// line 76: return format_list(extract_tb(tb, limit))
					πF.SetLineno(76)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
						continue
					}
					πTemp002[0] = µtb
					if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
						continue
					}
					πTemp002[1] = µlimit
					if πTemp003, πE = πg.ResolveGlobal(πF, ßextract_tb); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßformat_list); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßformat_tb.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 78: def extract_tb(tb, limit = None):
			πF.SetLineno(78)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "tb", Def: nil}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[1] = πg.Param{Name: "limit", Def: πTemp009}
			πTemp008 = πg.NewFunction(πg.NewCode("extract_tb", "build/src/__python__/traceback.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtb *πg.Object = πArgs[0]; _ = µtb
				var µlimit *πg.Object = πArgs[1]; _ = µlimit
				var µlist *πg.Object = πg.UnboundLocal; _ = µlist
				var µn *πg.Object = πg.UnboundLocal; _ = µn
				var µf *πg.Object = πg.UnboundLocal; _ = µf
				var µlineno *πg.Object = πg.UnboundLocal; _ = µlineno
				var µco *πg.Object = πg.UnboundLocal; _ = µco
				var µfilename *πg.Object = πg.UnboundLocal; _ = µfilename
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µline *πg.Object = πg.UnboundLocal; _ = µline
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 5: goto Label5
					case 6: goto Label6
					default: panic("unexpected function state")
					}
					// line 79: """Return list of up to limit pre-processed entries from traceback.
					πF.SetLineno(79)
					if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µlimit == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 89: if limit is None:
					πF.SetLineno(89)
				Label1:
					πTemp004 = πF.MakeArgs(2)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					πTemp004[1] = ßtracebacklimit.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label3
					}
					goto Label4
					// line 90: if hasattr(sys, 'tracebacklimit'):
					πF.SetLineno(90)
				Label3:
					// line 91: limit = sys.tracebacklimit
					πF.SetLineno(91)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtracebacklimit, nil); πE != nil {
						continue
					}
					µlimit = πTemp002
					goto Label4
				Label4:
					goto Label2
				Label2:
					// line 92: list = []
					πF.SetLineno(92)
					πTemp004 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp004...).ToObject()
					µlist = πTemp001
					// line 93: n = 0
					πF.SetLineno(93)
					µn = πg.NewInt(0).ToObject()
					// line 94: while tb is not None and (limit is None or n < limit):
					πF.SetLineno(94)
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
					if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µtb != πTemp007).ToObject()
					πTemp001 = πTemp002
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label8
					}
					if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp007 = πg.GetBool(µlimit == πTemp009).ToObject()
					πTemp002 = πTemp007
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label9
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.LT(πF, µn, µlimit); πE != nil {
						continue
					}
					πTemp002 = πTemp007
				Label9:
					πTemp001 = πTemp002
				Label8:
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(5)            
					// line 95: f = tb.tb_frame
					πF.SetLineno(95)
					if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µtb, ßtb_frame, nil); πE != nil {
						continue
					}
					µf = πTemp001
					// line 96: lineno = tb.tb_lineno
					πF.SetLineno(96)
					if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µtb, ßtb_lineno, nil); πE != nil {
						continue
					}
					µlineno = πTemp001
					// line 97: co = f.f_code
					πF.SetLineno(97)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßf_code, nil); πE != nil {
						continue
					}
					µco = πTemp001
					// line 98: filename = co.co_filename
					πF.SetLineno(98)
					if πE = πg.CheckLocal(πF, µco, "co"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µco, ßco_filename, nil); πE != nil {
						continue
					}
					µfilename = πTemp001
					// line 99: name = co.co_name
					πF.SetLineno(99)
					if πE = πg.CheckLocal(πF, µco, "co"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µco, ßco_name, nil); πE != nil {
						continue
					}
					µname = πTemp001
					// line 100: linecache.checkcache(filename)
					πF.SetLineno(100)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp004[0] = µfilename
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlinecache); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßcheckcache, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 101: line = linecache.getline(filename, lineno, f.f_globals)
					πF.SetLineno(101)
					πTemp004 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp004[0] = µfilename
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					πTemp004[1] = µlineno
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßf_globals, nil); πE != nil {
						continue
					}
					πTemp004[2] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlinecache); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßgetline, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µline = πTemp001
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µline); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label10
					}
					goto Label11
					// line 102: if line: line = line.strip()
					πF.SetLineno(102)
				Label10:
					// line 102: if line: line = line.strip()
					πF.SetLineno(102)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µline, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µline = πTemp002
					goto Label12
				Label11:
					// line 103: else: line = None
					πF.SetLineno(103)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µline = πTemp001
					goto Label12
				Label12:
					// line 104: list.append((filename, lineno, name, line))
					πF.SetLineno(104)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple4(µfilename, µlineno, µname, µline).ToObject()
					πTemp004[0] = πTemp001
					if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µlist, ßappend, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 105: tb = tb.tb_next
					πF.SetLineno(105)
					if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µtb, ßtb_next, nil); πE != nil {
						continue
					}
					µtb = πTemp001
					// line 106: n = n+1
					πF.SetLineno(106)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µn = πTemp001
					continue
				Label6:
					if πE != nil || πR != nil {
						continue
					}
				Label7:
					// line 107: return list
					πF.SetLineno(107)
					if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
						continue
					}
					πR = µlist
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßextract_tb.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 110: def print_exception(etype, value, tb, limit=None, file=None):
			πF.SetLineno(110)
			πTemp003 = make([]πg.Param, 5)
			πTemp003[0] = πg.Param{Name: "etype", Def: nil}
			πTemp003[1] = πg.Param{Name: "value", Def: nil}
			πTemp003[2] = πg.Param{Name: "tb", Def: nil}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[3] = πg.Param{Name: "limit", Def: πTemp010}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[4] = πg.Param{Name: "file", Def: πTemp010}
			πTemp009 = πg.NewFunction(πg.NewCode("print_exception", "build/src/__python__/traceback.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µetype *πg.Object = πArgs[0]; _ = µetype
				var µvalue *πg.Object = πArgs[1]; _ = µvalue
				var µtb *πg.Object = πArgs[2]; _ = µtb
				var µlimit *πg.Object = πArgs[3]; _ = µlimit
				var µfile *πg.Object = πArgs[4]; _ = µfile
				var µlines *πg.Object = πg.UnboundLocal; _ = µlines
				var µline *πg.Object = πg.UnboundLocal; _ = µline
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 []*πg.Object
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
					// line 111: """Print exception up to 'limit' stack trace entries from 'tb' to 'file'.
					πF.SetLineno(111)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µfile == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 121: if file is None:
					πF.SetLineno(121)
				Label1:
					// line 123: file = open('/dev/stderr', 'w')
					πF.SetLineno(123)
					πTemp004 = πF.MakeArgs(2)
					πTemp004[0] = πg.NewStr("/dev/stderr").ToObject()
					πTemp004[1] = ßw.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µfile = πTemp002
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µtb); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label3
					}
					goto Label4
					// line 125: if tb:
					πF.SetLineno(125)
				Label3:
					// line 126: _print(file, 'Traceback (most recent call last):')
					πF.SetLineno(126)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp004[0] = µfile
					πTemp004[1] = πg.NewStr("Traceback (most recent call last):").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_print); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 127: print_tb(tb, limit, file)
					πF.SetLineno(127)
					πTemp004 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
						continue
					}
					πTemp004[0] = µtb
					if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
						continue
					}
					πTemp004[1] = µlimit
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp004[2] = µfile
					if πTemp001, πE = πg.ResolveGlobal(πF, ßprint_tb); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					goto Label4
				Label4:
					// line 128: lines = format_exception_only(etype, value)
					πF.SetLineno(128)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µetype, "etype"); πE != nil {
						continue
					}
					πTemp004[0] = µetype
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp004[1] = µvalue
					if πTemp001, πE = πg.ResolveGlobal(πF, ßformat_exception_only); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µlines = πTemp002
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µlines); πE != nil {
						continue
					}
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
						µline = πTemp002
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(5)            
					// line 130: _print(file, line, '')
					πF.SetLineno(130)
					πTemp004 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp004[0] = µfile
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp004[1] = µline
					πTemp004[2] = ß.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_print); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					continue
				Label6:
					if πE != nil || πR != nil {
						continue
					}
				Label7:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßprint_exception.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 132: def format_exception(etype, value, tb, limit = None):
			πF.SetLineno(132)
			πTemp003 = make([]πg.Param, 4)
			πTemp003[0] = πg.Param{Name: "etype", Def: nil}
			πTemp003[1] = πg.Param{Name: "value", Def: nil}
			πTemp003[2] = πg.Param{Name: "tb", Def: nil}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[3] = πg.Param{Name: "limit", Def: πTemp011}
			πTemp010 = πg.NewFunction(πg.NewCode("format_exception", "build/src/__python__/traceback.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µetype *πg.Object = πArgs[0]; _ = µetype
				var µvalue *πg.Object = πArgs[1]; _ = µvalue
				var µtb *πg.Object = πArgs[2]; _ = µtb
				var µlimit *πg.Object = πArgs[3]; _ = µlimit
				var µlist *πg.Object = πg.UnboundLocal; _ = µlist
				var πTemp001 bool
				_ = πTemp001
				var πTemp002 []*πg.Object
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
					// line 133: """Format a stack trace and the exception information.
					πF.SetLineno(133)
					if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µtb); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label1
					}
					goto Label2
					// line 141: if tb:
					πF.SetLineno(141)
				Label1:
					// line 142: list = ['Traceback (most recent call last):\n']
					πF.SetLineno(142)
					πTemp002 = make([]*πg.Object, 1)
					πTemp002[0] = πg.NewStr("Traceback (most recent call last):\n").ToObject()
					πTemp003 = πg.NewList(πTemp002...).ToObject()
					µlist = πTemp003
					// line 143: list = list + format_tb(tb, limit)
					πF.SetLineno(143)
					if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
						continue
					}
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
						continue
					}
					πTemp002[0] = µtb
					if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
						continue
					}
					πTemp002[1] = µlimit
					if πTemp004, πE = πg.ResolveGlobal(πF, ßformat_tb); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.Add(πF, µlist, πTemp005); πE != nil {
						continue
					}
					µlist = πTemp003
					goto Label3
				Label2:
					// line 145: list = []
					πF.SetLineno(145)
					πTemp002 = make([]*πg.Object, 0)
					πTemp003 = πg.NewList(πTemp002...).ToObject()
					µlist = πTemp003
					goto Label3
				Label3:
					// line 146: list = list + format_exception_only(etype, value)
					πF.SetLineno(146)
					if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
						continue
					}
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µetype, "etype"); πE != nil {
						continue
					}
					πTemp002[0] = µetype
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp002[1] = µvalue
					if πTemp004, πE = πg.ResolveGlobal(πF, ßformat_exception_only); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.Add(πF, µlist, πTemp005); πE != nil {
						continue
					}
					µlist = πTemp003
					// line 147: return list
					πF.SetLineno(147)
					if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
						continue
					}
					πR = µlist
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßformat_exception.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 149: def format_exception_only(etype, value):
			πF.SetLineno(149)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "etype", Def: nil}
			πTemp003[1] = πg.Param{Name: "value", Def: nil}
			πTemp011 = πg.NewFunction(πg.NewCode("format_exception_only", "build/src/__python__/traceback.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µetype *πg.Object = πArgs[0]; _ = µetype
				var µvalue *πg.Object = πArgs[1]; _ = µvalue
				var µstype *πg.Object = πg.UnboundLocal; _ = µstype
				var µlines *πg.Object = πg.UnboundLocal; _ = µlines
				var µmsg *πg.Object = πg.UnboundLocal; _ = µmsg
				var µfilename *πg.Object = πg.UnboundLocal; _ = µfilename
				var µlineno *πg.Object = πg.UnboundLocal; _ = µlineno
				var µoffset *πg.Object = πg.UnboundLocal; _ = µoffset
				var µbadline *πg.Object = πg.UnboundLocal; _ = µbadline
				var µcaretspace *πg.Object = πg.UnboundLocal; _ = µcaretspace
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 []πg.Param
				_ = πTemp010
				var πTemp011 *πg.BaseException
				_ = πTemp011
				var πTemp012 *πg.Traceback
				_ = πTemp012
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					// line 150: """Format the exception part of a traceback.
					πF.SetLineno(150)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µetype, "etype"); πE != nil {
						continue
					}
					πTemp003[0] = µetype
					if πTemp004, πE = πg.ResolveGlobal(πF, ßBaseException); πE != nil {
						continue
					}
					πTemp003[1] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001 = πTemp005
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µetype, "etype"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(µetype == πTemp005).ToObject()
					πTemp001 = πTemp004
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µetype, "etype"); πE != nil {
						continue
					}
					πTemp003[0] = µetype
					if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(πTemp006 == πTemp005).ToObject()
					πTemp001 = πTemp004
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 172: if (isinstance(etype, BaseException) or
					πF.SetLineno(172)
				Label2:
					// line 175: return [_format_final_exc_line(etype, value)]
					πF.SetLineno(175)
					πTemp003 = make([]*πg.Object, 1)
					πTemp007 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µetype, "etype"); πE != nil {
						continue
					}
					πTemp007[0] = µetype
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp007[1] = µvalue
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_format_final_exc_line); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp003[0] = πTemp004
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					πR = πTemp001
					continue
					goto Label3
				Label3:
					// line 177: stype = etype.__name__
					πF.SetLineno(177)
					if πE = πg.CheckLocal(πF, µetype, "etype"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µetype, ß__name__, nil); πE != nil {
						continue
					}
					µstype = πTemp001
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µetype, "etype"); πE != nil {
						continue
					}
					πTemp003[0] = µetype
					if πTemp004, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
						continue
					}
					πTemp003[1] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					goto Label5
					// line 179: if not issubclass(etype, SyntaxError):
					πF.SetLineno(179)
				Label4:
					// line 180: return [_format_final_exc_line(stype, value)]
					πF.SetLineno(180)
					πTemp003 = make([]*πg.Object, 1)
					πTemp007 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µstype, "stype"); πE != nil {
						continue
					}
					πTemp007[0] = µstype
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp007[1] = µvalue
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_format_final_exc_line); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp003[0] = πTemp004
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					πR = πTemp001
					continue
					goto Label5
				Label5:
					// line 183: lines = []
					πF.SetLineno(183)
					πTemp003 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					µlines = πTemp001
					// line 184: try:
					πF.SetLineno(184)
					πF.PushCheckpoint(7)
					// line 185: msg, (filename, lineno, offset, badline) = value.args
					πF.SetLineno(185)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µvalue, ßargs, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}}}}}, πTemp001); πE != nil {
						continue
					}
					µmsg = πTemp004
					µfilename = πTemp005
					µlineno = πTemp006
					µoffset = πTemp008
					µbadline = πTemp009
					πF.PopCheckpoint()
					// line 189: filename = filename or "<string>"
					πF.SetLineno(189)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp001 = µfilename
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label8
					}
					πTemp001 = πg.NewStr("<string>").ToObject()
				Label8:
					µfilename = πTemp001
					// line 190: lines.append('  File "%s", line %d\n' % (filename, lineno))
					πF.SetLineno(190)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple2(µfilename, µlineno).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("  File \"%s\", line %d\n").ToObject(), πTemp004); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µlines, ßappend, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.CheckLocal(πF, µbadline, "badline"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µbadline != πTemp004).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label9
					}
					goto Label10
					// line 191: if badline is not None:
					πF.SetLineno(191)
				Label9:
					// line 192: lines.append('    %s\n' % badline.strip())
					πF.SetLineno(192)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µbadline, "badline"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µbadline, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("    %s\n").ToObject(), πTemp005); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µlines, ßappend, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µoffset != πTemp004).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label11
					}
					goto Label12
					// line 193: if offset is not None:
					πF.SetLineno(193)
				Label11:
					// line 194: caretspace = badline.rstrip('\n')
					πF.SetLineno(194)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("\n").ToObject()
					if πE = πg.CheckLocal(πF, µbadline, "badline"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µbadline, ßrstrip, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µcaretspace = πTemp004
					// line 195: offset = min(len(caretspace), offset) - 1
					πF.SetLineno(195)
					πTemp003 = πF.MakeArgs(2)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcaretspace, "caretspace"); πE != nil {
						continue
					}
					πTemp007[0] = µcaretspace
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp003[0] = πTemp005
					if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
						continue
					}
					πTemp003[1] = µoffset
					if πTemp004, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp001, πE = πg.Sub(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µoffset = πTemp001
					// line 196: caretspace = caretspace[:offset].lstrip()
					πF.SetLineno(196)
					if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µoffset, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcaretspace, "caretspace"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µcaretspace, πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp004, ßlstrip, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µcaretspace = πTemp004
					// line 198: caretspace = ((c.isspace() and c or ' ') for c in caretspace)
					πF.SetLineno(198)
					πTemp010 = make([]πg.Param, 0)
					πTemp001 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/traceback.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µc *πg.Object = πg.UnboundLocal; _ = µc
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
						var πTemp006 bool
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
								case 1: goto Label1
								case 2: goto Label2
								case 6: goto Label6
								default: panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µcaretspace, "caretspace"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µcaretspace); πE != nil {
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
									µc = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 198: caretspace = ((c.isspace() and c or ' ') for c in caretspace)
								πF.SetLineno(198)
								if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, µc, ßisspace, nil); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
									continue
								}
								πTemp005 = πTemp008
								if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								if !πTemp006 {
									goto Label5
								}
								if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
									continue
								}
								πTemp005 = µc
							Label5:
								πTemp004 = πTemp005
								if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label4
								}
								πTemp004 = πg.NewStr(" ").ToObject()
							Label4:
								πF.PushCheckpoint(6)
								return πTemp004, nil
							Label6:
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
					if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µcaretspace = πTemp004
					// line 199: lines.append('    %s^\n' % ''.join(caretspace))
					πF.SetLineno(199)
					πTemp003 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcaretspace, "caretspace"); πE != nil {
						continue
					}
					πTemp007[0] = µcaretspace
					if πTemp005, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πTemp004, πE = πg.Mod(πF, πg.NewStr("    %s^\n").ToObject(), πTemp006); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µlines, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label12
				Label12:
					goto Label10
				Label10:
					// line 200: value = msg
					πF.SetLineno(200)
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					µvalue = µmsg
					goto Label6
				Label7:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp011, πTemp012 = πF.ExcInfo()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp004); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label13
					}
					πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
					continue
					// line 186: except Exception:
					πF.SetLineno(186)
				Label13:
					// line 187: pass
					πF.SetLineno(187)
					πF.RestoreExc(nil, nil)
					goto Label6
				Label6:
					// line 202: lines.append(_format_final_exc_line(stype, value))
					πF.SetLineno(202)
					πTemp003 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µstype, "stype"); πE != nil {
						continue
					}
					πTemp007[0] = µstype
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp007[1] = µvalue
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_format_final_exc_line); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp003[0] = πTemp005
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µlines, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 203: return lines
					πF.SetLineno(203)
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					πR = µlines
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßformat_exception_only.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 205: def _format_final_exc_line(etype, value):
			πF.SetLineno(205)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "etype", Def: nil}
			πTemp003[1] = πg.Param{Name: "value", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("_format_final_exc_line", "build/src/__python__/traceback.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µetype *πg.Object = πArgs[0]; _ = µetype
				var µvalue *πg.Object = πArgs[1]; _ = µvalue
				var µvaluestr *πg.Object = πg.UnboundLocal; _ = µvaluestr
				var µline *πg.Object = πg.UnboundLocal; _ = µline
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 206: """Return a list of a single line -- normal case for format_exception_only"""
					πF.SetLineno(206)
					// line 207: valuestr = _some_str(value)
					πF.SetLineno(207)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[0] = µvalue
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_some_str); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µvaluestr = πTemp003
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µvalue == πTemp005).ToObject()
					πTemp002 = πTemp003
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µvaluestr, "valuestr"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µvaluestr); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp006).ToObject()
					πTemp002 = πTemp003
				Label1:
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label2
					}
					goto Label3
					// line 208: if value is None or not valuestr:
					πF.SetLineno(208)
				Label2:
					// line 209: line = "%s\n" % etype
					πF.SetLineno(209)
					if πE = πg.CheckLocal(πF, µetype, "etype"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s\n").ToObject(), µetype); πE != nil {
						continue
					}
					µline = πTemp002
					goto Label4
				Label3:
					// line 211: line = "%s: %s\n" % (etype, valuestr)
					πF.SetLineno(211)
					if πE = πg.CheckLocal(πF, µetype, "etype"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µvaluestr, "valuestr"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µetype, µvaluestr).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s: %s\n").ToObject(), πTemp003); πE != nil {
						continue
					}
					µline = πTemp002
					goto Label4
				Label4:
					// line 212: return line
					πF.SetLineno(212)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πR = µline
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_format_final_exc_line.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 214: def _some_str(value):
			πF.SetLineno(214)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "value", Def: nil}
			πTemp013 = πg.NewFunction(πg.NewCode("_some_str", "build/src/__python__/traceback.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µvalue *πg.Object = πArgs[0]; _ = µvalue
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.BaseException
				_ = πTemp004
				var πTemp005 *πg.Traceback
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
					case 2: goto Label2
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 215: try:
					πF.SetLineno(215)
					πF.PushCheckpoint(2)
					// line 216: return str(value)
					πF.SetLineno(216)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[0] = µvalue
					if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 217: except Exception:
					πF.SetLineno(217)
				Label3:
					// line 218: pass
					πF.SetLineno(218)
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 219: try:
					πF.SetLineno(219)
					πF.PushCheckpoint(5)
					// line 220: value = unicode(value)
					πF.SetLineno(220)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[0] = µvalue
					if πTemp002, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µvalue = πTemp003
					// line 221: return value.encode("ascii", "backslashreplace")
					πF.SetLineno(221)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = ßascii.ToObject()
					πTemp001[1] = ßbackslashreplace.ToObject()
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µvalue, ßencode, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
					πF.PopCheckpoint()
					goto Label4
				Label5:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label6
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 222: except Exception:
					πF.SetLineno(222)
				Label6:
					// line 223: pass
					πF.SetLineno(223)
					πF.RestoreExc(nil, nil)
					goto Label4
				Label4:
					// line 224: return '<unprintable %s object>' % type(value).__name__
					πF.SetLineno(224)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[0] = µvalue
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.GetAttr(πF, πTemp007, ß__name__, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("<unprintable %s object>").ToObject(), πTemp003); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ß_some_str.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 227: def print_exc(limit=None, file=None):
			πF.SetLineno(227)
			πTemp003 = make([]πg.Param, 2)
			if πTemp015, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[0] = πg.Param{Name: "limit", Def: πTemp015}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[1] = πg.Param{Name: "file", Def: πTemp015}
			πTemp014 = πg.NewFunction(πg.NewCode("print_exc", "build/src/__python__/traceback.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µlimit *πg.Object = πArgs[0]; _ = µlimit
				var µfile *πg.Object = πArgs[1]; _ = µfile
				var µetype *πg.Object = πg.UnboundLocal; _ = µetype
				var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
				var µtb *πg.Object = πg.UnboundLocal; _ = µtb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.BaseException
				_ = πTemp007
				var πTemp008 *πg.Traceback
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 3: goto Label3
					default: panic("unexpected function state")
					}
					// line 228: """Shorthand for 'print_exception(sys.exc_type, sys.exc_value, sys.exc_traceback, limit, file)'.
					πF.SetLineno(228)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µfile == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 231: if file is None:
					πF.SetLineno(231)
				Label1:
					// line 233: file = open('/dev/stderr', 'w')
					πF.SetLineno(233)
					πTemp004 = πF.MakeArgs(2)
					πTemp004[0] = πg.NewStr("/dev/stderr").ToObject()
					πTemp004[1] = ßw.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µfile = πTemp002
					goto Label2
				Label2:
					// line 235: try:
					πF.SetLineno(235)
					πF.PushCheckpoint(3)
					// line 236: etype, value, tb = sys.exc_info()
					πF.SetLineno(236)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßexc_info, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp001); πE != nil {
						continue
					}
					µetype = πTemp002
					µvalue = πTemp005
					µtb = πTemp006
					// line 237: print_exception(etype, value, tb, limit, file)
					πF.SetLineno(237)
					πTemp004 = πF.MakeArgs(5)
					if πE = πg.CheckLocal(πF, µetype, "etype"); πE != nil {
						continue
					}
					πTemp004[0] = µetype
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp004[1] = µvalue
					if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
						continue
					}
					πTemp004[2] = µtb
					if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
						continue
					}
					πTemp004[3] = µlimit
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp004[4] = µfile
					if πTemp001, πE = πg.ResolveGlobal(πF, ßprint_exception); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πF.PopCheckpoint()
					goto Label3
				Label3:
					πTemp007, πTemp008 = πF.RestoreExc(nil, nil)
					// line 239: etype = value = tb = None
					πF.SetLineno(239)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µetype = πTemp001
					µvalue = πTemp001
					µtb = πTemp001
					if πTemp007 != nil {
						πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
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
			if πE = πF.Globals().SetItem(πF, ßprint_exc.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 242: def format_exc(limit=None):
			πF.SetLineno(242)
			πTemp003 = make([]πg.Param, 1)
			if πTemp016, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[0] = πg.Param{Name: "limit", Def: πTemp016}
			πTemp015 = πg.NewFunction(πg.NewCode("format_exc", "build/src/__python__/traceback.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µlimit *πg.Object = πArgs[0]; _ = µlimit
				var µetype *πg.Object = πg.UnboundLocal; _ = µetype
				var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
				var µtb *πg.Object = πg.UnboundLocal; _ = µtb
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
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 *πg.BaseException
				_ = πTemp007
				var πTemp008 *πg.Traceback
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					default: panic("unexpected function state")
					}
					// line 243: """Like print_exc() but return a string."""
					πF.SetLineno(243)
					// line 244: try:
					πF.SetLineno(244)
					πF.PushCheckpoint(1)
					// line 245: etype, value, tb = sys.exc_info()
					πF.SetLineno(245)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßexc_info, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
						continue
					}
					µetype = πTemp002
					µvalue = πTemp003
					µtb = πTemp004
					// line 246: return ''.join(format_exception(etype, value, tb, limit))
					πF.SetLineno(246)
					πTemp005 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µetype, "etype"); πE != nil {
						continue
					}
					πTemp006[0] = µetype
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp006[1] = µvalue
					if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
						continue
					}
					πTemp006[2] = µtb
					if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
						continue
					}
					πTemp006[3] = µlimit
					if πTemp001, πE = πg.ResolveGlobal(πF, ßformat_exception); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp005[0] = πTemp002
					if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πR = πTemp002
					continue
					πF.PopCheckpoint()
					goto Label1
				Label1:
					πTemp007, πTemp008 = πF.RestoreExc(nil, nil)
					// line 248: etype = value = tb = None
					πF.SetLineno(248)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µetype = πTemp001
					µvalue = πTemp001
					µtb = πTemp001
					if πTemp007 != nil {
						πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
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
			if πE = πF.Globals().SetItem(πF, ßformat_exc.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 251: def print_last(limit=None, file=None):
			πF.SetLineno(251)
			πTemp003 = make([]πg.Param, 2)
			if πTemp017, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[0] = πg.Param{Name: "limit", Def: πTemp017}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[1] = πg.Param{Name: "file", Def: πTemp017}
			πTemp016 = πg.NewFunction(πg.NewCode("print_last", "build/src/__python__/traceback.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µlimit *πg.Object = πArgs[0]; _ = µlimit
				var µfile *πg.Object = πArgs[1]; _ = µfile
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
					// line 252: """This is a shorthand for 'print_exception(sys.last_type,
					πF.SetLineno(252)
					πTemp002 = πF.MakeArgs(2)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					πTemp002[1] = ßlast_type.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
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
					// line 254: if not hasattr(sys, "last_type"):
					πF.SetLineno(254)
				Label1:
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("no last exception").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 255: raise ValueError("no last exception")
					πF.SetLineno(255)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µfile == πTemp003).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label3
					}
					goto Label4
					// line 256: if file is None:
					πF.SetLineno(256)
				Label3:
					// line 257: file = sys.stderr
					πF.SetLineno(257)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßstderr, nil); πE != nil {
						continue
					}
					µfile = πTemp003
					goto Label4
				Label4:
					// line 258: print_exception(sys.last_type, sys.last_value, sys.last_traceback,
					πF.SetLineno(258)
					πTemp002 = πF.MakeArgs(5)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßlast_type, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßlast_value, nil); πE != nil {
						continue
					}
					πTemp002[1] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßlast_traceback, nil); πE != nil {
						continue
					}
					πTemp002[2] = πTemp003
					if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
						continue
					}
					πTemp002[3] = µlimit
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp002[4] = µfile
					if πTemp001, πE = πg.ResolveGlobal(πF, ßprint_exception); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßprint_last.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 262: def print_stack(f=None, limit=None, file=None):
			πF.SetLineno(262)
			πTemp003 = make([]πg.Param, 3)
			if πTemp018, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[0] = πg.Param{Name: "f", Def: πTemp018}
			if πTemp018, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[1] = πg.Param{Name: "limit", Def: πTemp018}
			if πTemp018, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[2] = πg.Param{Name: "file", Def: πTemp018}
			πTemp017 = πg.NewFunction(πg.NewCode("print_stack", "build/src/__python__/traceback.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µf *πg.Object = πArgs[0]; _ = µf
				var µlimit *πg.Object = πArgs[1]; _ = µlimit
				var µfile *πg.Object = πArgs[2]; _ = µfile
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 *πg.BaseException
				_ = πTemp004
				var πTemp005 *πg.Traceback
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πTemp009 []*πg.Object
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					// line 263: """Print a stack trace from its invocation point.
					πF.SetLineno(263)
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
						goto Label1
					}
					goto Label2
					// line 269: if f is None:
					πF.SetLineno(269)
				Label1:
					// line 270: try:
					πF.SetLineno(270)
					πF.PushCheckpoint(4)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßZeroDivisionError); πE != nil {
						continue
					}
					// line 271: raise ZeroDivisionError
					πF.SetLineno(271)
					πE = πF.Raise(πTemp001, nil, nil)
					continue
					πF.PopCheckpoint()
					goto Label3
				Label4:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßZeroDivisionError); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label5
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 272: except ZeroDivisionError:
					πF.SetLineno(272)
				Label5:
					// line 273: f = sys.exc_info()[2].tb_frame.f_back
					πF.SetLineno(273)
					πTemp001 = πg.NewInt(2).ToObject()
					if πTemp006, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßexc_info, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp007.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßtb_frame, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßf_back, nil); πE != nil {
						continue
					}
					µf = πTemp002
					πF.RestoreExc(nil, nil)
					goto Label3
				Label3:
					goto Label2
				Label2:
					// line 274: print_list(extract_stack(f, limit), file)
					πF.SetLineno(274)
					πTemp008 = πF.MakeArgs(2)
					πTemp009 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					πTemp009[0] = µf
					if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
						continue
					}
					πTemp009[1] = µlimit
					if πTemp001, πE = πg.ResolveGlobal(πF, ßextract_stack); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp008[0] = πTemp002
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp008[1] = µfile
					if πTemp001, πE = πg.ResolveGlobal(πF, ßprint_list); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßprint_stack.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 276: def format_stack(f=None, limit=None):
			πF.SetLineno(276)
			πTemp003 = make([]πg.Param, 2)
			if πTemp019, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[0] = πg.Param{Name: "f", Def: πTemp019}
			if πTemp019, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[1] = πg.Param{Name: "limit", Def: πTemp019}
			πTemp018 = πg.NewFunction(πg.NewCode("format_stack", "build/src/__python__/traceback.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µf *πg.Object = πArgs[0]; _ = µf
				var µlimit *πg.Object = πArgs[1]; _ = µlimit
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 *πg.BaseException
				_ = πTemp004
				var πTemp005 *πg.Traceback
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πTemp009 []*πg.Object
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					// line 277: """Shorthand for 'format_list(extract_stack(f, limit))'."""
					πF.SetLineno(277)
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
						goto Label1
					}
					goto Label2
					// line 278: if f is None:
					πF.SetLineno(278)
				Label1:
					// line 279: try:
					πF.SetLineno(279)
					πF.PushCheckpoint(4)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßZeroDivisionError); πE != nil {
						continue
					}
					// line 280: raise ZeroDivisionError
					πF.SetLineno(280)
					πE = πF.Raise(πTemp001, nil, nil)
					continue
					πF.PopCheckpoint()
					goto Label3
				Label4:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßZeroDivisionError); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label5
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 281: except ZeroDivisionError:
					πF.SetLineno(281)
				Label5:
					// line 282: f = sys.exc_info()[2].tb_frame.f_back
					πF.SetLineno(282)
					πTemp001 = πg.NewInt(2).ToObject()
					if πTemp006, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßexc_info, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp007.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßtb_frame, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßf_back, nil); πE != nil {
						continue
					}
					µf = πTemp002
					πF.RestoreExc(nil, nil)
					goto Label3
				Label3:
					goto Label2
				Label2:
					// line 283: return format_list(extract_stack(f, limit))
					πF.SetLineno(283)
					πTemp008 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					πTemp009[0] = µf
					if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
						continue
					}
					πTemp009[1] = µlimit
					if πTemp001, πE = πg.ResolveGlobal(πF, ßextract_stack); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp008[0] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßformat_list); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
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
			if πE = πF.Globals().SetItem(πF, ßformat_stack.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 285: def extract_stack(f=None, limit = None):
			πF.SetLineno(285)
			πTemp003 = make([]πg.Param, 2)
			if πTemp020, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[0] = πg.Param{Name: "f", Def: πTemp020}
			if πTemp020, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[1] = πg.Param{Name: "limit", Def: πTemp020}
			πTemp019 = πg.NewFunction(πg.NewCode("extract_stack", "build/src/__python__/traceback.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µf *πg.Object = πArgs[0]; _ = µf
				var µlimit *πg.Object = πArgs[1]; _ = µlimit
				var µlist *πg.Object = πg.UnboundLocal; _ = µlist
				var µn *πg.Object = πg.UnboundLocal; _ = µn
				var µlineno *πg.Object = πg.UnboundLocal; _ = µlineno
				var µco *πg.Object = πg.UnboundLocal; _ = µco
				var µfilename *πg.Object = πg.UnboundLocal; _ = µfilename
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µline *πg.Object = πg.UnboundLocal; _ = µline
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 *πg.BaseException
				_ = πTemp004
				var πTemp005 *πg.Traceback
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πTemp009 bool
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
					case 10: goto Label10
					case 11: goto Label11
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					// line 286: """Extract the raw traceback from the current stack frame.
					πF.SetLineno(286)
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
						goto Label1
					}
					goto Label2
					// line 294: if f is None:
					πF.SetLineno(294)
				Label1:
					// line 295: try:
					πF.SetLineno(295)
					πF.PushCheckpoint(4)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßZeroDivisionError); πE != nil {
						continue
					}
					// line 296: raise ZeroDivisionError
					πF.SetLineno(296)
					πE = πF.Raise(πTemp001, nil, nil)
					continue
					πF.PopCheckpoint()
					goto Label3
				Label4:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßZeroDivisionError); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label5
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 297: except ZeroDivisionError:
					πF.SetLineno(297)
				Label5:
					// line 298: f = sys.exc_info()[2].tb_frame.f_back
					πF.SetLineno(298)
					πTemp001 = πg.NewInt(2).ToObject()
					if πTemp006, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßexc_info, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp007.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßtb_frame, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßf_back, nil); πE != nil {
						continue
					}
					µf = πTemp002
					πF.RestoreExc(nil, nil)
					goto Label3
				Label3:
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µlimit == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label6
					}
					goto Label7
					// line 299: if limit is None:
					πF.SetLineno(299)
				Label6:
					πTemp008 = πF.MakeArgs(2)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					πTemp008[0] = πTemp001
					πTemp008[1] = ßtracebacklimit.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label8
					}
					goto Label9
					// line 300: if hasattr(sys, 'tracebacklimit'):
					πF.SetLineno(300)
				Label8:
					// line 301: limit = sys.tracebacklimit
					πF.SetLineno(301)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtracebacklimit, nil); πE != nil {
						continue
					}
					µlimit = πTemp002
					goto Label9
				Label9:
					goto Label7
				Label7:
					// line 302: list = []
					πF.SetLineno(302)
					πTemp008 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp008...).ToObject()
					µlist = πTemp001
					// line 303: n = 0
					πF.SetLineno(303)
					µn = πg.NewInt(0).ToObject()
					// line 304: while f is not None and (limit is None or n < limit):
					πF.SetLineno(304)
					πF.PushCheckpoint(11)
					πTemp003 = false
				Label10:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp003 {
						πF.PopCheckpoint()
						goto Label12
					}
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µf != πTemp006).ToObject()
					πTemp001 = πTemp002
					if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp010 {
						goto Label13
					}
					if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp006 = πg.GetBool(µlimit == πTemp007).ToObject()
					πTemp002 = πTemp006
					if πTemp011, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp011 {
						goto Label14
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.LT(πF, µn, µlimit); πE != nil {
						continue
					}
					πTemp002 = πTemp006
				Label14:
					πTemp001 = πTemp002
				Label13:
					if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(10)            
					// line 305: lineno = f.f_lineno
					πF.SetLineno(305)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßf_lineno, nil); πE != nil {
						continue
					}
					µlineno = πTemp001
					// line 306: co = f.f_code
					πF.SetLineno(306)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßf_code, nil); πE != nil {
						continue
					}
					µco = πTemp001
					// line 307: filename = co.co_filename
					πF.SetLineno(307)
					if πE = πg.CheckLocal(πF, µco, "co"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µco, ßco_filename, nil); πE != nil {
						continue
					}
					µfilename = πTemp001
					// line 308: name = co.co_name
					πF.SetLineno(308)
					if πE = πg.CheckLocal(πF, µco, "co"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µco, ßco_name, nil); πE != nil {
						continue
					}
					µname = πTemp001
					// line 309: linecache.checkcache(filename)
					πF.SetLineno(309)
					πTemp008 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp008[0] = µfilename
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlinecache); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßcheckcache, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					// line 310: line = linecache.getline(filename, lineno, f.f_globals)
					πF.SetLineno(310)
					πTemp008 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp008[0] = µfilename
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					πTemp008[1] = µlineno
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßf_globals, nil); πE != nil {
						continue
					}
					πTemp008[2] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlinecache); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßgetline, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					µline = πTemp001
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, µline); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label15
					}
					goto Label16
					// line 311: if line: line = line.strip()
					πF.SetLineno(311)
				Label15:
					// line 311: if line: line = line.strip()
					πF.SetLineno(311)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µline, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µline = πTemp002
					goto Label17
				Label16:
					// line 312: else: line = None
					πF.SetLineno(312)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µline = πTemp001
					goto Label17
				Label17:
					// line 313: list.append((filename, lineno, name, line))
					πF.SetLineno(313)
					πTemp008 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple4(µfilename, µlineno, µname, µline).ToObject()
					πTemp008[0] = πTemp001
					if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µlist, ßappend, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					// line 314: f = f.f_back
					πF.SetLineno(314)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßf_back, nil); πE != nil {
						continue
					}
					µf = πTemp001
					// line 315: n = n+1
					πF.SetLineno(315)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µn = πTemp001
					continue
				Label11:
					if πE != nil || πR != nil {
						continue
					}
				Label12:
					// line 316: list.reverse()
					πF.SetLineno(316)
					if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µlist, ßreverse, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 317: return list
					πF.SetLineno(317)
					if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
						continue
					}
					πR = µlist
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßextract_stack.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 319: def tb_lineno(tb):
			πF.SetLineno(319)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "tb", Def: nil}
			πTemp020 = πg.NewFunction(πg.NewCode("tb_lineno", "build/src/__python__/traceback.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtb *πg.Object = πArgs[0]; _ = µtb
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 320: """Calculate correct line number of traceback given in tb.
					πF.SetLineno(320)
					// line 324: return tb.tb_lineno
					πF.SetLineno(324)
					if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µtb, ßtb_lineno, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßtb_lineno.ToObject(), πTemp020); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("traceback", Code)
}
