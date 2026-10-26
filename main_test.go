package main

import(
	"testing"
	"ghostops/read"
	"ghostops/analyze"
)

func TestStats(t *testing.T) {
	filePath := "sample.log"
	stats := analyze.NewStats()
	err := read.ReadFile(filePath, stats)
	numberOfLines := 6000	
	if err != nil {
		t.Fatalf("error in reading file: %v", err)
	}
	sumValid := stats.ValidLinesCount + stats.InvalidLinesCount
	if sumValid != numberOfLines {
		t.Errorf("there are %d line missing in validating", sumValid)
	}
	sumSeverity := 0
	for _,value := range stats.SeverityCounts{
		sumSeverity += value
	}
	if sumSeverity + stats.InvalidLinesCount != numberOfLines {
		t.Errorf("there are %d line missing in seperating the severities", sumSeverity)
	}
	sumService := 0
	for _,value := range stats.ServiceCounts{
		sumService += value
	}
	if sumService + stats.InvalidLinesCount != numberOfLines {
		t.Errorf("there are %d line missing in seperating the services", sumService)
	}

}