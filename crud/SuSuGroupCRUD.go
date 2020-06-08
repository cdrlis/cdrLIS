package crud

import (
	"errors"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/jinzhu/gorm"
	"time"
)

type SuSuGroupCRUD struct {
	DB *gorm.DB
}

func (crud SuSuGroupCRUD) Read(where ...interface{}) (interface{}, error) {
	var suSuGroup ladm.SuSuGroup
	if where != nil {
		reader := crud.DB.Where("whole = ? AND "+
			"part = ? AND "+
			"endlifespanversion IS NULL", where[0].(common.Oid).String(), where[1].(common.Oid).String()).
			Preload("Whole", "endlifespanversion IS NULL").
			Preload("Part", "endlifespanversion IS NULL").
			First(&suSuGroup)
		if reader.RowsAffected == 0 {
			return nil, errors.New("Entity not found")
		}
		return suSuGroup, nil
	}
	return nil, nil
}

func (crud SuSuGroupCRUD) Create(suSuGroupIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	suSuGroup := suSuGroupIn.(ladm.SuSuGroup)
	currentTime := time.Now()
	suSuGroup.BeginLifespanVersion = currentTime
	suSuGroup.EndLifespanVersion = nil
	suSuGroup.WholeID = suSuGroup.Whole.SugID.String()
	suSuGroup.WholeBeginLifespanVersion = suSuGroup.Whole.BeginLifespanVersion
	suSuGroup.PartID = suSuGroup.Part.SuID.String()
	suSuGroup.PartBeginLifespanVersion = suSuGroup.Part.BeginLifespanVersion
	writer := tx.Set("gorm:save_associations", false).Create(&suSuGroup)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return nil, commit.Error
	}
	return &suSuGroup, nil
}

func (crud SuSuGroupCRUD) ReadAll(where ...interface{}) (interface{}, error) {
	var bawholeAsParties []ladm.SuSuGroup
	if crud.DB.Where("endlifespanversion IS NULL").
		Preload("Whole", "endlifespanversion IS NULL").
		Preload("Part", "endlifespanversion IS NULL").Find(&bawholeAsParties).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	return &bawholeAsParties, nil
}

func (crud SuSuGroupCRUD) Update(suSuGroupIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	suSuGroup := suSuGroupIn.(*ladm.SuSuGroup)
	currentTime := time.Now()
	var oldSuSuGroup ladm.SuSuGroup
	reader := tx.Where("whole = ? AND "+
		"part = ? AND "+
		"endlifespanversion IS NULL", suSuGroup.Whole.SugID.String(), suSuGroup.Part.SuID.String()).
		First(&oldSuSuGroup)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	oldSuSuGroup.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldSuSuGroup)
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	suSuGroup.BeginLifespanVersion = currentTime
	suSuGroup.EndLifespanVersion = nil
	suSuGroup.WholeID = suSuGroup.Whole.SugID.String()
	suSuGroup.WholeBeginLifespanVersion = suSuGroup.Whole.BeginLifespanVersion
	suSuGroup.PartID = suSuGroup.Part.SuID.String()
	suSuGroup.PartBeginLifespanVersion = suSuGroup.Part.BeginLifespanVersion
	writer = tx.Set("gorm:save_associations", false).Create(&suSuGroup)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return nil, commit.Error
	}
	return suSuGroup, nil
}

func (crud SuSuGroupCRUD) Delete(suSuGroupIn interface{}) error {
	tx := crud.DB.Begin()
	suSuGroup := suSuGroupIn.(ladm.SuSuGroup)
	currentTime := time.Now()
	var oldSuSuGroup ladm.SuSuGroup
	reader := tx.Where("whole = ? AND "+
		"part = ? AND "+
		"endlifespanversion IS NULL", suSuGroup.WholeID, suSuGroup.PartID).First(&oldSuSuGroup)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	oldSuSuGroup.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldSuSuGroup)
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return commit.Error
	}
	return nil
}
