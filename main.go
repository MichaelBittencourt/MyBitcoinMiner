package main

import (
	"fmt"
	"os"
	"strconv"
)

var Version = "0.0.0"

func help() {
	fmt.Println("Usage:")
	fmt.Printf("    %s [ help | version | host | port | username | password ]\n", os.Args[0])
	fmt.Println("\t -h | help\tShow help message")
	fmt.Println("\t -v | version\tShow current version")
	fmt.Println("\t -H | host\t\tSet stratum+tcp host")
	fmt.Println("\t -p | port\t\tSet server port")
	fmt.Println("\t -u | username\t\tSet wallet username")
	fmt.Println("\t -P | password\t\tSet wallet password")
	fmt.Println()
	fmt.Println("    Eg:")
	fmt.Printf("\t%s help\n", os.Args[0])
	fmt.Printf("\t%s -h\n", os.Args[0])
	fmt.Printf("\t%s version\n", os.Args[0])
	fmt.Printf("\t%s -v\n", os.Args[0])
	fmt.Printf("\t%s host sha256.auto.nicehash.com port 9200 username 31kA8ZZBBBCb57ScxyE9QopkQKinXS4Wuk password x\n", os.Args[0])
	fmt.Printf("\t%s -H sha256.auto.nicehash.com -p 9200 -u 31kA8ZZBBBCb57ScxyE9QopkQKinXS4Wuk -P x\n", os.Args[0])
}

func version() {
	fmt.Printf("Go Test Log Version: %s\n", Version)
}

func invalidParam(param string) {
	fmt.Printf("Invalid Param: %s\n", param)
	fmt.Printf("Try: %s help\n", os.Args[0])
}

func main() {
	host := "sha256.auto.nicehash.com"
	port := 9200
	username := "31kA8ZZBBBCb57ScxyE9QopkQKinXS4Wuk"
	password := "x"
	argv := os.Args[1:]
	for i := 0; i < len(argv); i++ {
		var paramError error = nil
		switch argv[i] {
		case "version", "-v":
			version()
			os.Exit(0)
		case "help", "-h":
			help()
			os.Exit(0)
		case "host", "-H":
			i++
			host = argv[i]
		case "port", "-p":
			i++
			port, paramError = strconv.Atoi(argv[i])
			if paramError != nil {
				invalidParam(argv[i])
				os.Exit(1)
			}
		case "username", "-u":
			i++
			username = argv[i]
		case "password", "-P":
			i++
			password = argv[i]
		default:
			invalidParam(argv[i])
			os.Exit(1)
		}
	}
    fmt.Printf("Starting communication on [host, username, password]: [%s:%d, %s, %s]\n", host, port, username, password)
	runStratumTCPCommunicationJob(host, port, username, password)
}
