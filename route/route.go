package route
import (
	"/api"
	"/repository"

	"github.com/globalsign/mgo"

	"github.com/gin-gonic/gin"
)
func NewRouteProduct(route *gin.Engine, connectionDB *mgo.Session) {
	userRepository := repository.UserRepositoryMongo{
		ConnectionDB: connectionDB,
	}
	userAPI := api.UserAPI{
		UserRepository: &userRepository,
	}
	route.GET("api/v1/user", userAPI.UserListHandler)
}