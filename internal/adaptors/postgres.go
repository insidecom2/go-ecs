package adaptors

import (
	"context"
	"go-ecs/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Postgres struct {
	db *gorm.DB
}

type PostgresInt interface {
	UpsertBatch(ctx context.Context, data []models.Symbols) error
}

func NewPostgres(db *gorm.DB) *Postgres {
	return &Postgres{
		db: db,
	}
}

func (p *Postgres) UpsertBatch(ctx context.Context, data []models.Symbols) error {
	return p.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "symbol"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"name",
				"circulating_supply",
				"total_supply",
				"max_supply",
				"last_price",
			}),
		}).
		CreateInBatches(&data, 500).Error
}
