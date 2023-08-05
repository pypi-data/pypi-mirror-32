package md5
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/md5.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß_md5 := πg.InternStr("_md5")
		ßdigest_size := πg.InternStr("digest_size")
		ßmd5 := πg.InternStr("md5")
		ßnew := πg.InternStr("new")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 11: import _md5
			πF.SetLineno(11)
			if πTemp002, πE = πg.ImportModule(πF, "_md5"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_md5.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 13: new = _md5.new
			πF.SetLineno(13)
			if πTemp001, πE = πg.ResolveGlobal(πF, ß_md5); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßnew, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßnew.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 14: md5 = _md5.new
			πF.SetLineno(14)
			if πTemp001, πE = πg.ResolveGlobal(πF, ß_md5); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßnew, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmd5.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 15: digest_size = 16
			πF.SetLineno(15)
			if πE = πF.Globals().SetItem(πF, ßdigest_size.ToObject(), πg.NewInt(16).ToObject()); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("md5", Code)
}
