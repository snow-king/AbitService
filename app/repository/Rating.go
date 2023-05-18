package repository

import (
	"AbitService/app/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CompGroups список поступащих групп по статусам
func CompGroups(status int) ([]models.CompGroup, error) {
	var groups []models.CompGroup
	err := models.DbAbit.Preload("Potok", "pot_status_id = ?", status).Preload("Plan").Preload("SpecSoots", func(db *gorm.DB) *gorm.DB {
		return db.Preload("AbitCard", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Marks", func(db *gorm.DB) *gorm.DB {
				return db.Preload(clause.Associations)
			}).Preload("LgotSoot").Preload("LgotSoot.LgotVidDoc.Lgot", "parent_id = ?", 16)
		})
	}).Find(&groups).Error
	return groups, err
}

// SpecSootsByAbit - получение заявлений аббитуриента
func SpecSootsByAbit(abitId int) []models.SpecSoot {
	var list []models.SpecSoot
	models.DbAbit.Preload("Contract").Preload("CompGroups", func(db *gorm.DB) *gorm.DB {
		return db.Preload(clause.Associations)
	}).Where("abit_id", abitId).Find(&list)
	return list
}
