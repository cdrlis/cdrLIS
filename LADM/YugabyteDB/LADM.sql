--CREATE USER "LADM" WITH
--	LOGIN
--	SUPERUSER
--	CREATEDB
--	CREATEROLE
--	INHERIT
--	REPLICATION
--	CONNECTION LIMIT -1
--	PASSWORD 'LADM';

CREATE SCHEMA IF NOT EXISTS "LADM" AUTHORIZATION yugabyte;
ALTER ROLE yugabyte IN DATABASE yugabyte SET search_path TO "LADM", public;

--DROP TABLE IF EXISTS "VersionedObject";

DROP TABLE IF EXISTS "baunitAsParty";
DROP TABLE IF EXISTS "LA_PartyMember";

DROP TABLE IF EXISTS "pointBfs";
DROP TABLE IF EXISTS "pointPb";
DROP TABLE IF EXISTS "LA_Point";
DROP TABLE IF EXISTS "minus";
DROP TABLE IF EXISTS "plus";
DROP TABLE IF EXISTS "LA_BoundaryFaceString";

DROP TYPE IF EXISTS "LA_InterpolationType";
DROP TYPE IF EXISTS "LA_MonumentationType";
DROP TYPE IF EXISTS "LA_PointType";
DROP TYPE IF EXISTS "LI_Lineage";
DROP TYPE IF EXISTS "LA_Transformation";
DROP TYPE IF EXISTS "CC_OperationMethod";
DROP TYPE IF EXISTS "CC_Formula";

DROP TABLE IF EXISTS "LA_LegalSpaceBuildingUnit";
DROP TABLE IF EXISTS "LA_RequiredRelationshipSpatialUnit";

DROP TABLE IF EXISTS "suBaunit";
DROP TABLE IF EXISTS "suSuGroup";
DROP TABLE IF EXISTS "suHierarchy";
DROP TABLE IF EXISTS "LA_SpatialUnit";
DROP TABLE IF EXISTS "suGroupHierarchy";
DROP TABLE IF EXISTS "LA_SpatialUnitGroup";
DROP TABLE IF EXISTS "LA_Level";

DROP TABLE IF EXISTS "LA_RequiredRelationshipBAUnit";
DROP TABLE IF EXISTS "LA_BAUnit" CASCADE;
DROP TYPE IF EXISTS "LA_BAUnitType";
DROP TABLE IF EXISTS "mortgageRight";
DROP TABLE IF EXISTS "LA_Mortgage";
DROP TYPE IF EXISTS "LA_MortgageType";
DROP TABLE IF EXISTS "LA_Right";
DROP TYPE IF EXISTS "LA_RightType";
DROP TABLE IF EXISTS "LA_Restriction";
DROP TYPE IF EXISTS "LA_RestrictionType";
DROP TABLE IF EXISTS "LA_Responsibility";
DROP TYPE IF EXISTS "LA_ResponsibilityType";
DROP TABLE IF EXISTS "LA_Party";
DROP TYPE IF EXISTS "LA_PartyType";
DROP TABLE IF EXISTS "LA_GroupParty";
DROP TYPE IF EXISTS "LA_PartyRoleType";
DROP TYPE IF EXISTS "LA_GroupPartyType";

--DROP TABLE IF EXISTS "LA_RRR";
DROP TYPE IF EXISTS "Fraction";

DROP TABLE IF EXISTS "PolygonBoundary";
DROP TABLE IF EXISTS "PolygonSpatialUnit";
DROP TABLE IF EXISTS "PolygonSpatialUnitGroup";
DROP TABLE IF EXISTS "PolygonLevel";

DROP TYPE IF EXISTS "CI_ResponsibleParty";
DROP TYPE IF EXISTS "Oid";
DROP TYPE IF EXISTS "CI_RoleCode";
DROP TYPE IF EXISTS "CI_Contact";
DROP TYPE IF EXISTS "CI_OnlineResource";
DROP TYPE IF EXISTS "CI_OnLineFunctionCode";
DROP TYPE IF EXISTS "CI_Address";
DROP TYPE IF EXISTS "CI_Telephone";

CREATE TYPE "Oid" AS (
                         localId		VARCHAR,
                         namespace  VARCHAR
                     );

--
-- "DQ_Element" -- Omitted for simplicity
--

--
-- "CI_ResponsibleParty"
--
CREATE TYPE "CI_Telephone" AS (
                                  voice 		VARCHAR,
                                  fascimile	VARCHAR
                              );
CREATE TYPE "CI_Address" AS (
                                deliveryPoint 			VARCHAR,
                                city					VARCHAR,
                                administrativeArea		VARCHAR,
                                postalCode				VARCHAR,
                                country					VARCHAR,
                                electronicMailAddress	VARCHAR
                            );

CREATE TYPE "CI_OnLineFunctionCode" AS ENUM (
    'download', 'information', 'offlineAccess', 'order', 'search'
    );

CREATE TYPE "CI_OnlineResource" AS (
                                       linkage				VARCHAR,
                                       protocol			VARCHAR,
                                       applicationProfile	VARCHAR,
                                       name				VARCHAR,
                                       description			VARCHAR,
                                       function			"CI_OnLineFunctionCode"
                                   );

CREATE TYPE "CI_Contact" AS (
                                phone				"CI_Telephone",
                                address				"CI_Address",
                                onlineResource		"CI_OnlineResource",
                                hoursOfService		VARCHAR(16),
                                contactInstructions	TEXT
                            );

CREATE TYPE "CI_RoleCode" AS ENUM (
    'resourceProvider', 'custodian', 'owner', 'user', 'distributor', 'originator',
    'pointOfContact', 'principalInvestigator', 'processor', 'publisher', 'author'
    );

CREATE TYPE "CI_ResponsibleParty" AS (
                                         individualName		VARCHAR,
                                         organisationName 	VARCHAR,
                                         positionName		VARCHAR,
                                         contactInfo			"CI_Contact",
                                         role				"CI_RoleCode"
                                     );
--
-- End of "ResponsibleParty"
--

--
-- Class VersionedObject is introduced in the LADM to manage and maintain historical data in the database.
-- History requires, that inserted and superseded data, are given a time-stamp. In this way, the contents of the
-- database can be reconstructed, as they were at any historical moment. For more on history and dynamic
-- aspects of LA systems, see Annex N.
-- NOTE: Inheritance is not yet supported in YugabyteDB
-- See https://github.com/YugaByte/yugabyte-db/issues/1129. Click '+' on the description to raise its priority
--
--CREATE TABLE "VersionedObject" (
--	beginLifeSpanVersion 	TIMESTAMP NOT NULL,
--	endLifeSpanVersion		TIMESTAMP DEFAULT '-infinity'::timestamp,
--	quality					"DQ_Element",
--	source 					"CI_ResponsibleParty"
--);

