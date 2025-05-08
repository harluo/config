package finder

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newEnvironment,
	).Group("finders").Build().Apply()
}
