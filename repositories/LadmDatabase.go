package repositories

import "github.com/jinzhu/gorm"

type LadmDatabase struct {
	DB *gorm.DB
}

func (ctx LadmDatabase) Create(value interface{}) error {
	ctx.DB.Create(value)
	return nil // TODO error handling
}

func (ctx LadmDatabase) Read(out interface{}) error {
	ctx.DB.First(out)
	return nil // TODO error handling
}

func (ctx LadmDatabase) ReadAll(out interface{}) error {
	ctx.DB.Find(out)
	return nil // TODO error handling
}

func (ctx LadmDatabase) Update(value interface{}) error {
	ctx.DB.Save(value)
	return nil // TODO error handling
}

func (ctx LadmDatabase) Delete(value interface{}) error {
	ctx.DB.Delete(value)
	return nil // TODO error handling
}
