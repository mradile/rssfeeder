package storage

import (
	"time"

	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	"github.com/mradile/rssfeeder"
)

type FeedEntry struct {
	ID         int    `storm:"id,increment"`
	Login      string `storm:"index"`
	Category   string `storm:"index"`
	URI        string
	CreateDate time.Time
}

func toDBFeedEntry(fe *rssfeeder.FeedEntry) *FeedEntry {
	return &FeedEntry{
		ID:         fe.ID,
		Login:      fe.Login,
		Category:   fe.Category,
		URI:        fe.URI,
		CreateDate: fe.CreateDate,
	}
}

func (fe *FeedEntry) fromDBFeedEntry() *rssfeeder.FeedEntry {
	return &rssfeeder.FeedEntry{
		ID:         fe.ID,
		Login:      fe.Login,
		Category:   fe.Category,
		URI:        fe.URI,
		CreateDate: fe.CreateDate,
	}
}

type feedEntryStorage struct {
	db *storm.DB
}

func NewFeedEntryStorage(db *storm.DB) rssfeeder.FeedEntryStorage {
	return &feedEntryStorage{db: db}
}

func (s *feedEntryStorage) Add(entry *rssfeeder.FeedEntry) error {
	e := toDBFeedEntry(entry)
	err := s.db.Save(e)
	if err != nil {
		return err
	}
	entry.ID = e.ID
	return nil
}

func (s *feedEntryStorage) Delete(id int) error {
	return s.db.DeleteStruct(&rssfeeder.FeedEntry{ID: id})
}

func (s *feedEntryStorage) AllByLoginAndCategory(login string, category string) ([]*rssfeeder.FeedEntry, error) {
	var entries []*FeedEntry
	query := s.db.Select(q.And(
		q.Eq("Login", login),
		q.Eq("Category", category),
	)).OrderBy("CreateDate").Reverse()

	err := query.Find(&entries)
	if err != nil {
		return nil, err
	}
	fentries := make([]*rssfeeder.FeedEntry, 0, len(entries))
	for _, e := range entries {
		fentries = append(fentries, e.fromDBFeedEntry())
	}
	return fentries, err
}

func (s *feedEntryStorage) ExistsEntry(id int) (bool, error) {
	var e *FeedEntry
	if err := s.db.One("ID", id, &e); err != nil {
		if err == storm.ErrNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *feedEntryStorage) EntryBelongsToLogin(id int, login string) (bool, error) {
	var e *FeedEntry
	query := s.db.Select(q.And(
		q.Eq("ID", id),
		q.Eq("Login", login),
	))

	if err := query.First(e); err != nil {
		if err == storm.ErrNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *feedEntryStorage) GetCategories(login string) ([]string, error) {
	query := s.db.Select(q.And(
		q.Eq("Login", login),
	)).OrderBy("Category")

	var entries []*FeedEntry
	if err := query.Find(&entries); err != nil {
		return nil, err
	}

	allCats := make(map[string]bool)
	for _, entry := range entries {
		allCats[entry.Category] = true
	}

	categories := make([]string, 0, len(allCats))
	for cat := range allCats {
		categories = append(categories, cat)
	}

	return categories, nil
}
