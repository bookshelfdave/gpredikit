package runtime

type ChkRegistry struct {
	fns    map[string]*ChkDef
	groups map[string]*ChkDef
	tools  map[string]*ChkDef
}

func NewChkRegistry() *ChkRegistry {
	return &ChkRegistry{
		fns:    make(map[string]*ChkDef),
		groups: make(map[string]*ChkDef),
		tools:  make(map[string]*ChkDef),
	}
}

func (c *ChkRegistry) IsValidFn(id string) bool {
	_, ok := c.fns[id]
	return ok
}

func (c *ChkRegistry) GetFn(id string) *ChkDef {
	return c.fns[id]
}

func (c *ChkRegistry) RegisterFn(cd *ChkDef) {
	if err := cd.Validate(); err != nil {
		// any errors here should be raised during development
		panic(err)
	}

	// keep track of groups separately
	if cd.IsGroup {
		c.groups[cd.Name] = cd
	}

	// but always add it to Fns
	c.fns[cd.Name] = cd
}
