package base

import "github.com/viebiz/cqrs/application/model"

type ICommand[TRequest model.IRequest] interface {
	GetCommand() model.Command
	GetRequest() TRequest
}
