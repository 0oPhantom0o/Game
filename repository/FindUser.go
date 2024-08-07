package repository

import (
	"fmt"
	"game/constants"
	"game/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *ConRepository) FindUserIdByPhone(phone string) (string, error) {
	var user domain.UserId
	collection := repo.mongodb.Database(constants.Database).Collection(constants.UserCollection)

	filter := bson.D{{"phone", phone}}
	//find _id based on phone
	err := collection.FindOne(repo.ctx, filter).Decode(&user)
	if err != nil {
		return "", err
	}
	return user.ID.Hex(), nil
}
func (repo *ConRepository) FindUserByID(primitiveId primitive.ObjectID) (int, error) {
	var user domain.InternalUser
	collection := repo.mongodb.Database(constants.Database).Collection(constants.UserCollection)

	filter := bson.D{{"_id", primitiveId}}
	//find _id based on phone
	err := collection.FindOne(repo.ctx, filter).Decode(&user)
	if err != nil {
		return 0, err
	}
	return user.NickNameLimit, nil
}

func (repo *ConRepository) FindStoredOtp(id string) (string, error) {

	otp, err := repo.redisdb.Get(repo.ctx, id).Result()
	if err != nil {
		return "", fmt.Errorf("user doesnt exist")
	}
	return otp, nil

}

func (repo *ConRepository) FindStoredAnswer(id string) (string, error) {
	answer, err := repo.redisdb.Get(repo.ctx, id).Result()
	if err != nil {
		return "", fmt.Errorf("answer doesnt exist")
	}
	return answer, nil

}
