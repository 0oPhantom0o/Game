package logic

import (
	"fmt"
	"game/repository"
	"math/rand"
)

func Question(id string) (string, error) {

	question, answer := QaGenerator()
	err := repository.InsertAnswer(id, answer)
	if err != nil {
		return "", fmt.Errorf("data didnt inserted")
	}

	return question, nil
}

func QaGenerator() (string, string) {
	//generate question and answer
	number1 := rand.Intn(10)
	number2 := rand.Intn(10)
	result := number1 + number2
	question, answer := convertIntegerToString(number1, number2, result)
	fmt.Println(question)
	return question, answer
}
