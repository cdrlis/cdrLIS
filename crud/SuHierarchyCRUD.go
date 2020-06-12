package crud

import (
	"errors"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/jinzhu/gorm"
	"time"
)

type SuHierarchyCRUD struct {
	DB *gorm.DB
}

func (crud SuHierarchyCRUD) Read(where ...interface{}) (interface{}, error) {
	var suHierarchy ladm.SuHierarchy
	if where != nil {
		reader := crud.DB.Where("parent = ? AND "+
			"child = ? AND "+
			"endlifespanversion IS NULL", where[0].(common.Oid).String(), where[1].(common.Oid).String()).
			Preload("Parent", "endlifespanversion IS NULL").
			Preload("Child", "endlifespanversion IS NULL").
			First(&suHierarchy)
		if reader.RowsAffected == 0 {
			return nil, errors.New("Entity not found")
		}
		return suHierarchy, nil
	}
	return nil, nil
}

func (crud SuHierarchyCRUD) Create(suHierarchyIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	suHierarchy := suHierarchyIn.(ladm.SuHierarchy)
	existing := 0
	reader := tx.Model(&ladm.SuHierarchy{}).Where("parent = ? AND "+
		"child = ? AND "+
		"endlifespanversion IS NULL",
		suHierarchy.Parent.SuID.String(),
		suHierarchy.Child.SuID.String()).
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
	suHierarchy.BeginLifespanVersion = currentTime
	suHierarchy.EndLifespanVersion = nil
	suHierarchy.ParentID = suHierarchy.Parent.SuID.String()
	suHierarchy.ParentBeginLifespanVersion = suHierarchy.Parent.BeginLifespanVersion
	suHierarchy.ChildID = suHierarchy.Child.SuID.String()
	suHierarchy.ChildBeginLifespanVersion = suHierarchy.Child.BeginLifespanVersion
	writer := tx.Set("gorm:save_associations", false).Create(&suHierarchy)
	if writer.Error != nil {
		tx.Rollback()
		return nil, writer.Error
	}
	commit := tx.Commit()
	if commit.Error != nil {
		return nil, commit.Error
	}
	return &suHierarchy, nil
}

func (crud SuHierarchyCRUD) ReadAll(where ...interface{}) (interface{}, error) {
	var suHierarchies []ladm.SuHierarchy
	if crud.DB.Where("endlifespanversion IS NULL").
		Preload("Parent", "endlifespanversion IS NULL").
		Preload("Child", "endlifespanversion IS NULL").Find(&suHierarchies).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	return &suHierarchies, nil
}

func (crud SuHierarchyCRUD) Update(suHierarchyIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	suHierarchy := suHierarchyIn.(*ladm.SuHierarchy)
	currentTime := time.Now()
	var oldSuHierarchy ladm.SuHierarchy
	reader := tx.Where("parent = ? AND "+
		"child = ? AND "+
		"endlifespanversion IS NULL", suHierarchy.Parent.SuID.String(), suHierarchy.Child.SuID.String()).
		First(&oldSuHierarchy)
	if reader.Error != nil{
		tx.Rollback()
		return nil, reader.Error
	}
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	oldSuHierarchy.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldSuHierarchy)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	suHierarchy.BeginLifespanVersion = currentTime
	suHierarchy.EndLifespanVersion = nil
	suHierarchy.ParentID = suHierarchy.Parent.SuID.String()
	suHierarchy.ParentBeginLifespanVersion = suHierarchy.Parent.BeginLifespanVersion
	suHierarchy.ChildID = suHierarchy.Child.SuID.String()
	suHierarchy.ChildBeginLifespanVersion = suHierarchy.Child.BeginLifespanVersion
	writer = tx.Set("gorm:save_associations", false).Create(&suHierarchy)
	if writer.Error != nil {
		tx.Rollback()
		return nil, writer.Error
	}
	commit := tx.Commit()
	if commit.Error != nil {
		return nil, commit.Error
	}
	return suHierarchy, nil
}

func (crud SuHierarchyCRUD) Delete(suHierarchyIn interface{}) error {
	tx := crud.DB.Begin()
	suHierarchy := suHierarchyIn.(ladm.SuHierarchy)
	currentTime := time.Now()
	var oldSuHierarchy ladm.SuHierarchy
	reader := tx.Where("parent = ? AND "+
		"child = ? AND "+
		"endlifespanversion IS NULL", suHierarchy.ParentID, suHierarchy.ChildID).First(&oldSuHierarchy)
	if reader.Error != nil{
		tx.Rollback()
		return reader.Error
	}
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	oldSuHierarchy.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldSuHierarchy)
	if writer.Error != nil{
		tx.Rollback()
		return writer.Error
	}
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
