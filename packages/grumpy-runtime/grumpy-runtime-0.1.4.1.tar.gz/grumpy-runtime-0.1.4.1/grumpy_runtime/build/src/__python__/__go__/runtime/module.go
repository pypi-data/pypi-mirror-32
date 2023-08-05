package runtime
import (
	"grumpy"
	"reflect"
	mod "runtime"
)
func fun(f *grumpy.Frame, _ []*grumpy.Object) (*grumpy.Object, *grumpy.BaseException) {
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BlockProfile)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BlockProfile", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.BlockProfileRecord
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "BlockProfileRecord", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Breakpoint)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Breakpoint", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.CPUProfile)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CPUProfile", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Caller)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Caller", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Callers)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Callers", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.CallersFrames)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CallersFrames", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Compiler)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Compiler", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Frame
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Frame", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.Frames
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Frames", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.Func
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Func", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.FuncForPC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FuncForPC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.GC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.GOARCH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GOARCH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.GOMAXPROCS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GOMAXPROCS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.GOOS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GOOS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.GOROOT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GOROOT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Goexit)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Goexit", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.GoroutineProfile)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GoroutineProfile", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Gosched)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Gosched", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.KeepAlive)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "KeepAlive", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.LockOSThread)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LockOSThread", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MemProfile)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MemProfile", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MemProfileRate)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MemProfileRate", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.MemProfileRecord
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "MemProfileRecord", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.MemStats
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "MemStats", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MutexProfile)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MutexProfile", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NumCPU)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NumCPU", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NumCgoCall)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NumCgoCall", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NumGoroutine)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NumGoroutine", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ReadMemStats)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ReadMemStats", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ReadTrace)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ReadTrace", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetBlockProfileRate)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetBlockProfileRate", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetCPUProfileRate)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetCPUProfileRate", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetCgoTraceback)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetCgoTraceback", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetFinalizer)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetFinalizer", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetMutexProfileFraction)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetMutexProfileFraction", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Stack)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Stack", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.StackRecord
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "StackRecord", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.StartTrace)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "StartTrace", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.StopTrace)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "StopTrace", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ThreadCreateProfile)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ThreadCreateProfile", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.TypeAssertionError
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "TypeAssertionError", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.UnlockOSThread)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "UnlockOSThread", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Version)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Version", o); raised != nil {
		return nil, raised
	}

	return nil, nil
}
var Code = grumpy.NewCode("<module>", "runtime", nil, 0, fun)
func init() {
	grumpy.RegisterModule("__go__/runtime", Code)
}
