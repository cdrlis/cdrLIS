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
			Preload("Rights", "endlifespanversion IS NULL").
			Preload("Rights.Party", "endlifespanversion IS NULL").
			Preload("Responsibilities", "endlifespanversion IS NULL").
			Preload("Responsibilities.Party", "endlifespanversion IS NULL").
			Preload("Restrictions", "endlifespanversion IS NULL").
			Preload("Restrictions.Party", "endlifespanversion IS NULL").
			Preload("SU", "endlifespanversion IS NULL").
			Preload("SU.SU", "endlifespanversion IS NULL").
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
	crud.DB.Set("gorm:save_associations", false).Create(&baUnit)
	return &baUnit, nil
}

func (crud LABAUnitCRUD) ReadAll(where ...interface{}) (interface{}, error) {
	var baUnits []ladm.LABAUnit
	if crud.DB.Where("endlifespanversion IS NULL").Find(&baUnits).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	return &baUnits, nil
}

func (crud LABAUnitCRUD) Update(baunitIn interface{}) (interface{}, error) {
	baunit := baunitIn.(*ladm.LABAUnit)
	currentTime := time.Now()
	var oldBaunit ladm.LABAUnit
	reader := crud.DB.Where("uid = ?::\"Oid\" AND endlifespanversion IS NULL", baunit.UID).
		First(&oldBaunit)
	if reader.RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	oldBaunit.EndLifespanVersion = &currentTime
	crud.DB.Set("gorm:save_associations", false).Save(&oldBaunit)

	baunit.ID = fmt.Sprintf("%v-%v", baunit.UID.Namespace, baunit.UID.LocalID)
	baunit.BeginLifespanVersion = currentTime
	baunit.EndLifespanVersion = nil
	crud.DB.Set("gorm:save_associations", false).Create(&baunit)

	reader = crud.DB.Where("uid = ?::\"Oid\" AND endlifespanversion = ?", baunit.UID, currentTime).
		Preload("SU", "endlifespanversion IS NULL").
		Preload("Rights", "endlifespanversion IS NULL").
		Preload("Responsibilities", "endlifespanversion IS NULL").
		Preload("Restrictions", "endlifespanversion IS NULL").
		First(&oldBaunit)
	if reader.RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	for _, su := range oldBaunit.SU {
		su.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&su)
		su.BeginLifespanVersion = currentTime
		su.EndLifespanVersion = nil
		su.BaUnitBeginLifespanVersion = currentTime
		crud.DB.Set("gorm:save_associations", false).Create(&su)
	}

	for _, right := range oldBaunit.Rights {
		right.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&right)
		right.BeginLifespanVersion = currentTime
		right.EndLifespanVersion = nil
		right.UnitBeginLifespanVersion = currentTime
		crud.DB.Set("gorm:save_associations", false).Create(&right)
	}

	for _, responsibility := range oldBaunit.Responsibilities {
		responsibility.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&responsibility)
		responsibility.BeginLifespanVersion = currentTime
		responsibility.EndLifespanVersion = nil
		responsibility.UnitBeginLifespanVersion = currentTime
		crud.DB.Set("gorm:save_associations", false).Create(&responsibility)
	}

	for _, restriction := range oldBaunit.Restrictions {
		restriction.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&restriction)
		restriction.BeginLifespanVersion = currentTime
		restriction.EndLifespanVersion = nil
		restriction.UnitBeginLifespanVersion = currentTime
		crud.DB.Set("gorm:save_associations", false).Create(&restriction)
	}
	return baunit, nil
}

func (crud LABAUnitCRUD) Delete(baunitIn interface{}) error {
	baunit := baunitIn.(ladm.LABAUnit)
	currentTime := time.Now()
	var oldBaunit ladm.LABAUnit
	reader := crud.DB.Where("uid = ?::\"Oid\" AND endlifespanversion IS NULL", baunit.UID).First(&oldBaunit)
	if reader.RowsAffected == 0 {
		return errors.New("Entity not found")
	}
	oldBaunit.EndLifespanVersion = &currentTime
	crud.DB.Set("gorm:save_associations", false).Save(&oldBaunit)

	reader = crud.DB.Where("uid = ?::\"Oid\" AND endlifespanversion = ?", baunit.UID, currentTime).
		Preload("SU", "endlifespanversion IS NULL").
		Preload("Rights", "endlifespanversion IS NULL").
		Preload("Responsibilities", "endlifespanversion IS NULL").
		Preload("Restrictions", "endlifespanversion IS NULL").
		First(&oldBaunit)
	if reader.RowsAffected == 0 {
		return errors.New("Entity not found")
	}
	for _, su := range oldBaunit.SU {
		su.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&su)
	}

	for _, right := range oldBaunit.Rights {
		right.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&right)
	}

	for _, responsibility := range oldBaunit.Responsibilities {
		responsibility.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&responsibility)
	}

	for _, restriction := range oldBaunit.Restrictions {
		restriction.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&restriction)
	}
	return nil
}
