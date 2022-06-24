package user

import (
	"GDColumn/app/models"
	"GDColumn/app/models/column"
	"GDColumn/pkg/database"
	"GDColumn/pkg/hash"
	"github.com/spf13/cast"
)

type User struct {
	ID           string `gorm:"column:id;primaryKey;autoIncrement;" json:"id"`
	ColumnID  	 string `json:"-"`
	AvatarID     string `json:"-"`
	NickName     string `json:"nick_name,omitempty"`
	Description  string `json:"description,omitempty"`
	Email    	 string `json:"email,omitempty"`
	Phone 		 string `json:"-"`
	Password	 string `json:"-"`

	Avatar 	     *Image  `json:"avatar,omitempty"`
	Column 		 *Column  `json:"-"`

	models.CommonTimestampsField
}

type Image struct {
	ID  string `json:"id,omitempty"`
	URL string `json:"url,omitempty"`
}

type Column struct {
	column.Column
}

func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

// GetStringID 获取 ID 的字符串格式
func (userModel *User) GetStringID() string {
	return cast.ToString(userModel.ID)
}

func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}

func (userModel *User) Update() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}

//func (avatarModel *Avatar) Save() (rowsAffected int64) {
//	result := database.DB.Save(&avatarModel)
//	return result.RowsAffected
//}