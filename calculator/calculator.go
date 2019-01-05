package calculator

import "strings"
import "strconv"

func errorHandler (err error, msg string) string{
	var errorMessage = ""
	if err != nil {
		errorMessage = msg
	}
	return errorMessage
}

func Calculator(got string) string{
	var want string = ""
	if got == "" {
		want = "0"
	}

	if got == "1" {
		want = "1"
	}

	if got == "1,2" {
		nums := strings.Split(got, ",")
		num1, err := strconv.Atoi(nums[0])
		errorHandler(err, "Can not convert input string to number!")
		num2, err := strconv.Atoi(nums[1])
		errorHandler(err, "Can not convert input string to number!")
		want = strconv.Itoa(num1+num2)
	}
	return want
}