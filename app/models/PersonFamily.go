package models

type PersonFamily struct {
	ID           int
	PersonId     int    `gorm:"column:person_id" json:"parent"`
	Person2Id    int    `gorm:"column:person2_id" json:"child"`
	Child        Person `gorm:"foreignKey:person2_id;"`
	FamilyTypeId int    `gorm:"column:family_type_id"`
	Status       int
}

func (PersonFamily) TableName() string {
	return "persons_family"
}

type FamilyStatus struct {
	Status int `json:"status,omitempty"`
	Person struct {
		Id       int    `json:"id,omitempty" `
		FullName string `json:"fullName,omitempty" `
		Token    string `json:"token,omitempty"`
	} `json:"person"`
}
