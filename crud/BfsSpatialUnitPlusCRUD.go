package crud

import (
	"errors"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/jinzhu/gorm"
	"time"
)

type BfsSpatialUnitPlusCRUD struct {
	DB *gorm.DB
}

func (crud BfsSpatialUnitPlusCRUD) Read(where ...interface{}) (interface{}, error) {
	var bfsSpatialUnit ladm.BfsSpatialUnitPlus
	if where != nil {
		reader := crud.DB.Where("su = ? AND "+
			"bfs = ? AND "+
			"endlifespanversion IS NULL", where[0].(common.Oid).String(), where[1].(common.Oid).String()).
			Preload("Su", "endlifespanversion IS NULL").
			Preload("Bfs", "endlifespanversion IS NULL").
			First(&bfsSpatialUnit)
		if reader.RowsAffected == 0 {
			return nil, errors.New("Entity not found")
		}
		return bfsSpatialUnit, nil
	}
	return nil, nil
}

func (crud BfsSpatialUnitPlusCRUD) Create(bfsSpatialUnitIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	bfsSpatialUnit := bfsSpatialUnitIn.(ladm.BfsSpatialUnitPlus)
	currentTime := time.Now()
	bfsSpatialUnit.BeginLifespanVersion = currentTime
	bfsSpatialUnit.EndLifespanVersion = nil
	bfsSpatialUnit.SuID = bfsSpatialUnit.Su.SuID.String()
	bfsSpatialUnit.SuBeginLifespanVersion = bfsSpatialUnit.Su.BeginLifespanVersion
	bfsSpatialUnit.BfsID = bfsSpatialUnit.Bfs.BfsID.String()
	bfsSpatialUnit.BfsBeginLifespanVersion = bfsSpatialUnit.Bfs.BeginLifespanVersion
	writer := crud.DB.Set("gorm:save_associations", false).Create(&bfsSpatialUnit)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return nil, commit.Error
	}
	return &bfsSpatialUnit, nil
}

func (crud BfsSpatialUnitPlusCRUD) ReadAll(where ...interface{}) (interface{}, error) {
	var bfsSpatialUnits []ladm.BfsSpatialUnitPlus
	if crud.DB.Where("endlifespanversion IS NULL").
		Preload("Su", "endlifespanversion IS NULL").
		Preload("Bfs", "endlifespanversion IS NULL").Find(&bfsSpatialUnits).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	return &bfsSpatialUnits, nil
}

func (crud BfsSpatialUnitPlusCRUD) Update(bfsSpatialUnitIn interface{}) (interface{}, error) {
	bfsSpatialUnit := bfsSpatialUnitIn.(*ladm.BfsSpatialUnitPlus)
	return bfsSpatialUnit, nil
}

func (crud BfsSpatialUnitPlusCRUD) Delete(bfsSpatialUnitIn interface{}) error {
	tx := crud.DB.Begin()
	bfsSpatialUnit := bfsSpatialUnitIn.(ladm.BfsSpatialUnitPlus)
	currentTime := time.Now()
	var oldBfsSpatialUnit ladm.BfsSpatialUnitPlus
	reader := crud.DB.Where("su = ? AND "+
		"bfs = ? AND "+
		"endlifespanversion IS NULL", bfsSpatialUnit.Su.SuID.String(), bfsSpatialUnit.Bfs.BfsID.String()).First(&oldBfsSpatialUnit)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	oldBfsSpatialUnit.EndLifespanVersion = &currentTime
	crud.DB.Set("gorm:save_associations", false).Save(&oldBfsSpatialUnit)
	commit := tx.Commit()
	if commit.Error != nil{
		return commit.Error
	}
	return nil
}
