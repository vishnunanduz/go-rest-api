package cache

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/vishnunanduz/go-rest-api/entity"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) PostCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) getRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",       // no password set
		DB:       cache.db, // use default DB
	})
}

func (cache *redisCache) Set(key string, value *entity.Post) {
	rdc := cache.getRedisClient()

	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	rdc.Set(key, json, cache.expires*time.Second)

}
func (cache *redisCache) Get(key string) *entity.Post {
	rdc := cache.getRedisClient()

	val, err := rdc.Get(key).Result()
	if err != nil {
		return nil
	}

	post := entity.Post{}
	err = json.Unmarshal([]byte(val), &post)
	if err != nil {
		panic(err)
	}

	return &post

}
