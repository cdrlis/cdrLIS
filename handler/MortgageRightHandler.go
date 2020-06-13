package handler

import (
	"encoding/json"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type MortgageRightHandler struct {
	MortgageRightCRUD CRUDer
	RrrCRUD         CRUDer
}

func (handler *MortgageRightHandler) GetMortgageRight(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uId := common.Oid{Namespace: p.ByName("mortgageNamespace"), LocalID: p.ByName("mortgageLocalId")}
	pId := common.Oid{Namespace: p.ByName("rightNamespace"), LocalID: p.ByName("rightLocalId")}
	mortgageRight, err := handler.MortgageRightCRUD.Read(uId, pId)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	respondJSON(w, 200, mortgageRight)
}

func (handler *MortgageRightHandler) GetMortgageRights(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	mortgageRights, err := handler.MortgageRightCRUD.ReadAll()
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, mortgageRights)
}

func (handler *MortgageRightHandler) CreateMortgageRight(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var mortgageRight ladm.MortgageRight
	err := decoder.Decode(&mortgageRight)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	mortgage, err := handler.RrrCRUD.Read(mortgageRight.Mortgage.RID)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	if mortgage.(ladm.LARRR).Restriction == nil || mortgage.(ladm.LARRR).Restriction.Mortgage == nil{
		respondError(w, 404, "Mortgage not found")
	}
	right, err := handler.RrrCRUD.Read(mortgageRight.Right.RID)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	if right.(ladm.LARRR).Right == nil{
		respondError(w, 404, "Mortgage not found")
	}
	mortgageRight.Mortgage.BeginLifespanVersion = mortgage.(ladm.LARRR).Restriction.Mortgage.BeginLifespanVersion
	mortgageRight.Right.BeginLifespanVersion = right.(ladm.LARRR).Right.BeginLifespanVersion
	createdMortgageRight, err := handler.MortgageRightCRUD.Create(mortgageRight)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 201, createdMortgageRight)
}

func (handler *MortgageRightHandler) UpdateMortgageRight(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pId := common.Oid{Namespace: p.ByName("mortgageNamespace"), LocalID: p.ByName("mortgageLocalId")}
	rightId := common.Oid{Namespace: p.ByName("rightNamespace"), LocalID: p.ByName("rightLocalId")}
	decoder := json.NewDecoder(r.Body)
	mortgageRight, err := handler.MortgageRightCRUD.Read(pId, rightId)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	var newMortgageRight ladm.MortgageRight
	err = decoder.Decode(&newMortgageRight)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	newMortgageRight.Mortgage = mortgageRight.(ladm.MortgageRight).Mortgage
	newMortgageRight.Right = mortgageRight.(ladm.MortgageRight).Right
	updatedMortgageRight, err := handler.MortgageRightCRUD.Update(&newMortgageRight)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 200, updatedMortgageRight)
}

func (handler *MortgageRightHandler) DeleteMortgageRight(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uId := common.Oid{Namespace: p.ByName("mortgageNamespace"), LocalID: p.ByName("mortgageLocalId")}
	pId := common.Oid{Namespace: p.ByName("rightNamespace"), LocalID: p.ByName("rightLocalId")}
	mortgageRight, err := handler.MortgageRightCRUD.Read(uId, pId)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	err = handler.MortgageRightCRUD.Delete(mortgageRight)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondEmpty(w, 204)
}
