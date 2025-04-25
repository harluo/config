package core

import (
	"github.com/goexl/gox"
	"github.com/harluo/config/internal/internal/config"
	"github.com/harluo/config/internal/kernel"
)

type Builder struct {
	original *config.Getter
	config   *config.Getter

	override bool
}

func newBuilder(original *config.Getter) (builder *Builder) {
	builder = new(Builder)
	builder.original = original
	*builder.config = *original // 复制原始配置

	builder.override = false

	return
}

func (b *Builder) Nullable() (builder *Builder) {
	b.config.Nullable = true
	builder = b

	return
}

func (b *Builder) Required() (builder *Builder) {
	b.config.Nullable = false
	builder = b

	return
}

func (b *Builder) Default() (builder *Builder) {
	b.config.Default = true
	builder = b

	return
}

func (b *Builder) Filepath(required string, optionals ...string) (builder *Builder) {
	b.config.Paths = append(b.config.Paths, required)
	b.config.Paths = append(b.config.Paths, optionals...)
	builder = b

	return
}

func (b *Builder) Loader(required kernel.Loader, optionals ...kernel.Loader) (builder *Builder) {
	b.config.Loaders = append(b.config.Loaders, required)
	b.config.Loaders = append(b.config.Loaders, optionals...)
	builder = b

	return
}

func (b *Builder) Getter(required kernel.Getter, optionals ...kernel.Getter) (builder *Builder) {
	b.config.Getters[required] = new(gox.Empty)
	for _, _getter := range optionals {
		b.config.Getters[_getter] = new(gox.Empty)
	}
	builder = b

	return
}

func (b *Builder) Override() (builder *Builder) {
	b.override = true
	builder = b

	return
}

func (b *Builder) Build() (getter *Getter) {
	if b.override {
		*b.original = *b.config
	} else {
		getter = new(Getter)
		getter.init(b.config)
	}

	return
}
