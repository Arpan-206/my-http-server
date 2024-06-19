package main

import (
	"fmt"
	"strings"

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

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	req := make([]byte, 1024)
	conn.Read(req)
	match, _ := regexp.MatchString("GET / HTTP/1.1", string(req))
	match2, _ := regexp.MatchString("^GET /echo/[A-Za-z0-9\\-._~%]+ HTTP/1\\.1", string(req))
	match3, _ := regexp.MatchString("^GET /user-agent HTTP/1\\.1", string(req))
	match4, _ := regexp.MatchString("^GET /files/[A-Za-z0-9\\-._~%]+ HTTP/1\\.1", string(req))
	match5, _ := regexp.MatchString("^POST /files/[A-Za-z0-9\\-._~%]+ HTTP/1\\.1", string(req))
	fmt.Println(string(req))
	if match {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else if match2 {
		// Get the path from the request
		path := regexp.MustCompile("^GET /echo/([A-Za-z0-9\\-._~%]+) HTTP/1\\.1").FindStringSubmatch(string(req))[1]
		fmt.Println(path)
		conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: " + fmt.Sprintf("%v", len(path)) + "\r\n\r\n" + path))
	} else if match3 {
		user_agent := regexp.MustCompile("User-Agent: (.*)").FindStringSubmatch(string(req))[1]
		user_agent = strings.Trim(user_agent, "\r\n")
		conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: " + fmt.Sprintf("%v", len(user_agent)) + "\r\n\r\n" + user_agent))
	} else if match4 {
		path := regexp.MustCompile("^GET /files/([A-Za-z0-9\\-._~%]+) HTTP/1\\.1").FindStringSubmatch(string(req))[1]
		dir := os.Args[2]
		data, err := os.ReadFile(dir + path)
		if err != nil {
			conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		} else {
			conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Type: application/octet-stream\r\nContent-Length: " + fmt.Sprintf("%v", len(data)) + "\r\n\r\n" + string(data)))
		}
	} else if match5 {
		path := regexp.MustCompile("^POST /files/([A-Za-z0-9\\-._~%]+) HTTP/1\\.1").FindStringSubmatch(string(req))[1]
		dir := os.Args[2]
		body := regexp.MustCompile("\r\n\r\n(.*)").FindStringSubmatch(string(req))[1]
		data := body
		data = strings.Trim(data, "\x00")
		err := os.WriteFile(dir+path, []byte(data), 0666)
		if err != nil {
			conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		} else {
			conn.Write([]byte("HTTP/1.1 201 Created\r\n\r\n"))
		}
	} else {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}
}
