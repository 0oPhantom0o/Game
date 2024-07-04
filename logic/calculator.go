package logic

import (
	"fmt"
	"game/repository"
	"math/rand"
)

func Question(id string) (string, error) {
	question, answer := QuestionGenerator()
	err := repository.RedisDataSet(id, answer, "1")
	if err != nil {
		return "", err
	}
	boolError := repository.ExpireTime(id)

	if boolError != true {
		return "", fmt.Errorf("error in expire section")
	}
	return question, nil
}

func QuestionGenerator() (string, string) {
	number1 := rand.Intn(10)
	number2 := rand.Intn(10)
	answer := number1 + number2
	question, result := ConvertIntegerToString(number1, number2, answer)
	fmt.Println(question)
	return question, result
}
