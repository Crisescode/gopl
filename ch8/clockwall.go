package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	done := make(chan struct{})
	for _, args := range os.Args[1:] {
		address := strings.Split(args, "=")[1]
		fmt.Println("address:", address)
		go connTCP(address, done)
	}

	//for {
	//	time.Sleep(10 * time.Second)
	//}

	<- done

}

func connTCP(address string, ch chan struct{}) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	ch <- struct {
	}{}
	defer conn.Close()
	distMustCopy(os.Stdout, conn)
}

func distMustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err !=nil {
		log.Fatal(err)
	}
}