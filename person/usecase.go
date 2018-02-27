package person

import (
	"github.com/JrFarias/go-clean/common"
)

// UseCaseGetAll return all
func UseCaseGetAll() ([]Person, common.Error) {
	return RepositoryGetAll()
}

// UseCaseGetByID get by id
func UseCaseGetByID(id int) (Person, common.Error) {
	return RepositoryGetByID(id)
}

// UseCaseCreate create
func UseCaseCreate(person Person) (Person, common.Error) {
	return ResitoryCreate(person)
}

// UseCaseUpdate update
func UseCaseUpdate(person Person) (Person, common.Error) {
	return RepositoryUpdate(person)
}
