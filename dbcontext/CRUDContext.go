package dbcontext

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type CRUDContext struct {
	DB *gorm.DB
}

func (ctx CRUDContext) Create(value interface{}) error {
	ctx.DB.Create(value)
	return nil // TODO error handling
}

func (ctx CRUDContext) Read(out interface{}, where ...interface{}) error {
	if where != nil {
		if ctx.DB.Set("gorm:auto_preload", true).Where(where[0], where[1:]).First(out).RowsAffected == 0 {
			return errors.New("Entity not found")
		}
		return nil
	}
	ctx.DB.First(out)
	return nil // TODO error handling
}

func (ctx CRUDContext) ReadAll(out interface{}, where ...interface{}) error {
	if where != nil {
		if ctx.DB.Where(where[0], where[1:]).Find(out).RowsAffected == 0 {
			return errors.New("Entity not found")
		}
		return nil
	}
	ctx.DB.Find(out)
	return nil // TODO error handling
}

func (ctx CRUDContext) Update(value interface{}, where ...interface{}) error {
	ctx.DB.Save(value)
	return nil // TODO error handling
}

func (ctx CRUDContext) Delete(value interface{}, where ...interface{}) error {
	ctx.DB.Delete(value)
	return nil // TODO error handling
}
