package deleting

import (
	"github.com/mradile/rssfeeder"
)

type deleter struct {
	entries rssfeeder.FeedEntryStorage
	feeds   rssfeeder.FeedStorage
}

func NewDeletingService(entries rssfeeder.FeedEntryStorage, feeds rssfeeder.FeedStorage) rssfeeder.DeletingService {
	a := &deleter{
		entries: entries,
		feeds:   feeds,
	}
	return a
}

func (d *deleter) DeleteFeed(id int, login string) error {
	if login == "" {
		return rssfeeder.ErrEmptyLogin
	}

	feed, err := d.feeds.Get(id)
	if err != nil {
		return err
	}

	if feed == nil {
		return rssfeeder.ErrFeedMissing
	}

	if login != feed.Login {
		return rssfeeder.ErrNotAllowed
	}

	return d.feeds.Delete(id)
}

func (d *deleter) DeleteFeedEntry(id int, login string) error {
	if login == "" {
		return rssfeeder.ErrEmptyLogin
	}

	if existsForLogin, err := d.entries.EntryBelongsToLogin(id, login); err != nil {
		return err
	} else if !existsForLogin {
		return rssfeeder.ErrEntryMissing
	}

	if err := d.entries.Delete(id); err != nil {
		return err
	}

	return nil
}
