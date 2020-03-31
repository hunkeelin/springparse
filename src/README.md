## Configurations
Configurations are set during startup. 

## Environment Variables
* `AWS_ACCESS_KEY_ID`
* `AWS_SECRET_ACCESS_KEY`
* `AWS_SESSION_TOKEN`
* `AWS_REGION`: 
* `AWS_ELASTICSEARCH_URL`: URL for the elasticsearch domain. (required)
* `LOG_PREFIX`: The log prefix when indexing to elasticsearch. (required)
* `TAIL_BIN`: The binary for `tail`. (required)
* `LOG_DIRECTORY`: The logs directory the software listens to. (required) If you mount the volumne to `/var/log/` it would be `/var/log/containers`.
* `SERVICE_REGEX`: The regex of the log you want `springparse` to listen to. E.g `foo&bar, aaa&bbb` means any log file that matches the regex `*foo*bar*` or `*aaa*bbb*`. (required)
