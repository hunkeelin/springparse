package springparse

import (
	"context"
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
	RawLog     string    `json:"rawlog"`     // RawLog
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
		// Ignoring that part of the log
		return nil
	}
	err = getkubeInfo(getkubeInfoInput{
		fileName: s.fileName,
		es:       &out.content,
	})
	rDate := fmt.Sprintf(time.Now().UTC().Format("2006-01-02"))
	client, err := newElasticClient(awsCredentials)
	if err != nil {
		return err
	}
	ctx := context.Background()
	put, err := client.Index().
		Index(logPrefix + "-" + rDate).
		Type("springparse").
		Id(out.id).
		BodyJson(out.content).
		Do(ctx)
	if err != nil {
		return err
	}
	log.Info(fmt.Sprintf("Index %s created with id %v", put.Index, put.Id))
	return nil
}
