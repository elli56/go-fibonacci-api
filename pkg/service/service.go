package service

import (
	"github.com/elli56/fibo-api/pkg/repository"
)

type Calculation interface {
	FiboSlice(x, y int64) (map[int64]int64, error)
}

type Service struct {
	Calculation
}

func NewService(redisDb *repository.Repository) *Service {
	return &Service{
		Calculation: NewCalculateService(redisDb.PostCache),
	}
}
