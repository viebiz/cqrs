package sender

import (
	"github.com/viebiz/cqrs/application/bus"
)

type impl struct {
	queryBus   bus.QueryBus
	commandBus bus.CommandBus

	sendCommand func[TRequest any, TResponse any](request TRequest) (TResponse, error)
}

func New(
	queryBus bus.QueryBus,
	commandBus bus.CommandBus,
) Sender {
	return &impl{
		queryBus:   queryBus,
		commandBus: commandBus,
	}
}

//func (i *impl[TRequest, TResponse]) SendCommand(command base.ICommand[TRequest]) (TResponse, error) {
//	processorFn := i.commandBus[command.GetCommand()]
//	processor := processorFn.(base.ICommandProcessor[TRequest, TResponse])
//	return SendCommand[TRequest, TResponse](processor, command.GetRequest())
//}
