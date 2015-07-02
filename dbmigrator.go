// Package dbmigrator is yet another library to migrate database schemas.
//
// It saves the current migration version directly in the database. There is also
// support for the go-bindata program.
package dbmigrator

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"sort"
)

const defaultSettingTableCreateStatement = `
CREATE TABLE setting (
    ` + "`key`" + ` VARCHAR(100) NOT NULL,
    value TEXT NOT NULL,
    PRIMARY KEY (` + "`key`" + ` )
)`
const versionKey = "migration_version"

// ErrCorruptDatabase is returned when the setting table does not have the
// migration_version key or another database error occured while retrieving
// the migration version.
var ErrCorruptDatabase = errors.New("failed to retrieve migration version from setting table (database not empty)")

// ErrNoMigrationFiles is returned when there are no files in the directory
// that match the required filename pattern.
var ErrNoMigrationFiles = errors.New("migration directory does not contain any valid migration files")

var filenamePattern = regexp.MustCompile(`^([0-9][0-9][0-9][0-9]-[0-1][0-9]-[0-3][0-9] [0-2][0-9][0-5][0-9][0-5][0-9]) (.*)\.sql$`)

// MigrationConfig contains the configuration for the migrator.
// The use of Asset, AssetDir and AssetInfo functions is optional, but if you
// set a value for either of them, you have to set values for all of them. If
// you leave SettingTableCreateStatement empty, the default create statement
// will be used.
type MigrationConfig struct {
	Asset                       func(name string) ([]byte, error)
	AssetDir                    func(name string) ([]string, error)
	AssetInfo                   func(name string) (os.FileInfo, error)
	DataSourceName              string
	DriverName                  string
	MigrationsPath              string
	Quiet                       bool
	SettingTableCreateStatement string
}

// Migrator contains the migrations to apply to the database.
type Migrator struct {
	config     MigrationConfig
	migrations migrations
}

type migration struct {
	description string
	filename    string
	basename    string
	version     string
}

type migrations []migration

func (a migrations) Len() int           { return len(a) }
func (a migrations) Less(i, j int) bool { return a[i].basename < a[j].basename }
func (a migrations) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (m *migration) isNew(version string) bool {
	return m.version > version
}

func (m *migration) apply(db *sql.DB, readFile func(name string) ([]byte, error)) error {
	query, err := readFile(m.filename)

	if err != nil {
		return err
	}

	_, err = db.Exec(string(query))

	if err != nil {
		return fmt.Errorf("failed to apply %s: %s", m.basename, err)
	}

	err = m.updateVersion(db)

	if err != nil {
		return fmt.Errorf("failed to update version to %s: %s", m.version, err)
	}

	return nil
}

func (m *migration) updateVersion(db *sql.DB) error {
	version, err := json.Marshal(m.version)

	if err != nil {
		return err
	}

	_, err = db.Exec(`
        REPLACE INTO
            setting ( `+"`key`"+` , value )
        VALUES
            ( ? , ? )`, versionKey, version)

	return err
}

