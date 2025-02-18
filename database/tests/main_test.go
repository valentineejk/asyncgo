package dbq

import (
	"asyncgo/config"
	"asyncgo/database"
	dbq "asyncgo/database/sqlc"
	"context"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"os"
	"path/filepath"
	"testing"
)



var testQueries *dbq.Queries

func TestMain(m *testing.M) {
	if err := godotenv.Load("../../.env"); err != nil {
		fmt.Println("Warning: .env file not loaded properly, using system environment variables")
	}

	//if os.Getenv("ENV") == "TEST" {
	//	err := os.Setenv("ENV", string(config.ENV_TEST))
	//	if err != nil {
	//		fmt.Println("❌ Test Env error:", err)
	//		os.Exit(1)
	//	}
	//}
	env := os.Getenv("ENV")
	if env != "TEST" {
		fmt.Printf("❌ Tests can only run in TEST environment (current ENV=%s)\n", env)
		os.Exit(1)
	}

	conf, err := config.New()
	if err != nil {
		fmt.Println("❌ Config error:", err)
		os.Exit(1)
	}

	fmt.Printf(">> db: %s\n", conf.DatabaseUrl())

	db, err := database.NewPostgres(conf)
	if err != nil {
		fmt.Printf("❌ Database connection error: %v\n", err)
		os.Exit(1) // Stop execution
	}

	testQueries = dbq.New(db)

	// 🟢 Run database migrations before tests
	err = runMigrations(conf)
	if err != nil {
		fmt.Println("❌ Migration error:", err)
		os.Exit(1)
	}

	// Run the tests
	code := m.Run()

	// Clean up after tests
	if err := db.Close(); err != nil {
		fmt.Printf("❌ Error closing database: %v\n", err)
	}
	os.Exit(code) // Ensure this always runs
	//return &TestEnv{
	//	Config: conf,
	//	Db:     dbq.New(db),
	//}
}

func runMigrations(conf *config.Config) error {

	// Get the absolute path of the migrations directory
	absPath, err := filepath.Abs("../migrations")
	if err != nil {
		fmt.Printf("❌ Failed to resolve absolute path: %v\n", err)
		os.Exit(1)
	}

	migrationPath := fmt.Sprintf("file:///%s", absPath)
	fmt.Println("🔍 Using migration path:", migrationPath)

	// Run database migrations
	mig, err := migrate.New(
		migrationPath, // Path to your migration files
		conf.DatabaseUrl())
	if err != nil {
		fmt.Println("❌ Migration error:", err)
		os.Exit(1)
	}

	err = mig.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		fmt.Println("❌ Migration error:", err)
		os.Exit(1)
	}

	fmt.Println("✅ Migrations applied successfully")
	return nil
}

func TearDownTestDb() error {

	err := testQueries.TruncateTables(context.Background())
	//require.NoError(t, err)
	if err != nil {
		return err
	}
	fmt.Println("✅ Database Teardown Completed")
	return nil
}
