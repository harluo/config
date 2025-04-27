package loader

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Get().Dependency().Puts(
		newJson,
		newXml,
	).Build().Apply()
}
