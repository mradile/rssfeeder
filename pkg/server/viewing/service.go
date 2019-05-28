package viewing

import (
	"errors"

	"github.com/mradile/rssfeeder"
)

type ViewingService interface {
	GetCategories(login string) ([]string, error)
	GetFeed(category string, login string) ([]*rssfeeder.FeedEntry, error)
}

type viewer struct {
	entries rssfeeder.FeedEntryStorage
}

func NewViewingService(entryStore rssfeeder.FeedEntryStorage) ViewingService {
	v := &viewer{
		entries: entryStore,
	}
	return v
}

func (v *viewer) GetCategories(login string) ([]string, error) {
	if login == "" {
		return nil, errors.New("invalid login")
	}
	categories, err := v.entries.GetCategories(login)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (v *viewer) GetFeed(category string, login string) ([]*rssfeeder.FeedEntry, error) {
	if login == "" {
		return nil, errors.New("invalid login")
	}
	if category == "" {
		category = rssfeeder.DefaultCategory
	}
	return v.entries.AllByLoginAndCategory(login, category)
}
