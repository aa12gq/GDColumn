package factories

import (
	"GDColumn/app/models/user"
	"GDColumn/pkg/helpers"
	"GDColumn/pkg/snowflake"

	"github.com/bxcodec/faker/v3"
)

func MakeUsers(times int) []user.User {

	var objs []user.User

	// 设置唯一值
	faker.SetGenerateUniqueValues(true)
	_id,_ := snowflake.GetID()
	cid,_ := snowflake.GetID()
	for i := 0; i < times; i++ {
		model := user.User{
			NickName:      faker.Username(),
			UserID:			_id,
			Email:    		faker.Email(),
			Description: 	"this is description",
			Column:     	cid,
			Avatar:  		"this is avatar",
			Phone:    		helpers.RandomNumber(11),
			Password: 		"$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe",
		}
		objs = append(objs, model)
	}

	return objs
}
