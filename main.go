package main

import (
	"flag"
	"log"
	"time"

	"strings"

	"github.com/influxdata/influxdb/client/v2"
)

func main() {
	//	Our command line options:
	influxServer := flag.String("server", "http://localhost:8086", "Full path to the influxdb server, including credentials")
	influxDatabase := flag.String("database", "events", "InfluxDB database to write the annotation to")
	measurementName := flag.String("name", "build", "Measurement name")
	measurementValue := flag.String("value", "your_value_here", "Measurement value")
	measurementTags := flag.String("tags", "", "Comma seperated list of key=value tags")
	flag.Parse()

	//	Create a new InfluxDB client
	c, err := client.NewHTTPClient(client.HTTPConfig{Addr: *influxServer})
	if err != nil {
		log.Fatal(err)
	}

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  *influxDatabase,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	//	Update our tags
	tags := make(map[string]string)
	kvPairs := strings.Split(*measurementTags, ",")

	for _, kvPair := range kvPairs {
		nameValue := strings.Split(kvPair, "=")
		tags[nameValue[0]] = nameValue[1]
	}

	//	Update our measurement
	measurements := map[string]interface{}{"value": *measurementValue}

	// Create a point and add to batch
	pt, err := client.NewPoint(*measurementName, tags, measurements, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	// Write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}
}
