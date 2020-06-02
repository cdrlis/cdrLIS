package crud

import (
	"errors"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/jinzhu/gorm"
	"time"
)

type LAGroupPartyCRUD struct {
	DB *gorm.DB
}

func (crud LAGroupPartyCRUD) Read(where ...interface{}) (interface{}, error) {
	var party ladm.LAGroupParty
	if where != nil {
		reader := crud.DB.Where("pid = ?::\"Oid\" AND endlifespanversion IS NULL", where).
//			Preload("Groups", "endlifespanversion IS NULL").
//			Preload("Groups.Group", "endlifespanversion IS NULL").
//			Preload("Rights", "endlifespanversion IS NULL").
//			Preload("Rights.Unit", "endlifespanversion IS NULL").
//			Preload("Responsibilities", "endlifespanversion IS NULL").
//			Preload("Responsibilities.Unit", "endlifespanversion IS NULL").
//			Preload("Restrictions", "endlifespanversion IS NULL").
//			Preload("Restrictions.Unit", "endlifespanversion IS NULL").
			First(&party)
		if reader.RowsAffected == 0 {
			return nil, errors.New("Entity not found")
		}
		return party, nil
	}
	return nil, nil
}

func (crud LAGroupPartyCRUD) Create(groupPartyIn interface{}) (interface{}, error) {
	groupParty := groupPartyIn.(ladm.LAGroupParty)
	currentTime := time.Now()
	groupParty.ID = groupParty.PID.String()
	groupParty.BeginLifespanVersion = currentTime
	groupParty.EndLifespanVersion = nil
	writer := crud.DB.Set("gorm:save_associations", false).Create(&groupParty)
	if writer.Error != nil{
		return nil, writer.Error
	}
	return &groupParty, nil
}

func (crud LAGroupPartyCRUD) ReadAll(where ...interface{}) (interface{}, error) {
	var groupParties []ladm.LAGroupParty
	if crud.DB.Where("endlifespanversion IS NULL").Find(&groupParties).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	return &groupParties, nil
}

func (crud LAGroupPartyCRUD) Update(partyIn interface{}) (interface{}, error) {
	groupParty := partyIn.(*ladm.LAGroupParty)
	currentTime := time.Now()
	var oldGroupParty ladm.LAGroupParty
	reader := crud.DB.Where("groupid = ?::\"Oid\" AND endlifespanversion IS NULL", groupParty.PID).
		First(&oldGroupParty)
	if reader.RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	oldGroupParty.EndLifespanVersion = &currentTime
	crud.DB.Set("gorm:save_associations", false).Save(&oldGroupParty)

	groupParty.ID = groupParty.PID.String()
	groupParty.BeginLifespanVersion = currentTime
	groupParty.EndLifespanVersion = nil
	crud.DB.Set("gorm:save_associations", false).Create(&groupParty)
/*
	reader = crud.DB.Where("groupid = ?::\"Oid\" AND endlifespanversion = ?", groupParty.PID, currentTime).
		Preload("Groups", "endlifespanversion IS NULL").
		Preload("Rights", "endlifespanversion IS NULL").
		Preload("Responsibilities", "endlifespanversion IS NULL").
		Preload("Restrictions", "endlifespanversion IS NULL").
		First(&oldGroupParty)
	if reader.RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	for _, group := range oldGroupParty.Groups {
		group.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&group)
		group.BeginLifespanVersion = currentTime
		group.EndLifespanVersion = nil
		group.PartyBeginLifespanVersion = currentTime
		crud.DB.Set("gorm:save_associations", false).Create(&group)
	}

	for _, right := range oldGroupParty.Rights {
		right.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&right)
		right.BeginLifespanVersion = currentTime
		right.EndLifespanVersion = nil
		right.PartyBeginLifespanVersion = currentTime
		crud.DB.Set("gorm:save_associations", false).Create(&right)
	}

	for _, responsibility := range oldGroupParty.Responsibilities {
		responsibility.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&responsibility)
		responsibility.BeginLifespanVersion = currentTime
		responsibility.EndLifespanVersion = nil
		responsibility.PartyBeginLifespanVersion = currentTime
		crud.DB.Set("gorm:save_associations", false).Create(&responsibility)
	}

	for _, restriction := range oldGroupParty.Restrictions {
		restriction.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&restriction)
		restriction.BeginLifespanVersion = currentTime
		restriction.EndLifespanVersion = nil
		restriction.PartyBeginLifespanVersion = currentTime
		crud.DB.Set("gorm:save_associations", false).Create(&restriction)
	}
 */
	return groupParty, nil
}

func (crud LAGroupPartyCRUD) Delete(groupPartyIn interface{}) error {
	groupParty := groupPartyIn.(ladm.LAGroupParty)
	currentTime := time.Now()
	var oldGroupParty ladm.LAGroupParty
	reader := crud.DB.Where("groupid = ?::\"Oid\" AND endlifespanversion IS NULL", groupParty.PID).First(&oldGroupParty)
	if reader.RowsAffected == 0 {
		return errors.New("Entity not found")
	}
	oldGroupParty.EndLifespanVersion = &currentTime
	crud.DB.Set("gorm:save_associations", false).Save(&oldGroupParty)
/*
	reader = crud.DB.Where("groupid = ?::\"Oid\" AND endlifespanversion = ?", groupParty.PID, currentTime).
		Preload("Groups", "endlifespanversion IS NULL").
		Preload("Rights", "endlifespanversion IS NULL").
		Preload("Responsibilities", "endlifespanversion IS NULL").
		Preload("Restrictions", "endlifespanversion IS NULL").
		First(&oldGroupParty)
	if reader.RowsAffected == 0 {
		return errors.New("Entity not found")
	}
	for _, group := range oldGroupParty.Groups {
		group.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&group)
	}

	for _, right := range oldGroupParty.Rights {
		right.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&right)
	}

	for _, responsibility := range oldGroupParty.Responsibilities {
		responsibility.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&responsibility)
	}

	for _, restriction := range oldGroupParty.Restrictions {
		restriction.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&restriction)
	}
 */
	return nil
}
