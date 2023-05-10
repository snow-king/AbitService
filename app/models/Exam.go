package models

import "time"

type Mark struct {
	Id        int
	AbitId    int
	PredmetId int
	SpecId    int
	Mark      int
	CreateAt  time.Time
	UpdateAt  time.Time
	Subject   Subject `gorm:"foreignKey:Id;references:PredmetId"`
}

func (Mark) TableName() string {
	return "marks"
}

type Subject struct {
	Id         int
	Name       string
	Vid        int
	BallMin    int
	EgeId      int
	Priority   int
	Changeable int
	Active     int
}

func (Subject) TableName() string {
	return "predmets"
}
