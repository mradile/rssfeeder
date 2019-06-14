package deleting

import (
	"github.com/mradile/rssfeeder"
)

type adder struct {
	entries rssfeeder.FeedEntryStorage
}

func NewDeletingService(entryStore rssfeeder.FeedEntryStorage) rssfeeder.DeletingService {
	a := &adder{
		entries: entryStore,
	}
	return a
}

func (a *adder) DeleteFeedEntry(id int, login string) error {
	panic("not implemented")
	/*
		if login == "" {
			return rssfeeder.ErrEmptyLogin
		}

		if belongs, err := a.entries.EntryBelongsToLogin(id, login); err != nil {
			return err
		} else if !belongs {
			return rssfeeder.ErrNotAllowed
		}

		if exists, err := a.entries.ExistsEntry(id); err != nil {
			return err
		} else if !exists {
			return rssfeeder.ErrEntryMissing
		}
		if err := a.entries.Delete(id); err != nil {
			return err
		}

		return nil

	*/
}
