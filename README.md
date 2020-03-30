# RedisProxy 
[![CircleCI](https://circleci.com/gh/hunkeelin/springparse.svg?style=shield)](https://circleci.com/gh/hunkeelin/springparse)
[![Go Report Card](https://goreportcard.com/badge/github.com/hunkeelin/springparse)](https://goreportcard.com/report/github.com/hunkeelin/springparse)
[![GoDoc](https://godoc.org/github.com/hunkeelin/springparse/src?status.svg)](https://godoc.org/github.com/hunkeelin/springparse/src)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hunkeelin/springparse/master/LICENSE)


## Motivation
As of now, we can use fluentd daemonset to have a pod on every eks node. Mount the `/var/log/containers` and tail any incoming logs and send it to elasticsearch. However, the log format isn't ideal. For example, the below log will simply stick the sprintboot logs to the key value `log`. The log level, executor, thread, and loggername are not searchable in elasticsearch that way. `fluentd` have log transformation via `regexp` but that plugin is not suffice for the use case as regexing logs for transformation is too difficult or simply not possible. 
```
{"log":"2020-03-30 19:57:33.702  INFO 1 --- [       Thread-3] c.v.b.t.fooWebhookEventConsumer      : FOOLOG","stream":"stdout","time":"2020-03-30T19:57:33.703403463Z"}
```

## Golang version

`springparse` is currently compatible with golang version from 1.12+.