--
--***************************************************************************
-- Party package														    *
--***************************************************************************
--
-- Party::LA_Party
--
-- An instance of class LA_Party is a party. A party is associated to zero or more [0..*] instances of a subclass of
-- LA_RRR. LA_Party is also associated to LA_BAUnit, to cater for the fact that a basic administrative unit can
-- be a party (e.g. a basic administrative unit holding an easement on another basic administrative unit). A party
-- may be associated to zero or more [0..*] administrative sources (i.e. the author of a transfer document is
-- defined as a party playing the role of conveyancer in a source). A party may be associated to zero or more
-- [0..*] spatial sources (i.e. the author of a survey document is defined as a party playing the role of surveyor in
-- a source); see Figure 9.
CREATE TYPE "Fraction" AS(
                             numerator	SMALLINT,
                             denominator	SMALLINT
                         );

CREATE TYPE "LA_PartyRoleType" AS ENUM (
    'bank', 'certifiedSurveyor', 'citizen', 'conveyancer', 'employee', 'farmer',
    'moneyProvider', 'notary', 'stateAdminsitration', 'surveyor', 'writer'
    );
CREATE TYPE "LA_PartyType" AS ENUM (
    'baunit', 'group', 'naturalPerson', 'nonNaturalPerson'
    );
CREATE TABLE "LA_Party" (
                            id						VARCHAR NOT NULL,							--  pID.namespace || '-' || pID.localId
                            extPID					"Oid",
                            name					VARCHAR,
                            pID						"Oid" NOT NULL,
                            role					"LA_PartyRoleType"[],
                            type					"LA_PartyType" NOT NULL,
                            beginLifeSpanVersion 	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            endLifeSpanVersion		TIMESTAMP DEFAULT '-infinity'::timestamp,
--	quality					"DQ_Element",								-- Omitted for simplicity
--	source 					"CI_ResponsibleParty",						-- Omitted for simplicity
                            PRIMARY KEY(id, beginLifeSpanVersion)
);
-- INHERITS ("VersionedObject"); -- INHERITS not supported yet

--
-- Party::LA_GroupParty
--
-- An instance of class LA_GroupParty is a group party. Class LA_GroupParty is a subclass of LA_Party, thus
-- allowing instances of class LA_GroupParty to have an association with instances of class LA_RRR (and
-- thereby also to class LA_BAUnit). A group party consists of two or more [2..*] parties, but also of other group
-- parties (that is to say, a group party of group parties). Conversely, a party is a member of zero or more [0..*]
-- group parties, see Figure 9.
CREATE TYPE "LA_GroupPartyType" AS ENUM (
    'association', 'baunitGroup', 'family', 'tribe'
    );
CREATE TABLE "LA_GroupParty" (
                                 id                   	VARCHAR NOT NULL,                    --  pID.namespace || '-' || pID.localId
                                 extPID                	"Oid",
                                 name              		VARCHAR,
                                 pID               		"Oid" NOT NULL,
                                 role              		"LA_PartyRoleType"[],
                                 type              		"LA_PartyType" NOT NULL,
                                 groupType          		"LA_GroupPartyType" NOT NULL,
                                 beginLifeSpanVersion  	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                 endLifeSpanVersion      TIMESTAMP,
-- quality                "DQ_Element",                       -- Omitted for simplicity
-- source                 "CI_ResponsibleParty",                -- Omitted for simplicity
                                 PRIMARY KEY(id, beginLifeSpanVersion)
);
-- INHERITS ("VersionedObject"); -- INHERITS not supported yet

--
-- Party::LA_PartyMember
--
-- An instance of class LA_PartyMember is a party member. Class LA_PartyMember is an optional association
-- class between LA_Party and LA_GroupParty, see Figure 9.
CREATE TABLE "LA_PartyMember" (
                                  parties                 	VARCHAR NOT NULL,           -- rID.namespace || '-' || rID.localId
                                  groups                  	VARCHAR NOT NULL,           -- rID.namespace || '-' || rID.localId
                                  partiesBeginLifeSpanVersion	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                  groupsBeginLifeSpanVersion	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                  fraction             		"Fraction",
-- quality                   "DQ_Element",                 -- Omitted for simplicity
-- source                    "CI_ResponsibleParty",          -- Omitted for simplicity

                                  beginLifeSpanVersion		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                  endLifeSpanVersion			TIMESTAMP,
                                  PRIMARY KEY(parties, groups, beginLifeSpanVersion),
                                  UNIQUE(parties, partiesBeginLifeSpanVersion, groups, groupsBeginLifeSpanVersion),
                                  FOREIGN KEY (parties, partiesBeginLifeSpanVersion) REFERENCES "LA_Party"(id, beginLifeSpanVersion),
                                  FOREIGN KEY (groups, groupsBeginLifeSpanVersion) REFERENCES "LA_GroupParty"(id, beginLifeSpanVersion)
);
-- INHERITS ("VersionedObject"); -- INHERITS not supported yet

--***************************************************************************
-- Administrative Package													*
--***************************************************************************

--
-- Administrative::LA_RRR
-- LA_RRR is an abstract class with three specialization classes:
-- 1) LA_Right, with rights as instances. Rights are primarily in the domain of private or customary law.
--    Ownership rights are generally based on (national) legislation, and code lists in the LADM are in
--    support of this, see Annex J.
-- 2) LA_Restriction, with restrictions as instances. Restrictions usually "run with the land", meaning that
--    they remain valid, even when the right to the land is transferred after the right was created (and
--    registered). A mortgage, an instance of class LA_Mortgage, is a special restriction of the ownership
--    right. It concerns the conveyance of a property by a debtor to a creditor, as a security for a financial
--    loan, with the condition that the property is returned, when the loan is paid off.
-- 3) LA_Responsibility, with responsibilities as instances.
--CREATE TABLE "LA_RRR" (
--	id			VARCHAR NOT NULL,				-- rID.namespace || '-' || rID.localId
--	description	VARCHAR,
--	rID			"Oid" NOT NULL, 				-- PRIMARY KEY containing column of type 'user_defined_type' not yet supported in YigabyteDB
--	share		"Fraction",
--	shareCheck	BOOLEAN DEFAULT TRUE,
--	timeSpec	INTERVAL[],
--	beginLifeSpanVersion 	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--	endLifeSpanVersion		TIMESTAMP DEFAULT '-infinity'::timestamp,
----	quality					"DQ_Element", 					-- Omitted for simplicity
----	source 					"CI_ResponsibleParty",			-- Omitted for simplicity
--	CONSTRAINT r_pk PRIMARY KEY(id, beginLifeSpanVersion, endLifeSpanVersion),
--	CONSTRAINT share_chk CHECK ((share).numerator > 0 AND (share).denominator > 0 AND (share).numerator <= (share).denominator),
--	FOREIGN KEY (id, beginLifeSpanVersion) REFERENCES "LA_Party"(id, beginLifeSpanVersion),
--	FOREIGN KEY (id, beginLifeSpanVersion) REFERENCES "LA_GroupParty"(id, beginLifeSpanVersion)
--);

