package sqlstore

import (
	"context"
	"fmt"
	"strings"

	"simple-wallet/config"
	"simple-wallet/internal/infrastructure/sqlstore/migration"
	"simple-wallet/internal/infrastructure/sqlstore/seeders"
	"simple-wallet/pkg/db"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Migrator struct {
	db       *gorm.DB
	migrator *gormigrate.Gormigrate
}

func openDBWrapper(dbConfig db.Config) (*gorm.DB, error) {
	masterDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&allowCleartextPasswords=1",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)
	sqlDB, err := gorm.Open(mysql.Open(masterDsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}
	return sqlDB, nil
}

func NewMigrator() (*Migrator, error) {
	config := *config.All()
	db, err := openDBWrapper(config.Database.Main)
	if err != nil {
		return nil, err
	}

	migrator := gormigrate.New(db, gormigrate.DefaultOptions, GetMigration())

	return &Migrator{
		db:       db,
		migrator: migrator,
	}, nil
}

func GetMigration() []*gormigrate.Migration {
	return append(migration.GetMigrations(), seeders.GetSeeders()...)
}

// Migrate executes all migrations exists
func (m *Migrator) Migrate() error {

	if err := m.migrator.Migrate(); err != nil {
		return err
	}
	return nil
}

var migrationIDs []string
var migrationLeft int
var totalMigration = len(GetMigration())

var remainingMigration int

// Run migration
func (m *Migrator) MigrateAll(ctx context.Context) (err error) {

	if remainingMigration, err = m.countMigrationLeft(); err != nil {
		return err
	}

	if remainingMigration == 0 {
		fmt.Println("\nNo new migration.")
		return nil
	}

	for i := totalMigration - migrationLeft; i < totalMigration; i++ {
		migrationID := GetMigration()[i].ID

		if err := m.migrator.MigrateTo(migrationID); err != nil {
			return fmt.Errorf("%s when run %s", err.Error(), migrationID)
		}

		fmt.Println("migrated:", migrationID)
	}

	fmt.Println("\nMigrate run successfully.")

	return nil
}

// Rollback migration(s) that start from the last till N step backward.
func (m *Migrator) Rollback(step int) error {
	for _, id := range *m.getMigrationIDs(step) {
		if err := m.migrator.RollbackLast(); err != nil {
			return fmt.Errorf("%s %s", err.Error(), id)
		}

		fmt.Printf("Reverted: %s\n", id)
	}

	fmt.Println("\nRollback successfully.")

	return nil
}

// Rollback all database migrations.
func (m *Migrator) Reset() error {
	step, err := m.countTotalMigrationInDb()

	if err != nil {
		return err
	}

	return m.Rollback(step)
}

// Get migration id(s) that start from the last till N step backward.
func (m *Migrator) getMigrationIDs(step int) *[]string {
	if len(migrationIDs) > 0 {
		return &migrationIDs
	}

	// Get all migration from database.
	migrationString, err := m.getStringMigration()

	if err != nil {
		return nil
	}

	// Filter migration from backward.
	for i := len(GetMigration()) - 1; i >= 0; i-- {
		if step == 0 {
			break
		}

		migrationID := GetMigration()[i].ID

		// We can only rollback migration if they are already exist in database.
		// So we need to check if the ID exists in the database.
		if strings.Contains(*migrationString, migrationID) {
			migrationIDs = append(migrationIDs, migrationID)
			step--
		}
	}

	return &migrationIDs
}

// Get all migrations in the database and join it (Separated by space).
func (m *Migrator) getStringMigration() (*string, error) {
	if err := m.initSchema(); err != nil {
		return nil, err
	}

	rows, err := m.db.Raw("select id from migrations").Rows()

	if err != nil {
		return nil, err
	}

	var id string
	var migrationIDs []string

	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}

		migrationIDs = append(migrationIDs, id)
	}

	migrationString := strings.Join(migrationIDs, " ")

	return &migrationString, nil
}

// Count how many migrations left in definition before we run it.
func (m *Migrator) countMigrationLeft() (int, error) {
	if migrationLeft != 0 {
		return migrationLeft, nil
	}

	migratedNumber, err := m.countTotalMigrationInDb()

	if err != nil {
		return 0, err
	}
	migrationLeft = len(GetMigration()) - migratedNumber

	return migrationLeft, nil
}

// Count total migration in database.
func (m *Migrator) countTotalMigrationInDb() (int, error) {
	migrationString, err := m.getStringMigration()
	if *migrationString == "" {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	migratedNumber := len(strings.Split(*migrationString, " "))
	return migratedNumber, nil
}

func (m *Migrator) initSchema() error {
	if m.db.Migrator().HasTable("migrations") {
		return nil
	}

	sql := fmt.Sprintf("CREATE TABLE %s (%s VARCHAR(%d) PRIMARY KEY)", gormigrate.DefaultOptions.TableName, gormigrate.DefaultOptions.IDColumnName, gormigrate.DefaultOptions.IDColumnSize)

	return m.db.Exec(sql).Error
}

// Reset and re-run all migrations
func (m *Migrator) Refresh() error {
	if err := m.Reset(); err != nil {
		return err
	}

	fmt.Println("") // Add new line

	if err := m.migrator.Migrate(); err != nil {
		return err
	}

	return nil
}

type migrationType int

const (
	Migration migrationType = iota
	Seeder
)

func (mt *migrationType) String() string {
	types := map[migrationType]string{
		Migration: "migration",
		Seeder:    "seeders",
	}

	return types[*mt]
}
