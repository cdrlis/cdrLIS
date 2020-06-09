package handler

import (
	"encoding/json"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type SuSuGroupHandler struct {
	SuSuGroupCRUD CRUDer
	SuGroupCRUD   CRUDer
	SuCRUD        CRUDer
}

func (handler *SuSuGroupHandler) GetSuSuGroup(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uId := common.Oid{Namespace: p.ByName("sugNamespace"), LocalID: p.ByName("sugLocalId")}
	pId := common.Oid{Namespace: p.ByName("suNamespace"), LocalID: p.ByName("suLocalId")}
	suSuGroup, err := handler.SuSuGroupCRUD.Read(uId, pId)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	respondJSON(w, 200, suSuGroup)
}

func (handler *SuSuGroupHandler) GetSuSuGroups(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	suSuGroups, err := handler.SuSuGroupCRUD.ReadAll()
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, suSuGroups)
}

func (handler *SuSuGroupHandler) CreateSuSuGroup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var suSuGroup ladm.SuSuGroup
	err := decoder.Decode(&suSuGroup)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	sug, err := handler.SuGroupCRUD.Read(suSuGroup.Whole.SugID)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	su, err := handler.SuCRUD.Read(suSuGroup.Part.SuID)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	suSuGroup.Whole.BeginLifespanVersion = sug.(ladm.LASpatialUnitGroup).BeginLifespanVersion
	suSuGroup.Part.BeginLifespanVersion = su.(ladm.LASpatialUnit).BeginLifespanVersion
	createdSuSuGroup, err := handler.SuSuGroupCRUD.Create(suSuGroup)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 201, createdSuSuGroup)
}

func (handler *SuSuGroupHandler) UpdateSuSuGroup(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pId := common.Oid{Namespace: p.ByName("sugNamespace"), LocalID: p.ByName("sugLocalId")}
	suId := common.Oid{Namespace: p.ByName("suNamespace"), LocalID: p.ByName("suLocalId")}
	decoder := json.NewDecoder(r.Body)
	suSuGroup, err := handler.SuSuGroupCRUD.Read(pId, suId)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	var newSuSuGroup ladm.SuSuGroup
	err = decoder.Decode(&newSuSuGroup)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	newSuSuGroup.Whole = suSuGroup.(ladm.SuSuGroup).Whole
	newSuSuGroup.Part = suSuGroup.(ladm.SuSuGroup).Part
	handler.SuSuGroupCRUD.Update(&newSuSuGroup)
	respondJSON(w, 200, newSuSuGroup)
}

func (handler *SuSuGroupHandler) DeleteSuSuGroup(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uId := common.Oid{Namespace: p.ByName("sugNamespace"), LocalID: p.ByName("sugLocalId")}
	pId := common.Oid{Namespace: p.ByName("suNamespace"), LocalID: p.ByName("suLocalId")}
	suSuGroup, err := handler.SuSuGroupCRUD.Read(uId, pId)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	handler.SuSuGroupCRUD.Delete(suSuGroup)
	respondEmpty(w, 204)
}
