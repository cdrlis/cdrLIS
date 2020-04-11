package logic

import (
	ladm "github.com/cdrlis/cdrLIS/LADM"
)

type LAPartyService struct {
	Context IDatabase
}

func (service LAPartyService) ReadParty() (*ladm.LAParty, error) {
	var party ladm.LAParty
	err := service.Context.Read(&party)
	if err != nil {
		return nil, err
	}
	return &party, nil
}

func (service LAPartyService) ReadPartyList() (*[]ladm.LAParty, error) {
	var parties []ladm.LAParty
	err := service.Context.ReadAll(&parties)
	if err != nil {
		return nil, err
	}
	return &parties, nil
}

func (service LAPartyService) UpdateParty(party ladm.LAParty) error {
	err := service.Context.Update(&party)
	if err != nil {
		return err
	}
	return nil
}
