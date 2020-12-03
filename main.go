package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/swaggo/cli"
	"github.com/yrisob/database_migrations/config"
	"github.com/yrisob/database_migrations/database"
	"github.com/yrisob/database_migrations/utils"
)

func getPathWithSources(input string) string {
	trimingInput := strings.TrimSpace(input)
	if trimingInput == "" {
		config, err := config.GetConfig()
		if err == nil && strings.TrimSpace(config.Sources) != "" {
			return strings.TrimSpace(config.Sources)
		}
		return "./db/migrations/"
	}
	return trimingInput
}

func getDataSource(source string) (string, error) {
	if source == "" {
		config, err := config.GetConfig()
		if err != nil {
			return "", err
		}
		return config.GetConnectionString(), nil
	}
	return source, nil
}

func upgradeDatabase(source, connectionString string) error {
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

func main() {
	app := &cli.App{
		Name:  "database_migrations",
		Usage: "application create (create) template file for sql migration or execute migrations (exec)",
		Commands: []cli.Command{
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "create new template file for migration with sql format",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "source, s",
						Usage: "path for database migration files",
					},
					cli.StringFlag{
						Name:  "name, n",
						Usage: "Name of migration",
					},
				},
				Action: func(c *cli.Context) error {
					source := getPathWithSources(c.String("source"))
					name := strings.TrimSpace(c.String("name"))
					if name == "" {
						return errors.New("you should get name for your migration")
					}

					filePath, err := utils.CreateMigrationFile(source, name)
					if err != nil {
						return err
					}
					fmt.Println("file was created: ", filePath)
					return nil
				},
			},
			{
				Name:    "exec",
				Aliases: []string{"exc"},
				Usage:   "execute migrations from source into database",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "source, s",
						Usage: "path for database migration files",
					},
					cli.StringFlag{
						Name:  "datasource, d",
						Usage: "database connection string, like: postgres://user_name:password@host:port/database?sslmode=disable",
					},
				},
				Action: func(c *cli.Context) error {
					datasource, err := getDataSource(strings.TrimSpace(c.String("datasource")))
					if err != nil {
						return err
					}
					source := getPathWithSources(c.String("source"))
					return upgradeDatabase(source, datasource)
				},
			},
			{
				Name:    "show",
				Aliases: []string{"s"},
				Usage:   "show version of database",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "datasource, d",
						Usage: "database connection string, like: postgres://user_name:password@host:port/database?sslmode=disable",
					},
				},
				Action: func(c *cli.Context) error {
					datasource, err := getDataSource(strings.TrimSpace(c.String("datasource")))
					if err != nil {
						return err
					}
					version, err := database.GetMigrationVersion(datasource)
					if err != nil {
						return err
					}
					fmt.Println("Database version is:", version)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil && !(err.Error() == "file does not exist") {
		panic(err)
	}
}
