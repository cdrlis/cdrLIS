package repositories

import "github.com/jinzhu/gorm"

type LadmRepository struct {
	DB *gorm.DB
}

func (ctx LadmRepository) Create(value interface{}) error {
	ctx.DB.Create(value)
	return nil // TODO error handling
}

func (ctx LadmRepository) Read(out interface{}) error {
	ctx.DB.First(out)
	return nil // TODO error handling
}

func (ctx LadmRepository) ReadAll(out interface{}) error {
	ctx.DB.Find(out)
	return nil // TODO error handling
}

func (ctx LadmRepository) Update(value interface{}) error {
	ctx.DB.Save(value)
	return nil // TODO error handling
}

func (ctx LadmRepository) Delete(value interface{}) error {
	ctx.DB.Delete(value)
	return nil // TODO error handling
}
