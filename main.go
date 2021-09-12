package main

import (
	"context"
	"fmt"
	"os"

	"github.com/lukibahr/windhager-biowin2t-client-go/biowin2t"
	log "github.com/sirupsen/logrus"
)

func main() {

	ctx := context.Background()

	url, isSet := os.LookupEnv("WH_ENDPOINT")
	if !isSet {
		log.Panicf("WH_ENDPOINT not set, exiting")
	}
	username, isSet := os.LookupEnv("WH_USERNAME")
	if !isSet {
		log.Panicf("WH_USERNAME not set, exiting")
	}
	password, isSet := os.LookupEnv("WH_PASSWORD")
	if !isSet {
		log.Panicf("WH_PASSWORD not set, exiting")
	}
	c := biowin2t.NewWindhagerClient(url,
		username, password)
	result, err := c.GetTimeUntilNextMajorMaintenanceInHours(ctx)
	if err != nil {
		log.Panicf("Exiting: %s", err)
	}
	fmt.Print(result)

}
