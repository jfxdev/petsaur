package migrations

import (
	"github.com/jfxdev/petsaur/src/entities"
	"github.com/jfxdev/petsaur/src/infra"

	"gorm.io/gorm"
)

type migration struct {
	db *gorm.DB
}

func New(db *gorm.DB) infra.Migration {
	return &migration{
		db: db,
	}
}

func (m migration) Run() error {
	return m.db.AutoMigrate(&entities.Pet{})
}
