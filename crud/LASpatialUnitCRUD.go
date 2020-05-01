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
			Preload("Level", "endlifespanversion IS NULL").
			Preload("Baunit", "endlifespanversion IS NULL").
			Preload("Baunit.BaUnit", "endlifespanversion IS NULL").
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

func (crud LASpatialUnitCRUD) Update(sunitIn interface{}) (interface{}, error) {
	sUnit := sunitIn.(*ladm.LASpatialUnit)
	currentTime := time.Now()
	var oldSUnit ladm.LASpatialUnit
	reader := crud.DB.Where("suid = ?::\"Oid\" AND endlifespanversion IS NULL", sUnit.SuID).
		First(&oldSUnit)
	if reader.RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	oldSUnit.EndLifespanVersion = &currentTime
	crud.DB.Set("gorm:save_associations", false).Save(&oldSUnit)

	sUnit.ID = fmt.Sprintf("%v-%v", sUnit.SuID.Namespace, sUnit.SuID.LocalID)
	sUnit.BeginLifespanVersion = currentTime
	sUnit.EndLifespanVersion = nil
	sUnit.LevelID = oldSUnit.LevelID
	sUnit.LevelBeginLifespanVersion = oldSUnit.LevelBeginLifespanVersion
	crud.DB.Set("gorm:save_associations", false).Create(&sUnit)

	reader = crud.DB.Where("suid = ?::\"Oid\" AND endlifespanversion = ?", sUnit.SuID, currentTime).
		Preload("Baunit", "endlifespanversion IS NULL").
		First(&oldSUnit)
	if reader.RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	for _, baUnit := range oldSUnit.Baunit {
		baUnit.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&baUnit)
		baUnit.BeginLifespanVersion = currentTime
		baUnit.EndLifespanVersion = nil
		baUnit.SUBeginLifespanVersion = currentTime
		crud.DB.Set("gorm:save_associations", false).Create(&baUnit)
	}
	return sUnit, nil
}

func (crud LASpatialUnitCRUD) Delete(sunitIn interface{}) error {
	baunit := sunitIn.(ladm.LASpatialUnit)
	currentTime := time.Now()
	var oldBaunit ladm.LASpatialUnit
	reader := crud.DB.Where("suid = ?::\"Oid\" AND endlifespanversion IS NULL", baunit.SuID).First(&oldBaunit)
	if reader.RowsAffected == 0 {
		return errors.New("Entity not found")
	}
	oldBaunit.EndLifespanVersion = &currentTime
	crud.DB.Set("gorm:save_associations", false).Save(&oldBaunit)

	reader = crud.DB.Where("suid = ?::\"Oid\" AND endlifespanversion = ?", baunit.SuID, currentTime).
		Preload("Baunit", "endlifespanversion IS NULL").
		First(&oldBaunit)
	if reader.RowsAffected == 0 {
		return errors.New("Entity not found")
	}
	for _, baUnit := range oldBaunit.Baunit {
		baUnit.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&baUnit)
	}
	return nil
}
