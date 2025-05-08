package internal

import (
	"github.com/harluo/config/internal/kernel"
	"github.com/harluo/di"
)

type Put struct {
	di.Put

	Environment kernel.Finder `group:"mappers"`
}
