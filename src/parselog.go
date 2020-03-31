package springparse

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type rawLog struct {
	Log    string    `json:"log"`
	Stream string    `json:"stream"`
	Time   time.Time `json:"time"`
}

type parseLogOutput struct {
	id      string
	content elasticOut
}

func parseLog(s []byte) (parseLogOutput, error) {
	var p rawLog
	var level, processId, thread, loggerName string
	s = bytes.Replace(s, []byte("\n"), []byte(""), -1)
	err := json.Unmarshal(s, &p)
	if err != nil {
		return parseLogOutput{}, err
	}
	logSplit := strings.Split(p.Log, " ")
	for _, word := range logSplit {
		// Get level
		if word == "INFO" || word == "DEBUG" || word == "ERROR" || word == "WARN" && level == "" {
			level = word
		}

		/// Get process Id
		_, err := strconv.Atoi(word)
		if err == nil && processId == "" {
			processId = word
		}
	}

	thread = getThread(p.Log)
	loggerName = getLoggerName(p.Log)
	return parseLogOutput{
		id: "",
		content: elasticOut{
			TimeStamp:  p.Time,
			RawLog:     p.Log,
			Thread:     thread,
			LoggerName: loggerName,
			ProcessId:  processId,
			LogLevel:   level,
		},
	}, nil
}

func getThread(s string) string {
	var begin, end int
	for idx, char := range s {
		if string(char) == "[" && begin == 0 {
			begin = idx + 1
		}
		if string(char) == "]" && end == 0 && begin != 0 {
			end = idx
		}
	}
	// Get thread
	if end > begin {
		return strings.Replace(s[begin:end], " ", "", -1)
	}
	return ""
}
func getLoggerName(s string) string {
	var begin, end, loggerbegin, loggerend, spacecount int
	for idx, char := range s {
		if string(char) == "[" && begin == 0 {
			begin = idx + 1
		}
		if string(char) == "]" && end == 0 && begin != 0 && loggerbegin == 0 {
			end = idx
			loggerbegin = idx + 1
		}
		if string(char) == " " && loggerbegin != 0 {
			spacecount++
		}
		if spacecount == 2 && loggerend == 0 {
			loggerend = idx
		}
	}
	// Get thread
	if loggerend > loggerbegin {
		return strings.Replace(s[loggerbegin:loggerend], " ", "", -1)
	}
	return ""
}
