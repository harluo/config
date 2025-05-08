package core

import (
	"math"

	"github.com/harluo/config/internal/core/internal/core/internal/checker"
)

type Typer struct {
	data any
}

func NewTyper(data any) *Typer {
	return &Typer{
		data: data,
	}
}

func (t *Typer) Order() (order uint64) {
	switch converted := t.data.(type) {
	case checker.Order[int]:
		order = uint64(converted.Order())
	case checker.Order[int8]:
		order = uint64(converted.Order())
	case checker.Order[int16]:
		order = uint64(converted.Order())
	case checker.Order[int32]:
		order = uint64(converted.Order())
	case checker.Order[int64]:
		order = uint64(converted.Order())
	case checker.Order[uint]:
		order = uint64(converted.Order())
	case checker.Order[uint8]:
		order = uint64(converted.Order())
	case checker.Order[uint16]:
		order = uint64(converted.Order())
	case checker.Order[uint32]:
		order = uint64(converted.Order())
	case checker.Order[uint64]:
		order = converted.Order()
	default:
		order = math.MaxUint64
	}

	return
}
