package bolt

import (
	"errors"
	db "github.com/boltdb/bolt"
)

type openCalback func(*db.DB) error
type ScanCallback func(bucket, key, value string) error

func openDB(filename string, callback openCalback) error {
	db, err := db.Open(filename, 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	return callback(db)

}

func Put(filename, bucket, key, value string) error {
	return openDB(filename, func(storage *db.DB) error {
		return storage.Update(func(tx *db.Tx) error {
			b, err := tx.CreateBucket([]byte(bucket))
			if err != nil {
				return err
			}

			return b.Put([]byte(key), []byte(value))
		})
	})
}

func Get(filename, bucket, key string) (string, error) {
	var result string
	err := openDB(filename, func(storage *db.DB) error {
		return storage.View(func(tx *db.Tx) error {
			b := tx.Bucket([]byte(bucket))
			bs := b.Get([]byte(key))
			if bs == nil {
				return errors.New("Key not Found")
			}
			result = string(bs)
			return nil
		})
	})
	if err != nil {
		return "", err
	}
	return result, nil
}

func Scan(filename, bucket string, callback ScanCallback) error {
	return openDB(filename, func(storage *db.DB) error {
		return storage.View(func(tx *db.Tx) error {
			b := tx.Bucket([]byte(bucket))
			c := b.Cursor()
			for k, v := c.First(); k != nil; k, v = c.Next() {

				callback(bucket, string(k), string(v))
			}

			return nil
		})
	})
}
