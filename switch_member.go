package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/coreos/go-etcd/etcd"
)

func stressSwitch(endpoints []string) {
	client := etcd.NewClient(endpoints)
	etcd.SetLogger(log.New(os.Stderr, "go-etcd", log.LstdFlags))
	client.OpenCURL()
	for i := 0; i < 10000; i++ {
		key := fmt.Sprintf("%d", i)
		_, err := client.Set(key, key, 0)
		if err != nil {
			log.Printf("error setting key %s", key)
			time.Sleep(time.Second)
			continue
		}
		time.Sleep(time.Second)
	}
}
