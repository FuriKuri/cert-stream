package main

import (
	"github.com/CaliDog/certstream-go"
	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("example")

func main() {
	// The false flag specifies that we don't want heartbeat messages.
	stream, errStream := certstream.CertStreamEventStream(false)
	for {
		select {
		case jq := <-stream:
			messageType, err := jq.String("message_type")
			domains, err := jq.String("data", "leaf_cert", "all_domains", "0")

			if err == nil {
				log.Info("Message type -> ", messageType, "Domains: -> ", domains)
			}

		case err := <-errStream:
			log.Error(err)
		}
	}
}
