package matchrule

import (
	"ghostops/shared"
	"ghostops/keyword"
)

func RuleMatcher(line shared.Line) shared.Rule {
	message := line.Message
	messageKeywords := keyword.KeywordFinder(message)
	for _, rule := range shared.Rules {
		count := make(map[string]bool)
		for _, message := range messageKeywords {
			if rule.Keywords[message] {
				count[message] = true
			}
		}
		if len(count) == len(rule.Keywords) {
			return rule
		}
	}
	return shared.Rule{}
}