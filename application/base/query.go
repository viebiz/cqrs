package base

import "github.com/viebiz/cqrs/application/model"

type IQuery[TRequest model.IRequest] interface {
	GetRequest() TRequest
}
