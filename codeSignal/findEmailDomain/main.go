package main

import "fmt"
import "regexp"

func main() {
	fmt.Println("vim-go")
}

func findEmailDomain(email string) string {

	validMailRegex := regexp.MustCompile(`^[A-z]+\@[A-z]+\.[A-z]+$`)
	match := validMailRegex.FindString(email)

	// CHECK IS MAIL IS VALID
	if match != "" {
		domainRegex := regexp.MustCompile(	
	}

	return ""
}
