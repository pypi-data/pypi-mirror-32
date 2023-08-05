package unittest
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/unittest/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßBaseTestSuite := πg.InternStr("BaseTestSuite")
		ßFunctionTestCase := πg.InternStr("FunctionTestCase")
		ßSkipTest := πg.InternStr("SkipTest")
		ßTestCase := πg.InternStr("TestCase")
		ßTestLoader := πg.InternStr("TestLoader")
		ßTestResult := πg.InternStr("TestResult")
		ßTestSuite := πg.InternStr("TestSuite")
		ßTextTestResult := πg.InternStr("TextTestResult")
		ßTextTestRunner := πg.InternStr("TextTestRunner")
		ßTrue := πg.InternStr("True")
		ß_TextTestResult := πg.InternStr("_TextTestResult")
		ß__all__ := πg.InternStr("__all__")
		ß__unittest := πg.InternStr("__unittest")
		ßdefaultTestLoader := πg.InternStr("defaultTestLoader")
		ßexpectedFailure := πg.InternStr("expectedFailure")
		ßfindTestCases := πg.InternStr("findTestCases")
		ßgetTestCaseNames := πg.InternStr("getTestCaseNames")
		ßinstallHandler := πg.InternStr("installHandler")
		ßmakeSuite := πg.InternStr("makeSuite")
		ßregisterResult := πg.InternStr("registerResult")
		ßremoveHandler := πg.InternStr("removeHandler")
		ßremoveResult := πg.InternStr("removeResult")
		ßskip := πg.InternStr("skip")
		ßskipIf := πg.InternStr("skipIf")
		ßskipUnless := πg.InternStr("skipUnless")
		ßunittest_case := πg.InternStr("unittest_case")
		ßunittest_loader := πg.InternStr("unittest_loader")
		ßunittest_result := πg.InternStr("unittest_result")
		ßunittest_runner := πg.InternStr("unittest_runner")
		ßunittest_signals := πg.InternStr("unittest_signals")
		ßunittest_suite := πg.InternStr("unittest_suite")
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
		var πTemp008 *πg.Object
		_ = πTemp008
		var πTemp009 *πg.Object
		_ = πTemp009
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """
			πF.SetLineno(1)
			// line 47: __all__ = ['TestResult', 'TestCase', 'TestSuite',
			πF.SetLineno(47)
			πTemp001 = make([]*πg.Object, 17)
			πTemp001[0] = ßTestResult.ToObject()
			πTemp001[1] = ßTestCase.ToObject()
			πTemp001[2] = ßTestSuite.ToObject()
			πTemp001[3] = ßTextTestRunner.ToObject()
			πTemp001[4] = ßTestLoader.ToObject()
			πTemp001[5] = ßFunctionTestCase.ToObject()
			πTemp001[6] = ßdefaultTestLoader.ToObject()
			πTemp001[7] = ßSkipTest.ToObject()
			πTemp001[8] = ßskip.ToObject()
			πTemp001[9] = ßskipIf.ToObject()
			πTemp001[10] = ßskipUnless.ToObject()
			πTemp001[11] = ßexpectedFailure.ToObject()
			πTemp001[12] = ßTextTestResult.ToObject()
			πTemp001[13] = ßinstallHandler.ToObject()
			πTemp001[14] = ßregisterResult.ToObject()
			πTemp001[15] = ßremoveResult.ToObject()
			πTemp001[16] = ßremoveHandler.ToObject()
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 55: __all__ += (['getTestCaseNames', 'makeSuite', 'findTestCases'])
			πF.SetLineno(55)
			if πTemp002, πE = πg.ResolveGlobal(πF, ß__all__); πE != nil {
				continue
			}
			πTemp001 = make([]*πg.Object, 3)
			πTemp001[0] = ßgetTestCaseNames.ToObject()
			πTemp001[1] = ßmakeSuite.ToObject()
			πTemp001[2] = ßfindTestCases.ToObject()
			πTemp003 = πg.NewList(πTemp001...).ToObject()
			if πTemp004, πE = πg.IAdd(πF, πTemp002, πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 57: __unittest = True
			πF.SetLineno(57)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__unittest.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 59: import unittest_result
			πF.SetLineno(59)
			if πTemp001, πE = πg.ImportModule(πF, "unittest_result"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßunittest_result.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 60: import unittest_case
			πF.SetLineno(60)
			if πTemp001, πE = πg.ImportModule(πF, "unittest_case"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßunittest_case.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 61: import unittest_suite
			πF.SetLineno(61)
			if πTemp001, πE = πg.ImportModule(πF, "unittest_suite"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßunittest_suite.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 62: import unittest_loader
			πF.SetLineno(62)
			if πTemp001, πE = πg.ImportModule(πF, "unittest_loader"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßunittest_loader.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 64: import unittest_runner
			πF.SetLineno(64)
			if πTemp001, πE = πg.ImportModule(πF, "unittest_runner"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßunittest_runner.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 65: import unittest_signals
			πF.SetLineno(65)
			if πTemp001, πE = πg.ImportModule(πF, "unittest_signals"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßunittest_signals.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 77: TestResult = unittest_result.TestResult
			πF.SetLineno(77)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßunittest_result); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßTestResult, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTestResult.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 78: TestCase, FunctionTestCase, SkipTest, skip, skipIf, skipUnless, expectedFailure = \
			πF.SetLineno(78)
			πTemp001 = make([]*πg.Object, 7)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßunittest_case); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTestCase, nil); πE != nil {
				continue
			}
			πTemp001[0] = πTemp004
			if πTemp003, πE = πg.ResolveGlobal(πF, ßunittest_case); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßFunctionTestCase, nil); πE != nil {
				continue
			}
			πTemp001[1] = πTemp004
			if πTemp003, πE = πg.ResolveGlobal(πF, ßunittest_case); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßSkipTest, nil); πE != nil {
				continue
			}
			πTemp001[2] = πTemp004
			if πTemp003, πE = πg.ResolveGlobal(πF, ßunittest_case); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßskip, nil); πE != nil {
				continue
			}
			πTemp001[3] = πTemp004
			if πTemp003, πE = πg.ResolveGlobal(πF, ßunittest_case); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßskipIf, nil); πE != nil {
				continue
			}
			πTemp001[4] = πTemp004
			if πTemp003, πE = πg.ResolveGlobal(πF, ßunittest_case); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßskipUnless, nil); πE != nil {
				continue
			}
			πTemp001[5] = πTemp004
			if πTemp003, πE = πg.ResolveGlobal(πF, ßunittest_case); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßexpectedFailure, nil); πE != nil {
				continue
			}
			πTemp001[6] = πTemp004
			πTemp002 = πg.NewTuple(πTemp001...).ToObject()
			if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}}}, πTemp002); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTestCase.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFunctionTestCase.ToObject(), πTemp004); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSkipTest.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßskip.ToObject(), πTemp006); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßskipIf.ToObject(), πTemp007); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßskipUnless.ToObject(), πTemp008); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßexpectedFailure.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 82: BaseTestSuite, TestSuite = unittest_suite.BaseTestSuite, unittest_suite.TestSuite
			πF.SetLineno(82)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßunittest_suite); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßBaseTestSuite, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßunittest_suite); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßTestSuite, nil); πE != nil {
				continue
			}
			πTemp002 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
			if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBaseTestSuite.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTestSuite.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 83: TestLoader, defaultTestLoader, makeSuite, getTestCaseNames, findTestCases = \
			πF.SetLineno(83)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßunittest_loader); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTestLoader, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßunittest_loader); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßdefaultTestLoader, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßunittest_loader); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßmakeSuite, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßunittest_loader); πE != nil {
				continue
			}
			if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßgetTestCaseNames, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßunittest_loader); πE != nil {
				continue
			}
			if πTemp008, πE = πg.GetAttr(πF, πTemp003, ßfindTestCases, nil); πE != nil {
				continue
			}
			πTemp002 = πg.NewTuple5(πTemp004, πTemp005, πTemp006, πTemp007, πTemp008).ToObject()
			if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp002); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTestLoader.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdefaultTestLoader.ToObject(), πTemp004); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmakeSuite.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßgetTestCaseNames.ToObject(), πTemp006); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßfindTestCases.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 87: TextTestRunner, TextTestResult = unittest_runner.TextTestRunner, unittest_runner.TextTestResult
			πF.SetLineno(87)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßunittest_runner); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTextTestRunner, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßunittest_runner); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßTextTestResult, nil); πE != nil {
				continue
			}
			πTemp002 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
			if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTextTestRunner.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTextTestResult.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 88: installHandler, registerResult, removeResult, removeHandler = \
			πF.SetLineno(88)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßunittest_signals); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßinstallHandler, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßunittest_signals); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßregisterResult, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßunittest_signals); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßremoveResult, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßunittest_signals); πE != nil {
				continue
			}
			if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßremoveHandler, nil); πE != nil {
				continue
			}
			πTemp002 = πg.NewTuple4(πTemp004, πTemp005, πTemp006, πTemp007).ToObject()
			if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßinstallHandler.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßregisterResult.ToObject(), πTemp004); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßremoveResult.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßremoveHandler.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 93: _TextTestResult = TextTestResult
			πF.SetLineno(93)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßTextTestResult); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_TextTestResult.ToObject(), πTemp002); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("unittest", Code)
}
