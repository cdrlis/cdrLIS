package handler

import (
	"encoding/json"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PointBfsHandler struct {
	PointBfsCRUD           CRUDer
	PointCRUD              CRUDer
	BoundaryFaceStringCRUD CRUDer
}

func (handler *PointBfsHandler) GetPointBfs(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pId := common.Oid{Namespace: p.ByName("pointNamespace"), LocalID: p.ByName("pointLocalId")}
	bfsId := common.Oid{Namespace: p.ByName("bfsNamespace"), LocalID: p.ByName("bfsLocalId")}
	pointBfs, err := handler.PointBfsCRUD.Read(pId, bfsId)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	respondJSON(w, 200, pointBfs)
}

func (handler *PointBfsHandler) GetPointBfss(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	pointBfss, err := handler.PointBfsCRUD.ReadAll()
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, pointBfss)
}

func (handler *PointBfsHandler) CreatePointBfs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var pointBfs ladm.PointBfs
	err := decoder.Decode(&pointBfs)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	point, err := handler.PointCRUD.Read(pointBfs.Point.PID)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	bfs, err := handler.BoundaryFaceStringCRUD.Read(pointBfs.Bfs.BfsID)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	pointBfs.Point.BeginLifespanVersion = point.(ladm.LASpatialUnit).BeginLifespanVersion
	pointBfs.Bfs.BeginLifespanVersion = bfs.(ladm.LABoundaryFaceString).BeginLifespanVersion
	createdPointBfs, err := handler.PointBfsCRUD.Create(pointBfs)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 201, createdPointBfs)
}

func (handler *PointBfsHandler) UpdatePointBfs(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pId := common.Oid{Namespace: p.ByName("pointNamespace"), LocalID: p.ByName("pointLocalId")}
	bfsId := common.Oid{Namespace: p.ByName("bfsNamespace"), LocalID: p.ByName("bfsLocalId")}
	decoder := json.NewDecoder(r.Body)
	pointBfs, err := handler.PointBfsCRUD.Read(pId, bfsId)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	var newPointBfs ladm.PointBfs
	err = decoder.Decode(&newPointBfs)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	newPointBfs.Point = pointBfs.(ladm.PointBfs).Point
	newPointBfs.Bfs = pointBfs.(ladm.PointBfs).Bfs
	handler.PointBfsCRUD.Update(&newPointBfs)
	respondJSON(w, 200, newPointBfs)
}

func (handler *PointBfsHandler) DeletePointBfs(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pId := common.Oid{Namespace: p.ByName("pointNamespace"), LocalID: p.ByName("pointLocalId")}
	bfsId := common.Oid{Namespace: p.ByName("bfsNamespace"), LocalID: p.ByName("bfsLocalId")}
	pointBfs, err := handler.PointBfsCRUD.Read(pId, bfsId)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	handler.PointBfsCRUD.Delete(pointBfs)
	respondEmpty(w, 204)
}
