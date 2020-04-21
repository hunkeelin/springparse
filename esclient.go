package springparse

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/olivere/elastic"
	"github.com/sha1sum/aws_signing_client"
)

var (
	esClient    *elastic.Client
	ctx         context.Context
	bulkRequest *elastic.BulkService
)

func newElasticClient(awsCredentials *credentials.Credentials) (*elastic.Client, error) {
	signer := v4.NewSigner(awsCredentials)
	awsClient, err := aws_signing_client.New(signer, nil, "es", awsRegion)
	if err != nil {
		return nil, fmt.Errorf("Unable to sign client %v", err)
	}
	return elastic.NewClient(
		elastic.SetURL(awsElasticSearchURL),
		elastic.SetScheme("https"),
		elastic.SetHttpClient(awsClient),
		elastic.SetSniff(false),
	)
}
