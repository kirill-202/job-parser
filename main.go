package main

import (
	"fmt"
	//"net/http"
	//"net/smtp"
	"os"
	//"golang.org/x/net/html"
	"bufio"
	"strings"
)

//https://startup.jobs 



type Job struct {
	Name string
	WebAddress string
}



// func (page LandingPage) FindJobDescrpition(JobDescription string) (Job, error) {
// 	fmt.Println("not implemented yet")
// }

func traverse() {

}

func getInput(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		if scanner.Scan() {
			line := scanner.Text()
			trimmed_line := strings.TrimSpace(line)
			if len(trimmed_line) > 0 {
				return trimmed_line
			}
			fmt.Println("Input cannot be empty. Please try again.")
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}
	}
}



func main() {

	var KeyWords []string
	var BaseWebsite string

	fmt.Println("Program has started...")

	KeyWords = strings.Fields(getInput("Please provide key words for job search: "))
	BaseWebsite = getInput("Please provide the job website to parse in the format 'website.com': ")

	fmt.Println("Scanned words:", KeyWords)
	fmt.Println("Scanned website:", BaseWebsite)
	fmt.Println("Parsing has started...")

}