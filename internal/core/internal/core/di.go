package core

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newPath,
		newPaths,
		newGetter,
		newLoader,
		newEnvironment,
	).Build().Apply()
}
