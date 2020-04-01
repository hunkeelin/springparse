package springparse

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
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
type parseLogInput struct {
	rawLog   []byte
	fileName string
}
type elasticOut struct {
	TimeStamp  time.Time `json:"@timestamp"` // TimeStamp
	LogLevel   string    `json:"level"`      // LogLevel the log level of the log
	Thread     string    `json:"thread"`     // Thread
	LoggerName string    `json:"loggername"` // LoggerName
	ProcessId  string    `json:"processid"`  // ProcessId
	RawLog     string    `json:"rawlog"`     // RawLog
	FileName   string    `json:"filename"`   //FileName
}

func (r *Runner) parseLog(s parseLogInput) (parseLogOutput, error) {
	var p rawLog
	var level, processId, thread, loggerName string
	rawlogNoESC := bytes.Replace(s.rawLog, []byte("\n"), []byte(""), -1)
	err := json.Unmarshal(rawlogNoESC, &p)
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
		id: logHash(s.rawLog),
		content: elasticOut{
			TimeStamp:  p.Time,
			RawLog:     p.Log,
			Thread:     thread,
			LoggerName: loggerName,
			ProcessId:  processId,
			FileName:   s.fileName,
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

func logHash(rawInput []byte) string {
	hasher := md5.New()
	hasher.Write(rawInput)
	return hex.EncodeToString(hasher.Sum(nil))[0:20]
}
