package comp

import "fmt"

const (
	ParamTypeString   = "String"
	ParamTypeInt      = "Int"
	ParamTypeBool     = "Bool"
	ParamTypeDuration = "Duration"
)

type ActualParam struct {
	Name      string
	ParamType string

	s *string
	i *int
	b *bool
	d *int64
}

func NewUnnamedParamString(news string) *ActualParam {
	return &ActualParam{s: &news, ParamType: ParamTypeString}
}

func NewUnnamedParamInt(newi int) *ActualParam {
	return &ActualParam{i: &newi, ParamType: ParamTypeInt}
}

func NewUnnamedParamBool(newb bool) *ActualParam {
	return &ActualParam{b: &newb, ParamType: ParamTypeBool}
}

func NewUnnamedParamDuration(newd int64) *ActualParam {
	return &ActualParam{d: &newd, ParamType: ParamTypeDuration}
}

func (ap *ActualParam) GetString() (string, error) {
	if ap.ParamType != ParamTypeString {
		return "", fmt.Errorf("cannot get string value from param of type %s", ap.ParamType)
	}
	return *ap.s, nil
}

func (ap *ActualParam) GetInt() (int, error) {
	if ap.ParamType != ParamTypeInt {
		return 0, fmt.Errorf("cannot get int value from param of type %s", ap.ParamType)
	}
	return *ap.i, nil
}

func (ap *ActualParam) GetBool() (bool, error) {
	if ap.ParamType != ParamTypeBool {
		return false, fmt.Errorf("cannot get bool value from param of type %s", ap.ParamType)
	}
	return *ap.b, nil
}

func (ap *ActualParam) GetDuration() (int64, error) {
	if ap.ParamType != ParamTypeDuration {
		return 0, fmt.Errorf("cannot get duration value from param of type %s", ap.ParamType)
	}
	return *ap.d, nil
}

func (ap *ActualParam) String() string {
	switch ap.ParamType {
	case ParamTypeString:
		return fmt.Sprintf("String(%s)", *ap.s)
	case ParamTypeInt:
		return fmt.Sprintf("Int(%d)", *ap.i)
	case ParamTypeBool:
		return fmt.Sprintf("Bool(%t)", *ap.b)
	case ParamTypeDuration:
		return fmt.Sprintf("Duration(%d)", *ap.d)
	default:
		return fmt.Sprintf("<unknown param type %s>", ap.ParamType)
	}
}
