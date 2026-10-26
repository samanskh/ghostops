package matchrule

import (
	"ghostops/shared"
	"ghostops/keyword"
)

func IncidentMatcher(line shared.Line) shared.Incident {
	message := line.Message
	messageKeywords := keyword.KeywordFinder(message)
	for _, Incident := range shared.Incidents {
		count := make(map[string]bool)
		for _, message := range messageKeywords {
			if Incident.Keywords[message] {
				count[message] = true
			}
		}
		if len(count) == len(Incident.Keywords) {
			return Incident
		}
	}
	return shared.Incident{}
}