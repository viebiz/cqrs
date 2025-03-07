package bus

import (
	"github.com/viebiz/cqrs/application/model"
)

type CommandBus map[model.Command]interface{}
