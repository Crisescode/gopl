package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	for _, args := range os.Args[1:] {
		address := strings.Split(args, "=")[1]
		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go mustCopy(os.Stdout, conn)
	}

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err !=nil {
		log.Fatal(err)
	}
}