CREATE TYPE "LA_RightType" AS ENUM (
    'agriActivity', 'belowTheDepth', 'boatHarbour', 'commonwealthAcquisition', 'covenant', 'easement',
    'excludedArea', 'forest', 'freeholding', 'grazing', 'housingLand', 'industrialState', 'landsLease'
        'leae', 'mainRoad', 'marinePark', 'mineTenure', 'nationalPark', 'occupation', 'ownership',
    'portAuthority', 'profitPrendre', 'railway', 'reserve', 'stateForest', 'stateLand', 'timberReserve',
    'transferredProperty', 'waterResource', 'waterRights'
    );
CREATE TABLE "LA_Right" (
                            id							VARCHAR NOT NULL,				-- rID.namespace || '-' || rID.localId
                            description					VARCHAR,
                            rID							"Oid" NOT NULL, 				-- PRIMARY KEY containing column of type 'user_defined_type' not yet supported in YugabyteDB
                            share						"Fraction",
                            shareCheck					BOOLEAN DEFAULT TRUE,
                            timeSpec					INTERVAL[],
                            type						"LA_RightType" NOT NULL DEFAULT 'ownership',
--	quality						"DQ_Element", 					-- Omitted for simplicity
--	source 						"CI_ResponsibleParty",			-- Omitted for simplicity
                            beginLifeSpanVersion 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            endLifeSpanVersion			TIMESTAMP DEFAULT '-infinity'::timestamp,
                            party						VARCHAR NOT NULL,				--  pID.namespace || '-' || pID.localId
                            partyBeginLifeSpanVersion 	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            unit						VARCHAR NOT NULL,				--  uID.namespace || '-' || uID.localId
                            unitBeginLifeSpanVersion 	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            PRIMARY KEY(id, beginLifeSpanVersion),
                            CONSTRAINT right_share_chk CHECK ((share).numerator > 0 AND (share).denominator > 0 AND (share).numerator <= (share).denominator),
                            FOREIGN KEY (party, partyBeginLifeSpanVersion) REFERENCES "LA_Party"(id, beginLifeSpanVersion)
);
-- INHERITS ("LA_RRR"); -- INHERITS not supported yet

CREATE TYPE "LA_RestrictionType" AS ENUM (
    'adminPublicServitude', 'monument', 'monumentPartly', 'noBuilding', 'servitude', 'servitudePartly'
    );
CREATE TABLE "LA_Restriction" (
                                  id						VARCHAR NOT NULL,				-- rID.namespace || '-' || rID.localId
                                  description				VARCHAR,
                                  rID						"Oid" NOT NULL, 				-- PRIMARY KEY containing column of type 'user_defined_type' not yet supported in YigabyteDB
                                  share					"Fraction",
                                  shareCheck				BOOLEAN DEFAULT TRUE,
                                  timeSpec				INTERVAL[],
                                  partyRequired			BOOLEAN DEFAULT TRUE,
                                  type					"LA_RestrictionType" NOT NULL,
--	quality					"DQ_Element", 					-- Omitted for simplicity
--	source 					"CI_ResponsibleParty",			-- Omitted for simplicity
                                  beginLifeSpanVersion 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                  endLifeSpanVersion			TIMESTAMP DEFAULT '-infinity'::timestamp,
                                  party						VARCHAR NOT NULL,				--  pID.namespace || '-' || pID.localId
                                  partyBeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                                  unit						VARCHAR NOT NULL,				--  uID.namespace || '-' || uID.localId
                                  unitBeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                                  PRIMARY KEY(id, beginLifeSpanVersion),
                                  CONSTRAINT restriction_share_chk CHECK ((share).numerator > 0 AND (share).denominator > 0 AND (share).numerator <= (share).denominator),
                                  FOREIGN KEY (party, partyBeginLifeSpanVersion) REFERENCES "LA_Party"(id, beginLifeSpanVersion)
);
-- INHERITS ("LA_RRR"); -- INHERITS not supported yet

CREATE TYPE "LA_ResponsibilityType" AS ENUM (
    'monumentMaintenance', 'waterwayMaintenance'
    );
CREATE TABLE "LA_Responsibility" (
                                     id						VARCHAR NOT NULL,				-- rID.namespace || '-' || rID.localId
                                     description				VARCHAR,
                                     rID						"Oid" NOT NULL, 				-- PRIMARY KEY containing column of type 'user_defined_type' not yet supported in YigabyteDB
                                     share					"Fraction",
                                     shareCheck				BOOLEAN DEFAULT TRUE,
                                     timeSpec				INTERVAL[],
                                     type					"LA_ResponsibilityType" NOT NULL,
--	quality					"DQ_Element", 					-- Omitted for simplicity
--	source 					"CI_ResponsibleParty",			-- Omitted for simplicity
                                     beginLifeSpanVersion 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     endLifeSpanVersion			TIMESTAMP DEFAULT '-infinity'::timestamp,
                                     party						VARCHAR NOT NULL,				--  pID.namespace || '-' || pID.localId
                                     partyBeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                                     unit						VARCHAR NOT NULL,				--  uID.namespace || '-' || uID.localId
                                     unitBeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                                     PRIMARY KEY(id, beginLifeSpanVersion),
                                     CONSTRAINT responsibility_share_chk CHECK ((share).numerator > 0 AND (share).denominator > 0 AND (share).numerator <= (share).denominator),
                                     FOREIGN KEY (party, partyBeginLifeSpanVersion) REFERENCES "LA_Party"(id, beginLifeSpanVersion)
);
-- INHERITS ("LA_RRR"); -- INHERITS not supported yet

-- An instance of class LA_Mortgage is a mortgage. LA_Mortgage is a subclass of LA_Restriction. LA_Mortgage
-- is associated to class LA_Right (the right that is the basis for the mortgage). A mortgage can be associated to
-- zero or more [0..*] rights (i.e. a mortgage can be associated specifically to the right which is the object of the
-- mortgage). In all cases, the mortgage is associated, through LA_Restriction and LA_RRR, to the basic
-- administrative unit which is affected by the mortgage; see Figure 10.
CREATE TYPE "LA_MortgageType" AS ENUM (
    'levelPayment', 'linear', 'microcredit'
    );
