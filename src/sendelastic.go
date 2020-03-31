package springparse

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

func sendElasticSearch(s []byte) error {
	out, err := parseLog(s)
	if err != nil {
		return err
	}
	rDate := fmt.Sprintf(time.Now().UTC().Format("2006-01-02"))
	client, err := newElasticClient(awsCredentials)
	if err != nil {
		return err
	}
	ctx := context.Background()
	put, err := client.Index().
		Index("springparse" + "-" + rDate).
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
