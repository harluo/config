package internal

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Get().Dependency().Puts(
		newFiller,
	).Build().Apply()
}
