package comp

type ParamTypeable[T any] interface {
	TypeName() string
	GetValue() T
	SetValue(T)
}

type ParamString struct {
	s string
}

type ParamInt struct {
	i int
}

type ParamBool struct {
	b bool
}

type ParamDuration struct {
	d int64
}

func NewParamString(news string) ParamTypeable[string] {
	return &ParamString{s: news}
}

func NewParamInt(newi int) ParamTypeable[int] {
	return &ParamInt{i: newi}
}

func NewParamBool(newb bool) ParamTypeable[bool] {
	return &ParamBool{b: newb}
}

func NewParamDuration(newd int64) ParamTypeable[int64] {
	return &ParamDuration{d: newd}
}

func (ps *ParamString) TypeName() string {
	return "string"
}

func (ps *ParamString) GetValue() string {
	return ps.s
}

func (ps *ParamString) SetValue(s string) {
	ps.s = s
}

func (pi *ParamInt) TypeName() string {
	return "int"
}

func (pi *ParamInt) GetValue() int {
	return pi.i
}

func (pi *ParamInt) SetValue(i int) {
	pi.i = i
}

func (pb *ParamBool) TypeName() string {
	return "bool"
}

func (pb *ParamBool) GetValue() bool {
	return pb.b
}

func (pb *ParamBool) SetValue(b bool) {
	pb.b = b
}

func (pd *ParamDuration) TypeName() string {
	return "duration"
}

func (pd *ParamDuration) GetValue() int64 {
	return pd.d
}

func (pd *ParamDuration) SetValue(d int64) {
	pd.d = d
}
