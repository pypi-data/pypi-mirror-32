package grumpy
import (
	"grumpy"
	"reflect"
	mod "grumpy"
)
func fun(f *grumpy.Frame, _ []*grumpy.Object) (*grumpy.Object, *grumpy.BaseException) {
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Abs)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Abs", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Add)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Add", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.And)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "And", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Args
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Args", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ArithmeticErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ArithmeticErrorType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Assert)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Assert", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AssertionErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AssertionErrorType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AttributeErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AttributeErrorType", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.BaseException
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "BaseException", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BaseExceptionType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BaseExceptionType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BaseStringType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BaseStringType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BoolType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BoolType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Builtins)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Builtins", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.ByteArray
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "ByteArray", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ByteArrayType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ByteArrayType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BytesWarningType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BytesWarningType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.CheckLocal)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CheckLocal", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ClassMethodType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ClassMethodType", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Code
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Code", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.CodeFlag
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "CodeFlag", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.CodeFlagKWArg))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CodeFlagKWArg", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint(mod.CodeFlagVarArg))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CodeFlagVarArg", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.CodeType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CodeType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Compare)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Compare", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Complex
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Complex", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ComplexType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ComplexType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Contains)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Contains", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DelAttr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DelAttr", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DelItem)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DelItem", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DelVar)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DelVar", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DeprecationWarningType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DeprecationWarningType", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Dict
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Dict", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DictType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DictType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Div)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Div", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DivMod)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DivMod", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EOFErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EOFErrorType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Ellipsis)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Ellipsis", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EllipsisType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EllipsisType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EncodeDefault)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EncodeDefault", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EncodeIgnore)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EncodeIgnore", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EncodeReplace)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EncodeReplace", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EncodeStrict)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EncodeStrict", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EnvironmentErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EnvironmentErrorType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Eq)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Eq", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ExceptionType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ExceptionType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ExceptionTypes)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ExceptionTypes", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.False)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "False", o); raised != nil {
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
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.FileType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FileType", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Float
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Float", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.FloatType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FloatType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.FloorDiv)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FloorDiv", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.FormatExc)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FormatExc", o); raised != nil {
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
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.FrameType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FrameType", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.FrozenSet
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "FrozenSet", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.FrozenSetType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FrozenSetType", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Func
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Func", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.Function
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Function", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.FunctionType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FunctionType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.FutureWarningType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FutureWarningType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.GE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.GT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GT", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Generator
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Generator", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.GeneratorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GeneratorType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.GetAttr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GetAttr", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.GetBool)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GetBool", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.GetItem)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GetItem", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Hash)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Hash", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Hex)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Hex", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IAdd)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IAdd", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IAnd)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IAnd", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IDiv)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IDiv", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFloorDiv)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFloorDiv", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ILShift)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ILShift", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IMod)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IMod", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IMul)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IMul", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IOErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IOErrorType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IOr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IOr", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPow)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPow", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IRShift)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IRShift", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ISub)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ISub", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IXor)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IXor", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ImportErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ImportErrorType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ImportModule)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ImportModule", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ImportWarningType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ImportWarningType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Index)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Index", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IndexErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IndexErrorType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IndexInt)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IndexInt", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Int
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Int", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IntType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IntType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.InternStr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "InternStr", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Invert)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Invert", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Invoke)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Invoke", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IsInstance)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IsInstance", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IsSubclass)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IsSubclass", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IsTrue)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IsTrue", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Iter)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Iter", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.KWArg
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "KWArg", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.KWArgs
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "KWArgs", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.KeyErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "KeyErrorType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.KeyboardInterruptType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "KeyboardInterruptType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.LE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.LShift)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LShift", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.LT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Len)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Len", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.List
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "List", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ListType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ListType", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Long
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Long", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.LongType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LongType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.LookupErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LookupErrorType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(uint64(mod.MaxInt))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MaxInt", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MemoryErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MemoryErrorType", o); raised != nil {
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
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MethodType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MethodType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(int64(mod.MinInt))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MinInt", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Mod)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mod", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Module
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Module", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.ModuleInit
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "ModuleInit", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ModuleType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ModuleType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Mul)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mul", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NameErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NameErrorType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Neg)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Neg", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewCode)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewCode", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewComplex)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewComplex", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewDict)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewDict", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewFileFromFD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewFileFromFD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewFloat)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewFloat", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewFunction)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewFunction", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewGenerator)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewGenerator", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewInt)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewInt", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewList)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewList", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewLong)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewLong", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewLongFromBytes)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewLongFromBytes", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewParamSpec)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewParamSpec", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewRootFrame)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewRootFrame", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewSet)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewSet", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewStr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewStr", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewTryableMutex)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewTryableMutex", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewTuple)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewTuple", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewTuple0)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewTuple0", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewTuple1)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewTuple1", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewTuple2)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewTuple2", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewTuple3)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewTuple3", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewTuple4)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewTuple4", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewTuple5)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewTuple5", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewTuple6)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewTuple6", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewUnicode)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewUnicode", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewUnicodeFromRunes)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewUnicodeFromRunes", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Next)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Next", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.None)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "None", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NoneType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NoneType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NotImplemented)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NotImplemented", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NotImplementedErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NotImplementedErrorType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NotImplementedType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NotImplementedType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.OSErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "OSErrorType", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Object
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Object", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ObjectType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ObjectType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Oct)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Oct", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Or)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Or", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.OverflowErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "OverflowErrorType", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Param
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Param", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.ParamSpec
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "ParamSpec", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PendingDeprecationWarningType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PendingDeprecationWarningType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Pos)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Pos", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Pow)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Pow", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Print)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Print", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Property
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Property", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PropertyType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PropertyType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RShift)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RShift", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ReferenceErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ReferenceErrorType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RegisterModule)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RegisterModule", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Repr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Repr", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ResolveClass)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ResolveClass", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ResolveGlobal)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ResolveGlobal", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RunMain)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RunMain", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.RunState
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "RunState", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RuntimeErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RuntimeErrorType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RuntimeWarningType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RuntimeWarningType", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Set
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Set", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetAttr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetAttr", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetItem)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetItem", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetType", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Slice
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Slice", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SliceType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SliceType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.StandardErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "StandardErrorType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.StartThread)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "StartThread", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.StaticMethodType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "StaticMethodType", o); raised != nil {
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
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.StopIterationType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "StopIterationType", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Str
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Str", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.StrType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "StrType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Sub)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sub", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SyntaxErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SyntaxErrorType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SyntaxWarningType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SyntaxWarningType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SysModules)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SysModules", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SystemErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SystemErrorType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SystemExitType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SystemExitType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ThreadCount)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ThreadCount", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Tie)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Tie", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.TieTarget
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "TieTarget", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ToInt)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ToInt", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ToIntValue)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ToIntValue", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ToNative)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ToNative", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ToStr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ToStr", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Traceback
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Traceback", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TracebackType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TracebackType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.True)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "True", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.TryableMutex
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "TryableMutex", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.Tuple
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Tuple", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TupleType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TupleType", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Type
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Type", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TypeErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TypeErrorType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TypeType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TypeType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.UnboundLocal)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "UnboundLocal", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.UnboundLocalErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "UnboundLocalErrorType", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Unicode
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Unicode", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.UnicodeDecodeErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "UnicodeDecodeErrorType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.UnicodeEncodeErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "UnicodeEncodeErrorType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.UnicodeErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "UnicodeErrorType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.UnicodeType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "UnicodeType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.UnicodeWarningType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "UnicodeWarningType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.UserWarningType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "UserWarningType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ValueErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ValueErrorType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.WarningType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "WarningType", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.WeakRef
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "WeakRef", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.WeakRefType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "WeakRefType", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.WrapNative)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "WrapNative", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Xor)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Xor", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ZeroDivisionErrorType)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ZeroDivisionErrorType", o); raised != nil {
		return nil, raised
	}

	return nil, nil
}
var Code = grumpy.NewCode("<module>", "grumpy", nil, 0, fun)
func init() {
	grumpy.RegisterModule("__go__/grumpy", Code)
}
