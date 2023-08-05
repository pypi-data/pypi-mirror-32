package ioutil
import (
	"grumpy"
	"reflect"
	mod "io/ioutil"
)
func fun(f *grumpy.Frame, _ []*grumpy.Object) (*grumpy.Object, *grumpy.BaseException) {
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Discard)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Discard", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NopCloser)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NopCloser", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ReadAll)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ReadAll", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ReadDir)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ReadDir", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ReadFile)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ReadFile", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TempDir)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TempDir", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TempFile)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TempFile", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.WriteFile)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "WriteFile", o); raised != nil {
		return nil, raised
	}

	return nil, nil
}
var Code = grumpy.NewCode("<module>", "io/ioutil", nil, 0, fun)
func init() {
	grumpy.RegisterModule("__go__/io/ioutil", Code)
}
