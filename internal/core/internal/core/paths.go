package core

type Paths struct {
	paths []string
}

func newPaths() *Paths {
	return &Paths{
		paths: make([]string, 0),
	}
}

func (p *Paths) Get() []string {
	return p.paths
}

func (p *Paths) Add(required string, optionals ...string) {
	p.paths = append(p.paths, required)
	p.paths = append(p.paths, optionals...)
}
