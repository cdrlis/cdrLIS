package crud

import (
	"errors"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/jinzhu/gorm"
	"time"
)

type LASpatialUnitCRUD struct {
	DB *gorm.DB
}

func (crud LASpatialUnitCRUD) Read(where ...interface{}) (interface{}, error) {
	var spatialUnit ladm.LASpatialUnit
	if where != nil {
		reader := crud.DB.Where("suid = ?::\"Oid\" AND endlifespanversion IS NULL", where).
			Preload("Level", "endlifespanversion IS NULL").
			Preload("Baunit", "endlifespanversion IS NULL").
			Preload("Baunit.BaUnit", "endlifespanversion IS NULL").
			Preload("PlusBfs", "endlifespanversion IS NULL").
			Preload("PlusBfs.Bfs", "endlifespanversion IS NULL").
			Preload("MinusBfs", "endlifespanversion IS NULL").
			Preload("MinusBfs.Bfs", "endlifespanversion IS NULL").
			First(&spatialUnit)
		if reader.RowsAffected == 0 {
			return nil, errors.New("Entity not found")
		}
		return spatialUnit, nil
	}
	return nil, nil
}

func (crud LASpatialUnitCRUD) Create(spatialUnitIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	spatialUnit := spatialUnitIn.(ladm.LASpatialUnit)
	currentTime := time.Now()
	spatialUnit.ID = spatialUnit.SuID.String()
	spatialUnit.BeginLifespanVersion = currentTime
	spatialUnit.EndLifespanVersion = nil
	writer := tx.Set("gorm:save_associations", false).Create(&spatialUnit)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return nil, commit.Error
	}
	return &spatialUnit, nil
}

func (crud LASpatialUnitCRUD) ReadAll(where ...interface{}) (interface{}, error) {
	var spatialUnits []ladm.LASpatialUnit
	if crud.DB.Where("endlifespanversion IS NULL").Find(&spatialUnits).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	return &spatialUnits, nil
}

func (crud LASpatialUnitCRUD) Update(spatialUnitIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	spatialUnit := spatialUnitIn.(*ladm.LASpatialUnit)
	currentTime := time.Now()
	var oldSUnit ladm.LASpatialUnit
	reader := tx.Where("suid = ?::\"Oid\" AND endlifespanversion IS NULL", spatialUnit.SuID).
		First(&oldSUnit)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	oldSUnit.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldSUnit)
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	spatialUnit.ID = spatialUnit.SuID.String()
	spatialUnit.BeginLifespanVersion = currentTime
	spatialUnit.EndLifespanVersion = nil
	spatialUnit.LevelID = oldSUnit.LevelID
	spatialUnit.LevelBeginLifespanVersion = oldSUnit.LevelBeginLifespanVersion
	writer = tx.Set("gorm:save_associations", false).Create(&spatialUnit)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	reader = tx.Where("suid = ?::\"Oid\" AND endlifespanversion = ?", spatialUnit.SuID, currentTime).
		Preload("Baunit", "endlifespanversion IS NULL").
		Preload("PlusBfs", "endlifespanversion IS NULL").
		Preload("MinusBfs", "endlifespanversion IS NULL").
		First(&oldSUnit)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	for _, baUnit := range oldSUnit.Baunit {
		baUnit.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&baUnit)
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		baUnit.BeginLifespanVersion = currentTime
		baUnit.EndLifespanVersion = nil
		baUnit.SUBeginLifespanVersion = currentTime
		writer = tx.Set("gorm:save_associations", false).Create(&baUnit)
		if writer.Error != nil{
			tx.Rollback()
			return nil, writer.Error
		}
	}
	for _, plusBfs := range oldSUnit.PlusBfs {
		plusBfs.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&plusBfs)
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		plusBfs.BeginLifespanVersion = currentTime
		plusBfs.EndLifespanVersion = nil
		plusBfs.SuBeginLifespanVersion = currentTime
		writer = tx.Set("gorm:save_associations", false).Create(&plusBfs)
		if writer.Error != nil{
			tx.Rollback()
			return nil, writer.Error
		}
	}
	for _, minusBfs := range oldSUnit.MinusBfs {
		minusBfs.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&minusBfs)
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		minusBfs.BeginLifespanVersion = currentTime
		minusBfs.EndLifespanVersion = nil
		minusBfs.SuBeginLifespanVersion = currentTime
		writer = tx.Set("gorm:save_associations", false).Create(&minusBfs)
		if writer.Error != nil{
			tx.Rollback()
			return nil, writer.Error
		}
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return nil, commit.Error
	}
	return spatialUnit, nil
}

func (crud LASpatialUnitCRUD) Delete(spatialUnitIn interface{}) error {
	tx := crud.DB.Begin()
	spatialUnit := spatialUnitIn.(ladm.LASpatialUnit)
	currentTime := time.Now()
	var oldSpatialUnit ladm.LASpatialUnit
	reader := tx.Where("suid = ?::\"Oid\" AND endlifespanversion IS NULL", spatialUnit.SuID).First(&oldSpatialUnit)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	oldSpatialUnit.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldSpatialUnit)
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	reader = tx.Where("suid = ?::\"Oid\" AND endlifespanversion = ?", spatialUnit.SuID, currentTime).
		Preload("Baunit", "endlifespanversion IS NULL").
		Preload("PlusBfs", "endlifespanversion IS NULL").
		First(&oldSpatialUnit)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	for _, baUnit := range oldSpatialUnit.Baunit {
		baUnit.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&baUnit)
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("Entity not found")
		}
	}
	for _, plusBfs := range oldSpatialUnit.PlusBfs {
		plusBfs.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&plusBfs)
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("Entity not found")
		}
	}
	for _, minusBfs := range oldSpatialUnit.MinusBfs {
		minusBfs.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&minusBfs)
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("Entity not found")
		}
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return commit.Error
	}
	return nil
}
