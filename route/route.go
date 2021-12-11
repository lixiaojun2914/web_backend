package route

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Route() {
	r := gin.Default()
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
	fmt.Println(authorized)
}
