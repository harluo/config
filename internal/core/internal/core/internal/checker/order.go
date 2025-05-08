package checker

import (
	"github.com/harluo/config/internal/core/internal/core/internal/internal/constraint"
)

type Order[T constraint.Order] interface {
	Order() T
}
