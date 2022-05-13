package main

import (
	"GDColumn/bootstrap"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()

	bootstrap.SetupRoute(router)

	err := router.Run(":3000")
	if err != nil {
		fmt.Println(err.Error())
	}
}