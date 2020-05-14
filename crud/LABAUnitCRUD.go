package crud

import (
	"errors"
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
			Preload("Su", "endlifespanversion IS NULL").
			Preload("Su.Su", "endlifespanversion IS NULL").
			First(&baUnit)
		if reader.RowsAffected == 0 {
			return nil, errors.New("Entity not found")
		}
		return baUnit, nil
	}
	return nil, nil
}

func (crud LABAUnitCRUD) Create(baUnitIn interface{}) (interface{}, error) {
	baUnit := baUnitIn.(ladm.LABAUnit)
	currentTime := time.Now()
	baUnit.ID = baUnit.UID.String()
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

func (crud LABAUnitCRUD) Update(baUnitIn interface{}) (interface{}, error) {
	baUnit := baUnitIn.(*ladm.LABAUnit)
	currentTime := time.Now()
	var oldBaUnit ladm.LABAUnit
	reader := crud.DB.Where("uid = ?::\"Oid\" AND endlifespanversion IS NULL", baUnit.UID).
		First(&oldBaUnit)
	if reader.RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	oldBaUnit.EndLifespanVersion = &currentTime
	crud.DB.Set("gorm:save_associations", false).Save(&oldBaUnit)

	baUnit.ID = baUnit.UID.String()
	baUnit.BeginLifespanVersion = currentTime
	baUnit.EndLifespanVersion = nil
	crud.DB.Set("gorm:save_associations", false).Create(&baUnit)

	reader = crud.DB.Where("uid = ?::\"Oid\" AND endlifespanversion = ?", baUnit.UID, currentTime).
		Preload("Su", "endlifespanversion IS NULL").
		Preload("Rights", "endlifespanversion IS NULL").
		Preload("Responsibilities", "endlifespanversion IS NULL").
		Preload("Restrictions", "endlifespanversion IS NULL").
		First(&oldBaUnit)
	if reader.RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	for _, su := range oldBaUnit.SU {
		su.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&su)
		su.BeginLifespanVersion = currentTime
		su.EndLifespanVersion = nil
		su.BaUnitBeginLifespanVersion = currentTime
		crud.DB.Set("gorm:save_associations", false).Create(&su)
	}

	for _, right := range oldBaUnit.Rights {
		right.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&right)
		right.BeginLifespanVersion = currentTime
		right.EndLifespanVersion = nil
		right.UnitBeginLifespanVersion = currentTime
		crud.DB.Set("gorm:save_associations", false).Create(&right)
	}

	for _, responsibility := range oldBaUnit.Responsibilities {
		responsibility.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&responsibility)
		responsibility.BeginLifespanVersion = currentTime
		responsibility.EndLifespanVersion = nil
		responsibility.UnitBeginLifespanVersion = currentTime
		crud.DB.Set("gorm:save_associations", false).Create(&responsibility)
	}

	for _, restriction := range oldBaUnit.Restrictions {
		restriction.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&restriction)
		restriction.BeginLifespanVersion = currentTime
		restriction.EndLifespanVersion = nil
		restriction.UnitBeginLifespanVersion = currentTime
		crud.DB.Set("gorm:save_associations", false).Create(&restriction)
	}
	return baUnit, nil
}

func (crud LABAUnitCRUD) Delete(baUnitIn interface{}) error {
	baUnit := baUnitIn.(ladm.LABAUnit)
	currentTime := time.Now()
	var oldBaUnit ladm.LABAUnit
	reader := crud.DB.Where("uid = ?::\"Oid\" AND endlifespanversion IS NULL", baUnit.UID).First(&oldBaUnit)
	if reader.RowsAffected == 0 {
		return errors.New("Entity not found")
	}
	oldBaUnit.EndLifespanVersion = &currentTime
	crud.DB.Set("gorm:save_associations", false).Save(&oldBaUnit)

	reader = crud.DB.Where("uid = ?::\"Oid\" AND endlifespanversion = ?", baUnit.UID, currentTime).
		Preload("Su", "endlifespanversion IS NULL").
		Preload("Rights", "endlifespanversion IS NULL").
		Preload("Responsibilities", "endlifespanversion IS NULL").
		Preload("Restrictions", "endlifespanversion IS NULL").
		First(&oldBaUnit)
	if reader.RowsAffected == 0 {
		return errors.New("Entity not found")
	}
	for _, su := range oldBaUnit.SU {
		su.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&su)
	}

	for _, right := range oldBaUnit.Rights {
		right.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&right)
	}

	for _, responsibility := range oldBaUnit.Responsibilities {
		responsibility.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&responsibility)
	}

	for _, restriction := range oldBaUnit.Restrictions {
		restriction.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&restriction)
	}
	return nil
}
