package crud

import (
	"errors"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/jinzhu/gorm"
	"time"
)

type BAUnitAsPartyCRUD struct {
	DB *gorm.DB
}

func (crud BAUnitAsPartyCRUD) Read(where ...interface{}) (interface{}, error) {
	var baunitAsParty ladm.BAUnitAsParty
	if where != nil {
		reader := crud.DB.Where("unit = ? AND "+
			"party = ? AND "+
			"endlifespanversion IS NULL", where[0].(common.Oid).String(), where[1].(common.Oid).String()).
			Preload("Unit", "endlifespanversion IS NULL").
			Preload("Party", "endlifespanversion IS NULL").
			First(&baunitAsParty)
		if reader.RowsAffected == 0 {
			return nil, errors.New("Entity not found")
		}
		return baunitAsParty, nil
	}
	return nil, nil
}

func (crud BAUnitAsPartyCRUD) Create(baunitAsPartyIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	baunitAsParty := baunitAsPartyIn.(ladm.BAUnitAsParty)
	currentTime := time.Now()
	baunitAsParty.BeginLifespanVersion = currentTime
	baunitAsParty.EndLifespanVersion = nil
	baunitAsParty.UnitID = baunitAsParty.Unit.UID.String()
	baunitAsParty.UnitBeginLifespanVersion = baunitAsParty.Unit.BeginLifespanVersion
	baunitAsParty.PartyID = baunitAsParty.Party.PID.String()
	baunitAsParty.PartyBeginLifespanVersion = baunitAsParty.Party.BeginLifespanVersion
	writer := tx.Set("gorm:save_associations", false).Create(&baunitAsParty)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return nil, commit.Error
	}
	return &baunitAsParty, nil
}

func (crud BAUnitAsPartyCRUD) ReadAll(where ...interface{}) (interface{}, error) {
	var baunitAsParties []ladm.BAUnitAsParty
	if crud.DB.Where("endlifespanversion IS NULL").
		Preload("Unit", "endlifespanversion IS NULL").
		Preload("Party", "endlifespanversion IS NULL").Find(&baunitAsParties).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	return &baunitAsParties, nil
}

func (crud BAUnitAsPartyCRUD) Update(baunitAsPartyIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	baunitAsParty := baunitAsPartyIn.(*ladm.BAUnitAsParty)
	currentTime := time.Now()
	var oldBaunitAsParty ladm.BAUnitAsParty
	reader := tx.Where("unit = ? AND "+
		"party = ? AND "+
		"endlifespanversion IS NULL", baunitAsParty.Unit.UID.String(), baunitAsParty.Party.PID.String()).
		First(&oldBaunitAsParty)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	oldBaunitAsParty.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldBaunitAsParty)
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	baunitAsParty.BeginLifespanVersion = currentTime
	baunitAsParty.EndLifespanVersion = nil
	baunitAsParty.UnitID = baunitAsParty.Unit.UID.String()
	baunitAsParty.UnitBeginLifespanVersion = baunitAsParty.Unit.BeginLifespanVersion
	baunitAsParty.PartyID = baunitAsParty.Party.PID.String()
	baunitAsParty.PartyBeginLifespanVersion = baunitAsParty.Party.BeginLifespanVersion
	writer = tx.Set("gorm:save_associations", false).Create(&baunitAsParty)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return nil, commit.Error
	}
	return baunitAsParty, nil
}

func (crud BAUnitAsPartyCRUD) Delete(baunitAsPartyIn interface{}) error {
	tx := crud.DB.Begin()
	baunitAsParty := baunitAsPartyIn.(ladm.BAUnitAsParty)
	currentTime := time.Now()
	var oldBaunitAsParty ladm.BAUnitAsParty
	reader := tx.Where("unit = ? AND "+
		"party = ? AND "+
		"endlifespanversion IS NULL", baunitAsParty.UnitID, baunitAsParty.PartyID).First(&oldBaunitAsParty)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	oldBaunitAsParty.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldBaunitAsParty)
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
