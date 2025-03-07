package base

import "github.com/viebiz/cqrs/application/model"

type IQueryExecutor[TRequest model.IRequest, TResponse model.IResponse] interface {
	Executing(request TRequest) (TResponse, error)
}
