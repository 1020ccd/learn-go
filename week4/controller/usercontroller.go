package controller

import(
	model "firstgin/model"
	"net/http"
	"github.com/gin-gonic/gin"
)

func FindUser(c *gin.Context){
	var user model.TestUser
	result := user.FindUser()

	c.JSON(http.StatusOK, gin.H{
		"code":1,
		"msg":"success",
		"data":result,
		"err":"",
	})

}
