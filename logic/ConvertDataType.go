package logic

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

func ConvertStringToPrimivite(Id string) (primitive.ObjectID, error) {
	ID, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return ID, nil
}

func ConvertBsonDToScoreBoard(data []bson.D) ([]string, error) {

	var formattedData []string

	for _, entry := range data {
		var nickName string
		var point int

		for _, item := range entry {
			switch item.Key {
			case "nickName":
				nickName = item.Value.(string)
			case "point":
				point = int(item.Value.(int32))
			}
		}

		// Format and append to the slice
		formattedData = append(formattedData, fmt.Sprintf("Name : %s, Point : %d", nickName, point))
	}

	return formattedData, nil
}

func ConvertStringToInteger(stringNumber string) (int64, error) {
	number, err := strconv.Atoi(stringNumber)
	if err != nil {
		return 0, fmt.Errorf("can't convert this to an int")
	} else {
		fmt.Println(number)
	}
	return int64(number), nil
}
