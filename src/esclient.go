package springparse

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/olivere/elastic"
	"github.com/sha1sum/aws_signing_client"
	"time"
)

type elasticOut struct {
	TimeStamp  time.Time `json:"@timestamp"` // TimeStamp
	LogLevel   string    `json:"level"`      // LogLevel the log level of the log
	Thread     string    `json:"thread"`     // Thread
	LoggerName string    `json:"loggername"` // LoggerName
	ProcessId  string    `json:"processid"`  // ProcessId
	RawLog     string    `json:"rawlog"`     // RawLog
}

func newElasticClient(awsCredentials *credentials.Credentials) (*elastic.Client, error) {
	signer := v4.NewSigner(awsCredentials)
	awsClient, err := aws_signing_client.New(signer, nil, "es", awsRegion)
	if err != nil {
		return nil, err
	}
	return elastic.NewClient(
		elastic.SetURL(awsElasticSearchURL),
		elastic.SetScheme("https"),
		elastic.SetHttpClient(awsClient),
		elastic.SetSniff(false),
	)
}
