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
			Preload("Groups").
			Preload("Groups.Group").
			Preload("Rights").
			Preload("Rights.Unit").
			Preload("Responsibilities").
			Preload("Responsibilities.Unit").
			Preload("Restrictions").
			Preload("Restrictions.Unit").
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
	if crud.DB.Set("gorm:auto_preload", true).Where("pid = ?::\"Oid\" AND endlifespanversion IS NULL", party.PID).First(&oldParty).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	oldParty.EndLifespanVersion = &currentTime
	crud.DB.Save(&oldParty)
	party.ID = fmt.Sprintf("%v-%v", party.PID.Namespace, party.PID.LocalID)
	party.BeginLifespanVersion = currentTime
	party.EndLifespanVersion = nil
	crud.DB.Create(&party)
	return party, nil
}

func (crud LAPartyCRUD) Delete(partyIn interface{}) error {
	party := partyIn.(ladm.LAParty)
	currentTime := time.Now()
	var oldParty ladm.LAParty
	if crud.DB.Set("gorm:auto_preload", true).Where("pid = ?::\"Oid\" AND endlifespanversion IS NULL", party.PID).First(&oldParty).RowsAffected == 0 {
		return errors.New("Entity not found")
	}
	oldParty.EndLifespanVersion = &currentTime
	crud.DB.Save(&oldParty)
	return nil
}
