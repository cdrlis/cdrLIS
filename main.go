package main

import (
	"github.com/cdrlis/cdrLIS/logic"
	"github.com/cdrlis/cdrLIS/dbcontext"

	"net/http"

	"github.com/cdrlis/cdrLIS/handler"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// PostgreSQL
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=ladm password=123456vV sslmode=disable")
	// YugabyteDB
	// db, err := gorm.Open("postgres", "host=localhost port=5433 user=yugabyte dbname=yugabyte password=yugabyte sslmode=disable")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	ladmDB := dbcontext.CRUDContext{DB: db}
	service := logic.LAPartyService{Context: ladmDB}
	partyHandler := handler.PartyHandler{Service: service}

	router := httprouter.New()

	router.GET("/party", partyHandler.GetParties)
	router.POST("/party", partyHandler.CreateParty)
	router.GET("/party/:namespace/:localId", partyHandler.GetParty)
	router.PUT("/party/:namespace/:localId", partyHandler.UpdateParty)
	router.DELETE("/party/:namespace/:localId", partyHandler.DeleteParty)
	
	router.GET("/type/party", partyHandler.GetPartyTypes)
	router.GET("/role/party", partyHandler.GetPartyRoles)

	router.GET("/groupParty", partyHandler.GetGroupParties)

	http.ListenAndServe(":3000", router)

}
