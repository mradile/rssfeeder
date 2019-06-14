package storage

import (
	"github.com/asdine/storm"
	"github.com/mradile/rssfeeder/pkg/server/configuration"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func Test_NewDB(t *testing.T) {
	db, err := getDB()
	assert.Nil(t, err)
	assert.NotNil(t, db)
}

func getDB() (*storm.DB, error) {
	tmpPath := os.TempDir()
	dbFilePath := path.Join(os.TempDir(), dbFileName)
	_ = os.Remove(dbFilePath)
	cfg := &configuration.Configuration{
		DBPath: tmpPath,
	}
	return NewStormDB(cfg)
}
