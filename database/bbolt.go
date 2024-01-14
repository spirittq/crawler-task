package database

import (
	"crawler-task/utils"
	"encoding/json"
	"errors"

	"github.com/rs/zerolog/log"

	"go.etcd.io/bbolt"
)

var Database *bbolt.DB
var DB_BUCKET_NAME = utils.GetEnvOrDefault("DB_BUCKET_NAME", "")

func InitDB() {
	var err error
	Database, err = bbolt.Open(utils.GetEnvOrDefault("DB_NAME", ""), 0600, nil)
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
			return errors.New("record already exists in database")
		}

		json, err := json.Marshal(data)
		if err != nil {
			return err
		}
		return bucket.Put(key, json)
	})
	return err
}
