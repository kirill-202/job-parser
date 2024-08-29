package main

import (
	"bufio"
	"strings"
	"os"
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

// good score for https://www.datamuse.com/api/
const score int = 20000000

//logic related to scanning input from users

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



//logic related to adding titles


type JobTitle struct {
	Name string
	Synonyms []DataMuseSynonym
}


type DataMuseSynonym struct {
	Word string `json:"word"`
	Score int `json:"score"`
}

func (job *JobTitle) getSynonyms() error {
	if job.Name == "" {
		return fmt.Errorf("no title found in job search: %v", job)
	}


	synonyms, err := fetchSynonymsFromAPI(job.Name)
	if err != nil {
		return err
	}


	job.Synonyms = filterSynonyms(synonyms, score)
	fmt.Printf("Filtered synonyms: %v\n", job.Synonyms)

	return nil
}


func fetchSynonymsFromAPI(title string) ([]DataMuseSynonym, error) {
	processedTitle := strings.Replace(title, " ", "+", -1)
	url := fmt.Sprintf("https://api.datamuse.com/words?ml=%s", processedTitle)

	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching synonyms: %w", err)
	}
	
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	var synonyms []DataMuseSynonym
	if err := json.Unmarshal(body, &synonyms); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	fmt.Printf("Unfiltered synonyms: %v\n", synonyms)
	return synonyms, nil
}


func filterSynonyms(synonyms []DataMuseSynonym, minScore int) []DataMuseSynonym {
	var filtered []DataMuseSynonym
	for _, syn := range synonyms {
		if syn.Score > minScore {
			filtered = append(filtered, syn)
		}
	}
	return filtered
}
