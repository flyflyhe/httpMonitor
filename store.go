package httpMonitor

import (
	"errors"
	"github.com/boltdb/bolt"
	"os"
	"sync"
)

var (
	lock           sync.Mutex
	once           sync.Once
	db             *bolt.DB
	err            error
	bucketProxyPre string
)

func init() {
	bucketProxyPre = "bucketProxy"
	once.Do(func() {
		db, err = bolt.Open("./monitor.db", os.ModePerm, nil)
		if err != nil {
			panic(err)
		}
	})
}

//proxy eg socks5:127.0.0.1:9888

func SetUrlProxy(monitorURL string, proxy string) error {
	lock.Lock()
	defer lock.Unlock()
	return GetDb().Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketProxyPre + monitorURL))
		if err != nil {
			return err
		}

		return bucket.Put([]byte(proxy), []byte(proxy))
	})
}

func GetUrlProxyList(monitorURL string) (proxyArr []string, err error) {
	err = GetDb().View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketProxyPre + monitorURL))

		err := bucket.ForEach(func(k, v []byte) error {
			proxyArr = append(proxyArr, string(v))

			return nil
		})

		return err
	})

	return
}

func GetDb() *bolt.DB {
	if db == nil {
		panic(errors.New("db 必须预先初始化"))
	}

	return db
}