CREATE TABLE "LA_Mortgage" (
                               id						VARCHAR NOT NULL,				-- rID.namespace || '-' || rID.localId
                               amount					REAL,
                               interestRate			REAL,
                               ranking					SMALLINT,
                               type					"LA_MortgageType",
                               description				VARCHAR,
                               rID						"Oid" NOT NULL, 				-- PRIMARY KEY containing column of type 'user_defined_type' not yet supported in YigabyteDB
                               share					"Fraction",
                               shareCheck				BOOLEAN DEFAULT TRUE,
                               timeSpec				INTERVAL[],
                               partyRequired			BOOLEAN DEFAULT TRUE,
                               restrictionType			"LA_RestrictionType" NOT NULL,
--	quality					"DQ_Element", 					-- Omitted for simplicity
--	source 					"CI_ResponsibleParty",			-- Omitted for simplicity
                               beginLifeSpanVersion 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               endLifeSpanVersion			TIMESTAMP DEFAULT '-infinity'::timestamp,
                               party						VARCHAR NOT NULL,				--  pID.namespace || '-' || pID.localId
                               partyBeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                               unit						VARCHAR NOT NULL,				--  uID.namespace || '-' || uID.localId
                               unitBeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                               PRIMARY KEY(id, beginLifeSpanVersion),
                               CONSTRAINT mortgage_share_chk CHECK ((share).numerator > 0 AND (share).denominator > 0 AND (share).numerator <= (share).denominator),
                               FOREIGN KEY (party, partyBeginLifeSpanVersion) REFERENCES "LA_Party"(id, beginLifeSpanVersion)
);
-- INHERITS ("LA_Restriction"); -- INHERITS not supported yet

CREATE TABLE "mortgageRight" (
                                 mortgage						VARCHAR NOT NULL,			-- LA_Mortgage.id
                                 right_							VARCHAR NOT NULL,			-- LA_Right.id
                                 mortgageBeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                                 right_BeginLifeSpanVersion 		TIMESTAMP NOT NULL,
                                 PRIMARY KEY(mortgage, mortgageBeginLifeSpanVersion, right_, right_BeginLifeSpanVersion),
                                 FOREIGN KEY (mortgage, mortgageBeginLifeSpanVersion) REFERENCES "LA_Mortgage"(id, beginLifeSpanVersion),
                                 FOREIGN KEY (right_, right_BeginLifeSpanVersion) REFERENCES "LA_Right"(id, beginLifeSpanVersion)
);


--
-- Administrative::LA_BAUnit
--
-- An instance of class LA_BAUnit is a basic administrative unit. LA_BAUnit is associated to class LA_Party (a
-- party may be a basic administrative unit). A basic administrative unit is associated to zero or more [0..*] spatial
-- units. A basic administrative unit shall be associated to one or more [1..*] instances of right, restriction or
-- responsibility (i.e. a basic administrative unit cannot exist if there is not at least one right, restriction or
-- responsibility associated to it). A basic administrative unit can be spatially related, through a required
-- relationship, to zero or more [0..*] other basic administrative units (i.e. create an explicit spatial relationship
-- between two basic administrative units when the geometry is missing or inaccurate to provide reliable implicit
-- results). Basic administrative units do not need to be related explicitly. However, if an explicit required
-- relationship is specified, a basic administrative unit shall be associated to one or more [1..*] other basic
-- administrative units. A basic administrative unit can be associated to zero or more [0..*] administrative sources
-- (i.e. the basic administrative unit is usually described as the object affected by the right, restriction or
-- responsibility in the administrative source). A basic administrative unit can be associated to zero or more [0..*]
-- spatial sources (i.e. the extent – part of – of a basic administrative unit can be described on a spatial source).
-- See Figure 10.

CREATE TYPE "LA_BAUnitType" AS ENUM (
    'basicPropertyUnit', 'leasedUnit', 'rightOfUseUnit'
    );
CREATE TABLE "LA_BAUnit" (
                             id			VARCHAR NOT NULL,				-- uID.namespace || '-' || uID.localID
                             type		"LA_BAUnitType" NOT NULL,
                             uID			"Oid" NOT NULL, 				-- PRIMARY KEY containing column of type 'user_defined_type' not yet supported in YigabyteDB
--	quality					"DQ_Element", 					-- Omitted for simplicity
--	source 					"CI_ResponsibleParty",			-- Omitted for simplicity
                             beginLifeSpanVersion 	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             endLifeSpanVersion		TIMESTAMP DEFAULT '-infinity'::timestamp,
                             PRIMARY KEY (id, beginLifeSpanVersion)
);
ALTER TABLE "LA_Right"
    ADD CONSTRAINT right_baunit_fk
        FOREIGN KEY (unit, unitBeginLifeSpanVersion) REFERENCES "LA_BAUnit"(id, beginLifeSpanVersion);
ALTER TABLE "LA_Restriction"
    ADD CONSTRAINT right_baunit_fk
        FOREIGN KEY (unit, unitBeginLifeSpanVersion) REFERENCES "LA_BAUnit"(id, beginLifeSpanVersion);
ALTER TABLE "LA_Responsibility"
    ADD CONSTRAINT right_baunit_fk
        FOREIGN KEY (unit, unitBeginLifeSpanVersion) REFERENCES "LA_BAUnit"(id, beginLifeSpanVersion);
ALTER TABLE "LA_Mortgage"
    ADD CONSTRAINT mortgage_baunit_fk
        FOREIGN KEY (unit, unitBeginLifeSpanVersion) REFERENCES "LA_BAUnit"(id, beginLifeSpanVersion);


CREATE TABLE "LA_RequiredRelationshipBAUnit" (
                                                 id							VARCHAR NOT NULL,
                                                 relationship				VARCHAR NOT NULL,
--	quality						"DQ_Element",								-- Omitted for simplicity
--	source 						"CI_ResponsibleParty",						-- Omitted for simplicity
                                                 beginLifeSpanVersion 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                                 endLifeSpanVersion			TIMESTAMP DEFAULT '-infinity'::timestamp,
                                                 unit1						VARCHAR NOT NULL,			-- LA_BAUnit.id
                                                 unit1BeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                                                 unit2						VARCHAR NOT NULL,			-- LA_BAUnit.id
                                                 unit2BeginLifeSpanVersion	TIMESTAMP NOT NULL,
                                                 PRIMARY KEY(id, beginLifeSpanVersion),
                                                 UNIQUE(unit1, unit1BeginLifeSpanVersion, unit2, unit2BeginLifeSpanVersion),
                                                 FOREIGN KEY (unit1, unit1BeginLifeSpanVersion) REFERENCES "LA_BAUnit" (id, beginLifeSpanVersion),
                                                 FOREIGN KEY (unit2, unit2BeginLifeSpanVersion) REFERENCES "LA_BAUnit" (id, beginLifeSpanVersion)
);
-- INHERITS ("VersionedObject"); -- INHERITS not supported yet


