package person

import (
	"github.com/JrFarias/go-clean/common"
	"github.com/gin-gonic/gin"
)

//Person model
type Person struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
}

// Router get all routes from domain
func Router(route *gin.Engine) gin.IRoutes {
	return route.Handle("GET", common.Prefix()+"/person", UseCaseGetAll).
		Handle("GET", common.Prefix()+"/person/:id", UseCaseGetByID).
		Handle("POST", common.Prefix()+"/person", UseCaseCreate).
		Handle("PUT", common.Prefix()+"/person/:id", UseCaseUpdate)
}
