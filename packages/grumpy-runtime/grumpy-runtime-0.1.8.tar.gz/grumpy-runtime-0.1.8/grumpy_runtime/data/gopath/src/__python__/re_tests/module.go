package re_tests
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/re_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ß1234abc5678 := πg.InternStr("1234abc5678")
		ß123abc := πg.InternStr("123abc")
		ß9 := πg.InternStr("9")
		ßA := πg.InternStr("A")
		ßAA := πg.InternStr("AA")
		ßAABBABC := πg.InternStr("AABBABC")
		ßAABC := πg.InternStr("AABC")
		ßAAC := πg.InternStr("AAC")
		ßAB := πg.InternStr("AB")
		ßABABC := πg.InternStr("ABABC")
		ßABBBBC := πg.InternStr("ABBBBC")
		ßABBBCD := πg.InternStr("ABBBCD")
		ßABBC := πg.InternStr("ABBC")
		ßABC := πg.InternStr("ABC")
		ßABCABC := πg.InternStr("ABCABC")
		ßABCC := πg.InternStr("ABCC")
		ßABCD := πg.InternStr("ABCD")
		ßABCDE := πg.InternStr("ABCDE")
		ßABCDEF := πg.InternStr("ABCDEF")
		ßABCDEFG := πg.InternStr("ABCDEFG")
		ßABD := πg.InternStr("ABD")
		ßABH := πg.InternStr("ABH")
		ßABQ := πg.InternStr("ABQ")
		ßABX := πg.InternStr("ABX")
		ßAC := πg.InternStr("AC")
		ßACE := πg.InternStr("ACE")
		ßADC := πg.InternStr("ADC")
		ßADCDCDE := πg.InternStr("ADCDCDE")
		ßAED := πg.InternStr("AED")
		ßALPHA := πg.InternStr("ALPHA")
		ßAXC := πg.InternStr("AXC")
		ßAXYZC := πg.InternStr("AXYZC")
		ßAXYZD := πg.InternStr("AXYZD")
		ßAssertionError := πg.InternStr("AssertionError")
		ßB := πg.InternStr("B")
		ßBCDD := πg.InternStr("BCDD")
		ßC := πg.InternStr("C")
		ßCDE := πg.InternStr("CDE")
		ßDEF := πg.InternStr("DEF")
		ßE := πg.InternStr("E")
		ßEF := πg.InternStr("EF")
		ßEFFG := πg.InternStr("EFFG")
		ßEFFGZ := πg.InternStr("EFFGZ")
		ßException := πg.InternStr("Exception")
		ßFAIL := πg.InternStr("FAIL")
		ßHIJ := πg.InternStr("HIJ")
		ßIJ := πg.InternStr("IJ")
		ßMMM := πg.InternStr("MMM")
		ßREFFGZ := πg.InternStr("REFFGZ")
		ßSUCCEED := πg.InternStr("SUCCEED")
		ßSYNTAX_ERROR := πg.InternStr("SYNTAX_ERROR")
		ßW := πg.InternStr("W")
		ßXABCY := πg.InternStr("XABCY")
		ßXABYABBBZ := πg.InternStr("XABYABBBZ")
		ßXAYABBBZ := πg.InternStr("XAYABBBZ")
		ßXBC := πg.InternStr("XBC")
		ßa := πg.InternStr("a")
		ßaa := πg.InternStr("aa")
		ßaaa := πg.InternStr("aaa")
		ßaaaa := πg.InternStr("aaaa")
		ßaaaaa := πg.InternStr("aaaaa")
		ßaaax := πg.InternStr("aaax")
		ßaabbabc := πg.InternStr("aabbabc")
		ßaabc := πg.InternStr("aabc")
		ßaac := πg.InternStr("aac")
		ßaacx := πg.InternStr("aacx")
		ßab := πg.InternStr("ab")
		ßabNNxyz := πg.InternStr("abNNxyz")
		ßabNNxyzN := πg.InternStr("abNNxyzN")
		ßab_cd0123 := πg.InternStr("ab_cd0123")
		ßaba := πg.InternStr("aba")
		ßababc := πg.InternStr("ababc")
		ßabad := πg.InternStr("abad")
		ßabbbbc := πg.InternStr("abbbbc")
		ßabbbcd := πg.InternStr("abbbcd")
		ßabbc := πg.InternStr("abbc")
		ßabc := πg.InternStr("abc")
		ßabcabc := πg.InternStr("abcabc")
		ßabcc := πg.InternStr("abcc")
		ßabcd := πg.InternStr("abcd")
		ßabcde := πg.InternStr("abcde")
		ßabcdef := πg.InternStr("abcdef")
		ßabcdefdof := πg.InternStr("abcdefdof")
		ßabcdefg := πg.InternStr("abcdefg")
		ßabcdefghijklk9 := πg.InternStr("abcdefghijklk9")
		ßabcx := πg.InternStr("abcx")
		ßabd := πg.InternStr("abd")
		ßabh := πg.InternStr("abh")
		ßabq := πg.InternStr("abq")
		ßabx := πg.InternStr("abx")
		ßac := πg.InternStr("ac")
		ßacb := πg.InternStr("acb")
		ßace := πg.InternStr("ace")
		ßad := πg.InternStr("ad")
		ßadc := πg.InternStr("adc")
		ßadcdcde := πg.InternStr("adcdcde")
		ßaed := πg.InternStr("aed")
		ßalpha := πg.InternStr("alpha")
		ßaxc := πg.InternStr("axc")
		ßaxyzc := πg.InternStr("axyzc")
		ßaxyzd := πg.InternStr("axyzd")
		ßb := πg.InternStr("b")
		ßbcdd := πg.InternStr("bcdd")
		ßcde := πg.InternStr("cde")
		ßce := πg.InternStr("ce")
		ßceghijkmopqyz := πg.InternStr("ceghijkmopqyz")
		ßchr := πg.InternStr("chr")
		ßdef := πg.InternStr("def")
		ßdof := πg.InternStr("dof")
		ße := πg.InternStr("e")
		ßef := πg.InternStr("ef")
		ßeffg := πg.InternStr("effg")
		ßeffgz := πg.InternStr("effgz")
		ßfoo := πg.InternStr("foo")
		ßfound := πg.InternStr("found")
		ßg1 := πg.InternStr("g1")
		ßg10 := πg.InternStr("g10")
		ßhij := πg.InternStr("hij")
		ßij := πg.InternStr("ij")
		ßlaser_beam := πg.InternStr("laser_beam")
		ßpattern := πg.InternStr("pattern")
		ßrange := πg.InternStr("range")
		ßre := πg.InternStr("re")
		ßreffgz := πg.InternStr("reffgz")
		ßs := πg.InternStr("s")
		ßsearch := πg.InternStr("search")
		ßsmil := πg.InternStr("smil")
		ßstatus := πg.InternStr("status")
		ßtest := πg.InternStr("test")
		ßtests := πg.InternStr("tests")
		ßw := πg.InternStr("w")
		ßwxyz := πg.InternStr("wxyz")
		ßx := πg.InternStr("x")
		ßxabcy := πg.InternStr("xabcy")
		ßxabyabbbz := πg.InternStr("xabyabbbz")
		ßxayabbbz := πg.InternStr("xayabbbz")
		ßxbc := πg.InternStr("xbc")
		ßxg1y := πg.InternStr("xg1y")
		ßxy := πg.InternStr("xy")
		ßxyz := πg.InternStr("xyz")
		ßyz := πg.InternStr("yz")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.Object
		_ = πTemp004
		var πTemp005 *πg.Object
		_ = πTemp005
		var πTemp006 []*πg.Object
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
		var πTemp013 bool
		_ = πTemp013
		var πTemp014 bool
		_ = πTemp014
		var πTemp015 *πg.BaseException
		_ = πTemp015
		var πTemp016 *πg.Traceback
		_ = πTemp016
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 1: goto Label1
			case 2: goto Label2
			case 9: goto Label9
			default: panic("unexpected function state")
			}
			// line 4: import re
			πF.SetLineno(4)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 9: [SUCCEED, FAIL, SYNTAX_ERROR] = range(3)
			πF.SetLineno(9)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewInt(3).ToObject()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSUCCEED.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFAIL.ToObject(), πTemp004); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSYNTAX_ERROR.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 58: tests = [
			πF.SetLineno(58)
			πTemp002 = make([]*πg.Object, 505)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?P<foo_123").ToObject(), ß.ToObject(), πTemp003).ToObject()
			πTemp002[0] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?P<1>a)").ToObject(), ß.ToObject(), πTemp003).ToObject()
			πTemp002[1] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?P<!>a)").ToObject(), ß.ToObject(), πTemp003).ToObject()
			πTemp002[2] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?P<foo!>a)").ToObject(), ß.ToObject(), πTemp003).ToObject()
			πTemp002[3] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?P<foo_123>a)(?P=foo_123").ToObject(), ßaa.ToObject(), πTemp003).ToObject()
			πTemp002[4] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?P<foo_123>a)(?P=1)").ToObject(), ßaa.ToObject(), πTemp003).ToObject()
			πTemp002[5] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?P<foo_123>a)(?P=!)").ToObject(), ßaa.ToObject(), πTemp003).ToObject()
			πTemp002[6] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?P<foo_123>a)(?P=foo_124").ToObject(), ßaa.ToObject(), πTemp003).ToObject()
			πTemp002[7] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?P<foo_123>a)").ToObject(), ßa.ToObject(), πTemp003, ßg1.ToObject(), ßa.ToObject()).ToObject()
			πTemp002[8] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?P<foo_123>a)(?P=foo_123)").ToObject(), ßaa.ToObject(), πTemp003, ßg1.ToObject(), ßa.ToObject()).ToObject()
			πTemp002[9] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("\\1").ToObject(), ßa.ToObject(), πTemp003).ToObject()
			πTemp002[10] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("[\\1]").ToObject(), πg.NewStr("\x01").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("\x01").ToObject()).ToObject()
			πTemp002[11] = πTemp001
			πTemp006 = πF.MakeArgs(1)
			πTemp006[0] = πg.NewInt(0).ToObject()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp006)
			if πTemp003, πE = πg.Add(πF, πTemp005, ß9.ToObject()); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp006 = πF.MakeArgs(1)
			πTemp006[0] = πg.NewInt(0).ToObject()
			if πTemp007, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
				continue
			}
			if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp006)
			if πTemp005, πE = πg.Add(πF, πTemp008, ß9.ToObject()); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("\\09").ToObject(), πTemp003, πTemp004, ßfound.ToObject(), πTemp005).ToObject()
			πTemp002[12] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("\\141").ToObject(), ßa.ToObject(), πTemp003, ßfound.ToObject(), ßa.ToObject()).ToObject()
			πTemp002[13] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a)(b)(c)(d)(e)(f)(g)(h)(i)(j)(k)(l)\\119").ToObject(), ßabcdefghijklk9.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g11").ToObject(), πg.NewStr("abcdefghijklk9-k").ToObject()).ToObject()
			πTemp002[14] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("\\0").ToObject(), πg.NewStr("\x00").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("\x00").ToObject()).ToObject()
			πTemp002[15] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("[\\0a]").ToObject(), πg.NewStr("\x00").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("\x00").ToObject()).ToObject()
			πTemp002[16] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("[a\\0]").ToObject(), πg.NewStr("\x00").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("\x00").ToObject()).ToObject()
			πTemp002[17] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("[^a\\0]").ToObject(), πg.NewStr("\x00").ToObject(), πTemp003).ToObject()
			πTemp002[18] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("\\a[\\b]\\f\\n\\r\\t\\v").ToObject(), πg.NewStr("\x07\x08\x0c\n\r\t\x0b").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("\x07\x08\x0c\n\r\t\x0b").ToObject()).ToObject()
			πTemp002[19] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("[\\a][\\b][\\f][\\n][\\r][\\t][\\v]").ToObject(), πg.NewStr("\x07\x08\x0c\n\r\t\x0b").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("\x07\x08\x0c\n\r\t\x0b").ToObject()).ToObject()
			πTemp002[20] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("\\c\\e\\g\\h\\i\\j\\k\\m\\o\\p\\q\\y\\z").ToObject(), ßceghijkmopqyz.ToObject(), πTemp003, ßfound.ToObject(), ßceghijkmopqyz.ToObject()).ToObject()
			πTemp002[21] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp006 = πF.MakeArgs(1)
			πTemp006[0] = πg.NewInt(255).ToObject()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp006)
			πTemp001 = πg.NewTuple5(πg.NewStr("\\xff").ToObject(), πg.NewStr("\xff").ToObject(), πTemp003, ßfound.ToObject(), πTemp005).ToObject()
			πTemp002[22] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp006 = πF.MakeArgs(1)
			πTemp006[0] = πg.NewInt(255).ToObject()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp006)
			πTemp001 = πg.NewTuple5(πg.NewStr("\\x00ffffffffffffff").ToObject(), πg.NewStr("\xff").ToObject(), πTemp003, ßfound.ToObject(), πTemp005).ToObject()
			πTemp002[23] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp006 = πF.MakeArgs(1)
			πTemp006[0] = πg.NewInt(15).ToObject()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp006)
			πTemp001 = πg.NewTuple5(πg.NewStr("\\x00f").ToObject(), πg.NewStr("\x0f").ToObject(), πTemp003, ßfound.ToObject(), πTemp005).ToObject()
			πTemp002[24] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp006 = πF.MakeArgs(1)
			πTemp006[0] = πg.NewInt(254).ToObject()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp006)
			πTemp001 = πg.NewTuple5(πg.NewStr("\\x00fe").ToObject(), πg.NewStr("\xfe").ToObject(), πTemp003, ßfound.ToObject(), πTemp005).ToObject()
			πTemp002[25] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("^\\w+=(\\\\[\\000-\\277]|[^\\n\\\\])*").ToObject(), πg.NewStr("SRC=eval.c g.c blah blah blah \\\\\n\tapes.c").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("SRC=eval.c g.c blah blah blah \\\\").ToObject()).ToObject()
			πTemp002[26] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a.b").ToObject(), ßacb.ToObject(), πTemp003, ßfound.ToObject(), ßacb.ToObject()).ToObject()
			πTemp002[27] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a.b").ToObject(), πg.NewStr("a\nb").ToObject(), πTemp003).ToObject()
			πTemp002[28] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a.*b").ToObject(), πg.NewStr("acc\nccb").ToObject(), πTemp003).ToObject()
			πTemp002[29] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a.{4,5}b").ToObject(), πg.NewStr("acc\nccb").ToObject(), πTemp003).ToObject()
			πTemp002[30] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a.b").ToObject(), πg.NewStr("a\rb").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("a\rb").ToObject()).ToObject()
			πTemp002[31] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a.b(?s)").ToObject(), πg.NewStr("a\nb").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("a\nb").ToObject()).ToObject()
			πTemp002[32] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a.*(?s)b").ToObject(), πg.NewStr("acc\nccb").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("acc\nccb").ToObject()).ToObject()
			πTemp002[33] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?s)a.{4,5}b").ToObject(), πg.NewStr("acc\nccb").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("acc\nccb").ToObject()).ToObject()
			πTemp002[34] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?s)a.b").ToObject(), πg.NewStr("a\nb").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("a\nb").ToObject()).ToObject()
			πTemp002[35] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr(")").ToObject(), ß.ToObject(), πTemp003).ToObject()
			πTemp002[36] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(ß.ToObject(), ß.ToObject(), πTemp003, ßfound.ToObject(), ß.ToObject()).ToObject()
			πTemp002[37] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(ßabc.ToObject(), ßabc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[38] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(ßabc.ToObject(), ßxbc.ToObject(), πTemp003).ToObject()
			πTemp002[39] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(ßabc.ToObject(), ßaxc.ToObject(), πTemp003).ToObject()
			πTemp002[40] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(ßabc.ToObject(), ßabx.ToObject(), πTemp003).ToObject()
			πTemp002[41] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(ßabc.ToObject(), ßxabcy.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[42] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(ßabc.ToObject(), ßababc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[43] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab*c").ToObject(), ßabc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[44] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab*bc").ToObject(), ßabc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[45] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab*bc").ToObject(), ßabbc.ToObject(), πTemp003, ßfound.ToObject(), ßabbc.ToObject()).ToObject()
			πTemp002[46] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab*bc").ToObject(), ßabbbbc.ToObject(), πTemp003, ßfound.ToObject(), ßabbbbc.ToObject()).ToObject()
			πTemp002[47] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab+bc").ToObject(), ßabbc.ToObject(), πTemp003, ßfound.ToObject(), ßabbc.ToObject()).ToObject()
			πTemp002[48] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("ab+bc").ToObject(), ßabc.ToObject(), πTemp003).ToObject()
			πTemp002[49] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("ab+bc").ToObject(), ßabq.ToObject(), πTemp003).ToObject()
			πTemp002[50] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab+bc").ToObject(), ßabbbbc.ToObject(), πTemp003, ßfound.ToObject(), ßabbbbc.ToObject()).ToObject()
			πTemp002[51] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab?bc").ToObject(), ßabbc.ToObject(), πTemp003, ßfound.ToObject(), ßabbc.ToObject()).ToObject()
			πTemp002[52] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab?bc").ToObject(), ßabc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[53] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("ab?bc").ToObject(), ßabbbbc.ToObject(), πTemp003).ToObject()
			πTemp002[54] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab?c").ToObject(), ßabc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[55] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("^abc$").ToObject(), ßabc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[56] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("^abc$").ToObject(), ßabcc.ToObject(), πTemp003).ToObject()
			πTemp002[57] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("^abc").ToObject(), ßabcc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[58] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("^abc$").ToObject(), ßaabc.ToObject(), πTemp003).ToObject()
			πTemp002[59] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("abc$").ToObject(), ßaabc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[60] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("^").ToObject(), ßabc.ToObject(), πTemp003, πg.NewStr("found+\"-\"").ToObject(), πg.NewStr("-").ToObject()).ToObject()
			πTemp002[61] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("$").ToObject(), ßabc.ToObject(), πTemp003, πg.NewStr("found+\"-\"").ToObject(), πg.NewStr("-").ToObject()).ToObject()
			πTemp002[62] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a.c").ToObject(), ßabc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[63] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a.c").ToObject(), ßaxc.ToObject(), πTemp003, ßfound.ToObject(), ßaxc.ToObject()).ToObject()
			πTemp002[64] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a.*c").ToObject(), ßaxyzc.ToObject(), πTemp003, ßfound.ToObject(), ßaxyzc.ToObject()).ToObject()
			πTemp002[65] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a.*c").ToObject(), ßaxyzd.ToObject(), πTemp003).ToObject()
			πTemp002[66] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a[bc]d").ToObject(), ßabc.ToObject(), πTemp003).ToObject()
			πTemp002[67] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a[bc]d").ToObject(), ßabd.ToObject(), πTemp003, ßfound.ToObject(), ßabd.ToObject()).ToObject()
			πTemp002[68] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a[b-d]e").ToObject(), ßabd.ToObject(), πTemp003).ToObject()
			πTemp002[69] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a[b-d]e").ToObject(), ßace.ToObject(), πTemp003, ßfound.ToObject(), ßace.ToObject()).ToObject()
			πTemp002[70] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a[b-d]").ToObject(), ßaac.ToObject(), πTemp003, ßfound.ToObject(), ßac.ToObject()).ToObject()
			πTemp002[71] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a[-b]").ToObject(), πg.NewStr("a-").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("a-").ToObject()).ToObject()
			πTemp002[72] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a[\\-b]").ToObject(), πg.NewStr("a-").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("a-").ToObject()).ToObject()
			πTemp002[73] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a[]b").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[74] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a[").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[75] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a\\").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[76] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("abc)").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[77] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(abc").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[78] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a]").ToObject(), πg.NewStr("a]").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("a]").ToObject()).ToObject()
			πTemp002[79] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a[]]b").ToObject(), πg.NewStr("a]b").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("a]b").ToObject()).ToObject()
			πTemp002[80] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a[\\]]b").ToObject(), πg.NewStr("a]b").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("a]b").ToObject()).ToObject()
			πTemp002[81] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a[^bc]d").ToObject(), ßaed.ToObject(), πTemp003, ßfound.ToObject(), ßaed.ToObject()).ToObject()
			πTemp002[82] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a[^bc]d").ToObject(), ßabd.ToObject(), πTemp003).ToObject()
			πTemp002[83] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a[^-b]c").ToObject(), ßadc.ToObject(), πTemp003, ßfound.ToObject(), ßadc.ToObject()).ToObject()
			πTemp002[84] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a[^-b]c").ToObject(), πg.NewStr("a-c").ToObject(), πTemp003).ToObject()
			πTemp002[85] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a[^]b]c").ToObject(), πg.NewStr("a]c").ToObject(), πTemp003).ToObject()
			πTemp002[86] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a[^]b]c").ToObject(), ßadc.ToObject(), πTemp003, ßfound.ToObject(), ßadc.ToObject()).ToObject()
			πTemp002[87] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("\\ba\\b").ToObject(), πg.NewStr("a-").ToObject(), πTemp003, πg.NewStr("\"-\"").ToObject(), πg.NewStr("-").ToObject()).ToObject()
			πTemp002[88] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("\\ba\\b").ToObject(), πg.NewStr("-a").ToObject(), πTemp003, πg.NewStr("\"-\"").ToObject(), πg.NewStr("-").ToObject()).ToObject()
			πTemp002[89] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("\\ba\\b").ToObject(), πg.NewStr("-a-").ToObject(), πTemp003, πg.NewStr("\"-\"").ToObject(), πg.NewStr("-").ToObject()).ToObject()
			πTemp002[90] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("\\by\\b").ToObject(), ßxy.ToObject(), πTemp003).ToObject()
			πTemp002[91] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("\\by\\b").ToObject(), ßyz.ToObject(), πTemp003).ToObject()
			πTemp002[92] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("\\by\\b").ToObject(), ßxyz.ToObject(), πTemp003).ToObject()
			πTemp002[93] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("x\\b").ToObject(), ßxyz.ToObject(), πTemp003).ToObject()
			πTemp002[94] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("x\\B").ToObject(), ßxyz.ToObject(), πTemp003, πg.NewStr("\"-\"").ToObject(), πg.NewStr("-").ToObject()).ToObject()
			πTemp002[95] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("\\Bz").ToObject(), ßxyz.ToObject(), πTemp003, πg.NewStr("\"-\"").ToObject(), πg.NewStr("-").ToObject()).ToObject()
			πTemp002[96] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("z\\B").ToObject(), ßxyz.ToObject(), πTemp003).ToObject()
			πTemp002[97] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("\\Bx").ToObject(), ßxyz.ToObject(), πTemp003).ToObject()
			πTemp002[98] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("\\Ba\\B").ToObject(), πg.NewStr("a-").ToObject(), πTemp003, πg.NewStr("\"-\"").ToObject(), πg.NewStr("-").ToObject()).ToObject()
			πTemp002[99] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("\\Ba\\B").ToObject(), πg.NewStr("-a").ToObject(), πTemp003, πg.NewStr("\"-\"").ToObject(), πg.NewStr("-").ToObject()).ToObject()
			πTemp002[100] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("\\Ba\\B").ToObject(), πg.NewStr("-a-").ToObject(), πTemp003, πg.NewStr("\"-\"").ToObject(), πg.NewStr("-").ToObject()).ToObject()
			πTemp002[101] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("\\By\\B").ToObject(), ßxy.ToObject(), πTemp003).ToObject()
			πTemp002[102] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("\\By\\B").ToObject(), ßyz.ToObject(), πTemp003).ToObject()
			πTemp002[103] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("\\By\\b").ToObject(), ßxy.ToObject(), πTemp003, πg.NewStr("\"-\"").ToObject(), πg.NewStr("-").ToObject()).ToObject()
			πTemp002[104] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("\\by\\B").ToObject(), ßyz.ToObject(), πTemp003, πg.NewStr("\"-\"").ToObject(), πg.NewStr("-").ToObject()).ToObject()
			πTemp002[105] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("\\By\\B").ToObject(), ßxyz.ToObject(), πTemp003, πg.NewStr("\"-\"").ToObject(), πg.NewStr("-").ToObject()).ToObject()
			πTemp002[106] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab|cd").ToObject(), ßabc.ToObject(), πTemp003, ßfound.ToObject(), ßab.ToObject()).ToObject()
			πTemp002[107] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab|cd").ToObject(), ßabcd.ToObject(), πTemp003, ßfound.ToObject(), ßab.ToObject()).ToObject()
			πTemp002[108] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("()ef").ToObject(), ßdef.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("ef-").ToObject()).ToObject()
			πTemp002[109] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("$b").ToObject(), ßb.ToObject(), πTemp003).ToObject()
			πTemp002[110] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a\\(b").ToObject(), πg.NewStr("a(b").ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("a(b-Error").ToObject()).ToObject()
			πTemp002[111] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a\\(*b").ToObject(), ßab.ToObject(), πTemp003, ßfound.ToObject(), ßab.ToObject()).ToObject()
			πTemp002[112] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a\\(*b").ToObject(), πg.NewStr("a((b").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("a((b").ToObject()).ToObject()
			πTemp002[113] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a\\\\b").ToObject(), πg.NewStr("a\\b").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("a\\b").ToObject()).ToObject()
			πTemp002[114] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("((a))").ToObject(), ßabc.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("a-a-a").ToObject()).ToObject()
			πTemp002[115] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a)b(c)").ToObject(), ßabc.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("abc-a-c").ToObject()).ToObject()
			πTemp002[116] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a+b+c").ToObject(), ßaabbabc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[117] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a+|b)*").ToObject(), ßab.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("ab-b").ToObject()).ToObject()
			πTemp002[118] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a+|b)+").ToObject(), ßab.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("ab-b").ToObject()).ToObject()
			πTemp002[119] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a+|b)?").ToObject(), ßab.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("a-a").ToObject()).ToObject()
			πTemp002[120] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr(")(").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[121] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("[^ab]*").ToObject(), ßcde.ToObject(), πTemp003, ßfound.ToObject(), ßcde.ToObject()).ToObject()
			πTemp002[122] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(ßabc.ToObject(), ß.ToObject(), πTemp003).ToObject()
			πTemp002[123] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a*").ToObject(), ß.ToObject(), πTemp003, ßfound.ToObject(), ß.ToObject()).ToObject()
			πTemp002[124] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a|b|c|d|e").ToObject(), ße.ToObject(), πTemp003, ßfound.ToObject(), ße.ToObject()).ToObject()
			πTemp002[125] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a|b|c|d|e)f").ToObject(), ßef.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("ef-e").ToObject()).ToObject()
			πTemp002[126] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("abcd*efg").ToObject(), ßabcdefg.ToObject(), πTemp003, ßfound.ToObject(), ßabcdefg.ToObject()).ToObject()
			πTemp002[127] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab*").ToObject(), ßxabyabbbz.ToObject(), πTemp003, ßfound.ToObject(), ßab.ToObject()).ToObject()
			πTemp002[128] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab*").ToObject(), ßxayabbbz.ToObject(), πTemp003, ßfound.ToObject(), ßa.ToObject()).ToObject()
			πTemp002[129] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(ab|cd)e").ToObject(), ßabcde.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("cde-cd").ToObject()).ToObject()
			πTemp002[130] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("[abhgefdc]ij").ToObject(), ßhij.ToObject(), πTemp003, ßfound.ToObject(), ßhij.ToObject()).ToObject()
			πTemp002[131] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("^(ab|cd)e").ToObject(), ßabcde.ToObject(), πTemp003, ßxg1y.ToObject(), ßxy.ToObject()).ToObject()
			πTemp002[132] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(abc|)ef").ToObject(), ßabcdef.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("ef-").ToObject()).ToObject()
			πTemp002[133] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a|b)c*d").ToObject(), ßabcd.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("bcd-b").ToObject()).ToObject()
			πTemp002[134] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(ab|ab*)bc").ToObject(), ßabc.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("abc-a").ToObject()).ToObject()
			πTemp002[135] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a([bc]*)c*").ToObject(), ßabc.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("abc-bc").ToObject()).ToObject()
			πTemp002[136] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a([bc]*)(c*d)").ToObject(), ßabcd.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("abcd-bc-d").ToObject()).ToObject()
			πTemp002[137] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a([bc]+)(c*d)").ToObject(), ßabcd.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("abcd-bc-d").ToObject()).ToObject()
			πTemp002[138] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a([bc]*)(c+d)").ToObject(), ßabcd.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("abcd-b-cd").ToObject()).ToObject()
			πTemp002[139] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a[bcd]*dcdcde").ToObject(), ßadcdcde.ToObject(), πTemp003, ßfound.ToObject(), ßadcdcde.ToObject()).ToObject()
			πTemp002[140] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a[bcd]+dcdcde").ToObject(), ßadcdcde.ToObject(), πTemp003).ToObject()
			πTemp002[141] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(ab|a)b*c").ToObject(), ßabc.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("abc-ab").ToObject()).ToObject()
			πTemp002[142] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("((a)(b)c)(d)").ToObject(), ßabcd.ToObject(), πTemp003, πg.NewStr("g1+\"-\"+g2+\"-\"+g3+\"-\"+g4").ToObject(), πg.NewStr("abc-a-b-d").ToObject()).ToObject()
			πTemp002[143] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("[a-zA-Z_][a-zA-Z0-9_]*").ToObject(), ßalpha.ToObject(), πTemp003, ßfound.ToObject(), ßalpha.ToObject()).ToObject()
			πTemp002[144] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("^a(bc+|b[eh])g|.h$").ToObject(), ßabh.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("bh-None").ToObject()).ToObject()
			πTemp002[145] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(bc+d$|ef*g.|h?i(j|k))").ToObject(), ßeffgz.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("effgz-effgz-None").ToObject()).ToObject()
			πTemp002[146] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(bc+d$|ef*g.|h?i(j|k))").ToObject(), ßij.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("ij-ij-j").ToObject()).ToObject()
			πTemp002[147] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(bc+d$|ef*g.|h?i(j|k))").ToObject(), ßeffg.ToObject(), πTemp003).ToObject()
			πTemp002[148] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(bc+d$|ef*g.|h?i(j|k))").ToObject(), ßbcdd.ToObject(), πTemp003).ToObject()
			πTemp002[149] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(bc+d$|ef*g.|h?i(j|k))").ToObject(), ßreffgz.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("effgz-effgz-None").ToObject()).ToObject()
			πTemp002[150] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(((((((((a)))))))))").ToObject(), ßa.ToObject(), πTemp003, ßfound.ToObject(), ßa.ToObject()).ToObject()
			πTemp002[151] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("multiple words of text").ToObject(), πg.NewStr("uh-uh").ToObject(), πTemp003).ToObject()
			πTemp002[152] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("multiple words").ToObject(), πg.NewStr("multiple words, yeah").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("multiple words").ToObject()).ToObject()
			πTemp002[153] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(.*)c(.*)").ToObject(), ßabcde.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("abcde-ab-de").ToObject()).ToObject()
			πTemp002[154] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("\\((.*), (.*)\\)").ToObject(), πg.NewStr("(a, b)").ToObject(), πTemp003, πg.NewStr("g2+\"-\"+g1").ToObject(), πg.NewStr("b-a").ToObject()).ToObject()
			πTemp002[155] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("[k]").ToObject(), ßab.ToObject(), πTemp003).ToObject()
			πTemp002[156] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a[-]?c").ToObject(), ßac.ToObject(), πTemp003, ßfound.ToObject(), ßac.ToObject()).ToObject()
			πTemp002[157] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(abc)\\1").ToObject(), ßabcabc.ToObject(), πTemp003, ßg1.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[158] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("([a-c]*)\\1").ToObject(), ßabcabc.ToObject(), πTemp003, ßg1.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[159] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("^(.+)?B").ToObject(), ßAB.ToObject(), πTemp003, ßg1.ToObject(), ßA.ToObject()).ToObject()
			πTemp002[160] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a+).\\1$").ToObject(), ßaaaaa.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("aaaaa-aa").ToObject()).ToObject()
			πTemp002[161] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("^(a+).\\1$").ToObject(), ßaaaa.ToObject(), πTemp003).ToObject()
			πTemp002[162] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(abc)\\1").ToObject(), ßabcabc.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("abcabc-abc").ToObject()).ToObject()
			πTemp002[163] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("([a-c]+)\\1").ToObject(), ßabcabc.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("abcabc-abc").ToObject()).ToObject()
			πTemp002[164] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a)\\1").ToObject(), ßaa.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("aa-a").ToObject()).ToObject()
			πTemp002[165] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a+)\\1").ToObject(), ßaa.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("aa-a").ToObject()).ToObject()
			πTemp002[166] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a+)+\\1").ToObject(), ßaa.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("aa-a").ToObject()).ToObject()
			πTemp002[167] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a).+\\1").ToObject(), ßaba.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("aba-a").ToObject()).ToObject()
			πTemp002[168] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a)ba*\\1").ToObject(), ßaba.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("aba-a").ToObject()).ToObject()
			πTemp002[169] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(aa|a)a\\1$").ToObject(), ßaaa.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("aaa-a").ToObject()).ToObject()
			πTemp002[170] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a|aa)a\\1$").ToObject(), ßaaa.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("aaa-a").ToObject()).ToObject()
			πTemp002[171] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a+)a\\1$").ToObject(), ßaaa.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("aaa-a").ToObject()).ToObject()
			πTemp002[172] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("([abc]*)\\1").ToObject(), ßabcabc.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("abcabc-abc").ToObject()).ToObject()
			πTemp002[173] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a)(b)c|ab").ToObject(), ßab.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("ab-None-None").ToObject()).ToObject()
			πTemp002[174] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a)+x").ToObject(), ßaaax.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("aaax-a").ToObject()).ToObject()
			πTemp002[175] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("([ac])+x").ToObject(), ßaacx.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("aacx-c").ToObject()).ToObject()
			πTemp002[176] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("([^/]*/)*sub1/").ToObject(), πg.NewStr("d:msgs/tdir/sub1/trial/away.cpp").ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("d:msgs/tdir/sub1/-tdir/").ToObject()).ToObject()
			πTemp002[177] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("([^.]*)\\.([^:]*):[T ]+(.*)").ToObject(), πg.NewStr("track1.title:TBlah blah blah").ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2+\"-\"+g3").ToObject(), πg.NewStr("track1.title:TBlah blah blah-track1-title-Blah blah blah").ToObject()).ToObject()
			πTemp002[178] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("([^N]*N)+").ToObject(), ßabNNxyzN.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("abNNxyzN-xyzN").ToObject()).ToObject()
			πTemp002[179] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("([^N]*N)+").ToObject(), ßabNNxyz.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("abNN-N").ToObject()).ToObject()
			πTemp002[180] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("([abc]*)x").ToObject(), ßabcx.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("abcx-abc").ToObject()).ToObject()
			πTemp002[181] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("([abc]*)x").ToObject(), ßabc.ToObject(), πTemp003).ToObject()
			πTemp002[182] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("([xyz]*)x").ToObject(), ßabcx.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("x-").ToObject()).ToObject()
			πTemp002[183] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a)+b|aac").ToObject(), ßaac.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("aac-None").ToObject()).ToObject()
			πTemp002[184] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?P<i d>aaa)a").ToObject(), ßaaaa.ToObject(), πTemp003).ToObject()
			πTemp002[185] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?P<id>aaa)a").ToObject(), ßaaaa.ToObject(), πTemp003, πg.NewStr("found+\"-\"+id").ToObject(), πg.NewStr("aaaa-aaa").ToObject()).ToObject()
			πTemp002[186] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?P<id>aa)(?P=id)").ToObject(), ßaaaa.ToObject(), πTemp003, πg.NewStr("found+\"-\"+id").ToObject(), πg.NewStr("aaaa-aa").ToObject()).ToObject()
			πTemp002[187] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?P<id>aa)(?P=xd)").ToObject(), ßaaaa.ToObject(), πTemp003).ToObject()
			πTemp002[188] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("\\1").ToObject(), ßa.ToObject(), πTemp003).ToObject()
			πTemp002[189] = πTemp001
			πTemp006 = πF.MakeArgs(1)
			πTemp006[0] = πg.NewInt(0).ToObject()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp006)
			if πTemp003, πE = πg.Add(πF, πTemp005, ß9.ToObject()); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp006 = πF.MakeArgs(1)
			πTemp006[0] = πg.NewInt(0).ToObject()
			if πTemp007, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
				continue
			}
			if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp006)
			if πTemp005, πE = πg.Add(πF, πTemp008, ß9.ToObject()); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("\\09").ToObject(), πTemp003, πTemp004, ßfound.ToObject(), πTemp005).ToObject()
			πTemp002[190] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("\\141").ToObject(), ßa.ToObject(), πTemp003, ßfound.ToObject(), ßa.ToObject()).ToObject()
			πTemp002[191] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a)(b)(c)(d)(e)(f)(g)(h)(i)(j)(k)(l)\\119").ToObject(), ßabcdefghijklk9.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g11").ToObject(), πg.NewStr("abcdefghijklk9-k").ToObject()).ToObject()
			πTemp002[192] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(ßabc.ToObject(), ßabc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[193] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(ßabc.ToObject(), ßxbc.ToObject(), πTemp003).ToObject()
			πTemp002[194] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(ßabc.ToObject(), ßaxc.ToObject(), πTemp003).ToObject()
			πTemp002[195] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(ßabc.ToObject(), ßabx.ToObject(), πTemp003).ToObject()
			πTemp002[196] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(ßabc.ToObject(), ßxabcy.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[197] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(ßabc.ToObject(), ßababc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[198] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab*c").ToObject(), ßabc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[199] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab*bc").ToObject(), ßabc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[200] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab*bc").ToObject(), ßabbc.ToObject(), πTemp003, ßfound.ToObject(), ßabbc.ToObject()).ToObject()
			πTemp002[201] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab*bc").ToObject(), ßabbbbc.ToObject(), πTemp003, ßfound.ToObject(), ßabbbbc.ToObject()).ToObject()
			πTemp002[202] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab{0,}bc").ToObject(), ßabbbbc.ToObject(), πTemp003, ßfound.ToObject(), ßabbbbc.ToObject()).ToObject()
			πTemp002[203] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab+bc").ToObject(), ßabbc.ToObject(), πTemp003, ßfound.ToObject(), ßabbc.ToObject()).ToObject()
			πTemp002[204] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("ab+bc").ToObject(), ßabc.ToObject(), πTemp003).ToObject()
			πTemp002[205] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("ab+bc").ToObject(), ßabq.ToObject(), πTemp003).ToObject()
			πTemp002[206] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("ab{1,}bc").ToObject(), ßabq.ToObject(), πTemp003).ToObject()
			πTemp002[207] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab+bc").ToObject(), ßabbbbc.ToObject(), πTemp003, ßfound.ToObject(), ßabbbbc.ToObject()).ToObject()
			πTemp002[208] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab{1,}bc").ToObject(), ßabbbbc.ToObject(), πTemp003, ßfound.ToObject(), ßabbbbc.ToObject()).ToObject()
			πTemp002[209] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab{1,3}bc").ToObject(), ßabbbbc.ToObject(), πTemp003, ßfound.ToObject(), ßabbbbc.ToObject()).ToObject()
			πTemp002[210] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab{3,4}bc").ToObject(), ßabbbbc.ToObject(), πTemp003, ßfound.ToObject(), ßabbbbc.ToObject()).ToObject()
			πTemp002[211] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("ab{4,5}bc").ToObject(), ßabbbbc.ToObject(), πTemp003).ToObject()
			πTemp002[212] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab?bc").ToObject(), ßabbc.ToObject(), πTemp003, ßfound.ToObject(), ßabbc.ToObject()).ToObject()
			πTemp002[213] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab?bc").ToObject(), ßabc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[214] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab{0,1}bc").ToObject(), ßabc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[215] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("ab?bc").ToObject(), ßabbbbc.ToObject(), πTemp003).ToObject()
			πTemp002[216] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab?c").ToObject(), ßabc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[217] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab{0,1}c").ToObject(), ßabc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[218] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("^abc$").ToObject(), ßabc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[219] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("^abc$").ToObject(), ßabcc.ToObject(), πTemp003).ToObject()
			πTemp002[220] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("^abc").ToObject(), ßabcc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[221] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("^abc$").ToObject(), ßaabc.ToObject(), πTemp003).ToObject()
			πTemp002[222] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("abc$").ToObject(), ßaabc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[223] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("^").ToObject(), ßabc.ToObject(), πTemp003, ßfound.ToObject(), ß.ToObject()).ToObject()
			πTemp002[224] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("$").ToObject(), ßabc.ToObject(), πTemp003, ßfound.ToObject(), ß.ToObject()).ToObject()
			πTemp002[225] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a.c").ToObject(), ßabc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[226] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a.c").ToObject(), ßaxc.ToObject(), πTemp003, ßfound.ToObject(), ßaxc.ToObject()).ToObject()
			πTemp002[227] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a.*c").ToObject(), ßaxyzc.ToObject(), πTemp003, ßfound.ToObject(), ßaxyzc.ToObject()).ToObject()
			πTemp002[228] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a.*c").ToObject(), ßaxyzd.ToObject(), πTemp003).ToObject()
			πTemp002[229] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a[bc]d").ToObject(), ßabc.ToObject(), πTemp003).ToObject()
			πTemp002[230] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a[bc]d").ToObject(), ßabd.ToObject(), πTemp003, ßfound.ToObject(), ßabd.ToObject()).ToObject()
			πTemp002[231] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a[b-d]e").ToObject(), ßabd.ToObject(), πTemp003).ToObject()
			πTemp002[232] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a[b-d]e").ToObject(), ßace.ToObject(), πTemp003, ßfound.ToObject(), ßace.ToObject()).ToObject()
			πTemp002[233] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a[b-d]").ToObject(), ßaac.ToObject(), πTemp003, ßfound.ToObject(), ßac.ToObject()).ToObject()
			πTemp002[234] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a[-b]").ToObject(), πg.NewStr("a-").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("a-").ToObject()).ToObject()
			πTemp002[235] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a[b-]").ToObject(), πg.NewStr("a-").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("a-").ToObject()).ToObject()
			πTemp002[236] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a[b-a]").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[237] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a[]b").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[238] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a[").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[239] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a]").ToObject(), πg.NewStr("a]").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("a]").ToObject()).ToObject()
			πTemp002[240] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a[]]b").ToObject(), πg.NewStr("a]b").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("a]b").ToObject()).ToObject()
			πTemp002[241] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a[^bc]d").ToObject(), ßaed.ToObject(), πTemp003, ßfound.ToObject(), ßaed.ToObject()).ToObject()
			πTemp002[242] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a[^bc]d").ToObject(), ßabd.ToObject(), πTemp003).ToObject()
			πTemp002[243] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a[^-b]c").ToObject(), ßadc.ToObject(), πTemp003, ßfound.ToObject(), ßadc.ToObject()).ToObject()
			πTemp002[244] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a[^-b]c").ToObject(), πg.NewStr("a-c").ToObject(), πTemp003).ToObject()
			πTemp002[245] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a[^]b]c").ToObject(), πg.NewStr("a]c").ToObject(), πTemp003).ToObject()
			πTemp002[246] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a[^]b]c").ToObject(), ßadc.ToObject(), πTemp003, ßfound.ToObject(), ßadc.ToObject()).ToObject()
			πTemp002[247] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab|cd").ToObject(), ßabc.ToObject(), πTemp003, ßfound.ToObject(), ßab.ToObject()).ToObject()
			πTemp002[248] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab|cd").ToObject(), ßabcd.ToObject(), πTemp003, ßfound.ToObject(), ßab.ToObject()).ToObject()
			πTemp002[249] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("()ef").ToObject(), ßdef.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("ef-").ToObject()).ToObject()
			πTemp002[250] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("*a").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[251] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(*)b").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[252] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("$b").ToObject(), ßb.ToObject(), πTemp003).ToObject()
			πTemp002[253] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a\\").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[254] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a\\(b").ToObject(), πg.NewStr("a(b").ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("a(b-Error").ToObject()).ToObject()
			πTemp002[255] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a\\(*b").ToObject(), ßab.ToObject(), πTemp003, ßfound.ToObject(), ßab.ToObject()).ToObject()
			πTemp002[256] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a\\(*b").ToObject(), πg.NewStr("a((b").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("a((b").ToObject()).ToObject()
			πTemp002[257] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a\\\\b").ToObject(), πg.NewStr("a\\b").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("a\\b").ToObject()).ToObject()
			πTemp002[258] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("abc)").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[259] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(abc").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[260] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("((a))").ToObject(), ßabc.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("a-a-a").ToObject()).ToObject()
			πTemp002[261] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a)b(c)").ToObject(), ßabc.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("abc-a-c").ToObject()).ToObject()
			πTemp002[262] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a+b+c").ToObject(), ßaabbabc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[263] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a{1,}b{1,}c").ToObject(), ßaabbabc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[264] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a**").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[265] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a.+?c").ToObject(), ßabcabc.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[266] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a+|b)*").ToObject(), ßab.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("ab-b").ToObject()).ToObject()
			πTemp002[267] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a+|b){0,}").ToObject(), ßab.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("ab-b").ToObject()).ToObject()
			πTemp002[268] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a+|b)+").ToObject(), ßab.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("ab-b").ToObject()).ToObject()
			πTemp002[269] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a+|b){1,}").ToObject(), ßab.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("ab-b").ToObject()).ToObject()
			πTemp002[270] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a+|b)?").ToObject(), ßab.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("a-a").ToObject()).ToObject()
			πTemp002[271] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a+|b){0,1}").ToObject(), ßab.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("a-a").ToObject()).ToObject()
			πTemp002[272] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr(")(").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[273] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("[^ab]*").ToObject(), ßcde.ToObject(), πTemp003, ßfound.ToObject(), ßcde.ToObject()).ToObject()
			πTemp002[274] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(ßabc.ToObject(), ß.ToObject(), πTemp003).ToObject()
			πTemp002[275] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a*").ToObject(), ß.ToObject(), πTemp003, ßfound.ToObject(), ß.ToObject()).ToObject()
			πTemp002[276] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("([abc])*d").ToObject(), ßabbbcd.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("abbbcd-c").ToObject()).ToObject()
			πTemp002[277] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("([abc])*bcd").ToObject(), ßabcd.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("abcd-a").ToObject()).ToObject()
			πTemp002[278] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a|b|c|d|e").ToObject(), ße.ToObject(), πTemp003, ßfound.ToObject(), ße.ToObject()).ToObject()
			πTemp002[279] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a|b|c|d|e)f").ToObject(), ßef.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("ef-e").ToObject()).ToObject()
			πTemp002[280] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("abcd*efg").ToObject(), ßabcdefg.ToObject(), πTemp003, ßfound.ToObject(), ßabcdefg.ToObject()).ToObject()
			πTemp002[281] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab*").ToObject(), ßxabyabbbz.ToObject(), πTemp003, ßfound.ToObject(), ßab.ToObject()).ToObject()
			πTemp002[282] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("ab*").ToObject(), ßxayabbbz.ToObject(), πTemp003, ßfound.ToObject(), ßa.ToObject()).ToObject()
			πTemp002[283] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(ab|cd)e").ToObject(), ßabcde.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("cde-cd").ToObject()).ToObject()
			πTemp002[284] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("[abhgefdc]ij").ToObject(), ßhij.ToObject(), πTemp003, ßfound.ToObject(), ßhij.ToObject()).ToObject()
			πTemp002[285] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("^(ab|cd)e").ToObject(), ßabcde.ToObject(), πTemp003).ToObject()
			πTemp002[286] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(abc|)ef").ToObject(), ßabcdef.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("ef-").ToObject()).ToObject()
			πTemp002[287] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(a|b)c*d").ToObject(), ßabcd.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("bcd-b").ToObject()).ToObject()
			πTemp002[288] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(ab|ab*)bc").ToObject(), ßabc.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("abc-a").ToObject()).ToObject()
			πTemp002[289] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a([bc]*)c*").ToObject(), ßabc.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("abc-bc").ToObject()).ToObject()
			πTemp002[290] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a([bc]*)(c*d)").ToObject(), ßabcd.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("abcd-bc-d").ToObject()).ToObject()
			πTemp002[291] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a([bc]+)(c*d)").ToObject(), ßabcd.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("abcd-bc-d").ToObject()).ToObject()
			πTemp002[292] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a([bc]*)(c+d)").ToObject(), ßabcd.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("abcd-b-cd").ToObject()).ToObject()
			πTemp002[293] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a[bcd]*dcdcde").ToObject(), ßadcdcde.ToObject(), πTemp003, ßfound.ToObject(), ßadcdcde.ToObject()).ToObject()
			πTemp002[294] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a[bcd]+dcdcde").ToObject(), ßadcdcde.ToObject(), πTemp003).ToObject()
			πTemp002[295] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(ab|a)b*c").ToObject(), ßabc.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("abc-ab").ToObject()).ToObject()
			πTemp002[296] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("((a)(b)c)(d)").ToObject(), ßabcd.ToObject(), πTemp003, πg.NewStr("g1+\"-\"+g2+\"-\"+g3+\"-\"+g4").ToObject(), πg.NewStr("abc-a-b-d").ToObject()).ToObject()
			πTemp002[297] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("[a-zA-Z_][a-zA-Z0-9_]*").ToObject(), ßalpha.ToObject(), πTemp003, ßfound.ToObject(), ßalpha.ToObject()).ToObject()
			πTemp002[298] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("^a(bc+|b[eh])g|.h$").ToObject(), ßabh.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("bh-None").ToObject()).ToObject()
			πTemp002[299] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(bc+d$|ef*g.|h?i(j|k))").ToObject(), ßeffgz.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("effgz-effgz-None").ToObject()).ToObject()
			πTemp002[300] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(bc+d$|ef*g.|h?i(j|k))").ToObject(), ßij.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("ij-ij-j").ToObject()).ToObject()
			πTemp002[301] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(bc+d$|ef*g.|h?i(j|k))").ToObject(), ßeffg.ToObject(), πTemp003).ToObject()
			πTemp002[302] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(bc+d$|ef*g.|h?i(j|k))").ToObject(), ßbcdd.ToObject(), πTemp003).ToObject()
			πTemp002[303] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(bc+d$|ef*g.|h?i(j|k))").ToObject(), ßreffgz.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("effgz-effgz-None").ToObject()).ToObject()
			πTemp002[304] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("((((((((((a))))))))))").ToObject(), ßa.ToObject(), πTemp003, ßg10.ToObject(), ßa.ToObject()).ToObject()
			πTemp002[305] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("((((((((((a))))))))))\\10").ToObject(), ßaa.ToObject(), πTemp003, ßfound.ToObject(), ßaa.ToObject()).ToObject()
			πTemp002[306] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("((((((((((a))))))))))\\41").ToObject(), ß.ToObject(), πTemp003).ToObject()
			πTemp002[307] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)((((((((((a))))))))))\\41").ToObject(), ß.ToObject(), πTemp003).ToObject()
			πTemp002[308] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(((((((((a)))))))))").ToObject(), ßa.ToObject(), πTemp003, ßfound.ToObject(), ßa.ToObject()).ToObject()
			πTemp002[309] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("multiple words of text").ToObject(), πg.NewStr("uh-uh").ToObject(), πTemp003).ToObject()
			πTemp002[310] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("multiple words").ToObject(), πg.NewStr("multiple words, yeah").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("multiple words").ToObject()).ToObject()
			πTemp002[311] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(.*)c(.*)").ToObject(), ßabcde.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("abcde-ab-de").ToObject()).ToObject()
			πTemp002[312] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("\\((.*), (.*)\\)").ToObject(), πg.NewStr("(a, b)").ToObject(), πTemp003, πg.NewStr("g2+\"-\"+g1").ToObject(), πg.NewStr("b-a").ToObject()).ToObject()
			πTemp002[313] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("[k]").ToObject(), ßab.ToObject(), πTemp003).ToObject()
			πTemp002[314] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a[-]?c").ToObject(), ßac.ToObject(), πTemp003, ßfound.ToObject(), ßac.ToObject()).ToObject()
			πTemp002[315] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(abc)\\1").ToObject(), ßabcabc.ToObject(), πTemp003, ßg1.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[316] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("([a-c]*)\\1").ToObject(), ßabcabc.ToObject(), πTemp003, ßg1.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[317] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)abc").ToObject(), ßABC.ToObject(), πTemp003, ßfound.ToObject(), ßABC.ToObject()).ToObject()
			πTemp002[318] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)abc").ToObject(), ßXBC.ToObject(), πTemp003).ToObject()
			πTemp002[319] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)abc").ToObject(), ßAXC.ToObject(), πTemp003).ToObject()
			πTemp002[320] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)abc").ToObject(), ßABX.ToObject(), πTemp003).ToObject()
			πTemp002[321] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)abc").ToObject(), ßXABCY.ToObject(), πTemp003, ßfound.ToObject(), ßABC.ToObject()).ToObject()
			πTemp002[322] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)abc").ToObject(), ßABABC.ToObject(), πTemp003, ßfound.ToObject(), ßABC.ToObject()).ToObject()
			πTemp002[323] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)ab*c").ToObject(), ßABC.ToObject(), πTemp003, ßfound.ToObject(), ßABC.ToObject()).ToObject()
			πTemp002[324] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)ab*bc").ToObject(), ßABC.ToObject(), πTemp003, ßfound.ToObject(), ßABC.ToObject()).ToObject()
			πTemp002[325] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)ab*bc").ToObject(), ßABBC.ToObject(), πTemp003, ßfound.ToObject(), ßABBC.ToObject()).ToObject()
			πTemp002[326] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)ab*?bc").ToObject(), ßABBBBC.ToObject(), πTemp003, ßfound.ToObject(), ßABBBBC.ToObject()).ToObject()
			πTemp002[327] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)ab{0,}?bc").ToObject(), ßABBBBC.ToObject(), πTemp003, ßfound.ToObject(), ßABBBBC.ToObject()).ToObject()
			πTemp002[328] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)ab+?bc").ToObject(), ßABBC.ToObject(), πTemp003, ßfound.ToObject(), ßABBC.ToObject()).ToObject()
			πTemp002[329] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)ab+bc").ToObject(), ßABC.ToObject(), πTemp003).ToObject()
			πTemp002[330] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)ab+bc").ToObject(), ßABQ.ToObject(), πTemp003).ToObject()
			πTemp002[331] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)ab{1,}bc").ToObject(), ßABQ.ToObject(), πTemp003).ToObject()
			πTemp002[332] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)ab+bc").ToObject(), ßABBBBC.ToObject(), πTemp003, ßfound.ToObject(), ßABBBBC.ToObject()).ToObject()
			πTemp002[333] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)ab{1,}?bc").ToObject(), ßABBBBC.ToObject(), πTemp003, ßfound.ToObject(), ßABBBBC.ToObject()).ToObject()
			πTemp002[334] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)ab{1,3}?bc").ToObject(), ßABBBBC.ToObject(), πTemp003, ßfound.ToObject(), ßABBBBC.ToObject()).ToObject()
			πTemp002[335] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)ab{3,4}?bc").ToObject(), ßABBBBC.ToObject(), πTemp003, ßfound.ToObject(), ßABBBBC.ToObject()).ToObject()
			πTemp002[336] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)ab{4,5}?bc").ToObject(), ßABBBBC.ToObject(), πTemp003).ToObject()
			πTemp002[337] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)ab??bc").ToObject(), ßABBC.ToObject(), πTemp003, ßfound.ToObject(), ßABBC.ToObject()).ToObject()
			πTemp002[338] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)ab??bc").ToObject(), ßABC.ToObject(), πTemp003, ßfound.ToObject(), ßABC.ToObject()).ToObject()
			πTemp002[339] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)ab{0,1}?bc").ToObject(), ßABC.ToObject(), πTemp003, ßfound.ToObject(), ßABC.ToObject()).ToObject()
			πTemp002[340] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)ab??bc").ToObject(), ßABBBBC.ToObject(), πTemp003).ToObject()
			πTemp002[341] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)ab??c").ToObject(), ßABC.ToObject(), πTemp003, ßfound.ToObject(), ßABC.ToObject()).ToObject()
			πTemp002[342] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)ab{0,1}?c").ToObject(), ßABC.ToObject(), πTemp003, ßfound.ToObject(), ßABC.ToObject()).ToObject()
			πTemp002[343] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)^abc$").ToObject(), ßABC.ToObject(), πTemp003, ßfound.ToObject(), ßABC.ToObject()).ToObject()
			πTemp002[344] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)^abc$").ToObject(), ßABCC.ToObject(), πTemp003).ToObject()
			πTemp002[345] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)^abc").ToObject(), ßABCC.ToObject(), πTemp003, ßfound.ToObject(), ßABC.ToObject()).ToObject()
			πTemp002[346] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)^abc$").ToObject(), ßAABC.ToObject(), πTemp003).ToObject()
			πTemp002[347] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)abc$").ToObject(), ßAABC.ToObject(), πTemp003, ßfound.ToObject(), ßABC.ToObject()).ToObject()
			πTemp002[348] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)^").ToObject(), ßABC.ToObject(), πTemp003, ßfound.ToObject(), ß.ToObject()).ToObject()
			πTemp002[349] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)$").ToObject(), ßABC.ToObject(), πTemp003, ßfound.ToObject(), ß.ToObject()).ToObject()
			πTemp002[350] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a.c").ToObject(), ßABC.ToObject(), πTemp003, ßfound.ToObject(), ßABC.ToObject()).ToObject()
			πTemp002[351] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a.c").ToObject(), ßAXC.ToObject(), πTemp003, ßfound.ToObject(), ßAXC.ToObject()).ToObject()
			πTemp002[352] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a.*?c").ToObject(), ßAXYZC.ToObject(), πTemp003, ßfound.ToObject(), ßAXYZC.ToObject()).ToObject()
			πTemp002[353] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)a.*c").ToObject(), ßAXYZD.ToObject(), πTemp003).ToObject()
			πTemp002[354] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)a[bc]d").ToObject(), ßABC.ToObject(), πTemp003).ToObject()
			πTemp002[355] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a[bc]d").ToObject(), ßABD.ToObject(), πTemp003, ßfound.ToObject(), ßABD.ToObject()).ToObject()
			πTemp002[356] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)a[b-d]e").ToObject(), ßABD.ToObject(), πTemp003).ToObject()
			πTemp002[357] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a[b-d]e").ToObject(), ßACE.ToObject(), πTemp003, ßfound.ToObject(), ßACE.ToObject()).ToObject()
			πTemp002[358] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a[b-d]").ToObject(), ßAAC.ToObject(), πTemp003, ßfound.ToObject(), ßAC.ToObject()).ToObject()
			πTemp002[359] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a[-b]").ToObject(), πg.NewStr("A-").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("A-").ToObject()).ToObject()
			πTemp002[360] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a[b-]").ToObject(), πg.NewStr("A-").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("A-").ToObject()).ToObject()
			πTemp002[361] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)a[b-a]").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[362] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)a[]b").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[363] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)a[").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[364] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a]").ToObject(), πg.NewStr("A]").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("A]").ToObject()).ToObject()
			πTemp002[365] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a[]]b").ToObject(), πg.NewStr("A]B").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("A]B").ToObject()).ToObject()
			πTemp002[366] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a[^bc]d").ToObject(), ßAED.ToObject(), πTemp003, ßfound.ToObject(), ßAED.ToObject()).ToObject()
			πTemp002[367] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)a[^bc]d").ToObject(), ßABD.ToObject(), πTemp003).ToObject()
			πTemp002[368] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a[^-b]c").ToObject(), ßADC.ToObject(), πTemp003, ßfound.ToObject(), ßADC.ToObject()).ToObject()
			πTemp002[369] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)a[^-b]c").ToObject(), πg.NewStr("A-C").ToObject(), πTemp003).ToObject()
			πTemp002[370] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)a[^]b]c").ToObject(), πg.NewStr("A]C").ToObject(), πTemp003).ToObject()
			πTemp002[371] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a[^]b]c").ToObject(), ßADC.ToObject(), πTemp003, ßfound.ToObject(), ßADC.ToObject()).ToObject()
			πTemp002[372] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)ab|cd").ToObject(), ßABC.ToObject(), πTemp003, ßfound.ToObject(), ßAB.ToObject()).ToObject()
			πTemp002[373] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)ab|cd").ToObject(), ßABCD.ToObject(), πTemp003, ßfound.ToObject(), ßAB.ToObject()).ToObject()
			πTemp002[374] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)()ef").ToObject(), ßDEF.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("EF-").ToObject()).ToObject()
			πTemp002[375] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)*a").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[376] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)(*)b").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[377] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)$b").ToObject(), ßB.ToObject(), πTemp003).ToObject()
			πTemp002[378] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)a\\").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[379] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a\\(b").ToObject(), πg.NewStr("A(B").ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("A(B-Error").ToObject()).ToObject()
			πTemp002[380] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a\\(*b").ToObject(), ßAB.ToObject(), πTemp003, ßfound.ToObject(), ßAB.ToObject()).ToObject()
			πTemp002[381] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a\\(*b").ToObject(), πg.NewStr("A((B").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("A((B").ToObject()).ToObject()
			πTemp002[382] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a\\\\b").ToObject(), πg.NewStr("A\\B").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("A\\B").ToObject()).ToObject()
			πTemp002[383] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)abc)").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[384] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)(abc").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[385] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)((a))").ToObject(), ßABC.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("A-A-A").ToObject()).ToObject()
			πTemp002[386] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)(a)b(c)").ToObject(), ßABC.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("ABC-A-C").ToObject()).ToObject()
			πTemp002[387] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a+b+c").ToObject(), ßAABBABC.ToObject(), πTemp003, ßfound.ToObject(), ßABC.ToObject()).ToObject()
			πTemp002[388] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a{1,}b{1,}c").ToObject(), ßAABBABC.ToObject(), πTemp003, ßfound.ToObject(), ßABC.ToObject()).ToObject()
			πTemp002[389] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)a**").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[390] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a.+?c").ToObject(), ßABCABC.ToObject(), πTemp003, ßfound.ToObject(), ßABC.ToObject()).ToObject()
			πTemp002[391] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a.*?c").ToObject(), ßABCABC.ToObject(), πTemp003, ßfound.ToObject(), ßABC.ToObject()).ToObject()
			πTemp002[392] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a.{0,5}?c").ToObject(), ßABCABC.ToObject(), πTemp003, ßfound.ToObject(), ßABC.ToObject()).ToObject()
			πTemp002[393] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)(a+|b)*").ToObject(), ßAB.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("AB-B").ToObject()).ToObject()
			πTemp002[394] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)(a+|b){0,}").ToObject(), ßAB.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("AB-B").ToObject()).ToObject()
			πTemp002[395] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)(a+|b)+").ToObject(), ßAB.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("AB-B").ToObject()).ToObject()
			πTemp002[396] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)(a+|b){1,}").ToObject(), ßAB.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("AB-B").ToObject()).ToObject()
			πTemp002[397] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)(a+|b)?").ToObject(), ßAB.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("A-A").ToObject()).ToObject()
			πTemp002[398] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)(a+|b){0,1}").ToObject(), ßAB.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("A-A").ToObject()).ToObject()
			πTemp002[399] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)(a+|b){0,1}?").ToObject(), ßAB.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("-None").ToObject()).ToObject()
			πTemp002[400] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i))(").ToObject(), πg.NewStr("-").ToObject(), πTemp003).ToObject()
			πTemp002[401] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)[^ab]*").ToObject(), ßCDE.ToObject(), πTemp003, ßfound.ToObject(), ßCDE.ToObject()).ToObject()
			πTemp002[402] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)abc").ToObject(), ß.ToObject(), πTemp003).ToObject()
			πTemp002[403] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a*").ToObject(), ß.ToObject(), πTemp003, ßfound.ToObject(), ß.ToObject()).ToObject()
			πTemp002[404] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)([abc])*d").ToObject(), ßABBBCD.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("ABBBCD-C").ToObject()).ToObject()
			πTemp002[405] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)([abc])*bcd").ToObject(), ßABCD.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("ABCD-A").ToObject()).ToObject()
			πTemp002[406] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a|b|c|d|e").ToObject(), ßE.ToObject(), πTemp003, ßfound.ToObject(), ßE.ToObject()).ToObject()
			πTemp002[407] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)(a|b|c|d|e)f").ToObject(), ßEF.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("EF-E").ToObject()).ToObject()
			πTemp002[408] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)abcd*efg").ToObject(), ßABCDEFG.ToObject(), πTemp003, ßfound.ToObject(), ßABCDEFG.ToObject()).ToObject()
			πTemp002[409] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)ab*").ToObject(), ßXABYABBBZ.ToObject(), πTemp003, ßfound.ToObject(), ßAB.ToObject()).ToObject()
			πTemp002[410] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)ab*").ToObject(), ßXAYABBBZ.ToObject(), πTemp003, ßfound.ToObject(), ßA.ToObject()).ToObject()
			πTemp002[411] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)(ab|cd)e").ToObject(), ßABCDE.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("CDE-CD").ToObject()).ToObject()
			πTemp002[412] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)[abhgefdc]ij").ToObject(), ßHIJ.ToObject(), πTemp003, ßfound.ToObject(), ßHIJ.ToObject()).ToObject()
			πTemp002[413] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)^(ab|cd)e").ToObject(), ßABCDE.ToObject(), πTemp003).ToObject()
			πTemp002[414] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)(abc|)ef").ToObject(), ßABCDEF.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("EF-").ToObject()).ToObject()
			πTemp002[415] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)(a|b)c*d").ToObject(), ßABCD.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("BCD-B").ToObject()).ToObject()
			πTemp002[416] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)(ab|ab*)bc").ToObject(), ßABC.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("ABC-A").ToObject()).ToObject()
			πTemp002[417] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a([bc]*)c*").ToObject(), ßABC.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("ABC-BC").ToObject()).ToObject()
			πTemp002[418] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a([bc]*)(c*d)").ToObject(), ßABCD.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("ABCD-BC-D").ToObject()).ToObject()
			πTemp002[419] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a([bc]+)(c*d)").ToObject(), ßABCD.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("ABCD-BC-D").ToObject()).ToObject()
			πTemp002[420] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a([bc]*)(c+d)").ToObject(), ßABCD.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("ABCD-B-CD").ToObject()).ToObject()
			πTemp002[421] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a[bcd]*dcdcde").ToObject(), ßADCDCDE.ToObject(), πTemp003, ßfound.ToObject(), ßADCDCDE.ToObject()).ToObject()
			πTemp002[422] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)a[bcd]+dcdcde").ToObject(), ßADCDCDE.ToObject(), πTemp003).ToObject()
			πTemp002[423] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)(ab|a)b*c").ToObject(), ßABC.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("ABC-AB").ToObject()).ToObject()
			πTemp002[424] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)((a)(b)c)(d)").ToObject(), ßABCD.ToObject(), πTemp003, πg.NewStr("g1+\"-\"+g2+\"-\"+g3+\"-\"+g4").ToObject(), πg.NewStr("ABC-A-B-D").ToObject()).ToObject()
			πTemp002[425] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)[a-zA-Z_][a-zA-Z0-9_]*").ToObject(), ßALPHA.ToObject(), πTemp003, ßfound.ToObject(), ßALPHA.ToObject()).ToObject()
			πTemp002[426] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)^a(bc+|b[eh])g|.h$").ToObject(), ßABH.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1").ToObject(), πg.NewStr("BH-None").ToObject()).ToObject()
			πTemp002[427] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)(bc+d$|ef*g.|h?i(j|k))").ToObject(), ßEFFGZ.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("EFFGZ-EFFGZ-None").ToObject()).ToObject()
			πTemp002[428] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)(bc+d$|ef*g.|h?i(j|k))").ToObject(), ßIJ.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("IJ-IJ-J").ToObject()).ToObject()
			πTemp002[429] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)(bc+d$|ef*g.|h?i(j|k))").ToObject(), ßEFFG.ToObject(), πTemp003).ToObject()
			πTemp002[430] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)(bc+d$|ef*g.|h?i(j|k))").ToObject(), ßBCDD.ToObject(), πTemp003).ToObject()
			πTemp002[431] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)(bc+d$|ef*g.|h?i(j|k))").ToObject(), ßREFFGZ.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("EFFGZ-EFFGZ-None").ToObject()).ToObject()
			πTemp002[432] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)((((((((((a))))))))))").ToObject(), ßA.ToObject(), πTemp003, ßg10.ToObject(), ßA.ToObject()).ToObject()
			πTemp002[433] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)((((((((((a))))))))))\\10").ToObject(), ßAA.ToObject(), πTemp003, ßfound.ToObject(), ßAA.ToObject()).ToObject()
			πTemp002[434] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)(((((((((a)))))))))").ToObject(), ßA.ToObject(), πTemp003, ßfound.ToObject(), ßA.ToObject()).ToObject()
			πTemp002[435] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)(?:(?:(?:(?:(?:(?:(?:(?:(?:(a))))))))))").ToObject(), ßA.ToObject(), πTemp003, ßg1.ToObject(), ßA.ToObject()).ToObject()
			πTemp002[436] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)(?:(?:(?:(?:(?:(?:(?:(?:(?:(a|b|c))))))))))").ToObject(), ßC.ToObject(), πTemp003, ßg1.ToObject(), ßC.ToObject()).ToObject()
			πTemp002[437] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)multiple words of text").ToObject(), πg.NewStr("UH-UH").ToObject(), πTemp003).ToObject()
			πTemp002[438] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)multiple words").ToObject(), πg.NewStr("MULTIPLE WORDS, YEAH").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("MULTIPLE WORDS").ToObject()).ToObject()
			πTemp002[439] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)(.*)c(.*)").ToObject(), ßABCDE.ToObject(), πTemp003, πg.NewStr("found+\"-\"+g1+\"-\"+g2").ToObject(), πg.NewStr("ABCDE-AB-DE").ToObject()).ToObject()
			πTemp002[440] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)\\((.*), (.*)\\)").ToObject(), πg.NewStr("(A, B)").ToObject(), πTemp003, πg.NewStr("g2+\"-\"+g1").ToObject(), πg.NewStr("B-A").ToObject()).ToObject()
			πTemp002[441] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(?i)[k]").ToObject(), ßAB.ToObject(), πTemp003).ToObject()
			πTemp002[442] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)a[-]?c").ToObject(), ßAC.ToObject(), πTemp003, ßfound.ToObject(), ßAC.ToObject()).ToObject()
			πTemp002[443] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)(abc)\\1").ToObject(), ßABCABC.ToObject(), πTemp003, ßg1.ToObject(), ßABC.ToObject()).ToObject()
			πTemp002[444] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)([a-c]*)\\1").ToObject(), ßABCABC.ToObject(), πTemp003, ßg1.ToObject(), ßABC.ToObject()).ToObject()
			πTemp002[445] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a(?!b).").ToObject(), ßabad.ToObject(), πTemp003, ßfound.ToObject(), ßad.ToObject()).ToObject()
			πTemp002[446] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a(?=d).").ToObject(), ßabad.ToObject(), πTemp003, ßfound.ToObject(), ßad.ToObject()).ToObject()
			πTemp002[447] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a(?=c|d).").ToObject(), ßabad.ToObject(), πTemp003, ßfound.ToObject(), ßad.ToObject()).ToObject()
			πTemp002[448] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a(?:b|c|d)(.)").ToObject(), ßace.ToObject(), πTemp003, ßg1.ToObject(), ße.ToObject()).ToObject()
			πTemp002[449] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a(?:b|c|d)*(.)").ToObject(), ßace.ToObject(), πTemp003, ßg1.ToObject(), ße.ToObject()).ToObject()
			πTemp002[450] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a(?:b|c|d)+?(.)").ToObject(), ßace.ToObject(), πTemp003, ßg1.ToObject(), ße.ToObject()).ToObject()
			πTemp002[451] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a(?:b|(c|e){1,2}?|d)+?(.)").ToObject(), ßace.ToObject(), πTemp003, πg.NewStr("g1 + g2").ToObject(), ßce.ToObject()).ToObject()
			πTemp002[452] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("^(.+)?B").ToObject(), ßAB.ToObject(), πTemp003, ßg1.ToObject(), ßA.ToObject()).ToObject()
			πTemp002[453] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?<!-):(.*?)(?<!-):").ToObject(), πg.NewStr("a:bc-:de:f").ToObject(), πTemp003, ßg1.ToObject(), πg.NewStr("bc-:de").ToObject()).ToObject()
			πTemp002[454] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?<!\\\\):(.*?)(?<!\\\\):").ToObject(), πg.NewStr("a:bc\\:de:f").ToObject(), πTemp003, ßg1.ToObject(), πg.NewStr("bc\\:de").ToObject()).ToObject()
			πTemp002[455] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?<!\\?)'(.*?)(?<!\\?)'").ToObject(), πg.NewStr("a'bc?'de'f").ToObject(), πTemp003, ßg1.ToObject(), πg.NewStr("bc?'de").ToObject()).ToObject()
			πTemp002[456] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("w(?# comment").ToObject(), ßw.ToObject(), πTemp003).ToObject()
			πTemp002[457] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("w(?# comment 1)xy(?# comment 2)z").ToObject(), ßwxyz.ToObject(), πTemp003, ßfound.ToObject(), ßwxyz.ToObject()).ToObject()
			πTemp002[458] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("w(?i)").ToObject(), ßW.ToObject(), πTemp003, ßfound.ToObject(), ßW.ToObject()).ToObject()
			πTemp002[459] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?x)w# comment 1\n        x y\n        # comment 2\n        z").ToObject(), ßwxyz.ToObject(), πTemp003, ßfound.ToObject(), ßwxyz.ToObject()).ToObject()
			πTemp002[460] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("^abc").ToObject(), πg.NewStr("jkl\nabc\nxyz").ToObject(), πTemp003).ToObject()
			πTemp002[461] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?m)^abc").ToObject(), πg.NewStr("jkl\nabc\nxyz").ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[462] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?m)abc$").ToObject(), πg.NewStr("jkl\nxyzabc\n123").ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[463] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a.b").ToObject(), πg.NewStr("a\nb").ToObject(), πTemp003).ToObject()
			πTemp002[464] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?s)a.b").ToObject(), πg.NewStr("a\nb").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("a\nb").ToObject()).ToObject()
			πTemp002[465] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("\\w+").ToObject(), πg.NewStr("--ab_cd0123--").ToObject(), πTemp003, ßfound.ToObject(), ßab_cd0123.ToObject()).ToObject()
			πTemp002[466] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("[\\w]+").ToObject(), πg.NewStr("--ab_cd0123--").ToObject(), πTemp003, ßfound.ToObject(), ßab_cd0123.ToObject()).ToObject()
			πTemp002[467] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("\\D+").ToObject(), ß1234abc5678.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[468] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("[\\D]+").ToObject(), ß1234abc5678.ToObject(), πTemp003, ßfound.ToObject(), ßabc.ToObject()).ToObject()
			πTemp002[469] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("[\\da-fA-F]+").ToObject(), ß123abc.ToObject(), πTemp003, ßfound.ToObject(), ß123abc.ToObject()).ToObject()
			πTemp002[470] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("([\\s]*)([\\S]*)([\\s]*)").ToObject(), πg.NewStr(" testing!1972").ToObject(), πTemp003, πg.NewStr("g3+g2+g1").ToObject(), πg.NewStr("testing!1972 ").ToObject()).ToObject()
			πTemp002[471] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(\\s*)(\\S*)(\\s*)").ToObject(), πg.NewStr(" testing!1972").ToObject(), πTemp003, πg.NewStr("g3+g2+g1").ToObject(), πg.NewStr("testing!1972 ").ToObject()).ToObject()
			πTemp002[472] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp006 = πF.MakeArgs(1)
			πTemp006[0] = πg.NewInt(255).ToObject()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp006)
			πTemp001 = πg.NewTuple5(πg.NewStr("\\xff").ToObject(), πg.NewStr("\xff").ToObject(), πTemp003, ßfound.ToObject(), πTemp005).ToObject()
			πTemp002[473] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("\\x00ff").ToObject(), πg.NewStr("\xff").ToObject(), πTemp003).ToObject()
			πTemp002[474] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("\\t\\n\\v\\r\\f\\a\\g").ToObject(), πg.NewStr("\t\n\x0b\r\x0c\x07g").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("\t\n\x0b\r\x0c\x07g").ToObject()).ToObject()
			πTemp002[475] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("\t\n\x0b\r\x0c\x07\\g").ToObject(), πg.NewStr("\t\n\x0b\r\x0c\x07g").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("\t\n\x0b\r\x0c\x07g").ToObject()).ToObject()
			πTemp002[476] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp006 = πF.MakeArgs(1)
			πTemp006[0] = πg.NewInt(9).ToObject()
			if πTemp010, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp010.Call(πF, πTemp006, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp006)
			πTemp006 = πF.MakeArgs(1)
			πTemp006[0] = πg.NewInt(10).ToObject()
			if πTemp010, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
				continue
			}
			if πTemp012, πE = πTemp010.Call(πF, πTemp006, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp006)
			if πTemp009, πE = πg.Add(πF, πTemp011, πTemp012); πE != nil {
				continue
			}
			πTemp006 = πF.MakeArgs(1)
			πTemp006[0] = πg.NewInt(11).ToObject()
			if πTemp010, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp010.Call(πF, πTemp006, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp006)
			if πTemp008, πE = πg.Add(πF, πTemp009, πTemp011); πE != nil {
				continue
			}
			πTemp006 = πF.MakeArgs(1)
			πTemp006[0] = πg.NewInt(13).ToObject()
			if πTemp009, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
				continue
			}
			if πTemp010, πE = πTemp009.Call(πF, πTemp006, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp006)
			if πTemp007, πE = πg.Add(πF, πTemp008, πTemp010); πE != nil {
				continue
			}
			πTemp006 = πF.MakeArgs(1)
			πTemp006[0] = πg.NewInt(12).ToObject()
			if πTemp008, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
				continue
			}
			if πTemp009, πE = πTemp008.Call(πF, πTemp006, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp006)
			if πTemp005, πE = πg.Add(πF, πTemp007, πTemp009); πE != nil {
				continue
			}
			πTemp006 = πF.MakeArgs(1)
			πTemp006[0] = πg.NewInt(7).ToObject()
			if πTemp007, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
				continue
			}
			if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp006)
			if πTemp004, πE = πg.Add(πF, πTemp005, πTemp008); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("\\t\\n\\v\\r\\f\\a").ToObject(), πg.NewStr("\t\n\x0b\r\x0c\x07").ToObject(), πTemp003, ßfound.ToObject(), πTemp004).ToObject()
			πTemp002[477] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("[\\t][\\n][\\v][\\r][\\f][\\b]").ToObject(), πg.NewStr("\t\n\x0b\r\x0c\x08").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("\t\n\x0b\r\x0c\x08").ToObject()).ToObject()
			πTemp002[478] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(([a-z]+):)?([a-z]+)$").ToObject(), ßsmil.ToObject(), πTemp003, πg.NewStr("g1+\"-\"+g2+\"-\"+g3").ToObject(), πg.NewStr("None-None-smil").ToObject()).ToObject()
			πTemp002[479] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("((.)\\1+)").ToObject(), ß.ToObject(), πTemp003).ToObject()
			πTemp002[480] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr(".*d").ToObject(), πg.NewStr("abc\nabd").ToObject(), πTemp003, ßfound.ToObject(), ßabd.ToObject()).ToObject()
			πTemp002[481] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("(").ToObject(), ß.ToObject(), πTemp003).ToObject()
			πTemp002[482] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("[\\41]").ToObject(), πg.NewStr("!").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("!").ToObject()).ToObject()
			πTemp002[483] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(x?)?").ToObject(), ßx.ToObject(), πTemp003, ßfound.ToObject(), ßx.ToObject()).ToObject()
			πTemp002[484] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr(" (?x)foo ").ToObject(), ßfoo.ToObject(), πTemp003, ßfound.ToObject(), ßfoo.ToObject()).ToObject()
			πTemp002[485] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?<!abc)(d.f)").ToObject(), ßabcdefdof.ToObject(), πTemp003, ßfound.ToObject(), ßdof.ToObject()).ToObject()
			πTemp002[486] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("[\\w-]+").ToObject(), ßlaser_beam.ToObject(), πTemp003, ßfound.ToObject(), ßlaser_beam.ToObject()).ToObject()
			πTemp002[487] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr(".*?\\S *:").ToObject(), πg.NewStr("xx:").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("xx:").ToObject()).ToObject()
			πTemp002[488] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a[ ]*?\\ (\\d+).*").ToObject(), πg.NewStr("a   10").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("a   10").ToObject()).ToObject()
			πTemp002[489] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("a[ ]*?\\ (\\d+).*").ToObject(), πg.NewStr("a    10").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("a    10").ToObject()).ToObject()
			πTemp002[490] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?ms).*?x\\s*\\Z(.*)").ToObject(), πg.NewStr("xx\nx\n").ToObject(), πTemp003, ßg1.ToObject(), ß.ToObject()).ToObject()
			πTemp002[491] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)M+").ToObject(), ßMMM.ToObject(), πTemp003, ßfound.ToObject(), ßMMM.ToObject()).ToObject()
			πTemp002[492] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)m+").ToObject(), ßMMM.ToObject(), πTemp003, ßfound.ToObject(), ßMMM.ToObject()).ToObject()
			πTemp002[493] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)[M]+").ToObject(), ßMMM.ToObject(), πTemp003, ßfound.ToObject(), ßMMM.ToObject()).ToObject()
			πTemp002[494] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("(?i)[m]+").ToObject(), ßMMM.ToObject(), πTemp003, ßfound.ToObject(), ßMMM.ToObject()).ToObject()
			πTemp002[495] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("^*").ToObject(), ß.ToObject(), πTemp003).ToObject()
			πTemp002[496] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("\"(?:\\\\\"|[^\"])*?\"").ToObject(), πg.NewStr("\"\\\"\"").ToObject(), πTemp003, ßfound.ToObject(), πg.NewStr("\"\\\"\"").ToObject()).ToObject()
			πTemp002[497] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("^.*?$").ToObject(), πg.NewStr("one\ntwo\nthree\n").ToObject(), πTemp003).ToObject()
			πTemp002[498] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("a[^>]*?b").ToObject(), πg.NewStr("a>b").ToObject(), πTemp003).ToObject()
			πTemp002[499] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple3(πg.NewStr("^a*?$").ToObject(), ßfoo.ToObject(), πTemp003).ToObject()
			πTemp002[500] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("^((a)c)?(ab)$").ToObject(), ßab.ToObject(), πTemp003, πg.NewStr("g1+\"-\"+g2+\"-\"+g3").ToObject(), πg.NewStr("None-None-ab").ToObject()).ToObject()
			πTemp002[501] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("^([ab]*?)(?=(b)?)c").ToObject(), ßabc.ToObject(), πTemp003, πg.NewStr("g1+\"-\"+g2").ToObject(), πg.NewStr("ab-None").ToObject()).ToObject()
			πTemp002[502] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("^([ab]*?)(?!(b))c").ToObject(), ßabc.ToObject(), πTemp003, πg.NewStr("g1+\"-\"+g2").ToObject(), πg.NewStr("ab-None").ToObject()).ToObject()
			πTemp002[503] = πTemp001
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πg.NewStr("^([ab]*?)(?<!(a))c").ToObject(), ßabc.ToObject(), πTemp003, πg.NewStr("g1+\"-\"+g2").ToObject(), πg.NewStr("ab-None").ToObject()).ToObject()
			πTemp002[504] = πTemp001
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßtests.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßtests); πE != nil {
				continue
			}
			if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
				continue
			}
			πF.PushCheckpoint(2)
			πTemp013 = false
		Label1:
			if πE != nil || πR != nil {
				continue
			}
			if πTemp013 {
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
				πTemp014 = !isStop
			} else {
				πTemp014 = true
				if πE = πF.Globals().SetItem(πF, ßtest.ToObject(), πTemp003); πE != nil {
					continue
				}
			}
			if πE != nil || !πTemp014 {
				continue
			}
			πF.PushCheckpoint(1)            
			// line 679: pattern, s, status = test[:3]
			πF.SetLineno(679)
			if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(3).ToObject(), πg.None}, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßtest); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp007}}}, πTemp004); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßpattern.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßs.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßstatus.ToObject(), πTemp007); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßstatus); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßSUCCEED); πE != nil {
				continue
			}
			if πTemp003, πE = πg.Eq(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			if πTemp014, πE = πg.IsTrue(πF, πTemp003); πE != nil {
				continue
			}
			if πTemp014 {
				goto Label4
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßstatus); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßFAIL); πE != nil {
				continue
			}
			if πTemp003, πE = πg.Eq(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			if πTemp014, πE = πg.IsTrue(πF, πTemp003); πE != nil {
				continue
			}
			if πTemp014 {
				goto Label5
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßstatus); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßSYNTAX_ERROR); πE != nil {
				continue
			}
			if πTemp003, πE = πg.Eq(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			if πTemp014, πE = πg.IsTrue(πF, πTemp003); πE != nil {
				continue
			}
			if πTemp014 {
				goto Label6
			}
			goto Label7
			// line 680: if status == SUCCEED:
			πF.SetLineno(680)
		Label4:
			// line 681: assert re.search(pattern, s)
			πF.SetLineno(681)
			πTemp002 = πF.MakeArgs(2)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßpattern); πE != nil {
				continue
			}
			πTemp002[0] = πTemp003
			if πTemp003, πE = πg.ResolveGlobal(πF, ßs); πE != nil {
				continue
			}
			πTemp002[1] = πTemp003
			if πTemp003, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsearch, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πg.Assert(πF, πTemp003, nil); πE != nil {
				continue
			}
			goto Label7
			// line 682: elif status == FAIL:
			πF.SetLineno(682)
		Label5:
			// line 683: assert not re.search(pattern, s)
			πF.SetLineno(683)
			πTemp002 = πF.MakeArgs(2)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßpattern); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			if πTemp004, πE = πg.ResolveGlobal(πF, ßs); πE != nil {
				continue
			}
			πTemp002[1] = πTemp004
			if πTemp004, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßsearch, nil); πE != nil {
				continue
			}
			if πTemp004, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πTemp014, πE = πg.IsTrue(πF, πTemp004); πE != nil {
				continue
			}
			πTemp003 = πg.GetBool(!πTemp014).ToObject()
			if πE = πg.Assert(πF, πTemp003, nil); πE != nil {
				continue
			}
			goto Label7
			// line 684: elif status == SYNTAX_ERROR:
			πF.SetLineno(684)
		Label6:
			// line 685: try:
			πF.SetLineno(685)
			πF.PushCheckpoint(9)
			// line 686: re.search(pattern, s)
			πF.SetLineno(686)
			πTemp002 = πF.MakeArgs(2)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßpattern); πE != nil {
				continue
			}
			πTemp002[0] = πTemp003
			if πTemp003, πE = πg.ResolveGlobal(πF, ßs); πE != nil {
				continue
			}
			πTemp002[1] = πTemp003
			if πTemp003, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsearch, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 687: assert AssertionError
			πF.SetLineno(687)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
				continue
			}
			if πE = πg.Assert(πF, πTemp003, nil); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label8
		Label9:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp015, πTemp016 = πF.ExcInfo()
			if πTemp003, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
				continue
			}
			if πTemp014, πE = πg.IsInstance(πF, πTemp015.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πTemp014 {
				goto Label10
			}
			πE = πF.Raise(πTemp015.ToObject(), nil, πTemp016.ToObject())
			continue
			// line 688: except Exception:  # pylint: disable=broad-except
			πF.SetLineno(688)
		Label10:
			// line 689: pass
			πF.SetLineno(689)
			πF.RestoreExc(nil, nil)
			goto Label8
		Label8:
			goto Label7
		Label7:
			continue
		Label2:
			if πE != nil || πR != nil {
				continue
			}
		Label3:
		}
		return nil, πE
	})
	πg.RegisterModule("re_tests", Code)
}
