package finder

import (
	"github.com/goexl/env"
	"github.com/harluo/config/internal/kernel"
)

type Environment struct {
	// 无字段
}

func newEnvironment() kernel.Finder {
	return &Environment{
		// 无字段
	}
}

func (*Environment) Find(key string) string {
	return env.Get(key)
}
