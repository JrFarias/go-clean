package person

import (
	"log"

	"github.com/JrFarias/go-clean/common"
)

// RepositoryGetAll call database
func RepositoryGetAll() []Person {
	defer common.DataBase().Close()

	common.DataBase().AutoMigrate(&Person{})
	var person []Person
	if err := common.DataBase().Find(&person).Error; err != nil {
		log.Fatal(err)
	}

	return person
}
