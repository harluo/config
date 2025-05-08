package core

import (
	_ "github.com/harluo/config/internal/internal/finder" // 预加载内部映射器
	"github.com/harluo/config/internal/kernel"

	"github.com/harluo/config/internal/core/internal/core/internal/get"
)

type Getter struct {
	finders []kernel.Finder
}

func newGetter(get get.Finders) *Getter {
	return &Getter{
		finders: get.Sorted(),
	}
}

func (g *Getter) Get(key string) (value string) {
	for _, finder := range g.finders {
		value = finder.Find(key)
		if "" != value { // 及时回退
			break
		}
	}

	return
}
