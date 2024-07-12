package logic

import (
	"fmt"
	"game/constants"
	"math/rand"
	"time"
)

func RandomCode() (string, error) {

	code := CreateRandomCode()
	fmt.Println(code)
	return code, nil
}

func CreateRandomCode() string {
	numbers := constants.OtpCharacters
	codeLength := constants.OtpCodeLength
	code := make([]byte, codeLength)

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	for i := 0; i < codeLength; i++ {
		randCharNum := rng.Intn(len(numbers))
		code[i] = numbers[randCharNum]
	}
	return string(code)

}
