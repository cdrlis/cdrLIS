package crud

import (
	"errors"
	"fmt"
	ladm "github.com/cdrlis/cdrLIS/LADM"
	"github.com/jinzhu/gorm"
	"time"
)

type LASpatialUnitCRUD struct {
	DB *gorm.DB
}

func (crud LASpatialUnitCRUD) Read(where ...interface{}) (interface{}, error) {
	var sUnit ladm.LASpatialUnit
	if where != nil {
		reader := crud.DB.Where("suid = ?::\"Oid\" AND endlifespanversion IS NULL", where).
			Preload("Level").
			Preload("Baunit").
			Preload("Baunit.BaUnit").
			First(&sUnit)
		if reader.RowsAffected == 0 {
			return nil, errors.New("Entity not found")
		}
		return sUnit, nil
	}
	return nil, nil
}

func (crud LASpatialUnitCRUD) Create(partyIn interface{}) (interface{}, error) {
	sUnit := partyIn.(ladm.LASpatialUnit)
	currentTime := time.Now()
	sUnit.ID = fmt.Sprintf("%v-%v", sUnit.SuID.Namespace, sUnit.SuID.LocalID)
	sUnit.BeginLifespanVersion = currentTime
	sUnit.EndLifespanVersion = nil
	crud.DB.Create(&sUnit)
	return &sUnit, nil
}

func (crud LASpatialUnitCRUD) ReadAll(where ...interface{}) (interface{}, error) {
	var sUnits []ladm.LASpatialUnit
	if crud.DB.Where("endlifespanversion IS NULL").Find(&sUnits).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	return &sUnits, nil
}

func (crud LASpatialUnitCRUD) Update(partyIn interface{}) (interface{}, error) {
	sUnit := partyIn.(*ladm.LASpatialUnit)
	currentTime := time.Now()
	var oldSUnit ladm.LASpatialUnit
	if crud.DB.Set("gorm:auto_preload", true).Where("uid = ?::\"Oid\" AND endlifespanversion IS NULL", sUnit.SuID).First(&oldSUnit).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	oldSUnit.EndLifespanVersion = &currentTime
	crud.DB.Save(&oldSUnit)
	sUnit.ID = fmt.Sprintf("%v-%v", sUnit.SuID.Namespace, sUnit.SuID.LocalID)
	sUnit.BeginLifespanVersion = currentTime
	sUnit.EndLifespanVersion = nil
	crud.DB.Create(&sUnit)
	return sUnit, nil
}

func (crud LASpatialUnitCRUD) Delete(partyIn interface{}) error {
	sUnit := partyIn.(ladm.LASpatialUnit)
	currentTime := time.Now()
	var oldSUnit ladm.LASpatialUnit
	if crud.DB.Set("gorm:auto_preload", true).Where("uid = ?::\"Oid\" AND endlifespanversion IS NULL", sUnit.SuID).First(&oldSUnit).RowsAffected == 0 {
		return errors.New("Entity not found")
	}
	oldSUnit.EndLifespanVersion = &currentTime
	crud.DB.Save(&oldSUnit)
	return nil
}
