package main

import (
	_ "embed"
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
)

//go:embed day.go.template
var dayTemplate string

// TODO: Refactor this to parse flags then hand off to helper client or something testable
// TODO: Implement tests for helper client
// TODO: Order should be 1) create directory, 2) create files, 3) write files
func main() {
	day := flag.Int("day", 0, "day number to initialize")
	flag.Parse()

	if *day < 1 || *day > 25 {
		log.Fatalf("invalid day number provided (must be between 1 and 25)")
	}

	err := os.Mkdir(fmt.Sprintf("src/day%02d", *day), 0755)
	if err != nil {
		log.Fatalf("failed to create new day directory: %v", err)
	}

	t := template.Must(template.New("day").Parse(dayTemplate))
	t.Execute(os.Stdout, struct {
		DayNum int
	}{
		DayNum: *day,
	})
}
