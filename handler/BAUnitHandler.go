package handler

import (
	"encoding/json"
	ladm "github.com/cdrlis/cdrLIS/LADM"
	"github.com/cdrlis/cdrLIS/LADM/common"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type BAUnitHandler struct {
	BAUnitCRUD CRUDer
}

func (handler *BAUnitHandler) GetBAUnit(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uid := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	baUnit, err := handler.BAUnitCRUD.Read(uid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	respondJSON(w, 200, baUnit)
}

func (handler *BAUnitHandler) GetBAUnits(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	baUnits, err := handler.BAUnitCRUD.ReadAll()
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, baUnits)
}

func (handler *BAUnitHandler) CreateBAUnit(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var baunit ladm.LABAUnit
	err := decoder.Decode(&baunit)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	createdBaUnit, err := handler.BAUnitCRUD.Create(baunit)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 201, createdBaUnit)
}

func (handler *BAUnitHandler) UpdateBAUnit(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uid := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	decoder := json.NewDecoder(r.Body)
	_, err := handler.BAUnitCRUD.Read(uid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	var newBaUnit ladm.LABAUnit
	err = decoder.Decode(&newBaUnit)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	handler.BAUnitCRUD.Update(&newBaUnit)
	respondJSON(w, 200, newBaUnit)
}

func (handler *BAUnitHandler) DeleteBAUnit(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uid := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	baUnit, err := handler.BAUnitCRUD.Read(uid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	handler.BAUnitCRUD.Delete(baUnit)
	respondEmpty(w, 204)
}