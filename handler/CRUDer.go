package handler

// CRUDer SpatialUnitCRUD
type CRUDer interface {
	Create(value interface{}) (interface{}, error)
	Read(where ...interface{}) (interface{}, error)
	ReadAll(where ...interface{}) (interface{}, error)
	Update(value interface{}) (interface{}, error)
	Delete(value interface{}) error
}
