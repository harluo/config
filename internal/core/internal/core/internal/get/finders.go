package get

import (
	"sort"

	"github.com/harluo/config/internal/core/internal/core/internal/core"
	"github.com/harluo/config/internal/kernel"
	"github.com/harluo/di"
)

type Finders struct {
	di.Get

	Finders []kernel.Finder `group:"finders"`
}

func (f *Finders) Sorted() []kernel.Finder {
	sort.Slice(f.Finders, func(i, j int) bool {
		it := core.NewTyper(f.Finders[i])
		jt := core.NewTyper(f.Finders[j])

		return it.Order() < jt.Order()
	})

	return f.Finders
}
