package multiplegroupsindifferentfile

import "github.com/gin-gonic/gin"

type routes struct {
	router *gin.Engine
}

func CreateRoutesAndRunServer() {
	r := routes{
		router: gin.Default(),
	}

	api := r.router.Group("/api")
	r.addAdminFunction(api)
	r.addUsersFunction(api)

	r.router.Run("localhost:8080")
}

// usage:
// url1 : http://localhost:8080/api/users  GET
// url2 : http://localhost:8080/api/admin  GET
