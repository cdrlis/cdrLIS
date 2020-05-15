package crud

import (
	"errors"
	ladm "github.com/cdrlis/cdrLIS/LADM"
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
			First(&spatialUnit)
		if reader.RowsAffected == 0 {
			return nil, errors.New("Entity not found")
		}
		return spatialUnit, nil
	}
	return nil, nil
}

func (crud LASpatialUnitCRUD) Create(spatialUnitIn interface{}) (interface{}, error) {
	spatialUnit := spatialUnitIn.(ladm.LASpatialUnit)
	currentTime := time.Now()
	spatialUnit.ID = spatialUnit.SuID.String()
	spatialUnit.BeginLifespanVersion = currentTime
	spatialUnit.EndLifespanVersion = nil
	crud.DB.Set("gorm:save_associations", false).Create(&spatialUnit)
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
	spatialUnit := spatialUnitIn.(*ladm.LASpatialUnit)
	currentTime := time.Now()
	var oldSUnit ladm.LASpatialUnit
	reader := crud.DB.Where("suid = ?::\"Oid\" AND endlifespanversion IS NULL", spatialUnit.SuID).
		First(&oldSUnit)
	if reader.RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	oldSUnit.EndLifespanVersion = &currentTime
	crud.DB.Set("gorm:save_associations", false).Save(&oldSUnit)

	spatialUnit.ID = spatialUnit.SuID.String()
	spatialUnit.BeginLifespanVersion = currentTime
	spatialUnit.EndLifespanVersion = nil
	spatialUnit.LevelID = oldSUnit.LevelID
	spatialUnit.LevelBeginLifespanVersion = oldSUnit.LevelBeginLifespanVersion
	crud.DB.Set("gorm:save_associations", false).Create(&spatialUnit)

	reader = crud.DB.Where("suid = ?::\"Oid\" AND endlifespanversion = ?", spatialUnit.SuID, currentTime).
		Preload("Baunit", "endlifespanversion IS NULL").
		Preload("PlusBfs", "endlifespanversion IS NULL").
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
	for _, plusBfs := range oldSUnit.PlusBfs {
		plusBfs.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&plusBfs)
		plusBfs.BeginLifespanVersion = currentTime
		plusBfs.EndLifespanVersion = nil
		plusBfs.SuBeginLifespanVersion = currentTime
		crud.DB.Set("gorm:save_associations", false).Create(&plusBfs)
	}
	return spatialUnit, nil
}

func (crud LASpatialUnitCRUD) Delete(spatialUnitIn interface{}) error {
	spatialUnit := spatialUnitIn.(ladm.LASpatialUnit)
	currentTime := time.Now()
	var oldSpatialUnit ladm.LASpatialUnit
	reader := crud.DB.Where("suid = ?::\"Oid\" AND endlifespanversion IS NULL", spatialUnit.SuID).First(&oldSpatialUnit)
	if reader.RowsAffected == 0 {
		return errors.New("Entity not found")
	}
	oldSpatialUnit.EndLifespanVersion = &currentTime
	crud.DB.Set("gorm:save_associations", false).Save(&oldSpatialUnit)

	reader = crud.DB.Where("suid = ?::\"Oid\" AND endlifespanversion = ?", spatialUnit.SuID, currentTime).
		Preload("Baunit", "endlifespanversion IS NULL").
		Preload("PlusBfs", "endlifespanversion IS NULL").
		First(&oldSpatialUnit)
	if reader.RowsAffected == 0 {
		return errors.New("Entity not found")
	}
	for _, baUnit := range oldSpatialUnit.Baunit {
		baUnit.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&baUnit)
	}
	for _, plusBfs := range oldSpatialUnit.PlusBfs {
		plusBfs.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&plusBfs)
	}
	return nil
}
