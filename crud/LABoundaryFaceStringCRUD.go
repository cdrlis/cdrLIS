package crud

import (
	"errors"
	ladm "github.com/cdrlis/cdrLIS/LADM"
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
	boundaryFaceString := boundaryFaceStringIn.(ladm.LABoundaryFaceString)
	currentTime := time.Now()
	boundaryFaceString.ID = boundaryFaceString.BfsID.String()
	boundaryFaceString.BeginLifespanVersion = currentTime
	boundaryFaceString.EndLifespanVersion = nil
	writer := crud.DB.Set("gorm:save_associations", false).Create(&boundaryFaceString)
	if writer.Error != nil{
		return nil, writer.Error
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
	boundaryFaceString := boundaryFaceStringIn.(*ladm.LABoundaryFaceString)
	currentTime := time.Now()
	var oldBoundaryFaceString ladm.LABoundaryFaceString
	reader := crud.DB.Where("bfsid = ?::\"Oid\" AND endlifespanversion IS NULL", boundaryFaceString.BfsID).
		First(&oldBoundaryFaceString)
	if reader.RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	oldBoundaryFaceString.EndLifespanVersion = &currentTime
	crud.DB.Set("gorm:save_associations", false).Save(&oldBoundaryFaceString)

	boundaryFaceString.ID = boundaryFaceString.BfsID.String()
	boundaryFaceString.BeginLifespanVersion = currentTime
	boundaryFaceString.EndLifespanVersion = nil
	crud.DB.Set("gorm:save_associations", false).Create(&boundaryFaceString)

	reader = crud.DB.Where("bfsid = ?::\"Oid\" AND endlifespanversion = ?", boundaryFaceString.BfsID, currentTime).
		Preload("PlusSu", "endlifespanversion IS NULL").
		First(&oldBoundaryFaceString)
	if reader.RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	for _, plusSu := range oldBoundaryFaceString.PlusSu {
		plusSu.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&plusSu)
		plusSu.BeginLifespanVersion = currentTime
		plusSu.EndLifespanVersion = nil
		plusSu.BfsBeginLifespanVersion = currentTime
		crud.DB.Set("gorm:save_associations", false).Create(&plusSu)
	}
	return boundaryFaceString, nil
}

func (crud LABoundaryFaceStringCRUD) Delete(boundaryFaceStringIn interface{}) error {
	boundaryFaceString := boundaryFaceStringIn.(ladm.LABoundaryFaceString)
	currentTime := time.Now()
	var oldBoundaryFaceString ladm.LABoundaryFaceString
	reader := crud.DB.Where("bfsid = ?::\"Oid\" AND endlifespanversion IS NULL", boundaryFaceString.BfsID).First(&oldBoundaryFaceString)
	if reader.RowsAffected == 0 {
		return errors.New("Entity not found")
	}
	oldBoundaryFaceString.EndLifespanVersion = &currentTime
	crud.DB.Set("gorm:save_associations", false).Save(&oldBoundaryFaceString)

	reader = crud.DB.Where("bfsid = ?::\"Oid\" AND endlifespanversion = ?", boundaryFaceString.BfsID, currentTime).
		Preload("PlusSu", "endlifespanversion IS NULL").
		First(&oldBoundaryFaceString)
	if reader.RowsAffected == 0 {
		return errors.New("Entity not found")
	}
	for _, plusSu := range oldBoundaryFaceString.PlusSu {
		plusSu.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&plusSu)
	}
	return nil
}
