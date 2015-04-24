package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

func main() {
	c := flag.Int("c", 100, "concurrent client")
	action := flag.String("action", "", "stress action")
	prefix := flag.String("prefix", "/stress", "stress prefix")
	endpoint := flag.String("end-points", "http://host01-rack10:2379,http://host02-rack10:2379,http://host17-rack11:2379", "end-point string")
	interval := flag.Duration("interval", time.Second, "write interval")
	flag.Parse()

	endpoints := strings.Split(*endpoint, ",")
	switch *action {
	case "watch":
		stressWatch(endpoints, *c, *prefix)
	case "write":
		fmt.Println("starting write-stress")
		stressWrite(endpoints, *c, *prefix, *interval)
	case "mem":
		stressMem(endpoints)
	case "switch":
		stressSwitch(endpoints)
	}
	select {}
}
