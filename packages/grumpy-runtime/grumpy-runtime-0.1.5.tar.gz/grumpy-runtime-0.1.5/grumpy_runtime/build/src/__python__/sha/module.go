package sha
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/sha.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß_sha := πg.InternStr("_sha")
		ßblocksize := πg.InternStr("blocksize")
		ßdigest_size := πg.InternStr("digest_size")
		ßdigestsize := πg.InternStr("digestsize")
		ßnew := πg.InternStr("new")
		ßsha := πg.InternStr("sha")
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
			// line 10: import _sha
			πF.SetLineno(10)
			if πTemp002, πE = πg.ImportModule(πF, "_sha"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_sha.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 12: sha = _sha.new
			πF.SetLineno(12)
			if πTemp001, πE = πg.ResolveGlobal(πF, ß_sha); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßnew, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßsha.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 13: new = _sha.new
			πF.SetLineno(13)
			if πTemp001, πE = πg.ResolveGlobal(πF, ß_sha); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßnew, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßnew.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 15: blocksize = 1        # legacy value (wrong in any useful sense)
			πF.SetLineno(15)
			if πE = πF.Globals().SetItem(πF, ßblocksize.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			// line 16: digest_size = 20
			πF.SetLineno(16)
			if πE = πF.Globals().SetItem(πF, ßdigest_size.ToObject(), πg.NewInt(20).ToObject()); πE != nil {
				continue
			}
			// line 17: digestsize = 20
			πF.SetLineno(17)
			if πE = πF.Globals().SetItem(πF, ßdigestsize.ToObject(), πg.NewInt(20).ToObject()); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("sha", Code)
}
