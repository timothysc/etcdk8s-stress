package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

func main() {
	c := flag.Int("c", 10, "concurrent client")
	action := flag.String("action", "", "stress action")
	prefix := flag.String("prefix", "/stress", "stress prefix")
	endpoint := flag.String("end-points", "http://127.0.0.1:4001", "end-point string")
	interval := flag.Duration("interval", time.Second, "write interval")
	flag.Parse()

	endpoints := strings.Split(*endpoint, ",")
	switch *action {
	case "watch":
		stressWatch(endpoints, *c, *prefix)
	case "write":
		fmt.Println("starting write-stress")
		stressWrite(endpoints, *c, *prefix, *interval)
	}
	select {}
}
