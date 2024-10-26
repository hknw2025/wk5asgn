# Week 5 Asignment - Scraping the Web

This tool takes a list of URLs and scrapes the text data from the webpages and saves it in the form of a JSON lines file. 

## How to use this tool:
* After cloning this git repository you can 
`go run main.go 
* you will have a new file in the repository with jl extension with the text data from the pages desired.

## Example:

urls := []string{
		"https://en.wikipedia.org/wiki/Robotics",
}

Excerpt from saved file:

{"Text":"Robotics is the interdisciplinary study and practice of the design, construction, operation, and use of robots.[1]\nWithin mechanical engineering, etc... }

I was incredibly impressed by the processing time on this program. It was much faster than similar programs I have run in the past in both R and Python. I was also impressed with how easy to use the web scraping tools were in Go. I really enjoyed the syntax they were written in. Having a collector that takes certain actions for each URL passed to it seemed to not only be a fast processing structure, but it was also very easy to understand and write. 



