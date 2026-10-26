package read

import (
	"bufio"
	// "fmt"
	"os"
	// "ghostops/shared"
	// "ghostops/parser"
	"ghostops/shared"
	"ghostops/keyword"
)

type LineAnalyzer interface {
    Analyze(line string)
}

func ReadFile(path string, analyzer LineAnalyzer) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	defer f.Close()
	shared.KeyWords = keyword.MakeTotalKeywords(shared.Incidents)
	// stats := analyze.NewStats()
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 0, 64*1024), 10*1024*1024)
	for scanner.Scan() {
        line := scanner.Text()
		analyzer.Analyze(line)
		
        // fmt.Println(line)
    }
	if err := scanner.Err(); err != nil {
        return err
    }

	return nil 
}