package unicode
import (
	"grumpy"
	"reflect"
	mod "unicode"
)
func fun(f *grumpy.Frame, _ []*grumpy.Object) (*grumpy.Object, *grumpy.BaseException) {
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ASCII_Hex_Digit)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ASCII_Hex_Digit", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Adlam)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Adlam", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Ahom)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Ahom", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Anatolian_Hieroglyphs)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Anatolian_Hieroglyphs", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Arabic)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Arabic", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Armenian)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Armenian", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Avestan)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Avestan", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AzeriCase)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AzeriCase", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Balinese)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Balinese", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Bamum)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Bamum", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Bassa_Vah)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Bassa_Vah", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Batak)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Batak", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Bengali)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Bengali", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Bhaiksuki)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Bhaiksuki", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Bidi_Control)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Bidi_Control", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Bopomofo)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Bopomofo", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Brahmi)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Brahmi", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Braille)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Braille", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Buginese)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Buginese", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Buhid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Buhid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.C)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "C", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Canadian_Aboriginal)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Canadian_Aboriginal", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Carian)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Carian", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.CaseRange
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "CaseRange", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.CaseRanges)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CaseRanges", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Categories)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Categories", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Caucasian_Albanian)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Caucasian_Albanian", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Cc)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Cc", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Cf)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Cf", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Chakma)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Chakma", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Cham)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Cham", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Cherokee)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Cherokee", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Co)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Co", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Common)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Common", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Coptic)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Coptic", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Cs)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Cs", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Cuneiform)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Cuneiform", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Cypriot)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Cypriot", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Cyrillic)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Cyrillic", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Dash)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Dash", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Deprecated)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Deprecated", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Deseret)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Deseret", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Devanagari)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Devanagari", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Diacritic)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Diacritic", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Digit)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Digit", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Duployan)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Duployan", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Egyptian_Hieroglyphs)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Egyptian_Hieroglyphs", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Elbasan)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Elbasan", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Ethiopic)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Ethiopic", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Extender)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Extender", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.FoldCategory)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FoldCategory", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.FoldScript)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FoldScript", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Georgian)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Georgian", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Glagolitic)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Glagolitic", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Gothic)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Gothic", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Grantha)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Grantha", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.GraphicRanges)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GraphicRanges", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Greek)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Greek", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Gujarati)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Gujarati", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Gurmukhi)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Gurmukhi", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Han)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Han", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Hangul)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Hangul", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Hanunoo)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Hanunoo", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Hatran)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Hatran", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Hebrew)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Hebrew", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Hex_Digit)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Hex_Digit", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Hiragana)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Hiragana", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Hyphen)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Hyphen", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IDS_Binary_Operator)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IDS_Binary_Operator", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IDS_Trinary_Operator)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IDS_Trinary_Operator", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Ideographic)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Ideographic", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Imperial_Aramaic)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Imperial_Aramaic", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.In)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "In", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Inherited)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Inherited", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Inscriptional_Pahlavi)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Inscriptional_Pahlavi", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Inscriptional_Parthian)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Inscriptional_Parthian", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Is)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Is", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IsControl)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IsControl", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IsDigit)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IsDigit", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IsGraphic)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IsGraphic", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IsLetter)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IsLetter", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IsLower)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IsLower", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IsMark)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IsMark", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IsNumber)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IsNumber", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IsOneOf)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IsOneOf", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IsPrint)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IsPrint", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IsPunct)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IsPunct", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IsSpace)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IsSpace", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IsSymbol)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IsSymbol", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IsTitle)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IsTitle", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IsUpper)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IsUpper", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Javanese)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Javanese", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Join_Control)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Join_Control", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Kaithi)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Kaithi", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Kannada)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Kannada", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Katakana)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Katakana", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Kayah_Li)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Kayah_Li", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Kharoshthi)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Kharoshthi", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Khmer)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Khmer", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Khojki)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Khojki", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Khudawadi)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Khudawadi", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.L)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "L", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Lao)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Lao", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Latin)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Latin", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Lepcha)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Lepcha", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Letter)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Letter", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Limbu)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Limbu", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Linear_A)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Linear_A", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Linear_B)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Linear_B", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Lisu)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Lisu", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Ll)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Ll", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Lm)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Lm", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Lo)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Lo", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Logical_Order_Exception)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Logical_Order_Exception", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Lower)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Lower", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.LowerCase))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LowerCase", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Lt)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Lt", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Lu)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Lu", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Lycian)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Lycian", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Lydian)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Lydian", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.M)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "M", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Mahajani)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mahajani", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Malayalam)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Malayalam", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Mandaic)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mandaic", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Manichaean)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Manichaean", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Marchen)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Marchen", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Mark)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mark", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint32(mod.MaxASCII))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MaxASCII", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.MaxCase))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MaxCase", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint32(mod.MaxLatin1))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MaxLatin1", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint32(mod.MaxRune))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MaxRune", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Mc)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mc", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Me)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Me", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Meetei_Mayek)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Meetei_Mayek", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Mende_Kikakui)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mende_Kikakui", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Meroitic_Cursive)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Meroitic_Cursive", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Meroitic_Hieroglyphs)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Meroitic_Hieroglyphs", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Miao)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Miao", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Mn)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mn", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Modi)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Modi", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Mongolian)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mongolian", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Mro)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mro", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Multani)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Multani", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Myanmar)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Myanmar", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.N)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "N", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Nabataean)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Nabataean", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Nd)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Nd", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.New_Tai_Lue)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "New_Tai_Lue", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Newa)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Newa", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Nko)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Nko", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Nl)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Nl", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.No)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "No", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Noncharacter_Code_Point)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Noncharacter_Code_Point", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Number)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Number", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Ogham)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Ogham", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Ol_Chiki)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Ol_Chiki", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Old_Hungarian)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Old_Hungarian", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Old_Italic)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Old_Italic", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Old_North_Arabian)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Old_North_Arabian", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Old_Permic)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Old_Permic", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Old_Persian)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Old_Persian", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Old_South_Arabian)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Old_South_Arabian", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Old_Turkic)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Old_Turkic", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Oriya)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Oriya", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Osage)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Osage", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Osmanya)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Osmanya", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Other)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Other", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Other_Alphabetic)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Other_Alphabetic", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Other_Default_Ignorable_Code_Point)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Other_Default_Ignorable_Code_Point", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Other_Grapheme_Extend)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Other_Grapheme_Extend", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Other_ID_Continue)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Other_ID_Continue", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Other_ID_Start)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Other_ID_Start", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Other_Lowercase)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Other_Lowercase", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Other_Math)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Other_Math", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Other_Uppercase)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Other_Uppercase", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.P)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "P", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Pahawh_Hmong)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Pahawh_Hmong", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Palmyrene)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Palmyrene", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Pattern_Syntax)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Pattern_Syntax", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Pattern_White_Space)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Pattern_White_Space", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Pau_Cin_Hau)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Pau_Cin_Hau", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Pc)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Pc", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Pd)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Pd", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Pe)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Pe", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Pf)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Pf", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Phags_Pa)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Phags_Pa", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Phoenician)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Phoenician", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Pi)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Pi", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Po)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Po", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Prepended_Concatenation_Mark)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Prepended_Concatenation_Mark", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PrintRanges)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PrintRanges", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Properties)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Properties", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Ps)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Ps", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Psalter_Pahlavi)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Psalter_Pahlavi", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Punct)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Punct", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Quotation_Mark)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Quotation_Mark", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Radical)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Radical", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Range16
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Range16", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.Range32
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Range32", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.RangeTable
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "RangeTable", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Rejang)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Rejang", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint32(mod.ReplacementChar))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ReplacementChar", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Runic)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Runic", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.STerm)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "STerm", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Samaritan)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Samaritan", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Saurashtra)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Saurashtra", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Sc)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sc", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Scripts)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Scripts", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Sentence_Terminal)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sentence_Terminal", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Sharada)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sharada", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Shavian)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Shavian", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Siddham)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Siddham", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SignWriting)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SignWriting", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SimpleFold)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SimpleFold", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Sinhala)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sinhala", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Sk)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sk", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Sm)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sm", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.So)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "So", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Soft_Dotted)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Soft_Dotted", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Sora_Sompeng)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sora_Sompeng", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Space)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Space", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.SpecialCase
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "SpecialCase", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Sundanese)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sundanese", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Syloti_Nagri)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Syloti_Nagri", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Symbol)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Symbol", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Syriac)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Syriac", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Tagalog)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Tagalog", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Tagbanwa)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Tagbanwa", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Tai_Le)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Tai_Le", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Tai_Tham)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Tai_Tham", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Tai_Viet)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Tai_Viet", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Takri)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Takri", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Tamil)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Tamil", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Tangut)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Tangut", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Telugu)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Telugu", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Terminal_Punctuation)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Terminal_Punctuation", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Thaana)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Thaana", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Thai)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Thai", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Tibetan)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Tibetan", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Tifinagh)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Tifinagh", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Tirhuta)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Tirhuta", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Title)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Title", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.TitleCase))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TitleCase", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.To)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "To", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ToLower)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ToLower", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ToTitle)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ToTitle", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ToUpper)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ToUpper", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TurkishCase)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TurkishCase", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Ugaritic)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Ugaritic", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Unified_Ideograph)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Unified_Ideograph", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Upper)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Upper", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.UpperCase))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "UpperCase", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint32(mod.UpperLower))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "UpperLower", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Vai)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Vai", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Variation_Selector)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Variation_Selector", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Version)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Version", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Warang_Citi)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Warang_Citi", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.White_Space)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "White_Space", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Yi)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Yi", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Z)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Z", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Zl)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Zl", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Zp)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Zp", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Zs)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Zs", o); raised != nil {
		return nil, raised
	}

	return nil, nil
}
var Code = grumpy.NewCode("<module>", "unicode", nil, 0, fun)
func init() {
	grumpy.RegisterModule("__go__/unicode", Code)
}
