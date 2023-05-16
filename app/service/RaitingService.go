package service

import (
	arrayUtils "AbitService/app/utils"
	"golang.org/x/exp/slices"
)

//go:generate go run github.com/vektra/mockery/v2@latest --name=IRating
type IRating interface {
	ByGroup(id int) (RatingList, error)
	TopPriority(abitId int) int
}

type RatingCalc struct {
	Rating []RatingList
}

func NewRatingCalc(rating []RatingList) *RatingCalc {
	return &RatingCalc{Rating: rating}
}

// ByGroup формирование рейтинга в группе, с учётом абсолютного рейтинга
func (r RatingCalc) ByGroup(id int) (RatingList, error) {
	group := r.Rating[slices.IndexFunc(r.Rating, func(list RatingList) bool {
		return list.GroupId == id
	})]
	for id, people := range group.Rating {
		absolute := r.TopPriority(people.AbitId)
		if group.GroupId == absolute {
			group.Rating[id].AbsoluteRating = true
		}
	}
	return group, nil
}

// TopPriority высчитывание высшего приоритета у поступающего
func (r RatingCalc) TopPriority(abitId int) int {
	groups := arrayUtils.Filter(r.Rating, func(list RatingList) bool {
		return slices.IndexFunc(list.Rating, func(rating Rating) bool {
			return abitId == rating.AbitId
		}) != -1
	})
	slices.SortFunc(groups, func(a, b RatingList) bool {
		aRat := slices.IndexFunc(a.Rating, func(rating Rating) bool {
			return rating.AbitId == abitId
		})
		bRat := slices.IndexFunc(b.Rating, func(rating Rating) bool {
			return rating.AbitId == abitId
		})
		return a.Rating[aRat].Prior > b.Rating[bRat].Prior
	})
	absolute := 0
	for _, list := range groups {
		if list.TotalPositions >= slices.IndexFunc(list.Rating, func(rating Rating) bool {
			return rating.AbitId == abitId
		}) {
			absolute = list.GroupId
			break
		}
	}
	return absolute
}
