package internal

import (
	"context"
	"os"
	"reflect"

	"github.com/drone/envsubst"
	"github.com/goexl/gox"
	"github.com/harluo/config/internal/core/internal/core"
	"github.com/harluo/config/internal/core/internal/internal"
	"github.com/harluo/config/internal/core/internal/internal/get"
	"github.com/harluo/config/internal/kernel"
	"github.com/harluo/config/internal/runtime"
)

type Filler struct {
	paths  *core.Paths
	loader *core.Loader
	getter *core.Getter

	targets []runtime.Pointer
}

func newFiller(get get.Filler) *Filler {
	return &Filler{
		paths:  get.Paths,
		loader: get.Loader,
		getter: get.Getter,

		targets: make([]runtime.Pointer, 0),
	}
}

func (l *Filler) Load(target runtime.Pointer) (err error) {
	for _, path := range l.paths.Get() {
		err = l.load(path, target)
	}

	return
}

func (l *Filler) load(path string, target runtime.Pointer) (err error) {
	if ctx, has, lce := l.loadLocalContext(path); nil != lce {
		err = lce
	} else {
		err = l.fill(ctx, target, has)
	}

	return
}

func (l *Filler) Wrote() {
	for _, target := range l.targets {
		newTarget := reflect.New(reflect.TypeOf(target)).Elem().Interface()
		if le := l.Load(newTarget); nil != le {
			// todo
		} else if !reflect.DeepEqual(target, newTarget) { // 如果配置有变化
			// todo
		}
	}
}
func (l *Filler) loadLocalContext(path string) (ctx context.Context, populated bool, err error) {
	if bytes, rfe := l.read(path); nil != rfe {
		err = rfe
	} else if eval, ee := envsubst.Eval(string(bytes), l.getter.Get); nil != ee {
		err = ee
	} else {
		ctx = context.Background()
		ctx = context.WithValue(ctx, kernel.ContextFilepath, path)
		ctx = context.WithValue(ctx, kernel.ContextBytes, []byte(eval))
		populated = 0 != len(bytes)
	}

	return
}

func (l *Filler) read(path string) (bytes []byte, err error) {
	if "" != path {
		bytes, err = os.ReadFile(path)
	}

	return
}

func (l *Filler) fill(localContext context.Context, target runtime.Pointer, populated bool) (err error) {
	networkContext := context.Background()
	modules := l.modules(target)
	for _, loader := range l.loader.Get() {
		if loader.Local() && !populated {
			continue
		}

		ctx := localContext
		if !loader.Local() { // 默认为本地上下文，如果为网络加载器，切换为网络上下文
			ctx = networkContext
		}

		value := make(map[string]any)
		if loaded, le := loader.Load(ctx, &value, modules); nil != le {
			err = le
		} else if loaded && 0 != len(value) { // 确实加载了配置数据
			err = internal.NewDecoder(&value).Decode(target)
		}

		if nil != err {
			break
		}
	}

	if nil == err {
		l.targets = append(l.targets, target)
	}

	return
}

func (l *Filler) modules(target runtime.Pointer) (modules []string) {
	typeOfTarget := reflect.TypeOf(target)
	if reflect.Ptr == typeOfTarget.Kind() {
		typeOfTarget = typeOfTarget.Elem()
	}
	if 1 != typeOfTarget.NumField() {
		return
	}

	names := make(map[string]*gox.Empty)
	name := typeOfTarget.Field(0).Name

	swh := gox.String(name).Switch()
	names[swh.Strike().Build().Case()] = nil
	names[swh.Underscore().Build().Case()] = nil
	names[swh.Camel().Build().Case()] = nil

	modules = make([]string, 0, len(names))
	for key := range names {
		modules = append(modules, key)
	}

	return
}
