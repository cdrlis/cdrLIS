package handler

import (
	"encoding/json"
	ladm "github.com/cdrlis/cdrLIS/LADM"
	"github.com/cdrlis/cdrLIS/logic"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PartyHandler struct {
	Service logic.LAPartyService
}

func (handler *PartyHandler) GetParty(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	party, err := handler.Service.GetParty()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	respondJSON(w,200,party)
}

func (handler *PartyHandler) GetParties(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	parties, err := handler.Service.GetPartyList()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	var partyList []string
	for _, party := range *parties {
		name := "-"
		if party.Name != nil {
			name = *party.Name
		}
		partyList = append(partyList, name)
	}
	respondJSON(w,200,parties)
}

func (handler *PartyHandler) CreateParty(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var party ladm.LAParty
	err := decoder.Decode(&party)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	handler.Service.CreateParty(party)
	respondJSON(w,201,party)
}

func (handler *PartyHandler) UpdateParty(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	party, err := handler.Service.GetParty()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	newName := *party.Name + "1"
	party.Name = &newName
	handler.Service.UpdateParty(*party)
	respondJSON(w,200,party)
}

func (handler *PartyHandler) DeleteParty(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	party, err := handler.Service.GetParty()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	handler.Service.DeleteParty(*party)
	respondJSON(w,200,party)
}