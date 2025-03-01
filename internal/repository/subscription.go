package repository

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"subscription-back/config/storage"
)

type SubscriptionRepository struct {
	db  *pgxpool.Pool
	ctx *gin.Context
}

func NewSubscriptionRepository(ctx *gin.Context) *SubscriptionRepository {
	return &SubscriptionRepository{
		db:  storage.GetDB(),
		ctx: ctx,
	}
}

func (repository *SubscriptionRepository) CreateSubscription(email string) error {
	query := "insert into subscription(email, hash, created) values ($1, $2, NOW())"

	uuidHash, _ := uuid.NewUUID()

	_, err := repository.db.Exec(repository.ctx, query, email, uuidHash)
	if err != nil {
		return err
	}

	return nil
}

func (repository *SubscriptionRepository) FindByEmail(email string) (int, error) {
	var id int

	query := "select id from subscription where email = $1"

	err := repository.db.QueryRow(repository.ctx, query, email).Scan(&id)

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return id, err
	}

	return id, nil
}
