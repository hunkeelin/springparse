package springparse

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/url"
	"strconv"
)

// ValidateAwsEnv validates all aws related configration and set defaults from environment variables
func (r *Client) ValidateAwsEnv() error {
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
func (r *Client) ValidateOtherEnv() error {
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
	if elasticType == "" {
		elasticType = "springparse"
	}
	if flushCycle == "" {
		flushCycleInt = 65
	} else {
		var err error
		flushCycleInt, err = strconv.Atoi(flushCycle)
		if err != nil {
			return err
		}
	}
	if batchCount == "" {
		batchCountInt = 250
	} else {
		var err error
		batchCountInt, err = strconv.Atoi(batchCount)
		if err != nil {
			return err
		}
	}
	return nil
}

// ShowConfiguration shows the environment variables being set
func (r *Client) ShowConfiguration() {
	log.Info("Starting springparse with the following configuration")
	log.Info(fmt.Sprintf("Elasticsearch URL: %v", awsElasticSearchURL))
	log.Info(fmt.Sprintf("Log prefix: %v", logPrefix))
	log.Info(fmt.Sprintf("Tailing directory:  %v", logDirectory))
	log.Info(fmt.Sprintf("Service regex: %v", serviceRegex))
	log.Info(fmt.Sprintf("Aws Region: %v", awsRegion))
	log.Info(fmt.Sprintf("Batch count: %v", batchCountInt))
	log.Info(fmt.Sprintf("Flush interval: %v", flushCycleInt))
}
