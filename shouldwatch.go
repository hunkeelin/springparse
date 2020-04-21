package springparse

import (
	"strings"
)

type shouldWatchInput struct {
	logFile string
}
type shouldWatchOutput struct {
	watch bool
	err   error
}

func (r *Client) shouldWatch(s shouldWatchInput) shouldWatchOutput {
	for _, service := range serviceRegexList {
		serviceDetail := strings.Split(service, "&")
		for _, detail := range serviceDetail {
			if strings.Contains(detail, "!") {
				if strings.Contains(s.logFile, detail[1:]) {
					return shouldWatchOutput{
						watch: false,
					}
				}
			}
		}
	}
	for _, service := range serviceRegexList {
		serviceDetail := strings.Split(service, "&")
		var didmatch bool
		for _, detail := range serviceDetail {
			match := strings.Contains(s.logFile, detail)
			if !match {
				didmatch = false
				break
			}
			didmatch = true
		}
		if didmatch {
			return shouldWatchOutput{
				watch: true,
				err:   nil,
			}
		}
	}
	return shouldWatchOutput{
		watch: false,
	}
}
