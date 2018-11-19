package cache

import (
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
