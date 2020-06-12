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
	reader = tx.Where("sugid = ?::\"Oid\" AND endlifespanversion = ?", suGroup.SugID, currentTime).
		First(&oldSuGroup)
	if reader.Error != nil{
		tx.Rollback()
		return nil, reader.Error
	}
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
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
	reader := tx.Where("sugid = ?::\"Oid\" AND endlifespanversion IS NULL", suGroup.SugID).First(&oldSuGroup)
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
	reader = tx.Where("sugid = ?::\"Oid\" AND endlifespanversion = ?", suGroup.SugID, currentTime).
		Preload("Su", "endlifespanversion IS NULL").
		First(&oldSuGroup)
	if reader.Error != nil{
		tx.Rollback()
		return reader.Error
	}
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return commit.Error
	}
	return nil
}
