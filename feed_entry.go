package rssfeeder

import "time"

type FeedEntry struct {
	ID         int
	Login      string
	Category   string
	URI        string
	CreateDate time.Time
}

const DefaultCategory = "default"

type FeedEntryStorage interface {
	Add(entry *FeedEntry) error
	Delete(id int) error
	AllByLoginAndCategory(login string, category string) ([]*FeedEntry, error)
	ExistsEntry(id int) (bool, error)
	EntryBelongsToLogin(id int, login string) (bool, error)
	GetCategories(login string) ([]string, error)
}
