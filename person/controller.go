package person

import (
	"net/http"
	"strconv"

	"github.com/JrFarias/go-clean/common"
	"github.com/gin-gonic/gin"
)

// ControllerGetAll call UseCaseGetAll
func ControllerGetAll(context *gin.Context) {
	persons, err := UseCaseGetAll()

	if err.Code != 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	context.JSON(http.StatusOK, persons)
}

// ControllerGetByID call UseCaseGetByID
func ControllerGetByID(context *gin.Context) {
	id, _ := strconv.Atoi(context.Params.ByName("id"))
	person, err := UseCaseGetByID(id)

	if err.Message != "" {
		context.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	context.JSON(http.StatusOK, person)
}

// ControllerCreate call UseCaseCreate
func ControllerCreate(context *gin.Context) {
	var person Person
	context.BindJSON(&person)
	personID, _ := strconv.ParseInt(strconv.FormatUint(person.ID, 10), 10, 64)
	_, err := UseCaseGetByID(int(personID))

	if err.Code == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, common.ErrorMessage()["CustomerAlreadyExists"])
		return
	}

	newPerson, _ := UseCaseCreate(person)
	context.JSON(http.StatusOK, newPerson)
}

// ControllerUpdate call UseCaseUpdate
func ControllerUpdate(context *gin.Context) {
	var person Person
	id, _ := strconv.Atoi(context.Params.ByName("id"))
	context.BindJSON(&person)

	oldPerson, err := UseCaseGetByID(id)

	if err.Code != 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}
	person.ID = oldPerson.ID
	newPerson, _ := UseCaseUpdate(person)

	context.JSON(http.StatusOK, newPerson)
}
