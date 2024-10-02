package services

import (
	"context"
	"errors"
	"httpServer/config"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type PingService struct {
	mutex             sync.Mutex
	counter           map[string]int
	RateLimitDuration time.Duration
	MaxRequests       int
}

func NewPingService() *PingService {
	return &PingService{
		counter:           make(map[string]int),
		RateLimitDuration: time.Second * 60,
		MaxRequests:       2,
	}
}

var ctx = context.Background()

func (s *PingService) HandlePing(userID string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// handle rate limit
	key := "rate_limit_api_ping_" + userID

	currentRequests, err := config.RedisRateLimitClient.Incr(ctx, key).Result()
	if err != nil {
		return err
	}

	if currentRequests == 1 {
		err = config.RedisRateLimitClient.Expire(ctx, key, s.RateLimitDuration).Err()
		if err != nil {
			return err
		}
	}

	if currentRequests > int64(s.MaxRequests) {
		return errors.New("rate limit exceeded: you can only call this API 2 times per 60 seconds")
	}

	// handle increment counter
	s.counter[userID]++

	// count number of pings
	err = config.RedisRateLimitClient.ZIncrBy(ctx, "ping_call_counts", 1, userID).Err()
	if err != nil {
		return err
	}

	// count number of users call api
	err = config.RedisRateLimitClient.PFAdd(ctx, "ping_user_count", userID).Err()
	if err != nil {
		return err
	}

	time.Sleep(time.Second * 5)

	return nil
}

func (s *PingService) GetPingCount(userID string) int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.counter[userID]
}

func (s *PingService) GetTopUsers() ([]redis.Z, error) {
	topUsers, err := config.RedisRateLimitClient.ZRevRangeWithScores(ctx, "ping_call_counts", 0, 9).Result()
	if err != nil {
		return nil, err
	}

	return topUsers, nil
}

func (s *PingService) GetApproximateUserCount() (int64, error) {
	count, err := config.RedisRateLimitClient.PFCount(ctx, "ping_user_count").Result()
	if err != nil {
		return 0, err
	}
	return count, nil
}
