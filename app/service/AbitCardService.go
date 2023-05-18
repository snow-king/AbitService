package service

import (
	"AbitService/app/models"
	arrayUtils "AbitService/app/utils"
	"golang.org/x/exp/slices"
)

type AbitCardService struct {
	Abit models.AbitCard
}
type AbitCardServiceI interface {
	getSumPoints() Points
	SumLgots() int
}

func NewAbitCardService(abit models.AbitCard) *AbitCardService {
	return &AbitCardService{Abit: abit}
}

type Points struct {
	First      int
	Second     int
	Three      int
	Additional int
	Summary    int
}

// GetSumPoints сумма баллов у поступающего
func (a AbitCardService) GetSumPoints() Points {
	var sum int
	var arr []models.Mark
	var points []models.Mark
	var point Points
	for i := 1; i < 6; i++ {
		arr = arrayUtils.Filter(a.Abit.Marks, func(t models.Mark) bool {
			return t.Subject.Priority == i
		})
		slices.SortFunc(arr, func(a, b models.Mark) bool {
			return a.Mark > b.Mark
		})
		if len(arr) > 0 {
			sum += arr[0].Mark
			points = append(points, arr[0])
		}
	}
	switch len(points) {
	case 1:
		point = Points{
			First:      points[0].Mark,
			Second:     0,
			Three:      0,
			Additional: a.SumLgots(),
			Summary:    points[0].Mark,
		}
	case 2:
		point = Points{
			First:      points[0].Mark,
			Second:     points[1].Mark,
			Three:      0,
			Additional: a.SumLgots(),
			Summary:    points[0].Mark + points[1].Mark,
		}
	case 3:
		point = Points{
			First:      points[0].Mark,
			Second:     points[1].Mark,
			Three:      points[2].Mark,
			Additional: a.SumLgots(),
			Summary:    points[0].Mark + points[1].Mark + points[2].Mark,
		}
	default:
		point = Points{
			First:      0,
			Second:     0,
			Three:      0,
			Additional: 0,
		}
	}
	//if len(points) > 0 {
	//	fmt.Println(point)
	//}
	return point
}

// SumLgots сумма баллов по льготам
func (a AbitCardService) SumLgots() int {
	sum := 0
	for _, soot := range a.Abit.LgotSoot {
		sum += soot.Mark
	}
	if sum > 10 {
		return 10
	} else {
		return sum
	}
}
