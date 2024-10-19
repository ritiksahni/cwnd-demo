package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func logWithTimestamp(message string) {
	fmt.Printf("%s - %s\n", time.Now().Format("2006-01-02 15:04:05"), message)
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		logWithTimestamp(fmt.Sprintf("Error connecting to server: %v", err))
		return
	}
	defer conn.Close()

	logWithTimestamp("Connected to server. Waiting for data...")

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			logWithTimestamp("Connection closed by server.")
			return
		}

		logWithTimestamp(fmt.Sprintf("Received: %s", message))

		_, err = fmt.Fprintf(conn, "ACK\n")
		if err != nil {
			logWithTimestamp(fmt.Sprintf("Error sending ACK: %v", err))
			return
		}

		logWithTimestamp("Sent: ACK")

		time.Sleep(500 * time.Millisecond)
	}
}

