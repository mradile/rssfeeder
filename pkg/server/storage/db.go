package storage

import (
	"github.com/asdine/storm"
	"github.com/mradile/rssfeeder/pkg/server/configuration"
	"github.com/sirupsen/logrus"
	"path"
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
