package rand
import (
	"grumpy"
	"reflect"
	mod "math/rand"
)
func fun(f *grumpy.Frame, _ []*grumpy.Object) (*grumpy.Object, *grumpy.BaseException) {
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ExpFloat64)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ExpFloat64", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Float32)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Float32", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Float64)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Float64", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Int)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Int", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Int31)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Int31", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Int31n)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Int31n", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Int63)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Int63", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Int63n)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Int63n", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Intn)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Intn", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.New)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "New", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewSource)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewSource", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewZipf)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewZipf", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NormFloat64)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NormFloat64", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Perm)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Perm", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Rand
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Rand", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Read)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Read", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Seed)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Seed", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Uint32)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Uint32", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Uint64)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Uint64", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Zipf
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Zipf", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}

	return nil, nil
}
var Code = grumpy.NewCode("<module>", "math/rand", nil, 0, fun)
func init() {
	grumpy.RegisterModule("__go__/math/rand", Code)
}
