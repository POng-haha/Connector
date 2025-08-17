package client

import (
	"encoding/binary" // For length prefixing
	"fmt"             // For error formatting
	"io"
	"net"  // For TCP connections
	"time" // For timeouts

	tis620 "connectorapi-go/pkg/tis620"
)

// TCPSocketClient defines the interface for a TCP socket client.
type TCPSocketClient interface {
	SendAndReceive(address string, requestPayloadUTF8 []byte) ([]byte, error)
}

// BasicTCPSocketClient implements TCPSocketClient with length-prefixing and TIS-620 encoding.
type BasicTCPSocketClient struct {
	DialTimeout      time.Duration // Timeout for establishing the connection
	ReadWriteTimeout time.Duration // Timeout for read/write operations
}

// NewBasicTCPSocketClient creates a new instance of BasicTCPSocketClient.
func NewBasicTCPSocketClient(dialTimeout, readWriteTimeout time.Duration) *BasicTCPSocketClient {
	return &BasicTCPSocketClient{
		DialTimeout:      dialTimeout,
		ReadWriteTimeout: readWriteTimeout,
	}
}

// SendAndReceive connects to a TCP server, sends length-prefixed data, and receives a response.
func (c *BasicTCPSocketClient) SendAndReceive(address string, requestPayloadUTF8 []byte) ([]byte, error) {
	// 1. Establish TCP connection with a dial timeout.
	conn, err := net.DialTimeout("tcp", address, c.DialTimeout)
	if err != nil {
		return nil, fmt.Errorf("failed to dial TCP address %s: %w", address, err)
	}
	defer conn.Close()

	// --- Encoding: Convert requestPayload (UTF-8 []byte) to TIS-620 bytes using tis620.ToTIS620 ---
	requestPayloadTIS620 := tis620.ToTIS620(requestPayloadUTF8)

	// 2. Prepare the request payload with a 4-byte length prefix (BigEndian).
	length := uint32(len(requestPayloadTIS620))
	lengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes, length)

	// Combine length prefix and actual payload.
	fullRequest := append(lengthBytes, requestPayloadTIS620...) // Send TIS-620 bytes

	// 3. Set a write deadline to prevent hanging indefinitely.
	if err := conn.SetWriteDeadline(time.Now().Add(c.ReadWriteTimeout)); err != nil {
		return nil, fmt.Errorf("failed to set write deadline: %w", err)
	}

	// 4. Write the full request (length + payload) to the connection.
	_, err = conn.Write(fullRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to write data to TCP connection: %w", err)
	}

	// 5. Set a read deadline for reading the response.
	if err := conn.SetReadDeadline(time.Now().Add(c.ReadWriteTimeout)); err != nil {
		return nil, fmt.Errorf("failed to set read deadline: %w", err)
	}

	// 6. Read the 4-byte length prefix of the response first.
	responseLengthBytes := make([]byte, 4)
	_, err = io.ReadFull(conn, responseLengthBytes) // ReadFull ensures all 4 bytes are read.
	if err != nil {
		return nil, fmt.Errorf("failed to read response length prefix: %w", err)
	}
	responseLength := binary.BigEndian.Uint32(responseLengthBytes)

	// 7. Read the actual response payload based on the length received.
	responsePayloadTIS620 := make([]byte, responseLength)
	_, err = io.ReadFull(conn, responsePayloadTIS620) // ReadFull ensures all bytes are read.
	if err != nil {
		return nil, fmt.Errorf("failed to read response payload: %w", err)
	}

	responsePayloadUTF8 := tis620.ToUTF8(responsePayloadTIS620)

	return responsePayloadUTF8, nil
}
