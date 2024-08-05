package main

import (
	httpOrderService "applicationDesignTest/internal/api/http/order"
	orderRepository "applicationDesignTest/internal/repository/order"
	roomRepository "applicationDesignTest/internal/repository/room"
	orderService "applicationDesignTest/internal/service/order"
	"applicationDesignTest/internal/tx"
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log := slog.Default()
	orderRepo := orderRepository.NewRepository()
	roomRepo := roomRepository.NewRepository()
	tm := tx.NewInMemoryTxManager()

	orderSvc := orderService.NewService(tm, orderRepo, roomRepo)
	httpOrderSvc := httpOrderService.NewHttpService(orderSvc, log)

	router := chi.NewRouter()
	router.Post("/orders", httpOrderSvc.CreateOrder)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		<-ctx.Done()
		if err := srv.Shutdown(ctx); err != nil {
			log.Error("Server shutdown err:", err)
			return
		}
		log.Info("Server gracefully stopped")
	}()

	log.Info("Server listening on localhost:8080")
	err := srv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		log.Info("Server closed")
	} else if err != nil {
		log.Error("Server failed: ", err)
		os.Exit(1)
	}
}
