package loader

import (
	"context"
	"encoding/xml"
	"path/filepath"
	"strings"

	"github.com/goexl/exception"
	"github.com/goexl/gox/field"
	"github.com/harluo/config/internal/internal/loader/internal/constant"
	"github.com/harluo/config/internal/kernel"
	"github.com/harluo/config/internal/runtime"
)

var _ kernel.Loader = (*Xml)(nil)

type Xml struct {
	targets map[runtime.Pointer]bool
}

func newXml() *Xml {
	return &Xml{
		targets: make(map[runtime.Pointer]bool),
	}
}

func (*Xml) Local() bool {
	return true
}

func (*Xml) Extensions() []string {
	return []string{
		constant.ExtensionXml,
	}
}

func (x *Xml) Load(ctx context.Context, target *map[string]any, _ []string) (loaded bool, err error) {
	if path, pok := ctx.Value(kernel.ContextFilepath).(string); !pok {
		err = exception.New().Message("未指定配置文件路径").Field(field.New("loader", "xml")).Build()
	} else if bytes, bok := ctx.Value(kernel.ContextBytes).([]byte); !bok {
		err = exception.New().Message("配置文件无内容").Field(field.New("loader", "xml")).Build()
	} else {
		loaded, err = x.load(&path, &bytes, target)
	}

	return
}

func (x *Xml) load(path *string, bytes *[]byte, target *map[string]any) (loaded bool, err error) {
	loadable := false
	if constant.ExtensionXml == strings.ToLower(filepath.Ext(*path)) {
		loadable = true
		err = xml.Unmarshal(*bytes, target)
	}
	if nil == err && loadable {
		loaded = true
	}

	return
}
