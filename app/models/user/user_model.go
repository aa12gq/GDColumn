package user

import "GDColumn/app/models"

type User struct {
	models.BaseModel

	NickName     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}
