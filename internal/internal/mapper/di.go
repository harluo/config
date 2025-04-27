package mapper

import (
	"github.com/harluo/config/internal/internal/mapper/internal"
	"github.com/harluo/di"
)

func init() {
	di.New().Get().Dependency().Puts(
		newEnvironment,
		func(environment *Environment) internal.Put {
			return internal.Put{
				Environment: environment,
			}
		},
	).Build().Apply()
}
