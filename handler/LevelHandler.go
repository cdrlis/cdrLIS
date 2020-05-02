package handler

import (
	"encoding/json"
	ladm "github.com/cdrlis/cdrLIS/LADM"
	"github.com/cdrlis/cdrLIS/LADM/common"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type LevelHandler struct {
	CRUD CRUDer
}

func (handler *LevelHandler) GetLevel(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	lid := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	level, err := handler.CRUD.Read(lid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	respondJSON(w, 200, level)
}

func (handler *LevelHandler) GetLevels(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	levels, err := handler.CRUD.ReadAll()
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, levels)
}

func (handler *LevelHandler) CreateLevel(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var level ladm.LALevel
	err := decoder.Decode(&level)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	createdLevel, err := handler.CRUD.Create(level)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 201, createdLevel)
}

func (handler *LevelHandler) UpdateLevel(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	lid := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	decoder := json.NewDecoder(r.Body)
	_, err := handler.CRUD.Read(lid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	var newLevel ladm.LALevel
	err = decoder.Decode(&newLevel)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	handler.CRUD.Update(&newLevel)
	respondJSON(w, 200, newLevel)
}

func (handler *LevelHandler) DeleteLevel(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	lid := common.Oid{ Namespace: p.ByName("namespace"), LocalID:p.ByName("localId")}
	level, err := handler.CRUD.Read(lid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	handler.CRUD.Delete(level)
	respondEmpty(w, 204)
}