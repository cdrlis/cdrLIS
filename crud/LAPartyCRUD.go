package crud

import (
	"errors"
	"fmt"
	ladm "github.com/cdrlis/cdrLIS/LADM"
	"github.com/jinzhu/gorm"
	"time"
)

type LAPartyCRUD struct {
	DB *gorm.DB
}

func (crud LAPartyCRUD) Read(where ...interface{}) (interface{}, error) {
	var party ladm.LAParty
	if where != nil {
		reader := crud.DB.Where("pid = ?::\"Oid\" AND endlifespanversion IS NULL", where).
			Preload("Groups", "endlifespanversion IS NULL").
			Preload("Groups.Group", "endlifespanversion IS NULL").
			Preload("Rights", "endlifespanversion IS NULL").
			Preload("Rights.Unit", "endlifespanversion IS NULL").
			Preload("Responsibilities", "endlifespanversion IS NULL").
			Preload("Responsibilities.Unit", "endlifespanversion IS NULL").
			Preload("Restrictions", "endlifespanversion IS NULL").
			Preload("Restrictions.Unit", "endlifespanversion IS NULL").
			First(&party)
		if reader.RowsAffected == 0 {
			return nil, errors.New("Entity not found")
		}
		return party, nil
	}
	return nil, nil
}

func (crud LAPartyCRUD) Create(partyIn interface{}) (interface{}, error) {
	party := partyIn.(ladm.LAParty)
	currentTime := time.Now()
	party.ID = fmt.Sprintf("%v-%v", party.PID.Namespace, party.PID.LocalID)
	party.BeginLifespanVersion = currentTime
	party.EndLifespanVersion = nil
	crud.DB.Create(&party)
	return &party, nil
}

func (crud LAPartyCRUD) ReadAll(where ...interface{}) (interface{}, error) {
	var parties []ladm.LAParty
	if crud.DB.Where("endlifespanversion IS NULL").Find(&parties).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	return &parties, nil
}

func (crud LAPartyCRUD) Update(partyIn interface{}) (interface{}, error) {
	party := partyIn.(*ladm.LAParty)
	currentTime := time.Now()
	var oldParty ladm.LAParty
	reader := crud.DB.Where("pid = ?::\"Oid\" AND endlifespanversion IS NULL", party.PID).
		First(&oldParty)
	if reader.RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	oldParty.EndLifespanVersion = &currentTime
	crud.DB.Save(&oldParty)

	party.ID = fmt.Sprintf("%v-%v", party.PID.Namespace, party.PID.LocalID)
	party.BeginLifespanVersion = currentTime
	party.EndLifespanVersion = nil
	crud.DB.Create(&party)

	reader = crud.DB.Where("pid = ?::\"Oid\" AND endlifespanversion = ?", party.PID, currentTime).
		Preload("Groups", "endlifespanversion IS NULL").
		Preload("Rights", "endlifespanversion IS NULL").
		Preload("Responsibilities", "endlifespanversion IS NULL").
		Preload("Restrictions", "endlifespanversion IS NULL").
		First(&oldParty)
	if reader.RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	for _, group := range oldParty.Groups {
		group.EndLifespanVersion = &currentTime
		crud.DB.Save(&group)
		group.BeginLifespanVersion = currentTime
		group.EndLifespanVersion = nil
		group.PartyBeginLifespanVersion = currentTime
		crud.DB.Create(&group)
	}

	for _, right := range oldParty.Rights {
		right.EndLifespanVersion = &currentTime
		crud.DB.Save(&right)
		right.BeginLifespanVersion = currentTime
		right.EndLifespanVersion = nil
		right.PartyBeginLifespanVersion = currentTime
		crud.DB.Create(&right)
	}

	for _, responsibility := range oldParty.Responsibilities {
		responsibility.EndLifespanVersion = &currentTime
		crud.DB.Save(&responsibility)
		responsibility.BeginLifespanVersion = currentTime
		responsibility.EndLifespanVersion = nil
		responsibility.PartyBeginLifespanVersion = currentTime
		crud.DB.Create(&responsibility)
	}

	for _, restriction := range oldParty.Restrictions {
		restriction.EndLifespanVersion = &currentTime
		crud.DB.Save(&restriction)
		restriction.BeginLifespanVersion = currentTime
		restriction.EndLifespanVersion = nil
		restriction.PartyBeginLifespanVersion = currentTime
		crud.DB.Create(&restriction)
	}
	return party, nil
}

func (crud LAPartyCRUD) Delete(partyIn interface{}) error {
	party := partyIn.(ladm.LAParty)
	currentTime := time.Now()
	var oldParty ladm.LAParty
	reader := crud.DB.Where("pid = ?::\"Oid\" AND endlifespanversion IS NULL", party.PID).First(&oldParty)
	if reader.RowsAffected == 0 {
		return errors.New("Entity not found")
	}
	oldParty.EndLifespanVersion = &currentTime
	crud.DB.Save(&oldParty)

	reader = crud.DB.Where("pid = ?::\"Oid\" AND endlifespanversion = ?", party.PID, currentTime).
		Preload("Groups", "endlifespanversion IS NULL").
		Preload("Rights", "endlifespanversion IS NULL").
		Preload("Responsibilities", "endlifespanversion IS NULL").
		Preload("Restrictions", "endlifespanversion IS NULL").
		First(&oldParty)
	if reader.RowsAffected == 0 {
		return errors.New("Entity not found")
	}
	for _, group := range oldParty.Groups {
		group.EndLifespanVersion = &currentTime
		crud.DB.Save(&group)
	}

	for _, right := range oldParty.Rights {
		right.EndLifespanVersion = &currentTime
		crud.DB.Save(&right)
	}

	for _, responsibility := range oldParty.Responsibilities {
		responsibility.EndLifespanVersion = &currentTime
		crud.DB.Save(&responsibility)
	}

	for _, restriction := range oldParty.Restrictions {
		restriction.EndLifespanVersion = &currentTime
		crud.DB.Save(&restriction)
	}
	return nil
}
