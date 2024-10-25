package common

import (
	"context"
	"domain_blog/infrastructure/database/entity"
	"encoding/json"
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

func (r *RedisClient) GetAbouts(key string) ([]entity.About, error) {
	value, err := r.GetSet(key)
	if err != nil {
		return nil, err
	}
	abouts := make([]entity.About, 0)
	bytes := []byte(value)
	err = json.Unmarshal(bytes, &abouts)
	if err != nil {
		return nil, err
	}
	return abouts, nil
}

func (r *RedisClient) SetAbouts(key string, abouts []entity.About) error {
	bytes, err := json.Marshal(&abouts)
	if err != nil {
		return err
	}
	err = r.Set(key, string(bytes))
	if err != nil {
		return err
	}
	return nil
}
