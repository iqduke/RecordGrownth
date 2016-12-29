package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":7777"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		fmt.Println("1")
		conn, err := listener.Accept()
		fmt.Println(conn.RemoteAddr().String())
		if err != nil {
			continue
		}

		go handleClient(conn)
	}
}
func handleClient(conn net.Conn) {
	defer conn.Close()
	daytime := time.Now().String()
	conn.Write([]byte(daytime)) // don't care about return value
	// we're finished with this client
}

// func handleClient(conn net.Conn) {
// 	fmt.Println(conn.RemoteAddr().String())

// 	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
// 	request := make([]byte, 128)
// 	defer conn.Close()

// 	for {
// 		read_len, err := conn.Read(request)
// 		fmt.Println(string(request))
// 		if err != nil {
// 			fmt.Println(err)
// 			break
// 		}

// 		fmt.Println(read_len)
// 		if read_len == 0 {
// 			break
// 		} else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
// 			daytime := strconv.FormatInt(time.Now().Unix(), 10)
// 			conn.Write([]byte(daytime))
// 		} else {
// 			daytime := time.Now().String()
// 			fmt.Println(daytime)
// 			conn.Write([]byte(daytime))
// 		}

// 		request = make([]byte, 128)
// 	}

// }

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
