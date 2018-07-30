package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/sdeoras/scratchpad/proto/mesg"
	"k8s.io/api/core/v1"
)

func main() {
	data := new(mesg.Data)
	data.Pod = new(v1.Pod)
	data.Pod.Name = "myPod"

	b, err := proto.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	d := new(mesg.Data)
	if err := proto.Unmarshal(b, d); err != nil {
		log.Fatal(err)
	}

	fmt.Println(d.Pod.Name)
}
