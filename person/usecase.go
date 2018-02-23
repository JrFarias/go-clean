package person

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UseCaseGetAll return all
func UseCaseGetAll(context *gin.Context) {
	context.JSON(http.StatusOK, RepositoryGetAll())
}
