package test_select
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/test/test_select.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßAlmost := πg.InternStr("Almost")
		ßEOF := πg.InternStr("EOF")
		ßNone := πg.InternStr("None")
		ßNope := πg.InternStr("Nope")
		ßSelectTestCase := πg.InternStr("SelectTestCase")
		ßTestCase := πg.InternStr("TestCase")
		ßTypeError := πg.InternStr("TypeError")
		ßValueError := πg.InternStr("ValueError")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertIsNot := πg.InternStr("assertIsNot")
		ßassertRaises := πg.InternStr("assertRaises")
		ßclose := πg.InternStr("close")
		ßfail := πg.InternStr("fail")
		ßfileno := πg.InternStr("fileno")
		ßobject := πg.InternStr("object")
		ßos := πg.InternStr("os")
		ßpopen := πg.InternStr("popen")
		ßr := πg.InternStr("r")
		ßreadline := πg.InternStr("readline")
		ßreap_children := πg.InternStr("reap_children")
		ßrepr := πg.InternStr("repr")
		ßrun_unittest := πg.InternStr("run_unittest")
		ßselect := πg.InternStr("select")
		ßstdout := πg.InternStr("stdout")
		ßsys := πg.InternStr("sys")
		ßtest_error_conditions := πg.InternStr("test_error_conditions")
		ßtest_main := πg.InternStr("test_main")
		ßtest_returned_list_identity := πg.InternStr("test_returned_list_identity")
		ßtest_select := πg.InternStr("test_select")
		ßtest_select_mutated := πg.InternStr("test_select_mutated")
		ßtest_support := πg.InternStr("test_support")
		ßunittest := πg.InternStr("unittest")
		ßverbose := πg.InternStr("verbose")
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
		var πTemp006 *πg.Object
		_ = πTemp006
		var πTemp007 []πg.Param
		_ = πTemp007
		var πTemp008 bool
		_ = πTemp008
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: from test import test_support
			πF.SetLineno(1)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßtest_support.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 2: import unittest
			πF.SetLineno(2)
			if πTemp002, πE = πg.ImportModule(πF, "unittest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 3: import select_ as select
			πF.SetLineno(3)
			if πTemp002, πE = πg.ImportModule(πF, "select_"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßselect.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 4: import os
			πF.SetLineno(4)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 5: import sys
			πF.SetLineno(5)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 7: @unittest.skipIf(sys.platform[:3] in ('win', 'os2', 'riscos'),
			πF.SetLineno(7)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßTestCase, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("SelectTestCase", "build/src/__python__/test/test_select.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Dict
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 []πg.Param
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 11: class Nope(object):
					πF.SetLineno(11)
					πTemp003 = make([]*πg.Object, 1)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßobject); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					πTemp001 = πg.NewDict()
					if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
						continue
					}
					_, πE = πg.NewCode("Nope", "build/src/__python__/test/test_select.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
						πClass := πTemp001
						_ = πClass
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 12: pass
							πF.SetLineno(12)
						}
						return nil, nil
					}).Eval(πF, πF.Globals(), nil, nil)
					if πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
						continue
					}
					if πTemp004 == nil {
						πTemp004 = πg.TypeType.ToObject()
					}
					if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("Nope").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßNope.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 14: class Almost(object):
					πF.SetLineno(14)
					πTemp003 = make([]*πg.Object, 1)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßobject); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					πTemp001 = πg.NewDict()
					if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
						continue
					}
					_, πE = πg.NewCode("Almost", "build/src/__python__/test/test_select.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
						πClass := πTemp001
						_ = πClass
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []πg.Param
						_ = πTemp002
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 15: def fileno(self):
							πF.SetLineno(15)
							πTemp002 = make([]πg.Param, 1)
							πTemp002[0] = πg.Param{Name: "self", Def: nil}
							πTemp001 = πg.NewFunction(πg.NewCode("fileno", "build/src/__python__/test/test_select.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 16: return 'fileno'
									πF.SetLineno(16)
									πR = ßfileno.ToObject()
									continue
								}
								if πE != nil {
									πR = nil
								} else if πR == nil {
									πR = πg.None
								}
								return πR, πE
							}), πF.Globals()).ToObject()
							if πE = πClass.SetItem(πF, ßfileno.ToObject(), πTemp001); πE != nil {
								continue
							}
						}
						return nil, nil
					}).Eval(πF, πF.Globals(), nil, nil)
					if πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
						continue
					}
					if πTemp004 == nil {
						πTemp004 = πg.TypeType.ToObject()
					}
					if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("Almost").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßAlmost.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 18: def test_error_conditions(self):
					πF.SetLineno(18)
					πTemp006 = make([]πg.Param, 1)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("test_error_conditions", "build/src/__python__/test/test_select.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 []*πg.Object
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
							// line 19: self.assertRaises(TypeError, select.select, 1, 2, 3)
							πF.SetLineno(19)
							πTemp001 = πF.MakeArgs(5)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßselect); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßselect, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewInt(1).ToObject()
							πTemp001[3] = πg.NewInt(2).ToObject()
							πTemp001[4] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 20: self.assertRaises(TypeError, select.select, [self.Nope()], [], [])
							πF.SetLineno(20)
							πTemp001 = πF.MakeArgs(5)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßselect); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßselect, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp004 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßNope, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							πTemp002 = πg.NewList(πTemp004...).ToObject()
							πTemp001[2] = πTemp002
							πTemp004 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp004...).ToObject()
							πTemp001[3] = πTemp002
							πTemp004 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp004...).ToObject()
							πTemp001[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 21: self.assertRaises(TypeError, select.select, [self.Almost()], [], [])
							πF.SetLineno(21)
							πTemp001 = πF.MakeArgs(5)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßselect); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßselect, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp004 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßAlmost, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							πTemp002 = πg.NewList(πTemp004...).ToObject()
							πTemp001[2] = πTemp002
							πTemp004 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp004...).ToObject()
							πTemp001[3] = πTemp002
							πTemp004 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp004...).ToObject()
							πTemp001[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 22: self.assertRaises(ValueError, select.select, [], [], [], "not a number")
							πF.SetLineno(22)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßselect); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßselect, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp004 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp004...).ToObject()
							πTemp001[2] = πTemp002
							πTemp004 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp004...).ToObject()
							πTemp001[3] = πTemp002
							πTemp004 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp004...).ToObject()
							πTemp001[4] = πTemp002
							πTemp001[5] = πg.NewStr("not a number").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_error_conditions.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 24: def test_returned_list_identity(self):
					πF.SetLineno(24)
					πTemp006 = make([]πg.Param, 1)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_returned_list_identity", "build/src/__python__/test/test_select.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µr *πg.Object = πg.UnboundLocal; _ = µr
						var µw *πg.Object = πg.UnboundLocal; _ = µw
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
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
							// line 26: r, w, x = select.select([], [], [], 1)
							πF.SetLineno(26)
							πTemp001 = πF.MakeArgs(4)
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[2] = πTemp003
							πTemp001[3] = πg.NewInt(1).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßselect); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßselect, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp003); πE != nil {
								continue
							}
							µr = πTemp004
							µw = πTemp005
							µx = πTemp006
							// line 27: self.assertIsNot(r, w)
							πF.SetLineno(27)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
								continue
							}
							πTemp001[0] = µr
							if πE = πg.CheckLocal(πF, µw, "w"); πE != nil {
								continue
							}
							πTemp001[1] = µw
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertIsNot, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 28: self.assertIsNot(r, x)
							πF.SetLineno(28)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
								continue
							}
							πTemp001[0] = µr
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[1] = µx
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertIsNot, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 29: self.assertIsNot(w, x)
							πF.SetLineno(29)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µw, "w"); πE != nil {
								continue
							}
							πTemp001[0] = µw
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[1] = µx
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertIsNot, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_returned_list_identity.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 31: def test_select(self):
					πF.SetLineno(31)
					πTemp006 = make([]πg.Param, 1)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_select", "build/src/__python__/test/test_select.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µcmd *πg.Object = πg.UnboundLocal; _ = µcmd
						var µp *πg.Object = πg.UnboundLocal; _ = µp
						var µtout *πg.Object = πg.UnboundLocal; _ = µtout
						var µrfd *πg.Object = πg.UnboundLocal; _ = µrfd
						var µwfd *πg.Object = πg.UnboundLocal; _ = µwfd
						var µxfd *πg.Object = πg.UnboundLocal; _ = µxfd
						var µline *πg.Object = πg.UnboundLocal; _ = µline
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
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πTemp011 *πg.Object
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
							// line 32: cmd = 'for i in 0 1 2 3 4 5 6 7 8 9; do echo testing...; sleep 1; done'
							πF.SetLineno(32)
							µcmd = πg.NewStr("for i in 0 1 2 3 4 5 6 7 8 9; do echo testing...; sleep 1; done").ToObject()
							// line 33: p = os.popen(cmd, 'r')
							πF.SetLineno(33)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µcmd, "cmd"); πE != nil {
								continue
							}
							πTemp001[0] = µcmd
							πTemp001[1] = ßr.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpopen, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µp = πTemp002
							πTemp004 = πg.NewTuple6(πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject(), πg.NewInt(4).ToObject(), πg.NewInt(8).ToObject(), πg.NewInt(16).ToObject()).ToObject()
							if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp006 = πg.NewTuple1(πTemp007).ToObject()
							if πTemp005, πE = πg.Mul(πF, πTemp006, πg.NewInt(10).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								µtout = πTemp003
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßverbose, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label4
							}
							goto Label5
							// line 35: if test_support.verbose:
							πF.SetLineno(35)
						Label4:
							// line 36: print 'timeout =', tout
							πF.SetLineno(36)
							πTemp001 = make([]*πg.Object, 2)
							πTemp001[0] = πg.NewStr("timeout =").ToObject()
							if πE = πg.CheckLocal(πF, µtout, "tout"); πE != nil {
								continue
							}
							πTemp001[1] = µtout
							if πE = πg.Print(πF, πTemp001, true); πE != nil {
								continue
							}
							goto Label5
						Label5:
							// line 37: rfd, wfd, xfd = select.select([p], [], [], tout)
							πF.SetLineno(37)
							πTemp001 = πF.MakeArgs(4)
							πTemp010 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp010[0] = µp
							πTemp003 = πg.NewList(πTemp010...).ToObject()
							πTemp001[0] = πTemp003
							πTemp010 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp010...).ToObject()
							πTemp001[1] = πTemp003
							πTemp010 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp010...).ToObject()
							πTemp001[2] = πTemp003
							if πE = πg.CheckLocal(πF, µtout, "tout"); πE != nil {
								continue
							}
							πTemp001[3] = µtout
							if πTemp003, πE = πg.ResolveGlobal(πF, ßselect); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßselect, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp003); πE != nil {
								continue
							}
							µrfd = πTemp004
							µwfd = πTemp005
							µxfd = πTemp006
							if πE = πg.CheckLocal(πF, µrfd, "rfd"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µwfd, "wfd"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µxfd, "xfd"); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(µrfd, µwfd, µxfd).ToObject()
							πTemp001 = make([]*πg.Object, 0)
							πTemp006 = πg.NewList(πTemp001...).ToObject()
							πTemp001 = make([]*πg.Object, 0)
							πTemp007 = πg.NewList(πTemp001...).ToObject()
							πTemp001 = make([]*πg.Object, 0)
							πTemp011 = πg.NewList(πTemp001...).ToObject()
							πTemp005 = πg.NewTuple3(πTemp006, πTemp007, πTemp011).ToObject()
							if πTemp003, πE = πg.Eq(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label6
							}
							goto Label7
							// line 38: if (rfd, wfd, xfd) == ([], [], []):
							πF.SetLineno(38)
						Label6:
							// line 39: continue
							πF.SetLineno(39)
							continue
							goto Label7
						Label7:
							if πE = πg.CheckLocal(πF, µrfd, "rfd"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µwfd, "wfd"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µxfd, "xfd"); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(µrfd, µwfd, µxfd).ToObject()
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp001[0] = µp
							πTemp006 = πg.NewList(πTemp001...).ToObject()
							πTemp001 = make([]*πg.Object, 0)
							πTemp007 = πg.NewList(πTemp001...).ToObject()
							πTemp001 = make([]*πg.Object, 0)
							πTemp011 = πg.NewList(πTemp001...).ToObject()
							πTemp005 = πg.NewTuple3(πTemp006, πTemp007, πTemp011).ToObject()
							if πTemp003, πE = πg.Eq(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label8
							}
							goto Label9
							// line 40: if (rfd, wfd, xfd) == ([p], [], []):
							πF.SetLineno(40)
						Label8:
							// line 41: line = p.readline()
							πF.SetLineno(41)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µp, ßreadline, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µline = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßverbose, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label10
							}
							goto Label11
							// line 42: if test_support.verbose:
							πF.SetLineno(42)
						Label10:
							// line 43: print repr(line)
							πF.SetLineno(43)
							πTemp001 = make([]*πg.Object, 1)
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp010[0] = µline
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp001[0] = πTemp004
							if πE = πg.Print(πF, πTemp001, true); πE != nil {
								continue
							}
							goto Label11
						Label11:
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, µline); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp009).ToObject()
							if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label12
							}
							goto Label13
							// line 44: if not line:
							πF.SetLineno(44)
						Label12:
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßverbose, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label14
							}
							goto Label15
							// line 45: if test_support.verbose:
							πF.SetLineno(45)
						Label14:
							// line 46: print 'EOF'
							πF.SetLineno(46)
							πTemp001 = make([]*πg.Object, 1)
							πTemp001[0] = ßEOF.ToObject()
							if πE = πg.Print(πF, πTemp001, true); πE != nil {
								continue
							}
							goto Label15
						Label15:
							// line 47: break
							πF.SetLineno(47)
							πTemp008 = true
							continue
							goto Label13
						Label13:
							// line 48: continue
							πF.SetLineno(48)
							continue
							goto Label9
						Label9:
							// line 49: self.fail('Unexpected return values from select():', rfd, wfd, xfd)
							πF.SetLineno(49)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewStr("Unexpected return values from select():").ToObject()
							if πE = πg.CheckLocal(πF, µrfd, "rfd"); πE != nil {
								continue
							}
							πTemp001[1] = µrfd
							if πE = πg.CheckLocal(πF, µwfd, "wfd"); πE != nil {
								continue
							}
							πTemp001[2] = µwfd
							if πE = πg.CheckLocal(πF, µxfd, "xfd"); πE != nil {
								continue
							}
							πTemp001[3] = µxfd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
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
							// line 50: p.close()
							πF.SetLineno(50)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µp, ßclose, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_select.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 53: def test_select_mutated(self):
					πF.SetLineno(53)
					πTemp006 = make([]πg.Param, 1)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_select_mutated", "build/src/__python__/test/test_select.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πg.UnboundLocal; _ = µa
						var µF *πg.Object = πg.UnboundLocal; _ = µF
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Dict
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
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
							default: panic("unexpected function state")
							}
							// line 54: a = []
							πF.SetLineno(54)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µa = πTemp002
							// line 55: class F(object):
							πF.SetLineno(55)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							πTemp003 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("F", "build/src/__python__/test/test_select.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp003
								_ = πClass
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 []πg.Param
								_ = πTemp002
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 56: def fileno(self):
									πF.SetLineno(56)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("fileno", "build/src/__python__/test/test_select.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
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
											// line 57: del a[-1]
											πF.SetLineno(57)
											if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
												continue
											}
											if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
												continue
											}
											πTemp001 = πTemp002
											if πE = πg.DelItem(πF, µa, πTemp001); πE != nil {
												continue
											}
											// line 58: return sys.stdout.fileno()
											πF.SetLineno(58)
											if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstdout, nil); πE != nil {
												continue
											}
											if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßfileno, nil); πE != nil {
												continue
											}
											if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
									if πE = πClass.SetItem(πF, ßfileno.ToObject(), πTemp001); πE != nil {
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("F").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
								continue
							}
							µF = πTemp005
							// line 59: a[:] = [F()] * 10
							πF.SetLineno(59)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µF, "F"); πE != nil {
								continue
							}
							if πTemp004, πE = µF.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp004 = πg.NewList(πTemp001...).ToObject()
							if πTemp002, πE = πg.Mul(πF, πTemp004, πg.NewInt(10).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, µa, πTemp005, πTemp004); πE != nil {
								continue
							}
							// line 60: self.assertEqual(select.select([], a, []), ([], a[:5], []))
							πF.SetLineno(60)
							πTemp001 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(3)
							πTemp007 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp007...).ToObject()
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp006[1] = µa
							πTemp007 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp007...).ToObject()
							πTemp006[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßselect); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßselect, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[0] = πTemp002
							πTemp006 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp006...).ToObject()
							if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(5).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µa, πTemp005); πE != nil {
								continue
							}
							πTemp006 = make([]*πg.Object, 0)
							πTemp005 = πg.NewList(πTemp006...).ToObject()
							πTemp002 = πg.NewTuple3(πTemp004, πTemp008, πTemp005).ToObject()
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_select_mutated.ToObject(), πTemp007); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("SelectTestCase").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSelectTestCase.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 62: def test_main():
			πF.SetLineno(62)
			πTemp007 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("test_main", "build/src/__python__/test/test_select.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 63: test_support.run_unittest(SelectTestCase)
					πF.SetLineno(63)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSelectTestCase); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrun_unittest, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 64: test_support.reap_children()
					πF.SetLineno(64)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßreap_children, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßtest_main.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp004, πE = πg.Eq(πF, πTemp005, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
				continue
			}
			if πTemp008 {
				goto Label1
			}
			goto Label2
			// line 66: if __name__ == "__main__":
			πF.SetLineno(66)
		Label1:
			// line 67: test_main()
			πF.SetLineno(67)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßtest_main); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("test.test_select", Code)
}
