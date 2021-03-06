package handler

import (
	"encoding/json"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type SpatialUnitHandler struct {
	SpatialUnitCRUD CRUDer
	LevelCRUD       CRUDer
}

func (handler *SpatialUnitHandler) GetSpatialUnit(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uid := common.Oid{Namespace: p.ByName("namespace"), LocalID: p.ByName("localId")}
	suUnit, err := handler.SpatialUnitCRUD.Read(uid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	respondJSON(w, 200, suUnit)
}

func (handler *SpatialUnitHandler) GetSpatialUnits(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	baUnits, err := handler.SpatialUnitCRUD.ReadAll()
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, baUnits)
}

func (handler *SpatialUnitHandler) CreateSpatialUnit(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var spatialUnit ladm.LASpatialUnit
	err := decoder.Decode(&spatialUnit)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	level, err := handler.LevelCRUD.Read(spatialUnit.Level.LID)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	spatialUnit.LevelID = level.(ladm.LALevel).ID
	spatialUnit.LevelBeginLifespanVersion = level.(ladm.LALevel).BeginLifespanVersion
	createdSpatialUnit, err := handler.SpatialUnitCRUD.Create(spatialUnit)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 201, createdSpatialUnit)
}

func (handler *SpatialUnitHandler) UpdateSpatialUnit(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	suid := common.Oid{Namespace: p.ByName("namespace"), LocalID: p.ByName("localId")}
	decoder := json.NewDecoder(r.Body)
	_, err := handler.SpatialUnitCRUD.Read(suid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	var newSpatialUnit ladm.LASpatialUnit
	err = decoder.Decode(&newSpatialUnit)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	newSpatialUnit.SuID = suid
	updatedSpatialUnit, err := handler.SpatialUnitCRUD.Update(&newSpatialUnit)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 200, updatedSpatialUnit)
}

func (handler *SpatialUnitHandler) DeleteSpatialUnit(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	suid := common.Oid{Namespace: p.ByName("namespace"), LocalID: p.ByName("localId")}
	spatialUnit, err := handler.SpatialUnitCRUD.Read(suid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	err = handler.SpatialUnitCRUD.Delete(spatialUnit)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondEmpty(w, 204)
}

func (handler *SpatialUnitHandler) GetSpatialUnitGeometry(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uid := common.Oid{Namespace: p.ByName("namespace"), LocalID: p.ByName("localId")}
	suUnit, err := handler.SpatialUnitCRUD.Read(uid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	respondJSON(w, 200, suUnit.(ladm.LASpatialUnit).CreateArea())
}

func (handler *SpatialUnitHandler) GetSpatialUnitArea(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uid := common.Oid{Namespace: p.ByName("namespace"), LocalID: p.ByName("localId")}
	suUnit, err := handler.SpatialUnitCRUD.Read(uid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	respondJSON(w, 200, suUnit.(ladm.LASpatialUnit).ComputeArea())
}