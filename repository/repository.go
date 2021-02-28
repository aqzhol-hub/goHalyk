package repository

import "gorm.io/gorm"

type Repository interface {
	Model(value interface{}) *gorm.DB
	Select(query interface{}, args ...interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	Exec(sql string, values ...interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	Raw(sql string, values ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Update(col string, value interface{}) *gorm.DB
	Delete(value interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	AutoMigrate(value interface{}) error
	Table(name string, args ...interface{}) *gorm.DB
}

type repository struct {
	db *gorm.DB
}
