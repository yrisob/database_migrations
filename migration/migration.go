package migration

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/yrisob/database_migrations/database"
	"github.com/yrisob/database_migrations/utils"
)

func UpgradeDatabase(source, connectionString string) error {
	version, _ := database.GetMigrationVersion(connectionString)
	files, err := ioutil.ReadDir(source)
	if err != nil {
		return err
	}

	for _, file := range utils.SortFilesByNameAsc(files) {
		fileVersion := utils.GetVersionByFileName(file.Name())
		if !file.IsDir() && version < fileVersion {
			bytes, err := ioutil.ReadFile(path.Join(source, file.Name()))
			if err != nil {
				return err
			}
			insertMigration := database.ExecuteScript(string(bytes))
			updateVersion := database.UpdateMigrationVersion(fileVersion)
			err = database.ExecMigrationOperations(connectionString, insertMigration, updateVersion)
			if err != nil {
				return err
			}
		}
	}
	version, err = database.GetMigrationVersion(connectionString)
	if err != nil {
		return err
	}

	fmt.Println("New DB version is:", version)
	return nil
}
