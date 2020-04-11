package logic

// IRepository CRUD
type IRepository interface {
	Create(value interface{}) error
	Read(out interface{}) error
	ReadAll(out interface{}) error
	Update(value interface{}) error
	Delete(value interface{}) error
}
