package cache

import (
	"github.com/guitarpawat/portscan/api/model"
	"go.etcd.io/bbolt"
	"time"
)

const bucketName = "portscan"
const timeout = 1 * time.Second

var db *bolt.DB

func getDB() (*bolt.DB, error) {
	if db == nil {
		conn, err := bolt.Open("portscan.db", 0600, &bolt.Options{Timeout: timeout})
		if err != nil {
			return nil, err
		}
		db = conn
	}
	return db, nil
}

func PutNewToken(id string, ip ...string) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}

		val := model.MakeGetOutput(ip...)

		b, err := val.Marshal()
		if err != nil {
			return err
		}

		return bucket.Put([]byte(id), b)
	})
}