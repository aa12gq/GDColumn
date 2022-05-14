package user

import (
	"GDColumn/app/models"
	"GDColumn/pkg/database"
)

type User struct {
	models.BaseModel

	NickName     string `json:"name,omitempty"`
	Email    	 string `json:"-"`
	Phone 		 string `json:"-"`
	Password	 string `json:"-"`

	models.CommonTimestampsField
}

func (userModel *User) Create() {
	database.DB.Create(&userModel)
}