// Migrate applies all new migrations to the database. After each successful
// migration it will update the "migration_version" key in the setting table
// to the corresponding migration version. If there is an error while applying
// a migration, the operation is canceled (but not rolled back).
func (m *Migrator) Migrate() error {
	db, err := m.getDB()
	defer db.Close()

	if err != nil {
		return err
	}

	empty := isDatabaseEmpty(db)
	var migrations migrations

	if empty {
		if !m.config.Quiet {
			//fmt.Println("Empty database, creating setting table and applying all migrations.")
		}

		createSettingTable(db, m.config.SettingTableCreateStatement)
		migrations = m.migrations
	} else {
		version, err := getCurrentMigrationVersion(db)

		if err != nil {
			return err
		}

		migrations = getNewMigrations(version, m.migrations)
	}

	var readFile func(name string) ([]byte, error)

	if m.config.Asset == nil {
		readFile = ioutil.ReadFile
	} else {
		readFile = m.config.Asset
	}

	for _, migration := range migrations {
		if !m.config.Quiet {
			//fmt.Printf("Applying %s\n", migration.basename)
		}

		err := migration.apply(db, readFile)

		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Migrator) getDB() (*sql.DB, error) {
	return sql.Open(m.config.DriverName, m.config.DataSourceName)
}

// NewMigrator returns a new migrator. It reads all migrations (files) in the migration
// path that match the "YYYY-MM-DD hhmmss *.sql" filename format and sorts them.
func NewMigrator(config MigrationConfig) (Migrator, error) {
	if (config.Asset != nil || config.AssetDir != nil || config.AssetInfo != nil) &&
		(config.Asset == nil || config.AssetDir == nil || config.AssetInfo == nil) {
		return Migrator{}, errors.New("when using Asset, AssetDir or AssetInfo, neither of them can be nil")
	}

	migrations, err := getAllMigrations(config)

	if err != nil {
		return Migrator{}, err
	}

	if len(migrations) == 0 {
		return Migrator{}, ErrNoMigrationFiles
	}

	migrator := Migrator{
		config:     config,
		migrations: migrations,
	}

	return migrator, nil
}

func createSettingTable(db *sql.DB, stmt string) error {
	if stmt == "" {
		stmt = defaultSettingTableCreateStatement
	}

	_, err := db.Exec(stmt)
	return err
}

func fileInfoAdapter(
	assetDir func(name string) ([]string, error),
	assetInfo func(name string) (os.FileInfo, error)) func(name string) ([]os.FileInfo, error) {

	return func(name string) ([]os.FileInfo, error) {
		entries, err := assetDir(name)

		if err != nil {
			return nil, err
		}

		result := make([]os.FileInfo, len(entries))

		for i, entry := range entries {
			result[i], err = assetInfo(path.Join(name, entry))

			if err != nil {
				return nil, err
			}
		}

		return result, nil
	}
}

func getAllMigrations(config MigrationConfig) (migrations, error) {
	var readDir func(name string) ([]os.FileInfo, error)

	if config.AssetDir == nil {
		readDir = ioutil.ReadDir
	} else {
		readDir = fileInfoAdapter(config.AssetDir, config.AssetInfo)
	}

	entries, err := readDir(config.MigrationsPath)

	if err != nil {
		return nil, err
	}

	migrations := make(migrations, 0)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := path.Base(entry.Name())
		match := filenamePattern.FindStringSubmatch(name)

		if match == nil {
			continue
		}

		migration := migration{
			basename:    name,
			description: match[2],
			filename:    path.Join(config.MigrationsPath, name),
			version:     match[1],
		}

		migrations = append(migrations, migration)
	}

	sort.Sort(migrations)

	return migrations, nil
}

func getCurrentMigrationVersion(db *sql.DB) (string, error) {
	var version string
	var versionRaw []byte

	err := db.QueryRow(`
        SELECT
            value
        FROM
            setting
        WHERE
            `+"`key`"+` = ?
        LIMIT
            1`, versionKey).Scan(&versionRaw)

	switch {
	case err == sql.ErrNoRows:
		fallthrough
	case err != nil:
		return "", ErrCorruptDatabase
	}

	err = json.Unmarshal(versionRaw, &version)

	if err != nil {
		return "", err
	}

	return version, nil
}

func getNewMigrations(version string, migs migrations) migrations {
	result := make(migrations, 0)

	for _, migration := range migs {
		if migration.isNew(version) {
			result = append(result, migration)
		}
	}

	return result
}

func isDatabaseEmpty(db *sql.DB) bool {
	var table string

	err := db.QueryRow(`SHOW TABLES`).Scan(&table)

	switch {
	case err == sql.ErrNoRows:
		return true
	case err != nil:
		panic(err)
	default:
		return false
	}
}
