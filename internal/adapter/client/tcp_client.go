package client

import (
	"connectorapi-go/internal/adapter/utils"
	"encoding/binary" // For length prefixing
	"fmt"             // For error formatting
	"io"
	"net"  // For TCP connections
	"time" // For timeouts
)

// TCPSocketClient defines the interface for a TCP socket client.
type TCPSocketClient interface {
	SendAndReceive(address string, requestPayloadUTF8 string) (string, error)
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
func (c *BasicTCPSocketClient) SendAndReceive(address string, requestPayloadUTF8 string) (string, error) {

	conn, err := net.DialTimeout("tcp", address, c.DialTimeout)
	if err != nil {
		return "", fmt.Errorf("failed to dial TCP address %s: %w", address, err)
	}
	defer conn.Close()

	requestPayload874, err := utils.Utf8ToCP874(requestPayloadUTF8)
	if err != nil {
		return "", fmt.Errorf("failed to encode request to CP874: %w", err)
	}

	length := uint32(len(requestPayload874))
	lengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes, length)

	encodedRequest := append(lengthBytes, requestPayload874...)

	if err := conn.SetWriteDeadline(time.Now().Add(c.ReadWriteTimeout)); err != nil {
		return "", fmt.Errorf("failed to set write deadline: %w", err)
	}

	_, err = conn.Write(encodedRequest)
	if err != nil {
		return "", fmt.Errorf("failed to write data to TCP connection: %w", err)
	}

	if err := conn.SetReadDeadline(time.Now().Add(c.ReadWriteTimeout)); err != nil {
		return "", fmt.Errorf("failed to set read deadline: %w", err)
	}

	responseBytes := make([]byte, 4096)
	n, err := conn.Read(responseBytes)
	if err != nil {
		if err == io.EOF {
			return "", fmt.Errorf("connection closed unexpectedly: %w", err)
		}
		return "", fmt.Errorf("connection closed unexpectedly: %w", err)
	}
	responseBytes = responseBytes[:n]

	// Decode from CP874 to UTF-8
	decoded, err := utils.DecodeCP874(responseBytes)
	if err != nil {
		return "", fmt.Errorf("failed to set read deadline: %w", err)
	}

	return decoded, nil
}
