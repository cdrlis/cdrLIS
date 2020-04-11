package logic

// IRepository CRUD
type IRepository interface {
	Create(value interface{}) error
	Get(out interface{}, where ...interface{}) error
	GetAll(out interface{}, where ...interface{}) error
	Update(value interface{}, where ...interface{}) error
	Delete(value interface{}, where ...interface{}) error
}
