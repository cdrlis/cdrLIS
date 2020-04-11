package repositories

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type LadmDatabase struct {
	DB *gorm.DB
}

func (ctx LadmDatabase) Create(value interface{}) error {
	ctx.DB.Create(value)
	return nil // TODO error handling
}

func (ctx LadmDatabase) Get(out interface{}, where ...interface{}) error {
	if where != nil {
		if ctx.DB.First(out, where).RowsAffected == 0 {
			return errors.New("Entity not found")
		}
		return nil
	}
	ctx.DB.First(out)
	return nil // TODO error handling
}

func (ctx LadmDatabase) GetAll(out interface{}, where ...interface{}) error {
	ctx.DB.Find(out)
	return nil // TODO error handling
}

func (ctx LadmDatabase) Update(value interface{}, where ...interface{}) error {
	ctx.DB.Save(value)
	return nil // TODO error handling
}

func (ctx LadmDatabase) Delete(value interface{}, where ...interface{}) error {
	ctx.DB.Delete(value)
	return nil // TODO error handling
}
