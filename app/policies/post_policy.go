package policies

import (
	"GDColumn/pkg/auth"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func CanModifyPost(c *gin.Context, author uint64) bool {
	return auth.CurrentUID(c) == cast.ToString(author)
}
