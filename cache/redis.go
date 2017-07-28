package cache

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"
	"github.com/garyburd/redigo/redis"
	"syscall"
)

type RedisPool struct {
	Dial func() (Conn, error)

	MaxIdle int

	MaxActive int

	IdelTimeout time.Duration

	Wait bool
}

var (
	Pool *redis.Pool
	redisServer = flag.String("redisServer", ":6379", "")
	redisPassword = flag.String("redisPassword", "123456", "")
)

func init() {
	redisHost := ":6379"
	Pool = newPool(redisHost)
	close()
}

func newPool(server, password string) *redis.Pool {
	return *Redis.Pool{
		MaxIdle: 3,
		MaxActive: 5,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server) 
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		}

	}
}

func close() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<- c
		Pool.Close()
		os.Exit(0)
	}()
}

func main() {
}


func SET(key string, value string, expire int) {
	conn := Pool.Get()
	defer conn.Close()
	
	res, err := redis.Bytes(conn.Do("SET", key, value, expire))
	if err != nil {
		return res, fmt.Errorf("error set key %s : %v", key, err)
	}
	return res, err
}

func GET(key string) {
    conn := Pool.Get()
	defer conn.Close()

	var data[] byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error get key %s: %v", key, err)
	}	
	return data, err
}

func HSET(hash string, key string, value string) {
	conn := Pool.Get()
	defer conn.Close()
	
	res, err := redis.Bytes(conn.DO("HSET", hash, key, value))
	if err != nil {
		return res, "HSET error"
	}
	return res, err
}

func DEL(key string) {
	conn := Pool.Get()
	defer conn.Close()

	res, err := redis.Bytes(conn.Do("DEL", key))
	if err != nil {
		return res, "DEL key error"
	}
	return res, err
}

//用户