package crud

import (
	"errors"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/jinzhu/gorm"
	"time"
)

type MortgageRightCRUD struct {
	DB *gorm.DB
}

func (crud MortgageRightCRUD) Read(where ...interface{}) (interface{}, error) {
	var mortgageRight ladm.MortgageRight
	if where != nil {
		reader := crud.DB.Where("mortgage = ? AND "+
			"right_ = ? AND "+
			"endlifespanversion IS NULL", where[0].(common.Oid).String(), where[1].(common.Oid).String()).
			Preload("Mortgage", "endlifespanversion IS NULL").
			Preload("Right", "endlifespanversion IS NULL").
			First(&mortgageRight)
		if reader.RowsAffected == 0 {
			return nil, errors.New("Entity not found")
		}
		return mortgageRight, nil
	}
	return nil, nil
}

func (crud MortgageRightCRUD) Create(mortgageRightIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	mortgageRight := mortgageRightIn.(ladm.MortgageRight)
	currentTime := time.Now()
	mortgageRight.BeginLifespanVersion = currentTime
	mortgageRight.EndLifespanVersion = nil
	mortgageRight.MortgageID = mortgageRight.Mortgage.RID.String()
	mortgageRight.MortgageBeginLifespanVersion = mortgageRight.Mortgage.BeginLifespanVersion
	mortgageRight.RightID = mortgageRight.Right.RID.String()
	mortgageRight.RightBeginLifespanVersion = mortgageRight.Right.BeginLifespanVersion
	writer := tx.Set("gorm:save_associations", false).Create(&mortgageRight)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return nil, commit.Error
	}
	return &mortgageRight, nil
}

func (crud MortgageRightCRUD) ReadAll(where ...interface{}) (interface{}, error) {
	var mortgageRights []ladm.MortgageRight
	if crud.DB.Where("endlifespanversion IS NULL").
		Preload("Mortgage", "endlifespanversion IS NULL").
		Preload("Right", "endlifespanversion IS NULL").Find(&mortgageRights).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	return &mortgageRights, nil
}

func (crud MortgageRightCRUD) Update(mortgageRightIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	mortgageRight := mortgageRightIn.(*ladm.MortgageRight)
	currentTime := time.Now()
	var oldMortgageRight ladm.MortgageRight
	reader := tx.Where("mortgage = ? AND "+
		"right_ = ? AND "+
		"endlifespanversion IS NULL", mortgageRight.Mortgage.RID.String(), mortgageRight.Right.RID.String()).
		First(&oldMortgageRight)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	oldMortgageRight.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldMortgageRight)
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	mortgageRight.BeginLifespanVersion = currentTime
	mortgageRight.EndLifespanVersion = nil
	mortgageRight.MortgageID = mortgageRight.Mortgage.RID.String()
	mortgageRight.MortgageBeginLifespanVersion = mortgageRight.Mortgage.BeginLifespanVersion
	mortgageRight.RightID = mortgageRight.Right.RID.String()
	mortgageRight.RightBeginLifespanVersion = mortgageRight.Right.BeginLifespanVersion
	writer = tx.Set("gorm:save_associations", false).Create(&mortgageRight)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return nil, commit.Error
	}
	return mortgageRight, nil
}

func (crud MortgageRightCRUD) Delete(mortgageRightIn interface{}) error {
	tx := crud.DB.Begin()
	mortgageRight := mortgageRightIn.(ladm.MortgageRight)
	currentTime := time.Now()
	var oldMortgageRight ladm.MortgageRight
	reader := tx.Where("mortgage = ? AND "+
		"right_ = ? AND "+
		"endlifespanversion IS NULL", mortgageRight.MortgageID, mortgageRight.RightID).First(&oldMortgageRight)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	oldMortgageRight.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldMortgageRight)
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return commit.Error
	}
	return nil
}
