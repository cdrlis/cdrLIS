package crud

import (
	"errors"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/jinzhu/gorm"
	"time"
)

type LARRRCRUD struct {
	DB *gorm.DB
}

func (crud LARRRCRUD) Read(where ...interface{}) (interface{}, error) {
	var rrr ladm.LARRR
	if where != nil {
		reader := crud.DB.Where("rid = ?::\"Oid\" AND endlifespanversion IS NULL", where).
			Preload("Right", "endlifespanversion IS NULL").
			Preload("Responsibility", "endlifespanversion IS NULL").
			Preload("Restriction", "endlifespanversion IS NULL").
			Preload("Restriction.Mortgage", "endlifespanversion IS NULL").
			Preload("Unit", "endlifespanversion IS NULL").
			Preload("Party", "endlifespanversion IS NULL").
			First(&rrr)
		if reader.RowsAffected == 0 {
			return nil, errors.New("Entity not found")
		}
		return rrr, nil
	}
	return nil, nil
}

func (crud LARRRCRUD) Create(rrrIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	rrr := rrrIn.(ladm.LARRR)
	existing := 0
	reader := tx.Model(&ladm.LARRR{}).Where("rid = ?::\"Oid\" AND endlifespanversion IS NULL", rrr.RID.String()).
		Count(&existing)
	if reader.Error != nil{
		tx.Rollback()
		return nil, reader.Error
	}
	if existing != 0 {
		tx.Rollback()
		return nil, errors.New("Entity already exists")
	}
	rCount := 0
	if rrr.Right != nil {
		rCount += 1
	}
	if rrr.Restriction != nil {
		rCount += 1
	}
	if rrr.Responsibility != nil {
		rCount += 1
	}
	if rCount != 1 {
		tx.Rollback()
		return nil, errors.New("RRR must be one of: Right, Restriction, Responsibility")
	}
	currentTime := time.Now()
	rrr.ID = rrr.RID.String()
	rrr.BeginLifespanVersion = currentTime
	rrr.EndLifespanVersion = nil

	rrr.PartyID = rrr.Party.ID
	rrr.PartyBeginLifespanVersion = rrr.Party.BeginLifespanVersion
	rrr.UnitID = rrr.Unit.ID
	rrr.UnitBeginLifespanVersion = rrr.Unit.BeginLifespanVersion
	writer := tx.Set("gorm:save_associations", false).Create(&rrr)
	if writer.Error != nil {
		tx.Rollback()
		return nil, writer.Error
	}
	if right := rrr.Right; right != nil {
		right.ID = rrr.ID
		right.RID = rrr.RID
		right.BeginLifespanVersion = rrr.BeginLifespanVersion
		right.EndLifespanVersion = rrr.EndLifespanVersion
		writer = tx.Set("gorm:save_associations", false).Create(right)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
	}
	if responsibility := rrr.Responsibility; responsibility != nil {
		responsibility.ID = rrr.ID
		responsibility.RID = rrr.RID
		responsibility.BeginLifespanVersion = rrr.BeginLifespanVersion
		responsibility.EndLifespanVersion = rrr.EndLifespanVersion
		writer = tx.Set("gorm:save_associations", false).Create(responsibility)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
	}
	if restriction := rrr.Restriction; restriction != nil {
		restriction.ID = rrr.ID
		restriction.RID = rrr.RID
		restriction.BeginLifespanVersion = rrr.BeginLifespanVersion
		restriction.EndLifespanVersion = rrr.EndLifespanVersion
		writer = tx.Set("gorm:save_associations", false).Create(restriction)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
		if mortgage := restriction.Mortgage; mortgage != nil{
			mortgage.ID = rrr.ID
			mortgage.RID = rrr.RID
			mortgage.BeginLifespanVersion = rrr.BeginLifespanVersion
			mortgage.EndLifespanVersion = rrr.EndLifespanVersion
			writer = tx.Set("gorm:save_associations", false).Create(mortgage)
			if writer.Error != nil {
				tx.Rollback()
				return nil, writer.Error
			}
		}
	}
	commit := tx.Commit()
	if commit.Error != nil {
		return nil, commit.Error
	}
	return &rrr, nil
}

func (crud LARRRCRUD) ReadAll(where ...interface{}) (interface{}, error) {
	var rrrs []ladm.LARRR
	if crud.DB.Where("endlifespanversion IS NULL").
		Preload("Right", "endlifespanversion IS NULL").
		Preload("Responsibility", "endlifespanversion IS NULL").
		Preload("Restriction", "endlifespanversion IS NULL").
		Preload("Restriction.Mortgage", "endlifespanversion IS NULL").
		Find(&rrrs).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	return &rrrs, nil
}

