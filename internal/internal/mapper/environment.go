package mapper

import (
	"github.com/goexl/env"
)

type Environment struct {
	// 无字段
}

func newEnvironment() *Environment {
	return &Environment{
		// 无字段
	}
}

func (*Environment) Find(key string) string {
	return env.Get(key)
}
