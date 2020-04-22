package logic

import (
	ladm "github.com/cdrlis/cdrLIS/LADM"
	"github.com/cdrlis/cdrLIS/LADM/common"
	"github.com/google/uuid"
	"time"
)

type LAPartyService struct {
	Context IDatabase
}

func (service LAPartyService) GetParty(id common.Oid) (*ladm.LAParty, error) {
	var party ladm.LAParty
	err := service.Context.Read(&party, "pid = ?::\"Oid\" AND endlifespanversion IS NULL", id)
	if err != nil {
		return nil, err
	}
	return &party, nil
}

func (service LAPartyService) CreateParty(party ladm.LAParty) (*ladm.LAParty, error) {
	currentTime := time.Now()
	party.ID = uuid.New().String()
	party.BeginLifespanVersion = currentTime
	party.EndLifespanVersion = nil
	err := service.Context.Create(&party)
	if err != nil {
		return nil, err
	}
	return &party, nil
}

func (service LAPartyService) GetPartyList() (*[]ladm.LAParty, error) {
	var parties []ladm.LAParty
	err := service.Context.ReadAll(&parties, "endlifespanversion IS NULL")
	if err != nil {
		return nil, err
	}
	return &parties, nil
}

func (service LAPartyService) UpdateParty(party ladm.LAParty) error {
	currentTime := time.Now()
	var oldParty ladm.LAParty
	err := service.Context.Read(&oldParty, "pid = ?::\"Oid\" AND endlifespanversion IS NULL", party.Pid)
	if err != nil {
		return err
	}
	oldParty.EndLifespanVersion = &currentTime
	service.Context.Update(&oldParty)
	party.ID = uuid.New().String()
	party.BeginLifespanVersion = currentTime
	party.EndLifespanVersion = nil
	service.Context.Create(&party)
	return nil
}

func (service LAPartyService) DeleteParty(party ladm.LAParty) error {
	currentTime := time.Now()
	var oldParty ladm.LAParty
	err := service.Context.Read(&oldParty, "pid = ?::\"Oid\" AND endlifespanversion IS NULL", party.Pid)
	if err != nil {
		return err
	}
	oldParty.EndLifespanVersion = &currentTime
	service.Context.Update(&oldParty)
	return nil
}
