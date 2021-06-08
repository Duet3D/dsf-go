package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Duet3D/dsf-go/dsf/commands"
	"github.com/Duet3D/dsf-go/dsf/connection"
	"github.com/Duet3D/dsf-go/dsf/connection/initmessages"
	"github.com/Duet3D/dsf-go/dsf/types"
)

func usage(name string) {
	fmt.Printf("usage: %s SOCKET_FILE subscribe|intercept|command GCODE...\n", name)
}

func main() {
	if len(os.Args) <= 2 {
		usage(os.Args[0])
		os.Exit(1)
	}

	var socket_file string = os.Args[1]

	switch os.Args[2] {
	case "subscribe":
		subscribe(socket_file)
	case "intercept":
		intercept(socket_file)
	case "command":
		if len(os.Args) > 2 {
			for _, c := range os.Args[2:] {
				command(socket_file, c)
			}
		} else {
			command(socket_file, "")
		}
	default:
		usage(os.Args[0])
		os.Exit(1)
	}
}

func command(socket_file string, code string) {
	cc := connection.CommandConnection{}
	err := cc.Connect(socket_file)
	if err != nil {
		panic(err)
	}
	defer cc.Close()
	if code != "" {
		r, err := cc.PerformSimpleCode(code, types.SPI)
		if err != nil {
			log.Panic(err)
		}
		log.Println(r)
	} else {
		mm, err := cc.GetSerializedMachineModel()
		if err != nil {
			log.Panic(err)
		}
		log.Println(string(mm))
	}
}

func subscribe(socket_file string) {
	sc := connection.SubscribeConnection{}
	sc.Debug = true
	err := sc.Connect(initmessages.SubscriptionModePatch, "heat/**", socket_file)
	if err != nil {
		log.Panic(err)
	}
	defer sc.Close()
	m, err := sc.GetMachineModelPatch()
	if err != nil {
		log.Panic(err)
	}
	log.Println(m)
}

func intercept(socket_file string) {
	ic := connection.InterceptConnection{}
	ic.Debug = true
	err := ic.Connect(initmessages.InterceptionModePre, socket_file)
	if err != nil {
		log.Panic(err)
	}
	defer ic.Close()
	for {
		c, err := ic.ReceiveCode()
		if err != nil {
			log.Panic(err)
		}
		cc := c.Clone()
		cc.Flags |= commands.Asynchronous
		ic.PerformCode(cc)
		// log.Println(c)
		err = ic.IgnoreCode()
		if err != nil {
			log.Panic(err)
		}
	}
}
