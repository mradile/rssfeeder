package storage

import (
	"testing"
	"time"

	"github.com/mradile/rssfeeder"
	"github.com/stretchr/testify/assert"
)

func setupNewFeedEntryStorage(t *testing.T) rssfeeder.FeedEntryStorage {
	db, err := getDB()
	assert.Nil(t, err)
	return NewFeedEntryStorage(db)
}

func TestFeedEntryStorage(t *testing.T) {
	fes := setupNewFeedEntryStorage(t)
	defer fes.(*feedEntryStorage).db.Close()

	loginA := "a"
	loginZ := "z"

	//add some entries
	ti := time.Now().Add(time.Second * 100)
	fe1, err := testAddFeedEntry(loginA, "a", "a1", &ti, fes)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, fe1.ID)
	fe2, err := testAddFeedEntry(loginA, "a", "a2", nil, fes)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, fe2.ID)
	fe3, err := testAddFeedEntry(loginA, "b", "b1", nil, fes)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, fe3.ID)

	fe4, err := testAddFeedEntry(loginZ, "z", "z1", nil, fes)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, fe4.ID)
	fe5, err := testAddFeedEntry(loginZ, "z", "z2", nil, fes)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, fe5.ID)

	//exists
	exists, err := fes.ExistsEntry(fe1.ID)
	assert.Nil(t, err)
	assert.True(t, exists)
	exists, err = fes.ExistsEntry(fe4.ID)
	assert.Nil(t, err)
	assert.True(t, exists)

	//delete
	err = fes.Delete(fe5.ID)
	assert.Nil(t, err)

	//does not exist
	exists, err = fes.ExistsEntry(fe5.ID)
	assert.Nil(t, err)
	assert.False(t, exists)

	//belongs to login
	if ok, err := fes.EntryBelongsToLogin(fe1.ID, loginA); err != nil {
		assert.Fail(t, "no error expected ", err)
	} else if !ok {
		assert.Fail(t, "should belong to login", err)
	}
	if ok, err := fes.EntryBelongsToLogin(fe4.ID, loginZ); err != nil {
		assert.Fail(t, "no error expected ", err)
	} else if !ok {
		assert.Fail(t, "should belong to login")
	}

	//belongs not to login
	if ok, err := fes.EntryBelongsToLogin(fe1.ID, loginZ); err != nil {
		assert.Fail(t, "no error expected ", err)
	} else if ok {
		assert.Fail(t, "should not belong to login")
	}

	//get feed content
	if e, err := fes.AllByLoginAndFeedName(loginA, "a"); err != nil {
		assert.Fail(t, "error: ", err)
	} else {
		assert.Equal(t, 2, len(e))
		assert.Equal(t, fe1.ID, e[0].ID)
		assert.Equal(t, fe2.ID, e[1].ID)
	}

	//get feed content
	if e, err := fes.AllByLoginAndFeedName(loginZ, "z"); err != nil {
		assert.Fail(t, "error: ", err)
	} else {
		assert.Equal(t, 1, len(e))
		assert.Equal(t, fe4.ID, e[0].ID)
	}

	e, err := fes.AllByLoginAndFeedName("", "")
	assert.Nil(t, err)
	assert.Nil(t, e)
}

func TestFeedEntryStorage_Add(t *testing.T) {

	type args struct {
		login    string
		feedName string
		uri      string
	}
	var tests = []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				login:    "",
				feedName: "",
				uri:      "",
			},
		},
		{
			name: "2",
			args: args{
				login:    "a",
				feedName: "",
				uri:      "",
			},
		},
		{
			name: "3",
			args: args{
				login:    "a",
				feedName: "a",
				uri:      "",
			},
		},
		{
			name: "4",
			args: args{
				login:    "a",
				feedName: "",
				uri:      "a",
			},
		},
		{
			name: "5",
			args: args{
				login:    "",
				feedName: "a",
				uri:      "a",
			},
		},
		{
			name: "6",
			args: args{
				login:    "",
				feedName: "",
				uri:      "a",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fes := setupNewFeedEntryStorage(t)
			defer fes.(*feedEntryStorage).db.Close()

			fe, err := testAddFeedEntry(
				tt.args.login,
				tt.args.feedName,
				tt.args.uri,
				nil,
				fes,
			)
			assert.Error(t, err)
			assert.Equal(t, 0, fe.ID)

		})
	}
}

func testAddFeedEntry(login, feedName, uri string, date *time.Time, fes rssfeeder.FeedEntryStorage) (*rssfeeder.FeedEntry, error) {
	if date == nil {
		t := time.Now()
		date = &t
	}

	fe := &rssfeeder.FeedEntry{
		Login:      login,
		FeedName:   feedName,
		URI:        uri,
		CreateDate: *date,
	}
	return fe, fes.Add(fe)
}
