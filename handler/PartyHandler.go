package handler

import (
	"github.com/cdrlis/cdrLIS/logic"
	"net/http"
)

type PartyHandler struct {
	Service logic.LAPartyService
}

func (handler *PartyHandler) GetParty(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		respondError(w, 405,"Method not supported.")
		http.Error(w, http.StatusText(405), 405)
		return
	}
	party, err := handler.Service.GetParty()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	respondJSON(w,200,party)
}

func (handler *PartyHandler) GetParties(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
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

func (handler *PartyHandler) UpdateParty(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
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
