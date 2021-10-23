package repository

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

type Config struct {
	Host     string
	Port     string
	Password string
	Db       int
	Exp      time.Duration
}

type RedisCache struct {
	host     string
	port     string
	password string
	db       int
	// время исчезновения для всех хранящихся элементов
	expires time.Duration
}

func NewRedisCache(Host, Port, Password string, Db int, Exp time.Duration) *RedisCache {
	return &RedisCache{
		host:     Host,
		port:     Port,
		password: Password,
		db:       Db,
		expires:  Exp,
	}
}

// create new redis client
func (cache *RedisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cache.host, cache.port),
		Password: cache.password,
		DB:       cache.db,
	})
}

func (cache *RedisCache) Set(key, value int64) {
	client := cache.getClient()

	key_int := int(key)
	client.Set(strconv.Itoa(key_int), value, cache.expires*time.Hour)
}

func (cache *RedisCache) Get(key int64) (int64, error) {
	client := cache.getClient()

	key_int := int(key)
	val, err := client.Get(strconv.Itoa(key_int)).Result()
	if err != nil {
		return 0, fmt.Errorf("no such key. Writing key=%d to Redis", key_int)
	}

	val_int, err := strconv.Atoi(val)
	if err != nil {
		return 0, fmt.Errorf("error while conversion value: %d from String to int", val_int)
	}

	return int64(val_int), nil
}
