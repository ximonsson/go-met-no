package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"yr"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal(fmt.Sprintf("usage %s <lat> <lon>", os.Args[0]))
	}

	lat, e := strconv.ParseFloat(os.Args[1], 64)
	if e != nil {
		log.Fatal("lat needs to be a valid float")
	}

	lon, e := strconv.ParseFloat(os.Args[2], 64)
	if e != nil {
		log.Fatal("lon needs to be a valid float")
	}

	f, e := yr.Compact(lat, lon)
	if e != nil {
		log.Fatal(e)
	}

	ts := f.Properties.TimeSeries[0].Data
	sum := ts.Next1H.Summary.SymbolCode
	temp := ts.Instant.Details.AirTemp

	fmt.Println(fmt.Sprintf("%.0f:%s", temp, sum))
}
