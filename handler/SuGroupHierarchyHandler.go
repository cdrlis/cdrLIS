package handler

import (
	"encoding/json"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type SuGroupHierarchyHandler struct {
	SuGroupHierarchyCRUD CRUDer
	SpatialUnitGroupCRUD CRUDer
}

func (handler *SuGroupHierarchyHandler) GetSuGroupHierarchy(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	setId := common.Oid{Namespace: p.ByName("setNamespace"), LocalID: p.ByName("setLocalId")}
	elementId := common.Oid{Namespace: p.ByName("elementNamespace"), LocalID: p.ByName("elementLocalId")}
	suGroupHierarchy, err := handler.SuGroupHierarchyCRUD.Read(setId, elementId)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	respondJSON(w, 200, suGroupHierarchy)
}

func (handler *SuGroupHierarchyHandler) GetSuGroupHierarchys(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	suGroupHierarchies, err := handler.SuGroupHierarchyCRUD.ReadAll()
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, suGroupHierarchies)
}

func (handler *SuGroupHierarchyHandler) CreateSuGroupHierarchy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var suGroupHierarchy ladm.SuGroupHierarchy
	err := decoder.Decode(&suGroupHierarchy)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	set, err := handler.SpatialUnitGroupCRUD.Read(suGroupHierarchy.Set.SugID)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	element, err := handler.SpatialUnitGroupCRUD.Read(suGroupHierarchy.Element.SugID)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	suGroupHierarchy.Set.BeginLifespanVersion = set.(ladm.LASpatialUnit).BeginLifespanVersion
	suGroupHierarchy.Element.BeginLifespanVersion = element.(ladm.LASpatialUnit).BeginLifespanVersion
	createdSuGroupHierarchy, err := handler.SuGroupHierarchyCRUD.Create(suGroupHierarchy)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 201, createdSuGroupHierarchy)
}

func (handler *SuGroupHierarchyHandler) UpdateSuGroupHierarchy(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	setId := common.Oid{Namespace: p.ByName("setNamespace"), LocalID: p.ByName("setLocalId")}
	elementId := common.Oid{Namespace: p.ByName("elementNamespace"), LocalID: p.ByName("elementLocalId")}
	decoder := json.NewDecoder(r.Body)
	suGroupHierarchy, err := handler.SuGroupHierarchyCRUD.Read(setId, elementId)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	var newSuGroupHierarchy ladm.SuGroupHierarchy
	err = decoder.Decode(&newSuGroupHierarchy)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	newSuGroupHierarchy.Set = suGroupHierarchy.(ladm.SuGroupHierarchy).Set
	newSuGroupHierarchy.Element = suGroupHierarchy.(ladm.SuGroupHierarchy).Element
	updatedSuGroupHierarchy, err := handler.SuGroupHierarchyCRUD.Update(&newSuGroupHierarchy)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 200, updatedSuGroupHierarchy)
}

func (handler *SuGroupHierarchyHandler) DeleteSuGroupHierarchy(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	setId := common.Oid{Namespace: p.ByName("setNamespace"), LocalID: p.ByName("setLocalId")}
	elementId := common.Oid{Namespace: p.ByName("elementNamespace"), LocalID: p.ByName("elementLocalId")}
	suGroupHierarchy, err := handler.SuGroupHierarchyCRUD.Read(setId, elementId)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	err = handler.SuGroupHierarchyCRUD.Delete(suGroupHierarchy)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondEmpty(w, 204)
}
