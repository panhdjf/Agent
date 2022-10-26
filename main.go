package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	uptime time.Time
}

func UptimeServer() string {
	readFile, err := os.Open("/proc/uptime")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()

	a := fileScanner.Text()
	s := strings.Split(a, " ")
	uptime := s[0]

	readFile.Close()

	return uptime
}

func main() {
	router := gin.Default()
	router.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok", "uptime": UptimeServer()})
	})
	router.Run()
}
