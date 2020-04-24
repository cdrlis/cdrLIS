package dblogic

// CRUDer CRUD
type CRUDer interface {
	Create(value interface{}) error
	Read(out interface{}, where ...interface{}) error
	ReadAll(out interface{}, where ...interface{}) error
	Update(value interface{}, where ...interface{}) error
	Delete(value interface{}, where ...interface{}) error
}
