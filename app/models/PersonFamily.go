package models

type PersonFamily struct {
	ID        int
	PersonId  int    `gorm:"column:person_id" json:"parent"`
	Person2Id int    `gorm:"column:person2_id" json:"child"`
	Person    Person `gorm:"foreignKey:person2_id;"`
	Status    int
}

func (PersonFamily) TableName() string {
	return "persons_family"
}

type FamilyStatus struct {
	Status int
	Person struct {
		Id       int    `json:"id,omitempty"`
		FullName string `json:"fullName,omitempty"`
		Token    string `json:"token,omitempty"`
	}
}
