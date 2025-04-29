package internal

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newFiller,
	).Build().Apply()
}
