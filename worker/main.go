package main

import (
	"log"
	"net/http"

	"github.com/babafemikuku/temporal-ip-geolocation/iplocate"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	client, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer client.Close()

	w := worker.New(client, iplocate.TaskQueueName, worker.Options{})

	activities := &iplocate.IPActivities{
		HTTPClient: http.DefaultClient,
	}

	w.RegisterWorkflow(iplocate.GetAddressFromIp)
	w.RegisterActivity(activities)

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
