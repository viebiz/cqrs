package sender

import (
	"github.com/viebiz/cqrs/application/base"
	"github.com/viebiz/cqrs/application/model"
)

type SendCommand[TRequest model.IRequest, TResponse model.IResponse] func(processor base.ICommandProcessor[TRequest, TResponse], request TRequest) (TResponse, error)
