package springparse

import (
	"io/ioutil"
	"testing"
)

func TestLoghash(t *testing.T) {
	s := []byte("{\"log\":\"2020-03-30 19:57:34.061  INFO 3 --- [       Thread-3] b.c.c.NetworkCardTransactionEventHandler : Ignoring ISO8583 network management request with eventId:f9de2711-1206-4aee-b562-48a2a741ce4d\n\",\"stream\":\"stdout\",\"time\":\"2020-03-30T19:57:34.062113419Z\"}")
	g := []byte("{\"log\":\"2020-03-30 19:57:34.061  INFO 3 --- [       Thread-3] b.c.c.NetworkCardTransactionEventHandler : Ignoring ISO8583 network management request with eventId:f9de2711-1206-4aee-b562-48a2a741ce4d\n\",\"stream\":\"stdout\",\"time\":\"2020-03-30T19:57:34.062113429Z\"}")
	if logHash(s) == logHash(g) {
		t.Errorf("loghash is not hashing unique id")
	}
}

func TestParselog(t *testing.T) {
	s := []byte("{\"log\":\"2020-03-30 19:57:34.061  INFO 3 --- [       Thread-3] b.c.c.NetworkCardTransactionEventHandler : Ignoring ISO8583 network management request with eventId:f9de2711-1206-4aee-b562-48a2a741ce4d\n\",\"stream\":\"stdout\",\"time\":\"2020-03-30T19:57:34.062113419Z\"}")
	r := Runner{}
	result, err := r.parseLog(parseLogInput{
		rawLog: s,
	})
	if err != nil {
		t.Errorf(err.Error())
	}
	if result.content.Thread != "Thread-3" || result.content.LogLevel != "INFO" || result.content.LoggerName != "b.c.c.NetworkCardTransactionEventHandler" || result.content.ProcessId != "3" {
		t.Errorf("Parselog failed")
	}
	f, err := ioutil.ReadFile("testlog1")
	if err != nil {
		return
	}
	result, err = r.parseLog(parseLogInput{
		rawLog: f,
	})
	if err != nil {
		t.Errorf(err.Error())
	}
	if result.content.Thread != "Thread-3" || result.content.LogLevel != "INFO" || result.content.LoggerName != "c.v.b.t.LeeroyWebhookEventConsumer" || result.content.ProcessId != "1" {
		t.Errorf("Parselog failed")
	}
}
