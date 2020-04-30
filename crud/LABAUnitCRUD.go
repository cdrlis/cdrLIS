package crud

import (
	"errors"
	"fmt"
	ladm "github.com/cdrlis/cdrLIS/LADM"
	"github.com/jinzhu/gorm"
	"time"
)

type LABAUnitCRUD struct {
	DB *gorm.DB
}

func (crud LABAUnitCRUD) Read(where ...interface{}) (interface{}, error) {
	var baUnit ladm.LABAUnit
	if where != nil {
		reader := crud.DB.Where("uid = ?::\"Oid\" AND endlifespanversion IS NULL", where).
			Preload("Rights").
			Preload("Rights.Party").
			Preload("Responsibilities").
			Preload("Responsibilities.Party").
			Preload("Restrictions").
			Preload("Restrictions.Party").
			Preload("SU").
			Preload("SU.SU").
			First(&baUnit)
		if reader.RowsAffected == 0 {
			return nil, errors.New("Entity not found")
		}
		return baUnit, nil
	}
	return nil, nil
}

func (crud LABAUnitCRUD) Create(partyIn interface{}) (interface{}, error) {
	baUnit := partyIn.(ladm.LABAUnit)
	currentTime := time.Now()
	baUnit.ID = fmt.Sprintf("%v-%v", baUnit.UID.Namespace, baUnit.UID.LocalID)
	baUnit.BeginLifespanVersion = currentTime
	baUnit.EndLifespanVersion = nil
	crud.DB.Create(&baUnit)
	return &baUnit, nil
}

func (crud LABAUnitCRUD) ReadAll(where ...interface{}) (interface{}, error) {
	var baUnits []ladm.LABAUnit
	if crud.DB.Where("endlifespanversion IS NULL").Find(&baUnits).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	return &baUnits, nil
}

func (crud LABAUnitCRUD) Update(partyIn interface{}) (interface{}, error) {
	baUnit := partyIn.(*ladm.LABAUnit)
	currentTime := time.Now()
	var oldBaUnit ladm.LAParty
	if crud.DB.Set("gorm:auto_preload", true).Where("uid = ?::\"Oid\" AND endlifespanversion IS NULL", baUnit.UID).First(&oldBaUnit).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	oldBaUnit.EndLifespanVersion = &currentTime
	crud.DB.Save(&oldBaUnit)
	baUnit.ID = fmt.Sprintf("%v-%v", baUnit.UID.Namespace, baUnit.UID.LocalID)
	baUnit.BeginLifespanVersion = currentTime
	baUnit.EndLifespanVersion = nil
	crud.DB.Create(&baUnit)
	return baUnit, nil
}

func (crud LABAUnitCRUD) Delete(partyIn interface{}) error {
	baUnit := partyIn.(ladm.LABAUnit)
	currentTime := time.Now()
	var oldBaUnit ladm.LABAUnit
	if crud.DB.Set("gorm:auto_preload", true).Where("uid = ?::\"Oid\" AND endlifespanversion IS NULL", baUnit.UID).First(&oldBaUnit).RowsAffected == 0 {
		return errors.New("Entity not found")
	}
	oldBaUnit.EndLifespanVersion = &currentTime
	crud.DB.Save(&oldBaUnit)
	return nil
}
