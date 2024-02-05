package storage

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type RedisStorage struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisStorage(client *redis.Client) *RedisStorage {
	return &RedisStorage{
		client: client,
		ctx:    context.Background(),
	}
}

// 实现 storage.Storage 接口的方法
func (s *RedisStorage) Init() error {
	// 初始化存储，如果需要的话
	return nil
}

func (s *RedisStorage) Visited(requestID uint64) error {
	// 记录已访问的URL
	return s.client.Set(s.ctx, s.getKey(requestID), "1", 0).Err()
}

func (s *RedisStorage) IsVisited(requestID uint64) (bool, error) {
	// 检查URL是否已访问
	exists, err := s.client.Exists(s.ctx, s.getKey(requestID)).Result()
	return exists > 0, err
}

func (s *RedisStorage) Clear() error {
	// 清除所有记录
	return s.client.FlushDB(s.ctx).Err()
}

func (s *RedisStorage) getKey(requestID uint64) string {
	// 生成存储键名
	return "visited:" + strconv.FormatUint(requestID, 10)
}
