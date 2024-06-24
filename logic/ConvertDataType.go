package logic

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertStringToPrimivite(Id string) (primitive.ObjectID, error) {
	ID, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return ID, nil
}

func ConvertBsonDToScoreBoard(results []bson.D) ([]string, error) {

	var mergedData []string

	var nickName string
	var point int
	type KeyValue struct {
		Key   string      `json:"Key"`
		Value interface{} `json:"Value"`
	}
	var data []KeyValue
	for _, result := range results {
		jsonData, _ := json.Marshal(result)
		strData := string(jsonData)
		_ = json.Unmarshal([]byte(strData), &data)
		for _, item := range data {
			switch item.Key {
			case "nickName":
				if val, ok := item.Value.(string); ok {
					nickName = val
				}
			case "point":
				if val, ok := item.Value.(float64); ok {
					point = int(val)
				}
			}
		}
		if nickName != "" && point != 0 {
			mergedData = append(mergedData, fmt.Sprintf("NickName : %s, Point : %d", nickName, point))
		}

	}

	return mergedData, nil
}
