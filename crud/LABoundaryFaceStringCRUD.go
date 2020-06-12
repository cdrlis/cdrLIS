package crud

import (
	"errors"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/jinzhu/gorm"
	"time"
)

type LABoundaryFaceStringCRUD struct {
	DB *gorm.DB
}

func (crud LABoundaryFaceStringCRUD) Read(where ...interface{}) (interface{}, error) {
	var boundaryFaceString ladm.LABoundaryFaceString
	if where != nil {
		reader := crud.DB.Where("bfsid = ?::\"Oid\" AND endlifespanversion IS NULL", where).
			Preload("Point", "endlifespanversion IS NULL").
			Preload("Point.Point", "endlifespanversion IS NULL").
			First(&boundaryFaceString)
		if reader.RowsAffected == 0 {
			return nil, errors.New("Entity not found")
		}
		return boundaryFaceString, nil
	}
	return nil, nil
}

func (crud LABoundaryFaceStringCRUD) Create(boundaryFaceStringIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	boundaryFaceString := boundaryFaceStringIn.(ladm.LABoundaryFaceString)
	existing := 0
	reader := tx.Model(&ladm.LABoundaryFaceString{}).Where("bfsid = ?::\"Oid\" AND endlifespanversion IS NULL", boundaryFaceString.BfsID).
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
	boundaryFaceString.ID = boundaryFaceString.BfsID.String()
	boundaryFaceString.BeginLifespanVersion = currentTime
	boundaryFaceString.EndLifespanVersion = nil
	writer := tx.Set("gorm:save_associations", false).Create(&boundaryFaceString)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return nil, commit.Error
	}
	return &boundaryFaceString, nil
}

func (crud LABoundaryFaceStringCRUD) ReadAll(where ...interface{}) (interface{}, error) {
	var boundaryFaceStrings []ladm.LABoundaryFaceString
	if crud.DB.Where("endlifespanversion IS NULL").Find(&boundaryFaceStrings).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	return &boundaryFaceStrings, nil
}

func (crud LABoundaryFaceStringCRUD) Update(boundaryFaceStringIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	boundaryFaceString := boundaryFaceStringIn.(*ladm.LABoundaryFaceString)
	currentTime := time.Now()
	var oldBoundaryFaceString ladm.LABoundaryFaceString
	reader := tx.Where("bfsid = ?::\"Oid\" AND endlifespanversion IS NULL", boundaryFaceString.BfsID).
		Preload("PlusSu", "endlifespanversion IS NULL").
		Preload("MinusSu", "endlifespanversion IS NULL").
		First(&oldBoundaryFaceString)
	if reader.Error != nil{
		tx.Rollback()
		return nil, reader.Error
	}
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	oldBoundaryFaceString.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldBoundaryFaceString)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	boundaryFaceString.ID = boundaryFaceString.BfsID.String()
	boundaryFaceString.BeginLifespanVersion = currentTime
	boundaryFaceString.EndLifespanVersion = nil
	writer = tx.Set("gorm:save_associations", false).Create(&boundaryFaceString)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	for _, plusSu := range oldBoundaryFaceString.PlusSu {
		plusSu.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&plusSu)
		if writer.Error != nil{
			tx.Rollback()
			return nil, writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		plusSu.BeginLifespanVersion = currentTime
		plusSu.EndLifespanVersion = nil
		plusSu.BfsBeginLifespanVersion = currentTime
		writer = tx.Set("gorm:save_associations", false).Create(&plusSu)
		if writer.Error != nil{
			tx.Rollback()
			return nil, writer.Error
		}
	}
	for _, minusSu := range oldBoundaryFaceString.MinusSu {
		minusSu.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&minusSu)
		if writer.Error != nil{
			tx.Rollback()
			return nil, writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		minusSu.BeginLifespanVersion = currentTime
		minusSu.EndLifespanVersion = nil
		minusSu.BfsBeginLifespanVersion = currentTime
		writer = tx.Set("gorm:save_associations", false).Create(&minusSu)
		if writer.Error != nil{
			tx.Rollback()
			return nil, writer.Error
		}
	}
	for _, point := range oldBoundaryFaceString.Point {
		point.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&point)
		if writer.Error != nil{
			tx.Rollback()
			return nil, writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		point.BeginLifespanVersion = currentTime
		point.EndLifespanVersion = nil
		point.BfsBeginLifespanVersion = currentTime
		writer = tx.Set("gorm:save_associations", false).Create(&point)
		if writer.Error != nil{
			tx.Rollback()
			return nil, writer.Error
		}
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return nil, commit.Error
	}
	return boundaryFaceString, nil
}

func (crud LABoundaryFaceStringCRUD) Delete(boundaryFaceStringIn interface{}) error {
	tx := crud.DB.Begin()
	boundaryFaceString := boundaryFaceStringIn.(ladm.LABoundaryFaceString)
	currentTime := time.Now()
	var oldBoundaryFaceString ladm.LABoundaryFaceString
	reader := tx.Where("bfsid = ?::\"Oid\" AND endlifespanversion IS NULL", boundaryFaceString.BfsID).
		Preload("PlusSu", "endlifespanversion IS NULL").
		Preload("MinusSu", "endlifespanversion IS NULL").
		Preload("Point", "endlifespanversion IS NULL").
		First(&oldBoundaryFaceString)
	if reader.Error != nil{
		tx.Rollback()
		return reader.Error
	}
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	oldBoundaryFaceString.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldBoundaryFaceString)
	if writer.Error != nil{
		tx.Rollback()
		return writer.Error
	}
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	reader = tx.Where("bfsid = ?::\"Oid\" AND endlifespanversion = ?", boundaryFaceString.BfsID, currentTime).
		Preload("PlusSu", "endlifespanversion IS NULL").
		Preload("MinusSu", "endlifespanversion IS NULL").
		Preload("Point", "endlifespanversion IS NULL").
		First(&oldBoundaryFaceString)
	if reader.Error != nil{
		tx.Rollback()
		return reader.Error
	}
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	for _, plusSu := range oldBoundaryFaceString.PlusSu {
		plusSu.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&plusSu)
		if writer.Error != nil{
			tx.Rollback()
			return writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("Entity not found")
		}
	}
	for _, minusSu := range oldBoundaryFaceString.MinusSu {
		minusSu.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&minusSu)
		if writer.Error != nil{
			tx.Rollback()
			return writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("Entity not found")
		}
	}
	for _, point := range oldBoundaryFaceString.Point {
		point.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&point)
		if writer.Error != nil{
			tx.Rollback()
			return writer.Error
		}
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
