package springparse

import (
	"context"
	"fmt"
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
				log.Info(fmt.Sprintf("Flushing %v records via buffer limit", len(tosend)))
				err := batchSendDo(tosend)
				if err != nil {
					panic(err)
				}
				putSuccess.Inc()
				tosend = nil
			} else {
				log.Info("Not flushing because buffer is nil")
			}
		case <-flushSignal:
			if tosend != nil {
				log.Info(fmt.Sprintf("Flushing %v records via time limit", len(tosend)))
				err := batchSendDo(tosend)
				if err != nil {
					panic(err)
				}
				putFlushSuccess.Inc()
				tosend = nil
			} else {
				log.Info("Not flushing because buffer is nil time")
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
	_, err = bulkRequest.Do(ctx)
	if err != nil {
		return err
	}
	log.Info(fmt.Sprintf("Successfully send %v records to elasticsearch", len(tosend)))
	// purely for debug
	if len(tosend) == 1 {
		log.Info("This is the id with length 1 ", tosend[0].id)
	}
	return nil
	// Clear up the array
}
