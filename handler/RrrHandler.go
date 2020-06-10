package handler

import (
	"encoding/json"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type RrrHandler struct {
	RrrCRUD    CRUDer
	PartyCRUD  CRUDer
	BAUnitCRUD CRUDer
}

func (handler *RrrHandler) GetRrr(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rid := common.Oid{Namespace: p.ByName("namespace"), LocalID: p.ByName("localId")}
	rrr, err := handler.RrrCRUD.Read(rid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	respondJSON(w, 200, rrr)
}

func (handler *RrrHandler) GetRrrs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rrrs, err := handler.RrrCRUD.ReadAll()
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, rrrs)
}

func (handler *RrrHandler) CreateRrr(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var rrr ladm.LARRR
	err := decoder.Decode(&rrr)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	partyInterface, err := handler.PartyCRUD.Read(rrr.Party.PID)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	unitInterface, err := handler.BAUnitCRUD.Read(rrr.Unit.UID)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	party := partyInterface.(ladm.LAParty)
	unit := unitInterface.(ladm.LABAUnit)
	rrr.Party = &party
	rrr.Unit = &unit

	createdRrr, err := handler.RrrCRUD.Create(rrr)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 201, createdRrr)
}

func (handler *RrrHandler) UpdateRrr(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rid := common.Oid{Namespace: p.ByName("namespace"), LocalID: p.ByName("localId")}
	decoder := json.NewDecoder(r.Body)
	_, err := handler.RrrCRUD.Read(rid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	var newRrr ladm.LARRR
	err = decoder.Decode(&newRrr)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	handler.RrrCRUD.Update(&newRrr)
	respondJSON(w, 200, newRrr)
}

func (handler *RrrHandler) DeleteRrr(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rid := common.Oid{Namespace: p.ByName("namespace"), LocalID: p.ByName("localId")}
	rrr, err := handler.RrrCRUD.Read(rid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	handler.RrrCRUD.Delete(rrr)
	respondEmpty(w, 204)
}
