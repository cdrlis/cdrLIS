package crud

import (
	"errors"
	ladm "github.com/cdrlis/cdrLIS/LADM"
	"github.com/cdrlis/cdrLIS/LADM/common"
	"github.com/jinzhu/gorm"
	"time"
)

type BfsSpatialUnitMinusCRUD struct {
	DB *gorm.DB
}

func (crud BfsSpatialUnitMinusCRUD) Read(where ...interface{}) (interface{}, error) {
	var bfsSpatialUnit ladm.BfsSpatialUnitMinus
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

func (crud BfsSpatialUnitMinusCRUD) Create(bfsSpatialUnitIn interface{}) (interface{}, error) {
	bfsSpatialUnit := bfsSpatialUnitIn.(ladm.BfsSpatialUnitMinus)
	currentTime := time.Now()
	bfsSpatialUnit.BeginLifespanVersion = currentTime
	bfsSpatialUnit.EndLifespanVersion = nil
	bfsSpatialUnit.SuID = bfsSpatialUnit.Su.SuID.String()
	bfsSpatialUnit.SuBeginLifespanVersion = bfsSpatialUnit.Su.BeginLifespanVersion
	bfsSpatialUnit.BfsID = bfsSpatialUnit.Bfs.BfsID.String()
	bfsSpatialUnit.BfsBeginLifespanVersion = bfsSpatialUnit.Bfs.BeginLifespanVersion
	writer := crud.DB.Set("gorm:save_associations", false).Create(&bfsSpatialUnit)
	if writer.Error != nil{
		return nil, writer.Error
	}
	return &bfsSpatialUnit, nil
}

func (crud BfsSpatialUnitMinusCRUD) ReadAll(where ...interface{}) (interface{}, error) {
	var bfsSpatialUnits []ladm.BfsSpatialUnitMinus
	if crud.DB.Where("endlifespanversion IS NULL").
		Preload("Su", "endlifespanversion IS NULL").
		Preload("Bfs", "endlifespanversion IS NULL").Find(&bfsSpatialUnits).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	return &bfsSpatialUnits, nil
}

func (crud BfsSpatialUnitMinusCRUD) Update(bfsSpatialUnitIn interface{}) (interface{}, error) {
	bfsSpatialUnit := bfsSpatialUnitIn.(*ladm.BfsSpatialUnitMinus)
	return bfsSpatialUnit, nil
}

func (crud BfsSpatialUnitMinusCRUD) Delete(bfsSpatialUnitIn interface{}) error {
	bfsSpatialUnit := bfsSpatialUnitIn.(ladm.BfsSpatialUnitMinus)
	currentTime := time.Now()
	var oldBfsSpatialUnit ladm.BfsSpatialUnitMinus
	reader := crud.DB.Where("su = ? AND "+
		"bfs = ? AND "+
		"endlifespanversion IS NULL", bfsSpatialUnit.Su.SuID.String(), bfsSpatialUnit.Bfs.BfsID.String()).First(&oldBfsSpatialUnit)
	if reader.RowsAffected == 0 {
		return errors.New("Entity not found")
	}
	oldBfsSpatialUnit.EndLifespanVersion = &currentTime
	crud.DB.Set("gorm:save_associations", false).Save(&oldBfsSpatialUnit)
	return nil
}
