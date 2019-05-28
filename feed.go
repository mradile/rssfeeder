package rssfeeder

type RSSFeed struct {
	ID     int
	UserID int
	Name   string
	Key    string
}

type FeedStorage interface {
	Add(feed *RSSFeed) error
	Delete(id int) error
	Update(feed *RSSFeed) error
	AllByLogin(login string) ([]*RSSFeed, error)
}
