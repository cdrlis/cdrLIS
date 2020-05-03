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
	groupId := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	party, err := handler.GroupPartyCRUD.Read(groupId)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	respondJSON(w, 200, party)
}

func (handler *GroupPartyHandler) GetParties(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	groupParties, err := handler.GroupPartyCRUD.ReadAll()
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, groupParties)
}

func (handler *GroupPartyHandler) CreateParty(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var groupParty ladm.LAGroupParty
	err := decoder.Decode(&groupParty)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	createdGroupParty, err := handler.GroupPartyCRUD.Create(groupParty)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 201, createdGroupParty)
}

func (handler *GroupPartyHandler) UpdateParty(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	groupId := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	decoder := json.NewDecoder(r.Body)
	_, err := handler.GroupPartyCRUD.Read(groupId)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	var newGroupParty ladm.LAGroupParty
	err = decoder.Decode(&newGroupParty)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	handler.GroupPartyCRUD.Update(&newGroupParty)
	respondJSON(w, 200, newGroupParty)
}

func (handler *GroupPartyHandler) DeleteParty(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	groupId := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	groupParty, err := handler.GroupPartyCRUD.Read(groupId)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	handler.GroupPartyCRUD.Delete(groupParty)
	respondEmpty(w, 204)
}