CREATE TABLE "baunitAsParty" (
                                 unit						VARCHAR NOT NULL,		-- LA_BAUnit.id
                                 party						VARCHAR NOT NULL,		-- LA_Party.id
                                 unitBeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                                 partyBeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                                 PRIMARY KEY (unit, unitBeginLifeSpanVersion, party, partyBeginLifeSpanVersion),
                                 FOREIGN KEY (unit, unitBeginLifeSpanVersion) REFERENCES "LA_BAUnit"(id, beginLifeSpanVersion),
                                 FOREIGN KEY (party, partyBeginLifeSpanVersion) REFERENCES "LA_Party"(id, beginLifeSpanVersion)
);

--***************************************************************************
-- Spatial Unit Package														*
--***************************************************************************
DROP TYPE IF EXISTS "ExtPhysicalBuildingUnit";
DROP TYPE IF EXISTS "ExtAddress";
DROP TYPE IF EXISTS "LA_AreaValue";
DROP TYPE IF EXISTS "LA_AreaType";
DROP TYPE IF EXISTS "LA_DimensionType";
DROP TYPE IF EXISTS "LA_SurfaceRelationType";
DROP TYPE IF EXISTS "LA_RegisterType";
DROP TYPE IF EXISTS "LA_StructureType";
DROP TYPE IF EXISTS "LA_LevelContentType";

CREATE TYPE "ExtAddress" AS (
                                addressAreaName		VARCHAR,
                                addressCoordinate	geometry(POINT),
                                addressId			Oid,
                                buildingName		VARCHAR,
                                buildingNumber		VARCHAR,
                                city				VARCHAR,
                                country				VARCHAR,
                                postalCode			VARCHAR,
                                postBox				VARCHAR,
                                state				VARCHAR,
                                streetName			VARCHAR
                            );
--
-- Code lists for Spatial Unit Package
--
CREATE TYPE "LA_AreaType" AS ENUM (
    'officialArea', 'nonOfficialArea', 'calculatedArea', 'surveyedArea'
    );
CREATE TYPE "LA_AreaValue" AS (
                                  areaSize		REAL,
                                  type			"LA_AreaType"
                              );
CREATE TYPE "LA_DimensionType" AS ENUM (
    '0D', '1D', '2D', '3D', 'liminal'
    );
CREATE TYPE "LA_SurfaceRelationType" AS ENUM (
    'mixed', 'below', 'above', 'onSurface'
    );
CREATE TYPE "LA_RegisterType" AS ENUM (
    'urban', 'rural', 'mining',	'publicSpace', 'forest', 'all'
    );
CREATE TYPE "LA_StructureType" AS ENUM (
    'point', 'polygon', 'text',	'topological', 'unstructuredLine', 'sketch'
    );
CREATE TYPE "LA_LevelContentType" AS ENUM (
    'building', 'customary', 'mixed', 'network',	'primaryRight',	'responsibility', 'restriction', 'informal'
    );



--
-- Spatial Unit::LA_SpatialUnitGroup

-- Any number of spatial units (4.1.23), considered as an entity.
-- An instance of class LA_SpatialUnitGroup is a spatial unit group. A spatial unit group is made of one or more
-- [1..*] parts/elements (which can be spatial units, or spatial unit groups, or a combination of spatial units and
-- spatial unit groups). A spatial unit group is part of zero or one [0..1] larger spatial unit group, which again can
-- even be part of zero or one [0..1] larger spatial unit group, and so on. See Figure 11.
--
CREATE TABLE "LA_SpatialUnitGroup" (
                                       id						VARCHAR NOT NULL,				-- sugID.namespace || '-' || sugID.localId
                                       hierarchyLevel			INTEGER NOT NULL DEFAULT 0,
                                       label					VARCHAR,
                                       name					VARCHAR,
                                       referencePoint			geometry(POINT),
                                       sugID					"Oid" NOT NULL, 				-- PRIMARY KEY containing column of type 'user_defined_type' not yet supported in YugabyteDB
                                       beginLifeSpanVersion 	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                       endLifeSpanVersion		TIMESTAMP DEFAULT '-infinity'::timestamp,
--	quality					"DQ_Element", 					-- Omitted for simplicity
--	source 					"CI_ResponsibleParty",			-- Omitted for simplicity
                                       PRIMARY KEY (id, beginLifeSpanVersion)
);
-- INHERITS ("VersionedObject"); -- INHERITS not supported yet

--
-- Spatial Unit::LA_Level - Set of spatial units (4.1.23), with a geometric, and/or topological, and/or thematic coherence
--
CREATE TABLE "LA_Level" (
                            id						VARCHAR NOT NULL,				--  lID.namespace || '-' || lID.localId
                            lID						"Oid" NOT NULL,					-- PRIMARY KEY containing column of type 'user_defined_type' not yet supported in YugabyteDB
                            name					VARCHAR,
                            registerType			"LA_RegisterType",
                            structure				"LA_StructureType",
                            type					"LA_LevelContentType",
                            beginLifeSpanVersion 	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            endLifeSpanVersion		TIMESTAMP DEFAULT '-infinity'::timestamp,
--	quality					"DQ_Element",					-- Omitted for simplicity
--	source 					"CI_ResponsibleParty",			-- Omitted for simplicity
                            PRIMARY KEY (id, beginLifeSpanVersion)
);
-- INHERITS ("VersionedObject"); -- INHERITS not supported yet

--
-- Spatial Unit::LA_SpatialUnit - single area (or multiple areas) of land (4.1.9) and/or water, or a single volume (or multiple volumes) of space
--
CREATE TABLE "LA_SpatialUnit" (
                                  id							VARCHAR NOT NULL,				--  suID.namespace || '-' || suID.localId
                                  extAddressID				"ExtAddress",
                                  area						"LA_AreaValue",
                                  dimension					"LA_DimensionType" DEFAULT '2D',
                                  label						VARCHAR,
                                  referencePoint				geometry(POINT),
                                  suID						"Oid" NOT NULL, 				-- PRIMARY KEY containing column of type 'user_defined_type' not yet supported in YugabyteDB
                                  surfaceRelation				"LA_SurfaceRelationType",
                                  level						VARCHAR NOT NULL,				--  suID.namespace || '-' || suID.localId
                                  levelBeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                                  beginLifeSpanVersion 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                  endLifeSpanVersion			TIMESTAMP DEFAULT '-infinity'::timestamp,
--	quality					"DQ_Element",					-- Omitted for simplicity
--	source 					"CI_ResponsibleParty",			-- Omitted for simplicity
                                  PRIMARY KEY (id, beginLifeSpanVersion),
                                  FOREIGN KEY (level, levelBeginLifeSpanVersion) REFERENCES "LA_Level"(id, beginLifeSpanVersion)
);
-- INHERITS ("VersionedObject"); -- INHERITS not supported yet

