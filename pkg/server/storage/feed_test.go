package storage

import (
	"sort"
	"testing"

	"github.com/mradile/rssfeeder"
	"github.com/stretchr/testify/assert"
)

func setupNewFeedStorage(t *testing.T) rssfeeder.FeedStorage {
	db, err := getDB()
	assert.Nil(t, err)
	return NewFeedStorage(db)
}

func Test_Feeds(t *testing.T) {
	fs := setupNewFeedStorage(t)
	defer fs.(*feedStorage).db.Close()

	//add a feed
	f1, err := testAddFeed("ccc", "a", "", fs)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, f1.ID)
	assert.NotEqual(t, "", f1.Token)

	f2, err := testAddFeed("aaa", "a", "", fs)
	f3, err := testAddFeed("fff", "a", "", fs)
	f4, err := testAddFeed("ggg", "g", "", fs)

	//get a feed
	nilFeed, err := fs.Get(0)
	assert.Nil(t, err)
	assert.Nil(t, nilFeed)
	gotF1, err := fs.Get(f1.ID)
	assert.Nil(t, err)
	assert.Equal(t, f1, gotF1)

	//exists f1, f2, f3
	exists, err := fs.Exists(f1.Login, f1.Name)
	assert.Nil(t, err)
	assert.True(t, exists)

	exists, err = fs.Exists(f2.Login, f2.Name)
	assert.Nil(t, err)
	assert.True(t, exists)

	exists, err = fs.Exists(f3.Login, f3.Name)
	assert.Nil(t, err)
	assert.True(t, exists)

	exists, err = fs.Exists(f4.Login, f4.Name)
	assert.Nil(t, err)
	assert.True(t, exists)

	//get by name and login
	ff1, err := fs.GetByNameAndLogin(f1.Login, f1.Name)
	assert.Nil(t, err)
	assert.Equal(t, f1, ff1)
	ff2, err := fs.GetByNameAndLogin(f2.Login, f2.Name)
	assert.Nil(t, err)
	assert.Equal(t, f2, ff2)
	ff4, err := fs.GetByNameAndLogin(f4.Login, f4.Name)
	assert.Nil(t, err)
	assert.Equal(t, f4, ff4)

	feedsExp := []*rssfeeder.Feed{f1, f2, f3}
	sort.Slice(feedsExp, func(i, j int) bool {
		return feedsExp[i].Name < feedsExp[j].Name
	})
	feedsGot, err := fs.GetFeedsByLogin(f1.Login)
	assert.Nil(t, err)
	assert.Equal(t, feedsExp, feedsGot)
}

func Test_Feeds_Delete(t *testing.T) {
	fs := setupNewFeedStorage(t)
	fes := NewFeedEntryStorage(fs.(*feedStorage).db)
	defer fs.(*feedStorage).db.Close()

	//delete not existing feed
	assert.Error(t, fs.Delete(0)) //does not exist

	//delete feed without entries
	fe, _ := testAddFeed("z", "z", "", fs)
	assert.NoError(t, fs.Delete(fe.ID))
	deletedFab, err := fs.Get(fe.ID)
	assert.Nil(t, err)
	assert.Nil(t, deletedFab)

	faa, _ := testAddFeed("a", "a", "", fs)
	assert.NoError(t, fes.Add(&rssfeeder.FeedEntry{
		Login:    "a",
		FeedName: "a",
		URI:      "a",
	}))
	fba, _ := testAddFeed("a", "b", "", fs)
	fbaEntry := &rssfeeder.FeedEntry{
		Login:    "b",
		FeedName: "a",
		URI:      "a",
	}
	assert.NoError(t, fes.Add(fbaEntry))

	assert.NoError(t, fs.Delete(faa.ID))
	emptyList, err := fes.AllByLoginAndFeedName(faa.Login, faa.Name)
	assert.Nil(t, err)
	assert.Nil(t, emptyList)

	fbaGot, err := fs.Get(fba.ID)
	assert.Nil(t, err)
	assert.Equal(t, fba, fbaGot)
	fbaEntries, err := fes.AllByLoginAndFeedName(fba.Login, fba.Name)
	assert.Nil(t, err)
	assert.Equal(t, fbaEntry.ID, fbaEntries[0].ID)

}

func testAddFeed(name, login, token string, fs rssfeeder.FeedStorage) (*rssfeeder.Feed, error) {
	f := &rssfeeder.Feed{
		Name:  name,
		Login: login,
		Token: token,
	}
	err := fs.Add(f)
	return f, err
}

func Test_feedStorage_NoFound(t *testing.T) {

	type args struct {
		login string
		name  string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "empty",
			args: args{},
		},
		{
			name: "login empty",
			args: args{
				name: "a",
			},
		},
		{
			name: "name empty",
			args: args{
				login: "a",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := setupNewFeedStorage(t)
			defer fs.(*feedStorage).db.Close()

			exists, err := fs.Exists(tt.args.login, tt.args.name)
			assert.Nil(t, err)
			assert.False(t, exists)

			f, err := fs.GetByNameAndLogin(tt.args.login, tt.args.name)
			assert.Nil(t, err)
			assert.Nil(t, f)

			feeds, err := fs.GetFeedsByLogin(tt.args.login)
			assert.Nil(t, err)
			assert.Nil(t, feeds)
		})
	}
}
