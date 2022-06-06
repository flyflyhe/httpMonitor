package httpMonitor

import (
	"errors"
	"github.com/boltdb/bolt"
	"os"
	"sync"
)

const BucketProxy = "bucketProxy"
const BucketUrl = "bucketUrl"

var (
	once sync.Once
	db   *bolt.DB
	err  error
)

func init() {
	once.Do(func() {
		db, err = bolt.Open("./monitor.db", os.ModePerm, nil)
		if err != nil {
			panic(err)
		}
	})
}

//proxy eg socks5:127.0.0.1:9888

func SetUrlProxy(proxy string) error {
	return GetDb().Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(BucketProxy))
		if err != nil {
			return err
		}

		return bucket.Put([]byte(proxy), []byte(proxy))
	})
}

func SetUrl(url string) error {
	return GetDb().Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(BucketUrl))
		if err != nil {
			return err
		}

		return bucket.Put([]byte(url), []byte(url))
	})
}

func GetByBucket(bucketName string) (arr []string, err error) {
	err = GetDb().View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))

		if bucket == nil {
			return nil
		}

		err := bucket.ForEach(func(k, v []byte) error {
			arr = append(arr, string(v))
			return nil
		})

		return err
	})

	return
}

func GetByBucketAndKey(bucketName, key string) (v string, err error) {
	err = GetDb().View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))

		if bucket == nil {
			return nil
		}

		result := bucket.Get([]byte(key))

		if result != nil {
			v = string(result)
		}

		return nil
	})

	return
}

func Delete(bucketName, key string) error {
	return GetDb().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))

		if bucket == nil {
			return nil
		}

		return bucket.Delete([]byte(key))
	})
}

func GetDb() *bolt.DB {
	if db == nil {
		panic(errors.New("db 必须预先初始化"))
	}

	return db
}
