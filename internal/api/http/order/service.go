package order

import (
	"applicationDesignTest/internal/entity/order"
	"context"
	"net/http"
)

type svc interface {
	CreateOrder(ctx context.Context, order order.Order) error
}

type logger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
}

type HttpService struct {
	orderSvc svc
	logger   logger
}

func NewHttpService(orderSvc svc, logger logger) HttpService {
	return HttpService{
		orderSvc: orderSvc,
		logger:   logger,
	}
}

func (s *HttpService) writeError(w http.ResponseWriter, err error, status int) {
	http.Error(w, err.Error(), status)
	s.logger.Error(err.Error())
}
