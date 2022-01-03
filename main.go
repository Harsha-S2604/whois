package main

import (
	"github.com/likexian/whois"
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Print("Enter the domain name(example google.com): ")
	var domainName string
	fmt.Scanf("%s", &domainName)
	fmt.Println(domainName)

	fmt.Println("Fetching the domain information...")
	results, err := whois.Whois(domainName)
	if err != nil {
		fmt.Println("Error in fetching the information. Please try again later", err.Error())
	} else {
		err := ioutil.WriteFile("info.txt", []byte(results), 0644)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("Fetching possible information from social media's...")
	}

}
