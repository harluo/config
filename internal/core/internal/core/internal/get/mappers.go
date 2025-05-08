package get

import (
	"github.com/harluo/config/internal/kernel"
	"github.com/harluo/di"
)

type Finders struct {
	di.Get

	Finders []kernel.Finder `group:"finders"`
}
