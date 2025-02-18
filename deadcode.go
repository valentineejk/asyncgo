package main

//var testQueries *dbq.Queries
//
//func TestMain(m *testing.M) {
//	if err := godotenv.Load("../.env"); err != nil {
//		fmt.Println("Warning: .env file not loaded properly, using system environment variables")
//	}
//
//	if os.Getenv("ENV") == "TEST" {
//		os.Setenv("ENV", string(config.ENV_TEST))
//	}
//
//	// Get the absolute path of the migrations directory
//	absPath, err := filepath.Abs("migrations")
//	if err != nil {
//		fmt.Printf("❌ Failed to resolve absolute path: %v\n", err)
//		os.Exit(1)
//	}
//
//	migrationPath := fmt.Sprintf("file:///%s", absPath)
//	fmt.Println("🔍 Using migration path:", migrationPath)
//
//	conf, err := config.New()
//	if err != nil {
//		fmt.Println("❌ Config error:", err)
//		os.Exit(1)
//	}
//	fmt.Printf(">> db: %s\n", conf.DatabaseUrl())
//
//	db, err := NewPostgres(conf)
//	if err != nil {
//		_ = fmt.Errorf("databse connection error: %w", err)
//	}
//
//
//	//require.NoError(m, err) // Ensure database connection works
//
//	// Run database migrations
//	mig, err := migrate.New(
//		migrationPath, // Path to your migration files
//		conf.DatabaseUrl())
//	if err != nil {
//		fmt.Println("❌ Migration error:", err)
//		os.Exit(1)
//	}
//	//require.NoError(m, err) // Ensure migrations are loaded
//
//	err = mig.Up()
//	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
//		fmt.Println("❌ Migration error:", err)
//		os.Exit(1)
//	}
//
//	testQueries = dbq.New(db)
//
//	// Run the tests
//	code := m.Run()
//
//	// Close DB connection before exiting
//	db.Close()
//	os.Exit(code)
//}
