package springparse

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

type sendElasticSearchInput struct {
	rawLog   []byte
	fileName string
}
type elasticOut struct {
	TimeStamp  time.Time `json:"@timestamp"` // TimeStamp
	LogLevel   string    `json:"level"`      // LogLevel the log level of the log
	Thread     string    `json:"thread"`     // Thread
	LoggerName string    `json:"loggername"` // LoggerName
	ProcessId  string    `json:"processid"`  // ProcessId
	RawLog     string    `json:"log"`        // RawLog
	FileName   string    `json:"filename"`   // FileName
	KubeInfo   kubeInfo  `json:"kubernetes"` // KubeInfo
}

func (r *Runner) sendElasticSearch(s sendElasticSearchInput) error {
	out, err := r.parseLog(parseLogInput{
		fileName: s.fileName,
		rawLog:   s.rawLog,
	})
	if err != nil {
		return err
	}
	if out.content.LogLevel == "" {
		if r.Buffer != nil {
			r.Buffer.RawLog = r.Buffer.RawLog + "\n" + out.content.RawLog
		} else {
			// Ignoring that part of the log
			log.Info(fmt.Sprintf("It seems %v isn't a springboot log", s.fileName))
		}
		return nil
	}
	err = getkubeInfo(getkubeInfoInput{
		fileName: s.fileName,
		es:       &out.content,
	})
	if err != nil {
		return err
	}
	if r.Buffer == nil {
		r.Buffer = &out.content
		r.BufferId = out.id
		return nil
	}

	// Sending it to bulk
	rDate := fmt.Sprintf(time.Now().UTC().Format("2006-01-02"))
	sendItems <- elasticItem{
		index:    logPrefix + "-" + rDate,
		id:       r.BufferId,
		bodyJson: *r.Buffer,
	}
	r.Buffer = &out.content
	r.BufferId = out.id
	return nil
}
