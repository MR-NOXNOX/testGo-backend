package api

import (
	"/model"
	"/repository"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)
type UserAPI struct {
	UserRepository repository.UserRepository
}
func (api UserAPI) UserListHandler(context *gin.Context) {
	var UsersInfo model.UserInfo
	users, err := api.UserRepository.GetAllProduct()
	if err != nil {
		log.Println("error UserListHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	UsersInfo.Product = users
	context.JSON(http.StatusOK, UsersInfo)
}