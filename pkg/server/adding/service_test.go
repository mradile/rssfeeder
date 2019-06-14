package adding

import (
	"github.com/golang/mock/gomock"
	"github.com/mradile/rssfeeder"
	"github.com/mradile/rssfeeder/pkg/server/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_adder_AddFeedEntry(t *testing.T) {
	type args struct {
		entry *rssfeeder.FeedEntry
	}
	tests := []struct {
		name    string
		args    args
		exp     func(fes *mock.MockFeedEntryStorage, fs *mock.MockFeedStorage)
		ass     func(e *rssfeeder.FeedEntry)
		wantErr bool
	}{
		{
			name: "empty login",
			args: args{
				entry: &rssfeeder.FeedEntry{
					Login:    "",
					FeedName: "a",
					URI:      "b",
				},
			},
			wantErr: true,
		},
		{
			name: "empty uri",
			args: args{
				entry: &rssfeeder.FeedEntry{
					Login:    "a",
					FeedName: "b",
					URI:      "",
				},
			},
			wantErr: true,
		},
		{
			name: "default feed",
			args: args{
				entry: &rssfeeder.FeedEntry{
					Login: "a",
					URI:   "b",
				},
			},
			exp: func(fes *mock.MockFeedEntryStorage, fs *mock.MockFeedStorage) {
				fs.EXPECT().Exists(gomock.Any(), gomock.Any()).Return(false, nil)
				fs.EXPECT().Add(gomock.Any()).Return(nil)
				fes.EXPECT().Add(gomock.Any()).Return(nil)
			},
			ass: func(e *rssfeeder.FeedEntry) {
				assert.Equal(t, rssfeeder.DefaultFeedName, e.FeedName)
			},
			wantErr: false,
		},
		{
			name: "feed exists",
			args: args{
				entry: &rssfeeder.FeedEntry{
					Login:    "a",
					FeedName: "b",
					URI:      "c",
				},
			},
			exp: func(fes *mock.MockFeedEntryStorage, fs *mock.MockFeedStorage) {
				fs.EXPECT().Exists(gomock.Any(), gomock.Any()).Return(true, nil)
				fes.EXPECT().Add(gomock.Any()).Return(nil)
			},
			ass: func(e *rssfeeder.FeedEntry) {
				assert.Equal(t, "b", e.FeedName)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			fe := mock.NewMockFeedStorage(ctrl)
			fes := mock.NewMockFeedEntryStorage(ctrl)
			defer ctrl.Finish()

			if tt.exp != nil {
				tt.exp(fes, fe)
			}

			s := NewAddingService(fes, fe)

			if err := s.AddFeedEntry(tt.args.entry); (err != nil) != tt.wantErr {
				t.Errorf("adder.AddFeedEntry() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.ass != nil {
				tt.ass(tt.args.entry)
			}
		})
	}
}
