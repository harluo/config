package kernel

import (
	"github.com/harluo/config/internal/runtime"
)

type Getter interface {
	Get(config runtime.Pointer) error
}
