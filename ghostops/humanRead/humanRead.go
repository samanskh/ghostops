package humanread

import (
	"fmt"
	"ghostops/analyze"
)

func maxLogSeverity(LogSeverity any) string{
	switch logsever := LogSeverity.(type) {
	case string:
		return logsever
	case map[string]int:
		var maxKey string
		var maxValue int
		for key,val := range logsever{
			if val > maxValue{
				maxKey = key
			}
		}
		return maxKey
	}
	return LogSeverity.(string)
}


func HumanRead(stats analyze.Stats){
	fmt.Printf("there were %d valid lines \n",stats.ValidLinesCount)
	fmt.Printf("there were %d invalid lines \n",stats.InvalidLinesCount)
	if stats.InvalidLinesCount != 0{
		fmt.Printf("%d lines had invalid timestamp field \n",stats.InvalidFieldCount["timestamp"])
		fmt.Printf("%d lines had invalid severity field \n",stats.InvalidFieldCount["severity"])
		fmt.Printf("%d lines had invalid service field \n",stats.InvalidFieldCount["service"])
		fmt.Printf("%d lines had invalid message field \n",stats.InvalidFieldCount["message"])
	}
	fmt.Println("======= severity stats =======")
	for key,value := range stats.SeverityCounts{
		fmt.Printf("there were %d lines with %v severity \n",value,key)
	}

	fmt.Println("======= service stats =======")
	for key,value := range stats.ServiceCounts{
		fmt.Printf("there were %d lines with %v service \n",value,key[:len(key) - 8]) // to remove -service at the end of each service
	}

	fmt.Println("======= incident stats =======")
	for key,value := range stats.RuleStats{
		fmt.Printf("incident %v happend with %v severity with %v log severity %d times (first happened at %v last happened at %v) \n",key,value.Severity,maxLogSeverity(value.LogSeverity),value.Count,value.StartTime.String()[:19],value.LastTime.String()[:19])
	}

}