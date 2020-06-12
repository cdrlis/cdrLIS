package crud

import (
	"errors"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/jinzhu/gorm"
	"time"
)

type LAPartyMemberCRUD struct {
	DB *gorm.DB
}

func (crud LAPartyMemberCRUD) Read(where ...interface{}) (interface{}, error) {
	var partyMember ladm.LAPartyMember
	if where != nil {
		reader := crud.DB.Where("parties = ? AND "+
			"groups = ? AND "+
			"endlifespanversion IS NULL", where[0].(common.Oid).String(), where[1].(common.Oid).String()).
			Preload("Party", "endlifespanversion IS NULL").
			Preload("Group", "endlifespanversion IS NULL").
			First(&partyMember)
		if reader.RowsAffected == 0 {
			return nil, errors.New("Entity not found")
		}
		return partyMember, nil
	}
	return nil, nil
}

func (crud LAPartyMemberCRUD) Create(partyMemberIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	partyMember := partyMemberIn.(ladm.LAPartyMember)
	if partyMember.Group == nil{
		tx.Rollback()
		return nil, errors.New("Group not found")
	}
	existing := 0
	reader := tx.Model(&ladm.LAPartyMember{}).Where("parties = ? AND "+
		"groups = ? AND "+
		"endlifespanversion IS NULL",
		partyMember.Party.PID.String(),
		partyMember.Group.PID.String()).
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
	partyMember.BeginLifespanVersion = currentTime
	partyMember.EndLifespanVersion = nil
	partyMember.PartyID = partyMember.Party.PID.String()
	partyMember.PartyBeginLifespanVersion = partyMember.Party.BeginLifespanVersion
	partyMember.GroupID = partyMember.Group.PID.String()
	partyMember.GroupBeginLifespanVersion = partyMember.Group.BeginLifespanVersion
	writer := tx.Set("gorm:save_associations", false).Create(&partyMember)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return nil, commit.Error
	}
	return &partyMember, nil
}

func (crud LAPartyMemberCRUD) ReadAll(where ...interface{}) (interface{}, error) {
	var partyMembers []ladm.LAPartyMember
	if crud.DB.Where("endlifespanversion IS NULL").
		Preload("Party", "endlifespanversion IS NULL").
		Preload("Group", "endlifespanversion IS NULL").Find(&partyMembers).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	return &partyMembers, nil
}

func (crud LAPartyMemberCRUD) Update(partyMemberIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	partyMember := partyMemberIn.(*ladm.LAPartyMember)
	currentTime := time.Now()
	var oldPartyMember ladm.LAPartyMember
	reader := tx.Where("parties = ? AND "+
		"groups = ? AND "+
		"endlifespanversion IS NULL", partyMember.Party.PID.String(), partyMember.Group.PID.String()).
		First(&oldPartyMember)
	if reader.Error != nil{
		tx.Rollback()
		return nil, reader.Error
	}
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	oldPartyMember.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldPartyMember)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	partyMember.BeginLifespanVersion = currentTime
	partyMember.EndLifespanVersion = nil
	partyMember.PartyID = partyMember.Party.PID.String()
	partyMember.PartyBeginLifespanVersion = partyMember.Party.BeginLifespanVersion
	partyMember.GroupID = partyMember.Group.PID.String()
	partyMember.GroupBeginLifespanVersion = partyMember.Group.BeginLifespanVersion
	writer = tx.Set("gorm:save_associations", false).Create(&partyMember)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return nil, commit.Error
	}
	return partyMember, nil
}

func (crud LAPartyMemberCRUD) Delete(partyMemberIn interface{}) error {
	tx := crud.DB.Begin()
	partyMember := partyMemberIn.(ladm.LAPartyMember)
	currentTime := time.Now()
	var oldPartyMember ladm.LAPartyMember
	reader := tx.Where("parties = ? AND "+
		"groups = ? AND "+
		"endlifespanversion IS NULL", partyMember.PartyID, partyMember.GroupID).First(&oldPartyMember)
	if reader.Error != nil{
		tx.Rollback()
		return reader.Error
	}
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	oldPartyMember.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldPartyMember)
	if writer.Error != nil{
		tx.Rollback()
		return writer.Error
	}
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
