package main

import (
	"context"
	"flag"
	"log"
	"time"
)

func main() {
	id := flag.Uint64("id", 1, "node id")
	flag.Parse()
	log.Printf("I'am node %v\n", *id)

	cluster := map[uint64]string{
		1: "http://127.0.0.1:22210",
		2: "http://127.0.0.1:22220",
		3: "http://127.0.0.1:22230",
	}

	log.Println("id is ", *id)
	n := newRaftNode(*id, cluster)

	if *id == 1 {
		time.Sleep(5 * time.Second)
		for {
			log.Printf("Propose on node %v\n", *id)
			n.node.Propose(context.TODO(), []byte("hello"))
			time.Sleep(time.Second)
		}

	}

	select {}

}
