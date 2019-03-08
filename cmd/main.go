package main

import (
	"fmt"

	auth "github.com/zerospiel/i3-gmail-checker/pkg"
)

func main() {
	str, _ := auth.GenerateAuthURL()
	fmt.Println(str)
}
