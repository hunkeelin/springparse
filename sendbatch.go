package springparse

import (
	"context"
	"github.com/olivere/elastic"
)

var (
	stopSig   chan bool
	sendItems chan elasticItem
)

type elasticItem struct {
	index    string
	id       string
	bodyJson elasticOut
}

func sendBatch(item <-chan elasticItem, stopSignal <-chan bool) {
	var tosend []elasticItem
	for {
		select {
		case i := <-item:
			tosend = append(tosend, i)
			if len(tosend) >= batchCountInt {
				esClient, err := newElasticClient(awsCredentials)
				if err != nil {
					// halting the program if unable to create es client
					panic(err)
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
					panic(err)
				}
				putSuccess.Inc()
				// Clear up the array
				tosend = nil
			}
		case <-stopSignal:
			return
		}
	}
}