-- Required relationships are explicit spatial relationships between spatial units, and instances of class
-- LA_RequiredRelationshipSpatialUnit. Sometimes there is a need for these explicit spatial relationships, when
-- the geometry of the spatial units is not accurate enough to give reliable results, when applying geospatial
-- overlaying techniques (e.g. a building, in reality inside a parcel, is reported to fall outside the parcel; the same
-- applies to the geometry of a right, e.g. an easement). Required relationships override implicit relationships,
-- established through geospatial overlaying techniques.
DROP TYPE IF EXISTS "ISO19125_Type";
CREATE TYPE "ISO19125_Type" AS ENUM (
    'ST_Equals', 'ST_Disjoint', 'ST_Intersects', 'ST_Touches', 'ST_Crosses', 'ST_Within', 'ST_Contains', 'ST_Overlaps'
    );
CREATE TABLE "LA_RequiredRelationshipSpatialUnit" (
                                                      id								VARCHAR NOT NULL,
--	quality							"DQ_Element",								-- Omitted for simplicity
--	source 							"CI_ResponsibleParty",						-- Omitted for simplicity
                                                      su1								VARCHAR NOT NULL,			-- LA_SpatialUnit.id
                                                      su1BeginLifeSpanVersion 		TIMESTAMP NOT NULL,
                                                      su2								VARCHAR NOT NULL,			-- LA_SpatialUnit.id
                                                      su2BeginLifeSpanVersion 		TIMESTAMP NOT NULL,
                                                      relationship					"ISO19125_Type",
                                                      beginLifeSpanVersion 			TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                                      endLifeSpanVersion				TIMESTAMP DEFAULT '-infinity'::timestamp,
                                                      PRIMARY KEY (id, beginLifeSpanVersion),
                                                      UNIQUE(su1, su1BeginLifeSpanVersion, su2, su2BeginLifeSpanVersion),
                                                      FOREIGN KEY (su1, su1BeginLifeSpanVersion) REFERENCES "LA_SpatialUnit"(id, beginLifeSpanVersion),
                                                      FOREIGN KEY (su2, su2BeginLifeSpanVersion) REFERENCES "LA_SpatialUnit"(id, beginLifeSpanVersion)
);
-- INHERITS ("VersionedObject"); -- INHERITS not supported yet

--
-- Spatial Unit::LA_LegalSpaceBuildingUnit
--
-- LA_LegalSpaceBuildingUnit is a subclass of class LA_SpatialUnit.
-- An instance of class LA_LegalSpaceBuildingUnit is a building unit, i.e, - component of building
--
DROP TYPE IF EXISTS "LA_BuildingUnitType";
CREATE TYPE "LA_BuildingUnitType" AS ENUM (
    'individual', 'shared'
    );
CREATE TYPE "ExtPhysicalBuildingUnit" AS (
    extAddressID	"ExtAddress"
                                         );
CREATE TABLE "LA_LegalSpaceBuildingUnit" (
                                             id							VARCHAR NOT NULL,							--  lsbuid.namespace || '-' || lsbsuid.localId
                                             extAddressID				"ExtAddress",
                                             area						"LA_AreaValue",
                                             dimension					"LA_DimensionType" DEFAULT '2D',
                                             label						VARCHAR,
                                             referencePoint				geometry(POINT),
                                             lsbuID						"Oid" NOT NULL,
                                             surfaceRelation				"LA_SurfaceRelationType",
--	quality						"DQ_Element",								-- Omitted for simplicity
--	source 						"CI_ResponsibleParty",						-- Omitted for simplicity
                                             extPhysicalBuildingUnitID	"ExtPhysicalBuildingUnit",
                                             type 						"LA_BuildingUnitType",
                                             beginLifeSpanVersion 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                             endLifeSpanVersion			TIMESTAMP DEFAULT '-infinity'::timestamp,
                                             PRIMARY KEY (id, beginLifeSpanVersion)
);
-- INHERITS ("VersionedObject"); -- INHERITS not supported yet

CREATE TABLE "suSuGroup" (
                             part						VARCHAR NOT NULL,		-- suID.namespace || '-' || suID.localId
                             whole						VARCHAR NOT NULL,		-- sugID.namespace || '-' || sugID.localId
                             partBeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                             wholeBeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                             PRIMARY KEY (part, partBeginLifeSpanVersion, whole, wholeBeginLifeSpanVersion),
                             FOREIGN KEY (part, partBeginLifeSpanVersion) REFERENCES "LA_SpatialUnit"(id, beginLifeSpanVersion),
                             FOREIGN KEY (whole, wholeBeginLifeSpanVersion) REFERENCES "LA_BAUnit"(id, beginLifeSpanVersion)
);

CREATE TABLE "suBaunit" (
                            su							VARCHAR NOT NULL,		-- sugID.namespace || '-' || sugID.localId
                            baunit						VARCHAR NOT NULL,		-- LA_BAUnit.id
                            suBeginLifeSpanVersion 		TIMESTAMP NOT NULL,
                            baunitBeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                            beginLifeSpanVersion 	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            endLifeSpanVersion		TIMESTAMP DEFAULT '-infinity'::timestamp,
                            PRIMARY KEY (su, suBeginLifeSpanVersion, baunit, baunitBeginLifeSpanVersion, beginLifeSpanVersion),
                            FOREIGN KEY (su, suBeginLifeSpanVersion) REFERENCES "LA_SpatialUnit"(id, beginLifeSpanVersion),
                            FOREIGN KEY (baunit, baunitBeginLifeSpanVersion) REFERENCES "LA_BAUnit"(id, beginLifeSpanVersion)
);

CREATE TABLE "suHierarchy" (
                               child						VARCHAR NOT NULL,		-- childID.namespace || '-' || childID.localId
                               parent						VARCHAR NOT NULL,		-- parentID.namespace || '-' || parentID.localId
                               childBeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                               parentBeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                               PRIMARY KEY (child, childBeginLifeSpanVersion),
                               FOREIGN KEY (child, childBeginLifeSpanVersion) REFERENCES "LA_SpatialUnit"(id, beginLifeSpanVersion),
                               FOREIGN KEY (parent, parentBeginLifeSpanVersion) REFERENCES "LA_SpatialUnit"(id, beginLifeSpanVersion)
);

CREATE TABLE "suGroupHierarchy" (
                                    element						VARCHAR NOT NULL,		-- childID.namespace || '-' || childID.localId
                                    set							VARCHAR NOT NULL,		-- parentID.namespace || '-' || parentID.localId
                                    elementBeginLifeSpanVersion TIMESTAMP NOT NULL,
                                    setBeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                                    PRIMARY KEY (element, elementBeginLifeSpanVersion),
                                    FOREIGN KEY (element, elementBeginLifeSpanVersion) REFERENCES "LA_SpatialUnitGroup"(id, beginLifeSpanVersion),
                                    FOREIGN KEY (set, setBeginLifeSpanVersion) REFERENCES "LA_SpatialUnitGroup"(id, beginLifeSpanVersion)
);



