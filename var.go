package springparse

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"os"
	"strings"
)

var (
	awsCredentials      = credentials.NewStaticCredentials(awsAccessKeyId, awsSecretAccessKey, awsSessionToken)
	serviceRegexList    = strings.Split(serviceRegex, ",")
	batchCountInt       int
	awsAccessKeyId      = os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKey  = os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsSessionToken     = os.Getenv("AWS_SESSION_TOKEN")
	awsElasticSearchURL = os.Getenv("AWS_ELASTICSEARCH_URL")
	logPrefix           = os.Getenv("LOG_PREFIX")
	awsbucketPrefix     = os.Getenv("S3_PREFIX")
	hostPort            = os.Getenv("HOST_PORT")
	awsbucketName       = os.Getenv("AWS_S3_BUCKET")
	logDirectory        = os.Getenv("LOG_DIRECTORY")
	serviceRegex        = os.Getenv("SERVICE_REGEX")
	awsRegion           = os.Getenv("AWS_REGION")
	batchCount          = os.Getenv("BATCH_COUNT")
	elasticType         = os.Getenv("ELASTIC_TYPE")
)
