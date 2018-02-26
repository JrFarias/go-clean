package person

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/JrFarias/go-clean/common"

	"github.com/gin-gonic/gin"
)

// UseCaseGetAll return all
func UseCaseGetAll(context *gin.Context) {
	persons, err := RepositoryGetAll()

	if err.Code != 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	context.JSON(http.StatusOK, persons)
}

// UseCaseGetByID get by id
func UseCaseGetByID(context *gin.Context) {
	id, _ := strconv.Atoi(context.Params.ByName("id"))
	fmt.Println("id", id)
	person, err := RepositoryGetByID(id)

	if err.Message != "" {
		context.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	context.JSON(http.StatusOK, person)
}

// UseCaseCreate create
func UseCaseCreate(context *gin.Context) {
	var person Person
	context.BindJSON(&person)
	personID, _ := strconv.ParseInt(strconv.FormatUint(person.ID, 10), 10, 64)
	_, err := RepositoryGetByID(int(personID))

	if err.Code == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, common.ErrorMessage()["CustomerAlreadyExists"])
		return
	}

	newPerson, _ := ResitoryCreate(person)
	context.JSON(http.StatusOK, newPerson)
}

// UseCaseUpdate update
func UseCaseUpdate(context *gin.Context) {
	var person Person
	id, _ := strconv.Atoi(context.Params.ByName("id"))
	context.BindJSON(&person)

	oldPerson, err := RepositoryGetByID(id)

	if err.Code != 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}
	person.ID = oldPerson.ID

	context.JSON(http.StatusOK, person)
}
