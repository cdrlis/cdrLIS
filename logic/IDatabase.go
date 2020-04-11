package logic

// Interface IDatabase CRUD
type IDatabase interface {
	Create(value interface{}) error
	Read(out interface{}) error
	ReadAll(out interface{}) error
	Update(value interface{}) error
	Delete(value interface{}) error
}
