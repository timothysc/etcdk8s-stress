package main

import (
	"fmt"
	"log"

	"github.com/coreos/go-etcd/etcd"
)

func stressMem(endpoints []string) {
	client := etcd.NewClient(endpoints)
	data := ""
	for i := 0; i < 128; i++ {
		data += "0"
	}
	total := 0
	for i := 0; i < 1000000; i++ {
		key := fmt.Sprintf("%d", i)
		_, err := client.Set(key, data, 0)
		if err != nil {
			log.Printf("error setting key %s", key)
			continue
		}
		total += 138
		log.Printf("set key %s %dbytes", key, total)
	}
}