--
--***************************************************************************
-- Surveying and Representation Subpackage														*
--***************************************************************************

--
-- Surveying and Representation::LA_BoundaryFaceString
--
-- boundary face string: boundary (4.1.3) forming part of the outside of a spatial unit (4.1.23).
-- NOTE Boundary face strings are used to represent the boundaries of spatial units by means of line strings in 2D.
-- This 2D representation is a 2D boundary in a 2D land administration (4.1.10) system. In a 3D land administration system
-- it represents a series of vertical boundary faces (4.1.4) where an unbounded volume is assumed, surrounded by
-- boundary faces which intersect the Earth’s surface (such as traditionally depicted in the cadastral map).
-- An instance of class LA_BoundaryFaceString is a boundary face string. LA_BoundaryFaceString is associated
-- to class LA_Point and class LA_SpatialSource to document the origin of the geometry. In the case of a
-- location by text, a boundary face string would not be defined by points. However, in all other cases, a
-- boundary face string shall be defined by two or more [2..*] points (i.e. as a minimum a boundary starts and
-- ends at a point, i.e. a straight line).

CREATE TABLE "LA_BoundaryFaceString"(
                                        id						VARCHAR NOT NULL,						--  bfsid.namespace || '-' || bfsid.localId
                                        bfsID					"Oid" NOT NULL,							-- PRIMARY KEY containing column of type 'user_defined_type' not yet supported in YugabyteDB
                                        geometry				geometry(MULTILINESTRING),
                                        locationByText			VARCHAR,								-- Ignored - we assume 2D polygon-based spatial units
--	quality					"DQ_Element",							-- Omitted for simplicity
--	source 					"CI_ResponsibleParty"					-- Omitted for simplicity
                                        beginLifeSpanVersion 	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                        endLifeSpanVersion		TIMESTAMP DEFAULT '-infinity'::timestamp,
                                        PRIMARY KEY (id, beginLifeSpanVersion)
);
-- INHERITS ("VersionedObject"); -- INHERITS not supported yet

CREATE TABLE "minus" (
                         bfs						VARCHAR NOT NULL,		-- bfsid.namespace || '-' || bfsid.localId
                         su						VARCHAR NOT NULL,		-- suID.namespace || '-' || suID.localId
                         bfsBeginLifeSpanVersion TIMESTAMP NOT NULL,
                         suBeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                         beginLifeSpanVersion 	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         endLifeSpanVersion		TIMESTAMP DEFAULT '-infinity'::timestamp,
                         PRIMARY KEY (bfs, bfsBeginLifeSpanVersion, su, suBeginLifeSpanVersion, beginLifeSpanVersion),
                         FOREIGN KEY (bfs, bfsBeginLifeSpanVersion) REFERENCES "LA_BoundaryFaceString"(id, beginLifeSpanVersion),
                         FOREIGN KEY (su, suBeginLifeSpanVersion) REFERENCES "LA_SpatialUnit"(id, beginLifeSpanVersion)
);

CREATE TABLE "plus" (
                        bfs						VARCHAR NOT NULL,		-- bfsid.namespace || '-' || bfsid.localId
                        su						VARCHAR NOT NULL,		-- suID.namespace || '-' || suID.localId
                        bfsBeginLifeSpanVersion TIMESTAMP NOT NULL,
                        suBeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                        beginLifeSpanVersion 	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        endLifeSpanVersion		TIMESTAMP DEFAULT '-infinity'::timestamp,
                        PRIMARY KEY (bfs, bfsBeginLifeSpanVersion, su, suBeginLifeSpanVersion, beginLifeSpanVersion),
                        FOREIGN KEY (bfs, bfsBeginLifeSpanVersion) REFERENCES "LA_BoundaryFaceString"(id, beginLifeSpanVersion),
                        FOREIGN KEY (su, suBeginLifeSpanVersion) REFERENCES "LA_SpatialUnit"(id, beginLifeSpanVersion)
);



--
-- Surveying and Representation::LA_Point
--
-- An instance of class LA_Point is a point. A point may be associated to zero or one [0..1] spatial units (i.e. the
-- point may be used as the reference point to describe the position of a spatial unit). A point may be associated
-- to zero or more [0..*] boundary faces (i.e. a point may be used to define a vertex of the side of a 3D parcel). A
-- point may be associated to zero or more [0..*] boundary face strings (i.e. a point can be used to define the
-- start, end or vertex of a boundary). A point should be associated to zero or more [0..*] spatial sources. See Figure 12.
--
CREATE TYPE "LA_InterpolationType" AS ENUM (
    'end', 'isolated', 'mid', 'midArc', 'start'
    );
CREATE TYPE "LA_MonumentationType" AS ENUM (
    'beacon', 'cornerstone', 'marker', 'notMarked'
    );
CREATE TYPE "LA_PointType" AS ENUM (
    'control', 'noSource', 'source'
    );
CREATE TYPE "LI_Lineage" AS (
    statement	VARCHAR --,
--	scope		MD_Scope,
--	additionalDocumentation	CI_Citation
                            );
CREATE TYPE "CC_Formula" AS (
    formula			VARCHAR --,
--	formulaCitation	CI_Citation
                            );
CREATE TYPE "CC_OperationMethod" AS (
                                        formulaReference	"CC_Formula",
                                        sourceDimensions	SMALLINT,
                                        targetDimensions	SMALLINT
                                    );
CREATE TYPE "LA_Transformation" AS (
                                       transformation		"CC_OperationMethod",
                                       transformedLocation	geometry(POINT)
                                   );
CREATE TABLE "LA_Point" (
                            id							VARCHAR NOT NULL,						--  pid.namespace || '-' || pid.localId
                            interpolationRole			"LA_InterpolationType" NOT NULL DEFAULT 'mid',
                            monumentation				"LA_MonumentationType",
                            originalLocation			geometry(POINT) NOT NULL,
                            pID							"Oid" NOT NULL,
                            pointType					"LA_PointType" NOT NULL DEFAULT 'control',
                            productionMethod			"LI_Lineage",
                            transAndResult				"LA_Transformation",
--	quality						"DQ_Element",						-- Omitted for simplicity
--	source 						"CI_ResponsibleParty"				-- Omitted for simplicity
                            beginLifeSpanVersion 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            endLifeSpanVersion			TIMESTAMP DEFAULT '-infinity'::timestamp,
                            PRIMARY KEY (id, beginLifeSpanVersion)
);
-- INHERITS ("VersionedObject"); -- INHERITS not supported yet

