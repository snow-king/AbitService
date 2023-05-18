package repository

import "AbitService/app/models"

type PersonRepo struct {
}

func (p PersonRepo) Index(id int) models.Person {
	var person models.Person
	models.DbAbit.Where("id", id).First(&person)
	return person
}
func (p PersonRepo) Family(id int) []models.PersonFamily {
	var family []models.PersonFamily
	models.DbAbit.Preload("Child").Where("person_id", id).Find(&family)
	return family
}
