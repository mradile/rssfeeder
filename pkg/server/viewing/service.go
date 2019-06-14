package viewing

import (
	"github.com/mradile/rssfeeder"
)

type viewer struct {
	entries rssfeeder.FeedEntryStorage
	feeds   rssfeeder.FeedStorage
}

func NewViewingService(entryStore rssfeeder.FeedEntryStorage, feeds rssfeeder.FeedStorage) rssfeeder.ViewingService {
	v := &viewer{
		entries: entryStore,
		feeds:   feeds,
	}
	return v
}

func (v *viewer) GetFeeds(login string) ([]*rssfeeder.Feed, error) {
	if login == "" {
		return nil, rssfeeder.ErrEmptyLogin
	}

	feeds, err := v.feeds.GetFeedsByLogin(login)
	if err != nil {
		return nil, err
	}
	return feeds, nil
}

func (v *viewer) GetFeed(feedName string, login string) ([]*rssfeeder.FeedEntry, error) {
	if login == "" {
		return nil, rssfeeder.ErrEmptyLogin
	}
	if feedName == "" {
		feedName = rssfeeder.DefaultFeedName
	}
	return v.entries.AllByLoginAndFeedName(login, feedName)
}
