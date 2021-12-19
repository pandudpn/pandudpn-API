package logger

import (
	"context"
	"fmt"
	"runtime"
	"strings"
)

type logger struct{}

var Log *logger

func (l *logger) Errorf(ctx context.Context, format string, args ...interface{}) {
	var (
		messages []LogMessage
		file     string
	)
	
	value, ok := extract(ctx)
	if !ok {
		return
	}
	
	// for get filename and line when developer called this method
	_, fileName, line, _ := runtime.Caller(1)
	files := strings.Split(fileName, "/")
	if len(files) > 3 {
		file = fmt.Sprintf("%s:%d", strings.Join(files[len(files)-2:], "/"), line)
	} else {
		file = fmt.Sprintf("%s:%d", strings.Join(files, "/"), line)
	}
	
	tmp, ok := value.LoadAndDelete(_LogMessages)
	if ok {
		messages = tmp.([]LogMessage)
	}
	
	message := LogMessage{
		File:    file,
		Type:    err,
		Message: fmt.Sprintf(format, args...),
	}
	
	messages = append(messages, message)
	
	value.Set(_LogMessages, messages)
}

func (l *logger) Error(ctx context.Context, args ...interface{}) {
	var (
		messages []LogMessage
		file     string
	)
	
	value, ok := extract(ctx)
	if !ok {
		return
	}
	
	// for get filename and line when developer called this method
	_, fileName, line, _ := runtime.Caller(1)
	files := strings.Split(fileName, "/")
	if len(files) > 3 {
		file = fmt.Sprintf("%s:%d", strings.Join(files[len(files)-2:], "/"), line)
	} else {
		file = fmt.Sprintf("%s:%d", strings.Join(files, "/"), line)
	}
	
	tmp, ok := value.LoadAndDelete(_LogMessages)
	if ok {
		messages = tmp.([]LogMessage)
	}
	
	message := LogMessage{
		File:    file,
		Type:    err,
		Message: fmt.Sprint(args...),
	}
	
	messages = append(messages, message)
	
	value.Set(_LogMessages, messages)
}

func (l *logger) Debugf(ctx context.Context, format string, args ...interface{}) {
	var (
		messages []LogMessage
		file     string
	)
	
	value, ok := extract(ctx)
	if !ok {
		return
	}
	
	// for get filename and line when developer called this method
	_, fileName, line, _ := runtime.Caller(1)
	files := strings.Split(fileName, "/")
	if len(files) > 3 {
		file = fmt.Sprintf("%s:%d", strings.Join(files[len(files)-2:], "/"), line)
	} else {
		file = fmt.Sprintf("%s:%d", strings.Join(files, "/"), line)
	}
	
	tmp, ok := value.LoadAndDelete(_LogMessages)
	if ok {
		messages = tmp.([]LogMessage)
	}
	
	message := LogMessage{
		File:    file,
		Type:    debug,
		Message: fmt.Sprintf(format, args...),
	}
	
	messages = append(messages, message)
	
	value.Set(_LogMessages, messages)
}

func (l *logger) Debug(ctx context.Context, args ...interface{}) {
	var (
		messages []LogMessage
		file     string
	)
	
	value, ok := extract(ctx)
	if !ok {
		return
	}
	
	// for get filename and line when developer called this method
	_, fileName, line, _ := runtime.Caller(1)
	files := strings.Split(fileName, "/")
	if len(files) > 3 {
		file = fmt.Sprintf("%s:%d", strings.Join(files[len(files)-2:], "/"), line)
	} else {
		file = fmt.Sprintf("%s:%d", strings.Join(files, "/"), line)
	}
	
	tmp, ok := value.LoadAndDelete(_LogMessages)
	if ok {
		messages = tmp.([]LogMessage)
	}
	
	message := LogMessage{
		File:    file,
		Type:    debug,
		Message: fmt.Sprint(args...),
	}
	
	messages = append(messages, message)
	
	value.Set(_LogMessages, messages)
}

func (l *logger) Printf(ctx context.Context, format string, args ...interface{}) {
	var (
		messages []LogMessage
		file     string
	)
	
	value, ok := extract(ctx)
	if !ok {
		return
	}
	
	// for get filename and line when developer called this method
	_, fileName, line, _ := runtime.Caller(1)
	files := strings.Split(fileName, "/")
	if len(files) > 3 {
		file = fmt.Sprintf("%s:%d", strings.Join(files[len(files)-2:], "/"), line)
	} else {
		file = fmt.Sprintf("%s:%d", strings.Join(files, "/"), line)
	}
	
	tmp, ok := value.LoadAndDelete(_LogMessages)
	if ok {
		messages = tmp.([]LogMessage)
	}
	
	message := LogMessage{
		File:    file,
		Type:    print,
		Message: fmt.Sprintf(format, args...),
	}
	
	messages = append(messages, message)
	
	value.Set(_LogMessages, messages)
}

func (l *logger) Print(ctx context.Context, args ...interface{}) {
	var (
		messages []LogMessage
		file     string
	)
	
	value, ok := extract(ctx)
	if !ok {
		return
	}
	
	// for get filename and line when developer called this method
	_, fileName, line, _ := runtime.Caller(1)
	files := strings.Split(fileName, "/")
	if len(files) > 3 {
		file = fmt.Sprintf("%s:%d", strings.Join(files[len(files)-2:], "/"), line)
	} else {
		file = fmt.Sprintf("%s:%d", strings.Join(files, "/"), line)
	}
	
	tmp, ok := value.LoadAndDelete(_LogMessages)
	if ok {
		messages = tmp.([]LogMessage)
	}
	
	message := LogMessage{
		File:    file,
		Type:    print,
		Message: fmt.Sprint(args...),
	}
	
	messages = append(messages, message)
	
	value.Set(_LogMessages, messages)
}
