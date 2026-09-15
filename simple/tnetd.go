package main

import (
//	"context"
//	"io"
	"log"

	"github.com/globalcyberalliance/telnet-go"
)

func main() {
	if err := telnet.ListenAndServe("10.0.14.1:2323", YourHandlerFunc); err != nil {
		log.Fatalf("error -- listen: 5v\n", err)
	}
}

func YourHandlerFunc(session *telnet.Session) {
	if err := session.WriteLine("Welcome!\r\n"); err != nil {
		return
	}

	for {
		line, err := session.ReadLine()
		if err != nil {
			return
		}

		if len(line) == 0 {
			if err = session.WriteLine("Goodbye!\r\n"); err != nil {
				return
			}
			return
		}

		if err = session.WriteLine("You wrote: "+line+"\r\n"); err != nil {
			return
		}
	}
}
