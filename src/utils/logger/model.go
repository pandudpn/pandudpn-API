package logger

import (
	"context"
	"sync"
	"time"
)

// Locker is container data
type Locker struct {
	data sync.Map
}

type (
	// Key of context
	Key int
	// Flags is key for store context
	Flags string
	// ServiceType is type for data logging
	ServiceType string
)

const (
	// logKey is key context for request http rest api
	logKey = Key(31)

	// rest is type for logging http rest api
	rest ServiceType = "http_rest_api"

	_StatusCode   Flags = "StatusCode"
	_Response     Flags = "Response"
	_LogMessages  Flags = "LogMessages"
	_ThirdParties Flags = "ThirdParties"
	_ErrorMessage Flags = "ErrorMessage"

	// list type of logger
	debug   = "DEBUG"
	print   = "PRINT"
	err     = "ERROR"
	success = "Success request"
)

// DataLogger is standard output to terminal
type DataLogger struct {
	RequestId     string                 `json:"requestId"`
	Type          ServiceType            `json:"type"`
	TimeStart     time.Time              `json:"timeStart"`
	Service       string                 `json:"service"`
	Host          string                 `json:"host"`
	Endpoint      string                 `json:"endpoint"`
	RequestMethod string                 `json:"requestMethod"`
	RequestHeader map[string]interface{} `json:"requestHeader"`
	RequestBody   map[string]interface{} `json:"requestBody"`
	StatusCode    int                    `json:"statusCode"`
	Response      interface{}            `json:"response"`
	ErrorMessage  string                 `json:"errorMessage"`
	ExecTime      float64                `json:"execTime"`
	LogMessages   []LogMessage           `json:"logMessage"`
	ThirdParties  []ThirdParty           `json:"outgoing_log"`
}

// LogMessage is data logging for developer want to debug or error
type LogMessage struct {
	File    string `json:"file"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

// ThirdParty is data logging for any request to third party
type ThirdParty struct {
	ServiceTarget string      `json:"serviceTarget"`
	URL           string      `json:"url"`
	RequestHeader interface{} `json:"requestHeader"`
	RequestBody   interface{} `json:"requestBody"`
	Response      interface{} `json:"response"`
	Method        string      `json:"method"`
	StatusCode    int         `json:"statusCode"`
	ExecTime      float64     `json:"execTime"`
}

// iface for get unsafe context data
type iface struct {
	itab, data uintptr
}

// valueCtx for slicing key and value of context
type valueCtx struct {
	context.Context
	key, val interface{}
}

// String convert
func (st ServiceType) String() string {
	return string(st)
}
