package policies

import (
	"GDColumn/app/models/post"
	"GDColumn/pkg/auth"
	"github.com/gin-gonic/gin"
)

func CanModifyPost(c *gin.Context, _post post.Post) bool {
	return auth.CurrentUID(c) == _post.UserID
}
