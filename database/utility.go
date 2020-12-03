package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type MigrationOperation = func(tx *gorm.DB) error

func GetDB(connectionString string) (*gorm.DB, error) {
	if db, err := gorm.Open("postgres", connectionString); err != nil {
		return nil, err
	} else {
		return db.LogMode(true), nil
	}

}

//CheckDB -  check that database configuration right and database is exist
func CheckDB(connectionString string) error {
	if db, err := GetDB(connectionString); err != nil {
		return err
	} else {
		db.Close()
		return nil
	}
}

func ExecMigrationOperations(connectionString string, operations ...MigrationOperation) error {
	db, err := GetDB(connectionString)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Transaction(func(tx *gorm.DB) error {
		for _, operation := range operations {
			if err := operation(tx); err != nil {
				return err
			}
		}
		return nil
	})
}
