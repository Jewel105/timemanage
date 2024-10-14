package factory

import (
	"context"
	"fmt"
	"gin_study/config"
	"gin_study/logger"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var redisClient *redis.Client

func RedisStart() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:         config.Config.Redis.Host + ":" + strconv.Itoa(config.Config.Redis.Port),
		Password:     config.Config.Redis.Password,
		DB:           config.Config.Redis.Db,
		ReadTimeout:  time.Duration(config.Config.Redis.IdleTimeout) * time.Second,
		WriteTimeout: time.Duration(config.Config.Redis.IdleTimeout) * time.Second,
	})
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		logger.Error(map[string]interface{}{"redis init error": err.Error})
	}
}

// RedisSet 设置
func RedisSet(ctx context.Context, key, value string, expiration time.Duration) error {
	return redisClient.Set(ctx, key, value, expiration).Err()
}

// RedisGet 读取
func RedisGet(ctx context.Context, key string) (string, error) {
	return redisClient.Get(ctx, key).Result()
}

// RedisDel 删除
func RedisDel(ctx context.Context, key string) error {
	return redisClient.Del(ctx, key).Err()
}

// RedisHSet 设置
func RedisHSet(ctx context.Context, key, field, value string) error {
	return redisClient.HSet(ctx, key, field, value).Err()
}

// RedisHGet 读取
func RedisHGet(ctx context.Context, key, field string) (string, error) {
	return redisClient.HGet(ctx, key, field).Result()
}

// RedisHGetAll 读取全部数据
func RedisHGetAll(ctx context.Context, key string) (map[string]string, error) {
	return redisClient.HGetAll(ctx, key).Result()
}

// RedisHDel 删除
func RedisHDel(ctx context.Context, key, field string) error {
	return redisClient.HDel(ctx, key, field).Err()
}

// RedisExpire 续期
func RedisExpire(ctx context.Context, key string, expiration time.Duration) error {
	result, err := redisClient.Expire(ctx, key, expiration).Result()
	if err != nil || !result {
		return fmt.Errorf("failed to set expiration: %w", err)
	}
	return nil
}
