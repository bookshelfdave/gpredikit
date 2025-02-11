package comp

import "fmt"

const (
	ParamTypeString   = "String"
	ParamTypeInt      = "Int"
	ParamTypeBool     = "Bool"
	ParamTypeDuration = "Duration"
	ParamTypeTypeName = "TypeName"
)

type ActualParam struct {
	Name      string
	ParamType string

	s       *string
	i       *int
	b       *bool
	d       *int64
	Address ContentAddress
}

func NewUnnamedParamString(newS string) *ActualParam {
	return &ActualParam{s: &newS, ParamType: ParamTypeString}
}

func NewUnnamedParamInt(newI int) *ActualParam {
	return &ActualParam{i: &newI, ParamType: ParamTypeInt}
}

func NewUnnamedParamBool(newB bool) *ActualParam {
	return &ActualParam{b: &newB, ParamType: ParamTypeBool}
}

func NewUnnamedParamDuration(newD int64) *ActualParam {
	return &ActualParam{d: &newD, ParamType: ParamTypeDuration}
}

func NewUnnamedParamTypeName(newS string) *ActualParam {
	return &ActualParam{s: &newS, ParamType: ParamTypeTypeName}
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
		return fmt.Sprintf("%s: String(%s) Address(%v)", ap.Name, *ap.s, ap.Address)
	case ParamTypeInt:
		return fmt.Sprintf("%s: Int(%d) Address(%v)", ap.Name, *ap.i, ap.Address)
	case ParamTypeBool:
		return fmt.Sprintf("%s: Bool(%t) Address(%v)", ap.Name, *ap.b, ap.Address)
	case ParamTypeDuration:
		return fmt.Sprintf("%s: Duration(%d) Address(%v)", ap.Name, *ap.d, ap.Address)
	case ParamTypeTypeName:
		return fmt.Sprintf("%s: TypeName(%s) Address(%v)", ap.Name, *ap.s, ap.Address)
	default:
		return fmt.Sprintf("%s: <unknown param type %s> Address(%v)", ap.Name, ap.ParamType, ap.Address)
	}
}
