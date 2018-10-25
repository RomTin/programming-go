package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	tzs := os.Args[1:]

	for _, tz := range tzs {
		r := strings.Split(tz, "=")
		tzName := r[0]
		target := r[1]
		conn, err := net.Dial("tcp", target)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go mustCopy(os.Stdout, conn, tzName)
	}
	for {
		time.Sleep(time.Minute)
	}
}

func mustCopy(dst io.Writer, src io.Reader, tzName string) {
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		fmt.Fprintf(dst, "%s %s\n", tzName, scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Println("error")
	}
}
