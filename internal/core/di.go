package core

import (
	"github.com/harluo/config/internal/core/internal/argument"
	"github.com/harluo/config/internal/core/internal/put"
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newDefault,
		func(config *argument.Config) put.Arguments {
			return put.Arguments{
				Config: config,
			}
		},
	).Build().Apply()
}
