package redis

import (
	"github.com/go-redis/redis/v8"
	"weikang_database/nacos"
)

func InitRedis(callback func(db *redis.Client) (interface{}, error)) (interface{}, error) {
	nac := nacos.RpcNac.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: nac.Addr,
	})
	i, err := callback(rdb)
	if err != nil {
		return nil, err
	}
	defer rdb.Close()
	return i, err
}
