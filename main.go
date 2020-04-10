package main

import (
	"github.com/cdrlis/cdrLIS/logic"
	"github.com/cdrlis/cdrLIS/repositories"

	"net/http"

	"github.com/cdrlis/cdrLIS/handlers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=ladm password=123456vV sslmode=disable")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	ladmRepository := repositories.LadmRepository{DB: db}
	service := logic.LAPartyService{Context: ladmRepository}
	partyHandler := handlers.PartyHandler{Service: service}

	http.HandleFunc("/party", partyHandler.GetParty)
	http.HandleFunc("/party/list", partyHandler.GetParties)
	http.HandleFunc("/party/update", partyHandler.UpdateParty)
	http.ListenAndServe(":3000", nil)
}
