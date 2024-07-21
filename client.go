package main

import (
	"errors"
	"net"
	"os"
	"strconv"
    "fmt"
//    "time"
)

/*
func main() {

	var host, err = getHostFromArgs()
	// Connect to the server
	if err == nil {
		connectToNetwork(host)
	} else {
		fmt.Println("Error: ", err)
	}
}
*/

func connectToNetwork(host string, username string, password string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", host)
	println("tcpAddr: ", tcpAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

    conn, err := net.DialTCP("tcp", nil, tcpAddr)
	// conn, err := net.Dial("tcp", host)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}
    println("Sending mining.subscribe")
	message := "{\"id\": 1, \"method\": \"mining.subscribe\", \"params\": []}\n"
    sendMessage(message, conn)
    println("Sending mining.authorize")
    message = fmt.Sprintf("{\"params\": [\"%s\", \"%s\"], \"id\": 1, \"method\": \"mining.authorize\"}\n", username, password)
    go sendMessage(message, conn)
    println("Starting reading responses")
    continousProcessResponses(conn)

    defer conn.Close()
}

func sendMessage(message string, conn *(net.TCPConn)) {
    _, err := conn.Write([]byte(message))
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("write to server = ", message)	
}

func getResponse(conn *(net.TCPConn)) {
    reply := make([]byte, 10240)

    _, err := conn.Read(reply)
	if err != nil {
		println("Response from server failed:", err.Error())
		os.Exit(1)
	}

	println("reply from server=", string(reply))
}

func continousProcessResponses(conn *(net.TCPConn)) {
    for true {
        println("Getting Response")
        getResponse(conn)
    }
}

func runStratumTCPCommunicationJob(host string, port int, username string, password string) {
	connectToNetwork(getAddress(host, port), username, password)
}

func getAddress(host string, port int) string {
	stringPort := strconv.Itoa(port)
	return host + ":" + stringPort
}

func getHostFromArgs() (string, error) {
	var args = os.Args[1:]
	var ret = ""
	var err error = nil
	if len(args) == 2 {
		ret += args[0] + ":" + args[1]
	} else {
		err = errors.New("Wrong params")
	}
	return ret, err
}
