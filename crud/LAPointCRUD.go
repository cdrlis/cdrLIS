package crud

import (
	"errors"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/jinzhu/gorm"
	"time"
)

type LAPointCRUD struct {
	DB *gorm.DB
}

func (crud LAPointCRUD) Read(where ...interface{}) (interface{}, error) {
	var point ladm.LAPoint
	if where != nil {
		reader := crud.DB.Where("pid = ?::\"Oid\" AND endlifespanversion IS NULL", where).
			First(&point)
		if reader.RowsAffected == 0 {
			return nil, errors.New("Entity not found")
		}
		return point, nil
	}
	return nil, nil
}

func (crud LAPointCRUD) Create(pointIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	point := pointIn.(ladm.LAPoint)
	currentTime := time.Now()
	point.ID = point.PID.String()
	point.BeginLifespanVersion = currentTime
	point.EndLifespanVersion = nil
	writer := tx.Set("gorm:save_associations", false).Create(&point)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return nil, commit.Error
	}
	return &point, nil
}

func (crud LAPointCRUD) ReadAll(where ...interface{}) (interface{}, error) {
	var points []ladm.LAPoint
	if crud.DB.Where("endlifespanversion IS NULL").Find(&points).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	return &points, nil
}

func (crud LAPointCRUD) Update(pointIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	point := pointIn.(*ladm.LAPoint)
	currentTime := time.Now()
	var oldPoint ladm.LAPoint
	reader := tx.Where("pid = ?::\"Oid\" AND endlifespanversion IS NULL", point.PID).
		First(&oldPoint)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	oldPoint.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldPoint)
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	point.ID = point.PID.String()
	point.BeginLifespanVersion = currentTime
	point.EndLifespanVersion = nil
	writer = tx.Set("gorm:save_associations", false).Create(&point)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	reader = tx.Where("pid = ?::\"Oid\" AND endlifespanversion = ?", point.PID, currentTime).
		Preload("SU", "endlifespanversion IS NULL").
		First(&oldPoint)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	for _, bfs := range oldPoint.Bfs {
		bfs.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&bfs)
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		bfs.BeginLifespanVersion = currentTime
		bfs.EndLifespanVersion = nil
		bfs.PointBeginLifespanVersion = currentTime
		writer = tx.Set("gorm:save_associations", false).Create(&bfs)
		if writer.Error != nil{
			tx.Rollback()
			return nil, writer.Error
		}
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return nil, commit.Error
	}
	return point, nil
}

func (crud LAPointCRUD) Delete(pointIn interface{}) error {
	tx := crud.DB.Begin()
	point := pointIn.(ladm.LAPoint)
	currentTime := time.Now()
	var oldPoint ladm.LAPoint
	reader := tx.Where("pid = ?::\"Oid\" AND endlifespanversion IS NULL", point.PID).First(&oldPoint)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	oldPoint.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldPoint)
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	reader = tx.Where("pid = ?::\"Oid\" AND endlifespanversion = ?", point.PID, currentTime).
		Preload("SU", "endlifespanversion IS NULL").
		First(&oldPoint)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	for _, bfs := range oldPoint.Bfs {
		bfs.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&bfs)
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
