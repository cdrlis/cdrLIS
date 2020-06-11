package crud

import (
	"errors"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/jinzhu/gorm"
	"time"
)

type LARequiredRelationshipBAUnitCRUD struct {
	DB *gorm.DB
}

func (crud LARequiredRelationshipBAUnitCRUD) Read(where ...interface{}) (interface{}, error) {
	var relationshipBAUnit ladm.LARequiredRelationshipBAUnit
	if where != nil {
		reader := crud.DB.Where("unit1 = ? AND "+
			"unit2 = ? AND "+
			"endlifespanversion IS NULL", where[0].(common.Oid).String(), where[1].(common.Oid).String()).
			Preload("Unit1", "endlifespanversion IS NULL").
			Preload("Unit2", "endlifespanversion IS NULL").
			First(&relationshipBAUnit)
		if reader.RowsAffected == 0 {
			return nil, errors.New("Entity not found")
		}
		return relationshipBAUnit, nil
	}
	return nil, nil
}

func (crud LARequiredRelationshipBAUnitCRUD) Create(relationshipBAUnitIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	relationshipBAUnit := relationshipBAUnitIn.(ladm.LARequiredRelationshipBAUnit)
	existing := 0
	reader := tx.Model(&ladm.LARequiredRelationshipBAUnit{}).Where("unit1 = ? AND "+
		"unit2 = ? AND "+
		"endlifespanversion IS NULL",
		relationshipBAUnit.Unit1.SuID.String(),
		relationshipBAUnit.Unit2.SuID.String()).
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
	relationshipBAUnit.BeginLifespanVersion = currentTime
	relationshipBAUnit.EndLifespanVersion = nil
	relationshipBAUnit.Unit1ID = relationshipBAUnit.Unit1.SuID.String()
	relationshipBAUnit.Unit1BeginLifespanVersion = relationshipBAUnit.Unit1.BeginLifespanVersion
	relationshipBAUnit.Unit2ID = relationshipBAUnit.Unit2.SuID.String()
	relationshipBAUnit.Unit2BeginLifespanVersion = relationshipBAUnit.Unit2.BeginLifespanVersion
	writer := tx.Set("gorm:save_associations", false).Create(&relationshipBAUnit)
	if writer.Error != nil {
		tx.Rollback()
		return nil, writer.Error
	}
	commit := tx.Commit()
	if commit.Error != nil {
		return nil, commit.Error
	}
	return &relationshipBAUnit, nil
}

func (crud LARequiredRelationshipBAUnitCRUD) ReadAll(where ...interface{}) (interface{}, error) {
	var relationshipBAUnits []ladm.LARequiredRelationshipBAUnit
	if crud.DB.Where("endlifespanversion IS NULL").
		Preload("Unit1", "endlifespanversion IS NULL").
		Preload("Unit2", "endlifespanversion IS NULL").Find(&relationshipBAUnits).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	return &relationshipBAUnits, nil
}

func (crud LARequiredRelationshipBAUnitCRUD) Update(relationshipBAUnitIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	relationshipBAUnit := relationshipBAUnitIn.(*ladm.LARequiredRelationshipBAUnit)
	currentTime := time.Now()
	var oldRelationshipBAUnit ladm.LARequiredRelationshipBAUnit
	reader := tx.Where("unit1 = ? AND "+
		"unit2 = ? AND "+
		"endlifespanversion IS NULL", relationshipBAUnit.Unit1.SuID.String(), relationshipBAUnit.Unit2.SuID.String()).
		First(&oldRelationshipBAUnit)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	oldRelationshipBAUnit.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldRelationshipBAUnit)
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	relationshipBAUnit.BeginLifespanVersion = currentTime
	relationshipBAUnit.EndLifespanVersion = nil
	relationshipBAUnit.Unit1ID = relationshipBAUnit.Unit1.SuID.String()
	relationshipBAUnit.Unit1BeginLifespanVersion = relationshipBAUnit.Unit1.BeginLifespanVersion
	relationshipBAUnit.Unit2ID = relationshipBAUnit.Unit2.SuID.String()
	relationshipBAUnit.Unit2BeginLifespanVersion = relationshipBAUnit.Unit2.BeginLifespanVersion
	writer = tx.Set("gorm:save_associations", false).Create(&relationshipBAUnit)
	if writer.Error != nil {
		tx.Rollback()
		return nil, writer.Error
	}
	commit := tx.Commit()
	if commit.Error != nil {
		return nil, commit.Error
	}
	return relationshipBAUnit, nil
}

func (crud LARequiredRelationshipBAUnitCRUD) Delete(relationshipBAUnitIn interface{}) error {
	tx := crud.DB.Begin()
	relationshipBAUnit := relationshipBAUnitIn.(ladm.LARequiredRelationshipBAUnit)
	currentTime := time.Now()
	var oldRelationshipBAUnit ladm.LARequiredRelationshipBAUnit
	reader := tx.Where("unit1 = ? AND "+
		"unit2 = ? AND "+
		"endlifespanversion IS NULL", relationshipBAUnit.Unit1ID, relationshipBAUnit.Unit2ID).First(&oldRelationshipBAUnit)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	oldRelationshipBAUnit.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldRelationshipBAUnit)
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	commit := tx.Commit()
	if commit.Error != nil {
		return commit.Error
	}
	return nil
}
