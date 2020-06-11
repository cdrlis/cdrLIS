package crud

import (
	"errors"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/jinzhu/gorm"
	"time"
)

type PointBfsCRUD struct {
	DB *gorm.DB
}

func (crud PointBfsCRUD) Read(where ...interface{}) (interface{}, error) {
	var pointBfs ladm.PointBfs
	if where != nil {
		reader := crud.DB.Where("point = ? AND "+
			"bfs = ? AND "+
			"endlifespanversion IS NULL", where[0].(common.Oid).String(), where[1].(common.Oid).String()).
			Preload("Point", "endlifespanversion IS NULL").
			Preload("Bfs", "endlifespanversion IS NULL").
			First(&pointBfs)
		if reader.RowsAffected == 0 {
			return nil, errors.New("Entity not found")
		}
		return pointBfs, nil
	}
	return nil, nil
}

func (crud PointBfsCRUD) Create(pointBfsIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	pointBfs := pointBfsIn.(ladm.PointBfs)
	existing := 0
	reader := tx.Model(&ladm.PointBfs{}).Where("point = ? AND "+
		"bfs = ? AND "+
		"endlifespanversion IS NULL",
		pointBfs.Point.PID.String(),
		pointBfs.Bfs.BfsID.String()).
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
	pointBfs.BeginLifespanVersion = currentTime
	pointBfs.EndLifespanVersion = nil
	pointBfs.PointID = pointBfs.Point.PID.String()
	pointBfs.PointBeginLifespanVersion = pointBfs.Point.BeginLifespanVersion
	pointBfs.BfsID = pointBfs.Bfs.BfsID.String()
	pointBfs.BfsBeginLifespanVersion = pointBfs.Bfs.BeginLifespanVersion
	writer := tx.Set("gorm:save_associations", false).Create(&pointBfs)
	if writer.Error != nil {
		tx.Rollback()
		return nil, writer.Error
	}
	commit := tx.Commit()
	if commit.Error != nil {
		return nil, commit.Error
	}
	return &pointBfs, nil
}

func (crud PointBfsCRUD) ReadAll(where ...interface{}) (interface{}, error) {
	var pointBfss []ladm.PointBfs
	if crud.DB.Where("endlifespanversion IS NULL").
		Preload("Point", "endlifespanversion IS NULL").
		Preload("Bfs", "endlifespanversion IS NULL").Find(&pointBfss).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	return &pointBfss, nil
}

func (crud PointBfsCRUD) Update(pointBfsIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	pointBfs := pointBfsIn.(*ladm.PointBfs)
	currentTime := time.Now()
	var oldRelationshipBAUnit ladm.PointBfs
	reader := tx.Where("point = ? AND "+
		"bfs = ? AND "+
		"endlifespanversion IS NULL", pointBfs.Point.PID.String(), pointBfs.Bfs.BfsID.String()).
		First(&oldRelationshipBAUnit)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	oldRelationshipBAUnit.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldRelationshipBAUnit)
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	pointBfs.BeginLifespanVersion = currentTime
	pointBfs.EndLifespanVersion = nil
	pointBfs.PointID = pointBfs.Point.PID.String()
	pointBfs.PointBeginLifespanVersion = pointBfs.Point.BeginLifespanVersion
	pointBfs.BfsID = pointBfs.Bfs.BfsID.String()
	pointBfs.BfsBeginLifespanVersion = pointBfs.Bfs.BeginLifespanVersion
	writer = tx.Set("gorm:save_associations", false).Create(&pointBfs)
	if writer.Error != nil {
		tx.Rollback()
		return nil, writer.Error
	}
	commit := tx.Commit()
	if commit.Error != nil {
		return nil, commit.Error
	}
	return pointBfs, nil
}

func (crud PointBfsCRUD) Delete(pointBfsIn interface{}) error {
	tx := crud.DB.Begin()
	pointBfs := pointBfsIn.(ladm.PointBfs)
	currentTime := time.Now()
	var oldRelationshipBAUnit ladm.PointBfs
	reader := tx.Where("point = ? AND "+
		"bfs = ? AND "+
		"endlifespanversion IS NULL", pointBfs.PointID, pointBfs.BfsID).First(&oldRelationshipBAUnit)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	oldRelationshipBAUnit.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldRelationshipBAUnit)
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	commit := tx.Commit()
	if commit.Error != nil {
		return commit.Error
	}
	return nil
}
