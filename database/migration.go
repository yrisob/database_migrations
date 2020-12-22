package database

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type (
	Migrations []Migration
	Migration  struct {
		Version   string `gorm:"primary_key;not null;"`
		UpdatedAt time.Time
		Faild     bool `gorm:"type:boolean;not null;default:false"`
	}
)

// IsSuccessfullyExecuted проверяем наличие версии в базе, если нет или она была с ошибками, проверка выдаст отсутствие записи
func IsSuccessfullyExecuted(version string, connectionString string) (bool, error) {
	db, err := GetDB(connectionString)
	if err != nil {
		return false, err
	}
	defer db.Close()
	migration := &Migration{}
	tableExists := db.HasTable(migration)
	if !tableExists {
		db.AutoMigrate(migration)
	}
	if err := db.First(migration, version).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, err
	}
	return !migration.Faild, nil
}

func ExecuteScript(scriptText string) MigrationOperation {
	return func(tx *gorm.DB) error {
		return tx.Exec(scriptText).Error
	}
}

func CreateOrUpdateMigrationVersion(version string, faild bool) MigrationOperation {
	return func(tx *gorm.DB) error {
		migration := &Migration{Version: version, Faild: faild}
		return tx.Save(migration).Error
	}
}

func GetMigrationVersion(connectionString string) (*Migrations, error) {
	db, err := GetDB(connectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	migrations := &Migrations{}
	err = db.Find(migrations).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return nil, err
	} else if err != nil && gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("isn't any migrations")
	}
	return migrations, nil
}
