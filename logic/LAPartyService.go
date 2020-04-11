package logic

import (
	ladm "github.com/cdrlis/cdrLIS/LADM"
)

type LAPartyService struct {
	Context IRepository
}

func (service LAPartyService) GetParty() (*ladm.LAParty, error) {
	var party ladm.LAParty
	err := service.Context.Get(&party)
	if err != nil {
		return nil, err
	}
	return &party, nil
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