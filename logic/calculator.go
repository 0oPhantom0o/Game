package logic

import (
	"fmt"
	"game/repository"
	"math/rand"
)

func Question(id string) (string, error) {
	question, answer := QuestionGenerator()
	err := repository.RedisDataSet(id, answer, "question")
	if err != nil {
		return "", fmt.Errorf("data didnt inserted")
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
