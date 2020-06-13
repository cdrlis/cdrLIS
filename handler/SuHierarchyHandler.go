package handler

import (
	"encoding/json"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type SuHierarchyHandler struct {
	SuHierarchyCRUD CRUDer
	SpatialUnitCRUD CRUDer
}

func (handler *SuHierarchyHandler) GetSuHierarchy(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	parentId := common.Oid{Namespace: p.ByName("parentNamespace"), LocalID: p.ByName("parentLocalId")}
	childId := common.Oid{Namespace: p.ByName("childNamespace"), LocalID: p.ByName("childLocalId")}
	suHierarchy, err := handler.SuHierarchyCRUD.Read(parentId, childId)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	respondJSON(w, 200, suHierarchy)
}

func (handler *SuHierarchyHandler) GetSuHierarchys(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	suHierarchies, err := handler.SuHierarchyCRUD.ReadAll()
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, suHierarchies)
}

func (handler *SuHierarchyHandler) CreateSuHierarchy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var suHierarchy ladm.SuHierarchy
	err := decoder.Decode(&suHierarchy)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	parent, err := handler.SpatialUnitCRUD.Read(suHierarchy.Parent.SuID)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	child, err := handler.SpatialUnitCRUD.Read(suHierarchy.Child.SuID)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	suHierarchy.Parent.BeginLifespanVersion = parent.(ladm.LASpatialUnit).BeginLifespanVersion
	suHierarchy.Child.BeginLifespanVersion = child.(ladm.LASpatialUnit).BeginLifespanVersion
	createdSuHierarchy, err := handler.SuHierarchyCRUD.Create(suHierarchy)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 201, createdSuHierarchy)
}

func (handler *SuHierarchyHandler) UpdateSuHierarchy(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	parentId := common.Oid{Namespace: p.ByName("parentNamespace"), LocalID: p.ByName("parentLocalId")}
	childId := common.Oid{Namespace: p.ByName("childNamespace"), LocalID: p.ByName("childLocalId")}
	decoder := json.NewDecoder(r.Body)
	suHierarchy, err := handler.SuHierarchyCRUD.Read(parentId, childId)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	var newSuHierarchy ladm.SuHierarchy
	err = decoder.Decode(&newSuHierarchy)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	newSuHierarchy.Parent = suHierarchy.(ladm.SuHierarchy).Parent
	newSuHierarchy.Child = suHierarchy.(ladm.SuHierarchy).Child
	updatedSuHierarchy, err := handler.SuHierarchyCRUD.Update(&newSuHierarchy)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 200, updatedSuHierarchy)
}

func (handler *SuHierarchyHandler) DeleteSuHierarchy(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	parentId := common.Oid{Namespace: p.ByName("parentNamespace"), LocalID: p.ByName("parentLocalId")}
	childId := common.Oid{Namespace: p.ByName("childNamespace"), LocalID: p.ByName("childLocalId")}
	parentBfs, err := handler.SuHierarchyCRUD.Read(parentId, childId)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	err = handler.SuHierarchyCRUD.Delete(parentBfs)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondEmpty(w, 204)
}
