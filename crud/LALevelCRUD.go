package crud

import (
	"errors"
	"fmt"
	ladm "github.com/cdrlis/cdrLIS/LADM"
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
	level := levelIn.(ladm.LALevel)
	currentTime := time.Now()
	level.ID = fmt.Sprintf("%v-%v", level.LID.Namespace, level.LID.LocalID)
	level.BeginLifespanVersion = currentTime
	level.EndLifespanVersion = nil
	crud.DB.Set("gorm:save_associations", false).Create(&level)
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
	level := levelIn.(*ladm.LALevel)
	currentTime := time.Now()
	var oldLevel ladm.LALevel
	reader := crud.DB.Where("lid = ?::\"Oid\" AND endlifespanversion IS NULL", level.LID).
		First(&oldLevel)
	if reader.RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	oldLevel.EndLifespanVersion = &currentTime
	crud.DB.Set("gorm:save_associations", false).Save(&oldLevel)

	level.ID = fmt.Sprintf("%v-%v", level.LID.Namespace, level.LID.LocalID)
	level.BeginLifespanVersion = currentTime
	level.EndLifespanVersion = nil
	crud.DB.Set("gorm:save_associations", false).Create(&level)

	reader = crud.DB.Where("lid = ?::\"Oid\" AND endlifespanversion = ?", level.LID, currentTime).
		Preload("SU", "endlifespanversion IS NULL").
		First(&oldLevel)
	if reader.RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	for _, su := range oldLevel.SU {
		su.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&su)
		su.BeginLifespanVersion = currentTime
		su.EndLifespanVersion = nil
		su.LevelBeginLifespanVersion = currentTime
		crud.DB.Set("gorm:save_associations", false).Create(&su)
	}
	return level, nil
}

func (crud LALevelCRUD) Delete(levelIn interface{}) error {
	level := levelIn.(ladm.LALevel)
	currentTime := time.Now()
	var oldLevel ladm.LALevel
	reader := crud.DB.Where("lid = ?::\"Oid\" AND endlifespanversion IS NULL", level.LID).First(&oldLevel)
	if reader.RowsAffected == 0 {
		return errors.New("Entity not found")
	}
	oldLevel.EndLifespanVersion = &currentTime
	crud.DB.Set("gorm:save_associations", false).Save(&oldLevel)

	reader = crud.DB.Where("lid = ?::\"Oid\" AND endlifespanversion = ?", level.LID, currentTime).
		Preload("SU", "endlifespanversion IS NULL").
		First(&oldLevel)
	if reader.RowsAffected == 0 {
		return errors.New("Entity not found")
	}
	for _, su := range oldLevel.SU {
		su.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&su)
	}
	return nil
}
