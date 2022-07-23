package database

import (
	"context"
	"log"
	"stilyng94/fiber-crm/ent"
	"stilyng94/fiber-crm/ent/migrate"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"ariga.io/atlas/sql/sqltool"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
)

var Db *ent.Client

func InitializeDatabase() {
	client, err := ent.Open("postgres", viper.GetString("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	Db = client
}

func Migrate() {
	//docker://postgres

	ctx := context.Background()
	// Create a local migration directory able to understand golang-migrate migration files for replay.
	dir, err := sqltool.NewGolangMigrateDir("migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}
	// Write migration diff.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.Postgres),        // Ent dialect to use
		schema.WithFormatter(sqltool.GolangMigrateFormatter),
	}
	err = migrate.NamedDiff(ctx, viper.GetString("DATABASE_URL"), "my_migration", opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
