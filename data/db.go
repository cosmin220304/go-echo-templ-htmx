package data

import (
	"os"

	"github.com/cosmin220304/go-echo-templ-htmx/data/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	var db, err = gorm.Open(postgres.Open(os.Getenv("LOCAL_DB")), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Todo{})

	return db
}
