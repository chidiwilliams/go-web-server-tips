package server

import (
	"github.com/tidwall/buntdb"
)

func connectToInMemoryDB() (*buntdb.DB, error) {
	return buntdb.Open(":memory:")
}
