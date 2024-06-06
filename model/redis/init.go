package redis

import (
	"github.com/15733012783/weikang/nacos"
	"github.com/go-redis/redis/v8"
)

func InitRedis(callback func(db *redis.Client) (interface{}, error)) (interface{}, error) {
	nac := nacos.RpcNac.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     nac.Addr,
		Password: nac.Password,
	})
	i, err := callback(rdb)
	if err != nil {
		return nil, err
	}
	defer rdb.Close()
	return i, err
}
