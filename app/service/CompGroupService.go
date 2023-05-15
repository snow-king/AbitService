package service

import (
	"AbitService/app/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

type Comp struct {
	Id       int
	Name     string
	Date     time.Time
	EducForm string
	Plans    models.ShortPlans
}
type RatingList struct {
	GroupId        int
	Name           string
	TotalPositions int
	Rating         []Rating
}

func GetGroups(status int) []Comp {
	var groups []models.CompGroup
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
			Id:       group.Id,
			Name:     group.Name,
			Date:     group.Potok.DateStart,
			EducForm: group.Plan.StudyForm.Name,
			Plans: models.ShortPlans{
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

func GetList(status int) []RatingList {
	var groups []models.CompGroup
	err := models.DbAbit.Preload("Potok", func(db *gorm.DB) *gorm.DB {
		return db.Where("potoks.pot_status_id = ?", status)
	}).Preload("Plan").Preload("SpecSoots", func(db *gorm.DB) *gorm.DB {
		return db.Preload("AbitCard", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Marks", func(db *gorm.DB) *gorm.DB {
				return db.Preload(clause.Associations)
			})
		})
	}).Find(&groups).Error
	if err != nil {
		log.Println(err)
	}
	var rating []RatingList
	for _, group := range groups {
		list := getListApplicants(group.SpecSoots)
		var plan int
		if group.Plan.Plan1 == 0 {
			plan = group.Plan.Plan1
		} else {
			plan = group.Plan.Plan2
		}
		rating = append(rating, RatingList{
			GroupId:        group.Id,
			Name:           group.SsName,
			Rating:         list,
			TotalPositions: plan,
		})
	}
	return rating
}
