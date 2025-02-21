package monitor

import (
	"time"

	"github.com/kehrlann/gonitor/config"
)

// Monitor takes a resource and polls the given HTTP url for errors , and emits failure / recovery messages accordingly
func Monitor(resources []config.Resource) <-chan *StateChangeMessage {
	messages := make(chan *StateChangeMessage)
	for _, resource := range resources {
		go monitor(resource, messages)
	}
	return messages
}

func monitor(resource config.Resource, messages chan<- *StateChangeMessage) {
	responseCodes := make(chan int)

	go analyze(resource, responseCodes, messages)

	// TODO : do we want to test that ?? if so, we need to mock the duration interval (val'd ?) and the fetch method
	for range time.Tick(time.Duration(resource.IntervalInSeconds) * time.Second) {
		responseCodes <- Fetch(resource)
	}
}
