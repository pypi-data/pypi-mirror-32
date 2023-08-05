package sre_constants
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/sre_constants.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßANY := πg.InternStr("ANY")
		ßANY_ALL := πg.InternStr("ANY_ALL")
		ßASSERT := πg.InternStr("ASSERT")
		ßASSERT_NOT := πg.InternStr("ASSERT_NOT")
		ßAT := πg.InternStr("AT")
		ßATCODES := πg.InternStr("ATCODES")
		ßAT_BEGINNING := πg.InternStr("AT_BEGINNING")
		ßAT_BEGINNING_LINE := πg.InternStr("AT_BEGINNING_LINE")
		ßAT_BEGINNING_STRING := πg.InternStr("AT_BEGINNING_STRING")
		ßAT_BOUNDARY := πg.InternStr("AT_BOUNDARY")
		ßAT_END := πg.InternStr("AT_END")
		ßAT_END_LINE := πg.InternStr("AT_END_LINE")
		ßAT_END_STRING := πg.InternStr("AT_END_STRING")
		ßAT_LOCALE := πg.InternStr("AT_LOCALE")
		ßAT_LOC_BOUNDARY := πg.InternStr("AT_LOC_BOUNDARY")
		ßAT_LOC_NON_BOUNDARY := πg.InternStr("AT_LOC_NON_BOUNDARY")
		ßAT_MULTILINE := πg.InternStr("AT_MULTILINE")
		ßAT_NON_BOUNDARY := πg.InternStr("AT_NON_BOUNDARY")
		ßAT_UNICODE := πg.InternStr("AT_UNICODE")
		ßAT_UNI_BOUNDARY := πg.InternStr("AT_UNI_BOUNDARY")
		ßAT_UNI_NON_BOUNDARY := πg.InternStr("AT_UNI_NON_BOUNDARY")
		ßBIGCHARSET := πg.InternStr("BIGCHARSET")
		ßBRANCH := πg.InternStr("BRANCH")
		ßCALL := πg.InternStr("CALL")
		ßCATEGORY := πg.InternStr("CATEGORY")
		ßCATEGORY_DIGIT := πg.InternStr("CATEGORY_DIGIT")
		ßCATEGORY_LINEBREAK := πg.InternStr("CATEGORY_LINEBREAK")
		ßCATEGORY_LOC_NOT_WORD := πg.InternStr("CATEGORY_LOC_NOT_WORD")
		ßCATEGORY_LOC_WORD := πg.InternStr("CATEGORY_LOC_WORD")
		ßCATEGORY_NOT_DIGIT := πg.InternStr("CATEGORY_NOT_DIGIT")
		ßCATEGORY_NOT_LINEBREAK := πg.InternStr("CATEGORY_NOT_LINEBREAK")
		ßCATEGORY_NOT_SPACE := πg.InternStr("CATEGORY_NOT_SPACE")
		ßCATEGORY_NOT_WORD := πg.InternStr("CATEGORY_NOT_WORD")
		ßCATEGORY_SPACE := πg.InternStr("CATEGORY_SPACE")
		ßCATEGORY_UNI_DIGIT := πg.InternStr("CATEGORY_UNI_DIGIT")
		ßCATEGORY_UNI_LINEBREAK := πg.InternStr("CATEGORY_UNI_LINEBREAK")
		ßCATEGORY_UNI_NOT_DIGIT := πg.InternStr("CATEGORY_UNI_NOT_DIGIT")
		ßCATEGORY_UNI_NOT_LINEBREAK := πg.InternStr("CATEGORY_UNI_NOT_LINEBREAK")
		ßCATEGORY_UNI_NOT_SPACE := πg.InternStr("CATEGORY_UNI_NOT_SPACE")
		ßCATEGORY_UNI_NOT_WORD := πg.InternStr("CATEGORY_UNI_NOT_WORD")
		ßCATEGORY_UNI_SPACE := πg.InternStr("CATEGORY_UNI_SPACE")
		ßCATEGORY_UNI_WORD := πg.InternStr("CATEGORY_UNI_WORD")
		ßCATEGORY_WORD := πg.InternStr("CATEGORY_WORD")
		ßCHARSET := πg.InternStr("CHARSET")
		ßCHCODES := πg.InternStr("CHCODES")
		ßCH_LOCALE := πg.InternStr("CH_LOCALE")
		ßCH_UNICODE := πg.InternStr("CH_UNICODE")
		ßException := πg.InternStr("Exception")
		ßFAILURE := πg.InternStr("FAILURE")
		ßGROUPREF := πg.InternStr("GROUPREF")
		ßGROUPREF_EXISTS := πg.InternStr("GROUPREF_EXISTS")
		ßGROUPREF_IGNORE := πg.InternStr("GROUPREF_IGNORE")
		ßIN := πg.InternStr("IN")
		ßINFO := πg.InternStr("INFO")
		ßIN_IGNORE := πg.InternStr("IN_IGNORE")
		ßJUMP := πg.InternStr("JUMP")
		ßLITERAL := πg.InternStr("LITERAL")
		ßLITERAL_IGNORE := πg.InternStr("LITERAL_IGNORE")
		ßMAGIC := πg.InternStr("MAGIC")
		ßMARK := πg.InternStr("MARK")
		ßMAXCODE := πg.InternStr("MAXCODE")
		ßMAXREPEAT := πg.InternStr("MAXREPEAT")
		ßMAX_REPEAT := πg.InternStr("MAX_REPEAT")
		ßMAX_UNTIL := πg.InternStr("MAX_UNTIL")
		ßMIN_REPEAT := πg.InternStr("MIN_REPEAT")
		ßMIN_REPEAT_ONE := πg.InternStr("MIN_REPEAT_ONE")
		ßMIN_UNTIL := πg.InternStr("MIN_UNTIL")
		ßNEGATE := πg.InternStr("NEGATE")
		ßNOT_LITERAL := πg.InternStr("NOT_LITERAL")
		ßNOT_LITERAL_IGNORE := πg.InternStr("NOT_LITERAL_IGNORE")
		ßOPCODES := πg.InternStr("OPCODES")
		ßOP_IGNORE := πg.InternStr("OP_IGNORE")
		ßRANGE := πg.InternStr("RANGE")
		ßREPEAT := πg.InternStr("REPEAT")
		ßREPEAT_ONE := πg.InternStr("REPEAT_ONE")
		ßSRE_FLAG_DEBUG := πg.InternStr("SRE_FLAG_DEBUG")
		ßSRE_FLAG_DOTALL := πg.InternStr("SRE_FLAG_DOTALL")
		ßSRE_FLAG_IGNORECASE := πg.InternStr("SRE_FLAG_IGNORECASE")
		ßSRE_FLAG_LOCALE := πg.InternStr("SRE_FLAG_LOCALE")
		ßSRE_FLAG_MULTILINE := πg.InternStr("SRE_FLAG_MULTILINE")
		ßSRE_FLAG_TEMPLATE := πg.InternStr("SRE_FLAG_TEMPLATE")
		ßSRE_FLAG_UNICODE := πg.InternStr("SRE_FLAG_UNICODE")
		ßSRE_FLAG_VERBOSE := πg.InternStr("SRE_FLAG_VERBOSE")
		ßSRE_INFO_CHARSET := πg.InternStr("SRE_INFO_CHARSET")
		ßSRE_INFO_LITERAL := πg.InternStr("SRE_INFO_LITERAL")
		ßSRE_INFO_PREFIX := πg.InternStr("SRE_INFO_PREFIX")
		ßSUBPATTERN := πg.InternStr("SUBPATTERN")
		ßSUCCESS := πg.InternStr("SUCCESS")
		ß__all__ := πg.InternStr("__all__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßany := πg.InternStr("any")
		ßany_all := πg.InternStr("any_all")
		ßassert := πg.InternStr("assert")
		ßassert_not := πg.InternStr("assert_not")
		ßat := πg.InternStr("at")
		ßat_beginning := πg.InternStr("at_beginning")
		ßat_beginning_line := πg.InternStr("at_beginning_line")
		ßat_beginning_string := πg.InternStr("at_beginning_string")
		ßat_boundary := πg.InternStr("at_boundary")
		ßat_end := πg.InternStr("at_end")
		ßat_end_line := πg.InternStr("at_end_line")
		ßat_end_string := πg.InternStr("at_end_string")
		ßat_loc_boundary := πg.InternStr("at_loc_boundary")
		ßat_loc_non_boundary := πg.InternStr("at_loc_non_boundary")
		ßat_non_boundary := πg.InternStr("at_non_boundary")
		ßat_uni_boundary := πg.InternStr("at_uni_boundary")
		ßat_uni_non_boundary := πg.InternStr("at_uni_non_boundary")
		ßbigcharset := πg.InternStr("bigcharset")
		ßbranch := πg.InternStr("branch")
		ßcall := πg.InternStr("call")
		ßcategory := πg.InternStr("category")
		ßcategory_digit := πg.InternStr("category_digit")
		ßcategory_linebreak := πg.InternStr("category_linebreak")
		ßcategory_loc_not_word := πg.InternStr("category_loc_not_word")
		ßcategory_loc_word := πg.InternStr("category_loc_word")
		ßcategory_not_digit := πg.InternStr("category_not_digit")
		ßcategory_not_linebreak := πg.InternStr("category_not_linebreak")
		ßcategory_not_space := πg.InternStr("category_not_space")
		ßcategory_not_word := πg.InternStr("category_not_word")
		ßcategory_space := πg.InternStr("category_space")
		ßcategory_uni_digit := πg.InternStr("category_uni_digit")
		ßcategory_uni_linebreak := πg.InternStr("category_uni_linebreak")
		ßcategory_uni_not_digit := πg.InternStr("category_uni_not_digit")
		ßcategory_uni_not_linebreak := πg.InternStr("category_uni_not_linebreak")
		ßcategory_uni_not_space := πg.InternStr("category_uni_not_space")
		ßcategory_uni_not_word := πg.InternStr("category_uni_not_word")
		ßcategory_uni_space := πg.InternStr("category_uni_space")
		ßcategory_uni_word := πg.InternStr("category_uni_word")
		ßcategory_word := πg.InternStr("category_word")
		ßcharset := πg.InternStr("charset")
		ßerror := πg.InternStr("error")
		ßfailure := πg.InternStr("failure")
		ßgroupref := πg.InternStr("groupref")
		ßgroupref_exists := πg.InternStr("groupref_exists")
		ßgroupref_ignore := πg.InternStr("groupref_ignore")
		ßin := πg.InternStr("in")
		ßin_ignore := πg.InternStr("in_ignore")
		ßinfo := πg.InternStr("info")
		ßjump := πg.InternStr("jump")
		ßliteral := πg.InternStr("literal")
		ßliteral_ignore := πg.InternStr("literal_ignore")
		ßmakedict := πg.InternStr("makedict")
		ßmark := πg.InternStr("mark")
		ßmax_repeat := πg.InternStr("max_repeat")
		ßmax_until := πg.InternStr("max_until")
		ßmin_repeat := πg.InternStr("min_repeat")
		ßmin_repeat_one := πg.InternStr("min_repeat_one")
		ßmin_until := πg.InternStr("min_until")
		ßnegate := πg.InternStr("negate")
		ßnot_literal := πg.InternStr("not_literal")
		ßnot_literal_ignore := πg.InternStr("not_literal_ignore")
		ßrange := πg.InternStr("range")
		ßrepeat := πg.InternStr("repeat")
		ßrepeat_one := πg.InternStr("repeat_one")
		ßsubpattern := πg.InternStr("subpattern")
		ßsuccess := πg.InternStr("success")
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
		var πTemp006 []πg.Param
		_ = πTemp006
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 12: """Internal support module for sre"""
			πF.SetLineno(12)
			// line 13: __all__ = [
			πF.SetLineno(13)
			πTemp001 = make([]*πg.Object, 88)
			πTemp001[0] = ßANY.ToObject()
			πTemp001[1] = ßANY_ALL.ToObject()
			πTemp001[2] = ßASSERT.ToObject()
			πTemp001[3] = ßASSERT_NOT.ToObject()
			πTemp001[4] = ßAT.ToObject()
			πTemp001[5] = ßATCODES.ToObject()
			πTemp001[6] = ßAT_BEGINNING.ToObject()
			πTemp001[7] = ßAT_BEGINNING_LINE.ToObject()
			πTemp001[8] = ßAT_BEGINNING_STRING.ToObject()
			πTemp001[9] = ßAT_BOUNDARY.ToObject()
			πTemp001[10] = ßAT_END.ToObject()
			πTemp001[11] = ßAT_END_LINE.ToObject()
			πTemp001[12] = ßAT_END_STRING.ToObject()
			πTemp001[13] = ßAT_LOCALE.ToObject()
			πTemp001[14] = ßAT_LOC_BOUNDARY.ToObject()
			πTemp001[15] = ßAT_LOC_NON_BOUNDARY.ToObject()
			πTemp001[16] = ßAT_MULTILINE.ToObject()
			πTemp001[17] = ßAT_NON_BOUNDARY.ToObject()
			πTemp001[18] = ßAT_UNICODE.ToObject()
			πTemp001[19] = ßAT_UNI_BOUNDARY.ToObject()
			πTemp001[20] = ßAT_UNI_NON_BOUNDARY.ToObject()
			πTemp001[21] = ßBIGCHARSET.ToObject()
			πTemp001[22] = ßBRANCH.ToObject()
			πTemp001[23] = ßCALL.ToObject()
			πTemp001[24] = ßCATEGORY.ToObject()
			πTemp001[25] = ßCATEGORY_DIGIT.ToObject()
			πTemp001[26] = ßCATEGORY_LINEBREAK.ToObject()
			πTemp001[27] = ßCATEGORY_LOC_NOT_WORD.ToObject()
			πTemp001[28] = ßCATEGORY_LOC_WORD.ToObject()
			πTemp001[29] = ßCATEGORY_NOT_DIGIT.ToObject()
			πTemp001[30] = ßCATEGORY_NOT_LINEBREAK.ToObject()
			πTemp001[31] = ßCATEGORY_NOT_SPACE.ToObject()
			πTemp001[32] = ßCATEGORY_NOT_WORD.ToObject()
			πTemp001[33] = ßCATEGORY_SPACE.ToObject()
			πTemp001[34] = ßCATEGORY_UNI_DIGIT.ToObject()
			πTemp001[35] = ßCATEGORY_UNI_LINEBREAK.ToObject()
			πTemp001[36] = ßCATEGORY_UNI_NOT_DIGIT.ToObject()
			πTemp001[37] = ßCATEGORY_UNI_NOT_LINEBREAK.ToObject()
			πTemp001[38] = ßCATEGORY_UNI_NOT_SPACE.ToObject()
			πTemp001[39] = ßCATEGORY_UNI_NOT_WORD.ToObject()
			πTemp001[40] = ßCATEGORY_UNI_SPACE.ToObject()
			πTemp001[41] = ßCATEGORY_UNI_WORD.ToObject()
			πTemp001[42] = ßCATEGORY_WORD.ToObject()
			πTemp001[43] = ßCHARSET.ToObject()
			πTemp001[44] = ßCHCODES.ToObject()
			πTemp001[45] = ßCH_LOCALE.ToObject()
			πTemp001[46] = ßCH_UNICODE.ToObject()
			πTemp001[47] = ßFAILURE.ToObject()
			πTemp001[48] = ßGROUPREF.ToObject()
			πTemp001[49] = ßGROUPREF_EXISTS.ToObject()
			πTemp001[50] = ßGROUPREF_IGNORE.ToObject()
			πTemp001[51] = ßIN.ToObject()
			πTemp001[52] = ßINFO.ToObject()
			πTemp001[53] = ßIN_IGNORE.ToObject()
			πTemp001[54] = ßJUMP.ToObject()
			πTemp001[55] = ßLITERAL.ToObject()
			πTemp001[56] = ßLITERAL_IGNORE.ToObject()
			πTemp001[57] = ßMAGIC.ToObject()
			πTemp001[58] = ßMARK.ToObject()
			πTemp001[59] = ßMAXREPEAT.ToObject()
			πTemp001[60] = ßMAX_REPEAT.ToObject()
			πTemp001[61] = ßMAX_UNTIL.ToObject()
			πTemp001[62] = ßMIN_REPEAT.ToObject()
			πTemp001[63] = ßMIN_REPEAT_ONE.ToObject()
			πTemp001[64] = ßMIN_UNTIL.ToObject()
			πTemp001[65] = ßNEGATE.ToObject()
			πTemp001[66] = ßNOT_LITERAL.ToObject()
			πTemp001[67] = ßNOT_LITERAL_IGNORE.ToObject()
			πTemp001[68] = ßOPCODES.ToObject()
			πTemp001[69] = ßOP_IGNORE.ToObject()
			πTemp001[70] = ßRANGE.ToObject()
			πTemp001[71] = ßREPEAT.ToObject()
			πTemp001[72] = ßREPEAT_ONE.ToObject()
			πTemp001[73] = ßSRE_FLAG_DOTALL.ToObject()
			πTemp001[74] = ßSRE_FLAG_IGNORECASE.ToObject()
			πTemp001[75] = ßSRE_FLAG_LOCALE.ToObject()
			πTemp001[76] = ßSRE_FLAG_MULTILINE.ToObject()
			πTemp001[77] = ßSRE_FLAG_TEMPLATE.ToObject()
			πTemp001[78] = ßSRE_FLAG_UNICODE.ToObject()
			πTemp001[79] = ßSRE_FLAG_VERBOSE.ToObject()
			πTemp001[80] = ßSRE_INFO_CHARSET.ToObject()
			πTemp001[81] = ßSRE_INFO_LITERAL.ToObject()
			πTemp001[82] = ßSRE_INFO_PREFIX.ToObject()
			πTemp001[83] = ßSUBPATTERN.ToObject()
			πTemp001[84] = ßSUCCESS.ToObject()
			πTemp001[85] = ßSRE_FLAG_DEBUG.ToObject()
			πTemp001[86] = ßMAXCODE.ToObject()
			πTemp001[87] = ßerror.ToObject()
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 39: MAGIC = 20031017
			πF.SetLineno(39)
			if πE = πF.Globals().SetItem(πF, ßMAGIC.ToObject(), πg.NewInt(20031017).ToObject()); πE != nil {
				continue
			}
			// line 41: MAXCODE = 65535
			πF.SetLineno(41)
			if πE = πF.Globals().SetItem(πF, ßMAXCODE.ToObject(), πg.NewInt(65535).ToObject()); πE != nil {
				continue
			}
			// line 48: MAXREPEAT = 65535
			πF.SetLineno(48)
			if πE = πF.Globals().SetItem(πF, ßMAXREPEAT.ToObject(), πg.NewInt(65535).ToObject()); πE != nil {
				continue
			}
			// line 53: class error(Exception):
			πF.SetLineno(53)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
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
			_, πE = πg.NewCode("error", "build/src/__python__/sre_constants.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 54: pass
					πF.SetLineno(54)
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("error").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßerror.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 58: FAILURE = "failure"
			πF.SetLineno(58)
			if πE = πF.Globals().SetItem(πF, ßFAILURE.ToObject(), ßfailure.ToObject()); πE != nil {
				continue
			}
			// line 59: SUCCESS = "success"
			πF.SetLineno(59)
			if πE = πF.Globals().SetItem(πF, ßSUCCESS.ToObject(), ßsuccess.ToObject()); πE != nil {
				continue
			}
			// line 61: ANY = "any"
			πF.SetLineno(61)
			if πE = πF.Globals().SetItem(πF, ßANY.ToObject(), ßany.ToObject()); πE != nil {
				continue
			}
			// line 62: ANY_ALL = "any_all"
			πF.SetLineno(62)
			if πE = πF.Globals().SetItem(πF, ßANY_ALL.ToObject(), ßany_all.ToObject()); πE != nil {
				continue
			}
			// line 63: ASSERT = "assert"
			πF.SetLineno(63)
			if πE = πF.Globals().SetItem(πF, ßASSERT.ToObject(), ßassert.ToObject()); πE != nil {
				continue
			}
			// line 64: ASSERT_NOT = "assert_not"
			πF.SetLineno(64)
			if πE = πF.Globals().SetItem(πF, ßASSERT_NOT.ToObject(), ßassert_not.ToObject()); πE != nil {
				continue
			}
			// line 65: AT = "at"
			πF.SetLineno(65)
			if πE = πF.Globals().SetItem(πF, ßAT.ToObject(), ßat.ToObject()); πE != nil {
				continue
			}
			// line 66: BIGCHARSET = "bigcharset"
			πF.SetLineno(66)
			if πE = πF.Globals().SetItem(πF, ßBIGCHARSET.ToObject(), ßbigcharset.ToObject()); πE != nil {
				continue
			}
			// line 67: BRANCH = "branch"
			πF.SetLineno(67)
			if πE = πF.Globals().SetItem(πF, ßBRANCH.ToObject(), ßbranch.ToObject()); πE != nil {
				continue
			}
			// line 68: CALL = "call"
			πF.SetLineno(68)
			if πE = πF.Globals().SetItem(πF, ßCALL.ToObject(), ßcall.ToObject()); πE != nil {
				continue
			}
			// line 69: CATEGORY = "category"
			πF.SetLineno(69)
			if πE = πF.Globals().SetItem(πF, ßCATEGORY.ToObject(), ßcategory.ToObject()); πE != nil {
				continue
			}
			// line 70: CHARSET = "charset"
			πF.SetLineno(70)
			if πE = πF.Globals().SetItem(πF, ßCHARSET.ToObject(), ßcharset.ToObject()); πE != nil {
				continue
			}
			// line 71: GROUPREF = "groupref"
			πF.SetLineno(71)
			if πE = πF.Globals().SetItem(πF, ßGROUPREF.ToObject(), ßgroupref.ToObject()); πE != nil {
				continue
			}
			// line 72: GROUPREF_IGNORE = "groupref_ignore"
			πF.SetLineno(72)
			if πE = πF.Globals().SetItem(πF, ßGROUPREF_IGNORE.ToObject(), ßgroupref_ignore.ToObject()); πE != nil {
				continue
			}
			// line 73: GROUPREF_EXISTS = "groupref_exists"
			πF.SetLineno(73)
			if πE = πF.Globals().SetItem(πF, ßGROUPREF_EXISTS.ToObject(), ßgroupref_exists.ToObject()); πE != nil {
				continue
			}
			// line 74: IN = "in"
			πF.SetLineno(74)
			if πE = πF.Globals().SetItem(πF, ßIN.ToObject(), ßin.ToObject()); πE != nil {
				continue
			}
			// line 75: IN_IGNORE = "in_ignore"
			πF.SetLineno(75)
			if πE = πF.Globals().SetItem(πF, ßIN_IGNORE.ToObject(), ßin_ignore.ToObject()); πE != nil {
				continue
			}
			// line 76: INFO = "info"
			πF.SetLineno(76)
			if πE = πF.Globals().SetItem(πF, ßINFO.ToObject(), ßinfo.ToObject()); πE != nil {
				continue
			}
			// line 77: JUMP = "jump"
			πF.SetLineno(77)
			if πE = πF.Globals().SetItem(πF, ßJUMP.ToObject(), ßjump.ToObject()); πE != nil {
				continue
			}
			// line 78: LITERAL = "literal"
			πF.SetLineno(78)
			if πE = πF.Globals().SetItem(πF, ßLITERAL.ToObject(), ßliteral.ToObject()); πE != nil {
				continue
			}
			// line 79: LITERAL_IGNORE = "literal_ignore"
			πF.SetLineno(79)
			if πE = πF.Globals().SetItem(πF, ßLITERAL_IGNORE.ToObject(), ßliteral_ignore.ToObject()); πE != nil {
				continue
			}
			// line 80: MARK = "mark"
			πF.SetLineno(80)
			if πE = πF.Globals().SetItem(πF, ßMARK.ToObject(), ßmark.ToObject()); πE != nil {
				continue
			}
			// line 81: MAX_REPEAT = "max_repeat"
			πF.SetLineno(81)
			if πE = πF.Globals().SetItem(πF, ßMAX_REPEAT.ToObject(), ßmax_repeat.ToObject()); πE != nil {
				continue
			}
			// line 82: MAX_UNTIL = "max_until"
			πF.SetLineno(82)
			if πE = πF.Globals().SetItem(πF, ßMAX_UNTIL.ToObject(), ßmax_until.ToObject()); πE != nil {
				continue
			}
			// line 83: MIN_REPEAT = "min_repeat"
			πF.SetLineno(83)
			if πE = πF.Globals().SetItem(πF, ßMIN_REPEAT.ToObject(), ßmin_repeat.ToObject()); πE != nil {
				continue
			}
			// line 84: MIN_UNTIL = "min_until"
			πF.SetLineno(84)
			if πE = πF.Globals().SetItem(πF, ßMIN_UNTIL.ToObject(), ßmin_until.ToObject()); πE != nil {
				continue
			}
			// line 85: NEGATE = "negate"
			πF.SetLineno(85)
			if πE = πF.Globals().SetItem(πF, ßNEGATE.ToObject(), ßnegate.ToObject()); πE != nil {
				continue
			}
			// line 86: NOT_LITERAL = "not_literal"
			πF.SetLineno(86)
			if πE = πF.Globals().SetItem(πF, ßNOT_LITERAL.ToObject(), ßnot_literal.ToObject()); πE != nil {
				continue
			}
			// line 87: NOT_LITERAL_IGNORE = "not_literal_ignore"
			πF.SetLineno(87)
			if πE = πF.Globals().SetItem(πF, ßNOT_LITERAL_IGNORE.ToObject(), ßnot_literal_ignore.ToObject()); πE != nil {
				continue
			}
			// line 88: RANGE = "range"
			πF.SetLineno(88)
			if πE = πF.Globals().SetItem(πF, ßRANGE.ToObject(), ßrange.ToObject()); πE != nil {
				continue
			}
			// line 89: REPEAT = "repeat"
			πF.SetLineno(89)
			if πE = πF.Globals().SetItem(πF, ßREPEAT.ToObject(), ßrepeat.ToObject()); πE != nil {
				continue
			}
			// line 90: REPEAT_ONE = "repeat_one"
			πF.SetLineno(90)
			if πE = πF.Globals().SetItem(πF, ßREPEAT_ONE.ToObject(), ßrepeat_one.ToObject()); πE != nil {
				continue
			}
			// line 91: SUBPATTERN = "subpattern"
			πF.SetLineno(91)
			if πE = πF.Globals().SetItem(πF, ßSUBPATTERN.ToObject(), ßsubpattern.ToObject()); πE != nil {
				continue
			}
			// line 92: MIN_REPEAT_ONE = "min_repeat_one"
			πF.SetLineno(92)
			if πE = πF.Globals().SetItem(πF, ßMIN_REPEAT_ONE.ToObject(), ßmin_repeat_one.ToObject()); πE != nil {
				continue
			}
			// line 95: AT_BEGINNING = "at_beginning"
			πF.SetLineno(95)
			if πE = πF.Globals().SetItem(πF, ßAT_BEGINNING.ToObject(), ßat_beginning.ToObject()); πE != nil {
				continue
			}
			// line 96: AT_BEGINNING_LINE = "at_beginning_line"
			πF.SetLineno(96)
			if πE = πF.Globals().SetItem(πF, ßAT_BEGINNING_LINE.ToObject(), ßat_beginning_line.ToObject()); πE != nil {
				continue
			}
			// line 97: AT_BEGINNING_STRING = "at_beginning_string"
			πF.SetLineno(97)
			if πE = πF.Globals().SetItem(πF, ßAT_BEGINNING_STRING.ToObject(), ßat_beginning_string.ToObject()); πE != nil {
				continue
			}
			// line 98: AT_BOUNDARY = "at_boundary"
			πF.SetLineno(98)
			if πE = πF.Globals().SetItem(πF, ßAT_BOUNDARY.ToObject(), ßat_boundary.ToObject()); πE != nil {
				continue
			}
			// line 99: AT_NON_BOUNDARY = "at_non_boundary"
			πF.SetLineno(99)
			if πE = πF.Globals().SetItem(πF, ßAT_NON_BOUNDARY.ToObject(), ßat_non_boundary.ToObject()); πE != nil {
				continue
			}
			// line 100: AT_END = "at_end"
			πF.SetLineno(100)
			if πE = πF.Globals().SetItem(πF, ßAT_END.ToObject(), ßat_end.ToObject()); πE != nil {
				continue
			}
			// line 101: AT_END_LINE = "at_end_line"
			πF.SetLineno(101)
			if πE = πF.Globals().SetItem(πF, ßAT_END_LINE.ToObject(), ßat_end_line.ToObject()); πE != nil {
				continue
			}
			// line 102: AT_END_STRING = "at_end_string"
			πF.SetLineno(102)
			if πE = πF.Globals().SetItem(πF, ßAT_END_STRING.ToObject(), ßat_end_string.ToObject()); πE != nil {
				continue
			}
			// line 103: AT_LOC_BOUNDARY = "at_loc_boundary"
			πF.SetLineno(103)
			if πE = πF.Globals().SetItem(πF, ßAT_LOC_BOUNDARY.ToObject(), ßat_loc_boundary.ToObject()); πE != nil {
				continue
			}
			// line 104: AT_LOC_NON_BOUNDARY = "at_loc_non_boundary"
			πF.SetLineno(104)
			if πE = πF.Globals().SetItem(πF, ßAT_LOC_NON_BOUNDARY.ToObject(), ßat_loc_non_boundary.ToObject()); πE != nil {
				continue
			}
			// line 105: AT_UNI_BOUNDARY = "at_uni_boundary"
			πF.SetLineno(105)
			if πE = πF.Globals().SetItem(πF, ßAT_UNI_BOUNDARY.ToObject(), ßat_uni_boundary.ToObject()); πE != nil {
				continue
			}
			// line 106: AT_UNI_NON_BOUNDARY = "at_uni_non_boundary"
			πF.SetLineno(106)
			if πE = πF.Globals().SetItem(πF, ßAT_UNI_NON_BOUNDARY.ToObject(), ßat_uni_non_boundary.ToObject()); πE != nil {
				continue
			}
			// line 109: CATEGORY_DIGIT = "category_digit"
			πF.SetLineno(109)
			if πE = πF.Globals().SetItem(πF, ßCATEGORY_DIGIT.ToObject(), ßcategory_digit.ToObject()); πE != nil {
				continue
			}
			// line 110: CATEGORY_NOT_DIGIT = "category_not_digit"
			πF.SetLineno(110)
			if πE = πF.Globals().SetItem(πF, ßCATEGORY_NOT_DIGIT.ToObject(), ßcategory_not_digit.ToObject()); πE != nil {
				continue
			}
			// line 111: CATEGORY_SPACE = "category_space"
			πF.SetLineno(111)
			if πE = πF.Globals().SetItem(πF, ßCATEGORY_SPACE.ToObject(), ßcategory_space.ToObject()); πE != nil {
				continue
			}
			// line 112: CATEGORY_NOT_SPACE = "category_not_space"
			πF.SetLineno(112)
			if πE = πF.Globals().SetItem(πF, ßCATEGORY_NOT_SPACE.ToObject(), ßcategory_not_space.ToObject()); πE != nil {
				continue
			}
			// line 113: CATEGORY_WORD = "category_word"
			πF.SetLineno(113)
			if πE = πF.Globals().SetItem(πF, ßCATEGORY_WORD.ToObject(), ßcategory_word.ToObject()); πE != nil {
				continue
			}
			// line 114: CATEGORY_NOT_WORD = "category_not_word"
			πF.SetLineno(114)
			if πE = πF.Globals().SetItem(πF, ßCATEGORY_NOT_WORD.ToObject(), ßcategory_not_word.ToObject()); πE != nil {
				continue
			}
			// line 115: CATEGORY_LINEBREAK = "category_linebreak"
			πF.SetLineno(115)
			if πE = πF.Globals().SetItem(πF, ßCATEGORY_LINEBREAK.ToObject(), ßcategory_linebreak.ToObject()); πE != nil {
				continue
			}
			// line 116: CATEGORY_NOT_LINEBREAK = "category_not_linebreak"
			πF.SetLineno(116)
			if πE = πF.Globals().SetItem(πF, ßCATEGORY_NOT_LINEBREAK.ToObject(), ßcategory_not_linebreak.ToObject()); πE != nil {
				continue
			}
			// line 117: CATEGORY_LOC_WORD = "category_loc_word"
			πF.SetLineno(117)
			if πE = πF.Globals().SetItem(πF, ßCATEGORY_LOC_WORD.ToObject(), ßcategory_loc_word.ToObject()); πE != nil {
				continue
			}
			// line 118: CATEGORY_LOC_NOT_WORD = "category_loc_not_word"
			πF.SetLineno(118)
			if πE = πF.Globals().SetItem(πF, ßCATEGORY_LOC_NOT_WORD.ToObject(), ßcategory_loc_not_word.ToObject()); πE != nil {
				continue
			}
			// line 119: CATEGORY_UNI_DIGIT = "category_uni_digit"
			πF.SetLineno(119)
			if πE = πF.Globals().SetItem(πF, ßCATEGORY_UNI_DIGIT.ToObject(), ßcategory_uni_digit.ToObject()); πE != nil {
				continue
			}
			// line 120: CATEGORY_UNI_NOT_DIGIT = "category_uni_not_digit"
			πF.SetLineno(120)
			if πE = πF.Globals().SetItem(πF, ßCATEGORY_UNI_NOT_DIGIT.ToObject(), ßcategory_uni_not_digit.ToObject()); πE != nil {
				continue
			}
			// line 121: CATEGORY_UNI_SPACE = "category_uni_space"
			πF.SetLineno(121)
			if πE = πF.Globals().SetItem(πF, ßCATEGORY_UNI_SPACE.ToObject(), ßcategory_uni_space.ToObject()); πE != nil {
				continue
			}
			// line 122: CATEGORY_UNI_NOT_SPACE = "category_uni_not_space"
			πF.SetLineno(122)
			if πE = πF.Globals().SetItem(πF, ßCATEGORY_UNI_NOT_SPACE.ToObject(), ßcategory_uni_not_space.ToObject()); πE != nil {
				continue
			}
			// line 123: CATEGORY_UNI_WORD = "category_uni_word"
			πF.SetLineno(123)
			if πE = πF.Globals().SetItem(πF, ßCATEGORY_UNI_WORD.ToObject(), ßcategory_uni_word.ToObject()); πE != nil {
				continue
			}
			// line 124: CATEGORY_UNI_NOT_WORD = "category_uni_not_word"
			πF.SetLineno(124)
			if πE = πF.Globals().SetItem(πF, ßCATEGORY_UNI_NOT_WORD.ToObject(), ßcategory_uni_not_word.ToObject()); πE != nil {
				continue
			}
			// line 125: CATEGORY_UNI_LINEBREAK = "category_uni_linebreak"
			πF.SetLineno(125)
			if πE = πF.Globals().SetItem(πF, ßCATEGORY_UNI_LINEBREAK.ToObject(), ßcategory_uni_linebreak.ToObject()); πE != nil {
				continue
			}
			// line 126: CATEGORY_UNI_NOT_LINEBREAK = "category_uni_not_linebreak"
			πF.SetLineno(126)
			if πE = πF.Globals().SetItem(πF, ßCATEGORY_UNI_NOT_LINEBREAK.ToObject(), ßcategory_uni_not_linebreak.ToObject()); πE != nil {
				continue
			}
			// line 128: OPCODES = [
			πF.SetLineno(128)
			πTemp001 = make([]*πg.Object, 32)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßFAILURE); πE != nil {
				continue
			}
			πTemp001[0] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßSUCCESS); πE != nil {
				continue
			}
			πTemp001[1] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßANY); πE != nil {
				continue
			}
			πTemp001[2] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßANY_ALL); πE != nil {
				continue
			}
			πTemp001[3] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßASSERT); πE != nil {
				continue
			}
			πTemp001[4] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßASSERT_NOT); πE != nil {
				continue
			}
			πTemp001[5] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßAT); πE != nil {
				continue
			}
			πTemp001[6] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßBRANCH); πE != nil {
				continue
			}
			πTemp001[7] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßCALL); πE != nil {
				continue
			}
			πTemp001[8] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßCATEGORY); πE != nil {
				continue
			}
			πTemp001[9] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßCHARSET); πE != nil {
				continue
			}
			πTemp001[10] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßBIGCHARSET); πE != nil {
				continue
			}
			πTemp001[11] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßGROUPREF); πE != nil {
				continue
			}
			πTemp001[12] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßGROUPREF_EXISTS); πE != nil {
				continue
			}
			πTemp001[13] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßGROUPREF_IGNORE); πE != nil {
				continue
			}
			πTemp001[14] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßIN); πE != nil {
				continue
			}
			πTemp001[15] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßIN_IGNORE); πE != nil {
				continue
			}
			πTemp001[16] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßINFO); πE != nil {
				continue
			}
			πTemp001[17] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßJUMP); πE != nil {
				continue
			}
			πTemp001[18] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
				continue
			}
			πTemp001[19] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßLITERAL_IGNORE); πE != nil {
				continue
			}
			πTemp001[20] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßMARK); πE != nil {
				continue
			}
			πTemp001[21] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßMAX_UNTIL); πE != nil {
				continue
			}
			πTemp001[22] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßMIN_UNTIL); πE != nil {
				continue
			}
			πTemp001[23] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßNOT_LITERAL); πE != nil {
				continue
			}
			πTemp001[24] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßNOT_LITERAL_IGNORE); πE != nil {
				continue
			}
			πTemp001[25] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßNEGATE); πE != nil {
				continue
			}
			πTemp001[26] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßRANGE); πE != nil {
				continue
			}
			πTemp001[27] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßREPEAT); πE != nil {
				continue
			}
			πTemp001[28] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßREPEAT_ONE); πE != nil {
				continue
			}
			πTemp001[29] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßSUBPATTERN); πE != nil {
				continue
			}
			πTemp001[30] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßMIN_REPEAT_ONE); πE != nil {
				continue
			}
			πTemp001[31] = πTemp002
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßOPCODES.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 158: ATCODES = [
			πF.SetLineno(158)
			πTemp001 = make([]*πg.Object, 12)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßAT_BEGINNING); πE != nil {
				continue
			}
			πTemp001[0] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßAT_BEGINNING_LINE); πE != nil {
				continue
			}
			πTemp001[1] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßAT_BEGINNING_STRING); πE != nil {
				continue
			}
			πTemp001[2] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßAT_BOUNDARY); πE != nil {
				continue
			}
			πTemp001[3] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßAT_NON_BOUNDARY); πE != nil {
				continue
			}
			πTemp001[4] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßAT_END); πE != nil {
				continue
			}
			πTemp001[5] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßAT_END_LINE); πE != nil {
				continue
			}
			πTemp001[6] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßAT_END_STRING); πE != nil {
				continue
			}
			πTemp001[7] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßAT_LOC_BOUNDARY); πE != nil {
				continue
			}
			πTemp001[8] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßAT_LOC_NON_BOUNDARY); πE != nil {
				continue
			}
			πTemp001[9] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßAT_UNI_BOUNDARY); πE != nil {
				continue
			}
			πTemp001[10] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßAT_UNI_NON_BOUNDARY); πE != nil {
				continue
			}
			πTemp001[11] = πTemp002
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßATCODES.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 165: CHCODES = [
			πF.SetLineno(165)
			πTemp001 = make([]*πg.Object, 18)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßCATEGORY_DIGIT); πE != nil {
				continue
			}
			πTemp001[0] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßCATEGORY_NOT_DIGIT); πE != nil {
				continue
			}
			πTemp001[1] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßCATEGORY_SPACE); πE != nil {
				continue
			}
			πTemp001[2] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßCATEGORY_NOT_SPACE); πE != nil {
				continue
			}
			πTemp001[3] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßCATEGORY_WORD); πE != nil {
				continue
			}
			πTemp001[4] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßCATEGORY_NOT_WORD); πE != nil {
				continue
			}
			πTemp001[5] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßCATEGORY_LINEBREAK); πE != nil {
				continue
			}
			πTemp001[6] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßCATEGORY_NOT_LINEBREAK); πE != nil {
				continue
			}
			πTemp001[7] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßCATEGORY_LOC_WORD); πE != nil {
				continue
			}
			πTemp001[8] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßCATEGORY_LOC_NOT_WORD); πE != nil {
				continue
			}
			πTemp001[9] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßCATEGORY_UNI_DIGIT); πE != nil {
				continue
			}
			πTemp001[10] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßCATEGORY_UNI_NOT_DIGIT); πE != nil {
				continue
			}
			πTemp001[11] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßCATEGORY_UNI_SPACE); πE != nil {
				continue
			}
			πTemp001[12] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßCATEGORY_UNI_NOT_SPACE); πE != nil {
				continue
			}
			πTemp001[13] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßCATEGORY_UNI_WORD); πE != nil {
				continue
			}
			πTemp001[14] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßCATEGORY_UNI_NOT_WORD); πE != nil {
				continue
			}
			πTemp001[15] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßCATEGORY_UNI_LINEBREAK); πE != nil {
				continue
			}
			πTemp001[16] = πTemp002
			if πTemp002, πE = πg.ResolveGlobal(πF, ßCATEGORY_UNI_NOT_LINEBREAK); πE != nil {
				continue
			}
			πTemp001[17] = πTemp002
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßCHCODES.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 175: def makedict(list):
			πF.SetLineno(175)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "list", Def: nil}
			πTemp002 = πg.NewFunction(πg.NewCode("makedict", "build/src/__python__/sre_constants.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µlist *πg.Object = πArgs[0]; _ = µlist
				var µd *πg.Object = πg.UnboundLocal; _ = µd
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µitem *πg.Object = πg.UnboundLocal; _ = µitem
				var πTemp001 *πg.Dict
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 176: d = {}
					πF.SetLineno(176)
					πTemp001 = πg.NewDict()
					πTemp002 = πTemp001.ToObject()
					µd = πTemp002
					// line 177: i = 0
					πF.SetLineno(177)
					µi = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µlist); πE != nil {
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
						µitem = πTemp005
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 179: d[item] = i
					πF.SetLineno(179)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, µi); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp006 = µitem
					if πE = πg.SetItem(πF, µd, πTemp006, πTemp005); πE != nil {
						continue
					}
					// line 180: i = i + 1
					πF.SetLineno(180)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µi = πTemp005
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 181: return d
					πF.SetLineno(181)
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					πR = µd
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßmakedict.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 183: OPCODES = makedict(OPCODES)
			πF.SetLineno(183)
			πTemp001 = πF.MakeArgs(1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
				continue
			}
			πTemp001[0] = πTemp004
			if πTemp004, πE = πg.ResolveGlobal(πF, ßmakedict); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßOPCODES.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 184: ATCODES = makedict(ATCODES)
			πF.SetLineno(184)
			πTemp001 = πF.MakeArgs(1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßATCODES); πE != nil {
				continue
			}
			πTemp001[0] = πTemp004
			if πTemp004, πE = πg.ResolveGlobal(πF, ßmakedict); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßATCODES.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 185: CHCODES = makedict(CHCODES)
			πF.SetLineno(185)
			πTemp001 = πF.MakeArgs(1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßCHCODES); πE != nil {
				continue
			}
			πTemp001[0] = πTemp004
			if πTemp004, πE = πg.ResolveGlobal(πF, ßmakedict); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πF.Globals().SetItem(πF, ßCHCODES.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 188: OP_IGNORE = {
			πF.SetLineno(188)
			πTemp003 = πg.NewDict()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßGROUPREF); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßGROUPREF_IGNORE); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßIN); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßIN_IGNORE); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßLITERAL_IGNORE); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßNOT_LITERAL); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNOT_LITERAL_IGNORE); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			πTemp004 = πTemp003.ToObject()
			if πE = πF.Globals().SetItem(πF, ßOP_IGNORE.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 195: AT_MULTILINE = {
			πF.SetLineno(195)
			πTemp003 = πg.NewDict()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßAT_BEGINNING); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßAT_BEGINNING_LINE); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßAT_END); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßAT_END_LINE); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			πTemp004 = πTemp003.ToObject()
			if πE = πF.Globals().SetItem(πF, ßAT_MULTILINE.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 200: AT_LOCALE = {
			πF.SetLineno(200)
			πTemp003 = πg.NewDict()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßAT_BOUNDARY); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßAT_LOC_BOUNDARY); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßAT_NON_BOUNDARY); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßAT_LOC_NON_BOUNDARY); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			πTemp004 = πTemp003.ToObject()
			if πE = πF.Globals().SetItem(πF, ßAT_LOCALE.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 205: AT_UNICODE = {
			πF.SetLineno(205)
			πTemp003 = πg.NewDict()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßAT_BOUNDARY); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßAT_UNI_BOUNDARY); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßAT_NON_BOUNDARY); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßAT_UNI_NON_BOUNDARY); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			πTemp004 = πTemp003.ToObject()
			if πE = πF.Globals().SetItem(πF, ßAT_UNICODE.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 210: CH_LOCALE = {
			πF.SetLineno(210)
			πTemp003 = πg.NewDict()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßCATEGORY_DIGIT); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßCATEGORY_DIGIT); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßCATEGORY_NOT_DIGIT); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßCATEGORY_NOT_DIGIT); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßCATEGORY_SPACE); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßCATEGORY_SPACE); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßCATEGORY_NOT_SPACE); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßCATEGORY_NOT_SPACE); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßCATEGORY_WORD); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßCATEGORY_LOC_WORD); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßCATEGORY_NOT_WORD); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßCATEGORY_LOC_NOT_WORD); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßCATEGORY_LINEBREAK); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßCATEGORY_LINEBREAK); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßCATEGORY_NOT_LINEBREAK); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßCATEGORY_NOT_LINEBREAK); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			πTemp004 = πTemp003.ToObject()
			if πE = πF.Globals().SetItem(πF, ßCH_LOCALE.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 221: CH_UNICODE = {
			πF.SetLineno(221)
			πTemp003 = πg.NewDict()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßCATEGORY_DIGIT); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßCATEGORY_UNI_DIGIT); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßCATEGORY_NOT_DIGIT); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßCATEGORY_UNI_NOT_DIGIT); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßCATEGORY_SPACE); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßCATEGORY_UNI_SPACE); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßCATEGORY_NOT_SPACE); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßCATEGORY_UNI_NOT_SPACE); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßCATEGORY_WORD); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßCATEGORY_UNI_WORD); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßCATEGORY_NOT_WORD); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßCATEGORY_UNI_NOT_WORD); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßCATEGORY_LINEBREAK); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßCATEGORY_UNI_LINEBREAK); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßCATEGORY_NOT_LINEBREAK); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßCATEGORY_UNI_NOT_LINEBREAK); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πTemp004, πTemp005); πE != nil {
				continue
			}
			πTemp004 = πTemp003.ToObject()
			if πE = πF.Globals().SetItem(πF, ßCH_UNICODE.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 233: SRE_FLAG_TEMPLATE = 1 # template mode (disable backtracking)
			πF.SetLineno(233)
			if πE = πF.Globals().SetItem(πF, ßSRE_FLAG_TEMPLATE.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			// line 234: SRE_FLAG_IGNORECASE = 2 # case insensitive
			πF.SetLineno(234)
			if πE = πF.Globals().SetItem(πF, ßSRE_FLAG_IGNORECASE.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
				continue
			}
			// line 235: SRE_FLAG_LOCALE = 4 # honour system locale
			πF.SetLineno(235)
			if πE = πF.Globals().SetItem(πF, ßSRE_FLAG_LOCALE.ToObject(), πg.NewInt(4).ToObject()); πE != nil {
				continue
			}
			// line 236: SRE_FLAG_MULTILINE = 8 # treat target as multiline string
			πF.SetLineno(236)
			if πE = πF.Globals().SetItem(πF, ßSRE_FLAG_MULTILINE.ToObject(), πg.NewInt(8).ToObject()); πE != nil {
				continue
			}
			// line 237: SRE_FLAG_DOTALL = 16 # treat target as a single string
			πF.SetLineno(237)
			if πE = πF.Globals().SetItem(πF, ßSRE_FLAG_DOTALL.ToObject(), πg.NewInt(16).ToObject()); πE != nil {
				continue
			}
			// line 238: SRE_FLAG_UNICODE = 32 # use unicode locale
			πF.SetLineno(238)
			if πE = πF.Globals().SetItem(πF, ßSRE_FLAG_UNICODE.ToObject(), πg.NewInt(32).ToObject()); πE != nil {
				continue
			}
			// line 239: SRE_FLAG_VERBOSE = 64 # ignore whitespace and comments
			πF.SetLineno(239)
			if πE = πF.Globals().SetItem(πF, ßSRE_FLAG_VERBOSE.ToObject(), πg.NewInt(64).ToObject()); πE != nil {
				continue
			}
			// line 240: SRE_FLAG_DEBUG = 128 # debugging
			πF.SetLineno(240)
			if πE = πF.Globals().SetItem(πF, ßSRE_FLAG_DEBUG.ToObject(), πg.NewInt(128).ToObject()); πE != nil {
				continue
			}
			// line 243: SRE_INFO_PREFIX = 1 # has prefix
			πF.SetLineno(243)
			if πE = πF.Globals().SetItem(πF, ßSRE_INFO_PREFIX.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			// line 244: SRE_INFO_LITERAL = 2 # entire pattern is literal (given by prefix)
			πF.SetLineno(244)
			if πE = πF.Globals().SetItem(πF, ßSRE_INFO_LITERAL.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
				continue
			}
			// line 245: SRE_INFO_CHARSET = 4 # pattern starts with character from given set
			πF.SetLineno(245)
			if πE = πF.Globals().SetItem(πF, ßSRE_INFO_CHARSET.ToObject(), πg.NewInt(4).ToObject()); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("sre_constants", Code)
}
