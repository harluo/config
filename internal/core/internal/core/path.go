package core

type Path struct {
	path string
}

func newPath() *Path {
	return &Path{
		path: "",
	}
}

func (p *Path) Get() string {
	return p.path
}

func (p *Path) Bind() *string {
	return &p.path
}
