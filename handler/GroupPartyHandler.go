package handler

import (
	"encoding/json"
	ladm "github.com/cdrlis/cdrLIS/LADM"
	"github.com/cdrlis/cdrLIS/LADM/common"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type GroupPartyHandler struct {
	GroupPartyCRUD CRUDer
}

func (handler *GroupPartyHandler) GetParty(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pid := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	party, err := handler.GroupPartyCRUD.Read(pid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	respondJSON(w, 200, party)
}

func (handler *GroupPartyHandler) GetParties(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	parties, err := handler.GroupPartyCRUD.ReadAll()
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, parties)
}

func (handler *GroupPartyHandler) CreateParty(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var party ladm.LAParty
	err := decoder.Decode(&party)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	createdParty, err := handler.GroupPartyCRUD.Create(party)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 201, createdParty)
}

func (handler *GroupPartyHandler) UpdateParty(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pid := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	decoder := json.NewDecoder(r.Body)
	_, err := handler.GroupPartyCRUD.Read(pid)
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
	handler.GroupPartyCRUD.Update(&newParty)
	respondJSON(w, 200, newParty)
}

func (handler *GroupPartyHandler) DeleteParty(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pid := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	party, err := handler.GroupPartyCRUD.Read(pid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	handler.GroupPartyCRUD.Delete(party)
	respondEmpty(w, 204)
}