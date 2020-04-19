## Configurations
Configurations are set during startup. 

## Environment Variables
* `AWS_ACCESS_KEY_ID`
* `AWS_SECRET_ACCESS_KEY`
* `AWS_SESSION_TOKEN`
* `AWS_REGION`: 
* `AWS_ELASTICSEARCH_URL`: URL for the elasticsearch domain. (required)
* `HOST_PORT`: The default port the server is hosted on. Default `8080`
* `LOG_PREFIX`: The log prefix when indexing to elasticsearch. (required)
* `LOG_DIRECTORY`: The logs directory the software listens to. (required) If you mount the volumne to `/var/log/` it would be `/var/log/containers`.
* `SERVICE_REGEX`: The regex of the log you want `springparse` to listen to. E.g `foo&bar, aaa&bbb` means any log file that matches the regex `*foo*bar*` or `*aaa*bbb*`. (required)
* `BATCH_COUNT`: The limit of batch requests before its send to elasticsearch
* `FLUSH_CYCLE`: The time limit before it flushes the request to elasticsearch


## prometheus metrics
* `springparse_putsuccess`: The total number batch send due to buffer overload
* `springparse_putflushsuccess`: The total of batch sent due to time flushes
* `springparse_putfailed`: The number of failed puts to elasticsearch index
