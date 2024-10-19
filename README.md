# TCP Slow Start Simulation in Go

This repository contains a simple simulation of TCP's **slow-start** congestion control mechanism, implemented in Go. It demonstrates how the congestion window (cwnd) grows exponentially during the slow-start phase and how the server reacts to packet loss by reducing the window size.

This is inspired by [https://hpbn.co/building-blocks-of-tcp/#slow-start](https://hpbn.co/building-blocks-of-tcp/#slow-start)

It's a brilliant piece on the building blocks of TCP and how "slow start" works.

## Overview

### Key Concepts:
- **Slow Start**: TCP starts with a small congestion window (cwnd) and exponentially increases it after every successful acknowledgment (ACK), until a packet loss occurs.
- **Congestion Window (cwnd)**: Controls the number of packets the server can send without receiving an ACK. Starts small and grows during the slow-start phase.
- **Packet Loss**: Simulated with a probability threshold, causing the congestion window to reduce by half.

The server sends data packets based on the current congestion window size and waits for ACKs from the client. On receiving ACKs, the congestion window increases. If a packet loss is detected, the window size is reduced.

### Features:
- Logs with timestamps for better visualization of the slow-start process and packet loss handling.
- Simulated round-trip times (RTT) and network delays.
- A simple random loss simulator to emulate packet drops in a network.

## Usage

### 1. Run the Server
Start the server, which listens for incoming client connections and begins sending packets while applying TCP slow-start behavior.

```bash
go run server.go
```

### 2. Run the Client
The client connects to the server and receives data while sending ACKs back for each batch of packets received.

```bash
go run client.go
```

### 3. Output
Both server and client log events such as packet transmissions, ACKs, congestion window updates, and packet loss. Logs include timestamps to show the progression of each round-trip time (RTT) cycle.

## Example Logs

**Server:**
```
2024-10-19 12:30:01 - Server is listening on localhost:8080
2024-10-19 12:30:05 - Current congestion window (cwnd): 1
2024-10-19 12:30:06 - Sent: Packet 1
2024-10-19 12:30:06 - Received ACK from client.
2024-10-19 12:30:06 - ACK successful, increasing cwnd. New cwnd = 2
...
```

**Client:**
```
2024-10-19 12:30:05 - Connected to server. Waiting for data...
2024-10-19 12:30:05 - Received: Packet 1
2024-10-19 12:30:05 - Sent: ACK
...
```
