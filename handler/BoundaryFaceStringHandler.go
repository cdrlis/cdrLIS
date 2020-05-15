package handler

import (
	"encoding/json"
	ladm "github.com/cdrlis/cdrLIS/LADM"
	"github.com/cdrlis/cdrLIS/LADM/common"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type BoundaryFaceStringHandler struct {
	BoundaryFaceStringCRUD CRUDer
}

func (handler *BoundaryFaceStringHandler) GetBoundaryFaceString(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bfsid := common.Oid{Namespace: p.ByName("namespace"), LocalID: p.ByName("localId")}
	boundaryFaceString, err := handler.BoundaryFaceStringCRUD.Read(bfsid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	respondJSON(w, 200, boundaryFaceString)
}

func (handler *BoundaryFaceStringHandler) GetBoundaryFaceStrings(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	boundaryFaceStrings, err := handler.BoundaryFaceStringCRUD.ReadAll()
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, boundaryFaceStrings)
}

func (handler *BoundaryFaceStringHandler) CreateBoundaryFaceString(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var boundaryFaceString ladm.LABoundaryFaceString
	err := decoder.Decode(&boundaryFaceString)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	createdBoundaryFaceString, err := handler.BoundaryFaceStringCRUD.Create(boundaryFaceString)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 201, createdBoundaryFaceString)
}

func (handler *BoundaryFaceStringHandler) UpdateBoundaryFaceString(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bfsid := common.Oid{Namespace: p.ByName("namespace"), LocalID: p.ByName("localId")}
	decoder := json.NewDecoder(r.Body)
	_, err := handler.BoundaryFaceStringCRUD.Read(bfsid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	var newBoundaryFaceString ladm.LABoundaryFaceString
	err = decoder.Decode(&newBoundaryFaceString)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	handler.BoundaryFaceStringCRUD.Update(&newBoundaryFaceString)
	respondJSON(w, 200, newBoundaryFaceString)
}

func (handler *BoundaryFaceStringHandler) DeleteBoundaryFaceString(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bfsid := common.Oid{Namespace: p.ByName("namespace"), LocalID: p.ByName("localId")}
	boundaryFaceString, err := handler.BoundaryFaceStringCRUD.Read(bfsid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	handler.BoundaryFaceStringCRUD.Delete(boundaryFaceString)
	respondEmpty(w, 204)
}
