package springparse

import (
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
	if len(tosend) < 1 {
		// nothing to send
		return nil
	}
	for _, i := range tosend {
		tmpRequest := elastic.NewBulkIndexRequest().
			Index(i.index).
			Type(elasticType).
			Id(i.id).
			Doc(i.bodyJson)
		bulkRequest = bulkRequest.Add(tmpRequest)
	}
	// Successfully populate bulk request now sending it to elasticSearch
	bulkDo, err := bulkRequest.Do(ctx)
	if err != nil {
		return err
	}
	log.Info(fmt.Sprintf("Sending batch, this should usually be %v apart, unless current length %v == %v. %v got indexed", flushCycleInt, len(tosend), batchCountInt, len(bulkDo.Items)))
	return nil
	// Clear up the array
}
