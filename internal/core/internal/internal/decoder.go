package internal

import (
	"strings"

	"github.com/goexl/gox"
	"github.com/goexl/structer"
	"github.com/harluo/config/internal/runtime"
)

type Decoder struct {
	from runtime.Pointer
}

func NewDecoder(from runtime.Pointer) *Decoder {
	return &Decoder{
		from: from,
	}
}

func (d *Decoder) Decode(target runtime.Pointer) (err error) {
	builder := structer.Copy().Mapper(d.mapper)

	environments := make(map[string]any) // !通过中间亦是的引入，防止零值被复制到了目标变量
	if fee := builder.From(d.from).To(&environments).Build().Apply(); nil != fee {
		err = fee
	} else if ete := builder.From(environments).To(target).Build().Apply(); nil != ete {
		err = ete
	}

	return
}

func (d *Decoder) mapper(key string, field string) (mapped bool) {
	if key == field {
		mapped = true
	} else {
		mapped = d.variants(key, field)
	}

	return
}

func (d *Decoder) variants(key string, field string) (mapped bool) {
	from := gox.String(field).Switch()
	checked := strings.ToLower(key) // 全部转换成小写，避免复杂的判断逻辑
	if strings.ToLower(from.Camel().Build().Case()) == checked {
		mapped = true
	} else if strings.ToLower(from.Underscore().Build().Case()) == checked {
		mapped = true
	} else if strings.ToLower(from.Strike().Build().Case()) == checked {
		mapped = true
	}

	return
}
