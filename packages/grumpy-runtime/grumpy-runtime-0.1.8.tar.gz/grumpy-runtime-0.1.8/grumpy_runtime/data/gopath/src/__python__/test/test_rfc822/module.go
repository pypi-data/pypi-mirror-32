package test_rfc822
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/test/test_rfc822.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßFirst := πg.InternStr("First")
		ßImportError := πg.InternStr("ImportError")
		ßIndexError := πg.InternStr("IndexError")
		ßMessage := πg.InternStr("Message")
		ßMessageTestCase := πg.InternStr("MessageTestCase")
		ßNone := πg.InternStr("None")
		ßSecond := πg.InternStr("Second")
		ßStringIO := πg.InternStr("StringIO")
		ßTO := πg.InternStr("TO")
		ßTestCase := πg.InternStr("TestCase")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertTrue := πg.InternStr("assertTrue")
		ßcc := πg.InternStr("cc")
		ßcheck := πg.InternStr("check")
		ßcreate_message := πg.InternStr("create_message")
		ßdate := πg.InternStr("date")
		ßfrom := πg.InternStr("from")
		ßget := πg.InternStr("get")
		ßgetaddrlist := πg.InternStr("getaddrlist")
		ßgetdate := πg.InternStr("getdate")
		ßgetheader := πg.InternStr("getheader")
		ßhas_key := πg.InternStr("has_key")
		ßparseaddr := πg.InternStr("parseaddr")
		ßquote := πg.InternStr("quote")
		ßrepr := πg.InternStr("repr")
		ßrfc822 := πg.InternStr("rfc822")
		ßrun_unittest := πg.InternStr("run_unittest")
		ßsetdefault := πg.InternStr("setdefault")
		ßsort := πg.InternStr("sort")
		ßsorted := πg.InternStr("sorted")
		ßsubject := πg.InternStr("subject")
		ßtest_2getaddrlist := πg.InternStr("test_2getaddrlist")
		ßtest_addr_ipquad := πg.InternStr("test_addr_ipquad")
		ßtest_basic := πg.InternStr("test_basic")
		ßtest_bogus_to_header := πg.InternStr("test_bogus_to_header")
		ßtest_commas_in_full_name := πg.InternStr("test_commas_in_full_name")
		ßtest_doublecomment := πg.InternStr("test_doublecomment")
		ßtest_get := πg.InternStr("test_get")
		ßtest_invalid_headers := πg.InternStr("test_invalid_headers")
		ßtest_iter := πg.InternStr("test_iter")
		ßtest_main := πg.InternStr("test_main")
		ßtest_parseaddr := πg.InternStr("test_parseaddr")
		ßtest_quote_unquote := πg.InternStr("test_quote_unquote")
		ßtest_quoted_name := πg.InternStr("test_quoted_name")
		ßtest_rfc2822_phrases := πg.InternStr("test_rfc2822_phrases")
		ßtest_setdefault := πg.InternStr("test_setdefault")
		ßtest_support := πg.InternStr("test_support")
		ßtest_twisted := πg.InternStr("test_twisted")
		ßto := πg.InternStr("to")
		ßunittest := πg.InternStr("unittest")
		ßunquote := πg.InternStr("unquote")
		ßval := πg.InternStr("val")
		ßval2 := πg.InternStr("val2")
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
		var πTemp007 *πg.Dict
		_ = πTemp007
		var πTemp008 *πg.Object
		_ = πTemp008
		var πTemp009 *πg.Object
		_ = πTemp009
		var πTemp010 []πg.Param
		_ = πTemp010
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 2: goto Label2
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
			// line 2: from test import test_support
			πF.SetLineno(2)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßtest_support.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 5: import rfc822
			πF.SetLineno(5)
			if πTemp002, πE = πg.ImportModule(πF, "rfc822"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßrfc822.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 7: try:
			πF.SetLineno(7)
			πF.PushCheckpoint(2)
			// line 8: from cStringIO import StringIO
			πF.SetLineno(8)
			if πTemp002, πE = πg.ImportModule(πF, "cStringIO"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßStringIO, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStringIO.ToObject(), πTemp003); πE != nil {
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
			// line 9: except ImportError:
			πF.SetLineno(9)
		Label3:
			// line 10: from StringIO import StringIO
			πF.SetLineno(10)
			if πTemp002, πE = πg.ImportModule(πF, "StringIO"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßStringIO, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStringIO.ToObject(), πTemp003); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label1
		Label1:
			// line 13: class MessageTestCase(unittest.TestCase):
			πF.SetLineno(13)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßTestCase, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp007 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("MessageTestCase", "build/src/__python__/test/test_rfc822.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp007
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
				var πTemp017 *πg.Object
				_ = πTemp017
				var πTemp018 *πg.Object
				_ = πTemp018
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 14: def create_message(self, msg):
					πF.SetLineno(14)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "msg", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("create_message", "build/src/__python__/test/test_rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmsg *πg.Object = πArgs[1]; _ = µmsg
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
							// line 15: return rfc822.Message(StringIO(msg))
							πF.SetLineno(15)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp002[0] = µmsg
							if πTemp003, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrfc822); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßMessage, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcreate_message.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 17: def test_get(self):
					πF.SetLineno(17)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_get", "build/src/__python__/test/test_rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmsg *πg.Object = πg.UnboundLocal; _ = µmsg
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							// line 18: msg = self.create_message(
							πF.SetLineno(18)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("To: \"last, first\" <userid@foo.net>\n\ntest\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcreate_message, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmsg = πTemp003
							// line 20: self.assertTrue(msg.get("to") == '"last, first" <userid@foo.net>')
							πF.SetLineno(20)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßto.ToObject()
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmsg, ßget, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Eq(πF, πTemp005, πg.NewStr("\"last, first\" <userid@foo.net>").ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 21: self.assertTrue(msg.get("TO") == '"last, first" <userid@foo.net>')
							πF.SetLineno(21)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßTO.ToObject()
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmsg, ßget, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Eq(πF, πTemp005, πg.NewStr("\"last, first\" <userid@foo.net>").ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 22: self.assertTrue(msg.get("No-Such-Header") is None)
							πF.SetLineno(22)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("No-Such-Header").ToObject()
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmsg, ßget, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp005 == πTemp003).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 23: self.assertTrue(msg.get("No-Such-Header", "No-Such-Value")
							πF.SetLineno(23)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("No-Such-Header").ToObject()
							πTemp004[1] = πg.NewStr("No-Such-Value").ToObject()
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmsg, ßget, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Eq(πF, πTemp005, πg.NewStr("No-Such-Value").ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_get.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 26: def test_setdefault(self):
					πF.SetLineno(26)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_setdefault", "build/src/__python__/test/test_rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmsg *πg.Object = πg.UnboundLocal; _ = µmsg
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
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
							default: panic("unexpected function state")
							}
							// line 27: msg = self.create_message(
							πF.SetLineno(27)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("To: \"last, first\" <userid@foo.net>\n\ntest\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcreate_message, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmsg = πTemp003
							// line 29: self.assertTrue(not msg.has_key("New-Header"))
							πF.SetLineno(29)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("New-Header").ToObject()
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmsg, ßhas_key, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp006).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 30: self.assertTrue(msg.setdefault("New-Header", "New-Value") == "New-Value")
							πF.SetLineno(30)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("New-Header").ToObject()
							πTemp004[1] = πg.NewStr("New-Value").ToObject()
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmsg, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Eq(πF, πTemp005, πg.NewStr("New-Value").ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 31: self.assertTrue(msg.setdefault("New-Header", "Different-Value")
							πF.SetLineno(31)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("New-Header").ToObject()
							πTemp004[1] = πg.NewStr("Different-Value").ToObject()
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmsg, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Eq(πF, πTemp005, πg.NewStr("New-Value").ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 33: self.assertTrue(msg["new-header"] == "New-Value")
							πF.SetLineno(33)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πg.NewStr("new-header").ToObject()
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µmsg, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp005, πg.NewStr("New-Value").ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 35: self.assertTrue(msg.setdefault("Another-Header") == "")
							πF.SetLineno(35)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("Another-Header").ToObject()
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmsg, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Eq(πF, πTemp005, ß.ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 36: self.assertTrue(msg["another-header"] == "")
							πF.SetLineno(36)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πg.NewStr("another-header").ToObject()
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µmsg, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp005, ß.ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_setdefault.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 38: def check(self, msg, results):
					πF.SetLineno(38)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "msg", Def: nil}
					πTemp002[2] = πg.Param{Name: "results", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("check", "build/src/__python__/test/test_rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmsg *πg.Object = πArgs[1]; _ = µmsg
						var µresults *πg.Object = πArgs[2]; _ = µresults
						var µm *πg.Object = πg.UnboundLocal; _ = µm
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var µa *πg.Object = πg.UnboundLocal; _ = µa
						var µmn *πg.Object = πg.UnboundLocal; _ = µmn
						var µma *πg.Object = πg.UnboundLocal; _ = µma
						var µout *πg.Object = πg.UnboundLocal; _ = µout
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
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 *πg.BaseException
						_ = πTemp011
						var πTemp012 *πg.Traceback
						_ = πTemp012
						var πTemp013 []*πg.Object
						_ = πTemp013
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							case 5: goto Label5
							default: panic("unexpected function state")
							}
							// line 39: """Check addresses and the date."""
							πF.SetLineno(39)
							// line 40: m = self.create_message(msg)
							πF.SetLineno(40)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp001[0] = µmsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcreate_message, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µm = πTemp003
							// line 41: i = 0
							πF.SetLineno(41)
							µi = πg.NewInt(0).ToObject()
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßto.ToObject()
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µm, ßgetaddrlist, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßcc.ToObject()
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µm, ßgetaddrlist, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.Add(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								µn = πTemp004
								µa = πTemp005
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 43: try:
							πF.SetLineno(43)
							πF.PushCheckpoint(5)
							// line 44: mn, ma = results[i][0], results[i][1]
							πF.SetLineno(44)
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp006 = µi
							if πE = πg.CheckLocal(πF, µresults, "results"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µresults, πTemp006); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
								continue
							}
							πTemp004 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp009 = µi
							if πE = πg.CheckLocal(πF, µresults, "results"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µresults, πTemp009); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp010, πTemp004); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
								continue
							}
							µmn = πTemp004
							µma = πTemp005
							πF.PopCheckpoint()
							goto Label4
						Label5:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp011, πTemp012 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label6
							}
							πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
							continue
							// line 45: except IndexError:
							πF.SetLineno(45)
						Label6:
							// line 46: print 'extra parsed address:', repr(n), repr(a)
							πF.SetLineno(46)
							πTemp001 = make([]*πg.Object, 3)
							πTemp001[0] = πg.NewStr("extra parsed address:").ToObject()
							πTemp013 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp013[0] = µn
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp013, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp013)
							πTemp001[1] = πTemp004
							πTemp013 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp013[0] = µa
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp013, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp013)
							πTemp001[2] = πTemp004
							if πE = πg.Print(πF, πTemp001, true); πE != nil {
								continue
							}
							// line 47: continue
							πF.SetLineno(47)
							continue
							πF.RestoreExc(nil, nil)
							goto Label4
						Label4:
							// line 48: i = i + 1
							πF.SetLineno(48)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µi = πTemp003
							// line 49: self.assertEqual(mn, n,
							πF.SetLineno(49)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µmn, "mn"); πE != nil {
								continue
							}
							πTemp001[0] = µmn
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[1] = µn
							if πE = πg.CheckLocal(πF, µmn, "mn"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(µmn, µn).ToObject()
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("Un-expected name: %r != %r").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp001[2] = πTemp003
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
							// line 51: self.assertEqual(ma, a,
							πF.SetLineno(51)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µma, "ma"); πE != nil {
								continue
							}
							πTemp001[0] = µma
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[1] = µa
							if πE = πg.CheckLocal(πF, µma, "ma"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(µma, µa).ToObject()
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("Un-expected address: %r != %r").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp001[2] = πTemp003
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
							if πE = πg.CheckLocal(πF, µmn, "mn"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Eq(πF, µmn, µn); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp008 {
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µma, "ma"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Eq(πF, µma, µa); πE != nil {
								continue
							}
							πTemp003 = πTemp004
						Label7:
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label8
							}
							goto Label9
							// line 53: if mn == n and ma == a:
							πF.SetLineno(53)
						Label8:
							// line 54: pass
							πF.SetLineno(54)
							goto Label10
						Label9:
							// line 56: print 'not found:', repr(n), repr(a)
							πF.SetLineno(56)
							πTemp001 = make([]*πg.Object, 3)
							πTemp001[0] = πg.NewStr("not found:").ToObject()
							πTemp013 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp013[0] = µn
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp013, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp013)
							πTemp001[1] = πTemp004
							πTemp013 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp013[0] = µa
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp013, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp013)
							πTemp001[2] = πTemp004
							if πE = πg.Print(πF, πTemp001, true); πE != nil {
								continue
							}
							goto Label10
						Label10:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 58: out = m.getdate('date')
							πF.SetLineno(58)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßdate.ToObject()
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µm, ßgetdate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µout = πTemp003
							if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µout); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label11
							}
							goto Label12
							// line 59: if out:
							πF.SetLineno(59)
						Label11:
							// line 60: self.assertEqual(out,
							πF.SetLineno(60)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
								continue
							}
							πTemp001[0] = µout
							πTemp013 = make([]*πg.Object, 9)
							πTemp013[0] = πg.NewInt(1999).ToObject()
							πTemp013[1] = πg.NewInt(1).ToObject()
							πTemp013[2] = πg.NewInt(13).ToObject()
							πTemp013[3] = πg.NewInt(23).ToObject()
							πTemp013[4] = πg.NewInt(57).ToObject()
							πTemp013[5] = πg.NewInt(35).ToObject()
							πTemp013[6] = πg.NewInt(0).ToObject()
							πTemp013[7] = πg.NewInt(1).ToObject()
							πTemp013[8] = πg.NewInt(0).ToObject()
							πTemp002 = πg.NewTuple(πTemp013...).ToObject()
							πTemp001[1] = πTemp002
							πTemp001[2] = πg.NewStr("date conversion failed").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label12
						Label12:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcheck.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 68: def test_basic(self):
					πF.SetLineno(68)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_basic", "build/src/__python__/test/test_rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 69: self.check(
							πF.SetLineno(69)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("Date:    Wed, 13 Jan 1999 23:57:35 -0500\nFrom:    Guido van Rossum <guido@CNRI.Reston.VA.US>\nTo:      \"Guido van\n\t : Rossum\" <guido@python.org>\nSubject: test2\n\ntest2\n").ToObject()
							πTemp002 = make([]*πg.Object, 1)
							πTemp003 = πg.NewTuple2(πg.NewStr("Guido van\n\t : Rossum").ToObject(), πg.NewStr("guido@python.org").ToObject()).ToObject()
							πTemp002[0] = πTemp003
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheck, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 79: self.check(
							πF.SetLineno(79)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("From: Barry <bwarsaw@python.org\nTo: guido@python.org (Guido: the Barbarian)\nSubject: nonsense\nDate: Wednesday, January 13 1999 23:57:35 -0500\n\ntest").ToObject()
							πTemp002 = make([]*πg.Object, 1)
							πTemp003 = πg.NewTuple2(πg.NewStr("Guido: the Barbarian").ToObject(), πg.NewStr("guido@python.org").ToObject()).ToObject()
							πTemp002[0] = πTemp003
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheck, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 88: self.check(
							πF.SetLineno(88)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("From: Barry <bwarsaw@python.org\nTo: guido@python.org (Guido: the Barbarian)\nCc: \"Guido: the Madman\" <guido@python.org>\nDate:  13-Jan-1999 23:57:35 EST\n\ntest").ToObject()
							πTemp002 = make([]*πg.Object, 2)
							πTemp003 = πg.NewTuple2(πg.NewStr("Guido: the Barbarian").ToObject(), πg.NewStr("guido@python.org").ToObject()).ToObject()
							πTemp002[0] = πTemp003
							πTemp003 = πg.NewTuple2(πg.NewStr("Guido: the Madman").ToObject(), πg.NewStr("guido@python.org").ToObject()).ToObject()
							πTemp002[1] = πTemp003
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheck, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 99: self.check(
							πF.SetLineno(99)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("To: \"The monster with\n     the very long name: Guido\" <guido@python.org>\nDate:    Wed, 13 Jan 1999 23:57:35 -0500\n\ntest").ToObject()
							πTemp002 = make([]*πg.Object, 1)
							πTemp003 = πg.NewTuple2(πg.NewStr("The monster with\n     the very long name: Guido").ToObject(), πg.NewStr("guido@python.org").ToObject()).ToObject()
							πTemp002[0] = πTemp003
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheck, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 108: self.check(
							πF.SetLineno(108)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("To: \"Amit J. Patel\" <amitp@Theory.Stanford.EDU>\nCC: Mike Fletcher <mfletch@vrtelecom.com>,\n        \"'string-sig@python.org'\" <string-sig@python.org>\nCc: fooz@bat.com, bart@toof.com\nCc: goit@lip.com\nDate:    Wed, 13 Jan 1999 23:57:35 -0500\n\ntest").ToObject()
							πTemp002 = make([]*πg.Object, 6)
							πTemp003 = πg.NewTuple2(πg.NewStr("Amit J. Patel").ToObject(), πg.NewStr("amitp@Theory.Stanford.EDU").ToObject()).ToObject()
							πTemp002[0] = πTemp003
							πTemp003 = πg.NewTuple2(πg.NewStr("Mike Fletcher").ToObject(), πg.NewStr("mfletch@vrtelecom.com").ToObject()).ToObject()
							πTemp002[1] = πTemp003
							πTemp003 = πg.NewTuple2(πg.NewStr("'string-sig@python.org'").ToObject(), πg.NewStr("string-sig@python.org").ToObject()).ToObject()
							πTemp002[2] = πTemp003
							πTemp003 = πg.NewTuple2(ß.ToObject(), πg.NewStr("fooz@bat.com").ToObject()).ToObject()
							πTemp002[3] = πTemp003
							πTemp003 = πg.NewTuple2(ß.ToObject(), πg.NewStr("bart@toof.com").ToObject()).ToObject()
							πTemp002[4] = πTemp003
							πTemp003 = πg.NewTuple2(ß.ToObject(), πg.NewStr("goit@lip.com").ToObject()).ToObject()
							πTemp002[5] = πTemp003
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheck, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 125: self.check(
							πF.SetLineno(125)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("To: Some One <someone@dom.ain>\nFrom: Anudder Persin <subuddy.else@dom.ain>\nDate:\n\ntest").ToObject()
							πTemp002 = make([]*πg.Object, 1)
							πTemp003 = πg.NewTuple2(πg.NewStr("Some One").ToObject(), πg.NewStr("someone@dom.ain").ToObject()).ToObject()
							πTemp002[0] = πTemp003
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheck, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 133: self.check(
							πF.SetLineno(133)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("To: person@dom.ain (User J. Person)\n\n").ToObject()
							πTemp002 = make([]*πg.Object, 1)
							πTemp003 = πg.NewTuple2(πg.NewStr("User J. Person").ToObject(), πg.NewStr("person@dom.ain").ToObject()).ToObject()
							πTemp002[0] = πTemp003
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheck, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_basic.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 137: def test_doublecomment(self):
					πF.SetLineno(137)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_doublecomment", "build/src/__python__/test/test_rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 139: self.check(
							πF.SetLineno(139)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("To: person@dom.ain ((User J. Person)), John Doe <foo@bar.com>\n\n").ToObject()
							πTemp002 = make([]*πg.Object, 2)
							πTemp003 = πg.NewTuple2(πg.NewStr("User J. Person").ToObject(), πg.NewStr("person@dom.ain").ToObject()).ToObject()
							πTemp002[0] = πTemp003
							πTemp003 = πg.NewTuple2(πg.NewStr("John Doe").ToObject(), πg.NewStr("foo@bar.com").ToObject()).ToObject()
							πTemp002[1] = πTemp003
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheck, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_doublecomment.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 143: def test_twisted(self):
					πF.SetLineno(143)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("test_twisted", "build/src/__python__/test/test_rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 147: self.check(
							πF.SetLineno(147)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("To: <[smtp:dd47@mail.xxx.edu]_at_hmhq@hdq-mdm1-imgout.companay.com>\nDate:    Wed, 13 Jan 1999 23:57:35 -0500\n\ntest").ToObject()
							πTemp002 = make([]*πg.Object, 3)
							πTemp003 = πg.NewTuple2(ß.ToObject(), ß.ToObject()).ToObject()
							πTemp002[0] = πTemp003
							πTemp003 = πg.NewTuple2(ß.ToObject(), πg.NewStr("dd47@mail.xxx.edu").ToObject()).ToObject()
							πTemp002[1] = πTemp003
							πTemp003 = πg.NewTuple2(ß.ToObject(), πg.NewStr("_at_hmhq@hdq-mdm1-imgout.companay.com").ToObject()).ToObject()
							πTemp002[2] = πTemp003
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheck, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_twisted.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 157: def test_commas_in_full_name(self):
					πF.SetLineno(157)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("test_commas_in_full_name", "build/src/__python__/test/test_rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 161: self.check(
							πF.SetLineno(161)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("To: \"last, first\" <userid@foo.net>\n\ntest").ToObject()
							πTemp002 = make([]*πg.Object, 1)
							πTemp003 = πg.NewTuple2(πg.NewStr("last, first").ToObject(), πg.NewStr("userid@foo.net").ToObject()).ToObject()
							πTemp002[0] = πTemp003
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheck, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_commas_in_full_name.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 167: def test_quoted_name(self):
					πF.SetLineno(167)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("test_quoted_name", "build/src/__python__/test/test_rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 168: self.check(
							πF.SetLineno(168)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("To: (Comment stuff) \"Quoted name\"@somewhere.com\n\ntest").ToObject()
							πTemp002 = make([]*πg.Object, 1)
							πTemp003 = πg.NewTuple2(πg.NewStr("Comment stuff").ToObject(), πg.NewStr("\"Quoted name\"@somewhere.com").ToObject()).ToObject()
							πTemp002[0] = πTemp003
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheck, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_quoted_name.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 174: def test_bogus_to_header(self):
					πF.SetLineno(174)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("test_bogus_to_header", "build/src/__python__/test/test_rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 175: self.check(
							πF.SetLineno(175)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("To: :\nCc: goit@lip.com\nDate:    Wed, 13 Jan 1999 23:57:35 -0500\n\ntest").ToObject()
							πTemp002 = make([]*πg.Object, 1)
							πTemp003 = πg.NewTuple2(ß.ToObject(), πg.NewStr("goit@lip.com").ToObject()).ToObject()
							πTemp002[0] = πTemp003
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheck, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_bogus_to_header.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 183: def test_addr_ipquad(self):
					πF.SetLineno(183)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("test_addr_ipquad", "build/src/__python__/test/test_rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 184: self.check(
							πF.SetLineno(184)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("To: guido@[132.151.1.21]\n\nfoo").ToObject()
							πTemp002 = make([]*πg.Object, 1)
							πTemp003 = πg.NewTuple2(ß.ToObject(), πg.NewStr("guido@[132.151.1.21]").ToObject()).ToObject()
							πTemp002[0] = πTemp003
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheck, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_addr_ipquad.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 190: def test_iter(self):
					πF.SetLineno(190)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("test_iter", "build/src/__python__/test/test_rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µm *πg.Object = πg.UnboundLocal; _ = µm
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
							// line 191: m = rfc822.Message(StringIO(
							πF.SetLineno(191)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("Date:    Wed, 13 Jan 1999 23:57:35 -0500\nFrom:    Guido van Rossum <guido@CNRI.Reston.VA.US>\nTo:      \"Guido van\n\t : Rossum\" <guido@python.org>\nSubject: test2\n\ntest2\n").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrfc822); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßMessage, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µm = πTemp003
							// line 199: self.assertEqual(sorted(m), ['date', 'from', 'subject', 'to'])
							πF.SetLineno(199)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							πTemp002[0] = µm
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp002 = make([]*πg.Object, 4)
							πTemp002[0] = ßdate.ToObject()
							πTemp002[1] = ßfrom.ToObject()
							πTemp002[2] = ßsubject.ToObject()
							πTemp002[3] = ßto.ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_iter.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 201: def test_rfc2822_phrases(self):
					πF.SetLineno(201)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("test_rfc2822_phrases", "build/src/__python__/test/test_rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 207: self.check('To: User J. Person <person@dom.ain>\n\n',
							πF.SetLineno(207)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("To: User J. Person <person@dom.ain>\n\n").ToObject()
							πTemp002 = make([]*πg.Object, 1)
							πTemp003 = πg.NewTuple2(πg.NewStr("User J. Person").ToObject(), πg.NewStr("person@dom.ain").ToObject()).ToObject()
							πTemp002[0] = πTemp003
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheck, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_rfc2822_phrases.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 218: def test_2getaddrlist(self):
					πF.SetLineno(218)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("test_2getaddrlist", "build/src/__python__/test/test_rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µeq *πg.Object = πg.UnboundLocal; _ = µeq
						var µmsg *πg.Object = πg.UnboundLocal; _ = µmsg
						var µccs *πg.Object = πg.UnboundLocal; _ = µccs
						var µaddrs *πg.Object = πg.UnboundLocal; _ = µaddrs
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []πg.Param
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
							// line 219: eq = self.assertEqual
							πF.SetLineno(219)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							µeq = πTemp001
							// line 220: msg = self.create_message("""\
							πF.SetLineno(220)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("To: aperson@dom.ain\nCc: bperson@dom.ain\nCc: cperson@dom.ain\nCc: dperson@dom.ain\n\nA test message.\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcreate_message, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µmsg = πTemp003
							// line 228: ccs = [('', a) for a in
							πF.SetLineno(228)
							πTemp004 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "build/src/__python__/test/test_rfc822.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µa *πg.Object = πg.UnboundLocal; _ = µa
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
										πTemp002 = make([]*πg.Object, 3)
										πTemp002[0] = πg.NewStr("bperson@dom.ain").ToObject()
										πTemp002[1] = πg.NewStr("cperson@dom.ain").ToObject()
										πTemp002[2] = πg.NewStr("dperson@dom.ain").ToObject()
										πTemp003 = πg.NewList(πTemp002...).ToObject()
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
											µa = πTemp003
										}
										if πE != nil || !πTemp005 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 228: ccs = [('', a) for a in
										πF.SetLineno(228)
										if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
											continue
										}
										πTemp003 = πg.NewTuple2(ß.ToObject(), µa).ToObject()
										πF.PushCheckpoint(4)
										return πTemp003, nil
									Label4:
										πTemp006 = πSent
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
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
								continue
							}
							µccs = πTemp001
							// line 230: addrs = msg.getaddrlist('cc')
							πF.SetLineno(230)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßcc.ToObject()
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmsg, ßgetaddrlist, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µaddrs = πTemp005
							// line 231: addrs.sort()
							πF.SetLineno(231)
							if πE = πg.CheckLocal(πF, µaddrs, "addrs"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µaddrs, ßsort, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 232: eq(addrs, ccs)
							πF.SetLineno(232)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µaddrs, "addrs"); πE != nil {
								continue
							}
							πTemp002[0] = µaddrs
							if πE = πg.CheckLocal(πF, µccs, "ccs"); πE != nil {
								continue
							}
							πTemp002[1] = µccs
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 234: addrs = msg.getaddrlist('cc')
							πF.SetLineno(234)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßcc.ToObject()
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmsg, ßgetaddrlist, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µaddrs = πTemp005
							// line 235: addrs.sort()
							πF.SetLineno(235)
							if πE = πg.CheckLocal(πF, µaddrs, "addrs"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µaddrs, ßsort, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 236: eq(addrs, ccs)
							πF.SetLineno(236)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µaddrs, "addrs"); πE != nil {
								continue
							}
							πTemp002[0] = µaddrs
							if πE = πg.CheckLocal(πF, µccs, "ccs"); πE != nil {
								continue
							}
							πTemp002[1] = µccs
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_2getaddrlist.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 238: def test_parseaddr(self):
					πF.SetLineno(238)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("test_parseaddr", "build/src/__python__/test/test_rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µeq *πg.Object = πg.UnboundLocal; _ = µeq
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 239: eq = self.assertEqual
							πF.SetLineno(239)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							µeq = πTemp001
							// line 240: eq(rfc822.parseaddr('<>'), ('', ''))
							πF.SetLineno(240)
							πTemp002 = πF.MakeArgs(2)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("<>").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßrfc822); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßparseaddr, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002[0] = πTemp001
							πTemp001 = πg.NewTuple2(ß.ToObject(), ß.ToObject()).ToObject()
							πTemp002[1] = πTemp001
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 241: eq(rfc822.parseaddr('aperson@dom.ain'), ('', 'aperson@dom.ain'))
							πF.SetLineno(241)
							πTemp002 = πF.MakeArgs(2)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("aperson@dom.ain").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßrfc822); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßparseaddr, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002[0] = πTemp001
							πTemp001 = πg.NewTuple2(ß.ToObject(), πg.NewStr("aperson@dom.ain").ToObject()).ToObject()
							πTemp002[1] = πTemp001
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 242: eq(rfc822.parseaddr('bperson@dom.ain (Bea A. Person)'),
							πF.SetLineno(242)
							πTemp002 = πF.MakeArgs(2)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("bperson@dom.ain (Bea A. Person)").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßrfc822); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßparseaddr, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002[0] = πTemp001
							πTemp001 = πg.NewTuple2(πg.NewStr("Bea A. Person").ToObject(), πg.NewStr("bperson@dom.ain").ToObject()).ToObject()
							πTemp002[1] = πTemp001
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 244: eq(rfc822.parseaddr('Cynthia Person <cperson@dom.ain>'),
							πF.SetLineno(244)
							πTemp002 = πF.MakeArgs(2)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("Cynthia Person <cperson@dom.ain>").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßrfc822); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßparseaddr, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002[0] = πTemp001
							πTemp001 = πg.NewTuple2(πg.NewStr("Cynthia Person").ToObject(), πg.NewStr("cperson@dom.ain").ToObject()).ToObject()
							πTemp002[1] = πTemp001
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_parseaddr.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 247: def test_quote_unquote(self):
					πF.SetLineno(247)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("test_quote_unquote", "build/src/__python__/test/test_rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µeq *πg.Object = πg.UnboundLocal; _ = µeq
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 248: eq = self.assertEqual
							πF.SetLineno(248)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							µeq = πTemp001
							// line 249: eq(rfc822.quote('foo\\wacky"name'), 'foo\\\\wacky\\"name')
							πF.SetLineno(249)
							πTemp002 = πF.MakeArgs(2)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("foo\\wacky\"name").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßrfc822); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßquote, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002[0] = πTemp001
							πTemp002[1] = πg.NewStr("foo\\\\wacky\\\"name").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 250: eq(rfc822.unquote('"foo\\\\wacky\\"name"'), 'foo\\wacky"name')
							πF.SetLineno(250)
							πTemp002 = πF.MakeArgs(2)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("\"foo\\\\wacky\\\"name\"").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßrfc822); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßunquote, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002[0] = πTemp001
							πTemp002[1] = πg.NewStr("foo\\wacky\"name").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_quote_unquote.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 252: def test_invalid_headers(self):
					πF.SetLineno(252)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("test_invalid_headers", "build/src/__python__/test/test_rfc822.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µeq *πg.Object = πg.UnboundLocal; _ = µeq
						var µmsg *πg.Object = πg.UnboundLocal; _ = µmsg
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
							// line 253: eq = self.assertEqual
							πF.SetLineno(253)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							µeq = πTemp001
							// line 254: msg = self.create_message("First: val\n: otherval\nSecond: val2\n")
							πF.SetLineno(254)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("First: val\n: otherval\nSecond: val2\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcreate_message, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µmsg = πTemp003
							// line 255: eq(msg.getheader('First'), 'val')
							πF.SetLineno(255)
							πTemp002 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßFirst.ToObject()
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmsg, ßgetheader, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002[0] = πTemp003
							πTemp002[1] = ßval.ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 256: eq(msg.getheader('Second'), 'val2')
							πF.SetLineno(256)
							πTemp002 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßSecond.ToObject()
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmsg, ßgetheader, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002[0] = πTemp003
							πTemp002[1] = ßval2.ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_invalid_headers.ToObject(), πTemp018); πE != nil {
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
			if πTemp008, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("MessageTestCase").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMessageTestCase.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 259: def test_main():
			πF.SetLineno(259)
			πTemp010 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("test_main", "build/src/__python__/test/test_rfc822.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 260: test_support.run_unittest(MessageTestCase)
					πF.SetLineno(260)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßMessageTestCase); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßtest_main.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp003, πE = πg.Eq(πF, πTemp008, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label4
			}
			goto Label5
			// line 263: if __name__ == "__main__":
			πF.SetLineno(263)
		Label4:
			// line 264: test_main()
			πF.SetLineno(264)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_main); πE != nil {
				continue
			}
			if πTemp008, πE = πTemp003.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label5
		Label5:
		}
		return nil, πE
	})
	πg.RegisterModule("test.test_rfc822", Code)
}
