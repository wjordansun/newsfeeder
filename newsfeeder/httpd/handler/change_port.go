package handler

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/gin-gonic/gin"
)

var (
	ipAddress string = "192.168.0.229:8000"
)

func sendChangePortMessage() {
	fmt.Println("test")
	conn, err := net.Dial("tcp", ipAddress)

	if err != nil {
		log.Panic(err)
	}

	b := []byte("changeport")

	_, err = io.Copy(conn, bytes.NewReader(b))
	if err != nil {
		log.Panic(err)
	}
}

func ChangePort() gin.HandlerFunc {
	return func(c *gin.Context) {
		sendChangePortMessage()
	}
}