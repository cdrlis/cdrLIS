package crud

import (
	"errors"
	"github.com/cdrlis/cdrLIS/ladm"
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
			Preload("RRR", "endlifespanversion IS NULL").
			Preload("RRR.Party", "endlifespanversion IS NULL").
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

func (crud LABAUnitCRUD) Create(baUnitIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	baUnit := baUnitIn.(ladm.LABAUnit)
	existing := 0
	reader := tx.Model(&ladm.LABAUnit{}).Where("uid = ?::\"Oid\" AND endlifespanversion IS NULL", baUnit.UID).
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
	baUnit.ID = baUnit.UID.String()
	baUnit.BeginLifespanVersion = currentTime
	baUnit.EndLifespanVersion = nil
	writer := tx.Set("gorm:save_associations", false).Create(&baUnit)
	if writer.Error != nil {
		tx.Rollback()
		return nil, writer.Error
	}
	commit := tx.Commit()
	if commit.Error != nil {
		return nil, commit.Error
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
	tx := crud.DB.Begin()
	baUnit := baUnitIn.(*ladm.LABAUnit)
	currentTime := time.Now()
	var oldBaUnit ladm.LABAUnit
	reader := tx.Where("uid = ?::\"Oid\" AND endlifespanversion IS NULL", baUnit.UID).
		First(&oldBaUnit)
	if reader.RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	oldBaUnit.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldBaUnit)
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}

	baUnit.ID = baUnit.UID.String()
	baUnit.BeginLifespanVersion = currentTime
	baUnit.EndLifespanVersion = nil
	writer = tx.Set("gorm:save_associations", false).Create(&baUnit)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	reader = tx.Where("uid = ?::\"Oid\" AND endlifespanversion = ?", baUnit.UID, currentTime).
		Preload("Su", "endlifespanversion IS NULL").
		Preload("RRR", "endlifespanversion IS NULL").
		Preload("RRR.Right", "endlifespanversion IS NULL").
		Preload("RRR.Responsibility", "endlifespanversion IS NULL").
		Preload("RRR.Restriction", "endlifespanversion IS NULL").
		Preload("RRR.Restriction.Mortgage", "endlifespanversion IS NULL").
		First(&oldBaUnit)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	for _, su := range oldBaUnit.SU {
		su.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&su)
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		su.BeginLifespanVersion = currentTime
		su.EndLifespanVersion = nil
		su.BaUnitBeginLifespanVersion = currentTime
		writer = tx.Set("gorm:save_associations", false).Create(&su)
		if writer.Error != nil{
			tx.Rollback()
			return nil, writer.Error
		}
	}

	for _, rrr := range oldBaUnit.RRR {
		rrr.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&rrr)
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		rrr.BeginLifespanVersion = currentTime
		rrr.EndLifespanVersion = nil
		rrr.PartyBeginLifespanVersion = currentTime
		writer = tx.Set("gorm:save_associations", false).Create(&rrr)
		if writer.Error != nil{
			tx.Rollback()
			return nil, writer.Error
		}
		if right := rrr.Right; right != nil {
			right.EndLifespanVersion = &currentTime
			writer = tx.Set("gorm:save_associations", false).Save(right)
			if writer.RowsAffected == 0 {
				tx.Rollback()
				return nil, errors.New("Entity not found")
			}
			right.BeginLifespanVersion = currentTime
			right.EndLifespanVersion = nil
			writer = tx.Set("gorm:save_associations", false).Create(right)
			if writer.Error != nil{
				tx.Rollback()
				return nil, writer.Error
			}
		}
		if restriction := rrr.Restriction; restriction != nil {
			restriction.EndLifespanVersion = &currentTime
			writer = tx.Set("gorm:save_associations", false).Save(restriction)
			if writer.RowsAffected == 0 {
				tx.Rollback()
				return nil, errors.New("Entity not found")
			}
			restriction.BeginLifespanVersion = currentTime
			restriction.EndLifespanVersion = nil
			writer = tx.Set("gorm:save_associations", false).Create(restriction)
			if writer.Error != nil{
				tx.Rollback()
				return nil, writer.Error
			}
			if mortgage := rrr.Restriction.Mortgage; mortgage != nil {
				mortgage.EndLifespanVersion = &currentTime
				writer = tx.Set("gorm:save_associations", false).Save(mortgage)
				if writer.RowsAffected == 0 {
					tx.Rollback()
					return nil, errors.New("Entity not found")
				}
				mortgage.BeginLifespanVersion = currentTime
				mortgage.EndLifespanVersion = nil
				writer = tx.Set("gorm:save_associations", false).Create(mortgage)
				if writer.Error != nil{
					tx.Rollback()
					return nil, writer.Error
				}
			}
		}
		if responsibility := rrr.Responsibility; responsibility != nil {
			responsibility.EndLifespanVersion = &currentTime
			writer = tx.Set("gorm:save_associations", false).Save(responsibility)
			if writer.RowsAffected == 0 {
				tx.Rollback()
				return nil, errors.New("Entity not found")
			}
			responsibility.BeginLifespanVersion = currentTime
			responsibility.EndLifespanVersion = nil
			writer = tx.Set("gorm:save_associations", false).Create(responsibility)
			if writer.Error != nil{
				tx.Rollback()
				return nil, writer.Error
			}
		}
	}
	commit := tx.Commit()
	if commit.Error != nil {
		return nil, commit.Error
	}
	return baUnit, nil
}

