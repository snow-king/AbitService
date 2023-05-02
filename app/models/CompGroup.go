package models

import "time"

type CompGroups struct {
	Id                int       `json:"id"`
	SpecId            int       `json:"specId"`
	PlanId            int       `json:"planId"`
	ContactId         int       `json:"contactId"`
	FormaObucheniyaId int       `json:"formaObucheniyaId"`
	FilialId          int       `json:"filialId,omitempty"`
	Status            int       `json:"status,omitempty"`
	Active            int       `json:"active,omitempty"`
	Name              string    `json:"name,omitempty"`
	SsName            string    `json:"ssName,omitempty"`
	FisId             int       `json:"fisId,omitempty"`
	Type              int       `json:"type,omitempty"`
	SyncSs            int       `json:"syncSs,omitempty"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	PotokId           int       `json:"potokId,omitempty"`
	Potok             Potok     `gorm:"foreignKey:Id;references:PotokId"`
	Plan              Plan      `gorm:"foreignKey:Id;references:PlanId"`
}

func (CompGroups) TableName() string {
	return "comp_groups"
}

type Potok struct {
	Id          int
	PotStatusId int
	Name        string
	DateStart   time.Time
	DateEnd     time.Time
}

func (Potok) TableName() string {
	return "potoks"
}

type Plan struct {
	Id                int
	PotokId           int
	SpecId            int
	FormaObucheniyaId int
	FilialId          int
	Status            int
	Active            int
	Plan1             int `gorm:"plan1"`
	Plan2             int `gorm:"plan2"`
	Plan3             int `gorm:"plan3"`
	Plan4             int `gorm:"plan4"`
	Plan5             int `gorm:"plan5"`
	Plan6             int `gorm:"plan6"`
	Plan7             int `gorm:"plan7"`
	Type              int
	CreatedAt         time.Time
	UpdatedAt         time.Time
	StudyForm         StudyForm `gorm:"foreignKey:Id;references:FormaObucheniyaId"`
}

func (Plan) TableName() string {
	return "plans"
}

type StudyForm struct {
	Id    int
	Name  string
	CYear int
}

func (StudyForm) TableName() string {
	return "studyforms"
}
