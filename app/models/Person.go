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

type AbitCard struct {
	Id       int
	PersonID int
	Marks    []Mark     `gorm:"foreignKey:AbitId;references:Id"`
	LgotSoot []LgotSoot `gorm:"foreignKey:AbitId;references:Id"`
}

type LgotSoot struct {
	Id         int
	AbitId     int        `json:"abitId,omitempty"`
	SDoc       string     `json:"SDoc,omitempty"`
	NDoc       string     `json:"NDoc,omitempty"`
	LocDoc     string     `json:"locDoc,omitempty"`
	DateDoc    time.Time  `json:"dateDoc"`
	LgotVidId  int        `json:"lgotVidId,omitempty"`
	Confirm    int        `json:"confirm,omitempty"`
	Mark       int        `json:"mark,omitempty"`
	Active     int        `json:"active,omitempty"`
	CreateAt   time.Time  `json:"create_At"`
	UpdateAt   time.Time  `json:"update_At"`
	LgotVidDoc LgotVidDoc `gorm:"foreignKey:Id;references:LgotVidId"`
}

func (LgotSoot) TableName() string {
	return "lgot_soot"
}

type LgotVidDoc struct {
	Id              int       `json:"id,omitempty"`
	LgotId          int       `json:"lgotId,omitempty"`
	DocName         string    `json:"docName,omitempty"`
	GtypeId         int       `json:"gtypeId,omitempty"`
	MarkRecommended int       `json:"markRecommended,omitempty"`
	Active          int       `json:"active,omitempty"`
	SsTypeId        int       `json:"ssTypeId,omitempty"`
	SsCategoryId    int       `json:"ssCategoryId,omitempty"`
	UpdateAt        time.Time `json:"updateAt"`
	Lgot            Lgot      `gorm:"foreignKey:Id;references:LgotId"`
}

func (LgotVidDoc) TableName() string {
	return "lgot_vid_doc"
}

type Lgot struct {
	Id         int
	Name       string
	ParentID   string
	LgotTypeID int
	Active     int
}

func (Lgot) TableName() string {
	return "lgots"
}
