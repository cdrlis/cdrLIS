package handler

import (
	"encoding/json"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PointHandler struct {
	PointCRUD CRUDer
}

func (handler *PointHandler) GetPoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	lid := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	point, err := handler.PointCRUD.Read(lid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	respondJSON(w, 200, point)
}

func (handler *PointHandler) GetPoints(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	points, err := handler.PointCRUD.ReadAll()
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, points)
}

func (handler *PointHandler) CreatePoint(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var point ladm.LAPoint
	err := decoder.Decode(&point)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	createdPoint, err := handler.PointCRUD.Create(point)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 201, createdPoint)
}

func (handler *PointHandler) UpdatePoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pid := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	decoder := json.NewDecoder(r.Body)
	_, err := handler.PointCRUD.Read(pid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	var newPoint ladm.LAPoint
	err = decoder.Decode(&newPoint)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	newPoint.PID = pid
	updatedPoint ,err := handler.PointCRUD.Update(&newPoint)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 200, updatedPoint)
}

func (handler *PointHandler) DeletePoint(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	lid := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	point, err := handler.PointCRUD.Read(lid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	err = handler.PointCRUD.Delete(point)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondEmpty(w, 204)
}