package main

import (
	"github.com/cdrlis/cdrLIS/crud"
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

	partyCRUD := crud.LAPartyCRUD{DB: db}
	levelCRUD := crud.LALevelCRUD{DB: db}
	sunitCRUD := crud.LASpatialUnitCRUD{DB: db}
	baunitCRUD := crud.LABAUnitCRUD{DB: db}

	partyHandler := handler.PartyHandler{PartyCRUD: partyCRUD}
	baunitHandler := handler.BAUnitHandler{BAUnitCRUD: baunitCRUD}
	sunitHandler := handler.SpatialUnitHandler{SpatialUnitCRUD: sunitCRUD, LevelCRUD: levelCRUD}
	levelHandler := handler.LevelHandler{LevelCRUD: levelCRUD}

	router := httprouter.New()

	router.GET("/party", partyHandler.GetParties)
	router.POST("/party", partyHandler.CreateParty)
	router.GET("/party/:namespace/:localId", partyHandler.GetParty)
	router.PUT("/party/:namespace/:localId", partyHandler.UpdateParty)
	router.DELETE("/party/:namespace/:localId", partyHandler.DeleteParty)

	router.GET("/baunit", baunitHandler.GetBAUnits)
	router.POST("/baunit", baunitHandler.CreateBAUnit)
	router.GET("/baunit/:namespace/:localId", baunitHandler.GetBAUnit)
	router.PUT("/baunit/:namespace/:localId", baunitHandler.UpdateBAUnit)
	router.DELETE("/baunit/:namespace/:localId", baunitHandler.DeleteBAUnit)

	router.GET("/spatialunit", sunitHandler.GetSpatialUnits)
	router.POST("/spatialunit", sunitHandler.CreateSpatialUnit)
	router.GET("/spatialunit/:namespace/:localId", sunitHandler.GetSpatialUnit)
	router.PUT("/spatialunit/:namespace/:localId", sunitHandler.UpdateSpatialUnit)
	router.DELETE("/spatialunit/:namespace/:localId", sunitHandler.DeleteSpatialUnit)

	router.GET("/level", levelHandler.GetLevels)
	router.POST("/level", levelHandler.CreateLevel)
	router.GET("/level/:namespace/:localId", levelHandler.GetLevel)
	router.PUT("/level/:namespace/:localId", levelHandler.UpdateLevel)
	router.DELETE("/level/:namespace/:localId", levelHandler.DeleteLevel)

	http.ListenAndServe(":3000", router)

}
