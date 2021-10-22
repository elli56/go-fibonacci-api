package service

import (
	"errors"
	"math"
)

type CalculateService struct {
}

func NewCalculateService() *CalculateService {
	return &CalculateService{}
}

func (s *CalculateService) FiboSlice(x, y int64) (map[int64]int64, error) {
	var m = make(map[int64]int64)
	if x >= y {
		return m, errors.New("incorrect diapasone. 'x' should be less than 'y'")
	}
	for x < y {
		n := s.fibonacciFindByNumber(float64(x))
		m[x] = int64(n)
		x++
	}
	return m, nil
}

// formula
func (s *CalculateService) fibonacciFindByNumber(n float64) float64 {
	return (math.Pow(float64((1.0+math.Sqrt(5.0))/2), n) - math.Pow(float64((1.0-math.Sqrt(5.0))/2), n)) / math.Sqrt(5.0)
}
