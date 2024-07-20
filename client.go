package main

import (
	"errors"
	"net"
	"os"
	"strconv"
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

func connectToNetwork(host string) {
	strEcho := "{\"id\": 1, \"method\": \"mining.subscribe\", \"params\": []}\n"
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

	_, err = conn.Write([]byte(strEcho))
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("write to server = ", strEcho)

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("reply from server=", string(reply))

	defer conn.Close()
}

func runStratumTCPCommunicationJob(host string, port int) {
	connectToNetwork(getAddress(host, port))
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
