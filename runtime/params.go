package runtime

import "fmt"

const (
	ParamTypeString   = "String"
	ParamTypeInt      = "Int"
	ParamTypeBool     = "Bool"
	ParamTypeDuration = "Duration"
	ParamTypeTypeName = "TypeName"
)

type FormalParam struct {
	Name      string
	Required  bool
	ParamType string
	Default   *ActualParam
}

type FormalParamsBuilder struct {
	fps map[string]*FormalParam
}

// I might like Build_ better than New_ here, don't punch me
func BuildFormalParams() *FormalParamsBuilder {
	return &FormalParamsBuilder{
		fps: make(map[string]*FormalParam),
	}
}

func (builder *FormalParamsBuilder) Build() map[string]*FormalParam {
	return builder.fps
}

func (builder *FormalParamsBuilder) WithRequiredString(name string) *FormalParamsBuilder {
	builder.fps[name] = &FormalParam{
		Name:      name,
		Required:  true,
		ParamType: ParamTypeString,
		Default:   nil,
	}
	return builder
}

func (builder *FormalParamsBuilder) WithOptionalString(name string, defaultValue string) *FormalParamsBuilder {
	builder.fps[name] = &FormalParam{
		Name:      name,
		Required:  false,
		ParamType: ParamTypeString,
		Default:   NewUnnamedParamString(defaultValue),
	}
	return builder
}

func (builder *FormalParamsBuilder) WithRequiredInt(name string) *FormalParamsBuilder {
	builder.fps[name] = &FormalParam{
		Name:      name,
		Required:  true,
		ParamType: ParamTypeInt,
		Default:   nil,
	}
	return builder
}

func (builder *FormalParamsBuilder) WithOptionalInt(name string, defaultValue int) *FormalParamsBuilder {
	builder.fps[name] = &FormalParam{
		Name:      name,
		Required:  false,
		ParamType: ParamTypeInt,
		Default:   NewUnnamedParamInt(defaultValue),
	}
	return builder
}

func (builder *FormalParamsBuilder) WithRequiredBool(name string) *FormalParamsBuilder {
	builder.fps[name] = &FormalParam{
		Name:      name,
		Required:  true,
		ParamType: ParamTypeBool,
		Default:   nil,
	}
	return builder
}

func (builder *FormalParamsBuilder) WithOptionalBool(name string, defaultValue bool) *FormalParamsBuilder {
	builder.fps[name] = &FormalParam{
		Name:      name,
		Required:  false,
		ParamType: ParamTypeBool,
		Default:   NewUnnamedParamBool(defaultValue),
	}
	return builder
}

func (builder *FormalParamsBuilder) WithRequiredDuration(name string) *FormalParamsBuilder {
	builder.fps[name] = &FormalParam{
		Name:      name,
		Required:  true,
		ParamType: ParamTypeDuration,
		Default:   nil,
	}
	return builder
}

func (builder *FormalParamsBuilder) WithOptionalDuration(name string, defaultValue int64) *FormalParamsBuilder {
	builder.fps[name] = &FormalParam{
		Name:      name,
		Required:  false,
		ParamType: ParamTypeDuration,
		Default:   NewUnnamedParamDuration(defaultValue),
	}
	return builder
}

func (builder *FormalParamsBuilder) WithRequiredTypeName(name string) *FormalParamsBuilder {
	builder.fps[name] = &FormalParam{
		Name:      name,
		Required:  true,
		ParamType: ParamTypeTypeName,
		Default:   nil,
	}
	return builder
}

func (builder *FormalParamsBuilder) WithOptionalTypeName(name string, defaultValue string) *FormalParamsBuilder {
	builder.fps[name] = &FormalParam{
		Name:      name,
		Required:  false,
		ParamType: ParamTypeTypeName,
		Default:   NewUnnamedParamTypeName(defaultValue),
	}
	return builder
}

type ActualParam struct {
	Name      string
	ParamType string

	s       *string
	i       *int
	b       *bool
	d       *int64
	Address ContentAddress
}

// this is probably going to be used mostly for testing
type ActualParamsBuilder struct {
	aps map[string]*ActualParam
}

func BuildActualParams() *ActualParamsBuilder {
	return &ActualParamsBuilder{
		aps: make(map[string]*ActualParam),
	}
}

func (builder *ActualParamsBuilder) Build() map[string]*ActualParam {
	return builder.aps
}

func (builder *ActualParamsBuilder) NewNamedParamString(name string, newS string) *ActualParamsBuilder {
	builder.aps[name] = &ActualParam{Name: name, s: &newS, ParamType: ParamTypeString}
	return builder
}
func (builder *ActualParamsBuilder) NewNamedParamInt(name string, newI int) *ActualParamsBuilder {
	builder.aps[name] = &ActualParam{Name: name, i: &newI, ParamType: ParamTypeInt}
	return builder
}
func (builder *ActualParamsBuilder) NewNamedParamBool(name string, newB bool) *ActualParamsBuilder {
	builder.aps[name] = &ActualParam{Name: name, b: &newB, ParamType: ParamTypeBool}
	return builder
}
func (builder *ActualParamsBuilder) NewNamedParamDuration(name string, newD int64) *ActualParamsBuilder {
	builder.aps[name] = &ActualParam{Name: name, d: &newD, ParamType: ParamTypeDuration}
	return builder
}
func (builder *ActualParamsBuilder) NewNamedParamTypeName(name string, newS string) *ActualParamsBuilder {
	builder.aps[name] = &ActualParam{Name: name, s: &newS, ParamType: ParamTypeTypeName}
	return builder
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
