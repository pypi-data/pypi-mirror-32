package test_stat
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "build/src/__python__/test/test_stat.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßBLK := πg.InternStr("BLK")
		ßCHR := πg.InternStr("CHR")
		ßDIR := πg.InternStr("DIR")
		ßFIFO := πg.InternStr("FIFO")
		ßFalse := πg.InternStr("False")
		ßLNK := πg.InternStr("LNK")
		ßNone := πg.InternStr("None")
		ßNotImplementedError := πg.InternStr("NotImplementedError")
		ßOSError := πg.InternStr("OSError")
		ßREG := πg.InternStr("REG")
		ßSF_APPEND := πg.InternStr("SF_APPEND")
		ßSF_ARCHIVED := πg.InternStr("SF_ARCHIVED")
		ßSF_IMMUTABLE := πg.InternStr("SF_IMMUTABLE")
		ßSF_NOUNLINK := πg.InternStr("SF_NOUNLINK")
		ßSF_SNAPSHOT := πg.InternStr("SF_SNAPSHOT")
		ßST_ATIME := πg.InternStr("ST_ATIME")
		ßST_CTIME := πg.InternStr("ST_CTIME")
		ßST_DEV := πg.InternStr("ST_DEV")
		ßST_GID := πg.InternStr("ST_GID")
		ßST_INO := πg.InternStr("ST_INO")
		ßST_MODE := πg.InternStr("ST_MODE")
		ßST_MTIME := πg.InternStr("ST_MTIME")
		ßST_NLINK := πg.InternStr("ST_NLINK")
		ßST_SIZE := πg.InternStr("ST_SIZE")
		ßST_UID := πg.InternStr("ST_UID")
		ßS_ENFMT := πg.InternStr("S_ENFMT")
		ßS_IEXEC := πg.InternStr("S_IEXEC")
		ßS_IFBLK := πg.InternStr("S_IFBLK")
		ßS_IFCHR := πg.InternStr("S_IFCHR")
		ßS_IFDIR := πg.InternStr("S_IFDIR")
		ßS_IFIFO := πg.InternStr("S_IFIFO")
		ßS_IFLNK := πg.InternStr("S_IFLNK")
		ßS_IFMT := πg.InternStr("S_IFMT")
		ßS_IFREG := πg.InternStr("S_IFREG")
		ßS_IFSOCK := πg.InternStr("S_IFSOCK")
		ßS_IMODE := πg.InternStr("S_IMODE")
		ßS_IREAD := πg.InternStr("S_IREAD")
		ßS_IRGRP := πg.InternStr("S_IRGRP")
		ßS_IROTH := πg.InternStr("S_IROTH")
		ßS_IRUSR := πg.InternStr("S_IRUSR")
		ßS_IRWXG := πg.InternStr("S_IRWXG")
		ßS_IRWXO := πg.InternStr("S_IRWXO")
		ßS_IRWXU := πg.InternStr("S_IRWXU")
		ßS_IS := πg.InternStr("S_IS")
		ßS_ISBLK := πg.InternStr("S_ISBLK")
		ßS_ISCHR := πg.InternStr("S_ISCHR")
		ßS_ISDIR := πg.InternStr("S_ISDIR")
		ßS_ISFIFO := πg.InternStr("S_ISFIFO")
		ßS_ISGID := πg.InternStr("S_ISGID")
		ßS_ISLNK := πg.InternStr("S_ISLNK")
		ßS_ISREG := πg.InternStr("S_ISREG")
		ßS_ISSOCK := πg.InternStr("S_ISSOCK")
		ßS_ISUID := πg.InternStr("S_ISUID")
		ßS_ISVTX := πg.InternStr("S_ISVTX")
		ßS_IWGRP := πg.InternStr("S_IWGRP")
		ßS_IWOTH := πg.InternStr("S_IWOTH")
		ßS_IWRITE := πg.InternStr("S_IWRITE")
		ßS_IWUSR := πg.InternStr("S_IWUSR")
		ßS_IXGRP := πg.InternStr("S_IXGRP")
		ßS_IXOTH := πg.InternStr("S_IXOTH")
		ßS_IXUSR := πg.InternStr("S_IXUSR")
		ßSkipTest := πg.InternStr("SkipTest")
		ßTESTFN := πg.InternStr("TESTFN")
		ßTestCase := πg.InternStr("TestCase")
		ßTestFilemode := πg.InternStr("TestFilemode")
		ßUF_APPEND := πg.InternStr("UF_APPEND")
		ßUF_COMPRESSED := πg.InternStr("UF_COMPRESSED")
		ßUF_HIDDEN := πg.InternStr("UF_HIDDEN")
		ßUF_IMMUTABLE := πg.InternStr("UF_IMMUTABLE")
		ßUF_NODUMP := πg.InternStr("UF_NODUMP")
		ßUF_NOUNLINK := πg.InternStr("UF_NOUNLINK")
		ßUF_OPAQUE := πg.InternStr("UF_OPAQUE")
		ßValueError := πg.InternStr("ValueError")
		ß__enter__ := πg.InternStr("__enter__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertFalse := πg.InternStr("assertFalse")
		ßassertIsInstance := πg.InternStr("assertIsInstance")
		ßassertS_IS := πg.InternStr("assertS_IS")
		ßassertTrue := πg.InternStr("assertTrue")
		ßcallable := πg.InternStr("callable")
		ßchmod := πg.InternStr("chmod")
		ßdevnull := πg.InternStr("devnull")
		ßexists := πg.InternStr("exists")
		ßfile_flags := πg.InternStr("file_flags")
		ßformat_funcs := πg.InternStr("format_funcs")
		ßformats := πg.InternStr("formats")
		ßget_mode := πg.InternStr("get_mode")
		ßgetattr := πg.InternStr("getattr")
		ßgetcwd := πg.InternStr("getcwd")
		ßgrumpy := πg.InternStr("grumpy")
		ßhasattr := πg.InternStr("hasattr")
		ßint := πg.InternStr("int")
		ßitems := πg.InternStr("items")
		ßmkdir := πg.InternStr("mkdir")
		ßmkfifo := πg.InternStr("mkfifo")
		ßname := πg.InternStr("name")
		ßopen := πg.InternStr("open")
		ßos := πg.InternStr("os")
		ßpath := πg.InternStr("path")
		ßpermission_bits := πg.InternStr("permission_bits")
		ßposix := πg.InternStr("posix")
		ßremove := πg.InternStr("remove")
		ßrmdir := πg.InternStr("rmdir")
		ßrun_unittest := πg.InternStr("run_unittest")
		ßsetUp := πg.InternStr("setUp")
		ßskip := πg.InternStr("skip")
		ßskipUnless := πg.InternStr("skipUnless")
		ßst_mode := πg.InternStr("st_mode")
		ßstat := πg.InternStr("stat")
		ßstat_struct := πg.InternStr("stat_struct")
		ßstr := πg.InternStr("str")
		ßsymlink := πg.InternStr("symlink")
		ßtearDown := πg.InternStr("tearDown")
		ßtest_devices := πg.InternStr("test_devices")
		ßtest_directory := πg.InternStr("test_directory")
		ßtest_fifo := πg.InternStr("test_fifo")
		ßtest_link := πg.InternStr("test_link")
		ßtest_main := πg.InternStr("test_main")
		ßtest_mode := πg.InternStr("test_mode")
		ßtest_module_attributes := πg.InternStr("test_module_attributes")
		ßunittest := πg.InternStr("unittest")
		ßw := πg.InternStr("w")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.Dict
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
			// line 1: import unittest
			πF.SetLineno(1)
			if πTemp002, πE = πg.ImportModule(πF, "unittest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 2: import os
			πF.SetLineno(2)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 3: from test.test_support import TESTFN, run_unittest
			πF.SetLineno(3)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßTESTFN, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTESTFN.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßrun_unittest, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßrun_unittest.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 4: import stat
			πF.SetLineno(4)
			if πTemp002, πE = πg.ImportModule(πF, "stat"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßstat.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 6: class TestFilemode(unittest.TestCase):
			πF.SetLineno(6)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßTestCase, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TestFilemode", "build/src/__python__/test/test_stat.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Set
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Dict
				_ = πTemp003
				var πTemp004 []πg.Param
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 []*πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 *πg.Object
				_ = πTemp012
				var πTemp013 []*πg.Object
				_ = πTemp013
				var πTemp014 *πg.Object
				_ = πTemp014
				var πTemp015 *πg.Object
				_ = πTemp015
				var πTemp016 *πg.Object
				_ = πTemp016
				var πTemp017 *πg.Object
				_ = πTemp017
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 7: file_flags = {'SF_APPEND', 'SF_ARCHIVED', 'SF_IMMUTABLE', 'SF_NOUNLINK',
					πF.SetLineno(7)
					πTemp001 = πg.NewSet()
					if _, πE = πTemp001.Add(πF, ßSF_APPEND.ToObject()); πE != nil {
						continue
					}
					if _, πE = πTemp001.Add(πF, ßSF_ARCHIVED.ToObject()); πE != nil {
						continue
					}
					if _, πE = πTemp001.Add(πF, ßSF_IMMUTABLE.ToObject()); πE != nil {
						continue
					}
					if _, πE = πTemp001.Add(πF, ßSF_NOUNLINK.ToObject()); πE != nil {
						continue
					}
					if _, πE = πTemp001.Add(πF, ßSF_SNAPSHOT.ToObject()); πE != nil {
						continue
					}
					if _, πE = πTemp001.Add(πF, ßUF_APPEND.ToObject()); πE != nil {
						continue
					}
					if _, πE = πTemp001.Add(πF, ßUF_COMPRESSED.ToObject()); πE != nil {
						continue
					}
					if _, πE = πTemp001.Add(πF, ßUF_HIDDEN.ToObject()); πE != nil {
						continue
					}
					if _, πE = πTemp001.Add(πF, ßUF_IMMUTABLE.ToObject()); πE != nil {
						continue
					}
					if _, πE = πTemp001.Add(πF, ßUF_NODUMP.ToObject()); πE != nil {
						continue
					}
					if _, πE = πTemp001.Add(πF, ßUF_NOUNLINK.ToObject()); πE != nil {
						continue
					}
					if _, πE = πTemp001.Add(πF, ßUF_OPAQUE.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp001.ToObject()
					if πE = πClass.SetItem(πF, ßfile_flags.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 11: formats = {'S_IFBLK', 'S_IFCHR', 'S_IFDIR', 'S_IFIFO', 'S_IFLNK',
					πF.SetLineno(11)
					πTemp001 = πg.NewSet()
					if _, πE = πTemp001.Add(πF, ßS_IFBLK.ToObject()); πE != nil {
						continue
					}
					if _, πE = πTemp001.Add(πF, ßS_IFCHR.ToObject()); πE != nil {
						continue
					}
					if _, πE = πTemp001.Add(πF, ßS_IFDIR.ToObject()); πE != nil {
						continue
					}
					if _, πE = πTemp001.Add(πF, ßS_IFIFO.ToObject()); πE != nil {
						continue
					}
					if _, πE = πTemp001.Add(πF, ßS_IFLNK.ToObject()); πE != nil {
						continue
					}
					if _, πE = πTemp001.Add(πF, ßS_IFREG.ToObject()); πE != nil {
						continue
					}
					if _, πE = πTemp001.Add(πF, ßS_IFSOCK.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp001.ToObject()
					if πE = πClass.SetItem(πF, ßformats.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 14: format_funcs = {'S_ISBLK', 'S_ISCHR', 'S_ISDIR', 'S_ISFIFO', 'S_ISLNK',
					πF.SetLineno(14)
					πTemp001 = πg.NewSet()
					if _, πE = πTemp001.Add(πF, ßS_ISBLK.ToObject()); πE != nil {
						continue
					}
					if _, πE = πTemp001.Add(πF, ßS_ISCHR.ToObject()); πE != nil {
						continue
					}
					if _, πE = πTemp001.Add(πF, ßS_ISDIR.ToObject()); πE != nil {
						continue
					}
					if _, πE = πTemp001.Add(πF, ßS_ISFIFO.ToObject()); πE != nil {
						continue
					}
					if _, πE = πTemp001.Add(πF, ßS_ISLNK.ToObject()); πE != nil {
						continue
					}
					if _, πE = πTemp001.Add(πF, ßS_ISREG.ToObject()); πE != nil {
						continue
					}
					if _, πE = πTemp001.Add(πF, ßS_ISSOCK.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp001.ToObject()
					if πE = πClass.SetItem(πF, ßformat_funcs.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 17: stat_struct = {
					πF.SetLineno(17)
					πTemp003 = πg.NewDict()
					if πE = πTemp003.SetItem(πF, ßST_MODE.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßST_INO.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßST_DEV.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßST_NLINK.ToObject(), πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßST_UID.ToObject(), πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßST_GID.ToObject(), πg.NewInt(5).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßST_SIZE.ToObject(), πg.NewInt(6).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßST_ATIME.ToObject(), πg.NewInt(7).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßST_MTIME.ToObject(), πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßST_CTIME.ToObject(), πg.NewInt(9).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003.ToObject()
					if πE = πClass.SetItem(πF, ßstat_struct.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 30: permission_bits = {
					πF.SetLineno(30)
					πTemp003 = πg.NewDict()
					if πE = πTemp003.SetItem(πF, ßS_ISUID.ToObject(), πg.NewInt(2048).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßS_ISGID.ToObject(), πg.NewInt(1024).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßS_ENFMT.ToObject(), πg.NewInt(1024).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßS_ISVTX.ToObject(), πg.NewInt(512).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßS_IRWXU.ToObject(), πg.NewInt(448).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßS_IRUSR.ToObject(), πg.NewInt(256).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßS_IREAD.ToObject(), πg.NewInt(256).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßS_IWUSR.ToObject(), πg.NewInt(128).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßS_IWRITE.ToObject(), πg.NewInt(128).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßS_IXUSR.ToObject(), πg.NewInt(64).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßS_IEXEC.ToObject(), πg.NewInt(64).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßS_IRWXG.ToObject(), πg.NewInt(56).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßS_IRGRP.ToObject(), πg.NewInt(32).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßS_IWGRP.ToObject(), πg.NewInt(16).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßS_IXGRP.ToObject(), πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßS_IRWXO.ToObject(), πg.NewInt(7).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßS_IROTH.ToObject(), πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßS_IWOTH.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßS_IXOTH.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003.ToObject()
					if πE = πClass.SetItem(πF, ßpermission_bits.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 51: def setUp(self):
					πF.SetLineno(51)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("setUp", "build/src/__python__/test/test_stat.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							case 5: goto Label5
							default: panic("unexpected function state")
							}
							// line 52: try:
							πF.SetLineno(52)
							πF.PushCheckpoint(2)
							// line 53: os.remove(TESTFN)
							πF.SetLineno(53)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTESTFN); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßremove, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
							// line 54: except OSError:
							πF.SetLineno(54)
						Label3:
							// line 55: try:
							πF.SetLineno(55)
							πF.PushCheckpoint(5)
							// line 56: os.rmdir(TESTFN)
							πF.SetLineno(56)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTESTFN); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrmdir, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πF.PopCheckpoint()
							goto Label4
						Label5:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp007, πTemp005 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label6
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 57: except OSError:
							πF.SetLineno(57)
						Label6:
							// line 58: pass
							πF.SetLineno(58)
							πF.RestoreExc(nil, nil)
							goto Label4
						Label4:
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
					if πE = πClass.SetItem(πF, ßsetUp.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 59: tearDown = setUp
					πF.SetLineno(59)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßsetUp); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßtearDown.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 61: def get_mode(self, fname=TESTFN): #, lstat=True):
					πF.SetLineno(61)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßTESTFN); πE != nil {
						continue
					}
					πTemp004[1] = πg.Param{Name: "fname", Def: πTemp006}
					πTemp005 = πg.NewFunction(πg.NewCode("get_mode", "build/src/__python__/test/test_stat.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfname *πg.Object = πArgs[1]; _ = µfname
						var µst_mode *πg.Object = πg.UnboundLocal; _ = µst_mode
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
							// line 65: st_mode = os.stat(fname).st_mode
							πF.SetLineno(65)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfname, "fname"); πE != nil {
								continue
							}
							πTemp001[0] = µfname
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstat, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßst_mode, nil); πE != nil {
								continue
							}
							µst_mode = πTemp003
							// line 66: return st_mode
							πF.SetLineno(66)
							if πE = πg.CheckLocal(πF, µst_mode, "st_mode"); πE != nil {
								continue
							}
							πR = µst_mode
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßget_mode.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 68: def assertS_IS(self, name, mode):
					πF.SetLineno(68)
					πTemp004 = make([]πg.Param, 3)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "name", Def: nil}
					πTemp004[2] = πg.Param{Name: "mode", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("assertS_IS", "build/src/__python__/test/test_stat.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
						var µmode *πg.Object = πArgs[2]; _ = µmode
						var µtestname *πg.Object = πg.UnboundLocal; _ = µtestname
						var µfuncname *πg.Object = πg.UnboundLocal; _ = µfuncname
						var µfunc *πg.Object = πg.UnboundLocal; _ = µfunc
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
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
							// line 73: testname = "S_IS" + name
							πF.SetLineno(73)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, ßS_IS.ToObject(), µname); πE != nil {
								continue
							}
							µtestname = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßformat_funcs, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
								µfuncname = πTemp002
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 75: func = getattr(stat, funcname, None)
							πF.SetLineno(75)
							πTemp005 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µfuncname, "funcname"); πE != nil {
								continue
							}
							πTemp005[1] = µfuncname
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µfunc = πTemp006
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µfunc == πTemp006).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 76: if func is None:
							πF.SetLineno(76)
						Label4:
							if πE = πg.CheckLocal(πF, µfuncname, "funcname"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtestname, "testname"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µfuncname, µtestname); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							goto Label7
							// line 77: if funcname == testname:
							πF.SetLineno(77)
						Label6:
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfuncname, "funcname"); πE != nil {
								continue
							}
							πTemp005[0] = µfuncname
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 78: raise ValueError(funcname)
							πF.SetLineno(78)
							πE = πF.Raise(πTemp006, nil, nil)
							continue
							goto Label7
						Label7:
							// line 79: continue
							πF.SetLineno(79)
							continue
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µfuncname, "funcname"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtestname, "testname"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µfuncname, µtestname); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							goto Label9
							// line 80: if funcname == testname:
							πF.SetLineno(80)
						Label8:
							// line 81: self.assertTrue(func(mode))
							πF.SetLineno(81)
							πTemp005 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
								continue
							}
							πTemp007[0] = µmode
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							if πTemp002, πE = µfunc.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label10
						Label9:
							// line 83: self.assertFalse(func(mode))
							πF.SetLineno(83)
							πTemp005 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
								continue
							}
							πTemp007[0] = µmode
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							if πTemp002, πE = µfunc.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label10
						Label10:
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
					if πE = πClass.SetItem(πF, ßassertS_IS.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 86: def test_mode(self):
					πF.SetLineno(86)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_mode", "build/src/__python__/test/test_stat.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µst_mode *πg.Object = πg.UnboundLocal; _ = µst_mode
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
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
						_ = πTemp008
						var πTemp009 *πg.Type
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 87: with open(TESTFN, 'w'):
							πF.SetLineno(87)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTESTFN); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßw.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp004.Call(πF, πg.Args{πTemp003}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(1)
							// line 88: pass
							πF.SetLineno(88)
							πF.PopCheckpoint()
						Label1:
							πTemp007, πTemp008 = nil, nil
							if πE != nil {
								πTemp007, πTemp008 = πF.ExcInfo()
							}
							if πTemp007 != nil {
								πTemp009 = πTemp007.Type()
								if πTemp005, πE = πTemp002.Call(πF, πg.Args{πTemp003, πTemp009.ToObject(), πTemp007.ToObject(), πTemp008.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp005, πE = πTemp002.Call(πF, πg.Args{πTemp003, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp007 != nil && πTemp006 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßname, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp004, ßposix.ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label2
							}
							goto Label3
							// line 89: if os.name == 'posix':
							πF.SetLineno(89)
						Label2:
							// line 90: os.chmod(TESTFN, 0o700)
							πF.SetLineno(90)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTESTFN); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(448).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßchmod, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 91: st_mode = self.get_mode()
							πF.SetLineno(91)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßget_mode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µst_mode = πTemp003
							// line 92: self.assertS_IS("REG", st_mode)
							πF.SetLineno(92)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßREG.ToObject()
							if πE = πg.CheckLocal(πF, µst_mode, "st_mode"); πE != nil {
								continue
							}
							πTemp001[1] = µst_mode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertS_IS, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 93: self.assertEqual(stat.S_IMODE(st_mode),
							πF.SetLineno(93)
							πTemp001 = πF.MakeArgs(2)
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µst_mode, "st_mode"); πE != nil {
								continue
							}
							πTemp010[0] = µst_mode
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßS_IMODE, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßS_IRWXU, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
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
							// line 96: os.chmod(TESTFN, 0o070)
							πF.SetLineno(96)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTESTFN); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(56).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßchmod, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 97: st_mode = self.get_mode()
							πF.SetLineno(97)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßget_mode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µst_mode = πTemp003
							// line 98: self.assertS_IS("REG", st_mode)
							πF.SetLineno(98)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßREG.ToObject()
							if πE = πg.CheckLocal(πF, µst_mode, "st_mode"); πE != nil {
								continue
							}
							πTemp001[1] = µst_mode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertS_IS, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 99: self.assertEqual(stat.S_IMODE(st_mode),
							πF.SetLineno(99)
							πTemp001 = πF.MakeArgs(2)
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µst_mode, "st_mode"); πE != nil {
								continue
							}
							πTemp010[0] = µst_mode
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßS_IMODE, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßS_IRWXG, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
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
							// line 102: os.chmod(TESTFN, 0o007)
							πF.SetLineno(102)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTESTFN); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(7).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßchmod, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 103: st_mode = self.get_mode()
							πF.SetLineno(103)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßget_mode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µst_mode = πTemp003
							// line 104: self.assertS_IS("REG", st_mode)
							πF.SetLineno(104)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßREG.ToObject()
							if πE = πg.CheckLocal(πF, µst_mode, "st_mode"); πE != nil {
								continue
							}
							πTemp001[1] = µst_mode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertS_IS, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 105: self.assertEqual(stat.S_IMODE(st_mode),
							πF.SetLineno(105)
							πTemp001 = πF.MakeArgs(2)
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µst_mode, "st_mode"); πE != nil {
								continue
							}
							πTemp010[0] = µst_mode
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßS_IMODE, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßS_IRWXO, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
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
							// line 108: os.chmod(TESTFN, 0o444)
							πF.SetLineno(108)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTESTFN); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(292).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßchmod, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 109: st_mode = self.get_mode()
							πF.SetLineno(109)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßget_mode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µst_mode = πTemp003
							// line 110: self.assertS_IS("REG", st_mode)
							πF.SetLineno(110)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßREG.ToObject()
							if πE = πg.CheckLocal(πF, µst_mode, "st_mode"); πE != nil {
								continue
							}
							πTemp001[1] = µst_mode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertS_IS, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 111: self.assertEqual(stat.S_IMODE(st_mode), 0o444)
							πF.SetLineno(111)
							πTemp001 = πF.MakeArgs(2)
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µst_mode, "st_mode"); πE != nil {
								continue
							}
							πTemp010[0] = µst_mode
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßS_IMODE, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(292).ToObject()
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
							goto Label4
						Label3:
							// line 113: os.chmod(TESTFN, 0o700)
							πF.SetLineno(113)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTESTFN); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(448).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßchmod, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 114: st_mode = self.get_mode()
							πF.SetLineno(114)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßget_mode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µst_mode = πTemp003
							// line 115: self.assertS_IS("REG", st_mode)
							πF.SetLineno(115)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßREG.ToObject()
							if πE = πg.CheckLocal(πF, µst_mode, "st_mode"); πE != nil {
								continue
							}
							πTemp001[1] = µst_mode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertS_IS, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 116: self.assertEqual(stat.S_IFMT(st_mode),
							πF.SetLineno(116)
							πTemp001 = πF.MakeArgs(2)
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µst_mode, "st_mode"); πE != nil {
								continue
							}
							πTemp010[0] = µst_mode
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßS_IFMT, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßS_IFREG, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
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
					if πE = πClass.SetItem(πF, ßtest_mode.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 85: @unittest.skip('grumpy')
					πF.SetLineno(85)
					πTemp008 = πF.MakeArgs(1)
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßtest_mode); πE != nil {
						continue
					}
					πTemp008[0] = πTemp009
					πTemp010 = πF.MakeArgs(1)
					πTemp010[0] = ßgrumpy.ToObject()
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp009, ßskip, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp011.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp011, πE = πTemp009.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πE = πClass.SetItem(πF, ßtest_mode.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 119: def test_directory(self):
					πF.SetLineno(119)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("test_directory", "build/src/__python__/test/test_stat.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µst_mode *πg.Object = πg.UnboundLocal; _ = µst_mode
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
							// line 120: os.mkdir(TESTFN)
							πF.SetLineno(120)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTESTFN); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmkdir, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 121: os.chmod(TESTFN, 0o700)
							πF.SetLineno(121)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTESTFN); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(448).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßchmod, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 122: st_mode = self.get_mode()
							πF.SetLineno(122)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßget_mode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µst_mode = πTemp003
							// line 123: self.assertS_IS("DIR", st_mode)
							πF.SetLineno(123)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßDIR.ToObject()
							if πE = πg.CheckLocal(πF, µst_mode, "st_mode"); πE != nil {
								continue
							}
							πTemp001[1] = µst_mode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertS_IS, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_directory.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 127: def test_link(self):
					πF.SetLineno(127)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("test_link", "build/src/__python__/test/test_stat.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µerr *πg.Object = πg.UnboundLocal; _ = µerr
						var µst_mode *πg.Object = πg.UnboundLocal; _ = µst_mode
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 128: try:
							πF.SetLineno(128)
							πF.PushCheckpoint(2)
							// line 129: os.symlink(os.getcwd(), TESTFN)
							πF.SetLineno(129)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßgetcwd, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTESTFN); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsymlink, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πF.PopCheckpoint()
							// line 133: st_mode = self.get_mode()
							πF.SetLineno(133)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßget_mode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µst_mode = πTemp003
							// line 134: self.assertS_IS("LNK", st_mode)
							πF.SetLineno(134)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßLNK.ToObject()
							if πE = πg.CheckLocal(πF, µst_mode, "st_mode"); πE != nil {
								continue
							}
							πTemp001[1] = µst_mode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertS_IS, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, πTemp006).ToObject()
							if πTemp007, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label3
							}
							πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 130: except (OSError, NotImplementedError) as err:
							πF.SetLineno(130)
						Label3:
							µerr = πTemp004.ToObject()
							πTemp001 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp008[0] = µerr
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßSkipTest, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 131: raise unittest.SkipTest(str(err))
							πF.SetLineno(131)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
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
					if πE = πClass.SetItem(πF, ßtest_link.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 125: @unittest.skip('grumpy')
					πF.SetLineno(125)
					πTemp008 = πF.MakeArgs(1)
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßtest_link); πE != nil {
						continue
					}
					πTemp008[0] = πTemp012
					πTemp010 = πF.MakeArgs(2)
					πTemp013 = πF.MakeArgs(2)
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					πTemp013[0] = πTemp012
					πTemp013[1] = ßsymlink.ToObject()
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßhasattr); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp012.Call(πF, πTemp013, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp013)
					πTemp010[0] = πTemp014
					πTemp010[1] = πg.NewStr("os.symlink not available").ToObject()
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp014, πE = πg.GetAttr(πF, πTemp012, ßskipUnless, nil); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp014.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp014, πE = πTemp012.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πE = πClass.SetItem(πF, ßtest_link.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 125: @unittest.skip('grumpy')
					πF.SetLineno(125)
					πTemp008 = πF.MakeArgs(1)
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßtest_link); πE != nil {
						continue
					}
					πTemp008[0] = πTemp012
					πTemp010 = πF.MakeArgs(1)
					πTemp010[0] = ßgrumpy.ToObject()
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp014, πE = πg.GetAttr(πF, πTemp012, ßskip, nil); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp014.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp014, πE = πTemp012.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πE = πClass.SetItem(πF, ßtest_link.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 138: def test_fifo(self):
					πF.SetLineno(138)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("test_fifo", "build/src/__python__/test/test_stat.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µst_mode *πg.Object = πg.UnboundLocal; _ = µst_mode
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
							// line 139: os.mkfifo(TESTFN, 0o700)
							πF.SetLineno(139)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTESTFN); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(448).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmkfifo, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 140: st_mode = self.get_mode()
							πF.SetLineno(140)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßget_mode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µst_mode = πTemp003
							// line 141: self.assertS_IS("FIFO", st_mode)
							πF.SetLineno(141)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßFIFO.ToObject()
							if πE = πg.CheckLocal(πF, µst_mode, "st_mode"); πE != nil {
								continue
							}
							πTemp001[1] = µst_mode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertS_IS, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_fifo.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 136: @unittest.skip('grumpy')
					πF.SetLineno(136)
					πTemp008 = πF.MakeArgs(1)
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßtest_fifo); πE != nil {
						continue
					}
					πTemp008[0] = πTemp014
					πTemp010 = πF.MakeArgs(2)
					πTemp013 = πF.MakeArgs(2)
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					πTemp013[0] = πTemp014
					πTemp013[1] = ßmkfifo.ToObject()
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßhasattr); πE != nil {
						continue
					}
					if πTemp015, πE = πTemp014.Call(πF, πTemp013, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp013)
					πTemp010[0] = πTemp015
					πTemp010[1] = πg.NewStr("os.mkfifo not available").ToObject()
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp015, πE = πg.GetAttr(πF, πTemp014, ßskipUnless, nil); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp015.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp015, πE = πTemp014.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πE = πClass.SetItem(πF, ßtest_fifo.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 136: @unittest.skip('grumpy')
					πF.SetLineno(136)
					πTemp008 = πF.MakeArgs(1)
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßtest_fifo); πE != nil {
						continue
					}
					πTemp008[0] = πTemp014
					πTemp010 = πF.MakeArgs(1)
					πTemp010[0] = ßgrumpy.ToObject()
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp015, πE = πg.GetAttr(πF, πTemp014, ßskip, nil); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp015.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp015, πE = πTemp014.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πE = πClass.SetItem(πF, ßtest_fifo.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 145: def test_devices(self):
					πF.SetLineno(145)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("test_devices", "build/src/__python__/test/test_stat.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µst_mode *πg.Object = πg.UnboundLocal; _ = µst_mode
						var µblockdev *πg.Object = πg.UnboundLocal; _ = µblockdev
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 πg.KWArgs
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 3: goto Label3
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdevnull, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßexists, nil); πE != nil {
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
							// line 146: if os.path.exists(os.devnull):
							πF.SetLineno(146)
						Label1:
							// line 147: st_mode = self.get_mode(os.devnull, lstat=False)
							πF.SetLineno(147)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdevnull, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"lstat", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßget_mode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µst_mode = πTemp003
							// line 148: self.assertS_IS("CHR", st_mode)
							πF.SetLineno(148)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßCHR.ToObject()
							if πE = πg.CheckLocal(πF, µst_mode, "st_mode"); πE != nil {
								continue
							}
							πTemp001[1] = µst_mode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertS_IS, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label2
						Label2:
							πTemp003 = πg.NewTuple2(πg.NewStr("/dev/sda").ToObject(), πg.NewStr("/dev/hda").ToObject()).ToObject()
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
								continue
							}
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
							if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µblockdev = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(3)            
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µblockdev, "blockdev"); πE != nil {
								continue
							}
							πTemp001[0] = µblockdev
							if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp007, ßexists, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp006, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label6
							}
							goto Label7
							// line 151: if os.path.exists(blockdev):
							πF.SetLineno(151)
						Label6:
							// line 152: st_mode = self.get_mode(blockdev, lstat=False)
							πF.SetLineno(152)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µblockdev, "blockdev"); πE != nil {
								continue
							}
							πTemp001[0] = µblockdev
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"lstat", πTemp003},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßget_mode, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp001, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µst_mode = πTemp007
							// line 153: self.assertS_IS("BLK", st_mode)
							πF.SetLineno(153)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßBLK.ToObject()
							if πE = πg.CheckLocal(πF, µst_mode, "st_mode"); πE != nil {
								continue
							}
							πTemp001[1] = µst_mode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertS_IS, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 154: break
							πF.SetLineno(154)
							πTemp004 = true
							continue
							goto Label7
						Label7:
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
					if πE = πClass.SetItem(πF, ßtest_devices.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 143: @unittest.skip('grumpy')
					πF.SetLineno(143)
					πTemp008 = πF.MakeArgs(1)
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßtest_devices); πE != nil {
						continue
					}
					πTemp008[0] = πTemp015
					πTemp010 = πF.MakeArgs(2)
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					if πTemp017, πE = πg.GetAttr(πF, πTemp016, ßname, nil); πE != nil {
						continue
					}
					if πTemp015, πE = πg.Eq(πF, πTemp017, ßposix.ToObject()); πE != nil {
						continue
					}
					πTemp010[0] = πTemp015
					πTemp010[1] = πg.NewStr("requires Posix").ToObject()
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp016, πE = πg.GetAttr(πF, πTemp015, ßskipUnless, nil); πE != nil {
						continue
					}
					if πTemp015, πE = πTemp016.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp016, πE = πTemp015.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πE = πClass.SetItem(πF, ßtest_devices.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 143: @unittest.skip('grumpy')
					πF.SetLineno(143)
					πTemp008 = πF.MakeArgs(1)
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßtest_devices); πE != nil {
						continue
					}
					πTemp008[0] = πTemp015
					πTemp010 = πF.MakeArgs(1)
					πTemp010[0] = ßgrumpy.ToObject()
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp016, πE = πg.GetAttr(πF, πTemp015, ßskip, nil); πE != nil {
						continue
					}
					if πTemp015, πE = πTemp016.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp016, πE = πTemp015.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πE = πClass.SetItem(πF, ßtest_devices.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 157: def test_module_attributes(self):
					πF.SetLineno(157)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("test_module_attributes", "build/src/__python__/test/test_stat.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
						var µmodvalue *πg.Object = πg.UnboundLocal; _ = µmodvalue
						var µfunc *πg.Object = πg.UnboundLocal; _ = µfunc
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
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
							case 10: goto Label10
							case 11: goto Label11
							case 13: goto Label13
							case 14: goto Label14
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstat_struct, nil); πE != nil {
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
								µkey = πTemp003
								µvalue = πTemp006
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 159: modvalue = getattr(stat, key)
							πF.SetLineno(159)
							πTemp007 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp007[1] = µkey
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µmodvalue = πTemp003
							// line 160: self.assertEqual(value, modvalue, key)
							πF.SetLineno(160)
							πTemp007 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp007[0] = µvalue
							if πE = πg.CheckLocal(πF, µmodvalue, "modvalue"); πE != nil {
								continue
							}
							πTemp007[1] = µmodvalue
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp007[2] = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpermission_bits, nil); πE != nil {
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
							πF.PushCheckpoint(5)
							πTemp004 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label6
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
								µkey = πTemp003
								µvalue = πTemp006
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 162: modvalue = getattr(stat, key)
							πF.SetLineno(162)
							πTemp007 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp007[1] = µkey
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µmodvalue = πTemp003
							// line 163: self.assertEqual(value, modvalue, key)
							πF.SetLineno(163)
							πTemp007 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp007[0] = µvalue
							if πE = πg.CheckLocal(πF, µmodvalue, "modvalue"); πE != nil {
								continue
							}
							πTemp007[1] = µmodvalue
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp007[2] = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfile_flags, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
								continue
							}
							πF.PushCheckpoint(8)
							πTemp004 = false
						Label7:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label9
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
								µkey = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(7)            
							// line 165: modvalue = getattr(stat, key)
							πF.SetLineno(165)
							πTemp007 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp007[1] = µkey
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µmodvalue = πTemp003
							// line 166: self.assertIsInstance(modvalue, int)
							πF.SetLineno(166)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmodvalue, "modvalue"); πE != nil {
								continue
							}
							πTemp007[0] = µmodvalue
							if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							πTemp007[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIsInstance, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							continue
						Label8:
							if πE != nil || πR != nil {
								continue
							}
						Label9:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßformats, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
								continue
							}
							πF.PushCheckpoint(11)
							πTemp004 = false
						Label10:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label12
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
								µkey = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(10)            
							// line 168: modvalue = getattr(stat, key)
							πF.SetLineno(168)
							πTemp007 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp007[1] = µkey
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µmodvalue = πTemp003
							// line 169: self.assertIsInstance(modvalue, int)
							πF.SetLineno(169)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmodvalue, "modvalue"); πE != nil {
								continue
							}
							πTemp007[0] = µmodvalue
							if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							πTemp007[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIsInstance, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							continue
						Label11:
							if πE != nil || πR != nil {
								continue
							}
						Label12:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßformat_funcs, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
								continue
							}
							πF.PushCheckpoint(14)
							πTemp004 = false
						Label13:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label15
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
								µkey = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(13)            
							// line 171: func = getattr(stat, key)
							πF.SetLineno(171)
							πTemp007 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstat); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp007[1] = µkey
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µfunc = πTemp003
							// line 172: self.assertTrue(callable(func))
							πF.SetLineno(172)
							πTemp007 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							πTemp008[0] = µfunc
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcallable); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp007[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 173: self.assertEqual(func(0), 0)
							πF.SetLineno(173)
							πTemp007 = πF.MakeArgs(2)
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							if πTemp002, πE = µfunc.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp007[0] = πTemp002
							πTemp007[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							continue
						Label14:
							if πE != nil || πR != nil {
								continue
							}
						Label15:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_module_attributes.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 156: @unittest.skip('grumpy')
					πF.SetLineno(156)
					πTemp008 = πF.MakeArgs(1)
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßtest_module_attributes); πE != nil {
						continue
					}
					πTemp008[0] = πTemp016
					πTemp010 = πF.MakeArgs(1)
					πTemp010[0] = ßgrumpy.ToObject()
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp017, πE = πg.GetAttr(πF, πTemp016, ßskip, nil); πE != nil {
						continue
					}
					if πTemp016, πE = πTemp017.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp017, πE = πTemp016.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πE = πClass.SetItem(πF, ßtest_module_attributes.ToObject(), πTemp017); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("TestFilemode").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTestFilemode.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 176: def test_main():
			πF.SetLineno(176)
			πTemp007 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("test_main", "build/src/__python__/test/test_stat.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 177: run_unittest(TestFilemode)
					πF.SetLineno(177)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTestFilemode); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßrun_unittest); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßtest_main.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp003, πE = πg.Eq(πF, πTemp005, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
				continue
			}
			if πTemp008 {
				goto Label1
			}
			goto Label2
			// line 179: if __name__ == '__main__':
			πF.SetLineno(179)
		Label1:
			// line 180: test_main()
			πF.SetLineno(180)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_main); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("test.test_stat", Code)
}
