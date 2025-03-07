package writable

import (
	"github.com/viebiz/cqrs/application/bus"
	"github.com/viebiz/cqrs/sample/internal/cqrs/command"
	"github.com/viebiz/cqrs/sample/internal/cqrs/processor"
)

// New returns a new registered bus.CommandBus
func New() bus.CommandBus {
	return CommandBus
}

var (
	CommandBus = bus.CommandBus{
		LoginCommand: processor.LoginCommandProcessor[command.LoginCommandRequest, command.LoginCommandResponse]{},
	}
)
