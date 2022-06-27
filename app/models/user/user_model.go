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
	ColumnID  	 string `json:"columnId"`
	AvatarID     string `json:"avatar_id"`
	NickName     string `json:"nickName,omitempty"`
	Description  string `json:"description,omitempty"`
	Email    	 string `json:"email,omitempty"`
	Phone 		 string `json:"-"`
	Password	 string `json:"-"`

	Avatar 	     *Image   `json:"avatar"`
	Column 		 *Column  `json:"-"`

	models.CommonTimestampsField
}

type Image struct {
	ID  string `json:"id"`
	URL string `json:"url"`
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

func (userModel *User) Updates(id,aId,name,des string) (rowsAffected int64) {

	if aId != "" && name != "" && des != ""{
		result := database.DB.Model(&userModel).
					Select("avatar_id","nick_name","description").
					Updates(map[string]interface{}{"avatar_id":aId,"nick_name":name,"description":des})
		return result.RowsAffected
	}else if aId != ""{
		result := database.DB.Model(&userModel).
			Select("avatar_id").
			Updates(map[string]interface{}{"avatar_id":aId})
		return result.RowsAffected
	}else if name != ""{
		result := database.DB.Model(&userModel).
			Select("nick_name").
			Updates(map[string]interface{}{"nick_name":name})
		return result.RowsAffected
	}else if des != "" {
		result := database.DB.Model(&userModel).
			Select("description").
			Updates(map[string]interface{}{"description":des})
		return result.RowsAffected
	}
	return 0
}
