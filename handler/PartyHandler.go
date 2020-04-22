package handler

import (
	"encoding/json"
	ladm "github.com/cdrlis/cdrLIS/LADM"
	"github.com/cdrlis/cdrLIS/LADM/common"
	"github.com/cdrlis/cdrLIS/logic"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PartyHandler struct {
	Service logic.LAPartyService
}

func (handler *PartyHandler) GetParty(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pid := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	party, err := handler.Service.GetParty(pid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	respondJSON(w, 200, party)
}

func (handler *PartyHandler) GetParties(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	parties, err := handler.Service.GetPartyList()
	if err != nil {
		respondError(w, 500, err.Error())
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
	respondJSON(w, 200, parties)
}

func (handler *PartyHandler) CreateParty(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var party ladm.LAParty
	err := decoder.Decode(&party)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	createdParty, err := handler.Service.CreateParty(party)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 201, createdParty)
}

func (handler *PartyHandler) UpdateParty(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pid := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	decoder := json.NewDecoder(r.Body)
	party, err := handler.Service.GetParty(pid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	err = decoder.Decode(&party)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	handler.Service.UpdateParty(*party)
	respondJSON(w, 200, party)
}

func (handler *PartyHandler) DeleteParty(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pid := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	party, err := handler.Service.GetParty(pid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	handler.Service.DeleteParty(*party)
	respondEmpty(w, 204)
}
