package response

import (
	"gorm.io/gorm"
)

type UserAuths struct {
	gorm.Model
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status   string `json:"status"`
	RolesID  uint   `json:"role_id"`
	Roles    Roles  `gorm:"foreignKey:RolesID"`
}

type Roles struct {
	gorm.Model
	RoleName string `json:"role_name"`
}

// SigninResponse is the tructure of the response of the user signin
type SigninResponse struct {
	Token    string
	Username string
	ID       int
	Role     string
}
