package main

import (
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
	// ./bytestream {host:port} {number_of_bytes} [delay] [seed] [send_length]
	// delay is in milliseconds (default 100) optional
	// seed is for the random number generator (default 0) optional
	// send_length is the number of bytes to send at a time (default 1) optional
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s host:port number_of_bytes [delay] [seed] [send_length]", os.Args[0])
	}

	service := os.Args[1]
	nbytes, err := strconv.Atoi(os.Args[2])
	checkError(err)

	//if nbytes is not 1, then round up to the nearest 100 bytes
	if nbytes != 1 {
		nbytes = (nbytes/100 + 1) * 100
	}

	conn, err := net.Dial("tcp", service)
	checkError(err)

	// generate a byte slice filling it with random bytes and send it over the connection number_of_bytes times
	buf := make([]byte, nbytes)

	rand.Read(buf)

	delay := 100
	seed := 0
	send_length := 1
	if len(os.Args) > 3 {
		delay, err = strconv.Atoi(os.Args[3])
		checkError(err)
		seed, err = strconv.Atoi(os.Args[4])
		checkError(err)
		send_length, err = strconv.Atoi(os.Args[5])
		checkError(err)
	}

	rand.Seed(int64(seed))

	for i := 0; i < nbytes; i += send_length {
		conn.Write(buf[i : i+send_length])
		time.Sleep(time.Duration(delay) * time.Millisecond)
		log.Printf("Sent data: %d", buf[i:i+send_length])
	}

}
