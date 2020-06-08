package crud

import (
	"errors"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/jinzhu/gorm"
	"time"
)

type LARequiredRelationshipSpatialUnitCRUD struct {
	DB *gorm.DB
}

func (crud LARequiredRelationshipSpatialUnitCRUD) Read(where ...interface{}) (interface{}, error) {
	var relationshipSu ladm.LARequiredRelationshipSpatialUnit
	if where != nil {
		reader := crud.DB.Where("su1 = ? AND "+
			"su2 = ? AND "+
			"endlifespanversion IS NULL", where[0].(common.Oid).String(), where[1].(common.Oid).String()).
			Preload("Su1", "endlifespanversion IS NULL").
			Preload("Su2", "endlifespanversion IS NULL").
			First(&relationshipSu)
		if reader.RowsAffected == 0 {
			return nil, errors.New("Entity not found")
		}
		return relationshipSu, nil
	}
	return nil, nil
}

func (crud LARequiredRelationshipSpatialUnitCRUD) Create(relationshipSuIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	relationshipSu := relationshipSuIn.(ladm.LARequiredRelationshipSpatialUnit)
	currentTime := time.Now()
	relationshipSu.BeginLifespanVersion = currentTime
	relationshipSu.EndLifespanVersion = nil
	relationshipSu.Su1ID = relationshipSu.Su1.SuID.String()
	relationshipSu.Su1BeginLifespanVersion = relationshipSu.Su1.BeginLifespanVersion
	relationshipSu.Su2ID = relationshipSu.Su2.SuID.String()
	relationshipSu.Su2BeginLifespanVersion = relationshipSu.Su2.BeginLifespanVersion
	writer := tx.Set("gorm:save_associations", false).Create(&relationshipSu)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return nil, commit.Error
	}
	return &relationshipSu, nil
}

func (crud LARequiredRelationshipSpatialUnitCRUD) ReadAll(where ...interface{}) (interface{}, error) {
	var relationshipSus []ladm.LARequiredRelationshipSpatialUnit
	if crud.DB.Where("endlifespanversion IS NULL").
		Preload("Su1", "endlifespanversion IS NULL").
		Preload("Su2", "endlifespanversion IS NULL").Find(&relationshipSus).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	return &relationshipSus, nil
}

func (crud LARequiredRelationshipSpatialUnitCRUD) Update(relationshipSuIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	relationshipSu := relationshipSuIn.(*ladm.LARequiredRelationshipSpatialUnit)
	currentTime := time.Now()
	var oldRelationshipSu ladm.LARequiredRelationshipSpatialUnit
	reader := tx.Where("su1 = ? AND "+
		"su2 = ? AND "+
		"endlifespanversion IS NULL", relationshipSu.Su1.SuID.String(), relationshipSu.Su2.SuID.String()).
		First(&oldRelationshipSu)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	oldRelationshipSu.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldRelationshipSu)
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	relationshipSu.BeginLifespanVersion = currentTime
	relationshipSu.EndLifespanVersion = nil
	relationshipSu.Su1ID = relationshipSu.Su1.SuID.String()
	relationshipSu.Su1BeginLifespanVersion = relationshipSu.Su1.BeginLifespanVersion
	relationshipSu.Su2ID = relationshipSu.Su2.SuID.String()
	relationshipSu.Su2BeginLifespanVersion = relationshipSu.Su2.BeginLifespanVersion
	writer = tx.Set("gorm:save_associations", false).Create(&relationshipSu)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return nil, commit.Error
	}
	return relationshipSu, nil
}

func (crud LARequiredRelationshipSpatialUnitCRUD) Delete(relationshipSuIn interface{}) error {
	tx := crud.DB.Begin()
	relationshipSu := relationshipSuIn.(ladm.LARequiredRelationshipSpatialUnit)
	currentTime := time.Now()
	var oldRelationshipSu ladm.LARequiredRelationshipSpatialUnit
	reader := tx.Where("su1 = ? AND "+
		"su2 = ? AND "+
		"endlifespanversion IS NULL", relationshipSu.Su1ID, relationshipSu.Su2ID).First(&oldRelationshipSu)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	oldRelationshipSu.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldRelationshipSu)
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
