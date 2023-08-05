package os
import (
	"grumpy"
	"reflect"
	mod "os"
)
func fun(f *grumpy.Frame, _ []*grumpy.Object) (*grumpy.Object, *grumpy.BaseException) {
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Args)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Args", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Chdir)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Chdir", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Chmod)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Chmod", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Chown)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Chown", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Chtimes)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Chtimes", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Clearenv)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Clearenv", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Create)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Create", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DevNull)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DevNull", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Environ)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Environ", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ErrClosed)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ErrClosed", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ErrExist)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ErrExist", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ErrInvalid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ErrInvalid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ErrNotExist)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ErrNotExist", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ErrPermission)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ErrPermission", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Executable)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Executable", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Exit)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Exit", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Expand)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Expand", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ExpandEnv)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ExpandEnv", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.File
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "File", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.FileMode
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "FileMode", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.FindProcess)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FindProcess", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getegid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getegid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getenv)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getenv", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Geteuid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Geteuid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getgid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getgid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getgroups)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getgroups", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getpagesize)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getpagesize", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getpid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getpid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getppid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getppid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getuid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getuid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getwd)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getwd", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Hostname)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Hostname", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Interrupt)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Interrupt", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IsExist)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IsExist", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IsNotExist)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IsNotExist", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IsPathSeparator)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IsPathSeparator", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IsPermission)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IsPermission", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Kill)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Kill", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Lchown)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Lchown", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Link)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Link", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.LinkError
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "LinkError", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.LookupEnv)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LookupEnv", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Lstat)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Lstat", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Mkdir)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mkdir", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MkdirAll)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MkdirAll", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ModeAppend)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ModeAppend", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ModeCharDevice)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ModeCharDevice", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ModeDevice)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ModeDevice", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ModeDir)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ModeDir", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ModeExclusive)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ModeExclusive", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ModeNamedPipe)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ModeNamedPipe", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ModePerm)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ModePerm", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ModeSetgid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ModeSetgid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ModeSetuid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ModeSetuid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ModeSocket)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ModeSocket", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ModeSticky)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ModeSticky", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ModeSymlink)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ModeSymlink", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ModeTemporary)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ModeTemporary", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ModeType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ModeType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewFile)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewFile", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewSyscallError)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewSyscallError", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_APPEND)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_APPEND", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_CREATE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_CREATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_EXCL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_EXCL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_RDONLY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_RDONLY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_RDWR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_RDWR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_SYNC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_SYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_TRUNC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_TRUNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_WRONLY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_WRONLY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Open)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Open", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.OpenFile)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "OpenFile", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.PathError
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "PathError", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PathListSeparator)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PathListSeparator", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PathSeparator)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PathSeparator", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Pipe)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Pipe", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.ProcAttr
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "ProcAttr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.Process
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Process", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.ProcessState
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "ProcessState", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Readlink)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Readlink", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Remove)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Remove", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RemoveAll)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RemoveAll", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Rename)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Rename", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SEEK_CUR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SEEK_CUR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SEEK_END)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SEEK_END", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SEEK_SET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SEEK_SET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SameFile)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SameFile", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Setenv)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setenv", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.StartProcess)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "StartProcess", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Stat)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Stat", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Stderr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Stderr", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Stdin)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Stdin", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Stdout)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Stdout", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Symlink)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Symlink", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.SyscallError
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "SyscallError", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TempDir)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TempDir", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Truncate)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Truncate", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Unsetenv)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Unsetenv", o); raised != nil {
		return nil, raised
	}

	return nil, nil
}
var Code = grumpy.NewCode("<module>", "os", nil, 0, fun)
func init() {
	grumpy.RegisterModule("__go__/os", Code)
}
