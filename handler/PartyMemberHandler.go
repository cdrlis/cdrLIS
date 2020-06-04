package handler

import (
	"encoding/json"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PartyMemberHandler struct {
	PartyMemberCRUD CRUDer
	PartyCRUD       CRUDer
}

func (handler *PartyMemberHandler) GetPartyMember(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pid := common.Oid{ Namespace: p.ByName("partyNamespace"), LocalID:p.ByName("partyLocalId")}
	gid := common.Oid{ Namespace: p.ByName("groupNamespace"), LocalID:p.ByName("groupLocalId")}
	partyMember, err := handler.PartyMemberCRUD.Read(pid, gid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	respondJSON(w, 200, partyMember)
}

func (handler *PartyMemberHandler) GetPartyMembers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	partyMembers, err := handler.PartyMemberCRUD.ReadAll()
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	respondJSON(w, 200, partyMembers)
}

func (handler *PartyMemberHandler) CreatePartyMember(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var partyMember ladm.LAPartyMember
	err := decoder.Decode(&partyMember)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	partyInterface, err := handler.PartyCRUD.Read(partyMember.Party.PID)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	groupInterface, err := handler.PartyCRUD.Read(partyMember.Group.PID)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	party := partyInterface.(ladm.LAParty)
	group := groupInterface.(ladm.LAParty).GroupParty
	partyMember.Party = &party
	partyMember.Group = group
	createdPartyMember, err := handler.PartyMemberCRUD.Create(partyMember)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 201, createdPartyMember)
}

func (handler *PartyMemberHandler) UpdatePartyMember(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pid := common.Oid{ Namespace: p.ByName("partyNamespace"), LocalID:p.ByName("partyLocalId")}
	gid := common.Oid{ Namespace: p.ByName("groupNamespace"), LocalID:p.ByName("groupLocalId")}
	decoder := json.NewDecoder(r.Body)
	partyMember, err := handler.PartyMemberCRUD.Read(pid, gid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	var newPartyMember ladm.LAPartyMember
	err = decoder.Decode(&newPartyMember)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	newPartyMember.Party = partyMember.(ladm.LAPartyMember).Party
	newPartyMember.Group = partyMember.(ladm.LAPartyMember).Group
	handler.PartyMemberCRUD.Update(&newPartyMember)
	respondJSON(w, 200, newPartyMember)
}

func (handler *PartyMemberHandler) DeletePartyMember(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pid := common.Oid{ Namespace: p.ByName("partyNamespace"), LocalID:p.ByName("partyLocalId")}
	gid := common.Oid{ Namespace: p.ByName("groupNamespace"), LocalID:p.ByName("groupLocalId")}
	partyMember, err := handler.PartyMemberCRUD.Read(pid, gid)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	handler.PartyMemberCRUD.Delete(partyMember)
	respondEmpty(w, 204)
}