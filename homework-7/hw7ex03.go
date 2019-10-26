package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	close := make(chan int)
	ex := make([]byte, 4)

	go func() {
		for {
			os.Stdin.Read(ex)
			if string(ex) == "exit" {
				break
			}
		}
		close <- 1
	}()

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Print(err)
				continue
			}
			go handleConn(conn)
		}
	}()

	<-close
	listener.Close()
	fmt.Println("Server has been closed!")
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}

}
