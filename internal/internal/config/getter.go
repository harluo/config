package config

import (
	"github.com/goexl/env"
	"github.com/goexl/gox"
	"github.com/harluo/config/internal/internal/config/internal"
	"github.com/harluo/config/internal/internal/config/internal/getter"
	"github.com/harluo/config/internal/internal/loader"
	"github.com/harluo/config/internal/kernel"
)

type Getter struct {
	// 是否允许设置默认值
	Default bool
	// 是否要验证配置数据
	Validate bool
	// 是否可以没有配置文件
	Nullable bool
	// 是否可刷新配置
	Refreshable bool
	// 配置文件列表
	Paths []string

	// 标签
	Tag *internal.Tag
	// 环境变量获取器
	Getters map[kernel.Getter]*gox.Empty

	Loaders  []kernel.Loader
	Changers []kernel.Changer
}

func NewGetter() *Getter {
	return &Getter{
		Default:     true,
		Validate:    true,
		Nullable:    true,
		Refreshable: true,
		Paths:       make([]string, 0), // 默认没有配置文件

		Tag: internal.NewTag(),
		Getters: map[kernel.Getter]*gox.Empty{
			getter.NewDefault(env.Get): new(gox.Empty),
		},

		Loaders: []kernel.Loader{
			loader.NewJson(),
			loader.NewXml(),
		},
		Changers: make([]kernel.Changer, 0),
	}
}

func (g *Getter) Get(key string) (value string) {
	for _getter := range g.Getters {
		value = _getter.Get(key)
		if "" != value { // 及时回退
			break
		}
	}

	return
}
