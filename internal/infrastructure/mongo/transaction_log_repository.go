package mongo

import (
	"context"
	"time"

	"github.com/ghitufnine/my-go/internal/domain/entity"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionLogRepository struct {
	collection *mongo.Collection
}

func NewTransactionLogRepository(
	db *mongo.Database,
) *TransactionLogRepository {

	return &TransactionLogRepository{
		collection: db.Collection("transaction_logs"),
	}
}

func (r *TransactionLogRepository) Insert(
	ctx context.Context,
	topic string,
	payload string,
) error {

	log := entity.TransactionLog{
		ID:        uuid.New().String(),
		Topic:     topic,
		Payload:   payload,
		CreatedAt: time.Now(),
	}

	_, err := r.collection.InsertOne(ctx, log)
	return err
}
