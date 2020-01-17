package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"
)

func main() {
	fmt.Println("TCP D20 listing on port 8889")
	ln, _ := net.Listen("tcp", ":8889")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("error creating TCP socket connection")
			fmt.Println(err.Error())
			panic("exiting...")
		}
		fmt.Println("Connection established...")
		conn.Write(getGreeting())
		go handleConnection(conn)
	}
}

func getGreeting() []byte {
	greeting := make([]byte, 1024)
	copy(greeting[0:29], "********* Welcome! ********* \n")
	copy(greeting[30:43], "How it works: \n")
	copy(greeting[44:], "Send us any bytes via TCP and we'll send back your D20 dice roll resultst \n")
	return greeting
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		data := make([]byte, 1024)
		_, err := conn.Read(data)

		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Disconnected.")
			break
		}

		conn.Write(getD20Roll())
	}
}

func getD20Roll() []byte {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(20) + 1
	result := []byte("You rolled:   ")
	copy(result[12:], strconv.Itoa(num))
	return result
}
