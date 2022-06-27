package httpMonitor

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/boltdb/bolt"
	"github.com/rs/zerolog/log"
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
		path, err := os.Getwd()
		if err != nil {
			log.Error().Msg(err.Error())
		}
		log.Debug().Str("store.path", path).Send()
		db, err = bolt.Open(path+string(os.PathSeparator)+"monitor.db", os.ModePerm, nil)
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

func SetUrl(url string, interval int32) error {
	return GetDb().Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(BucketUrl))
		if err != nil {
			return err
		}

		return bucket.Put([]byte(url), IntToBytes(interval))
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

func GetAllUrls() (map[string]int32, error) {
	m := make(map[string]int32)
	err = GetDb().View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketUrl))

		if bucket == nil {
			return nil
		}

		err := bucket.ForEach(func(k, v []byte) error {
			m[string(k)] = BytesToInt(v)
			return nil
		})

		return err
	})

	return m, err
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

//整形转换成字节

func IntToBytes(n int32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}

//字节转换成整形

func BytesToInt(b []byte) int32 {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	_ = binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}
