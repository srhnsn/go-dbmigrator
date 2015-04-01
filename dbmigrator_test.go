package dbmigrator

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/srhnsn/go-dbmigrator/test-assets"
)

const testDefaultDatabaseName = "dbmigrator_test"
const testDefaultDatabaseUser = "root"

var testDatabaseName string
var testDSN string
var testDSNWithDatabase string

func TestMain(m *testing.M) {
	testDatabaseName = os.Getenv("MYSQL_DATABASE")
	testDatabasePassword := os.Getenv("MYSQL_PASSWORD")
	testDatabaseUser := os.Getenv("MYSQL_USER")

	if testDatabaseName == "" {
		testDatabaseName = testDefaultDatabaseName
	}

	if testDatabaseUser == "" {
		testDatabaseUser = testDefaultDatabaseUser
	}

	baseDSN := "%s:%s@/%s?collation=utf8_unicode_ci&loc=Local&parseTime=true"

	testDSN = fmt.Sprintf(baseDSN, testDatabaseUser, testDatabasePassword, "")
	testDSNWithDatabase = fmt.Sprintf(baseDSN, testDatabaseUser, testDatabasePassword, testDatabaseName)

	os.Exit(m.Run())
}

func TestGetCurrentMigrationVersion(t *testing.T) {
	dropTestDatabase()
	createNewTestDatabase()
	migrator := getTestMigrator()

	db, err := migrator.getDB()

	if err != nil {
		t.Errorf("getDB() failed: %s", err)
		return
	}

	version, err := getCurrentMigrationVersion(db)

	if err != ErrCorruptDatabase {
		t.Errorf("getCurrentMigrationVersion() should have returned ErrCorruptDatabase, got %s", err)
		return
	}

	empty := isDatabaseEmpty(db)

	if !empty {
		t.Error("isDatabaseEmpty() should have returned true")
		return
	}

	createSettingTable(db, "")

	version, err = getCurrentMigrationVersion(db)

	if err != ErrCorruptDatabase {
		t.Errorf("getCurrentMigrationVersion() should have returned ErrCorruptDatabase, got %s", err)
		return
	}

	empty = isDatabaseEmpty(db)

	if empty {
		t.Error("isDatabaseEmpty() should have returned false")
		return
	}

	if len(version) != 0 {
		t.Errorf("version should be empty, is %s", version)
		return
	}
}

func TestCompleteMigration(t *testing.T) {
	dropTestDatabase()
	createNewTestDatabase()
	migrator := getTestMigrator()

	err := migrator.Migrate()

	if err != nil {
		t.Errorf("Migrate() failed: %s", err)
		return
	}
}

func TestIncrementalMigration(t *testing.T) {
	dropTestDatabase()
	createNewTestDatabase()
	migrator := getTestAssetMigrator()

	db, err := migrator.getDB()

	if err != nil {
		t.Errorf("getDB() failed: %s", err)
		return
	}

	allMigrations := migrator.migrations

	if len(allMigrations) != 6 {
		t.Errorf("Expected 6 test migrations, got %d", len(allMigrations))
		return
	}

	migrator.migrations = allMigrations[0:1]
	err = migrator.Migrate()

	if err != nil {
		t.Errorf("Migrate() failed: %s", err)
		return
	}

	version, err := getCurrentMigrationVersion(db)

	if err != nil {
		t.Errorf("getCurrentMigrationVersion() failed: %s", err)
		return
	}

	if version != "2015-01-01 120000" {
		t.Errorf("getCurrentMigrationVersion() should have returned 2015-01-01 120000, returned %s", version)
		return
	}

	migrator.migrations = allMigrations[1:]
	err = migrator.Migrate()

	if err != nil {
		t.Errorf("Migrate() failed: %s", err)
		return
	}

	// This should return the latest version.
	version, err = getCurrentMigrationVersion(db)

	if err != nil {
		t.Errorf("getCurrentMigrationVersion() failed: %s", err)
		return
	}

	if version != "2015-01-01 130200" {
		t.Errorf("getCurrentMigrationVersion() should have returned 2015-01-01 130200, returned %s", version)
		return
	}
}

func TestEmptyMigrator(t *testing.T) {
	_, err := getEmptyTestMigrator()

	if err != ErrNoMigrationFiles {
		t.Errorf("getEmptyTestMigrator() should have returned ErrNoMigrationFiles, returned %s", err)
		return
	}
}

func createNewTestDatabase() {
	db, err := sql.Open("mysql", testDSN)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS `" + testDatabaseName + "`")

	if err != nil {
		log.Fatal(err)
	}
}

func dropTestDatabase() {
	db, err := sql.Open("mysql", testDSN)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("DROP DATABASE IF EXISTS `" + testDatabaseName + "`")

	if err != nil {
		log.Fatal(err)
	}
}

func getTestAssetMigrator() Migrator {
	migrator, err := NewMigrator(MigrationConfig{
		Asset:          assets.Asset,
		AssetDir:       assets.AssetDir,
		AssetInfo:      assets.AssetInfo,
		DataSourceName: testDSNWithDatabase,
		DriverName:     "mysql",
		MigrationsPath: "test-migrations",
	})

	if err != nil {
		log.Fatal(err)
	}

	return migrator
}

func getTestMigrator() Migrator {
	migrator, err := NewMigrator(MigrationConfig{
		DataSourceName: testDSNWithDatabase,
		DriverName:     "mysql",
		MigrationsPath: "test-migrations",
	})

	if err != nil {
		log.Fatal(err)
	}

	return migrator
}

func getEmptyTestMigrator() (Migrator, error) {
	return NewMigrator(MigrationConfig{
		DataSourceName: testDSNWithDatabase,
		DriverName:     "mysql",
		MigrationsPath: "test-assets", // No valid migration files there.
	})
}
