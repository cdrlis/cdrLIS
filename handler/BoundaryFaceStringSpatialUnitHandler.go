package handler

import (
	"encoding/json"
	ladm "github.com/cdrlis/cdrLIS/LADM"
	"github.com/cdrlis/cdrLIS/LADM/common"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type BoundaryFaceStringSpatialUnitHandler struct {
	BfsSpatialUnitCRUD     CRUDer
	SpatialUnitCRUD        CRUDer
	BoundaryFaceStringCRUD CRUDer
}

func (handler *BoundaryFaceStringSpatialUnitHandler) GetBfsSpatialUnit(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	suId := common.Oid{ Namespace: p.ByName("suNamespace"), LocalID:p.ByName("suLocalId")}
	bfsId := common.Oid{ Namespace: p.ByName("bfsNamespace"), LocalID:p.ByName("bfsLocalId")}
	bfsSpatialUnit, err := handler.BfsSpatialUnitCRUD.Read(suId, bfsId)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	respondJSON(w, 200, bfsSpatialUnit)
}

func (handler *BoundaryFaceStringSpatialUnitHandler) GetBfsSpatialUnits(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bfsSpatialUnits, err := handler.BfsSpatialUnitCRUD.ReadAll()
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, bfsSpatialUnits)
}

func (handler *BoundaryFaceStringSpatialUnitHandler) CreateBfsSpatialUnit(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var bfsSpatialUnit ladm.BfsSpatialUnitPlus
	err := decoder.Decode(&bfsSpatialUnit)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	spatialUnit, err := handler.SpatialUnitCRUD.Read(bfsSpatialUnit.Su.SuID)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	bfs, err := handler.BoundaryFaceStringCRUD.Read(bfsSpatialUnit.Bfs.BfsID)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	bfsSpatialUnit.Su.BeginLifespanVersion = spatialUnit.(ladm.LASpatialUnit).BeginLifespanVersion
	bfsSpatialUnit.Bfs.BeginLifespanVersion = bfs.(ladm.LABoundaryFaceString).BeginLifespanVersion
	createdBfsSpatialUnit, err := handler.BfsSpatialUnitCRUD.Create(bfsSpatialUnit)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 201, createdBfsSpatialUnit)
}

func (handler *BoundaryFaceStringSpatialUnitHandler) UpdateBfsSpatialUnit(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	suId := common.Oid{ Namespace: p.ByName("suNamespace"), LocalID:p.ByName("suLocalId")}
	bfsId := common.Oid{ Namespace: p.ByName("bfsNamespace"), LocalID:p.ByName("bfsLocalId")}
	decoder := json.NewDecoder(r.Body)
	bfsSpatialUnit, err := handler.BfsSpatialUnitCRUD.Read(suId, bfsId)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	var newBfsSpatialUnit ladm.BfsSpatialUnitPlus
	err = decoder.Decode(&newBfsSpatialUnit)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	newBfsSpatialUnit.Su = bfsSpatialUnit.(ladm.BfsSpatialUnitPlus).Su
	newBfsSpatialUnit.Bfs = bfsSpatialUnit.(ladm.BfsSpatialUnitPlus).Bfs
	handler.BfsSpatialUnitCRUD.Update(&newBfsSpatialUnit)
	respondJSON(w, 200, newBfsSpatialUnit)
}

func (handler *BoundaryFaceStringSpatialUnitHandler) DeleteBfsSpatialUnit(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	suId := common.Oid{ Namespace: p.ByName("suNamespace"), LocalID:p.ByName("suLocalId")}
	bfsId := common.Oid{ Namespace: p.ByName("bfsNamespace"), LocalID:p.ByName("bfsLocalId")}
	bfsSpatialUnit, err := handler.BfsSpatialUnitCRUD.Read(suId, bfsId)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	handler.BfsSpatialUnitCRUD.Delete(bfsSpatialUnit)
	respondEmpty(w, 204)
}