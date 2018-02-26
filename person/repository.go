package person

import (
	"github.com/JrFarias/go-clean/common"
)

var person Person
var error common.Error

// RepositoryGetAll call database
func RepositoryGetAll() ([]Person, common.Error) {
	var person []Person
	if err := common.DB.Find(&person).Error; err != nil {
		return person, common.ErrorMessage()["NoCustomer"]
	}

	return person, error
}

// RepositoryGetByID access db and get by id
func RepositoryGetByID(id int) (Person, common.Error) {
	if err := common.DB.Where("id = ?", id).First(&person).Error; err != nil {
		return person, common.ErrorMessage()["CustomerNotFound"]
	}

	return person, error
}

// ResitoryCreate create model
func ResitoryCreate(person Person) (Person, common.Error) {
	if err := common.DB.Create(&person).Error; err != nil {
		return person, error
	}

	return person, error
}
