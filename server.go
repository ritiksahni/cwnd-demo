package main

import (
	"fmt"
	"net"
	"time"
)

const (
	initialCwnd = 1    // Initial congestion window size (packets)
	maxCwnd     = 10   // Maximum congestion window size
)

func logWithTimestamp(message string) {
	fmt.Printf("%s - %s\n", time.Now().Format("2006-01-02 15:04:05"), message)
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	cwnd := initialCwnd
	lossThreshold := 0.2  // Simulate 20% chance of packet loss
	packetCount := 0

	logWithTimestamp("New client connected, starting data transmission.")

	for {
		logWithTimestamp(fmt.Sprintf("Current congestion window (cwnd): %d", cwnd))
		for i := 0; i < cwnd; i++ {
			packetCount++
			// Write data (packet) to the client
			_, err := fmt.Fprintf(conn, "Packet %d\n", packetCount)
			if err != nil {
				logWithTimestamp("Connection closed by client.")
				return
			}
			logWithTimestamp(fmt.Sprintf("Sent: Packet %d", packetCount))
		}

		time.Sleep(1 * time.Second)

		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			logWithTimestamp("Connection lost.")
			return
		}
		logWithTimestamp("Received ACK from client.")

		// Randomly simulate packet loss
		if simulateLoss(lossThreshold) {
			logWithTimestamp("Packet loss detected! Reducing congestion window.")
			cwnd = max(1, cwnd/2)  // On loss, halve the cwnd
		} else {
			// On success, increase cwnd exponentially until maxCwnd
			if cwnd < maxCwnd {
				cwnd *= 2
				if cwnd > maxCwnd {
					cwnd = maxCwnd
				}
			}
			logWithTimestamp(fmt.Sprintf("ACK successful, increasing cwnd. New cwnd = %d", cwnd))
		}

		logWithTimestamp("---- End of RTT cycle ----")
	}
}

func simulateLoss(threshold float64) bool {
	return float64(time.Now().UnixNano()%100) < threshold*100
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		logWithTimestamp(fmt.Sprintf("Error starting server: %v", err))
		return
	}
	defer listener.Close()

	logWithTimestamp("Server is listening on localhost:8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			logWithTimestamp(fmt.Sprintf("Error accepting connection: %v", err))
			continue
		}
		go handleConnection(conn)
	}
}

