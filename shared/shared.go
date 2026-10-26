package shared

import "time"

type Line struct {
	Timestamp string
	Severity  string
	Service   string
	Message   string
}

type Incident struct {
	IncidentName         string
	Keywords         map[string]bool
	IncidentSeverity string
}

type IncidentStat struct {
	Count       int
	LogSeverity any //string or map[string]int
	Severity    string
	StartTime   time.Time 
	LastTime    time.Time
}

var Incidents = []Incident{
	{
		IncidentName: "DB_POOL_EXHAUSTED",
		Keywords: map[string]bool{"database": true, "connection": true, "pool": true, "exhausted": true},
		IncidentSeverity: "CRITICAL",
	},
	{
		IncidentName: "DB_CONNECTION_FAILURE",
		Keywords: map[string]bool{"failed": true, "connect": true, "database": true},
		IncidentSeverity: "HIGH",
	},
	{
		IncidentName: "PAYMENT_GATEWAY_TIMEOUT",
		Keywords: map[string]bool{"payment": true, "gateway": true, "timeout": true},
		IncidentSeverity: "HIGH",
	},
	{
		IncidentName: "PAYMENT_GATEWAY_502",
		Keywords: map[string]bool{"payment": true, "gateway": true, "502": true},
		IncidentSeverity: "HIGH",
	},
	{
		IncidentName: "SQL_SYNTAX_ERROR",
		Keywords: map[string]bool{"sql": true, "syntax": true, "error": true},
		IncidentSeverity: "HIGH",
	},
	{
		IncidentName: "NULL_POINTER_EXCEPTION",
		Keywords: map[string]bool{"nullpointerexception": true},
		IncidentSeverity: "HIGH",
	},
	{
		IncidentName: "DB_POOL_THRESHOLD",
		Keywords: map[string]bool{"database": true, "connection": true, "pool": true, "threshold": true},
		IncidentSeverity: "MEDIUM",
	},
	{
		IncidentName: "HIGH_MEMORY_USAGE",
		Keywords: map[string]bool{"high": true, "memory": true, "usage": true},
		IncidentSeverity: "MEDIUM",
	},
	{
		IncidentName: "RATE_LIMIT_WARNING",
		Keywords: map[string]bool{"rate": true, "limit": true},
		IncidentSeverity: "MEDIUM",
	},
	{
		IncidentName: "DEPRECATED_API_USAGE",
		Keywords: map[string]bool{"deprecated": true, "api": true},
		IncidentSeverity: "LOW",
	},
}

// var KeyWords = map[string]bool{"database":true,"connection":true,"pool":true,"exhausted":true,
// 								"failed":true, "connect":true,"payment":true , "gateway":true , "timeout":true,
// 							"502":true, "sql":true, "syntax":true , "error":true, "nullpointerexception":true , "threshold":true,
// 							"high":true, "memory":true, "usage":true, "rate":true, "limit":true, "deprecated":true, "api":true}

var KeyWords map[string]bool // I built this with code from rule list not hard code like above