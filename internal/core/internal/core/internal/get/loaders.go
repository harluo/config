package get

import (
	"github.com/harluo/config/internal/kernel"
	"github.com/harluo/di"
)

type Loaders struct {
	di.Get

	Loaders []kernel.Loader `group:"loaders"`
}
