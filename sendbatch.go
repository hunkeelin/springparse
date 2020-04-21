package springparse

import (
	"context"
	"github.com/olivere/elastic"
	log "github.com/sirupsen/logrus"
)

var (
	flushSig  chan bool
	sendItems chan elasticItem
)

type elasticItem struct {
	index    string
	id       string
	bodyJson elasticOut
}

func sendBatch(item <-chan elasticItem, flushSignal <-chan bool) {
	var tosend []elasticItem
	for {
		select {
		case i := <-item:
			tosend = append(tosend, i)
			if len(tosend) >= batchCountInt {
				err := batchSendDo(tosend)
				if err != nil {
					panic(err)
				}
				putSuccess.Inc()
				tosend = nil
			}
		case <-flushSignal:
			if tosend != nil {
				err := batchSendDo(tosend)
				if err != nil {
					panic(err)
				}
				putFlushSuccess.Inc()
				tosend = nil
			}
		}
	}
}

func batchSendDo(tosend []elasticItem) error {
	esClient, err := newElasticClient(awsCredentials)
	if err != nil {
		return err
	}
	ctx := context.Background()
	bulkRequest := esClient.Bulk()
	for _, i := range tosend {
		tmpRequest := elastic.NewBulkIndexRequest().
			Index(i.index).
			Type(elasticType).
			Id(i.id).
			Doc(i.bodyJson)
		bulkRequest = bulkRequest.Add(tmpRequest)
	}
	// Successfully populate bulk request now sending it to elasticSearch
	log.Info("Sending batch this should usually be time cycle(45 seconds) apart unless length is at limit ", len(tosend))
	_, err = bulkRequest.Do(ctx)
	if err != nil {
		return err
	}
	if len(tosend) == 1 {
		log.Debug("This is the id with length 1 investigate ", tosend[0].id)
	}
	return nil
	// Clear up the array
}
