package user_domain

type UserRepository interface {
	GetUser(id int64) (*User, error)
	GetAll() ([]User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id int64) error
}
