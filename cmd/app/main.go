package main

import (
	"context"
	"log"

	"teklif/internal/api"
	"teklif/internal/config"
	"teklif/internal/pkg/db"
	loggerpkg "teklif/internal/pkg/logger"
	repo "teklif/internal/repository/offer"
	offersrv "teklif/internal/service/offer"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, err := loggerpkg.New(cfg.Logger.LogFilePath)
	if err != nil {
		log.Fatal("Ошибка загрузки лог-файла:", err)
	}

	pool, err := db.OpenDB(ctx, cfg.Database)
	if err != nil {
		logger.Fatal(err)
	}
	defer pool.Close()

	offerService := offersrv.NewService(repo.New(pool))

	// Запуск воркера для автоматического снятия просроченных предложений
	go offerService.StartOfferExpiryWorker(ctx)

	// Запуск API
	r, err := api.New(logger, offerService)
	if err != nil {
		log.Fatal(err)
	}

	logger.Info("Сервер успешно запущен на порту 8080")
	logger.Error(r.Run("0.0.0.0:8080"))
}
