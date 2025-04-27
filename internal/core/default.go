package core

import (
	"github.com/goexl/gfx"
	"github.com/goexl/mengpo"
	"github.com/goexl/xiren"
	"github.com/harluo/config/internal/core/internal"
	"github.com/harluo/config/internal/core/internal/core"
	"github.com/harluo/config/internal/core/internal/get"
	"github.com/harluo/config/internal/kernel"
	"github.com/harluo/config/internal/runtime"
)

type Default struct {
	path  string
	paths *core.Paths

	environment *core.Environment
	getter      *core.Getter
	loader      *core.Loader

	filler *internal.Filler
}

func newDefault(get get.Default) kernel.Getter {
	return &Default{
		paths: get.Paths,

		environment: get.Environment,
		getter:      get.Getter,
		loader:      get.Loader,
		filler:      get.Filler,
	}
}

func (g *Default) Get(target runtime.Pointer) (err error) {
	if dfe := g.detectPath(); nil != dfe { // 探测所有的配置文件路径
		err = dfe
	} else if fe := g.fill(target); nil != fe { // 加载数据
		err = fe
	}

	return
}

func (g *Default) fill(target runtime.Pointer) (err error) {
	if le := g.filler.Fill(target); nil != le { // 从路径中加载数据
		err = le
	}

	if nil == err { // 处理默认值
		// !此处逻辑不能往前，原因是如果对象里面包含指针，那么只能在包含指针的结构体被解析后才能去设置默认值，不然指针将被会设置成空值
		err = mengpo.New().Getter(g.getter).Build().Set(target)
	}

	if nil == err { // 从环境变量中加载配置
		err = g.environment.Process(target)
	}

	if nil == err { // 数据验证
		err = xiren.Struct(target)
	}

	return
}

func (g *Default) detectPath() (err error) {
	list := gfx.List().Filepath(g.path) // 加入默认从命令行和配置项而来的配置文件
	list.Limit().File().Build()         // 限制只探测文件
	list.Filename("*")                  // 探测所有可能的文件
	// 配置所有可能的配置目录
	list.Directory("config")
	list.Directory("conf")
	list.Directory("configuration")
	list.Directory(".") // 当前目录

	// Unix类操作系统目录
	list.Directory("conf.d")
	list.Directory("config.d")
	list.Directory("configuration.d")
	list.Directory("etc")

	g.loader.Foreach(func(loader kernel.Loader) { // 限制扩展名
		extensions := loader.Extensions()
		if 0 != len(extensions) {
			list.Extension(extensions[0], extensions[1:]...).Reset() // 当有加载器时，需要去掉默认参数
		}
	})

	if paths := list.Build().All(); 0 != len(paths) {
		g.paths.Add(paths[0], paths[1:]...)
		// TODO g.getLogger().Debug("使用配置文件进行配置加载", field.New("paths", *g.paths))
	}

	return
}
