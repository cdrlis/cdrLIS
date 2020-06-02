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
	currentTime := time.Now()
	boundaryFaceString.ID = boundaryFaceString.BfsID.String()
	boundaryFaceString.BeginLifespanVersion = currentTime
	boundaryFaceString.EndLifespanVersion = nil
	writer := crud.DB.Set("gorm:save_associations", false).Create(&boundaryFaceString)
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
	reader := crud.DB.Where("bfsid = ?::\"Oid\" AND endlifespanversion IS NULL", boundaryFaceString.BfsID).
		First(&oldBoundaryFaceString)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	oldBoundaryFaceString.EndLifespanVersion = &currentTime
	writer := crud.DB.Set("gorm:save_associations", false).Save(&oldBoundaryFaceString)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	boundaryFaceString.ID = boundaryFaceString.BfsID.String()
	boundaryFaceString.BeginLifespanVersion = currentTime
	boundaryFaceString.EndLifespanVersion = nil
	writer = crud.DB.Set("gorm:save_associations", false).Create(&boundaryFaceString)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	reader = crud.DB.Where("bfsid = ?::\"Oid\" AND endlifespanversion = ?", boundaryFaceString.BfsID, currentTime).
		Preload("PlusSu", "endlifespanversion IS NULL").
		First(&oldBoundaryFaceString)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	for _, plusSu := range oldBoundaryFaceString.PlusSu {
		plusSu.EndLifespanVersion = &currentTime
		writer = crud.DB.Set("gorm:save_associations", false).Save(&plusSu)
		if reader.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		plusSu.BeginLifespanVersion = currentTime
		plusSu.EndLifespanVersion = nil
		plusSu.BfsBeginLifespanVersion = currentTime
		writer = crud.DB.Set("gorm:save_associations", false).Create(&plusSu)
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
	reader := crud.DB.Where("bfsid = ?::\"Oid\" AND endlifespanversion IS NULL", boundaryFaceString.BfsID).First(&oldBoundaryFaceString)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	oldBoundaryFaceString.EndLifespanVersion = &currentTime
	writer := crud.DB.Set("gorm:save_associations", false).Save(&oldBoundaryFaceString)
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	reader = crud.DB.Where("bfsid = ?::\"Oid\" AND endlifespanversion = ?", boundaryFaceString.BfsID, currentTime).
		Preload("PlusSu", "endlifespanversion IS NULL").
		First(&oldBoundaryFaceString)
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	for _, plusSu := range oldBoundaryFaceString.PlusSu {
		plusSu.EndLifespanVersion = &currentTime
		writer = crud.DB.Set("gorm:save_associations", false).Save(&plusSu)
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
