package common

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisMode int

const (
	RedisModeOfSignal RedisMode = iota
	RedisModeOfCluster
)

const (
	Offline = iota
	Online
)

var Client *RedisClient

type RedisClient struct {
	Client interface{}
	Type   RedisMode
	Ctx    context.Context
}

type RedisConfig struct {
	Address  []string
	Password string
	Type     RedisMode
}

func InitRedis(conf RedisConfig) error {
	if Client == nil {
		Client = new(RedisClient)
		Client.Ctx = context.Background()
	}
	if conf.Type == RedisModeOfSignal {
		opt := &redis.Options{
			Addr:     conf.Address[0],
			Password: conf.Password,
		}
		Client.Client = redis.NewClient(opt)
		Client.Type = RedisModeOfSignal
		_, err := Client.Client.(*redis.Client).Ping(Client.Ctx).Result()
		if err != nil {
			return err
		}
	} else {
		opt := &redis.ClusterOptions{
			Addrs:    conf.Address,
			Password: conf.Password,
		}
		Client.Client = redis.NewClusterClient(opt)
		Client.Type = RedisModeOfCluster
		_, err := Client.Client.(*redis.ClusterClient).Ping(Client.Ctx).Result()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RedisClient) Set(key, value string) error {
	if r.Type == RedisModeOfSignal {
		err := r.Client.(*redis.Client).Set(Client.Ctx, key, value, 0).Err()
		if err != nil {
			return err
		}
	} else {
		err := r.Client.(*redis.ClusterClient).Set(Client.Ctx, key, value, 0).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RedisClient) GetSet(key string) (string, error) {
	if r.Type == RedisModeOfSignal {
		value, err := r.Client.(*redis.Client).Get(Client.Ctx, key).Result()
		return value, err
	} else {
		value, err := r.Client.(*redis.ClusterClient).Get(Client.Ctx, key).Result()
		return value, err
	}
}
