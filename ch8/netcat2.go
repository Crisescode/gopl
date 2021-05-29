package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy2(os.Stdout, conn)
}

func mustCopy2(dst io.Writer, src io.Reader) {
	fmt.Println("src:", &src)
	if _, err := io.Copy(dst, src); err !=nil {
		log.Fatal(err)
	}
}