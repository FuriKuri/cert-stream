package main

import (
	"github.com/CaliDog/certstream-go"
	logging "github.com/op/go-logging"
	"os"
	"strings"
)

var log = logging.MustGetLogger("example")

func main() {
	domain := os.Args[1]
	stream, errStream := certstream.CertStreamEventStream(false)
	for {
		select {
		case jq := <-stream:
			messageType, err := jq.String("message_type")
			domains, err := jq.String("data", "leaf_cert", "all_domains", "0")

			if err == nil && strings.HasSuffix(domains, domain) {
				log.Info("Message type -> ", messageType, "Domains: -> ", domains)
			}

		case err := <-errStream:
			log.Error(err)
		}
	}
}
