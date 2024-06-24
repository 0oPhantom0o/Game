package logic

import "go.mongodb.org/mongo-driver/bson/primitive"

func ConvertStringToPrimivite(Id string) (primitive.ObjectID, error) {
	ID, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return ID, nil
}
