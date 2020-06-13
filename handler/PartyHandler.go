package handler

import (
	"encoding/json"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PartyHandler struct {
	PartyCRUD CRUDer
}

func (handler *PartyHandler) GetParty(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pid := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	party, err := handler.PartyCRUD.Read(pid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	respondJSON(w, 200, party)
}

func (handler *PartyHandler) GetParties(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	parties, err := handler.PartyCRUD.ReadAll()
	if err != nil {
		respondError(w, 500, err.Error())
		return
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
	createdParty, err := handler.PartyCRUD.Create(party)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 201, createdParty)
}

func (handler *PartyHandler) UpdateParty(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pid := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	decoder := json.NewDecoder(r.Body)
	_, err := handler.PartyCRUD.Read(pid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	var newParty ladm.LAParty
	err = decoder.Decode(&newParty)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	newParty.PID = pid
	updatedParty, err := handler.PartyCRUD.Update(&newParty)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 200, updatedParty)
}

func (handler *PartyHandler) DeleteParty(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pid := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	party, err := handler.PartyCRUD.Read(pid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	err = handler.PartyCRUD.Delete(party)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondEmpty(w, 204)
}