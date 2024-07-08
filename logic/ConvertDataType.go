package logic

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

func ConvertStringToPrimitive(Id string) (primitive.ObjectID, error) {
	ID, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return ID, nil
}

//
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

func ConvertStringToInteger(stringNumber string) (int64, error) {
	number, err := strconv.Atoi(stringNumber)
	if err != nil {
		return 0, fmt.Errorf("can't convert this to an int ")
	} else {
		fmt.Println(number)
	}
	return int64(number), nil
}

func ConvertIntegerToString(n1, n2, answer int) (string, string) {
	number1 := strconv.Itoa(n1)
	number2 := strconv.Itoa(n2)
	result := strconv.Itoa(answer)
	question := number1 + " + " + number2
	return question, result
}

//type IntToStr struct {
//	number1 int
//	number2 int
//	result int
//}
//
//func (i IntToStr) conv() (n1,n2,r string) {
//	n1 := strconv.Itoa(i.number1)
//	n2 := strconv.Itoa(i.number2)
//	r := strconv.Itoa(i.result)
//	return n1,n2,r
//
//}
