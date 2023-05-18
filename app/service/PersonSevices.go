package service

import (
	"AbitService/app/models"
	"AbitService/app/repository"
)

type PersonService struct {
	person models.Person
	family []models.PersonFamily
}

func (p *PersonService) Show(id int) models.Person {
	models.DbAbit.Where("id", id).First(&p.person)
	return p.person
}
func (p *PersonService) GetFamily(id int) []models.FamilyStatus {
	repo := new(repository.PersonRepo)
	p.family = repo.Family(id)
	var children []models.FamilyStatus
	for _, family := range p.family {
		children = append(children, models.FamilyStatus{
			Status: family.Status,
			Person: struct {
				Id       int    `json:"id,omitempty" `
				FullName string `json:"fullName,omitempty" `
				Token    string `json:"token,omitempty"`
			}(struct {
				Id       int
				FullName string
				Token    string
			}{Id: family.Child.ID, FullName: family.Child.Name + " " + family.Child.LastName + " " + family.Child.Surname, Token: family.Child.Token}),
		})
	}
	return children
}
