package loader

import (
	"github.com/harluo/config/internal/internal/loader/internal"
	"github.com/harluo/di"
)

func init() {
	di.New().Get().Dependency().Puts(
		newJson,
		newXml,
		func(json *Json, xml *Xml) internal.Put {
			return internal.Put{
				Json: json,
				Xml:  xml,
			}
		},
	).Build().Apply()
}
