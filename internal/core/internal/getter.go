package internal

import (
	"github.com/harluo/config/internal/core/internal/core"
	"github.com/harluo/di"
)

type Getter struct {
	di.Get

	Paths       *core.Paths
	Environment *core.Environment
	Getter      *core.Getter
	Loader      *core.Loader

	Filler *Filler
}
