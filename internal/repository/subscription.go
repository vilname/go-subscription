package repository

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"subscription-back/config/storage"
	"subscription-back/internal/model"
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

func (repository *SubscriptionRepository) CreateSubscription(email string) (string, error) {
	query := "insert into subscription(email, hash, created) values ($1, $2, NOW())"

	uuidHash, _ := uuid.NewUUID()

	_, err := repository.db.Exec(repository.ctx, query, email, uuidHash)
	if err != nil {
		return uuidHash.String(), err
	}

	return uuidHash.String(), nil
}

func (repository *SubscriptionRepository) GetByEmail(email string) (int, error) {
	var id int

	query := "select id from subscription where email = $1"

	err := repository.db.QueryRow(repository.ctx, query, email).Scan(&id)

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return id, err
	}

	return id, nil
}

func (repository *SubscriptionRepository) GetByHash(hash string) (model.SubscriptionModel, error) {
	var subscriptionModel model.SubscriptionModel

	query := "select id, email from subscription where hash = $1"

	err := repository.db.QueryRow(repository.ctx, query, hash).Scan(&subscriptionModel.Id, &subscriptionModel.Email)

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return subscriptionModel, err
	}

	return subscriptionModel, nil
}

func (repository *SubscriptionRepository) GetSerialNumber(hash string) (string, error) {
	//isExist := false
	var cardUuid *string

	query := "select card_uuid from subscription where hash = $1"

	err := repository.db.QueryRow(repository.ctx, query, hash).Scan(&cardUuid)

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return *cardUuid, err
	}

	if cardUuid == nil {
		return "", nil
	}

	return *cardUuid, nil
}

func (repository *SubscriptionRepository) UpdateCardUuid(hash string, serialNumber string) error {
	query := "update subscription set card_uuid = $1 where hash = $2"

	_, err := repository.db.Exec(repository.ctx, query, serialNumber, hash)

	if err != nil {
		return err
	}

	return nil
}
