package runtime

import (
	"fmt"
	"time"
)

type ParamTypeName int32

const (
	ParamTypeString ParamTypeName = iota
	ParamTypeInt
	ParamTypeBool
	ParamTypeDuration
	ParamTypeTypeName
)

func (pt ParamTypeName) String() string {
	switch pt {
	case ParamTypeString:
		return "String"
	case ParamTypeInt:
		return "Int"
	case ParamTypeBool:
		return "Bool"
	case ParamTypeDuration:
		return "Duration"
	case ParamTypeTypeName:
		return "TypeName"
	default:
		return fmt.Sprintf("<unknown param type %d>", pt)
	}
}

type FormalParam struct {
	Name      string
	Required  bool
	ParamType ParamTypeName
	Default   any
}

func (fp *FormalParam) String() string {
	if fp.Required {
		if fp.Default == nil {
			return fmt.Sprintf("(%s %s Required)", fp.ParamType, fp.Name)
		} else {
			return fmt.Sprintf("(%s %s Required Default %s)", fp.ParamType, fp.Name, fp.Default)
		}
	} else {
		if fp.Default == nil {
			return fmt.Sprintf("(%s %s Optional)", fp.ParamType, fp.Name)
		} else {
			return fmt.Sprintf("(%s %s Optional Default %s)", fp.ParamType, fp.Name, fp.Default)
		}
	}
}

func (fp *FormalParam) CoercibleFromString() bool {
	switch fp.ParamType {
	case ParamTypeDuration:
		return true
	default:
		return false
	}
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

func (builder *FormalParamsBuilder) WithOptionalDuration(name string, defaultValue time.Duration) *FormalParamsBuilder {
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

func (builder *FormalParamsBuilder) WithOptionalTypeName(name string, defaultValue ParamTypeName) *FormalParamsBuilder {
	builder.fps[name] = &FormalParam{
		Name:      name,
		Required:  false,
		ParamType: ParamTypeTypeName,
		Default:   NewUnnamedParamTypeName(defaultValue),
	}
	return builder
}

type ActualParams struct {
	Params map[string]*ActualParam
	Def    *ChkDef
}

func NewActualParams(cd *ChkDef, params map[string]*ActualParam) *ActualParams {
	return &ActualParams{Params: params, Def: cd}
}

// a convenience method so you don't have to reach inside ap.Params
func (ap *ActualParams) Get(key string) (*ActualParam, bool) {
	v, ok := ap.Params[key]
	return v, ok
}

func (ap *ActualParams) GetStringOrDefault(name string) (string, error) {
	v, ok := ap.Params[name]
	if ok {
		return v.GetString()
	}

	// param name not in actual params, try getting the default
	fp, ok := ap.Def.FormalParams[name]
	if !ok {
		msg := fmt.Errorf("Unknown parameter %s", name)
		return "", msg
	}
	return fp.Default.(string), nil
}

func (ap *ActualParams) GetIntOrDefault(name string) (int, error) {
	v, ok := ap.Params[name]
	if ok {
		return v.GetInt()
	}

	// param name not in actual params, try getting the default
	fp, ok := ap.Def.FormalParams[name]
	if !ok {
		msg := fmt.Errorf("Unknown parameter %s", name)
		return 0, msg
	}
	return fp.Default.(int), nil
}

func (ap *ActualParams) GetDurationOrDefault(name string) (time.Duration, error) {
	v, ok := ap.Params[name]
	if ok {
		return v.GetDuration()
	}

	// param name not in actual params, try getting the default
	fmt.Printf("FORMAL PARAMS = %+v\n", ap.Def.FormalParams)
	fp, ok := ap.Def.FormalParams[name]
	if !ok {
		msg := fmt.Errorf("Unknown parameter %s", name)
		return 0, msg
	}
	return fp.Default.(time.Duration), nil
}

type ActualParam struct {
	Name      string
	ParamType ParamTypeName

	s       *string
	i       *int
	b       *bool
	d       *time.Duration
	tn      *ParamTypeName
	Address ContentAddress
}

func (ap *ActualParam) String() string {
	switch ap.ParamType {
	case ParamTypeString:
		return fmt.Sprintf("(%s: String %s)", ap.Name, *ap.s)
	case ParamTypeInt:
		return fmt.Sprintf("(%s: Int %d)", ap.Name, *ap.i)
	case ParamTypeBool:
		return fmt.Sprintf("(%s: Bool %t)", ap.Name, *ap.b)
	case ParamTypeDuration:
		return fmt.Sprintf("(%s: Duration %d)", ap.Name, *ap.d)
	case ParamTypeTypeName:
		return fmt.Sprintf("(%s: TypeName %s)", ap.Name, *ap.tn)
	default:
		return fmt.Sprintf("(Unknown %s)", ap.ParamType)
	}
}

type TestParams map[string]any

func MakeTestParams(cd *ChkDef, p TestParams) *ActualParams {
	builder := BuildParams()
	for key, value := range p {
		switch v := value.(type) {
		case string:
			builder.NewNamedParamString(key, v)
		case int:
			builder.NewNamedParamInt(key, v)
		case bool:
			builder.NewNamedParamBool(key, v)
		case time.Duration:
			builder.NewNamedParamDuration(key, v)
		}
	}
	return builder.Build(cd)
}
func (ap *ActualParam) CoercibleToString() bool {
	switch ap.ParamType {
	case ParamTypeDuration:
		return true
	default:
		return false
	}
}

// Useful for testing
type ActualParamsBuilder struct {
	aps map[string]*ActualParam
}

func BuildParams() *ActualParamsBuilder {
	return &ActualParamsBuilder{
		aps: make(map[string]*ActualParam),
	}
}

func (builder *ActualParamsBuilder) Build(cd *ChkDef) *ActualParams {
	return NewActualParams(cd, builder.aps)
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
func (builder *ActualParamsBuilder) NewNamedParamDuration(name string, newD time.Duration) *ActualParamsBuilder {
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

func NewUnnamedParamDuration(newD time.Duration) *ActualParam {
	return &ActualParam{d: &newD, ParamType: ParamTypeDuration}
}

func NewUnnamedParamTypeName(newTN ParamTypeName) *ActualParam {
	return &ActualParam{tn: &newTN, ParamType: ParamTypeTypeName}
}

func (ap *ActualParam) GetString() (string, error) {
	if ap.ParamType != ParamTypeString {
		return "", fmt.Errorf("cannot get string value from param of type %s", ap.ParamType)
	}
	return *ap.s, nil
}

func (ap *ActualParam) GetTypeName() (ParamTypeName, error) {
	if ap.ParamType != ParamTypeTypeName {
		return 0, fmt.Errorf("cannot get typename value from param of type %s", ap.ParamType)
	}
	return *ap.tn, nil
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

func (ap *ActualParam) GetDuration() (time.Duration, error) {
	if ap.ParamType != ParamTypeDuration {
		return 0, fmt.Errorf("cannot get duration value from param of type %s", ap.ParamType)
	}
	return *ap.d, nil
}

func TypeNameToType(s string) (ParamTypeName, error) {
	switch s {
	case "String":
		return ParamTypeString, nil
	case "Int":
		return ParamTypeInt, nil
	case "Bool":
		return ParamTypeBool, nil
	case "Duration":
		return ParamTypeDuration, nil
	case "TypeName":
		return ParamTypeTypeName, nil
	default:
		return ParamTypeString, fmt.Errorf("unknown type name: %s", s)
	}
}
