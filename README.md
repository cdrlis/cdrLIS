## ***cdrLIS*** -- generic and extensible core of a consistent, distributed and resilient land information system (LIS).

**cdrLIS** is a generic and extensible LADM-compliant core framework implemented in [Go](https://golang.org) programming language.

The [Land Administration Domain Model (LADM)](https://www.iso.org/standard/51206.html) is a well-known international standard for the domain of land administration. It is a conceptual, generic domain model for designing and building  LIS and has been already extended and adapted to a number of particular profiles. The model covers the basic data related components of land administration: (i) party related data; (ii) data on rights, restrictions and responsibilities (RRR) and the basic administrative units where RRR apply to; and (iii) data on spatial units and on surveying and topology/geometry. LADM provides an extensible basis for the development and refinement of LIS, based on a Model Driven Architecture (MDA). 

**cdrLIS** relies on [gogeos](https://github.com/paulsmith/gogeos/tree/master/geos) packages to manage and use geospatial data types and operations on them. The gogeos packages provides bindings to the [GEOS](%28https://trac.osgeo.org/geos%29) library that implements the geometry model and API according to the [OGC Simple Features Specification for SQL.](https://www.ogc.org/standards/sfs) In order to isolate object-oriented paradigm from the relational paradigm, we use the object-relational mapping package [GORM](https://gorm.io).

**cdrLIS** core framework can be extended in two directions: 

 1. To support specific types of spatial unit, i.e. spatial profiles  (polygon based, topological based, or similar); and 
 2. To support implementation of a specific country profile.

The framework is also generic in the sense that it can be used in building of both distributed and centralized LIS, as well as other applications in the land administration domain including data acquisition tools.

**cdrLIS** has been intensively and sucessfuly tested on [YugaByteDB](https://www.yugabyte.com) -- distributed NewSQL DBMS for global, internet-scale applications with low query latency and extreme resilience against failure. However, it can be reused and extended in developing either centralized or distributed LIS provided that underlying DBMS implements OGC Simple Features Specification for SQL.

