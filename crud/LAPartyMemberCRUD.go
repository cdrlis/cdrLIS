package crud

import (
	"errors"
	"fmt"
	ladm "github.com/cdrlis/cdrLIS/LADM"
	"github.com/cdrlis/cdrLIS/LADM/common"
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
	partyMember := partyMemberIn.(ladm.LAPartyMember)
	currentTime := time.Now()
	partyMember.ID = fmt.Sprintf("%v-%v", partyMember.PartyID, partyMember.GroupID)
	partyMember.BeginLifespanVersion = currentTime
	partyMember.EndLifespanVersion = nil
	partyMember.PartyID = partyMember.Party.PID.String()
	partyMember.PartyBeginLifespanVersion = partyMember.Party.BeginLifespanVersion
	partyMember.GroupID = partyMember.Group.GroupID.String()
	partyMember.GroupBeginLifespanVersion = partyMember.Group.BeginLifespanVersion
	crud.DB.Set("gorm:save_associations", false).Create(&partyMember)
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
	partyMember := partyMemberIn.(*ladm.LAPartyMember)
	currentTime := time.Now()
	var oldPartyMember ladm.LAPartyMember
	reader := crud.DB.Where("parties = ? AND "+
		"groups = ? AND "+
		"endlifespanversion IS NULL", partyMember.Party.PID.String(), partyMember.Group.GroupID.String()).
		First(&oldPartyMember)
	if reader.RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	oldPartyMember.EndLifespanVersion = &currentTime
	crud.DB.Set("gorm:save_associations", false).Save(&oldPartyMember)

	partyMember.ID = fmt.Sprintf("%v-%v", partyMember.PartyID, partyMember.GroupID)
	partyMember.BeginLifespanVersion = currentTime
	partyMember.EndLifespanVersion = nil
	partyMember.PartyID = partyMember.Party.PID.String()
	partyMember.PartyBeginLifespanVersion = partyMember.Party.BeginLifespanVersion
	partyMember.GroupID = partyMember.Group.GroupID.String()
	partyMember.GroupBeginLifespanVersion = partyMember.Group.BeginLifespanVersion
	crud.DB.Set("gorm:save_associations", false).Create(&partyMember)

	return partyMember, nil
}

func (crud LAPartyMemberCRUD) Delete(partyMemberIn interface{}) error {
	partyMember := partyMemberIn.(ladm.LAPartyMember)
	currentTime := time.Now()
	var oldPartyMember ladm.LAPartyMember
	reader := crud.DB.Where("parties = ? AND "+
		"groups = ? AND "+
		"endlifespanversion IS NULL", partyMember.PartyID, partyMember.GroupID).First(&oldPartyMember)
	if reader.RowsAffected == 0 {
		return errors.New("Entity not found")
	}
	oldPartyMember.EndLifespanVersion = &currentTime
	crud.DB.Set("gorm:save_associations", false).Save(&oldPartyMember)
	return nil
}
