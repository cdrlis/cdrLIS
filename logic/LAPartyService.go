package logic

import (
	ladm "github.com/cdrlis/cdrLIS/LADM"
)

type LAPartyService struct {
	Context IRepository
}

func (service LAPartyService) GetParty(id string) (*ladm.LAParty, error) {
	var party ladm.LAParty
	err := service.Context.Get(&party,id)
	if err != nil {
		return nil, err
	}
	return &party, nil
}

func (service LAPartyService) CreateParty(party ladm.LAParty) error {
	service.Context.Create(&party)
	return nil
}

func (service LAPartyService) GetPartyList() (*[]ladm.LAParty, error) {
	var parties []ladm.LAParty
	err := service.Context.GetAll(&parties)
	if err != nil {
		return nil, err
	}
	return &parties, nil
}

func (service LAPartyService) UpdateParty(party ladm.LAParty) error {
	service.Context.Update(&party)
	return nil
}

func (service LAPartyService) DeleteParty(party ladm.LAParty) error {
	service.Context.Delete(&party)
	return nil
}
