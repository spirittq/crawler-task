package database

import (
	"encoding/json"
	"shared/utils"
	"sync"

	"github.com/rs/zerolog/log"

	"go.etcd.io/bbolt"
)

var Database *bbolt.DB
var DB_BUCKET_NAME = utils.GetEnvOrDefault("DB_BUCKET_NAME", "")
var DB_NAME = utils.GetEnvOrDefault("DB_NAME", "")
var mu = sync.Mutex{}

func InitDB() {
	var err error
	Database, err = bbolt.Open(DB_NAME, 0600, nil)
	if err != nil {
		log.Fatal().Msgf("could not open database: %v", err)
	}
}

func SaveToDB[T any](data T, key []byte) error {
	err := Database.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(DB_BUCKET_NAME))
		if err != nil {
			return err
		}

		item := bucket.Get(key)
		if item != nil {
			log.Warn().Msg("record already exists in database")
			return nil
		}

		json, err := json.Marshal(data)
		if err != nil {
			return err
		}
		err = bucket.Put(key, json)
		return err
	})
	return err
}

func FetchAllFromDB() ([][]byte, error) {
	var data [][]byte
	err := Database.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(DB_BUCKET_NAME))
		if err != nil {
			return err
		}
		err = bucket.ForEach(func(k, v []byte) error {
			data = append(data, v)
			return nil
		})
		return err
	})
	return data, err
}
