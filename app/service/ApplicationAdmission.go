package service

import (
	"AbitService/app/models"
	"AbitService/app/utils"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AdmissionList struct {
	Name           string
	Position       int
	TotalPositions int
	Contract       string
}

type ApplicationAdmission struct {
	AbitId int
	Points int
}

func NewApplicationAdmission(abitId int) *ApplicationAdmission {
	return &ApplicationAdmission{AbitId: abitId}
}

func (a ApplicationAdmission) List() []AdmissionList {
	var list []models.SpecSoot
	var admissions []AdmissionList
	models.DbAbit.Preload("Contract").Preload("CompGroups", func(db *gorm.DB) *gorm.DB {
		return db.Preload(clause.Associations)
	}).Where("abit_id", a.AbitId).Find(&list)
	for _, soot := range list {
		admissions = append(admissions, AdmissionList{
			Name:           soot.CompGroups.SsName,
			Position:       a.CalcPosition(soot.CompGroupId),
			TotalPositions: selectPlan(soot.ContractId, soot.CompGroups.Plan),
			Contract:       soot.Contract.Name,
		})
	}
	return admissions
}
func selectPlan(contractId int, plans models.Plan) int {
	switch contractId {
	case 1:
		return plans.Plan2
	case 2:
		return plans.Plan1
	case 3:
		return plans.Plan3
	default:
		return 0
	}
}

type Rating struct {
	AbitId         int
	Points         int
	Prior          int
	AbsoluteRating bool
}

func (a ApplicationAdmission) CalcPosition(compGroupId int) int {
	var soots []models.SpecSoot
	models.DbAbit.Where("comp_group_id = ?", compGroupId).Find(&soots)
	var points []Rating
	points = getListApplicants(soots)
	return slices.IndexFunc(points, func(rating Rating) bool {
		return rating.AbitId == a.AbitId
	})
}
func getListApplicants(soots []models.SpecSoot) []Rating {
	var points []Rating
	for _, soot := range soots {
		points = append(points, Rating{
			AbitId: soot.AbitId,
			Points: getSumPoints(soot.AbitCard),
			Prior:  soot.Raiting,
		})
	}
	slices.SortFunc(points, func(a, b Rating) bool {
		return a.Points > b.Points
	})
	return points
}
func getSumPoints(abit models.AbitCard) int {
	var sum int
	var arr []models.Mark
	for i := 1; i < 6; i++ {
		arr = arrayUtils.Filter(abit.Marks, func(t models.Mark) bool {
			return t.Subject.Priority == i
		})
		slices.SortFunc(arr, func(a, b models.Mark) bool {
			return a.Mark > b.Mark
		})
		if len(arr) > 0 {
			sum += arr[0].Mark
		}
	}
	return sum
}
