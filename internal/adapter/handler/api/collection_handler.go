package client

import (
	"connectorapi-go/internal/adapter/utils"
	"fmt"             
	"io"
	"net"  // For TCP connections
	"time" // For timeouts
)

// TCPSocketClient defines the interface for a TCP socket client.
type TCPSocketClient interface {
	SendAndReceive(address string, combinedPayloadString string) (string, error)
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
/*func (c *BasicTCPSocketClient) SendAndReceive(address string, requestPayloadUTF8 string) (string, error) {
	
	conn, err := net.DialTimeout("tcp", address, c.DialTimeout)
	if err != nil {
		return "", fmt.Errorf("failed to dial TCP address %s: %w", address, err)
	}
	defer conn.Close()

	requestPayload874, err := utils.Utf8ToCP874(requestPayloadUTF8)
	if err != nil {
		return "", fmt.Errorf("failed to encode request to CP874: %w", err)
	}

	fmt.Printf("Sending TCP request payload 874 : % X\n", requestPayload874)

	//length := uint32(len(requestPayload874))
	//lengthBytes := make([]byte, 4)
	//binary.BigEndian.PutUint32(lengthBytes, length)

	//encodedRequest := append(lengthBytes, requestPayload874...)

	//if err := conn.SetWriteDeadline(time.Now().Add(c.ReadWriteTimeout)); err != nil {
	//	return "", fmt.Errorf("failed to set write deadline: %w", err)
	//}

	_, err = conn.Write(requestPayload874)
	if err != nil {
		return "", fmt.Errorf("failed to write data to TCP connection: %w", err)
	}

	//if err := conn.SetReadDeadline(time.Now().Add(c.ReadWriteTimeout)); err != nil {
	//	return "", fmt.Errorf("failed to set read deadline: %w", err)
	//}

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

	fmt.Printf("Receiving TCP reponse payload UTF8 : % X\n", decoded)

	return decoded, nil
}*/

func (c *BasicTCPSocketClient) SendAndReceive(address string, combinedPayloadString string) (string, error) {

	//fmt.Println("Param -> requestPayloadUTF8 :",combinedPayloadString)

	// requestMsg := "AEON_WF   INQ_CUST_COSINF001RQ2025081820064204782025081801064200021                                                        4090610164364407DMEDฉันลองยิงเอพีไอมาน่ะ                                                                                                    ฉันลองยิงเอพีไอมาน่ะ2                                                                                                   ฉันลองยิงเอพีไอมาน่ะ3                                                                                                   ฉันลองยิงเอพีไอมาน่ะ4                                                                                                   ฉันลองยิงเอพีไอมาน่ะ5                                                                                                   20250820122600GSYS1234567891"

	fmt.Println("Connecting to the server...")
	conn, err := net.DialTimeout("tcp", address, c.DialTimeout)
	if err != nil {
		return "", fmt.Errorf("ER040: " + err.Error())
	}
	defer conn.Close()

	conn.SetDeadline(time.Now().Add(c.DialTimeout))

	fmt.Println("Encoding request to CP874...")
	encodedRequest, err := utils.Utf8ToCP874(combinedPayloadString)
	if err != nil {
		return "", fmt.Errorf("ER040: Failed to encode request to CP874: " + err.Error())
	}

	//fmt.Printf("Sending bytes: % X\n", encodedRequest)
	_, err = conn.Write(encodedRequest)
	if err != nil {
		return "", fmt.Errorf("ER060: " + err.Error())
	}

	fmt.Println("Request sent (UTF-8):", combinedPayloadString)

	fmt.Println("Reading response...")
	responseBytes := make([]byte, 4096)
	n, err := conn.Read(responseBytes)
	if err != nil {
		if err == io.EOF {
			return "", fmt.Errorf("ER060: Connection closed unexpectedly")
		}
		return "", fmt.Errorf("ER060: " + err.Error())
	}
	responseBytes = responseBytes[:n]

	// Decode from CP874 to UTF-8
	decoded, err := utils.DecodeCP874(responseBytes)
	if err != nil {
		return "", fmt.Errorf("ER040: Failed to decode response from CP874: " + err.Error())
	}

	fmt.Println("Final result (UTF-8):", decoded)
	return decoded, nil
}
