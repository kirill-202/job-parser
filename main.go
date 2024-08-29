package main

import (
	"fmt"

	//"net/smtp"
	// "golang.org/x/net/html"
)

//https://startup.jobs



type JobSearch struct {
	Title JobTitle
	WebAddress string
}


// func FindJobDescrpition(jobDescription string) (Job, error) {
// 	html.

// 	fmt.Println("not implemented yet")
// }

// func traverse() {

// }




func main() {

	var job JobSearch

	fmt.Println("Program has started...")

	job.Title.Name = getInput("Please provide key words for job search: ")
	job.WebAddress = getInput("Please provide the job website to parse in the format 'website.com': ")

	fmt.Println("Scanned words:", job.Title.Name)
	fmt.Println("Scanned website:", job.WebAddress)
	fmt.Println("Parsing has started...")
	job.Title.getSynonyms()

}