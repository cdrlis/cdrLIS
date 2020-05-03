package handler

import (
	"encoding/json"
	ladm "github.com/cdrlis/cdrLIS/LADM"
	"github.com/cdrlis/cdrLIS/LADM/common"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PartyMemberHandler struct {
	PartyMemberCRUD CRUDer
	PartyCRUD       CRUDer
	GroupPartyCRUD  CRUDer
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
	party, err := handler.PartyCRUD.Read(partyMember.Party.PID)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	group, err := handler.GroupPartyCRUD.Read(partyMember.Group.GroupID)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	partyMember.Party.BeginLifespanVersion = party.(ladm.LAParty).BeginLifespanVersion
	partyMember.Group.BeginLifespanVersion = group.(ladm.LAGroupParty).BeginLifespanVersion
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