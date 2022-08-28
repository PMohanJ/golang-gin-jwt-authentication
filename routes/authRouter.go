package routes

import (
	controller "github.com/pmohanj/golang-jwt-project/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/singUP", controller.Singup())
	incomingRoutes.POST("users/login", controller.Login())
}
