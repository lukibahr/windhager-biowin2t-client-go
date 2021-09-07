package main

import (
	"context"
	"fmt"

	//"net/http"
	//"github.com/prometheus/client_golang/prometheus"
	//"github.com/prometheus/client_golang/prometheus/promhttp"
	windhager "github.com/lukibahr/windhager-biowin2t-client-go/biowin2t"
	log "github.com/sirupsen/logrus"
)

func main() {

	ctx := context.Background()

	c := windhager.NewWindhagerClient("user", "password")
	result, err := c.GetTimeUntilNextMajorMaintenanceInHours(ctx)
	if err != nil {
		log.Panicf("Exiting: %s", err)
	}
	fmt.Printf("%s \n", result.Value)
	result, err1 := c.GetTimeUntilNextMaintenanceInHours(ctx)
	if err1 != nil {
		log.Panicf("Exiting: %s", err)
	}
	fmt.Print("\n")
	result, err2 := c.GetCountOfBurningUnit(ctx)
	if err2 != nil {
		log.Panicf("Exiting: %s", err)
	}
	fmt.Printf("%s\n", result.Value)
	result, err3 := c.GetTotalOperationalRuntime(ctx)
	if err3 != nil {
		log.Panicf("Exiting: %s", err)
	}
	//log.Printf("result: %s", result.Value)

	//Create a new instance of the foocollector and
	//register it with the prometheus client.
	// coll := newWindhagerCollector()
	// prometheus.MustRegister(coll)

	// //This section will start the HTTP server and expose
	// //any metrics on the /metrics endpoint.
	// http.Handle("/metrics", promhttp.Handler())
	// log.Info("Beginning to serve on port :8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
