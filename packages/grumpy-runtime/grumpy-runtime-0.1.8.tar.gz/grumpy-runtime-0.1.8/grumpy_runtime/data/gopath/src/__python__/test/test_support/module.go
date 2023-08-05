package test_support
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/test/test_support.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßAssertionError := πg.InternStr("AssertionError")
		ßAttributeError := πg.InternStr("AttributeError")
		ßBasicTestRunner := πg.InternStr("BasicTestRunner")
		ßCleanImport := πg.InternStr("CleanImport")
		ßDeprecationWarning := πg.InternStr("DeprecationWarning")
		ßDictMixin := πg.InternStr("DictMixin")
		ßEnvironmentVarGuard := πg.InternStr("EnvironmentVarGuard")
		ßError := πg.InternStr("Error")
		ßException := πg.InternStr("Exception")
		ßFS_NONASCII := πg.InternStr("FS_NONASCII")
		ßFalse := πg.InternStr("False")
		ßI := πg.InternStr("I")
		ßImportError := πg.InternStr("ImportError")
		ßMAX_Py_ssize_t := πg.InternStr("MAX_Py_ssize_t")
		ßNameError := πg.InternStr("NameError")
		ßNone := πg.InternStr("None")
		ßOSError := πg.InternStr("OSError")
		ßRuntimeWarning := πg.InternStr("RuntimeWarning")
		ßTESTFN := πg.InternStr("TESTFN")
		ßTestCase := πg.InternStr("TestCase")
		ßTestFailed := πg.InternStr("TestFailed")
		ßTestResult := πg.InternStr("TestResult")
		ßTestSuite := πg.InternStr("TestSuite")
		ßTextTestRunner := πg.InternStr("TextTestRunner")
		ßTrue := πg.InternStr("True")
		ßUserDict := πg.InternStr("UserDict")
		ßValueError := πg.InternStr("ValueError")
		ßWNOHANG := πg.InternStr("WNOHANG")
		ßWarning := πg.InternStr("Warning")
		ßWarningMessage := πg.InternStr("WarningMessage")
		ßWarningsRecorder := πg.InternStr("WarningsRecorder")
		ß_WARNING_DETAILS := πg.InternStr("_WARNING_DETAILS")
		ß__all__ := πg.InternStr("__all__")
		ß__class__ := πg.InternStr("__class__")
		ß__delitem__ := πg.InternStr("__delitem__")
		ß__enter__ := πg.InternStr("__enter__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__getattr__ := πg.InternStr("__getattr__")
		ß__getitem__ := πg.InternStr("__getitem__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__setitem__ := πg.InternStr("__setitem__")
		ß_changed := πg.InternStr("_changed")
		ß_count := πg.InternStr("_count")
		ß_environ := πg.InternStr("_environ")
		ß_filterwarnings := πg.InternStr("_filterwarnings")
		ß_id := πg.InternStr("_id")
		ß_last := πg.InternStr("_last")
		ß_rmdir := πg.InternStr("_rmdir")
		ß_rmtree := πg.InternStr("_rmtree")
		ß_run_suite := πg.InternStr("_run_suite")
		ß_unlink := πg.InternStr("_unlink")
		ß_waitfor := πg.InternStr("_waitfor")
		ß_warnings := πg.InternStr("_warnings")
		ßaddTest := πg.InternStr("addTest")
		ßalways := πg.InternStr("always")
		ßappend := πg.InternStr("append")
		ßcatch_warnings := πg.InternStr("catch_warnings")
		ßcheck_py3k_warnings := πg.InternStr("check_py3k_warnings")
		ßcheck_warnings := πg.InternStr("check_warnings")
		ßcontextlib := πg.InternStr("contextlib")
		ßcontextmanager := πg.InternStr("contextmanager")
		ßcopy := πg.InternStr("copy")
		ßcpython_only := πg.InternStr("cpython_only")
		ßenviron := πg.InternStr("environ")
		ßerrors := πg.InternStr("errors")
		ßfailures := πg.InternStr("failures")
		ßfindTestCases := πg.InternStr("findTestCases")
		ßget := πg.InternStr("get")
		ßgetattr := πg.InternStr("getattr")
		ßgetpid := πg.InternStr("getpid")
		ßhasattr := πg.InternStr("hasattr")
		ßhave_unicode := πg.InternStr("have_unicode")
		ßisdir := πg.InternStr("isdir")
		ßisinstance := πg.InternStr("isinstance")
		ßissubclass := πg.InternStr("issubclass")
		ßitems := πg.InternStr("items")
		ßjava := πg.InternStr("java")
		ßjoin := πg.InternStr("join")
		ßkeys := πg.InternStr("keys")
		ßlen := πg.InternStr("len")
		ßlistdir := πg.InternStr("listdir")
		ßmakeSuite := πg.InternStr("makeSuite")
		ßmatch := πg.InternStr("match")
		ßmaxsize := πg.InternStr("maxsize")
		ßmessage := πg.InternStr("message")
		ßmodules := πg.InternStr("modules")
		ßname := πg.InternStr("name")
		ßobject := πg.InternStr("object")
		ßoriginal_modules := πg.InternStr("original_modules")
		ßos := πg.InternStr("os")
		ßpath := πg.InternStr("path")
		ßplatform := πg.InternStr("platform")
		ßproperty := πg.InternStr("property")
		ßpy3kwarning := πg.InternStr("py3kwarning")
		ßquiet := πg.InternStr("quiet")
		ßrange := πg.InternStr("range")
		ßre := πg.InternStr("re")
		ßreap_children := πg.InternStr("reap_children")
		ßremove := πg.InternStr("remove")
		ßrequires_unicode := πg.InternStr("requires_unicode")
		ßreset := πg.InternStr("reset")
		ßriscos := πg.InternStr("riscos")
		ßrmdir := πg.InternStr("rmdir")
		ßrun := πg.InternStr("run")
		ßrun_unittest := πg.InternStr("run_unittest")
		ßset := πg.InternStr("set")
		ßsimplefilter := πg.InternStr("simplefilter")
		ßskipUnless := πg.InternStr("skipUnless")
		ßsleep := πg.InternStr("sleep")
		ßsplit := πg.InternStr("split")
		ßstartswith := πg.InternStr("startswith")
		ßstdout := πg.InternStr("stdout")
		ßstr := πg.InternStr("str")
		ßsys := πg.InternStr("sys")
		ßtestfile := πg.InternStr("testfile")
		ßthread := πg.InternStr("thread")
		ßthreading_cleanup := πg.InternStr("threading_cleanup")
		ßthreading_setup := πg.InternStr("threading_setup")
		ßtime := πg.InternStr("time")
		ßunicode := πg.InternStr("unicode")
		ßunittest := πg.InternStr("unittest")
		ßunlink := πg.InternStr("unlink")
		ßunset := πg.InternStr("unset")
		ßupdate := πg.InternStr("update")
		ßverbose := πg.InternStr("verbose")
		ßwaitpid := πg.InternStr("waitpid")
		ßwarn := πg.InternStr("warn")
		ßwarnings := πg.InternStr("warnings")
		ßwasSuccessful := πg.InternStr("wasSuccessful")
		ßwin := πg.InternStr("win")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.BaseException
		_ = πTemp003
		var πTemp004 *πg.Traceback
		_ = πTemp004
		var πTemp005 bool
		_ = πTemp005
		var πTemp006 *πg.Dict
		_ = πTemp006
		var πTemp007 *πg.Object
		_ = πTemp007
		var πTemp008 *πg.Object
		_ = πTemp008
		var πTemp009 []πg.Param
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
		var πTemp021 *πg.Object
		_ = πTemp021
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 8: goto Label8
			case 2: goto Label2
			default: panic("unexpected function state")
			}
			// line 1: """Supporting definitions for the Python regression tests."""
			πF.SetLineno(1)
			// line 6: import contextlib
			πF.SetLineno(6)
			if πTemp002, πE = πg.ImportModule(πF, "contextlib"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßcontextlib.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 11: import sys
			πF.SetLineno(11)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 12: import os
			πF.SetLineno(12)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 15: import warnings
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "warnings"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßwarnings.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 16: import unittest
			πF.SetLineno(16)
			if πTemp002, πE = πg.ImportModule(πF, "unittest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: import UserDict
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "UserDict"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßUserDict.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 23: try:
			πF.SetLineno(23)
			πF.PushCheckpoint(2)
			// line 24: import thread
			πF.SetLineno(24)
			if πTemp002, πE = πg.ImportModule(πF, "thread"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßthread.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label1
		Label2:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp003, πTemp004 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsInstance(πF, πTemp003.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label3
			}
			πE = πF.Raise(πTemp003.ToObject(), nil, πTemp004.ToObject())
			continue
			// line 25: except ImportError:
			πF.SetLineno(25)
		Label3:
			// line 26: thread = None
			πF.SetLineno(26)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßthread.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label1
		Label1:
			// line 28: __all__ = [
			πF.SetLineno(28)
			πTemp002 = make([]*πg.Object, 9)
			πTemp002[0] = ßError.ToObject()
			πTemp002[1] = ßTestFailed.ToObject()
			πTemp002[2] = ßhave_unicode.ToObject()
			πTemp002[3] = ßBasicTestRunner.ToObject()
			πTemp002[4] = ßrun_unittest.ToObject()
			πTemp002[5] = ßcheck_warnings.ToObject()
			πTemp002[6] = ßcheck_py3k_warnings.ToObject()
			πTemp002[7] = ßCleanImport.ToObject()
			πTemp002[8] = ßEnvironmentVarGuard.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 51: class Error(Exception):
			πF.SetLineno(51)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Error", "build/src/__python__/test/test_support.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 52: """Base class for regression test exceptions."""
					πF.SetLineno(52)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp007, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp007 == nil {
				πTemp007 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp007.Call(πF, []*πg.Object{πg.NewStr("Error").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßError.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 54: class TestFailed(Error):
			πF.SetLineno(54)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßError); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TestFailed", "build/src/__python__/test/test_support.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 55: """Test failed."""
					πF.SetLineno(55)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp007, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp007 == nil {
				πTemp007 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp007.Call(πF, []*πg.Object{πg.NewStr("TestFailed").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTestFailed.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 168: verbose = 1              # Flag set to 0 by regrtest.py
			πF.SetLineno(168)
			if πE = πF.Globals().SetItem(πF, ßverbose.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = ßwin.ToObject()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp007, πE = πg.GetAttr(πF, πTemp001, ßplatform, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πg.GetAttr(πF, πTemp007, ßstartswith, nil); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πTemp005, πE = πg.IsTrue(πF, πTemp007); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label4
			}
			goto Label5
			// line 191: if sys.platform.startswith("win"):
			πF.SetLineno(191)
		Label4:
			// line 192: def _waitfor(func, pathname, waitall=False):
			πF.SetLineno(192)
			πTemp009 = make([]πg.Param, 3)
			πTemp009[0] = πg.Param{Name: "func", Def: nil}
			πTemp009[1] = πg.Param{Name: "pathname", Def: nil}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp009[2] = πg.Param{Name: "waitall", Def: πTemp007}
			πTemp001 = πg.NewFunction(πg.NewCode("_waitfor", "build/src/__python__/test/test_support.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfunc *πg.Object = πArgs[0]; _ = µfunc
				var µpathname *πg.Object = πArgs[1]; _ = µpathname
				var µwaitall *πg.Object = πArgs[2]; _ = µwaitall
				var µdirname *πg.Object = πg.UnboundLocal; _ = µdirname
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µtimeout *πg.Object = πg.UnboundLocal; _ = µtimeout
				var µL *πg.Object = πg.UnboundLocal; _ = µL
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 πg.KWArgs
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 5: goto Label5
					case 6: goto Label6
					default: panic("unexpected function state")
					}
					// line 194: func(pathname)
					πF.SetLineno(194)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpathname, "pathname"); πE != nil {
						continue
					}
					πTemp001[0] = µpathname
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp002, πE = µfunc.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µwaitall, "waitall"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µwaitall); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 196: if waitall:
					πF.SetLineno(196)
				Label1:
					// line 197: dirname = pathname
					πF.SetLineno(197)
					if πE = πg.CheckLocal(πF, µpathname, "pathname"); πE != nil {
						continue
					}
					µdirname = µpathname
					goto Label3
				Label2:
					// line 199: dirname, name = os.path.split(pathname)
					πF.SetLineno(199)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpathname, "pathname"); πE != nil {
						continue
					}
					πTemp001[0] = µpathname
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp004, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp005}}}, πTemp004); πE != nil {
						continue
					}
					µdirname = πTemp002
					µname = πTemp005
					// line 200: dirname = dirname or '.'
					πF.SetLineno(200)
					if πE = πg.CheckLocal(πF, µdirname, "dirname"); πE != nil {
						continue
					}
					πTemp002 = µdirname
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label4
					}
					πTemp002 = πg.NewStr(".").ToObject()
				Label4:
					µdirname = πTemp002
					goto Label3
				Label3:
					// line 207: timeout = 0.001
					πF.SetLineno(207)
					µtimeout = πg.NewFloat(0.001).ToObject()
					// line 208: while timeout < 1.0:
					πF.SetLineno(208)
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
					if πE = πg.CheckLocal(πF, µtimeout, "timeout"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LT(πF, µtimeout, πg.NewFloat(1.0).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(5)            
					// line 216: L = os.listdir(dirname)
					πF.SetLineno(216)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdirname, "dirname"); πE != nil {
						continue
					}
					πTemp001[0] = µdirname
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßlistdir, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µL = πTemp002
					if πE = πg.CheckLocal(πF, µwaitall, "waitall"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µwaitall); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label8
					}
					if πE = πg.CheckLocal(πF, µL, "L"); πE != nil {
						continue
					}
					πTemp004 = µL
					goto Label9
				Label8:
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µL, "L"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Contains(πF, µL, µname); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(πTemp006).ToObject()
					πTemp004 = πTemp005
				Label9:
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label10
					}
					goto Label11
					// line 217: if not (L if waitall else name in L):
					πF.SetLineno(217)
				Label10:
					// line 218: return
					πF.SetLineno(218)
					πR = πg.None
					continue
					goto Label11
				Label11:
					// line 220: time.sleep(timeout)
					πF.SetLineno(220)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtimeout, "timeout"); πE != nil {
						continue
					}
					πTemp001[0] = µtimeout
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßsleep, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 221: timeout *= 2
					πF.SetLineno(221)
					if πE = πg.CheckLocal(πF, µtimeout, "timeout"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IMul(πF, µtimeout, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					µtimeout = πTemp002
					continue
				Label6:
					if πE != nil || πR != nil {
						continue
					}
				Label7:
					// line 222: warnings.warn('tests may fail, delete still pending for ' + pathname,
					πF.SetLineno(222)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpathname, "pathname"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πg.NewStr("tests may fail, delete still pending for ").ToObject(), µpathname); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßRuntimeWarning); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					πTemp007 = πg.KWArgs{
						{"stacklevel", πg.NewInt(4).ToObject()},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßwarnings); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßwarn, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, πTemp001, πTemp007); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_waitfor.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 225: def _unlink(filename):
			πF.SetLineno(225)
			πTemp009 = make([]πg.Param, 1)
			πTemp009[0] = πg.Param{Name: "filename", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("_unlink", "build/src/__python__/test/test_support.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfilename *πg.Object = πArgs[0]; _ = µfilename
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
					// line 226: _waitfor(os.unlink, filename)
					πF.SetLineno(226)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunlink, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp001[1] = µfilename
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_waitfor); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_unlink.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 228: def _rmdir(dirname):
			πF.SetLineno(228)
			πTemp009 = make([]πg.Param, 1)
			πTemp009[0] = πg.Param{Name: "dirname", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("_rmdir", "build/src/__python__/test/test_support.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µdirname *πg.Object = πArgs[0]; _ = µdirname
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
					// line 229: _waitfor(os.rmdir, dirname)
					πF.SetLineno(229)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrmdir, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µdirname, "dirname"); πE != nil {
						continue
					}
					πTemp001[1] = µdirname
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_waitfor); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_rmdir.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 231: def _rmtree(path):
			πF.SetLineno(231)
			πTemp009 = make([]πg.Param, 1)
			πTemp009[0] = πg.Param{Name: "path", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("_rmtree", "build/src/__python__/test/test_support.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πArgs[0]; _ = µpath
				var µ_rmtree_inner *πg.Object = πg.UnboundLocal; _ = µ_rmtree_inner
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 πg.KWArgs
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
					// line 232: def _rmtree_inner(path):
					πF.SetLineno(232)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "path", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("_rmtree_inner", "build/src/__python__/test/test_support.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µpath *πg.Object = πArgs[0]; _ = µpath
						var µname *πg.Object = πg.UnboundLocal; _ = µname
						var µfullname *πg.Object = πg.UnboundLocal; _ = µfullname
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
						var πTemp007 πg.KWArgs
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
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp002[0] = µpath
							if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßlistdir, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								µname = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 234: fullname = os.path.join(path, name)
							πF.SetLineno(234)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp002[0] = µpath
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp002[1] = µname
							if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µfullname = πTemp004
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfullname, "fullname"); πE != nil {
								continue
							}
							πTemp002[0] = µfullname
							if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßisdir, nil); πE != nil {
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
							// line 235: if os.path.isdir(fullname):
							πF.SetLineno(235)
						Label4:
							// line 236: _waitfor(_rmtree_inner, fullname, waitall=True)
							πF.SetLineno(236)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µ_rmtree_inner, "_rmtree_inner"); πE != nil {
								continue
							}
							πTemp002[0] = µ_rmtree_inner
							if πE = πg.CheckLocal(πF, µfullname, "fullname"); πE != nil {
								continue
							}
							πTemp002[1] = µfullname
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp007 = πg.KWArgs{
								{"waitall", πTemp003},
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_waitfor); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, πTemp007); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 237: os.rmdir(fullname)
							πF.SetLineno(237)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfullname, "fullname"); πE != nil {
								continue
							}
							πTemp002[0] = µfullname
							if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßrmdir, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label6
						Label5:
							// line 239: os.unlink(fullname)
							πF.SetLineno(239)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfullname, "fullname"); πE != nil {
								continue
							}
							πTemp002[0] = µfullname
							if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßunlink, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
					µ_rmtree_inner = πTemp001
					// line 240: _waitfor(_rmtree_inner, path, waitall=True)
					πF.SetLineno(240)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µ_rmtree_inner, "_rmtree_inner"); πE != nil {
						continue
					}
					πTemp003[0] = µ_rmtree_inner
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp003[1] = µpath
					if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp005 = πg.KWArgs{
						{"waitall", πTemp004},
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_waitfor); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp003, πTemp005); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 241: _waitfor(os.rmdir, path)
					πF.SetLineno(241)
					πTemp003 = πF.MakeArgs(2)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßrmdir, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp006
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp003[1] = µpath
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_waitfor); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_rmtree.ToObject(), πTemp010); πE != nil {
				continue
			}
			goto Label6
		Label5:
			// line 243: _unlink = os.unlink
			πF.SetLineno(243)
			if πTemp011, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßunlink, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_unlink.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 244: _rmdir = os.rmdir
			πF.SetLineno(244)
			if πTemp011, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßrmdir, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_rmdir.ToObject(), πTemp012); πE != nil {
				continue
			}
			goto Label6
		Label6:
			// line 247: def unlink(filename):
			πF.SetLineno(247)
			πTemp009 = make([]πg.Param, 1)
			πTemp009[0] = πg.Param{Name: "filename", Def: nil}
			πTemp011 = πg.NewFunction(πg.NewCode("unlink", "build/src/__python__/test/test_support.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfilename *πg.Object = πArgs[0]; _ = µfilename
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 248: try:
					πF.SetLineno(248)
					πF.PushCheckpoint(2)
					// line 249: _unlink(filename)
					πF.SetLineno(249)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp001[0] = µfilename
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_unlink); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
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
					// line 250: except OSError:
					πF.SetLineno(250)
				Label3:
					// line 251: pass
					πF.SetLineno(251)
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßunlink.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 566: try:
			πF.SetLineno(566)
			πF.PushCheckpoint(8)
			// line 567: unicode
			πF.SetLineno(567)
			if πTemp012, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
				continue
			}
			// line 568: have_unicode = True
			πF.SetLineno(568)
			if πTemp012, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßhave_unicode.ToObject(), πTemp012); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label7
		Label8:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp003, πTemp004 = πF.ExcInfo()
			if πTemp012, πE = πg.ResolveGlobal(πF, ßNameError); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsInstance(πF, πTemp003.ToObject(), πTemp012); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label9
			}
			πE = πF.Raise(πTemp003.ToObject(), nil, πTemp004.ToObject())
			continue
			// line 569: except NameError:
			πF.SetLineno(569)
		Label9:
			// line 570: have_unicode = False
			πF.SetLineno(570)
			if πTemp012, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßhave_unicode.ToObject(), πTemp012); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label7
		Label7:
			// line 572: requires_unicode = unittest.skipUnless(have_unicode, 'no unicode support')
			πF.SetLineno(572)
			πTemp002 = πF.MakeArgs(2)
			if πTemp012, πE = πg.ResolveGlobal(πF, ßhave_unicode); πE != nil {
				continue
			}
			πTemp002[0] = πTemp012
			πTemp002[1] = πg.NewStr("no unicode support").ToObject()
			if πTemp012, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßskipUnless, nil); πE != nil {
				continue
			}
			if πTemp012, πE = πTemp013.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßrequires_unicode.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 579: FS_NONASCII = None
			πF.SetLineno(579)
			if πTemp012, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFS_NONASCII.ToObject(), πTemp012); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßname, nil); πE != nil {
				continue
			}
			if πTemp012, πE = πg.Eq(πF, πTemp014, ßjava.ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsTrue(πF, πTemp012); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label10
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßname, nil); πE != nil {
				continue
			}
			if πTemp012, πE = πg.Eq(πF, πTemp014, ßriscos.ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsTrue(πF, πTemp012); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label11
			}
			goto Label12
			// line 624: if os.name == 'java':
			πF.SetLineno(624)
		Label10:
			// line 626: TESTFN = '$test'
			πF.SetLineno(626)
			if πE = πF.Globals().SetItem(πF, ßTESTFN.ToObject(), πg.NewStr("$test").ToObject()); πE != nil {
				continue
			}
			goto Label13
			// line 627: elif os.name == 'riscos':
			πF.SetLineno(627)
		Label11:
			// line 628: TESTFN = 'testfile'
			πF.SetLineno(628)
			if πE = πF.Globals().SetItem(πF, ßTESTFN.ToObject(), ßtestfile.ToObject()); πE != nil {
				continue
			}
			goto Label13
		Label12:
			// line 630: TESTFN = '@test'
			πF.SetLineno(630)
			if πE = πF.Globals().SetItem(πF, ßTESTFN.ToObject(), πg.NewStr("@test").ToObject()); πE != nil {
				continue
			}
			goto Label13
		Label13:
			// line 671: TESTFN = "%s_%s_tmp" % (TESTFN, os.getpid())
			πF.SetLineno(671)
			if πTemp014, πE = πg.ResolveGlobal(πF, ßTESTFN); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp016, πE = πg.GetAttr(πF, πTemp015, ßgetpid, nil); πE != nil {
				continue
			}
			if πTemp015, πE = πTemp016.Call(πF, nil, nil); πE != nil {
				continue
			}
			πTemp013 = πg.NewTuple2(πTemp014, πTemp015).ToObject()
			if πTemp012, πE = πg.Mod(πF, πg.NewStr("%s_%s_tmp").ToObject(), πTemp013); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTESTFN.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 821: class WarningsRecorder(object):
			πF.SetLineno(821)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp014, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp014
			πTemp006 = πg.NewDict()
			if πTemp012, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp012); πE != nil {
				continue
			}
			_, πE = πg.NewCode("WarningsRecorder", "build/src/__python__/test/test_support.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 822: """Convenience wrapper for the warnings list returned on
					πF.SetLineno(822)
					// line 825: def __init__(self, warnings_list):
					πF.SetLineno(825)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "warnings_list", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/test/test_support.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µwarnings_list *πg.Object = πArgs[1]; _ = µwarnings_list
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 826: self._warnings = warnings_list
							πF.SetLineno(826)
							if πE = πg.CheckLocal(πF, µwarnings_list, "warnings_list"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µwarnings_list); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_warnings, πTemp001); πE != nil {
								continue
							}
							// line 827: self._last = 0
							πF.SetLineno(827)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_last, πTemp001); πE != nil {
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
					// line 829: def __getattr__(self, attr):
					πF.SetLineno(829)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "attr", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__getattr__", "build/src/__python__/test/test_support.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µattr *πg.Object = πArgs[1]; _ = µattr
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
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_warnings, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_last, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GT(πF, πTemp004, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßwarnings); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßWarningMessage, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ß_WARNING_DETAILS, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, πTemp003, µattr); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label2
							}
							goto Label3
							// line 830: if len(self._warnings) > self._last:
							πF.SetLineno(830)
						Label1:
							// line 831: return getattr(self._warnings[-1], attr)
							πF.SetLineno(831)
							πTemp002 = πF.MakeArgs(2)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_warnings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp002[1] = µattr
							if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πR = πTemp003
							continue
							goto Label3
							// line 832: elif attr in warnings.WarningMessage._WARNING_DETAILS:
							πF.SetLineno(832)
						Label2:
							// line 833: return None
							πF.SetLineno(833)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label3
						Label3:
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(µself, µattr).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%r has no attribute %r").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 834: raise AttributeError("%r has no attribute %r" % (self, attr))
							πF.SetLineno(834)
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
					if πE = πClass.SetItem(πF, ß__getattr__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 837: def warnings(self):
					πF.SetLineno(837)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("warnings", "build/src/__python__/test/test_support.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
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
							// line 838: return self._warnings[self._last:]
							πF.SetLineno(838)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_last, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πTemp002, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_warnings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßwarnings.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 836: @property
					πF.SetLineno(836)
					πTemp005 = πF.MakeArgs(1)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßwarnings); πE != nil {
						continue
					}
					πTemp005[0] = πTemp006
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πClass.SetItem(πF, ßwarnings.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 840: def reset(self):
					πF.SetLineno(840)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("reset", "build/src/__python__/test/test_support.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 841: self._last = len(self._warnings)
							πF.SetLineno(841)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_warnings, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ß_last, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßreset.ToObject(), πTemp006); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp013, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp013 == nil {
				πTemp013 = πg.TypeType.ToObject()
			}
			if πTemp014, πE = πTemp013.Call(πF, []*πg.Object{πg.NewStr("WarningsRecorder").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßWarningsRecorder.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 844: def _filterwarnings(filters, quiet=False):
			πF.SetLineno(844)
			πTemp009 = make([]πg.Param, 2)
			πTemp009[0] = πg.Param{Name: "filters", Def: nil}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp009[1] = πg.Param{Name: "quiet", Def: πTemp013}
			πTemp012 = πg.NewFunction(πg.NewCode("_filterwarnings", "build/src/__python__/test/test_support.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfilters *πg.Object = πArgs[0]; _ = µfilters
				var µquiet *πg.Object = πArgs[1]; _ = µquiet
				var µw *πg.Object = πg.UnboundLocal; _ = µw
				var µreraise *πg.Object = πg.UnboundLocal; _ = µreraise
				var µmissing *πg.Object = πg.UnboundLocal; _ = µmissing
				var µmsg *πg.Object = πg.UnboundLocal; _ = µmsg
				var µcat *πg.Object = πg.UnboundLocal; _ = µcat
				var µseen *πg.Object = πg.UnboundLocal; _ = µseen
				var µexc *πg.Object = πg.UnboundLocal; _ = µexc
				var µmessage *πg.Object = πg.UnboundLocal; _ = µmessage
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 πg.KWArgs
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
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 bool
				_ = πTemp010
				var πTemp011 *πg.BaseException
				_ = πTemp011
				var πTemp012 *πg.Traceback
				_ = πTemp012
				var πTemp013 *πg.Type
				_ = πTemp013
				var πTemp014 []πg.Param
				_ = πTemp014
				var πTemp015 bool
				_ = πTemp015
				var πTemp016 bool
				_ = πTemp016
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 1: goto Label1
						case 2: goto Label2
						case 3: goto Label3
						case 4: goto Label4
						case 6: goto Label6
						case 7: goto Label7
						default: panic("unexpected function state")
						}
						// line 845: """Catch the warnings, then check if all the expected
						πF.SetLineno(845)
						// line 855: with warnings.catch_warnings(record=True) as w:
						πF.SetLineno(855)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						πTemp002 = πg.KWArgs{
							{"record", πTemp001},
						}
						if πTemp001, πE = πg.ResolveGlobal(πF, ßwarnings); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßcatch_warnings, nil); πE != nil {
							continue
						}
						if πTemp001, πE = πTemp003.Call(πF, nil, πTemp002); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp001.Type().ToObject(), ß__exit__, nil); πE != nil {
							continue
						}
						if πTemp004, πE = πg.GetAttr(πF, πTemp001.Type().ToObject(), ß__enter__, nil); πE != nil {
							continue
						}
						if πTemp004, πE = πTemp004.Call(πF, πg.Args{πTemp001}, nil); πE != nil {
							continue
						}
						πF.PushCheckpoint(1)
						µw = πTemp004
						// line 859: sys.modules['warnings'].simplefilter("always")
						πF.SetLineno(859)
						πTemp005 = πF.MakeArgs(1)
						πTemp005[0] = ßalways.ToObject()
						πTemp006 = ßwarnings.ToObject()
						if πTemp008, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßmodules, nil); πE != nil {
							continue
						}
						if πTemp007, πE = πg.GetItem(πF, πTemp009, πTemp006); πE != nil {
							continue
						}
						if πTemp006, πE = πg.GetAttr(πF, πTemp007, ßsimplefilter, nil); πE != nil {
							continue
						}
						if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp005)
						// line 860: yield WarningsRecorder(w)
						πF.SetLineno(860)
						πTemp005 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µw, "w"); πE != nil {
							continue
						}
						πTemp005[0] = µw
						if πTemp006, πE = πg.ResolveGlobal(πF, ßWarningsRecorder); πE != nil {
							continue
						}
						if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp005)
						πF.PushCheckpoint(2)
						return πTemp007, nil
					Label2:
						πTemp006 = πSent
						πF.PopCheckpoint()
					Label1:
						πTemp011, πTemp012 = nil, nil
						if πE != nil {
							πTemp011, πTemp012 = πF.ExcInfo()
						}
						if πTemp011 != nil {
							πTemp013 = πTemp011.Type()
							if πTemp006, πE = πTemp003.Call(πF, πg.Args{πTemp001, πTemp013.ToObject(), πTemp011.ToObject(), πTemp012.ToObject()}, nil); πE != nil {
								continue
							}
						} else {
							if πTemp006, πE = πTemp003.Call(πF, πg.Args{πTemp001, πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
						}
						if πTemp010, πE = πg.IsTrue(πF, πTemp006); πE != nil {
							continue
						}
						if πTemp011 != nil && πTemp010 != true {
							πE = πF.Raise(nil, nil, nil)
							continue
						}
						if πR != nil {
							continue
						}
						// line 862: reraise = [warning.message for warning in w]
						πF.SetLineno(862)
						πTemp014 = make([]πg.Param, 0)
						πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/test/test_support.py", πTemp014, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
							var µwarning *πg.Object = πg.UnboundLocal; _ = µwarning
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
									if πE = πg.CheckLocal(πF, µw, "w"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.Iter(πF, µw); πE != nil {
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
										µwarning = πTemp004
									}
									if πE != nil || !πTemp003 {
										continue
									}
									πF.PushCheckpoint(1)            
									// line 862: reraise = [warning.message for warning in w]
									πF.SetLineno(862)
									if πE = πg.CheckLocal(πF, µwarning, "warning"); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetAttr(πF, µwarning, ßmessage, nil); πE != nil {
										continue
									}
									πF.PushCheckpoint(4)
									return πTemp004, nil
								Label4:
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
						if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
							continue
						}
						µreraise = πTemp001
						// line 863: missing = []
						πF.SetLineno(863)
						πTemp005 = make([]*πg.Object, 0)
						πTemp001 = πg.NewList(πTemp005...).ToObject()
						µmissing = πTemp001
						if πE = πg.CheckLocal(πF, µfilters, "filters"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.Iter(πF, µfilters); πE != nil {
							continue
						}
						πF.PushCheckpoint(4)
						πTemp010 = false
					Label3:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp010 {
							πF.PopCheckpoint()
							goto Label5
						}
						if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
							isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
							if exc != nil {
								πE = exc
							} else if isStop {
								πE = nil
								πF.RestoreExc(nil, nil)
							}
							πTemp015 = !isStop
						} else {
							πTemp015 = true
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp008}}}, πTemp004); πE != nil {
								continue
							}
							µmsg = πTemp006
							µcat = πTemp008
						}
						if πE != nil || !πTemp015 {
							continue
						}
						πF.PushCheckpoint(3)            
						// line 865: seen = False
						πF.SetLineno(865)
						if πTemp004, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
							continue
						}
						µseen = πTemp004
						if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µreraise, "reraise"); πE != nil {
							continue
						}
						if πTemp008, πE = πg.GetItem(πF, µreraise, πTemp006); πE != nil {
							continue
						}
						if πTemp004, πE = πg.Iter(πF, πTemp008); πE != nil {
							continue
						}
						πF.PushCheckpoint(7)
						πTemp015 = false
					Label6:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp015 {
							πF.PopCheckpoint()
							goto Label8
						}
						if πTemp006, πE = πg.Next(πF, πTemp004); πE != nil {
							isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
							if exc != nil {
								πE = exc
							} else if isStop {
								πE = nil
								πF.RestoreExc(nil, nil)
							}
							πTemp016 = !isStop
						} else {
							πTemp016 = true
							µexc = πTemp006
						}
						if πE != nil || !πTemp016 {
							continue
						}
						πF.PushCheckpoint(6)            
						// line 867: message = str(exc)
						πF.SetLineno(867)
						πTemp005 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µexc, "exc"); πE != nil {
							continue
						}
						πTemp005[0] = µexc
						if πTemp006, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
							continue
						}
						if πTemp008, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp005)
						µmessage = πTemp008
						πTemp005 = πF.MakeArgs(3)
						if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
							continue
						}
						πTemp005[0] = µmsg
						if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
							continue
						}
						πTemp005[1] = µmessage
						if πTemp008, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßI, nil); πE != nil {
							continue
						}
						πTemp005[2] = πTemp009
						if πTemp008, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßmatch, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πTemp009.Call(πF, πTemp005, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp005)
						πTemp006 = πTemp008
						if πTemp016, πE = πg.IsTrue(πF, πTemp006); πE != nil {
							continue
						}
						if !πTemp016 {
							goto Label9
						}
						πTemp005 = πF.MakeArgs(2)
						if πE = πg.CheckLocal(πF, µexc, "exc"); πE != nil {
							continue
						}
						if πTemp008, πE = πg.GetAttr(πF, µexc, ß__class__, nil); πE != nil {
							continue
						}
						πTemp005[0] = πTemp008
						if πE = πg.CheckLocal(πF, µcat, "cat"); πE != nil {
							continue
						}
						πTemp005[1] = µcat
						if πTemp008, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
							continue
						}
						if πTemp009, πE = πTemp008.Call(πF, πTemp005, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp005)
						πTemp006 = πTemp009
					Label9:
						if πTemp016, πE = πg.IsTrue(πF, πTemp006); πE != nil {
							continue
						}
						if πTemp016 {
							goto Label10
						}
						goto Label11
						// line 869: if (re.match(msg, message, re.I) and
						πF.SetLineno(869)
					Label10:
						// line 871: seen = True
						πF.SetLineno(871)
						if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						µseen = πTemp006
						// line 872: reraise.remove(exc)
						πF.SetLineno(872)
						πTemp005 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µexc, "exc"); πE != nil {
							continue
						}
						πTemp005[0] = µexc
						if πE = πg.CheckLocal(πF, µreraise, "reraise"); πE != nil {
							continue
						}
						if πTemp006, πE = πg.GetAttr(πF, µreraise, ßremove, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp005)
						goto Label11
					Label11:
						continue
					Label7:
						if πE != nil || πR != nil {
							continue
						}
					Label8:
						if πE = πg.CheckLocal(πF, µseen, "seen"); πE != nil {
							continue
						}
						if πTemp016, πE = πg.IsTrue(πF, µseen); πE != nil {
							continue
						}
						πTemp006 = πg.GetBool(!πTemp016).ToObject()
						πTemp004 = πTemp006
						if πTemp015, πE = πg.IsTrue(πF, πTemp004); πE != nil {
							continue
						}
						if !πTemp015 {
							goto Label12
						}
						if πE = πg.CheckLocal(πF, µquiet, "quiet"); πE != nil {
							continue
						}
						if πTemp016, πE = πg.IsTrue(πF, µquiet); πE != nil {
							continue
						}
						πTemp006 = πg.GetBool(!πTemp016).ToObject()
						πTemp004 = πTemp006
					Label12:
						if πTemp015, πE = πg.IsTrue(πF, πTemp004); πE != nil {
							continue
						}
						if πTemp015 {
							goto Label13
						}
						goto Label14
						// line 873: if not seen and not quiet:
						πF.SetLineno(873)
					Label13:
						// line 875: missing.append((msg, cat.__name__))
						πF.SetLineno(875)
						πTemp005 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µcat, "cat"); πE != nil {
							continue
						}
						if πTemp006, πE = πg.GetAttr(πF, µcat, ß__name__, nil); πE != nil {
							continue
						}
						πTemp004 = πg.NewTuple2(µmsg, πTemp006).ToObject()
						πTemp005[0] = πTemp004
						if πE = πg.CheckLocal(πF, µmissing, "missing"); πE != nil {
							continue
						}
						if πTemp004, πE = πg.GetAttr(πF, µmissing, ßappend, nil); πE != nil {
							continue
						}
						if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp005)
						goto Label14
					Label14:
						continue
					Label4:
						if πE != nil || πR != nil {
							continue
						}
					Label5:
						if πE = πg.CheckLocal(πF, µreraise, "reraise"); πE != nil {
							continue
						}
						if πTemp010, πE = πg.IsTrue(πF, µreraise); πE != nil {
							continue
						}
						if πTemp010 {
							goto Label15
						}
						goto Label16
						// line 876: if reraise:
						πF.SetLineno(876)
					Label15:
						πTemp005 = πF.MakeArgs(1)
						πTemp004 = πg.NewInt(0).ToObject()
						if πE = πg.CheckLocal(πF, µreraise, "reraise"); πE != nil {
							continue
						}
						if πTemp006, πE = πg.GetItem(πF, µreraise, πTemp004); πE != nil {
							continue
						}
						if πTemp001, πE = πg.Mod(πF, πg.NewStr("unhandled warning %r").ToObject(), πTemp006); πE != nil {
							continue
						}
						πTemp005[0] = πTemp001
						if πTemp001, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
							continue
						}
						if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp005)
						// line 877: raise AssertionError("unhandled warning %r" % reraise[0])
						πF.SetLineno(877)
						πE = πF.Raise(πTemp004, nil, nil)
						continue
						goto Label16
					Label16:
						if πE = πg.CheckLocal(πF, µmissing, "missing"); πE != nil {
							continue
						}
						if πTemp010, πE = πg.IsTrue(πF, µmissing); πE != nil {
							continue
						}
						if πTemp010 {
							goto Label17
						}
						goto Label18
						// line 878: if missing:
						πF.SetLineno(878)
					Label17:
						πTemp005 = πF.MakeArgs(1)
						πTemp004 = πg.NewInt(0).ToObject()
						if πE = πg.CheckLocal(πF, µmissing, "missing"); πE != nil {
							continue
						}
						if πTemp006, πE = πg.GetItem(πF, µmissing, πTemp004); πE != nil {
							continue
						}
						if πTemp001, πE = πg.Mod(πF, πg.NewStr("filter (%r, %s) did not catch any warning").ToObject(), πTemp006); πE != nil {
							continue
						}
						πTemp005[0] = πTemp001
						if πTemp001, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
							continue
						}
						if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp005)
						// line 879: raise AssertionError("filter (%r, %s) did not catch any warning" %
						πF.SetLineno(879)
						πE = πF.Raise(πTemp004, nil, nil)
						continue
						goto Label18
					Label18:
					}
					return nil, πE
				}).ToObject(), nil
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_filterwarnings.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 884: def check_warnings(*filters, **kwargs):
			πF.SetLineno(884)
			πTemp009 = make([]πg.Param, 0)
			πTemp013 = πg.NewFunction(πg.NewCode("check_warnings", "build/src/__python__/test/test_support.py", πTemp009, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfilters *πg.Object = πArgs[0]; _ = µfilters
				var µkwargs *πg.Object = πArgs[1]; _ = µkwargs
				var µquiet *πg.Object = πg.UnboundLocal; _ = µquiet
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
					// line 885: """Context manager to silence warnings.
					πF.SetLineno(885)
					// line 898: quiet = kwargs.get('quiet')
					πF.SetLineno(898)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ßquiet.ToObject()
					if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µkwargs, ßget, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µquiet = πTemp003
					if πE = πg.CheckLocal(πF, µfilters, "filters"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µfilters); πE != nil {
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
					// line 899: if not filters:
					πF.SetLineno(899)
				Label1:
					// line 900: filters = (("", Warning),)
					πF.SetLineno(900)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßWarning); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(ß.ToObject(), πTemp005).ToObject()
					πTemp002 = πg.NewTuple1(πTemp003).ToObject()
					µfilters = πTemp002
					if πE = πg.CheckLocal(πF, µquiet, "quiet"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µquiet == πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 902: if quiet is None:
					πF.SetLineno(902)
				Label3:
					// line 903: quiet = True
					πF.SetLineno(903)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					µquiet = πTemp002
					goto Label4
				Label4:
					goto Label2
				Label2:
					// line 904: return _filterwarnings(filters, quiet)
					πF.SetLineno(904)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfilters, "filters"); πE != nil {
						continue
					}
					πTemp001[0] = µfilters
					if πE = πg.CheckLocal(πF, µquiet, "quiet"); πE != nil {
						continue
					}
					πTemp001[1] = µquiet
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_filterwarnings); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßcheck_warnings.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 883: @contextlib.contextmanager
			πF.SetLineno(883)
			πTemp002 = πF.MakeArgs(1)
			if πTemp014, πE = πg.ResolveGlobal(πF, ßcheck_warnings); πE != nil {
				continue
			}
			πTemp002[0] = πTemp014
			if πTemp014, πE = πg.ResolveGlobal(πF, ßcontextlib); πE != nil {
				continue
			}
			if πTemp015, πE = πg.GetAttr(πF, πTemp014, ßcontextmanager, nil); πE != nil {
				continue
			}
			if πTemp014, πE = πTemp015.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßcheck_warnings.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 908: def check_py3k_warnings(*filters, **kwargs):
			πF.SetLineno(908)
			πTemp009 = make([]πg.Param, 0)
			πTemp014 = πg.NewFunction(πg.NewCode("check_py3k_warnings", "build/src/__python__/test/test_support.py", πTemp009, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfilters *πg.Object = πArgs[0]; _ = µfilters
				var µkwargs *πg.Object = πArgs[1]; _ = µkwargs
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 909: """Context manager to silence py3k warnings.
					πF.SetLineno(909)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpy3kwarning, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 921: if sys.py3kwarning:
					πF.SetLineno(921)
				Label1:
					if πE = πg.CheckLocal(πF, µfilters, "filters"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µfilters); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp003).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label4
					}
					goto Label5
					// line 922: if not filters:
					πF.SetLineno(922)
				Label4:
					// line 923: filters = (("", DeprecationWarning),)
					πF.SetLineno(923)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßDeprecationWarning); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(ß.ToObject(), πTemp004).ToObject()
					πTemp001 = πg.NewTuple1(πTemp002).ToObject()
					µfilters = πTemp001
					goto Label5
				Label5:
					goto Label3
				Label2:
					// line 926: filters = ()
					πF.SetLineno(926)
					πTemp001 = πg.NewTuple0().ToObject()
					µfilters = πTemp001
					goto Label3
				Label3:
					// line 927: return _filterwarnings(filters, kwargs.get('quiet'))
					πF.SetLineno(927)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfilters, "filters"); πE != nil {
						continue
					}
					πTemp005[0] = µfilters
					πTemp006 = πF.MakeArgs(1)
					πTemp006[0] = ßquiet.ToObject()
					if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µkwargs, ßget, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp005[1] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_filterwarnings); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßcheck_py3k_warnings.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 907: @contextlib.contextmanager
			πF.SetLineno(907)
			πTemp002 = πF.MakeArgs(1)
			if πTemp015, πE = πg.ResolveGlobal(πF, ßcheck_py3k_warnings); πE != nil {
				continue
			}
			πTemp002[0] = πTemp015
			if πTemp015, πE = πg.ResolveGlobal(πF, ßcontextlib); πE != nil {
				continue
			}
			if πTemp016, πE = πg.GetAttr(πF, πTemp015, ßcontextmanager, nil); πE != nil {
				continue
			}
			if πTemp015, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßcheck_py3k_warnings.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 930: class CleanImport(object):
			πF.SetLineno(930)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp017, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp017
			πTemp006 = πg.NewDict()
			if πTemp015, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp015); πE != nil {
				continue
			}
			_, πE = πg.NewCode("CleanImport", "build/src/__python__/test/test_support.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 931: """Context manager to force import to return a new module reference.
					πF.SetLineno(931)
					// line 942: def __init__(self, *module_names):
					πF.SetLineno(942)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/test/test_support.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmodule_names *πg.Object = πArgs[1]; _ = µmodule_names
						var µmodule_name *πg.Object = πg.UnboundLocal; _ = µmodule_name
						var µmodule *πg.Object = πg.UnboundLocal; _ = µmodule
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 943: self.original_modules = sys.modules.copy()
							πF.SetLineno(943)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmodules, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßcopy, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßoriginal_modules, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmodule_names, "module_names"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µmodule_names); πE != nil {
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
								µmodule_name = πTemp002
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µmodule_name, "module_name"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßmodules, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp006, µmodule_name); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 945: if module_name in sys.modules:
							πF.SetLineno(945)
						Label4:
							// line 946: module = sys.modules[module_name]
							πF.SetLineno(946)
							if πE = πg.CheckLocal(πF, µmodule_name, "module_name"); πE != nil {
								continue
							}
							πTemp002 = µmodule_name
							if πTemp006, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßmodules, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
								continue
							}
							µmodule = πTemp005
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µmodule, ß__name__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmodule_name, "module_name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.NE(πF, πTemp005, µmodule_name); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							goto Label7
							// line 951: if module.__name__ != module_name:
							πF.SetLineno(951)
						Label6:
							// line 952: del sys.modules[module.__name__]
							πF.SetLineno(952)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßmodules, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µmodule, ß__name__, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp006
							if πE = πg.DelItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							goto Label7
						Label7:
							// line 953: del sys.modules[module_name]
							πF.SetLineno(953)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßmodules, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmodule_name, "module_name"); πE != nil {
								continue
							}
							πTemp002 = µmodule_name
							if πE = πg.DelItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							goto Label5
						Label5:
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 955: def __enter__(self):
					πF.SetLineno(955)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__enter__", "build/src/__python__/test/test_support.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 956: return self
							πF.SetLineno(956)
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
					if πE = πClass.SetItem(πF, ß__enter__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 958: def __exit__(self, *ignore_exc):
					πF.SetLineno(958)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__exit__", "build/src/__python__/test/test_support.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µignore_exc *πg.Object = πArgs[1]; _ = µignore_exc
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
							// line 959: sys.modules.update(self.original_modules)
							πF.SetLineno(959)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoriginal_modules, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmodules, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßupdate, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__exit__.ToObject(), πTemp004); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp016, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp016 == nil {
				πTemp016 = πg.TypeType.ToObject()
			}
			if πTemp017, πE = πTemp016.Call(πF, []*πg.Object{πg.NewStr("CleanImport").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCleanImport.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 962: class EnvironmentVarGuard(UserDict.DictMixin):
			πF.SetLineno(962)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp017, πE = πg.ResolveGlobal(πF, ßUserDict); πE != nil {
				continue
			}
			if πTemp018, πE = πg.GetAttr(πF, πTemp017, ßDictMixin, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp018
			πTemp006 = πg.NewDict()
			if πTemp015, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp015); πE != nil {
				continue
			}
			_, πE = πg.NewCode("EnvironmentVarGuard", "build/src/__python__/test/test_support.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 964: """Class to help protect the environment variable properly.  Can be used as
					πF.SetLineno(964)
					// line 967: def __init__(self):
					πF.SetLineno(967)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "build/src/__python__/test/test_support.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Dict
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 968: self._environ = os.environ
							πF.SetLineno(968)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßenviron, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_environ, πTemp001); πE != nil {
								continue
							}
							// line 969: self._changed = {}
							πF.SetLineno(969)
							πTemp003 = πg.NewDict()
							πTemp001 = πTemp003.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_changed, πTemp002); πE != nil {
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
					// line 971: def __getitem__(self, envvar):
					πF.SetLineno(971)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "envvar", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__getitem__", "build/src/__python__/test/test_support.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µenvvar *πg.Object = πArgs[1]; _ = µenvvar
						var πTemp001 *πg.Object
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
							// line 972: return self._environ[envvar]
							πF.SetLineno(972)
							if πE = πg.CheckLocal(πF, µenvvar, "envvar"); πE != nil {
								continue
							}
							πTemp001 = µenvvar
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_environ, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__getitem__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 974: def __setitem__(self, envvar, value):
					πF.SetLineno(974)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "envvar", Def: nil}
					πTemp002[2] = πg.Param{Name: "value", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__setitem__", "build/src/__python__/test/test_support.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µenvvar *πg.Object = πArgs[1]; _ = µenvvar
						var µvalue *πg.Object = πArgs[2]; _ = µvalue
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µenvvar, "envvar"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_changed, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, µenvvar); πE != nil {
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
							// line 976: if envvar not in self._changed:
							πF.SetLineno(976)
						Label1:
							// line 977: self._changed[envvar] = self._environ.get(envvar)
							πF.SetLineno(977)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µenvvar, "envvar"); πE != nil {
								continue
							}
							πTemp004[0] = µenvvar
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_environ, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_changed, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µenvvar, "envvar"); πE != nil {
								continue
							}
							πTemp006 = µenvvar
							if πE = πg.SetItem(πF, πTemp005, πTemp006, πTemp002); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 978: self._environ[envvar] = value
							πF.SetLineno(978)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_environ, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µenvvar, "envvar"); πE != nil {
								continue
							}
							πTemp005 = µenvvar
							if πE = πg.SetItem(πF, πTemp002, πTemp005, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__setitem__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 980: def __delitem__(self, envvar):
					πF.SetLineno(980)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "envvar", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("__delitem__", "build/src/__python__/test/test_support.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µenvvar *πg.Object = πArgs[1]; _ = µenvvar
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µenvvar, "envvar"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_changed, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, µenvvar); πE != nil {
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
							// line 982: if envvar not in self._changed:
							πF.SetLineno(982)
						Label1:
							// line 983: self._changed[envvar] = self._environ.get(envvar)
							πF.SetLineno(983)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µenvvar, "envvar"); πE != nil {
								continue
							}
							πTemp004[0] = µenvvar
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_environ, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_changed, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µenvvar, "envvar"); πE != nil {
								continue
							}
							πTemp006 = µenvvar
							if πE = πg.SetItem(πF, πTemp005, πTemp006, πTemp002); πE != nil {
								continue
							}
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µenvvar, "envvar"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_environ, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, µenvvar); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 984: if envvar in self._environ:
							πF.SetLineno(984)
						Label3:
							// line 985: del self._environ[envvar]
							πF.SetLineno(985)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_environ, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µenvvar, "envvar"); πE != nil {
								continue
							}
							πTemp002 = µenvvar
							if πE = πg.DelItem(πF, πTemp001, πTemp002); πE != nil {
								continue
							}
							goto Label4
						Label4:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__delitem__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 987: def keys(self):
					πF.SetLineno(987)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("keys", "build/src/__python__/test/test_support.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 988: return self._environ.keys()
							πF.SetLineno(988)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_environ, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßkeys.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 990: def set(self, envvar, value):
					πF.SetLineno(990)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "envvar", Def: nil}
					πTemp002[2] = πg.Param{Name: "value", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("set", "build/src/__python__/test/test_support.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µenvvar *πg.Object = πArgs[1]; _ = µenvvar
						var µvalue *πg.Object = πArgs[2]; _ = µvalue
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
							// line 991: self[envvar] = value
							πF.SetLineno(991)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µenvvar, "envvar"); πE != nil {
								continue
							}
							πTemp002 = µenvvar
							if πE = πg.SetItem(πF, µself, πTemp002, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 993: def unset(self, envvar):
					πF.SetLineno(993)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "envvar", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("unset", "build/src/__python__/test/test_support.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µenvvar *πg.Object = πArgs[1]; _ = µenvvar
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 994: del self[envvar]
							πF.SetLineno(994)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µenvvar, "envvar"); πE != nil {
								continue
							}
							πTemp001 = µenvvar
							if πE = πg.DelItem(πF, µself, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßunset.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 996: def __enter__(self):
					πF.SetLineno(996)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("__enter__", "build/src/__python__/test/test_support.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 997: return self
							πF.SetLineno(997)
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
					if πE = πClass.SetItem(πF, ß__enter__.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 999: def __exit__(self, *ignore_exc):
					πF.SetLineno(999)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("__exit__", "build/src/__python__/test/test_support.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µignore_exc *πg.Object = πArgs[1]; _ = µignore_exc
						var µk *πg.Object = πg.UnboundLocal; _ = µk
						var µv *πg.Object = πg.UnboundLocal; _ = µv
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
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
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_changed, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßitems, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
									continue
								}
								µk = πTemp003
								µv = πTemp006
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µv == πTemp003).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 1001: if v is None:
							πF.SetLineno(1001)
						Label4:
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_environ, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, πTemp003, µk); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label7
							}
							goto Label8
							// line 1002: if k in self._environ:
							πF.SetLineno(1002)
						Label7:
							// line 1003: del self._environ[k]
							πF.SetLineno(1003)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_environ, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp003 = µk
							if πE = πg.DelItem(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							goto Label8
						Label8:
							goto Label6
						Label5:
							// line 1005: self._environ[k] = v
							πF.SetLineno(1005)
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µv); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_environ, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp006 = µk
							if πE = πg.SetItem(πF, πTemp003, πTemp006, πTemp002); πE != nil {
								continue
							}
							goto Label6
						Label6:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 1006: os.environ = self._environ
							πF.SetLineno(1006)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_environ, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp003, ßenviron, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__exit__.ToObject(), πTemp010); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp016, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp016 == nil {
				πTemp016 = πg.TypeType.ToObject()
			}
			if πTemp017, πE = πTemp016.Call(πF, []*πg.Object{πg.NewStr("EnvironmentVarGuard").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßEnvironmentVarGuard.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 1276: MAX_Py_ssize_t = sys.maxsize
			πF.SetLineno(1276)
			if πTemp015, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp016, πE = πg.GetAttr(πF, πTemp015, ßmaxsize, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMAX_Py_ssize_t.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 1374: class BasicTestRunner(object):
			πF.SetLineno(1374)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp017, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp017
			πTemp006 = πg.NewDict()
			if πTemp015, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp015); πE != nil {
				continue
			}
			_, πE = πg.NewCode("BasicTestRunner", "build/src/__python__/test/test_support.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
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
					// line 1375: def run(self, test):
					πF.SetLineno(1375)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "test", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("run", "build/src/__python__/test/test_support.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtest *πg.Object = πArgs[1]; _ = µtest
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1376: result = unittest.TestResult()
							πF.SetLineno(1376)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßTestResult, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µresult = πTemp001
							// line 1377: test(result)
							πF.SetLineno(1377)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp003[0] = µresult
							if πE = πg.CheckLocal(πF, µtest, "test"); πE != nil {
								continue
							}
							if πTemp001, πE = µtest.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 1378: return result
							πF.SetLineno(1378)
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
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp016, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp016 == nil {
				πTemp016 = πg.TypeType.ToObject()
			}
			if πTemp017, πE = πTemp016.Call(πF, []*πg.Object{πg.NewStr("BasicTestRunner").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBasicTestRunner.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 1380: def _id(obj):
			πF.SetLineno(1380)
			πTemp009 = make([]πg.Param, 1)
			πTemp009[0] = πg.Param{Name: "obj", Def: nil}
			πTemp015 = πg.NewFunction(πg.NewCode("_id", "build/src/__python__/test/test_support.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1381: return obj
					πF.SetLineno(1381)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πR = µobj
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_id.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 1391: def cpython_only(test):
			πF.SetLineno(1391)
			πTemp009 = make([]πg.Param, 1)
			πTemp009[0] = πg.Param{Name: "test", Def: nil}
			πTemp016 = πg.NewFunction(πg.NewCode("cpython_only", "build/src/__python__/test/test_support.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtest *πg.Object = πArgs[0]; _ = µtest
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1392: return lambda *arg, **kw: None
					πF.SetLineno(1392)
					πTemp002 = make([]πg.Param, 0)
					πTemp001 = πg.NewFunction(πg.NewCode("<lambda>", "build/src/__python__/test/test_support.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µarg *πg.Object = πArgs[0]; _ = µarg
						var µkw *πg.Object = πArgs[1]; _ = µkw
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1392: return lambda *arg, **kw: None
							πF.SetLineno(1392)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßcpython_only.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 1435: def _run_suite(suite):
			πF.SetLineno(1435)
			πTemp009 = make([]πg.Param, 1)
			πTemp009[0] = πg.Param{Name: "suite", Def: nil}
			πTemp017 = πg.NewFunction(πg.NewCode("_run_suite", "build/src/__python__/test/test_support.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsuite *πg.Object = πArgs[0]; _ = µsuite
				var µrunner *πg.Object = πg.UnboundLocal; _ = µrunner
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µerr *πg.Object = πg.UnboundLocal; _ = µerr
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 πg.KWArgs
				_ = πTemp005
				var πTemp006 *πg.Object
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
					default: panic("unexpected function state")
					}
					// line 1436: """Run tests from a unittest.TestSuite-derived class."""
					πF.SetLineno(1436)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßverbose); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 1437: if verbose:
					πF.SetLineno(1437)
				Label1:
					// line 1438: runner = unittest.TextTestRunner(sys.stdout, verbosity=2)
					πF.SetLineno(1438)
					πTemp003 = πF.MakeArgs(1)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßstdout, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					πTemp005 = πg.KWArgs{
						{"verbosity", πg.NewInt(2).ToObject()},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßTextTestRunner, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, πTemp003, πTemp005); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µrunner = πTemp001
					goto Label3
				Label2:
					// line 1440: runner = BasicTestRunner()
					πF.SetLineno(1440)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßBasicTestRunner); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µrunner = πTemp004
					goto Label3
				Label3:
					// line 1442: result = runner.run(suite)
					πF.SetLineno(1442)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsuite, "suite"); πE != nil {
						continue
					}
					πTemp003[0] = µsuite
					if πE = πg.CheckLocal(πF, µrunner, "runner"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µrunner, ßrun, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µresult = πTemp004
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µresult, ßwasSuccessful, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp006); πE != nil {
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
					// line 1443: if not result.wasSuccessful():
					πF.SetLineno(1443)
				Label4:
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µresult, ßerrors, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp006
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp004, πE = πg.Eq(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp004
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label6
					}
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µresult, ßfailures, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(!πTemp008).ToObject()
					πTemp001 = πTemp004
				Label6:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label7
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µresult, ßfailures, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp006
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp004, πE = πg.Eq(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp004
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label8
					}
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µresult, ßerrors, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(!πTemp008).ToObject()
					πTemp001 = πTemp004
				Label8:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label9
					}
					goto Label10
					// line 1444: if len(result.errors) == 1 and not result.failures:
					πF.SetLineno(1444)
				Label7:
					// line 1445: err = result.errors[0][1]
					πF.SetLineno(1445)
					πTemp001 = πg.NewInt(1).ToObject()
					πTemp006 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, µresult, ßerrors, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, πTemp009, πTemp006); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp007, πTemp001); πE != nil {
						continue
					}
					µerr = πTemp004
					goto Label11
					// line 1446: elif len(result.failures) == 1 and not result.errors:
					πF.SetLineno(1446)
				Label9:
					// line 1447: err = result.failures[0][1]
					πF.SetLineno(1447)
					πTemp001 = πg.NewInt(1).ToObject()
					πTemp006 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, µresult, ßfailures, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, πTemp009, πTemp006); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp007, πTemp001); πE != nil {
						continue
					}
					µerr = πTemp004
					goto Label11
				Label10:
					// line 1449: err = "multiple errors occurred"
					πF.SetLineno(1449)
					µerr = πg.NewStr("multiple errors occurred").ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßverbose); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label12
					}
					goto Label13
					// line 1450: if not verbose:
					πF.SetLineno(1450)
				Label12:
					// line 1451: err += "; run in verbose mode for details"
					πF.SetLineno(1451)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µerr, πg.NewStr("; run in verbose mode for details").ToObject()); πE != nil {
						continue
					}
					µerr = πTemp001
					goto Label13
				Label13:
					goto Label11
				Label11:
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					πTemp003[0] = µerr
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTestFailed); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 1452: raise TestFailed(err)
					πF.SetLineno(1452)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					goto Label5
				Label5:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_run_suite.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 1455: def run_unittest(*classes):
			πF.SetLineno(1455)
			πTemp009 = make([]πg.Param, 0)
			πTemp018 = πg.NewFunction(πg.NewCode("run_unittest", "build/src/__python__/test/test_support.py", πTemp009, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µclasses *πg.Object = πArgs[0]; _ = µclasses
				var µvalid_types *πg.Object = πg.UnboundLocal; _ = µvalid_types
				var µsuite *πg.Object = πg.UnboundLocal; _ = µsuite
				var µcls *πg.Object = πg.UnboundLocal; _ = µcls
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
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
				var πTemp008 []*πg.Object
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
					// line 1456: """Run tests from unittest.TestCase-derived classes."""
					πF.SetLineno(1456)
					// line 1457: valid_types = (unittest.TestSuite, unittest.TestCase)
					πF.SetLineno(1457)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßTestSuite, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßTestCase, nil); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					µvalid_types = πTemp001
					// line 1458: suite = unittest.TestSuite()
					πF.SetLineno(1458)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßTestSuite, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µsuite = πTemp001
					if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µclasses); πE != nil {
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
					if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
						µcls = πTemp002
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(1)            
					πTemp007 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					πTemp007[0] = µcls
					if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					πTemp007[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label4
					}
					πTemp007 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					πTemp007[0] = µcls
					if πE = πg.CheckLocal(πF, µvalid_types, "valid_types"); πE != nil {
						continue
					}
					πTemp007[1] = µvalid_types
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label5
					}
					goto Label6
					// line 1460: if isinstance(cls, str):
					πF.SetLineno(1460)
				Label4:
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßmodules, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Contains(πF, πTemp004, µcls); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label8
					}
					goto Label9
					// line 1461: if cls in sys.modules:
					πF.SetLineno(1461)
				Label8:
					// line 1462: suite.addTest(unittest.findTestCases(sys.modules[cls]))
					πF.SetLineno(1462)
					πTemp007 = πF.MakeArgs(1)
					πTemp008 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					πTemp002 = µcls
					if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp004, ßmodules, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp009, πTemp002); πE != nil {
						continue
					}
					πTemp008[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßfindTestCases, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					πTemp007[0] = πTemp002
					if πE = πg.CheckLocal(πF, µsuite, "suite"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsuite, ßaddTest, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					goto Label10
				Label9:
					πTemp007 = πF.MakeArgs(1)
					πTemp007[0] = πg.NewStr("str arguments must be keys in sys.modules").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 1464: raise ValueError("str arguments must be keys in sys.modules")
					πF.SetLineno(1464)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label10
				Label10:
					goto Label7
					// line 1465: elif isinstance(cls, valid_types):
					πF.SetLineno(1465)
				Label5:
					// line 1466: suite.addTest(cls)
					πF.SetLineno(1466)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					πTemp007[0] = µcls
					if πE = πg.CheckLocal(πF, µsuite, "suite"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsuite, ßaddTest, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					goto Label7
				Label6:
					// line 1468: suite.addTest(unittest.makeSuite(cls))
					πF.SetLineno(1468)
					πTemp007 = πF.MakeArgs(1)
					πTemp008 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					πTemp008[0] = µcls
					if πTemp002, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmakeSuite, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					πTemp007[0] = πTemp002
					if πE = πg.CheckLocal(πF, µsuite, "suite"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsuite, ßaddTest, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					goto Label7
				Label7:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 1469: _run_suite(suite)
					πF.SetLineno(1469)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsuite, "suite"); πE != nil {
						continue
					}
					πTemp007[0] = µsuite
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_run_suite); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßrun_unittest.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 1525: def threading_setup():
			πF.SetLineno(1525)
			πTemp009 = make([]πg.Param, 0)
			πTemp019 = πg.NewFunction(πg.NewCode("threading_setup", "build/src/__python__/test/test_support.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
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
					if πTemp001, πE = πg.ResolveGlobal(πF, ßthread); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 1526: if thread:
					πF.SetLineno(1526)
				Label1:
					// line 1527: return (thread._count(),)
					πF.SetLineno(1527)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßthread); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß_count, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple1(πTemp003).ToObject()
					πR = πTemp001
					continue
					goto Label3
				Label2:
					// line 1529: return (1,)
					πF.SetLineno(1529)
					πTemp001 = πg.NewTuple1(πg.NewInt(1).ToObject()).ToObject()
					πR = πTemp001
					continue
					goto Label3
				Label3:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßthreading_setup.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 1531: def threading_cleanup(nb_threads):
			πF.SetLineno(1531)
			πTemp009 = make([]πg.Param, 1)
			πTemp009[0] = πg.Param{Name: "nb_threads", Def: nil}
			πTemp020 = πg.NewFunction(πg.NewCode("threading_cleanup", "build/src/__python__/test/test_support.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µnb_threads *πg.Object = πArgs[0]; _ = µnb_threads
				var µ_MAX_COUNT *πg.Object = πg.UnboundLocal; _ = µ_MAX_COUNT
				var µcount *πg.Object = πg.UnboundLocal; _ = µcount
				var µn *πg.Object = πg.UnboundLocal; _ = µn
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
				var πTemp006 bool
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 3: goto Label3
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßthread); πE != nil {
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
					// line 1532: if not thread:
					πF.SetLineno(1532)
				Label1:
					// line 1533: return
					πF.SetLineno(1533)
					πR = πg.None
					continue
					goto Label2
				Label2:
					// line 1535: _MAX_COUNT = 10
					πF.SetLineno(1535)
					µ_MAX_COUNT = πg.NewInt(10).ToObject()
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µ_MAX_COUNT, "_MAX_COUNT"); πE != nil {
						continue
					}
					πTemp004[0] = µ_MAX_COUNT
					if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
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
						πTemp006 = !isStop
					} else {
						πTemp006 = true
						µcount = πTemp002
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(3)            
					// line 1537: n = thread._count()
					πF.SetLineno(1537)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßthread); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp002, ß_count, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					µn = πTemp002
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µnb_threads, "nb_threads"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µn, µnb_threads); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label6
					}
					goto Label7
					// line 1538: if n == nb_threads:
					πF.SetLineno(1538)
				Label6:
					// line 1539: break
					πF.SetLineno(1539)
					πTemp003 = true
					continue
					goto Label7
				Label7:
					// line 1540: time.sleep(0.1)
					πF.SetLineno(1540)
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewFloat(0.1).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßsleep, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
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
			if πE = πF.Globals().SetItem(πF, ßthreading_cleanup.ToObject(), πTemp020); πE != nil {
				continue
			}
			// line 1560: def reap_children():
			πF.SetLineno(1560)
			πTemp009 = make([]πg.Param, 0)
			πTemp021 = πg.NewFunction(πg.NewCode("reap_children", "build/src/__python__/test/test_support.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µany_process *πg.Object = πg.UnboundLocal; _ = µany_process
				var µpid *πg.Object = πg.UnboundLocal; _ = µpid
				var µstatus *πg.Object = πg.UnboundLocal; _ = µstatus
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
					case 4: goto Label4
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					// line 1561: """Use this function at the end of test_main() whenever sub-processes
					πF.SetLineno(1561)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					πTemp001[1] = ßwaitpid.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 1569: if hasattr(os, 'waitpid'):
					πF.SetLineno(1569)
				Label1:
					// line 1570: any_process = -1
					πF.SetLineno(1570)
					if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µany_process = πTemp002
					// line 1571: while True:
					πF.SetLineno(1571)
					πF.PushCheckpoint(4)
					πTemp004 = false
				Label3:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label5
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
					πF.PushCheckpoint(3)            
					// line 1572: try:
					πF.SetLineno(1572)
					πF.PushCheckpoint(7)
					// line 1574: pid, status = os.waitpid(any_process, os.WNOHANG)
					πF.SetLineno(1574)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µany_process, "any_process"); πE != nil {
						continue
					}
					πTemp001[0] = µany_process
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßWNOHANG, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßwaitpid, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
						continue
					}
					µpid = πTemp003
					µstatus = πTemp006
					if πE = πg.CheckLocal(πF, µpid, "pid"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µpid, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label8
					}
					goto Label9
					// line 1575: if pid == 0:
					πF.SetLineno(1575)
				Label8:
					// line 1576: break
					πF.SetLineno(1576)
					πTemp004 = true
					continue
					goto Label9
				Label9:
					πF.PopCheckpoint()
					goto Label6
				Label7:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					goto Label10
					// line 1577: except:
					πF.SetLineno(1577)
				Label10:
					// line 1578: break
					πF.SetLineno(1578)
					πTemp004 = true
					continue
					πF.RestoreExc(nil, nil)
					goto Label6
				Label6:
					continue
				Label4:
					if πE != nil || πR != nil {
						continue
					}
				Label5:
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
			if πE = πF.Globals().SetItem(πF, ßreap_children.ToObject(), πTemp021); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("test.test_support", Code)
}
