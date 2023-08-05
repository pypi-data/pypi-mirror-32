package sync
import (
	"grumpy"
	"reflect"
	mod "sync"
)
func fun(f *grumpy.Frame, _ []*grumpy.Object) (*grumpy.Object, *grumpy.BaseException) {
	if true {
		var x mod.Cond
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Cond", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.Map
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Map", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.Mutex
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Mutex", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewCond)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewCond", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Once
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Once", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.Pool
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Pool", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.RWMutex
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "RWMutex", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.WaitGroup
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "WaitGroup", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}

	return nil, nil
}
var Code = grumpy.NewCode("<module>", "sync", nil, 0, fun)
func init() {
	grumpy.RegisterModule("__go__/sync", Code)
}
