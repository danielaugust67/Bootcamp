package main

import (
	"github.com/danielaugust67/hr-api/internal/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// Connect ke DB (schema sudah dimigrasi)
	dsn := "host=localhost user=postgres password=admin dbname=hr_db_dev port=5433 sslmode=disable search_path=hr"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Init generator
	g := gen.NewGenerator(gen.Config{
		OutPath:           "../../internal/domain/query",
		OutFile:           "gen.go",
		ModelPkgPath:      "github.com/danielaugust67/hr-api/internal/domain/models",
		WithUnitTest:      false,
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:     true, // Nullable fields jadi pointer
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	// Assign DB untuk introspeksi schema
	g.UseDB(db)

	// Exclude tabel dengan strategy (return "" untuk ignore)
	g.WithTableNameStrategy(func(tableName string) (targetTableName string) {
		if tableName == "schema_migrations" {
			return "" // Ignore tabel ini
		}
		return tableName // Generate yang lain
	})

	// Apply existing models
	g.ApplyBasic(
		models.Region{},
		models.Country{},
		models.Location{},
		models.Department{},
		models.Job{},
		models.Dependent{},
		models.EmployeeDocument{},
		models.Employee{},
		models.User{},
		models.UserRole{},
		models.Role{},
	)

	// Execute ke file
	g.Execute()

}
