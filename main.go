package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

// import needed packages

type ScrapedText struct {
	Text string `selector:"p"`
}

func main() {

	urls := []string{
		"https://en.wikipedia.org/wiki/Robotics",
		"https://en.wikipedia.org/wiki/Robot",
		"https://en.wikipedia.org/wiki/Reinforcement_learning",
		"https://en.wikipedia.org/wiki/Robot_Operating_System",
		"https://en.wikipedia.org/wiki/Intelligent_agent",
		"https://en.wikipedia.org/wiki/Software_agent",
		"https://en.wikipedia.org/wiki/Robotic_process_automation",
		"https://en.wikipedia.org/wiki/Chatbot",
		"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
		"https://en.wikipedia.org/wiki/Android_(robot)",
	}

	var holder []ScrapedText

	c := colly.NewCollector()

	c.OnHTML("div.mw-body-content", func(e *colly.HTMLElement) {

		text := ScrapedText{}

		text.Text = e.ChildText("p")

		if text != (ScrapedText{}) {
			holder = append(holder, text)
		}

	})

	for index, url := range urls {
		fmt.Println(url)
		fmt.Println(index)
		c.Visit(url)
	}

	// convert an array of structs to JSON using marshaling functions from the encoding/json package
	jsonData, err := json.MarshalIndent(holder, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	// Open the JSONL file that will be written to
	file, err := os.Create("scraped.jl")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a buffered writer for efficient writing
	writer := bufio.NewWriter(file)

	var jsonArray []map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &jsonArray)
	if err != nil {
		panic(err)
	}

	for _, element := range jsonArray {
		data, err := json.Marshal(element)
		if err != nil {
			panic(err)
		}

		// Append a newline after each JSON object
		data = append(data, '\n')

		// write each element in json array to jsonl file
		if _, err := writer.Write(data); err != nil {
			panic(err)
		}
	}
	writer.Flush()

}
