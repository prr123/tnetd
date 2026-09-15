package main

import (
	"fmt"
	"log"

	"github.com/prr123/telnet-go"
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

	cn := session.GetConn()
	rm := cn.RemoteAddr()
	fmt.Printf("session conn: \n")

	fmt.Printf("Remote: %s\n", rm.String())


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
