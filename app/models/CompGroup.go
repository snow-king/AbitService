package models

import "time"

type CompGroup struct {
	Id                int        `json:"id"`
	SpecId            int        `json:"specId"`
	PlanId            int        `json:"planId"`
	ContactId         int        `json:"contactId"`
	FormaObucheniyaId int        `json:"formaObucheniyaId"`
	FilialId          int        `json:"filialId,omitempty"`
	Status            int        `json:"status,omitempty"`
	Active            int        `json:"active,omitempty"`
	Name              string     `json:"name,omitempty"`
	SsName            string     `json:"ssName,omitempty"`
	FisId             int        `json:"fisId,omitempty"`
	Type              int        `json:"type,omitempty"`
	SyncSs            int        `json:"syncSs,omitempty"`
	CreatedAt         time.Time  `json:"createdAt"`
	UpdatedAt         time.Time  `json:"updatedAt"`
	PotokId           int        `json:"potokId,omitempty"`
	Potok             Potok      `gorm:"foreignKey:Id;references:PotokId"`
	Plan              Plan       `gorm:"foreignKey:Id;references:PlanId"`
	SpecSoots         []SpecSoot `gorm:"foreignKey:CompGroupId;references:Id"`
}

func (CompGroup) TableName() string {
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

type SpecSoot struct {
	Id                 int
	AbitId             int
	ApplicationDate    time.Time
	SpecId             int
	FormId             int
	ContractId         int
	UserId             int
	NapravId           int
	CompGroupId        int
	AbitDocEdu         int
	Original           int
	PassExam           int
	Potok              string
	Professionalism    string
	AbortDate          time.Time
	Raiting            int
	Kurs               string
	RegDate            time.Time
	ReseptionCondition int
	FilialId           int
	CompGroups         CompGroup `gorm:"foreignKey:Id;references:CompGroupId"`
	Contract           Contract  `gorm:"foreignKey:Id;references:ContractId"`
	AbitCard           AbitCard  `gorm:"foreignKey:Id;references:AbitId"`
}

func (SpecSoot) TableName() string {
	return "spec_soot"
}

type Contract struct {
	Id   int
	Name string
}

func (c Contract) TableName() string {
	return "contracts"
}

type ShortPlans struct {
	PlanOK   int
	PlanPVZ  int
	PlanC    int
	PlanK    int
	PlanI    int
	PlanKvot int
}
