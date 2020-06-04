package crud

import (
	"errors"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/jinzhu/gorm"
	"time"
)

type LALevelCRUD struct {
	DB *gorm.DB
}

func (crud LALevelCRUD) Read(where ...interface{}) (interface{}, error) {
	var level ladm.LALevel
	if where != nil {
		reader := crud.DB.Where("lid = ?::\"Oid\" AND endlifespanversion IS NULL", where).
			First(&level)
		if reader.RowsAffected == 0 {
			return nil, errors.New("Entity not found")
		}
		return level, nil
	}
	return nil, nil
}

func (crud LALevelCRUD) Create(levelIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	level := levelIn.(ladm.LALevel)
	currentTime := time.Now()
	level.ID = level.LID.String()
	level.BeginLifespanVersion = currentTime
	level.EndLifespanVersion = nil
	writer := tx.Set("gorm:save_associations", false).Create(&level)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return nil, commit.Error
	}
	return &level, nil
}

func (crud LALevelCRUD) ReadAll(where ...interface{}) (interface{}, error) {
	var levels []ladm.LALevel
	if crud.DB.Where("endlifespanversion IS NULL").Find(&levels).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	return &levels, nil
}

func (crud LALevelCRUD) Update(levelIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	level := levelIn.(*ladm.LALevel)
	currentTime := time.Now()
	var oldLevel ladm.LALevel
	reader := tx.Where("lid = ?::\"Oid\" AND endlifespanversion IS NULL", level.LID).
		First(&oldLevel)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	oldLevel.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldLevel)
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	level.ID = level.LID.String()
	level.BeginLifespanVersion = currentTime
	level.EndLifespanVersion = nil
	writer = tx.Set("gorm:save_associations", false).Create(&level)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	reader = tx.Where("lid = ?::\"Oid\" AND endlifespanversion = ?", level.LID, currentTime).
		Preload("Su", "endlifespanversion IS NULL").
		First(&oldLevel)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	for _, su := range oldLevel.SU {
		su.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&su)
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		su.BeginLifespanVersion = currentTime
		su.EndLifespanVersion = nil
		su.LevelBeginLifespanVersion = currentTime
		writer = tx.Set("gorm:save_associations", false).Create(&su)
		if writer.Error != nil{
			tx.Rollback()
			return nil, writer.Error
		}
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return nil, commit.Error
	}
	return level, nil
}

func (crud LALevelCRUD) Delete(levelIn interface{}) error {
	tx := crud.DB.Begin()
	level := levelIn.(ladm.LALevel)
	currentTime := time.Now()
	var oldLevel ladm.LALevel
	reader := tx.Where("lid = ?::\"Oid\" AND endlifespanversion IS NULL", level.LID).First(&oldLevel)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	oldLevel.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldLevel)
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	reader = tx.Where("lid = ?::\"Oid\" AND endlifespanversion = ?", level.LID, currentTime).
		Preload("Su", "endlifespanversion IS NULL").
		First(&oldLevel)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	for _, su := range oldLevel.SU {
		su.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&su)
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
