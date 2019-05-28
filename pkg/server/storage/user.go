package storage

import (
	"errors"

	"github.com/asdine/storm"
	"github.com/mradile/rssfeeder"
)

type userStorage struct {
	db *storm.DB
}

type User struct {
	ID       int    `storm:"id,increment"`
	Login    string `storm:"unique"`
	Password string
}

func (u *User) toUser() *rssfeeder.User {
	return &rssfeeder.User{
		ID:       u.ID,
		Login:    u.Login,
		Password: u.Password,
	}
}

func toDBUser(u *rssfeeder.User) *User {
	return &User{
		ID:       u.ID,
		Login:    u.Login,
		Password: u.Password,
	}
}

func NewUserStorage(db *storm.DB) rssfeeder.UserStorage {
	return &userStorage{db: db}
}

func (s *userStorage) Add(user *rssfeeder.User) error {
	u := toDBUser(user)
	err := s.db.Save(u)
	if err != nil {
		return err
	}
	user.ID = u.ID
	return nil
}

func (s *userStorage) Update(user *rssfeeder.User) error {
	if user.ID == 0 {
		return errors.New("zero value for id")
	}
	return s.db.Update(toDBUser(user))
}

func (s *userStorage) Delete(login string) error {
	user, err := s.Get(login)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user does not exist")
	}
	return s.db.DeleteStruct(&User{ID: user.ID})
}

func (s *userStorage) Get(login string) (*rssfeeder.User, error) {
	var u User
	if err := s.db.One("Login", login, &u); err != nil {
		if err == storm.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}

	return u.toUser(), nil
}
