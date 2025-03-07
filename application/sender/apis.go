package sender

import (
	"github.com/viebiz/cqrs/application/base"
	"github.com/viebiz/cqrs/application/model"
)

//func SendCommand[TRequest model.IRequest, TResponse model.IResponse](processor base.ICommandProcessor[TRequest, TResponse], request TRequest) (TResponse, error) {
//	return processor.Processing(request)
//}

func SendQuery[TRequest model.IRequest, TResponse model.IResponse](executorFn interface{}, query base.IQuery[TRequest]) (TResponse, error) {
	executor := executorFn.(base.IQueryExecutor[TRequest, TResponse])
	return executor.Executing(query.GetRequest())
}
