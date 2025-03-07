package bus

import (
	"github.com/viebiz/cqrs/application/model"
)

type QueryBus map[model.Query]interface{}
