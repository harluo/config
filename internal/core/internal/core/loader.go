package core

import (
	_ "github.com/harluo/config/internal/internal/loader" // 预加载内部加载器

	"github.com/harluo/config/internal/core/internal/core/internal/get"
	"github.com/harluo/config/internal/kernel"
	"github.com/harluo/di"
)

type Loader struct {
	// 无字段
}

func newLoader() *Loader {
	return &Loader{
		// 无字段
	}
}

func (*Loader) Get() (loaders []kernel.Loader) {
	di.New().Get().Dependency().Get(func(get get.Loaders) {
		loaders = get.Loaders
	}).Build().Build().Apply()

	return
}

func (*Loader) Foreach(each func(loader kernel.Loader)) {
	di.New().Get().Dependency().Get(func(get get.Loaders) {
		for _, loader := range get.Loaders {
			each(loader)
		}
	}).Build().Build().Apply()
}
