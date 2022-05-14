package user

import (
	"GDColumn/app/models"
	"GDColumn/pkg/database"
	"GDColumn/pkg/hash"
	"github.com/spf13/cast"
)

type User struct {
	models.BaseModel

	UserID 		 uint64 `json:"user_id"`
	NickName     string `json:"name,omitempty"`
	Email    	 string `json:"-"`
	Phone 		 string `json:"-"`
	Password	 string `json:"-"`

	models.CommonTimestampsField
}

func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

// GetStringID 获取 ID 的字符串格式
func (userModel *User) GetStringID() string {
	return cast.ToString(userModel.UserID)
}