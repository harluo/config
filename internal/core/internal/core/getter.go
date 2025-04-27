package core

import (
	"github.com/harluo/config/internal/core/internal/core/internal/get"
	"github.com/harluo/di"
)

type Getter struct{}

func newGetter() *Getter {
	return new(Getter)
}

func (g *Getter) Get(key string) (value string) {
	di.New().Get().Dependency().Get(func(get get.Mappers) {
		for _, mapper := range get.Mappers {
			value = mapper.Get(key)
			if "" != value { // 及时回退
				break
			}
		}
	}).Build().Build().Apply()

	return
}
