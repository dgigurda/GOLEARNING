package payment

import (
	"github.com/google/uuid"

	"golearning/payment/internal/repository"
	servicepkg "golearning/payment/internal/service"
)

type Service struct {
	repository repository.PaymentRepository
	newUUID    func() string
}

func NewService(repository repository.PaymentRepository) servicepkg.PaymentService {
	return &Service{repository: repository, newUUID: uuid.NewString}
}
