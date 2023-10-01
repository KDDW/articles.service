package db

import "gorm.io/gorm"

type DB struct {
	Db *gorm.DB
}

func NewDB() *DB {
	return &DB{
		Db: newGORMDB(),
	}
}
