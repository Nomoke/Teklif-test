package offer

import (
	"context"
	"log"
	"time"
)

// запускает воркер для автоматического снятия просроченных предложений
func (s *Service) StartOfferExpiryWorker(ctx context.Context) {
	ticker := time.NewTicker(8 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("Запуск авто-снятия просроченных предложений...")
			if err := s.offerRepo.ExpireOldOffers(ctx); err != nil {
				log.Printf("Ошибка при авто-снятии предложений: %v", err)
			}
		case <-ctx.Done():
			log.Println("Остановка воркера авто-снятия предложений")
			return
		}
	}
}
