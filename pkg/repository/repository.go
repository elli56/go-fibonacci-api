package repository

import "time"

type PostCache interface {
	Set(key, value int64)
	Get(key int64) (int64, error)
}

type Repository struct {
	PostCache
}

// Конструктор просто создает экземпляр структуры RedisCache которая содержит методы и подключается к Redis сама. Передавать ничего не нужно
func NewRepository(Host, Port, Password string, Db int, Exp time.Duration) *Repository {
	return &Repository{
		PostCache: NewRedisCache(Host, Port, Password, Db, Exp),
	}
}
