package processor

import (
	"github.com/viebiz/cqrs/application/base"
	"github.com/viebiz/cqrs/sample/internal/cqrs/command"
)

type LoginCommandProcessor[TRequest command.LoginCommandRequest, TResponse command.LoginCommandResponse] struct {
	base.ICommandProcessor[TRequest, TResponse]
}

func (p LoginCommandProcessor[TRequest, TResponse]) Processing(request TRequest) (TResponse, error) {
	return TResponse{}, nil
}
