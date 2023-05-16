package multiplegroupsinonefile

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// This demo shows how to create multiple groups
//
//	There is a big issue: as the number of url increases, this file will explode with a big size
//	How can we solve this issue?
func CreateRoutesAndRunServer() {
	var router = gin.Default()
	api := router.Group("/api")
	{
		admin := api.Group("/admin")
		{
			admin.GET("/", adminFunction)
			// more from admin
		}
		users := api.Group("/users")
		{
			users.GET("/", usersFunction)
			// more for users
		}
	}
	//router.Run(":8080")
	router.Run("localhost:8080")
}
func adminFunction(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"adminFunction": "adminFunction content"})
}
func usersFunction(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"usersFunction": "usersFunction content"})
}

// usage:
// url1 : http://localhost:8080/api/users  GET
// url2 : http://localhost:8080/api/admin  GET
