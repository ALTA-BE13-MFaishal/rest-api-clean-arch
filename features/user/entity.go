package user

import "time"

type Core struct {
	ID        uint
	Name      string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
	Phone     string `validate:"required"`
	Address   string
	Role      string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServiceInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int) error
	Delete(id int) error
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core) (row int, err error)
	GetById(id int) (data Core, err error)
	Update(input Core, id int) (row int, err error)
	Delete(id int) (row int, err error)
}
