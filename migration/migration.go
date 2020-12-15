package migration

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/yrisob/database_migrations/database"
	"github.com/yrisob/database_migrations/utils"
)

func UpgradeDatabase(source, connectionString string) error {

	files, err := ioutil.ReadDir(source)
	if err != nil {
		return err
	}

	//var version string
	executedVersions := []string{}
	for _, file := range utils.SortFilesByNameAsc(files) {
		//version, _ = database.GetMigrationVersion(connectionString)
		fileVersion := utils.GetVersionByFileName(file.Name())
		isVersionSuccessfullyAdd, err := database.IsSuccessfullyExecuted(fileVersion, connectionString)
		if err != nil {
			return err
		}
		if !file.IsDir() && !isVersionSuccessfullyAdd {
			executedVersions = append(executedVersions, fileVersion)
			bytes, err := ioutil.ReadFile(path.Join(source, file.Name()))
			if err != nil {
				return err
			}
			insertMigration := database.ExecuteScript(string(bytes))
			updateVersion := database.CreateOrUpdateMigrationVersion(fileVersion, false)
			err = database.ExecMigrationOperations(connectionString, insertMigration, updateVersion)
			if err != nil {
				if err = database.ExecMigrationOperations(connectionString, database.CreateOrUpdateMigrationVersion(fileVersion, true)); err != nil {
					return err
				}
			}
		}
	}
	// version, err = database.GetMigrationVersion(connectionString)
	// if err != nil {
	// 	return err
	// }

	fmt.Printf("Executed this versions: %v \n", executedVersions)
	return nil
}
