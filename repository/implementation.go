package repository

import (
	"gorm.io/gorm"
)

func (rep *repository) Model(value interface{}) *gorm.DB {
	return rep.db.Model(value)
}

func (rep *repository) Select(query interface{}, args ...interface{}) *gorm.DB {
	return rep.db.Select(query, args...)
}

func (rep *repository) Find(out interface{}, where ...interface{}) *gorm.DB {
	return rep.db.Find(out, where...)
}

func (rep *repository) Exec(sql string, values ...interface{}) *gorm.DB {
	return rep.db.Exec(sql, values...)
}

func (rep *repository) First(out interface{}, where ...interface{}) *gorm.DB {
	return rep.db.First(out, where...)
}

func (rep *repository) Raw(sql string, values ...interface{}) *gorm.DB {
	return rep.db.Raw(sql, values...)
}

func (rep *repository) Create(value interface{}) *gorm.DB {
	return rep.db.Create(value)
}

func (rep *repository) Save(value interface{}) *gorm.DB {
	return rep.db.Save(value)
}

func (rep *repository) Update(col string, value interface{}) *gorm.DB {
	return rep.db.Update(col, value)
}

func (rep *repository) Delete(value interface{}) *gorm.DB {
	return rep.db.Delete(value)
}

func (rep *repository) Where(query interface{}, args ...interface{}) *gorm.DB {
	return rep.db.Where(query, args...)
}

func (rep *repository) AutoMigrate(value interface{}) error {
	return rep.db.AutoMigrate(value)
}

func (rep *repository) Table(name string, args ...interface{}) *gorm.DB {
	return rep.db.Table(name, args...)
}
