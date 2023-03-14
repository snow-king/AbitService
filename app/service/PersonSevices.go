package service

import "AbitService/app/models"

type PersonService struct {
	person models.Person
	family []models.PersonFamily
}

func (p *PersonService) Show(id int) models.Person {
	models.DbAbit.Where("id", id).First(&p.person)
	return p.person
}
func (p *PersonService) GetFamily(id int) []models.FamilyStatus {
	models.DbAbit.Where("person_id", id).Find(&p.family)
	var children []models.FamilyStatus
	for _, family := range p.family {
		var child models.Person
		models.DbAbit.Where("id", family.Person2Id).Find(&child)
		children = append(children, models.FamilyStatus{
			Status: family.Status,
			Person: struct {
				Id       int
				FullName string
				Token    string
			}{Id: child.ID, FullName: child.Name + " " + child.LastName + " " + child.Surname, Token: child.Token},
		})
	}
	return children
}
