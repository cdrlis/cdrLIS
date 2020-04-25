package handler

import (
	"encoding/json"
	ladm "github.com/cdrlis/cdrLIS/LADM"
	"github.com/cdrlis/cdrLIS/LADM/common"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PartyHandler struct {
	Service CRUDer
}

func (handler *PartyHandler) GetParty(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pid := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	party, err := handler.Service.Read(pid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	respondJSON(w, 200, party)
}

func (handler *PartyHandler) GetParties(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	parties, err := handler.Service.ReadAll()
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
	createdParty, err := handler.Service.Create(party)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 201, createdParty)
}

func (handler *PartyHandler) UpdateParty(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pid := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	decoder := json.NewDecoder(r.Body)
	_, err := handler.Service.Read(pid)
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
	handler.Service.Update(&newParty)
	respondJSON(w, 200, newParty)
}

func (handler *PartyHandler) DeleteParty(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pid := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	party, err := handler.Service.Read(pid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	handler.Service.Delete(party)
	respondEmpty(w, 204)
}