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

func RedisStart() error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:         config.Config.Redis.Host + ":" + strconv.Itoa(config.Config.Redis.Port),
		Password:     config.Config.Redis.Password,
		DB:           config.Config.Redis.Db,
		ReadTimeout:  time.Duration(config.Config.Redis.IdleTimeout) * time.Second,
		WriteTimeout: time.Duration(config.Config.Redis.IdleTimeout) * time.Second,
	})
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		logger.Error(map[string]interface{}{"redis init error": err.Error()})
		return err
	}
	return nil
}

// RedisSet 设置
func RedisSet(key, value string, expiration time.Duration) error {
	return redisClient.Set(ctx, key, value, expiration).Err()
}

// RedisGet 读取
func RedisGet(key string) (string, error) {
	return redisClient.Get(ctx, key).Result()
}

// RedisDel 删除
func RedisDel(key string) error {
	return redisClient.Del(ctx, key).Err()
}

// RedisHSet 设置
func RedisHSet(key, field, value string) error {
	return redisClient.HSet(ctx, key, field, value).Err()
}

// RedisHGet 读取
func RedisHGet(key, field string) (string, error) {
	return redisClient.HGet(ctx, key, field).Result()
}

// RedisHGetAll 读取全部数据
func RedisHGetAll(key string) (map[string]string, error) {
	return redisClient.HGetAll(ctx, key).Result()
}

// RedisHDel 删除
func RedisHDel(key, field string) error {
	return redisClient.HDel(ctx, key, field).Err()
}

// RedisExpire 续期
func RedisExpire(key string, expiration time.Duration) error {
	result, err := redisClient.Expire(ctx, key, expiration).Result()
	if err != nil || !result {
		return fmt.Errorf("failed to set expiration: %w", err)
	}
	return nil
}
