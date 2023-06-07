package main

import (
	communication "TopLearn/communication/proto"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"

	"github.com/golang/protobuf/proto"
)

func main() {
	op := flag.String("op", "s", "s for Server and c for Client.")
	flag.Parse()

	switch strings.ToLower(*op) {
	case "s":
		RunServer()
	case "c":
		RunClient()
	}

}

func RunServer() {
	l, err := net.Listen("tcp", ":8283")
	if err != nil {
		log.Fatal(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		defer l.Close()
		go func(cn net.Conn) {
			defer c.Close()
			data, err := ioutil.ReadAll(c)
			if err != nil {
				log.Fatal(err)
			}
			p := &communication.Person{}
			proto.Unmarshal(data, p)

			fmt.Println(p)
		}(c)
	}

}

//RunClient Start ClientType Application
func RunClient() {
	p := &communication.Person{
		Id:     proto.Int(1),
		Name:   proto.String("Mohammad"),
		Family: proto.String("Ghari"),
	}
	data, err := proto.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	SendData(data)

	p = &communication.Person{
		Id:     proto.Int(2),
		Name:   proto.String("Ali"),
		Family: proto.String("Ghari"),
	}
	data, err = proto.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	SendData(data)

}

func SendData(data []byte) {
	c, err := net.Dial("tcp", "127.0.0.1:8283")

	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	c.Write(data)

}
