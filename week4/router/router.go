package router

import(
	. "firstgin/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter()*gin.Engine{
	router := gin.Default()
	router.GET("/test",FindUser)
	return router
}
