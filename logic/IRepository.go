package logic

// IRepository CRUD
type IRepository interface {
	Create(value interface{}) error
	Get(out interface{}) error
	GetAll(out interface{}) error
	Update(value interface{}) error
	Delete(value interface{}) error
}
