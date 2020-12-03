package database

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Migration struct {
	Version   string
	UpdatedAt time.Time
}

func ExecuteScript(scriptText string) MigrationOperation {
	return func(tx *gorm.DB) error {
		return tx.Exec(scriptText).Error
	}
}

func UpdateMigrationVersion(version string) MigrationOperation {
	return func(tx *gorm.DB) error {
		migration := &Migration{}
		tx.AutoMigrate(migration)
		err := tx.First(migration).Error
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			return err
		} else if err != nil && gorm.IsRecordNotFoundError(err) {
			migration.Version = version
			return tx.Save(migration).Error
		} else {
			return tx.Model(migration).Update("version", version).Error
		}

	}
}

func GetMigrationVersion(connectionString string) (string, error) {
	db, err := GetDB(connectionString)
	if err != nil {
		return "", err
	}
	defer db.Close()
	migration := &Migration{}
	err = db.First(migration).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return "", err
	} else if err != nil && gorm.IsRecordNotFoundError(err) {
		return "", errors.New("isn't any migrations")
	}
	return migration.Version, nil
}
