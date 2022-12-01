package delivery

import "14-api-clean-arch/features/user"

type InsertRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Role     string `json:"role" form:"role"`
}

type UpdateRequest struct {
	ID       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Role     string `json:"role" form:"role"`
}

func toCore(i interface{}) user.Core {
	switch i.(type) {
	case InsertRequest:
		cnv := i.(InsertRequest)
		return user.Core{
			Name:     cnv.Name,
			Email:    cnv.Email,
			Password: cnv.Password,
			Phone:    cnv.Phone,
			Address:  cnv.Address,
			Role:     cnv.Role,
		}

	case UpdateRequest:
		cnv := i.(UpdateRequest)
		return user.Core{
			ID:       cnv.ID,
			Name:     cnv.Name,
			Email:    cnv.Email,
			Password: cnv.Password,
			Phone:    cnv.Phone,
			Address:  cnv.Address,
			Role:     cnv.Role,
		}
	}

	return user.Core{}
}
