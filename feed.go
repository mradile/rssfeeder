package rssfeeder

import "time"

//go:generate mockgen -destination=pkg/server/mock/feed.go -package=mock github.com/mradile/rssfeeder FeedStorage,FeedEntryStorage,AddingService,DeletingService,ViewingService

//Feed represents a RSS feed
type Feed struct {
	ID    int
	Name  string
	Token string
	Login string
}

//FeedStorage manages persistence of a RSS feed
type FeedStorage interface {
	Add(feed *Feed) error
	Get(id int) (*Feed, error)
	Exists(login, name string) (bool, error)
	GetByNameAndLogin(login, name string) (*Feed, error)
	Delete(id int) error
	GetFeedsByLogin(login string) ([]*Feed, error)
}

//FeedEntry represents a single item in a RSS feed
type FeedEntry struct {
	ID         int
	Login      string
	FeedName   string
	URI        string
	CreateDate time.Time
}

//DefaultFeedName is the default name of a feed where entries are stored
const DefaultFeedName = "default"

//FeedEntryStorage manages persistence of feed entries
type FeedEntryStorage interface {
	Add(entry *FeedEntry) error
	Delete(id int) error
	AllByLoginAndFeedName(login string, feedName string) ([]*FeedEntry, error)
	ExistsEntry(id int) (bool, error)
	EntryBelongsToLogin(id int, login string) (bool, error)
}

type AddingService interface {
	AddFeedEntry(entry *FeedEntry) error
}

type DeletingService interface {
	DeleteFeed(id int, login string) error
	DeleteFeedEntry(id int, login string) error
}

type ViewingService interface {
	GetFeeds(login string) ([]*Feed, error)
	GetFeed(feedName string, login string) ([]*FeedEntry, error)
}
