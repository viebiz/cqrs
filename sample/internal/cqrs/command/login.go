package command

import (
	"github.com/viebiz/cqrs/application/base"
	"github.com/viebiz/cqrs/application/model"
)

type LoginCommand struct {
	base.ICommand[LoginCommandRequest]

	Request LoginCommandRequest
}

func (c LoginCommand) GetRequest() LoginCommandRequest {
	return c.Request
}

type LoginCommandRequest struct {
	model.IRequest
}

type LoginCommandResponse struct {
	model.IResponse
}
