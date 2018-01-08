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

        if err != nil{
          log.Fatal("Error decoding jq string")
        }

        log.Info("Message type -> ", messageType)
        log.Info("recv: ", jq)
      
      case err := <-errStream:
        log.Error(err)
    }
  }
}