package service

import (
	"AbitService/app/models"
)

type Family struct {
	Parent int
}

func NewFamily(parent int) *Family {
	return &Family{Parent: parent}
}

func (f Family) GetAccess(token string) error {
	var person models.Person
	err := models.DbAbit.Where("token = ?", token).First(&person).Error
	if err != nil {
		return err
	}
	family := models.PersonFamily{
		PersonId:     f.Parent,
		Person2Id:    person.ID,
		Status:       0,
		FamilyTypeId: 1,
	}
	models.DbAbit.Where(family).FirstOrCreate(&family)
	return nil
}
