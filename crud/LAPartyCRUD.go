package crud

import (
	"errors"
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
			Preload("GroupParty","endlifespanversion IS NULL").
			Preload("Groups", "endlifespanversion IS NULL").
			Preload("Groups.Group", "endlifespanversion IS NULL").
			Preload("RRR", "endlifespanversion IS NULL").
			Preload("RRR.Unit", "endlifespanversion IS NULL").
			Preload("RRR.Right", "endlifespanversion IS NULL").
			Preload("RRR.Responsibility", "endlifespanversion IS NULL").
			Preload("RRR.Restriction", "endlifespanversion IS NULL").
			Preload("RRR.Restriction.Mortgage", "endlifespanversion IS NULL").
			First(&party)
		if reader.RowsAffected == 0 {
			return nil, errors.New("Entity not found")
		}
		return party, nil
	}
	return nil, nil
}

func (crud LAPartyCRUD) Create(partyIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	party := partyIn.(ladm.LAParty)
	reader := crud.DB.Where("pid = ?::\"Oid\" AND endlifespanversion IS NULL", party.PID).First(&party)
	if reader.RowsAffected != 0 {
		tx.Rollback()
		return nil, errors.New("Entity already exists")
	}
	currentTime := time.Now()
	party.ID = party.PID.String()
	party.BeginLifespanVersion = currentTime
	party.EndLifespanVersion = nil
	writer := crud.DB.Set("gorm:save_associations", false).Create(&party)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	if groupParty := party.GroupParty; groupParty != nil {
		groupParty.PID = party.PID
		groupParty.ID = groupParty.PID.String()
		groupParty.BeginLifespanVersion = currentTime
		groupParty.EndLifespanVersion = nil
		writer = crud.DB.Set("gorm:save_associations", false).Create(&groupParty)
		if writer.Error != nil{
			tx.Rollback()
			return nil, writer.Error
		}
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return nil, commit.Error
	}
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
	tx := crud.DB.Begin()
	currentTime := time.Now()
	var oldParty ladm.LAParty
	reader := crud.DB.Where("pid = ?::\"Oid\" AND endlifespanversion IS NULL", party.PID).
		First(&oldParty)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	oldParty.EndLifespanVersion = &currentTime
	crud.DB.Set("gorm:save_associations", false).Save(&oldParty)

	party.ID = party.PID.String()
	party.BeginLifespanVersion = currentTime
	party.EndLifespanVersion = nil
	writer := crud.DB.Set("gorm:save_associations", false).Create(&party)
	if writer.Error != nil{
		tx.Rollback()
		return nil, writer.Error
	}
	reader = crud.DB.Where("pid = ?::\"Oid\" AND endlifespanversion = ?", party.PID, currentTime).
		Preload("GroupParty","endlifespanversion IS NULL").
		Preload("Groups", "endlifespanversion IS NULL").
		Preload("RRR", "endlifespanversion IS NULL").
		Preload("RRR.Right", "endlifespanversion IS NULL").
		Preload("RRR.Responsibility", "endlifespanversion IS NULL").
		Preload("RRR.Restriction", "endlifespanversion IS NULL").
		Preload("RRR.Restriction.Mortgage", "endlifespanversion IS NULL").
		First(&oldParty)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}

	if groupParty := oldParty.GroupParty; groupParty != nil{
		groupParty.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&groupParty)
		groupParty.BeginLifespanVersion = currentTime
		groupParty.EndLifespanVersion = nil
		crud.DB.Set("gorm:save_associations", false).Create(&groupParty)
	}

	for _, group := range oldParty.Groups {
		group.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&group)
		group.BeginLifespanVersion = currentTime
		group.EndLifespanVersion = nil
		group.PartyBeginLifespanVersion = currentTime
		crud.DB.Set("gorm:save_associations", false).Create(&group)
	}

	for _, rrr := range oldParty.RRR {
		rrr.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&rrr)
		rrr.BeginLifespanVersion = currentTime
		rrr.EndLifespanVersion = nil
		rrr.PartyBeginLifespanVersion = currentTime
		crud.DB.Set("gorm:save_associations", false).Create(&rrr)
		if right := rrr.Right; right != nil{
			right.EndLifespanVersion = &currentTime
			crud.DB.Set("gorm:save_associations", false).Save(right)
			right.BeginLifespanVersion = currentTime
			right.EndLifespanVersion = nil
			crud.DB.Set("gorm:save_associations", false).Create(right)
		}
		if restriction := rrr.Restriction; restriction != nil{
			restriction.EndLifespanVersion = &currentTime
			crud.DB.Set("gorm:save_associations", false).Save(restriction)
			restriction.BeginLifespanVersion = currentTime
			restriction.EndLifespanVersion = nil
			crud.DB.Set("gorm:save_associations", false).Create(restriction)
			if mortgage := rrr.Restriction.Mortgage; mortgage != nil{
				mortgage.EndLifespanVersion = &currentTime
				crud.DB.Set("gorm:save_associations", false).Save(mortgage)
				mortgage.BeginLifespanVersion = currentTime
				mortgage.EndLifespanVersion = nil
				crud.DB.Set("gorm:save_associations", false).Create(mortgage)
			}
		}
		if responsibility := rrr.Responsibility; responsibility != nil{
			responsibility.EndLifespanVersion = &currentTime
			crud.DB.Set("gorm:save_associations", false).Save(responsibility)
			responsibility.BeginLifespanVersion = currentTime
			responsibility.EndLifespanVersion = nil
			crud.DB.Set("gorm:save_associations", false).Create(responsibility)
		}
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return nil, commit.Error
	}
	return party, nil
}

