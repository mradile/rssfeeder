package deleting

import (
	"errors"

	"github.com/mradile/rssfeeder"
)

var (
	ErrNotAllowed = errors.New("not allowed")
)

type DeletingService interface {
	DeleteFeedEntry(id int, login string) error
}

type adder struct {
	entries rssfeeder.FeedEntryStorage
}

func NewDeletingService(entryStore rssfeeder.FeedEntryStorage) DeletingService {
	a := &adder{
		entries: entryStore,
	}
	return a
}

func (a *adder) DeleteFeedEntry(id int, login string) error {
	if login == "" {
		return errors.New("invalid login")
	}

	if exists, err := a.entries.ExistsEntry(id); err != nil {
		return err
	} else if !exists {
		return errors.New("entry does not exist")
	}

	if belongs, err := a.entries.EntryBelongsToLogin(id, login); err != nil {
		return err
	} else if !belongs {
		return ErrNotAllowed
	}

	if err := a.entries.Delete(id); err != nil {
		return err
	}

	return nil
}