func (crud LABAUnitCRUD) Delete(baUnitIn interface{}) error {
	tx := crud.DB.Begin()
	baUnit := baUnitIn.(ladm.LABAUnit)
	currentTime := time.Now()
	var oldBaUnit ladm.LABAUnit
	reader := tx.Where("uid = ?::\"Oid\" AND endlifespanversion IS NULL", baUnit.UID).First(&oldBaUnit)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	oldBaUnit.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldBaUnit)
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}

	reader = tx.Where("uid = ?::\"Oid\" AND endlifespanversion = ?", baUnit.UID, currentTime).
		Preload("Su", "endlifespanversion IS NULL").
		Preload("RRR", "endlifespanversion IS NULL").
		Preload("RRR.Right", "endlifespanversion IS NULL").
		Preload("RRR.Responsibility", "endlifespanversion IS NULL").
		Preload("RRR.Restriction", "endlifespanversion IS NULL").
		Preload("RRR.Restriction.Mortgage", "endlifespanversion IS NULL").
		First(&oldBaUnit)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	for _, su := range oldBaUnit.SU {
		su.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&su)
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("Entity not found")
		}
	}

	for _, rrr := range oldBaUnit.RRR {
		rrr.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&rrr)
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("Entity not found")
		}
		if right := rrr.Right; right != nil {
			right.EndLifespanVersion = &currentTime
			writer = tx.Set("gorm:save_associations", false).Save(right)
			if writer.RowsAffected == 0 {
				tx.Rollback()
				return errors.New("Entity not found")
			}
		}
		if restriction := rrr.Restriction; restriction != nil {
			restriction.EndLifespanVersion = &currentTime
			writer = tx.Set("gorm:save_associations", false).Save(restriction)
			if writer.RowsAffected == 0 {
				tx.Rollback()
				return errors.New("Entity not found")
			}
			if mortgage := rrr.Restriction.Mortgage; mortgage != nil {
				mortgage.EndLifespanVersion = &currentTime
				writer = tx.Set("gorm:save_associations", false).Save(mortgage)
				if writer.RowsAffected == 0 {
					tx.Rollback()
					return errors.New("Entity not found")
				}
			}
		}
		if responsibility := rrr.Responsibility; responsibility != nil {
			responsibility.EndLifespanVersion = &currentTime
			writer = tx.Set("gorm:save_associations", false).Save(responsibility)
			if writer.RowsAffected == 0 {
				tx.Rollback()
				return errors.New("Entity not found")
			}
		}
	}
	commit := tx.Commit()
	if commit.Error != nil {
		return commit.Error
	}
	return nil
}
