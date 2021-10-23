package service

import (
	"errors"
	"math"

	"github.com/elli56/fibo-api/pkg/repository"
	"github.com/sirupsen/logrus"
)

type CalculateService struct {
	redisDb repository.PostCache
}

func NewCalculateService(redisdb repository.PostCache) *CalculateService {
	return &CalculateService{redisDb: redisdb}
}

func (s *CalculateService) FiboSlice(x, y int64) (map[int64]int64, error) {
	var m = make(map[int64]int64)
	if x >= y {
		return m, errors.New("incorrect diapasone. 'x' should be less than 'y'")
	}
	for x < y {
		// проверяем есть ли значение в Redis
		redisValue, err := s.redisDb.Get(x)
		// если не нашли то считаем и пишем в map и в Redis
		if err != nil {
			logrus.Error(err)
			n := s.fibonacciFindByNumber(float64(x))
			m[x] = int64(n)
			s.redisDb.Set(x, int64(n))
			x++
		} else {
			// если нашли то пишем в map
			m[x] = redisValue
			x++
		}
	}
	return m, nil
}

// formula
func (s *CalculateService) fibonacciFindByNumber(n float64) float64 {
	return (math.Pow(float64((1.0+math.Sqrt(5.0))/2), n) - math.Pow(float64((1.0-math.Sqrt(5.0))/2), n)) / math.Sqrt(5.0)
}
