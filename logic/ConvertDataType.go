package logic

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

func convertStringToPrimitive(Id string) (primitive.ObjectID, error) {
	primitiveId, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return primitiveId, nil
}

func convertStringToInteger(stringNumber string) (int64, error) {
	number, err := strconv.Atoi(stringNumber)
	if err != nil {
		return 0, fmt.Errorf("can't convert this to an int ")
	}
	return int64(number), nil
}

func convertIntegerToString(n1, n2, answer int) (string, string) {
	number1 := strconv.Itoa(n1)
	number2 := strconv.Itoa(n2)
	result := strconv.Itoa(answer)
	question := number1 + " + " + number2
	return question, result
}

//func ConvertStructToString(data []domain.InternalUser) ([]string, error) {
//
//	var formattedData []string
//
//	for _, entry := range data {
//
//		formattedData = append(formattedData, fmt.Sprintf("Name : %s, Point : %d", entry.NickName, entry.Point))
//
//	}
//	return formattedData, nil
//}
