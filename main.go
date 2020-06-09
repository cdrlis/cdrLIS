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
	partyMemberCRUD := crud.LAPartyMemberCRUD{DB: db}
	baunitCRUD := crud.LABAUnitCRUD{DB: db}
	baUnitAsPartyCRUD := crud.BAUnitAsPartyCRUD{DB: db}
	rrrCRUD := crud.LARRRCRUD{DB: db}
	levelCRUD := crud.LALevelCRUD{DB: db}
	sunitCRUD := crud.LASpatialUnitCRUD{DB: db}
	suHierarchyCRUD := crud.SuHierarchyCRUD{DB: db}
	suGroupCRUD := crud.LASpatialUnitGroupCRUD{DB: db}
	suGroupHierarchyCRUD := crud.SuGroupHierarchyCRUD{DB: db}
	suSuGroupCRUD := crud.SuSuGroupCRUD{DB: db}
	boundaryFaceStringCRUD := crud.LABoundaryFaceStringCRUD{DB: db}
	bfsSpatialUnitPlusCRUD := crud.BfsSpatialUnitPlusCRUD{DB: db}
	bfsSpatialUnitMinusCRUD := crud.BfsSpatialUnitMinusCRUD{DB: db}
	pointCRUD := crud.LAPointCRUD{DB: db}
	pointBfsCRUD := crud.PointBfsCRUD{DB: db}

	partyHandler := handler.PartyHandler{PartyCRUD: partyCRUD}
	partyMemberHandler := handler.PartyMemberHandler{PartyMemberCRUD: partyMemberCRUD, PartyCRUD: partyCRUD}
	baunitHandler := handler.BAUnitHandler{BAUnitCRUD: baunitCRUD}
	baUnitAsPartyHandler := handler.BAUnitAsPartyHandler{BAUnitAsPartyCRUD: baUnitAsPartyCRUD, BAUnitCRUD: baunitCRUD, PartyCRUD: partyCRUD}
	rrrHandler := handler.RrrHandler{RrrCRUD: rrrCRUD, PartyCRUD: partyCRUD, BAUnitCRUD: baunitCRUD}
	levelHandler := handler.LevelHandler{LevelCRUD: levelCRUD}
	sunitHandler := handler.SpatialUnitHandler{SpatialUnitCRUD: sunitCRUD, LevelCRUD: levelCRUD}
	suHierarchyHandler := handler.SuHierarchyHandler{SuHierarchyCRUD: suHierarchyCRUD, SpatialUnitCRUD: sunitCRUD}
	sunitGroupHandler := handler.SpatialUnitGroupHandler{SuGroupCRUD: suGroupCRUD}
	suGroupHierarchyHandler := handler.SuGroupHierarchyHandler{SuGroupHierarchyCRUD: suGroupHierarchyCRUD, SpatialUnitGroupCRUD: suGroupCRUD}
	suSuGroupHandler := handler.SuSuGroupHandler{SuSuGroupCRUD: suSuGroupCRUD, SuGroupCRUD: suGroupCRUD, SuCRUD: sunitCRUD}
	boundaryFaceStringHandler := handler.BoundaryFaceStringHandler{BoundaryFaceStringCRUD: boundaryFaceStringCRUD}
	bfsSpatialUnitPlusHandler := handler.BoundaryFaceStringSpatialUnitPlusHandler{SpatialUnitCRUD: sunitCRUD, BoundaryFaceStringCRUD: boundaryFaceStringCRUD, BfsSpatialUnitCRUD: bfsSpatialUnitPlusCRUD}
	bfsSpatialUnitMinusHandler := handler.BoundaryFaceStringSpatialUnitMinusHandler{SpatialUnitCRUD: sunitCRUD, BoundaryFaceStringCRUD: boundaryFaceStringCRUD, BfsSpatialUnitCRUD: bfsSpatialUnitMinusCRUD}
	pointHandler := handler.PointHandler{PointCRUD: pointCRUD}
	pointBfsHandler := handler.PointBfsHandler{PointBfsCRUD: pointBfsCRUD, PointCRUD: pointCRUD, BoundaryFaceStringCRUD: boundaryFaceStringCRUD}

	router := httprouter.New()

	router.GET("/party", partyHandler.GetParties)
	router.POST("/party", partyHandler.CreateParty)
	router.GET("/party/:namespace/:localId", partyHandler.GetParty)
	router.PUT("/party/:namespace/:localId", partyHandler.UpdateParty)
	router.DELETE("/party/:namespace/:localId", partyHandler.DeleteParty)

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

	router.GET("/baunitasparty", baUnitAsPartyHandler.GetBAUnitAsPartys)
	router.POST("/baunitasparty", baUnitAsPartyHandler.CreateBAUnitAsParty)
	router.GET("/baunitasparty/:baunitNamespace/:baunitLocalId/:partyNamespace/:partyLocalId", baUnitAsPartyHandler.GetBAUnitAsParty)
	router.PUT("/baunitasparty/:baunitNamespace/:baunitLocalId/:partyNamespace/:partyLocalId", baUnitAsPartyHandler.UpdateBAUnitAsParty)
	router.DELETE("/baunitasparty/:baunitNamespace/:baunitLocalId/:partyNamespace/:partyLocalId", baUnitAsPartyHandler.DeleteBAUnitAsParty)

	router.GET("/rrr", rrrHandler.GetRrrs)
	router.POST("/rrr", rrrHandler.CreateRrr)
	router.GET("/rrr/:namespace/:localId", rrrHandler.GetRrr)
	router.PUT("/rrr/:namespace/:localId", rrrHandler.UpdateRrr)
	router.DELETE("/rrr/:namespace/:localId", rrrHandler.DeleteRrr)

	router.GET("/level", levelHandler.GetLevels)
	router.POST("/level", levelHandler.CreateLevel)
	router.GET("/level/:namespace/:localId", levelHandler.GetLevel)
	router.PUT("/level/:namespace/:localId", levelHandler.UpdateLevel)
	router.DELETE("/level/:namespace/:localId", levelHandler.DeleteLevel)

	router.GET("/spatialunit", sunitHandler.GetSpatialUnits)
	router.POST("/spatialunit", sunitHandler.CreateSpatialUnit)
	router.GET("/spatialunit/:namespace/:localId", sunitHandler.GetSpatialUnit)
	router.PUT("/spatialunit/:namespace/:localId", sunitHandler.UpdateSpatialUnit)
	router.DELETE("/spatialunit/:namespace/:localId", sunitHandler.DeleteSpatialUnit)

	router.GET("/spatialunit/:namespace/:localId/geometry", sunitHandler.GetSpatialUnitGeometry)
	router.GET("/spatialunit/:namespace/:localId/area", sunitHandler.GetSpatialUnitArea)

	router.GET("/suhierarchy", suHierarchyHandler.GetSuHierarchys)
	router.POST("/suhierarchy", suHierarchyHandler.CreateSuHierarchy)
	router.GET("/suhierarchy/:parentNamespace/:parentLocalId/:childNamespace/:childLocalId", suHierarchyHandler.GetSuHierarchy)
	router.PUT("/suhierarchy/:parentNamespace/:parentLocalId/:childNamespace/:childLocalId", suHierarchyHandler.UpdateSuHierarchy)
	router.DELETE("/suhierarchy/:parentNamespace/:parentLocalId/:childNamespace/:childLocalId", suHierarchyHandler.DeleteSuHierarchy)

	router.GET("/spatialunitgroup", sunitGroupHandler.GetSuGroups)
	router.POST("/spatialunitgroup", sunitGroupHandler.CreateSuGroup)
	router.GET("/spatialunitgroup/:namespace/:localId", sunitGroupHandler.GetSuGroup)
	router.PUT("/spatialunitgroup/:namespace/:localId", sunitGroupHandler.UpdateSuGroup)
	router.DELETE("/spatialunitgroup/:namespace/:localId", sunitGroupHandler.DeleteSuGroup)

	router.GET("/sugrouphierarchy", suGroupHierarchyHandler.GetSuGroupHierarchys)
	router.POST("/sugrouphierarchy", suGroupHierarchyHandler.CreateSuGroupHierarchy)
	router.GET("/sugrouphierarchy/:setNamespace/:setLocalId/:elementNamespace/:elementLocalId", suGroupHierarchyHandler.GetSuGroupHierarchy)
	router.PUT("/sugrouphierarchy/:setNamespace/:setLocalId/:elementNamespace/:elementLocalId", suGroupHierarchyHandler.UpdateSuGroupHierarchy)
	router.DELETE("/sugrouphierarchy/:setNamespace/:setLocalId/:elementNamespace/:elementLocalId", suGroupHierarchyHandler.DeleteSuGroupHierarchy)

	router.GET("/susugroup", suSuGroupHandler.GetSuSuGroups)
	router.POST("/susugroup", suSuGroupHandler.CreateSuSuGroup)
	router.GET("/susugroup/:sugNamespace/:sugLocalId/:suNamespace/:suLocalId", suSuGroupHandler.GetSuSuGroup)
	router.PUT("/susugroup/:sugNamespace/:sugLocalId/:suNamespace/:suLocalId", suSuGroupHandler.UpdateSuSuGroup)
	router.DELETE("/susugroup/:sugNamespace/:sugLocalId/:suNamespace/:suLocalId", suSuGroupHandler.DeleteSuSuGroup)

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

	router.GET("/point", pointHandler.GetPoints)
	router.POST("/point", pointHandler.CreatePoint)
	router.GET("/point/:namespace/:localId", pointHandler.GetPoint)
	router.PUT("/point/:namespace/:localId", pointHandler.UpdatePoint)
	router.DELETE("/level/:namespace/:localId", pointHandler.DeletePoint)

	router.GET("/pointbfs", pointBfsHandler.GetPointBfss)
	router.POST("/pointbfs", pointBfsHandler.CreatePointBfs)
	router.GET("/pointbfs/:sugNamespace/:sugLocalId/:suNamespace/:suLocalId", pointBfsHandler.GetPointBfs)
	router.PUT("/pointbfs/:sugNamespace/:sugLocalId/:suNamespace/:suLocalId", pointBfsHandler.UpdatePointBfs)
	router.DELETE("/pointbfs/:sugNamespace/:sugLocalId/:suNamespace/:suLocalId", pointBfsHandler.DeletePointBfs)

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
