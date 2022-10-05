package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %s", err.Error())
	}
}

func main() {
	// ./bytestream host:port number_of_bytes
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s {host:port} {number_of_bytes}", os.Args[0])
	}
	service := os.Args[1]
	nbytes, err := strconv.Atoi(os.Args[2])
	checkError(err)

	conn, err := net.Dial("tcp", service)
	checkError(err)

	// generate a byte slice filling it with random bytes and send it over the connection number_of_bytes times
	buf := make([]byte, nbytes)

	rand.Read(buf)

	// send bytes one by one over the connection
	for i := 0; i < nbytes; i++ {
		// send one byte
		// log bytes sent
		// sleep for 1 second
		fmt.Printf("Sending byte %d: %d \n", i, buf[i])
		conn.Write(buf[i : i+1])
		time.Sleep(1 * time.Second)
	}

}
