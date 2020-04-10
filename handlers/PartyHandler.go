package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cdrlis/cdrLIS/logic"
)

type PartyHandler struct {
	Service logic.LAPartyService
}

func (handler *PartyHandler) GetParty(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	party, err := handler.Service.GetParty()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	fmt.Fprintf(w, "Party(%s), ExtPid(%s.%s)", *party.Name, party.ExtPid.Namespace, party.ExtPid.LocalID)
}

func (handler *PartyHandler) GetParties(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	parties, err := handler.Service.GetPartyList()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	var partyList []string
	for _, party := range *parties {
		name := "-"
		if party.Name != nil {
			name = *party.Name
		}
		partyList = append(partyList, name)
	}
	fmt.Fprintf(w, "Party list ( count %d ):\n%s", len(*parties), strings.Join(partyList, "\n"))
}

func (handler *PartyHandler) UpdateParty(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	party, err := handler.Service.GetParty()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	newName := *party.Name + "1"
	party.Name = &newName
	handler.Service.UpdateParty(*party)
	fmt.Fprintf(w, "Updated \"%s\"", *party.Name)
}
