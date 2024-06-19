package main

import (
	"fmt"
	// Uncomment this block to pass the first stage
	"net"
	"os"
	"regexp"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	req := make([]byte, 1024)
	conn.Read(req)
	match, _ := regexp.MatchString("GET / HTTP/1.1", string(req))
	match2, _ := regexp.MatchString("^GET /echo/[A-Za-z0-9\\-._~%]+ HTTP/1\\.1", string(req))
	fmt.Println(string(req))
	if match {
		conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Length: 0\r\n\r\n"))
	} else if match2 {
		// Get the path from the request
		path := regexp.MustCompile("^GET /echo/([A-Za-z0-9\\-._~%]+) HTTP/1\\.1").FindStringSubmatch(string(req))[1]
		fmt.Println(path)
		conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: " + fmt.Sprintf("%v", len(path)) + "\r\n\r\n" + path))
	} else {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}
}
