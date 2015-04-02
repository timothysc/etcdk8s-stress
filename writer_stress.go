package main

import (
	"fmt"
	"log"
	"time"

	"github.com/coreos/go-etcd/etcd"
)

func stressWrite(endpoints []string, c int, prefix string, interval time.Duration) {
	for i := 0; i < c; i++ {
		client := etcd.NewClient(endpoints)
		go func(i int) {
			time.Sleep(time.Duration(int(interval) / c * i))
			for {
				key := fmt.Sprintf("%s/%v", prefix, i)
				_, err := client.Set(key, key, 0)
				if err != nil {
					log.Printf("error setting key %s", key)
					continue
				}
				log.Printf("set key %s", key)
				time.Sleep(interval)
			}
		}(i)
	}
}
