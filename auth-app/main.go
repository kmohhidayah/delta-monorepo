package main

import (
	"delta-monorepo/auth-app/api"
	"delta-monorepo/auth-app/models"
	"delta-monorepo/auth-app/user"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main(){
  gotenv.Load()
  models.ConnectToDB()
  models.RegisterModels()

  userRepository := user.NewRepository(models.ConnectToDB())
  userService := user.NewService(userRepository)
  userAPI := api.NewUserAPI(userService) 

  router := gin.Default()
  v1 := router.Group("api/v1")
  v1.POST("/users",userAPI.CreateUser)
  v1.POST("/users/login",userAPI.Login)
  v1.GET("/claims",userAPI.Claims)

  router.Run(":8080")

}
