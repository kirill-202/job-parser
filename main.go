package main

import (
	"fmt"
	"io"

	"net/http"
	"strings"

	"golang.org/x/net/html"
)

//https://startup.jobs
type JobSearch struct {
	Title JobTitle
	WebAddress string
}




func (js *JobSearch) getWebPage() ([]string, error) {
	url := "https://" + js.WebAddress

	response, err := http.Get(url)
	if err != nil {
		return nil , fmt.Errorf("error fetching the url: %s. Http code %v", url, response.StatusCode)
	}

	defer response.Body.Close()
	
	topNode, err := html.Parse(response.Body)
	if err != nil {
		return nil , fmt.Errorf("error reading the response body: %e", err)
	}
	

	searchTerms := []string{js.Title.Name}
	for _, synonym := range js.Title.Synonyms {
		searchTerms = append(searchTerms, synonym.Word)
	}

	var results []string

	for _, term := range searchTerms {
		temp_results := traverse(topNode, term)

		if len(temp_results) > 0 {
			results = append(results, temp_results...)
		}
	}

	fmt.Println("the full list of elems with term is: ", results)
	return results, nil
	

}

func (js *JobSearch) getWebPageFull() error {
	url := "https://" + js.WebAddress

	response, err := http.Get(url)
	if err != nil {
		return  fmt.Errorf("error fetching the url: %s. Http code %v", url, response.StatusCode)
	}

	defer response.Body.Close()
	
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading the response body: %e", err)
	}
	fmt.Println(string(body))
	return nil
}


// func parsePrintTest(path string) error {
// 	file, err := os.Open(path)

// 	if err != nil {
// 		return fmt.Errorf("error pasrsing file: %e", err)
// 	}
// 	topNode, _ := html.Parse(file)
	

// 	var results []string

// 	test_terms := []string{"good term", "bad term"}
// 	for _, term := range test_terms {
// 		temp_results := traverse(topNode, term)

// 		if len(temp_results) > 0 {
// 			results = append(results, temp_results...)
// 		}
// 	}
	
// 	if len(results) == 0 {
// 		return fmt.Errorf("no terms found in the HTML file")
// 	}
// 	fmt.Printf("These are the terms found %v\n", results)
	
// 	return nil
// }


func traverse(node *html.Node, term string) []string {
	var nodeLists []string

	if node.Type == html.TextNode && strings.Contains(node.Data, term) {
		nodeLists = append(nodeLists, node.Data)
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		childNodeLists := traverse(child, term)
		nodeLists = append(nodeLists, childNodeLists...)
	}

	return nodeLists
}

func main() {

	var job JobSearch

	fmt.Println("Program has started...")

	job.Title.Name = getInput("Please provide key words for job search: ")
	job.WebAddress = getInput("Please provide the job website to parse in the format 'website.com': ")

	fmt.Println("Scanned words:", job.Title.Name)
	fmt.Println("Scanned website:", job.WebAddress)

	fmt.Println("Parsing has started...")

	job.Title.getSynonyms()

	eer := job.getWebPageFull()
	fmt.Println(eer)


}