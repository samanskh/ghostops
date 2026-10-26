package analyze

import (
	"ghostops/matchRule"
	"ghostops/parser"
	"ghostops/shared"
	"time"
	// "strings"
)
type Stats struct{
	ValidLinesCount int
	InvalidLinesCount int
	InvalidFieldCount map[string]int 
	SeverityCounts map[string]int
	ServiceCounts map[string]int
	RuleStats map[string]shared.IncidentStat
}

func NewStats() *Stats{
	var statsInfo = Stats{ValidLinesCount:0, InvalidLinesCount:0, InvalidFieldCount:make(map[string]int),
	 SeverityCounts:make(map[string]int), ServiceCounts:make(map[string]int),RuleStats: make(map[string]shared.IncidentStat),}
	return &statsInfo
}
// func NewRuleStat() *map[*shared.Rule]int{
// 	var ruleStat = make(map[*shared.Rule]int)
// 	return &ruleStat
// }

func (s *Stats) Analyze(line string){
	event,bools := parser.Parser(line)
	if !(bools[0] && bools[1] && bools[2] && bools[3]){
		s.InvalidLinesCount ++
	}else{
		s.ValidLinesCount ++
		severityKey := event.Severity
		s.SeverityCounts[severityKey] ++
		serviceKey := event.Service
		s.ServiceCounts[serviceKey] ++
		rule := matchrule.RuleMatcher(event)
		if rule.RuleName != "" {
			ts, err := time.Parse(time.RFC3339, event.Timestamp)
			if err != nil {
				ts = time.Time{}
			}
			
			stat := s.RuleStats[rule.RuleName]
			if stat.Count == 0{
				stat = shared.IncidentStat{Count:0,Severity: rule.IncidentSeverity, LogSeverity:make(map[string]int), StartTime: ts,LastTime: ts,}
			}else{
				if ts.Before(stat.StartTime){
					stat.StartTime = ts
				}
				if ts.After(stat.LastTime){
					stat.LastTime = ts
				}
			}
			stat.Count++
			logSeverity := stat.LogSeverity.(map[string]int)
			logSeverity[severityKey]++
			s.RuleStats[rule.RuleName] = stat
		}

		
	}
	if !bools[0]{
		s.InvalidFieldCount["timestamp"] ++
	}
	if !bools[1]{
		s.InvalidFieldCount["severity"] ++
	}
	if !bools[2]{
		s.InvalidFieldCount["service"] ++
	}
	if !bools[3]{
		s.InvalidFieldCount["message"] ++
	}

}


