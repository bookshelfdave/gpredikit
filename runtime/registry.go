package runtime

type ChkRegistry struct {
	Fns    map[string]*ChkDef
	Groups map[string]*ChkDef
	Tools  map[string]*ChkDef
}

func NewChkRegistry() *ChkRegistry {
	return &ChkRegistry{
		Fns:    make(map[string]*ChkDef),
		Groups: make(map[string]*ChkDef),
		Tools:  make(map[string]*ChkDef),
	}
}
