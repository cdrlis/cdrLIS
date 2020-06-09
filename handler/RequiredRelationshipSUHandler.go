package handler

import (
	"encoding/json"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type RequiredRelationshipSUHandler struct {
	RequiredRelationshipSUCRUD CRUDer
	SpatialUnitCRUD            CRUDer
}

func (handler *RequiredRelationshipSUHandler) GetRequiredRelationshipSU(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	su1Id := common.Oid{Namespace: p.ByName("su1Namespace"), LocalID: p.ByName("su1LocalId")}
	su2Id := common.Oid{Namespace: p.ByName("su2Namespace"), LocalID: p.ByName("su2LocalId")}
	relationshioSU, err := handler.RequiredRelationshipSUCRUD.Read(su1Id, su2Id)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	respondJSON(w, 200, relationshioSU)
}

func (handler *RequiredRelationshipSUHandler) GetRequiredRelationshipSUs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	suHierarchies, err := handler.RequiredRelationshipSUCRUD.ReadAll()
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, suHierarchies)
}

func (handler *RequiredRelationshipSUHandler) CreateRequiredRelationshipSU(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var relationshioSU ladm.LARequiredRelationshipSpatialUnit
	err := decoder.Decode(&relationshioSU)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	su1, err := handler.SpatialUnitCRUD.Read(relationshioSU.Su1.SuID)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	su2, err := handler.SpatialUnitCRUD.Read(relationshioSU.Su2.SuID)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	relationshioSU.Su1.BeginLifespanVersion = su1.(ladm.LASpatialUnit).BeginLifespanVersion
	relationshioSU.Su2.BeginLifespanVersion = su2.(ladm.LASpatialUnit).BeginLifespanVersion
	createdRequiredRelationshipSU, err := handler.RequiredRelationshipSUCRUD.Create(relationshioSU)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 201, createdRequiredRelationshipSU)
}

func (handler *RequiredRelationshipSUHandler) UpdateRequiredRelationshipSU(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	su1Id := common.Oid{Namespace: p.ByName("su1Namespace"), LocalID: p.ByName("su1LocalId")}
	su2Id := common.Oid{Namespace: p.ByName("su2Namespace"), LocalID: p.ByName("su2LocalId")}
	decoder := json.NewDecoder(r.Body)
	relationshioSU, err := handler.RequiredRelationshipSUCRUD.Read(su1Id, su2Id)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	var newRequiredRelationshipSU ladm.LARequiredRelationshipSpatialUnit
	err = decoder.Decode(&newRequiredRelationshipSU)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	newRequiredRelationshipSU.Su1 = relationshioSU.(ladm.LARequiredRelationshipSpatialUnit).Su1
	newRequiredRelationshipSU.Su2 = relationshioSU.(ladm.LARequiredRelationshipSpatialUnit).Su2
	handler.RequiredRelationshipSUCRUD.Update(&newRequiredRelationshipSU)
	respondJSON(w, 200, newRequiredRelationshipSU)
}

func (handler *RequiredRelationshipSUHandler) DeleteRequiredRelationshipSU(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	su1Id := common.Oid{Namespace: p.ByName("su1Namespace"), LocalID: p.ByName("su1LocalId")}
	su2Id := common.Oid{Namespace: p.ByName("su2Namespace"), LocalID: p.ByName("su2LocalId")}
	su1Bfs, err := handler.RequiredRelationshipSUCRUD.Read(su1Id, su2Id)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	handler.RequiredRelationshipSUCRUD.Delete(su1Bfs)
	respondEmpty(w, 204)
}
