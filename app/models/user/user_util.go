package user

import "GDColumn/pkg/database"

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
	database.DB.Where("user_id", idstr).First(&userModel)
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