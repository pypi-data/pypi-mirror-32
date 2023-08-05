package cStringIO
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/cStringIO.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßEINVAL := πg.InternStr("EINVAL")
		ßFalse := πg.InternStr("False")
		ßIOError := πg.InternStr("IOError")
		ßImportError := πg.InternStr("ImportError")
		ßNone := πg.InternStr("None")
		ßRead := πg.InternStr("Read")
		ßRuntimeError := πg.InternStr("RuntimeError")
		ßStopIteration := πg.InternStr("StopIteration")
		ßStringIO := πg.InternStr("StringIO")
		ßTrue := πg.InternStr("True")
		ßValueError := πg.InternStr("ValueError")
		ß__all__ := πg.InternStr("__all__")
		ß__init__ := πg.InternStr("__init__")
		ß__iter__ := πg.InternStr("__iter__")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_complain_ifclosed := πg.InternStr("_complain_ifclosed")
		ßappend := πg.InternStr("append")
		ßargv := πg.InternStr("argv")
		ßbasestring := πg.InternStr("basestring")
		ßbuf := πg.InternStr("buf")
		ßbuflist := πg.InternStr("buflist")
		ßclose := πg.InternStr("close")
		ßclosed := πg.InternStr("closed")
		ßerrno := πg.InternStr("errno")
		ßfind := πg.InternStr("find")
		ßflush := πg.InternStr("flush")
		ßgetvalue := πg.InternStr("getvalue")
		ßisatty := πg.InternStr("isatty")
		ßisinstance := πg.InternStr("isinstance")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßmax := πg.InternStr("max")
		ßmin := πg.InternStr("min")
		ßnext := πg.InternStr("next")
		ßobject := πg.InternStr("object")
		ßopen := πg.InternStr("open")
		ßpos := πg.InternStr("pos")
		ßr := πg.InternStr("r")
		ßread := πg.InternStr("read")
		ßreadline := πg.InternStr("readline")
		ßreadlines := πg.InternStr("readlines")
		ßrepr := πg.InternStr("repr")
		ßseek := πg.InternStr("seek")
		ßsoftspace := πg.InternStr("softspace")
		ßstr := πg.InternStr("str")
		ßtell := πg.InternStr("tell")
		ßtest := πg.InternStr("test")
		ßtruncate := πg.InternStr("truncate")
		ßwrite := πg.InternStr("write")
		ßwritelines := πg.InternStr("writelines")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.BaseException
		_ = πTemp004
		var πTemp005 *πg.Traceback
		_ = πTemp005
		var πTemp006 bool
		_ = πTemp006
		var πTemp007 []πg.Param
		_ = πTemp007
		var πTemp008 *πg.Dict
		_ = πTemp008
		var πTemp009 *πg.Object
		_ = πTemp009
		var πTemp010 *πg.Object
		_ = πTemp010
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 2: goto Label2
			default: panic("unexpected function state")
			}
			// line 1: r"""File-like objects that read from or write to a string buffer.
			πF.SetLineno(1)
			// line 31: try:
			πF.SetLineno(31)
			πF.PushCheckpoint(2)
			// line 32: import errno
			πF.SetLineno(32)
			if πTemp002, πE = πg.ImportModule(πF, "errno"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßerrno.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 33: EINVAL = errno.EINVAL
			πF.SetLineno(33)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßerrno); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßEINVAL, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßEINVAL.ToObject(), πTemp003); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label1
		Label2:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp004, πTemp005 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label3
			}
			πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
			continue
			// line 34: except ImportError:
			πF.SetLineno(34)
		Label3:
			// line 35: EINVAL = 22
			πF.SetLineno(35)
			if πE = πF.Globals().SetItem(πF, ßEINVAL.ToObject(), πg.NewInt(22).ToObject()); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label1
		Label1:
			// line 37: __all__ = ["StringIO"]
			πF.SetLineno(37)
			πTemp002 = make([]*πg.Object, 1)
			πTemp002[0] = ßStringIO.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 39: def _complain_ifclosed(closed):
			πF.SetLineno(39)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "closed", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("_complain_ifclosed", "build/src/__python__/cStringIO.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µclosed *πg.Object = πArgs[0]; _ = µclosed
				var πTemp001 bool
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
					if πE = πg.CheckLocal(πF, µclosed, "closed"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µclosed); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label1
					}
					goto Label2
					// line 40: if closed:
					πF.SetLineno(40)
				Label1:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					// line 41: raise ValueError, "I/O operation on closed file"
					πF.SetLineno(41)
					πE = πF.Raise(πTemp002, πg.NewStr("I/O operation on closed file").ToObject(), nil)
					continue
					goto Label2
				Label2:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_complain_ifclosed.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 43: class StringIO(object):
			πF.SetLineno(43)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp010, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp010
			πTemp008 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("StringIO", "build/src/__python__/cStringIO.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp008
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 44: """class StringIO([buffer])
					πF.SetLineno(44)
					// line 55: def __init__(self, buf = ''):
					πF.SetLineno(55)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "buf", Def: ß.ToObject()}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/cStringIO.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µbuf *πg.Object = πArgs[1]; _ = µbuf
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
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							πTemp002[0] = µbuf
							if πTemp003, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
							// line 57: if not isinstance(buf, basestring):
							πF.SetLineno(57)
						Label1:
							// line 58: buf = str(buf)
							πF.SetLineno(58)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							πTemp002[0] = µbuf
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µbuf = πTemp003
							goto Label2
						Label2:
							// line 59: self.buf = buf
							πF.SetLineno(59)
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µbuf); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbuf, πTemp001); πE != nil {
								continue
							}
							// line 60: self.len = len(buf)
							πF.SetLineno(60)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
								continue
							}
							πTemp002[0] = µbuf
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlen, πTemp001); πE != nil {
								continue
							}
							// line 61: self.buflist = []
							πF.SetLineno(61)
							πTemp002 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbuflist, πTemp003); πE != nil {
								continue
							}
							// line 62: self.pos = 0
							πF.SetLineno(62)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp001); πE != nil {
								continue
							}
							// line 63: self.closed = False
							πF.SetLineno(63)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßclosed, πTemp003); πE != nil {
								continue
							}
							// line 64: self.softspace = 0
							πF.SetLineno(64)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsoftspace, πTemp001); πE != nil {
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
					// line 66: def __iter__(self):
					πF.SetLineno(66)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__iter__", "build/src/__python__/cStringIO.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 67: return self
							πF.SetLineno(67)
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
					// line 69: def next(self):
					πF.SetLineno(69)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("next", "build/src/__python__/cStringIO.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µr *πg.Object = πg.UnboundLocal; _ = µr
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 70: """A file object is its own iterator, for example iter(f) returns f
							πF.SetLineno(70)
							// line 76: _complain_ifclosed(self.closed)
							πF.SetLineno(76)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßclosed, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_complain_ifclosed); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 77: r = self.readline()
							πF.SetLineno(77)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßreadline, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µr = πTemp003
							if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µr); πE != nil {
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
							// line 78: if not r:
							πF.SetLineno(78)
						Label1:
							if πTemp002, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
								continue
							}
							// line 79: raise StopIteration
							πF.SetLineno(79)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label2
						Label2:
							// line 80: return r
							πF.SetLineno(80)
							if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
								continue
							}
							πR = µr
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
					// line 82: def close(self):
					πF.SetLineno(82)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("close", "build/src/__python__/cStringIO.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 83: """Free the memory buffer.
							πF.SetLineno(83)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßclosed, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 85: if not self.closed:
							πF.SetLineno(85)
						Label1:
							// line 86: self.closed = True
							πF.SetLineno(86)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßclosed, πTemp002); πE != nil {
								continue
							}
							// line 87: del self.buf, self.pos
							πF.SetLineno(87)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.DelAttr(πF, µself, ßbuf); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.DelAttr(πF, µself, ßpos); πE != nil {
								continue
							}
							goto Label2
						Label2:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßclose.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 89: def isatty(self):
					πF.SetLineno(89)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("isatty", "build/src/__python__/cStringIO.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 90: """Returns False because StringIO objects are not connected to a
							πF.SetLineno(90)
							// line 93: _complain_ifclosed(self.closed)
							πF.SetLineno(93)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßclosed, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_complain_ifclosed); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 94: return False
							πF.SetLineno(94)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
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
					if πE = πClass.SetItem(πF, ßisatty.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 96: def seek(self, pos, mode = 0):
					πF.SetLineno(96)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "pos", Def: nil}
					πTemp002[2] = πg.Param{Name: "mode", Def: πg.NewInt(0).ToObject()}
					πTemp007 = πg.NewFunction(πg.NewCode("seek", "build/src/__python__/cStringIO.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µpos *πg.Object = πArgs[1]; _ = µpos
						var µmode *πg.Object = πArgs[2]; _ = µmode
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 97: """Set the file's current position.
							πF.SetLineno(97)
							// line 105: _complain_ifclosed(self.closed)
							πF.SetLineno(105)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßclosed, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_complain_ifclosed); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbuflist, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 106: if self.buflist:
							πF.SetLineno(106)
						Label1:
							// line 107: self.buf += ''.join(self.buflist)
							πF.SetLineno(107)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbuf, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbuflist, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.IAdd(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbuf, πTemp003); πE != nil {
								continue
							}
							// line 108: self.buflist = []
							πF.SetLineno(108)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbuflist, πTemp003); πE != nil {
								continue
							}
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µmode, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µmode, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 109: if mode == 1:
							πF.SetLineno(109)
						Label3:
							// line 110: pos += self.pos
							πF.SetLineno(110)
							if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, µpos, πTemp002); πE != nil {
								continue
							}
							µpos = πTemp003
							goto Label5
							// line 111: elif mode == 2:
							πF.SetLineno(111)
						Label4:
							// line 112: pos += self.len
							πF.SetLineno(112)
							if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlen, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, µpos, πTemp002); πE != nil {
								continue
							}
							µpos = πTemp003
							goto Label5
						Label5:
							// line 113: self.pos = max(0, pos)
							πF.SetLineno(113)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
								continue
							}
							πTemp001[1] = µpos
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßseek.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 115: def tell(self):
					πF.SetLineno(115)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("tell", "build/src/__python__/cStringIO.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 116: """Return the file's current position."""
							πF.SetLineno(116)
							// line 117: _complain_ifclosed(self.closed)
							πF.SetLineno(117)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßclosed, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_complain_ifclosed); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 118: return self.pos
							πF.SetLineno(118)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtell.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 120: def read(self, n = -1):
					πF.SetLineno(120)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp010, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "n", Def: πTemp010}
					πTemp009 = πg.NewFunction(πg.NewCode("read", "build/src/__python__/cStringIO.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µn *πg.Object = πArgs[1]; _ = µn
						var µnewpos *πg.Object = πg.UnboundLocal; _ = µnewpos
						var µr *πg.Object = πg.UnboundLocal; _ = µr
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 121: """Read at most size bytes from the file
							πF.SetLineno(121)
							// line 128: _complain_ifclosed(self.closed)
							πF.SetLineno(128)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßclosed, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_complain_ifclosed); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbuflist, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 129: if self.buflist:
							πF.SetLineno(129)
						Label1:
							// line 130: self.buf += ''.join(self.buflist)
							πF.SetLineno(130)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbuf, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbuflist, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.IAdd(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbuf, πTemp003); πE != nil {
								continue
							}
							// line 131: self.buflist = []
							πF.SetLineno(131)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbuflist, πTemp003); πE != nil {
								continue
							}
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µn == πTemp005).ToObject()
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LT(πF, µn, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
						Label3:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 132: if n is None or n < 0:
							πF.SetLineno(132)
						Label4:
							// line 133: newpos = self.len
							πF.SetLineno(133)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlen, nil); πE != nil {
								continue
							}
							µnewpos = πTemp002
							goto Label6
						Label5:
							// line 135: newpos = min(self.pos+n, self.len)
							πF.SetLineno(135)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, µn); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlen, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µnewpos = πTemp003
							goto Label6
						Label6:
							// line 136: r = self.buf[self.pos:newpos]
							πF.SetLineno(136)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnewpos, "newpos"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, µnewpos, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßbuf, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							µr = πTemp003
							// line 137: self.pos = newpos
							πF.SetLineno(137)
							if πE = πg.CheckLocal(πF, µnewpos, "newpos"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µnewpos); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp002); πE != nil {
								continue
							}
							// line 138: return r
							πF.SetLineno(138)
							if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
								continue
							}
							πR = µr
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßread.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 140: def readline(self, length=None):
					πF.SetLineno(140)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "length", Def: πTemp011}
					πTemp010 = πg.NewFunction(πg.NewCode("readline", "build/src/__python__/cStringIO.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlength *πg.Object = πArgs[1]; _ = µlength
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µnewpos *πg.Object = πg.UnboundLocal; _ = µnewpos
						var µr *πg.Object = πg.UnboundLocal; _ = µr
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 141: r"""Read one entire line from the file.
							πF.SetLineno(141)
							// line 153: _complain_ifclosed(self.closed)
							πF.SetLineno(153)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßclosed, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_complain_ifclosed); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbuflist, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 154: if self.buflist:
							πF.SetLineno(154)
						Label1:
							// line 155: self.buf += ''.join(self.buflist)
							πF.SetLineno(155)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbuf, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbuflist, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.IAdd(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbuf, πTemp003); πE != nil {
								continue
							}
							// line 156: self.buflist = []
							πF.SetLineno(156)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbuflist, πTemp003); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 157: i = self.buf.find('\n', self.pos)
							πF.SetLineno(157)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbuf, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßfind, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µi = πTemp002
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, µi, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 158: if i < 0:
							πF.SetLineno(158)
						Label3:
							// line 159: newpos = self.len
							πF.SetLineno(159)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlen, nil); πE != nil {
								continue
							}
							µnewpos = πTemp002
							goto Label5
						Label4:
							// line 161: newpos = i+1
							πF.SetLineno(161)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µnewpos = πTemp002
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µlength != πTemp005).ToObject()
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GE(πF, µlength, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
						Label6:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							goto Label8
							// line 162: if length is not None and length >= 0:
							πF.SetLineno(162)
						Label7:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp005, µlength); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnewpos, "newpos"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, πTemp003, µnewpos); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label9
							}
							goto Label10
							// line 163: if self.pos + length < newpos:
							πF.SetLineno(163)
						Label9:
							// line 164: newpos = self.pos + length
							πF.SetLineno(164)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, µlength); πE != nil {
								continue
							}
							µnewpos = πTemp002
							goto Label10
						Label10:
							goto Label8
						Label8:
							// line 165: r = self.buf[self.pos:newpos]
							πF.SetLineno(165)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnewpos, "newpos"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, µnewpos, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßbuf, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							µr = πTemp003
							// line 166: self.pos = newpos
							πF.SetLineno(166)
							if πE = πg.CheckLocal(πF, µnewpos, "newpos"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µnewpos); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp002); πE != nil {
								continue
							}
							// line 167: return r
							πF.SetLineno(167)
							if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
								continue
							}
							πR = µr
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßreadline.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 169: def readlines(self, sizehint = 0):
					πF.SetLineno(169)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "sizehint", Def: πg.NewInt(0).ToObject()}
					πTemp011 = πg.NewFunction(πg.NewCode("readlines", "build/src/__python__/cStringIO.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsizehint *πg.Object = πArgs[1]; _ = µsizehint
						var µtotal *πg.Object = πg.UnboundLocal; _ = µtotal
						var µlines *πg.Object = πg.UnboundLocal; _ = µlines
						var µline *πg.Object = πg.UnboundLocal; _ = µline
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 170: """Read until EOF using readline() and return a list containing the
							πF.SetLineno(170)
							// line 177: total = 0
							πF.SetLineno(177)
							µtotal = πg.NewInt(0).ToObject()
							// line 178: lines = []
							πF.SetLineno(178)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µlines = πTemp002
							// line 179: line = self.readline()
							πF.SetLineno(179)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßreadline, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µline = πTemp003
							// line 180: while line:
							πF.SetLineno(180)
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
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µline); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 181: lines.append(line)
							πF.SetLineno(181)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp001[0] = µline
							if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlines, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 182: total += len(line)
							πF.SetLineno(182)
							if πE = πg.CheckLocal(πF, µtotal, "total"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp001[0] = µline
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.IAdd(πF, µtotal, πTemp003); πE != nil {
								continue
							}
							µtotal = πTemp002
							if πE = πg.CheckLocal(πF, µsizehint, "sizehint"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, πg.NewInt(0).ToObject(), µsizehint); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µtotal, "total"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LE(πF, µsizehint, µtotal); πE != nil {
								continue
							}
						Label4:
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							goto Label6
							// line 183: if 0 < sizehint <= total:
							πF.SetLineno(183)
						Label5:
							// line 184: break
							πF.SetLineno(184)
							πTemp004 = true
							continue
							goto Label6
						Label6:
							// line 185: line = self.readline()
							πF.SetLineno(185)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßreadline, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µline = πTemp003
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 186: return lines
							πF.SetLineno(186)
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
					if πE = πClass.SetItem(πF, ßreadlines.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 188: def truncate(self, size=None):
					πF.SetLineno(188)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "size", Def: πTemp013}
					πTemp012 = πg.NewFunction(πg.NewCode("truncate", "build/src/__python__/cStringIO.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsize *πg.Object = πArgs[1]; _ = µsize
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 189: """Truncate the file's size.
							πF.SetLineno(189)
							// line 199: _complain_ifclosed(self.closed)
							πF.SetLineno(199)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßclosed, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_complain_ifclosed); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µsize == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, µsize, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, µsize, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 200: if size is None:
							πF.SetLineno(200)
						Label1:
							// line 201: size = self.pos
							πF.SetLineno(201)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							µsize = πTemp002
							goto Label4
							// line 202: elif size < 0:
							πF.SetLineno(202)
						Label2:
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßEINVAL); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("Negative size not allowed").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 203: raise IOError(EINVAL, "Negative size not allowed")
							πF.SetLineno(203)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label4
							// line 204: elif size < self.pos:
							πF.SetLineno(204)
						Label3:
							// line 205: self.pos = size
							πF.SetLineno(205)
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µsize); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp002); πE != nil {
								continue
							}
							goto Label4
						Label4:
							// line 206: self.buf = self.getvalue()[:size]
							πF.SetLineno(206)
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µsize, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßgetvalue, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbuf, πTemp002); πE != nil {
								continue
							}
							// line 207: self.len = size
							πF.SetLineno(207)
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µsize); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlen, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtruncate.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 209: def write(self, s):
					πF.SetLineno(209)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "s", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("write", "build/src/__python__/cStringIO.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µs *πg.Object = πArgs[1]; _ = µs
						var µspos *πg.Object = πg.UnboundLocal; _ = µspos
						var µslen *πg.Object = πg.UnboundLocal; _ = µslen
						var µnewpos *πg.Object = πg.UnboundLocal; _ = µnewpos
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 210: """Write a string to the file.
							πF.SetLineno(210)
							// line 214: _complain_ifclosed(self.closed)
							πF.SetLineno(214)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßclosed, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_complain_ifclosed); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µs); πE != nil {
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
							// line 215: if not s: return
							πF.SetLineno(215)
						Label1:
							// line 215: if not s: return
							πF.SetLineno(215)
							πR = πg.None
							continue
							goto Label2
						Label2:
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 217: if not isinstance(s, basestring):
							πF.SetLineno(217)
						Label3:
							// line 218: s = str(s)
							πF.SetLineno(218)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp003
							goto Label4
						Label4:
							// line 219: spos = self.pos
							πF.SetLineno(219)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpos, nil); πE != nil {
								continue
							}
							µspos = πTemp002
							// line 220: slen = self.len
							πF.SetLineno(220)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlen, nil); πE != nil {
								continue
							}
							µslen = πTemp002
							if πE = πg.CheckLocal(πF, µspos, "spos"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µslen, "slen"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µspos, µslen); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							goto Label6
							// line 221: if spos == slen:
							πF.SetLineno(221)
						Label5:
							// line 222: self.buflist.append(s)
							πF.SetLineno(222)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbuflist, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 223: self.len = self.pos = spos + len(s)
							πF.SetLineno(223)
							if πE = πg.CheckLocal(πF, µspos, "spos"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Add(πF, µspos, πTemp005); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlen, πTemp003); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp003); πE != nil {
								continue
							}
							// line 224: return
							πF.SetLineno(224)
							πR = πg.None
							continue
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µspos, "spos"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µslen, "slen"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GT(πF, µspos, µslen); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							goto Label8
							// line 225: if spos > slen:
							πF.SetLineno(225)
						Label7:
							// line 226: self.buflist.append('\0'*(spos - slen))
							πF.SetLineno(226)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µspos, "spos"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µslen, "slen"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, µspos, µslen); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mul(πF, πg.NewStr("\x00").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbuflist, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 227: slen = spos
							πF.SetLineno(227)
							if πE = πg.CheckLocal(πF, µspos, "spos"); πE != nil {
								continue
							}
							µslen = µspos
							goto Label8
						Label8:
							// line 228: newpos = spos + len(s)
							πF.SetLineno(228)
							if πE = πg.CheckLocal(πF, µspos, "spos"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Add(πF, µspos, πTemp005); πE != nil {
								continue
							}
							µnewpos = πTemp002
							if πE = πg.CheckLocal(πF, µspos, "spos"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µslen, "slen"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, µspos, µslen); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label9
							}
							goto Label10
							// line 229: if spos < slen:
							πF.SetLineno(229)
						Label9:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbuflist, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label12
							}
							goto Label13
							// line 230: if self.buflist:
							πF.SetLineno(230)
						Label12:
							// line 231: self.buf += ''.join(self.buflist)
							πF.SetLineno(231)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbuf, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbuflist, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.IAdd(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbuf, πTemp003); πE != nil {
								continue
							}
							goto Label13
						Label13:
							// line 232: self.buflist = [self.buf[:spos], s, self.buf[newpos:]]
							πF.SetLineno(232)
							πTemp001 = make([]*πg.Object, 3)
							if πE = πg.CheckLocal(πF, µspos, "spos"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µspos, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßbuf, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							if πE = πg.CheckLocal(πF, µnewpos, "newpos"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µnewpos, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßbuf, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							πTemp001[2] = πTemp003
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbuflist, πTemp003); πE != nil {
								continue
							}
							// line 233: self.buf = ''
							πF.SetLineno(233)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ß.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbuf, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnewpos, "newpos"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µslen, "slen"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GT(πF, µnewpos, µslen); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label14
							}
							goto Label15
							// line 234: if newpos > slen:
							πF.SetLineno(234)
						Label14:
							// line 235: slen = newpos
							πF.SetLineno(235)
							if πE = πg.CheckLocal(πF, µnewpos, "newpos"); πE != nil {
								continue
							}
							µslen = µnewpos
							goto Label15
						Label15:
							goto Label11
						Label10:
							// line 237: self.buflist.append(s)
							πF.SetLineno(237)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbuflist, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 238: slen = newpos
							πF.SetLineno(238)
							if πE = πg.CheckLocal(πF, µnewpos, "newpos"); πE != nil {
								continue
							}
							µslen = µnewpos
							goto Label11
						Label11:
							// line 239: self.len = slen
							πF.SetLineno(239)
							if πE = πg.CheckLocal(πF, µslen, "slen"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µslen); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlen, πTemp002); πE != nil {
								continue
							}
							// line 240: self.pos = newpos
							πF.SetLineno(240)
							if πE = πg.CheckLocal(πF, µnewpos, "newpos"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µnewpos); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßwrite.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 242: def writelines(self, iterable):
					πF.SetLineno(242)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "iterable", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("writelines", "build/src/__python__/cStringIO.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µiterable *πg.Object = πArgs[1]; _ = µiterable
						var µwrite *πg.Object = πg.UnboundLocal; _ = µwrite
						var µline *πg.Object = πg.UnboundLocal; _ = µline
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
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 243: """Write a sequence of strings to the file. The sequence can be any
							πF.SetLineno(243)
							// line 250: write = self.write
							πF.SetLineno(250)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							µwrite = πTemp001
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
								µline = πTemp004
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 252: write(line)
							πF.SetLineno(252)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp005[0] = µline
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp004, πE = µwrite.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
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
					if πE = πClass.SetItem(πF, ßwritelines.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 254: def flush(self):
					πF.SetLineno(254)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("flush", "build/src/__python__/cStringIO.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 255: """Flush the internal buffer
							πF.SetLineno(255)
							// line 257: _complain_ifclosed(self.closed)
							πF.SetLineno(257)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßclosed, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_complain_ifclosed); πE != nil {
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
					if πE = πClass.SetItem(πF, ßflush.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 259: def getvalue(self):
					πF.SetLineno(259)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("getvalue", "build/src/__python__/cStringIO.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 260: """
							πF.SetLineno(260)
							// line 270: _complain_ifclosed(self.closed)
							πF.SetLineno(270)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßclosed, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_complain_ifclosed); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbuflist, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 271: if self.buflist:
							πF.SetLineno(271)
						Label1:
							// line 272: self.buf += ''.join(self.buflist)
							πF.SetLineno(272)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbuf, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbuflist, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.IAdd(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbuf, πTemp003); πE != nil {
								continue
							}
							// line 273: self.buflist = []
							πF.SetLineno(273)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbuflist, πTemp003); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 274: return self.buf
							πF.SetLineno(274)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbuf, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßgetvalue.ToObject(), πTemp016); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp009, πE = πTemp008.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp009 == nil {
				πTemp009 = πg.TypeType.ToObject()
			}
			if πTemp010, πE = πTemp009.Call(πF, []*πg.Object{πg.NewStr("StringIO").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp008.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStringIO.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 279: def test():
			πF.SetLineno(279)
			πTemp007 = make([]πg.Param, 0)
			πTemp003 = πg.NewFunction(πg.NewCode("test", "build/src/__python__/cStringIO.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsys *πg.Object = πg.UnboundLocal; _ = µsys
				var µfile *πg.Object = πg.UnboundLocal; _ = µfile
				var µlines *πg.Object = πg.UnboundLocal; _ = µlines
				var µtext *πg.Object = πg.UnboundLocal; _ = µtext
				var µf *πg.Object = πg.UnboundLocal; _ = µf
				var µline *πg.Object = πg.UnboundLocal; _ = µline
				var µlength *πg.Object = πg.UnboundLocal; _ = µlength
				var µline2 *πg.Object = πg.UnboundLocal; _ = µline2
				var µlist *πg.Object = πg.UnboundLocal; _ = µlist
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
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 4: goto Label4
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 280: import sys
					πF.SetLineno(280)
					if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					µsys = πTemp001
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
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 281: if sys.argv[1:]:
					πF.SetLineno(281)
				Label1:
					// line 282: file = sys.argv[1]
					πF.SetLineno(282)
					πTemp001 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µsys, ßargv, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
						continue
					}
					µfile = πTemp003
					goto Label3
				Label2:
					// line 284: file = '/etc/passwd'
					πF.SetLineno(284)
					µfile = πg.NewStr("/etc/passwd").ToObject()
					goto Label3
				Label3:
					// line 285: lines = open(file, 'r').readlines()
					πF.SetLineno(285)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp002[0] = µfile
					πTemp002[1] = ßr.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßreadlines, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µlines = πTemp003
					// line 286: text = open(file, 'r').read()
					πF.SetLineno(286)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp002[0] = µfile
					πTemp002[1] = ßr.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßread, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µtext = πTemp003
					// line 287: f = StringIO()
					πF.SetLineno(287)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µf = πTemp003
					if πTemp004, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp004, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µlines, πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
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
					πF.PushCheckpoint(4)            
					// line 289: f.write(line)
					πF.SetLineno(289)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp002[0] = µline
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µf, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					continue
				Label5:
					if πE != nil || πR != nil {
						continue
					}
				Label6:
					// line 290: f.writelines(lines[-2:])
					πF.SetLineno(290)
					πTemp002 = πF.MakeArgs(1)
					if πTemp003, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µlines, πTemp001); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßwritelines, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µf, ßgetvalue, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, πTemp004, µtext); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label7
					}
					goto Label8
					// line 291: if f.getvalue() != text:
					πF.SetLineno(291)
				Label7:
					if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
						continue
					}
					// line 292: raise RuntimeError, 'write failed'
					πF.SetLineno(292)
					πE = πF.Raise(πTemp001, πg.NewStr("write failed").ToObject(), nil)
					continue
					goto Label8
				Label8:
					// line 293: length = f.tell()
					πF.SetLineno(293)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßtell, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µlength = πTemp003
					// line 294: print 'File length =', length
					πF.SetLineno(294)
					πTemp002 = make([]*πg.Object, 2)
					πTemp002[0] = πg.NewStr("File length =").ToObject()
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					πTemp002[1] = µlength
					if πE = πg.Print(πF, πTemp002, true); πE != nil {
						continue
					}
					// line 295: f.seek(len(lines[0]))
					πF.SetLineno(295)
					πTemp002 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					πTemp001 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µlines, πTemp001); πE != nil {
						continue
					}
					πTemp007[0] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp002[0] = πTemp003
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßseek, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 296: f.write(lines[1])
					πF.SetLineno(296)
					πTemp002 = πF.MakeArgs(1)
					πTemp001 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µlines, πTemp001); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 297: f.seek(0)
					πF.SetLineno(297)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßseek, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 298: print 'First line =', repr(f.readline())
					πF.SetLineno(298)
					πTemp002 = make([]*πg.Object, 2)
					πTemp002[0] = πg.NewStr("First line =").ToObject()
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßreadline, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp007[0] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp002[1] = πTemp003
					if πE = πg.Print(πF, πTemp002, true); πE != nil {
						continue
					}
					// line 299: print 'Position =', f.tell()
					πF.SetLineno(299)
					πTemp002 = make([]*πg.Object, 2)
					πTemp002[0] = πg.NewStr("Position =").ToObject()
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßtell, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[1] = πTemp003
					if πE = πg.Print(πF, πTemp002, true); πE != nil {
						continue
					}
					// line 300: line = f.readline()
					πF.SetLineno(300)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßreadline, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µline = πTemp003
					// line 301: print 'Second line =', repr(line)
					πF.SetLineno(301)
					πTemp002 = make([]*πg.Object, 2)
					πTemp002[0] = πg.NewStr("Second line =").ToObject()
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp007[0] = µline
					if πTemp001, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp002[1] = πTemp003
					if πE = πg.Print(πF, πTemp002, true); πE != nil {
						continue
					}
					// line 302: f.seek(-len(line), 1)
					πF.SetLineno(302)
					πTemp002 = πF.MakeArgs(2)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp007[0] = µline
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πTemp001, πE = πg.Neg(πF, πTemp004); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					πTemp002[1] = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßseek, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 303: line2 = f.read(len(line))
					πF.SetLineno(303)
					πTemp002 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp007[0] = µline
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp002[0] = πTemp003
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßread, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µline2 = πTemp003
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µline2, "line2"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, µline, µline2); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label9
					}
					goto Label10
					// line 304: if line != line2:
					πF.SetLineno(304)
				Label9:
					if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
						continue
					}
					// line 305: raise RuntimeError, 'bad result after seek back'
					πF.SetLineno(305)
					πE = πF.Raise(πTemp001, πg.NewStr("bad result after seek back").ToObject(), nil)
					continue
					goto Label10
				Label10:
					// line 306: f.seek(len(line2), 1)
					πF.SetLineno(306)
					πTemp002 = πF.MakeArgs(2)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µline2, "line2"); πE != nil {
						continue
					}
					πTemp007[0] = µline2
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp002[0] = πTemp003
					πTemp002[1] = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßseek, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 307: list = f.readlines()
					πF.SetLineno(307)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßreadlines, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µlist = πTemp003
					// line 308: line = list[-1]
					πF.SetLineno(308)
					if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µlist, πTemp001); πE != nil {
						continue
					}
					µline = πTemp003
					// line 309: f.seek(f.tell() - len(line))
					πF.SetLineno(309)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µf, ßtell, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp007[0] = µline
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πTemp001, πE = πg.Sub(πF, πTemp004, πTemp008); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßseek, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 310: line2 = f.read()
					πF.SetLineno(310)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßread, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µline2 = πTemp003
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µline2, "line2"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, µline, µline2); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label11
					}
					goto Label12
					// line 311: if line != line2:
					πF.SetLineno(311)
				Label11:
					if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
						continue
					}
					// line 312: raise RuntimeError, 'bad result after seek back from EOF'
					πF.SetLineno(312)
					πE = πF.Raise(πTemp001, πg.NewStr("bad result after seek back from EOF").ToObject(), nil)
					continue
					goto Label12
				Label12:
					// line 313: print 'Read', len(list), 'more lines'
					πF.SetLineno(313)
					πTemp002 = make([]*πg.Object, 3)
					πTemp002[0] = ßRead.ToObject()
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
						continue
					}
					πTemp007[0] = µlist
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp002[1] = πTemp003
					πTemp002[2] = πg.NewStr("more lines").ToObject()
					if πE = πg.Print(πF, πTemp002, true); πE != nil {
						continue
					}
					// line 314: print 'File length =', f.tell()
					πF.SetLineno(314)
					πTemp002 = make([]*πg.Object, 2)
					πTemp002[0] = πg.NewStr("File length =").ToObject()
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßtell, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[1] = πTemp003
					if πE = πg.Print(πF, πTemp002, true); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µf, ßtell, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, πTemp004, µlength); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label13
					}
					goto Label14
					// line 315: if f.tell() != length:
					πF.SetLineno(315)
				Label13:
					if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
						continue
					}
					// line 316: raise RuntimeError, 'bad length'
					πF.SetLineno(316)
					πE = πF.Raise(πTemp001, πg.NewStr("bad length").ToObject(), nil)
					continue
					goto Label14
				Label14:
					// line 317: f.truncate(length/2)
					πF.SetLineno(317)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Div(πF, µlength, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßtruncate, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 318: f.seek(0, 2)
					πF.SetLineno(318)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewInt(0).ToObject()
					πTemp002[1] = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßseek, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 319: print 'Truncated length =', f.tell()
					πF.SetLineno(319)
					πTemp002 = make([]*πg.Object, 2)
					πTemp002[0] = πg.NewStr("Truncated length =").ToObject()
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßtell, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[1] = πTemp003
					if πE = πg.Print(πF, πTemp002, true); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µf, ßtell, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Div(πF, µlength, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, πTemp004, πTemp003); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label15
					}
					goto Label16
					// line 320: if f.tell() != length/2:
					πF.SetLineno(320)
				Label15:
					if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
						continue
					}
					// line 321: raise RuntimeError, 'truncate did not adjust length'
					πF.SetLineno(321)
					πE = πF.Raise(πTemp001, πg.NewStr("truncate did not adjust length").ToObject(), nil)
					continue
					goto Label16
				Label16:
					// line 322: f.close()
					πF.SetLineno(322)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µf, ßclose, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßtest.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp009, πE = πg.Eq(πF, πTemp010, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp009); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label4
			}
			goto Label5
			// line 324: if __name__ == '__main__':
			πF.SetLineno(324)
		Label4:
			// line 325: test()
			πF.SetLineno(325)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßtest); πE != nil {
				continue
			}
			if πTemp010, πE = πTemp009.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label5
		Label5:
		}
		return nil, πE
	})
	πg.RegisterModule("cStringIO", Code)
}
