package argument

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newConfig,
	).Build().Apply()
}
