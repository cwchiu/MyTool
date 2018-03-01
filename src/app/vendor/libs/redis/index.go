package redis

import (
	// "errors"
	db "gopkg.in/redis.v5"
	"time"
)

type openCalback func(*db.Client) error

func openDB(addr string, callback openCalback) error {
	client := db.NewClient(&db.Options{
		Addr:         addr,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return err
	}

	defer client.Close()

	return callback(client)

}

func Put(addr, key, value string) error {
	return openDB(addr, func(client *db.Client) error {
		return client.Set(key, value, 0).Err()
	})
}

func Get(addr, key string) (string, error) {
	var result string
	err := openDB(addr, func(client *db.Client) error {
		result = client.Get(key).Val()
		return nil
	})
	if err != nil {
		return "", err
	}
	return result, nil
}
