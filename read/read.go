package read

import (
	"bufio"
	// "fmt"
	"os"
	// "ghostops/shared"
	// "ghostops/parser"
	"ghostops/analyze"
	"ghostops/shared"
	"ghostops/keyword"
)

func ReadFile(path string) (analyze.Stats, error) {
	f, err := os.Open(path)
	if err != nil {
		return analyze.Stats{}, err
	}

	defer f.Close()
	shared.KeyWords = keyword.MakeTotalKeywords(shared.Incidents)
	stats := analyze.NewStats()
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 0, 64*1024), 10*1024*1024)
	for scanner.Scan() {
        line := scanner.Text()
		stats.Analyze(line)
		
        // fmt.Println(line)
    }
	if err := scanner.Err(); err != nil {
        return analyze.Stats{},err
    }

	return *stats,nil //گفتم خود استس رو برگردونم نه آدرسشو که مثلا از مین کپیشو بگیرن و دسترسی برای تغییرش نداشته باشن 
}