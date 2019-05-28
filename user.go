package rssfeeder

type User struct {
	ID       int
	Login    string
	Password string
}

type UserStorage interface {
	Add(user *User) error
	Update(user *User) error
	Delete(login string) error
	Get(login string) (*User, error)
}
