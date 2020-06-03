package main

import (
	"encoding/json"
	"fmt"
	"github.com/cdrlis/cdrLIS/crud"
	"github.com/rs/cors"
	"net/http"
	"os"

	"github.com/cdrlis/cdrLIS/handler"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
)

func main() {
	config := LoadConfiguration("config.json")
	fmt.Println(config.Database.ConnectionString())
	// config.json (PostgreSQL)
	db, err := gorm.Open(config.Database.Dialect, config.Database.ConnectionString())
	// YugabyteDB
	// db, err := gorm.Open("postgres", "host=localhost port=5433 user=yugabyte dbname=yugabyte password=yugabyte sslmode=disable")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	db.LogMode(config.Database.Log)

	partyCRUD := crud.LAPartyCRUD{DB: db}
	groupPartyCRUD := crud.LAGroupPartyCRUD{DB: db}
	partyMemberCRUD := crud.LAPartyMemberCRUD{DB: db}
	levelCRUD := crud.LALevelCRUD{DB: db}
	sunitCRUD := crud.LASpatialUnitCRUD{DB: db}
	baunitCRUD := crud.LABAUnitCRUD{DB: db}
	boundaryFaceStringCRUD := crud.LABoundaryFaceStringCRUD{DB: db}
	bfsSpatialUnitPlusCRUD := crud.BfsSpatialUnitPlusCRUD{DB: db}
	bfsSpatialUnitMinusCRUD := crud.BfsSpatialUnitMinusCRUD{DB: db}

	partyHandler := handler.PartyHandler{PartyCRUD: partyCRUD}
	groupPartyHandler := handler.GroupPartyHandler{GroupPartyCRUD: groupPartyCRUD}
	partyMemberHandler := handler.PartyMemberHandler{PartyMemberCRUD: partyMemberCRUD, PartyCRUD: partyCRUD}
	baunitHandler := handler.BAUnitHandler{BAUnitCRUD: baunitCRUD}
	sunitHandler := handler.SpatialUnitHandler{SpatialUnitCRUD: sunitCRUD, LevelCRUD: levelCRUD}
	levelHandler := handler.LevelHandler{LevelCRUD: levelCRUD}
	boundaryFaceStringHandler := handler.BoundaryFaceStringHandler{BoundaryFaceStringCRUD: boundaryFaceStringCRUD}
	bfsSpatialUnitPlusHandler := handler.BoundaryFaceStringSpatialUnitPlusHandler{SpatialUnitCRUD: sunitCRUD, BoundaryFaceStringCRUD: boundaryFaceStringCRUD, BfsSpatialUnitCRUD: bfsSpatialUnitPlusCRUD}
	bfsSpatialUnitMinusHandler := handler.BoundaryFaceStringSpatialUnitMinusHandler{SpatialUnitCRUD: sunitCRUD, BoundaryFaceStringCRUD: boundaryFaceStringCRUD, BfsSpatialUnitCRUD: bfsSpatialUnitMinusCRUD}

	router := httprouter.New()

	router.GET("/party", partyHandler.GetParties)
	router.POST("/party", partyHandler.CreateParty)
	router.GET("/party/:namespace/:localId", partyHandler.GetParty)
	router.PUT("/party/:namespace/:localId", partyHandler.UpdateParty)
	router.DELETE("/party/:namespace/:localId", partyHandler.DeleteParty)

	router.GET("/groupparty", groupPartyHandler.GetParties)
	router.POST("/groupparty", groupPartyHandler.CreateParty)
	router.GET("/groupparty/:namespace/:localId", groupPartyHandler.GetParty)
	router.PUT("/groupparty/:namespace/:localId", groupPartyHandler.UpdateParty)
	router.DELETE("/groupparty/:namespace/:localId", groupPartyHandler.DeleteParty)

	router.GET("/partymember", partyMemberHandler.GetPartyMembers)
	router.POST("/partymember", partyMemberHandler.CreatePartyMember)
	router.GET("/partymember/:partyNamespace/:partyLocalId/:groupNamespace/:groupLocalId", partyMemberHandler.GetPartyMember)
	router.PUT("/partymember/:partyNamespace/:partyLocalId/:groupNamespace/:groupLocalId", partyMemberHandler.UpdatePartyMember)
	router.DELETE("/partymember/:partyNamespace/:partyLocalId/:groupNamespace/:groupLocalId", partyMemberHandler.DeletePartyMember)

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

	router.GET("/spatialunit/:namespace/:localId/geometry", sunitHandler.GetSpatialUnitGeometry)
	router.GET("/spatialunit/:namespace/:localId/area", sunitHandler.GetSpatialUnitArea)

	router.GET("/level", levelHandler.GetLevels)
	router.POST("/level", levelHandler.CreateLevel)
	router.GET("/level/:namespace/:localId", levelHandler.GetLevel)
	router.PUT("/level/:namespace/:localId", levelHandler.UpdateLevel)
	router.DELETE("/level/:namespace/:localId", levelHandler.DeleteLevel)

	router.GET("/boundaryfacestring", boundaryFaceStringHandler.GetBoundaryFaceStrings)
	router.POST("/boundaryfacestring", boundaryFaceStringHandler.CreateBoundaryFaceString)
	router.GET("/boundaryfacestring/:namespace/:localId", boundaryFaceStringHandler.GetBoundaryFaceString)
	router.PUT("/boundaryfacestring/:namespace/:localId", boundaryFaceStringHandler.UpdateBoundaryFaceString)
	router.DELETE("/boundaryfacestring/:namespace/:localId", boundaryFaceStringHandler.DeleteBoundaryFaceString)

	router.GET("/boundaryfacestring-plus", bfsSpatialUnitPlusHandler.GetBfsSpatialUnits)
	router.POST("/boundaryfacestring-plus", bfsSpatialUnitPlusHandler.CreateBfsSpatialUnit)
	router.GET("/boundaryfacestring-plus/:suNamespace/:suLocalId/:bfsNamespace/:bfsLocalId", bfsSpatialUnitPlusHandler.GetBfsSpatialUnit)
	router.PUT("/boundaryfacestring-plus/:suNamespace/:suLocalId/:bfsNamespace/:bfsLocalId", bfsSpatialUnitPlusHandler.UpdateBfsSpatialUnit)
	router.DELETE("/boundaryfacestring-plus/:suNamespace/:suLocalId/:bfsNamespace/:bfsLocalId", bfsSpatialUnitPlusHandler.DeleteBfsSpatialUnit)

	router.GET("/boundaryfacestring-minus", bfsSpatialUnitMinusHandler.GetBfsSpatialUnits)
	router.POST("/boundaryfacestring-minus", bfsSpatialUnitMinusHandler.CreateBfsSpatialUnit)
	router.GET("/boundaryfacestring-minus/:suNamespace/:suLocalId/:bfsNamespace/:bfsLocalId", bfsSpatialUnitMinusHandler.GetBfsSpatialUnit)
	router.PUT("/boundaryfacestring-minus/:suNamespace/:suLocalId/:bfsNamespace/:bfsLocalId", bfsSpatialUnitMinusHandler.UpdateBfsSpatialUnit)
	router.DELETE("/boundaryfacestring-minus/:suNamespace/:suLocalId/:bfsNamespace/:bfsLocalId", bfsSpatialUnitMinusHandler.DeleteBfsSpatialUnit)

	handler := cors.Default().Handler(router)
	http.ListenAndServe(":3000", handler)

}

type DatabaseSettings struct {
	Host              string
	Port              int
	User              string
	DbName            string
	Password          string
	AdditionalOptions string
	Dialect           string
	Log               bool
}

func (settings DatabaseSettings) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s %s", settings.Host, settings.Port,
		settings.User, settings.DbName, settings.Password, settings.AdditionalOptions)
}

type Configuration struct {
	Database DatabaseSettings
}

func LoadConfiguration(configFilePath string) Configuration {
	configFile, err := os.Open(configFilePath)
	defer configFile.Close()
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(configFile)
	var config Configuration
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	return config
}
