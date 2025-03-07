package readonly

import (
	"github.com/viebiz/cqrs/application/bus"
)

// New returns a new registered bus.QueryBus
func New() bus.QueryBus {
	return QueryBus
}

var (
	QueryBus = bus.QueryBus{}
)
