package get

import (
	"github.com/harluo/config/internal/core/internal/core"
	"github.com/harluo/di"
)

type Filler struct {
	di.Get

	Paths  *core.Paths
	Loader *core.Loader
	Getter *core.Getter
}