func (crud LAPartyCRUD) Delete(partyIn interface{}) error {
	party := partyIn.(ladm.LAParty)
	tx := crud.DB.Begin()
	currentTime := time.Now()
	var oldParty ladm.LAParty
	reader := crud.DB.Where("pid = ?::\"Oid\" AND endlifespanversion IS NULL", party.PID).First(&oldParty)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	oldParty.EndLifespanVersion = &currentTime
	crud.DB.Set("gorm:save_associations", false).Save(&oldParty)

	reader = crud.DB.Where("pid = ?::\"Oid\" AND endlifespanversion = ?", party.PID, currentTime).
		Preload("GroupParty","endlifespanversion IS NULL").
		Preload("Groups", "endlifespanversion IS NULL").
		Preload("RRR", "endlifespanversion IS NULL").
		Preload("RRR.Right", "endlifespanversion IS NULL").
		Preload("RRR.Responsibility", "endlifespanversion IS NULL").
		Preload("RRR.Restriction", "endlifespanversion IS NULL").
		Preload("RRR.Restriction.Mortgage", "endlifespanversion IS NULL").
		First(&oldParty)
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	if groupParty := oldParty.GroupParty; groupParty != nil{
		groupParty.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&groupParty)
	}

	for _, group := range oldParty.Groups {
		group.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&group)
	}

	for _, rrr := range oldParty.RRR {
		rrr.EndLifespanVersion = &currentTime
		crud.DB.Set("gorm:save_associations", false).Save(&rrr)
		if right := rrr.Right; right != nil{
			right.EndLifespanVersion = &currentTime
			crud.DB.Set("gorm:save_associations", false).Save(right)
		}
		if restriction := rrr.Restriction; restriction != nil{
			restriction.EndLifespanVersion = &currentTime
			crud.DB.Set("gorm:save_associations", false).Save(restriction)
			if mortgage := rrr.Restriction.Mortgage; mortgage != nil{
				mortgage.EndLifespanVersion = &currentTime
				crud.DB.Set("gorm:save_associations", false).Save(mortgage)
			}
		}
		if responsibility := rrr.Responsibility; responsibility != nil{
			responsibility.EndLifespanVersion = &currentTime
			crud.DB.Set("gorm:save_associations", false).Save(responsibility)
		}
	}
	commit := tx.Commit()
	if commit.Error != nil{
		return commit.Error
	}
	return nil
}
