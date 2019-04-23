package db

import "github.com/go-redis/redis"

func (root *DatabaseRoot) GetRedis() (*redis.Client, error) {

	client := root.RedisClient

	_, err := client.Ping().Result()
	if err != nil {

		err = root.initRedisClient()
		if err != nil {
			return nil, err
		}

	}

	return client, nil
}