func (crud LARRRCRUD) Update(rrrIn interface{}) (interface{}, error) {
	rrr := rrrIn.(*ladm.LARRR)
	tx := crud.DB.Begin()
	rCount := 0
	if rrr.Right != nil {
		rCount += 1
	}
	if rrr.Restriction != nil {
		rCount += 1
	}
	if rrr.Responsibility != nil {
		rCount += 1
	}
	if rCount != 1 {
		tx.Rollback()
		return nil, errors.New("RRR must be one of: Right, Restriction, Responsibility")
	}
	currentTime := time.Now()
	var oldRrr ladm.LARRR
	reader := tx.Where("rid = ?::\"Oid\" AND endlifespanversion IS NULL", rrr.RID).
		Preload("Right", "endlifespanversion IS NULL").
		Preload("Responsibility", "endlifespanversion IS NULL").
		Preload("Restriction", "endlifespanversion IS NULL").
		Preload("Restriction.Mortgage", "endlifespanversion IS NULL").
		First(&oldRrr)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	oldRrr.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldRrr)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	rrr.ID = rrr.RID.String()
	rrr.BeginLifespanVersion = currentTime
	rrr.EndLifespanVersion = nil
	writer = tx.Set("gorm:save_associations", false).Create(&rrr)
	if writer.Error != nil {
		tx.Rollback()
		return nil, writer.Error
	}
	reader = tx.Where("lid = ?::\"Oid\" AND endlifespanversion = ?", rrr.RID, currentTime).
		Preload("Right", "endlifespanversion IS NULL").
		Preload("Responsibility", "endlifespanversion IS NULL").
		Preload("Restriction", "endlifespanversion IS NULL").
		Preload("Restriction.Mortgage", "endlifespanversion IS NULL").
		First(&oldRrr)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}

	if right := rrr.Right; right != nil {
		right.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&right)
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		right.BeginLifespanVersion = currentTime
		right.EndLifespanVersion = nil
		writer = tx.Set("gorm:save_associations", false).Create(right)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
	}
	if responsibility := rrr.Responsibility; responsibility != nil {
		responsibility.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&responsibility)
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		responsibility.BeginLifespanVersion = currentTime
		responsibility.EndLifespanVersion = nil
		writer = tx.Set("gorm:save_associations", false).Create(responsibility)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
	}
	if restriction := rrr.Restriction; restriction != nil {
		restriction.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&restriction)
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		restriction.BeginLifespanVersion = currentTime
		restriction.EndLifespanVersion = nil
		writer = tx.Set("gorm:save_associations", false).Create(restriction)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
		if mortgage := restriction.Mortgage; mortgage != nil{
			mortgage.EndLifespanVersion = &currentTime
			writer = tx.Set("gorm:save_associations", false).Save(&mortgage)
			if writer.RowsAffected == 0 {
				tx.Rollback()
				return nil, errors.New("Entity not found")
			}
			mortgage.BeginLifespanVersion = currentTime
			mortgage.EndLifespanVersion = nil
			writer = tx.Set("gorm:save_associations", false).Create(mortgage)
			if writer.Error != nil {
				tx.Rollback()
				return nil, writer.Error
			}
		}
	}
	commit := tx.Commit()
	if commit.Error != nil {
		return nil, commit.Error
	}
	return rrr, nil
}

func (crud LARRRCRUD) Delete(rrrIn interface{}) error {
	rrr := rrrIn.(ladm.LARRR)
	tx := crud.DB.Begin()
	currentTime := time.Now()
	var oldRrr ladm.LARRR
	reader := tx.Where("rid = ?::\"Oid\" AND endlifespanversion IS NULL", rrr.RID).
		Preload("Right", "endlifespanversion IS NULL").
		Preload("Responsibility", "endlifespanversion IS NULL").
		Preload("Restriction", "endlifespanversion IS NULL").
		Preload("Restriction.Mortgage", "endlifespanversion IS NULL").
		First(&oldRrr)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	oldRrr.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldRrr)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}

	if right := oldRrr.Right; right != nil {
		right.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&right)
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("Entity not found")
		}
	}
	if responsibility := oldRrr.Responsibility; responsibility != nil {
		responsibility.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&responsibility)
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("Entity not found")
		}
	}
	if restriction := oldRrr.Restriction; restriction != nil {
		restriction.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&restriction)
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("Entity not found")
		}
		if mortgage := restriction.Mortgage; mortgage != nil{
			mortgage.EndLifespanVersion = &currentTime
			writer = tx.Set("gorm:save_associations", false).Save(&mortgage)
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
