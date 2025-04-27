package get

import (
	"github.com/harluo/config/internal/kernel"
	"github.com/harluo/di"
)

type Mappers struct {
	di.Get

	Mappers []kernel.Mapper `group:"mappers"`
}
