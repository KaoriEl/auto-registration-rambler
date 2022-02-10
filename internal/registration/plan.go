package registration

import (
	"fmt"
	"main/internal/structures"
	"pkg.re/essentialkaos/translit.v2"
	"regexp"
)

func Plan(i structures.AccInfo) structures.AccInfo {

	i.Email = translit.EncodeToICAO(i.Name)

	r := regexp.MustCompile("\\s+")

	i.Email = r.ReplaceAllString(i.Email, "")

	i.Email = i.Email + GenerateNum(10)
	fmt.Println(i.Email)

	return i
}
