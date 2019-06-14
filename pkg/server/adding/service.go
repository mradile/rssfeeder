package adding

import (
	"github.com/mradile/rssfeeder"
	"github.com/pkg/errors"
)

type adder struct {
	entries rssfeeder.FeedEntryStorage
	feeds   rssfeeder.FeedStorage
}

func NewAddingService(entries rssfeeder.FeedEntryStorage, feeds rssfeeder.FeedStorage) rssfeeder.AddingService {
	a := &adder{
		entries: entries,
		feeds:   feeds,
	}
	return a
}

func (a *adder) AddFeedEntry(entry *rssfeeder.FeedEntry) error {
	if entry.Login == "" {
		return rssfeeder.ErrEmptyLogin
	}
	if entry.URI == "" {
		return rssfeeder.ErrEmptyURI
	}

	if entry.FeedName == "" {
		entry.FeedName = rssfeeder.DefaultFeedName
	}

	exists, err := a.feeds.Exists(entry.FeedName, entry.Login)
	if err != nil {
		return errors.Wrap(err, "could not fetch feed for adding feed entry")
	}
	if !exists {
		feed := &rssfeeder.Feed{
			Name:  entry.FeedName,
			Login: entry.Login,
		}
		if err := a.feeds.Add(feed); err != nil {
			return errors.Wrap(err, "could not create new feed for adding feed entry")
		}
	}

	if err := a.entries.Add(entry); err != nil {
		return err
	}

	return nil
}
