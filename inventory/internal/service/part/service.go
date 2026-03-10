package part

import (
	"golearning/inventory/internal/repository"
	servicepkg "golearning/inventory/internal/service"
)

type Service struct {
	repository repository.PartRepository
}

func NewService(repository repository.PartRepository) servicepkg.PartService {
	return &Service{repository: repository}
}
