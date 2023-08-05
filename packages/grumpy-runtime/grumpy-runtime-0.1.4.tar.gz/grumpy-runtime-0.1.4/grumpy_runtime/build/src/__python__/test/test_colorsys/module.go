package test_colorsys
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/test/test_colorsys.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßColorsysTest := πg.InternStr("ColorsysTest")
		ßTestCase := πg.InternStr("TestCase")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßassertAlmostEqual := πg.InternStr("assertAlmostEqual")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertTripleEqual := πg.InternStr("assertTripleEqual")
		ßcolorsys := πg.InternStr("colorsys")
		ßfrange := πg.InternStr("frange")
		ßhls_to_rgb := πg.InternStr("hls_to_rgb")
		ßhsv_to_rgb := πg.InternStr("hsv_to_rgb")
		ßlen := πg.InternStr("len")
		ßrgb_to_hls := πg.InternStr("rgb_to_hls")
		ßrgb_to_hsv := πg.InternStr("rgb_to_hsv")
		ßrgb_to_yiq := πg.InternStr("rgb_to_yiq")
		ßrun_unittest := πg.InternStr("run_unittest")
		ßtest_hls_roundtrip := πg.InternStr("test_hls_roundtrip")
		ßtest_hls_values := πg.InternStr("test_hls_values")
		ßtest_hsv_roundtrip := πg.InternStr("test_hsv_roundtrip")
		ßtest_hsv_values := πg.InternStr("test_hsv_values")
		ßtest_main := πg.InternStr("test_main")
		ßtest_support := πg.InternStr("test_support")
		ßtest_yiq_roundtrip := πg.InternStr("test_yiq_roundtrip")
		ßtest_yiq_values := πg.InternStr("test_yiq_values")
		ßunittest := πg.InternStr("unittest")
		ßyiq_to_rgb := πg.InternStr("yiq_to_rgb")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 []πg.Param
		_ = πTemp003
		var πTemp004 *πg.Dict
		_ = πTemp004
		var πTemp005 *πg.Object
		_ = πTemp005
		var πTemp006 *πg.Object
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
			// line 1: import unittest
			πF.SetLineno(1)
			if πTemp002, πE = πg.ImportModule(πF, "unittest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 2: import colorsys
			πF.SetLineno(2)
			if πTemp002, πE = πg.ImportModule(πF, "colorsys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßcolorsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 3: from test import test_support
			πF.SetLineno(3)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßtest_support.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 5: def frange(start, stop, step):
			πF.SetLineno(5)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "start", Def: nil}
			πTemp003[1] = πg.Param{Name: "stop", Def: nil}
			πTemp003[2] = πg.Param{Name: "step", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("frange", "build/src/__python__/test/test_colorsys.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µstart *πg.Object = πArgs[0]; _ = µstart
				var µstop *πg.Object = πArgs[1]; _ = µstop
				var µstep *πg.Object = πArgs[2]; _ = µstep
				var πTemp001 bool
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 1: goto Label1
						case 2: goto Label2
						case 4: goto Label4
						default: panic("unexpected function state")
						}
						// line 6: while start <= stop:
						πF.SetLineno(6)
						πF.PushCheckpoint(2)
						πTemp001 = false
					Label1:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp001 {
							πF.PopCheckpoint()
							goto Label3
						}
						if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µstop, "stop"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.LE(πF, µstart, µstop); πE != nil {
							continue
						}
						if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
							continue
						}
						if πE != nil || !πTemp002 {
							continue
						}
						πF.PushCheckpoint(1)            
						// line 7: yield start
						πF.SetLineno(7)
						if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
							continue
						}
						πF.PushCheckpoint(4)
						return µstart, nil
					Label4:
						πTemp003 = πSent
						// line 8: start += step
						πF.SetLineno(8)
						if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µstep, "step"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.IAdd(πF, µstart, µstep); πE != nil {
							continue
						}
						µstart = πTemp003
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
			if πE = πF.Globals().SetItem(πF, ßfrange.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 10: class ColorsysTest(unittest.TestCase):
			πF.SetLineno(10)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßTestCase, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp004 = πg.NewDict()
			if πTemp005, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp005); πE != nil {
				continue
			}
			_, πE = πg.NewCode("ColorsysTest", "build/src/__python__/test/test_colorsys.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
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
					// line 12: def assertTripleEqual(self, tr1, tr2):
					πF.SetLineno(12)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "tr1", Def: nil}
					πTemp002[2] = πg.Param{Name: "tr2", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("assertTripleEqual", "build/src/__python__/test/test_colorsys.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtr1 *πg.Object = πArgs[1]; _ = µtr1
						var µtr2 *πg.Object = πArgs[2]; _ = µtr2
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
							// line 13: self.assertEqual(len(tr1), 3)
							πF.SetLineno(13)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtr1, "tr1"); πE != nil {
								continue
							}
							πTemp002[0] = µtr1
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 14: self.assertEqual(len(tr2), 3)
							πF.SetLineno(14)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtr2, "tr2"); πE != nil {
								continue
							}
							πTemp002[0] = µtr2
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 15: self.assertAlmostEqual(tr1[0], tr2[0])
							πF.SetLineno(15)
							πTemp001 = πF.MakeArgs(2)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µtr1, "tr1"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µtr1, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µtr2, "tr2"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µtr2, πTemp003); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertAlmostEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 16: self.assertAlmostEqual(tr1[1], tr2[1])
							πF.SetLineno(16)
							πTemp001 = πF.MakeArgs(2)
							πTemp003 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µtr1, "tr1"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µtr1, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp003 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µtr2, "tr2"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µtr2, πTemp003); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertAlmostEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 17: self.assertAlmostEqual(tr1[2], tr2[2])
							πF.SetLineno(17)
							πTemp001 = πF.MakeArgs(2)
							πTemp003 = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µtr1, "tr1"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µtr1, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp003 = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µtr2, "tr2"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µtr2, πTemp003); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertAlmostEqual, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßassertTripleEqual.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 19: def test_hsv_roundtrip(self):
					πF.SetLineno(19)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_hsv_roundtrip", "build/src/__python__/test/test_colorsys.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µr *πg.Object = πg.UnboundLocal; _ = µr
						var µg *πg.Object = πg.UnboundLocal; _ = µg
						var µb *πg.Object = πg.UnboundLocal; _ = µb
						var µrgb *πg.Object = πg.UnboundLocal; _ = µrgb
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
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 bool
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
							case 4: goto Label4
							case 5: goto Label5
							case 7: goto Label7
							case 8: goto Label8
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(3)
							πTemp002[0] = πg.NewFloat(0.0).ToObject()
							πTemp002[1] = πg.NewFloat(1.0).ToObject()
							πTemp002[2] = πg.NewFloat(0.2).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfrange); πE != nil {
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
								µr = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							πTemp002 = πF.MakeArgs(3)
							πTemp002[0] = πg.NewFloat(0.0).ToObject()
							πTemp002[1] = πg.NewFloat(1.0).ToObject()
							πTemp002[2] = πg.NewFloat(0.2).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßfrange); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.Iter(πF, πTemp007); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp006 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label6
							}
							if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
								µg = πTemp004
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(4)            
							πTemp002 = πF.MakeArgs(3)
							πTemp002[0] = πg.NewFloat(0.0).ToObject()
							πTemp002[1] = πg.NewFloat(1.0).ToObject()
							πTemp002[2] = πg.NewFloat(0.2).ToObject()
							if πTemp007, πE = πg.ResolveGlobal(πF, ßfrange); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp004, πE = πg.Iter(πF, πTemp009); πE != nil {
								continue
							}
							πF.PushCheckpoint(8)
							πTemp008 = false
						Label7:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label9
							}
							if πTemp007, πE = πg.Next(πF, πTemp004); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp010 = !isStop
							} else {
								πTemp010 = true
								µb = πTemp007
							}
							if πE != nil || !πTemp010 {
								continue
							}
							πF.PushCheckpoint(7)            
							// line 23: rgb = (r, g, b)
							πF.SetLineno(23)
							if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							πTemp007 = πg.NewTuple3(µr, µg, µb).ToObject()
							µrgb = πTemp007
							// line 24: self.assertTripleEqual(
							πF.SetLineno(24)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µrgb, "rgb"); πE != nil {
								continue
							}
							πTemp002[0] = µrgb
							if πE = πg.CheckLocal(πF, µrgb, "rgb"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßcolorsys); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßrgb_to_hsv, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Invoke(πF, πTemp009, nil, µrgb, nil, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.ResolveGlobal(πF, ßcolorsys); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, πTemp009, ßhsv_to_rgb, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Invoke(πF, πTemp011, nil, πTemp007, nil, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp009
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßassertTripleEqual, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label8:
							if πE != nil || πR != nil {
								continue
							}
						Label9:
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßtest_hsv_roundtrip.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 29: def test_hsv_values(self):
					πF.SetLineno(29)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_hsv_values", "build/src/__python__/test/test_colorsys.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µvalues *πg.Object = πg.UnboundLocal; _ = µvalues
						var µrgb *πg.Object = πg.UnboundLocal; _ = µrgb
						var µhsv *πg.Object = πg.UnboundLocal; _ = µhsv
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
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
							// line 30: values = [
							πF.SetLineno(30)
							πTemp001 = make([]*πg.Object, 9)
							πTemp003 = πg.NewTuple3(πg.NewFloat(0.0).ToObject(), πg.NewFloat(0.0).ToObject(), πg.NewFloat(0.0).ToObject()).ToObject()
							πTemp004 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewFloat(0.0).ToObject(), πg.NewFloat(0.0).ToObject()).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[0] = πTemp002
							πTemp003 = πg.NewTuple3(πg.NewFloat(0.0).ToObject(), πg.NewFloat(0.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
							if πTemp005, πE = πg.Div(πF, πg.NewFloat(4.0).ToObject(), πg.NewFloat(6.0).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πTemp005, πg.NewFloat(1.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[1] = πTemp002
							πTemp003 = πg.NewTuple3(πg.NewFloat(0.0).ToObject(), πg.NewFloat(1.0).ToObject(), πg.NewFloat(0.0).ToObject()).ToObject()
							if πTemp005, πE = πg.Div(πF, πg.NewFloat(2.0).ToObject(), πg.NewFloat(6.0).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πTemp005, πg.NewFloat(1.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[2] = πTemp002
							πTemp003 = πg.NewTuple3(πg.NewFloat(0.0).ToObject(), πg.NewFloat(1.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
							if πTemp005, πE = πg.Div(πF, πg.NewFloat(3.0).ToObject(), πg.NewFloat(6.0).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πTemp005, πg.NewFloat(1.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[3] = πTemp002
							πTemp003 = πg.NewTuple3(πg.NewFloat(1.0).ToObject(), πg.NewFloat(0.0).ToObject(), πg.NewFloat(0.0).ToObject()).ToObject()
							πTemp004 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewFloat(1.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[4] = πTemp002
							πTemp003 = πg.NewTuple3(πg.NewFloat(1.0).ToObject(), πg.NewFloat(0.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
							if πTemp005, πE = πg.Div(πF, πg.NewFloat(5.0).ToObject(), πg.NewFloat(6.0).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πTemp005, πg.NewFloat(1.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[5] = πTemp002
							πTemp003 = πg.NewTuple3(πg.NewFloat(1.0).ToObject(), πg.NewFloat(1.0).ToObject(), πg.NewFloat(0.0).ToObject()).ToObject()
							if πTemp005, πE = πg.Div(πF, πg.NewFloat(1.0).ToObject(), πg.NewFloat(6.0).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πTemp005, πg.NewFloat(1.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[6] = πTemp002
							πTemp003 = πg.NewTuple3(πg.NewFloat(1.0).ToObject(), πg.NewFloat(1.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
							πTemp004 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewFloat(0.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[7] = πTemp002
							πTemp003 = πg.NewTuple3(πg.NewFloat(0.5).ToObject(), πg.NewFloat(0.5).ToObject(), πg.NewFloat(0.5).ToObject()).ToObject()
							πTemp004 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewFloat(0.0).ToObject(), πg.NewFloat(0.5).ToObject()).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[8] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µvalues = πTemp002
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µvalues); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp006 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
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
								πTemp007 = !isStop
							} else {
								πTemp007 = true
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
									continue
								}
								µrgb = πTemp004
								µhsv = πTemp005
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 43: self.assertTripleEqual(hsv, colorsys.rgb_to_hsv(*rgb))
							πF.SetLineno(43)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µhsv, "hsv"); πE != nil {
								continue
							}
							πTemp001[0] = µhsv
							if πE = πg.CheckLocal(πF, µrgb, "rgb"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßcolorsys); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßrgb_to_hsv, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp004, nil, µrgb, nil, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertTripleEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 44: self.assertTripleEqual(rgb, colorsys.hsv_to_rgb(*hsv))
							πF.SetLineno(44)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µrgb, "rgb"); πE != nil {
								continue
							}
							πTemp001[0] = µrgb
							if πE = πg.CheckLocal(πF, µhsv, "hsv"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßcolorsys); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßhsv_to_rgb, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp004, nil, µhsv, nil, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertTripleEqual, nil); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_hsv_values.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 46: def test_hls_roundtrip(self):
					πF.SetLineno(46)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_hls_roundtrip", "build/src/__python__/test/test_colorsys.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µr *πg.Object = πg.UnboundLocal; _ = µr
						var µg *πg.Object = πg.UnboundLocal; _ = µg
						var µb *πg.Object = πg.UnboundLocal; _ = µb
						var µrgb *πg.Object = πg.UnboundLocal; _ = µrgb
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
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 bool
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
							case 4: goto Label4
							case 5: goto Label5
							case 7: goto Label7
							case 8: goto Label8
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(3)
							πTemp002[0] = πg.NewFloat(0.0).ToObject()
							πTemp002[1] = πg.NewFloat(1.0).ToObject()
							πTemp002[2] = πg.NewFloat(0.2).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfrange); πE != nil {
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
								µr = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							πTemp002 = πF.MakeArgs(3)
							πTemp002[0] = πg.NewFloat(0.0).ToObject()
							πTemp002[1] = πg.NewFloat(1.0).ToObject()
							πTemp002[2] = πg.NewFloat(0.2).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßfrange); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.Iter(πF, πTemp007); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp006 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label6
							}
							if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
								µg = πTemp004
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(4)            
							πTemp002 = πF.MakeArgs(3)
							πTemp002[0] = πg.NewFloat(0.0).ToObject()
							πTemp002[1] = πg.NewFloat(1.0).ToObject()
							πTemp002[2] = πg.NewFloat(0.2).ToObject()
							if πTemp007, πE = πg.ResolveGlobal(πF, ßfrange); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp004, πE = πg.Iter(πF, πTemp009); πE != nil {
								continue
							}
							πF.PushCheckpoint(8)
							πTemp008 = false
						Label7:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label9
							}
							if πTemp007, πE = πg.Next(πF, πTemp004); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp010 = !isStop
							} else {
								πTemp010 = true
								µb = πTemp007
							}
							if πE != nil || !πTemp010 {
								continue
							}
							πF.PushCheckpoint(7)            
							// line 50: rgb = (r, g, b)
							πF.SetLineno(50)
							if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							πTemp007 = πg.NewTuple3(µr, µg, µb).ToObject()
							µrgb = πTemp007
							// line 51: self.assertTripleEqual(
							πF.SetLineno(51)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µrgb, "rgb"); πE != nil {
								continue
							}
							πTemp002[0] = µrgb
							if πE = πg.CheckLocal(πF, µrgb, "rgb"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßcolorsys); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßrgb_to_hls, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Invoke(πF, πTemp009, nil, µrgb, nil, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.ResolveGlobal(πF, ßcolorsys); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, πTemp009, ßhls_to_rgb, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Invoke(πF, πTemp011, nil, πTemp007, nil, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp009
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßassertTripleEqual, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label8:
							if πE != nil || πR != nil {
								continue
							}
						Label9:
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßtest_hls_roundtrip.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 56: def test_hls_values(self):
					πF.SetLineno(56)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_hls_values", "build/src/__python__/test/test_colorsys.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µvalues *πg.Object = πg.UnboundLocal; _ = µvalues
						var µrgb *πg.Object = πg.UnboundLocal; _ = µrgb
						var µhls *πg.Object = πg.UnboundLocal; _ = µhls
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
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
							// line 57: values = [
							πF.SetLineno(57)
							πTemp001 = make([]*πg.Object, 9)
							πTemp003 = πg.NewTuple3(πg.NewFloat(0.0).ToObject(), πg.NewFloat(0.0).ToObject(), πg.NewFloat(0.0).ToObject()).ToObject()
							πTemp004 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewFloat(0.0).ToObject(), πg.NewFloat(0.0).ToObject()).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[0] = πTemp002
							πTemp003 = πg.NewTuple3(πg.NewFloat(0.0).ToObject(), πg.NewFloat(0.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
							if πTemp005, πE = πg.Div(πF, πg.NewFloat(4.0).ToObject(), πg.NewFloat(6.0).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πTemp005, πg.NewFloat(0.5).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[1] = πTemp002
							πTemp003 = πg.NewTuple3(πg.NewFloat(0.0).ToObject(), πg.NewFloat(1.0).ToObject(), πg.NewFloat(0.0).ToObject()).ToObject()
							if πTemp005, πE = πg.Div(πF, πg.NewFloat(2.0).ToObject(), πg.NewFloat(6.0).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πTemp005, πg.NewFloat(0.5).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[2] = πTemp002
							πTemp003 = πg.NewTuple3(πg.NewFloat(0.0).ToObject(), πg.NewFloat(1.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
							if πTemp005, πE = πg.Div(πF, πg.NewFloat(3.0).ToObject(), πg.NewFloat(6.0).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πTemp005, πg.NewFloat(0.5).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[3] = πTemp002
							πTemp003 = πg.NewTuple3(πg.NewFloat(1.0).ToObject(), πg.NewFloat(0.0).ToObject(), πg.NewFloat(0.0).ToObject()).ToObject()
							πTemp004 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewFloat(0.5).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[4] = πTemp002
							πTemp003 = πg.NewTuple3(πg.NewFloat(1.0).ToObject(), πg.NewFloat(0.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
							if πTemp005, πE = πg.Div(πF, πg.NewFloat(5.0).ToObject(), πg.NewFloat(6.0).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πTemp005, πg.NewFloat(0.5).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[5] = πTemp002
							πTemp003 = πg.NewTuple3(πg.NewFloat(1.0).ToObject(), πg.NewFloat(1.0).ToObject(), πg.NewFloat(0.0).ToObject()).ToObject()
							if πTemp005, πE = πg.Div(πF, πg.NewFloat(1.0).ToObject(), πg.NewFloat(6.0).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πTemp005, πg.NewFloat(0.5).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[6] = πTemp002
							πTemp003 = πg.NewTuple3(πg.NewFloat(1.0).ToObject(), πg.NewFloat(1.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
							πTemp004 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewFloat(1.0).ToObject(), πg.NewFloat(0.0).ToObject()).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[7] = πTemp002
							πTemp003 = πg.NewTuple3(πg.NewFloat(0.5).ToObject(), πg.NewFloat(0.5).ToObject(), πg.NewFloat(0.5).ToObject()).ToObject()
							πTemp004 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewFloat(0.5).ToObject(), πg.NewFloat(0.0).ToObject()).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[8] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µvalues = πTemp002
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µvalues); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp006 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
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
								πTemp007 = !isStop
							} else {
								πTemp007 = true
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
									continue
								}
								µrgb = πTemp004
								µhls = πTemp005
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 70: self.assertTripleEqual(hls, colorsys.rgb_to_hls(*rgb))
							πF.SetLineno(70)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µhls, "hls"); πE != nil {
								continue
							}
							πTemp001[0] = µhls
							if πE = πg.CheckLocal(πF, µrgb, "rgb"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßcolorsys); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßrgb_to_hls, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp004, nil, µrgb, nil, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertTripleEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 71: self.assertTripleEqual(rgb, colorsys.hls_to_rgb(*hls))
							πF.SetLineno(71)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µrgb, "rgb"); πE != nil {
								continue
							}
							πTemp001[0] = µrgb
							if πE = πg.CheckLocal(πF, µhls, "hls"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßcolorsys); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßhls_to_rgb, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp004, nil, µhls, nil, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertTripleEqual, nil); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_hls_values.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 73: def test_yiq_roundtrip(self):
					πF.SetLineno(73)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_yiq_roundtrip", "build/src/__python__/test/test_colorsys.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µr *πg.Object = πg.UnboundLocal; _ = µr
						var µg *πg.Object = πg.UnboundLocal; _ = µg
						var µb *πg.Object = πg.UnboundLocal; _ = µb
						var µrgb *πg.Object = πg.UnboundLocal; _ = µrgb
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
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 bool
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
							case 4: goto Label4
							case 5: goto Label5
							case 7: goto Label7
							case 8: goto Label8
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(3)
							πTemp002[0] = πg.NewFloat(0.0).ToObject()
							πTemp002[1] = πg.NewFloat(1.0).ToObject()
							πTemp002[2] = πg.NewFloat(0.2).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfrange); πE != nil {
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
								µr = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							πTemp002 = πF.MakeArgs(3)
							πTemp002[0] = πg.NewFloat(0.0).ToObject()
							πTemp002[1] = πg.NewFloat(1.0).ToObject()
							πTemp002[2] = πg.NewFloat(0.2).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßfrange); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.Iter(πF, πTemp007); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp006 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label6
							}
							if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
								µg = πTemp004
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(4)            
							πTemp002 = πF.MakeArgs(3)
							πTemp002[0] = πg.NewFloat(0.0).ToObject()
							πTemp002[1] = πg.NewFloat(1.0).ToObject()
							πTemp002[2] = πg.NewFloat(0.2).ToObject()
							if πTemp007, πE = πg.ResolveGlobal(πF, ßfrange); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp004, πE = πg.Iter(πF, πTemp009); πE != nil {
								continue
							}
							πF.PushCheckpoint(8)
							πTemp008 = false
						Label7:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label9
							}
							if πTemp007, πE = πg.Next(πF, πTemp004); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp010 = !isStop
							} else {
								πTemp010 = true
								µb = πTemp007
							}
							if πE != nil || !πTemp010 {
								continue
							}
							πF.PushCheckpoint(7)            
							// line 77: rgb = (r, g, b)
							πF.SetLineno(77)
							if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							πTemp007 = πg.NewTuple3(µr, µg, µb).ToObject()
							µrgb = πTemp007
							// line 78: self.assertTripleEqual(
							πF.SetLineno(78)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µrgb, "rgb"); πE != nil {
								continue
							}
							πTemp002[0] = µrgb
							if πE = πg.CheckLocal(πF, µrgb, "rgb"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßcolorsys); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßrgb_to_yiq, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Invoke(πF, πTemp009, nil, µrgb, nil, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.ResolveGlobal(πF, ßcolorsys); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, πTemp009, ßyiq_to_rgb, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Invoke(πF, πTemp011, nil, πTemp007, nil, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp009
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßassertTripleEqual, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label8:
							if πE != nil || πR != nil {
								continue
							}
						Label9:
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßtest_yiq_roundtrip.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 83: def test_yiq_values(self):
					πF.SetLineno(83)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("test_yiq_values", "build/src/__python__/test/test_colorsys.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µvalues *πg.Object = πg.UnboundLocal; _ = µvalues
						var µrgb *πg.Object = πg.UnboundLocal; _ = µrgb
						var µyiq *πg.Object = πg.UnboundLocal; _ = µyiq
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
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 bool
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
							// line 84: values = [
							πF.SetLineno(84)
							πTemp001 = make([]*πg.Object, 9)
							πTemp003 = πg.NewTuple3(πg.NewFloat(0.0).ToObject(), πg.NewFloat(0.0).ToObject(), πg.NewFloat(0.0).ToObject()).ToObject()
							πTemp004 = πg.NewTuple3(πg.NewFloat(0.0).ToObject(), πg.NewFloat(0.0).ToObject(), πg.NewFloat(0.0).ToObject()).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[0] = πTemp002
							πTemp003 = πg.NewTuple3(πg.NewFloat(0.0).ToObject(), πg.NewFloat(0.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
							if πTemp005, πE = πg.Neg(πF, πg.NewFloat(0.3217).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πg.NewFloat(0.11).ToObject(), πTemp005, πg.NewFloat(0.3121).ToObject()).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[1] = πTemp002
							πTemp003 = πg.NewTuple3(πg.NewFloat(0.0).ToObject(), πg.NewFloat(1.0).ToObject(), πg.NewFloat(0.0).ToObject()).ToObject()
							if πTemp005, πE = πg.Neg(πF, πg.NewFloat(0.2773).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewFloat(0.5251).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πg.NewFloat(0.59).ToObject(), πTemp005, πTemp006).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[2] = πTemp002
							πTemp003 = πg.NewTuple3(πg.NewFloat(0.0).ToObject(), πg.NewFloat(1.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
							if πTemp005, πE = πg.Neg(πF, πg.NewFloat(0.599).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewFloat(0.213).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πg.NewFloat(0.7).ToObject(), πTemp005, πTemp006).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[3] = πTemp002
							πTemp003 = πg.NewTuple3(πg.NewFloat(1.0).ToObject(), πg.NewFloat(0.0).ToObject(), πg.NewFloat(0.0).ToObject()).ToObject()
							πTemp004 = πg.NewTuple3(πg.NewFloat(0.3).ToObject(), πg.NewFloat(0.599).ToObject(), πg.NewFloat(0.213).ToObject()).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[4] = πTemp002
							πTemp003 = πg.NewTuple3(πg.NewFloat(1.0).ToObject(), πg.NewFloat(0.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
							πTemp004 = πg.NewTuple3(πg.NewFloat(0.41).ToObject(), πg.NewFloat(0.2773).ToObject(), πg.NewFloat(0.5251).ToObject()).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[5] = πTemp002
							πTemp003 = πg.NewTuple3(πg.NewFloat(1.0).ToObject(), πg.NewFloat(1.0).ToObject(), πg.NewFloat(0.0).ToObject()).ToObject()
							if πTemp005, πE = πg.Neg(πF, πg.NewFloat(0.3121).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πg.NewFloat(0.89).ToObject(), πg.NewFloat(0.3217).ToObject(), πTemp005).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[6] = πTemp002
							πTemp003 = πg.NewTuple3(πg.NewFloat(1.0).ToObject(), πg.NewFloat(1.0).ToObject(), πg.NewFloat(1.0).ToObject()).ToObject()
							πTemp004 = πg.NewTuple3(πg.NewFloat(1.0).ToObject(), πg.NewFloat(0.0).ToObject(), πg.NewFloat(0.0).ToObject()).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[7] = πTemp002
							πTemp003 = πg.NewTuple3(πg.NewFloat(0.5).ToObject(), πg.NewFloat(0.5).ToObject(), πg.NewFloat(0.5).ToObject()).ToObject()
							πTemp004 = πg.NewTuple3(πg.NewFloat(0.5).ToObject(), πg.NewFloat(0.0).ToObject(), πg.NewFloat(0.0).ToObject()).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[8] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µvalues = πTemp002
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µvalues); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp007 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
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
								πTemp008 = !isStop
							} else {
								πTemp008 = true
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
									continue
								}
								µrgb = πTemp004
								µyiq = πTemp005
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 97: self.assertTripleEqual(yiq, colorsys.rgb_to_yiq(*rgb))
							πF.SetLineno(97)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µyiq, "yiq"); πE != nil {
								continue
							}
							πTemp001[0] = µyiq
							if πE = πg.CheckLocal(πF, µrgb, "rgb"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßcolorsys); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßrgb_to_yiq, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp004, nil, µrgb, nil, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertTripleEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 98: self.assertTripleEqual(rgb, colorsys.yiq_to_rgb(*yiq))
							πF.SetLineno(98)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µrgb, "rgb"); πE != nil {
								continue
							}
							πTemp001[0] = µrgb
							if πE = πg.CheckLocal(πF, µyiq, "yiq"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßcolorsys); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßyiq_to_rgb, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp004, nil, µyiq, nil, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertTripleEqual, nil); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_yiq_values.ToObject(), πTemp008); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("ColorsysTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßColorsysTest.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 100: def test_main():
			πF.SetLineno(100)
			πTemp003 = make([]πg.Param, 0)
			πTemp005 = πg.NewFunction(πg.NewCode("test_main", "build/src/__python__/test/test_colorsys.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 101: test_support.run_unittest(ColorsysTest)
					πF.SetLineno(101)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßColorsysTest); πE != nil {
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
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßtest_main.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp006, πE = πg.Eq(πF, πTemp007, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp009, πE = πg.IsTrue(πF, πTemp006); πE != nil {
				continue
			}
			if πTemp009 {
				goto Label1
			}
			goto Label2
			// line 103: if __name__ == "__main__":
			πF.SetLineno(103)
		Label1:
			// line 104: test_main()
			πF.SetLineno(104)
			if πTemp006, πE = πg.ResolveGlobal(πF, ßtest_main); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("test.test_colorsys", Code)
}
