package put

import (
	"github.com/harluo/boot"
	"github.com/harluo/di"
)

type Arguments struct {
	di.Put

	Config boot.Argument `group:"settings"`
}
