package main

import (
	"github.com/cdrlis/cdrLIS/logic"
	"github.com/cdrlis/cdrLIS/repositories"

	"net/http"

	"github.com/cdrlis/cdrLIS/handler"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
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
	partyHandler := handler.PartyHandler{Service: service}

	router := httprouter.New()

	router.GET("/party", partyHandler.GetParties)
	router.POST("/party", partyHandler.CreateParty)
	router.GET("/party/:id", partyHandler.GetParty)
	router.PUT("/party/:id", partyHandler.UpdateParty)
	router.DELETE("/party/:id", partyHandler.DeleteParty)

	http.ListenAndServe(":3000", router)
}
