package reflect
import (
	"grumpy"
	"reflect"
	mod "reflect"
)
func fun(f *grumpy.Frame, _ []*grumpy.Object) (*grumpy.Object, *grumpy.BaseException) {
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Append)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Append", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AppendSlice)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AppendSlice", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.Array))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Array", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ArrayOf)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ArrayOf", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.Bool))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Bool", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.BothDir))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BothDir", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.Chan))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Chan", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.ChanDir
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "ChanDir", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ChanOf)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ChanOf", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.Complex128))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Complex128", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.Complex64))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Complex64", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Copy)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Copy", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DeepEqual)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DeepEqual", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.Float32))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Float32", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.Float64))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Float64", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.Func))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Func", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.FuncOf)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FuncOf", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Indirect)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Indirect", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.Int))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Int", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.Int16))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Int16", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.Int32))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Int32", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.Int64))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Int64", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.Int8))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Int8", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.Interface))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Interface", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.Invalid))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Invalid", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Kind
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Kind", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MakeChan)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MakeChan", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MakeFunc)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MakeFunc", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MakeMap)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MakeMap", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MakeMapWithSize)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MakeMapWithSize", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MakeSlice)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MakeSlice", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.Map))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Map", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MapOf)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MapOf", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Method
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Method", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.New)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "New", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewAt)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewAt", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.Ptr))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Ptr", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PtrTo)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PtrTo", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.RecvDir))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RecvDir", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Select)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Select", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.SelectCase
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "SelectCase", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.SelectDefault))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SelectDefault", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.SelectDir
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "SelectDir", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.SelectRecv))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SelectRecv", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.SelectSend))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SelectSend", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.SendDir))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SendDir", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.Slice))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Slice", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.SliceHeader
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "SliceHeader", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SliceOf)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SliceOf", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.String))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "String", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.StringHeader
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "StringHeader", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.Struct))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Struct", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.StructField
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "StructField", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.StructOf)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "StructOf", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.StructTag
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "StructTag", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Swapper)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Swapper", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TypeOf)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TypeOf", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.Uint))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Uint", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.Uint16))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Uint16", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.Uint32))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Uint32", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.Uint64))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Uint64", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.Uint8))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Uint8", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.Uintptr))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Uintptr", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.UnsafePointer))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "UnsafePointer", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Value
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Value", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.ValueError
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "ValueError", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ValueOf)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ValueOf", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Zero)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Zero", o); raised != nil {
		return nil, raised
	}

	return nil, nil
}
var Code = grumpy.NewCode("<module>", "reflect", nil, 0, fun)
func init() {
	grumpy.RegisterModule("__go__/reflect", Code)
}
