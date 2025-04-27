package get

import (
	"github.com/harluo/config/internal/core/internal"
	"github.com/harluo/config/internal/core/internal/core"
	"github.com/harluo/di"
)

type Default struct {
	di.Get

	Paths       *core.Paths
	Environment *core.Environment
	Getter      *core.Getter
	Loader      *core.Loader

	Filler *internal.Filler
}
