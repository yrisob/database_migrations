package main

import (
	"testing"

	"github.com/yrisob/database_migrations/database"
)

const connectionString = "postgres://postgres:gfhjkm1986@127.0.0.1:5432/tc_user?sslmode=disable"

const queryInsert = `
	CREATE TABLE "public".table_a(id integer, name varchar(100));
	CREATE TABLE "public".table_b(id integer, name varchar(100));
`

const queryDROP = `
	DROP TABLE IF EXISTS "public".table_a;
	DROP TABLE IF EXISTS "public".table_b;
`

// func TestTransactions(t *testing.T) {
// 	insertMigrations := database.ExecuteScript(queryInsert)
// 	updateVersion := database.UpdateMigrationVersion("1234")
// 	err := database.ExecMigrationOperations(connectionString, insertMigrations, updateVersion)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log("migrations executed success")
// }

func TestGetMigrationVersion(t *testing.T) {
	version, err := database.GetMigrationVersion(connectionString)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(version)
}
