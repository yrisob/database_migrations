package main

import (
	"io/ioutil"
	"testing"

	"github.com/yrisob/database_migrations/utils"
)

const source = "./db/migrations"

func TestGetFileName(t *testing.T) {
	filePartName := "inserUser"
	t.Log(utils.CreateFileName(filePartName))
}

func TestSortFiles(t *testing.T) {
	files, err := ioutil.ReadDir(source)
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range utils.SortFilesByNameAsc(files) {
		t.Log(file.Name())
	}
}

func TestCreateMigrationFile(t *testing.T) {
	filePath, err := utils.CreateMigrationFile(source, "add_User")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("file was created: ", filePath)
}
