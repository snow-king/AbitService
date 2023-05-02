package service

import (
	"AbitService/app/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

type Comp struct {
	Name     string
	Date     time.Time
	EducForm string
	Plans    struct {
		PlanOK   int
		PlanPVZ  int
		PlanC    int
		PlanK    int
		PlanI    int
		PlanKvot int
	}
}

func GetGroups(status int) []Comp {
	var groups []models.CompGroups
	var comp []Comp
	err := models.DbAbit.Preload("Potok", func(db *gorm.DB) *gorm.DB {
		return db.Where("potoks.pot_status_id = ?", status)
	}).Preload("Plan", func(db *gorm.DB) *gorm.DB {
		return db.Preload(clause.Associations)
	}).Find(&groups).Error
	if err != nil {
		log.Println(err)
		return []Comp{}
	}
	for _, group := range groups {
		comp = append(comp, Comp{
			Name:     group.Name,
			Date:     group.Potok.DateStart,
			EducForm: group.Plan.StudyForm.Name,
			Plans: struct {
				PlanOK   int
				PlanPVZ  int
				PlanC    int
				PlanK    int
				PlanI    int
				PlanKvot int
			}{
				PlanOK:   group.Plan.Plan1,
				PlanPVZ:  group.Plan.Plan2,
				PlanC:    group.Plan.Plan3,
				PlanK:    group.Plan.Plan4,
				PlanI:    group.Plan.Plan5,
				PlanKvot: group.Plan.Plan6,
			},
		})
	}
	return comp
}
