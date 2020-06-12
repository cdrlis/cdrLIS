package crud

import (
	"errors"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/jinzhu/gorm"
	"time"
)

type SuGroupHierarchyCRUD struct {
	DB *gorm.DB
}

func (crud SuGroupHierarchyCRUD) Read(where ...interface{}) (interface{}, error) {
	var suGroupHierarchy ladm.SuGroupHierarchy
	if where != nil {
		reader := crud.DB.Where("set = ? AND "+
			"element = ? AND "+
			"endlifespanversion IS NULL", where[0].(common.Oid).String(), where[1].(common.Oid).String()).
			Preload("Set", "endlifespanversion IS NULL").
			Preload("Element", "endlifespanversion IS NULL").
			First(&suGroupHierarchy)
		if reader.RowsAffected == 0 {
			return nil, errors.New("Entity not found")
		}
		return suGroupHierarchy, nil
	}
	return nil, nil
}

func (crud SuGroupHierarchyCRUD) Create(suGroupHierarchyIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	suGroupHierarchy := suGroupHierarchyIn.(ladm.SuGroupHierarchy)
	existing := 0
	reader := tx.Model(&ladm.SuGroupHierarchy{}).Where("set = ? AND "+
		"element = ? AND "+
		"endlifespanversion IS NULL",
		suGroupHierarchy.Set.SugID.String(),
		suGroupHierarchy.Element.SugID.String()).
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
	suGroupHierarchy.BeginLifespanVersion = currentTime
	suGroupHierarchy.EndLifespanVersion = nil
	suGroupHierarchy.SetID = suGroupHierarchy.Set.SugID.String()
	suGroupHierarchy.SetBeginLifespanVersion = suGroupHierarchy.Set.BeginLifespanVersion
	suGroupHierarchy.ElementID = suGroupHierarchy.Element.SugID.String()
	suGroupHierarchy.ElementBeginLifespanVersion = suGroupHierarchy.Element.BeginLifespanVersion
	writer := tx.Set("gorm:save_associations", false).Create(&suGroupHierarchy)
	if writer.Error != nil {
		tx.Rollback()
		return nil, writer.Error
	}
	commit := tx.Commit()
	if commit.Error != nil {
		return nil, commit.Error
	}
	return &suGroupHierarchy, nil
}

func (crud SuGroupHierarchyCRUD) ReadAll(where ...interface{}) (interface{}, error) {
	var suGroupHierarchies []ladm.SuGroupHierarchy
	if crud.DB.Where("endlifespanversion IS NULL").
		Preload("Set", "endlifespanversion IS NULL").
		Preload("Element", "endlifespanversion IS NULL").Find(&suGroupHierarchies).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	return &suGroupHierarchies, nil
}

func (crud SuGroupHierarchyCRUD) Update(suGroupHierarchyIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	suGroupHierarchy := suGroupHierarchyIn.(*ladm.SuGroupHierarchy)
	currentTime := time.Now()
	var oldSuGroupHierarchy ladm.SuGroupHierarchy
	reader := tx.Where("set = ? AND "+
		"element = ? AND "+
		"endlifespanversion IS NULL", suGroupHierarchy.Set.SugID.String(), suGroupHierarchy.Element.SugID.String()).
		First(&oldSuGroupHierarchy)
	if reader.Error != nil{
		tx.Rollback()
		return nil, reader.Error
	}
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	oldSuGroupHierarchy.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldSuGroupHierarchy)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	suGroupHierarchy.BeginLifespanVersion = currentTime
	suGroupHierarchy.EndLifespanVersion = nil
	suGroupHierarchy.SetID = suGroupHierarchy.Set.SugID.String()
	suGroupHierarchy.SetBeginLifespanVersion = suGroupHierarchy.Set.BeginLifespanVersion
	suGroupHierarchy.ElementID = suGroupHierarchy.Element.SugID.String()
	suGroupHierarchy.ElementBeginLifespanVersion = suGroupHierarchy.Element.BeginLifespanVersion
	writer = tx.Set("gorm:save_associations", false).Create(&suGroupHierarchy)
	if writer.Error != nil {
		tx.Rollback()
		return nil, writer.Error
	}
	commit := tx.Commit()
	if commit.Error != nil {
		return nil, commit.Error
	}
	return suGroupHierarchy, nil
}

func (crud SuGroupHierarchyCRUD) Delete(suGroupHierarchyIn interface{}) error {
	tx := crud.DB.Begin()
	suGroupHierarchy := suGroupHierarchyIn.(ladm.SuGroupHierarchy)
	currentTime := time.Now()
	var oldSuGroupHierarchy ladm.SuGroupHierarchy
	reader := tx.Where("set = ? AND "+
		"element = ? AND "+
		"endlifespanversion IS NULL", suGroupHierarchy.SetID, suGroupHierarchy.ElementID).First(&oldSuGroupHierarchy)
	if reader.Error != nil{
		tx.Rollback()
		return reader.Error
	}
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	oldSuGroupHierarchy.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldSuGroupHierarchy)
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
