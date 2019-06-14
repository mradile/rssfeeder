package storage

import (
	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	"github.com/mradile/rssfeeder"
	"github.com/pkg/errors"
	"math/rand"
)

type Feed struct {
	ID    int `storm:"id,increment"`
	Name  string
	Token string `storm:"index,unique"`
	Login string `storm:"index"`
}

func toDBFeed(f *rssfeeder.Feed) *Feed {
	return &Feed{
		ID:    f.ID,
		Name:  f.Name,
		Token: f.Token,
		Login: f.Login,
	}
}

func (f *Feed) fromDBFeed() *rssfeeder.Feed {
	return &rssfeeder.Feed{
		ID:    f.ID,
		Name:  f.Name,
		Token: f.Token,
		Login: f.Login,
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

const tokenLength = 12

type feedStorage struct {
	db *storm.DB
}

func NewFeedStorage(db *storm.DB) rssfeeder.FeedStorage {
	return &feedStorage{db: db}
}

func (fs *feedStorage) Add(feed *rssfeeder.Feed) error {
	if feed.Token == "" {
		var token string
		for {
			token = getToken()
			if exists, err := fs.tokenExists(token); err != nil {
				return errors.Wrap(err, "could not query existing token")
			} else if !exists {
				feed.Token = token
				break
			}
		}
	}
	e := toDBFeed(feed)
	err := fs.db.Save(e)
	if err != nil {
		return err
	}
	feed.ID = e.ID
	return nil
}

func (fs *feedStorage) Exists(login, name string) (bool, error) {
	f, err := fs.GetByNameAndLogin(login, name)
	if err != nil {
		return false, err
	}
	if f == nil {
		return false, err
	}
	return true, nil
}

func (fs *feedStorage) GetByNameAndLogin(login, name string) (*rssfeeder.Feed, error) {
	var f Feed
	query := fs.db.Select(q.And(
		q.Eq("Name", name),
		q.Eq("Login", login),
	))
	if err := query.First(&f); err != nil {
		if err == storm.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return f.fromDBFeed(), nil
}

func (fs *feedStorage) Delete(id int) error {
	panic("implement me")
}

func (fs *feedStorage) GetFeedsByLogin(login string) ([]*rssfeeder.Feed, error) {
	var dbFeeds []*Feed

	query := fs.db.Select(q.And(
		q.Eq("Login", login),
	)).OrderBy("Name")

	if err := query.Find(&dbFeeds); err != nil {
		if err == storm.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}

	feeds := make([]*rssfeeder.Feed, 0, len(dbFeeds))
	for _, f := range dbFeeds {
		feeds = append(feeds, f.fromDBFeed())
	}

	return feeds, nil
}

func (fs *feedStorage) tokenExists(token string) (bool, error) {
	var feed Feed
	err := fs.db.One("Token", token, &feed)
	if err != nil {
		if err == storm.ErrNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func getToken() string {
	b := make([]rune, tokenLength)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
