package core

import (
	"github.com/harluo/boot"
	"github.com/harluo/config/internal/core/internal/core"
	"github.com/harluo/config/internal/core/internal/put"
)

func newArguments(path *core.Path) put.Arguments {
	return put.Arguments{
		Config: boot.NewArgument("config", path.Bind()).
			Usage("指定配置文件`路径`").
			Aliases("c", "conf", "configuration").
			Build(),
	}
}
