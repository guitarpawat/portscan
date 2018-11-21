package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/guitarpawat/portscan/api/model"
	"go.etcd.io/bbolt"
	"os"
	"os/signal"
	"time"
)

const bucketName = "portscan"
const timeout = 1 * time.Second

var db *bolt.DB

func init() {
	_, err := getDB()
	if err != nil {
		panic(err)
	}

	// Close the db if OS interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	go func() {
		select {
		case <-c:
			CloseDB()
			os.Exit(2)
		}
	}()
}

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

// CloseDB closes the database connection
func CloseDB() error {
	fmt.Println("db closed")
	return db.Close()
}

// PutNewToken register new token to database
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

// GetTokenInfo gets the results of the token id from database
func GetTokenInfo(id string) (model.GetOutput, error) {
	db, err := getDB()
	if err != nil {
		return model.GetOutput{}, err
	}

	var out model.GetOutput
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return errors.New("no db bucket found")
		}
		b := bucket.Get([]byte(id))
		if b == nil {
			return errors.New("token id not found")
		}
		err := json.Unmarshal(b, &out)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return model.GetOutput{}, err
	}

	return out, nil
}

// UpdateTokenInfo updates the results for specified token id
func UpdateTokenInfo(id string, result model.Result) error {
	res, err := GetTokenInfo(id)
	if err != nil {
		return err
	}

	res.Results = append(res.Results, result)
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if err != nil {
			return err
		}


		b, err := res.Marshal()
		if err != nil {
			return err
		}

		return bucket.Put([]byte(id), b)
	})
}