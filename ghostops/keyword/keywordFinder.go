package keyword

import (
	"ghostops/shared"
	"strings"
)

func MakeTotalKeywords(ruleList []shared.Rule) map[string]bool{
	output := make(map[string]bool)
	for _,r := range ruleList{
		for key,_ := range r.Keywords{
			output[key] = true
		}
	}	
	return output
}

func KeywordFinder(message string) []string{
	stringLower := strings.ToLower(message)
	words := strings.Fields(stringLower)
	output := make([]string,0,len(shared.KeyWords))
	for _,word := range words{
		if shared.KeyWords[word]{
			output = append(output,word)
			}	
	}
	return output
}

