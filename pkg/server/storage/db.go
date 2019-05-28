package storage

import (
	"path"

	"github.com/sirupsen/logrus"

	"github.com/asdine/storm"
	"github.com/mradile/rssfeeder/pkg/server/configuration"
)

const dbFileName = "rssfeeder.db"

func NewStormDB(cfg *configuration.Configuration) (*storm.DB, error) {
	dbFilePath := path.Join(cfg.DBPath, dbFileName)
	db, err := storm.Open(dbFilePath)
	if err != nil {
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"db_file_path": dbFilePath,
	}).Info("opened db file")

	return db, nil
}
