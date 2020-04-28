package server

import (
	"gopkg.in/mgo.v2"
)

var mongoIndexes = map[string][]mgo.Index{}

func connectToMongo() (*mgo.Database, error) {
	session, err := mgo.Dial("mongodb://localhost:27017/")
	if err != nil {
		return nil, err
	}

	return session.DB("go"), nil
}

func ensureMongoIndexes(db *mgo.Database) error {
	for coll, indexes := range mongoIndexes {
		for _, index := range indexes {
			err := db.C(coll).EnsureIndex(index)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
