package offer

type Service struct {
	offerRepo offerRepo
}

func NewService(r offerRepo) *Service {
	return &Service{
		offerRepo: r,
	}
}
