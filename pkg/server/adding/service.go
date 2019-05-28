package adding

import (
	"errors"
	"github.com/mradile/rssfeeder"
)

type AddingService interface {
	AddFeedEntry(entry *rssfeeder.FeedEntry) error
}

type adder struct {
	entries rssfeeder.FeedEntryStorage
}

func NewAddingService(entryStore rssfeeder.FeedEntryStorage) AddingService {
	a := &adder{
		entries: entryStore,
	}
	return a
}

func (a *adder) AddFeedEntry(entry *rssfeeder.FeedEntry) error {
	if entry.Login == "" {
		return errors.New("invalid login")
	}
	if entry.Category == "" {
		entry.Category = rssfeeder.DefaultCategory
	}

	if err := a.entries.Add(entry); err != nil {
		return err
	}

	return nil
}
