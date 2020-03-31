package springparse

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/url"
)

// ValidateAwsEnv validates all aws related configration and set defaults from environment variables
func ValidateAwsEnv() error {
	if awsElasticSearchURL == "" {
		return fmt.Errorf("Please set AWS_ELASTICSEARCH_URL")
	}
	_, err := url.Parse(awsElasticSearchURL)
	if err != nil {
		return err
	}
	if awsRegion == "" {
		awsRegion = "us-west-2"
	}
	return nil
}

// ValidateOtherEnv validates all system info from environment variables.
func ValidateOtherEnv() error {
	if tailBinary == "" {
		return fmt.Errorf("Please set TAIL_BIN")
	}
	if logPrefix == "" {
		return fmt.Errorf("Please specify LOG_PREFIX for elasticsearch prefix")
	}
	if serviceRegex == "" {
		return fmt.Errorf("Please specify SERVICE_REGEX")
	}
	if logDirectory == "" {
		return fmt.Errorf("Please specify LOG_DIRECTORY")
	}
	if hostPort == "" {
		hostPort = "8080"
	}
	return nil
}

// ShowConfiguration shows the environment variables being set
func ShowConfiguration() {
	log.Info("Starting springparse with the following configuration")
	log.Info(fmt.Sprintf("AWS_ELASTICSEARCH_URL: %v", awsElasticSearchURL))
	log.Info(fmt.Sprintf("LOG_PREFIX: %v", logPrefix))
	log.Info(fmt.Sprintf("TAIL_BIN: %v", tailBinary))
	log.Info(fmt.Sprintf("LOG_DIRECTORY:  %v", logDirectory))
	log.Info(fmt.Sprintf("SERVICE_REGEX: %v", serviceRegex))
	log.Info(fmt.Sprintf("AWS_REGION: %v", awsRegion))
}
