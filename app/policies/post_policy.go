package policies

import (
	"GDColumn/pkg/auth"
	"github.com/gin-gonic/gin"
)

func CanModifyPost(c *gin.Context, author string) bool {
	return auth.CurrentUID(c) == author
}
