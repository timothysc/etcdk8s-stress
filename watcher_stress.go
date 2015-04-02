package main

import (
	"fmt"
	"log"

	"github.com/coreos/go-etcd/etcd"
)

func stressWatch(endpoints []string, c int, prefix string) {
	for i := 0; i < c; i++ {
		client := etcd.NewClient(endpoints)
		go func(i int) {
			ch := make(chan *etcd.Response, 8)
			key := fmt.Sprintf("%s/%v", prefix, i)
			go client.Watch(key, 0, true, ch, nil)
			log.Printf("watching on %s", key)
			for {
				resp := <-ch
				log.Printf("recv: %s, %s", resp.Node.Key, resp.Node.Value)
			}
		}(i)
	}
}
