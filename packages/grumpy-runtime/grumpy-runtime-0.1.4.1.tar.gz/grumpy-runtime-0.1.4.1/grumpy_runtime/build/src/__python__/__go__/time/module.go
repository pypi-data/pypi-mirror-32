package time
import (
	"grumpy"
	"reflect"
	mod "time"
)
func fun(f *grumpy.Frame, _ []*grumpy.Object) (*grumpy.Object, *grumpy.BaseException) {
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ANSIC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ANSIC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.After)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "After", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AfterFunc)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AfterFunc", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.April)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "April", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.August)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "August", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Date)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Date", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.December)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "December", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Duration
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Duration", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.February)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "February", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.FixedZone)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FixedZone", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Friday)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Friday", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Hour)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Hour", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.January)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "January", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.July)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "July", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.June)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "June", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Kitchen)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Kitchen", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.LoadLocation)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LoadLocation", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Local)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Local", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Location
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Location", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.March)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "March", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.May)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "May", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Microsecond)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Microsecond", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Millisecond)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Millisecond", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Minute)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Minute", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Monday)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Monday", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Month
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Month", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Nanosecond)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Nanosecond", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewTicker)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewTicker", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NewTimer)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewTimer", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.November)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "November", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Now)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Now", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.October)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "October", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Parse)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Parse", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ParseDuration)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ParseDuration", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.ParseError
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "ParseError", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ParseInLocation)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ParseInLocation", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RFC1123)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RFC1123", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RFC1123Z)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RFC1123Z", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RFC3339)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RFC3339", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RFC3339Nano)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RFC3339Nano", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RFC822)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RFC822", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RFC822Z)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RFC822Z", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RFC850)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RFC850", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RubyDate)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RubyDate", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Saturday)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Saturday", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Second)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Second", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.September)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "September", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Since)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Since", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Sleep)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sleep", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Stamp)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Stamp", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.StampMicro)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "StampMicro", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.StampMilli)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "StampMilli", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.StampNano)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "StampNano", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Sunday)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sunday", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Thursday)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Thursday", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Tick)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Tick", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Ticker
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Ticker", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.Time
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Time", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.Timer
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Timer", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Tuesday)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Tuesday", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.UTC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "UTC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Unix)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Unix", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.UnixDate)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "UnixDate", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Until)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Until", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Wednesday)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Wednesday", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Weekday
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Weekday", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}

	return nil, nil
}
var Code = grumpy.NewCode("<module>", "time", nil, 0, fun)
func init() {
	grumpy.RegisterModule("__go__/time", Code)
}
