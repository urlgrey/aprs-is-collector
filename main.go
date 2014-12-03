package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

type RawAprsPacket struct {
	Data   string `json:"data"`
	IsAX25 bool   `json:"is_ax25"`
}

func main() {
	apiKey := os.Getenv("APRS_DASHBOARD_API_KEY")
	aprsDashHost := os.Getenv("APRS_DASHBOARD_HOST")
	if aprsDashHost == "" {
		log.Fatal("Set the APRS_DASHBOARD_HOST environment variable")
	}
	aprsDashUrl := "http://" + aprsDashHost + "/api/v1/message"

	conn, _ := net.Dial("tcp", "rotate.aprs.net:23")
	fmt.Fprintf(conn, "user N0CALL pass -1 vers goAPRS 0.00\r\n")

	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		sendPacketToDashboard(aprsDashUrl, apiKey, line)
	}
}

func sendPacketToDashboard(aprsDashUrl string, apiKey string, line string) {
	log.Println("Sending: ", line)
	packet := RawAprsPacket{Data: line, IsAX25: false}
	body, _ := json.Marshal(packet)

	req, err := http.NewRequest("PUT", aprsDashUrl, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if apiKey != "" {
		req.Header.Set("X-API-KEY", apiKey)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
}
