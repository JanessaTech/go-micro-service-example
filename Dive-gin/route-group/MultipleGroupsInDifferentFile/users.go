package multiplegroupsindifferentfile

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *routes) addUsersFunction(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	users.GET("/", usersFunction)
}

func usersFunction(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"usersFunction": "usersFunction content"})
}
