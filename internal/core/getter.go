package core

import (
	"github.com/goexl/exception"
	"github.com/goexl/gfx"
	"github.com/goexl/gox"
	"github.com/goexl/log"
	"github.com/goexl/mengpo"
	"github.com/goexl/xiren"
	"github.com/harluo/config/internal/core/internal/core"
	"github.com/harluo/config/internal/internal/config"
	"github.com/harluo/config/internal/runtime"
	"github.com/harluo/di"
)

type Getter struct {
	path   string
	paths  *gox.Slice[string]
	config *config.Getter

	environment *core.Environment
	loader      *core.Loader
	watcher     *core.Watcher

	logger log.Logger
}

func newGetter() (getter *Getter) {
	getter = new(Getter)
	getter.init(config.NewGetter())

	return
}

func (g *Getter) Get(target runtime.Pointer) (err error) {
	if dfe := g.detectPath(); nil != dfe { // 探测所有的配置文件路径
		err = dfe
	} else if fe := g.fill(target); nil != fe { // 加载数据
		err = fe
	}

	return
}

func (g *Getter) Rebuild() *Builder {
	return newBuilder(g.config)
}

func (g *Getter) init(config *config.Getter) {
	g.paths = gox.NewPointer(gox.NewSlice[string](g.path))
	g.config = config

	g.environment = core.NewEnvironment()
	g.loader = core.NewLoader(g.paths, g.config)
	g.watcher = core.NewWatch(g.loader)

	return
}

func (g *Getter) fill(target runtime.Pointer) (err error) {
	if le := g.loader.Load(target); nil != le { // 从路径中加载数据
		err = le
	}

	if nil == err && g.config.Default { // 处理默认值
		// !此处逻辑不能往前，原因是如果对象里面包含指针，那么只能在包含指针的结构体被解析后才能去设置默认值，不然指针将被会设置成空值
		err = mengpo.New().Tag(g.config.Tag.Default).Getter(g.config).Build().Set(target)
	}

	// 从环境变量中加载配置
	if nil == err {
		err = g.environment.Process(target)
	}

	if nil == err && g.config.Validate { // 数据验证
		err = xiren.Struct(target)
	}

	return
}

func (g *Getter) detectPath() (err error) {
	list := gfx.List().Filepath(g.path, g.config.Paths...) // 加入默认从命令行和配置项而来的配置文件
	list.Limit().File().Build()                            // 限制只探测文件
	list.Filename("*")                                     // 探测所有可能的文件
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

	// 限制扩展名
	for _, loader := range g.config.Loaders {
		extensions := loader.Extensions()
		if 0 != len(extensions) {
			list.Extension(extensions[0], extensions[1:]...).Reset() // 当有加载器时，需要去掉默认参数
		}
	}

	if paths := list.Build().All(); 0 != len(paths) {
		*g.paths = gox.NewSlice(paths...) // !一定要使用指针修改原来的值，而不是传新的指针
		// TODO g.getLogger().Debug("使用配置文件进行配置加载", field.New("paths", *g.paths))
	} else if !g.config.Nullable {
		err = exception.New().Message("配置文件不存在").Build()
	}

	return
}

func (g *Getter) getLogger() log.Logger {
	if nil == g.logger {
		di.New().Get().Dependency().Get(g.setLogger).Build().Build().Apply()
	}

	return g.logger
}

func (g *Getter) setLogger(logger log.Logger) {
	g.logger = logger
}
