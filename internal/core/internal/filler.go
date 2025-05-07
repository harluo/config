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

func (f *Filler) Fill(target runtime.Pointer) (err error) {
	for _, path := range f.paths.Get() {
		err = f.load(path, target)
	}

	return
}

func (f *Filler) load(path string, target runtime.Pointer) (err error) {
	if ctx, has, lce := f.loadLocalContext(path); nil != lce {
		err = lce
	} else {
		err = f.fill(ctx, target, has)
	}

	return
}

func (f *Filler) Wrote() {
	for _, target := range f.targets {
		newTarget := reflect.New(reflect.TypeOf(target)).Elem().Interface()
		if le := f.Fill(newTarget); nil != le {
			// todo
		} else if !reflect.DeepEqual(target, newTarget) { // 如果配置有变化
			// todo
		}
	}
}
func (f *Filler) loadLocalContext(path string) (ctx context.Context, populated bool, err error) {
	if bytes, rfe := f.read(path); nil != rfe {
		err = rfe
	} else if eval, ee := envsubst.Eval(string(bytes), f.getter.Get); nil != ee {
		err = ee
	} else {
		ctx = context.Background()
		ctx = context.WithValue(ctx, kernel.ContextFilepath, path)
		ctx = context.WithValue(ctx, kernel.ContextBytes, []byte(eval))
		populated = 0 != len(bytes)
	}

	return
}

func (f *Filler) read(path string) (bytes []byte, err error) {
	if "" != path {
		bytes, err = os.ReadFile(path)
	}

	return
}

func (f *Filler) fill(localContext context.Context, target runtime.Pointer, populated bool) (err error) {
	networkContext := context.Background()
	modules := f.modules(target)
	for _, loader := range f.loader.Get() {
		if loader.Local() && !populated {
			continue
		}

		ctx := localContext
		if !loader.Local() { // 默认为本地上下文，如果为网络加载器，切换为网络上下文
			ctx = networkContext
		}

		newToLoad := reflect.New(reflect.TypeOf(target).Elem()).Elem().Interface() // 创建一份新的值，用于从配置文件中加载
		if loaded, le := loader.Load(ctx, &newToLoad, modules); nil != le {
			err = le
		} else if loaded { // 确实加载了配置数据
			err = internal.NewDecoder(newToLoad).Decode(target)
		}

		if nil != err {
			break
		}
	}

	if nil == err {
		f.targets = append(f.targets, target)
	}

	return
}

func (f *Filler) modules(target runtime.Pointer) (modules []string) {
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
