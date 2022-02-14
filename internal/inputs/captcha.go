package inputs

import "fmt"

func WaitInput(text string) string {
	var captcha string
	fmt.Println(text)
	fmt.Scanf("%s", &captcha)

	return captcha
}