CREATE TABLE "pointBfs" (
                            point						VARCHAR NOT NULL,		-- pID.namespace || '-' || pID.localId
                            bfs							VARCHAR NOT NULL,		-- bfsid.namespace || '-' || bfsid.localId
                            pointBeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                            bfsBeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                            PRIMARY KEY (point, pointBeginLifeSpanVersion, bfs, bfsBeginLifeSpanVersion),
                            FOREIGN KEY (point, pointBeginLifeSpanVersion) REFERENCES "LA_Point"(id, beginLifeSpanVersion),
                            FOREIGN KEY (bfs, bfsBeginLifeSpanVersion) REFERENCES "LA_BoundaryFaceString"(id, beginLifeSpanVersion)
);


--***************************************************************************
-- LADM 2D Polygon based profile											*
--***************************************************************************

CREATE TABLE "PolygonLevel" (
                                id						VARCHAR NOT NULL,				--  lID.namespace || '-' || lID.localId
                                lID						"Oid" NOT NULL,					-- PRIMARY KEY containing column of type 'user_defined_type' not yet supported in YugabyteDB
                                name					VARCHAR,
                                registerType			"LA_RegisterType",
                                structure				"LA_StructureType" DEFAULT 'polygon',
                                type					"LA_LevelContentType",
                                beginLifeSpanVersion 	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                endLifeSpanVersion		TIMESTAMP DEFAULT '-infinity'::timestamp,
--	quality					"DQ_Element",					-- Omitted for simplicity
--	source 					"CI_ResponsibleParty",			-- Omitted for simplicity
                                PRIMARY KEY (id, beginLifeSpanVersion)
);
-- INHERITS ("VersionedObject"); -- INHERITS not supported yet

--
-- Polygon_Profile::Polygon_SpatialUnitGroup
--
CREATE TABLE "PolygonSpatialUnitGroup" (
                                           id						VARCHAR NOT NULL,							--  sugID.namespace || '-' || sugID.localId
                                           sugID					"Oid" NOT NULL,								-- PRIMARY KEY containing column of type 'user_defined_type' not yet supported in YugabyteDB
                                           hierarchyLevel			INTEGER NOT NULL DEFAULT 1,
                                           label					VARCHAR,
                                           name					VARCHAR,
--	element					"Oid",
                                           beginLifeSpanVersion 	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                           endLifeSpanVersion		TIMESTAMP DEFAULT '-infinity'::timestamp,
--	quality					"DQ_Element",								-- Omitted for simplicity
--	source 					"CI_ResponsibleParty",						-- Omitted for simplicity
                                           PRIMARY KEY (id, beginLifeSpanVersion)
);
-- INHERITS ("VersionedObject"); -- INHERITS not supported yet

--
-- Polygon_Profile::Polygon_SpatialUnit
--
-- a ‘polygon based’ spatial unit (polygon spatial unit) is used when every spatial unit is recorded as a
-- separate entity. There is no topological connection between neighbouring spatial units (and no
-- boundaries shared), and so any constraint, enforcing a complete coverage, shall be applied by the
-- originating and receiving software. In the 2D representation there is exactly one link to a closed boundary
-- face string for every ring of the polygon (or set of boundary face strings, that together form a closed ring).
CREATE TABLE "PolygonSpatialUnit" (
                                      id							VARCHAR NOT NULL,							--  suID.namespace || '-' || suID.localId
                                      extAddressID				"ExtAddress",
                                      area						"LA_AreaValue",
                                      dimension					"LA_DimensionType" DEFAULT '2D',
                                      label						VARCHAR,
                                      referencePoint				geometry(POINT),
                                      suID						"Oid",
                                      surfaceRelation				"LA_SurfaceRelationType",
--	quality						"DQ_Element",								-- Omitted for simplicity
--	source 						"CI_ResponsibleParty",						-- Omitted for simplicity
                                      level						VARCHAR NOT NULL,							--  lID.namespace || '-' || lID.localId
                                      levelBeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                                      beginLifeSpanVersion 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                      endLifeSpanVersion			TIMESTAMP DEFAULT '-infinity'::timestamp,
                                      PRIMARY KEY (id, beginLifeSpanVersion),
                                      FOREIGN KEY (level, levelBeginLifeSpanVersion) REFERENCES "PolygonLevel"(id, beginLifeSpanVersion)
);
-- INHERITS ("VersionedObject"); -- INHERITS not supported yet
ALTER TABLE "suBaunit" ADD CONSTRAINT suBaunit_polygonSpatialUnit_fk FOREIGN KEY (su, suBeginLifeSpanVersion) REFERENCES "PolygonSpatialUnit"(id, beginLifeSpanVersion);


--
-- Polygon_Profile::Polygon_Boundary
--
CREATE TABLE "PolygonBoundary" (
                                   id							VARCHAR NOT NULL,				--  bfsID.namespace || '-' || bfsID.localId
                                   bfsID						"Oid",
                                   geometry					geometry(MULTILINESTRING) NOT NULL,
                                   beginLifeSpanVersion 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                   endLifeSpanVersion			TIMESTAMP DEFAULT '-infinity'::timestamp,
--	quality						"DQ_Element",					-- Omitted for simplicity
--	source 						"CI_ResponsibleParty",			-- Omitted for simplicity
                                   psu							VARCHAR NOT NULL,				--  suID.namespace || '-' || suID.localId
                                   psuBeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                                   psug						VARCHAR NOT NULL,				--  sugID.namespace || '-' || sugID.localId
                                   psugBeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                                   PRIMARY KEY (id, beginLifeSpanVersion),
                                   FOREIGN KEY (psu, psuBeginLifeSpanVersion) REFERENCES "PolygonSpatialUnit"(id, beginLifeSpanVersion),
                                   FOREIGN KEY (psug, psugBeginLifeSpanVersion) REFERENCES "PolygonSpatialUnitGroup"(id, beginLifeSpanVersion)
);
-- INHERITS ("VersionedObject"); -- INHERITS not supported yet

CREATE TABLE "pointPb" (
                           point						VARCHAR NOT NULL,		-- pID.namespace || '-' || pID.localId
                           pb							VARCHAR NOT NULL,		-- bfsid.namespace || '-' || bfsid.localId
                           pointBeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                           pbBeginLifeSpanVersion 	TIMESTAMP NOT NULL,
                           PRIMARY KEY (point, pointBeginLifeSpanVersion, pb, pbBeginLifeSpanVersion),
                           FOREIGN KEY (point, pointBeginLifeSpanVersion) REFERENCES "LA_Point"(id, beginLifeSpanVersion),
                           FOREIGN KEY (pb, pbBeginLifeSpanVersion) REFERENCES "PolygonBoundary"(id, beginLifeSpanVersion)
);


--COMMIT;

