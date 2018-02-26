package person

import (
	"log"

	"github.com/JrFarias/go-clean/common"
)

// RepositoryGetAll call database
func RepositoryGetAll() []Person {
	var person []Person
	if err := common.DB.Find(&person).Error; err != nil {
		log.Fatal(err)
	}

	return person
}
