package handler

// CRUDer CRUD
type CRUDer interface {
	Create(value interface{}) (interface{}, error)
	Read(where ...interface{}) (interface{}, error)
	ReadAll(where ...interface{}) (interface{}, error)
	Update(value interface{}, where ...interface{}) (interface{}, error)
	Delete(value interface{}, where ...interface{}) error
}
