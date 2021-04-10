package redis

import "github.com/garyburd/redigo/redis"

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxActive: 15, //TODO
		MaxIdle:   5,  //TODO
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				"localhost:6379",           //TODO
				redis.DialPassword("carton&*()"), //TODO
			)
		},
	}
	conn := pool.Get()
	defer conn.Close()
	if _, err := conn.Do("ping"); err != nil {
		panic(err)
	}
}
