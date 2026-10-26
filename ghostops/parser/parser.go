package parser

import (
	"ghostops/shared"
	"strings"
	"time"
	"regexp"
)

func Parser(line string) (shared.Line,[]bool){
	fields := strings.SplitN(line, " ", 4)
	var timestamp,severity, service, message string
	if len(fields) > 0 {
		timestamp = fields[0]
	}
	if len(fields) > 1 {
		severity = fields[1]
	}
	if len(fields) > 2 {
		service = fields[2]
	}
	if len(fields) > 3 {
		message = fields[3]
	}
	event := shared.Line{Timestamp: timestamp,Severity: severity,Service: service,Message: message}
	validate := totalValidate(event)
	return event,validate
}

func totalValidate(line shared.Line) []bool {
	timestamp := validateTimestamp(line.Timestamp)
	severity := validateSeverity(line.Severity)
	service := validateService(line.Service)
	message := validateMessage(line.Message)
	bools := []bool{timestamp,severity,service,message}
	return bools
}

func validateTimestamp(s string) bool {
    _, err := time.Parse(time.RFC3339, s)
    return err == nil
}

var severity = map[string]bool{
    "INFO": true, "WARN": true, "ERROR": true, "FATAL": true,
	 "**INFO**":true, "**WARN**": true, "**ERROR**": true, "**FATAL**": true,
}

func validateSeverity(s string) bool {
    return severity[s]
}

var serviceRe = regexp.MustCompile(`^[a-z][a-z0-9-]*-service$`)

func validateService(s string) bool {
    return serviceRe.MatchString(s)
}

func validateMessage(s string) bool{
	return s != ""
}