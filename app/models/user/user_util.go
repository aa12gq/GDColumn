package user

import (
	"GDColumn/pkg/app"
	"GDColumn/pkg/database"
	"GDColumn/pkg/paginator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

func IsPhoneExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", email).Count(&count)
	return count > 0
}

func GetByPhone(phone string) (userModel User) {
	database.DB.Where("phone = ?", phone).First(&userModel)
	return
}

func GetByMulti(loginID string) (userModel User) {
	database.DB.
		Where("phone = ?", loginID).
		Or("email = ?", loginID).
		Or("nick_name = ?", loginID).
		First(&userModel)
	return
}

func Get(idstr string) (userModel User) {
	database.DB.Preload(clause.Associations).Where("id", idstr).First(&userModel)
	return
}

func GetAvatar(idstr string) (avatarModel Image) {
	database.DB.Where("id", idstr).First(&avatarModel)
	return
}

func GetByEmail(email string) (userModel User) {
	database.DB.Where("email = ?", email).First(&userModel)
	return
}

func All() (users []User) {
	database.DB.Find(&users)
	return
}

func Paginate(c *gin.Context, perPage int) (users []User, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(User{}),
		&users,
		app.V1URL(database.TableName(&User{})),
		perPage,
	)
	return
}