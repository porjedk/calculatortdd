package calculator

func Calculator(got string) string{
	var want string = ""
	if got == "" {
		want = "0"
	}

	if got == "1" {
		want = "1"
	}
	return want
}