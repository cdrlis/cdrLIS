package handler

import (
	"encoding/json"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type SpatialUnitGroupHandler struct {
	SuGroupCRUD CRUDer
}

func (handler *SpatialUnitGroupHandler) GetSuGroup(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	lid := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	suGroup, err := handler.SuGroupCRUD.Read(lid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	respondJSON(w, 200, suGroup)
}

func (handler *SpatialUnitGroupHandler) GetSuGroups(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	suGroups, err := handler.SuGroupCRUD.ReadAll()
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, suGroups)
}

func (handler *SpatialUnitGroupHandler) CreateSuGroup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var suGroup ladm.LASpatialUnitGroup
	err := decoder.Decode(&suGroup)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	createdSuGroup, err := handler.SuGroupCRUD.Create(suGroup)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 201, createdSuGroup)
}

func (handler *SpatialUnitGroupHandler) UpdateSuGroup(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	sugid := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	decoder := json.NewDecoder(r.Body)
	_, err := handler.SuGroupCRUD.Read(sugid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	var newSuGroup ladm.LASpatialUnitGroup
	err = decoder.Decode(&newSuGroup)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	newSuGroup.SugID = sugid
	updatedSuGroup, err := handler.SuGroupCRUD.Update(&newSuGroup)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 200, updatedSuGroup)
}

func (handler *SpatialUnitGroupHandler) DeleteSuGroup(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	lid := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	suGroup, err := handler.SuGroupCRUD.Read(lid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	err = handler.SuGroupCRUD.Delete(suGroup)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondEmpty(w, 204)
}