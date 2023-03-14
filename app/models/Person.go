package models

import "time"

type Person struct {
	ID            int       `gorm:"primary_key" json:"id,omitempty"`
	Name          string    `gorm:"column:name1" json:"name,omitempty"`
	LastName      string    `gorm:"column:name2" json:"last-name,omitempty"`
	Surname       string    `gorm:"column:name3" json:"surname,omitempty"`
	Dob           string    `json:"dob,omitempty"`
	SexID         int       `json:"sex-id,omitempty"`
	CountryID     int       `json:"country-id,omitempty"`
	BirLoc        string    `json:"bir-loc,omitempty"`
	Phone         string    `json:"phone,omitempty"`
	Status        int       `json:"status,omitempty"`
	Email         string    `json:"email,omitempty"`
	SNILS         string    `json:"snils,omitempty"`
	INN           string    `json:"inn,omitempty"`
	SMS           int       `json:"sms,omitempty"`
	Address       string    `json:"address,omitempty"`
	AddrGuidFias  string    `json:"addr-guid-fias,omitempty"`
	HouseGuidFias string    `json:"house-guid-fias,omitempty"`
	RoomGuidFias  string    `json:"room-guid-fias,omitempty"`
	Room          string    `json:"room,omitempty"`
	SmallNations  int       `json:"small-nations,omitempty"`
	Compatriot    int       `json:"compatriot,omitempty"`
	Login         string    `json:"login,omitempty"`
	Password      string    `json:"password,omitempty"`
	SertDate      time.Time `json:"sert-date"`
	SertPWD       string    `json:"sert-pwd,omitempty"`
	AdLogin       string    `json:"ad-login,omitempty"`
	Token         string    `json:"token,omitempty"`
	Active        int       `json:"active,omitempty"`
	CreatedAt     time.Time `json:"created-at"`
	UpdatedAt     time.Time `json:"updated-at"`
}

func (Person) TableName() string {
	return "persons"
}
