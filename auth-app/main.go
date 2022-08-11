package main

import (
	"delta-monorepo/auth-app/models"
	"delta-monorepo/auth-app/user"

	"github.com/subosito/gotenv"
)

func main(){
  gotenv.Load()
  models.ConnectToDB()
  models.RegisterModels()

  userRepository := user.NewRepository(models.ConnectToDB())
  user := models.User{
    Name: "Hidh",
    Phone: "0192",
  } 
  userRepository.CreateUser(user)
  }
