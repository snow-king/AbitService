package service

import (
	"AbitService/app/models"
	"AbitService/app/repository"
	"golang.org/x/exp/slices"
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
	var admissions []AdmissionList
	list := repository.SpecSootsByAbit(a.AbitId)
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
	Points         Points
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
		abit := NewAbitCardService(soot.AbitCard)
		points = append(points, Rating{
			AbitId: soot.AbitId,
			Points: abit.GetSumPoints(),
			Prior:  soot.Raiting,
		})
	}
	slices.SortFunc(points, func(a, b Rating) bool {
		return a.Points.Summary > b.Points.Summary
	})
	return points
}
