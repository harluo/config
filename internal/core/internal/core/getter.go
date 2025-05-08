package core

import (
	_ "github.com/harluo/config/internal/internal/mapper" // 预加载内部映射器

	"github.com/harluo/config/internal/core/internal/core/internal/get"
	"github.com/harluo/di"
)

type Getter struct{}

func newGetter() *Getter {
	return new(Getter)
}

func (g *Getter) Get(key string) (value string) {
	di.New().Instance().Get(func(get get.Finders) {
		for _, finder := range get.Finders {
			value = finder.Find(key)
			if "" != value { // 及时回退
				break
			}
		}
	}).Build().Apply()

	return
}
