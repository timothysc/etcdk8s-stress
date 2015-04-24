package main

import (
	"fmt"
	"log"
	"time"

	"github.com/coreos/go-etcd/etcd"
)

// copying the exact k8's flow
type EtcdGetSet interface {
	GetCluster() []string
	Get(key string, sort, recursive bool) (*etcd.Response, error)
	Set(key, value string, ttl uint64) (*etcd.Response, error)
	Create(key, value string, ttl uint64) (*etcd.Response, error)
	Delete(key string, recursive bool) (*etcd.Response, error)
	CompareAndSwap(key, value string, ttl uint64, prevValue string, prevIndex uint64) (*etcd.Response, error)
	Watch(prefix string, waitIndex uint64, recursive bool, receiver chan *etcd.Response, stop chan bool) (*etcd.Response, error)
}

func writer(client EtcdGetSet, prefix string, interval time.Duration, c int, i int) {
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
}

func stressWrite(endpoints []string, c int, prefix string, interval time.Duration) {
	var client EtcdGetSet
	client = etcd.NewClient(endpoints)
	for i := 0; i < c; i++ {
		go writer(client, prefix, interval, c, i)
		/*go func(i int) {
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
		}(i)*/
	}
}
