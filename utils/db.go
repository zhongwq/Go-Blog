package utils

import (
	"github.com/boltdb/bolt"
)

type DB struct {
	conn *bolt.DB
}

func (db *DB) getConn() *bolt.DB {
	conn, err := bolt.Open("data.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	db.conn = conn
	return db.conn
}

func (db *DB) Get(bucket string, key string) []byte {
	conn := db.getConn()
	defer conn.Close()
	var value []byte
	err := conn.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			panic(err)
		}
		val := b.Get([]byte(key))
		value = make([]byte, len(val))
		copy(value, val)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return value
}

func (db *DB) Set(bucket string, key string, value string) error {
	conn := db.getConn()
	defer conn.Close()
	var ret error
	err := conn.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			panic(err)
		}
		ret = b.Put([]byte(key), []byte(value))
		return nil
	})
	if err != nil {
		panic(err)
	}
	return ret
}

func (db *DB) GenerateID(bucket string) int {
	conn := db.getConn()
	defer conn.Close()
	var id int
	err := conn.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			panic(err)
		}
		ret, _ := b.NextSequence()
		id = int(ret)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return id
}

func (db *DB) Scan(bucket string) map[string]string {
	conn := db.getConn()
	defer conn.Close()
	var value = make(map[string]string)
	err := conn.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			panic(err)
		}
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tempk := string(k)
			tempv := string(v)
			value[tempk] = tempv
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return value
}
