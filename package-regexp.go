package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eto"))
	fmt.Println(regex.MatchString("eDo"))

	fmt.Println(regex.FindAllString("eko edo eRo eso Eda", 10))

	regexMail := regexp.MustCompile(`([@])`)
	fmt.Println(regexMail.MatchString("refo@email.com"))
	fmt.Println(regexMail.MatchString("refoemail.com"))
}
