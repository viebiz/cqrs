package auth

import (
	"net/http"

	"github.com/viebiz/cqrs/application/sender"
	"github.com/viebiz/cqrs/sample/internal/cqrs/business/writable"
	"github.com/viebiz/cqrs/sample/internal/cqrs/command"
	"github.com/viebiz/lit"
)

type impl struct {
	sender sender.Sender
}

func New(router lit.Router) {
	ctrl := &impl{}
	router.Get("/login", ctrl.Login)
}

func (i *impl) Login(ctx lit.Context) error {
	response, err := sender.SendCommand[
		command.LoginCommandRequest,
		command.LoginCommandResponse,
	](
		command.LoginCommand{
			Request: command.LoginCommandRequest{},
		},
	)
	if err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, response)
	return nil
}
