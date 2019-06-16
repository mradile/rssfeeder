package deleting

import (
	"github.com/golang/mock/gomock"
	"github.com/mradile/rssfeeder/pkg/server/mock"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/mradile/rssfeeder"
)

func Test_deleter_DeleteFeed(t *testing.T) {

	type args struct {
		id    int
		login string
	}
	tests := []struct {
		name      string
		args      args
		wantErr   bool
		errorType error
		exp       func(fes *mock.MockFeedEntryStorage, fs *mock.MockFeedStorage)
	}{
		{
			name: "empty login and id",
			args: args{
				id:    0,
				login: "",
			},
			errorType: rssfeeder.ErrEmptyLogin,
			wantErr:   true,
		},
		{
			name: "empty id",
			args: args{
				id:    0,
				login: "a",
			},
			errorType: rssfeeder.ErrFeedMissing,
			wantErr:   true,
			exp: func(fes *mock.MockFeedEntryStorage, fs *mock.MockFeedStorage) {
				fs.EXPECT().Get(gomock.Eq(0))
			},
		},
		{
			name: "not allowed",
			args: args{
				id:    10,
				login: "a",
			},
			errorType: rssfeeder.ErrNotAllowed,
			wantErr:   true,
			exp: func(fes *mock.MockFeedEntryStorage, fs *mock.MockFeedStorage) {
				fs.EXPECT().Get(gomock.Eq(10)).Return(&rssfeeder.Feed{
					ID:    10,
					Name:  "a",
					Login: "b",
				}, nil)
			},
		},
		{
			name: "ok",
			args: args{
				id:    10,
				login: "a",
			},
			wantErr: false,
			exp: func(fes *mock.MockFeedEntryStorage, fs *mock.MockFeedStorage) {
				fs.EXPECT().Get(gomock.Eq(10)).Return(&rssfeeder.Feed{
					ID:    10,
					Name:  "a",
					Login: "a",
				}, nil)
				fs.EXPECT().Delete(gomock.Eq(10)).Return(nil)
			},
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
			ds := NewDeletingService(fes, fe)
			err := ds.DeleteFeed(tt.args.id, tt.args.login)
			if tt.wantErr {
				assert.Equal(t, tt.errorType, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func Test_deleter_DeleteFeedEntry(t *testing.T) {
	type args struct {
		id    int
		login string
	}
	tests := []struct {
		name      string
		args      args
		wantErr   bool
		errorType error
		exp       func(fes *mock.MockFeedEntryStorage)
	}{
		{
			name: "empty login and id",
			args: args{
				id:    0,
				login: "",
			},
			wantErr:   true,
			errorType: rssfeeder.ErrEmptyLogin,
		},
		{
			name: "not exists",
			args: args{
				id:    0,
				login: "a",
			},
			wantErr:   true,
			errorType: rssfeeder.ErrEntryMissing,
			exp: func(fes *mock.MockFeedEntryStorage) {
				fes.
					EXPECT().EntryBelongsToLogin(gomock.Eq(0), gomock.Eq("a")).
					Return(false, nil)
			},
		},
		{
			name: "ok",
			args: args{
				id:    10,
				login: "a",
			},
			wantErr: false,
			exp: func(fes *mock.MockFeedEntryStorage) {
				fes.EXPECT().
					EntryBelongsToLogin(gomock.Eq(10), gomock.Eq("a")).
					Return(true, nil)
				fes.EXPECT().
					Delete(gomock.Eq(10))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			fe := mock.NewMockFeedStorage(ctrl)
			fes := mock.NewMockFeedEntryStorage(ctrl)
			defer ctrl.Finish()

			if tt.exp != nil {
				tt.exp(fes)
			}
			ds := NewDeletingService(fes, fe)
			err := ds.DeleteFeedEntry(tt.args.id, tt.args.login)
			if tt.wantErr {
				assert.Equal(t, tt.errorType, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
