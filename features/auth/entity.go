package auth

import "time"

type Core struct {
	ID        int
	FullName  string
	Email     string `validate:"required,email"`
	Password  string `valudate:"required"`
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServiceInterface interface {
	Login(input Core) (token string, err error)
}

type RepositoryInterface interface {
	FindUser(email string) (result Core, err error)
}
