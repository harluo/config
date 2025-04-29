package mapper

import (
	"github.com/harluo/config/internal/internal/mapper/internal"
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newEnvironment,
		func(environment *Environment) internal.Put {
			return internal.Put{
				Environment: environment,
			}
		},
	).Build().Apply()
}
