package storage

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mradile/rssfeeder/pkg/server/configuration"

	"github.com/asdine/storm"
	"github.com/mradile/rssfeeder"
)

func getDB() (*storm.DB, error) {
	tmpPath := os.TempDir()
	dbFilePath := path.Join(os.TempDir(), dbFileName)
	os.Remove(dbFilePath)
	cfg := &configuration.Configuration{
		DBPath: tmpPath,
	}
	return NewStormDB(cfg)
}

func Test_userStorage(t *testing.T) {
	db, err := getDB()
	defer db.Close()
	assert.Nil(t, err)
	s := NewUserStorage(db)

	//add user1
	user1 := &rssfeeder.User{
		Login:    "bla",
		Password: "blub",
	}
	assert.NoError(t, s.Add(user1))

	//add user2
	user2 := &rssfeeder.User{
		Login:    "bla2",
		Password: "blub2",
	}
	assert.NoError(t, s.Add(user2))

	//test unique constraint
	user1Duplicate := &rssfeeder.User{
		Login: "bla",
	}
	assert.Error(t, s.Add(user1Duplicate))

	//get user
	getUser1, err := s.Get(user1.Login)
	assert.Nil(t, err)
	assert.Equal(t, user1, getUser1)

	//get not existing user
	notExists, err := s.Get("asdsad")
	assert.Nil(t, err)
	assert.Nil(t, notExists)

	//update
	user1.Password = "changed"
	assert.NoError(t, s.Update(user1))
	updated, err := s.Get(user1.Login)
	assert.Nil(t, err)
	assert.Equal(t, user1, updated)

	//update user without id
	assert.Error(t, s.Update(&rssfeeder.User{
		Login:    user2.Login,
		Password: "changed",
	}))

	//delete
	assert.NoError(t, s.Delete(user1.Login))
	deleted, err := s.Get(user1.Login)
	assert.Nil(t, err)
	assert.Nil(t, deleted)

	//delete not existing user
	assert.Error(t, s.Delete(user1.Login))
}
