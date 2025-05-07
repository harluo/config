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
	middle := make(map[string]any) // !通过中间亦是的引入，防止零值被复制到了目标变量
	if tme := structer.Copy().From(d.from).To(&middle).Mapper(d.mapper).Build().Apply(); nil != tme {
		err = tme
	} else if fme := structer.Copy().From(middle).To(target).Mapper(d.mapper).Build().Apply(); nil != fme {
		err = fme
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
