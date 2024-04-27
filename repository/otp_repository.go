package repository

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type OtpRepository interface {
	CreateOtp() string
	StoreOtp(key string, value interface{}, expTime time.Duration) error
	FetchOtp(key string) (string, error)
}

type otpRepository struct {
	redisClient *redis.Client
}

func (o *otpRepository) CreateOtp() string {
	rand.Seed(time.Now().UnixNano())
	low := 100000
	hi := 999999
	otpInt := low + rand.Intn(hi-low)
	return strconv.Itoa(otpInt)
}

func (o *otpRepository) StoreOtp(key string, value interface{}, expTime time.Duration) error {
	result := o.redisClient.Set(context.Background(), key, value, expTime)
	return result.Err()
}

func (o *otpRepository) FetchOtp(key string) (string, error) {
	return o.redisClient.Get(context.Background(), key).Result()
}

func NewOtpRepository(redisClient *redis.Client) OtpRepository {
	repo := new(otpRepository)
	repo.redisClient = redisClient
	return repo
}
