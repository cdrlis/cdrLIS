package handler

import (
	"encoding/json"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type BAUnitAsPartyHandler struct {
	BAUnitAsPartyCRUD           CRUDer
	BAUnitCRUD              CRUDer
	PartyCRUD CRUDer
}

func (handler *BAUnitAsPartyHandler) GetBAUnitAsParty(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uId := common.Oid{Namespace: p.ByName("baunitNamespace"), LocalID: p.ByName("baunitLocalId")}
	pId := common.Oid{Namespace: p.ByName("partyNamespace"), LocalID: p.ByName("partyLocalId")}
	baUnitAsParty, err := handler.BAUnitAsPartyCRUD.Read(uId, pId)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	respondJSON(w, 200, baUnitAsParty)
}

func (handler *BAUnitAsPartyHandler) GetBAUnitAsPartys(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	baUnitAsPartys, err := handler.BAUnitAsPartyCRUD.ReadAll()
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, baUnitAsPartys)
}

func (handler *BAUnitAsPartyHandler) CreateBAUnitAsParty(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var baUnitAsParty ladm.BAUnitAsParty
	err := decoder.Decode(&baUnitAsParty)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	baunit, err := handler.BAUnitCRUD.Read(baUnitAsParty.Unit.UID)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	party, err := handler.PartyCRUD.Read(baUnitAsParty.Party.PID)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	baUnitAsParty.Unit.BeginLifespanVersion = baunit.(ladm.LASpatialUnit).BeginLifespanVersion
	baUnitAsParty.Party.BeginLifespanVersion = party.(ladm.LAParty).BeginLifespanVersion
	createdBAUnitAsParty, err := handler.BAUnitAsPartyCRUD.Create(baUnitAsParty)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 201, createdBAUnitAsParty)
}

func (handler *BAUnitAsPartyHandler) UpdateBAUnitAsParty(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pId := common.Oid{Namespace: p.ByName("baunitNamespace"), LocalID: p.ByName("baunitLocalId")}
	partyId := common.Oid{Namespace: p.ByName("partyNamespace"), LocalID: p.ByName("partyLocalId")}
	decoder := json.NewDecoder(r.Body)
	baUnitAsParty, err := handler.BAUnitAsPartyCRUD.Read(pId, partyId)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	var newBAUnitAsParty ladm.BAUnitAsParty
	err = decoder.Decode(&newBAUnitAsParty)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	newBAUnitAsParty.Unit = baUnitAsParty.(ladm.BAUnitAsParty).Unit
	newBAUnitAsParty.Party = baUnitAsParty.(ladm.BAUnitAsParty).Party
	updatedBAUnitAsParty, err := handler.BAUnitAsPartyCRUD.Update(&newBAUnitAsParty)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 200, updatedBAUnitAsParty)
}

func (handler *BAUnitAsPartyHandler) DeleteBAUnitAsParty(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uId := common.Oid{Namespace: p.ByName("baunitNamespace"), LocalID: p.ByName("baunitLocalId")}
	pId := common.Oid{Namespace: p.ByName("partyNamespace"), LocalID: p.ByName("partyLocalId")}
	baUnitAsParty, err := handler.BAUnitAsPartyCRUD.Read(uId, pId)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	err = handler.BAUnitAsPartyCRUD.Delete(baUnitAsParty)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondEmpty(w, 204)
}
