package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var session = make(map[string]*Session)

type Session struct {
	Conn net.Conn
	ID string
}

func main(){
	startServer()
	startClient()

}

func beacon(){

}

func startServer(){
	//open listener
	//make beacon
	beacon()
}