package viewing

import (
	"github.com/golang/mock/gomock"
	"github.com/mradile/rssfeeder"
	"github.com/mradile/rssfeeder/pkg/server/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_viewer_GetFeeds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fe := mock.NewMockFeedStorage(ctrl)
	fes := mock.NewMockFeedEntryStorage(ctrl)
	s := NewViewingService(fes, fe)

	//empty login
	got, err := s.GetFeeds("")
	assert.Equal(t, rssfeeder.ErrEmptyLogin, err)
	assert.Nil(t, got)
}

func Test_viewer_GetFeed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fe := mock.NewMockFeedStorage(ctrl)
	fes := mock.NewMockFeedEntryStorage(ctrl)
	s := NewViewingService(fes, fe)

	//empty login
	got, err := s.GetFeed("a", "")
	assert.Equal(t, rssfeeder.ErrEmptyLogin, err)
	assert.Nil(t, got)

	//default feed
	fes.EXPECT().
		AllByLoginAndFeedName(
			gomock.Eq("a"),
			gomock.Eq(rssfeeder.DefaultFeedName),
		).
		Return([]*rssfeeder.FeedEntry{
			{
				ID:       1,
				Login:    "a",
				FeedName: rssfeeder.DefaultFeedName,
				URI:      "",
			},
		}, nil)
	got, err = s.GetFeed("", "a")
	assert.Nil(t, err)
	assert.NotNil(t, got)
}
