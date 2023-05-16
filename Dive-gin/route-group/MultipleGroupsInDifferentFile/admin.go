package multiplegroupsindifferentfile

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *routes) addAdminFunction(rg *gin.RouterGroup) {
	admin := rg.Group("/admin")
	admin.GET("/", adminFunction)
}

func adminFunction(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"adminFunction": "adminFunction content"})
}
