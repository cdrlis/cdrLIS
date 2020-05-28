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
	writer := crud.DB.Set("gorm:save_associations", false).Create(&baUnit)
	if writer.Error != nil{
		return nil, writer.Error
	}
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
		Preload("RRR", "endlifespanversion IS NULL").
		Preload("RRR.Right", "endlifespanversion IS NULL").
		Preload("RRR.Responsibility", "endlifespanversion IS NULL").
		Preload("RRR.Restriction", "endlifespanversion IS NULL").
		Preload("RRR.Restriction.Mortgage", "endlifespanversion IS NULL").
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

	for _, rrr := range oldBaUnit.RRR {
		rrr.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&rrr)
		rrr.BeginLifespanVersion = currentTime
		rrr.EndLifespanVersion = nil
		rrr.PartyBeginLifespanVersion = currentTime
		crud.DB.Set("gorm:save_associations", false).Create(&rrr)
		if right := rrr.Right; right != nil{
			right.EndLifespanVersion = &currentTime
			crud.DB.Set("gorm:save_associations", false).Save(right)
			right.BeginLifespanVersion = currentTime
			right.EndLifespanVersion = nil
			crud.DB.Set("gorm:save_associations", false).Create(right)
		}
		if restriction := rrr.Restriction; restriction != nil{
			restriction.EndLifespanVersion = &currentTime
			crud.DB.Set("gorm:save_associations", false).Save(restriction)
			restriction.BeginLifespanVersion = currentTime
			restriction.EndLifespanVersion = nil
			crud.DB.Set("gorm:save_associations", false).Create(restriction)
			if mortgage := rrr.Restriction.Mortgage; mortgage != nil{
				mortgage.EndLifespanVersion = &currentTime
				crud.DB.Set("gorm:save_associations", false).Save(mortgage)
				mortgage.BeginLifespanVersion = currentTime
				mortgage.EndLifespanVersion = nil
				crud.DB.Set("gorm:save_associations", false).Create(mortgage)
			}
		}
		if responsibility := rrr.Responsibility; responsibility != nil{
			responsibility.EndLifespanVersion = &currentTime
			crud.DB.Set("gorm:save_associations", false).Save(responsibility)
			responsibility.BeginLifespanVersion = currentTime
			responsibility.EndLifespanVersion = nil
			crud.DB.Set("gorm:save_associations", false).Create(responsibility)
		}
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
		Preload("RRR", "endlifespanversion IS NULL").
		Preload("RRR.Right", "endlifespanversion IS NULL").
		Preload("RRR.Responsibility", "endlifespanversion IS NULL").
		Preload("RRR.Restriction", "endlifespanversion IS NULL").
		Preload("RRR.Restriction.Mortgage", "endlifespanversion IS NULL").
		First(&oldBaUnit)
	if reader.RowsAffected == 0 {
		return errors.New("Entity not found")
	}
	for _, su := range oldBaUnit.SU {
		su.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&su)
	}

	for _, rrr := range oldBaUnit.RRR {
		rrr.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&rrr)
		if right := rrr.Right; right != nil{
			right.EndLifespanVersion = &currentTime
			crud.DB.Set("gorm:save_associations", false).Save(right)
		}
		if restriction := rrr.Restriction; restriction != nil{
			restriction.EndLifespanVersion = &currentTime
			crud.DB.Set("gorm:save_associations", false).Save(restriction)
			if mortgage := rrr.Restriction.Mortgage; mortgage != nil{
				mortgage.EndLifespanVersion = &currentTime
				crud.DB.Set("gorm:save_associations", false).Save(mortgage)
			}
		}
		if responsibility := rrr.Responsibility; responsibility != nil{
			responsibility.EndLifespanVersion = &currentTime
			crud.DB.Set("gorm:save_associations", false).Save(responsibility)
		}
	}
	return nil
}
