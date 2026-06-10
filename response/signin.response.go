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
	RolesID  uint
	Roles    Roles `gorm:"foreignKey:RolesID"`
}

type Roles struct {
	gorm.Model
	RoleName string `json:"role_name"`
}
