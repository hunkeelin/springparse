package springparse

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"os"
	"strings"
)

var (
	awsCredentials      = credentials.NewStaticCredentials(awsAccessKeyId, awsSecretAccessKey, awsSessionToken)
	serviceRegexList    = strings.Split(serviceRegex, ",")
	awsAccessKeyId      = os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKey  = os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsSessionToken     = os.Getenv("AWS_SESSION_TOKEN")
	awsElasticSearchURL = os.Getenv("AWS_ELASTICSEARCH_URL")
	logPrefix           = os.Getenv("LOGSTASH_PREFIX")
	awsbucketPrefix     = os.Getenv("S3_PREFIX")
	awsbucketName       = os.Getenv("AWS_S3_BUCKET")
	tailBinary          = os.Getenv("TAIL_BIN")
	logDirectory        = os.Getenv("LOG_DIRECTORY")
	serviceRegex        = os.Getenv("SERVICE_REGEX")
)