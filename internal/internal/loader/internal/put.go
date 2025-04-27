package internal

import (
	"github.com/harluo/config/internal/kernel"
	"github.com/harluo/di"
)

type Put struct {
	di.Put

	Json kernel.Loader `group:"loaders"`
	Xml  kernel.Loader `group:"loaders"`
}
