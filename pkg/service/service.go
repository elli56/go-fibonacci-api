package service

type Calculation interface {
	FiboSlice(x, y int64) (map[int64]int64, error)
}

type Service struct {
	Calculation
}

func NewService() *Service {
	return &Service{
		Calculation: NewCalculateService(),
	}
}
