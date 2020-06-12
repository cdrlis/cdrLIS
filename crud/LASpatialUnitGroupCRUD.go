package crud

import (
	"errors"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/jinzhu/gorm"
	"time"
)

type LASpatialUnitGroupCRUD struct {
	DB *gorm.DB
}

func (crud LASpatialUnitGroupCRUD) Read(where ...interface{}) (interface{}, error) {
	var suGroup ladm.LASpatialUnitGroup
	if where != nil {
		reader := crud.DB.Where("sugid = ?::\"Oid\" AND endlifespanversion IS NULL", where).
			Preload("SpatialUnits","endlifespanversion IS NULL").
			Preload("SpatialUnits.Part","endlifespanversion IS NULL").
			Preload("SuGroupHierarchySet","endlifespanversion IS NULL").
			Preload("SuGroupHierarchySet.Set","endlifespanversion IS NULL").
			Preload("SuGroupHierarchyElements","endlifespanversion IS NULL").
			Preload("SuGroupHierarchyElements.Element","endlifespanversion IS NULL").
			First(&suGroup)
		if reader.RowsAffected == 0 {
			return nil, errors.New("Entity not found")
		}
		return suGroup, nil
	}
	return nil, nil
}

func (crud LASpatialUnitGroupCRUD) Create(suGroupIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	suGroup := suGroupIn.(ladm.LASpatialUnitGroup)
	existing := 0
	reader := tx.Model(&ladm.LASpatialUnitGroup{}).Where("sugid = ?::\"Oid\" AND endlifespanversion IS NULL", suGroup.SugID).
		Count(&existing)
	if reader.Error != nil{
		tx.Rollback()
		return nil, reader.Error
	}
	if existing != 0 {
		tx.Rollback()
		return nil, errors.New("Entity already exists")
	}
	currentTime := time.Now()
	suGroup.ID = suGroup.SugID.String()
	suGroup.BeginLifespanVersion = currentTime
	suGroup.EndLifespanVersion = nil
	writer := tx.Set("gorm:save_associations", false).Create(&suGroup)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return nil, commit.Error
	}
	return &suGroup, nil
}

func (crud LASpatialUnitGroupCRUD) ReadAll(where ...interface{}) (interface{}, error) {
	var suGroups []ladm.LASpatialUnitGroup
	if crud.DB.Where("endlifespanversion IS NULL").Find(&suGroups).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	return &suGroups, nil
}

func (crud LASpatialUnitGroupCRUD) Update(suGroupIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	suGroup := suGroupIn.(*ladm.LASpatialUnitGroup)
	currentTime := time.Now()
	var oldSuGroup ladm.LASpatialUnitGroup
	reader := tx.Where("sugid = ?::\"Oid\" AND endlifespanversion IS NULL", suGroup.SugID).
		Preload("SpatialUnits","endlifespanversion IS NULL").
		Preload("SuGroupHierarchySet","endlifespanversion IS NULL").
		Preload("SuGroupHierarchyElements","endlifespanversion IS NULL").
		First(&oldSuGroup)
	if reader.Error != nil{
		tx.Rollback()
		return nil, reader.Error
	}
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	oldSuGroup.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldSuGroup)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	suGroup.ID = suGroup.SugID.String()
	suGroup.BeginLifespanVersion = currentTime
	suGroup.EndLifespanVersion = nil
	writer = tx.Set("gorm:save_associations", false).Create(&suGroup)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	for _, spatilUnit := range oldSuGroup.SpatialUnits{
		spatilUnit.EndLifespanVersion = &currentTime
		writer := tx.Set("gorm:save_associations", false).Save(&spatilUnit)
		if writer.Error != nil{
			tx.Rollback()
			return nil, writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		spatilUnit.BeginLifespanVersion = currentTime
		spatilUnit.EndLifespanVersion = nil
		spatilUnit.WholeBeginLifespanVersion = suGroup.BeginLifespanVersion
		writer = tx.Set("gorm:save_associations", false).Create(&spatilUnit)
		if writer.Error != nil{
			tx.Rollback()
			return nil, writer.Error
		}
	}
	if hierarchySet := oldSuGroup.SuGroupHierarchySet; hierarchySet != nil {
		hierarchySet.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(hierarchySet)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		hierarchySet.BeginLifespanVersion = currentTime
		hierarchySet.EndLifespanVersion = nil
		hierarchySet.ElementBeginLifespanVersion = currentTime
		writer = tx.Set("gorm:save_associations", false).Create(hierarchySet)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
	}
	for _, element := range oldSuGroup.SuGroupHierarchyElements {
		element.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&element)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		element.BeginLifespanVersion = currentTime
		element.EndLifespanVersion = nil
		element.SetBeginLifespanVersion = currentTime
		writer = tx.Set("gorm:save_associations", false).Create(&element)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return nil, commit.Error
	}
	return suGroup, nil
}

func (crud LASpatialUnitGroupCRUD) Delete(suGroupIn interface{}) error {
	tx := crud.DB.Begin()
	suGroup := suGroupIn.(ladm.LASpatialUnitGroup)
	currentTime := time.Now()
	var oldSuGroup ladm.LASpatialUnitGroup
	reader := tx.Where("sugid = ?::\"Oid\" AND endlifespanversion IS NULL", suGroup.SugID).
		Preload("SpatialUnits","endlifespanversion IS NULL").
		Preload("SuGroupHierarchySet","endlifespanversion IS NULL").
		Preload("SuGroupHierarchyElements","endlifespanversion IS NULL").
		First(&oldSuGroup)
	if reader.Error != nil{
		tx.Rollback()
		return reader.Error
	}
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	oldSuGroup.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldSuGroup)
	if writer.Error != nil{
		tx.Rollback()
		return writer.Error
	}
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	for _, spatilUnit := range oldSuGroup.SpatialUnits {
		spatilUnit.EndLifespanVersion = &currentTime
		writer := tx.Set("gorm:save_associations", false).Save(&spatilUnit)
		if writer.Error != nil {
			tx.Rollback()
			return writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("Entity not found")
		}
	}
	if hierarchySet := oldSuGroup.SuGroupHierarchySet; hierarchySet != nil {
		hierarchySet.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(hierarchySet)
		if writer.Error != nil {
			tx.Rollback()
			return writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("Entity not found")
		}
	}
	for _, element := range oldSuGroup.SuGroupHierarchyElements {
		element.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&element)
		if writer.Error != nil {
			tx.Rollback()
			return writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("Entity not found")
		}
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return commit.Error
	}
	return nil
}
