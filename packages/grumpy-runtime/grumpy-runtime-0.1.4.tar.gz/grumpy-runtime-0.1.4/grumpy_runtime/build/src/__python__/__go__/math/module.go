package math
import (
	"grumpy"
	"reflect"
	mod "math"
)
func fun(f *grumpy.Frame, _ []*grumpy.Object) (*grumpy.Object, *grumpy.BaseException) {
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Abs)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Abs", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Acos)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Acos", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Acosh)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Acosh", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Asin)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Asin", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Asinh)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Asinh", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Atan)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Atan", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Atan2)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Atan2", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Atanh)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Atanh", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Cbrt)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Cbrt", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Ceil)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Ceil", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Copysign)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Copysign", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Cos)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Cos", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Cosh)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Cosh", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Dim)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Dim", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(float64(mod.E))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "E", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Erf)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Erf", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Erfc)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Erfc", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Exp)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Exp", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Exp2)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Exp2", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Expm1)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Expm1", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Float32bits)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Float32bits", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Float32frombits)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Float32frombits", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Float64bits)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Float64bits", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Float64frombits)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Float64frombits", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Floor)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Floor", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Frexp)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Frexp", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Gamma)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Gamma", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Hypot)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Hypot", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Ilogb)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Ilogb", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Inf)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Inf", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IsInf)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IsInf", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IsNaN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IsNaN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.J0)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "J0", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.J1)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "J1", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Jn)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Jn", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Ldexp)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Ldexp", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Lgamma)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Lgamma", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(float64(mod.Ln10))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Ln10", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(float64(mod.Ln2))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Ln2", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Log)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Log", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Log10)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Log10", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(float64(mod.Log10E))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Log10E", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Log1p)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Log1p", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Log2)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Log2", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(float64(mod.Log2E))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Log2E", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Logb)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Logb", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Max)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Max", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(float64(mod.MaxFloat32))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MaxFloat32", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(float64(mod.MaxFloat64))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MaxFloat64", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MaxInt16)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MaxInt16", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MaxInt32)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MaxInt32", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MaxInt64)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MaxInt64", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MaxInt8)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MaxInt8", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MaxUint16)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MaxUint16", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MaxUint32)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MaxUint32", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint64(mod.MaxUint64))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MaxUint64", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MaxUint8)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MaxUint8", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Min)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Min", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MinInt16)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MinInt16", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MinInt32)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MinInt32", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MinInt64)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MinInt64", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MinInt8)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MinInt8", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Mod)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mod", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Modf)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Modf", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NaN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NaN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Nextafter)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Nextafter", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Nextafter32)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Nextafter32", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(float64(mod.Phi))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Phi", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(float64(mod.Pi))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Pi", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Pow)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Pow", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Pow10)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Pow10", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Remainder)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Remainder", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Signbit)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Signbit", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Sin)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sin", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Sincos)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sincos", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Sinh)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sinh", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(float64(mod.SmallestNonzeroFloat32))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SmallestNonzeroFloat32", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(float64(mod.SmallestNonzeroFloat64))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SmallestNonzeroFloat64", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Sqrt)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sqrt", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(float64(mod.Sqrt2))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sqrt2", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(float64(mod.SqrtE))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SqrtE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(float64(mod.SqrtPhi))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SqrtPhi", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(float64(mod.SqrtPi))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SqrtPi", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Tan)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Tan", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Tanh)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Tanh", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Trunc)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Trunc", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Y0)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Y0", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Y1)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Y1", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Yn)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Yn", o); raised != nil {
		return nil, raised
	}

	return nil, nil
}
var Code = grumpy.NewCode("<module>", "math", nil, 0, fun)
func init() {
	grumpy.RegisterModule("__go__/math", Code)
}
