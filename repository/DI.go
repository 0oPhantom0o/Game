package repository

import (
	"context"
	"game/domain"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateUser(phone string) (string, error)
	DeleteAnswer(id string) error
	ExpireOtpTime(key string) bool
	ExpireWrongedAnswerTime(key string) bool
	FindUserIdByPhone(phone string) (string, error)
	FindUserByID(primitiveId primitive.ObjectID) (int, error)
	FindStoredOtp(id string) (string, error)
	FindStoredAnswer(id string) (string, error)
	ChangePoint(id primitive.ObjectID, point int) error
	InsertOtp(phone, code string) error
	InsertAnswer(id, answer string) error
	OtpLimit(phone string) (int64, string)
	OTPAnswerLimit(phone string) (int64, string)
	ShowUsers(page, limit int64) ([]domain.TopPlayers, error)
	UpdateNickName(id primitive.ObjectID, nickname string) error
}

type ConRepository struct {
	redisdb *redis.Client
	mongodb *mongo.Collection
	ctx     context.Context
}

func NewMongoRepository(redisdb *redis.Client, mongodb *mongo.Collection, ctx context.Context) *ConRepository {
	return &ConRepository{
		redisdb: redisdb, mongodb: mongodb, ctx: ctx}
}
