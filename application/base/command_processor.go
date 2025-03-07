package base

import "github.com/viebiz/cqrs/application/model"

type ICommandProcessor[TRequest model.IRequest, TResponse model.IResponse] interface {
	Processing(request TRequest) (TResponse, error)
}
