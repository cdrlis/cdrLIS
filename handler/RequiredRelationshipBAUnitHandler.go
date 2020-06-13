package handler

import (
	"encoding/json"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type RequiredRelationshipBAUnitHandler struct {
	RequiredRelationshipBAUnitCRUD CRUDer
	BAUnitCRUD CRUDer
}

func (handler *RequiredRelationshipBAUnitHandler) GetRequiredRelationshipBAUnit(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	unit1Id := common.Oid{Namespace: p.ByName("unit1Namespace"), LocalID: p.ByName("unit1LocalId")}
	unit2Id := common.Oid{Namespace: p.ByName("unit2Namespace"), LocalID: p.ByName("unit2LocalId")}
	relationshipBAUnit, err := handler.RequiredRelationshipBAUnitCRUD.Read(unit1Id, unit2Id)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	respondJSON(w, 200, relationshipBAUnit)
}

func (handler *RequiredRelationshipBAUnitHandler) GetRequiredRelationshipBAUnits(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	relationshipBAUnits, err := handler.RequiredRelationshipBAUnitCRUD.ReadAll()
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, relationshipBAUnits)
}

func (handler *RequiredRelationshipBAUnitHandler) CreateRequiredRelationshipBAUnit(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var relationshipBAUnit ladm.LARequiredRelationshipBAUnit
	err := decoder.Decode(&relationshipBAUnit)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	unit1, err := handler.BAUnitCRUD.Read(relationshipBAUnit.Unit1.UID)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	unit2, err := handler.BAUnitCRUD.Read(relationshipBAUnit.Unit2.UID)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	relationshipBAUnit.Unit1.BeginLifespanVersion = unit1.(ladm.LABAUnit).BeginLifespanVersion
	relationshipBAUnit.Unit2.BeginLifespanVersion = unit2.(ladm.LABAUnit).BeginLifespanVersion
	createdRequiredRelationshipBAUnit, err := handler.RequiredRelationshipBAUnitCRUD.Create(relationshipBAUnit)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 201, createdRequiredRelationshipBAUnit)
}

func (handler *RequiredRelationshipBAUnitHandler) UpdateRequiredRelationshipBAUnit(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	unit1Id := common.Oid{Namespace: p.ByName("unit1Namespace"), LocalID: p.ByName("unit1LocalId")}
	unit2Id := common.Oid{Namespace: p.ByName("unit2Namespace"), LocalID: p.ByName("unit2LocalId")}
	decoder := json.NewDecoder(r.Body)
	relationshipBAUnit, err := handler.RequiredRelationshipBAUnitCRUD.Read(unit1Id, unit2Id)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	var newRelationshipBAUnit ladm.LARequiredRelationshipBAUnit
	err = decoder.Decode(&newRelationshipBAUnit)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	newRelationshipBAUnit.Unit1 = relationshipBAUnit.(ladm.LARequiredRelationshipBAUnit).Unit1
	newRelationshipBAUnit.Unit2 = relationshipBAUnit.(ladm.LARequiredRelationshipBAUnit).Unit2
	updatedRelationshipBAUnit, err := handler.RequiredRelationshipBAUnitCRUD.Update(&newRelationshipBAUnit)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 200, updatedRelationshipBAUnit)
}

func (handler *RequiredRelationshipBAUnitHandler) DeleteRequiredRelationshipBAUnit(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	unit1Id := common.Oid{Namespace: p.ByName("unit1Namespace"), LocalID: p.ByName("unit1LocalId")}
	unit2Id := common.Oid{Namespace: p.ByName("unit2Namespace"), LocalID: p.ByName("unit2LocalId")}
	relationshipBAUnit, err := handler.RequiredRelationshipBAUnitCRUD.Read(unit1Id, unit2Id)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	err = handler.RequiredRelationshipBAUnitCRUD.Delete(relationshipBAUnit)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondEmpty(w, 204)
}
